package main

import (
	"testing"
	"runtime"
)

// 2つのgoroutine間でチャネルを使って「ピンポン」させ、
// 強制的にコンテキストスイッチを発生させる
func BenchmarkGoroutineContextSwitch(b *testing.B) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		// b.N：人間が考えなくていいように適切な値が取れるまで実行する回数
		for i := 0; i < b.N; i++ {
			ch1 <- struct{}{} // 送信してブロック（スイッチ発生）
			<-ch2             // 受信待ちでブロック（スイッチ発生）
		}
	}()

	// 準備にかかった時間は含めずこの先のループ時間だけを計測するという意味。
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-ch1             // 受信待ち
		ch2 <- struct{}{} // 送信
	}
}

// OSスレッド同士のコンテキストスイッチを計測
func BenchmarkOSThreadContextSwitch(b *testing.B) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// 1つ目のスレッドを固定して起動
	go func() {
		// 本来はgoruntimeがOSスレッドを複数保持してその中でgoroutineを管理して割り当てている
		runtime.LockOSThread() // が、goroutineをOSスレッドに固定できるやつ
		defer runtime.UnlockOSThread()
		for i := 0; i < b.N; i++ {
			ch1 <- struct{}{}
			<-ch2
		}
	}()

	// メインのベンチマークgoroutineも別のOSスレッドに固定
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-ch1
		ch2 <- struct{}{}
	}
}