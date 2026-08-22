// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFunctionalUnitPositiveTransparency] class.
var (
	_ICScannerFunctionalUnitPositiveTransparencyClass     ICScannerFunctionalUnitPositiveTransparencyClass
	_ICScannerFunctionalUnitPositiveTransparencyClassOnce sync.Once
)

func getICScannerFunctionalUnitPositiveTransparencyClass() ICScannerFunctionalUnitPositiveTransparencyClass {
	_ICScannerFunctionalUnitPositiveTransparencyClassOnce.Do(func() {
		_ICScannerFunctionalUnitPositiveTransparencyClass = ICScannerFunctionalUnitPositiveTransparencyClass{class: objc.GetClass("ICScannerFunctionalUnitPositiveTransparency")}
	})
	return _ICScannerFunctionalUnitPositiveTransparencyClass
}

// GetICScannerFunctionalUnitPositiveTransparencyClass returns the class object for ICScannerFunctionalUnitPositiveTransparency.
func GetICScannerFunctionalUnitPositiveTransparencyClass() ICScannerFunctionalUnitPositiveTransparencyClass {
	return getICScannerFunctionalUnitPositiveTransparencyClass()
}

type ICScannerFunctionalUnitPositiveTransparencyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFunctionalUnitPositiveTransparencyClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFunctionalUnitPositiveTransparencyClass) Alloc() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the transparency unit for scanning positives on
// the scanner.
//
// # Instance Properties
//
//   - [ICScannerFunctionalUnitPositiveTransparency.DocumentSize]
//   - [ICScannerFunctionalUnitPositiveTransparency.DocumentType]
//   - [ICScannerFunctionalUnitPositiveTransparency.SetDocumentType]
//   - [ICScannerFunctionalUnitPositiveTransparency.SupportedDocumentTypes]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitPositiveTransparency
type ICScannerFunctionalUnitPositiveTransparency struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitPositiveTransparencyFromID constructs a [ICScannerFunctionalUnitPositiveTransparency] from an objc.ID.
//
// An object that represents the transparency unit for scanning positives on
// the scanner.
func ICScannerFunctionalUnitPositiveTransparencyFromID(id objc.ID) ICScannerFunctionalUnitPositiveTransparency {
	return ICScannerFunctionalUnitPositiveTransparency{ICScannerFunctionalUnit: ICScannerFunctionalUnitFromID(id)}
}

// NOTE: ICScannerFunctionalUnitPositiveTransparency adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFunctionalUnitPositiveTransparency] class.
//
// # Instance Properties
//
//   - [IICScannerFunctionalUnitPositiveTransparency.DocumentSize]
//   - [IICScannerFunctionalUnitPositiveTransparency.DocumentType]
//   - [IICScannerFunctionalUnitPositiveTransparency.SetDocumentType]
//   - [IICScannerFunctionalUnitPositiveTransparency.SupportedDocumentTypes]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitPositiveTransparency
type IICScannerFunctionalUnitPositiveTransparency interface {
	IICScannerFunctionalUnit

	// Topic: Instance Properties

	DocumentSize() corefoundation.CGSize
	DocumentType() ICScannerDocumentType
	SetDocumentType(value ICScannerDocumentType)
	SupportedDocumentTypes() foundation.NSIndexSet
}

// Init initializes the instance.
func (s ICScannerFunctionalUnitPositiveTransparency) Init() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFunctionalUnitPositiveTransparency) Autorelease() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitPositiveTransparency creates a new ICScannerFunctionalUnitPositiveTransparency instance.
func NewICScannerFunctionalUnitPositiveTransparency() ICScannerFunctionalUnitPositiveTransparency {
	class := getICScannerFunctionalUnitPositiveTransparencyClass()
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitPositiveTransparency/documentSize
func (s ICScannerFunctionalUnitPositiveTransparency) DocumentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("documentSize"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitPositiveTransparency/documentType
func (s ICScannerFunctionalUnitPositiveTransparency) DocumentType() ICScannerDocumentType {
	rv := objc.Send[ICScannerDocumentType](s.ID, objc.Sel("documentType"))
	return ICScannerDocumentType(rv)
}
func (s ICScannerFunctionalUnitPositiveTransparency) SetDocumentType(value ICScannerDocumentType) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentType:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitPositiveTransparency/supportedDocumentTypes
func (s ICScannerFunctionalUnitPositiveTransparency) SupportedDocumentTypes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedDocumentTypes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
