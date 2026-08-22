// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the interface for section objects vended by a fetched results controller.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo
type NSFetchedResultsSectionInfo interface {
	objectivec.IObject

	// The number of objects (rows) in the section.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/numberOfObjects
	NumberOfObjects() uint

	// The array of objects in the section.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/objects
	Objects() foundation.INSArray

	// The name of the section.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/name
	Name() string

	// The index title of the section.
	//
	// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/indexTitle
	IndexTitle() string
}

// NSFetchedResultsSectionInfoObject wraps an existing Objective-C object that conforms to the NSFetchedResultsSectionInfo protocol.
type NSFetchedResultsSectionInfoObject struct {
	objectivec.Object
}

func (o NSFetchedResultsSectionInfoObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFetchedResultsSectionInfoObjectFromID constructs a [NSFetchedResultsSectionInfoObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFetchedResultsSectionInfoObjectFromID(id objc.ID) NSFetchedResultsSectionInfoObject {
	return NSFetchedResultsSectionInfoObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The number of objects (rows) in the section.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/numberOfObjects
func (o NSFetchedResultsSectionInfoObject) NumberOfObjects() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("numberOfObjects"))
	return uint(rv)
}

// The array of objects in the section.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/objects
func (o NSFetchedResultsSectionInfoObject) Objects() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objects"))
	return foundation.NSArrayFromID(rv)
}

// The name of the section.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/name
func (o NSFetchedResultsSectionInfoObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The index title of the section.
//
// # Discussion
//
// This is used when displaying the index.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchedResultsSectionInfo/indexTitle
func (o NSFetchedResultsSectionInfoObject) IndexTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indexTitle"))
	return foundation.NSStringFromID(rv).String()
}
