// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioTraceDataHelper] class.
var (
	_GTMioTraceDataHelperClass     GTMioTraceDataHelperClass
	_GTMioTraceDataHelperClassOnce sync.Once
)

func getGTMioTraceDataHelperClass() GTMioTraceDataHelperClass {
	_GTMioTraceDataHelperClassOnce.Do(func() {
		_GTMioTraceDataHelperClass = GTMioTraceDataHelperClass{class: objc.GetClass("GTMioTraceDataHelper")}
	})
	return _GTMioTraceDataHelperClass
}

// GetGTMioTraceDataHelperClass returns the class object for GTMioTraceDataHelper.
func GetGTMioTraceDataHelperClass() GTMioTraceDataHelperClass {
	return getGTMioTraceDataHelperClass()
}

type GTMioTraceDataHelperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioTraceDataHelperClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioTraceDataHelperClass) Alloc() GTMioTraceDataHelper {
	rv := objc.Send[GTMioTraceDataHelper](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioTraceDataHelper._cachePerEncoderShaderTracks]
//   - [GTMioTraceDataHelper.DoNotMergeProgramTypes]
//   - [GTMioTraceDataHelper.SetDoNotMergeProgramTypes]
//   - [GTMioTraceDataHelper.GenerateAggregatedCliqueTrackForUSC]
//   - [GTMioTraceDataHelper.GenerateAggregatedDrawTrackForEncoder]
//   - [GTMioTraceDataHelper.GenerateAggregatedPerBinaryShaderTrackForEncoder]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForDataMaster]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForDataMasterWithBinaries]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForDataMasterWithCliques]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForEncoder]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateDataMaster]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndex]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyDataMaster]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyProgramType]
//   - [GTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateProgramType]
//   - [GTMioTraceDataHelper.GenerateBinaryTrack]
//   - [GTMioTraceDataHelper.GenerateCliqueInstructionTracksForShaderProgramTypeUscIndex]
//   - [GTMioTraceDataHelper.GenerateCliqueInstructionTracksForUSC]
//   - [GTMioTraceDataHelper.GenerateCliqueTracksForUSC]
//   - [GTMioTraceDataHelper.GenerateEncoderShaderTracks]
//   - [GTMioTraceDataHelper.GenerateIndividualShaderTrackForProgramTypes]
//   - [GTMioTraceDataHelper.GenerateKickTracksForMGPU]
//   - [GTMioTraceDataHelper.GenerateKickTracksForUSC]
//   - [GTMioTraceDataHelper.GenerateShaderTrackForProgramType]
//   - [GTMioTraceDataHelper.GenerateShaderTrackForProgramTypes]
//   - [GTMioTraceDataHelper.GenerateShaderTrackForProgramTypesForUSC]
//   - [GTMioTraceDataHelper.GenerateShaderTracksForPipelineState]
//   - [GTMioTraceDataHelper.GenerateTileTracksForUSC]
//   - [GTMioTraceDataHelper.GenerateTopBinaryTracks]
//   - [GTMioTraceDataHelper.GenerateTopDrawTracks]
//   - [GTMioTraceDataHelper.GenerateTopKickTracks]
//   - [GTMioTraceDataHelper.GenerateTopRIATracks]
//   - [GTMioTraceDataHelper.GenerateTrackForCliqueIndexesCountGroup]
//   - [GTMioTraceDataHelper.ShowDriverInternalShaders]
//   - [GTMioTraceDataHelper.SetShowDriverInternalShaders]
//   - [GTMioTraceDataHelper.ShowDriverIntersectionShaders]
//   - [GTMioTraceDataHelper.SetShowDriverIntersectionShaders]
//   - [GTMioTraceDataHelper.ShowESLShaders]
//   - [GTMioTraceDataHelper.SetShowESLShaders]
//   - [GTMioTraceDataHelper.Stats]
//   - [GTMioTraceDataHelper.TraceData]
//   - [GTMioTraceDataHelper.InitWithTraceData]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper
type GTMioTraceDataHelper struct {
	objectivec.Object
}

// GTMioTraceDataHelperFromID constructs a [GTMioTraceDataHelper] from an objc.ID.
func GTMioTraceDataHelperFromID(id objc.ID) GTMioTraceDataHelper {
	return GTMioTraceDataHelper{objectivec.Object{ID: id}}
}
// Ensure GTMioTraceDataHelper implements IGTMioTraceDataHelper.
var _ IGTMioTraceDataHelper = GTMioTraceDataHelper{}

// An interface definition for the [GTMioTraceDataHelper] class.
//
// # Methods
//
//   - [IGTMioTraceDataHelper._cachePerEncoderShaderTracks]
//   - [IGTMioTraceDataHelper.DoNotMergeProgramTypes]
//   - [IGTMioTraceDataHelper.SetDoNotMergeProgramTypes]
//   - [IGTMioTraceDataHelper.GenerateAggregatedCliqueTrackForUSC]
//   - [IGTMioTraceDataHelper.GenerateAggregatedDrawTrackForEncoder]
//   - [IGTMioTraceDataHelper.GenerateAggregatedPerBinaryShaderTrackForEncoder]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForDataMaster]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForDataMasterWithBinaries]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForDataMasterWithCliques]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForEncoder]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateDataMaster]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndex]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyDataMaster]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyProgramType]
//   - [IGTMioTraceDataHelper.GenerateAggregatedShaderTrackForPipelineStateProgramType]
//   - [IGTMioTraceDataHelper.GenerateBinaryTrack]
//   - [IGTMioTraceDataHelper.GenerateCliqueInstructionTracksForShaderProgramTypeUscIndex]
//   - [IGTMioTraceDataHelper.GenerateCliqueInstructionTracksForUSC]
//   - [IGTMioTraceDataHelper.GenerateCliqueTracksForUSC]
//   - [IGTMioTraceDataHelper.GenerateEncoderShaderTracks]
//   - [IGTMioTraceDataHelper.GenerateIndividualShaderTrackForProgramTypes]
//   - [IGTMioTraceDataHelper.GenerateKickTracksForMGPU]
//   - [IGTMioTraceDataHelper.GenerateKickTracksForUSC]
//   - [IGTMioTraceDataHelper.GenerateShaderTrackForProgramType]
//   - [IGTMioTraceDataHelper.GenerateShaderTrackForProgramTypes]
//   - [IGTMioTraceDataHelper.GenerateShaderTrackForProgramTypesForUSC]
//   - [IGTMioTraceDataHelper.GenerateShaderTracksForPipelineState]
//   - [IGTMioTraceDataHelper.GenerateTileTracksForUSC]
//   - [IGTMioTraceDataHelper.GenerateTopBinaryTracks]
//   - [IGTMioTraceDataHelper.GenerateTopDrawTracks]
//   - [IGTMioTraceDataHelper.GenerateTopKickTracks]
//   - [IGTMioTraceDataHelper.GenerateTopRIATracks]
//   - [IGTMioTraceDataHelper.GenerateTrackForCliqueIndexesCountGroup]
//   - [IGTMioTraceDataHelper.ShowDriverInternalShaders]
//   - [IGTMioTraceDataHelper.SetShowDriverInternalShaders]
//   - [IGTMioTraceDataHelper.ShowDriverIntersectionShaders]
//   - [IGTMioTraceDataHelper.SetShowDriverIntersectionShaders]
//   - [IGTMioTraceDataHelper.ShowESLShaders]
//   - [IGTMioTraceDataHelper.SetShowESLShaders]
//   - [IGTMioTraceDataHelper.Stats]
//   - [IGTMioTraceDataHelper.TraceData]
//   - [IGTMioTraceDataHelper.InitWithTraceData]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper
type IGTMioTraceDataHelper interface {
	objectivec.IObject

	// Topic: Methods

	_cachePerEncoderShaderTracks()
	DoNotMergeProgramTypes() bool
	SetDoNotMergeProgramTypes(value bool)
	GenerateAggregatedCliqueTrackForUSC(usc uint32) objectivec.IObject
	GenerateAggregatedDrawTrackForEncoder(encoder uint32) objectivec.IObject
	GenerateAggregatedPerBinaryShaderTrackForEncoder(encoder uint32) objectivec.IObject
	GenerateAggregatedShaderTrackForDataMaster(master uint16) objectivec.IObject
	GenerateAggregatedShaderTrackForDataMasterWithBinaries(binaries uint16) objectivec.IObject
	GenerateAggregatedShaderTrackForDataMasterWithCliques(cliques uint16) objectivec.IObject
	GenerateAggregatedShaderTrackForEncoder(encoder uint32) objectivec.IObject
	GenerateAggregatedShaderTrackForPipelineStateDataMaster(state uint64, master uint16) objectivec.IObject
	GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndex(state uint64, index uint32) objectivec.IObject
	GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyDataMaster(state uint64, index uint32, key uint64, master uint16) objectivec.IObject
	GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyProgramType(state uint64, index uint32, key uint64, type_ uint16) objectivec.IObject
	GenerateAggregatedShaderTrackForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject
	GenerateBinaryTrack(track uint32) objectivec.IObject
	GenerateCliqueInstructionTracksForShaderProgramTypeUscIndex(shader uint64, type_ uint16, index uint32) objectivec.IObject
	GenerateCliqueInstructionTracksForUSC(usc uint32) objectivec.IObject
	GenerateCliqueTracksForUSC(usc uint32) objectivec.IObject
	GenerateEncoderShaderTracks(tracks uint32) objectivec.IObject
	GenerateIndividualShaderTrackForProgramTypes() objectivec.IObject
	GenerateKickTracksForMGPU(mgpu uint32) objectivec.IObject
	GenerateKickTracksForUSC(usc uint32) objectivec.IObject
	GenerateShaderTrackForProgramType(type_ uint16) objectivec.IObject
	GenerateShaderTrackForProgramTypes() objectivec.IObject
	GenerateShaderTrackForProgramTypesForUSC(usc uint32) objectivec.IObject
	GenerateShaderTracksForPipelineState(state uint64) objectivec.IObject
	GenerateTileTracksForUSC(usc uint32) objectivec.IObject
	GenerateTopBinaryTracks() objectivec.IObject
	GenerateTopDrawTracks() objectivec.IObject
	GenerateTopKickTracks() objectivec.IObject
	GenerateTopRIATracks() objectivec.IObject
	GenerateTrackForCliqueIndexesCountGroup(indexes unsafe.Pointer, count uint64, group VoidHandler) objectivec.IObject
	ShowDriverInternalShaders() bool
	SetShowDriverInternalShaders(value bool)
	ShowDriverIntersectionShaders() bool
	SetShowDriverIntersectionShaders(value bool)
	ShowESLShaders() bool
	SetShowESLShaders(value bool)
	Stats() objectivec.IObject
	TraceData() objectivec.IObject
	InitWithTraceData(data objectivec.IObject) GTMioTraceDataHelper
}

// Init initializes the instance.
func (g GTMioTraceDataHelper) Init() GTMioTraceDataHelper {
	rv := objc.Send[GTMioTraceDataHelper](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioTraceDataHelper) Autorelease() GTMioTraceDataHelper {
	rv := objc.Send[GTMioTraceDataHelper](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioTraceDataHelper creates a new GTMioTraceDataHelper instance.
func NewGTMioTraceDataHelper() GTMioTraceDataHelper {
	class := getGTMioTraceDataHelperClass()
	rv := objc.Send[GTMioTraceDataHelper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/initWithTraceData:
func NewGTMioTraceDataHelperWithTraceData(data objectivec.IObject) GTMioTraceDataHelper {
	instance := getGTMioTraceDataHelperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTraceData:"), data)
	return GTMioTraceDataHelperFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/_cachePerEncoderShaderTracks
func (g GTMioTraceDataHelper) _cachePerEncoderShaderTracks() {
	objc.Send[objc.ID](g.ID, objc.Sel("_cachePerEncoderShaderTracks"))
}

// CachePerEncoderShaderTracks is an exported wrapper for the private method _cachePerEncoderShaderTracks.
func (g GTMioTraceDataHelper) CachePerEncoderShaderTracks() {
	g._cachePerEncoderShaderTracks()
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedCliqueTrackForUSC:
func (g GTMioTraceDataHelper) GenerateAggregatedCliqueTrackForUSC(usc uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedCliqueTrackForUSC:"), usc)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedDrawTrackForEncoder:
func (g GTMioTraceDataHelper) GenerateAggregatedDrawTrackForEncoder(encoder uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedDrawTrackForEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedPerBinaryShaderTrackForEncoder:
func (g GTMioTraceDataHelper) GenerateAggregatedPerBinaryShaderTrackForEncoder(encoder uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedPerBinaryShaderTrackForEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForDataMaster:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForDataMaster(master uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForDataMaster:"), master)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForDataMasterWithBinaries:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForDataMasterWithBinaries(binaries uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForDataMasterWithBinaries:"), binaries)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForDataMasterWithCliques:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForDataMasterWithCliques(cliques uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForDataMasterWithCliques:"), cliques)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForEncoder:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForEncoder(encoder uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForPipelineState:dataMaster:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForPipelineStateDataMaster(state uint64, master uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForPipelineState:dataMaster:"), state, master)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForPipelineState:encoderFunctionIndex:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndex(state uint64, index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForPipelineState:encoderFunctionIndex:"), state, index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForPipelineState:encoderFunctionIndex:cacheKey:dataMaster:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyDataMaster(state uint64, index uint32, key uint64, master uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForPipelineState:encoderFunctionIndex:cacheKey:dataMaster:"), state, index, key, master)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForPipelineState:encoderFunctionIndex:cacheKey:programType:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForPipelineStateEncoderFunctionIndexCacheKeyProgramType(state uint64, index uint32, key uint64, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForPipelineState:encoderFunctionIndex:cacheKey:programType:"), state, index, key, type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateAggregatedShaderTrackForPipelineState:programType:
func (g GTMioTraceDataHelper) GenerateAggregatedShaderTrackForPipelineStateProgramType(state uint64, type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateAggregatedShaderTrackForPipelineState:programType:"), state, type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateBinaryTrack:
func (g GTMioTraceDataHelper) GenerateBinaryTrack(track uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateBinaryTrack:"), track)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateCliqueInstructionTracksForShader:programType:uscIndex:
func (g GTMioTraceDataHelper) GenerateCliqueInstructionTracksForShaderProgramTypeUscIndex(shader uint64, type_ uint16, index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateCliqueInstructionTracksForShader:programType:uscIndex:"), shader, type_, index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateCliqueInstructionTracksForUSC:
func (g GTMioTraceDataHelper) GenerateCliqueInstructionTracksForUSC(usc uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateCliqueInstructionTracksForUSC:"), usc)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateCliqueTracksForUSC:
func (g GTMioTraceDataHelper) GenerateCliqueTracksForUSC(usc uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateCliqueTracksForUSC:"), usc)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateEncoderShaderTracks:
func (g GTMioTraceDataHelper) GenerateEncoderShaderTracks(tracks uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateEncoderShaderTracks:"), tracks)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateIndividualShaderTrackForProgramTypes
func (g GTMioTraceDataHelper) GenerateIndividualShaderTrackForProgramTypes() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateIndividualShaderTrackForProgramTypes"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateKickTracksForMGPU:
func (g GTMioTraceDataHelper) GenerateKickTracksForMGPU(mgpu uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateKickTracksForMGPU:"), mgpu)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateKickTracksForUSC:
func (g GTMioTraceDataHelper) GenerateKickTracksForUSC(usc uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateKickTracksForUSC:"), usc)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateShaderTrackForProgramType:
func (g GTMioTraceDataHelper) GenerateShaderTrackForProgramType(type_ uint16) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateShaderTrackForProgramType:"), type_)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateShaderTrackForProgramTypes
func (g GTMioTraceDataHelper) GenerateShaderTrackForProgramTypes() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateShaderTrackForProgramTypes"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateShaderTrackForProgramTypesForUSC:
func (g GTMioTraceDataHelper) GenerateShaderTrackForProgramTypesForUSC(usc uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateShaderTrackForProgramTypesForUSC:"), usc)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateShaderTracksForPipelineState:
func (g GTMioTraceDataHelper) GenerateShaderTracksForPipelineState(state uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateShaderTracksForPipelineState:"), state)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateTileTracksForUSC:
func (g GTMioTraceDataHelper) GenerateTileTracksForUSC(usc uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTileTracksForUSC:"), usc)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateTopBinaryTracks
func (g GTMioTraceDataHelper) GenerateTopBinaryTracks() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTopBinaryTracks"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateTopDrawTracks
func (g GTMioTraceDataHelper) GenerateTopDrawTracks() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTopDrawTracks"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateTopKickTracks
func (g GTMioTraceDataHelper) GenerateTopKickTracks() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTopKickTracks"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateTopRIATracks
func (g GTMioTraceDataHelper) GenerateTopRIATracks() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTopRIATracks"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/generateTrackForCliqueIndexes:count:group:
func (g GTMioTraceDataHelper) GenerateTrackForCliqueIndexesCountGroup(indexes unsafe.Pointer, count uint64, group VoidHandler) objectivec.IObject {
_block2, _ := NewVoidBlock(group)
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTrackForCliqueIndexes:count:group:"), indexes, count, _block2)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/stats
func (g GTMioTraceDataHelper) Stats() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stats"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/initWithTraceData:
func (g GTMioTraceDataHelper) InitWithTraceData(data objectivec.IObject) GTMioTraceDataHelper {
	rv := objc.Send[GTMioTraceDataHelper](g.ID, objc.Sel("initWithTraceData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/doNotMergeProgramTypes
func (g GTMioTraceDataHelper) DoNotMergeProgramTypes() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("doNotMergeProgramTypes"))
	return rv
}
func (g GTMioTraceDataHelper) SetDoNotMergeProgramTypes(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setDoNotMergeProgramTypes:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/showDriverInternalShaders
func (g GTMioTraceDataHelper) ShowDriverInternalShaders() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("showDriverInternalShaders"))
	return rv
}
func (g GTMioTraceDataHelper) SetShowDriverInternalShaders(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setShowDriverInternalShaders:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/showDriverIntersectionShaders
func (g GTMioTraceDataHelper) ShowDriverIntersectionShaders() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("showDriverIntersectionShaders"))
	return rv
}
func (g GTMioTraceDataHelper) SetShowDriverIntersectionShaders(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setShowDriverIntersectionShaders:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/showESLShaders
func (g GTMioTraceDataHelper) ShowESLShaders() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("showESLShaders"))
	return rv
}
func (g GTMioTraceDataHelper) SetShowESLShaders(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setShowESLShaders:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioTraceDataHelper/traceData
func (g GTMioTraceDataHelper) TraceData() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("traceData"))
	return objectivec.Object{ID: rv}
}

// GenerateTrackForCliqueIndexesCountGroupSync is a synchronous wrapper around [GTMioTraceDataHelper.GenerateTrackForCliqueIndexesCountGroup].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTMioTraceDataHelper) GenerateTrackForCliqueIndexesCountGroupSync(ctx context.Context, indexes unsafe.Pointer, count uint64) error {
	done := make(chan struct{}, 1)
	g.GenerateTrackForCliqueIndexesCountGroup(indexes, count, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

