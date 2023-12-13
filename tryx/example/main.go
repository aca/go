package main

import (
	"log/slog"
	"os"

	"github.com/aca/go/must"
	try "github.com/aca/go/tryx"
)

func main(){
	err := x()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Error("hello", slog.Any("error", err))
}

func x() ( ferr error ) {
	defer try.Handle(&ferr)
	try.E(must.Equal("3", "4"))
	return nil
}


