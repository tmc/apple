// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("gtshaderprofiler: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("gtshaderprofiler: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("gtshaderprofiler: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("gtshaderprofiler: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _agxps_aps_clique_instruction_trace_get_execution_events func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxps_aps_clique_instruction_trace_get_execution_eventsErr error

func tryAgxps_aps_clique_instruction_trace_get_execution_events(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxps_aps_clique_instruction_trace_get_execution_events == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_execution_events", "", _agxps_aps_clique_instruction_trace_get_execution_eventsErr)
	}
	return _agxps_aps_clique_instruction_trace_get_execution_events(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_execution_events.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_execution_events
func Agxps_aps_clique_instruction_trace_get_execution_events(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxps_aps_clique_instruction_trace_get_execution_events(trace)
}

var _agxps_aps_clique_instruction_trace_get_execution_events_num func(trace AGXPSCliqueInstructionTraceRef) uint64
var _agxps_aps_clique_instruction_trace_get_execution_events_numErr error

func tryAgxps_aps_clique_instruction_trace_get_execution_events_num(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	if _agxps_aps_clique_instruction_trace_get_execution_events_num == nil {
		return 0, symbolCallError("agxps_aps_clique_instruction_trace_get_execution_events_num", "", _agxps_aps_clique_instruction_trace_get_execution_events_numErr)
	}
	return _agxps_aps_clique_instruction_trace_get_execution_events_num(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_execution_events_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_execution_events_num
func Agxps_aps_clique_instruction_trace_get_execution_events_num(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	return tryAgxps_aps_clique_instruction_trace_get_execution_events_num(trace)
}

var _agxps_aps_clique_instruction_trace_get_instruction_stats func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxps_aps_clique_instruction_trace_get_instruction_statsErr error

func tryAgxps_aps_clique_instruction_trace_get_instruction_stats(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxps_aps_clique_instruction_trace_get_instruction_stats == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_instruction_stats", "", _agxps_aps_clique_instruction_trace_get_instruction_statsErr)
	}
	return _agxps_aps_clique_instruction_trace_get_instruction_stats(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_instruction_stats.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_instruction_stats
func Agxps_aps_clique_instruction_trace_get_instruction_stats(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxps_aps_clique_instruction_trace_get_instruction_stats(trace)
}

var _agxps_aps_clique_instruction_trace_get_pc_advances func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxps_aps_clique_instruction_trace_get_pc_advancesErr error

func tryAgxps_aps_clique_instruction_trace_get_pc_advances(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxps_aps_clique_instruction_trace_get_pc_advances == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_pc_advances", "", _agxps_aps_clique_instruction_trace_get_pc_advancesErr)
	}
	return _agxps_aps_clique_instruction_trace_get_pc_advances(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_pc_advances.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_pc_advances
func Agxps_aps_clique_instruction_trace_get_pc_advances(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxps_aps_clique_instruction_trace_get_pc_advances(trace)
}

var _agxps_aps_clique_instruction_trace_get_pc_advances_num func(trace AGXPSCliqueInstructionTraceRef) uint64
var _agxps_aps_clique_instruction_trace_get_pc_advances_numErr error

func tryAgxps_aps_clique_instruction_trace_get_pc_advances_num(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	if _agxps_aps_clique_instruction_trace_get_pc_advances_num == nil {
		return 0, symbolCallError("agxps_aps_clique_instruction_trace_get_pc_advances_num", "", _agxps_aps_clique_instruction_trace_get_pc_advances_numErr)
	}
	return _agxps_aps_clique_instruction_trace_get_pc_advances_num(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_pc_advances_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_pc_advances_num
func Agxps_aps_clique_instruction_trace_get_pc_advances_num(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	return tryAgxps_aps_clique_instruction_trace_get_pc_advances_num(trace)
}

var _agxps_aps_clique_instruction_trace_get_timestamp_references func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxps_aps_clique_instruction_trace_get_timestamp_referencesErr error

func tryAgxps_aps_clique_instruction_trace_get_timestamp_references(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxps_aps_clique_instruction_trace_get_timestamp_references == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_timestamp_references", "", _agxps_aps_clique_instruction_trace_get_timestamp_referencesErr)
	}
	return _agxps_aps_clique_instruction_trace_get_timestamp_references(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_timestamp_references.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_timestamp_references
func Agxps_aps_clique_instruction_trace_get_timestamp_references(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxps_aps_clique_instruction_trace_get_timestamp_references(trace)
}

var _agxps_aps_clique_instruction_trace_get_timestamp_references_num func(trace AGXPSCliqueInstructionTraceRef) uint64
var _agxps_aps_clique_instruction_trace_get_timestamp_references_numErr error

func tryAgxps_aps_clique_instruction_trace_get_timestamp_references_num(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	if _agxps_aps_clique_instruction_trace_get_timestamp_references_num == nil {
		return 0, symbolCallError("agxps_aps_clique_instruction_trace_get_timestamp_references_num", "", _agxps_aps_clique_instruction_trace_get_timestamp_references_numErr)
	}
	return _agxps_aps_clique_instruction_trace_get_timestamp_references_num(trace), nil
}

// Agxps_aps_clique_instruction_trace_get_timestamp_references_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_timestamp_references_num
func Agxps_aps_clique_instruction_trace_get_timestamp_references_num(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	return tryAgxps_aps_clique_instruction_trace_get_timestamp_references_num(trace)
}

var _agxps_aps_clique_time_stats_create func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueTimeStatsRef
var _agxps_aps_clique_time_stats_createErr error

func tryAgxps_aps_clique_time_stats_create(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueTimeStatsRef, error) {
	if _agxps_aps_clique_time_stats_create == nil {
		return 0, symbolCallError("agxps_aps_clique_time_stats_create", "", _agxps_aps_clique_time_stats_createErr)
	}
	return _agxps_aps_clique_time_stats_create(profileData, cliqueIndex), nil
}

// Agxps_aps_clique_time_stats_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_time_stats_create
func Agxps_aps_clique_time_stats_create(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueTimeStatsRef, error) {
	return tryAgxps_aps_clique_time_stats_create(profileData, cliqueIndex)
}

var _agxps_aps_descriptor_create func(descriptor unsafe.Pointer) AGXPSDescriptorRef
var _agxps_aps_descriptor_createErr error

func tryAgxps_aps_descriptor_create(descriptor unsafe.Pointer) (AGXPSDescriptorRef, error) {
	if _agxps_aps_descriptor_create == nil {
		return 0, symbolCallError("agxps_aps_descriptor_create", "", _agxps_aps_descriptor_createErr)
	}
	return _agxps_aps_descriptor_create(descriptor), nil
}

// Agxps_aps_descriptor_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_descriptor_create
func Agxps_aps_descriptor_create(descriptor unsafe.Pointer) (AGXPSDescriptorRef, error) {
	return tryAgxps_aps_descriptor_create(descriptor)
}

var _agxps_aps_gpu_is_supported func(gpu AGXPSGPU) bool
var _agxps_aps_gpu_is_supportedErr error

func tryAgxps_aps_gpu_is_supported(gpu AGXPSGPU) (bool, error) {
	if _agxps_aps_gpu_is_supported == nil {
		return false, symbolCallError("agxps_aps_gpu_is_supported", "", _agxps_aps_gpu_is_supportedErr)
	}
	return _agxps_aps_gpu_is_supported(gpu), nil
}

// Agxps_aps_gpu_is_supported.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_gpu_is_supported
func Agxps_aps_gpu_is_supported(gpu AGXPSGPU) (bool, error) {
	return tryAgxps_aps_gpu_is_supported(gpu)
}

var _agxps_aps_parser_create func(descriptor AGXPSDescriptorRef) AGXPSParserHandle
var _agxps_aps_parser_createErr error

func tryAgxps_aps_parser_create(descriptor AGXPSDescriptorRef) (AGXPSParserHandle, error) {
	if _agxps_aps_parser_create == nil {
		return *new(AGXPSParserHandle), symbolCallError("agxps_aps_parser_create", "", _agxps_aps_parser_createErr)
	}
	return _agxps_aps_parser_create(descriptor), nil
}

// Agxps_aps_parser_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_create
func Agxps_aps_parser_create(descriptor AGXPSDescriptorRef) (AGXPSParserHandle, error) {
	return tryAgxps_aps_parser_create(descriptor)
}

var _agxps_aps_parser_destroy func(parser AGXPSParserHandle)
var _agxps_aps_parser_destroyErr error

func tryAgxps_aps_parser_destroy(parser AGXPSParserHandle) error {
	if _agxps_aps_parser_destroy == nil {
		return symbolCallError("agxps_aps_parser_destroy", "", _agxps_aps_parser_destroyErr)
	}
	_agxps_aps_parser_destroy(parser)
	return nil
}

// Agxps_aps_parser_destroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_destroy
func Agxps_aps_parser_destroy(parser AGXPSParserHandle) error {
	return tryAgxps_aps_parser_destroy(parser)
}

var _agxps_aps_parser_is_valid func(parser AGXPSParserHandle) bool
var _agxps_aps_parser_is_validErr error

func tryAgxps_aps_parser_is_valid(parser AGXPSParserHandle) (bool, error) {
	if _agxps_aps_parser_is_valid == nil {
		return false, symbolCallError("agxps_aps_parser_is_valid", "", _agxps_aps_parser_is_validErr)
	}
	return _agxps_aps_parser_is_valid(parser), nil
}

// Agxps_aps_parser_is_valid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_is_valid
func Agxps_aps_parser_is_valid(parser AGXPSParserHandle) (bool, error) {
	return tryAgxps_aps_parser_is_valid(parser)
}

var _agxps_aps_parser_parse func(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) int
var _agxps_aps_parser_parseErr error

func tryAgxps_aps_parser_parse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) (int, error) {
	if _agxps_aps_parser_parse == nil {
		return 0, symbolCallError("agxps_aps_parser_parse", "", _agxps_aps_parser_parseErr)
	}
	return _agxps_aps_parser_parse(parser, data, size, profileDataOut), nil
}

// Agxps_aps_parser_parse.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_parse
func Agxps_aps_parser_parse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) (int, error) {
	return tryAgxps_aps_parser_parse(parser, data, size, profileDataOut)
}

var _agxps_aps_profile_data_destroy func(profileData AGXPSProfileData)
var _agxps_aps_profile_data_destroyErr error

func tryAgxps_aps_profile_data_destroy(profileData AGXPSProfileData) error {
	if _agxps_aps_profile_data_destroy == nil {
		return symbolCallError("agxps_aps_profile_data_destroy", "", _agxps_aps_profile_data_destroyErr)
	}
	_agxps_aps_profile_data_destroy(profileData)
	return nil
}

// Agxps_aps_profile_data_destroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_destroy
func Agxps_aps_profile_data_destroy(profileData AGXPSProfileData) error {
	return tryAgxps_aps_profile_data_destroy(profileData)
}

var _agxps_aps_profile_data_get_esl_clique_clique_id func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_esl_clique_clique_idErr error

func tryAgxps_aps_profile_data_get_esl_clique_clique_id(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_esl_clique_clique_id == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_clique_id", "", _agxps_aps_profile_data_get_esl_clique_clique_idErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_clique_id(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_clique_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_clique_id
func Agxps_aps_profile_data_get_esl_clique_clique_id(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_clique_id(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_end func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_esl_clique_endErr error

func tryAgxps_aps_profile_data_get_esl_clique_end(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_esl_clique_end == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_end", "", _agxps_aps_profile_data_get_esl_clique_endErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_end(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_end
func Agxps_aps_profile_data_get_esl_clique_end(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_end(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_esl_id func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_esl_clique_esl_idErr error

func tryAgxps_aps_profile_data_get_esl_clique_esl_id(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_esl_clique_esl_id == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_esl_id", "", _agxps_aps_profile_data_get_esl_clique_esl_idErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_esl_id(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_esl_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_esl_id
func Agxps_aps_profile_data_get_esl_clique_esl_id(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_esl_id(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_instruction_trace func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef
var _agxps_aps_profile_data_get_esl_clique_instruction_traceErr error

func tryAgxps_aps_profile_data_get_esl_clique_instruction_trace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	if _agxps_aps_profile_data_get_esl_clique_instruction_trace == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_instruction_trace", "", _agxps_aps_profile_data_get_esl_clique_instruction_traceErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_instruction_trace(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_instruction_trace.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_instruction_trace
func Agxps_aps_profile_data_get_esl_clique_instruction_trace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_instruction_trace(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_kick_id func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_esl_clique_kick_idErr error

func tryAgxps_aps_profile_data_get_esl_clique_kick_id(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_esl_clique_kick_id == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_kick_id", "", _agxps_aps_profile_data_get_esl_clique_kick_idErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_kick_id(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_kick_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_kick_id
func Agxps_aps_profile_data_get_esl_clique_kick_id(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_kick_id(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_missing_end func(profileData AGXPSProfileData, cliqueIndex uint64) bool
var _agxps_aps_profile_data_get_esl_clique_missing_endErr error

func tryAgxps_aps_profile_data_get_esl_clique_missing_end(profileData AGXPSProfileData, cliqueIndex uint64) (bool, error) {
	if _agxps_aps_profile_data_get_esl_clique_missing_end == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_missing_end", "", _agxps_aps_profile_data_get_esl_clique_missing_endErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_missing_end(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_missing_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_missing_end
func Agxps_aps_profile_data_get_esl_clique_missing_end(profileData AGXPSProfileData, cliqueIndex uint64) (bool, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_missing_end(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_clique_start func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_esl_clique_startErr error

func tryAgxps_aps_profile_data_get_esl_clique_start(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_esl_clique_start == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_start", "", _agxps_aps_profile_data_get_esl_clique_startErr)
	}
	return _agxps_aps_profile_data_get_esl_clique_start(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_esl_clique_start.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_start
func Agxps_aps_profile_data_get_esl_clique_start(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_esl_clique_start(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_esl_cliques_num func(profileData AGXPSProfileData) uint64
var _agxps_aps_profile_data_get_esl_cliques_numErr error

func tryAgxps_aps_profile_data_get_esl_cliques_num(profileData AGXPSProfileData) (uint64, error) {
	if _agxps_aps_profile_data_get_esl_cliques_num == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_cliques_num", "", _agxps_aps_profile_data_get_esl_cliques_numErr)
	}
	return _agxps_aps_profile_data_get_esl_cliques_num(profileData), nil
}

// Agxps_aps_profile_data_get_esl_cliques_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_cliques_num
func Agxps_aps_profile_data_get_esl_cliques_num(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxps_aps_profile_data_get_esl_cliques_num(profileData)
}

var _agxps_aps_profile_data_get_kick_end func(profileData AGXPSProfileData, kickIndex uint64) uint64
var _agxps_aps_profile_data_get_kick_endErr error

func tryAgxps_aps_profile_data_get_kick_end(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_kick_end == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kick_end", "", _agxps_aps_profile_data_get_kick_endErr)
	}
	return _agxps_aps_profile_data_get_kick_end(profileData, kickIndex), nil
}

// Agxps_aps_profile_data_get_kick_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_end
func Agxps_aps_profile_data_get_kick_end(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_kick_end(profileData, kickIndex)
}

var _agxps_aps_profile_data_get_kick_id func(profileData AGXPSProfileData, kickIndex uint64) uint64
var _agxps_aps_profile_data_get_kick_idErr error

func tryAgxps_aps_profile_data_get_kick_id(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_kick_id == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kick_id", "", _agxps_aps_profile_data_get_kick_idErr)
	}
	return _agxps_aps_profile_data_get_kick_id(profileData, kickIndex), nil
}

// Agxps_aps_profile_data_get_kick_id.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_id
func Agxps_aps_profile_data_get_kick_id(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_kick_id(profileData, kickIndex)
}

var _agxps_aps_profile_data_get_kick_start func(profileData AGXPSProfileData, kickIndex uint64) uint64
var _agxps_aps_profile_data_get_kick_startErr error

func tryAgxps_aps_profile_data_get_kick_start(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_kick_start == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kick_start", "", _agxps_aps_profile_data_get_kick_startErr)
	}
	return _agxps_aps_profile_data_get_kick_start(profileData, kickIndex), nil
}

// Agxps_aps_profile_data_get_kick_start.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_start
func Agxps_aps_profile_data_get_kick_start(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_kick_start(profileData, kickIndex)
}

var _agxps_aps_profile_data_get_kicks_num func(profileData AGXPSProfileData) uint64
var _agxps_aps_profile_data_get_kicks_numErr error

func tryAgxps_aps_profile_data_get_kicks_num(profileData AGXPSProfileData) (uint64, error) {
	if _agxps_aps_profile_data_get_kicks_num == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kicks_num", "", _agxps_aps_profile_data_get_kicks_numErr)
	}
	return _agxps_aps_profile_data_get_kicks_num(profileData), nil
}

// Agxps_aps_profile_data_get_kicks_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kicks_num
func Agxps_aps_profile_data_get_kicks_num(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxps_aps_profile_data_get_kicks_num(profileData)
}

var _agxps_aps_profile_data_get_work_clique_end func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_work_clique_endErr error

func tryAgxps_aps_profile_data_get_work_clique_end(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_work_clique_end == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_clique_end", "", _agxps_aps_profile_data_get_work_clique_endErr)
	}
	return _agxps_aps_profile_data_get_work_clique_end(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_work_clique_end.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_end
func Agxps_aps_profile_data_get_work_clique_end(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_work_clique_end(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_work_clique_instruction_trace func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef
var _agxps_aps_profile_data_get_work_clique_instruction_traceErr error

func tryAgxps_aps_profile_data_get_work_clique_instruction_trace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	if _agxps_aps_profile_data_get_work_clique_instruction_trace == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_clique_instruction_trace", "", _agxps_aps_profile_data_get_work_clique_instruction_traceErr)
	}
	return _agxps_aps_profile_data_get_work_clique_instruction_trace(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_work_clique_instruction_trace.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_instruction_trace
func Agxps_aps_profile_data_get_work_clique_instruction_trace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	return tryAgxps_aps_profile_data_get_work_clique_instruction_trace(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_work_clique_start func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxps_aps_profile_data_get_work_clique_startErr error

func tryAgxps_aps_profile_data_get_work_clique_start(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxps_aps_profile_data_get_work_clique_start == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_clique_start", "", _agxps_aps_profile_data_get_work_clique_startErr)
	}
	return _agxps_aps_profile_data_get_work_clique_start(profileData, cliqueIndex), nil
}

// Agxps_aps_profile_data_get_work_clique_start.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_start
func Agxps_aps_profile_data_get_work_clique_start(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxps_aps_profile_data_get_work_clique_start(profileData, cliqueIndex)
}

var _agxps_aps_profile_data_get_work_cliques_num func(profileData AGXPSProfileData) uint64
var _agxps_aps_profile_data_get_work_cliques_numErr error

func tryAgxps_aps_profile_data_get_work_cliques_num(profileData AGXPSProfileData) (uint64, error) {
	if _agxps_aps_profile_data_get_work_cliques_num == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_cliques_num", "", _agxps_aps_profile_data_get_work_cliques_numErr)
	}
	return _agxps_aps_profile_data_get_work_cliques_num(profileData), nil
}

// Agxps_aps_profile_data_get_work_cliques_num.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_cliques_num
func Agxps_aps_profile_data_get_work_cliques_num(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxps_aps_profile_data_get_work_cliques_num(profileData)
}

var _agxps_aps_profile_data_is_valid func(profileData AGXPSProfileData) bool
var _agxps_aps_profile_data_is_validErr error

func tryAgxps_aps_profile_data_is_valid(profileData AGXPSProfileData) (bool, error) {
	if _agxps_aps_profile_data_is_valid == nil {
		return false, symbolCallError("agxps_aps_profile_data_is_valid", "", _agxps_aps_profile_data_is_validErr)
	}
	return _agxps_aps_profile_data_is_valid(profileData), nil
}

// Agxps_aps_profile_data_is_valid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_is_valid
func Agxps_aps_profile_data_is_valid(profileData AGXPSProfileData) (bool, error) {
	return tryAgxps_aps_profile_data_is_valid(profileData)
}

var _agxps_aps_timing_analyzer_get_num_commands func(analyzer uintptr) uint64
var _agxps_aps_timing_analyzer_get_num_commandsErr error

func tryAgxps_aps_timing_analyzer_get_num_commands(analyzer uintptr) (uint64, error) {
	if _agxps_aps_timing_analyzer_get_num_commands == nil {
		return 0, symbolCallError("agxps_aps_timing_analyzer_get_num_commands", "", _agxps_aps_timing_analyzer_get_num_commandsErr)
	}
	return _agxps_aps_timing_analyzer_get_num_commands(analyzer), nil
}

// Agxps_aps_timing_analyzer_get_num_commands.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_num_commands
func Agxps_aps_timing_analyzer_get_num_commands(analyzer uintptr) (uint64, error) {
	return tryAgxps_aps_timing_analyzer_get_num_commands(analyzer)
}

var _agxps_aps_timing_analyzer_get_num_work_cliques func(analyzer uintptr) uint64
var _agxps_aps_timing_analyzer_get_num_work_cliquesErr error

func tryAgxps_aps_timing_analyzer_get_num_work_cliques(analyzer uintptr) (uint64, error) {
	if _agxps_aps_timing_analyzer_get_num_work_cliques == nil {
		return 0, symbolCallError("agxps_aps_timing_analyzer_get_num_work_cliques", "", _agxps_aps_timing_analyzer_get_num_work_cliquesErr)
	}
	return _agxps_aps_timing_analyzer_get_num_work_cliques(analyzer), nil
}

// Agxps_aps_timing_analyzer_get_num_work_cliques.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_num_work_cliques
func Agxps_aps_timing_analyzer_get_num_work_cliques(analyzer uintptr) (uint64, error) {
	return tryAgxps_aps_timing_analyzer_get_num_work_cliques(analyzer)
}

var _agxps_aps_timing_analyzer_get_work_cliques_average_duration func(analyzer uintptr) float64
var _agxps_aps_timing_analyzer_get_work_cliques_average_durationErr error

func tryAgxps_aps_timing_analyzer_get_work_cliques_average_duration(analyzer uintptr) (float64, error) {
	if _agxps_aps_timing_analyzer_get_work_cliques_average_duration == nil {
		return 0.0, symbolCallError("agxps_aps_timing_analyzer_get_work_cliques_average_duration", "", _agxps_aps_timing_analyzer_get_work_cliques_average_durationErr)
	}
	return _agxps_aps_timing_analyzer_get_work_cliques_average_duration(analyzer), nil
}

// Agxps_aps_timing_analyzer_get_work_cliques_average_duration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_average_duration
func Agxps_aps_timing_analyzer_get_work_cliques_average_duration(analyzer uintptr) (float64, error) {
	return tryAgxps_aps_timing_analyzer_get_work_cliques_average_duration(analyzer)
}

var _agxps_aps_timing_analyzer_get_work_cliques_max_duration func(analyzer uintptr) float64
var _agxps_aps_timing_analyzer_get_work_cliques_max_durationErr error

func tryAgxps_aps_timing_analyzer_get_work_cliques_max_duration(analyzer uintptr) (float64, error) {
	if _agxps_aps_timing_analyzer_get_work_cliques_max_duration == nil {
		return 0.0, symbolCallError("agxps_aps_timing_analyzer_get_work_cliques_max_duration", "", _agxps_aps_timing_analyzer_get_work_cliques_max_durationErr)
	}
	return _agxps_aps_timing_analyzer_get_work_cliques_max_duration(analyzer), nil
}

// Agxps_aps_timing_analyzer_get_work_cliques_max_duration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_max_duration
func Agxps_aps_timing_analyzer_get_work_cliques_max_duration(analyzer uintptr) (float64, error) {
	return tryAgxps_aps_timing_analyzer_get_work_cliques_max_duration(analyzer)
}

var _agxps_aps_timing_analyzer_get_work_cliques_min_duration func(analyzer uintptr) float64
var _agxps_aps_timing_analyzer_get_work_cliques_min_durationErr error

func tryAgxps_aps_timing_analyzer_get_work_cliques_min_duration(analyzer uintptr) (float64, error) {
	if _agxps_aps_timing_analyzer_get_work_cliques_min_duration == nil {
		return 0.0, symbolCallError("agxps_aps_timing_analyzer_get_work_cliques_min_duration", "", _agxps_aps_timing_analyzer_get_work_cliques_min_durationErr)
	}
	return _agxps_aps_timing_analyzer_get_work_cliques_min_duration(analyzer), nil
}

// Agxps_aps_timing_analyzer_get_work_cliques_min_duration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_min_duration
func Agxps_aps_timing_analyzer_get_work_cliques_min_duration(analyzer uintptr) (float64, error) {
	return tryAgxps_aps_timing_analyzer_get_work_cliques_min_duration(analyzer)
}

var _agxps_gpu_create func(gen uint, variant uint, rev uint) AGXPSGPU
var _agxps_gpu_createErr error

func tryAgxps_gpu_create(gen uint, variant uint, rev uint) (AGXPSGPU, error) {
	if _agxps_gpu_create == nil {
		return *new(AGXPSGPU), symbolCallError("agxps_gpu_create", "", _agxps_gpu_createErr)
	}
	return _agxps_gpu_create(gen, variant, rev), nil
}

// Agxps_gpu_create.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_create
func Agxps_gpu_create(gen uint, variant uint, rev uint) (AGXPSGPU, error) {
	return tryAgxps_gpu_create(gen, variant, rev)
}

var _agxps_gpu_destroy func(gpu AGXPSGPU)
var _agxps_gpu_destroyErr error

func tryAgxps_gpu_destroy(gpu AGXPSGPU) error {
	if _agxps_gpu_destroy == nil {
		return symbolCallError("agxps_gpu_destroy", "", _agxps_gpu_destroyErr)
	}
	_agxps_gpu_destroy(gpu)
	return nil
}

// Agxps_gpu_destroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_destroy
func Agxps_gpu_destroy(gpu AGXPSGPU) error {
	return tryAgxps_gpu_destroy(gpu)
}

var _agxps_gpu_format_name func(gpu AGXPSGPU, buf *byte, size uint64) int
var _agxps_gpu_format_nameErr error

func tryAgxps_gpu_format_name(gpu AGXPSGPU, buf *byte, size uint64) (int, error) {
	if _agxps_gpu_format_name == nil {
		return 0, symbolCallError("agxps_gpu_format_name", "", _agxps_gpu_format_nameErr)
	}
	return _agxps_gpu_format_name(gpu, buf, size), nil
}

// Agxps_gpu_format_name.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_format_name
func Agxps_gpu_format_name(gpu AGXPSGPU, buf *byte, size uint64) (int, error) {
	return tryAgxps_gpu_format_name(gpu, buf, size)
}

var _agxps_gpu_get_gen func(gpu AGXPSGPU) uint
var _agxps_gpu_get_genErr error

func tryAgxps_gpu_get_gen(gpu AGXPSGPU) (uint, error) {
	if _agxps_gpu_get_gen == nil {
		return 0, symbolCallError("agxps_gpu_get_gen", "", _agxps_gpu_get_genErr)
	}
	return _agxps_gpu_get_gen(gpu), nil
}

// Agxps_gpu_get_gen.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_gen
func Agxps_gpu_get_gen(gpu AGXPSGPU) (uint, error) {
	return tryAgxps_gpu_get_gen(gpu)
}

var _agxps_gpu_get_rev func(gpu AGXPSGPU) uint
var _agxps_gpu_get_revErr error

func tryAgxps_gpu_get_rev(gpu AGXPSGPU) (uint, error) {
	if _agxps_gpu_get_rev == nil {
		return 0, symbolCallError("agxps_gpu_get_rev", "", _agxps_gpu_get_revErr)
	}
	return _agxps_gpu_get_rev(gpu), nil
}

// Agxps_gpu_get_rev.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_rev
func Agxps_gpu_get_rev(gpu AGXPSGPU) (uint, error) {
	return tryAgxps_gpu_get_rev(gpu)
}

var _agxps_gpu_get_variant func(gpu AGXPSGPU) uint
var _agxps_gpu_get_variantErr error

func tryAgxps_gpu_get_variant(gpu AGXPSGPU) (uint, error) {
	if _agxps_gpu_get_variant == nil {
		return 0, symbolCallError("agxps_gpu_get_variant", "", _agxps_gpu_get_variantErr)
	}
	return _agxps_gpu_get_variant(gpu), nil
}

// Agxps_gpu_get_variant.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_variant
func Agxps_gpu_get_variant(gpu AGXPSGPU) (uint, error) {
	return tryAgxps_gpu_get_variant(gpu)
}

var _agxps_gpu_is_valid func(gpu AGXPSGPU) bool
var _agxps_gpu_is_validErr error

func tryAgxps_gpu_is_valid(gpu AGXPSGPU) (bool, error) {
	if _agxps_gpu_is_valid == nil {
		return false, symbolCallError("agxps_gpu_is_valid", "", _agxps_gpu_is_validErr)
	}
	return _agxps_gpu_is_valid(gpu), nil
}

// Agxps_gpu_is_valid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_is_valid
func Agxps_gpu_is_valid(gpu AGXPSGPU) (bool, error) {
	return tryAgxps_gpu_is_valid(gpu)
}

var _agxps_initialize func() int
var _agxps_initializeErr error

func tryAgxps_initialize() (int, error) {
	if _agxps_initialize == nil {
		return 0, symbolCallError("agxps_initialize", "", _agxps_initializeErr)
	}
	return _agxps_initialize(), nil
}

// Agxps_initialize.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_initialize
func Agxps_initialize() (int, error) {
	return tryAgxps_initialize()
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_agxps_aps_clique_instruction_trace_get_execution_events, &_agxps_aps_clique_instruction_trace_get_execution_eventsErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_execution_events", "")
	registerFunc(&_agxps_aps_clique_instruction_trace_get_execution_events_num, &_agxps_aps_clique_instruction_trace_get_execution_events_numErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_execution_events_num", "")
	registerFunc(&_agxps_aps_clique_instruction_trace_get_instruction_stats, &_agxps_aps_clique_instruction_trace_get_instruction_statsErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_instruction_stats", "")
	registerFunc(&_agxps_aps_clique_instruction_trace_get_pc_advances, &_agxps_aps_clique_instruction_trace_get_pc_advancesErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_pc_advances", "")
	registerFunc(&_agxps_aps_clique_instruction_trace_get_pc_advances_num, &_agxps_aps_clique_instruction_trace_get_pc_advances_numErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_pc_advances_num", "")
	registerFunc(&_agxps_aps_clique_instruction_trace_get_timestamp_references, &_agxps_aps_clique_instruction_trace_get_timestamp_referencesErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_timestamp_references", "")
	registerFunc(&_agxps_aps_clique_instruction_trace_get_timestamp_references_num, &_agxps_aps_clique_instruction_trace_get_timestamp_references_numErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_timestamp_references_num", "")
	registerFunc(&_agxps_aps_clique_time_stats_create, &_agxps_aps_clique_time_stats_createErr, frameworkHandle, "agxps_aps_clique_time_stats_create", "")
	registerFunc(&_agxps_aps_descriptor_create, &_agxps_aps_descriptor_createErr, frameworkHandle, "agxps_aps_descriptor_create", "")
	registerFunc(&_agxps_aps_gpu_is_supported, &_agxps_aps_gpu_is_supportedErr, frameworkHandle, "agxps_aps_gpu_is_supported", "")
	registerFunc(&_agxps_aps_parser_create, &_agxps_aps_parser_createErr, frameworkHandle, "agxps_aps_parser_create", "")
	registerFunc(&_agxps_aps_parser_destroy, &_agxps_aps_parser_destroyErr, frameworkHandle, "agxps_aps_parser_destroy", "")
	registerFunc(&_agxps_aps_parser_is_valid, &_agxps_aps_parser_is_validErr, frameworkHandle, "agxps_aps_parser_is_valid", "")
	registerFunc(&_agxps_aps_parser_parse, &_agxps_aps_parser_parseErr, frameworkHandle, "agxps_aps_parser_parse", "")
	registerFunc(&_agxps_aps_profile_data_destroy, &_agxps_aps_profile_data_destroyErr, frameworkHandle, "agxps_aps_profile_data_destroy", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_clique_id, &_agxps_aps_profile_data_get_esl_clique_clique_idErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_clique_id", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_end, &_agxps_aps_profile_data_get_esl_clique_endErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_end", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_esl_id, &_agxps_aps_profile_data_get_esl_clique_esl_idErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_esl_id", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_instruction_trace, &_agxps_aps_profile_data_get_esl_clique_instruction_traceErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_instruction_trace", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_kick_id, &_agxps_aps_profile_data_get_esl_clique_kick_idErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_kick_id", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_missing_end, &_agxps_aps_profile_data_get_esl_clique_missing_endErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_missing_end", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_clique_start, &_agxps_aps_profile_data_get_esl_clique_startErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_start", "")
	registerFunc(&_agxps_aps_profile_data_get_esl_cliques_num, &_agxps_aps_profile_data_get_esl_cliques_numErr, frameworkHandle, "agxps_aps_profile_data_get_esl_cliques_num", "")
	registerFunc(&_agxps_aps_profile_data_get_kick_end, &_agxps_aps_profile_data_get_kick_endErr, frameworkHandle, "agxps_aps_profile_data_get_kick_end", "")
	registerFunc(&_agxps_aps_profile_data_get_kick_id, &_agxps_aps_profile_data_get_kick_idErr, frameworkHandle, "agxps_aps_profile_data_get_kick_id", "")
	registerFunc(&_agxps_aps_profile_data_get_kick_start, &_agxps_aps_profile_data_get_kick_startErr, frameworkHandle, "agxps_aps_profile_data_get_kick_start", "")
	registerFunc(&_agxps_aps_profile_data_get_kicks_num, &_agxps_aps_profile_data_get_kicks_numErr, frameworkHandle, "agxps_aps_profile_data_get_kicks_num", "")
	registerFunc(&_agxps_aps_profile_data_get_work_clique_end, &_agxps_aps_profile_data_get_work_clique_endErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_end", "")
	registerFunc(&_agxps_aps_profile_data_get_work_clique_instruction_trace, &_agxps_aps_profile_data_get_work_clique_instruction_traceErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_instruction_trace", "")
	registerFunc(&_agxps_aps_profile_data_get_work_clique_start, &_agxps_aps_profile_data_get_work_clique_startErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_start", "")
	registerFunc(&_agxps_aps_profile_data_get_work_cliques_num, &_agxps_aps_profile_data_get_work_cliques_numErr, frameworkHandle, "agxps_aps_profile_data_get_work_cliques_num", "")
	registerFunc(&_agxps_aps_profile_data_is_valid, &_agxps_aps_profile_data_is_validErr, frameworkHandle, "agxps_aps_profile_data_is_valid", "")
	registerFunc(&_agxps_aps_timing_analyzer_get_num_commands, &_agxps_aps_timing_analyzer_get_num_commandsErr, frameworkHandle, "agxps_aps_timing_analyzer_get_num_commands", "")
	registerFunc(&_agxps_aps_timing_analyzer_get_num_work_cliques, &_agxps_aps_timing_analyzer_get_num_work_cliquesErr, frameworkHandle, "agxps_aps_timing_analyzer_get_num_work_cliques", "")
	registerFunc(&_agxps_aps_timing_analyzer_get_work_cliques_average_duration, &_agxps_aps_timing_analyzer_get_work_cliques_average_durationErr, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_average_duration", "")
	registerFunc(&_agxps_aps_timing_analyzer_get_work_cliques_max_duration, &_agxps_aps_timing_analyzer_get_work_cliques_max_durationErr, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_max_duration", "")
	registerFunc(&_agxps_aps_timing_analyzer_get_work_cliques_min_duration, &_agxps_aps_timing_analyzer_get_work_cliques_min_durationErr, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_min_duration", "")
	registerFunc(&_agxps_gpu_create, &_agxps_gpu_createErr, frameworkHandle, "agxps_gpu_create", "")
	registerFunc(&_agxps_gpu_destroy, &_agxps_gpu_destroyErr, frameworkHandle, "agxps_gpu_destroy", "")
	registerFunc(&_agxps_gpu_format_name, &_agxps_gpu_format_nameErr, frameworkHandle, "agxps_gpu_format_name", "")
	registerFunc(&_agxps_gpu_get_gen, &_agxps_gpu_get_genErr, frameworkHandle, "agxps_gpu_get_gen", "")
	registerFunc(&_agxps_gpu_get_rev, &_agxps_gpu_get_revErr, frameworkHandle, "agxps_gpu_get_rev", "")
	registerFunc(&_agxps_gpu_get_variant, &_agxps_gpu_get_variantErr, frameworkHandle, "agxps_gpu_get_variant", "")
	registerFunc(&_agxps_gpu_is_valid, &_agxps_gpu_is_validErr, frameworkHandle, "agxps_gpu_is_valid", "")
	registerFunc(&_agxps_initialize, &_agxps_initializeErr, frameworkHandle, "agxps_initialize", "")
}
