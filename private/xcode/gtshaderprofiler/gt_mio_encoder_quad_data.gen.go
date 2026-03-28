// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioEncoderQuadData] class.
var (
	_GTMioEncoderQuadDataClass     GTMioEncoderQuadDataClass
	_GTMioEncoderQuadDataClassOnce sync.Once
)

func getGTMioEncoderQuadDataClass() GTMioEncoderQuadDataClass {
	_GTMioEncoderQuadDataClassOnce.Do(func() {
		_GTMioEncoderQuadDataClass = GTMioEncoderQuadDataClass{class: objc.GetClass("GTMioEncoderQuadData")}
	})
	return _GTMioEncoderQuadDataClass
}

// GetGTMioEncoderQuadDataClass returns the class object for GTMioEncoderQuadData.
func GetGTMioEncoderQuadDataClass() GTMioEncoderQuadDataClass {
	return getGTMioEncoderQuadDataClass()
}

type GTMioEncoderQuadDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioEncoderQuadDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioEncoderQuadDataClass) Alloc() GTMioEncoderQuadData {
	rv := objc.Send[GTMioEncoderQuadData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioEncoderQuadData._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [GTMioEncoderQuadData._handleClique]
//   - [GTMioEncoderQuadData.BuildEncoderFunctionIndexCliqueFilter]
//   - [GTMioEncoderQuadData.CliqueIndexesForQuadCount]
//   - [GTMioEncoderQuadData.CliqueIndexesForQuadLocationCount]
//   - [GTMioEncoderQuadData.ContainsDraw]
//   - [GTMioEncoderQuadData.Depth]
//   - [GTMioEncoderQuadData.DrawCount]
//   - [GTMioEncoderQuadData.DrawIndexes]
//   - [GTMioEncoderQuadData.DrawIndexesForQuad]
//   - [GTMioEncoderQuadData.DrawIndexesForQuadLocation]
//   - [GTMioEncoderQuadData.EncoderInfo]
//   - [GTMioEncoderQuadData.EnumerateCliquesForQuadEnumerator]
//   - [GTMioEncoderQuadData.EnumerateCliquesForQuadLocationEnumerator]
//   - [GTMioEncoderQuadData.EnumerateOrderedQuads]
//   - [GTMioEncoderQuadData.HeatmapType]
//   - [GTMioEncoderQuadData.Height]
//   - [GTMioEncoderQuadData.InstructionsExecutedForQuadLocationThreadInstructionsExectuedActiveThreadInstructionsExecutedTotalInstructionsExectuedNumActiveThreads]
//   - [GTMioEncoderQuadData.MaxCost]
//   - [GTMioEncoderQuadData.MaxTimestamp]
//   - [GTMioEncoderQuadData.MinCost]
//   - [GTMioEncoderQuadData.MinTimestamp]
//   - [GTMioEncoderQuadData.Options]
//   - [GTMioEncoderQuadData.ProgramType]
//   - [GTMioEncoderQuadData.QuadCount]
//   - [GTMioEncoderQuadData.QuadIndexForQuad]
//   - [GTMioEncoderQuadData.Quads]
//   - [GTMioEncoderQuadData.ReferenceComputePosition]
//   - [GTMioEncoderQuadData.TraceData]
//   - [GTMioEncoderQuadData.Width]
//   - [GTMioEncoderQuadData.InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions]
//   - [GTMioEncoderQuadData.InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions]
//   - [GTMioEncoderQuadData.InitWithTraceDataEncoderFunctionIndexProgramTypeOptions]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData
type GTMioEncoderQuadData struct {
	objectivec.Object
}

// GTMioEncoderQuadDataFromID constructs a [GTMioEncoderQuadData] from an objc.ID.
func GTMioEncoderQuadDataFromID(id objc.ID) GTMioEncoderQuadData {
	return GTMioEncoderQuadData{objectivec.Object{ID: id}}
}
// Ensure GTMioEncoderQuadData implements IGTMioEncoderQuadData.
var _ IGTMioEncoderQuadData = GTMioEncoderQuadData{}

// An interface definition for the [GTMioEncoderQuadData] class.
//
// # Methods
//
//   - [IGTMioEncoderQuadData._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [IGTMioEncoderQuadData._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [IGTMioEncoderQuadData._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter]
//   - [IGTMioEncoderQuadData._handleClique]
//   - [IGTMioEncoderQuadData.BuildEncoderFunctionIndexCliqueFilter]
//   - [IGTMioEncoderQuadData.CliqueIndexesForQuadCount]
//   - [IGTMioEncoderQuadData.CliqueIndexesForQuadLocationCount]
//   - [IGTMioEncoderQuadData.ContainsDraw]
//   - [IGTMioEncoderQuadData.Depth]
//   - [IGTMioEncoderQuadData.DrawCount]
//   - [IGTMioEncoderQuadData.DrawIndexes]
//   - [IGTMioEncoderQuadData.DrawIndexesForQuad]
//   - [IGTMioEncoderQuadData.DrawIndexesForQuadLocation]
//   - [IGTMioEncoderQuadData.EncoderInfo]
//   - [IGTMioEncoderQuadData.EnumerateCliquesForQuadEnumerator]
//   - [IGTMioEncoderQuadData.EnumerateCliquesForQuadLocationEnumerator]
//   - [IGTMioEncoderQuadData.EnumerateOrderedQuads]
//   - [IGTMioEncoderQuadData.HeatmapType]
//   - [IGTMioEncoderQuadData.Height]
//   - [IGTMioEncoderQuadData.InstructionsExecutedForQuadLocationThreadInstructionsExectuedActiveThreadInstructionsExecutedTotalInstructionsExectuedNumActiveThreads]
//   - [IGTMioEncoderQuadData.MaxCost]
//   - [IGTMioEncoderQuadData.MaxTimestamp]
//   - [IGTMioEncoderQuadData.MinCost]
//   - [IGTMioEncoderQuadData.MinTimestamp]
//   - [IGTMioEncoderQuadData.Options]
//   - [IGTMioEncoderQuadData.ProgramType]
//   - [IGTMioEncoderQuadData.QuadCount]
//   - [IGTMioEncoderQuadData.QuadIndexForQuad]
//   - [IGTMioEncoderQuadData.Quads]
//   - [IGTMioEncoderQuadData.ReferenceComputePosition]
//   - [IGTMioEncoderQuadData.TraceData]
//   - [IGTMioEncoderQuadData.Width]
//   - [IGTMioEncoderQuadData.InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions]
//   - [IGTMioEncoderQuadData.InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions]
//   - [IGTMioEncoderQuadData.InitWithTraceDataEncoderFunctionIndexProgramTypeOptions]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData
type IGTMioEncoderQuadData interface {
	objectivec.IObject

	// Topic: Methods

	_buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter(cliques objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool
	_buildComputeEncoderFunctionIndexProgramTypeCliqueFilter(compute objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool
	_buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter(fragment objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool
	_handleClique(clique unsafe.Pointer)
	BuildEncoderFunctionIndexCliqueFilter(build objectivec.IObject, index uint32, filter VoidHandler) bool
	CliqueIndexesForQuadCount(quad []objectivec.IObject, count unsafe.Pointer) unsafe.Pointer
	CliqueIndexesForQuadLocationCount(location uint64, count unsafe.Pointer) unsafe.Pointer
	ContainsDraw(draw uint32) bool
	Depth() uint32
	DrawCount() uint64
	DrawIndexes() unsafe.Pointer
	DrawIndexesForQuad(quad unsafe.Pointer) objectivec.IObject
	DrawIndexesForQuadLocation(location uint64) objectivec.IObject
	EncoderInfo() unsafe.Pointer
	EnumerateCliquesForQuadEnumerator(quad unsafe.Pointer, enumerator VoidHandler)
	EnumerateCliquesForQuadLocationEnumerator(location uint64, enumerator VoidHandler)
	EnumerateOrderedQuads(quads VoidHandler)
	HeatmapType() uint64
	Height() uint32
	InstructionsExecutedForQuadLocationThreadInstructionsExectuedActiveThreadInstructionsExecutedTotalInstructionsExectuedNumActiveThreads(location uint64, exectued unsafe.Pointer, executed unsafe.Pointer, exectued2 unsafe.Pointer, threads unsafe.Pointer)
	MaxCost() float64
	MaxTimestamp() uint64
	MinCost() float64
	MinTimestamp() uint64
	Options() uint64
	ProgramType() uint16
	QuadCount() uint64
	QuadIndexForQuad(quad unsafe.Pointer) uint32
	Quads() unsafe.Pointer
	ReferenceComputePosition() unsafe.Pointer
	TraceData() objectivec.IObject
	Width() uint32
	InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions(data objectivec.IObject, index uint32, index2 uint32, type_ uint16, options uint64) GTMioEncoderQuadData
	InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions(data objectivec.IObject, index uint32, id uint64, type_ uint16, options uint64) GTMioEncoderQuadData
	InitWithTraceDataEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, index uint32, type_ uint16, options uint64) GTMioEncoderQuadData
}

// Init initializes the instance.
func (g GTMioEncoderQuadData) Init() GTMioEncoderQuadData {
	rv := objc.Send[GTMioEncoderQuadData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioEncoderQuadData) Autorelease() GTMioEncoderQuadData {
	rv := objc.Send[GTMioEncoderQuadData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioEncoderQuadData creates a new GTMioEncoderQuadData instance.
func NewGTMioEncoderQuadData() GTMioEncoderQuadData {
	class := getGTMioEncoderQuadDataClass()
	rv := objc.Send[GTMioEncoderQuadData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:
func NewGTMioEncoderQuadDataWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions(data objectivec.IObject, index uint32, index2 uint32, type_ uint16, options uint64) GTMioEncoderQuadData {
	instance := getGTMioEncoderQuadDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:"), data, index, index2, type_, options)
	return GTMioEncoderQuadDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:
func NewGTMioEncoderQuadDataWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions(data objectivec.IObject, index uint32, id uint64, type_ uint16, options uint64) GTMioEncoderQuadData {
	instance := getGTMioEncoderQuadDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:"), data, index, id, type_, options)
	return GTMioEncoderQuadDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/initWithTraceData:encoderFunctionIndex:programType:options:
func NewGTMioEncoderQuadDataWithTraceDataEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, index uint32, type_ uint16, options uint64) GTMioEncoderQuadData {
	instance := getGTMioEncoderQuadDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:programType:options:"), data, index, type_, options)
	return GTMioEncoderQuadDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/_buildCliques:encoderFunctionIndex:programType:cliqueFilter:
func (g GTMioEncoderQuadData) _buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter(cliques objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool {
_block3, _ := NewVoidBlock(filter)
	rv := objc.Send[bool](g.ID, objc.Sel("_buildCliques:encoderFunctionIndex:programType:cliqueFilter:"), cliques, index, type_, _block3)
	return rv
}

// BuildCliquesEncoderFunctionIndexProgramTypeCliqueFilter is an exported wrapper for the private method _buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter.
func (g GTMioEncoderQuadData) BuildCliquesEncoderFunctionIndexProgramTypeCliqueFilter(cliques objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool {
	return g._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter(cliques, index, type_, filter)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/_buildCompute:encoderFunctionIndex:programType:cliqueFilter:
func (g GTMioEncoderQuadData) _buildComputeEncoderFunctionIndexProgramTypeCliqueFilter(compute objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool {
_block3, _ := NewVoidBlock(filter)
	rv := objc.Send[bool](g.ID, objc.Sel("_buildCompute:encoderFunctionIndex:programType:cliqueFilter:"), compute, index, type_, _block3)
	return rv
}

// BuildComputeEncoderFunctionIndexProgramTypeCliqueFilter is an exported wrapper for the private method _buildComputeEncoderFunctionIndexProgramTypeCliqueFilter.
func (g GTMioEncoderQuadData) BuildComputeEncoderFunctionIndexProgramTypeCliqueFilter(compute objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool {
	return g._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter(compute, index, type_, filter)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/_buildFragment:encoderFunctionIndex:programType:cliqueFilter:
func (g GTMioEncoderQuadData) _buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter(fragment objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool {
_block3, _ := NewVoidBlock(filter)
	rv := objc.Send[bool](g.ID, objc.Sel("_buildFragment:encoderFunctionIndex:programType:cliqueFilter:"), fragment, index, type_, _block3)
	return rv
}

// BuildFragmentEncoderFunctionIndexProgramTypeCliqueFilter is an exported wrapper for the private method _buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter.
func (g GTMioEncoderQuadData) BuildFragmentEncoderFunctionIndexProgramTypeCliqueFilter(fragment objectivec.IObject, index uint32, type_ uint16, filter VoidHandler) bool {
	return g._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter(fragment, index, type_, filter)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/_handleClique:
func (g GTMioEncoderQuadData) _handleClique(clique unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("_handleClique:"), clique)
}

// HandleClique is an exported wrapper for the private method _handleClique.
func (g GTMioEncoderQuadData) HandleClique(clique unsafe.Pointer) {
	g._handleClique(clique)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/build:encoderFunctionIndex:cliqueFilter:
func (g GTMioEncoderQuadData) BuildEncoderFunctionIndexCliqueFilter(build objectivec.IObject, index uint32, filter VoidHandler) bool {
_block2, _ := NewVoidBlock(filter)
	rv := objc.Send[bool](g.ID, objc.Sel("build:encoderFunctionIndex:cliqueFilter:"), build, index, _block2)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/cliqueIndexesForQuad:count:
func (g GTMioEncoderQuadData) CliqueIndexesForQuadCount(quad []objectivec.IObject, count unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cliqueIndexesForQuad:count:"), objc.CArray(quad), count)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/cliqueIndexesForQuadLocation:count:
func (g GTMioEncoderQuadData) CliqueIndexesForQuadLocationCount(location uint64, count unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("cliqueIndexesForQuadLocation:count:"), location, count)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/containsDraw:
func (g GTMioEncoderQuadData) ContainsDraw(draw uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("containsDraw:"), draw)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/drawIndexesForQuad:
func (g GTMioEncoderQuadData) DrawIndexesForQuad(quad unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("drawIndexesForQuad:"), quad)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/drawIndexesForQuadLocation:
func (g GTMioEncoderQuadData) DrawIndexesForQuadLocation(location uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("drawIndexesForQuadLocation:"), location)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/enumerateCliquesForQuad:enumerator:
func (g GTMioEncoderQuadData) EnumerateCliquesForQuadEnumerator(quad unsafe.Pointer, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateCliquesForQuad:enumerator:"), quad, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/enumerateCliquesForQuadLocation:enumerator:
func (g GTMioEncoderQuadData) EnumerateCliquesForQuadLocationEnumerator(location uint64, enumerator VoidHandler) {
_block1, _ := NewVoidBlock(enumerator)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateCliquesForQuadLocation:enumerator:"), location, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/enumerateOrderedQuads:
func (g GTMioEncoderQuadData) EnumerateOrderedQuads(quads VoidHandler) {
_block0, _ := NewVoidBlock(quads)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateOrderedQuads:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/instructionsExecutedForQuadLocation:threadInstructionsExectued:activeThreadInstructionsExecuted:totalInstructionsExectued:numActiveThreads:
func (g GTMioEncoderQuadData) InstructionsExecutedForQuadLocationThreadInstructionsExectuedActiveThreadInstructionsExecutedTotalInstructionsExectuedNumActiveThreads(location uint64, exectued unsafe.Pointer, executed unsafe.Pointer, exectued2 unsafe.Pointer, threads unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("instructionsExecutedForQuadLocation:threadInstructionsExectued:activeThreadInstructionsExecuted:totalInstructionsExectued:numActiveThreads:"), location, exectued, executed, exectued2, threads)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/quadIndexForQuad:
func (g GTMioEncoderQuadData) QuadIndexForQuad(quad unsafe.Pointer) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("quadIndexForQuad:"), quad)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:
func (g GTMioEncoderQuadData) InitWithTraceDataEncoderFunctionIndexDrawIndexProgramTypeOptions(data objectivec.IObject, index uint32, index2 uint32, type_ uint16, options uint64) GTMioEncoderQuadData {
	rv := objc.Send[GTMioEncoderQuadData](g.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:drawIndex:programType:options:"), data, index, index2, type_, options)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:
func (g GTMioEncoderQuadData) InitWithTraceDataEncoderFunctionIndexPipelineStateIdProgramTypeOptions(data objectivec.IObject, index uint32, id uint64, type_ uint16, options uint64) GTMioEncoderQuadData {
	rv := objc.Send[GTMioEncoderQuadData](g.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:pipelineStateId:programType:options:"), data, index, id, type_, options)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/initWithTraceData:encoderFunctionIndex:programType:options:
func (g GTMioEncoderQuadData) InitWithTraceDataEncoderFunctionIndexProgramTypeOptions(data objectivec.IObject, index uint32, type_ uint16, options uint64) GTMioEncoderQuadData {
	rv := objc.Send[GTMioEncoderQuadData](g.ID, objc.Sel("initWithTraceData:encoderFunctionIndex:programType:options:"), data, index, type_, options)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/depth
func (g GTMioEncoderQuadData) Depth() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("depth"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/drawCount
func (g GTMioEncoderQuadData) DrawCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("drawCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/drawIndexes
func (g GTMioEncoderQuadData) DrawIndexes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("drawIndexes"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/encoderInfo
func (g GTMioEncoderQuadData) EncoderInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoderInfo"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/heatmapType
func (g GTMioEncoderQuadData) HeatmapType() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("heatmapType"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/height
func (g GTMioEncoderQuadData) Height() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("height"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/maxCost
func (g GTMioEncoderQuadData) MaxCost() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("maxCost"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/maxTimestamp
func (g GTMioEncoderQuadData) MaxTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("maxTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/minCost
func (g GTMioEncoderQuadData) MinCost() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("minCost"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/minTimestamp
func (g GTMioEncoderQuadData) MinTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("minTimestamp"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/options
func (g GTMioEncoderQuadData) Options() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("options"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/programType
func (g GTMioEncoderQuadData) ProgramType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("programType"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/quadCount
func (g GTMioEncoderQuadData) QuadCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("quadCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/quads
func (g GTMioEncoderQuadData) Quads() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("quads"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/referenceComputePosition
func (g GTMioEncoderQuadData) ReferenceComputePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("referenceComputePosition"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/traceData
func (g GTMioEncoderQuadData) TraceData() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("traceData"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioEncoderQuadData/width
func (g GTMioEncoderQuadData) Width() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("width"))
	return rv
}

// _buildCliquesEncoderFunctionIndexProgramTypeCliqueFilterSync is a synchronous wrapper around [GTMioEncoderQuadData._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) _buildCliquesEncoderFunctionIndexProgramTypeCliqueFilterSync(ctx context.Context, cliques objectivec.IObject, index uint32, type_ uint16) error {
	done := make(chan struct{}, 1)
	g._buildCliquesEncoderFunctionIndexProgramTypeCliqueFilter(cliques, index, type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _buildComputeEncoderFunctionIndexProgramTypeCliqueFilterSync is a synchronous wrapper around [GTMioEncoderQuadData._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) _buildComputeEncoderFunctionIndexProgramTypeCliqueFilterSync(ctx context.Context, compute objectivec.IObject, index uint32, type_ uint16) error {
	done := make(chan struct{}, 1)
	g._buildComputeEncoderFunctionIndexProgramTypeCliqueFilter(compute, index, type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// _buildFragmentEncoderFunctionIndexProgramTypeCliqueFilterSync is a synchronous wrapper around [GTMioEncoderQuadData._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) _buildFragmentEncoderFunctionIndexProgramTypeCliqueFilterSync(ctx context.Context, fragment objectivec.IObject, index uint32, type_ uint16) error {
	done := make(chan struct{}, 1)
	g._buildFragmentEncoderFunctionIndexProgramTypeCliqueFilter(fragment, index, type_, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// BuildEncoderFunctionIndexCliqueFilterSync is a synchronous wrapper around [GTMioEncoderQuadData.BuildEncoderFunctionIndexCliqueFilter].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) BuildEncoderFunctionIndexCliqueFilterSync(ctx context.Context, build objectivec.IObject, index uint32) error {
	done := make(chan struct{}, 1)
	g.BuildEncoderFunctionIndexCliqueFilter(build, index, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateCliquesForQuadEnumeratorSync is a synchronous wrapper around [GTMioEncoderQuadData.EnumerateCliquesForQuadEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) EnumerateCliquesForQuadEnumeratorSync(ctx context.Context, quad unsafe.Pointer) error {
	done := make(chan struct{}, 1)
	g.EnumerateCliquesForQuadEnumerator(quad, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateCliquesForQuadLocationEnumeratorSync is a synchronous wrapper around [GTMioEncoderQuadData.EnumerateCliquesForQuadLocationEnumerator].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) EnumerateCliquesForQuadLocationEnumeratorSync(ctx context.Context, location uint64) error {
	done := make(chan struct{}, 1)
	g.EnumerateCliquesForQuadLocationEnumerator(location, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateOrderedQuadsSync is a synchronous wrapper around [GTMioEncoderQuadData.EnumerateOrderedQuads].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioEncoderQuadData) EnumerateOrderedQuadsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateOrderedQuads(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

