// Command statecache keeps mutable state resident on the Neural Engine across
// executions, and across two separately compiled programs, through one retained
// E5RT inout port.
//
// Every other ANE example here is a pure function: the host writes an input,
// the engine computes, the host reads an output. A decoder is not that. It
// carries state between steps, and if that state has to be staged back to the
// host after every token then the state, not the arithmetic, sets the pace.
// [e5rt.Lib.OperationRetainInoutPort] is the interface that avoids it: one
// caller-owned buffer object serves as both the read and the write side of a
// MIL state parameter, and it stays bound across executions.
//
// That capability was proven by the package's own probes. Nothing had used it.
//
// # The state buffer is not packed
//
// The first thing this example had to discover is that a bound state buffer is
// not laid out the way the MIL text reads. Each channel of a
// [1, channels, 1, dim] fp16 state starts on a 64-byte boundary, so the buffer
// a caller must allocate is channels*round_up(dim*2, 64) bytes, not
// channels*dim*2. At dim=4 that is 256 bytes where the packed size is 32.
//
// A caller who allocates the packed size does not get an error. The engine
// reads and writes the padded offsets regardless, which for the state of an
// accumulate program means writing 224 bytes past the end of the allocation.
// The first version of this example did exactly that, and the symptom was that
// one channel round-tripped and the rest came back zero.
//
// Run with -layout to measure it rather than take it from this comment: that
// mode binds an oversized state buffer, fills every 2-byte slot with its own
// index, and prints the slot each tensor element was read from.
//
// [e5rt.StateLayout] now owns the arithmetic, and [e5rt.Lib.BindStatePort] is a
// bind that asks the runtime how large the buffer is and refuses one too small
// for the shape. This example uses both, so it doubles as their exercise: the
// raw retain-and-bind pair accepts an undersized buffer without complaint,
// which is why the checked path exists.
//
// The ordinary input and output ports are packed. Only the state carries the
// padding, and the -layout mode plus the exact recurrence below are what
// establish that asymmetry.
//
// # What this is, and what it is not
//
// This is one resident mutable tensor that a program reads, modifies, and
// writes back in place. It is the mechanism a KV cache needs, exercised on its
// own.
//
// It is NOT a KV cache. A cache is position-indexed — token t writes row t and
// leaves rows 0..t-1 alone — and nothing here indexes by position. The disjoint
// control below shows that updates to separate regions do not disturb each
// other, which is a necessary property of a cache and not a sufficient one.
// Writing an actual per-layer, position-indexed cache and proving it against a
// full-prefix oracle is a separate, larger piece of work.
//
// # What each check would catch
//
//   - A float64 recurrence, checked at every step. The values are chosen so the
//     whole recurrence is exact in fp16, so this requires exact agreement rather
//     than a tolerance. A program that dropped an update, applied one twice, or
//     accumulated in the wrong order fails here.
//
//   - A host-visibility probe, run before anything depends on it. This program
//     zeroes the state between arms by writing through the host pointer, which
//     is only a reset if the engine actually observes host writes to a bound
//     inout buffer. That is checked by writing a distinctive pattern and reading
//     it back through the engine, not assumed. It is also what catches a wrong
//     buffer layout, which is how the padding above was found.
//
//   - A separately compiled reader, bound to the same buffer object. The update
//     program returns the new state as an ordinary tensor output, so its own
//     output cannot distinguish stored state from an echo of what was just
//     computed. A different program that only reads can.
//
//   - The same reader program bound to a different buffer, required to return
//     zeros. Without this arm, the reader agreeing with the reference would be
//     consistent with the reader seeing the right answer by some route other
//     than the shared buffer. This is the negative polarity of the check above,
//     and the two together are what make the sharing claim load-bearing.
//
//   - Two states advanced in an interleaved order with different values, each
//     required to follow its own recurrence. A single global state behind the
//     inout port would pass every single-state check and fail this one.
//
//   - Disjoint regions updated in separate steps. An update that replaced the
//     state rather than accumulating into it, or that wrote at the wrong offset,
//     leaves a detectable hole.
//
//   - A state name the program does not declare, required to be refused, so
//     that "the inout port was retained" is evidence of resolution rather than
//     evidence that the call returns non-nil.
//
// # Why exact agreement is the right bar
//
// The state is fp16 and the update is an fp16 add, so an accumulation would
// normally drift away from a float64 reference and need a tolerance. A
// tolerance is a weaker check, and here it is avoidable: every value is a
// multiple of 0.5 and every partial sum stays well under 1024, and multiples of
// 0.5 in that range are exactly representable in binary16. A correctly rounded
// add of two exactly representable values whose exact sum is also representable
// is exact. So the entire recurrence is exact, and any disagreement at all is a
// defect rather than rounding. The flag guards below keep the inputs inside
// that range; step counts outside it are refused rather than silently switching
// to a weaker comparison.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

