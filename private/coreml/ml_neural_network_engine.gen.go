// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkEngine] class.
var (
	_MLNeuralNetworkEngineClass     MLNeuralNetworkEngineClass
	_MLNeuralNetworkEngineClassOnce sync.Once
)

func getMLNeuralNetworkEngineClass() MLNeuralNetworkEngineClass {
	_MLNeuralNetworkEngineClassOnce.Do(func() {
		_MLNeuralNetworkEngineClass = MLNeuralNetworkEngineClass{class: objc.GetClass("MLNeuralNetworkEngine")}
	})
	return _MLNeuralNetworkEngineClass
}

// GetMLNeuralNetworkEngineClass returns the class object for MLNeuralNetworkEngine.
func GetMLNeuralNetworkEngineClass() MLNeuralNetworkEngineClass {
	return getMLNeuralNetworkEngineClass()
}

type MLNeuralNetworkEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkEngineClass) Alloc() MLNeuralNetworkEngine {
	rv := objc.Send[MLNeuralNetworkEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralNetworkEngine._addCompiledNetworkOrProgramToPlanError]
//   - [MLNeuralNetworkEngine._addNetworkToPlanError]
//   - [MLNeuralNetworkEngine._deallocContextAndPlan]
//   - [MLNeuralNetworkEngine._espressoDeviceForConfigurationError]
//   - [MLNeuralNetworkEngine._espressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray]
//   - [MLNeuralNetworkEngine._handleAddNetworkToPlanStatusError]
//   - [MLNeuralNetworkEngine._matchEngineToOptionsError]
//   - [MLNeuralNetworkEngine._pixelBufferFromEbufDescriptionError]
//   - [MLNeuralNetworkEngine._setMultiArrayOutputBackingForOutputFeatureNameToEbufError]
//   - [MLNeuralNetworkEngine._setMultipleBuffersOnPlanError]
//   - [MLNeuralNetworkEngine._setupContextAndPlanWithConfigurationPriorityError]
//   - [MLNeuralNetworkEngine._setupContextAndPlanWithConfigurationUsingCPUPriorityError]
//   - [MLNeuralNetworkEngine._setupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError]
//   - [MLNeuralNetworkEngine._setupContextAndPlanWithForceCPUPriorityError]
//   - [MLNeuralNetworkEngine.ActiveFunction]
//   - [MLNeuralNetworkEngine.SetActiveFunction]
//   - [MLNeuralNetworkEngine.AddClassifierInformationToOutputOptionsError]
//   - [MLNeuralNetworkEngine.AvailableOutputBlobList]
//   - [MLNeuralNetworkEngine.BindDirectlyInputFeatureNamedPixelBufferCleanUpBlocksBoundDirectlyError]
//   - [MLNeuralNetworkEngine.BindDynamicOutputBuffersError]
//   - [MLNeuralNetworkEngine.BindInputFeatureNamedConvertingMultiArrayBufferIndexError]
//   - [MLNeuralNetworkEngine.BindInputFeatureNamedFeatureValueBufferIndexCleanUpBlocksBoundDirectlyError]
//   - [MLNeuralNetworkEngine.BindInputFeatureNamedPixelBufferCleanUpBlocksError]
//   - [MLNeuralNetworkEngine.BindInputFeaturesBufferIndexCleanUpBlocksDirectlyBoundFeatureNamesError]
//   - [MLNeuralNetworkEngine.BindInputsAndOutputsCleanUpBlocksBufferIndexOptionsError]
//   - [MLNeuralNetworkEngine.BindOutputBuffersOutputBackingsAutomaticOutputBackingModeDirectlyBoundOutputFeatureNamesError]
//   - [MLNeuralNetworkEngine.BufferSemaphore]
//   - [MLNeuralNetworkEngine.SetBufferSemaphore]
//   - [MLNeuralNetworkEngine.ClassLabels]
//   - [MLNeuralNetworkEngine.SetClassLabels]
//   - [MLNeuralNetworkEngine.ClassScoreVectorName]
//   - [MLNeuralNetworkEngine.SetClassScoreVectorName]
//   - [MLNeuralNetworkEngine.ClassifyOptionsError]
//   - [MLNeuralNetworkEngine.CollectParametersFromContainerConfigurationError]
//   - [MLNeuralNetworkEngine.CompilerVersionInfo]
//   - [MLNeuralNetworkEngine.CompleteOutputBackingsAutomaticOutputBackingModeError]
//   - [MLNeuralNetworkEngine.Context]
//   - [MLNeuralNetworkEngine.SetContext]
//   - [MLNeuralNetworkEngine.ConvertPredictionToClassifierResultWithOptionsError]
//   - [MLNeuralNetworkEngine.CopyEbufOfPixelTypeToPixelBufferError]
//   - [MLNeuralNetworkEngine.CopyImagePreprocessingParametersToError]
//   - [MLNeuralNetworkEngine.CopyPixelBufferByApplyingImagePreprocessingToPixelBuffer]
//   - [MLNeuralNetworkEngine.CopyPixelBufferByApplyingImagePreprocessingForFeatureNamedToPixelBuffer]
//   - [MLNeuralNetworkEngine.CopyPixelBufferFromPixelBufferUsingPixelFormat]
//   - [MLNeuralNetworkEngine.DefaultOptionalValues]
//   - [MLNeuralNetworkEngine.SetDefaultOptionalValues]
//   - [MLNeuralNetworkEngine.DumpTestVectorsToPath]
//   - [MLNeuralNetworkEngine.Engine]
//   - [MLNeuralNetworkEngine.SetEngine]
//   - [MLNeuralNetworkEngine.EspressoInputShapes]
//   - [MLNeuralNetworkEngine.SetEspressoInputShapes]
//   - [MLNeuralNetworkEngine.EspressoInputStrides]
//   - [MLNeuralNetworkEngine.SetEspressoInputStrides]
//   - [MLNeuralNetworkEngine.EspressoProfileInfo]
//   - [MLNeuralNetworkEngine.SetEspressoProfileInfo]
//   - [MLNeuralNetworkEngine.EspressoQueue]
//   - [MLNeuralNetworkEngine.SetEspressoQueue]
//   - [MLNeuralNetworkEngine.EvaluateError]
//   - [MLNeuralNetworkEngine.EvaluateBatchOptionsError]
//   - [MLNeuralNetworkEngine.EvaluateInputsBufferIndexOptionsError]
//   - [MLNeuralNetworkEngine.EvaluateInputsOptionsError]
//   - [MLNeuralNetworkEngine.EvaluateInputsOptionsVerifyInputsError]
//   - [MLNeuralNetworkEngine.ExecutePlanError]
//   - [MLNeuralNetworkEngine.HardwareFallbackDetected]
//   - [MLNeuralNetworkEngine.SetHardwareFallbackDetected]
//   - [MLNeuralNetworkEngine.HasBidirectionalLayer]
//   - [MLNeuralNetworkEngine.SetHasBidirectionalLayer]
//   - [MLNeuralNetworkEngine.HasOptionalInputSequenceConcat]
//   - [MLNeuralNetworkEngine.SetHasOptionalInputSequenceConcat]
//   - [MLNeuralNetworkEngine.ImageFeatureValueFromEbufBackingCVPixelBufferDescriptionError]
//   - [MLNeuralNetworkEngine.ImageFeatureValueFromPixelBufferUsingPixelFormat]
//   - [MLNeuralNetworkEngine.ImagePreprocessingParameters]
//   - [MLNeuralNetworkEngine.SetImagePreprocessingParameters]
//   - [MLNeuralNetworkEngine.InputBindStateForFeatureValueError]
//   - [MLNeuralNetworkEngine.InputBlobNameToLastBackingMode]
//   - [MLNeuralNetworkEngine.InputFeatureConformer]
//   - [MLNeuralNetworkEngine.InputLayers]
//   - [MLNeuralNetworkEngine.IsANEPathForbidden]
//   - [MLNeuralNetworkEngine.SetIsANEPathForbidden]
//   - [MLNeuralNetworkEngine.IsEspressoBiasPreprocessingShared]
//   - [MLNeuralNetworkEngine.SetIsEspressoBiasPreprocessingShared]
//   - [MLNeuralNetworkEngine.IsGPUPathForbidden]
//   - [MLNeuralNetworkEngine.SetIsGPUPathForbidden]
//   - [MLNeuralNetworkEngine.LockPixelBufferCleanUpBlocksError]
//   - [MLNeuralNetworkEngine.ModelFilePath]
//   - [MLNeuralNetworkEngine.SetModelFilePath]
//   - [MLNeuralNetworkEngine.ModelIsEncrypted]
//   - [MLNeuralNetworkEngine.ModelIsMIL]
//   - [MLNeuralNetworkEngine.SetModelIsMIL]
//   - [MLNeuralNetworkEngine.ModelVersionInfo]
//   - [MLNeuralNetworkEngine.MultiArrayFeatureValueFromEbufBackingMultiArrayDescriptionOutputNameError]
//   - [MLNeuralNetworkEngine.NdArrayInterpretation]
//   - [MLNeuralNetworkEngine.SetNdArrayInterpretation]
//   - [MLNeuralNetworkEngine.Network]
//   - [MLNeuralNetworkEngine.SetNetwork]
//   - [MLNeuralNetworkEngine.NumInputs]
//   - [MLNeuralNetworkEngine.NumOutputs]
//   - [MLNeuralNetworkEngine.ObtainBuffer]
//   - [MLNeuralNetworkEngine.OpacifyAndPermutePixelBufferBufferContainsBGRAError]
//   - [MLNeuralNetworkEngine.OptionalInputTypes]
//   - [MLNeuralNetworkEngine.OutputBackingMultiArrayForFeatureName]
//   - [MLNeuralNetworkEngine.OutputBlobNameToLastBackingMode]
//   - [MLNeuralNetworkEngine.OutputLayers]
//   - [MLNeuralNetworkEngine.PixelBufferBackedMultiArrayWithShape]
//   - [MLNeuralNetworkEngine.PixelBufferFromOutputBackingForFeature]
//   - [MLNeuralNetworkEngine.PixelBufferPool]
//   - [MLNeuralNetworkEngine.Plan]
//   - [MLNeuralNetworkEngine.SetPlan]
//   - [MLNeuralNetworkEngine.PopulateMultiArrayShapeStridesForEbufFeatureDescriptionNdArrayInterpretation]
//   - [MLNeuralNetworkEngine.PopulateOutputsOutputBackingsDirectlyBoundOutputFeatureNamesError]
//   - [MLNeuralNetworkEngine.Precision]
//   - [MLNeuralNetworkEngine.SetPrecision]
//   - [MLNeuralNetworkEngine.PredictionsQueue]
//   - [MLNeuralNetworkEngine.SetPredictionsQueue]
//   - [MLNeuralNetworkEngine.PrepareBlobNamedForNewBlobBackingModeBindMode]
//   - [MLNeuralNetworkEngine.Priority]
//   - [MLNeuralNetworkEngine.SetPriority]
//   - [MLNeuralNetworkEngine.ProbabilityDictionarySharedKeySet]
//   - [MLNeuralNetworkEngine.SetProbabilityDictionarySharedKeySet]
//   - [MLNeuralNetworkEngine.Qos]
//   - [MLNeuralNetworkEngine.SetQos]
//   - [MLNeuralNetworkEngine.RebuildPlan]
//   - [MLNeuralNetworkEngine.RebuildPlanError]
//   - [MLNeuralNetworkEngine.RegressOptionsError]
//   - [MLNeuralNetworkEngine.ReleaseBuffer]
//   - [MLNeuralNetworkEngine.ResetSizesError]
//   - [MLNeuralNetworkEngine.ResetSizesNoAutoReleaseError]
//   - [MLNeuralNetworkEngine.ResetSizesWithEspressoConfigurationsError]
//   - [MLNeuralNetworkEngine.SequenceConcatConsumesOptionalInputNamed]
//   - [MLNeuralNetworkEngine.SequenceNamed]
//   - [MLNeuralNetworkEngine.SetEspressoBlobShapesWidthsHeightsKsBatchesSequencesRanksError]
//   - [MLNeuralNetworkEngine.SortBatchByShapeWithMapError]
//   - [MLNeuralNetworkEngine.SubmitSemaphore]
//   - [MLNeuralNetworkEngine.SetSubmitSemaphore]
//   - [MLNeuralNetworkEngine.SupportFromEspressoLayerInfo]
//   - [MLNeuralNetworkEngine.SupportFromEspressoPlatform]
//   - [MLNeuralNetworkEngine.TransferOneComponent16HalfPixelBufferToPixelBufferWithScaleBias]
//   - [MLNeuralNetworkEngine.TransferPixelBufferToPixelBuffer]
//   - [MLNeuralNetworkEngine.TryToSetOutputBackingForFeatureNameToEbufReportPointerFlagsError]
//   - [MLNeuralNetworkEngine.UpdateDynamicOutputBlobIndicatorCacheAndReturnError]
//   - [MLNeuralNetworkEngine.UsingCPU]
//   - [MLNeuralNetworkEngine.SetUsingCPU]
//   - [MLNeuralNetworkEngine.UsingEspressoConfigurations]
//   - [MLNeuralNetworkEngine.VerifyInputsError]
//   - [MLNeuralNetworkEngine.InitWithContainerConfigurationError]
//   - [MLNeuralNetworkEngine.InitWithContainerError]
type MLNeuralNetworkEngine struct {
	MLModelEngine
}

// MLNeuralNetworkEngineFromID constructs a [MLNeuralNetworkEngine] from an objc.ID.
func MLNeuralNetworkEngineFromID(id objc.ID) MLNeuralNetworkEngine {
	return MLNeuralNetworkEngine{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLNeuralNetworkEngine implements IMLNeuralNetworkEngine.
var _ IMLNeuralNetworkEngine = MLNeuralNetworkEngine{}

// An interface definition for the [MLNeuralNetworkEngine] class.
//
// # Methods
//
//   - [IMLNeuralNetworkEngine._addCompiledNetworkOrProgramToPlanError]
//   - [IMLNeuralNetworkEngine._addNetworkToPlanError]
//   - [IMLNeuralNetworkEngine._deallocContextAndPlan]
//   - [IMLNeuralNetworkEngine._espressoDeviceForConfigurationError]
//   - [IMLNeuralNetworkEngine._espressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray]
//   - [IMLNeuralNetworkEngine._handleAddNetworkToPlanStatusError]
//   - [IMLNeuralNetworkEngine._matchEngineToOptionsError]
//   - [IMLNeuralNetworkEngine._pixelBufferFromEbufDescriptionError]
//   - [IMLNeuralNetworkEngine._setMultiArrayOutputBackingForOutputFeatureNameToEbufError]
//   - [IMLNeuralNetworkEngine._setMultipleBuffersOnPlanError]
//   - [IMLNeuralNetworkEngine._setupContextAndPlanWithConfigurationPriorityError]
//   - [IMLNeuralNetworkEngine._setupContextAndPlanWithConfigurationUsingCPUPriorityError]
//   - [IMLNeuralNetworkEngine._setupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError]
//   - [IMLNeuralNetworkEngine._setupContextAndPlanWithForceCPUPriorityError]
//   - [IMLNeuralNetworkEngine.ActiveFunction]
//   - [IMLNeuralNetworkEngine.SetActiveFunction]
//   - [IMLNeuralNetworkEngine.AddClassifierInformationToOutputOptionsError]
//   - [IMLNeuralNetworkEngine.AvailableOutputBlobList]
//   - [IMLNeuralNetworkEngine.BindDirectlyInputFeatureNamedPixelBufferCleanUpBlocksBoundDirectlyError]
//   - [IMLNeuralNetworkEngine.BindDynamicOutputBuffersError]
//   - [IMLNeuralNetworkEngine.BindInputFeatureNamedConvertingMultiArrayBufferIndexError]
//   - [IMLNeuralNetworkEngine.BindInputFeatureNamedFeatureValueBufferIndexCleanUpBlocksBoundDirectlyError]
//   - [IMLNeuralNetworkEngine.BindInputFeatureNamedPixelBufferCleanUpBlocksError]
//   - [IMLNeuralNetworkEngine.BindInputFeaturesBufferIndexCleanUpBlocksDirectlyBoundFeatureNamesError]
//   - [IMLNeuralNetworkEngine.BindInputsAndOutputsCleanUpBlocksBufferIndexOptionsError]
//   - [IMLNeuralNetworkEngine.BindOutputBuffersOutputBackingsAutomaticOutputBackingModeDirectlyBoundOutputFeatureNamesError]
//   - [IMLNeuralNetworkEngine.BufferSemaphore]
//   - [IMLNeuralNetworkEngine.SetBufferSemaphore]
//   - [IMLNeuralNetworkEngine.ClassLabels]
//   - [IMLNeuralNetworkEngine.SetClassLabels]
//   - [IMLNeuralNetworkEngine.ClassScoreVectorName]
//   - [IMLNeuralNetworkEngine.SetClassScoreVectorName]
//   - [IMLNeuralNetworkEngine.ClassifyOptionsError]
//   - [IMLNeuralNetworkEngine.CollectParametersFromContainerConfigurationError]
//   - [IMLNeuralNetworkEngine.CompilerVersionInfo]
//   - [IMLNeuralNetworkEngine.CompleteOutputBackingsAutomaticOutputBackingModeError]
//   - [IMLNeuralNetworkEngine.Context]
//   - [IMLNeuralNetworkEngine.SetContext]
//   - [IMLNeuralNetworkEngine.ConvertPredictionToClassifierResultWithOptionsError]
//   - [IMLNeuralNetworkEngine.CopyEbufOfPixelTypeToPixelBufferError]
//   - [IMLNeuralNetworkEngine.CopyImagePreprocessingParametersToError]
//   - [IMLNeuralNetworkEngine.CopyPixelBufferByApplyingImagePreprocessingToPixelBuffer]
//   - [IMLNeuralNetworkEngine.CopyPixelBufferByApplyingImagePreprocessingForFeatureNamedToPixelBuffer]
//   - [IMLNeuralNetworkEngine.CopyPixelBufferFromPixelBufferUsingPixelFormat]
//   - [IMLNeuralNetworkEngine.DefaultOptionalValues]
//   - [IMLNeuralNetworkEngine.SetDefaultOptionalValues]
//   - [IMLNeuralNetworkEngine.DumpTestVectorsToPath]
//   - [IMLNeuralNetworkEngine.Engine]
//   - [IMLNeuralNetworkEngine.SetEngine]
//   - [IMLNeuralNetworkEngine.EspressoInputShapes]
//   - [IMLNeuralNetworkEngine.SetEspressoInputShapes]
//   - [IMLNeuralNetworkEngine.EspressoInputStrides]
//   - [IMLNeuralNetworkEngine.SetEspressoInputStrides]
//   - [IMLNeuralNetworkEngine.EspressoProfileInfo]
//   - [IMLNeuralNetworkEngine.SetEspressoProfileInfo]
//   - [IMLNeuralNetworkEngine.EspressoQueue]
//   - [IMLNeuralNetworkEngine.SetEspressoQueue]
//   - [IMLNeuralNetworkEngine.EvaluateError]
//   - [IMLNeuralNetworkEngine.EvaluateBatchOptionsError]
//   - [IMLNeuralNetworkEngine.EvaluateInputsBufferIndexOptionsError]
//   - [IMLNeuralNetworkEngine.EvaluateInputsOptionsError]
//   - [IMLNeuralNetworkEngine.EvaluateInputsOptionsVerifyInputsError]
//   - [IMLNeuralNetworkEngine.ExecutePlanError]
//   - [IMLNeuralNetworkEngine.HardwareFallbackDetected]
//   - [IMLNeuralNetworkEngine.SetHardwareFallbackDetected]
//   - [IMLNeuralNetworkEngine.HasBidirectionalLayer]
//   - [IMLNeuralNetworkEngine.SetHasBidirectionalLayer]
//   - [IMLNeuralNetworkEngine.HasOptionalInputSequenceConcat]
//   - [IMLNeuralNetworkEngine.SetHasOptionalInputSequenceConcat]
//   - [IMLNeuralNetworkEngine.ImageFeatureValueFromEbufBackingCVPixelBufferDescriptionError]
//   - [IMLNeuralNetworkEngine.ImageFeatureValueFromPixelBufferUsingPixelFormat]
//   - [IMLNeuralNetworkEngine.ImagePreprocessingParameters]
//   - [IMLNeuralNetworkEngine.SetImagePreprocessingParameters]
//   - [IMLNeuralNetworkEngine.InputBindStateForFeatureValueError]
//   - [IMLNeuralNetworkEngine.InputBlobNameToLastBackingMode]
//   - [IMLNeuralNetworkEngine.InputFeatureConformer]
//   - [IMLNeuralNetworkEngine.InputLayers]
//   - [IMLNeuralNetworkEngine.IsANEPathForbidden]
//   - [IMLNeuralNetworkEngine.SetIsANEPathForbidden]
//   - [IMLNeuralNetworkEngine.IsEspressoBiasPreprocessingShared]
//   - [IMLNeuralNetworkEngine.SetIsEspressoBiasPreprocessingShared]
//   - [IMLNeuralNetworkEngine.IsGPUPathForbidden]
//   - [IMLNeuralNetworkEngine.SetIsGPUPathForbidden]
//   - [IMLNeuralNetworkEngine.LockPixelBufferCleanUpBlocksError]
//   - [IMLNeuralNetworkEngine.ModelFilePath]
//   - [IMLNeuralNetworkEngine.SetModelFilePath]
//   - [IMLNeuralNetworkEngine.ModelIsEncrypted]
//   - [IMLNeuralNetworkEngine.ModelIsMIL]
//   - [IMLNeuralNetworkEngine.SetModelIsMIL]
//   - [IMLNeuralNetworkEngine.ModelVersionInfo]
//   - [IMLNeuralNetworkEngine.MultiArrayFeatureValueFromEbufBackingMultiArrayDescriptionOutputNameError]
//   - [IMLNeuralNetworkEngine.NdArrayInterpretation]
//   - [IMLNeuralNetworkEngine.SetNdArrayInterpretation]
//   - [IMLNeuralNetworkEngine.Network]
//   - [IMLNeuralNetworkEngine.SetNetwork]
//   - [IMLNeuralNetworkEngine.NumInputs]
//   - [IMLNeuralNetworkEngine.NumOutputs]
//   - [IMLNeuralNetworkEngine.ObtainBuffer]
//   - [IMLNeuralNetworkEngine.OpacifyAndPermutePixelBufferBufferContainsBGRAError]
//   - [IMLNeuralNetworkEngine.OptionalInputTypes]
//   - [IMLNeuralNetworkEngine.OutputBackingMultiArrayForFeatureName]
//   - [IMLNeuralNetworkEngine.OutputBlobNameToLastBackingMode]
//   - [IMLNeuralNetworkEngine.OutputLayers]
//   - [IMLNeuralNetworkEngine.PixelBufferBackedMultiArrayWithShape]
//   - [IMLNeuralNetworkEngine.PixelBufferFromOutputBackingForFeature]
//   - [IMLNeuralNetworkEngine.PixelBufferPool]
//   - [IMLNeuralNetworkEngine.Plan]
//   - [IMLNeuralNetworkEngine.SetPlan]
//   - [IMLNeuralNetworkEngine.PopulateMultiArrayShapeStridesForEbufFeatureDescriptionNdArrayInterpretation]
//   - [IMLNeuralNetworkEngine.PopulateOutputsOutputBackingsDirectlyBoundOutputFeatureNamesError]
//   - [IMLNeuralNetworkEngine.Precision]
//   - [IMLNeuralNetworkEngine.SetPrecision]
//   - [IMLNeuralNetworkEngine.PredictionsQueue]
//   - [IMLNeuralNetworkEngine.SetPredictionsQueue]
//   - [IMLNeuralNetworkEngine.PrepareBlobNamedForNewBlobBackingModeBindMode]
//   - [IMLNeuralNetworkEngine.Priority]
//   - [IMLNeuralNetworkEngine.SetPriority]
//   - [IMLNeuralNetworkEngine.ProbabilityDictionarySharedKeySet]
//   - [IMLNeuralNetworkEngine.SetProbabilityDictionarySharedKeySet]
//   - [IMLNeuralNetworkEngine.Qos]
//   - [IMLNeuralNetworkEngine.SetQos]
//   - [IMLNeuralNetworkEngine.RebuildPlan]
//   - [IMLNeuralNetworkEngine.RebuildPlanError]
//   - [IMLNeuralNetworkEngine.RegressOptionsError]
//   - [IMLNeuralNetworkEngine.ReleaseBuffer]
//   - [IMLNeuralNetworkEngine.ResetSizesError]
//   - [IMLNeuralNetworkEngine.ResetSizesNoAutoReleaseError]
//   - [IMLNeuralNetworkEngine.ResetSizesWithEspressoConfigurationsError]
//   - [IMLNeuralNetworkEngine.SequenceConcatConsumesOptionalInputNamed]
//   - [IMLNeuralNetworkEngine.SequenceNamed]
//   - [IMLNeuralNetworkEngine.SetEspressoBlobShapesWidthsHeightsKsBatchesSequencesRanksError]
//   - [IMLNeuralNetworkEngine.SortBatchByShapeWithMapError]
//   - [IMLNeuralNetworkEngine.SubmitSemaphore]
//   - [IMLNeuralNetworkEngine.SetSubmitSemaphore]
//   - [IMLNeuralNetworkEngine.SupportFromEspressoLayerInfo]
//   - [IMLNeuralNetworkEngine.SupportFromEspressoPlatform]
//   - [IMLNeuralNetworkEngine.TransferOneComponent16HalfPixelBufferToPixelBufferWithScaleBias]
//   - [IMLNeuralNetworkEngine.TransferPixelBufferToPixelBuffer]
//   - [IMLNeuralNetworkEngine.TryToSetOutputBackingForFeatureNameToEbufReportPointerFlagsError]
//   - [IMLNeuralNetworkEngine.UpdateDynamicOutputBlobIndicatorCacheAndReturnError]
//   - [IMLNeuralNetworkEngine.UsingCPU]
//   - [IMLNeuralNetworkEngine.SetUsingCPU]
//   - [IMLNeuralNetworkEngine.UsingEspressoConfigurations]
//   - [IMLNeuralNetworkEngine.VerifyInputsError]
//   - [IMLNeuralNetworkEngine.InitWithContainerConfigurationError]
//   - [IMLNeuralNetworkEngine.InitWithContainerError]
type IMLNeuralNetworkEngine interface {
	IMLModelEngine

	// Topic: Methods

	_addCompiledNetworkOrProgramToPlanError(plan unsafe.Pointer) (bool, error)
	_addNetworkToPlanError(plan unsafe.Pointer) (bool, error)
	_deallocContextAndPlan()
	_espressoDeviceForConfigurationError(configuration objectivec.IObject) (int, error)
	_espressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray(name objectivec.IObject, array objectivec.IObject) bool
	_handleAddNetworkToPlanStatusError(status int) (bool, error)
	_matchEngineToOptionsError(options objectivec.IObject) (bool, error)
	_pixelBufferFromEbufDescriptionError(ebuf unsafe.Pointer, description objectivec.IObject) (corevideo.CVImageBufferRef, error)
	_setMultiArrayOutputBackingForOutputFeatureNameToEbufError(backing objectivec.IObject, name objectivec.IObject, ebuf unsafe.Pointer) (bool, error)
	_setMultipleBuffersOnPlanError(plan unsafe.Pointer) (bool, error)
	_setupContextAndPlanWithConfigurationPriorityError(configuration objectivec.IObject, priority int) (bool, error)
	_setupContextAndPlanWithConfigurationUsingCPUPriorityError(configuration objectivec.IObject, cpu bool, priority int) (bool, error)
	_setupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError(configuration objectivec.IObject, cpu bool, priority int, container bool) (bool, error)
	_setupContextAndPlanWithForceCPUPriorityError(cpu bool, priority int) (bool, error)
	ActiveFunction() string
	SetActiveFunction(value string)
	AddClassifierInformationToOutputOptionsError(output objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	AvailableOutputBlobList() objectivec.IObject
	BindDirectlyInputFeatureNamedPixelBufferCleanUpBlocksBoundDirectlyError(named objectivec.IObject, buffer corevideo.CVImageBufferRef, blocks objectivec.IObject) (bool, error)
	BindDynamicOutputBuffersError(buffers unsafe.Pointer) (bool, error)
	BindInputFeatureNamedConvertingMultiArrayBufferIndexError(named objectivec.IObject, array objectivec.IObject, index uint64) (bool, error)
	BindInputFeatureNamedFeatureValueBufferIndexCleanUpBlocksBoundDirectlyError(named objectivec.IObject, value objectivec.IObject, index uint64, blocks objectivec.IObject) (bool, error)
	BindInputFeatureNamedPixelBufferCleanUpBlocksError(named objectivec.IObject, buffer corevideo.CVImageBufferRef, blocks objectivec.IObject) (bool, error)
	BindInputFeaturesBufferIndexCleanUpBlocksDirectlyBoundFeatureNamesError(features objectivec.IObject, index uint64, blocks objectivec.IObject, names []objectivec.IObject) (bool, error)
	BindInputsAndOutputsCleanUpBlocksBufferIndexOptionsError(outputs objectivec.IObject, blocks objectivec.IObject, index uint64, options objectivec.IObject) (bool, error)
	BindOutputBuffersOutputBackingsAutomaticOutputBackingModeDirectlyBoundOutputFeatureNamesError(buffers unsafe.Pointer, backings objectivec.IObject, mode objectivec.IObject, names []objectivec.IObject) (bool, error)
	BufferSemaphore() objectivec.Object
	SetBufferSemaphore(value objectivec.Object)
	ClassLabels() foundation.INSArray
	SetClassLabels(value foundation.INSArray)
	ClassScoreVectorName() string
	SetClassScoreVectorName(value string)
	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	CollectParametersFromContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (bool, error)
	CompilerVersionInfo() IMLVersionInfo
	CompleteOutputBackingsAutomaticOutputBackingModeError(backings objectivec.IObject, mode objectivec.IObject) (objectivec.IObject, error)
	Context() unsafe.Pointer
	SetContext(value unsafe.Pointer)
	ConvertPredictionToClassifierResultWithOptionsError(result objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	CopyEbufOfPixelTypeToPixelBufferError(ebuf unsafe.Pointer, type_ uint64, buffer corevideo.CVImageBufferRef) (bool, error)
	CopyImagePreprocessingParametersToError(to unsafe.Pointer) (bool, error)
	CopyPixelBufferByApplyingImagePreprocessingToPixelBuffer(preprocessing unsafe.Pointer, buffer corevideo.CVImageBufferRef) corevideo.CVImageBufferRef
	CopyPixelBufferByApplyingImagePreprocessingForFeatureNamedToPixelBuffer(named objectivec.IObject, buffer corevideo.CVImageBufferRef) corevideo.CVImageBufferRef
	CopyPixelBufferFromPixelBufferUsingPixelFormat(buffer corevideo.CVImageBufferRef, format uint32) corevideo.CVImageBufferRef
	DefaultOptionalValues() foundation.INSDictionary
	SetDefaultOptionalValues(value foundation.INSDictionary)
	DumpTestVectorsToPath(path objectivec.IObject)
	Engine() int
	SetEngine(value int)
	EspressoInputShapes() foundation.INSDictionary
	SetEspressoInputShapes(value foundation.INSDictionary)
	EspressoInputStrides() foundation.INSDictionary
	SetEspressoInputStrides(value foundation.INSDictionary)
	EspressoProfileInfo() objectivec.IObject
	SetEspressoProfileInfo(value objectivec.IObject)
	EspressoQueue() objectivec.Object
	SetEspressoQueue(value objectivec.Object)
	EvaluateError(evaluate objectivec.IObject) (objectivec.IObject, error)
	EvaluateBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	EvaluateInputsBufferIndexOptionsError(inputs objectivec.IObject, index uint64, options objectivec.IObject) (objectivec.IObject, error)
	EvaluateInputsOptionsError(inputs objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	EvaluateInputsOptionsVerifyInputsError(inputs objectivec.IObject, options objectivec.IObject, inputs2 bool) (objectivec.IObject, error)
	ExecutePlanError(plan unsafe.Pointer) (bool, error)
	HardwareFallbackDetected() bool
	SetHardwareFallbackDetected(value bool)
	HasBidirectionalLayer() bool
	SetHasBidirectionalLayer(value bool)
	HasOptionalInputSequenceConcat() bool
	SetHasOptionalInputSequenceConcat(value bool)
	ImageFeatureValueFromEbufBackingCVPixelBufferDescriptionError(ebuf unsafe.Pointer, buffer corevideo.CVImageBufferRef, description objectivec.IObject) (objectivec.IObject, error)
	ImageFeatureValueFromPixelBufferUsingPixelFormat(buffer corevideo.CVImageBufferRef, format uint32) objectivec.IObject
	ImagePreprocessingParameters() foundation.INSDictionary
	SetImagePreprocessingParameters(value foundation.INSDictionary)
	InputBindStateForFeatureValueError(value objectivec.IObject) (int64, error)
	InputBlobNameToLastBackingMode() foundation.INSDictionary
	InputFeatureConformer() IMLFeatureProviderConformer
	InputLayers() foundation.INSArray
	IsANEPathForbidden() bool
	SetIsANEPathForbidden(value bool)
	IsEspressoBiasPreprocessingShared() bool
	SetIsEspressoBiasPreprocessingShared(value bool)
	IsGPUPathForbidden() bool
	SetIsGPUPathForbidden(value bool)
	LockPixelBufferCleanUpBlocksError(buffer corevideo.CVImageBufferRef, blocks objectivec.IObject) (bool, error)
	ModelFilePath() string
	SetModelFilePath(value string)
	ModelIsEncrypted() bool
	ModelIsMIL() bool
	SetModelIsMIL(value bool)
	ModelVersionInfo() IMLVersionInfo
	MultiArrayFeatureValueFromEbufBackingMultiArrayDescriptionOutputNameError(ebuf unsafe.Pointer, array objectivec.IObject, description objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error)
	NdArrayInterpretation() bool
	SetNdArrayInterpretation(value bool)
	Network() unsafe.Pointer
	SetNetwork(value unsafe.Pointer)
	NumInputs() uint64
	NumOutputs() uint64
	ObtainBuffer() uint64
	OpacifyAndPermutePixelBufferBufferContainsBGRAError(buffer corevideo.CVImageBufferRef, bgra bool) (bool, error)
	OptionalInputTypes() foundation.INSDictionary
	OutputBackingMultiArrayForFeatureName(name objectivec.IObject) objectivec.IObject
	OutputBlobNameToLastBackingMode() foundation.INSDictionary
	OutputLayers() foundation.INSArray
	PixelBufferBackedMultiArrayWithShape(shape objectivec.IObject) objectivec.IObject
	PixelBufferFromOutputBackingForFeature(backing objectivec.IObject, feature objectivec.IObject) corevideo.CVImageBufferRef
	PixelBufferPool() IMLPixelBufferPool
	Plan() unsafe.Pointer
	SetPlan(value unsafe.Pointer)
	PopulateMultiArrayShapeStridesForEbufFeatureDescriptionNdArrayInterpretation(shape []objectivec.IObject, strides []objectivec.IObject, ebuf unsafe.Pointer, description objectivec.IObject, interpretation bool)
	PopulateOutputsOutputBackingsDirectlyBoundOutputFeatureNamesError(outputs uint64, backings objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error)
	Precision() int
	SetPrecision(value int)
	PredictionsQueue() objectivec.Object
	SetPredictionsQueue(value objectivec.Object)
	PrepareBlobNamedForNewBlobBackingModeBindMode(named objectivec.IObject, mode int64, mode2 int)
	Priority() int
	SetPriority(value int)
	ProbabilityDictionarySharedKeySet() objectivec.IObject
	SetProbabilityDictionarySharedKeySet(value objectivec.IObject)
	Qos() int
	SetQos(value int)
	RebuildPlan(plan []objectivec.IObject) bool
	RebuildPlanError(plan bool) (bool, error)
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ReleaseBuffer(buffer uint64)
	ResetSizesError(sizes objectivec.IObject) (bool, error)
	ResetSizesNoAutoReleaseError(release objectivec.IObject) (bool, error)
	ResetSizesWithEspressoConfigurationsError(configurations objectivec.IObject) (bool, error)
	SequenceConcatConsumesOptionalInputNamed(named objectivec.IObject) bool
	SequenceNamed(named objectivec.IObject) int
	SetEspressoBlobShapesWidthsHeightsKsBatchesSequencesRanksError(shapes unsafe.Pointer, widths unsafe.Pointer, heights unsafe.Pointer, ks unsafe.Pointer, batches unsafe.Pointer, sequences unsafe.Pointer, ranks unsafe.Pointer) (bool, error)
	SortBatchByShapeWithMapError(shape objectivec.IObject, map_ []objectivec.IObject) (objectivec.IObject, error)
	SubmitSemaphore() objectivec.Object
	SetSubmitSemaphore(value objectivec.Object)
	SupportFromEspressoLayerInfo(info objectivec.IObject) uint64
	SupportFromEspressoPlatform(platform int) uint64
	TransferOneComponent16HalfPixelBufferToPixelBufferWithScaleBias(buffer corevideo.CVImageBufferRef, buffer2 corevideo.CVImageBufferRef, scale float32, bias float32) bool
	TransferPixelBufferToPixelBuffer(buffer corevideo.CVImageBufferRef, buffer2 corevideo.CVImageBufferRef) bool
	TryToSetOutputBackingForFeatureNameToEbufReportPointerFlagsError(backing objectivec.IObject, name objectivec.IObject, ebuf unsafe.Pointer) (int, error)
	UpdateDynamicOutputBlobIndicatorCacheAndReturnError() (bool, error)
	UsingCPU() bool
	SetUsingCPU(value bool)
	UsingEspressoConfigurations() bool
	VerifyInputsError(inputs objectivec.IObject) (objectivec.IObject, error)
	InitWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkEngine, error)
	InitWithContainerError(container objectivec.IObject) (MLNeuralNetworkEngine, error)
}

// Init initializes the instance.
func (m MLNeuralNetworkEngine) Init() MLNeuralNetworkEngine {
	rv := objc.Send[MLNeuralNetworkEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworkEngine) Autorelease() MLNeuralNetworkEngine {
	rv := objc.Send[MLNeuralNetworkEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkEngine creates a new MLNeuralNetworkEngine instance.
func NewMLNeuralNetworkEngine() MLNeuralNetworkEngine {
	class := getMLNeuralNetworkEngineClass()
	rv := objc.Send[MLNeuralNetworkEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNeuralNetworkEngineWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkEngineFromID(rv), nil
}

func NewNeuralNetworkEngineWithContainerError(container objectivec.IObject) (MLNeuralNetworkEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkEngineFromID(rv), nil
}

func NewNeuralNetworkEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkEngine {
	instance := getMLNeuralNetworkEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkEngineFromID(rv)
}

func NewNeuralNetworkEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkEngine {
	instance := getMLNeuralNetworkEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkEngineFromID(rv)
}

func (m MLNeuralNetworkEngine) _addCompiledNetworkOrProgramToPlanError(plan unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_addCompiledNetworkOrProgramToPlan:error:"), plan, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_addCompiledNetworkOrProgramToPlan:error: returned NO with nil NSError")
	}
	return rv, nil

}

// AddCompiledNetworkOrProgramToPlanError is an exported wrapper for the private method _addCompiledNetworkOrProgramToPlanError.
func (m MLNeuralNetworkEngine) AddCompiledNetworkOrProgramToPlanError(plan unsafe.Pointer) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_addCompiledNetworkOrProgramToPlan:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_addCompiledNetworkOrProgramToPlan:error:"}
		return false, err
	}
	return m._addCompiledNetworkOrProgramToPlanError(plan)
}

// CanAddCompiledNetworkOrProgramToPlanError reports whether the receiver responds to the private selector _addCompiledNetworkOrProgramToPlan:error:.
func (m MLNeuralNetworkEngine) CanAddCompiledNetworkOrProgramToPlanError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_addCompiledNetworkOrProgramToPlan:error:"))
}
func (m MLNeuralNetworkEngine) _addNetworkToPlanError(plan unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_addNetworkToPlan:error:"), plan, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_addNetworkToPlan:error: returned NO with nil NSError")
	}
	return rv, nil

}

// AddNetworkToPlanError is an exported wrapper for the private method _addNetworkToPlanError.
func (m MLNeuralNetworkEngine) AddNetworkToPlanError(plan unsafe.Pointer) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_addNetworkToPlan:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_addNetworkToPlan:error:"}
		return false, err
	}
	return m._addNetworkToPlanError(plan)
}

