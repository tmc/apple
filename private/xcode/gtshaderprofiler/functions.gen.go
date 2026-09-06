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

// SymbolAddress returns the address of name in gtshaderprofiler, whether or not
// this package generated a binding for it.
//
// What is generated is bounded by what is documented, and for a private
// framework that is whatever a manifest happened to enumerate. The dylib
// usually exports far more. Without this, a symbol nobody wrote down is
// unreachable from a package that has already loaded the image holding it, and
// the generated surface becomes a ceiling instead of a floor.
//
// The lookup is scoped to this framework's handle, not RTLD_DEFAULT, so a
// symbol some other loaded image exports is not reported as this one's.
func SymbolAddress(name string) (uintptr, error) {
	if frameworkHandle == 0 {
		return 0, fmt.Errorf("gtshaderprofiler: symbol %s unavailable because the framework could not be loaded", name)
	}
	sym, err := purego.Dlsym(frameworkHandle, name)
	if err != nil || sym == 0 {
		return 0, missingSymbolError(name, "", err)
	}
	return sym, nil
}

// BindFunc binds the gtshaderprofiler symbol name into fptr, which must be a
// pointer to a func variable.
//
// The caller supplies the signature, and nothing checks it. A dylib records no
// argument count or types for a C symbol, so a wrong signature here is not a
// type error: it is the wrong number of machine words moved on a live stack,
// and the failure surfaces somewhere else entirely. Prefer a generated binding,
// whose signature carries recorded evidence, and reach for this only for a
// symbol that has none.
// purego.RegisterFunc panics on a signature it cannot lower; that is recovered
// and returned, because an escape hatch that takes down the process on a
// mistyped experiment is not one anybody can experiment with.
func BindFunc(fptr any, name string) (err error) {
	sym, err := SymbolAddress(name)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("gtshaderprofiler: bind symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	return nil
}

var _agxpsApsCliqueInstructionTraceGetExecutionEvents func(trace AGXPSCliqueInstructionTraceRef) unsafe.Pointer
var _agxpsApsCliqueInstructionTraceGetExecutionEventsErr error

func tryAgxpsApsCliqueInstructionTraceGetExecutionEvents(trace AGXPSCliqueInstructionTraceRef) (unsafe.Pointer, error) {
	if _agxpsApsCliqueInstructionTraceGetExecutionEvents == nil {
		return nil, symbolCallError("agxps_aps_clique_instruction_trace_get_execution_events", "", _agxpsApsCliqueInstructionTraceGetExecutionEventsErr)
	}
	return _agxpsApsCliqueInstructionTraceGetExecutionEvents(trace), nil
}

// AgxpsApsCliqueInstructionTraceGetExecutionEvents signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueInstructionTraceGetExecutionEventsNum signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueInstructionTraceGetInstructionStats signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueInstructionTraceGetPCAdvances signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueInstructionTraceGetPCAdvancesNum signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueInstructionTraceGetTimestampReferences signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueInstructionTraceGetTimestampReferencesNum signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsCliqueTimeStatsCreate signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func AgxpsApsCliqueTimeStatsCreate(profileData AGXPSProfileData, cliqueIndex uint64) (AGXPSCliqueTimeStatsRef, error) {
	return tryAgxpsApsCliqueTimeStatsCreate(profileData, cliqueIndex)
}

var _agxpsApsGPUIsSupported func(gpu AGXPSGPU) bool
var _agxpsApsGPUIsSupportedErr error

func tryAgxpsApsGPUIsSupported(gpu AGXPSGPU) (bool, error) {
	if _agxpsApsGPUIsSupported == nil {
		return false, symbolCallError("agxps_aps_gpu_is_supported", "", _agxpsApsGPUIsSupportedErr)
	}
	return _agxpsApsGPUIsSupported(gpu), nil
}

