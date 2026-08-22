// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AUAudioUnitBusArray] class.
var (
	_AUAudioUnitBusArrayClass     AUAudioUnitBusArrayClass
	_AUAudioUnitBusArrayClassOnce sync.Once
)

func getAUAudioUnitBusArrayClass() AUAudioUnitBusArrayClass {
	_AUAudioUnitBusArrayClassOnce.Do(func() {
		_AUAudioUnitBusArrayClass = AUAudioUnitBusArrayClass{class: objc.GetClass("AUAudioUnitBusArray")}
	})
	return _AUAudioUnitBusArrayClass
}

// GetAUAudioUnitBusArrayClass returns the class object for AUAudioUnitBusArray.
func GetAUAudioUnitBusArrayClass() AUAudioUnitBusArrayClass {
	return getAUAudioUnitBusArrayClass()
}

type AUAudioUnitBusArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUAudioUnitBusArrayClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUAudioUnitBusArrayClass) Alloc() AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that defines a container for an audio unit’s input or output
// busses.
//
// # Overview
//
// Hosts can observe a bus property across all busses by using KVO on a bus
// array object, without having to observe it on each individual bus. Some
// audio units (e.g. mixers) support variable numbers of busses, via
// subclassing. When the bus count changes, a KVO notification is sent on the
// audio unit’s [AUAudioUnit.InputBusses] or [AUAudioUnit.OutputBusses]
// property, as appropriate.
//
// This version 3 class is bridged to the version 2
// `kAudioUnitProperty_ElementCount` API.
//
// # Initialization
//
//   - [AUAudioUnitBusArray.InitWithAudioUnitBusType]: Initializes an empty bus array.
//   - [AUAudioUnitBusArray.InitWithAudioUnitBusTypeBusses]: Initializes a bus array by making a copy of the supplied busses.
//
// # Bus Array Methods and Properties
//
//   - [AUAudioUnitBusArray.Count]: The number of busses in the array.
//   - [AUAudioUnitBusArray.IsCountChangeable]: Determines whether the array can have a variable number of busses.
//   - [AUAudioUnitBusArray.OwnerAudioUnit]: The audio unit that owns the bus array.
//   - [AUAudioUnitBusArray.BusType]: Determines whether the bus array is for input or output.
//   - [AUAudioUnitBusArray.ObjectAtIndexedSubscript]: Returns the bus at the specified index.
//   - [AUAudioUnitBusArray.SetBusCountError]: Changes the number of busses in the array.
//   - [AUAudioUnitBusArray.AddObserverToAllBussesForKeyPathOptionsContext]: Adds a KVO observer for a given property on all busses in the array.
//   - [AUAudioUnitBusArray.RemoveObserverFromAllBussesForKeyPathContext]: Removes a KVO observer for a given property on all busses in the array.
//
// # Audio Unit Implementations
//
//   - [AUAudioUnitBusArray.ReplaceBusses]: Replaces the current bus array with a copy of the supplied bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray
type AUAudioUnitBusArray struct {
	objectivec.Object
}

// AUAudioUnitBusArrayFromID constructs a [AUAudioUnitBusArray] from an objc.ID.
//
// A class that defines a container for an audio unit’s input or output
// busses.
func AUAudioUnitBusArrayFromID(id objc.ID) AUAudioUnitBusArray {
	return AUAudioUnitBusArray{objectivec.Object{ID: id}}
}

