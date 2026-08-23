package e5rt

import (
	"errors"
	"fmt"
)

// A StateLayout describes how a rank-4 fp16 MIL state tensor of shape
// [1, Channels, 1, Dim] occupies a bound buffer object.
//
// It is not the packed element count. Each channel starts on a 64-byte
// boundary, so a state of shape [1, 4, 1, 4] occupies 256 bytes where its 16
// fp16 elements would pack into 32. Use [StateLayout.Size] to allocate and
// [StateLayout.Offset] to address elements; the trailing bytes of each row are
// padding the engine does not read.
//
// This matters because getting it wrong is silent. A caller who allocates the
// packed size receives no error from [Lib.BufferObjectAlloc] or from
// [Lib.IOPortBindBufferObject]. The engine addresses the padded offsets
// regardless, so a program that writes its state writes past the end of the
// allocation — and because a reader reads back the same out-of-bounds offsets,
// the round trip is self-consistent and the corruption does not show up as a
// wrong answer. That is exactly how it went unnoticed in this package's own
// probes. [Lib.BindStatePort] is the checked bind that refuses it.
//
// # What has been measured
//
// The 64-byte row alignment was measured on macOS 26.x on one machine, for
// fp16 states of shape [1, C, 1, D], at (C, D) of (4,4), (4,32), (4,40),
// (2,64) and (8,2), by binding an oversized buffer, filling every two-byte slot
// with its own index, and reading back which slot each element came from.
// Shapes whose row already fills whole 64-byte lines, such as D=32, are packed
// by coincidence; code developed at those shapes works and corrupts memory at
// others.
//
// Element types other than fp16, ranks other than 4, and leading extents other
// than 1 are UNMEASURED. So is whether the 64-byte figure is a property of this
// hardware generation, this OS version, or the E5RT ABI. Callers should compute
// sizes with this type rather than hard-coding them, and
// examples/ane/statecache -layout re-runs the measurement.
//
// Ordinary input and output ports are packed. Only the state carries padding.
type StateLayout struct {
	Channels int
	Dim      int
}

// stateRowAlign is the boundary each channel of a state tensor starts on.
const stateRowAlign = 64

// RowBytes is the distance in bytes between the starts of consecutive channels.
func (l StateLayout) RowBytes() int {
	return max((l.Dim*2+stateRowAlign-1)&^(stateRowAlign-1), stateRowAlign)
}

// Size is the number of bytes a buffer object bound to this state must hold.
func (l StateLayout) Size() int { return l.Channels * l.RowBytes() }

// Values is the number of logical elements, Channels*Dim.
func (l StateLayout) Values() int { return l.Channels * l.Dim }

// Offset returns the byte offset of element i of the given channel.
// It panics if either index is out of range, which is a programming error
// rather than a runtime condition.
func (l StateLayout) Offset(channel, i int) int {
	if channel < 0 || channel >= l.Channels {
		panic(fmt.Sprintf("e5rt: channel %d out of range [0, %d)", channel, l.Channels))
	}
	if i < 0 || i >= l.Dim {
		panic(fmt.Sprintf("e5rt: index %d out of range [0, %d)", i, l.Dim))
	}
	return channel*l.RowBytes() + i*2
}

// Packed reports whether the layout leaves no padding, which happens only when
// a channel row exactly fills a whole number of 64-byte lines. A caller must
// not rely on this being true for a particular shape; it is offered so that a
// diagnostic can say which case it is in.
func (l StateLayout) Packed() bool { return l.Dim*2 == l.RowBytes() }

// valid reports whether the layout describes a usable tensor.
func (l StateLayout) valid() error {
	if l.Channels <= 0 || l.Dim <= 0 {
		return fmt.Errorf("e5rt: state layout [1, %d, 1, %d] has a non-positive extent", l.Channels, l.Dim)
	}
	if l.Dim > 1<<20 || l.Channels > 1<<20 {
		return fmt.Errorf("e5rt: state layout [1, %d, 1, %d] is implausibly large", l.Channels, l.Dim)
	}
	return nil
}

// BindStatePort retains the named inout port of op and binds buf to it, after
// checking that buf is large enough for lay.
//
// This is [Lib.OperationRetainInoutPort] followed by
// [Lib.IOPortBindBufferObject] with the size check that neither of them
// performs. The size comes from [Lib.BufferObjectGetSize], so the caller does
// not have to track it alongside the buffer.
//
// On success the returned port is the caller's to release. On any failure
// nothing is left retained.
//
// The check is one-sided: a buffer larger than the layout requires is accepted,
// because binding a subregion of a larger allocation is a reasonable thing to
// do and this package has no way to know it is not intended.
func (l *Lib) BindStatePort(op uintptr, portName string, buf uintptr, lay StateLayout) (uintptr, error) {
	if err := lay.valid(); err != nil {
		return 0, err
	}
	size, err := l.BufferObjectGetSize(buf)
	if err != nil {
		return 0, fmt.Errorf("e5rt: get size of the buffer for state port %q: %w", portName, err)
	}
	if want := lay.Size(); size < uintptr(want) {
		return 0, fmt.Errorf("e5rt: state port %q needs %d bytes for shape [1, %d, 1, %d] "+
			"(%d channels of %d), but the buffer holds %d; a packed %d is the usual mistake",
			portName, want, lay.Channels, lay.Dim, lay.Channels, lay.RowBytes(), size, lay.Values()*2)
	}
	port, err := l.OperationRetainInoutPort(op, portName)
	if err != nil {
		return 0, fmt.Errorf("e5rt: retain state port %q: %w", portName, err)
	}
	if err := l.IOPortBindBufferObject(port, buf); err != nil {
		return 0, errors.Join(fmt.Errorf("e5rt: bind state port %q: %w", portName, err), l.IOPortRelease(port))
	}
	return port, nil
}
