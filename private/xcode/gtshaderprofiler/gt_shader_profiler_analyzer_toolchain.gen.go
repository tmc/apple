// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerAnalyzerToolchain] class.
var (
	_GTShaderProfilerAnalyzerToolchainClass     GTShaderProfilerAnalyzerToolchainClass
	_GTShaderProfilerAnalyzerToolchainClassOnce sync.Once
)

func getGTShaderProfilerAnalyzerToolchainClass() GTShaderProfilerAnalyzerToolchainClass {
	_GTShaderProfilerAnalyzerToolchainClassOnce.Do(func() {
		_GTShaderProfilerAnalyzerToolchainClass = GTShaderProfilerAnalyzerToolchainClass{class: objc.GetClass("GTShaderProfilerAnalyzerToolchain")}
	})
	return _GTShaderProfilerAnalyzerToolchainClass
}

// GetGTShaderProfilerAnalyzerToolchainClass returns the class object for GTShaderProfilerAnalyzerToolchain.
func GetGTShaderProfilerAnalyzerToolchainClass() GTShaderProfilerAnalyzerToolchainClass {
	return getGTShaderProfilerAnalyzerToolchainClass()
}

type GTShaderProfilerAnalyzerToolchainClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerAnalyzerToolchainClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerAnalyzerToolchainClass) Alloc() GTShaderProfilerAnalyzerToolchain {
	rv := objc.Send[GTShaderProfilerAnalyzerToolchain](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerAnalyzerToolchain.McaCommandLineToolPath]
//   - [GTShaderProfilerAnalyzerToolchain.SetMcaCommandLineToolPath]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzerToolchain
type GTShaderProfilerAnalyzerToolchain struct {
	objectivec.Object
}

// GTShaderProfilerAnalyzerToolchainFromID constructs a [GTShaderProfilerAnalyzerToolchain] from an objc.ID.
func GTShaderProfilerAnalyzerToolchainFromID(id objc.ID) GTShaderProfilerAnalyzerToolchain {
	return GTShaderProfilerAnalyzerToolchain{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerAnalyzerToolchain implements IGTShaderProfilerAnalyzerToolchain.
var _ IGTShaderProfilerAnalyzerToolchain = GTShaderProfilerAnalyzerToolchain{}

// An interface definition for the [GTShaderProfilerAnalyzerToolchain] class.
//
// # Methods
//
//   - [IGTShaderProfilerAnalyzerToolchain.McaCommandLineToolPath]
//   - [IGTShaderProfilerAnalyzerToolchain.SetMcaCommandLineToolPath]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzerToolchain
type IGTShaderProfilerAnalyzerToolchain interface {
	objectivec.IObject

	// Topic: Methods

	McaCommandLineToolPath() string
	SetMcaCommandLineToolPath(value string)
}

// Init initializes the instance.
func (g GTShaderProfilerAnalyzerToolchain) Init() GTShaderProfilerAnalyzerToolchain {
	rv := objc.Send[GTShaderProfilerAnalyzerToolchain](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerAnalyzerToolchain) Autorelease() GTShaderProfilerAnalyzerToolchain {
	rv := objc.Send[GTShaderProfilerAnalyzerToolchain](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerAnalyzerToolchain creates a new GTShaderProfilerAnalyzerToolchain instance.
func NewGTShaderProfilerAnalyzerToolchain() GTShaderProfilerAnalyzerToolchain {
	class := getGTShaderProfilerAnalyzerToolchainClass()
	rv := objc.Send[GTShaderProfilerAnalyzerToolchain](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzerToolchain/mcaCommandLineToolPath
func (g GTShaderProfilerAnalyzerToolchain) McaCommandLineToolPath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mcaCommandLineToolPath"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerAnalyzerToolchain) SetMcaCommandLineToolPath(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setMcaCommandLineToolPath:"), objc.String(value))
}
