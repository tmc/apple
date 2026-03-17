// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerSessionRequest] class.
var (
	_GTShaderProfilerSessionRequestClass     GTShaderProfilerSessionRequestClass
	_GTShaderProfilerSessionRequestClassOnce sync.Once
)

func getGTShaderProfilerSessionRequestClass() GTShaderProfilerSessionRequestClass {
	_GTShaderProfilerSessionRequestClassOnce.Do(func() {
		_GTShaderProfilerSessionRequestClass = GTShaderProfilerSessionRequestClass{class: objc.GetClass("GTShaderProfilerSessionRequest")}
	})
	return _GTShaderProfilerSessionRequestClass
}

// GetGTShaderProfilerSessionRequestClass returns the class object for GTShaderProfilerSessionRequest.
func GetGTShaderProfilerSessionRequestClass() GTShaderProfilerSessionRequestClass {
	return getGTShaderProfilerSessionRequestClass()
}

type GTShaderProfilerSessionRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerSessionRequestClass) Alloc() GTShaderProfilerSessionRequest {
	rv := objc.Send[GTShaderProfilerSessionRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerSessionRequest.ExecutionMode]
//   - [GTShaderProfilerSessionRequest.SetExecutionMode]
//   - [GTShaderProfilerSessionRequest.PerformanceState]
//   - [GTShaderProfilerSessionRequest.SetPerformanceState]
//   - [GTShaderProfilerSessionRequest.ProfilerMode]
//   - [GTShaderProfilerSessionRequest.SetProfilerMode]
//   - [GTShaderProfilerSessionRequest.StreamDataToLoad]
//   - [GTShaderProfilerSessionRequest.SetStreamDataToLoad]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerSessionRequest
type GTShaderProfilerSessionRequest struct {
	objectivec.Object
}

// GTShaderProfilerSessionRequestFromID constructs a [GTShaderProfilerSessionRequest] from an objc.ID.
func GTShaderProfilerSessionRequestFromID(id objc.ID) GTShaderProfilerSessionRequest {
	return GTShaderProfilerSessionRequest{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerSessionRequest implements IGTShaderProfilerSessionRequest.
var _ IGTShaderProfilerSessionRequest = GTShaderProfilerSessionRequest{}

// An interface definition for the [GTShaderProfilerSessionRequest] class.
//
// # Methods
//
//   - [IGTShaderProfilerSessionRequest.ExecutionMode]
//   - [IGTShaderProfilerSessionRequest.SetExecutionMode]
//   - [IGTShaderProfilerSessionRequest.PerformanceState]
//   - [IGTShaderProfilerSessionRequest.SetPerformanceState]
//   - [IGTShaderProfilerSessionRequest.ProfilerMode]
//   - [IGTShaderProfilerSessionRequest.SetProfilerMode]
//   - [IGTShaderProfilerSessionRequest.StreamDataToLoad]
//   - [IGTShaderProfilerSessionRequest.SetStreamDataToLoad]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerSessionRequest
type IGTShaderProfilerSessionRequest interface {
	objectivec.IObject

	// Topic: Methods

	ExecutionMode() uint32
	SetExecutionMode(value uint32)
	PerformanceState() uint32
	SetPerformanceState(value uint32)
	ProfilerMode() uint32
	SetProfilerMode(value uint32)
	StreamDataToLoad() foundation.INSURL
	SetStreamDataToLoad(value foundation.INSURL)
}

// Init initializes the instance.
func (g GTShaderProfilerSessionRequest) Init() GTShaderProfilerSessionRequest {
	rv := objc.Send[GTShaderProfilerSessionRequest](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerSessionRequest) Autorelease() GTShaderProfilerSessionRequest {
	rv := objc.Send[GTShaderProfilerSessionRequest](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerSessionRequest creates a new GTShaderProfilerSessionRequest instance.
func NewGTShaderProfilerSessionRequest() GTShaderProfilerSessionRequest {
	class := getGTShaderProfilerSessionRequestClass()
	rv := objc.Send[GTShaderProfilerSessionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerSessionRequest/executionMode
func (g GTShaderProfilerSessionRequest) ExecutionMode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("executionMode"))
	return rv
}
func (g GTShaderProfilerSessionRequest) SetExecutionMode(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setExecutionMode:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerSessionRequest/performanceState
func (g GTShaderProfilerSessionRequest) PerformanceState() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("performanceState"))
	return rv
}
func (g GTShaderProfilerSessionRequest) SetPerformanceState(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setPerformanceState:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerSessionRequest/profilerMode
func (g GTShaderProfilerSessionRequest) ProfilerMode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profilerMode"))
	return rv
}
func (g GTShaderProfilerSessionRequest) SetProfilerMode(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setProfilerMode:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerSessionRequest/streamDataToLoad
func (g GTShaderProfilerSessionRequest) StreamDataToLoad() foundation.INSURL {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("streamDataToLoad"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (g GTShaderProfilerSessionRequest) SetStreamDataToLoad(value foundation.INSURL) {
	objc.Send[struct{}](g.ID, objc.Sel("setStreamDataToLoad:"), value)
}

