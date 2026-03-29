// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTermOfAddress] class.
var (
	_NSTermOfAddressClass     NSTermOfAddressClass
	_NSTermOfAddressClassOnce sync.Once
)

func getNSTermOfAddressClass() NSTermOfAddressClass {
	_NSTermOfAddressClassOnce.Do(func() {
		_NSTermOfAddressClass = NSTermOfAddressClass{class: objc.GetClass("NSTermOfAddress")}
	})
	return _NSTermOfAddressClass
}

// GetNSTermOfAddressClass returns the class object for NSTermOfAddress.
func GetNSTermOfAddressClass() NSTermOfAddressClass {
	return getNSTermOfAddressClass()
}

type NSTermOfAddressClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTermOfAddressClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTermOfAddressClass) Alloc() NSTermOfAddress {
	rv := objc.Send[NSTermOfAddress](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The type for representing grammatical gender in localized text.
//
// # Overview
// 
// Many languages rely on gender for their grammar. Without knowing the
// subject’s gender or pronoun preferences, some localized strings may have
// grammatical errors, resulting in a poor user experience.
// 
// [TermOfAddress] is a type that enables the system to make pronoun
// substitutions in localized text based on gender. You don’t create
// instances of this type directly. Instead, use the predefined types to
// specify the gender to use when referring to people in translated text. Or
// define your own pronoun terms for a specific language when the predefined
// types are insufficient.
// 
// For example, to substitute the masculine pronoun , for the neutral pronoun
// , do the following:
// 
// If the [NSTermOfAddress.Masculine], [NSTermOfAddress.Feminine], and [NSTermOfAddress.Neutral] terms of address are
// insufficient, create your own term of address specifying the pronouns and
// language.
// 
// For examples of how to use terms of address, see:
// 
// - [AttributeScopes.FoundationAttributes.ReferentConceptAttribute] -
// [AttributeScopes.FoundationAttributes.AgreementConceptAttribute]
//
// [AttributeScopes.FoundationAttributes.AgreementConceptAttribute]: https://developer.apple.com/documentation/Foundation/AttributeScopes/FoundationAttributes/AgreementConceptAttribute
// [AttributeScopes.FoundationAttributes.ReferentConceptAttribute]: https://developer.apple.com/documentation/Foundation/AttributeScopes/FoundationAttributes/ReferentConceptAttribute
//
// # Defining your own terms of address
//
//   - [NSTermOfAddress.Pronouns]: A list of pronouns for a localized term of address
//
// # Instance Properties
//
//   - [NSTermOfAddress.LanguageIdentifier]: The ISO language code if this is a localized term of address
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress
type NSTermOfAddress struct {
	objectivec.Object
}

// NSTermOfAddressFromID constructs a [NSTermOfAddress] from an objc.ID.
//
// The type for representing grammatical gender in localized text.
func NSTermOfAddressFromID(id objc.ID) NSTermOfAddress {
	return NSTermOfAddress{objectivec.Object{ID: id}}
}
// NOTE: NSTermOfAddress adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTermOfAddress] class.
//
// # Defining your own terms of address
//
//   - [INSTermOfAddress.Pronouns]: A list of pronouns for a localized term of address
//
// # Instance Properties
//
//   - [INSTermOfAddress.LanguageIdentifier]: The ISO language code if this is a localized term of address
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress
type INSTermOfAddress interface {
	objectivec.IObject
	NSCopying
	NSSecureCoding

	// Topic: Defining your own terms of address

	// A list of pronouns for a localized term of address
	Pronouns() []NSMorphologyPronoun

	// Topic: Instance Properties

	// The ISO language code if this is a localized term of address
	LanguageIdentifier() string

	InitWithCoder(coder INSCoder) NSTermOfAddress
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}

// Init initializes the instance.
func (t NSTermOfAddress) Init() NSTermOfAddress {
	rv := objc.Send[NSTermOfAddress](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTermOfAddress) Autorelease() NSTermOfAddress {
	rv := objc.Send[NSTermOfAddress](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTermOfAddress creates a new NSTermOfAddress instance.
func NewNSTermOfAddress() NSTermOfAddress {
	class := getNSTermOfAddressClass()
	rv := objc.Send[NSTermOfAddress](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (t NSTermOfAddress) InitWithCoder(coder INSCoder) NSTermOfAddress {
	rv := objc.Send[NSTermOfAddress](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (t NSTermOfAddress) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Term of address that uses feminine pronouns (e.g. she/her/hers in English),
// and a feminine grammatical gender when inflecting verbs and adjectives
// referring to the person
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/feminine
func (_NSTermOfAddressClass NSTermOfAddressClass) Feminine() NSTermOfAddress {
	rv := objc.Send[objc.ID](objc.ID(_NSTermOfAddressClass.class), objc.Sel("feminine"))
	return NSTermOfAddressFromID(rv)
}
// Term of address that uses masculine pronouns (e.g. he/him/his in English),
// and a masculine grammatical gender when inflecting verbs and adjectives
// referring to the person
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/masculine
func (_NSTermOfAddressClass NSTermOfAddressClass) Masculine() NSTermOfAddress {
	rv := objc.Send[objc.ID](objc.ID(_NSTermOfAddressClass.class), objc.Sel("masculine"))
	return NSTermOfAddressFromID(rv)
}
// Term of address that uses gender-neutral pronouns (e.g. they/them/theirs in
// English), and an epicene grammatical gender when inflecting verbs and
// adjectives referring to the person
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/neutral
func (_NSTermOfAddressClass NSTermOfAddressClass) Neutral() NSTermOfAddress {
	rv := objc.Send[objc.ID](objc.ID(_NSTermOfAddressClass.class), objc.Sel("neutral"))
	return NSTermOfAddressFromID(rv)
}
// The term of address that should be used for addressing the user
//
// # Discussion
// 
// This term of address will only compare equal to another `+[NSTermOfAddress
// currentUser]`
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/currentUser
func (_NSTermOfAddressClass NSTermOfAddressClass) CurrentUser() NSTermOfAddress {
	rv := objc.Send[objc.ID](objc.ID(_NSTermOfAddressClass.class), objc.Sel("currentUser"))
	return NSTermOfAddressFromID(rv)
}
// A term of address restricted to a given language
//
// language: ISO language code identifier for the language
//
// pronouns: A list of pronouns in the target language that can be used to refer to the
// person.
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/localizedForLanguageIdentifier:withPronouns:
func (_NSTermOfAddressClass NSTermOfAddressClass) LocalizedForLanguageIdentifierWithPronouns(language string, pronouns []NSMorphologyPronoun) NSTermOfAddress {
	rv := objc.Send[objc.ID](objc.ID(_NSTermOfAddressClass.class), objc.Sel("localizedForLanguageIdentifier:withPronouns:"), objc.String(language), objectivec.IObjectSliceToNSArray(pronouns))
	return NSTermOfAddressFromID(rv)
}

// A list of pronouns for a localized term of address
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/pronouns
func (t NSTermOfAddress) Pronouns() []NSMorphologyPronoun {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("pronouns"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMorphologyPronoun {
		return NSMorphologyPronounFromID(id)
	})
}
// The ISO language code if this is a localized term of address
//
// See: https://developer.apple.com/documentation/Foundation/NSTermOfAddress/languageIdentifier
func (t NSTermOfAddress) LanguageIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("languageIdentifier"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

