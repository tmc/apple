// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSigningKey] class.
var (
	_SLSSigningKeyClass     SLSSigningKeyClass
	_SLSSigningKeyClassOnce sync.Once
)

func getSLSSigningKeyClass() SLSSigningKeyClass {
	_SLSSigningKeyClassOnce.Do(func() {
		_SLSSigningKeyClass = SLSSigningKeyClass{class: objc.GetClass("SLSSigningKey")}
	})
	return _SLSSigningKeyClass
}

// GetSLSSigningKeyClass returns the class object for SLSSigningKey.
func GetSLSSigningKeyClass() SLSSigningKeyClass {
	return getSLSSigningKeyClass()
}

type SLSSigningKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSigningKeyClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSigningKeyClass) Alloc() SLSSigningKey {
	rv := objc.Send[SLSSigningKey](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSigningKey.CreateSignatureForMessage]
//   - [SLSSigningKey.EncodeWithCoder]
//   - [SLSSigningKey.SigningContext]
//   - [SLSSigningKey.InitWithCoder]
//   - [SLSSigningKey.InitWithData]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey
type SLSSigningKey struct {
	objectivec.Object
}

// SLSSigningKeyFromID constructs a [SLSSigningKey] from an objc.ID.
func SLSSigningKeyFromID(id objc.ID) SLSSigningKey {
	return SLSSigningKey{objectivec.Object{ID: id}}
}

// Ensure SLSSigningKey implements ISLSSigningKey.
var _ ISLSSigningKey = SLSSigningKey{}

// An interface definition for the [SLSSigningKey] class.
//
// # Methods
//
//   - [ISLSSigningKey.CreateSignatureForMessage]
//   - [ISLSSigningKey.EncodeWithCoder]
//   - [ISLSSigningKey.SigningContext]
//   - [ISLSSigningKey.InitWithCoder]
//   - [ISLSSigningKey.InitWithData]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey
type ISLSSigningKey interface {
	objectivec.IObject

	// Topic: Methods

	CreateSignatureForMessage(message objectivec.IObject) objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	SigningContext() objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) SLSSigningKey
	InitWithData(data objectivec.IObject) SLSSigningKey
}

// Init initializes the instance.
func (s SLSSigningKey) Init() SLSSigningKey {
	rv := objc.Send[SLSSigningKey](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSigningKey) Autorelease() SLSSigningKey {
	rv := objc.Send[SLSSigningKey](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSigningKey creates a new SLSSigningKey instance.
func NewSLSSigningKey() SLSSigningKey {
	class := getSLSSigningKeyClass()
	rv := objc.Send[SLSSigningKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/initWithCoder:
func NewSLSSigningKeyWithCoder(coder objectivec.IObject) SLSSigningKey {
	instance := getSLSSigningKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSSigningKeyFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/initWithData:
func NewSLSSigningKeyWithData(data objectivec.IObject) SLSSigningKey {
	instance := getSLSSigningKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return SLSSigningKeyFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/createSignatureForMessage:
func (s SLSSigningKey) CreateSignatureForMessage(message objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createSignatureForMessage:"), message)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/encodeWithCoder:
func (s SLSSigningKey) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/signingContext
func (s SLSSigningKey) SigningContext() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("signingContext"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/initWithCoder:
func (s SLSSigningKey) InitWithCoder(coder foundation.INSCoder) SLSSigningKey {
	rv := objc.Send[SLSSigningKey](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/initWithData:
func (s SLSSigningKey) InitWithData(data objectivec.IObject) SLSSigningKey {
	rv := objc.Send[SLSSigningKey](s.ID, objc.Sel("initWithData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/key
func (_SLSSigningKeyClass SLSSigningKeyClass) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSSigningKeyClass.class), objc.Sel("key"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/keyWithData:
func (_SLSSigningKeyClass SLSSigningKeyClass) KeyWithData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSSigningKeyClass.class), objc.Sel("keyWithData:"), data)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSigningKey/supportsSecureCoding
func (_SLSSigningKeyClass SLSSigningKeyClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLSSigningKeyClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
