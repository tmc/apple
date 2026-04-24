// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLVirtualDisplayConfiguration] class.
var (
	_SLVirtualDisplayConfigurationClass     SLVirtualDisplayConfigurationClass
	_SLVirtualDisplayConfigurationClassOnce sync.Once
)

func getSLVirtualDisplayConfigurationClass() SLVirtualDisplayConfigurationClass {
	_SLVirtualDisplayConfigurationClassOnce.Do(func() {
		_SLVirtualDisplayConfigurationClass = SLVirtualDisplayConfigurationClass{class: objc.GetClass("SLVirtualDisplayConfiguration")}
	})
	return _SLVirtualDisplayConfigurationClass
}

// GetSLVirtualDisplayConfigurationClass returns the class object for SLVirtualDisplayConfiguration.
func GetSLVirtualDisplayConfigurationClass() SLVirtualDisplayConfigurationClass {
	return getSLVirtualDisplayConfigurationClass()
}

type SLVirtualDisplayConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLVirtualDisplayConfigurationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLVirtualDisplayConfigurationClass) Alloc() SLVirtualDisplayConfiguration {
	rv := objc.Send[SLVirtualDisplayConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLVirtualDisplayConfiguration.Chromaticities]
//   - [SLVirtualDisplayConfiguration.DictionaryRepresentation]
//   - [SLVirtualDisplayConfiguration.MaximumSizeInPixels]
//   - [SLVirtualDisplayConfiguration.Name]
//   - [SLVirtualDisplayConfiguration.Options]
//   - [SLVirtualDisplayConfiguration.SetOptions]
//   - [SLVirtualDisplayConfiguration.ProductID]
//   - [SLVirtualDisplayConfiguration.SerialNumber]
//   - [SLVirtualDisplayConfiguration.SizeInMillimeters]
//   - [SLVirtualDisplayConfiguration.Subtype]
//   - [SLVirtualDisplayConfiguration.SetSubtype]
//   - [SLVirtualDisplayConfiguration.Type]
//   - [SLVirtualDisplayConfiguration.SetType]
//   - [SLVirtualDisplayConfiguration.Uti]
//   - [SLVirtualDisplayConfiguration.SetUti]
//   - [SLVirtualDisplayConfiguration.VendorID]
//   - [SLVirtualDisplayConfiguration.InitWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration
type SLVirtualDisplayConfiguration struct {
	objectivec.Object
}

// SLVirtualDisplayConfigurationFromID constructs a [SLVirtualDisplayConfiguration] from an objc.ID.
func SLVirtualDisplayConfigurationFromID(id objc.ID) SLVirtualDisplayConfiguration {
	return SLVirtualDisplayConfiguration{objectivec.Object{ID: id}}
}

// Ensure SLVirtualDisplayConfiguration implements ISLVirtualDisplayConfiguration.
var _ ISLVirtualDisplayConfiguration = SLVirtualDisplayConfiguration{}

// An interface definition for the [SLVirtualDisplayConfiguration] class.
//
// # Methods
//
//   - [ISLVirtualDisplayConfiguration.Chromaticities]
//   - [ISLVirtualDisplayConfiguration.DictionaryRepresentation]
//   - [ISLVirtualDisplayConfiguration.MaximumSizeInPixels]
//   - [ISLVirtualDisplayConfiguration.Name]
//   - [ISLVirtualDisplayConfiguration.Options]
//   - [ISLVirtualDisplayConfiguration.SetOptions]
//   - [ISLVirtualDisplayConfiguration.ProductID]
//   - [ISLVirtualDisplayConfiguration.SerialNumber]
//   - [ISLVirtualDisplayConfiguration.SizeInMillimeters]
//   - [ISLVirtualDisplayConfiguration.Subtype]
//   - [ISLVirtualDisplayConfiguration.SetSubtype]
//   - [ISLVirtualDisplayConfiguration.Type]
//   - [ISLVirtualDisplayConfiguration.SetType]
//   - [ISLVirtualDisplayConfiguration.Uti]
//   - [ISLVirtualDisplayConfiguration.SetUti]
//   - [ISLVirtualDisplayConfiguration.VendorID]
//   - [ISLVirtualDisplayConfiguration.InitWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError]
//
// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration
type ISLVirtualDisplayConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	Chromaticities() objectivec.IObject
	DictionaryRepresentation() objectivec.IObject
	MaximumSizeInPixels() objectivec.IObject
	Name() string
	Options() uint64
	SetOptions(value uint64)
	ProductID() uint64
	SerialNumber() uint64
	SizeInMillimeters() objectivec.IObject
	Subtype() uint64
	SetSubtype(value uint64)
	Type() uint64
	SetType(value uint64)
	Uti() string
	SetUti(value string)
	VendorID() uint64
	InitWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError(name objectivec.IObject, id uint64, id2 uint64, number uint64, millimeters objectivec.IObject, pixels objectivec.IObject, chromaticities objectivec.IObject) (SLVirtualDisplayConfiguration, error)
}

