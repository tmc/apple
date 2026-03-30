// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PBCodable] class.
var (
	_PBCodableClass     PBCodableClass
	_PBCodableClassOnce sync.Once
)

func getPBCodableClass() PBCodableClass {
	_PBCodableClassOnce.Do(func() {
		_PBCodableClass = PBCodableClass{class: objc.GetClass("PBCodable")}
	})
	return _PBCodableClass
}

// GetPBCodableClass returns the class object for PBCodable.
func GetPBCodableClass() PBCodableClass {
	return getPBCodableClass()
}

type PBCodableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PBCodableClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PBCodableClass) Alloc() PBCodable {
	rv := objc.Send[PBCodable](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other CoreML classes. [Full Topic]
type PBCodable struct {
	objectivec.Object
}

// PBCodableFromID constructs a [PBCodable] from an objc.ID.
//
// A parent class referenced by other CoreML classes.
func PBCodableFromID(id objc.ID) PBCodable {
	return PBCodable{objectivec.Object{ID: id}}
}

// Ensure PBCodable implements IPBCodable.
var _ IPBCodable = PBCodable{}

// An interface definition for the [PBCodable] class.
type IPBCodable interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p PBCodable) Init() PBCodable {
	rv := objc.Send[PBCodable](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PBCodable) Autorelease() PBCodable {
	rv := objc.Send[PBCodable](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPBCodable creates a new PBCodable instance.
func NewPBCodable() PBCodable {
	class := getPBCodableClass()
	rv := objc.Send[PBCodable](objc.ID(class.class), objc.Sel("new"))
	return rv
}
