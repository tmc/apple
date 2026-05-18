// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKShareRequestAccessOperation] class.
var (
	_CKShareRequestAccessOperationClass     CKShareRequestAccessOperationClass
	_CKShareRequestAccessOperationClassOnce sync.Once
)

func getCKShareRequestAccessOperationClass() CKShareRequestAccessOperationClass {
	_CKShareRequestAccessOperationClassOnce.Do(func() {
		_CKShareRequestAccessOperationClass = CKShareRequestAccessOperationClass{class: objc.GetClass("CKShareRequestAccessOperation")}
	})
	return _CKShareRequestAccessOperationClass
}

// GetCKShareRequestAccessOperationClass returns the class object for CKShareRequestAccessOperation.
func GetCKShareRequestAccessOperationClass() CKShareRequestAccessOperationClass {
	return getCKShareRequestAccessOperationClass()
}

type CKShareRequestAccessOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKShareRequestAccessOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKShareRequestAccessOperationClass) Alloc() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [CKShareRequestAccessOperation.InitWithShareURLs]: Creates a share request access operation configured with specified share URLs.
//
// # Instance Properties
//
//   - [CKShareRequestAccessOperation.PerShareAccessRequestResultBlock]: The closure to execute when CloudKit processes a share access request.
//   - [CKShareRequestAccessOperation.SetPerShareAccessRequestResultBlock]
//   - [CKShareRequestAccessOperation.ShareAccessRequestResultBlock]: The closure to execute after CloudKit processes each share access request.
//   - [CKShareRequestAccessOperation.SetShareAccessRequestResultBlock]
//   - [CKShareRequestAccessOperation.ShareURLs]: The URLs of the shares to request access to.
//   - [CKShareRequestAccessOperation.SetShareURLs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation
type CKShareRequestAccessOperation struct {
	CKOperation
}

// CKShareRequestAccessOperationFromID constructs a [CKShareRequestAccessOperation] from an objc.ID.
func CKShareRequestAccessOperationFromID(id objc.ID) CKShareRequestAccessOperation {
	return CKShareRequestAccessOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKShareRequestAccessOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKShareRequestAccessOperation] class.
//
// # Initializers
//
//   - [ICKShareRequestAccessOperation.InitWithShareURLs]: Creates a share request access operation configured with specified share URLs.
//
// # Instance Properties
//
//   - [ICKShareRequestAccessOperation.PerShareAccessRequestResultBlock]: The closure to execute when CloudKit processes a share access request.
//   - [ICKShareRequestAccessOperation.SetPerShareAccessRequestResultBlock]
//   - [ICKShareRequestAccessOperation.ShareAccessRequestResultBlock]: The closure to execute after CloudKit processes each share access request.
//   - [ICKShareRequestAccessOperation.SetShareAccessRequestResultBlock]
//   - [ICKShareRequestAccessOperation.ShareURLs]: The URLs of the shares to request access to.
//   - [ICKShareRequestAccessOperation.SetShareURLs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation
type ICKShareRequestAccessOperation interface {
	ICKOperation

	// Topic: Initializers

	// Creates a share request access operation configured with specified share URLs.
	InitWithShareURLs(shareURLs []foundation.NSURL) CKShareRequestAccessOperation

	// Topic: Instance Properties

	// The closure to execute when CloudKit processes a share access request.
	PerShareAccessRequestResultBlock() unsafe.Pointer
	SetPerShareAccessRequestResultBlock(value unsafe.Pointer)
	// The closure to execute after CloudKit processes each share access request.
	ShareAccessRequestResultBlock() unsafe.Pointer
	SetShareAccessRequestResultBlock(value unsafe.Pointer)
	// The URLs of the shares to request access to.
	ShareURLs() []foundation.NSURL
	SetShareURLs(value []foundation.NSURL)
}

// Init initializes the instance.
func (c CKShareRequestAccessOperation) Init() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKShareRequestAccessOperation) Autorelease() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareRequestAccessOperation creates a new CKShareRequestAccessOperation instance.
func NewCKShareRequestAccessOperation() CKShareRequestAccessOperation {
	class := getCKShareRequestAccessOperationClass()
	rv := objc.Send[CKShareRequestAccessOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a share request access operation configured with specified share
// URLs.
//
// shareURLs: An array of [NSURL] objects representing the shares to request access to.
//
// # Return Value
//
// A configured [CKShareRequestAccessOperation] instance.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/init(shareURLs:)
func NewCKShareRequestAccessOperationWithShareURLs(shareURLs []foundation.NSURL) CKShareRequestAccessOperation {
	instance := getCKShareRequestAccessOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithShareURLs:"), objectivec.IObjectSliceToNSArray(shareURLs))
	return CKShareRequestAccessOperationFromID(rv)
}

// Creates a share request access operation configured with specified share
// URLs.
//
// shareURLs: An array of [NSURL] objects representing the shares to request access to.
//
// # Return Value
//
// A configured [CKShareRequestAccessOperation] instance.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/init(shareURLs:)
func (c CKShareRequestAccessOperation) InitWithShareURLs(shareURLs []foundation.NSURL) CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c.ID, objc.Sel("initWithShareURLs:"), objectivec.IObjectSliceToNSArray(shareURLs))
	return rv
}

// The closure to execute when CloudKit processes a share access request.
//
// See: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/pershareaccessrequestresultblock
func (c CKShareRequestAccessOperation) PerShareAccessRequestResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perShareAccessRequestResultBlock"))
	return rv
}
func (c CKShareRequestAccessOperation) SetPerShareAccessRequestResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerShareAccessRequestResultBlock:"), value)
}

// The closure to execute after CloudKit processes each share access request.
//
// See: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/shareaccessrequestresultblock
func (c CKShareRequestAccessOperation) ShareAccessRequestResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("shareAccessRequestResultBlock"))
	return rv
}
func (c CKShareRequestAccessOperation) SetShareAccessRequestResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setShareAccessRequestResultBlock:"), value)
}

// The URLs of the shares to request access to.
//
// # Discussion
//
// Include multiple URLs to request access to multiple shares simultaneously.
// The server processes each URL independently.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareURLs
func (c CKShareRequestAccessOperation) ShareURLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shareURLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}
func (c CKShareRequestAccessOperation) SetShareURLs(value []foundation.NSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setShareURLs:"), objectivec.IObjectSliceToNSArray(value))
}
