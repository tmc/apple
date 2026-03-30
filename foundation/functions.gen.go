// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Foundation: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Foundation: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Foundation: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _nSAllHashTableObjects func(table *NSHashTable) NSArray

// NSAllHashTableObjects returns all of the elements in the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllHashTableObjects(_:)
func NSAllHashTableObjects(table *NSHashTable) NSArray {
	if _nSAllHashTableObjects == nil {
		panic("Foundation: symbol NSAllHashTableObjects not loaded")
	}
	return _nSAllHashTableObjects(table)
}

var _nSAllMapTableKeys func(table *NSMapTable) NSArray

// NSAllMapTableKeys returns all of the keys in the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllMapTableKeys(_:)
func NSAllMapTableKeys(table *NSMapTable) NSArray {
	if _nSAllMapTableKeys == nil {
		panic("Foundation: symbol NSAllMapTableKeys not loaded")
	}
	return _nSAllMapTableKeys(table)
}

var _nSAllMapTableValues func(table *NSMapTable) NSArray

// NSAllMapTableValues returns all of the values in the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllMapTableValues(_:)
func NSAllMapTableValues(table *NSMapTable) NSArray {
	if _nSAllMapTableValues == nil {
		panic("Foundation: symbol NSAllMapTableValues not loaded")
	}
	return _nSAllMapTableValues(table)
}

var _nSAllocateCollectable func(size uint, options uint) unsafe.Pointer

// NSAllocateCollectable allocates collectable memory.
//
// Deprecated: Garbage collection is deprecated in OS X v10.8; instead,you should use AutomaticReference Counting—see [Transitioning to ARC Release Notes](<https://developer.apple.com/library/archive/releasenotes/ObjectiveC/RN-TransitioningToARC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011226>).
//
// See: https://developer.apple.com/documentation/Foundation/NSAllocateCollectable
func NSAllocateCollectable(size uint, options uint) unsafe.Pointer {
	if _nSAllocateCollectable == nil {
		panic("Foundation: symbol NSAllocateCollectable not loaded")
	}
	return _nSAllocateCollectable(size, options)
}

var _nSAllocateMemoryPages func(bytes uint) unsafe.Pointer

// NSAllocateMemoryPages allocates a new block of memory.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllocateMemoryPages(_:)
func NSAllocateMemoryPages(bytes uint) unsafe.Pointer {
	if _nSAllocateMemoryPages == nil {
		panic("Foundation: symbol NSAllocateMemoryPages not loaded")
	}
	return _nSAllocateMemoryPages(bytes)
}

var _nSAllocateObject func(aClass objc.Class, extraBytes uint, zone *NSZone) objectivec.Object

// NSAllocateObject creates and returns a new instance of a given class.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllocateObject
func NSAllocateObject(aClass objc.Class, extraBytes uint, zone *NSZone) objectivec.Object {
	if _nSAllocateObject == nil {
		panic("Foundation: symbol NSAllocateObject not loaded")
	}
	return _nSAllocateObject(aClass, extraBytes, zone)
}

var _nSClassFromString func(aClassName NSString) objc.Class

// NSClassFromString obtains a class by name.
//
// See: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName NSString) objc.Class {
	if _nSClassFromString == nil {
		panic("Foundation: symbol NSClassFromString not loaded")
	}
	return _nSClassFromString(aClassName)
}

var _nSCompareHashTables func(table1 *NSHashTable, table2 *NSHashTable) bool

// NSCompareHashTables returns a Boolean value that indicates whether the elements of two hash tables are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompareHashTables(_:_:)
func NSCompareHashTables(table1 *NSHashTable, table2 *NSHashTable) bool {
	if _nSCompareHashTables == nil {
		panic("Foundation: symbol NSCompareHashTables not loaded")
	}
	return _nSCompareHashTables(table1, table2)
}

var _nSCompareMapTables func(table1 *NSMapTable, table2 *NSMapTable) bool

// NSCompareMapTables compares the elements of two map tables for equality.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompareMapTables(_:_:)
func NSCompareMapTables(table1 *NSMapTable, table2 *NSMapTable) bool {
	if _nSCompareMapTables == nil {
		panic("Foundation: symbol NSCompareMapTables not loaded")
	}
	return _nSCompareMapTables(table1, table2)
}

// NSContainsRect returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// See: https://developer.apple.com/documentation/Foundation/NSContainsRect(_:_:)
func NSContainsRect(aRect corefoundation.CGRect, bRect corefoundation.CGRect) bool {
	return nsContainsRect(aRect, bRect)
}

var _nSCopyHashTableWithZone func(table *NSHashTable, zone *NSZone) *NSHashTable

// NSCopyHashTableWithZone performs a shallow copy of the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyHashTableWithZone(_:_:)
func NSCopyHashTableWithZone(table *NSHashTable, zone *NSZone) *NSHashTable {
	if _nSCopyHashTableWithZone == nil {
		panic("Foundation: symbol NSCopyHashTableWithZone not loaded")
	}
	return _nSCopyHashTableWithZone(table, zone)
}

var _nSCopyMapTableWithZone func(table *NSMapTable, zone *NSZone) *NSMapTable

// NSCopyMapTableWithZone performs a shallow copy of the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyMapTableWithZone(_:_:)
func NSCopyMapTableWithZone(table *NSMapTable, zone *NSZone) *NSMapTable {
	if _nSCopyMapTableWithZone == nil {
		panic("Foundation: symbol NSCopyMapTableWithZone not loaded")
	}
	return _nSCopyMapTableWithZone(table, zone)
}

var _nSCopyMemoryPages func(source unsafe.Pointer, dest unsafe.Pointer, bytes uint)

// NSCopyMemoryPages copies a block of memory.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyMemoryPages(_:_:_:)
func NSCopyMemoryPages(source unsafe.Pointer, dest unsafe.Pointer, bytes uint) {
	if _nSCopyMemoryPages == nil {
		panic("Foundation: symbol NSCopyMemoryPages not loaded")
	}
	_nSCopyMemoryPages(source, dest, bytes)
}

var _nSCopyObject func(object objectivec.Object, extraBytes uint, zone *NSZone) objectivec.Object

// NSCopyObject creates an exact copy of an object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyObject
func NSCopyObject(object objectivec.Object, extraBytes uint, zone *NSZone) objectivec.Object {
	if _nSCopyObject == nil {
		panic("Foundation: symbol NSCopyObject not loaded")
	}
	return _nSCopyObject(object, extraBytes, zone)
}

var _nSCountFrames func() uint

// NSCountFrames returns the number of call frames on the stack.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() uint {
	if _nSCountFrames == nil {
		panic("Foundation: symbol NSCountFrames not loaded")
	}
	return _nSCountFrames()
}

var _nSCountHashTable func(table *NSHashTable) uint

// NSCountHashTable returns the number of elements in a hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountHashTable(_:)
func NSCountHashTable(table *NSHashTable) uint {
	if _nSCountHashTable == nil {
		panic("Foundation: symbol NSCountHashTable not loaded")
	}
	return _nSCountHashTable(table)
}

var _nSCountMapTable func(table *NSMapTable) uint

// NSCountMapTable returns the number of elements in a map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountMapTable(_:)
func NSCountMapTable(table *NSMapTable) uint {
	if _nSCountMapTable == nil {
		panic("Foundation: symbol NSCountMapTable not loaded")
	}
	return _nSCountMapTable(table)
}

var _nSCreateHashTable func(callBacks NSHashTableCallBacks, capacity uint) *NSHashTable

// NSCreateHashTable creates and returns a new hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateHashTable(_:_:)
func NSCreateHashTable(callBacks NSHashTableCallBacks, capacity uint) *NSHashTable {
	if _nSCreateHashTable == nil {
		panic("Foundation: symbol NSCreateHashTable not loaded")
	}
	return _nSCreateHashTable(callBacks, capacity)
}

