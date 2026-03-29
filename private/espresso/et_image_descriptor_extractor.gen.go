// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETImageDescriptorExtractor] class.
var (
	_ETImageDescriptorExtractorClass     ETImageDescriptorExtractorClass
	_ETImageDescriptorExtractorClassOnce sync.Once
)

func getETImageDescriptorExtractorClass() ETImageDescriptorExtractorClass {
	_ETImageDescriptorExtractorClassOnce.Do(func() {
		_ETImageDescriptorExtractorClass = ETImageDescriptorExtractorClass{class: objc.GetClass("ETImageDescriptorExtractor")}
	})
	return _ETImageDescriptorExtractorClass
}

// GetETImageDescriptorExtractorClass returns the class object for ETImageDescriptorExtractor.
func GetETImageDescriptorExtractorClass() ETImageDescriptorExtractorClass {
	return getETImageDescriptorExtractorClass()
}

type ETImageDescriptorExtractorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETImageDescriptorExtractorClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETImageDescriptorExtractorClass) Alloc() ETImageDescriptorExtractor {
	rv := objc.Send[ETImageDescriptorExtractor](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETImageDescriptorExtractor.Brightness_range]
//   - [ETImageDescriptorExtractor.SetBrightness_range]
//   - [ETImageDescriptorExtractor.Contrast_range]
//   - [ETImageDescriptorExtractor.SetContrast_range]
//   - [ETImageDescriptorExtractor.CropResizeInputImage]
//   - [ETImageDescriptorExtractor.Descriptors_file_cache_size]
//   - [ETImageDescriptorExtractor.SetDescriptors_file_cache_size]
//   - [ETImageDescriptorExtractor.Descriptors_mem_cache_size]
//   - [ETImageDescriptorExtractor.SetDescriptors_mem_cache_size]
//   - [ETImageDescriptorExtractor.DoBatchnormTuning]
//   - [ETImageDescriptorExtractor.SetDoBatchnormTuning]
//   - [ETImageDescriptorExtractor.ExtractDescriptorForDataPointFreeWhenDone]
//   - [ETImageDescriptorExtractor.ExtractForDataPoint]
//   - [ETImageDescriptorExtractor.Horizontal_flip]
//   - [ETImageDescriptorExtractor.SetHorizontal_flip]
//   - [ETImageDescriptorExtractor.NAugmentations]
//   - [ETImageDescriptorExtractor.SetNAugmentations]
//   - [ETImageDescriptorExtractor.NumberOfChannels]
//   - [ETImageDescriptorExtractor.Rotation_range]
//   - [ETImageDescriptorExtractor.SetRotation_range]
//   - [ETImageDescriptorExtractor.Shear_range]
//   - [ETImageDescriptorExtractor.SetShear_range]
//   - [ETImageDescriptorExtractor.Zoom_range]
//   - [ETImageDescriptorExtractor.SetZoom_range]
//   - [ETImageDescriptorExtractor.InitWithNetwork]
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor
type ETImageDescriptorExtractor struct {
	objectivec.Object
}

// ETImageDescriptorExtractorFromID constructs a [ETImageDescriptorExtractor] from an objc.ID.
func ETImageDescriptorExtractorFromID(id objc.ID) ETImageDescriptorExtractor {
	return ETImageDescriptorExtractor{objectivec.Object{ID: id}}
}
// Ensure ETImageDescriptorExtractor implements IETImageDescriptorExtractor.
var _ IETImageDescriptorExtractor = ETImageDescriptorExtractor{}

// An interface definition for the [ETImageDescriptorExtractor] class.
//
// # Methods
//
//   - [IETImageDescriptorExtractor.Brightness_range]
//   - [IETImageDescriptorExtractor.SetBrightness_range]
//   - [IETImageDescriptorExtractor.Contrast_range]
//   - [IETImageDescriptorExtractor.SetContrast_range]
//   - [IETImageDescriptorExtractor.CropResizeInputImage]
//   - [IETImageDescriptorExtractor.Descriptors_file_cache_size]
//   - [IETImageDescriptorExtractor.SetDescriptors_file_cache_size]
//   - [IETImageDescriptorExtractor.Descriptors_mem_cache_size]
//   - [IETImageDescriptorExtractor.SetDescriptors_mem_cache_size]
//   - [IETImageDescriptorExtractor.DoBatchnormTuning]
//   - [IETImageDescriptorExtractor.SetDoBatchnormTuning]
//   - [IETImageDescriptorExtractor.ExtractDescriptorForDataPointFreeWhenDone]
//   - [IETImageDescriptorExtractor.ExtractForDataPoint]
//   - [IETImageDescriptorExtractor.Horizontal_flip]
//   - [IETImageDescriptorExtractor.SetHorizontal_flip]
//   - [IETImageDescriptorExtractor.NAugmentations]
//   - [IETImageDescriptorExtractor.SetNAugmentations]
//   - [IETImageDescriptorExtractor.NumberOfChannels]
//   - [IETImageDescriptorExtractor.Rotation_range]
//   - [IETImageDescriptorExtractor.SetRotation_range]
//   - [IETImageDescriptorExtractor.Shear_range]
//   - [IETImageDescriptorExtractor.SetShear_range]
//   - [IETImageDescriptorExtractor.Zoom_range]
//   - [IETImageDescriptorExtractor.SetZoom_range]
//   - [IETImageDescriptorExtractor.InitWithNetwork]
//
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor
type IETImageDescriptorExtractor interface {
	objectivec.IObject

	// Topic: Methods

	Brightness_range() float32
	SetBrightness_range(value float32)
	Contrast_range() float32
	SetContrast_range(value float32)
	CropResizeInputImage(image unsafe.Pointer) unsafe.Pointer
	Descriptors_file_cache_size() uint64
	SetDescriptors_file_cache_size(value uint64)
	Descriptors_mem_cache_size() uint64
	SetDescriptors_mem_cache_size(value uint64)
	DoBatchnormTuning() int
	SetDoBatchnormTuning(value int)
	ExtractDescriptorForDataPointFreeWhenDone(point unsafe.Pointer, done bool) unsafe.Pointer
	ExtractForDataPoint(point objectivec.IObject)
	Horizontal_flip() float32
	SetHorizontal_flip(value float32)
	NAugmentations() int
	SetNAugmentations(value int)
	NumberOfChannels() int
	Rotation_range() float32
	SetRotation_range(value float32)
	Shear_range() float32
	SetShear_range(value float32)
	Zoom_range() float32
	SetZoom_range(value float32)
	InitWithNetwork(network objectivec.IObject) ETImageDescriptorExtractor
}

// Init initializes the instance.
func (e ETImageDescriptorExtractor) Init() ETImageDescriptorExtractor {
	rv := objc.Send[ETImageDescriptorExtractor](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETImageDescriptorExtractor) Autorelease() ETImageDescriptorExtractor {
	rv := objc.Send[ETImageDescriptorExtractor](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETImageDescriptorExtractor creates a new ETImageDescriptorExtractor instance.
func NewETImageDescriptorExtractor() ETImageDescriptorExtractor {
	class := getETImageDescriptorExtractorClass()
	rv := objc.Send[ETImageDescriptorExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/initWithNetwork:
func NewETImageDescriptorExtractorWithNetwork(network objectivec.IObject) ETImageDescriptorExtractor {
	instance := getETImageDescriptorExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETImageDescriptorExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/cropResizeInputImage:
func (e ETImageDescriptorExtractor) CropResizeInputImage(image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("cropResizeInputImage:"), image)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/extractDescriptorForDataPoint:freeWhenDone:
func (e ETImageDescriptorExtractor) ExtractDescriptorForDataPointFreeWhenDone(point unsafe.Pointer, done bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("extractDescriptorForDataPoint:freeWhenDone:"), point, done)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/extractForDataPoint:
func (e ETImageDescriptorExtractor) ExtractForDataPoint(point objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("extractForDataPoint:"), point)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/numberOfChannels
func (e ETImageDescriptorExtractor) NumberOfChannels() int {
	rv := objc.Send[int](e.ID, objc.Sel("numberOfChannels"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/initWithNetwork:
func (e ETImageDescriptorExtractor) InitWithNetwork(network objectivec.IObject) ETImageDescriptorExtractor {
	rv := objc.Send[ETImageDescriptorExtractor](e.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/VisionSceneNet_iOS10_Extractor
func (_ETImageDescriptorExtractorClass ETImageDescriptorExtractorClass) VisionSceneNet_iOS10_Extractor() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETImageDescriptorExtractorClass.class), objc.Sel("VisionSceneNet_iOS10_Extractor"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/VisionSmartCamNet_iOS11_Extractor
func (_ETImageDescriptorExtractorClass ETImageDescriptorExtractorClass) VisionSmartCamNet_iOS11_Extractor() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETImageDescriptorExtractorClass.class), objc.Sel("VisionSmartCamNet_iOS11_Extractor"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/inception_v3_Extractor
func (_ETImageDescriptorExtractorClass ETImageDescriptorExtractorClass) Inception_v3_Extractor() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETImageDescriptorExtractorClass.class), objc.Sel("inception_v3_Extractor"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/passthroughExtractor
func (_ETImageDescriptorExtractorClass ETImageDescriptorExtractorClass) PassthroughExtractor() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETImageDescriptorExtractorClass.class), objc.Sel("passthroughExtractor"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/brightness_range
func (e ETImageDescriptorExtractor) Brightness_range() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("brightness_range"))
	return rv
}
func (e ETImageDescriptorExtractor) SetBrightness_range(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setBrightness_range:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/contrast_range
func (e ETImageDescriptorExtractor) Contrast_range() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("contrast_range"))
	return rv
}
func (e ETImageDescriptorExtractor) SetContrast_range(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setContrast_range:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/descriptors_file_cache_size
func (e ETImageDescriptorExtractor) Descriptors_file_cache_size() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("descriptors_file_cache_size"))
	return rv
}
func (e ETImageDescriptorExtractor) SetDescriptors_file_cache_size(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setDescriptors_file_cache_size:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/descriptors_mem_cache_size
func (e ETImageDescriptorExtractor) Descriptors_mem_cache_size() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("descriptors_mem_cache_size"))
	return rv
}
func (e ETImageDescriptorExtractor) SetDescriptors_mem_cache_size(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setDescriptors_mem_cache_size:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/doBatchnormTuning
func (e ETImageDescriptorExtractor) DoBatchnormTuning() int {
	rv := objc.Send[int](e.ID, objc.Sel("doBatchnormTuning"))
	return rv
}
func (e ETImageDescriptorExtractor) SetDoBatchnormTuning(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setDoBatchnormTuning:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/horizontal_flip
func (e ETImageDescriptorExtractor) Horizontal_flip() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("horizontal_flip"))
	return rv
}
func (e ETImageDescriptorExtractor) SetHorizontal_flip(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setHorizontal_flip:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/nAugmentations
func (e ETImageDescriptorExtractor) NAugmentations() int {
	rv := objc.Send[int](e.ID, objc.Sel("nAugmentations"))
	return rv
}
func (e ETImageDescriptorExtractor) SetNAugmentations(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setNAugmentations:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/rotation_range
func (e ETImageDescriptorExtractor) Rotation_range() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("rotation_range"))
	return rv
}
func (e ETImageDescriptorExtractor) SetRotation_range(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setRotation_range:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/shear_range
func (e ETImageDescriptorExtractor) Shear_range() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("shear_range"))
	return rv
}
func (e ETImageDescriptorExtractor) SetShear_range(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setShear_range:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETImageDescriptorExtractor/zoom_range
func (e ETImageDescriptorExtractor) Zoom_range() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("zoom_range"))
	return rv
}
func (e ETImageDescriptorExtractor) SetZoom_range(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setZoom_range:"), value)
}

