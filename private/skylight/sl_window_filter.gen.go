// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLWindowFilter] class.
var (
	_SLWindowFilterClass     SLWindowFilterClass
	_SLWindowFilterClassOnce sync.Once
)

func getSLWindowFilterClass() SLWindowFilterClass {
	_SLWindowFilterClassOnce.Do(func() {
		_SLWindowFilterClass = SLWindowFilterClass{class: objc.GetClass("SLWindowFilter")}
	})
	return _SLWindowFilterClass
}

// GetSLWindowFilterClass returns the class object for SLWindowFilter.
func GetSLWindowFilterClass() SLWindowFilterClass {
	return getSLWindowFilterClass()
}

type SLWindowFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLWindowFilterClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLWindowFilterClass) Alloc() SLWindowFilter {
	rv := objc.Send[SLWindowFilter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLWindowFilter.DictionaryRepresentation]
//   - [SLWindowFilter.Enforce_sharing_type]
//   - [SLWindowFilter.SetEnforce_sharing_type]
//   - [SLWindowFilter.ExcludedApplications]
//   - [SLWindowFilter.SetExcludedApplications]
//   - [SLWindowFilter.ExcludedPIDS]
//   - [SLWindowFilter.SetExcludedPIDS]
//   - [SLWindowFilter.ExcludedWindows]
//   - [SLWindowFilter.SetExcludedWindows]
//   - [SLWindowFilter.FilterPolicy]
//   - [SLWindowFilter.SetFilterPolicy]
//   - [SLWindowFilter.Hide_menu_bar]
//   - [SLWindowFilter.SetHide_menu_bar]
//   - [SLWindowFilter.IncludedApplications]
//   - [SLWindowFilter.SetIncludedApplications]
//   - [SLWindowFilter.IncludedPIDS]
//   - [SLWindowFilter.SetIncludedPIDS]
//   - [SLWindowFilter.IncludedWindows]
//   - [SLWindowFilter.SetIncludedWindows]
//   - [SLWindowFilter.IsEqualToWindowFilter]
//   - [SLWindowFilter.ShieldWindow]
//   - [SLWindowFilter.SetShieldWindow]
//   - [SLWindowFilter.InitFromDictionaryRepresentation]
//
// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter
type SLWindowFilter struct {
	objectivec.Object
}

// SLWindowFilterFromID constructs a [SLWindowFilter] from an objc.ID.
func SLWindowFilterFromID(id objc.ID) SLWindowFilter {
	return SLWindowFilter{objectivec.Object{ID: id}}
}

// Ensure SLWindowFilter implements ISLWindowFilter.
var _ ISLWindowFilter = SLWindowFilter{}

// An interface definition for the [SLWindowFilter] class.
//
// # Methods
//
//   - [ISLWindowFilter.DictionaryRepresentation]
//   - [ISLWindowFilter.Enforce_sharing_type]
//   - [ISLWindowFilter.SetEnforce_sharing_type]
//   - [ISLWindowFilter.ExcludedApplications]
//   - [ISLWindowFilter.SetExcludedApplications]
//   - [ISLWindowFilter.ExcludedPIDS]
//   - [ISLWindowFilter.SetExcludedPIDS]
//   - [ISLWindowFilter.ExcludedWindows]
//   - [ISLWindowFilter.SetExcludedWindows]
//   - [ISLWindowFilter.FilterPolicy]
//   - [ISLWindowFilter.SetFilterPolicy]
//   - [ISLWindowFilter.Hide_menu_bar]
//   - [ISLWindowFilter.SetHide_menu_bar]
//   - [ISLWindowFilter.IncludedApplications]
//   - [ISLWindowFilter.SetIncludedApplications]
//   - [ISLWindowFilter.IncludedPIDS]
//   - [ISLWindowFilter.SetIncludedPIDS]
//   - [ISLWindowFilter.IncludedWindows]
//   - [ISLWindowFilter.SetIncludedWindows]
//   - [ISLWindowFilter.IsEqualToWindowFilter]
//   - [ISLWindowFilter.ShieldWindow]
//   - [ISLWindowFilter.SetShieldWindow]
//   - [ISLWindowFilter.InitFromDictionaryRepresentation]
//
// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter
type ISLWindowFilter interface {
	objectivec.IObject

	// Topic: Methods

	DictionaryRepresentation() objectivec.IObject
	Enforce_sharing_type() bool
	SetEnforce_sharing_type(value bool)
	ExcludedApplications() foundation.INSSet
	SetExcludedApplications(value foundation.INSSet)
	ExcludedPIDS() foundation.INSSet
	SetExcludedPIDS(value foundation.INSSet)
	ExcludedWindows() foundation.INSSet
	SetExcludedWindows(value foundation.INSSet)
	FilterPolicy() foundation.NSNumber
	SetFilterPolicy(value foundation.NSNumber)
	Hide_menu_bar() bool
	SetHide_menu_bar(value bool)
	IncludedApplications() foundation.INSSet
	SetIncludedApplications(value foundation.INSSet)
	IncludedPIDS() foundation.INSSet
	SetIncludedPIDS(value foundation.INSSet)
	IncludedWindows() foundation.INSSet
	SetIncludedWindows(value foundation.INSSet)
	IsEqualToWindowFilter(filter objectivec.IObject) bool
	ShieldWindow() foundation.NSNumber
	SetShieldWindow(value foundation.NSNumber)
	InitFromDictionaryRepresentation(representation objectivec.IObject) SLWindowFilter
}

// Init initializes the instance.
func (s SLWindowFilter) Init() SLWindowFilter {
	rv := objc.Send[SLWindowFilter](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLWindowFilter) Autorelease() SLWindowFilter {
	rv := objc.Send[SLWindowFilter](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLWindowFilter creates a new SLWindowFilter instance.
func NewSLWindowFilter() SLWindowFilter {
	class := getSLWindowFilterClass()
	rv := objc.Send[SLWindowFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/initFromDictionaryRepresentation:
func NewSLWindowFilterFromDictionaryRepresentation(representation objectivec.IObject) SLWindowFilter {
	instance := getSLWindowFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromDictionaryRepresentation:"), representation)
	return SLWindowFilterFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/dictionaryRepresentation
func (s SLWindowFilter) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/isEqualToWindowFilter:
func (s SLWindowFilter) IsEqualToWindowFilter(filter objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEqualToWindowFilter:"), filter)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/initFromDictionaryRepresentation:
func (s SLWindowFilter) InitFromDictionaryRepresentation(representation objectivec.IObject) SLWindowFilter {
	rv := objc.Send[SLWindowFilter](s.ID, objc.Sel("initFromDictionaryRepresentation:"), representation)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/enforce_sharing_type
func (s SLWindowFilter) Enforce_sharing_type() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("enforce_sharing_type"))
	return rv
}
func (s SLWindowFilter) SetEnforce_sharing_type(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnforce_sharing_type:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/excludedApplications
func (s SLWindowFilter) ExcludedApplications() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("excludedApplications"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetExcludedApplications(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setExcludedApplications:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/excludedPIDS
func (s SLWindowFilter) ExcludedPIDS() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("excludedPIDS"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetExcludedPIDS(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setExcludedPIDS:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/excludedWindows
func (s SLWindowFilter) ExcludedWindows() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("excludedWindows"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetExcludedWindows(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setExcludedWindows:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/filterPolicy
func (s SLWindowFilter) FilterPolicy() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filterPolicy"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetFilterPolicy(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setFilterPolicy:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/hide_menu_bar
func (s SLWindowFilter) Hide_menu_bar() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hide_menu_bar"))
	return rv
}
func (s SLWindowFilter) SetHide_menu_bar(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHide_menu_bar:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/includedApplications
func (s SLWindowFilter) IncludedApplications() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("includedApplications"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetIncludedApplications(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncludedApplications:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/includedPIDS
func (s SLWindowFilter) IncludedPIDS() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("includedPIDS"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetIncludedPIDS(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncludedPIDS:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/includedWindows
func (s SLWindowFilter) IncludedWindows() foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("includedWindows"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetIncludedWindows(value foundation.INSSet) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncludedWindows:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowFilter/shieldWindow
func (s SLWindowFilter) ShieldWindow() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("shieldWindow"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLWindowFilter) SetShieldWindow(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setShieldWindow:"), value)
}
