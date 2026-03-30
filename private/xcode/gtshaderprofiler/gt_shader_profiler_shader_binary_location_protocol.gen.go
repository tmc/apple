// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderBinaryLocation protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation
type GTShaderProfilerShaderBinaryLocation interface {
	objectivec.IObject

	// Column protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/column
	Column() int

	// FileIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/fileIndex
	FileIndex() uint64

	// FunctionNameIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/functionNameIndex
	FunctionNameIndex() uint64

	// Line protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/line
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/binary
func (o GTShaderProfilerShaderBinaryLocationObject) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/column
func (o GTShaderProfilerShaderBinaryLocationObject) Column() int {
	rv := objc.Send[int](o.ID, objc.Sel("column"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/fileIndex
func (o GTShaderProfilerShaderBinaryLocationObject) FileIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("fileIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/fullPath
func (o GTShaderProfilerShaderBinaryLocationObject) FullPath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fullPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/functionName
func (o GTShaderProfilerShaderBinaryLocationObject) FunctionName() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("functionName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/functionNameIndex
func (o GTShaderProfilerShaderBinaryLocationObject) FunctionNameIndex() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("functionNameIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryLocation/line
func (o GTShaderProfilerShaderBinaryLocationObject) Line() int {
	rv := objc.Send[int](o.ID, objc.Sel("line"))
	return rv
}
