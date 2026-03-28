// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderBinaryInfo] class.
var (
	_GTAGX2ShaderBinaryInfoClass     GTAGX2ShaderBinaryInfoClass
	_GTAGX2ShaderBinaryInfoClassOnce sync.Once
)

func getGTAGX2ShaderBinaryInfoClass() GTAGX2ShaderBinaryInfoClass {
	_GTAGX2ShaderBinaryInfoClassOnce.Do(func() {
		_GTAGX2ShaderBinaryInfoClass = GTAGX2ShaderBinaryInfoClass{class: objc.GetClass("GTAGX2ShaderBinaryInfo")}
	})
	return _GTAGX2ShaderBinaryInfoClass
}

// GetGTAGX2ShaderBinaryInfoClass returns the class object for GTAGX2ShaderBinaryInfo.
func GetGTAGX2ShaderBinaryInfoClass() GTAGX2ShaderBinaryInfoClass {
	return getGTAGX2ShaderBinaryInfoClass()
}

type GTAGX2ShaderBinaryInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderBinaryInfoClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderBinaryInfoClass) Alloc() GTAGX2ShaderBinaryInfo {
	rv := objc.Send[GTAGX2ShaderBinaryInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderBinaryInfo.AluBlockCount]
//   - [GTAGX2ShaderBinaryInfo.AnalysisResult]
//   - [GTAGX2ShaderBinaryInfo.Binary]
//   - [GTAGX2ShaderBinaryInfo.SetBinary]
//   - [GTAGX2ShaderBinaryInfo.Dylib]
//   - [GTAGX2ShaderBinaryInfo.SetDylib]
//   - [GTAGX2ShaderBinaryInfo.Key]
//   - [GTAGX2ShaderBinaryInfo.SetKey]
//   - [GTAGX2ShaderBinaryInfo.Type]
//   - [GTAGX2ShaderBinaryInfo.SetType]
//   - [GTAGX2ShaderBinaryInfo.UscSamples]
//   - [GTAGX2ShaderBinaryInfo.SetUscSamples]
//   - [GTAGX2ShaderBinaryInfo.InitWithKeyBinaryTypeDylibAnalysisResult]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo
type GTAGX2ShaderBinaryInfo struct {
	objectivec.Object
}

// GTAGX2ShaderBinaryInfoFromID constructs a [GTAGX2ShaderBinaryInfo] from an objc.ID.
func GTAGX2ShaderBinaryInfoFromID(id objc.ID) GTAGX2ShaderBinaryInfo {
	return GTAGX2ShaderBinaryInfo{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderBinaryInfo implements IGTAGX2ShaderBinaryInfo.
var _ IGTAGX2ShaderBinaryInfo = GTAGX2ShaderBinaryInfo{}

// An interface definition for the [GTAGX2ShaderBinaryInfo] class.
//
// # Methods
//
//   - [IGTAGX2ShaderBinaryInfo.AluBlockCount]
//   - [IGTAGX2ShaderBinaryInfo.AnalysisResult]
//   - [IGTAGX2ShaderBinaryInfo.Binary]
//   - [IGTAGX2ShaderBinaryInfo.SetBinary]
//   - [IGTAGX2ShaderBinaryInfo.Dylib]
//   - [IGTAGX2ShaderBinaryInfo.SetDylib]
//   - [IGTAGX2ShaderBinaryInfo.Key]
//   - [IGTAGX2ShaderBinaryInfo.SetKey]
//   - [IGTAGX2ShaderBinaryInfo.Type]
//   - [IGTAGX2ShaderBinaryInfo.SetType]
//   - [IGTAGX2ShaderBinaryInfo.UscSamples]
//   - [IGTAGX2ShaderBinaryInfo.SetUscSamples]
//   - [IGTAGX2ShaderBinaryInfo.InitWithKeyBinaryTypeDylibAnalysisResult]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo
type IGTAGX2ShaderBinaryInfo interface {
	objectivec.IObject

	// Topic: Methods

	AluBlockCount() uint32
	AnalysisResult() IGTShaderProfilerBinaryAnalysisResult
	Binary() foundation.INSData
	SetBinary(value foundation.INSData)
	Dylib() foundation.NSNumber
	SetDylib(value foundation.NSNumber)
	Key() string
	SetKey(value string)
	Type() string
	SetType(value string)
	UscSamples() foundation.NSMutableData
	SetUscSamples(value foundation.NSMutableData)
	InitWithKeyBinaryTypeDylibAnalysisResult(key objectivec.IObject, binary objectivec.IObject, type_ objectivec.IObject, dylib objectivec.IObject, result objectivec.IObject) GTAGX2ShaderBinaryInfo
}

// Init initializes the instance.
func (g GTAGX2ShaderBinaryInfo) Init() GTAGX2ShaderBinaryInfo {
	rv := objc.Send[GTAGX2ShaderBinaryInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderBinaryInfo) Autorelease() GTAGX2ShaderBinaryInfo {
	rv := objc.Send[GTAGX2ShaderBinaryInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderBinaryInfo creates a new GTAGX2ShaderBinaryInfo instance.
func NewGTAGX2ShaderBinaryInfo() GTAGX2ShaderBinaryInfo {
	class := getGTAGX2ShaderBinaryInfoClass()
	rv := objc.Send[GTAGX2ShaderBinaryInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/initWithKey:binary:type:dylib:analysisResult:
func NewGTAGX2ShaderBinaryInfoWithKeyBinaryTypeDylibAnalysisResult(key objectivec.IObject, binary objectivec.IObject, type_ objectivec.IObject, dylib objectivec.IObject, result objectivec.IObject) GTAGX2ShaderBinaryInfo {
	instance := getGTAGX2ShaderBinaryInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:binary:type:dylib:analysisResult:"), key, binary, type_, dylib, result)
	return GTAGX2ShaderBinaryInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/initWithKey:binary:type:dylib:analysisResult:
func (g GTAGX2ShaderBinaryInfo) InitWithKeyBinaryTypeDylibAnalysisResult(key objectivec.IObject, binary objectivec.IObject, type_ objectivec.IObject, dylib objectivec.IObject, result objectivec.IObject) GTAGX2ShaderBinaryInfo {
	rv := objc.Send[GTAGX2ShaderBinaryInfo](g.ID, objc.Sel("initWithKey:binary:type:dylib:analysisResult:"), key, binary, type_, dylib, result)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/aluBlockCount
func (g GTAGX2ShaderBinaryInfo) AluBlockCount() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("aluBlockCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/analysisResult
func (g GTAGX2ShaderBinaryInfo) AnalysisResult() IGTShaderProfilerBinaryAnalysisResult {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("analysisResult"))
	return GTShaderProfilerBinaryAnalysisResultFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/binary
func (g GTAGX2ShaderBinaryInfo) Binary() foundation.INSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binary"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTAGX2ShaderBinaryInfo) SetBinary(value foundation.INSData) {
	objc.Send[struct{}](g.ID, objc.Sel("setBinary:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/dylib
func (g GTAGX2ShaderBinaryInfo) Dylib() foundation.NSNumber {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dylib"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (g GTAGX2ShaderBinaryInfo) SetDylib(value foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setDylib:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/key
func (g GTAGX2ShaderBinaryInfo) Key() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryInfo) SetKey(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setKey:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/type
func (g GTAGX2ShaderBinaryInfo) Type() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderBinaryInfo) SetType(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setType:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderBinaryInfo/uscSamples
func (g GTAGX2ShaderBinaryInfo) UscSamples() foundation.NSMutableData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("uscSamples"))
	return foundation.NSMutableDataFromID(objc.ID(rv))
}
func (g GTAGX2ShaderBinaryInfo) SetUscSamples(value foundation.NSMutableData) {
	objc.Send[struct{}](g.ID, objc.Sel("setUscSamples:"), value)
}

