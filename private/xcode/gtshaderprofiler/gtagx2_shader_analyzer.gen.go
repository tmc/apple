// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderAnalyzer] class.
var (
	_GTAGX2ShaderAnalyzerClass     GTAGX2ShaderAnalyzerClass
	_GTAGX2ShaderAnalyzerClassOnce sync.Once
)

func getGTAGX2ShaderAnalyzerClass() GTAGX2ShaderAnalyzerClass {
	_GTAGX2ShaderAnalyzerClassOnce.Do(func() {
		_GTAGX2ShaderAnalyzerClass = GTAGX2ShaderAnalyzerClass{class: objc.GetClass("GTAGX2ShaderAnalyzer")}
	})
	return _GTAGX2ShaderAnalyzerClass
}

// GetGTAGX2ShaderAnalyzerClass returns the class object for GTAGX2ShaderAnalyzer.
func GetGTAGX2ShaderAnalyzerClass() GTAGX2ShaderAnalyzerClass {
	return getGTAGX2ShaderAnalyzerClass()
}

type GTAGX2ShaderAnalyzerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderAnalyzerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderAnalyzerClass) Alloc() GTAGX2ShaderAnalyzer {
	rv := objc.Send[GTAGX2ShaderAnalyzer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderAnalyzer._calculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary]
//   - [GTAGX2ShaderAnalyzer._calculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo]
//   - [GTAGX2ShaderAnalyzer.AnalyzedBinaryProcessedUscSamplesTargetIndexWithALUBlocksBinaryInfo]
//   - [GTAGX2ShaderAnalyzer.DisassembleBinary]
//   - [GTAGX2ShaderAnalyzer.GetShaderBinaryForTargetIndexBinaryInfo]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer
type GTAGX2ShaderAnalyzer struct {
	objectivec.Object
}

// GTAGX2ShaderAnalyzerFromID constructs a [GTAGX2ShaderAnalyzer] from an objc.ID.
func GTAGX2ShaderAnalyzerFromID(id objc.ID) GTAGX2ShaderAnalyzer {
	return GTAGX2ShaderAnalyzer{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderAnalyzer implements IGTAGX2ShaderAnalyzer.
var _ IGTAGX2ShaderAnalyzer = GTAGX2ShaderAnalyzer{}

// An interface definition for the [GTAGX2ShaderAnalyzer] class.
//
// # Methods
//
//   - [IGTAGX2ShaderAnalyzer._calculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary]
//   - [IGTAGX2ShaderAnalyzer._calculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo]
//   - [IGTAGX2ShaderAnalyzer.AnalyzedBinaryProcessedUscSamplesTargetIndexWithALUBlocksBinaryInfo]
//   - [IGTAGX2ShaderAnalyzer.DisassembleBinary]
//   - [IGTAGX2ShaderAnalyzer.GetShaderBinaryForTargetIndexBinaryInfo]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer
type IGTAGX2ShaderAnalyzer interface {
	objectivec.IObject

	// Topic: Methods

	_calculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary(data objectivec.IObject, index int, binary objectivec.IObject)
	_calculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo(data objectivec.IObject, result objectivec.IObject, index int, aLUBlocks unsafe.Pointer, info objectivec.IObject) objectivec.IObject
	AnalyzedBinaryProcessedUscSamplesTargetIndexWithALUBlocksBinaryInfo(binary objectivec.IObject, samples objectivec.IObject, index int, aLUBlocks unsafe.Pointer, info objectivec.IObject) objectivec.IObject
	DisassembleBinary(binary objectivec.IObject) objectivec.IObject
	GetShaderBinaryForTargetIndexBinaryInfo(binary objectivec.IObject, index int, info objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (g GTAGX2ShaderAnalyzer) Init() GTAGX2ShaderAnalyzer {
	rv := objc.Send[GTAGX2ShaderAnalyzer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderAnalyzer) Autorelease() GTAGX2ShaderAnalyzer {
	rv := objc.Send[GTAGX2ShaderAnalyzer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderAnalyzer creates a new GTAGX2ShaderAnalyzer instance.
func NewGTAGX2ShaderAnalyzer() GTAGX2ShaderAnalyzer {
	class := getGTAGX2ShaderAnalyzerClass()
	rv := objc.Send[GTAGX2ShaderAnalyzer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer/_calculatePerDrawCallWithGRCSampleData:targetIndex:shaderBinary:
func (g GTAGX2ShaderAnalyzer) _calculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary(data objectivec.IObject, index int, binary objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_calculatePerDrawCallWithGRCSampleData:targetIndex:shaderBinary:"), data, index, binary)
}

// CalculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary is an exported wrapper for the private method _calculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary.
func (g GTAGX2ShaderAnalyzer) CalculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary(data objectivec.IObject, index int, binary objectivec.IObject) {
	g._calculatePerDrawCallWithGRCSampleDataTargetIndexShaderBinary(data, index, binary)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer/_calculatePerLineCostWithSampleData:analysisResult:targetIndex:withALUBlocks:binaryInfo:
func (g GTAGX2ShaderAnalyzer) _calculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo(data objectivec.IObject, result objectivec.IObject, index int, aLUBlocks unsafe.Pointer, info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_calculatePerLineCostWithSampleData:analysisResult:targetIndex:withALUBlocks:binaryInfo:"), data, result, index, aLUBlocks, info)
	return objectivec.Object{ID: rv}
}

// CalculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo is an exported wrapper for the private method _calculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo.
func (g GTAGX2ShaderAnalyzer) CalculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo(data objectivec.IObject, result objectivec.IObject, index int, aLUBlocks unsafe.Pointer, info objectivec.IObject) objectivec.IObject {
	return g._calculatePerLineCostWithSampleDataAnalysisResultTargetIndexWithALUBlocksBinaryInfo(data, result, index, aLUBlocks, info)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer/analyzedBinary:processedUscSamples:targetIndex:withALUBlocks:binaryInfo:
func (g GTAGX2ShaderAnalyzer) AnalyzedBinaryProcessedUscSamplesTargetIndexWithALUBlocksBinaryInfo(binary objectivec.IObject, samples objectivec.IObject, index int, aLUBlocks unsafe.Pointer, info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("analyzedBinary:processedUscSamples:targetIndex:withALUBlocks:binaryInfo:"), binary, samples, index, aLUBlocks, info)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer/disassembleBinary:
func (g GTAGX2ShaderAnalyzer) DisassembleBinary(binary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("disassembleBinary:"), binary)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderAnalyzer/getShaderBinary:forTargetIndex:binaryInfo:
func (g GTAGX2ShaderAnalyzer) GetShaderBinaryForTargetIndexBinaryInfo(binary objectivec.IObject, index int, info objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getShaderBinary:forTargetIndex:binaryInfo:"), binary, index, info)
	return objectivec.Object{ID: rv}
}

