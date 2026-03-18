// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRegularExpression] class.
var (
	_NSRegularExpressionClass     NSRegularExpressionClass
	_NSRegularExpressionClassOnce sync.Once
)

func getNSRegularExpressionClass() NSRegularExpressionClass {
	_NSRegularExpressionClassOnce.Do(func() {
		_NSRegularExpressionClass = NSRegularExpressionClass{class: objc.GetClass("NSRegularExpression")}
	})
	return _NSRegularExpressionClass
}

// GetNSRegularExpressionClass returns the class object for NSRegularExpression.
func GetNSRegularExpressionClass() NSRegularExpressionClass {
	return getNSRegularExpressionClass()
}

type NSRegularExpressionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRegularExpressionClass) Alloc() NSRegularExpression {
	rv := objc.Send[NSRegularExpression](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An immutable representation of a compiled regular expression that you apply
// to Unicode strings.
//
// # Overview
// 
// The fundamental matching method for [NSRegularExpression] is a Block
// iterator method that allows clients to supply a Block object which will be
// invoked each time the regular expression matches a portion of the target
// string. There are additional convenience methods for returning all the
// matches as an array, the total number of matches, the first match, and the
// range of the first match.
// 
// An individual match is represented by an instance of the
// [NSTextCheckingResult] class, which carries information about the overall
// matched range (via its [NSRegularExpression.Range] property), and the range of each individual
// capture group (via the [RangeAtIndex] method). For basic
// [NSRegularExpression] objects, these match results will be of type
// [TextCheckingTypeRegularExpression], but subclasses may use other types.
// 
// # Examples Using NSRegularExpression
// 
// What follows are a set of graduated examples for using the
// [NSRegularExpression] class. All these examples use the regular expression
// `\\b(a|b)(c|d)\\b` as their regular expression.
// 
// This snippet creates a regular expression to match two-letter words, in
// which the first letter is “a” or “b” and the second letter is
// “c” or “d”. Specifying [RegularExpressionCaseInsensitive] means
// that matches will be case-insensitive, so this will match “BC”,
// “aD”, and so forth, as well as their lower-case equivalents.
// 
// The [NSRegularExpression.NumberOfMatchesInStringOptionsRange] method provides a simple
// mechanism for counting the number of matches in a given range of a string.
// 
// If you are interested only in the overall range of the first match, the
// [NSRegularExpression.RangeOfFirstMatchInStringOptionsRange] method provides it for you. Some
// regular expressions (though not the example pattern) can successfully match
// a zero-length range, so the comparison of the resulting range with
// `{NSNotFound, 0}` is the most reliable way to determine whether there was a
// match or not.
// 
// The example regular expression contains two capture groups, corresponding
// to the two sets of parentheses, one for the first letter, and one for the
// second. If you are interested in more than just the overall matched range,
// you want to obtain an [NSTextCheckingResult] object corresponding to a
// given match. This object provides information about the overall matched
// range, via its [NSRegularExpression.Range] property, and also supplies the capture group
// ranges, via the [RangeAtIndex] method. The first capture group range is
// given by `[result 1]`, the second by `[result 2]`. Sending a result the
// [RangeAtIndex] message and passing `0` is equivalent to `[result range]`.
// 
// If the result returned is non-`nil`, then `[result range]` will always be a
// valid range, so it is not necessary to compare it against `{NSNotFound,
// 0}`. However, for some regular expressions (though not the example pattern)
// some capture groups may or may not participate in a given match. If a given
// capture group does not participate in a given match, then `[result idx]`
// will return `{NSNotFound, 0}`.
// 
// The [NSRegularExpression.MatchesInStringOptionsRange] returns all the matching results.
// 
// The [NSRegularExpression.FirstMatchInStringOptionsRange] method is similar to
// [NSRegularExpression.MatchesInStringOptionsRange] but it returns only the first match.
// 
// The Block enumeration method
// [NSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock] is the most general and
// flexible of the matching methods of [NSRegularExpression]. It allows you to
// iterate through matches in a string, performing arbitrary actions on each
// as specified by the code in the Block and to stop partway through if
// desired. In the following example case, the iteration is stopped after a
// certain number of matches have been found.
// 
// If neither of the special options [MatchingReportProgress] or
// [MatchingReportCompletion] is specified, then the result argument to the
// Block is guaranteed to be non-`nil`, and as mentioned before, it is
// guaranteed to have a valid overall range. See
// [NSRegularExpression.MatchingOptions] for the significance of
// [MatchingReportProgress] or [MatchingReportCompletion].
// 
// [NSRegularExpression] also provides simple methods for performing
// find-and-replace operations on a string. The following example returns a
// modified copy, but there is a corresponding method for modifying a mutable
// string in place. The template specifies what is to be used to replace each
// match, with `$0` representing the contents of the overall matched range,
// `$1` representing the contents of the first capture group, and so on. In
// this case, the template reverses the two letters of the word.
// 
// # Concurrency and Thread Safety
// 
// [NSRegularExpression] is designed to be immutable and thread safe, so that
// a single instance can be used in matching operations on multiple threads at
// once. However, the string on which it is operating should not be mutated
// during the course of a matching operation, whether from another thread or
// from within the Block used in the iteration.
// 
// # Regular Expression Syntax
// 
// The following tables describe the character expressions used by the regular
// expression to match patterns within a string, the pattern operators that
// specify how many times a pattern is matched and additional matching
// restrictions, and the last table specifies flags that can be included in
// the regular expression pattern that specify search behavior over multiple
// lines (these flags can also be specified using the
// [NSRegularExpression.Options] option flags.
// 
// # Regular Expression Metacharacters
// 
// Table 1: Character sequences used to match characters within a string.
// 
// [Table data omitted]
// 
// # Regular Expression Operators
// 
// Table 2: Regular expression operators.
// 
// [Table data omitted]
// 
// # Template Matching Format
// 
// The [NSRegularExpression] class provides find-and-replace methods for both
// immutable and mutable strings using the technique of template matching.
// 
// Table 3: Find-and-replace syntax.
// 
// [Table data omitted]
// 
// The replacement string is treated as a template, with `$0` being replaced
// by the contents of the matched range, `$1` by the contents of the first
// capture group, and so on. Additional digits beyond the maximum required to
// represent the number of capture groups will be treated as ordinary
// characters, as will a `$` not followed by digits. Backslash will escape
// both `$` and `\`.
// 
// # Flag Options
// 
// The following flags control various aspects of regular expression matching.
// These flag values may be specified within the pattern using the
// `(?ismx-ismx)` pattern options. Equivalent behaviors can be specified for
// the entire pattern when an [NSRegularExpression] is initialized, using the
// [NSRegularExpression.Options] option flags.
// 
// Table 4: Regular expression matching flags.
// 
// [Table data omitted]
// 
// # Performance
// 
// [NSRegularExpression] implements a nondeterministic finite automaton
// matching engine. As such, complex regular expression patterns containing
// multiple `*` or `+` operators may result in poor performance when
// attempting to perform matches — particularly failing to match a given
// input. For more information, see the [“Performance Tips” section of the
// ICU User Guide].
// 
// # ICU License
// 
// Tables 1, 2, 3, and 4 are reproduced from the ICU User Guide, Copyright (c)
// 2000 - 2009 IBM and Others, which are licensed under the following terms:
// 
// COPYRIGHT AND PERMISSION NOTICE
// 
// Copyright (c) 1995-2009 International Business Machines Corporation and
// others. All rights reserved.
// 
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// “Software”), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, and/or sell copies of the Software, and to permit persons to
// whom the Software is furnished to do so, provided that the above copyright
// notice(s) and this permission notice appear in all copies of the Software
// and that both the above copyright notice(s) and this permission notice
// appear in supporting documentation.
// 
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT OF THIRD PARTY
// RIGHTS. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR HOLDERS INCLUDED IN THIS
// NOTICE BE LIABLE FOR ANY CLAIM, OR ANY SPECIAL INDIRECT OR CONSEQUENTIAL
// DAMAGES, OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR
// PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS
// ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS
// SOFTWARE.
// 
// Except as contained in this notice, the name of a copyright holder shall
// not be used in advertising or otherwise to promote the sale, use or other
// dealings in this Software without prior written authorization of the
// copyright holder.
// 
// All trademarks and registered trademarks mentioned herein are the property
// of their respective owners.
//
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
// [NSRegularExpression.Options]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
// [“Performance Tips” section of the ICU User Guide]: http://userguide.icu-project.org/strings/regexp#TOC-Performance-Tips
//
// # Creating Regular Expressions
//
//   - [NSRegularExpression.InitWithPatternOptionsError]: Returns an initialized NSRegularExpression instance with the specified regular expression pattern and options.
//
// # Getting the Regular Expression and Options
//
//   - [NSRegularExpression.Pattern]: Returns the regular expression pattern.
//   - [NSRegularExpression.Options]: Returns the options used when the regular expression option was created.
//   - [NSRegularExpression.NumberOfCaptureGroups]: Returns the number of capture groups in the regular expression.
//
// # Searching Strings Using Regular Expressions
//
//   - [NSRegularExpression.NumberOfMatchesInStringOptionsRange]: Returns the number of matches of the regular expression within the specified range of the string.
//   - [NSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock]: Enumerates the string allowing the Block to handle each regular expression match.
//   - [NSRegularExpression.MatchesInStringOptionsRange]: Returns an array containing all the matches of the regular expression in the string.
//   - [NSRegularExpression.FirstMatchInStringOptionsRange]: Returns the first match of the regular expression within the specified range of the string.
//   - [NSRegularExpression.RangeOfFirstMatchInStringOptionsRange]: Returns the range of the first match of the regular expression within the specified range of the string.
//
// # Replacing Strings Using Regular Expressions
//
//   - [NSRegularExpression.ReplaceMatchesInStringOptionsRangeWithTemplate]: Replaces regular expression matches within the mutable string using the template string.
//   - [NSRegularExpression.StringByReplacingMatchesInStringOptionsRangeWithTemplate]: Returns a new string containing matching regular expressions replaced with the template string.
//
// # Custom Replace Functionality
//
//   - [NSRegularExpression.ReplacementStringForResultInStringOffsetTemplate]: Used to perform template substitution for a single result for clients implementing their own replace functionality.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression
type NSRegularExpression struct {
	objectivec.Object
}

// NSRegularExpressionFromID constructs a [NSRegularExpression] from an objc.ID.
//
// An immutable representation of a compiled regular expression that you apply
// to Unicode strings.
func NSRegularExpressionFromID(id objc.ID) NSRegularExpression {
	return NSRegularExpression{objectivec.Object{ID: id}}
}
// NOTE: NSRegularExpression adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRegularExpression] class.
//
// # Creating Regular Expressions
//
//   - [INSRegularExpression.InitWithPatternOptionsError]: Returns an initialized NSRegularExpression instance with the specified regular expression pattern and options.
//
// # Getting the Regular Expression and Options
//
//   - [INSRegularExpression.Pattern]: Returns the regular expression pattern.
//   - [INSRegularExpression.Options]: Returns the options used when the regular expression option was created.
//   - [INSRegularExpression.NumberOfCaptureGroups]: Returns the number of capture groups in the regular expression.
//
// # Searching Strings Using Regular Expressions
//
//   - [INSRegularExpression.NumberOfMatchesInStringOptionsRange]: Returns the number of matches of the regular expression within the specified range of the string.
//   - [INSRegularExpression.EnumerateMatchesInStringOptionsRangeUsingBlock]: Enumerates the string allowing the Block to handle each regular expression match.
//   - [INSRegularExpression.MatchesInStringOptionsRange]: Returns an array containing all the matches of the regular expression in the string.
//   - [INSRegularExpression.FirstMatchInStringOptionsRange]: Returns the first match of the regular expression within the specified range of the string.
//   - [INSRegularExpression.RangeOfFirstMatchInStringOptionsRange]: Returns the range of the first match of the regular expression within the specified range of the string.
//
// # Replacing Strings Using Regular Expressions
//
//   - [INSRegularExpression.ReplaceMatchesInStringOptionsRangeWithTemplate]: Replaces regular expression matches within the mutable string using the template string.
//   - [INSRegularExpression.StringByReplacingMatchesInStringOptionsRangeWithTemplate]: Returns a new string containing matching regular expressions replaced with the template string.
//
// # Custom Replace Functionality
//
//   - [INSRegularExpression.ReplacementStringForResultInStringOffsetTemplate]: Used to perform template substitution for a single result for clients implementing their own replace functionality.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression
type INSRegularExpression interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating Regular Expressions

	// Returns an initialized NSRegularExpression instance with the specified regular expression pattern and options.
	InitWithPatternOptionsError(pattern string, options NSRegularExpressionOptions) (NSRegularExpression, error)

	// Topic: Getting the Regular Expression and Options

	// Returns the regular expression pattern.
	Pattern() string
	// Returns the options used when the regular expression option was created.
	Options() NSRegularExpressionOptions
	// Returns the number of capture groups in the regular expression.
	NumberOfCaptureGroups() uint

	// Topic: Searching Strings Using Regular Expressions

	// Returns the number of matches of the regular expression within the specified range of the string.
	NumberOfMatchesInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) uint
	// Enumerates the string allowing the Block to handle each regular expression match.
	EnumerateMatchesInStringOptionsRangeUsingBlock(string_ string, options NSMatchingOptions, range_ NSRange, block unsafe.Pointer)
	// Returns an array containing all the matches of the regular expression in the string.
	MatchesInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) []NSTextCheckingResult
	// Returns the first match of the regular expression within the specified range of the string.
	FirstMatchInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) INSTextCheckingResult
	// Returns the range of the first match of the regular expression within the specified range of the string.
	RangeOfFirstMatchInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) NSRange

	// Topic: Replacing Strings Using Regular Expressions

	// Replaces regular expression matches within the mutable string using the template string.
	ReplaceMatchesInStringOptionsRangeWithTemplate(string_ INSMutableString, options NSMatchingOptions, range_ NSRange, templ string) uint
	// Returns a new string containing matching regular expressions replaced with the template string.
	StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ string, options NSMatchingOptions, range_ NSRange, templ string) string

	// Topic: Custom Replace Functionality

	// Used to perform template substitution for a single result for clients implementing their own replace functionality.
	ReplacementStringForResultInStringOffsetTemplate(result INSTextCheckingResult, string_ string, offset int, templ string) string

	// A value indicating that a requested item couldn’t be found or doesn’t exist.
	NSNotFound() int
	SetNSNotFound(value int)
	// Returns the range of the result that the receiver represents.
	Range() NSRange
	SetRange(value NSRange)
	// Matches a regular expression.
	RegularExpression() NSTextCheckingType
	SetRegularExpression(value NSTextCheckingType)
	// Call the Block once after the completion of any matching. This option has no effect for methods other than 
	ReportCompletion() NSMatchingOptions
	SetReportCompletion(value NSMatchingOptions)
	// Call the Block periodically during long-running match operations. This option has no effect for methods other than 
	ReportProgress() NSMatchingOptions
	SetReportProgress(value NSMatchingOptions)
}

