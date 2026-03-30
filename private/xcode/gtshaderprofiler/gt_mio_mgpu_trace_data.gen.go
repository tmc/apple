// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioMGPUTraceData] class.
var (
	_GTMioMGPUTraceDataClass     GTMioMGPUTraceDataClass
	_GTMioMGPUTraceDataClassOnce sync.Once
)

func getGTMioMGPUTraceDataClass() GTMioMGPUTraceDataClass {
	_GTMioMGPUTraceDataClassOnce.Do(func() {
		_GTMioMGPUTraceDataClass = GTMioMGPUTraceDataClass{class: objc.GetClass("GTMioMGPUTraceData")}
	})
	return _GTMioMGPUTraceDataClass
}

// GetGTMioMGPUTraceDataClass returns the class object for GTMioMGPUTraceData.
func GetGTMioMGPUTraceDataClass() GTMioMGPUTraceDataClass {
	return getGTMioMGPUTraceDataClass()
}

type GTMioMGPUTraceDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioMGPUTraceDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioMGPUTraceDataClass) Alloc() GTMioMGPUTraceData {
	rv := objc.Send[GTMioMGPUTraceData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioMGPUTraceData.CostCount]
//   - [GTMioMGPUTraceData.CostForScopeScopeIdentifierCost]
//   - [GTMioMGPUTraceData.Costs]
//   - [GTMioMGPUTraceData.Index]
//   - [GTMioMGPUTraceData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [GTMioMGPUTraceData.Kicks]
//   - [GTMioMGPUTraceData.KicksCount]
//   - [GTMioMGPUTraceData.TimelineCounters]
//   - [GTMioMGPUTraceData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [GTMioMGPUTraceData.TraceData]
//   - [GTMioMGPUTraceData.InitWithMGPUDataParent]
//   - [GTMioMGPUTraceData.DebugDescription]
//   - [GTMioMGPUTraceData.Description]
//   - [GTMioMGPUTraceData.Hash]
//   - [GTMioMGPUTraceData.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData
type GTMioMGPUTraceData struct {
	objectivec.Object
}

// GTMioMGPUTraceDataFromID constructs a [GTMioMGPUTraceData] from an objc.ID.
func GTMioMGPUTraceDataFromID(id objc.ID) GTMioMGPUTraceData {
	return GTMioMGPUTraceData{objectivec.Object{ID: id}}
}

// Ensure GTMioMGPUTraceData implements IGTMioMGPUTraceData.
var _ IGTMioMGPUTraceData = GTMioMGPUTraceData{}

// An interface definition for the [GTMioMGPUTraceData] class.
//
// # Methods
//
//   - [IGTMioMGPUTraceData.CostCount]
//   - [IGTMioMGPUTraceData.CostForScopeScopeIdentifierCost]
//   - [IGTMioMGPUTraceData.Costs]
//   - [IGTMioMGPUTraceData.Index]
//   - [IGTMioMGPUTraceData.InstructionCountForScopeScopeIdentifierDataMaster]
//   - [IGTMioMGPUTraceData.Kicks]
//   - [IGTMioMGPUTraceData.KicksCount]
//   - [IGTMioMGPUTraceData.TimelineCounters]
//   - [IGTMioMGPUTraceData.TotalCostForScopeScopeIdentifierDataMaster]
//   - [IGTMioMGPUTraceData.TraceData]
//   - [IGTMioMGPUTraceData.InitWithMGPUDataParent]
//   - [IGTMioMGPUTraceData.DebugDescription]
//   - [IGTMioMGPUTraceData.Description]
//   - [IGTMioMGPUTraceData.Hash]
//   - [IGTMioMGPUTraceData.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData
type IGTMioMGPUTraceData interface {
	objectivec.IObject

	// Topic: Methods

	CostCount() uint64
	CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool
	Costs() unsafe.Pointer
	Index() uint64
	InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64
	Kicks() unsafe.Pointer
	KicksCount() uint64
	TimelineCounters() IGTMioTimelineCounters
	TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64
	TraceData() objectivec.IObject
	InitWithMGPUDataParent(mGPUData unsafe.Pointer, parent objectivec.IObject) GTMioMGPUTraceData
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTMioMGPUTraceData) Init() GTMioMGPUTraceData {
	rv := objc.Send[GTMioMGPUTraceData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioMGPUTraceData) Autorelease() GTMioMGPUTraceData {
	rv := objc.Send[GTMioMGPUTraceData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioMGPUTraceData creates a new GTMioMGPUTraceData instance.
func NewGTMioMGPUTraceData() GTMioMGPUTraceData {
	class := getGTMioMGPUTraceDataClass()
	rv := objc.Send[GTMioMGPUTraceData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/initWithMGPUData:parent:
func NewGTMioMGPUTraceDataWithMGPUDataParent(mGPUData unsafe.Pointer, parent objectivec.IObject) GTMioMGPUTraceData {
	instance := getGTMioMGPUTraceDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMGPUData:parent:"), mGPUData, parent)
	return GTMioMGPUTraceDataFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/costForScope:scopeIdentifier:cost:
func (g GTMioMGPUTraceData) CostForScopeScopeIdentifierCost(scope uint16, identifier uint64, cost unsafe.Pointer) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("costForScope:scopeIdentifier:cost:"), scope, identifier, cost)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/instructionCountForScope:scopeIdentifier:dataMaster:
func (g GTMioMGPUTraceData) InstructionCountForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("instructionCountForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/totalCostForScope:scopeIdentifier:dataMaster:
func (g GTMioMGPUTraceData) TotalCostForScopeScopeIdentifierDataMaster(scope uint16, identifier uint64, master uint16) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("totalCostForScope:scopeIdentifier:dataMaster:"), scope, identifier, master)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/initWithMGPUData:parent:
func (g GTMioMGPUTraceData) InitWithMGPUDataParent(mGPUData unsafe.Pointer, parent objectivec.IObject) GTMioMGPUTraceData {
	rv := objc.Send[GTMioMGPUTraceData](g.ID, objc.Sel("initWithMGPUData:parent:"), mGPUData, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/costCount
func (g GTMioMGPUTraceData) CostCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("costCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/costs
func (g GTMioMGPUTraceData) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("costs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/debugDescription
func (g GTMioMGPUTraceData) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/description
func (g GTMioMGPUTraceData) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/hash
func (g GTMioMGPUTraceData) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/index
func (g GTMioMGPUTraceData) Index() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/kicks
func (g GTMioMGPUTraceData) Kicks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("kicks"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/kicksCount
func (g GTMioMGPUTraceData) KicksCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("kicksCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/superclass
func (g GTMioMGPUTraceData) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/timelineCounters
func (g GTMioMGPUTraceData) TimelineCounters() IGTMioTimelineCounters {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineCounters"))
	return GTMioTimelineCountersFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioMGPUTraceData/traceData
func (g GTMioMGPUTraceData) TraceData() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("traceData"))
	return objectivec.Object{ID: rv}
}