var (
	channels  = flag.Int("chan", 4, "channel extent of the state tensor")
	dim       = flag.Int("dim", 4, "trailing extent of the state tensor")
	steps     = flag.Int("steps", 6, "number of updates to apply (2..8, to keep the recurrence exact in fp16)")
	mapLayout = flag.Bool("layout", false, "map which host offsets the engine reads the state from, then exit")
	report    = flag.Bool("report", false, "print a machine-readable JSON layout report for mailing back (implies -layout)")
)

// stateName is the MIL state parameter these programs share. Both the update
// program and the reader declare it, which is what lets one buffer serve both.
const stateName = "kv"

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "FAIL:", err)
		os.Exit(1)
	}
	fmt.Println("\nOK")
}

func run() error {
	if err := checkFlags(); err != nil {
		return err
	}
	n := *channels * *dim
	shape := [4]int{1, *channels, 1, *dim}
	lay := e5rt.StateLayout{Channels: *channels, Dim: *dim}
	bytes := n * 2

	fmt.Printf("retained state: chan=%d dim=%d (%d values) steps=%d\n\n", *channels, *dim, n, *steps)

	lib, err := e5rt.Open()
	if err != nil {
		return fmt.Errorf("open e5rt: %w", err)
	}

	root, err := os.MkdirTemp("", "statecache")
	if err != nil {
		return err
	}
	defer os.RemoveAll(root)

	if *mapLayout || *report {
		return layoutProbe(lib, root, shape)
	}

	// The state buffer is allocated by this program and outlives every
	// operation bound to it. That ownership is the whole point: the engine
	// keeps reading and writing this one allocation.
	//
	// It is sized by the layout rather than by the packed element count. The
	// two differ whenever a channel row does not fill a whole 64-byte line,
	// and the packed size is the smaller one.
	state, err := allocBuffer(lib, lay.Size())
	if err != nil {
		return fmt.Errorf("allocate state buffer: %w", err)
	}
	defer state.close()
	reportLayout(lay)
	reportInitialContents(state, lay)

	update, err := openStateOp(lib, filepath.Join(root, "update"),
		mil.GenAccumulateState(stateName, shape), state, lay,
		[]e5rt.Port{{Name: "value", Size: bytes}},
		[]e5rt.Port{{Name: "y", Size: bytes}})
	if err != nil {
		return fmt.Errorf("open update program: %w", err)
	}
	defer update.close()

	reader, err := openStateOp(lib, filepath.Join(root, "reader"),
		mil.GenReadState(stateName, shape), state, lay,
		nil, []e5rt.Port{{Name: "y", Size: bytes}})
	if err != nil {
		return fmt.Errorf("open reader program: %w", err)
	}
	defer reader.close()

	fmt.Println()
	if err := hostVisibility(reader, state, lay); err != nil {
		return err
	}

	fmt.Println()
	want, err := recurrenceArm(update, state, lay)
	if err != nil {
		return err
	}

	fmt.Println()
	if err := readerArms(lib, root, shape, lay, reader, want); err != nil {
		return err
	}

	fmt.Println()
	if err := disjointArm(update, state, n); err != nil {
		return err
	}

	fmt.Println()
	if err := independenceArm(lib, root, shape, lay, update, state, n); err != nil {
		return err
	}

	fmt.Println()
	return refusalControls(update)
}

func checkFlags() error {
	if *channels < 1 || *dim < 2 {
		return errors.New("-chan must be at least 1 and -dim at least 2")
	}
	if *channels > 1<<20 || *dim > 1<<20 || *channels**dim > 1<<20 {
		return errors.New("-chan times -dim must not exceed 1048576")
	}
	if (*channels**dim)%2 != 0 {
		// The disjoint control splits the value vector in half.
		return errors.New("-chan times -dim must be even")
	}
	if *steps < 2 || *steps > 8 {
		// Outside this range the fp16 partial sums are no longer guaranteed
		// exact, and the exact comparison below would become a false failure.
		// Refusing is better than silently loosening the bar.
		return errors.New("-steps must be between 2 and 8 to keep the recurrence exact in fp16")
	}
	return nil
}

