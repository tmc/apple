// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSpellChecker] class.
var (
	_NSSpellCheckerClass     NSSpellCheckerClass
	_NSSpellCheckerClassOnce sync.Once
)

func getNSSpellCheckerClass() NSSpellCheckerClass {
	_NSSpellCheckerClassOnce.Do(func() {
		_NSSpellCheckerClass = NSSpellCheckerClass{class: objc.GetClass("NSSpellChecker")}
	})
	return _NSSpellCheckerClass
}

// GetNSSpellCheckerClass returns the class object for NSSpellChecker.
func GetNSSpellCheckerClass() NSSpellCheckerClass {
	return getNSSpellCheckerClass()
}

type NSSpellCheckerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSpellCheckerClass) Alloc() NSSpellChecker {
	rv := objc.Send[NSSpellChecker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An interface to the Cocoa spell-checking service.
//
// # Overview
// 
// To handle all its spell checking, an app needs only one instance of
// [NSSpellChecker], known as the spell checker. Using the spell checker you
// manage the Spelling panel, in which the user can specify decisions about
// words that are suspect. The spell checker also offers the ability to
// provide word completions to augment the text completion system.
//
// # Configuring Spell Checkers Languages
//
//   - [NSSpellChecker.AvailableLanguages]: Provides a list of all available languages.
//   - [NSSpellChecker.UserPreferredLanguages]: Provides a subset of the available languages to be used for spell checking.
//   - [NSSpellChecker.AutomaticallyIdentifiesLanguages]: Sets whether the spell checker will automatically identify languages.
//   - [NSSpellChecker.SetAutomaticallyIdentifiesLanguages]
//   - [NSSpellChecker.Language]: Returns the current language used in spell checking.
//   - [NSSpellChecker.SetLanguage]: Returns whether the specified language is in the Spelling pop-up list.
//
// # Managing Panels
//
//   - [NSSpellChecker.SpellingPanel]: Returns the spell checker’s panel.
//   - [NSSpellChecker.SubstitutionsPanel]: Returns the substitutions panel.
//   - [NSSpellChecker.UpdateSpellingPanelWithGrammarStringDetail]: Specifies a grammar-analysis detail to highlight in the Spelling panel.
//   - [NSSpellChecker.UpdatePanels]: Updates the available panels to account for user changes.
//   - [NSSpellChecker.AccessoryView]: Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//   - [NSSpellChecker.SetAccessoryView]
//   - [NSSpellChecker.SubstitutionsPanelAccessoryViewController]: Sets the substitutions panel’s accessory view.
//   - [NSSpellChecker.SetSubstitutionsPanelAccessoryViewController]
//
// # Checking Strings for Spelling and Grammar
//
//   - [NSSpellChecker.CountWordsInStringLanguage]: Returns the number of words in the specified string.
//   - [NSSpellChecker.CheckSpellingOfStringStartingAt]: Starts the search for a misspelled word in `stringToCheck` starting at `startingOffset` within the string object.
//   - [NSSpellChecker.CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount]: Starts the search for a misspelled word in a string starting at specified offset within the string.
//   - [NSSpellChecker.CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails]: Initiates a grammatical analysis of a given string.
//   - [NSSpellChecker.CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount]: Requests unified text checking for the given range of the given string.
//   - [NSSpellChecker.RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler]: Requests that the string be checked in the background.
//   - [NSSpellChecker.GuessesForWordRangeInStringLanguageInSpellDocumentWithTag]: Returns an array of possible substitutions for the specified string.
//
// # Managing the Spell-Checking Process
//
//   - [NSSpellChecker.CloseSpellDocumentWithTag]: Notifies the receiver that the user has finished with the tagged document.
//   - [NSSpellChecker.IgnoreWordInSpellDocumentWithTag]: Instructs the spell checker to ignore all future occurrences of `wordToIgnore` in the document identified by `tag`.
//   - [NSSpellChecker.IgnoredWordsInSpellDocumentWithTag]: Returns the array of ignored words for a document identified by `tag`.
//   - [NSSpellChecker.SetIgnoredWordsInSpellDocumentWithTag]: Initializes the ignored-words document (a dictionary identified by `tag` with `someWords`), an array of words to ignore.
//   - [NSSpellChecker.SetWordFieldStringValue]: Sets the string that appears in the misspelled word field, using the string object `aString`.
//   - [NSSpellChecker.UpdateSpellingPanelWithMisspelledWord]: Causes the spell checker to update the Spelling panel’s misspelled-word field to reflect `word`.
//   - [NSSpellChecker.CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag]: Provides a list of complete words that the user might be trying to type based on a partial word in a given string.
//   - [NSSpellChecker.HasLearnedWord]: Indicates whether the spell checker has learned a given word.
//   - [NSSpellChecker.UnlearnWord]: Tells the spell checker to unlearn a given word.
//   - [NSSpellChecker.LearnWord]: Adds the word to the spell checker dictionary.
//   - [NSSpellChecker.UserQuotesArrayForLanguage]: Returns the default values for quote replacement.
//   - [NSSpellChecker.UserReplacementsDictionary]: Returns the dictionary used when replacing words.
//
// # Data Detector Interaction
//
//   - [NSSpellChecker.MenuForResultStringOptionsAtLocationInView]: Provides a menu containing contextual menu items suitable for certain kinds of detected results.
//
// # Automatic Spelling Correction
//
//   - [NSSpellChecker.CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag]: Returns a single proposed correction if a word is mis-spelled.
//   - [NSSpellChecker.RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag]: Records the user response to the correction indicator being displayed.
//   - [NSSpellChecker.DismissCorrectionIndicatorForView]: Dismisses the correction indicator for the specified view.
//
// # Instance Methods
//
//   - [NSSpellChecker.DeletesAutospaceBetweenStringAndStringLanguage]
//   - [NSSpellChecker.LanguageForWordRangeInStringOrthography]
//   - [NSSpellChecker.PreventsAutocorrectionBeforeStringLanguage]
//   - [NSSpellChecker.RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
//   - [NSSpellChecker.ShowInlinePredictionForCandidatesClient]
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker
type NSSpellChecker struct {
	objectivec.Object
}

// NSSpellCheckerFromID constructs a [NSSpellChecker] from an objc.ID.
//
// An interface to the Cocoa spell-checking service.
func NSSpellCheckerFromID(id objc.ID) NSSpellChecker {
	return NSSpellChecker{objectivec.Object{id}}
}
// NOTE: NSSpellChecker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSSpellChecker] class.
//
// # Configuring Spell Checkers Languages
//
//   - [INSSpellChecker.AvailableLanguages]: Provides a list of all available languages.
//   - [INSSpellChecker.UserPreferredLanguages]: Provides a subset of the available languages to be used for spell checking.
//   - [INSSpellChecker.AutomaticallyIdentifiesLanguages]: Sets whether the spell checker will automatically identify languages.
//   - [INSSpellChecker.SetAutomaticallyIdentifiesLanguages]
//   - [INSSpellChecker.Language]: Returns the current language used in spell checking.
//   - [INSSpellChecker.SetLanguage]: Returns whether the specified language is in the Spelling pop-up list.
//
// # Managing Panels
//
//   - [INSSpellChecker.SpellingPanel]: Returns the spell checker’s panel.
//   - [INSSpellChecker.SubstitutionsPanel]: Returns the substitutions panel.
//   - [INSSpellChecker.UpdateSpellingPanelWithGrammarStringDetail]: Specifies a grammar-analysis detail to highlight in the Spelling panel.
//   - [INSSpellChecker.UpdatePanels]: Updates the available panels to account for user changes.
//   - [INSSpellChecker.AccessoryView]: Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
//   - [INSSpellChecker.SetAccessoryView]
//   - [INSSpellChecker.SubstitutionsPanelAccessoryViewController]: Sets the substitutions panel’s accessory view.
//   - [INSSpellChecker.SetSubstitutionsPanelAccessoryViewController]
//
// # Checking Strings for Spelling and Grammar
//
//   - [INSSpellChecker.CountWordsInStringLanguage]: Returns the number of words in the specified string.
//   - [INSSpellChecker.CheckSpellingOfStringStartingAt]: Starts the search for a misspelled word in `stringToCheck` starting at `startingOffset` within the string object.
//   - [INSSpellChecker.CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount]: Starts the search for a misspelled word in a string starting at specified offset within the string.
//   - [INSSpellChecker.CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails]: Initiates a grammatical analysis of a given string.
//   - [INSSpellChecker.CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount]: Requests unified text checking for the given range of the given string.
//   - [INSSpellChecker.RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler]: Requests that the string be checked in the background.
//   - [INSSpellChecker.GuessesForWordRangeInStringLanguageInSpellDocumentWithTag]: Returns an array of possible substitutions for the specified string.
//
// # Managing the Spell-Checking Process
//
//   - [INSSpellChecker.CloseSpellDocumentWithTag]: Notifies the receiver that the user has finished with the tagged document.
//   - [INSSpellChecker.IgnoreWordInSpellDocumentWithTag]: Instructs the spell checker to ignore all future occurrences of `wordToIgnore` in the document identified by `tag`.
//   - [INSSpellChecker.IgnoredWordsInSpellDocumentWithTag]: Returns the array of ignored words for a document identified by `tag`.
//   - [INSSpellChecker.SetIgnoredWordsInSpellDocumentWithTag]: Initializes the ignored-words document (a dictionary identified by `tag` with `someWords`), an array of words to ignore.
//   - [INSSpellChecker.SetWordFieldStringValue]: Sets the string that appears in the misspelled word field, using the string object `aString`.
//   - [INSSpellChecker.UpdateSpellingPanelWithMisspelledWord]: Causes the spell checker to update the Spelling panel’s misspelled-word field to reflect `word`.
//   - [INSSpellChecker.CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag]: Provides a list of complete words that the user might be trying to type based on a partial word in a given string.
//   - [INSSpellChecker.HasLearnedWord]: Indicates whether the spell checker has learned a given word.
//   - [INSSpellChecker.UnlearnWord]: Tells the spell checker to unlearn a given word.
//   - [INSSpellChecker.LearnWord]: Adds the word to the spell checker dictionary.
//   - [INSSpellChecker.UserQuotesArrayForLanguage]: Returns the default values for quote replacement.
//   - [INSSpellChecker.UserReplacementsDictionary]: Returns the dictionary used when replacing words.
//
// # Data Detector Interaction
//
//   - [INSSpellChecker.MenuForResultStringOptionsAtLocationInView]: Provides a menu containing contextual menu items suitable for certain kinds of detected results.
//
// # Automatic Spelling Correction
//
//   - [INSSpellChecker.CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag]: Returns a single proposed correction if a word is mis-spelled.
//   - [INSSpellChecker.RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag]: Records the user response to the correction indicator being displayed.
//   - [INSSpellChecker.DismissCorrectionIndicatorForView]: Dismisses the correction indicator for the specified view.
//
// # Instance Methods
//
//   - [INSSpellChecker.DeletesAutospaceBetweenStringAndStringLanguage]
//   - [INSSpellChecker.LanguageForWordRangeInStringOrthography]
//   - [INSSpellChecker.PreventsAutocorrectionBeforeStringLanguage]
//   - [INSSpellChecker.RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler]
//   - [INSSpellChecker.ShowInlinePredictionForCandidatesClient]
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker
type INSSpellChecker interface {
	objectivec.IObject

	// Topic: Configuring Spell Checkers Languages

	// Provides a list of all available languages.
	AvailableLanguages() []string
	// Provides a subset of the available languages to be used for spell checking.
	UserPreferredLanguages() []string
	// Sets whether the spell checker will automatically identify languages.
	AutomaticallyIdentifiesLanguages() bool
	SetAutomaticallyIdentifiesLanguages(value bool)
	// Returns the current language used in spell checking.
	Language() string
	// Returns whether the specified language is in the Spelling pop-up list.
	SetLanguage(language string) bool

	// Topic: Managing Panels

	// Returns the spell checker’s panel.
	SpellingPanel() INSPanel
	// Returns the substitutions panel.
	SubstitutionsPanel() INSPanel
	// Specifies a grammar-analysis detail to highlight in the Spelling panel.
	UpdateSpellingPanelWithGrammarStringDetail(string_ string, detail foundation.INSDictionary)
	// Updates the available panels to account for user changes.
	UpdatePanels()
	// Makes a view an accessory of the Spelling panel by making it a subview of the panel’s content view.
	AccessoryView() INSView
	SetAccessoryView(value INSView)
	// Sets the substitutions panel’s accessory view.
	SubstitutionsPanelAccessoryViewController() INSViewController
	SetSubstitutionsPanelAccessoryViewController(value INSViewController)

	// Topic: Checking Strings for Spelling and Grammar

	// Returns the number of words in the specified string.
	CountWordsInStringLanguage(stringToCount string, language string) int
	// Starts the search for a misspelled word in `stringToCheck` starting at `startingOffset` within the string object.
	CheckSpellingOfStringStartingAt(stringToCheck string, startingOffset int) foundation.NSRange
	// Starts the search for a misspelled word in a string starting at specified offset within the string.
	CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, wordCount int) foundation.NSRange
	// Initiates a grammatical analysis of a given string.
	CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, details foundation.INSDictionary) foundation.NSRange
	// Requests unified text checking for the given range of the given string.
	CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, tag int, orthography foundation.NSOrthography, wordCount int) []foundation.NSTextCheckingResult
	// Requests that the string be checked in the background.
	RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, tag int, completionHandler ErrorHandler) int
	// Returns an array of possible substitutions for the specified string.
	GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.NSRange, string_ string, language string, tag int) []string

	// Topic: Managing the Spell-Checking Process

	// Notifies the receiver that the user has finished with the tagged document.
	CloseSpellDocumentWithTag(tag int)
	// Instructs the spell checker to ignore all future occurrences of `wordToIgnore` in the document identified by `tag`.
	IgnoreWordInSpellDocumentWithTag(wordToIgnore string, tag int)
	// Returns the array of ignored words for a document identified by `tag`.
	IgnoredWordsInSpellDocumentWithTag(tag int) []string
	// Initializes the ignored-words document (a dictionary identified by `tag` with `someWords`), an array of words to ignore.
	SetIgnoredWordsInSpellDocumentWithTag(words []string, tag int)
	// Sets the string that appears in the misspelled word field, using the string object `aString`.
	SetWordFieldStringValue(string_ string)
	// Causes the spell checker to update the Spelling panel’s misspelled-word field to reflect `word`.
	UpdateSpellingPanelWithMisspelledWord(word string)
	// Provides a list of complete words that the user might be trying to type based on a partial word in a given string.
	CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.NSRange, string_ string, language string, tag int) []string
	// Indicates whether the spell checker has learned a given word.
	HasLearnedWord(word string) bool
	// Tells the spell checker to unlearn a given word.
	UnlearnWord(word string)
	// Adds the word to the spell checker dictionary.
	LearnWord(word string)
	// Returns the default values for quote replacement.
	UserQuotesArrayForLanguage(language string) []string
	// Returns the dictionary used when replacing words.
	UserReplacementsDictionary() foundation.INSDictionary

	// Topic: Data Detector Interaction

	// Provides a menu containing contextual menu items suitable for certain kinds of detected results.
	MenuForResultStringOptionsAtLocationInView(result foundation.NSTextCheckingResult, checkedString string, options foundation.INSDictionary, location corefoundation.CGPoint, view INSView) INSMenu

	// Topic: Automatic Spelling Correction

	// Returns a single proposed correction if a word is mis-spelled.
	CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.NSRange, string_ string, language string, tag int) string
	// Records the user response to the correction indicator being displayed.
	RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response NSCorrectionResponse, correction string, word string, language string, tag int)
	// Dismisses the correction indicator for the specified view.
	DismissCorrectionIndicatorForView(view INSView)

	// Topic: Instance Methods

	DeletesAutospaceBetweenStringAndStringLanguage(precedingString string, followingString string, language string) bool
	LanguageForWordRangeInStringOrthography(range_ foundation.NSRange, string_ string, orthography foundation.NSOrthography) string
	PreventsAutocorrectionBeforeStringLanguage(string_ string, language string) bool
	RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.NSRange, stringToCheck string, checkingTypes uint64, options foundation.INSDictionary, tag int, completionHandler ArrayHandler) int
	ShowInlinePredictionForCandidatesClient(candidates []foundation.NSTextCheckingResult, client NSTextInputClient)
}





