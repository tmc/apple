package plist

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Marshal encodes a Go value into plist data in the specified format.
// Struct fields are encoded using the "plist" struct tag. If no tag is
// present, the field name is used. The "omitempty" option causes the
// field to be omitted when it has a zero value.
func Marshal(v any, format Format) ([]byte, error) {
	intermediate, err := marshalValue(reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}
	switch format {
	case FormatBinary:
		return encodeBinary(intermediate)
	case FormatXML:
		return encodeXML(intermediate)
	case FormatJSON:
		return encodeJSON(intermediate)
	default:
		return nil, fmt.Errorf("unsupported format for marshal: %d", format)
	}
}

// marshalValue converts a Go reflect.Value into a plist intermediate
// representation (map[string]any, []any, string, int64, float64, bool,
// []byte, time.Time).
func marshalValue(v reflect.Value) (any, error) {
	if !v.IsValid() {
		return nil, nil
	}

	// Dereference pointers.
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil, nil
		}
		v = v.Elem()
	}

	// Dereference interfaces.
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return nil, nil
		}
		v = v.Elem()
	}

	// Preserve UID values as-is.
	if v.Type() == reflect.TypeOf(UID(0)) {
		return v.Interface().(UID), nil
	}

	switch v.Kind() {
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(time.Time{}) {
			return v.Interface().(time.Time), nil
		}
		return marshalStruct(v)

	case reflect.Map:
		return marshalMap(v)

	case reflect.Slice:
		// []byte → data
		if v.Type().Elem().Kind() == reflect.Uint8 {
			if v.IsNil() {
				return nil, nil
			}
			return v.Bytes(), nil
		}
		return marshalSlice(v)

	case reflect.String:
		return v.String(), nil

	case reflect.Bool:
		return v.Bool(), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint()), nil

	case reflect.Float32, reflect.Float64:
		return v.Float(), nil

	default:
		return nil, fmt.Errorf("unsupported type: %s", v.Type())
	}
}

func marshalStruct(v reflect.Value) (map[string]any, error) {
	m := make(map[string]any)
	if err := marshalStructFields(v, m); err != nil {
		return nil, err
	}
	return m, nil
}

func marshalStructFields(v reflect.Value, m map[string]any) error {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		fv := v.Field(i)
		// Flatten anonymous (embedded) struct fields.
		if f.Anonymous {
			for fv.Kind() == reflect.Ptr {
				if fv.IsNil() {
					break
				}
				fv = fv.Elem()
			}
			if fv.Kind() == reflect.Struct {
				if err := marshalStructFields(fv, m); err != nil {
					return err
				}
				continue
			}
		}
		name, opts := parseTag(f)
		if name == "-" {
			continue
		}
		if opts.omitempty && isZero(fv) {
			continue
		}
		val, err := marshalValue(fv)
		if err != nil {
			return fmt.Errorf("field %s: %w", f.Name, err)
		}
		if val == nil {
			continue
		}
		m[name] = val
	}
	return nil
}

func marshalMap(v reflect.Value) (map[string]any, error) {
	if v.IsNil() {
		return nil, nil
	}
	m := make(map[string]any)
	for _, key := range v.MapKeys() {
		k := fmt.Sprintf("%v", key.Interface())
		val, err := marshalValue(v.MapIndex(key))
		if err != nil {
			return nil, fmt.Errorf("map key %s: %w", k, err)
		}
		if val == nil {
			continue
		}
		m[k] = val
	}
	return m, nil
}

func marshalSlice(v reflect.Value) ([]any, error) {
	if v.IsNil() {
		return nil, nil
	}
	result := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		val, err := marshalValue(v.Index(i))
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}
		result[i] = val
	}
	return result, nil
}

type tagOpts struct {
	omitempty bool
}

func parseTag(f reflect.StructField) (string, tagOpts) {
	tag := f.Tag.Get("plist")
	if tag == "" {
		return f.Name, tagOpts{}
	}
	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = f.Name
	}
	var opts tagOpts
	for _, p := range parts[1:] {
		if p == "omitempty" {
			opts.omitempty = true
		}
	}
	return name, opts
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Struct:
		if v.Type() == reflect.TypeOf(time.Time{}) {
			return v.Interface().(time.Time).IsZero()
		}
		return false
	default:
		return false
	}
}

// Unmarshal decodes plist data into a Go value. The value must be a
// pointer to a struct, map, slice, or other supported type. Returns
// the detected format.
func Unmarshal(data []byte, v any) (Format, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return 0, fmt.Errorf("unmarshal requires a non-nil pointer")
	}

	format := DetectFormat(data)
	parsed, err := ParseBytes(data)
	if err != nil {
		return format, err
	}

	if err := unmarshalValue(parsed, rv.Elem()); err != nil {
		return format, err
	}
	return format, nil
}

