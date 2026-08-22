// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICScannerFunctionalUnit] class.
var (
	_ICScannerFunctionalUnitClass     ICScannerFunctionalUnitClass
	_ICScannerFunctionalUnitClassOnce sync.Once
)

func getICScannerFunctionalUnitClass() ICScannerFunctionalUnitClass {
	_ICScannerFunctionalUnitClassOnce.Do(func() {
		_ICScannerFunctionalUnitClass = ICScannerFunctionalUnitClass{class: objc.GetClass("ICScannerFunctionalUnit")}
	})
	return _ICScannerFunctionalUnitClass
}

// GetICScannerFunctionalUnitClass returns the class object for ICScannerFunctionalUnit.
func GetICScannerFunctionalUnitClass() ICScannerFunctionalUnitClass {
	return getICScannerFunctionalUnitClass()
}

type ICScannerFunctionalUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFunctionalUnitClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFunctionalUnitClass) Alloc() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that represents a scanner functional unit.
//
// # Overview
//
// The ImageCaptureCore framework defines four concrete subclasses of
// functional units:
//
// - [ICScannerFunctionalUnitDocumentFeeder] -
// [ICScannerFunctionalUnitFlatbed] -
// [ICScannerFunctionalUnitPositiveTransparency] -
// [ICScannerFunctionalUnitNegativeTransparency]
//
// [ICScannerDevice] creates instances of these subclasses.
//
// # Instance Properties
//
//   - [ICScannerFunctionalUnit.AcceptsThresholdForBlackAndWhiteScanning]
//   - [ICScannerFunctionalUnit.BitDepth]
//   - [ICScannerFunctionalUnit.SetBitDepth]
//   - [ICScannerFunctionalUnit.CanPerformOverviewScan]
//   - [ICScannerFunctionalUnit.DefaultThresholdForBlackAndWhiteScanning]
//   - [ICScannerFunctionalUnit.MeasurementUnit]
//   - [ICScannerFunctionalUnit.SetMeasurementUnit]
//   - [ICScannerFunctionalUnit.NativeXResolution]
//   - [ICScannerFunctionalUnit.NativeYResolution]
//   - [ICScannerFunctionalUnit.OverviewImage]
//   - [ICScannerFunctionalUnit.OverviewResolution]
//   - [ICScannerFunctionalUnit.SetOverviewResolution]
//   - [ICScannerFunctionalUnit.OverviewScanInProgress]
//   - [ICScannerFunctionalUnit.PhysicalSize]
//   - [ICScannerFunctionalUnit.PixelDataType]
//   - [ICScannerFunctionalUnit.SetPixelDataType]
//   - [ICScannerFunctionalUnit.PreferredResolutions]
//   - [ICScannerFunctionalUnit.PreferredScaleFactors]
//   - [ICScannerFunctionalUnit.Resolution]
//   - [ICScannerFunctionalUnit.SetResolution]
//   - [ICScannerFunctionalUnit.ScaleFactor]
//   - [ICScannerFunctionalUnit.SetScaleFactor]
//   - [ICScannerFunctionalUnit.ScanArea]
//   - [ICScannerFunctionalUnit.SetScanArea]
//   - [ICScannerFunctionalUnit.ScanAreaOrientation]
//   - [ICScannerFunctionalUnit.SetScanAreaOrientation]
//   - [ICScannerFunctionalUnit.ScanInProgress]
//   - [ICScannerFunctionalUnit.ScanProgressPercentDone]
//   - [ICScannerFunctionalUnit.State]
//   - [ICScannerFunctionalUnit.SupportedBitDepths]
//   - [ICScannerFunctionalUnit.SupportedMeasurementUnits]
//   - [ICScannerFunctionalUnit.SupportedResolutions]
//   - [ICScannerFunctionalUnit.SupportedScaleFactors]
//   - [ICScannerFunctionalUnit.Templates]
//   - [ICScannerFunctionalUnit.ThresholdForBlackAndWhiteScanning]
//   - [ICScannerFunctionalUnit.SetThresholdForBlackAndWhiteScanning]
//   - [ICScannerFunctionalUnit.Type]
//   - [ICScannerFunctionalUnit.UsesThresholdForBlackAndWhiteScanning]
//   - [ICScannerFunctionalUnit.SetUsesThresholdForBlackAndWhiteScanning]
//   - [ICScannerFunctionalUnit.VendorFeatures]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit
type ICScannerFunctionalUnit struct {
	objectivec.Object
}

