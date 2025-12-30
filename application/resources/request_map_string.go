package resources

import (
	"github.com/francoispqt/gojay"
)

type MapString map[string]string

func (req *MapString) Validate() error {
	return nil
}

// UnmarshalJSONObject implementing Unmarshaler for MapString
func (req *MapString) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	var value string
	err := dec.String(&value)
	if err != nil {
		return err
	}
	(*req)[key] = value
	return nil
}

// NKeys return 0, it tells the Decoder to decode all keys
func (req *MapString) NKeys() int {
	return 0
}

// MarshalJSONObject Implement for MapString
func (req MapString) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range req {
		enc.AddStringKey(k, v)
	}
}

// IsNil Implement for MapString
func (req MapString) IsNil() bool {
	return req == nil
}
