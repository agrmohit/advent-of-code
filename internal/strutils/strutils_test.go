package strutils

import "testing"

func TestReverse(t *testing.T) {
	t.Run("test ascii string", func(t *testing.T) {
		input := "Hello World!"
		want := "!dlroW olleH"
		got := Reverse(input)

		if got != want {
			t.Errorf("got '%v', want '%v', given '%v'", got, want, input)
		}
	})

	t.Run("test utf-8 string", func(t *testing.T) {
		input := "The quick brown 狐 jumped over the lazy 犬"
		want := "犬 yzal eht revo depmuj 狐 nworb kciuq ehT"
		got := Reverse(input)

		if got != want {
			t.Errorf("got '%v', want '%v', given '%v'", got, want, input)
		}
	})
}
