// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerCounterInfo] class.
var (
	_GTShaderProfilerCounterInfoClass     GTShaderProfilerCounterInfoClass
	_GTShaderProfilerCounterInfoClassOnce sync.Once
)

func getGTShaderProfilerCounterInfoClass() GTShaderProfilerCounterInfoClass {
	_GTShaderProfilerCounterInfoClassOnce.Do(func() {
		_GTShaderProfilerCounterInfoClass = GTShaderProfilerCounterInfoClass{class: objc.GetClass("GTShaderProfilerCounterInfo")}
	})
	return _GTShaderProfilerCounterInfoClass
}

// GetGTShaderProfilerCounterInfoClass returns the class object for GTShaderProfilerCounterInfo.
func GetGTShaderProfilerCounterInfoClass() GTShaderProfilerCounterInfoClass {
	return getGTShaderProfilerCounterInfoClass()
}

type GTShaderProfilerCounterInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerCounterInfoClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerCounterInfoClass) Alloc() GTShaderProfilerCounterInfo {
	rv := objc.Send[GTShaderProfilerCounterInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerCounterInfo.BatchIdFilterable]
//   - [GTShaderProfilerCounterInfo.DataType]
//   - [GTShaderProfilerCounterInfo.SetDataType]
//   - [GTShaderProfilerCounterInfo.DisplayStyle]
//   - [GTShaderProfilerCounterInfo.ExternallyVisible]
//   - [GTShaderProfilerCounterInfo.FormatTimeValueUnitValue]
//   - [GTShaderProfilerCounterInfo.FormatValueContainsUnit]
//   - [GTShaderProfilerCounterInfo.GroupTag]
//   - [GTShaderProfilerCounterInfo.MaskInCompute]
//   - [GTShaderProfilerCounterInfo.MaskInDraw]
//   - [GTShaderProfilerCounterInfo.Name]
//   - [GTShaderProfilerCounterInfo.NonOverlappingEncoderDataIndex]
//   - [GTShaderProfilerCounterInfo.SetNonOverlappingEncoderDataIndex]
//   - [GTShaderProfilerCounterInfo.NonOverlappingGPUCommandDataIndex]
//   - [GTShaderProfilerCounterInfo.SetNonOverlappingGPUCommandDataIndex]
//   - [GTShaderProfilerCounterInfo.RequiresBatchIDFiltering]
//   - [GTShaderProfilerCounterInfo.SetRequiresBatchIDFiltering]
//   - [GTShaderProfilerCounterInfo.StringFromMemoryByteCount]
//   - [GTShaderProfilerCounterInfo.ToolsCounterName]
//   - [GTShaderProfilerCounterInfo.Unit]
//   - [GTShaderProfilerCounterInfo.VendorCounterNames]
//   - [GTShaderProfilerCounterInfo.InitWithSpecParent]
//   - [GTShaderProfilerCounterInfo.Description]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo
type GTShaderProfilerCounterInfo struct {
	objectivec.Object
}

// GTShaderProfilerCounterInfoFromID constructs a [GTShaderProfilerCounterInfo] from an objc.ID.
func GTShaderProfilerCounterInfoFromID(id objc.ID) GTShaderProfilerCounterInfo {
	return GTShaderProfilerCounterInfo{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerCounterInfo implements IGTShaderProfilerCounterInfo.
var _ IGTShaderProfilerCounterInfo = GTShaderProfilerCounterInfo{}

// An interface definition for the [GTShaderProfilerCounterInfo] class.
//
// # Methods
//
//   - [IGTShaderProfilerCounterInfo.BatchIdFilterable]
//   - [IGTShaderProfilerCounterInfo.DataType]
//   - [IGTShaderProfilerCounterInfo.SetDataType]
//   - [IGTShaderProfilerCounterInfo.DisplayStyle]
//   - [IGTShaderProfilerCounterInfo.ExternallyVisible]
//   - [IGTShaderProfilerCounterInfo.FormatTimeValueUnitValue]
//   - [IGTShaderProfilerCounterInfo.FormatValueContainsUnit]
//   - [IGTShaderProfilerCounterInfo.GroupTag]
//   - [IGTShaderProfilerCounterInfo.MaskInCompute]
//   - [IGTShaderProfilerCounterInfo.MaskInDraw]
//   - [IGTShaderProfilerCounterInfo.Name]
//   - [IGTShaderProfilerCounterInfo.NonOverlappingEncoderDataIndex]
//   - [IGTShaderProfilerCounterInfo.SetNonOverlappingEncoderDataIndex]
//   - [IGTShaderProfilerCounterInfo.NonOverlappingGPUCommandDataIndex]
//   - [IGTShaderProfilerCounterInfo.SetNonOverlappingGPUCommandDataIndex]
//   - [IGTShaderProfilerCounterInfo.RequiresBatchIDFiltering]
//   - [IGTShaderProfilerCounterInfo.SetRequiresBatchIDFiltering]
//   - [IGTShaderProfilerCounterInfo.StringFromMemoryByteCount]
//   - [IGTShaderProfilerCounterInfo.ToolsCounterName]
//   - [IGTShaderProfilerCounterInfo.Unit]
//   - [IGTShaderProfilerCounterInfo.VendorCounterNames]
//   - [IGTShaderProfilerCounterInfo.InitWithSpecParent]
//   - [IGTShaderProfilerCounterInfo.Description]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo
type IGTShaderProfilerCounterInfo interface {
	objectivec.IObject

	// Topic: Methods

	BatchIdFilterable() bool
	DataType() uint64
	SetDataType(value uint64)
	DisplayStyle() uint64
	ExternallyVisible() bool
	FormatTimeValueUnitValue(value int64, value2 int64) objectivec.IObject
	FormatValueContainsUnit(value float32, unit unsafe.Pointer) objectivec.IObject
	GroupTag() string
	MaskInCompute() bool
	MaskInDraw() bool
	Name() string
	NonOverlappingEncoderDataIndex() int
	SetNonOverlappingEncoderDataIndex(value int)
	NonOverlappingGPUCommandDataIndex() int
	SetNonOverlappingGPUCommandDataIndex(value int)
	RequiresBatchIDFiltering() bool
	SetRequiresBatchIDFiltering(value bool)
	StringFromMemoryByteCount(count uint64) objectivec.IObject
	ToolsCounterName() string
	Unit() string
	VendorCounterNames() foundation.INSArray
	InitWithSpecParent(spec objectivec.IObject, parent objectivec.IObject) GTShaderProfilerCounterInfo
	Description() string
}

// Init initializes the instance.
func (g GTShaderProfilerCounterInfo) Init() GTShaderProfilerCounterInfo {
	rv := objc.Send[GTShaderProfilerCounterInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerCounterInfo) Autorelease() GTShaderProfilerCounterInfo {
	rv := objc.Send[GTShaderProfilerCounterInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerCounterInfo creates a new GTShaderProfilerCounterInfo instance.
func NewGTShaderProfilerCounterInfo() GTShaderProfilerCounterInfo {
	class := getGTShaderProfilerCounterInfoClass()
	rv := objc.Send[GTShaderProfilerCounterInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/initWithSpec:parent:
func NewGTShaderProfilerCounterInfoWithSpecParent(spec objectivec.IObject, parent objectivec.IObject) GTShaderProfilerCounterInfo {
	instance := getGTShaderProfilerCounterInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpec:parent:"), spec, parent)
	return GTShaderProfilerCounterInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/formatTimeValue:unitValue:
func (g GTShaderProfilerCounterInfo) FormatTimeValueUnitValue(value int64, value2 int64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("formatTimeValue:unitValue:"), value, value2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/formatValue:containsUnit:
func (g GTShaderProfilerCounterInfo) FormatValueContainsUnit(value float32, unit unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("formatValue:containsUnit:"), value, unit)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/stringFromMemoryByteCount:
func (g GTShaderProfilerCounterInfo) StringFromMemoryByteCount(count uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("stringFromMemoryByteCount:"), count)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/initWithSpec:parent:
func (g GTShaderProfilerCounterInfo) InitWithSpecParent(spec objectivec.IObject, parent objectivec.IObject) GTShaderProfilerCounterInfo {
	rv := objc.Send[GTShaderProfilerCounterInfo](g.ID, objc.Sel("initWithSpec:parent:"), spec, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/stringToDataType:
func (_GTShaderProfilerCounterInfoClass GTShaderProfilerCounterInfoClass) StringToDataType(type_ objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_GTShaderProfilerCounterInfoClass.class), objc.Sel("stringToDataType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/stringToDisplayStyle:
func (_GTShaderProfilerCounterInfoClass GTShaderProfilerCounterInfoClass) StringToDisplayStyle(style objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_GTShaderProfilerCounterInfoClass.class), objc.Sel("stringToDisplayStyle:"), style)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/batchIdFilterable
func (g GTShaderProfilerCounterInfo) BatchIdFilterable() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("batchIdFilterable"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/dataType
func (g GTShaderProfilerCounterInfo) DataType() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("dataType"))
	return rv
}
func (g GTShaderProfilerCounterInfo) SetDataType(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setDataType:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/description
func (g GTShaderProfilerCounterInfo) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/displayStyle
func (g GTShaderProfilerCounterInfo) DisplayStyle() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("displayStyle"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/externallyVisible
func (g GTShaderProfilerCounterInfo) ExternallyVisible() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("externallyVisible"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/groupTag
func (g GTShaderProfilerCounterInfo) GroupTag() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("groupTag"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/maskInCompute
func (g GTShaderProfilerCounterInfo) MaskInCompute() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("maskInCompute"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/maskInDraw
func (g GTShaderProfilerCounterInfo) MaskInDraw() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("maskInDraw"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/name
func (g GTShaderProfilerCounterInfo) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/nonOverlappingEncoderDataIndex
func (g GTShaderProfilerCounterInfo) NonOverlappingEncoderDataIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("nonOverlappingEncoderDataIndex"))
	return rv
}
func (g GTShaderProfilerCounterInfo) SetNonOverlappingEncoderDataIndex(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setNonOverlappingEncoderDataIndex:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/nonOverlappingGPUCommandDataIndex
func (g GTShaderProfilerCounterInfo) NonOverlappingGPUCommandDataIndex() int {
	rv := objc.Send[int](g.ID, objc.Sel("nonOverlappingGPUCommandDataIndex"))
	return rv
}
func (g GTShaderProfilerCounterInfo) SetNonOverlappingGPUCommandDataIndex(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setNonOverlappingGPUCommandDataIndex:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/requiresBatchIDFiltering
func (g GTShaderProfilerCounterInfo) RequiresBatchIDFiltering() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("requiresBatchIDFiltering"))
	return rv
}
func (g GTShaderProfilerCounterInfo) SetRequiresBatchIDFiltering(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setRequiresBatchIDFiltering:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/toolsCounterName
func (g GTShaderProfilerCounterInfo) ToolsCounterName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("toolsCounterName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/unit
func (g GTShaderProfilerCounterInfo) Unit() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unit"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerCounterInfo/vendorCounterNames
func (g GTShaderProfilerCounterInfo) VendorCounterNames() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("vendorCounterNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
