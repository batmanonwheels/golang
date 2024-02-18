package main

import (
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want DictionaryErrs) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDefine(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("get definition of known word", func(t *testing.T) {
		got, _ := dictionary.Define("test")

		assertStrings(t, got, "this is just a test")
	})

	t.Run("get definition of unknown word", func(t *testing.T) {
		_, got := dictionary.Define("trust")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("new word", func(t *testing.T) {
		got, err := dict.Add("new", "a novel addition")
		want := "a novel addition"

		if err != "" {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		_, got := dict.Add("test", "this is just a test")

		assertError(t, got, ErrDuplicateWord)
	})

}

func TestUpdate(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("update exisiting word", func(t *testing.T) {
		got, _ := dict.Update("test", "a novel addition")
		want, _ := dict.Define("test")

		assertStrings(t, string(got), want)
	})

	t.Run("update non-existant word", func(t *testing.T) {
		_, got := dict.Update("toast", "this is just a toast")

		assertError(t, got, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := Word("test")
	dict := Dictionary{word: "this is just a test"}
	dict.Delete(word)
	_, err := dict.Define(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}

}
