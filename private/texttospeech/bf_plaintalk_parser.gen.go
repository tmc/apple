// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFPlaintalkParser] class.
var (
	_BFPlaintalkParserClass     BFPlaintalkParserClass
	_BFPlaintalkParserClassOnce sync.Once
)

func getBFPlaintalkParserClass() BFPlaintalkParserClass {
	_BFPlaintalkParserClassOnce.Do(func() {
		_BFPlaintalkParserClass = BFPlaintalkParserClass{class: objc.GetClass("BFPlaintalkParser")}
	})
	return _BFPlaintalkParserClass
}

// GetBFPlaintalkParserClass returns the class object for BFPlaintalkParser.
func GetBFPlaintalkParserClass() BFPlaintalkParserClass {
	return getBFPlaintalkParserClass()
}

type BFPlaintalkParserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFPlaintalkParserClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFPlaintalkParserClass) Alloc() BFPlaintalkParser {
	rv := objc.SendIfResponds[BFPlaintalkParser](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFPlaintalkParser._nextCommandAtStartOfTail]
//   - [BFPlaintalkParser._nextNumericStringRange]
//   - [BFPlaintalkParser._nextTextStringRange]
//   - [BFPlaintalkParser._parse]
//   - [BFPlaintalkParser._parseCommand]
//   - [BFPlaintalkParser._parseText]
//   - [BFPlaintalkParser.ContextSkipStart]
//   - [BFPlaintalkParser.SetContextSkipStart]
//   - [BFPlaintalkParser.ContextSkipperState]
//   - [BFPlaintalkParser.SetContextSkipperState]
//   - [BFPlaintalkParser.CurrentIndex]
//   - [BFPlaintalkParser.SetCurrentIndex]
//   - [BFPlaintalkParser.CurrentProsodicState]
//   - [BFPlaintalkParser.SetCurrentProsodicState]
//   - [BFPlaintalkParser.FirstNonCommandCharacterIndex]
//   - [BFPlaintalkParser.SetFirstNonCommandCharacterIndex]
//   - [BFPlaintalkParser.InitialProsodicState]
//   - [BFPlaintalkParser.SetInitialProsodicState]
//   - [BFPlaintalkParser.LastNonCommandCharacterIndex]
//   - [BFPlaintalkParser.SetLastNonCommandCharacterIndex]
//   - [BFPlaintalkParser.Parse]
//   - [BFPlaintalkParser.PlaintalkString]
//   - [BFPlaintalkParser.SetPlaintalkString]
//   - [BFPlaintalkParser.Product]
//   - [BFPlaintalkParser.SetProduct]
//   - [BFPlaintalkParser.Tail]
//   - [BFPlaintalkParser.InitWithTextAndProsodicState]
type BFPlaintalkParser struct {
	objectivec.Object
}

// BFPlaintalkParserFromID constructs a [BFPlaintalkParser] from an objc.ID.
func BFPlaintalkParserFromID(id objc.ID) BFPlaintalkParser {
	return BFPlaintalkParser{objectivec.Object{ID: id}}
}

// Ensure BFPlaintalkParser implements IBFPlaintalkParser.
var _ IBFPlaintalkParser = BFPlaintalkParser{}

// An interface definition for the [BFPlaintalkParser] class.
//
// # Methods
//
//   - [IBFPlaintalkParser._nextCommandAtStartOfTail]
//   - [IBFPlaintalkParser._nextNumericStringRange]
//   - [IBFPlaintalkParser._nextTextStringRange]
//   - [IBFPlaintalkParser._parse]
//   - [IBFPlaintalkParser._parseCommand]
//   - [IBFPlaintalkParser._parseText]
//   - [IBFPlaintalkParser.ContextSkipStart]
//   - [IBFPlaintalkParser.SetContextSkipStart]
//   - [IBFPlaintalkParser.ContextSkipperState]
//   - [IBFPlaintalkParser.SetContextSkipperState]
//   - [IBFPlaintalkParser.CurrentIndex]
//   - [IBFPlaintalkParser.SetCurrentIndex]
//   - [IBFPlaintalkParser.CurrentProsodicState]
//   - [IBFPlaintalkParser.SetCurrentProsodicState]
//   - [IBFPlaintalkParser.FirstNonCommandCharacterIndex]
//   - [IBFPlaintalkParser.SetFirstNonCommandCharacterIndex]
//   - [IBFPlaintalkParser.InitialProsodicState]
//   - [IBFPlaintalkParser.SetInitialProsodicState]
//   - [IBFPlaintalkParser.LastNonCommandCharacterIndex]
//   - [IBFPlaintalkParser.SetLastNonCommandCharacterIndex]
//   - [IBFPlaintalkParser.Parse]
//   - [IBFPlaintalkParser.PlaintalkString]
//   - [IBFPlaintalkParser.SetPlaintalkString]
//   - [IBFPlaintalkParser.Product]
//   - [IBFPlaintalkParser.SetProduct]
//   - [IBFPlaintalkParser.Tail]
//   - [IBFPlaintalkParser.InitWithTextAndProsodicState]
type IBFPlaintalkParser interface {
	objectivec.IObject

	// Topic: Methods

	_nextCommandAtStartOfTail() foundation.NSRange
	_nextNumericStringRange() foundation.NSRange
	_nextTextStringRange() foundation.NSRange
	_parse()
	_parseCommand()
	_parseText()
	ContextSkipStart() uint64
	SetContextSkipStart(value uint64)
	ContextSkipperState() uint64
	SetContextSkipperState(value uint64)
	CurrentIndex() int64
	SetCurrentIndex(value int64)
	CurrentProsodicState() IBFProsodicState
	SetCurrentProsodicState(value IBFProsodicState)
	FirstNonCommandCharacterIndex() int64
	SetFirstNonCommandCharacterIndex(value int64)
	InitialProsodicState() IBFProsodicState
	SetInitialProsodicState(value IBFProsodicState)
	LastNonCommandCharacterIndex() int64
	SetLastNonCommandCharacterIndex(value int64)
	Parse() objectivec.IObject
	PlaintalkString() string
	SetPlaintalkString(value string)
	Product() foundation.INSArray
	SetProduct(value foundation.INSArray)
	Tail() objectivec.IObject
	InitWithTextAndProsodicState(text objectivec.IObject, state objectivec.IObject) BFPlaintalkParser
}

// Init initializes the instance.
func (b BFPlaintalkParser) Init() BFPlaintalkParser {
	rv := objc.SendIfResponds[BFPlaintalkParser](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFPlaintalkParser) Autorelease() BFPlaintalkParser {
	rv := objc.SendIfResponds[BFPlaintalkParser](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFPlaintalkParser creates a new BFPlaintalkParser instance.
func NewBFPlaintalkParser() BFPlaintalkParser {
	class := getBFPlaintalkParserClass()
	rv := objc.SendIfResponds[BFPlaintalkParser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFPlaintalkParserWithTextAndProsodicState(text objectivec.IObject, state objectivec.IObject) BFPlaintalkParser {
	instance := getBFPlaintalkParserClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithText:andProsodicState:"), text, state)
	return BFPlaintalkParserFromID(rv)
}

func (b BFPlaintalkParser) _nextCommandAtStartOfTail() foundation.NSRange {
	rv := objc.SendIfResponds[foundation.NSRange](b.ID, objc.Sel("_nextCommandAtStartOfTail"))
	return foundation.NSRange(rv)
}

// NextCommandAtStartOfTail is an exported wrapper for the private method _nextCommandAtStartOfTail.
func (b BFPlaintalkParser) NextCommandAtStartOfTail() (foundation.NSRange, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_nextCommandAtStartOfTail")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_nextCommandAtStartOfTail"}
		return foundation.NSRange{}, err
	}
	return b._nextCommandAtStartOfTail(), nil
}

// CanNextCommandAtStartOfTail reports whether the receiver responds to the private selector _nextCommandAtStartOfTail.
func (b BFPlaintalkParser) CanNextCommandAtStartOfTail() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_nextCommandAtStartOfTail"))
}
func (b BFPlaintalkParser) _nextNumericStringRange() foundation.NSRange {
	rv := objc.SendIfResponds[foundation.NSRange](b.ID, objc.Sel("_nextNumericStringRange"))
	return foundation.NSRange(rv)
}

// NextNumericStringRange is an exported wrapper for the private method _nextNumericStringRange.
func (b BFPlaintalkParser) NextNumericStringRange() (foundation.NSRange, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_nextNumericStringRange")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_nextNumericStringRange"}
		return foundation.NSRange{}, err
	}
	return b._nextNumericStringRange(), nil
}

// CanNextNumericStringRange reports whether the receiver responds to the private selector _nextNumericStringRange.
func (b BFPlaintalkParser) CanNextNumericStringRange() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_nextNumericStringRange"))
}
func (b BFPlaintalkParser) _nextTextStringRange() foundation.NSRange {
	rv := objc.SendIfResponds[foundation.NSRange](b.ID, objc.Sel("_nextTextStringRange"))
	return foundation.NSRange(rv)
}

