package resources

import (
	"time"

	"github.com/francoispqt/gojay"
)

type MapResponse map[string]interface{}
type ListArray []interface{}

func (req *MapResponse) Validate() error {
	return nil
}

// MarshalJSONObject Implement for ListMap
func (req MapResponse) MarshalJSONObject(enc *gojay.Encoder) {
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
func (req MapResponse) IsNil() bool {
	return req == nil
}


// MarshalJSONArray Implement for ListArray
func (req ListArray) MarshalJSONArray(enc *gojay.Encoder) {
	for _, v := range req {
		switch vt := v.(type) {
		case map[string]interface{}:
			enc.Object(MapResponse(vt))
		case []interface{}:
			enc.AddArray(ListArray(vt))
		default:
			enc.AddInterface(v)
		}
	}
}
func (req ListArray) IsNil() bool {
	return req == nil
}