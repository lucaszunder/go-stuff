package main

import "testing"

//func TestSearch(t *testing.T) {
//	dictionary := map[string]string{"test": "This is a test"}
//
//	output := Search(dictionary, "test")
//	expectation := "This is a test"
//
//	if output != expectation {
//		t.Errorf("output `%s`, expected `%s`", output, expectation)
//	}
//}

func TestSearch(t *testing.T) {
	compareStrings := func(t *testing.T, output, expectation string) {
		t.Helper()
		if output != expectation {
			t.Errorf("output `%s`, expected `%s`", output, expectation)
		}
	}

	dictionary := Dictionary{"test": "This is a test"}

	t.Run("if we found a word", func(t *testing.T) {
		output, _ := dictionary.Search("test")
		expectation := "This is a test"
		compareStrings(t, output, expectation)
	})

	t.Run("if we didn't found a word", func(t *testing.T) {
		_, err := dictionary.Search("test1")

		if err == nil {
			t.Errorf("Expected to found one error")
		}
	})
}

func TestAdd(t *testing.T) {
	compareStrings := func(t *testing.T, output, expectation string) {
		t.Helper()
		if output != expectation {
			t.Errorf("output `%s`, expected `%s`", output, expectation)
		}
	}
	dictionary := Dictionary{"test": "This is a test"}

	t.Run("if we found a word", func(t *testing.T) {
		err := dictionary.Add("test2", "secondTest")
		if err != nil {
			t.Errorf("Wasn't expected to found one error")
		}

		output, err := dictionary.Search("test2")

		if err != nil {
			t.Errorf("Wasn't expected to found one error")
		}

		expectation := "secondTest"
		compareStrings(t, output, expectation)
	})

}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}

	t.Run("deleting one word", func(t *testing.T) {
		err := dictionary.Delete("test")
		if err != nil {
			t.Errorf("Wasn't expected to found one error")
		}

		_, err = dictionary.Search("test")

		if err == nil {
			t.Errorf("Was expect to found one error")
		}
	})

}
