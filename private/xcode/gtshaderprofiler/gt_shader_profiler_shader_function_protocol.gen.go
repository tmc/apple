// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderFunction protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction
type GTShaderProfilerShaderFunction interface {
	objectivec.IObject

	// Index protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/index
	Index() uint32

	// LibraryObjectId protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/libraryObjectId
	LibraryObjectId() uint64

	// ObjectId protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/objectId
	ObjectId() uint64

	// PointerId protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/pointerId
	PointerId() uint64

	// Type protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/type
	Type() uint32
}

// GTShaderProfilerShaderFunctionObject wraps an existing Objective-C object that conforms to the GTShaderProfilerShaderFunction protocol.
type GTShaderProfilerShaderFunctionObject struct {
	objectivec.Object
}

func (o GTShaderProfilerShaderFunctionObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerShaderFunctionObjectFromID constructs a [GTShaderProfilerShaderFunctionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerShaderFunctionObjectFromID(id objc.ID) GTShaderProfilerShaderFunctionObject {
	return GTShaderProfilerShaderFunctionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/filePath
func (o GTShaderProfilerShaderFunctionObject) FilePath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filePath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/index
func (o GTShaderProfilerShaderFunctionObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/libraryObjectId
func (o GTShaderProfilerShaderFunctionObject) LibraryObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("libraryObjectId"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/name
func (o GTShaderProfilerShaderFunctionObject) Name() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/objectId
func (o GTShaderProfilerShaderFunctionObject) ObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("objectId"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/pointerId
func (o GTShaderProfilerShaderFunctionObject) PointerId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pointerId"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderFunction/type
func (o GTShaderProfilerShaderFunctionObject) Type() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("type"))
	return rv
}