// NOTE: AUAudioUnitBusArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUAudioUnitBusArray] class.
//
// # Initialization
//
//   - [IAUAudioUnitBusArray.InitWithAudioUnitBusType]: Initializes an empty bus array.
//   - [IAUAudioUnitBusArray.InitWithAudioUnitBusTypeBusses]: Initializes a bus array by making a copy of the supplied busses.
//
// # Bus Array Methods and Properties
//
//   - [IAUAudioUnitBusArray.Count]: The number of busses in the array.
//   - [IAUAudioUnitBusArray.IsCountChangeable]: Determines whether the array can have a variable number of busses.
//   - [IAUAudioUnitBusArray.OwnerAudioUnit]: The audio unit that owns the bus array.
//   - [IAUAudioUnitBusArray.BusType]: Determines whether the bus array is for input or output.
//   - [IAUAudioUnitBusArray.ObjectAtIndexedSubscript]: Returns the bus at the specified index.
//   - [IAUAudioUnitBusArray.SetBusCountError]: Changes the number of busses in the array.
//   - [IAUAudioUnitBusArray.AddObserverToAllBussesForKeyPathOptionsContext]: Adds a KVO observer for a given property on all busses in the array.
//   - [IAUAudioUnitBusArray.RemoveObserverFromAllBussesForKeyPathContext]: Removes a KVO observer for a given property on all busses in the array.
//
// # Audio Unit Implementations
//
//   - [IAUAudioUnitBusArray.ReplaceBusses]: Replaces the current bus array with a copy of the supplied bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray
type IAUAudioUnitBusArray interface {
	objectivec.IObject

	// Topic: Initialization

	// Initializes an empty bus array.
	InitWithAudioUnitBusType(owner IAUAudioUnit, busType AUAudioUnitBusType) AUAudioUnitBusArray
	// Initializes a bus array by making a copy of the supplied busses.
	InitWithAudioUnitBusTypeBusses(owner IAUAudioUnit, busType AUAudioUnitBusType, busArray []AUAudioUnitBus) AUAudioUnitBusArray

	// Topic: Bus Array Methods and Properties

	// The number of busses in the array.
	Count() uint
	// Determines whether the array can have a variable number of busses.
	IsCountChangeable() bool
	// The audio unit that owns the bus array.
	OwnerAudioUnit() IAUAudioUnit
	// Determines whether the bus array is for input or output.
	BusType() AUAudioUnitBusType
	// Returns the bus at the specified index.
	ObjectAtIndexedSubscript(index uint) IAUAudioUnitBus
	// Changes the number of busses in the array.
	SetBusCountError(count uint) (bool, error)
	// Adds a KVO observer for a given property on all busses in the array.
	AddObserverToAllBussesForKeyPathOptionsContext(observer objectivec.NSObject, keyPath string, options uint, context unsafe.Pointer)
	// Removes a KVO observer for a given property on all busses in the array.
	RemoveObserverFromAllBussesForKeyPathContext(observer objectivec.NSObject, keyPath string, context unsafe.Pointer)

	// Topic: Audio Unit Implementations

	// Replaces the current bus array with a copy of the supplied bus array.
	ReplaceBusses(busArray []AUAudioUnitBus)
}

// Init initializes the instance.
func (a AUAudioUnitBusArray) Init() AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AUAudioUnitBusArray) Autorelease() AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUAudioUnitBusArray creates a new AUAudioUnitBusArray instance.
func NewAUAudioUnitBusArray() AUAudioUnitBusArray {
	class := getAUAudioUnitBusArrayClass()
	rv := objc.Send[AUAudioUnitBusArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an empty bus array.
//
// owner: The audio unit that owns the bus array.
//
// busType: Determines whether the bus array is for input or output.
//
// # Return Value
//
// A newly-initialized bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:)
func NewAudioUnitBusArrayWithAudioUnitBusType(owner IAUAudioUnit, busType AUAudioUnitBusType) AUAudioUnitBusArray {
	instance := getAUAudioUnitBusArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioUnit:busType:"), owner, busType)
	return AUAudioUnitBusArrayFromID(rv)
}

// Initializes a bus array by making a copy of the supplied busses.
//
// owner: The audio unit that owns the bus array.
//
// busType: Determines whether the busses are for input or output.
//
// busArray: An array of busses.
//
// # Return Value
//
// A newly-initialized bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:busses:)
func NewAudioUnitBusArrayWithAudioUnitBusTypeBusses(owner IAUAudioUnit, busType AUAudioUnitBusType, busArray []AUAudioUnitBus) AUAudioUnitBusArray {
	instance := getAUAudioUnitBusArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioUnit:busType:busses:"), owner, busType, objectivec.IObjectSliceToNSArray(busArray))
	return AUAudioUnitBusArrayFromID(rv)
}

// Initializes an empty bus array.
//
// owner: The audio unit that owns the bus array.
//
// busType: Determines whether the bus array is for input or output.
//
// # Return Value
//
// A newly-initialized bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:)
func (a AUAudioUnitBusArray) InitWithAudioUnitBusType(owner IAUAudioUnit, busType AUAudioUnitBusType) AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](a.ID, objc.Sel("initWithAudioUnit:busType:"), owner, busType)
	return rv
}

