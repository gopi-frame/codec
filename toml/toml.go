package toml

import (
	"bytes"

	"github.com/pelletier/go-toml/v2"
)

// Codec codec
type Codec struct {
	TableInline     bool
	ArraysMultiline bool
	IndentSymbol    string
	IndentTables    bool
}

// Marshal marshal
func (c *Codec) Marshal(data any) ([]byte, error) {
	buffer := bytes.NewBufferString("")
	encoder := toml.NewEncoder(buffer)
	encoder.SetTablesInline(c.TableInline)
	encoder.SetArraysMultiline(c.ArraysMultiline)
	encoder.SetIndentSymbol(c.IndentSymbol)
	encoder.SetIndentTables(c.IndentTables)
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

// Unmarshal unmarshal
func (c *Codec) Unmarshal(data []byte, dest any) error {
	return toml.Unmarshal(data, dest)
}
