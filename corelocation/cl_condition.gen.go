// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLCondition] class.
var (
	_CLConditionClass     CLConditionClass
	_CLConditionClassOnce sync.Once
)

func getCLConditionClass() CLConditionClass {
	_CLConditionClassOnce.Do(func() {
		_CLConditionClass = CLConditionClass{class: objc.GetClass("CLCondition")}
	})
	return _CLConditionClass
}

// GetCLConditionClass returns the class object for CLCondition.
func GetCLConditionClass() CLConditionClass {
	return getCLConditionClass()
}

type CLConditionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLConditionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLConditionClass) Alloc() CLCondition {
	rv := objc.Send[CLCondition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class that all other conditions derive from.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCondition-c.class
type CLCondition struct {
	objectivec.Object
}

// CLConditionFromID constructs a [CLCondition] from an objc.ID.
//
// The abstract base class that all other conditions derive from.
func CLConditionFromID(id objc.ID) CLCondition {
	return CLCondition{objectivec.Object{ID: id}}
}

// NOTE: CLCondition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLCondition] class.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCondition-c.class
type ICLCondition interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CLCondition) Init() CLCondition {
	rv := objc.Send[CLCondition](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CLCondition) Autorelease() CLCondition {
	rv := objc.Send[CLCondition](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLCondition creates a new CLCondition instance.
func NewCLCondition() CLCondition {
	class := getCLConditionClass()
	rv := objc.Send[CLCondition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CLCondition) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
