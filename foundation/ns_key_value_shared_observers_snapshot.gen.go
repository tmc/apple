// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSKeyValueSharedObserversSnapshot] class.
var (
	_NSKeyValueSharedObserversSnapshotClass     NSKeyValueSharedObserversSnapshotClass
	_NSKeyValueSharedObserversSnapshotClassOnce sync.Once
)

func getNSKeyValueSharedObserversSnapshotClass() NSKeyValueSharedObserversSnapshotClass {
	_NSKeyValueSharedObserversSnapshotClassOnce.Do(func() {
		_NSKeyValueSharedObserversSnapshotClass = NSKeyValueSharedObserversSnapshotClass{class: objc.GetClass("NSKeyValueSharedObserversSnapshot")}
	})
	return _NSKeyValueSharedObserversSnapshotClass
}

// GetNSKeyValueSharedObserversSnapshotClass returns the class object for NSKeyValueSharedObserversSnapshot.
func GetNSKeyValueSharedObserversSnapshotClass() NSKeyValueSharedObserversSnapshotClass {
	return getNSKeyValueSharedObserversSnapshotClass()
}

type NSKeyValueSharedObserversSnapshotClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSKeyValueSharedObserversSnapshotClass) Alloc() NSKeyValueSharedObserversSnapshot {
	rv := objc.Send[NSKeyValueSharedObserversSnapshot](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A collection of key-value observations which may be registered with
// multiple observable objects. Create using `-[NSKeyValueSharedObservers
// snapshot]`
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObserversSnapshot
type NSKeyValueSharedObserversSnapshot struct {
	objectivec.Object
}

// NSKeyValueSharedObserversSnapshotFromID constructs a [NSKeyValueSharedObserversSnapshot] from an objc.ID.
//
// A collection of key-value observations which may be registered with
// multiple observable objects. Create using `-[NSKeyValueSharedObservers
// snapshot]`
func NSKeyValueSharedObserversSnapshotFromID(id objc.ID) NSKeyValueSharedObserversSnapshot {
	return NSKeyValueSharedObserversSnapshot{objectivec.Object{ID: id}}
}
// NOTE: NSKeyValueSharedObserversSnapshot adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSKeyValueSharedObserversSnapshot] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObserversSnapshot
type INSKeyValueSharedObserversSnapshot interface {
	objectivec.IObject
}

// Init initializes the instance.
func (k NSKeyValueSharedObserversSnapshot) Init() NSKeyValueSharedObserversSnapshot {
	rv := objc.Send[NSKeyValueSharedObserversSnapshot](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k NSKeyValueSharedObserversSnapshot) Autorelease() NSKeyValueSharedObserversSnapshot {
	rv := objc.Send[NSKeyValueSharedObserversSnapshot](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSKeyValueSharedObserversSnapshot creates a new NSKeyValueSharedObserversSnapshot instance.
func NewNSKeyValueSharedObserversSnapshot() NSKeyValueSharedObserversSnapshot {
	class := getNSKeyValueSharedObserversSnapshotClass()
	rv := objc.Send[NSKeyValueSharedObserversSnapshot](objc.ID(class.class), objc.Sel("new"))
	return rv
}

