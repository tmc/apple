
// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

// Package gtshaderprofiler provides Go bindings for the gtshaderprofiler framework.
//
// # Key Types
//
//   - [GTShaderProfilerStreamData]
//   - [GTShaderProfilerBinaryAnalysisResult]
//   - [GTShaderProfilerCounterInfo]
//   - [GTShaderProfilerStreamDataProcessor]
//   - [GTShaderProfilerProcessedData]
//   - [GTShaderProfilerCounterGroupInfo]
//   - [GTShaderProfilerDeviceInfo]
//   - [GTShaderProfilerMCABinary]
//   - [GTShaderProfilerCounterSpec]
//   - [GTShaderProfilerDiassemblyRegisterPressure]
package gtshaderprofiler

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/Applications/Xcode.app/Contents/PlugIns/GPUDebugger.ideplugin/Contents/Frameworks/GTShaderProfiler.framework/Versions/A/GTShaderProfiler"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: gtshaderprofiler: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

