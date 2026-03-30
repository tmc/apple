// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SymbolDescriptor] class.
var (
	_SymbolDescriptorClass     SymbolDescriptorClass
	_SymbolDescriptorClassOnce sync.Once
)

func getSymbolDescriptorClass() SymbolDescriptorClass {
	_SymbolDescriptorClassOnce.Do(func() {
		_SymbolDescriptorClass = SymbolDescriptorClass{class: objc.GetClass("symbolDescriptor")}
	})
	return _SymbolDescriptorClass
}

// GetSymbolDescriptorClass returns the class object for symbolDescriptor.
func GetSymbolDescriptorClass() SymbolDescriptorClass {
	return getSymbolDescriptorClass()
}

type SymbolDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SymbolDescriptorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SymbolDescriptorClass) Alloc() SymbolDescriptor {
	rv := objc.Send[SymbolDescriptor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/symbolDescriptor-c.ivar
type SymbolDescriptor struct {
	objectivec.Object
}

// SymbolDescriptorFromID constructs a [SymbolDescriptor] from an objc.ID.
func SymbolDescriptorFromID(id objc.ID) SymbolDescriptor {
	return SymbolDescriptor{objectivec.Object{ID: id}}
}

// Ensure SymbolDescriptor implements ISymbolDescriptor.
var _ ISymbolDescriptor = SymbolDescriptor{}

// An interface definition for the [SymbolDescriptor] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/symbolDescriptor-c.ivar
type ISymbolDescriptor interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SymbolDescriptor) Init() SymbolDescriptor {
	rv := objc.Send[SymbolDescriptor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SymbolDescriptor) Autorelease() SymbolDescriptor {
	rv := objc.Send[SymbolDescriptor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSymbolDescriptor creates a new SymbolDescriptor instance.
func NewSymbolDescriptor() SymbolDescriptor {
	class := getSymbolDescriptorClass()
	rv := objc.Send[SymbolDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
