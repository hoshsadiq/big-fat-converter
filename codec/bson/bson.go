// Package bson is a basic bson converter.
package bson

import (
	"errors"
	"io"
)

func init() {
	//nolint:gocritic
	//converter.Register([]string{"bson"}, &Converter{})
}

// Converter for bson.
type Converter struct{}

// Decode for bson.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	return nil, errors.New("decoding bson files is not implemented yet")
}

// Encode for bson.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	return errors.New("encoding bson files is not implemented yet")
}
