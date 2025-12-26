package resources

import (
	"time"

	"github.com/francoispqt/gojay"
)

type MapRequest map[string]interface{}
type List []interface{}

func (req *MapRequest) Validate() error {
	return nil
}

// Implementing Unmarshaler
func (req *MapRequest) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	var value interface{}
	err := dec.Interface(&value)
	if err != nil {
		return err
	}
	(*req)[key] = value
	return nil
}
// we return 0, it tells the Decoder to decode all keys
func (req *MapRequest) NKeys() int {
	return 0
}

// MarshalJSONObject Implement for ListMap
func (req MapRequest) MarshalJSONObject(enc *gojay.Encoder) {
	for k, v := range req {
		switch vt := v.(type) {
		case map[string]interface{}:
			enc.ObjectKey(k, MapResponse(vt))
		case time.Time:
			enc.TimeKey(k, &vt, "2006-01-02T15:04:05")
		case []interface{}:
			enc.AddArrayKey(k, ListArray(vt))
		default:
			enc.AddInterfaceKey(k, v)
		}
	}
}
func (req MapRequest) IsNil() bool {
	return req == nil
}


// MarshalJSONArray Implement for ListArray
func (req List) MarshalJSONArray(enc *gojay.Encoder) {
	for _, v := range req {
		switch vt := v.(type) {
		case map[string]interface{}:
			enc.Object(MapResponse(vt))
		case []interface{}:
			enc.AddArray(List(vt))
		default:
			enc.AddInterface(v)
		}
	}
}
func (req List) IsNil() bool {
	return req == nil
}