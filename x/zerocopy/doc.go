// Package zerocopy verifies at runtime that a claimed zero-copy handoff
// really shares memory rather than copying it.
//
// The probe is the only evidence that distinguishes an alias from a copy:
// a fast copy has the same timing as an alias, and APIs can fall back to
// copying silently. Check writes sentinel bytes through the producer's
// view and confirms each write is observed through the consumer's view:
//
//	if err := zerocopy.Check(producer, consumer); err != nil {
//		log.Fatalf("handoff is a copy: %v", err)
//	}
//
// For consumers whose view must be re-read or synchronized after a write
// (a GPU array, another process), use CheckFunc with a read function that
// performs the synchronization and returns the consumer's current bytes.
//
// The probe proves visibility, not mechanism: a read function that
// re-copies from the producer's memory on every call is indistinguishable
// from an alias. That is impossible across a real device or process
// boundary, which is where the probe is meant to run.
package zerocopy
