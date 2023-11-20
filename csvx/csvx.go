package csvx

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/noborus/trdsql"
)

type Encoder[T any] struct {
	TRDSQL *trdsql.TRDSQL
	Writer io.Writer
	Delimiter string
}

func NewEncoder[T any](w io.Writer) *Encoder[T] {
	return &Encoder[T]{
		TRDSQL: &trdsql.TRDSQL{},
		Writer: w,
		Delimiter: ",",
	}
}

func (e *Encoder[T]) SetDelimiter(delim string) {
	e.Delimiter = delim
}

func (e *Encoder[T]) Encode(arr []T) error {
	tmpfile, err := os.CreateTemp(os.TempDir(), "csvx")
	if err != nil {
		return nil
	}
	defer func() {
		// error handle
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}()

	err = json.NewEncoder(tmpfile).Encode(arr)
	if err != nil {
		return err
	}

	importer := trdsql.NewImporter(
		trdsql.InFormat(trdsql.JSON),
		trdsql.InHeader(true),
	)

	writer := trdsql.NewWriter(
		trdsql.OutFormat(trdsql.CSV),
		trdsql.OutDelimiter(e.Delimiter),
		trdsql.OutHeader(true),
		trdsql.OutStream(e.Writer),
	)

	exporter := trdsql.NewExporter(writer)
	trd := trdsql.NewTRDSQL(importer, exporter)
	query := fmt.Sprintf("SELECT * FROM %s", tmpfile.Name())
	err = trd.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
