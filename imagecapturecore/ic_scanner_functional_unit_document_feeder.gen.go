// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFunctionalUnitDocumentFeeder] class.
var (
	_ICScannerFunctionalUnitDocumentFeederClass     ICScannerFunctionalUnitDocumentFeederClass
	_ICScannerFunctionalUnitDocumentFeederClassOnce sync.Once
)

func getICScannerFunctionalUnitDocumentFeederClass() ICScannerFunctionalUnitDocumentFeederClass {
	_ICScannerFunctionalUnitDocumentFeederClassOnce.Do(func() {
		_ICScannerFunctionalUnitDocumentFeederClass = ICScannerFunctionalUnitDocumentFeederClass{class: objc.GetClass("ICScannerFunctionalUnitDocumentFeeder")}
	})
	return _ICScannerFunctionalUnitDocumentFeederClass
}

// GetICScannerFunctionalUnitDocumentFeederClass returns the class object for ICScannerFunctionalUnitDocumentFeeder.
func GetICScannerFunctionalUnitDocumentFeederClass() ICScannerFunctionalUnitDocumentFeederClass {
	return getICScannerFunctionalUnitDocumentFeederClass()
}

type ICScannerFunctionalUnitDocumentFeederClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFunctionalUnitDocumentFeederClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFunctionalUnitDocumentFeederClass) Alloc() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the document feeder unit on a scanner.
//
// # Instance Properties
//
//   - [ICScannerFunctionalUnitDocumentFeeder.DocumentLoaded]
//   - [ICScannerFunctionalUnitDocumentFeeder.DocumentSize]
//   - [ICScannerFunctionalUnitDocumentFeeder.DocumentType]
//   - [ICScannerFunctionalUnitDocumentFeeder.SetDocumentType]
//   - [ICScannerFunctionalUnitDocumentFeeder.DuplexScanningEnabled]
//   - [ICScannerFunctionalUnitDocumentFeeder.SetDuplexScanningEnabled]
//   - [ICScannerFunctionalUnitDocumentFeeder.EvenPageOrientation]
//   - [ICScannerFunctionalUnitDocumentFeeder.SetEvenPageOrientation]
//   - [ICScannerFunctionalUnitDocumentFeeder.OddPageOrientation]
//   - [ICScannerFunctionalUnitDocumentFeeder.SetOddPageOrientation]
//   - [ICScannerFunctionalUnitDocumentFeeder.ReverseFeederPageOrder]
//   - [ICScannerFunctionalUnitDocumentFeeder.SupportedDocumentTypes]
//   - [ICScannerFunctionalUnitDocumentFeeder.SupportsDuplexScanning]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder
type ICScannerFunctionalUnitDocumentFeeder struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitDocumentFeederFromID constructs a [ICScannerFunctionalUnitDocumentFeeder] from an objc.ID.
//
// An object that represents the document feeder unit on a scanner.
func ICScannerFunctionalUnitDocumentFeederFromID(id objc.ID) ICScannerFunctionalUnitDocumentFeeder {
	return ICScannerFunctionalUnitDocumentFeeder{ICScannerFunctionalUnit: ICScannerFunctionalUnitFromID(id)}
}

