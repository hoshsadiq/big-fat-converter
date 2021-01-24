// Package dhall is a dhall lang converter.
package dhall

import (
	"errors"
	"io"
)

func init() {
	//nolint:gocritic
	//converter.Register([]string{"dhall"}, &Converter{})
}

// Converter struct for dhall.
type Converter struct{}

// Decode struct for dhall.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	return nil, errors.New("decoding dhall files is not implemented yet")
}

// Encode struct for dhall.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	return errors.New("encoding dhall files is not implemented yet")
}
