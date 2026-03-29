// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZFramebufferRemoteSessionNotifier] class.
var (
	_VZFramebufferRemoteSessionNotifierClass     VZFramebufferRemoteSessionNotifierClass
	_VZFramebufferRemoteSessionNotifierClassOnce sync.Once
)

func getVZFramebufferRemoteSessionNotifierClass() VZFramebufferRemoteSessionNotifierClass {
	_VZFramebufferRemoteSessionNotifierClassOnce.Do(func() {
		_VZFramebufferRemoteSessionNotifierClass = VZFramebufferRemoteSessionNotifierClass{class: objc.GetClass("_VZFramebufferRemoteSessionNotifier")}
	})
	return _VZFramebufferRemoteSessionNotifierClass
}

// GetVZFramebufferRemoteSessionNotifierClass returns the class object for _VZFramebufferRemoteSessionNotifier.
func GetVZFramebufferRemoteSessionNotifierClass() VZFramebufferRemoteSessionNotifierClass {
	return getVZFramebufferRemoteSessionNotifierClass()
}

type VZFramebufferRemoteSessionNotifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZFramebufferRemoteSessionNotifierClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFramebufferRemoteSessionNotifierClass) Alloc() VZFramebufferRemoteSessionNotifier {
	rv := objc.Send[VZFramebufferRemoteSessionNotifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferRemoteSessionNotifier
type VZFramebufferRemoteSessionNotifier struct {
	objectivec.Object
}

// VZFramebufferRemoteSessionNotifierFromID constructs a [VZFramebufferRemoteSessionNotifier] from an objc.ID.
func VZFramebufferRemoteSessionNotifierFromID(id objc.ID) VZFramebufferRemoteSessionNotifier {
	return VZFramebufferRemoteSessionNotifier{objectivec.Object{ID: id}}
}
// Ensure VZFramebufferRemoteSessionNotifier implements IVZFramebufferRemoteSessionNotifier.
var _ IVZFramebufferRemoteSessionNotifier = VZFramebufferRemoteSessionNotifier{}

// An interface definition for the [VZFramebufferRemoteSessionNotifier] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferRemoteSessionNotifier
type IVZFramebufferRemoteSessionNotifier interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v VZFramebufferRemoteSessionNotifier) Init() VZFramebufferRemoteSessionNotifier {
	rv := objc.Send[VZFramebufferRemoteSessionNotifier](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFramebufferRemoteSessionNotifier) Autorelease() VZFramebufferRemoteSessionNotifier {
	rv := objc.Send[VZFramebufferRemoteSessionNotifier](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFramebufferRemoteSessionNotifier creates a new VZFramebufferRemoteSessionNotifier instance.
func NewVZFramebufferRemoteSessionNotifier() VZFramebufferRemoteSessionNotifier {
	class := getVZFramebufferRemoteSessionNotifierClass()
	rv := objc.Send[VZFramebufferRemoteSessionNotifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