// Init initializes the instance.
func (r NSRegularExpression) Init() NSRegularExpression {
	rv := objc.Send[NSRegularExpression](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRegularExpression) Autorelease() NSRegularExpression {
	rv := objc.Send[NSRegularExpression](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRegularExpression creates a new NSRegularExpression instance.
func NewNSRegularExpression() NSRegularExpression {
	class := getNSRegularExpressionClass()
	rv := objc.Send[NSRegularExpression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewRegularExpressionWithCoder(coder INSCoder) NSRegularExpression {
	instance := getNSRegularExpressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSRegularExpressionFromID(rv)
}

// Returns an initialized NSRegularExpression instance with the specified
// regular expression pattern and options.
//
// pattern: The regular expression pattern to compile.
//
// options: The regular expression options that are applied to the expression during
// matching. See [NSRegularExpression.Options] for possible values.
// //
// [NSRegularExpression.Options]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
//
// # Return Value
// 
// An instance of [NSRegularExpression] for the specified regular expression
// and options.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/init(pattern:options:)
func NewRegularExpressionWithPatternOptionsError(pattern string, options NSRegularExpressionOptions) (NSRegularExpression, error) {
	var errorPtr objc.ID
	instance := getNSRegularExpressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPattern:options:error:"), objc.String(pattern), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSRegularExpressionFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSRegularExpressionFromID(rv), nil
}

// Returns an initialized NSRegularExpression instance with the specified
// regular expression pattern and options.
//
// pattern: The regular expression pattern to compile.
//
// options: The regular expression options that are applied to the expression during
// matching. See [NSRegularExpression.Options] for possible values.
// //
// [NSRegularExpression.Options]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
//
// # Return Value
// 
// An instance of [NSRegularExpression] for the specified regular expression
// and options.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/init(pattern:options:)
func (r NSRegularExpression) InitWithPatternOptionsError(pattern string, options NSRegularExpressionOptions) (NSRegularExpression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("initWithPattern:options:error:"), objc.String(pattern), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSRegularExpression{}, NSErrorFrom(errorPtr)
	}
	return NSRegularExpressionFromID(rv), nil

}

// Returns the number of matches of the regular expression within the
// specified range of the string.
//
// string: The string to search.
//
// options: The matching options to use. See [NSRegularExpression.MatchingOptions] for
// possible values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to search.
//
// # Return Value
// 
// The number of matches of the regular expression.
//
// # Discussion
// 
// This is a convenience method that calls
// [EnumerateMatchesInStringOptionsRangeUsingBlock].
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/numberOfMatches(in:options:range:)
func (r NSRegularExpression) NumberOfMatchesInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) uint {
	rv := objc.Send[uint](r.ID, objc.Sel("numberOfMatchesInString:options:range:"), objc.String(string_), options, range_)
	return rv
}

// Enumerates the string allowing the Block to handle each regular expression
// match.
//
// string: The string.
//
// options: The matching options to report. See [NSRegularExpression.MatchingOptions]
// for the supported values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to test.
//
// block: The Block enumerates the matches of the regular expression in the string.
// 
// The block takes three arguments:
// 
// result: An [NSTextCheckingResult] specifying the match. This result gives
// the overall matched range via its [Range] property, and the range of each
// individual capture group via its [RangeAtIndex] method. The range
// {[NSNotFound], 0} is returned if one of the capture groups did not
// participate in this particular match. flags: The current state of the
// matching progress. See [NSRegularExpression.MatchingFlags] for the possible
// values. stop: A reference to a Boolean value. The Block can set the value
// to [true] to stop further processing of the array. The stop argument is an
// out-only argument. You should only ever set this Boolean to [true] within
// the Block.
// 
// The Block returns void.
// //
// [NSRegularExpression.MatchingFlags]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the fundamental matching method for regular expressions and
// is suitable for overriding by subclassers. There are additional convenience
// methods for returning all the matches as an array, the total number of
// matches, the first match, and the range of the first match.
// 
// By default, the Block iterator method calls the Block precisely once for
// each match, with a non-`nil` `result` and the appropriate `flags`. The
// client may then stop the operation by setting the contents of `stop` to
// [true]. The `stop` argument is an out-only argument. You should only ever
// set this Boolean to [true] within the Block.
// 
// If the [MatchingReportProgress] matching option is specified, the Block
// will also be called periodically during long-running match operations, with
// `nil` result and [MatchingProgress] matching flag set in the Block’s
// `flags` parameter, at which point the client may again stop the operation
// by setting the contents of stop to [true].
// 
// If the [MatchingReportCompletion] matching option is specified, the Block
// object will be called once after matching is complete, with `nil` result
// and the [MatchingCompleted] matching flag is set in the `flags` passed to
// the Block, plus any additional relevant [NSRegularExpression.MatchingFlags]
// from among [MatchingHitEnd], [MatchingRequiredEnd], or
// [MatchingInternalError].
// 
// [MatchingProgress] and [MatchingCompleted] matching flags have no effect
// for methods other than this method.
// 
// The [MatchingHitEnd] matching flag is set in the `flags` passed to the
// Block if the current match operation reached the end of the search range.
// The [MatchingRequiredEnd] matching flag is set in the `flags` passed to the
// Block if the current match depended on the location of the end of the
// search range.
// 
// The [NSRegularExpression.MatchingFlags] matching flag is set in the `flags`
// passed to the block if matching failed due to an internal error (such as an
// expression requiring exponential memory allocations) without examining the
// entire search range.
// 
// The [MatchingAnchored], [MatchingWithTransparentBounds], and
// [MatchingWithoutAnchoringBounds] regular expression options, specified in
// the [Options] property specified when the regular expression instance is
// created, can apply to any match or replace method.
// 
// If [MatchingAnchored] matching option is specified, matches are limited to
// those at the start of the search range.
// 
// If [MatchingWithTransparentBounds] matching option is specified, matching
// may examine parts of the string beyond the bounds of the search range, for
// purposes such as word boundary detection, lookahead, etc.
// 
// If [MatchingWithoutAnchoringBounds] matching option is specified, `^` and
// `$` will not automatically match the beginning and end of the search range,
// but will still match the beginning and end of the entire string.
// 
// [MatchingWithTransparentBounds] and [MatchingWithoutAnchoringBounds]
// matching options have no effect if the search range covers the entire
// string.
//
// [NSRegularExpression.MatchingFlags]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/enumerateMatches(in:options:range:using:)
func (r NSRegularExpression) EnumerateMatchesInStringOptionsRangeUsingBlock(string_ string, options NSMatchingOptions, range_ NSRange, block unsafe.Pointer) {
	objc.Send[objc.ID](r.ID, objc.Sel("enumerateMatchesInString:options:range:usingBlock:"), objc.String(string_), options, range_, block)
}

// Returns an array containing all the matches of the regular expression in
// the string.
//
// string: The string to search.
//
// options: The matching options to use. See [NSRegularExpression.MatchingOptions] for
// possible values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to search.
//
// # Return Value
// 
// An array of [NSTextCheckingResult] objects. Each result gives the overall
// matched range via its [Range] property, and the range of each individual
// capture group via its [RangeAtIndex] method. The range {[NSNotFound], 0} is
// returned if one of the capture groups did not participate in this
// particular match.
//
// # Discussion
// 
// This is a convenience method that calls
// [EnumerateMatchesInStringOptionsRangeUsingBlock] passing the appropriate
// string, options, and range.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/matches(in:options:range:)
func (r NSRegularExpression) MatchesInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) []NSTextCheckingResult {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("matchesInString:options:range:"), objc.String(string_), options, range_)
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextCheckingResult {
		return NSTextCheckingResultFromID(id)
	})
}

