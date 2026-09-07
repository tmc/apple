// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSLocalDefaults] class.
var (
	_WSLocalDefaultsClass     WSLocalDefaultsClass
	_WSLocalDefaultsClassOnce sync.Once
)

func getWSLocalDefaultsClass() WSLocalDefaultsClass {
	_WSLocalDefaultsClassOnce.Do(func() {
		_WSLocalDefaultsClass = WSLocalDefaultsClass{class: objc.GetClass("WSLocalDefaults")}
	})
	return _WSLocalDefaultsClass
}

// GetWSLocalDefaultsClass returns the class object for WSLocalDefaults.
func GetWSLocalDefaultsClass() WSLocalDefaultsClass {
	return getWSLocalDefaultsClass()
}

type WSLocalDefaultsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSLocalDefaultsClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSLocalDefaultsClass) Alloc() WSLocalDefaults {
	rv := objc.SendIfResponds[WSLocalDefaults](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSLocalDefaults._init]
//   - [WSLocalDefaults.SystemGestures]
type WSLocalDefaults struct {
	objectivec.Object
}

// WSLocalDefaultsFromID constructs a [WSLocalDefaults] from an objc.ID.
func WSLocalDefaultsFromID(id objc.ID) WSLocalDefaults {
	return WSLocalDefaults{objectivec.Object{ID: id}}
}

// Ensure WSLocalDefaults implements IWSLocalDefaults.
var _ IWSLocalDefaults = WSLocalDefaults{}

// An interface definition for the [WSLocalDefaults] class.
//
// # Methods
//
//   - [IWSLocalDefaults._init]
//   - [IWSLocalDefaults.SystemGestures]
type IWSLocalDefaults interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	SystemGestures() IWSSystemGestureDefaults
}

// Init initializes the instance.
func (w WSLocalDefaults) Init() WSLocalDefaults {
	rv := objc.SendIfResponds[WSLocalDefaults](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSLocalDefaults) Autorelease() WSLocalDefaults {
	rv := objc.SendIfResponds[WSLocalDefaults](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSLocalDefaults creates a new WSLocalDefaults instance.
func NewWSLocalDefaults() WSLocalDefaults {
	class := getWSLocalDefaultsClass()
	rv := objc.SendIfResponds[WSLocalDefaults](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSLocalDefaults) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (_WSLocalDefaultsClass WSLocalDefaultsClass) SharedInstance() WSLocalDefaults {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_WSLocalDefaultsClass.class), objc.Sel("sharedInstance"))
	return WSLocalDefaultsFromID(rv)
}

func (w WSLocalDefaults) SystemGestures() IWSSystemGestureDefaults {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("systemGestures"))
	return WSSystemGestureDefaultsFromID(objc.ID(rv))
}
