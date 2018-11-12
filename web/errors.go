package web

import (
	"errors"
	"fmt"
)

type ErrDocDoesNotExist struct {
	docName string
}

func (e ErrDocDoesNotExist) Error() string {
	return fmt.Sprintf("Document '%s' does not exist", e.docName)
}

var ErrGeneric = errors.New("Something went wrong")