// Returns the first match of the regular expression within the specified
// range of the string.
//
// string: The string to search.
//
// options: The matching options to use. See [NSRegularExpression.MatchingOptions] for
// possible values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to search.
//
// # Return Value
// 
// An [NSTextCheckingResult] object. This result gives the overall matched
// range via its [Range] property, and the range of each individual capture
// group via its [RangeAtIndex] method. The range {[NSNotFound], 0} is
// returned if one of the capture groups did not participate in this
// particular match.
//
// # Discussion
// 
// This is a convenience method that calls
// [EnumerateMatchesInStringOptionsRangeUsingBlock].
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/firstMatch(in:options:range:)
func (r NSRegularExpression) FirstMatchInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) INSTextCheckingResult {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("firstMatchInString:options:range:"), objc.String(string_), options, range_)
	return NSTextCheckingResultFromID(rv)
}

// Returns the range of the first match of the regular expression within the
// specified range of the string.
//
// string: The string to search.
//
// options: The matching options to use. See [NSRegularExpression.MatchingOptions] for
// possible values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to search.
//
// # Return Value
// 
// The range of the first match. Returns `{NSNotFound, 0}` if no match is
// found.
//
// # Discussion
// 
// This is a convenience method that calls
// [EnumerateMatchesInStringOptionsRangeUsingBlock].
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/rangeOfFirstMatch(in:options:range:)
func (r NSRegularExpression) RangeOfFirstMatchInStringOptionsRange(string_ string, options NSMatchingOptions, range_ NSRange) NSRange {
	rv := objc.Send[NSRange](r.ID, objc.Sel("rangeOfFirstMatchInString:options:range:"), objc.String(string_), options, range_)
	return NSRange(rv)
}

