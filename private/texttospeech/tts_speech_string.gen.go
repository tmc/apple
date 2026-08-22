// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[TTSSpeechString](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

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
	rv := objc.SendIfResponds[TTSSpeechString](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechString) Autorelease() TTSSpeechString {
	rv := objc.SendIfResponds[TTSSpeechString](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechString creates a new TTSSpeechString instance.
func NewTTSSpeechString() TTSSpeechString {
	class := getTTSSpeechStringClass()
	rv := objc.SendIfResponds[TTSSpeechString](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTTSSpeechStringWithOriginalString(string_ objectivec.IObject) TTSSpeechString {
	instance := getTTSSpeechStringClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOriginalString:"), string_)
	return TTSSpeechStringFromID(rv)
}

func NewTTSSpeechStringWithParentSpeechString(string_ objectivec.IObject) TTSSpeechString {
	instance := getTTSSpeechStringClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParentSpeechString:"), string_)
	return TTSSpeechStringFromID(rv)
}

func NewTTSSpeechStringWithSSMLString(sSMLString objectivec.IObject) TTSSpeechString {
	instance := getTTSSpeechStringClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSSMLString:"), sSMLString)
	return TTSSpeechStringFromID(rv)
}

func (t TTSSpeechString) _insertTransformationForEncapsulatedTerminator(transformation objectivec.IObject, terminator bool) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_insertTransformation:forEncapsulatedTerminator:"), transformation, terminator)
}

// InsertTransformationForEncapsulatedTerminator is an exported wrapper for the private method _insertTransformationForEncapsulatedTerminator.
func (t TTSSpeechString) InsertTransformationForEncapsulatedTerminator(transformation objectivec.IObject, terminator bool) error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_insertTransformation:forEncapsulatedTerminator:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_insertTransformation:forEncapsulatedTerminator:"}
		return err
	}
	t._insertTransformationForEncapsulatedTerminator(transformation, terminator)
	return nil
}

// CanInsertTransformationForEncapsulatedTerminator reports whether the receiver responds to the private selector _insertTransformation:forEncapsulatedTerminator:.
func (t TTSSpeechString) CanInsertTransformationForEncapsulatedTerminator() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_insertTransformation:forEncapsulatedTerminator:"))
}
func (t TTSSpeechString) _rangeIsValid(valid foundation.NSRange) bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("_rangeIsValid:"), valid)
	return rv
}

// RangeIsValid is an exported wrapper for the private method _rangeIsValid.
func (t TTSSpeechString) RangeIsValid(valid foundation.NSRange) (bool, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_rangeIsValid:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_rangeIsValid:"}
		return false, err
	}
	return t._rangeIsValid(valid), nil
}

// CanRangeIsValid reports whether the receiver responds to the private selector _rangeIsValid:.
func (t TTSSpeechString) CanRangeIsValid() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_rangeIsValid:"))
}
func (t TTSSpeechString) _transformedStringNonMutating() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_transformedStringNonMutating"))
	return objectivec.Object{ID: rv}
}

// TransformedStringNonMutating is an exported wrapper for the private method _transformedStringNonMutating.
func (t TTSSpeechString) TransformedStringNonMutating() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_transformedStringNonMutating")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_transformedStringNonMutating"}
		return nil, err
	}
	return t._transformedStringNonMutating(), nil
}

// CanTransformedStringNonMutating reports whether the receiver responds to the private selector _transformedStringNonMutating.
func (t TTSSpeechString) CanTransformedStringNonMutating() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_transformedStringNonMutating"))
}
func (t TTSSpeechString) _translateRangeInTransformedStringWithParent(string_ foundation.NSRange, parent objectivec.IObject) foundation.NSRange {
	rv := objc.SendIfResponds[foundation.NSRange](t.ID, objc.Sel("_translateRangeInTransformedString:withParent:"), string_, parent)
	return foundation.NSRange(rv)
}

