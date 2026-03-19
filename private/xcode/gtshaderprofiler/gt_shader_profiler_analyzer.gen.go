// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
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

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerAnalyzerClass) Alloc() GTShaderProfilerAnalyzer {
	rv := objc.Send[GTShaderProfilerAnalyzer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
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
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/generateFullMCAReport:
func (g GTShaderProfilerAnalyzer) GenerateFullMCAReport(mCAReport VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(mCAReport)
	defer _cleanup0()
	objc.Send[objc.ID](g.ID, objc.Sel("generateFullMCAReport:"), _block0)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/generateMCAOutput:callback:
func (g GTShaderProfilerAnalyzer) GenerateMCAOutputCallback(mCAOutput bool, callback VoidHandler) {
_block1, _cleanup1 := NewVoidBlock(callback)
	defer _cleanup1()
	objc.Send[objc.ID](g.ID, objc.Sel("generateMCAOutput:callback:"), mCAOutput, _block1)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerAnalyzer/generateRegisterPressureView:
func (g GTShaderProfilerAnalyzer) GenerateRegisterPressureView(view VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(view)
	defer _cleanup0()
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

