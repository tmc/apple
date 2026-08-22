// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphExecutionDescriptor] class.
var (
	_MPSGraphExecutionDescriptorClass     MPSGraphExecutionDescriptorClass
	_MPSGraphExecutionDescriptorClassOnce sync.Once
)

func getMPSGraphExecutionDescriptorClass() MPSGraphExecutionDescriptorClass {
	_MPSGraphExecutionDescriptorClassOnce.Do(func() {
		_MPSGraphExecutionDescriptorClass = MPSGraphExecutionDescriptorClass{class: objc.GetClass("MPSGraphExecutionDescriptor")}
	})
	return _MPSGraphExecutionDescriptorClass
}

// GetMPSGraphExecutionDescriptorClass returns the class object for MPSGraphExecutionDescriptor.
func GetMPSGraphExecutionDescriptorClass() MPSGraphExecutionDescriptorClass {
	return getMPSGraphExecutionDescriptorClass()
}

type MPSGraphExecutionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphExecutionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphExecutionDescriptorClass) Alloc() MPSGraphExecutionDescriptor {
	rv := objc.Send[MPSGraphExecutionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that consists of all the levers to synchronize and schedule graph
// execution.
//
// # Instance Properties
//
//   - [MPSGraphExecutionDescriptor.CompilationDescriptor]: The compilation descriptor for the graph.
//   - [MPSGraphExecutionDescriptor.SetCompilationDescriptor]
//   - [MPSGraphExecutionDescriptor.CompletionHandler]: The handler that graph calls at the completion of the execution.
//   - [MPSGraphExecutionDescriptor.SetCompletionHandler]
//   - [MPSGraphExecutionDescriptor.ScheduledHandler]: The handler that graph calls when it schedules the execution.
//   - [MPSGraphExecutionDescriptor.SetScheduledHandler]
//   - [MPSGraphExecutionDescriptor.WaitUntilCompleted]: The flag that blocks the execution call until the entire execution is complete.
//   - [MPSGraphExecutionDescriptor.SetWaitUntilCompleted]
//
// # Instance Methods
//
//   - [MPSGraphExecutionDescriptor.SignalEventAtExecutionEventValue]: Executable signals these shared events at execution stage and immediately proceeds.
//   - [MPSGraphExecutionDescriptor.WaitForEventValue]: Executable waits on these shared events before scheduling execution on the HW, this does not include encoding which can still continue.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor
type MPSGraphExecutionDescriptor struct {
	MPSGraphObject
}

// MPSGraphExecutionDescriptorFromID constructs a [MPSGraphExecutionDescriptor] from an objc.ID.
//
// A class that consists of all the levers to synchronize and schedule graph
// execution.
func MPSGraphExecutionDescriptorFromID(id objc.ID) MPSGraphExecutionDescriptor {
	return MPSGraphExecutionDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphExecutionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphExecutionDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphExecutionDescriptor.CompilationDescriptor]: The compilation descriptor for the graph.
//   - [IMPSGraphExecutionDescriptor.SetCompilationDescriptor]
//   - [IMPSGraphExecutionDescriptor.CompletionHandler]: The handler that graph calls at the completion of the execution.
//   - [IMPSGraphExecutionDescriptor.SetCompletionHandler]
//   - [IMPSGraphExecutionDescriptor.ScheduledHandler]: The handler that graph calls when it schedules the execution.
//   - [IMPSGraphExecutionDescriptor.SetScheduledHandler]
//   - [IMPSGraphExecutionDescriptor.WaitUntilCompleted]: The flag that blocks the execution call until the entire execution is complete.
//   - [IMPSGraphExecutionDescriptor.SetWaitUntilCompleted]
//
// # Instance Methods
//
//   - [IMPSGraphExecutionDescriptor.SignalEventAtExecutionEventValue]: Executable signals these shared events at execution stage and immediately proceeds.
//   - [IMPSGraphExecutionDescriptor.WaitForEventValue]: Executable waits on these shared events before scheduling execution on the HW, this does not include encoding which can still continue.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor
type IMPSGraphExecutionDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The compilation descriptor for the graph.
	CompilationDescriptor() IMPSGraphCompilationDescriptor
	SetCompilationDescriptor(value IMPSGraphCompilationDescriptor)
	// The handler that graph calls at the completion of the execution.
	CompletionHandler() MPSGraphCompletionHandler
	SetCompletionHandler(value objc.ID)
	// The handler that graph calls when it schedules the execution.
	ScheduledHandler() MPSGraphScheduledHandler
	SetScheduledHandler(value objc.ID)
	// The flag that blocks the execution call until the entire execution is complete.
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)

	// Topic: Instance Methods

	// Executable signals these shared events at execution stage and immediately proceeds.
	SignalEventAtExecutionEventValue(event metal.MTLSharedEvent, executionStage MPSGraphExecutionStage, value uint64)
	// Executable waits on these shared events before scheduling execution on the HW, this does not include encoding which can still continue.
	WaitForEventValue(event metal.MTLSharedEvent, value uint64)
}