// layoutProbe reports which host byte offsets the engine reads each element of
// the state tensor from.
//
// It binds a state buffer far larger than the tensor and fills every 2-byte
// slot with its own index, so each value the reader returns names the host slot
// the engine took it from. A host-packed tensor would return 0, 1, 2, ..., n-1.
// Anything else is the layout a caller has to write through, and there is no
// way to learn it from the MIL text.
func layoutProbe(lib *e5rt.Lib, root string, shape [4]int) error {
	n := shape[1] * shape[3]
	// Give each channel four times the space its packed row needs, rounded up
	// to a 64-byte boundary, so a stride of one, two, or four cache lines is
	// visible rather than running off the end of the buffer.
	perChannel := 4 * max((shape[3]*2+63)&^63, 64)
	slots := shape[1] * perChannel / 2
	if float32(slots) > 2048 {
		// Slot indices are carried as fp16, which represents integers exactly
		// only up to 2048. Past that the probe would mislabel offsets.
		return fmt.Errorf("layout probe needs %d slots, past the %d that fp16 indexes exactly; use a smaller -chan/-dim", slots, 2048)
	}
	buf, err := allocBuffer(lib, slots*2)
	if err != nil {
		return err
	}
	defer buf.close()
	marks := make([]float32, slots)
	for i := range marks {
		marks[i] = float32(i)
	}
	if err := buf.writeSlots(marks); err != nil {
		return err
	}
	// The probe's buffer is deliberately four times any plausible layout, so
	// the size check in BindStatePort cannot refuse it whatever the real
	// stride turns out to be. Passing the computed layout here is a lower
	// bound for that check, not an input to the measurement.
	probeLay := e5rt.StateLayout{Channels: shape[1], Dim: shape[3]}
	reader, err := openStateOp(lib, filepath.Join(root, "layout"),
		mil.GenReadState(stateName, shape), buf, probeLay,
		nil, []e5rt.Port{{Name: "y", Size: n * 2}})
	if err != nil {
		return fmt.Errorf("open the layout reader: %w", err)
	}
	defer reader.close()

	got, err := reader.run("y", n)
	if err != nil {
		return err
	}
	fmt.Printf("  state shape [1, %d, 1, %d] over a %d-byte buffer\n", shape[1], shape[3], slots*2)
	fmt.Println("  each value below is the host 2-byte slot the engine read that element from:")
	for c := range shape[1] {
		fmt.Printf("    channel %d: %v\n", c, got[c*shape[3]:(c+1)*shape[3]])
	}
	packed := true
	for i, v := range got {
		if v != float32(i) {
			packed = false
			break
		}
	}
	if packed {
		fmt.Println("  the state tensor is host-packed: element i comes from slot i")
	} else {
		fmt.Println("  the state tensor is NOT host-packed; a caller must write through the layout above")
	}
	if *report {
		printLayoutReport(shape, slots, packed, got)
	}
	return nil
}

// layoutReport is what -report mails back. Slot maps are raw observations; the
// prediction fields restate what e5rt.StateLayout computes for this shape so a
// recipient can check agreement without importing anything.
type layoutReport struct {
	Date              string  `json:"date"`
	GOOS              string  `json:"goos"`
	GOArch            string  `json:"goarch"`
	OSVersion         string  `json:"os_version"`
	HWModel           string  `json:"hw_model"`
	Channels          int     `json:"channels"`
	Dim               int     `json:"dim"`
	RowBytes          int     `json:"predicted_row_bytes"`
	SizeBytes         int     `json:"predicted_size_bytes"`
	BufferSlots       int     `json:"buffer_slots"`
	Packed            bool    `json:"host_packed"`
	MatchesPrediction bool    `json:"matches_state_layout_prediction"`
	SlotMap           [][]int `json:"slot_map"`
}

func printLayoutReport(shape [4]int, slots int, packed bool, got []float32) {
	rep := layoutReport{
		Date:        time.Now().UTC().Format(time.RFC3339),
		GOOS:        runtime.GOOS,
		GOArch:      runtime.GOARCH,
		OSVersion:   sysctlString("kern.osproductversion"),
		HWModel:     sysctlString("hw.model"),
		Channels:    shape[1],
		Dim:         shape[3],
		BufferSlots: slots,
		Packed:      packed,
	}
	lay := e5rt.StateLayout{Channels: shape[1], Dim: shape[3]}
	rep.RowBytes = lay.RowBytes()
	rep.SizeBytes = lay.Size()
	mismatch := false
	for c := range shape[1] {
		row := make([]int, shape[3])
		for i := range row {
			v := int(got[c*shape[3]+i])
			row[i] = v
			if v != lay.Offset(c, i)/2 {
				mismatch = true
			}
		}
		rep.SlotMap = append(rep.SlotMap, row)
	}
	rep.MatchesPrediction = !mismatch
	blob, err := json.Marshal(rep)
	if err != nil {
		fmt.Printf("REPORT %v\n", err)
		return
	}
	fmt.Println("REPORT " + string(blob))
}

// sysctlString reads one string-valued sysctl, returning "" rather than failing:
// a mailed-back report should degrade gracefully, not refuse to exist because
// one identifier is unavailable.
func sysctlString(name string) string {
	v, err := syscall.Sysctl(name)
	if err != nil {
		return ""
	}
	return v
}

