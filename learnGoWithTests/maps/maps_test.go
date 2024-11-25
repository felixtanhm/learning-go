package maps

import (
	"testing"
)

func checkString(t testing.TB, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("Result %s | Expected %s", result, expected)
	}
}

func checkError(t testing.TB, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("Result %s | Expected %s", result, expected)
	}
}

func checkDefinition(t testing.TB, key, expected string, dict Dictionary) {
	t.Helper()

	result, err := dict.Search(key)
	if err != nil {
		t.Fatal("could not find word", err)
	}

	checkString(t, result, expected)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}
	t.Run("Known Word: ", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is just a test"

		checkString(t, result, expected)
	})

	t.Run("Unknown Word: ", func(t *testing.T) {
		_, err := dictionary.Search("hello")
		expected := ErrNotFound

		if err == nil {
			t.Fatal("expected an error but didn't get one.")
		}

		checkError(t, err, expected)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add new word: ", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		expected := "this is just a test"
		err := dictionary.Add(word, expected)

		checkError(t, err, nil)
		checkDefinition(t, word, expected, dictionary)
	})

	t.Run("Add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		checkError(t, err, ErrWordExists)
		checkDefinition(t, word, definition, dictionary)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update new word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("newTest", "new test")

		checkError(t, err, ErrWordDoesNotExist)
		checkDefinition(t, "test", "this is just a test", dictionary)
	})

	t.Run("Update existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}
		err := dictionary.Update(word, "new test")

		checkError(t, err, nil)
		checkDefinition(t, word, "new test", dictionary)
	})
}
