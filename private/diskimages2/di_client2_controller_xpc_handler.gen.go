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

// The class instance for the [DIClient2Controller_XPCHandler] class.
var (
	_DIClient2Controller_XPCHandlerClass     DIClient2Controller_XPCHandlerClass
	_DIClient2Controller_XPCHandlerClassOnce sync.Once
)

func getDIClient2Controller_XPCHandlerClass() DIClient2Controller_XPCHandlerClass {
	_DIClient2Controller_XPCHandlerClassOnce.Do(func() {
		_DIClient2Controller_XPCHandlerClass = DIClient2Controller_XPCHandlerClass{class: objc.GetClass("DIClient2Controller_XPCHandler")}
	})
	return _DIClient2Controller_XPCHandlerClass
}

// GetDIClient2Controller_XPCHandlerClass returns the class object for DIClient2Controller_XPCHandler.
func GetDIClient2Controller_XPCHandlerClass() DIClient2Controller_XPCHandlerClass {
	return getDIClient2Controller_XPCHandlerClass()
}

type DIClient2Controller_XPCHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIClient2Controller_XPCHandlerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIClient2Controller_XPCHandlerClass) Alloc() DIClient2Controller_XPCHandler {
	rv := objc.Send[DIClient2Controller_XPCHandler](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIClient2Controller_XPCHandler.GUIAskForPassphraseWithEncryptionFrontendUsageError]
//   - [DIClient2Controller_XPCHandler.CreateAndStoreInSystemKeychainWithCreatorAccountError]
//   - [DIClient2Controller_XPCHandler.KeychainUnlockWithEncryptionUnlockerError]
//   - [DIClient2Controller_XPCHandler.NewAttachWithParamsError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIClient2Controller_XPCHandler
type DIClient2Controller_XPCHandler struct {
	DIBaseXPCHandler
}

// DIClient2Controller_XPCHandlerFromID constructs a [DIClient2Controller_XPCHandler] from an objc.ID.
func DIClient2Controller_XPCHandlerFromID(id objc.ID) DIClient2Controller_XPCHandler {
	return DIClient2Controller_XPCHandler{DIBaseXPCHandler: DIBaseXPCHandlerFromID(id)}
}

// Ensure DIClient2Controller_XPCHandler implements IDIClient2Controller_XPCHandler.
var _ IDIClient2Controller_XPCHandler = DIClient2Controller_XPCHandler{}

// An interface definition for the [DIClient2Controller_XPCHandler] class.
//
// # Methods
//
//   - [IDIClient2Controller_XPCHandler.GUIAskForPassphraseWithEncryptionFrontendUsageError]
//   - [IDIClient2Controller_XPCHandler.CreateAndStoreInSystemKeychainWithCreatorAccountError]
//   - [IDIClient2Controller_XPCHandler.KeychainUnlockWithEncryptionUnlockerError]
//   - [IDIClient2Controller_XPCHandler.NewAttachWithParamsError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIClient2Controller_XPCHandler
type IDIClient2Controller_XPCHandler interface {
	IDIBaseXPCHandler

	// Topic: Methods

	GUIAskForPassphraseWithEncryptionFrontendUsageError(frontend objectivec.IObject, usage int64) (bool, error)
	CreateAndStoreInSystemKeychainWithCreatorAccountError(creator objectivec.IObject, account objectivec.IObject) (bool, error)
	KeychainUnlockWithEncryptionUnlockerError(unlocker objectivec.IObject) (bool, error)
	NewAttachWithParamsError(params objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (d DIClient2Controller_XPCHandler) Init() DIClient2Controller_XPCHandler {
	rv := objc.Send[DIClient2Controller_XPCHandler](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIClient2Controller_XPCHandler) Autorelease() DIClient2Controller_XPCHandler {
	rv := objc.Send[DIClient2Controller_XPCHandler](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIClient2Controller_XPCHandler creates a new DIClient2Controller_XPCHandler instance.
func NewDIClient2Controller_XPCHandler() DIClient2Controller_XPCHandler {
	class := getDIClient2Controller_XPCHandlerClass()
	rv := objc.Send[DIClient2Controller_XPCHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2Controller_XPCHandler/GUIAskForPassphraseWithEncryptionFrontend:usage:error:
func (d DIClient2Controller_XPCHandler) GUIAskForPassphraseWithEncryptionFrontendUsageError(frontend objectivec.IObject, usage int64) (bool, error) {
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

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2Controller_XPCHandler/createAndStoreInSystemKeychainWithCreator:account:error:
func (d DIClient2Controller_XPCHandler) CreateAndStoreInSystemKeychainWithCreatorAccountError(creator objectivec.IObject, account objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2Controller_XPCHandler/keychainUnlockWithEncryptionUnlocker:error:
func (d DIClient2Controller_XPCHandler) KeychainUnlockWithEncryptionUnlockerError(unlocker objectivec.IObject) (bool, error) {
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

// See: https://developer.apple.com/documentation/DiskImages2/DIClient2Controller_XPCHandler/newAttachWithParams:error:
func (d DIClient2Controller_XPCHandler) NewAttachWithParamsError(params objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("newAttachWithParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
