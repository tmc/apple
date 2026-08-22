// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("CoreText: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreText: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("CoreText: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("CoreText: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _cTFontCollectionCopyExclusionDescriptors func(collection CTFontCollectionRef) corefoundation.CFArrayRef
var _cTFontCollectionCopyExclusionDescriptorsErr error

func tryCTFontCollectionCopyExclusionDescriptors(collection CTFontCollectionRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCopyExclusionDescriptors == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCopyExclusionDescriptors", "10.7", _cTFontCollectionCopyExclusionDescriptorsErr)
	}
	return _cTFontCollectionCopyExclusionDescriptors(collection), nil
}

// CTFontCollectionCopyExclusionDescriptors retrieves the array of descriptors to exclude from the match.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyExclusionDescriptors(_:)
func CTFontCollectionCopyExclusionDescriptors(collection CTFontCollectionRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCopyExclusionDescriptors(collection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCopyFontAttribute func(collection CTFontCollectionRef, attributeName corefoundation.CFStringRef, options CTFontCollectionCopyOptions) corefoundation.CFArrayRef
var _cTFontCollectionCopyFontAttributeErr error

func tryCTFontCollectionCopyFontAttribute(collection CTFontCollectionRef, attributeName corefoundation.CFStringRef, options CTFontCollectionCopyOptions) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCopyFontAttribute == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCopyFontAttribute", "10.7", _cTFontCollectionCopyFontAttributeErr)
	}
	return _cTFontCollectionCopyFontAttribute(collection, attributeName, options), nil
}

// CTFontCollectionCopyFontAttribute retrieves an array of font descriptor attribute values.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyFontAttribute(_:_:_:)
func CTFontCollectionCopyFontAttribute(collection CTFontCollectionRef, attributeName corefoundation.CFStringRef, options CTFontCollectionCopyOptions) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCopyFontAttribute(collection, attributeName, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCopyFontAttributes func(collection CTFontCollectionRef, attributeNames corefoundation.CFSetRef, options CTFontCollectionCopyOptions) corefoundation.CFArrayRef
var _cTFontCollectionCopyFontAttributesErr error

func tryCTFontCollectionCopyFontAttributes(collection CTFontCollectionRef, attributeNames corefoundation.CFSetRef, options CTFontCollectionCopyOptions) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCopyFontAttributes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCopyFontAttributes", "10.7", _cTFontCollectionCopyFontAttributesErr)
	}
	return _cTFontCollectionCopyFontAttributes(collection, attributeNames, options), nil
}

// CTFontCollectionCopyFontAttributes retrieves an array of dictionaries containing font descriptor attribute values.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyFontAttributes(_:_:_:)
func CTFontCollectionCopyFontAttributes(collection CTFontCollectionRef, attributeNames corefoundation.CFSetRef, options CTFontCollectionCopyOptions) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCopyFontAttributes(collection, attributeNames, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCopyQueryDescriptors func(collection CTFontCollectionRef) corefoundation.CFArrayRef
var _cTFontCollectionCopyQueryDescriptorsErr error

func tryCTFontCollectionCopyQueryDescriptors(collection CTFontCollectionRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCopyQueryDescriptors == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCopyQueryDescriptors", "10.7", _cTFontCollectionCopyQueryDescriptorsErr)
	}
	return _cTFontCollectionCopyQueryDescriptors(collection), nil
}

// CTFontCollectionCopyQueryDescriptors retrieves the array of descriptors for font matching.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCopyQueryDescriptors(_:)
func CTFontCollectionCopyQueryDescriptors(collection CTFontCollectionRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCopyQueryDescriptors(collection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateCopyWithFontDescriptors func(original CTFontCollectionRef, queryDescriptors corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) CTFontCollectionRef
var _cTFontCollectionCreateCopyWithFontDescriptorsErr error

func tryCTFontCollectionCreateCopyWithFontDescriptors(original CTFontCollectionRef, queryDescriptors corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) (CTFontCollectionRef, error) {
	if _cTFontCollectionCreateCopyWithFontDescriptors == nil {
		return *new(CTFontCollectionRef), symbolCallError("CTFontCollectionCreateCopyWithFontDescriptors", "10.5", _cTFontCollectionCreateCopyWithFontDescriptorsErr)
	}
	return _cTFontCollectionCreateCopyWithFontDescriptors(original, queryDescriptors, options), nil
}

// CTFontCollectionCreateCopyWithFontDescriptors returns a copy of the original collection augmented with the given new font descriptors.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateCopyWithFontDescriptors(_:_:_:)
func CTFontCollectionCreateCopyWithFontDescriptors(original CTFontCollectionRef, queryDescriptors corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) CTFontCollectionRef {
	result, callErr := tryCTFontCollectionCreateCopyWithFontDescriptors(original, queryDescriptors, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateFromAvailableFonts func(options corefoundation.CFDictionaryRef) CTFontCollectionRef
var _cTFontCollectionCreateFromAvailableFontsErr error

func tryCTFontCollectionCreateFromAvailableFonts(options corefoundation.CFDictionaryRef) (CTFontCollectionRef, error) {
	if _cTFontCollectionCreateFromAvailableFonts == nil {
		return *new(CTFontCollectionRef), symbolCallError("CTFontCollectionCreateFromAvailableFonts", "10.5", _cTFontCollectionCreateFromAvailableFontsErr)
	}
	return _cTFontCollectionCreateFromAvailableFonts(options), nil
}

// CTFontCollectionCreateFromAvailableFonts returns a new font collection containing all available fonts.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateFromAvailableFonts(_:)
func CTFontCollectionCreateFromAvailableFonts(options corefoundation.CFDictionaryRef) CTFontCollectionRef {
	result, callErr := tryCTFontCollectionCreateFromAvailableFonts(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateMatchingFontDescriptors func(collection CTFontCollectionRef) corefoundation.CFArrayRef
var _cTFontCollectionCreateMatchingFontDescriptorsErr error

func tryCTFontCollectionCreateMatchingFontDescriptors(collection CTFontCollectionRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCreateMatchingFontDescriptors == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCreateMatchingFontDescriptors", "10.5", _cTFontCollectionCreateMatchingFontDescriptorsErr)
	}
	return _cTFontCollectionCreateMatchingFontDescriptors(collection), nil
}

// CTFontCollectionCreateMatchingFontDescriptors returns an array of font descriptors matching the collection.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptors(_:)
func CTFontCollectionCreateMatchingFontDescriptors(collection CTFontCollectionRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCreateMatchingFontDescriptors(collection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateMatchingFontDescriptorsForFamily func(collection CTFontCollectionRef, familyName corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef
var _cTFontCollectionCreateMatchingFontDescriptorsForFamilyErr error

func tryCTFontCollectionCreateMatchingFontDescriptorsForFamily(collection CTFontCollectionRef, familyName corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCreateMatchingFontDescriptorsForFamily == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCreateMatchingFontDescriptorsForFamily", "10.7", _cTFontCollectionCreateMatchingFontDescriptorsForFamilyErr)
	}
	return _cTFontCollectionCreateMatchingFontDescriptorsForFamily(collection, familyName, options), nil
}

// CTFontCollectionCreateMatchingFontDescriptorsForFamily retrieves an array of font descriptors that match the specified family, one descriptor for each style in the collection.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsForFamily(_:_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsForFamily(collection CTFontCollectionRef, familyName corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCreateMatchingFontDescriptorsForFamily(collection, familyName, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback func(collection CTFontCollectionRef, sortCallback CTFontCollectionSortDescriptorsCallback, refCon uintptr) corefoundation.CFArrayRef
var _cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallbackErr error

func tryCTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection CTFontCollectionRef, sortCallback CTFontCollectionSortDescriptorsCallback, refCon uintptr) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback", "10.5", _cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallbackErr)
	}
	return _cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection, sortCallback, refCon), nil
}

// CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback returns the array of matching font descriptors sorted with the callback function.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(_:_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection CTFontCollectionRef, sortCallback CTFontCollectionSortDescriptorsCallback, refCon uintptr) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback(collection, sortCallback, refCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateMatchingFontDescriptorsWithOptions func(collection CTFontCollectionRef, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef
var _cTFontCollectionCreateMatchingFontDescriptorsWithOptionsErr error

func tryCTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection CTFontCollectionRef, options corefoundation.CFDictionaryRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCollectionCreateMatchingFontDescriptorsWithOptions == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCollectionCreateMatchingFontDescriptorsWithOptions", "10.7", _cTFontCollectionCreateMatchingFontDescriptorsWithOptionsErr)
	}
	return _cTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection, options), nil
}

