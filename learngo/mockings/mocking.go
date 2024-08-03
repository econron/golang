package mockings

import (
	"io"
	"fmt"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func(d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(buf io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buf, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(buf, finalWord)
}