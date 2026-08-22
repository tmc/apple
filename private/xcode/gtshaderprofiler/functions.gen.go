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
func AgxpsApsParserIsValid(parser AGXPSParserHandle) (bool, error) {
	return tryAgxpsApsParserIsValid(parser)
}

var _agxpsApsParserParse func(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) int32
var _agxpsApsParserParseErr error

func tryAgxpsApsParserParse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) (int32, error) {
	if _agxpsApsParserParse == nil {
		return 0, symbolCallError("agxps_aps_parser_parse", "", _agxpsApsParserParseErr)
	}
	return _agxpsApsParserParse(parser, data, size, profileDataOut), nil
}

// AgxpsApsParserParse.
func AgxpsApsParserParse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, profileDataOut *AGXPSProfileData) (int32, error) {
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
func AgxpsApsProfileDataDestroy(profileData AGXPSProfileData) error {
	return tryAgxpsApsProfileDataDestroy(profileData)
}

var _agxpsApsProfileDataGetCounterGroupID func(profileData AGXPSProfileData, out *byte, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterGroupIDErr error

func tryAgxpsApsProfileDataGetCounterGroupID(profileData AGXPSProfileData, out []byte, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterGroupID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_group_id", "", _agxpsApsProfileDataGetCounterGroupIDErr)
	}
	return _agxpsApsProfileDataGetCounterGroupID(profileData, unsafe.SliceData(out), first, count), nil
}

// AgxpsApsProfileDataGetCounterGroupID.
func AgxpsApsProfileDataGetCounterGroupID(profileData AGXPSProfileData, out []byte, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterGroupID(profileData, out, first, count)
}

var _agxpsApsProfileDataGetCounterGroupMetadata func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterGroupMetadataErr error

func tryAgxpsApsProfileDataGetCounterGroupMetadata(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterGroupMetadata == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_group_metadata", "", _agxpsApsProfileDataGetCounterGroupMetadataErr)
	}
	return _agxpsApsProfileDataGetCounterGroupMetadata(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetCounterGroupMetadata.
func AgxpsApsProfileDataGetCounterGroupMetadata(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterGroupMetadata(profileData, out, first, count)
}

var _agxpsApsProfileDataGetCounterNames func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterNamesErr error

func tryAgxpsApsProfileDataGetCounterNames(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterNames == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_names", "", _agxpsApsProfileDataGetCounterNamesErr)
	}
	return _agxpsApsProfileDataGetCounterNames(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetCounterNames.
func AgxpsApsProfileDataGetCounterNames(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterNames(profileData, out, first, count)
}

var _agxpsApsProfileDataGetCounterValues func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterValuesErr error

func tryAgxpsApsProfileDataGetCounterValues(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterValues == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_values", "", _agxpsApsProfileDataGetCounterValuesErr)
	}
	return _agxpsApsProfileDataGetCounterValues(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetCounterValues.
func AgxpsApsProfileDataGetCounterValues(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterValues(profileData, out, first, count)
}

var _agxpsApsProfileDataGetCounterValuesNum func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterValuesNumErr error

func tryAgxpsApsProfileDataGetCounterValuesNum(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterValuesNum == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_values_num", "", _agxpsApsProfileDataGetCounterValuesNumErr)
	}
	return _agxpsApsProfileDataGetCounterValuesNum(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetCounterValuesNum.
func AgxpsApsProfileDataGetCounterValuesNum(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterValuesNum(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueCliqueID func(profileData AGXPSProfileData, out *byte, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueCliqueIDErr error

func tryAgxpsApsProfileDataGetEslCliqueCliqueID(profileData AGXPSProfileData, out []byte, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueCliqueID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_clique_id", "", _agxpsApsProfileDataGetEslCliqueCliqueIDErr)
	}
	return _agxpsApsProfileDataGetEslCliqueCliqueID(profileData, unsafe.SliceData(out), first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueCliqueID.
func AgxpsApsProfileDataGetEslCliqueCliqueID(profileData AGXPSProfileData, out []byte, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueCliqueID(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueEnd func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueEndErr error

func tryAgxpsApsProfileDataGetEslCliqueEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueEnd == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_end", "", _agxpsApsProfileDataGetEslCliqueEndErr)
	}
	return _agxpsApsProfileDataGetEslCliqueEnd(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueEnd.
func AgxpsApsProfileDataGetEslCliqueEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueEnd(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueEslID func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueEslIDErr error

func tryAgxpsApsProfileDataGetEslCliqueEslID(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueEslID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_esl_id", "", _agxpsApsProfileDataGetEslCliqueEslIDErr)
	}
	return _agxpsApsProfileDataGetEslCliqueEslID(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueEslID.
func AgxpsApsProfileDataGetEslCliqueEslID(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueEslID(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueInstructionTrace func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueInstructionTraceErr error

func tryAgxpsApsProfileDataGetEslCliqueInstructionTrace(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueInstructionTrace == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_instruction_trace", "", _agxpsApsProfileDataGetEslCliqueInstructionTraceErr)
	}
	return _agxpsApsProfileDataGetEslCliqueInstructionTrace(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueInstructionTrace.
func AgxpsApsProfileDataGetEslCliqueInstructionTrace(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueInstructionTrace(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueKickID func(profileData AGXPSProfileData, out *uint32, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueKickIDErr error

func tryAgxpsApsProfileDataGetEslCliqueKickID(profileData AGXPSProfileData, out *uint32, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueKickID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_kick_id", "", _agxpsApsProfileDataGetEslCliqueKickIDErr)
	}
	return _agxpsApsProfileDataGetEslCliqueKickID(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueKickID.
func AgxpsApsProfileDataGetEslCliqueKickID(profileData AGXPSProfileData, out *uint32, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueKickID(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueMissingEnd func(profileData AGXPSProfileData, out *byte, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueMissingEndErr error

func tryAgxpsApsProfileDataGetEslCliqueMissingEnd(profileData AGXPSProfileData, out []byte, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueMissingEnd == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_missing_end", "", _agxpsApsProfileDataGetEslCliqueMissingEndErr)
	}
	return _agxpsApsProfileDataGetEslCliqueMissingEnd(profileData, unsafe.SliceData(out), first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueMissingEnd.
func AgxpsApsProfileDataGetEslCliqueMissingEnd(profileData AGXPSProfileData, out []byte, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueMissingEnd(profileData, out, first, count)
}

var _agxpsApsProfileDataGetEslCliqueStart func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetEslCliqueStartErr error

func tryAgxpsApsProfileDataGetEslCliqueStart(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetEslCliqueStart == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_esl_clique_start", "", _agxpsApsProfileDataGetEslCliqueStartErr)
	}
	return _agxpsApsProfileDataGetEslCliqueStart(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetEslCliqueStart.
func AgxpsApsProfileDataGetEslCliqueStart(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetEslCliqueStart(profileData, out, first, count)
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
func AgxpsApsProfileDataGetEslCliquesNum(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxpsApsProfileDataGetEslCliquesNum(profileData)
}

var _agxpsApsProfileDataGetKickEnd func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetKickEndErr error

func tryAgxpsApsProfileDataGetKickEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetKickEnd == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_kick_end", "", _agxpsApsProfileDataGetKickEndErr)
	}
	return _agxpsApsProfileDataGetKickEnd(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetKickEnd.
func AgxpsApsProfileDataGetKickEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetKickEnd(profileData, out, first, count)
}

var _agxpsApsProfileDataGetKickID func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetKickIDErr error

func tryAgxpsApsProfileDataGetKickID(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetKickID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_kick_id", "", _agxpsApsProfileDataGetKickIDErr)
	}
	return _agxpsApsProfileDataGetKickID(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetKickID.
func AgxpsApsProfileDataGetKickID(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetKickID(profileData, out, first, count)
}

var _agxpsApsProfileDataGetKickKickSlot func(profileData AGXPSProfileData, out *uint16, first uint64, count uint64) bool
var _agxpsApsProfileDataGetKickKickSlotErr error

func tryAgxpsApsProfileDataGetKickKickSlot(profileData AGXPSProfileData, out *uint16, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetKickKickSlot == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_kick_kick_slot", "", _agxpsApsProfileDataGetKickKickSlotErr)
	}
	return _agxpsApsProfileDataGetKickKickSlot(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetKickKickSlot.
func AgxpsApsProfileDataGetKickKickSlot(profileData AGXPSProfileData, out *uint16, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetKickKickSlot(profileData, out, first, count)
}

var _agxpsApsProfileDataGetKickSoftwareID func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetKickSoftwareIDErr error

func tryAgxpsApsProfileDataGetKickSoftwareID(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetKickSoftwareID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_kick_software_id", "", _agxpsApsProfileDataGetKickSoftwareIDErr)
	}
	return _agxpsApsProfileDataGetKickSoftwareID(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetKickSoftwareID.
func AgxpsApsProfileDataGetKickSoftwareID(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetKickSoftwareID(profileData, out, first, count)
}

var _agxpsApsProfileDataGetKickStart func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetKickStartErr error

func tryAgxpsApsProfileDataGetKickStart(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetKickStart == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_kick_start", "", _agxpsApsProfileDataGetKickStartErr)
	}
	return _agxpsApsProfileDataGetKickStart(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetKickStart.
func AgxpsApsProfileDataGetKickStart(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetKickStart(profileData, out, first, count)
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
func AgxpsApsProfileDataGetKicksNum(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxpsApsProfileDataGetKicksNum(profileData)
}

var _agxpsApsProfileDataGetSystemTimestamps func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetSystemTimestampsErr error

func tryAgxpsApsProfileDataGetSystemTimestamps(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetSystemTimestamps == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_system_timestamps", "", _agxpsApsProfileDataGetSystemTimestampsErr)
	}
	return _agxpsApsProfileDataGetSystemTimestamps(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetSystemTimestamps.
func AgxpsApsProfileDataGetSystemTimestamps(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetSystemTimestamps(profileData, out, first, count)
}

var _agxpsApsProfileDataGetWorkCliqueEnd func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetWorkCliqueEndErr error

func tryAgxpsApsProfileDataGetWorkCliqueEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetWorkCliqueEnd == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_work_clique_end", "", _agxpsApsProfileDataGetWorkCliqueEndErr)
	}
	return _agxpsApsProfileDataGetWorkCliqueEnd(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetWorkCliqueEnd.
func AgxpsApsProfileDataGetWorkCliqueEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetWorkCliqueEnd(profileData, out, first, count)
}

var _agxpsApsProfileDataGetWorkCliqueInstructionTrace func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetWorkCliqueInstructionTraceErr error

func tryAgxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetWorkCliqueInstructionTrace == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_work_clique_instruction_trace", "", _agxpsApsProfileDataGetWorkCliqueInstructionTraceErr)
	}
	return _agxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetWorkCliqueInstructionTrace.
func AgxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetWorkCliqueInstructionTrace(profileData, out, first, count)
}

var _agxpsApsProfileDataGetWorkCliqueStart func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetWorkCliqueStartErr error

func tryAgxpsApsProfileDataGetWorkCliqueStart(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetWorkCliqueStart == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_work_clique_start", "", _agxpsApsProfileDataGetWorkCliqueStartErr)
	}
	return _agxpsApsProfileDataGetWorkCliqueStart(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetWorkCliqueStart.
func AgxpsApsProfileDataGetWorkCliqueStart(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetWorkCliqueStart(profileData, out, first, count)
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
func AgxpsApsProfileDataIsValid(profileData AGXPSProfileData) (bool, error) {
	return tryAgxpsApsProfileDataIsValid(profileData)
}

var _agxpsApsSystemTimestampToNanoseconds func(timestamp uint64) float64
var _agxpsApsSystemTimestampToNanosecondsErr error

func tryAgxpsApsSystemTimestampToNanoseconds(timestamp uint64) (float64, error) {
	if _agxpsApsSystemTimestampToNanoseconds == nil {
		return 0.0, symbolCallError("agxps_aps_system_timestamp_to_nanoseconds", "", _agxpsApsSystemTimestampToNanosecondsErr)
	}
	return _agxpsApsSystemTimestampToNanoseconds(timestamp), nil
}

// AgxpsApsSystemTimestampToNanoseconds.
func AgxpsApsSystemTimestampToNanoseconds(timestamp uint64) (float64, error) {
	return tryAgxpsApsSystemTimestampToNanoseconds(timestamp)
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
func AgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer uintptr) (float64, error) {
	return tryAgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer)
}

var _agxpsGPUCreate func(gen uint32, variant uint32, rev uint32) AGXPSGPU
var _agxpsGPUCreateErr error

func tryAgxpsGPUCreate(gen uint32, variant uint32, rev uint32) (AGXPSGPU, error) {
	if _agxpsGPUCreate == nil {
		return *new(AGXPSGPU), symbolCallError("agxps_gpu_create", "", _agxpsGPUCreateErr)
	}
	return _agxpsGPUCreate(gen, variant, rev), nil
}

// AgxpsGPUCreate.
func AgxpsGPUCreate(gen uint32, variant uint32, rev uint32) (AGXPSGPU, error) {
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
func AgxpsGPUDestroy(gpu AGXPSGPU) error {
	return tryAgxpsGPUDestroy(gpu)
}

var _agxpsGPUFormatName func(gpu AGXPSGPU, buf *byte, size uint64) int32
var _agxpsGPUFormatNameErr error

func tryAgxpsGPUFormatName(gpu AGXPSGPU, buf *byte, size uint64) (int32, error) {
	if _agxpsGPUFormatName == nil {
		return 0, symbolCallError("agxps_gpu_format_name", "", _agxpsGPUFormatNameErr)
	}
	return _agxpsGPUFormatName(gpu, buf, size), nil
}

// AgxpsGPUFormatName.
func AgxpsGPUFormatName(gpu AGXPSGPU, buf *byte, size uint64) (int32, error) {
	return tryAgxpsGPUFormatName(gpu, buf, size)
}

var _agxpsGPUGetGen func(gpu AGXPSGPU) uint32
var _agxpsGPUGetGenErr error

func tryAgxpsGPUGetGen(gpu AGXPSGPU) (uint32, error) {
	if _agxpsGPUGetGen == nil {
		return 0, symbolCallError("agxps_gpu_get_gen", "", _agxpsGPUGetGenErr)
	}
	return _agxpsGPUGetGen(gpu), nil
}

// AgxpsGPUGetGen.
func AgxpsGPUGetGen(gpu AGXPSGPU) (uint32, error) {
	return tryAgxpsGPUGetGen(gpu)
}

var _agxpsGPUGetRev func(gpu AGXPSGPU) uint32
var _agxpsGPUGetRevErr error

func tryAgxpsGPUGetRev(gpu AGXPSGPU) (uint32, error) {
	if _agxpsGPUGetRev == nil {
		return 0, symbolCallError("agxps_gpu_get_rev", "", _agxpsGPUGetRevErr)
	}
	return _agxpsGPUGetRev(gpu), nil
}

// AgxpsGPUGetRev.
func AgxpsGPUGetRev(gpu AGXPSGPU) (uint32, error) {
	return tryAgxpsGPUGetRev(gpu)
}

var _agxpsGPUGetVariant func(gpu AGXPSGPU) uint32
var _agxpsGPUGetVariantErr error

func tryAgxpsGPUGetVariant(gpu AGXPSGPU) (uint32, error) {
	if _agxpsGPUGetVariant == nil {
		return 0, symbolCallError("agxps_gpu_get_variant", "", _agxpsGPUGetVariantErr)
	}
	return _agxpsGPUGetVariant(gpu), nil
}

// AgxpsGPUGetVariant.
func AgxpsGPUGetVariant(gpu AGXPSGPU) (uint32, error) {
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
func AgxpsGPUIsValid(gpu AGXPSGPU) (bool, error) {
	return tryAgxpsGPUIsValid(gpu)
}

var _agxpsInitialize func() int32
var _agxpsInitializeErr error

func tryAgxpsInitialize() (int32, error) {
	if _agxpsInitialize == nil {
		return 0, symbolCallError("agxps_initialize", "", _agxpsInitializeErr)
	}
	return _agxpsInitialize(), nil
}

// AgxpsInitialize.
func AgxpsInitialize() (int32, error) {
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
	registerFunc(&_agxpsApsProfileDataGetCounterGroupID, &_agxpsApsProfileDataGetCounterGroupIDErr, frameworkHandle, "agxps_aps_profile_data_get_counter_group_id", "")
	registerFunc(&_agxpsApsProfileDataGetCounterGroupMetadata, &_agxpsApsProfileDataGetCounterGroupMetadataErr, frameworkHandle, "agxps_aps_profile_data_get_counter_group_metadata", "")
	registerFunc(&_agxpsApsProfileDataGetCounterNames, &_agxpsApsProfileDataGetCounterNamesErr, frameworkHandle, "agxps_aps_profile_data_get_counter_names", "")
	registerFunc(&_agxpsApsProfileDataGetCounterValues, &_agxpsApsProfileDataGetCounterValuesErr, frameworkHandle, "agxps_aps_profile_data_get_counter_values", "")
	registerFunc(&_agxpsApsProfileDataGetCounterValuesNum, &_agxpsApsProfileDataGetCounterValuesNumErr, frameworkHandle, "agxps_aps_profile_data_get_counter_values_num", "")
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
	registerFunc(&_agxpsApsProfileDataGetKickKickSlot, &_agxpsApsProfileDataGetKickKickSlotErr, frameworkHandle, "agxps_aps_profile_data_get_kick_kick_slot", "")
	registerFunc(&_agxpsApsProfileDataGetKickSoftwareID, &_agxpsApsProfileDataGetKickSoftwareIDErr, frameworkHandle, "agxps_aps_profile_data_get_kick_software_id", "")
	registerFunc(&_agxpsApsProfileDataGetKickStart, &_agxpsApsProfileDataGetKickStartErr, frameworkHandle, "agxps_aps_profile_data_get_kick_start", "")
	registerFunc(&_agxpsApsProfileDataGetKicksNum, &_agxpsApsProfileDataGetKicksNumErr, frameworkHandle, "agxps_aps_profile_data_get_kicks_num", "")
	registerFunc(&_agxpsApsProfileDataGetSystemTimestamps, &_agxpsApsProfileDataGetSystemTimestampsErr, frameworkHandle, "agxps_aps_profile_data_get_system_timestamps", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliqueEnd, &_agxpsApsProfileDataGetWorkCliqueEndErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_end", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliqueInstructionTrace, &_agxpsApsProfileDataGetWorkCliqueInstructionTraceErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_instruction_trace", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliqueStart, &_agxpsApsProfileDataGetWorkCliqueStartErr, frameworkHandle, "agxps_aps_profile_data_get_work_clique_start", "")
	registerFunc(&_agxpsApsProfileDataGetWorkCliquesNum, &_agxpsApsProfileDataGetWorkCliquesNumErr, frameworkHandle, "agxps_aps_profile_data_get_work_cliques_num", "")
	registerFunc(&_agxpsApsProfileDataIsValid, &_agxpsApsProfileDataIsValidErr, frameworkHandle, "agxps_aps_profile_data_is_valid", "")
	registerFunc(&_agxpsApsSystemTimestampToNanoseconds, &_agxpsApsSystemTimestampToNanosecondsErr, frameworkHandle, "agxps_aps_system_timestamp_to_nanoseconds", "")
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
