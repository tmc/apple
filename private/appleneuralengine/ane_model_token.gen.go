// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEModelToken] class.
var (
	_ANEModelTokenClass     ANEModelTokenClass
	_ANEModelTokenClassOnce sync.Once
)

func getANEModelTokenClass() ANEModelTokenClass {
	_ANEModelTokenClassOnce.Do(func() {
		_ANEModelTokenClass = ANEModelTokenClass{class: objc.GetClass("_ANEModelToken")}
	})
	return _ANEModelTokenClass
}

// GetANEModelTokenClass returns the class object for _ANEModelToken.
func GetANEModelTokenClass() ANEModelTokenClass {
	return getANEModelTokenClass()
}

type ANEModelTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEModelTokenClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEModelTokenClass) Alloc() ANEModelToken {
	rv := objc.Send[ANEModelToken](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEModelToken.CsIdentity]
//   - [ANEModelToken.ModelIdentifier]
//   - [ANEModelToken.ProcessIdentifier]
//   - [ANEModelToken.TeamIdentity]
//   - [ANEModelToken.InitWithAuditTokenModelIdentifierProcessIdentifier]
//   - [ANEModelToken.InitWithCsIdentityTeamIdentityModelIdentifierProcessIdentifier]
type ANEModelToken struct {
	objectivec.Object
}

// ANEModelTokenFromID constructs a [ANEModelToken] from an objc.ID.
func ANEModelTokenFromID(id objc.ID) ANEModelToken {
	return ANEModelToken{objectivec.Object{ID: id}}
}

// Ensure ANEModelToken implements IANEModelToken.
var _ IANEModelToken = ANEModelToken{}

// An interface definition for the [ANEModelToken] class.
//
// # Methods
//
//   - [IANEModelToken.CsIdentity]
//   - [IANEModelToken.ModelIdentifier]
//   - [IANEModelToken.ProcessIdentifier]
//   - [IANEModelToken.TeamIdentity]
//   - [IANEModelToken.InitWithAuditTokenModelIdentifierProcessIdentifier]
//   - [IANEModelToken.InitWithCsIdentityTeamIdentityModelIdentifierProcessIdentifier]
type IANEModelToken interface {
	objectivec.IObject

	// Topic: Methods

	CsIdentity() string
	ModelIdentifier() string
	ProcessIdentifier() int
	TeamIdentity() string
	InitWithAuditTokenModelIdentifierProcessIdentifier(token unsafe.Pointer, identifier objectivec.IObject, identifier2 int) ANEModelToken
	InitWithCsIdentityTeamIdentityModelIdentifierProcessIdentifier(identity objectivec.IObject, identity2 objectivec.IObject, identifier objectivec.IObject, identifier2 int) ANEModelToken
}

// Init initializes the instance.
func (a ANEModelToken) Init() ANEModelToken {
	rv := objc.Send[ANEModelToken](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEModelToken) Autorelease() ANEModelToken {
	rv := objc.Send[ANEModelToken](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEModelToken creates a new ANEModelToken instance.
func NewANEModelToken() ANEModelToken {
	class := getANEModelTokenClass()
	rv := objc.Send[ANEModelToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEModelTokenWithAuditTokenModelIdentifierProcessIdentifier(token unsafe.Pointer, identifier objectivec.IObject, identifier2 int) ANEModelToken {
	instance := getANEModelTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAuditToken:modelIdentifier:processIdentifier:"), token, identifier, identifier2)
	return ANEModelTokenFromID(rv)
}

func NewANEModelTokenWithCsIdentityTeamIdentityModelIdentifierProcessIdentifier(identity objectivec.IObject, identity2 objectivec.IObject, identifier objectivec.IObject, identifier2 int) ANEModelToken {
	instance := getANEModelTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCsIdentity:teamIdentity:modelIdentifier:processIdentifier:"), identity, identity2, identifier, identifier2)
	return ANEModelTokenFromID(rv)
}

func (a ANEModelToken) InitWithAuditTokenModelIdentifierProcessIdentifier(token unsafe.Pointer, identifier objectivec.IObject, identifier2 int) ANEModelToken {
	rv := objc.Send[ANEModelToken](a.ID, objc.Sel("initWithAuditToken:modelIdentifier:processIdentifier:"), token, identifier, identifier2)
	return rv
}
func (a ANEModelToken) InitWithCsIdentityTeamIdentityModelIdentifierProcessIdentifier(identity objectivec.IObject, identity2 objectivec.IObject, identifier objectivec.IObject, identifier2 int) ANEModelToken {
	rv := objc.Send[ANEModelToken](a.ID, objc.Sel("initWithCsIdentity:teamIdentity:modelIdentifier:processIdentifier:"), identity, identity2, identifier, identifier2)
	return rv
}

func (_ANEModelTokenClass ANEModelTokenClass) CodeSigningIDForProcessIdentifier(iDFor unsafe.Pointer, identifier int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelTokenClass.class), objc.Sel("codeSigningIDFor:processIdentifier:"), iDFor, identifier)
	return objectivec.Object{ID: rv}
}
func (_ANEModelTokenClass ANEModelTokenClass) ProcessNameForIdentifier(for_ unsafe.Pointer, identifier int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelTokenClass.class), objc.Sel("processNameFor:identifier:"), for_, identifier)
	return objectivec.Object{ID: rv}
}
func (_ANEModelTokenClass ANEModelTokenClass) TeamIDForProcessIdentifier(iDFor unsafe.Pointer, identifier int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelTokenClass.class), objc.Sel("teamIDFor:processIdentifier:"), iDFor, identifier)
	return objectivec.Object{ID: rv}
}
func (_ANEModelTokenClass ANEModelTokenClass) TokenWithAuditTokenModelIdentifierProcessIdentifier(token unsafe.Pointer, identifier objectivec.IObject, identifier2 int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelTokenClass.class), objc.Sel("tokenWithAuditToken:modelIdentifier:processIdentifier:"), token, identifier, identifier2)
	return objectivec.Object{ID: rv}
}
func (_ANEModelTokenClass ANEModelTokenClass) TokenWithCsIdentityTeamIdentityModelIdentifierProcessIdentifier(identity objectivec.IObject, identity2 objectivec.IObject, identifier objectivec.IObject, identifier2 int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelTokenClass.class), objc.Sel("tokenWithCsIdentity:teamIdentity:modelIdentifier:processIdentifier:"), identity, identity2, identifier, identifier2)
	return objectivec.Object{ID: rv}
}

func (a ANEModelToken) CsIdentity() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("csIdentity"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEModelToken) ModelIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEModelToken) ProcessIdentifier() int {
	rv := objc.Send[int](a.ID, objc.Sel("processIdentifier"))
	return rv
}
func (a ANEModelToken) TeamIdentity() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("teamIdentity"))
	return foundation.NSStringFromID(rv).String()
}
