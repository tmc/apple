// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLContentFilter] class.
var (
	_SLContentFilterClass     SLContentFilterClass
	_SLContentFilterClassOnce sync.Once
)

func getSLContentFilterClass() SLContentFilterClass {
	_SLContentFilterClassOnce.Do(func() {
		_SLContentFilterClass = SLContentFilterClass{class: objc.GetClass("SLContentFilter")}
	})
	return _SLContentFilterClass
}

// GetSLContentFilterClass returns the class object for SLContentFilter.
func GetSLContentFilterClass() SLContentFilterClass {
	return getSLContentFilterClass()
}

type SLContentFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLContentFilterClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLContentFilterClass) Alloc() SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLContentFilter.ApplicationID]
//   - [SLContentFilter.DisplayID]
//   - [SLContentFilter.EncodeWithCoder]
//   - [SLContentFilter.ExcludeMenuBar]
//   - [SLContentFilter.ExcludedApplications]
//   - [SLContentFilter.ExcludedPIDS]
//   - [SLContentFilter.SetExcludedPIDS]
//   - [SLContentFilter.ExcludedWindows]
//   - [SLContentFilter.FilterType]
//   - [SLContentFilter.GetFilterType]
//   - [SLContentFilter.HideMenuBar]
//   - [SLContentFilter.SetHideMenuBar]
//   - [SLContentFilter.IncludedApplications]
//   - [SLContentFilter.IncludedPIDS]
//   - [SLContentFilter.SetIncludedPIDS]
//   - [SLContentFilter.IncludedWindows]
//   - [SLContentFilter.ShareAll]
//   - [SLContentFilter.WindowID]
//   - [SLContentFilter.InitWithCoder]
//   - [SLContentFilter.InitWithDesktopIndependentWindow]
//   - [SLContentFilter.InitWithDisplay]
//   - [SLContentFilter.InitWithDisplayApplication]
//   - [SLContentFilter.InitWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications]
//   - [SLContentFilter.InitWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS]
//   - [SLContentFilter.InitWithDisplayWindow]
type SLContentFilter struct {
	objectivec.Object
}

// SLContentFilterFromID constructs a [SLContentFilter] from an objc.ID.
func SLContentFilterFromID(id objc.ID) SLContentFilter {
	return SLContentFilter{objectivec.Object{ID: id}}
}

// Ensure SLContentFilter implements ISLContentFilter.
var _ ISLContentFilter = SLContentFilter{}

// An interface definition for the [SLContentFilter] class.
//
// # Methods
//
//   - [ISLContentFilter.ApplicationID]
//   - [ISLContentFilter.DisplayID]
//   - [ISLContentFilter.EncodeWithCoder]
//   - [ISLContentFilter.ExcludeMenuBar]
//   - [ISLContentFilter.ExcludedApplications]
//   - [ISLContentFilter.ExcludedPIDS]
//   - [ISLContentFilter.SetExcludedPIDS]
//   - [ISLContentFilter.ExcludedWindows]
//   - [ISLContentFilter.FilterType]
//   - [ISLContentFilter.GetFilterType]
//   - [ISLContentFilter.HideMenuBar]
//   - [ISLContentFilter.SetHideMenuBar]
//   - [ISLContentFilter.IncludedApplications]
//   - [ISLContentFilter.IncludedPIDS]
//   - [ISLContentFilter.SetIncludedPIDS]
//   - [ISLContentFilter.IncludedWindows]
//   - [ISLContentFilter.ShareAll]
//   - [ISLContentFilter.WindowID]
//   - [ISLContentFilter.InitWithCoder]
//   - [ISLContentFilter.InitWithDesktopIndependentWindow]
//   - [ISLContentFilter.InitWithDisplay]
//   - [ISLContentFilter.InitWithDisplayApplication]
//   - [ISLContentFilter.InitWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications]
//   - [ISLContentFilter.InitWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS]
//   - [ISLContentFilter.InitWithDisplayWindow]
type ISLContentFilter interface {
	objectivec.IObject

	// Topic: Methods

	ApplicationID() string
	DisplayID() uint32
	EncodeWithCoder(coder foundation.INSCoder)
	ExcludeMenuBar(bar bool)
	ExcludedApplications() foundation.INSSet
	ExcludedPIDS() foundation.INSSet
	SetExcludedPIDS(value foundation.INSSet)
	ExcludedWindows() foundation.INSSet
	FilterType() uint32
	GetFilterType() uint32
	HideMenuBar() bool
	SetHideMenuBar(value bool)
	IncludedApplications() foundation.INSSet
	IncludedPIDS() foundation.INSSet
	SetIncludedPIDS(value foundation.INSSet)
	IncludedWindows() foundation.INSSet
	ShareAll() bool
	WindowID() uint32
	InitWithCoder(coder foundation.INSCoder) SLContentFilter
	InitWithDesktopIndependentWindow(window uint32) SLContentFilter
	InitWithDisplay(display uint32) SLContentFilter
	InitWithDisplayApplication(display uint32, application objectivec.IObject) SLContentFilter
	InitWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject) SLContentFilter
	InitWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, pids objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject, pids2 objectivec.IObject) SLContentFilter
	InitWithDisplayWindow(display uint32, window uint32) SLContentFilter
}

