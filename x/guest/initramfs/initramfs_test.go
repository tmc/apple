package initramfs

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

// cpioEntry is a parsed newc record used to verify archive contents.
type cpioEntry struct {
	name string
	mode uint32
	data []byte
}

// parseNewc decodes a newc cpio archive into entries (stopping at the trailer).
func parseNewc(t *testing.T, b []byte) []cpioEntry {
	t.Helper()
	var entries []cpioEntry
	pos := 0
	for pos < len(b) {
		if pos+110 > len(b) {
			t.Fatalf("truncated header at %d", pos)
		}
		if string(b[pos:pos+6]) != "070701" {
			t.Fatalf("bad magic at %d: %q", pos, b[pos:pos+6])
		}
		field := func(i int) uint32 {
			start := pos + 6 + i*8
			var v uint32
			for _, c := range b[start : start+8] {
				v <<= 4
				switch {
				case c >= '0' && c <= '9':
					v |= uint32(c - '0')
				case c >= 'a' && c <= 'f':
					v |= uint32(c-'a') + 10
				case c >= 'A' && c <= 'F':
					v |= uint32(c-'A') + 10
				}
			}
			return v
		}
		mode := field(1)
		filesize := field(6)
		namesize := field(11)
		nameStart := pos + 110
		name := string(b[nameStart : nameStart+int(namesize)-1])
		dataStart := nameStart + int(namesize)
		dataStart += (4 - (dataStart % 4)) % 4
		data := b[dataStart : dataStart+int(filesize)]
		next := dataStart + int(filesize)
		next += (4 - (next % 4)) % 4
		pos = next
		if name == "TRAILER!!!" {
			break
		}
		entries = append(entries, cpioEntry{name: name, mode: mode, data: append([]byte(nil), data...)})
	}
	return entries
}

func find(entries []cpioEntry, name string) (cpioEntry, bool) {
	for _, e := range entries {
		if e.name == name {
			return e, true
		}
	}
	return cpioEntry{}, false
}

func TestPackTree(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "etc"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "etc", "hostname"), []byte("vm"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("etc/hostname", filepath.Join(root, "link")); err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	extra := Entry{Name: "config.json", Mode: 0o100644, Data: []byte("{}")}
	if err := PackTree(&buf, root, extra); err != nil {
		t.Fatalf("PackTree: %v", err)
	}

	entries := parseNewc(t, buf.Bytes())
	if e, ok := find(entries, "etc/hostname"); !ok || string(e.data) != "vm" {
		t.Errorf("etc/hostname = %q, ok=%v", e.data, ok)
	}
	if e, ok := find(entries, "link"); !ok || string(e.data) != "etc/hostname" || e.mode&0o120000 == 0 {
		t.Errorf("symlink link = %q mode=%o ok=%v", e.data, e.mode, ok)
	}
	if e, ok := find(entries, "config.json"); !ok || string(e.data) != "{}" {
		t.Errorf("extra config.json = %q ok=%v", e.data, ok)
	}
	// The extra must come after the walked entries.
	if entries[len(entries)-1].name != "config.json" {
		t.Errorf("last entry = %q, want config.json (extras append last)", entries[len(entries)-1].name)
	}
}

func TestPackTreeDeterministic(t *testing.T) {
	root := t.TempDir()
	for _, n := range []string{"a", "b", "c"} {
		if err := os.WriteFile(filepath.Join(root, n), []byte(n), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	var a, b bytes.Buffer
	if err := PackTree(&a, root); err != nil {
		t.Fatal(err)
	}
	if err := PackTree(&b, root); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(a.Bytes(), b.Bytes()) {
		t.Error("PackTree is not deterministic")
	}
}

func buildTar(t *testing.T, gz bool) []byte {
	t.Helper()
	var raw bytes.Buffer
	tw := tar.NewWriter(&raw)
	writeReg := func(name, body string) {
		tw.WriteHeader(&tar.Header{Name: name, Typeflag: tar.TypeReg, Mode: 0o644, Size: int64(len(body))})
		tw.Write([]byte(body))
	}
	tw.WriteHeader(&tar.Header{Name: "dir/", Typeflag: tar.TypeDir, Mode: 0o755})
	writeReg("dir/file", "hello")
	tw.WriteHeader(&tar.Header{Name: "dev/null", Typeflag: tar.TypeChar, Mode: 0o666, Devmajor: 1, Devminor: 3})
	tw.Close()
	if !gz {
		return raw.Bytes()
	}
	var zbuf bytes.Buffer
	zw := gzip.NewWriter(&zbuf)
	zw.Write(raw.Bytes())
	zw.Close()
	return zbuf.Bytes()
}

func TestPackTar(t *testing.T) {
	for _, gz := range []bool{false, true} {
		name := "plain"
		if gz {
			name = "gzip"
		}
		t.Run(name, func(t *testing.T) {
			data := buildTar(t, gz)
			var buf bytes.Buffer
			extra := Entry{Name: "init", Mode: 0o100755, Data: []byte("#!/bin/sh")}
			if err := PackTar(&buf, bytes.NewReader(data), extra); err != nil {
				t.Fatalf("PackTar: %v", err)
			}
			entries := parseNewc(t, buf.Bytes())
			if e, ok := find(entries, "dir/file"); !ok || string(e.data) != "hello" {
				t.Errorf("dir/file = %q ok=%v", e.data, ok)
			}
			if e, ok := find(entries, "dev/null"); !ok || e.mode&0o020000 == 0 {
				t.Errorf("dev/null mode=%o ok=%v, want char device", e.mode, ok)
			}
			if e, ok := find(entries, "init"); !ok || e.mode != 0o100755 {
				t.Errorf("init mode=%o ok=%v", e.mode, ok)
			}
		})
	}
}

func TestTarReaderDetectsGzip(t *testing.T) {
	data := buildTar(t, true)
	tr, err := TarReader(bufio.NewReader(bytes.NewReader(data)))
	if err != nil {
		t.Fatal(err)
	}
	hdr, err := tr.Next()
	if err != nil {
		t.Fatalf("Next: %v", err)
	}
	if hdr.Name != "dir/" {
		t.Errorf("first entry = %q, want dir/", hdr.Name)
	}
}

func TestCleanName(t *testing.T) {
	tests := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{"foo/bar", "foo/bar", false},
		{"/foo/bar", "foo/bar", false},
		{"./foo", "foo", false},
		{"foo/../bar", "bar", false},
		{"", "", true},
		{".", "", true},
		{"foo\x00bar", "", true},
	}
	for _, tt := range tests {
		got, err := CleanName(tt.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("CleanName(%q) err = %v, wantErr %v", tt.in, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("CleanName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestPackDefaultsZeroMode(t *testing.T) {
	var buf bytes.Buffer
	if err := Pack(&buf, []Entry{{Name: "f", Data: []byte("x")}}); err != nil {
		t.Fatal(err)
	}
	entries := parseNewc(t, buf.Bytes())
	if e, ok := find(entries, "f"); !ok || e.mode != 0o100644 {
		t.Errorf("zero-mode entry got mode %o, want 0100644", e.mode)
	}
}
