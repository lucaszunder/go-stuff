package main

import "testing"

func TestIterate(t *testing.T) {
	t.Run("Validate if the function has the correct number of results", func(t *testing.T) {
		expectation := "test"
		var output []string

		x := struct {
			Name string
		}{expectation}

		Iterate(x, func(incoming string) {
			output = append(output, incoming)
		})

		if len(output) != 1 {
			t.Errorf("incorrect number of calls: output %d, expectation %d", len(output), 1)
		}

		if expectation != output[0] {
			t.Errorf("expectation %s | output %s", expectation, output[0])
		}
	})
}
