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
	rv := objc.SendIfResponds[SLVirtualDisplayConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
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
type ISLVirtualDisplayConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	Chromaticities() unsafe.Pointer
	DictionaryRepresentation() objectivec.IObject
	MaximumSizeInPixels() unsafe.Pointer
	Name() string
	Options() uint64
	SetOptions(value uint64)
	ProductID() uint64
	SerialNumber() uint64
	SizeInMillimeters() unsafe.Pointer
	Subtype() uint64
	SetSubtype(value uint64)
	Type() uint64
	SetType(value uint64)
	Uti() string
	SetUti(value string)
	VendorID() uint64
	InitWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError(name objectivec.IObject, id uint64, id2 uint64, number uint64, millimeters unsafe.Pointer, pixels unsafe.Pointer, chromaticities unsafe.Pointer) (SLVirtualDisplayConfiguration, error)
}

// Init initializes the instance.
func (s SLVirtualDisplayConfiguration) Init() SLVirtualDisplayConfiguration {
	rv := objc.SendIfResponds[SLVirtualDisplayConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLVirtualDisplayConfiguration) Autorelease() SLVirtualDisplayConfiguration {
	rv := objc.SendIfResponds[SLVirtualDisplayConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLVirtualDisplayConfiguration creates a new SLVirtualDisplayConfiguration instance.
func NewSLVirtualDisplayConfiguration() SLVirtualDisplayConfiguration {
	class := getSLVirtualDisplayConfigurationClass()
	rv := objc.SendIfResponds[SLVirtualDisplayConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLVirtualDisplayConfigurationWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError(name objectivec.IObject, id uint64, id2 uint64, number uint64, millimeters unsafe.Pointer, pixels unsafe.Pointer, chromaticities unsafe.Pointer) (SLVirtualDisplayConfiguration, error) {
	var errorPtr objc.ID
	instance := getSLVirtualDisplayConfigurationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:vendorID:productID:serialNumber:sizeInMillimeters:maximumSizeInPixels:chromaticities:error:"), name, id, id2, number, millimeters, pixels, chromaticities, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayConfiguration{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SLVirtualDisplayConfiguration{}, objc.ErrInitFailed
	}
	return SLVirtualDisplayConfigurationFromID(rv), nil
}

func (s SLVirtualDisplayConfiguration) DictionaryRepresentation() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
func (s SLVirtualDisplayConfiguration) InitWithNameVendorIDProductIDSerialNumberSizeInMillimetersMaximumSizeInPixelsChromaticitiesError(name objectivec.IObject, id uint64, id2 uint64, number uint64, millimeters unsafe.Pointer, pixels unsafe.Pointer, chromaticities unsafe.Pointer) (SLVirtualDisplayConfiguration, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithName:vendorID:productID:serialNumber:sizeInMillimeters:maximumSizeInPixels:chromaticities:error:"), name, id, id2, number, millimeters, pixels, chromaticities, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SLVirtualDisplayConfiguration{}, foundation.NSErrorFrom(errorPtr)
	}
	return SLVirtualDisplayConfigurationFromID(rv), nil

}

func (_SLVirtualDisplayConfigurationClass SLVirtualDisplayConfigurationClass) ConfigurationWithBackendOptions(options objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLVirtualDisplayConfigurationClass.class), objc.Sel("configurationWithBackendOptions:"), options)
	return objectivec.Object{ID: rv}
}
func (_SLVirtualDisplayConfigurationClass SLVirtualDisplayConfigurationClass) ConfigurationWithDictionaryRepresentation(representation objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLVirtualDisplayConfigurationClass.class), objc.Sel("configurationWithDictionaryRepresentation:"), representation)
	return objectivec.Object{ID: rv}
}
func (_SLVirtualDisplayConfigurationClass SLVirtualDisplayConfigurationClass) ConfigurationWithDisplayInfo(info objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SLVirtualDisplayConfigurationClass.class), objc.Sel("configurationWithDisplayInfo:"), info)
	return objectivec.Object{ID: rv}
}

func (s SLVirtualDisplayConfiguration) Chromaticities() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("chromaticities"))
	return rv
}
func (s SLVirtualDisplayConfiguration) MaximumSizeInPixels() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("maximumSizeInPixels"))
	return rv
}
func (s SLVirtualDisplayConfiguration) Name() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLVirtualDisplayConfiguration) Options() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("options"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SetOptions(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setOptions:"), value)
}
func (s SLVirtualDisplayConfiguration) ProductID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("productID"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SerialNumber() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("serialNumber"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SizeInMillimeters() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("sizeInMillimeters"))
	return rv
}
func (s SLVirtualDisplayConfiguration) Subtype() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("subtype"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SetSubtype(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setSubtype:"), value)
}
func (s SLVirtualDisplayConfiguration) Type() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("type"))
	return rv
}
func (s SLVirtualDisplayConfiguration) SetType(value uint64) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setType:"), value)
}
func (s SLVirtualDisplayConfiguration) Uti() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("uti"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLVirtualDisplayConfiguration) SetUti(value string) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setUti:"), objc.String(value))
}
func (s SLVirtualDisplayConfiguration) VendorID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("vendorID"))
	return rv
}
