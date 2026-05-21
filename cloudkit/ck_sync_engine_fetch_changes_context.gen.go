// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFetchChangesContext] class.
var (
	_CKSyncEngineFetchChangesContextClass     CKSyncEngineFetchChangesContextClass
	_CKSyncEngineFetchChangesContextClassOnce sync.Once
)

func getCKSyncEngineFetchChangesContextClass() CKSyncEngineFetchChangesContextClass {
	_CKSyncEngineFetchChangesContextClassOnce.Do(func() {
		_CKSyncEngineFetchChangesContextClass = CKSyncEngineFetchChangesContextClass{class: objc.GetClass("CKSyncEngineFetchChangesContext")}
	})
	return _CKSyncEngineFetchChangesContextClass
}

// GetCKSyncEngineFetchChangesContextClass returns the class object for CKSyncEngineFetchChangesContext.
func GetCKSyncEngineFetchChangesContextClass() CKSyncEngineFetchChangesContextClass {
	return getCKSyncEngineFetchChangesContextClass()
}

type CKSyncEngineFetchChangesContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchChangesContextClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchChangesContextClass) Alloc() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The context of an attempt to fetch changes from the server.
//
// # Overview
//
// The sync engine might attempt to fetch changes to the server for many
// reasons. For example, if you call
// [CKSyncEngine.FetchChangesWithCompletionHandler], it tries to fetch changes
// immediately. Or if it receives a push notification, it schedules a sync and
// fetch changes when the scheduler task runs.
//
// # Instance Properties
//
//   - [CKSyncEngineFetchChangesContext.Options]
//   - [CKSyncEngineFetchChangesContext.Reason]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext
type CKSyncEngineFetchChangesContext struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesContextFromID constructs a [CKSyncEngineFetchChangesContext] from an objc.ID.
//
// The context of an attempt to fetch changes from the server.
func CKSyncEngineFetchChangesContextFromID(id objc.ID) CKSyncEngineFetchChangesContext {
	return CKSyncEngineFetchChangesContext{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineFetchChangesContext implements ICKSyncEngineFetchChangesContext.
var _ ICKSyncEngineFetchChangesContext = CKSyncEngineFetchChangesContext{}

// An interface definition for the [CKSyncEngineFetchChangesContext] class.
//
// # Instance Properties
//
//   - [ICKSyncEngineFetchChangesContext.Options]
//   - [ICKSyncEngineFetchChangesContext.Reason]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext
type ICKSyncEngineFetchChangesContext interface {
	objectivec.IObject

	// Topic: Instance Properties

	Options() ICKSyncEngineFetchChangesOptions
	Reason() CKSyncEngineSyncReason
}

// Init initializes the instance.
func (c CKSyncEngineFetchChangesContext) Init() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchChangesContext) Autorelease() CKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesContext creates a new CKSyncEngineFetchChangesContext instance.
func NewCKSyncEngineFetchChangesContext() CKSyncEngineFetchChangesContext {
	class := getCKSyncEngineFetchChangesContextClass()
	rv := objc.Send[CKSyncEngineFetchChangesContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext/options
func (c CKSyncEngineFetchChangesContext) Options() ICKSyncEngineFetchChangesOptions {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("options"))
	return CKSyncEngineFetchChangesOptionsFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesContext/reason
func (c CKSyncEngineFetchChangesContext) Reason() CKSyncEngineSyncReason {
	rv := objc.Send[CKSyncEngineSyncReason](c.ID, objc.Sel("reason"))
	return CKSyncEngineSyncReason(rv)
}
