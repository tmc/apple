// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOUtteranceResultsFile] class.
var (
	_SOUtteranceResultsFileClass     SOUtteranceResultsFileClass
	_SOUtteranceResultsFileClassOnce sync.Once
)

func getSOUtteranceResultsFileClass() SOUtteranceResultsFileClass {
	_SOUtteranceResultsFileClassOnce.Do(func() {
		_SOUtteranceResultsFileClass = SOUtteranceResultsFileClass{class: objc.GetClass("SOUtteranceResultsFile")}
	})
	return _SOUtteranceResultsFileClass
}

// GetSOUtteranceResultsFileClass returns the class object for SOUtteranceResultsFile.
func GetSOUtteranceResultsFileClass() SOUtteranceResultsFileClass {
	return getSOUtteranceResultsFileClass()
}

type SOUtteranceResultsFileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOUtteranceResultsFileClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOUtteranceResultsFileClass) Alloc() SOUtteranceResultsFile {
	rv := objc.Send[SOUtteranceResultsFile](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SOUtteranceResultsFile._dictionary]
//   - [SOUtteranceResultsFile.AppendUtteranceResult]
//   - [SOUtteranceResultsFile.AudioFileData]
//   - [SOUtteranceResultsFile.SetAudioFileData]
//   - [SOUtteranceResultsFile.CreationDate]
//   - [SOUtteranceResultsFile.SetCreationDate]
//   - [SOUtteranceResultsFile.CreatorArguments]
//   - [SOUtteranceResultsFile.SetCreatorArguments]
//   - [SOUtteranceResultsFile.CreatorName]
//   - [SOUtteranceResultsFile.SetCreatorName]
//   - [SOUtteranceResultsFile.ExportAudioDataToFile]
//   - [SOUtteranceResultsFile.ExportRXGrammarToFile]
//   - [SOUtteranceResultsFile.FirstUtteranceStartTime]
//   - [SOUtteranceResultsFile.SetFirstUtteranceStartTime]
//   - [SOUtteranceResultsFile.ImportAudioDataFromFile]
//   - [SOUtteranceResultsFile.ImportRXGrammarFromFile]
//   - [SOUtteranceResultsFile.LocaleIdentifier]
//   - [SOUtteranceResultsFile.SetLocaleIdentifier]
//   - [SOUtteranceResultsFile.RecognizerType]
//   - [SOUtteranceResultsFile.SetRecognizerType]
//   - [SOUtteranceResultsFile.RxGrammar]
//   - [SOUtteranceResultsFile.SetRxGrammar]
//   - [SOUtteranceResultsFile.StreamDescription]
//   - [SOUtteranceResultsFile.SetStreamDescription]
//   - [SOUtteranceResultsFile.UtteranceResults]
//   - [SOUtteranceResultsFile.SetUtteranceResults]
//   - [SOUtteranceResultsFile.WriteToFile]
//   - [SOUtteranceResultsFile.InitWithContentsOfFile]
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile
type SOUtteranceResultsFile struct {
	objectivec.Object
}

// SOUtteranceResultsFileFromID constructs a [SOUtteranceResultsFile] from an objc.ID.
func SOUtteranceResultsFileFromID(id objc.ID) SOUtteranceResultsFile {
	return SOUtteranceResultsFile{objectivec.Object{ID: id}}
}
// Ensure SOUtteranceResultsFile implements ISOUtteranceResultsFile.
var _ ISOUtteranceResultsFile = SOUtteranceResultsFile{}

// An interface definition for the [SOUtteranceResultsFile] class.
//
// # Methods
//
//   - [ISOUtteranceResultsFile._dictionary]
//   - [ISOUtteranceResultsFile.AppendUtteranceResult]
//   - [ISOUtteranceResultsFile.AudioFileData]
//   - [ISOUtteranceResultsFile.SetAudioFileData]
//   - [ISOUtteranceResultsFile.CreationDate]
//   - [ISOUtteranceResultsFile.SetCreationDate]
//   - [ISOUtteranceResultsFile.CreatorArguments]
//   - [ISOUtteranceResultsFile.SetCreatorArguments]
//   - [ISOUtteranceResultsFile.CreatorName]
//   - [ISOUtteranceResultsFile.SetCreatorName]
//   - [ISOUtteranceResultsFile.ExportAudioDataToFile]
//   - [ISOUtteranceResultsFile.ExportRXGrammarToFile]
//   - [ISOUtteranceResultsFile.FirstUtteranceStartTime]
//   - [ISOUtteranceResultsFile.SetFirstUtteranceStartTime]
//   - [ISOUtteranceResultsFile.ImportAudioDataFromFile]
//   - [ISOUtteranceResultsFile.ImportRXGrammarFromFile]
//   - [ISOUtteranceResultsFile.LocaleIdentifier]
//   - [ISOUtteranceResultsFile.SetLocaleIdentifier]
//   - [ISOUtteranceResultsFile.RecognizerType]
//   - [ISOUtteranceResultsFile.SetRecognizerType]
//   - [ISOUtteranceResultsFile.RxGrammar]
//   - [ISOUtteranceResultsFile.SetRxGrammar]
//   - [ISOUtteranceResultsFile.StreamDescription]
//   - [ISOUtteranceResultsFile.SetStreamDescription]
//   - [ISOUtteranceResultsFile.UtteranceResults]
//   - [ISOUtteranceResultsFile.SetUtteranceResults]
//   - [ISOUtteranceResultsFile.WriteToFile]
//   - [ISOUtteranceResultsFile.InitWithContentsOfFile]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile
type ISOUtteranceResultsFile interface {
	objectivec.IObject

	// Topic: Methods

	_dictionary() objectivec.IObject
	AppendUtteranceResult(result objectivec.IObject)
	AudioFileData() foundation.INSData
	SetAudioFileData(value foundation.INSData)
	CreationDate() float64
	SetCreationDate(value float64)
	CreatorArguments() string
	SetCreatorArguments(value string)
	CreatorName() string
	SetCreatorName(value string)
	ExportAudioDataToFile(file objectivec.IObject) bool
	ExportRXGrammarToFile(file objectivec.IObject) bool
	FirstUtteranceStartTime() float64
	SetFirstUtteranceStartTime(value float64)
	ImportAudioDataFromFile(file objectivec.IObject) bool
	ImportRXGrammarFromFile(file objectivec.IObject) bool
	LocaleIdentifier() string
	SetLocaleIdentifier(value string)
	RecognizerType() string
	SetRecognizerType(value string)
	RxGrammar() string
	SetRxGrammar(value string)
	StreamDescription() objectivec.IObject
	SetStreamDescription(value objectivec.IObject)
	UtteranceResults() foundation.INSArray
	SetUtteranceResults(value foundation.INSArray)
	WriteToFile(file objectivec.IObject) bool
	InitWithContentsOfFile(file objectivec.IObject) SOUtteranceResultsFile
}

// Init initializes the instance.
func (s SOUtteranceResultsFile) Init() SOUtteranceResultsFile {
	rv := objc.Send[SOUtteranceResultsFile](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOUtteranceResultsFile) Autorelease() SOUtteranceResultsFile {
	rv := objc.Send[SOUtteranceResultsFile](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOUtteranceResultsFile creates a new SOUtteranceResultsFile instance.
func NewSOUtteranceResultsFile() SOUtteranceResultsFile {
	class := getSOUtteranceResultsFileClass()
	rv := objc.Send[SOUtteranceResultsFile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/initWithContentsOfFile:
func NewSOUtteranceResultsFileWithContentsOfFile(file objectivec.IObject) SOUtteranceResultsFile {
	instance := getSOUtteranceResultsFileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), file)
	return SOUtteranceResultsFileFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/_dictionary
func (s SOUtteranceResultsFile) _dictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_dictionary"))
	return objectivec.Object{ID: rv}
}

// Dictionary is an exported wrapper for the private method _dictionary.
func (s SOUtteranceResultsFile) Dictionary() objectivec.IObject {
	return s._dictionary()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/appendUtteranceResult:
func (s SOUtteranceResultsFile) AppendUtteranceResult(result objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("appendUtteranceResult:"), result)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/exportAudioDataToFile:
func (s SOUtteranceResultsFile) ExportAudioDataToFile(file objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("exportAudioDataToFile:"), file)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/exportRXGrammarToFile:
func (s SOUtteranceResultsFile) ExportRXGrammarToFile(file objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("exportRXGrammarToFile:"), file)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/importAudioDataFromFile:
func (s SOUtteranceResultsFile) ImportAudioDataFromFile(file objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("importAudioDataFromFile:"), file)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/importRXGrammarFromFile:
func (s SOUtteranceResultsFile) ImportRXGrammarFromFile(file objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("importRXGrammarFromFile:"), file)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/writeToFile:
func (s SOUtteranceResultsFile) WriteToFile(file objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("writeToFile:"), file)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/initWithContentsOfFile:
func (s SOUtteranceResultsFile) InitWithContentsOfFile(file objectivec.IObject) SOUtteranceResultsFile {
	rv := objc.Send[SOUtteranceResultsFile](s.ID, objc.Sel("initWithContentsOfFile:"), file)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/audioFileData
func (s SOUtteranceResultsFile) AudioFileData() foundation.INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("audioFileData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (s SOUtteranceResultsFile) SetAudioFileData(value foundation.INSData) {
	objc.Send[struct{}](s.ID, objc.Sel("setAudioFileData:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/creationDate
func (s SOUtteranceResultsFile) CreationDate() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("creationDate"))
	return rv
}
func (s SOUtteranceResultsFile) SetCreationDate(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setCreationDate:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/creatorArguments
func (s SOUtteranceResultsFile) CreatorArguments() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("creatorArguments"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResultsFile) SetCreatorArguments(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setCreatorArguments:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/creatorName
func (s SOUtteranceResultsFile) CreatorName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("creatorName"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResultsFile) SetCreatorName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setCreatorName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/firstUtteranceStartTime
func (s SOUtteranceResultsFile) FirstUtteranceStartTime() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("firstUtteranceStartTime"))
	return rv
}
func (s SOUtteranceResultsFile) SetFirstUtteranceStartTime(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setFirstUtteranceStartTime:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/localeIdentifier
func (s SOUtteranceResultsFile) LocaleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localeIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResultsFile) SetLocaleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLocaleIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/recognizerType
func (s SOUtteranceResultsFile) RecognizerType() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("recognizerType"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResultsFile) SetRecognizerType(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setRecognizerType:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/rxGrammar
func (s SOUtteranceResultsFile) RxGrammar() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("rxGrammar"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOUtteranceResultsFile) SetRxGrammar(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setRxGrammar:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/streamDescription
func (s SOUtteranceResultsFile) StreamDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("streamDescription"))
	return objectivec.Object{ID: rv}
}
func (s SOUtteranceResultsFile) SetStreamDescription(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setStreamDescription:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOUtteranceResultsFile/utteranceResults
func (s SOUtteranceResultsFile) UtteranceResults() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("utteranceResults"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SOUtteranceResultsFile) SetUtteranceResults(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setUtteranceResults:"), value)
}

