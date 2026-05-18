// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKShareBlockedIdentity] class.
var (
	_CKShareBlockedIdentityClass     CKShareBlockedIdentityClass
	_CKShareBlockedIdentityClassOnce sync.Once
)

func getCKShareBlockedIdentityClass() CKShareBlockedIdentityClass {
	_CKShareBlockedIdentityClassOnce.Do(func() {
		_CKShareBlockedIdentityClass = CKShareBlockedIdentityClass{class: objc.GetClass("CKShareBlockedIdentity")}
	})
	return _CKShareBlockedIdentityClass
}

// GetCKShareBlockedIdentityClass returns the class object for CKShareBlockedIdentity.
func GetCKShareBlockedIdentityClass() CKShareBlockedIdentityClass {
	return getCKShareBlockedIdentityClass()
}

type CKShareBlockedIdentityClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKShareBlockedIdentityClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKShareBlockedIdentityClass) Alloc() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [CKShareBlockedIdentity.Contact]: A displayable [CNContact] representing the blocked user.
//   - [CKShareBlockedIdentity.UserIdentity]: The identity of the user who has been blocked from requesting access to the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity
type CKShareBlockedIdentity struct {
	objectivec.Object
}

// CKShareBlockedIdentityFromID constructs a [CKShareBlockedIdentity] from an objc.ID.
func CKShareBlockedIdentityFromID(id objc.ID) CKShareBlockedIdentity {
	return CKShareBlockedIdentity{objectivec.Object{ID: id}}
}

// NOTE: CKShareBlockedIdentity adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKShareBlockedIdentity] class.
//
// # Instance Properties
//
//   - [ICKShareBlockedIdentity.Contact]: A displayable [CNContact] representing the blocked user.
//   - [ICKShareBlockedIdentity.UserIdentity]: The identity of the user who has been blocked from requesting access to the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity
type ICKShareBlockedIdentity interface {
	objectivec.IObject

	// Topic: Instance Properties

	// A displayable [CNContact] representing the blocked user.
	Contact() unsafe.Pointer
	// The identity of the user who has been blocked from requesting access to the share.
	UserIdentity() ICKUserIdentity

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKShareBlockedIdentity) Init() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKShareBlockedIdentity) Autorelease() CKShareBlockedIdentity {
	rv := objc.Send[CKShareBlockedIdentity](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareBlockedIdentity creates a new CKShareBlockedIdentity instance.
func NewCKShareBlockedIdentity() CKShareBlockedIdentity {
	class := getCKShareBlockedIdentityClass()
	rv := objc.Send[CKShareBlockedIdentity](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKShareBlockedIdentity) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A displayable [CNContact] representing the blocked user.
//
// # Discussion
//
// If the blocked identity does not exist in the user’s contacts or is not
// accessible, returns a newly created [CNContact]. This provides formatted
// blocked identity information suitable for display in the application’s
// UI.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity/contact
func (c CKShareBlockedIdentity) Contact() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("contact"))
	return rv
}

// The identity of the user who has been blocked from requesting access to the
// share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/BlockedIdentity/userIdentity
func (c CKShareBlockedIdentity) UserIdentity() ICKUserIdentity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentity"))
	return CKUserIdentityFromID(objc.ID(rv))
}
