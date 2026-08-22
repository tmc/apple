// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLObjectWithLabel] class.
var (
	_MTLObjectWithLabelClass     MTLObjectWithLabelClass
	_MTLObjectWithLabelClassOnce sync.Once
)

func getMTLObjectWithLabelClass() MTLObjectWithLabelClass {
	_MTLObjectWithLabelClassOnce.Do(func() {
		_MTLObjectWithLabelClass = MTLObjectWithLabelClass{class: objc.GetClass("_MTLObjectWithLabel")}
	})
	return _MTLObjectWithLabelClass
}

// GetMTLObjectWithLabelClass returns the class object for _MTLObjectWithLabel.
func GetMTLObjectWithLabelClass() MTLObjectWithLabelClass {
	return getMTLObjectWithLabelClass()
}

type MTLObjectWithLabelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLObjectWithLabelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLObjectWithLabelClass) Alloc() MTLObjectWithLabel {
	rv := objc.SendIfResponds[MTLObjectWithLabel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other iogpu classes. [Full Topic]
type MTLObjectWithLabel struct {
	objectivec.Object
}

// MTLObjectWithLabelFromID constructs a [MTLObjectWithLabel] from an objc.ID.
//
// A parent class referenced by other iogpu classes.
func MTLObjectWithLabelFromID(id objc.ID) MTLObjectWithLabel {
	return MTLObjectWithLabel{objectivec.Object{ID: id}}
}

// Ensure MTLObjectWithLabel implements IMTLObjectWithLabel.
var _ IMTLObjectWithLabel = MTLObjectWithLabel{}

// An interface definition for the [MTLObjectWithLabel] class.
type IMTLObjectWithLabel interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MTLObjectWithLabel) Init() MTLObjectWithLabel {
	rv := objc.SendIfResponds[MTLObjectWithLabel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLObjectWithLabel) Autorelease() MTLObjectWithLabel {
	rv := objc.SendIfResponds[MTLObjectWithLabel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLObjectWithLabel creates a new MTLObjectWithLabel instance.
func NewMTLObjectWithLabel() MTLObjectWithLabel {
	class := getMTLObjectWithLabelClass()
	rv := objc.SendIfResponds[MTLObjectWithLabel](objc.ID(class.class), objc.Sel("new"))
	return rv
}
