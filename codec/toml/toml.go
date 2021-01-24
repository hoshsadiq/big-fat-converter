// Package toml is a TOML converter
package toml

import (
	"fmt"
	"io"

	"github.com/pelletier/go-toml"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"toml"}, &Converter{})
}

// Converter for TOML.
type Converter struct{}

// Decode for TOML.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	dec := toml.NewDecoder(r)

	var data interface{}
	err := dec.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("faild to decode TOML file: %w", err)
	}

	return data, nil
}

// Encode for TOML.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	enc := toml.NewEncoder(w)

	return enc.Encode(data)
}