// ICScannerFunctionalUnitFromID constructs a [ICScannerFunctionalUnit] from an objc.ID.
//
// An abstract class that represents a scanner functional unit.
func ICScannerFunctionalUnitFromID(id objc.ID) ICScannerFunctionalUnit {
	return ICScannerFunctionalUnit{objectivec.Object{ID: id}}
}

// NOTE: ICScannerFunctionalUnit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFunctionalUnit] class.
//
// # Instance Properties
//
//   - [IICScannerFunctionalUnit.AcceptsThresholdForBlackAndWhiteScanning]
//   - [IICScannerFunctionalUnit.BitDepth]
//   - [IICScannerFunctionalUnit.SetBitDepth]
//   - [IICScannerFunctionalUnit.CanPerformOverviewScan]
//   - [IICScannerFunctionalUnit.DefaultThresholdForBlackAndWhiteScanning]
//   - [IICScannerFunctionalUnit.MeasurementUnit]
//   - [IICScannerFunctionalUnit.SetMeasurementUnit]
//   - [IICScannerFunctionalUnit.NativeXResolution]
//   - [IICScannerFunctionalUnit.NativeYResolution]
//   - [IICScannerFunctionalUnit.OverviewImage]
//   - [IICScannerFunctionalUnit.OverviewResolution]
//   - [IICScannerFunctionalUnit.SetOverviewResolution]
//   - [IICScannerFunctionalUnit.OverviewScanInProgress]
//   - [IICScannerFunctionalUnit.PhysicalSize]
//   - [IICScannerFunctionalUnit.PixelDataType]
//   - [IICScannerFunctionalUnit.SetPixelDataType]
//   - [IICScannerFunctionalUnit.PreferredResolutions]
//   - [IICScannerFunctionalUnit.PreferredScaleFactors]
//   - [IICScannerFunctionalUnit.Resolution]
//   - [IICScannerFunctionalUnit.SetResolution]
//   - [IICScannerFunctionalUnit.ScaleFactor]
//   - [IICScannerFunctionalUnit.SetScaleFactor]
//   - [IICScannerFunctionalUnit.ScanArea]
//   - [IICScannerFunctionalUnit.SetScanArea]
//   - [IICScannerFunctionalUnit.ScanAreaOrientation]
//   - [IICScannerFunctionalUnit.SetScanAreaOrientation]
//   - [IICScannerFunctionalUnit.ScanInProgress]
//   - [IICScannerFunctionalUnit.ScanProgressPercentDone]
//   - [IICScannerFunctionalUnit.State]
//   - [IICScannerFunctionalUnit.SupportedBitDepths]
//   - [IICScannerFunctionalUnit.SupportedMeasurementUnits]
//   - [IICScannerFunctionalUnit.SupportedResolutions]
//   - [IICScannerFunctionalUnit.SupportedScaleFactors]
//   - [IICScannerFunctionalUnit.Templates]
//   - [IICScannerFunctionalUnit.ThresholdForBlackAndWhiteScanning]
//   - [IICScannerFunctionalUnit.SetThresholdForBlackAndWhiteScanning]
//   - [IICScannerFunctionalUnit.Type]
//   - [IICScannerFunctionalUnit.UsesThresholdForBlackAndWhiteScanning]
//   - [IICScannerFunctionalUnit.SetUsesThresholdForBlackAndWhiteScanning]
//   - [IICScannerFunctionalUnit.VendorFeatures]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit
type IICScannerFunctionalUnit interface {
	objectivec.IObject

	// Topic: Instance Properties

	AcceptsThresholdForBlackAndWhiteScanning() bool
	BitDepth() ICScannerBitDepth
	SetBitDepth(value ICScannerBitDepth)
	CanPerformOverviewScan() bool
	DefaultThresholdForBlackAndWhiteScanning() byte
	MeasurementUnit() ICScannerMeasurementUnit
	SetMeasurementUnit(value ICScannerMeasurementUnit)
	NativeXResolution() uint
	NativeYResolution() uint
	OverviewImage() coregraphics.CGImageRef
	OverviewResolution() uint
	SetOverviewResolution(value uint)
	OverviewScanInProgress() bool
	PhysicalSize() corefoundation.CGSize
	PixelDataType() ICScannerPixelDataType
	SetPixelDataType(value ICScannerPixelDataType)
	PreferredResolutions() foundation.NSIndexSet
	PreferredScaleFactors() foundation.NSIndexSet
	Resolution() uint
	SetResolution(value uint)
	ScaleFactor() uint
	SetScaleFactor(value uint)
	ScanArea() corefoundation.CGRect
	SetScanArea(value corefoundation.CGRect)
	ScanAreaOrientation() ICEXIFOrientationType
	SetScanAreaOrientation(value ICEXIFOrientationType)
	ScanInProgress() bool
	ScanProgressPercentDone() float64
	State() ICScannerFunctionalUnitState
	SupportedBitDepths() foundation.NSIndexSet
	SupportedMeasurementUnits() foundation.NSIndexSet
	SupportedResolutions() foundation.NSIndexSet
	SupportedScaleFactors() foundation.NSIndexSet
	Templates() []ICScannerFeatureTemplate
	ThresholdForBlackAndWhiteScanning() byte
	SetThresholdForBlackAndWhiteScanning(value byte)
	Type() ICScannerFunctionalUnitType
	UsesThresholdForBlackAndWhiteScanning() bool
	SetUsesThresholdForBlackAndWhiteScanning(value bool)
	VendorFeatures() []ICScannerFeature
}

