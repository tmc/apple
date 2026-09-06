//go:build cgo && darwin

package signpostmsg

/*
#include <os/log.h>
#include <os/signpost.h>
#include <stdlib.h>

// The os_signpost operations are macros that require both the signpost name
// and the format string to be compile-time literals, so each name needs its
// own call site. The switch below supplies them; the message is the single
// run-time argument.

static os_log_t spmsg_log_create(const char *subsystem, const char *category) {
	return os_log_create(subsystem, category);
}

static void spmsg_log_release(os_log_t log) {
	os_release(log);
}

static os_signpost_id_t spmsg_id_generate(os_log_t log) {
	return os_signpost_id_generate(log);
}

static bool spmsg_enabled(os_log_t log) {
	return os_signpost_enabled(log);
}

// spmsg_emit dispatches on op (0 begin, 1 end, 2 event) and name (0 Model,
// 1 Layer, 2 Op). The nine call sites exist because the literals cannot be
// hoisted into variables.
static void spmsg_emit(os_log_t log, os_signpost_id_t spid, int op, int name, const char *msg) {
	switch (op) {
	case 0:
		switch (name) {
		case 0: os_signpost_interval_begin(log, spid, "Model", "%{public}s", msg); break;
		case 1: os_signpost_interval_begin(log, spid, "Layer", "%{public}s", msg); break;
		case 2: os_signpost_interval_begin(log, spid, "Op",    "%{public}s", msg); break;
		}
		break;
	case 1:
		switch (name) {
		case 0: os_signpost_interval_end(log, spid, "Model", "%{public}s", msg); break;
		case 1: os_signpost_interval_end(log, spid, "Layer", "%{public}s", msg); break;
		case 2: os_signpost_interval_end(log, spid, "Op",    "%{public}s", msg); break;
		}
		break;
	case 2:
		switch (name) {
		case 0: os_signpost_event_emit(log, spid, "Model", "%{public}s", msg); break;
		case 1: os_signpost_event_emit(log, spid, "Layer", "%{public}s", msg); break;
		case 2: os_signpost_event_emit(log, spid, "Op",    "%{public}s", msg); break;
		}
		break;
	}
}
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// ID identifies a signpost so that a begin can be paired with its end. It
// mirrors os_signpost_id_t. The zero value is not usable; obtain one from
// [Logger.NewID].
type ID uint64

// Reserved signpost id values (see <os/signpost.h>).
const (
	idNull    ID = 0
	idInvalid ID = ^ID(0)
)

// IDExclusive is a shared id usable when at most one interval with a given
// name is in flight at a time on a log, avoiding the need to thread an [ID]
// through the code between begin and end.
const IDExclusive ID = 0xEEEEB0B5B2B2EEEE

// Logger emits signposts against a single os_log handle. It is created with
// [New] and is safe for concurrent use. The zero value is not usable.
type Logger struct {
	handle C.os_log_t
}

// New returns a Logger that emits signposts under the given subsystem and
// category. Use [PointsOfInterest] as the category to have intervals appear in
// the Instruments Points of Interest track. New never returns nil.
func New(subsystem, category string) *Logger {
	sub := C.CString(subsystem)
	defer C.free(unsafe.Pointer(sub))
	cat := C.CString(category)
	defer C.free(unsafe.Pointer(cat))

	l := &Logger{handle: C.spmsg_log_create(sub, cat)}
	runtime.AddCleanup(l, func(h C.os_log_t) { C.spmsg_log_release(h) }, l.handle)
	return l
}

// Enabled reports whether signposts are being recorded for this log. Emitting
// while disabled is harmless but wasteful, so hot paths may check first.
func (l *Logger) Enabled() bool {
	if l == nil || l.handle == nil {
		return false
	}
	return bool(C.spmsg_enabled(l.handle))
}

// NewID returns an ID that is unique among signposts logged to this Logger.
// Pair the returned ID's begin and end to mark an interval.
func (l *Logger) NewID() ID {
	if l == nil || l.handle == nil {
		return idNull
	}
	return ID(C.spmsg_id_generate(l.handle))
}

// IntervalBegin marks the start of an interval identified by id, tagged with
// name and carrying msg as its message. Pair it with an [Logger.IntervalEnd]
// call using the same id and name.
func (l *Logger) IntervalBegin(id ID, name Name, msg string) {
	l.emit(opBegin, id, name, msg)
}

// IntervalEnd marks the end of the interval begun with the same id and name.
func (l *Logger) IntervalEnd(id ID, name Name, msg string) {
	l.emit(opEnd, id, name, msg)
}

// Event emits a single point-in-time signpost tagged with name and carrying
// msg as its message.
func (l *Logger) Event(id ID, name Name, msg string) {
	l.emit(opEvent, id, name, msg)
}

// Signpost operations, matching the op dispatch in spmsg_emit.
const (
	opBegin C.int = 0
	opEnd   C.int = 1
	opEvent C.int = 2
)

func (l *Logger) emit(op C.int, id ID, name Name, msg string) {
	if l == nil || l.handle == nil {
		return
	}
	// The macros skip emission for the reserved ids; match that.
	if id == idNull || id == idInvalid {
		return
	}
	if name < Model || name > Op {
		return
	}
	if !bool(C.spmsg_enabled(l.handle)) {
		return
	}
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))
	C.spmsg_emit(l.handle, C.os_signpost_id_t(id), op, C.int(name), cmsg)
}
