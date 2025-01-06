package main

import "testing"

// requirements:
// 1 - if we call hello with a empty string, should return "Hello World",
// if we call it with a string should return "Hello STRING"
// 2 - The second parameter should determine the language, if one empty string should return in english

func TestHello(t *testing.T) {
	validateMessage := func(t *testing.T, expectation, result string) {
		t.Helper()
		if result != expectation {
			t.Errorf("result '%s', expected '%s'", result, expectation)
		}
	}

	t.Run("Say hello to people in portuguese", func(t *testing.T) {
		result := Hello("Teste", "portuguese")
		expectation := "Oi Teste"

		validateMessage(t, expectation, result)
	})

	t.Run("Say hello to people in spanish", func(t *testing.T) {
		result := Hello("Teste", "spanish")
		expectation := "Ola Teste"

		validateMessage(t, expectation, result)
	})

	t.Run("Say hello to people in english due the default language", func(t *testing.T) {
		result := Hello("Teste", "")
		expectation := "Hello Teste"

		validateMessage(t, expectation, result)
	})

	t.Run("Return the default message", func(t *testing.T) {
		result := Hello("", "")
		expectation := "Hello World"

		validateMessage(t, expectation, result)
	})

	t.Run("Return the default message in portuguese", func(t *testing.T) {
		result := Hello("", "portuguese")
		expectation := "Oi Mundo"

		validateMessage(t, expectation, result)
	})
}
