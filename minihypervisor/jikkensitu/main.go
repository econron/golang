package main

import "fmt"

// 定数の定義
const (
	nrbits   = 8
	typebits = 8
	sizebits = 14
	dirbits  = 2

	nrmask   = (1 << nrbits) - 1   // 255 (0b11111111)
	sizemask = (1 << sizebits) - 1 // 16383 (0b11111111111111)
	dirmask  = (1 << dirbits) - 1  // 3 (0b11)

	nrshift   = 0
	typeshift = nrshift + nrbits   // 8
	sizeshift = typeshift + typebits // 16
	dirshift  = sizeshift + sizebits // 30
)

func decodeCommand(cmd uintptr) (dir, nr, size, kvmio uintptr) {
	// 各フィールドを抽出するためにビットシフトとビットマスクを使う
	dir = (cmd >> dirshift) & dirmask    // ディレクション (読み書きの方向)
	size = (cmd >> sizeshift) & sizemask // データサイズ
	kvmio = (cmd >> typeshift) & nrmask  // KVMIO (デバイスの種類)
	nr = (cmd >> nrshift) & nrmask       // 操作番号

	return dir, nr, size, kvmio
}

func main() {
	// 例として 0xAE00 を分解する
	cmd := uintptr(0xAE00)
	dir, nr, size, kvmio := decodeCommand(cmd)

	fmt.Printf("Command: 0x%X\n", cmd)
	fmt.Printf("Direction (dir): %d\n", dir)
	fmt.Printf("Operation Number (nr): %d\n", nr)
	fmt.Printf("Size: %d\n", size)
	fmt.Printf("Device Type (KVMIO): 0x%X\n", kvmio)
}