var _nSCreateHashTableWithZone func(callBacks NSHashTableCallBacks, capacity uint, zone *NSZone) *NSHashTable

// NSCreateHashTableWithZone creates a new hash table in a given zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateHashTableWithZone(_:_:_:)
func NSCreateHashTableWithZone(callBacks NSHashTableCallBacks, capacity uint, zone *NSZone) *NSHashTable {
	if _nSCreateHashTableWithZone == nil {
		panic("Foundation: symbol NSCreateHashTableWithZone not loaded")
	}
	return _nSCreateHashTableWithZone(callBacks, capacity, zone)
}

var _nSCreateMapTable func(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint) *NSMapTable

// NSCreateMapTable creates a new map table in the default zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateMapTable(_:_:_:)
func NSCreateMapTable(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint) *NSMapTable {
	if _nSCreateMapTable == nil {
		panic("Foundation: symbol NSCreateMapTable not loaded")
	}
	return _nSCreateMapTable(keyCallBacks, valueCallBacks, capacity)
}

var _nSCreateMapTableWithZone func(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint, zone *NSZone) *NSMapTable

// NSCreateMapTableWithZone creates a new map table in the specified zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateMapTableWithZone(_:_:_:_:)
func NSCreateMapTableWithZone(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint, zone *NSZone) *NSMapTable {
	if _nSCreateMapTableWithZone == nil {
		panic("Foundation: symbol NSCreateMapTableWithZone not loaded")
	}
	return _nSCreateMapTableWithZone(keyCallBacks, valueCallBacks, capacity, zone)
}

var _nSCreateZone func(startSize uint, granularity uint, canFree bool) *NSZone

// NSCreateZone creates a new zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateZone
func NSCreateZone(startSize uint, granularity uint, canFree bool) *NSZone {
	if _nSCreateZone == nil {
		panic("Foundation: symbol NSCreateZone not loaded")
	}
	return _nSCreateZone(startSize, granularity, canFree)
}

var _nSDeallocateMemoryPages func(ptr unsafe.Pointer, bytes uint)

// NSDeallocateMemoryPages deallocates the specified block of memory.
//
// See: https://developer.apple.com/documentation/Foundation/NSDeallocateMemoryPages(_:_:)
func NSDeallocateMemoryPages(ptr unsafe.Pointer, bytes uint) {
	if _nSDeallocateMemoryPages == nil {
		panic("Foundation: symbol NSDeallocateMemoryPages not loaded")
	}
	_nSDeallocateMemoryPages(ptr, bytes)
}

var _nSDeallocateObject func(object objectivec.Object)

// NSDeallocateObject destroys an existing object.
//
// See: https://developer.apple.com/documentation/Foundation/NSDeallocateObject
func NSDeallocateObject(object objectivec.Object) {
	if _nSDeallocateObject == nil {
		panic("Foundation: symbol NSDeallocateObject not loaded")
	}
	_nSDeallocateObject(object)
}

var _nSDecimalAdd func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalAdd adds two decimal values.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalAdd(_:_:_:_:)
func NSDecimalAdd(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalAdd == nil {
		panic("Foundation: symbol NSDecimalAdd not loaded")
	}
	return _nSDecimalAdd(result, leftOperand, rightOperand, roundingMode)
}

var _nSDecimalCompact func(number *NSDecimal)

// NSDecimalCompact compacts the decimal structure for efficiency.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number *NSDecimal) {
	if _nSDecimalCompact == nil {
		panic("Foundation: symbol NSDecimalCompact not loaded")
	}
	_nSDecimalCompact(number)
}

var _nSDecimalCompare func(leftOperand *NSDecimal, rightOperand *NSDecimal) NSComparisonResult

// NSDecimalCompare compares two decimal values.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalCompare(_:_:)
func NSDecimalCompare(leftOperand *NSDecimal, rightOperand *NSDecimal) NSComparisonResult {
	if _nSDecimalCompare == nil {
		panic("Foundation: symbol NSDecimalCompare not loaded")
	}
	return _nSDecimalCompare(leftOperand, rightOperand)
}

var _nSDecimalCopy func(destination *NSDecimal, source *NSDecimal)

// NSDecimalCopy copies the value of a decimal number.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalCopy(_:_:)
func NSDecimalCopy(destination *NSDecimal, source *NSDecimal) {
	if _nSDecimalCopy == nil {
		panic("Foundation: symbol NSDecimalCopy not loaded")
	}
	_nSDecimalCopy(destination, source)
}

var _nSDecimalDivide func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalDivide divides one decimal value by another.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalDivide(_:_:_:_:)
func NSDecimalDivide(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalDivide == nil {
		panic("Foundation: symbol NSDecimalDivide not loaded")
	}
	return _nSDecimalDivide(result, leftOperand, rightOperand, roundingMode)
}

var _nSDecimalMultiply func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalMultiply multiplies two decimal numbers together.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalMultiply(_:_:_:_:)
func NSDecimalMultiply(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalMultiply == nil {
		panic("Foundation: symbol NSDecimalMultiply not loaded")
	}
	return _nSDecimalMultiply(result, leftOperand, rightOperand, roundingMode)
}

