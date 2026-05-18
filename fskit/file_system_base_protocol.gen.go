// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol containing functionality supplied by FSKit to file system implementations.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase
type FSFileSystemBase interface {
	objectivec.IObject

	// Wipes existing file systems on the specified resource.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase/wipe(_:completionHandler:)
	WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion ErrorHandler)

	// The status of the file system container, indicating its readiness and activity.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSFileSystemBase/containerStatus
	ContainerStatus() IFSContainerStatus
	SetContainerStatus(value IFSContainerStatus)
}

// FSFileSystemBaseObject wraps an existing Objective-C object that conforms to the FSFileSystemBase protocol.
type FSFileSystemBaseObject struct {
	objectivec.Object
}

func (o FSFileSystemBaseObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSFileSystemBaseObjectFromID constructs a [FSFileSystemBaseObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSFileSystemBaseObjectFromID(id objc.ID) FSFileSystemBaseObject {
	return FSFileSystemBaseObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o FSFileSystemBaseObject) WipeResourceCompletionHandler(resource IFSBlockDeviceResource, completion ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("wipeResource:completionHandler:"), resource, completion)
}

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
func (o FSFileSystemBaseObject) ContainerStatus() IFSContainerStatus {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("containerStatus"))
	return FSContainerStatusFromID(rv)
}

func (o FSFileSystemBaseObject) SetContainerStatus(value IFSContainerStatus) {
	objc.Send[struct{}](o.ID, objc.Sel("setContainerStatus:"), value)
}
