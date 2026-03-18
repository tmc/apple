// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileSecurity] class.
var (
	_NSFileSecurityClass     NSFileSecurityClass
	_NSFileSecurityClassOnce sync.Once
)

func getNSFileSecurityClass() NSFileSecurityClass {
	_NSFileSecurityClassOnce.Do(func() {
		_NSFileSecurityClass = NSFileSecurityClass{class: objc.GetClass("NSFileSecurity")}
	})
	return _NSFileSecurityClass
}

// GetNSFileSecurityClass returns the class object for NSFileSecurity.
func GetNSFileSecurityClass() NSFileSecurityClass {
	return getNSFileSecurityClass()
}

type NSFileSecurityClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileSecurityClass) Alloc() NSFileSecurity {
	rv := objc.Send[NSFileSecurity](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A stub class that encapsulates security information about a file.
//
// # Overview
// 
// [NSFileSecurity] contains no methods of its own. Instead, it is
// transparently bridged to [CFFileSecurity].
//
// [CFFileSecurity]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurity
//
// See: https://developer.apple.com/documentation/Foundation/NSFileSecurity
type NSFileSecurity struct {
	objectivec.Object
}

// NSFileSecurityFromID constructs a [NSFileSecurity] from an objc.ID.
//
// A stub class that encapsulates security information about a file.
func NSFileSecurityFromID(id objc.ID) NSFileSecurity {
	return NSFileSecurity{objectivec.Object{ID: id}}
}
// NOTE: NSFileSecurity adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileSecurity] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileSecurity
type INSFileSecurity interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding
}

// Init initializes the instance.
func (f NSFileSecurity) Init() NSFileSecurity {
	rv := objc.Send[NSFileSecurity](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileSecurity) Autorelease() NSFileSecurity {
	rv := objc.Send[NSFileSecurity](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileSecurity creates a new NSFileSecurity instance.
func NewNSFileSecurity() NSFileSecurity {
	class := getNSFileSecurityClass()
	rv := objc.Send[NSFileSecurity](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSFileSecurity/init(coder:)
func NewFileSecurityWithCoder(coder INSCoder) NSFileSecurity {
	instance := getNSFileSecurityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSFileSecurityFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSFileSecurity/init(coder:)
func (f NSFileSecurity) InitWithCoder(coder INSCoder) NSFileSecurity {
	rv := objc.Send[NSFileSecurity](f.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (f NSFileSecurity) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

