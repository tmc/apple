//go:build darwin

package objcinspect

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync/atomic"

	"github.com/tmc/apple/objc"
)

var enabled atomic.Bool

func init() {
	if v := os.Getenv("OBJCCHECK"); v != "" && v != "0" {
		enabled.Store(true)
	}
}

// SetEnabled turns the SendChecked and CallChecked preflight on or off at run
// time and reports the previous setting. It defaults to on when the OBJCCHECK
// environment variable is set to a value other than "" or "0".
func SetEnabled(on bool) bool { return enabled.Swap(on) }

// Enabled reports whether the preflight is on.
func Enabled() bool { return enabled.Load() }

// SendChecked sends sel to id and returns the result. When the preflight is
// enabled it first runs Check against id's live class; if the call would not
// match—a missing selector, the wrong argument count, or a return or argument
// type in the wrong register family—it returns the zero value of T and the
// Check error instead of dispatching. When the preflight is off it is objc.Send
// plus one atomic load and always returns a nil error.
//
// T is the type the caller will read the result as, so the preflight verifies
// the method's return against it. For a method whose result is void or ignored,
// use CallChecked. Check's limits carry over: it is family- not width-level, so
// SendChecked does not catch integer width truncation, and it accepts a struct
// where an object is expected.
func SendChecked[T any](id objc.ID, sel objc.SEL, args ...any) (T, error) {
	if enabled.Load() {
		if err := Check(id, sel, reflect.TypeFor[T](), args...); err != nil {
			var zero T
			return zero, err
		}
	}
	return objc.Send[T](id, sel, args...), nil
}

// CallChecked sends sel to id for its effect, discarding the result. When the
// preflight is enabled it runs Check with no return-type constraint, so it
// verifies the selector and arguments but not the return. Use it for void
// methods and for calls whose result you do not read.
func CallChecked(id objc.ID, sel objc.SEL, args ...any) error {
	if enabled.Load() {
		if err := Check(id, sel, nil, args...); err != nil {
			return err
		}
	}
	objc.Send[objc.ID](id, sel, args...)
	return nil
}

// A Call describes one intended objc.Send: a live receiver, the selector, the
// return type the caller expects (nil to skip the return check), and the
// arguments it will pass. It is the unit of a conformance sweep.
type Call struct {
	ID     objc.ID
	Sel    objc.SEL
	Return reflect.Type
	Args   []any
}

// CheckAll runs Check on every call and joins the failures, returning nil when
// all pass. It does not dispatch any calls or depend on the enabled flag.
func CheckAll(calls ...Call) error {
	var errs []error
	for i, c := range calls {
		if err := Check(c.ID, c.Sel, c.Return, c.Args...); err != nil {
			errs = append(errs, fmt.Errorf("call %d %v: %w", i, c.Sel, err))
		}
	}
	return errors.Join(errs...)
}
