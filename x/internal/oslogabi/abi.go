// Package oslogabi builds the argument buffers of the unified logging
// system's ABI, shared by the os_log (x/oslog) and os_signpost (x/signpost)
// emitters. The buffer layout is produced by clang's __builtin_os_log_format
// in C; the oracle test validates this encoder against clang itself.
package oslogabi

import (
	"fmt"
	"math"
	"unsafe"
)

// os_log argument buffer format (verified against __builtin_os_log_format):
//
//	byte 0: summary flags — OR of per-argument flags across all arguments
//	        0x01 if any argument is private, 0x02 if any is non-scalar (string/pointer)
//	byte 1: argument count
//	per argument: [descriptor][size][payload]
//	  descriptor = visibility | (kind << 4)
//	    visibility: 0x01 private, 0x02 public, 0x00 default
//	    kind:       0x0 scalar, 0x2 string (payload is an 8-byte pointer)
//	  size:    payload length in bytes (4/8 for scalars, 8 for a string pointer)
//	  payload: little-endian scalar, or an 8-byte pointer to a NUL-terminated string
const (
	visPrivate = 0x01
	visPublic  = 0x02
	kindScalar = 0x00
	kindString = 0x20

	summaryPrivate   = 0x01
	summaryNonScalar = 0x02
)

// Encode builds the os_log argument buffer for format and args. It returns
// the buffer and a slice of values that must be kept alive for the duration
// of the emitting call (the C strings that buffer pointers refer into).
func Encode(format string, args []any) (buf []byte, pins []any) {
	specs := parseSpecs(format)
	// One byte summary + one byte count, then the arguments.
	buf = make([]byte, 2, 2+len(args)*10)
	var summary byte
	n := 0
	for i, s := range specs {
		if i >= len(args) {
			break
		}
		a := args[i]
		vis := s.visibility
		if s.isString {
			cs := cstringArg(a)
			pins = append(pins, cs)
			ptr := uintptr(unsafe.Pointer(&cs[0]))
			buf = append(buf, byte(vis)|kindString, 8)
			buf = appendUintptr(buf, ptr)
			summary |= summaryNonScalar
		} else {
			// The payload width follows the format specifier (%d=4, %ld=8),
			// not the Go type; the value is sign/zero-extended or truncated.
			var v uint64
			if s.isFloat {
				v = floatBits(a)
			} else {
				v, _ = scalarBits(a)
			}
			size := s.size
			if size == 0 {
				size = 4
			}
			buf = append(buf, byte(vis)|kindScalar, byte(size))
			buf = appendScalar(buf, v, size)
		}
		if vis == visPrivate {
			summary |= summaryPrivate
		}
		n++
	}
	buf[0] = summary
	buf[1] = byte(n)
	return buf, pins
}

// spec is one parsed os_log format specifier.
type spec struct {
	isString   bool
	isFloat    bool // %f/%e/%g family: the payload is IEEE754, not an integer
	visibility int  // visPrivate, visPublic, or 0 for default
	size       int  // scalar payload width the format expects: 4, or 8 for l/ll/z/q or %p
}

