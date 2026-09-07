package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"

	"github.com/tmc/apple/x/machdebug"
)

// decision is the outcome of inspecting one validation.
type decision struct {
	patch  bool
	reason string
}

// decide chooses whether to patch a verdict, failing closed on every doubt.
// It patches only when the validator object is non-null, its signing
// identifier was read cleanly, and that identifier equals the nominated
// identity exactly. A tool that patches when it is confused silently disables
// code signing for the whole machine, so anything short of a clean, exact
// match leaves the verdict standing.
func decide(self uint64, identity string, identErr error, want string) decision {
	switch {
	case self == 0:
		return decision{false, "validator object pointer is null"}
	case identErr != nil:
		return decision{false, fmt.Sprintf("signing identifier unreadable: %v", identErr)}
	case identity != want:
		return decision{false, fmt.Sprintf("signing identifier %q does not match %q", identity, want)}
	default:
		return decision{true, fmt.Sprintf("signing identifier %q matches", identity)}
	}
}

// hookLoop drives the one-hook design: break at the return of
// validateWithError:, read the validated binary's signing identifier, and — for
// the one nominated identity — write the result ivars and the return register
// directly. It re-arms after every non-matching verdict and returns once it
// patches, or with an error once the watchdog or a signal detaches the target
// (which unblocks Wait).
func hookLoop(p *machdebug.Process, a *amfi, want string) (bool, error) {
	for {
		// Arm entry as a one-shot each round; this needs only machdebug's
		// one-shot primitive, not persistent breakpoints. At entry, x0 is the
		// receiver and LR is the return address.
		if _, err := p.SetOneShot(a.imp); err != nil {
			return false, fmt.Errorf("arm entry breakpoint: %w", err)
		}
		entry, err := p.Wait()
		if err != nil {
			return false, err
		}
		if entry.Kind != machdebug.EventBreakpoint || entry.Addr != a.imp {
			continue
		}
		eregs, err := entry.Thread.Regs()
		if err != nil {
			return false, fmt.Errorf("read registers at entry: %w", err)
		}
		self := eregs.X[0]
		entrySP := eregs.SP
		ret := eregs.LR

		// A one-shot at the return address, matched by the entry SP: when the
		// frame unwinds, SP is restored to its entry value, so an SP match
		// identifies our own frame returning.
		if _, err := p.SetOneShot(ret); err != nil {
			return false, fmt.Errorf("arm return breakpoint: %w", err)
		}
		rev, err := p.Wait()
		if err != nil {
			return false, err
		}
		if rev.Kind != machdebug.EventBreakpoint || rev.Addr != ret {
			// Some other validation stopped here; re-arm and keep waiting.
			continue
		}
		rregs, err := rev.Thread.Regs()
		if err != nil {
			return false, fmt.Errorf("read registers at return: %w", err)
		}
		if rregs.SP != entrySP {
			// A different frame returning to the same address; not ours.
			continue
		}

		identPtr, err := readPtr(p, self+uint64(a.offSigningIdentifier))
		var ident string
		var identErr error
		if err != nil {
			identErr = err
		} else {
			ident, identErr = readSigningIdentifier(p, identPtr)
		}

		d := decide(self, ident, identErr, want)
		log.Printf("validation: self=%#x identity=%q -> %s", self, ident, d.reason)
		if !d.patch {
			continue
		}

		if err := a.patch(p, self, rev.Thread, rregs); err != nil {
			return false, fmt.Errorf("patch verdict: %w", err)
		}
		return true, nil
	}
}

// patch overwrites the validator's result ivars so the verdict reads
// Apple-signed and unrestricted, then sets the return register so the immediate
// caller sees success. Writes to shared-cache-backed ivars are private to the
// target via copy-on-write inside machdebug.
func (a *amfi) patch(p *machdebug.Process, self uint64, t machdebug.Thread, regs *machdebug.Regs) error {
	writes := []struct {
		off int64
		val byte
	}{
		{a.offIsValid, 1},
		{a.offIsApple, 1},
		{a.offShouldUnrestrict, 1},
		{a.offHasRestrictedEntitlements, 0},
	}
	for _, w := range writes {
		if _, err := p.WriteAt([]byte{w.val}, int64(self)+w.off); err != nil {
			return fmt.Errorf("write ivar at +%d: %w", w.off, err)
		}
	}
	// Set x0=1 so the caller that reads the return value sees success. Mutate
	// the register state read at this stop and write it back, per machdebug's
	// PAC rule; never synthesize the struct.
	regs.X[0] = 1
	if err := t.SetRegs(regs); err != nil {
		return fmt.Errorf("set x0: %w", err)
	}
	return nil
}

// readPtr reads an 8-byte little-endian pointer at addr from r.
func readPtr(r io.ReaderAt, addr uint64) (uint64, error) {
	var b [8]byte
	if _, err := r.ReadAt(b[:], int64(addr)); err != nil {
		return 0, fmt.Errorf("read pointer at %#x: %w", addr, err)
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}
