package di

import (
	"fmt"
	"io"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}