// Init initializes the instance.
func (s NSSpellChecker) Init() NSSpellChecker {
	rv := objc.Send[NSSpellChecker](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSpellChecker) Autorelease() NSSpellChecker {
	rv := objc.Send[NSSpellChecker](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSpellChecker creates a new NSSpellChecker instance.
func NewNSSpellChecker() NSSpellChecker {
	class := getNSSpellCheckerClass()
	rv := objc.Send[NSSpellChecker](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the current language used in spell checking.
//
// # Return Value
// 
// The current spell checking language, as a string.
//
// # Discussion
// 
// The result string specifies the language using the language and regional
// designations described in Language and Locale Designations in
// [Internationalization and Localization Guide].
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/language()
func (s NSSpellChecker) Language() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}

// Returns whether the specified language is in the Spelling pop-up list.
//
// language: The requested language.
//
// # Return Value
// 
// [true] if the language is available in the pop-up list, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The code listing below shows how languages can be specified in language. If
// the language specified is listed in the user’s list of preferred
// languages, the spell checker uses that language to accomplish its task.
// 
// Listing 1. Specifying the spell checker language
// 
// To learn about the strings you can use to specify a language in `language`,
// see Language and Locale Designations in [Internationalization and
// Localization Guide].
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setLanguage(_:)
func (s NSSpellChecker) SetLanguage(language string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setLanguage:"), objc.String(language))
	return rv
}

// Specifies a grammar-analysis detail to highlight in the Spelling panel.
//
// string: Problematic grammatical unit identified by
// [CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails].
//
// detail: One of the grammar-analysis details provided by
// [CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails].
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updateSpellingPanel(withGrammarString:detail:)
func (s NSSpellChecker) UpdateSpellingPanelWithGrammarStringDetail(string_ string, detail foundation.INSDictionary) {
	objc.Send[objc.ID](s.ID, objc.Sel("updateSpellingPanelWithGrammarString:detail:"), objc.String(string_), detail)
}

// Updates the available panels to account for user changes.
//
// # Discussion
// 
// This method should be called when a client changes some relevant setting,
// such as what kind of spelling, grammar checking, or substitutions it uses.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updatePanels()
func (s NSSpellChecker) UpdatePanels() {
	objc.Send[objc.ID](s.ID, objc.Sel("updatePanels"))
}

// Returns the number of words in the specified string.
//
// stringToCount: The string to count the words in.
//
// language: The language of the string.
//
// # Return Value
// 
// The number of words in the string or `-1` if word counting is unavailable
// or has not occurred for some reason.
//
// # Discussion
// 
// If `language` is `nil`, the current selection in the Spelling panel’s
// pop-up menu is used.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/countWords(in:language:)
func (s NSSpellChecker) CountWordsInStringLanguage(stringToCount string, language string) int {
	rv := objc.Send[int](s.ID, objc.Sel("countWordsInString:language:"), objc.String(stringToCount), objc.String(language))
	return rv
}

// Starts the search for a misspelled word in `stringToCheck` starting at
// `startingOffset` within the string object.
//
// stringToCheck: The string to spell check.
//
// startingOffset: The offset at which to start checking.
//
// # Return Value
// 
// Returns the range of the first misspelled word.
//
// # Discussion
// 
// Wrapping occurs, but no ignored-words dictionary is used.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:)
func (s NSSpellChecker) CheckSpellingOfStringStartingAt(stringToCheck string, startingOffset int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](s.ID, objc.Sel("checkSpellingOfString:startingAt:"), objc.String(stringToCheck), startingOffset)
	return foundation.NSRange(rv)
}

// Starts the search for a misspelled word in a string starting at specified
// offset within the string.
//
// stringToCheck: The string object containing the words to spellcheck.
//
// startingOffset: The offset within `stringToCheck` at which to begin spellchecking.
//
// language: The language of the words in the string. If `language` is `nil`, or if you
// obtain the value by sending [Language] to `self`, the current selection in
// the Spelling panel’s pop-up menu is used. Do not pass in an empty string
// for `language`.
//
// wrapFlag: [true] to indicate that spell checking should continue at the beginning of
// the string when the end of the string is reached; [false] to indicate that
// spellchecking should stop at the end of the document.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// wordCount: Returns by indirection a count of the words spell-checked up to and
// including the first error (if any), or -1 if the spell checker fails or
// does not support word counting. Specify [NULL] if you do not want this word
// count.
//
// # Return Value
// 
// The range of the first misspelled word and optionally (and by reference)
// the count of words spellchecked in the string in `wordCount`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkSpelling(of:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:)
func (s NSSpellChecker) CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, wordCount int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](s.ID, objc.Sel("checkSpellingOfString:startingAt:language:wrap:inSpellDocumentWithTag:wordCount:"), objc.String(stringToCheck), startingOffset, objc.String(language), wrapFlag, tag, wordCount)
	return foundation.NSRange(rv)
}

// Initiates a grammatical analysis of a given string.
//
// stringToCheck: String to analyze.
//
// startingOffset: Location within `string` at which to start the analysis.
//
// language: Language use in `string`. When `nil`, the language selected in the Spelling
// panel is used.
//
// wrapFlag: [true] to specify that the analysis continue to the beginning of string
// when the end is reached.
// 
// [false] to have the analysis stop at the end of string.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// details: On output, dictionaries describing grammar-analysis details within the
// flagged grammatical unit. See the [NSSpellServer] class for information
// about these dictionaries.
// //
// [NSSpellServer]: https://developer.apple.com/documentation/Foundation/NSSpellServer
//
// # Return Value
// 
// Location of the first flagged grammatical unit.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/checkGrammar(of:startingAt:language:wrap:inSpellDocumentWithTag:details:)
func (s NSSpellChecker) CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails(stringToCheck string, startingOffset int, language string, wrapFlag bool, tag int, details foundation.INSDictionary) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](s.ID, objc.Sel("checkGrammarOfString:startingAt:language:wrap:inSpellDocumentWithTag:details:"), objc.String(stringToCheck), startingOffset, objc.String(language), wrapFlag, tag, details)
	return foundation.NSRange(rv)
}

