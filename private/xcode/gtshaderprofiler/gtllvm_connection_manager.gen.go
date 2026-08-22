// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[GTLLVMConnectionManager](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

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
//   - [GTLLVMConnectionManager.ShaderProfilerBinaryInfo]
//   - [GTLLVMConnectionManager.TargetIndex]
//   - [GTLLVMConnectionManager.InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients]
//   - [GTLLVMConnectionManager.Version]
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
//   - [IGTLLVMConnectionManager.ShaderProfilerBinaryInfo]
//   - [IGTLLVMConnectionManager.TargetIndex]
//   - [IGTLLVMConnectionManager.InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients]
//   - [IGTLLVMConnectionManager.Version]
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
	BinaryInfo(info uint32) GTAPSBinaryInfo
	BinarySize(size uint32) uint32
	CreateLLMVAnalyzerForBinaryForKey(binary objectivec.IObject, key uint32) uint32
	CreateLLMVAnalyzerForFilePath(path objectivec.IObject) uint32
	DumpDebugInfoRanges(ranges uint32) objectivec.IObject
	DumpFileInstructionOutput(output uint32) objectivec.IObject
	EstablishConnectionWithLLVMHosts(lLVMHosts objectivec.IObject) bool
	GpuName() string
	IsLLVMValid(lLVMValid uint32) bool
	NLLVMClients() uint32
	ShaderProfilerBinaryInfo(info uint32) GTShaderProfilerBinaryInfo
	TargetIndex() int
	InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients(gPUName objectivec.IObject, index int, path objectivec.IObject, gen byte, name objectivec.IObject, clients uint32) GTLLVMConnectionManager
	Version() uint32
}

// Init initializes the instance.
func (g GTLLVMConnectionManager) Init() GTLLVMConnectionManager {
	rv := objc.SendIfResponds[GTLLVMConnectionManager](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTLLVMConnectionManager) Autorelease() GTLLVMConnectionManager {
	rv := objc.SendIfResponds[GTLLVMConnectionManager](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTLLVMConnectionManager creates a new GTLLVMConnectionManager instance.
func NewGTLLVMConnectionManager() GTLLVMConnectionManager {
	class := getGTLLVMConnectionManagerClass()
	rv := objc.SendIfResponds[GTLLVMConnectionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTLLVMConnectionManagerWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients(gPUName objectivec.IObject, index int, path objectivec.IObject, gen byte, name objectivec.IObject, clients uint32) GTLLVMConnectionManager {
	instance := getGTLLVMConnectionManagerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithGPUName:withTargetIndex:binaryPath:withGen:withSocketName:forNumClients:"), gPUName, index, path, gen, name, clients)
	return GTLLVMConnectionManagerFromID(rv)
}

func (g GTLLVMConnectionManager) _acquireAllHosts() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_acquireAllHosts"))
}

// AcquireAllHosts is an exported wrapper for the private method _acquireAllHosts.
func (g GTLLVMConnectionManager) AcquireAllHosts() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_acquireAllHosts")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_acquireAllHosts"}
		return err
	}
	g._acquireAllHosts()
	return nil
}

// CanAcquireAllHosts reports whether the receiver responds to the private selector _acquireAllHosts.
func (g GTLLVMConnectionManager) CanAcquireAllHosts() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_acquireAllHosts"))
}
func (g GTLLVMConnectionManager) _acquireAvailableClient() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("_acquireAvailableClient"))
	return rv
}

// AcquireAvailableClient is an exported wrapper for the private method _acquireAvailableClient.
func (g GTLLVMConnectionManager) AcquireAvailableClient() (uint64, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_acquireAvailableClient")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_acquireAvailableClient"}
		return 0, err
	}
	return g._acquireAvailableClient(), nil
}

// CanAcquireAvailableClient reports whether the receiver responds to the private selector _acquireAvailableClient.
func (g GTLLVMConnectionManager) CanAcquireAvailableClient() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_acquireAvailableClient"))
}
func (g GTLLVMConnectionManager) _acquireAvailableHost() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("_acquireAvailableHost"))
	return rv
}

// AcquireAvailableHost is an exported wrapper for the private method _acquireAvailableHost.
func (g GTLLVMConnectionManager) AcquireAvailableHost() (uint32, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_acquireAvailableHost")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_acquireAvailableHost"}
		return 0, err
	}
	return g._acquireAvailableHost(), nil
}

// CanAcquireAvailableHost reports whether the receiver responds to the private selector _acquireAvailableHost.
func (g GTLLVMConnectionManager) CanAcquireAvailableHost() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_acquireAvailableHost"))
}
func (g GTLLVMConnectionManager) _anyHostBusy() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("_anyHostBusy"))
	return rv
}

// AnyHostBusy is an exported wrapper for the private method _anyHostBusy.
func (g GTLLVMConnectionManager) AnyHostBusy() (bool, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_anyHostBusy")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_anyHostBusy"}
		return false, err
	}
	return g._anyHostBusy(), nil
}

// CanAnyHostBusy reports whether the receiver responds to the private selector _anyHostBusy.
func (g GTLLVMConnectionManager) CanAnyHostBusy() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_anyHostBusy"))
}
func (g GTLLVMConnectionManager) _queryVersion() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("_queryVersion"))
	return rv
}

// QueryVersion is an exported wrapper for the private method _queryVersion.
func (g GTLLVMConnectionManager) QueryVersion() (uint32, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_queryVersion")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_queryVersion"}
		return 0, err
	}
	return g._queryVersion(), nil
}