// NOTE: ICScannerFunctionalUnitDocumentFeeder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFunctionalUnitDocumentFeeder] class.
//
// # Instance Properties
//
//   - [IICScannerFunctionalUnitDocumentFeeder.DocumentLoaded]
//   - [IICScannerFunctionalUnitDocumentFeeder.DocumentSize]
//   - [IICScannerFunctionalUnitDocumentFeeder.DocumentType]
//   - [IICScannerFunctionalUnitDocumentFeeder.SetDocumentType]
//   - [IICScannerFunctionalUnitDocumentFeeder.DuplexScanningEnabled]
//   - [IICScannerFunctionalUnitDocumentFeeder.SetDuplexScanningEnabled]
//   - [IICScannerFunctionalUnitDocumentFeeder.EvenPageOrientation]
//   - [IICScannerFunctionalUnitDocumentFeeder.SetEvenPageOrientation]
//   - [IICScannerFunctionalUnitDocumentFeeder.OddPageOrientation]
//   - [IICScannerFunctionalUnitDocumentFeeder.SetOddPageOrientation]
//   - [IICScannerFunctionalUnitDocumentFeeder.ReverseFeederPageOrder]
//   - [IICScannerFunctionalUnitDocumentFeeder.SupportedDocumentTypes]
//   - [IICScannerFunctionalUnitDocumentFeeder.SupportsDuplexScanning]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder
type IICScannerFunctionalUnitDocumentFeeder interface {
	IICScannerFunctionalUnit

	// Topic: Instance Properties

	DocumentLoaded() bool
	DocumentSize() corefoundation.CGSize
	DocumentType() ICScannerDocumentType
	SetDocumentType(value ICScannerDocumentType)
	DuplexScanningEnabled() bool
	SetDuplexScanningEnabled(value bool)
	EvenPageOrientation() ICEXIFOrientationType
	SetEvenPageOrientation(value ICEXIFOrientationType)
	OddPageOrientation() ICEXIFOrientationType
	SetOddPageOrientation(value ICEXIFOrientationType)
	ReverseFeederPageOrder() bool
	SupportedDocumentTypes() foundation.NSIndexSet
	SupportsDuplexScanning() bool
}

// Init initializes the instance.
func (s ICScannerFunctionalUnitDocumentFeeder) Init() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFunctionalUnitDocumentFeeder) Autorelease() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitDocumentFeeder creates a new ICScannerFunctionalUnitDocumentFeeder instance.
func NewICScannerFunctionalUnitDocumentFeeder() ICScannerFunctionalUnitDocumentFeeder {
	class := getICScannerFunctionalUnitDocumentFeederClass()
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/documentLoaded
func (s ICScannerFunctionalUnitDocumentFeeder) DocumentLoaded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("documentLoaded"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/documentSize
func (s ICScannerFunctionalUnitDocumentFeeder) DocumentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("documentSize"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/documentType
func (s ICScannerFunctionalUnitDocumentFeeder) DocumentType() ICScannerDocumentType {
	rv := objc.Send[ICScannerDocumentType](s.ID, objc.Sel("documentType"))
	return ICScannerDocumentType(rv)
}
func (s ICScannerFunctionalUnitDocumentFeeder) SetDocumentType(value ICScannerDocumentType) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentType:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/duplexScanningEnabled
func (s ICScannerFunctionalUnitDocumentFeeder) DuplexScanningEnabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("duplexScanningEnabled"))
	return rv
}
func (s ICScannerFunctionalUnitDocumentFeeder) SetDuplexScanningEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDuplexScanningEnabled:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/evenPageOrientation
func (s ICScannerFunctionalUnitDocumentFeeder) EvenPageOrientation() ICEXIFOrientationType {
	rv := objc.Send[ICEXIFOrientationType](s.ID, objc.Sel("evenPageOrientation"))
	return ICEXIFOrientationType(rv)
}
func (s ICScannerFunctionalUnitDocumentFeeder) SetEvenPageOrientation(value ICEXIFOrientationType) {
	objc.Send[struct{}](s.ID, objc.Sel("setEvenPageOrientation:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/oddPageOrientation
func (s ICScannerFunctionalUnitDocumentFeeder) OddPageOrientation() ICEXIFOrientationType {
	rv := objc.Send[ICEXIFOrientationType](s.ID, objc.Sel("oddPageOrientation"))
	return ICEXIFOrientationType(rv)
}
func (s ICScannerFunctionalUnitDocumentFeeder) SetOddPageOrientation(value ICEXIFOrientationType) {
	objc.Send[struct{}](s.ID, objc.Sel("setOddPageOrientation:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/reverseFeederPageOrder
func (s ICScannerFunctionalUnitDocumentFeeder) ReverseFeederPageOrder() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("reverseFeederPageOrder"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/supportedDocumentTypes
func (s ICScannerFunctionalUnitDocumentFeeder) SupportedDocumentTypes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedDocumentTypes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder/supportsDuplexScanning
func (s ICScannerFunctionalUnitDocumentFeeder) SupportsDuplexScanning() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("supportsDuplexScanning"))
	return rv
}