// Replaces regular expression matches within the mutable string using the
// template string.
//
// string: The mutable string to search and replace values within.
//
// options: The matching options to use. See [NSRegularExpression.MatchingOptions] for
// possible values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to search.
//
// templ: The substitution template used when replacing matching instances.
//
// # Return Value
// 
// The number of matches.
//
// # Discussion
// 
// See [NSRegularExpression] for the format of `templ`.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/replaceMatches(in:options:range:withTemplate:)
func (r NSRegularExpression) ReplaceMatchesInStringOptionsRangeWithTemplate(string_ INSMutableString, options NSMatchingOptions, range_ NSRange, templ string) uint {
	rv := objc.Send[uint](r.ID, objc.Sel("replaceMatchesInString:options:range:withTemplate:"), string_, options, range_, objc.String(templ))
	return rv
}

// Returns a new string containing matching regular expressions replaced with
// the template string.
//
// string: The string to search for values within.
//
// options: The matching options to use. See [NSRegularExpression.MatchingOptions] for
// possible values.
// //
// [NSRegularExpression.MatchingOptions]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
//
// range: The range of the string to search.
//
// templ: The substitution template used when replacing matching instances.
//
// # Return Value
// 
// A string with matching regular expressions replaced by the template string.
//
// # Discussion
// 
// See [NSRegularExpression] for the format of `templ`.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/stringByReplacingMatches(in:options:range:withTemplate:)
func (r NSRegularExpression) StringByReplacingMatchesInStringOptionsRangeWithTemplate(string_ string, options NSMatchingOptions, range_ NSRange, templ string) string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("stringByReplacingMatchesInString:options:range:withTemplate:"), objc.String(string_), options, range_, objc.String(templ))
	return NSStringFromID(rv).String()
}