var _nSDecimalMultiplyByPowerOf10 func(result *NSDecimal, number *NSDecimal, power int16, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalMultiplyByPowerOf10 multiplies a decimal by the specified power of 10.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalMultiplyByPowerOf10(_:_:_:_:)
func NSDecimalMultiplyByPowerOf10(result *NSDecimal, number *NSDecimal, power int16, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalMultiplyByPowerOf10 == nil {
		panic("Foundation: symbol NSDecimalMultiplyByPowerOf10 not loaded")
	}
	return _nSDecimalMultiplyByPowerOf10(result, number, power, roundingMode)
}

var _nSDecimalNormalize func(number1 *NSDecimal, number2 *NSDecimal, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalNormalize normalizes the internal format of two decimal numbers to simplify later operations.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNormalize(_:_:_:)
func NSDecimalNormalize(number1 *NSDecimal, number2 *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalNormalize == nil {
		panic("Foundation: symbol NSDecimalNormalize not loaded")
	}
	return _nSDecimalNormalize(number1, number2, roundingMode)
}

var _nSDecimalPower func(result *NSDecimal, number *NSDecimal, power uint, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalPower raises the decimal value to the specified power.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result *NSDecimal, number *NSDecimal, power uint, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalPower == nil {
		panic("Foundation: symbol NSDecimalPower not loaded")
	}
	return _nSDecimalPower(result, number, power, roundingMode)
}

var _nSDecimalRound func(result *NSDecimal, number *NSDecimal, scale int, roundingMode NSRoundingMode)

// NSDecimalRound rounds off the decimal value.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalRound(_:_:_:_:)
func NSDecimalRound(result *NSDecimal, number *NSDecimal, scale int, roundingMode NSRoundingMode) {
	if _nSDecimalRound == nil {
		panic("Foundation: symbol NSDecimalRound not loaded")
	}
	_nSDecimalRound(result, number, scale, roundingMode)
}

var _nSDecimalString func(dcm *NSDecimal, locale objectivec.Object) NSString

// NSDecimalString returns a string representation of the decimal value appropriate for the specified locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalString(_:_:)
func NSDecimalString(dcm *NSDecimal, locale objectivec.Object) NSString {
	if _nSDecimalString == nil {
		panic("Foundation: symbol NSDecimalString not loaded")
	}
	return _nSDecimalString(dcm, locale)
}

var _nSDecimalSubtract func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError

// NSDecimalSubtract subtracts one decimal value from another.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalSubtract(_:_:_:_:)
func NSDecimalSubtract(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	if _nSDecimalSubtract == nil {
		panic("Foundation: symbol NSDecimalSubtract not loaded")
	}
	return _nSDecimalSubtract(result, leftOperand, rightOperand, roundingMode)
}

var _nSDecrementExtraRefCountWasZero func(object objectivec.Object) bool

// NSDecrementExtraRefCountWasZero decrements the specified object’s reference count.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecrementExtraRefCountWasZero
func NSDecrementExtraRefCountWasZero(object objectivec.Object) bool {
	if _nSDecrementExtraRefCountWasZero == nil {
		panic("Foundation: symbol NSDecrementExtraRefCountWasZero not loaded")
	}
	return _nSDecrementExtraRefCountWasZero(object)
}

var _nSDefaultMallocZone func() *NSZone

// NSDefaultMallocZone returns the default zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSDefaultMallocZone
func NSDefaultMallocZone() *NSZone {
	if _nSDefaultMallocZone == nil {
		panic("Foundation: symbol NSDefaultMallocZone not loaded")
	}
	return _nSDefaultMallocZone()
}

// NSDivideRect divides a rectangle into two new rectangles.
//
// See: https://developer.apple.com/documentation/Foundation/NSDivideRect(_:_:_:_:_:)
func NSDivideRect(inRect corefoundation.CGRect, slice *corefoundation.CGRect, rem *corefoundation.CGRect, amount float64, edge NSRectEdge) {
	nsDivideRect(inRect, slice, rem, amount, edge)
}

// NSEdgeInsetsEqual returns a Boolean value that indicates whether two edge insets structures are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSEdgeInsetsEqual(_:_:)
func NSEdgeInsetsEqual(aInsets NSEdgeInsets, bInsets NSEdgeInsets) bool {
	return nsEdgeInsetsEqual(aInsets, bInsets)
}

var _nSEndHashTableEnumeration func(enumerator *NSHashEnumerator)

// NSEndHashTableEnumeration used when finished with an enumerator.
//
// See: https://developer.apple.com/documentation/Foundation/NSEndHashTableEnumeration(_:)
func NSEndHashTableEnumeration(enumerator *NSHashEnumerator) {
	if _nSEndHashTableEnumeration == nil {
		panic("Foundation: symbol NSEndHashTableEnumeration not loaded")
	}
	_nSEndHashTableEnumeration(enumerator)
}

var _nSEndMapTableEnumeration func(enumerator *NSMapEnumerator)

// NSEndMapTableEnumeration used when finished with an enumerator.
//
// See: https://developer.apple.com/documentation/Foundation/NSEndMapTableEnumeration(_:)
func NSEndMapTableEnumeration(enumerator *NSMapEnumerator) {
	if _nSEndMapTableEnumeration == nil {
		panic("Foundation: symbol NSEndMapTableEnumeration not loaded")
	}
	_nSEndMapTableEnumeration(enumerator)
}

var _nSEnumerateHashTable func(table *NSHashTable) NSHashEnumerator

// NSEnumerateHashTable creates an enumerator for the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerateHashTable(_:)
func NSEnumerateHashTable(table *NSHashTable) NSHashEnumerator {
	if _nSEnumerateHashTable == nil {
		panic("Foundation: symbol NSEnumerateHashTable not loaded")
	}
	return _nSEnumerateHashTable(table)
}

var _nSEnumerateMapTable func(table *NSMapTable) NSMapEnumerator

// NSEnumerateMapTable creates an enumerator for the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerateMapTable(_:)
func NSEnumerateMapTable(table *NSMapTable) NSMapEnumerator {
	if _nSEnumerateMapTable == nil {
		panic("Foundation: symbol NSEnumerateMapTable not loaded")
	}
	return _nSEnumerateMapTable(table)
}

// NSEqualPoints returns a Boolean value that indicates whether two points are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSEqualPoints(_:_:)
func NSEqualPoints(aPoint corefoundation.CGPoint, bPoint corefoundation.CGPoint) bool {
	return nsEqualPoints(aPoint, bPoint)
}

// NSEqualRects returns a Boolean value that indicates whether the two rectangles are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSEqualRects(_:_:)
func NSEqualRects(aRect corefoundation.CGRect, bRect corefoundation.CGRect) bool {
	return nsEqualRects(aRect, bRect)
}

// NSEqualSizes returns a Boolean that indicates whether two size values are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSEqualSizes(_:_:)
func NSEqualSizes(aSize corefoundation.CGSize, bSize corefoundation.CGSize) bool {
	return nsEqualSizes(aSize, bSize)
}

var _nSExtraRefCount func(object objectivec.Object) uint

// NSExtraRefCount returns the specified object’s reference count.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtraRefCount
func NSExtraRefCount(object objectivec.Object) uint {
	if _nSExtraRefCount == nil {
		panic("Foundation: symbol NSExtraRefCount not loaded")
	}
	return _nSExtraRefCount(object)
}

var _nSFileTypeForHFSTypeCode func(hfsFileTypeCode uint32) NSString

// NSFileTypeForHFSTypeCode returns a string encoding a file type code.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileTypeForHFSTypeCode(_:)
func NSFileTypeForHFSTypeCode(hfsFileTypeCode uint32) NSString {
	if _nSFileTypeForHFSTypeCode == nil {
		panic("Foundation: symbol NSFileTypeForHFSTypeCode not loaded")
	}
	return _nSFileTypeForHFSTypeCode(hfsFileTypeCode)
}

var _nSFrameAddress func(frame uint) unsafe.Pointer

// NSFrameAddress returns the value of the frame pointer of the specified frame.
//
// See: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame uint) unsafe.Pointer {
	if _nSFrameAddress == nil {
		panic("Foundation: symbol NSFrameAddress not loaded")
	}
	return _nSFrameAddress(frame)
}

var _nSFreeHashTable func(table *NSHashTable)

// NSFreeHashTable deletes the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSFreeHashTable(_:)
func NSFreeHashTable(table *NSHashTable) {
	if _nSFreeHashTable == nil {
		panic("Foundation: symbol NSFreeHashTable not loaded")
	}
	_nSFreeHashTable(table)
}

var _nSFreeMapTable func(table *NSMapTable)

// NSFreeMapTable deletes the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSFreeMapTable(_:)
func NSFreeMapTable(table *NSMapTable) {
	if _nSFreeMapTable == nil {
		panic("Foundation: symbol NSFreeMapTable not loaded")
	}
	_nSFreeMapTable(table)
}

var _nSFullUserName func() NSString

// NSFullUserName returns a string containing the full name of the current user.
//
// See: https://developer.apple.com/documentation/Foundation/NSFullUserName()
func NSFullUserName() NSString {
	if _nSFullUserName == nil {
		panic("Foundation: symbol NSFullUserName not loaded")
	}
	return _nSFullUserName()
}

var _nSGetSizeAndAlignment func(typePtr string, sizep *uint, alignp *uint) *byte

// NSGetSizeAndAlignment obtains the actual size and the aligned size of an encoded type.
//
// See: https://developer.apple.com/documentation/Foundation/NSGetSizeAndAlignment(_:_:_:)
func NSGetSizeAndAlignment(typePtr string, sizep *uint, alignp *uint) *byte {
	if _nSGetSizeAndAlignment == nil {
		panic("Foundation: symbol NSGetSizeAndAlignment not loaded")
	}
	return _nSGetSizeAndAlignment(typePtr, sizep, alignp)
}

var _nSGetUncaughtExceptionHandler func() func(unsafe.Pointer)

// NSGetUncaughtExceptionHandler returns the top-level error handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSGetUncaughtExceptionHandler()
func NSGetUncaughtExceptionHandler() func(unsafe.Pointer) {
	if _nSGetUncaughtExceptionHandler == nil {
		panic("Foundation: symbol NSGetUncaughtExceptionHandler not loaded")
	}
	return _nSGetUncaughtExceptionHandler()
}

var _nSHFSTypeCodeFromFileType func(fileTypeString NSString) uint32

// NSHFSTypeCodeFromFileType returns a file type code.
//
// See: https://developer.apple.com/documentation/Foundation/NSHFSTypeCodeFromFileType(_:)
func NSHFSTypeCodeFromFileType(fileTypeString NSString) uint32 {
	if _nSHFSTypeCodeFromFileType == nil {
		panic("Foundation: symbol NSHFSTypeCodeFromFileType not loaded")
	}
	return _nSHFSTypeCodeFromFileType(fileTypeString)
}

var _nSHFSTypeOfFile func(fullFilePath NSString) NSString

// NSHFSTypeOfFile returns a string encoding a file type.
//
// See: https://developer.apple.com/documentation/Foundation/NSHFSTypeOfFile(_:)
func NSHFSTypeOfFile(fullFilePath NSString) NSString {
	if _nSHFSTypeOfFile == nil {
		panic("Foundation: symbol NSHFSTypeOfFile not loaded")
	}
	return _nSHFSTypeOfFile(fullFilePath)
}

var _nSHashGet func(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer

// NSHashGet returns an element of the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashGet(_:_:)
func NSHashGet(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer {
	if _nSHashGet == nil {
		panic("Foundation: symbol NSHashGet not loaded")
	}
	return _nSHashGet(table, pointer)
}

var _nSHashInsert func(table *NSHashTable, pointer unsafe.Pointer)

// NSHashInsert adds an element to the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashInsert(_:_:)
func NSHashInsert(table *NSHashTable, pointer unsafe.Pointer) {
	if _nSHashInsert == nil {
		panic("Foundation: symbol NSHashInsert not loaded")
	}
	_nSHashInsert(table, pointer)
}

var _nSHashInsertIfAbsent func(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer

// NSHashInsertIfAbsent adds an element to the specified hash table only if the table does not already contain the element.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashInsertIfAbsent(_:_:)
func NSHashInsertIfAbsent(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer {
	if _nSHashInsertIfAbsent == nil {
		panic("Foundation: symbol NSHashInsertIfAbsent not loaded")
	}
	return _nSHashInsertIfAbsent(table, pointer)
}

var _nSHashInsertKnownAbsent func(table *NSHashTable, pointer unsafe.Pointer)

// NSHashInsertKnownAbsent adds an element to the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashInsertKnownAbsent(_:_:)
func NSHashInsertKnownAbsent(table *NSHashTable, pointer unsafe.Pointer) {
	if _nSHashInsertKnownAbsent == nil {
		panic("Foundation: symbol NSHashInsertKnownAbsent not loaded")
	}
	_nSHashInsertKnownAbsent(table, pointer)
}

var _nSHashRemove func(table *NSHashTable, pointer unsafe.Pointer)

// NSHashRemove removes an element from the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashRemove(_:_:)
func NSHashRemove(table *NSHashTable, pointer unsafe.Pointer) {
	if _nSHashRemove == nil {
		panic("Foundation: symbol NSHashRemove not loaded")
	}
	_nSHashRemove(table, pointer)
}

var _nSHomeDirectory func() NSString

// NSHomeDirectory returns the path to either the user’s or application’s home directory, depending on the platform.
//
// See: https://developer.apple.com/documentation/Foundation/NSHomeDirectory()
func NSHomeDirectory() NSString {
	if _nSHomeDirectory == nil {
		panic("Foundation: symbol NSHomeDirectory not loaded")
	}
	return _nSHomeDirectory()
}

var _nSHomeDirectoryForUser func(userName NSString) NSString

// NSHomeDirectoryForUser returns the path to a given user’s home directory.
//
// See: https://developer.apple.com/documentation/Foundation/NSHomeDirectoryForUser(_:)
func NSHomeDirectoryForUser(userName NSString) NSString {
	if _nSHomeDirectoryForUser == nil {
		panic("Foundation: symbol NSHomeDirectoryForUser not loaded")
	}
	return _nSHomeDirectoryForUser(userName)
}

var _nSIncrementExtraRefCount func(object objectivec.Object)

// NSIncrementExtraRefCount increments the specified object’s reference count.
//
// See: https://developer.apple.com/documentation/Foundation/NSIncrementExtraRefCount
func NSIncrementExtraRefCount(object objectivec.Object) {
	if _nSIncrementExtraRefCount == nil {
		panic("Foundation: symbol NSIncrementExtraRefCount not loaded")
	}
	_nSIncrementExtraRefCount(object)
}

// NSInsetRect insets a rectangle by a specified amount.
//
// See: https://developer.apple.com/documentation/Foundation/NSInsetRect(_:_:_:)
func NSInsetRect(aRect corefoundation.CGRect, dX float64, dY float64) corefoundation.CGRect {
	return nsInsetRect(aRect, dX, dY)
}

// NSIntegralRect adjusts the sides of a rectangle to integer values.
//
// See: https://developer.apple.com/documentation/Foundation/NSIntegralRect(_:)
func NSIntegralRect(aRect corefoundation.CGRect) corefoundation.CGRect {
	return nsIntegralRect(aRect)
}

// NSIntegralRectWithOptions adjusts the sides of a rectangle to integral values using the specified options.
//
// See: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func NSIntegralRectWithOptions(aRect corefoundation.CGRect, opts NSAlignmentOptions) corefoundation.CGRect {
	return nsIntegralRectWithOptions(aRect, opts)
}

// NSIntersectionRange returns the intersection of the specified ranges.
//
// See: https://developer.apple.com/documentation/Foundation/NSIntersectionRange(_:_:)
func NSIntersectionRange(range1 NSRange, range2 NSRange) NSRange {
	return nsIntersectionRange(range1, range2)
}

// NSIntersectionRect calculates the intersection of two rectangles.
//
// See: https://developer.apple.com/documentation/Foundation/NSIntersectionRect(_:_:)
func NSIntersectionRect(aRect corefoundation.CGRect, bRect corefoundation.CGRect) corefoundation.CGRect {
	return nsIntersectionRect(aRect, bRect)
}

// NSIntersectsRect returns a Boolean value that indicates whether two rectangles intersect.
//
// See: https://developer.apple.com/documentation/Foundation/NSIntersectsRect(_:_:)
func NSIntersectsRect(aRect corefoundation.CGRect, bRect corefoundation.CGRect) bool {
	return nsIntersectsRect(aRect, bRect)
}

// NSIsEmptyRect returns a Boolean value that indicates whether a given rectangle is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSIsEmptyRect(_:)
func NSIsEmptyRect(aRect corefoundation.CGRect) bool {
	return nsIsEmptyRect(aRect)
}

var _nSIsFreedObject func(anObject objectivec.Object) bool

// NSIsFreedObject returns a Boolean indicating whether the specified object has been freed.
//
// See: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject objectivec.Object) bool {
	if _nSIsFreedObject == nil {
		panic("Foundation: symbol NSIsFreedObject not loaded")
	}
	return _nSIsFreedObject(anObject)
}

var _nSLog func(format NSString)

// NSLog logs an error message to the Apple System Log facility.
//
// See: https://developer.apple.com/documentation/Foundation/NSLog
func NSLog(format NSString) {
	if _nSLog == nil {
		panic("Foundation: symbol NSLog not loaded")
	}
	_nSLog(format)
}

var _nSLogPageSize func() uint

// NSLogPageSize returns the binary log of the page size.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogPageSize()
func NSLogPageSize() uint {
	if _nSLogPageSize == nil {
		panic("Foundation: symbol NSLogPageSize not loaded")
	}
	return _nSLogPageSize()
}

var _nSLogv func(format NSString, args kernel.Va_list)

// NSLogv logs an error message to the Apple System Log facility.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogv(_:_:)
func NSLogv(format NSString, args kernel.Va_list) {
	if _nSLogv == nil {
		panic("Foundation: symbol NSLogv not loaded")
	}
	_nSLogv(format, args)
}

var _nSMapGet func(table *NSMapTable, key unsafe.Pointer) unsafe.Pointer

// NSMapGet returns a map table value for the specified key.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapGet(_:_:)
func NSMapGet(table *NSMapTable, key unsafe.Pointer) unsafe.Pointer {
	if _nSMapGet == nil {
		panic("Foundation: symbol NSMapGet not loaded")
	}
	return _nSMapGet(table, key)
}

var _nSMapInsert func(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer)

// NSMapInsert inserts a key-value pair into the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapInsert(_:_:_:)
func NSMapInsert(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) {
	if _nSMapInsert == nil {
		panic("Foundation: symbol NSMapInsert not loaded")
	}
	_nSMapInsert(table, key, value)
}

var _nSMapInsertIfAbsent func(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer

// NSMapInsertIfAbsent inserts a key-value pair into the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapInsertIfAbsent(_:_:_:)
func NSMapInsertIfAbsent(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	if _nSMapInsertIfAbsent == nil {
		panic("Foundation: symbol NSMapInsertIfAbsent not loaded")
	}
	return _nSMapInsertIfAbsent(table, key, value)
}

var _nSMapInsertKnownAbsent func(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer)

// NSMapInsertKnownAbsent inserts a key-value pair into the specified table if the pair had not been previously added.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapInsertKnownAbsent(_:_:_:)
func NSMapInsertKnownAbsent(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) {
	if _nSMapInsertKnownAbsent == nil {
		panic("Foundation: symbol NSMapInsertKnownAbsent not loaded")
	}
	_nSMapInsertKnownAbsent(table, key, value)
}

var _nSMapMember func(table *NSMapTable, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) bool

// NSMapMember indicates whether a given table contains a given key.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapMember(_:_:_:_:)
func NSMapMember(table *NSMapTable, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) bool {
	if _nSMapMember == nil {
		panic("Foundation: symbol NSMapMember not loaded")
	}
	return _nSMapMember(table, key, originalKey, value)
}

var _nSMapRemove func(table *NSMapTable, key unsafe.Pointer)

// NSMapRemove removes a key and corresponding value from the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapRemove(_:_:)
func NSMapRemove(table *NSMapTable, key unsafe.Pointer) {
	if _nSMapRemove == nil {
		panic("Foundation: symbol NSMapRemove not loaded")
	}
	_nSMapRemove(table, key)
}

// NSMouseInRect returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSMouseInRect(_:_:_:)
func NSMouseInRect(aPoint corefoundation.CGPoint, aRect corefoundation.CGRect, flipped bool) bool {
	return nsMouseInRect(aPoint, aRect, flipped)
}

var _nSNextHashEnumeratorItem func(enumerator *NSHashEnumerator) unsafe.Pointer

// NSNextHashEnumeratorItem returns the next hash-table element in the enumeration.
//
// See: https://developer.apple.com/documentation/Foundation/NSNextHashEnumeratorItem(_:)
func NSNextHashEnumeratorItem(enumerator *NSHashEnumerator) unsafe.Pointer {
	if _nSNextHashEnumeratorItem == nil {
		panic("Foundation: symbol NSNextHashEnumeratorItem not loaded")
	}
	return _nSNextHashEnumeratorItem(enumerator)
}

var _nSNextMapEnumeratorPair func(enumerator *NSMapEnumerator, key unsafe.Pointer, value unsafe.Pointer) bool

// NSNextMapEnumeratorPair returns a Boolean value that indicates whether the next map-table pair in the enumeration are set.
//
// See: https://developer.apple.com/documentation/Foundation/NSNextMapEnumeratorPair(_:_:_:)
func NSNextMapEnumeratorPair(enumerator *NSMapEnumerator, key unsafe.Pointer, value unsafe.Pointer) bool {
	if _nSNextMapEnumeratorPair == nil {
		panic("Foundation: symbol NSNextMapEnumeratorPair not loaded")
	}
	return _nSNextMapEnumeratorPair(enumerator, key, value)
}

// NSOffsetRect offsets the rectangle by the specified amount.
//
// See: https://developer.apple.com/documentation/Foundation/NSOffsetRect(_:_:_:)
func NSOffsetRect(aRect corefoundation.CGRect, dX float64, dY float64) corefoundation.CGRect {
	return nsOffsetRect(aRect, dX, dY)
}

var _nSOpenStepRootDirectory func() NSString

// NSOpenStepRootDirectory returns the root directory of the user’s system.
//
// See: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() NSString {
	if _nSOpenStepRootDirectory == nil {
		panic("Foundation: symbol NSOpenStepRootDirectory not loaded")
	}
	return _nSOpenStepRootDirectory()
}

var _nSPageSize func() uint

// NSPageSize returns the number of bytes in a page.
//
// See: https://developer.apple.com/documentation/Foundation/NSPageSize()
func NSPageSize() uint {
	if _nSPageSize == nil {
		panic("Foundation: symbol NSPageSize not loaded")
	}
	return _nSPageSize()
}

var _nSPointFromString func(aString NSString) corefoundation.CGPoint

// NSPointFromString returns a point from a text-based representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointFromString(_:)
func NSPointFromString(aString NSString) corefoundation.CGPoint {
	if _nSPointFromString == nil {
		panic("Foundation: symbol NSPointFromString not loaded")
	}
	return _nSPointFromString(aString)
}

// NSPointInRect returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointInRect(_:_:)
func NSPointInRect(aPoint corefoundation.CGPoint, aRect corefoundation.CGRect) bool {
	return nsPointInRect(aPoint, aRect)
}

var _nSProtocolFromString func(namestr NSString) **objectivec.Protocol

// NSProtocolFromString returns a the protocol with a given name.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolFromString(_:)
func NSProtocolFromString(namestr NSString) **objectivec.Protocol {
	if _nSProtocolFromString == nil {
		panic("Foundation: symbol NSProtocolFromString not loaded")
	}
	return _nSProtocolFromString(namestr)
}

var _nSRangeFromString func(aString NSString) NSRange

// NSRangeFromString returns a range from a textual representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeFromString(_:)
func NSRangeFromString(aString NSString) NSRange {
	if _nSRangeFromString == nil {
		panic("Foundation: symbol NSRangeFromString not loaded")
	}
	return _nSRangeFromString(aString)
}

var _nSReallocateCollectable func(ptr unsafe.Pointer, size uint, options uint) unsafe.Pointer

// NSReallocateCollectable reallocates collectable memory.
//
// Deprecated: Garbage collection is deprecated in OS X v10.8; instead,you should use AutomaticReference Counting—see [Transitioning to ARC Release Notes](<https://developer.apple.com/library/archive/releasenotes/ObjectiveC/RN-TransitioningToARC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011226>).
//
// See: https://developer.apple.com/documentation/Foundation/NSReallocateCollectable
func NSReallocateCollectable(ptr unsafe.Pointer, size uint, options uint) unsafe.Pointer {
	if _nSReallocateCollectable == nil {
		panic("Foundation: symbol NSReallocateCollectable not loaded")
	}
	return _nSReallocateCollectable(ptr, size, options)
}

var _nSRecordAllocationEvent func(eventType int, object objectivec.Object)

// NSRecordAllocationEvent notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// See: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType int, object objectivec.Object) {
	if _nSRecordAllocationEvent == nil {
		panic("Foundation: symbol NSRecordAllocationEvent not loaded")
	}
	_nSRecordAllocationEvent(eventType, object)
}

var _nSRectFromString func(aString NSString) corefoundation.CGRect

// NSRectFromString returns a rectangle from a text-based representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSRectFromString(_:)
func NSRectFromString(aString NSString) corefoundation.CGRect {
	if _nSRectFromString == nil {
		panic("Foundation: symbol NSRectFromString not loaded")
	}
	return _nSRectFromString(aString)
}

var _nSRecycleZone func(zone *NSZone)

// NSRecycleZone frees memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSRecycleZone
func NSRecycleZone(zone *NSZone) {
	if _nSRecycleZone == nil {
		panic("Foundation: symbol NSRecycleZone not loaded")
	}
	_nSRecycleZone(zone)
}

var _nSResetHashTable func(table *NSHashTable)

// NSResetHashTable deletes the elements of the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSResetHashTable(_:)
func NSResetHashTable(table *NSHashTable) {
	if _nSResetHashTable == nil {
		panic("Foundation: symbol NSResetHashTable not loaded")
	}
	_nSResetHashTable(table)
}

var _nSResetMapTable func(table *NSMapTable)

// NSResetMapTable deletes the elements of the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSResetMapTable(_:)
func NSResetMapTable(table *NSMapTable) {
	if _nSResetMapTable == nil {
		panic("Foundation: symbol NSResetMapTable not loaded")
	}
	_nSResetMapTable(table)
}

var _nSReturnAddress func(frame uint) unsafe.Pointer

// NSReturnAddress returns the value of the return address of the specified frame.
//
// See: https://developer.apple.com/documentation/Foundation/NSReturnAddress
func NSReturnAddress(frame uint) unsafe.Pointer {
	if _nSReturnAddress == nil {
		panic("Foundation: symbol NSReturnAddress not loaded")
	}
	return _nSReturnAddress(frame)
}

var _nSRoundDownToMultipleOfPageSize func(bytes uint) uint

// NSRoundDownToMultipleOfPageSize returns the specified number of bytes rounded down to a multiple of the page size.
//
// See: https://developer.apple.com/documentation/Foundation/NSRoundDownToMultipleOfPageSize(_:)
func NSRoundDownToMultipleOfPageSize(bytes uint) uint {
	if _nSRoundDownToMultipleOfPageSize == nil {
		panic("Foundation: symbol NSRoundDownToMultipleOfPageSize not loaded")
	}
	return _nSRoundDownToMultipleOfPageSize(bytes)
}

var _nSRoundUpToMultipleOfPageSize func(bytes uint) uint

// NSRoundUpToMultipleOfPageSize returns the specified number of bytes rounded up to a multiple of the page size.
//
// See: https://developer.apple.com/documentation/Foundation/NSRoundUpToMultipleOfPageSize(_:)
func NSRoundUpToMultipleOfPageSize(bytes uint) uint {
	if _nSRoundUpToMultipleOfPageSize == nil {
		panic("Foundation: symbol NSRoundUpToMultipleOfPageSize not loaded")
	}
	return _nSRoundUpToMultipleOfPageSize(bytes)
}

var _nSSearchPathForDirectoriesInDomains func(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, expandTilde bool) []NSString

// NSSearchPathForDirectoriesInDomains creates a list of directory search paths.
//
// See: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, expandTilde bool) []NSString {
	if _nSSearchPathForDirectoriesInDomains == nil {
		panic("Foundation: symbol NSSearchPathForDirectoriesInDomains not loaded")
	}
	return _nSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
}

var _nSSelectorFromString func(aSelectorName NSString) objectivec.SEL

// NSSelectorFromString returns the selector with a given name.
//
// See: https://developer.apple.com/documentation/Foundation/NSSelectorFromString(_:)
func NSSelectorFromString(aSelectorName NSString) objectivec.SEL {
	if _nSSelectorFromString == nil {
		panic("Foundation: symbol NSSelectorFromString not loaded")
	}
	return _nSSelectorFromString(aSelectorName)
}

var _nSSetUncaughtExceptionHandler func(arg0 func(*NSException))

// NSSetUncaughtExceptionHandler changes the top-level error handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetUncaughtExceptionHandler(_:)
func NSSetUncaughtExceptionHandler(arg0 func(*NSException)) {
	if _nSSetUncaughtExceptionHandler == nil {
		panic("Foundation: symbol NSSetUncaughtExceptionHandler not loaded")
	}
	_nSSetUncaughtExceptionHandler(arg0)
}

var _nSSetZoneName func(zone *NSZone, name NSString)

// NSSetZoneName sets the name of the specified zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetZoneName
func NSSetZoneName(zone *NSZone, name NSString) {
	if _nSSetZoneName == nil {
		panic("Foundation: symbol NSSetZoneName not loaded")
	}
	_nSSetZoneName(zone, name)
}

var _nSShouldRetainWithZone func(anObject objectivec.Object, requestedZone *NSZone) bool

// NSShouldRetainWithZone indicates whether an object should be retained.
//
// See: https://developer.apple.com/documentation/Foundation/NSShouldRetainWithZone
func NSShouldRetainWithZone(anObject objectivec.Object, requestedZone *NSZone) bool {
	if _nSShouldRetainWithZone == nil {
		panic("Foundation: symbol NSShouldRetainWithZone not loaded")
	}
	return _nSShouldRetainWithZone(anObject, requestedZone)
}

var _nSSizeFromString func(aString NSString) corefoundation.CGSize

// NSSizeFromString returns an [NSSize] from a text-based representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString NSString) corefoundation.CGSize {
	if _nSSizeFromString == nil {
		panic("Foundation: symbol NSSizeFromString not loaded")
	}
	return _nSSizeFromString(aString)
}

