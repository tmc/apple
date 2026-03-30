// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/AssetsResult
type NLContextualEmbeddingAssetsResult int

const (
	// NLContextualEmbeddingAssetsResultAvailable: A result that indicates that the assets are present on-device.
	NLContextualEmbeddingAssetsResultAvailable NLContextualEmbeddingAssetsResult = 0
	// NLContextualEmbeddingAssetsResultError: A result that indicates the framework encounters an error.
	NLContextualEmbeddingAssetsResultError NLContextualEmbeddingAssetsResult = 2
	// NLContextualEmbeddingAssetsResultNotAvailable: A result that indicates that the assets aren’t present on-device.
	NLContextualEmbeddingAssetsResultNotAvailable NLContextualEmbeddingAssetsResult = 1
)

func (e NLContextualEmbeddingAssetsResult) String() string {
	switch e {
	case NLContextualEmbeddingAssetsResultAvailable:
		return "NLContextualEmbeddingAssetsResultAvailable"
	case NLContextualEmbeddingAssetsResultError:
		return "NLContextualEmbeddingAssetsResultError"
	case NLContextualEmbeddingAssetsResultNotAvailable:
		return "NLContextualEmbeddingAssetsResultNotAvailable"
	default:
		return fmt.Sprintf("NLContextualEmbeddingAssetsResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NaturalLanguage/NLDistanceType
type NLDistanceType int

const (
	// NLDistanceTypeCosine: A method of calculating distance by using cosine similarity.
	NLDistanceTypeCosine NLDistanceType = 0
)

func (e NLDistanceType) String() string {
	switch e {
	case NLDistanceTypeCosine:
		return "NLDistanceTypeCosine"
	default:
		return fmt.Sprintf("NLDistanceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/ModelType
type NLModelType int

const (
	// NLModelTypeClassifier: A classifier model type that tags text at the phrase, sentence, paragraph, or higher level.
	NLModelTypeClassifier NLModelType = 0
	// NLModelTypeSequence: A sequence model type that tags text at the token level.
	NLModelTypeSequence NLModelType = 1
)

func (e NLModelType) String() string {
	switch e {
	case NLModelTypeClassifier:
		return "NLModelTypeClassifier"
	case NLModelTypeSequence:
		return "NLModelTypeSequence"
	default:
		return fmt.Sprintf("NLModelType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/AssetsResult
type NLTaggerAssetsResult int

const (
	// NLTaggerAssetsResultAvailable: The asset is now available and loaded onto the device.
	NLTaggerAssetsResultAvailable NLTaggerAssetsResult = 0
	// NLTaggerAssetsResultError: The framework couldn’t load the asset due to an error.
	NLTaggerAssetsResultError NLTaggerAssetsResult = 2
	// NLTaggerAssetsResultNotAvailable: The asset is unavailable on the device.
	NLTaggerAssetsResultNotAvailable NLTaggerAssetsResult = 1
)

func (e NLTaggerAssetsResult) String() string {
	switch e {
	case NLTaggerAssetsResultAvailable:
		return "NLTaggerAssetsResultAvailable"
	case NLTaggerAssetsResultError:
		return "NLTaggerAssetsResultError"
	case NLTaggerAssetsResultNotAvailable:
		return "NLTaggerAssetsResultNotAvailable"
	default:
		return fmt.Sprintf("NLTaggerAssetsResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NaturalLanguage/NLTagger/Options
type NLTaggerOptions uint

const (
	// NLTaggerJoinContractions: Contractions will be returned as one token.
	NLTaggerJoinContractions NLTaggerOptions = 32
	// NLTaggerJoinNames: Typically, multiple-word names will be returned as multiple tokens, following the standard tokenization practice of the tagger.
	NLTaggerJoinNames NLTaggerOptions = 16
	// NLTaggerOmitOther: Omit tokens of type other (non-linguistic items, such as symbols).
	NLTaggerOmitOther NLTaggerOptions = 8
	// NLTaggerOmitPunctuation: Omit tokens of type punctuation (all punctuation).
	NLTaggerOmitPunctuation NLTaggerOptions = 2
	// NLTaggerOmitWhitespace: Omit tokens of type whitespace (whitespace of all sorts).
	NLTaggerOmitWhitespace NLTaggerOptions = 4
	// NLTaggerOmitWords: Omit tokens of type word (items considered to be words).
	NLTaggerOmitWords NLTaggerOptions = 1
)

func (e NLTaggerOptions) String() string {
	switch e {
	case NLTaggerJoinContractions:
		return "NLTaggerJoinContractions"
	case NLTaggerJoinNames:
		return "NLTaggerJoinNames"
	case NLTaggerOmitOther:
		return "NLTaggerOmitOther"
	case NLTaggerOmitPunctuation:
		return "NLTaggerOmitPunctuation"
	case NLTaggerOmitWhitespace:
		return "NLTaggerOmitWhitespace"
	case NLTaggerOmitWords:
		return "NLTaggerOmitWords"
	default:
		return fmt.Sprintf("NLTaggerOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenUnit
type NLTokenUnit int

const (
	// NLTokenUnitDocument: The document in its entirety.
	NLTokenUnitDocument NLTokenUnit = 3
	// NLTokenUnitParagraph: An individual paragraph.
	NLTokenUnitParagraph NLTokenUnit = 2
	// NLTokenUnitSentence: An individual sentence.
	NLTokenUnitSentence NLTokenUnit = 1
	// NLTokenUnitWord: An individual word.
	NLTokenUnitWord NLTokenUnit = 0
)

func (e NLTokenUnit) String() string {
	switch e {
	case NLTokenUnitDocument:
		return "NLTokenUnitDocument"
	case NLTokenUnitParagraph:
		return "NLTokenUnitParagraph"
	case NLTokenUnitSentence:
		return "NLTokenUnitSentence"
	case NLTokenUnitWord:
		return "NLTokenUnitWord"
	default:
		return fmt.Sprintf("NLTokenUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/Attributes
type NLTokenizerAttributes uint

const (
	// NLTokenizerAttributeEmoji: The string contains emoji.
	NLTokenizerAttributeEmoji NLTokenizerAttributes = 4
	// NLTokenizerAttributeNumeric: The string contains numbers.
	NLTokenizerAttributeNumeric NLTokenizerAttributes = 1
	// NLTokenizerAttributeSymbolic: The string contains symbols.
	NLTokenizerAttributeSymbolic NLTokenizerAttributes = 2
)

func (e NLTokenizerAttributes) String() string {
	switch e {
	case NLTokenizerAttributeEmoji:
		return "NLTokenizerAttributeEmoji"
	case NLTokenizerAttributeNumeric:
		return "NLTokenizerAttributeNumeric"
	case NLTokenizerAttributeSymbolic:
		return "NLTokenizerAttributeSymbolic"
	default:
		return fmt.Sprintf("NLTokenizerAttributes(%d)", e)
	}
}
