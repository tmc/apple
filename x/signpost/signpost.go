package signpost

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Type identifies a signpost operation. It mirrors os_signpost_type_t.
type Type uint8

const (
	typeEvent         Type = 0x00 // OS_SIGNPOST_EVENT
	typeIntervalBegin Type = 0x01 // OS_SIGNPOST_INTERVAL_BEGIN
	typeIntervalEnd   Type = 0x02 // OS_SIGNPOST_INTERVAL_END
)

// ID identifies a signpost so that a begin can be paired with its end. It
// mirrors os_signpost_id_t. The zero value is not usable; obtain one from
// [Logger.NewID].
type ID uint64

// Reserved signpost id values (see <os/signpost.h>).
const (
	idNull      ID = 0                  // OS_SIGNPOST_ID_NULL
	idInvalid   ID = ^ID(0)             // OS_SIGNPOST_ID_INVALID
	idExclusive ID = 0xEEEEB0B5B2B2EEEE // OS_SIGNPOST_ID_EXCLUSIVE
)

// IDExclusive is a shared id usable when at most one interval with a given
// name is in flight at a time on a log, avoiding the need to thread an [ID]
// through the code between begin and end.
const IDExclusive = idExclusive

// Category names understood by Instruments and the logging system.
const (
	// PointsOfInterest is the category Instruments displays in the Points of
	// Interest track. It maps to OS_LOG_CATEGORY_POINTS_OF_INTEREST.
	PointsOfInterest = "PointsOfInterest"
	// DynamicTracing is a category whose signposts are disabled until a tool
	// such as Instruments enables them. It maps to
	// OS_LOG_CATEGORY_DYNAMIC_TRACING.
	DynamicTracing = "DynamicTracing"
)

// These are the os_signpost primitives. The interval and event operations are
// C macros with no exported symbol; they lower to
// _os_signpost_emit_with_name_impl, which this package calls directly.
var (
	initOnce sync.Once

	osLogCreate       func(subsystem, category *byte) uintptr
	osSignpostIDGen   func(log uintptr) uint64
	osSignpostEnabled func(log uintptr) bool
	osSignpostEmit    func(dso unsafe.Pointer, log uintptr, typ uint8, spid uint64, name, format *byte, buf *byte, size uint32)
)

func loadSymbols() {
	initOnce.Do(func() {
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_log_create"); err == nil && sym != 0 {
			purego.RegisterFunc(&osLogCreate, sym)
		}
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_signpost_id_generate"); err == nil && sym != 0 {
			purego.RegisterFunc(&osSignpostIDGen, sym)
		}
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_signpost_enabled"); err == nil && sym != 0 {
			purego.RegisterFunc(&osSignpostEnabled, sym)
		}
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "_os_signpost_emit_with_name_impl"); err == nil && sym != 0 {
			purego.RegisterFunc(&osSignpostEmit, sym)
		}
	})
}

// Logger emits signposts against a single os_log handle. It is created with
// [New] and is safe for concurrent use. The zero value is not usable.
type Logger struct {
	handle uintptr
}

// New returns a Logger that emits signposts under the given subsystem and
// category. Use [PointsOfInterest] as the category to have intervals appear in
// the Instruments Points of Interest track. New never returns nil; if the
// signpost symbols cannot be resolved the returned Logger's methods are
// no-ops and [Logger.Enabled] reports false.
func New(subsystem, category string) *Logger {
	loadSymbols()
	l := &Logger{}
	if osLogCreate != nil {
		sub := cString(subsystem)
		cat := cString(category)
		l.handle = osLogCreate(&sub[0], &cat[0])
		runtime.KeepAlive(sub)
		runtime.KeepAlive(cat)
	}
	return l
}

// Enabled reports whether signposts are being recorded for this log. Emitting
// while disabled is harmless but wasteful, so hot paths may check first.
func (l *Logger) Enabled() bool {
	if l == nil || l.handle == 0 || osSignpostEnabled == nil {
		return false
	}
	return osSignpostEnabled(l.handle)
}

// NewID returns an ID that is unique among signposts logged to this Logger.
// Pair the returned ID's begin and end to mark an interval.
func (l *Logger) NewID() ID {
	if l == nil || l.handle == 0 || osSignpostIDGen == nil {
		return idNull
	}
	return ID(osSignpostIDGen(l.handle))
}

// IntervalBegin marks the start of a named interval identified by id. Pair it
// with an [Logger.IntervalEnd] call using the same id and name.
func (l *Logger) IntervalBegin(id ID, name string) {
	l.emit(typeIntervalBegin, id, name)
}

// IntervalEnd marks the end of the interval begun with the same id and name.
func (l *Logger) IntervalEnd(id ID, name string) {
	l.emit(typeIntervalEnd, id, name)
}

// Event emits a single point-in-time signpost with the given name.
func (l *Logger) Event(id ID, name string) {
	l.emit(typeEvent, id, name)
}

// emitBuf is a minimal, empty os_log format buffer: a zero summary byte and a
// zero argument count. The underlying symbol requires a non-nil buffer even
// when no formatted arguments are present.
var emitBuf = [2]byte{0, 0}

func (l *Logger) emit(typ Type, id ID, name string) {
	if l == nil || l.handle == 0 || osSignpostEmit == nil {
		return
	}
	// The macros skip emission for the reserved ids; match that.
	if id == idNull || id == idInvalid {
		return
	}
	if !osSignpostEnabled(l.handle) {
		return
	}
	format := [1]byte{0}
	cname := cString(name)
	osSignpostEmit(
		dsoHandle(),
		l.handle,
		uint8(typ),
		uint64(id),
		&cname[0],
		&format[0],
		&emitBuf[0],
		uint32(len(emitBuf)),
	)
	// Keep the buffers passed by pointer alive until the call returns.
	runtime.KeepAlive(cname)
	runtime.KeepAlive(format)
}

// cString returns a NUL-terminated copy of s as a byte slice suitable for
// passing its first element as a C string. Callers must keep the slice alive
// (via runtime.KeepAlive) for the duration of any call that reads it.
func cString(s string) []byte {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return b
}