var _nSStringFromClass func(aClass objc.Class) NSString

// NSStringFromClass returns the name of a class as a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromClass(_:)
func NSStringFromClass(aClass objc.Class) NSString {
	if _nSStringFromClass == nil {
		panic("Foundation: symbol NSStringFromClass not loaded")
	}
	return _nSStringFromClass(aClass)
}

var _nSStringFromHashTable func(table *NSHashTable) NSString

// NSStringFromHashTable returns a string describing the hash table’s contents.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromHashTable(_:)
func NSStringFromHashTable(table *NSHashTable) NSString {
	if _nSStringFromHashTable == nil {
		panic("Foundation: symbol NSStringFromHashTable not loaded")
	}
	return _nSStringFromHashTable(table)
}

var _nSStringFromMapTable func(table *NSMapTable) NSString

// NSStringFromMapTable returns a string describing the map table’s contents.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromMapTable(_:)
func NSStringFromMapTable(table *NSMapTable) NSString {
	if _nSStringFromMapTable == nil {
		panic("Foundation: symbol NSStringFromMapTable not loaded")
	}
	return _nSStringFromMapTable(table)
}

var _nSStringFromPoint func(aPoint corefoundation.CGPoint) NSString

// NSStringFromPoint returns a string representation of a point.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromPoint(_:)
func NSStringFromPoint(aPoint corefoundation.CGPoint) NSString {
	if _nSStringFromPoint == nil {
		panic("Foundation: symbol NSStringFromPoint not loaded")
	}
	return _nSStringFromPoint(aPoint)
}