// CanAddNetworkToPlanError reports whether the receiver responds to the private selector _addNetworkToPlan:error:.
func (m MLNeuralNetworkEngine) CanAddNetworkToPlanError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_addNetworkToPlan:error:"))
}
func (m MLNeuralNetworkEngine) _deallocContextAndPlan() {
	objc.Send[objc.ID](m.ID, objc.Sel("_deallocContextAndPlan"))
}

// DeallocContextAndPlan is an exported wrapper for the private method _deallocContextAndPlan.
func (m MLNeuralNetworkEngine) DeallocContextAndPlan() error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_deallocContextAndPlan")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_deallocContextAndPlan"}
		return err
	}
	m._deallocContextAndPlan()
	return nil
}

// CanDeallocContextAndPlan reports whether the receiver responds to the private selector _deallocContextAndPlan.
func (m MLNeuralNetworkEngine) CanDeallocContextAndPlan() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_deallocContextAndPlan"))
}
func (m MLNeuralNetworkEngine) _espressoDeviceForConfigurationError(configuration objectivec.IObject) (int, error) {
	var errorPtr objc.ID
	rv := objc.Send[int](m.ID, objc.Sel("_espressoDeviceForConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// EspressoDeviceForConfigurationError is an exported wrapper for the private method _espressoDeviceForConfigurationError.
func (m MLNeuralNetworkEngine) EspressoDeviceForConfigurationError(configuration objectivec.IObject) (int, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_espressoDeviceForConfiguration:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_espressoDeviceForConfiguration:error:"}
		return 0, err
	}
	return m._espressoDeviceForConfigurationError(configuration)
}

// CanEspressoDeviceForConfigurationError reports whether the receiver responds to the private selector _espressoDeviceForConfiguration:error:.
func (m MLNeuralNetworkEngine) CanEspressoDeviceForConfigurationError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_espressoDeviceForConfiguration:error:"))
}
func (m MLNeuralNetworkEngine) _espressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray(name objectivec.IObject, array objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_espressoOutputShapeForFeatureName:matchesShapeOfMLMultiArray:"), name, array)
	return rv
}

// EspressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray is an exported wrapper for the private method _espressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray.
func (m MLNeuralNetworkEngine) EspressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray(name objectivec.IObject, array objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_espressoOutputShapeForFeatureName:matchesShapeOfMLMultiArray:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_espressoOutputShapeForFeatureName:matchesShapeOfMLMultiArray:"}
		return false, err
	}
	return m._espressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray(name, array), nil
}

// CanEspressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray reports whether the receiver responds to the private selector _espressoOutputShapeForFeatureName:matchesShapeOfMLMultiArray:.
func (m MLNeuralNetworkEngine) CanEspressoOutputShapeForFeatureNameMatchesShapeOfMLMultiArray() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_espressoOutputShapeForFeatureName:matchesShapeOfMLMultiArray:"))
}
func (m MLNeuralNetworkEngine) _handleAddNetworkToPlanStatusError(status int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_handleAddNetworkToPlanStatus:error:"), status, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_handleAddNetworkToPlanStatus:error: returned NO with nil NSError")
	}
	return rv, nil

}

// HandleAddNetworkToPlanStatusError is an exported wrapper for the private method _handleAddNetworkToPlanStatusError.
func (m MLNeuralNetworkEngine) HandleAddNetworkToPlanStatusError(status int) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_handleAddNetworkToPlanStatus:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_handleAddNetworkToPlanStatus:error:"}
		return false, err
	}
	return m._handleAddNetworkToPlanStatusError(status)
}

// CanHandleAddNetworkToPlanStatusError reports whether the receiver responds to the private selector _handleAddNetworkToPlanStatus:error:.
func (m MLNeuralNetworkEngine) CanHandleAddNetworkToPlanStatusError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_handleAddNetworkToPlanStatus:error:"))
}
func (m MLNeuralNetworkEngine) _matchEngineToOptionsError(options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_matchEngineToOptions:error:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_matchEngineToOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}