// Requests unified text checking for the given range of the given string.
//
// stringToCheck: The string to check.
//
// range: The range of the string to check.
//
// checkingTypes: The type of checking to be performed. The possible constants are listed in
// [NSTextCheckingResult.CheckingType] and can be combined using the C
// bit-wise [OR] operator to perform multiple checks at the same time.
// //
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
//
// options: The options dictionary specifying the types of checking to perform. See
// `Spell Checking Option Dictionary Keys` for the possible keys and expected
// values.
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// orthography: Returns by-reference, the orthography of the range of the string. See
// [NSOrthography] for more information.
// //
// [NSOrthography]: https://developer.apple.com/documentation/Foundation/NSOrthography
//
// wordCount: Returns by-reference, the word count for the range of the string.
//
// # Return Value
// 
// An array of [NSTextCheckingResult] objects describing particular items
// found during checking and their individual ranges, sorted by range origin,
// then range end, then result type.
//
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/check(_:range:types:options:inSpellDocumentWithTag:orthography:wordCount:)
func (s NSSpellChecker) CheckStringRangeTypesOptionsInSpellDocumentWithTagOrthographyWordCount(stringToCheck string, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, tag int, orthography foundation.NSOrthography, wordCount int) []foundation.NSTextCheckingResult {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("checkString:range:types:options:inSpellDocumentWithTag:orthography:wordCount:"), objc.String(stringToCheck), range_, checkingTypes, options, tag, orthography, wordCount)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSTextCheckingResult {
		return foundation.NSTextCheckingResultFromID(id)
	})
}

