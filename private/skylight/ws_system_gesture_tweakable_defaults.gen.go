// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSSystemGestureTweakableDefaults] class.
var (
	_WSSystemGestureTweakableDefaultsClass     WSSystemGestureTweakableDefaultsClass
	_WSSystemGestureTweakableDefaultsClassOnce sync.Once
)

func getWSSystemGestureTweakableDefaultsClass() WSSystemGestureTweakableDefaultsClass {
	_WSSystemGestureTweakableDefaultsClassOnce.Do(func() {
		_WSSystemGestureTweakableDefaultsClass = WSSystemGestureTweakableDefaultsClass{class: objc.GetClass("WSSystemGestureTweakableDefaults")}
	})
	return _WSSystemGestureTweakableDefaultsClass
}

// GetWSSystemGestureTweakableDefaultsClass returns the class object for WSSystemGestureTweakableDefaults.
func GetWSSystemGestureTweakableDefaultsClass() WSSystemGestureTweakableDefaultsClass {
	return getWSSystemGestureTweakableDefaultsClass()
}

type WSSystemGestureTweakableDefaultsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSystemGestureTweakableDefaultsClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSystemGestureTweakableDefaultsClass) Alloc() WSSystemGestureTweakableDefaults {
	rv := objc.SendIfResponds[WSSystemGestureTweakableDefaults](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSSystemGestureTweakableDefaults._bindAndRegisterDefaults]
//   - [WSSystemGestureTweakableDefaults.PrimaryGateHysteresis]
//   - [WSSystemGestureTweakableDefaults.SetPrimaryGateHysteresis]
//   - [WSSystemGestureTweakableDefaults.PrimaryGatePointCountLimit]
//   - [WSSystemGestureTweakableDefaults.SetPrimaryGatePointCountLimit]
//   - [WSSystemGestureTweakableDefaults.PrimaryGateShouldCancelEvents]
//   - [WSSystemGestureTweakableDefaults.SetPrimaryGateShouldCancelEvents]
//   - [WSSystemGestureTweakableDefaults.PrimaryGateTimeoutMilliseconds]
//   - [WSSystemGestureTweakableDefaults.SetPrimaryGateTimeoutMilliseconds]
//   - [WSSystemGestureTweakableDefaults.SecondaryGateEnabled]
//   - [WSSystemGestureTweakableDefaults.SetSecondaryGateEnabled]
//   - [WSSystemGestureTweakableDefaults.SecondaryGateHysteresis]
//   - [WSSystemGestureTweakableDefaults.SetSecondaryGateHysteresis]
//   - [WSSystemGestureTweakableDefaults.SecondaryGatePointCountLimit]
//   - [WSSystemGestureTweakableDefaults.SetSecondaryGatePointCountLimit]
//   - [WSSystemGestureTweakableDefaults.SecondaryGateShouldCancelEvents]
//   - [WSSystemGestureTweakableDefaults.SetSecondaryGateShouldCancelEvents]
//   - [WSSystemGestureTweakableDefaults.SecondaryGateTimeoutMilliseconds]
//   - [WSSystemGestureTweakableDefaults.SetSecondaryGateTimeoutMilliseconds]
type WSSystemGestureTweakableDefaults struct {
	objectivec.Object
}

// WSSystemGestureTweakableDefaultsFromID constructs a [WSSystemGestureTweakableDefaults] from an objc.ID.
func WSSystemGestureTweakableDefaultsFromID(id objc.ID) WSSystemGestureTweakableDefaults {
	return WSSystemGestureTweakableDefaults{objectivec.Object{ID: id}}
}

// NOTE: WSSystemGestureTweakableDefaults embeds objectivec.Object because the parent type is
// unavailable, but IWSSystemGestureTweakableDefaults embeds IBSAbstractDefaultDomain, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [WSSystemGestureTweakableDefaults] class.
//
// # Methods
//
//   - [IWSSystemGestureTweakableDefaults._bindAndRegisterDefaults]
//   - [IWSSystemGestureTweakableDefaults.PrimaryGateHysteresis]
//   - [IWSSystemGestureTweakableDefaults.SetPrimaryGateHysteresis]
//   - [IWSSystemGestureTweakableDefaults.PrimaryGatePointCountLimit]
//   - [IWSSystemGestureTweakableDefaults.SetPrimaryGatePointCountLimit]
//   - [IWSSystemGestureTweakableDefaults.PrimaryGateShouldCancelEvents]
//   - [IWSSystemGestureTweakableDefaults.SetPrimaryGateShouldCancelEvents]
//   - [IWSSystemGestureTweakableDefaults.PrimaryGateTimeoutMilliseconds]
//   - [IWSSystemGestureTweakableDefaults.SetPrimaryGateTimeoutMilliseconds]
//   - [IWSSystemGestureTweakableDefaults.SecondaryGateEnabled]
//   - [IWSSystemGestureTweakableDefaults.SetSecondaryGateEnabled]
//   - [IWSSystemGestureTweakableDefaults.SecondaryGateHysteresis]
//   - [IWSSystemGestureTweakableDefaults.SetSecondaryGateHysteresis]
//   - [IWSSystemGestureTweakableDefaults.SecondaryGatePointCountLimit]
//   - [IWSSystemGestureTweakableDefaults.SetSecondaryGatePointCountLimit]
//   - [IWSSystemGestureTweakableDefaults.SecondaryGateShouldCancelEvents]
//   - [IWSSystemGestureTweakableDefaults.SetSecondaryGateShouldCancelEvents]
//   - [IWSSystemGestureTweakableDefaults.SecondaryGateTimeoutMilliseconds]
//   - [IWSSystemGestureTweakableDefaults.SetSecondaryGateTimeoutMilliseconds]
type IWSSystemGestureTweakableDefaults interface {
	IBSAbstractDefaultDomain

	// Topic: Methods

	_bindAndRegisterDefaults()
	PrimaryGateHysteresis() float64
	SetPrimaryGateHysteresis(value float64)
	PrimaryGatePointCountLimit() int64
	SetPrimaryGatePointCountLimit(value int64)
	PrimaryGateShouldCancelEvents() bool
	SetPrimaryGateShouldCancelEvents(value bool)
	PrimaryGateTimeoutMilliseconds() float64
	SetPrimaryGateTimeoutMilliseconds(value float64)
	SecondaryGateEnabled() bool
	SetSecondaryGateEnabled(value bool)
	SecondaryGateHysteresis() float64
	SetSecondaryGateHysteresis(value float64)
	SecondaryGatePointCountLimit() int64
	SetSecondaryGatePointCountLimit(value int64)
	SecondaryGateShouldCancelEvents() bool
	SetSecondaryGateShouldCancelEvents(value bool)
	SecondaryGateTimeoutMilliseconds() float64
	SetSecondaryGateTimeoutMilliseconds(value float64)
}

// Init initializes the instance.
func (w WSSystemGestureTweakableDefaults) Init() WSSystemGestureTweakableDefaults {
	rv := objc.SendIfResponds[WSSystemGestureTweakableDefaults](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSystemGestureTweakableDefaults) Autorelease() WSSystemGestureTweakableDefaults {
	rv := objc.SendIfResponds[WSSystemGestureTweakableDefaults](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSystemGestureTweakableDefaults creates a new WSSystemGestureTweakableDefaults instance.
func NewWSSystemGestureTweakableDefaults() WSSystemGestureTweakableDefaults {
	class := getWSSystemGestureTweakableDefaultsClass()
	rv := objc.SendIfResponds[WSSystemGestureTweakableDefaults](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSSystemGestureTweakableDefaults) _bindAndRegisterDefaults() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_bindAndRegisterDefaults"))
}

// BindAndRegisterDefaults is an exported wrapper for the private method _bindAndRegisterDefaults.
func (w WSSystemGestureTweakableDefaults) BindAndRegisterDefaults() error {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_bindAndRegisterDefaults")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_bindAndRegisterDefaults"}
		return err
	}
	w._bindAndRegisterDefaults()
	return nil
}

// CanBindAndRegisterDefaults reports whether the receiver responds to the private selector _bindAndRegisterDefaults.
func (w WSSystemGestureTweakableDefaults) CanBindAndRegisterDefaults() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_bindAndRegisterDefaults"))
}

func (w WSSystemGestureTweakableDefaults) PrimaryGateHysteresis() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("primaryGateHysteresis"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetPrimaryGateHysteresis(value float64) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setPrimaryGateHysteresis:"), value)
}
func (w WSSystemGestureTweakableDefaults) PrimaryGatePointCountLimit() int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("primaryGatePointCountLimit"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetPrimaryGatePointCountLimit(value int64) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setPrimaryGatePointCountLimit:"), value)
}
func (w WSSystemGestureTweakableDefaults) PrimaryGateShouldCancelEvents() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("primaryGateShouldCancelEvents"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetPrimaryGateShouldCancelEvents(value bool) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setPrimaryGateShouldCancelEvents:"), value)
}
func (w WSSystemGestureTweakableDefaults) PrimaryGateTimeoutMilliseconds() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("primaryGateTimeoutMilliseconds"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetPrimaryGateTimeoutMilliseconds(value float64) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setPrimaryGateTimeoutMilliseconds:"), value)
}
func (w WSSystemGestureTweakableDefaults) SecondaryGateEnabled() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("secondaryGateEnabled"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetSecondaryGateEnabled(value bool) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setSecondaryGateEnabled:"), value)
}
func (w WSSystemGestureTweakableDefaults) SecondaryGateHysteresis() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("secondaryGateHysteresis"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetSecondaryGateHysteresis(value float64) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setSecondaryGateHysteresis:"), value)
}
func (w WSSystemGestureTweakableDefaults) SecondaryGatePointCountLimit() int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("secondaryGatePointCountLimit"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetSecondaryGatePointCountLimit(value int64) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setSecondaryGatePointCountLimit:"), value)
}
func (w WSSystemGestureTweakableDefaults) SecondaryGateShouldCancelEvents() bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("secondaryGateShouldCancelEvents"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetSecondaryGateShouldCancelEvents(value bool) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setSecondaryGateShouldCancelEvents:"), value)
}
func (w WSSystemGestureTweakableDefaults) SecondaryGateTimeoutMilliseconds() float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("secondaryGateTimeoutMilliseconds"))
	return rv
}
func (w WSSystemGestureTweakableDefaults) SetSecondaryGateTimeoutMilliseconds(value float64) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setSecondaryGateTimeoutMilliseconds:"), value)
}
