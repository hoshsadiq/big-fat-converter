// Package cue is a cuelang converter
package cue

import (
	"errors"
	"io"
)

func init() {
	//nolint:gocritic
	//converter.Register([]string{"cue", "cuelang"}, &Converter{})
}

// Converter struct for cuelang.
type Converter struct{}

// Decode for cuelang.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	return nil, errors.New("decoding Cue files is not implemented yet")
}

// Encode for cuelang.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	return errors.New("encoding Cue files is not implemented yet")
}