// hostVisibility establishes that a host write to the bound state buffer is
// what the engine subsequently reads. Every later arm resets the state by
// writing zeros through this pointer, so if the engine held a private copy
// those resets would be silent no-ops and the arms after them would be
// measuring something other than what they claim.
//
// The probe writes a distinctive pattern rather than zeros: zeros would also be
// consistent with the engine reading an untouched allocation.
func hostVisibility(reader *stateOp, state *buffer, lay e5rt.StateLayout) error {
	n := lay.Values()
	pattern := make([]float32, n)
	for i := range pattern {
		pattern[i] = float32(i%7) - 3
	}
	if err := state.writeTensor(lay, pattern); err != nil {
		return err
	}
	got, err := reader.run("y", n)
	if err != nil {
		return fmt.Errorf("read back the host-written pattern: %w", err)
	}
	diff, err := compare(got, pattern)
	if err != nil {
		return fmt.Errorf("host visibility probe: %w", err)
	}
	if diff != 0 {
		return fmt.Errorf("the engine did not read a host write to the bound state buffer (max diff %g); "+
			"host-side zeroing is not a reset and the remaining arms would be meaningless\n"+
			"    wrote %v\n    read  %v", diff, pattern, got)
	}
	fmt.Println("  a host write to the bound state buffer is what the engine reads back")
	return nil
}

// recurrenceArm applies the updates one at a time and checks the state after
// each against a float64 running sum. Checking every step rather than only the
// last one is what distinguishes a correct recurrence from one that happens to
// arrive at the right total.
func recurrenceArm(update *stateOp, state *buffer, lay e5rt.StateLayout) ([]float32, error) {
	n := lay.Values()
	if err := state.zero(); err != nil {
		return nil, err
	}
	sum := make([]float64, n)
	for t := range *steps {
		v := values(t, n)
		for i := range sum {
			sum[i] += float64(v[i])
		}
		got, err := update.step("value", v, "y", n)
		if err != nil {
			return nil, fmt.Errorf("step %d: %w", t, err)
		}
		diff, err := compare(got, toFloat32(sum))
		if err != nil {
			return nil, fmt.Errorf("step %d: %w", t, err)
		}
		if diff != 0 {
			return nil, fmt.Errorf("step %d differs from the float64 recurrence by %g; "+
				"the values were chosen so this recurrence is exact in fp16", t, diff)
		}
	}
	want := toFloat32(sum)

	// A later arm requires a reader on a different buffer to disagree with
	// this result. If the accumulated state were all zeros, that arm would
	// pass on a completely broken implementation.
	if allZero(want) {
		return nil, errors.New("the accumulated state is all zeros, which would make the fresh-buffer control vacuous")
	}

	fmt.Printf("  %d updates match a float64 recurrence exactly, checked after every step\n", *steps)

	// The state buffer's own host pointer should now show the accumulated
	// values. This is the reverse direction of the visibility probe.
	direct, err := state.readTensor(lay)
	if err != nil {
		return nil, err
	}
	diff, err := compare(direct, want)
	if err != nil {
		return nil, fmt.Errorf("reading the state buffer directly: %w", err)
	}
	if diff != 0 {
		return nil, fmt.Errorf("the state buffer's host pointer shows values %g away from what the engine computed", diff)
	}
	fmt.Println("  the engine's writes are visible through the state buffer's own host pointer")
	return want, nil
}

// readerArms run the separately compiled reader against the shared state, then
// against a buffer it has never touched. The pair is the point: the first alone
// would be consistent with the reader obtaining the right answer by some route
// that has nothing to do with the buffer being shared.
func readerArms(lib *e5rt.Lib, root string, shape [4]int, lay e5rt.StateLayout, reader *stateOp, want []float32) error {
	got, err := reader.run("y", len(want))
	if err != nil {
		return fmt.Errorf("read the shared state: %w", err)
	}
	diff, err := compare(got, want)
	if err != nil {
		return fmt.Errorf("reader on the shared buffer: %w", err)
	}
	if diff != 0 {
		return fmt.Errorf("a separately compiled reader bound to the same buffer disagrees by %g", diff)
	}
	fmt.Println("  a separately compiled reader on the same buffer returns the accumulated state")

	// The control: same program text, same compiled shape, different buffer.
	other, err := allocBuffer(lib, lay.Size())
	if err != nil {
		return err
	}
	defer other.close()
	if err := other.zero(); err != nil {
		return err
	}
	control, err := openStateOp(lib, filepath.Join(root, "reader-control"),
		mil.GenReadState(stateName, shape), other, lay,
		nil, []e5rt.Port{{Name: "y", Size: len(want) * 2}})
	if err != nil {
		return fmt.Errorf("open the fresh-buffer reader: %w", err)
	}
	defer control.close()

	fresh, err := control.run("y", len(want))
	if err != nil {
		return fmt.Errorf("read the fresh buffer: %w", err)
	}
	if _, err := compare(fresh, want); err != nil {
		return fmt.Errorf("reader on a fresh buffer: %w", err)
	}
	if !allZero(fresh) {
		return fmt.Errorf("a reader bound to a buffer this program zeroed and never updated returned %v", fresh)
	}
	fmt.Println("  the same reader program bound to a different buffer returns zeros, as it must")
	return nil
}

