package encodingx

import (
	"io"

	"golang.org/x/text/encoding/korean"
)

var euckrDecoder = korean.EUCKR.NewDecoder()
var euckrEncoder = korean.EUCKR.NewEncoder()

func Decode_EUCKR_Bytes(b []byte) ([]byte, error) {
<<<<<<<< HEAD:encodingx/euckr.go
	return korean.EUCKR.NewDecoder().Bytes(b)
========
	defer func() {
		euckrDecoder.Reset()
	}()
	return euckrDecoder.Bytes(b)
>>>>>>>> ad8d2cd ():encodingx/encodingx.go
}

func Decode_EUCKR_String(s string) (string, error) {
	return korean.EUCKR.NewDecoder().String(s)
}

func Decode_EUCKR_Reader(r io.Reader) (io.Reader) {
	return korean.EUCKR.NewDecoder().Reader(r)
}

func Encode_EUCKR_Bytes(b []byte) ([]byte, error) {
	return korean.EUCKR.NewEncoder().Bytes(b)
}

func Encode_EUCKR_String(s string) (string, error) {
	return korean.EUCKR.NewEncoder().String(s)
}

func Encode_EUCKR_Reader(w io.Writer) (io.Writer) {
	return korean.EUCKR.NewEncoder().Writer(w)
}
