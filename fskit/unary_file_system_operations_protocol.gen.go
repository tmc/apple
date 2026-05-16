// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Operations performed by a unary file system.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations
type FSUnaryFileSystemOperations interface {
	objectivec.IObject

	// Requests that the file system load a resource and present it as a volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/loadResource(resource:options:replyHandler:)
	LoadResourceOptionsReplyHandler(resource IFSResource, options IFSTaskOptions, reply FSVolumeErrorHandler)

	// Requests that the file system unload the specified resource.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/unloadResource(resource:options:replyHandler:)
	UnloadResourceOptionsReplyHandler(resource IFSResource, options IFSTaskOptions, reply ErrorHandler)

	// Requests that the file system probe the specified resource.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/probeResource(resource:replyHandler:)
	ProbeResourceReplyHandler(resource IFSResource, reply FSProbeResultErrorHandler)
}

// FSUnaryFileSystemOperationsObject wraps an existing Objective-C object that conforms to the FSUnaryFileSystemOperations protocol.
type FSUnaryFileSystemOperationsObject struct {
	objectivec.Object
}

func (o FSUnaryFileSystemOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSUnaryFileSystemOperationsObjectFromID constructs a [FSUnaryFileSystemOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSUnaryFileSystemOperationsObjectFromID(id objc.ID) FSUnaryFileSystemOperationsObject {
	return FSUnaryFileSystemOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Requests that the file system load a resource and present it as a volume.
//
// resource: An [FSResource] to load.
//
// options: An [FSTaskOptions] object specifying options to apply when loading the
// resource. An [FSUnaryFileSystem] supports two options: `-f` for “force”
// and `--rdonly` for read-only. The file system must remember if the
// read-only option is present.
//
// reply: A block or closure that your implementation invokes when it finishes
// setting up or encounters an error. Pass a subclass of [FSVolume] as the
// first parameter if loading succeeds. If loading fails, pass an error as the
// second parameter.
//
// # Discussion
//
// Implement this method by inspecting the provided resource and verifying it
// uses a supported format. If the resource does use a supported format,
// create a subclass of [FSVolume], clear the container error state, and
// invoke the `reply` callback, passing your volume as a parameter. If loading
// can’t proceed, invoke `reply` and send an appropriate error as the second
// parameter.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/loadResource(resource:options:replyHandler:)
func (o FSUnaryFileSystemOperationsObject) LoadResourceOptionsReplyHandler(resource IFSResource, options IFSTaskOptions, reply FSVolumeErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("loadResource:options:replyHandler:"), resource, options, reply)
}

// Requests that the file system unload the specified resource.
//
// resource: An [FSResource] to unload.
//
// options: An [FSTaskOptions] object specifying options to apply when unloading the
// resource.
//
// reply: A block or closure that your implementation invokes when it finishes
// unloading or encounters an error. If unloading fails, pass an error as the
// parameter to describe the problem. Otherwise, pass `nil`.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/unloadResource(resource:options:replyHandler:)
func (o FSUnaryFileSystemOperationsObject) UnloadResourceOptionsReplyHandler(resource IFSResource, options IFSTaskOptions, reply ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("unloadResource:options:replyHandler:"), resource, options, reply)
}

// Requests that the file system probe the specified resource.
//
// resource: The [FSResource] to probe.
//
// reply: A block or closure that your implementation invokes when it finishes the
// probe or encounters an error. Pass an instance of [FSProbeResult] with
// probe results as the first parameter if your probe operation succeeds. If
// probing fails, pass an error as the second parameter.
//
// # Discussion
//
// Implement this method to indicate whether the resource is recognizable and
// usable.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/probeResource(resource:replyHandler:)
func (o FSUnaryFileSystemOperationsObject) ProbeResourceReplyHandler(resource IFSResource, reply FSProbeResultErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("probeResource:replyHandler:"), resource, reply)
}

// Notifies you that the system finished loading your file system extension.
//
// # Discussion
//
// The system performs this callback after the main run loop starts and before
// receiving the first message from the FSKit daemon.
//
// Implement this method if you want to perform any setup prior to receiving
// FSKit callbacks.
//
// See: https://developer.apple.com/documentation/FSKit/FSUnaryFileSystemOperations/didFinishLoading()
func (o FSUnaryFileSystemOperationsObject) DidFinishLoading() {
	objc.Send[struct{}](o.ID, objc.Sel("didFinishLoading"))
}
