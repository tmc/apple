// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderFunction protocol.
type GTShaderProfilerShaderFunction interface {
	objectivec.IObject

	// Index protocol.
	Index() uint32

	// LibraryObjectId protocol.
	LibraryObjectId() uint64

	// ObjectId protocol.
	ObjectId() uint64

	// PointerId protocol.
	PointerId() uint64

	// Type protocol.
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

func (o GTShaderProfilerShaderFunctionObject) FilePath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filePath"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderFunctionObject) Index() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("index"))
	return rv
}
func (o GTShaderProfilerShaderFunctionObject) LibraryObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("libraryObjectId"))
	return rv
}
func (o GTShaderProfilerShaderFunctionObject) Name() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderFunctionObject) ObjectId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("objectId"))
	return rv
}
func (o GTShaderProfilerShaderFunctionObject) PointerId() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("pointerId"))
	return rv
}
func (o GTShaderProfilerShaderFunctionObject) Type() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("type"))
	return rv
}
