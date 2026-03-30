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
func (f MLFairPlayKeyLoadingSession) Init() MLFairPlayKeyLoadingSession {
	rv := objc.Send[MLFairPlayKeyLoadingSession](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFairPlayKeyLoadingSession) Autorelease() MLFairPlayKeyLoadingSession {
	rv := objc.Send[MLFairPlayKeyLoadingSession](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFairPlayKeyLoadingSession creates a new MLFairPlayKeyLoadingSession instance.
func NewMLFairPlayKeyLoadingSession() MLFairPlayKeyLoadingSession {
	class := getMLFairPlayKeyLoadingSessionClass()
	rv := objc.Send[MLFairPlayKeyLoadingSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/generateKeyRequestForKeyIdentifier:teamIdentifier:error:
func (f MLFairPlayKeyLoadingSession) GenerateKeyRequestForKeyIdentifierTeamIdentifierError(identifier objectivec.IObject, identifier2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("generateKeyRequestForKeyIdentifier:teamIdentifier:error:"), identifier, identifier2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/generatePersistentKeyBlobFromKeyResponse:error:
func (f MLFairPlayKeyLoadingSession) GeneratePersistentKeyBlobFromKeyResponseError(response objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("generatePersistentKeyBlobFromKeyResponse:error:"), response, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/transformKeyIdentifier:error:
func (f MLFairPlayKeyLoadingSession) TransformKeyIdentifierError(identifier objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("transformKeyIdentifier:error:"), identifier, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/keyIdentifier
func (f MLFairPlayKeyLoadingSession) KeyIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("keyIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (f MLFairPlayKeyLoadingSession) SetKeyIdentifier(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setKeyIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/sessionContext
func (f MLFairPlayKeyLoadingSession) SessionContext() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sessionContext"))
	return objectivec.Object{ID: rv}
}
func (f MLFairPlayKeyLoadingSession) SetSessionContext(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setSessionContext:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayKeyLoadingSession/sessionID
func (f MLFairPlayKeyLoadingSession) SessionID() uint32 {
	rv := objc.Send[uint32](f.ID, objc.Sel("sessionID"))
	return rv
}
func (f MLFairPlayKeyLoadingSession) SetSessionID(value uint32) {
	objc.Send[struct{}](f.ID, objc.Sel("setSessionID:"), value)
}
