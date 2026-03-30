// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

// Package gtshaderprofiler provides Go bindings for the gtshaderprofiler framework.
//
// # Key Types
//
//   - [GTMioTraceTimelineData]
//   - [GTMioTraceData]
//   - [XRGPUAPSDataProcessor]
//   - [GTShaderProfilerStreamData]
//   - [GTMioShaderBinaryData]
//   - [GTMioUSCTraceData]
//   - [GTMioTraceDataHelper]
//   - [GTShaderProfilerBinaryAnalysisResult]
//   - [GTAGX2ShaderBinary]
//   - [GTAGX2ShaderProfilerResult]
package gtshaderprofiler

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the gtshaderprofiler library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/Applications/Xcode.app/Contents/PlugIns/GPUDebugger.ideplugin/Contents/Frameworks/GTShaderProfiler.framework/Versions/A/GTShaderProfiler"}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: gtshaderprofiler: failed to load framework from any known path\n")
}