// MatchEngineToOptionsError is an exported wrapper for the private method _matchEngineToOptionsError.
func (m MLNeuralNetworkEngine) MatchEngineToOptionsError(options objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_matchEngineToOptions:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_matchEngineToOptions:error:"}
		return false, err
	}
	return m._matchEngineToOptionsError(options)
}

// CanMatchEngineToOptionsError reports whether the receiver responds to the private selector _matchEngineToOptions:error:.
func (m MLNeuralNetworkEngine) CanMatchEngineToOptionsError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_matchEngineToOptions:error:"))
}
func (m MLNeuralNetworkEngine) _pixelBufferFromEbufDescriptionError(ebuf unsafe.Pointer, description objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("_pixelBufferFromEbuf:description:error:"), ebuf, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corevideo.CVImageBufferRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// PixelBufferFromEbufDescriptionError is an exported wrapper for the private method _pixelBufferFromEbufDescriptionError.
func (m MLNeuralNetworkEngine) PixelBufferFromEbufDescriptionError(ebuf unsafe.Pointer, description objectivec.IObject) (corevideo.CVImageBufferRef, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_pixelBufferFromEbuf:description:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_pixelBufferFromEbuf:description:error:"}
		return *new(corevideo.CVImageBufferRef), err
	}
	return m._pixelBufferFromEbufDescriptionError(ebuf, description)
}

