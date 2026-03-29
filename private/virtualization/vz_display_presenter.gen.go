// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDisplayPresenter] class.
var (
	_VZDisplayPresenterClass     VZDisplayPresenterClass
	_VZDisplayPresenterClassOnce sync.Once
)

func getVZDisplayPresenterClass() VZDisplayPresenterClass {
	_VZDisplayPresenterClassOnce.Do(func() {
		_VZDisplayPresenterClass = VZDisplayPresenterClass{class: objc.GetClass("_VZDisplayPresenter")}
	})
	return _VZDisplayPresenterClass
}

// GetVZDisplayPresenterClass returns the class object for _VZDisplayPresenter.
func GetVZDisplayPresenterClass() VZDisplayPresenterClass {
	return getVZDisplayPresenterClass()
}

type VZDisplayPresenterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDisplayPresenterClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDisplayPresenterClass) Alloc() VZDisplayPresenter {
	rv := objc.Send[VZDisplayPresenter](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDisplayPresenter
type VZDisplayPresenter struct {
	objectivec.Object
}

// VZDisplayPresenterFromID constructs a [VZDisplayPresenter] from an objc.ID.
func VZDisplayPresenterFromID(id objc.ID) VZDisplayPresenter {
	return VZDisplayPresenter{objectivec.Object{ID: id}}
}
// Ensure VZDisplayPresenter implements IVZDisplayPresenter.
var _ IVZDisplayPresenter = VZDisplayPresenter{}

// An interface definition for the [VZDisplayPresenter] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDisplayPresenter
type IVZDisplayPresenter interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZDisplayPresenter) Init() VZDisplayPresenter {
	rv := objc.Send[VZDisplayPresenter](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDisplayPresenter) Autorelease() VZDisplayPresenter {
	rv := objc.Send[VZDisplayPresenter](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDisplayPresenter creates a new VZDisplayPresenter instance.
func NewVZDisplayPresenter() VZDisplayPresenter {
	class := getVZDisplayPresenterClass()
	rv := objc.Send[VZDisplayPresenter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

