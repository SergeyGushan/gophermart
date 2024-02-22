package entity

import "fmt"

type DuplicateError struct {
	Field string
	Value string
}

func (e DuplicateError) Error() string {
	return fmt.Sprintf("%s %s already exists", e.Field, e.Value)
}