// Requests that the string be checked in the background.
//
// stringToCheck: The string to check.
//
// range: The range of the string to check.
//
// checkingTypes: The type of checking to be performed. The possible constants are listed in
// [NSTextCheckingResult.CheckingType] and can be combined using the C
// bit-wise [OR] operator to perform multiple checks at the same time.
// //
// [NSTextCheckingResult.CheckingType]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
//
// options: The options dictionary specifying the types of checking to perform. See
// [NSTextCheckingOptionKey] for the possible keys and expected values.
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// completionHandler: The completion handler block object will be called (in an arbitrary
// context) when results are available, with the sequence number and results.
// 
// The block takes four arguments:
// 
// sequenceNumber: A monotonically increasing sequence number. results: An
// array of [NSTextCheckingResult] objects describing particular items found
// during checking and their individual ranges, sorted by range origin, then
// range end, then result type. orthography: The orthography of the string.
// wordCount: The number of words in the range of the string.
// //
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// # Return Value
// 
// The return value is a monotonically increasing sequence number that can be
// used to keep track of requests in flight.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestChecking(of:range:types:options:inSpellDocumentWithTag:completionHandler:)
func (s NSSpellChecker) RequestCheckingOfStringRangeTypesOptionsInSpellDocumentWithTagCompletionHandler(stringToCheck string, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, tag int, completionHandler ErrorHandler) int {
		_block5, _cleanup5 := NewErrorBlock(completionHandler)
	defer _cleanup5()
		rv := objc.Send[int](s.ID, objc.Sel("requestCheckingOfString:range:types:options:inSpellDocumentWithTag:completionHandler:"), objc.String(stringToCheck), range_, checkingTypes, options, tag, _block5)
	return rv
}

