package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/aca/go/flat"
)

type Person struct {
	Name string
	Info Info
}

type Info struct {
	ID string
	Gender string
	Slices []string
}

func main() {
	x := Person {
		Name: "james", 
		Info: Info{
			Gender: "girl",
			Slices: []string{"wer", "wefcc"},
		},
	}

	spew.Dump(flat.Flat(x))
	spew.Dump(flat.Flat([]string{"werrwe","werwre"}))

	spew.Dump(flat.Flat(map[string]string{"hello": "werw3r"}))
}
