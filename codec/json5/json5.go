// Package json5 is a json5 converter.
package json5

import (
	"fmt"
	"io"

	"github.com/yosuke-furukawa/json5/encoding/json5"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"json"}, &Converter{})
}

// Converter for json5.
type Converter struct{}

// Decode for json5.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	dec := json5.NewDecoder(r)

	var data interface{}
	err := dec.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("faild to decode json5 file: %w", err)
	}

	return data, nil
}

// Encode for json5.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	enc := json5.NewEncoder(w)

	return enc.Encode(data)
}
