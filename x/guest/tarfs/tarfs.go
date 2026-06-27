// Package tarfs extracts tar streams into a directory safely, with optional
// OCI image-layer (whiteout) semantics.
//
// All extraction rejects entries that would escape the destination directory
// via ".." traversal or symlinked parents (Zip Slip). [Unpack] extracts a plain
// (optionally gzipped) tar; [UnpackLayer] additionally applies OCI whiteouts so
// a sequence of UnpackLayer calls into the same directory reconstructs an image
// rootfs.
package tarfs

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Unpack extracts the tar stream in r into dest. If the stream is gzipped it is
// decompressed first. Entries that would escape dest are rejected.
func Unpack(ctx context.Context, r io.Reader, dest string) error {
	return extract(ctx, r, dest, false)
}

// UnpackLayer is like [Unpack] but applies OCI whiteouts: an entry named
// ".wh.<name>" deletes the sibling <name> already present in dest, and an entry
// named ".wh..wh..opq" clears the existing contents of its directory. Applying
// image layers in order with UnpackLayer reconstructs the merged rootfs.
func UnpackLayer(ctx context.Context, r io.Reader, dest string) error {
	return extract(ctx, r, dest, true)
}

func extract(ctx context.Context, r io.Reader, dest string, whiteouts bool) error {
	br, err := maybeGunzip(r)
	if err != nil {
		return err
	}
	tr := tar.NewReader(br)
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if whiteouts {
			done, err := applyWhiteout(dest, header.Name)
			if err != nil {
				return err
			}
			if done {
				continue
			}
		}

		if err := writeEntry(dest, header, tr); err != nil {
			return err
		}
	}
	return nil
}

// applyWhiteout handles OCI whiteout markers. It reports whether name was a
// whiteout entry (and thus fully handled); a false return means name is a
// regular entry the caller should extract.
func applyWhiteout(dest, name string) (bool, error) {
	base := filepath.Base(name)
	dir := filepath.Dir(name)

	// Opaque whiteout: clear the directory's existing contents.
	if base == ".wh..wh..opq" {
		targetDir, err := safeJoin(dest, dir)
		if err != nil {
			return false, err
		}
		entries, err := os.ReadDir(targetDir)
		if err == nil {
			for _, entry := range entries {
				if err := os.RemoveAll(filepath.Join(targetDir, entry.Name())); err != nil {
					return false, err
				}
			}
		}
		return true, nil
	}

	// Explicit whiteout: delete the named sibling.
	if strings.HasPrefix(base, ".wh.") {
		toDelete := strings.TrimPrefix(base, ".wh.")
		target, err := safeJoin(dest, filepath.Join(dir, toDelete))
		if err != nil {
			return false, err
		}
		if err := os.RemoveAll(target); err != nil && !os.IsNotExist(err) {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func writeEntry(dest string, header *tar.Header, tr io.Reader) error {
	target, err := safeJoin(dest, header.Name)
	if err != nil {
		return err
	}

	switch header.Typeflag {
	case tar.TypeDir:
		if err := mkdirAll(dest, header.Name, os.FileMode(header.Mode)); err != nil {
			return err
		}

	case tar.TypeReg, tar.TypeRegA:
		if err := mkdirAll(dest, filepath.Dir(header.Name), 0o755); err != nil {
			return err
		}
		if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
			return err
		}
		f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(header.Mode))
		if err != nil {
			return err
		}
		if _, err := io.Copy(f, tr); err != nil {
			f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}

	case tar.TypeSymlink:
		if err := mkdirAll(dest, filepath.Dir(header.Name), 0o755); err != nil {
			return err
		}
		if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := os.Symlink(header.Linkname, target); err != nil {
			return err
		}

	case tar.TypeLink:
		if err := mkdirAll(dest, filepath.Dir(header.Name), 0o755); err != nil {
			return err
		}
		if err := os.Remove(target); err != nil && !os.IsNotExist(err) {
			return err
		}
		// Hard-link target is relative to the tar root; resolve under dest.
		linkTarget, err := safeJoin(dest, header.Linkname)
		if err != nil {
			return err
		}
		if err := os.Link(linkTarget, target); err != nil {
			return err
		}

	default:
		// Unsupported type (fifo, device, etc.): skip, matching donor behavior.
		return nil
	}

	applyMetadata(target, header)
	return nil
}

// applyMetadata restores ownership and modification time on an extracted entry.
// Both are best-effort: extracting an image rootfs as an unprivileged user
// cannot chown to arbitrary uid/gid (EPERM), and that is expected, not fatal.
// Timestamps are not applied to symlinks (os.Chtimes follows the link).
func applyMetadata(target string, header *tar.Header) {
	os.Lchown(target, header.Uid, header.Gid)
	if header.Typeflag != tar.TypeSymlink {
		os.Chtimes(target, header.AccessTime, header.ModTime)
	}
}

// maybeGunzip wraps r in a gzip reader if it begins with the gzip magic bytes.
func maybeGunzip(r io.Reader) (io.Reader, error) {
	// Peek the first two bytes without consuming them.
	var magic [2]byte
	n, err := io.ReadFull(r, magic[:])
	switch {
	case err == io.EOF:
		return strings.NewReader(""), nil
	case err == io.ErrUnexpectedEOF:
		return io.MultiReader(bytes.NewReader(magic[:n]), r), nil
	case err != nil:
		return nil, err
	}
	combined := io.MultiReader(bytes.NewReader(magic[:]), r)
	if magic[0] == 0x1f && magic[1] == 0x8b {
		return gzip.NewReader(combined)
	}
	return combined, nil
}

// safeJoin returns dest/name, rejecting any name that would escape dest via
// ".." traversal (Zip Slip).
func safeJoin(dest, name string) (string, error) {
	target := filepath.Join(dest, name)
	if !strings.HasPrefix(filepath.Clean(target)+string(filepath.Separator), filepath.Clean(dest)+string(filepath.Separator)) {
		return "", fmt.Errorf("tarfs: entry %q would escape destination directory", name)
	}
	return target, nil
}

// mkdirAll creates dest/name and its parents, refusing to follow a symlinked
// path component (which could redirect a write outside dest).
func mkdirAll(dest, name string, mode os.FileMode) error {
	target, err := safeJoin(dest, name)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(filepath.Clean(dest), filepath.Clean(target))
	if err != nil {
		return fmt.Errorf("tarfs: relative path %s: %w", target, err)
	}
	if rel == "." {
		return nil
	}
	current := filepath.Clean(dest)
	for _, elem := range strings.Split(rel, string(filepath.Separator)) {
		current = filepath.Join(current, elem)
		info, err := os.Lstat(current)
		if os.IsNotExist(err) {
			if err := os.Mkdir(current, mode); err != nil && !os.IsExist(err) {
				return fmt.Errorf("tarfs: mkdir %s: %w", current, err)
			}
			continue
		}
		if err != nil {
			return fmt.Errorf("tarfs: lstat %s: %w", current, err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("tarfs: refusing to follow symlink %s", current)
		}
		if !info.IsDir() {
			return fmt.Errorf("tarfs: not a directory: %s", current)
		}
	}
	return nil
}
