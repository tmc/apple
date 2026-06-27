package initramfs

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Entry is one filesystem entry in a newc initramfs archive. Mode is a full
// st_mode (type bits OR permission bits); for device nodes set RDevMajor and
// RDevMinor. Data holds a regular file's contents or a symlink's target.
type Entry struct {
	Name      string
	Mode      uint32
	UID       uint32
	GID       uint32
	RDevMajor uint32
	RDevMinor uint32
	Data      []byte
}

// PackTree writes the directory tree at root as a newc initramfs, then appends
// extras (e.g. config files or injected binaries) before the trailer.
func PackTree(w io.Writer, root string, extras ...Entry) error {
	var entries []Entry
	err := filepath.WalkDir(root, func(name string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if name == root {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, name)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		mode := uint32(info.Mode().Perm())
		switch {
		case info.Mode().IsDir():
			mode |= 0o040000
			entries = append(entries, Entry{Name: rel, Mode: mode})
		case info.Mode().Type() == os.ModeSymlink:
			target, err := os.Readlink(name)
			if err != nil {
				return err
			}
			mode |= 0o120000
			entries = append(entries, Entry{Name: rel, Mode: mode, Data: []byte(target)})
		case info.Mode().IsRegular():
			data, err := os.ReadFile(name)
			if err != nil {
				return err
			}
			mode |= 0o100000
			entries = append(entries, Entry{Name: rel, Mode: mode, Data: data})
		}
		return nil
	})
	if err != nil {
		return err
	}
	entries = append(entries, extras...)
	return Pack(w, entries)
}

// PackTar writes a (possibly gzipped) tar rootfs from r as a newc initramfs,
// then appends extras before the trailer. Regular files are streamed rather
// than buffered.
func PackTar(w io.Writer, r io.Reader, extras ...Entry) error {
	br := bufio.NewReader(r)
	tr, err := TarReader(br)
	if err != nil {
		return err
	}
	ino := uint32(1)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read rootfs tar: %w", err)
		}
		name, err := CleanName(hdr.Name)
		if err != nil {
			if isEmptyEntryName(err) {
				continue
			}
			return err
		}
		mode := uint32(hdr.Mode & 0o7777)
		switch hdr.Typeflag {
		case tar.TypeDir:
			mode |= 0o040000
			if err := writeNewcEntry(w, ino, name, mode, uint32(hdr.Uid), uint32(hdr.Gid), 0, 0, nil); err != nil {
				return err
			}
		case tar.TypeReg:
			if hdr.Size < 0 || hdr.Size > int64(^uint32(0)) {
				return fmt.Errorf("rootfs tar entry %q has unsupported size %d", name, hdr.Size)
			}
			mode |= 0o100000
			if err := writeNewcEntryHeader(w, ino, name, mode, uint32(hdr.Uid), uint32(hdr.Gid), 0, 0, uint32(hdr.Size)); err != nil {
				return err
			}
			n, err := io.CopyN(w, tr, hdr.Size)
			if err != nil {
				return fmt.Errorf("read rootfs tar entry %q: %w", name, err)
			}
			if n != hdr.Size {
				return fmt.Errorf("read rootfs tar entry %q: short read", name)
			}
			if err := writePad(w, int(hdr.Size)); err != nil {
				return err
			}
		case tar.TypeSymlink:
			mode |= 0o120000
			if err := writeNewcEntry(w, ino, name, mode, uint32(hdr.Uid), uint32(hdr.Gid), 0, 0, []byte(hdr.Linkname)); err != nil {
				return err
			}
		case tar.TypeLink:
			mode |= 0o120000
			if err := writeNewcEntry(w, ino, name, mode, uint32(hdr.Uid), uint32(hdr.Gid), 0, 0, []byte(hdr.Linkname)); err != nil {
				return err
			}
		case tar.TypeChar:
			ent, err := deviceEntry(name, mode|0o020000, hdr)
			if err != nil {
				return err
			}
			if err := writeNewcEntry(w, ino, ent.Name, ent.Mode, ent.UID, ent.GID, ent.RDevMajor, ent.RDevMinor, nil); err != nil {
				return err
			}
		case tar.TypeBlock:
			ent, err := deviceEntry(name, mode|0o060000, hdr)
			if err != nil {
				return err
			}
			if err := writeNewcEntry(w, ino, ent.Name, ent.Mode, ent.UID, ent.GID, ent.RDevMajor, ent.RDevMinor, nil); err != nil {
				return err
			}
		case tar.TypeFifo:
			if err := writeNewcEntry(w, ino, name, mode|0o010000, uint32(hdr.Uid), uint32(hdr.Gid), 0, 0, nil); err != nil {
				return err
			}
		default:
			continue
		}
		ino++
	}
	for _, ent := range extras {
		name, err := CleanName(ent.Name)
		if err != nil {
			return err
		}
		if err := writeNewcEntry(w, ino, name, entryMode(ent.Mode), ent.UID, ent.GID, ent.RDevMajor, ent.RDevMinor, ent.Data); err != nil {
			return err
		}
		ino++
	}
	return writeNewcEntry(w, 0, "TRAILER!!!", 0, 0, 0, 0, 0, nil)
}