// CanQueryVersion reports whether the receiver responds to the private selector _queryVersion.
func (g GTLLVMConnectionManager) CanQueryVersion() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_queryVersion"))
}
func (g GTLLVMConnectionManager) _releaseAllHosts() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_releaseAllHosts"))
}

// ReleaseAllHosts is an exported wrapper for the private method _releaseAllHosts.
func (g GTLLVMConnectionManager) ReleaseAllHosts() error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_releaseAllHosts")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_releaseAllHosts"}
		return err
	}
	g._releaseAllHosts()
	return nil
}

// CanReleaseAllHosts reports whether the receiver responds to the private selector _releaseAllHosts.
func (g GTLLVMConnectionManager) CanReleaseAllHosts() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_releaseAllHosts"))
}
func (g GTLLVMConnectionManager) _releaseAvailableClientAtIndex(index uint64) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_releaseAvailableClientAtIndex:"), index)
}

// ReleaseAvailableClientAtIndex is an exported wrapper for the private method _releaseAvailableClientAtIndex.
func (g GTLLVMConnectionManager) ReleaseAvailableClientAtIndex(index uint64) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_releaseAvailableClientAtIndex:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_releaseAvailableClientAtIndex:"}
		return err
	}
	g._releaseAvailableClientAtIndex(index)
	return nil
}

// CanReleaseAvailableClientAtIndex reports whether the receiver responds to the private selector _releaseAvailableClientAtIndex:.
func (g GTLLVMConnectionManager) CanReleaseAvailableClientAtIndex() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_releaseAvailableClientAtIndex:"))
}
func (g GTLLVMConnectionManager) _releaseHost(host uint32) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("_releaseHost:"), host)
}

// ReleaseHost is an exported wrapper for the private method _releaseHost.
func (g GTLLVMConnectionManager) ReleaseHost(host uint32) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_releaseHost:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_releaseHost:"}
		return err
	}
	g._releaseHost(host)
	return nil
}

// CanReleaseHost reports whether the receiver responds to the private selector _releaseHost:.
func (g GTLLVMConnectionManager) CanReleaseHost() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_releaseHost:"))
}
func (g GTLLVMConnectionManager) _tryAcquireHost(host uint32) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("_tryAcquireHost:"), host)
	return rv
}

// TryAcquireHost is an exported wrapper for the private method _tryAcquireHost.
func (g GTLLVMConnectionManager) TryAcquireHost(host uint32) (bool, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_tryAcquireHost:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_tryAcquireHost:"}
		return false, err
	}
	return g._tryAcquireHost(host), nil
}

// CanTryAcquireHost reports whether the receiver responds to the private selector _tryAcquireHost:.
func (g GTLLVMConnectionManager) CanTryAcquireHost() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_tryAcquireHost:"))
}
func (g GTLLVMConnectionManager) BinaryInfo(info uint32) GTAPSBinaryInfo {
	rv := objc.SendIfResponds[GTAPSBinaryInfo](g.ID, objc.Sel("binaryInfo:"), info)
	return GTAPSBinaryInfo(rv)
}
func (g GTLLVMConnectionManager) BinarySize(size uint32) uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("binarySize:"), size)
	return rv
}
func (g GTLLVMConnectionManager) CreateLLMVAnalyzerForBinaryForKey(binary objectivec.IObject, key uint32) uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("createLLMVAnalyzerForBinary:forKey:"), binary, key)
	return rv
}
func (g GTLLVMConnectionManager) CreateLLMVAnalyzerForFilePath(path objectivec.IObject) uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("createLLMVAnalyzerForFilePath:"), path)
	return rv
}
func (g GTLLVMConnectionManager) DumpDebugInfoRanges(ranges uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("dumpDebugInfoRanges:"), ranges)
	return objectivec.Object{ID: rv}
}
func (g GTLLVMConnectionManager) DumpFileInstructionOutput(output uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("dumpFileInstructionOutput:"), output)
	return objectivec.Object{ID: rv}
}
func (g GTLLVMConnectionManager) EstablishConnectionWithLLVMHosts(lLVMHosts objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("establishConnectionWithLLVMHosts:"), lLVMHosts)
	return rv
}
func (g GTLLVMConnectionManager) IsLLVMValid(lLVMValid uint32) bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("isLLVMValid:"), lLVMValid)
	return rv
}
func (g GTLLVMConnectionManager) ShaderProfilerBinaryInfo(info uint32) GTShaderProfilerBinaryInfo {
	rv := objc.SendIfResponds[GTShaderProfilerBinaryInfo](g.ID, objc.Sel("shaderProfilerBinaryInfo:"), info)
	return GTShaderProfilerBinaryInfo(rv)
}
func (g GTLLVMConnectionManager) InitWithGPUNameWithTargetIndexBinaryPathWithGenWithSocketNameForNumClients(gPUName objectivec.IObject, index int, path objectivec.IObject, gen byte, name objectivec.IObject, clients uint32) GTLLVMConnectionManager {
	rv := objc.SendIfResponds[GTLLVMConnectionManager](g.ID, objc.Sel("initWithGPUName:withTargetIndex:binaryPath:withGen:withSocketName:forNumClients:"), gPUName, index, path, gen, name, clients)
	return rv
}

func (g GTLLVMConnectionManager) GpuName() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("gpuName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTLLVMConnectionManager) NLLVMClients() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("nLLVMClients"))
	return rv
}
func (g GTLLVMConnectionManager) TargetIndex() int {
	rv := objc.SendIfResponds[int](g.ID, objc.Sel("targetIndex"))
	return rv
}
func (g GTLLVMConnectionManager) Version() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("version"))
	return rv
}