var _nSStringFromProtocol func(proto **objectivec.Protocol) NSString

// NSStringFromProtocol returns the name of a protocol as a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromProtocol(_:)
func NSStringFromProtocol(proto **objectivec.Protocol) NSString {
	if _nSStringFromProtocol == nil {
		panic("Foundation: symbol NSStringFromProtocol not loaded")
	}
	return _nSStringFromProtocol(proto)
}

var _nSStringFromRange func(range_ NSRange) NSString

// NSStringFromRange returns a string representation of a range.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromRange(_:)
func NSStringFromRange(range_ NSRange) NSString {
	if _nSStringFromRange == nil {
		panic("Foundation: symbol NSStringFromRange not loaded")
	}
	return _nSStringFromRange(range_)
}

var _nSStringFromRect func(aRect corefoundation.CGRect) NSString

// NSStringFromRect returns a string representation of a rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromRect(_:)
func NSStringFromRect(aRect corefoundation.CGRect) NSString {
	if _nSStringFromRect == nil {
		panic("Foundation: symbol NSStringFromRect not loaded")
	}
	return _nSStringFromRect(aRect)
}

var _nSStringFromSelector func(aSelector objectivec.SEL) NSString

// NSStringFromSelector returns a string representation of a given selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromSelector(_:)
func NSStringFromSelector(aSelector objectivec.SEL) NSString {
	if _nSStringFromSelector == nil {
		panic("Foundation: symbol NSStringFromSelector not loaded")
	}
	return _nSStringFromSelector(aSelector)
}

