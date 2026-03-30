package plist

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"
)

func TestBinaryParseUID(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte // UID marker + data bytes
		want    UID
	}{
		{
			name:    "1-byte uid 0",
			payload: []byte{0x80, 0x00},
			want:    0,
		},
		{
			name:    "1-byte uid 5",
			payload: []byte{0x80, 0x05},
			want:    5,
		},
		{
			name:    "1-byte uid 255",
			payload: []byte{0x80, 0xFF},
			want:    255,
		},
		{
			name:    "2-byte uid 256",
			payload: []byte{0x81, 0x01, 0x00},
			want:    256,
		},
		{
			name:    "2-byte uid 1000",
			payload: []byte{0x81, 0x03, 0xE8},
			want:    1000,
		},
		{
			name:    "4-byte uid",
			payload: []byte{0x83, 0x00, 0x01, 0x00, 0x00},
			want:    65536,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := buildBinaryPlist(tt.payload)
			v, err := ParseBytes(data)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}
			uid, ok := v.(UID)
			if !ok {
				t.Fatalf("got %T (%v), want UID", v, v)
			}
			if uid != tt.want {
				t.Errorf("uid = %d, want %d", uid, tt.want)
			}
		})
	}
}

func TestBinaryUIDRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		uid  UID
	}{
		{"zero", 0},
		{"small", 7},
		{"byte boundary", 255},
		{"two byte", 1000},
		{"large", 100000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := map[string]any{"ref": tt.uid}
			data, err := Marshal(orig, FormatBinary)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			parsed, err := ParseBytes(data)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}
			m := parsed.(map[string]any)
			got, ok := m["ref"].(UID)
			if !ok {
				t.Fatalf("ref is %T (%v), want UID", m["ref"], m["ref"])
			}
			if got != tt.uid {
				t.Errorf("uid = %d, want %d", got, tt.uid)
			}
		})
	}
}

func TestUnmarshalNSKeyedArchive(t *testing.T) {
	// Build a minimal NSKeyedArchiver-like binary plist:
	// {
	//   "$archiver": "NSKeyedArchiver",
	//   "$top": {"root": UID(1)},
	//   "$objects": ["$null", "hello world"],
	// }
	archive := map[string]any{
		"$archiver": "NSKeyedArchiver",
		"$top":      map[string]any{"root": UID(1)},
		"$objects":  []any{"$null", "hello world"},
	}

	data, err := Marshal(archive, FormatBinary)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	// Unmarshal into map[string]any and verify structure.
	var got map[string]any
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	archiver, ok := got["$archiver"].(string)
	if !ok || archiver != "NSKeyedArchiver" {
		t.Errorf("$archiver = %v", got["$archiver"])
	}

	top, ok := got["$top"].(map[string]any)
	if !ok {
		t.Fatalf("$top is %T, want map[string]any", got["$top"])
	}
	rootRef, ok := top["root"].(UID)
	if !ok {
		t.Fatalf("$top.root is %T (%v), want UID", top["root"], top["root"])
	}
	if rootRef != 1 {
		t.Errorf("$top.root = %d, want 1", rootRef)
	}

	objects, ok := got["$objects"].([]any)
	if !ok {
		t.Fatalf("$objects is %T, want []any", got["$objects"])
	}
	if len(objects) != 2 {
		t.Fatalf("len($objects) = %d, want 2", len(objects))
	}

	// Dereference the UID.
	resolved := objects[rootRef]
	if resolved != "hello world" {
		t.Errorf("$objects[root] = %v, want 'hello world'", resolved)
	}
}

func TestUnmarshalUIDIntoStruct(t *testing.T) {
	type Ref struct {
		Root UID `plist:"root"`
	}

	orig := Ref{Root: 42}
	data, err := Marshal(orig, FormatBinary)
	if err != nil {
		t.Fatal(err)
	}

	var got Ref
	if _, err := Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got.Root != 42 {
		t.Errorf("Root = %d, want 42", got.Root)
	}
}

func TestUIDTypeOf(t *testing.T) {
	if got := TypeOf(UID(5)); got != "uid" {
		t.Errorf("TypeOf(UID) = %q, want 'uid'", got)
	}
}

