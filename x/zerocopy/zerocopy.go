package zerocopy

import "fmt"

// Check verifies that w and r are views of the same memory: it writes
// sentinel bytes through w at several offsets and confirms each write is
// observed through r, restoring w's original contents before returning.
// A non-nil error means r holds a copy (or a stale snapshot) of w.
//
// r must be at least as long as w; probed offsets refer to the same
// underlying byte in both views.
func Check(w, r []byte) error {
	return CheckFunc(w, func() ([]byte, error) { return r, nil })
}

// CheckFunc is Check for consumers whose view must be re-read or
// synchronized after each write. read performs whatever synchronization
// the consumer needs (evaluate a lazy array, message another process)
// and returns the consumer's current view of the same bytes as w.
// read is called at least twice per probed offset, so it must be
// repeatable.
func CheckFunc(w []byte, read func() ([]byte, error)) error {
	if len(w) == 0 {
		return fmt.Errorf("empty buffer")
	}
	for _, off := range probeOffsets(len(w)) {
		orig := w[off]
		w[off] = ^orig
		if err := expect(read, off, ^orig, "sentinel write"); err != nil {
			w[off] = orig
			return err
		}
		w[off] = orig
		// Verifying the restore too catches a reader that snapshots
		// lazily: its first read could observe the sentinel by luck
		// and then never move again.
		if err := expect(read, off, orig, "restore"); err != nil {
			return err
		}
	}
	return nil
}

func expect(read func() ([]byte, error), off int, want byte, step string) error {
	r, err := read()
	if err != nil {
		return fmt.Errorf("read after %s at offset %d: %w", step, off, err)
	}
	if off >= len(r) {
		return fmt.Errorf("reader view too short: %d bytes, need offset %d", len(r), off)
	}
	if r[off] != want {
		return fmt.Errorf("%s at offset %d not observed: wrote %#02x, read %#02x — reader holds a copy", step, off, want, r[off])
	}
	return nil
}

// probeOffsets picks the bytes to probe: both ends plus interior points,
// deduplicated and ordered.
func probeOffsets(n int) []int {
	offs := []int{0, n / 3, n / 2, n - 1}
	var out []int
	for _, o := range offs {
		if len(out) == 0 || o > out[len(out)-1] {
			out = append(out, o)
		}
	}
	return out
}
