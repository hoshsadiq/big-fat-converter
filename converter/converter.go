// Package converter defines an interface for writing generic converters.
package converter

import (
	"errors"
	"io"
)

// Converter interface that allows different converts to be defined.
// A converter should be able to decode, and encode.
type Converter interface {
	Decode(io.ReadSeeker) (interface{}, error)
	Encode(interface{}, io.Writer) error
}

var converters = make(map[string]Converter)

// Register a converter for a file format.
func Register(extensions []string, converter Converter) {
	for _, ext := range extensions {
		converters[ext] = converter
	}
}

// Get a converter for a file format.
func Get(ext string) (Converter, error) {
	converter, ok := converters[ext]
	if !ok {
		return nil, errors.New("no converter found for filetype")
	}

	return converter, nil
}

// GetFormats gets a list of all format extensions available.
func GetFormats() []string {
	var formats []string //nolint:prealloc
	for format := range converters {
		formats = append(formats, format)
	}

	return formats
}
