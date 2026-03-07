// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextCheckingResult] class.
var (
	_NSTextCheckingResultClass     NSTextCheckingResultClass
	_NSTextCheckingResultClassOnce sync.Once
)

func getNSTextCheckingResultClass() NSTextCheckingResultClass {
	_NSTextCheckingResultClassOnce.Do(func() {
		_NSTextCheckingResultClass = NSTextCheckingResultClass{class: objc.GetClass("NSTextCheckingResult")}
	})
	return _NSTextCheckingResultClass
}

// GetNSTextCheckingResultClass returns the class object for NSTextCheckingResult.
func GetNSTextCheckingResultClass() NSTextCheckingResultClass {
	return getNSTextCheckingResultClass()
}

type NSTextCheckingResultClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextCheckingResultClass) Alloc() NSTextCheckingResult {
	rv := objc.Send[NSTextCheckingResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An occurrence of textual content found during the analysis of a block of
// text, such as when matching a regular expression.
//
// # Overview
// 
// On both iOS and macOS, instances of [NSTextCheckingResult] are returned by
// the [NSRegularExpression] class and the [NSDataDetector] class to indicate
// the discovery of content. In those cases, what is found may be a match for
// a regular expression or a date, address, phone number, and so on. In macOS,
// instances of [NSTextCheckingResult] are returned by the [NSSpellChecker]
// object to describe the results of spelling, grammar, or text-substitution
// actions.
//
// [NSSpellChecker]: https://developer.apple.com/documentation/AppKit/NSSpellChecker
//
// # Text Checking Type Range and Type
//
//   - [NSTextCheckingResult.Range]: Returns the range of the result that the receiver represents.
//   - [NSTextCheckingResult.ResultType]: Returns the text checking result type that the receiver represents.
//   - [NSTextCheckingResult.NumberOfRanges]: Returns the number of ranges.
//   - [NSTextCheckingResult.RangeAtIndex]: Returns the result type that the range represents.
//
// # Text Checking Results for Text Replacement
//
//   - [NSTextCheckingResult.ReplacementString]: A replacement string from one of a number of replacement checking results.
//
// # Text Checking Results for Regular Expressions
//
//   - [NSTextCheckingResult.RegularExpression]: The regular expression of a type checking result.
//
// # Text Checking Result Components
//
//   - [NSTextCheckingResult.Components]: A dictionary containing the components of a type checking result.
//
// # Text Checking Results for URLs
//
//   - [NSTextCheckingResult.URL]: The URL of a type checking result.
//
// # Text Checking Results for Addresses
//
//   - [NSTextCheckingResult.AddressComponents]: The address dictionary of a type checking result.
//
// # Text Checking Results for Phone Numbers
//
//   - [NSTextCheckingResult.PhoneNumber]: The phone number of a type checking result.
//
// # Text Checking Results for Dates and Times
//
//   - [NSTextCheckingResult.Date]: The date component of a type checking result.
//   - [NSTextCheckingResult.Duration]: The duration component of a type checking result.
//   - [NSTextCheckingResult.TimeZone]: The time zone component of a type checking result.
//
// # Text Checking Results for Orthography
//
//   - [NSTextCheckingResult.Orthography]: The detected orthography of a type checking result.
//
// # Text Checking Results for Grammar
//
//   - [NSTextCheckingResult.GrammarDetails]: The details of a located grammatical type checking result.
//
// # Adjusting the Ranges of a Text Checking Result
//
//   - [NSTextCheckingResult.ResultByAdjustingRangesWithOffset]: Returns a new text checking result after adjusting the ranges as specified by the offset.
//
// # Instance Properties
//
//   - [NSTextCheckingResult.AlternativeStrings]
//
// # Instance Methods
//
//   - [NSTextCheckingResult.RangeWithName]
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
type NSTextCheckingResult struct {
	objectivec.Object
}

// NSTextCheckingResultFromID constructs a [NSTextCheckingResult] from an objc.ID.
//
// An occurrence of textual content found during the analysis of a block of
// text, such as when matching a regular expression.
func NSTextCheckingResultFromID(id objc.ID) NSTextCheckingResult {
	return NSTextCheckingResult{objectivec.Object{id}}
}
// NOTE: NSTextCheckingResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextCheckingResult] class.
//
// # Text Checking Type Range and Type
//
//   - [INSTextCheckingResult.Range]: Returns the range of the result that the receiver represents.
//   - [INSTextCheckingResult.ResultType]: Returns the text checking result type that the receiver represents.
//   - [INSTextCheckingResult.NumberOfRanges]: Returns the number of ranges.
//   - [INSTextCheckingResult.RangeAtIndex]: Returns the result type that the range represents.
//
// # Text Checking Results for Text Replacement
//
//   - [INSTextCheckingResult.ReplacementString]: A replacement string from one of a number of replacement checking results.
//
// # Text Checking Results for Regular Expressions
//
//   - [INSTextCheckingResult.RegularExpression]: The regular expression of a type checking result.
//
// # Text Checking Result Components
//
//   - [INSTextCheckingResult.Components]: A dictionary containing the components of a type checking result.
//
// # Text Checking Results for URLs
//
//   - [INSTextCheckingResult.URL]: The URL of a type checking result.
//
// # Text Checking Results for Addresses
//
//   - [INSTextCheckingResult.AddressComponents]: The address dictionary of a type checking result.
//
// # Text Checking Results for Phone Numbers
//
//   - [INSTextCheckingResult.PhoneNumber]: The phone number of a type checking result.
//
// # Text Checking Results for Dates and Times
//
//   - [INSTextCheckingResult.Date]: The date component of a type checking result.
//   - [INSTextCheckingResult.Duration]: The duration component of a type checking result.
//   - [INSTextCheckingResult.TimeZone]: The time zone component of a type checking result.
//
// # Text Checking Results for Orthography
//
//   - [INSTextCheckingResult.Orthography]: The detected orthography of a type checking result.
//
// # Text Checking Results for Grammar
//
//   - [INSTextCheckingResult.GrammarDetails]: The details of a located grammatical type checking result.
//
// # Adjusting the Ranges of a Text Checking Result
//
//   - [INSTextCheckingResult.ResultByAdjustingRangesWithOffset]: Returns a new text checking result after adjusting the ranges as specified by the offset.
//
// # Instance Properties
//
//   - [INSTextCheckingResult.AlternativeStrings]
//
// # Instance Methods
//
//   - [INSTextCheckingResult.RangeWithName]
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
type INSTextCheckingResult interface {
	objectivec.IObject
	NSCopying

	// Topic: Text Checking Type Range and Type

	// Returns the range of the result that the receiver represents.
	Range() NSRange
	// Returns the text checking result type that the receiver represents.
	ResultType() NSTextCheckingType
	// Returns the number of ranges.
	NumberOfRanges() uint
	// Returns the result type that the range represents.
	RangeAtIndex(idx uint) NSRange

	// Topic: Text Checking Results for Text Replacement

	// A replacement string from one of a number of replacement checking results.
	ReplacementString() string

	// Topic: Text Checking Results for Regular Expressions

	// The regular expression of a type checking result.
	RegularExpression() INSRegularExpression

	// Topic: Text Checking Result Components

	// A dictionary containing the components of a type checking result.
	Components() INSDictionary

	// Topic: Text Checking Results for URLs

	// The URL of a type checking result.
	URL() INSURL

	// Topic: Text Checking Results for Addresses

	// The address dictionary of a type checking result.
	AddressComponents() INSDictionary

	// Topic: Text Checking Results for Phone Numbers

	// The phone number of a type checking result.
	PhoneNumber() string

	// Topic: Text Checking Results for Dates and Times

	// The date component of a type checking result.
	Date() INSDate
	// The duration component of a type checking result.
	Duration() float64
	// The time zone component of a type checking result.
	TimeZone() INSTimeZone

	// Topic: Text Checking Results for Orthography

	// The detected orthography of a type checking result.
	Orthography() INSOrthography

	// Topic: Text Checking Results for Grammar

	// The details of a located grammatical type checking result.
	GrammarDetails() INSDictionary

	// Topic: Adjusting the Ranges of a Text Checking Result

	// Returns a new text checking result after adjusting the ranges as specified by the offset.
	ResultByAdjustingRangesWithOffset(offset int) INSTextCheckingResult

	// Topic: Instance Properties

	AlternativeStrings() []string

	// Topic: Instance Methods

	RangeWithName(name string) NSRange

	// A value indicating that a requested item couldn’t be found or doesn’t exist.
	NSNotFound() int
	SetNSNotFound(value int)
	InitWithCoder(coder INSCoder) NSTextCheckingResult
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}





// Init initializes the instance.
func (t NSTextCheckingResult) Init() NSTextCheckingResult {
	rv := objc.Send[NSTextCheckingResult](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextCheckingResult) Autorelease() NSTextCheckingResult {
	rv := objc.Send[NSTextCheckingResult](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextCheckingResult creates a new NSTextCheckingResult instance.
func NewNSTextCheckingResult() NSTextCheckingResult {
	class := getNSTextCheckingResultClass()
	rv := objc.Send[NSTextCheckingResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewTextCheckingResultWithCoder(coder INSCoder) NSTextCheckingResult {
	instance := getNSTextCheckingResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextCheckingResultFromID(rv)
}







// Returns the result type that the range represents.
//
// idx: The index of the result.
//
// # Return Value
// 
// The range of the result.
//
// # Discussion
// 
// A result must have at least one range, but may optionally have more, for
// example, to represent regular expression capture groups.
// 
// Passing [RangeAtIndex] the value `0` always returns the value of the
// [Range] property. Additional ranges, if any, will have indexes from `1` to
// `numberOfRanges``-1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range(at:)
func (t NSTextCheckingResult) RangeAtIndex(idx uint) NSRange {
	rv := objc.Send[NSRange](t.ID, objc.Sel("rangeAtIndex:"), idx)
	return NSRange(rv)
}

// Returns a new text checking result after adjusting the ranges as specified
// by the offset.
//
// offset: The amount the ranges are adjusted.
//
// # Return Value
// 
// A new [NSTextCheckingResult] instance with the adjusted range or ranges.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/adjustingRanges(offset:)
func (t NSTextCheckingResult) ResultByAdjustingRangesWithOffset(offset int) INSTextCheckingResult {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resultByAdjustingRangesWithOffset:"), offset)
	return NSTextCheckingResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range(withName:)
func (t NSTextCheckingResult) RangeWithName(name string) NSRange {
	rv := objc.Send[NSRange](t.ID, objc.Sel("rangeWithName:"), objc.String(name))
	return NSRange(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (t NSTextCheckingResult) InitWithCoder(coder INSCoder) NSTextCheckingResult {
	rv := objc.Send[NSTextCheckingResult](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (t NSTextCheckingResult) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Creates and returns a text checking result with the specified replacement
// string.
//
// range: The range of the detected result.
//
// replacementString: The replacement string.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeReplacement].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/replacementCheckingResult(range:replacementString:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) ReplacementCheckingResultWithRangeReplacementString(range_ NSRange, replacementString string) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("replacementCheckingResultWithRange:replacementString:"), range_, objc.String(replacementString))
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a type checking result with the specified regular
// expression data.
//
// ranges: A C array of ranges, which must have at least one element, and the first
// element represents the overall range.
//
// count: The number of items in the `ranges` array.
//
// regularExpression: The regular expression.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeRegularExpression].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpressionCheckingResult(ranges:count:regularExpression:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) RegularExpressionCheckingResultWithRangesCountRegularExpression(ranges []NSRangePointer, count uint, regularExpression INSRegularExpression) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("regularExpressionCheckingResultWithRanges:count:regularExpression:"), objc.CArray(ranges), count, regularExpression)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified URL.
//
// range: The range of the detected result.
//
// url: The URL.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeLink].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/linkCheckingResult(range:url:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) LinkCheckingResultWithRangeURL(range_ NSRange, url INSURL) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("linkCheckingResultWithRange:URL:"), range_, url)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified address
// components.
//
// range: The range of the detected result.
//
// components: A dictionary containing the address components. The dictionary keys are
// described in [Keys for Address Components].
// //
// [Keys for Address Components]: https://developer.apple.com/documentation/Foundation/keys-for-address-components
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeAddress].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/addressCheckingResult(range:components:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) AddressCheckingResultWithRangeComponents(range_ NSRange, components INSDictionary) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("addressCheckingResultWithRange:components:"), range_, components)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified transit
// information.
//
// range: The range of the detected result.
//
// components: A dictionary containing the transit components. The currently supported
// keys are [airline] and [flight].
// //
// [airline]: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/airline
// [flight]: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey/flight
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeTransitInformation].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/transitInformationCheckingResult(range:components:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) TransitInformationCheckingResultWithRangeComponents(range_ NSRange, components INSDictionary) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("transitInformationCheckingResultWithRange:components:"), range_, components)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified phone number.
//
// range: The range of the detected result.
//
// phoneNumber: The phone number.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypePhoneNumber].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumberCheckingResult(range:phoneNumber:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) PhoneNumberCheckingResultWithRangePhoneNumber(range_ NSRange, phoneNumber string) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("phoneNumberCheckingResultWithRange:phoneNumber:"), range_, objc.String(phoneNumber))
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified date.
//
// range: The range of the detected result.
//
// date: The detected date.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeDate].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/dateCheckingResult(range:date:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) DateCheckingResultWithRangeDate(range_ NSRange, date INSDate) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("dateCheckingResultWithRange:date:"), range_, date)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified date, time
// zone, and duration.
//
// range: The range of the detected result.
//
// date: The detected date.
//
// timeZone: The detected time zone.
//
// duration: The detected duration.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeDate].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/dateCheckingResult(range:date:timeZone:duration:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) DateCheckingResultWithRangeDateTimeZoneDuration(range_ NSRange, date INSDate, timeZone INSTimeZone, duration float64) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("dateCheckingResultWithRange:date:timeZone:duration:"), range_, date, timeZone, duration)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified dash
// corrected replacement string.
//
// range: The range of the detected result.
//
// replacementString: The replacement string.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeDash].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/dashCheckingResult(range:replacementString:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) DashCheckingResultWithRangeReplacementString(range_ NSRange, replacementString string) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("dashCheckingResultWithRange:replacementString:"), range_, objc.String(replacementString))
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified
// quote-balanced replacement string.
//
// range: The range of the detected result.
//
// replacementString: The replacement string.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeQuote].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/quoteCheckingResult(range:replacementString:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) QuoteCheckingResultWithRangeReplacementString(range_ NSRange, replacementString string) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("quoteCheckingResultWithRange:replacementString:"), range_, objc.String(replacementString))
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the range of a misspelled
// word.
//
// range: The range of the detected result.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeSpelling].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/spellCheckingResult(range:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) SpellCheckingResultWithRange(range_ NSRange) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("spellCheckingResultWithRange:"), range_)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result after detecting a possible
// correction.
//
// range: The range of the detected result.
//
// replacementString: The suggested replacement string.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeSpelling].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/correctionCheckingResult(range:replacementString:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) CorrectionCheckingResultWithRangeReplacementString(range_ NSRange, replacementString string) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("correctionCheckingResultWithRange:replacementString:"), range_, objc.String(replacementString))
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified orthography.
//
// range: The range of the detected result.
//
// orthography: An orthography object that describes the script.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeOrthography].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/orthographyCheckingResult(range:orthography:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) OrthographyCheckingResultWithRangeOrthography(range_ NSRange, orthography INSOrthography) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("orthographyCheckingResultWithRange:orthography:"), range_, orthography)
	return NSTextCheckingResultFromID(rv)
}

// Creates and returns a text checking result with the specified array of
// grammatical errors.
//
// range: The range of the detected result.
//
// details: An array of details regarding the grammatical errors. This array of strings
// is suitable for presenting to the user.
//
// # Return Value
// 
// Returns an [NSTextCheckingResult] with the specified [Range] and a
// [ResultType] of [TextCheckingTypeGrammar].
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/grammarCheckingResult(range:details:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) GrammarCheckingResultWithRangeDetails(range_ NSRange, details INSDictionary) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("grammarCheckingResultWithRange:details:"), range_, details)
	return NSTextCheckingResultFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/correctionCheckingResult(range:replacementString:alternativeStrings:)
func (_NSTextCheckingResultClass NSTextCheckingResultClass) CorrectionCheckingResultWithRangeReplacementStringAlternativeStrings(range_ NSRange, replacementString string, alternativeStrings []string) NSTextCheckingResult {
	rv := objc.Send[objc.ID](objc.ID(_NSTextCheckingResultClass.class), objc.Sel("correctionCheckingResultWithRange:replacementString:alternativeStrings:"), range_, objc.String(replacementString), objectivec.StringSliceToNSArray(alternativeStrings))
	return NSTextCheckingResultFromID(rv)
}








// Returns the range of the result that the receiver represents.
//
// # Discussion
// 
// This property will be present for all returned [NSTextCheckingResult]
// instances.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/range
func (t NSTextCheckingResult) Range() NSRange {
	rv := objc.Send[NSRange](t.ID, objc.Sel("range"))
	return NSRange(rv)
}



// Returns the text checking result type that the receiver represents.
//
// # Discussion
// 
// The possible result types for the built in checking capabilities are
// described in [NSTextCheckingResult.CheckingType].
// 
// This property will be present for all returned [NSTextCheckingResult]
// instances.
//
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/resultType
func (t NSTextCheckingResult) ResultType() NSTextCheckingType {
	rv := objc.Send[NSTextCheckingType](t.ID, objc.Sel("resultType"))
	return NSTextCheckingType(rv)
}



// Returns the number of ranges.
//
// # Discussion
// 
// A result must have at least one range, but may optionally have more (for
// example, to represent regular expression capture groups).
// 
// Passing [RangeAtIndex] the value `0` always returns the value of the the
// [Range] property. Additional ranges, if any, will have indexes from `1` to
// `numberOfRanges``-1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/numberOfRanges
func (t NSTextCheckingResult) NumberOfRanges() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("numberOfRanges"))
	return rv
}



// A replacement string from one of a number of replacement checking results.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/replacementString
func (t NSTextCheckingResult) ReplacementString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("replacementString"))
	return NSStringFromID(rv).String()
}



// The regular expression of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/regularExpression
func (t NSTextCheckingResult) RegularExpression() INSRegularExpression {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("regularExpression"))
	return NSRegularExpressionFromID(objc.ID(rv))
}



// A dictionary containing the components of a type checking result.
//
// # Discussion
// 
// Currently used by the transit checking result. The supported keys are
// located in [Keys for Transit Components].
//
// [Keys for Transit Components]: https://developer.apple.com/documentation/Foundation/keys-for-transit-components
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/components
func (t NSTextCheckingResult) Components() INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("components"))
	return NSDictionaryFromID(objc.ID(rv))
}