// Init initializes the instance.
func (s ICScannerFunctionalUnit) Init() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFunctionalUnit) Autorelease() ICScannerFunctionalUnit {
	rv := objc.Send[ICScannerFunctionalUnit](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnit creates a new ICScannerFunctionalUnit instance.
func NewICScannerFunctionalUnit() ICScannerFunctionalUnit {
	class := getICScannerFunctionalUnitClass()
	rv := objc.Send[ICScannerFunctionalUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/acceptsThresholdForBlackAndWhiteScanning
func (s ICScannerFunctionalUnit) AcceptsThresholdForBlackAndWhiteScanning() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("acceptsThresholdForBlackAndWhiteScanning"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/bitDepth
func (s ICScannerFunctionalUnit) BitDepth() ICScannerBitDepth {
	rv := objc.Send[ICScannerBitDepth](s.ID, objc.Sel("bitDepth"))
	return ICScannerBitDepth(rv)
}
func (s ICScannerFunctionalUnit) SetBitDepth(value ICScannerBitDepth) {
	objc.Send[struct{}](s.ID, objc.Sel("setBitDepth:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/canPerformOverviewScan
func (s ICScannerFunctionalUnit) CanPerformOverviewScan() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canPerformOverviewScan"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/defaultThresholdForBlackAndWhiteScanning
func (s ICScannerFunctionalUnit) DefaultThresholdForBlackAndWhiteScanning() byte {
	rv := objc.Send[byte](s.ID, objc.Sel("defaultThresholdForBlackAndWhiteScanning"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/measurementUnit
func (s ICScannerFunctionalUnit) MeasurementUnit() ICScannerMeasurementUnit {
	rv := objc.Send[ICScannerMeasurementUnit](s.ID, objc.Sel("measurementUnit"))
	return ICScannerMeasurementUnit(rv)
}
func (s ICScannerFunctionalUnit) SetMeasurementUnit(value ICScannerMeasurementUnit) {
	objc.Send[struct{}](s.ID, objc.Sel("setMeasurementUnit:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/nativeXResolution
func (s ICScannerFunctionalUnit) NativeXResolution() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("nativeXResolution"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/nativeYResolution
func (s ICScannerFunctionalUnit) NativeYResolution() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("nativeYResolution"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/overviewImage
func (s ICScannerFunctionalUnit) OverviewImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](s.ID, objc.Sel("overviewImage"))
	return coregraphics.CGImageRef(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/overviewResolution
func (s ICScannerFunctionalUnit) OverviewResolution() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("overviewResolution"))
	return rv
}
func (s ICScannerFunctionalUnit) SetOverviewResolution(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setOverviewResolution:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/overviewScanInProgress
func (s ICScannerFunctionalUnit) OverviewScanInProgress() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("overviewScanInProgress"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/physicalSize
func (s ICScannerFunctionalUnit) PhysicalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("physicalSize"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/pixelDataType
func (s ICScannerFunctionalUnit) PixelDataType() ICScannerPixelDataType {
	rv := objc.Send[ICScannerPixelDataType](s.ID, objc.Sel("pixelDataType"))
	return ICScannerPixelDataType(rv)
}
func (s ICScannerFunctionalUnit) SetPixelDataType(value ICScannerPixelDataType) {
	objc.Send[struct{}](s.ID, objc.Sel("setPixelDataType:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/preferredResolutions
func (s ICScannerFunctionalUnit) PreferredResolutions() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("preferredResolutions"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/preferredScaleFactors
func (s ICScannerFunctionalUnit) PreferredScaleFactors() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("preferredScaleFactors"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/resolution
func (s ICScannerFunctionalUnit) Resolution() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("resolution"))
	return rv
}
func (s ICScannerFunctionalUnit) SetResolution(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setResolution:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/scaleFactor
func (s ICScannerFunctionalUnit) ScaleFactor() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("scaleFactor"))
	return rv
}
func (s ICScannerFunctionalUnit) SetScaleFactor(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setScaleFactor:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/scanArea
func (s ICScannerFunctionalUnit) ScanArea() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("scanArea"))
	return corefoundation.CGRect(rv)
}
func (s ICScannerFunctionalUnit) SetScanArea(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setScanArea:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/scanAreaOrientation
func (s ICScannerFunctionalUnit) ScanAreaOrientation() ICEXIFOrientationType {
	rv := objc.Send[ICEXIFOrientationType](s.ID, objc.Sel("scanAreaOrientation"))
	return ICEXIFOrientationType(rv)
}
func (s ICScannerFunctionalUnit) SetScanAreaOrientation(value ICEXIFOrientationType) {
	objc.Send[struct{}](s.ID, objc.Sel("setScanAreaOrientation:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/scanInProgress
func (s ICScannerFunctionalUnit) ScanInProgress() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("scanInProgress"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/scanProgressPercentDone
func (s ICScannerFunctionalUnit) ScanProgressPercentDone() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("scanProgressPercentDone"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/state
func (s ICScannerFunctionalUnit) State() ICScannerFunctionalUnitState {
	rv := objc.Send[ICScannerFunctionalUnitState](s.ID, objc.Sel("state"))
	return ICScannerFunctionalUnitState(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/supportedBitDepths
func (s ICScannerFunctionalUnit) SupportedBitDepths() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedBitDepths"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/supportedMeasurementUnits
func (s ICScannerFunctionalUnit) SupportedMeasurementUnits() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedMeasurementUnits"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/supportedResolutions
func (s ICScannerFunctionalUnit) SupportedResolutions() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedResolutions"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/supportedScaleFactors
func (s ICScannerFunctionalUnit) SupportedScaleFactors() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedScaleFactors"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/templates
func (s ICScannerFunctionalUnit) Templates() []ICScannerFeatureTemplate {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("templates"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICScannerFeatureTemplate {
		return ICScannerFeatureTemplateFromID(id)
	})
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/thresholdForBlackAndWhiteScanning
func (s ICScannerFunctionalUnit) ThresholdForBlackAndWhiteScanning() byte {
	rv := objc.Send[byte](s.ID, objc.Sel("thresholdForBlackAndWhiteScanning"))
	return rv
}
func (s ICScannerFunctionalUnit) SetThresholdForBlackAndWhiteScanning(value byte) {
	objc.Send[struct{}](s.ID, objc.Sel("setThresholdForBlackAndWhiteScanning:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/type
func (s ICScannerFunctionalUnit) Type() ICScannerFunctionalUnitType {
	rv := objc.Send[ICScannerFunctionalUnitType](s.ID, objc.Sel("type"))
	return ICScannerFunctionalUnitType(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/usesThresholdForBlackAndWhiteScanning
func (s ICScannerFunctionalUnit) UsesThresholdForBlackAndWhiteScanning() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("usesThresholdForBlackAndWhiteScanning"))
	return rv
}
func (s ICScannerFunctionalUnit) SetUsesThresholdForBlackAndWhiteScanning(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setUsesThresholdForBlackAndWhiteScanning:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnit/vendorFeatures
func (s ICScannerFunctionalUnit) VendorFeatures() []ICScannerFeature {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("vendorFeatures"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICScannerFeature {
		return ICScannerFeatureFromID(id)
	})
}
