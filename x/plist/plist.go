package plist

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Format represents a plist format.
type Format int

const (
	FormatXML Format = iota
	FormatBinary
	FormatJSON
)

// DetectFormat detects the plist format from data.
func DetectFormat(data []byte) Format {
	if len(data) >= 6 && string(data[:6]) == "bplist" {
		return FormatBinary
	}
	// Check for JSON (starts with { or [)
	for _, b := range data {
		if b == ' ' || b == '\t' || b == '\n' || b == '\r' {
			continue
		}
		if b == '{' || b == '[' {
			return FormatJSON
		}
		break
	}
	return FormatXML
}

// Parse parses plist data from a reader (auto-detects format).
func Parse(r io.Reader) (any, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

// ParseFile parses a plist file (auto-detects format).
func ParseFile(path string) (any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

// ParseBytes parses plist data (auto-detects format).
func ParseBytes(data []byte) (any, error) {
	switch DetectFormat(data) {
	case FormatBinary:
		return parseBinary(data)
	case FormatJSON:
		return parseJSON(data)
	default:
		return parseXML(data)
	}
}

// parseJSON parses JSON plist data.
func parseJSON(data []byte) (any, error) {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	// Convert float64 to int64 where appropriate
	return normalizeJSON(v), nil
}

func normalizeJSON(v any) any {
	switch val := v.(type) {
	case map[string]any:
		m := make(map[string]any)
		for k, vv := range val {
			m[k] = normalizeJSON(vv)
		}
		return m
	case []any:
		a := make([]any, len(val))
		for i, vv := range val {
			a[i] = normalizeJSON(vv)
		}
		return a
	case float64:
		// Convert to int64 if it's a whole number
		if val == math.Trunc(val) && val >= math.MinInt64 && val <= math.MaxInt64 {
			return int64(val)
		}
		return val
	default:
		return v
	}
}

// parseXML parses XML plist data.
func parseXML(data []byte) (any, error) {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))

	// Find <plist> element
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		if se, ok := tok.(xml.StartElement); ok && se.Name.Local == "plist" {
			break
		}
	}

	return parseXMLValue(decoder)
}

func parseXMLValue(decoder *xml.Decoder) (any, error) {
	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "dict":
				return parseXMLDict(decoder)
			case "array":
				return parseXMLArray(decoder)
			case "string":
				return parseXMLText(decoder, "string")
			case "integer":
				text, err := parseXMLText(decoder, "integer")
				if err != nil {
					return nil, err
				}
				return strconv.ParseInt(text, 10, 64)
			case "real":
				text, err := parseXMLText(decoder, "real")
				if err != nil {
					return nil, err
				}
				return strconv.ParseFloat(text, 64)
			case "true":
				skipXMLToEnd(decoder, "true")
				return true, nil
			case "false":
				skipXMLToEnd(decoder, "false")
				return false, nil
			case "data":
				text, err := parseXMLText(decoder, "data")
				if err != nil {
					return nil, err
				}
				// Strip ALL whitespace (Apple plists often have embedded newlines in base64)
				text = strings.Map(func(r rune) rune {
					if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
						return -1 // remove
					}
					return r
				}, text)
				return base64.StdEncoding.DecodeString(text)
			case "date":
				text, err := parseXMLText(decoder, "date")
				if err != nil {
					return nil, err
				}
				return time.Parse(time.RFC3339, text)
			}
		case xml.EndElement:
			return nil, nil
		}
	}
}

func parseXMLDict(decoder *xml.Decoder) (map[string]any, error) {
	result := make(map[string]any)

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.EndElement:
			if t.Name.Local == "dict" {
				return result, nil
			}
		case xml.StartElement:
			if t.Name.Local == "key" {
				key, err := parseXMLText(decoder, "key")
				if err != nil {
					return nil, err
				}
				value, err := parseXMLValue(decoder)
				if err != nil {
					return nil, err
				}
				result[key] = value
			}
		}
	}
}

func parseXMLArray(decoder *xml.Decoder) ([]any, error) {
	var result []any

	for {
		val, err := parseXMLValue(decoder)
		if err != nil {
			return nil, err
		}
		if val == nil {
			return result, nil
		}
		result = append(result, val)
	}
}

func parseXMLText(decoder *xml.Decoder, endTag string) (string, error) {
	var text strings.Builder

	for {
		tok, err := decoder.Token()
		if err != nil {
			return "", err
		}

		switch t := tok.(type) {
		case xml.CharData:
			text.Write(t)
		case xml.EndElement:
			if t.Name.Local == endTag {
				return text.String(), nil
			}
		}
	}
}

func skipXMLToEnd(decoder *xml.Decoder, endTag string) {
	for {
		tok, err := decoder.Token()
		if err != nil {
			return
		}
		if ee, ok := tok.(xml.EndElement); ok && ee.Name.Local == endTag {
			return
		}
	}
}

