// Code generated from Apple documentation. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// DictionaryHandler is the signature for a completion handler block.
//
// Used by:
//   - [AUMessageChannel.SetCallHostBlock]
type DictionaryHandler = func(*foundation.INSDictionary)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [TTSWrappedAudioQueue.PlayBufferCompletionHandler]
//   - [TTSWrappedAudioQueue.ScheduleBufferCompletionHandlerLastBuffer]
//   - [TTSWrappedAudioQueue.ScheduleBufferCompletionHandler]
//   - [TTSWrappedAudioQueueBuffer.SetCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.PauseSpeakingAtCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopSpeakingAtCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler]
//   - [TextToSpeechTTSAURenderer.FormatForVoiceCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentLocaleIdentifiersWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentSystemLocaleIdentifierWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentSystemLocaleWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.FallbackForVoiceCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForIdentifierCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForIdentifierPreferringLanguageCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForLocaleCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForLocaleIdentifierCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TTSWrappedAudioQueue.PlayBufferCompletionHandler]
//   - [TTSWrappedAudioQueue.ScheduleBufferCompletionHandlerLastBuffer]
//   - [TTSWrappedAudioQueue.ScheduleBufferCompletionHandler]
//   - [TTSWrappedAudioQueueBuffer.SetCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.PauseSpeakingAtCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopSpeakingAtCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler]
//   - [TextToSpeechTTSAURenderer.FormatForVoiceCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentLocaleIdentifiersWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentSystemLocaleIdentifierWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentSystemLocaleWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.FallbackForVoiceCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForIdentifierCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForIdentifierPreferringLanguageCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForLocaleCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForLocaleIdentifierCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [TTSAUMessagingAU.SetCallHostBlock]
//   - [TTSAUMessagingAU.SetHostBlock]
//   - [TTSAXResourceManager._performBlockOnObservers]
//   - [TTSApplebetMapperRule.SetMatchRule]
//   - [TTSEmojiUtilities.EnumerateEmojiCharactersInStringLanguageCodeWithBlock]
//   - [TTSExceptionCatcher.CatchExceptionError]
//   - [TTSRegex.EnumerateMatchesInCStringLengthUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringRangesUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock]
//   - [TTSRuleReplacement.SetPostMatch]
//   - [TTSRulesetRunner.SetMatchLogger]
//   - [TTSRulesetRunner.SetPostRuleWriter]
//   - [TTSRulesetRunner.SetPreRuleWriter]
//   - [TTSSiriAssetManager.DownloadAssetProgressHandler]
//   - [TTSSiriAssetManager.DownloadVoiceResourceForLanguageCompletion]
//   - [TTSSpeechAction.SetAudioBufferCallback]
//   - [TTSSpeechAction.SetCompletionCallback]
//   - [TTSSpeechAction.SetMarkerCallback]
//   - [TTSSpeechAction.SetOnMarkerCallback]
//   - [TTSSpeechAction.SetOnPauseCallback]
//   - [TTSSpeechAction.SetOnResumeCallback]
//   - [TTSSpeechAction.SetOnSpeechStartCallback]
//   - [TTSSpeechAction.SetOnWillSpeakRangeCallback]
//   - [TTSSpeechManager.SetRequestWillStart]
//   - [TTSSpeechManager.Test_actionStartTap]
//   - [TTSSpeechRequest.SetAudioBufferCallback]
//   - [TTSSpeechRequest.SetLatencyCallback]
//   - [TTSSpeechSynthesizer.SetAudioBufferCallback]
//   - [TTSStreamingZipReader.EnumerateFiles]
//   - [TTSVoiceResourceManager.EnumerateLoadableResourcesInAssetUsingBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [TTSAUMessagingAU.SetCallHostBlock]
//   - [TTSAUMessagingAU.SetHostBlock]
//   - [TTSAXResourceManager._performBlockOnObservers]
//   - [TTSApplebetMapperRule.SetMatchRule]
//   - [TTSEmojiUtilities.EnumerateEmojiCharactersInStringLanguageCodeWithBlock]
//   - [TTSExceptionCatcher.CatchExceptionError]
//   - [TTSRegex.EnumerateMatchesInCStringLengthUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringRangesUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock]
//   - [TTSRuleReplacement.SetPostMatch]
//   - [TTSRulesetRunner.SetMatchLogger]
//   - [TTSRulesetRunner.SetPostRuleWriter]
//   - [TTSRulesetRunner.SetPreRuleWriter]
//   - [TTSSiriAssetManager.DownloadAssetProgressHandler]
//   - [TTSSiriAssetManager.DownloadVoiceResourceForLanguageCompletion]
//   - [TTSSpeechAction.SetAudioBufferCallback]
//   - [TTSSpeechAction.SetCompletionCallback]
//   - [TTSSpeechAction.SetMarkerCallback]
//   - [TTSSpeechAction.SetOnMarkerCallback]
//   - [TTSSpeechAction.SetOnPauseCallback]
//   - [TTSSpeechAction.SetOnResumeCallback]
//   - [TTSSpeechAction.SetOnSpeechStartCallback]
//   - [TTSSpeechAction.SetOnWillSpeakRangeCallback]
//   - [TTSSpeechManager.SetRequestWillStart]
//   - [TTSSpeechManager.Test_actionStartTap]
//   - [TTSSpeechRequest.SetAudioBufferCallback]
//   - [TTSSpeechRequest.SetLatencyCallback]
//   - [TTSSpeechSynthesizer.SetAudioBufferCallback]
//   - [TTSStreamingZipReader.EnumerateFiles]
//   - [TTSVoiceResourceManager.EnumerateLoadableResourcesInAssetUsingBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

