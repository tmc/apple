// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLModel] class.
var (
	_NLModelClass     NLModelClass
	_NLModelClassOnce sync.Once
)

func getNLModelClass() NLModelClass {
	_NLModelClassOnce.Do(func() {
		_NLModelClass = NLModelClass{class: objc.GetClass("NLModel")}
	})
	return _NLModelClass
}

// GetNLModelClass returns the class object for NLModel.
func GetNLModelClass() NLModelClass {
	return getNLModelClass()
}

type NLModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NLModelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLModelClass) Alloc() NLModel {
	rv := objc.Send[NLModel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A custom model trained to classify or tag natural language text.
//
// # Overview
// 
// With [Natural Language], you can create text classifier
// ([MLTextClassifier]) or word tagger ([MLWordTagger]) models. Use [NLModel]
// to integrate those models into your app. This integration ensures that your
// tokenization and tagger configurations are identical when you train your
// model and use it in your app.
// 
// If you create a text classifier as described in
// doc:creating-a-text-classifier-model, you can integrate that model into
// your app and use it to make predictions like this:
// 
// If you create a custom word tagger as described in
// doc:creating-a-word-tagger-model, you can integrate that model into your
// app and generate tags for new text input like this:
//
// [MLTextClassifier]: https://developer.apple.com/documentation/CreateML/MLTextClassifier
// [MLWordTagger]: https://developer.apple.com/documentation/CreateML/MLWordTagger
// [Natural Language]: https://developer.apple.com/documentation/NaturalLanguage
//
// # Making predictions
//
//   - [NLModel.PredictedLabelForString]: Predicts a label for the given input string.
//   - [NLModel.PredictedLabelsForTokens]: Predicts a label for each string in the given array.
//
// # Inspecting a model
//
//   - [NLModel.Configuration]: A configuration describing the natural language model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel
type NLModel struct {
	objectivec.Object
}

// NLModelFromID constructs a [NLModel] from an objc.ID.
//
// A custom model trained to classify or tag natural language text.
func NLModelFromID(id objc.ID) NLModel {
	return NLModel{objectivec.Object{ID: id}}
}
// NOTE: NLModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLModel] class.
//
// # Making predictions
//
//   - [INLModel.PredictedLabelForString]: Predicts a label for the given input string.
//   - [INLModel.PredictedLabelsForTokens]: Predicts a label for each string in the given array.
//
// # Inspecting a model
//
//   - [INLModel.Configuration]: A configuration describing the natural language model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel
type INLModel interface {
	objectivec.IObject

	// Topic: Making predictions

	// Predicts a label for the given input string.
	PredictedLabelForString(string_ string) string
	// Predicts a label for each string in the given array.
	PredictedLabelsForTokens(tokens []string) []string

	// Topic: Inspecting a model

	// A configuration describing the natural language model.
	Configuration() INLModelConfiguration

	// Predicts multiple possible labels for the given input string.
	PredictedLabelHypothesesForStringMaximumCount(string_ string, maximumCount uint) foundation.INSDictionary
	// Predicts multiple possible labels for each string in the given array.
	PredictedLabelHypothesesForTokensMaximumCount(tokens []string, maximumCount uint) foundation.INSDictionary
}

// Init initializes the instance.
func (m NLModel) Init() NLModel {
	rv := objc.Send[NLModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NLModel) Autorelease() NLModel {
	rv := objc.Send[NLModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLModel creates a new NLModel instance.
func NewNLModel() NLModel {
	class := getNLModelClass()
	rv := objc.Send[NLModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new natural language model based on a compiled Core ML model at
// the given URL.
//
// url: The location of the Core ML model file in the file system (ending with
// `XCUIElementTypeMlmodelc`) that’s the basis for this natural language
// model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(contentsOf:)
func NewModelWithContentsOfURLError(url foundation.INSURL) (NLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getNLModelClass().class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLModelFromID(rv), nil
}

// Creates a new natural language model based on the given Core ML model
// instance.
//
// mlModel: A Core ML model instance that’s the basis for this natural language
// model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/init(mlModel:)
func NewModelWithMLModelError(mlModel coreml.MLModel) (NLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getNLModelClass().class), objc.Sel("modelWithMLModel:error:"), mlModel, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return NLModelFromID(rv), nil
}

// Predicts a label for the given input string.
//
// string: The input text for the model to analyze.
//
// # Return Value
// 
// The model’s predicted label for the input string.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabel(for:)
func (m NLModel) PredictedLabelForString(string_ string) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedLabelForString:"), objc.String(string_))
	return foundation.NSStringFromID(rv).String()
}
// Predicts a label for each string in the given array.
//
// tokens: An array of input strings for the model to analyze.
//
// # Return Value
// 
// The model’s predicted labels for each of the input strings.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabels(forTokens:)
func (m NLModel) PredictedLabelsForTokens(tokens []string) []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("predictedLabelsForTokens:"), objectivec.StringSliceToNSArray(tokens))
	return objc.ConvertSliceToStrings(rv)
}
// Predicts multiple possible labels for the given input string.
//
// string: The input string for the model to analyze.
//
// maximumCount: The maximum number of label predictions to return for each input string.
//
// # Return Value
// 
// A dictionary of label hypotheses. Each dictionary entry is a predicted
// label with its associated probability score. These labels are the top
// candidates proposed as possible labels for the input string. The dictionary
// contains up to `maximumCount` entries.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForString:maximumCount:
func (m NLModel) PredictedLabelHypothesesForStringMaximumCount(string_ string, maximumCount uint) foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedLabelHypothesesForString:maximumCount:"), objc.String(string_), maximumCount)
	return foundation.NSDictionaryFromID(rv)
}
// Predicts multiple possible labels for each string in the given array.
//
// tokens: An array of input tokens for the model to analyze.
//
// maximumCount: The maximum number of label predictions to return for each input string.
//
// # Return Value
// 
// An array of dictionaries. Each dictionary corresponds to the token at the
// same index in the input array `tokens`. Within each dictionary, each entry
// is a predicted label with its associated probability score. These labels
// are the top candidates proposed as possible labels for the token. Each
// dictionary contains up to `maximumCount` entries.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/predictedLabelHypothesesForTokens:maximumCount:
func (m NLModel) PredictedLabelHypothesesForTokensMaximumCount(tokens []string, maximumCount uint) foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedLabelHypothesesForTokens:maximumCount:"), objectivec.StringSliceToNSArray(tokens), maximumCount)
	return foundation.NSDictionaryFromID(rv)
}

// A configuration describing the natural language model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModel/configuration
func (m NLModel) Configuration() INLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return NLModelConfigurationFromID(objc.ID(rv))
}

