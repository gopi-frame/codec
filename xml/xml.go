package xml

import (
	"bytes"
	"encoding/xml"
)

// Codec codec
type Codec struct {
	Prefix string
	Indent string
}

// Marshal marshal
func (c *Codec) Marshal(data any) ([]byte, error) {
	buffer := bytes.NewBufferString("")
	encoder := xml.NewEncoder(buffer)
	encoder.Indent(c.Prefix, c.Indent)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

// Unmarshal unmarshal
func (c *Codec) Unmarshal(data []byte, dest any) error {
	return xml.Unmarshal(data, dest)
}
