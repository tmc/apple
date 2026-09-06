//go:build !cgo || !darwin

package signpostmsg

import "github.com/tmc/apple/x/signpost"

// ID identifies a signpost so that a begin can be paired with its end. The
// zero value is not usable; obtain one from [Logger.NewID].
type ID uint64

// Reserved signpost id values (see <os/signpost.h>).
const (
	idNull    ID = 0
	idInvalid ID = ^ID(0)
)

// IDExclusive is a shared id usable when at most one interval with a given
// name is in flight at a time on a log.
const IDExclusive ID = ID(signpost.IDExclusive)

// Logger emits signposts against a single os_log handle. The message is
// delivered as a public formatted argument ("%{public}s"), matching the cgo
// implementation, via [signpost.Logger.IntervalBeginMessage] and friends.
type Logger struct {
	log *signpost.Logger
}

// New returns a Logger that emits signposts under the given subsystem and
// category. New never returns nil.
func New(subsystem, category string) *Logger {
	return &Logger{log: signpost.New(subsystem, category)}
}

// Enabled reports whether signposts are being recorded for this log.
func (l *Logger) Enabled() bool {
	if l == nil {
		return false
	}
	return l.log.Enabled()
}

// NewID returns an ID that is unique among signposts logged to this Logger.
func (l *Logger) NewID() ID {
	if l == nil {
		return idNull
	}
	return ID(l.log.NewID())
}

// IntervalBegin marks the start of an interval identified by id, tagged with
// name and carrying msg. Pair it with an [Logger.IntervalEnd] call using the
// same id, name and msg.
func (l *Logger) IntervalBegin(id ID, name Name, msg string) {
	if l == nil {
		return
	}
	l.log.IntervalBeginMessage(signpost.ID(id), name.String(), msg)
}

// IntervalEnd marks the end of the interval begun with the same id, name and
// msg.
func (l *Logger) IntervalEnd(id ID, name Name, msg string) {
	if l == nil {
		return
	}
	l.log.IntervalEndMessage(signpost.ID(id), name.String(), msg)
}

// Event emits a single point-in-time signpost tagged with name and carrying
// msg.
func (l *Logger) Event(id ID, name Name, msg string) {
	if l == nil {
		return
	}
	l.log.EventMessage(signpost.ID(id), name.String(), msg)
}