// Returns an array of possible substitutions for the specified string.
//
// range: The range of the string to check.
//
// string: The string to guess.
//
// language: The language of the string.
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// # Return Value
// 
// An array of strings containing possible replacement words.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/guesses(forWordRange:in:language:inSpellDocumentWithTag:)
func (s NSSpellChecker) GuessesForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.NSRange, string_ string, language string, tag int) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("guessesForWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return objc.ConvertSliceToStrings(rv)
}

// Notifies the receiver that the user has finished with the tagged document.
//
// # Discussion
// 
// The spell checker will release any resources associated with the document,
// including but not necessarily limited to, ignored words.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/closeSpellDocument(withTag:)
func (s NSSpellChecker) CloseSpellDocumentWithTag(tag int) {
	objc.Send[objc.ID](s.ID, objc.Sel("closeSpellDocumentWithTag:"), tag)
}

// Instructs the spell checker to ignore all future occurrences of
// `wordToIgnore` in the document identified by `tag`.
//
// # Discussion
// 
// You should invoke this method from within your implementation of the
// NSIgnoreMisspelledWords protocol’s [IgnoreSpelling] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/ignoreWord(_:inSpellDocumentWithTag:)
func (s NSSpellChecker) IgnoreWordInSpellDocumentWithTag(wordToIgnore string, tag int) {
	objc.Send[objc.ID](s.ID, objc.Sel("ignoreWord:inSpellDocumentWithTag:"), objc.String(wordToIgnore), tag)
}

