// Code generated from Apple documentation. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
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
		KANEFAOTCacheUrlIdentifierKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFBaseModelIdentifierKey"); err == nil && ptr != 0 {
		KANEFBaseModelIdentifierKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFCompilationInitiatedByE5MLKey"); err == nil && ptr != 0 {
		KANEFCompilationInitiatedByE5MLKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFCompilerOptionsFilenameKey"); err == nil && ptr != 0 {
		KANEFCompilerOptionsFilenameKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFConstantSurfaceAlignmentKey"); err == nil && ptr != 0 {
		KANEFConstantSurfaceAlignmentKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFConstantSurfaceIDKey"); err == nil && ptr != 0 {
		KANEFConstantSurfaceIDKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFDisableIOFencesUseSharedEventsKey"); err == nil && ptr != 0 {
		KANEFDisableIOFencesUseSharedEventsKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEnableFWToFWSignal"); err == nil && ptr != 0 {
		KANEFEnableFWToFWSignal = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEnableLateLatchKey"); err == nil && ptr != 0 {
		KANEFEnableLateLatchKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEnablePowerSavingKey"); err == nil && ptr != 0 {
		KANEFEnablePowerSavingKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFEspressoFileResourcesKey"); err == nil && ptr != 0 {
		KANEFEspressoFileResourcesKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintEnergyEfficientWorkloadKey"); err == nil && ptr != 0 {
		KANEFHintEnergyEfficientWorkloadKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintReportResidentPagesKey"); err == nil && ptr != 0 {
		KANEFHintReportResidentPagesKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintReportSessionStatusKey"); err == nil && ptr != 0 {
		KANEFHintReportSessionStatusKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintReportTotalPagesKey"); err == nil && ptr != 0 {
		KANEFHintReportTotalPagesKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionAbort"); err == nil && ptr != 0 {
		KANEFHintSessionAbort = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionInfo"); err == nil && ptr != 0 {
		KANEFHintSessionInfo = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionStart"); err == nil && ptr != 0 {
		KANEFHintSessionStart = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFHintSessionStop"); err == nil && ptr != 0 {
		KANEFHintSessionStop = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFInMemoryModelIdentifierKey"); err == nil && ptr != 0 {
		KANEFInMemoryModelIdentifierKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFInMemoryModelIsCachedKey"); err == nil && ptr != 0 {
		KANEFInMemoryModelIsCachedKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFIntermediateBufferHandleKey"); err == nil && ptr != 0 {
		KANEFIntermediateBufferHandleKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFKeepModelMemoryWiredKey"); err == nil && ptr != 0 {
		KANEFKeepModelMemoryWiredKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFMemoryPoolIDKey"); err == nil && ptr != 0 {
		KANEFMemoryPoolIDKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelANECIRValue"); err == nil && ptr != 0 {
		KANEFModelANECIRValue = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelCacheIdentifierUsingSourceURLKey"); err == nil && ptr != 0 {
		KANEFModelCacheIdentifierUsingSourceURLKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelCoreMLValue"); err == nil && ptr != 0 {
		KANEFModelCoreMLValue = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelDescriptionKey"); err == nil && ptr != 0 {
		KANEFModelDescriptionKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelHasCacheURLIdentifierKey"); err == nil && ptr != 0 {
		KANEFModelHasCacheURLIdentifierKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelIdentityStrKey"); err == nil && ptr != 0 {
		KANEFModelIdentityStrKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInput16KAlignmentArrayKey"); err == nil && ptr != 0 {
		KANEFModelInput16KAlignmentArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInputSymbolIndexArrayKey"); err == nil && ptr != 0 {
		KANEFModelInputSymbolIndexArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInputSymbolsArrayKey"); err == nil && ptr != 0 {
		KANEFModelInputSymbolsArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelInstanceParameters"); err == nil && ptr != 0 {
		KANEFModelInstanceParameters = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelIsEncryptedKey"); err == nil && ptr != 0 {
		KANEFModelIsEncryptedKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelLLIRBundleValue"); err == nil && ptr != 0 {
		KANEFModelLLIRBundleValue = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelLoadPerformanceStatsKey"); err == nil && ptr != 0 {
		KANEFModelLoadPerformanceStatsKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelMILValue"); err == nil && ptr != 0 {
		KANEFModelMILValue = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelMLIRValue"); err == nil && ptr != 0 {
		KANEFModelMLIRValue = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelOutput16KAlignmentArrayKey"); err == nil && ptr != 0 {
		KANEFModelOutput16KAlignmentArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelOutputSymbolIndexArrayKey"); err == nil && ptr != 0 {
		KANEFModelOutputSymbolIndexArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelOutputSymbolsArrayKey"); err == nil && ptr != 0 {
		KANEFModelOutputSymbolsArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelPreCompiledValue"); err == nil && ptr != 0 {
		KANEFModelPreCompiledValue = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProcedureIDKey"); err == nil && ptr != 0 {
		KANEFModelProcedureIDKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProcedureNameToIDMapKey"); err == nil && ptr != 0 {
		KANEFModelProcedureNameToIDMapKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProcedureNameToStatsSizeMapKey"); err == nil && ptr != 0 {
		KANEFModelProcedureNameToStatsSizeMapKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelProceduresArrayKey"); err == nil && ptr != 0 {
		KANEFModelProceduresArrayKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFModelTypeKey"); err == nil && ptr != 0 {
		KANEFModelTypeKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFNetPlistFilenameKey"); err == nil && ptr != 0 {
		KANEFNetPlistFilenameKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFPerformanceStatsMaskKey"); err == nil && ptr != 0 {
		KANEFPerformanceStatsMaskKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFRetainModelsWithoutSourceURLKey"); err == nil && ptr != 0 {
		KANEFRetainModelsWithoutSourceURLKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kANEFSkipPreparePhaseKey"); err == nil && ptr != 0 {
		KANEFSkipPreparePhaseKey = objectivec.ObjectFromID(objc.IDValueAt(ptr))
	}

}

