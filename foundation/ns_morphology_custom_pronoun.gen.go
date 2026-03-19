// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMorphologyCustomPronoun] class.
var (
	_NSMorphologyCustomPronounClass     NSMorphologyCustomPronounClass
	_NSMorphologyCustomPronounClassOnce sync.Once
)

func getNSMorphologyCustomPronounClass() NSMorphologyCustomPronounClass {
	_NSMorphologyCustomPronounClassOnce.Do(func() {
		_NSMorphologyCustomPronounClass = NSMorphologyCustomPronounClass{class: objc.GetClass("NSMorphologyCustomPronoun")}
	})
	return _NSMorphologyCustomPronounClass
}

// GetNSMorphologyCustomPronounClass returns the class object for NSMorphologyCustomPronoun.
func GetNSMorphologyCustomPronounClass() NSMorphologyCustomPronounClass {
	return getNSMorphologyCustomPronounClass()
}

type NSMorphologyCustomPronounClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMorphologyCustomPronounClass) Alloc() NSMorphologyCustomPronoun {
	rv := objc.Send[NSMorphologyCustomPronoun](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A custom pronoun behavior for use in a specific langauge.
//
// # Overview
// 
// Set a [NSMorphologyCustomPronoun] instance on a [NSMorphology] instance
// when you want to provide a langauge-specific customization of pronoun use
// in that language. Different languages have different requirements for the
// grammatical information needed to apply a custom pronoun, so you set custom
// pronoun behavior on a per-language basis.
// 
// The example below shows how to create English “ze” and “hir” custom
// pronouns:
// 
// [NSMorphologyCustomPronoun] only supports third-person pronouns. Use this
// feature when your app needs to refer to third parties with a specific
// pronoun.
//
// # Determining Pronoun Forms
//
//   - [NSMorphologyCustomPronoun.SubjectForm]: The subject pronoun form to apply when using this custom pronoun behavior.
//   - [NSMorphologyCustomPronoun.SetSubjectForm]
//   - [NSMorphologyCustomPronoun.ObjectForm]: The object pronoun form to apply when using this custom pronoun behavior.
//   - [NSMorphologyCustomPronoun.SetObjectForm]
//   - [NSMorphologyCustomPronoun.PossessiveForm]: The posessive pronoun form to apply when using this custom pronoun behavior.
//   - [NSMorphologyCustomPronoun.SetPossessiveForm]
//   - [NSMorphologyCustomPronoun.PossessiveAdjectiveForm]: The posessive adjective pronoun form to apply when using this custom pronoun behavior.
//   - [NSMorphologyCustomPronoun.SetPossessiveAdjectiveForm]
//   - [NSMorphologyCustomPronoun.ReflexiveForm]: The reflexive pronoun form to apply when using this custom pronoun behavior.
//   - [NSMorphologyCustomPronoun.SetReflexiveForm]
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun
type NSMorphologyCustomPronoun struct {
	objectivec.Object
}

// NSMorphologyCustomPronounFromID constructs a [NSMorphologyCustomPronoun] from an objc.ID.
//
// A custom pronoun behavior for use in a specific langauge.
func NSMorphologyCustomPronounFromID(id objc.ID) NSMorphologyCustomPronoun {
	return NSMorphologyCustomPronoun{objectivec.Object{ID: id}}
}
// NOTE: NSMorphologyCustomPronoun adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMorphologyCustomPronoun] class.
//
// # Determining Pronoun Forms
//
//   - [INSMorphologyCustomPronoun.SubjectForm]: The subject pronoun form to apply when using this custom pronoun behavior.
//   - [INSMorphologyCustomPronoun.SetSubjectForm]
//   - [INSMorphologyCustomPronoun.ObjectForm]: The object pronoun form to apply when using this custom pronoun behavior.
//   - [INSMorphologyCustomPronoun.SetObjectForm]
//   - [INSMorphologyCustomPronoun.PossessiveForm]: The posessive pronoun form to apply when using this custom pronoun behavior.
//   - [INSMorphologyCustomPronoun.SetPossessiveForm]
//   - [INSMorphologyCustomPronoun.PossessiveAdjectiveForm]: The posessive adjective pronoun form to apply when using this custom pronoun behavior.
//   - [INSMorphologyCustomPronoun.SetPossessiveAdjectiveForm]
//   - [INSMorphologyCustomPronoun.ReflexiveForm]: The reflexive pronoun form to apply when using this custom pronoun behavior.
//   - [INSMorphologyCustomPronoun.SetReflexiveForm]
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun
type INSMorphologyCustomPronoun interface {
	objectivec.IObject
	NSCopying
	NSSecureCoding

	// Topic: Determining Pronoun Forms

	// The subject pronoun form to apply when using this custom pronoun behavior.
	SubjectForm() string
	SetSubjectForm(value string)
	// The object pronoun form to apply when using this custom pronoun behavior.
	ObjectForm() string
	SetObjectForm(value string)
	// The posessive pronoun form to apply when using this custom pronoun behavior.
	PossessiveForm() string
	SetPossessiveForm(value string)
	// The posessive adjective pronoun form to apply when using this custom pronoun behavior.
	PossessiveAdjectiveForm() string
	SetPossessiveAdjectiveForm(value string)
	// The reflexive pronoun form to apply when using this custom pronoun behavior.
	ReflexiveForm() string
	SetReflexiveForm(value string)

	InitWithCoder(coder INSCoder) NSMorphologyCustomPronoun
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}

// Init initializes the instance.
func (m NSMorphologyCustomPronoun) Init() NSMorphologyCustomPronoun {
	rv := objc.Send[NSMorphologyCustomPronoun](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMorphologyCustomPronoun) Autorelease() NSMorphologyCustomPronoun {
	rv := objc.Send[NSMorphologyCustomPronoun](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMorphologyCustomPronoun creates a new NSMorphologyCustomPronoun instance.
func NewNSMorphologyCustomPronoun() NSMorphologyCustomPronoun {
	class := getNSMorphologyCustomPronounClass()
	rv := objc.Send[NSMorphologyCustomPronoun](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (m NSMorphologyCustomPronoun) InitWithCoder(coder INSCoder) NSMorphologyCustomPronoun {
	rv := objc.Send[NSMorphologyCustomPronoun](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (m NSMorphologyCustomPronoun) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The subject pronoun form to apply when using this custom pronoun behavior.
//
// # Discussion
// 
// In the English phrase “she reads,” the pronoun “she” exhibits the
// subject form.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/subjectForm
func (m NSMorphologyCustomPronoun) SubjectForm() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("subjectForm"))
	return NSStringFromID(rv).String()
}
func (m NSMorphologyCustomPronoun) SetSubjectForm(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setSubjectForm:"), objc.String(value))
}
// The object pronoun form to apply when using this custom pronoun behavior.
//
// # Discussion
// 
// In the English phrase “ask him,” the pronoun “him” exhibits the
// object form.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/objectForm
func (m NSMorphologyCustomPronoun) ObjectForm() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectForm"))
	return NSStringFromID(rv).String()
}
func (m NSMorphologyCustomPronoun) SetObjectForm(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectForm:"), objc.String(value))
}
// The posessive pronoun form to apply when using this custom pronoun
// behavior.
//
// # Discussion
// 
// In the English phrase “the bike is hers,” the pronoun “hers”
// exhibits the posessive form.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/possessiveForm
func (m NSMorphologyCustomPronoun) PossessiveForm() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("possessiveForm"))
	return NSStringFromID(rv).String()
}
func (m NSMorphologyCustomPronoun) SetPossessiveForm(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setPossessiveForm:"), objc.String(value))
}
// The posessive adjective pronoun form to apply when using this custom
// pronoun behavior.
//
// # Discussion
// 
// In the English phrase “his bike,” the pronoun “his” exhibits the
// possesive adjective form.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/possessiveAdjectiveForm
func (m NSMorphologyCustomPronoun) PossessiveAdjectiveForm() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("possessiveAdjectiveForm"))
	return NSStringFromID(rv).String()
}
func (m NSMorphologyCustomPronoun) SetPossessiveAdjectiveForm(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setPossessiveAdjectiveForm:"), objc.String(value))
}
// The reflexive pronoun form to apply when using this custom pronoun
// behavior.
//
// # Discussion
// 
// The English pronoun “herself” is an example of the reflexive form.
//
// See: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun/reflexiveForm
func (m NSMorphologyCustomPronoun) ReflexiveForm() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("reflexiveForm"))
	return NSStringFromID(rv).String()
}
func (m NSMorphologyCustomPronoun) SetReflexiveForm(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setReflexiveForm:"), objc.String(value))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