// AgxpsApsGPUIsSupported signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsParserCreate signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsParserDestroy signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsParserIsValid signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func AgxpsApsParserIsValid(parser AGXPSParserHandle) (bool, error) {
	return tryAgxpsApsParserIsValid(parser)
}

var _agxpsApsParserParse func(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, flags uint32, parseErrorOut *uint32) AGXPSProfileData
var _agxpsApsParserParseErr error

func tryAgxpsApsParserParse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, flags uint32, parseErrorOut *uint32) (AGXPSProfileData, error) {
	if _agxpsApsParserParse == nil {
		return *new(AGXPSProfileData), symbolCallError("agxps_aps_parser_parse", "", _agxpsApsParserParseErr)
	}
	return _agxpsApsParserParse(parser, data, size, flags, parseErrorOut), nil
}

// AgxpsApsParserParse signature evidence: otool -arch arm64 -tvV GTShaderProfiler at 0x4ea634 (Xcode 26.4): tail call preserves x0..x4, null branch stores error through x4; gputrace decoded all 40 Counters_f_*.raw shards of parity-asymmetric-perfdata.gputrace with this shape.
func AgxpsApsParserParse(parser AGXPSParserHandle, data unsafe.Pointer, size uint64, flags uint32, parseErrorOut *uint32) (AGXPSProfileData, error) {
	return tryAgxpsApsParserParse(parser, data, size, flags, parseErrorOut)
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

// AgxpsApsProfileDataDestroy signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsProfileDataGetCounterGroupID signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsProfileDataGetCounterGroupMetadata signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsProfileDataGetCounterNames signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
func AgxpsApsProfileDataGetCounterNames(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterNames(profileData, out, first, count)
}

var _agxpsApsProfileDataGetCounterNum func(profileData AGXPSProfileData) uint64
var _agxpsApsProfileDataGetCounterNumErr error

func tryAgxpsApsProfileDataGetCounterNum(profileData AGXPSProfileData) (uint64, error) {
	if _agxpsApsProfileDataGetCounterNum == nil {
		return 0, symbolCallError("agxps_aps_profile_data_get_counter_num", "", _agxpsApsProfileDataGetCounterNumErr)
	}
	return _agxpsApsProfileDataGetCounterNum(profileData), nil
}

// AgxpsApsProfileDataGetCounterNum signature evidence: otool -arch arm64 -tvV GTShaderProfiler at 0x4ed7b4 (Xcode 26.4): (end-begin)>>3 over the counter-ID vector at +0x371b8/+0x371c0.
func AgxpsApsProfileDataGetCounterNum(profileData AGXPSProfileData) (uint64, error) {
	return tryAgxpsApsProfileDataGetCounterNum(profileData)
}

var _agxpsApsProfileDataGetCounterValues func(profileData AGXPSProfileData, beginPointersOut *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterValuesErr error

func tryAgxpsApsProfileDataGetCounterValues(profileData AGXPSProfileData, beginPointersOut *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterValues == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_values", "", _agxpsApsProfileDataGetCounterValuesErr)
	}
	return _agxpsApsProfileDataGetCounterValues(profileData, beginPointersOut, first, count), nil
}

// AgxpsApsProfileDataGetCounterValues signature evidence: otool -arch arm64 -tvV GTShaderProfiler at 0x4edce4 (Xcode 26.4): stores one begin() pointer per counter from the 24-byte record table at +0x30f48; get_counter_values_num stores (end-begin)>>3 from the same record.
func AgxpsApsProfileDataGetCounterValues(profileData AGXPSProfileData, beginPointersOut *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetCounterValues(profileData, beginPointersOut, first, count)
}

var _agxpsApsProfileDataGetCounterValuesNum func(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) bool
var _agxpsApsProfileDataGetCounterValuesNumErr error

func tryAgxpsApsProfileDataGetCounterValuesNum(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetCounterValuesNum == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_counter_values_num", "", _agxpsApsProfileDataGetCounterValuesNumErr)
	}
	return _agxpsApsProfileDataGetCounterValuesNum(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetCounterValuesNum signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsProfileDataGetEslCliqueCliqueID signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4ecba0 ldr x10, 0x4ecba4 strb w10,[x1],#1 -- 1-byte store.
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

// AgxpsApsProfileDataGetEslCliqueEnd signature evidence: gputrace internal/agxps/rawprobe_manual_test.go, live capture; bulk range shape confirmed.
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

// AgxpsApsProfileDataGetEslCliqueEslID signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4ec9bc ldr x12, 0x4ec9c8 str x12,[x1],#8.
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

// AgxpsApsProfileDataGetEslCliqueInstructionTrace signature evidence: gputrace internal/agxps/rawprobe_manual_test.go, live capture; bulk range shape confirmed.
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

// AgxpsApsProfileDataGetEslCliqueKickID signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4eca74 ldr w12, 0x4eca80 str w12,[x1],#4 -- 4-byte store.
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

// AgxpsApsProfileDataGetEslCliqueMissingEnd signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4ecb24 computes bit, 0x4ecb30 strb w12,[x1],#1.
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

// AgxpsApsProfileDataGetEslCliqueStart signature evidence: gputrace internal/agxps/rawprobe_manual_test.go, live capture; bulk range shape confirmed.
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

// AgxpsApsProfileDataGetEslCliquesNum signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsProfileDataGetKickEnd signature evidence: gputrace internal/agxps/rawprobe_manual_test.go, live capture; bulk range shape confirmed.
func AgxpsApsProfileDataGetKickEnd(profileData AGXPSProfileData, out *uint64, first uint64, count uint64) (bool, error) {
	return tryAgxpsApsProfileDataGetKickEnd(profileData, out, first, count)
}

var _agxpsApsProfileDataGetKickID func(profileData AGXPSProfileData, out *uint32, first uint64, count uint64) bool
var _agxpsApsProfileDataGetKickIDErr error

func tryAgxpsApsProfileDataGetKickID(profileData AGXPSProfileData, out *uint32, first uint64, count uint64) (bool, error) {
	if _agxpsApsProfileDataGetKickID == nil {
		return false, symbolCallError("agxps_aps_profile_data_get_kick_id", "", _agxpsApsProfileDataGetKickIDErr)
	}
	return _agxpsApsProfileDataGetKickID(profileData, out, first, count), nil
}

// AgxpsApsProfileDataGetKickID signature evidence: gputrace poisoned-buffer probe on Counters_f_0.raw: nk=3792, first unwritten index 1896 = nk/2, so 4 bytes/element; gputrace docs/research/agxps-signatures.yaml bytes:4 verified:runtime.
func AgxpsApsProfileDataGetKickID(profileData AGXPSProfileData, out *uint32, first uint64, count uint64) (bool, error) {
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

// AgxpsApsProfileDataGetKickKickSlot signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsProfileDataGetKickSoftwareID signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsProfileDataGetKickStart signature evidence: gputrace internal/agxps/rawprobe_manual_test.go, live capture; bulk range shape confirmed.
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

// AgxpsApsProfileDataGetKicksNum signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsProfileDataGetSystemTimestamps signature evidence: gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsProfileDataGetWorkCliqueEnd signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4ec204 ldr x12, 0x4ec210 str x12,[x1],#8.
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

// AgxpsApsProfileDataGetWorkCliqueInstructionTrace signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4ec6d0 ldr x10, 0x4ec6d4 str x10,[x1],#8.
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

// AgxpsApsProfileDataGetWorkCliqueStart signature evidence: IPSW disassembly, GTShaderProfiler arm64 UUID 4115D609-CBAF-3394-8F8E-564F03F8588E: 0x4ec14c ldr x12, 0x4ec158 str x12,[x1],#8.
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

// AgxpsApsProfileDataGetWorkCliquesNum signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsProfileDataIsValid signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsSystemTimestampToNanoseconds signature evidence: IPSW disassembly 0x4ee29c, floating-point return; gputrace internal/agxps/counterprobe_manual_test.go, live capture.
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

// AgxpsApsTimingAnalyzerGetNumCommands signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsTimingAnalyzerGetNumWorkCliques signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsTimingAnalyzerGetWorkCliquesAverageDuration signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsTimingAnalyzerGetWorkCliquesMaxDuration signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsApsTimingAnalyzerGetWorkCliquesMinDuration signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
func AgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer uintptr) (float64, error) {
	return tryAgxpsApsTimingAnalyzerGetWorkCliquesMinDuration(analyzer)
}

var _agxpsCounterComputeDerivedCounters func(gpu AGXPSGPU, desc unsafe.Pointer, rawInput unsafe.Pointer, rawCount uint32, constants unsafe.Pointer, constCount uint32, output *float64, outCount uint32) int32
var _agxpsCounterComputeDerivedCountersErr error

func tryAgxpsCounterComputeDerivedCounters(gpu AGXPSGPU, desc unsafe.Pointer, rawInput unsafe.Pointer, rawCount uint32, constants unsafe.Pointer, constCount uint32, output []float64, outCount uint32) (int32, error) {
	if _agxpsCounterComputeDerivedCounters == nil {
		return 0, symbolCallError("agxps_counter_compute_derived_counters", "", _agxpsCounterComputeDerivedCountersErr)
	}
	return _agxpsCounterComputeDerivedCounters(gpu, desc, rawInput, rawCount, constants, constCount, unsafe.SliceData(output), outCount), nil
}

// AgxpsCounterComputeDerivedCounters signature evidence: GTShaderProfiler.framework symbol @ 0x558f6c; evaluates 174 derived limiters from raw counters and constants.
func AgxpsCounterComputeDerivedCounters(gpu AGXPSGPU, desc unsafe.Pointer, rawInput unsafe.Pointer, rawCount uint32, constants unsafe.Pointer, constCount uint32, output []float64, outCount uint32) (int32, error) {
	return tryAgxpsCounterComputeDerivedCounters(gpu, desc, rawInput, rawCount, constants, constCount, output, outCount)
}

var _agxpsCounterObfuscatedName func(name string) *byte
var _agxpsCounterObfuscatedNameErr error

func tryAgxpsCounterObfuscatedName(name string) (*byte, error) {
	if _agxpsCounterObfuscatedName == nil {
		return nil, symbolCallError("agxps_counter_obfuscated_name", "", _agxpsCounterObfuscatedNameErr)
	}
	return _agxpsCounterObfuscatedName(name), nil
}

// AgxpsCounterObfuscatedName signature evidence: otool -arch arm64 -tvV GTShaderProfiler (Xcode 26.4): 0x4adcd8 null-checks x0 and hands it straight to the std::string(const char *) constructor, so the argument is a name string and not a counter ident.
func AgxpsCounterObfuscatedName(name string) (*byte, error) {
	return tryAgxpsCounterObfuscatedName(name)
}

var _agxpsGPUCreate func(gen uint32, variant uint32, rev uint32, exact bool) AGXPSGPU
var _agxpsGPUCreateErr error

func tryAgxpsGPUCreate(gen uint32, variant uint32, rev uint32, exact bool) (AGXPSGPU, error) {
	if _agxpsGPUCreate == nil {
		return *new(AGXPSGPU), symbolCallError("agxps_gpu_create", "", _agxpsGPUCreateErr)
	}
	return _agxpsGPUCreate(gen, variant, rev, exact), nil
}

// AgxpsGPUCreate signature evidence: otool -arch arm64 -tvV GTShaderProfiler (Xcode 26.4): 0x49b528 "mov x23, x3" reads a fourth argument and 0x49b5a8 "tbnz w23, #0x0" tests only its bit 0, so it is a bool gating the revision fallback.
func AgxpsGPUCreate(gen uint32, variant uint32, rev uint32, exact bool) (AGXPSGPU, error) {
	return tryAgxpsGPUCreate(gen, variant, rev, exact)
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

// AgxpsGPUDestroy signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsGPUFormatName signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsGPUGetGen signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsGPUGetRev signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsGPUGetVariant signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsGPUIsValid signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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

// AgxpsInitialize signature evidence: none recorded; the argument count, order, and types are unverified and may all be wrong.
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
	registerFunc(&_agxpsApsGPUIsSupported, &_agxpsApsGPUIsSupportedErr, frameworkHandle, "agxps_aps_gpu_is_supported", "")
	registerFunc(&_agxpsApsParserCreate, &_agxpsApsParserCreateErr, frameworkHandle, "agxps_aps_parser_create", "")
	registerFunc(&_agxpsApsParserDestroy, &_agxpsApsParserDestroyErr, frameworkHandle, "agxps_aps_parser_destroy", "")
	registerFunc(&_agxpsApsParserIsValid, &_agxpsApsParserIsValidErr, frameworkHandle, "agxps_aps_parser_is_valid", "")
	registerFunc(&_agxpsApsParserParse, &_agxpsApsParserParseErr, frameworkHandle, "agxps_aps_parser_parse", "")
	registerFunc(&_agxpsApsProfileDataDestroy, &_agxpsApsProfileDataDestroyErr, frameworkHandle, "agxps_aps_profile_data_destroy", "")
	registerFunc(&_agxpsApsProfileDataGetCounterGroupID, &_agxpsApsProfileDataGetCounterGroupIDErr, frameworkHandle, "agxps_aps_profile_data_get_counter_group_id", "")
	registerFunc(&_agxpsApsProfileDataGetCounterGroupMetadata, &_agxpsApsProfileDataGetCounterGroupMetadataErr, frameworkHandle, "agxps_aps_profile_data_get_counter_group_metadata", "")
	registerFunc(&_agxpsApsProfileDataGetCounterNames, &_agxpsApsProfileDataGetCounterNamesErr, frameworkHandle, "agxps_aps_profile_data_get_counter_names", "")
	registerFunc(&_agxpsApsProfileDataGetCounterNum, &_agxpsApsProfileDataGetCounterNumErr, frameworkHandle, "agxps_aps_profile_data_get_counter_num", "")
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
	registerFunc(&_agxpsCounterComputeDerivedCounters, &_agxpsCounterComputeDerivedCountersErr, frameworkHandle, "agxps_counter_compute_derived_counters", "")
	registerFunc(&_agxpsCounterObfuscatedName, &_agxpsCounterObfuscatedNameErr, frameworkHandle, "agxps_counter_obfuscated_name", "")
	registerFunc(&_agxpsGPUCreate, &_agxpsGPUCreateErr, frameworkHandle, "agxps_gpu_create", "")
	registerFunc(&_agxpsGPUDestroy, &_agxpsGPUDestroyErr, frameworkHandle, "agxps_gpu_destroy", "")
	registerFunc(&_agxpsGPUFormatName, &_agxpsGPUFormatNameErr, frameworkHandle, "agxps_gpu_format_name", "")
	registerFunc(&_agxpsGPUGetGen, &_agxpsGPUGetGenErr, frameworkHandle, "agxps_gpu_get_gen", "")
	registerFunc(&_agxpsGPUGetRev, &_agxpsGPUGetRevErr, frameworkHandle, "agxps_gpu_get_rev", "")
	registerFunc(&_agxpsGPUGetVariant, &_agxpsGPUGetVariantErr, frameworkHandle, "agxps_gpu_get_variant", "")
	registerFunc(&_agxpsGPUIsValid, &_agxpsGPUIsValidErr, frameworkHandle, "agxps_gpu_is_valid", "")
	registerFunc(&_agxpsInitialize, &_agxpsInitializeErr, frameworkHandle, "agxps_initialize", "")
}
