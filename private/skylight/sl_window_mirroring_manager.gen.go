// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLWindowMirroringManager] class.
var (
	_SLWindowMirroringManagerClass     SLWindowMirroringManagerClass
	_SLWindowMirroringManagerClassOnce sync.Once
)

func getSLWindowMirroringManagerClass() SLWindowMirroringManagerClass {
	_SLWindowMirroringManagerClassOnce.Do(func() {
		_SLWindowMirroringManagerClass = SLWindowMirroringManagerClass{class: objc.GetClass("SLWindowMirroringManager")}
	})
	return _SLWindowMirroringManagerClass
}

// GetSLWindowMirroringManagerClass returns the class object for SLWindowMirroringManager.
func GetSLWindowMirroringManagerClass() SLWindowMirroringManagerClass {
	return getSLWindowMirroringManagerClass()
}

type SLWindowMirroringManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLWindowMirroringManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLWindowMirroringManagerClass) Alloc() SLWindowMirroringManager {
	rv := objc.Send[SLWindowMirroringManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLWindowMirroringManager.AddAppTo]
//   - [SLWindowMirroringManager.AddWindowTo]
//   - [SLWindowMirroringManager.AppsMirroredToDisplay]
//   - [SLWindowMirroringManager.ContextForDisplay]
//   - [SLWindowMirroringManager.Delegate]
//   - [SLWindowMirroringManager.SetDelegate]
//   - [SLWindowMirroringManager.Extend]
//   - [SLWindowMirroringManager.FilterForDisplay]
//   - [SLWindowMirroringManager.InvokeDelegateContentChanged]
//   - [SLWindowMirroringManager.MirrorTo]
//   - [SLWindowMirroringManager.MirrorWithIncludedAppsTo]
//   - [SLWindowMirroringManager.MirrorWithIncludedWindowsAndIncludedAppsTo]
//   - [SLWindowMirroringManager.MirrorWithIncludedWindowsTo]
//   - [SLWindowMirroringManager.RemoveAllContentsFrom]
//   - [SLWindowMirroringManager.RemoveAppFrom]
//   - [SLWindowMirroringManager.RemoveWindowFrom]
//   - [SLWindowMirroringManager.ReplaceAllContentsWithAppTo]
//   - [SLWindowMirroringManager.ReplaceAllContentsWithWindowTo]
//   - [SLWindowMirroringManager.ReplaceAllWindowsTo]
//   - [SLWindowMirroringManager.ReplaceWindowWithTo]
//   - [SLWindowMirroringManager.ResetSession]
//   - [SLWindowMirroringManager.ShieldWindowForDisplay]
//   - [SLWindowMirroringManager.SupportWindowMirroring]
//   - [SLWindowMirroringManager.UpdateWithShieldWindowTo]
//   - [SLWindowMirroringManager.WindowMirroringDisplays]
//   - [SLWindowMirroringManager.WindowsMirroredToDisplay]
//
// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager
type SLWindowMirroringManager struct {
	objectivec.Object
}

// SLWindowMirroringManagerFromID constructs a [SLWindowMirroringManager] from an objc.ID.
func SLWindowMirroringManagerFromID(id objc.ID) SLWindowMirroringManager {
	return SLWindowMirroringManager{objectivec.Object{ID: id}}
}

// Ensure SLWindowMirroringManager implements ISLWindowMirroringManager.
var _ ISLWindowMirroringManager = SLWindowMirroringManager{}

// An interface definition for the [SLWindowMirroringManager] class.
//
// # Methods
//
//   - [ISLWindowMirroringManager.AddAppTo]
//   - [ISLWindowMirroringManager.AddWindowTo]
//   - [ISLWindowMirroringManager.AppsMirroredToDisplay]
//   - [ISLWindowMirroringManager.ContextForDisplay]
//   - [ISLWindowMirroringManager.Delegate]
//   - [ISLWindowMirroringManager.SetDelegate]
//   - [ISLWindowMirroringManager.Extend]
//   - [ISLWindowMirroringManager.FilterForDisplay]
//   - [ISLWindowMirroringManager.InvokeDelegateContentChanged]
//   - [ISLWindowMirroringManager.MirrorTo]
//   - [ISLWindowMirroringManager.MirrorWithIncludedAppsTo]
//   - [ISLWindowMirroringManager.MirrorWithIncludedWindowsAndIncludedAppsTo]
//   - [ISLWindowMirroringManager.MirrorWithIncludedWindowsTo]
//   - [ISLWindowMirroringManager.RemoveAllContentsFrom]
//   - [ISLWindowMirroringManager.RemoveAppFrom]
//   - [ISLWindowMirroringManager.RemoveWindowFrom]
//   - [ISLWindowMirroringManager.ReplaceAllContentsWithAppTo]
//   - [ISLWindowMirroringManager.ReplaceAllContentsWithWindowTo]
//   - [ISLWindowMirroringManager.ReplaceAllWindowsTo]
//   - [ISLWindowMirroringManager.ReplaceWindowWithTo]
//   - [ISLWindowMirroringManager.ResetSession]
//   - [ISLWindowMirroringManager.ShieldWindowForDisplay]
//   - [ISLWindowMirroringManager.SupportWindowMirroring]
//   - [ISLWindowMirroringManager.UpdateWithShieldWindowTo]
//   - [ISLWindowMirroringManager.WindowMirroringDisplays]
//   - [ISLWindowMirroringManager.WindowsMirroredToDisplay]
//
// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager
type ISLWindowMirroringManager interface {
	objectivec.IObject

	// Topic: Methods

	AddAppTo(app objectivec.IObject, to objectivec.IObject) bool
	AddWindowTo(window objectivec.IObject, to objectivec.IObject) bool
	AppsMirroredToDisplay(display objectivec.IObject) objectivec.IObject
	ContextForDisplay(display objectivec.IObject) objectivec.IObject
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	Extend(extend objectivec.IObject) bool
	FilterForDisplay(display objectivec.IObject) objectivec.IObject
	InvokeDelegateContentChanged(changed objectivec.IObject)
	MirrorTo(mirror objectivec.IObject, to objectivec.IObject) bool
	MirrorWithIncludedAppsTo(mirror objectivec.IObject, apps objectivec.IObject, to objectivec.IObject) bool
	MirrorWithIncludedWindowsAndIncludedAppsTo(mirror objectivec.IObject, windows objectivec.IObject, apps objectivec.IObject, to objectivec.IObject) bool
	MirrorWithIncludedWindowsTo(mirror objectivec.IObject, windows objectivec.IObject, to objectivec.IObject) bool
	RemoveAllContentsFrom(from objectivec.IObject) bool
	RemoveAppFrom(app objectivec.IObject, from objectivec.IObject) bool
	RemoveWindowFrom(window objectivec.IObject, from objectivec.IObject) bool
	ReplaceAllContentsWithAppTo(app objectivec.IObject, to objectivec.IObject) bool
	ReplaceAllContentsWithWindowTo(window objectivec.IObject, to objectivec.IObject) bool
	ReplaceAllWindowsTo(windows objectivec.IObject, to objectivec.IObject) bool
	ReplaceWindowWithTo(window objectivec.IObject, with objectivec.IObject, to objectivec.IObject) bool
	ResetSession(session objectivec.IObject)
	ShieldWindowForDisplay(display objectivec.IObject) objectivec.IObject
	SupportWindowMirroring() bool
	UpdateWithShieldWindowTo(update objectivec.IObject, window objectivec.IObject, to objectivec.IObject) bool
	WindowMirroringDisplays() objectivec.IObject
	WindowsMirroredToDisplay(display objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (s SLWindowMirroringManager) Init() SLWindowMirroringManager {
	rv := objc.Send[SLWindowMirroringManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLWindowMirroringManager) Autorelease() SLWindowMirroringManager {
	rv := objc.Send[SLWindowMirroringManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLWindowMirroringManager creates a new SLWindowMirroringManager instance.
func NewSLWindowMirroringManager() SLWindowMirroringManager {
	class := getSLWindowMirroringManagerClass()
	rv := objc.Send[SLWindowMirroringManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/addApp:to:
func (s SLWindowMirroringManager) AddAppTo(app objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("addApp:to:"), app, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/addWindow:to:
func (s SLWindowMirroringManager) AddWindowTo(window objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("addWindow:to:"), window, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/appsMirroredToDisplay:
func (s SLWindowMirroringManager) AppsMirroredToDisplay(display objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("appsMirroredToDisplay:"), display)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/contextForDisplay:
func (s SLWindowMirroringManager) ContextForDisplay(display objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("contextForDisplay:"), display)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/extend:
func (s SLWindowMirroringManager) Extend(extend objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("extend:"), extend)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/filterForDisplay:
func (s SLWindowMirroringManager) FilterForDisplay(display objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filterForDisplay:"), display)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/invokeDelegateContentChanged:
func (s SLWindowMirroringManager) InvokeDelegateContentChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("invokeDelegateContentChanged:"), changed)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/mirror:to:
func (s SLWindowMirroringManager) MirrorTo(mirror objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mirror:to:"), mirror, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/mirror:withIncludedApps:to:
func (s SLWindowMirroringManager) MirrorWithIncludedAppsTo(mirror objectivec.IObject, apps objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mirror:withIncludedApps:to:"), mirror, apps, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/mirror:withIncludedWindows:andIncludedApps:to:
func (s SLWindowMirroringManager) MirrorWithIncludedWindowsAndIncludedAppsTo(mirror objectivec.IObject, windows objectivec.IObject, apps objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mirror:withIncludedWindows:andIncludedApps:to:"), mirror, windows, apps, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/mirror:withIncludedWindows:to:
func (s SLWindowMirroringManager) MirrorWithIncludedWindowsTo(mirror objectivec.IObject, windows objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mirror:withIncludedWindows:to:"), mirror, windows, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/removeAllContentsFrom:
func (s SLWindowMirroringManager) RemoveAllContentsFrom(from objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("removeAllContentsFrom:"), from)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/removeApp:from:
func (s SLWindowMirroringManager) RemoveAppFrom(app objectivec.IObject, from objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("removeApp:from:"), app, from)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/removeWindow:from:
func (s SLWindowMirroringManager) RemoveWindowFrom(window objectivec.IObject, from objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("removeWindow:from:"), window, from)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/replaceAllContentsWithApp:to:
func (s SLWindowMirroringManager) ReplaceAllContentsWithAppTo(app objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("replaceAllContentsWithApp:to:"), app, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/replaceAllContentsWithWindow:to:
func (s SLWindowMirroringManager) ReplaceAllContentsWithWindowTo(window objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("replaceAllContentsWithWindow:to:"), window, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/replaceAllWindows:to:
func (s SLWindowMirroringManager) ReplaceAllWindowsTo(windows objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("replaceAllWindows:to:"), windows, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/replaceWindow:with:to:
func (s SLWindowMirroringManager) ReplaceWindowWithTo(window objectivec.IObject, with objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("replaceWindow:with:to:"), window, with, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/resetSession:
func (s SLWindowMirroringManager) ResetSession(session objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("resetSession:"), session)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/shieldWindowForDisplay:
func (s SLWindowMirroringManager) ShieldWindowForDisplay(display objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("shieldWindowForDisplay:"), display)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/supportWindowMirroring
func (s SLWindowMirroringManager) SupportWindowMirroring() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("supportWindowMirroring"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/update:withShieldWindow:to:
func (s SLWindowMirroringManager) UpdateWithShieldWindowTo(update objectivec.IObject, window objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("update:withShieldWindow:to:"), update, window, to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/windowMirroringDisplays
func (s SLWindowMirroringManager) WindowMirroringDisplays() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windowMirroringDisplays"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/windowsMirroredToDisplay:
func (s SLWindowMirroringManager) WindowsMirroredToDisplay(display objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("windowsMirroredToDisplay:"), display)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/shared
func (_SLWindowMirroringManagerClass SLWindowMirroringManagerClass) Shared() SLWindowMirroringManager {
	rv := objc.Send[objc.ID](objc.ID(_SLWindowMirroringManagerClass.class), objc.Sel("shared"))
	return SLWindowMirroringManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringManager/delegate
func (s SLWindowMirroringManager) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (s SLWindowMirroringManager) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
