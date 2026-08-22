// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXBrailleTable] class.
var (
	_AXBrailleTableClass     AXBrailleTableClass
	_AXBrailleTableClassOnce sync.Once
)

func getAXBrailleTableClass() AXBrailleTableClass {
	_AXBrailleTableClassOnce.Do(func() {
		_AXBrailleTableClass = AXBrailleTableClass{class: objc.GetClass("AXBrailleTable")}
	})
	return _AXBrailleTableClass
}

// GetAXBrailleTableClass returns the class object for AXBrailleTable.
func GetAXBrailleTableClass() AXBrailleTableClass {
	return getAXBrailleTableClass()
}

type AXBrailleTableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXBrailleTableClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXBrailleTableClass) Alloc() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A rule for translating print text to Braille, and back-translating Braille
// to print text.
//
// # Initializers
//
//   - [AXBrailleTable.InitWithCoder]
//   - [AXBrailleTable.InitWithIdentifier]: Returns nil if there is no table with the given identifier.
//
// # Instance Properties
//
//   - [AXBrailleTable.Identifier]: A unique string that identifies this table.
//   - [AXBrailleTable.IsEightDot]: Returns true if this table makes use of eight dots as opposed to six dots.
//   - [AXBrailleTable.Language]: The 3-character code from ISO 639-2 for the language this Braille table pertains to.
//   - [AXBrailleTable.Locales]: All locales this table supports.
//   - [AXBrailleTable.LocalizedName]: The localized name of this table for user display.
//   - [AXBrailleTable.LocalizedProviderName]: The localized name of the provider of this table for user display.
//   - [AXBrailleTable.ProviderIdentifier]: The identifier of the provider of this table.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable
type AXBrailleTable struct {
	objectivec.Object
}

// AXBrailleTableFromID constructs a [AXBrailleTable] from an objc.ID.
//
// A rule for translating print text to Braille, and back-translating Braille
// to print text.
func AXBrailleTableFromID(id objc.ID) AXBrailleTable {
	return AXBrailleTable{objectivec.Object{ID: id}}
}

// NOTE: AXBrailleTable adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXBrailleTable] class.
//
// # Initializers
//
//   - [IAXBrailleTable.InitWithCoder]
//   - [IAXBrailleTable.InitWithIdentifier]: Returns nil if there is no table with the given identifier.
//
// # Instance Properties
//
//   - [IAXBrailleTable.Identifier]: A unique string that identifies this table.
//   - [IAXBrailleTable.IsEightDot]: Returns true if this table makes use of eight dots as opposed to six dots.
//   - [IAXBrailleTable.Language]: The 3-character code from ISO 639-2 for the language this Braille table pertains to.
//   - [IAXBrailleTable.Locales]: All locales this table supports.
//   - [IAXBrailleTable.LocalizedName]: The localized name of this table for user display.
//   - [IAXBrailleTable.LocalizedProviderName]: The localized name of the provider of this table for user display.
//   - [IAXBrailleTable.ProviderIdentifier]: The identifier of the provider of this table.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable
type IAXBrailleTable interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) AXBrailleTable
	// Returns nil if there is no table with the given identifier.
	InitWithIdentifier(identifier string) AXBrailleTable

	// Topic: Instance Properties

	// A unique string that identifies this table.
	Identifier() string
	// Returns true if this table makes use of eight dots as opposed to six dots.
	IsEightDot() bool
	// The 3-character code from ISO 639-2 for the language this Braille table pertains to.
	Language() string
	// All locales this table supports.
	Locales() foundation.INSSet
	// The localized name of this table for user display.
	LocalizedName() string
	// The localized name of the provider of this table for user display.
	LocalizedProviderName() string
	// The identifier of the provider of this table.
	ProviderIdentifier() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AXBrailleTable) Init() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXBrailleTable) Autorelease() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleTable creates a new AXBrailleTable instance.
func NewAXBrailleTable() AXBrailleTable {
	class := getAXBrailleTableClass()
	rv := objc.Send[AXBrailleTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/init(coder:)
func NewAXBrailleTableWithCoder(coder foundation.INSCoder) AXBrailleTable {
	instance := getAXBrailleTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AXBrailleTableFromID(rv)
}

// Returns nil if there is no table with the given identifier.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/init(identifier:)
func NewAXBrailleTableWithIdentifier(identifier string) AXBrailleTable {
	instance := getAXBrailleTableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return AXBrailleTableFromID(rv)
}

// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/init(coder:)
func (a AXBrailleTable) InitWithCoder(coder foundation.INSCoder) AXBrailleTable {
	rv := objc.Send[AXBrailleTable](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns nil if there is no table with the given identifier.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/init(identifier:)
func (a AXBrailleTable) InitWithIdentifier(identifier string) AXBrailleTable {
	rv := objc.Send[AXBrailleTable](a.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return rv
}
func (a AXBrailleTable) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The default table that provides translations for the given locale’s
// language. Returns nil if there is none.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/defaultTable(for:)
func (_AXBrailleTableClass AXBrailleTableClass) DefaultTableForLocale(locale foundation.NSLocale) AXBrailleTable {
	rv := objc.Send[objc.ID](objc.ID(_AXBrailleTableClass.class), objc.Sel("defaultTableForLocale:"), locale)
	return AXBrailleTableFromID(rv)
}

// All tables that are not specific to any language.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/languageAgnosticTables()
func (_AXBrailleTableClass AXBrailleTableClass) LanguageAgnosticTables() foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_AXBrailleTableClass.class), objc.Sel("languageAgnosticTables"))
	return foundation.NSSetFromID(rv)
}

// All locales supported by existing tables.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/supportedLocales()
func (_AXBrailleTableClass AXBrailleTableClass) SupportedLocales() foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_AXBrailleTableClass.class), objc.Sel("supportedLocales"))
	return foundation.NSSetFromID(rv)
}

// All tables that provide translations for the given locale’s language.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/tables(for:)
func (_AXBrailleTableClass AXBrailleTableClass) TablesForLocale(locale foundation.NSLocale) foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_AXBrailleTableClass.class), objc.Sel("tablesForLocale:"), locale)
	return foundation.NSSetFromID(rv)
}

// A unique string that identifies this table.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/identifier
func (a AXBrailleTable) Identifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns true if this table makes use of eight dots as opposed to six dots.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/isEightDot
func (a AXBrailleTable) IsEightDot() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEightDot"))
	return rv
}

// The 3-character code from ISO 639-2 for the language this Braille table
// pertains to.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/language-3570f
func (a AXBrailleTable) Language() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}

// All locales this table supports.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/locales
func (a AXBrailleTable) Locales() foundation.INSSet {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("locales"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The localized name of this table for user display.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedName
func (a AXBrailleTable) LocalizedName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// The localized name of the provider of this table for user display.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedProviderName
func (a AXBrailleTable) LocalizedProviderName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("localizedProviderName"))
	return foundation.NSStringFromID(rv).String()
}

// The identifier of the provider of this table.
//
// See: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/providerIdentifier
func (a AXBrailleTable) ProviderIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("providerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
