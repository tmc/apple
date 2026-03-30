// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSecureModelDecryptCredential] class.
var (
	_MLSecureModelDecryptCredentialClass     MLSecureModelDecryptCredentialClass
	_MLSecureModelDecryptCredentialClassOnce sync.Once
)

func getMLSecureModelDecryptCredentialClass() MLSecureModelDecryptCredentialClass {
	_MLSecureModelDecryptCredentialClassOnce.Do(func() {
		_MLSecureModelDecryptCredentialClass = MLSecureModelDecryptCredentialClass{class: objc.GetClass("MLSecureModelDecryptCredential")}
	})
	return _MLSecureModelDecryptCredentialClass
}

// GetMLSecureModelDecryptCredentialClass returns the class object for MLSecureModelDecryptCredential.
func GetMLSecureModelDecryptCredentialClass() MLSecureModelDecryptCredentialClass {
	return getMLSecureModelDecryptCredentialClass()
}

type MLSecureModelDecryptCredentialClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSecureModelDecryptCredentialClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSecureModelDecryptCredentialClass) Alloc() MLSecureModelDecryptCredential {
	rv := objc.Send[MLSecureModelDecryptCredential](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSecureModelDecryptCredential.CryptoKey]
//   - [MLSecureModelDecryptCredential.SetCryptoKey]
//   - [MLSecureModelDecryptCredential.EncodeWithCoder]
//   - [MLSecureModelDecryptCredential.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential
type MLSecureModelDecryptCredential struct {
	objectivec.Object
}

// MLSecureModelDecryptCredentialFromID constructs a [MLSecureModelDecryptCredential] from an objc.ID.
func MLSecureModelDecryptCredentialFromID(id objc.ID) MLSecureModelDecryptCredential {
	return MLSecureModelDecryptCredential{objectivec.Object{ID: id}}
}

// Ensure MLSecureModelDecryptCredential implements IMLSecureModelDecryptCredential.
var _ IMLSecureModelDecryptCredential = MLSecureModelDecryptCredential{}

// An interface definition for the [MLSecureModelDecryptCredential] class.
//
// # Methods
//
//   - [IMLSecureModelDecryptCredential.CryptoKey]
//   - [IMLSecureModelDecryptCredential.SetCryptoKey]
//   - [IMLSecureModelDecryptCredential.EncodeWithCoder]
//   - [IMLSecureModelDecryptCredential.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential
type IMLSecureModelDecryptCredential interface {
	objectivec.IObject

	// Topic: Methods

	CryptoKey() int64
	SetCryptoKey(value int64)
	EncodeWithCoder(coder foundation.INSCoder)
	InitWithCoder(coder foundation.INSCoder) MLSecureModelDecryptCredential
}

// Init initializes the instance.
func (s MLSecureModelDecryptCredential) Init() MLSecureModelDecryptCredential {
	rv := objc.Send[MLSecureModelDecryptCredential](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSecureModelDecryptCredential) Autorelease() MLSecureModelDecryptCredential {
	rv := objc.Send[MLSecureModelDecryptCredential](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSecureModelDecryptCredential creates a new MLSecureModelDecryptCredential instance.
func NewMLSecureModelDecryptCredential() MLSecureModelDecryptCredential {
	class := getMLSecureModelDecryptCredentialClass()
	rv := objc.Send[MLSecureModelDecryptCredential](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential/initWithCoder:
func NewSecureModelDecryptCredentialWithCoder(coder objectivec.IObject) MLSecureModelDecryptCredential {
	instance := getMLSecureModelDecryptCredentialClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLSecureModelDecryptCredentialFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential/encodeWithCoder:
func (s MLSecureModelDecryptCredential) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential/initWithCoder:
func (s MLSecureModelDecryptCredential) InitWithCoder(coder foundation.INSCoder) MLSecureModelDecryptCredential {
	rv := objc.Send[MLSecureModelDecryptCredential](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential/supportsSecureCoding
func (_MLSecureModelDecryptCredentialClass MLSecureModelDecryptCredentialClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLSecureModelDecryptCredentialClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSecureModelDecryptCredential/cryptoKey
func (s MLSecureModelDecryptCredential) CryptoKey() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("cryptoKey"))
	return rv
}
func (s MLSecureModelDecryptCredential) SetCryptoKey(value int64) {
	objc.Send[struct{}](s.ID, objc.Sel("setCryptoKey:"), value)
}