// parseSpecs scans an os_log format string and returns its specifiers in order.
// It recognizes the %{public}/%{private} annotations and classifies each
// conversion as string (%s, %@) or scalar (everything else numeric).
func parseSpecs(format string) []spec {
	var specs []spec
	for i := 0; i < len(format); i++ {
		if format[i] != '%' {
			continue
		}
		i++
		if i >= len(format) || format[i] == '%' {
			continue // "%%" is a literal percent, no argument
		}
		vis := 0
		if format[i] == '{' {
			// %{key}conv — read the annotation up to '}'.
			end := i + 1
			for end < len(format) && format[end] != '}' {
				end++
			}
			ann := format[i+1 : min(end, len(format))]
			switch {
			case containsWord(ann, "private"):
				vis = visPrivate
			case containsWord(ann, "public"):
				vis = visPublic
			}
			i = end + 1
		}
		// Skip flags, width, and precision; note length modifiers (l/ll/z/q/j/t)
		// which widen a scalar to 8 bytes, then stop at the conversion.
		size := 4
		for i < len(format) && !isConv(format[i]) {
			switch format[i] {
			case 'l', 'z', 'q', 'j', 't':
				size = 8
			}
			i++
		}
		if i >= len(format) {
			break
		}
		switch format[i] {
		case 's', '@':
			specs = append(specs, spec{isString: true, visibility: vis})
		case 'p':
			specs = append(specs, spec{isString: false, visibility: vis, size: 8})
		case 'f', 'F', 'e', 'E', 'g', 'G', 'a', 'A':
			// A float argument promotes to double in a variadic call, so the
			// payload is always 8 bytes of IEEE754 whatever the modifiers say.
			specs = append(specs, spec{isString: false, isFloat: true, visibility: vis, size: 8})
		default:
			specs = append(specs, spec{isString: false, visibility: vis, size: size})
		}
	}
	return specs
}

func isConv(c byte) bool {
	switch c {
	case 'd', 'i', 'u', 'x', 'X', 'o', 'p', 'c', 's', '@':
		return true
	case 'f', 'F', 'e', 'E', 'g', 'G', 'a', 'A':
		return true
	}
	return false
}

func containsWord(s, w string) bool {
	for i := 0; i+len(w) <= len(s); i++ {
		if s[i:i+len(w)] == w {
			return true
		}
	}
	return false
}

// scalarBits returns the little-endian value and byte size for a scalar arg.
// Integers wider than 32 bits and pointers use 8 bytes; everything else 4.
func scalarBits(a any) (v uint64, size int) {
	switch x := a.(type) {
	case int:
		return uint64(int64(x)), 8
	case int8:
		return uint64(int64(x)), 4
	case int16:
		return uint64(int64(x)), 4
	case int32:
		return uint64(int64(x)), 4
	case int64:
		return uint64(x), 8
	case uint:
		return uint64(x), 8
	case uint8:
		return uint64(x), 4
	case uint16:
		return uint64(x), 4
	case uint32:
		return uint64(x), 4
	case uint64:
		return x, 8
	case uintptr:
		return uint64(x), 8
	case bool:
		if x {
			return 1, 4
		}
		return 0, 4
	default:
		// Fall back to the string form for unknown scalar types.
		return 0, 4
	}
}

// floatBits returns the IEEE754 double bits for a %f/%e/%g argument. A variadic
// C call promotes float to double, so the payload is always binary64 even when
// the Go value is a float32. An integer given to a float specifier converts, so
// that log lines keep the value the caller passed rather than its bit pattern.
func floatBits(a any) uint64 {
	var f float64
	switch x := a.(type) {
	case float64:
		f = x
	case float32:
		f = float64(x)
	case int:
		f = float64(x)
	case int8:
		f = float64(x)
	case int16:
		f = float64(x)
	case int32:
		f = float64(x)
	case int64:
		f = float64(x)
	case uint:
		f = float64(x)
	case uint8:
		f = float64(x)
	case uint16:
		f = float64(x)
	case uint32:
		f = float64(x)
	case uint64:
		f = float64(x)
	}
	return math.Float64bits(f)
}

// cstringArg renders a as a NUL-terminated C string for a %s/%@ argument.
func cstringArg(a any) []byte {
	var s string
	switch x := a.(type) {
	case string:
		s = x
	case fmt.Stringer:
		s = x.String()
	case error:
		s = x.Error()
	default:
		s = fmt.Sprint(x)
	}
	return CString(s)
}

// CString returns a NUL-terminated copy of s. Callers keep the slice alive
// (via runtime.KeepAlive) for the duration of any call that reads it.
func CString(s string) []byte {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return b
}

func appendScalar(b []byte, v uint64, size int) []byte {
	for i := range size {
		b = append(b, byte(v>>(8*i)))
	}
	return b
}

func appendUintptr(b []byte, p uintptr) []byte {
	return appendScalar(b, uint64(p), 8)
}