// Initializes a bus array by making a copy of the supplied busses.
//
// owner: The audio unit that owns the bus array.
//
// busType: Determines whether the busses are for input or output.
//
// busArray: An array of busses.
//
// # Return Value
//
// A newly-initialized bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/init(audioUnit:busType:busses:)
func (a AUAudioUnitBusArray) InitWithAudioUnitBusTypeBusses(owner IAUAudioUnit, busType AUAudioUnitBusType, busArray []AUAudioUnitBus) AUAudioUnitBusArray {
	rv := objc.Send[AUAudioUnitBusArray](a.ID, objc.Sel("initWithAudioUnit:busType:busses:"), owner, busType, objectivec.IObjectSliceToNSArray(busArray))
	return rv
}

// Returns the bus at the specified index.
//
// index: An index corresponding to a bus in the array.
//
// # Return Value
//
// The bus located at the specified index.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/subscript(_:)
func (a AUAudioUnitBusArray) ObjectAtIndexedSubscript(index uint) IAUAudioUnitBus {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return AUAudioUnitBusFromID(rv)
}

// Changes the number of busses in the array.
//
// count: The new number of busses in the array.
//
// # Discussion
//
// - false if the operation failed.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/setBusCount(_:)
func (a AUAudioUnitBusArray) SetBusCountError(count uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setBusCount:error:"), count, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setBusCount:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Adds a KVO observer for a given property on all busses in the array.
//
// observer: The KVO observer.
//
// keyPath: The property’s key path.
//
// options: The KVO options.
//
// context: The KVO context.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/addObserver(toAllBusses:forKeyPath:options:context:)
func (a AUAudioUnitBusArray) AddObserverToAllBussesForKeyPathOptionsContext(observer objectivec.NSObject, keyPath string, options uint, context unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("addObserverToAllBusses:forKeyPath:options:context:"), observer, objc.String(keyPath), options, context)
}

// Removes a KVO observer for a given property on all busses in the array.
//
// observer: The KVO observer.
//
// keyPath: The property’s key path.
//
// context: The KVO context.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/removeObserver(fromAllBusses:forKeyPath:context:)
func (a AUAudioUnitBusArray) RemoveObserverFromAllBussesForKeyPathContext(observer objectivec.NSObject, keyPath string, context unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeObserverFromAllBusses:forKeyPath:context:"), observer, objc.String(keyPath), context)
}

// Replaces the current bus array with a copy of the supplied bus array.
//
// busArray: The new bus array.
//
// # Discussion
//
// The base class issues KVO notifications.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/replaceBusses(_:)
func (a AUAudioUnitBusArray) ReplaceBusses(busArray []AUAudioUnitBus) {
	objc.Send[objc.ID](a.ID, objc.Sel("replaceBusses:"), objectivec.IObjectSliceToNSArray(busArray))
}

// The number of busses in the array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/count
func (a AUAudioUnitBusArray) Count() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("count"))
	return rv
}

// Determines whether the array can have a variable number of busses.
//
// # Discussion
//
// The base implementation default value is false.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/isCountChangeable
func (a AUAudioUnitBusArray) IsCountChangeable() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isCountChangeable"))
	return rv
}

// The audio unit that owns the bus array.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/ownerAudioUnit
func (a AUAudioUnitBusArray) OwnerAudioUnit() IAUAudioUnit {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("ownerAudioUnit"))
	return AUAudioUnitFromID(objc.ID(rv))
}

// Determines whether the bus array is for input or output.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBusArray/busType
func (a AUAudioUnitBusArray) BusType() AUAudioUnitBusType {
	rv := objc.Send[AUAudioUnitBusType](a.ID, objc.Sel("busType"))
	return AUAudioUnitBusType(rv)
}
