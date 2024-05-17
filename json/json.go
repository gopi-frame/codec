package json

import (
	"bytes"
	"encoding/json"
)

// Codec json codec
type Codec struct {
	Prefix     string
	Indent     string
	EscapeHTML bool
}

// Marshal marshal
func (c *Codec) Marshal(data any) ([]byte, error) {
	buffer := bytes.NewBufferString("")
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(c.EscapeHTML)
	encoder.SetIndent(c.Prefix, c.Indent)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

// Unmarshal unmarshal
func (c *Codec) Unmarshal(data []byte, dest any) error {
	return json.Unmarshal(data, dest)
}