// unmarshalValue assigns a parsed plist value into a reflect.Value.
func unmarshalValue(src any, dst reflect.Value) error {
	if src == nil {
		return nil
	}

	// Dereference/allocate pointers.
	for dst.Kind() == reflect.Ptr {
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		dst = dst.Elem()
	}

	// Handle UID destination type.
	if dst.Type() == reflect.TypeOf(UID(0)) {
		if u, ok := src.(UID); ok {
			dst.Set(reflect.ValueOf(u))
			return nil
		}
		return fmt.Errorf("expected UID, got %T", src)
	}

	switch dst.Kind() {
	case reflect.Struct:
		if dst.Type() == reflect.TypeOf(time.Time{}) {
			if t, ok := src.(time.Time); ok {
				dst.Set(reflect.ValueOf(t))
				return nil
			}
			return fmt.Errorf("expected time.Time, got %T", src)
		}
		m, ok := src.(map[string]any)
		if !ok {
			return fmt.Errorf("expected dict for struct, got %T", src)
		}
		return unmarshalStruct(m, dst)

	case reflect.Map:
		m, ok := src.(map[string]any)
		if !ok {
			return fmt.Errorf("expected dict for map, got %T", src)
		}
		return unmarshalMap(m, dst)

	case reflect.Slice:
		return unmarshalSlice(src, dst)

	case reflect.String:
		switch s := src.(type) {
		case string:
			dst.SetString(s)
		default:
			dst.SetString(fmt.Sprintf("%v", src))
		}

	case reflect.Bool:
		if b, ok := src.(bool); ok {
			dst.SetBool(b)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch n := src.(type) {
		case int64:
			dst.SetInt(n)
		case int:
			dst.SetInt(int64(n))
		case float64:
			dst.SetInt(int64(n))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch n := src.(type) {
		case int64:
			dst.SetUint(uint64(n))
		case int:
			dst.SetUint(uint64(n))
		case float64:
			dst.SetUint(uint64(n))
		}

	case reflect.Float32, reflect.Float64:
		switch n := src.(type) {
		case float64:
			dst.SetFloat(n)
		case int64:
			dst.SetFloat(float64(n))
		case int:
			dst.SetFloat(float64(n))
		}

	case reflect.Interface:
		dst.Set(reflect.ValueOf(src))

	default:
		return fmt.Errorf("cannot unmarshal into %s", dst.Type())
	}
	return nil
}

func unmarshalStruct(m map[string]any, dst reflect.Value) error {
	fields := collectStructFields(dst.Type())
	for key, val := range m {
		fi, ok := fields[key]
		if !ok {
			continue
		}
		fv := fieldByIndex(dst, fi.index)
		if err := unmarshalValue(val, fv); err != nil {
			return fmt.Errorf("field %s: %w", key, err)
		}
	}
	return nil
}

type fieldInfo struct {
	index []int // path of field indices for nested embedded structs
}

// collectStructFields builds a map from plist tag name to field index path,
// flattening anonymous (embedded) struct fields.
func collectStructFields(t reflect.Type) map[string]fieldInfo {
	fields := make(map[string]fieldInfo)
	collectFields(t, nil, fields)
	return fields
}

func collectFields(t reflect.Type, path []int, fields map[string]fieldInfo) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		idx := append(append([]int(nil), path...), i)
		if f.Anonymous {
			ft := f.Type
			for ft.Kind() == reflect.Ptr {
				ft = ft.Elem()
			}
			if ft.Kind() == reflect.Struct {
				collectFields(ft, idx, fields)
				continue
			}
		}
		name, _ := parseTag(f)
		if name == "-" {
			continue
		}
		if _, exists := fields[name]; !exists {
			fields[name] = fieldInfo{index: idx}
		}
	}
}

// fieldByIndex traverses the index path, allocating embedded pointer
// structs as needed.
func fieldByIndex(v reflect.Value, index []int) reflect.Value {
	for _, i := range index {
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				v.Set(reflect.New(v.Type().Elem()))
			}
			v = v.Elem()
		}
		v = v.Field(i)
	}
	return v
}

func unmarshalMap(m map[string]any, dst reflect.Value) error {
	if dst.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("map key must be string, got %s", dst.Type().Key())
	}
	if dst.IsNil() {
		dst.Set(reflect.MakeMap(dst.Type()))
	}
	elemType := dst.Type().Elem()
	for key, val := range m {
		elem := reflect.New(elemType).Elem()
		if err := unmarshalValue(val, elem); err != nil {
			return fmt.Errorf("map key %s: %w", key, err)
		}
		dst.SetMapIndex(reflect.ValueOf(key), elem)
	}
	return nil
}

func unmarshalSlice(src any, dst reflect.Value) error {
	// []byte ← []byte (data)
	if dst.Type().Elem().Kind() == reflect.Uint8 {
		if b, ok := src.([]byte); ok {
			dst.SetBytes(b)
			return nil
		}
		return fmt.Errorf("expected data for []byte, got %T", src)
	}

	arr, ok := src.([]any)
	if !ok {
		return fmt.Errorf("expected array for slice, got %T", src)
	}

	slice := reflect.MakeSlice(dst.Type(), len(arr), len(arr))
	for i, val := range arr {
		if err := unmarshalValue(val, slice.Index(i)); err != nil {
			return fmt.Errorf("index %d: %w", i, err)
		}
	}
	dst.Set(slice)
	return nil
}

// encodeXML marshals an intermediate plist value to XML bytes.
func encodeXML(v any) ([]byte, error) {
	var buf strings.Builder
	if err := WriteXML(&buf, v); err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}

// encodeJSON marshals an intermediate plist value to JSON bytes.
// Note: []byte data values are base64-encoded and time.Time values
// are RFC 3339 strings, matching the WriteJSON behavior.
func encodeJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	if err := WriteJSON(&buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
