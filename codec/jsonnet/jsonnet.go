// Package jsonnet is a jsonnet converter.
package jsonnet

import (
	"errors"
	"io"
)

func init() {
	//nolint:gocritic
	//converter.Register([]string{"jsonnet"}, &Converter{})
}

// Converter for jsonnet.
type Converter struct{}

// Decode for jsonnet.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	return nil, errors.New("decoding JSONNet files is not implemented yet")
}

// Encode for jsonnet.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	return errors.New("encoding JSONNet files is not implemented yet")
}