// Binary plist parsing
func parseBinary(data []byte) (any, error) {
	if len(data) < 32 {
		return nil, fmt.Errorf("binary plist too short")
	}
	if string(data[:6]) != "bplist" {
		return nil, fmt.Errorf("not a binary plist")
	}

	// Parse trailer (last 32 bytes)
	trailer := data[len(data)-32:]
	offsetSize := int(trailer[6])
	objectRefSize := int(trailer[7])
	numObjects := int(binary.BigEndian.Uint64(trailer[8:16]))
	topObject := int(binary.BigEndian.Uint64(trailer[16:24]))
	offsetTableOffset := int(binary.BigEndian.Uint64(trailer[24:32]))

	// Parse offset table
	offsets := make([]int, numObjects)
	for i := 0; i < numObjects; i++ {
		offset := 0
		for j := 0; j < offsetSize; j++ {
			offset = offset<<8 | int(data[offsetTableOffset+i*offsetSize+j])
		}
		offsets[i] = offset
	}

	bp := &binaryParser{data: data, offsets: offsets, objectRefSize: objectRefSize}
	return bp.parseObject(topObject)
}

type binaryParser struct {
	data          []byte
	offsets       []int
	objectRefSize int
}

func (bp *binaryParser) parseObject(index int) (any, error) {
	if index >= len(bp.offsets) {
		return nil, fmt.Errorf("invalid object index: %d", index)
	}
	offset := bp.offsets[index]
	if offset >= len(bp.data) {
		return nil, fmt.Errorf("invalid offset: %d", offset)
	}

	marker := bp.data[offset]
	objType := marker >> 4
	objInfo := int(marker & 0x0F)

	switch objType {
	case 0x0: // null, bool, fill
		switch objInfo {
		case 0x0:
			return nil, nil
		case 0x8:
			return false, nil
		case 0x9:
			return true, nil
		}
	case 0x1: // int
		size := 1 << objInfo
		return bp.parseInt(offset+1, size)
	case 0x2: // real
		size := 1 << objInfo
		return bp.parseReal(offset+1, size)
	case 0x3: // date
		return bp.parseDate(offset + 1)
	case 0x4: // data
		length, dataOffset := bp.parseLength(offset, objInfo)
		return bp.data[dataOffset : dataOffset+length], nil
	case 0x5: // ascii string
		length, dataOffset := bp.parseLength(offset, objInfo)
		return string(bp.data[dataOffset : dataOffset+length]), nil
	case 0x6: // unicode string
		length, dataOffset := bp.parseLength(offset, objInfo)
		return bp.parseUTF16(dataOffset, length)
	case 0xA: // array
		length, dataOffset := bp.parseLength(offset, objInfo)
		return bp.parseArray(dataOffset, length)
	case 0xD: // dict
		length, dataOffset := bp.parseLength(offset, objInfo)
		return bp.parseDict(dataOffset, length)
	}

	return nil, fmt.Errorf("unknown object type: %x at offset %d", marker, offset)
}

func (bp *binaryParser) parseLength(offset, objInfo int) (int, int) {
	if objInfo == 0x0F {
		// Length is in next byte(s)
		intMarker := bp.data[offset+1]
		intSize := 1 << (intMarker & 0x0F)
		length, _ := bp.parseInt(offset+2, intSize)
		return int(length), offset + 2 + intSize
	}
	return objInfo, offset + 1
}

func (bp *binaryParser) parseInt(offset, size int) (int64, error) {
	if offset+size > len(bp.data) {
		return 0, fmt.Errorf("int out of bounds")
	}
	var val int64
	for i := 0; i < size; i++ {
		val = val<<8 | int64(bp.data[offset+i])
	}
	return val, nil
}

func (bp *binaryParser) parseReal(offset, size int) (float64, error) {
	if size == 4 {
		bits := binary.BigEndian.Uint32(bp.data[offset : offset+4])
		return float64(math.Float32frombits(bits)), nil
	} else if size == 8 {
		bits := binary.BigEndian.Uint64(bp.data[offset : offset+8])
		return math.Float64frombits(bits), nil
	}
	return 0, fmt.Errorf("invalid real size: %d", size)
}

func (bp *binaryParser) parseDate(offset int) (time.Time, error) {
	bits := binary.BigEndian.Uint64(bp.data[offset : offset+8])
	secs := math.Float64frombits(bits)
	// Apple epoch is 2001-01-01
	appleEpoch := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	return appleEpoch.Add(time.Duration(secs * float64(time.Second))), nil
}

