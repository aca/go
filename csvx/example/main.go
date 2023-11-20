package main

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/jszwec/csvutil"
	"github.com/samber/lo"
)

type Article struct {
	Name string
	XXX  int
	// ZZZ  []string
}

type Encoder[T any] struct {
	Encoder csvutil.Encoder
	Writer  csv.Writer
}

func NewEncoder[T any](w io.Writer) *Encoder[T] {
	writer := csv.NewWriter(w)
	enc := csvutil.NewEncoder(writer)
	return &Encoder[T]{
		Encoder: *enc,
		Writer:  *writer,
	}
}

func (e *Encoder[T]) Encode(arr []T) error {
	defer e.Writer.Flush()

	for _, elem := range arr {
		err := e.Encoder.Encode(elem)
		if err != nil {
		  return err
		}
	}

	return nil
}

func (e *Encoder[T]) SetDelimiter(x rune) {
	e.Writer.Comma = x
}

func main() {
	arr := []Article{
		{
			Name: "rw",
			XXX:  3,
			// ZZZ:  []string{"wer", "er"},
		},
		{
			Name: "r33w",
			XXX:  5,
		},
	}

    enc := NewEncoder[Article](os.Stdout)
	enc.SetDelimiter('.')
	lo.Must0(enc.Encode(arr))
}
