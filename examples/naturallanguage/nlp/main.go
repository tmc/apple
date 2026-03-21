// Command nlp provides Unix-native text analysis using Apple's NaturalLanguage framework.
//
// Usage:
//
//	nlp <command> [flags] [file | -]
//
// Commands:
//
//	language     Detect dominant language
//	tokens       Tokenize text (words, sentences, or paragraphs)
//	entities     Extract named entities (people, places, organizations)
//	sentiment    Analyze sentiment (-1.0 to 1.0)
//	lemma        Get base forms of words
//
// Text can be a file path, "-" for stdin, or passed via -text flag.
// Output is JSON.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/naturallanguage"
	"github.com/tmc/apple/objc"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "language":
		err = runLanguage(args)
	case "tokens":
		err = runTokens(args)
	case "entities":
		err = runEntities(args)
	case "sentiment":
		err = runSentiment(args)
	case "lemma":
		err = runLemma(args)
	case "-h", "-help", "--help", "help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "nlp: unknown command %q\n", cmd)
		usage()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "nlp %s: %v\n", cmd, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: nlp <command> [flags] [file | -]

Commands:
  language     Detect dominant language
  tokens       Tokenize text (words, sentences, or paragraphs)
  entities     Extract named entities (people, places, organizations)
  sentiment    Analyze sentiment (-1.0 to 1.0)
  lemma        Get base forms of words