// TranslateRangeInTransformedStringWithParent is an exported wrapper for the private method _translateRangeInTransformedStringWithParent.
func (t TTSSpeechString) TranslateRangeInTransformedStringWithParent(string_ foundation.NSRange, parent objectivec.IObject) (foundation.NSRange, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_translateRangeInTransformedString:withParent:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_translateRangeInTransformedString:withParent:"}
		return foundation.NSRange{}, err
	}
	return t._translateRangeInTransformedStringWithParent(string_, parent), nil
}

// CanTranslateRangeInTransformedStringWithParent reports whether the receiver responds to the private selector _translateRangeInTransformedString:withParent:.
func (t TTSSpeechString) CanTranslateRangeInTransformedStringWithParent() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_translateRangeInTransformedString:withParent:"))
}
func (t TTSSpeechString) EncapsulateSubstringAtRangeWithPrefixAndSuffix(range_ foundation.NSRange, prefix objectivec.IObject, suffix objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("encapsulateSubstringAtRange:withPrefix:andSuffix:"), range_, prefix, suffix)
	return rv
}
func (t TTSSpeechString) InsertAtLocationString(location uint64, string_ objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("insertAtLocation:string:"), location, string_)
	return rv
}
func (t TTSSpeechString) ReplaceOccurencesOfStringWithString(string_ objectivec.IObject, string_2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("replaceOccurencesOfString:withString:"), string_, string_2)
}
func (t TTSSpeechString) TransformRangeTo(range_ foundation.NSRange, to objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("transformRange:to:"), range_, to)
	return rv
}
func (t TTSSpeechString) TranslateRangeInTransformedString(string_ foundation.NSRange) foundation.NSRange {
	rv := objc.SendIfResponds[foundation.NSRange](t.ID, objc.Sel("translateRangeInTransformedString:"), string_)
	return foundation.NSRange(rv)
}
func (t TTSSpeechString) XmlEscaped() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("xmlEscaped"))
	return objectivec.Object{ID: rv}
}
func (t TTSSpeechString) XmlUnescaped() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("xmlUnescaped"))
	return objectivec.Object{ID: rv}
}
func (t TTSSpeechString) InitWithOriginalString(string_ objectivec.IObject) TTSSpeechString {
	rv := objc.SendIfResponds[TTSSpeechString](t.ID, objc.Sel("initWithOriginalString:"), string_)
	return rv
}
func (t TTSSpeechString) InitWithParentSpeechString(string_ objectivec.IObject) TTSSpeechString {
	rv := objc.SendIfResponds[TTSSpeechString](t.ID, objc.Sel("initWithParentSpeechString:"), string_)
	return rv
}
func (t TTSSpeechString) InitWithSSMLString(sSMLString objectivec.IObject) TTSSpeechString {
	rv := objc.SendIfResponds[TTSSpeechString](t.ID, objc.Sel("initWithSSMLString:"), sSMLString)
	return rv
}

func (t TTSSpeechString) DefrostedTransformedString() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("defrostedTransformedString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechString) Finalized() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("finalized"))
	return rv
}
func (t TTSSpeechString) OriginalString() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("originalString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechString) SetOriginalString(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setOriginalString:"), objc.String(value))
}
func (t TTSSpeechString) ParentString() ITTSSpeechString {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("parentString"))
	return TTSSpeechStringFromID(objc.ID(rv))
}
func (t TTSSpeechString) SetParentString(value ITTSSpeechString) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setParentString:"), value)
}
func (t TTSSpeechString) Transformations() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("transformations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechString) SetTransformations(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setTransformations:"), value)
}
func (t TTSSpeechString) TransformedString() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("transformedString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechString) SetTransformedString(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setTransformedString:"), objc.String(value))
}
func (t TTSSpeechString) Type() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("type"))
	return rv
}
func (t TTSSpeechString) SetType(value uint64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setType:"), value)
}
