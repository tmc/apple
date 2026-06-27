package tarfs

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// entry describes one tar member for the test builder.
type entry struct {
	name     string
	typeflag byte
	mode     int64
	body     string
	linkname string
	modTime  time.Time
}

func buildTar(t *testing.T, entries []entry, gz bool) []byte {
	t.Helper()
	var buf bytes.Buffer
	var w *tar.Writer
	var gzw *gzip.Writer
	if gz {
		gzw = gzip.NewWriter(&buf)
		w = tar.NewWriter(gzw)
	} else {
		w = tar.NewWriter(&buf)
	}
	for _, e := range entries {
		mode := e.mode
		if mode == 0 {
			mode = 0o644
		}
		h := &tar.Header{
			Name:     e.name,
			Typeflag: e.typeflag,
			Mode:     mode,
			Size:     int64(len(e.body)),
			Linkname: e.linkname,
			ModTime:  e.modTime,
		}
		if e.typeflag == tar.TypeDir && e.mode == 0 {
			h.Mode = 0o755
		}
		if err := w.WriteHeader(h); err != nil {
			t.Fatalf("WriteHeader %q: %v", e.name, err)
		}
		if len(e.body) > 0 {
			if _, err := w.Write([]byte(e.body)); err != nil {
				t.Fatalf("Write %q: %v", e.name, err)
			}
		}
	}
	if err := w.Close(); err != nil {
		t.Fatalf("tar close: %v", err)
	}
	if gz {
		if err := gzw.Close(); err != nil {
			t.Fatalf("gzip close: %v", err)
		}
	}
	return buf.Bytes()
}

func TestUnpack(t *testing.T) {
	for _, gz := range []bool{false, true} {
		gz := gz
		name := "plain"
		if gz {
			name = "gzip"
		}
		t.Run(name, func(t *testing.T) {
			dest := t.TempDir()
			data := buildTar(t, []entry{
				{name: "dir/", typeflag: tar.TypeDir},
				{name: "dir/file.txt", typeflag: tar.TypeReg, body: "hello", mode: 0o600},
				{name: "dir/link", typeflag: tar.TypeSymlink, linkname: "file.txt"},
			}, gz)

			if err := Unpack(context.Background(), bytes.NewReader(data), dest); err != nil {
				t.Fatalf("Unpack: %v", err)
			}

			got, err := os.ReadFile(filepath.Join(dest, "dir", "file.txt"))
			if err != nil {
				t.Fatalf("read file: %v", err)
			}
			if string(got) != "hello" {
				t.Errorf("file body = %q, want %q", got, "hello")
			}
			fi, err := os.Stat(filepath.Join(dest, "dir", "file.txt"))
			if err != nil {
				t.Fatal(err)
			}
			if fi.Mode().Perm() != 0o600 {
				t.Errorf("file mode = %o, want 600", fi.Mode().Perm())
			}
			link, err := os.Readlink(filepath.Join(dest, "dir", "link"))
			if err != nil {
				t.Fatalf("readlink: %v", err)
			}
			if link != "file.txt" {
				t.Errorf("symlink target = %q, want %q", link, "file.txt")
			}
		})
	}
}

func TestUnpackRejectsEscape(t *testing.T) {
	tests := []struct {
		name  string
		entry entry
	}{
		{"dotdot file", entry{name: "../escape.txt", typeflag: tar.TypeReg, body: "x"}},
		{"dotdot dir", entry{name: "../evil/", typeflag: tar.TypeDir}},
		{"nested dotdot", entry{name: "a/../../escape.txt", typeflag: tar.TypeReg, body: "x"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dest := t.TempDir()
			data := buildTar(t, []entry{tt.entry}, false)
			err := Unpack(context.Background(), bytes.NewReader(data), dest)
			if err == nil {
				t.Fatalf("expected escape rejection, got nil")
			}
			// Ensure nothing was written outside dest.
			if _, statErr := os.Stat(filepath.Join(filepath.Dir(dest), "escape.txt")); statErr == nil {
				t.Fatalf("escape.txt was written outside dest")
			}
		})
	}
}

