// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTLLVMConnectionManager] class.
var (
	_GTLLVMConnectionManagerClass     GTLLVMConnectionManagerClass
	_GTLLVMConnectionManagerClassOnce sync.Once
)

func getGTLLVMConnectionManagerClass() GTLLVMConnectionManagerClass {
	_GTLLVMConnectionManagerClassOnce.Do(func() {
		_GTLLVMConnectionManagerClass = GTLLVMConnectionManagerClass{class: objc.GetClass("GTLLVMConnectionManager")}
	})
	return _GTLLVMConnectionManagerClass
}

// GetGTLLVMConnectionManagerClass returns the class object for GTLLVMConnectionManager.
func GetGTLLVMConnectionManagerClass() GTLLVMConnectionManagerClass {
	return getGTLLVMConnectionManagerClass()
}

type GTLLVMConnectionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTLLVMConnectionManagerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTLLVMConnectionManagerClass) Alloc() GTLLVMConnectionManager {
	rv := objc.Send[GTLLVMConnectionManager](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTLLVMConnectionManager._acquireAllHosts]
//   - [GTLLVMConnectionManager._acquireAvailableClient]
//   - [GTLLVMConnectionManager._acquireAvailableHost]
//   - [GTLLVMConnectionManager._anyHostBusy]
//   - [GTLLVMConnectionManager._queryVersion]
//   - [GTLLVMConnectionManager._releaseAllHosts]
//   - [GTLLVMConnectionManager._releaseAvailableClientAtIndex]
//   - [GTLLVMConnectionManager._releaseHost]
//   - [GTLLVMConnectionManager._tryAcquireHost]
//   - [GTLLVMConnectionManager.BinaryInfo]
//   - [GTLLVMConnectionManager.BinarySize]
//   - [GTLLVMConnectionManager.CreateLLMVAnalyzerForBinaryForKey]
//   - [GTLLVMConnectionManager.CreateLLMVAnalyzerForFilePath]
//   - [GTLLVMConnectionManager.DumpDebugInfoRanges]
//   - [GTLLVMConnectionManager.DumpFileInstructionOutput]
//   - [GTLLVMConnectionManager.EstablishConnectionWithLLVMHosts]
//   - [GTLLVMConnectionManager.GpuName]
//   - [GTLLVMConnectionManager.IsLLVMValid]
//   - [GTLLVMConnectionManager.NLLVMClients]
//   - [GTLLVMConnectionManager.ProcessInstructionTraceForBinariesWithNoTimestamp]
//   - [GTLLVMConnectionManager.ShaderProfilerBinaryInfo]
//   - [GTLLVMConnectionManager.TargetIndex]
//   - [GTLLVMConnectionManager.InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients]
//   - [GTLLVMConnectionManager.Version]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager
type GTLLVMConnectionManager struct {
	objectivec.Object
}

// GTLLVMConnectionManagerFromID constructs a [GTLLVMConnectionManager] from an objc.ID.
func GTLLVMConnectionManagerFromID(id objc.ID) GTLLVMConnectionManager {
	return GTLLVMConnectionManager{objectivec.Object{ID: id}}
}
// Ensure GTLLVMConnectionManager implements IGTLLVMConnectionManager.
var _ IGTLLVMConnectionManager = GTLLVMConnectionManager{}

// An interface definition for the [GTLLVMConnectionManager] class.
//
// # Methods
//
//   - [IGTLLVMConnectionManager._acquireAllHosts]
//   - [IGTLLVMConnectionManager._acquireAvailableClient]
//   - [IGTLLVMConnectionManager._acquireAvailableHost]
//   - [IGTLLVMConnectionManager._anyHostBusy]
//   - [IGTLLVMConnectionManager._queryVersion]
//   - [IGTLLVMConnectionManager._releaseAllHosts]
//   - [IGTLLVMConnectionManager._releaseAvailableClientAtIndex]
//   - [IGTLLVMConnectionManager._releaseHost]
//   - [IGTLLVMConnectionManager._tryAcquireHost]
//   - [IGTLLVMConnectionManager.BinaryInfo]
//   - [IGTLLVMConnectionManager.BinarySize]
//   - [IGTLLVMConnectionManager.CreateLLMVAnalyzerForBinaryForKey]
//   - [IGTLLVMConnectionManager.CreateLLMVAnalyzerForFilePath]
//   - [IGTLLVMConnectionManager.DumpDebugInfoRanges]
//   - [IGTLLVMConnectionManager.DumpFileInstructionOutput]
//   - [IGTLLVMConnectionManager.EstablishConnectionWithLLVMHosts]
//   - [IGTLLVMConnectionManager.GpuName]
//   - [IGTLLVMConnectionManager.IsLLVMValid]
//   - [IGTLLVMConnectionManager.NLLVMClients]
//   - [IGTLLVMConnectionManager.ProcessInstructionTraceForBinariesWithNoTimestamp]
//   - [IGTLLVMConnectionManager.ShaderProfilerBinaryInfo]
//   - [IGTLLVMConnectionManager.TargetIndex]
//   - [IGTLLVMConnectionManager.InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients]
//   - [IGTLLVMConnectionManager.Version]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager
type IGTLLVMConnectionManager interface {
	objectivec.IObject

	// Topic: Methods

	_acquireAllHosts()
	_acquireAvailableClient() uint64
	_acquireAvailableHost() uint32
	_anyHostBusy() bool
	_queryVersion() uint32
	_releaseAllHosts()
	_releaseAvailableClientAtIndex(index uint64)
	_releaseHost(host uint32)
	_tryAcquireHost(host uint32) bool
	BinaryInfo(info uint32) objectivec.IObject
	BinarySize(size uint32) uint32
	CreateLLMVAnalyzerForBinaryForKey(binary objectivec.IObject, key uint32) uint32
	CreateLLMVAnalyzerForFilePath(path objectivec.IObject) uint32
	DumpDebugInfoRanges(ranges uint32) objectivec.IObject
	DumpFileInstructionOutput(output uint32) objectivec.IObject
	EstablishConnectionWithLLVMHosts(lLVMHosts objectivec.IObject) bool
	GpuName() string
	IsLLVMValid(lLVMValid uint32) bool
	NLLVMClients() uint32
	ProcessInstructionTraceForBinariesWithNoTimestamp(binaries unsafe.Pointer, timestamp bool) objectivec.IObject
	ShaderProfilerBinaryInfo(info uint32) objectivec.IObject
	TargetIndex() int
	InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients(gPUName objectivec.IObject, index int, path objectivec.IObject, gen byte, name objectivec.IObject, clients uint32) GTLLVMConnectionManager
	Version() uint32
}

// Init initializes the instance.
func (g GTLLVMConnectionManager) Init() GTLLVMConnectionManager {
	rv := objc.Send[GTLLVMConnectionManager](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTLLVMConnectionManager) Autorelease() GTLLVMConnectionManager {
	rv := objc.Send[GTLLVMConnectionManager](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTLLVMConnectionManager creates a new GTLLVMConnectionManager instance.
func NewGTLLVMConnectionManager() GTLLVMConnectionManager {
	class := getGTLLVMConnectionManagerClass()
	rv := objc.Send[GTLLVMConnectionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/initWithGPUName:withTargetIndex:binaryPath:withGen:withSocketName:forNumClients:
func NewGTLLVMConnectionManagerWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients(gPUName objectivec.IObject, index int, path objectivec.IObject, gen byte, name objectivec.IObject, clients uint32) GTLLVMConnectionManager {
	instance := getGTLLVMConnectionManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGPUName:withTargetIndex:binaryPath:withGen:withSocketName:forNumClients:"), gPUName, index, path, gen, name, clients)
	return GTLLVMConnectionManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_acquireAllHosts
func (g GTLLVMConnectionManager) _acquireAllHosts() {
	objc.Send[objc.ID](g.ID, objc.Sel("_acquireAllHosts"))
}

// AcquireAllHosts is an exported wrapper for the private method _acquireAllHosts.
func (g GTLLVMConnectionManager) AcquireAllHosts() {
	g._acquireAllHosts()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_acquireAvailableClient
func (g GTLLVMConnectionManager) _acquireAvailableClient() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("_acquireAvailableClient"))
	return rv
}

// AcquireAvailableClient is an exported wrapper for the private method _acquireAvailableClient.
func (g GTLLVMConnectionManager) AcquireAvailableClient() uint64 {
	return g._acquireAvailableClient()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_acquireAvailableHost
func (g GTLLVMConnectionManager) _acquireAvailableHost() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("_acquireAvailableHost"))
	return rv
}

// AcquireAvailableHost is an exported wrapper for the private method _acquireAvailableHost.
func (g GTLLVMConnectionManager) AcquireAvailableHost() uint32 {
	return g._acquireAvailableHost()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_anyHostBusy
func (g GTLLVMConnectionManager) _anyHostBusy() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_anyHostBusy"))
	return rv
}

// AnyHostBusy is an exported wrapper for the private method _anyHostBusy.
func (g GTLLVMConnectionManager) AnyHostBusy() bool {
	return g._anyHostBusy()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_queryVersion
func (g GTLLVMConnectionManager) _queryVersion() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("_queryVersion"))
	return rv
}

// QueryVersion is an exported wrapper for the private method _queryVersion.
func (g GTLLVMConnectionManager) QueryVersion() uint32 {
	return g._queryVersion()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_releaseAllHosts
func (g GTLLVMConnectionManager) _releaseAllHosts() {
	objc.Send[objc.ID](g.ID, objc.Sel("_releaseAllHosts"))
}

// ReleaseAllHosts is an exported wrapper for the private method _releaseAllHosts.
func (g GTLLVMConnectionManager) ReleaseAllHosts() {
	g._releaseAllHosts()
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_releaseAvailableClientAtIndex:
func (g GTLLVMConnectionManager) _releaseAvailableClientAtIndex(index uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("_releaseAvailableClientAtIndex:"), index)
}

// ReleaseAvailableClientAtIndex is an exported wrapper for the private method _releaseAvailableClientAtIndex.
func (g GTLLVMConnectionManager) ReleaseAvailableClientAtIndex(index uint64) {
	g._releaseAvailableClientAtIndex(index)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_releaseHost:
func (g GTLLVMConnectionManager) _releaseHost(host uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("_releaseHost:"), host)
}

// ReleaseHost is an exported wrapper for the private method _releaseHost.
func (g GTLLVMConnectionManager) ReleaseHost(host uint32) {
	g._releaseHost(host)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/_tryAcquireHost:
func (g GTLLVMConnectionManager) _tryAcquireHost(host uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("_tryAcquireHost:"), host)
	return rv
}

// TryAcquireHost is an exported wrapper for the private method _tryAcquireHost.
func (g GTLLVMConnectionManager) TryAcquireHost(host uint32) bool {
	return g._tryAcquireHost(host)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/binaryInfo:
func (g GTLLVMConnectionManager) BinaryInfo(info uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binaryInfo:"), info)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/binarySize:
func (g GTLLVMConnectionManager) BinarySize(size uint32) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("binarySize:"), size)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/createLLMVAnalyzerForBinary:forKey:
func (g GTLLVMConnectionManager) CreateLLMVAnalyzerForBinaryForKey(binary objectivec.IObject, key uint32) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("createLLMVAnalyzerForBinary:forKey:"), binary, key)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/createLLMVAnalyzerForFilePath:
func (g GTLLVMConnectionManager) CreateLLMVAnalyzerForFilePath(path objectivec.IObject) uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("createLLMVAnalyzerForFilePath:"), path)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/dumpDebugInfoRanges:
func (g GTLLVMConnectionManager) DumpDebugInfoRanges(ranges uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dumpDebugInfoRanges:"), ranges)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/dumpFileInstructionOutput:
func (g GTLLVMConnectionManager) DumpFileInstructionOutput(output uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dumpFileInstructionOutput:"), output)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/establishConnectionWithLLVMHosts:
func (g GTLLVMConnectionManager) EstablishConnectionWithLLVMHosts(lLVMHosts objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("establishConnectionWithLLVMHosts:"), lLVMHosts)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/isLLVMValid:
func (g GTLLVMConnectionManager) IsLLVMValid(lLVMValid uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isLLVMValid:"), lLVMValid)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/processInstructionTraceForBinaries:withNoTimestamp:
func (g GTLLVMConnectionManager) ProcessInstructionTraceForBinariesWithNoTimestamp(binaries unsafe.Pointer, timestamp bool) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("processInstructionTraceForBinaries:withNoTimestamp:"), binaries, timestamp)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/shaderProfilerBinaryInfo:
func (g GTLLVMConnectionManager) ShaderProfilerBinaryInfo(info uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shaderProfilerBinaryInfo:"), info)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/initWithGPUName:withTargetIndex:binaryPath:withGen:withSocketName:forNumClients:
func (g GTLLVMConnectionManager) InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients(gPUName objectivec.IObject, index int, path objectivec.IObject, gen byte, name objectivec.IObject, clients uint32) GTLLVMConnectionManager {
	rv := objc.Send[GTLLVMConnectionManager](g.ID, objc.Sel("initWithGPUName:withTargetIndex:binaryPath:withGen:withSocketName:forNumClients:"), gPUName, index, path, gen, name, clients)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/gpuName
func (g GTLLVMConnectionManager) GpuName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/nLLVMClients
func (g GTLLVMConnectionManager) NLLVMClients() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("nLLVMClients"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/targetIndex
func (g GTLLVMConnectionManager) TargetIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("targetIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTLLVMConnectionManager/version
func (g GTLLVMConnectionManager) Version() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("version"))
	return rv
}

