// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XRGPUShaderInfo] class.
var (
	_XRGPUShaderInfoClass     XRGPUShaderInfoClass
	_XRGPUShaderInfoClassOnce sync.Once
)

func getXRGPUShaderInfoClass() XRGPUShaderInfoClass {
	_XRGPUShaderInfoClassOnce.Do(func() {
		_XRGPUShaderInfoClass = XRGPUShaderInfoClass{class: objc.GetClass("XRGPUShaderInfo")}
	})
	return _XRGPUShaderInfoClass
}

// GetXRGPUShaderInfoClass returns the class object for XRGPUShaderInfo.
func GetXRGPUShaderInfoClass() XRGPUShaderInfoClass {
	return getXRGPUShaderInfoClass()
}

type XRGPUShaderInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUShaderInfoClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUShaderInfoClass) Alloc() XRGPUShaderInfo {
	rv := objc.Send[XRGPUShaderInfo](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [XRGPUShaderInfo.Address]
//   - [XRGPUShaderInfo.SetAddress]
//   - [XRGPUShaderInfo.AvailableTime]
//   - [XRGPUShaderInfo.SetAvailableTime]
//   - [XRGPUShaderInfo.Label]
//   - [XRGPUShaderInfo.SetLabel]
//   - [XRGPUShaderInfo.Length]
//   - [XRGPUShaderInfo.SetLength]
//   - [XRGPUShaderInfo.Name]
//   - [XRGPUShaderInfo.SetName]
//   - [XRGPUShaderInfo.Pid]
//   - [XRGPUShaderInfo.SetPid]
//   - [XRGPUShaderInfo.PipelineLabel]
//   - [XRGPUShaderInfo.SetPipelineLabel]
//   - [XRGPUShaderInfo.PipelineStateId]
//   - [XRGPUShaderInfo.SetPipelineStateId]
//   - [XRGPUShaderInfo.ProgramType]
//   - [XRGPUShaderInfo.SetProgramType]
//   - [XRGPUShaderInfo.ShaderId]
//   - [XRGPUShaderInfo.SetShaderId]
//   - [XRGPUShaderInfo.InitWithShaderIdAddressLengthName]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo
type XRGPUShaderInfo struct {
	objectivec.Object
}

// XRGPUShaderInfoFromID constructs a [XRGPUShaderInfo] from an objc.ID.
func XRGPUShaderInfoFromID(id objc.ID) XRGPUShaderInfo {
	return XRGPUShaderInfo{objectivec.Object{ID: id}}
}

// Ensure XRGPUShaderInfo implements IXRGPUShaderInfo.
var _ IXRGPUShaderInfo = XRGPUShaderInfo{}

// An interface definition for the [XRGPUShaderInfo] class.
//
// # Methods
//
//   - [IXRGPUShaderInfo.Address]
//   - [IXRGPUShaderInfo.SetAddress]
//   - [IXRGPUShaderInfo.AvailableTime]
//   - [IXRGPUShaderInfo.SetAvailableTime]
//   - [IXRGPUShaderInfo.Label]
//   - [IXRGPUShaderInfo.SetLabel]
//   - [IXRGPUShaderInfo.Length]
//   - [IXRGPUShaderInfo.SetLength]
//   - [IXRGPUShaderInfo.Name]
//   - [IXRGPUShaderInfo.SetName]
//   - [IXRGPUShaderInfo.Pid]
//   - [IXRGPUShaderInfo.SetPid]
//   - [IXRGPUShaderInfo.PipelineLabel]
//   - [IXRGPUShaderInfo.SetPipelineLabel]
//   - [IXRGPUShaderInfo.PipelineStateId]
//   - [IXRGPUShaderInfo.SetPipelineStateId]
//   - [IXRGPUShaderInfo.ProgramType]
//   - [IXRGPUShaderInfo.SetProgramType]
//   - [IXRGPUShaderInfo.ShaderId]
//   - [IXRGPUShaderInfo.SetShaderId]
//   - [IXRGPUShaderInfo.InitWithShaderIdAddressLengthName]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo
type IXRGPUShaderInfo interface {
	objectivec.IObject

	// Topic: Methods

	Address() uint64
	SetAddress(value uint64)
	AvailableTime() uint64
	SetAvailableTime(value uint64)
	Label() string
	SetLabel(value string)
	Length() uint64
	SetLength(value uint64)
	Name() string
	SetName(value string)
	Pid() int
	SetPid(value int)
	PipelineLabel() string
	SetPipelineLabel(value string)
	PipelineStateId() uint64
	SetPipelineStateId(value uint64)
	ProgramType() string
	SetProgramType(value string)
	ShaderId() uint32
	SetShaderId(value uint32)
	InitWithShaderIdAddressLengthName(id uint32, address uint64, length uint64, name objectivec.IObject) XRGPUShaderInfo
}

// Init initializes the instance.
func (x XRGPUShaderInfo) Init() XRGPUShaderInfo {
	rv := objc.Send[XRGPUShaderInfo](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUShaderInfo) Autorelease() XRGPUShaderInfo {
	rv := objc.Send[XRGPUShaderInfo](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUShaderInfo creates a new XRGPUShaderInfo instance.
func NewXRGPUShaderInfo() XRGPUShaderInfo {
	class := getXRGPUShaderInfoClass()
	rv := objc.Send[XRGPUShaderInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/initWithShaderId:address:length:name:
func NewXRGPUShaderInfoWithShaderIdAddressLengthName(id uint32, address uint64, length uint64, name objectivec.IObject) XRGPUShaderInfo {
	instance := getXRGPUShaderInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShaderId:address:length:name:"), id, address, length, name)
	return XRGPUShaderInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/initWithShaderId:address:length:name:
func (x XRGPUShaderInfo) InitWithShaderIdAddressLengthName(id uint32, address uint64, length uint64, name objectivec.IObject) XRGPUShaderInfo {
	rv := objc.Send[XRGPUShaderInfo](x.ID, objc.Sel("initWithShaderId:address:length:name:"), id, address, length, name)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/address
func (x XRGPUShaderInfo) Address() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("address"))
	return rv
}
func (x XRGPUShaderInfo) SetAddress(value uint64) {
	objc.Send[struct{}](x.ID, objc.Sel("setAddress:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/availableTime
func (x XRGPUShaderInfo) AvailableTime() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("availableTime"))
	return rv
}
func (x XRGPUShaderInfo) SetAvailableTime(value uint64) {
	objc.Send[struct{}](x.ID, objc.Sel("setAvailableTime:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/label
func (x XRGPUShaderInfo) Label() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (x XRGPUShaderInfo) SetLabel(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/length
func (x XRGPUShaderInfo) Length() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("length"))
	return rv
}
func (x XRGPUShaderInfo) SetLength(value uint64) {
	objc.Send[struct{}](x.ID, objc.Sel("setLength:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/name
func (x XRGPUShaderInfo) Name() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (x XRGPUShaderInfo) SetName(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/pid
func (x XRGPUShaderInfo) Pid() int {
	rv := objc.Send[int](x.ID, objc.Sel("pid"))
	return rv
}
func (x XRGPUShaderInfo) SetPid(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setPid:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/pipelineLabel
func (x XRGPUShaderInfo) PipelineLabel() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("pipelineLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (x XRGPUShaderInfo) SetPipelineLabel(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setPipelineLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/pipelineStateId
func (x XRGPUShaderInfo) PipelineStateId() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("pipelineStateId"))
	return rv
}
func (x XRGPUShaderInfo) SetPipelineStateId(value uint64) {
	objc.Send[struct{}](x.ID, objc.Sel("setPipelineStateId:"), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/programType
func (x XRGPUShaderInfo) ProgramType() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("programType"))
	return foundation.NSStringFromID(rv).String()
}
func (x XRGPUShaderInfo) SetProgramType(value string) {
	objc.Send[struct{}](x.ID, objc.Sel("setProgramType:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUShaderInfo/shaderId
func (x XRGPUShaderInfo) ShaderId() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("shaderId"))
	return rv
}
func (x XRGPUShaderInfo) SetShaderId(value uint32) {
	objc.Send[struct{}](x.ID, objc.Sel("setShaderId:"), value)
}
