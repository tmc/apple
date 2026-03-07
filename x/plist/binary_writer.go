package plist

import (
	"bytes"
	"encoding/binary"
	"math"
	"sort"
	"time"
)

// encodeBinary encodes a plist intermediate representation to binary
// plist format (bplist00).
func encodeBinary(v any) ([]byte, error) {
	w := &binaryWriter{
		scalarIndex: make(map[any]int),
	}
	w.flatten(v)
	return w.encode()
}

// binaryWriter builds and serializes a binary plist object table.
//
// The flatten phase does a depth-first walk of the value tree, assigning
// each unique object a sequential index. Dict keys and array elements
// store their children's indices rather than the values themselves.
type binaryWriter struct {
	objects     []bpObject
	scalarIndex map[any]int // dedup for hashable scalars
}

// bpObject is a flattened binary plist object. For dicts it stores
// sorted key/value index pairs; for arrays it stores child indices.
type bpObject struct {
	value    any   // scalar value (string, int64, bool, []byte, etc.) or nil for containers
	kind     byte  // 'd' dict, 'a' array, 's' scalar
	children []int // for arrays: [valIdx...]; for dicts: [keyIdx, valIdx, keyIdx, valIdx, ...]
}

func (w *binaryWriter) flatten(v any) {
	w.addObject(v)
}

// addObject adds v to the object table and returns its index.
func (w *binaryWriter) addObject(v any) int {
	switch val := v.(type) {
	case map[string]any:
		idx := len(w.objects)
		w.objects = append(w.objects, bpObject{kind: 'd'})

		keys := sortedKeys(val)
		children := make([]int, 0, len(keys)*2)
		for _, k := range keys {
			ki := w.addObject(k)
			vi := w.addObject(val[k])
			children = append(children, ki, vi)
		}
		w.objects[idx].children = children
		return idx

	case []any:
		idx := len(w.objects)
		w.objects = append(w.objects, bpObject{kind: 'a'})

		children := make([]int, len(val))
		for i, item := range val {
			children[i] = w.addObject(item)
		}
		w.objects[idx].children = children
		return idx

	default:
		// Deduplicate hashable scalar values.
		if isHashable(v) {
			if idx, ok := w.scalarIndex[v]; ok {
				return idx
			}
		}
		idx := len(w.objects)
		w.objects = append(w.objects, bpObject{kind: 's', value: v})
		if isHashable(v) {
			w.scalarIndex[v] = idx
		}
		return idx
	}
}

func (w *binaryWriter) encode() ([]byte, error) {
	numObjects := len(w.objects)
	objectRefSize := intSize(numObjects)

	var buf bytes.Buffer
	buf.WriteString("bplist00")

	offsets := make([]int, numObjects)
	for i := range w.objects {
		offsets[i] = buf.Len()
		if err := w.writeObject(&buf, &w.objects[i], objectRefSize); err != nil {
			return nil, err
		}
	}

	// Offset table.
	offsetTableOffset := buf.Len()
	offsetSize := intSize(offsetTableOffset)
	for _, off := range offsets {
		writeIntN(&buf, off, offsetSize)
	}

	// Trailer (32 bytes).
	var trailer [32]byte
	trailer[6] = byte(offsetSize)
	trailer[7] = byte(objectRefSize)
	binary.BigEndian.PutUint64(trailer[8:16], uint64(numObjects))
	binary.BigEndian.PutUint64(trailer[16:24], 0) // top object
	binary.BigEndian.PutUint64(trailer[24:32], uint64(offsetTableOffset))
	buf.Write(trailer[:])

	return buf.Bytes(), nil
}

func (w *binaryWriter) writeObject(buf *bytes.Buffer, obj *bpObject, refSize int) error {
	switch obj.kind {
	case 'd':
		n := len(obj.children) / 2
		writeLength(buf, 0xD0, n)
		// Key refs.
		for i := 0; i < len(obj.children); i += 2 {
			writeIntN(buf, obj.children[i], refSize)
		}
		// Value refs.
		for i := 1; i < len(obj.children); i += 2 {
			writeIntN(buf, obj.children[i], refSize)
		}

	case 'a':
		writeLength(buf, 0xA0, len(obj.children))
		for _, ci := range obj.children {
			writeIntN(buf, ci, refSize)
		}

	case 's':
		return writeScalar(buf, obj.value)
	}
	return nil
}

