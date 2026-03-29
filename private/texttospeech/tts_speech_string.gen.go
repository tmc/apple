// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSpeechString] class.
var (
	_TTSSpeechStringClass     TTSSpeechStringClass
	_TTSSpeechStringClassOnce sync.Once
)

func getTTSSpeechStringClass() TTSSpeechStringClass {
	_TTSSpeechStringClassOnce.Do(func() {
		_TTSSpeechStringClass = TTSSpeechStringClass{class: objc.GetClass("TTSSpeechString")}
	})
	return _TTSSpeechStringClass
}

// GetTTSSpeechStringClass returns the class object for TTSSpeechString.
func GetTTSSpeechStringClass() TTSSpeechStringClass {
	return getTTSSpeechStringClass()
}

type TTSSpeechStringClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechStringClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechStringClass) Alloc() TTSSpeechString {
	rv := objc.Send[TTSSpeechString](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSSpeechString._insertTransformationForEncapsulatedTerminator]
//   - [TTSSpeechString._rangeIsValid]
//   - [TTSSpeechString._transformedStringNonMutating]
//   - [TTSSpeechString._translateRangeInTransformedStringWithParent]
//   - [TTSSpeechString.DefrostedTransformedString]
//   - [TTSSpeechString.EncapsulateSubstringAtRangeWithPrefixAndSuffix]
//   - [TTSSpeechString.Finalized]
//   - [TTSSpeechString.InsertAtLocationString]
//   - [TTSSpeechString.OriginalString]
//   - [TTSSpeechString.SetOriginalString]
//   - [TTSSpeechString.ParentString]
//   - [TTSSpeechString.SetParentString]
//   - [TTSSpeechString.ReplaceOccurencesOfStringWithString]
//   - [TTSSpeechString.TransformRangeTo]
//   - [TTSSpeechString.Transformations]
//   - [TTSSpeechString.SetTransformations]
//   - [TTSSpeechString.TransformedString]
//   - [TTSSpeechString.SetTransformedString]
//   - [TTSSpeechString.TranslateRangeInTransformedString]
//   - [TTSSpeechString.Type]
//   - [TTSSpeechString.SetType]
//   - [TTSSpeechString.XmlEscaped]
//   - [TTSSpeechString.XmlUnescaped]
//   - [TTSSpeechString.InitWithOriginalString]
//   - [TTSSpeechString.InitWithParentSpeechString]
//   - [TTSSpeechString.InitWithSSMLString]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString
type TTSSpeechString struct {
	objectivec.Object
}

// TTSSpeechStringFromID constructs a [TTSSpeechString] from an objc.ID.
func TTSSpeechStringFromID(id objc.ID) TTSSpeechString {
	return TTSSpeechString{objectivec.Object{ID: id}}
}
// Ensure TTSSpeechString implements ITTSSpeechString.
var _ ITTSSpeechString = TTSSpeechString{}

// An interface definition for the [TTSSpeechString] class.
//
// # Methods
//
//   - [ITTSSpeechString._insertTransformationForEncapsulatedTerminator]
//   - [ITTSSpeechString._rangeIsValid]
//   - [ITTSSpeechString._transformedStringNonMutating]
//   - [ITTSSpeechString._translateRangeInTransformedStringWithParent]
//   - [ITTSSpeechString.DefrostedTransformedString]
//   - [ITTSSpeechString.EncapsulateSubstringAtRangeWithPrefixAndSuffix]
//   - [ITTSSpeechString.Finalized]
//   - [ITTSSpeechString.InsertAtLocationString]
//   - [ITTSSpeechString.OriginalString]
//   - [ITTSSpeechString.SetOriginalString]
//   - [ITTSSpeechString.ParentString]
//   - [ITTSSpeechString.SetParentString]
//   - [ITTSSpeechString.ReplaceOccurencesOfStringWithString]
//   - [ITTSSpeechString.TransformRangeTo]
//   - [ITTSSpeechString.Transformations]
//   - [ITTSSpeechString.SetTransformations]
//   - [ITTSSpeechString.TransformedString]
//   - [ITTSSpeechString.SetTransformedString]
//   - [ITTSSpeechString.TranslateRangeInTransformedString]
//   - [ITTSSpeechString.Type]
//   - [ITTSSpeechString.SetType]
//   - [ITTSSpeechString.XmlEscaped]
//   - [ITTSSpeechString.XmlUnescaped]
//   - [ITTSSpeechString.InitWithOriginalString]
//   - [ITTSSpeechString.InitWithParentSpeechString]
//   - [ITTSSpeechString.InitWithSSMLString]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString
type ITTSSpeechString interface {
	objectivec.IObject

	// Topic: Methods

	_insertTransformationForEncapsulatedTerminator(transformation objectivec.IObject, terminator bool)
	_rangeIsValid(valid foundation.NSRange) bool
	_transformedStringNonMutating() objectivec.IObject
	_translateRangeInTransformedStringWithParent(string_ foundation.NSRange, parent objectivec.IObject) foundation.NSRange
	DefrostedTransformedString() string
	EncapsulateSubstringAtRangeWithPrefixAndSuffix(range_ foundation.NSRange, prefix objectivec.IObject, suffix objectivec.IObject) bool
	Finalized() bool
	InsertAtLocationString(location uint64, string_ objectivec.IObject) bool
	OriginalString() string
	SetOriginalString(value string)
	ParentString() ITTSSpeechString
	SetParentString(value ITTSSpeechString)
	ReplaceOccurencesOfStringWithString(string_ objectivec.IObject, string_2 objectivec.IObject)
	TransformRangeTo(range_ foundation.NSRange, to objectivec.IObject) bool
	Transformations() foundation.INSArray
	SetTransformations(value foundation.INSArray)
	TransformedString() string
	SetTransformedString(value string)
	TranslateRangeInTransformedString(string_ foundation.NSRange) foundation.NSRange
	Type() uint64
	SetType(value uint64)
	XmlEscaped() objectivec.IObject
	XmlUnescaped() objectivec.IObject
	InitWithOriginalString(string_ objectivec.IObject) TTSSpeechString
	InitWithParentSpeechString(string_ objectivec.IObject) TTSSpeechString
	InitWithSSMLString(sSMLString objectivec.IObject) TTSSpeechString
}

// Init initializes the instance.
func (t TTSSpeechString) Init() TTSSpeechString {
	rv := objc.Send[TTSSpeechString](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechString) Autorelease() TTSSpeechString {
	rv := objc.Send[TTSSpeechString](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechString creates a new TTSSpeechString instance.
func NewTTSSpeechString() TTSSpeechString {
	class := getTTSSpeechStringClass()
	rv := objc.Send[TTSSpeechString](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/initWithOriginalString:
func NewTTSSpeechStringWithOriginalString(string_ objectivec.IObject) TTSSpeechString {
	instance := getTTSSpeechStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOriginalString:"), string_)
	return TTSSpeechStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/initWithParentSpeechString:
func NewTTSSpeechStringWithParentSpeechString(string_ objectivec.IObject) TTSSpeechString {
	instance := getTTSSpeechStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParentSpeechString:"), string_)
	return TTSSpeechStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/initWithSSMLString:
func NewTTSSpeechStringWithSSMLString(sSMLString objectivec.IObject) TTSSpeechString {
	instance := getTTSSpeechStringClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSSMLString:"), sSMLString)
	return TTSSpeechStringFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/_insertTransformation:forEncapsulatedTerminator:
func (t TTSSpeechString) _insertTransformationForEncapsulatedTerminator(transformation objectivec.IObject, terminator bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("_insertTransformation:forEncapsulatedTerminator:"), transformation, terminator)
}

// InsertTransformationForEncapsulatedTerminator is an exported wrapper for the private method _insertTransformationForEncapsulatedTerminator.
func (t TTSSpeechString) InsertTransformationForEncapsulatedTerminator(transformation objectivec.IObject, terminator bool) {
	t._insertTransformationForEncapsulatedTerminator(transformation, terminator)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/_rangeIsValid:
func (t TTSSpeechString) _rangeIsValid(valid foundation.NSRange) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_rangeIsValid:"), valid)
	return rv
}

// RangeIsValid is an exported wrapper for the private method _rangeIsValid.
func (t TTSSpeechString) RangeIsValid(valid foundation.NSRange) bool {
	return t._rangeIsValid(valid)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/_transformedStringNonMutating
func (t TTSSpeechString) _transformedStringNonMutating() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_transformedStringNonMutating"))
	return objectivec.Object{ID: rv}
}

// TransformedStringNonMutating is an exported wrapper for the private method _transformedStringNonMutating.
func (t TTSSpeechString) TransformedStringNonMutating() objectivec.IObject {
	return t._transformedStringNonMutating()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/_translateRangeInTransformedString:withParent:
func (t TTSSpeechString) _translateRangeInTransformedStringWithParent(string_ foundation.NSRange, parent objectivec.IObject) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("_translateRangeInTransformedString:withParent:"), string_, parent)
	return foundation.NSRange(rv)
}

// TranslateRangeInTransformedStringWithParent is an exported wrapper for the private method _translateRangeInTransformedStringWithParent.
func (t TTSSpeechString) TranslateRangeInTransformedStringWithParent(string_ foundation.NSRange, parent objectivec.IObject) foundation.NSRange {
	return t._translateRangeInTransformedStringWithParent(string_, parent)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/encapsulateSubstringAtRange:withPrefix:andSuffix:
func (t TTSSpeechString) EncapsulateSubstringAtRangeWithPrefixAndSuffix(range_ foundation.NSRange, prefix objectivec.IObject, suffix objectivec.IObject) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("encapsulateSubstringAtRange:withPrefix:andSuffix:"), range_, prefix, suffix)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/insertAtLocation:string:
func (t TTSSpeechString) InsertAtLocationString(location uint64, string_ objectivec.IObject) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("insertAtLocation:string:"), location, string_)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/replaceOccurencesOfString:withString:
func (t TTSSpeechString) ReplaceOccurencesOfStringWithString(string_ objectivec.IObject, string_2 objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceOccurencesOfString:withString:"), string_, string_2)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/transformRange:to:
func (t TTSSpeechString) TransformRangeTo(range_ foundation.NSRange, to objectivec.IObject) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("transformRange:to:"), range_, to)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/translateRangeInTransformedString:
func (t TTSSpeechString) TranslateRangeInTransformedString(string_ foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("translateRangeInTransformedString:"), string_)
	return foundation.NSRange(rv)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/xmlEscaped
func (t TTSSpeechString) XmlEscaped() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("xmlEscaped"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/xmlUnescaped
func (t TTSSpeechString) XmlUnescaped() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("xmlUnescaped"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/initWithOriginalString:
func (t TTSSpeechString) InitWithOriginalString(string_ objectivec.IObject) TTSSpeechString {
	rv := objc.Send[TTSSpeechString](t.ID, objc.Sel("initWithOriginalString:"), string_)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/initWithParentSpeechString:
func (t TTSSpeechString) InitWithParentSpeechString(string_ objectivec.IObject) TTSSpeechString {
	rv := objc.Send[TTSSpeechString](t.ID, objc.Sel("initWithParentSpeechString:"), string_)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/initWithSSMLString:
func (t TTSSpeechString) InitWithSSMLString(sSMLString objectivec.IObject) TTSSpeechString {
	rv := objc.Send[TTSSpeechString](t.ID, objc.Sel("initWithSSMLString:"), sSMLString)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/defrostedTransformedString
func (t TTSSpeechString) DefrostedTransformedString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defrostedTransformedString"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/finalized
func (t TTSSpeechString) Finalized() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("finalized"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/originalString
func (t TTSSpeechString) OriginalString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("originalString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechString) SetOriginalString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setOriginalString:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/parentString
func (t TTSSpeechString) ParentString() ITTSSpeechString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("parentString"))
	return TTSSpeechStringFromID(objc.ID(rv))
}
func (t TTSSpeechString) SetParentString(value ITTSSpeechString) {
	objc.Send[struct{}](t.ID, objc.Sel("setParentString:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/transformations
func (t TTSSpeechString) Transformations() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("transformations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechString) SetTransformations(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setTransformations:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/transformedString
func (t TTSSpeechString) TransformedString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("transformedString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechString) SetTransformedString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setTransformedString:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechString/type
func (t TTSSpeechString) Type() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("type"))
	return rv
}
func (t TTSSpeechString) SetType(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setType:"), value)
}

