package main

import "testing"

func TestGeometry(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, esperado float64) {
		t.Helper()
		resultado := shape.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("square", func(t *testing.T) {
		square := Square{12.0, 6.0}
		checkArea(t, square, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func TestArea(t *testing.T) {
	testsArea := []struct {
		shape       Shape
		expectation float64
	}{
		{Square{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range testsArea {
		output := tt.shape.Area()
		if output != tt.expectation {
			t.Errorf("resultado %.2f, esperado %.2f", output, tt.expectation)
		}
	}
}
