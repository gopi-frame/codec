package yaml

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// Codec codec
type Codec struct {
	Space int
}

// Marshal marshal
func (c *Codec) Marshal(data any) ([]byte, error) {
	buffer := bytes.NewBufferString("")
	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(c.Space)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

// Unmarshal unmarshal
func (c *Codec) Unmarshal(data []byte, dest any) error {
	return yaml.Unmarshal(data, dest)
}
