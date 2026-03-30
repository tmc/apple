// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVContentKeySpecifier] class.
var (
	_AVContentKeySpecifierClass     AVContentKeySpecifierClass
	_AVContentKeySpecifierClassOnce sync.Once
)

func getAVContentKeySpecifierClass() AVContentKeySpecifierClass {
	_AVContentKeySpecifierClassOnce.Do(func() {
		_AVContentKeySpecifierClass = AVContentKeySpecifierClass{class: objc.GetClass("AVContentKeySpecifier")}
	})
	return _AVContentKeySpecifierClass
}

// GetAVContentKeySpecifierClass returns the class object for AVContentKeySpecifier.
func GetAVContentKeySpecifierClass() AVContentKeySpecifierClass {
	return getAVContentKeySpecifierClass()
}

type AVContentKeySpecifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVContentKeySpecifierClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVContentKeySpecifierClass) Alloc() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that uniquely identifies a content key.
//
// # Creating a specifier
//
//   - [AVContentKeySpecifier.InitForKeySystemIdentifierOptions]: Creates a content key specifier.
//
// # Inspecting a specifier
//
//   - [AVContentKeySpecifier.Identifier]: The container and protocol-specific key identifier.
//   - [AVContentKeySpecifier.KeySystem]: The key system that generates content keys.
//   - [AVContentKeySpecifier.Options]: A dictionary of options with which you initialized the specifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier
type AVContentKeySpecifier struct {
	objectivec.Object
}

// AVContentKeySpecifierFromID constructs a [AVContentKeySpecifier] from an objc.ID.
//
// An object that uniquely identifies a content key.
func AVContentKeySpecifierFromID(id objc.ID) AVContentKeySpecifier {
	return AVContentKeySpecifier{objectivec.Object{ID: id}}
}

// NOTE: AVContentKeySpecifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVContentKeySpecifier] class.
//
// # Creating a specifier
//
//   - [IAVContentKeySpecifier.InitForKeySystemIdentifierOptions]: Creates a content key specifier.
//
// # Inspecting a specifier
//
//   - [IAVContentKeySpecifier.Identifier]: The container and protocol-specific key identifier.
//   - [IAVContentKeySpecifier.KeySystem]: The key system that generates content keys.
//   - [IAVContentKeySpecifier.Options]: A dictionary of options with which you initialized the specifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier
type IAVContentKeySpecifier interface {
	objectivec.IObject

	// Topic: Creating a specifier

	// Creates a content key specifier.
	InitForKeySystemIdentifierOptions(keySystem AVContentKeySystem, contentKeyIdentifier objectivec.IObject, options foundation.INSDictionary) AVContentKeySpecifier

	// Topic: Inspecting a specifier

	// The container and protocol-specific key identifier.
	Identifier() objectivec.IObject
	// The key system that generates content keys.
	KeySystem() AVContentKeySystem
	// A dictionary of options with which you initialized the specifier.
	Options() foundation.INSDictionary
}

// Init initializes the instance.
func (c AVContentKeySpecifier) Init() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVContentKeySpecifier) Autorelease() AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContentKeySpecifier creates a new AVContentKeySpecifier instance.
func NewAVContentKeySpecifier() AVContentKeySpecifier {
	class := getAVContentKeySpecifierClass()
	rv := objc.Send[AVContentKeySpecifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a content key specifier.
//
// keySystem: The key system to use to generate content keys.
//
// contentKeyIdentifier: The container and protocol-specific key identifier.
//
// options: Additional information necessary to obtain the key. Pass `nil` to indicate
// no additional options.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/init(forKeySystem:identifier:options:)
func NewContentKeySpecifierForKeySystemIdentifierOptions(keySystem AVContentKeySystem, contentKeyIdentifier objectivec.IObject, options foundation.INSDictionary) AVContentKeySpecifier {
	instance := getAVContentKeySpecifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForKeySystem:identifier:options:"), objc.String(string(keySystem)), contentKeyIdentifier, options)
	return AVContentKeySpecifierFromID(rv)
}

// Creates a content key specifier.
//
// keySystem: The key system to use to generate content keys.
//
// contentKeyIdentifier: The container and protocol-specific key identifier.
//
// options: Additional information necessary to obtain the key. Pass `nil` to indicate
// no additional options.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/init(forKeySystem:identifier:options:)
func (c AVContentKeySpecifier) InitForKeySystemIdentifierOptions(keySystem AVContentKeySystem, contentKeyIdentifier objectivec.IObject, options foundation.INSDictionary) AVContentKeySpecifier {
	rv := objc.Send[AVContentKeySpecifier](c.ID, objc.Sel("initForKeySystem:identifier:options:"), objc.String(string(keySystem)), contentKeyIdentifier, options)
	return rv
}

// A convenience initializer to create a content key specifier.
//
// keySystem: The key system to use to generate content keys.
//
// contentKeyIdentifier: The container and protocol-specific key identifier.
//
// options: Additional information necessary to obtain the key. Pass `nil` to indicate
// no additional options.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/contentKeySpecifierForKeySystem:identifier:options:
func (_AVContentKeySpecifierClass AVContentKeySpecifierClass) ContentKeySpecifierForKeySystemIdentifierOptions(keySystem AVContentKeySystem, contentKeyIdentifier objectivec.IObject, options foundation.INSDictionary) AVContentKeySpecifier {
	rv := objc.Send[objc.ID](objc.ID(_AVContentKeySpecifierClass.class), objc.Sel("contentKeySpecifierForKeySystem:identifier:options:"), objc.String(string(keySystem)), contentKeyIdentifier, options)
	return AVContentKeySpecifierFromID(rv)
}

// The container and protocol-specific key identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/identifier
func (c AVContentKeySpecifier) Identifier() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return objectivec.Object{ID: rv}
}

// The key system that generates content keys.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/keySystem
func (c AVContentKeySpecifier) KeySystem() AVContentKeySystem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("keySystem"))
	return AVContentKeySystem(foundation.NSStringFromID(rv).String())
}

// A dictionary of options with which you initialized the specifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier/options
func (c AVContentKeySpecifier) Options() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("options"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
