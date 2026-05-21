// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKReference] class.
var (
	_CKReferenceClass     CKReferenceClass
	_CKReferenceClassOnce sync.Once
)

func getCKReferenceClass() CKReferenceClass {
	_CKReferenceClassOnce.Do(func() {
		_CKReferenceClass = CKReferenceClass{class: objc.GetClass("CKReference")}
	})
	return _CKReferenceClass
}

// GetCKReferenceClass returns the class object for CKReference.
func GetCKReferenceClass() CKReferenceClass {
	return getCKReferenceClass()
}

type CKReferenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKReferenceClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKReferenceClass) Alloc() CKReference {
	rv := objc.Send[CKReference](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A relationship between two records in a record zone.
//
// # Overview
//
// A [CKReference] object creates a many-to-one relationship between records
// in your database. Each reference object stores information about the one
// record that is the target of the reference. You then save the reference
// object in the fields of one or more records to create a link from those
// records to the target. You may only create references between records
// within the same zone of the same database.
//
// References create a stronger relationship between records than just saving
// the ID of a record as a string. Specifically, you can use references to
// create an ownership model between two records. When the reference
// object’s action is [CKRecord.ReferenceAction.deleteSelf], the target of
// the reference—that is, the record in the reference’s
// [CKReference.RecordID] property—becomes the owner of the source record.
// Deleting the target (owner) record deletes all its source records. The
// deletion of any owned records can trigger further deletions if those
// records are the owners of other records. If a record contains two or more
// [CKReference] objects with an action of
// [CKRecord.ReferenceAction.deleteSelf], CloudKit deletes the record when it
// deletes any of the objects it references.
//
// To save multiple records that contain references between them, save the
// target records first or save all the records in one batch operation using
// [CKModifyRecordsOperation].
//
// # Interacting with Reference Objects
//
// You use reference objects to create strong links between two records and to
// search for related fields. When you create new records, you create
// reference objects and assign them to fields of your records. The only other
// time you create reference objects is when you build a search predicate to
// search for related records.
//
// # Linking to Another Record
//
// To link records together and create a strong relationship between them,
// create a new [CKReference] object, initialize it with the owner record, and
// assign that reference object to a field of the owned record. When you
// design the relationships among your own records, make the owner the more
// important of two related records. The owner record rarely depends on any
// records that point to it. The owner record is also the one that you
// typically fetch first from the database.
//
// The figure below shows an example of a relationship between a to-do list
// record and a set of item records that represent individual items to
// complete. The to-do list is the primary record, or owner, in the
// relationship because it represents the entire to-do list, including all
// items on the list. As a result, each item record has a field that contains
// a [CKReference] object that points to the owning to-do list record.
//
// [media-1965777]
//
// The following code sample shows how to create the reference object for each
// item record and configure it to point at the list record:
//
// An ownership type of organization is useful even if one object doesn’t
// explicitly own another. Ownership helps establish the relationships between
// records and how you search for them in the database. Ownership doesn’t
// require the deletion of the owned records when you delete their owner
// record. You can prevent such deletions by specifying the
// [CKRecord.ReferenceAction.none] action when you create a [CKReference]
// object.
//
// # Searching for Related Records
//
// When you want to find records for a single owner object, you create a
// [CKReference] object and use it to build your search predicate. When you
// use reference objects in search predicates, the search code looks only at
// the ID value in the reference object. It matches the ID in records of the
// specified type with the ID you provide in the [CKReference] object.
//
// The code sample below shows how to use a reference object to construct a
// query for the records in the figure above. The `listID` variable is a
// placeholder for the record ID of the list with the items you want to
// retrieve. The predicate tells the query object to search the `owningList`
// field of the target records and compare the reference object there with the
// one in the `recordToMatch` variable. Executing the query operation object
// returns the matching records asynchronously to the completion block you
// provide.
//
// # Creating a Reference
//
//   - [CKReference.InitWithRecordIDAction]: Creates a reference object that points to the record with the specified ID.
//   - [CKReference.InitWithRecordAction]: Creates a reference object that points to the specified record object.
//
// # Getting the Reference Attributes
//
//   - [CKReference.ReferenceAction]: The ownership behavior for the records.
//   - [CKReference.RecordID]: The ID of the referenced record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference
//
// [CKRecord.ReferenceAction.deleteSelf]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
// [CKRecord.ReferenceAction.none]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/none
type CKReference struct {
	objectivec.Object
}

// CKReferenceFromID constructs a [CKReference] from an objc.ID.
//
// A relationship between two records in a record zone.
func CKReferenceFromID(id objc.ID) CKReference {
	return CKReference{objectivec.Object{ID: id}}
}

// NOTE: CKReference adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKReference] class.
//
// # Creating a Reference
//
//   - [ICKReference.InitWithRecordIDAction]: Creates a reference object that points to the record with the specified ID.
//   - [ICKReference.InitWithRecordAction]: Creates a reference object that points to the specified record object.
//
// # Getting the Reference Attributes
//
//   - [ICKReference.ReferenceAction]: The ownership behavior for the records.
//   - [ICKReference.RecordID]: The ID of the referenced record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference
type ICKReference interface {
	objectivec.IObject

	// Topic: Creating a Reference

	// Creates a reference object that points to the record with the specified ID.
	InitWithRecordIDAction(recordID ICKRecordID, action CKReferenceAction) CKReference
	// Creates a reference object that points to the specified record object.
	InitWithRecordAction(record ICKRecord, action CKReferenceAction) CKReference

	// Topic: Getting the Reference Attributes

	// The ownership behavior for the records.
	ReferenceAction() CKReferenceAction
	// The ID of the referenced record.
	RecordID() ICKRecordID

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKReference) Init() CKReference {
	rv := objc.Send[CKReference](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKReference) Autorelease() CKReference {
	rv := objc.Send[CKReference](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKReference creates a new CKReference instance.
func NewCKReference() CKReference {
	class := getCKReferenceClass()
	rv := objc.Send[CKReference](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a reference object that points to the specified record object.
//
// record: The target record of the reference.
//
// action: The ownership options to use for the records. If you specify the
// [CKRecord.ReferenceAction.deleteSelf] option, the object that the
// `recordID` parameter references becomes the owner of (or acts as the parent
// of) any objects that use this reference object. For a list of possible
// values, see [CKRecord.ReferenceAction].
//
// # Return Value
//
// An initialized reference object that points to the specified record.
//
// # Discussion
//
// Use this method to initialize a reference to a local record object. You can
// reference a local record that you create, or one that you fetch from the
// server.
//
// When you create a reference object for use in a search predicate, the
// predicate ignores the value in the `action` parameter. Search predicates
// use only the ID of the record during their comparison.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(record:action:)
//
// [CKRecord.ReferenceAction.deleteSelf]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
// [CKRecord.ReferenceAction]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction
func NewCKReferenceWithRecordAction(record ICKRecord, action CKReferenceAction) CKReference {
	instance := getCKReferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecord:action:"), record, action)
	return CKReferenceFromID(rv)
}

// Creates a reference object that points to the record with the specified ID.
//
// recordID: The ID of the target record. This method throws an exception if you specify
// `nil` for this parameter.
//
// action: The ownership option use between the target record and any records that
// incorporate this reference object. If you specify the
// [CKRecord.ReferenceAction.deleteSelf] option, the record that the
// `recordID` parameter references becomes the owner of (or acts as the parent
// of) any objects that use this reference object. For a list of possible
// values, see [CKRecord.ReferenceAction].
//
// # Return Value
//
// An initialized reference object that points to the specified record.
//
// # Discussion
//
// Use this method when you have only the ID of the record for the target of a
// link. You might use this method if you save only the ID of the record to a
// local data cache.
//
// When you create a reference object for use in a search predicate, the
// predicate ignores the value in the `action` parameter. Search predicates
// use only the ID of the record during their comparison.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(recordID:action:)
//
// [CKRecord.ReferenceAction.deleteSelf]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
// [CKRecord.ReferenceAction]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction
func NewCKReferenceWithRecordIDAction(recordID ICKRecordID, action CKReferenceAction) CKReference {
	instance := getCKReferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordID:action:"), recordID, action)
	return CKReferenceFromID(rv)
}

// Creates a reference object that points to the record with the specified ID.
//
// recordID: The ID of the target record. This method throws an exception if you specify
// `nil` for this parameter.
//
// action: The ownership option use between the target record and any records that
// incorporate this reference object. If you specify the
// [CKRecord.ReferenceAction.deleteSelf] option, the record that the
// `recordID` parameter references becomes the owner of (or acts as the parent
// of) any objects that use this reference object. For a list of possible
// values, see [CKRecord.ReferenceAction].
//
// # Return Value
//
// An initialized reference object that points to the specified record.
//
// # Discussion
//
// Use this method when you have only the ID of the record for the target of a
// link. You might use this method if you save only the ID of the record to a
// local data cache.
//
// When you create a reference object for use in a search predicate, the
// predicate ignores the value in the `action` parameter. Search predicates
// use only the ID of the record during their comparison.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(recordID:action:)
//
// [CKRecord.ReferenceAction.deleteSelf]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
// [CKRecord.ReferenceAction]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction
func (c CKReference) InitWithRecordIDAction(recordID ICKRecordID, action CKReferenceAction) CKReference {
	rv := objc.Send[CKReference](c.ID, objc.Sel("initWithRecordID:action:"), recordID, action)
	return rv
}

// Creates a reference object that points to the specified record object.
//
// record: The target record of the reference.
//
// action: The ownership options to use for the records. If you specify the
// [CKRecord.ReferenceAction.deleteSelf] option, the object that the
// `recordID` parameter references becomes the owner of (or acts as the parent
// of) any objects that use this reference object. For a list of possible
// values, see [CKRecord.ReferenceAction].
//
// # Return Value
//
// An initialized reference object that points to the specified record.
//
// # Discussion
//
// Use this method to initialize a reference to a local record object. You can
// reference a local record that you create, or one that you fetch from the
// server.
//
// When you create a reference object for use in a search predicate, the
// predicate ignores the value in the `action` parameter. Search predicates
// use only the ID of the record during their comparison.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/init(record:action:)
//
// [CKRecord.ReferenceAction.deleteSelf]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
// [CKRecord.ReferenceAction]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction
func (c CKReference) InitWithRecordAction(record ICKRecord, action CKReferenceAction) CKReference {
	rv := objc.Send[CKReference](c.ID, objc.Sel("initWithRecord:action:"), record, action)
	return rv
}
func (c CKReference) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The ownership behavior for the records.
//
// # Discussion
//
// The value in this property determines which action, if any, to take when
// deleting the target of the reference object — that is, the object that
// the [CKReference.RecordID] property points to. When this property is
// [CKRecord.ReferenceAction.deleteSelf], deleting the target object deletes
// any records that contain that reference in one of their fields. When this
// property is [CKRecord.ReferenceAction.none], deleting the target object
// doesn’t delete any additional objects.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/action-swift.property
//
// [CKRecord.ReferenceAction.deleteSelf]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/deleteSelf
// [CKRecord.ReferenceAction.none]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/none
func (c CKReference) ReferenceAction() CKReferenceAction {
	rv := objc.Send[CKReferenceAction](c.ID, objc.Sel("referenceAction"))
	return CKReferenceAction(rv)
}

// The ID of the referenced record.
//
// # Discussion
//
// Use the ID in this property to fetch the record on the other end of the
// link.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/Reference/recordID
func (c CKReference) RecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordID"))
	return CKRecordIDFromID(objc.ID(rv))
}
