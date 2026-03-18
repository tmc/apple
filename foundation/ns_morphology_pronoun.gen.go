// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMorphologyPronoun] class.
var (
	_NSMorphologyPronounClass     NSMorphologyPronounClass
	_NSMorphologyPronounClassOnce sync.Once
)

func getNSMorphologyPronounClass() NSMorphologyPronounClass {
	_NSMorphologyPronounClassOnce.Do(func() {
		_NSMorphologyPronounClass = NSMorphologyPronounClass{class: objc.GetClass("NSMorphologyPronoun")}
	})
	return _NSMorphologyPronounClass
}

// GetNSMorphologyPronounClass returns the class object for NSMorphologyPronoun.
func GetNSMorphologyPronounClass() NSMorphologyPronounClass {
	return getNSMorphologyPronounClass()
}

type NSMorphologyPronounClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMorphologyPronounClass) Alloc() NSMorphologyPronoun {
	rv := objc.Send[NSMorphologyPronoun](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A custom pronoun for referring to a third person.
//
// # Overview
// 
// Create instances of [NSMorphologyPronoun] when you need to define custom
// pronouns for a localized term of address.
// 
// For examples of how to create custom pronouns, see [TermOfAddress].
//
// [TermOfAddress]: https://developer.apple.com/documentation/Foundation/TermOfAddress
//
// # Using pronouns
//
//   - [NSMorphologyPronoun.Pronoun]
//   - [NSMorphologyPronoun.Morphology]
//   - [NSMorphologyPronoun.DependentMorphology]
//
// # Instance Methods
//
//   - [NSMorphologyPronoun.InitWithPronounMorphologyDependentMorphology]
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun
type NSMorphologyPronoun struct {
	objectivec.Object
}

// NSMorphologyPronounFromID constructs a [NSMorphologyPronoun] from an objc.ID.
//
// A custom pronoun for referring to a third person.
func NSMorphologyPronounFromID(id objc.ID) NSMorphologyPronoun {
	return NSMorphologyPronoun{objectivec.Object{ID: id}}
}
// NOTE: NSMorphologyPronoun adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMorphologyPronoun] class.
//
// # Using pronouns
//
//   - [INSMorphologyPronoun.Pronoun]
//   - [INSMorphologyPronoun.Morphology]
//   - [INSMorphologyPronoun.DependentMorphology]
//
// # Instance Methods
//
//   - [INSMorphologyPronoun.InitWithPronounMorphologyDependentMorphology]
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun
type INSMorphologyPronoun interface {
	objectivec.IObject
	NSCopying
	NSSecureCoding

	// Topic: Using pronouns

	Pronoun() string
	Morphology() INSMorphology
	DependentMorphology() INSMorphology

	// Topic: Instance Methods

	InitWithPronounMorphologyDependentMorphology(pronoun string, morphology INSMorphology, dependentMorphology INSMorphology) NSMorphologyPronoun

	InitWithCoder(coder INSCoder) NSMorphologyPronoun
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}

// Init initializes the instance.
func (m NSMorphologyPronoun) Init() NSMorphologyPronoun {
	rv := objc.Send[NSMorphologyPronoun](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMorphologyPronoun) Autorelease() NSMorphologyPronoun {
	rv := objc.Send[NSMorphologyPronoun](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMorphologyPronoun creates a new NSMorphologyPronoun instance.
func NewNSMorphologyPronoun() NSMorphologyPronoun {
	class := getNSMorphologyPronounClass()
	rv := objc.Send[NSMorphologyPronoun](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/initWithPronoun:morphology:dependentMorphology:
func NewMorphologyPronounWithPronounMorphologyDependentMorphology(pronoun string, morphology INSMorphology, dependentMorphology INSMorphology) NSMorphologyPronoun {
	instance := getNSMorphologyPronounClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPronoun:morphology:dependentMorphology:"), objc.String(pronoun), morphology, dependentMorphology)
	return NSMorphologyPronounFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/initWithPronoun:morphology:dependentMorphology:
func (m NSMorphologyPronoun) InitWithPronounMorphologyDependentMorphology(pronoun string, morphology INSMorphology, dependentMorphology INSMorphology) NSMorphologyPronoun {
	rv := objc.Send[NSMorphologyPronoun](m.ID, objc.Sel("initWithPronoun:morphology:dependentMorphology:"), objc.String(pronoun), morphology, dependentMorphology)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (m NSMorphologyPronoun) InitWithCoder(coder INSCoder) NSMorphologyPronoun {
	rv := objc.Send[NSMorphologyPronoun](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (m NSMorphologyPronoun) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/pronoun
func (m NSMorphologyPronoun) Pronoun() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pronoun"))
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/morphology
func (m NSMorphologyPronoun) Morphology() INSMorphology {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("morphology"))
	return NSMorphologyFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun/dependentMorphology
func (m NSMorphologyPronoun) DependentMorphology() INSMorphology {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dependentMorphology"))
	return NSMorphologyFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

