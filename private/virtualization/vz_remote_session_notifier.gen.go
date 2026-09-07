// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZRemoteSessionNotifier] class.
var (
	_VZRemoteSessionNotifierClass     VZRemoteSessionNotifierClass
	_VZRemoteSessionNotifierClassOnce sync.Once
)

func getVZRemoteSessionNotifierClass() VZRemoteSessionNotifierClass {
	_VZRemoteSessionNotifierClassOnce.Do(func() {
		_VZRemoteSessionNotifierClass = VZRemoteSessionNotifierClass{class: objc.GetClass("_VZRemoteSessionNotifier")}
	})
	return _VZRemoteSessionNotifierClass
}

// GetVZRemoteSessionNotifierClass returns the class object for _VZRemoteSessionNotifier.
func GetVZRemoteSessionNotifierClass() VZRemoteSessionNotifierClass {
	return getVZRemoteSessionNotifierClass()
}

type VZRemoteSessionNotifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZRemoteSessionNotifierClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZRemoteSessionNotifierClass) Alloc() VZRemoteSessionNotifier {
	rv := objc.SendIfResponds[VZRemoteSessionNotifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZRemoteSessionNotifier struct {
	objectivec.Object
}

// VZRemoteSessionNotifierFromID constructs a [VZRemoteSessionNotifier] from an objc.ID.
func VZRemoteSessionNotifierFromID(id objc.ID) VZRemoteSessionNotifier {
	return VZRemoteSessionNotifier{objectivec.Object{ID: id}}
}

// Ensure VZRemoteSessionNotifier implements IVZRemoteSessionNotifier.
var _ IVZRemoteSessionNotifier = VZRemoteSessionNotifier{}

// An interface definition for the [VZRemoteSessionNotifier] class.
type IVZRemoteSessionNotifier interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZRemoteSessionNotifier) Init() VZRemoteSessionNotifier {
	rv := objc.SendIfResponds[VZRemoteSessionNotifier](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZRemoteSessionNotifier) Autorelease() VZRemoteSessionNotifier {
	rv := objc.SendIfResponds[VZRemoteSessionNotifier](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZRemoteSessionNotifier creates a new VZRemoteSessionNotifier instance.
func NewVZRemoteSessionNotifier() VZRemoteSessionNotifier {
	class := getVZRemoteSessionNotifierClass()
	rv := objc.SendIfResponds[VZRemoteSessionNotifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}
