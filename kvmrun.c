#include <fcntl.h>
#include <linux/kvm.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/mman.h>
#include <unistd.h>

#define GUEST_CODE_START 0x1000
#define KVM_EXIT_IO_INSTRUCTION 0x80

void setup_kvm(int kvm_fd, int vm_fd, int vcpu_fd) {
    // 記憶用のKVM関連の設定

    // 割り込み用に使用される構造体
    struct kvm_interrupt kvm_irq;
    kvm_irq.irq = 0x01; // 割り込み番号 0x01

    // 割り込みを発行する
    if (ioctl(vcpu_fd, KVM_INTERRUPT, &kvm_irq) < 0) {
        perror("Failed to inject interrupt");
        exit(1);
    }
}

int main() {
    int kvm_fd, vm_fd, vcpu_fd;
    struct kvm_run *run;
    int vcpu_mmap_size;
    
    // KVMデバイスをオープン
    kvm_fd = open("/dev/kvm", O_RDWR | O_CLOEXEC);
    if (kvm_fd < 0) {
        perror("Failed to open /dev/kvm");
        exit(1);
    }

    // 仮想マシン（VM）を作成
    vm_fd = ioctl(kvm_fd, KVM_CREATE_VM, 0);
    if (vm_fd < 0) {
        perror("Failed to create VM");
        exit(1);
    }

    // VCPUの作成
    vcpu_fd = ioctl(vm_fd, KVM_CREATE_VCPU, 0);
    if (vcpu_fd < 0) {
        perror("Failed to create VCPU");
        exit(1);
    }

    // VCPUの実行情報のメモリマッピング
    vcpu_mmap_size = ioctl(kvm_fd, KVM_GET_VCPU_MMAP_SIZE, 0);
    if (vcpu_mmap_size < 0) {
        perror("Failed to get VCPU mmap size");
        exit(1);
    }

    run = mmap(NULL, vcpu_mmap_size, PROT_READ | PROT_WRITE, MAP_SHARED, vcpu_fd, 0);
    if (run == MAP_FAILED) {
        perror("Failed to mmap VCPU");
        exit(1);
    }

    // 仮想マシンを設定して割り込みを発生させる
    setup_kvm(kvm_fd, vm_fd, vcpu_fd);

    // VCPUを実行
    while (1) {
        // VCPUを実行する
        if (ioctl(vcpu_fd, KVM_RUN, 0) < 0) {
            perror("Failed to run VCPU");
            exit(1);
        }

        // 割り込みを処理する
        if (run->exit_reason == KVM_EXIT_IO) {
            // 割り込みの発生を検知し、適切にハンドリングする
            if (run->io.direction == KVM_EXIT_IO_INSTRUCTION) {
                printf("Handling IO instruction interrupt\n");
            }
        } else if (run->exit_reason == KVM_EXIT_INTR) {
            // 外部割り込みを処理する
            printf("Handling external interrupt\n");
        } else {
            printf("Unhandled exit reason: %d\n", run->exit_reason);
        }
    }

    // クリーンアップ
    munmap(run, vcpu_mmap_size);
    close(vcpu_fd);
    close(vm_fd);
    close(kvm_fd);

    return 0;
}