// Used to perform template substitution for a single result for clients
// implementing their own replace functionality.
//
// result: The result of the single match.
//
// string: The string from which the result was matched.
//
// offset: The offset to be added to the location of the result in the string.
//
// templ: See [NSRegularExpression] for the format of `template`.
//
// # Return Value
// 
// A replacement string.
//
// # Discussion
// 
// For clients implementing their own replace functionality, this is a method
// to perform the template substitution for a single result, given the string
// from which the result was matched, an offset to be added to the location of
// the result in the string (for example, in cases that modifications to the
// string moved the result since it was matched), and a replacement template.
// 
// This is an advanced method that is used only if you wanted to iterate
// through a list of matches yourself and do the template replacement for each
// one, plus maybe some other calculation that you want to do in code, then
// you would use this at each step.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/replacementString(for:in:offset:template:)
func (r NSRegularExpression) ReplacementStringForResultInStringOffsetTemplate(result INSTextCheckingResult, string_ string, offset int, templ string) string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("replacementStringForResult:inString:offset:template:"), result, objc.String(string_), offset, objc.String(templ))
	return NSStringFromID(rv).String()
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (r NSRegularExpression) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (r NSRegularExpression) InitWithCoder(coder INSCoder) NSRegularExpression {
	rv := objc.Send[NSRegularExpression](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns a template string by adding backslash escapes as necessary to
// protect any characters that would match as pattern metacharacters
//
// string: The template string
//
// # Return Value
// 
// The escaped template string.
//
// # Discussion
// 
// Returns a string by adding backslash escapes as necessary to the given
// string, to escape any characters that would otherwise be treated as pattern
// metacharacters. You typically use this method to match on a particular
// string within a larger pattern.
// 
// For example, the string `"(N/A)"` contains the pattern metacharacters `(`,
// `/`, and `)`. The result of adding backslash escapes to this string is
// `"\\(N\\/A\\)"`.
// 
// See [NSRegularExpression] for the format of the resulting template string.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/escapedTemplate(for:)
func (_NSRegularExpressionClass NSRegularExpressionClass) EscapedTemplateForString(string_ string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSRegularExpressionClass.class), objc.Sel("escapedTemplateForString:"), objc.String(string_))
	return NSStringFromID(rv).String()
}

