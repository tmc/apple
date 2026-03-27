// Command audiosay speaks text using Apple's speech synthesis APIs.
//
// By default it uses AVSpeechSynthesizer (public API). With -siri it uses
// the private TextToSpeech framework to speak with a Siri voice.
//
// Usage:
//
//	audiosay [-voice lang] [-rate n] [-list] [text | -]
//	audiosay -siri [text | -]
//	audiosay -private [-voice id] [-list] [text | -]
//	audiosay -list-all [-siri]
//	audiosay -preview <voice-id>
//
// Examples:
//
//	audiosay "Hello, world!"
//	echo "pipe me" | audiosay -
//	audiosay -voice en-AU "Good day, mate"
//	audiosay -list
//	audiosay -rate 0.6 "Speaking slowly"
//	audiosay -siri "Hello from Siri"
//	audiosay -private -list
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/corefoundation"
)

var verbose bool

func logf(format string, args ...any) {
	if verbose {
		fmt.Fprintf(os.Stderr, "audiosay: "+format+"\n", args...)
	}
}

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	voice := flag.String("voice", "", "BCP 47 language code or voice identifier")
	rate := flag.Float64("rate", -1, "speech rate (0.0–1.0, default ~0.5)")
	list := flag.Bool("list", false, "list available voices and exit")
	listAll := flag.Bool("list-all", false, "list full voice catalog including not-yet-downloaded voices")
	preview := flag.String("preview", "", "preview a voice by speaking a sample phrase")
	private := flag.Bool("private", false, "use private TextToSpeech framework")
	siri := flag.Bool("siri", false, "speak with Siri voice (uses private framework)")
	flag.BoolVar(&verbose, "verbose", false, "print debug information")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: audiosay [-voice lang] [-rate n] [-list] [text | -]\n")
		fmt.Fprintf(os.Stderr, "       audiosay -siri [text | -]\n")
		fmt.Fprintf(os.Stderr, "       audiosay -private [-voice id] [-list] [text | -]\n")
		fmt.Fprintf(os.Stderr, "       audiosay -list-all [-siri]\n")
		fmt.Fprintf(os.Stderr, "       audiosay -preview <voice-id>\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *siri {
		*private = true
	}

	if *preview != "" {
		if err := previewVoice(*preview); err != nil {
			fmt.Fprintf(os.Stderr, "audiosay: %v\n", err)
			os.Exit(1)
		}
		return
	}
	if *listAll {
		listCatalogVoices(*siri)
		return
	}
	if *list {
		if *private {
			listPrivateVoices(false)
		} else {
			listVoices()
		}
		return
	}

	text, err := loadText(flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "audiosay: %v\n", err)
		os.Exit(1)
	}
	if text == "" {
		flag.Usage()
		os.Exit(2)
	}

	if *siri {
		voiceID := *voice
		if voiceID == "" {
			voiceID, err = findSiriVoice()
			if err != nil {
				fmt.Fprintf(os.Stderr, "audiosay: %v\n", err)
				os.Exit(1)
			}
		}
		err = sayPrivate(text, voiceID)
	} else if *private {
		err = sayPrivate(text, *voice)
	} else {
		err = say(text, *voice, *rate)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "audiosay: %v\n", err)
		os.Exit(1)
	}
}

func say(text, lang string, rate float64) error {
	synth := avfaudio.NewAVSpeechSynthesizer()
	utterance := avfaudio.NewSpeechUtteranceWithString(text)

	if lang != "" {
		v := avfaudio.NewSpeechSynthesisVoiceWithLanguage(lang)
		if v.ID == 0 {
			return fmt.Errorf("unknown voice %q", lang)
		}
		utterance.SetVoice(v)
	}
	if rate >= 0 {
		utterance.SetRate(float32(rate))
	}

	done := make(chan struct{})
	delegate := avfaudio.NewAVSpeechSynthesizerDelegate(avfaudio.AVSpeechSynthesizerDelegateConfig{
		SpeechSynthesizerDidFinishSpeechUtterance: func(_ avfaudio.AVSpeechSynthesizer, _ avfaudio.AVSpeechUtterance) {
			close(done)
		},
		SpeechSynthesizerDidCancelSpeechUtterance: func(_ avfaudio.AVSpeechSynthesizer, _ avfaudio.AVSpeechUtterance) {
			close(done)
		},
	})
	synth.SetDelegate(delegate)
	synth.SpeakUtterance(utterance)

	// Pump the run loop so delegate callbacks fire.
	for {
		select {
		case <-done:
			return nil
		default:
			corefoundation.CFRunLoopRunInMode(corefoundation.KCFRunLoopDefaultMode, 0.1, false)
		}
	}
}

func listVoices() {
	voices := avfaudio.GetAVSpeechSynthesisVoiceClass().SpeechVoices()
	for _, v := range voices {
		fmt.Printf("%-8s %s\n", v.Language(), v.Name())
	}
}

func loadText(args []string) (string, error) {
	if len(args) == 0 {
		return "", nil
	}
	if args[0] == "-" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("read stdin: %w", err)
		}
		return strings.TrimSpace(string(data)), nil
	}
	return strings.Join(args, " "), nil
}
