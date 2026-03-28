// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerDeviceInfo] class.
var (
	_GTShaderProfilerDeviceInfoClass     GTShaderProfilerDeviceInfoClass
	_GTShaderProfilerDeviceInfoClassOnce sync.Once
)

func getGTShaderProfilerDeviceInfoClass() GTShaderProfilerDeviceInfoClass {
	_GTShaderProfilerDeviceInfoClassOnce.Do(func() {
		_GTShaderProfilerDeviceInfoClass = GTShaderProfilerDeviceInfoClass{class: objc.GetClass("GTShaderProfilerDeviceInfo")}
	})
	return _GTShaderProfilerDeviceInfoClass
}

// GetGTShaderProfilerDeviceInfoClass returns the class object for GTShaderProfilerDeviceInfo.
func GetGTShaderProfilerDeviceInfoClass() GTShaderProfilerDeviceInfoClass {
	return getGTShaderProfilerDeviceInfoClass()
}

type GTShaderProfilerDeviceInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerDeviceInfoClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerDeviceInfoClass) Alloc() GTShaderProfilerDeviceInfo {
	rv := objc.Send[GTShaderProfilerDeviceInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerDeviceInfo.Build]
//   - [GTShaderProfilerDeviceInfo.SetBuild]
//   - [GTShaderProfilerDeviceInfo.EncodeWithCoder]
//   - [GTShaderProfilerDeviceInfo.MetalVersion]
//   - [GTShaderProfilerDeviceInfo.SetMetalVersion]
//   - [GTShaderProfilerDeviceInfo.Name]
//   - [GTShaderProfilerDeviceInfo.SetName]
//   - [GTShaderProfilerDeviceInfo.Platform]
//   - [GTShaderProfilerDeviceInfo.SetPlatform]
//   - [GTShaderProfilerDeviceInfo.InitWithCoder]
//   - [GTShaderProfilerDeviceInfo.Version]
//   - [GTShaderProfilerDeviceInfo.SetVersion]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo
type GTShaderProfilerDeviceInfo struct {
	objectivec.Object
}

// GTShaderProfilerDeviceInfoFromID constructs a [GTShaderProfilerDeviceInfo] from an objc.ID.
func GTShaderProfilerDeviceInfoFromID(id objc.ID) GTShaderProfilerDeviceInfo {
	return GTShaderProfilerDeviceInfo{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerDeviceInfo implements IGTShaderProfilerDeviceInfo.
var _ IGTShaderProfilerDeviceInfo = GTShaderProfilerDeviceInfo{}

// An interface definition for the [GTShaderProfilerDeviceInfo] class.
//
// # Methods
//
//   - [IGTShaderProfilerDeviceInfo.Build]
//   - [IGTShaderProfilerDeviceInfo.SetBuild]
//   - [IGTShaderProfilerDeviceInfo.EncodeWithCoder]
//   - [IGTShaderProfilerDeviceInfo.MetalVersion]
//   - [IGTShaderProfilerDeviceInfo.SetMetalVersion]
//   - [IGTShaderProfilerDeviceInfo.Name]
//   - [IGTShaderProfilerDeviceInfo.SetName]
//   - [IGTShaderProfilerDeviceInfo.Platform]
//   - [IGTShaderProfilerDeviceInfo.SetPlatform]
//   - [IGTShaderProfilerDeviceInfo.InitWithCoder]
//   - [IGTShaderProfilerDeviceInfo.Version]
//   - [IGTShaderProfilerDeviceInfo.SetVersion]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo
type IGTShaderProfilerDeviceInfo interface {
	objectivec.IObject

	// Topic: Methods

	Build() string
	SetBuild(value string)
	EncodeWithCoder(coder foundation.INSCoder)
	MetalVersion() string
	SetMetalVersion(value string)
	Name() string
	SetName(value string)
	Platform() int
	SetPlatform(value int)
	InitWithCoder(coder foundation.INSCoder) GTShaderProfilerDeviceInfo
	Version() string
	SetVersion(value string)
}

// Init initializes the instance.
func (g GTShaderProfilerDeviceInfo) Init() GTShaderProfilerDeviceInfo {
	rv := objc.Send[GTShaderProfilerDeviceInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerDeviceInfo) Autorelease() GTShaderProfilerDeviceInfo {
	rv := objc.Send[GTShaderProfilerDeviceInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerDeviceInfo creates a new GTShaderProfilerDeviceInfo instance.
func NewGTShaderProfilerDeviceInfo() GTShaderProfilerDeviceInfo {
	class := getGTShaderProfilerDeviceInfoClass()
	rv := objc.Send[GTShaderProfilerDeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/initWithCoder:
func NewGTShaderProfilerDeviceInfoWithCoder(coder objectivec.IObject) GTShaderProfilerDeviceInfo {
	instance := getGTShaderProfilerDeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerDeviceInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/encodeWithCoder:
func (g GTShaderProfilerDeviceInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/initWithCoder:
func (g GTShaderProfilerDeviceInfo) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerDeviceInfo {
	rv := objc.Send[GTShaderProfilerDeviceInfo](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/supportsSecureCoding
func (_GTShaderProfilerDeviceInfoClass GTShaderProfilerDeviceInfoClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTShaderProfilerDeviceInfoClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/build
func (g GTShaderProfilerDeviceInfo) Build() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("build"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerDeviceInfo) SetBuild(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setBuild:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/metalVersion
func (g GTShaderProfilerDeviceInfo) MetalVersion() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalVersion"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerDeviceInfo) SetMetalVersion(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setMetalVersion:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/name
func (g GTShaderProfilerDeviceInfo) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerDeviceInfo) SetName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/platform
func (g GTShaderProfilerDeviceInfo) Platform() int {
	rv := objc.Send[int](g.ID, objc.Sel("platform"))
	return rv
}
func (g GTShaderProfilerDeviceInfo) SetPlatform(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setPlatform:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDeviceInfo/version
func (g GTShaderProfilerDeviceInfo) Version() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerDeviceInfo) SetVersion(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setVersion:"), objc.String(value))
}

