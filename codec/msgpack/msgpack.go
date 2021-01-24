// Package msgpack is a converter for MessagePack.
package msgpack

import (
	"fmt"
	"io"

	"github.com/vmihailenco/msgpack/v5"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"msgpack"}, &Converter{})
}

// Converter for MessagePack.
type Converter struct{}

// Decode for MessagePack.
func (c Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	dec := msgpack.NewDecoder(r)

	var data interface{}
	err := dec.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("faild to decode MessagePack file: %w", err)
	}

	return data, nil
}

// Encode for MessagePack.
func (c Converter) Encode(data interface{}, w io.Writer) error {
	enc := msgpack.NewEncoder(w)

	return enc.Encode(data)
}
