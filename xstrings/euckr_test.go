package xstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	x := []byte{0xbe, 0xc8, 0xb3, 0xe7}
	t.Log(string(x))
	t.Log(Decode_EUCKR_Bytes(x))
}

func TestAppendString(t *testing.T) {
	s := "안녕123"
	s_euckr, err := Encode_EUCKR_String(s)
	if err != nil {
		t.Fatal(err)
	}
	s += "456"
	s_euckr += "456"

	s2, err := Decode_EUCKR_String(s_euckr) 
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, s, s2)
}