// Init initializes the instance.
func (s SLVirtualDisplayConfiguration) Init() SLVirtualDisplayConfiguration {
	rv := objc.Send[SLVirtualDisplayConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplayConfiguration) Autorelease() SLVirtualDisplayConfiguration {
	rv := objc.Send[SLVirtualDisplayConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplayConfiguration creates a new SLVirtualDisplayConfiguration instance.
func NewSLVirtualDisplayConfiguration() SLVirtualDisplayConfiguration {
	class := getSLVirtualDisplayConfigurationClass()
	rv := objc.Send[SLVirtualDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/initWithName:vendorID:productID:serialNumber:sizeInMillimeters:maximumSizeInPixels:chromaticities:error:
func NewSLVirtualDisplayConfigurationWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError(name objectivec.IObject, id uint64, id2 uint64, number uint64, millimeters objectivec.IObject, pixels objectivec.IObject, chromaticities objectivec.IObject) (SLVirtualDisplayConfiguration, error) {
	var errorPtr objc.ID
	instance := getSLVirtualDisplayConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:vendorID:productID:serialNumber:sizeInMillimeters:maximumSizeInPixels:chromaticities:error:"), name, id, id2, number, millimeters, pixels, chromaticities, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayConfiguration{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayConfigurationFromID(rv), nil
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/dictionaryRepresentation
func (s SLVirtualDisplayConfiguration) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/initWithName:vendorID:productID:serialNumber:sizeInMillimeters:maximumSizeInPixels:chromaticities:error:
func (s SLVirtualDisplayConfiguration) InitWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError(name objectivec.IObject, id uint64, id2 uint64, number uint64, millimeters objectivec.IObject, pixels objectivec.IObject, chromaticities objectivec.IObject) (SLVirtualDisplayConfiguration, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithName:vendorID:productID:serialNumber:sizeInMillimeters:maximumSizeInPixels:chromaticities:error:"), name, id, id2, number, millimeters, pixels, chromaticities, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayConfiguration{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayConfigurationFromID(rv), nil

}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/configurationWithBackendOptions:
func (_SLVirtualDisplayConfigurationClass SLVirtualDisplayConfigurationClass) ConfigurationWithBackendOptions(options objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplayConfigurationClass.class), objc.Sel("configurationWithBackendOptions:"), options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/configurationWithDictionaryRepresentation:
func (_SLVirtualDisplayConfigurationClass SLVirtualDisplayConfigurationClass) ConfigurationWithDictionaryRepresentation(representation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplayConfigurationClass.class), objc.Sel("configurationWithDictionaryRepresentation:"), representation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/configurationWithDisplayInfo:
func (_SLVirtualDisplayConfigurationClass SLVirtualDisplayConfigurationClass) ConfigurationWithDisplayInfo(info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLVirtualDisplayConfigurationClass.class), objc.Sel("configurationWithDisplayInfo:"), info)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/chromaticities
func (s SLVirtualDisplayConfiguration) Chromaticities() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("chromaticities"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/maximumSizeInPixels
func (s SLVirtualDisplayConfiguration) MaximumSizeInPixels() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("maximumSizeInPixels"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/name
func (s SLVirtualDisplayConfiguration) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/options
func (s SLVirtualDisplayConfiguration) Options() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("options"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SetOptions(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setOptions:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/productID
func (s SLVirtualDisplayConfiguration) ProductID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("productID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/serialNumber
func (s SLVirtualDisplayConfiguration) SerialNumber() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("serialNumber"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/sizeInMillimeters
func (s SLVirtualDisplayConfiguration) SizeInMillimeters() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sizeInMillimeters"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/subtype
func (s SLVirtualDisplayConfiguration) Subtype() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("subtype"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SetSubtype(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setSubtype:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/type
func (s SLVirtualDisplayConfiguration) Type() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("type"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SetType(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setType:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/uti
func (s SLVirtualDisplayConfiguration) Uti() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uti"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLVirtualDisplayConfiguration) SetUti(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setUti:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/SkyLight/SLVirtualDisplayConfiguration/vendorID
func (s SLVirtualDisplayConfiguration) VendorID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("vendorID"))
	return rv
}
