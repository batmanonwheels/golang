package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeats ten times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeats zero times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "a"
		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
