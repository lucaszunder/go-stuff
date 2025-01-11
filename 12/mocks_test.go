package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

type TimeSpy struct {
	duration time.Duration
}

func (t *TimeSpy) pause(duration time.Duration) {
	t.duration = duration
}

const Printer = "print"
const Hold = "sleep"

type Actions struct {
	operations []string
}

func (a *Actions) Write(p []byte) (n int, err error) {
	a.operations = append(a.operations, Printer)
	return
}

func (a *Actions) Sleep() {
	a.operations = append(a.operations, Hold)
}

func TestCounter(t *testing.T) {
	t.Run("Testing if everything run", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}
		Counter(&buffer, sleeperSpy)

		output := buffer.String()
		expectation := `3
2
1
Go!`

		if output != expectation {
			t.Errorf("Output '%s', expectation '%s'", output, expectation)
		}

		if sleeperSpy.Calls != 4 {
			t.Errorf("There isn't enough calls, expected 4, output %d", sleeperSpy.Calls)
		}
	})

	t.Run("Testing if everything run in the correct order", func(t *testing.T) {
		recorder := &Actions{}
		Counter(recorder, recorder)

		expectation := []string{
			Hold,
			Printer,
			Hold,
			Printer,
			Hold,
			Printer,
			Hold,
			Printer,
		}
		if !reflect.DeepEqual(recorder.operations, expectation) {
			t.Errorf("expectation %v calls, output %v", expectation, recorder.operations)
		}
	})
}

func TestSleeper(t *testing.T) {
	duration := 5 * time.Second

	durationSpy := &TimeSpy{}
	sleeper := SleeperDefault{duration, durationSpy.pause}
	sleeper.Pause()

	if durationSpy.duration != duration {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", duration, durationSpy.duration)
	}
}
