package main

import (
	"fmt"
	"strings"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	pavfaudio "github.com/tmc/apple/private/avfaudio"
	"github.com/tmc/apple/private/texttospeech"
)

// sayPrivate speaks text using AVSpeechSynthesizer with private/Siri voices.
// It uses _speechVoicesIncludingSiri to resolve voices and setIsInternalSynth:
// to allow AVSpeechSynthesizer to use them.
func sayPrivate(text, voiceID string) error {
	utterance := avfaudio.NewSpeechUtteranceWithString(text)
	if utterance.ID == 0 {
		return fmt.Errorf("create utterance")
	}
	logf("utterance: %q", text)

	if voiceID != "" {
		voice, err := resolvePrivateVoice(voiceID)
		if err != nil {
			return err
		}
		utterance.SetVoice(voice)
		logf("voice: %s", voiceID)
	}

	synth := avfaudio.NewAVSpeechSynthesizer()
	// Required: without this the synthesizer silently ignores private/Siri voices.
	pavfaudio.AVSpeechSynthesizerFromID(synth.ID).SetIsInternalSynth(true)

	done := make(chan struct{})
	delegate := avfaudio.NewAVSpeechSynthesizerDelegate(avfaudio.AVSpeechSynthesizerDelegateConfig{
		SpeechSynthesizerDidFinishSpeechUtterance: func(_ avfaudio.AVSpeechSynthesizer, _ avfaudio.AVSpeechUtterance) {
			logf("finished")
			close(done)
		},
		SpeechSynthesizerDidCancelSpeechUtterance: func(_ avfaudio.AVSpeechSynthesizer, _ avfaudio.AVSpeechUtterance) {
			logf("cancelled")
			close(done)
		},
	})
	synth.SetDelegate(delegate)
	synth.SpeakUtterance(utterance)

	for {
		select {
		case <-done:
			return nil
		default:
			corefoundation.CFRunLoopRunInMode(corefoundation.KCFRunLoopDefaultMode, 0.1, false)
		}
	}
}

// resolvePrivateVoice finds an AVSpeechSynthesisVoice by identifier from the
// private voice list (which includes Siri voices).
func resolvePrivateVoice(voiceID string) (avfaudio.AVSpeechSynthesisVoice, error) {
	arr := privateVoices()
	for i := range arr.Count() {
		voice := avfaudio.AVSpeechSynthesisVoiceFromID(arr.ObjectAtIndex(i).GetID())
		if voice.Identifier() == voiceID {
			return voice, nil
		}
	}
	return avfaudio.AVSpeechSynthesisVoice{}, fmt.Errorf("voice %q not found", voiceID)
}

// findSiriVoice returns the identifier of the first available Siri voice.
func findSiriVoice() (string, error) {
	arr := privateVoices()
	for i := range arr.Count() {
		voice := avfaudio.AVSpeechSynthesisVoiceFromID(arr.ObjectAtIndex(i).GetID())
		id := voice.Identifier()
		idLower := strings.ToLower(id)
		if strings.Contains(idLower, "siri") || strings.Contains(idLower, "gryphon") {
			logf("found siri voice: %s", id)
			return id, nil
		}
	}
	return "", fmt.Errorf("no siri voice installed")
}

// privateVoices returns all voices including Siri via the private API.
func privateVoices() foundation.NSArray {
	obj := pavfaudio.GetAVSpeechSynthesisVoiceClass().SpeechVoicesIncludingSiri()
	return foundation.NSArrayFromID(obj.GetID())
}

// previewVoice speaks a sample phrase using the given voice identifier.
func previewVoice(voiceID string) error {
	voice, err := resolvePrivateVoice(voiceID)
	if err != nil {
		return err
	}
	lang := voice.Language()
	sample := samplePhrase(lang)
	logf("preview: voice=%s lang=%s", voiceID, lang)
	return sayPrivate(sample, voiceID)
}

