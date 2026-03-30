// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DITemporaryPassphrase] class.
var (
	_DITemporaryPassphraseClass     DITemporaryPassphraseClass
	_DITemporaryPassphraseClassOnce sync.Once
)

func getDITemporaryPassphraseClass() DITemporaryPassphraseClass {
	_DITemporaryPassphraseClassOnce.Do(func() {
		_DITemporaryPassphraseClass = DITemporaryPassphraseClass{class: objc.GetClass("DITemporaryPassphrase")}
	})
	return _DITemporaryPassphraseClass
}

// GetDITemporaryPassphraseClass returns the class object for DITemporaryPassphrase.
func GetDITemporaryPassphraseClass() DITemporaryPassphraseClass {
	return getDITemporaryPassphraseClass()
}

type DITemporaryPassphraseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DITemporaryPassphraseClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DITemporaryPassphraseClass) Alloc() DITemporaryPassphrase {
	rv := objc.Send[DITemporaryPassphrase](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DITemporaryPassphrase.Buf]
//   - [DITemporaryPassphrase.SetBuf]
//   - [DITemporaryPassphrase.InitWithPassphrase]
//
// See: https://developer.apple.com/documentation/DiskImages2/DITemporaryPassphrase
type DITemporaryPassphrase struct {
	objectivec.Object
}

// DITemporaryPassphraseFromID constructs a [DITemporaryPassphrase] from an objc.ID.
func DITemporaryPassphraseFromID(id objc.ID) DITemporaryPassphrase {
	return DITemporaryPassphrase{objectivec.Object{ID: id}}
}

// Ensure DITemporaryPassphrase implements IDITemporaryPassphrase.
var _ IDITemporaryPassphrase = DITemporaryPassphrase{}

// An interface definition for the [DITemporaryPassphrase] class.
//
// # Methods
//
//   - [IDITemporaryPassphrase.Buf]
//   - [IDITemporaryPassphrase.SetBuf]
//   - [IDITemporaryPassphrase.InitWithPassphrase]
//
// See: https://developer.apple.com/documentation/DiskImages2/DITemporaryPassphrase
type IDITemporaryPassphrase interface {
	objectivec.IObject

	// Topic: Methods

	Buf() string
	SetBuf(value string)
	InitWithPassphrase(passphrase string) DITemporaryPassphrase
}

// Init initializes the instance.
func (d DITemporaryPassphrase) Init() DITemporaryPassphrase {
	rv := objc.Send[DITemporaryPassphrase](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DITemporaryPassphrase) Autorelease() DITemporaryPassphrase {
	rv := objc.Send[DITemporaryPassphrase](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDITemporaryPassphrase creates a new DITemporaryPassphrase instance.
func NewDITemporaryPassphrase() DITemporaryPassphrase {
	class := getDITemporaryPassphraseClass()
	rv := objc.Send[DITemporaryPassphrase](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DITemporaryPassphrase/initWithPassphrase:
func NewDITemporaryPassphraseWithPassphrase(passphrase string) DITemporaryPassphrase {
	instance := getDITemporaryPassphraseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPassphrase:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")))
	return DITemporaryPassphraseFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DITemporaryPassphrase/initWithPassphrase:
func (d DITemporaryPassphrase) InitWithPassphrase(passphrase string) DITemporaryPassphrase {
	rv := objc.Send[DITemporaryPassphrase](d.ID, objc.Sel("initWithPassphrase:"), unsafe.Pointer(unsafe.StringData(passphrase+"\x00")))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DITemporaryPassphrase/buf
func (d DITemporaryPassphrase) Buf() string {
	rv := objc.Send[*byte](d.ID, objc.Sel("buf"))
	return objc.GoString(rv)
}
func (d DITemporaryPassphrase) SetBuf(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setBuf:"), objc.String(value))
}
