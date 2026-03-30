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

// The class instance for the [DIEncryptionFrontend] class.
var (
	_DIEncryptionFrontendClass     DIEncryptionFrontendClass
	_DIEncryptionFrontendClassOnce sync.Once
)

func getDIEncryptionFrontendClass() DIEncryptionFrontendClass {
	_DIEncryptionFrontendClassOnce.Do(func() {
		_DIEncryptionFrontendClass = DIEncryptionFrontendClass{class: objc.GetClass("DIEncryptionFrontend")}
	})
	return _DIEncryptionFrontendClass
}

// GetDIEncryptionFrontendClass returns the class object for DIEncryptionFrontend.
func GetDIEncryptionFrontendClass() DIEncryptionFrontendClass {
	return getDIEncryptionFrontendClass()
}

type DIEncryptionFrontendClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIEncryptionFrontendClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIEncryptionFrontendClass) Alloc() DIEncryptionFrontend {
	rv := objc.Send[DIEncryptionFrontend](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIEncryptionFrontend.CLIPassphrasePromptCreate]
//   - [DIEncryptionFrontend.CLIPassphrasePromptUnlock]
//   - [DIEncryptionFrontend.CLIVerifyPassphrasePromptCreate]
//   - [DIEncryptionFrontend.GUIAskForPassphraseWithPassphraseUsageError]
//   - [DIEncryptionFrontend.GUIPassphraseLabelCreate]
//   - [DIEncryptionFrontend.GUIPassphraseLabelUnlock]
//   - [DIEncryptionFrontend.GUIPassphrasePromptCreate]
//   - [DIEncryptionFrontend.GUIPassphrasePromptUnlock]
//   - [DIEncryptionFrontend.GUIVerifyPassphraseLabelCreate]
//   - [DIEncryptionFrontend.AddPassphraseEntryWithXpcHandlerFlagsUsageError]
//   - [DIEncryptionFrontend.AllowStoringInKeychain]
//   - [DIEncryptionFrontend.SetAllowStoringInKeychain]
//   - [DIEncryptionFrontend.AskPermissionWithRememberPasswordError]
//   - [DIEncryptionFrontend.CheckAuthEntriesWithHasPassphraseEntryHasPublicKeyEntryError]
//   - [DIEncryptionFrontend.CheckWithHasIcloudKeychainError]
//   - [DIEncryptionFrontend.CheckWithItemRefIsSystemKeychainError]
//   - [DIEncryptionFrontend.ConsoleAskForPassphraseWithUseStdinUsageError]
//   - [DIEncryptionFrontend.DiParams]
//   - [DIEncryptionFrontend.EncodeWithCoder]
//   - [DIEncryptionFrontend.EncryptionUUID]
//   - [DIEncryptionFrontend.Flags]
//   - [DIEncryptionFrontend.GenerateAuthTableWithError]
//   - [DIEncryptionFrontend.GetCertificateWithCertificatePathError]
//   - [DIEncryptionFrontend.GetCertificateWithPublicKeyError]
//   - [DIEncryptionFrontend.GetSerializerWithAuthTable]
//   - [DIEncryptionFrontend.KeychainUnlockWithError]
//   - [DIEncryptionFrontend.KeychainUnlockWithIsSystemKeychainError]
//   - [DIEncryptionFrontend.LookupLegacyKeychainWithXpcHandlerError]
//   - [DIEncryptionFrontend.SetPassphraseError]
//   - [DIEncryptionFrontend.StoreInKeychainWithPassphraseForceSystemKeychainError]
//   - [DIEncryptionFrontend.UnlockUsingPublicKeyWithError]
//   - [DIEncryptionFrontend.UnlockUsingSaksWithError]
//   - [DIEncryptionFrontend.UnlockUsingSymmetricKeyWithError]
//   - [DIEncryptionFrontend.UnlockWithPassphraseError]
//   - [DIEncryptionFrontend.UnlockWithXpcHandlerError]
//   - [DIEncryptionFrontend.UpdateDiskImageParamsWithFrontendError]
//   - [DIEncryptionFrontend.ValidateDeserializationWithError]
//   - [DIEncryptionFrontend.InitWithCoder]
//   - [DIEncryptionFrontend.InitWithParams]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend
type DIEncryptionFrontend struct {
	objectivec.Object
}

// DIEncryptionFrontendFromID constructs a [DIEncryptionFrontend] from an objc.ID.
func DIEncryptionFrontendFromID(id objc.ID) DIEncryptionFrontend {
	return DIEncryptionFrontend{objectivec.Object{ID: id}}
}

// Ensure DIEncryptionFrontend implements IDIEncryptionFrontend.
var _ IDIEncryptionFrontend = DIEncryptionFrontend{}

// An interface definition for the [DIEncryptionFrontend] class.
//
// # Methods
//
//   - [IDIEncryptionFrontend.CLIPassphrasePromptCreate]
//   - [IDIEncryptionFrontend.CLIPassphrasePromptUnlock]
//   - [IDIEncryptionFrontend.CLIVerifyPassphrasePromptCreate]
//   - [IDIEncryptionFrontend.GUIAskForPassphraseWithPassphraseUsageError]
//   - [IDIEncryptionFrontend.GUIPassphraseLabelCreate]
//   - [IDIEncryptionFrontend.GUIPassphraseLabelUnlock]
//   - [IDIEncryptionFrontend.GUIPassphrasePromptCreate]
//   - [IDIEncryptionFrontend.GUIPassphrasePromptUnlock]
//   - [IDIEncryptionFrontend.GUIVerifyPassphraseLabelCreate]
//   - [IDIEncryptionFrontend.AddPassphraseEntryWithXpcHandlerFlagsUsageError]
//   - [IDIEncryptionFrontend.AllowStoringInKeychain]
//   - [IDIEncryptionFrontend.SetAllowStoringInKeychain]
//   - [IDIEncryptionFrontend.AskPermissionWithRememberPasswordError]
//   - [IDIEncryptionFrontend.CheckAuthEntriesWithHasPassphraseEntryHasPublicKeyEntryError]
//   - [IDIEncryptionFrontend.CheckWithHasIcloudKeychainError]
//   - [IDIEncryptionFrontend.CheckWithItemRefIsSystemKeychainError]
//   - [IDIEncryptionFrontend.ConsoleAskForPassphraseWithUseStdinUsageError]
//   - [IDIEncryptionFrontend.DiParams]
//   - [IDIEncryptionFrontend.EncodeWithCoder]
//   - [IDIEncryptionFrontend.EncryptionUUID]
//   - [IDIEncryptionFrontend.Flags]
//   - [IDIEncryptionFrontend.GenerateAuthTableWithError]
//   - [IDIEncryptionFrontend.GetCertificateWithCertificatePathError]
//   - [IDIEncryptionFrontend.GetCertificateWithPublicKeyError]
//   - [IDIEncryptionFrontend.GetSerializerWithAuthTable]
//   - [IDIEncryptionFrontend.KeychainUnlockWithError]
//   - [IDIEncryptionFrontend.KeychainUnlockWithIsSystemKeychainError]
//   - [IDIEncryptionFrontend.LookupLegacyKeychainWithXpcHandlerError]
//   - [IDIEncryptionFrontend.SetPassphraseError]
//   - [IDIEncryptionFrontend.StoreInKeychainWithPassphraseForceSystemKeychainError]
//   - [IDIEncryptionFrontend.UnlockUsingPublicKeyWithError]
//   - [IDIEncryptionFrontend.UnlockUsingSaksWithError]
//   - [IDIEncryptionFrontend.UnlockUsingSymmetricKeyWithError]
//   - [IDIEncryptionFrontend.UnlockWithPassphraseError]
//   - [IDIEncryptionFrontend.UnlockWithXpcHandlerError]
//   - [IDIEncryptionFrontend.UpdateDiskImageParamsWithFrontendError]
//   - [IDIEncryptionFrontend.ValidateDeserializationWithError]
//   - [IDIEncryptionFrontend.InitWithCoder]
//   - [IDIEncryptionFrontend.InitWithParams]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend
type IDIEncryptionFrontend interface {
	objectivec.IObject

	// Topic: Methods

	CLIPassphrasePromptCreate() string
	CLIPassphrasePromptUnlock() string
	CLIVerifyPassphrasePromptCreate() string
	GUIAskForPassphraseWithPassphraseUsageError(usage int64) (bool, error)
	GUIPassphraseLabelCreate() string
	GUIPassphraseLabelUnlock() string
	GUIPassphrasePromptCreate() string
	GUIPassphrasePromptUnlock() string
	GUIVerifyPassphraseLabelCreate() string
	AddPassphraseEntryWithXpcHandlerFlagsUsageError(handler objectivec.IObject, flags uint64, usage int64) (bool, error)
	AllowStoringInKeychain() bool
	SetAllowStoringInKeychain(value bool)
	AskPermissionWithRememberPasswordError() (bool, error)
	CheckAuthEntriesWithHasPassphraseEntryHasPublicKeyEntryError() (bool, bool, error)
	CheckWithHasIcloudKeychainError() (bool, error)
	CheckWithItemRefIsSystemKeychainError(ref objectivec.IObject) (bool, error)
	ConsoleAskForPassphraseWithUseStdinUsageError(stdin bool, usage int64) (bool, error)
	DiParams() IDIBaseParams
	EncodeWithCoder(coder foundation.INSCoder)
	EncryptionUUID() foundation.NSUUID
	Flags() uint64
	GenerateAuthTableWithError() (unsafe.Pointer, error)
	GetCertificateWithCertificatePathError(path objectivec.IObject) (objectivec.IObject, error)
	GetCertificateWithPublicKeyError(key objectivec.IObject) (objectivec.IObject, error)
	GetSerializerWithAuthTable(table unsafe.Pointer) objectivec.IObject
	KeychainUnlockWithError() (bool, error)
	KeychainUnlockWithIsSystemKeychainError(keychain bool) (bool, error)
	LookupLegacyKeychainWithXpcHandlerError(handler objectivec.IObject) (bool, error)
	SetPassphraseError(passphrase string) (bool, error)
	StoreInKeychainWithPassphraseForceSystemKeychainError(passphrase objectivec.IObject, keychain bool) (bool, error)
	UnlockUsingPublicKeyWithError() (bool, error)
	UnlockUsingSaksWithError() (bool, error)
	UnlockUsingSymmetricKeyWithError() (bool, error)
	UnlockWithPassphraseError(passphrase string) (bool, error)
	UnlockWithXpcHandlerError(handler objectivec.IObject) (bool, error)
	UpdateDiskImageParamsWithFrontendError(frontend objectivec.IObject) (bool, error)
	ValidateDeserializationWithError() (bool, error)
	InitWithCoder(coder foundation.INSCoder) DIEncryptionFrontend
	InitWithParams(params objectivec.IObject) DIEncryptionFrontend
}

// Init initializes the instance.
func (d DIEncryptionFrontend) Init() DIEncryptionFrontend {
	rv := objc.Send[DIEncryptionFrontend](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIEncryptionFrontend) Autorelease() DIEncryptionFrontend {
	rv := objc.Send[DIEncryptionFrontend](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIEncryptionFrontend creates a new DIEncryptionFrontend instance.
func NewDIEncryptionFrontend() DIEncryptionFrontend {
	class := getDIEncryptionFrontendClass()
	rv := objc.Send[DIEncryptionFrontend](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithCoder:
func NewDIEncryptionFrontendWithCoder(coder objectivec.IObject) DIEncryptionFrontend {
	instance := getDIEncryptionFrontendClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIEncryptionFrontendFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithParams:
func NewDIEncryptionFrontendWithParams(params objectivec.IObject) DIEncryptionFrontend {
	instance := getDIEncryptionFrontendClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return DIEncryptionFrontendFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/GUIAskForPassphraseWithPassphraseUsage:error:
func (d DIEncryptionFrontend) GUIAskForPassphraseWithPassphraseUsageError(usage int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("GUIAskForPassphraseWithPassphraseUsage:error:"), usage, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("GUIAskForPassphraseWithPassphraseUsage:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/addPassphraseEntryWithXpcHandler:flags:usage:error:
func (d DIEncryptionFrontend) AddPassphraseEntryWithXpcHandlerFlagsUsageError(handler objectivec.IObject, flags uint64, usage int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("addPassphraseEntryWithXpcHandler:flags:usage:error:"), handler, flags, usage, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addPassphraseEntryWithXpcHandler:flags:usage:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/askPermissionWithRememberPassword:error:
func (d DIEncryptionFrontend) AskPermissionWithRememberPasswordError() (bool, error) {
	var password bool
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("askPermissionWithRememberPassword:error:"), unsafe.Pointer(&password), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("askPermissionWithRememberPassword:error: returned NO with nil NSError")
	}
	return password, nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/checkAuthEntriesWithHasPassphraseEntry:hasPublicKeyEntry:error:
func (d DIEncryptionFrontend) CheckAuthEntriesWithHasPassphraseEntryHasPublicKeyEntryError() (bool, bool, error) {
	var entry bool
	var entry2 bool
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("checkAuthEntriesWithHasPassphraseEntry:hasPublicKeyEntry:error:"), unsafe.Pointer(&entry), unsafe.Pointer(&entry2), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, false, errors.New("checkAuthEntriesWithHasPassphraseEntry:hasPublicKeyEntry:error: returned NO with nil NSError")
	}
	return entry, entry2, nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/checkWithHasIcloudKeychain:error:
func (d DIEncryptionFrontend) CheckWithHasIcloudKeychainError() (bool, error) {
	var keychain bool
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("checkWithHasIcloudKeychain:error:"), unsafe.Pointer(&keychain), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkWithHasIcloudKeychain:error: returned NO with nil NSError")
	}
	return keychain, nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/checkWithItemRef:isSystemKeychain:error:
func (d DIEncryptionFrontend) CheckWithItemRefIsSystemKeychainError(ref objectivec.IObject) (bool, error) {
	var keychain bool
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("checkWithItemRef:isSystemKeychain:error:"), ref, unsafe.Pointer(&keychain), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkWithItemRef:isSystemKeychain:error: returned NO with nil NSError")
	}
	return keychain, nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/consoleAskForPassphraseWithUseStdin:usage:error:
func (d DIEncryptionFrontend) ConsoleAskForPassphraseWithUseStdinUsageError(stdin bool, usage int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("consoleAskForPassphraseWithUseStdin:usage:error:"), stdin, usage, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("consoleAskForPassphraseWithUseStdin:usage:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/encodeWithCoder:
func (d DIEncryptionFrontend) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/generateAuthTableWithError:
func (d DIEncryptionFrontend) GenerateAuthTableWithError() (unsafe.Pointer, error) {
	var errorPtr objc.ID
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("generateAuthTableWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/getCertificateWithCertificatePath:error:
func (d DIEncryptionFrontend) GetCertificateWithCertificatePathError(path objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getCertificateWithCertificatePath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/getCertificateWithPublicKey:error:
func (d DIEncryptionFrontend) GetCertificateWithPublicKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getCertificateWithPublicKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/getSerializerWithAuthTable:
func (d DIEncryptionFrontend) GetSerializerWithAuthTable(table unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("getSerializerWithAuthTable:"), table)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/keychainUnlockWithError:
func (d DIEncryptionFrontend) KeychainUnlockWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("keychainUnlockWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("keychainUnlockWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/keychainUnlockWithIsSystemKeychain:error:
func (d DIEncryptionFrontend) KeychainUnlockWithIsSystemKeychainError(keychain bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("keychainUnlockWithIsSystemKeychain:error:"), keychain, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("keychainUnlockWithIsSystemKeychain:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/lookupLegacyKeychainWithXpcHandler:error:
func (d DIEncryptionFrontend) LookupLegacyKeychainWithXpcHandlerError(handler objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("lookupLegacyKeychainWithXpcHandler:error:"), handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("lookupLegacyKeychainWithXpcHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/setPassphrase:error:
func (d DIEncryptionFrontend) SetPassphraseError(passphrase string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setPassphrase:error:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setPassphrase:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/storeInKeychainWithPassphrase:forceSystemKeychain:error:
func (d DIEncryptionFrontend) StoreInKeychainWithPassphraseForceSystemKeychainError(passphrase objectivec.IObject, keychain bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("storeInKeychainWithPassphrase:forceSystemKeychain:error:"), passphrase, keychain, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("storeInKeychainWithPassphrase:forceSystemKeychain:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/unlockUsingPublicKeyWithError:
func (d DIEncryptionFrontend) UnlockUsingPublicKeyWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unlockUsingPublicKeyWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unlockUsingPublicKeyWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/unlockUsingSaksWithError:
func (d DIEncryptionFrontend) UnlockUsingSaksWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unlockUsingSaksWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unlockUsingSaksWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/unlockUsingSymmetricKeyWithError:
func (d DIEncryptionFrontend) UnlockUsingSymmetricKeyWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unlockUsingSymmetricKeyWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unlockUsingSymmetricKeyWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/unlockWithPassphrase:error:
func (d DIEncryptionFrontend) UnlockWithPassphraseError(passphrase string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unlockWithPassphrase:error:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unlockWithPassphrase:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/unlockWithXpcHandler:error:
func (d DIEncryptionFrontend) UnlockWithXpcHandlerError(handler objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("unlockWithXpcHandler:error:"), handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unlockWithXpcHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/updateDiskImageParamsWithFrontend:error:
func (d DIEncryptionFrontend) UpdateDiskImageParamsWithFrontendError(frontend objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("updateDiskImageParamsWithFrontend:error:"), frontend, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateDiskImageParamsWithFrontend:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/validateDeserializationWithError:
func (d DIEncryptionFrontend) ValidateDeserializationWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("validateDeserializationWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateDeserializationWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithCoder:
func (d DIEncryptionFrontend) InitWithCoder(coder foundation.INSCoder) DIEncryptionFrontend {
	rv := objc.Send[DIEncryptionFrontend](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/initWithParams:
func (d DIEncryptionFrontend) InitWithParams(params objectivec.IObject) DIEncryptionFrontend {
	rv := objc.Send[DIEncryptionFrontend](d.ID, objc.Sel("initWithParams:"), params)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/hasGUIaccess
func (_DIEncryptionFrontendClass DIEncryptionFrontendClass) HasGUIaccess() bool {
	rv := objc.Send[bool](objc.ID(_DIEncryptionFrontendClass.class), objc.Sel("hasGUIaccess"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/supportsSecureCoding
func (_DIEncryptionFrontendClass DIEncryptionFrontendClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIEncryptionFrontendClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/updateSystemKeychainAttrWithDict:isStoring:error:
func (_DIEncryptionFrontendClass DIEncryptionFrontendClass) UpdateSystemKeychainAttrWithDictIsStoringError(dict objectivec.IObject, storing bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIEncryptionFrontendClass.class), objc.Sel("updateSystemKeychainAttrWithDict:isStoring:error:"), dict, storing, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateSystemKeychainAttrWithDict:isStoring:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/CLIPassphrasePromptCreate
func (d DIEncryptionFrontend) CLIPassphrasePromptCreate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("CLIPassphrasePromptCreate"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/CLIPassphrasePromptUnlock
func (d DIEncryptionFrontend) CLIPassphrasePromptUnlock() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("CLIPassphrasePromptUnlock"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/CLIVerifyPassphrasePromptCreate
func (d DIEncryptionFrontend) CLIVerifyPassphrasePromptCreate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("CLIVerifyPassphrasePromptCreate"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/GUIPassphraseLabelCreate
func (d DIEncryptionFrontend) GUIPassphraseLabelCreate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("GUIPassphraseLabelCreate"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/GUIPassphraseLabelUnlock
func (d DIEncryptionFrontend) GUIPassphraseLabelUnlock() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("GUIPassphraseLabelUnlock"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/GUIPassphrasePromptCreate
func (d DIEncryptionFrontend) GUIPassphrasePromptCreate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("GUIPassphrasePromptCreate"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/GUIPassphrasePromptUnlock
func (d DIEncryptionFrontend) GUIPassphrasePromptUnlock() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("GUIPassphrasePromptUnlock"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/GUIVerifyPassphraseLabelCreate
func (d DIEncryptionFrontend) GUIVerifyPassphraseLabelCreate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("GUIVerifyPassphraseLabelCreate"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/allowStoringInKeychain
func (d DIEncryptionFrontend) AllowStoringInKeychain() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowStoringInKeychain"))
	return rv
}
func (d DIEncryptionFrontend) SetAllowStoringInKeychain(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setAllowStoringInKeychain:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/diParams
func (d DIEncryptionFrontend) DiParams() IDIBaseParams {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("diParams"))
	return DIBaseParamsFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/encryptionUUID
func (d DIEncryptionFrontend) EncryptionUUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("encryptionUUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionFrontend/flags
func (d DIEncryptionFrontend) Flags() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("flags"))
	return rv
}
