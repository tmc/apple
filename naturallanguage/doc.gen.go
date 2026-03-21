
// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

// Package naturallanguage provides Go bindings for the NaturalLanguage framework.
//
// Analyze natural language text and deduce its language-specific metadata.
//
// The Natural Language framework provides a variety of natural language processing (NLP) functionality with support for many different languages and scripts. Use this framework to segment natural language text into paragraphs, sentences, or words, and tag information about those segments, such as part of speech, lexical class, lemma, script, and language.
//
// # Tokenization
//
//   - Tokenizing natural language text: Enumerate the words in a string.
//   - NLTokenizer: A tokenizer that segments natural language text into semantic units.
//
// # Language identification
//
//   - Identifying the language in text: Detect the language in a piece of text by using a language recognizer.
//   - NLLanguageRecognizer: The language of a body of text.
//   - NLLanguage: The languages that the Natural Language framework supports.
//
// # Linguistic tags
//
//   - Identifying parts of speech: Classify nouns, verbs, adjectives, and other parts of speech in a string.
//   - Identifying people, places, and organizations: Use a linguistic tagger to perform named entity recognition on a string.
//   - NLTagger: A tagger that analyzes natural language text. ([NLTagScheme], [NLTag], [NLTokenUnit], [NLGazetteer])
//
// # Text embedding
//
//   - Finding similarities between pieces of text: Calculate the semantic distance between words or sentences.
//   - NLEmbedding: A map of strings to vectors, which locates neighboring, similar strings. ([NLDistance])
//
// # Contextual embedding
//
//   - NLContextualEmbedding: A model that computes sequences of embedding vectors for natural language utterances. ([NLContextualEmbeddingResult])
//   - NLContextualEmbeddingKey: Contextual embedding keys.
//   - NLScript: The writing scripts that the Natural Language framework supports.
//
// # Natural language models
//
//   - Creating a text classifier model: Train a machine learning model to classify natural language text.
//   - Creating a word tagger model: Train a machine learning model to tag individual words in natural language text.
//   - NLModel: A custom model trained to classify or tag natural language text. ([NLModelConfiguration])
//
// # Key Types
//
//   - [NLEmbedding] - A map of strings to vectors, which locates neighboring, similar strings.
//   - [NLTagger] - A tagger that analyzes natural language text.
//   - [NLContextualEmbedding] - A model that computes sequences of embedding vectors for natural language utterances.
//   - [NLGazetteer] - A collection of terms and their labels, which take precedence over a word tagger.
//   - [NLLanguageRecognizer] - The language of a body of text.
//   - [NLTokenizer] - A tokenizer that segments natural language text into semantic units.
//   - [NLModel] - A custom model trained to classify or tag natural language text.
//   - [NLModelConfiguration] - The configuration parameters of a natural language model.
//   - [NLContextualEmbeddingResult] - An object that represents the embedding vector result from applying a contextual embedding to a string.
//
// [NaturalLanguage Documentation]: https://developer.apple.com/documentation/NaturalLanguage
package naturallanguage

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/NaturalLanguage.framework/NaturalLanguage"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: NaturalLanguage: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

