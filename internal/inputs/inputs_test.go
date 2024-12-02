package inputs

import (
	"slices"
	"testing"
)

func TestExtractIntPairs(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {
		input := ``
		_, _, err := ExtractIntPairs(input)

		if err == nil {
			t.Errorf("expected an error, didnt get one")
		}
	})

	t.Run("invalid line", func(t *testing.T) {
		input := `123 456 789`
		_, _, err := ExtractIntPairs(input)

		if err == nil {
			t.Errorf("expected an error, didnt get one")
		}
	})

	t.Run("invalid left column", func(t *testing.T) {
		input := `abc 123`
		_, _, err := ExtractIntPairs(input)

		if err == nil {
			t.Errorf("expected an error, didnt get one")
		}
	})

	t.Run("invalid right column", func(t *testing.T) {
		input := `123 abc`
		_, _, err := ExtractIntPairs(input)

		if err == nil {
			t.Errorf("expected an error, didnt get one")
		}
	})

	t.Run("test slice content", func(t *testing.T) {
		input := "111 222\n333 444"
		leftWant := []int{111, 333}
		rightWant := []int{222, 444}
		leftGot, rightGot, _ := ExtractIntPairs(input)

		if !slices.Equal(leftGot, leftWant) || !slices.Equal(rightGot, rightWant) {
			t.Errorf("Expected %v, %v, got %v, %v", leftWant, leftGot, rightWant, rightGot)
		}
	})
}
