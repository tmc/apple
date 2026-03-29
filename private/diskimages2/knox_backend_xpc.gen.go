// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [KNOXBackendXPC] class.
var (
	_KNOXBackendXPCClass     KNOXBackendXPCClass
	_KNOXBackendXPCClassOnce sync.Once
)

func getKNOXBackendXPCClass() KNOXBackendXPCClass {
	_KNOXBackendXPCClassOnce.Do(func() {
		_KNOXBackendXPCClass = KNOXBackendXPCClass{class: objc.GetClass("KNOXBackendXPC")}
	})
	return _KNOXBackendXPCClass
}

// GetKNOXBackendXPCClass returns the class object for KNOXBackendXPC.
func GetKNOXBackendXPCClass() KNOXBackendXPCClass {
	return getKNOXBackendXPCClass()
}

type KNOXBackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (kc KNOXBackendXPCClass) Class() objc.Class {
	return kc.class
}

// Alloc allocates memory for a new instance of the class.
func (kc KNOXBackendXPCClass) Alloc() KNOXBackendXPC {
	rv := objc.Send[KNOXBackendXPC](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [KNOXBackendXPC.URL]
//   - [KNOXBackendXPC.SetURL]
//   - [KNOXBackendXPC.Key]
//   - [KNOXBackendXPC.InitWithURLKey]
// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC
type KNOXBackendXPC struct {
	BackendXPC
}

// KNOXBackendXPCFromID constructs a [KNOXBackendXPC] from an objc.ID.
func KNOXBackendXPCFromID(id objc.ID) KNOXBackendXPC {
	return KNOXBackendXPC{BackendXPC: BackendXPCFromID(id)}
}
// Ensure KNOXBackendXPC implements IKNOXBackendXPC.
var _ IKNOXBackendXPC = KNOXBackendXPC{}

// An interface definition for the [KNOXBackendXPC] class.
//
// # Methods
//
//   - [IKNOXBackendXPC.URL]
//   - [IKNOXBackendXPC.SetURL]
//   - [IKNOXBackendXPC.Key]
//   - [IKNOXBackendXPC.InitWithURLKey]
//
// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC
type IKNOXBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	URL() IDIURL
	SetURL(value IDIURL)
	Key() objectivec.IObject
	InitWithURLKey(url foundation.INSURL, key unsafe.Pointer) KNOXBackendXPC
}

// Init initializes the instance.
func (k KNOXBackendXPC) Init() KNOXBackendXPC {
	rv := objc.Send[KNOXBackendXPC](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k KNOXBackendXPC) Autorelease() KNOXBackendXPC {
	rv := objc.Send[KNOXBackendXPC](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewKNOXBackendXPC creates a new KNOXBackendXPC instance.
func NewKNOXBackendXPC() KNOXBackendXPC {
	class := getKNOXBackendXPCClass()
	rv := objc.Send[KNOXBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC/initWithCoder:
func NewKNOXBackendXPCWithCoder(coder objectivec.IObject) KNOXBackendXPC {
	instance := getKNOXBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return KNOXBackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC/initWithURL:key:
func NewKNOXBackendXPCWithURLKey(url foundation.INSURL, key unsafe.Pointer) KNOXBackendXPC {
	instance := getKNOXBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:key:"), url, key)
	return KNOXBackendXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC/initWithURL:key:
func (k KNOXBackendXPC) InitWithURLKey(url foundation.INSURL, key unsafe.Pointer) KNOXBackendXPC {
	rv := objc.Send[KNOXBackendXPC](k.ID, objc.Sel("initWithURL:key:"), url, key)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC/URL
func (k KNOXBackendXPC) URL() IDIURL {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("URL"))
	return DIURLFromID(objc.ID(rv))
}
func (k KNOXBackendXPC) SetURL(value IDIURL) {
	objc.Send[struct{}](k.ID, objc.Sel("setURL:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/KNOXBackendXPC/key
func (k KNOXBackendXPC) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}

