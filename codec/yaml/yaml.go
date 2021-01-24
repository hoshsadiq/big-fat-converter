// Package yaml is a YAML converter.
package yaml

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v2"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"yml", "yaml"}, &Converter{})
}

// Converter for YAML.
type Converter struct{}

// Decode for YAML.
func (y *Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	dec := yaml.NewDecoder(r)

	var data interface{}
	err := dec.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode input: %w", err)
	}
	data = fix(data)

	return data, nil
}

// Encode for YAML.
func (y *Converter) Encode(data interface{}, w io.Writer) error {
	enc := yaml.NewEncoder(w)

	return enc.Encode(data)
}

// YAML creates types like map[interface{}]interface{}, which are not a valid JSON types.
// This function is a best-effort of fixing representation of YAML-encoded value, so it
// can be marshaled to JSON types.
func fix(v interface{}) interface{} {
	switch v := v.(type) {
	case *interface{}:
		return fix(*v)
	case map[string]interface{}:
		for k, w := range v {
			v[k] = fix(w)
		}

		return v
	case map[interface{}]interface{}:
		fixedV := make(map[string]interface{}, len(v))

		for k, v := range v {
			fixedV[fmt.Sprintf("%v", k)] = fix(v)
		}

		return fixedV
	case []interface{}:
		fixedV := make([]interface{}, len(v))

		for i := range v {
			fixedV[i] = fix(v[i])
		}

		return fixedV
	default:
		return v
	}
}
