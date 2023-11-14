package encodingx

import (
	"io"

	"golang.org/x/text/encoding/korean"
)

var euckrDecoder = korean.EUCKR.NewDecoder()
var euckrEncoder = korean.EUCKR.NewEncoder()

func Decode_EUCKR_Bytes(b []byte) ([]byte, error) {
	return euckrDecoder.Bytes(b)
}

func Decode_EUCKR_String(s string) (string, error) {
	return euckrDecoder.String(s)
}

func Decode_EUCKR_Reader(r io.Reader) (io.Reader) {
	return euckrDecoder.Reader(r)
}

func Encode_EUCKR_Bytes(b []byte) ([]byte, error) {
	return euckrEncoder.Bytes(b)
}

func Encode_EUCKR_String(s string) (string, error) {
	return euckrEncoder.String(s)
}

func Encode_EUCKR_Writer(w io.Writer) (io.Writer) {
	return euckrEncoder.Writer(w)
}
