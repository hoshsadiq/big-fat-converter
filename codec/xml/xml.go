// Package xml is an XML converter.
package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

func init() {
	converter.Register([]string{"xml"}, &Converter{})
}

// Converter for XML.
type Converter struct{}

// Doc represents an XML tag and children. This is extremely primitive to keep things simple,
// thus this likely doesn't capture all edge cases.
type Doc struct {
	Tag        string            `json:"tag_name,omitempty" yaml:"tag_name,omitempty" xml:"tag_name,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty" xml:"attributes,omitempty"`
	Children   []*Doc            `json:"children,omitempty" yaml:"children,omitempty" xml:"children,omitempty"`
	Value      string            `json:"value,omitempty" yaml:"value,omitempty" xml:"value,innerxml"`
}

func (md *Doc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error { //nolint:golint
	md.Tag = start.Name.Local
	md.Attributes = make(map[string]string, len(start.Attr))
	for _, attr := range start.Attr {
		md.Attributes[attr.Name.Local] = attr.Value
	}

	var innerxml string
	for {
		t, err := d.Token()
		if err != nil {
			return fmt.Errorf("failed to get XML token: %w", err)
		}
		switch tt := t.(type) {
		case xml.StartElement:
			item := new(Doc)
			if err = d.DecodeElement(item, &tt); err != nil {
				return fmt.Errorf("failed to decode XML element: %w", err)
			}
			md.Children = append(md.Children, item)
		case xml.CharData:
			innerxml = string(tt)
		case xml.EndElement:
			if tt == start.End() {
				if len(md.Children) == 0 && innerxml != "" {
					md.Value = innerxml
				}

				return nil
			}
		}
	}
}

// Decode for XML.
func (y *Converter) Decode(r io.ReadSeeker) (interface{}, error) {
	var data Doc
	x := xml.NewDecoder(r)
	x.CharsetReader = charset.NewReaderLabel
	err := x.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	return data, nil
}

// Encode for XML.
func (y *Converter) Encode(data interface{}, w io.Writer) error {
	// todo, we need to convert the interface{} to an actual object then pass that into xml encoder
	return errors.New("decoding XML files is not implemented yet")
}
