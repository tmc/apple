// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Components] class.
var (
	_ComponentsClass     ComponentsClass
	_ComponentsClassOnce sync.Once
)

func getComponentsClass() ComponentsClass {
	_ComponentsClassOnce.Do(func() {
		_ComponentsClass = ComponentsClass{class: objc.GetClass("components")}
	})
	return _ComponentsClass
}

// GetComponentsClass returns the class object for components.
func GetComponentsClass() ComponentsClass {
	return getComponentsClass()
}

type ComponentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc ComponentsClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc ComponentsClass) Alloc() Components {
	rv := objc.Send[Components](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/components-c.ivar
type Components struct {
	objectivec.Object
}

// ComponentsFromID constructs a [Components] from an objc.ID.
func ComponentsFromID(id objc.ID) Components {
	return Components{objectivec.Object{ID: id}}
}

// Ensure Components implements IComponents.
var _ IComponents = Components{}

// An interface definition for the [Components] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/components-c.ivar
type IComponents interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c Components) Init() Components {
	rv := objc.Send[Components](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c Components) Autorelease() Components {
	rv := objc.Send[Components](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewComponents creates a new Components instance.
func NewComponents() Components {
	class := getComponentsClass()
	rv := objc.Send[Components](objc.ID(class.class), objc.Sel("new"))
	return rv
}
