// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKDatabaseOperation] class.
var (
	_CKDatabaseOperationClass     CKDatabaseOperationClass
	_CKDatabaseOperationClassOnce sync.Once
)

func getCKDatabaseOperationClass() CKDatabaseOperationClass {
	_CKDatabaseOperationClassOnce.Do(func() {
		_CKDatabaseOperationClass = CKDatabaseOperationClass{class: objc.GetClass("CKDatabaseOperation")}
	})
	return _CKDatabaseOperationClass
}

// GetCKDatabaseOperationClass returns the class object for CKDatabaseOperation.
func GetCKDatabaseOperationClass() CKDatabaseOperationClass {
	return getCKDatabaseOperationClass()
}

type CKDatabaseOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKDatabaseOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKDatabaseOperationClass) Alloc() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for operations that act upon databases in CloudKit.
//
// # Overview
//
// Database operations typically involve fetching and saving records and other
// database objects, as well as executing queries on the contents of the
// database. Use this class’s [CKDatabaseOperation.Database] property to tell the operation
// which database to use when you execute it. Don’t subclass this class or
// create instances of it. Instead, create instances of one of its concrete
// subclasses.
//
// # Accessing the Database
//
//   - [CKDatabaseOperation.Database]: The database that the operation uses.
//   - [CKDatabaseOperation.SetDatabase]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseOperation
type CKDatabaseOperation struct {
	CKOperation
}

// CKDatabaseOperationFromID constructs a [CKDatabaseOperation] from an objc.ID.
//
// The abstract base class for operations that act upon databases in CloudKit.
func CKDatabaseOperationFromID(id objc.ID) CKDatabaseOperation {
	return CKDatabaseOperation{CKOperation: CKOperationFromID(id)}
}

// NOTE: CKDatabaseOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKDatabaseOperation] class.
//
// # Accessing the Database
//
//   - [ICKDatabaseOperation.Database]: The database that the operation uses.
//   - [ICKDatabaseOperation.SetDatabase]
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseOperation
type ICKDatabaseOperation interface {
	ICKOperation

	// Topic: Accessing the Database

	// The database that the operation uses.
	Database() ICKDatabase
	SetDatabase(value ICKDatabase)
}

// Init initializes the instance.
func (c CKDatabaseOperation) Init() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKDatabaseOperation) Autorelease() CKDatabaseOperation {
	rv := objc.Send[CKDatabaseOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKDatabaseOperation creates a new CKDatabaseOperation instance.
func NewCKDatabaseOperation() CKDatabaseOperation {
	class := getCKDatabaseOperationClass()
	rv := objc.Send[CKDatabaseOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The database that the operation uses.
//
// # Discussion
//
// For operations that you execute in a custom queue, use this property to
// specify the target database. Setting the database also sets the
// corresponding container, which it inherits from [CKOperation]. If this
// property’s value is `nil`, the operation targets the user’s private
// database.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKDatabaseOperation/database
func (c CKDatabaseOperation) Database() ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("database"))
	return CKDatabaseFromID(objc.ID(rv))
}
func (c CKDatabaseOperation) SetDatabase(value ICKDatabase) {
	objc.Send[struct{}](c.ID, objc.Sel("setDatabase:"), value)
}
