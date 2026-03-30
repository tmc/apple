// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HTTPBackendXPC] class.
var (
	_HTTPBackendXPCClass     HTTPBackendXPCClass
	_HTTPBackendXPCClassOnce sync.Once
)

func getHTTPBackendXPCClass() HTTPBackendXPCClass {
	_HTTPBackendXPCClassOnce.Do(func() {
		_HTTPBackendXPCClass = HTTPBackendXPCClass{class: objc.GetClass("HTTPBackendXPC")}
	})
	return _HTTPBackendXPCClass
}

// GetHTTPBackendXPCClass returns the class object for HTTPBackendXPC.
func GetHTTPBackendXPCClass() HTTPBackendXPCClass {
	return getHTTPBackendXPCClass()
}

type HTTPBackendXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (hc HTTPBackendXPCClass) Class() objc.Class {
	return hc.class
}

// Alloc allocates memory for a new instance of the class.
func (hc HTTPBackendXPCClass) Alloc() HTTPBackendXPC {
	rv := objc.Send[HTTPBackendXPC](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [HTTPBackendXPC.URL]
//   - [HTTPBackendXPC.SetURL]
//   - [HTTPBackendXPC.AddToURLWithUserPassword]
//   - [HTTPBackendXPC.AskForPasswordUsingGUI]
//   - [HTTPBackendXPC.AskForPasswordUsingTTY]
//   - [HTTPBackendXPC.CreateBackend]
//   - [HTTPBackendXPC.CreateBackendWithSize]
//   - [HTTPBackendXPC.LookupPasswordInKeychainWithRealm]
//   - [HTTPBackendXPC.InitWithURL]
//
// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC
type HTTPBackendXPC struct {
	BackendXPC
}

// HTTPBackendXPCFromID constructs a [HTTPBackendXPC] from an objc.ID.
func HTTPBackendXPCFromID(id objc.ID) HTTPBackendXPC {
	return HTTPBackendXPC{BackendXPC: BackendXPCFromID(id)}
}

// Ensure HTTPBackendXPC implements IHTTPBackendXPC.
var _ IHTTPBackendXPC = HTTPBackendXPC{}

// An interface definition for the [HTTPBackendXPC] class.
//
// # Methods
//
//   - [IHTTPBackendXPC.URL]
//   - [IHTTPBackendXPC.SetURL]
//   - [IHTTPBackendXPC.AddToURLWithUserPassword]
//   - [IHTTPBackendXPC.AskForPasswordUsingGUI]
//   - [IHTTPBackendXPC.AskForPasswordUsingTTY]
//   - [IHTTPBackendXPC.CreateBackend]
//   - [IHTTPBackendXPC.CreateBackendWithSize]
//   - [IHTTPBackendXPC.LookupPasswordInKeychainWithRealm]
//   - [IHTTPBackendXPC.InitWithURL]
//
// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC
type IHTTPBackendXPC interface {
	IBackendXPC

	// Topic: Methods

	URL() IDIURL
	SetURL(value IDIURL)
	AddToURLWithUserPassword(user objectivec.IObject, password objectivec.IObject)
	AskForPasswordUsingGUI() bool
	AskForPasswordUsingTTY() bool
	CreateBackend()
	CreateBackendWithSize(size uint64)
	LookupPasswordInKeychainWithRealm(realm objectivec.IObject) bool
	InitWithURL(url foundation.INSURL) HTTPBackendXPC
}

// Init initializes the instance.
func (h HTTPBackendXPC) Init() HTTPBackendXPC {
	rv := objc.Send[HTTPBackendXPC](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HTTPBackendXPC) Autorelease() HTTPBackendXPC {
	rv := objc.Send[HTTPBackendXPC](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPBackendXPC creates a new HTTPBackendXPC instance.
func NewHTTPBackendXPC() HTTPBackendXPC {
	class := getHTTPBackendXPCClass()
	rv := objc.Send[HTTPBackendXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/initWithCoder:
func NewHTTPBackendXPCWithCoder(coder objectivec.IObject) HTTPBackendXPC {
	instance := getHTTPBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return HTTPBackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/initWithURL:
func NewHTTPBackendXPCWithURL(url foundation.INSURL) HTTPBackendXPC {
	instance := getHTTPBackendXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return HTTPBackendXPCFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/addToURLWithUser:password:
func (h HTTPBackendXPC) AddToURLWithUserPassword(user objectivec.IObject, password objectivec.IObject) {
	objc.Send[objc.ID](h.ID, objc.Sel("addToURLWithUser:password:"), user, password)
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/askForPasswordUsingGUI
func (h HTTPBackendXPC) AskForPasswordUsingGUI() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("askForPasswordUsingGUI"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/askForPasswordUsingTTY
func (h HTTPBackendXPC) AskForPasswordUsingTTY() bool {
	rv := objc.Send[bool](h.ID, objc.Sel("askForPasswordUsingTTY"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/createBackend
func (h HTTPBackendXPC) CreateBackend() {
	objc.Send[objc.ID](h.ID, objc.Sel("createBackend"))
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/createBackendWithSize:
func (h HTTPBackendXPC) CreateBackendWithSize(size uint64) {
	objc.Send[objc.ID](h.ID, objc.Sel("createBackendWithSize:"), size)
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/lookupPasswordInKeychainWithRealm:
func (h HTTPBackendXPC) LookupPasswordInKeychainWithRealm(realm objectivec.IObject) bool {
	rv := objc.Send[bool](h.ID, objc.Sel("lookupPasswordInKeychainWithRealm:"), realm)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/initWithURL:
func (h HTTPBackendXPC) InitWithURL(url foundation.INSURL) HTTPBackendXPC {
	rv := objc.Send[HTTPBackendXPC](h.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/HTTPBackendXPC/URL
func (h HTTPBackendXPC) URL() IDIURL {
	rv := objc.Send[objc.ID](h.ID, objc.Sel("URL"))
	return DIURLFromID(objc.ID(rv))
}
func (h HTTPBackendXPC) SetURL(value IDIURL) {
	objc.Send[struct{}](h.ID, objc.Sel("setURL:"), value)
}
