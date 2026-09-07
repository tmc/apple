package plist

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestUnsignedIdentifierRoundTrip(t *testing.T) {
	for _, format := range []Format{FormatXML, FormatBinary} {
		for _, want := range []uint64{0, 1, 1<<63 - 1, 1 << 63, ^uint64(0)} {
			t.Run(fmt.Sprintf("%d/%d", format, want), func(t *testing.T) {
				type identity struct {
					ECID uint64 `plist:"ECID"`
				}
				data, err := Marshal(identity{want}, format)
				if err != nil {
					t.Fatal(err)
				}
				var got identity
				if _, err := Unmarshal(data, &got); err != nil || got.ECID != want {
					t.Fatalf("got %x want %x err %v", got.ECID, want, err)
				}
			})
		}
	}
}
func TestUnsignedIdentifierOverflow(t *testing.T) {
	for _, value := range []any{int64(-1), uint64(1 << 32), ^uint64(0)} {
		data, err := Marshal(map[string]any{"ID": value}, FormatXML)
		if err != nil {
			t.Fatal(err)
		}
		var d struct{ ID uint32 }
		if _, err := Unmarshal(data, &d); err == nil {
			t.Fatalf("accepted %v as uint32", value)
		}
	}
}

// Fixtures were independently generated with Python's standard plistlib.
// The tests themselves do not require Python.
func TestUnsignedBinaryFixtures(t *testing.T) {
	for _, tt := range []struct {
		name, hex string
		want      uint64
	}{
		{"high bit", "62706c69737430301400000000000000008000000000000000080000000000000101000000000000000100000000000000000000000000000019", 1 << 63},
		{"maximum", "62706c6973743030140000000000000000ffffffffffffffff080000000000000101000000000000000100000000000000000000000000000019", ^uint64(0)},
	} {
		t.Run(tt.name, func(t *testing.T) {
			data, err := hex.DecodeString(tt.hex)
			if err != nil {
				t.Fatal(err)
			}
			value, err := ParseBytes(data)
			if err != nil || value != tt.want {
				t.Fatal(value, err)
			}
			encoded, err := Marshal(tt.want, FormatBinary)
			if err != nil || !bytes.Equal(encoded, data) {
				t.Fatalf("fixture mismatch: %x, %v", encoded, err)
			}
			data[9] = 1
			if _, err := ParseBytes(data); err == nil {
				t.Fatal("accepted integer larger than uint64")
			}
		})
	}
}
func TestSignedIntegerOverflow(t *testing.T) {
	for _, value := range []any{int64(128), int64(-129), uint64(1 << 63)} {
		data, err := Marshal(value, FormatXML)
		if err != nil {
			t.Fatal(err)
		}
		var small int8
		if _, err := Unmarshal(data, &small); err == nil {
			t.Fatalf("accepted %v as int8", value)
		}
	}
}