// Returns the array of ignored words for a document identified by `tag`.
//
// # Discussion
// 
// Invoke this method before [CloseSpellDocumentWithTag] if you want to store
// the ignored words.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/ignoredWords(inSpellDocumentWithTag:)
func (s NSSpellChecker) IgnoredWordsInSpellDocumentWithTag(tag int) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("ignoredWordsInSpellDocumentWithTag:"), tag)
	return objc.ConvertSliceToStrings(rv)
}

// Initializes the ignored-words document (a dictionary identified by `tag`
// with `someWords`), an array of words to ignore.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setIgnoredWords(_:inSpellDocumentWithTag:)
func (s NSSpellChecker) SetIgnoredWordsInSpellDocumentWithTag(words []string, tag int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setIgnoredWords:inSpellDocumentWithTag:"), objectivec.StringSliceToNSArray(words), tag)
}

// Sets the string that appears in the misspelled word field, using the string
// object `aString`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/setWordFieldStringValue(_:)
func (s NSSpellChecker) SetWordFieldStringValue(string_ string) {
	objc.Send[objc.ID](s.ID, objc.Sel("setWordFieldStringValue:"), objc.String(string_))
}

// Causes the spell checker to update the Spelling panel’s misspelled-word
// field to reflect `word`.
//
// # Discussion
// 
// You are responsible for highlighting `word` in the document and for
// extracting it from the document using the range returned by the `...`
// methods. Pass the empty string as `word` to have the system beep,
// indicating no misspelled words were found.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/updateSpellingPanel(withMisspelledWord:)
func (s NSSpellChecker) UpdateSpellingPanelWithMisspelledWord(word string) {
	objc.Send[objc.ID](s.ID, objc.Sel("updateSpellingPanelWithMisspelledWord:"), objc.String(word))
}

// Provides a list of complete words that the user might be trying to type
// based on a partial word in a given string.
//
// range: Range that identifies a partial word in `string`.
//
// string: String with the partial word from which to generate the result.
//
// language: Language to used in `string`. When `nil`, this method uses the language
// selected in the Spelling panel.
//
// tag: Identifies the spell document with ignored words to use.
//
// # Return Value
// 
// List of complete words from the spell checker dictionary in the order they
// should be presented to the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/completions(forPartialWordRange:in:language:inSpellDocumentWithTag:)
func (s NSSpellChecker) CompletionsForPartialWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.NSRange, string_ string, language string, tag int) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("completionsForPartialWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return objc.ConvertSliceToStrings(rv)
}

// Indicates whether the spell checker has learned a given word.
//
// word: Word in question.
//
// # Return Value
// 
// [true] when the spell checker has learned word, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/hasLearnedWord(_:)
func (s NSSpellChecker) HasLearnedWord(word string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasLearnedWord:"), objc.String(word))
	return rv
}

// Tells the spell checker to unlearn a given word.
//
// word: Word to unlearn.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/unlearnWord(_:)
func (s NSSpellChecker) UnlearnWord(word string) {
	objc.Send[objc.ID](s.ID, objc.Sel("unlearnWord:"), objc.String(word))
}

// Adds the word to the spell checker dictionary.
//
// word: The word to add.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/learnWord(_:)
func (s NSSpellChecker) LearnWord(word string) {
	objc.Send[objc.ID](s.ID, objc.Sel("learnWord:"), objc.String(word))
}

// Returns the default values for quote replacement.
//
// language: The language for quote replacement.
//
// # Return Value
// 
// An array of quote replacements used by the [quotes] key-value pair.
//
// [quotes]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/quotes
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userQuotesArray(forLanguage:)
func (s NSSpellChecker) UserQuotesArrayForLanguage(language string) []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("userQuotesArrayForLanguage:"), objc.String(language))
	return objc.ConvertSliceToStrings(rv)
}

// Provides a menu containing contextual menu items suitable for certain kinds
// of detected results.
//
// result: The [NSTextCheckingResult] instance for the checked string.
// //
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// checkedString: The string that has been checked.
//
// options: The options dictionary allows clients to pass in information associated
// with the document. See `Spell Checking Option Dictionary Keys` for possible
// key-value pairs.
//
// location: The location, in the view’s coordinate system, to display the menu.
//
// view: The view object over which to display the contextual menu.
//
// # Return Value
// 
// A menu suitable for displaying as a contextual menu, or adding to another
// contextual menu as a submenu.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/menu(for:string:options:atLocation:in:)
func (s NSSpellChecker) MenuForResultStringOptionsAtLocationInView(result foundation.NSTextCheckingResult, checkedString string, options foundation.INSDictionary, location corefoundation.CGPoint, view INSView) INSMenu {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("menuForResult:string:options:atLocation:inView:"), result, objc.String(checkedString), options, location, view)
	return NSMenuFromID(rv)
}

