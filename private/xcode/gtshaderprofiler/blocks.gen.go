// Code generated from Apple documentation. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
)

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [GTShaderProfilerAnalyzer.GenerateFullMCAReport]
//   - [GTShaderProfilerAnalyzer.GenerateMCAOutputCallback]
//   - [GTShaderProfilerAnalyzer.GenerateRegisterPressureView]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [GTShaderProfilerAnalyzer.GenerateFullMCAReport]
//   - [GTShaderProfilerAnalyzer.GenerateMCAOutputCallback]
//   - [GTShaderProfilerAnalyzer.GenerateRegisterPressureView]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

