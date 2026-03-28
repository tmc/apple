// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioHeatmapHistogram] class.
var (
	_GTMioHeatmapHistogramClass     GTMioHeatmapHistogramClass
	_GTMioHeatmapHistogramClassOnce sync.Once
)

func getGTMioHeatmapHistogramClass() GTMioHeatmapHistogramClass {
	_GTMioHeatmapHistogramClassOnce.Do(func() {
		_GTMioHeatmapHistogramClass = GTMioHeatmapHistogramClass{class: objc.GetClass("GTMioHeatmapHistogram")}
	})
	return _GTMioHeatmapHistogramClass
}

// GetGTMioHeatmapHistogramClass returns the class object for GTMioHeatmapHistogram.
func GetGTMioHeatmapHistogramClass() GTMioHeatmapHistogramClass {
	return getGTMioHeatmapHistogramClass()
}

type GTMioHeatmapHistogramClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioHeatmapHistogramClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioHeatmapHistogramClass) Alloc() GTMioHeatmapHistogram {
	rv := objc.Send[GTMioHeatmapHistogram](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioHeatmapHistogram._generate]
//   - [GTMioHeatmapHistogram.GenerateImageColor]
//   - [GTMioHeatmapHistogram.MaxCount]
//   - [GTMioHeatmapHistogram.MaxValue]
//   - [GTMioHeatmapHistogram.MinCount]
//   - [GTMioHeatmapHistogram.MinValue]
//   - [GTMioHeatmapHistogram.NumBuckets]
//   - [GTMioHeatmapHistogram.Values]
//   - [GTMioHeatmapHistogram.InitWithHeatmapMinValueMaxValueNumBuckets]
//   - [GTMioHeatmapHistogram.InitWithHeatmapNumBuckets]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram
type GTMioHeatmapHistogram struct {
	objectivec.Object
}

// GTMioHeatmapHistogramFromID constructs a [GTMioHeatmapHistogram] from an objc.ID.
func GTMioHeatmapHistogramFromID(id objc.ID) GTMioHeatmapHistogram {
	return GTMioHeatmapHistogram{objectivec.Object{ID: id}}
}
// Ensure GTMioHeatmapHistogram implements IGTMioHeatmapHistogram.
var _ IGTMioHeatmapHistogram = GTMioHeatmapHistogram{}

// An interface definition for the [GTMioHeatmapHistogram] class.
//
// # Methods
//
//   - [IGTMioHeatmapHistogram._generate]
//   - [IGTMioHeatmapHistogram.GenerateImageColor]
//   - [IGTMioHeatmapHistogram.MaxCount]
//   - [IGTMioHeatmapHistogram.MaxValue]
//   - [IGTMioHeatmapHistogram.MinCount]
//   - [IGTMioHeatmapHistogram.MinValue]
//   - [IGTMioHeatmapHistogram.NumBuckets]
//   - [IGTMioHeatmapHistogram.Values]
//   - [IGTMioHeatmapHistogram.InitWithHeatmapMinValueMaxValueNumBuckets]
//   - [IGTMioHeatmapHistogram.InitWithHeatmapNumBuckets]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram
type IGTMioHeatmapHistogram interface {
	objectivec.IObject

	// Topic: Methods

	_generate()
	GenerateImageColor(image corefoundation.CGSize, color *coregraphics.CGColorRef) *coregraphics.CGImageRef
	MaxCount() uint64
	MaxValue() uint64
	MinCount() uint64
	MinValue() uint64
	NumBuckets() uint32
	Values() unsafe.Pointer
	InitWithHeatmapMinValueMaxValueNumBuckets(heatmap objectivec.IObject, value uint64, value2 uint64, buckets uint32) GTMioHeatmapHistogram
	InitWithHeatmapNumBuckets(heatmap objectivec.IObject, buckets uint32) GTMioHeatmapHistogram
}

// Init initializes the instance.
func (g GTMioHeatmapHistogram) Init() GTMioHeatmapHistogram {
	rv := objc.Send[GTMioHeatmapHistogram](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioHeatmapHistogram) Autorelease() GTMioHeatmapHistogram {
	rv := objc.Send[GTMioHeatmapHistogram](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioHeatmapHistogram creates a new GTMioHeatmapHistogram instance.
func NewGTMioHeatmapHistogram() GTMioHeatmapHistogram {
	class := getGTMioHeatmapHistogramClass()
	rv := objc.Send[GTMioHeatmapHistogram](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/initWithHeatmap:minValue:maxValue:numBuckets:
func NewGTMioHeatmapHistogramWithHeatmapMinValueMaxValueNumBuckets(heatmap objectivec.IObject, value uint64, value2 uint64, buckets uint32) GTMioHeatmapHistogram {
	instance := getGTMioHeatmapHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHeatmap:minValue:maxValue:numBuckets:"), heatmap, value, value2, buckets)
	return GTMioHeatmapHistogramFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/initWithHeatmap:numBuckets:
func NewGTMioHeatmapHistogramWithHeatmapNumBuckets(heatmap objectivec.IObject, buckets uint32) GTMioHeatmapHistogram {
	instance := getGTMioHeatmapHistogramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHeatmap:numBuckets:"), heatmap, buckets)
	return GTMioHeatmapHistogramFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/_generate
func (g GTMioHeatmapHistogram) _generate() {
	objc.Send[objc.ID](g.ID, objc.Sel("_generate"))
}

// Generate is an exported wrapper for the private method _generate.
func (g GTMioHeatmapHistogram) Generate() {
	g._generate()
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/generateImage:color:
func (g GTMioHeatmapHistogram) GenerateImageColor(image corefoundation.CGSize, color *coregraphics.CGColorRef) *coregraphics.CGImageRef {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("generateImage:color:"), image, color)
		return (*coregraphics.CGImageRef)(rv)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/initWithHeatmap:minValue:maxValue:numBuckets:
func (g GTMioHeatmapHistogram) InitWithHeatmapMinValueMaxValueNumBuckets(heatmap objectivec.IObject, value uint64, value2 uint64, buckets uint32) GTMioHeatmapHistogram {
	rv := objc.Send[GTMioHeatmapHistogram](g.ID, objc.Sel("initWithHeatmap:minValue:maxValue:numBuckets:"), heatmap, value, value2, buckets)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/initWithHeatmap:numBuckets:
func (g GTMioHeatmapHistogram) InitWithHeatmapNumBuckets(heatmap objectivec.IObject, buckets uint32) GTMioHeatmapHistogram {
	rv := objc.Send[GTMioHeatmapHistogram](g.ID, objc.Sel("initWithHeatmap:numBuckets:"), heatmap, buckets)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/maxCount
func (g GTMioHeatmapHistogram) MaxCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("maxCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/maxValue
func (g GTMioHeatmapHistogram) MaxValue() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("maxValue"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/minCount
func (g GTMioHeatmapHistogram) MinCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("minCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/minValue
func (g GTMioHeatmapHistogram) MinValue() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("minValue"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/numBuckets
func (g GTMioHeatmapHistogram) NumBuckets() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("numBuckets"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapHistogram/values
func (g GTMioHeatmapHistogram) Values() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("values"))
	return rv
}