// CanPixelBufferFromEbufDescriptionError reports whether the receiver responds to the private selector _pixelBufferFromEbuf:description:error:.
func (m MLNeuralNetworkEngine) CanPixelBufferFromEbufDescriptionError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_pixelBufferFromEbuf:description:error:"))
}
func (m MLNeuralNetworkEngine) _setMultiArrayOutputBackingForOutputFeatureNameToEbufError(backing objectivec.IObject, name objectivec.IObject, ebuf unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setMultiArrayOutputBacking:forOutputFeatureName:toEbuf:error:"), backing, name, ebuf, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setMultiArrayOutputBacking:forOutputFeatureName:toEbuf:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetMultiArrayOutputBackingForOutputFeatureNameToEbufError is an exported wrapper for the private method _setMultiArrayOutputBackingForOutputFeatureNameToEbufError.
func (m MLNeuralNetworkEngine) SetMultiArrayOutputBackingForOutputFeatureNameToEbufError(backing objectivec.IObject, name objectivec.IObject, ebuf unsafe.Pointer) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setMultiArrayOutputBacking:forOutputFeatureName:toEbuf:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setMultiArrayOutputBacking:forOutputFeatureName:toEbuf:error:"}
		return false, err
	}
	return m._setMultiArrayOutputBackingForOutputFeatureNameToEbufError(backing, name, ebuf)
}

