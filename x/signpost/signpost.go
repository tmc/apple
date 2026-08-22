package signpost

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"

	"github.com/tmc/apple/x/internal/oslogabi"
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

// messageFormat is the format string the Message variants emit; a pooled
// in-image copy (see oslogstrings_cgo.go and pool.go) lets it decode.
const messageFormat = "%{public}s"

// These are the os_signpost primitives. The interval and event operations are
// C macros with no exported symbol; they lower to
// _os_signpost_emit_with_name_impl, which this package calls directly.
var (
	initOnce sync.Once

	osLogCreate       func(subsystem, category *byte) uintptr
	osRelease         func(obj uintptr)
	osSignpostIDGen   func(log uintptr) uint64
	osSignpostEnabled func(log uintptr) bool
	osSignpostEmit    func(dso unsafe.Pointer, log uintptr, typ uint8, spid uint64, name, format *byte, buf *byte, size uint32)
)

func loadSymbols() {
	initOnce.Do(func() {
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_log_create"); err == nil && sym != 0 {
			purego.RegisterFunc(&osLogCreate, sym)
		}
		if sym, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_release"); err == nil && sym != 0 {
			purego.RegisterFunc(&osRelease, sym)
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
		if l.handle != 0 && osRelease != nil {
			runtime.AddCleanup(l, func(h uintptr) { osRelease(h) }, l.handle)
		}
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

// IntervalBeginMessage marks the start of a named interval identified by id,
// carrying msg as a public formatted message. Instruments groups intervals by
// name and shows the message as detail, so a small fixed set of names with a
// descriptive message yields one track per name rather than one per span.
func (l *Logger) IntervalBeginMessage(id ID, name, msg string) {
	l.emitMessage(typeIntervalBegin, id, name, msg)
}

// IntervalEndMessage marks the end of the interval begun with the same id and
// name, carrying msg as a public formatted message.
func (l *Logger) IntervalEndMessage(id ID, name, msg string) {
	l.emitMessage(typeIntervalEnd, id, name, msg)
}

// EventMessage emits a single point-in-time signpost with the given name,
// carrying msg as a public formatted message.
func (l *Logger) EventMessage(id ID, name, msg string) {
	l.emitMessage(typeEvent, id, name, msg)
}

// IntervalBeginf begins a named interval whose message is built from an
// os_log format string and typed arguments (%d/%u/%x and l-prefixed 64-bit
// forms, %p, %f family, %s/%@; %{public}/%{private} set visibility). The
// format string must be a literal pooled by signpostnames for the message to
// decode; the arguments are serialized into the tracepoint at emit time.
func (l *Logger) IntervalBeginf(id ID, name, format string, args ...any) {
	l.emitf(typeIntervalBegin, id, name, format, args)
}

// IntervalEndf ends the interval begun with the same id and name, with a
// formatted message as in [Logger.IntervalBeginf].
func (l *Logger) IntervalEndf(id ID, name, format string, args ...any) {
	l.emitf(typeIntervalEnd, id, name, format, args)
}

// Eventf emits a point-in-time signpost with a formatted message as in
// [Logger.IntervalBeginf].
func (l *Logger) Eventf(id ID, name, format string, args ...any) {
	l.emitf(typeEvent, id, name, format, args)
}

func (l *Logger) emitMessage(typ Type, id ID, name, msg string) {
	if msg == "" {
		l.emit(typ, id, name)
		return
	}
	l.emitf(typ, id, name, messageFormat, []any{msg})
}

func (l *Logger) emit(typ Type, id ID, name string) {
	l.emitf(typ, id, name, "", nil)
}

// emitf is the single emission path. The argument buffer the C macros build
// at compile time via __builtin_os_log_format comes from oslogabi.Encode;
// the name and format strings decode by offset into the dso image and must
// be pooled (see loadnames.go), while the encoded arguments are serialized
// into the tracepoint at emit time and may live anywhere.
func (l *Logger) emitf(typ Type, id ID, name, format string, args []any) {
	if l == nil || l.handle == 0 || osSignpostEmit == nil || osSignpostEnabled == nil {
		return
	}
	// The macros skip emission for the reserved ids; match that.
	if id == idNull || id == idInvalid {
		return
	}
	if !osSignpostEnabled(l.handle) {
		return
	}
	nameStr, formatStr, dso, nameOK := lookup(name, format)
	var cfmt, cname []byte
	if !nameOK {
		warnUnpooled(name)
	}
	if nameStr == nil {
		cname = cString(name)
		nameStr = &cname[0]
	}
	if formatStr == nil {
		cfmt = cString(format)
		formatStr = &cfmt[0]
	}
	if dso == nil {
		dso = dsoHandle()
	}
	buf, pins := oslogabi.Encode(format, args)
	osSignpostEmit(
		dso,
		l.handle,
		uint8(typ),
		uint64(id),
		nameStr,
		formatStr,
		&buf[0],
		uint32(len(buf)),
	)
	// Keep everything passed by pointer alive until the call returns.
	runtime.KeepAlive(cname)
	runtime.KeepAlive(cfmt)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(pins)
}

// cString returns a NUL-terminated copy of s as a byte slice suitable for
// passing its first element as a C string. Callers must keep the slice alive
// (via runtime.KeepAlive) for the duration of any call that reads it.
func cString(s string) []byte {
	b := make([]byte, len(s)+1)
	copy(b, s)
	return b
}
