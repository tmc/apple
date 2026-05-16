// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSProbeResult] class.
var (
	_FSProbeResultClass     FSProbeResultClass
	_FSProbeResultClassOnce sync.Once
)

func getFSProbeResultClass() FSProbeResultClass {
	_FSProbeResultClassOnce.Do(func() {
		_FSProbeResultClass = FSProbeResultClass{class: objc.GetClass("FSProbeResult")}
	})
	return _FSProbeResultClass
}

// GetFSProbeResultClass returns the class object for FSProbeResult.
func GetFSProbeResultClass() FSProbeResultClass {
	return getFSProbeResultClass()
}

type FSProbeResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSProbeResultClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSProbeResultClass) Alloc() FSProbeResult {
	rv := objc.Send[FSProbeResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the results of a specific probe.
//
// # Overview
//
// For any [FSProbeResult.Result] value other than [FSMatchResult.notRecognized], ensure the
// [FSProbeResult.Name] and [FSProbeResult.ContainerID] values are non-`nil`. When a container or volume
// format doesn’t use a name, return an empty string. Also use an empty
// string in the case in which the format supports a name, but the value
// isn’t set yet.
//
// Some container or volume formats may lack a durable UUID on which to base a
// container identifier. This situation is only valid for unary file systems.
// In such a case, return a random UUID.
//
// With a block device resource, a probe operation may successfully get a
// result but encounter an error reading the name or UUID. If this happens,
// use whatever information is available, and provide an empty string or
// random UUID for the name or container ID, respectively.
//
// # Working with result properties
//
//   - [FSProbeResult.ContainerID]: The container identifier, as found during the probe operation.
//   - [FSProbeResult.Name]: The resource name, as found during the probe operation.
//   - [FSProbeResult.Result]: The match result, representing the recognition and usability of a probed resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult
//
// [FSMatchResult.notRecognized]: https://developer.apple.com/documentation/FSKit/FSMatchResult/notRecognized
type FSProbeResult struct {
	objectivec.Object
}

// FSProbeResultFromID constructs a [FSProbeResult] from an objc.ID.
//
// An object that represents the results of a specific probe.
func FSProbeResultFromID(id objc.ID) FSProbeResult {
	return FSProbeResult{objectivec.Object{ID: id}}
}

// NOTE: FSProbeResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSProbeResult] class.
//
// # Working with result properties
//
//   - [IFSProbeResult.ContainerID]: The container identifier, as found during the probe operation.
//   - [IFSProbeResult.Name]: The resource name, as found during the probe operation.
//   - [IFSProbeResult.Result]: The match result, representing the recognition and usability of a probed resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult
type IFSProbeResult interface {
	objectivec.IObject

	// Topic: Working with result properties

	// The container identifier, as found during the probe operation.
	ContainerID() IFSContainerIdentifier
	// The resource name, as found during the probe operation.
	Name() string
	// The match result, representing the recognition and usability of a probed resource.
	Result() FSMatchResult

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p FSProbeResult) Init() FSProbeResult {
	rv := objc.Send[FSProbeResult](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p FSProbeResult) Autorelease() FSProbeResult {
	rv := objc.Send[FSProbeResult](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSProbeResult creates a new FSProbeResult instance.
func NewFSProbeResult() FSProbeResult {
	class := getFSProbeResultClass()
	rv := objc.Send[FSProbeResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p FSProbeResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a probe result for a recognized file system.
//
// name: The resource name, as found during the probe operation. If the file system
// doesn’t support names, or is awaiting naming, use an empty string.
//
// containerID: The container identifier, as found during the probe operation. If the file
// system doesn’t support durable identifiers, use a random UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/recognized(name:containerID:)
func (_FSProbeResultClass FSProbeResultClass) RecognizedProbeResultWithNameContainerID(name string, containerID IFSContainerIdentifier) FSProbeResult {
	rv := objc.Send[objc.ID](objc.ID(_FSProbeResultClass.class), objc.Sel("recognizedProbeResultWithName:containerID:"), objc.String(name), containerID)
	return FSProbeResultFromID(rv)
}

// Creates a probe result for a recognized and usable file system.
//
// name: The resource name, as found during the probe operation. If the file system
// doesn’t support names, or is awaiting naming, use an empty string.
//
// containerID: The container identifier, as found during the probe operation. If the file
// system doesn’t support durable identifiers, use a random UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/usable(name:containerID:)
func (_FSProbeResultClass FSProbeResultClass) UsableProbeResultWithNameContainerID(name string, containerID IFSContainerIdentifier) FSProbeResult {
	rv := objc.Send[objc.ID](objc.ID(_FSProbeResultClass.class), objc.Sel("usableProbeResultWithName:containerID:"), objc.String(name), containerID)
	return FSProbeResultFromID(rv)
}

// Creates a probe result for a recognized file system that is usable, but
// with limited capabilities.
//
// name: The resource name, as found during the probe operation. If the file system
// doesn’t support names, or is awaiting naming, use an empty string.
//
// containerID: The container identifier, as found during the probe operation. If the file
// system doesn’t support durable identifiers, use a random UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited(name:containerID:)
func (_FSProbeResultClass FSProbeResultClass) UsableButLimitedProbeResultWithNameContainerID(name string, containerID IFSContainerIdentifier) FSProbeResult {
	rv := objc.Send[objc.ID](objc.ID(_FSProbeResultClass.class), objc.Sel("usableButLimitedProbeResultWithName:containerID:"), objc.String(name), containerID)
	return FSProbeResultFromID(rv)
}

// The container identifier, as found during the probe operation.
//
// # Discussion
//
// This value is non-`nil` unless the result is [FSMatchResult.notRecognized].
// For formats that lack a durable UUID on which to base a container
// identifier — which is only legal for a [FSUnaryFileSystem] — this value
// may be a random UUID.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/containerID
//
// [FSMatchResult.notRecognized]: https://developer.apple.com/documentation/FSKit/FSMatchResult/notRecognized
func (p FSProbeResult) ContainerID() IFSContainerIdentifier {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("containerID"))
	return FSContainerIdentifierFromID(objc.ID(rv))
}

// The resource name, as found during the probe operation.
//
// # Discussion
//
// This value is non-`nil` unless the [Result] is
// “FSMatchResult/notRecognized`. For formats that lack a name, this value
// may be an empty string. This value can also be an empty string if the
// format supports a name, but the value isn’t set yet.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/name
func (p FSProbeResult) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The match result, representing the recognition and usability of a probed
// resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/result
func (p FSProbeResult) Result() FSMatchResult {
	rv := objc.Send[FSMatchResult](p.ID, objc.Sel("result"))
	return FSMatchResult(rv)
}

// A probe result for a recognized file system that is usable, but with
// limited capabilities.
//
// # Discussion
//
// This kind of probe result lacks the [Name], [ContainerID], or both. Don’t
// return this result from probing a resource that isn’t limited.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/usableButLimited
func (_FSProbeResultClass FSProbeResultClass) UsableButLimitedProbeResult() FSProbeResult {
	rv := objc.Send[objc.ID](objc.ID(_FSProbeResultClass.class), objc.Sel("usableButLimitedProbeResult"))
	return FSProbeResultFromID(objc.ID(rv))
}

// A probe result for an unrecognized file system.
//
// # Discussion
//
// An unrecognized probe result contains `nil` for its [Name] and
// [ContainerID] properties.
//
// See: https://developer.apple.com/documentation/FSKit/FSProbeResult/notRecognized
func (_FSProbeResultClass FSProbeResultClass) NotRecognizedProbeResult() FSProbeResult {
	rv := objc.Send[objc.ID](objc.ID(_FSProbeResultClass.class), objc.Sel("notRecognizedProbeResult"))
	return FSProbeResultFromID(objc.ID(rv))
}
