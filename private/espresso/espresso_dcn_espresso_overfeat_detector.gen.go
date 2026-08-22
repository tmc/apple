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
	rv := objc.SendIfResponds[EspressoDCNEspressoOverfeatDetector](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [EspressoDCNEspressoOverfeatDetector.CommonInit]
//   - [EspressoDCNEspressoOverfeatDetector.CompareObjectWithObjectError]
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
type IEspressoDCNEspressoOverfeatDetector interface {
	objectivec.IObject

	// Topic: Methods

	CommonInit()
	CompareObjectWithObjectError(object objectivec.IObject, object2 objectivec.IObject) (float64, error)
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
	rv := objc.SendIfResponds[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDCNEspressoOverfeatDetector) Autorelease() EspressoDCNEspressoOverfeatDetector {
	rv := objc.SendIfResponds[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDCNEspressoOverfeatDetector creates a new EspressoDCNEspressoOverfeatDetector instance.
func NewEspressoDCNEspressoOverfeatDetector() EspressoDCNEspressoOverfeatDetector {
	class := getEspressoDCNEspressoOverfeatDetectorClass()
	rv := objc.SendIfResponds[EspressoDCNEspressoOverfeatDetector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewEspressoDCNEspressoOverfeatDetectorWithNetwork(network objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	instance := getEspressoDCNEspressoOverfeatDetectorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return EspressoDCNEspressoOverfeatDetectorFromID(rv)
}

func NewEspressoDCNEspressoOverfeatDetectorWithOptions(options objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	instance := getEspressoDCNEspressoOverfeatDetectorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return EspressoDCNEspressoOverfeatDetectorFromID(rv)
}

func (e EspressoDCNEspressoOverfeatDetector) CommonInit() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("commonInit"))
}
func (e EspressoDCNEspressoOverfeatDetector) CompareObjectWithObjectError(object objectivec.IObject, object2 objectivec.IObject) (float64, error) {
	var errorPtr objc.ID
	rv := objc.Send[float64](e.ID, objc.Sel("compareObject:withObject:error:"), object, object2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (e EspressoDCNEspressoOverfeatDetector) FillFaceList() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("fillFaceList"))
}
func (e EspressoDCNEspressoOverfeatDetector) GetDescription() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("getDescription"))
	return objectivec.Object{ID: rv}
}
func (e EspressoDCNEspressoOverfeatDetector) GetFacesFromNetworkResultOriginalWidthOriginalHeight(width float32, height float32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("getFacesFromNetworkResultOriginalWidth:originalHeight:"), width, height)
	return objectivec.Object{ID: rv}
}
func (e EspressoDCNEspressoOverfeatDetector) MergeFaceList() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("mergeFaceList"))
}
func (e EspressoDCNEspressoOverfeatDetector) InitWithNetwork(network objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	rv := objc.SendIfResponds[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}
func (e EspressoDCNEspressoOverfeatDetector) InitWithOptions(options objectivec.IObject) EspressoDCNEspressoOverfeatDetector {
	rv := objc.SendIfResponds[EspressoDCNEspressoOverfeatDetector](e.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

func (e EspressoDCNEspressoOverfeatDetector) ConfidenceThreshold() float64 {
	rv := objc.SendIfResponds[float64](e.ID, objc.Sel("confidenceThreshold"))
	return rv
}
func (e EspressoDCNEspressoOverfeatDetector) SetConfidenceThreshold(value float64) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setConfidenceThreshold:"), value)
}
func (e EspressoDCNEspressoOverfeatDetector) Enet() IEspressoFDOverfeatNetwork {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("enet"))
	return EspressoFDOverfeatNetworkFromID(objc.ID(rv))
}
func (e EspressoDCNEspressoOverfeatDetector) SetEnet(value IEspressoFDOverfeatNetwork) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setEnet:"), value)
}
func (e EspressoDCNEspressoOverfeatDetector) MinBoundingBoxThreshold() float64 {
	rv := objc.SendIfResponds[float64](e.ID, objc.Sel("minBoundingBoxThreshold"))
	return rv
}
func (e EspressoDCNEspressoOverfeatDetector) SetMinBoundingBoxThreshold(value float64) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setMinBoundingBoxThreshold:"), value)
}
