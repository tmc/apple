// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMorphology] class.
var (
	_NSMorphologyClass     NSMorphologyClass
	_NSMorphologyClassOnce sync.Once
)

func getNSMorphologyClass() NSMorphologyClass {
	_NSMorphologyClassOnce.Do(func() {
		_NSMorphologyClass = NSMorphologyClass{class: objc.GetClass("NSMorphology")}
	})
	return _NSMorphologyClass
}

// GetNSMorphologyClass returns the class object for NSMorphology.
func GetNSMorphologyClass() NSMorphologyClass {
	return getNSMorphologyClass()
}

type NSMorphologyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMorphologyClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMorphologyClass) Alloc() NSMorphology {
	rv := objc.Send[NSMorphology](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A description of the grammatical properties of a string.
//
// # Overview
//
// Use a morphology with an [NSInflectionRule] to specify how to interpret a
// specific word when inflecting an [NSAttributedString]. This affects
// grammatical agreement with traits like number and gender, as well as
// declaring the word’s part of speech.
//
// The [NSMorphology] type’s design is language-independent; the concepts it
// can specify encompass the spectrum of what languages can do. Even for
// languages that don’t have one or more of those properties benefit the
// system as hints to make appropriate choices even when an exact inflection
// isn’t possible. Examples of properties absent from languages include
// Spanish’s lack of a grammatical gender of neuter, or the nonexistence of
// a paucal (plural few) pronoun in English.
//
// # Accessing Grammatical Properties
//
//   - [NSMorphology.Unspecified]: A Boolean value that indicates whether this instance specifies no particular grammar.
//   - [NSMorphology.GrammaticalGender]: The grammatical gender used for inflecting strings with this morphology.
//   - [NSMorphology.SetGrammaticalGender]
//   - [NSMorphology.Number]: The grammatical number used for inflecting strings with this morphology.
//   - [NSMorphology.SetNumber]
//   - [NSMorphology.PartOfSpeech]: The grammatical part of speech used for inflecting strings with this morphology.
//   - [NSMorphology.SetPartOfSpeech]
//
// # Instance Properties
//
//   - [NSMorphology.Definiteness]
//   - [NSMorphology.SetDefiniteness]
//   - [NSMorphology.Determination]
//   - [NSMorphology.SetDetermination]
//   - [NSMorphology.GrammaticalCase]
//   - [NSMorphology.SetGrammaticalCase]
//   - [NSMorphology.GrammaticalPerson]
//   - [NSMorphology.SetGrammaticalPerson]
//   - [NSMorphology.PronounType]
//   - [NSMorphology.SetPronounType]
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology
type NSMorphology struct {
	objectivec.Object
}

// NSMorphologyFromID constructs a [NSMorphology] from an objc.ID.
//
// A description of the grammatical properties of a string.
func NSMorphologyFromID(id objc.ID) NSMorphology {
	return NSMorphology{objectivec.Object{ID: id}}
}

// NOTE: NSMorphology adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMorphology] class.
//
// # Accessing Grammatical Properties
//
//   - [INSMorphology.Unspecified]: A Boolean value that indicates whether this instance specifies no particular grammar.
//   - [INSMorphology.GrammaticalGender]: The grammatical gender used for inflecting strings with this morphology.
//   - [INSMorphology.SetGrammaticalGender]
//   - [INSMorphology.Number]: The grammatical number used for inflecting strings with this morphology.
//   - [INSMorphology.SetNumber]
//   - [INSMorphology.PartOfSpeech]: The grammatical part of speech used for inflecting strings with this morphology.
//   - [INSMorphology.SetPartOfSpeech]
//
// # Instance Properties
//
//   - [INSMorphology.Definiteness]
//   - [INSMorphology.SetDefiniteness]
//   - [INSMorphology.Determination]
//   - [INSMorphology.SetDetermination]
//   - [INSMorphology.GrammaticalCase]
//   - [INSMorphology.SetGrammaticalCase]
//   - [INSMorphology.GrammaticalPerson]
//   - [INSMorphology.SetGrammaticalPerson]
//   - [INSMorphology.PronounType]
//   - [INSMorphology.SetPronounType]
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology
type INSMorphology interface {
	objectivec.IObject
	NSCopying
	NSSecureCoding

	// Topic: Accessing Grammatical Properties

	// A Boolean value that indicates whether this instance specifies no particular grammar.
	Unspecified() bool
	// The grammatical gender used for inflecting strings with this morphology.
	GrammaticalGender() NSGrammaticalGender
	SetGrammaticalGender(value NSGrammaticalGender)
	// The grammatical number used for inflecting strings with this morphology.
	Number() NSGrammaticalNumber
	SetNumber(value NSGrammaticalNumber)
	// The grammatical part of speech used for inflecting strings with this morphology.
	PartOfSpeech() NSGrammaticalPartOfSpeech
	SetPartOfSpeech(value NSGrammaticalPartOfSpeech)

	// Topic: Instance Properties

	Definiteness() NSGrammaticalDefiniteness
	SetDefiniteness(value NSGrammaticalDefiniteness)
	Determination() NSGrammaticalDetermination
	SetDetermination(value NSGrammaticalDetermination)
	GrammaticalCase() NSGrammaticalCase
	SetGrammaticalCase(value NSGrammaticalCase)
	GrammaticalPerson() NSGrammaticalPerson
	SetGrammaticalPerson(value NSGrammaticalPerson)
	PronounType() NSGrammaticalPronounType
	SetPronounType(value NSGrammaticalPronounType)

	InitWithCoder(coder INSCoder) NSMorphology
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}

// Init initializes the instance.
func (m NSMorphology) Init() NSMorphology {
	rv := objc.Send[NSMorphology](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMorphology) Autorelease() NSMorphology {
	rv := objc.Send[NSMorphology](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMorphology creates a new NSMorphology instance.
func NewNSMorphology() NSMorphology {
	class := getNSMorphologyClass()
	rv := objc.Send[NSMorphology](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (m NSMorphology) InitWithCoder(coder INSCoder) NSMorphology {
	rv := objc.Send[NSMorphology](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (m NSMorphology) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether this instance specifies no
// particular grammar.
//
// # Discussion
//
// This value is equivalent to having set none of the properties in this
// [NSMorphology]. This occurs when the user hasn’t specified preferences,
// or chose not to share them with this app. When the morphology is
// unspecified, inflecting a string with this morphology does nothing.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology/unspecified
func (m NSMorphology) Unspecified() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isUnspecified"))
	return rv
}

// The grammatical gender used for inflecting strings with this morphology.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalGender
func (m NSMorphology) GrammaticalGender() NSGrammaticalGender {
	rv := objc.Send[NSGrammaticalGender](m.ID, objc.Sel("grammaticalGender"))
	return NSGrammaticalGender(rv)
}
func (m NSMorphology) SetGrammaticalGender(value NSGrammaticalGender) {
	objc.Send[struct{}](m.ID, objc.Sel("setGrammaticalGender:"), value)
}

// The grammatical number used for inflecting strings with this morphology.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology/number
func (m NSMorphology) Number() NSGrammaticalNumber {
	rv := objc.Send[NSGrammaticalNumber](m.ID, objc.Sel("number"))
	return NSGrammaticalNumber(rv)
}
func (m NSMorphology) SetNumber(value NSGrammaticalNumber) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumber:"), value)
}

// The grammatical part of speech used for inflecting strings with this
// morphology.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology/partOfSpeech
func (m NSMorphology) PartOfSpeech() NSGrammaticalPartOfSpeech {
	rv := objc.Send[NSGrammaticalPartOfSpeech](m.ID, objc.Sel("partOfSpeech"))
	return NSGrammaticalPartOfSpeech(rv)
}
func (m NSMorphology) SetPartOfSpeech(value NSGrammaticalPartOfSpeech) {
	objc.Send[struct{}](m.ID, objc.Sel("setPartOfSpeech:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphology/definiteness
func (m NSMorphology) Definiteness() NSGrammaticalDefiniteness {
	rv := objc.Send[NSGrammaticalDefiniteness](m.ID, objc.Sel("definiteness"))
	return NSGrammaticalDefiniteness(rv)
}
func (m NSMorphology) SetDefiniteness(value NSGrammaticalDefiniteness) {
	objc.Send[struct{}](m.ID, objc.Sel("setDefiniteness:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphology/determination
func (m NSMorphology) Determination() NSGrammaticalDetermination {
	rv := objc.Send[NSGrammaticalDetermination](m.ID, objc.Sel("determination"))
	return NSGrammaticalDetermination(rv)
}
func (m NSMorphology) SetDetermination(value NSGrammaticalDetermination) {
	objc.Send[struct{}](m.ID, objc.Sel("setDetermination:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalCase
func (m NSMorphology) GrammaticalCase() NSGrammaticalCase {
	rv := objc.Send[NSGrammaticalCase](m.ID, objc.Sel("grammaticalCase"))
	return NSGrammaticalCase(rv)
}
func (m NSMorphology) SetGrammaticalCase(value NSGrammaticalCase) {
	objc.Send[struct{}](m.ID, objc.Sel("setGrammaticalCase:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphology/grammaticalPerson
func (m NSMorphology) GrammaticalPerson() NSGrammaticalPerson {
	rv := objc.Send[NSGrammaticalPerson](m.ID, objc.Sel("grammaticalPerson"))
	return NSGrammaticalPerson(rv)
}
func (m NSMorphology) SetGrammaticalPerson(value NSGrammaticalPerson) {
	objc.Send[struct{}](m.ID, objc.Sel("setGrammaticalPerson:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphology/pronounType
func (m NSMorphology) PronounType() NSGrammaticalPronounType {
	rv := objc.Send[NSGrammaticalPronounType](m.ID, objc.Sel("pronounType"))
	return NSGrammaticalPronounType(rv)
}
func (m NSMorphology) SetPronounType(value NSGrammaticalPronounType) {
	objc.Send[struct{}](m.ID, objc.Sel("setPronounType:"), value)
}

// The addressing preferences of the current user.
//
// # Discussion
//
// If the user hasn’t specified preferences, or chose not to share them with
// this app, the [Unspecified] property is `true`.
//
// This value doesn’t change throughout the lifetime of the process.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphology/userMorphology
func (_NSMorphologyClass NSMorphologyClass) UserMorphology() NSMorphology {
	rv := objc.Send[objc.ID](objc.ID(_NSMorphologyClass.class), objc.Sel("userMorphology"))
	return NSMorphologyFromID(objc.ID(rv))
}

// Protocol methods for NSCopying

// Protocol methods for NSSecureCoding
