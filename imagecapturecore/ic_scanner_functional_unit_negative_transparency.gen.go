// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFunctionalUnitNegativeTransparency] class.
var (
	_ICScannerFunctionalUnitNegativeTransparencyClass     ICScannerFunctionalUnitNegativeTransparencyClass
	_ICScannerFunctionalUnitNegativeTransparencyClassOnce sync.Once
)

func getICScannerFunctionalUnitNegativeTransparencyClass() ICScannerFunctionalUnitNegativeTransparencyClass {
	_ICScannerFunctionalUnitNegativeTransparencyClassOnce.Do(func() {
		_ICScannerFunctionalUnitNegativeTransparencyClass = ICScannerFunctionalUnitNegativeTransparencyClass{class: objc.GetClass("ICScannerFunctionalUnitNegativeTransparency")}
	})
	return _ICScannerFunctionalUnitNegativeTransparencyClass
}

// GetICScannerFunctionalUnitNegativeTransparencyClass returns the class object for ICScannerFunctionalUnitNegativeTransparency.
func GetICScannerFunctionalUnitNegativeTransparencyClass() ICScannerFunctionalUnitNegativeTransparencyClass {
	return getICScannerFunctionalUnitNegativeTransparencyClass()
}

type ICScannerFunctionalUnitNegativeTransparencyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFunctionalUnitNegativeTransparencyClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFunctionalUnitNegativeTransparencyClass) Alloc() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the transparency unit for scanning negatives on
// the scanner.
//
// # Instance Properties
//
//   - [ICScannerFunctionalUnitNegativeTransparency.DocumentSize]
//   - [ICScannerFunctionalUnitNegativeTransparency.DocumentType]
//   - [ICScannerFunctionalUnitNegativeTransparency.SetDocumentType]
//   - [ICScannerFunctionalUnitNegativeTransparency.SupportedDocumentTypes]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitNegativeTransparency
type ICScannerFunctionalUnitNegativeTransparency struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitNegativeTransparencyFromID constructs a [ICScannerFunctionalUnitNegativeTransparency] from an objc.ID.
//
// An object that represents the transparency unit for scanning negatives on
// the scanner.
func ICScannerFunctionalUnitNegativeTransparencyFromID(id objc.ID) ICScannerFunctionalUnitNegativeTransparency {
	return ICScannerFunctionalUnitNegativeTransparency{ICScannerFunctionalUnit: ICScannerFunctionalUnitFromID(id)}
}

// NOTE: ICScannerFunctionalUnitNegativeTransparency adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFunctionalUnitNegativeTransparency] class.
//
// # Instance Properties
//
//   - [IICScannerFunctionalUnitNegativeTransparency.DocumentSize]
//   - [IICScannerFunctionalUnitNegativeTransparency.DocumentType]
//   - [IICScannerFunctionalUnitNegativeTransparency.SetDocumentType]
//   - [IICScannerFunctionalUnitNegativeTransparency.SupportedDocumentTypes]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitNegativeTransparency
type IICScannerFunctionalUnitNegativeTransparency interface {
	IICScannerFunctionalUnit

	// Topic: Instance Properties

	DocumentSize() corefoundation.CGSize
	DocumentType() ICScannerDocumentType
	SetDocumentType(value ICScannerDocumentType)
	SupportedDocumentTypes() foundation.NSIndexSet
}

// Init initializes the instance.
func (s ICScannerFunctionalUnitNegativeTransparency) Init() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFunctionalUnitNegativeTransparency) Autorelease() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitNegativeTransparency creates a new ICScannerFunctionalUnitNegativeTransparency instance.
func NewICScannerFunctionalUnitNegativeTransparency() ICScannerFunctionalUnitNegativeTransparency {
	class := getICScannerFunctionalUnitNegativeTransparencyClass()
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitNegativeTransparency/documentSize
func (s ICScannerFunctionalUnitNegativeTransparency) DocumentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("documentSize"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitNegativeTransparency/documentType
func (s ICScannerFunctionalUnitNegativeTransparency) DocumentType() ICScannerDocumentType {
	rv := objc.Send[ICScannerDocumentType](s.ID, objc.Sel("documentType"))
	return ICScannerDocumentType(rv)
}
func (s ICScannerFunctionalUnitNegativeTransparency) SetDocumentType(value ICScannerDocumentType) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentType:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitNegativeTransparency/supportedDocumentTypes
func (s ICScannerFunctionalUnitNegativeTransparency) SupportedDocumentTypes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedDocumentTypes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
