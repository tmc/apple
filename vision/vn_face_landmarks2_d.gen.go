// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNFaceLandmarks2D] class.
var (
	_VNFaceLandmarks2DClass     VNFaceLandmarks2DClass
	_VNFaceLandmarks2DClassOnce sync.Once
)

func getVNFaceLandmarks2DClass() VNFaceLandmarks2DClass {
	_VNFaceLandmarks2DClassOnce.Do(func() {
		_VNFaceLandmarks2DClass = VNFaceLandmarks2DClass{class: objc.GetClass("VNFaceLandmarks2D")}
	})
	return _VNFaceLandmarks2DClass
}

// GetVNFaceLandmarks2DClass returns the class object for VNFaceLandmarks2D.
func GetVNFaceLandmarks2DClass() VNFaceLandmarks2DClass {
	return getVNFaceLandmarks2DClass()
}

type VNFaceLandmarks2DClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNFaceLandmarks2DClass) Alloc() VNFaceLandmarks2D {
	rv := objc.Send[VNFaceLandmarks2D](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A collection of facial features that a request detects.
//
// # Overview
// 
// This class represents the set of all detectable 2D face landmarks and
// regions, exposed as properties. The coordinates of the face landmarks are
// normalized to the dimensions of the face observation’s [VNFaceLandmarks2D.BoundingBox],
// with the origin at the bounding box’s lower-left corner. Use the
// [VNImagePointForFaceLandmarkPoint] function to convert normalized face
// landmark points into absolute points within the image’s coordinate
// system.
//
// # Face Landmark Points
//
//   - [VNFaceLandmarks2D.AllPoints]: The region containing all face landmark points.
//   - [VNFaceLandmarks2D.FaceContour]: The region containing points that trace the face contour from the left cheek, over the chin, to the right cheek.
//   - [VNFaceLandmarks2D.LeftEye]: The region containing points that outline the left eye.
//   - [VNFaceLandmarks2D.RightEye]: The region containing points that outline the right eye.
//   - [VNFaceLandmarks2D.LeftEyebrow]: The region containing points that trace the left eyebrow.
//   - [VNFaceLandmarks2D.RightEyebrow]: The region containing points that trace the right eyebrow.
//   - [VNFaceLandmarks2D.Nose]: The region containing points that outline the nose.
//   - [VNFaceLandmarks2D.NoseCrest]: The region containing points that trace the center crest of the nose.
//   - [VNFaceLandmarks2D.MedianLine]: The region containing points that trace a vertical line down the center of the face.
//   - [VNFaceLandmarks2D.OuterLips]: The region containing points that outline the outside of the lips.
//   - [VNFaceLandmarks2D.InnerLips]: The region containing points that outline the space between the lips.
//   - [VNFaceLandmarks2D.LeftPupil]: The region containing the point where the left pupil is located.
//   - [VNFaceLandmarks2D.RightPupil]: The region containing the point where the right pupil is located.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D
type VNFaceLandmarks2D struct {
	VNFaceLandmarks
}

// VNFaceLandmarks2DFromID constructs a [VNFaceLandmarks2D] from an objc.ID.
//
// A collection of facial features that a request detects.
func VNFaceLandmarks2DFromID(id objc.ID) VNFaceLandmarks2D {
	return VNFaceLandmarks2D{VNFaceLandmarks: VNFaceLandmarksFromID(id)}
}
// NOTE: VNFaceLandmarks2D adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNFaceLandmarks2D] class.
//
// # Face Landmark Points
//
//   - [IVNFaceLandmarks2D.AllPoints]: The region containing all face landmark points.
//   - [IVNFaceLandmarks2D.FaceContour]: The region containing points that trace the face contour from the left cheek, over the chin, to the right cheek.
//   - [IVNFaceLandmarks2D.LeftEye]: The region containing points that outline the left eye.
//   - [IVNFaceLandmarks2D.RightEye]: The region containing points that outline the right eye.
//   - [IVNFaceLandmarks2D.LeftEyebrow]: The region containing points that trace the left eyebrow.
//   - [IVNFaceLandmarks2D.RightEyebrow]: The region containing points that trace the right eyebrow.
//   - [IVNFaceLandmarks2D.Nose]: The region containing points that outline the nose.
//   - [IVNFaceLandmarks2D.NoseCrest]: The region containing points that trace the center crest of the nose.
//   - [IVNFaceLandmarks2D.MedianLine]: The region containing points that trace a vertical line down the center of the face.
//   - [IVNFaceLandmarks2D.OuterLips]: The region containing points that outline the outside of the lips.
//   - [IVNFaceLandmarks2D.InnerLips]: The region containing points that outline the space between the lips.
//   - [IVNFaceLandmarks2D.LeftPupil]: The region containing the point where the left pupil is located.
//   - [IVNFaceLandmarks2D.RightPupil]: The region containing the point where the right pupil is located.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D
type IVNFaceLandmarks2D interface {
	IVNFaceLandmarks

	// Topic: Face Landmark Points

	// The region containing all face landmark points.
	AllPoints() IVNFaceLandmarkRegion2D
	// The region containing points that trace the face contour from the left cheek, over the chin, to the right cheek.
	FaceContour() IVNFaceLandmarkRegion2D
	// The region containing points that outline the left eye.
	LeftEye() IVNFaceLandmarkRegion2D
	// The region containing points that outline the right eye.
	RightEye() IVNFaceLandmarkRegion2D
	// The region containing points that trace the left eyebrow.
	LeftEyebrow() IVNFaceLandmarkRegion2D
	// The region containing points that trace the right eyebrow.
	RightEyebrow() IVNFaceLandmarkRegion2D
	// The region containing points that outline the nose.
	Nose() IVNFaceLandmarkRegion2D
	// The region containing points that trace the center crest of the nose.
	NoseCrest() IVNFaceLandmarkRegion2D
	// The region containing points that trace a vertical line down the center of the face.
	MedianLine() IVNFaceLandmarkRegion2D
	// The region containing points that outline the outside of the lips.
	OuterLips() IVNFaceLandmarkRegion2D
	// The region containing points that outline the space between the lips.
	InnerLips() IVNFaceLandmarkRegion2D
	// The region containing the point where the left pupil is located.
	LeftPupil() IVNFaceLandmarkRegion2D
	// The region containing the point where the right pupil is located.
	RightPupil() IVNFaceLandmarkRegion2D

	// The bounding box of the object that the request detects.
	BoundingBox() corefoundation.CGRect
	SetBoundingBox(value corefoundation.CGRect)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (f VNFaceLandmarks2D) Init() VNFaceLandmarks2D {
	rv := objc.Send[VNFaceLandmarks2D](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VNFaceLandmarks2D) Autorelease() VNFaceLandmarks2D {
	rv := objc.Send[VNFaceLandmarks2D](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNFaceLandmarks2D creates a new VNFaceLandmarks2D instance.
func NewVNFaceLandmarks2D() VNFaceLandmarks2D {
	class := getVNFaceLandmarks2DClass()
	rv := objc.Send[VNFaceLandmarks2D](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (f VNFaceLandmarks2D) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The region containing all face landmark points.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/allPoints
func (f VNFaceLandmarks2D) AllPoints() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("allPoints"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that trace the face contour from the left
// cheek, over the chin, to the right cheek.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/faceContour
func (f VNFaceLandmarks2D) FaceContour() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("faceContour"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that outline the left eye.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/leftEye
func (f VNFaceLandmarks2D) LeftEye() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("leftEye"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that outline the right eye.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/rightEye
func (f VNFaceLandmarks2D) RightEye() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("rightEye"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that trace the left eyebrow.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/leftEyebrow
func (f VNFaceLandmarks2D) LeftEyebrow() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("leftEyebrow"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that trace the right eyebrow.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/rightEyebrow
func (f VNFaceLandmarks2D) RightEyebrow() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("rightEyebrow"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that outline the nose.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/nose
func (f VNFaceLandmarks2D) Nose() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("nose"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that trace the center crest of the nose.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/noseCrest
func (f VNFaceLandmarks2D) NoseCrest() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("noseCrest"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that trace a vertical line down the center of
// the face.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/medianLine
func (f VNFaceLandmarks2D) MedianLine() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("medianLine"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that outline the outside of the lips.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/outerLips
func (f VNFaceLandmarks2D) OuterLips() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("outerLips"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing points that outline the space between the lips.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/innerLips
func (f VNFaceLandmarks2D) InnerLips() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("innerLips"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing the point where the left pupil is located.
//
// # Discussion
// 
// This value may be inaccurate if the eye is blinking.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/leftPupil
func (f VNFaceLandmarks2D) LeftPupil() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("leftPupil"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The region containing the point where the right pupil is located.
//
// # Discussion
// 
// This value may be inaccurate if the eye is blinking.
//
// See: https://developer.apple.com/documentation/Vision/VNFaceLandmarks2D/rightPupil
func (f VNFaceLandmarks2D) RightPupil() IVNFaceLandmarkRegion2D {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("rightPupil"))
	return VNFaceLandmarkRegion2DFromID(objc.ID(rv))
}



// The bounding box of the object that the request detects.
//
// See: https://developer.apple.com/documentation/vision/vndetectedobjectobservation/boundingbox
func (f VNFaceLandmarks2D) BoundingBox() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](f.ID, objc.Sel("boundingBox"))
	return corefoundation.CGRect(rv)
}
func (f VNFaceLandmarks2D) SetBoundingBox(value corefoundation.CGRect) {
	objc.Send[struct{}](f.ID, objc.Sel("setBoundingBox:"), value)
}



























