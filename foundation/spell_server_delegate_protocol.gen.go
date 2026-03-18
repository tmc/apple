// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The optional methods implemented by the delegate of a spell server.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate
type NSSpellServerDelegate interface {
	objectivec.IObject
}

// NSSpellServerDelegateObject wraps an existing Objective-C object that conforms to the NSSpellServerDelegate protocol.
type NSSpellServerDelegateObject struct {
	objectivec.Object
}
func (o NSSpellServerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSpellServerDelegateObjectFromID constructs a [NSSpellServerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSpellServerDelegateObjectFromID(id objc.ID) NSSpellServerDelegateObject {
	return NSSpellServerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Gives the delegate the opportunity to analyze both the spelling and grammar
// simultaneously, which is more efficient.
//
// sender: Spell server making the analysis request.
//
// stringToCheck: String to analyze.
//
// offset: The offset in the string.
//
// checkingTypes: The text checking types to perform.
//
// options: A dictionary defining the actions to be taken while checking this string.
// See Constants in [NSSpellChecker] for the possible keys.
// //
// [NSSpellChecker]: https://developer.apple.com/documentation/AppKit/NSSpellChecker
//
// orthography: The identified orthography of `stringToCheck`. See [NSOrthography] for more
// information.
//
// wordCount: On output, returns by-reference the number of words from the beginning of
// the string object until the misspelled word (or the end of string).
//
// # Return Value
// 
// An array of NSTextCheckingResult instances of the spelling, grammar, or
// correction types, depending on the `checkingTypes` requested.
//
// # Discussion
// 
// This method is optional, but if implemented it will be called during the
// course of unified text checking via the [NSSpellChecker]
// [checkSpelling(of:startingAt:)] and
// [requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)]
// methods. This allows spelling and grammar checking to be performed
// simultaneously, which can be significantly more efficient, and allows the
// delegate to return autocorrection results as well.
// 
// If this method is not implemented, then unified text checking will call the
// separate spelling and grammar checking methods instead.
// 
// This method may be called repeatedly with strings representing different
// subranges of the string that was originally requested to be checked; the
// offset argument represents the offset of the portion passed in to this
// method within that original string, and should be added to the origin of
// the range in any [NSTextCheckingResult] returned.
//
// [checkSpelling(of:startingAt:)]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:)
// [requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:check:offset:types:options:orthography:wordCount:)

func (o NSSpellServerDelegateObject) SpellServerCheckStringOffsetTypesOptionsOrthographyWordCount(sender INSSpellServer, stringToCheck string, offset uint, checkingTypes NSTextCheckingTypes, options INSDictionary, orthography INSOrthography, wordCount unsafe.Pointer) []NSTextCheckingResult {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("spellServer:checkString:offset:types:options:orthography:wordCount:"), sender, objc.String(stringToCheck), offset, checkingTypes, options, orthography, wordCount)
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextCheckingResult {
		return NSTextCheckingResultFromID(id)
	})
	}

// Gives the delegate the opportunity to suggest guesses to the sender for the
// correct spelling of the given misspelled word in the specified language.
//
// sender: The [NSSpellServer] object that sent this message.
//
// word: The misspelled word.
//
// language: The language to use for the guesses.
//
// # Return Value
// 
// An array of [NSString] objects indicating possible correct spellings.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:suggestGuessesForWord:inLanguage:)

func (o NSSpellServerDelegateObject) SpellServerSuggestGuessesForWordInLanguage(sender INSSpellServer, word string, language string) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("spellServer:suggestGuessesForWord:inLanguage:"), sender, objc.String(word), objc.String(language))
	return objc.ConvertSliceToStrings(rv)
	}

// Gives the delegate the opportunity to customize the grammatical analysis of
// a given string.
//
// sender: Spell server satisfying a grammatical analysis request.
//
// stringToCheck: String to analyze.
//
// language: Language use in `string`. When `nil`, the language selected in the Spelling
// panel is used.
//
// details: On output, dictionaries describing grammar-analysis details within the
// flagged grammatical unit. See the [NSSpellServer] class for information
// about these dictionaries.
//
// # Return Value
// 
// Location of the first flagged grammatical unit within `string`.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:checkGrammarIn:language:details:)

func (o NSSpellServerDelegateObject) SpellServerCheckGrammarInStringLanguageDetails(sender INSSpellServer, stringToCheck string, language string, details INSDictionary) NSRange {
	
	rv := objc.Send[NSRange](o.ID, objc.Sel("spellServer:checkGrammarInString:language:details:"), sender, objc.String(stringToCheck), objc.String(language), details)
	return rv
	}