// samplePhrase returns a short demo phrase for a language.
func samplePhrase(lang string) string {
	// Use the primary language subtag.
	primary := lang
	if i := strings.IndexByte(lang, '-'); i > 0 {
		primary = lang[:i]
	}
	switch primary {
	case "zh":
		return "你好，这是一个语音预览。"
	case "ja":
		return "こんにちは、これは音声プレビューです。"
	case "ko":
		return "안녕하세요, 음성 미리듣기입니다."
	case "de":
		return "Hallo, dies ist eine Sprachvorschau."
	case "fr":
		return "Bonjour, ceci est un aperçu vocal."
	case "es":
		return "Hola, esta es una vista previa de voz."
	case "it":
		return "Ciao, questa è un'anteprima vocale."
	case "pt":
		return "Olá, esta é uma prévia de voz."
	case "ru":
		return "Здравствуйте, это предварительный просмотр голоса."
	case "ar":
		return "مرحبا، هذه معاينة صوتية."
	case "he":
		return "שלום, זו תצוגה מקדימה של קול."
	case "nl":
		return "Hallo, dit is een spraakvoorbeeld."
	case "sv":
		return "Hej, det här är en röstförhandsvisning."
	case "da":
		return "Hej, dette er en stemmeforhåndsvisning."
	case "nb", "no":
		return "Hei, dette er en taleforhåndsvisning."
	case "fi":
		return "Hei, tämä on äänen esikatselu."
	case "tr":
		return "Merhaba, bu bir ses önizlemesidir."
	case "th":
		return "สวัสดี นี่คือตัวอย่างเสียง"
	case "vi":
		return "Xin chào, đây là bản xem trước giọng nói."
	case "ms":
		return "Hai, ini ialah pratonton suara."
	default:
		return "Hello! This is a preview of how I sound. I can read anything you type."
	}
}

// listCatalogVoices lists the full voice catalog (726+ voices) including
// voices that are not yet downloaded, using the private TextToSpeech framework.
func listCatalogVoices(siriOnly bool) {
	shim := texttospeech.NewTextToSpeechCoreSynthesisVoiceShim()
	obj := shim.ResourceVoicesWithOnlyInstalled(false)
	arr := foundation.NSArrayFromID(obj.GetID())
	for i := range arr.Count() {
		item := arr.ObjectAtIndex(i).GetID()
		idRv := objc.Send[objc.ID](item, objc.Sel("identifier"))
		if idRv == 0 {
			continue
		}
		identifier := foundation.NSStringFromID(idRv).String()
		if siriOnly {
			lower := strings.ToLower(identifier)
			if !strings.Contains(lower, "siri") && !strings.Contains(lower, "gryphon") {
				continue
			}
		}
		nameRv := objc.Send[objc.ID](item, objc.Sel("name"))
		name := ""
		if nameRv != 0 {
			name = foundation.NSStringFromID(nameRv).String()
		}
		langRv := objc.Send[objc.ID](item, objc.Sel("language"))
		lang := ""
		if langRv != 0 {
			lang = foundation.NSStringFromID(langRv).String()
		}
		installed := objc.Send[bool](item, objc.Sel("isInstalled"))
		status := " "
		if installed {
			status = "*"
		}
		fmt.Printf("%s %-8s %-30s %s\n", status, lang, name, identifier)
	}
}

// listPrivateVoices lists voices from the private voice list.
func listPrivateVoices(siriOnly bool) {
	arr := privateVoices()
	for i := range arr.Count() {
		voice := avfaudio.AVSpeechSynthesisVoiceFromID(arr.ObjectAtIndex(i).GetID())
		id := voice.Identifier()
		if siriOnly {
			idLower := strings.ToLower(id)
			if !strings.Contains(idLower, "siri") && !strings.Contains(idLower, "gryphon") {
				continue
			}
		}
		fmt.Printf("%-8s %-30s %s\n", voice.Language(), voice.Name(), id)
	}
}
