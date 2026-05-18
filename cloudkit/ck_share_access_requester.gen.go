// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKShareAccessRequester] class.
var (
	_CKShareAccessRequesterClass     CKShareAccessRequesterClass
	_CKShareAccessRequesterClassOnce sync.Once
)

func getCKShareAccessRequesterClass() CKShareAccessRequesterClass {
	_CKShareAccessRequesterClassOnce.Do(func() {
		_CKShareAccessRequesterClass = CKShareAccessRequesterClass{class: objc.GetClass("CKShareAccessRequester")}
	})
	return _CKShareAccessRequesterClass
}

// GetCKShareAccessRequesterClass returns the class object for CKShareAccessRequester.
func GetCKShareAccessRequesterClass() CKShareAccessRequesterClass {
	return getCKShareAccessRequesterClass()
}

type CKShareAccessRequesterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKShareAccessRequesterClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKShareAccessRequesterClass) Alloc() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [CKShareAccessRequester.Contact]: A displayable [CNContact] representing the requester.
//   - [CKShareAccessRequester.ParticipantLookupInfo]: Lookup information for the requester.
//   - [CKShareAccessRequester.UserIdentity]: The identity of the user requesting access to the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester
type CKShareAccessRequester struct {
	objectivec.Object
}

// CKShareAccessRequesterFromID constructs a [CKShareAccessRequester] from an objc.ID.
func CKShareAccessRequesterFromID(id objc.ID) CKShareAccessRequester {
	return CKShareAccessRequester{objectivec.Object{ID: id}}
}

// NOTE: CKShareAccessRequester adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKShareAccessRequester] class.
//
// # Instance Properties
//
//   - [ICKShareAccessRequester.Contact]: A displayable [CNContact] representing the requester.
//   - [ICKShareAccessRequester.ParticipantLookupInfo]: Lookup information for the requester.
//   - [ICKShareAccessRequester.UserIdentity]: The identity of the user requesting access to the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester
type ICKShareAccessRequester interface {
	objectivec.IObject

	// Topic: Instance Properties

	// A displayable [CNContact] representing the requester.
	Contact() unsafe.Pointer
	// Lookup information for the requester.
	ParticipantLookupInfo() ICKUserIdentityLookupInfo
	// The identity of the user requesting access to the share.
	UserIdentity() ICKUserIdentity

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKShareAccessRequester) Init() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKShareAccessRequester) Autorelease() CKShareAccessRequester {
	rv := objc.Send[CKShareAccessRequester](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareAccessRequester creates a new CKShareAccessRequester instance.
func NewCKShareAccessRequester() CKShareAccessRequester {
	class := getCKShareAccessRequesterClass()
	rv := objc.Send[CKShareAccessRequester](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKShareAccessRequester) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A displayable [CNContact] representing the requester.
//
// # Discussion
//
// If the requester doesn’t exist in the user’s contacts or is not
// accessible, returns a newly created [CNContact]. This provides formatted
// requester information suitable for display in the application’s UI.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/contact
func (c CKShareAccessRequester) Contact() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("contact"))
	return rv
}

// Lookup information for the requester.
//
// # Discussion
//
// Use this lookup info with [CKFetchShareParticipantsOperation] to fetch the
// corresponding participant. Once fetched, add the participant to the share
// to approve the requester.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/participantLookupInfo
func (c CKShareAccessRequester) ParticipantLookupInfo() ICKUserIdentityLookupInfo {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("participantLookupInfo"))
	return CKUserIdentityLookupInfoFromID(objc.ID(rv))
}

// The identity of the user requesting access to the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/AccessRequester/userIdentity
func (c CKShareAccessRequester) UserIdentity() ICKUserIdentity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentity"))
	return CKUserIdentityFromID(objc.ID(rv))
}
