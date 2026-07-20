// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
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

// # Methods
//
//   - [EspressoFaceLandmarkDetector.Newface]
//   - [EspressoFaceLandmarkDetector.SetNewface]
//   - [EspressoFaceLandmarkDetector.InitWithNetworkAtPathContextObjCPlatformComputePath]
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
//   - [IEspressoFaceLandmarkDetector.Newface]
//   - [IEspressoFaceLandmarkDetector.SetNewface]
//   - [IEspressoFaceLandmarkDetector.InitWithNetworkAtPathContextObjCPlatformComputePath]
type IEspressoFaceLandmarkDetector interface {
	objectivec.IObject

	// Topic: Methods

	Newface() corefoundation.CGRect
	SetNewface(value corefoundation.CGRect)
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

func NewEspressoFaceLandmarkDetectorWithNetworkAtPathContextObjCPlatformComputePath(path objectivec.IObject, c objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector {
	instance := getEspressoFaceLandmarkDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkAtPath:contextObjC:platform:computePath:"), path, c, platform, path2)
	return EspressoFaceLandmarkDetectorFromID(rv)
}

func NewEspressoFaceLandmarkDetectorWithNetworkAtPathContextPlatformComputePath(path objectivec.IObject, context unsafe.Pointer, platform int, path2 int) EspressoFaceLandmarkDetector {
	instance := getEspressoFaceLandmarkDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkAtPath:context:platform:computePath:"), path, context, platform, path2)
	return EspressoFaceLandmarkDetectorFromID(rv)
}

func (e EspressoFaceLandmarkDetector) InitWithNetworkAtPathContextObjCPlatformComputePath(path objectivec.IObject, c objectivec.IObject, platform int, path2 int) EspressoFaceLandmarkDetector {
	rv := objc.Send[EspressoFaceLandmarkDetector](e.ID, objc.Sel("initWithNetworkAtPath:contextObjC:platform:computePath:"), path, c, platform, path2)
	return rv
}

func (e EspressoFaceLandmarkDetector) Newface() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](e.ID, objc.Sel("newface"))
	return corefoundation.CGRect(rv)
}
func (e EspressoFaceLandmarkDetector) SetNewface(value corefoundation.CGRect) {
	objc.Send[struct{}](e.ID, objc.Sel("setNewface:"), value)
}
