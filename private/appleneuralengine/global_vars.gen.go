// Code generated from Apple documentation. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objectivec"
)

var KANEFAOTCacheUrlIdentifierKey objectivec.Object

var KANEFBaseModelIdentifierKey objectivec.Object

var KANEFCompilationInitiatedByE5MLKey objectivec.Object

var KANEFCompilerOptionsFilenameKey objectivec.Object

var KANEFConstantSurfaceAlignmentKey objectivec.Object

var KANEFConstantSurfaceIDKey objectivec.Object

var KANEFDisableIOFencesUseSharedEventsKey objectivec.Object

var KANEFEnableFWToFWSignal objectivec.Object

var KANEFEnableLateLatchKey objectivec.Object

var KANEFEnablePowerSavingKey objectivec.Object

var KANEFEspressoFileResourcesKey objectivec.Object

var KANEFHintEnergyEfficientWorkloadKey objectivec.Object

var KANEFHintReportResidentPagesKey objectivec.Object

var KANEFHintReportSessionStatusKey objectivec.Object

var KANEFHintReportTotalPagesKey objectivec.Object

var KANEFHintSessionAbort objectivec.Object

var KANEFHintSessionInfo objectivec.Object

var KANEFHintSessionStart objectivec.Object

var KANEFHintSessionStop objectivec.Object

var KANEFInMemoryModelIdentifierKey objectivec.Object

var KANEFInMemoryModelIsCachedKey objectivec.Object

var KANEFIntermediateBufferHandleKey objectivec.Object

var KANEFKeepModelMemoryWiredKey objectivec.Object

var KANEFMemoryPoolIDKey objectivec.Object

var KANEFModelANECIRValue objectivec.Object

var KANEFModelCacheIdentifierUsingSourceURLKey objectivec.Object

var KANEFModelCoreMLValue objectivec.Object

var KANEFModelDescriptionKey objectivec.Object

var KANEFModelHasCacheURLIdentifierKey objectivec.Object

var KANEFModelIdentityStrKey objectivec.Object

var KANEFModelInput16KAlignmentArrayKey objectivec.Object

var KANEFModelInputSymbolIndexArrayKey objectivec.Object

var KANEFModelInputSymbolsArrayKey objectivec.Object

var KANEFModelInstanceParameters objectivec.Object

var KANEFModelIsEncryptedKey objectivec.Object

var KANEFModelLLIRBundleValue objectivec.Object

var KANEFModelLoadPerformanceStatsKey objectivec.Object

var KANEFModelMILValue objectivec.Object

var KANEFModelMLIRValue objectivec.Object

var KANEFModelOutput16KAlignmentArrayKey objectivec.Object

var KANEFModelOutputSymbolIndexArrayKey objectivec.Object

var KANEFModelOutputSymbolsArrayKey objectivec.Object

var KANEFModelPreCompiledValue objectivec.Object

var KANEFModelProcedureIDKey objectivec.Object

var KANEFModelProcedureNameToIDMapKey objectivec.Object

var KANEFModelProcedureNameToStatsSizeMapKey objectivec.Object

var KANEFModelProceduresArrayKey objectivec.Object

var KANEFModelTypeKey objectivec.Object

var KANEFNetPlistFilenameKey objectivec.Object

var KANEFPerformanceStatsMaskKey objectivec.Object

var KANEFRetainModelsWithoutSourceURLKey objectivec.Object

