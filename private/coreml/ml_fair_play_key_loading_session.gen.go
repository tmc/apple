// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFairPlayKeyLoadingSession] class.
var (
	_MLFairPlayKeyLoadingSessionClass     MLFairPlayKeyLoadingSessionClass
	_MLFairPlayKeyLoadingSessionClassOnce sync.Once
)

func getMLFairPlayKeyLoadingSessionClass() MLFairPlayKeyLoadingSessionClass {
	_MLFairPlayKeyLoadingSessionClassOnce.Do(func() {
		_MLFairPlayKeyLoadingSessionClass = MLFairPlayKeyLoadingSessionClass{class: objc.GetClass("MLFairPlayKeyLoadingSession")}
	})
	return _MLFairPlayKeyLoadingSessionClass
}

// GetMLFairPlayKeyLoadingSessionClass returns the class object for MLFairPlayKeyLoadingSession.
func GetMLFairPlayKeyLoadingSessionClass() MLFairPlayKeyLoadingSessionClass {
	return getMLFairPlayKeyLoadingSessionClass()
}

type MLFairPlayKeyLoadingSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFairPlayKeyLoadingSessionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFairPlayKeyLoadingSessionClass) Alloc() MLFairPlayKeyLoadingSession {
	rv := objc.Send[MLFairPlayKeyLoadingSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFairPlayKeyLoadingSession.GenerateKeyRequestForKeyIdentifierTeamIdentifierError]
//   - [MLFairPlayKeyLoadingSession.GeneratePersistentKeyBlobFromKeyResponseError]
//   - [MLFairPlayKeyLoadingSession.KeyIdentifier]
//   - [MLFairPlayKeyLoadingSession.SetKeyIdentifier]
//   - [MLFairPlayKeyLoadingSession.SessionContext]
//   - [MLFairPlayKeyLoadingSession.SetSessionContext]
//   - [MLFairPlayKeyLoadingSession.SessionID]
//   - [MLFairPlayKeyLoadingSession.SetSessionID]
//   - [MLFairPlayKeyLoadingSession.TransformKeyIdentifierError]
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession
type MLFairPlayKeyLoadingSession struct {
	objectivec.Object
}

// MLFairPlayKeyLoadingSessionFromID constructs a [MLFairPlayKeyLoadingSession] from an objc.ID.
func MLFairPlayKeyLoadingSessionFromID(id objc.ID) MLFairPlayKeyLoadingSession {
	return MLFairPlayKeyLoadingSession{objectivec.Object{ID: id}}
}

// Ensure MLFairPlayKeyLoadingSession implements IMLFairPlayKeyLoadingSession.
var _ IMLFairPlayKeyLoadingSession = MLFairPlayKeyLoadingSession{}

// An interface definition for the [MLFairPlayKeyLoadingSession] class.
//
// # Methods
//
//   - [IMLFairPlayKeyLoadingSession.GenerateKeyRequestForKeyIdentifierTeamIdentifierError]
//   - [IMLFairPlayKeyLoadingSession.GeneratePersistentKeyBlobFromKeyResponseError]
//   - [IMLFairPlayKeyLoadingSession.KeyIdentifier]
//   - [IMLFairPlayKeyLoadingSession.SetKeyIdentifier]
//   - [IMLFairPlayKeyLoadingSession.SessionContext]
//   - [IMLFairPlayKeyLoadingSession.SetSessionContext]
//   - [IMLFairPlayKeyLoadingSession.SessionID]
//   - [IMLFairPlayKeyLoadingSession.SetSessionID]
//   - [IMLFairPlayKeyLoadingSession.TransformKeyIdentifierError]
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession
type IMLFairPlayKeyLoadingSession interface {
	objectivec.IObject

	// Topic: Methods

	GenerateKeyRequestForKeyIdentifierTeamIdentifierError(identifier objectivec.IObject, identifier2 objectivec.IObject) (objectivec.IObject, error)
	GeneratePersistentKeyBlobFromKeyResponseError(response objectivec.IObject) (objectivec.IObject, error)
	KeyIdentifier() string
	SetKeyIdentifier(value string)
	SessionContext() objectivec.IObject
	SetSessionContext(value objectivec.IObject)
	SessionID() uint32
	SetSessionID(value uint32)
	TransformKeyIdentifierError(identifier objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (m MLFairPlayKeyLoadingSession) Init() MLFairPlayKeyLoadingSession {
	rv := objc.Send[MLFairPlayKeyLoadingSession](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLFairPlayKeyLoadingSession) Autorelease() MLFairPlayKeyLoadingSession {
	rv := objc.Send[MLFairPlayKeyLoadingSession](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFairPlayKeyLoadingSession creates a new MLFairPlayKeyLoadingSession instance.
func NewMLFairPlayKeyLoadingSession() MLFairPlayKeyLoadingSession {
	class := getMLFairPlayKeyLoadingSessionClass()
	rv := objc.Send[MLFairPlayKeyLoadingSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/generateKeyRequestForKeyIdentifier:teamIdentifier:error:
func (m MLFairPlayKeyLoadingSession) GenerateKeyRequestForKeyIdentifierTeamIdentifierError(identifier objectivec.IObject, identifier2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("generateKeyRequestForKeyIdentifier:teamIdentifier:error:"), identifier, identifier2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/generatePersistentKeyBlobFromKeyResponse:error:
func (m MLFairPlayKeyLoadingSession) GeneratePersistentKeyBlobFromKeyResponseError(response objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("generatePersistentKeyBlobFromKeyResponse:error:"), response, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/transformKeyIdentifier:error:
func (m MLFairPlayKeyLoadingSession) TransformKeyIdentifierError(identifier objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("transformKeyIdentifier:error:"), identifier, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/keyIdentifier
func (m MLFairPlayKeyLoadingSession) KeyIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keyIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLFairPlayKeyLoadingSession) SetKeyIdentifier(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setKeyIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/sessionContext
func (m MLFairPlayKeyLoadingSession) SessionContext() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sessionContext"))
	return objectivec.Object{ID: rv}
}
func (m MLFairPlayKeyLoadingSession) SetSessionContext(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setSessionContext:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/sessionID
func (m MLFairPlayKeyLoadingSession) SessionID() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("sessionID"))
	return rv
}
func (m MLFairPlayKeyLoadingSession) SetSessionID(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setSessionID:"), value)
}