// Asks the delegate to search for a misspelled word in a given string, using
// the specified language, and marking the first misspelled word found by
// returning its range within the string.
//
// sender: The [NSSpellServer] object that sent this message.
//
// stringToCheck: The string to search for the misspelled word.
//
// language: The language to use for the search.
//
// wordCount: On output, returns by reference the number of words from the beginning of
// the string object until the misspelled word (or the end of string).
//
// countOnly: If [true], the method only counts the words in the string object and does
// not spell checking.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The range of the misspelled word within the given string.
//
// # Discussion
// 
// Send [IsWordInUserDictionariesCaseSensitive] to the spelling server to
// determine if the word exists in the user’s language dictionaries.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:findMisspelledWordIn:language:wordCount:countOnly:)

func (o NSSpellServerDelegateObject) SpellServerFindMisspelledWordInStringLanguageWordCountCountOnly(sender INSSpellServer, stringToCheck string, language string, wordCount unsafe.Pointer, countOnly bool) NSRange {
	
	rv := objc.Send[NSRange](o.ID, objc.Sel("spellServer:findMisspelledWordInString:language:wordCount:countOnly:"), sender, objc.String(stringToCheck), objc.String(language), wordCount, countOnly)
	return rv
	}

// Notifies the delegate that the sender has removed the specified word from
// the user’s list of acceptable words in the specified language.
//
// sender: The [NSSpellServer] object that removed the word.
//
// word: The word that was removed.
//
// language: The language of the removed word.
//
// # Discussion
// 
// If your delegate maintains a similar auxiliary word list, you may wish to
// edit the list accordingly.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:didForgetWord:inLanguage:)

func (o NSSpellServerDelegateObject) SpellServerDidForgetWordInLanguage(sender INSSpellServer, word string, language string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("spellServer:didForgetWord:inLanguage:"), sender, objc.String(word), objc.String(language))
	}

// Notifies the delegate that the sender has added the specified word to the
// user’s list of acceptable words in the specified language.
//
// sender: The [NSSpellServer] object that added the word.
//
// word: The word that was added.
//
// language: The language of the added word.
//
// # Discussion
// 
// If your delegate maintains a similar auxiliary word list, you may wish to
// edit the list accordingly.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:didLearnWord:inLanguage:)

func (o NSSpellServerDelegateObject) SpellServerDidLearnWordInLanguage(sender INSSpellServer, word string, language string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("spellServer:didLearnWord:inLanguage:"), sender, objc.String(word), objc.String(language))
	}

// This delegate method returns an array of possible word completions from the
// spell checker, based on a partially completed string and a given range.
//
// sender: The [NSSpellServer] object that sent this message.
//
// range: The range of the partially completed word.
//
// string: The string containing the partial word range.
//
// language: The language to use for the completion.
//
// # Return Value
// 
// An array of [NSString] objects indicating possible completions.
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:suggestCompletionsForPartialWordRange:in:language:)

func (o NSSpellServerDelegateObject) SpellServerSuggestCompletionsForPartialWordRangeInStringLanguage(sender INSSpellServer, range_ NSRange, string_ string, language string) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("spellServer:suggestCompletionsForPartialWordRange:inString:language:"), sender, range_, objc.String(string_), objc.String(language))
	return objc.ConvertSliceToStrings(rv)
	}

// Notifies the spell checker of the users’s response to a correction.
//
// sender: The spell server.
//
// response: The user’s response.
//
// correction: The corrected word. This should match the original correction.
//
// word: The original word. This should match the original correction.
//
// language: The language being edited. This should match the original correction.
//
// # Discussion
// 
// When the user accepts, rejects, or edits an autocorrection, the view
// notifies the [NSSpellChecker] class of what happened in the client
// application, and [NSSpellChecker] then invokes this method, so that it can
// record that and modify future autocorrection behavior based on what it has
// learned from the user’s actions.
//
// [NSSpellChecker]: https://developer.apple.com/documentation/AppKit/NSSpellChecker
//
// See: https://developer.apple.com/documentation/Foundation/NSSpellServerDelegate/spellServer(_:recordResponse:toCorrection:forWord:language:)

func (o NSSpellServerDelegateObject) SpellServerRecordResponseToCorrectionForWordLanguage(sender INSSpellServer, response uint, correction string, word string, language string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("spellServer:recordResponse:toCorrection:forWord:language:"), sender, response, objc.String(correction), objc.String(word), objc.String(language))
	}

