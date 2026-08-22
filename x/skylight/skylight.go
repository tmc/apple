package skylight

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/tmc/apple/applicationservices"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/private/skylight"
)

// Space represents a macOS display space (virtual desktop).
type Space uint64

// Window represents a window managed by the WindowServer.
type Window uint32

// ProcessSerialNumber uniquely identifies a process in the macOS window system.
type ProcessSerialNumber = applicationservices.ProcessSerialNumber

var (
	connID     skylight.CGSConnectionID
	connErr    error
	connIDOnce sync.Once
)

func connection() (skylight.CGSConnectionID, error) {
	connIDOnce.Do(func() {
		connID, connErr = skylight.CGSMainConnectionID()
		if connErr != nil || connID == 0 {
			// Fallback attempt with SLSMainConnectionID
			connID, connErr = skylight.SLSMainConnectionID()
		}
	})
	if connErr != nil {
		return 0, fmt.Errorf("skylight: failed to get main connection ID: %w", connErr)
	}
	if connID == 0 {
		return 0, fmt.Errorf("skylight: main connection ID is zero")
	}
	return connID, nil
}

// ActiveSpace returns the identifier for the currently active space.
func ActiveSpace() (Space, error) {
	cid, err := connection()
	if err != nil {
		return 0, err
	}
	sp, err := skylight.SLSGetActiveSpace(cid)
	if err != nil {
		return 0, fmt.Errorf("skylight: SLSGetActiveSpace failed: %w", err)
	}
	return Space(sp), nil
}

// SpacesForWindow returns all space IDs that contain the specified window.
func SpacesForWindow(w Window) ([]Space, error) {
	cid, err := connection()
	if err != nil {
		return nil, err
	}

	// SLSCopySpacesForWindows takes an array of CFNumbers, not of raw ids: the
	// elements are CoreFoundation objects it dereferences. Passing the window id
	// itself as the element is a bus error when the id is read as a pointer.
	wID := int32(w)
	number := corefoundation.CFNumberCreate(0, corefoundation.KCFNumberSInt32Type, unsafe.Pointer(&wID))
	if number == 0 {
		return nil, fmt.Errorf("skylight: CFNumberCreate failed for window %d", w)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(number))

	values := unsafe.Pointer(uintptr(number))
	cfArray := corefoundation.CFArrayCreate(0, unsafe.Pointer(&values), 1, &corefoundation.KCFTypeArrayCallBacks)
	if cfArray == 0 {
		return nil, fmt.Errorf("skylight: CFArrayCreate failed for window %d", w)
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(cfArray))

	spacesArray, err := skylight.SLSCopySpacesForWindows(cid, uint(skylight.CGSAllSpacesMask), cfArray)
	if err != nil {
		return nil, fmt.Errorf("skylight: SLSCopySpacesForWindows failed: %w", err)
	}
	if spacesArray == 0 {
		return nil, nil
	}
	defer corefoundation.CFRelease(corefoundation.CFTypeRef(spacesArray))

	count := corefoundation.CFArrayGetCount(spacesArray)
	if count <= 0 {
		return nil, nil
	}

	// The result is CFNumbers as well, for the same reason.
	spaces := make([]Space, 0, count)
	for i := 0; i < count; i++ {
		val := corefoundation.CFArrayGetValueAtIndex(spacesArray, i)
		if val == nil {
			continue
		}
		var id int64
		if !corefoundation.CFNumberGetValue(corefoundation.CFNumberRef(uintptr(val)), corefoundation.KCFNumberSInt64Type, unsafe.Pointer(&id)) {
			return nil, fmt.Errorf("skylight: CFNumberGetValue failed for space %d of window %d", i, w)
		}
		spaces = append(spaces, Space(id))
	}
	return spaces, nil
}

// IsWindowOffSpace reports whether the window is present on the current active space.
func IsWindowOffSpace(w Window) (bool, error) {
	active, err := ActiveSpace()
	if err != nil {
		return false, err
	}

	spaces, err := SpacesForWindow(w)
	if err != nil {
		return false, err
	}

	for _, s := range spaces {
		if s == active {
			return false, nil
		}
	}
	return true, nil
}

// WindowOwnerPSN returns the Process Serial Number of the process owning the window.
func WindowOwnerPSN(w Window) (ProcessSerialNumber, error) {
	cid, err := connection()
	if err != nil {
		return ProcessSerialNumber{}, err
	}

	var ownerConn skylight.CGSConnectionID
	cgErr, err := skylight.SLSGetWindowOwner(cid, coregraphics.CGWindowID(w), &ownerConn)
	if err != nil {
		return ProcessSerialNumber{}, fmt.Errorf("skylight: SLSGetWindowOwner failed: %w", err)
	}
	if cgErr != 0 {
		return ProcessSerialNumber{}, fmt.Errorf("skylight: SLSGetWindowOwner returned CGError %d", cgErr)
	}

	var psn ProcessSerialNumber
	cgErr, err = skylight.SLSGetConnectionPSN(ownerConn, &psn)
	if err != nil {
		return ProcessSerialNumber{}, fmt.Errorf("skylight: SLSGetConnectionPSN failed: %w", err)
	}
	if cgErr != 0 {
		return ProcessSerialNumber{}, fmt.Errorf("skylight: SLSGetConnectionPSN returned CGError %d", cgErr)
	}

	return psn, nil
}