// disjointArm updates two halves of the state in separate steps. Accumulating
// into a region must leave the other region alone, which is the property a
// position-indexed cache would be built on. It is checked here on packed
// regions, which is weaker than checking it on cache positions.
func disjointArm(update *stateOp, state *buffer, n int) error {
	if err := state.zero(); err != nil {
		return err
	}
	half := n / 2
	lower := make([]float32, n)
	upper := make([]float32, n)
	for i := range half {
		lower[i] = float32(i%5) + 1
		upper[half+i] = -(float32(i%3) + 1)
	}

	if _, err := update.step("value", lower, "y", n); err != nil {
		return fmt.Errorf("lower-half update: %w", err)
	}
	got, err := update.step("value", upper, "y", n)
	if err != nil {
		return fmt.Errorf("upper-half update: %w", err)
	}
	want := make([]float32, n)
	for i := range want {
		want[i] = lower[i] + upper[i]
	}
	diff, err := compare(got, want)
	if err != nil {
		return fmt.Errorf("disjoint control: %w", err)
	}
	if diff != 0 {
		return fmt.Errorf("updating the upper half disturbed the lower half, or landed at the wrong offset (max diff %g)", diff)
	}
	fmt.Printf("  updating %d values in the upper half left the lower half untouched\n", half)
	return nil
}

// independenceArm advances two states in an interleaved order. Every check
// above uses one state, and a single process-wide buffer behind the inout port
// would satisfy all of them. Two states with different values, updated in
// alternation and each required to follow its own recurrence, would not.
func independenceArm(lib *e5rt.Lib, root string, shape [4]int, lay e5rt.StateLayout, a *stateOp, stateA *buffer, n int) error {
	stateB, err := allocBuffer(lib, lay.Size())
	if err != nil {
		return err
	}
	defer stateB.close()
	if err := stateB.zero(); err != nil {
		return err
	}
	b, err := openStateOp(lib, filepath.Join(root, "update-b"),
		mil.GenAccumulateState(stateName, shape), stateB, lay,
		[]e5rt.Port{{Name: "value", Size: n * 2}},
		[]e5rt.Port{{Name: "y", Size: n * 2}})
	if err != nil {
		return fmt.Errorf("open the second update program: %w", err)
	}
	defer b.close()

	if err := stateA.zero(); err != nil {
		return err
	}
	sumA := make([]float64, n)
	sumB := make([]float64, n)
	for t := range 3 {
		for _, arm := range []struct {
			name string
			op   *stateOp
			sum  []float64
			v    []float32
		}{
			{"A", a, sumA, values(t, n)},
			{"B", b, sumB, negate(values(t+4, n))},
		} {
			for i := range arm.sum {
				arm.sum[i] += float64(arm.v[i])
			}
			got, err := arm.op.step("value", arm.v, "y", n)
			if err != nil {
				return fmt.Errorf("state %s step %d: %w", arm.name, t, err)
			}
			diff, err := compare(got, toFloat32(arm.sum))
			if err != nil {
				return fmt.Errorf("state %s step %d: %w", arm.name, t, err)
			}
			if diff != 0 {
				return fmt.Errorf("state %s step %d differs by %g; the two states are not independent", arm.name, t, diff)
			}
		}
	}
	if sameValues(toFloat32(sumA), toFloat32(sumB)) {
		return errors.New("the two states ended up equal, which would make the independence control vacuous")
	}
	fmt.Println("  two states advanced in alternation each followed their own recurrence")
	return nil
}

// refusalControls check that retaining an inout port resolves a name rather
// than returning whatever it is handed.
func refusalControls(op *stateOp) error {
	_, err := op.lib.OperationRetainInoutPort(op.op, "no_such_state")
	if err == nil {
		return errors.New(`retaining an inout port named "no_such_state" succeeded; the name is not being resolved`)
	}
	fmt.Printf("  refused, as it must be: an inout port the program does not declare\n    %v\n", err)
	// The positive polarity, on the same operation: the declared name still
	// resolves, so the refusal above is about the name and not about the call
	// having stopped working.
	port, err := op.lib.OperationRetainInoutPort(op.op, stateName)
	if err != nil {
		return fmt.Errorf("retaining the declared inout port %q now fails: %w", stateName, err)
	}
	defer op.lib.IOPortRelease(port)
	fmt.Printf("  and the declared port %q still resolves on the same operation\n", stateName)
	return nil
}

