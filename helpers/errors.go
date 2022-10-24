package helpers

import (
	"errors"
	"fmt"
)

type DuplicateError struct {
	Message string
}

func (d *DuplicateError) Error() string {
	return d.Message
}

func getError() error {
	return errors.New("new Error")
}

func main() {
	err := getError()

	fmt.Println("err", err)
}