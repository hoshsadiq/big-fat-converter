// Package plist is a converter for plist.
package plist

import (
	"fmt"
	"io"

	"howett.net/plist"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"plist"}, &Converter{})
}

// Converter for plist.
type Converter struct{}

// Decode for plist.
func (y *Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	d := plist.NewDecoder(r)

	var data interface{}
	err := d.Decode(data)
	if err != nil {
		return nil, fmt.Errorf("faild to decode plist file: %w", err)
	}

	return data, nil
}

// Encode for plist.
func (y *Converter) Encode(data interface{}, w io.Writer) error {
	e := plist.NewEncoder(w)

	return e.Encode(data)
}
