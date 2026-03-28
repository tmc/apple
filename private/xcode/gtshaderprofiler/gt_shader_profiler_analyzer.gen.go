// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerAnalyzer] class.
var (
	_GTShaderProfilerAnalyzerClass     GTShaderProfilerAnalyzerClass
	_GTShaderProfilerAnalyzerClassOnce sync.Once
)

func getGTShaderProfilerAnalyzerClass() GTShaderProfilerAnalyzerClass {
	_GTShaderProfilerAnalyzerClassOnce.Do(func() {
		_GTShaderProfilerAnalyzerClass = GTShaderProfilerAnalyzerClass{class: objc.GetClass("GTShaderProfilerAnalyzer")}
	})
	return _GTShaderProfilerAnalyzerClass
}

// GetGTShaderProfilerAnalyzerClass returns the class object for GTShaderProfilerAnalyzer.
func GetGTShaderProfilerAnalyzerClass() GTShaderProfilerAnalyzerClass {
	return getGTShaderProfilerAnalyzerClass()
}

type GTShaderProfilerAnalyzerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerAnalyzerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerAnalyzerClass) Alloc() GTShaderProfilerAnalyzer {
	rv := objc.Send[GTShaderProfilerAnalyzer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerAnalyzer._executeTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError]
//   - [GTShaderProfilerAnalyzer._generateMCAOutputSync]
//   - [GTShaderProfilerAnalyzer.Binary]
//   - [GTShaderProfilerAnalyzer.GenerateFullMCAReport]
//   - [GTShaderProfilerAnalyzer.GenerateMCAOutputCallback]
//   - [GTShaderProfilerAnalyzer.GenerateRegisterPressureView]
//   - [GTShaderProfilerAnalyzer.InitWithToolchainBinaryGpu]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer
type GTShaderProfilerAnalyzer struct {
	objectivec.Object
}

// GTShaderProfilerAnalyzerFromID constructs a [GTShaderProfilerAnalyzer] from an objc.ID.
func GTShaderProfilerAnalyzerFromID(id objc.ID) GTShaderProfilerAnalyzer {
	return GTShaderProfilerAnalyzer{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerAnalyzer implements IGTShaderProfilerAnalyzer.
var _ IGTShaderProfilerAnalyzer = GTShaderProfilerAnalyzer{}

// An interface definition for the [GTShaderProfilerAnalyzer] class.
//
// # Methods
//
//   - [IGTShaderProfilerAnalyzer._executeTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError]
//   - [IGTShaderProfilerAnalyzer._generateMCAOutputSync]
//   - [IGTShaderProfilerAnalyzer.Binary]
//   - [IGTShaderProfilerAnalyzer.GenerateFullMCAReport]
//   - [IGTShaderProfilerAnalyzer.GenerateMCAOutputCallback]
//   - [IGTShaderProfilerAnalyzer.GenerateRegisterPressureView]
//   - [IGTShaderProfilerAnalyzer.InitWithToolchainBinaryGpu]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer
type IGTShaderProfilerAnalyzer interface {
	objectivec.IObject

	// Topic: Methods

	_executeTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError(task objectivec.IObject, arguments objectivec.IObject, environment objectivec.IObject, output []objectivec.IObject, directory objectivec.IObject, description objectivec.IObject) (bool, error)
	_generateMCAOutputSync(sync bool) objectivec.IObject
	Binary() objectivec.IObject
	GenerateFullMCAReport(mCAReport VoidHandler)
	GenerateMCAOutputCallback(mCAOutput bool, callback VoidHandler)
	GenerateRegisterPressureView(view VoidHandler)
	InitWithToolchainBinaryGpu(toolchain objectivec.IObject, binary objectivec.IObject, gpu uint32) GTShaderProfilerAnalyzer
}

// Init initializes the instance.
func (g GTShaderProfilerAnalyzer) Init() GTShaderProfilerAnalyzer {
	rv := objc.Send[GTShaderProfilerAnalyzer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerAnalyzer) Autorelease() GTShaderProfilerAnalyzer {
	rv := objc.Send[GTShaderProfilerAnalyzer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerAnalyzer creates a new GTShaderProfilerAnalyzer instance.
func NewGTShaderProfilerAnalyzer() GTShaderProfilerAnalyzer {
	class := getGTShaderProfilerAnalyzerClass()
	rv := objc.Send[GTShaderProfilerAnalyzer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/initWithToolchain:binary:gpu:
func NewGTShaderProfilerAnalyzerWithToolchainBinaryGpu(toolchain objectivec.IObject, binary objectivec.IObject, gpu uint32) GTShaderProfilerAnalyzer {
	instance := getGTShaderProfilerAnalyzerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithToolchain:binary:gpu:"), toolchain, binary, gpu)
	return GTShaderProfilerAnalyzerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/_executeTask:arguments:environment:standardOutput:workingDirectory:description:error:
func (g GTShaderProfilerAnalyzer) _executeTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError(task objectivec.IObject, arguments objectivec.IObject, environment objectivec.IObject, output []objectivec.IObject, directory objectivec.IObject, description objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("_executeTask:arguments:environment:standardOutput:workingDirectory:description:error:"), task, arguments, environment, objectivec.IObjectSliceToNSArray(output), directory, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_executeTask:arguments:environment:standardOutput:workingDirectory:description:error: returned NO with nil NSError")
	}
	return rv, nil

}

// ExecuteTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError is an exported wrapper for the private method _executeTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError.
func (g GTShaderProfilerAnalyzer) ExecuteTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError(task objectivec.IObject, arguments objectivec.IObject, environment objectivec.IObject, output []objectivec.IObject, directory objectivec.IObject, description objectivec.IObject) (bool, error) {
	return g._executeTaskArgumentsEnvironmentStandardOutputWorkingDirectoryDescriptionError(task, arguments, environment, output, directory, description)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/_generateMCAOutputSync:
func (g GTShaderProfilerAnalyzer) _generateMCAOutputSync(sync bool) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_generateMCAOutputSync:"), sync)
	return objectivec.Object{ID: rv}
}

// GenerateMCAOutputSync is an exported wrapper for the private method _generateMCAOutputSync.
func (g GTShaderProfilerAnalyzer) GenerateMCAOutputSync(sync bool) objectivec.IObject {
	return g._generateMCAOutputSync(sync)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/generateFullMCAReport:
func (g GTShaderProfilerAnalyzer) GenerateFullMCAReport(mCAReport VoidHandler) {
_block0, _ := NewVoidBlock(mCAReport)
	objc.Send[objc.ID](g.ID, objc.Sel("generateFullMCAReport:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/generateMCAOutput:callback:
func (g GTShaderProfilerAnalyzer) GenerateMCAOutputCallback(mCAOutput bool, callback VoidHandler) {
_block1, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](g.ID, objc.Sel("generateMCAOutput:callback:"), mCAOutput, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/generateRegisterPressureView:
func (g GTShaderProfilerAnalyzer) GenerateRegisterPressureView(view VoidHandler) {
_block0, _ := NewVoidBlock(view)
	objc.Send[objc.ID](g.ID, objc.Sel("generateRegisterPressureView:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/initWithToolchain:binary:gpu:
func (g GTShaderProfilerAnalyzer) InitWithToolchainBinaryGpu(toolchain objectivec.IObject, binary objectivec.IObject, gpu uint32) GTShaderProfilerAnalyzer {
	rv := objc.Send[GTShaderProfilerAnalyzer](g.ID, objc.Sel("initWithToolchain:binary:gpu:"), toolchain, binary, gpu)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/binary
func (g GTShaderProfilerAnalyzer) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}

// GenerateFullMCAReportSync is a synchronous wrapper around [GTShaderProfilerAnalyzer.GenerateFullMCAReport].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTShaderProfilerAnalyzer) GenerateFullMCAReportSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.GenerateFullMCAReport(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GenerateMCAOutputCallbackSync is a synchronous wrapper around [GTShaderProfilerAnalyzer.GenerateMCAOutputCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTShaderProfilerAnalyzer) GenerateMCAOutputCallbackSync(ctx context.Context, mCAOutput bool) error {
	done := make(chan struct{}, 1)
	g.GenerateMCAOutputCallback(mCAOutput, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GenerateRegisterPressureViewSync is a synchronous wrapper around [GTShaderProfilerAnalyzer.GenerateRegisterPressureView].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTShaderProfilerAnalyzer) GenerateRegisterPressureViewSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.GenerateRegisterPressureView(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

