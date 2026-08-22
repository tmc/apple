// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSFileSystem] class.
var (
	_FSFileSystemClass     FSFileSystemClass
	_FSFileSystemClassOnce sync.Once
)

func getFSFileSystemClass() FSFileSystemClass {
	_FSFileSystemClassOnce.Do(func() {
		_FSFileSystemClass = FSFileSystemClass{class: objc.GetClass("FSFileSystem")}
	})
	return _FSFileSystemClass
}

// GetFSFileSystemClass returns the class object for FSFileSystem.
func GetFSFileSystemClass() FSFileSystemClass {
	return getFSFileSystemClass()
}

type FSFileSystemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSFileSystemClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSFileSystemClass) Alloc() FSFileSystem {
	rv := objc.Send[FSFileSystem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for implementing a full-featured file system.
//
// # Overview
//
// [FSFileSystem] is a full-featured file system, which works with one or more
// [FSResource] instances and presents one or more [FSVolume] references to
// callers.
//
// Implement your app extension by providing a subclass of [FSFileSystem] as a
// delegate object. Your delegate also needs to implement the
// [FSFileSystemOperations] protocol so that it can probe, load, and unload
// resources.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystem
type FSFileSystem struct {
	objectivec.Object
}

// FSFileSystemFromID constructs a [FSFileSystem] from an objc.ID.
//
// An abstract base class for implementing a full-featured file system.
func FSFileSystemFromID(id objc.ID) FSFileSystem {
	return FSFileSystem{objectivec.Object{ID: id}}
}

// NOTE: FSFileSystem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSFileSystem] class.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystem
type IFSFileSystem interface {
	objectivec.IObject

	// The status of the file system container, indicating its readiness and activity.
	ContainerStatus() IFSContainerStatus
	// Wipes existing file systems on the specified resource.
	WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion ErrorHandler)
}

// Init initializes the instance.
func (f FSFileSystem) Init() FSFileSystem {
	rv := objc.Send[FSFileSystem](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FSFileSystem) Autorelease() FSFileSystem {
	rv := objc.Send[FSFileSystem](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSFileSystem creates a new FSFileSystem instance.
func NewFSFileSystem() FSFileSystem {
	class := getFSFileSystemClass()
	rv := objc.Send[FSFileSystem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The status of the file system container, indicating its readiness and
// activity.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase/containerStatus
func (f FSFileSystem) ContainerStatus() IFSContainerStatus {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("containerStatus"))
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
func (f FSFileSystem) WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion ErrorHandler) {
	_block1, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](f.ID, objc.Sel("wipeResource:completionHandler:"), resource, _block1)
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
func (o FSFileSystem) SetContainerStatus(value IFSContainerStatus) {
	objc.Send[struct{}](o.ID, objc.Sel("setContainerStatus:"), value)
}

// WipeResource is a synchronous wrapper around [FSFileSystem.WipeResourceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f FSFileSystem) WipeResource(ctx context.Context, resource IFSBlockDeviceResource) error {
	done := make(chan error, 1)
	f.WipeResourceCompletionHandler(resource, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
