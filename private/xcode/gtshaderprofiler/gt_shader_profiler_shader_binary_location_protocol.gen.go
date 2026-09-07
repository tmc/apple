// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderBinaryLocation protocol.
type GTShaderProfilerShaderBinaryLocation interface {
	objectivec.IObject

	// Binary protocol.
	Binary() objectivec.IObject

	// Column protocol.
	Column() int

	// FileIndex protocol.
	FileIndex() uint64

	// FullPath protocol.
	FullPath() objectivec.IObject

	// FunctionName protocol.
	FunctionName() objectivec.IObject

	// FunctionNameIndex protocol.
	FunctionNameIndex() uint64

	// Line protocol.
	Line() int
}

// GTShaderProfilerShaderBinaryLocationObject wraps an existing Objective-C object that conforms to the GTShaderProfilerShaderBinaryLocation protocol.
type GTShaderProfilerShaderBinaryLocationObject struct {
	objectivec.Object
}

func (o GTShaderProfilerShaderBinaryLocationObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerShaderBinaryLocationObjectFromID constructs a [GTShaderProfilerShaderBinaryLocationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerShaderBinaryLocationObjectFromID(id objc.ID) GTShaderProfilerShaderBinaryLocationObject {
	return GTShaderProfilerShaderBinaryLocationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTShaderProfilerShaderBinaryLocationObject) Binary() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryLocationObject) Column() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("column"))
	return rv
}
func (o GTShaderProfilerShaderBinaryLocationObject) FileIndex() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("fileIndex"))
	return rv
}
func (o GTShaderProfilerShaderBinaryLocationObject) FullPath() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("fullPath"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryLocationObject) FunctionName() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("functionName"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryLocationObject) FunctionNameIndex() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("functionNameIndex"))
	return rv
}
func (o GTShaderProfilerShaderBinaryLocationObject) Line() int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("line"))
	return rv
}
