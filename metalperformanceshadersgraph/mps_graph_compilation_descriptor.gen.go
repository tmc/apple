// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphCompilationDescriptor] class.
var (
	_MPSGraphCompilationDescriptorClass     MPSGraphCompilationDescriptorClass
	_MPSGraphCompilationDescriptorClassOnce sync.Once
)

func getMPSGraphCompilationDescriptorClass() MPSGraphCompilationDescriptorClass {
	_MPSGraphCompilationDescriptorClassOnce.Do(func() {
		_MPSGraphCompilationDescriptorClass = MPSGraphCompilationDescriptorClass{class: objc.GetClass("MPSGraphCompilationDescriptor")}
	})
	return _MPSGraphCompilationDescriptorClass
}

// GetMPSGraphCompilationDescriptorClass returns the class object for MPSGraphCompilationDescriptor.
func GetMPSGraphCompilationDescriptorClass() MPSGraphCompilationDescriptorClass {
	return getMPSGraphCompilationDescriptorClass()
}

type MPSGraphCompilationDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphCompilationDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphCompilationDescriptorClass) Alloc() MPSGraphCompilationDescriptor {
	rv := objc.Send[MPSGraphCompilationDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that consists of all the levers for compiling graphs.
//
// # Instance Properties
//
//   - [MPSGraphCompilationDescriptor.Callables]: The dictionary used during runtime to lookup the [MPSGraphExecutable](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable>) which correspond to the `symbolName`.
//   - [MPSGraphCompilationDescriptor.SetCallables]
//   - [MPSGraphCompilationDescriptor.CompilationCompletionHandler]: The handler that the graph calls when the compilation completes.
//   - [MPSGraphCompilationDescriptor.SetCompilationCompletionHandler]
//   - [MPSGraphCompilationDescriptor.DispatchQueue]: The dispatch queue used for the compilation.
//   - [MPSGraphCompilationDescriptor.SetDispatchQueue]
//   - [MPSGraphCompilationDescriptor.OptimizationLevel]: The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.
//   - [MPSGraphCompilationDescriptor.SetOptimizationLevel]
//   - [MPSGraphCompilationDescriptor.OptimizationProfile]: The optimization profile for the graph optimization.
//   - [MPSGraphCompilationDescriptor.SetOptimizationProfile]
//   - [MPSGraphCompilationDescriptor.ReducedPrecisionFastMath]: Across the executable allow reduced precision fast math optimizations.
//   - [MPSGraphCompilationDescriptor.SetReducedPrecisionFastMath]
//   - [MPSGraphCompilationDescriptor.WaitForCompilationCompletion]: Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.
//   - [MPSGraphCompilationDescriptor.SetWaitForCompilationCompletion]
//
// # Instance Methods
//
//   - [MPSGraphCompilationDescriptor.DisableTypeInference]: Turns off type inference and relies on type inference during runtime.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor
type MPSGraphCompilationDescriptor struct {
	MPSGraphObject
}

// MPSGraphCompilationDescriptorFromID constructs a [MPSGraphCompilationDescriptor] from an objc.ID.
//
// A class that consists of all the levers for compiling graphs.
func MPSGraphCompilationDescriptorFromID(id objc.ID) MPSGraphCompilationDescriptor {
	return MPSGraphCompilationDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphCompilationDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphCompilationDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphCompilationDescriptor.Callables]: The dictionary used during runtime to lookup the [MPSGraphExecutable](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable>) which correspond to the `symbolName`.
//   - [IMPSGraphCompilationDescriptor.SetCallables]
//   - [IMPSGraphCompilationDescriptor.CompilationCompletionHandler]: The handler that the graph calls when the compilation completes.
//   - [IMPSGraphCompilationDescriptor.SetCompilationCompletionHandler]
//   - [IMPSGraphCompilationDescriptor.DispatchQueue]: The dispatch queue used for the compilation.
//   - [IMPSGraphCompilationDescriptor.SetDispatchQueue]
//   - [IMPSGraphCompilationDescriptor.OptimizationLevel]: The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.
//   - [IMPSGraphCompilationDescriptor.SetOptimizationLevel]
//   - [IMPSGraphCompilationDescriptor.OptimizationProfile]: The optimization profile for the graph optimization.
//   - [IMPSGraphCompilationDescriptor.SetOptimizationProfile]
//   - [IMPSGraphCompilationDescriptor.ReducedPrecisionFastMath]: Across the executable allow reduced precision fast math optimizations.
//   - [IMPSGraphCompilationDescriptor.SetReducedPrecisionFastMath]
//   - [IMPSGraphCompilationDescriptor.WaitForCompilationCompletion]: Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.
//   - [IMPSGraphCompilationDescriptor.SetWaitForCompilationCompletion]
//
// # Instance Methods
//
//   - [IMPSGraphCompilationDescriptor.DisableTypeInference]: Turns off type inference and relies on type inference during runtime.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor
type IMPSGraphCompilationDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The dictionary used during runtime to lookup the [MPSGraphExecutable](<https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutable>) which correspond to the `symbolName`.
	Callables() MPSGraphCallableMap
	SetCallables(value MPSGraphCallableMap)
	// The handler that the graph calls when the compilation completes.
	CompilationCompletionHandler() MPSGraphExecutableErrorHandler
	SetCompilationCompletionHandler(value MPSGraphExecutableErrorHandler)
	// The dispatch queue used for the compilation.
	DispatchQueue() dispatch.Queue
	SetDispatchQueue(value dispatch.Queue)
	// The optimization level for the graph execution, default is MPSGraphOptimizationLevel1.
	OptimizationLevel() MPSGraphOptimization
	SetOptimizationLevel(value MPSGraphOptimization)
	// The optimization profile for the graph optimization.
	OptimizationProfile() MPSGraphOptimizationProfile
	SetOptimizationProfile(value MPSGraphOptimizationProfile)
	// Across the executable allow reduced precision fast math optimizations.
	ReducedPrecisionFastMath() MPSGraphReducedPrecisionFastMath
	SetReducedPrecisionFastMath(value MPSGraphReducedPrecisionFastMath)
	// Flag that makes the compile or specialize call blocking till the entire compilation is complete, defaults to NO.
	WaitForCompilationCompletion() bool
	SetWaitForCompilationCompletion(value bool)

	// Topic: Instance Methods

	// Turns off type inference and relies on type inference during runtime.
	DisableTypeInference()
}

// Init initializes the instance.
func (g MPSGraphCompilationDescriptor) Init() MPSGraphCompilationDescriptor {
	rv := objc.Send[MPSGraphCompilationDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphCompilationDescriptor) Autorelease() MPSGraphCompilationDescriptor {
	rv := objc.Send[MPSGraphCompilationDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphCompilationDescriptor creates a new MPSGraphCompilationDescriptor instance.
func NewMPSGraphCompilationDescriptor() MPSGraphCompilationDescriptor {
	class := getMPSGraphCompilationDescriptorClass()
	rv := objc.Send[MPSGraphCompilationDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Turns off type inference and relies on type inference during runtime.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/disableTypeInference()
func (g MPSGraphCompilationDescriptor) DisableTypeInference() {
	objc.Send[objc.ID](g.ID, objc.Sel("disableTypeInference"))
}

// The dictionary used during runtime to lookup the [MPSGraphExecutable] which
// correspond to the `symbolName`.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/callables
func (g MPSGraphCompilationDescriptor) Callables() MPSGraphCallableMap {
	rv := objc.Send[MPSGraphCallableMap](g.ID, objc.Sel("callables"))
	return MPSGraphCallableMap(rv)
}
func (g MPSGraphCompilationDescriptor) SetCallables(value MPSGraphCallableMap) {
	objc.Send[struct{}](g.ID, objc.Sel("setCallables:"), value)
}

// The handler that the graph calls when the compilation completes.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/compilationCompletionHandler
func (g MPSGraphCompilationDescriptor) CompilationCompletionHandler() MPSGraphExecutableErrorHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("compilationCompletionHandler"))
	_ = rv
	return nil
}
func (g MPSGraphCompilationDescriptor) SetCompilationCompletionHandler(value MPSGraphExecutableErrorHandler) {
	block, cleanup := NewMPSGraphExecutableErrorBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setCompilationCompletionHandler:"), block)
}

// The dispatch queue used for the compilation.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/dispatchQueue
func (g MPSGraphCompilationDescriptor) DispatchQueue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("dispatchQueue"))
	return dispatch.QueueFromHandle(rv)
}
func (g MPSGraphCompilationDescriptor) SetDispatchQueue(value dispatch.Queue) {
	objc.Send[struct{}](g.ID, objc.Sel("setDispatchQueue:"), uintptr(value.Handle()))
}

// The optimization level for the graph execution, default is
// MPSGraphOptimizationLevel1.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationLevel
func (g MPSGraphCompilationDescriptor) OptimizationLevel() MPSGraphOptimization {
	rv := objc.Send[MPSGraphOptimization](g.ID, objc.Sel("optimizationLevel"))
	return MPSGraphOptimization(rv)
}
func (g MPSGraphCompilationDescriptor) SetOptimizationLevel(value MPSGraphOptimization) {
	objc.Send[struct{}](g.ID, objc.Sel("setOptimizationLevel:"), value)
}

// The optimization profile for the graph optimization.
//
// # Discussion
//
// Default is MPSGraphOptimizationProfilePerformance.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/optimizationProfile
func (g MPSGraphCompilationDescriptor) OptimizationProfile() MPSGraphOptimizationProfile {
	rv := objc.Send[MPSGraphOptimizationProfile](g.ID, objc.Sel("optimizationProfile"))
	return MPSGraphOptimizationProfile(rv)
}
func (g MPSGraphCompilationDescriptor) SetOptimizationProfile(value MPSGraphOptimizationProfile) {
	objc.Send[struct{}](g.ID, objc.Sel("setOptimizationProfile:"), value)
}

// Across the executable allow reduced precision fast math optimizations.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/reducedPrecisionFastMath
func (g MPSGraphCompilationDescriptor) ReducedPrecisionFastMath() MPSGraphReducedPrecisionFastMath {
	rv := objc.Send[MPSGraphReducedPrecisionFastMath](g.ID, objc.Sel("reducedPrecisionFastMath"))
	return MPSGraphReducedPrecisionFastMath(rv)
}
func (g MPSGraphCompilationDescriptor) SetReducedPrecisionFastMath(value MPSGraphReducedPrecisionFastMath) {
	objc.Send[struct{}](g.ID, objc.Sel("setReducedPrecisionFastMath:"), value)
}

// Flag that makes the compile or specialize call blocking till the entire
// compilation is complete, defaults to NO.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphCompilationDescriptor/waitForCompilationCompletion
func (g MPSGraphCompilationDescriptor) WaitForCompilationCompletion() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("waitForCompilationCompletion"))
	return rv
}
func (g MPSGraphCompilationDescriptor) SetWaitForCompilationCompletion(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setWaitForCompilationCompletion:"), value)
}
