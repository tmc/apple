// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVContentKey] class.
var (
	_AVContentKeyClass     AVContentKeyClass
	_AVContentKeyClassOnce sync.Once
)

func getAVContentKeyClass() AVContentKeyClass {
	_AVContentKeyClassOnce.Do(func() {
		_AVContentKeyClass = AVContentKeyClass{class: objc.GetClass("AVContentKey")}
	})
	return _AVContentKeyClass
}

// GetAVContentKeyClass returns the class object for AVContentKey.
func GetAVContentKeyClass() AVContentKeyClass {
	return getAVContentKeyClass()
}

type AVContentKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVContentKeyClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVContentKeyClass) Alloc() AVContentKey {
	rv := objc.Send[AVContentKey](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the content key decryptor.
//
// # Accessing the specifier
//
//   - [AVContentKey.ContentKeySpecifier]: The content key’s unique identifier.
//
// # Inspecting protection status
//
//   - [AVContentKey.ExternalContentProtectionStatus]: The external protection status for the content key based on all attached displays.
//   - [AVContentKey.Revoke]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKey
type AVContentKey struct {
	objectivec.Object
}

// AVContentKeyFromID constructs a [AVContentKey] from an objc.ID.
//
// An object that represents the content key decryptor.
func AVContentKeyFromID(id objc.ID) AVContentKey {
	return AVContentKey{objectivec.Object{ID: id}}
}
// NOTE: AVContentKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVContentKey] class.
//
// # Accessing the specifier
//
//   - [IAVContentKey.ContentKeySpecifier]: The content key’s unique identifier.
//
// # Inspecting protection status
//
//   - [IAVContentKey.ExternalContentProtectionStatus]: The external protection status for the content key based on all attached displays.
//   - [IAVContentKey.Revoke]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKey
type IAVContentKey interface {
	objectivec.IObject

	// Topic: Accessing the specifier

	// The content key’s unique identifier.
	ContentKeySpecifier() IAVContentKeySpecifier

	// Topic: Inspecting protection status

	// The external protection status for the content key based on all attached displays.
	ExternalContentProtectionStatus() AVExternalContentProtectionStatus
	Revoke()
}

// Init initializes the instance.
func (c AVContentKey) Init() AVContentKey {
	rv := objc.Send[AVContentKey](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVContentKey) Autorelease() AVContentKey {
	rv := objc.Send[AVContentKey](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContentKey creates a new AVContentKey instance.
func NewAVContentKey() AVContentKey {
	class := getAVContentKeyClass()
	rv := objc.Send[AVContentKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVContentKey/revoke()
func (c AVContentKey) Revoke() {
	objc.Send[objc.ID](c.ID, objc.Sel("revoke"))
}

// The content key’s unique identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKey/contentKeySpecifier
func (c AVContentKey) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("contentKeySpecifier"))
	return AVContentKeySpecifierFromID(objc.ID(rv))
}
// The external protection status for the content key based on all attached
// displays.
//
// # Discussion
// 
// This property isn’t key-value observable. Instead, use the
// [ContentKeySessionExternalProtectionStatusDidChangeForContentKey] delegate
// method to monitor changes to this value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKey/externalContentProtectionStatus
func (c AVContentKey) ExternalContentProtectionStatus() AVExternalContentProtectionStatus {
	rv := objc.Send[AVExternalContentProtectionStatus](c.ID, objc.Sel("externalContentProtectionStatus"))
	return AVExternalContentProtectionStatus(rv)
}

