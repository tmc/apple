// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FetchRestoreKeysCryptoKitWrapper] class.
var (
	_FetchRestoreKeysCryptoKitWrapperClass     FetchRestoreKeysCryptoKitWrapperClass
	_FetchRestoreKeysCryptoKitWrapperClassOnce sync.Once
)

func getFetchRestoreKeysCryptoKitWrapperClass() FetchRestoreKeysCryptoKitWrapperClass {
	_FetchRestoreKeysCryptoKitWrapperClassOnce.Do(func() {
		_FetchRestoreKeysCryptoKitWrapperClass = FetchRestoreKeysCryptoKitWrapperClass{class: objc.GetClass("FetchRestoreKeys.CryptoKitWrapper")}
	})
	return _FetchRestoreKeysCryptoKitWrapperClass
}

// GetFetchRestoreKeysCryptoKitWrapperClass returns the class object for FetchRestoreKeys.CryptoKitWrapper.
func GetFetchRestoreKeysCryptoKitWrapperClass() FetchRestoreKeysCryptoKitWrapperClass {
	return getFetchRestoreKeysCryptoKitWrapperClass()
}

type FetchRestoreKeysCryptoKitWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FetchRestoreKeysCryptoKitWrapperClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FetchRestoreKeysCryptoKitWrapperClass) Alloc() FetchRestoreKeysCryptoKitWrapper {
	rv := objc.Send[FetchRestoreKeysCryptoKitWrapper](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/FetchRestoreKeys.CryptoKitWrapper
type FetchRestoreKeysCryptoKitWrapper struct {
	objectivec.Object
}

// FetchRestoreKeysCryptoKitWrapperFromID constructs a [FetchRestoreKeysCryptoKitWrapper] from an objc.ID.
func FetchRestoreKeysCryptoKitWrapperFromID(id objc.ID) FetchRestoreKeysCryptoKitWrapper {
	return FetchRestoreKeysCryptoKitWrapper{objectivec.Object{ID: id}}
}
// Ensure FetchRestoreKeysCryptoKitWrapper implements IFetchRestoreKeysCryptoKitWrapper.
var _ IFetchRestoreKeysCryptoKitWrapper = FetchRestoreKeysCryptoKitWrapper{}

// An interface definition for the [FetchRestoreKeysCryptoKitWrapper] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/FetchRestoreKeys.CryptoKitWrapper
type IFetchRestoreKeysCryptoKitWrapper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (f FetchRestoreKeysCryptoKitWrapper) Init() FetchRestoreKeysCryptoKitWrapper {
	rv := objc.Send[FetchRestoreKeysCryptoKitWrapper](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FetchRestoreKeysCryptoKitWrapper) Autorelease() FetchRestoreKeysCryptoKitWrapper {
	rv := objc.Send[FetchRestoreKeysCryptoKitWrapper](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFetchRestoreKeysCryptoKitWrapper creates a new FetchRestoreKeysCryptoKitWrapper instance.
func NewFetchRestoreKeysCryptoKitWrapper() FetchRestoreKeysCryptoKitWrapper {
	class := getFetchRestoreKeysCryptoKitWrapperClass()
	rv := objc.Send[FetchRestoreKeysCryptoKitWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/FetchRestoreKeys.CryptoKitWrapper/convertPrivateKeyToPEMWithX963PrivateKey:error:
func (_FetchRestoreKeysCryptoKitWrapperClass FetchRestoreKeysCryptoKitWrapperClass) ConvertPrivateKeyToPEMWithX963PrivateKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_FetchRestoreKeysCryptoKitWrapperClass.class), objc.Sel("convertPrivateKeyToPEMWithX963PrivateKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/FetchRestoreKeys.CryptoKitWrapper/convertPrivateKeyTox963WithPemPrivateKey:error:
func (_FetchRestoreKeysCryptoKitWrapperClass FetchRestoreKeysCryptoKitWrapperClass) ConvertPrivateKeyTox963WithPemPrivateKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_FetchRestoreKeysCryptoKitWrapperClass.class), objc.Sel("convertPrivateKeyTox963WithPemPrivateKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/FetchRestoreKeys.CryptoKitWrapper/unwrapEncryptionKeyWithWrappedKey:encapsulatedKey:privateKey:error:
func (_FetchRestoreKeysCryptoKitWrapperClass FetchRestoreKeysCryptoKitWrapperClass) UnwrapEncryptionKeyWithWrappedKeyEncapsulatedKeyPrivateKeyError(key objectivec.IObject, key2 objectivec.IObject, key3 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_FetchRestoreKeysCryptoKitWrapperClass.class), objc.Sel("unwrapEncryptionKeyWithWrappedKey:encapsulatedKey:privateKey:error:"), key, key2, key3, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/FetchRestoreKeys.CryptoKitWrapper/wrappedDataDictionaryWithCertWithPlainText:certificate:error:
func (_FetchRestoreKeysCryptoKitWrapperClass FetchRestoreKeysCryptoKitWrapperClass) WrappedDataDictionaryWithCertWithPlainTextCertificateError(text objectivec.IObject, certificate objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_FetchRestoreKeysCryptoKitWrapperClass.class), objc.Sel("wrappedDataDictionaryWithCertWithPlainText:certificate:error:"), text, certificate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

