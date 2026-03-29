// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIEncryptionCreator] class.
var (
	_DIEncryptionCreatorClass     DIEncryptionCreatorClass
	_DIEncryptionCreatorClassOnce sync.Once
)

func getDIEncryptionCreatorClass() DIEncryptionCreatorClass {
	_DIEncryptionCreatorClassOnce.Do(func() {
		_DIEncryptionCreatorClass = DIEncryptionCreatorClass{class: objc.GetClass("DIEncryptionCreator")}
	})
	return _DIEncryptionCreatorClass
}

// GetDIEncryptionCreatorClass returns the class object for DIEncryptionCreator.
func GetDIEncryptionCreatorClass() DIEncryptionCreatorClass {
	return getDIEncryptionCreatorClass()
}

type DIEncryptionCreatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIEncryptionCreatorClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIEncryptionCreatorClass) Alloc() DIEncryptionCreator {
	rv := objc.Send[DIEncryptionCreator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIEncryptionCreator.AddPublicKeyEntryWithXpcHandlerError]
//   - [DIEncryptionCreator.AddSymmetricKeyEntryWithError]
//   - [DIEncryptionCreator.CreateAndStoreInSystemKeychainWithAccountError]
//   - [DIEncryptionCreator.CreateParams]
//   - [DIEncryptionCreator.SetCreateParams]
//   - [DIEncryptionCreator.CreateWithXpcHandlerError]
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator
type DIEncryptionCreator struct {
	DIEncryptionFrontend
}

// DIEncryptionCreatorFromID constructs a [DIEncryptionCreator] from an objc.ID.
func DIEncryptionCreatorFromID(id objc.ID) DIEncryptionCreator {
	return DIEncryptionCreator{DIEncryptionFrontend: DIEncryptionFrontendFromID(id)}
}
// Ensure DIEncryptionCreator implements IDIEncryptionCreator.
var _ IDIEncryptionCreator = DIEncryptionCreator{}

// An interface definition for the [DIEncryptionCreator] class.
//
// # Methods
//
//   - [IDIEncryptionCreator.AddPublicKeyEntryWithXpcHandlerError]
//   - [IDIEncryptionCreator.AddSymmetricKeyEntryWithError]
//   - [IDIEncryptionCreator.CreateAndStoreInSystemKeychainWithAccountError]
//   - [IDIEncryptionCreator.CreateParams]
//   - [IDIEncryptionCreator.SetCreateParams]
//   - [IDIEncryptionCreator.CreateWithXpcHandlerError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator
type IDIEncryptionCreator interface {
	IDIEncryptionFrontend

	// Topic: Methods

	AddPublicKeyEntryWithXpcHandlerError(handler objectivec.IObject) (bool, error)
	AddSymmetricKeyEntryWithError() (bool, error)
	CreateAndStoreInSystemKeychainWithAccountError(account objectivec.IObject) (bool, error)
	CreateParams() IDICreateParams
	SetCreateParams(value IDICreateParams)
	CreateWithXpcHandlerError(handler objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (d DIEncryptionCreator) Init() DIEncryptionCreator {
	rv := objc.Send[DIEncryptionCreator](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIEncryptionCreator) Autorelease() DIEncryptionCreator {
	rv := objc.Send[DIEncryptionCreator](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIEncryptionCreator creates a new DIEncryptionCreator instance.
func NewDIEncryptionCreator() DIEncryptionCreator {
	class := getDIEncryptionCreatorClass()
	rv := objc.Send[DIEncryptionCreator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithCoder:
func NewDIEncryptionCreatorWithCoder(coder objectivec.IObject) DIEncryptionCreator {
	instance := getDIEncryptionCreatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIEncryptionCreatorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithParams:
func NewDIEncryptionCreatorWithParams(params objectivec.IObject) DIEncryptionCreator {
	instance := getDIEncryptionCreatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return DIEncryptionCreatorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator/addPublicKeyEntryWithXpcHandler:error:
func (d DIEncryptionCreator) AddPublicKeyEntryWithXpcHandlerError(handler objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addPublicKeyEntryWithXpcHandler:error:"), handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addPublicKeyEntryWithXpcHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator/addSymmetricKeyEntryWithError:
func (d DIEncryptionCreator) AddSymmetricKeyEntryWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addSymmetricKeyEntryWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addSymmetricKeyEntryWithError: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator/createAndStoreInSystemKeychainWithAccount:error:
func (d DIEncryptionCreator) CreateAndStoreInSystemKeychainWithAccountError(account objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createAndStoreInSystemKeychainWithAccount:error:"), account, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createAndStoreInSystemKeychainWithAccount:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator/createWithXpcHandler:error:
func (d DIEncryptionCreator) CreateWithXpcHandlerError(handler objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("createWithXpcHandler:error:"), handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createWithXpcHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator/getPublicKeyWithCertificate:error:
func (_DIEncryptionCreatorClass DIEncryptionCreatorClass) GetPublicKeyWithCertificateError(certificate objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIEncryptionCreatorClass.class), objc.Sel("getPublicKeyWithCertificate:error:"), certificate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionCreator/createParams
func (d DIEncryptionCreator) CreateParams() IDICreateParams {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("createParams"))
	return DICreateParamsFromID(objc.ID(rv))
}
func (d DIEncryptionCreator) SetCreateParams(value IDICreateParams) {
	objc.Send[struct{}](d.ID, objc.Sel("setCreateParams:"), value)
}

