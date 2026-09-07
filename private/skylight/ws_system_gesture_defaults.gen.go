// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSSystemGestureDefaults] class.
var (
	_WSSystemGestureDefaultsClass     WSSystemGestureDefaultsClass
	_WSSystemGestureDefaultsClassOnce sync.Once
)

func getWSSystemGestureDefaultsClass() WSSystemGestureDefaultsClass {
	_WSSystemGestureDefaultsClassOnce.Do(func() {
		_WSSystemGestureDefaultsClass = WSSystemGestureDefaultsClass{class: objc.GetClass("WSSystemGestureDefaults")}
	})
	return _WSSystemGestureDefaultsClass
}

// GetWSSystemGestureDefaultsClass returns the class object for WSSystemGestureDefaults.
func GetWSSystemGestureDefaultsClass() WSSystemGestureDefaultsClass {
	return getWSSystemGestureDefaultsClass()
}

type WSSystemGestureDefaultsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSystemGestureDefaultsClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSystemGestureDefaultsClass) Alloc() WSSystemGestureDefaults {
	rv := objc.SendIfResponds[WSSystemGestureDefaults](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSSystemGestureDefaults.PrimaryGateHysteresis]
//   - [WSSystemGestureDefaults.PrimaryGatePointCountLimit]
//   - [WSSystemGestureDefaults.PrimaryGateShouldCancelEvents]
//   - [WSSystemGestureDefaults.PrimaryGateTimeoutMilliseconds]
//   - [WSSystemGestureDefaults.SecondaryGateEnabled]
//   - [WSSystemGestureDefaults.SecondaryGateHysteresis]
//   - [WSSystemGestureDefaults.SecondaryGatePointCountLimit]
//   - [WSSystemGestureDefaults.SecondaryGateShouldCancelEvents]
//   - [WSSystemGestureDefaults.SecondaryGateTimeoutMilliseconds]
type WSSystemGestureDefaults struct {
	objectivec.Object
}

// WSSystemGestureDefaultsFromID constructs a [WSSystemGestureDefaults] from an objc.ID.
func WSSystemGestureDefaultsFromID(id objc.ID) WSSystemGestureDefaults {
	return WSSystemGestureDefaults{objectivec.Object{ID: id}}
}

// Ensure WSSystemGestureDefaults implements IWSSystemGestureDefaults.
var _ IWSSystemGestureDefaults = WSSystemGestureDefaults{}

// An interface definition for the [WSSystemGestureDefaults] class.
//
// # Methods
//
//   - [IWSSystemGestureDefaults.PrimaryGateHysteresis]
//   - [IWSSystemGestureDefaults.PrimaryGatePointCountLimit]
//   - [IWSSystemGestureDefaults.PrimaryGateShouldCancelEvents]
//   - [IWSSystemGestureDefaults.PrimaryGateTimeoutMilliseconds]
//   - [IWSSystemGestureDefaults.SecondaryGateEnabled]
//   - [IWSSystemGestureDefaults.SecondaryGateHysteresis]
//   - [IWSSystemGestureDefaults.SecondaryGatePointCountLimit]
//   - [IWSSystemGestureDefaults.SecondaryGateShouldCancelEvents]
//   - [IWSSystemGestureDefaults.SecondaryGateTimeoutMilliseconds]
type IWSSystemGestureDefaults interface {
	objectivec.IObject

	// Topic: Methods

	PrimaryGateHysteresis() float64
	PrimaryGatePointCountLimit() int64
	PrimaryGateShouldCancelEvents() bool
	PrimaryGateTimeoutMilliseconds() float64
	SecondaryGateEnabled() bool
	SecondaryGateHysteresis() float64
	SecondaryGatePointCountLimit() int64
	SecondaryGateShouldCancelEvents() bool
	SecondaryGateTimeoutMilliseconds() float64
}

// Init initializes the instance.
func (w WSSystemGestureDefaults) Init() WSSystemGestureDefaults {
	rv := objc.SendIfResponds[WSSystemGestureDefaults](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSystemGestureDefaults) Autorelease() WSSystemGestureDefaults {
	rv := objc.SendIfResponds[WSSystemGestureDefaults](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSystemGestureDefaults creates a new WSSystemGestureDefaults instance.
func NewWSSystemGestureDefaults() WSSystemGestureDefaults {
	class := getWSSystemGestureDefaultsClass()
	rv := objc.SendIfResponds[WSSystemGestureDefaults](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSSystemGestureDefaults) PrimaryGateHysteresis() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("primaryGateHysteresis"))
	return rv
}
func (w WSSystemGestureDefaults) PrimaryGatePointCountLimit() int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("primaryGatePointCountLimit"))
	return rv
}
func (w WSSystemGestureDefaults) PrimaryGateShouldCancelEvents() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("primaryGateShouldCancelEvents"))
	return rv
}
func (w WSSystemGestureDefaults) PrimaryGateTimeoutMilliseconds() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("primaryGateTimeoutMilliseconds"))
	return rv
}
func (w WSSystemGestureDefaults) SecondaryGateEnabled() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("secondaryGateEnabled"))
	return rv
}
func (w WSSystemGestureDefaults) SecondaryGateHysteresis() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("secondaryGateHysteresis"))
	return rv
}
func (w WSSystemGestureDefaults) SecondaryGatePointCountLimit() int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("secondaryGatePointCountLimit"))
	return rv
}
func (w WSSystemGestureDefaults) SecondaryGateShouldCancelEvents() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("secondaryGateShouldCancelEvents"))
	return rv
}
func (w WSSystemGestureDefaults) SecondaryGateTimeoutMilliseconds() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("secondaryGateTimeoutMilliseconds"))
	return rv
}
