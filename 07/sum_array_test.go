package main

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	validateExpectations := func(t *testing.T, expectation, output int) {
		t.Helper()
		if !reflect.DeepEqual(output, expectation) {
			t.Errorf("The output for the method Sum is '%v' and the expectation was '%v'", output, expectation)
		}
	}

	t.Run("SumArray should sum all elements of an slice", func(t *testing.T) {
		output := SumArray([]int{1, 2, 3})
		expectation := 6

		validateExpectations(t, expectation, output)
	})

	t.Run("SumSlices should sum all slices of an slice", func(t *testing.T) {
		output := SumSlices([]int{1, 2, 3}, []int{1, 3, 3})
		expectation := []int{6, 7}

		if !reflect.DeepEqual(output, expectation) {
			t.Errorf("The output for the method Sum is '%v' and the expectation was '%v'", output, expectation)
		}
	})
}