// Init initializes the instance.
func (g MPSGraphExecutionDescriptor) Init() MPSGraphExecutionDescriptor {
	rv := objc.Send[MPSGraphExecutionDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphExecutionDescriptor) Autorelease() MPSGraphExecutionDescriptor {
	rv := objc.Send[MPSGraphExecutionDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphExecutionDescriptor creates a new MPSGraphExecutionDescriptor instance.
func NewMPSGraphExecutionDescriptor() MPSGraphExecutionDescriptor {
	class := getMPSGraphExecutionDescriptorClass()
	rv := objc.Send[MPSGraphExecutionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Executable signals these shared events at execution stage and immediately
// proceeds.
//
// event: Shared event to signal.
//
// executionStage: Execution stage to signal event at.
//
// value: Value for shared event to wait on.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/signal(_:atExecutionEvent:value:)
func (g MPSGraphExecutionDescriptor) SignalEventAtExecutionEventValue(event metal.MTLSharedEvent, executionStage MPSGraphExecutionStage, value uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("signalEvent:atExecutionEvent:value:"), event, executionStage, value)
}

// Executable waits on these shared events before scheduling execution on the
// HW, this does not include encoding which can still continue.
//
// event: Shared event graph waits on.
//
// value: Value of shared event graph waits on.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/wait(for:value:)
func (g MPSGraphExecutionDescriptor) WaitForEventValue(event metal.MTLSharedEvent, value uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("waitForEvent:value:"), event, value)
}

// The compilation descriptor for the graph.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/compilationDescriptor
func (g MPSGraphExecutionDescriptor) CompilationDescriptor() IMPSGraphCompilationDescriptor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("compilationDescriptor"))
	return MPSGraphCompilationDescriptorFromID(objc.ID(rv))
}
func (g MPSGraphExecutionDescriptor) SetCompilationDescriptor(value IMPSGraphCompilationDescriptor) {
	objc.Send[struct{}](g.ID, objc.Sel("setCompilationDescriptor:"), value)
}

// The handler that graph calls at the completion of the execution.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/completionHandler
func (g MPSGraphExecutionDescriptor) CompletionHandler() MPSGraphCompletionHandler {
	rv := objc.Send[MPSGraphCompletionHandler](g.ID, objc.Sel("completionHandler"))
	return MPSGraphCompletionHandler(rv)
}
func (g MPSGraphExecutionDescriptor) SetCompletionHandler(value objc.ID) {
	objc.Send[struct{}](g.ID, objc.Sel("setCompletionHandler:"), value)
}

// The handler that graph calls when it schedules the execution.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/scheduledHandler
func (g MPSGraphExecutionDescriptor) ScheduledHandler() MPSGraphScheduledHandler {
	rv := objc.Send[MPSGraphScheduledHandler](g.ID, objc.Sel("scheduledHandler"))
	return MPSGraphScheduledHandler(rv)
}
func (g MPSGraphExecutionDescriptor) SetScheduledHandler(value objc.ID) {
	objc.Send[struct{}](g.ID, objc.Sel("setScheduledHandler:"), value)
}

// The flag that blocks the execution call until the entire execution is
// complete.
//
// # Discussion
//
// Defaults to NO.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutionDescriptor/waitUntilCompleted
func (g MPSGraphExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("waitUntilCompleted"))
	return rv
}
func (g MPSGraphExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setWaitUntilCompleted:"), value)
}