var _nSStringFromSize func(aSize corefoundation.CGSize) NSString

// NSStringFromSize returns a string representation of a size.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromSize(_:)
func NSStringFromSize(aSize corefoundation.CGSize) NSString {
	if _nSStringFromSize == nil {
		panic("Foundation: symbol NSStringFromSize not loaded")
	}
	return _nSStringFromSize(aSize)
}

var _nSTemporaryDirectory func() NSString

// NSTemporaryDirectory returns the path of the temporary directory for the current user.
//
// See: https://developer.apple.com/documentation/Foundation/NSTemporaryDirectory()
func NSTemporaryDirectory() NSString {
	if _nSTemporaryDirectory == nil {
		panic("Foundation: symbol NSTemporaryDirectory not loaded")
	}
	return _nSTemporaryDirectory()
}

// NSUnionRange returns the union of the specified ranges.
//
// See: https://developer.apple.com/documentation/Foundation/NSUnionRange(_:_:)
func NSUnionRange(range1 NSRange, range2 NSRange) NSRange {
	return nsUnionRange(range1, range2)
}

// NSUnionRect calculates the union of two rectangles.
//
// See: https://developer.apple.com/documentation/Foundation/NSUnionRect(_:_:)
func NSUnionRect(aRect corefoundation.CGRect, bRect corefoundation.CGRect) corefoundation.CGRect {
	return nsUnionRect(aRect, bRect)
}

