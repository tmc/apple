// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerMCABinaryList] class.
var (
	_GTShaderProfilerMCABinaryListClass     GTShaderProfilerMCABinaryListClass
	_GTShaderProfilerMCABinaryListClassOnce sync.Once
)

func getGTShaderProfilerMCABinaryListClass() GTShaderProfilerMCABinaryListClass {
	_GTShaderProfilerMCABinaryListClassOnce.Do(func() {
		_GTShaderProfilerMCABinaryListClass = GTShaderProfilerMCABinaryListClass{class: objc.GetClass("GTShaderProfilerMCABinaryList")}
	})
	return _GTShaderProfilerMCABinaryListClass
}

// GetGTShaderProfilerMCABinaryListClass returns the class object for GTShaderProfilerMCABinaryList.
func GetGTShaderProfilerMCABinaryListClass() GTShaderProfilerMCABinaryListClass {
	return getGTShaderProfilerMCABinaryListClass()
}

type GTShaderProfilerMCABinaryListClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerMCABinaryListClass) Alloc() GTShaderProfilerMCABinaryList {
	rv := objc.Send[GTShaderProfilerMCABinaryList](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerMCABinaryList.AllocatedGPRCount]
//   - [GTShaderProfilerMCABinaryList.HighRegisterCount]
//   - [GTShaderProfilerMCABinaryList.McaBinaries]
//   - [GTShaderProfilerMCABinaryList.InitWithShaderProfilerResultPipelineStateIdProgramType]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList
type GTShaderProfilerMCABinaryList struct {
	objectivec.Object
}

// GTShaderProfilerMCABinaryListFromID constructs a [GTShaderProfilerMCABinaryList] from an objc.ID.
func GTShaderProfilerMCABinaryListFromID(id objc.ID) GTShaderProfilerMCABinaryList {
	return GTShaderProfilerMCABinaryList{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerMCABinaryList implements IGTShaderProfilerMCABinaryList.
var _ IGTShaderProfilerMCABinaryList = GTShaderProfilerMCABinaryList{}

// An interface definition for the [GTShaderProfilerMCABinaryList] class.
//
// # Methods
//
//   - [IGTShaderProfilerMCABinaryList.AllocatedGPRCount]
//   - [IGTShaderProfilerMCABinaryList.HighRegisterCount]
//   - [IGTShaderProfilerMCABinaryList.McaBinaries]
//   - [IGTShaderProfilerMCABinaryList.InitWithShaderProfilerResultPipelineStateIdProgramType]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList
type IGTShaderProfilerMCABinaryList interface {
	objectivec.IObject

	// Topic: Methods

	AllocatedGPRCount() int
	HighRegisterCount() int
	McaBinaries() foundation.INSArray
	InitWithShaderProfilerResultPipelineStateIdProgramType(result objectivec.IObject, id uint64, type_ uint32) GTShaderProfilerMCABinaryList
}

// Init initializes the instance.
func (g GTShaderProfilerMCABinaryList) Init() GTShaderProfilerMCABinaryList {
	rv := objc.Send[GTShaderProfilerMCABinaryList](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerMCABinaryList) Autorelease() GTShaderProfilerMCABinaryList {
	rv := objc.Send[GTShaderProfilerMCABinaryList](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerMCABinaryList creates a new GTShaderProfilerMCABinaryList instance.
func NewGTShaderProfilerMCABinaryList() GTShaderProfilerMCABinaryList {
	class := getGTShaderProfilerMCABinaryListClass()
	rv := objc.Send[GTShaderProfilerMCABinaryList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList/initWithShaderProfilerResult:pipelineStateId:programType:
func NewGTShaderProfilerMCABinaryListWithShaderProfilerResultPipelineStateIdProgramType(result objectivec.IObject, id uint64, type_ uint32) GTShaderProfilerMCABinaryList {
	instance := getGTShaderProfilerMCABinaryListClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShaderProfilerResult:pipelineStateId:programType:"), result, id, type_)
	return GTShaderProfilerMCABinaryListFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList/initWithShaderProfilerResult:pipelineStateId:programType:
func (g GTShaderProfilerMCABinaryList) InitWithShaderProfilerResultPipelineStateIdProgramType(result objectivec.IObject, id uint64, type_ uint32) GTShaderProfilerMCABinaryList {
	rv := objc.Send[GTShaderProfilerMCABinaryList](g.ID, objc.Sel("initWithShaderProfilerResult:pipelineStateId:programType:"), result, id, type_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList/allocatedGPRCount
func (g GTShaderProfilerMCABinaryList) AllocatedGPRCount() int {
	rv := objc.Send[int](g.ID, objc.Sel("allocatedGPRCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList/highRegisterCount
func (g GTShaderProfilerMCABinaryList) HighRegisterCount() int {
	rv := objc.Send[int](g.ID, objc.Sel("highRegisterCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerMCABinaryList/mcaBinaries
func (g GTShaderProfilerMCABinaryList) McaBinaries() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mcaBinaries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

