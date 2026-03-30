// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioHeatmapBuilder] class.
var (
	_GTMioHeatmapBuilderClass     GTMioHeatmapBuilderClass
	_GTMioHeatmapBuilderClassOnce sync.Once
)

func getGTMioHeatmapBuilderClass() GTMioHeatmapBuilderClass {
	_GTMioHeatmapBuilderClassOnce.Do(func() {
		_GTMioHeatmapBuilderClass = GTMioHeatmapBuilderClass{class: objc.GetClass("GTMioHeatmapBuilder")}
	})
	return _GTMioHeatmapBuilderClass
}

// GetGTMioHeatmapBuilderClass returns the class object for GTMioHeatmapBuilder.
func GetGTMioHeatmapBuilderClass() GTMioHeatmapBuilderClass {
	return getGTMioHeatmapBuilderClass()
}

type GTMioHeatmapBuilderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioHeatmapBuilderClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioHeatmapBuilderClass) Alloc() GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioHeatmapBuilder.BuildCliqueHeatmapWidthHeightGenerationOptions]
//   - [GTMioHeatmapBuilder.BuildHeatmapWidthHeightGenerationOptions]
//   - [GTMioHeatmapBuilder.BuildQuadData]
//   - [GTMioHeatmapBuilder.BuildTileHeatmapWidthHeightGenerationOptions]
//   - [GTMioHeatmapBuilder.Depth]
//   - [GTMioHeatmapBuilder.DrawIndex]
//   - [GTMioHeatmapBuilder.EncoderInfo]
//   - [GTMioHeatmapBuilder.HeatmapSizeHeightWrap]
//   - [GTMioHeatmapBuilder.Options]
//   - [GTMioHeatmapBuilder.PipelineStateId]
//   - [GTMioHeatmapBuilder.ProgramType]
//   - [GTMioHeatmapBuilder.QuadData]
//   - [GTMioHeatmapBuilder.InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions]
//   - [GTMioHeatmapBuilder.InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions]
//   - [GTMioHeatmapBuilder.InitWithTraceDataEncoderFunctionIndexProgramTypeOptions]
//   - [GTMioHeatmapBuilder.InitWithTraceDataInitPCEncoderFunctionIndexProgramTypeOptions]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder
type GTMioHeatmapBuilder struct {
	objectivec.Object
}

// GTMioHeatmapBuilderFromID constructs a [GTMioHeatmapBuilder] from an objc.ID.
func GTMioHeatmapBuilderFromID(id objc.ID) GTMioHeatmapBuilder {
	return GTMioHeatmapBuilder{objectivec.Object{ID: id}}
}

// Ensure GTMioHeatmapBuilder implements IGTMioHeatmapBuilder.
var _ IGTMioHeatmapBuilder = GTMioHeatmapBuilder{}

// An interface definition for the [GTMioHeatmapBuilder] class.
//
// # Methods
//
//   - [IGTMioHeatmapBuilder.BuildCliqueHeatmapWidthHeightGenerationOptions]
//   - [IGTMioHeatmapBuilder.BuildHeatmapWidthHeightGenerationOptions]
//   - [IGTMioHeatmapBuilder.BuildQuadData]
//   - [IGTMioHeatmapBuilder.BuildTileHeatmapWidthHeightGenerationOptions]
//   - [IGTMioHeatmapBuilder.Depth]
//   - [IGTMioHeatmapBuilder.DrawIndex]
//   - [IGTMioHeatmapBuilder.EncoderInfo]
//   - [IGTMioHeatmapBuilder.HeatmapSizeHeightWrap]
//   - [IGTMioHeatmapBuilder.Options]
//   - [IGTMioHeatmapBuilder.PipelineStateId]
//   - [IGTMioHeatmapBuilder.ProgramType]
//   - [IGTMioHeatmapBuilder.QuadData]
//   - [IGTMioHeatmapBuilder.InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions]
//   - [IGTMioHeatmapBuilder.InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions]
//   - [IGTMioHeatmapBuilder.InitWithTraceDataEncoderFunctionIndexProgramTypeOptions]
//   - [IGTMioHeatmapBuilder.InitWithTraceDataInitPCEncoderFunctionIndexProgramTypeOptions]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder
type IGTMioHeatmapBuilder interface {
	objectivec.IObject

	// Topic: Methods

	BuildCliqueHeatmapWidthHeightGenerationOptions(heatmap uint64, width uint64, height uint64, options uint) objectivec.IObject
	BuildHeatmapWidthHeightGenerationOptions(heatmap uint64, width uint64, height uint64, options uint) objectivec.IObject
	BuildQuadData()
	BuildTileHeatmapWidthHeightGenerationOptions(heatmap uint64, width uint64, height uint64, options uint) objectivec.IObject
	Depth() uint32
	DrawIndex() uint32
	EncoderInfo() unsafe.Pointer
	HeatmapSizeHeightWrap(size uint64, height uint64, wrap bool) corefoundation.CGSize
	Options() uint64
	PipelineStateId() uint64
	ProgramType() uint16
	QuadData() IGTMioEncoderQuadData
	InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions(data objectivec.IObject, index uint32, index2 uint32, type_ uint16, options uint64) GTMioHeatmapBuilder
	InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions(data objectivec.IObject, index uint32, id uint64, type_ uint16, options uint64) GTMioHeatmapBuilder
	InitWithTraceDataEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, index uint32, type_ uint16, options uint64) GTMioHeatmapBuilder
	InitWithTraceDataInitPCEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, pc uint64, index uint32, type_ uint16, options uint64) GTMioHeatmapBuilder
}