var _nSUserName func() NSString

// NSUserName returns the logon name of the current user.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserName()
func NSUserName() NSString {
	if _nSUserName == nil {
		panic("Foundation: symbol NSUserName not loaded")
	}
	return _nSUserName()
}

var _nSZoneCalloc func(zone *NSZone, numElems uint, byteSize uint) unsafe.Pointer

// NSZoneCalloc allocates memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneCalloc
func NSZoneCalloc(zone *NSZone, numElems uint, byteSize uint) unsafe.Pointer {
	if _nSZoneCalloc == nil {
		panic("Foundation: symbol NSZoneCalloc not loaded")
	}
	return _nSZoneCalloc(zone, numElems, byteSize)
}

var _nSZoneFree func(zone *NSZone, ptr unsafe.Pointer)

// NSZoneFree deallocates a block of memory in the specified zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneFree
func NSZoneFree(zone *NSZone, ptr unsafe.Pointer) {
	if _nSZoneFree == nil {
		panic("Foundation: symbol NSZoneFree not loaded")
	}
	_nSZoneFree(zone, ptr)
}

var _nSZoneFromPointer func(ptr unsafe.Pointer) *NSZone

// NSZoneFromPointer gets the zone for a given block of memory.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneFromPointer
func NSZoneFromPointer(ptr unsafe.Pointer) *NSZone {
	if _nSZoneFromPointer == nil {
		panic("Foundation: symbol NSZoneFromPointer not loaded")
	}
	return _nSZoneFromPointer(ptr)
}

var _nSZoneMalloc func(zone *NSZone, size uint) unsafe.Pointer

// NSZoneMalloc allocates memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneMalloc
func NSZoneMalloc(zone *NSZone, size uint) unsafe.Pointer {
	if _nSZoneMalloc == nil {
		panic("Foundation: symbol NSZoneMalloc not loaded")
	}
	return _nSZoneMalloc(zone, size)
}

var _nSZoneName func(zone *NSZone) NSString

// NSZoneName returns the name of the specified zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneName
func NSZoneName(zone *NSZone) NSString {
	if _nSZoneName == nil {
		panic("Foundation: symbol NSZoneName not loaded")
	}
	return _nSZoneName(zone)
}

var _nSZoneRealloc func(zone *NSZone, ptr unsafe.Pointer, size uint) unsafe.Pointer

// NSZoneRealloc allocates memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneRealloc
func NSZoneRealloc(zone *NSZone, ptr unsafe.Pointer, size uint) unsafe.Pointer {
	if _nSZoneRealloc == nil {
		panic("Foundation: symbol NSZoneRealloc not loaded")
	}
	return _nSZoneRealloc(zone, ptr, size)
}

var _nXReadNSObjectFromCoder func(decoder *NSCoder) *objectivec.Object