// Returns a string by adding backslash escapes as necessary to protect any
// characters that would match as pattern metacharacters.
//
// string: The string.
//
// # Return Value
// 
// The escaped string.
//
// # Discussion
// 
// Returns a string by adding backslash escapes as necessary to the given
// string, to escape any characters that would otherwise be treated as pattern
// metacharacters. You typically use this method to match on a particular
// string within a larger pattern.
// 
// For example, the string `"(N/A)"` contains the pattern metacharacters `(`,
// `/`, and `)`. The result of adding backslash escapes to this string is
// `"\\(N\\/A\\)"`.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/escapedPattern(for:)
func (_NSRegularExpressionClass NSRegularExpressionClass) EscapedPatternForString(string_ string) string {
	rv := objc.Send[objc.ID](objc.ID(_NSRegularExpressionClass.class), objc.Sel("escapedPatternForString:"), objc.String(string_))
	return NSStringFromID(rv).String()
}

// Creates an NSRegularExpression instance with the specified regular
// expression pattern and options.
//
// pattern: The regular expression pattern to compile.
//
// options: The matching options. See [NSRegularExpression.Options] for possible
// values. The values can be combined using the C-bitwise [OR] operator.
// //
// [NSRegularExpression.Options]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
//
// error: An out value that returns any error encountered during initialization.
// Returns an [NSError] object if the regular expression pattern is invalid;
// otherwise returns `nil`.
//
// # Return Value
// 
// An instance of [NSRegularExpression] for the specified regular expression
// and options.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/regularExpressionWithPattern:options:error:
func (_NSRegularExpressionClass NSRegularExpressionClass) RegularExpressionWithPatternOptionsError(pattern string, options NSRegularExpressionOptions) (NSRegularExpression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSRegularExpressionClass.class), objc.Sel("regularExpressionWithPattern:options:error:"), objc.String(pattern), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSRegularExpression{}, NSErrorFrom(errorPtr)
	}
	return NSRegularExpressionFromID(rv), nil

}

