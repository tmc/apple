// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXBrailleTranslator] class.
var (
	_AXBrailleTranslatorClass     AXBrailleTranslatorClass
	_AXBrailleTranslatorClassOnce sync.Once
)

func getAXBrailleTranslatorClass() AXBrailleTranslatorClass {
	_AXBrailleTranslatorClassOnce.Do(func() {
		_AXBrailleTranslatorClass = AXBrailleTranslatorClass{class: objc.GetClass("AXBrailleTranslator")}
	})
	return _AXBrailleTranslatorClass
}

// GetAXBrailleTranslatorClass returns the class object for AXBrailleTranslator.
func GetAXBrailleTranslatorClass() AXBrailleTranslatorClass {
	return getAXBrailleTranslatorClass()
}

type AXBrailleTranslatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXBrailleTranslatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXBrailleTranslatorClass) Alloc() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Translates print text to Braille and Braille to print text according to the
// given Braille table.
//
// # Initializers
//
//   - [AXBrailleTranslator.InitWithBrailleTable]
//
// # Instance Methods
//
//   - [AXBrailleTranslator.BackTranslateBraille]: Input Braille should use the unicode Braille characters (0x2800-0x28FF).
//   - [AXBrailleTranslator.TranslatePrintText]: Output Braille uses the unicode Braille characters (0x2800-0x28FF).
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator
type AXBrailleTranslator struct {
	objectivec.Object
}

// AXBrailleTranslatorFromID constructs a [AXBrailleTranslator] from an objc.ID.
//
// Translates print text to Braille and Braille to print text according to the
// given Braille table.
func AXBrailleTranslatorFromID(id objc.ID) AXBrailleTranslator {
	return AXBrailleTranslator{objectivec.Object{ID: id}}
}

// NOTE: AXBrailleTranslator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXBrailleTranslator] class.
//
// # Initializers
//
//   - [IAXBrailleTranslator.InitWithBrailleTable]
//
// # Instance Methods
//
//   - [IAXBrailleTranslator.BackTranslateBraille]: Input Braille should use the unicode Braille characters (0x2800-0x28FF).
//   - [IAXBrailleTranslator.TranslatePrintText]: Output Braille uses the unicode Braille characters (0x2800-0x28FF).
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator
type IAXBrailleTranslator interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithBrailleTable(brailleTable IAXBrailleTable) AXBrailleTranslator

	// Topic: Instance Methods

	// Input Braille should use the unicode Braille characters (0x2800-0x28FF).
	BackTranslateBraille(braille string) IAXBrailleTranslationResult
	// Output Braille uses the unicode Braille characters (0x2800-0x28FF).
	TranslatePrintText(printText string) IAXBrailleTranslationResult
}

// Init initializes the instance.
func (a AXBrailleTranslator) Init() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXBrailleTranslator) Autorelease() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleTranslator creates a new AXBrailleTranslator instance.
func NewAXBrailleTranslator() AXBrailleTranslator {
	class := getAXBrailleTranslatorClass()
	rv := objc.Send[AXBrailleTranslator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/init(brailleTable:)
func NewAXBrailleTranslatorWithBrailleTable(brailleTable IAXBrailleTable) AXBrailleTranslator {
	instance := getAXBrailleTranslatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBrailleTable:"), brailleTable)
	return AXBrailleTranslatorFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/init(brailleTable:)
func (a AXBrailleTranslator) InitWithBrailleTable(brailleTable IAXBrailleTable) AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](a.ID, objc.Sel("initWithBrailleTable:"), brailleTable)
	return rv
}

// Input Braille should use the unicode Braille characters (0x2800-0x28FF).
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/backTranslateBraille(_:)
func (a AXBrailleTranslator) BackTranslateBraille(braille string) IAXBrailleTranslationResult {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("backTranslateBraille:"), objc.String(braille))
	return AXBrailleTranslationResultFromID(rv)
}

// Output Braille uses the unicode Braille characters (0x2800-0x28FF).
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/translatePrintText(_:)
func (a AXBrailleTranslator) TranslatePrintText(printText string) IAXBrailleTranslationResult {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("translatePrintText:"), objc.String(printText))
	return AXBrailleTranslationResultFromID(rv)
}
