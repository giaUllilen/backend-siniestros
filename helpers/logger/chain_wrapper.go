package logger

import (
	"reflect"
	"strings"

	"github.com/francoispqt/onelog"
)

const jsonTag = "json"

type chainEntryWrapper struct {
	onelog.ChainEntry
}

func (chainEntry chainEntryWrapper) Object(key string, obj interface{}) chainEntryWrapper {

	chainEntry.ObjectFunc(key, chainEntry.writeObjectFields(obj))

	return chainEntry
}

func (chainEntry chainEntryWrapper) writeObjectFields(obj interface{}) func(entry onelog.Entry) {

	sType := reflect.TypeOf(obj)
	sValue := reflect.ValueOf(obj)

	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}

	if sType.Kind() == reflect.Struct {

		return func(enc onelog.Entry) {
			for i := 0; i < sType.NumField(); i++ {
				field := sType.Field(i)
				value := sValue.Field(i)
				chainEntry.writeField(enc, field, value)
			}
		}
	}

	if sType.Kind() == reflect.Map {
		//TODO
	}

	return func(enc onelog.Entry) {}
}

func (chainEntry chainEntryWrapper) writeField(entry onelog.Entry, field reflect.StructField, value reflect.Value) {

	fieldName := chainEntry.getFieldName(field)

	switch value.Kind() {

	case reflect.String:
		entry.String(fieldName, value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		entry.Int64(fieldName, value.Int())
	case reflect.Float32, reflect.Float64:
		entry.Float(fieldName, value.Float())
	case reflect.Bool:
		entry.Bool(fieldName, value.Bool())
	case reflect.Struct:
		entry.ObjectFunc(fieldName, chainEntry.writeObjectFields(value.Interface()))
	default:
		return
	}
}

func (chainEntry chainEntryWrapper) getFieldName(field reflect.StructField) string {

	if jsonTagValue, ok := field.Tag.Lookup(jsonTag); ok {
		return strings.Split(jsonTagValue, ",")[0]
	}
	return field.Name
}
