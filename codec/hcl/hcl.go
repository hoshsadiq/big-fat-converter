// Package hcl is a hcl converter.
package hcl

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/tmccombs/hcl2json/convert"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"hcl", "tf", "tfvars"}, &Converter{})
}

// Converter for hcl.
type Converter struct{}

// Decode for hcl.
func (y *Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("faild to read hcl file: %w", err)
	}

	file, diags := hclsyntax.ParseConfig(input, "input.tf", hcl.Pos{Line: 1, Column: 1}) //nolint:exhaustivestruct
	if diags.HasErrors() {
		return nil, fmt.Errorf("parse config: %v", diags.Errs())
	}

	convertedFile, err := convert.ConvertFile(file, convert.Options{
		Simplify: true,
	})
	if err != nil {
		return nil, fmt.Errorf("faild to decode hcl file: %w", err)
	}

	return map[string]interface{}(convertedFile), nil
}

// Encode for hcl.
func (y *Converter) Encode(data interface{}, w io.Writer) error {
	// We're unable to convert to a generic map.
	// JSON is a subset of HCL, maybe we should convert it JSON?
	return errors.New("encoding HCL files is not implemented yet")
}
