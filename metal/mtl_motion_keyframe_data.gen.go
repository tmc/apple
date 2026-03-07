// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLMotionKeyframeData] class.
var (
	_MTLMotionKeyframeDataClass     MTLMotionKeyframeDataClass
	_MTLMotionKeyframeDataClassOnce sync.Once
)

func getMTLMotionKeyframeDataClass() MTLMotionKeyframeDataClass {
	_MTLMotionKeyframeDataClassOnce.Do(func() {
		_MTLMotionKeyframeDataClass = MTLMotionKeyframeDataClass{class: objc.GetClass("MTLMotionKeyframeData")}
	})
	return _MTLMotionKeyframeDataClass
}

// GetMTLMotionKeyframeDataClass returns the class object for MTLMotionKeyframeData.
func GetMTLMotionKeyframeDataClass() MTLMotionKeyframeDataClass {
	return getMTLMotionKeyframeDataClass()
}

type MTLMotionKeyframeDataClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLMotionKeyframeDataClass) Alloc() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Geometry data for a specific keyframe to use in a moving instance.
//
// # Overview
// 
// An [MTLMotionKeyframeData] instance describes the location of geometry data
// for a keyframe. The exact type of data can vary, depending on which kind of
// motion descriptor you create. For an
// [MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor] instance, the
// buffer data is a list of bounding boxes. For an
// [MTLAccelerationStructureMotionTriangleGeometryDescriptor], the buffer data
// is a list of vertices.
//
// # Specifying the keyframe data
//
//   - [MTLMotionKeyframeData.Buffer]: The buffer that holds the geometry data.
//   - [MTLMotionKeyframeData.SetBuffer]
//   - [MTLMotionKeyframeData.Offset]: The offset, in bytes, to the keyframe data.
//   - [MTLMotionKeyframeData.SetOffset]
//
// See: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData
type MTLMotionKeyframeData struct {
	objectivec.Object
}

// MTLMotionKeyframeDataFromID constructs a [MTLMotionKeyframeData] from an objc.ID.
//
// Geometry data for a specific keyframe to use in a moving instance.
func MTLMotionKeyframeDataFromID(id objc.ID) MTLMotionKeyframeData {
	return MTLMotionKeyframeData{objectivec.Object{id}}
}
// NOTE: MTLMotionKeyframeData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLMotionKeyframeData] class.
//
// # Specifying the keyframe data
//
//   - [IMTLMotionKeyframeData.Buffer]: The buffer that holds the geometry data.
//   - [IMTLMotionKeyframeData.SetBuffer]
//   - [IMTLMotionKeyframeData.Offset]: The offset, in bytes, to the keyframe data.
//   - [IMTLMotionKeyframeData.SetOffset]
//
// See: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData
type IMTLMotionKeyframeData interface {
	objectivec.IObject

	// Topic: Specifying the keyframe data

	// The buffer that holds the geometry data.
	Buffer() MTLBuffer
	SetBuffer(value MTLBuffer)
	// The offset, in bytes, to the keyframe data.
	Offset() uint
	SetOffset(value uint)
}





// Init initializes the instance.
func (m MTLMotionKeyframeData) Init() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLMotionKeyframeData) Autorelease() MTLMotionKeyframeData {
	rv := objc.Send[MTLMotionKeyframeData](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLMotionKeyframeData creates a new MTLMotionKeyframeData instance.
func NewMTLMotionKeyframeData() MTLMotionKeyframeData {
	class := getMTLMotionKeyframeDataClass()
	rv := objc.Send[MTLMotionKeyframeData](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// Creates a new keyframe object.
//
// See: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/data
func (_MTLMotionKeyframeDataClass MTLMotionKeyframeDataClass) Data() MTLMotionKeyframeData {
	rv := objc.Send[objc.ID](objc.ID(_MTLMotionKeyframeDataClass.class), objc.Sel("data"))
	return MTLMotionKeyframeDataFromID(rv)
}








// The buffer that holds the geometry data.
//
// See: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/buffer
func (m MTLMotionKeyframeData) Buffer() MTLBuffer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("buffer"))
	return MTLBufferObjectFromID(rv)
}
func (m MTLMotionKeyframeData) SetBuffer(value MTLBuffer) {
	objc.Send[struct{}](m.ID, objc.Sel("setBuffer:"), value)
}



// The offset, in bytes, to the keyframe data.
//
// See: https://developer.apple.com/documentation/Metal/MTLMotionKeyframeData/offset
func (m MTLMotionKeyframeData) Offset() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("offset"))
	return rv
}
func (m MTLMotionKeyframeData) SetOffset(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setOffset:"), value)
}























