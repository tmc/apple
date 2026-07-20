package oslog

import (
	"bytes"
	"testing"
)

// The expected buffers were captured from __builtin_os_log_format via a C
// program that hooks _os_log_impl and dumps the bytes clang emits.
func TestEncode(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []any
		// want is the buffer with string-pointer payloads zeroed (addresses are
		// not reproducible); wantHasPtr marks that an 8-byte pointer follows.
		want []byte
	}{
		{
			name:   "public int %d",
			format: "count=%d",
			args:   []any{int32(1234)},
			want:   []byte{0x00, 0x01, 0x00, 0x04, 0xd2, 0x04, 0x00, 0x00},
		},
		{
			name:   "private int %{private}d",
			format: "n=%{private}d",
			args:   []any{int32(5)},
			want:   []byte{0x01, 0x01, 0x01, 0x04, 0x05, 0x00, 0x00, 0x00},
		},
		{
			name:   "wide %ld is 8 bytes",
			format: "big=%ld",
			args:   []any{int64(1) << 40},
			want:   []byte{0x00, 0x01, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := encode(tt.format, tt.args)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("encode(%q) = % x, want % x", tt.format, got, tt.want)
			}
		})
	}
}

// A string argument yields the string descriptor, size 8, and the non-scalar
// summary flag; the payload is a pointer we do not check.
func TestEncodeString(t *testing.T) {
	got, pins := encode("host=%{public}s", []any{"apple.com"})
	if len(pins) != 1 {
		t.Fatalf("expected 1 pinned C string, got %d", len(pins))
	}
	want := []byte{summaryNonScalar, 0x01, visPublic | kindString, 0x08}
	if !bytes.Equal(got[:4], want) {
		t.Errorf("string header = % x, want % x", got[:4], want)
	}
	if len(got) != 4+8 {
		t.Errorf("string buffer len = %d, want %d", len(got), 4+8)
	}
}
