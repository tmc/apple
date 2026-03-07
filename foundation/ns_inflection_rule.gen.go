// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSInflectionRule] class.
var (
	_NSInflectionRuleClass     NSInflectionRuleClass
	_NSInflectionRuleClassOnce sync.Once
)

func getNSInflectionRuleClass() NSInflectionRuleClass {
	_NSInflectionRuleClassOnce.Do(func() {
		_NSInflectionRuleClass = NSInflectionRuleClass{class: objc.GetClass("NSInflectionRule")}
	})
	return _NSInflectionRuleClass
}

// GetNSInflectionRuleClass returns the class object for NSInflectionRule.
func GetNSInflectionRuleClass() NSInflectionRuleClass {
	return getNSInflectionRuleClass()
}

type NSInflectionRuleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSInflectionRuleClass) Alloc() NSInflectionRule {
	rv := objc.Send[NSInflectionRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A rule that affects how an attributed string performs automatic grammatical
// agreement.
//
// # Overview
// 
// Most apps can rely on loading localized strings to perform automatic
// grammar agreement. Typically, your app’s strings files use the Markdown
// extension syntax to indicate portions of the string that may require
// inflection to agree grammatically. This transformation occurs when you load
// the attributed string with methods like [NSLocalizedAttributedString].
// 
// However, if the system lacks information about the words in the string, you
// may need to apply an inflection rule programmatically. For example, a
// social networking app may have gender information about other users that
// you want to apply at runtime. When performing manual inflection at runtime,
// you use an inflection rule to indicate to the system what portions of a
// string should be automatically edited, and what to match. Add the attribute
// [inflectionRule] with an [NSInflectionRule] on an [NSAttributedString],
// then call [AttributedStringByInflectingString] to perform the grammar
// agreement and produce an edited string.
//
// [inflectionRule]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/inflectionRule
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRule
type NSInflectionRule struct {
	objectivec.Object
}

// NSInflectionRuleFromID constructs a [NSInflectionRule] from an objc.ID.
//
// A rule that affects how an attributed string performs automatic grammatical
// agreement.
func NSInflectionRuleFromID(id objc.ID) NSInflectionRule {
	return NSInflectionRule{objectivec.Object{id}}
}
// NOTE: NSInflectionRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSInflectionRule] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRule
type INSInflectionRule interface {
	objectivec.IObject
	NSCopying

	InitWithCoder(coder INSCoder) NSInflectionRule
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}





// Init initializes the instance.
func (i NSInflectionRule) Init() NSInflectionRule {
	rv := objc.Send[NSInflectionRule](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSInflectionRule) Autorelease() NSInflectionRule {
	rv := objc.Send[NSInflectionRule](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSInflectionRule creates a new NSInflectionRule instance.
func NewNSInflectionRule() NSInflectionRule {
	class := getNSInflectionRuleClass()
	rv := objc.Send[NSInflectionRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}










//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (i NSInflectionRule) InitWithCoder(coder INSCoder) NSInflectionRule {
	rv := objc.Send[NSInflectionRule](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (i NSInflectionRule) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Returns a Boolean value that indicates whether the rule can inflect a given
// language.
//
// language: The language to apply inflection to, specified as a BCP 47 language code.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRule/canInflectLanguage:
func (_NSInflectionRuleClass NSInflectionRuleClass) CanInflectLanguage(language string) bool {
	rv := objc.Send[bool](objc.ID(_NSInflectionRuleClass.class), objc.Sel("canInflectLanguage:"), objc.String(language))
	return rv
}












// An inflection rule that performs automatic grammar agreement with default
// transformations.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRule/automaticRule
func (_NSInflectionRuleClass NSInflectionRuleClass) AutomaticRule() NSInflectionRule {
	rv := objc.Send[objc.ID](objc.ID(_NSInflectionRuleClass.class), objc.Sel("automaticRule"))
	return NSInflectionRuleFromID(objc.ID(rv))
}



// A Boolean value that indicates whether the rule can inflect the user’s
// current preferred localization.
//
// # Discussion
// 
// This value doesn’t change throughout the lifetime of a process.
//
// See: https://developer.apple.com/documentation/Foundation/NSInflectionRule/canInflectPreferredLocalization
func (_NSInflectionRuleClass NSInflectionRuleClass) CanInflectPreferredLocalization() bool {
	rv := objc.Send[bool](objc.ID(_NSInflectionRuleClass.class), objc.Sel("canInflectPreferredLocalization"))
	return rv
}




			// Protocol methods for NSCopying
			