// A buffer is one E5RT buffer object together with its host-visible storage.
//
// The state buffers are allocated here rather than by an operation, because
// several operations bind to the same one and it has to outlive all of them.
type buffer struct {
	lib  *e5rt.Lib
	obj  uintptr
	data []byte
}

func allocBuffer(lib *e5rt.Lib, n int) (*buffer, error) {
	if n <= 0 {
		return nil, fmt.Errorf("buffer size %d", n)
	}
	// E5RT buffers are allocated 64-byte aligned and rounded up elsewhere in
	// this module; matching that here keeps the demo from being the place a
	// stricter allocator requirement first shows up.
	obj, err := lib.BufferObjectAlloc(uintptr(max((n+63)&^63, 64)), 0)
	if err != nil {
		return nil, err
	}
	ptr, err := lib.BufferObjectGetDataPtr(obj)
	if err != nil {
		lib.BufferObjectRelease(obj)
		return nil, err
	}
	return &buffer{lib: lib, obj: obj, data: unsafe.Slice((*byte)(pointerAt(ptr)), n)}, nil
}

// pointerAt converts an address E5RT returned into a pointer. The address does
// not originate in Go, so the conversion is laundered through memory rather
// than written as unsafe.Pointer(uintptr), which go vet flags and which the
// garbage collector is entitled to treat as a non-pointer.
func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}

// writeTensor stores v, one channel row at a time, at the layout's offsets.
func (b *buffer) writeTensor(l e5rt.StateLayout, v []float32) error {
	if len(v) != l.Values() {
		return fmt.Errorf("writing %d values to a [1, %d, 1, %d] tensor", len(v), l.Channels, l.Dim)
	}
	if len(b.data) < l.Size() {
		return fmt.Errorf("writing a %d-byte tensor to a %d-byte buffer", l.Size(), len(b.data))
	}
	clear(b.data)
	for c := range l.Channels {
		row := b.data[c*l.RowBytes():]
		for i := range l.Dim {
			bits := ane.Float32ToFP16(v[c*l.Dim+i])
			row[2*i] = byte(bits)
			row[2*i+1] = byte(bits >> 8)
		}
	}
	return nil
}

// readTensor loads the layout's offsets into a packed slice.
func (b *buffer) readTensor(l e5rt.StateLayout) ([]float32, error) {
	if len(b.data) < l.Size() {
		return nil, fmt.Errorf("reading a %d-byte tensor from a %d-byte buffer", l.Size(), len(b.data))
	}
	out := make([]float32, l.Values())
	for c := range l.Channels {
		row := b.data[c*l.RowBytes():]
		for i := range l.Dim {
			out[c*l.Dim+i] = ane.FP16ToFloat32(uint16(row[2*i]) | uint16(row[2*i+1])<<8)
		}
	}
	return out, nil
}

// writeSlots and readSlots treat the buffer as consecutive fp16 values with no
// padding. The ordinary input and output ports use this: only the state buffer
// carries the padded layout above. That asymmetry is measured, not assumed —
// see the package comment.
func (b *buffer) writeSlots(v []float32) error {
	if len(v)*2 != len(b.data) {
		return fmt.Errorf("writing %d values to a %d-byte buffer", len(v), len(b.data))
	}
	for i, x := range v {
		bits := ane.Float32ToFP16(x)
		b.data[2*i] = byte(bits)
		b.data[2*i+1] = byte(bits >> 8)
	}
	return nil
}

func (b *buffer) readSlots(n int) ([]float32, error) {
	if n*2 != len(b.data) {
		return nil, fmt.Errorf("reading %d values from a %d-byte buffer", n, len(b.data))
	}
	out := make([]float32, n)
	for i := range out {
		out[i] = ane.FP16ToFloat32(uint16(b.data[2*i]) | uint16(b.data[2*i+1])<<8)
	}
	return out, nil
}

func (b *buffer) zero() error {
	clear(b.data)
	return nil
}

func (b *buffer) close() error {
	if b == nil || b.obj == 0 {
		return nil
	}
	err := b.lib.BufferObjectRelease(b.obj)
	b.obj, b.data = 0, nil
	return err
}

// A stateOp is a compiled MIL program whose state parameter is bound to a
// buffer the caller owns. Its ordinary input and output ports get buffers of
// their own; only the state is shared.
type stateOp struct {
	lib      *e5rt.Lib
	op       uintptr
	stream   uintptr
	in, out  map[string]*buffer
	releases []func() error
}