func TestUnpackLayerExplicitWhiteout(t *testing.T) {
	dest := t.TempDir()
	// Lower layer: a file that a later layer deletes.
	lower := buildTar(t, []entry{
		{name: "a.txt", typeflag: tar.TypeReg, body: "lower"},
		{name: "keep.txt", typeflag: tar.TypeReg, body: "keep"},
	}, false)
	if err := UnpackLayer(context.Background(), bytes.NewReader(lower), dest); err != nil {
		t.Fatalf("lower: %v", err)
	}
	// Upper layer: whiteout a.txt.
	upper := buildTar(t, []entry{
		{name: ".wh.a.txt", typeflag: tar.TypeReg},
	}, false)
	if err := UnpackLayer(context.Background(), bytes.NewReader(upper), dest); err != nil {
		t.Fatalf("upper: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dest, "a.txt")); !os.IsNotExist(err) {
		t.Errorf("a.txt should have been whited out, stat err = %v", err)
	}
	if _, err := os.Stat(filepath.Join(dest, "keep.txt")); err != nil {
		t.Errorf("keep.txt should survive, got %v", err)
	}
}

func TestUnpackLayerOpaqueWhiteout(t *testing.T) {
	dest := t.TempDir()
	lower := buildTar(t, []entry{
		{name: "d/", typeflag: tar.TypeDir},
		{name: "d/x.txt", typeflag: tar.TypeReg, body: "x"},
		{name: "d/y.txt", typeflag: tar.TypeReg, body: "y"},
	}, false)
	if err := UnpackLayer(context.Background(), bytes.NewReader(lower), dest); err != nil {
		t.Fatalf("lower: %v", err)
	}
	// Opaque whiteout clears d/, then adds a fresh file.
	upper := buildTar(t, []entry{
		{name: "d/.wh..wh..opq", typeflag: tar.TypeReg},
		{name: "d/z.txt", typeflag: tar.TypeReg, body: "z"},
	}, false)
	if err := UnpackLayer(context.Background(), bytes.NewReader(upper), dest); err != nil {
		t.Fatalf("upper: %v", err)
	}

	for _, gone := range []string{"d/x.txt", "d/y.txt"} {
		if _, err := os.Stat(filepath.Join(dest, gone)); !os.IsNotExist(err) {
			t.Errorf("%s should be cleared by opaque whiteout, stat err = %v", gone, err)
		}
	}
	if _, err := os.Stat(filepath.Join(dest, "d", "z.txt")); err != nil {
		t.Errorf("d/z.txt should exist after opaque whiteout, got %v", err)
	}
}

func TestUnpackPreservesModTime(t *testing.T) {
	dest := t.TempDir()
	want := time.Date(2021, 3, 4, 5, 6, 7, 0, time.UTC)
	data := buildTar(t, []entry{
		{name: "stamped.txt", typeflag: tar.TypeReg, body: "x", modTime: want},
	}, false)
	if err := Unpack(context.Background(), bytes.NewReader(data), dest); err != nil {
		t.Fatalf("Unpack: %v", err)
	}
	fi, err := os.Stat(filepath.Join(dest, "stamped.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if !fi.ModTime().Equal(want) {
		t.Errorf("mod time = %v, want %v", fi.ModTime().UTC(), want)
	}
}

func TestUnpackLegacyRegType(t *testing.T) {
	dest := t.TempDir()
	data := buildTar(t, []entry{
		{name: "old.txt", typeflag: tar.TypeRegA, body: "legacy"},
	}, false)
	if err := Unpack(context.Background(), bytes.NewReader(data), dest); err != nil {
		t.Fatalf("Unpack: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dest, "old.txt"))
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(got) != "legacy" {
		t.Errorf("body = %q, want %q", got, "legacy")
	}
}

func TestUnpackContextCancel(t *testing.T) {
	dest := t.TempDir()
	data := buildTar(t, []entry{{name: "a.txt", typeflag: tar.TypeReg, body: "x"}}, false)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Unpack(ctx, bytes.NewReader(data), dest); err == nil {
		t.Fatal("expected context cancellation error")
	}
}
