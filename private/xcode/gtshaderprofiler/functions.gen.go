// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: gtshaderprofiler: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: gtshaderprofiler: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: gtshaderprofiler: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _agxps_aps_clique_instruction_trace_get_execution_events func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer

// Agxps_aps_clique_instruction_trace_get_execution_events.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_execution_events
func Agxps_aps_clique_instruction_trace_get_execution_events(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer {
	if _agxps_aps_clique_instruction_trace_get_execution_events == nil {
		return nil
	}
	return _agxps_aps_clique_instruction_trace_get_execution_events(trace)
}

var _agxps_aps_clique_instruction_trace_get_execution_events_num func(trace AGXPSCliqueInstructionTraceRef) uint64

// Agxps_aps_clique_instruction_trace_get_execution_events_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_execution_events_num
func Agxps_aps_clique_instruction_trace_get_execution_events_num(trace AGXPSCliqueInstructionTraceRef) uint64 {
	if _agxps_aps_clique_instruction_trace_get_execution_events_num == nil {
		return 0
	}
	return _agxps_aps_clique_instruction_trace_get_execution_events_num(trace)
}

var _agxps_aps_clique_instruction_trace_get_instruction_stats func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer

// Agxps_aps_clique_instruction_trace_get_instruction_stats.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_instruction_stats
func Agxps_aps_clique_instruction_trace_get_instruction_stats(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer {
	if _agxps_aps_clique_instruction_trace_get_instruction_stats == nil {
		return nil
	}
	return _agxps_aps_clique_instruction_trace_get_instruction_stats(trace)
}

var _agxps_aps_clique_instruction_trace_get_pc_advances func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer

// Agxps_aps_clique_instruction_trace_get_pc_advances.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_pc_advances
func Agxps_aps_clique_instruction_trace_get_pc_advances(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer {
	if _agxps_aps_clique_instruction_trace_get_pc_advances == nil {
		return nil
	}
	return _agxps_aps_clique_instruction_trace_get_pc_advances(trace)
}

var _agxps_aps_clique_instruction_trace_get_pc_advances_num func(trace AGXPSCliqueInstructionTraceRef) uint64

// Agxps_aps_clique_instruction_trace_get_pc_advances_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_pc_advances_num
func Agxps_aps_clique_instruction_trace_get_pc_advances_num(trace AGXPSCliqueInstructionTraceRef) uint64 {
	if _agxps_aps_clique_instruction_trace_get_pc_advances_num == nil {
		return 0
	}
	return _agxps_aps_clique_instruction_trace_get_pc_advances_num(trace)
}

var _agxps_aps_clique_instruction_trace_get_timestamp_references func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer

// Agxps_aps_clique_instruction_trace_get_timestamp_references.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_timestamp_references
func Agxps_aps_clique_instruction_trace_get_timestamp_references(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer {
	if _agxps_aps_clique_instruction_trace_get_timestamp_references == nil {
		return nil
	}
	return _agxps_aps_clique_instruction_trace_get_timestamp_references(trace)
}

var _agxps_aps_clique_instruction_trace_get_timestamp_references_num func(trace AGXPSCliqueInstructionTraceRef) uint64

// Agxps_aps_clique_instruction_trace_get_timestamp_references_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_timestamp_references_num
func Agxps_aps_clique_instruction_trace_get_timestamp_references_num(trace AGXPSCliqueInstructionTraceRef) uint64 {
	if _agxps_aps_clique_instruction_trace_get_timestamp_references_num == nil {
		return 0
	}
	return _agxps_aps_clique_instruction_trace_get_timestamp_references_num(trace)
}

var _agxps_aps_clique_time_stats_create func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueTimeStatsRef

// Agxps_aps_clique_time_stats_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_time_stats_create
func Agxps_aps_clique_time_stats_create(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueTimeStatsRef {
	if _agxps_aps_clique_time_stats_create == nil {
		return 0
	}
	return _agxps_aps_clique_time_stats_create(profileData, cliqueIndex)
}

var _agxps_aps_descriptor_create func(descriptor unsafe.Pointer) AGXPSDescriptorRef

// Agxps_aps_descriptor_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_descriptor_create
func Agxps_aps_descriptor_create(descriptor unsafe.Pointer) AGXPSDescriptorRef {
	if _agxps_aps_descriptor_create == nil {
		return 0
	}
	return _agxps_aps_descriptor_create(descriptor)
}

var _agxps_aps_gpu_is_supported func(gpu AGXPSGPU) bool

// Agxps_aps_gpu_is_supported.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_gpu_is_supported
func Agxps_aps_gpu_is_supported(gpu AGXPSGPU) bool {
	if _agxps_aps_gpu_is_supported == nil {
		return false
	}
	return _agxps_aps_gpu_is_supported(gpu)
}

var _agxps_aps_parser_create func(descriptor AGXPSDescriptorRef) AGXPSParserHandle

// Agxps_aps_parser_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_create
func Agxps_aps_parser_create(descriptor AGXPSDescriptorRef) AGXPSParserHandle {
	if _agxps_aps_parser_create == nil {
		return 0
	}
	return _agxps_aps_parser_create(descriptor)
}

var _agxps_aps_parser_destroy func(parser AGXPSParserHandle)

// Agxps_aps_parser_destroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_destroy
func Agxps_aps_parser_destroy(parser AGXPSParserHandle) {
	if _agxps_aps_parser_destroy == nil {
		return
	}
	_agxps_aps_parser_destroy(parser)
}

var _agxps_aps_parser_is_valid func(parser AGXPSParserHandle) bool

// Agxps_aps_parser_is_valid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_is_valid
func Agxps_aps_parser_is_valid(parser AGXPSParserHandle) bool {
	if _agxps_aps_parser_is_valid == nil {
		return false
	}
	return _agxps_aps_parser_is_valid(parser)
}

var _agxps_aps_parser_parse func(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) int

// Agxps_aps_parser_parse.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_parse
func Agxps_aps_parser_parse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) int {
	if _agxps_aps_parser_parse == nil {
		return 0
	}
	return _agxps_aps_parser_parse(parser, data, size, profileDataOut)
}

var _agxps_aps_profile_data_destroy func(profileData AGXPSProfileData)

// Agxps_aps_profile_data_destroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_destroy
func Agxps_aps_profile_data_destroy(profileData AGXPSProfileData) {
	if _agxps_aps_profile_data_destroy == nil {
		return
	}
	_agxps_aps_profile_data_destroy(profileData)
}

var _agxps_aps_profile_data_get_esl_clique_clique_id func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_esl_clique_clique_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_clique_id
func Agxps_aps_profile_data_get_esl_clique_clique_id(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_esl_clique_clique_id == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_clique_clique_id(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_end func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_esl_clique_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_end
func Agxps_aps_profile_data_get_esl_clique_end(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_esl_clique_end == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_clique_end(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_esl_id func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_esl_clique_esl_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_esl_id
func Agxps_aps_profile_data_get_esl_clique_esl_id(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_esl_clique_esl_id == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_clique_esl_id(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_instruction_trace func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef

// Agxps_aps_profile_data_get_esl_clique_instruction_trace.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_instruction_trace
func Agxps_aps_profile_data_get_esl_clique_instruction_trace(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef {
	if _agxps_aps_profile_data_get_esl_clique_instruction_trace == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_clique_instruction_trace(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_kick_id func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_esl_clique_kick_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_kick_id
func Agxps_aps_profile_data_get_esl_clique_kick_id(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_esl_clique_kick_id == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_clique_kick_id(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_missing_end func(profileData AGXPSProfileData, cliqueIndex uint64) bool

// Agxps_aps_profile_data_get_esl_clique_missing_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_missing_end
func Agxps_aps_profile_data_get_esl_clique_missing_end(profileData AGXPSProfileData, cliqueIndex uint64) bool {
	if _agxps_aps_profile_data_get_esl_clique_missing_end == nil {
		return false
	}
	return _agxps_aps_profile_data_get_esl_clique_missing_end(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_start func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_esl_clique_start.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_start
func Agxps_aps_profile_data_get_esl_clique_start(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_esl_clique_start == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_clique_start(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_cliques_num func(profileData AGXPSProfileData) uint64

// Agxps_aps_profile_data_get_esl_cliques_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_cliques_num
func Agxps_aps_profile_data_get_esl_cliques_num(profileData AGXPSProfileData) uint64 {
	if _agxps_aps_profile_data_get_esl_cliques_num == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_esl_cliques_num(profileData)
}

var _agxps_aps_profile_data_get_kick_end func(profileData AGXPSProfileData, kickIndex uint64) uint64

// Agxps_aps_profile_data_get_kick_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_end
func Agxps_aps_profile_data_get_kick_end(profileData AGXPSProfileData, kickIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_kick_end == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_kick_end(profileData, kickIndex)
}

var _agxps_aps_profile_data_get_kick_id func(profileData AGXPSProfileData, kickIndex uint64) uint64

// Agxps_aps_profile_data_get_kick_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_id
func Agxps_aps_profile_data_get_kick_id(profileData AGXPSProfileData, kickIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_kick_id == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_kick_id(profileData, kickIndex)
}

var _agxps_aps_profile_data_get_kick_start func(profileData AGXPSProfileData, kickIndex uint64) uint64

// Agxps_aps_profile_data_get_kick_start.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_start
func Agxps_aps_profile_data_get_kick_start(profileData AGXPSProfileData, kickIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_kick_start == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_kick_start(profileData, kickIndex)
}

var _agxps_aps_profile_data_get_kicks_num func(profileData AGXPSProfileData) uint64

// Agxps_aps_profile_data_get_kicks_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kicks_num
func Agxps_aps_profile_data_get_kicks_num(profileData AGXPSProfileData) uint64 {
	if _agxps_aps_profile_data_get_kicks_num == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_kicks_num(profileData)
}

var _agxps_aps_profile_data_get_work_clique_end func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_work_clique_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_end
func Agxps_aps_profile_data_get_work_clique_end(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_work_clique_end == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_work_clique_end(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_work_clique_instruction_trace func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef

// Agxps_aps_profile_data_get_work_clique_instruction_trace.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_instruction_trace
func Agxps_aps_profile_data_get_work_clique_instruction_trace(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef {
	if _agxps_aps_profile_data_get_work_clique_instruction_trace == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_work_clique_instruction_trace(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_work_clique_start func(profileData AGXPSProfileData, cliqueIndex uint64) uint64

// Agxps_aps_profile_data_get_work_clique_start.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_start
func Agxps_aps_profile_data_get_work_clique_start(profileData AGXPSProfileData, cliqueIndex uint64) uint64 {
	if _agxps_aps_profile_data_get_work_clique_start == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_work_clique_start(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_work_cliques_num func(profileData AGXPSProfileData) uint64

// Agxps_aps_profile_data_get_work_cliques_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_cliques_num
func Agxps_aps_profile_data_get_work_cliques_num(profileData AGXPSProfileData) uint64 {
	if _agxps_aps_profile_data_get_work_cliques_num == nil {
		return 0
	}
	return _agxps_aps_profile_data_get_work_cliques_num(profileData)
}

var _agxps_aps_profile_data_is_valid func(profileData AGXPSProfileData) bool

// Agxps_aps_profile_data_is_valid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_is_valid
func Agxps_aps_profile_data_is_valid(profileData AGXPSProfileData) bool {
	if _agxps_aps_profile_data_is_valid == nil {
		return false
	}
	return _agxps_aps_profile_data_is_valid(profileData)
}

var _agxps_aps_timing_analyzer_get_num_commands func(analyzer uintptr) uint64

// Agxps_aps_timing_analyzer_get_num_commands.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_num_commands
func Agxps_aps_timing_analyzer_get_num_commands(analyzer uintptr) uint64 {
	if _agxps_aps_timing_analyzer_get_num_commands == nil {
		return 0
	}
	return _agxps_aps_timing_analyzer_get_num_commands(analyzer)
}

var _agxps_aps_timing_analyzer_get_num_work_cliques func(analyzer uintptr) uint64

// Agxps_aps_timing_analyzer_get_num_work_cliques.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_num_work_cliques
func Agxps_aps_timing_analyzer_get_num_work_cliques(analyzer uintptr) uint64 {
	if _agxps_aps_timing_analyzer_get_num_work_cliques == nil {
		return 0
	}
	return _agxps_aps_timing_analyzer_get_num_work_cliques(analyzer)
}

var _agxps_aps_timing_analyzer_get_work_cliques_average_duration func(analyzer uintptr) float64

// Agxps_aps_timing_analyzer_get_work_cliques_average_duration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_average_duration
func Agxps_aps_timing_analyzer_get_work_cliques_average_duration(analyzer uintptr) float64 {
	if _agxps_aps_timing_analyzer_get_work_cliques_average_duration == nil {
		return 0.0
	}
	return _agxps_aps_timing_analyzer_get_work_cliques_average_duration(analyzer)
}

var _agxps_aps_timing_analyzer_get_work_cliques_max_duration func(analyzer uintptr) float64

// Agxps_aps_timing_analyzer_get_work_cliques_max_duration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_max_duration
func Agxps_aps_timing_analyzer_get_work_cliques_max_duration(analyzer uintptr) float64 {
	if _agxps_aps_timing_analyzer_get_work_cliques_max_duration == nil {
		return 0.0
	}
	return _agxps_aps_timing_analyzer_get_work_cliques_max_duration(analyzer)
}

var _agxps_aps_timing_analyzer_get_work_cliques_min_duration func(analyzer uintptr) float64

// Agxps_aps_timing_analyzer_get_work_cliques_min_duration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_min_duration
func Agxps_aps_timing_analyzer_get_work_cliques_min_duration(analyzer uintptr) float64 {
	if _agxps_aps_timing_analyzer_get_work_cliques_min_duration == nil {
		return 0.0
	}
	return _agxps_aps_timing_analyzer_get_work_cliques_min_duration(analyzer)
}

var _agxps_gpu_create func(gen uint, variant uint, rev uint) AGXPSGPU

// Agxps_gpu_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_create
func Agxps_gpu_create(gen uint, variant uint, rev uint) AGXPSGPU {
	if _agxps_gpu_create == nil {
		return 0
	}
	return _agxps_gpu_create(gen, variant, rev)
}

var _agxps_gpu_destroy func(gpu AGXPSGPU)

// Agxps_gpu_destroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_destroy
func Agxps_gpu_destroy(gpu AGXPSGPU) {
	if _agxps_gpu_destroy == nil {
		return
	}
	_agxps_gpu_destroy(gpu)
}

var _agxps_gpu_format_name func(gpu AGXPSGPU, buf *byte, size uint64) int

// Agxps_gpu_format_name.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_format_name
func Agxps_gpu_format_name(gpu AGXPSGPU, buf *byte, size uint64) int {
	if _agxps_gpu_format_name == nil {
		return 0
	}
	return _agxps_gpu_format_name(gpu, buf, size)
}

var _agxps_gpu_get_gen func(gpu AGXPSGPU) uint

// Agxps_gpu_get_gen.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_gen
func Agxps_gpu_get_gen(gpu AGXPSGPU) uint {
	if _agxps_gpu_get_gen == nil {
		return 0
	}
	return _agxps_gpu_get_gen(gpu)
}

var _agxps_gpu_get_rev func(gpu AGXPSGPU) uint

// Agxps_gpu_get_rev.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_rev
func Agxps_gpu_get_rev(gpu AGXPSGPU) uint {
	if _agxps_gpu_get_rev == nil {
		return 0
	}
	return _agxps_gpu_get_rev(gpu)
}

var _agxps_gpu_get_variant func(gpu AGXPSGPU) uint

// Agxps_gpu_get_variant.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_variant
func Agxps_gpu_get_variant(gpu AGXPSGPU) uint {
	if _agxps_gpu_get_variant == nil {
		return 0
	}
	return _agxps_gpu_get_variant(gpu)
}

var _agxps_gpu_is_valid func(gpu AGXPSGPU) bool

// Agxps_gpu_is_valid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_is_valid
func Agxps_gpu_is_valid(gpu AGXPSGPU) bool {
	if _agxps_gpu_is_valid == nil {
		return false
	}
	return _agxps_gpu_is_valid(gpu)
}

var _agxps_initialize func() int

// Agxps_initialize.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_initialize
func Agxps_initialize() int {
	if _agxps_initialize == nil {
		return 0
	}
	return _agxps_initialize()
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_agxps_aps_clique_instruction_trace_get_execution_events, frameworkHandle, "agxps_aps_clique_instruction_trace_get_execution_events")
		registerFunc(&_agxps_aps_clique_instruction_trace_get_execution_events_num, frameworkHandle, "agxps_aps_clique_instruction_trace_get_execution_events_num")
		registerFunc(&_agxps_aps_clique_instruction_trace_get_instruction_stats, frameworkHandle, "agxps_aps_clique_instruction_trace_get_instruction_stats")
		registerFunc(&_agxps_aps_clique_instruction_trace_get_pc_advances, frameworkHandle, "agxps_aps_clique_instruction_trace_get_pc_advances")
		registerFunc(&_agxps_aps_clique_instruction_trace_get_pc_advances_num, frameworkHandle, "agxps_aps_clique_instruction_trace_get_pc_advances_num")
		registerFunc(&_agxps_aps_clique_instruction_trace_get_timestamp_references, frameworkHandle, "agxps_aps_clique_instruction_trace_get_timestamp_references")
		registerFunc(&_agxps_aps_clique_instruction_trace_get_timestamp_references_num, frameworkHandle, "agxps_aps_clique_instruction_trace_get_timestamp_references_num")
		registerFunc(&_agxps_aps_clique_time_stats_create, frameworkHandle, "agxps_aps_clique_time_stats_create")
		registerFunc(&_agxps_aps_descriptor_create, frameworkHandle, "agxps_aps_descriptor_create")
		registerFunc(&_agxps_aps_gpu_is_supported, frameworkHandle, "agxps_aps_gpu_is_supported")
		registerFunc(&_agxps_aps_parser_create, frameworkHandle, "agxps_aps_parser_create")
		registerFunc(&_agxps_aps_parser_destroy, frameworkHandle, "agxps_aps_parser_destroy")
		registerFunc(&_agxps_aps_parser_is_valid, frameworkHandle, "agxps_aps_parser_is_valid")
		registerFunc(&_agxps_aps_parser_parse, frameworkHandle, "agxps_aps_parser_parse")
		registerFunc(&_agxps_aps_profile_data_destroy, frameworkHandle, "agxps_aps_profile_data_destroy")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_clique_id, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_clique_id")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_end, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_end")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_esl_id, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_esl_id")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_instruction_trace, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_instruction_trace")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_kick_id, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_kick_id")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_missing_end, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_missing_end")
		registerFunc(&_agxps_aps_profile_data_get_esl_clique_start, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_start")
		registerFunc(&_agxps_aps_profile_data_get_esl_cliques_num, frameworkHandle, "agxps_aps_profile_data_get_esl_cliques_num")
		registerFunc(&_agxps_aps_profile_data_get_kick_end, frameworkHandle, "agxps_aps_profile_data_get_kick_end")
		registerFunc(&_agxps_aps_profile_data_get_kick_id, frameworkHandle, "agxps_aps_profile_data_get_kick_id")
		registerFunc(&_agxps_aps_profile_data_get_kick_start, frameworkHandle, "agxps_aps_profile_data_get_kick_start")
		registerFunc(&_agxps_aps_profile_data_get_kicks_num, frameworkHandle, "agxps_aps_profile_data_get_kicks_num")
		registerFunc(&_agxps_aps_profile_data_get_work_clique_end, frameworkHandle, "agxps_aps_profile_data_get_work_clique_end")
		registerFunc(&_agxps_aps_profile_data_get_work_clique_instruction_trace, frameworkHandle, "agxps_aps_profile_data_get_work_clique_instruction_trace")
		registerFunc(&_agxps_aps_profile_data_get_work_clique_start, frameworkHandle, "agxps_aps_profile_data_get_work_clique_start")
		registerFunc(&_agxps_aps_profile_data_get_work_cliques_num, frameworkHandle, "agxps_aps_profile_data_get_work_cliques_num")
		registerFunc(&_agxps_aps_profile_data_is_valid, frameworkHandle, "agxps_aps_profile_data_is_valid")
		registerFunc(&_agxps_aps_timing_analyzer_get_num_commands, frameworkHandle, "agxps_aps_timing_analyzer_get_num_commands")
		registerFunc(&_agxps_aps_timing_analyzer_get_num_work_cliques, frameworkHandle, "agxps_aps_timing_analyzer_get_num_work_cliques")
		registerFunc(&_agxps_aps_timing_analyzer_get_work_cliques_average_duration, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_average_duration")
		registerFunc(&_agxps_aps_timing_analyzer_get_work_cliques_max_duration, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_max_duration")
		registerFunc(&_agxps_aps_timing_analyzer_get_work_cliques_min_duration, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_min_duration")
		registerFunc(&_agxps_gpu_create, frameworkHandle, "agxps_gpu_create")
		registerFunc(&_agxps_gpu_destroy, frameworkHandle, "agxps_gpu_destroy")
		registerFunc(&_agxps_gpu_format_name, frameworkHandle, "agxps_gpu_format_name")
		registerFunc(&_agxps_gpu_get_gen, frameworkHandle, "agxps_gpu_get_gen")
		registerFunc(&_agxps_gpu_get_rev, frameworkHandle, "agxps_gpu_get_rev")
		registerFunc(&_agxps_gpu_get_variant, frameworkHandle, "agxps_gpu_get_variant")
		registerFunc(&_agxps_gpu_is_valid, frameworkHandle, "agxps_gpu_is_valid")
		registerFunc(&_agxps_initialize, frameworkHandle, "agxps_initialize")
	}

