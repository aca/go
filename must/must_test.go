package must

import "testing"

func TestEqual(t *testing.T) {
	err := Equal("john", "james")
	if err != nil {
		t.Log(err)
	}

	err = Equal2("john", "james", "names")
	if err != nil {
		t.Log(err)
	}
}
