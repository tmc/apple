// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSKeyValueSharedObservers] class.
var (
	_NSKeyValueSharedObserversClass     NSKeyValueSharedObserversClass
	_NSKeyValueSharedObserversClassOnce sync.Once
)

func getNSKeyValueSharedObserversClass() NSKeyValueSharedObserversClass {
	_NSKeyValueSharedObserversClassOnce.Do(func() {
		_NSKeyValueSharedObserversClass = NSKeyValueSharedObserversClass{class: objc.GetClass("NSKeyValueSharedObservers")}
	})
	return _NSKeyValueSharedObserversClass
}

// GetNSKeyValueSharedObserversClass returns the class object for NSKeyValueSharedObservers.
func GetNSKeyValueSharedObserversClass() NSKeyValueSharedObserversClass {
	return getNSKeyValueSharedObserversClass()
}

type NSKeyValueSharedObserversClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSKeyValueSharedObserversClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSKeyValueSharedObserversClass) Alloc() NSKeyValueSharedObservers {
	rv := objc.Send[NSKeyValueSharedObservers](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A collection of key-value observations which may be registered with
// multiple observable objects
//
// # Initializers
//
//   - [NSKeyValueSharedObservers.InitWithObservableClass]: A new collection of observables for an observable object of the given class
//
// # Instance Methods
//
//   - [NSKeyValueSharedObservers.AddSharedObserverForKeyOptionsContext]: Add a new observer to the collection.
//   - [NSKeyValueSharedObservers.Snapshot]: A momentary snapshot of all observers added to the collection thus far, that can be assigned to an observable using `-[NSObject ]`
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers
type NSKeyValueSharedObservers struct {
	objectivec.Object
}

// NSKeyValueSharedObserversFromID constructs a [NSKeyValueSharedObservers] from an objc.ID.
//
// A collection of key-value observations which may be registered with
// multiple observable objects
func NSKeyValueSharedObserversFromID(id objc.ID) NSKeyValueSharedObservers {
	return NSKeyValueSharedObservers{objectivec.Object{ID: id}}
}

// NOTE: NSKeyValueSharedObservers adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSKeyValueSharedObservers] class.
//
// # Initializers
//
//   - [INSKeyValueSharedObservers.InitWithObservableClass]: A new collection of observables for an observable object of the given class
//
// # Instance Methods
//
//   - [INSKeyValueSharedObservers.AddSharedObserverForKeyOptionsContext]: Add a new observer to the collection.
//   - [INSKeyValueSharedObservers.Snapshot]: A momentary snapshot of all observers added to the collection thus far, that can be assigned to an observable using `-[NSObject ]`
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers
type INSKeyValueSharedObservers interface {
	objectivec.IObject

	// Topic: Initializers

	// A new collection of observables for an observable object of the given class
	InitWithObservableClass(observableClass objc.Class) NSKeyValueSharedObservers

	// Topic: Instance Methods

	// Add a new observer to the collection.
	AddSharedObserverForKeyOptionsContext(observer objectivec.Object, key string, options uint, context unsafe.Pointer)
	// A momentary snapshot of all observers added to the collection thus far, that can be assigned to an observable using `-[NSObject ]`
	Snapshot() INSKeyValueSharedObserversSnapshot
}

// Init initializes the instance.
func (k NSKeyValueSharedObservers) Init() NSKeyValueSharedObservers {
	rv := objc.Send[NSKeyValueSharedObservers](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k NSKeyValueSharedObservers) Autorelease() NSKeyValueSharedObservers {
	rv := objc.Send[NSKeyValueSharedObservers](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSKeyValueSharedObservers creates a new NSKeyValueSharedObservers instance.
func NewNSKeyValueSharedObservers() NSKeyValueSharedObservers {
	class := getNSKeyValueSharedObserversClass()
	rv := objc.Send[NSKeyValueSharedObservers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A new collection of observables for an observable object of the given class
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/init(observableClass:)
func NewKeyValueSharedObserversWithObservableClass(observableClass objc.Class) NSKeyValueSharedObservers {
	instance := getNSKeyValueSharedObserversClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObservableClass:"), observableClass)
	return NSKeyValueSharedObserversFromID(rv)
}

// A new collection of observables for an observable object of the given class
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/init(observableClass:)
func (k NSKeyValueSharedObservers) InitWithObservableClass(observableClass objc.Class) NSKeyValueSharedObservers {
	rv := objc.Send[NSKeyValueSharedObservers](k.ID, objc.Sel("initWithObservableClass:"), observableClass)
	return rv
}

// Add a new observer to the collection.
//
// observer: The observer object to register for KVO notifications. The observer must
// implement the key-value observing method “
//
// key: Key of the property being observed. This cannot be a nested key path or a
// computed property
//
// options: A combination of NSKeyValueObservingOptions values that specify what is
// included in observation notifications. For possible values see
// NSKeyValueObservingOptions.
//
// context: Arbitrary data which is passed to the observer object
//
// # Discussion
//
// This method works like `-[NSObject ]`, but observations on nested and
// computed properties are disallowed. Observers are not registered until
// `setSharedObservers` is called on the observable.
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/addSharedObserver(_:forKey:options:context:)
func (k NSKeyValueSharedObservers) AddSharedObserverForKeyOptionsContext(observer objectivec.Object, key string, options uint, context unsafe.Pointer) {
	objc.Send[objc.ID](k.ID, objc.Sel("addSharedObserver:forKey:options:context:"), observer, objc.String(key), options, context)
}

// A momentary snapshot of all observers added to the collection thus far,
// that can be assigned to an observable using `-[NSObject ]`
//
// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSharedObservers/snapshot()
func (k NSKeyValueSharedObservers) Snapshot() INSKeyValueSharedObserversSnapshot {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("snapshot"))
	return NSKeyValueSharedObserversSnapshotFromID(rv)
}
