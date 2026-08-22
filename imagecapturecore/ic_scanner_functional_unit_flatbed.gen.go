// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ICScannerFunctionalUnitFlatbed] class.
var (
	_ICScannerFunctionalUnitFlatbedClass     ICScannerFunctionalUnitFlatbedClass
	_ICScannerFunctionalUnitFlatbedClassOnce sync.Once
)

func getICScannerFunctionalUnitFlatbedClass() ICScannerFunctionalUnitFlatbedClass {
	_ICScannerFunctionalUnitFlatbedClassOnce.Do(func() {
		_ICScannerFunctionalUnitFlatbedClass = ICScannerFunctionalUnitFlatbedClass{class: objc.GetClass("ICScannerFunctionalUnitFlatbed")}
	})
	return _ICScannerFunctionalUnitFlatbedClass
}

// GetICScannerFunctionalUnitFlatbedClass returns the class object for ICScannerFunctionalUnitFlatbed.
func GetICScannerFunctionalUnitFlatbedClass() ICScannerFunctionalUnitFlatbedClass {
	return getICScannerFunctionalUnitFlatbedClass()
}

type ICScannerFunctionalUnitFlatbedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICScannerFunctionalUnitFlatbedClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICScannerFunctionalUnitFlatbedClass) Alloc() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the flatbed unit on a scanner.
//
// # Instance Properties
//
//   - [ICScannerFunctionalUnitFlatbed.DocumentSize]
//   - [ICScannerFunctionalUnitFlatbed.DocumentType]
//   - [ICScannerFunctionalUnitFlatbed.SetDocumentType]
//   - [ICScannerFunctionalUnitFlatbed.SupportedDocumentTypes]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitFlatbed
type ICScannerFunctionalUnitFlatbed struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitFlatbedFromID constructs a [ICScannerFunctionalUnitFlatbed] from an objc.ID.
//
// An object that represents the flatbed unit on a scanner.
func ICScannerFunctionalUnitFlatbedFromID(id objc.ID) ICScannerFunctionalUnitFlatbed {
	return ICScannerFunctionalUnitFlatbed{ICScannerFunctionalUnit: ICScannerFunctionalUnitFromID(id)}
}

// NOTE: ICScannerFunctionalUnitFlatbed adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICScannerFunctionalUnitFlatbed] class.
//
// # Instance Properties
//
//   - [IICScannerFunctionalUnitFlatbed.DocumentSize]
//   - [IICScannerFunctionalUnitFlatbed.DocumentType]
//   - [IICScannerFunctionalUnitFlatbed.SetDocumentType]
//   - [IICScannerFunctionalUnitFlatbed.SupportedDocumentTypes]
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitFlatbed
type IICScannerFunctionalUnitFlatbed interface {
	IICScannerFunctionalUnit

	// Topic: Instance Properties

	DocumentSize() corefoundation.CGSize
	DocumentType() ICScannerDocumentType
	SetDocumentType(value ICScannerDocumentType)
	SupportedDocumentTypes() foundation.NSIndexSet
}

// Init initializes the instance.
func (s ICScannerFunctionalUnitFlatbed) Init() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s ICScannerFunctionalUnitFlatbed) Autorelease() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitFlatbed creates a new ICScannerFunctionalUnitFlatbed instance.
func NewICScannerFunctionalUnitFlatbed() ICScannerFunctionalUnitFlatbed {
	class := getICScannerFunctionalUnitFlatbedClass()
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitFlatbed/documentSize
func (s ICScannerFunctionalUnitFlatbed) DocumentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("documentSize"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitFlatbed/documentType
func (s ICScannerFunctionalUnitFlatbed) DocumentType() ICScannerDocumentType {
	rv := objc.Send[ICScannerDocumentType](s.ID, objc.Sel("documentType"))
	return ICScannerDocumentType(rv)
}
func (s ICScannerFunctionalUnitFlatbed) SetDocumentType(value ICScannerDocumentType) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentType:"), value)
}

// See: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitFlatbed/supportedDocumentTypes
func (s ICScannerFunctionalUnitFlatbed) SupportedDocumentTypes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedDocumentTypes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
