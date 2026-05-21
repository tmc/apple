// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIClient2ControllerXPCHandler] class.
var (
	_DIClient2ControllerXPCHandlerClass     DIClient2ControllerXPCHandlerClass
	_DIClient2ControllerXPCHandlerClassOnce sync.Once
)

func getDIClient2ControllerXPCHandlerClass() DIClient2ControllerXPCHandlerClass {
	_DIClient2ControllerXPCHandlerClassOnce.Do(func() {
		_DIClient2ControllerXPCHandlerClass = DIClient2ControllerXPCHandlerClass{class: objc.GetClass("DIClient2Controller_XPCHandler")}
	})
	return _DIClient2ControllerXPCHandlerClass
}

// GetDIClient2ControllerXPCHandlerClass returns the class object for DIClient2Controller_XPCHandler.
func GetDIClient2ControllerXPCHandlerClass() DIClient2ControllerXPCHandlerClass {
	return getDIClient2ControllerXPCHandlerClass()
}

type DIClient2ControllerXPCHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIClient2ControllerXPCHandlerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIClient2ControllerXPCHandlerClass) Alloc() DIClient2ControllerXPCHandler {
	rv := objc.Send[DIClient2ControllerXPCHandler](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIClient2ControllerXPCHandler.GUIAskForPassphraseWithEncryptionFrontendUsageError]
//   - [DIClient2ControllerXPCHandler.CreateAndStoreInSystemKeychainWithCreatorAccountError]
//   - [DIClient2ControllerXPCHandler.KeychainUnlockWithEncryptionUnlockerError]
//   - [DIClient2ControllerXPCHandler.NewAttachWithParamsError]
type DIClient2ControllerXPCHandler struct {
	DIBaseXPCHandler
}

// DIClient2ControllerXPCHandlerFromID constructs a [DIClient2ControllerXPCHandler] from an objc.ID.
func DIClient2ControllerXPCHandlerFromID(id objc.ID) DIClient2ControllerXPCHandler {
	return DIClient2ControllerXPCHandler{DIBaseXPCHandler: DIBaseXPCHandlerFromID(id)}
}

// DIClient2Controller_XPCHandlerFromID is an alias for [DIClient2ControllerXPCHandlerFromID] for cross-framework compatibility.
func DIClient2Controller_XPCHandlerFromID(id objc.ID) DIClient2ControllerXPCHandler {
	return DIClient2ControllerXPCHandlerFromID(id)
}

// Ensure DIClient2ControllerXPCHandler implements IDIClient2ControllerXPCHandler.
var _ IDIClient2ControllerXPCHandler = DIClient2ControllerXPCHandler{}

// An interface definition for the [DIClient2ControllerXPCHandler] class.
//
// # Methods
//
//   - [IDIClient2ControllerXPCHandler.GUIAskForPassphraseWithEncryptionFrontendUsageError]
//   - [IDIClient2ControllerXPCHandler.CreateAndStoreInSystemKeychainWithCreatorAccountError]
//   - [IDIClient2ControllerXPCHandler.KeychainUnlockWithEncryptionUnlockerError]
//   - [IDIClient2ControllerXPCHandler.NewAttachWithParamsError]
type IDIClient2ControllerXPCHandler interface {
	IDIBaseXPCHandler

	// Topic: Methods

	GUIAskForPassphraseWithEncryptionFrontendUsageError(frontend objectivec.IObject, usage int64) (bool, error)
	CreateAndStoreInSystemKeychainWithCreatorAccountError(creator objectivec.IObject, account objectivec.IObject) (bool, error)
	KeychainUnlockWithEncryptionUnlockerError(unlocker objectivec.IObject) (bool, error)
	NewAttachWithParamsError(params objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (d DIClient2ControllerXPCHandler) Init() DIClient2ControllerXPCHandler {
	rv := objc.Send[DIClient2ControllerXPCHandler](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIClient2ControllerXPCHandler) Autorelease() DIClient2ControllerXPCHandler {
	rv := objc.Send[DIClient2ControllerXPCHandler](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIClient2ControllerXPCHandler creates a new DIClient2ControllerXPCHandler instance.
func NewDIClient2ControllerXPCHandler() DIClient2ControllerXPCHandler {
	class := getDIClient2ControllerXPCHandlerClass()
	rv := objc.Send[DIClient2ControllerXPCHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (d DIClient2ControllerXPCHandler) GUIAskForPassphraseWithEncryptionFrontendUsageError(frontend objectivec.IObject, usage int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("GUIAskForPassphraseWithEncryptionFrontend:usage:error:"), frontend, usage, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("GUIAskForPassphraseWithEncryptionFrontend:usage:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIClient2ControllerXPCHandler) CreateAndStoreInSystemKeychainWithCreatorAccountError(creator objectivec.IObject, account objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createAndStoreInSystemKeychainWithCreator:account:error:"), creator, account, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createAndStoreInSystemKeychainWithCreator:account:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIClient2ControllerXPCHandler) KeychainUnlockWithEncryptionUnlockerError(unlocker objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("keychainUnlockWithEncryptionUnlocker:error:"), unlocker, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("keychainUnlockWithEncryptionUnlocker:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIClient2ControllerXPCHandler) NewAttachWithParamsError(params objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("newAttachWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
