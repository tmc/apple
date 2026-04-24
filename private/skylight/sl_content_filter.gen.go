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
	rv := objc.Send[SLContentFilter](objc.ID(sc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter
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
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLContentFilter) Autorelease() SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLContentFilter creates a new SLContentFilter instance.
func NewSLContentFilter() SLContentFilter {
	class := getSLContentFilterClass()
	rv := objc.Send[SLContentFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithCoder:
func NewSLContentFilterWithCoder(coder objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDesktopIndependentWindow:
func NewSLContentFilterWithDesktopIndependentWindow(window uint32) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:
func NewSLContentFilterWithDisplay(display uint32) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:"), display)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:application:
func NewSLContentFilterWithDisplayApplication(display uint32, application objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:application:"), display, application)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:shareAll:includedWindows:includedApplications:excludedWindows:excludedApplications:
func NewSLContentFilterWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:excludedWindows:excludedApplications:"), display, all, windows, applications, windows2, applications2)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:shareAll:includedWindows:includedApplications:includedPIDS:excludedWindows:excludedApplications:excludedPIDS:
func NewSLContentFilterWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, pids objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject, pids2 objectivec.IObject) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:includedPIDS:excludedWindows:excludedApplications:excludedPIDS:"), display, all, windows, applications, pids, windows2, applications2, pids2)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:window:
func NewSLContentFilterWithDisplayWindow(display uint32, window uint32) SLContentFilter {
	instance := getSLContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:window:"), display, window)
	return SLContentFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/encodeWithCoder:
func (s SLContentFilter) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/excludeMenuBar:
func (s SLContentFilter) ExcludeMenuBar(bar bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("excludeMenuBar:"), bar)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/getFilterType
func (s SLContentFilter) GetFilterType() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("getFilterType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithCoder:
func (s SLContentFilter) InitWithCoder(coder foundation.INSCoder) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDesktopIndependentWindow:
func (s SLContentFilter) InitWithDesktopIndependentWindow(window uint32) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:
func (s SLContentFilter) InitWithDisplay(display uint32) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithDisplay:"), display)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:application:
func (s SLContentFilter) InitWithDisplayApplication(display uint32, application objectivec.IObject) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithDisplay:application:"), display, application)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:shareAll:includedWindows:includedApplications:excludedWindows:excludedApplications:
func (s SLContentFilter) InitWithDisplayShareAllIncludedWindowsIncludedApplicationsExcludedWindowsExcludedApplications(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:excludedWindows:excludedApplications:"), display, all, windows, applications, windows2, applications2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:shareAll:includedWindows:includedApplications:includedPIDS:excludedWindows:excludedApplications:excludedPIDS:
func (s SLContentFilter) InitWithDisplayShareAllIncludedWindowsIncludedApplicationsIncludedPIDSExcludedWindowsExcludedApplicationsExcludedPIDS(display uint32, all bool, windows objectivec.IObject, applications objectivec.IObject, pids objectivec.IObject, windows2 objectivec.IObject, applications2 objectivec.IObject, pids2 objectivec.IObject) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithDisplay:shareAll:includedWindows:includedApplications:includedPIDS:excludedWindows:excludedApplications:excludedPIDS:"), display, all, windows, applications, pids, windows2, applications2, pids2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/initWithDisplay:window:
func (s SLContentFilter) InitWithDisplayWindow(display uint32, window uint32) SLContentFilter {
	rv := objc.Send[SLContentFilter](s.ID, objc.Sel("initWithDisplay:window:"), display, window)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/supportsSecureCoding
func (_SLContentFilterClass SLContentFilterClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLContentFilterClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/applicationID
func (s SLContentFilter) ApplicationID() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("applicationID"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/displayID
func (s SLContentFilter) DisplayID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("displayID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/excludedApplications
func (s SLContentFilter) ExcludedApplications() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("excludedApplications"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/excludedPIDS
func (s SLContentFilter) ExcludedPIDS() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("excludedPIDS"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) SetExcludedPIDS(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setExcludedPIDS:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/excludedWindows
func (s SLContentFilter) ExcludedWindows() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("excludedWindows"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/filterType
func (s SLContentFilter) FilterType() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("filterType"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/hideMenuBar
func (s SLContentFilter) HideMenuBar() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hideMenuBar"))
	return rv
}
func (s SLContentFilter) SetHideMenuBar(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHideMenuBar:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/includedApplications
func (s SLContentFilter) IncludedApplications() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("includedApplications"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/includedPIDS
func (s SLContentFilter) IncludedPIDS() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("includedPIDS"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLContentFilter) SetIncludedPIDS(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncludedPIDS:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/includedWindows
func (s SLContentFilter) IncludedWindows() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("includedWindows"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/shareAll
func (s SLContentFilter) ShareAll() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shareAll"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLContentFilter/windowID
func (s SLContentFilter) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
