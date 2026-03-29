// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOUtteranceResult] class.
var (
	_SOUtteranceResultClass     SOUtteranceResultClass
	_SOUtteranceResultClassOnce sync.Once
)

func getSOUtteranceResultClass() SOUtteranceResultClass {
	_SOUtteranceResultClassOnce.Do(func() {
		_SOUtteranceResultClass = SOUtteranceResultClass{class: objc.GetClass("SOUtteranceResult")}
	})
	return _SOUtteranceResultClass
}

// GetSOUtteranceResultClass returns the class object for SOUtteranceResult.
func GetSOUtteranceResultClass() SOUtteranceResultClass {
	return getSOUtteranceResultClass()
}

type SOUtteranceResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOUtteranceResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOUtteranceResultClass) Alloc() SOUtteranceResult {
	rv := objc.Send[SOUtteranceResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SOUtteranceResult._dictionary]
//   - [SOUtteranceResult._initWithDictionary]
//   - [SOUtteranceResult._normalizeTimesAgainstTimeInterval]
//   - [SOUtteranceResult.AudioFilePath]
//   - [SOUtteranceResult.SetAudioFilePath]
//   - [SOUtteranceResult.CommandIdentifier]
//   - [SOUtteranceResult.SetCommandIdentifier]
//   - [SOUtteranceResult.CreationDate]
//   - [SOUtteranceResult.SetCreationDate]
//   - [SOUtteranceResult.EndTime]
//   - [SOUtteranceResult.SetEndTime]
//   - [SOUtteranceResult.PrependedSilenceDuration]
//   - [SOUtteranceResult.SetPrependedSilenceDuration]
//   - [SOUtteranceResult.StartTime]
//   - [SOUtteranceResult.SetStartTime]
//   - [SOUtteranceResult.Text]
//   - [SOUtteranceResult.SetText]
//   - [SOUtteranceResult.TextVariants]
//   - [SOUtteranceResult.SetTextVariants]
//   - [SOUtteranceResult.Type]
//   - [SOUtteranceResult.SetType]
//   - [SOUtteranceResult.InitWithTypeStartTimeEndTimeTextTextVariantsCommandIdentifier]
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult
type SOUtteranceResult struct {
	objectivec.Object
}

// SOUtteranceResultFromID constructs a [SOUtteranceResult] from an objc.ID.
func SOUtteranceResultFromID(id objc.ID) SOUtteranceResult {
	return SOUtteranceResult{objectivec.Object{ID: id}}
}
// Ensure SOUtteranceResult implements ISOUtteranceResult.
var _ ISOUtteranceResult = SOUtteranceResult{}

// An interface definition for the [SOUtteranceResult] class.
//
// # Methods
//
//   - [ISOUtteranceResult._dictionary]
//   - [ISOUtteranceResult._initWithDictionary]
//   - [ISOUtteranceResult._normalizeTimesAgainstTimeInterval]
//   - [ISOUtteranceResult.AudioFilePath]
//   - [ISOUtteranceResult.SetAudioFilePath]
//   - [ISOUtteranceResult.CommandIdentifier]
//   - [ISOUtteranceResult.SetCommandIdentifier]
//   - [ISOUtteranceResult.CreationDate]
//   - [ISOUtteranceResult.SetCreationDate]
//   - [ISOUtteranceResult.EndTime]
//   - [ISOUtteranceResult.SetEndTime]
//   - [ISOUtteranceResult.PrependedSilenceDuration]
//   - [ISOUtteranceResult.SetPrependedSilenceDuration]
//   - [ISOUtteranceResult.StartTime]
//   - [ISOUtteranceResult.SetStartTime]
//   - [ISOUtteranceResult.Text]
//   - [ISOUtteranceResult.SetText]
//   - [ISOUtteranceResult.TextVariants]
//   - [ISOUtteranceResult.SetTextVariants]
//   - [ISOUtteranceResult.Type]
//   - [ISOUtteranceResult.SetType]
//   - [ISOUtteranceResult.InitWithTypeStartTimeEndTimeTextTextVariantsCommandIdentifier]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult
type ISOUtteranceResult interface {
	objectivec.IObject

	// Topic: Methods

	_dictionary() objectivec.IObject
	_initWithDictionary(dictionary objectivec.IObject) objectivec.IObject
	_normalizeTimesAgainstTimeInterval(interval float64)
	AudioFilePath() string
	SetAudioFilePath(value string)
	CommandIdentifier() string
	SetCommandIdentifier(value string)
	CreationDate() float64
	SetCreationDate(value float64)
	EndTime() float64
	SetEndTime(value float64)
	PrependedSilenceDuration() float64
	SetPrependedSilenceDuration(value float64)
	StartTime() float64
	SetStartTime(value float64)
	Text() string
	SetText(value string)
	TextVariants() foundation.INSArray
	SetTextVariants(value foundation.INSArray)
	Type() string
	SetType(value string)
	InitWithTypeStartTimeEndTimeTextTextVariantsCommandIdentifier(type_ objectivec.IObject, time float64, time2 float64, text objectivec.IObject, variants objectivec.IObject, identifier objectivec.IObject) SOUtteranceResult
}

// Init initializes the instance.
func (s SOUtteranceResult) Init() SOUtteranceResult {
	rv := objc.Send[SOUtteranceResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOUtteranceResult) Autorelease() SOUtteranceResult {
	rv := objc.Send[SOUtteranceResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOUtteranceResult creates a new SOUtteranceResult instance.
func NewSOUtteranceResult() SOUtteranceResult {
	class := getSOUtteranceResultClass()
	rv := objc.Send[SOUtteranceResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/initWithType:startTime:endTime:text:textVariants:commandIdentifier:
func NewSOUtteranceResultWithTypeStartTimeEndTimeTextTextVariantsCommandIdentifier(type_ objectivec.IObject, time float64, time2 float64, text objectivec.IObject, variants objectivec.IObject, identifier objectivec.IObject) SOUtteranceResult {
	instance := getSOUtteranceResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:startTime:endTime:text:textVariants:commandIdentifier:"), type_, time, time2, text, variants, identifier)
	return SOUtteranceResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/_dictionary
func (s SOUtteranceResult) _dictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_dictionary"))
	return objectivec.Object{ID: rv}
}

// Dictionary is an exported wrapper for the private method _dictionary.
func (s SOUtteranceResult) Dictionary() objectivec.IObject {
	return s._dictionary()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/_initWithDictionary:
func (s SOUtteranceResult) _initWithDictionary(dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_initWithDictionary:"), dictionary)
	return objectivec.Object{ID: rv}
}

// InitWithDictionary is an exported wrapper for the private method _initWithDictionary.
func (s SOUtteranceResult) InitWithDictionary(dictionary objectivec.IObject) objectivec.IObject {
	return s._initWithDictionary(dictionary)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/_normalizeTimesAgainstTimeInterval:
func (s SOUtteranceResult) _normalizeTimesAgainstTimeInterval(interval float64) {
	objc.Send[objc.ID](s.ID, objc.Sel("_normalizeTimesAgainstTimeInterval:"), interval)
}

// NormalizeTimesAgainstTimeInterval is an exported wrapper for the private method _normalizeTimesAgainstTimeInterval.
func (s SOUtteranceResult) NormalizeTimesAgainstTimeInterval(interval float64) {
	s._normalizeTimesAgainstTimeInterval(interval)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/initWithType:startTime:endTime:text:textVariants:commandIdentifier:
func (s SOUtteranceResult) InitWithTypeStartTimeEndTimeTextTextVariantsCommandIdentifier(type_ objectivec.IObject, time float64, time2 float64, text objectivec.IObject, variants objectivec.IObject, identifier objectivec.IObject) SOUtteranceResult {
	rv := objc.Send[SOUtteranceResult](s.ID, objc.Sel("initWithType:startTime:endTime:text:textVariants:commandIdentifier:"), type_, time, time2, text, variants, identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/audioFilePath
func (s SOUtteranceResult) AudioFilePath() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("audioFilePath"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResult) SetAudioFilePath(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setAudioFilePath:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/commandIdentifier
func (s SOUtteranceResult) CommandIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commandIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResult) SetCommandIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setCommandIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/creationDate
func (s SOUtteranceResult) CreationDate() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("creationDate"))
	return rv
}
func (s SOUtteranceResult) SetCreationDate(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setCreationDate:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/endTime
func (s SOUtteranceResult) EndTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("endTime"))
	return rv
}
func (s SOUtteranceResult) SetEndTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setEndTime:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/prependedSilenceDuration
func (s SOUtteranceResult) PrependedSilenceDuration() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("prependedSilenceDuration"))
	return rv
}
func (s SOUtteranceResult) SetPrependedSilenceDuration(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPrependedSilenceDuration:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/startTime
func (s SOUtteranceResult) StartTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("startTime"))
	return rv
}
func (s SOUtteranceResult) SetStartTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setStartTime:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/text
func (s SOUtteranceResult) Text() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResult) SetText(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setText:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/textVariants
func (s SOUtteranceResult) TextVariants() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("textVariants"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SOUtteranceResult) SetTextVariants(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setTextVariants:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResult/type
func (s SOUtteranceResult) Type() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResult) SetType(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setType:"), objc.String(value))
}

