// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKServerChangeToken] class.
var (
	_CKServerChangeTokenClass     CKServerChangeTokenClass
	_CKServerChangeTokenClassOnce sync.Once
)

func getCKServerChangeTokenClass() CKServerChangeTokenClass {
	_CKServerChangeTokenClassOnce.Do(func() {
		_CKServerChangeTokenClass = CKServerChangeTokenClass{class: objc.GetClass("CKServerChangeToken")}
	})
	return _CKServerChangeTokenClass
}

// GetCKServerChangeTokenClass returns the class object for CKServerChangeToken.
func GetCKServerChangeTokenClass() CKServerChangeTokenClass {
	return getCKServerChangeTokenClass()
}

type CKServerChangeTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKServerChangeTokenClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKServerChangeTokenClass) Alloc() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An opaque token that represents a specific point in a database’s history.
//
// # Overview
//
// CloudKit uses server change tokens to record significant events in a
// database’s history, such as record creation, modification, and deletion.
// Using change tokens helps reduce the cost of a fetch operation — both the
// time to execute the fetch and the overall number of records it returns.
//
// You don’t create change tokens. Instead,
// [CKFetchDatabaseChangesOperation] and [CKFetchRecordZoneChangesOperation]
// provide them during their execution and when they complete. Cache each
// token as you receive it, overwriting any previous token for the database or
// record zone you’re fetching from. Then, pass the cached token with your
// next fetch and CloudKit returns only the changes that occur after that
// point. Don’t infer any behavior or order from a token’s contents.
//
// The change tokens that [CKFetchDatabaseChangesOperation] provides aren’t
// compatible with [CKFetchRecordZoneChangesOperation] and vice versa, so
// segregate them in your cache.
//
// Change tokens conform to [NSSecureCoding] and are safe to cache on-disk, as
// the following example shows:
//
// # Initializers
//
//   - [CKServerChangeToken.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKServerChangeToken
//
// [NSSecureCoding]: https://developer.apple.com/documentation/Foundation/NSSecureCoding
type CKServerChangeToken struct {
	objectivec.Object
}

// CKServerChangeTokenFromID constructs a [CKServerChangeToken] from an objc.ID.
//
// An opaque token that represents a specific point in a database’s history.
func CKServerChangeTokenFromID(id objc.ID) CKServerChangeToken {
	return CKServerChangeToken{objectivec.Object{ID: id}}
}

// NOTE: CKServerChangeToken adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKServerChangeToken] class.
//
// # Initializers
//
//   - [ICKServerChangeToken.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKServerChangeToken
type ICKServerChangeToken interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CKServerChangeToken

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKServerChangeToken) Init() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKServerChangeToken) Autorelease() CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKServerChangeToken creates a new CKServerChangeToken instance.
func NewCKServerChangeToken() CKServerChangeToken {
	class := getCKServerChangeTokenClass()
	rv := objc.Send[CKServerChangeToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKServerChangeToken/init(coder:)
func NewCKServerChangeTokenWithCoder(coder foundation.INSCoder) CKServerChangeToken {
	instance := getCKServerChangeTokenClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CKServerChangeTokenFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKServerChangeToken/init(coder:)
func (c CKServerChangeToken) InitWithCoder(coder foundation.INSCoder) CKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CKServerChangeToken) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
