package must

import "testing"

func TestEqual(t *testing.T) {
	err := Equal("john", "james")
	if err != nil {
		t.Log(err)
	}

	err = EqualM("john", "james", "names")
	if err != nil {
		t.Log(err)
	}
}