// Returns the regular expression pattern.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/pattern
func (r NSRegularExpression) Pattern() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("pattern"))
	return NSStringFromID(rv).String()
}

// Returns the options used when the regular expression option was created.
//
// # Discussion
// 
// The options property specifies aspects of the regular expression matching
// that are always used when matching the regular expression. For example, if
// the expression is case sensitive, allows comments, ignores metacharacters,
// etc. See [NSRegularExpression.Options] for a complete discussion of the
// possible constants and their meanings.
//
// [NSRegularExpression.Options]: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/options-swift.property
func (r NSRegularExpression) Options() NSRegularExpressionOptions {
	rv := objc.Send[NSRegularExpressionOptions](r.ID, objc.Sel("options"))
	return NSRegularExpressionOptions(rv)
}

// Returns the number of capture groups in the regular expression.
//
// # Discussion
// 
// A capture group consists of each possible match within a regular
// expression. Each capture group can then be used in a replacement template
// to insert that value into a replacement string.
// 
// This value puts a limit on the values of `n` for `$n` in templates, and it
// determines the number of ranges in the returned [NSTextCheckingResult]
// instances returned in the `match...` methods.
// 
// An exception will be generated if you attempt to access a result with an
// index value exceeding `numberOfCaptureGroups``-1`.
//
// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/numberOfCaptureGroups
func (r NSRegularExpression) NumberOfCaptureGroups() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("numberOfCaptureGroups"))
	return rv
}

