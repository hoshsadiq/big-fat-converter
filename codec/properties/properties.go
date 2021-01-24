// Package properties is a properties file converter.
package properties

import (
	"errors"
	"io"
)

func init() {
	//nolint:gocritic
	//converter.Register([]string{"properties", "prop", "props"}, &Converter{})
}

// Converter for properties.
type Converter struct{}

// Decode for properties.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	return nil, errors.New("decoding properties files is not implemented yet")
}

// Encode for properties.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	return errors.New("decoding properties files is not implemented yet")
}
