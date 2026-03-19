// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerTimingInfo] class.
var (
	_GTShaderProfilerTimingInfoClass     GTShaderProfilerTimingInfoClass
	_GTShaderProfilerTimingInfoClassOnce sync.Once
)

func getGTShaderProfilerTimingInfoClass() GTShaderProfilerTimingInfoClass {
	_GTShaderProfilerTimingInfoClassOnce.Do(func() {
		_GTShaderProfilerTimingInfoClass = GTShaderProfilerTimingInfoClass{class: objc.GetClass("GTShaderProfilerTimingInfo")}
	})
	return _GTShaderProfilerTimingInfoClass
}

// GetGTShaderProfilerTimingInfoClass returns the class object for GTShaderProfilerTimingInfo.
func GetGTShaderProfilerTimingInfoClass() GTShaderProfilerTimingInfoClass {
	return getGTShaderProfilerTimingInfoClass()
}

type GTShaderProfilerTimingInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerTimingInfoClass) Alloc() GTShaderProfilerTimingInfo {
	rv := objc.Send[GTShaderProfilerTimingInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerTimingInfo.ComputeTime]
//   - [GTShaderProfilerTimingInfo.EncodeWithCoder]
//   - [GTShaderProfilerTimingInfo.FragmentTime]
//   - [GTShaderProfilerTimingInfo.Time]
//   - [GTShaderProfilerTimingInfo.VertexTime]
//   - [GTShaderProfilerTimingInfo.InitWithCoder]
//   - [GTShaderProfilerTimingInfo.InitWithTimeVertexTimeFragmentTimeComputeTime]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo
type GTShaderProfilerTimingInfo struct {
	objectivec.Object
}

// GTShaderProfilerTimingInfoFromID constructs a [GTShaderProfilerTimingInfo] from an objc.ID.
func GTShaderProfilerTimingInfoFromID(id objc.ID) GTShaderProfilerTimingInfo {
	return GTShaderProfilerTimingInfo{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerTimingInfo implements IGTShaderProfilerTimingInfo.
var _ IGTShaderProfilerTimingInfo = GTShaderProfilerTimingInfo{}

// An interface definition for the [GTShaderProfilerTimingInfo] class.
//
// # Methods
//
//   - [IGTShaderProfilerTimingInfo.ComputeTime]
//   - [IGTShaderProfilerTimingInfo.EncodeWithCoder]
//   - [IGTShaderProfilerTimingInfo.FragmentTime]
//   - [IGTShaderProfilerTimingInfo.Time]
//   - [IGTShaderProfilerTimingInfo.VertexTime]
//   - [IGTShaderProfilerTimingInfo.InitWithCoder]
//   - [IGTShaderProfilerTimingInfo.InitWithTimeVertexTimeFragmentTimeComputeTime]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo
type IGTShaderProfilerTimingInfo interface {
	objectivec.IObject

	// Topic: Methods

	ComputeTime() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	FragmentTime() uint64
	Time() uint64
	VertexTime() uint64
	InitWithCoder(coder foundation.INSCoder) GTShaderProfilerTimingInfo
	InitWithTimeVertexTimeFragmentTimeComputeTime(time uint64, time2 uint64, time3 uint64, time4 uint64) GTShaderProfilerTimingInfo
}

// Init initializes the instance.
func (g GTShaderProfilerTimingInfo) Init() GTShaderProfilerTimingInfo {
	rv := objc.Send[GTShaderProfilerTimingInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerTimingInfo) Autorelease() GTShaderProfilerTimingInfo {
	rv := objc.Send[GTShaderProfilerTimingInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerTimingInfo creates a new GTShaderProfilerTimingInfo instance.
func NewGTShaderProfilerTimingInfo() GTShaderProfilerTimingInfo {
	class := getGTShaderProfilerTimingInfoClass()
	rv := objc.Send[GTShaderProfilerTimingInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/initWithCoder:
func NewGTShaderProfilerTimingInfoWithCoder(coder objectivec.IObject) GTShaderProfilerTimingInfo {
	instance := getGTShaderProfilerTimingInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerTimingInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/initWithTime:vertexTime:fragmentTime:computeTime:
func NewGTShaderProfilerTimingInfoWithTimeVertexTimeFragmentTimeComputeTime(time uint64, time2 uint64, time3 uint64, time4 uint64) GTShaderProfilerTimingInfo {
	instance := getGTShaderProfilerTimingInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTime:vertexTime:fragmentTime:computeTime:"), time, time2, time3, time4)
	return GTShaderProfilerTimingInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/encodeWithCoder:
func (g GTShaderProfilerTimingInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/initWithCoder:
func (g GTShaderProfilerTimingInfo) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerTimingInfo {
	rv := objc.Send[GTShaderProfilerTimingInfo](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/initWithTime:vertexTime:fragmentTime:computeTime:
func (g GTShaderProfilerTimingInfo) InitWithTimeVertexTimeFragmentTimeComputeTime(time uint64, time2 uint64, time3 uint64, time4 uint64) GTShaderProfilerTimingInfo {
	rv := objc.Send[GTShaderProfilerTimingInfo](g.ID, objc.Sel("initWithTime:vertexTime:fragmentTime:computeTime:"), time, time2, time3, time4)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/supportsSecureCoding
func (_GTShaderProfilerTimingInfoClass GTShaderProfilerTimingInfoClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTShaderProfilerTimingInfoClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/computeTime
func (g GTShaderProfilerTimingInfo) ComputeTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("computeTime"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/fragmentTime
func (g GTShaderProfilerTimingInfo) FragmentTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("fragmentTime"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/time
func (g GTShaderProfilerTimingInfo) Time() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("time"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerTimingInfo/vertexTime
func (g GTShaderProfilerTimingInfo) VertexTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("vertexTime"))
	return rv
}