// TarReader returns a tar reader over r, transparently decompressing gzip input.
func TarReader(r *bufio.Reader) (*tar.Reader, error) {
	magic, err := r.Peek(2)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("read rootfs tar header: %w", err)
	}
	var rr io.Reader = r
	if len(magic) == 2 && magic[0] == 0x1f && magic[1] == 0x8b {
		gz, err := gzip.NewReader(r)
		if err != nil {
			return nil, fmt.Errorf("read rootfs gzip header: %w", err)
		}
		rr = gz
	}
	return tar.NewReader(rr), nil
}

func deviceEntry(name string, mode uint32, hdr *tar.Header) (Entry, error) {
	if hdr.Devmajor < 0 || hdr.Devminor < 0 {
		return Entry{}, fmt.Errorf("rootfs tar entry %q has negative device number", name)
	}
	return Entry{
		Name:      name,
		Mode:      mode,
		UID:       uint32(hdr.Uid),
		GID:       uint32(hdr.Gid),
		RDevMajor: uint32(hdr.Devmajor),
		RDevMinor: uint32(hdr.Devminor),
	}, nil
}

// Pack writes entries as a deterministic SVR4 newc cpio archive, in order,
// followed by the trailer.
func Pack(w io.Writer, entries []Entry) error {
	for i, ent := range entries {
		name, err := CleanName(ent.Name)
		if err != nil {
			return err
		}
		if err := writeNewcEntry(w, uint32(i+1), name, entryMode(ent.Mode), ent.UID, ent.GID, ent.RDevMajor, ent.RDevMinor, ent.Data); err != nil {
			return err
		}
	}
	return writeNewcEntry(w, 0, "TRAILER!!!", 0, 0, 0, 0, 0, nil)
}

// entryMode defaults a zero mode to a regular file with 0644 permissions.
func entryMode(mode uint32) uint32 {
	if mode == 0 {
		return 0o100644
	}
	return mode
}

// CleanName returns a cpio-safe relative entry name, rejecting empty names and
// names containing a NUL byte.
func CleanName(name string) (string, error) {
	name = strings.TrimPrefix(path.Clean("/"+name), "/")
	if name == "." || name == "" {
		return "", errEmptyEntryName
	}
	if strings.Contains(name, "\x00") {
		return "", fmt.Errorf("initramfs entry %q contains NUL", name)
	}
	return name, nil
}

var errEmptyEntryName = errors.New("initramfs entry name is empty")

func isEmptyEntryName(err error) bool {
	return err == errEmptyEntryName
}

func writeNewcEntry(w io.Writer, ino uint32, name string, mode uint32, uid uint32, gid uint32, rdevmajor uint32, rdevminor uint32, data []byte) error {
	if err := writeNewcEntryHeader(w, ino, name, mode, uid, gid, rdevmajor, rdevminor, uint32(len(data))); err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	return writePad(w, len(data))
}

func writeNewcEntryHeader(w io.Writer, ino uint32, name string, mode uint32, uid uint32, gid uint32, rdevmajor uint32, rdevminor uint32, filesize uint32) error {
	namesize := uint32(len(name) + 1)
	fields := []uint32{
		ino,
		mode,
		uid,
		gid,
		1, // nlink
		0, // mtime
		filesize,
		0, 0, rdevmajor, rdevminor,
		namesize,
		0,
	}
	if _, err := io.WriteString(w, "070701"); err != nil {
		return err
	}
	for _, field := range fields {
		if _, err := fmt.Fprintf(w, "%08x", field); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, name+"\x00"); err != nil {
		return err
	}
	return writePad(w, 110+int(namesize))
}

func writePad(w io.Writer, n int) error {
	pad := (4 - (n % 4)) % 4
	if pad == 0 {
		return nil
	}
	_, err := w.Write([]byte{0, 0, 0}[:pad])
	return err
}