// A value indicating that a requested item couldn’t be found or doesn’t
// exist.
//
// See: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (r NSRegularExpression) NSNotFound() int {
	rv := objc.Send[int](r.ID, objc.Sel("NSNotFound"))
	return rv
}
func (r NSRegularExpression) SetNSNotFound(value int) {
	objc.Send[struct{}](r.ID, objc.Sel("setNSNotFound:"), value)
}

// Returns the range of the result that the receiver represents.
//
// See: https://developer.apple.com/documentation/foundation/nstextcheckingresult/range
func (r NSRegularExpression) Range() NSRange {
	rv := objc.Send[NSRange](r.ID, objc.Sel("range"))
	return NSRange(rv)
}
func (r NSRegularExpression) SetRange(value NSRange) {
	objc.Send[struct{}](r.ID, objc.Sel("setRange:"), value)
}

// Matches a regular expression.
//
// See: https://developer.apple.com/documentation/foundation/nstextcheckingresult/checkingtype/regularexpression
func (r NSRegularExpression) RegularExpression() NSTextCheckingType {
	rv := objc.Send[NSTextCheckingType](r.ID, objc.Sel("NSTextCheckingTypeRegularExpression"))
	return NSTextCheckingType(rv)
}
func (r NSRegularExpression) SetRegularExpression(value NSTextCheckingType) {
	objc.Send[struct{}](r.ID, objc.Sel("setNSTextCheckingTypeRegularExpression:"), value)
}

// Call the Block once after the completion of any matching. This option has
// no effect for methods other than
//
// See: https://developer.apple.com/documentation/foundation/nsregularexpression/matchingoptions/reportcompletion
func (r NSRegularExpression) ReportCompletion() NSMatchingOptions {
	rv := objc.Send[NSMatchingOptions](r.ID, objc.Sel("NSMatchingReportCompletion"))
	return NSMatchingOptions(rv)
}
func (r NSRegularExpression) SetReportCompletion(value NSMatchingOptions) {
	objc.Send[struct{}](r.ID, objc.Sel("setNSMatchingReportCompletion:"), value)
}

// Call the Block periodically during long-running match operations. This
// option has no effect for methods other than
//
// See: https://developer.apple.com/documentation/foundation/nsregularexpression/matchingoptions/reportprogress
func (r NSRegularExpression) ReportProgress() NSMatchingOptions {
	rv := objc.Send[NSMatchingOptions](r.ID, objc.Sel("NSMatchingReportProgress"))
	return NSMatchingOptions(rv)
}
func (r NSRegularExpression) SetReportProgress(value NSMatchingOptions) {
	objc.Send[struct{}](r.ID, objc.Sel("setNSMatchingReportProgress:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