// CanSetMultiArrayOutputBackingForOutputFeatureNameToEbufError reports whether the receiver responds to the private selector _setMultiArrayOutputBacking:forOutputFeatureName:toEbuf:error:.
func (m MLNeuralNetworkEngine) CanSetMultiArrayOutputBackingForOutputFeatureNameToEbufError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setMultiArrayOutputBacking:forOutputFeatureName:toEbuf:error:"))
}
func (m MLNeuralNetworkEngine) _setMultipleBuffersOnPlanError(plan unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setMultipleBuffersOnPlan:error:"), plan, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setMultipleBuffersOnPlan:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetMultipleBuffersOnPlanError is an exported wrapper for the private method _setMultipleBuffersOnPlanError.
func (m MLNeuralNetworkEngine) SetMultipleBuffersOnPlanError(plan unsafe.Pointer) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setMultipleBuffersOnPlan:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setMultipleBuffersOnPlan:error:"}
		return false, err
	}
	return m._setMultipleBuffersOnPlanError(plan)
}

// CanSetMultipleBuffersOnPlanError reports whether the receiver responds to the private selector _setMultipleBuffersOnPlan:error:.
func (m MLNeuralNetworkEngine) CanSetMultipleBuffersOnPlanError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setMultipleBuffersOnPlan:error:"))
}
func (m MLNeuralNetworkEngine) _setupContextAndPlanWithConfigurationPriorityError(configuration objectivec.IObject, priority int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:priority:error:"), configuration, priority, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setupContextAndPlanWithConfiguration:priority:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetupContextAndPlanWithConfigurationPriorityError is an exported wrapper for the private method _setupContextAndPlanWithConfigurationPriorityError.
func (m MLNeuralNetworkEngine) SetupContextAndPlanWithConfigurationPriorityError(configuration objectivec.IObject, priority int) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:priority:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupContextAndPlanWithConfiguration:priority:error:"}
		return false, err
	}
	return m._setupContextAndPlanWithConfigurationPriorityError(configuration, priority)
}

// CanSetupContextAndPlanWithConfigurationPriorityError reports whether the receiver responds to the private selector _setupContextAndPlanWithConfiguration:priority:error:.
func (m MLNeuralNetworkEngine) CanSetupContextAndPlanWithConfigurationPriorityError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:priority:error:"))
}
func (m MLNeuralNetworkEngine) _setupContextAndPlanWithConfigurationUsingCPUPriorityError(configuration objectivec.IObject, cpu bool, priority int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:usingCPU:priority:error:"), configuration, cpu, priority, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setupContextAndPlanWithConfiguration:usingCPU:priority:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetupContextAndPlanWithConfigurationUsingCPUPriorityError is an exported wrapper for the private method _setupContextAndPlanWithConfigurationUsingCPUPriorityError.
func (m MLNeuralNetworkEngine) SetupContextAndPlanWithConfigurationUsingCPUPriorityError(configuration objectivec.IObject, cpu bool, priority int) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:usingCPU:priority:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupContextAndPlanWithConfiguration:usingCPU:priority:error:"}
		return false, err
	}
	return m._setupContextAndPlanWithConfigurationUsingCPUPriorityError(configuration, cpu, priority)
}

// CanSetupContextAndPlanWithConfigurationUsingCPUPriorityError reports whether the receiver responds to the private selector _setupContextAndPlanWithConfiguration:usingCPU:priority:error:.
func (m MLNeuralNetworkEngine) CanSetupContextAndPlanWithConfigurationUsingCPUPriorityError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:usingCPU:priority:error:"))
}
func (m MLNeuralNetworkEngine) _setupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError(configuration objectivec.IObject, cpu bool, priority int, container bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:usingCPU:priority:reshapeWithContainer:error:"), configuration, cpu, priority, container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setupContextAndPlanWithConfiguration:usingCPU:priority:reshapeWithContainer:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError is an exported wrapper for the private method _setupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError.
func (m MLNeuralNetworkEngine) SetupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError(configuration objectivec.IObject, cpu bool, priority int, container bool) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:usingCPU:priority:reshapeWithContainer:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupContextAndPlanWithConfiguration:usingCPU:priority:reshapeWithContainer:error:"}
		return false, err
	}
	return m._setupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError(configuration, cpu, priority, container)
}

// CanSetupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError reports whether the receiver responds to the private selector _setupContextAndPlanWithConfiguration:usingCPU:priority:reshapeWithContainer:error:.
func (m MLNeuralNetworkEngine) CanSetupContextAndPlanWithConfigurationUsingCPUPriorityReshapeWithContainerError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithConfiguration:usingCPU:priority:reshapeWithContainer:error:"))
}
func (m MLNeuralNetworkEngine) _setupContextAndPlanWithForceCPUPriorityError(cpu bool, priority int) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_setupContextAndPlanWithForceCPU:priority:error:"), cpu, priority, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_setupContextAndPlanWithForceCPU:priority:error: returned NO with nil NSError")
	}
	return rv, nil

}

// SetupContextAndPlanWithForceCPUPriorityError is an exported wrapper for the private method _setupContextAndPlanWithForceCPUPriorityError.
func (m MLNeuralNetworkEngine) SetupContextAndPlanWithForceCPUPriorityError(cpu bool, priority int) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithForceCPU:priority:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupContextAndPlanWithForceCPU:priority:error:"}
		return false, err
	}
	return m._setupContextAndPlanWithForceCPUPriorityError(cpu, priority)
}

// CanSetupContextAndPlanWithForceCPUPriorityError reports whether the receiver responds to the private selector _setupContextAndPlanWithForceCPU:priority:error:.
func (m MLNeuralNetworkEngine) CanSetupContextAndPlanWithForceCPUPriorityError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_setupContextAndPlanWithForceCPU:priority:error:"))
}
func (m MLNeuralNetworkEngine) AddClassifierInformationToOutputOptionsError(output objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addClassifierInformationToOutput:options:error:"), output, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) AvailableOutputBlobList() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableOutputBlobList"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkEngine) BindDirectlyInputFeatureNamedPixelBufferCleanUpBlocksBoundDirectlyError(named objectivec.IObject, buffer corevideo.CVImageBufferRef, blocks objectivec.IObject) (bool, error) {
	var directly bool
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindDirectlyInputFeatureNamed:pixelBuffer:cleanUpBlocks:boundDirectly:error:"), named, buffer, blocks, unsafe.Pointer(&directly), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindDirectlyInputFeatureNamed:pixelBuffer:cleanUpBlocks:boundDirectly:error: returned NO with nil NSError")
	}
	return directly, nil
}
func (m MLNeuralNetworkEngine) BindDynamicOutputBuffersError(buffers unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindDynamicOutputBuffers:error:"), buffers, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindDynamicOutputBuffers:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) BindInputFeatureNamedConvertingMultiArrayBufferIndexError(named objectivec.IObject, array objectivec.IObject, index uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindInputFeatureNamed:convertingMultiArray:bufferIndex:error:"), named, array, index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindInputFeatureNamed:convertingMultiArray:bufferIndex:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) BindInputFeatureNamedFeatureValueBufferIndexCleanUpBlocksBoundDirectlyError(named objectivec.IObject, value objectivec.IObject, index uint64, blocks objectivec.IObject) (bool, error) {
	var directly bool
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindInputFeatureNamed:featureValue:bufferIndex:cleanUpBlocks:boundDirectly:error:"), named, value, index, blocks, unsafe.Pointer(&directly), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindInputFeatureNamed:featureValue:bufferIndex:cleanUpBlocks:boundDirectly:error: returned NO with nil NSError")
	}
	return directly, nil
}
func (m MLNeuralNetworkEngine) BindInputFeatureNamedPixelBufferCleanUpBlocksError(named objectivec.IObject, buffer corevideo.CVImageBufferRef, blocks objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindInputFeatureNamed:pixelBuffer:cleanUpBlocks:error:"), named, buffer, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindInputFeatureNamed:pixelBuffer:cleanUpBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) BindInputFeaturesBufferIndexCleanUpBlocksDirectlyBoundFeatureNamesError(features objectivec.IObject, index uint64, blocks objectivec.IObject, names []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindInputFeatures:bufferIndex:cleanUpBlocks:directlyBoundFeatureNames:error:"), features, index, blocks, objectivec.IObjectSliceToNSArray(names), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindInputFeatures:bufferIndex:cleanUpBlocks:directlyBoundFeatureNames:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) BindInputsAndOutputsCleanUpBlocksBufferIndexOptionsError(outputs objectivec.IObject, blocks objectivec.IObject, index uint64, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindInputsAndOutputs:cleanUpBlocks:bufferIndex:options:error:"), outputs, blocks, index, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindInputsAndOutputs:cleanUpBlocks:bufferIndex:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) BindOutputBuffersOutputBackingsAutomaticOutputBackingModeDirectlyBoundOutputFeatureNamesError(buffers unsafe.Pointer, backings objectivec.IObject, mode objectivec.IObject, names []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("bindOutputBuffers:outputBackings:automaticOutputBackingMode:directlyBoundOutputFeatureNames:error:"), buffers, backings, mode, objectivec.IObjectSliceToNSArray(names), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("bindOutputBuffers:outputBackings:automaticOutputBackingMode:directlyBoundOutputFeatureNames:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) CollectParametersFromContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("collectParametersFromContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("collectParametersFromContainer:configuration:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) CompleteOutputBackingsAutomaticOutputBackingModeError(backings objectivec.IObject, mode objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("completeOutputBackings:automaticOutputBackingMode:error:"), backings, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) ConvertPredictionToClassifierResultWithOptionsError(result objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("convertPredictionToClassifierResult:withOptions:error:"), result, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) CopyEbufOfPixelTypeToPixelBufferError(ebuf unsafe.Pointer, type_ uint64, buffer corevideo.CVImageBufferRef) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("copyEbuf:ofPixelType:toPixelBuffer:error:"), ebuf, type_, buffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyEbuf:ofPixelType:toPixelBuffer:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) CopyImagePreprocessingParametersToError(to unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("copyImagePreprocessingParametersTo:error:"), to, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyImagePreprocessingParametersTo:error: returned NO with nil NSError")
	}
	return rv, nil

}

