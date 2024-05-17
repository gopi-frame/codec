package codec

import (
	"sync"

	"github.com/gopi-frame/codec/json"
	"github.com/gopi-frame/codec/toml"
	"github.com/gopi-frame/codec/yaml"
	"github.com/gopi-frame/contract/codec"
	"github.com/gopi-frame/exception"
)

// codec
const (
	JSON = "json"
	XML  = "xml"
	YAML = "yaml"
	TOML = "toml"
)

var once sync.Once

var _factory *Factory

// Marshal marshal
func Marshal(name string, data any) ([]byte, error) {
	return newFactory().Marshal(name, data)
}

// Unmarshal unmarshal
func Unmarshal(name string, data []byte, dest any) error {
	return newFactory().Unmarshal(name, data, dest)
}

// RegisterMarshaler register marshaler
func RegisterMarshaler(name string, marshaler codec.Marshaler) {
	newFactory().RegisterMarshaler(name, marshaler)
}

// RegisterUnmarshaler register unmarshaler
func RegisterUnmarshaler(name string, unmarshaler codec.Unmarshaler) {
	newFactory().RegisterUnmarshaler(name, unmarshaler)
}

// newFactory new codec factory
func newFactory() *Factory {
	once.Do(func() {
		jsonCodec := &json.Codec{EscapeHTML: true}
		yamlCodec := &yaml.Codec{}
		tomlCodec := &toml.Codec{IndentSymbol: "  "}
		xmlCodec := &toml.Codec{}
		_factory = &Factory{
			marshalers: map[string]codec.Marshaler{
				JSON: jsonCodec,
				YAML: yamlCodec,
				TOML: tomlCodec,
				XML:  xmlCodec,
			},
			unmarshalers: map[string]codec.Unmarshaler{
				JSON: jsonCodec,
				YAML: yamlCodec,
				TOML: tomlCodec,
				XML:  xmlCodec,
			},
		}
	})
	return _factory
}

// Factory factory
type Factory struct {
	marshalers   map[string]codec.Marshaler
	unmarshalers map[string]codec.Unmarshaler
}

// RegisterMarshaler register marshaler
func (f *Factory) RegisterMarshaler(key string, marshaler codec.Marshaler) {
	f.marshalers[key] = marshaler
}

// RegisterUnmarshaler register unmarshaler
func (f *Factory) RegisterUnmarshaler(key string, unmarshaler codec.Unmarshaler) {
	f.unmarshalers[key] = unmarshaler
}

// Marshal marshal
func (f *Factory) Marshal(name string, data any) ([]byte, error) {
	if marshaler, ok := f.marshalers[name]; ok {
		return marshaler.Marshal(data)
	}
	return nil, exception.NewUnsupportedException("unsupported")
}

// Unmarshal unmarshal
func (f *Factory) Unmarshal(name string, data []byte, dest any) error {
	if unmarshaler, ok := f.unmarshalers[name]; ok {
		return unmarshaler.Unmarshal(data, dest)
	}
	return exception.NewUnsupportedException("unsupported")
}
