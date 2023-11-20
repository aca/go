package main

import (
	"os"

	"github.com/aca/go/csvx"
	"github.com/samber/lo"
)

type Human struct {
	ID int
	TT string
}

type Article struct {
	Name string
	XXX  int
	ZZZ  []string
	Man Human
}

type String string

func main() {
	arr := []Article{
		{
			Name: "rw",
			XXX:  3,
			ZZZ:  []string{"wer", "er"},
		},
		{
			Name: "r33w",
			XXX:  5,
		},
	}

    enc := csvx.NewEncoder[Article](os.Stdout)
	lo.Must0(enc.Encode(arr))
}
