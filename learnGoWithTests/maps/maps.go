package maps

import "fmt"

type Dictionary map[string]string

var ErrNotFound = fmt.Errorf("no such word available")
var ErrWordExists = fmt.Errorf("word already exists")
var ErrWordDoesNotExist = fmt.Errorf("word does not exist")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, definition string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, definition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = definition
	default:
		return err
	}
	return nil
}
