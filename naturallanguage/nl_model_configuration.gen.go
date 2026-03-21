// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NLModelConfiguration] class.
var (
	_NLModelConfigurationClass     NLModelConfigurationClass
	_NLModelConfigurationClassOnce sync.Once
)

func getNLModelConfigurationClass() NLModelConfigurationClass {
	_NLModelConfigurationClassOnce.Do(func() {
		_NLModelConfigurationClass = NLModelConfigurationClass{class: objc.GetClass("NLModelConfiguration")}
	})
	return _NLModelConfigurationClass
}

// GetNLModelConfigurationClass returns the class object for NLModelConfiguration.
func GetNLModelConfigurationClass() NLModelConfigurationClass {
	return getNLModelConfigurationClass()
}

type NLModelConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NLModelConfigurationClass) Alloc() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The configuration parameters of a natural language model.
//
// # Accessing the configuration
//
//   - [NLModelConfiguration.Language]: The language the model supports.
//   - [NLModelConfiguration.Revision]: The version of the Natural Language framework that trained the model.
//   - [NLModelConfiguration.Type]: The natural language model type of the model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration
type NLModelConfiguration struct {
	objectivec.Object
}

// NLModelConfigurationFromID constructs a [NLModelConfiguration] from an objc.ID.
//
// The configuration parameters of a natural language model.
func NLModelConfigurationFromID(id objc.ID) NLModelConfiguration {
	return NLModelConfiguration{objectivec.Object{ID: id}}
}
// NOTE: NLModelConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NLModelConfiguration] class.
//
// # Accessing the configuration
//
//   - [INLModelConfiguration.Language]: The language the model supports.
//   - [INLModelConfiguration.Revision]: The version of the Natural Language framework that trained the model.
//   - [INLModelConfiguration.Type]: The natural language model type of the model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration
type INLModelConfiguration interface {
	objectivec.IObject

	// Topic: Accessing the configuration

	// The language the model supports.
	Language() NLLanguage
	// The version of the Natural Language framework that trained the model.
	Revision() uint
	// The natural language model type of the model.
	Type() NLModelType

	// A configuration describing the natural language model.
	Configuration() INLModelConfiguration
	SetConfiguration(value INLModelConfiguration)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m NLModelConfiguration) Init() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NLModelConfiguration) Autorelease() NLModelConfiguration {
	rv := objc.Send[NLModelConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLModelConfiguration creates a new NLModelConfiguration instance.
func NewNLModelConfiguration() NLModelConfiguration {
	class := getNLModelConfigurationClass()
	rv := objc.Send[NLModelConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m NLModelConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the versions of the Natural Language framework the OS supports.
//
// # Return Value
// 
// An index set of version numbers.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/supportedRevisions(for:)
func (_NLModelConfigurationClass NLModelConfigurationClass) SupportedRevisionsForType(type_ NLModelType) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](objc.ID(_NLModelConfigurationClass.class), objc.Sel("supportedRevisionsForType:"), type_)
	return foundation.NSIndexSetFromID(rv)
}
// Returns the current Natural Language framework version in the OS.
//
// # Return Value
// 
// A version number.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/currentRevision(for:)
func (_NLModelConfigurationClass NLModelConfigurationClass) CurrentRevisionForType(type_ NLModelType) uint {
	rv := objc.Send[uint](objc.ID(_NLModelConfigurationClass.class), objc.Sel("currentRevisionForType:"), type_)
	return rv
}

// The language the model supports.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/language
func (m NLModelConfiguration) Language() NLLanguage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("language"))
	return NLLanguage(foundation.NSStringFromID(rv).String())
}
// The version of the Natural Language framework that trained the model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/revision
func (m NLModelConfiguration) Revision() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("revision"))
	return rv
}
// The natural language model type of the model.
//
// See: https://developer.apple.com/documentation/NaturalLanguage/NLModelConfiguration/type
func (m NLModelConfiguration) Type() NLModelType {
	rv := objc.Send[NLModelType](m.ID, objc.Sel("type"))
	return NLModelType(rv)
}
// A configuration describing the natural language model.
//
// See: https://developer.apple.com/documentation/naturallanguage/nlmodel/configuration
func (m NLModelConfiguration) Configuration() INLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return NLModelConfigurationFromID(objc.ID(rv))
}
func (m NLModelConfiguration) SetConfiguration(value INLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfiguration:"), value)
}

