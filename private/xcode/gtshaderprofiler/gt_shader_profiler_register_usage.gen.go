// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerRegisterUsage] class.
var (
	_GTShaderProfilerRegisterUsageClass     GTShaderProfilerRegisterUsageClass
	_GTShaderProfilerRegisterUsageClassOnce sync.Once
)

func getGTShaderProfilerRegisterUsageClass() GTShaderProfilerRegisterUsageClass {
	_GTShaderProfilerRegisterUsageClassOnce.Do(func() {
		_GTShaderProfilerRegisterUsageClass = GTShaderProfilerRegisterUsageClass{class: objc.GetClass("GTShaderProfilerRegisterUsage")}
	})
	return _GTShaderProfilerRegisterUsageClass
}

// GetGTShaderProfilerRegisterUsageClass returns the class object for GTShaderProfilerRegisterUsage.
func GetGTShaderProfilerRegisterUsageClass() GTShaderProfilerRegisterUsageClass {
	return getGTShaderProfilerRegisterUsageClass()
}

type GTShaderProfilerRegisterUsageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerRegisterUsageClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerRegisterUsageClass) Alloc() GTShaderProfilerRegisterUsage {
	rv := objc.Send[GTShaderProfilerRegisterUsage](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerRegisterUsage.Combine]
//   - [GTShaderProfilerRegisterUsage.Size]
//   - [GTShaderProfilerRegisterUsage.Test]
//   - [GTShaderProfilerRegisterUsage.InitWithBitsetString]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage
type GTShaderProfilerRegisterUsage struct {
	objectivec.Object
}

// GTShaderProfilerRegisterUsageFromID constructs a [GTShaderProfilerRegisterUsage] from an objc.ID.
func GTShaderProfilerRegisterUsageFromID(id objc.ID) GTShaderProfilerRegisterUsage {
	return GTShaderProfilerRegisterUsage{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerRegisterUsage implements IGTShaderProfilerRegisterUsage.
var _ IGTShaderProfilerRegisterUsage = GTShaderProfilerRegisterUsage{}

// An interface definition for the [GTShaderProfilerRegisterUsage] class.
//
// # Methods
//
//   - [IGTShaderProfilerRegisterUsage.Combine]
//   - [IGTShaderProfilerRegisterUsage.Size]
//   - [IGTShaderProfilerRegisterUsage.Test]
//   - [IGTShaderProfilerRegisterUsage.InitWithBitsetString]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage
type IGTShaderProfilerRegisterUsage interface {
	objectivec.IObject

	// Topic: Methods

	Combine(combine objectivec.IObject)
	Size() uint64
	Test(test uint64) bool
	InitWithBitsetString(string_ objectivec.IObject) GTShaderProfilerRegisterUsage
}

// Init initializes the instance.
func (g GTShaderProfilerRegisterUsage) Init() GTShaderProfilerRegisterUsage {
	rv := objc.Send[GTShaderProfilerRegisterUsage](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerRegisterUsage) Autorelease() GTShaderProfilerRegisterUsage {
	rv := objc.Send[GTShaderProfilerRegisterUsage](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerRegisterUsage creates a new GTShaderProfilerRegisterUsage instance.
func NewGTShaderProfilerRegisterUsage() GTShaderProfilerRegisterUsage {
	class := getGTShaderProfilerRegisterUsageClass()
	rv := objc.Send[GTShaderProfilerRegisterUsage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage/initWithBitsetString:
func NewGTShaderProfilerRegisterUsageWithBitsetString(string_ objectivec.IObject) GTShaderProfilerRegisterUsage {
	instance := getGTShaderProfilerRegisterUsageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBitsetString:"), string_)
	return GTShaderProfilerRegisterUsageFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage/combine:
func (g GTShaderProfilerRegisterUsage) Combine(combine objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("combine:"), combine)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage/test:
func (g GTShaderProfilerRegisterUsage) Test(test uint64) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("test:"), test)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage/initWithBitsetString:
func (g GTShaderProfilerRegisterUsage) InitWithBitsetString(string_ objectivec.IObject) GTShaderProfilerRegisterUsage {
	rv := objc.Send[GTShaderProfilerRegisterUsage](g.ID, objc.Sel("initWithBitsetString:"), string_)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerRegisterUsage/size
func (g GTShaderProfilerRegisterUsage) Size() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("size"))
	return rv
}

