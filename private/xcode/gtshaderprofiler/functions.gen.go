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

var _agxpsApsCliqueInstructionTraceGetExecutionEvents func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxpsApsCliqueInstructionTraceGetExecutionEventsErr error

func tryAgxpsApsCliqueInstructionTraceGetExecutionEvents(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxpsApsCliqueInstructionTraceGetExecutionEvents == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_execution_events", "", _agxpsApsCliqueInstructionTraceGetExecutionEventsErr)
	}
	return _agxpsApsCliqueInstructionTraceGetExecutionEvents(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetExecutionEvents.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_execution_events
func AgxpsApsCliqueInstructionTraceGetExecutionEvents(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxpsApsCliqueInstructionTraceGetExecutionEvents(trace)
}

var _agxpsApsCliqueInstructionTraceGetExecutionEventsNum func(trace AGXPSCliqueInstructionTraceRef) uint64
var _agxpsApsCliqueInstructionTraceGetExecutionEventsNumErr error

func tryAgxpsApsCliqueInstructionTraceGetExecutionEventsNum(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	if _agxpsApsCliqueInstructionTraceGetExecutionEventsNum == nil {
		return 0, symbolCallError("agxps_aps_clique_instruction_trace_get_execution_events_num", "", _agxpsApsCliqueInstructionTraceGetExecutionEventsNumErr)
	}
	return _agxpsApsCliqueInstructionTraceGetExecutionEventsNum(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetExecutionEventsNum.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_execution_events_num
func AgxpsApsCliqueInstructionTraceGetExecutionEventsNum(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	return tryAgxpsApsCliqueInstructionTraceGetExecutionEventsNum(trace)
}

var _agxpsApsCliqueInstructionTraceGetInstructionStats func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxpsApsCliqueInstructionTraceGetInstructionStatsErr error

func tryAgxpsApsCliqueInstructionTraceGetInstructionStats(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxpsApsCliqueInstructionTraceGetInstructionStats == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_instruction_stats", "", _agxpsApsCliqueInstructionTraceGetInstructionStatsErr)
	}
	return _agxpsApsCliqueInstructionTraceGetInstructionStats(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetInstructionStats.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_instruction_stats
func AgxpsApsCliqueInstructionTraceGetInstructionStats(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxpsApsCliqueInstructionTraceGetInstructionStats(trace)
}

var _agxpsApsCliqueInstructionTraceGetPCAdvances func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxpsApsCliqueInstructionTraceGetPCAdvancesErr error

func tryAgxpsApsCliqueInstructionTraceGetPCAdvances(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxpsApsCliqueInstructionTraceGetPCAdvances == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_pc_advances", "", _agxpsApsCliqueInstructionTraceGetPCAdvancesErr)
	}
	return _agxpsApsCliqueInstructionTraceGetPCAdvances(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetPCAdvances.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_pc_advances
func AgxpsApsCliqueInstructionTraceGetPCAdvances(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxpsApsCliqueInstructionTraceGetPCAdvances(trace)
}

var _agxpsApsCliqueInstructionTraceGetPCAdvancesNum func(trace AGXPSCliqueInstructionTraceRef) uint64
var _agxpsApsCliqueInstructionTraceGetPCAdvancesNumErr error

func tryAgxpsApsCliqueInstructionTraceGetPCAdvancesNum(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	if _agxpsApsCliqueInstructionTraceGetPCAdvancesNum == nil {
		return 0, symbolCallError("agxps_aps_clique_instruction_trace_get_pc_advances_num", "", _agxpsApsCliqueInstructionTraceGetPCAdvancesNumErr)
	}
	return _agxpsApsCliqueInstructionTraceGetPCAdvancesNum(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetPCAdvancesNum.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_pc_advances_num
func AgxpsApsCliqueInstructionTraceGetPCAdvancesNum(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	return tryAgxpsApsCliqueInstructionTraceGetPCAdvancesNum(trace)
}

var _agxpsApsCliqueInstructionTraceGetTimestampReferences func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxpsApsCliqueInstructionTraceGetTimestampReferencesErr error

func tryAgxpsApsCliqueInstructionTraceGetTimestampReferences(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxpsApsCliqueInstructionTraceGetTimestampReferences == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_timestamp_references", "", _agxpsApsCliqueInstructionTraceGetTimestampReferencesErr)
	}
	return _agxpsApsCliqueInstructionTraceGetTimestampReferences(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetTimestampReferences.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_timestamp_references
func AgxpsApsCliqueInstructionTraceGetTimestampReferences(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	return tryAgxpsApsCliqueInstructionTraceGetTimestampReferences(trace)
}

var _agxpsApsCliqueInstructionTraceGetTimestampReferencesNum func(trace AGXPSCliqueInstructionTraceRef) uint64
var _agxpsApsCliqueInstructionTraceGetTimestampReferencesNumErr error

func tryAgxpsApsCliqueInstructionTraceGetTimestampReferencesNum(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	if _agxpsApsCliqueInstructionTraceGetTimestampReferencesNum == nil {
		return 0, symbolCallError("agxps_aps_clique_instruction_trace_get_timestamp_references_num", "", _agxpsApsCliqueInstructionTraceGetTimestampReferencesNumErr)
	}
	return _agxpsApsCliqueInstructionTraceGetTimestampReferencesNum(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetTimestampReferencesNum.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_instruction_trace_get_timestamp_references_num
func AgxpsApsCliqueInstructionTraceGetTimestampReferencesNum(trace AGXPSCliqueInstructionTraceRef) (uint64, error) {
	return tryAgxpsApsCliqueInstructionTraceGetTimestampReferencesNum(trace)
}

var _agxpsApsCliqueTimeStatsCreate func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueTimeStatsRef
var _agxpsApsCliqueTimeStatsCreateErr error

func tryAgxpsApsCliqueTimeStatsCreate(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueTimeStatsRef, error) {
	if _agxpsApsCliqueTimeStatsCreate == nil {
		return *new(AGXPSCliqueTimeStatsRef), symbolCallError("agxps_aps_clique_time_stats_create", "", _agxpsApsCliqueTimeStatsCreateErr)
	}
	return _agxpsApsCliqueTimeStatsCreate(profileData, cliqueIndex), nil
}

// AgxpsApsCliqueTimeStatsCreate.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_clique_time_stats_create
func AgxpsApsCliqueTimeStatsCreate(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueTimeStatsRef, error) {
	return tryAgxpsApsCliqueTimeStatsCreate(profileData, cliqueIndex)
}

var _agxpsApsDescriptorCreate func(descriptor unsafe.Pointer) AGXPSDescriptorRef
var _agxpsApsDescriptorCreateErr error

func tryAgxpsApsDescriptorCreate(descriptor unsafe.Pointer) (AGXPSDescriptorRef, error) {
	if _agxpsApsDescriptorCreate == nil {
		return *new(AGXPSDescriptorRef), symbolCallError("agxps_aps_descriptor_create", "", _agxpsApsDescriptorCreateErr)
	}
	return _agxpsApsDescriptorCreate(descriptor), nil
}

// AgxpsApsDescriptorCreate.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_descriptor_create
func AgxpsApsDescriptorCreate(descriptor unsafe.Pointer) (AGXPSDescriptorRef, error) {
	return tryAgxpsApsDescriptorCreate(descriptor)
}

var _agxpsApsGPUIsSupported func(gpu AGXPSGPU) bool
var _agxpsApsGPUIsSupportedErr error

func tryAgxpsApsGPUIsSupported(gpu AGXPSGPU) (bool, error) {
	if _agxpsApsGPUIsSupported == nil {
		return false, symbolCallError("agxps_aps_gpu_is_supported", "", _agxpsApsGPUIsSupportedErr)
	}
	return _agxpsApsGPUIsSupported(gpu), nil
}

// AgxpsApsGPUIsSupported.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_gpu_is_supported
func AgxpsApsGPUIsSupported(gpu AGXPSGPU) (bool, error) {
	return tryAgxpsApsGPUIsSupported(gpu)
}

var _agxpsApsParserCreate func(descriptor AGXPSDescriptorRef) AGXPSParserHandle
var _agxpsApsParserCreateErr error

func tryAgxpsApsParserCreate(descriptor AGXPSDescriptorRef) (AGXPSParserHandle, error) {
	if _agxpsApsParserCreate == nil {
		return *new(AGXPSParserHandle), symbolCallError("agxps_aps_parser_create", "", _agxpsApsParserCreateErr)
	}
	return _agxpsApsParserCreate(descriptor), nil
}

// AgxpsApsParserCreate.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_create
func AgxpsApsParserCreate(descriptor AGXPSDescriptorRef) (AGXPSParserHandle, error) {
	return tryAgxpsApsParserCreate(descriptor)
}

var _agxpsApsParserDestroy func(parser AGXPSParserHandle)
var _agxpsApsParserDestroyErr error

func tryAgxpsApsParserDestroy(parser AGXPSParserHandle) error {
	if _agxpsApsParserDestroy == nil {
		return symbolCallError("agxps_aps_parser_destroy", "", _agxpsApsParserDestroyErr)
	}
	_agxpsApsParserDestroy(parser)
	return nil
}

// AgxpsApsParserDestroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_destroy
func AgxpsApsParserDestroy(parser AGXPSParserHandle) error {
	return tryAgxpsApsParserDestroy(parser)
}

var _agxpsApsParserIsValid func(parser AGXPSParserHandle) bool
var _agxpsApsParserIsValidErr error

func tryAgxpsApsParserIsValid(parser AGXPSParserHandle) (bool, error) {
	if _agxpsApsParserIsValid == nil {
		return false, symbolCallError("agxps_aps_parser_is_valid", "", _agxpsApsParserIsValidErr)
	}
	return _agxpsApsParserIsValid(parser), nil
}

// AgxpsApsParserIsValid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_is_valid
func AgxpsApsParserIsValid(parser AGXPSParserHandle) (bool, error) {
	return tryAgxpsApsParserIsValid(parser)
}

var _agxpsApsParserParse func(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) int
var _agxpsApsParserParseErr error

func tryAgxpsApsParserParse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) (int, error) {
	if _agxpsApsParserParse == nil {
		return 0, symbolCallError("agxps_aps_parser_parse", "", _agxpsApsParserParseErr)
	}
	return _agxpsApsParserParse(parser, data, size, profileDataOut), nil
}

// AgxpsApsParserParse.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_parser_parse
func AgxpsApsParserParse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) (int, error) {
	return tryAgxpsApsParserParse(parser, data, size, profileDataOut)
}

var _agxpsApsProfileDataDestroy func(profileData AGXPSProfileData)
var _agxpsApsProfileDataDestroyErr error

func tryAgxpsApsProfileDataDestroy(profileData AGXPSProfileData) error {
	if _agxpsApsProfileDataDestroy == nil {
		return symbolCallError("agxps_aps_profile_data_destroy", "", _agxpsApsProfileDataDestroyErr)
	}
	_agxpsApsProfileDataDestroy(profileData)
	return nil
}

// AgxpsApsProfileDataDestroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_destroy
func AgxpsApsProfileDataDestroy(profileData AGXPSProfileData) error {
	return tryAgxpsApsProfileDataDestroy(profileData)
}

var _agxpsApsProfileDataGetEslCliqueCliqueID func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetEslCliqueCliqueIDErr error

func tryAgxpsApsProfileDataGetEslCliqueCliqueID(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetEslCliqueCliqueID == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_clique_id", "", _agxpsApsProfileDataGetEslCliqueCliqueIDErr)
	}
	return _agxpsApsProfileDataGetEslCliqueCliqueID(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueCliqueID.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_clique_id
func AgxpsApsProfileDataGetEslCliqueCliqueID(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliqueCliqueID(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliqueEnd func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetEslCliqueEndErr error

func tryAgxpsApsProfileDataGetEslCliqueEnd(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetEslCliqueEnd == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_end", "", _agxpsApsProfileDataGetEslCliqueEndErr)
	}
	return _agxpsApsProfileDataGetEslCliqueEnd(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueEnd.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_end
func AgxpsApsProfileDataGetEslCliqueEnd(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliqueEnd(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliqueEslID func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetEslCliqueEslIDErr error

func tryAgxpsApsProfileDataGetEslCliqueEslID(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetEslCliqueEslID == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_esl_id", "", _agxpsApsProfileDataGetEslCliqueEslIDErr)
	}
	return _agxpsApsProfileDataGetEslCliqueEslID(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueEslID.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_esl_id
func AgxpsApsProfileDataGetEslCliqueEslID(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliqueEslID(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliqueInstructionTrace func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef
var _agxpsApsProfileDataGetEslCliqueInstructionTraceErr error

func tryAgxpsApsProfileDataGetEslCliqueInstructionTrace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	if _agxpsApsProfileDataGetEslCliqueInstructionTrace == nil {
		return *new(AGXPSCliqueInstructionTraceRef), symbolCallError("agxps_aps_profile_data_get_esl_clique_instruction_trace", "", _agxpsApsProfileDataGetEslCliqueInstructionTraceErr)
	}
	return _agxpsApsProfileDataGetEslCliqueInstructionTrace(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueInstructionTrace.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_instruction_trace
func AgxpsApsProfileDataGetEslCliqueInstructionTrace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	return tryAgxpsApsProfileDataGetEslCliqueInstructionTrace(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliqueKickID func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetEslCliqueKickIDErr error

func tryAgxpsApsProfileDataGetEslCliqueKickID(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetEslCliqueKickID == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_kick_id", "", _agxpsApsProfileDataGetEslCliqueKickIDErr)
	}
	return _agxpsApsProfileDataGetEslCliqueKickID(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueKickID.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_kick_id
func AgxpsApsProfileDataGetEslCliqueKickID(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliqueKickID(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliqueMissingEnd func(profileData AGXPSProfileData, cliqueIndex uint64) bool
var _agxpsApsProfileDataGetEslCliqueMissingEndErr error

func tryAgxpsApsProfileDataGetEslCliqueMissingEnd(profileData AGXPSProfileData, cliqueIndex uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueMissingEnd == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_missing_end", "", _agxpsApsProfileDataGetEslCliqueMissingEndErr)
	}
	return _agxpsApsProfileDataGetEslCliqueMissingEnd(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueMissingEnd.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_missing_end
func AgxpsApsProfileDataGetEslCliqueMissingEnd(profileData AGXPSProfileData, cliqueIndex uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueMissingEnd(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliqueStart func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetEslCliqueStartErr error

func tryAgxpsApsProfileDataGetEslCliqueStart(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetEslCliqueStart == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_clique_start", "", _agxpsApsProfileDataGetEslCliqueStartErr)
	}
	return _agxpsApsProfileDataGetEslCliqueStart(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetEslCliqueStart.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_clique_start
func AgxpsApsProfileDataGetEslCliqueStart(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliqueStart(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetEslCliquesNum func(profileData AGXPSProfileData) uint64
var _agxpsApsProfileDataGetEslCliquesNumErr error

func tryAgxpsApsProfileDataGetEslCliquesNum(profileData AGXPSProfileData) (uint64, error) {
	if _agxpsApsProfileDataGetEslCliquesNum == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_esl_cliques_num", "", _agxpsApsProfileDataGetEslCliquesNumErr)
	}
	return _agxpsApsProfileDataGetEslCliquesNum(profileData), nil
}

// AgxpsApsProfileDataGetEslCliquesNum.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_esl_cliques_num
func AgxpsApsProfileDataGetEslCliquesNum(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliquesNum(profileData)
}

var _agxpsApsProfileDataGetKickEnd func(profileData AGXPSProfileData, kickIndex uint64) uint64
var _agxpsApsProfileDataGetKickEndErr error

func tryAgxpsApsProfileDataGetKickEnd(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetKickEnd == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kick_end", "", _agxpsApsProfileDataGetKickEndErr)
	}
	return _agxpsApsProfileDataGetKickEnd(profileData, kickIndex), nil
}

// AgxpsApsProfileDataGetKickEnd.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_end
func AgxpsApsProfileDataGetKickEnd(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetKickEnd(profileData, kickIndex)
}

var _agxpsApsProfileDataGetKickID func(profileData AGXPSProfileData, kickIndex uint64) uint64
var _agxpsApsProfileDataGetKickIDErr error

func tryAgxpsApsProfileDataGetKickID(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetKickID == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kick_id", "", _agxpsApsProfileDataGetKickIDErr)
	}
	return _agxpsApsProfileDataGetKickID(profileData, kickIndex), nil
}

// AgxpsApsProfileDataGetKickID.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_id
func AgxpsApsProfileDataGetKickID(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetKickID(profileData, kickIndex)
}

var _agxpsApsProfileDataGetKickStart func(profileData AGXPSProfileData, kickIndex uint64) uint64
var _agxpsApsProfileDataGetKickStartErr error

func tryAgxpsApsProfileDataGetKickStart(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetKickStart == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kick_start", "", _agxpsApsProfileDataGetKickStartErr)
	}
	return _agxpsApsProfileDataGetKickStart(profileData, kickIndex), nil
}

// AgxpsApsProfileDataGetKickStart.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kick_start
func AgxpsApsProfileDataGetKickStart(profileData AGXPSProfileData, kickIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetKickStart(profileData, kickIndex)
}

var _agxpsApsProfileDataGetKicksNum func(profileData AGXPSProfileData) uint64
var _agxpsApsProfileDataGetKicksNumErr error

func tryAgxpsApsProfileDataGetKicksNum(profileData AGXPSProfileData) (uint64, error) {
	if _agxpsApsProfileDataGetKicksNum == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_kicks_num", "", _agxpsApsProfileDataGetKicksNumErr)
	}
	return _agxpsApsProfileDataGetKicksNum(profileData), nil
}

// AgxpsApsProfileDataGetKicksNum.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_kicks_num
func AgxpsApsProfileDataGetKicksNum(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxpsApsProfileDataGetKicksNum(profileData)
}

var _agxpsApsProfileDataGetWorkCliqueEnd func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetWorkCliqueEndErr error

func tryAgxpsApsProfileDataGetWorkCliqueEnd(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetWorkCliqueEnd == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_clique_end", "", _agxpsApsProfileDataGetWorkCliqueEndErr)
	}
	return _agxpsApsProfileDataGetWorkCliqueEnd(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetWorkCliqueEnd.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_end
func AgxpsApsProfileDataGetWorkCliqueEnd(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetWorkCliqueEnd(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetWorkCliqueInstructionTrace func(profileData AGXPSProfileData, cliqueIndex uint64) AGXPSCliqueInstructionTraceRef
var _agxpsApsProfileDataGetWorkCliqueInstructionTraceErr error

func tryAgxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	if _agxpsApsProfileDataGetWorkCliqueInstructionTrace == nil {
		return *new(AGXPSCliqueInstructionTraceRef), symbolCallError("agxps_aps_profile_data_get_work_clique_instruction_trace", "", _agxpsApsProfileDataGetWorkCliqueInstructionTraceErr)
	}
	return _agxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetWorkCliqueInstructionTrace.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_instruction_trace
func AgxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueInstructionTraceRef, error) {
	return tryAgxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetWorkCliqueStart func(profileData AGXPSProfileData, cliqueIndex uint64) uint64
var _agxpsApsProfileDataGetWorkCliqueStartErr error

func tryAgxpsApsProfileDataGetWorkCliqueStart(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	if _agxpsApsProfileDataGetWorkCliqueStart == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_clique_start", "", _agxpsApsProfileDataGetWorkCliqueStartErr)
	}
	return _agxpsApsProfileDataGetWorkCliqueStart(profileData, cliqueIndex), nil
}

// AgxpsApsProfileDataGetWorkCliqueStart.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_clique_start
func AgxpsApsProfileDataGetWorkCliqueStart(profileData AGXPSProfileData, cliqueIndex uint64) (uint64, error) {
	return tryAgxpsApsProfileDataGetWorkCliqueStart(profileData, cliqueIndex)
}

var _agxpsApsProfileDataGetWorkCliquesNum func(profileData AGXPSProfileData) uint64
var _agxpsApsProfileDataGetWorkCliquesNumErr error

func tryAgxpsApsProfileDataGetWorkCliquesNum(profileData AGXPSProfileData) (uint64, error) {
	if _agxpsApsProfileDataGetWorkCliquesNum == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_work_cliques_num", "", _agxpsApsProfileDataGetWorkCliquesNumErr)
	}
	return _agxpsApsProfileDataGetWorkCliquesNum(profileData), nil
}

// AgxpsApsProfileDataGetWorkCliquesNum.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_get_work_cliques_num
func AgxpsApsProfileDataGetWorkCliquesNum(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxpsApsProfileDataGetWorkCliquesNum(profileData)
}

var _agxpsApsProfileDataIsValid func(profileData AGXPSProfileData) bool
var _agxpsApsProfileDataIsValidErr error

func tryAgxpsApsProfileDataIsValid(profileData AGXPSProfileData) (bool, error) {
	if _agxpsApsProfileDataIsValid == nil {
		return false, symbolCallError("agxps_aps_profile_data_is_valid", "", _agxpsApsProfileDataIsValidErr)
	}
	return _agxpsApsProfileDataIsValid(profileData), nil
}

// AgxpsApsProfileDataIsValid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_profile_data_is_valid
func AgxpsApsProfileDataIsValid(profileData AGXPSProfileData) (bool, error) {
	return tryAgxpsApsProfileDataIsValid(profileData)
}

var _agxpsApsTimingAnalyzerGetNumCommands func(analyzer uintptr) uint64
var _agxpsApsTimingAnalyzerGetNumCommandsErr error

func tryAgxpsApsTimingAnalyzerGetNumCommands(analyzer uintptr) (uint64, error) {
	if _agxpsApsTimingAnalyzerGetNumCommands == nil {
		return 0, symbolCallError("agxps_aps_timing_analyzer_get_num_commands", "", _agxpsApsTimingAnalyzerGetNumCommandsErr)
	}
	return _agxpsApsTimingAnalyzerGetNumCommands(analyzer), nil
}

// AgxpsApsTimingAnalyzerGetNumCommands.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_num_commands
func AgxpsApsTimingAnalyzerGetNumCommands(analyzer uintptr) (uint64, error) {
	return tryAgxpsApsTimingAnalyzerGetNumCommands(analyzer)
}

var _agxpsApsTimingAnalyzerGetNumWorkCliques func(analyzer uintptr) uint64
var _agxpsApsTimingAnalyzerGetNumWorkCliquesErr error

func tryAgxpsApsTimingAnalyzerGetNumWorkCliques(analyzer uintptr) (uint64, error) {
	if _agxpsApsTimingAnalyzerGetNumWorkCliques == nil {
		return 0, symbolCallError("agxps_aps_timing_analyzer_get_num_work_cliques", "", _agxpsApsTimingAnalyzerGetNumWorkCliquesErr)
	}
	return _agxpsApsTimingAnalyzerGetNumWorkCliques(analyzer), nil
}

// AgxpsApsTimingAnalyzerGetNumWorkCliques.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_num_work_cliques
func AgxpsApsTimingAnalyzerGetNumWorkCliques(analyzer uintptr) (uint64, error) {
	return tryAgxpsApsTimingAnalyzerGetNumWorkCliques(analyzer)
}

var _agxpsApsTimingAnalyzerGetWorkCliquesAverageDuration func(analyzer uintptr) float64
var _agxpsApsTimingAnalyzerGetWorkCliquesAverageDurationErr error

func tryAgxpsApsTimingAnalyzerGetWorkCliquesAverageDuration(analyzer uintptr) (float64, error) {
	if _agxpsApsTimingAnalyzerGetWorkCliquesAverageDuration == nil {
		return 0.0, symbolCallError("agxps_aps_timing_analyzer_get_work_cliques_average_duration", "", _agxpsApsTimingAnalyzerGetWorkCliquesAverageDurationErr)
	}
	return _agxpsApsTimingAnalyzerGetWorkCliquesAverageDuration(analyzer), nil
}

// AgxpsApsTimingAnalyzerGetWorkCliquesAverageDuration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_average_duration
func AgxpsApsTimingAnalyzerGetWorkCliquesAverageDuration(analyzer uintptr) (float64, error) {
	return tryAgxpsApsTimingAnalyzerGetWorkCliquesAverageDuration(analyzer)
}

var _agxpsApsTimingAnalyzerGetWorkCliquesMaxDuration func(analyzer uintptr) float64
var _agxpsApsTimingAnalyzerGetWorkCliquesMaxDurationErr error

func tryAgxpsApsTimingAnalyzerGetWorkCliquesMaxDuration(analyzer uintptr) (float64, error) {
	if _agxpsApsTimingAnalyzerGetWorkCliquesMaxDuration == nil {
		return 0.0, symbolCallError("agxps_aps_timing_analyzer_get_work_cliques_max_duration", "", _agxpsApsTimingAnalyzerGetWorkCliquesMaxDurationErr)
	}
	return _agxpsApsTimingAnalyzerGetWorkCliquesMaxDuration(analyzer), nil
}

// AgxpsApsTimingAnalyzerGetWorkCliquesMaxDuration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_max_duration
func AgxpsApsTimingAnalyzerGetWorkCliquesMaxDuration(analyzer uintptr) (float64, error) {
	return tryAgxpsApsTimingAnalyzerGetWorkCliquesMaxDuration(analyzer)
}

var _agxpsApsTimingAnalyzerGetWorkCliquesMinDuration func(analyzer uintptr) float64
var _agxpsApsTimingAnalyzerGetWorkCliquesMinDurationErr error

func tryAgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer uintptr) (float64, error) {
	if _agxpsApsTimingAnalyzerGetWorkCliquesMinDuration == nil {
		return 0.0, symbolCallError("agxps_aps_timing_analyzer_get_work_cliques_min_duration", "", _agxpsApsTimingAnalyzerGetWorkCliquesMinDurationErr)
	}
	return _agxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer), nil
}

// AgxpsApsTimingAnalyzerGetWorkCliquesMinDuration.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_aps_timing_analyzer_get_work_cliques_min_duration
func AgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer uintptr) (float64, error) {
	return tryAgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer)
}

var _agxpsGPUCreate func(gen uint, variant uint, rev uint) AGXPSGPU
var _agxpsGPUCreateErr error

func tryAgxpsGPUCreate(gen uint, variant uint, rev uint) (AGXPSGPU, error) {
	if _agxpsGPUCreate == nil {
		return *new(AGXPSGPU), symbolCallError("agxps_gpu_create", "", _agxpsGPUCreateErr)
	}
	return _agxpsGPUCreate(gen, variant, rev), nil
}

// AgxpsGPUCreate.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_create
func AgxpsGPUCreate(gen uint, variant uint, rev uint) (AGXPSGPU, error) {
	return tryAgxpsGPUCreate(gen, variant, rev)
}

var _agxpsGPUDestroy func(gpu AGXPSGPU)
var _agxpsGPUDestroyErr error

func tryAgxpsGPUDestroy(gpu AGXPSGPU) error {
	if _agxpsGPUDestroy == nil {
		return symbolCallError("agxps_gpu_destroy", "", _agxpsGPUDestroyErr)
	}
	_agxpsGPUDestroy(gpu)
	return nil
}

// AgxpsGPUDestroy.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_destroy
func AgxpsGPUDestroy(gpu AGXPSGPU) error {
	return tryAgxpsGPUDestroy(gpu)
}

var _agxpsGPUFormatName func(gpu AGXPSGPU, buf *byte, size uint64) int
var _agxpsGPUFormatNameErr error

func tryAgxpsGPUFormatName(gpu AGXPSGPU, buf *byte, size uint64) (int, error) {
	if _agxpsGPUFormatName == nil {
		return 0, symbolCallError("agxps_gpu_format_name", "", _agxpsGPUFormatNameErr)
	}
	return _agxpsGPUFormatName(gpu, buf, size), nil
}

// AgxpsGPUFormatName.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_format_name
func AgxpsGPUFormatName(gpu AGXPSGPU, buf *byte, size uint64) (int, error) {
	return tryAgxpsGPUFormatName(gpu, buf, size)
}

var _agxpsGPUGetGen func(gpu AGXPSGPU) uint
var _agxpsGPUGetGenErr error

func tryAgxpsGPUGetGen(gpu AGXPSGPU) (uint, error) {
	if _agxpsGPUGetGen == nil {
		return 0, symbolCallError("agxps_gpu_get_gen", "", _agxpsGPUGetGenErr)
	}
	return _agxpsGPUGetGen(gpu), nil
}

// AgxpsGPUGetGen.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_gen
func AgxpsGPUGetGen(gpu AGXPSGPU) (uint, error) {
	return tryAgxpsGPUGetGen(gpu)
}

var _agxpsGPUGetRev func(gpu AGXPSGPU) uint
var _agxpsGPUGetRevErr error

func tryAgxpsGPUGetRev(gpu AGXPSGPU) (uint, error) {
	if _agxpsGPUGetRev == nil {
		return 0, symbolCallError("agxps_gpu_get_rev", "", _agxpsGPUGetRevErr)
	}
	return _agxpsGPUGetRev(gpu), nil
}

// AgxpsGPUGetRev.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_rev
func AgxpsGPUGetRev(gpu AGXPSGPU) (uint, error) {
	return tryAgxpsGPUGetRev(gpu)
}

var _agxpsGPUGetVariant func(gpu AGXPSGPU) uint
var _agxpsGPUGetVariantErr error

func tryAgxpsGPUGetVariant(gpu AGXPSGPU) (uint, error) {
	if _agxpsGPUGetVariant == nil {
		return 0, symbolCallError("agxps_gpu_get_variant", "", _agxpsGPUGetVariantErr)
	}
	return _agxpsGPUGetVariant(gpu), nil
}

// AgxpsGPUGetVariant.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_get_variant
func AgxpsGPUGetVariant(gpu AGXPSGPU) (uint, error) {
	return tryAgxpsGPUGetVariant(gpu)
}

var _agxpsGPUIsValid func(gpu AGXPSGPU) bool
var _agxpsGPUIsValidErr error

func tryAgxpsGPUIsValid(gpu AGXPSGPU) (bool, error) {
	if _agxpsGPUIsValid == nil {
		return false, symbolCallError("agxps_gpu_is_valid", "", _agxpsGPUIsValidErr)
	}
	return _agxpsGPUIsValid(gpu), nil
}

// AgxpsGPUIsValid.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_gpu_is_valid
func AgxpsGPUIsValid(gpu AGXPSGPU) (bool, error) {
	return tryAgxpsGPUIsValid(gpu)
}

var _agxpsInitialize func() int
var _agxpsInitializeErr error

func tryAgxpsInitialize() (int, error) {
	if _agxpsInitialize == nil {
		return 0, symbolCallError("agxps_initialize", "", _agxpsInitializeErr)
	}
	return _agxpsInitialize(), nil
}

// AgxpsInitialize.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/agxps_initialize
func AgxpsInitialize() (int, error) {
	return tryAgxpsInitialize()
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_agxpsApsCliqueInstructionTraceGetExecutionEvents, &_agxpsApsCliqueInstructionTraceGetExecutionEventsErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_execution_events", "")
	registerFunc(&_agxpsApsCliqueInstructionTraceGetExecutionEventsNum, &_agxpsApsCliqueInstructionTraceGetExecutionEventsNumErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_execution_events_num", "")
	registerFunc(&_agxpsApsCliqueInstructionTraceGetInstructionStats, &_agxpsApsCliqueInstructionTraceGetInstructionStatsErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_instruction_stats", "")
	registerFunc(&_agxpsApsCliqueInstructionTraceGetPCAdvances, &_agxpsApsCliqueInstructionTraceGetPCAdvancesErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_pc_advances", "")
	registerFunc(&_agxpsApsCliqueInstructionTraceGetPCAdvancesNum, &_agxpsApsCliqueInstructionTraceGetPCAdvancesNumErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_pc_advances_num", "")
	registerFunc(&_agxpsApsCliqueInstructionTraceGetTimestampReferences, &_agxpsApsCliqueInstructionTraceGetTimestampReferencesErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_timestamp_references", "")
	registerFunc(&_agxpsApsCliqueInstructionTraceGetTimestampReferencesNum, &_agxpsApsCliqueInstructionTraceGetTimestampReferencesNumErr, frameworkHandle, "agxps_aps_clique_instruction_trace_get_timestamp_references_num", "")
	registerFunc(&_agxpsApsCliqueTimeStatsCreate, &_agxpsApsCliqueTimeStatsCreateErr, frameworkHandle, "agxps_aps_clique_time_stats_create", "")
	registerFunc(&_agxpsApsDescriptorCreate, &_agxpsApsDescriptorCreateErr, frameworkHandle, "agxps_aps_descriptor_create", "")
	registerFunc(&_agxpsApsGPUIsSupported, &_agxpsApsGPUIsSupportedErr, frameworkHandle, "agxps_aps_gpu_is_supported", "")
	registerFunc(&_agxpsApsParserCreate, &_agxpsApsParserCreateErr, frameworkHandle, "agxps_aps_parser_create", "")
	registerFunc(&_agxpsApsParserDestroy, &_agxpsApsParserDestroyErr, frameworkHandle, "agxps_aps_parser_destroy", "")
	registerFunc(&_agxpsApsParserIsValid, &_agxpsApsParserIsValidErr, frameworkHandle, "agxps_aps_parser_is_valid", "")
	registerFunc(&_agxpsApsParserParse, &_agxpsApsParserParseErr, frameworkHandle, "agxps_aps_parser_parse", "")
	registerFunc(&_agxpsApsProfileDataDestroy, &_agxpsApsProfileDataDestroyErr, frameworkHandle, "agxps_aps_profile_data_destroy", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueCliqueID, &_agxpsApsProfileDataGetEslCliqueCliqueIDErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_clique_id", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueEnd, &_agxpsApsProfileDataGetEslCliqueEndErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_end", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueEslID, &_agxpsApsProfileDataGetEslCliqueEslIDErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_esl_id", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueInstructionTrace, &_agxpsApsProfileDataGetEslCliqueInstructionTraceErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_instruction_trace", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueKickID, &_agxpsApsProfileDataGetEslCliqueKickIDErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_kick_id", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueMissingEnd, &_agxpsApsProfileDataGetEslCliqueMissingEndErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_missing_end", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliqueStart, &_agxpsApsProfileDataGetEslCliqueStartErr, frameworkHandle, "agxps_aps_profile_data_get_esl_clique_start", "")
	registerFunc(&_agxpsApsProfileDataGetEslCliquesNum, &_agxpsApsProfileDataGetEslCliquesNumErr, frameworkHandle, "agxps_aps_profile_data_get_esl_cliques_num", "")
	registerFunc(&_agxpsApsProfileDataGetKickEnd, &_agxpsApsProfileDataGetKickEndErr, frameworkHandle, "agxps_aps_profile_data_get_kick_end", "")
	registerFunc(&_agxpsApsProfileDataGetKickID, &_agxpsApsProfileDataGetKickIDErr, frameworkHandle, "agxps_aps_profile_data_get_kick_id", "")
	registerFunc(&_agxpsApsProfileDataGetKickStart, &_agxpsApsProfileDataGetKickStartErr, frameworkHandle, "agxps_aps_profile_data_get_kick_start", "")
	registerFunc(&_agxpsApsProfileDataGetKicksNum, &_agxpsApsProfileDataGetKicksNumErr, frameworkHandle, "agxps_aps_profile_data_get_kicks_num", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliqueEnd, &_agxpsApsProfileDataGetWorkCliqueEndErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_end", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliqueInstructionTrace, &_agxpsApsProfileDataGetWorkCliqueInstructionTraceErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_instruction_trace", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliqueStart, &_agxpsApsProfileDataGetWorkCliqueStartErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_start", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliquesNum, &_agxpsApsProfileDataGetWorkCliquesNumErr, frameworkHandle, "agxps_aps_profile_data_get_work_cliques_num", "")
	registerFunc(&_agxpsApsProfileDataIsValid, &_agxpsApsProfileDataIsValidErr, frameworkHandle, "agxps_aps_profile_data_is_valid", "")
	registerFunc(&_agxpsApsTimingAnalyzerGetNumCommands, &_agxpsApsTimingAnalyzerGetNumCommandsErr, frameworkHandle, "agxps_aps_timing_analyzer_get_num_commands", "")
	registerFunc(&_agxpsApsTimingAnalyzerGetNumWorkCliques, &_agxpsApsTimingAnalyzerGetNumWorkCliquesErr, frameworkHandle, "agxps_aps_timing_analyzer_get_num_work_cliques", "")
	registerFunc(&_agxpsApsTimingAnalyzerGetWorkCliquesAverageDuration, &_agxpsApsTimingAnalyzerGetWorkCliquesAverageDurationErr, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_average_duration", "")
	registerFunc(&_agxpsApsTimingAnalyzerGetWorkCliquesMaxDuration, &_agxpsApsTimingAnalyzerGetWorkCliquesMaxDurationErr, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_max_duration", "")
	registerFunc(&_agxpsApsTimingAnalyzerGetWorkCliquesMinDuration, &_agxpsApsTimingAnalyzerGetWorkCliquesMinDurationErr, frameworkHandle, "agxps_aps_timing_analyzer_get_work_cliques_min_duration", "")
	registerFunc(&_agxpsGPUCreate, &_agxpsGPUCreateErr, frameworkHandle, "agxps_gpu_create", "")
	registerFunc(&_agxpsGPUDestroy, &_agxpsGPUDestroyErr, frameworkHandle, "agxps_gpu_destroy", "")
	registerFunc(&_agxpsGPUFormatName, &_agxpsGPUFormatNameErr, frameworkHandle, "agxps_gpu_format_name", "")
	registerFunc(&_agxpsGPUGetGen, &_agxpsGPUGetGenErr, frameworkHandle, "agxps_gpu_get_gen", "")
	registerFunc(&_agxpsGPUGetRev, &_agxpsGPUGetRevErr, frameworkHandle, "agxps_gpu_get_rev", "")
	registerFunc(&_agxpsGPUGetVariant, &_agxpsGPUGetVariantErr, frameworkHandle, "agxps_gpu_get_variant", "")
	registerFunc(&_agxpsGPUIsValid, &_agxpsGPUIsValidErr, frameworkHandle, "agxps_gpu_is_valid", "")
	registerFunc(&_agxpsInitialize, &_agxpsInitializeErr, frameworkHandle, "agxps_initialize", "")
}