// Init initializes the instance.
func (s SLContentFilter) Init() SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLContentFilter) Autorelease() SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLContentFilter creates a new SLContentFilter instance.
func NewSLContentFilter() SLContentFilter {
	class := getSLContentFilterClass()
	rv := objc.SendIfResponds[SLContentFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLContentFilterWithCoder(coder objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLContentFilterFromID(rv)
}

func NewSLContentFilterWithDesktopIndependentWindow(window uint32) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	return SLContentFilterFromID(rv)
}

func NewSLContentFilterWithDisplay(display uint32) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplay:"), display)
	return SLContentFilterFromID(rv)
}

func NewSLContentFilterWithDisplayApplication(display uint32, application objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplay:application:"), display, application)
	return SLContentFilterFromID(rv)
}

func NewSLContentFilterWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:excludedWindows:excludedApplications:"), display, all, windows, applications, windows2, applications2)
	return SLContentFilterFromID(rv)
}

func NewSLContentFilterWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, pids objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject, pids2 objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:includedPIDS:excludedWindows:excludedApplications:excludedPIDS:"), display, all, windows, applications, pids, windows2, applications2, pids2)
	return SLContentFilterFromID(rv)
}

func NewSLContentFilterWithDisplayWindow(display uint32, window uint32) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDisplay:window:"), display, window)
	return SLContentFilterFromID(rv)
}

func (s SLContentFilter) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (s SLContentFilter) ExcludeMenuBar(bar bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("excludeMenuBar:"), bar)
}
func (s SLContentFilter) GetFilterType() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("getFilterType"))
	return rv
}
func (s SLContentFilter) InitWithCoder(coder foundation.INSCoder) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s SLContentFilter) InitWithDesktopIndependentWindow(window uint32) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	return rv
}
func (s SLContentFilter) InitWithDisplay(display uint32) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithDisplay:"), display)
	return rv
}
func (s SLContentFilter) InitWithDisplayApplication(display uint32, application objectivec.IObject) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithDisplay:application:"), display, application)
	return rv
}
func (s SLContentFilter) InitWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:excludedWindows:excludedApplications:"), display, all, windows, applications, windows2, applications2)
	return rv
}
func (s SLContentFilter) InitWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, pids objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject, pids2 objectivec.IObject) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:includedPIDS:excludedWindows:excludedApplications:excludedPIDS:"), display, all, windows, applications, pids, windows2, applications2, pids2)
	return rv
}
func (s SLContentFilter) InitWithDisplayWindow(display uint32, window uint32) SLContentFilter {
	rv := objc.SendIfResponds[SLContentFilter](s.ID, objc.Sel("initWithDisplay:window:"), display, window)
	return rv
}

func (_SLContentFilterClass SLContentFilterClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_SLContentFilterClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (s SLContentFilter) ApplicationID() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("applicationID"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLContentFilter) DisplayID() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("displayID"))
	return rv
}
func (s SLContentFilter) ExcludedApplications() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("excludedApplications"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) ExcludedPIDS() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("excludedPIDS"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) SetExcludedPIDS(value foundation.INSSet) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setExcludedPIDS:"), value)
}
func (s SLContentFilter) ExcludedWindows() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("excludedWindows"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) FilterType() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("filterType"))
	return rv
}
func (s SLContentFilter) HideMenuBar() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("hideMenuBar"))
	return rv
}
func (s SLContentFilter) SetHideMenuBar(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setHideMenuBar:"), value)
}
func (s SLContentFilter) IncludedApplications() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("includedApplications"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) IncludedPIDS() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("includedPIDS"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) SetIncludedPIDS(value foundation.INSSet) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setIncludedPIDS:"), value)
}
func (s SLContentFilter) IncludedWindows() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("includedWindows"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) ShareAll() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("shareAll"))
	return rv
}
func (s SLContentFilter) WindowID() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
