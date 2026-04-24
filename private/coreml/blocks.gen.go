// Code generated from Apple documentation. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [CoreMLMLOdieEngine.SubmitPredictionRequestCompletionHandler]
//   - [MLDelegateModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [MLDelegateModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLDelegateModel.PrepareWithCompletionHandler]
//   - [MLDelegateModel.SubmitPredictionRequestCompletionHandler]
//   - [MLDelegateModel._predictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLDelegateModel._schedulePredictionRequestCompletionHandler]
//   - [MLDelegateModel._submitPredictionRequestCompletionHandler]
//   - [MLDelegateUpdatableModel.SetUpdateProgressHandlersDispatchQueue]
//   - [MLE5Engine.SubmitPredictionRequestCompletionHandler]
//   - [MLE5Engine._predictionFromFeaturesOptionsCompletionHandler]
//   - [MLE5ExecutionStream.SubmitWithCompletionHandler]
//   - [MLGenericPredictionRequest.SubmitWithCompletionHandler]
//   - [MLKNearestNeighborsClassifier.SetProgressHandlers]
//   - [MLKNearestNeighborsClassifier.SetUpdateProgressHandlersDispatchQueue]
//   - [MLModel.PrepareWithCompletionHandler]
//   - [MLModel.SubmitPredictionRequestCompletionHandler]
//   - [MLModelAsset.ModelStructureWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.FunctionNamesWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ModelDescriptionOfFunctionNamedCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ModelDescriptionWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor._modelAssetDescriptionWithCompletionHandler]
//   - [MLModelAssetModelStructureVendor.ModelStructureWithCompletionHandler]
//   - [MLModelAssetModelStructureVendor._modelStructureWithCompletionHandler]
//   - [MLModelAssetModelVendor.ModelWithConfigurationCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelAssetDescriptionWithCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelStructureWithCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelWithConfigurationCompletionHandler]
//   - [MLModelCollection.BeginAccessingModelCollectionWithIdentifierCompletionHandler]
//   - [MLModelCollection.EndAccessingModelCollectionWithIdentifierCompletionHandler]
//   - [MLModelEngine.SubmitPredictionRequestCompletionHandler]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetProgressHandlers]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkUpdateEngine.SetProgressHandlers]
//   - [MLNeuralNetworkUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLPendingPrediction.InitWithPredictionRequestCompletionHandler]
//   - [MLPipelineClassifier.ClassifyOptionsCompletionHandler]
//   - [MLPipelineUpdateEngine.SetProgressHandlers]
//   - [MLPipelineUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetProgressHandlers]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLUpdatable.SetUpdateProgressHandlersDispatchQueue]
//   - [MLUpdateProgressHandlers.SetCompletionHandler]
//   - [MLUpdateTask.InitWithModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataProgressHandlersError]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CoreMLMLOdieEngine.SubmitPredictionRequestCompletionHandler]
//   - [MLDelegateModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [MLDelegateModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLDelegateModel.PrepareWithCompletionHandler]
//   - [MLDelegateModel.SubmitPredictionRequestCompletionHandler]
//   - [MLDelegateModel._predictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLDelegateModel._schedulePredictionRequestCompletionHandler]
//   - [MLDelegateModel._submitPredictionRequestCompletionHandler]
//   - [MLDelegateUpdatableModel.SetUpdateProgressHandlersDispatchQueue]
//   - [MLE5Engine.SubmitPredictionRequestCompletionHandler]
//   - [MLE5Engine._predictionFromFeaturesOptionsCompletionHandler]
//   - [MLE5ExecutionStream.SubmitWithCompletionHandler]
//   - [MLGenericPredictionRequest.SubmitWithCompletionHandler]
//   - [MLKNearestNeighborsClassifier.SetProgressHandlers]
//   - [MLKNearestNeighborsClassifier.SetUpdateProgressHandlersDispatchQueue]
//   - [MLModel.PrepareWithCompletionHandler]
//   - [MLModel.SubmitPredictionRequestCompletionHandler]
//   - [MLModelAsset.ModelStructureWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.FunctionNamesWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ModelDescriptionOfFunctionNamedCompletionHandler]
//   - [MLModelAssetDescriptionVendor.ModelDescriptionWithCompletionHandler]
//   - [MLModelAssetDescriptionVendor._modelAssetDescriptionWithCompletionHandler]
//   - [MLModelAssetModelStructureVendor.ModelStructureWithCompletionHandler]
//   - [MLModelAssetModelStructureVendor._modelStructureWithCompletionHandler]
//   - [MLModelAssetModelVendor.ModelWithConfigurationCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelAssetDescriptionWithCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelStructureWithCompletionHandler]
//   - [MLModelAssetResourceFactory.ModelWithConfigurationCompletionHandler]
//   - [MLModelCollection.BeginAccessingModelCollectionWithIdentifierCompletionHandler]
//   - [MLModelCollection.EndAccessingModelCollectionWithIdentifierCompletionHandler]
//   - [MLModelEngine.SubmitPredictionRequestCompletionHandler]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetProgressHandlers]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkUpdateEngine.SetProgressHandlers]
//   - [MLNeuralNetworkUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLPendingPrediction.InitWithPredictionRequestCompletionHandler]
//   - [MLPipelineClassifier.ClassifyOptionsCompletionHandler]
//   - [MLPipelineUpdateEngine.SetProgressHandlers]
//   - [MLPipelineUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetProgressHandlers]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLUpdatable.SetUpdateProgressHandlersDispatchQueue]
//   - [MLUpdateProgressHandlers.SetCompletionHandler]
//   - [MLUpdateTask.InitWithModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataProgressHandlersError]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLClassifierResultErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLAsyncClassifier.ClassifyOptionsCompletionHandler]
type MLClassifierResultErrorHandler = func(*MLClassifierResult, error)

