package main

import (
	"reflect"
	"testing"
)

func TestIterate(t *testing.T) {
	t.Run("Validate if the function has the correct number of results when is one", func(t *testing.T) {
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

	t.Run("Validate if the function has the correct number of results", func(t *testing.T) {
		cases := []struct {
			Name          string
			Entry         interface{}
			NumberOfCalls []string
		}{
			{
				"Struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var output []string
				Iterate(test.Entry, func(entry string) {
					output = append(output, entry)
				})

				if !reflect.DeepEqual(output, test.NumberOfCalls) {
					t.Errorf("output %v, expectation %v", output, test.NumberOfCalls)
				}
			})
		}
	})

	t.Run("when one of the values is not one string", func(t *testing.T) {
		cases := []struct {
			Name          string
			Entry         interface{}
			NumberOfCalls []string
		}{
			{
				"Struct with two string fields",
				struct {
					Name string
					Age  int
				}{"Chris", 13},
				[]string{"Chris"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var output []string
				Iterate(test.Entry, func(entry string) {
					output = append(output, entry)
				})

				if !reflect.DeepEqual(output, test.NumberOfCalls) {
					t.Errorf("output %v, expectation %v", output, test.NumberOfCalls)
				}
			})
		}
	})

	t.Run("when one of the values is a nested value", func(t *testing.T) {
		cases := []struct {
			Name          string
			Entry         interface{}
			NumberOfCalls []string
		}{
			{
				"Struct with two string fields",
				struct {
					Name    string
					Address struct {
						City   string
						Street string
					}
				}{"Chris", struct {
					City   string
					Street string
				}{"New York", "3000 ST"}},
				[]string{"Chris"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var output []string
				Iterate(test.Entry, func(entry string) {
					output = append(output, entry)
				})

				if !reflect.DeepEqual(output, test.NumberOfCalls) {
					t.Errorf("output %v, expectation %v", output, test.NumberOfCalls)
				}
			})
		}
	})
}
