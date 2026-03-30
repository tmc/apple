// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAttributedStringMarkdownParsingOptions] class.
var (
	_NSAttributedStringMarkdownParsingOptionsClass     NSAttributedStringMarkdownParsingOptionsClass
	_NSAttributedStringMarkdownParsingOptionsClassOnce sync.Once
)

func getNSAttributedStringMarkdownParsingOptionsClass() NSAttributedStringMarkdownParsingOptionsClass {
	_NSAttributedStringMarkdownParsingOptionsClassOnce.Do(func() {
		_NSAttributedStringMarkdownParsingOptionsClass = NSAttributedStringMarkdownParsingOptionsClass{class: objc.GetClass("NSAttributedStringMarkdownParsingOptions")}
	})
	return _NSAttributedStringMarkdownParsingOptionsClass
}

// GetNSAttributedStringMarkdownParsingOptionsClass returns the class object for NSAttributedStringMarkdownParsingOptions.
func GetNSAttributedStringMarkdownParsingOptionsClass() NSAttributedStringMarkdownParsingOptionsClass {
	return getNSAttributedStringMarkdownParsingOptionsClass()
}

type NSAttributedStringMarkdownParsingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAttributedStringMarkdownParsingOptionsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAttributedStringMarkdownParsingOptionsClass) Alloc() NSAttributedStringMarkdownParsingOptions {
	rv := objc.Send[NSAttributedStringMarkdownParsingOptions](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Options that affect the parsing of Markdown content into an attributed
// string.
//
// # Determining Markdown Parsing Options
//
//   - [NSAttributedStringMarkdownParsingOptions.AllowsExtendedAttributes]: A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes.
//   - [NSAttributedStringMarkdownParsingOptions.SetAllowsExtendedAttributes]
//   - [NSAttributedStringMarkdownParsingOptions.AppliesSourcePositionAttributes]: A Boolean value that indicates whether parsing applies attributes that indicate the position of attributed text in the original Markdown string.
//   - [NSAttributedStringMarkdownParsingOptions.SetAppliesSourcePositionAttributes]
//   - [NSAttributedStringMarkdownParsingOptions.FailurePolicy]: The policy for handling a parsing failure.
//   - [NSAttributedStringMarkdownParsingOptions.SetFailurePolicy]
//   - [NSAttributedStringMarkdownParsingOptions.InterpretedSyntax]: The syntax for intepreting a Markdown string.
//   - [NSAttributedStringMarkdownParsingOptions.SetInterpretedSyntax]
//   - [NSAttributedStringMarkdownParsingOptions.LanguageCode]: The BCP-47 language code for this document.
//   - [NSAttributedStringMarkdownParsingOptions.SetLanguageCode]
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions
type NSAttributedStringMarkdownParsingOptions struct {
	objectivec.Object
}

// NSAttributedStringMarkdownParsingOptionsFromID constructs a [NSAttributedStringMarkdownParsingOptions] from an objc.ID.
//
// Options that affect the parsing of Markdown content into an attributed
// string.
func NSAttributedStringMarkdownParsingOptionsFromID(id objc.ID) NSAttributedStringMarkdownParsingOptions {
	return NSAttributedStringMarkdownParsingOptions{objectivec.Object{ID: id}}
}

// NOTE: NSAttributedStringMarkdownParsingOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAttributedStringMarkdownParsingOptions] class.
//
// # Determining Markdown Parsing Options
//
//   - [INSAttributedStringMarkdownParsingOptions.AllowsExtendedAttributes]: A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes.
//   - [INSAttributedStringMarkdownParsingOptions.SetAllowsExtendedAttributes]
//   - [INSAttributedStringMarkdownParsingOptions.AppliesSourcePositionAttributes]: A Boolean value that indicates whether parsing applies attributes that indicate the position of attributed text in the original Markdown string.
//   - [INSAttributedStringMarkdownParsingOptions.SetAppliesSourcePositionAttributes]
//   - [INSAttributedStringMarkdownParsingOptions.FailurePolicy]: The policy for handling a parsing failure.
//   - [INSAttributedStringMarkdownParsingOptions.SetFailurePolicy]
//   - [INSAttributedStringMarkdownParsingOptions.InterpretedSyntax]: The syntax for intepreting a Markdown string.
//   - [INSAttributedStringMarkdownParsingOptions.SetInterpretedSyntax]
//   - [INSAttributedStringMarkdownParsingOptions.LanguageCode]: The BCP-47 language code for this document.
//   - [INSAttributedStringMarkdownParsingOptions.SetLanguageCode]
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions
type INSAttributedStringMarkdownParsingOptions interface {
	objectivec.IObject
	NSCopying

	// Topic: Determining Markdown Parsing Options

	// A Boolean value that indicates whether parsing allows extensions to Markdown that specify extended attributes.
	AllowsExtendedAttributes() bool
	SetAllowsExtendedAttributes(value bool)
	// A Boolean value that indicates whether parsing applies attributes that indicate the position of attributed text in the original Markdown string.
	AppliesSourcePositionAttributes() bool
	SetAppliesSourcePositionAttributes(value bool)
	// The policy for handling a parsing failure.
	FailurePolicy() NSAttributedStringMarkdownParsingFailurePolicy
	SetFailurePolicy(value NSAttributedStringMarkdownParsingFailurePolicy)
	// The syntax for intepreting a Markdown string.
	InterpretedSyntax() NSAttributedStringMarkdownInterpretedSyntax
	SetInterpretedSyntax(value NSAttributedStringMarkdownInterpretedSyntax)
	// The BCP-47 language code for this document.
	LanguageCode() string
	SetLanguageCode(value string)
}

// Init initializes the instance.
func (a NSAttributedStringMarkdownParsingOptions) Init() NSAttributedStringMarkdownParsingOptions {
	rv := objc.Send[NSAttributedStringMarkdownParsingOptions](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAttributedStringMarkdownParsingOptions) Autorelease() NSAttributedStringMarkdownParsingOptions {
	rv := objc.Send[NSAttributedStringMarkdownParsingOptions](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAttributedStringMarkdownParsingOptions creates a new NSAttributedStringMarkdownParsingOptions instance.
func NewNSAttributedStringMarkdownParsingOptions() NSAttributedStringMarkdownParsingOptions {
	class := getNSAttributedStringMarkdownParsingOptionsClass()
	rv := objc.Send[NSAttributedStringMarkdownParsingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether parsing allows extensions to
// Markdown that specify extended attributes.
//
// # Discussion
//
// If this value is [NO], the Markdown parser supports only the CommonMark
// syntax. The default is [NO].
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/allowsExtendedAttributes
func (a NSAttributedStringMarkdownParsingOptions) AllowsExtendedAttributes() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowsExtendedAttributes"))
	return rv
}
func (a NSAttributedStringMarkdownParsingOptions) SetAllowsExtendedAttributes(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAllowsExtendedAttributes:"), value)
}

// A Boolean value that indicates whether parsing applies attributes that
// indicate the position of attributed text in the original Markdown string.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/appliesSourcePositionAttributes
func (a NSAttributedStringMarkdownParsingOptions) AppliesSourcePositionAttributes() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("appliesSourcePositionAttributes"))
	return rv
}
func (a NSAttributedStringMarkdownParsingOptions) SetAppliesSourcePositionAttributes(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAppliesSourcePositionAttributes:"), value)
}

// The policy for handling a parsing failure.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/failurePolicy
func (a NSAttributedStringMarkdownParsingOptions) FailurePolicy() NSAttributedStringMarkdownParsingFailurePolicy {
	rv := objc.Send[NSAttributedStringMarkdownParsingFailurePolicy](a.ID, objc.Sel("failurePolicy"))
	return NSAttributedStringMarkdownParsingFailurePolicy(rv)
}
func (a NSAttributedStringMarkdownParsingOptions) SetFailurePolicy(value NSAttributedStringMarkdownParsingFailurePolicy) {
	objc.Send[struct{}](a.ID, objc.Sel("setFailurePolicy:"), value)
}

// The syntax for intepreting a Markdown string.
//
// # Discussion
//
// If your Markdown data uses syntax that this setting excludes, the parser
// still parses it and includes its text in the final result. However, the
// relevant text won’t have attributes.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/interpretedSyntax
func (a NSAttributedStringMarkdownParsingOptions) InterpretedSyntax() NSAttributedStringMarkdownInterpretedSyntax {
	rv := objc.Send[NSAttributedStringMarkdownInterpretedSyntax](a.ID, objc.Sel("interpretedSyntax"))
	return NSAttributedStringMarkdownInterpretedSyntax(rv)
}
func (a NSAttributedStringMarkdownParsingOptions) SetInterpretedSyntax(value NSAttributedStringMarkdownInterpretedSyntax) {
	objc.Send[struct{}](a.ID, objc.Sel("setInterpretedSyntax:"), value)
}

// The BCP-47 language code for this document.
//
// # Discussion
//
// If not `nil`, the string applies the [languageIdentifier] attribute to any
// range in the returned string that doesn’t otherwise specify a language
// attribute. The default is `nil`, which applies no attributes.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingOptions/languageCode
//
// [languageIdentifier]: https://developer.apple.com/documentation/Foundation/AttributeScopes/FoundationAttributes/languageIdentifier
func (a NSAttributedStringMarkdownParsingOptions) LanguageCode() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("languageCode"))
	return NSStringFromID(rv).String()
}
func (a NSAttributedStringMarkdownParsingOptions) SetLanguageCode(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}

// Protocol methods for NSCopying