func writeScalar(buf *bytes.Buffer, v any) error {
	switch val := v.(type) {
	case nil:
		buf.WriteByte(0x00)
	case bool:
		if val {
			buf.WriteByte(0x09)
		} else {
			buf.WriteByte(0x08)
		}
	case int64:
		writeInt(buf, val)
	case int:
		writeInt(buf, int64(val))
	case float64:
		buf.WriteByte(0x23)
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], math.Float64bits(val))
		buf.Write(b[:])
	case string:
		writeString(buf, val)
	case []byte:
		writeLength(buf, 0x40, len(val))
		buf.Write(val)
	case time.Time:
		buf.WriteByte(0x33)
		appleEpoch := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
		secs := val.Sub(appleEpoch).Seconds()
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], math.Float64bits(secs))
		buf.Write(b[:])
	default:
		buf.WriteByte(0x00) // unknown → null
	}
	return nil
}

func writeString(buf *bytes.Buffer, s string) {
	ascii := true
	for _, r := range s {
		if r > 127 {
			ascii = false
			break
		}
	}
	if ascii {
		writeLength(buf, 0x50, len(s))
		buf.WriteString(s)
	} else {
		// Encode as UTF-16BE. Code points above U+FFFF require surrogate pairs.
		utf16Units := stringToUTF16(s)
		writeLength(buf, 0x60, len(utf16Units))
		for _, u := range utf16Units {
			var b [2]byte
			binary.BigEndian.PutUint16(b[:], u)
			buf.Write(b[:])
		}
	}
}

// stringToUTF16 converts a Go string to a UTF-16 code unit slice,
// using surrogate pairs for code points above U+FFFF.
func stringToUTF16(s string) []uint16 {
	var out []uint16
	for _, r := range s {
		if r <= 0xFFFF {
			out = append(out, uint16(r))
		} else {
			r -= 0x10000
			out = append(out, uint16(0xD800+(r>>10)), uint16(0xDC00+(r&0x3FF)))
		}
	}
	return out
}

// writeInt writes an integer in the smallest encoding that fits.
func writeInt(buf *bytes.Buffer, val int64) {
	switch {
	case val >= 0 && val <= 0xFF:
		buf.WriteByte(0x10)
		buf.WriteByte(byte(val))
	case val >= 0 && val <= 0xFFFF:
		buf.WriteByte(0x11)
		var b [2]byte
		binary.BigEndian.PutUint16(b[:], uint16(val))
		buf.Write(b[:])
	case val >= 0 && val <= 0xFFFFFFFF:
		buf.WriteByte(0x12)
		var b [4]byte
		binary.BigEndian.PutUint32(b[:], uint32(val))
		buf.Write(b[:])
	default:
		buf.WriteByte(0x13)
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(val))
		buf.Write(b[:])
	}
}

// writeLength writes the type marker and length for variable-length objects.
func writeLength(buf *bytes.Buffer, marker byte, length int) {
	if length < 15 {
		buf.WriteByte(marker | byte(length))
	} else {
		buf.WriteByte(marker | 0x0F)
		writeInt(buf, int64(length))
	}
}

// writeIntN writes an integer in exactly n bytes (big-endian).
func writeIntN(buf *bytes.Buffer, val, n int) {
	for i := n - 1; i >= 0; i-- {
		buf.WriteByte(byte(val >> (uint(i) * 8)))
	}
}

// intSize returns the number of bytes needed to represent max.
func intSize(max int) int {
	switch {
	case max <= 0xFF:
		return 1
	case max <= 0xFFFF:
		return 2
	case max <= 0xFFFFFFFF:
		return 4
	default:
		return 8
	}
}

// isHashable reports whether v can be used as a map key.
func isHashable(v any) bool {
	switch v.(type) {
	case []byte, []any:
		return false
	default:
		return true
	}
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
