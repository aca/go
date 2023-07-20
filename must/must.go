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

func Equal2[T comparable](got T, want T, msg string) error {
	if want == got {
		return nil
	}
	
	return fmt.Errorf("%v: got %#v, want %#v", msg, got, want)
}
