package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

// Sleep SpySleeper implements the Sleep method of the Sleeper interface
func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep ConfigurableSleeper implements the Sleep method of the Sleeper interface
func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, finalWord)
}