// Returns a single proposed correction if a word is mis-spelled.
//
// range: The range of the word to be corrected.
//
// string: The string containing the proposed correction.
//
// language: The language.
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// # Return Value
// 
// The proposed correct string.
//
// # Discussion
// 
// While correction functionality is available starting in OS X v10.6 as part
// of unified text checking, for convenience this method makes it available
// separately starting in OS X v10.7.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/correction(forWordRange:in:language:inSpellDocumentWithTag:)
func (s NSSpellChecker) CorrectionForWordRangeInStringLanguageInSpellDocumentWithTag(range_ foundation.NSRange, string_ string, language string, tag int) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("correctionForWordRange:inString:language:inSpellDocumentWithTag:"), range_, objc.String(string_), objc.String(language), tag)
	return foundation.NSStringFromID(rv).String()
}

// Records the user response to the correction indicator being displayed.
//
// response: The user’s response. The possible values are shown in
// [NSSpellChecker.CorrectionResponse].
// //
// [NSSpellChecker.CorrectionResponse]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse
//
// correction: The corrected word. This should match the original correction.
//
// word: The original word. This should match the original correction.
//
// language: The language being edited. This should match the original correction.
//
// tag: An identifier unique within the application used to inform the spell
// checker which document that text is associated, potentially for many
// purposes, not necessarily just for ignored words. A value of 0 can be
// passed in for text not associated with a particular document.
//
// # Discussion
// 
// When a correction is automatically proposed, the user may respond in one of
// several ways. Clients may report this to the spell checker so that it can
// learn from the user’s response and adjust future correction behavior
// accordingly.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/record(_:toCorrection:forWord:language:inSpellDocumentWithTag:)
func (s NSSpellChecker) RecordResponseToCorrectionForWordLanguageInSpellDocumentWithTag(response NSCorrectionResponse, correction string, word string, language string, tag int) {
	objc.Send[objc.ID](s.ID, objc.Sel("recordResponse:toCorrection:forWord:language:inSpellDocumentWithTag:"), response, objc.String(correction), objc.String(word), objc.String(language), tag)
}

// Dismisses the correction indicator for the specified view.
//
// view: The view.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/dismissCorrectionIndicator(for:)
func (s NSSpellChecker) DismissCorrectionIndicatorForView(view INSView) {
	objc.Send[objc.ID](s.ID, objc.Sel("dismissCorrectionIndicatorForView:"), view)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/deletesAutospaceBetweenString(_:andString:language:)
func (s NSSpellChecker) DeletesAutospaceBetweenStringAndStringLanguage(precedingString string, followingString string, language string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("deletesAutospaceBetweenString:andString:language:"), objc.String(precedingString), objc.String(followingString), objc.String(language))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/language(forWordRange:in:orthography:)
func (s NSSpellChecker) LanguageForWordRangeInStringOrthography(range_ foundation.NSRange, string_ string, orthography foundation.NSOrthography) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageForWordRange:inString:orthography:"), range_, objc.String(string_), orthography)
	return foundation.NSStringFromID(rv).String()
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/preventsAutocorrection(before:language:)
func (s NSSpellChecker) PreventsAutocorrectionBeforeStringLanguage(string_ string, language string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("preventsAutocorrectionBeforeString:language:"), objc.String(string_), objc.String(language))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/requestCandidates(forSelectedRange:in:types:options:inSpellDocumentWithTag:completionHandler:)
func (s NSSpellChecker) RequestCandidatesForSelectedRangeInStringTypesOptionsInSpellDocumentWithTagCompletionHandler(selectedRange foundation.NSRange, stringToCheck string, checkingTypes uint64, options foundation.INSDictionary, tag int, completionHandler ArrayHandler) int {
		_block5, _cleanup5 := NewArrayBlock(completionHandler)
	defer _cleanup5()
		rv := objc.Send[int](s.ID, objc.Sel("requestCandidatesForSelectedRange:inString:types:options:inSpellDocumentWithTag:completionHandler:"), selectedRange, objc.String(stringToCheck), checkingTypes, options, tag, _block5)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/showInlinePrediction(forCandidates:client:)
func (s NSSpellChecker) ShowInlinePredictionForCandidatesClient(candidates []foundation.NSTextCheckingResult, client NSTextInputClient) {
	objc.Send[objc.ID](s.ID, objc.Sel("showInlinePredictionForCandidates:client:"), objectivec.IObjectSliceToNSArray(candidates), client)
}





// Returns a guaranteed unique tag to use as the spell-document tag for a
// document.
//
// # Return Value
// 
// Returns a unique tag to identified this spell checked object.
//
// # Discussion
// 
// Use this method to generate tags to avoid collisions with other objects
// that can be spell checked.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/uniqueSpellDocumentTag()
func (_NSSpellCheckerClass NSSpellCheckerClass) UniqueSpellDocumentTag() int {
	rv := objc.Send[int](objc.ID(_NSSpellCheckerClass.class), objc.Sel("uniqueSpellDocumentTag"))
	return rv
}








// Provides a list of all available languages.
//
// # Return Value
// 
// An array containing all the available spell checking languages. The
// languages are ordered in the user’s preferred order as set in the system
// preferences.
// 
// # Discussion
// 
// If [AutomaticallyIdentifiesLanguages] is [true], then text checking will
// automatically use this method as appropriate; otherwise, it will use the
// language set by [SetLanguage].
// 
// The older
// [CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount]
// and
// [CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails].
// methods will use the language set by [SetLanguage], if they are called with
// a `nil` language argument.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/availableLanguages
func (s NSSpellChecker) AvailableLanguages() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("availableLanguages"))
	return objc.ConvertSliceToStrings(rv)
}