// The URL of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/url
func (t NSTextCheckingResult) URL() INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("URL"))
	return NSURLFromID(objc.ID(rv))
}



// The address dictionary of a type checking result.
//
// # Discussion
// 
// The dictionary keys are described in [Keys for Address Components].
//
// [Keys for Address Components]: https://developer.apple.com/documentation/Foundation/keys-for-address-components
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/addressComponents
func (t NSTextCheckingResult) AddressComponents() INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("addressComponents"))
	return NSDictionaryFromID(objc.ID(rv))
}



// The phone number of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/phoneNumber
func (t NSTextCheckingResult) PhoneNumber() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("phoneNumber"))
	return NSStringFromID(rv).String()
}



// The date component of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/date
func (t NSTextCheckingResult) Date() INSDate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("date"))
	return NSDateFromID(objc.ID(rv))
}



// The duration component of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/duration
func (t NSTextCheckingResult) Duration() float64 {
	rv := objc.Send[NSTimeInterval](t.ID, objc.Sel("duration"))
	return float64(rv)
}



// The time zone component of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/timeZone
func (t NSTextCheckingResult) TimeZone() INSTimeZone {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("timeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}



// The detected orthography of a type checking result.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/orthography
func (t NSTextCheckingResult) Orthography() INSOrthography {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("orthography"))
	return NSOrthographyFromID(objc.ID(rv))
}



// The details of a located grammatical type checking result.
//
// # Discussion
// 
// This array of strings is suitable for presenting to the user.
//
// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/grammarDetails
func (t NSTextCheckingResult) GrammarDetails() INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("grammarDetails"))
	return NSDictionaryFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/alternativeStrings
func (t NSTextCheckingResult) AlternativeStrings() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("alternativeStrings"))
	return objc.ConvertSliceToStrings(rv)
}



// A value indicating that a requested item couldn’t be found or doesn’t
// exist.
//
// See: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (t NSTextCheckingResult) NSNotFound() int {
	rv := objc.Send[int](t.ID, objc.Sel("NSNotFound"))
	return rv
}
func (t NSTextCheckingResult) SetNSNotFound(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSNotFound:"), value)
}














			// Protocol methods for NSCopying
			















