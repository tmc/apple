// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNYOLOLossDescriptor] class.
var (
	_MPSCNNYOLOLossDescriptorClass     MPSCNNYOLOLossDescriptorClass
	_MPSCNNYOLOLossDescriptorClassOnce sync.Once
)

func getMPSCNNYOLOLossDescriptorClass() MPSCNNYOLOLossDescriptorClass {
	_MPSCNNYOLOLossDescriptorClassOnce.Do(func() {
		_MPSCNNYOLOLossDescriptorClass = MPSCNNYOLOLossDescriptorClass{class: objc.GetClass("MPSCNNYOLOLossDescriptor")}
	})
	return _MPSCNNYOLOLossDescriptorClass
}

// GetMPSCNNYOLOLossDescriptorClass returns the class object for MPSCNNYOLOLossDescriptor.
func GetMPSCNNYOLOLossDescriptorClass() MPSCNNYOLOLossDescriptorClass {
	return getMPSCNNYOLOLossDescriptorClass()
}

type MPSCNNYOLOLossDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNYOLOLossDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNYOLOLossDescriptorClass) Alloc() MPSCNNYOLOLossDescriptor {
	rv := objc.Send[MPSCNNYOLOLossDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies properties used by a YOLO loss kernel.
//
// # Instance Properties
//
//   - [MPSCNNYOLOLossDescriptor.AnchorBoxes]
//   - [MPSCNNYOLOLossDescriptor.SetAnchorBoxes]
//   - [MPSCNNYOLOLossDescriptor.ClassesLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.SetClassesLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.ConfidenceLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.SetConfidenceLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.MaxIOUForObjectAbsence]
//   - [MPSCNNYOLOLossDescriptor.SetMaxIOUForObjectAbsence]
//   - [MPSCNNYOLOLossDescriptor.MinIOUForObjectPresence]
//   - [MPSCNNYOLOLossDescriptor.SetMinIOUForObjectPresence]
//   - [MPSCNNYOLOLossDescriptor.NumberOfAnchorBoxes]
//   - [MPSCNNYOLOLossDescriptor.SetNumberOfAnchorBoxes]
//   - [MPSCNNYOLOLossDescriptor.ReduceAcrossBatch]
//   - [MPSCNNYOLOLossDescriptor.SetReduceAcrossBatch]
//   - [MPSCNNYOLOLossDescriptor.ReductionType]
//   - [MPSCNNYOLOLossDescriptor.SetReductionType]
//   - [MPSCNNYOLOLossDescriptor.Rescore]
//   - [MPSCNNYOLOLossDescriptor.SetRescore]
//   - [MPSCNNYOLOLossDescriptor.ScaleClass]
//   - [MPSCNNYOLOLossDescriptor.SetScaleClass]
//   - [MPSCNNYOLOLossDescriptor.ScaleNoObject]
//   - [MPSCNNYOLOLossDescriptor.SetScaleNoObject]
//   - [MPSCNNYOLOLossDescriptor.ScaleObject]
//   - [MPSCNNYOLOLossDescriptor.SetScaleObject]
//   - [MPSCNNYOLOLossDescriptor.ScaleWH]
//   - [MPSCNNYOLOLossDescriptor.SetScaleWH]
//   - [MPSCNNYOLOLossDescriptor.ScaleXY]
//   - [MPSCNNYOLOLossDescriptor.SetScaleXY]
//   - [MPSCNNYOLOLossDescriptor.WHLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.SetWHLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.XYLossDescriptor]
//   - [MPSCNNYOLOLossDescriptor.SetXYLossDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor
type MPSCNNYOLOLossDescriptor struct {
	objectivec.Object
}

// MPSCNNYOLOLossDescriptorFromID constructs a [MPSCNNYOLOLossDescriptor] from an objc.ID.
//
// An object that specifies properties used by a YOLO loss kernel.
func MPSCNNYOLOLossDescriptorFromID(id objc.ID) MPSCNNYOLOLossDescriptor {
	return MPSCNNYOLOLossDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSCNNYOLOLossDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNYOLOLossDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSCNNYOLOLossDescriptor.AnchorBoxes]
//   - [IMPSCNNYOLOLossDescriptor.SetAnchorBoxes]
//   - [IMPSCNNYOLOLossDescriptor.ClassesLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.SetClassesLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.ConfidenceLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.SetConfidenceLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.MaxIOUForObjectAbsence]
//   - [IMPSCNNYOLOLossDescriptor.SetMaxIOUForObjectAbsence]
//   - [IMPSCNNYOLOLossDescriptor.MinIOUForObjectPresence]
//   - [IMPSCNNYOLOLossDescriptor.SetMinIOUForObjectPresence]
//   - [IMPSCNNYOLOLossDescriptor.NumberOfAnchorBoxes]
//   - [IMPSCNNYOLOLossDescriptor.SetNumberOfAnchorBoxes]
//   - [IMPSCNNYOLOLossDescriptor.ReduceAcrossBatch]
//   - [IMPSCNNYOLOLossDescriptor.SetReduceAcrossBatch]
//   - [IMPSCNNYOLOLossDescriptor.ReductionType]
//   - [IMPSCNNYOLOLossDescriptor.SetReductionType]
//   - [IMPSCNNYOLOLossDescriptor.Rescore]
//   - [IMPSCNNYOLOLossDescriptor.SetRescore]
//   - [IMPSCNNYOLOLossDescriptor.ScaleClass]
//   - [IMPSCNNYOLOLossDescriptor.SetScaleClass]
//   - [IMPSCNNYOLOLossDescriptor.ScaleNoObject]
//   - [IMPSCNNYOLOLossDescriptor.SetScaleNoObject]
//   - [IMPSCNNYOLOLossDescriptor.ScaleObject]
//   - [IMPSCNNYOLOLossDescriptor.SetScaleObject]
//   - [IMPSCNNYOLOLossDescriptor.ScaleWH]
//   - [IMPSCNNYOLOLossDescriptor.SetScaleWH]
//   - [IMPSCNNYOLOLossDescriptor.ScaleXY]
//   - [IMPSCNNYOLOLossDescriptor.SetScaleXY]
//   - [IMPSCNNYOLOLossDescriptor.WHLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.SetWHLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.XYLossDescriptor]
//   - [IMPSCNNYOLOLossDescriptor.SetXYLossDescriptor]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor
type IMPSCNNYOLOLossDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	AnchorBoxes() foundation.NSData
	SetAnchorBoxes(value foundation.NSData)
	ClassesLossDescriptor() IMPSCNNLossDescriptor
	SetClassesLossDescriptor(value IMPSCNNLossDescriptor)
	ConfidenceLossDescriptor() IMPSCNNLossDescriptor
	SetConfidenceLossDescriptor(value IMPSCNNLossDescriptor)
	MaxIOUForObjectAbsence() float32
	SetMaxIOUForObjectAbsence(value float32)
	MinIOUForObjectPresence() float32
	SetMinIOUForObjectPresence(value float32)
	NumberOfAnchorBoxes() uint
	SetNumberOfAnchorBoxes(value uint)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	ReductionType() MPSCNNReductionType
	SetReductionType(value MPSCNNReductionType)
	Rescore() bool
	SetRescore(value bool)
	ScaleClass() float32
	SetScaleClass(value float32)
	ScaleNoObject() float32
	SetScaleNoObject(value float32)
	ScaleObject() float32
	SetScaleObject(value float32)
	ScaleWH() float32
	SetScaleWH(value float32)
	ScaleXY() float32
	SetScaleXY(value float32)
	WHLossDescriptor() IMPSCNNLossDescriptor
	SetWHLossDescriptor(value IMPSCNNLossDescriptor)
	XYLossDescriptor() IMPSCNNLossDescriptor
	SetXYLossDescriptor(value IMPSCNNLossDescriptor)
}

// Init initializes the instance.
func (c MPSCNNYOLOLossDescriptor) Init() MPSCNNYOLOLossDescriptor {
	rv := objc.Send[MPSCNNYOLOLossDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNYOLOLossDescriptor) Autorelease() MPSCNNYOLOLossDescriptor {
	rv := objc.Send[MPSCNNYOLOLossDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNYOLOLossDescriptor creates a new MPSCNNYOLOLossDescriptor instance.
func NewMPSCNNYOLOLossDescriptor() MPSCNNYOLOLossDescriptor {
	class := getMPSCNNYOLOLossDescriptorClass()
	rv := objc.Send[MPSCNNYOLOLossDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/cnnLossDescriptor(withXYLossType:whLossType:confidenceLossType:classesLossType:reductionType:anchorBoxes:numberOfAnchorBoxes:)
func (_MPSCNNYOLOLossDescriptorClass MPSCNNYOLOLossDescriptorClass) CnnLossDescriptorWithXYLossTypeWHLossTypeConfidenceLossTypeClassesLossTypeReductionTypeAnchorBoxesNumberOfAnchorBoxes(XYLossType MPSCNNLossType, WHLossType MPSCNNLossType, confidenceLossType MPSCNNLossType, classesLossType MPSCNNLossType, reductionType MPSCNNReductionType, anchorBoxes foundation.NSData, numberOfAnchorBoxes uint) MPSCNNYOLOLossDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MPSCNNYOLOLossDescriptorClass.class), objc.Sel("cnnLossDescriptorWithXYLossType:WHLossType:confidenceLossType:classesLossType:reductionType:anchorBoxes:numberOfAnchorBoxes:"), XYLossType, WHLossType, confidenceLossType, classesLossType, reductionType, anchorBoxes, numberOfAnchorBoxes)
	return MPSCNNYOLOLossDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/anchorBoxes
func (c MPSCNNYOLOLossDescriptor) AnchorBoxes() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("anchorBoxes"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c MPSCNNYOLOLossDescriptor) SetAnchorBoxes(value foundation.NSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setAnchorBoxes:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/classesLossDescriptor
func (c MPSCNNYOLOLossDescriptor) ClassesLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("classesLossDescriptor"))
	return MPSCNNLossDescriptorFromID(objc.ID(rv))
}
func (c MPSCNNYOLOLossDescriptor) SetClassesLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setClassesLossDescriptor:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/confidenceLossDescriptor
func (c MPSCNNYOLOLossDescriptor) ConfidenceLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("confidenceLossDescriptor"))
	return MPSCNNLossDescriptorFromID(objc.ID(rv))
}
func (c MPSCNNYOLOLossDescriptor) SetConfidenceLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfidenceLossDescriptor:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/maxIOUForObjectAbsence
func (c MPSCNNYOLOLossDescriptor) MaxIOUForObjectAbsence() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("maxIOUForObjectAbsence"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetMaxIOUForObjectAbsence(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxIOUForObjectAbsence:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/minIOUForObjectPresence
func (c MPSCNNYOLOLossDescriptor) MinIOUForObjectPresence() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("minIOUForObjectPresence"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetMinIOUForObjectPresence(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinIOUForObjectPresence:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/numberOfAnchorBoxes
func (c MPSCNNYOLOLossDescriptor) NumberOfAnchorBoxes() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("numberOfAnchorBoxes"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetNumberOfAnchorBoxes(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setNumberOfAnchorBoxes:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/reduceAcrossBatch
func (c MPSCNNYOLOLossDescriptor) ReduceAcrossBatch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetReduceAcrossBatch(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setReduceAcrossBatch:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/reductionType
func (c MPSCNNYOLOLossDescriptor) ReductionType() MPSCNNReductionType {
	rv := objc.Send[MPSCNNReductionType](c.ID, objc.Sel("reductionType"))
	return MPSCNNReductionType(rv)
}
func (c MPSCNNYOLOLossDescriptor) SetReductionType(value MPSCNNReductionType) {
	objc.Send[struct{}](c.ID, objc.Sel("setReductionType:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/rescore
func (c MPSCNNYOLOLossDescriptor) Rescore() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("rescore"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetRescore(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setRescore:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/scaleClass
func (c MPSCNNYOLOLossDescriptor) ScaleClass() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleClass"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetScaleClass(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setScaleClass:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/scaleNoObject
func (c MPSCNNYOLOLossDescriptor) ScaleNoObject() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleNoObject"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetScaleNoObject(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setScaleNoObject:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/scaleObject
func (c MPSCNNYOLOLossDescriptor) ScaleObject() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleObject"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetScaleObject(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setScaleObject:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/scaleWH
func (c MPSCNNYOLOLossDescriptor) ScaleWH() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleWH"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetScaleWH(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setScaleWH:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/scaleXY
func (c MPSCNNYOLOLossDescriptor) ScaleXY() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("scaleXY"))
	return rv
}
func (c MPSCNNYOLOLossDescriptor) SetScaleXY(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setScaleXY:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/whLossDescriptor
func (c MPSCNNYOLOLossDescriptor) WHLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("WHLossDescriptor"))
	return MPSCNNLossDescriptorFromID(objc.ID(rv))
}
func (c MPSCNNYOLOLossDescriptor) SetWHLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setWHLossDescriptor:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossDescriptor/xyLossDescriptor
func (c MPSCNNYOLOLossDescriptor) XYLossDescriptor() IMPSCNNLossDescriptor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("XYLossDescriptor"))
	return MPSCNNLossDescriptorFromID(objc.ID(rv))
}
func (c MPSCNNYOLOLossDescriptor) SetXYLossDescriptor(value IMPSCNNLossDescriptor) {
	objc.Send[struct{}](c.ID, objc.Sel("setXYLossDescriptor:"), value)
}
