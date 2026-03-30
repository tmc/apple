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

// The class instance for the [DIEncryptionChpass] class.
var (
	_DIEncryptionChpassClass     DIEncryptionChpassClass
	_DIEncryptionChpassClassOnce sync.Once
)

func getDIEncryptionChpassClass() DIEncryptionChpassClass {
	_DIEncryptionChpassClassOnce.Do(func() {
		_DIEncryptionChpassClass = DIEncryptionChpassClass{class: objc.GetClass("DIEncryptionChpass")}
	})
	return _DIEncryptionChpassClass
}

// GetDIEncryptionChpassClass returns the class object for DIEncryptionChpass.
func GetDIEncryptionChpassClass() DIEncryptionChpassClass {
	return getDIEncryptionChpassClass()
}

type DIEncryptionChpassClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIEncryptionChpassClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIEncryptionChpassClass) Alloc() DIEncryptionChpass {
	rv := objc.Send[DIEncryptionChpass](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIEncryptionChpass.PassEntryToChange]
//   - [DIEncryptionChpass.SetPassEntryToChange]
//   - [DIEncryptionChpass.ReplacePassWithXpcHandlerParamsError]
//   - [DIEncryptionChpass.ReplacePassphraseError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass
type DIEncryptionChpass struct {
	DIEncryptionFrontend
}

// DIEncryptionChpassFromID constructs a [DIEncryptionChpass] from an objc.ID.
func DIEncryptionChpassFromID(id objc.ID) DIEncryptionChpass {
	return DIEncryptionChpass{DIEncryptionFrontend: DIEncryptionFrontendFromID(id)}
}

// Ensure DIEncryptionChpass implements IDIEncryptionChpass.
var _ IDIEncryptionChpass = DIEncryptionChpass{}

// An interface definition for the [DIEncryptionChpass] class.
//
// # Methods
//
//   - [IDIEncryptionChpass.PassEntryToChange]
//   - [IDIEncryptionChpass.SetPassEntryToChange]
//   - [IDIEncryptionChpass.ReplacePassWithXpcHandlerParamsError]
//   - [IDIEncryptionChpass.ReplacePassphraseError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass
type IDIEncryptionChpass interface {
	IDIEncryptionFrontend

	// Topic: Methods

	PassEntryToChange() unsafe.Pointer
	SetPassEntryToChange(value unsafe.Pointer)
	ReplacePassWithXpcHandlerParamsError(handler objectivec.IObject, params objectivec.IObject) (bool, error)
	ReplacePassphraseError(passphrase string) (bool, error)
}

// Init initializes the instance.
func (d DIEncryptionChpass) Init() DIEncryptionChpass {
	rv := objc.Send[DIEncryptionChpass](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIEncryptionChpass) Autorelease() DIEncryptionChpass {
	rv := objc.Send[DIEncryptionChpass](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIEncryptionChpass creates a new DIEncryptionChpass instance.
func NewDIEncryptionChpass() DIEncryptionChpass {
	class := getDIEncryptionChpassClass()
	rv := objc.Send[DIEncryptionChpass](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass/initWithCoder:
func NewDIEncryptionChpassWithCoder(coder objectivec.IObject) DIEncryptionChpass {
	instance := getDIEncryptionChpassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIEncryptionChpassFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass/initWithParams:
func NewDIEncryptionChpassWithParams(params objectivec.IObject) DIEncryptionChpass {
	instance := getDIEncryptionChpassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParams:"), params)
	return DIEncryptionChpassFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass/replacePassWithXpcHandler:params:error:
func (d DIEncryptionChpass) ReplacePassWithXpcHandlerParamsError(handler objectivec.IObject, params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("replacePassWithXpcHandler:params:error:"), handler, params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("replacePassWithXpcHandler:params:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass/replacePassphrase:error:
func (d DIEncryptionChpass) ReplacePassphraseError(passphrase string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("replacePassphrase:error:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("replacePassphrase:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIEncryptionChpass/passEntryToChange
func (d DIEncryptionChpass) PassEntryToChange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d.ID, objc.Sel("passEntryToChange"))
	return rv
}
func (d DIEncryptionChpass) SetPassEntryToChange(value unsafe.Pointer) {
	objc.Send[struct{}](d.ID, objc.Sel("setPassEntryToChange:"), value)
}
