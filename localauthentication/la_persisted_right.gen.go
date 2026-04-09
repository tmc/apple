// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [LAPersistedRight] class.
var (
	_LAPersistedRightClass     LAPersistedRightClass
	_LAPersistedRightClassOnce sync.Once
)

func getLAPersistedRightClass() LAPersistedRightClass {
	_LAPersistedRightClassOnce.Do(func() {
		_LAPersistedRightClass = LAPersistedRightClass{class: objc.GetClass("LAPersistedRight")}
	})
	return _LAPersistedRightClass
}

// GetLAPersistedRightClass returns the class object for LAPersistedRight.
func GetLAPersistedRightClass() LAPersistedRightClass {
	return getLAPersistedRightClass()
}

type LAPersistedRightClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LAPersistedRightClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LAPersistedRightClass) Alloc() LAPersistedRight {
	rv := objc.Send[LAPersistedRight](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A right that gates access to a key and a secret.
//
// # Overview
//
// An [LAPersistedRight] is a right that’s backed by a unique key in the
// Secure Enclave with an access control list that matches the authorization
// requirements of the right. You can access the key that backs a right to
// perform cryptographic operations like encryption, decryption, signing, and
// verification.
//
// You can use the key that backs an [LAPersistedRight] to perform both public
// key and private key operations, but private key operations — like
// decryption, signing, and key exchange — are only available after you
// authorize the right. Public key operations like encryption and verification
// are always available.
//
// The following generates a right with the default authorization
// requirements, stores it in the [SharedStore] [LARightStore], and exports
// the public key so that you can use it to verify signatures that the
// corresponding private key produces:
//
// The following uses the private key associated with the right from the
// previous example to sign a challenge issued by a server:
//
// The signature operation occurs after verifying that the user has the proper
// authorization and confirming that the private key supports the given
// signing algorithm.
//
// # Accessing persistent data
//
//   - [LAPersistedRight.Key]: The private key that’s persisted by the right.
//   - [LAPersistedRight.Secret]: The data kept secret by the right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight
type LAPersistedRight struct {
	LARight
}

// LAPersistedRightFromID constructs a [LAPersistedRight] from an objc.ID.
//
// A right that gates access to a key and a secret.
func LAPersistedRightFromID(id objc.ID) LAPersistedRight {
	return LAPersistedRight{LARight: LARightFromID(id)}
}

// NOTE: LAPersistedRight adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LAPersistedRight] class.
//
// # Accessing persistent data
//
//   - [ILAPersistedRight.Key]: The private key that’s persisted by the right.
//   - [ILAPersistedRight.Secret]: The data kept secret by the right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight
type ILAPersistedRight interface {
	ILARight

	// Topic: Accessing persistent data

	// The private key that’s persisted by the right.
	Key() ILAPrivateKey
	// The data kept secret by the right.
	Secret() ILASecret
}

// Init initializes the instance.
func (p LAPersistedRight) Init() LAPersistedRight {
	rv := objc.Send[LAPersistedRight](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p LAPersistedRight) Autorelease() LAPersistedRight {
	rv := objc.Send[LAPersistedRight](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewLAPersistedRight creates a new LAPersistedRight instance.
func NewLAPersistedRight() LAPersistedRight {
	class := getLAPersistedRightClass()
	rv := objc.Send[LAPersistedRight](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a right with the authentication requirements you supply.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LARight/init(requirement:)
func NewPersistedRightWithRequirement(requirement ILAAuthenticationRequirement) LAPersistedRight {
	instance := getLAPersistedRightClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRequirement:"), requirement)
	return LAPersistedRightFromID(rv)
}

// The private key that’s persisted by the right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight/key
func (p LAPersistedRight) Key() ILAPrivateKey {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("key"))
	return LAPrivateKeyFromID(objc.ID(rv))
}

// The data kept secret by the right.
//
// See: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight/secret
func (p LAPersistedRight) Secret() ILASecret {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("secret"))
	return LASecretFromID(objc.ID(rv))
}

// A shared object that stores rights.
//
// See: https://developer.apple.com/documentation/localauthentication/larightstore/shared
func (_LAPersistedRightClass LAPersistedRightClass) Shared() LARightStore {
	rv := objc.Send[objc.ID](objc.ID(_LAPersistedRightClass.class), objc.Sel("sharedStore"))
	return LARightStoreFromID(objc.ID(rv))
}
func (_LAPersistedRightClass LAPersistedRightClass) SetShared(value LARightStore) {
	objc.Send[struct{}](objc.ID(_LAPersistedRightClass.class), objc.Sel("setSharedStore:"), value)
}
