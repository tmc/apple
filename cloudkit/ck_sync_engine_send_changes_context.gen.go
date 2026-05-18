// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineSendChangesContext] class.
var (
	_CKSyncEngineSendChangesContextClass     CKSyncEngineSendChangesContextClass
	_CKSyncEngineSendChangesContextClassOnce sync.Once
)

func getCKSyncEngineSendChangesContextClass() CKSyncEngineSendChangesContextClass {
	_CKSyncEngineSendChangesContextClassOnce.Do(func() {
		_CKSyncEngineSendChangesContextClass = CKSyncEngineSendChangesContextClass{class: objc.GetClass("CKSyncEngineSendChangesContext")}
	})
	return _CKSyncEngineSendChangesContextClass
}

// GetCKSyncEngineSendChangesContextClass returns the class object for CKSyncEngineSendChangesContext.
func GetCKSyncEngineSendChangesContextClass() CKSyncEngineSendChangesContextClass {
	return getCKSyncEngineSendChangesContextClass()
}

type CKSyncEngineSendChangesContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineSendChangesContextClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineSendChangesContextClass) Alloc() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The context of an attempt to send changes to the server.
//
// # Overview
//
// A sync engine has two ways to send changes to iCloud — periodically, in
// cooperation with the system scheduler, and manually, whenever your app
// invokes the [SendChangesWithCompletionHandler] method. This object provides
// information about a single attempt to send changes that includes both the
// reason for the attempt and any additional options in use by the attempt.
//
// # Accessing specific attributes
//
//   - [CKSyncEngineSendChangesContext.Reason]: The reason for the send operation.
//   - [CKSyncEngineSendChangesContext.Options]: The additional options for the send operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext
type CKSyncEngineSendChangesContext struct {
	objectivec.Object
}

// CKSyncEngineSendChangesContextFromID constructs a [CKSyncEngineSendChangesContext] from an objc.ID.
//
// The context of an attempt to send changes to the server.
func CKSyncEngineSendChangesContextFromID(id objc.ID) CKSyncEngineSendChangesContext {
	return CKSyncEngineSendChangesContext{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineSendChangesContext implements ICKSyncEngineSendChangesContext.
var _ ICKSyncEngineSendChangesContext = CKSyncEngineSendChangesContext{}

// An interface definition for the [CKSyncEngineSendChangesContext] class.
//
// # Accessing specific attributes
//
//   - [ICKSyncEngineSendChangesContext.Reason]: The reason for the send operation.
//   - [ICKSyncEngineSendChangesContext.Options]: The additional options for the send operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext
type ICKSyncEngineSendChangesContext interface {
	objectivec.IObject

	// Topic: Accessing specific attributes

	// The reason for the send operation.
	Reason() CKSyncEngineSyncReason
	// The additional options for the send operation.
	Options() ICKSyncEngineSendChangesOptions
}

// Init initializes the instance.
func (c CKSyncEngineSendChangesContext) Init() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineSendChangesContext) Autorelease() CKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesContext creates a new CKSyncEngineSendChangesContext instance.
func NewCKSyncEngineSendChangesContext() CKSyncEngineSendChangesContext {
	class := getCKSyncEngineSendChangesContextClass()
	rv := objc.Send[CKSyncEngineSendChangesContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The reason for the send operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext/reason
func (c CKSyncEngineSendChangesContext) Reason() CKSyncEngineSyncReason {
	rv := objc.Send[CKSyncEngineSyncReason](c.ID, objc.Sel("reason"))
	return CKSyncEngineSyncReason(rv)
}

// The additional options for the send operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesContext/options
func (c CKSyncEngineSendChangesContext) Options() ICKSyncEngineSendChangesOptions {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("options"))
	return CKSyncEngineSendChangesOptionsFromID(objc.ID(rv))
}
