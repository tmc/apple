// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCompoundPredicate] class.
var (
	_NSCompoundPredicateClass     NSCompoundPredicateClass
	_NSCompoundPredicateClassOnce sync.Once
)

func getNSCompoundPredicateClass() NSCompoundPredicateClass {
	_NSCompoundPredicateClassOnce.Do(func() {
		_NSCompoundPredicateClass = NSCompoundPredicateClass{class: objc.GetClass("NSCompoundPredicate")}
	})
	return _NSCompoundPredicateClass
}

// GetNSCompoundPredicateClass returns the class object for NSCompoundPredicate.
func GetNSCompoundPredicateClass() NSCompoundPredicateClass {
	return getNSCompoundPredicateClass()
}

type NSCompoundPredicateClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCompoundPredicateClass) Alloc() NSCompoundPredicate {
	rv := objc.Send[NSCompoundPredicate](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A specialized predicate that evaluates logical combinations of other
// predicates.
//
// # Overview
// 
// Use [NSCompoundPredicate] to create an [AND] or [OR] compound predicate of
// one or more other predicates, or the [NOT] of a single predicate. For the
// logical [AND] and [OR] operations:
// 
// - An [AND] predicate with no subpredicates evaluates to [true]. - An [OR]
// predicate with no subpredicates evaluates to [false]. - A compound
// predicate with one or more subpredicates evaluates to the truth of its
// subpredicates.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating Compound Predicates
//
//   - [NSCompoundPredicate.InitWithTypeSubpredicates]: Returns the receiver that a specified type initializes using predicates from a specified array.
//
// # Getting Information About a Compound Predicate
//
//   - [NSCompoundPredicate.CompoundPredicateType]: The predicate type for the receiver.
//   - [NSCompoundPredicate.Subpredicates]: The receiver’s subpredicates.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate
type NSCompoundPredicate struct {
	NSPredicate
}

// NSCompoundPredicateFromID constructs a [NSCompoundPredicate] from an objc.ID.
//
// A specialized predicate that evaluates logical combinations of other
// predicates.
func NSCompoundPredicateFromID(id objc.ID) NSCompoundPredicate {
	return NSCompoundPredicate{NSPredicate: NSPredicateFromID(id)}
}
// NOTE: NSCompoundPredicate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCompoundPredicate] class.
//
// # Creating Compound Predicates
//
//   - [INSCompoundPredicate.InitWithTypeSubpredicates]: Returns the receiver that a specified type initializes using predicates from a specified array.
//
// # Getting Information About a Compound Predicate
//
//   - [INSCompoundPredicate.CompoundPredicateType]: The predicate type for the receiver.
//   - [INSCompoundPredicate.Subpredicates]: The receiver’s subpredicates.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate
type INSCompoundPredicate interface {
	INSPredicate

	// Topic: Creating Compound Predicates

	// Returns the receiver that a specified type initializes using predicates from a specified array.
	InitWithTypeSubpredicates(type_ NSCompoundPredicateType, subpredicates []NSPredicate) NSCompoundPredicate

	// Topic: Getting Information About a Compound Predicate

	// The predicate type for the receiver.
	CompoundPredicateType() NSCompoundPredicateType
	// The receiver’s subpredicates.
	Subpredicates() INSArray
}





// Init initializes the instance.
func (c NSCompoundPredicate) Init() NSCompoundPredicate {
	rv := objc.Send[NSCompoundPredicate](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCompoundPredicate) Autorelease() NSCompoundPredicate {
	rv := objc.Send[NSCompoundPredicate](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCompoundPredicate creates a new NSCompoundPredicate instance.
func NewNSCompoundPredicate() NSCompoundPredicate {
	class := getNSCompoundPredicateClass()
	rv := objc.Send[NSCompoundPredicate](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Returns a new predicate that you form using an AND operation on the
// predicates in a specified array.
//
// subpredicates: An array of [NSPredicate] objects.
//
// # Return Value
// 
// A new predicate formed by AND-ing the predicates specified by
// `subpredicates`.
//
// # Discussion
// 
// An AND predicate with no subpredicates evaluates to TRUE.
// 
// # Special Considerations
// 
// For applications linked on macOS 10.5 or later, the `subpredicates` array
// is copied. For applications linked on OS X v10.4, the `subpredicates` array
// is retained (for binary compatibility).
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(andPredicateWithSubpredicates:)
func NewCompoundPredicateAndPredicateWithSubpredicates(subpredicates []NSPredicate) NSCompoundPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSCompoundPredicateClass().class), objc.Sel("andPredicateWithSubpredicates:"), objectivec.IObjectSliceToNSArray(subpredicates))
	return NSCompoundPredicateFromID(rv)
}


// Returns a new predicate that you form using a NOT operation on a specified
// predicate.
//
// predicate: A predicate.
//
// # Return Value
// 
// A new predicate formed by NOT-ing the predicate specified by `predicate`.
//
// # Discussion
// 
// For applications linked on macOS 10.5 or later, the `subpredicates` array
// is copied. For applications linked on OS X v10.4, the `subpredicates` array
// is retained (for binary compatibility).
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(notPredicateWithSubpredicate:)
func NewCompoundPredicateNotPredicateWithSubpredicate(predicate INSPredicate) NSCompoundPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSCompoundPredicateClass().class), objc.Sel("notPredicateWithSubpredicate:"), predicate)
	return NSCompoundPredicateFromID(rv)
}


// Returns a new predicate that you form using an OR operation on the
// predicates in a specified array.
//
// subpredicates: An array of [NSPredicate] objects.
//
// # Return Value
// 
// A new predicate formed by OR-ing the predicates specified by
// `subpredicates`.
//
// # Discussion
// 
// An OR predicate with no subpredicates evaluates to FALSE.
// 
// # Special Considerations
// 
// For applications linked on macOS 10.5 or later, the `subpredicates` array
// is copied. For applications linked on OS X v10.4, the `subpredicates` array
// is retained (for binary compatibility).
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(orPredicateWithSubpredicates:)
func NewCompoundPredicateOrPredicateWithSubpredicates(subpredicates []NSPredicate) NSCompoundPredicate {
	rv := objc.Send[objc.ID](objc.ID(getNSCompoundPredicateClass().class), objc.Sel("orPredicateWithSubpredicates:"), objectivec.IObjectSliceToNSArray(subpredicates))
	return NSCompoundPredicateFromID(rv)
}


// Creates a predicate by decoding from the coder you specify.
//
// coder: The coder to read data from.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(coder:)
func NewCompoundPredicateWithCoder(coder INSCoder) NSCompoundPredicate {
	instance := getNSCompoundPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCompoundPredicateFromID(rv)
}


// Returns the receiver that a specified type initializes using predicates
// from a specified array.
//
// type: The type of the new predicate.
//
// subpredicates: An array of [NSPredicate] objects.
//
// # Return Value
// 
// The receiver initialized with its type set to type and subpredicates array
// to `subpredicates`.
//
// # Discussion
// 
// For applications linked on macOS 10.5 or later, the `subpredicates` array
// is copied. For applications linked on OS X v10.4, the `subpredicates` array
// is retained (for binary compatibility).
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(type:subpredicates:)
func NewCompoundPredicateWithTypeSubpredicates(type_ NSCompoundPredicateType, subpredicates []NSPredicate) NSCompoundPredicate {
	instance := getNSCompoundPredicateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:subpredicates:"), type_, objectivec.IObjectSliceToNSArray(subpredicates))
	return NSCompoundPredicateFromID(rv)
}







// Returns the receiver that a specified type initializes using predicates
// from a specified array.
//
// type: The type of the new predicate.
//
// subpredicates: An array of [NSPredicate] objects.
//
// # Return Value
// 
// The receiver initialized with its type set to type and subpredicates array
// to `subpredicates`.
//
// # Discussion
// 
// For applications linked on macOS 10.5 or later, the `subpredicates` array
// is copied. For applications linked on OS X v10.4, the `subpredicates` array
// is retained (for binary compatibility).
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(type:subpredicates:)
func (c NSCompoundPredicate) InitWithTypeSubpredicates(type_ NSCompoundPredicateType, subpredicates []NSPredicate) NSCompoundPredicate {
	rv := objc.Send[NSCompoundPredicate](c.ID, objc.Sel("initWithType:subpredicates:"), type_, objectivec.IObjectSliceToNSArray(subpredicates))
	return rv
}












// The predicate type for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/compoundPredicateType
func (c NSCompoundPredicate) CompoundPredicateType() NSCompoundPredicateType {
	rv := objc.Send[NSCompoundPredicateType](c.ID, objc.Sel("compoundPredicateType"))
	return NSCompoundPredicateType(rv)
}



// The receiver’s subpredicates.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/subpredicates
func (c NSCompoundPredicate) Subpredicates() INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subpredicates"))
	return NSArrayFromID(objc.ID(rv))
}




