func openStateOp(lib *e5rt.Lib, dir, text string, state *buffer, lay e5rt.StateLayout, inputs, outputs []e5rt.Port) (_ *stateOp, err error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	modelPath := filepath.Join(dir, "model.mil")
	if err := os.WriteFile(modelPath, []byte(text), 0o644); err != nil {
		return nil, err
	}

	o := &stateOp{lib: lib, in: map[string]*buffer{}, out: map[string]*buffer{}}
	defer func() {
		if err != nil {
			err = errors.Join(err, o.close())
		}
	}()
	keep := func(f func() error) { o.releases = append(o.releases, f) }

	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		return nil, fmt.Errorf("create compiler config: %w", err)
	}
	keep(func() error { return lib.CompilerConfigOptionsRelease(config) })
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		return nil, fmt.Errorf("set cache location: %w", err)
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create compiler: %w", err)
	}
	keep(func() error { return lib.CompilerRelease(compiler) })
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		return nil, fmt.Errorf("create compiler options: %w", err)
	}
	keep(func() error { return lib.CompilerOptionsRelease(options) })
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		return nil, fmt.Errorf("set device mask: %w", err)
	}
	library, err := lib.CompilerCompile(compiler, modelPath, options)
	if err != nil {
		return nil, fmt.Errorf("compile: %w", err)
	}
	keep(func() error { return lib.ProgramLibraryRelease(library) })
	if err := requireANE(dir); err != nil {
		return nil, err
	}

	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		return nil, fmt.Errorf("retain function: %w", err)
	}
	keep(func() error { return lib.ProgramFunctionRelease(function) })
	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		return nil, fmt.Errorf("create operation options: %w", err)
	}
	keep(func() error { return lib.PrecompiledComputeOpOptionsRelease(opOptions) })
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		return nil, fmt.Errorf("set operation name: %w", err)
	}
	if err := lib.PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(opOptions, true); err != nil {
		return nil, fmt.Errorf("allocate intermediate buffers: %w", err)
	}
	if o.op, err = lib.OperationCreatePrecompiled(opOptions); err != nil {
		return nil, fmt.Errorf("create operation: %w", err)
	}
	keep(func() error { return lib.OperationRelease(o.op) })

	// The state port takes the caller's buffer. Everything else gets its own.
	//
	// BindStatePort is the checked bind: it asks the runtime how large the
	// buffer is and refuses one too small for the layout. The raw pair of
	// calls accepts an undersized buffer without complaint, which is how the
	// padding went unnoticed until this example ran.
	statePort, err := lib.BindStatePort(o.op, stateName, state.obj, lay)
	if err != nil {
		return nil, err
	}
	keep(func() error { return lib.IOPortRelease(statePort) })

	bind := func(spec e5rt.Port, input bool) error {
		var port uintptr
		var err error
		if input {
			port, err = lib.OperationRetainInputPort(o.op, spec.Name)
		} else {
			port, err = lib.OperationRetainOutputPort(o.op, spec.Name)
		}
		if err != nil {
			return fmt.Errorf("retain port %q: %w", spec.Name, err)
		}
		keep(func() error { return lib.IOPortRelease(port) })
		buf, err := allocBuffer(lib, spec.Size)
		if err != nil {
			return fmt.Errorf("allocate buffer for port %q: %w", spec.Name, err)
		}
		keep(buf.close)
		if err := lib.IOPortBindBufferObject(port, buf.obj); err != nil {
			return fmt.Errorf("bind buffer for port %q: %w", spec.Name, err)
		}
		if input {
			o.in[spec.Name] = buf
		} else {
			o.out[spec.Name] = buf
		}
		return nil
	}
	for _, spec := range inputs {
		if err := bind(spec, true); err != nil {
			return nil, err
		}
	}
	for _, spec := range outputs {
		if err := bind(spec, false); err != nil {
			return nil, err
		}
	}

	if o.stream, err = lib.ExecutionStreamCreate(); err != nil {
		return nil, fmt.Errorf("create execution stream: %w", err)
	}
	keep(func() error { return lib.ExecutionStreamRelease(o.stream) })
	if err := lib.EncodeOperation(o.stream, o.op); err != nil {
		return nil, fmt.Errorf("encode operation: %w", err)
	}
	return o, nil
}

// step writes one input, executes, and returns one output.
func (o *stateOp) step(inName string, v []float32, outName string, n int) ([]float32, error) {
	in, ok := o.in[inName]
	if !ok {
		return nil, fmt.Errorf("no input port %q", inName)
	}
	if err := in.writeSlots(v); err != nil {
		return nil, err
	}
	return o.run(outName, n)
}

// run executes and returns one output.
func (o *stateOp) run(outName string, n int) ([]float32, error) {
	if err := o.lib.ExecuteSync(o.stream); err != nil {
		return nil, fmt.Errorf("execute: %w", err)
	}
	out, ok := o.out[outName]
	if !ok {
		return nil, fmt.Errorf("no output port %q", outName)
	}
	return out.readSlots(n)
}