// Init initializes the instance.
func (g GTMioHeatmapBuilder) Init() GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioHeatmapBuilder) Autorelease() GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioHeatmapBuilder creates a new GTMioHeatmapBuilder instance.
func NewGTMioHeatmapBuilder() GTMioHeatmapBuilder {
	class := getGTMioHeatmapBuilderClass()
	rv := objc.Send[GTMioHeatmapBuilder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:
func NewGTMioHeatmapBuilderWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions(data objectivec.IObject, index uint32, index2 uint32, type_ uint16, options uint64) GTMioHeatmapBuilder {
	instance := getGTMioHeatmapBuilderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:"), data, index, index2, type_, options)
	return GTMioHeatmapBuilderFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:
func NewGTMioHeatmapBuilderWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions(data objectivec.IObject, index uint32, id uint64, type_ uint16, options uint64) GTMioHeatmapBuilder {
	instance := getGTMioHeatmapBuilderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:"), data, index, id, type_, options)
	return GTMioHeatmapBuilderFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:encoderFunctionIndex:programType:options:
func NewGTMioHeatmapBuilderWithTraceDataEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, index uint32, type_ uint16, options uint64) GTMioHeatmapBuilder {
	instance := getGTMioHeatmapBuilderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:programType:options:"), data, index, type_, options)
	return GTMioHeatmapBuilderFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:initPC:encoderFunctionIndex:programType:options:
func NewGTMioHeatmapBuilderWithTraceDataInitPCEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, pc uint64, index uint32, type_ uint16, options uint64) GTMioHeatmapBuilder {
	instance := getGTMioHeatmapBuilderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:initPC:encoderFunctionIndex:programType:options:"), data, pc, index, type_, options)
	return GTMioHeatmapBuilderFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/buildCliqueHeatmap:width:height:generationOptions:
func (g GTMioHeatmapBuilder) BuildCliqueHeatmapWidthHeightGenerationOptions(heatmap uint64, width uint64, height uint64, options uint) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buildCliqueHeatmap:width:height:generationOptions:"), heatmap, width, height, options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/buildHeatmap:width:height:generationOptions:
func (g GTMioHeatmapBuilder) BuildHeatmapWidthHeightGenerationOptions(heatmap uint64, width uint64, height uint64, options uint) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buildHeatmap:width:height:generationOptions:"), heatmap, width, height, options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/buildQuadData
func (g GTMioHeatmapBuilder) BuildQuadData() {
	objc.Send[objc.ID](g.ID, objc.Sel("buildQuadData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/buildTileHeatmap:width:height:generationOptions:
func (g GTMioHeatmapBuilder) BuildTileHeatmapWidthHeightGenerationOptions(heatmap uint64, width uint64, height uint64, options uint) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buildTileHeatmap:width:height:generationOptions:"), heatmap, width, height, options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/heatmapSize:height:wrap:
func (g GTMioHeatmapBuilder) HeatmapSizeHeightWrap(size uint64, height uint64, wrap bool) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](g.ID, objc.Sel("heatmapSize:height:wrap:"), size, height, wrap)
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:
func (g GTMioHeatmapBuilder) InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions(data objectivec.IObject, index uint32, index2 uint32, type_ uint16, options uint64) GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](g.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:"), data, index, index2, type_, options)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:
func (g GTMioHeatmapBuilder) InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions(data objectivec.IObject, index uint32, id uint64, type_ uint16, options uint64) GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](g.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:"), data, index, id, type_, options)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:encoderFunctionIndex:programType:options:
func (g GTMioHeatmapBuilder) InitWithTraceDataEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, index uint32, type_ uint16, options uint64) GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](g.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:programType:options:"), data, index, type_, options)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/initWithTraceData:initPC:encoderFunctionIndex:programType:options:
func (g GTMioHeatmapBuilder) InitWithTraceDataInitPCEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, pc uint64, index uint32, type_ uint16, options uint64) GTMioHeatmapBuilder {
	rv := objc.Send[GTMioHeatmapBuilder](g.ID, objc.Sel("initWithTraceData:initPC:encoderFunctionIndex:programType:options:"), data, pc, index, type_, options)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/depth
func (g GTMioHeatmapBuilder) Depth() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("depth"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/drawIndex
func (g GTMioHeatmapBuilder) DrawIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("drawIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/encoderInfo
func (g GTMioHeatmapBuilder) EncoderInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoderInfo"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/options
func (g GTMioHeatmapBuilder) Options() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/pipelineStateId
func (g GTMioHeatmapBuilder) PipelineStateId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateId"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/programType
func (g GTMioHeatmapBuilder) ProgramType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("programType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmapBuilder/quadData
func (g GTMioHeatmapBuilder) QuadData() IGTMioEncoderQuadData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("quadData"))
	return GTMioEncoderQuadDataFromID(objc.ID(rv))
}
