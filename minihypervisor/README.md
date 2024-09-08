## 注意

自分の手元のFedoraマシンで動作確認しただけ。

また、元ネタはこちら。

https://github.com/oreilly-japan/binary-hacks-rebooted/tree/main/ch03_os/30_kvm

これをなんとかGoで書き換えただけ。

あと、すでにgokvmというライブラリを作っている方がいて、Ioctlのラップであまり考える必要がなくて助かった。

https://github.com/bobuhiro11/gokvm/blob/main/kvm/ioctl.go

## 使い方

```
go mod tidy
go build .
make guest
./kvmgo countdown.bin
```

Enter Number:　と出てくるので適当に入れて実行