func (bp *binaryParser) parseUTF16(offset, length int) (string, error) {
	var runes []rune
	for i := 0; i < length; i++ {
		u := binary.BigEndian.Uint16(bp.data[offset+i*2:])
		if u >= 0xD800 && u <= 0xDBFF && i+1 < length {
			lo := binary.BigEndian.Uint16(bp.data[offset+(i+1)*2:])
			if lo >= 0xDC00 && lo <= 0xDFFF {
				r := rune((uint32(u)-0xD800)<<10 + (uint32(lo) - 0xDC00) + 0x10000)
				runes = append(runes, r)
				i++
				continue
			}
		}
		runes = append(runes, rune(u))
	}
	return string(runes), nil
}

func (bp *binaryParser) parseArray(offset, length int) ([]any, error) {
	result := make([]any, length)
	for i := 0; i < length; i++ {
		ref := bp.readRef(offset + i*bp.objectRefSize)
		val, err := bp.parseObject(ref)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}

func (bp *binaryParser) parseDict(offset, length int) (map[string]any, error) {
	result := make(map[string]any)
	keyOffset := offset
	valOffset := offset + length*bp.objectRefSize

	for i := 0; i < length; i++ {
		keyRef := bp.readRef(keyOffset + i*bp.objectRefSize)
		valRef := bp.readRef(valOffset + i*bp.objectRefSize)

		key, err := bp.parseObject(keyRef)
		if err != nil {
			return nil, err
		}
		keyStr, ok := key.(string)
		if !ok {
			return nil, fmt.Errorf("dict key is not a string")
		}

		val, err := bp.parseObject(valRef)
		if err != nil {
			return nil, err
		}
		result[keyStr] = val
	}
	return result, nil
}

func (bp *binaryParser) readRef(offset int) int {
	ref := 0
	for i := 0; i < bp.objectRefSize; i++ {
		ref = ref<<8 | int(bp.data[offset+i])
	}
	return ref
}

// Writing functions

// WriteXML writes plist data as XML to a writer.
func WriteXML(w io.Writer, v any) error {
	fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?>`)
	fmt.Fprintln(w, `<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">`)
	fmt.Fprintln(w, `<plist version="1.0">`)
	if err := writeXMLValue(w, v, 0); err != nil {
		return err
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, `</plist>`)
	return nil
}

func writeXMLValue(w io.Writer, v any, indent int) error {
	prefix := strings.Repeat("\t", indent)

	switch val := v.(type) {
	case map[string]any:
		fmt.Fprintln(w, prefix+"<dict>")
		// Sort keys for consistent output
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(w, "%s\t<key>%s</key>\n", prefix, xmlEscape(k))
			if err := writeXMLValue(w, val[k], indent+1); err != nil {
				return err
			}
			fmt.Fprintln(w)
		}
		fmt.Fprint(w, prefix+"</dict>")

	case []any:
		fmt.Fprintln(w, prefix+"<array>")
		for _, item := range val {
			if err := writeXMLValue(w, item, indent+1); err != nil {
				return err
			}
			fmt.Fprintln(w)
		}
		fmt.Fprint(w, prefix+"</array>")

	case string:
		fmt.Fprintf(w, "%s<string>%s</string>", prefix, xmlEscape(val))

	case int64:
		fmt.Fprintf(w, "%s<integer>%d</integer>", prefix, val)

	case int:
		fmt.Fprintf(w, "%s<integer>%d</integer>", prefix, val)

	case float64:
		fmt.Fprintf(w, "%s<real>%g</real>", prefix, val)

	case bool:
		if val {
			fmt.Fprint(w, prefix+"<true/>")
		} else {
			fmt.Fprint(w, prefix+"<false/>")
		}

	case []byte:
		encoded := base64.StdEncoding.EncodeToString(val)
		fmt.Fprintf(w, "%s<data>%s</data>", prefix, encoded)

	case time.Time:
		fmt.Fprintf(w, "%s<date>%s</date>", prefix, val.UTC().Format(time.RFC3339))

	case nil:
		// Skip nil values

	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	return nil
}

func xmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

// WriteJSON writes plist data as JSON to a writer.
func WriteJSON(w io.Writer, v any) error {
	v = convertForJSON(v)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func convertForJSON(v any) any {
	switch val := v.(type) {
	case map[string]any:
		m := make(map[string]any)
		for k, vv := range val {
			m[k] = convertForJSON(vv)
		}
		return m
	case []any:
		a := make([]any, len(val))
		for i, vv := range val {
			a[i] = convertForJSON(vv)
		}
		return a
	case []byte:
		return base64.StdEncoding.EncodeToString(val)
	case time.Time:
		return val.UTC().Format(time.RFC3339)
	case int64:
		return val
	default:
		return v
	}
}

// WriteBinary writes plist data as binary to a writer.
func WriteBinary(w io.Writer, v any) error {
	data, err := Marshal(v, FormatBinary)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// Lint validates plist data and returns any errors.
func Lint(data []byte) error {
	_, err := ParseBytes(data)
	return err
}

// LintFile validates a plist file and returns any errors.
func LintFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return Lint(data)
}

// TypeOf returns the plist type name for a value.
func TypeOf(v any) string {
	switch v.(type) {
	case string:
		return "string"
	case int64, int:
		return "integer"
	case float64:
		return "float"
	case bool:
		return "bool"
	case []byte:
		return "data"
	case time.Time:
		return "date"
	case []any:
		return "array"
	case map[string]any:
		return "dictionary"
	default:
		return "unknown"
	}
}

// Get retrieves a value at a key path (dot-separated or variadic).
func Get(v any, keys ...string) any {
	if len(keys) == 1 && strings.Contains(keys[0], ".") {
		keys = strings.Split(keys[0], ".")
	}

	current := v
	for _, key := range keys {
		switch c := current.(type) {
		case map[string]any:
			current = c[key]
		case []any:
			if i, err := strconv.Atoi(key); err == nil && i >= 0 && i < len(c) {
				current = c[i]
			} else {
				return nil
			}
		default:
			return nil
		}
	}
	return current
}

// Set sets a value at a key path. Returns the modified root value.
func Set(v any, keypath string, value any) (any, error) {
	keys := strings.Split(keypath, ".")
	if len(keys) == 0 {
		return value, nil
	}

	// Ensure root is a map
	root, ok := v.(map[string]any)
	if !ok {
		if v == nil {
			root = make(map[string]any)
		} else {
			return nil, fmt.Errorf("root is not a dictionary")
		}
	}

	// Navigate to parent
	current := root
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		next, ok := current[key].(map[string]any)
		if !ok {
			next = make(map[string]any)
			current[key] = next
		}
		current = next
	}

	// Set the value
	current[keys[len(keys)-1]] = value
	return root, nil
}

// Remove removes a value at a key path. Returns the modified root value.
func Remove(v any, keypath string) (any, error) {
	keys := strings.Split(keypath, ".")
	if len(keys) == 0 {
		return v, nil
	}

	root, ok := v.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("root is not a dictionary")
	}

	// Navigate to parent
	current := root
	for i := 0; i < len(keys)-1; i++ {
		key := keys[i]
		next, ok := current[key].(map[string]any)
		if !ok {
			return root, nil // Path doesn't exist
		}
		current = next
	}

	// Remove the key
	delete(current, keys[len(keys)-1])
	return root, nil
}

