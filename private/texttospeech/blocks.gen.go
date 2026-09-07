// Code generated from Apple documentation. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [TextToSpeechCoreSynthesisVoiceShim.CoreVoiceWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
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
//   - [TextToSpeechCoreSynthesisVoiceShim.CoreVoiceWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// INSDictionaryINSDictionaryHandler handles a primitive value and returns a primitive value.
//
// Used by:
//   - [AUMessageChannel.SetCallHostBlock]
type INSDictionaryINSDictionaryHandler = func(foundation.INSDictionary) foundation.INSDictionary

// NewINSDictionaryINSDictionaryBlock wraps a Go [INSDictionaryINSDictionaryHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AUMessageChannel.SetCallHostBlock]
func NewINSDictionaryINSDictionaryBlock(handler INSDictionaryINSDictionaryHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID) foundation.INSDictionary {
		var primitiveVal foundation.INSDictionary
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitiveVal = foundation.NSDictionaryFromID(primitiveID)
		}
		return handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [BFSpeechChannel.SetCfWordCallback]
//   - [BFSpeechChannel.SetErrorCallback]
//   - [BFSpeechChannel.SetPendingStopBlock]
//   - [BFSpeechChannel.SetPhonemeCallback]
//   - [BFSpeechChannel.SetSpeechDoneCallback]
//   - [BFSpeechChannel.SetSyncCallback]
//   - [BFSpeechChannel.SetWordCallback]
//   - [BabelFish.PerformForChannelBlock]
//   - [TTSAUMessagingAU.SetCallHostBlock]
//   - [TTSAUMessagingAU.SetHostBlock]
//   - [TTSAXResourceManager._performBlockOnObservers]
//   - [TTSApplebetMapperRule.SetMatchRule]
//   - [TTSEmojiUtilities.EnumerateEmojiCharactersInStringLanguageCodeWithBlock]
//   - [TTSExceptionCatcher.CatchExceptionError]
//   - [TTSRegex.EnumerateMatchesInCStringLengthUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringRangesUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock]
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
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackSynth]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynth]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [BFSpeechChannel.SetCfWordCallback]
//   - [BFSpeechChannel.SetErrorCallback]
//   - [BFSpeechChannel.SetPendingStopBlock]
//   - [BFSpeechChannel.SetPhonemeCallback]
//   - [BFSpeechChannel.SetSpeechDoneCallback]
//   - [BFSpeechChannel.SetSyncCallback]
//   - [BFSpeechChannel.SetWordCallback]
//   - [BabelFish.PerformForChannelBlock]
//   - [TTSAUMessagingAU.SetCallHostBlock]
//   - [TTSAUMessagingAU.SetHostBlock]
//   - [TTSAXResourceManager._performBlockOnObservers]
//   - [TTSApplebetMapperRule.SetMatchRule]
//   - [TTSEmojiUtilities.EnumerateEmojiCharactersInStringLanguageCodeWithBlock]
//   - [TTSExceptionCatcher.CatchExceptionError]
//   - [TTSRegex.EnumerateMatchesInCStringLengthUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringRangesUsingBlock]
//   - [TTSRegex.EnumerateMatchesInCStringStartOffsetLengthUsingBlock]
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
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackSynth]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynth]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