// preprocessing is a [*appleneuralengine.vimage2espresso_param].
func (m MLNeuralNetworkEngine) CopyPixelBufferByApplyingImagePreprocessingToPixelBuffer(preprocessing unsafe.Pointer, buffer corevideo.CVImageBufferRef) corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("copyPixelBufferByApplyingImagePreprocessing:toPixelBuffer:"), preprocessing, buffer)
	return corevideo.CVImageBufferRef(rv)
}
func (m MLNeuralNetworkEngine) CopyPixelBufferByApplyingImagePreprocessingForFeatureNamedToPixelBuffer(named objectivec.IObject, buffer corevideo.CVImageBufferRef) corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("copyPixelBufferByApplyingImagePreprocessingForFeatureNamed:toPixelBuffer:"), named, buffer)
	return corevideo.CVImageBufferRef(rv)
}
func (m MLNeuralNetworkEngine) CopyPixelBufferFromPixelBufferUsingPixelFormat(buffer corevideo.CVImageBufferRef, format uint32) corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("copyPixelBufferFromPixelBuffer:usingPixelFormat:"), buffer, format)
	return corevideo.CVImageBufferRef(rv)
}
func (m MLNeuralNetworkEngine) DumpTestVectorsToPath(path objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("dumpTestVectorsToPath:"), path)
}
func (m MLNeuralNetworkEngine) EvaluateError(evaluate objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluate:error:"), evaluate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) EvaluateBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) EvaluateInputsBufferIndexOptionsError(inputs objectivec.IObject, index uint64, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateInputs:bufferIndex:options:error:"), inputs, index, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) EvaluateInputsOptionsError(inputs objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateInputs:options:error:"), inputs, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) EvaluateInputsOptionsVerifyInputsError(inputs objectivec.IObject, options objectivec.IObject, inputs2 bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateInputs:options:verifyInputs:error:"), inputs, options, inputs2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) ExecutePlanError(plan unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("executePlan:error:"), plan, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("executePlan:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) ImageFeatureValueFromEbufBackingCVPixelBufferDescriptionError(ebuf unsafe.Pointer, buffer corevideo.CVImageBufferRef, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("imageFeatureValueFromEbuf:backingCVPixelBuffer:description:error:"), ebuf, buffer, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) ImageFeatureValueFromPixelBufferUsingPixelFormat(buffer corevideo.CVImageBufferRef, format uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("imageFeatureValueFromPixelBuffer:usingPixelFormat:"), buffer, format)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkEngine) InputBindStateForFeatureValueError(value objectivec.IObject) (int64, error) {
	var errorPtr objc.ID
	rv := objc.Send[int64](m.ID, objc.Sel("inputBindStateForFeatureValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) LockPixelBufferCleanUpBlocksError(buffer corevideo.CVImageBufferRef, blocks objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("lockPixelBuffer:cleanUpBlocks:error:"), buffer, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("lockPixelBuffer:cleanUpBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) MultiArrayFeatureValueFromEbufBackingMultiArrayDescriptionOutputNameError(ebuf unsafe.Pointer, array objectivec.IObject, description objectivec.IObject, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("multiArrayFeatureValueFromEbuf:backingMultiArray:description:outputName:error:"), ebuf, array, description, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) ObtainBuffer() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("obtainBuffer"))
	return rv
}
func (m MLNeuralNetworkEngine) OpacifyAndPermutePixelBufferBufferContainsBGRAError(buffer corevideo.CVImageBufferRef, bgra bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("opacifyAndPermutePixelBuffer:bufferContainsBGRA:error:"), buffer, bgra, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("opacifyAndPermutePixelBuffer:bufferContainsBGRA:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) OutputBackingMultiArrayForFeatureName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputBackingMultiArrayForFeatureName:"), name)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkEngine) PixelBufferBackedMultiArrayWithShape(shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelBufferBackedMultiArrayWithShape:"), shape)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkEngine) PixelBufferFromOutputBackingForFeature(backing objectivec.IObject, feature objectivec.IObject) corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("pixelBufferFromOutputBacking:forFeature:"), backing, feature)
	return corevideo.CVImageBufferRef(rv)
}
func (m MLNeuralNetworkEngine) PopulateMultiArrayShapeStridesForEbufFeatureDescriptionNdArrayInterpretation(shape []objectivec.IObject, strides []objectivec.IObject, ebuf unsafe.Pointer, description objectivec.IObject, interpretation bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("populateMultiArrayShape:strides:forEbuf:featureDescription:ndArrayInterpretation:"), objectivec.IObjectSliceToNSArray(shape), objectivec.IObjectSliceToNSArray(strides), ebuf, description, interpretation)
}
func (m MLNeuralNetworkEngine) PopulateOutputsOutputBackingsDirectlyBoundOutputFeatureNamesError(outputs uint64, backings objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("populateOutputs:outputBackings:directlyBoundOutputFeatureNames:error:"), outputs, backings, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) PrepareBlobNamedForNewBlobBackingModeBindMode(named objectivec.IObject, mode int64, mode2 int) {
	objc.Send[objc.ID](m.ID, objc.Sel("prepareBlobNamed:forNewBlobBackingMode:bindMode:"), named, mode, mode2)
}
func (m MLNeuralNetworkEngine) RebuildPlan(plan []objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("rebuildPlan:"), objectivec.IObjectSliceToNSArray(plan))
	return rv
}
func (m MLNeuralNetworkEngine) RebuildPlanError(plan bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("rebuildPlan:error:"), plan, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("rebuildPlan:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) ReleaseBuffer(buffer uint64) {
	objc.Send[objc.ID](m.ID, objc.Sel("releaseBuffer:"), buffer)
}
func (m MLNeuralNetworkEngine) ResetSizesError(sizes objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resetSizes:error:"), sizes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resetSizes:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) ResetSizesNoAutoReleaseError(release objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resetSizesNoAutoRelease:error:"), release, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resetSizesNoAutoRelease:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) ResetSizesWithEspressoConfigurationsError(configurations objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resetSizesWithEspressoConfigurations:error:"), configurations, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resetSizesWithEspressoConfigurations:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) SequenceConcatConsumesOptionalInputNamed(named objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("sequenceConcatConsumesOptionalInputNamed:"), named)
	return rv
}
func (m MLNeuralNetworkEngine) SequenceNamed(named objectivec.IObject) int {
	rv := objc.Send[int](m.ID, objc.Sel("sequenceNamed:"), named)
	return rv
}
func (m MLNeuralNetworkEngine) SetEspressoBlobShapesWidthsHeightsKsBatchesSequencesRanksError(shapes unsafe.Pointer, widths unsafe.Pointer, heights unsafe.Pointer, ks unsafe.Pointer, batches unsafe.Pointer, sequences unsafe.Pointer, ranks unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setEspressoBlobShapes:widths:heights:ks:batches:sequences:ranks:error:"), shapes, widths, heights, ks, batches, sequences, ranks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setEspressoBlobShapes:widths:heights:ks:batches:sequences:ranks:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) SortBatchByShapeWithMapError(shape objectivec.IObject, map_ []objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sortBatchByShape:withMap:error:"), shape, objectivec.IObjectSliceToNSArray(map_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) SupportFromEspressoLayerInfo(info objectivec.IObject) uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("supportFromEspressoLayerInfo:"), info)
	return rv
}
func (m MLNeuralNetworkEngine) SupportFromEspressoPlatform(platform int) uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("supportFromEspressoPlatform:"), platform)
	return rv
}
func (m MLNeuralNetworkEngine) TransferOneComponent16HalfPixelBufferToPixelBufferWithScaleBias(buffer corevideo.CVImageBufferRef, buffer2 corevideo.CVImageBufferRef, scale float32, bias float32) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("transferOneComponent16HalfPixelBuffer:toPixelBuffer:withScale:bias:"), buffer, buffer2, scale, bias)
	return rv
}
func (m MLNeuralNetworkEngine) TransferPixelBufferToPixelBuffer(buffer corevideo.CVImageBufferRef, buffer2 corevideo.CVImageBufferRef) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("transferPixelBuffer:toPixelBuffer:"), buffer, buffer2)
	return rv
}
func (m MLNeuralNetworkEngine) TryToSetOutputBackingForFeatureNameToEbufReportPointerFlagsError(backing objectivec.IObject, name objectivec.IObject, ebuf unsafe.Pointer) (int, error) {
	var flags int
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("tryToSetOutputBacking:forFeatureName:toEbuf:reportPointerFlags:error:"), backing, name, ebuf, unsafe.Pointer(&flags), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("tryToSetOutputBacking:forFeatureName:toEbuf:reportPointerFlags:error: returned NO with nil NSError")
	}
	return flags, nil
}
func (m MLNeuralNetworkEngine) UpdateDynamicOutputBlobIndicatorCacheAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updateDynamicOutputBlobIndicatorCacheAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateDynamicOutputBlobIndicatorCacheAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkEngine) UsingEspressoConfigurations() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usingEspressoConfigurations"))
	return rv
}
func (m MLNeuralNetworkEngine) VerifyInputsError(inputs objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("verifyInputs:error:"), inputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkEngine) InitWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkEngineFromID(rv), nil

}
func (m MLNeuralNetworkEngine) InitWithContainerError(container objectivec.IObject) (MLNeuralNetworkEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkEngineFromID(rv), nil

}

