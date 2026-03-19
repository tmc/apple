// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNHumanBodyRecognizedPoint3D] class.
var (
	_VNHumanBodyRecognizedPoint3DClass     VNHumanBodyRecognizedPoint3DClass
	_VNHumanBodyRecognizedPoint3DClassOnce sync.Once
)

func getVNHumanBodyRecognizedPoint3DClass() VNHumanBodyRecognizedPoint3DClass {
	_VNHumanBodyRecognizedPoint3DClassOnce.Do(func() {
		_VNHumanBodyRecognizedPoint3DClass = VNHumanBodyRecognizedPoint3DClass{class: objc.GetClass("VNHumanBodyRecognizedPoint3D")}
	})
	return _VNHumanBodyRecognizedPoint3DClass
}

// GetVNHumanBodyRecognizedPoint3DClass returns the class object for VNHumanBodyRecognizedPoint3D.
func GetVNHumanBodyRecognizedPoint3DClass() VNHumanBodyRecognizedPoint3DClass {
	return getVNHumanBodyRecognizedPoint3DClass()
}

type VNHumanBodyRecognizedPoint3DClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHumanBodyRecognizedPoint3DClass) Alloc() VNHumanBodyRecognizedPoint3D {
	rv := objc.Send[VNHumanBodyRecognizedPoint3D](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A recognized 3D point that includes a parent joint.
//
// # Getting the Position
//
//   - [VNHumanBodyRecognizedPoint3D.LocalPosition]: The three-dimensional position.
//
// # Getting the Parent Joint
//
//   - [VNHumanBodyRecognizedPoint3D.ParentJoint]: The parent joint in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D
type VNHumanBodyRecognizedPoint3D struct {
	VNRecognizedPoint3D
}

// VNHumanBodyRecognizedPoint3DFromID constructs a [VNHumanBodyRecognizedPoint3D] from an objc.ID.
//
// A recognized 3D point that includes a parent joint.
func VNHumanBodyRecognizedPoint3DFromID(id objc.ID) VNHumanBodyRecognizedPoint3D {
	return VNHumanBodyRecognizedPoint3D{VNRecognizedPoint3D: VNRecognizedPoint3DFromID(id)}
}
// NOTE: VNHumanBodyRecognizedPoint3D adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNHumanBodyRecognizedPoint3D] class.
//
// # Getting the Position
//
//   - [IVNHumanBodyRecognizedPoint3D.LocalPosition]: The three-dimensional position.
//
// # Getting the Parent Joint
//
//   - [IVNHumanBodyRecognizedPoint3D.ParentJoint]: The parent joint in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D
type IVNHumanBodyRecognizedPoint3D interface {
	IVNRecognizedPoint3D

	// Topic: Getting the Position

	// The three-dimensional position.
	LocalPosition() objectivec.IObject

	// Topic: Getting the Parent Joint

	// The parent joint in the observation.
	ParentJoint() VNHumanBodyPose3DObservationJointName
}

// Init initializes the instance.
func (h VNHumanBodyRecognizedPoint3D) Init() VNHumanBodyRecognizedPoint3D {
	rv := objc.Send[VNHumanBodyRecognizedPoint3D](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHumanBodyRecognizedPoint3D) Autorelease() VNHumanBodyRecognizedPoint3D {
	rv := objc.Send[VNHumanBodyRecognizedPoint3D](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHumanBodyRecognizedPoint3D creates a new VNHumanBodyRecognizedPoint3D instance.
func NewVNHumanBodyRecognizedPoint3D() VNHumanBodyRecognizedPoint3D {
	class := getVNHumanBodyRecognizedPoint3DClass()
	rv := objc.Send[VNHumanBodyRecognizedPoint3D](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a point object with the position you specify.
//
// position: The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNPoint3D/init(position:)
// position is a [simd.simd_float4x4].
func NewHumanBodyRecognizedPoint3DWithPosition(position objectivec.IObject) VNHumanBodyRecognizedPoint3D {
	instance := getVNHumanBodyRecognizedPoint3DClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPosition:"), position)
	return VNHumanBodyRecognizedPoint3DFromID(rv)
}

// The three-dimensional position.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D/localPosition
func (h VNHumanBodyRecognizedPoint3D) LocalPosition() objectivec.IObject {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("localPosition"))
	return objectivec.Object{ID: rv}
}
// The parent joint in the observation.
//
// See: https://developer.apple.com/documentation/Vision/VNHumanBodyRecognizedPoint3D/parentJoint
func (h VNHumanBodyRecognizedPoint3D) ParentJoint() VNHumanBodyPose3DObservationJointName {
	rv := objc.Send[VNHumanBodyPose3DObservationJointName](h.ID, objc.Sel("parentJoint"))
	return VNHumanBodyPose3DObservationJointName(rv)
}

