package oslog

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// An Activity groups related log messages under a common identifier. Every
// os_log message emitted while an activity's scope is active is tagged with the
// activity's id, so a later reader (Console.app, `log show`, Instruments) can
// collapse all the work of one logical operation together, even across threads.
//
// Create an activity, enter its scope around the work, and leave it — typically
// with defer:
//
//	act := oslog.NewActivity("handle request")
//	defer act.Enter()()
//	log.Info("started")   // tagged with act.ID()
type Activity struct {
	handle uintptr // os_activity_t
}

// ActivityFlag controls how a new activity relates to the current one.
type ActivityFlag uint32

const (
	// ActivityDefault nests the new activity under the current one.
	ActivityDefault ActivityFlag = 0
	// ActivityDetached makes the new activity a root, ignoring any current one.
	ActivityDetached ActivityFlag = 0x1
	// ActivityIfNonePresent creates the activity only if one is not already
	// current; otherwise the new activity is a no-op that adopts the current.
	ActivityIfNonePresent ActivityFlag = 0x2
)

var (
	activityOnce sync.Once

	osActivityCreate    func(dso unsafe.Pointer, desc *byte, parent uintptr, flags uint32) uintptr
	osActivityScopeEnt  func(act uintptr, state unsafe.Pointer)
	osActivityScopeLeav func(state unsafe.Pointer)
	osActivityGetID     func(act uintptr, parent *uint64) uint64

	activityCurrent uintptr // _os_activity_current sentinel
)

func loadActivitySymbols() {
	activityOnce.Do(func() {
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "_os_activity_create"); err == nil && s != 0 {
			purego.RegisterFunc(&osActivityCreate, s)
		}
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_activity_scope_enter"); err == nil && s != 0 {
			purego.RegisterFunc(&osActivityScopeEnt, s)
		}
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_activity_scope_leave"); err == nil && s != 0 {
			purego.RegisterFunc(&osActivityScopeLeav, s)
		}
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "os_activity_get_identifier"); err == nil && s != 0 {
			purego.RegisterFunc(&osActivityGetID, s)
		}
		if s, err := purego.Dlsym(purego.RTLD_DEFAULT, "_os_activity_current"); err == nil {
			activityCurrent = s
		}
	})
}

// NewActivity creates an activity with the given description, nested under the
// current activity. It never returns nil; if the symbols cannot be resolved the
// activity's methods are no-ops.
func NewActivity(description string) *Activity {
	return NewActivityFlags(description, ActivityDefault)
}

// NewActivityFlags creates an activity with an explicit flag (for example
// [ActivityDetached] to start a new root activity).
func NewActivityFlags(description string, flags ActivityFlag) *Activity {
	loadActivitySymbols()
	a := &Activity{}
	if osActivityCreate != nil {
		desc := cstring(description)
		a.handle = osActivityCreate(dsoHandle(), &desc[0], activityCurrent, uint32(flags))
		runtime.KeepAlive(desc)
	}
	return a
}

// ID returns the activity's identifier, or 0 if unavailable. It is the value
// that tags log entries emitted within the activity's scope.
func (a *Activity) ID() uint64 {
	if a == nil || a.handle == 0 || osActivityGetID == nil {
		return 0
	}
	var parent uint64
	return osActivityGetID(a.handle, &parent)
}

// Enter makes the activity current on the calling goroutine's thread and
// returns a function that leaves the scope. Because scope enter/leave must
// happen on the same OS thread, Enter locks the goroutine to its thread until
// the returned function is called; use it with defer:
//
//	defer act.Enter()()
func (a *Activity) Enter() (leave func()) {
	if a == nil || a.handle == 0 || osActivityScopeEnt == nil {
		return func() {}
	}
	runtime.LockOSThread()
	// os_activity_scope_state_s is uint64 opaque[2] (16 bytes). It must outlive
	// the scope, so it is heap-allocated and captured by the returned closure.
	state := new([2]uint64)
	osActivityScopeEnt(a.handle, unsafe.Pointer(state))
	return func() {
		osActivityScopeLeav(unsafe.Pointer(state))
		runtime.KeepAlive(state)
		runtime.UnlockOSThread()
	}
}