func (o *stateOp) close() error {
	var errs []error
	for i := len(o.releases) - 1; i >= 0; i-- {
		if err := o.releases[i](); err != nil {
			errs = append(errs, err)
		}
	}
	o.releases = nil
	return errors.Join(errs...)
}

// requireANE fails unless the compiler left an ANE artifact. The device mask is
// a permission and not a placement, so this directory is the only local
// evidence that the state actually lives on the engine rather than on the CPU,
// where none of this would be interesting.
func requireANE(cacheDir string) error {
	seen := map[string]bool{}
	filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() || filepath.Base(filepath.Dir(path)) != "main" {
			return nil
		}
		if suffix, ok := strings.CutPrefix(d.Name(), "main_"); ok {
			seen[suffix] = true
		}
		return nil
	})
	if !seen["ane"] {
		names := make([]string, 0, len(seen))
		for n := range seen {
			names = append(names, n)
		}
		return fmt.Errorf("the compiler emitted %v, not an ANE artifact", names)
	}
	return nil
}

// reportLayout states the state buffer's shape in bytes, because it is the one
// thing a caller has to get right that the MIL text does not tell them.
func reportLayout(l e5rt.StateLayout) {
	if l.Packed() {
		fmt.Printf("  state buffer: %d bytes, %d channels of %d bytes — packed at this shape\n",
			l.Size(), l.Channels, l.RowBytes())
		return
	}
	fmt.Printf("  state buffer: %d bytes, %d channels of %d bytes — a packed %d would be too small by %d\n",
		l.Size(), l.Channels, l.RowBytes(), l.Values()*2, l.Size()-l.Values()*2)
}

// reportInitialContents says what a fresh buffer object held. Nothing here
// depends on it — every arm zeroes the state explicitly — but an allocator that
// returns dirty memory is worth knowing about, and reporting it is free.
func reportInitialContents(b *buffer, lay e5rt.StateLayout) {
	v, err := b.readTensor(lay)
	if err != nil {
		fmt.Printf("  freshly allocated buffer: UNREAD (%v)\n", err)
		return
	}
	if allZero(v) {
		fmt.Println("  a freshly allocated buffer object read back as all zeros (observed, not relied on)")
		return
	}
	fmt.Printf("  a freshly allocated buffer object was NOT zeroed: %v\n", v)
}

// values returns step t's update. Every element is a multiple of 0.5 with
// magnitude at most 2, so eight steps cannot take any partial sum past 16 and
// the whole recurrence stays exact in fp16. See the package comment.
//
// The residues are deliberately asymmetric about zero. An earlier version used
// (7t+3i) mod 9, whose values over six steps sum to exactly zero for every i —
// which would have left the accumulated state all zeros and made the
// fresh-buffer control below pass against a completely broken implementation.
// The guard in recurrenceArm caught it; this generator is the repair.
func values(t, n int) []float32 {
	v := make([]float32, n)
	for i := range v {
		v[i] = float32((t*3+i*5)%7-2) * 0.5
	}
	return v
}

func negate(v []float32) []float32 {
	out := make([]float32, len(v))
	for i, x := range v {
		out[i] = -x
	}
	return out
}

func toFloat32(v []float64) []float32 {
	out := make([]float32, len(v))
	for i, x := range v {
		out[i] = float32(x)
	}
	return out
}

func allZero(v []float32) bool {
	for _, x := range v {
		if x != 0 {
			return false
		}
	}
	return true
}

func sameValues(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// compare returns the largest absolute difference between got and want.
//
// It refuses unequal lengths and non-finite values. Both refusals are load
// bearing: comparing a prefix reports a truncated result as identical, and
// NaN > max is false, so an all-NaN result reported a maximum difference of
// zero in an earlier version of a neighbouring example. Infinity fails loudly
// on its own; NaN is the silent one, and NaN is what an fp16 overflow in an
// accumulating state would produce.
func compare(got, want []float32) (float64, error) {
	if len(got) != len(want) {
		return 0, fmt.Errorf("comparing %d values against %d", len(got), len(want))
	}
	if len(got) == 0 {
		return 0, errors.New("nothing to compare")
	}
	var maxDiff float64
	for i := range got {
		g, w := float64(got[i]), float64(want[i])
		if math.IsNaN(g) || math.IsInf(g, 0) {
			return 0, fmt.Errorf("result element %d is %v", i, g)
		}
		if math.IsNaN(w) || math.IsInf(w, 0) {
			return 0, fmt.Errorf("reference element %d is %v", i, w)
		}
		if d := math.Abs(g - w); d > maxDiff {
			maxDiff = d
		}
	}
	return maxDiff, nil
}
