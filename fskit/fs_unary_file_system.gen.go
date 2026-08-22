// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSUnaryFileSystem] class.
var (
	_FSUnaryFileSystemClass     FSUnaryFileSystemClass
	_FSUnaryFileSystemClassOnce sync.Once
)

func getFSUnaryFileSystemClass() FSUnaryFileSystemClass {
	_FSUnaryFileSystemClassOnce.Do(func() {
		_FSUnaryFileSystemClass = FSUnaryFileSystemClass{class: objc.GetClass("FSUnaryFileSystem")}
	})
	return _FSUnaryFileSystemClass
}

// GetFSUnaryFileSystemClass returns the class object for FSUnaryFileSystem.
func GetFSUnaryFileSystemClass() FSUnaryFileSystemClass {
	return getFSUnaryFileSystemClass()
}

type FSUnaryFileSystemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSUnaryFileSystemClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSUnaryFileSystemClass) Alloc() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for implementing a minimal file system.
//
// # Overview
//
// [FSUnaryFileSystem] is a simplified file system, which works with one
// [FSResource] and presents it as one [FSVolume].
//
// The one volume and its container have a shared state and lifetime, a more
// constrained life cycle than the [FSFileSystem] design flow.
//
// Implement your app extension by providing a subclass of [FSUnaryFileSystem]
// as a delegate object. Your delegate also needs to implement the
// [FSUnaryFileSystemOperations] protocol so that it can load resources.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystem
type FSUnaryFileSystem struct {
	objectivec.Object
}

// FSUnaryFileSystemFromID constructs a [FSUnaryFileSystem] from an objc.ID.
//
// An abstract base class for implementing a minimal file system.
func FSUnaryFileSystemFromID(id objc.ID) FSUnaryFileSystem {
	return FSUnaryFileSystem{objectivec.Object{ID: id}}
}

// NOTE: FSUnaryFileSystem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSUnaryFileSystem] class.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystem
type IFSUnaryFileSystem interface {
	objectivec.IObject

	// The status of the file system container, indicating its readiness and activity.
	ContainerStatus() IFSContainerStatus
	// Wipes existing file systems on the specified resource.
	WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion ErrorHandler)
}

// Init initializes the instance.
func (u FSUnaryFileSystem) Init() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u FSUnaryFileSystem) Autorelease() FSUnaryFileSystem {
	rv := objc.Send[FSUnaryFileSystem](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSUnaryFileSystem creates a new FSUnaryFileSystem instance.
func NewFSUnaryFileSystem() FSUnaryFileSystem {
	class := getFSUnaryFileSystemClass()
	rv := objc.Send[FSUnaryFileSystem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The status of the file system container, indicating its readiness and
// activity.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase/containerStatus
func (u FSUnaryFileSystem) ContainerStatus() IFSContainerStatus {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("containerStatus"))
	return FSContainerStatusFromID(rv)
}

// Wipes existing file systems on the specified resource.
//
// resource: The [FSBlockDeviceResource] to wipe.
//
// completion: A block or closure that executes after the wipe operation completes. The
// completion handler receives a single parameter indicating any error that
// occurs during the operation. If the value is `nil`, the wipe operation
// succeeded.
//
// # Discussion
//
// This method wraps the `wipefs` functionality from `libutil`. For more
// information, see the `man` page for `wipefs`.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase/wipe(_:completionHandler:)
func (u FSUnaryFileSystem) WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion ErrorHandler) {
	_block1, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](u.ID, objc.Sel("wipeResource:completionHandler:"), resource, _block1)
}

// Protocol methods for FSFileSystemBase

// The status of the file system container, indicating its readiness and
// activity.
//
// # Discussion
//
// A file system container starts in the [FSContainerStateNotReady] state, and
// then transitions to the other values of the [FSContainerState] enumeration.
// The following diagram illustrates the possible state transitions.
//
// [fs-file-system-base]
//
// Your file system implementation updates this property as it changes state.
// Many events and operations may trigger a state transition, and some
// transitions depend on a specific file system’s design.
//
// When using [FSBlockDeviceResource], implement the following common state
// transitions:
//
// - Calling `loadResource` transitions the state out of
// [FSContainerStateNotReady]. For all block device file systems, this
// operation changes the state to either [FSContainerStateReady] or
// [FSContainerStateBlocked]. - Calling `unloadResource` transitions to the
// [FSContainerStateNotReady] state, as does device termination. -
// Transitioning from [FSContainerStateBlocked] to [FSContainerStateReady]
// occurs as a result of resolving the underlying block favorably. -
// Transitioning from [FSContainerStateReady] to [FSContainerStateBlocked] is
// unusal, but valid. - Transitioning between [FSContainerStateReady] and
// [FSContainerStateActive] can result from maintenance operations such as
// [StartCheckWithTaskOptionsError]. For a [FSUnaryFileSystem], this
// transition can also occur when activating or deactivating the container’s
// single volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase/containerStatus
//
// [FSContainerState]: https://developer.apple.com/documentation/FSKit/FSContainerState
func (o FSUnaryFileSystem) SetContainerStatus(value IFSContainerStatus) {
	objc.Send[struct{}](o.ID, objc.Sel("setContainerStatus:"), value)
}

// WipeResource is a synchronous wrapper around [FSUnaryFileSystem.WipeResourceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u FSUnaryFileSystem) WipeResource(ctx context.Context, resource IFSBlockDeviceResource) error {
	done := make(chan error, 1)
	u.WipeResourceCompletionHandler(resource, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
