// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLSharedTextureHandle] class.
var (
	_MTLSharedTextureHandleClass     MTLSharedTextureHandleClass
	_MTLSharedTextureHandleClassOnce sync.Once
)

func getMTLSharedTextureHandleClass() MTLSharedTextureHandleClass {
	_MTLSharedTextureHandleClassOnce.Do(func() {
		_MTLSharedTextureHandleClass = MTLSharedTextureHandleClass{class: objc.GetClass("MTLSharedTextureHandle")}
	})
	return _MTLSharedTextureHandleClass
}

// GetMTLSharedTextureHandleClass returns the class object for MTLSharedTextureHandle.
func GetMTLSharedTextureHandleClass() MTLSharedTextureHandleClass {
	return getMTLSharedTextureHandleClass()
}

type MTLSharedTextureHandleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLSharedTextureHandleClass) Alloc() MTLSharedTextureHandle {
	rv := objc.Send[MTLSharedTextureHandle](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A texture handle that can be shared across process address space
// boundaries.
//
// # Overview
// 
// [MTLSharedTextureHandle] objects may be passed between processes using XPC
// connections and then used to create a reference to the texture in another
// process. The texture in the other process needs to be created using the
// same [MTLDevice] on which the shared texture was originally created. To
// identify which device it was created on, you can use the [MTLSharedTextureHandle.Device] property
// of the [MTLSharedTextureHandle] object.
//
// # Identifying the shared texture handle
//
//   - [MTLSharedTextureHandle.Device]: The device object that created the texture.
//   - [MTLSharedTextureHandle.Label]: A string that identifies the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedTextureHandle
type MTLSharedTextureHandle struct {
	objectivec.Object
}

// MTLSharedTextureHandleFromID constructs a [MTLSharedTextureHandle] from an objc.ID.
//
// A texture handle that can be shared across process address space
// boundaries.
func MTLSharedTextureHandleFromID(id objc.ID) MTLSharedTextureHandle {
	return MTLSharedTextureHandle{objectivec.Object{id}}
}
// NOTE: MTLSharedTextureHandle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLSharedTextureHandle] class.
//
// # Identifying the shared texture handle
//
//   - [IMTLSharedTextureHandle.Device]: The device object that created the texture.
//   - [IMTLSharedTextureHandle.Label]: A string that identifies the texture.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedTextureHandle
type IMTLSharedTextureHandle interface {
	objectivec.IObject

	// Topic: Identifying the shared texture handle

	// The device object that created the texture.
	Device() MTLDevice
	// A string that identifies the texture.
	Label() string

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s MTLSharedTextureHandle) Init() MTLSharedTextureHandle {
	rv := objc.Send[MTLSharedTextureHandle](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLSharedTextureHandle) Autorelease() MTLSharedTextureHandle {
	rv := objc.Send[MTLSharedTextureHandle](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLSharedTextureHandle creates a new MTLSharedTextureHandle instance.
func NewMTLSharedTextureHandle() MTLSharedTextureHandle {
	class := getMTLSharedTextureHandleClass()
	rv := objc.Send[MTLSharedTextureHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (s MTLSharedTextureHandle) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The device object that created the texture.
//
// # Discussion
// 
// A texture is always associated with the [MTLDevice] that created it and can
// be used only with that device.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedTextureHandle/device
func (s MTLSharedTextureHandle) Device() MTLDevice {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}



// A string that identifies the texture.
//
// # Discussion
// 
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedTextureHandle/label
func (s MTLSharedTextureHandle) Label() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

























