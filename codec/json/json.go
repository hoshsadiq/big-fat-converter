// Package json is a json converter.
package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"json"}, &Converter{})
}

// Converter for json.
type Converter struct{}

// Decode converts any Reader to a go interface{}.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	dec := json.NewDecoder(r)

	var data interface{}
	err := dec.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("faild to decode json file: %w", err)
	}

	return data, nil
}

// Encode converts any structure to json and writes it out to an io.Writer.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	return enc.Encode(data)
}
