package must

import (
	"fmt"
)

func Equal[T comparable](got T, want T) error {
	if want == got {
		return nil
	}

	return fmt.Errorf("got %#v, want %#v", got, want)
}

func EqualM[T comparable](got T, want T, msg string) error {
	if want == got {
		return nil
	}

	return fmt.Errorf("got %#v, want %#v, msg: %s", got, want, msg)
}

func EqualE[T comparable](got T, want T, err error) error {
	if want == got {
		return nil
	}

	return fmt.Errorf("got %#v, want %#v, err: %w", got, want, err)
}
