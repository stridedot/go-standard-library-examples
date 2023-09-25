package testing

import (
	"testing"
)

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func TestSearch(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("exists word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("exists word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newValue := "new value"
		dictionary := Dictionary{key: value}
		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "new test"
		value := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test definition"}

	dictionary.Delete(key)
	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", key)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, want string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