// WindowOwnerPID returns the process ID of the window owner.
func WindowOwnerPID(w Window) (int, error) {
	psn, err := WindowOwnerPSN(w)
	if err != nil {
		return 0, err
	}
	var pid int32
	status := applicationservices.GetProcessPID(&psn, &pid)
	if status != 0 {
		return 0, fmt.Errorf("skylight: GetProcessPID failed with status %d", status)
	}
	return int(pid), nil
}

// spiError reports a SkyLight SPI failure. Both a non-nil err (the symbol could
// not be called) and a non-zero status (it was called and refused) are failures,
// but only one of the two carries information, so they are reported separately.
func spiError(name string, status int32, err error) error {
	if err != nil {
		return fmt.Errorf("skylight: %s failed: %w", name, err)
	}
	return fmt.Errorf("skylight: %s returned status %d", name, status)
}

// ownerPSN returns the serial number of the process owning w, falling back to
// the one for pid when the window is unknown to the window server. The window
// path is preferred because a process may own windows through more than one
// connection; the pid path keeps a stale or unregistered window id usable.
func ownerPSN(pid int, w Window) (ProcessSerialNumber, error) {
	psn, err := WindowOwnerPSN(w)
	if err == nil {
		return psn, nil
	}
	if pid <= 0 {
		return ProcessSerialNumber{}, err
	}
	var byPID ProcessSerialNumber
	if status := applicationservices.GetProcessForPID(int32(pid), &byPID); status != 0 {
		return ProcessSerialNumber{}, fmt.Errorf("skylight: no PSN for window %d (%v) or PID %d (status %d)", w, err, pid, status)
	}
	return byPID, nil
}

// FocusWithoutRaise makes targetWid the key window of its owning process
// without raising it or changing the window server's z-order, which is what
// makes it usable on a window that is not on the active space.
//
// It posts a defocus record to the process that is currently frontmost and a
// focus record naming targetWid to the target, following yabai's
// window_manager_focus_window_without_raise. It deliberately does not call
// SLPSSetFrontProcessWithOptions: that is what keeps Chromium's user-activation
// gate open, so synthetic input is still treated as user-generated.
//
// targetPID is used only if targetWid cannot be resolved to an owner.
func FocusWithoutRaise(targetPID int, targetWid Window) error {
	var prevPSN ProcessSerialNumber
	status, err := skylight.SLPSGetFrontProcess(&prevPSN)
	if err != nil || status != 0 {
		return spiError("SLPSGetFrontProcess", status, err)
	}

	targetPSN, err := ownerPSN(targetPID, targetWid)
	if err != nil {
		return err
	}

	var rec EventRecord
	rec.DeclaredLength = eventRecordLength
	rec.EventType = 0x0D // kCGSEventAppActive
	rec.WindowID = uint32(targetWid)
	recBytes := rec.Bytes()

	rec.ActivationState = 0x02 // defocus
	status, err = skylight.SLPSPostEventRecordTo(&prevPSN, recBytes)
	if err != nil || status != 0 {
		return spiError("SLPSPostEventRecordTo (defocus)", status, err)
	}

	rec.ActivationState = 0x01 // focus
	status, err = skylight.SLPSPostEventRecordTo(&targetPSN, recBytes)
	if err != nil || status != 0 {
		return spiError("SLPSPostEventRecordTo (focus)", status, err)
	}
	return nil
}

// ActivateForMenuShortcut makes the process owning targetWid frontmost to the
// window server, without raising its windows, so that a key equivalent posted
// to the HID tap reaches its menu. Callers should prefer
// WithMenuShortcutActivation, which restores the previous frontmost process.
//
// targetPID is used only if targetWid cannot be resolved to an owner.
func ActivateForMenuShortcut(targetPID int, targetWid Window) error {
	targetPSN, err := ownerPSN(targetPID, targetWid)
	if err != nil {
		return err
	}
	status, err := skylight.SLPSSetFrontProcessWithOptions(&targetPSN, uint32(targetWid), skylight.CPSModeNoWindows)
	if err != nil || status != 0 {
		return spiError("SLPSSetFrontProcessWithOptions", status, err)
	}
	return nil
}

// WithMenuShortcutActivation activates targetWid for menu shortcut dispatch,
// calls action, and restores the previously frontmost process. The restore runs
// even if action fails or panics. The menu still fires because the key event is
// enqueued on the target's run loop before the restore takes effect.
//
// Unlike the reference implementation in cua-driver, action is not called when
// activation fails: a caller that runs it anyway would post input to whichever
// process happened to be frontmost.
func WithMenuShortcutActivation(targetPID int, targetWid Window, action func() error) error {
	var prevPSN ProcessSerialNumber
	status, err := skylight.SLPSGetFrontProcess(&prevPSN)
	hasPrev := err == nil && status == 0

	if err := ActivateForMenuShortcut(targetPID, targetWid); err != nil {
		return err
	}
	defer func() {
		if hasPrev {
			_, _ = skylight.SLPSSetFrontProcessWithOptions(&prevPSN, 0, skylight.CPSModeNoWindows)
		}
	}()

	return action()
}