// NextTextStringRange is an exported wrapper for the private method _nextTextStringRange.
func (b BFPlaintalkParser) NextTextStringRange() (foundation.NSRange, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_nextTextStringRange")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_nextTextStringRange"}
		return foundation.NSRange{}, err
	}
	return b._nextTextStringRange(), nil
}

// CanNextTextStringRange reports whether the receiver responds to the private selector _nextTextStringRange.
func (b BFPlaintalkParser) CanNextTextStringRange() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_nextTextStringRange"))
}
func (b BFPlaintalkParser) _parse() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_parse"))
}
func (b BFPlaintalkParser) _parseCommand() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_parseCommand"))
}

// ParseCommand is an exported wrapper for the private method _parseCommand.
func (b BFPlaintalkParser) ParseCommand() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_parseCommand")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_parseCommand"}
		return err
	}
	b._parseCommand()
	return nil
}

// CanParseCommand reports whether the receiver responds to the private selector _parseCommand.
func (b BFPlaintalkParser) CanParseCommand() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_parseCommand"))
}
func (b BFPlaintalkParser) _parseText() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_parseText"))
}

// ParseText is an exported wrapper for the private method _parseText.
func (b BFPlaintalkParser) ParseText() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_parseText")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_parseText"}
		return err
	}
	b._parseText()
	return nil
}