// NXReadNSObjectFromCoder returns the next object from the coder.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/Foundation/NXReadNSObjectFromCoder
func NXReadNSObjectFromCoder(decoder *NSCoder) *objectivec.Object {
	if _nXReadNSObjectFromCoder == nil {
		panic("Foundation: symbol NXReadNSObjectFromCoder not loaded")
	}
	return _nXReadNSObjectFromCoder(decoder)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_nSAllHashTableObjects, frameworkHandle, "NSAllHashTableObjects")
	registerFunc(&_nSAllMapTableKeys, frameworkHandle, "NSAllMapTableKeys")
	registerFunc(&_nSAllMapTableValues, frameworkHandle, "NSAllMapTableValues")
	registerFunc(&_nSAllocateCollectable, frameworkHandle, "NSAllocateCollectable")
	registerFunc(&_nSAllocateMemoryPages, frameworkHandle, "NSAllocateMemoryPages")
	registerFunc(&_nSAllocateObject, frameworkHandle, "NSAllocateObject")
	registerFunc(&_nSClassFromString, frameworkHandle, "NSClassFromString")
	registerFunc(&_nSCompareHashTables, frameworkHandle, "NSCompareHashTables")
	registerFunc(&_nSCompareMapTables, frameworkHandle, "NSCompareMapTables")
	registerFunc(&_nSCopyHashTableWithZone, frameworkHandle, "NSCopyHashTableWithZone")
	registerFunc(&_nSCopyMapTableWithZone, frameworkHandle, "NSCopyMapTableWithZone")
	registerFunc(&_nSCopyMemoryPages, frameworkHandle, "NSCopyMemoryPages")
	registerFunc(&_nSCopyObject, frameworkHandle, "NSCopyObject")
	registerFunc(&_nSCountFrames, frameworkHandle, "NSCountFrames")
	registerFunc(&_nSCountHashTable, frameworkHandle, "NSCountHashTable")
	registerFunc(&_nSCountMapTable, frameworkHandle, "NSCountMapTable")
	registerFunc(&_nSCreateHashTable, frameworkHandle, "NSCreateHashTable")
	registerFunc(&_nSCreateHashTableWithZone, frameworkHandle, "NSCreateHashTableWithZone")
	registerFunc(&_nSCreateMapTable, frameworkHandle, "NSCreateMapTable")
	registerFunc(&_nSCreateMapTableWithZone, frameworkHandle, "NSCreateMapTableWithZone")
	registerFunc(&_nSCreateZone, frameworkHandle, "NSCreateZone")
	registerFunc(&_nSDeallocateMemoryPages, frameworkHandle, "NSDeallocateMemoryPages")
	registerFunc(&_nSDeallocateObject, frameworkHandle, "NSDeallocateObject")
	registerFunc(&_nSDecimalAdd, frameworkHandle, "NSDecimalAdd")
	registerFunc(&_nSDecimalCompact, frameworkHandle, "NSDecimalCompact")
	registerFunc(&_nSDecimalCompare, frameworkHandle, "NSDecimalCompare")
	registerFunc(&_nSDecimalCopy, frameworkHandle, "NSDecimalCopy")
	registerFunc(&_nSDecimalDivide, frameworkHandle, "NSDecimalDivide")
	registerFunc(&_nSDecimalMultiply, frameworkHandle, "NSDecimalMultiply")
	registerFunc(&_nSDecimalMultiplyByPowerOf10, frameworkHandle, "NSDecimalMultiplyByPowerOf10")
	registerFunc(&_nSDecimalNormalize, frameworkHandle, "NSDecimalNormalize")
	registerFunc(&_nSDecimalPower, frameworkHandle, "NSDecimalPower")
	registerFunc(&_nSDecimalRound, frameworkHandle, "NSDecimalRound")
	registerFunc(&_nSDecimalString, frameworkHandle, "NSDecimalString")
	registerFunc(&_nSDecimalSubtract, frameworkHandle, "NSDecimalSubtract")
	registerFunc(&_nSDecrementExtraRefCountWasZero, frameworkHandle, "NSDecrementExtraRefCountWasZero")
	registerFunc(&_nSDefaultMallocZone, frameworkHandle, "NSDefaultMallocZone")
	registerFunc(&_nSEndHashTableEnumeration, frameworkHandle, "NSEndHashTableEnumeration")
	registerFunc(&_nSEndMapTableEnumeration, frameworkHandle, "NSEndMapTableEnumeration")
	registerFunc(&_nSEnumerateHashTable, frameworkHandle, "NSEnumerateHashTable")
	registerFunc(&_nSEnumerateMapTable, frameworkHandle, "NSEnumerateMapTable")
	registerFunc(&_nSExtraRefCount, frameworkHandle, "NSExtraRefCount")
	registerFunc(&_nSFileTypeForHFSTypeCode, frameworkHandle, "NSFileTypeForHFSTypeCode")
	registerFunc(&_nSFrameAddress, frameworkHandle, "NSFrameAddress")
	registerFunc(&_nSFreeHashTable, frameworkHandle, "NSFreeHashTable")
	registerFunc(&_nSFreeMapTable, frameworkHandle, "NSFreeMapTable")
	registerFunc(&_nSFullUserName, frameworkHandle, "NSFullUserName")
	registerFunc(&_nSGetSizeAndAlignment, frameworkHandle, "NSGetSizeAndAlignment")
	registerFunc(&_nSGetUncaughtExceptionHandler, frameworkHandle, "NSGetUncaughtExceptionHandler")
	registerFunc(&_nSHFSTypeCodeFromFileType, frameworkHandle, "NSHFSTypeCodeFromFileType")
	registerFunc(&_nSHFSTypeOfFile, frameworkHandle, "NSHFSTypeOfFile")
	registerFunc(&_nSHashGet, frameworkHandle, "NSHashGet")
	registerFunc(&_nSHashInsert, frameworkHandle, "NSHashInsert")
	registerFunc(&_nSHashInsertIfAbsent, frameworkHandle, "NSHashInsertIfAbsent")
	registerFunc(&_nSHashInsertKnownAbsent, frameworkHandle, "NSHashInsertKnownAbsent")
	registerFunc(&_nSHashRemove, frameworkHandle, "NSHashRemove")
	registerFunc(&_nSHomeDirectory, frameworkHandle, "NSHomeDirectory")
	registerFunc(&_nSHomeDirectoryForUser, frameworkHandle, "NSHomeDirectoryForUser")
	registerFunc(&_nSIncrementExtraRefCount, frameworkHandle, "NSIncrementExtraRefCount")
	registerFunc(&_nSIsFreedObject, frameworkHandle, "NSIsFreedObject")
	registerFunc(&_nSLog, frameworkHandle, "NSLog")
	registerFunc(&_nSLogPageSize, frameworkHandle, "NSLogPageSize")
	registerFunc(&_nSLogv, frameworkHandle, "NSLogv")
	registerFunc(&_nSMapGet, frameworkHandle, "NSMapGet")
	registerFunc(&_nSMapInsert, frameworkHandle, "NSMapInsert")
	registerFunc(&_nSMapInsertIfAbsent, frameworkHandle, "NSMapInsertIfAbsent")
	registerFunc(&_nSMapInsertKnownAbsent, frameworkHandle, "NSMapInsertKnownAbsent")
	registerFunc(&_nSMapMember, frameworkHandle, "NSMapMember")
	registerFunc(&_nSMapRemove, frameworkHandle, "NSMapRemove")
	registerFunc(&_nSNextHashEnumeratorItem, frameworkHandle, "NSNextHashEnumeratorItem")
	registerFunc(&_nSNextMapEnumeratorPair, frameworkHandle, "NSNextMapEnumeratorPair")
	registerFunc(&_nSOpenStepRootDirectory, frameworkHandle, "NSOpenStepRootDirectory")
	registerFunc(&_nSPageSize, frameworkHandle, "NSPageSize")
	registerFunc(&_nSPointFromString, frameworkHandle, "NSPointFromString")
	registerFunc(&_nSProtocolFromString, frameworkHandle, "NSProtocolFromString")
	registerFunc(&_nSRangeFromString, frameworkHandle, "NSRangeFromString")
	registerFunc(&_nSReallocateCollectable, frameworkHandle, "NSReallocateCollectable")
	registerFunc(&_nSRecordAllocationEvent, frameworkHandle, "NSRecordAllocationEvent")
	registerFunc(&_nSRectFromString, frameworkHandle, "NSRectFromString")
	registerFunc(&_nSRecycleZone, frameworkHandle, "NSRecycleZone")
	registerFunc(&_nSResetHashTable, frameworkHandle, "NSResetHashTable")
	registerFunc(&_nSResetMapTable, frameworkHandle, "NSResetMapTable")
	registerFunc(&_nSReturnAddress, frameworkHandle, "NSReturnAddress")
	registerFunc(&_nSRoundDownToMultipleOfPageSize, frameworkHandle, "NSRoundDownToMultipleOfPageSize")
	registerFunc(&_nSRoundUpToMultipleOfPageSize, frameworkHandle, "NSRoundUpToMultipleOfPageSize")
	registerFunc(&_nSSearchPathForDirectoriesInDomains, frameworkHandle, "NSSearchPathForDirectoriesInDomains")
	registerFunc(&_nSSelectorFromString, frameworkHandle, "NSSelectorFromString")
	registerFunc(&_nSSetUncaughtExceptionHandler, frameworkHandle, "NSSetUncaughtExceptionHandler")
	registerFunc(&_nSSetZoneName, frameworkHandle, "NSSetZoneName")
	registerFunc(&_nSShouldRetainWithZone, frameworkHandle, "NSShouldRetainWithZone")
	registerFunc(&_nSSizeFromString, frameworkHandle, "NSSizeFromString")
	registerFunc(&_nSStringFromClass, frameworkHandle, "NSStringFromClass")
	registerFunc(&_nSStringFromHashTable, frameworkHandle, "NSStringFromHashTable")
	registerFunc(&_nSStringFromMapTable, frameworkHandle, "NSStringFromMapTable")
	registerFunc(&_nSStringFromPoint, frameworkHandle, "NSStringFromPoint")
	registerFunc(&_nSStringFromProtocol, frameworkHandle, "NSStringFromProtocol")
	registerFunc(&_nSStringFromRange, frameworkHandle, "NSStringFromRange")
	registerFunc(&_nSStringFromRect, frameworkHandle, "NSStringFromRect")
	registerFunc(&_nSStringFromSelector, frameworkHandle, "NSStringFromSelector")
	registerFunc(&_nSStringFromSize, frameworkHandle, "NSStringFromSize")
	registerFunc(&_nSTemporaryDirectory, frameworkHandle, "NSTemporaryDirectory")
	registerFunc(&_nSUserName, frameworkHandle, "NSUserName")
	registerFunc(&_nSZoneCalloc, frameworkHandle, "NSZoneCalloc")
	registerFunc(&_nSZoneFree, frameworkHandle, "NSZoneFree")
	registerFunc(&_nSZoneFromPointer, frameworkHandle, "NSZoneFromPointer")
	registerFunc(&_nSZoneMalloc, frameworkHandle, "NSZoneMalloc")
	registerFunc(&_nSZoneName, frameworkHandle, "NSZoneName")
	registerFunc(&_nSZoneRealloc, frameworkHandle, "NSZoneRealloc")
	registerFunc(&_nXReadNSObjectFromCoder, frameworkHandle, "NXReadNSObjectFromCoder")
}