var KANEFSkipPreparePhaseKey objectivec.Object

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFAOTCacheUrlIdentifierKey"); err == nil && ptr != 0 {
		KANEFAOTCacheUrlIdentifierKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFBaseModelIdentifierKey"); err == nil && ptr != 0 {
		KANEFBaseModelIdentifierKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFCompilationInitiatedByE5MLKey"); err == nil && ptr != 0 {
		KANEFCompilationInitiatedByE5MLKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFCompilerOptionsFilenameKey"); err == nil && ptr != 0 {
		KANEFCompilerOptionsFilenameKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFConstantSurfaceAlignmentKey"); err == nil && ptr != 0 {
		KANEFConstantSurfaceAlignmentKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFConstantSurfaceIDKey"); err == nil && ptr != 0 {
		KANEFConstantSurfaceIDKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFDisableIOFencesUseSharedEventsKey"); err == nil && ptr != 0 {
		KANEFDisableIOFencesUseSharedEventsKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEnableFWToFWSignal"); err == nil && ptr != 0 {
		KANEFEnableFWToFWSignal = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEnableLateLatchKey"); err == nil && ptr != 0 {
		KANEFEnableLateLatchKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEnablePowerSavingKey"); err == nil && ptr != 0 {
		KANEFEnablePowerSavingKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEspressoFileResourcesKey"); err == nil && ptr != 0 {
		KANEFEspressoFileResourcesKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintEnergyEfficientWorkloadKey"); err == nil && ptr != 0 {
		KANEFHintEnergyEfficientWorkloadKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintReportResidentPagesKey"); err == nil && ptr != 0 {
		KANEFHintReportResidentPagesKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintReportSessionStatusKey"); err == nil && ptr != 0 {
		KANEFHintReportSessionStatusKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintReportTotalPagesKey"); err == nil && ptr != 0 {
		KANEFHintReportTotalPagesKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionAbort"); err == nil && ptr != 0 {
		KANEFHintSessionAbort = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionInfo"); err == nil && ptr != 0 {
		KANEFHintSessionInfo = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionStart"); err == nil && ptr != 0 {
		KANEFHintSessionStart = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionStop"); err == nil && ptr != 0 {
		KANEFHintSessionStop = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFInMemoryModelIdentifierKey"); err == nil && ptr != 0 {
		KANEFInMemoryModelIdentifierKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFInMemoryModelIsCachedKey"); err == nil && ptr != 0 {
		KANEFInMemoryModelIsCachedKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFIntermediateBufferHandleKey"); err == nil && ptr != 0 {
		KANEFIntermediateBufferHandleKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFKeepModelMemoryWiredKey"); err == nil && ptr != 0 {
		KANEFKeepModelMemoryWiredKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFMemoryPoolIDKey"); err == nil && ptr != 0 {
		KANEFMemoryPoolIDKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelANECIRValue"); err == nil && ptr != 0 {
		KANEFModelANECIRValue = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelCacheIdentifierUsingSourceURLKey"); err == nil && ptr != 0 {
		KANEFModelCacheIdentifierUsingSourceURLKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelCoreMLValue"); err == nil && ptr != 0 {
		KANEFModelCoreMLValue = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelDescriptionKey"); err == nil && ptr != 0 {
		KANEFModelDescriptionKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelHasCacheURLIdentifierKey"); err == nil && ptr != 0 {
		KANEFModelHasCacheURLIdentifierKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelIdentityStrKey"); err == nil && ptr != 0 {
		KANEFModelIdentityStrKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInput16KAlignmentArrayKey"); err == nil && ptr != 0 {
		KANEFModelInput16KAlignmentArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInputSymbolIndexArrayKey"); err == nil && ptr != 0 {
		KANEFModelInputSymbolIndexArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInputSymbolsArrayKey"); err == nil && ptr != 0 {
		KANEFModelInputSymbolsArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInstanceParameters"); err == nil && ptr != 0 {
		KANEFModelInstanceParameters = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelIsEncryptedKey"); err == nil && ptr != 0 {
		KANEFModelIsEncryptedKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelLLIRBundleValue"); err == nil && ptr != 0 {
		KANEFModelLLIRBundleValue = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelLoadPerformanceStatsKey"); err == nil && ptr != 0 {
		KANEFModelLoadPerformanceStatsKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelMILValue"); err == nil && ptr != 0 {
		KANEFModelMILValue = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelMLIRValue"); err == nil && ptr != 0 {
		KANEFModelMLIRValue = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelOutput16KAlignmentArrayKey"); err == nil && ptr != 0 {
		KANEFModelOutput16KAlignmentArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelOutputSymbolIndexArrayKey"); err == nil && ptr != 0 {
		KANEFModelOutputSymbolIndexArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelOutputSymbolsArrayKey"); err == nil && ptr != 0 {
		KANEFModelOutputSymbolsArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelPreCompiledValue"); err == nil && ptr != 0 {
		KANEFModelPreCompiledValue = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProcedureIDKey"); err == nil && ptr != 0 {
		KANEFModelProcedureIDKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProcedureNameToIDMapKey"); err == nil && ptr != 0 {
		KANEFModelProcedureNameToIDMapKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProcedureNameToStatsSizeMapKey"); err == nil && ptr != 0 {
		KANEFModelProcedureNameToStatsSizeMapKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProceduresArrayKey"); err == nil && ptr != 0 {
		KANEFModelProceduresArrayKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelTypeKey"); err == nil && ptr != 0 {
		KANEFModelTypeKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFNetPlistFilenameKey"); err == nil && ptr != 0 {
		KANEFNetPlistFilenameKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFPerformanceStatsMaskKey"); err == nil && ptr != 0 {
		KANEFPerformanceStatsMaskKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFRetainModelsWithoutSourceURLKey"); err == nil && ptr != 0 {
		KANEFRetainModelsWithoutSourceURLKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFSkipPreparePhaseKey"); err == nil && ptr != 0 {
		KANEFSkipPreparePhaseKey = *(*objectivec.Object)(unsafe.Pointer(ptr))
	}

}

