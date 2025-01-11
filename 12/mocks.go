package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperDefault struct {
	duration time.Duration
	pause    func(time.Duration)
}

func (s SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *SleeperDefault) Pause() {
	c.pause(c.duration)
}

func Counter(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, "Go!")
}

func main() {
	sleeper := &SleeperDefault{1 * time.Second, time.Sleep}
	Counter(os.Stdout, sleeper)
}
