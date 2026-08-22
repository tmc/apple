// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXRemoteViewEventServiceSpecification] class.
var (
	_CPXRemoteViewEventServiceSpecificationClass     CPXRemoteViewEventServiceSpecificationClass
	_CPXRemoteViewEventServiceSpecificationClassOnce sync.Once
)

func getCPXRemoteViewEventServiceSpecificationClass() CPXRemoteViewEventServiceSpecificationClass {
	_CPXRemoteViewEventServiceSpecificationClassOnce.Do(func() {
		_CPXRemoteViewEventServiceSpecificationClass = CPXRemoteViewEventServiceSpecificationClass{class: objc.GetClass("CPXRemoteViewEventServiceSpecification")}
	})
	return _CPXRemoteViewEventServiceSpecificationClass
}

// GetCPXRemoteViewEventServiceSpecificationClass returns the class object for CPXRemoteViewEventServiceSpecification.
func GetCPXRemoteViewEventServiceSpecificationClass() CPXRemoteViewEventServiceSpecificationClass {
	return getCPXRemoteViewEventServiceSpecificationClass()
}

type CPXRemoteViewEventServiceSpecificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXRemoteViewEventServiceSpecificationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXRemoteViewEventServiceSpecificationClass) Alloc() CPXRemoteViewEventServiceSpecification {
	rv := objc.SendIfResponds[CPXRemoteViewEventServiceSpecification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

type CPXRemoteViewEventServiceSpecification struct {
	objectivec.Object
}

// CPXRemoteViewEventServiceSpecificationFromID constructs a [CPXRemoteViewEventServiceSpecification] from an objc.ID.
func CPXRemoteViewEventServiceSpecificationFromID(id objc.ID) CPXRemoteViewEventServiceSpecification {
	return CPXRemoteViewEventServiceSpecification{objectivec.Object{ID: id}}
}

// Ensure CPXRemoteViewEventServiceSpecification implements ICPXRemoteViewEventServiceSpecification.
var _ ICPXRemoteViewEventServiceSpecification = CPXRemoteViewEventServiceSpecification{}

// An interface definition for the [CPXRemoteViewEventServiceSpecification] class.
type ICPXRemoteViewEventServiceSpecification interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CPXRemoteViewEventServiceSpecification) Init() CPXRemoteViewEventServiceSpecification {
	rv := objc.SendIfResponds[CPXRemoteViewEventServiceSpecification](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXRemoteViewEventServiceSpecification) Autorelease() CPXRemoteViewEventServiceSpecification {
	rv := objc.SendIfResponds[CPXRemoteViewEventServiceSpecification](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXRemoteViewEventServiceSpecification creates a new CPXRemoteViewEventServiceSpecification instance.
func NewCPXRemoteViewEventServiceSpecification() CPXRemoteViewEventServiceSpecification {
	class := getCPXRemoteViewEventServiceSpecificationClass()
	rv := objc.SendIfResponds[CPXRemoteViewEventServiceSpecification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_CPXRemoteViewEventServiceSpecificationClass CPXRemoteViewEventServiceSpecificationClass) Identifier() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_CPXRemoteViewEventServiceSpecificationClass.class), objc.Sel("identifier"))
	return objectivec.Object{ID: rv}
}
func (_CPXRemoteViewEventServiceSpecificationClass CPXRemoteViewEventServiceSpecificationClass) Interface() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_CPXRemoteViewEventServiceSpecificationClass.class), objc.Sel("interface"))
	return objectivec.Object{ID: rv}
}
func (_CPXRemoteViewEventServiceSpecificationClass CPXRemoteViewEventServiceSpecificationClass) ServiceQuality() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_CPXRemoteViewEventServiceSpecificationClass.class), objc.Sel("serviceQuality"))
	return objectivec.Object{ID: rv}
}