// CanParseText reports whether the receiver responds to the private selector _parseText.
func (b BFPlaintalkParser) CanParseText() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_parseText"))
}
func (b BFPlaintalkParser) Parse() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("parse"))
	return objectivec.Object{ID: rv}
}
func (b BFPlaintalkParser) Tail() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("tail"))
	return objectivec.Object{ID: rv}
}
func (b BFPlaintalkParser) InitWithTextAndProsodicState(text objectivec.IObject, state objectivec.IObject) BFPlaintalkParser {
	rv := objc.SendIfResponds[BFPlaintalkParser](b.ID, objc.Sel("initWithText:andProsodicState:"), text, state)
	return rv
}

func (b BFPlaintalkParser) ContextSkipStart() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("contextSkipStart"))
	return rv
}
func (b BFPlaintalkParser) SetContextSkipStart(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setContextSkipStart:"), value)
}
func (b BFPlaintalkParser) ContextSkipperState() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("contextSkipperState"))
	return rv
}
func (b BFPlaintalkParser) SetContextSkipperState(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setContextSkipperState:"), value)
}
func (b BFPlaintalkParser) CurrentIndex() int64 {
	rv := objc.SendIfResponds[int64](b.ID, objc.Sel("currentIndex"))
	return rv
}
func (b BFPlaintalkParser) SetCurrentIndex(value int64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCurrentIndex:"), value)
}
func (b BFPlaintalkParser) CurrentProsodicState() IBFProsodicState {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("currentProsodicState"))
	return BFProsodicStateFromID(objc.ID(rv))
}
func (b BFPlaintalkParser) SetCurrentProsodicState(value IBFProsodicState) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCurrentProsodicState:"), value)
}
func (b BFPlaintalkParser) FirstNonCommandCharacterIndex() int64 {
	rv := objc.SendIfResponds[int64](b.ID, objc.Sel("firstNonCommandCharacterIndex"))
	return rv
}
func (b BFPlaintalkParser) SetFirstNonCommandCharacterIndex(value int64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setFirstNonCommandCharacterIndex:"), value)
}
func (b BFPlaintalkParser) InitialProsodicState() IBFProsodicState {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("initialProsodicState"))
	return BFProsodicStateFromID(objc.ID(rv))
}
func (b BFPlaintalkParser) SetInitialProsodicState(value IBFProsodicState) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setInitialProsodicState:"), value)
}
func (b BFPlaintalkParser) LastNonCommandCharacterIndex() int64 {
	rv := objc.SendIfResponds[int64](b.ID, objc.Sel("lastNonCommandCharacterIndex"))
	return rv
}
func (b BFPlaintalkParser) SetLastNonCommandCharacterIndex(value int64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setLastNonCommandCharacterIndex:"), value)
}
func (b BFPlaintalkParser) PlaintalkString() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("plaintalkString"))
	return foundation.NSStringFromID(rv).String()
}
func (b BFPlaintalkParser) SetPlaintalkString(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPlaintalkString:"), objc.String(value))
}
func (b BFPlaintalkParser) Product() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("product"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (b BFPlaintalkParser) SetProduct(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setProduct:"), value)
}
