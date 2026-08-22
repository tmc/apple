// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLWindowMirroringFilter] class.
var (
	_SLWindowMirroringFilterClass     SLWindowMirroringFilterClass
	_SLWindowMirroringFilterClassOnce sync.Once
)

func getSLWindowMirroringFilterClass() SLWindowMirroringFilterClass {
	_SLWindowMirroringFilterClassOnce.Do(func() {
		_SLWindowMirroringFilterClass = SLWindowMirroringFilterClass{class: objc.GetClass("SLWindowMirroringFilter")}
	})
	return _SLWindowMirroringFilterClass
}

// GetSLWindowMirroringFilterClass returns the class object for SLWindowMirroringFilter.
func GetSLWindowMirroringFilterClass() SLWindowMirroringFilterClass {
	return getSLWindowMirroringFilterClass()
}

type SLWindowMirroringFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLWindowMirroringFilterClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLWindowMirroringFilterClass) Alloc() SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLWindowMirroringFilter.SetShieldWindow]
//   - [SLWindowMirroringFilter.ShieldWindow]
//   - [SLWindowMirroringFilter.WindowFilter]
//   - [SLWindowMirroringFilter.InitWithFilter]
//   - [SLWindowMirroringFilter.InitWithIncludedApps]
//   - [SLWindowMirroringFilter.InitWithIncludedWindows]
//   - [SLWindowMirroringFilter.InitWithIncludedWindowsAndIncludedApps]
type SLWindowMirroringFilter struct {
	objectivec.Object
}

// SLWindowMirroringFilterFromID constructs a [SLWindowMirroringFilter] from an objc.ID.
func SLWindowMirroringFilterFromID(id objc.ID) SLWindowMirroringFilter {
	return SLWindowMirroringFilter{objectivec.Object{ID: id}}
}

// Ensure SLWindowMirroringFilter implements ISLWindowMirroringFilter.
var _ ISLWindowMirroringFilter = SLWindowMirroringFilter{}

// An interface definition for the [SLWindowMirroringFilter] class.
//
// # Methods
//
//   - [ISLWindowMirroringFilter.SetShieldWindow]
//   - [ISLWindowMirroringFilter.ShieldWindow]
//   - [ISLWindowMirroringFilter.WindowFilter]
//   - [ISLWindowMirroringFilter.InitWithFilter]
//   - [ISLWindowMirroringFilter.InitWithIncludedApps]
//   - [ISLWindowMirroringFilter.InitWithIncludedWindows]
//   - [ISLWindowMirroringFilter.InitWithIncludedWindowsAndIncludedApps]
type ISLWindowMirroringFilter interface {
	objectivec.IObject

	// Topic: Methods

	SetShieldWindow(window objectivec.IObject)
	ShieldWindow() objectivec.IObject
	WindowFilter() objectivec.IObject
	InitWithFilter(filter objectivec.IObject) SLWindowMirroringFilter
	InitWithIncludedApps(apps objectivec.IObject) SLWindowMirroringFilter
	InitWithIncludedWindows(windows objectivec.IObject) SLWindowMirroringFilter
	InitWithIncludedWindowsAndIncludedApps(windows objectivec.IObject, apps objectivec.IObject) SLWindowMirroringFilter
}

// Init initializes the instance.
func (s SLWindowMirroringFilter) Init() SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLWindowMirroringFilter) Autorelease() SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLWindowMirroringFilter creates a new SLWindowMirroringFilter instance.
func NewSLWindowMirroringFilter() SLWindowMirroringFilter {
	class := getSLWindowMirroringFilterClass()
	rv := objc.SendIfResponds[SLWindowMirroringFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLWindowMirroringFilterWithFilter(filter objectivec.IObject) SLWindowMirroringFilter {
	instance := getSLWindowMirroringFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFilter:"), filter)
	return SLWindowMirroringFilterFromID(rv)
}

func NewSLWindowMirroringFilterWithIncludedApps(apps objectivec.IObject) SLWindowMirroringFilter {
	instance := getSLWindowMirroringFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIncludedApps:"), apps)
	return SLWindowMirroringFilterFromID(rv)
}

func NewSLWindowMirroringFilterWithIncludedWindows(windows objectivec.IObject) SLWindowMirroringFilter {
	instance := getSLWindowMirroringFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIncludedWindows:"), windows)
	return SLWindowMirroringFilterFromID(rv)
}

func NewSLWindowMirroringFilterWithIncludedWindowsAndIncludedApps(windows objectivec.IObject, apps objectivec.IObject) SLWindowMirroringFilter {
	instance := getSLWindowMirroringFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIncludedWindows:andIncludedApps:"), windows, apps)
	return SLWindowMirroringFilterFromID(rv)
}

func (s SLWindowMirroringFilter) SetShieldWindow(window objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setShieldWindow:"), window)
}
func (s SLWindowMirroringFilter) ShieldWindow() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("shieldWindow"))
	return objectivec.Object{ID: rv}
}
func (s SLWindowMirroringFilter) WindowFilter() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("windowFilter"))
	return objectivec.Object{ID: rv}
}
func (s SLWindowMirroringFilter) InitWithFilter(filter objectivec.IObject) SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](s.ID, objc.Sel("initWithFilter:"), filter)
	return rv
}
func (s SLWindowMirroringFilter) InitWithIncludedApps(apps objectivec.IObject) SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](s.ID, objc.Sel("initWithIncludedApps:"), apps)
	return rv
}
func (s SLWindowMirroringFilter) InitWithIncludedWindows(windows objectivec.IObject) SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](s.ID, objc.Sel("initWithIncludedWindows:"), windows)
	return rv
}
func (s SLWindowMirroringFilter) InitWithIncludedWindowsAndIncludedApps(windows objectivec.IObject, apps objectivec.IObject) SLWindowMirroringFilter {
	rv := objc.SendIfResponds[SLWindowMirroringFilter](s.ID, objc.Sel("initWithIncludedWindows:andIncludedApps:"), windows, apps)
	return rv
}
