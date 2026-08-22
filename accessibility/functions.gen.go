// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
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
		return fmt.Sprintf("Accessibility: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Accessibility: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Accessibility: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Accessibility: register symbol %s: %v", name, r)
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

var _aXAnimatedImagesEnabled func() bool
var _aXAnimatedImagesEnabledErr error

func tryAXAnimatedImagesEnabled() (bool, error) {
	if _aXAnimatedImagesEnabled == nil {
		return false, symbolCallError("AXAnimatedImagesEnabled", "14.0", _aXAnimatedImagesEnabledErr)
	}
	return _aXAnimatedImagesEnabled(), nil
}

// AXAnimatedImagesEnabled.
//
// See: https://developer.apple.com/documentation/Accessibility/AXAnimatedImagesEnabled
func AXAnimatedImagesEnabled() bool {
	result, callErr := tryAXAnimatedImagesEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXAssistiveAccessEnabled func() bool
var _aXAssistiveAccessEnabledErr error

func tryAXAssistiveAccessEnabled() (bool, error) {
	if _aXAssistiveAccessEnabled == nil {
		return false, symbolCallError("AXAssistiveAccessEnabled", "15.0", _aXAssistiveAccessEnabledErr)
	}
	return _aXAssistiveAccessEnabled(), nil
}

// AXAssistiveAccessEnabled a Boolean value that indicates whether Assistive Access is running.
//
// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/isAssistiveAccessEnabled
func AXAssistiveAccessEnabled() bool {
	result, callErr := tryAXAssistiveAccessEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXNameFromColor func(color coregraphics.CGColorRef) *foundation.NSString
var _aXNameFromColorErr error

func tryAXNameFromColor(color coregraphics.CGColorRef) (*foundation.NSString, error) {
	if _aXNameFromColor == nil {
		return nil, symbolCallError("AXNameFromColor", "11.0", _aXNameFromColorErr)
	}
	return _aXNameFromColor(color), nil
}

// AXNameFromColor returns a localized description of the color to use in accessibility attributes.
//
// See: https://developer.apple.com/documentation/Accessibility/AXNameFromColor(_:)
func AXNameFromColor(color coregraphics.CGColorRef) *foundation.NSString {
	result, callErr := tryAXNameFromColor(color)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXOpenSettingsFeature func(feature AXSettingsFeature)
var _aXOpenSettingsFeatureErr error

func tryAXOpenSettingsFeature(feature AXSettingsFeature) error {
	if _aXOpenSettingsFeature == nil {
		return symbolCallError("AXOpenSettingsFeature", "15.0", _aXOpenSettingsFeatureErr)
	}
	_aXOpenSettingsFeature(feature)
	return nil
}

// AXOpenSettingsFeature.
//
// See: https://developer.apple.com/documentation/Accessibility/AXOpenSettingsFeature
func AXOpenSettingsFeature(feature AXSettingsFeature) {
	if callErr := tryAXOpenSettingsFeature(feature); callErr != nil {
		panic(callErr)
	}
}

var _aXOpenSettingsFeatureIsSupported func(feature AXSettingsFeature) bool
var _aXOpenSettingsFeatureIsSupportedErr error

func tryAXOpenSettingsFeatureIsSupported(feature AXSettingsFeature) (bool, error) {
	if _aXOpenSettingsFeatureIsSupported == nil {
		return false, symbolCallError("AXOpenSettingsFeatureIsSupported", "26.4", _aXOpenSettingsFeatureIsSupportedErr)
	}
	return _aXOpenSettingsFeatureIsSupported(feature), nil
}

// AXOpenSettingsFeatureIsSupported.
//
// See: https://developer.apple.com/documentation/Accessibility/AXOpenSettingsFeatureIsSupported
func AXOpenSettingsFeatureIsSupported(feature AXSettingsFeature) bool {
	result, callErr := tryAXOpenSettingsFeatureIsSupported(feature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXPrefersActionSliderAlternative func() bool
var _aXPrefersActionSliderAlternativeErr error

func tryAXPrefersActionSliderAlternative() (bool, error) {
	if _aXPrefersActionSliderAlternative == nil {
		return false, symbolCallError("AXPrefersActionSliderAlternative", "26.1", _aXPrefersActionSliderAlternativeErr)
	}
	return _aXPrefersActionSliderAlternative(), nil
}

// AXPrefersActionSliderAlternative.
//
// See: https://developer.apple.com/documentation/Accessibility/AXPrefersActionSliderAlternative
func AXPrefersActionSliderAlternative() bool {
	result, callErr := tryAXPrefersActionSliderAlternative()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXPrefersHorizontalTextLayout func() bool
var _aXPrefersHorizontalTextLayoutErr error

func tryAXPrefersHorizontalTextLayout() (bool, error) {
	if _aXPrefersHorizontalTextLayout == nil {
		return false, symbolCallError("AXPrefersHorizontalTextLayout", "14.0", _aXPrefersHorizontalTextLayoutErr)
	}
	return _aXPrefersHorizontalTextLayout(), nil
}

// AXPrefersHorizontalTextLayout.
//
// See: https://developer.apple.com/documentation/Accessibility/AXPrefersHorizontalTextLayout
func AXPrefersHorizontalTextLayout() bool {
	result, callErr := tryAXPrefersHorizontalTextLayout()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXPrefersNonBlinkingTextInsertionIndicator func() bool
var _aXPrefersNonBlinkingTextInsertionIndicatorErr error

func tryAXPrefersNonBlinkingTextInsertionIndicator() (bool, error) {
	if _aXPrefersNonBlinkingTextInsertionIndicator == nil {
		return false, symbolCallError("AXPrefersNonBlinkingTextInsertionIndicator", "15.0", _aXPrefersNonBlinkingTextInsertionIndicatorErr)
	}
	return _aXPrefersNonBlinkingTextInsertionIndicator(), nil
}

// AXPrefersNonBlinkingTextInsertionIndicator.
//
// See: https://developer.apple.com/documentation/Accessibility/AXPrefersNonBlinkingTextInsertionIndicator
func AXPrefersNonBlinkingTextInsertionIndicator() bool {
	result, callErr := tryAXPrefersNonBlinkingTextInsertionIndicator()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXReduceHighlightingEffectsEnabled func() bool
var _aXReduceHighlightingEffectsEnabledErr error

func tryAXReduceHighlightingEffectsEnabled() (bool, error) {
	if _aXReduceHighlightingEffectsEnabled == nil {
		return false, symbolCallError("AXReduceHighlightingEffectsEnabled", "26.4", _aXReduceHighlightingEffectsEnabledErr)
	}
	return _aXReduceHighlightingEffectsEnabled(), nil
}

// AXReduceHighlightingEffectsEnabled.
//
// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/isReduceHighlightingEffectsEnabled
func AXReduceHighlightingEffectsEnabled() bool {
	result, callErr := tryAXReduceHighlightingEffectsEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXShowBordersEnabled func() bool
var _aXShowBordersEnabledErr error

func tryAXShowBordersEnabled() (bool, error) {
	if _aXShowBordersEnabled == nil {
		return false, symbolCallError("AXShowBordersEnabled", "26.1", _aXShowBordersEnabledErr)
	}
	return _aXShowBordersEnabled(), nil
}

// AXShowBordersEnabled.
//
// See: https://developer.apple.com/documentation/Accessibility/AXShowBordersEnabled
func AXShowBordersEnabled() bool {
	result, callErr := tryAXShowBordersEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_aXAnimatedImagesEnabled, &_aXAnimatedImagesEnabledErr, frameworkHandle, "AXAnimatedImagesEnabled", "14.0")
	registerFunc(&_aXAssistiveAccessEnabled, &_aXAssistiveAccessEnabledErr, frameworkHandle, "AXAssistiveAccessEnabled", "15.0")
	registerFunc(&_aXNameFromColor, &_aXNameFromColorErr, frameworkHandle, "AXNameFromColor", "11.0")
	registerFunc(&_aXOpenSettingsFeature, &_aXOpenSettingsFeatureErr, frameworkHandle, "AXOpenSettingsFeature", "15.0")
	registerFunc(&_aXOpenSettingsFeatureIsSupported, &_aXOpenSettingsFeatureIsSupportedErr, frameworkHandle, "AXOpenSettingsFeatureIsSupported", "26.4")
	registerFunc(&_aXPrefersActionSliderAlternative, &_aXPrefersActionSliderAlternativeErr, frameworkHandle, "AXPrefersActionSliderAlternative", "26.1")
	registerFunc(&_aXPrefersHorizontalTextLayout, &_aXPrefersHorizontalTextLayoutErr, frameworkHandle, "AXPrefersHorizontalTextLayout", "14.0")
	registerFunc(&_aXPrefersNonBlinkingTextInsertionIndicator, &_aXPrefersNonBlinkingTextInsertionIndicatorErr, frameworkHandle, "AXPrefersNonBlinkingTextInsertionIndicator", "15.0")
	registerFunc(&_aXReduceHighlightingEffectsEnabled, &_aXReduceHighlightingEffectsEnabledErr, frameworkHandle, "AXReduceHighlightingEffectsEnabled", "26.4")
	registerFunc(&_aXShowBordersEnabled, &_aXShowBordersEnabledErr, frameworkHandle, "AXShowBordersEnabled", "26.1")
}
