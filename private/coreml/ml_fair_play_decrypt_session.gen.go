// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFairPlayDecryptSession] class.
var (
	_MLFairPlayDecryptSessionClass     MLFairPlayDecryptSessionClass
	_MLFairPlayDecryptSessionClassOnce sync.Once
)

func getMLFairPlayDecryptSessionClass() MLFairPlayDecryptSessionClass {
	_MLFairPlayDecryptSessionClassOnce.Do(func() {
		_MLFairPlayDecryptSessionClass = MLFairPlayDecryptSessionClass{class: objc.GetClass("MLFairPlayDecryptSession")}
	})
	return _MLFairPlayDecryptSessionClass
}

// GetMLFairPlayDecryptSessionClass returns the class object for MLFairPlayDecryptSession.
func GetMLFairPlayDecryptSessionClass() MLFairPlayDecryptSessionClass {
	return getMLFairPlayDecryptSessionClass()
}

type MLFairPlayDecryptSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFairPlayDecryptSessionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFairPlayDecryptSessionClass) Alloc() MLFairPlayDecryptSession {
	rv := objc.Send[MLFairPlayDecryptSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFairPlayDecryptSession.ModelPath]
//   - [MLFairPlayDecryptSession.SetModelPath]
//   - [MLFairPlayDecryptSession.XpcConnection]
//   - [MLFairPlayDecryptSession.SetXpcConnection]
//   - [MLFairPlayDecryptSession.XpcProxy]
//   - [MLFairPlayDecryptSession.SetXpcProxy]
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSession
type MLFairPlayDecryptSession struct {
	objectivec.Object
}

// MLFairPlayDecryptSessionFromID constructs a [MLFairPlayDecryptSession] from an objc.ID.
func MLFairPlayDecryptSessionFromID(id objc.ID) MLFairPlayDecryptSession {
	return MLFairPlayDecryptSession{objectivec.Object{ID: id}}
}

// Ensure MLFairPlayDecryptSession implements IMLFairPlayDecryptSession.
var _ IMLFairPlayDecryptSession = MLFairPlayDecryptSession{}

// An interface definition for the [MLFairPlayDecryptSession] class.
//
// # Methods
//
//   - [IMLFairPlayDecryptSession.ModelPath]
//   - [IMLFairPlayDecryptSession.SetModelPath]
//   - [IMLFairPlayDecryptSession.XpcConnection]
//   - [IMLFairPlayDecryptSession.SetXpcConnection]
//   - [IMLFairPlayDecryptSession.XpcProxy]
//   - [IMLFairPlayDecryptSession.SetXpcProxy]
//
// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSession
type IMLFairPlayDecryptSession interface {
	objectivec.IObject

	// Topic: Methods

	ModelPath() string
	SetModelPath(value string)
	XpcConnection() foundation.NSXPCConnection
	SetXpcConnection(value foundation.NSXPCConnection)
	XpcProxy() objectivec.Object
	SetXpcProxy(value objectivec.Object)
}

// Init initializes the instance.
func (f MLFairPlayDecryptSession) Init() MLFairPlayDecryptSession {
	rv := objc.Send[MLFairPlayDecryptSession](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFairPlayDecryptSession) Autorelease() MLFairPlayDecryptSession {
	rv := objc.Send[MLFairPlayDecryptSession](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFairPlayDecryptSession creates a new MLFairPlayDecryptSession instance.
func NewMLFairPlayDecryptSession() MLFairPlayDecryptSession {
	class := getMLFairPlayDecryptSessionClass()
	rv := objc.Send[MLFairPlayDecryptSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSession/decryptSessionForModelAtPath:usesCodeSigningIdentityForEncryption:keyBlob:error:
func (_MLFairPlayDecryptSessionClass MLFairPlayDecryptSessionClass) DecryptSessionForModelAtPathUsesCodeSigningIdentityForEncryptionKeyBlobError(path objectivec.IObject, encryption bool, blob objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFairPlayDecryptSessionClass.class), objc.Sel("decryptSessionForModelAtPath:usesCodeSigningIdentityForEncryption:keyBlob:error:"), path, encryption, blob, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSession/modelPath
func (f MLFairPlayDecryptSession) ModelPath() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("modelPath"))
	return foundation.NSStringFromID(rv).String()
}
func (f MLFairPlayDecryptSession) SetModelPath(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setModelPath:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSession/xpcConnection
func (f MLFairPlayDecryptSession) XpcConnection() foundation.NSXPCConnection {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("xpcConnection"))
	return foundation.NSXPCConnectionFromID(objc.ID(rv))
}
func (f MLFairPlayDecryptSession) SetXpcConnection(value foundation.NSXPCConnection) {
	objc.Send[struct{}](f.ID, objc.Sel("setXpcConnection:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLFairPlayDecryptSession/xpcProxy
func (f MLFairPlayDecryptSession) XpcProxy() objectivec.Object {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("xpcProxy"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (f MLFairPlayDecryptSession) SetXpcProxy(value objectivec.Object) {
	objc.Send[struct{}](f.ID, objc.Sel("setXpcProxy:"), value)
}