Text can be a file path, "-" for stdin, or passed via -text flag.
Output is JSON.
`)
	os.Exit(2)
}

// --- Commands ---

type languageResult struct {
	Language string `json:"language"`
	BCP47    string `json:"bcp47"`
}

func runLanguage(args []string) error {
	fs := flag.NewFlagSet("language", flag.ExitOnError)
	textFlag := fs.String("text", "", "text to analyze")
	fs.Parse(args)

	text, err := loadText(fs.Args(), textFlag)
	if err != nil {
		return err
	}

	lang := naturallanguage.GetNLLanguageRecognizerClass().DominantLanguageForString(text)
	return writeJSON(languageResult{
		Language: string(lang),
		BCP47:    string(lang),
	})
}

func runTokens(args []string) error {
	fs := flag.NewFlagSet("tokens", flag.ExitOnError)
	textFlag := fs.String("text", "", "text to analyze")
	unitFlag := fs.String("unit", "word", "tokenization unit: word, sentence, paragraph")
	fs.Parse(args)

	text, err := loadText(fs.Args(), textFlag)
	if err != nil {
		return err
	}

	unit := naturallanguage.NLTokenUnitWord
	switch *unitFlag {
	case "sentence":
		unit = naturallanguage.NLTokenUnitSentence
	case "paragraph":
		unit = naturallanguage.NLTokenUnitParagraph
	}

	opts := int(naturallanguage.NLTaggerOmitPunctuation | naturallanguage.NLTaggerOmitWhitespace)
	results := enumerateTags(text, naturallanguage.NLTagSchemes.TokenType, unit, opts)
	return writeJSON(results)
}

func runEntities(args []string) error {
	fs := flag.NewFlagSet("entities", flag.ExitOnError)
	textFlag := fs.String("text", "", "text to analyze")
	fs.Parse(args)

	text, err := loadText(fs.Args(), textFlag)
	if err != nil {
		return err
	}

	opts := int(naturallanguage.NLTaggerOmitPunctuation | naturallanguage.NLTaggerOmitWhitespace |
		naturallanguage.NLTaggerOmitOther | naturallanguage.NLTaggerJoinNames)
	results := enumerateTags(text, naturallanguage.NLTagSchemes.NameType, naturallanguage.NLTokenUnitWord, opts)

	var entities []tagResult
	for _, r := range results {
		switch r.Tag {
		case "PersonalName", "PlaceName", "OrganizationName":
			entities = append(entities, r)
		}
	}
	if entities == nil {
		entities = []tagResult{}
	}
	return writeJSON(entities)
}

type sentimentResult struct {
	Text  string  `json:"text"`
	Score float64 `json:"score"`
}

func runSentiment(args []string) error {
	fs := flag.NewFlagSet("sentiment", flag.ExitOnError)
	textFlag := fs.String("text", "", "text to analyze")
	fs.Parse(args)

	text, err := loadText(fs.Args(), textFlag)
	if err != nil {
		return err
	}

	results := enumerateTags(text, naturallanguage.NLTagSchemes.SentimentScore, naturallanguage.NLTokenUnitSentence, 0)

	var sentiments []sentimentResult
	for _, r := range results {
		score, _ := strconv.ParseFloat(r.Tag, 64)
		sentiments = append(sentiments, sentimentResult{
			Text:  r.Text,
			Score: score,
		})
	}
	if sentiments == nil {
		sentiments = []sentimentResult{}
	}
	return writeJSON(sentiments)
}

type lemmaResult struct {
	Word  string `json:"word"`
	Lemma string `json:"lemma"`
}

func runLemma(args []string) error {
	fs := flag.NewFlagSet("lemma", flag.ExitOnError)
	textFlag := fs.String("text", "", "text to analyze")
	fs.Parse(args)

	text, err := loadText(fs.Args(), textFlag)
	if err != nil {
		return err
	}

	opts := int(naturallanguage.NLTaggerOmitPunctuation | naturallanguage.NLTaggerOmitWhitespace)
	results := enumerateTags(text, naturallanguage.NLTagSchemes.Lemma, naturallanguage.NLTokenUnitWord, opts)

	var lemmas []lemmaResult
	for _, r := range results {
		lemma := r.Tag
		if lemma == "" {
			lemma = r.Text
		}
		lemmas = append(lemmas, lemmaResult{
			Word:  r.Text,
			Lemma: lemma,
		})
	}
	if lemmas == nil {
		lemmas = []lemmaResult{}
	}
	return writeJSON(lemmas)
}

// --- Tagger ---

type tagResult struct {
	Text  string `json:"text"`
	Tag   string `json:"tag,omitempty"`
	Start int    `json:"start"`
	Len   int    `json:"length"`
}

// enumerateTags runs the NLTagger enumerateTagsInRange:unit:scheme:options:usingBlock:
// method via raw objc.Send since the block-based enumeration isn't generated.
//
// The block callback receives (NLTag, NSRange, BOOL*) but purego can't pass
// structs through blocks, so NSRange is flattened to (location, length) uintptrs.
func enumerateTags(text string, scheme string, unit naturallanguage.NLTokenUnit, opts int) []tagResult {
	tagger := naturallanguage.NewTaggerWithTagSchemes([]string{scheme})
	tagger.SetString(text)

	var results []tagResult
	runes := []rune(text)
	textLen := len(runes)

	// Block signature: void(^)(NLTag tag, NSRange.location, NSRange.length, BOOL *stop)
	block := objc.NewBlock(func(_ objc.Block, tagID objc.ID, location uintptr, length uintptr, _ unsafe.Pointer) {
		tag := ""
		if tagID != 0 {
			tag = foundation.NSStringFromID(tagID).String()
		}
		start := int(location)
		ln := int(length)
		end := min(start+ln, len(runes))
		if start > len(runes) {
			return
		}
		substr := string(runes[start:end])

		results = append(results, tagResult{
			Text:  substr,
			Tag:   tag,
			Start: start,
			Len:   ln,
		})
	})
	defer block.Release()

	// NSRange is passed as two register values on arm64 (location, length).
	objc.Send[objc.ID](tagger.ID, objc.Sel("enumerateTagsInRange:unit:scheme:options:usingBlock:"),
		uintptr(0), uintptr(textLen),
		unit,
		objc.String(scheme),
		uintptr(opts),
		block,
	)

	return results
}

// --- I/O ---

func loadText(args []string, textFlag *string) (string, error) {
	if textFlag != nil && *textFlag != "" {
		return *textFlag, nil
	}
	var r io.Reader
	switch {
	case len(args) == 0 || args[0] == "-":
		r = os.Stdin
	default:
		f, err := os.Open(args[0])
		if err != nil {
			return "", fmt.Errorf("open %s: %w", args[0], err)
		}
		defer f.Close()
		r = f
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}
	return string(data), nil
}

func writeJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
