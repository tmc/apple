// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDCNEspressoOverfeatDetector] class.
var (
	_EspressoDCNEspressoOverfeatDetectorClass     EspressoDCNEspressoOverfeatDetectorClass
	_EspressoDCNEspressoOverfeatDetectorClassOnce sync.Once
)

func getEspressoDCNEspressoOverfeatDetectorClass() EspressoDCNEspressoOverfeatDetectorClass {
	_EspressoDCNEspressoOverfeatDetectorClassOnce.Do(func() {
		_EspressoDCNEspressoOverfeatDetectorClass = EspressoDCNEspressoOverfeatDetectorClass{class: objc.GetClass("EspressoDCNEspressoOverfeatDetector")}
	})
	return _EspressoDCNEspressoOverfeatDetectorClass
}

// GetEspressoDCNEspressoOverfeatDetectorClass returns the class object for EspressoDCNEspressoOverfeatDetector.
func GetEspressoDCNEspressoOverfeatDetectorClass() EspressoDCNEspressoOverfeatDetectorClass {
	return getEspressoDCNEspressoOverfeatDetectorClass()
}

type EspressoDCNEspressoOverfeatDetectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoDCNEspressoOverfeatDetectorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDCNEspressoOverfeatDetectorClass) Alloc() EspressoDCNEspressoOverfeatDetector {
	rv := objc.Send[EspressoDCNEspressoOverfeatDetector](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoDCNEspressoOverfeatDetector.CommonInit]
//   - [EspressoDCNEspressoOverfeatDetector.CompareObjectWithObjectError]
//   - [EspressoDCNEspressoOverfeatDetector.ComputeBBoxUsingProbBoxAndScalefactorPadXPadY]
//   - [EspressoDCNEspressoOverfeatDetector.ConfidenceThreshold]
//   - [EspressoDCNEspressoOverfeatDetector.SetConfidenceThreshold]
//   - [EspressoDCNEspressoOverfeatDetector.Enet]
//   - [EspressoDCNEspressoOverfeatDetector.SetEnet]
//   - [EspressoDCNEspressoOverfeatDetector.FillFaceList]
//   - [EspressoDCNEspressoOverfeatDetector.GetDescription]
//   - [EspressoDCNEspressoOverfeatDetector.GetFacesFromNetworkResultOriginalWidthOriginalHeight]
//   - [EspressoDCNEspressoOverfeatDetector.MergeFaceList]
//   - [EspressoDCNEspressoOverfeatDetector.MinBoundingBoxThreshold]
//   - [EspressoDCNEspressoOverfeatDetector.SetMinBoundingBoxThreshold]
//   - [EspressoDCNEspressoOverfeatDetector.InitWithNetwork]
//   - [EspressoDCNEspressoOverfeatDetector.InitWithOptions]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector
type EspressoDCNEspressoOverfeatDetector struct {
	objectivec.Object
}

// EspressoDCNEspressoOverfeatDetectorFromID constructs a [EspressoDCNEspressoOverfeatDetector] from an objc.ID.
func EspressoDCNEspressoOverfeatDetectorFromID(id objc.ID) EspressoDCNEspressoOverfeatDetector {
	return EspressoDCNEspressoOverfeatDetector{objectivec.Object{ID: id}}
}

// Ensure EspressoDCNEspressoOverfeatDetector implements IEspressoDCNEspressoOverfeatDetector.
var _ IEspressoDCNEspressoOverfeatDetector = EspressoDCNEspressoOverfeatDetector{}

// An interface definition for the [EspressoDCNEspressoOverfeatDetector] class.
//
// # Methods
//
//   - [IEspressoDCNEspressoOverfeatDetector.CommonInit]
//   - [IEspressoDCNEspressoOverfeatDetector.CompareObjectWithObjectError]
//   - [IEspressoDCNEspressoOverfeatDetector.ComputeBBoxUsingProbBoxAndScalefactorPadXPadY]
//   - [IEspressoDCNEspressoOverfeatDetector.ConfidenceThreshold]
//   - [IEspressoDCNEspressoOverfeatDetector.SetConfidenceThreshold]
//   - [IEspressoDCNEspressoOverfeatDetector.Enet]
//   - [IEspressoDCNEspressoOverfeatDetector.SetEnet]
//   - [IEspressoDCNEspressoOverfeatDetector.FillFaceList]
//   - [IEspressoDCNEspressoOverfeatDetector.GetDescription]
//   - [IEspressoDCNEspressoOverfeatDetector.GetFacesFromNetworkResultOriginalWidthOriginalHeight]
//   - [IEspressoDCNEspressoOverfeatDetector.MergeFaceList]
//   - [IEspressoDCNEspressoOverfeatDetector.MinBoundingBoxThreshold]
//   - [IEspressoDCNEspressoOverfeatDetector.SetMinBoundingBoxThreshold]
//   - [IEspressoDCNEspressoOverfeatDetector.InitWithNetwork]
//   - [IEspressoDCNEspressoOverfeatDetector.InitWithOptions]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector
type IEspressoDCNEspressoOverfeatDetector interface {
	objectivec.IObject

	// Topic: Methods

	CommonInit()
	CompareObjectWithObjectError(object objectivec.IObject, object2 objectivec.IObject) (float64, error)
	ComputeBBoxUsingProbBoxAndScalefactorPadXPadY(prob objectivec.IObject, box objectivec.IObject, scalefactor float32, x float32, y float32)
	ConfidenceThreshold() float64
	SetConfidenceThreshold(value float64)
	Enet() IEspressoFDOverfeatNetwork
	SetEnet(value IEspressoFDOverfeatNetwork)
	FillFaceList()
	GetDescription() objectivec.IObject
	GetFacesFromNetworkResultOriginalWidthOriginalHeight(width float32, height float32) objectivec.IObject
	MergeFaceList()
	MinBoundingBoxThreshold() float64
	SetMinBoundingBoxThreshold(value float64)
	InitWithNetwork(network objectivec.IObject) EspressoDCNEspressoOverfeatDetector
	InitWithOptions(options objectivec.IObject) EspressoDCNEspressoOverfeatDetector
}

// Init initializes the instance.
func (e EspressoDCNEspressoOverfeatDetector) Init() EspressoDCNEspressoOverfeatDetector {
	rv := objc.Send[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDCNEspressoOverfeatDetector) Autorelease() EspressoDCNEspressoOverfeatDetector {
	rv := objc.Send[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDCNEspressoOverfeatDetector creates a new EspressoDCNEspressoOverfeatDetector instance.
func NewEspressoDCNEspressoOverfeatDetector() EspressoDCNEspressoOverfeatDetector {
	class := getEspressoDCNEspressoOverfeatDetectorClass()
	rv := objc.Send[EspressoDCNEspressoOverfeatDetector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/initWithNetwork:
func NewEspressoDCNEspressoOverfeatDetectorWithNetwork(network objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	instance := getEspressoDCNEspressoOverfeatDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return EspressoDCNEspressoOverfeatDetectorFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/initWithOptions:
func NewEspressoDCNEspressoOverfeatDetectorWithOptions(options objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	instance := getEspressoDCNEspressoOverfeatDetectorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return EspressoDCNEspressoOverfeatDetectorFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/commonInit
func (e EspressoDCNEspressoOverfeatDetector) CommonInit() {
	objc.Send[objc.ID](e.ID, objc.Sel("commonInit"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/compareObject:withObject:error:
func (e EspressoDCNEspressoOverfeatDetector) CompareObjectWithObjectError(object objectivec.IObject, object2 objectivec.IObject) (float64, error) {
	var errorPtr objc.ID
	rv := objc.Send[float64](e.ID, objc.Sel("compareObject:withObject:error:"), object, object2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/computeBBoxUsingProb:box:andScalefactor:padX:padY:
func (e EspressoDCNEspressoOverfeatDetector) ComputeBBoxUsingProbBoxAndScalefactorPadXPadY(prob objectivec.IObject, box objectivec.IObject, scalefactor float32, x float32, y float32) {
	objc.Send[objc.ID](e.ID, objc.Sel("computeBBoxUsingProb:box:andScalefactor:padX:padY:"), prob, box, scalefactor, x, y)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/fillFaceList
func (e EspressoDCNEspressoOverfeatDetector) FillFaceList() {
	objc.Send[objc.ID](e.ID, objc.Sel("fillFaceList"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/getDescription
func (e EspressoDCNEspressoOverfeatDetector) GetDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getDescription"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/getFacesFromNetworkResultOriginalWidth:originalHeight:
func (e EspressoDCNEspressoOverfeatDetector) GetFacesFromNetworkResultOriginalWidthOriginalHeight(width float32, height float32) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getFacesFromNetworkResultOriginalWidth:originalHeight:"), width, height)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/mergeFaceList
func (e EspressoDCNEspressoOverfeatDetector) MergeFaceList() {
	objc.Send[objc.ID](e.ID, objc.Sel("mergeFaceList"))
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/initWithNetwork:
func (e EspressoDCNEspressoOverfeatDetector) InitWithNetwork(network objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	rv := objc.Send[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/initWithOptions:
func (e EspressoDCNEspressoOverfeatDetector) InitWithOptions(options objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	rv := objc.Send[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/confidenceThreshold
func (e EspressoDCNEspressoOverfeatDetector) ConfidenceThreshold() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("confidenceThreshold"))
	return rv
}
func (e EspressoDCNEspressoOverfeatDetector) SetConfidenceThreshold(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setConfidenceThreshold:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/enet
func (e EspressoDCNEspressoOverfeatDetector) Enet() IEspressoFDOverfeatNetwork {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("enet"))
	return EspressoFDOverfeatNetworkFromID(objc.ID(rv))
}
func (e EspressoDCNEspressoOverfeatDetector) SetEnet(value IEspressoFDOverfeatNetwork) {
	objc.Send[struct{}](e.ID, objc.Sel("setEnet:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDCNEspressoOverfeatDetector/minBoundingBoxThreshold
func (e EspressoDCNEspressoOverfeatDetector) MinBoundingBoxThreshold() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("minBoundingBoxThreshold"))
	return rv
}
func (e EspressoDCNEspressoOverfeatDetector) SetMinBoundingBoxThreshold(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setMinBoundingBoxThreshold:"), value)
}
