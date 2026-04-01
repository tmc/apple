// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("Foundation: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Foundation: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("Foundation: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("Foundation: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _nSAllHashTableObjects func(table *NSHashTable) NSArray
var _nSAllHashTableObjectsErr error

func tryNSAllHashTableObjects(table *NSHashTable) (NSArray, error) {
	if _nSAllHashTableObjects == nil {
		return NSArray{}, symbolCallError("NSAllHashTableObjects", "10.0", _nSAllHashTableObjectsErr)
	}
	return _nSAllHashTableObjects(table), nil
}

// NSAllHashTableObjects returns all of the elements in the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllHashTableObjects(_:)
func NSAllHashTableObjects(table *NSHashTable) NSArray {
	result, callErr := tryNSAllHashTableObjects(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAllMapTableKeys func(table *NSMapTable) NSArray
var _nSAllMapTableKeysErr error

func tryNSAllMapTableKeys(table *NSMapTable) (NSArray, error) {
	if _nSAllMapTableKeys == nil {
		return NSArray{}, symbolCallError("NSAllMapTableKeys", "10.0", _nSAllMapTableKeysErr)
	}
	return _nSAllMapTableKeys(table), nil
}

// NSAllMapTableKeys returns all of the keys in the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllMapTableKeys(_:)
func NSAllMapTableKeys(table *NSMapTable) NSArray {
	result, callErr := tryNSAllMapTableKeys(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAllMapTableValues func(table *NSMapTable) NSArray
var _nSAllMapTableValuesErr error

func tryNSAllMapTableValues(table *NSMapTable) (NSArray, error) {
	if _nSAllMapTableValues == nil {
		return NSArray{}, symbolCallError("NSAllMapTableValues", "10.0", _nSAllMapTableValuesErr)
	}
	return _nSAllMapTableValues(table), nil
}

// NSAllMapTableValues returns all of the values in the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllMapTableValues(_:)
func NSAllMapTableValues(table *NSMapTable) NSArray {
	result, callErr := tryNSAllMapTableValues(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAllocateCollectable func(size uint, options uint) unsafe.Pointer
var _nSAllocateCollectableErr error

func tryNSAllocateCollectable(size uint, options uint) (unsafe.Pointer, error) {
	if _nSAllocateCollectable == nil {
		return nil, symbolCallError("NSAllocateCollectable", "10.0", _nSAllocateCollectableErr)
	}
	return _nSAllocateCollectable(size, options), nil
}

// NSAllocateCollectable allocates collectable memory.
//
// Deprecated: Garbage collection is deprecated in OS X v10.8; instead,you should use AutomaticReference Counting—see [Transitioning to ARC Release Notes](<https://developer.apple.com/library/archive/releasenotes/ObjectiveC/RN-TransitioningToARC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011226>).
//
// See: https://developer.apple.com/documentation/Foundation/NSAllocateCollectable
func NSAllocateCollectable(size uint, options uint) unsafe.Pointer {
	result, callErr := tryNSAllocateCollectable(size, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAllocateMemoryPages func(bytes uint) unsafe.Pointer
var _nSAllocateMemoryPagesErr error

func tryNSAllocateMemoryPages(bytes uint) (unsafe.Pointer, error) {
	if _nSAllocateMemoryPages == nil {
		return nil, symbolCallError("NSAllocateMemoryPages", "10.0", _nSAllocateMemoryPagesErr)
	}
	return _nSAllocateMemoryPages(bytes), nil
}

// NSAllocateMemoryPages allocates a new block of memory.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllocateMemoryPages(_:)
func NSAllocateMemoryPages(bytes uint) unsafe.Pointer {
	result, callErr := tryNSAllocateMemoryPages(bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAllocateObject func(aClass objc.Class, extraBytes uint, zone *NSZone) objectivec.Object
var _nSAllocateObjectErr error

func tryNSAllocateObject(aClass objc.Class, extraBytes uint, zone *NSZone) (objectivec.Object, error) {
	if _nSAllocateObject == nil {
		return objectivec.Object{}, symbolCallError("NSAllocateObject", "10.0", _nSAllocateObjectErr)
	}
	return _nSAllocateObject(aClass, extraBytes, zone), nil
}

// NSAllocateObject creates and returns a new instance of a given class.
//
// See: https://developer.apple.com/documentation/Foundation/NSAllocateObject
func NSAllocateObject(aClass objc.Class, extraBytes uint, zone *NSZone) objectivec.Object {
	result, callErr := tryNSAllocateObject(aClass, extraBytes, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSClassFromString func(aClassName NSString) objc.Class
var _nSClassFromStringErr error

func tryNSClassFromString(aClassName NSString) (objc.Class, error) {
	if _nSClassFromString == nil {
		return 0, symbolCallError("NSClassFromString", "10.0", _nSClassFromStringErr)
	}
	return _nSClassFromString(aClassName), nil
}

// NSClassFromString obtains a class by name.
//
// See: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName NSString) objc.Class {
	result, callErr := tryNSClassFromString(aClassName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCompareHashTables func(table1 *NSHashTable, table2 *NSHashTable) bool
var _nSCompareHashTablesErr error

func tryNSCompareHashTables(table1 *NSHashTable, table2 *NSHashTable) (bool, error) {
	if _nSCompareHashTables == nil {
		return false, symbolCallError("NSCompareHashTables", "10.0", _nSCompareHashTablesErr)
	}
	return _nSCompareHashTables(table1, table2), nil
}

// NSCompareHashTables returns a Boolean value that indicates whether the elements of two hash tables are equal.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompareHashTables(_:_:)
func NSCompareHashTables(table1 *NSHashTable, table2 *NSHashTable) bool {
	result, callErr := tryNSCompareHashTables(table1, table2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCompareMapTables func(table1 *NSMapTable, table2 *NSMapTable) bool
var _nSCompareMapTablesErr error

func tryNSCompareMapTables(table1 *NSMapTable, table2 *NSMapTable) (bool, error) {
	if _nSCompareMapTables == nil {
		return false, symbolCallError("NSCompareMapTables", "10.0", _nSCompareMapTablesErr)
	}
	return _nSCompareMapTables(table1, table2), nil
}

// NSCompareMapTables compares the elements of two map tables for equality.
//
// See: https://developer.apple.com/documentation/Foundation/NSCompareMapTables(_:_:)
func NSCompareMapTables(table1 *NSMapTable, table2 *NSMapTable) bool {
	result, callErr := tryNSCompareMapTables(table1, table2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

// NSContainsRect returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// See: https://developer.apple.com/documentation/Foundation/NSContainsRect(_:_:)
func NSContainsRect(aRect corefoundation.CGRect, bRect corefoundation.CGRect) bool {
	return nsContainsRect(aRect, bRect)
}

var _nSCopyHashTableWithZone func(table *NSHashTable, zone *NSZone) *NSHashTable
var _nSCopyHashTableWithZoneErr error

func tryNSCopyHashTableWithZone(table *NSHashTable, zone *NSZone) (*NSHashTable, error) {
	if _nSCopyHashTableWithZone == nil {
		return nil, symbolCallError("NSCopyHashTableWithZone", "10.0", _nSCopyHashTableWithZoneErr)
	}
	return _nSCopyHashTableWithZone(table, zone), nil
}

// NSCopyHashTableWithZone performs a shallow copy of the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyHashTableWithZone(_:_:)
func NSCopyHashTableWithZone(table *NSHashTable, zone *NSZone) *NSHashTable {
	result, callErr := tryNSCopyHashTableWithZone(table, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCopyMapTableWithZone func(table *NSMapTable, zone *NSZone) *NSMapTable
var _nSCopyMapTableWithZoneErr error

func tryNSCopyMapTableWithZone(table *NSMapTable, zone *NSZone) (*NSMapTable, error) {
	if _nSCopyMapTableWithZone == nil {
		return nil, symbolCallError("NSCopyMapTableWithZone", "10.0", _nSCopyMapTableWithZoneErr)
	}
	return _nSCopyMapTableWithZone(table, zone), nil
}

// NSCopyMapTableWithZone performs a shallow copy of the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyMapTableWithZone(_:_:)
func NSCopyMapTableWithZone(table *NSMapTable, zone *NSZone) *NSMapTable {
	result, callErr := tryNSCopyMapTableWithZone(table, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCopyMemoryPages func(source unsafe.Pointer, dest unsafe.Pointer, bytes uint)
var _nSCopyMemoryPagesErr error

func tryNSCopyMemoryPages(source unsafe.Pointer, dest unsafe.Pointer, bytes uint) error {
	if _nSCopyMemoryPages == nil {
		return symbolCallError("NSCopyMemoryPages", "10.0", _nSCopyMemoryPagesErr)
	}
	_nSCopyMemoryPages(source, dest, bytes)
	return nil
}

// NSCopyMemoryPages copies a block of memory.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyMemoryPages(_:_:_:)
func NSCopyMemoryPages(source unsafe.Pointer, dest unsafe.Pointer, bytes uint) {
	if callErr := tryNSCopyMemoryPages(source, dest, bytes); callErr != nil {
		panic(callErr)
	}
}

var _nSCopyObject func(object objectivec.Object, extraBytes uint, zone *NSZone) objectivec.Object
var _nSCopyObjectErr error

func tryNSCopyObject(object objectivec.Object, extraBytes uint, zone *NSZone) (objectivec.Object, error) {
	if _nSCopyObject == nil {
		return objectivec.Object{}, symbolCallError("NSCopyObject", "10.0", _nSCopyObjectErr)
	}
	return _nSCopyObject(object, extraBytes, zone), nil
}

// NSCopyObject creates an exact copy of an object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/Foundation/NSCopyObject
func NSCopyObject(object objectivec.Object, extraBytes uint, zone *NSZone) objectivec.Object {
	result, callErr := tryNSCopyObject(object, extraBytes, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCountFrames func() uint
var _nSCountFramesErr error

func tryNSCountFrames() (uint, error) {
	if _nSCountFrames == nil {
		return 0, symbolCallError("NSCountFrames", "10.0", _nSCountFramesErr)
	}
	return _nSCountFrames(), nil
}

// NSCountFrames returns the number of call frames on the stack.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() uint {
	result, callErr := tryNSCountFrames()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCountHashTable func(table *NSHashTable) uint
var _nSCountHashTableErr error

func tryNSCountHashTable(table *NSHashTable) (uint, error) {
	if _nSCountHashTable == nil {
		return 0, symbolCallError("NSCountHashTable", "10.0", _nSCountHashTableErr)
	}
	return _nSCountHashTable(table), nil
}

// NSCountHashTable returns the number of elements in a hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountHashTable(_:)
func NSCountHashTable(table *NSHashTable) uint {
	result, callErr := tryNSCountHashTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCountMapTable func(table *NSMapTable) uint
var _nSCountMapTableErr error

func tryNSCountMapTable(table *NSMapTable) (uint, error) {
	if _nSCountMapTable == nil {
		return 0, symbolCallError("NSCountMapTable", "10.0", _nSCountMapTableErr)
	}
	return _nSCountMapTable(table), nil
}

// NSCountMapTable returns the number of elements in a map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCountMapTable(_:)
func NSCountMapTable(table *NSMapTable) uint {
	result, callErr := tryNSCountMapTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCreateHashTable func(callBacks NSHashTableCallBacks, capacity uint) *NSHashTable
var _nSCreateHashTableErr error

func tryNSCreateHashTable(callBacks NSHashTableCallBacks, capacity uint) (*NSHashTable, error) {
	if _nSCreateHashTable == nil {
		return nil, symbolCallError("NSCreateHashTable", "10.0", _nSCreateHashTableErr)
	}
	return _nSCreateHashTable(callBacks, capacity), nil
}

// NSCreateHashTable creates and returns a new hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateHashTable(_:_:)
func NSCreateHashTable(callBacks NSHashTableCallBacks, capacity uint) *NSHashTable {
	result, callErr := tryNSCreateHashTable(callBacks, capacity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCreateHashTableWithZone func(callBacks NSHashTableCallBacks, capacity uint, zone *NSZone) *NSHashTable
var _nSCreateHashTableWithZoneErr error

func tryNSCreateHashTableWithZone(callBacks NSHashTableCallBacks, capacity uint, zone *NSZone) (*NSHashTable, error) {
	if _nSCreateHashTableWithZone == nil {
		return nil, symbolCallError("NSCreateHashTableWithZone", "10.0", _nSCreateHashTableWithZoneErr)
	}
	return _nSCreateHashTableWithZone(callBacks, capacity, zone), nil
}

// NSCreateHashTableWithZone creates a new hash table in a given zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateHashTableWithZone(_:_:_:)
func NSCreateHashTableWithZone(callBacks NSHashTableCallBacks, capacity uint, zone *NSZone) *NSHashTable {
	result, callErr := tryNSCreateHashTableWithZone(callBacks, capacity, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCreateMapTable func(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint) *NSMapTable
var _nSCreateMapTableErr error

func tryNSCreateMapTable(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint) (*NSMapTable, error) {
	if _nSCreateMapTable == nil {
		return nil, symbolCallError("NSCreateMapTable", "10.0", _nSCreateMapTableErr)
	}
	return _nSCreateMapTable(keyCallBacks, valueCallBacks, capacity), nil
}

// NSCreateMapTable creates a new map table in the default zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateMapTable(_:_:_:)
func NSCreateMapTable(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint) *NSMapTable {
	result, callErr := tryNSCreateMapTable(keyCallBacks, valueCallBacks, capacity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCreateMapTableWithZone func(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint, zone *NSZone) *NSMapTable
var _nSCreateMapTableWithZoneErr error

func tryNSCreateMapTableWithZone(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint, zone *NSZone) (*NSMapTable, error) {
	if _nSCreateMapTableWithZone == nil {
		return nil, symbolCallError("NSCreateMapTableWithZone", "10.0", _nSCreateMapTableWithZoneErr)
	}
	return _nSCreateMapTableWithZone(keyCallBacks, valueCallBacks, capacity, zone), nil
}

// NSCreateMapTableWithZone creates a new map table in the specified zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateMapTableWithZone(_:_:_:_:)
func NSCreateMapTableWithZone(keyCallBacks NSMapTableKeyCallBacks, valueCallBacks NSMapTableValueCallBacks, capacity uint, zone *NSZone) *NSMapTable {
	result, callErr := tryNSCreateMapTableWithZone(keyCallBacks, valueCallBacks, capacity, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCreateZone func(startSize uint, granularity uint, canFree bool) *NSZone
var _nSCreateZoneErr error

func tryNSCreateZone(startSize uint, granularity uint, canFree bool) (*NSZone, error) {
	if _nSCreateZone == nil {
		return nil, symbolCallError("NSCreateZone", "10.0", _nSCreateZoneErr)
	}
	return _nSCreateZone(startSize, granularity, canFree), nil
}

// NSCreateZone creates a new zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSCreateZone
func NSCreateZone(startSize uint, granularity uint, canFree bool) *NSZone {
	result, callErr := tryNSCreateZone(startSize, granularity, canFree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDeallocateMemoryPages func(ptr unsafe.Pointer, bytes uint)
var _nSDeallocateMemoryPagesErr error

func tryNSDeallocateMemoryPages(ptr unsafe.Pointer, bytes uint) error {
	if _nSDeallocateMemoryPages == nil {
		return symbolCallError("NSDeallocateMemoryPages", "10.0", _nSDeallocateMemoryPagesErr)
	}
	_nSDeallocateMemoryPages(ptr, bytes)
	return nil
}

// NSDeallocateMemoryPages deallocates the specified block of memory.
//
// See: https://developer.apple.com/documentation/Foundation/NSDeallocateMemoryPages(_:_:)
func NSDeallocateMemoryPages(ptr unsafe.Pointer, bytes uint) {
	if callErr := tryNSDeallocateMemoryPages(ptr, bytes); callErr != nil {
		panic(callErr)
	}
}

var _nSDeallocateObject func(object objectivec.Object)
var _nSDeallocateObjectErr error

func tryNSDeallocateObject(object objectivec.Object) error {
	if _nSDeallocateObject == nil {
		return symbolCallError("NSDeallocateObject", "10.0", _nSDeallocateObjectErr)
	}
	_nSDeallocateObject(object)
	return nil
}

// NSDeallocateObject destroys an existing object.
//
// See: https://developer.apple.com/documentation/Foundation/NSDeallocateObject
func NSDeallocateObject(object objectivec.Object) {
	if callErr := tryNSDeallocateObject(object); callErr != nil {
		panic(callErr)
	}
}

var _nSDecimalAdd func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalAddErr error

func tryNSDecimalAdd(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalAdd == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalAdd", "10.0", _nSDecimalAddErr)
	}
	return _nSDecimalAdd(result, leftOperand, rightOperand, roundingMode), nil
}

// NSDecimalAdd adds two decimal values.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalAdd(_:_:_:_:)
func NSDecimalAdd(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	result0, callErr := tryNSDecimalAdd(result, leftOperand, rightOperand, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nSDecimalCompact func(number *NSDecimal)
var _nSDecimalCompactErr error

func tryNSDecimalCompact(number *NSDecimal) error {
	if _nSDecimalCompact == nil {
		return symbolCallError("NSDecimalCompact", "10.0", _nSDecimalCompactErr)
	}
	_nSDecimalCompact(number)
	return nil
}

// NSDecimalCompact compacts the decimal structure for efficiency.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number *NSDecimal) {
	if callErr := tryNSDecimalCompact(number); callErr != nil {
		panic(callErr)
	}
}

var _nSDecimalCompare func(leftOperand *NSDecimal, rightOperand *NSDecimal) NSComparisonResult
var _nSDecimalCompareErr error

func tryNSDecimalCompare(leftOperand *NSDecimal, rightOperand *NSDecimal) (NSComparisonResult, error) {
	if _nSDecimalCompare == nil {
		return *new(NSComparisonResult), symbolCallError("NSDecimalCompare", "10.0", _nSDecimalCompareErr)
	}
	return _nSDecimalCompare(leftOperand, rightOperand), nil
}

// NSDecimalCompare compares two decimal values.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalCompare(_:_:)
func NSDecimalCompare(leftOperand *NSDecimal, rightOperand *NSDecimal) NSComparisonResult {
	result, callErr := tryNSDecimalCompare(leftOperand, rightOperand)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDecimalCopy func(destination *NSDecimal, source *NSDecimal)
var _nSDecimalCopyErr error

func tryNSDecimalCopy(destination *NSDecimal, source *NSDecimal) error {
	if _nSDecimalCopy == nil {
		return symbolCallError("NSDecimalCopy", "10.0", _nSDecimalCopyErr)
	}
	_nSDecimalCopy(destination, source)
	return nil
}

// NSDecimalCopy copies the value of a decimal number.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalCopy(_:_:)
func NSDecimalCopy(destination *NSDecimal, source *NSDecimal) {
	if callErr := tryNSDecimalCopy(destination, source); callErr != nil {
		panic(callErr)
	}
}

var _nSDecimalDivide func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalDivideErr error

func tryNSDecimalDivide(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalDivide == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalDivide", "10.0", _nSDecimalDivideErr)
	}
	return _nSDecimalDivide(result, leftOperand, rightOperand, roundingMode), nil
}

// NSDecimalDivide divides one decimal value by another.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalDivide(_:_:_:_:)
func NSDecimalDivide(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	result0, callErr := tryNSDecimalDivide(result, leftOperand, rightOperand, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nSDecimalMultiply func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalMultiplyErr error

func tryNSDecimalMultiply(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalMultiply == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalMultiply", "10.0", _nSDecimalMultiplyErr)
	}
	return _nSDecimalMultiply(result, leftOperand, rightOperand, roundingMode), nil
}

// NSDecimalMultiply multiplies two decimal numbers together.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalMultiply(_:_:_:_:)
func NSDecimalMultiply(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	result0, callErr := tryNSDecimalMultiply(result, leftOperand, rightOperand, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nSDecimalMultiplyByPowerOf10 func(result *NSDecimal, number *NSDecimal, power int16, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalMultiplyByPowerOf10Err error

func tryNSDecimalMultiplyByPowerOf10(result *NSDecimal, number *NSDecimal, power int16, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalMultiplyByPowerOf10 == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalMultiplyByPowerOf10", "10.0", _nSDecimalMultiplyByPowerOf10Err)
	}
	return _nSDecimalMultiplyByPowerOf10(result, number, power, roundingMode), nil
}

// NSDecimalMultiplyByPowerOf10 multiplies a decimal by the specified power of 10.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalMultiplyByPowerOf10(_:_:_:_:)
func NSDecimalMultiplyByPowerOf10(result *NSDecimal, number *NSDecimal, power int16, roundingMode NSRoundingMode) NSCalculationError {
	result0, callErr := tryNSDecimalMultiplyByPowerOf10(result, number, power, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nSDecimalNormalize func(number1 *NSDecimal, number2 *NSDecimal, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalNormalizeErr error

func tryNSDecimalNormalize(number1 *NSDecimal, number2 *NSDecimal, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalNormalize == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalNormalize", "10.0", _nSDecimalNormalizeErr)
	}
	return _nSDecimalNormalize(number1, number2, roundingMode), nil
}

// NSDecimalNormalize normalizes the internal format of two decimal numbers to simplify later operations.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalNormalize(_:_:_:)
func NSDecimalNormalize(number1 *NSDecimal, number2 *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	result, callErr := tryNSDecimalNormalize(number1, number2, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDecimalPower func(result *NSDecimal, number *NSDecimal, power uint, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalPowerErr error

func tryNSDecimalPower(result *NSDecimal, number *NSDecimal, power uint, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalPower == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalPower", "10.0", _nSDecimalPowerErr)
	}
	return _nSDecimalPower(result, number, power, roundingMode), nil
}

// NSDecimalPower raises the decimal value to the specified power.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result *NSDecimal, number *NSDecimal, power uint, roundingMode NSRoundingMode) NSCalculationError {
	result0, callErr := tryNSDecimalPower(result, number, power, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nSDecimalRound func(result *NSDecimal, number *NSDecimal, scale int, roundingMode NSRoundingMode)
var _nSDecimalRoundErr error

func tryNSDecimalRound(result *NSDecimal, number *NSDecimal, scale int, roundingMode NSRoundingMode) error {
	if _nSDecimalRound == nil {
		return symbolCallError("NSDecimalRound", "10.0", _nSDecimalRoundErr)
	}
	_nSDecimalRound(result, number, scale, roundingMode)
	return nil
}

// NSDecimalRound rounds off the decimal value.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalRound(_:_:_:_:)
func NSDecimalRound(result *NSDecimal, number *NSDecimal, scale int, roundingMode NSRoundingMode) {
	if callErr := tryNSDecimalRound(result, number, scale, roundingMode); callErr != nil {
		panic(callErr)
	}
}

var _nSDecimalString func(dcm *NSDecimal, locale objectivec.Object) NSString
var _nSDecimalStringErr error

func tryNSDecimalString(dcm *NSDecimal, locale objectivec.Object) (NSString, error) {
	if _nSDecimalString == nil {
		return NSString{}, symbolCallError("NSDecimalString", "10.0", _nSDecimalStringErr)
	}
	return _nSDecimalString(dcm, locale), nil
}

// NSDecimalString returns a string representation of the decimal value appropriate for the specified locale.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalString(_:_:)
func NSDecimalString(dcm *NSDecimal, locale objectivec.Object) NSString {
	result, callErr := tryNSDecimalString(dcm, locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDecimalSubtract func(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError
var _nSDecimalSubtractErr error

func tryNSDecimalSubtract(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) (NSCalculationError, error) {
	if _nSDecimalSubtract == nil {
		return *new(NSCalculationError), symbolCallError("NSDecimalSubtract", "10.0", _nSDecimalSubtractErr)
	}
	return _nSDecimalSubtract(result, leftOperand, rightOperand, roundingMode), nil
}

// NSDecimalSubtract subtracts one decimal value from another.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecimalSubtract(_:_:_:_:)
func NSDecimalSubtract(result *NSDecimal, leftOperand *NSDecimal, rightOperand *NSDecimal, roundingMode NSRoundingMode) NSCalculationError {
	result0, callErr := tryNSDecimalSubtract(result, leftOperand, rightOperand, roundingMode)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _nSDecrementExtraRefCountWasZero func(object objectivec.Object) bool
var _nSDecrementExtraRefCountWasZeroErr error

func tryNSDecrementExtraRefCountWasZero(object objectivec.Object) (bool, error) {
	if _nSDecrementExtraRefCountWasZero == nil {
		return false, symbolCallError("NSDecrementExtraRefCountWasZero", "10.0", _nSDecrementExtraRefCountWasZeroErr)
	}
	return _nSDecrementExtraRefCountWasZero(object), nil
}

// NSDecrementExtraRefCountWasZero decrements the specified object’s reference count.
//
// See: https://developer.apple.com/documentation/Foundation/NSDecrementExtraRefCountWasZero
func NSDecrementExtraRefCountWasZero(object objectivec.Object) bool {
	result, callErr := tryNSDecrementExtraRefCountWasZero(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDefaultMallocZone func() *NSZone
var _nSDefaultMallocZoneErr error

func tryNSDefaultMallocZone() (*NSZone, error) {
	if _nSDefaultMallocZone == nil {
		return nil, symbolCallError("NSDefaultMallocZone", "10.0", _nSDefaultMallocZoneErr)
	}
	return _nSDefaultMallocZone(), nil
}

// NSDefaultMallocZone returns the default zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSDefaultMallocZone
func NSDefaultMallocZone() *NSZone {
	result, callErr := tryNSDefaultMallocZone()
	if callErr != nil {
		panic(callErr)
	}
	return result
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
var _nSEndHashTableEnumerationErr error

func tryNSEndHashTableEnumeration(enumerator *NSHashEnumerator) error {
	if _nSEndHashTableEnumeration == nil {
		return symbolCallError("NSEndHashTableEnumeration", "10.0", _nSEndHashTableEnumerationErr)
	}
	_nSEndHashTableEnumeration(enumerator)
	return nil
}

// NSEndHashTableEnumeration used when finished with an enumerator.
//
// See: https://developer.apple.com/documentation/Foundation/NSEndHashTableEnumeration(_:)
func NSEndHashTableEnumeration(enumerator *NSHashEnumerator) {
	if callErr := tryNSEndHashTableEnumeration(enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nSEndMapTableEnumeration func(enumerator *NSMapEnumerator)
var _nSEndMapTableEnumerationErr error

func tryNSEndMapTableEnumeration(enumerator *NSMapEnumerator) error {
	if _nSEndMapTableEnumeration == nil {
		return symbolCallError("NSEndMapTableEnumeration", "10.0", _nSEndMapTableEnumerationErr)
	}
	_nSEndMapTableEnumeration(enumerator)
	return nil
}

// NSEndMapTableEnumeration used when finished with an enumerator.
//
// See: https://developer.apple.com/documentation/Foundation/NSEndMapTableEnumeration(_:)
func NSEndMapTableEnumeration(enumerator *NSMapEnumerator) {
	if callErr := tryNSEndMapTableEnumeration(enumerator); callErr != nil {
		panic(callErr)
	}
}

var _nSEnumerateHashTable func(table *NSHashTable) NSHashEnumerator
var _nSEnumerateHashTableErr error

func tryNSEnumerateHashTable(table *NSHashTable) (NSHashEnumerator, error) {
	if _nSEnumerateHashTable == nil {
		return NSHashEnumerator{}, symbolCallError("NSEnumerateHashTable", "10.0", _nSEnumerateHashTableErr)
	}
	return _nSEnumerateHashTable(table), nil
}

// NSEnumerateHashTable creates an enumerator for the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerateHashTable(_:)
func NSEnumerateHashTable(table *NSHashTable) NSHashEnumerator {
	result, callErr := tryNSEnumerateHashTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSEnumerateMapTable func(table *NSMapTable) NSMapEnumerator
var _nSEnumerateMapTableErr error

func tryNSEnumerateMapTable(table *NSMapTable) (NSMapEnumerator, error) {
	if _nSEnumerateMapTable == nil {
		return NSMapEnumerator{}, symbolCallError("NSEnumerateMapTable", "10.0", _nSEnumerateMapTableErr)
	}
	return _nSEnumerateMapTable(table), nil
}

// NSEnumerateMapTable creates an enumerator for the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerateMapTable(_:)
func NSEnumerateMapTable(table *NSMapTable) NSMapEnumerator {
	result, callErr := tryNSEnumerateMapTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
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
var _nSExtraRefCountErr error

func tryNSExtraRefCount(object objectivec.Object) (uint, error) {
	if _nSExtraRefCount == nil {
		return 0, symbolCallError("NSExtraRefCount", "10.0", _nSExtraRefCountErr)
	}
	return _nSExtraRefCount(object), nil
}

// NSExtraRefCount returns the specified object’s reference count.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtraRefCount
func NSExtraRefCount(object objectivec.Object) uint {
	result, callErr := tryNSExtraRefCount(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSFileTypeForHFSTypeCode func(hfsFileTypeCode uint32) NSString
var _nSFileTypeForHFSTypeCodeErr error

func tryNSFileTypeForHFSTypeCode(hfsFileTypeCode uint32) (NSString, error) {
	if _nSFileTypeForHFSTypeCode == nil {
		return NSString{}, symbolCallError("NSFileTypeForHFSTypeCode", "10.0", _nSFileTypeForHFSTypeCodeErr)
	}
	return _nSFileTypeForHFSTypeCode(hfsFileTypeCode), nil
}

// NSFileTypeForHFSTypeCode returns a string encoding a file type code.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileTypeForHFSTypeCode(_:)
func NSFileTypeForHFSTypeCode(hfsFileTypeCode uint32) NSString {
	result, callErr := tryNSFileTypeForHFSTypeCode(hfsFileTypeCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSFrameAddress func(frame uint) unsafe.Pointer
var _nSFrameAddressErr error

func tryNSFrameAddress(frame uint) (unsafe.Pointer, error) {
	if _nSFrameAddress == nil {
		return nil, symbolCallError("NSFrameAddress", "10.0", _nSFrameAddressErr)
	}
	return _nSFrameAddress(frame), nil
}

// NSFrameAddress returns the value of the frame pointer of the specified frame.
//
// See: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame uint) unsafe.Pointer {
	result, callErr := tryNSFrameAddress(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSFreeHashTable func(table *NSHashTable)
var _nSFreeHashTableErr error

func tryNSFreeHashTable(table *NSHashTable) error {
	if _nSFreeHashTable == nil {
		return symbolCallError("NSFreeHashTable", "10.0", _nSFreeHashTableErr)
	}
	_nSFreeHashTable(table)
	return nil
}

// NSFreeHashTable deletes the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSFreeHashTable(_:)
func NSFreeHashTable(table *NSHashTable) {
	if callErr := tryNSFreeHashTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nSFreeMapTable func(table *NSMapTable)
var _nSFreeMapTableErr error

func tryNSFreeMapTable(table *NSMapTable) error {
	if _nSFreeMapTable == nil {
		return symbolCallError("NSFreeMapTable", "10.0", _nSFreeMapTableErr)
	}
	_nSFreeMapTable(table)
	return nil
}

// NSFreeMapTable deletes the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSFreeMapTable(_:)
func NSFreeMapTable(table *NSMapTable) {
	if callErr := tryNSFreeMapTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nSFullUserName func() NSString
var _nSFullUserNameErr error

func tryNSFullUserName() (NSString, error) {
	if _nSFullUserName == nil {
		return NSString{}, symbolCallError("NSFullUserName", "10.0", _nSFullUserNameErr)
	}
	return _nSFullUserName(), nil
}

// NSFullUserName returns a string containing the full name of the current user.
//
// See: https://developer.apple.com/documentation/Foundation/NSFullUserName()
func NSFullUserName() NSString {
	result, callErr := tryNSFullUserName()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSGetSizeAndAlignment func(typePtr string, sizep *uint, alignp *uint) *byte
var _nSGetSizeAndAlignmentErr error

func tryNSGetSizeAndAlignment(typePtr string, sizep *uint, alignp *uint) (*byte, error) {
	if _nSGetSizeAndAlignment == nil {
		return nil, symbolCallError("NSGetSizeAndAlignment", "10.0", _nSGetSizeAndAlignmentErr)
	}
	return _nSGetSizeAndAlignment(typePtr, sizep, alignp), nil
}

// NSGetSizeAndAlignment obtains the actual size and the aligned size of an encoded type.
//
// See: https://developer.apple.com/documentation/Foundation/NSGetSizeAndAlignment(_:_:_:)
func NSGetSizeAndAlignment(typePtr string, sizep *uint, alignp *uint) *byte {
	result, callErr := tryNSGetSizeAndAlignment(typePtr, sizep, alignp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSGetUncaughtExceptionHandler func() func(unsafe.Pointer)
var _nSGetUncaughtExceptionHandlerErr error

func tryNSGetUncaughtExceptionHandler() (func(unsafe.Pointer), error) {
	if _nSGetUncaughtExceptionHandler == nil {
		return nil, symbolCallError("NSGetUncaughtExceptionHandler", "10.0", _nSGetUncaughtExceptionHandlerErr)
	}
	return _nSGetUncaughtExceptionHandler(), nil
}

// NSGetUncaughtExceptionHandler returns the top-level error handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSGetUncaughtExceptionHandler()
func NSGetUncaughtExceptionHandler() func(unsafe.Pointer) {
	result, callErr := tryNSGetUncaughtExceptionHandler()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHFSTypeCodeFromFileType func(fileTypeString NSString) uint32
var _nSHFSTypeCodeFromFileTypeErr error

func tryNSHFSTypeCodeFromFileType(fileTypeString NSString) (uint32, error) {
	if _nSHFSTypeCodeFromFileType == nil {
		return 0, symbolCallError("NSHFSTypeCodeFromFileType", "10.0", _nSHFSTypeCodeFromFileTypeErr)
	}
	return _nSHFSTypeCodeFromFileType(fileTypeString), nil
}

// NSHFSTypeCodeFromFileType returns a file type code.
//
// See: https://developer.apple.com/documentation/Foundation/NSHFSTypeCodeFromFileType(_:)
func NSHFSTypeCodeFromFileType(fileTypeString NSString) uint32 {
	result, callErr := tryNSHFSTypeCodeFromFileType(fileTypeString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHFSTypeOfFile func(fullFilePath NSString) NSString
var _nSHFSTypeOfFileErr error

func tryNSHFSTypeOfFile(fullFilePath NSString) (NSString, error) {
	if _nSHFSTypeOfFile == nil {
		return NSString{}, symbolCallError("NSHFSTypeOfFile", "10.0", _nSHFSTypeOfFileErr)
	}
	return _nSHFSTypeOfFile(fullFilePath), nil
}

// NSHFSTypeOfFile returns a string encoding a file type.
//
// See: https://developer.apple.com/documentation/Foundation/NSHFSTypeOfFile(_:)
func NSHFSTypeOfFile(fullFilePath NSString) NSString {
	result, callErr := tryNSHFSTypeOfFile(fullFilePath)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHashGet func(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer
var _nSHashGetErr error

func tryNSHashGet(table *NSHashTable, pointer unsafe.Pointer) (unsafe.Pointer, error) {
	if _nSHashGet == nil {
		return nil, symbolCallError("NSHashGet", "10.0", _nSHashGetErr)
	}
	return _nSHashGet(table, pointer), nil
}

// NSHashGet returns an element of the hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashGet(_:_:)
func NSHashGet(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNSHashGet(table, pointer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHashInsert func(table *NSHashTable, pointer unsafe.Pointer)
var _nSHashInsertErr error

func tryNSHashInsert(table *NSHashTable, pointer unsafe.Pointer) error {
	if _nSHashInsert == nil {
		return symbolCallError("NSHashInsert", "10.0", _nSHashInsertErr)
	}
	_nSHashInsert(table, pointer)
	return nil
}

// NSHashInsert adds an element to the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashInsert(_:_:)
func NSHashInsert(table *NSHashTable, pointer unsafe.Pointer) {
	if callErr := tryNSHashInsert(table, pointer); callErr != nil {
		panic(callErr)
	}
}

var _nSHashInsertIfAbsent func(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer
var _nSHashInsertIfAbsentErr error

func tryNSHashInsertIfAbsent(table *NSHashTable, pointer unsafe.Pointer) (unsafe.Pointer, error) {
	if _nSHashInsertIfAbsent == nil {
		return nil, symbolCallError("NSHashInsertIfAbsent", "10.0", _nSHashInsertIfAbsentErr)
	}
	return _nSHashInsertIfAbsent(table, pointer), nil
}

// NSHashInsertIfAbsent adds an element to the specified hash table only if the table does not already contain the element.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashInsertIfAbsent(_:_:)
func NSHashInsertIfAbsent(table *NSHashTable, pointer unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNSHashInsertIfAbsent(table, pointer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHashInsertKnownAbsent func(table *NSHashTable, pointer unsafe.Pointer)
var _nSHashInsertKnownAbsentErr error

func tryNSHashInsertKnownAbsent(table *NSHashTable, pointer unsafe.Pointer) error {
	if _nSHashInsertKnownAbsent == nil {
		return symbolCallError("NSHashInsertKnownAbsent", "10.0", _nSHashInsertKnownAbsentErr)
	}
	_nSHashInsertKnownAbsent(table, pointer)
	return nil
}

// NSHashInsertKnownAbsent adds an element to the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashInsertKnownAbsent(_:_:)
func NSHashInsertKnownAbsent(table *NSHashTable, pointer unsafe.Pointer) {
	if callErr := tryNSHashInsertKnownAbsent(table, pointer); callErr != nil {
		panic(callErr)
	}
}

var _nSHashRemove func(table *NSHashTable, pointer unsafe.Pointer)
var _nSHashRemoveErr error

func tryNSHashRemove(table *NSHashTable, pointer unsafe.Pointer) error {
	if _nSHashRemove == nil {
		return symbolCallError("NSHashRemove", "10.0", _nSHashRemoveErr)
	}
	_nSHashRemove(table, pointer)
	return nil
}

// NSHashRemove removes an element from the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSHashRemove(_:_:)
func NSHashRemove(table *NSHashTable, pointer unsafe.Pointer) {
	if callErr := tryNSHashRemove(table, pointer); callErr != nil {
		panic(callErr)
	}
}

var _nSHomeDirectory func() NSString
var _nSHomeDirectoryErr error

func tryNSHomeDirectory() (NSString, error) {
	if _nSHomeDirectory == nil {
		return NSString{}, symbolCallError("NSHomeDirectory", "10.0", _nSHomeDirectoryErr)
	}
	return _nSHomeDirectory(), nil
}

// NSHomeDirectory returns the path to either the user’s or application’s home directory, depending on the platform.
//
// See: https://developer.apple.com/documentation/Foundation/NSHomeDirectory()
func NSHomeDirectory() NSString {
	result, callErr := tryNSHomeDirectory()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHomeDirectoryForUser func(userName NSString) NSString
var _nSHomeDirectoryForUserErr error

func tryNSHomeDirectoryForUser(userName NSString) (NSString, error) {
	if _nSHomeDirectoryForUser == nil {
		return NSString{}, symbolCallError("NSHomeDirectoryForUser", "10.0", _nSHomeDirectoryForUserErr)
	}
	return _nSHomeDirectoryForUser(userName), nil
}

// NSHomeDirectoryForUser returns the path to a given user’s home directory.
//
// See: https://developer.apple.com/documentation/Foundation/NSHomeDirectoryForUser(_:)
func NSHomeDirectoryForUser(userName NSString) NSString {
	result, callErr := tryNSHomeDirectoryForUser(userName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSIncrementExtraRefCount func(object objectivec.Object)
var _nSIncrementExtraRefCountErr error

func tryNSIncrementExtraRefCount(object objectivec.Object) error {
	if _nSIncrementExtraRefCount == nil {
		return symbolCallError("NSIncrementExtraRefCount", "10.0", _nSIncrementExtraRefCountErr)
	}
	_nSIncrementExtraRefCount(object)
	return nil
}

// NSIncrementExtraRefCount increments the specified object’s reference count.
//
// See: https://developer.apple.com/documentation/Foundation/NSIncrementExtraRefCount
func NSIncrementExtraRefCount(object objectivec.Object) {
	if callErr := tryNSIncrementExtraRefCount(object); callErr != nil {
		panic(callErr)
	}
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
var _nSIsFreedObjectErr error

func tryNSIsFreedObject(anObject objectivec.Object) (bool, error) {
	if _nSIsFreedObject == nil {
		return false, symbolCallError("NSIsFreedObject", "10.0", _nSIsFreedObjectErr)
	}
	return _nSIsFreedObject(anObject), nil
}

// NSIsFreedObject returns a Boolean indicating whether the specified object has been freed.
//
// See: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject objectivec.Object) bool {
	result, callErr := tryNSIsFreedObject(anObject)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSLog func(format NSString)
var _nSLogErr error

func tryNSLog(format NSString) error {
	if _nSLog == nil {
		return symbolCallError("NSLog", "10.0", _nSLogErr)
	}
	_nSLog(format)
	return nil
}

// NSLog logs an error message to the Apple System Log facility.
//
// See: https://developer.apple.com/documentation/Foundation/NSLog
func NSLog(format NSString) {
	if callErr := tryNSLog(format); callErr != nil {
		panic(callErr)
	}
}

var _nSLogPageSize func() uint
var _nSLogPageSizeErr error

func tryNSLogPageSize() (uint, error) {
	if _nSLogPageSize == nil {
		return 0, symbolCallError("NSLogPageSize", "10.0", _nSLogPageSizeErr)
	}
	return _nSLogPageSize(), nil
}

// NSLogPageSize returns the binary log of the page size.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogPageSize()
func NSLogPageSize() uint {
	result, callErr := tryNSLogPageSize()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSLogv func(format NSString, args kernel.Va_list)
var _nSLogvErr error

func tryNSLogv(format NSString, args kernel.Va_list) error {
	if _nSLogv == nil {
		return symbolCallError("NSLogv", "10.0", _nSLogvErr)
	}
	_nSLogv(format, args)
	return nil
}

// NSLogv logs an error message to the Apple System Log facility.
//
// See: https://developer.apple.com/documentation/Foundation/NSLogv(_:_:)
func NSLogv(format NSString, args kernel.Va_list) {
	if callErr := tryNSLogv(format, args); callErr != nil {
		panic(callErr)
	}
}

var _nSMapGet func(table *NSMapTable, key unsafe.Pointer) unsafe.Pointer
var _nSMapGetErr error

func tryNSMapGet(table *NSMapTable, key unsafe.Pointer) (unsafe.Pointer, error) {
	if _nSMapGet == nil {
		return nil, symbolCallError("NSMapGet", "10.0", _nSMapGetErr)
	}
	return _nSMapGet(table, key), nil
}

// NSMapGet returns a map table value for the specified key.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapGet(_:_:)
func NSMapGet(table *NSMapTable, key unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNSMapGet(table, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSMapInsert func(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer)
var _nSMapInsertErr error

func tryNSMapInsert(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) error {
	if _nSMapInsert == nil {
		return symbolCallError("NSMapInsert", "10.0", _nSMapInsertErr)
	}
	_nSMapInsert(table, key, value)
	return nil
}

// NSMapInsert inserts a key-value pair into the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapInsert(_:_:_:)
func NSMapInsert(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) {
	if callErr := tryNSMapInsert(table, key, value); callErr != nil {
		panic(callErr)
	}
}

var _nSMapInsertIfAbsent func(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer
var _nSMapInsertIfAbsentErr error

func tryNSMapInsertIfAbsent(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) (unsafe.Pointer, error) {
	if _nSMapInsertIfAbsent == nil {
		return nil, symbolCallError("NSMapInsertIfAbsent", "10.0", _nSMapInsertIfAbsentErr)
	}
	return _nSMapInsertIfAbsent(table, key, value), nil
}

// NSMapInsertIfAbsent inserts a key-value pair into the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapInsertIfAbsent(_:_:_:)
func NSMapInsertIfAbsent(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNSMapInsertIfAbsent(table, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSMapInsertKnownAbsent func(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer)
var _nSMapInsertKnownAbsentErr error

func tryNSMapInsertKnownAbsent(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) error {
	if _nSMapInsertKnownAbsent == nil {
		return symbolCallError("NSMapInsertKnownAbsent", "10.0", _nSMapInsertKnownAbsentErr)
	}
	_nSMapInsertKnownAbsent(table, key, value)
	return nil
}

// NSMapInsertKnownAbsent inserts a key-value pair into the specified table if the pair had not been previously added.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapInsertKnownAbsent(_:_:_:)
func NSMapInsertKnownAbsent(table *NSMapTable, key unsafe.Pointer, value unsafe.Pointer) {
	if callErr := tryNSMapInsertKnownAbsent(table, key, value); callErr != nil {
		panic(callErr)
	}
}

var _nSMapMember func(table *NSMapTable, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) bool
var _nSMapMemberErr error

func tryNSMapMember(table *NSMapTable, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) (bool, error) {
	if _nSMapMember == nil {
		return false, symbolCallError("NSMapMember", "10.0", _nSMapMemberErr)
	}
	return _nSMapMember(table, key, originalKey, value), nil
}

// NSMapMember indicates whether a given table contains a given key.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapMember(_:_:_:_:)
func NSMapMember(table *NSMapTable, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) bool {
	result, callErr := tryNSMapMember(table, key, originalKey, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSMapRemove func(table *NSMapTable, key unsafe.Pointer)
var _nSMapRemoveErr error

func tryNSMapRemove(table *NSMapTable, key unsafe.Pointer) error {
	if _nSMapRemove == nil {
		return symbolCallError("NSMapRemove", "10.0", _nSMapRemoveErr)
	}
	_nSMapRemove(table, key)
	return nil
}

// NSMapRemove removes a key and corresponding value from the specified table.
//
// See: https://developer.apple.com/documentation/Foundation/NSMapRemove(_:_:)
func NSMapRemove(table *NSMapTable, key unsafe.Pointer) {
	if callErr := tryNSMapRemove(table, key); callErr != nil {
		panic(callErr)
	}
}

// NSMouseInRect returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSMouseInRect(_:_:_:)
func NSMouseInRect(aPoint corefoundation.CGPoint, aRect corefoundation.CGRect, flipped bool) bool {
	return nsMouseInRect(aPoint, aRect, flipped)
}

var _nSNextHashEnumeratorItem func(enumerator *NSHashEnumerator) unsafe.Pointer
var _nSNextHashEnumeratorItemErr error

func tryNSNextHashEnumeratorItem(enumerator *NSHashEnumerator) (unsafe.Pointer, error) {
	if _nSNextHashEnumeratorItem == nil {
		return nil, symbolCallError("NSNextHashEnumeratorItem", "10.0", _nSNextHashEnumeratorItemErr)
	}
	return _nSNextHashEnumeratorItem(enumerator), nil
}

// NSNextHashEnumeratorItem returns the next hash-table element in the enumeration.
//
// See: https://developer.apple.com/documentation/Foundation/NSNextHashEnumeratorItem(_:)
func NSNextHashEnumeratorItem(enumerator *NSHashEnumerator) unsafe.Pointer {
	result, callErr := tryNSNextHashEnumeratorItem(enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSNextMapEnumeratorPair func(enumerator *NSMapEnumerator, key unsafe.Pointer, value unsafe.Pointer) bool
var _nSNextMapEnumeratorPairErr error

func tryNSNextMapEnumeratorPair(enumerator *NSMapEnumerator, key unsafe.Pointer, value unsafe.Pointer) (bool, error) {
	if _nSNextMapEnumeratorPair == nil {
		return false, symbolCallError("NSNextMapEnumeratorPair", "10.0", _nSNextMapEnumeratorPairErr)
	}
	return _nSNextMapEnumeratorPair(enumerator, key, value), nil
}

// NSNextMapEnumeratorPair returns a Boolean value that indicates whether the next map-table pair in the enumeration are set.
//
// See: https://developer.apple.com/documentation/Foundation/NSNextMapEnumeratorPair(_:_:_:)
func NSNextMapEnumeratorPair(enumerator *NSMapEnumerator, key unsafe.Pointer, value unsafe.Pointer) bool {
	result, callErr := tryNSNextMapEnumeratorPair(enumerator, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

// NSOffsetRect offsets the rectangle by the specified amount.
//
// See: https://developer.apple.com/documentation/Foundation/NSOffsetRect(_:_:_:)
func NSOffsetRect(aRect corefoundation.CGRect, dX float64, dY float64) corefoundation.CGRect {
	return nsOffsetRect(aRect, dX, dY)
}

var _nSOpenStepRootDirectory func() NSString
var _nSOpenStepRootDirectoryErr error

func tryNSOpenStepRootDirectory() (NSString, error) {
	if _nSOpenStepRootDirectory == nil {
		return NSString{}, symbolCallError("NSOpenStepRootDirectory", "10.0", _nSOpenStepRootDirectoryErr)
	}
	return _nSOpenStepRootDirectory(), nil
}

// NSOpenStepRootDirectory returns the root directory of the user’s system.
//
// See: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() NSString {
	result, callErr := tryNSOpenStepRootDirectory()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSPageSize func() uint
var _nSPageSizeErr error

func tryNSPageSize() (uint, error) {
	if _nSPageSize == nil {
		return 0, symbolCallError("NSPageSize", "10.0", _nSPageSizeErr)
	}
	return _nSPageSize(), nil
}

// NSPageSize returns the number of bytes in a page.
//
// See: https://developer.apple.com/documentation/Foundation/NSPageSize()
func NSPageSize() uint {
	result, callErr := tryNSPageSize()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSPointFromString func(aString NSString) corefoundation.CGPoint
var _nSPointFromStringErr error

func tryNSPointFromString(aString NSString) (corefoundation.CGPoint, error) {
	if _nSPointFromString == nil {
		return corefoundation.CGPoint{}, symbolCallError("NSPointFromString", "10.0", _nSPointFromStringErr)
	}
	return _nSPointFromString(aString), nil
}

// NSPointFromString returns a point from a text-based representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointFromString(_:)
func NSPointFromString(aString NSString) corefoundation.CGPoint {
	result, callErr := tryNSPointFromString(aString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

// NSPointInRect returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointInRect(_:_:)
func NSPointInRect(aPoint corefoundation.CGPoint, aRect corefoundation.CGRect) bool {
	return nsPointInRect(aPoint, aRect)
}

var _nSProtocolFromString func(namestr NSString) **objectivec.Protocol
var _nSProtocolFromStringErr error

func tryNSProtocolFromString(namestr NSString) (**objectivec.Protocol, error) {
	if _nSProtocolFromString == nil {
		return nil, symbolCallError("NSProtocolFromString", "10.5", _nSProtocolFromStringErr)
	}
	return _nSProtocolFromString(namestr), nil
}

// NSProtocolFromString returns a the protocol with a given name.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolFromString(_:)
func NSProtocolFromString(namestr NSString) **objectivec.Protocol {
	result, callErr := tryNSProtocolFromString(namestr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRangeFromString func(aString NSString) NSRange
var _nSRangeFromStringErr error

func tryNSRangeFromString(aString NSString) (NSRange, error) {
	if _nSRangeFromString == nil {
		return NSRange{}, symbolCallError("NSRangeFromString", "10.0", _nSRangeFromStringErr)
	}
	return _nSRangeFromString(aString), nil
}

// NSRangeFromString returns a range from a textual representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSRangeFromString(_:)
func NSRangeFromString(aString NSString) NSRange {
	result, callErr := tryNSRangeFromString(aString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSReallocateCollectable func(ptr unsafe.Pointer, size uint, options uint) unsafe.Pointer
var _nSReallocateCollectableErr error

func tryNSReallocateCollectable(ptr unsafe.Pointer, size uint, options uint) (unsafe.Pointer, error) {
	if _nSReallocateCollectable == nil {
		return nil, symbolCallError("NSReallocateCollectable", "10.0", _nSReallocateCollectableErr)
	}
	return _nSReallocateCollectable(ptr, size, options), nil
}

// NSReallocateCollectable reallocates collectable memory.
//
// Deprecated: Garbage collection is deprecated in OS X v10.8; instead,you should use AutomaticReference Counting—see [Transitioning to ARC Release Notes](<https://developer.apple.com/library/archive/releasenotes/ObjectiveC/RN-TransitioningToARC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011226>).
//
// See: https://developer.apple.com/documentation/Foundation/NSReallocateCollectable
func NSReallocateCollectable(ptr unsafe.Pointer, size uint, options uint) unsafe.Pointer {
	result, callErr := tryNSReallocateCollectable(ptr, size, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRecordAllocationEvent func(eventType int, object objectivec.Object)
var _nSRecordAllocationEventErr error

func tryNSRecordAllocationEvent(eventType int, object objectivec.Object) error {
	if _nSRecordAllocationEvent == nil {
		return symbolCallError("NSRecordAllocationEvent", "10.0", _nSRecordAllocationEventErr)
	}
	_nSRecordAllocationEvent(eventType, object)
	return nil
}

// NSRecordAllocationEvent notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// See: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType int, object objectivec.Object) {
	if callErr := tryNSRecordAllocationEvent(eventType, object); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFromString func(aString NSString) corefoundation.CGRect
var _nSRectFromStringErr error

func tryNSRectFromString(aString NSString) (corefoundation.CGRect, error) {
	if _nSRectFromString == nil {
		return corefoundation.CGRect{}, symbolCallError("NSRectFromString", "10.0", _nSRectFromStringErr)
	}
	return _nSRectFromString(aString), nil
}

// NSRectFromString returns a rectangle from a text-based representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSRectFromString(_:)
func NSRectFromString(aString NSString) corefoundation.CGRect {
	result, callErr := tryNSRectFromString(aString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRecycleZone func(zone *NSZone)
var _nSRecycleZoneErr error

func tryNSRecycleZone(zone *NSZone) error {
	if _nSRecycleZone == nil {
		return symbolCallError("NSRecycleZone", "10.0", _nSRecycleZoneErr)
	}
	_nSRecycleZone(zone)
	return nil
}

// NSRecycleZone frees memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSRecycleZone
func NSRecycleZone(zone *NSZone) {
	if callErr := tryNSRecycleZone(zone); callErr != nil {
		panic(callErr)
	}
}

var _nSResetHashTable func(table *NSHashTable)
var _nSResetHashTableErr error

func tryNSResetHashTable(table *NSHashTable) error {
	if _nSResetHashTable == nil {
		return symbolCallError("NSResetHashTable", "10.0", _nSResetHashTableErr)
	}
	_nSResetHashTable(table)
	return nil
}

// NSResetHashTable deletes the elements of the specified hash table.
//
// See: https://developer.apple.com/documentation/Foundation/NSResetHashTable(_:)
func NSResetHashTable(table *NSHashTable) {
	if callErr := tryNSResetHashTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nSResetMapTable func(table *NSMapTable)
var _nSResetMapTableErr error

func tryNSResetMapTable(table *NSMapTable) error {
	if _nSResetMapTable == nil {
		return symbolCallError("NSResetMapTable", "10.0", _nSResetMapTableErr)
	}
	_nSResetMapTable(table)
	return nil
}

// NSResetMapTable deletes the elements of the specified map table.
//
// See: https://developer.apple.com/documentation/Foundation/NSResetMapTable(_:)
func NSResetMapTable(table *NSMapTable) {
	if callErr := tryNSResetMapTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nSReturnAddress func(frame uint) unsafe.Pointer
var _nSReturnAddressErr error

func tryNSReturnAddress(frame uint) (unsafe.Pointer, error) {
	if _nSReturnAddress == nil {
		return nil, symbolCallError("NSReturnAddress", "10.0", _nSReturnAddressErr)
	}
	return _nSReturnAddress(frame), nil
}

// NSReturnAddress returns the value of the return address of the specified frame.
//
// See: https://developer.apple.com/documentation/Foundation/NSReturnAddress
func NSReturnAddress(frame uint) unsafe.Pointer {
	result, callErr := tryNSReturnAddress(frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRoundDownToMultipleOfPageSize func(bytes uint) uint
var _nSRoundDownToMultipleOfPageSizeErr error

func tryNSRoundDownToMultipleOfPageSize(bytes uint) (uint, error) {
	if _nSRoundDownToMultipleOfPageSize == nil {
		return 0, symbolCallError("NSRoundDownToMultipleOfPageSize", "10.0", _nSRoundDownToMultipleOfPageSizeErr)
	}
	return _nSRoundDownToMultipleOfPageSize(bytes), nil
}

// NSRoundDownToMultipleOfPageSize returns the specified number of bytes rounded down to a multiple of the page size.
//
// See: https://developer.apple.com/documentation/Foundation/NSRoundDownToMultipleOfPageSize(_:)
func NSRoundDownToMultipleOfPageSize(bytes uint) uint {
	result, callErr := tryNSRoundDownToMultipleOfPageSize(bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRoundUpToMultipleOfPageSize func(bytes uint) uint
var _nSRoundUpToMultipleOfPageSizeErr error

func tryNSRoundUpToMultipleOfPageSize(bytes uint) (uint, error) {
	if _nSRoundUpToMultipleOfPageSize == nil {
		return 0, symbolCallError("NSRoundUpToMultipleOfPageSize", "10.0", _nSRoundUpToMultipleOfPageSizeErr)
	}
	return _nSRoundUpToMultipleOfPageSize(bytes), nil
}

// NSRoundUpToMultipleOfPageSize returns the specified number of bytes rounded up to a multiple of the page size.
//
// See: https://developer.apple.com/documentation/Foundation/NSRoundUpToMultipleOfPageSize(_:)
func NSRoundUpToMultipleOfPageSize(bytes uint) uint {
	result, callErr := tryNSRoundUpToMultipleOfPageSize(bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSSearchPathForDirectoriesInDomains func(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, expandTilde bool) []NSString
var _nSSearchPathForDirectoriesInDomainsErr error

func tryNSSearchPathForDirectoriesInDomains(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, expandTilde bool) ([]NSString, error) {
	if _nSSearchPathForDirectoriesInDomains == nil {
		return nil, symbolCallError("NSSearchPathForDirectoriesInDomains", "10.0", _nSSearchPathForDirectoriesInDomainsErr)
	}
	return _nSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde), nil
}

// NSSearchPathForDirectoriesInDomains creates a list of directory search paths.
//
// See: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory NSSearchPathDirectory, domainMask NSSearchPathDomainMask, expandTilde bool) []NSString {
	result, callErr := tryNSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSSelectorFromString func(aSelectorName NSString) objectivec.SEL
var _nSSelectorFromStringErr error

func tryNSSelectorFromString(aSelectorName NSString) (objectivec.SEL, error) {
	if _nSSelectorFromString == nil {
		return *new(objectivec.SEL), symbolCallError("NSSelectorFromString", "10.0", _nSSelectorFromStringErr)
	}
	return _nSSelectorFromString(aSelectorName), nil
}

// NSSelectorFromString returns the selector with a given name.
//
// See: https://developer.apple.com/documentation/Foundation/NSSelectorFromString(_:)
func NSSelectorFromString(aSelectorName NSString) objectivec.SEL {
	result, callErr := tryNSSelectorFromString(aSelectorName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSSetUncaughtExceptionHandler func(arg0 func(*NSException))
var _nSSetUncaughtExceptionHandlerErr error

func tryNSSetUncaughtExceptionHandler(arg0 func(*NSException)) error {
	if _nSSetUncaughtExceptionHandler == nil {
		return symbolCallError("NSSetUncaughtExceptionHandler", "10.0", _nSSetUncaughtExceptionHandlerErr)
	}
	_nSSetUncaughtExceptionHandler(arg0)
	return nil
}

// NSSetUncaughtExceptionHandler changes the top-level error handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetUncaughtExceptionHandler(_:)
func NSSetUncaughtExceptionHandler(arg0 func(*NSException)) {
	if callErr := tryNSSetUncaughtExceptionHandler(arg0); callErr != nil {
		panic(callErr)
	}
}

var _nSSetZoneName func(zone *NSZone, name NSString)
var _nSSetZoneNameErr error

func tryNSSetZoneName(zone *NSZone, name NSString) error {
	if _nSSetZoneName == nil {
		return symbolCallError("NSSetZoneName", "10.0", _nSSetZoneNameErr)
	}
	_nSSetZoneName(zone, name)
	return nil
}

// NSSetZoneName sets the name of the specified zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetZoneName
func NSSetZoneName(zone *NSZone, name NSString) {
	if callErr := tryNSSetZoneName(zone, name); callErr != nil {
		panic(callErr)
	}
}

var _nSShouldRetainWithZone func(anObject objectivec.Object, requestedZone *NSZone) bool
var _nSShouldRetainWithZoneErr error

func tryNSShouldRetainWithZone(anObject objectivec.Object, requestedZone *NSZone) (bool, error) {
	if _nSShouldRetainWithZone == nil {
		return false, symbolCallError("NSShouldRetainWithZone", "10.0", _nSShouldRetainWithZoneErr)
	}
	return _nSShouldRetainWithZone(anObject, requestedZone), nil
}

// NSShouldRetainWithZone indicates whether an object should be retained.
//
// See: https://developer.apple.com/documentation/Foundation/NSShouldRetainWithZone
func NSShouldRetainWithZone(anObject objectivec.Object, requestedZone *NSZone) bool {
	result, callErr := tryNSShouldRetainWithZone(anObject, requestedZone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSSizeFromString func(aString NSString) corefoundation.CGSize
var _nSSizeFromStringErr error

func tryNSSizeFromString(aString NSString) (corefoundation.CGSize, error) {
	if _nSSizeFromString == nil {
		return corefoundation.CGSize{}, symbolCallError("NSSizeFromString", "10.0", _nSSizeFromStringErr)
	}
	return _nSSizeFromString(aString), nil
}

// NSSizeFromString returns an [NSSize] from a text-based representation.
//
// See: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString NSString) corefoundation.CGSize {
	result, callErr := tryNSSizeFromString(aString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromClass func(aClass objc.Class) NSString
var _nSStringFromClassErr error

func tryNSStringFromClass(aClass objc.Class) (NSString, error) {
	if _nSStringFromClass == nil {
		return NSString{}, symbolCallError("NSStringFromClass", "10.0", _nSStringFromClassErr)
	}
	return _nSStringFromClass(aClass), nil
}

// NSStringFromClass returns the name of a class as a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromClass(_:)
func NSStringFromClass(aClass objc.Class) NSString {
	result, callErr := tryNSStringFromClass(aClass)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromHashTable func(table *NSHashTable) NSString
var _nSStringFromHashTableErr error

func tryNSStringFromHashTable(table *NSHashTable) (NSString, error) {
	if _nSStringFromHashTable == nil {
		return NSString{}, symbolCallError("NSStringFromHashTable", "10.0", _nSStringFromHashTableErr)
	}
	return _nSStringFromHashTable(table), nil
}

// NSStringFromHashTable returns a string describing the hash table’s contents.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromHashTable(_:)
func NSStringFromHashTable(table *NSHashTable) NSString {
	result, callErr := tryNSStringFromHashTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromMapTable func(table *NSMapTable) NSString
var _nSStringFromMapTableErr error

func tryNSStringFromMapTable(table *NSMapTable) (NSString, error) {
	if _nSStringFromMapTable == nil {
		return NSString{}, symbolCallError("NSStringFromMapTable", "10.0", _nSStringFromMapTableErr)
	}
	return _nSStringFromMapTable(table), nil
}

// NSStringFromMapTable returns a string describing the map table’s contents.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromMapTable(_:)
func NSStringFromMapTable(table *NSMapTable) NSString {
	result, callErr := tryNSStringFromMapTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromPoint func(aPoint corefoundation.CGPoint) NSString
var _nSStringFromPointErr error

func tryNSStringFromPoint(aPoint corefoundation.CGPoint) (NSString, error) {
	if _nSStringFromPoint == nil {
		return NSString{}, symbolCallError("NSStringFromPoint", "10.0", _nSStringFromPointErr)
	}
	return _nSStringFromPoint(aPoint), nil
}

// NSStringFromPoint returns a string representation of a point.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromPoint(_:)
func NSStringFromPoint(aPoint corefoundation.CGPoint) NSString {
	result, callErr := tryNSStringFromPoint(aPoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromProtocol func(proto **objectivec.Protocol) NSString
var _nSStringFromProtocolErr error

func tryNSStringFromProtocol(proto **objectivec.Protocol) (NSString, error) {
	if _nSStringFromProtocol == nil {
		return NSString{}, symbolCallError("NSStringFromProtocol", "10.5", _nSStringFromProtocolErr)
	}
	return _nSStringFromProtocol(proto), nil
}

// NSStringFromProtocol returns the name of a protocol as a string.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromProtocol(_:)
func NSStringFromProtocol(proto **objectivec.Protocol) NSString {
	result, callErr := tryNSStringFromProtocol(proto)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromRange func(range_ NSRange) NSString
var _nSStringFromRangeErr error

func tryNSStringFromRange(range_ NSRange) (NSString, error) {
	if _nSStringFromRange == nil {
		return NSString{}, symbolCallError("NSStringFromRange", "10.0", _nSStringFromRangeErr)
	}
	return _nSStringFromRange(range_), nil
}

// NSStringFromRange returns a string representation of a range.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromRange(_:)
func NSStringFromRange(range_ NSRange) NSString {
	result, callErr := tryNSStringFromRange(range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromRect func(aRect corefoundation.CGRect) NSString
var _nSStringFromRectErr error

func tryNSStringFromRect(aRect corefoundation.CGRect) (NSString, error) {
	if _nSStringFromRect == nil {
		return NSString{}, symbolCallError("NSStringFromRect", "10.0", _nSStringFromRectErr)
	}
	return _nSStringFromRect(aRect), nil
}

// NSStringFromRect returns a string representation of a rectangle.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromRect(_:)
func NSStringFromRect(aRect corefoundation.CGRect) NSString {
	result, callErr := tryNSStringFromRect(aRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromSelector func(aSelector objectivec.SEL) NSString
var _nSStringFromSelectorErr error

func tryNSStringFromSelector(aSelector objectivec.SEL) (NSString, error) {
	if _nSStringFromSelector == nil {
		return NSString{}, symbolCallError("NSStringFromSelector", "10.0", _nSStringFromSelectorErr)
	}
	return _nSStringFromSelector(aSelector), nil
}

// NSStringFromSelector returns a string representation of a given selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromSelector(_:)
func NSStringFromSelector(aSelector objectivec.SEL) NSString {
	result, callErr := tryNSStringFromSelector(aSelector)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSStringFromSize func(aSize corefoundation.CGSize) NSString
var _nSStringFromSizeErr error

func tryNSStringFromSize(aSize corefoundation.CGSize) (NSString, error) {
	if _nSStringFromSize == nil {
		return NSString{}, symbolCallError("NSStringFromSize", "10.0", _nSStringFromSizeErr)
	}
	return _nSStringFromSize(aSize), nil
}

// NSStringFromSize returns a string representation of a size.
//
// See: https://developer.apple.com/documentation/Foundation/NSStringFromSize(_:)
func NSStringFromSize(aSize corefoundation.CGSize) NSString {
	result, callErr := tryNSStringFromSize(aSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSTemporaryDirectory func() NSString
var _nSTemporaryDirectoryErr error

func tryNSTemporaryDirectory() (NSString, error) {
	if _nSTemporaryDirectory == nil {
		return NSString{}, symbolCallError("NSTemporaryDirectory", "10.0", _nSTemporaryDirectoryErr)
	}
	return _nSTemporaryDirectory(), nil
}

// NSTemporaryDirectory returns the path of the temporary directory for the current user.
//
// See: https://developer.apple.com/documentation/Foundation/NSTemporaryDirectory()
func NSTemporaryDirectory() NSString {
	result, callErr := tryNSTemporaryDirectory()
	if callErr != nil {
		panic(callErr)
	}
	return result
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
var _nSUserNameErr error

func tryNSUserName() (NSString, error) {
	if _nSUserName == nil {
		return NSString{}, symbolCallError("NSUserName", "10.0", _nSUserNameErr)
	}
	return _nSUserName(), nil
}

// NSUserName returns the logon name of the current user.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserName()
func NSUserName() NSString {
	result, callErr := tryNSUserName()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSZoneCalloc func(zone *NSZone, numElems uint, byteSize uint) unsafe.Pointer
var _nSZoneCallocErr error

func tryNSZoneCalloc(zone *NSZone, numElems uint, byteSize uint) (unsafe.Pointer, error) {
	if _nSZoneCalloc == nil {
		return nil, symbolCallError("NSZoneCalloc", "10.0", _nSZoneCallocErr)
	}
	return _nSZoneCalloc(zone, numElems, byteSize), nil
}

// NSZoneCalloc allocates memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneCalloc
func NSZoneCalloc(zone *NSZone, numElems uint, byteSize uint) unsafe.Pointer {
	result, callErr := tryNSZoneCalloc(zone, numElems, byteSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSZoneFree func(zone *NSZone, ptr unsafe.Pointer)
var _nSZoneFreeErr error

func tryNSZoneFree(zone *NSZone, ptr unsafe.Pointer) error {
	if _nSZoneFree == nil {
		return symbolCallError("NSZoneFree", "10.0", _nSZoneFreeErr)
	}
	_nSZoneFree(zone, ptr)
	return nil
}

// NSZoneFree deallocates a block of memory in the specified zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneFree
func NSZoneFree(zone *NSZone, ptr unsafe.Pointer) {
	if callErr := tryNSZoneFree(zone, ptr); callErr != nil {
		panic(callErr)
	}
}

var _nSZoneFromPointer func(ptr unsafe.Pointer) *NSZone
var _nSZoneFromPointerErr error

func tryNSZoneFromPointer(ptr unsafe.Pointer) (*NSZone, error) {
	if _nSZoneFromPointer == nil {
		return nil, symbolCallError("NSZoneFromPointer", "10.0", _nSZoneFromPointerErr)
	}
	return _nSZoneFromPointer(ptr), nil
}

// NSZoneFromPointer gets the zone for a given block of memory.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneFromPointer
func NSZoneFromPointer(ptr unsafe.Pointer) *NSZone {
	result, callErr := tryNSZoneFromPointer(ptr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSZoneMalloc func(zone *NSZone, size uint) unsafe.Pointer
var _nSZoneMallocErr error

func tryNSZoneMalloc(zone *NSZone, size uint) (unsafe.Pointer, error) {
	if _nSZoneMalloc == nil {
		return nil, symbolCallError("NSZoneMalloc", "10.0", _nSZoneMallocErr)
	}
	return _nSZoneMalloc(zone, size), nil
}

// NSZoneMalloc allocates memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneMalloc
func NSZoneMalloc(zone *NSZone, size uint) unsafe.Pointer {
	result, callErr := tryNSZoneMalloc(zone, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSZoneName func(zone *NSZone) NSString
var _nSZoneNameErr error

func tryNSZoneName(zone *NSZone) (NSString, error) {
	if _nSZoneName == nil {
		return NSString{}, symbolCallError("NSZoneName", "10.0", _nSZoneNameErr)
	}
	return _nSZoneName(zone), nil
}

// NSZoneName returns the name of the specified zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneName
func NSZoneName(zone *NSZone) NSString {
	result, callErr := tryNSZoneName(zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSZoneRealloc func(zone *NSZone, ptr unsafe.Pointer, size uint) unsafe.Pointer
var _nSZoneReallocErr error

func tryNSZoneRealloc(zone *NSZone, ptr unsafe.Pointer, size uint) (unsafe.Pointer, error) {
	if _nSZoneRealloc == nil {
		return nil, symbolCallError("NSZoneRealloc", "10.0", _nSZoneReallocErr)
	}
	return _nSZoneRealloc(zone, ptr, size), nil
}

// NSZoneRealloc allocates memory in a zone.
//
// Deprecated: Zones are ignored on iOS and 64-bit runtime in macOS. You should not use zones in current development.
//
// See: https://developer.apple.com/documentation/Foundation/NSZoneRealloc
func NSZoneRealloc(zone *NSZone, ptr unsafe.Pointer, size uint) unsafe.Pointer {
	result, callErr := tryNSZoneRealloc(zone, ptr, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXReadNSObjectFromCoder func(decoder *NSCoder) *objectivec.Object
var _nXReadNSObjectFromCoderErr error

func tryNXReadNSObjectFromCoder(decoder *NSCoder) (*objectivec.Object, error) {
	if _nXReadNSObjectFromCoder == nil {
		return nil, symbolCallError("NXReadNSObjectFromCoder", "10.0", _nXReadNSObjectFromCoderErr)
	}
	return _nXReadNSObjectFromCoder(decoder), nil
}

// NXReadNSObjectFromCoder returns the next object from the coder.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/Foundation/NXReadNSObjectFromCoder
func NXReadNSObjectFromCoder(decoder *NSCoder) *objectivec.Object {
	result, callErr := tryNXReadNSObjectFromCoder(decoder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_nSAllHashTableObjects, &_nSAllHashTableObjectsErr, frameworkHandle, "NSAllHashTableObjects", "10.0")
	registerFunc(&_nSAllMapTableKeys, &_nSAllMapTableKeysErr, frameworkHandle, "NSAllMapTableKeys", "10.0")
	registerFunc(&_nSAllMapTableValues, &_nSAllMapTableValuesErr, frameworkHandle, "NSAllMapTableValues", "10.0")
	registerFunc(&_nSAllocateCollectable, &_nSAllocateCollectableErr, frameworkHandle, "NSAllocateCollectable", "10.0")
	registerFunc(&_nSAllocateMemoryPages, &_nSAllocateMemoryPagesErr, frameworkHandle, "NSAllocateMemoryPages", "10.0")
	registerFunc(&_nSAllocateObject, &_nSAllocateObjectErr, frameworkHandle, "NSAllocateObject", "10.0")
	registerFunc(&_nSClassFromString, &_nSClassFromStringErr, frameworkHandle, "NSClassFromString", "10.0")
	registerFunc(&_nSCompareHashTables, &_nSCompareHashTablesErr, frameworkHandle, "NSCompareHashTables", "10.0")
	registerFunc(&_nSCompareMapTables, &_nSCompareMapTablesErr, frameworkHandle, "NSCompareMapTables", "10.0")
	registerFunc(&_nSCopyHashTableWithZone, &_nSCopyHashTableWithZoneErr, frameworkHandle, "NSCopyHashTableWithZone", "10.0")
	registerFunc(&_nSCopyMapTableWithZone, &_nSCopyMapTableWithZoneErr, frameworkHandle, "NSCopyMapTableWithZone", "10.0")
	registerFunc(&_nSCopyMemoryPages, &_nSCopyMemoryPagesErr, frameworkHandle, "NSCopyMemoryPages", "10.0")
	registerFunc(&_nSCopyObject, &_nSCopyObjectErr, frameworkHandle, "NSCopyObject", "10.0")
	registerFunc(&_nSCountFrames, &_nSCountFramesErr, frameworkHandle, "NSCountFrames", "10.0")
	registerFunc(&_nSCountHashTable, &_nSCountHashTableErr, frameworkHandle, "NSCountHashTable", "10.0")
	registerFunc(&_nSCountMapTable, &_nSCountMapTableErr, frameworkHandle, "NSCountMapTable", "10.0")
	registerFunc(&_nSCreateHashTable, &_nSCreateHashTableErr, frameworkHandle, "NSCreateHashTable", "10.0")
	registerFunc(&_nSCreateHashTableWithZone, &_nSCreateHashTableWithZoneErr, frameworkHandle, "NSCreateHashTableWithZone", "10.0")
	registerFunc(&_nSCreateMapTable, &_nSCreateMapTableErr, frameworkHandle, "NSCreateMapTable", "10.0")
	registerFunc(&_nSCreateMapTableWithZone, &_nSCreateMapTableWithZoneErr, frameworkHandle, "NSCreateMapTableWithZone", "10.0")
	registerFunc(&_nSCreateZone, &_nSCreateZoneErr, frameworkHandle, "NSCreateZone", "10.0")
	registerFunc(&_nSDeallocateMemoryPages, &_nSDeallocateMemoryPagesErr, frameworkHandle, "NSDeallocateMemoryPages", "10.0")
	registerFunc(&_nSDeallocateObject, &_nSDeallocateObjectErr, frameworkHandle, "NSDeallocateObject", "10.0")
	registerFunc(&_nSDecimalAdd, &_nSDecimalAddErr, frameworkHandle, "NSDecimalAdd", "10.0")
	registerFunc(&_nSDecimalCompact, &_nSDecimalCompactErr, frameworkHandle, "NSDecimalCompact", "10.0")
	registerFunc(&_nSDecimalCompare, &_nSDecimalCompareErr, frameworkHandle, "NSDecimalCompare", "10.0")
	registerFunc(&_nSDecimalCopy, &_nSDecimalCopyErr, frameworkHandle, "NSDecimalCopy", "10.0")
	registerFunc(&_nSDecimalDivide, &_nSDecimalDivideErr, frameworkHandle, "NSDecimalDivide", "10.0")
	registerFunc(&_nSDecimalMultiply, &_nSDecimalMultiplyErr, frameworkHandle, "NSDecimalMultiply", "10.0")
	registerFunc(&_nSDecimalMultiplyByPowerOf10, &_nSDecimalMultiplyByPowerOf10Err, frameworkHandle, "NSDecimalMultiplyByPowerOf10", "10.0")
	registerFunc(&_nSDecimalNormalize, &_nSDecimalNormalizeErr, frameworkHandle, "NSDecimalNormalize", "10.0")
	registerFunc(&_nSDecimalPower, &_nSDecimalPowerErr, frameworkHandle, "NSDecimalPower", "10.0")
	registerFunc(&_nSDecimalRound, &_nSDecimalRoundErr, frameworkHandle, "NSDecimalRound", "10.0")
	registerFunc(&_nSDecimalString, &_nSDecimalStringErr, frameworkHandle, "NSDecimalString", "10.0")
	registerFunc(&_nSDecimalSubtract, &_nSDecimalSubtractErr, frameworkHandle, "NSDecimalSubtract", "10.0")
	registerFunc(&_nSDecrementExtraRefCountWasZero, &_nSDecrementExtraRefCountWasZeroErr, frameworkHandle, "NSDecrementExtraRefCountWasZero", "10.0")
	registerFunc(&_nSDefaultMallocZone, &_nSDefaultMallocZoneErr, frameworkHandle, "NSDefaultMallocZone", "10.0")
	registerFunc(&_nSEndHashTableEnumeration, &_nSEndHashTableEnumerationErr, frameworkHandle, "NSEndHashTableEnumeration", "10.0")
	registerFunc(&_nSEndMapTableEnumeration, &_nSEndMapTableEnumerationErr, frameworkHandle, "NSEndMapTableEnumeration", "10.0")
	registerFunc(&_nSEnumerateHashTable, &_nSEnumerateHashTableErr, frameworkHandle, "NSEnumerateHashTable", "10.0")
	registerFunc(&_nSEnumerateMapTable, &_nSEnumerateMapTableErr, frameworkHandle, "NSEnumerateMapTable", "10.0")
	registerFunc(&_nSExtraRefCount, &_nSExtraRefCountErr, frameworkHandle, "NSExtraRefCount", "10.0")
	registerFunc(&_nSFileTypeForHFSTypeCode, &_nSFileTypeForHFSTypeCodeErr, frameworkHandle, "NSFileTypeForHFSTypeCode", "10.0")
	registerFunc(&_nSFrameAddress, &_nSFrameAddressErr, frameworkHandle, "NSFrameAddress", "10.0")
	registerFunc(&_nSFreeHashTable, &_nSFreeHashTableErr, frameworkHandle, "NSFreeHashTable", "10.0")
	registerFunc(&_nSFreeMapTable, &_nSFreeMapTableErr, frameworkHandle, "NSFreeMapTable", "10.0")
	registerFunc(&_nSFullUserName, &_nSFullUserNameErr, frameworkHandle, "NSFullUserName", "10.0")
	registerFunc(&_nSGetSizeAndAlignment, &_nSGetSizeAndAlignmentErr, frameworkHandle, "NSGetSizeAndAlignment", "10.0")
	registerFunc(&_nSGetUncaughtExceptionHandler, &_nSGetUncaughtExceptionHandlerErr, frameworkHandle, "NSGetUncaughtExceptionHandler", "10.0")
	registerFunc(&_nSHFSTypeCodeFromFileType, &_nSHFSTypeCodeFromFileTypeErr, frameworkHandle, "NSHFSTypeCodeFromFileType", "10.0")
	registerFunc(&_nSHFSTypeOfFile, &_nSHFSTypeOfFileErr, frameworkHandle, "NSHFSTypeOfFile", "10.0")
	registerFunc(&_nSHashGet, &_nSHashGetErr, frameworkHandle, "NSHashGet", "10.0")
	registerFunc(&_nSHashInsert, &_nSHashInsertErr, frameworkHandle, "NSHashInsert", "10.0")
	registerFunc(&_nSHashInsertIfAbsent, &_nSHashInsertIfAbsentErr, frameworkHandle, "NSHashInsertIfAbsent", "10.0")
	registerFunc(&_nSHashInsertKnownAbsent, &_nSHashInsertKnownAbsentErr, frameworkHandle, "NSHashInsertKnownAbsent", "10.0")
	registerFunc(&_nSHashRemove, &_nSHashRemoveErr, frameworkHandle, "NSHashRemove", "10.0")
	registerFunc(&_nSHomeDirectory, &_nSHomeDirectoryErr, frameworkHandle, "NSHomeDirectory", "10.0")
	registerFunc(&_nSHomeDirectoryForUser, &_nSHomeDirectoryForUserErr, frameworkHandle, "NSHomeDirectoryForUser", "10.0")
	registerFunc(&_nSIncrementExtraRefCount, &_nSIncrementExtraRefCountErr, frameworkHandle, "NSIncrementExtraRefCount", "10.0")
	registerFunc(&_nSIsFreedObject, &_nSIsFreedObjectErr, frameworkHandle, "NSIsFreedObject", "10.0")
	registerFunc(&_nSLog, &_nSLogErr, frameworkHandle, "NSLog", "10.0")
	registerFunc(&_nSLogPageSize, &_nSLogPageSizeErr, frameworkHandle, "NSLogPageSize", "10.0")
	registerFunc(&_nSLogv, &_nSLogvErr, frameworkHandle, "NSLogv", "10.0")
	registerFunc(&_nSMapGet, &_nSMapGetErr, frameworkHandle, "NSMapGet", "10.0")
	registerFunc(&_nSMapInsert, &_nSMapInsertErr, frameworkHandle, "NSMapInsert", "10.0")
	registerFunc(&_nSMapInsertIfAbsent, &_nSMapInsertIfAbsentErr, frameworkHandle, "NSMapInsertIfAbsent", "10.0")
	registerFunc(&_nSMapInsertKnownAbsent, &_nSMapInsertKnownAbsentErr, frameworkHandle, "NSMapInsertKnownAbsent", "10.0")
	registerFunc(&_nSMapMember, &_nSMapMemberErr, frameworkHandle, "NSMapMember", "10.0")
	registerFunc(&_nSMapRemove, &_nSMapRemoveErr, frameworkHandle, "NSMapRemove", "10.0")
	registerFunc(&_nSNextHashEnumeratorItem, &_nSNextHashEnumeratorItemErr, frameworkHandle, "NSNextHashEnumeratorItem", "10.0")
	registerFunc(&_nSNextMapEnumeratorPair, &_nSNextMapEnumeratorPairErr, frameworkHandle, "NSNextMapEnumeratorPair", "10.0")
	registerFunc(&_nSOpenStepRootDirectory, &_nSOpenStepRootDirectoryErr, frameworkHandle, "NSOpenStepRootDirectory", "10.0")
	registerFunc(&_nSPageSize, &_nSPageSizeErr, frameworkHandle, "NSPageSize", "10.0")
	registerFunc(&_nSPointFromString, &_nSPointFromStringErr, frameworkHandle, "NSPointFromString", "10.0")
	registerFunc(&_nSProtocolFromString, &_nSProtocolFromStringErr, frameworkHandle, "NSProtocolFromString", "10.5")
	registerFunc(&_nSRangeFromString, &_nSRangeFromStringErr, frameworkHandle, "NSRangeFromString", "10.0")
	registerFunc(&_nSReallocateCollectable, &_nSReallocateCollectableErr, frameworkHandle, "NSReallocateCollectable", "10.0")
	registerFunc(&_nSRecordAllocationEvent, &_nSRecordAllocationEventErr, frameworkHandle, "NSRecordAllocationEvent", "10.0")
	registerFunc(&_nSRectFromString, &_nSRectFromStringErr, frameworkHandle, "NSRectFromString", "10.0")
	registerFunc(&_nSRecycleZone, &_nSRecycleZoneErr, frameworkHandle, "NSRecycleZone", "10.0")
	registerFunc(&_nSResetHashTable, &_nSResetHashTableErr, frameworkHandle, "NSResetHashTable", "10.0")
	registerFunc(&_nSResetMapTable, &_nSResetMapTableErr, frameworkHandle, "NSResetMapTable", "10.0")
	registerFunc(&_nSReturnAddress, &_nSReturnAddressErr, frameworkHandle, "NSReturnAddress", "10.0")
	registerFunc(&_nSRoundDownToMultipleOfPageSize, &_nSRoundDownToMultipleOfPageSizeErr, frameworkHandle, "NSRoundDownToMultipleOfPageSize", "10.0")
	registerFunc(&_nSRoundUpToMultipleOfPageSize, &_nSRoundUpToMultipleOfPageSizeErr, frameworkHandle, "NSRoundUpToMultipleOfPageSize", "10.0")
	registerFunc(&_nSSearchPathForDirectoriesInDomains, &_nSSearchPathForDirectoriesInDomainsErr, frameworkHandle, "NSSearchPathForDirectoriesInDomains", "10.0")
	registerFunc(&_nSSelectorFromString, &_nSSelectorFromStringErr, frameworkHandle, "NSSelectorFromString", "10.0")
	registerFunc(&_nSSetUncaughtExceptionHandler, &_nSSetUncaughtExceptionHandlerErr, frameworkHandle, "NSSetUncaughtExceptionHandler", "10.0")
	registerFunc(&_nSSetZoneName, &_nSSetZoneNameErr, frameworkHandle, "NSSetZoneName", "10.0")
	registerFunc(&_nSShouldRetainWithZone, &_nSShouldRetainWithZoneErr, frameworkHandle, "NSShouldRetainWithZone", "10.0")
	registerFunc(&_nSSizeFromString, &_nSSizeFromStringErr, frameworkHandle, "NSSizeFromString", "10.0")
	registerFunc(&_nSStringFromClass, &_nSStringFromClassErr, frameworkHandle, "NSStringFromClass", "10.0")
	registerFunc(&_nSStringFromHashTable, &_nSStringFromHashTableErr, frameworkHandle, "NSStringFromHashTable", "10.0")
	registerFunc(&_nSStringFromMapTable, &_nSStringFromMapTableErr, frameworkHandle, "NSStringFromMapTable", "10.0")
	registerFunc(&_nSStringFromPoint, &_nSStringFromPointErr, frameworkHandle, "NSStringFromPoint", "10.0")
	registerFunc(&_nSStringFromProtocol, &_nSStringFromProtocolErr, frameworkHandle, "NSStringFromProtocol", "10.5")
	registerFunc(&_nSStringFromRange, &_nSStringFromRangeErr, frameworkHandle, "NSStringFromRange", "10.0")
	registerFunc(&_nSStringFromRect, &_nSStringFromRectErr, frameworkHandle, "NSStringFromRect", "10.0")
	registerFunc(&_nSStringFromSelector, &_nSStringFromSelectorErr, frameworkHandle, "NSStringFromSelector", "10.0")
	registerFunc(&_nSStringFromSize, &_nSStringFromSizeErr, frameworkHandle, "NSStringFromSize", "10.0")
	registerFunc(&_nSTemporaryDirectory, &_nSTemporaryDirectoryErr, frameworkHandle, "NSTemporaryDirectory", "10.0")
	registerFunc(&_nSUserName, &_nSUserNameErr, frameworkHandle, "NSUserName", "10.0")
	registerFunc(&_nSZoneCalloc, &_nSZoneCallocErr, frameworkHandle, "NSZoneCalloc", "10.0")
	registerFunc(&_nSZoneFree, &_nSZoneFreeErr, frameworkHandle, "NSZoneFree", "10.0")
	registerFunc(&_nSZoneFromPointer, &_nSZoneFromPointerErr, frameworkHandle, "NSZoneFromPointer", "10.0")
	registerFunc(&_nSZoneMalloc, &_nSZoneMallocErr, frameworkHandle, "NSZoneMalloc", "10.0")
	registerFunc(&_nSZoneName, &_nSZoneNameErr, frameworkHandle, "NSZoneName", "10.0")
	registerFunc(&_nSZoneRealloc, &_nSZoneReallocErr, frameworkHandle, "NSZoneRealloc", "10.0")
	registerFunc(&_nXReadNSObjectFromCoder, &_nXReadNSObjectFromCoderErr, frameworkHandle, "NXReadNSObjectFromCoder", "10.0")
}