// Provides a subset of the available languages to be used for spell checking.
//
// # Return Value
// 
// An array containing the user’s preferred languages for spell checking.
// The order is set in the system preferences.
// 
// # Discussion
// 
// If [AutomaticallyIdentifiesLanguages] is [true], then text checking will
// automatically use this method as appropriate; otherwise, it will use the
// language set by [SetLanguage].
// 
// The older
// [CheckSpellingOfStringStartingAtLanguageWrapInSpellDocumentWithTagWordCount]
// and
// [CheckGrammarOfStringStartingAtLanguageWrapInSpellDocumentWithTagDetails].
// methods will use the language set by [SetLanguage], if they are called with
// a `nil` language argument.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userPreferredLanguages
func (s NSSpellChecker) UserPreferredLanguages() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("userPreferredLanguages"))
	return objc.ConvertSliceToStrings(rv)
}



// Sets whether the spell checker will automatically identify languages.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/automaticallyIdentifiesLanguages
func (s NSSpellChecker) AutomaticallyIdentifiesLanguages() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("automaticallyIdentifiesLanguages"))
	return rv
}
func (s NSSpellChecker) SetAutomaticallyIdentifiesLanguages(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutomaticallyIdentifiesLanguages:"), value)
}



// Returns the spell checker’s panel.
//
// # Return Value
// 
// The spell checking panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/spellingPanel
func (s NSSpellChecker) SpellingPanel() INSPanel {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spellingPanel"))
	return NSPanelFromID(objc.ID(rv))
}



// Returns the substitutions panel.
//
// # Return Value
// 
// The substitutions checking panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/substitutionsPanel
func (s NSSpellChecker) SubstitutionsPanel() INSPanel {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("substitutionsPanel"))
	return NSPanelFromID(objc.ID(rv))
}



// Makes a view an accessory of the Spelling panel by making it a subview of
// the panel’s content view.
//
// # Discussion
// 
// The accessory view can be any custom view you want to display with the
// spelling panel. The accessory view is displayed below the spelling checker
// and the panel automatically resizes to accommodate the accessory view.
// 
// This method posts a notification named [didResizeNotification] with the
// Spelling panel object to the default notification center.
//
// [didResizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResizeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/accessoryView
func (s NSSpellChecker) AccessoryView() INSView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (s NSSpellChecker) SetAccessoryView(value INSView) {
	objc.Send[struct{}](s.ID, objc.Sel("setAccessoryView:"), value)
}



// Sets the substitutions panel’s accessory view.
//
// # Discussion
// 
// The accessory view controller can accommodate any custom view you want to
// display with the substitutions panel. The accessory view controller’s
// view is displayed below the substitutions list and the panel automatically
// resizes to accommodate the accessory view.
// 
// This method posts a notification named [didResizeNotification] with the
// substitutions panel object to the default notification center.
//
// [didResizeNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/didResizeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/substitutionsPanelAccessoryViewController
func (s NSSpellChecker) SubstitutionsPanelAccessoryViewController() INSViewController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("substitutionsPanelAccessoryViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (s NSSpellChecker) SetSubstitutionsPanelAccessoryViewController(value INSViewController) {
	objc.Send[struct{}](s.ID, objc.Sel("setSubstitutionsPanelAccessoryViewController:"), value)
}



// Returns the dictionary used when replacing words.
//
// # Return Value
// 
// The dictionary.
// 
// # Discussion
// 
// The key-value pairs in this dictionary are used by the [quotes] when
// replacing characters and words.
//
// [quotes]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey/quotes
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/userReplacementsDictionary
func (s NSSpellChecker) UserReplacementsDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("userReplacementsDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}







// Returns the NSSpellChecker (one per application).
//
// # Return Value
// 
// The spelling checker shared by this application.
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/shared
func (_NSSpellCheckerClass NSSpellCheckerClass) SharedSpellChecker() NSSpellChecker {
	rv := objc.Send[objc.ID](objc.ID(_NSSpellCheckerClass.class), objc.Sel("sharedSpellChecker"))
	return NSSpellCheckerFromID(objc.ID(rv))
}



// Returns whether the application’s NSSpellChecker has already been
// created.
//
// # Return Value
// 
// [true] if the shared spell checker already exists, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/sharedSpellCheckerExists
func (_NSSpellCheckerClass NSSpellCheckerClass) SharedSpellCheckerExists() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("sharedSpellCheckerExists"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticCapitalizationEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticCapitalizationEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticCapitalizationEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticDashSubstitutionEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticInlinePredictionEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticInlinePredictionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticInlinePredictionEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticPeriodSubstitutionEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticPeriodSubstitutionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticPeriodSubstitutionEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticQuoteSubstitutionEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticSpellingCorrectionEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextCompletionEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}



// See: https://developer.apple.com/documentation/AppKit/NSSpellChecker/isAutomaticTextReplacementEnabled
func (_NSSpellCheckerClass NSSpellCheckerClass) AutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](objc.ID(_NSSpellCheckerClass.class), objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}



















