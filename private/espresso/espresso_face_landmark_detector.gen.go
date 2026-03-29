// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoFaceLandmarkDetector] class.
var (
	_EspressoFaceLandmarkDetectorClass     EspressoFaceLandmarkDetectorClass
	_EspressoFaceLandmarkDetectorClassOnce sync.Once
)

func getEspressoFaceLandmarkDetectorClass() EspressoFaceLandmarkDetectorClass {
	_EspressoFaceLandmarkDetectorClassOnce.Do(func() {
		_EspressoFaceLandmarkDetectorClass = EspressoFaceLandmarkDetectorClass{class: objc.GetClass("EspressoFaceLandmarkDetector")}
	})
	return _EspressoFaceLandmarkDetectorClass
}

// GetEspressoFaceLandmarkDetectorClass returns the class object for EspressoFaceLandmarkDetector.
func GetEspressoFaceLandmarkDetectorClass() EspressoFaceLandmarkDetectorClass {
	return getEspressoFaceLandmarkDetectorClass()
}

type EspressoFaceLandmarkDetectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoFaceLandmarkDetectorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoFaceLandmarkDetectorClass) Alloc() EspressoFaceLandmarkDetector {
	rv := objc.Send[EspressoFaceLandmarkDetector](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoFaceLandmarkDetector.CommonInitContextPlatformComputePath]
//   - [EspressoFaceLandmarkDetector.DetectFaceSublandmarkDoFaceRectFix]
//   - [EspressoFaceLandmarkDetector.DetectInImageRectFaceSublandmark]
//   - [EspressoFaceLandmarkDetector.ExtractSubLandmarksFromCrop192Centers192]
//   - [EspressoFaceLandmarkDetector.Newface]
//   - [EspressoFaceLandmarkDetector.SetNewface]
//   - [EspressoFaceLandmarkDetector.InitWithNetworkAtPathContextPlatformComputePath]
//   - [EspressoFaceLandmarkDetector.InitWithNetworkAtPathContextObjCPlatformComputePath]
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector
type EspressoFaceLandmarkDetector struct {
	objectivec.Object
}

// EspressoFaceLandmarkDetectorFromID constructs a [EspressoFaceLandmarkDetector] from an objc.ID.
func EspressoFaceLandmarkDetectorFromID(id objc.ID) EspressoFaceLandmarkDetector {
	return EspressoFaceLandmarkDetector{objectivec.Object{ID: id}}
}
// Ensure EspressoFaceLandmarkDetector implements IEspressoFaceLandmarkDetector.
var _ IEspressoFaceLandmarkDetector = EspressoFaceLandmarkDetector{}

// An interface definition for the [EspressoFaceLandmarkDetector] class.
//
// # Methods
//
//   - [IEspressoFaceLandmarkDetector.CommonInitContextPlatformComputePath]
//   - [IEspressoFaceLandmarkDetector.DetectFaceSublandmarkDoFaceRectFix]
//   - [IEspressoFaceLandmarkDetector.DetectInImageRectFaceSublandmark]
//   - [IEspressoFaceLandmarkDetector.ExtractSubLandmarksFromCrop192Centers192]
//   - [IEspressoFaceLandmarkDetector.Newface]
//   - [IEspressoFaceLandmarkDetector.SetNewface]
//   - [IEspressoFaceLandmarkDetector.InitWithNetworkAtPathContextPlatformComputePath]
//   - [IEspressoFaceLandmarkDetector.InitWithNetworkAtPathContextObjCPlatformComputePath]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector
type IEspressoFaceLandmarkDetector interface {
	objectivec.IObject

	// Topic: Methods

	CommonInitContextPlatformComputePath(init_ objectivec.IObject, context objectivec.IObject, platform int, path int)
	DetectFaceSublandmarkDoFaceRectFix(detect unsafe.Pointer, face corefoundation.CGRect, sublandmark bool, fix bool) objectivec.IObject
	DetectInImageRectFaceSublandmark(rect unsafe.Pointer, face corefoundation.CGRect, sublandmark bool) objectivec.IObject
	ExtractSubLandmarksFromCrop192Centers192(crop192 unsafe.Pointer, centers192 unsafe.Pointer) objectivec.IObject
	Newface() corefoundation.CGRect
	SetNewface(value corefoundation.CGRect)
	InitWithNetworkAtPathContextPlatformComputePath(path objectivec.IObject, context objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector
	InitWithNetworkAtPathContextObjCPlatformComputePath(path objectivec.IObject, c objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector
}

// Init initializes the instance.
func (e EspressoFaceLandmarkDetector) Init() EspressoFaceLandmarkDetector {
	rv := objc.Send[EspressoFaceLandmarkDetector](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoFaceLandmarkDetector) Autorelease() EspressoFaceLandmarkDetector {
	rv := objc.Send[EspressoFaceLandmarkDetector](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoFaceLandmarkDetector creates a new EspressoFaceLandmarkDetector instance.
func NewEspressoFaceLandmarkDetector() EspressoFaceLandmarkDetector {
	class := getEspressoFaceLandmarkDetectorClass()
	rv := objc.Send[EspressoFaceLandmarkDetector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/initWithNetworkAtPath:contextObjC:platform:computePath:
func NewEspressoFaceLandmarkDetectorWithNetworkAtPathContextObjCPlatformComputePath(path objectivec.IObject, c objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector {
	instance := getEspressoFaceLandmarkDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkAtPath:contextObjC:platform:computePath:"), path, c, platform, path2)
	return EspressoFaceLandmarkDetectorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/initWithNetworkAtPath:context:platform:computePath:
func NewEspressoFaceLandmarkDetectorWithNetworkAtPathContextPlatformComputePath(path objectivec.IObject, context objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector {
	instance := getEspressoFaceLandmarkDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkAtPath:context:platform:computePath:"), path, context, platform, path2)
	return EspressoFaceLandmarkDetectorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/commonInit:context:platform:computePath:
func (e EspressoFaceLandmarkDetector) CommonInitContextPlatformComputePath(init_ objectivec.IObject, context objectivec.IObject, platform int, path int) {
	objc.Send[objc.ID](e.ID, objc.Sel("commonInit:context:platform:computePath:"), init_, context, platform, path)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/detect:face:sublandmark:doFaceRectFix:
func (e EspressoFaceLandmarkDetector) DetectFaceSublandmarkDoFaceRectFix(detect unsafe.Pointer, face corefoundation.CGRect, sublandmark bool, fix bool) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("detect:face:sublandmark:doFaceRectFix:"), detect, face, sublandmark, fix)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/detectInImageRect:face:sublandmark:
func (e EspressoFaceLandmarkDetector) DetectInImageRectFaceSublandmark(rect unsafe.Pointer, face corefoundation.CGRect, sublandmark bool) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("detectInImageRect:face:sublandmark:"), rect, face, sublandmark)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/extractSubLandmarksFromCrop192:centers192:
func (e EspressoFaceLandmarkDetector) ExtractSubLandmarksFromCrop192Centers192(crop192 unsafe.Pointer, centers192 unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("extractSubLandmarksFromCrop192:centers192:"), crop192, centers192)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/initWithNetworkAtPath:context:platform:computePath:
func (e EspressoFaceLandmarkDetector) InitWithNetworkAtPathContextPlatformComputePath(path objectivec.IObject, context objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector {
	rv := objc.Send[EspressoFaceLandmarkDetector](e.ID, objc.Sel("initWithNetworkAtPath:context:platform:computePath:"), path, context, platform, path2)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/initWithNetworkAtPath:contextObjC:platform:computePath:
func (e EspressoFaceLandmarkDetector) InitWithNetworkAtPathContextObjCPlatformComputePath(path objectivec.IObject, c objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector {
	rv := objc.Send[EspressoFaceLandmarkDetector](e.ID, objc.Sel("initWithNetworkAtPath:contextObjC:platform:computePath:"), path, c, platform, path2)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoFaceLandmarkDetector/newface
func (e EspressoFaceLandmarkDetector) Newface() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](e.ID, objc.Sel("newface"))
	return corefoundation.CGRect(rv)
}
func (e EspressoFaceLandmarkDetector) SetNewface(value corefoundation.CGRect) {
	objc.Send[struct{}](e.ID, objc.Sel("setNewface:"), value)
}