// String retrieves a string at a key path.
func String(v any, keys ...string) string {
	if s, ok := Get(v, keys...).(string); ok {
		return s
	}
	return ""
}

// Int retrieves an integer at a key path.
func Int(v any, keys ...string) int {
	switch n := Get(v, keys...).(type) {
	case int64:
		return int(n)
	case int:
		return n
	case float64:
		return int(n)
	}
	return 0
}

// Int64 retrieves an int64 at a key path.
func Int64(v any, keys ...string) int64 {
	switch n := Get(v, keys...).(type) {
	case int64:
		return n
	case int:
		return int64(n)
	case float64:
		return int64(n)
	}
	return 0
}

// Float retrieves a float64 at a key path.
func Float(v any, keys ...string) float64 {
	switch n := Get(v, keys...).(type) {
	case float64:
		return n
	case int64:
		return float64(n)
	case int:
		return float64(n)
	}
	return 0
}

// Bool retrieves a bool at a key path.
func Bool(v any, keys ...string) bool {
	if b, ok := Get(v, keys...).(bool); ok {
		return b
	}
	return false
}

// Bytes retrieves base64-decoded data at a key path.
func Bytes(v any, keys ...string) []byte {
	if b, ok := Get(v, keys...).([]byte); ok {
		return b
	}
	return nil
}

// Array retrieves an array at a key path.
func Array(v any, keys ...string) []any {
	if a, ok := Get(v, keys...).([]any); ok {
		return a
	}
	return nil
}

// Dict retrieves a dict at a key path.
func Dict(v any, keys ...string) map[string]any {
	if d, ok := Get(v, keys...).(map[string]any); ok {
		return d
	}
	return nil
}

// CreateEmpty creates an empty plist (empty dictionary).
func CreateEmpty() map[string]any {
	return make(map[string]any)
}

// WriteToFile writes plist data to a file in the specified format.
func WriteToFile(path string, v any, format Format) error {
	var buf bytes.Buffer
	var err error

	switch format {
	case FormatXML:
		err = WriteXML(&buf, v)
	case FormatJSON:
		err = WriteJSON(&buf, v)
	case FormatBinary:
		err = WriteBinary(&buf, v)
	default:
		err = fmt.Errorf("unknown format")
	}

	if err != nil {
		return err
	}

	return os.WriteFile(path, buf.Bytes(), 0644)
}
