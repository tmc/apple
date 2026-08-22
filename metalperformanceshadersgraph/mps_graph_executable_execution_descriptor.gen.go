// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphExecutableExecutionDescriptor] class.
var (
	_MPSGraphExecutableExecutionDescriptorClass     MPSGraphExecutableExecutionDescriptorClass
	_MPSGraphExecutableExecutionDescriptorClassOnce sync.Once
)

func getMPSGraphExecutableExecutionDescriptorClass() MPSGraphExecutableExecutionDescriptorClass {
	_MPSGraphExecutableExecutionDescriptorClassOnce.Do(func() {
		_MPSGraphExecutableExecutionDescriptorClass = MPSGraphExecutableExecutionDescriptorClass{class: objc.GetClass("MPSGraphExecutableExecutionDescriptor")}
	})
	return _MPSGraphExecutableExecutionDescriptorClass
}

// GetMPSGraphExecutableExecutionDescriptorClass returns the class object for MPSGraphExecutableExecutionDescriptor.
func GetMPSGraphExecutableExecutionDescriptorClass() MPSGraphExecutableExecutionDescriptorClass {
	return getMPSGraphExecutableExecutionDescriptorClass()
}

type MPSGraphExecutableExecutionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphExecutableExecutionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphExecutableExecutionDescriptorClass) Alloc() MPSGraphExecutableExecutionDescriptor {
	rv := objc.Send[MPSGraphExecutableExecutionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that consists of all the levers to synchronize and schedule
// executable execution.
//
// # Instance Properties
//
//   - [MPSGraphExecutableExecutionDescriptor.CompletionHandler]: A notification that appears when graph-executable execution is finished.
//   - [MPSGraphExecutableExecutionDescriptor.SetCompletionHandler]
//   - [MPSGraphExecutableExecutionDescriptor.ScheduledHandler]: A notification that appears when graph-executable execution is scheduled.
//   - [MPSGraphExecutableExecutionDescriptor.SetScheduledHandler]
//   - [MPSGraphExecutableExecutionDescriptor.WaitUntilCompleted]: Flag for the graph executable to wait till the execution has completed.
//   - [MPSGraphExecutableExecutionDescriptor.SetWaitUntilCompleted]
//
// # Instance Methods
//
//   - [MPSGraphExecutableExecutionDescriptor.SignalEventAtExecutionEventValue]: Signals these shared events at execution stage and immediately proceeds.
//   - [MPSGraphExecutableExecutionDescriptor.WaitForEventValue]: Waits on these shared events before scheduling execution on the HW.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor
type MPSGraphExecutableExecutionDescriptor struct {
	MPSGraphObject
}

// MPSGraphExecutableExecutionDescriptorFromID constructs a [MPSGraphExecutableExecutionDescriptor] from an objc.ID.
//
// A class that consists of all the levers to synchronize and schedule
// executable execution.
func MPSGraphExecutableExecutionDescriptorFromID(id objc.ID) MPSGraphExecutableExecutionDescriptor {
	return MPSGraphExecutableExecutionDescriptor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphExecutableExecutionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphExecutableExecutionDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSGraphExecutableExecutionDescriptor.CompletionHandler]: A notification that appears when graph-executable execution is finished.
//   - [IMPSGraphExecutableExecutionDescriptor.SetCompletionHandler]
//   - [IMPSGraphExecutableExecutionDescriptor.ScheduledHandler]: A notification that appears when graph-executable execution is scheduled.
//   - [IMPSGraphExecutableExecutionDescriptor.SetScheduledHandler]
//   - [IMPSGraphExecutableExecutionDescriptor.WaitUntilCompleted]: Flag for the graph executable to wait till the execution has completed.
//   - [IMPSGraphExecutableExecutionDescriptor.SetWaitUntilCompleted]
//
// # Instance Methods
//
//   - [IMPSGraphExecutableExecutionDescriptor.SignalEventAtExecutionEventValue]: Signals these shared events at execution stage and immediately proceeds.
//   - [IMPSGraphExecutableExecutionDescriptor.WaitForEventValue]: Waits on these shared events before scheduling execution on the HW.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor
type IMPSGraphExecutableExecutionDescriptor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// A notification that appears when graph-executable execution is finished.
	CompletionHandler() MPSGraphTensorDataArrayErrorHandler
	SetCompletionHandler(value MPSGraphTensorDataArrayErrorHandler)
	// A notification that appears when graph-executable execution is scheduled.
	ScheduledHandler() MPSGraphTensorDataArrayErrorHandler
	SetScheduledHandler(value MPSGraphTensorDataArrayErrorHandler)
	// Flag for the graph executable to wait till the execution has completed.
	WaitUntilCompleted() bool
	SetWaitUntilCompleted(value bool)

	// Topic: Instance Methods

	// Signals these shared events at execution stage and immediately proceeds.
	SignalEventAtExecutionEventValue(event metal.MTLSharedEvent, executionStage MPSGraphExecutionStage, value uint64)
	// Waits on these shared events before scheduling execution on the HW.
	WaitForEventValue(event metal.MTLSharedEvent, value uint64)
}

// Init initializes the instance.
func (g MPSGraphExecutableExecutionDescriptor) Init() MPSGraphExecutableExecutionDescriptor {
	rv := objc.Send[MPSGraphExecutableExecutionDescriptor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphExecutableExecutionDescriptor) Autorelease() MPSGraphExecutableExecutionDescriptor {
	rv := objc.Send[MPSGraphExecutableExecutionDescriptor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphExecutableExecutionDescriptor creates a new MPSGraphExecutableExecutionDescriptor instance.
func NewMPSGraphExecutableExecutionDescriptor() MPSGraphExecutableExecutionDescriptor {
	class := getMPSGraphExecutableExecutionDescriptorClass()
	rv := objc.Send[MPSGraphExecutableExecutionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Signals these shared events at execution stage and immediately proceeds.
//
// event: Shared event to signal.
//
// executionStage: Execution stage to signal event at.
//
// value: Value for shared event to wait on.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/signal(_:atExecutionEvent:value:)
func (g MPSGraphExecutableExecutionDescriptor) SignalEventAtExecutionEventValue(event metal.MTLSharedEvent, executionStage MPSGraphExecutionStage, value uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("signalEvent:atExecutionEvent:value:"), event, executionStage, value)
}

// Waits on these shared events before scheduling execution on the HW.
//
// event: Shared event to wait on.
//
// value: Value for shared event to wait on.
//
// # Discussion
//
// This does not include encoding which can still continue.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/wait(for:value:)
func (g MPSGraphExecutableExecutionDescriptor) WaitForEventValue(event metal.MTLSharedEvent, value uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("waitForEvent:value:"), event, value)
}

// A notification that appears when graph-executable execution is finished.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/completionHandler
func (g MPSGraphExecutableExecutionDescriptor) CompletionHandler() MPSGraphTensorDataArrayErrorHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("completionHandler"))
	_ = rv
	return nil
}
func (g MPSGraphExecutableExecutionDescriptor) SetCompletionHandler(value MPSGraphTensorDataArrayErrorHandler) {
	block, cleanup := NewMPSGraphTensorDataArrayErrorBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setCompletionHandler:"), block)
}

// A notification that appears when graph-executable execution is scheduled.
//
// # Discussion
//
// Default value is nil.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/scheduledHandler
func (g MPSGraphExecutableExecutionDescriptor) ScheduledHandler() MPSGraphTensorDataArrayErrorHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scheduledHandler"))
	_ = rv
	return nil
}
func (g MPSGraphExecutableExecutionDescriptor) SetScheduledHandler(value MPSGraphTensorDataArrayErrorHandler) {
	block, cleanup := NewMPSGraphTensorDataArrayErrorBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setScheduledHandler:"), block)
}

// Flag for the graph executable to wait till the execution has completed.
//
// # Discussion
//
// Default value is false.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableExecutionDescriptor/waitUntilCompleted
func (g MPSGraphExecutableExecutionDescriptor) WaitUntilCompleted() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("waitUntilCompleted"))
	return rv
}
func (g MPSGraphExecutableExecutionDescriptor) SetWaitUntilCompleted(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setWaitUntilCompleted:"), value)
}