// CTFontCollectionCreateMatchingFontDescriptorsWithOptions creates an array of font descriptors that match the specified collection.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMatchingFontDescriptorsWithOptions(_:_:)
func CTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection CTFontCollectionRef, options corefoundation.CFDictionaryRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCollectionCreateMatchingFontDescriptorsWithOptions(collection, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateMutableCopy func(original CTFontCollectionRef) CTMutableFontCollectionRef
var _cTFontCollectionCreateMutableCopyErr error

func tryCTFontCollectionCreateMutableCopy(original CTFontCollectionRef) (CTMutableFontCollectionRef, error) {
	if _cTFontCollectionCreateMutableCopy == nil {
		return *new(CTMutableFontCollectionRef), symbolCallError("CTFontCollectionCreateMutableCopy", "10.7", _cTFontCollectionCreateMutableCopyErr)
	}
	return _cTFontCollectionCreateMutableCopy(original), nil
}

// CTFontCollectionCreateMutableCopy creates a mutable copy of the original collection.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateMutableCopy(_:)
func CTFontCollectionCreateMutableCopy(original CTFontCollectionRef) CTMutableFontCollectionRef {
	result, callErr := tryCTFontCollectionCreateMutableCopy(original)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionCreateWithFontDescriptors func(queryDescriptors corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) CTFontCollectionRef
var _cTFontCollectionCreateWithFontDescriptorsErr error

func tryCTFontCollectionCreateWithFontDescriptors(queryDescriptors corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) (CTFontCollectionRef, error) {
	if _cTFontCollectionCreateWithFontDescriptors == nil {
		return *new(CTFontCollectionRef), symbolCallError("CTFontCollectionCreateWithFontDescriptors", "10.5", _cTFontCollectionCreateWithFontDescriptorsErr)
	}
	return _cTFontCollectionCreateWithFontDescriptors(queryDescriptors, options), nil
}

// CTFontCollectionCreateWithFontDescriptors returns a new font collection based on the given array of font descriptors.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionCreateWithFontDescriptors(_:_:)
func CTFontCollectionCreateWithFontDescriptors(queryDescriptors corefoundation.CFArrayRef, options corefoundation.CFDictionaryRef) CTFontCollectionRef {
	result, callErr := tryCTFontCollectionCreateWithFontDescriptors(queryDescriptors, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionGetTypeID func() uint
var _cTFontCollectionGetTypeIDErr error

func tryCTFontCollectionGetTypeID() (uint, error) {
	if _cTFontCollectionGetTypeID == nil {
		return 0, symbolCallError("CTFontCollectionGetTypeID", "10.5", _cTFontCollectionGetTypeIDErr)
	}
	return _cTFontCollectionGetTypeID(), nil
}

// CTFontCollectionGetTypeID returns the type identifier for Core Text font collection references.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionGetTypeID()
func CTFontCollectionGetTypeID() uint {
	result, callErr := tryCTFontCollectionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCollectionSetExclusionDescriptors func(collection CTMutableFontCollectionRef, descriptors corefoundation.CFArrayRef)
var _cTFontCollectionSetExclusionDescriptorsErr error

func tryCTFontCollectionSetExclusionDescriptors(collection CTMutableFontCollectionRef, descriptors corefoundation.CFArrayRef) error {
	if _cTFontCollectionSetExclusionDescriptors == nil {
		return symbolCallError("CTFontCollectionSetExclusionDescriptors", "10.7", _cTFontCollectionSetExclusionDescriptorsErr)
	}
	_cTFontCollectionSetExclusionDescriptors(collection, descriptors)
	return nil
}

// CTFontCollectionSetExclusionDescriptors replaces the array of descriptors to exclude from the match.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionSetExclusionDescriptors(_:_:)
func CTFontCollectionSetExclusionDescriptors(collection CTMutableFontCollectionRef, descriptors corefoundation.CFArrayRef) {
	if callErr := tryCTFontCollectionSetExclusionDescriptors(collection, descriptors); callErr != nil {
		panic(callErr)
	}
}

var _cTFontCollectionSetQueryDescriptors func(collection CTMutableFontCollectionRef, descriptors corefoundation.CFArrayRef)
var _cTFontCollectionSetQueryDescriptorsErr error

func tryCTFontCollectionSetQueryDescriptors(collection CTMutableFontCollectionRef, descriptors corefoundation.CFArrayRef) error {
	if _cTFontCollectionSetQueryDescriptors == nil {
		return symbolCallError("CTFontCollectionSetQueryDescriptors", "10.7", _cTFontCollectionSetQueryDescriptorsErr)
	}
	_cTFontCollectionSetQueryDescriptors(collection, descriptors)
	return nil
}

// CTFontCollectionSetQueryDescriptors replaces the array of descriptors for font matching.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCollectionSetQueryDescriptors(_:_:)
func CTFontCollectionSetQueryDescriptors(collection CTMutableFontCollectionRef, descriptors corefoundation.CFArrayRef) {
	if callErr := tryCTFontCollectionSetQueryDescriptors(collection, descriptors); callErr != nil {
		panic(callErr)
	}
}

var _cTFontCopyAttribute func(font CTFontRef, attribute corefoundation.CFStringRef) corefoundation.CFTypeRef
var _cTFontCopyAttributeErr error

func tryCTFontCopyAttribute(font CTFontRef, attribute corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _cTFontCopyAttribute == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("CTFontCopyAttribute", "10.5", _cTFontCopyAttributeErr)
	}
	return _cTFontCopyAttribute(font, attribute), nil
}

// CTFontCopyAttribute returns the value associated with an arbitrary attribute of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyAttribute(_:_:)
func CTFontCopyAttribute(font CTFontRef, attribute corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryCTFontCopyAttribute(font, attribute)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyAvailableTables func(font CTFontRef, options CTFontTableOptions) corefoundation.CFArrayRef
var _cTFontCopyAvailableTablesErr error

func tryCTFontCopyAvailableTables(font CTFontRef, options CTFontTableOptions) (corefoundation.CFArrayRef, error) {
	if _cTFontCopyAvailableTables == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCopyAvailableTables", "10.5", _cTFontCopyAvailableTablesErr)
	}
	return _cTFontCopyAvailableTables(font, options), nil
}

// CTFontCopyAvailableTables returns an array of font table tags.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyAvailableTables(_:_:)
func CTFontCopyAvailableTables(font CTFontRef, options CTFontTableOptions) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCopyAvailableTables(font, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyCharacterSet func(font CTFontRef) corefoundation.CFCharacterSet
var _cTFontCopyCharacterSetErr error

func tryCTFontCopyCharacterSet(font CTFontRef) (corefoundation.CFCharacterSet, error) {
	if _cTFontCopyCharacterSet == nil {
		return *new(corefoundation.CFCharacterSet), symbolCallError("CTFontCopyCharacterSet", "10.5", _cTFontCopyCharacterSetErr)
	}
	return _cTFontCopyCharacterSet(font), nil
}

// CTFontCopyCharacterSet returns the Unicode character set of the font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyCharacterSet(_:)
func CTFontCopyCharacterSet(font CTFontRef) corefoundation.CFCharacterSet {
	result, callErr := tryCTFontCopyCharacterSet(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyDefaultCascadeListForLanguages func(font CTFontRef, languagePrefList corefoundation.CFArrayRef) corefoundation.CFArrayRef
var _cTFontCopyDefaultCascadeListForLanguagesErr error

func tryCTFontCopyDefaultCascadeListForLanguages(font CTFontRef, languagePrefList corefoundation.CFArrayRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCopyDefaultCascadeListForLanguages == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCopyDefaultCascadeListForLanguages", "10.8", _cTFontCopyDefaultCascadeListForLanguagesErr)
	}
	return _cTFontCopyDefaultCascadeListForLanguages(font, languagePrefList), nil
}

// CTFontCopyDefaultCascadeListForLanguages retrieves an ordered list of font substitution preferences.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyDefaultCascadeListForLanguages(_:_:)
func CTFontCopyDefaultCascadeListForLanguages(font CTFontRef, languagePrefList corefoundation.CFArrayRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCopyDefaultCascadeListForLanguages(font, languagePrefList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyDisplayName func(font CTFontRef) corefoundation.CFStringRef
var _cTFontCopyDisplayNameErr error

func tryCTFontCopyDisplayName(font CTFontRef) (corefoundation.CFStringRef, error) {
	if _cTFontCopyDisplayName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyDisplayName", "10.5", _cTFontCopyDisplayNameErr)
	}
	return _cTFontCopyDisplayName(font), nil
}

// CTFontCopyDisplayName returns the display name of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyDisplayName(_:)
func CTFontCopyDisplayName(font CTFontRef) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyDisplayName(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyFamilyName func(font CTFontRef) corefoundation.CFStringRef
var _cTFontCopyFamilyNameErr error

func tryCTFontCopyFamilyName(font CTFontRef) (corefoundation.CFStringRef, error) {
	if _cTFontCopyFamilyName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyFamilyName", "10.5", _cTFontCopyFamilyNameErr)
	}
	return _cTFontCopyFamilyName(font), nil
}

// CTFontCopyFamilyName returns the family name of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyFamilyName(_:)
func CTFontCopyFamilyName(font CTFontRef) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyFamilyName(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyFeatureSettings func(font CTFontRef) corefoundation.CFArrayRef
var _cTFontCopyFeatureSettingsErr error

func tryCTFontCopyFeatureSettings(font CTFontRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCopyFeatureSettings == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCopyFeatureSettings", "10.5", _cTFontCopyFeatureSettingsErr)
	}
	return _cTFontCopyFeatureSettings(font), nil
}

// CTFontCopyFeatureSettings returns an array of font feature-setting tuples.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyFeatureSettings(_:)
func CTFontCopyFeatureSettings(font CTFontRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCopyFeatureSettings(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyFeatures func(font CTFontRef) corefoundation.CFArrayRef
var _cTFontCopyFeaturesErr error

func tryCTFontCopyFeatures(font CTFontRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCopyFeatures == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCopyFeatures", "10.5", _cTFontCopyFeaturesErr)
	}
	return _cTFontCopyFeatures(font), nil
}

// CTFontCopyFeatures returns an array of font features.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyFeatures(_:)
func CTFontCopyFeatures(font CTFontRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCopyFeatures(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyFontDescriptor func(font CTFontRef) CTFontDescriptorRef
var _cTFontCopyFontDescriptorErr error

func tryCTFontCopyFontDescriptor(font CTFontRef) (CTFontDescriptorRef, error) {
	if _cTFontCopyFontDescriptor == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontCopyFontDescriptor", "10.5", _cTFontCopyFontDescriptorErr)
	}
	return _cTFontCopyFontDescriptor(font), nil
}

// CTFontCopyFontDescriptor returns the normalized font descriptor for the given font reference.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyFontDescriptor(_:)
func CTFontCopyFontDescriptor(font CTFontRef) CTFontDescriptorRef {
	result, callErr := tryCTFontCopyFontDescriptor(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyFullName func(font CTFontRef) corefoundation.CFStringRef
var _cTFontCopyFullNameErr error

func tryCTFontCopyFullName(font CTFontRef) (corefoundation.CFStringRef, error) {
	if _cTFontCopyFullName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyFullName", "10.5", _cTFontCopyFullNameErr)
	}
	return _cTFontCopyFullName(font), nil
}

// CTFontCopyFullName returns the full name of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyFullName(_:)
func CTFontCopyFullName(font CTFontRef) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyFullName(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyGraphicsFont func(font CTFontRef, attributes *CTFontDescriptorRef) coregraphics.CGFont
var _cTFontCopyGraphicsFontErr error

func tryCTFontCopyGraphicsFont(font CTFontRef, attributes *CTFontDescriptorRef) (coregraphics.CGFont, error) {
	if _cTFontCopyGraphicsFont == nil {
		return *new(coregraphics.CGFont), symbolCallError("CTFontCopyGraphicsFont", "10.5", _cTFontCopyGraphicsFontErr)
	}
	return _cTFontCopyGraphicsFont(font, attributes), nil
}

// CTFontCopyGraphicsFont returns a Core Graphics font reference and attributes.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyGraphicsFont(_:_:)
func CTFontCopyGraphicsFont(font CTFontRef, attributes *CTFontDescriptorRef) coregraphics.CGFont {
	result, callErr := tryCTFontCopyGraphicsFont(font, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyLocalizedName func(font CTFontRef, nameKey corefoundation.CFStringRef, actualLanguage *corefoundation.CFStringRef) corefoundation.CFStringRef
var _cTFontCopyLocalizedNameErr error

func tryCTFontCopyLocalizedName(font CTFontRef, nameKey corefoundation.CFStringRef, actualLanguage *corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _cTFontCopyLocalizedName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyLocalizedName", "10.5", _cTFontCopyLocalizedNameErr)
	}
	return _cTFontCopyLocalizedName(font, nameKey, actualLanguage), nil
}

// CTFontCopyLocalizedName returns a reference to a localized name for the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyLocalizedName(_:_:_:)
func CTFontCopyLocalizedName(font CTFontRef, nameKey corefoundation.CFStringRef, actualLanguage *corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyLocalizedName(font, nameKey, actualLanguage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyName func(font CTFontRef, nameKey corefoundation.CFStringRef) corefoundation.CFStringRef
var _cTFontCopyNameErr error

func tryCTFontCopyName(font CTFontRef, nameKey corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _cTFontCopyName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyName", "10.5", _cTFontCopyNameErr)
	}
	return _cTFontCopyName(font, nameKey), nil
}

// CTFontCopyName returns a reference to the requested name of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyName(_:_:)
func CTFontCopyName(font CTFontRef, nameKey corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyName(font, nameKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyNameForGlyph func(font CTFontRef, glyph uint16) corefoundation.CFStringRef
var _cTFontCopyNameForGlyphErr error

func tryCTFontCopyNameForGlyph(font CTFontRef, glyph uint16) (corefoundation.CFStringRef, error) {
	if _cTFontCopyNameForGlyph == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyNameForGlyph", "10.8", _cTFontCopyNameForGlyphErr)
	}
	return _cTFontCopyNameForGlyph(font, glyph), nil
}

// CTFontCopyNameForGlyph retrieves the name for the specified glyph.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyNameForGlyph(_:_:)
func CTFontCopyNameForGlyph(font CTFontRef, glyph uint16) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyNameForGlyph(font, glyph)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyPostScriptName func(font CTFontRef) corefoundation.CFStringRef
var _cTFontCopyPostScriptNameErr error

func tryCTFontCopyPostScriptName(font CTFontRef) (corefoundation.CFStringRef, error) {
	if _cTFontCopyPostScriptName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTFontCopyPostScriptName", "10.5", _cTFontCopyPostScriptNameErr)
	}
	return _cTFontCopyPostScriptName(font), nil
}

// CTFontCopyPostScriptName returns the PostScript name of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyPostScriptName(_:)
func CTFontCopyPostScriptName(font CTFontRef) corefoundation.CFStringRef {
	result, callErr := tryCTFontCopyPostScriptName(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopySupportedLanguages func(font CTFontRef) corefoundation.CFArrayRef
var _cTFontCopySupportedLanguagesErr error

func tryCTFontCopySupportedLanguages(font CTFontRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCopySupportedLanguages == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCopySupportedLanguages", "10.5", _cTFontCopySupportedLanguagesErr)
	}
	return _cTFontCopySupportedLanguages(font), nil
}

// CTFontCopySupportedLanguages returns an array of languages supported by the font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopySupportedLanguages(_:)
func CTFontCopySupportedLanguages(font CTFontRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCopySupportedLanguages(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyTable func(font CTFontRef, table CTFontTableTag, options CTFontTableOptions) corefoundation.CFDataRef
var _cTFontCopyTableErr error

func tryCTFontCopyTable(font CTFontRef, table CTFontTableTag, options CTFontTableOptions) (corefoundation.CFDataRef, error) {
	if _cTFontCopyTable == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("CTFontCopyTable", "10.5", _cTFontCopyTableErr)
	}
	return _cTFontCopyTable(font, table, options), nil
}

// CTFontCopyTable returns a reference to the font table data.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyTable(_:_:_:)
func CTFontCopyTable(font CTFontRef, table CTFontTableTag, options CTFontTableOptions) corefoundation.CFDataRef {
	result, callErr := tryCTFontCopyTable(font, table, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyTraits func(font CTFontRef) corefoundation.CFDictionaryRef
var _cTFontCopyTraitsErr error

func tryCTFontCopyTraits(font CTFontRef) (corefoundation.CFDictionaryRef, error) {
	if _cTFontCopyTraits == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CTFontCopyTraits", "10.5", _cTFontCopyTraitsErr)
	}
	return _cTFontCopyTraits(font), nil
}

// CTFontCopyTraits returns the traits dictionary of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyTraits(_:)
func CTFontCopyTraits(font CTFontRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCTFontCopyTraits(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyVariation func(font CTFontRef) corefoundation.CFDictionaryRef
var _cTFontCopyVariationErr error

func tryCTFontCopyVariation(font CTFontRef) (corefoundation.CFDictionaryRef, error) {
	if _cTFontCopyVariation == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CTFontCopyVariation", "10.5", _cTFontCopyVariationErr)
	}
	return _cTFontCopyVariation(font), nil
}

// CTFontCopyVariation returns a variation dictionary from the font reference.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyVariation(_:)
func CTFontCopyVariation(font CTFontRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCTFontCopyVariation(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCopyVariationAxes func(font CTFontRef) corefoundation.CFArrayRef
var _cTFontCopyVariationAxesErr error

func tryCTFontCopyVariationAxes(font CTFontRef) (corefoundation.CFArrayRef, error) {
	if _cTFontCopyVariationAxes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontCopyVariationAxes", "10.5", _cTFontCopyVariationAxesErr)
	}
	return _cTFontCopyVariationAxes(font), nil
}

// CTFontCopyVariationAxes returns an array of variation axes.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCopyVariationAxes(_:)
func CTFontCopyVariationAxes(font CTFontRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontCopyVariationAxes(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateCopyWithAttributes func(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, attributes CTFontDescriptorRef) CTFontRef
var _cTFontCreateCopyWithAttributesErr error

func tryCTFontCreateCopyWithAttributes(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, attributes CTFontDescriptorRef) (CTFontRef, error) {
	if _cTFontCreateCopyWithAttributes == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateCopyWithAttributes", "10.5", _cTFontCreateCopyWithAttributesErr)
	}
	return _cTFontCreateCopyWithAttributes(font, size, matrix, attributes), nil
}

// CTFontCreateCopyWithAttributes returns a new font with additional attributes based on the original font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithAttributes(_:_:_:_:)
func CTFontCreateCopyWithAttributes(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, attributes CTFontDescriptorRef) CTFontRef {
	result, callErr := tryCTFontCreateCopyWithAttributes(font, size, matrix, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateCopyWithFamily func(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, family corefoundation.CFStringRef) CTFontRef
var _cTFontCreateCopyWithFamilyErr error

func tryCTFontCreateCopyWithFamily(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, family corefoundation.CFStringRef) (CTFontRef, error) {
	if _cTFontCreateCopyWithFamily == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateCopyWithFamily", "10.5", _cTFontCreateCopyWithFamilyErr)
	}
	return _cTFontCreateCopyWithFamily(font, size, matrix, family), nil
}

// CTFontCreateCopyWithFamily returns a new font in the specified family based on the traits of the original font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithFamily(_:_:_:_:)
func CTFontCreateCopyWithFamily(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, family corefoundation.CFStringRef) CTFontRef {
	result, callErr := tryCTFontCreateCopyWithFamily(font, size, matrix, family)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateCopyWithSymbolicTraits func(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, symTraitValue CTFontSymbolicTraits, symTraitMask CTFontSymbolicTraits) CTFontRef
var _cTFontCreateCopyWithSymbolicTraitsErr error

func tryCTFontCreateCopyWithSymbolicTraits(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, symTraitValue CTFontSymbolicTraits, symTraitMask CTFontSymbolicTraits) (CTFontRef, error) {
	if _cTFontCreateCopyWithSymbolicTraits == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateCopyWithSymbolicTraits", "10.5", _cTFontCreateCopyWithSymbolicTraitsErr)
	}
	return _cTFontCreateCopyWithSymbolicTraits(font, size, matrix, symTraitValue, symTraitMask), nil
}

// CTFontCreateCopyWithSymbolicTraits returns a new font in the same font family as the original with the specified symbolic traits.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateCopyWithSymbolicTraits(_:_:_:_:_:)
func CTFontCreateCopyWithSymbolicTraits(font CTFontRef, size float64, matrix *corefoundation.CGAffineTransform, symTraitValue CTFontSymbolicTraits, symTraitMask CTFontSymbolicTraits) CTFontRef {
	result, callErr := tryCTFontCreateCopyWithSymbolicTraits(font, size, matrix, symTraitValue, symTraitMask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateForString func(currentFont CTFontRef, string_ corefoundation.CFStringRef, range_ corefoundation.CFRange) CTFontRef
var _cTFontCreateForStringErr error

func tryCTFontCreateForString(currentFont CTFontRef, string_ corefoundation.CFStringRef, range_ corefoundation.CFRange) (CTFontRef, error) {
	if _cTFontCreateForString == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateForString", "10.5", _cTFontCreateForStringErr)
	}
	return _cTFontCreateForString(currentFont, string_, range_), nil
}

// CTFontCreateForString returns a font reference that most accurately maps the string range based on the current font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateForString(_:_:_:)
func CTFontCreateForString(currentFont CTFontRef, string_ corefoundation.CFStringRef, range_ corefoundation.CFRange) CTFontRef {
	result, callErr := tryCTFontCreateForString(currentFont, string_, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateForStringWithLanguage func(currentFont CTFontRef, string_ corefoundation.CFStringRef, range_ corefoundation.CFRange, language corefoundation.CFStringRef) CTFontRef
var _cTFontCreateForStringWithLanguageErr error

func tryCTFontCreateForStringWithLanguage(currentFont CTFontRef, string_ corefoundation.CFStringRef, range_ corefoundation.CFRange, language corefoundation.CFStringRef) (CTFontRef, error) {
	if _cTFontCreateForStringWithLanguage == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateForStringWithLanguage", "10.9", _cTFontCreateForStringWithLanguageErr)
	}
	return _cTFontCreateForStringWithLanguage(currentFont, string_, range_, language), nil
}

// CTFontCreateForStringWithLanguage returns a font reference that most accurately maps the string range based on the current font and language.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateForStringWithLanguage(_:_:_:_:)
func CTFontCreateForStringWithLanguage(currentFont CTFontRef, string_ corefoundation.CFStringRef, range_ corefoundation.CFRange, language corefoundation.CFStringRef) CTFontRef {
	result, callErr := tryCTFontCreateForStringWithLanguage(currentFont, string_, range_, language)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreatePathForGlyph func(font CTFontRef, glyph uint16, matrix *corefoundation.CGAffineTransform) coregraphics.CGPathRef
var _cTFontCreatePathForGlyphErr error

func tryCTFontCreatePathForGlyph(font CTFontRef, glyph uint16, matrix *corefoundation.CGAffineTransform) (coregraphics.CGPathRef, error) {
	if _cTFontCreatePathForGlyph == nil {
		return *new(coregraphics.CGPathRef), symbolCallError("CTFontCreatePathForGlyph", "10.5", _cTFontCreatePathForGlyphErr)
	}
	return _cTFontCreatePathForGlyph(font, glyph, matrix), nil
}

// CTFontCreatePathForGlyph creates a path for the specified glyph.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreatePathForGlyph(_:_:_:)
func CTFontCreatePathForGlyph(font CTFontRef, glyph uint16, matrix *corefoundation.CGAffineTransform) coregraphics.CGPathRef {
	result, callErr := tryCTFontCreatePathForGlyph(font, glyph, matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateUIFontForLanguage func(uiType CTFontUIFontType, size float64, language corefoundation.CFStringRef) CTFontRef
var _cTFontCreateUIFontForLanguageErr error

func tryCTFontCreateUIFontForLanguage(uiType CTFontUIFontType, size float64, language corefoundation.CFStringRef) (CTFontRef, error) {
	if _cTFontCreateUIFontForLanguage == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateUIFontForLanguage", "10.5", _cTFontCreateUIFontForLanguageErr)
	}
	return _cTFontCreateUIFontForLanguage(uiType, size, language), nil
}

// CTFontCreateUIFontForLanguage returns the special user-interface font for the given language and user-interface type.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateUIFontForLanguage(_:_:_:)
func CTFontCreateUIFontForLanguage(uiType CTFontUIFontType, size float64, language corefoundation.CFStringRef) CTFontRef {
	result, callErr := tryCTFontCreateUIFontForLanguage(uiType, size, language)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateWithFontDescriptor func(descriptor CTFontDescriptorRef, size float64, matrix *corefoundation.CGAffineTransform) CTFontRef
var _cTFontCreateWithFontDescriptorErr error

func tryCTFontCreateWithFontDescriptor(descriptor CTFontDescriptorRef, size float64, matrix *corefoundation.CGAffineTransform) (CTFontRef, error) {
	if _cTFontCreateWithFontDescriptor == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateWithFontDescriptor", "10.5", _cTFontCreateWithFontDescriptorErr)
	}
	return _cTFontCreateWithFontDescriptor(descriptor, size, matrix), nil
}

// CTFontCreateWithFontDescriptor returns a new font reference that best matches the given font descriptor.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateWithFontDescriptor(_:_:_:)
func CTFontCreateWithFontDescriptor(descriptor CTFontDescriptorRef, size float64, matrix *corefoundation.CGAffineTransform) CTFontRef {
	result, callErr := tryCTFontCreateWithFontDescriptor(descriptor, size, matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateWithFontDescriptorAndOptions func(descriptor CTFontDescriptorRef, size float64, matrix *corefoundation.CGAffineTransform, options CTFontOptions) CTFontRef
var _cTFontCreateWithFontDescriptorAndOptionsErr error

func tryCTFontCreateWithFontDescriptorAndOptions(descriptor CTFontDescriptorRef, size float64, matrix *corefoundation.CGAffineTransform, options CTFontOptions) (CTFontRef, error) {
	if _cTFontCreateWithFontDescriptorAndOptions == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateWithFontDescriptorAndOptions", "10.6", _cTFontCreateWithFontDescriptorAndOptionsErr)
	}
	return _cTFontCreateWithFontDescriptorAndOptions(descriptor, size, matrix, options), nil
}

// CTFontCreateWithFontDescriptorAndOptions returns a new font reference that best matches the given font descriptor.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateWithFontDescriptorAndOptions(_:_:_:_:)
func CTFontCreateWithFontDescriptorAndOptions(descriptor CTFontDescriptorRef, size float64, matrix *corefoundation.CGAffineTransform, options CTFontOptions) CTFontRef {
	result, callErr := tryCTFontCreateWithFontDescriptorAndOptions(descriptor, size, matrix, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateWithGraphicsFont func(graphicsFont coregraphics.CGFont, size float64, matrix *corefoundation.CGAffineTransform, attributes CTFontDescriptorRef) CTFontRef
var _cTFontCreateWithGraphicsFontErr error

func tryCTFontCreateWithGraphicsFont(graphicsFont coregraphics.CGFont, size float64, matrix *corefoundation.CGAffineTransform, attributes CTFontDescriptorRef) (CTFontRef, error) {
	if _cTFontCreateWithGraphicsFont == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateWithGraphicsFont", "10.5", _cTFontCreateWithGraphicsFontErr)
	}
	return _cTFontCreateWithGraphicsFont(graphicsFont, size, matrix, attributes), nil
}

// CTFontCreateWithGraphicsFont creates a new font reference from an existing Core Graphics font reference.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateWithGraphicsFont(_:_:_:_:)
func CTFontCreateWithGraphicsFont(graphicsFont coregraphics.CGFont, size float64, matrix *corefoundation.CGAffineTransform, attributes CTFontDescriptorRef) CTFontRef {
	result, callErr := tryCTFontCreateWithGraphicsFont(graphicsFont, size, matrix, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateWithName func(name corefoundation.CFStringRef, size float64, matrix *corefoundation.CGAffineTransform) CTFontRef
var _cTFontCreateWithNameErr error

func tryCTFontCreateWithName(name corefoundation.CFStringRef, size float64, matrix *corefoundation.CGAffineTransform) (CTFontRef, error) {
	if _cTFontCreateWithName == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateWithName", "10.5", _cTFontCreateWithNameErr)
	}
	return _cTFontCreateWithName(name, size, matrix), nil
}

// CTFontCreateWithName returns a new font reference for the given name.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateWithName(_:_:_:)
func CTFontCreateWithName(name corefoundation.CFStringRef, size float64, matrix *corefoundation.CGAffineTransform) CTFontRef {
	result, callErr := tryCTFontCreateWithName(name, size, matrix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontCreateWithNameAndOptions func(name corefoundation.CFStringRef, size float64, matrix *corefoundation.CGAffineTransform, options CTFontOptions) CTFontRef
var _cTFontCreateWithNameAndOptionsErr error

func tryCTFontCreateWithNameAndOptions(name corefoundation.CFStringRef, size float64, matrix *corefoundation.CGAffineTransform, options CTFontOptions) (CTFontRef, error) {
	if _cTFontCreateWithNameAndOptions == nil {
		return *new(CTFontRef), symbolCallError("CTFontCreateWithNameAndOptions", "10.6", _cTFontCreateWithNameAndOptionsErr)
	}
	return _cTFontCreateWithNameAndOptions(name, size, matrix, options), nil
}

// CTFontCreateWithNameAndOptions returns a new font reference for the given name.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontCreateWithNameAndOptions(_:_:_:_:)
func CTFontCreateWithNameAndOptions(name corefoundation.CFStringRef, size float64, matrix *corefoundation.CGAffineTransform, options CTFontOptions) CTFontRef {
	result, callErr := tryCTFontCreateWithNameAndOptions(name, size, matrix, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCopyAttribute func(descriptor CTFontDescriptorRef, attribute corefoundation.CFStringRef) corefoundation.CFTypeRef
var _cTFontDescriptorCopyAttributeErr error

func tryCTFontDescriptorCopyAttribute(descriptor CTFontDescriptorRef, attribute corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _cTFontDescriptorCopyAttribute == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("CTFontDescriptorCopyAttribute", "10.5", _cTFontDescriptorCopyAttributeErr)
	}
	return _cTFontDescriptorCopyAttribute(descriptor, attribute), nil
}

// CTFontDescriptorCopyAttribute returns the value associated with an arbitrary attribute.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyAttribute(_:_:)
func CTFontDescriptorCopyAttribute(descriptor CTFontDescriptorRef, attribute corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryCTFontDescriptorCopyAttribute(descriptor, attribute)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCopyAttributes func(descriptor CTFontDescriptorRef) corefoundation.CFDictionaryRef
var _cTFontDescriptorCopyAttributesErr error

func tryCTFontDescriptorCopyAttributes(descriptor CTFontDescriptorRef) (corefoundation.CFDictionaryRef, error) {
	if _cTFontDescriptorCopyAttributes == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CTFontDescriptorCopyAttributes", "10.5", _cTFontDescriptorCopyAttributesErr)
	}
	return _cTFontDescriptorCopyAttributes(descriptor), nil
}

// CTFontDescriptorCopyAttributes returns the attributes dictionary of the font descriptor.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyAttributes(_:)
func CTFontDescriptorCopyAttributes(descriptor CTFontDescriptorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCTFontDescriptorCopyAttributes(descriptor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCopyLocalizedAttribute func(descriptor CTFontDescriptorRef, attribute corefoundation.CFStringRef, language *corefoundation.CFStringRef) corefoundation.CFTypeRef
var _cTFontDescriptorCopyLocalizedAttributeErr error

func tryCTFontDescriptorCopyLocalizedAttribute(descriptor CTFontDescriptorRef, attribute corefoundation.CFStringRef, language *corefoundation.CFStringRef) (corefoundation.CFTypeRef, error) {
	if _cTFontDescriptorCopyLocalizedAttribute == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("CTFontDescriptorCopyLocalizedAttribute", "10.5", _cTFontDescriptorCopyLocalizedAttributeErr)
	}
	return _cTFontDescriptorCopyLocalizedAttribute(descriptor, attribute, language), nil
}

// CTFontDescriptorCopyLocalizedAttribute returns a localized value for the requested attribute, if available.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCopyLocalizedAttribute(_:_:_:)
func CTFontDescriptorCopyLocalizedAttribute(descriptor CTFontDescriptorRef, attribute corefoundation.CFStringRef, language *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	result, callErr := tryCTFontDescriptorCopyLocalizedAttribute(descriptor, attribute, language)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateCopyWithAttributes func(original CTFontDescriptorRef, attributes corefoundation.CFDictionaryRef) CTFontDescriptorRef
var _cTFontDescriptorCreateCopyWithAttributesErr error

func tryCTFontDescriptorCreateCopyWithAttributes(original CTFontDescriptorRef, attributes corefoundation.CFDictionaryRef) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateCopyWithAttributes == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateCopyWithAttributes", "10.5", _cTFontDescriptorCreateCopyWithAttributesErr)
	}
	return _cTFontDescriptorCreateCopyWithAttributes(original, attributes), nil
}

// CTFontDescriptorCreateCopyWithAttributes creates a copy of the original font descriptor with new attributes.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithAttributes(_:_:)
func CTFontDescriptorCreateCopyWithAttributes(original CTFontDescriptorRef, attributes corefoundation.CFDictionaryRef) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateCopyWithAttributes(original, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateCopyWithFamily func(original CTFontDescriptorRef, family corefoundation.CFStringRef) CTFontDescriptorRef
var _cTFontDescriptorCreateCopyWithFamilyErr error

func tryCTFontDescriptorCreateCopyWithFamily(original CTFontDescriptorRef, family corefoundation.CFStringRef) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateCopyWithFamily == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateCopyWithFamily", "10.9", _cTFontDescriptorCreateCopyWithFamilyErr)
	}
	return _cTFontDescriptorCreateCopyWithFamily(original, family), nil
}

// CTFontDescriptorCreateCopyWithFamily creates a copy of the font descriptor in the specified family based on the traits of the original.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithFamily(_:_:)
func CTFontDescriptorCreateCopyWithFamily(original CTFontDescriptorRef, family corefoundation.CFStringRef) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateCopyWithFamily(original, family)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateCopyWithFeature func(original CTFontDescriptorRef, featureTypeIdentifier corefoundation.CFNumberRef, featureSelectorIdentifier corefoundation.CFNumberRef) CTFontDescriptorRef
var _cTFontDescriptorCreateCopyWithFeatureErr error

func tryCTFontDescriptorCreateCopyWithFeature(original CTFontDescriptorRef, featureTypeIdentifier corefoundation.CFNumberRef, featureSelectorIdentifier corefoundation.CFNumberRef) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateCopyWithFeature == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateCopyWithFeature", "10.5", _cTFontDescriptorCreateCopyWithFeatureErr)
	}
	return _cTFontDescriptorCreateCopyWithFeature(original, featureTypeIdentifier, featureSelectorIdentifier), nil
}

// CTFontDescriptorCreateCopyWithFeature copies a font descriptor with new feature settings.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithFeature(_:_:_:)
func CTFontDescriptorCreateCopyWithFeature(original CTFontDescriptorRef, featureTypeIdentifier corefoundation.CFNumberRef, featureSelectorIdentifier corefoundation.CFNumberRef) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateCopyWithFeature(original, featureTypeIdentifier, featureSelectorIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateCopyWithSymbolicTraits func(original CTFontDescriptorRef, symTraitValue CTFontSymbolicTraits, symTraitMask CTFontSymbolicTraits) CTFontDescriptorRef
var _cTFontDescriptorCreateCopyWithSymbolicTraitsErr error

func tryCTFontDescriptorCreateCopyWithSymbolicTraits(original CTFontDescriptorRef, symTraitValue CTFontSymbolicTraits, symTraitMask CTFontSymbolicTraits) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateCopyWithSymbolicTraits == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateCopyWithSymbolicTraits", "10.9", _cTFontDescriptorCreateCopyWithSymbolicTraitsErr)
	}
	return _cTFontDescriptorCreateCopyWithSymbolicTraits(original, symTraitValue, symTraitMask), nil
}

// CTFontDescriptorCreateCopyWithSymbolicTraits creates a copy of the font descriptor with the specified symbolic traits as the original.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithSymbolicTraits(_:_:_:)
func CTFontDescriptorCreateCopyWithSymbolicTraits(original CTFontDescriptorRef, symTraitValue CTFontSymbolicTraits, symTraitMask CTFontSymbolicTraits) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateCopyWithSymbolicTraits(original, symTraitValue, symTraitMask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateCopyWithVariation func(original CTFontDescriptorRef, variationIdentifier corefoundation.CFNumberRef, variationValue float64) CTFontDescriptorRef
var _cTFontDescriptorCreateCopyWithVariationErr error

func tryCTFontDescriptorCreateCopyWithVariation(original CTFontDescriptorRef, variationIdentifier corefoundation.CFNumberRef, variationValue float64) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateCopyWithVariation == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateCopyWithVariation", "10.5", _cTFontDescriptorCreateCopyWithVariationErr)
	}
	return _cTFontDescriptorCreateCopyWithVariation(original, variationIdentifier, variationValue), nil
}

// CTFontDescriptorCreateCopyWithVariation creates a copy of the original font descriptor with a new variation instance.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateCopyWithVariation(_:_:_:)
func CTFontDescriptorCreateCopyWithVariation(original CTFontDescriptorRef, variationIdentifier corefoundation.CFNumberRef, variationValue float64) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateCopyWithVariation(original, variationIdentifier, variationValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateMatchingFontDescriptor func(descriptor CTFontDescriptorRef, mandatoryAttributes corefoundation.CFSetRef) CTFontDescriptorRef
var _cTFontDescriptorCreateMatchingFontDescriptorErr error

func tryCTFontDescriptorCreateMatchingFontDescriptor(descriptor CTFontDescriptorRef, mandatoryAttributes corefoundation.CFSetRef) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateMatchingFontDescriptor == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateMatchingFontDescriptor", "10.5", _cTFontDescriptorCreateMatchingFontDescriptorErr)
	}
	return _cTFontDescriptorCreateMatchingFontDescriptor(descriptor, mandatoryAttributes), nil
}

// CTFontDescriptorCreateMatchingFontDescriptor returns the single preferred matching font descriptor based on the original descriptor and system precedence.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateMatchingFontDescriptor(_:_:)
func CTFontDescriptorCreateMatchingFontDescriptor(descriptor CTFontDescriptorRef, mandatoryAttributes corefoundation.CFSetRef) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateMatchingFontDescriptor(descriptor, mandatoryAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateMatchingFontDescriptors func(descriptor CTFontDescriptorRef, mandatoryAttributes corefoundation.CFSetRef) corefoundation.CFArrayRef
var _cTFontDescriptorCreateMatchingFontDescriptorsErr error

func tryCTFontDescriptorCreateMatchingFontDescriptors(descriptor CTFontDescriptorRef, mandatoryAttributes corefoundation.CFSetRef) (corefoundation.CFArrayRef, error) {
	if _cTFontDescriptorCreateMatchingFontDescriptors == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontDescriptorCreateMatchingFontDescriptors", "10.5", _cTFontDescriptorCreateMatchingFontDescriptorsErr)
	}
	return _cTFontDescriptorCreateMatchingFontDescriptors(descriptor, mandatoryAttributes), nil
}

// CTFontDescriptorCreateMatchingFontDescriptors returns an array of normalized font descriptors matching the provided descriptor.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateMatchingFontDescriptors(_:_:)
func CTFontDescriptorCreateMatchingFontDescriptors(descriptor CTFontDescriptorRef, mandatoryAttributes corefoundation.CFSetRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontDescriptorCreateMatchingFontDescriptors(descriptor, mandatoryAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateWithAttributes func(attributes corefoundation.CFDictionaryRef) CTFontDescriptorRef
var _cTFontDescriptorCreateWithAttributesErr error

func tryCTFontDescriptorCreateWithAttributes(attributes corefoundation.CFDictionaryRef) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateWithAttributes == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateWithAttributes", "10.5", _cTFontDescriptorCreateWithAttributesErr)
	}
	return _cTFontDescriptorCreateWithAttributes(attributes), nil
}

// CTFontDescriptorCreateWithAttributes creates a new font descriptor reference from a dictionary of attributes.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateWithAttributes(_:)
func CTFontDescriptorCreateWithAttributes(attributes corefoundation.CFDictionaryRef) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateWithAttributes(attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorCreateWithNameAndSize func(name corefoundation.CFStringRef, size float64) CTFontDescriptorRef
var _cTFontDescriptorCreateWithNameAndSizeErr error

func tryCTFontDescriptorCreateWithNameAndSize(name corefoundation.CFStringRef, size float64) (CTFontDescriptorRef, error) {
	if _cTFontDescriptorCreateWithNameAndSize == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontDescriptorCreateWithNameAndSize", "10.5", _cTFontDescriptorCreateWithNameAndSizeErr)
	}
	return _cTFontDescriptorCreateWithNameAndSize(name, size), nil
}

// CTFontDescriptorCreateWithNameAndSize creates a new font descriptor with the provided PostScript name and size.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorCreateWithNameAndSize(_:_:)
func CTFontDescriptorCreateWithNameAndSize(name corefoundation.CFStringRef, size float64) CTFontDescriptorRef {
	result, callErr := tryCTFontDescriptorCreateWithNameAndSize(name, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorGetTypeID func() uint
var _cTFontDescriptorGetTypeIDErr error

func tryCTFontDescriptorGetTypeID() (uint, error) {
	if _cTFontDescriptorGetTypeID == nil {
		return 0, symbolCallError("CTFontDescriptorGetTypeID", "10.5", _cTFontDescriptorGetTypeIDErr)
	}
	return _cTFontDescriptorGetTypeID(), nil
}

// CTFontDescriptorGetTypeID returns the type identifier for Core Text font descriptor references.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorGetTypeID()
func CTFontDescriptorGetTypeID() uint {
	result, callErr := tryCTFontDescriptorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDescriptorMatchFontDescriptorsWithProgressHandler func(descriptors corefoundation.CFArrayRef, mandatoryAttributes corefoundation.CFSetRef, progressBlock unsafe.Pointer) bool
var _cTFontDescriptorMatchFontDescriptorsWithProgressHandlerErr error

func tryCTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors corefoundation.CFArrayRef, mandatoryAttributes corefoundation.CFSetRef, progressBlock CTFontDescriptorProgressHandler) (bool, error) {
	if _cTFontDescriptorMatchFontDescriptorsWithProgressHandler == nil {
		return false, symbolCallError("CTFontDescriptorMatchFontDescriptorsWithProgressHandler", "10.9", _cTFontDescriptorMatchFontDescriptorsWithProgressHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 CTFontDescriptorMatchingState, blockArg1 corefoundation.CFDictionaryRef) bool {
		return progressBlock(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors, mandatoryAttributes, _block0), nil
}

// CTFontDescriptorMatchFontDescriptorsWithProgressHandler matches font descriptors and tracks progress with a progress handler.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDescriptorMatchFontDescriptorsWithProgressHandler(_:_:_:)
func CTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors corefoundation.CFArrayRef, mandatoryAttributes corefoundation.CFSetRef, progressBlock CTFontDescriptorProgressHandler) bool {
	result, callErr := tryCTFontDescriptorMatchFontDescriptorsWithProgressHandler(descriptors, mandatoryAttributes, progressBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontDrawGlyphs func(font CTFontRef, glyphs unsafe.Pointer, positions *corefoundation.CGPoint, count uintptr, context coregraphics.CGContextRef)
var _cTFontDrawGlyphsErr error

func tryCTFontDrawGlyphs(font CTFontRef, glyphs unsafe.Pointer, positions *corefoundation.CGPoint, count uintptr, context coregraphics.CGContextRef) error {
	if _cTFontDrawGlyphs == nil {
		return symbolCallError("CTFontDrawGlyphs", "10.7", _cTFontDrawGlyphsErr)
	}
	_cTFontDrawGlyphs(font, glyphs, positions, count, context)
	return nil
}

// CTFontDrawGlyphs renders the given glyphs of a font at the specified positions in the supplied graphics context.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDrawGlyphs(_:_:_:_:_:)
func CTFontDrawGlyphs(font CTFontRef, glyphs unsafe.Pointer, positions *corefoundation.CGPoint, count uintptr, context coregraphics.CGContextRef) {
	if callErr := tryCTFontDrawGlyphs(font, glyphs, positions, count, context); callErr != nil {
		panic(callErr)
	}
}

var _cTFontDrawImageFromAdaptiveImageProviderAtPoint func(font CTFontRef, provider CTAdaptiveImageProvidingObject, point corefoundation.CGPoint, context coregraphics.CGContextRef)
var _cTFontDrawImageFromAdaptiveImageProviderAtPointErr error

func tryCTFontDrawImageFromAdaptiveImageProviderAtPoint(font CTFontRef, provider CTAdaptiveImageProvidingObject, point corefoundation.CGPoint, context coregraphics.CGContextRef) error {
	if _cTFontDrawImageFromAdaptiveImageProviderAtPoint == nil {
		return symbolCallError("CTFontDrawImageFromAdaptiveImageProviderAtPoint", "15.0", _cTFontDrawImageFromAdaptiveImageProviderAtPointErr)
	}
	_cTFontDrawImageFromAdaptiveImageProviderAtPoint(font, provider, point, context)
	return nil
}

// CTFontDrawImageFromAdaptiveImageProviderAtPoint.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontDrawImageFromAdaptiveImageProviderAtPoint(_:_:_:_:)
func CTFontDrawImageFromAdaptiveImageProviderAtPoint(font CTFontRef, provider CTAdaptiveImageProvidingObject, point corefoundation.CGPoint, context coregraphics.CGContextRef) {
	if callErr := tryCTFontDrawImageFromAdaptiveImageProviderAtPoint(font, provider, point, context); callErr != nil {
		panic(callErr)
	}
}

var _cTFontGetAdvancesForGlyphs func(font CTFontRef, orientation CTFontOrientation, glyphs unsafe.Pointer, advances *corefoundation.CGSize, count int) float64
var _cTFontGetAdvancesForGlyphsErr error

func tryCTFontGetAdvancesForGlyphs(font CTFontRef, orientation CTFontOrientation, glyphs unsafe.Pointer, advances *corefoundation.CGSize, count int) (float64, error) {
	if _cTFontGetAdvancesForGlyphs == nil {
		return 0.0, symbolCallError("CTFontGetAdvancesForGlyphs", "10.5", _cTFontGetAdvancesForGlyphsErr)
	}
	return _cTFontGetAdvancesForGlyphs(font, orientation, glyphs, advances, count), nil
}

// CTFontGetAdvancesForGlyphs calculates the advances for an array of glyphs and returns the summed advance.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetAdvancesForGlyphs(_:_:_:_:_:)
func CTFontGetAdvancesForGlyphs(font CTFontRef, orientation CTFontOrientation, glyphs unsafe.Pointer, advances *corefoundation.CGSize, count int) float64 {
	result, callErr := tryCTFontGetAdvancesForGlyphs(font, orientation, glyphs, advances, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetAscent func(font CTFontRef) float64
var _cTFontGetAscentErr error

func tryCTFontGetAscent(font CTFontRef) (float64, error) {
	if _cTFontGetAscent == nil {
		return 0.0, symbolCallError("CTFontGetAscent", "10.5", _cTFontGetAscentErr)
	}
	return _cTFontGetAscent(font), nil
}

// CTFontGetAscent returns the scaled font-ascent metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetAscent(_:)
func CTFontGetAscent(font CTFontRef) float64 {
	result, callErr := tryCTFontGetAscent(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetBoundingBox func(font CTFontRef) corefoundation.CGRect
var _cTFontGetBoundingBoxErr error

func tryCTFontGetBoundingBox(font CTFontRef) (corefoundation.CGRect, error) {
	if _cTFontGetBoundingBox == nil {
		return corefoundation.CGRect{}, symbolCallError("CTFontGetBoundingBox", "10.5", _cTFontGetBoundingBoxErr)
	}
	return _cTFontGetBoundingBox(font), nil
}

// CTFontGetBoundingBox returns the scaled bounding box of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetBoundingBox(_:)
func CTFontGetBoundingBox(font CTFontRef) corefoundation.CGRect {
	result, callErr := tryCTFontGetBoundingBox(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetBoundingRectsForGlyphs func(font CTFontRef, orientation CTFontOrientation, glyphs unsafe.Pointer, boundingRects *corefoundation.CGRect, count int) corefoundation.CGRect
var _cTFontGetBoundingRectsForGlyphsErr error

func tryCTFontGetBoundingRectsForGlyphs(font CTFontRef, orientation CTFontOrientation, glyphs unsafe.Pointer, boundingRects *corefoundation.CGRect, count int) (corefoundation.CGRect, error) {
	if _cTFontGetBoundingRectsForGlyphs == nil {
		return corefoundation.CGRect{}, symbolCallError("CTFontGetBoundingRectsForGlyphs", "10.5", _cTFontGetBoundingRectsForGlyphsErr)
	}
	return _cTFontGetBoundingRectsForGlyphs(font, orientation, glyphs, boundingRects, count), nil
}

// CTFontGetBoundingRectsForGlyphs calculates the bounding rects for an array of glyphs and returns the overall bounding rectangle for the glyph run.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetBoundingRectsForGlyphs(_:_:_:_:_:)
func CTFontGetBoundingRectsForGlyphs(font CTFontRef, orientation CTFontOrientation, glyphs unsafe.Pointer, boundingRects *corefoundation.CGRect, count int) corefoundation.CGRect {
	result, callErr := tryCTFontGetBoundingRectsForGlyphs(font, orientation, glyphs, boundingRects, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetCapHeight func(font CTFontRef) float64
var _cTFontGetCapHeightErr error

func tryCTFontGetCapHeight(font CTFontRef) (float64, error) {
	if _cTFontGetCapHeight == nil {
		return 0.0, symbolCallError("CTFontGetCapHeight", "10.5", _cTFontGetCapHeightErr)
	}
	return _cTFontGetCapHeight(font), nil
}

// CTFontGetCapHeight returns the cap-height metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetCapHeight(_:)
func CTFontGetCapHeight(font CTFontRef) float64 {
	result, callErr := tryCTFontGetCapHeight(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetDescent func(font CTFontRef) float64
var _cTFontGetDescentErr error

func tryCTFontGetDescent(font CTFontRef) (float64, error) {
	if _cTFontGetDescent == nil {
		return 0.0, symbolCallError("CTFontGetDescent", "10.5", _cTFontGetDescentErr)
	}
	return _cTFontGetDescent(font), nil
}

// CTFontGetDescent returns the scaled font-descent metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetDescent(_:)
func CTFontGetDescent(font CTFontRef) float64 {
	result, callErr := tryCTFontGetDescent(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetGlyphCount func(font CTFontRef) int
var _cTFontGetGlyphCountErr error

func tryCTFontGetGlyphCount(font CTFontRef) (int, error) {
	if _cTFontGetGlyphCount == nil {
		return 0, symbolCallError("CTFontGetGlyphCount", "10.5", _cTFontGetGlyphCountErr)
	}
	return _cTFontGetGlyphCount(font), nil
}

// CTFontGetGlyphCount returns the number of glyphs of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphCount(_:)
func CTFontGetGlyphCount(font CTFontRef) int {
	result, callErr := tryCTFontGetGlyphCount(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetGlyphWithName func(font CTFontRef, glyphName corefoundation.CFStringRef) uint16
var _cTFontGetGlyphWithNameErr error

func tryCTFontGetGlyphWithName(font CTFontRef, glyphName corefoundation.CFStringRef) (uint16, error) {
	if _cTFontGetGlyphWithName == nil {
		return 0, symbolCallError("CTFontGetGlyphWithName", "10.5", _cTFontGetGlyphWithNameErr)
	}
	return _cTFontGetGlyphWithName(font, glyphName), nil
}

// CTFontGetGlyphWithName returns the glyph for the specified name.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphWithName(_:_:)
func CTFontGetGlyphWithName(font CTFontRef, glyphName corefoundation.CFStringRef) uint16 {
	result, callErr := tryCTFontGetGlyphWithName(font, glyphName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetGlyphsForCharacters func(font CTFontRef, characters *uint16, glyphs unsafe.Pointer, count int) bool
var _cTFontGetGlyphsForCharactersErr error

func tryCTFontGetGlyphsForCharacters(font CTFontRef, characters *uint16, glyphs unsafe.Pointer, count int) (bool, error) {
	if _cTFontGetGlyphsForCharacters == nil {
		return false, symbolCallError("CTFontGetGlyphsForCharacters", "10.5", _cTFontGetGlyphsForCharactersErr)
	}
	return _cTFontGetGlyphsForCharacters(font, characters, glyphs, count), nil
}

// CTFontGetGlyphsForCharacters performs basic character-to-glyph mapping.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetGlyphsForCharacters(_:_:_:_:)
func CTFontGetGlyphsForCharacters(font CTFontRef, characters *uint16, glyphs unsafe.Pointer, count int) bool {
	result, callErr := tryCTFontGetGlyphsForCharacters(font, characters, glyphs, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetLeading func(font CTFontRef) float64
var _cTFontGetLeadingErr error

func tryCTFontGetLeading(font CTFontRef) (float64, error) {
	if _cTFontGetLeading == nil {
		return 0.0, symbolCallError("CTFontGetLeading", "10.5", _cTFontGetLeadingErr)
	}
	return _cTFontGetLeading(font), nil
}

// CTFontGetLeading returns the scaled font-leading metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetLeading(_:)
func CTFontGetLeading(font CTFontRef) float64 {
	result, callErr := tryCTFontGetLeading(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetLigatureCaretPositions func(font CTFontRef, glyph uint16, positions *float64, maxPositions int) int
var _cTFontGetLigatureCaretPositionsErr error

func tryCTFontGetLigatureCaretPositions(font CTFontRef, glyph uint16, positions *float64, maxPositions int) (int, error) {
	if _cTFontGetLigatureCaretPositions == nil {
		return 0, symbolCallError("CTFontGetLigatureCaretPositions", "10.5", _cTFontGetLigatureCaretPositionsErr)
	}
	return _cTFontGetLigatureCaretPositions(font, glyph, positions, maxPositions), nil
}

// CTFontGetLigatureCaretPositions returns caret positions within a glyph.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetLigatureCaretPositions(_:_:_:_:)
func CTFontGetLigatureCaretPositions(font CTFontRef, glyph uint16, positions *float64, maxPositions int) int {
	result, callErr := tryCTFontGetLigatureCaretPositions(font, glyph, positions, maxPositions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetMatrix func(font CTFontRef) corefoundation.CGAffineTransform
var _cTFontGetMatrixErr error

func tryCTFontGetMatrix(font CTFontRef) (corefoundation.CGAffineTransform, error) {
	if _cTFontGetMatrix == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CTFontGetMatrix", "10.5", _cTFontGetMatrixErr)
	}
	return _cTFontGetMatrix(font), nil
}

// CTFontGetMatrix returns the transformation matrix of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetMatrix(_:)
func CTFontGetMatrix(font CTFontRef) corefoundation.CGAffineTransform {
	result, callErr := tryCTFontGetMatrix(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetOpticalBoundsForGlyphs func(font CTFontRef, glyphs unsafe.Pointer, boundingRects *corefoundation.CGRect, count int, options uint) corefoundation.CGRect
var _cTFontGetOpticalBoundsForGlyphsErr error

func tryCTFontGetOpticalBoundsForGlyphs(font CTFontRef, glyphs unsafe.Pointer, boundingRects *corefoundation.CGRect, count int, options uint) (corefoundation.CGRect, error) {
	if _cTFontGetOpticalBoundsForGlyphs == nil {
		return corefoundation.CGRect{}, symbolCallError("CTFontGetOpticalBoundsForGlyphs", "10.8", _cTFontGetOpticalBoundsForGlyphsErr)
	}
	return _cTFontGetOpticalBoundsForGlyphs(font, glyphs, boundingRects, count, options), nil
}

// CTFontGetOpticalBoundsForGlyphs calculates the optical bounds for an array of glyphs and returns the overall optical bounds for the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetOpticalBoundsForGlyphs(_:_:_:_:_:)
func CTFontGetOpticalBoundsForGlyphs(font CTFontRef, glyphs unsafe.Pointer, boundingRects *corefoundation.CGRect, count int, options uint) corefoundation.CGRect {
	result, callErr := tryCTFontGetOpticalBoundsForGlyphs(font, glyphs, boundingRects, count, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetSize func(font CTFontRef) float64
var _cTFontGetSizeErr error

func tryCTFontGetSize(font CTFontRef) (float64, error) {
	if _cTFontGetSize == nil {
		return 0.0, symbolCallError("CTFontGetSize", "10.5", _cTFontGetSizeErr)
	}
	return _cTFontGetSize(font), nil
}

// CTFontGetSize returns the point size of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetSize(_:)
func CTFontGetSize(font CTFontRef) float64 {
	result, callErr := tryCTFontGetSize(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetSlantAngle func(font CTFontRef) float64
var _cTFontGetSlantAngleErr error

func tryCTFontGetSlantAngle(font CTFontRef) (float64, error) {
	if _cTFontGetSlantAngle == nil {
		return 0.0, symbolCallError("CTFontGetSlantAngle", "10.5", _cTFontGetSlantAngleErr)
	}
	return _cTFontGetSlantAngle(font), nil
}

// CTFontGetSlantAngle returns the slant angle of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetSlantAngle(_:)
func CTFontGetSlantAngle(font CTFontRef) float64 {
	result, callErr := tryCTFontGetSlantAngle(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetStringEncoding func(font CTFontRef) uint32
var _cTFontGetStringEncodingErr error

func tryCTFontGetStringEncoding(font CTFontRef) (uint32, error) {
	if _cTFontGetStringEncoding == nil {
		return 0, symbolCallError("CTFontGetStringEncoding", "10.5", _cTFontGetStringEncodingErr)
	}
	return _cTFontGetStringEncoding(font), nil
}

// CTFontGetStringEncoding returns the best string encoding for legacy format support.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetStringEncoding(_:)
func CTFontGetStringEncoding(font CTFontRef) uint32 {
	result, callErr := tryCTFontGetStringEncoding(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetSymbolicTraits func(font CTFontRef) CTFontSymbolicTraits
var _cTFontGetSymbolicTraitsErr error

func tryCTFontGetSymbolicTraits(font CTFontRef) (CTFontSymbolicTraits, error) {
	if _cTFontGetSymbolicTraits == nil {
		return *new(CTFontSymbolicTraits), symbolCallError("CTFontGetSymbolicTraits", "10.5", _cTFontGetSymbolicTraitsErr)
	}
	return _cTFontGetSymbolicTraits(font), nil
}

// CTFontGetSymbolicTraits returns the symbolic traits of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetSymbolicTraits(_:)
func CTFontGetSymbolicTraits(font CTFontRef) CTFontSymbolicTraits {
	result, callErr := tryCTFontGetSymbolicTraits(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetTypeID func() uint
var _cTFontGetTypeIDErr error

func tryCTFontGetTypeID() (uint, error) {
	if _cTFontGetTypeID == nil {
		return 0, symbolCallError("CTFontGetTypeID", "10.5", _cTFontGetTypeIDErr)
	}
	return _cTFontGetTypeID(), nil
}

// CTFontGetTypeID returns the type identifier for Core Text font references.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetTypeID()
func CTFontGetTypeID() uint {
	result, callErr := tryCTFontGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetTypographicBoundsForAdaptiveImageProvider func(font CTFontRef, provider CTAdaptiveImageProvidingObject) corefoundation.CGRect
var _cTFontGetTypographicBoundsForAdaptiveImageProviderErr error

func tryCTFontGetTypographicBoundsForAdaptiveImageProvider(font CTFontRef, provider CTAdaptiveImageProvidingObject) (corefoundation.CGRect, error) {
	if _cTFontGetTypographicBoundsForAdaptiveImageProvider == nil {
		return corefoundation.CGRect{}, symbolCallError("CTFontGetTypographicBoundsForAdaptiveImageProvider", "15.0", _cTFontGetTypographicBoundsForAdaptiveImageProviderErr)
	}
	return _cTFontGetTypographicBoundsForAdaptiveImageProvider(font, provider), nil
}

// CTFontGetTypographicBoundsForAdaptiveImageProvider.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetTypographicBoundsForAdaptiveImageProvider(_:_:)
func CTFontGetTypographicBoundsForAdaptiveImageProvider(font CTFontRef, provider CTAdaptiveImageProvidingObject) corefoundation.CGRect {
	result, callErr := tryCTFontGetTypographicBoundsForAdaptiveImageProvider(font, provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetUIFontType func(font CTFontRef) CTFontUIFontType
var _cTFontGetUIFontTypeErr error

func tryCTFontGetUIFontType(font CTFontRef) (CTFontUIFontType, error) {
	if _cTFontGetUIFontType == nil {
		return *new(CTFontUIFontType), symbolCallError("CTFontGetUIFontType", "10.15", _cTFontGetUIFontTypeErr)
	}
	return _cTFontGetUIFontType(font), nil
}

// CTFontGetUIFontType.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetUIFontType(_:)
func CTFontGetUIFontType(font CTFontRef) CTFontUIFontType {
	result, callErr := tryCTFontGetUIFontType(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetUnderlinePosition func(font CTFontRef) float64
var _cTFontGetUnderlinePositionErr error

func tryCTFontGetUnderlinePosition(font CTFontRef) (float64, error) {
	if _cTFontGetUnderlinePosition == nil {
		return 0.0, symbolCallError("CTFontGetUnderlinePosition", "10.5", _cTFontGetUnderlinePositionErr)
	}
	return _cTFontGetUnderlinePosition(font), nil
}

// CTFontGetUnderlinePosition returns the scaled underline position of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetUnderlinePosition(_:)
func CTFontGetUnderlinePosition(font CTFontRef) float64 {
	result, callErr := tryCTFontGetUnderlinePosition(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetUnderlineThickness func(font CTFontRef) float64
var _cTFontGetUnderlineThicknessErr error

func tryCTFontGetUnderlineThickness(font CTFontRef) (float64, error) {
	if _cTFontGetUnderlineThickness == nil {
		return 0.0, symbolCallError("CTFontGetUnderlineThickness", "10.5", _cTFontGetUnderlineThicknessErr)
	}
	return _cTFontGetUnderlineThickness(font), nil
}

// CTFontGetUnderlineThickness returns the scaled underline-thickness metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetUnderlineThickness(_:)
func CTFontGetUnderlineThickness(font CTFontRef) float64 {
	result, callErr := tryCTFontGetUnderlineThickness(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetUnitsPerEm func(font CTFontRef) uint32
var _cTFontGetUnitsPerEmErr error

func tryCTFontGetUnitsPerEm(font CTFontRef) (uint32, error) {
	if _cTFontGetUnitsPerEm == nil {
		return 0, symbolCallError("CTFontGetUnitsPerEm", "10.5", _cTFontGetUnitsPerEmErr)
	}
	return _cTFontGetUnitsPerEm(font), nil
}

// CTFontGetUnitsPerEm returns the units-per-em metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetUnitsPerEm(_:)
func CTFontGetUnitsPerEm(font CTFontRef) uint32 {
	result, callErr := tryCTFontGetUnitsPerEm(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontGetVerticalTranslationsForGlyphs func(font CTFontRef, glyphs unsafe.Pointer, translations *corefoundation.CGSize, count int)
var _cTFontGetVerticalTranslationsForGlyphsErr error

func tryCTFontGetVerticalTranslationsForGlyphs(font CTFontRef, glyphs unsafe.Pointer, translations *corefoundation.CGSize, count int) error {
	if _cTFontGetVerticalTranslationsForGlyphs == nil {
		return symbolCallError("CTFontGetVerticalTranslationsForGlyphs", "10.5", _cTFontGetVerticalTranslationsForGlyphsErr)
	}
	_cTFontGetVerticalTranslationsForGlyphs(font, glyphs, translations, count)
	return nil
}

// CTFontGetVerticalTranslationsForGlyphs calculates the offset from the default (horizontal) origin to the vertical origin for an array of glyphs.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetVerticalTranslationsForGlyphs(_:_:_:_:)
func CTFontGetVerticalTranslationsForGlyphs(font CTFontRef, glyphs unsafe.Pointer, translations *corefoundation.CGSize, count int) {
	if callErr := tryCTFontGetVerticalTranslationsForGlyphs(font, glyphs, translations, count); callErr != nil {
		panic(callErr)
	}
}

var _cTFontGetXHeight func(font CTFontRef) float64
var _cTFontGetXHeightErr error

func tryCTFontGetXHeight(font CTFontRef) (float64, error) {
	if _cTFontGetXHeight == nil {
		return 0.0, symbolCallError("CTFontGetXHeight", "10.5", _cTFontGetXHeightErr)
	}
	return _cTFontGetXHeight(font), nil
}

// CTFontGetXHeight returns the x-height metric of the given font.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontGetXHeight(_:)
func CTFontGetXHeight(font CTFontRef) float64 {
	result, callErr := tryCTFontGetXHeight(font)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontHasTable func(font CTFontRef, tag CTFontTableTag) bool
var _cTFontHasTableErr error

func tryCTFontHasTable(font CTFontRef, tag CTFontTableTag) (bool, error) {
	if _cTFontHasTable == nil {
		return false, symbolCallError("CTFontHasTable", "10.15", _cTFontHasTableErr)
	}
	return _cTFontHasTable(font, tag), nil
}

// CTFontHasTable.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontHasTable(_:_:)
func CTFontHasTable(font CTFontRef, tag CTFontTableTag) bool {
	result, callErr := tryCTFontHasTable(font, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCompareFontFamilyNames func(family1 unsafe.Pointer, family2 unsafe.Pointer, context unsafe.Pointer) corefoundation.CFComparisonResult
var _cTFontManagerCompareFontFamilyNamesErr error

func tryCTFontManagerCompareFontFamilyNames(family1 unsafe.Pointer, family2 unsafe.Pointer, context unsafe.Pointer) (corefoundation.CFComparisonResult, error) {
	if _cTFontManagerCompareFontFamilyNames == nil {
		return *new(corefoundation.CFComparisonResult), symbolCallError("CTFontManagerCompareFontFamilyNames", "10.6", _cTFontManagerCompareFontFamilyNamesErr)
	}
	return _cTFontManagerCompareFontFamilyNames(family1, family2, context), nil
}

// CTFontManagerCompareFontFamilyNames a comparator function to compare font family names and sort them according to Apple guidelines.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCompareFontFamilyNames(_:_:_:)
func CTFontManagerCompareFontFamilyNames(family1 unsafe.Pointer, family2 unsafe.Pointer, context unsafe.Pointer) corefoundation.CFComparisonResult {
	result, callErr := tryCTFontManagerCompareFontFamilyNames(family1, family2, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCopyAvailableFontFamilyNames func() corefoundation.CFArrayRef
var _cTFontManagerCopyAvailableFontFamilyNamesErr error

func tryCTFontManagerCopyAvailableFontFamilyNames() (corefoundation.CFArrayRef, error) {
	if _cTFontManagerCopyAvailableFontFamilyNames == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontManagerCopyAvailableFontFamilyNames", "10.6", _cTFontManagerCopyAvailableFontFamilyNamesErr)
	}
	return _cTFontManagerCopyAvailableFontFamilyNames(), nil
}

// CTFontManagerCopyAvailableFontFamilyNames returns an array of visible font family names sorted for user interface display.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailableFontFamilyNames()
func CTFontManagerCopyAvailableFontFamilyNames() corefoundation.CFArrayRef {
	result, callErr := tryCTFontManagerCopyAvailableFontFamilyNames()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCopyAvailableFontURLs func() corefoundation.CFArrayRef
var _cTFontManagerCopyAvailableFontURLsErr error

func tryCTFontManagerCopyAvailableFontURLs() (corefoundation.CFArrayRef, error) {
	if _cTFontManagerCopyAvailableFontURLs == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontManagerCopyAvailableFontURLs", "10.6", _cTFontManagerCopyAvailableFontURLsErr)
	}
	return _cTFontManagerCopyAvailableFontURLs(), nil
}

// CTFontManagerCopyAvailableFontURLs returns an array of font URLs.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailableFontURLs()
func CTFontManagerCopyAvailableFontURLs() corefoundation.CFArrayRef {
	result, callErr := tryCTFontManagerCopyAvailableFontURLs()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCopyAvailablePostScriptNames func() corefoundation.CFArrayRef
var _cTFontManagerCopyAvailablePostScriptNamesErr error

func tryCTFontManagerCopyAvailablePostScriptNames() (corefoundation.CFArrayRef, error) {
	if _cTFontManagerCopyAvailablePostScriptNames == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontManagerCopyAvailablePostScriptNames", "10.6", _cTFontManagerCopyAvailablePostScriptNamesErr)
	}
	return _cTFontManagerCopyAvailablePostScriptNames(), nil
}

// CTFontManagerCopyAvailablePostScriptNames returns an array of unique PostScript font names for the fonts.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCopyAvailablePostScriptNames()
func CTFontManagerCopyAvailablePostScriptNames() corefoundation.CFArrayRef {
	result, callErr := tryCTFontManagerCopyAvailablePostScriptNames()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCreateFontDescriptorFromData func(data corefoundation.CFDataRef) CTFontDescriptorRef
var _cTFontManagerCreateFontDescriptorFromDataErr error

func tryCTFontManagerCreateFontDescriptorFromData(data corefoundation.CFDataRef) (CTFontDescriptorRef, error) {
	if _cTFontManagerCreateFontDescriptorFromData == nil {
		return *new(CTFontDescriptorRef), symbolCallError("CTFontManagerCreateFontDescriptorFromData", "10.7", _cTFontManagerCreateFontDescriptorFromDataErr)
	}
	return _cTFontManagerCreateFontDescriptorFromData(data), nil
}

// CTFontManagerCreateFontDescriptorFromData creates a font descriptor representing the font in the supplied data.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorFromData(_:)
func CTFontManagerCreateFontDescriptorFromData(data corefoundation.CFDataRef) CTFontDescriptorRef {
	result, callErr := tryCTFontManagerCreateFontDescriptorFromData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCreateFontDescriptorsFromData func(data corefoundation.CFDataRef) corefoundation.CFArrayRef
var _cTFontManagerCreateFontDescriptorsFromDataErr error

func tryCTFontManagerCreateFontDescriptorsFromData(data corefoundation.CFDataRef) (corefoundation.CFArrayRef, error) {
	if _cTFontManagerCreateFontDescriptorsFromData == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontManagerCreateFontDescriptorsFromData", "10.13", _cTFontManagerCreateFontDescriptorsFromDataErr)
	}
	return _cTFontManagerCreateFontDescriptorsFromData(data), nil
}

// CTFontManagerCreateFontDescriptorsFromData creates an array of font descriptors for the fonts in the supplied data.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorsFromData(_:)
func CTFontManagerCreateFontDescriptorsFromData(data corefoundation.CFDataRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontManagerCreateFontDescriptorsFromData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerCreateFontDescriptorsFromURL func(fileURL corefoundation.CFURLRef) corefoundation.CFArrayRef
var _cTFontManagerCreateFontDescriptorsFromURLErr error

func tryCTFontManagerCreateFontDescriptorsFromURL(fileURL corefoundation.CFURLRef) (corefoundation.CFArrayRef, error) {
	if _cTFontManagerCreateFontDescriptorsFromURL == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFontManagerCreateFontDescriptorsFromURL", "10.6", _cTFontManagerCreateFontDescriptorsFromURLErr)
	}
	return _cTFontManagerCreateFontDescriptorsFromURL(fileURL), nil
}

// CTFontManagerCreateFontDescriptorsFromURL returns an array of font descriptors representing each of the fonts in the specified URL.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerCreateFontDescriptorsFromURL(_:)
func CTFontManagerCreateFontDescriptorsFromURL(fileURL corefoundation.CFURLRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFontManagerCreateFontDescriptorsFromURL(fileURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerEnableFontDescriptors func(descriptors corefoundation.CFArrayRef, enable bool)
var _cTFontManagerEnableFontDescriptorsErr error

func tryCTFontManagerEnableFontDescriptors(descriptors corefoundation.CFArrayRef, enable bool) error {
	if _cTFontManagerEnableFontDescriptors == nil {
		return symbolCallError("CTFontManagerEnableFontDescriptors", "10.6", _cTFontManagerEnableFontDescriptorsErr)
	}
	_cTFontManagerEnableFontDescriptors(descriptors, enable)
	return nil
}

// CTFontManagerEnableFontDescriptors enables or disables the matching font descriptors for font descriptor matching.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerEnableFontDescriptors(_:_:)
func CTFontManagerEnableFontDescriptors(descriptors corefoundation.CFArrayRef, enable bool) {
	if callErr := tryCTFontManagerEnableFontDescriptors(descriptors, enable); callErr != nil {
		panic(callErr)
	}
}

var _cTFontManagerGetAutoActivationSetting func(bundleIdentifier corefoundation.CFStringRef) CTFontManagerAutoActivationSetting
var _cTFontManagerGetAutoActivationSettingErr error

func tryCTFontManagerGetAutoActivationSetting(bundleIdentifier corefoundation.CFStringRef) (CTFontManagerAutoActivationSetting, error) {
	if _cTFontManagerGetAutoActivationSetting == nil {
		return *new(CTFontManagerAutoActivationSetting), symbolCallError("CTFontManagerGetAutoActivationSetting", "10.6", _cTFontManagerGetAutoActivationSettingErr)
	}
	return _cTFontManagerGetAutoActivationSetting(bundleIdentifier), nil
}

// CTFontManagerGetAutoActivationSetting gets the auto-activation setting for the specified bundle identifier.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerGetAutoActivationSetting(_:)
func CTFontManagerGetAutoActivationSetting(bundleIdentifier corefoundation.CFStringRef) CTFontManagerAutoActivationSetting {
	result, callErr := tryCTFontManagerGetAutoActivationSetting(bundleIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerGetScopeForURL func(fontURL corefoundation.CFURLRef) CTFontManagerScope
var _cTFontManagerGetScopeForURLErr error

func tryCTFontManagerGetScopeForURL(fontURL corefoundation.CFURLRef) (CTFontManagerScope, error) {
	if _cTFontManagerGetScopeForURL == nil {
		return *new(CTFontManagerScope), symbolCallError("CTFontManagerGetScopeForURL", "10.6", _cTFontManagerGetScopeForURLErr)
	}
	return _cTFontManagerGetScopeForURL(fontURL), nil
}

// CTFontManagerGetScopeForURL returns the registration scope of the specified URL.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerGetScopeForURL(_:)
func CTFontManagerGetScopeForURL(fontURL corefoundation.CFURLRef) CTFontManagerScope {
	result, callErr := tryCTFontManagerGetScopeForURL(fontURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerIsSupportedFont func(fontURL corefoundation.CFURLRef) bool
var _cTFontManagerIsSupportedFontErr error

func tryCTFontManagerIsSupportedFont(fontURL corefoundation.CFURLRef) (bool, error) {
	if _cTFontManagerIsSupportedFont == nil {
		return false, symbolCallError("CTFontManagerIsSupportedFont", "10.6", _cTFontManagerIsSupportedFontErr)
	}
	return _cTFontManagerIsSupportedFont(fontURL), nil
}

// CTFontManagerIsSupportedFont determines whether a file is in a supported font format.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerIsSupportedFont(_:)
func CTFontManagerIsSupportedFont(fontURL corefoundation.CFURLRef) bool {
	result, callErr := tryCTFontManagerIsSupportedFont(fontURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerRegisterFontDescriptors func(fontDescriptors corefoundation.CFArrayRef, scope CTFontManagerScope, enabled bool, registrationHandler bool)
var _cTFontManagerRegisterFontDescriptorsErr error

func tryCTFontManagerRegisterFontDescriptors(fontDescriptors corefoundation.CFArrayRef, scope CTFontManagerScope, enabled bool, registrationHandler bool) error {
	if _cTFontManagerRegisterFontDescriptors == nil {
		return symbolCallError("CTFontManagerRegisterFontDescriptors", "10.15", _cTFontManagerRegisterFontDescriptorsErr)
	}
	_cTFontManagerRegisterFontDescriptors(fontDescriptors, scope, enabled, registrationHandler)
	return nil
}

// CTFontManagerRegisterFontDescriptors registers font descriptors with the font manager.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontDescriptors(_:_:_:_:)
func CTFontManagerRegisterFontDescriptors(fontDescriptors corefoundation.CFArrayRef, scope CTFontManagerScope, enabled bool, registrationHandler bool) {
	if callErr := tryCTFontManagerRegisterFontDescriptors(fontDescriptors, scope, enabled, registrationHandler); callErr != nil {
		panic(callErr)
	}
}

var _cTFontManagerRegisterFontURLs func(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, enabled bool, registrationHandler bool)
var _cTFontManagerRegisterFontURLsErr error

func tryCTFontManagerRegisterFontURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, enabled bool, registrationHandler bool) error {
	if _cTFontManagerRegisterFontURLs == nil {
		return symbolCallError("CTFontManagerRegisterFontURLs", "10.15", _cTFontManagerRegisterFontURLsErr)
	}
	_cTFontManagerRegisterFontURLs(fontURLs, scope, enabled, registrationHandler)
	return nil
}

// CTFontManagerRegisterFontURLs registers fonts from the specified font URLs with the font manager.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontURLs(_:_:_:_:)
func CTFontManagerRegisterFontURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, enabled bool, registrationHandler bool) {
	if callErr := tryCTFontManagerRegisterFontURLs(fontURLs, scope, enabled, registrationHandler); callErr != nil {
		panic(callErr)
	}
}

var _cTFontManagerRegisterFontsForURL func(fontURL corefoundation.CFURLRef, scope CTFontManagerScope, err *corefoundation.CFErrorRef) bool
var _cTFontManagerRegisterFontsForURLErr error

func tryCTFontManagerRegisterFontsForURL(fontURL corefoundation.CFURLRef, scope CTFontManagerScope, err *corefoundation.CFErrorRef) (bool, error) {
	if _cTFontManagerRegisterFontsForURL == nil {
		return false, symbolCallError("CTFontManagerRegisterFontsForURL", "10.6", _cTFontManagerRegisterFontsForURLErr)
	}
	return _cTFontManagerRegisterFontsForURL(fontURL, scope, err), nil
}

// CTFontManagerRegisterFontsForURL registers fonts from the specified font URL with the Font Manager.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsForURL(_:_:_:)
func CTFontManagerRegisterFontsForURL(fontURL corefoundation.CFURLRef, scope CTFontManagerScope, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryCTFontManagerRegisterFontsForURL(fontURL, scope, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerRegisterFontsForURLs func(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, errors *corefoundation.CFArrayRef) bool
var _cTFontManagerRegisterFontsForURLsErr error

func tryCTFontManagerRegisterFontsForURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, errors *corefoundation.CFArrayRef) (bool, error) {
	if _cTFontManagerRegisterFontsForURLs == nil {
		return false, symbolCallError("CTFontManagerRegisterFontsForURLs", "10.6", _cTFontManagerRegisterFontsForURLsErr)
	}
	return _cTFontManagerRegisterFontsForURLs(fontURLs, scope, errors), nil
}

// CTFontManagerRegisterFontsForURLs registers fonts from the specified array of font URLs with the Font Manager.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterFontsForURLs(_:_:_:)
func CTFontManagerRegisterFontsForURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, errors *corefoundation.CFArrayRef) bool {
	result, callErr := tryCTFontManagerRegisterFontsForURLs(fontURLs, scope, errors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerRegisterGraphicsFont func(font coregraphics.CGFont, err *corefoundation.CFErrorRef) bool
var _cTFontManagerRegisterGraphicsFontErr error

func tryCTFontManagerRegisterGraphicsFont(font coregraphics.CGFont, err *corefoundation.CFErrorRef) (bool, error) {
	if _cTFontManagerRegisterGraphicsFont == nil {
		return false, symbolCallError("CTFontManagerRegisterGraphicsFont", "10.8", _cTFontManagerRegisterGraphicsFontErr)
	}
	return _cTFontManagerRegisterGraphicsFont(font, err), nil
}

// CTFontManagerRegisterGraphicsFont registers the specified graphics font with the font manager.
//
// Deprecated: Deprecated since macOS 15.0. Use CTFontManagerCreateFontDescriptorsFromData or CTFontManagerRegisterFontsForURL
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerRegisterGraphicsFont(_:_:)
func CTFontManagerRegisterGraphicsFont(font coregraphics.CGFont, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryCTFontManagerRegisterGraphicsFont(font, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerSetAutoActivationSetting func(bundleIdentifier corefoundation.CFStringRef, setting CTFontManagerAutoActivationSetting)
var _cTFontManagerSetAutoActivationSettingErr error

func tryCTFontManagerSetAutoActivationSetting(bundleIdentifier corefoundation.CFStringRef, setting CTFontManagerAutoActivationSetting) error {
	if _cTFontManagerSetAutoActivationSetting == nil {
		return symbolCallError("CTFontManagerSetAutoActivationSetting", "10.6", _cTFontManagerSetAutoActivationSettingErr)
	}
	_cTFontManagerSetAutoActivationSetting(bundleIdentifier, setting)
	return nil
}

// CTFontManagerSetAutoActivationSetting sets the auto-activation setting for the specified bundle identifier.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerSetAutoActivationSetting(_:_:)
func CTFontManagerSetAutoActivationSetting(bundleIdentifier corefoundation.CFStringRef, setting CTFontManagerAutoActivationSetting) {
	if callErr := tryCTFontManagerSetAutoActivationSetting(bundleIdentifier, setting); callErr != nil {
		panic(callErr)
	}
}

var _cTFontManagerUnregisterFontDescriptors func(fontDescriptors corefoundation.CFArrayRef, scope CTFontManagerScope, registrationHandler bool)
var _cTFontManagerUnregisterFontDescriptorsErr error

func tryCTFontManagerUnregisterFontDescriptors(fontDescriptors corefoundation.CFArrayRef, scope CTFontManagerScope, registrationHandler bool) error {
	if _cTFontManagerUnregisterFontDescriptors == nil {
		return symbolCallError("CTFontManagerUnregisterFontDescriptors", "10.15", _cTFontManagerUnregisterFontDescriptorsErr)
	}
	_cTFontManagerUnregisterFontDescriptors(fontDescriptors, scope, registrationHandler)
	return nil
}

// CTFontManagerUnregisterFontDescriptors unregisters font descriptors with the font manager.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontDescriptors(_:_:_:)
func CTFontManagerUnregisterFontDescriptors(fontDescriptors corefoundation.CFArrayRef, scope CTFontManagerScope, registrationHandler bool) {
	if callErr := tryCTFontManagerUnregisterFontDescriptors(fontDescriptors, scope, registrationHandler); callErr != nil {
		panic(callErr)
	}
}

var _cTFontManagerUnregisterFontURLs func(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, registrationHandler bool)
var _cTFontManagerUnregisterFontURLsErr error

func tryCTFontManagerUnregisterFontURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, registrationHandler bool) error {
	if _cTFontManagerUnregisterFontURLs == nil {
		return symbolCallError("CTFontManagerUnregisterFontURLs", "10.15", _cTFontManagerUnregisterFontURLsErr)
	}
	_cTFontManagerUnregisterFontURLs(fontURLs, scope, registrationHandler)
	return nil
}

// CTFontManagerUnregisterFontURLs unregisters fonts from the specified font URLs with the font manager.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontURLs(_:_:_:)
func CTFontManagerUnregisterFontURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, registrationHandler bool) {
	if callErr := tryCTFontManagerUnregisterFontURLs(fontURLs, scope, registrationHandler); callErr != nil {
		panic(callErr)
	}
}

var _cTFontManagerUnregisterFontsForURL func(fontURL corefoundation.CFURLRef, scope CTFontManagerScope, err *corefoundation.CFErrorRef) bool
var _cTFontManagerUnregisterFontsForURLErr error

func tryCTFontManagerUnregisterFontsForURL(fontURL corefoundation.CFURLRef, scope CTFontManagerScope, err *corefoundation.CFErrorRef) (bool, error) {
	if _cTFontManagerUnregisterFontsForURL == nil {
		return false, symbolCallError("CTFontManagerUnregisterFontsForURL", "10.6", _cTFontManagerUnregisterFontsForURLErr)
	}
	return _cTFontManagerUnregisterFontsForURL(fontURL, scope, err), nil
}

// CTFontManagerUnregisterFontsForURL unregisters fonts from the specified font URL with the Font Manager.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontsForURL(_:_:_:)
func CTFontManagerUnregisterFontsForURL(fontURL corefoundation.CFURLRef, scope CTFontManagerScope, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryCTFontManagerUnregisterFontsForURL(fontURL, scope, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerUnregisterFontsForURLs func(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, errors *corefoundation.CFArrayRef) bool
var _cTFontManagerUnregisterFontsForURLsErr error

func tryCTFontManagerUnregisterFontsForURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, errors *corefoundation.CFArrayRef) (bool, error) {
	if _cTFontManagerUnregisterFontsForURLs == nil {
		return false, symbolCallError("CTFontManagerUnregisterFontsForURLs", "10.6", _cTFontManagerUnregisterFontsForURLsErr)
	}
	return _cTFontManagerUnregisterFontsForURLs(fontURLs, scope, errors), nil
}

// CTFontManagerUnregisterFontsForURLs unregisters fonts from the specified array of font URLs with the Font Manager.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterFontsForURLs(_:_:_:)
func CTFontManagerUnregisterFontsForURLs(fontURLs corefoundation.CFArrayRef, scope CTFontManagerScope, errors *corefoundation.CFArrayRef) bool {
	result, callErr := tryCTFontManagerUnregisterFontsForURLs(fontURLs, scope, errors)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFontManagerUnregisterGraphicsFont func(font coregraphics.CGFont, err *corefoundation.CFErrorRef) bool
var _cTFontManagerUnregisterGraphicsFontErr error

func tryCTFontManagerUnregisterGraphicsFont(font coregraphics.CGFont, err *corefoundation.CFErrorRef) (bool, error) {
	if _cTFontManagerUnregisterGraphicsFont == nil {
		return false, symbolCallError("CTFontManagerUnregisterGraphicsFont", "10.8", _cTFontManagerUnregisterGraphicsFontErr)
	}
	return _cTFontManagerUnregisterGraphicsFont(font, err), nil
}

// CTFontManagerUnregisterGraphicsFont unregisters the specified graphics font with the font manager.
//
// Deprecated: Deprecated since macOS 15.0. Use the API corresponding to the one used to register the font
//
// See: https://developer.apple.com/documentation/CoreText/CTFontManagerUnregisterGraphicsFont(_:_:)
func CTFontManagerUnregisterGraphicsFont(font coregraphics.CGFont, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryCTFontManagerUnregisterGraphicsFont(font, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFrameDraw func(frame CTFrameRef, context coregraphics.CGContextRef)
var _cTFrameDrawErr error

func tryCTFrameDraw(frame CTFrameRef, context coregraphics.CGContextRef) error {
	if _cTFrameDraw == nil {
		return symbolCallError("CTFrameDraw", "10.5", _cTFrameDrawErr)
	}
	_cTFrameDraw(frame, context)
	return nil
}

// CTFrameDraw draws an entire frame into a context.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameDraw(_:_:)
func CTFrameDraw(frame CTFrameRef, context coregraphics.CGContextRef) {
	if callErr := tryCTFrameDraw(frame, context); callErr != nil {
		panic(callErr)
	}
}

var _cTFrameGetFrameAttributes func(frame CTFrameRef) corefoundation.CFDictionaryRef
var _cTFrameGetFrameAttributesErr error

func tryCTFrameGetFrameAttributes(frame CTFrameRef) (corefoundation.CFDictionaryRef, error) {
	if _cTFrameGetFrameAttributes == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CTFrameGetFrameAttributes", "10.5", _cTFrameGetFrameAttributesErr)
	}
	return _cTFrameGetFrameAttributes(frame), nil
}

// CTFrameGetFrameAttributes returns the frame attributes used to create the frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetFrameAttributes(_:)
func CTFrameGetFrameAttributes(frame CTFrameRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCTFrameGetFrameAttributes(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFrameGetLineOrigins func(frame CTFrameRef, range_ corefoundation.CFRange, origins *corefoundation.CGPoint)
var _cTFrameGetLineOriginsErr error

func tryCTFrameGetLineOrigins(frame CTFrameRef, range_ corefoundation.CFRange, origins *corefoundation.CGPoint) error {
	if _cTFrameGetLineOrigins == nil {
		return symbolCallError("CTFrameGetLineOrigins", "10.5", _cTFrameGetLineOriginsErr)
	}
	_cTFrameGetLineOrigins(frame, range_, origins)
	return nil
}

// CTFrameGetLineOrigins copies a range of line origins for a frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetLineOrigins(_:_:_:)
func CTFrameGetLineOrigins(frame CTFrameRef, range_ corefoundation.CFRange, origins *corefoundation.CGPoint) {
	if callErr := tryCTFrameGetLineOrigins(frame, range_, origins); callErr != nil {
		panic(callErr)
	}
}

var _cTFrameGetLines func(frame CTFrameRef) corefoundation.CFArrayRef
var _cTFrameGetLinesErr error

func tryCTFrameGetLines(frame CTFrameRef) (corefoundation.CFArrayRef, error) {
	if _cTFrameGetLines == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTFrameGetLines", "10.5", _cTFrameGetLinesErr)
	}
	return _cTFrameGetLines(frame), nil
}

// CTFrameGetLines returns an array of lines stored in the frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetLines(_:)
func CTFrameGetLines(frame CTFrameRef) corefoundation.CFArrayRef {
	result, callErr := tryCTFrameGetLines(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFrameGetPath func(frame CTFrameRef) coregraphics.CGPathRef
var _cTFrameGetPathErr error

func tryCTFrameGetPath(frame CTFrameRef) (coregraphics.CGPathRef, error) {
	if _cTFrameGetPath == nil {
		return *new(coregraphics.CGPathRef), symbolCallError("CTFrameGetPath", "10.5", _cTFrameGetPathErr)
	}
	return _cTFrameGetPath(frame), nil
}

// CTFrameGetPath returns the path used to create the frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetPath(_:)
func CTFrameGetPath(frame CTFrameRef) coregraphics.CGPathRef {
	result, callErr := tryCTFrameGetPath(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFrameGetStringRange func(frame CTFrameRef) corefoundation.CFRange
var _cTFrameGetStringRangeErr error

func tryCTFrameGetStringRange(frame CTFrameRef) (corefoundation.CFRange, error) {
	if _cTFrameGetStringRange == nil {
		return corefoundation.CFRange{}, symbolCallError("CTFrameGetStringRange", "10.5", _cTFrameGetStringRangeErr)
	}
	return _cTFrameGetStringRange(frame), nil
}

// CTFrameGetStringRange returns the range of characters originally requested to fill the frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetStringRange(_:)
func CTFrameGetStringRange(frame CTFrameRef) corefoundation.CFRange {
	result, callErr := tryCTFrameGetStringRange(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFrameGetTypeID func() uint
var _cTFrameGetTypeIDErr error

func tryCTFrameGetTypeID() (uint, error) {
	if _cTFrameGetTypeID == nil {
		return 0, symbolCallError("CTFrameGetTypeID", "10.5", _cTFrameGetTypeIDErr)
	}
	return _cTFrameGetTypeID(), nil
}

// CTFrameGetTypeID returns the type identifier for the CTFrame opaque type.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetTypeID()
func CTFrameGetTypeID() uint {
	result, callErr := tryCTFrameGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFrameGetVisibleStringRange func(frame CTFrameRef) corefoundation.CFRange
var _cTFrameGetVisibleStringRangeErr error

func tryCTFrameGetVisibleStringRange(frame CTFrameRef) (corefoundation.CFRange, error) {
	if _cTFrameGetVisibleStringRange == nil {
		return corefoundation.CFRange{}, symbolCallError("CTFrameGetVisibleStringRange", "10.5", _cTFrameGetVisibleStringRangeErr)
	}
	return _cTFrameGetVisibleStringRange(frame), nil
}

// CTFrameGetVisibleStringRange returns the range of characters that actually fit in the frame.
//
// See: https://developer.apple.com/documentation/CoreText/CTFrameGetVisibleStringRange(_:)
func CTFrameGetVisibleStringRange(frame CTFrameRef) corefoundation.CFRange {
	result, callErr := tryCTFrameGetVisibleStringRange(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFramesetterCreateFrame func(framesetter CTFramesetterRef, stringRange corefoundation.CFRange, path coregraphics.CGPathRef, frameAttributes corefoundation.CFDictionaryRef) CTFrameRef
var _cTFramesetterCreateFrameErr error

func tryCTFramesetterCreateFrame(framesetter CTFramesetterRef, stringRange corefoundation.CFRange, path coregraphics.CGPathRef, frameAttributes corefoundation.CFDictionaryRef) (CTFrameRef, error) {
	if _cTFramesetterCreateFrame == nil {
		return *new(CTFrameRef), symbolCallError("CTFramesetterCreateFrame", "10.5", _cTFramesetterCreateFrameErr)
	}
	return _cTFramesetterCreateFrame(framesetter, stringRange, path, frameAttributes), nil
}

// CTFramesetterCreateFrame creates an immutable frame using a framesetter.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateFrame(_:_:_:_:)
func CTFramesetterCreateFrame(framesetter CTFramesetterRef, stringRange corefoundation.CFRange, path coregraphics.CGPathRef, frameAttributes corefoundation.CFDictionaryRef) CTFrameRef {
	result, callErr := tryCTFramesetterCreateFrame(framesetter, stringRange, path, frameAttributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFramesetterCreateWithAttributedString func(attrString corefoundation.CFAttributedStringRef) CTFramesetterRef
var _cTFramesetterCreateWithAttributedStringErr error

func tryCTFramesetterCreateWithAttributedString(attrString corefoundation.CFAttributedStringRef) (CTFramesetterRef, error) {
	if _cTFramesetterCreateWithAttributedString == nil {
		return *new(CTFramesetterRef), symbolCallError("CTFramesetterCreateWithAttributedString", "10.5", _cTFramesetterCreateWithAttributedStringErr)
	}
	return _cTFramesetterCreateWithAttributedString(attrString), nil
}

// CTFramesetterCreateWithAttributedString creates an immutable framesetter object from an attributed string.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateWithAttributedString(_:)
func CTFramesetterCreateWithAttributedString(attrString corefoundation.CFAttributedStringRef) CTFramesetterRef {
	result, callErr := tryCTFramesetterCreateWithAttributedString(attrString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFramesetterCreateWithTypesetter func(typesetter CTTypesetterRef) CTFramesetterRef
var _cTFramesetterCreateWithTypesetterErr error

func tryCTFramesetterCreateWithTypesetter(typesetter CTTypesetterRef) (CTFramesetterRef, error) {
	if _cTFramesetterCreateWithTypesetter == nil {
		return *new(CTFramesetterRef), symbolCallError("CTFramesetterCreateWithTypesetter", "10.14", _cTFramesetterCreateWithTypesetterErr)
	}
	return _cTFramesetterCreateWithTypesetter(typesetter), nil
}

// CTFramesetterCreateWithTypesetter creates a framesetter directly from a typesetter.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetterCreateWithTypesetter(_:)
func CTFramesetterCreateWithTypesetter(typesetter CTTypesetterRef) CTFramesetterRef {
	result, callErr := tryCTFramesetterCreateWithTypesetter(typesetter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFramesetterGetTypeID func() uint
var _cTFramesetterGetTypeIDErr error

func tryCTFramesetterGetTypeID() (uint, error) {
	if _cTFramesetterGetTypeID == nil {
		return 0, symbolCallError("CTFramesetterGetTypeID", "10.5", _cTFramesetterGetTypeIDErr)
	}
	return _cTFramesetterGetTypeID(), nil
}

// CTFramesetterGetTypeID returns the Core Foundation type identifier of the framesetter object.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetterGetTypeID()
func CTFramesetterGetTypeID() uint {
	result, callErr := tryCTFramesetterGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFramesetterGetTypesetter func(framesetter CTFramesetterRef) CTTypesetterRef
var _cTFramesetterGetTypesetterErr error

func tryCTFramesetterGetTypesetter(framesetter CTFramesetterRef) (CTTypesetterRef, error) {
	if _cTFramesetterGetTypesetter == nil {
		return *new(CTTypesetterRef), symbolCallError("CTFramesetterGetTypesetter", "10.5", _cTFramesetterGetTypesetterErr)
	}
	return _cTFramesetterGetTypesetter(framesetter), nil
}

// CTFramesetterGetTypesetter returns the typesetter object being used by the framesetter.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetterGetTypesetter(_:)
func CTFramesetterGetTypesetter(framesetter CTFramesetterRef) CTTypesetterRef {
	result, callErr := tryCTFramesetterGetTypesetter(framesetter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTFramesetterSuggestFrameSizeWithConstraints func(framesetter CTFramesetterRef, stringRange corefoundation.CFRange, frameAttributes corefoundation.CFDictionaryRef, constraints corefoundation.CGSize, fitRange *corefoundation.CFRange) corefoundation.CGSize
var _cTFramesetterSuggestFrameSizeWithConstraintsErr error

func tryCTFramesetterSuggestFrameSizeWithConstraints(framesetter CTFramesetterRef, stringRange corefoundation.CFRange, frameAttributes corefoundation.CFDictionaryRef, constraints corefoundation.CGSize, fitRange *corefoundation.CFRange) (corefoundation.CGSize, error) {
	if _cTFramesetterSuggestFrameSizeWithConstraints == nil {
		return corefoundation.CGSize{}, symbolCallError("CTFramesetterSuggestFrameSizeWithConstraints", "10.5", _cTFramesetterSuggestFrameSizeWithConstraintsErr)
	}
	return _cTFramesetterSuggestFrameSizeWithConstraints(framesetter, stringRange, frameAttributes, constraints, fitRange), nil
}

// CTFramesetterSuggestFrameSizeWithConstraints determines the frame size needed for a string range.
//
// See: https://developer.apple.com/documentation/CoreText/CTFramesetterSuggestFrameSizeWithConstraints(_:_:_:_:_:)
func CTFramesetterSuggestFrameSizeWithConstraints(framesetter CTFramesetterRef, stringRange corefoundation.CFRange, frameAttributes corefoundation.CFDictionaryRef, constraints corefoundation.CGSize, fitRange *corefoundation.CFRange) corefoundation.CGSize {
	result, callErr := tryCTFramesetterSuggestFrameSizeWithConstraints(framesetter, stringRange, frameAttributes, constraints, fitRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoCreateWithCharacterIdentifier func(cid uint16, collection CTCharacterCollection, baseString corefoundation.CFStringRef) CTGlyphInfoRef
var _cTGlyphInfoCreateWithCharacterIdentifierErr error

func tryCTGlyphInfoCreateWithCharacterIdentifier(cid uint16, collection CTCharacterCollection, baseString corefoundation.CFStringRef) (CTGlyphInfoRef, error) {
	if _cTGlyphInfoCreateWithCharacterIdentifier == nil {
		return *new(CTGlyphInfoRef), symbolCallError("CTGlyphInfoCreateWithCharacterIdentifier", "10.5", _cTGlyphInfoCreateWithCharacterIdentifierErr)
	}
	return _cTGlyphInfoCreateWithCharacterIdentifier(cid, collection, baseString), nil
}

// CTGlyphInfoCreateWithCharacterIdentifier creates an immutable glyph info object with a character identifier.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithCharacterIdentifier(_:_:_:)
func CTGlyphInfoCreateWithCharacterIdentifier(cid uint16, collection CTCharacterCollection, baseString corefoundation.CFStringRef) CTGlyphInfoRef {
	result, callErr := tryCTGlyphInfoCreateWithCharacterIdentifier(cid, collection, baseString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoCreateWithGlyph func(glyph uint16, font CTFontRef, baseString corefoundation.CFStringRef) CTGlyphInfoRef
var _cTGlyphInfoCreateWithGlyphErr error

func tryCTGlyphInfoCreateWithGlyph(glyph uint16, font CTFontRef, baseString corefoundation.CFStringRef) (CTGlyphInfoRef, error) {
	if _cTGlyphInfoCreateWithGlyph == nil {
		return *new(CTGlyphInfoRef), symbolCallError("CTGlyphInfoCreateWithGlyph", "10.5", _cTGlyphInfoCreateWithGlyphErr)
	}
	return _cTGlyphInfoCreateWithGlyph(glyph, font, baseString), nil
}

// CTGlyphInfoCreateWithGlyph creates an immutable glyph info object with a glyph index.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithGlyph(_:_:_:)
func CTGlyphInfoCreateWithGlyph(glyph uint16, font CTFontRef, baseString corefoundation.CFStringRef) CTGlyphInfoRef {
	result, callErr := tryCTGlyphInfoCreateWithGlyph(glyph, font, baseString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoCreateWithGlyphName func(glyphName corefoundation.CFStringRef, font CTFontRef, baseString corefoundation.CFStringRef) CTGlyphInfoRef
var _cTGlyphInfoCreateWithGlyphNameErr error

func tryCTGlyphInfoCreateWithGlyphName(glyphName corefoundation.CFStringRef, font CTFontRef, baseString corefoundation.CFStringRef) (CTGlyphInfoRef, error) {
	if _cTGlyphInfoCreateWithGlyphName == nil {
		return *new(CTGlyphInfoRef), symbolCallError("CTGlyphInfoCreateWithGlyphName", "10.5", _cTGlyphInfoCreateWithGlyphNameErr)
	}
	return _cTGlyphInfoCreateWithGlyphName(glyphName, font, baseString), nil
}

// CTGlyphInfoCreateWithGlyphName creates an immutable glyph info object with a glyph name.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoCreateWithGlyphName(_:_:_:)
func CTGlyphInfoCreateWithGlyphName(glyphName corefoundation.CFStringRef, font CTFontRef, baseString corefoundation.CFStringRef) CTGlyphInfoRef {
	result, callErr := tryCTGlyphInfoCreateWithGlyphName(glyphName, font, baseString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoGetCharacterCollection func(glyphInfo CTGlyphInfoRef) CTCharacterCollection
var _cTGlyphInfoGetCharacterCollectionErr error

func tryCTGlyphInfoGetCharacterCollection(glyphInfo CTGlyphInfoRef) (CTCharacterCollection, error) {
	if _cTGlyphInfoGetCharacterCollection == nil {
		return *new(CTCharacterCollection), symbolCallError("CTGlyphInfoGetCharacterCollection", "10.5", _cTGlyphInfoGetCharacterCollectionErr)
	}
	return _cTGlyphInfoGetCharacterCollection(glyphInfo), nil
}

// CTGlyphInfoGetCharacterCollection gets the character collection for a glyph info object.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetCharacterCollection(_:)
func CTGlyphInfoGetCharacterCollection(glyphInfo CTGlyphInfoRef) CTCharacterCollection {
	result, callErr := tryCTGlyphInfoGetCharacterCollection(glyphInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoGetCharacterIdentifier func(glyphInfo CTGlyphInfoRef) uint16
var _cTGlyphInfoGetCharacterIdentifierErr error

func tryCTGlyphInfoGetCharacterIdentifier(glyphInfo CTGlyphInfoRef) (uint16, error) {
	if _cTGlyphInfoGetCharacterIdentifier == nil {
		return 0, symbolCallError("CTGlyphInfoGetCharacterIdentifier", "10.5", _cTGlyphInfoGetCharacterIdentifierErr)
	}
	return _cTGlyphInfoGetCharacterIdentifier(glyphInfo), nil
}

// CTGlyphInfoGetCharacterIdentifier gets the character identifier for a glyph info object.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetCharacterIdentifier(_:)
func CTGlyphInfoGetCharacterIdentifier(glyphInfo CTGlyphInfoRef) uint16 {
	result, callErr := tryCTGlyphInfoGetCharacterIdentifier(glyphInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoGetGlyph func(glyphInfo CTGlyphInfoRef) uint16
var _cTGlyphInfoGetGlyphErr error

func tryCTGlyphInfoGetGlyph(glyphInfo CTGlyphInfoRef) (uint16, error) {
	if _cTGlyphInfoGetGlyph == nil {
		return 0, symbolCallError("CTGlyphInfoGetGlyph", "10.15", _cTGlyphInfoGetGlyphErr)
	}
	return _cTGlyphInfoGetGlyph(glyphInfo), nil
}

// CTGlyphInfoGetGlyph retrieves the glyph for a glyph info, if that object exists.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetGlyph(_:)
func CTGlyphInfoGetGlyph(glyphInfo CTGlyphInfoRef) uint16 {
	result, callErr := tryCTGlyphInfoGetGlyph(glyphInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoGetGlyphName func(glyphInfo CTGlyphInfoRef) corefoundation.CFStringRef
var _cTGlyphInfoGetGlyphNameErr error

func tryCTGlyphInfoGetGlyphName(glyphInfo CTGlyphInfoRef) (corefoundation.CFStringRef, error) {
	if _cTGlyphInfoGetGlyphName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTGlyphInfoGetGlyphName", "10.5", _cTGlyphInfoGetGlyphNameErr)
	}
	return _cTGlyphInfoGetGlyphName(glyphInfo), nil
}

// CTGlyphInfoGetGlyphName retrieves the glyph name for a glyph info object, if that object exists.
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetGlyphName(_:)
func CTGlyphInfoGetGlyphName(glyphInfo CTGlyphInfoRef) corefoundation.CFStringRef {
	result, callErr := tryCTGlyphInfoGetGlyphName(glyphInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTGlyphInfoGetTypeID func() uint
var _cTGlyphInfoGetTypeIDErr error

func tryCTGlyphInfoGetTypeID() (uint, error) {
	if _cTGlyphInfoGetTypeID == nil {
		return 0, symbolCallError("CTGlyphInfoGetTypeID", "10.5", _cTGlyphInfoGetTypeIDErr)
	}
	return _cTGlyphInfoGetTypeID(), nil
}

// CTGlyphInfoGetTypeID returns the Core Foundation type identifier of the glyph info object
//
// See: https://developer.apple.com/documentation/CoreText/CTGlyphInfoGetTypeID()
func CTGlyphInfoGetTypeID() uint {
	result, callErr := tryCTGlyphInfoGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineCreateJustifiedLine func(line CTLineRef, justificationFactor float64, justificationWidth float64) CTLineRef
var _cTLineCreateJustifiedLineErr error

func tryCTLineCreateJustifiedLine(line CTLineRef, justificationFactor float64, justificationWidth float64) (CTLineRef, error) {
	if _cTLineCreateJustifiedLine == nil {
		return *new(CTLineRef), symbolCallError("CTLineCreateJustifiedLine", "10.5", _cTLineCreateJustifiedLineErr)
	}
	return _cTLineCreateJustifiedLine(line, justificationFactor, justificationWidth), nil
}

// CTLineCreateJustifiedLine creates a justified line from an existing line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineCreateJustifiedLine(_:_:_:)
func CTLineCreateJustifiedLine(line CTLineRef, justificationFactor float64, justificationWidth float64) CTLineRef {
	result, callErr := tryCTLineCreateJustifiedLine(line, justificationFactor, justificationWidth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineCreateTruncatedLine func(line CTLineRef, width float64, truncationType CTLineTruncationType, truncationToken CTLineRef) CTLineRef
var _cTLineCreateTruncatedLineErr error

func tryCTLineCreateTruncatedLine(line CTLineRef, width float64, truncationType CTLineTruncationType, truncationToken CTLineRef) (CTLineRef, error) {
	if _cTLineCreateTruncatedLine == nil {
		return *new(CTLineRef), symbolCallError("CTLineCreateTruncatedLine", "10.5", _cTLineCreateTruncatedLineErr)
	}
	return _cTLineCreateTruncatedLine(line, width, truncationType, truncationToken), nil
}

// CTLineCreateTruncatedLine creates a truncated line from an existing line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineCreateTruncatedLine(_:_:_:_:)
func CTLineCreateTruncatedLine(line CTLineRef, width float64, truncationType CTLineTruncationType, truncationToken CTLineRef) CTLineRef {
	result, callErr := tryCTLineCreateTruncatedLine(line, width, truncationType, truncationToken)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineCreateWithAttributedString func(attrString corefoundation.CFAttributedStringRef) CTLineRef
var _cTLineCreateWithAttributedStringErr error

func tryCTLineCreateWithAttributedString(attrString corefoundation.CFAttributedStringRef) (CTLineRef, error) {
	if _cTLineCreateWithAttributedString == nil {
		return *new(CTLineRef), symbolCallError("CTLineCreateWithAttributedString", "10.5", _cTLineCreateWithAttributedStringErr)
	}
	return _cTLineCreateWithAttributedString(attrString), nil
}

// CTLineCreateWithAttributedString creates a single immutable line object from an attributed string.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineCreateWithAttributedString(_:)
func CTLineCreateWithAttributedString(attrString corefoundation.CFAttributedStringRef) CTLineRef {
	result, callErr := tryCTLineCreateWithAttributedString(attrString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineDraw func(line CTLineRef, context coregraphics.CGContextRef)
var _cTLineDrawErr error

func tryCTLineDraw(line CTLineRef, context coregraphics.CGContextRef) error {
	if _cTLineDraw == nil {
		return symbolCallError("CTLineDraw", "10.5", _cTLineDrawErr)
	}
	_cTLineDraw(line, context)
	return nil
}

// CTLineDraw draws a complete line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineDraw(_:_:)
func CTLineDraw(line CTLineRef, context coregraphics.CGContextRef) {
	if callErr := tryCTLineDraw(line, context); callErr != nil {
		panic(callErr)
	}
}

var _cTLineEnumerateCaretOffsets func(line CTLineRef)
var _cTLineEnumerateCaretOffsetsErr error

func tryCTLineEnumerateCaretOffsets(line CTLineRef) error {
	if _cTLineEnumerateCaretOffsets == nil {
		return symbolCallError("CTLineEnumerateCaretOffsets", "10.11", _cTLineEnumerateCaretOffsetsErr)
	}
	_cTLineEnumerateCaretOffsets(line)
	return nil
}

// CTLineEnumerateCaretOffsets enumerates caret offsets for characters in a line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineEnumerateCaretOffsets(_:_:)
func CTLineEnumerateCaretOffsets(line CTLineRef) {
	if callErr := tryCTLineEnumerateCaretOffsets(line); callErr != nil {
		panic(callErr)
	}
}

var _cTLineGetBoundsWithOptions func(line CTLineRef, options CTLineBoundsOptions) corefoundation.CGRect
var _cTLineGetBoundsWithOptionsErr error

func tryCTLineGetBoundsWithOptions(line CTLineRef, options CTLineBoundsOptions) (corefoundation.CGRect, error) {
	if _cTLineGetBoundsWithOptions == nil {
		return corefoundation.CGRect{}, symbolCallError("CTLineGetBoundsWithOptions", "10.8", _cTLineGetBoundsWithOptionsErr)
	}
	return _cTLineGetBoundsWithOptions(line, options), nil
}

// CTLineGetBoundsWithOptions calculates the bounds for a line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetBoundsWithOptions(_:_:)
func CTLineGetBoundsWithOptions(line CTLineRef, options CTLineBoundsOptions) corefoundation.CGRect {
	result, callErr := tryCTLineGetBoundsWithOptions(line, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetGlyphCount func(line CTLineRef) int
var _cTLineGetGlyphCountErr error

func tryCTLineGetGlyphCount(line CTLineRef) (int, error) {
	if _cTLineGetGlyphCount == nil {
		return 0, symbolCallError("CTLineGetGlyphCount", "10.5", _cTLineGetGlyphCountErr)
	}
	return _cTLineGetGlyphCount(line), nil
}

// CTLineGetGlyphCount returns the total glyph count for the line object.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetGlyphCount(_:)
func CTLineGetGlyphCount(line CTLineRef) int {
	result, callErr := tryCTLineGetGlyphCount(line)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetGlyphRuns func(line CTLineRef) corefoundation.CFArrayRef
var _cTLineGetGlyphRunsErr error

func tryCTLineGetGlyphRuns(line CTLineRef) (corefoundation.CFArrayRef, error) {
	if _cTLineGetGlyphRuns == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CTLineGetGlyphRuns", "10.5", _cTLineGetGlyphRunsErr)
	}
	return _cTLineGetGlyphRuns(line), nil
}

// CTLineGetGlyphRuns returns the array of glyph runs that make up the line object.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetGlyphRuns(_:)
func CTLineGetGlyphRuns(line CTLineRef) corefoundation.CFArrayRef {
	result, callErr := tryCTLineGetGlyphRuns(line)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetImageBounds func(line CTLineRef, context coregraphics.CGContextRef) corefoundation.CGRect
var _cTLineGetImageBoundsErr error

func tryCTLineGetImageBounds(line CTLineRef, context coregraphics.CGContextRef) (corefoundation.CGRect, error) {
	if _cTLineGetImageBounds == nil {
		return corefoundation.CGRect{}, symbolCallError("CTLineGetImageBounds", "10.5", _cTLineGetImageBoundsErr)
	}
	return _cTLineGetImageBounds(line, context), nil
}

// CTLineGetImageBounds calculates the image bounds for a line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetImageBounds(_:_:)
func CTLineGetImageBounds(line CTLineRef, context coregraphics.CGContextRef) corefoundation.CGRect {
	result, callErr := tryCTLineGetImageBounds(line, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetOffsetForStringIndex func(line CTLineRef, charIndex int, secondaryOffset *float64) float64
var _cTLineGetOffsetForStringIndexErr error

func tryCTLineGetOffsetForStringIndex(line CTLineRef, charIndex int, secondaryOffset *float64) (float64, error) {
	if _cTLineGetOffsetForStringIndex == nil {
		return 0.0, symbolCallError("CTLineGetOffsetForStringIndex", "10.5", _cTLineGetOffsetForStringIndexErr)
	}
	return _cTLineGetOffsetForStringIndex(line, charIndex, secondaryOffset), nil
}

// CTLineGetOffsetForStringIndex determines the graphical offset or offsets for a string index.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetOffsetForStringIndex(_:_:_:)
func CTLineGetOffsetForStringIndex(line CTLineRef, charIndex int, secondaryOffset *float64) float64 {
	result, callErr := tryCTLineGetOffsetForStringIndex(line, charIndex, secondaryOffset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetPenOffsetForFlush func(line CTLineRef, flushFactor float64, flushWidth float64) float64
var _cTLineGetPenOffsetForFlushErr error

func tryCTLineGetPenOffsetForFlush(line CTLineRef, flushFactor float64, flushWidth float64) (float64, error) {
	if _cTLineGetPenOffsetForFlush == nil {
		return 0.0, symbolCallError("CTLineGetPenOffsetForFlush", "10.5", _cTLineGetPenOffsetForFlushErr)
	}
	return _cTLineGetPenOffsetForFlush(line, flushFactor, flushWidth), nil
}

// CTLineGetPenOffsetForFlush gets the pen offset required to draw flush text.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetPenOffsetForFlush(_:_:_:)
func CTLineGetPenOffsetForFlush(line CTLineRef, flushFactor float64, flushWidth float64) float64 {
	result, callErr := tryCTLineGetPenOffsetForFlush(line, flushFactor, flushWidth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetStringIndexForPosition func(line CTLineRef, position corefoundation.CGPoint) int
var _cTLineGetStringIndexForPositionErr error

func tryCTLineGetStringIndexForPosition(line CTLineRef, position corefoundation.CGPoint) (int, error) {
	if _cTLineGetStringIndexForPosition == nil {
		return 0, symbolCallError("CTLineGetStringIndexForPosition", "10.5", _cTLineGetStringIndexForPositionErr)
	}
	return _cTLineGetStringIndexForPosition(line, position), nil
}

// CTLineGetStringIndexForPosition performs hit testing.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetStringIndexForPosition(_:_:)
func CTLineGetStringIndexForPosition(line CTLineRef, position corefoundation.CGPoint) int {
	result, callErr := tryCTLineGetStringIndexForPosition(line, position)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetStringRange func(line CTLineRef) corefoundation.CFRange
var _cTLineGetStringRangeErr error

func tryCTLineGetStringRange(line CTLineRef) (corefoundation.CFRange, error) {
	if _cTLineGetStringRange == nil {
		return corefoundation.CFRange{}, symbolCallError("CTLineGetStringRange", "10.5", _cTLineGetStringRangeErr)
	}
	return _cTLineGetStringRange(line), nil
}

// CTLineGetStringRange gets the range of characters that originally spawned the glyphs in the line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetStringRange(_:)
func CTLineGetStringRange(line CTLineRef) corefoundation.CFRange {
	result, callErr := tryCTLineGetStringRange(line)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetTrailingWhitespaceWidth func(line CTLineRef) float64
var _cTLineGetTrailingWhitespaceWidthErr error

func tryCTLineGetTrailingWhitespaceWidth(line CTLineRef) (float64, error) {
	if _cTLineGetTrailingWhitespaceWidth == nil {
		return 0.0, symbolCallError("CTLineGetTrailingWhitespaceWidth", "10.5", _cTLineGetTrailingWhitespaceWidthErr)
	}
	return _cTLineGetTrailingWhitespaceWidth(line), nil
}

// CTLineGetTrailingWhitespaceWidth returns the trailing whitespace width for a line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetTrailingWhitespaceWidth(_:)
func CTLineGetTrailingWhitespaceWidth(line CTLineRef) float64 {
	result, callErr := tryCTLineGetTrailingWhitespaceWidth(line)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetTypeID func() uint
var _cTLineGetTypeIDErr error

func tryCTLineGetTypeID() (uint, error) {
	if _cTLineGetTypeID == nil {
		return 0, symbolCallError("CTLineGetTypeID", "10.5", _cTLineGetTypeIDErr)
	}
	return _cTLineGetTypeID(), nil
}

// CTLineGetTypeID returns the Core Foundation type identifier of the line object.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetTypeID()
func CTLineGetTypeID() uint {
	result, callErr := tryCTLineGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTLineGetTypographicBounds func(line CTLineRef, ascent *float64, descent *float64, leading *float64) float64
var _cTLineGetTypographicBoundsErr error

func tryCTLineGetTypographicBounds(line CTLineRef, ascent *float64, descent *float64, leading *float64) (float64, error) {
	if _cTLineGetTypographicBounds == nil {
		return 0.0, symbolCallError("CTLineGetTypographicBounds", "10.5", _cTLineGetTypographicBoundsErr)
	}
	return _cTLineGetTypographicBounds(line, ascent, descent, leading), nil
}

// CTLineGetTypographicBounds calculates the typographic bounds of a line.
//
// See: https://developer.apple.com/documentation/CoreText/CTLineGetTypographicBounds(_:_:_:_:)
func CTLineGetTypographicBounds(line CTLineRef, ascent *float64, descent *float64, leading *float64) float64 {
	result, callErr := tryCTLineGetTypographicBounds(line, ascent, descent, leading)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTParagraphStyleCreate func(settings *CTParagraphStyleSetting, settingCount uintptr) CTParagraphStyleRef
var _cTParagraphStyleCreateErr error

func tryCTParagraphStyleCreate(settings *CTParagraphStyleSetting, settingCount uintptr) (CTParagraphStyleRef, error) {
	if _cTParagraphStyleCreate == nil {
		return *new(CTParagraphStyleRef), symbolCallError("CTParagraphStyleCreate", "10.5", _cTParagraphStyleCreateErr)
	}
	return _cTParagraphStyleCreate(settings, settingCount), nil
}

// CTParagraphStyleCreate creates an immutable paragraph style.
//
// See: https://developer.apple.com/documentation/CoreText/CTParagraphStyleCreate(_:_:)
func CTParagraphStyleCreate(settings *CTParagraphStyleSetting, settingCount uintptr) CTParagraphStyleRef {
	result, callErr := tryCTParagraphStyleCreate(settings, settingCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTParagraphStyleCreateCopy func(paragraphStyle CTParagraphStyleRef) CTParagraphStyleRef
var _cTParagraphStyleCreateCopyErr error

func tryCTParagraphStyleCreateCopy(paragraphStyle CTParagraphStyleRef) (CTParagraphStyleRef, error) {
	if _cTParagraphStyleCreateCopy == nil {
		return *new(CTParagraphStyleRef), symbolCallError("CTParagraphStyleCreateCopy", "10.5", _cTParagraphStyleCreateCopyErr)
	}
	return _cTParagraphStyleCreateCopy(paragraphStyle), nil
}

// CTParagraphStyleCreateCopy creates an immutable copy of a paragraph style.
//
// See: https://developer.apple.com/documentation/CoreText/CTParagraphStyleCreateCopy(_:)
func CTParagraphStyleCreateCopy(paragraphStyle CTParagraphStyleRef) CTParagraphStyleRef {
	result, callErr := tryCTParagraphStyleCreateCopy(paragraphStyle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTParagraphStyleGetTypeID func() uint
var _cTParagraphStyleGetTypeIDErr error

func tryCTParagraphStyleGetTypeID() (uint, error) {
	if _cTParagraphStyleGetTypeID == nil {
		return 0, symbolCallError("CTParagraphStyleGetTypeID", "10.5", _cTParagraphStyleGetTypeIDErr)
	}
	return _cTParagraphStyleGetTypeID(), nil
}

// CTParagraphStyleGetTypeID returns the Core Foundation type identifier of the paragraph style object.
//
// See: https://developer.apple.com/documentation/CoreText/CTParagraphStyleGetTypeID()
func CTParagraphStyleGetTypeID() uint {
	result, callErr := tryCTParagraphStyleGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTParagraphStyleGetValueForSpecifier func(paragraphStyle CTParagraphStyleRef, spec CTParagraphStyleSpecifier, valueBufferSize uintptr, valueBuffer unsafe.Pointer) bool
var _cTParagraphStyleGetValueForSpecifierErr error

func tryCTParagraphStyleGetValueForSpecifier(paragraphStyle CTParagraphStyleRef, spec CTParagraphStyleSpecifier, valueBufferSize uintptr, valueBuffer unsafe.Pointer) (bool, error) {
	if _cTParagraphStyleGetValueForSpecifier == nil {
		return false, symbolCallError("CTParagraphStyleGetValueForSpecifier", "10.5", _cTParagraphStyleGetValueForSpecifierErr)
	}
	return _cTParagraphStyleGetValueForSpecifier(paragraphStyle, spec, valueBufferSize, valueBuffer), nil
}

// CTParagraphStyleGetValueForSpecifier obtains the current value for a single setting specifier.
//
// See: https://developer.apple.com/documentation/CoreText/CTParagraphStyleGetValueForSpecifier(_:_:_:_:)
func CTParagraphStyleGetValueForSpecifier(paragraphStyle CTParagraphStyleRef, spec CTParagraphStyleSpecifier, valueBufferSize uintptr, valueBuffer unsafe.Pointer) bool {
	result, callErr := tryCTParagraphStyleGetValueForSpecifier(paragraphStyle, spec, valueBufferSize, valueBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationCreate func(alignment CTRubyAlignment, overhang CTRubyOverhang, sizeFactor float64, text corefoundation.CFStringRef) CTRubyAnnotationRef
var _cTRubyAnnotationCreateErr error

func tryCTRubyAnnotationCreate(alignment CTRubyAlignment, overhang CTRubyOverhang, sizeFactor float64, text corefoundation.CFStringRef) (CTRubyAnnotationRef, error) {
	if _cTRubyAnnotationCreate == nil {
		return *new(CTRubyAnnotationRef), symbolCallError("CTRubyAnnotationCreate", "10.10", _cTRubyAnnotationCreateErr)
	}
	return _cTRubyAnnotationCreate(alignment, overhang, sizeFactor, text), nil
}

// CTRubyAnnotationCreate creates an immutable ruby annotation object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreate(_:_:_:_:)
func CTRubyAnnotationCreate(alignment CTRubyAlignment, overhang CTRubyOverhang, sizeFactor float64, text corefoundation.CFStringRef) CTRubyAnnotationRef {
	result, callErr := tryCTRubyAnnotationCreate(alignment, overhang, sizeFactor, text)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationCreateCopy func(rubyAnnotation CTRubyAnnotationRef) CTRubyAnnotationRef
var _cTRubyAnnotationCreateCopyErr error

func tryCTRubyAnnotationCreateCopy(rubyAnnotation CTRubyAnnotationRef) (CTRubyAnnotationRef, error) {
	if _cTRubyAnnotationCreateCopy == nil {
		return *new(CTRubyAnnotationRef), symbolCallError("CTRubyAnnotationCreateCopy", "10.10", _cTRubyAnnotationCreateCopyErr)
	}
	return _cTRubyAnnotationCreateCopy(rubyAnnotation), nil
}

// CTRubyAnnotationCreateCopy creates an immutable copy of a ruby annotation object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreateCopy(_:)
func CTRubyAnnotationCreateCopy(rubyAnnotation CTRubyAnnotationRef) CTRubyAnnotationRef {
	result, callErr := tryCTRubyAnnotationCreateCopy(rubyAnnotation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationCreateWithAttributes func(alignment CTRubyAlignment, overhang CTRubyOverhang, position CTRubyPosition, string_ corefoundation.CFStringRef, attributes corefoundation.CFDictionaryRef) CTRubyAnnotationRef
var _cTRubyAnnotationCreateWithAttributesErr error

func tryCTRubyAnnotationCreateWithAttributes(alignment CTRubyAlignment, overhang CTRubyOverhang, position CTRubyPosition, string_ corefoundation.CFStringRef, attributes corefoundation.CFDictionaryRef) (CTRubyAnnotationRef, error) {
	if _cTRubyAnnotationCreateWithAttributes == nil {
		return *new(CTRubyAnnotationRef), symbolCallError("CTRubyAnnotationCreateWithAttributes", "10.12", _cTRubyAnnotationCreateWithAttributesErr)
	}
	return _cTRubyAnnotationCreateWithAttributes(alignment, overhang, position, string_, attributes), nil
}

// CTRubyAnnotationCreateWithAttributes creates an immutable ruby annotation object with the specified attributes.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationCreateWithAttributes(_:_:_:_:_:)
func CTRubyAnnotationCreateWithAttributes(alignment CTRubyAlignment, overhang CTRubyOverhang, position CTRubyPosition, string_ corefoundation.CFStringRef, attributes corefoundation.CFDictionaryRef) CTRubyAnnotationRef {
	result, callErr := tryCTRubyAnnotationCreateWithAttributes(alignment, overhang, position, string_, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationGetAlignment func(rubyAnnotation CTRubyAnnotationRef) CTRubyAlignment
var _cTRubyAnnotationGetAlignmentErr error

func tryCTRubyAnnotationGetAlignment(rubyAnnotation CTRubyAnnotationRef) (CTRubyAlignment, error) {
	if _cTRubyAnnotationGetAlignment == nil {
		return *new(CTRubyAlignment), symbolCallError("CTRubyAnnotationGetAlignment", "10.10", _cTRubyAnnotationGetAlignmentErr)
	}
	return _cTRubyAnnotationGetAlignment(rubyAnnotation), nil
}

// CTRubyAnnotationGetAlignment retrieves the alignment value of a ruby annotation object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetAlignment(_:)
func CTRubyAnnotationGetAlignment(rubyAnnotation CTRubyAnnotationRef) CTRubyAlignment {
	result, callErr := tryCTRubyAnnotationGetAlignment(rubyAnnotation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationGetOverhang func(rubyAnnotation CTRubyAnnotationRef) CTRubyOverhang
var _cTRubyAnnotationGetOverhangErr error

func tryCTRubyAnnotationGetOverhang(rubyAnnotation CTRubyAnnotationRef) (CTRubyOverhang, error) {
	if _cTRubyAnnotationGetOverhang == nil {
		return *new(CTRubyOverhang), symbolCallError("CTRubyAnnotationGetOverhang", "10.10", _cTRubyAnnotationGetOverhangErr)
	}
	return _cTRubyAnnotationGetOverhang(rubyAnnotation), nil
}

// CTRubyAnnotationGetOverhang retrieves the overhang value of a ruby annotation object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetOverhang(_:)
func CTRubyAnnotationGetOverhang(rubyAnnotation CTRubyAnnotationRef) CTRubyOverhang {
	result, callErr := tryCTRubyAnnotationGetOverhang(rubyAnnotation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationGetSizeFactor func(rubyAnnotation CTRubyAnnotationRef) float64
var _cTRubyAnnotationGetSizeFactorErr error

func tryCTRubyAnnotationGetSizeFactor(rubyAnnotation CTRubyAnnotationRef) (float64, error) {
	if _cTRubyAnnotationGetSizeFactor == nil {
		return 0.0, symbolCallError("CTRubyAnnotationGetSizeFactor", "10.10", _cTRubyAnnotationGetSizeFactorErr)
	}
	return _cTRubyAnnotationGetSizeFactor(rubyAnnotation), nil
}

// CTRubyAnnotationGetSizeFactor retrieves the size factor of a ruby annotation object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetSizeFactor(_:)
func CTRubyAnnotationGetSizeFactor(rubyAnnotation CTRubyAnnotationRef) float64 {
	result, callErr := tryCTRubyAnnotationGetSizeFactor(rubyAnnotation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationGetTextForPosition func(rubyAnnotation CTRubyAnnotationRef, position CTRubyPosition) corefoundation.CFStringRef
var _cTRubyAnnotationGetTextForPositionErr error

func tryCTRubyAnnotationGetTextForPosition(rubyAnnotation CTRubyAnnotationRef, position CTRubyPosition) (corefoundation.CFStringRef, error) {
	if _cTRubyAnnotationGetTextForPosition == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CTRubyAnnotationGetTextForPosition", "10.10", _cTRubyAnnotationGetTextForPositionErr)
	}
	return _cTRubyAnnotationGetTextForPosition(rubyAnnotation, position), nil
}

// CTRubyAnnotationGetTextForPosition retrieves the ruby text for a particular position in a ruby annotation.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetTextForPosition(_:_:)
func CTRubyAnnotationGetTextForPosition(rubyAnnotation CTRubyAnnotationRef, position CTRubyPosition) corefoundation.CFStringRef {
	result, callErr := tryCTRubyAnnotationGetTextForPosition(rubyAnnotation, position)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRubyAnnotationGetTypeID func() uint
var _cTRubyAnnotationGetTypeIDErr error

func tryCTRubyAnnotationGetTypeID() (uint, error) {
	if _cTRubyAnnotationGetTypeID == nil {
		return 0, symbolCallError("CTRubyAnnotationGetTypeID", "10.10", _cTRubyAnnotationGetTypeIDErr)
	}
	return _cTRubyAnnotationGetTypeID(), nil
}

// CTRubyAnnotationGetTypeID retrieves the type of the ruby annotation object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRubyAnnotationGetTypeID()
func CTRubyAnnotationGetTypeID() uint {
	result, callErr := tryCTRubyAnnotationGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunDelegateCreate func(callbacks *CTRunDelegateCallbacks, refCon uintptr) CTRunDelegateRef
var _cTRunDelegateCreateErr error

func tryCTRunDelegateCreate(callbacks *CTRunDelegateCallbacks, refCon uintptr) (CTRunDelegateRef, error) {
	if _cTRunDelegateCreate == nil {
		return *new(CTRunDelegateRef), symbolCallError("CTRunDelegateCreate", "10.5", _cTRunDelegateCreateErr)
	}
	return _cTRunDelegateCreate(callbacks, refCon), nil
}

// CTRunDelegateCreate creates an immutable instance of a run delegate.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateCreate(_:_:)
func CTRunDelegateCreate(callbacks *CTRunDelegateCallbacks, refCon uintptr) CTRunDelegateRef {
	result, callErr := tryCTRunDelegateCreate(callbacks, refCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunDelegateGetRefCon func(runDelegate CTRunDelegateRef) unsafe.Pointer
var _cTRunDelegateGetRefConErr error

func tryCTRunDelegateGetRefCon(runDelegate CTRunDelegateRef) (unsafe.Pointer, error) {
	if _cTRunDelegateGetRefCon == nil {
		return nil, symbolCallError("CTRunDelegateGetRefCon", "10.5", _cTRunDelegateGetRefConErr)
	}
	return _cTRunDelegateGetRefCon(runDelegate), nil
}

// CTRunDelegateGetRefCon returns a run delegate’s “refCon” value.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetRefCon(_:)
func CTRunDelegateGetRefCon(runDelegate CTRunDelegateRef) unsafe.Pointer {
	result, callErr := tryCTRunDelegateGetRefCon(runDelegate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunDelegateGetTypeID func() uint
var _cTRunDelegateGetTypeIDErr error

func tryCTRunDelegateGetTypeID() (uint, error) {
	if _cTRunDelegateGetTypeID == nil {
		return 0, symbolCallError("CTRunDelegateGetTypeID", "10.5", _cTRunDelegateGetTypeIDErr)
	}
	return _cTRunDelegateGetTypeID(), nil
}

// CTRunDelegateGetTypeID returns the type of CTRunDelegate objects.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDelegateGetTypeID()
func CTRunDelegateGetTypeID() uint {
	result, callErr := tryCTRunDelegateGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunDraw func(run CTRunRef, context coregraphics.CGContextRef, range_ corefoundation.CFRange)
var _cTRunDrawErr error

func tryCTRunDraw(run CTRunRef, context coregraphics.CGContextRef, range_ corefoundation.CFRange) error {
	if _cTRunDraw == nil {
		return symbolCallError("CTRunDraw", "10.5", _cTRunDrawErr)
	}
	_cTRunDraw(run, context, range_)
	return nil
}

// CTRunDraw draws a complete run or part of one.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunDraw(_:_:_:)
func CTRunDraw(run CTRunRef, context coregraphics.CGContextRef, range_ corefoundation.CFRange) {
	if callErr := tryCTRunDraw(run, context, range_); callErr != nil {
		panic(callErr)
	}
}

var _cTRunGetAdvances func(run CTRunRef, range_ corefoundation.CFRange, buffer *corefoundation.CGSize)
var _cTRunGetAdvancesErr error

func tryCTRunGetAdvances(run CTRunRef, range_ corefoundation.CFRange, buffer *corefoundation.CGSize) error {
	if _cTRunGetAdvances == nil {
		return symbolCallError("CTRunGetAdvances", "10.5", _cTRunGetAdvancesErr)
	}
	_cTRunGetAdvances(run, range_, buffer)
	return nil
}

// CTRunGetAdvances copies a range of glyph advances into a user-provided buffer.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetAdvances(_:_:_:)
func CTRunGetAdvances(run CTRunRef, range_ corefoundation.CFRange, buffer *corefoundation.CGSize) {
	if callErr := tryCTRunGetAdvances(run, range_, buffer); callErr != nil {
		panic(callErr)
	}
}

var _cTRunGetAdvancesPtr func(run CTRunRef) *corefoundation.CGSize
var _cTRunGetAdvancesPtrErr error

func tryCTRunGetAdvancesPtr(run CTRunRef) (*corefoundation.CGSize, error) {
	if _cTRunGetAdvancesPtr == nil {
		return nil, symbolCallError("CTRunGetAdvancesPtr", "10.5", _cTRunGetAdvancesPtrErr)
	}
	return _cTRunGetAdvancesPtr(run), nil
}

// CTRunGetAdvancesPtr returns a direct pointer for the glyph advance array stored in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetAdvancesPtr(_:)
func CTRunGetAdvancesPtr(run CTRunRef) *corefoundation.CGSize {
	result, callErr := tryCTRunGetAdvancesPtr(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetAttributes func(run CTRunRef) corefoundation.CFDictionaryRef
var _cTRunGetAttributesErr error

func tryCTRunGetAttributes(run CTRunRef) (corefoundation.CFDictionaryRef, error) {
	if _cTRunGetAttributes == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CTRunGetAttributes", "10.5", _cTRunGetAttributesErr)
	}
	return _cTRunGetAttributes(run), nil
}

// CTRunGetAttributes returns the attribute dictionary that was used to create the glyph run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetAttributes(_:)
func CTRunGetAttributes(run CTRunRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCTRunGetAttributes(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetBaseAdvancesAndOrigins func(runRef CTRunRef, range_ corefoundation.CFRange, advancesBuffer *corefoundation.CGSize, originsBuffer *corefoundation.CGPoint)
var _cTRunGetBaseAdvancesAndOriginsErr error

func tryCTRunGetBaseAdvancesAndOrigins(runRef CTRunRef, range_ corefoundation.CFRange, advancesBuffer *corefoundation.CGSize, originsBuffer *corefoundation.CGPoint) error {
	if _cTRunGetBaseAdvancesAndOrigins == nil {
		return symbolCallError("CTRunGetBaseAdvancesAndOrigins", "10.11", _cTRunGetBaseAdvancesAndOriginsErr)
	}
	_cTRunGetBaseAdvancesAndOrigins(runRef, range_, advancesBuffer, originsBuffer)
	return nil
}

// CTRunGetBaseAdvancesAndOrigins copies a range of base advances and origins into user-provided buffers.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetBaseAdvancesAndOrigins(_:_:_:_:)
func CTRunGetBaseAdvancesAndOrigins(runRef CTRunRef, range_ corefoundation.CFRange, advancesBuffer *corefoundation.CGSize, originsBuffer *corefoundation.CGPoint) {
	if callErr := tryCTRunGetBaseAdvancesAndOrigins(runRef, range_, advancesBuffer, originsBuffer); callErr != nil {
		panic(callErr)
	}
}

var _cTRunGetGlyphCount func(run CTRunRef) int
var _cTRunGetGlyphCountErr error

func tryCTRunGetGlyphCount(run CTRunRef) (int, error) {
	if _cTRunGetGlyphCount == nil {
		return 0, symbolCallError("CTRunGetGlyphCount", "10.5", _cTRunGetGlyphCountErr)
	}
	return _cTRunGetGlyphCount(run), nil
}

// CTRunGetGlyphCount gets the glyph count for the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphCount(_:)
func CTRunGetGlyphCount(run CTRunRef) int {
	result, callErr := tryCTRunGetGlyphCount(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetGlyphs func(run CTRunRef, range_ corefoundation.CFRange, buffer unsafe.Pointer)
var _cTRunGetGlyphsErr error

func tryCTRunGetGlyphs(run CTRunRef, range_ corefoundation.CFRange, buffer unsafe.Pointer) error {
	if _cTRunGetGlyphs == nil {
		return symbolCallError("CTRunGetGlyphs", "10.5", _cTRunGetGlyphsErr)
	}
	_cTRunGetGlyphs(run, range_, buffer)
	return nil
}

// CTRunGetGlyphs copies a range of glyphs into a user-provided buffer.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphs(_:_:_:)
func CTRunGetGlyphs(run CTRunRef, range_ corefoundation.CFRange, buffer unsafe.Pointer) {
	if callErr := tryCTRunGetGlyphs(run, range_, buffer); callErr != nil {
		panic(callErr)
	}
}

var _cTRunGetGlyphsPtr func(run CTRunRef) unsafe.Pointer
var _cTRunGetGlyphsPtrErr error

func tryCTRunGetGlyphsPtr(run CTRunRef) (unsafe.Pointer, error) {
	if _cTRunGetGlyphsPtr == nil {
		return nil, symbolCallError("CTRunGetGlyphsPtr", "10.5", _cTRunGetGlyphsPtrErr)
	}
	return _cTRunGetGlyphsPtr(run), nil
}

// CTRunGetGlyphsPtr returns a direct pointer for the glyph array stored in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetGlyphsPtr(_:)
func CTRunGetGlyphsPtr(run CTRunRef) unsafe.Pointer {
	result, callErr := tryCTRunGetGlyphsPtr(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetImageBounds func(run CTRunRef, context coregraphics.CGContextRef, range_ corefoundation.CFRange) corefoundation.CGRect
var _cTRunGetImageBoundsErr error

func tryCTRunGetImageBounds(run CTRunRef, context coregraphics.CGContextRef, range_ corefoundation.CFRange) (corefoundation.CGRect, error) {
	if _cTRunGetImageBounds == nil {
		return corefoundation.CGRect{}, symbolCallError("CTRunGetImageBounds", "10.5", _cTRunGetImageBoundsErr)
	}
	return _cTRunGetImageBounds(run, context, range_), nil
}

// CTRunGetImageBounds calculates the image bounds for a glyph range.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetImageBounds(_:_:_:)
func CTRunGetImageBounds(run CTRunRef, context coregraphics.CGContextRef, range_ corefoundation.CFRange) corefoundation.CGRect {
	result, callErr := tryCTRunGetImageBounds(run, context, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetPositions func(run CTRunRef, range_ corefoundation.CFRange, buffer *corefoundation.CGPoint)
var _cTRunGetPositionsErr error

func tryCTRunGetPositions(run CTRunRef, range_ corefoundation.CFRange, buffer *corefoundation.CGPoint) error {
	if _cTRunGetPositions == nil {
		return symbolCallError("CTRunGetPositions", "10.5", _cTRunGetPositionsErr)
	}
	_cTRunGetPositions(run, range_, buffer)
	return nil
}

// CTRunGetPositions copies a range of glyph positions into a user-provided buffer.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetPositions(_:_:_:)
func CTRunGetPositions(run CTRunRef, range_ corefoundation.CFRange, buffer *corefoundation.CGPoint) {
	if callErr := tryCTRunGetPositions(run, range_, buffer); callErr != nil {
		panic(callErr)
	}
}

var _cTRunGetPositionsPtr func(run CTRunRef) *corefoundation.CGPoint
var _cTRunGetPositionsPtrErr error

func tryCTRunGetPositionsPtr(run CTRunRef) (*corefoundation.CGPoint, error) {
	if _cTRunGetPositionsPtr == nil {
		return nil, symbolCallError("CTRunGetPositionsPtr", "10.5", _cTRunGetPositionsPtrErr)
	}
	return _cTRunGetPositionsPtr(run), nil
}

// CTRunGetPositionsPtr returns a direct pointer for the glyph position array stored in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetPositionsPtr(_:)
func CTRunGetPositionsPtr(run CTRunRef) *corefoundation.CGPoint {
	result, callErr := tryCTRunGetPositionsPtr(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetStatus func(run CTRunRef) CTRunStatus
var _cTRunGetStatusErr error

func tryCTRunGetStatus(run CTRunRef) (CTRunStatus, error) {
	if _cTRunGetStatus == nil {
		return *new(CTRunStatus), symbolCallError("CTRunGetStatus", "10.5", _cTRunGetStatusErr)
	}
	return _cTRunGetStatus(run), nil
}

// CTRunGetStatus returns the run’s status.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetStatus(_:)
func CTRunGetStatus(run CTRunRef) CTRunStatus {
	result, callErr := tryCTRunGetStatus(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetStringIndices func(run CTRunRef, range_ corefoundation.CFRange, buffer *int)
var _cTRunGetStringIndicesErr error

func tryCTRunGetStringIndices(run CTRunRef, range_ corefoundation.CFRange, buffer *int) error {
	if _cTRunGetStringIndices == nil {
		return symbolCallError("CTRunGetStringIndices", "10.5", _cTRunGetStringIndicesErr)
	}
	_cTRunGetStringIndices(run, range_, buffer)
	return nil
}

// CTRunGetStringIndices copies a range of string indices into a user-provided buffer.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetStringIndices(_:_:_:)
func CTRunGetStringIndices(run CTRunRef, range_ corefoundation.CFRange, buffer *int) {
	if callErr := tryCTRunGetStringIndices(run, range_, buffer); callErr != nil {
		panic(callErr)
	}
}

var _cTRunGetStringIndicesPtr func(run CTRunRef) *int
var _cTRunGetStringIndicesPtrErr error

func tryCTRunGetStringIndicesPtr(run CTRunRef) (*int, error) {
	if _cTRunGetStringIndicesPtr == nil {
		return nil, symbolCallError("CTRunGetStringIndicesPtr", "10.5", _cTRunGetStringIndicesPtrErr)
	}
	return _cTRunGetStringIndicesPtr(run), nil
}

// CTRunGetStringIndicesPtr returns a direct pointer for the string indices stored in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetStringIndicesPtr(_:)
func CTRunGetStringIndicesPtr(run CTRunRef) *int {
	result, callErr := tryCTRunGetStringIndicesPtr(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetStringRange func(run CTRunRef) corefoundation.CFRange
var _cTRunGetStringRangeErr error

func tryCTRunGetStringRange(run CTRunRef) (corefoundation.CFRange, error) {
	if _cTRunGetStringRange == nil {
		return corefoundation.CFRange{}, symbolCallError("CTRunGetStringRange", "10.5", _cTRunGetStringRangeErr)
	}
	return _cTRunGetStringRange(run), nil
}

// CTRunGetStringRange gets the range of characters that originally spawned the glyphs in the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetStringRange(_:)
func CTRunGetStringRange(run CTRunRef) corefoundation.CFRange {
	result, callErr := tryCTRunGetStringRange(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetTextMatrix func(run CTRunRef) corefoundation.CGAffineTransform
var _cTRunGetTextMatrixErr error

func tryCTRunGetTextMatrix(run CTRunRef) (corefoundation.CGAffineTransform, error) {
	if _cTRunGetTextMatrix == nil {
		return corefoundation.CGAffineTransform{}, symbolCallError("CTRunGetTextMatrix", "10.5", _cTRunGetTextMatrixErr)
	}
	return _cTRunGetTextMatrix(run), nil
}

// CTRunGetTextMatrix returns the text matrix needed to draw this run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetTextMatrix(_:)
func CTRunGetTextMatrix(run CTRunRef) corefoundation.CGAffineTransform {
	result, callErr := tryCTRunGetTextMatrix(run)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetTypeID func() uint
var _cTRunGetTypeIDErr error

func tryCTRunGetTypeID() (uint, error) {
	if _cTRunGetTypeID == nil {
		return 0, symbolCallError("CTRunGetTypeID", "10.5", _cTRunGetTypeIDErr)
	}
	return _cTRunGetTypeID(), nil
}

// CTRunGetTypeID returns the Core Foundation type identifier of the run object.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetTypeID()
func CTRunGetTypeID() uint {
	result, callErr := tryCTRunGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTRunGetTypographicBounds func(run CTRunRef, range_ corefoundation.CFRange, ascent *float64, descent *float64, leading *float64) float64
var _cTRunGetTypographicBoundsErr error

func tryCTRunGetTypographicBounds(run CTRunRef, range_ corefoundation.CFRange, ascent *float64, descent *float64, leading *float64) (float64, error) {
	if _cTRunGetTypographicBounds == nil {
		return 0.0, symbolCallError("CTRunGetTypographicBounds", "10.5", _cTRunGetTypographicBoundsErr)
	}
	return _cTRunGetTypographicBounds(run, range_, ascent, descent, leading), nil
}

// CTRunGetTypographicBounds gets the typographic bounds of the run.
//
// See: https://developer.apple.com/documentation/CoreText/CTRunGetTypographicBounds(_:_:_:_:_:)
func CTRunGetTypographicBounds(run CTRunRef, range_ corefoundation.CFRange, ascent *float64, descent *float64, leading *float64) float64 {
	result, callErr := tryCTRunGetTypographicBounds(run, range_, ascent, descent, leading)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTextTabCreate func(alignment CTTextAlignment, location float64, options corefoundation.CFDictionaryRef) CTTextTabRef
var _cTTextTabCreateErr error

func tryCTTextTabCreate(alignment CTTextAlignment, location float64, options corefoundation.CFDictionaryRef) (CTTextTabRef, error) {
	if _cTTextTabCreate == nil {
		return *new(CTTextTabRef), symbolCallError("CTTextTabCreate", "10.5", _cTTextTabCreateErr)
	}
	return _cTTextTabCreate(alignment, location, options), nil
}

// CTTextTabCreate creates and initializes a new text tab object.
//
// See: https://developer.apple.com/documentation/CoreText/CTTextTabCreate(_:_:_:)
func CTTextTabCreate(alignment CTTextAlignment, location float64, options corefoundation.CFDictionaryRef) CTTextTabRef {
	result, callErr := tryCTTextTabCreate(alignment, location, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTextTabGetAlignment func(tab CTTextTabRef) CTTextAlignment
var _cTTextTabGetAlignmentErr error

func tryCTTextTabGetAlignment(tab CTTextTabRef) (CTTextAlignment, error) {
	if _cTTextTabGetAlignment == nil {
		return *new(CTTextAlignment), symbolCallError("CTTextTabGetAlignment", "10.5", _cTTextTabGetAlignmentErr)
	}
	return _cTTextTabGetAlignment(tab), nil
}

// CTTextTabGetAlignment returns the text alignment of the tab.
//
// See: https://developer.apple.com/documentation/CoreText/CTTextTabGetAlignment(_:)
func CTTextTabGetAlignment(tab CTTextTabRef) CTTextAlignment {
	result, callErr := tryCTTextTabGetAlignment(tab)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTextTabGetLocation func(tab CTTextTabRef) float64
var _cTTextTabGetLocationErr error

func tryCTTextTabGetLocation(tab CTTextTabRef) (float64, error) {
	if _cTTextTabGetLocation == nil {
		return 0.0, symbolCallError("CTTextTabGetLocation", "10.5", _cTTextTabGetLocationErr)
	}
	return _cTTextTabGetLocation(tab), nil
}

// CTTextTabGetLocation returns the tab’s ruler location.
//
// See: https://developer.apple.com/documentation/CoreText/CTTextTabGetLocation(_:)
func CTTextTabGetLocation(tab CTTextTabRef) float64 {
	result, callErr := tryCTTextTabGetLocation(tab)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTextTabGetOptions func(tab CTTextTabRef) corefoundation.CFDictionaryRef
var _cTTextTabGetOptionsErr error

func tryCTTextTabGetOptions(tab CTTextTabRef) (corefoundation.CFDictionaryRef, error) {
	if _cTTextTabGetOptions == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CTTextTabGetOptions", "10.5", _cTTextTabGetOptionsErr)
	}
	return _cTTextTabGetOptions(tab), nil
}

// CTTextTabGetOptions returns the dictionary of attributes associated with the tab.
//
// See: https://developer.apple.com/documentation/CoreText/CTTextTabGetOptions(_:)
func CTTextTabGetOptions(tab CTTextTabRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCTTextTabGetOptions(tab)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTextTabGetTypeID func() uint
var _cTTextTabGetTypeIDErr error

func tryCTTextTabGetTypeID() (uint, error) {
	if _cTTextTabGetTypeID == nil {
		return 0, symbolCallError("CTTextTabGetTypeID", "10.5", _cTTextTabGetTypeIDErr)
	}
	return _cTTextTabGetTypeID(), nil
}

// CTTextTabGetTypeID returns the Core Foundation type identifier of the text tab object.
//
// See: https://developer.apple.com/documentation/CoreText/CTTextTabGetTypeID()
func CTTextTabGetTypeID() uint {
	result, callErr := tryCTTextTabGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterCreateLine func(typesetter CTTypesetterRef, stringRange corefoundation.CFRange) CTLineRef
var _cTTypesetterCreateLineErr error

func tryCTTypesetterCreateLine(typesetter CTTypesetterRef, stringRange corefoundation.CFRange) (CTLineRef, error) {
	if _cTTypesetterCreateLine == nil {
		return *new(CTLineRef), symbolCallError("CTTypesetterCreateLine", "10.5", _cTTypesetterCreateLineErr)
	}
	return _cTTypesetterCreateLine(typesetter, stringRange), nil
}

// CTTypesetterCreateLine creates an immutable line from the typesetter.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateLine(_:_:)
func CTTypesetterCreateLine(typesetter CTTypesetterRef, stringRange corefoundation.CFRange) CTLineRef {
	result, callErr := tryCTTypesetterCreateLine(typesetter, stringRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterCreateLineWithOffset func(typesetter CTTypesetterRef, stringRange corefoundation.CFRange, offset float64) CTLineRef
var _cTTypesetterCreateLineWithOffsetErr error

func tryCTTypesetterCreateLineWithOffset(typesetter CTTypesetterRef, stringRange corefoundation.CFRange, offset float64) (CTLineRef, error) {
	if _cTTypesetterCreateLineWithOffset == nil {
		return *new(CTLineRef), symbolCallError("CTTypesetterCreateLineWithOffset", "10.6", _cTTypesetterCreateLineWithOffsetErr)
	}
	return _cTTypesetterCreateLineWithOffset(typesetter, stringRange, offset), nil
}

// CTTypesetterCreateLineWithOffset creates an immutable line from the typesetter at a specified line offset.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateLineWithOffset(_:_:_:)
func CTTypesetterCreateLineWithOffset(typesetter CTTypesetterRef, stringRange corefoundation.CFRange, offset float64) CTLineRef {
	result, callErr := tryCTTypesetterCreateLineWithOffset(typesetter, stringRange, offset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterCreateWithAttributedString func(string_ corefoundation.CFAttributedStringRef) CTTypesetterRef
var _cTTypesetterCreateWithAttributedStringErr error

func tryCTTypesetterCreateWithAttributedString(string_ corefoundation.CFAttributedStringRef) (CTTypesetterRef, error) {
	if _cTTypesetterCreateWithAttributedString == nil {
		return *new(CTTypesetterRef), symbolCallError("CTTypesetterCreateWithAttributedString", "10.5", _cTTypesetterCreateWithAttributedStringErr)
	}
	return _cTTypesetterCreateWithAttributedString(string_), nil
}

// CTTypesetterCreateWithAttributedString creates an immutable typesetter object using an attributed string.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateWithAttributedString(_:)
func CTTypesetterCreateWithAttributedString(string_ corefoundation.CFAttributedStringRef) CTTypesetterRef {
	result, callErr := tryCTTypesetterCreateWithAttributedString(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterCreateWithAttributedStringAndOptions func(string_ corefoundation.CFAttributedStringRef, options corefoundation.CFDictionaryRef) CTTypesetterRef
var _cTTypesetterCreateWithAttributedStringAndOptionsErr error

func tryCTTypesetterCreateWithAttributedStringAndOptions(string_ corefoundation.CFAttributedStringRef, options corefoundation.CFDictionaryRef) (CTTypesetterRef, error) {
	if _cTTypesetterCreateWithAttributedStringAndOptions == nil {
		return *new(CTTypesetterRef), symbolCallError("CTTypesetterCreateWithAttributedStringAndOptions", "10.5", _cTTypesetterCreateWithAttributedStringAndOptionsErr)
	}
	return _cTTypesetterCreateWithAttributedStringAndOptions(string_, options), nil
}

// CTTypesetterCreateWithAttributedStringAndOptions creates an immutable typesetter object using an attributed string and a dictionary of options.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterCreateWithAttributedStringAndOptions(_:_:)
func CTTypesetterCreateWithAttributedStringAndOptions(string_ corefoundation.CFAttributedStringRef, options corefoundation.CFDictionaryRef) CTTypesetterRef {
	result, callErr := tryCTTypesetterCreateWithAttributedStringAndOptions(string_, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterGetTypeID func() uint
var _cTTypesetterGetTypeIDErr error

func tryCTTypesetterGetTypeID() (uint, error) {
	if _cTTypesetterGetTypeID == nil {
		return 0, symbolCallError("CTTypesetterGetTypeID", "10.5", _cTTypesetterGetTypeIDErr)
	}
	return _cTTypesetterGetTypeID(), nil
}

// CTTypesetterGetTypeID returns the Core Foundation type identifier of the typesetter object.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterGetTypeID()
func CTTypesetterGetTypeID() uint {
	result, callErr := tryCTTypesetterGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterSuggestClusterBreak func(typesetter CTTypesetterRef, startIndex int, width float64) int
var _cTTypesetterSuggestClusterBreakErr error

func tryCTTypesetterSuggestClusterBreak(typesetter CTTypesetterRef, startIndex int, width float64) (int, error) {
	if _cTTypesetterSuggestClusterBreak == nil {
		return 0, symbolCallError("CTTypesetterSuggestClusterBreak", "10.5", _cTTypesetterSuggestClusterBreakErr)
	}
	return _cTTypesetterSuggestClusterBreak(typesetter, startIndex, width), nil
}

// CTTypesetterSuggestClusterBreak suggests a cluster line breakpoint based on the width provided.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestClusterBreak(_:_:_:)
func CTTypesetterSuggestClusterBreak(typesetter CTTypesetterRef, startIndex int, width float64) int {
	result, callErr := tryCTTypesetterSuggestClusterBreak(typesetter, startIndex, width)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterSuggestClusterBreakWithOffset func(typesetter CTTypesetterRef, startIndex int, width float64, offset float64) int
var _cTTypesetterSuggestClusterBreakWithOffsetErr error

func tryCTTypesetterSuggestClusterBreakWithOffset(typesetter CTTypesetterRef, startIndex int, width float64, offset float64) (int, error) {
	if _cTTypesetterSuggestClusterBreakWithOffset == nil {
		return 0, symbolCallError("CTTypesetterSuggestClusterBreakWithOffset", "10.6", _cTTypesetterSuggestClusterBreakWithOffsetErr)
	}
	return _cTTypesetterSuggestClusterBreakWithOffset(typesetter, startIndex, width, offset), nil
}

// CTTypesetterSuggestClusterBreakWithOffset suggests a cluster line breakpoint based on the specified width and line offset.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestClusterBreakWithOffset(_:_:_:_:)
func CTTypesetterSuggestClusterBreakWithOffset(typesetter CTTypesetterRef, startIndex int, width float64, offset float64) int {
	result, callErr := tryCTTypesetterSuggestClusterBreakWithOffset(typesetter, startIndex, width, offset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterSuggestLineBreak func(typesetter CTTypesetterRef, startIndex int, width float64) int
var _cTTypesetterSuggestLineBreakErr error

func tryCTTypesetterSuggestLineBreak(typesetter CTTypesetterRef, startIndex int, width float64) (int, error) {
	if _cTTypesetterSuggestLineBreak == nil {
		return 0, symbolCallError("CTTypesetterSuggestLineBreak", "10.5", _cTTypesetterSuggestLineBreakErr)
	}
	return _cTTypesetterSuggestLineBreak(typesetter, startIndex, width), nil
}

// CTTypesetterSuggestLineBreak suggests a contextual line breakpoint based on the width provided.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestLineBreak(_:_:_:)
func CTTypesetterSuggestLineBreak(typesetter CTTypesetterRef, startIndex int, width float64) int {
	result, callErr := tryCTTypesetterSuggestLineBreak(typesetter, startIndex, width)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cTTypesetterSuggestLineBreakWithOffset func(typesetter CTTypesetterRef, startIndex int, width float64, offset float64) int
var _cTTypesetterSuggestLineBreakWithOffsetErr error

func tryCTTypesetterSuggestLineBreakWithOffset(typesetter CTTypesetterRef, startIndex int, width float64, offset float64) (int, error) {
	if _cTTypesetterSuggestLineBreakWithOffset == nil {
		return 0, symbolCallError("CTTypesetterSuggestLineBreakWithOffset", "10.6", _cTTypesetterSuggestLineBreakWithOffsetErr)
	}
	return _cTTypesetterSuggestLineBreakWithOffset(typesetter, startIndex, width, offset), nil
}

// CTTypesetterSuggestLineBreakWithOffset suggests a contextual line breakpoint based on the width provided and the specified offset.
//
// See: https://developer.apple.com/documentation/CoreText/CTTypesetterSuggestLineBreakWithOffset(_:_:_:_:)
func CTTypesetterSuggestLineBreakWithOffset(typesetter CTTypesetterRef, startIndex int, width float64, offset float64) int {
	result, callErr := tryCTTypesetterSuggestLineBreakWithOffset(typesetter, startIndex, width, offset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cTFontCollectionCopyExclusionDescriptors, &_cTFontCollectionCopyExclusionDescriptorsErr, frameworkHandle, "CTFontCollectionCopyExclusionDescriptors", "10.7")
	registerFunc(&_cTFontCollectionCopyFontAttribute, &_cTFontCollectionCopyFontAttributeErr, frameworkHandle, "CTFontCollectionCopyFontAttribute", "10.7")
	registerFunc(&_cTFontCollectionCopyFontAttributes, &_cTFontCollectionCopyFontAttributesErr, frameworkHandle, "CTFontCollectionCopyFontAttributes", "10.7")
	registerFunc(&_cTFontCollectionCopyQueryDescriptors, &_cTFontCollectionCopyQueryDescriptorsErr, frameworkHandle, "CTFontCollectionCopyQueryDescriptors", "10.7")
	registerFunc(&_cTFontCollectionCreateCopyWithFontDescriptors, &_cTFontCollectionCreateCopyWithFontDescriptorsErr, frameworkHandle, "CTFontCollectionCreateCopyWithFontDescriptors", "10.5")
	registerFunc(&_cTFontCollectionCreateFromAvailableFonts, &_cTFontCollectionCreateFromAvailableFontsErr, frameworkHandle, "CTFontCollectionCreateFromAvailableFonts", "10.5")
	registerFunc(&_cTFontCollectionCreateMatchingFontDescriptors, &_cTFontCollectionCreateMatchingFontDescriptorsErr, frameworkHandle, "CTFontCollectionCreateMatchingFontDescriptors", "10.5")
	registerFunc(&_cTFontCollectionCreateMatchingFontDescriptorsForFamily, &_cTFontCollectionCreateMatchingFontDescriptorsForFamilyErr, frameworkHandle, "CTFontCollectionCreateMatchingFontDescriptorsForFamily", "10.7")
	registerFunc(&_cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback, &_cTFontCollectionCreateMatchingFontDescriptorsSortedWithCallbackErr, frameworkHandle, "CTFontCollectionCreateMatchingFontDescriptorsSortedWithCallback", "10.5")
	registerFunc(&_cTFontCollectionCreateMatchingFontDescriptorsWithOptions, &_cTFontCollectionCreateMatchingFontDescriptorsWithOptionsErr, frameworkHandle, "CTFontCollectionCreateMatchingFontDescriptorsWithOptions", "10.7")
	registerFunc(&_cTFontCollectionCreateMutableCopy, &_cTFontCollectionCreateMutableCopyErr, frameworkHandle, "CTFontCollectionCreateMutableCopy", "10.7")
	registerFunc(&_cTFontCollectionCreateWithFontDescriptors, &_cTFontCollectionCreateWithFontDescriptorsErr, frameworkHandle, "CTFontCollectionCreateWithFontDescriptors", "10.5")
	registerFunc(&_cTFontCollectionGetTypeID, &_cTFontCollectionGetTypeIDErr, frameworkHandle, "CTFontCollectionGetTypeID", "10.5")
	registerFunc(&_cTFontCollectionSetExclusionDescriptors, &_cTFontCollectionSetExclusionDescriptorsErr, frameworkHandle, "CTFontCollectionSetExclusionDescriptors", "10.7")
	registerFunc(&_cTFontCollectionSetQueryDescriptors, &_cTFontCollectionSetQueryDescriptorsErr, frameworkHandle, "CTFontCollectionSetQueryDescriptors", "10.7")
	registerFunc(&_cTFontCopyAttribute, &_cTFontCopyAttributeErr, frameworkHandle, "CTFontCopyAttribute", "10.5")
	registerFunc(&_cTFontCopyAvailableTables, &_cTFontCopyAvailableTablesErr, frameworkHandle, "CTFontCopyAvailableTables", "10.5")
	registerFunc(&_cTFontCopyCharacterSet, &_cTFontCopyCharacterSetErr, frameworkHandle, "CTFontCopyCharacterSet", "10.5")
	registerFunc(&_cTFontCopyDefaultCascadeListForLanguages, &_cTFontCopyDefaultCascadeListForLanguagesErr, frameworkHandle, "CTFontCopyDefaultCascadeListForLanguages", "10.8")
	registerFunc(&_cTFontCopyDisplayName, &_cTFontCopyDisplayNameErr, frameworkHandle, "CTFontCopyDisplayName", "10.5")
	registerFunc(&_cTFontCopyFamilyName, &_cTFontCopyFamilyNameErr, frameworkHandle, "CTFontCopyFamilyName", "10.5")
	registerFunc(&_cTFontCopyFeatureSettings, &_cTFontCopyFeatureSettingsErr, frameworkHandle, "CTFontCopyFeatureSettings", "10.5")
	registerFunc(&_cTFontCopyFeatures, &_cTFontCopyFeaturesErr, frameworkHandle, "CTFontCopyFeatures", "10.5")
	registerFunc(&_cTFontCopyFontDescriptor, &_cTFontCopyFontDescriptorErr, frameworkHandle, "CTFontCopyFontDescriptor", "10.5")
	registerFunc(&_cTFontCopyFullName, &_cTFontCopyFullNameErr, frameworkHandle, "CTFontCopyFullName", "10.5")
	registerFunc(&_cTFontCopyGraphicsFont, &_cTFontCopyGraphicsFontErr, frameworkHandle, "CTFontCopyGraphicsFont", "10.5")
	registerFunc(&_cTFontCopyLocalizedName, &_cTFontCopyLocalizedNameErr, frameworkHandle, "CTFontCopyLocalizedName", "10.5")
	registerFunc(&_cTFontCopyName, &_cTFontCopyNameErr, frameworkHandle, "CTFontCopyName", "10.5")
	registerFunc(&_cTFontCopyNameForGlyph, &_cTFontCopyNameForGlyphErr, frameworkHandle, "CTFontCopyNameForGlyph", "10.8")
	registerFunc(&_cTFontCopyPostScriptName, &_cTFontCopyPostScriptNameErr, frameworkHandle, "CTFontCopyPostScriptName", "10.5")
	registerFunc(&_cTFontCopySupportedLanguages, &_cTFontCopySupportedLanguagesErr, frameworkHandle, "CTFontCopySupportedLanguages", "10.5")
	registerFunc(&_cTFontCopyTable, &_cTFontCopyTableErr, frameworkHandle, "CTFontCopyTable", "10.5")
	registerFunc(&_cTFontCopyTraits, &_cTFontCopyTraitsErr, frameworkHandle, "CTFontCopyTraits", "10.5")
	registerFunc(&_cTFontCopyVariation, &_cTFontCopyVariationErr, frameworkHandle, "CTFontCopyVariation", "10.5")
	registerFunc(&_cTFontCopyVariationAxes, &_cTFontCopyVariationAxesErr, frameworkHandle, "CTFontCopyVariationAxes", "10.5")
	registerFunc(&_cTFontCreateCopyWithAttributes, &_cTFontCreateCopyWithAttributesErr, frameworkHandle, "CTFontCreateCopyWithAttributes", "10.5")
	registerFunc(&_cTFontCreateCopyWithFamily, &_cTFontCreateCopyWithFamilyErr, frameworkHandle, "CTFontCreateCopyWithFamily", "10.5")
	registerFunc(&_cTFontCreateCopyWithSymbolicTraits, &_cTFontCreateCopyWithSymbolicTraitsErr, frameworkHandle, "CTFontCreateCopyWithSymbolicTraits", "10.5")
	registerFunc(&_cTFontCreateForString, &_cTFontCreateForStringErr, frameworkHandle, "CTFontCreateForString", "10.5")
	registerFunc(&_cTFontCreateForStringWithLanguage, &_cTFontCreateForStringWithLanguageErr, frameworkHandle, "CTFontCreateForStringWithLanguage", "10.9")
	registerFunc(&_cTFontCreatePathForGlyph, &_cTFontCreatePathForGlyphErr, frameworkHandle, "CTFontCreatePathForGlyph", "10.5")
	registerFunc(&_cTFontCreateUIFontForLanguage, &_cTFontCreateUIFontForLanguageErr, frameworkHandle, "CTFontCreateUIFontForLanguage", "10.5")
	registerFunc(&_cTFontCreateWithFontDescriptor, &_cTFontCreateWithFontDescriptorErr, frameworkHandle, "CTFontCreateWithFontDescriptor", "10.5")
	registerFunc(&_cTFontCreateWithFontDescriptorAndOptions, &_cTFontCreateWithFontDescriptorAndOptionsErr, frameworkHandle, "CTFontCreateWithFontDescriptorAndOptions", "10.6")
	registerFunc(&_cTFontCreateWithGraphicsFont, &_cTFontCreateWithGraphicsFontErr, frameworkHandle, "CTFontCreateWithGraphicsFont", "10.5")
	registerFunc(&_cTFontCreateWithName, &_cTFontCreateWithNameErr, frameworkHandle, "CTFontCreateWithName", "10.5")
	registerFunc(&_cTFontCreateWithNameAndOptions, &_cTFontCreateWithNameAndOptionsErr, frameworkHandle, "CTFontCreateWithNameAndOptions", "10.6")
	registerFunc(&_cTFontDescriptorCopyAttribute, &_cTFontDescriptorCopyAttributeErr, frameworkHandle, "CTFontDescriptorCopyAttribute", "10.5")
	registerFunc(&_cTFontDescriptorCopyAttributes, &_cTFontDescriptorCopyAttributesErr, frameworkHandle, "CTFontDescriptorCopyAttributes", "10.5")
	registerFunc(&_cTFontDescriptorCopyLocalizedAttribute, &_cTFontDescriptorCopyLocalizedAttributeErr, frameworkHandle, "CTFontDescriptorCopyLocalizedAttribute", "10.5")
	registerFunc(&_cTFontDescriptorCreateCopyWithAttributes, &_cTFontDescriptorCreateCopyWithAttributesErr, frameworkHandle, "CTFontDescriptorCreateCopyWithAttributes", "10.5")
	registerFunc(&_cTFontDescriptorCreateCopyWithFamily, &_cTFontDescriptorCreateCopyWithFamilyErr, frameworkHandle, "CTFontDescriptorCreateCopyWithFamily", "10.9")
	registerFunc(&_cTFontDescriptorCreateCopyWithFeature, &_cTFontDescriptorCreateCopyWithFeatureErr, frameworkHandle, "CTFontDescriptorCreateCopyWithFeature", "10.5")
	registerFunc(&_cTFontDescriptorCreateCopyWithSymbolicTraits, &_cTFontDescriptorCreateCopyWithSymbolicTraitsErr, frameworkHandle, "CTFontDescriptorCreateCopyWithSymbolicTraits", "10.9")
	registerFunc(&_cTFontDescriptorCreateCopyWithVariation, &_cTFontDescriptorCreateCopyWithVariationErr, frameworkHandle, "CTFontDescriptorCreateCopyWithVariation", "10.5")
	registerFunc(&_cTFontDescriptorCreateMatchingFontDescriptor, &_cTFontDescriptorCreateMatchingFontDescriptorErr, frameworkHandle, "CTFontDescriptorCreateMatchingFontDescriptor", "10.5")
	registerFunc(&_cTFontDescriptorCreateMatchingFontDescriptors, &_cTFontDescriptorCreateMatchingFontDescriptorsErr, frameworkHandle, "CTFontDescriptorCreateMatchingFontDescriptors", "10.5")
	registerFunc(&_cTFontDescriptorCreateWithAttributes, &_cTFontDescriptorCreateWithAttributesErr, frameworkHandle, "CTFontDescriptorCreateWithAttributes", "10.5")
	registerFunc(&_cTFontDescriptorCreateWithNameAndSize, &_cTFontDescriptorCreateWithNameAndSizeErr, frameworkHandle, "CTFontDescriptorCreateWithNameAndSize", "10.5")
	registerFunc(&_cTFontDescriptorGetTypeID, &_cTFontDescriptorGetTypeIDErr, frameworkHandle, "CTFontDescriptorGetTypeID", "10.5")
	registerFunc(&_cTFontDescriptorMatchFontDescriptorsWithProgressHandler, &_cTFontDescriptorMatchFontDescriptorsWithProgressHandlerErr, frameworkHandle, "CTFontDescriptorMatchFontDescriptorsWithProgressHandler", "10.9")
	registerFunc(&_cTFontDrawGlyphs, &_cTFontDrawGlyphsErr, frameworkHandle, "CTFontDrawGlyphs", "10.7")
	registerFunc(&_cTFontDrawImageFromAdaptiveImageProviderAtPoint, &_cTFontDrawImageFromAdaptiveImageProviderAtPointErr, frameworkHandle, "CTFontDrawImageFromAdaptiveImageProviderAtPoint", "15.0")
	registerFunc(&_cTFontGetAdvancesForGlyphs, &_cTFontGetAdvancesForGlyphsErr, frameworkHandle, "CTFontGetAdvancesForGlyphs", "10.5")
	registerFunc(&_cTFontGetAscent, &_cTFontGetAscentErr, frameworkHandle, "CTFontGetAscent", "10.5")
	registerFunc(&_cTFontGetBoundingBox, &_cTFontGetBoundingBoxErr, frameworkHandle, "CTFontGetBoundingBox", "10.5")
	registerFunc(&_cTFontGetBoundingRectsForGlyphs, &_cTFontGetBoundingRectsForGlyphsErr, frameworkHandle, "CTFontGetBoundingRectsForGlyphs", "10.5")
	registerFunc(&_cTFontGetCapHeight, &_cTFontGetCapHeightErr, frameworkHandle, "CTFontGetCapHeight", "10.5")
	registerFunc(&_cTFontGetDescent, &_cTFontGetDescentErr, frameworkHandle, "CTFontGetDescent", "10.5")
	registerFunc(&_cTFontGetGlyphCount, &_cTFontGetGlyphCountErr, frameworkHandle, "CTFontGetGlyphCount", "10.5")
	registerFunc(&_cTFontGetGlyphWithName, &_cTFontGetGlyphWithNameErr, frameworkHandle, "CTFontGetGlyphWithName", "10.5")
	registerFunc(&_cTFontGetGlyphsForCharacters, &_cTFontGetGlyphsForCharactersErr, frameworkHandle, "CTFontGetGlyphsForCharacters", "10.5")
	registerFunc(&_cTFontGetLeading, &_cTFontGetLeadingErr, frameworkHandle, "CTFontGetLeading", "10.5")
	registerFunc(&_cTFontGetLigatureCaretPositions, &_cTFontGetLigatureCaretPositionsErr, frameworkHandle, "CTFontGetLigatureCaretPositions", "10.5")
	registerFunc(&_cTFontGetMatrix, &_cTFontGetMatrixErr, frameworkHandle, "CTFontGetMatrix", "10.5")
	registerFunc(&_cTFontGetOpticalBoundsForGlyphs, &_cTFontGetOpticalBoundsForGlyphsErr, frameworkHandle, "CTFontGetOpticalBoundsForGlyphs", "10.8")
	registerFunc(&_cTFontGetSize, &_cTFontGetSizeErr, frameworkHandle, "CTFontGetSize", "10.5")
	registerFunc(&_cTFontGetSlantAngle, &_cTFontGetSlantAngleErr, frameworkHandle, "CTFontGetSlantAngle", "10.5")
	registerFunc(&_cTFontGetStringEncoding, &_cTFontGetStringEncodingErr, frameworkHandle, "CTFontGetStringEncoding", "10.5")
	registerFunc(&_cTFontGetSymbolicTraits, &_cTFontGetSymbolicTraitsErr, frameworkHandle, "CTFontGetSymbolicTraits", "10.5")
	registerFunc(&_cTFontGetTypeID, &_cTFontGetTypeIDErr, frameworkHandle, "CTFontGetTypeID", "10.5")
	registerFunc(&_cTFontGetTypographicBoundsForAdaptiveImageProvider, &_cTFontGetTypographicBoundsForAdaptiveImageProviderErr, frameworkHandle, "CTFontGetTypographicBoundsForAdaptiveImageProvider", "15.0")
	registerFunc(&_cTFontGetUIFontType, &_cTFontGetUIFontTypeErr, frameworkHandle, "CTFontGetUIFontType", "10.15")
	registerFunc(&_cTFontGetUnderlinePosition, &_cTFontGetUnderlinePositionErr, frameworkHandle, "CTFontGetUnderlinePosition", "10.5")
	registerFunc(&_cTFontGetUnderlineThickness, &_cTFontGetUnderlineThicknessErr, frameworkHandle, "CTFontGetUnderlineThickness", "10.5")
	registerFunc(&_cTFontGetUnitsPerEm, &_cTFontGetUnitsPerEmErr, frameworkHandle, "CTFontGetUnitsPerEm", "10.5")
	registerFunc(&_cTFontGetVerticalTranslationsForGlyphs, &_cTFontGetVerticalTranslationsForGlyphsErr, frameworkHandle, "CTFontGetVerticalTranslationsForGlyphs", "10.5")
	registerFunc(&_cTFontGetXHeight, &_cTFontGetXHeightErr, frameworkHandle, "CTFontGetXHeight", "10.5")
	registerFunc(&_cTFontHasTable, &_cTFontHasTableErr, frameworkHandle, "CTFontHasTable", "10.15")
	registerFunc(&_cTFontManagerCompareFontFamilyNames, &_cTFontManagerCompareFontFamilyNamesErr, frameworkHandle, "CTFontManagerCompareFontFamilyNames", "10.6")
	registerFunc(&_cTFontManagerCopyAvailableFontFamilyNames, &_cTFontManagerCopyAvailableFontFamilyNamesErr, frameworkHandle, "CTFontManagerCopyAvailableFontFamilyNames", "10.6")
	registerFunc(&_cTFontManagerCopyAvailableFontURLs, &_cTFontManagerCopyAvailableFontURLsErr, frameworkHandle, "CTFontManagerCopyAvailableFontURLs", "10.6")
	registerFunc(&_cTFontManagerCopyAvailablePostScriptNames, &_cTFontManagerCopyAvailablePostScriptNamesErr, frameworkHandle, "CTFontManagerCopyAvailablePostScriptNames", "10.6")
	registerFunc(&_cTFontManagerCreateFontDescriptorFromData, &_cTFontManagerCreateFontDescriptorFromDataErr, frameworkHandle, "CTFontManagerCreateFontDescriptorFromData", "10.7")
	registerFunc(&_cTFontManagerCreateFontDescriptorsFromData, &_cTFontManagerCreateFontDescriptorsFromDataErr, frameworkHandle, "CTFontManagerCreateFontDescriptorsFromData", "10.13")
	registerFunc(&_cTFontManagerCreateFontDescriptorsFromURL, &_cTFontManagerCreateFontDescriptorsFromURLErr, frameworkHandle, "CTFontManagerCreateFontDescriptorsFromURL", "10.6")
	registerFunc(&_cTFontManagerEnableFontDescriptors, &_cTFontManagerEnableFontDescriptorsErr, frameworkHandle, "CTFontManagerEnableFontDescriptors", "10.6")
	registerFunc(&_cTFontManagerGetAutoActivationSetting, &_cTFontManagerGetAutoActivationSettingErr, frameworkHandle, "CTFontManagerGetAutoActivationSetting", "10.6")
	registerFunc(&_cTFontManagerGetScopeForURL, &_cTFontManagerGetScopeForURLErr, frameworkHandle, "CTFontManagerGetScopeForURL", "10.6")
	registerFunc(&_cTFontManagerIsSupportedFont, &_cTFontManagerIsSupportedFontErr, frameworkHandle, "CTFontManagerIsSupportedFont", "10.6")
	registerFunc(&_cTFontManagerRegisterFontDescriptors, &_cTFontManagerRegisterFontDescriptorsErr, frameworkHandle, "CTFontManagerRegisterFontDescriptors", "10.15")
	registerFunc(&_cTFontManagerRegisterFontURLs, &_cTFontManagerRegisterFontURLsErr, frameworkHandle, "CTFontManagerRegisterFontURLs", "10.15")
	registerFunc(&_cTFontManagerRegisterFontsForURL, &_cTFontManagerRegisterFontsForURLErr, frameworkHandle, "CTFontManagerRegisterFontsForURL", "10.6")
	registerFunc(&_cTFontManagerRegisterFontsForURLs, &_cTFontManagerRegisterFontsForURLsErr, frameworkHandle, "CTFontManagerRegisterFontsForURLs", "10.6")
	registerFunc(&_cTFontManagerRegisterGraphicsFont, &_cTFontManagerRegisterGraphicsFontErr, frameworkHandle, "CTFontManagerRegisterGraphicsFont", "10.8")
	registerFunc(&_cTFontManagerSetAutoActivationSetting, &_cTFontManagerSetAutoActivationSettingErr, frameworkHandle, "CTFontManagerSetAutoActivationSetting", "10.6")
	registerFunc(&_cTFontManagerUnregisterFontDescriptors, &_cTFontManagerUnregisterFontDescriptorsErr, frameworkHandle, "CTFontManagerUnregisterFontDescriptors", "10.15")
	registerFunc(&_cTFontManagerUnregisterFontURLs, &_cTFontManagerUnregisterFontURLsErr, frameworkHandle, "CTFontManagerUnregisterFontURLs", "10.15")
	registerFunc(&_cTFontManagerUnregisterFontsForURL, &_cTFontManagerUnregisterFontsForURLErr, frameworkHandle, "CTFontManagerUnregisterFontsForURL", "10.6")
	registerFunc(&_cTFontManagerUnregisterFontsForURLs, &_cTFontManagerUnregisterFontsForURLsErr, frameworkHandle, "CTFontManagerUnregisterFontsForURLs", "10.6")
	registerFunc(&_cTFontManagerUnregisterGraphicsFont, &_cTFontManagerUnregisterGraphicsFontErr, frameworkHandle, "CTFontManagerUnregisterGraphicsFont", "10.8")
	registerFunc(&_cTFrameDraw, &_cTFrameDrawErr, frameworkHandle, "CTFrameDraw", "10.5")
	registerFunc(&_cTFrameGetFrameAttributes, &_cTFrameGetFrameAttributesErr, frameworkHandle, "CTFrameGetFrameAttributes", "10.5")
	registerFunc(&_cTFrameGetLineOrigins, &_cTFrameGetLineOriginsErr, frameworkHandle, "CTFrameGetLineOrigins", "10.5")
	registerFunc(&_cTFrameGetLines, &_cTFrameGetLinesErr, frameworkHandle, "CTFrameGetLines", "10.5")
	registerFunc(&_cTFrameGetPath, &_cTFrameGetPathErr, frameworkHandle, "CTFrameGetPath", "10.5")
	registerFunc(&_cTFrameGetStringRange, &_cTFrameGetStringRangeErr, frameworkHandle, "CTFrameGetStringRange", "10.5")
	registerFunc(&_cTFrameGetTypeID, &_cTFrameGetTypeIDErr, frameworkHandle, "CTFrameGetTypeID", "10.5")
	registerFunc(&_cTFrameGetVisibleStringRange, &_cTFrameGetVisibleStringRangeErr, frameworkHandle, "CTFrameGetVisibleStringRange", "10.5")
	registerFunc(&_cTFramesetterCreateFrame, &_cTFramesetterCreateFrameErr, frameworkHandle, "CTFramesetterCreateFrame", "10.5")
	registerFunc(&_cTFramesetterCreateWithAttributedString, &_cTFramesetterCreateWithAttributedStringErr, frameworkHandle, "CTFramesetterCreateWithAttributedString", "10.5")
	registerFunc(&_cTFramesetterCreateWithTypesetter, &_cTFramesetterCreateWithTypesetterErr, frameworkHandle, "CTFramesetterCreateWithTypesetter", "10.14")
	registerFunc(&_cTFramesetterGetTypeID, &_cTFramesetterGetTypeIDErr, frameworkHandle, "CTFramesetterGetTypeID", "10.5")
	registerFunc(&_cTFramesetterGetTypesetter, &_cTFramesetterGetTypesetterErr, frameworkHandle, "CTFramesetterGetTypesetter", "10.5")
	registerFunc(&_cTFramesetterSuggestFrameSizeWithConstraints, &_cTFramesetterSuggestFrameSizeWithConstraintsErr, frameworkHandle, "CTFramesetterSuggestFrameSizeWithConstraints", "10.5")
	registerFunc(&_cTGlyphInfoCreateWithCharacterIdentifier, &_cTGlyphInfoCreateWithCharacterIdentifierErr, frameworkHandle, "CTGlyphInfoCreateWithCharacterIdentifier", "10.5")
	registerFunc(&_cTGlyphInfoCreateWithGlyph, &_cTGlyphInfoCreateWithGlyphErr, frameworkHandle, "CTGlyphInfoCreateWithGlyph", "10.5")
	registerFunc(&_cTGlyphInfoCreateWithGlyphName, &_cTGlyphInfoCreateWithGlyphNameErr, frameworkHandle, "CTGlyphInfoCreateWithGlyphName", "10.5")
	registerFunc(&_cTGlyphInfoGetCharacterCollection, &_cTGlyphInfoGetCharacterCollectionErr, frameworkHandle, "CTGlyphInfoGetCharacterCollection", "10.5")
	registerFunc(&_cTGlyphInfoGetCharacterIdentifier, &_cTGlyphInfoGetCharacterIdentifierErr, frameworkHandle, "CTGlyphInfoGetCharacterIdentifier", "10.5")
	registerFunc(&_cTGlyphInfoGetGlyph, &_cTGlyphInfoGetGlyphErr, frameworkHandle, "CTGlyphInfoGetGlyph", "10.15")
	registerFunc(&_cTGlyphInfoGetGlyphName, &_cTGlyphInfoGetGlyphNameErr, frameworkHandle, "CTGlyphInfoGetGlyphName", "10.5")
	registerFunc(&_cTGlyphInfoGetTypeID, &_cTGlyphInfoGetTypeIDErr, frameworkHandle, "CTGlyphInfoGetTypeID", "10.5")
	registerFunc(&_cTLineCreateJustifiedLine, &_cTLineCreateJustifiedLineErr, frameworkHandle, "CTLineCreateJustifiedLine", "10.5")
	registerFunc(&_cTLineCreateTruncatedLine, &_cTLineCreateTruncatedLineErr, frameworkHandle, "CTLineCreateTruncatedLine", "10.5")
	registerFunc(&_cTLineCreateWithAttributedString, &_cTLineCreateWithAttributedStringErr, frameworkHandle, "CTLineCreateWithAttributedString", "10.5")
	registerFunc(&_cTLineDraw, &_cTLineDrawErr, frameworkHandle, "CTLineDraw", "10.5")
	registerFunc(&_cTLineEnumerateCaretOffsets, &_cTLineEnumerateCaretOffsetsErr, frameworkHandle, "CTLineEnumerateCaretOffsets", "10.11")
	registerFunc(&_cTLineGetBoundsWithOptions, &_cTLineGetBoundsWithOptionsErr, frameworkHandle, "CTLineGetBoundsWithOptions", "10.8")
	registerFunc(&_cTLineGetGlyphCount, &_cTLineGetGlyphCountErr, frameworkHandle, "CTLineGetGlyphCount", "10.5")
	registerFunc(&_cTLineGetGlyphRuns, &_cTLineGetGlyphRunsErr, frameworkHandle, "CTLineGetGlyphRuns", "10.5")
	registerFunc(&_cTLineGetImageBounds, &_cTLineGetImageBoundsErr, frameworkHandle, "CTLineGetImageBounds", "10.5")
	registerFunc(&_cTLineGetOffsetForStringIndex, &_cTLineGetOffsetForStringIndexErr, frameworkHandle, "CTLineGetOffsetForStringIndex", "10.5")
	registerFunc(&_cTLineGetPenOffsetForFlush, &_cTLineGetPenOffsetForFlushErr, frameworkHandle, "CTLineGetPenOffsetForFlush", "10.5")
	registerFunc(&_cTLineGetStringIndexForPosition, &_cTLineGetStringIndexForPositionErr, frameworkHandle, "CTLineGetStringIndexForPosition", "10.5")
	registerFunc(&_cTLineGetStringRange, &_cTLineGetStringRangeErr, frameworkHandle, "CTLineGetStringRange", "10.5")
	registerFunc(&_cTLineGetTrailingWhitespaceWidth, &_cTLineGetTrailingWhitespaceWidthErr, frameworkHandle, "CTLineGetTrailingWhitespaceWidth", "10.5")
	registerFunc(&_cTLineGetTypeID, &_cTLineGetTypeIDErr, frameworkHandle, "CTLineGetTypeID", "10.5")
	registerFunc(&_cTLineGetTypographicBounds, &_cTLineGetTypographicBoundsErr, frameworkHandle, "CTLineGetTypographicBounds", "10.5")
	registerFunc(&_cTParagraphStyleCreate, &_cTParagraphStyleCreateErr, frameworkHandle, "CTParagraphStyleCreate", "10.5")
	registerFunc(&_cTParagraphStyleCreateCopy, &_cTParagraphStyleCreateCopyErr, frameworkHandle, "CTParagraphStyleCreateCopy", "10.5")
	registerFunc(&_cTParagraphStyleGetTypeID, &_cTParagraphStyleGetTypeIDErr, frameworkHandle, "CTParagraphStyleGetTypeID", "10.5")
	registerFunc(&_cTParagraphStyleGetValueForSpecifier, &_cTParagraphStyleGetValueForSpecifierErr, frameworkHandle, "CTParagraphStyleGetValueForSpecifier", "10.5")
	registerFunc(&_cTRubyAnnotationCreate, &_cTRubyAnnotationCreateErr, frameworkHandle, "CTRubyAnnotationCreate", "10.10")
	registerFunc(&_cTRubyAnnotationCreateCopy, &_cTRubyAnnotationCreateCopyErr, frameworkHandle, "CTRubyAnnotationCreateCopy", "10.10")
	registerFunc(&_cTRubyAnnotationCreateWithAttributes, &_cTRubyAnnotationCreateWithAttributesErr, frameworkHandle, "CTRubyAnnotationCreateWithAttributes", "10.12")
	registerFunc(&_cTRubyAnnotationGetAlignment, &_cTRubyAnnotationGetAlignmentErr, frameworkHandle, "CTRubyAnnotationGetAlignment", "10.10")
	registerFunc(&_cTRubyAnnotationGetOverhang, &_cTRubyAnnotationGetOverhangErr, frameworkHandle, "CTRubyAnnotationGetOverhang", "10.10")
	registerFunc(&_cTRubyAnnotationGetSizeFactor, &_cTRubyAnnotationGetSizeFactorErr, frameworkHandle, "CTRubyAnnotationGetSizeFactor", "10.10")
	registerFunc(&_cTRubyAnnotationGetTextForPosition, &_cTRubyAnnotationGetTextForPositionErr, frameworkHandle, "CTRubyAnnotationGetTextForPosition", "10.10")
	registerFunc(&_cTRubyAnnotationGetTypeID, &_cTRubyAnnotationGetTypeIDErr, frameworkHandle, "CTRubyAnnotationGetTypeID", "10.10")
	registerFunc(&_cTRunDelegateCreate, &_cTRunDelegateCreateErr, frameworkHandle, "CTRunDelegateCreate", "10.5")
	registerFunc(&_cTRunDelegateGetRefCon, &_cTRunDelegateGetRefConErr, frameworkHandle, "CTRunDelegateGetRefCon", "10.5")
	registerFunc(&_cTRunDelegateGetTypeID, &_cTRunDelegateGetTypeIDErr, frameworkHandle, "CTRunDelegateGetTypeID", "10.5")
	registerFunc(&_cTRunDraw, &_cTRunDrawErr, frameworkHandle, "CTRunDraw", "10.5")
	registerFunc(&_cTRunGetAdvances, &_cTRunGetAdvancesErr, frameworkHandle, "CTRunGetAdvances", "10.5")
	registerFunc(&_cTRunGetAdvancesPtr, &_cTRunGetAdvancesPtrErr, frameworkHandle, "CTRunGetAdvancesPtr", "10.5")
	registerFunc(&_cTRunGetAttributes, &_cTRunGetAttributesErr, frameworkHandle, "CTRunGetAttributes", "10.5")
	registerFunc(&_cTRunGetBaseAdvancesAndOrigins, &_cTRunGetBaseAdvancesAndOriginsErr, frameworkHandle, "CTRunGetBaseAdvancesAndOrigins", "10.11")
	registerFunc(&_cTRunGetGlyphCount, &_cTRunGetGlyphCountErr, frameworkHandle, "CTRunGetGlyphCount", "10.5")
	registerFunc(&_cTRunGetGlyphs, &_cTRunGetGlyphsErr, frameworkHandle, "CTRunGetGlyphs", "10.5")
	registerFunc(&_cTRunGetGlyphsPtr, &_cTRunGetGlyphsPtrErr, frameworkHandle, "CTRunGetGlyphsPtr", "10.5")
	registerFunc(&_cTRunGetImageBounds, &_cTRunGetImageBoundsErr, frameworkHandle, "CTRunGetImageBounds", "10.5")
	registerFunc(&_cTRunGetPositions, &_cTRunGetPositionsErr, frameworkHandle, "CTRunGetPositions", "10.5")
	registerFunc(&_cTRunGetPositionsPtr, &_cTRunGetPositionsPtrErr, frameworkHandle, "CTRunGetPositionsPtr", "10.5")
	registerFunc(&_cTRunGetStatus, &_cTRunGetStatusErr, frameworkHandle, "CTRunGetStatus", "10.5")
	registerFunc(&_cTRunGetStringIndices, &_cTRunGetStringIndicesErr, frameworkHandle, "CTRunGetStringIndices", "10.5")
	registerFunc(&_cTRunGetStringIndicesPtr, &_cTRunGetStringIndicesPtrErr, frameworkHandle, "CTRunGetStringIndicesPtr", "10.5")
	registerFunc(&_cTRunGetStringRange, &_cTRunGetStringRangeErr, frameworkHandle, "CTRunGetStringRange", "10.5")
	registerFunc(&_cTRunGetTextMatrix, &_cTRunGetTextMatrixErr, frameworkHandle, "CTRunGetTextMatrix", "10.5")
	registerFunc(&_cTRunGetTypeID, &_cTRunGetTypeIDErr, frameworkHandle, "CTRunGetTypeID", "10.5")
	registerFunc(&_cTRunGetTypographicBounds, &_cTRunGetTypographicBoundsErr, frameworkHandle, "CTRunGetTypographicBounds", "10.5")
	registerFunc(&_cTTextTabCreate, &_cTTextTabCreateErr, frameworkHandle, "CTTextTabCreate", "10.5")
	registerFunc(&_cTTextTabGetAlignment, &_cTTextTabGetAlignmentErr, frameworkHandle, "CTTextTabGetAlignment", "10.5")
	registerFunc(&_cTTextTabGetLocation, &_cTTextTabGetLocationErr, frameworkHandle, "CTTextTabGetLocation", "10.5")
	registerFunc(&_cTTextTabGetOptions, &_cTTextTabGetOptionsErr, frameworkHandle, "CTTextTabGetOptions", "10.5")
	registerFunc(&_cTTextTabGetTypeID, &_cTTextTabGetTypeIDErr, frameworkHandle, "CTTextTabGetTypeID", "10.5")
	registerFunc(&_cTTypesetterCreateLine, &_cTTypesetterCreateLineErr, frameworkHandle, "CTTypesetterCreateLine", "10.5")
	registerFunc(&_cTTypesetterCreateLineWithOffset, &_cTTypesetterCreateLineWithOffsetErr, frameworkHandle, "CTTypesetterCreateLineWithOffset", "10.6")
	registerFunc(&_cTTypesetterCreateWithAttributedString, &_cTTypesetterCreateWithAttributedStringErr, frameworkHandle, "CTTypesetterCreateWithAttributedString", "10.5")
	registerFunc(&_cTTypesetterCreateWithAttributedStringAndOptions, &_cTTypesetterCreateWithAttributedStringAndOptionsErr, frameworkHandle, "CTTypesetterCreateWithAttributedStringAndOptions", "10.5")
	registerFunc(&_cTTypesetterGetTypeID, &_cTTypesetterGetTypeIDErr, frameworkHandle, "CTTypesetterGetTypeID", "10.5")
	registerFunc(&_cTTypesetterSuggestClusterBreak, &_cTTypesetterSuggestClusterBreakErr, frameworkHandle, "CTTypesetterSuggestClusterBreak", "10.5")
	registerFunc(&_cTTypesetterSuggestClusterBreakWithOffset, &_cTTypesetterSuggestClusterBreakWithOffsetErr, frameworkHandle, "CTTypesetterSuggestClusterBreakWithOffset", "10.6")
	registerFunc(&_cTTypesetterSuggestLineBreak, &_cTTypesetterSuggestLineBreakErr, frameworkHandle, "CTTypesetterSuggestLineBreak", "10.5")
	registerFunc(&_cTTypesetterSuggestLineBreakWithOffset, &_cTTypesetterSuggestLineBreakWithOffsetErr, frameworkHandle, "CTTypesetterSuggestLineBreakWithOffset", "10.6")
}