// NewMLClassifierResultErrorBlock wraps a Go [MLClassifierResultErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLAsyncClassifier.ClassifyOptionsCompletionHandler]
func NewMLClassifierResultErrorBlock(handler MLClassifierResultErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MLClassifierResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLClassifierResultFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLFeatureProviderErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLModeling.SubmitPredictionRequestCompletionHandler]
//   - [MLPredictionRequest.SubmitWithCompletionHandler]
type MLFeatureProviderErrorHandler = func(*objectivec.Object, error)

// NewMLFeatureProviderErrorBlock wraps a Go [MLFeatureProviderErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModeling.SubmitPredictionRequestCompletionHandler]
//   - [MLPredictionRequest.SubmitWithCompletionHandler]
func NewMLFeatureProviderErrorBlock(handler MLFeatureProviderErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *objectivec.Object
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := objectivec.ObjectFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLFeatureValueErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [CoreMLModelSecurityServiceToClientProtocol.ClientFeatureValueForNameUniqueKeyForProviderWithReply]
type MLFeatureValueErrorHandler = func(*MLFeatureValue, error)

// NewMLFeatureValueErrorBlock wraps a Go [MLFeatureValueErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CoreMLModelSecurityServiceToClientProtocol.ClientFeatureValueForNameUniqueKeyForProviderWithReply]
func NewMLFeatureValueErrorBlock(handler MLFeatureValueErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MLFeatureValue
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLFeatureValueFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// NSNumberArrayHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLMultiArray.GetMutableBytesWithHandler]
type NSNumberArrayHandler = func(*[]foundation.NSNumber)

// SetErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [CoreMLModelSecurityServiceToClientProtocol.ClientFeatureNamesWithReply]
type SetErrorHandler = func(*foundation.INSSet, error)

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [CoreMLModelSecurityServiceToClient.ClientFeatureNamesWithReply]
//   - [CoreMLModelSecurityServiceToClient.ClientFeatureValueForNameUniqueKeyForProviderWithReply]
//   - [MLAppleAudioFeatureExtractor.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleGazetteer.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleImageFeatureExtractor.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleSoundAnalysisPreprocessing.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleTextClassifier.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleWordEmbedding.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleWordTagger.NewStateForFeatureNamedInitializerBlock]
//   - [MLArrayFeatureExtractor.NewStateForFeatureNamedInitializerBlock]
//   - [MLClassConfidenceThresholding.NewStateForFeatureNamedInitializerBlock]
//   - [MLCustomModelWrapper.NewStateForFeatureNamedInitializerBlock]
//   - [MLDelegateModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLDelegateUpdatableModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithIRProgramContainerConfigurationDeallocator]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration]
//   - [MLIdentity.NewStateForFeatureNamedInitializerBlock]
//   - [MLItemSimilarityRecommender.NewStateForFeatureNamedInitializerBlock]
//   - [MLKNearestNeighborsClassifier.NewStateForFeatureNamedInitializerBlock]
//   - [MLLinkedModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLLocalOutlierFactor.NewStateForFeatureNamedInitializerBlock]
//   - [MLMetalDeviceObserver.CopyAllMTLDevicesWithHandlerDeviceObserver]
//   - [MLMetalDeviceObserver.StartObservingWithBlockDeviceObserver]
//   - [MLModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLModel.PredictionsFromSubbatchingBatchMaxSubbatchLengthPredictionBlockOptionsError]
//   - [MLModelCollection._downloadWithProgress]
//   - [MLMultiArray.GetContiguousFirstMajorFloat32BufferWithHandler]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [MLMultiArray.InitWithPixelBufferShapeStridesMutableShapedBufferProvider]
//   - [MLMultiArrayView.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [MLMultiArrayView.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [MLMultiArrayView.InitWithPixelBufferShapeStridesMutableShapedBufferProvider]
//   - [MLNeuralNetworkCompiler.NewStateForFeatureNamedInitializerBlock]
//   - [MLSecureModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLState.GetMultiArrayWithHandler]
//   - [MLState.InternalGetMultiArrayWithHandler]
//   - [MLUpdateProgressHandlers.SetProgressHandler]
//   - [MLWrappedModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLWritableWrappedModel.NewStateForFeatureNamedInitializerBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CoreMLModelSecurityServiceToClient.ClientFeatureNamesWithReply]
//   - [CoreMLModelSecurityServiceToClient.ClientFeatureValueForNameUniqueKeyForProviderWithReply]
//   - [MLAppleAudioFeatureExtractor.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleGazetteer.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleImageFeatureExtractor.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleSoundAnalysisPreprocessing.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleTextClassifier.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleWordEmbedding.NewStateForFeatureNamedInitializerBlock]
//   - [MLAppleWordTagger.NewStateForFeatureNamedInitializerBlock]
//   - [MLArrayFeatureExtractor.NewStateForFeatureNamedInitializerBlock]
//   - [MLClassConfidenceThresholding.NewStateForFeatureNamedInitializerBlock]
//   - [MLCustomModelWrapper.NewStateForFeatureNamedInitializerBlock]
//   - [MLDelegateModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLDelegateUpdatableModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithIRProgramContainerConfigurationDeallocator]
//   - [MLE5ProgramLibraryOnDeviceAOTCompilationImpl.InitWithMILTextAtURLIrProgramDeallocatorContainerConfiguration]
//   - [MLIdentity.NewStateForFeatureNamedInitializerBlock]
//   - [MLItemSimilarityRecommender.NewStateForFeatureNamedInitializerBlock]
//   - [MLKNearestNeighborsClassifier.NewStateForFeatureNamedInitializerBlock]
//   - [MLLinkedModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLLocalOutlierFactor.NewStateForFeatureNamedInitializerBlock]
//   - [MLMetalDeviceObserver.CopyAllMTLDevicesWithHandlerDeviceObserver]
//   - [MLMetalDeviceObserver.StartObservingWithBlockDeviceObserver]
//   - [MLModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLModel.PredictionsFromSubbatchingBatchMaxSubbatchLengthPredictionBlockOptionsError]
//   - [MLModelCollection._downloadWithProgress]
//   - [MLMultiArray.GetContiguousFirstMajorFloat32BufferWithHandler]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [MLMultiArray.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [MLMultiArray.InitWithPixelBufferShapeStridesMutableShapedBufferProvider]
//   - [MLMultiArrayView.InitWithBytesNoCopyShapeDataTypeStridesDeallocatorMutableShapedBufferProviderError]
//   - [MLMultiArrayView.InitWithBytesNoCopyShapeDataTypeStridesMutableShapedBufferProvider]
//   - [MLMultiArrayView.InitWithPixelBufferShapeStridesMutableShapedBufferProvider]
//   - [MLNeuralNetworkCompiler.NewStateForFeatureNamedInitializerBlock]
//   - [MLSecureModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLState.GetMultiArrayWithHandler]
//   - [MLState.InternalGetMultiArrayWithHandler]
//   - [MLUpdateProgressHandlers.SetProgressHandler]
//   - [MLWrappedModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLWritableWrappedModel.NewStateForFeatureNamedInitializerBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
