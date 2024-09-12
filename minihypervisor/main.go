package main

import(
	"fmt"
	"syscall"
	"os"
	"unsafe"
	"errors"
	// "encoding/binary"
	// "io"
	// "log"
	"golang.org/x/sys/unix"
)

const (
	kvmGetAPIVersion     = 0x00
	kvmCreateVM          = 0x1
	kvmGetVCPUMMapSize   = 0x04
	kvmCreateVCPU          = 0x41
	kvmSetUserMemoryRegion = 0x46
	kvmRun       = 0x80
	kvmGetRegs   = 0x81
	kvmSetRegs   = 0x82
	kvmGetSregs  = 0x83
	kvmSetSregs  = 0x84
	numInterrupts   = 0x100
)

type UserspaceMemoryRegion struct {
	Slot          uint32
	Flags         uint32
	GuestPhysAddr uint64
	MemorySize    uint64
	UserspaceAddr uint64
}

func main() {
	fmt.Println("Hello, World!")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <binary file>\n", os.Args[0])
		os.Exit(1)
	}

	binary := os.Args[1]

	kvm, err := os.OpenFile("/dev/kvm", unix.O_RDWR, 0o644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer kvm.Close()
	
	// api version
	apiVersion, err := GetAPIVersion(kvm.Fd())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("API Version: ", apiVersion)

	// create vm
	vmFd, err := CreateVM(kvm.Fd())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VM FD: ", vmFd)

	// assign memory
	mem, err := syscall.Mmap(-1, 0, 1024000, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED|syscall.MAP_ANONYMOUS|syscall.MAP_NORESERVE)
	if err != nil {
		fmt.Println(err)
		return
	}

	SetUserMemoryRegion(vmFd, &UserspaceMemoryRegion{
		Slot: 0, Flags: 0, GuestPhysAddr: 0, MemorySize: 1024000,
		UserspaceAddr: uint64(uintptr(unsafe.Pointer(&mem[0]))),
	})

	// load program to memory
	file, err := os.Open(binary)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fopen failed: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Read(mem)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fread failed: %v\n", err)
		os.Exit(1)
	}

	// create vcpu
	vcpuFd, err := CreateVCPU(vmFd, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("VCPU FD: ", vcpuFd)
	
	// vcpu memory map
	vcpuMmapSize, err := GetVCPUMMmapSize(kvm.Fd())
	if err != nil {
		fmt.Println(err)
		return
	}
	kvmRunPtr, err := syscall.Mmap(int(vcpuFd), 0, int(vcpuMmapSize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("VCPU MMAP SIZE: ", vcpuMmapSize)

	// initialize global registers
	regs, err := GetRegs(vcpuFd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("REGS: ", regs)
	regs.RFLAGS = 1 << 1
	regs.RIP = 0x0
	regs.RSP = 0xffffffff
	regs.RBP = 0

	err = SetRegs(vcpuFd, regs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Initial Registers: RIP=0x%x, RSP=0x%x, RFLAGS=0x%x\n", regs.RIP, regs.RSP, regs.RFLAGS)

	// initialize segment registers
	sregs, err := GetSregs(vcpuFd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SREGS: ", sregs)
	sregs.CS.Base = 0
	sregs.CS.Selector = 0
	err = SetSregs(vcpuFd, sregs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Initial Segment Registers: CS.Base=0x%x, CS.Selector=0x%x\n", sregs.CS.Base, sregs.CS.Selector)

	// run vm
	for {
		if err = Run(vcpuFd); err != nil {
			if errors.Is(err, syscall.EAGAIN) || errors.Is(err, syscall.EINTR) {
				continue
			}
			fmt.Println(err)
			return
		}

		run := (*RunData)(unsafe.Pointer(&kvmRunPtr[0]))

		regsAfterRun, err := GetRegs(vcpuFd)
		if err != nil {
			fmt.Println("Error fetching regs after KVM_RUN:", err)
		} else {
			fmt.Printf("Registers after KVM_RUN: RIP=0x%x, RSP=0x%x, RAX=0x%x\n", regsAfterRun.RIP, regsAfterRun.RSP, regsAfterRun.RAX)
		}


		switch ExitType(run.ExitReason) {
		case EXITHLT:
			fmt.Println("KVM_EXIT_HLT")

		case EXITIO:
			direction, size, port, count, offset := run.IO()
			fmt.Printf("KVM_EXIT_IO: direction=%d, size=%d, port=%d, count=%d, offset=%d\n", direction, size, port, count, offset)
			if direction == EXITIOIN {
				fmt.Printf("Enter a number: ")
				var input int
				fmt.Scanf("%d", &input)
				mem[offset] = byte(input)
			} else {
				fmt.Printf("Output: %d\n", mem[offset])
			}
		case EXITDCR,
			EXITINTR,
			EXITEXCEPTION,
			EXITFAILENTRY,
			EXITHYPERCALL,
			EXITINTERNALERROR,
			EXITIRQWINDOWOPEN,
			EXITMMIO,
			EXITNMI,
			EXITS390RESET,
			EXITS390SIEIC,
			EXITSETTPR,
			EXITSHUTDOWN,
			EXITTPRACCESS,
			EXITUNKNOWN:
			fmt.Printf("KVM_RUN exit reason: 0x%x\n", run.ExitReason)
			return
		default:
			fmt.Printf("KVM_RUN exit reason: 0x%x\n", run.ExitReason)
			return
		}
	}

}

const (
	nrbits   = 8
	typebits = 8
	sizebits = 14
	dirbits  = 2

	nrmask   = (1 << nrbits) - 1
	sizemask = (1 << sizebits) - 1
	dirmask  = (1 << dirbits) - 1

	none      = 0
	write     = 1
	read      = 2
	readwrite = 3

	nrshift   = 0
	typeshift = nrshift + nrbits
	sizeshift = typeshift + typebits
	dirshift  = sizeshift + sizebits
)

// KVMIO is for the KVMIO ioctl.
const KVMIO = 0xAE

// IIOWR creates an IIOWR ioctl.
func IIOWR(nr, size uintptr) uintptr {
	return IIOC(readwrite, nr, size)
}

// IIOR creates an IIOR ioctl.
// 読み取り専用のioctlコマンド
func IIOR(nr, size uintptr) uintptr {
	return IIOC(read, nr, size)
}

// IIOW creates an IIOW ioctl.
// 書き込み専用のioctlコマンド
func IIOW(nr, size uintptr) uintptr {
	return IIOC(write, nr, size)
}

// IIO creates an IIOC ioctl from a number.
// データ転送を行わないicotlコマンド
func IIO(nr uintptr) uintptr {
	return IIOC(none, nr, 0)
}

// IIOC creates an IIOC ioctl from a direction, nr, and size.
func IIOC(dir, nr, size uintptr) uintptr {
	// This is another case of forced wrapping which is considered an anti-pattern in Google.
	return ((dir & dirmask) << dirshift) | (KVMIO << typeshift) |
		((nr & nrmask) << nrshift) | ((size & sizemask) << sizeshift)
}

func Ioctl(fd, op, arg uintptr) (uintptr, error) {
	res, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, op, arg)
	if errno != 0 {
		fmt.Println("errno: ", errno)
		return res, errno
	}

	return res, nil
}

func GetAPIVersion(kvmFd uintptr) (uintptr, error) {
	return Ioctl(kvmFd, IIO(kvmGetAPIVersion), uintptr(0))
}

func CreateVM(kvmFd uintptr) (uintptr, error) {
	return Ioctl(kvmFd, IIO(kvmCreateVM), uintptr(0))
}

func SetUserMemoryRegion(vmFd uintptr, region *UserspaceMemoryRegion) error {
	_, err := Ioctl(vmFd, IIOW(kvmSetUserMemoryRegion, unsafe.Sizeof(UserspaceMemoryRegion{})),
		uintptr(unsafe.Pointer(region)))

	fmt.Println("SetUserMemoryRegion: ", err)
	return err
}

func CreateVCPU(vmFd uintptr, vcpuID int) (uintptr, error) {
	return Ioctl(vmFd, IIO(kvmCreateVCPU), uintptr(vcpuID))
}

func GetVCPUMMmapSize(kvmFd uintptr) (uintptr, error) {
	return Ioctl(kvmFd, IIO(kvmGetVCPUMMapSize), uintptr(0))
}

func Run(vcpuFd uintptr) error {
	_, err := Ioctl(vcpuFd, IIO(kvmRun), uintptr(0))
	if err != nil {
		// refs: https://github.com/kvmtool/kvmtool/blob/415f92c33a227c02f6719d4594af6fad10f07abf/kvm-cpu.c#L44
		if errors.Is(err, syscall.EAGAIN) || errors.Is(err, syscall.EINTR) {
			return nil
		}
	}
	fmt.Println("Run: ", err)

	return err
}

// Segment is an x86 segment descriptor.
type Segment struct {
	Base     uint64
	Limit    uint32
	Selector uint16
	Typ      uint8
	Present  uint8
	DPL      uint8
	DB       uint8
	S        uint8
	L        uint8
	G        uint8
	AVL      uint8
	Unusable uint8
	_        uint8
}

type Regs struct {
	RAX    uint64
	RBX    uint64
	RCX    uint64
	RDX    uint64
	RSI    uint64
	RDI    uint64
	RSP    uint64
	RBP    uint64
	R8     uint64
	R9     uint64
	R10    uint64
	R11    uint64
	R12    uint64
	R13    uint64
	R14    uint64
	R15    uint64
	RIP    uint64
	RFLAGS uint64
}

// GetRegs gets the general purpose registers for a vcpu.
func GetRegs(vcpuFd uintptr) (*Regs, error) {
	regs := &Regs{}
	_, err := Ioctl(vcpuFd, IIOR(kvmGetRegs, unsafe.Sizeof(Regs{})), uintptr(unsafe.Pointer(regs)))

	return regs, err
}

// SetRegs sets the general purpose registers for a vcpu.
func SetRegs(vcpuFd uintptr, regs *Regs) error {
	_, err := Ioctl(vcpuFd, IIOW(kvmSetRegs, unsafe.Sizeof(Regs{})), uintptr(unsafe.Pointer(regs)))

	return err
}

type Descriptor struct {
	Base  uint64
	Limit uint16
	_     [3]uint16
}

// Sregs are control registers, for memory mapping for the most part.
type Sregs struct {
	CS              Segment
	DS              Segment
	ES              Segment
	FS              Segment
	GS              Segment
	SS              Segment
	TR              Segment
	LDT             Segment
	GDT             Descriptor
	IDT             Descriptor
	CR0             uint64
	CR2             uint64
	CR3             uint64
	CR4             uint64
	CR8             uint64
	EFER            uint64
	ApicBase        uint64
	InterruptBitmap [(numInterrupts + 63) / 64]uint64
}

// GetSRegs gets the special registers for a vcpu.
func GetSregs(vcpuFd uintptr) (*Sregs, error) {
	sregs := &Sregs{}
	_, err := Ioctl(vcpuFd, IIOR(kvmGetSregs, unsafe.Sizeof(Sregs{})), uintptr(unsafe.Pointer(sregs)))

	return sregs, err
}

// SetSRegs sets the special registers for a vcpu.
func SetSregs(vcpuFd uintptr, sregs *Sregs) error {
	_, err := Ioctl(vcpuFd, IIOW(kvmSetSregs, unsafe.Sizeof(Sregs{})), uintptr(unsafe.Pointer(sregs)))

	return err
}

type ExitType uint

const (
	EXITUNKNOWN       ExitType = 0
	EXITEXCEPTION     ExitType = 1
	EXITIO            ExitType = 2
	EXITHYPERCALL     ExitType = 3
	EXITDEBUG         ExitType = 4
	EXITHLT           ExitType = 5
	EXITMMIO          ExitType = 6
	EXITIRQWINDOWOPEN ExitType = 7
	EXITSHUTDOWN      ExitType = 8
	EXITFAILENTRY     ExitType = 9
	EXITINTR          ExitType = 10
	EXITSETTPR        ExitType = 11
	EXITTPRACCESS     ExitType = 12
	EXITS390SIEIC     ExitType = 13
	EXITS390RESET     ExitType = 14
	EXITDCR           ExitType = 15
	EXITNMI           ExitType = 16
	EXITINTERNALERROR ExitType = 17

	EXITIOIN  = 0
	EXITIOOUT = 1
)

// RunData defines the data used to run a VM.
type RunData struct {
	RequestInterruptWindow     uint8
	ImmediateExit              uint8
	_                          [6]uint8
	ExitReason                 uint32
	ReadyForInterruptInjection uint8
	IfFlag                     uint8
	_                          [2]uint8
	CR8                        uint64
	ApicBase                   uint64
	Data                       [32]uint64
}

// IO interprets IO requests from a VM, by unpacking RunData.Data[0:1].
func (r *RunData) IO() (uint64, uint64, uint64, uint64, uint64) {
	direction := r.Data[0] & 0xFF
	size := (r.Data[0] >> 8) & 0xFF
	port := (r.Data[0] >> 16) & 0xFFFF
	count := (r.Data[0] >> 32) & 0xFFFFFFFF
	offset := r.Data[1]

	return direction, size, port, count, offset
}