func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) ContainerClass() objectivec.Class {
	rv := objc.Send[objectivec.Class](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("containerClass"))
	return objectivec.Class(rv)
}
func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) GpuEngine() int {
	rv := objc.Send[int](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("gpuEngine"))
	return rv
}
func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) GpuPrecision() int {
	rv := objc.Send[int](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("gpuPrecision"))
	return rv
}
func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) LoadModelAssetDescriptionFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("loadModelAssetDescriptionFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) NeuralNetworkFromContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("neuralNetworkFromContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkEngineClass MLNeuralNetworkEngineClass) NeuralNetworkFromContainerError(container objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkEngineClass.class), objc.Sel("neuralNetworkFromContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLNeuralNetworkEngine) ActiveFunction() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("activeFunction"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkEngine) SetActiveFunction(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setActiveFunction:"), objc.String(value))
}
func (m MLNeuralNetworkEngine) BufferSemaphore() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("bufferSemaphore"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetBufferSemaphore(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setBufferSemaphore:"), value)
}
func (m MLNeuralNetworkEngine) ClassLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetClassLabels(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassLabels:"), value)
}
func (m MLNeuralNetworkEngine) ClassScoreVectorName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classScoreVectorName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkEngine) SetClassScoreVectorName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassScoreVectorName:"), objc.String(value))
}
func (m MLNeuralNetworkEngine) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("context"))
	return rv
}
func (m MLNeuralNetworkEngine) SetContext(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setContext:"), value)
}
func (m MLNeuralNetworkEngine) DefaultOptionalValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultOptionalValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetDefaultOptionalValues(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setDefaultOptionalValues:"), value)
}
func (m MLNeuralNetworkEngine) Engine() int {
	rv := objc.Send[int](m.ID, objc.Sel("engine"))
	return rv
}
func (m MLNeuralNetworkEngine) SetEngine(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setEngine:"), value)
}
func (m MLNeuralNetworkEngine) EspressoInputShapes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("espressoInputShapes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetEspressoInputShapes(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setEspressoInputShapes:"), value)
}
func (m MLNeuralNetworkEngine) EspressoInputStrides() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("espressoInputStrides"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetEspressoInputStrides(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setEspressoInputStrides:"), value)
}
func (m MLNeuralNetworkEngine) EspressoProfileInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("espressoProfileInfo"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkEngine) SetEspressoProfileInfo(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setEspressoProfileInfo:"), value)
}
func (m MLNeuralNetworkEngine) EspressoQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("espressoQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetEspressoQueue(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setEspressoQueue:"), value)
}
func (m MLNeuralNetworkEngine) HardwareFallbackDetected() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hardwareFallbackDetected"))
	return rv
}
func (m MLNeuralNetworkEngine) SetHardwareFallbackDetected(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHardwareFallbackDetected:"), value)
}
func (m MLNeuralNetworkEngine) HasBidirectionalLayer() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasBidirectionalLayer"))
	return rv
}
func (m MLNeuralNetworkEngine) SetHasBidirectionalLayer(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHasBidirectionalLayer:"), value)
}
func (m MLNeuralNetworkEngine) HasOptionalInputSequenceConcat() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasOptionalInputSequenceConcat"))
	return rv
}
func (m MLNeuralNetworkEngine) SetHasOptionalInputSequenceConcat(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHasOptionalInputSequenceConcat:"), value)
}
func (m MLNeuralNetworkEngine) ImagePreprocessingParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("imagePreprocessingParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetImagePreprocessingParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setImagePreprocessingParameters:"), value)
}
func (m MLNeuralNetworkEngine) InputBlobNameToLastBackingMode() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputBlobNameToLastBackingMode"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) InputFeatureConformer() IMLFeatureProviderConformer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputFeatureConformer"))
	return MLFeatureProviderConformerFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) InputLayers() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputLayers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) IsANEPathForbidden() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isANEPathForbidden"))
	return rv
}
func (m MLNeuralNetworkEngine) SetIsANEPathForbidden(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setIsANEPathForbidden:"), value)
}
func (m MLNeuralNetworkEngine) IsEspressoBiasPreprocessingShared() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEspressoBiasPreprocessingShared"))
	return rv
}
func (m MLNeuralNetworkEngine) SetIsEspressoBiasPreprocessingShared(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setIsEspressoBiasPreprocessingShared:"), value)
}
func (m MLNeuralNetworkEngine) IsGPUPathForbidden() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isGPUPathForbidden"))
	return rv
}
func (m MLNeuralNetworkEngine) SetIsGPUPathForbidden(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setIsGPUPathForbidden:"), value)
}
func (m MLNeuralNetworkEngine) ModelFilePath() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelFilePath"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkEngine) SetModelFilePath(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelFilePath:"), objc.String(value))
}
func (m MLNeuralNetworkEngine) ModelIsEncrypted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("modelIsEncrypted"))
	return rv
}
func (m MLNeuralNetworkEngine) ModelIsMIL() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("modelIsMIL"))
	return rv
}
func (m MLNeuralNetworkEngine) SetModelIsMIL(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelIsMIL:"), value)
}
func (m MLNeuralNetworkEngine) ModelVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) NdArrayInterpretation() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("ndArrayInterpretation"))
	return rv
}
func (m MLNeuralNetworkEngine) SetNdArrayInterpretation(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setNdArrayInterpretation:"), value)
}
func (m MLNeuralNetworkEngine) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("network"))
	return rv
}
func (m MLNeuralNetworkEngine) SetNetwork(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setNetwork:"), value)
}
func (m MLNeuralNetworkEngine) NumInputs() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numInputs"))
	return rv
}
func (m MLNeuralNetworkEngine) NumOutputs() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numOutputs"))
	return rv
}
func (m MLNeuralNetworkEngine) OptionalInputTypes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("optionalInputTypes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) OutputBlobNameToLastBackingMode() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputBlobNameToLastBackingMode"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) OutputLayers() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputLayers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) Plan() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("plan"))
	return rv
}
func (m MLNeuralNetworkEngine) SetPlan(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setPlan:"), value)
}
func (m MLNeuralNetworkEngine) Precision() int {
	rv := objc.Send[int](m.ID, objc.Sel("precision"))
	return rv
}
func (m MLNeuralNetworkEngine) SetPrecision(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrecision:"), value)
}
func (m MLNeuralNetworkEngine) PredictionsQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetPredictionsQueue(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setPredictionsQueue:"), value)
}
func (m MLNeuralNetworkEngine) Priority() int {
	rv := objc.Send[int](m.ID, objc.Sel("priority"))
	return rv
}
func (m MLNeuralNetworkEngine) SetPriority(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setPriority:"), value)
}
func (m MLNeuralNetworkEngine) ProbabilityDictionarySharedKeySet() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("probabilityDictionarySharedKeySet"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkEngine) SetProbabilityDictionarySharedKeySet(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setProbabilityDictionarySharedKeySet:"), value)
}
func (m MLNeuralNetworkEngine) Qos() int {
	rv := objc.Send[int](m.ID, objc.Sel("qos"))
	return rv
}
func (m MLNeuralNetworkEngine) SetQos(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setQos:"), value)
}
func (m MLNeuralNetworkEngine) SubmitSemaphore() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("submitSemaphore"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNeuralNetworkEngine) SetSubmitSemaphore(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setSubmitSemaphore:"), value)
}
func (m MLNeuralNetworkEngine) UsingCPU() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("usingCPU"))
	return rv
}
func (m MLNeuralNetworkEngine) SetUsingCPU(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setUsingCPU:"), value)
}
