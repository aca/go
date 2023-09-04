package xstrings

import (
	"io"

	"golang.org/x/text/encoding/korean"
)

var euckrDecoder = korean.EUCKR.NewDecoder()
var euckrEncoder = korean.EUCKR.NewEncoder()

func Decode_EUCKR_Bytes(b []byte) ([]byte, error) {
	// reset if error?
	defer func() {
		euckrDecoder.Reset()
	}()
	return euckrDecoder.Bytes(b)
}

func Decode_EUCKR_Strings(s string) (string, error) {
	defer func() {
		euckrDecoder.Reset()
	}()
	return euckrDecoder.String(s)
}

func Decode_EUCKR_Reader(r io.Reader) (io.Reader) {
	// reuse or not ?
	return korean.EUCKR.NewDecoder().Reader(r)
}

func Encode_EUCKR_Bytes(b []byte) ([]byte, error) {
	defer func() {
		euckrEncoder.Reset()
	}()
	return euckrEncoder.Bytes(b)
}

func Encode_EUCKR_String(s string) (string, error) {
	defer func() {
		euckrEncoder.Reset()
	}()
	return euckrEncoder.String(s)
}

func Encode_EUCKR_Reader(w io.Writer) (io.Writer) {
	// reuse or not ?
	return korean.EUCKR.NewEncoder().Writer(w)
}