func TestBinaryIntegerSigns(t *testing.T) {
	tests := []struct {
		name string
		val  int64
	}{
		{"zero", 0},
		{"positive small", 42},
		{"positive large", 1<<31 - 1},
		{"negative small", -1},
		{"negative large", -1000000},
		{"max int64", 1<<63 - 1},
		{"min int64", -1 << 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := map[string]any{"n": tt.val}
			data, err := Marshal(orig, FormatBinary)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var got map[string]any
			if _, err := Unmarshal(data, &got); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			n, ok := got["n"].(int64)
			if !ok {
				t.Fatalf("n is %T (%v), want int64", got["n"], got["n"])
			}
			if n != tt.val {
				t.Errorf("n = %d, want %d", n, tt.val)
			}
		})
	}
}

// Verify interface compliance at compile time.
var (
	_ fmt.Stringer             = UID(0)
	_ encoding.TextMarshaler   = UID(0)
	_ encoding.TextUnmarshaler = (*UID)(nil)
	_ json.Marshaler           = UID(0)
	_ json.Unmarshaler         = (*UID)(nil)
)

func TestUIDString(t *testing.T) {
	tests := []struct {
		uid  UID
		want string
	}{
		{0, "UID(0)"},
		{42, "UID(42)"},
		{1<<32 - 1, "UID(4294967295)"},
	}
	for _, tt := range tests {
		if got := tt.uid.String(); got != tt.want {
			t.Errorf("UID(%d).String() = %q, want %q", tt.uid, got, tt.want)
		}
	}
}

func TestUIDTextMarshal(t *testing.T) {
	u := UID(12345)
	b, err := u.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "12345" {
		t.Errorf("MarshalText = %q, want '12345'", b)
	}

	var got UID
	if err := got.UnmarshalText(b); err != nil {
		t.Fatal(err)
	}
	if got != u {
		t.Errorf("UnmarshalText round-trip = %d, want %d", got, u)
	}
}

func TestUIDTextUnmarshalError(t *testing.T) {
	var u UID
	if err := u.UnmarshalText([]byte("notanumber")); err == nil {
		t.Error("expected error for invalid text")
	}
}

func TestUIDJSON(t *testing.T) {
	// Direct json.Marshal of UID.
	u := UID(99)
	b, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "99" {
		t.Errorf("json.Marshal(UID(99)) = %q, want '99'", b)
	}

	var got UID
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if got != 99 {
		t.Errorf("json.Unmarshal = %d, want 99", got)
	}
}

func TestUIDJSONInMap(t *testing.T) {
	// Verify json.Marshal works on map[string]any containing UID values.
	m := map[string]any{
		"root": UID(3),
		"name": "test",
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("json.Marshal map with UID: %v", err)
	}
	// Should contain "root":3 (not a nested object or string).
	if !bytes.Contains(b, []byte(`"root":3`)) {
		t.Errorf("json output = %s, want root:3", b)
	}
}

func TestUIDJSONUnmarshalError(t *testing.T) {
	var u UID
	if err := json.Unmarshal([]byte(`"notanumber"`), &u); err == nil {
		t.Error("expected error for invalid JSON")
	}
}

// buildBinaryPlist wraps a single object payload in a minimal valid bplist00.
func buildBinaryPlist(objectPayload []byte) []byte {
	var buf bytes.Buffer

	// Header.
	buf.WriteString("bplist00")

	// Object data starts at offset 8.
	objectOffset := buf.Len()
	buf.Write(objectPayload)

	// Offset table.
	offsetTableOffset := buf.Len()
	buf.WriteByte(byte(objectOffset))

	// Trailer (32 bytes).
	var trailer [32]byte
	trailer[6] = 1                                                        // offset size
	trailer[7] = 1                                                        // object ref size
	binary.BigEndian.PutUint64(trailer[8:16], 1)                          // num objects
	binary.BigEndian.PutUint64(trailer[16:24], 0)                         // top object
	binary.BigEndian.PutUint64(trailer[24:32], uint64(offsetTableOffset)) // offset table offset
	buf.Write(trailer[:])

	return buf.Bytes()
}
