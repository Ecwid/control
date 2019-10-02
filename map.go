package witness

import (
	"encoding/json"
	"errors"
	"reflect"
)

type bytes []byte

// Map ...
type Map map[string]interface{}

func (b bytes) MarshalJSON() ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}
	return b, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (b *bytes) UnmarshalJSON(data []byte) error {
	if b == nil {
		return errors.New("bytes: UnmarshalJSON on nil pointer")
	}
	*b = append((*b)[0:0], data...)
	return nil
}

func (b bytes) Unmarshal(to interface{}) error {
	return json.Unmarshal(b, to)
}

func (m Map) String(key string) string {
	if s, ok := m[key]; ok {
		val, ok := s.(string)
		if !ok {
			return ""
		}
		return val
	}
	return ""
}

// Int ...
func (m Map) Int(key string) int64 {
	return int64(m.Float(key))
}

// Float ...
func (m Map) Float(key string) float64 {
	if s, ok := m[key]; ok {
		return s.(float64)
	}
	return 0
}

// Bool ...
func (m Map) Bool(key string) bool {
	if s, ok := m[key]; ok {
		return s.(bool)
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func (m Map) omitempty() {
	for k, v := range m {
		if isEmptyValue(reflect.ValueOf(v)) {
			delete(m, k)
		}
	}
}

func (b bytes) json() Map {
	to := Map{}
	if err := b.Unmarshal(&to); err != nil {
		panic(err)
	}
	return to
}
