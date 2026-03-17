// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerCounterGroupInfo] class.
var (
	_GTShaderProfilerCounterGroupInfoClass     GTShaderProfilerCounterGroupInfoClass
	_GTShaderProfilerCounterGroupInfoClassOnce sync.Once
)

func getGTShaderProfilerCounterGroupInfoClass() GTShaderProfilerCounterGroupInfoClass {
	_GTShaderProfilerCounterGroupInfoClassOnce.Do(func() {
		_GTShaderProfilerCounterGroupInfoClass = GTShaderProfilerCounterGroupInfoClass{class: objc.GetClass("GTShaderProfilerCounterGroupInfo")}
	})
	return _GTShaderProfilerCounterGroupInfoClass
}

// GetGTShaderProfilerCounterGroupInfoClass returns the class object for GTShaderProfilerCounterGroupInfo.
func GetGTShaderProfilerCounterGroupInfoClass() GTShaderProfilerCounterGroupInfoClass {
	return getGTShaderProfilerCounterGroupInfoClass()
}

type GTShaderProfilerCounterGroupInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerCounterGroupInfoClass) Alloc() GTShaderProfilerCounterGroupInfo {
	rv := objc.Send[GTShaderProfilerCounterGroupInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerCounterGroupInfo.BasisCondition]
//   - [GTShaderProfilerCounterGroupInfo.BasisCounter]
//   - [GTShaderProfilerCounterGroupInfo.Counters]
//   - [GTShaderProfilerCounterGroupInfo.DisplayStyle]
//   - [GTShaderProfilerCounterGroupInfo.MaskInCompute]
//   - [GTShaderProfilerCounterGroupInfo.Name]
//   - [GTShaderProfilerCounterGroupInfo.ResourceLink]
//   - [GTShaderProfilerCounterGroupInfo.SumCounterIndex]
//   - [GTShaderProfilerCounterGroupInfo.ValueRange]
//   - [GTShaderProfilerCounterGroupInfo.InitWithSpecParent]
//   - [GTShaderProfilerCounterGroupInfo.Description]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo
type GTShaderProfilerCounterGroupInfo struct {
	objectivec.Object
}

// GTShaderProfilerCounterGroupInfoFromID constructs a [GTShaderProfilerCounterGroupInfo] from an objc.ID.
func GTShaderProfilerCounterGroupInfoFromID(id objc.ID) GTShaderProfilerCounterGroupInfo {
	return GTShaderProfilerCounterGroupInfo{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerCounterGroupInfo implements IGTShaderProfilerCounterGroupInfo.
var _ IGTShaderProfilerCounterGroupInfo = GTShaderProfilerCounterGroupInfo{}

// An interface definition for the [GTShaderProfilerCounterGroupInfo] class.
//
// # Methods
//
//   - [IGTShaderProfilerCounterGroupInfo.BasisCondition]
//   - [IGTShaderProfilerCounterGroupInfo.BasisCounter]
//   - [IGTShaderProfilerCounterGroupInfo.Counters]
//   - [IGTShaderProfilerCounterGroupInfo.DisplayStyle]
//   - [IGTShaderProfilerCounterGroupInfo.MaskInCompute]
//   - [IGTShaderProfilerCounterGroupInfo.Name]
//   - [IGTShaderProfilerCounterGroupInfo.ResourceLink]
//   - [IGTShaderProfilerCounterGroupInfo.SumCounterIndex]
//   - [IGTShaderProfilerCounterGroupInfo.ValueRange]
//   - [IGTShaderProfilerCounterGroupInfo.InitWithSpecParent]
//   - [IGTShaderProfilerCounterGroupInfo.Description]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo
type IGTShaderProfilerCounterGroupInfo interface {
	objectivec.IObject

	// Topic: Methods

	BasisCondition() int
	BasisCounter() string
	Counters() foundation.INSArray
	DisplayStyle() uint64
	MaskInCompute() bool
	Name() string
	ResourceLink() uint64
	SumCounterIndex() uint64
	ValueRange() foundation.NSRange
	InitWithSpecParent(spec objectivec.IObject, parent objectivec.IObject) GTShaderProfilerCounterGroupInfo
	Description() string
}

// Init initializes the instance.
func (g GTShaderProfilerCounterGroupInfo) Init() GTShaderProfilerCounterGroupInfo {
	rv := objc.Send[GTShaderProfilerCounterGroupInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerCounterGroupInfo) Autorelease() GTShaderProfilerCounterGroupInfo {
	rv := objc.Send[GTShaderProfilerCounterGroupInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerCounterGroupInfo creates a new GTShaderProfilerCounterGroupInfo instance.
func NewGTShaderProfilerCounterGroupInfo() GTShaderProfilerCounterGroupInfo {
	class := getGTShaderProfilerCounterGroupInfoClass()
	rv := objc.Send[GTShaderProfilerCounterGroupInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/initWithSpec:parent:
func NewGTShaderProfilerCounterGroupInfoWithSpecParent(spec objectivec.IObject, parent objectivec.IObject) GTShaderProfilerCounterGroupInfo {
	instance := getGTShaderProfilerCounterGroupInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpec:parent:"), spec, parent)
	return GTShaderProfilerCounterGroupInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/initWithSpec:parent:
func (g GTShaderProfilerCounterGroupInfo) InitWithSpecParent(spec objectivec.IObject, parent objectivec.IObject) GTShaderProfilerCounterGroupInfo {
	rv := objc.Send[GTShaderProfilerCounterGroupInfo](g.ID, objc.Sel("initWithSpec:parent:"), spec, parent)
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/stringToResourceLink:
func (_GTShaderProfilerCounterGroupInfoClass GTShaderProfilerCounterGroupInfoClass) StringToResourceLink(link objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_GTShaderProfilerCounterGroupInfoClass.class), objc.Sel("stringToResourceLink:"), link)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/basisCondition
func (g GTShaderProfilerCounterGroupInfo) BasisCondition() int {
	rv := objc.Send[int](g.ID, objc.Sel("basisCondition"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/basisCounter
func (g GTShaderProfilerCounterGroupInfo) BasisCounter() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("basisCounter"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/counters
func (g GTShaderProfilerCounterGroupInfo) Counters() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("counters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/description
func (g GTShaderProfilerCounterGroupInfo) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/displayStyle
func (g GTShaderProfilerCounterGroupInfo) DisplayStyle() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("displayStyle"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/maskInCompute
func (g GTShaderProfilerCounterGroupInfo) MaskInCompute() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("maskInCompute"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/name
func (g GTShaderProfilerCounterGroupInfo) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/resourceLink
func (g GTShaderProfilerCounterGroupInfo) ResourceLink() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("resourceLink"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/sumCounterIndex
func (g GTShaderProfilerCounterGroupInfo) SumCounterIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("sumCounterIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterGroupInfo/valueRange
func (g GTShaderProfilerCounterGroupInfo) ValueRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](g.ID, objc.Sel("valueRange"))
	return foundation.NSRange(rv)
}

