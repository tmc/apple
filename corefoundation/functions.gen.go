// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/kernel"
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
		return fmt.Sprintf("CoreFoundation: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreFoundation: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreFoundation: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreFoundation: register symbol %s: %v", name, r)
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

var _cFAbsoluteTimeGetCurrent func() CFAbsoluteTime
var _cFAbsoluteTimeGetCurrentErr error

func tryCFAbsoluteTimeGetCurrent() (CFAbsoluteTime, error) {
	if _cFAbsoluteTimeGetCurrent == nil {
		return *new(CFAbsoluteTime), symbolCallError("CFAbsoluteTimeGetCurrent", "", _cFAbsoluteTimeGetCurrentErr)
	}
	return _cFAbsoluteTimeGetCurrent(), nil
}

// CFAbsoluteTimeGetCurrent returns the current system absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetCurrent()
func CFAbsoluteTimeGetCurrent() CFAbsoluteTime {
	result, callErr := tryCFAbsoluteTimeGetCurrent()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorAllocate func(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer
var _cFAllocatorAllocateErr error

func tryCFAllocatorAllocate(allocator CFAllocatorRef, size int, hint uint64) (unsafe.Pointer, error) {
	if _cFAllocatorAllocate == nil {
		return nil, symbolCallError("CFAllocatorAllocate", "", _cFAllocatorAllocateErr)
	}
	return _cFAllocatorAllocate(allocator, size, hint), nil
}

// CFAllocatorAllocate allocates memory using the specified allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocate(_:_:_:)
func CFAllocatorAllocate(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer {
	result, callErr := tryCFAllocatorAllocate(allocator, size, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorAllocateBytes func(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer
var _cFAllocatorAllocateBytesErr error

func tryCFAllocatorAllocateBytes(allocator CFAllocatorRef, size int, hint uint64) (unsafe.Pointer, error) {
	if _cFAllocatorAllocateBytes == nil {
		return nil, symbolCallError("CFAllocatorAllocateBytes", "15.0", _cFAllocatorAllocateBytesErr)
	}
	return _cFAllocatorAllocateBytes(allocator, size, hint), nil
}

// CFAllocatorAllocateBytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateBytes(_:_:_:)
func CFAllocatorAllocateBytes(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer {
	result, callErr := tryCFAllocatorAllocateBytes(allocator, size, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorAllocateTyped func(allocator CFAllocatorRef, size int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer
var _cFAllocatorAllocateTypedErr error

func tryCFAllocatorAllocateTyped(allocator CFAllocatorRef, size int, descriptor CFAllocatorTypeID, hint uint64) (unsafe.Pointer, error) {
	if _cFAllocatorAllocateTyped == nil {
		return nil, symbolCallError("CFAllocatorAllocateTyped", "15.0", _cFAllocatorAllocateTypedErr)
	}
	return _cFAllocatorAllocateTyped(allocator, size, descriptor, hint), nil
}

// CFAllocatorAllocateTyped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateTyped(_:_:_:_:)
func CFAllocatorAllocateTyped(allocator CFAllocatorRef, size int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer {
	result, callErr := tryCFAllocatorAllocateTyped(allocator, size, descriptor, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorCreate func(allocator CFAllocatorRef, context *CFAllocatorContext) CFAllocatorRef
var _cFAllocatorCreateErr error

func tryCFAllocatorCreate(allocator CFAllocatorRef, context *CFAllocatorContext) (CFAllocatorRef, error) {
	if _cFAllocatorCreate == nil {
		return 0, symbolCallError("CFAllocatorCreate", "", _cFAllocatorCreateErr)
	}
	return _cFAllocatorCreate(allocator, context), nil
}

// CFAllocatorCreate creates an allocator object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreate(_:_:)
func CFAllocatorCreate(allocator CFAllocatorRef, context *CFAllocatorContext) CFAllocatorRef {
	result, callErr := tryCFAllocatorCreate(allocator, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorCreateWithZone func(allocator CFAllocatorRef, zone unsafe.Pointer) CFAllocatorRef
var _cFAllocatorCreateWithZoneErr error

func tryCFAllocatorCreateWithZone(allocator CFAllocatorRef, zone unsafe.Pointer) (CFAllocatorRef, error) {
	if _cFAllocatorCreateWithZone == nil {
		return 0, symbolCallError("CFAllocatorCreateWithZone", "", _cFAllocatorCreateWithZoneErr)
	}
	return _cFAllocatorCreateWithZone(allocator, zone), nil
}

// CFAllocatorCreateWithZone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreateWithZone
func CFAllocatorCreateWithZone(allocator CFAllocatorRef, zone unsafe.Pointer) CFAllocatorRef {
	result, callErr := tryCFAllocatorCreateWithZone(allocator, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorDeallocate func(allocator CFAllocatorRef, ptr unsafe.Pointer)
var _cFAllocatorDeallocateErr error

func tryCFAllocatorDeallocate(allocator CFAllocatorRef, ptr unsafe.Pointer) error {
	if _cFAllocatorDeallocate == nil {
		return symbolCallError("CFAllocatorDeallocate", "", _cFAllocatorDeallocateErr)
	}
	_cFAllocatorDeallocate(allocator, ptr)
	return nil
}

// CFAllocatorDeallocate deallocates a block of memory with a given allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocate(_:_:)
func CFAllocatorDeallocate(allocator CFAllocatorRef, ptr unsafe.Pointer) {
	if callErr := tryCFAllocatorDeallocate(allocator, ptr); callErr != nil {
		panic(callErr)
	}
}

var _cFAllocatorGetContext func(allocator CFAllocatorRef, context *CFAllocatorContext)
var _cFAllocatorGetContextErr error

func tryCFAllocatorGetContext(allocator CFAllocatorRef, context *CFAllocatorContext) error {
	if _cFAllocatorGetContext == nil {
		return symbolCallError("CFAllocatorGetContext", "", _cFAllocatorGetContextErr)
	}
	_cFAllocatorGetContext(allocator, context)
	return nil
}

// CFAllocatorGetContext obtains the context of the specified allocator or of the default allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetContext(_:_:)
func CFAllocatorGetContext(allocator CFAllocatorRef, context *CFAllocatorContext) {
	if callErr := tryCFAllocatorGetContext(allocator, context); callErr != nil {
		panic(callErr)
	}
}

var _cFAllocatorGetDefault func() CFAllocatorRef
var _cFAllocatorGetDefaultErr error

func tryCFAllocatorGetDefault() (CFAllocatorRef, error) {
	if _cFAllocatorGetDefault == nil {
		return 0, symbolCallError("CFAllocatorGetDefault", "", _cFAllocatorGetDefaultErr)
	}
	return _cFAllocatorGetDefault(), nil
}

// CFAllocatorGetDefault gets the default allocator object for the current thread.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetDefault()
func CFAllocatorGetDefault() CFAllocatorRef {
	result, callErr := tryCFAllocatorGetDefault()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorGetPreferredSizeForSize func(allocator CFAllocatorRef, size int, hint uint64) int
var _cFAllocatorGetPreferredSizeForSizeErr error

func tryCFAllocatorGetPreferredSizeForSize(allocator CFAllocatorRef, size int, hint uint64) (int, error) {
	if _cFAllocatorGetPreferredSizeForSize == nil {
		return 0, symbolCallError("CFAllocatorGetPreferredSizeForSize", "", _cFAllocatorGetPreferredSizeForSizeErr)
	}
	return _cFAllocatorGetPreferredSizeForSize(allocator, size, hint), nil
}

// CFAllocatorGetPreferredSizeForSize obtains the number of bytes likely to be allocated upon a specific request.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetPreferredSizeForSize(_:_:_:)
func CFAllocatorGetPreferredSizeForSize(allocator CFAllocatorRef, size int, hint uint64) int {
	result, callErr := tryCFAllocatorGetPreferredSizeForSize(allocator, size, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorGetTypeID func() uint
var _cFAllocatorGetTypeIDErr error

func tryCFAllocatorGetTypeID() (uint, error) {
	if _cFAllocatorGetTypeID == nil {
		return 0, symbolCallError("CFAllocatorGetTypeID", "", _cFAllocatorGetTypeIDErr)
	}
	return _cFAllocatorGetTypeID(), nil
}

// CFAllocatorGetTypeID returns the type identifier for the CFAllocator opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetTypeID()
func CFAllocatorGetTypeID() uint {
	result, callErr := tryCFAllocatorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorReallocate func(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer
var _cFAllocatorReallocateErr error

func tryCFAllocatorReallocate(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) (unsafe.Pointer, error) {
	if _cFAllocatorReallocate == nil {
		return nil, symbolCallError("CFAllocatorReallocate", "", _cFAllocatorReallocateErr)
	}
	return _cFAllocatorReallocate(allocator, ptr, newsize, hint), nil
}

// CFAllocatorReallocate reallocates memory using the specified allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocate(_:_:_:_:)
func CFAllocatorReallocate(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer {
	result, callErr := tryCFAllocatorReallocate(allocator, ptr, newsize, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorReallocateBytes func(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer
var _cFAllocatorReallocateBytesErr error

func tryCFAllocatorReallocateBytes(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) (unsafe.Pointer, error) {
	if _cFAllocatorReallocateBytes == nil {
		return nil, symbolCallError("CFAllocatorReallocateBytes", "15.0", _cFAllocatorReallocateBytesErr)
	}
	return _cFAllocatorReallocateBytes(allocator, ptr, newsize, hint), nil
}

// CFAllocatorReallocateBytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateBytes(_:_:_:_:)
func CFAllocatorReallocateBytes(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer {
	result, callErr := tryCFAllocatorReallocateBytes(allocator, ptr, newsize, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorReallocateTyped func(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer
var _cFAllocatorReallocateTypedErr error

func tryCFAllocatorReallocateTyped(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, descriptor CFAllocatorTypeID, hint uint64) (unsafe.Pointer, error) {
	if _cFAllocatorReallocateTyped == nil {
		return nil, symbolCallError("CFAllocatorReallocateTyped", "15.0", _cFAllocatorReallocateTypedErr)
	}
	return _cFAllocatorReallocateTyped(allocator, ptr, newsize, descriptor, hint), nil
}

// CFAllocatorReallocateTyped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateTyped(_:_:_:_:_:)
func CFAllocatorReallocateTyped(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer {
	result, callErr := tryCFAllocatorReallocateTyped(allocator, ptr, newsize, descriptor, hint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAllocatorSetDefault func(allocator CFAllocatorRef)
var _cFAllocatorSetDefaultErr error

func tryCFAllocatorSetDefault(allocator CFAllocatorRef) error {
	if _cFAllocatorSetDefault == nil {
		return symbolCallError("CFAllocatorSetDefault", "", _cFAllocatorSetDefaultErr)
	}
	_cFAllocatorSetDefault(allocator)
	return nil
}

// CFAllocatorSetDefault sets the given allocator as the default for the current thread.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorSetDefault(_:)
func CFAllocatorSetDefault(allocator CFAllocatorRef) {
	if callErr := tryCFAllocatorSetDefault(allocator); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayAppendArray func(theArray CFMutableArrayRef, otherArray CFArrayRef, otherRange CFRange)
var _cFArrayAppendArrayErr error

func tryCFArrayAppendArray(theArray CFMutableArrayRef, otherArray CFArrayRef, otherRange CFRange) error {
	if _cFArrayAppendArray == nil {
		return symbolCallError("CFArrayAppendArray", "", _cFArrayAppendArrayErr)
	}
	_cFArrayAppendArray(theArray, otherArray, otherRange)
	return nil
}

// CFArrayAppendArray adds the values from one array to another array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendArray(_:_:_:)
func CFArrayAppendArray(theArray CFMutableArrayRef, otherArray CFArrayRef, otherRange CFRange) {
	if callErr := tryCFArrayAppendArray(theArray, otherArray, otherRange); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayAppendValue func(theArray CFMutableArrayRef, value unsafe.Pointer)
var _cFArrayAppendValueErr error

func tryCFArrayAppendValue(theArray CFMutableArrayRef, value unsafe.Pointer) error {
	if _cFArrayAppendValue == nil {
		return symbolCallError("CFArrayAppendValue", "", _cFArrayAppendValueErr)
	}
	_cFArrayAppendValue(theArray, value)
	return nil
}

// CFArrayAppendValue adds a value to an array giving it the new largest index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendValue(_:_:)
func CFArrayAppendValue(theArray CFMutableArrayRef, value unsafe.Pointer) {
	if callErr := tryCFArrayAppendValue(theArray, value); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayApplyFunction func(theArray CFArrayRef, range_ CFRange, applier CFArrayApplierFunction, context unsafe.Pointer)
var _cFArrayApplyFunctionErr error

func tryCFArrayApplyFunction(theArray CFArrayRef, range_ CFRange, applier CFArrayApplierFunction, context unsafe.Pointer) error {
	if _cFArrayApplyFunction == nil {
		return symbolCallError("CFArrayApplyFunction", "", _cFArrayApplyFunctionErr)
	}
	_cFArrayApplyFunction(theArray, range_, applier, context)
	return nil
}

// CFArrayApplyFunction calls a function once for each element in range in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplyFunction(_:_:_:_:)
func CFArrayApplyFunction(theArray CFArrayRef, range_ CFRange, applier CFArrayApplierFunction, context unsafe.Pointer) {
	if callErr := tryCFArrayApplyFunction(theArray, range_, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayBSearchValues func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer, comparator CFComparatorFunction, context unsafe.Pointer) int
var _cFArrayBSearchValuesErr error

func tryCFArrayBSearchValues(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer, comparator CFComparatorFunction, context unsafe.Pointer) (int, error) {
	if _cFArrayBSearchValues == nil {
		return 0, symbolCallError("CFArrayBSearchValues", "", _cFArrayBSearchValuesErr)
	}
	return _cFArrayBSearchValues(theArray, range_, value, comparator, context), nil
}

// CFArrayBSearchValues searches an array for a value using a binary search algorithm.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayBSearchValues(_:_:_:_:_:)
func CFArrayBSearchValues(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer, comparator CFComparatorFunction, context unsafe.Pointer) int {
	result, callErr := tryCFArrayBSearchValues(theArray, range_, value, comparator, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayContainsValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) bool
var _cFArrayContainsValueErr error

func tryCFArrayContainsValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) (bool, error) {
	if _cFArrayContainsValue == nil {
		return false, symbolCallError("CFArrayContainsValue", "", _cFArrayContainsValueErr)
	}
	return _cFArrayContainsValue(theArray, range_, value), nil
}

// CFArrayContainsValue reports whether or not a value is in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayContainsValue(_:_:_:)
func CFArrayContainsValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) bool {
	result, callErr := tryCFArrayContainsValue(theArray, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayCreate func(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFArrayCallBacks) CFArrayRef
var _cFArrayCreateErr error

func tryCFArrayCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFArrayCallBacks) (CFArrayRef, error) {
	if _cFArrayCreate == nil {
		return 0, symbolCallError("CFArrayCreate", "", _cFArrayCreateErr)
	}
	return _cFArrayCreate(allocator, values, numValues, callBacks), nil
}

// CFArrayCreate creates a new immutable array with the given values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreate(_:_:_:_:)
func CFArrayCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFArrayCallBacks) CFArrayRef {
	result, callErr := tryCFArrayCreate(allocator, values, numValues, callBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayCreateCopy func(allocator CFAllocatorRef, theArray CFArrayRef) CFArrayRef
var _cFArrayCreateCopyErr error

func tryCFArrayCreateCopy(allocator CFAllocatorRef, theArray CFArrayRef) (CFArrayRef, error) {
	if _cFArrayCreateCopy == nil {
		return 0, symbolCallError("CFArrayCreateCopy", "", _cFArrayCreateCopyErr)
	}
	return _cFArrayCreateCopy(allocator, theArray), nil
}

// CFArrayCreateCopy creates a new immutable array with the values from another array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateCopy(_:_:)
func CFArrayCreateCopy(allocator CFAllocatorRef, theArray CFArrayRef) CFArrayRef {
	result, callErr := tryCFArrayCreateCopy(allocator, theArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayCreateMutable func(allocator CFAllocatorRef, capacity int, callBacks *CFArrayCallBacks) CFMutableArrayRef
var _cFArrayCreateMutableErr error

func tryCFArrayCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFArrayCallBacks) (CFMutableArrayRef, error) {
	if _cFArrayCreateMutable == nil {
		return 0, symbolCallError("CFArrayCreateMutable", "", _cFArrayCreateMutableErr)
	}
	return _cFArrayCreateMutable(allocator, capacity, callBacks), nil
}

// CFArrayCreateMutable creates a new empty mutable array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutable(_:_:_:)
func CFArrayCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFArrayCallBacks) CFMutableArrayRef {
	result, callErr := tryCFArrayCreateMutable(allocator, capacity, callBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theArray CFArrayRef) CFMutableArrayRef
var _cFArrayCreateMutableCopyErr error

func tryCFArrayCreateMutableCopy(allocator CFAllocatorRef, capacity int, theArray CFArrayRef) (CFMutableArrayRef, error) {
	if _cFArrayCreateMutableCopy == nil {
		return 0, symbolCallError("CFArrayCreateMutableCopy", "", _cFArrayCreateMutableCopyErr)
	}
	return _cFArrayCreateMutableCopy(allocator, capacity, theArray), nil
}

// CFArrayCreateMutableCopy creates a new mutable array with the values from another array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutableCopy(_:_:_:)
func CFArrayCreateMutableCopy(allocator CFAllocatorRef, capacity int, theArray CFArrayRef) CFMutableArrayRef {
	result, callErr := tryCFArrayCreateMutableCopy(allocator, capacity, theArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayExchangeValuesAtIndices func(theArray CFMutableArrayRef, idx1 int, idx2 int)
var _cFArrayExchangeValuesAtIndicesErr error

func tryCFArrayExchangeValuesAtIndices(theArray CFMutableArrayRef, idx1 int, idx2 int) error {
	if _cFArrayExchangeValuesAtIndices == nil {
		return symbolCallError("CFArrayExchangeValuesAtIndices", "", _cFArrayExchangeValuesAtIndicesErr)
	}
	_cFArrayExchangeValuesAtIndices(theArray, idx1, idx2)
	return nil
}

// CFArrayExchangeValuesAtIndices exchanges the values at two indices of an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayExchangeValuesAtIndices(_:_:_:)
func CFArrayExchangeValuesAtIndices(theArray CFMutableArrayRef, idx1 int, idx2 int) {
	if callErr := tryCFArrayExchangeValuesAtIndices(theArray, idx1, idx2); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayGetCount func(theArray CFArrayRef) int
var _cFArrayGetCountErr error

func tryCFArrayGetCount(theArray CFArrayRef) (int, error) {
	if _cFArrayGetCount == nil {
		return 0, symbolCallError("CFArrayGetCount", "", _cFArrayGetCountErr)
	}
	return _cFArrayGetCount(theArray), nil
}

// CFArrayGetCount returns the number of values currently in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCount(_:)
func CFArrayGetCount(theArray CFArrayRef) int {
	result, callErr := tryCFArrayGetCount(theArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayGetCountOfValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int
var _cFArrayGetCountOfValueErr error

func tryCFArrayGetCountOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) (int, error) {
	if _cFArrayGetCountOfValue == nil {
		return 0, symbolCallError("CFArrayGetCountOfValue", "", _cFArrayGetCountOfValueErr)
	}
	return _cFArrayGetCountOfValue(theArray, range_, value), nil
}

// CFArrayGetCountOfValue counts the number of times a given value occurs in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCountOfValue(_:_:_:)
func CFArrayGetCountOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int {
	result, callErr := tryCFArrayGetCountOfValue(theArray, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayGetFirstIndexOfValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int
var _cFArrayGetFirstIndexOfValueErr error

func tryCFArrayGetFirstIndexOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) (int, error) {
	if _cFArrayGetFirstIndexOfValue == nil {
		return 0, symbolCallError("CFArrayGetFirstIndexOfValue", "", _cFArrayGetFirstIndexOfValueErr)
	}
	return _cFArrayGetFirstIndexOfValue(theArray, range_, value), nil
}

// CFArrayGetFirstIndexOfValue searches an array forward for a value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetFirstIndexOfValue(_:_:_:)
func CFArrayGetFirstIndexOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int {
	result, callErr := tryCFArrayGetFirstIndexOfValue(theArray, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayGetLastIndexOfValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int
var _cFArrayGetLastIndexOfValueErr error

func tryCFArrayGetLastIndexOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) (int, error) {
	if _cFArrayGetLastIndexOfValue == nil {
		return 0, symbolCallError("CFArrayGetLastIndexOfValue", "", _cFArrayGetLastIndexOfValueErr)
	}
	return _cFArrayGetLastIndexOfValue(theArray, range_, value), nil
}

// CFArrayGetLastIndexOfValue searches an array backward for a value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetLastIndexOfValue(_:_:_:)
func CFArrayGetLastIndexOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int {
	result, callErr := tryCFArrayGetLastIndexOfValue(theArray, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayGetTypeID func() uint
var _cFArrayGetTypeIDErr error

func tryCFArrayGetTypeID() (uint, error) {
	if _cFArrayGetTypeID == nil {
		return 0, symbolCallError("CFArrayGetTypeID", "", _cFArrayGetTypeIDErr)
	}
	return _cFArrayGetTypeID(), nil
}

// CFArrayGetTypeID returns the type identifier for the CFArray opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetTypeID()
func CFArrayGetTypeID() uint {
	result, callErr := tryCFArrayGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayGetValueAtIndex func(theArray CFArrayRef, idx int) unsafe.Pointer
var _cFArrayGetValueAtIndexErr error

func tryCFArrayGetValueAtIndex(theArray CFArrayRef, idx int) (unsafe.Pointer, error) {
	if _cFArrayGetValueAtIndex == nil {
		return nil, symbolCallError("CFArrayGetValueAtIndex", "", _cFArrayGetValueAtIndexErr)
	}
	return _cFArrayGetValueAtIndex(theArray, idx), nil
}

// CFArrayGetValueAtIndex retrieves a value at a given index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValueAtIndex(_:_:)
func CFArrayGetValueAtIndex(theArray CFArrayRef, idx int) unsafe.Pointer {
	result, callErr := tryCFArrayGetValueAtIndex(theArray, idx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFArrayGetValues func(theArray CFArrayRef, range_ CFRange, values unsafe.Pointer)
var _cFArrayGetValuesErr error

func tryCFArrayGetValues(theArray CFArrayRef, range_ CFRange, values unsafe.Pointer) error {
	if _cFArrayGetValues == nil {
		return symbolCallError("CFArrayGetValues", "", _cFArrayGetValuesErr)
	}
	_cFArrayGetValues(theArray, range_, values)
	return nil
}

// CFArrayGetValues fills a buffer with values from an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValues(_:_:_:)
func CFArrayGetValues(theArray CFArrayRef, range_ CFRange, values unsafe.Pointer) {
	if callErr := tryCFArrayGetValues(theArray, range_, values); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayInsertValueAtIndex func(theArray CFMutableArrayRef, idx int, value unsafe.Pointer)
var _cFArrayInsertValueAtIndexErr error

func tryCFArrayInsertValueAtIndex(theArray CFMutableArrayRef, idx int, value unsafe.Pointer) error {
	if _cFArrayInsertValueAtIndex == nil {
		return symbolCallError("CFArrayInsertValueAtIndex", "", _cFArrayInsertValueAtIndexErr)
	}
	_cFArrayInsertValueAtIndex(theArray, idx, value)
	return nil
}

// CFArrayInsertValueAtIndex inserts a value into an array at a given index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayInsertValueAtIndex(_:_:_:)
func CFArrayInsertValueAtIndex(theArray CFMutableArrayRef, idx int, value unsafe.Pointer) {
	if callErr := tryCFArrayInsertValueAtIndex(theArray, idx, value); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayRemoveAllValues func(theArray CFMutableArrayRef)
var _cFArrayRemoveAllValuesErr error

func tryCFArrayRemoveAllValues(theArray CFMutableArrayRef) error {
	if _cFArrayRemoveAllValues == nil {
		return symbolCallError("CFArrayRemoveAllValues", "", _cFArrayRemoveAllValuesErr)
	}
	_cFArrayRemoveAllValues(theArray)
	return nil
}

// CFArrayRemoveAllValues removes all the values from an array, making it empty.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveAllValues(_:)
func CFArrayRemoveAllValues(theArray CFMutableArrayRef) {
	if callErr := tryCFArrayRemoveAllValues(theArray); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayRemoveValueAtIndex func(theArray CFMutableArrayRef, idx int)
var _cFArrayRemoveValueAtIndexErr error

func tryCFArrayRemoveValueAtIndex(theArray CFMutableArrayRef, idx int) error {
	if _cFArrayRemoveValueAtIndex == nil {
		return symbolCallError("CFArrayRemoveValueAtIndex", "", _cFArrayRemoveValueAtIndexErr)
	}
	_cFArrayRemoveValueAtIndex(theArray, idx)
	return nil
}

// CFArrayRemoveValueAtIndex removes the value at a given index from an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveValueAtIndex(_:_:)
func CFArrayRemoveValueAtIndex(theArray CFMutableArrayRef, idx int) {
	if callErr := tryCFArrayRemoveValueAtIndex(theArray, idx); callErr != nil {
		panic(callErr)
	}
}

var _cFArrayReplaceValues func(theArray CFMutableArrayRef, range_ CFRange, newValues unsafe.Pointer, newCount int)
var _cFArrayReplaceValuesErr error

func tryCFArrayReplaceValues(theArray CFMutableArrayRef, range_ CFRange, newValues unsafe.Pointer, newCount int) error {
	if _cFArrayReplaceValues == nil {
		return symbolCallError("CFArrayReplaceValues", "", _cFArrayReplaceValuesErr)
	}
	_cFArrayReplaceValues(theArray, range_, newValues, newCount)
	return nil
}

// CFArrayReplaceValues replaces a range of values in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayReplaceValues(_:_:_:_:)
func CFArrayReplaceValues(theArray CFMutableArrayRef, range_ CFRange, newValues unsafe.Pointer, newCount int) {
	if callErr := tryCFArrayReplaceValues(theArray, range_, newValues, newCount); callErr != nil {
		panic(callErr)
	}
}

var _cFArraySetValueAtIndex func(theArray CFMutableArrayRef, idx int, value unsafe.Pointer)
var _cFArraySetValueAtIndexErr error

func tryCFArraySetValueAtIndex(theArray CFMutableArrayRef, idx int, value unsafe.Pointer) error {
	if _cFArraySetValueAtIndex == nil {
		return symbolCallError("CFArraySetValueAtIndex", "", _cFArraySetValueAtIndexErr)
	}
	_cFArraySetValueAtIndex(theArray, idx, value)
	return nil
}

// CFArraySetValueAtIndex changes the value at a given index in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArraySetValueAtIndex(_:_:_:)
func CFArraySetValueAtIndex(theArray CFMutableArrayRef, idx int, value unsafe.Pointer) {
	if callErr := tryCFArraySetValueAtIndex(theArray, idx, value); callErr != nil {
		panic(callErr)
	}
}

var _cFArraySortValues func(theArray CFMutableArrayRef, range_ CFRange, comparator CFComparatorFunction, context unsafe.Pointer)
var _cFArraySortValuesErr error

func tryCFArraySortValues(theArray CFMutableArrayRef, range_ CFRange, comparator CFComparatorFunction, context unsafe.Pointer) error {
	if _cFArraySortValues == nil {
		return symbolCallError("CFArraySortValues", "", _cFArraySortValuesErr)
	}
	_cFArraySortValues(theArray, range_, comparator, context)
	return nil
}

// CFArraySortValues sorts the values in an array using a given comparison function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArraySortValues(_:_:_:_:)
func CFArraySortValues(theArray CFMutableArrayRef, range_ CFRange, comparator CFComparatorFunction, context unsafe.Pointer) {
	if callErr := tryCFArraySortValues(theArray, range_, comparator, context); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringBeginEditing func(aStr CFMutableAttributedStringRef)
var _cFAttributedStringBeginEditingErr error

func tryCFAttributedStringBeginEditing(aStr CFMutableAttributedStringRef) error {
	if _cFAttributedStringBeginEditing == nil {
		return symbolCallError("CFAttributedStringBeginEditing", "", _cFAttributedStringBeginEditingErr)
	}
	_cFAttributedStringBeginEditing(aStr)
	return nil
}

// CFAttributedStringBeginEditing defers internal consistency-checking and coalescing for a mutable attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringBeginEditing(_:)
func CFAttributedStringBeginEditing(aStr CFMutableAttributedStringRef) {
	if callErr := tryCFAttributedStringBeginEditing(aStr); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringCreate func(alloc CFAllocatorRef, str CFStringRef, attributes CFDictionaryRef) CFAttributedStringRef
var _cFAttributedStringCreateErr error

func tryCFAttributedStringCreate(alloc CFAllocatorRef, str CFStringRef, attributes CFDictionaryRef) (CFAttributedStringRef, error) {
	if _cFAttributedStringCreate == nil {
		return 0, symbolCallError("CFAttributedStringCreate", "", _cFAttributedStringCreateErr)
	}
	return _cFAttributedStringCreate(alloc, str, attributes), nil
}

// CFAttributedStringCreate creates an attributed string with specified string and attributes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreate(_:_:_:)
func CFAttributedStringCreate(alloc CFAllocatorRef, str CFStringRef, attributes CFDictionaryRef) CFAttributedStringRef {
	result, callErr := tryCFAttributedStringCreate(alloc, str, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringCreateCopy func(alloc CFAllocatorRef, aStr CFAttributedStringRef) CFAttributedStringRef
var _cFAttributedStringCreateCopyErr error

func tryCFAttributedStringCreateCopy(alloc CFAllocatorRef, aStr CFAttributedStringRef) (CFAttributedStringRef, error) {
	if _cFAttributedStringCreateCopy == nil {
		return 0, symbolCallError("CFAttributedStringCreateCopy", "", _cFAttributedStringCreateCopyErr)
	}
	return _cFAttributedStringCreateCopy(alloc, aStr), nil
}

// CFAttributedStringCreateCopy creates an immutable copy of an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateCopy(_:_:)
func CFAttributedStringCreateCopy(alloc CFAllocatorRef, aStr CFAttributedStringRef) CFAttributedStringRef {
	result, callErr := tryCFAttributedStringCreateCopy(alloc, aStr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringCreateMutable func(alloc CFAllocatorRef, maxLength int) CFMutableAttributedStringRef
var _cFAttributedStringCreateMutableErr error

func tryCFAttributedStringCreateMutable(alloc CFAllocatorRef, maxLength int) (CFMutableAttributedStringRef, error) {
	if _cFAttributedStringCreateMutable == nil {
		return 0, symbolCallError("CFAttributedStringCreateMutable", "", _cFAttributedStringCreateMutableErr)
	}
	return _cFAttributedStringCreateMutable(alloc, maxLength), nil
}

// CFAttributedStringCreateMutable creates a mutable attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutable(_:_:)
func CFAttributedStringCreateMutable(alloc CFAllocatorRef, maxLength int) CFMutableAttributedStringRef {
	result, callErr := tryCFAttributedStringCreateMutable(alloc, maxLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringCreateMutableCopy func(alloc CFAllocatorRef, maxLength int, aStr CFAttributedStringRef) CFMutableAttributedStringRef
var _cFAttributedStringCreateMutableCopyErr error

func tryCFAttributedStringCreateMutableCopy(alloc CFAllocatorRef, maxLength int, aStr CFAttributedStringRef) (CFMutableAttributedStringRef, error) {
	if _cFAttributedStringCreateMutableCopy == nil {
		return 0, symbolCallError("CFAttributedStringCreateMutableCopy", "", _cFAttributedStringCreateMutableCopyErr)
	}
	return _cFAttributedStringCreateMutableCopy(alloc, maxLength, aStr), nil
}

// CFAttributedStringCreateMutableCopy creates a mutable copy of an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutableCopy(_:_:_:)
func CFAttributedStringCreateMutableCopy(alloc CFAllocatorRef, maxLength int, aStr CFAttributedStringRef) CFMutableAttributedStringRef {
	result, callErr := tryCFAttributedStringCreateMutableCopy(alloc, maxLength, aStr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringCreateWithSubstring func(alloc CFAllocatorRef, aStr CFAttributedStringRef, range_ CFRange) CFAttributedStringRef
var _cFAttributedStringCreateWithSubstringErr error

func tryCFAttributedStringCreateWithSubstring(alloc CFAllocatorRef, aStr CFAttributedStringRef, range_ CFRange) (CFAttributedStringRef, error) {
	if _cFAttributedStringCreateWithSubstring == nil {
		return 0, symbolCallError("CFAttributedStringCreateWithSubstring", "", _cFAttributedStringCreateWithSubstringErr)
	}
	return _cFAttributedStringCreateWithSubstring(alloc, aStr, range_), nil
}

// CFAttributedStringCreateWithSubstring creates a sub-attributed string from the specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateWithSubstring(_:_:_:)
func CFAttributedStringCreateWithSubstring(alloc CFAllocatorRef, aStr CFAttributedStringRef, range_ CFRange) CFAttributedStringRef {
	result, callErr := tryCFAttributedStringCreateWithSubstring(alloc, aStr, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringEndEditing func(aStr CFMutableAttributedStringRef)
var _cFAttributedStringEndEditingErr error

func tryCFAttributedStringEndEditing(aStr CFMutableAttributedStringRef) error {
	if _cFAttributedStringEndEditing == nil {
		return symbolCallError("CFAttributedStringEndEditing", "", _cFAttributedStringEndEditingErr)
	}
	_cFAttributedStringEndEditing(aStr)
	return nil
}

// CFAttributedStringEndEditing re-enables internal consistency-checking and coalescing for a mutable attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringEndEditing(_:)
func CFAttributedStringEndEditing(aStr CFMutableAttributedStringRef) {
	if callErr := tryCFAttributedStringEndEditing(aStr); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringGetAttribute func(aStr CFAttributedStringRef, loc int, attrName CFStringRef, effectiveRange *CFRange) CFTypeRef
var _cFAttributedStringGetAttributeErr error

func tryCFAttributedStringGetAttribute(aStr CFAttributedStringRef, loc int, attrName CFStringRef, effectiveRange *CFRange) (CFTypeRef, error) {
	if _cFAttributedStringGetAttribute == nil {
		return 0, symbolCallError("CFAttributedStringGetAttribute", "", _cFAttributedStringGetAttributeErr)
	}
	return _cFAttributedStringGetAttribute(aStr, loc, attrName, effectiveRange), nil
}

// CFAttributedStringGetAttribute returns the value of a given attribute of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttribute(_:_:_:_:)
func CFAttributedStringGetAttribute(aStr CFAttributedStringRef, loc int, attrName CFStringRef, effectiveRange *CFRange) CFTypeRef {
	result, callErr := tryCFAttributedStringGetAttribute(aStr, loc, attrName, effectiveRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetAttributeAndLongestEffectiveRange func(aStr CFAttributedStringRef, loc int, attrName CFStringRef, inRange CFRange, longestEffectiveRange *CFRange) CFTypeRef
var _cFAttributedStringGetAttributeAndLongestEffectiveRangeErr error

func tryCFAttributedStringGetAttributeAndLongestEffectiveRange(aStr CFAttributedStringRef, loc int, attrName CFStringRef, inRange CFRange, longestEffectiveRange *CFRange) (CFTypeRef, error) {
	if _cFAttributedStringGetAttributeAndLongestEffectiveRange == nil {
		return 0, symbolCallError("CFAttributedStringGetAttributeAndLongestEffectiveRange", "", _cFAttributedStringGetAttributeAndLongestEffectiveRangeErr)
	}
	return _cFAttributedStringGetAttributeAndLongestEffectiveRange(aStr, loc, attrName, inRange, longestEffectiveRange), nil
}

// CFAttributedStringGetAttributeAndLongestEffectiveRange returns the value of a given attribute of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributeAndLongestEffectiveRange(_:_:_:_:_:)
func CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr CFAttributedStringRef, loc int, attrName CFStringRef, inRange CFRange, longestEffectiveRange *CFRange) CFTypeRef {
	result, callErr := tryCFAttributedStringGetAttributeAndLongestEffectiveRange(aStr, loc, attrName, inRange, longestEffectiveRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetAttributes func(aStr CFAttributedStringRef, loc int, effectiveRange *CFRange) CFDictionaryRef
var _cFAttributedStringGetAttributesErr error

func tryCFAttributedStringGetAttributes(aStr CFAttributedStringRef, loc int, effectiveRange *CFRange) (CFDictionaryRef, error) {
	if _cFAttributedStringGetAttributes == nil {
		return 0, symbolCallError("CFAttributedStringGetAttributes", "", _cFAttributedStringGetAttributesErr)
	}
	return _cFAttributedStringGetAttributes(aStr, loc, effectiveRange), nil
}

// CFAttributedStringGetAttributes returns the attributes of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributes(_:_:_:)
func CFAttributedStringGetAttributes(aStr CFAttributedStringRef, loc int, effectiveRange *CFRange) CFDictionaryRef {
	result, callErr := tryCFAttributedStringGetAttributes(aStr, loc, effectiveRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetAttributesAndLongestEffectiveRange func(aStr CFAttributedStringRef, loc int, inRange CFRange, longestEffectiveRange *CFRange) CFDictionaryRef
var _cFAttributedStringGetAttributesAndLongestEffectiveRangeErr error

func tryCFAttributedStringGetAttributesAndLongestEffectiveRange(aStr CFAttributedStringRef, loc int, inRange CFRange, longestEffectiveRange *CFRange) (CFDictionaryRef, error) {
	if _cFAttributedStringGetAttributesAndLongestEffectiveRange == nil {
		return 0, symbolCallError("CFAttributedStringGetAttributesAndLongestEffectiveRange", "", _cFAttributedStringGetAttributesAndLongestEffectiveRangeErr)
	}
	return _cFAttributedStringGetAttributesAndLongestEffectiveRange(aStr, loc, inRange, longestEffectiveRange), nil
}

// CFAttributedStringGetAttributesAndLongestEffectiveRange returns the attributes of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributesAndLongestEffectiveRange(_:_:_:_:)
func CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr CFAttributedStringRef, loc int, inRange CFRange, longestEffectiveRange *CFRange) CFDictionaryRef {
	result, callErr := tryCFAttributedStringGetAttributesAndLongestEffectiveRange(aStr, loc, inRange, longestEffectiveRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetBidiLevelsAndResolvedDirections func(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool
var _cFAttributedStringGetBidiLevelsAndResolvedDirectionsErr error

func tryCFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) (bool, error) {
	if _cFAttributedStringGetBidiLevelsAndResolvedDirections == nil {
		return false, symbolCallError("CFAttributedStringGetBidiLevelsAndResolvedDirections", "", _cFAttributedStringGetBidiLevelsAndResolvedDirectionsErr)
	}
	return _cFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections), nil
}

// CFAttributedStringGetBidiLevelsAndResolvedDirections.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetBidiLevelsAndResolvedDirections(_:_:_:_:_:)
func CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool {
	result, callErr := tryCFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetLength func(aStr CFAttributedStringRef) int
var _cFAttributedStringGetLengthErr error

func tryCFAttributedStringGetLength(aStr CFAttributedStringRef) (int, error) {
	if _cFAttributedStringGetLength == nil {
		return 0, symbolCallError("CFAttributedStringGetLength", "", _cFAttributedStringGetLengthErr)
	}
	return _cFAttributedStringGetLength(aStr), nil
}

// CFAttributedStringGetLength returns the length of the attributed string in characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetLength(_:)
func CFAttributedStringGetLength(aStr CFAttributedStringRef) int {
	result, callErr := tryCFAttributedStringGetLength(aStr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetMutableString func(aStr CFMutableAttributedStringRef) CFMutableStringRef
var _cFAttributedStringGetMutableStringErr error

func tryCFAttributedStringGetMutableString(aStr CFMutableAttributedStringRef) (CFMutableStringRef, error) {
	if _cFAttributedStringGetMutableString == nil {
		return 0, symbolCallError("CFAttributedStringGetMutableString", "", _cFAttributedStringGetMutableStringErr)
	}
	return _cFAttributedStringGetMutableString(aStr), nil
}

// CFAttributedStringGetMutableString gets as a mutable string the string for an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetMutableString(_:)
func CFAttributedStringGetMutableString(aStr CFMutableAttributedStringRef) CFMutableStringRef {
	result, callErr := tryCFAttributedStringGetMutableString(aStr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetStatisticalWritingDirections func(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool
var _cFAttributedStringGetStatisticalWritingDirectionsErr error

func tryCFAttributedStringGetStatisticalWritingDirections(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) (bool, error) {
	if _cFAttributedStringGetStatisticalWritingDirections == nil {
		return false, symbolCallError("CFAttributedStringGetStatisticalWritingDirections", "26.0", _cFAttributedStringGetStatisticalWritingDirectionsErr)
	}
	return _cFAttributedStringGetStatisticalWritingDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections), nil
}

// CFAttributedStringGetStatisticalWritingDirections.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetStatisticalWritingDirections(_:_:_:_:_:)
func CFAttributedStringGetStatisticalWritingDirections(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool {
	result, callErr := tryCFAttributedStringGetStatisticalWritingDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetString func(aStr CFAttributedStringRef) CFStringRef
var _cFAttributedStringGetStringErr error

func tryCFAttributedStringGetString(aStr CFAttributedStringRef) (CFStringRef, error) {
	if _cFAttributedStringGetString == nil {
		return 0, symbolCallError("CFAttributedStringGetString", "", _cFAttributedStringGetStringErr)
	}
	return _cFAttributedStringGetString(aStr), nil
}

// CFAttributedStringGetString returns the string for an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetString(_:)
func CFAttributedStringGetString(aStr CFAttributedStringRef) CFStringRef {
	result, callErr := tryCFAttributedStringGetString(aStr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringGetTypeID func() uint
var _cFAttributedStringGetTypeIDErr error

func tryCFAttributedStringGetTypeID() (uint, error) {
	if _cFAttributedStringGetTypeID == nil {
		return 0, symbolCallError("CFAttributedStringGetTypeID", "", _cFAttributedStringGetTypeIDErr)
	}
	return _cFAttributedStringGetTypeID(), nil
}

// CFAttributedStringGetTypeID returns the type identifier for the CFAttributedString opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetTypeID()
func CFAttributedStringGetTypeID() uint {
	result, callErr := tryCFAttributedStringGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFAttributedStringRemoveAttribute func(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef)
var _cFAttributedStringRemoveAttributeErr error

func tryCFAttributedStringRemoveAttribute(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef) error {
	if _cFAttributedStringRemoveAttribute == nil {
		return symbolCallError("CFAttributedStringRemoveAttribute", "", _cFAttributedStringRemoveAttributeErr)
	}
	_cFAttributedStringRemoveAttribute(aStr, range_, attrName)
	return nil
}

// CFAttributedStringRemoveAttribute removes the value of a single attribute over a specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringRemoveAttribute(_:_:_:)
func CFAttributedStringRemoveAttribute(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef) {
	if callErr := tryCFAttributedStringRemoveAttribute(aStr, range_, attrName); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringReplaceAttributedString func(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFAttributedStringRef)
var _cFAttributedStringReplaceAttributedStringErr error

func tryCFAttributedStringReplaceAttributedString(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFAttributedStringRef) error {
	if _cFAttributedStringReplaceAttributedString == nil {
		return symbolCallError("CFAttributedStringReplaceAttributedString", "", _cFAttributedStringReplaceAttributedStringErr)
	}
	_cFAttributedStringReplaceAttributedString(aStr, range_, replacement)
	return nil
}

// CFAttributedStringReplaceAttributedString replaces the attributed substring over a range with another attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceAttributedString(_:_:_:)
func CFAttributedStringReplaceAttributedString(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFAttributedStringRef) {
	if callErr := tryCFAttributedStringReplaceAttributedString(aStr, range_, replacement); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringReplaceString func(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFStringRef)
var _cFAttributedStringReplaceStringErr error

func tryCFAttributedStringReplaceString(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFStringRef) error {
	if _cFAttributedStringReplaceString == nil {
		return symbolCallError("CFAttributedStringReplaceString", "", _cFAttributedStringReplaceStringErr)
	}
	_cFAttributedStringReplaceString(aStr, range_, replacement)
	return nil
}

// CFAttributedStringReplaceString modifies the string of an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceString(_:_:_:)
func CFAttributedStringReplaceString(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFStringRef) {
	if callErr := tryCFAttributedStringReplaceString(aStr, range_, replacement); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringSetAttribute func(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef, value CFTypeRef)
var _cFAttributedStringSetAttributeErr error

func tryCFAttributedStringSetAttribute(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef, value CFTypeRef) error {
	if _cFAttributedStringSetAttribute == nil {
		return symbolCallError("CFAttributedStringSetAttribute", "", _cFAttributedStringSetAttributeErr)
	}
	_cFAttributedStringSetAttribute(aStr, range_, attrName, value)
	return nil
}

// CFAttributedStringSetAttribute sets the value of a single attribute over the specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttribute(_:_:_:_:)
func CFAttributedStringSetAttribute(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef, value CFTypeRef) {
	if callErr := tryCFAttributedStringSetAttribute(aStr, range_, attrName, value); callErr != nil {
		panic(callErr)
	}
}

var _cFAttributedStringSetAttributes func(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFDictionaryRef, clearOtherAttributes bool)
var _cFAttributedStringSetAttributesErr error

func tryCFAttributedStringSetAttributes(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFDictionaryRef, clearOtherAttributes bool) error {
	if _cFAttributedStringSetAttributes == nil {
		return symbolCallError("CFAttributedStringSetAttributes", "", _cFAttributedStringSetAttributesErr)
	}
	_cFAttributedStringSetAttributes(aStr, range_, replacement, clearOtherAttributes)
	return nil
}

// CFAttributedStringSetAttributes sets the value of attributes of a mutable attributed string over a specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttributes(_:_:_:_:)
func CFAttributedStringSetAttributes(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFDictionaryRef, clearOtherAttributes bool) {
	if callErr := tryCFAttributedStringSetAttributes(aStr, range_, replacement, clearOtherAttributes); callErr != nil {
		panic(callErr)
	}
}

var _cFAutorelease func(arg CFTypeRef) CFTypeRef
var _cFAutoreleaseErr error

func tryCFAutorelease(arg CFTypeRef) (CFTypeRef, error) {
	if _cFAutorelease == nil {
		return 0, symbolCallError("CFAutorelease", "10.9", _cFAutoreleaseErr)
	}
	return _cFAutorelease(arg), nil
}

// CFAutorelease.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAutorelease
func CFAutorelease(arg CFTypeRef) CFTypeRef {
	result, callErr := tryCFAutorelease(arg)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagAddValue func(theBag CFMutableBagRef, value unsafe.Pointer)
var _cFBagAddValueErr error

func tryCFBagAddValue(theBag CFMutableBagRef, value unsafe.Pointer) error {
	if _cFBagAddValue == nil {
		return symbolCallError("CFBagAddValue", "", _cFBagAddValueErr)
	}
	_cFBagAddValue(theBag, value)
	return nil
}

// CFBagAddValue adds a value to a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagAddValue(_:_:)
func CFBagAddValue(theBag CFMutableBagRef, value unsafe.Pointer) {
	if callErr := tryCFBagAddValue(theBag, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBagApplyFunction func(theBag CFBagRef, applier CFBagApplierFunction, context unsafe.Pointer)
var _cFBagApplyFunctionErr error

func tryCFBagApplyFunction(theBag CFBagRef, applier CFBagApplierFunction, context unsafe.Pointer) error {
	if _cFBagApplyFunction == nil {
		return symbolCallError("CFBagApplyFunction", "", _cFBagApplyFunctionErr)
	}
	_cFBagApplyFunction(theBag, applier, context)
	return nil
}

// CFBagApplyFunction calls a function once for each value in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagApplyFunction(_:_:_:)
func CFBagApplyFunction(theBag CFBagRef, applier CFBagApplierFunction, context unsafe.Pointer) {
	if callErr := tryCFBagApplyFunction(theBag, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cFBagContainsValue func(theBag CFBagRef, value unsafe.Pointer) bool
var _cFBagContainsValueErr error

func tryCFBagContainsValue(theBag CFBagRef, value unsafe.Pointer) (bool, error) {
	if _cFBagContainsValue == nil {
		return false, symbolCallError("CFBagContainsValue", "", _cFBagContainsValueErr)
	}
	return _cFBagContainsValue(theBag, value), nil
}

// CFBagContainsValue reports whether or not a value is in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagContainsValue(_:_:)
func CFBagContainsValue(theBag CFBagRef, value unsafe.Pointer) bool {
	result, callErr := tryCFBagContainsValue(theBag, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagCreate func(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFBagCallBacks) CFBagRef
var _cFBagCreateErr error

func tryCFBagCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFBagCallBacks) (CFBagRef, error) {
	if _cFBagCreate == nil {
		return 0, symbolCallError("CFBagCreate", "", _cFBagCreateErr)
	}
	return _cFBagCreate(allocator, values, numValues, callBacks), nil
}

// CFBagCreate creates an immutable bag containing specified values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreate(_:_:_:_:)
func CFBagCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFBagCallBacks) CFBagRef {
	result, callErr := tryCFBagCreate(allocator, values, numValues, callBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagCreateCopy func(allocator CFAllocatorRef, theBag CFBagRef) CFBagRef
var _cFBagCreateCopyErr error

func tryCFBagCreateCopy(allocator CFAllocatorRef, theBag CFBagRef) (CFBagRef, error) {
	if _cFBagCreateCopy == nil {
		return 0, symbolCallError("CFBagCreateCopy", "", _cFBagCreateCopyErr)
	}
	return _cFBagCreateCopy(allocator, theBag), nil
}

// CFBagCreateCopy creates an immutable bag with the values of another bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateCopy(_:_:)
func CFBagCreateCopy(allocator CFAllocatorRef, theBag CFBagRef) CFBagRef {
	result, callErr := tryCFBagCreateCopy(allocator, theBag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagCreateMutable func(allocator CFAllocatorRef, capacity int, callBacks *CFBagCallBacks) CFMutableBagRef
var _cFBagCreateMutableErr error

func tryCFBagCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFBagCallBacks) (CFMutableBagRef, error) {
	if _cFBagCreateMutable == nil {
		return 0, symbolCallError("CFBagCreateMutable", "", _cFBagCreateMutableErr)
	}
	return _cFBagCreateMutable(allocator, capacity, callBacks), nil
}

// CFBagCreateMutable creates a new empty mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutable(_:_:_:)
func CFBagCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFBagCallBacks) CFMutableBagRef {
	result, callErr := tryCFBagCreateMutable(allocator, capacity, callBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theBag CFBagRef) CFMutableBagRef
var _cFBagCreateMutableCopyErr error

func tryCFBagCreateMutableCopy(allocator CFAllocatorRef, capacity int, theBag CFBagRef) (CFMutableBagRef, error) {
	if _cFBagCreateMutableCopy == nil {
		return 0, symbolCallError("CFBagCreateMutableCopy", "", _cFBagCreateMutableCopyErr)
	}
	return _cFBagCreateMutableCopy(allocator, capacity, theBag), nil
}

// CFBagCreateMutableCopy creates a new mutable bag with the values from another bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutableCopy(_:_:_:)
func CFBagCreateMutableCopy(allocator CFAllocatorRef, capacity int, theBag CFBagRef) CFMutableBagRef {
	result, callErr := tryCFBagCreateMutableCopy(allocator, capacity, theBag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagGetCount func(theBag CFBagRef) int
var _cFBagGetCountErr error

func tryCFBagGetCount(theBag CFBagRef) (int, error) {
	if _cFBagGetCount == nil {
		return 0, symbolCallError("CFBagGetCount", "", _cFBagGetCountErr)
	}
	return _cFBagGetCount(theBag), nil
}

// CFBagGetCount returns the number of values currently in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCount(_:)
func CFBagGetCount(theBag CFBagRef) int {
	result, callErr := tryCFBagGetCount(theBag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagGetCountOfValue func(theBag CFBagRef, value unsafe.Pointer) int
var _cFBagGetCountOfValueErr error

func tryCFBagGetCountOfValue(theBag CFBagRef, value unsafe.Pointer) (int, error) {
	if _cFBagGetCountOfValue == nil {
		return 0, symbolCallError("CFBagGetCountOfValue", "", _cFBagGetCountOfValueErr)
	}
	return _cFBagGetCountOfValue(theBag, value), nil
}

// CFBagGetCountOfValue returns the number of times a value occurs in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCountOfValue(_:_:)
func CFBagGetCountOfValue(theBag CFBagRef, value unsafe.Pointer) int {
	result, callErr := tryCFBagGetCountOfValue(theBag, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagGetTypeID func() uint
var _cFBagGetTypeIDErr error

func tryCFBagGetTypeID() (uint, error) {
	if _cFBagGetTypeID == nil {
		return 0, symbolCallError("CFBagGetTypeID", "", _cFBagGetTypeIDErr)
	}
	return _cFBagGetTypeID(), nil
}

// CFBagGetTypeID returns the type identifier for the CFBag opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetTypeID()
func CFBagGetTypeID() uint {
	result, callErr := tryCFBagGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagGetValue func(theBag CFBagRef, value unsafe.Pointer) unsafe.Pointer
var _cFBagGetValueErr error

func tryCFBagGetValue(theBag CFBagRef, value unsafe.Pointer) (unsafe.Pointer, error) {
	if _cFBagGetValue == nil {
		return nil, symbolCallError("CFBagGetValue", "", _cFBagGetValueErr)
	}
	return _cFBagGetValue(theBag, value), nil
}

// CFBagGetValue returns a requested value from a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValue(_:_:)
func CFBagGetValue(theBag CFBagRef, value unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryCFBagGetValue(theBag, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagGetValueIfPresent func(theBag CFBagRef, candidate unsafe.Pointer, value unsafe.Pointer) bool
var _cFBagGetValueIfPresentErr error

func tryCFBagGetValueIfPresent(theBag CFBagRef, candidate unsafe.Pointer, value unsafe.Pointer) (bool, error) {
	if _cFBagGetValueIfPresent == nil {
		return false, symbolCallError("CFBagGetValueIfPresent", "", _cFBagGetValueIfPresentErr)
	}
	return _cFBagGetValueIfPresent(theBag, candidate, value), nil
}

// CFBagGetValueIfPresent reports whether or not a value is in a bag, and returns that value indirectly if it exists.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValueIfPresent(_:_:_:)
func CFBagGetValueIfPresent(theBag CFBagRef, candidate unsafe.Pointer, value unsafe.Pointer) bool {
	result, callErr := tryCFBagGetValueIfPresent(theBag, candidate, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBagGetValues func(theBag CFBagRef, values unsafe.Pointer)
var _cFBagGetValuesErr error

func tryCFBagGetValues(theBag CFBagRef, values unsafe.Pointer) error {
	if _cFBagGetValues == nil {
		return symbolCallError("CFBagGetValues", "", _cFBagGetValuesErr)
	}
	_cFBagGetValues(theBag, values)
	return nil
}

// CFBagGetValues fills a buffer with values from a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValues(_:_:)
func CFBagGetValues(theBag CFBagRef, values unsafe.Pointer) {
	if callErr := tryCFBagGetValues(theBag, values); callErr != nil {
		panic(callErr)
	}
}

var _cFBagRemoveAllValues func(theBag CFMutableBagRef)
var _cFBagRemoveAllValuesErr error

func tryCFBagRemoveAllValues(theBag CFMutableBagRef) error {
	if _cFBagRemoveAllValues == nil {
		return symbolCallError("CFBagRemoveAllValues", "", _cFBagRemoveAllValuesErr)
	}
	_cFBagRemoveAllValues(theBag)
	return nil
}

// CFBagRemoveAllValues removes all values from a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveAllValues(_:)
func CFBagRemoveAllValues(theBag CFMutableBagRef) {
	if callErr := tryCFBagRemoveAllValues(theBag); callErr != nil {
		panic(callErr)
	}
}

var _cFBagRemoveValue func(theBag CFMutableBagRef, value unsafe.Pointer)
var _cFBagRemoveValueErr error

func tryCFBagRemoveValue(theBag CFMutableBagRef, value unsafe.Pointer) error {
	if _cFBagRemoveValue == nil {
		return symbolCallError("CFBagRemoveValue", "", _cFBagRemoveValueErr)
	}
	_cFBagRemoveValue(theBag, value)
	return nil
}

// CFBagRemoveValue removes a value from a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveValue(_:_:)
func CFBagRemoveValue(theBag CFMutableBagRef, value unsafe.Pointer) {
	if callErr := tryCFBagRemoveValue(theBag, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBagReplaceValue func(theBag CFMutableBagRef, value unsafe.Pointer)
var _cFBagReplaceValueErr error

func tryCFBagReplaceValue(theBag CFMutableBagRef, value unsafe.Pointer) error {
	if _cFBagReplaceValue == nil {
		return symbolCallError("CFBagReplaceValue", "", _cFBagReplaceValueErr)
	}
	_cFBagReplaceValue(theBag, value)
	return nil
}

// CFBagReplaceValue replaces a value in a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagReplaceValue(_:_:)
func CFBagReplaceValue(theBag CFMutableBagRef, value unsafe.Pointer) {
	if callErr := tryCFBagReplaceValue(theBag, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBagSetValue func(theBag CFMutableBagRef, value unsafe.Pointer)
var _cFBagSetValueErr error

func tryCFBagSetValue(theBag CFMutableBagRef, value unsafe.Pointer) error {
	if _cFBagSetValue == nil {
		return symbolCallError("CFBagSetValue", "", _cFBagSetValueErr)
	}
	_cFBagSetValue(theBag, value)
	return nil
}

// CFBagSetValue sets a value in a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagSetValue(_:_:)
func CFBagSetValue(theBag CFMutableBagRef, value unsafe.Pointer) {
	if callErr := tryCFBagSetValue(theBag, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBinaryHeapAddValue func(heap CFBinaryHeapRef, value unsafe.Pointer)
var _cFBinaryHeapAddValueErr error

func tryCFBinaryHeapAddValue(heap CFBinaryHeapRef, value unsafe.Pointer) error {
	if _cFBinaryHeapAddValue == nil {
		return symbolCallError("CFBinaryHeapAddValue", "", _cFBinaryHeapAddValueErr)
	}
	_cFBinaryHeapAddValue(heap, value)
	return nil
}

// CFBinaryHeapAddValue adds a value to a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapAddValue(_:_:)
func CFBinaryHeapAddValue(heap CFBinaryHeapRef, value unsafe.Pointer) {
	if callErr := tryCFBinaryHeapAddValue(heap, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBinaryHeapApplyFunction func(heap CFBinaryHeapRef, applier CFBinaryHeapApplierFunction, context unsafe.Pointer)
var _cFBinaryHeapApplyFunctionErr error

func tryCFBinaryHeapApplyFunction(heap CFBinaryHeapRef, applier CFBinaryHeapApplierFunction, context unsafe.Pointer) error {
	if _cFBinaryHeapApplyFunction == nil {
		return symbolCallError("CFBinaryHeapApplyFunction", "", _cFBinaryHeapApplyFunctionErr)
	}
	_cFBinaryHeapApplyFunction(heap, applier, context)
	return nil
}

// CFBinaryHeapApplyFunction iteratively applies a function to all the values in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplyFunction(_:_:_:)
func CFBinaryHeapApplyFunction(heap CFBinaryHeapRef, applier CFBinaryHeapApplierFunction, context unsafe.Pointer) {
	if callErr := tryCFBinaryHeapApplyFunction(heap, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cFBinaryHeapContainsValue func(heap CFBinaryHeapRef, value unsafe.Pointer) bool
var _cFBinaryHeapContainsValueErr error

func tryCFBinaryHeapContainsValue(heap CFBinaryHeapRef, value unsafe.Pointer) (bool, error) {
	if _cFBinaryHeapContainsValue == nil {
		return false, symbolCallError("CFBinaryHeapContainsValue", "", _cFBinaryHeapContainsValueErr)
	}
	return _cFBinaryHeapContainsValue(heap, value), nil
}

// CFBinaryHeapContainsValue returns whether a given value is in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapContainsValue(_:_:)
func CFBinaryHeapContainsValue(heap CFBinaryHeapRef, value unsafe.Pointer) bool {
	result, callErr := tryCFBinaryHeapContainsValue(heap, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapCreate func(allocator CFAllocatorRef, capacity int, callBacks *CFBinaryHeapCallBacks, compareContext *CFBinaryHeapCompareContext) CFBinaryHeapRef
var _cFBinaryHeapCreateErr error

func tryCFBinaryHeapCreate(allocator CFAllocatorRef, capacity int, callBacks *CFBinaryHeapCallBacks, compareContext *CFBinaryHeapCompareContext) (CFBinaryHeapRef, error) {
	if _cFBinaryHeapCreate == nil {
		return 0, symbolCallError("CFBinaryHeapCreate", "", _cFBinaryHeapCreateErr)
	}
	return _cFBinaryHeapCreate(allocator, capacity, callBacks, compareContext), nil
}

// CFBinaryHeapCreate creates a new mutable or fixed-mutable binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreate(_:_:_:_:)
func CFBinaryHeapCreate(allocator CFAllocatorRef, capacity int, callBacks *CFBinaryHeapCallBacks, compareContext *CFBinaryHeapCompareContext) CFBinaryHeapRef {
	result, callErr := tryCFBinaryHeapCreate(allocator, capacity, callBacks, compareContext)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapCreateCopy func(allocator CFAllocatorRef, capacity int, heap CFBinaryHeapRef) CFBinaryHeapRef
var _cFBinaryHeapCreateCopyErr error

func tryCFBinaryHeapCreateCopy(allocator CFAllocatorRef, capacity int, heap CFBinaryHeapRef) (CFBinaryHeapRef, error) {
	if _cFBinaryHeapCreateCopy == nil {
		return 0, symbolCallError("CFBinaryHeapCreateCopy", "", _cFBinaryHeapCreateCopyErr)
	}
	return _cFBinaryHeapCreateCopy(allocator, capacity, heap), nil
}

// CFBinaryHeapCreateCopy creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreateCopy(_:_:_:)
func CFBinaryHeapCreateCopy(allocator CFAllocatorRef, capacity int, heap CFBinaryHeapRef) CFBinaryHeapRef {
	result, callErr := tryCFBinaryHeapCreateCopy(allocator, capacity, heap)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapGetCount func(heap CFBinaryHeapRef) int
var _cFBinaryHeapGetCountErr error

func tryCFBinaryHeapGetCount(heap CFBinaryHeapRef) (int, error) {
	if _cFBinaryHeapGetCount == nil {
		return 0, symbolCallError("CFBinaryHeapGetCount", "", _cFBinaryHeapGetCountErr)
	}
	return _cFBinaryHeapGetCount(heap), nil
}

// CFBinaryHeapGetCount returns the number of values currently in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCount(_:)
func CFBinaryHeapGetCount(heap CFBinaryHeapRef) int {
	result, callErr := tryCFBinaryHeapGetCount(heap)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapGetCountOfValue func(heap CFBinaryHeapRef, value unsafe.Pointer) int
var _cFBinaryHeapGetCountOfValueErr error

func tryCFBinaryHeapGetCountOfValue(heap CFBinaryHeapRef, value unsafe.Pointer) (int, error) {
	if _cFBinaryHeapGetCountOfValue == nil {
		return 0, symbolCallError("CFBinaryHeapGetCountOfValue", "", _cFBinaryHeapGetCountOfValueErr)
	}
	return _cFBinaryHeapGetCountOfValue(heap, value), nil
}

// CFBinaryHeapGetCountOfValue counts the number of times a given value occurs in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCountOfValue(_:_:)
func CFBinaryHeapGetCountOfValue(heap CFBinaryHeapRef, value unsafe.Pointer) int {
	result, callErr := tryCFBinaryHeapGetCountOfValue(heap, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapGetMinimum func(heap CFBinaryHeapRef) unsafe.Pointer
var _cFBinaryHeapGetMinimumErr error

func tryCFBinaryHeapGetMinimum(heap CFBinaryHeapRef) (unsafe.Pointer, error) {
	if _cFBinaryHeapGetMinimum == nil {
		return nil, symbolCallError("CFBinaryHeapGetMinimum", "", _cFBinaryHeapGetMinimumErr)
	}
	return _cFBinaryHeapGetMinimum(heap), nil
}

// CFBinaryHeapGetMinimum returns the minimum value in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimum(_:)
func CFBinaryHeapGetMinimum(heap CFBinaryHeapRef) unsafe.Pointer {
	result, callErr := tryCFBinaryHeapGetMinimum(heap)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapGetMinimumIfPresent func(heap CFBinaryHeapRef, value unsafe.Pointer) bool
var _cFBinaryHeapGetMinimumIfPresentErr error

func tryCFBinaryHeapGetMinimumIfPresent(heap CFBinaryHeapRef, value unsafe.Pointer) (bool, error) {
	if _cFBinaryHeapGetMinimumIfPresent == nil {
		return false, symbolCallError("CFBinaryHeapGetMinimumIfPresent", "", _cFBinaryHeapGetMinimumIfPresentErr)
	}
	return _cFBinaryHeapGetMinimumIfPresent(heap, value), nil
}

// CFBinaryHeapGetMinimumIfPresent returns the minimum value in a binary heap, if present.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimumIfPresent(_:_:)
func CFBinaryHeapGetMinimumIfPresent(heap CFBinaryHeapRef, value unsafe.Pointer) bool {
	result, callErr := tryCFBinaryHeapGetMinimumIfPresent(heap, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapGetTypeID func() uint
var _cFBinaryHeapGetTypeIDErr error

func tryCFBinaryHeapGetTypeID() (uint, error) {
	if _cFBinaryHeapGetTypeID == nil {
		return 0, symbolCallError("CFBinaryHeapGetTypeID", "", _cFBinaryHeapGetTypeIDErr)
	}
	return _cFBinaryHeapGetTypeID(), nil
}

// CFBinaryHeapGetTypeID returns the type identifier of the [CFBinaryHeap] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetTypeID()
func CFBinaryHeapGetTypeID() uint {
	result, callErr := tryCFBinaryHeapGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBinaryHeapGetValues func(heap CFBinaryHeapRef, values unsafe.Pointer)
var _cFBinaryHeapGetValuesErr error

func tryCFBinaryHeapGetValues(heap CFBinaryHeapRef, values unsafe.Pointer) error {
	if _cFBinaryHeapGetValues == nil {
		return symbolCallError("CFBinaryHeapGetValues", "", _cFBinaryHeapGetValuesErr)
	}
	_cFBinaryHeapGetValues(heap, values)
	return nil
}

// CFBinaryHeapGetValues copies all the values from a binary heap into a sorted C array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetValues(_:_:)
func CFBinaryHeapGetValues(heap CFBinaryHeapRef, values unsafe.Pointer) {
	if callErr := tryCFBinaryHeapGetValues(heap, values); callErr != nil {
		panic(callErr)
	}
}

var _cFBinaryHeapRemoveAllValues func(heap CFBinaryHeapRef)
var _cFBinaryHeapRemoveAllValuesErr error

func tryCFBinaryHeapRemoveAllValues(heap CFBinaryHeapRef) error {
	if _cFBinaryHeapRemoveAllValues == nil {
		return symbolCallError("CFBinaryHeapRemoveAllValues", "", _cFBinaryHeapRemoveAllValuesErr)
	}
	_cFBinaryHeapRemoveAllValues(heap)
	return nil
}

// CFBinaryHeapRemoveAllValues removes all values from a binary heap, making it empty.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveAllValues(_:)
func CFBinaryHeapRemoveAllValues(heap CFBinaryHeapRef) {
	if callErr := tryCFBinaryHeapRemoveAllValues(heap); callErr != nil {
		panic(callErr)
	}
}

var _cFBinaryHeapRemoveMinimumValue func(heap CFBinaryHeapRef)
var _cFBinaryHeapRemoveMinimumValueErr error

func tryCFBinaryHeapRemoveMinimumValue(heap CFBinaryHeapRef) error {
	if _cFBinaryHeapRemoveMinimumValue == nil {
		return symbolCallError("CFBinaryHeapRemoveMinimumValue", "", _cFBinaryHeapRemoveMinimumValueErr)
	}
	_cFBinaryHeapRemoveMinimumValue(heap)
	return nil
}

// CFBinaryHeapRemoveMinimumValue removes the minimum value from a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveMinimumValue(_:)
func CFBinaryHeapRemoveMinimumValue(heap CFBinaryHeapRef) {
	if callErr := tryCFBinaryHeapRemoveMinimumValue(heap); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorContainsBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) bool
var _cFBitVectorContainsBitErr error

func tryCFBitVectorContainsBit(bv CFBitVectorRef, range_ CFRange, value CFBit) (bool, error) {
	if _cFBitVectorContainsBit == nil {
		return false, symbolCallError("CFBitVectorContainsBit", "", _cFBitVectorContainsBitErr)
	}
	return _cFBitVectorContainsBit(bv, range_, value), nil
}

// CFBitVectorContainsBit returns whether a bit vector contains a particular bit value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorContainsBit(_:_:_:)
func CFBitVectorContainsBit(bv CFBitVectorRef, range_ CFRange, value CFBit) bool {
	result, callErr := tryCFBitVectorContainsBit(bv, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorCreate func(allocator CFAllocatorRef, bytes *uint8, numBits int) CFBitVectorRef
var _cFBitVectorCreateErr error

func tryCFBitVectorCreate(allocator CFAllocatorRef, bytes *uint8, numBits int) (CFBitVectorRef, error) {
	if _cFBitVectorCreate == nil {
		return 0, symbolCallError("CFBitVectorCreate", "", _cFBitVectorCreateErr)
	}
	return _cFBitVectorCreate(allocator, bytes, numBits), nil
}

// CFBitVectorCreate creates an immutable bit vector from a block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreate(_:_:_:)
func CFBitVectorCreate(allocator CFAllocatorRef, bytes *uint8, numBits int) CFBitVectorRef {
	result, callErr := tryCFBitVectorCreate(allocator, bytes, numBits)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorCreateCopy func(allocator CFAllocatorRef, bv CFBitVectorRef) CFBitVectorRef
var _cFBitVectorCreateCopyErr error

func tryCFBitVectorCreateCopy(allocator CFAllocatorRef, bv CFBitVectorRef) (CFBitVectorRef, error) {
	if _cFBitVectorCreateCopy == nil {
		return 0, symbolCallError("CFBitVectorCreateCopy", "", _cFBitVectorCreateCopyErr)
	}
	return _cFBitVectorCreateCopy(allocator, bv), nil
}

// CFBitVectorCreateCopy creates an immutable bit vector that is a copy of another bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateCopy(_:_:)
func CFBitVectorCreateCopy(allocator CFAllocatorRef, bv CFBitVectorRef) CFBitVectorRef {
	result, callErr := tryCFBitVectorCreateCopy(allocator, bv)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorCreateMutable func(allocator CFAllocatorRef, capacity int) CFMutableBitVectorRef
var _cFBitVectorCreateMutableErr error

func tryCFBitVectorCreateMutable(allocator CFAllocatorRef, capacity int) (CFMutableBitVectorRef, error) {
	if _cFBitVectorCreateMutable == nil {
		return 0, symbolCallError("CFBitVectorCreateMutable", "", _cFBitVectorCreateMutableErr)
	}
	return _cFBitVectorCreateMutable(allocator, capacity), nil
}

// CFBitVectorCreateMutable creates a mutable bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutable(_:_:)
func CFBitVectorCreateMutable(allocator CFAllocatorRef, capacity int) CFMutableBitVectorRef {
	result, callErr := tryCFBitVectorCreateMutable(allocator, capacity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorCreateMutableCopy func(allocator CFAllocatorRef, capacity int, bv CFBitVectorRef) CFMutableBitVectorRef
var _cFBitVectorCreateMutableCopyErr error

func tryCFBitVectorCreateMutableCopy(allocator CFAllocatorRef, capacity int, bv CFBitVectorRef) (CFMutableBitVectorRef, error) {
	if _cFBitVectorCreateMutableCopy == nil {
		return 0, symbolCallError("CFBitVectorCreateMutableCopy", "", _cFBitVectorCreateMutableCopyErr)
	}
	return _cFBitVectorCreateMutableCopy(allocator, capacity, bv), nil
}

// CFBitVectorCreateMutableCopy creates a new mutable bit vector from a pre-existing bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutableCopy(_:_:_:)
func CFBitVectorCreateMutableCopy(allocator CFAllocatorRef, capacity int, bv CFBitVectorRef) CFMutableBitVectorRef {
	result, callErr := tryCFBitVectorCreateMutableCopy(allocator, capacity, bv)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorFlipBitAtIndex func(bv CFMutableBitVectorRef, idx int)
var _cFBitVectorFlipBitAtIndexErr error

func tryCFBitVectorFlipBitAtIndex(bv CFMutableBitVectorRef, idx int) error {
	if _cFBitVectorFlipBitAtIndex == nil {
		return symbolCallError("CFBitVectorFlipBitAtIndex", "", _cFBitVectorFlipBitAtIndexErr)
	}
	_cFBitVectorFlipBitAtIndex(bv, idx)
	return nil
}

// CFBitVectorFlipBitAtIndex flips a bit value in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBitAtIndex(_:_:)
func CFBitVectorFlipBitAtIndex(bv CFMutableBitVectorRef, idx int) {
	if callErr := tryCFBitVectorFlipBitAtIndex(bv, idx); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorFlipBits func(bv CFMutableBitVectorRef, range_ CFRange)
var _cFBitVectorFlipBitsErr error

func tryCFBitVectorFlipBits(bv CFMutableBitVectorRef, range_ CFRange) error {
	if _cFBitVectorFlipBits == nil {
		return symbolCallError("CFBitVectorFlipBits", "", _cFBitVectorFlipBitsErr)
	}
	_cFBitVectorFlipBits(bv, range_)
	return nil
}

// CFBitVectorFlipBits flips a range of bit values in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBits(_:_:)
func CFBitVectorFlipBits(bv CFMutableBitVectorRef, range_ CFRange) {
	if callErr := tryCFBitVectorFlipBits(bv, range_); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorGetBitAtIndex func(bv CFBitVectorRef, idx int) CFBit
var _cFBitVectorGetBitAtIndexErr error

func tryCFBitVectorGetBitAtIndex(bv CFBitVectorRef, idx int) (CFBit, error) {
	if _cFBitVectorGetBitAtIndex == nil {
		return *new(CFBit), symbolCallError("CFBitVectorGetBitAtIndex", "", _cFBitVectorGetBitAtIndexErr)
	}
	return _cFBitVectorGetBitAtIndex(bv, idx), nil
}

// CFBitVectorGetBitAtIndex returns the bit value at a given index in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBitAtIndex(_:_:)
func CFBitVectorGetBitAtIndex(bv CFBitVectorRef, idx int) CFBit {
	result, callErr := tryCFBitVectorGetBitAtIndex(bv, idx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorGetBits func(bv CFBitVectorRef, range_ CFRange, bytes *uint8)
var _cFBitVectorGetBitsErr error

func tryCFBitVectorGetBits(bv CFBitVectorRef, range_ CFRange, bytes *uint8) error {
	if _cFBitVectorGetBits == nil {
		return symbolCallError("CFBitVectorGetBits", "", _cFBitVectorGetBitsErr)
	}
	_cFBitVectorGetBits(bv, range_, bytes)
	return nil
}

// CFBitVectorGetBits returns the bit values in a range of indices in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBits(_:_:_:)
func CFBitVectorGetBits(bv CFBitVectorRef, range_ CFRange, bytes *uint8) {
	if callErr := tryCFBitVectorGetBits(bv, range_, bytes); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorGetCount func(bv CFBitVectorRef) int
var _cFBitVectorGetCountErr error

func tryCFBitVectorGetCount(bv CFBitVectorRef) (int, error) {
	if _cFBitVectorGetCount == nil {
		return 0, symbolCallError("CFBitVectorGetCount", "", _cFBitVectorGetCountErr)
	}
	return _cFBitVectorGetCount(bv), nil
}

// CFBitVectorGetCount returns the number of bit values in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCount(_:)
func CFBitVectorGetCount(bv CFBitVectorRef) int {
	result, callErr := tryCFBitVectorGetCount(bv)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorGetCountOfBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) int
var _cFBitVectorGetCountOfBitErr error

func tryCFBitVectorGetCountOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) (int, error) {
	if _cFBitVectorGetCountOfBit == nil {
		return 0, symbolCallError("CFBitVectorGetCountOfBit", "", _cFBitVectorGetCountOfBitErr)
	}
	return _cFBitVectorGetCountOfBit(bv, range_, value), nil
}

// CFBitVectorGetCountOfBit counts the number of times a certain bit value occurs within a range of bits in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCountOfBit(_:_:_:)
func CFBitVectorGetCountOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) int {
	result, callErr := tryCFBitVectorGetCountOfBit(bv, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorGetFirstIndexOfBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) int
var _cFBitVectorGetFirstIndexOfBitErr error

func tryCFBitVectorGetFirstIndexOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) (int, error) {
	if _cFBitVectorGetFirstIndexOfBit == nil {
		return 0, symbolCallError("CFBitVectorGetFirstIndexOfBit", "", _cFBitVectorGetFirstIndexOfBitErr)
	}
	return _cFBitVectorGetFirstIndexOfBit(bv, range_, value), nil
}

// CFBitVectorGetFirstIndexOfBit locates the first occurrence of a certain bit value within a range of bits in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetFirstIndexOfBit(_:_:_:)
func CFBitVectorGetFirstIndexOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) int {
	result, callErr := tryCFBitVectorGetFirstIndexOfBit(bv, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorGetLastIndexOfBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) int
var _cFBitVectorGetLastIndexOfBitErr error

func tryCFBitVectorGetLastIndexOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) (int, error) {
	if _cFBitVectorGetLastIndexOfBit == nil {
		return 0, symbolCallError("CFBitVectorGetLastIndexOfBit", "", _cFBitVectorGetLastIndexOfBitErr)
	}
	return _cFBitVectorGetLastIndexOfBit(bv, range_, value), nil
}

// CFBitVectorGetLastIndexOfBit locates the last occurrence of a certain bit value within a range of bits in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetLastIndexOfBit(_:_:_:)
func CFBitVectorGetLastIndexOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) int {
	result, callErr := tryCFBitVectorGetLastIndexOfBit(bv, range_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorGetTypeID func() uint
var _cFBitVectorGetTypeIDErr error

func tryCFBitVectorGetTypeID() (uint, error) {
	if _cFBitVectorGetTypeID == nil {
		return 0, symbolCallError("CFBitVectorGetTypeID", "", _cFBitVectorGetTypeIDErr)
	}
	return _cFBitVectorGetTypeID(), nil
}

// CFBitVectorGetTypeID returns the type identifier for the CFBitVector opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetTypeID()
func CFBitVectorGetTypeID() uint {
	result, callErr := tryCFBitVectorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBitVectorSetAllBits func(bv CFMutableBitVectorRef, value CFBit)
var _cFBitVectorSetAllBitsErr error

func tryCFBitVectorSetAllBits(bv CFMutableBitVectorRef, value CFBit) error {
	if _cFBitVectorSetAllBits == nil {
		return symbolCallError("CFBitVectorSetAllBits", "", _cFBitVectorSetAllBitsErr)
	}
	_cFBitVectorSetAllBits(bv, value)
	return nil
}

// CFBitVectorSetAllBits sets all bits in a bit vector to a particular value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetAllBits(_:_:)
func CFBitVectorSetAllBits(bv CFMutableBitVectorRef, value CFBit) {
	if callErr := tryCFBitVectorSetAllBits(bv, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorSetBitAtIndex func(bv CFMutableBitVectorRef, idx int, value CFBit)
var _cFBitVectorSetBitAtIndexErr error

func tryCFBitVectorSetBitAtIndex(bv CFMutableBitVectorRef, idx int, value CFBit) error {
	if _cFBitVectorSetBitAtIndex == nil {
		return symbolCallError("CFBitVectorSetBitAtIndex", "", _cFBitVectorSetBitAtIndexErr)
	}
	_cFBitVectorSetBitAtIndex(bv, idx, value)
	return nil
}

// CFBitVectorSetBitAtIndex sets the value of a particular bit in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBitAtIndex(_:_:_:)
func CFBitVectorSetBitAtIndex(bv CFMutableBitVectorRef, idx int, value CFBit) {
	if callErr := tryCFBitVectorSetBitAtIndex(bv, idx, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorSetBits func(bv CFMutableBitVectorRef, range_ CFRange, value CFBit)
var _cFBitVectorSetBitsErr error

func tryCFBitVectorSetBits(bv CFMutableBitVectorRef, range_ CFRange, value CFBit) error {
	if _cFBitVectorSetBits == nil {
		return symbolCallError("CFBitVectorSetBits", "", _cFBitVectorSetBitsErr)
	}
	_cFBitVectorSetBits(bv, range_, value)
	return nil
}

// CFBitVectorSetBits sets a range of bits in a bit vector to a particular value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBits(_:_:_:)
func CFBitVectorSetBits(bv CFMutableBitVectorRef, range_ CFRange, value CFBit) {
	if callErr := tryCFBitVectorSetBits(bv, range_, value); callErr != nil {
		panic(callErr)
	}
}

var _cFBitVectorSetCount func(bv CFMutableBitVectorRef, count int)
var _cFBitVectorSetCountErr error

func tryCFBitVectorSetCount(bv CFMutableBitVectorRef, count int) error {
	if _cFBitVectorSetCount == nil {
		return symbolCallError("CFBitVectorSetCount", "", _cFBitVectorSetCountErr)
	}
	_cFBitVectorSetCount(bv, count)
	return nil
}

// CFBitVectorSetCount changes the size of a mutable bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetCount(_:_:)
func CFBitVectorSetCount(bv CFMutableBitVectorRef, count int) {
	if callErr := tryCFBitVectorSetCount(bv, count); callErr != nil {
		panic(callErr)
	}
}

var _cFBooleanGetTypeID func() uint
var _cFBooleanGetTypeIDErr error

func tryCFBooleanGetTypeID() (uint, error) {
	if _cFBooleanGetTypeID == nil {
		return 0, symbolCallError("CFBooleanGetTypeID", "", _cFBooleanGetTypeIDErr)
	}
	return _cFBooleanGetTypeID(), nil
}

// CFBooleanGetTypeID returns the Core Foundation type identifier for the CFBoolean opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetTypeID()
func CFBooleanGetTypeID() uint {
	result, callErr := tryCFBooleanGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBooleanGetValue func(boolean CFBooleanRef) bool
var _cFBooleanGetValueErr error

func tryCFBooleanGetValue(boolean CFBooleanRef) (bool, error) {
	if _cFBooleanGetValue == nil {
		return false, symbolCallError("CFBooleanGetValue", "", _cFBooleanGetValueErr)
	}
	return _cFBooleanGetValue(boolean), nil
}

// CFBooleanGetValue returns the value of a CFBoolean object as a standard C type [Boolean].
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetValue(_:)
func CFBooleanGetValue(boolean CFBooleanRef) bool {
	result, callErr := tryCFBooleanGetValue(boolean)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyAuxiliaryExecutableURL func(bundle CFBundleRef, executableName CFStringRef) CFURLRef
var _cFBundleCopyAuxiliaryExecutableURLErr error

func tryCFBundleCopyAuxiliaryExecutableURL(bundle CFBundleRef, executableName CFStringRef) (CFURLRef, error) {
	if _cFBundleCopyAuxiliaryExecutableURL == nil {
		return 0, symbolCallError("CFBundleCopyAuxiliaryExecutableURL", "", _cFBundleCopyAuxiliaryExecutableURLErr)
	}
	return _cFBundleCopyAuxiliaryExecutableURL(bundle, executableName), nil
}

// CFBundleCopyAuxiliaryExecutableURL returns the location of a bundle’s auxiliary executable code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyAuxiliaryExecutableURL(_:_:)
func CFBundleCopyAuxiliaryExecutableURL(bundle CFBundleRef, executableName CFStringRef) CFURLRef {
	result, callErr := tryCFBundleCopyAuxiliaryExecutableURL(bundle, executableName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyBuiltInPlugInsURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopyBuiltInPlugInsURLErr error

func tryCFBundleCopyBuiltInPlugInsURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopyBuiltInPlugInsURL == nil {
		return 0, symbolCallError("CFBundleCopyBuiltInPlugInsURL", "", _cFBundleCopyBuiltInPlugInsURLErr)
	}
	return _cFBundleCopyBuiltInPlugInsURL(bundle), nil
}

// CFBundleCopyBuiltInPlugInsURL returns the location of a bundle’s built in plug-in.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBuiltInPlugInsURL(_:)
func CFBundleCopyBuiltInPlugInsURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopyBuiltInPlugInsURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyBundleLocalizations func(bundle CFBundleRef) CFArrayRef
var _cFBundleCopyBundleLocalizationsErr error

func tryCFBundleCopyBundleLocalizations(bundle CFBundleRef) (CFArrayRef, error) {
	if _cFBundleCopyBundleLocalizations == nil {
		return 0, symbolCallError("CFBundleCopyBundleLocalizations", "", _cFBundleCopyBundleLocalizationsErr)
	}
	return _cFBundleCopyBundleLocalizations(bundle), nil
}

// CFBundleCopyBundleLocalizations returns an array containing a bundle’s localizations.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleLocalizations(_:)
func CFBundleCopyBundleLocalizations(bundle CFBundleRef) CFArrayRef {
	result, callErr := tryCFBundleCopyBundleLocalizations(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyBundleURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopyBundleURLErr error

func tryCFBundleCopyBundleURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopyBundleURL == nil {
		return 0, symbolCallError("CFBundleCopyBundleURL", "", _cFBundleCopyBundleURLErr)
	}
	return _cFBundleCopyBundleURL(bundle), nil
}

// CFBundleCopyBundleURL returns the location of a bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleURL(_:)
func CFBundleCopyBundleURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopyBundleURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyExecutableArchitectures func(bundle CFBundleRef) CFArrayRef
var _cFBundleCopyExecutableArchitecturesErr error

func tryCFBundleCopyExecutableArchitectures(bundle CFBundleRef) (CFArrayRef, error) {
	if _cFBundleCopyExecutableArchitectures == nil {
		return 0, symbolCallError("CFBundleCopyExecutableArchitectures", "10.5", _cFBundleCopyExecutableArchitecturesErr)
	}
	return _cFBundleCopyExecutableArchitectures(bundle), nil
}

// CFBundleCopyExecutableArchitectures returns an array of CFNumbers representing the architectures a given bundle provides.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitectures(_:)
func CFBundleCopyExecutableArchitectures(bundle CFBundleRef) CFArrayRef {
	result, callErr := tryCFBundleCopyExecutableArchitectures(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyExecutableArchitecturesForURL func(url CFURLRef) CFArrayRef
var _cFBundleCopyExecutableArchitecturesForURLErr error

func tryCFBundleCopyExecutableArchitecturesForURL(url CFURLRef) (CFArrayRef, error) {
	if _cFBundleCopyExecutableArchitecturesForURL == nil {
		return 0, symbolCallError("CFBundleCopyExecutableArchitecturesForURL", "10.5", _cFBundleCopyExecutableArchitecturesForURLErr)
	}
	return _cFBundleCopyExecutableArchitecturesForURL(url), nil
}

// CFBundleCopyExecutableArchitecturesForURL returns an array of CFNumbers representing the architectures a given URL provides.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitecturesForURL(_:)
func CFBundleCopyExecutableArchitecturesForURL(url CFURLRef) CFArrayRef {
	result, callErr := tryCFBundleCopyExecutableArchitecturesForURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyExecutableURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopyExecutableURLErr error

func tryCFBundleCopyExecutableURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopyExecutableURL == nil {
		return 0, symbolCallError("CFBundleCopyExecutableURL", "", _cFBundleCopyExecutableURLErr)
	}
	return _cFBundleCopyExecutableURL(bundle), nil
}

// CFBundleCopyExecutableURL returns the location of a bundle’s main executable code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableURL(_:)
func CFBundleCopyExecutableURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopyExecutableURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyInfoDictionaryForURL func(url CFURLRef) CFDictionaryRef
var _cFBundleCopyInfoDictionaryForURLErr error

func tryCFBundleCopyInfoDictionaryForURL(url CFURLRef) (CFDictionaryRef, error) {
	if _cFBundleCopyInfoDictionaryForURL == nil {
		return 0, symbolCallError("CFBundleCopyInfoDictionaryForURL", "", _cFBundleCopyInfoDictionaryForURLErr)
	}
	return _cFBundleCopyInfoDictionaryForURL(url), nil
}

// CFBundleCopyInfoDictionaryForURL returns the information dictionary for a given URL location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryForURL(_:)
func CFBundleCopyInfoDictionaryForURL(url CFURLRef) CFDictionaryRef {
	result, callErr := tryCFBundleCopyInfoDictionaryForURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyInfoDictionaryInDirectory func(bundleURL CFURLRef) CFDictionaryRef
var _cFBundleCopyInfoDictionaryInDirectoryErr error

func tryCFBundleCopyInfoDictionaryInDirectory(bundleURL CFURLRef) (CFDictionaryRef, error) {
	if _cFBundleCopyInfoDictionaryInDirectory == nil {
		return 0, symbolCallError("CFBundleCopyInfoDictionaryInDirectory", "", _cFBundleCopyInfoDictionaryInDirectoryErr)
	}
	return _cFBundleCopyInfoDictionaryInDirectory(bundleURL), nil
}

// CFBundleCopyInfoDictionaryInDirectory returns a bundle’s information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryInDirectory(_:)
func CFBundleCopyInfoDictionaryInDirectory(bundleURL CFURLRef) CFDictionaryRef {
	result, callErr := tryCFBundleCopyInfoDictionaryInDirectory(bundleURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyLocalizationsForPreferences func(locArray CFArrayRef, prefArray CFArrayRef) CFArrayRef
var _cFBundleCopyLocalizationsForPreferencesErr error

func tryCFBundleCopyLocalizationsForPreferences(locArray CFArrayRef, prefArray CFArrayRef) (CFArrayRef, error) {
	if _cFBundleCopyLocalizationsForPreferences == nil {
		return 0, symbolCallError("CFBundleCopyLocalizationsForPreferences", "", _cFBundleCopyLocalizationsForPreferencesErr)
	}
	return _cFBundleCopyLocalizationsForPreferences(locArray, prefArray), nil
}

// CFBundleCopyLocalizationsForPreferences given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForPreferences(_:_:)
func CFBundleCopyLocalizationsForPreferences(locArray CFArrayRef, prefArray CFArrayRef) CFArrayRef {
	result, callErr := tryCFBundleCopyLocalizationsForPreferences(locArray, prefArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyLocalizationsForURL func(url CFURLRef) CFArrayRef
var _cFBundleCopyLocalizationsForURLErr error

func tryCFBundleCopyLocalizationsForURL(url CFURLRef) (CFArrayRef, error) {
	if _cFBundleCopyLocalizationsForURL == nil {
		return 0, symbolCallError("CFBundleCopyLocalizationsForURL", "", _cFBundleCopyLocalizationsForURLErr)
	}
	return _cFBundleCopyLocalizationsForURL(url), nil
}

// CFBundleCopyLocalizationsForURL returns an array containing the localizations for a bundle or executable at a particular location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForURL(_:)
func CFBundleCopyLocalizationsForURL(url CFURLRef) CFArrayRef {
	result, callErr := tryCFBundleCopyLocalizationsForURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyLocalizedString func(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef) CFStringRef
var _cFBundleCopyLocalizedStringErr error

func tryCFBundleCopyLocalizedString(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef) (CFStringRef, error) {
	if _cFBundleCopyLocalizedString == nil {
		return 0, symbolCallError("CFBundleCopyLocalizedString", "", _cFBundleCopyLocalizedStringErr)
	}
	return _cFBundleCopyLocalizedString(bundle, key, value, tableName), nil
}

// CFBundleCopyLocalizedString returns a localized string from a bundle’s strings file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedString(_:_:_:_:)
func CFBundleCopyLocalizedString(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef) CFStringRef {
	result, callErr := tryCFBundleCopyLocalizedString(bundle, key, value, tableName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyLocalizedStringForLocalizations func(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef, localizations CFArrayRef) CFStringRef
var _cFBundleCopyLocalizedStringForLocalizationsErr error

func tryCFBundleCopyLocalizedStringForLocalizations(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef, localizations CFArrayRef) (CFStringRef, error) {
	if _cFBundleCopyLocalizedStringForLocalizations == nil {
		return 0, symbolCallError("CFBundleCopyLocalizedStringForLocalizations", "15.4", _cFBundleCopyLocalizedStringForLocalizationsErr)
	}
	return _cFBundleCopyLocalizedStringForLocalizations(bundle, key, value, tableName, localizations), nil
}

// CFBundleCopyLocalizedStringForLocalizations returns a localized string from a bundle’s strings file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedStringForLocalizations(_:_:_:_:_:)
func CFBundleCopyLocalizedStringForLocalizations(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef, localizations CFArrayRef) CFStringRef {
	result, callErr := tryCFBundleCopyLocalizedStringForLocalizations(bundle, key, value, tableName, localizations)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyPreferredLocalizationsFromArray func(locArray CFArrayRef) CFArrayRef
var _cFBundleCopyPreferredLocalizationsFromArrayErr error

func tryCFBundleCopyPreferredLocalizationsFromArray(locArray CFArrayRef) (CFArrayRef, error) {
	if _cFBundleCopyPreferredLocalizationsFromArray == nil {
		return 0, symbolCallError("CFBundleCopyPreferredLocalizationsFromArray", "", _cFBundleCopyPreferredLocalizationsFromArrayErr)
	}
	return _cFBundleCopyPreferredLocalizationsFromArray(locArray), nil
}

// CFBundleCopyPreferredLocalizationsFromArray given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPreferredLocalizationsFromArray(_:)
func CFBundleCopyPreferredLocalizationsFromArray(locArray CFArrayRef) CFArrayRef {
	result, callErr := tryCFBundleCopyPreferredLocalizationsFromArray(locArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyPrivateFrameworksURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopyPrivateFrameworksURLErr error

func tryCFBundleCopyPrivateFrameworksURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopyPrivateFrameworksURL == nil {
		return 0, symbolCallError("CFBundleCopyPrivateFrameworksURL", "", _cFBundleCopyPrivateFrameworksURLErr)
	}
	return _cFBundleCopyPrivateFrameworksURL(bundle), nil
}

// CFBundleCopyPrivateFrameworksURL returns the location of a bundle’s private Frameworks directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPrivateFrameworksURL(_:)
func CFBundleCopyPrivateFrameworksURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopyPrivateFrameworksURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourceURL func(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef
var _cFBundleCopyResourceURLErr error

func tryCFBundleCopyResourceURL(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) (CFURLRef, error) {
	if _cFBundleCopyResourceURL == nil {
		return 0, symbolCallError("CFBundleCopyResourceURL", "", _cFBundleCopyResourceURLErr)
	}
	return _cFBundleCopyResourceURL(bundle, resourceName, resourceType, subDirName), nil
}

// CFBundleCopyResourceURL returns the location of a resource contained in the specified bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURL(_:_:_:_:)
func CFBundleCopyResourceURL(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef {
	result, callErr := tryCFBundleCopyResourceURL(bundle, resourceName, resourceType, subDirName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourceURLForLocalization func(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFURLRef
var _cFBundleCopyResourceURLForLocalizationErr error

func tryCFBundleCopyResourceURLForLocalization(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) (CFURLRef, error) {
	if _cFBundleCopyResourceURLForLocalization == nil {
		return 0, symbolCallError("CFBundleCopyResourceURLForLocalization", "", _cFBundleCopyResourceURLForLocalizationErr)
	}
	return _cFBundleCopyResourceURLForLocalization(bundle, resourceName, resourceType, subDirName, localizationName), nil
}

// CFBundleCopyResourceURLForLocalization returns the location of a localized resource in a bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLForLocalization(_:_:_:_:_:)
func CFBundleCopyResourceURLForLocalization(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFURLRef {
	result, callErr := tryCFBundleCopyResourceURLForLocalization(bundle, resourceName, resourceType, subDirName, localizationName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourceURLInDirectory func(bundleURL CFURLRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef
var _cFBundleCopyResourceURLInDirectoryErr error

func tryCFBundleCopyResourceURLInDirectory(bundleURL CFURLRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) (CFURLRef, error) {
	if _cFBundleCopyResourceURLInDirectory == nil {
		return 0, symbolCallError("CFBundleCopyResourceURLInDirectory", "", _cFBundleCopyResourceURLInDirectoryErr)
	}
	return _cFBundleCopyResourceURLInDirectory(bundleURL, resourceName, resourceType, subDirName), nil
}

// CFBundleCopyResourceURLInDirectory returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLInDirectory(_:_:_:_:)
func CFBundleCopyResourceURLInDirectory(bundleURL CFURLRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef {
	result, callErr := tryCFBundleCopyResourceURLInDirectory(bundleURL, resourceName, resourceType, subDirName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourceURLsOfType func(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef
var _cFBundleCopyResourceURLsOfTypeErr error

func tryCFBundleCopyResourceURLsOfType(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef) (CFArrayRef, error) {
	if _cFBundleCopyResourceURLsOfType == nil {
		return 0, symbolCallError("CFBundleCopyResourceURLsOfType", "", _cFBundleCopyResourceURLsOfTypeErr)
	}
	return _cFBundleCopyResourceURLsOfType(bundle, resourceType, subDirName), nil
}

// CFBundleCopyResourceURLsOfType assembles an array of URLs specifying all of the resources of the specified type found in a bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfType(_:_:_:)
func CFBundleCopyResourceURLsOfType(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef {
	result, callErr := tryCFBundleCopyResourceURLsOfType(bundle, resourceType, subDirName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourceURLsOfTypeForLocalization func(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFArrayRef
var _cFBundleCopyResourceURLsOfTypeForLocalizationErr error

func tryCFBundleCopyResourceURLsOfTypeForLocalization(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) (CFArrayRef, error) {
	if _cFBundleCopyResourceURLsOfTypeForLocalization == nil {
		return 0, symbolCallError("CFBundleCopyResourceURLsOfTypeForLocalization", "", _cFBundleCopyResourceURLsOfTypeForLocalizationErr)
	}
	return _cFBundleCopyResourceURLsOfTypeForLocalization(bundle, resourceType, subDirName, localizationName), nil
}

// CFBundleCopyResourceURLsOfTypeForLocalization returns an array containing copies of the URL locations for a specified bundle, resource, and localization name.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeForLocalization(_:_:_:_:)
func CFBundleCopyResourceURLsOfTypeForLocalization(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFArrayRef {
	result, callErr := tryCFBundleCopyResourceURLsOfTypeForLocalization(bundle, resourceType, subDirName, localizationName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourceURLsOfTypeInDirectory func(bundleURL CFURLRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef
var _cFBundleCopyResourceURLsOfTypeInDirectoryErr error

func tryCFBundleCopyResourceURLsOfTypeInDirectory(bundleURL CFURLRef, resourceType CFStringRef, subDirName CFStringRef) (CFArrayRef, error) {
	if _cFBundleCopyResourceURLsOfTypeInDirectory == nil {
		return 0, symbolCallError("CFBundleCopyResourceURLsOfTypeInDirectory", "", _cFBundleCopyResourceURLsOfTypeInDirectoryErr)
	}
	return _cFBundleCopyResourceURLsOfTypeInDirectory(bundleURL, resourceType, subDirName), nil
}

// CFBundleCopyResourceURLsOfTypeInDirectory returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeInDirectory(_:_:_:)
func CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL CFURLRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef {
	result, callErr := tryCFBundleCopyResourceURLsOfTypeInDirectory(bundleURL, resourceType, subDirName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopyResourcesDirectoryURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopyResourcesDirectoryURLErr error

func tryCFBundleCopyResourcesDirectoryURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopyResourcesDirectoryURL == nil {
		return 0, symbolCallError("CFBundleCopyResourcesDirectoryURL", "", _cFBundleCopyResourcesDirectoryURLErr)
	}
	return _cFBundleCopyResourcesDirectoryURL(bundle), nil
}

// CFBundleCopyResourcesDirectoryURL returns the location of a bundle’s Resources directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourcesDirectoryURL(_:)
func CFBundleCopyResourcesDirectoryURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopyResourcesDirectoryURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopySharedFrameworksURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopySharedFrameworksURLErr error

func tryCFBundleCopySharedFrameworksURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopySharedFrameworksURL == nil {
		return 0, symbolCallError("CFBundleCopySharedFrameworksURL", "", _cFBundleCopySharedFrameworksURLErr)
	}
	return _cFBundleCopySharedFrameworksURL(bundle), nil
}

// CFBundleCopySharedFrameworksURL returns the location of a bundle’s shared frameworks directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedFrameworksURL(_:)
func CFBundleCopySharedFrameworksURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopySharedFrameworksURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopySharedSupportURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopySharedSupportURLErr error

func tryCFBundleCopySharedSupportURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopySharedSupportURL == nil {
		return 0, symbolCallError("CFBundleCopySharedSupportURL", "", _cFBundleCopySharedSupportURLErr)
	}
	return _cFBundleCopySharedSupportURL(bundle), nil
}

// CFBundleCopySharedSupportURL returns the location of a bundle’s shared support files directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedSupportURL(_:)
func CFBundleCopySharedSupportURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopySharedSupportURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCopySupportFilesDirectoryURL func(bundle CFBundleRef) CFURLRef
var _cFBundleCopySupportFilesDirectoryURLErr error

func tryCFBundleCopySupportFilesDirectoryURL(bundle CFBundleRef) (CFURLRef, error) {
	if _cFBundleCopySupportFilesDirectoryURL == nil {
		return 0, symbolCallError("CFBundleCopySupportFilesDirectoryURL", "", _cFBundleCopySupportFilesDirectoryURLErr)
	}
	return _cFBundleCopySupportFilesDirectoryURL(bundle), nil
}

// CFBundleCopySupportFilesDirectoryURL returns the location of the bundle’s support files directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySupportFilesDirectoryURL(_:)
func CFBundleCopySupportFilesDirectoryURL(bundle CFBundleRef) CFURLRef {
	result, callErr := tryCFBundleCopySupportFilesDirectoryURL(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCreate func(allocator CFAllocatorRef, bundleURL CFURLRef) CFBundleRef
var _cFBundleCreateErr error

func tryCFBundleCreate(allocator CFAllocatorRef, bundleURL CFURLRef) (CFBundleRef, error) {
	if _cFBundleCreate == nil {
		return 0, symbolCallError("CFBundleCreate", "", _cFBundleCreateErr)
	}
	return _cFBundleCreate(allocator, bundleURL), nil
}

// CFBundleCreate creates a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreate(_:_:)
func CFBundleCreate(allocator CFAllocatorRef, bundleURL CFURLRef) CFBundleRef {
	result, callErr := tryCFBundleCreate(allocator, bundleURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleCreateBundlesFromDirectory func(allocator CFAllocatorRef, directoryURL CFURLRef, bundleType CFStringRef) CFArrayRef
var _cFBundleCreateBundlesFromDirectoryErr error

func tryCFBundleCreateBundlesFromDirectory(allocator CFAllocatorRef, directoryURL CFURLRef, bundleType CFStringRef) (CFArrayRef, error) {
	if _cFBundleCreateBundlesFromDirectory == nil {
		return 0, symbolCallError("CFBundleCreateBundlesFromDirectory", "", _cFBundleCreateBundlesFromDirectoryErr)
	}
	return _cFBundleCreateBundlesFromDirectory(allocator, directoryURL, bundleType), nil
}

// CFBundleCreateBundlesFromDirectory searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreateBundlesFromDirectory(_:_:_:)
func CFBundleCreateBundlesFromDirectory(allocator CFAllocatorRef, directoryURL CFURLRef, bundleType CFStringRef) CFArrayRef {
	result, callErr := tryCFBundleCreateBundlesFromDirectory(allocator, directoryURL, bundleType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetAllBundles func() CFArrayRef
var _cFBundleGetAllBundlesErr error

func tryCFBundleGetAllBundles() (CFArrayRef, error) {
	if _cFBundleGetAllBundles == nil {
		return 0, symbolCallError("CFBundleGetAllBundles", "", _cFBundleGetAllBundlesErr)
	}
	return _cFBundleGetAllBundles(), nil
}

// CFBundleGetAllBundles returns an array containing all of the bundles currently open in the application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetAllBundles()
func CFBundleGetAllBundles() CFArrayRef {
	result, callErr := tryCFBundleGetAllBundles()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetBundleWithIdentifier func(bundleID CFStringRef) CFBundleRef
var _cFBundleGetBundleWithIdentifierErr error

func tryCFBundleGetBundleWithIdentifier(bundleID CFStringRef) (CFBundleRef, error) {
	if _cFBundleGetBundleWithIdentifier == nil {
		return 0, symbolCallError("CFBundleGetBundleWithIdentifier", "", _cFBundleGetBundleWithIdentifierErr)
	}
	return _cFBundleGetBundleWithIdentifier(bundleID), nil
}

// CFBundleGetBundleWithIdentifier locate a bundle given its program-defined identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetBundleWithIdentifier(_:)
func CFBundleGetBundleWithIdentifier(bundleID CFStringRef) CFBundleRef {
	result, callErr := tryCFBundleGetBundleWithIdentifier(bundleID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetDataPointerForName func(bundle CFBundleRef, symbolName CFStringRef) unsafe.Pointer
var _cFBundleGetDataPointerForNameErr error

func tryCFBundleGetDataPointerForName(bundle CFBundleRef, symbolName CFStringRef) (unsafe.Pointer, error) {
	if _cFBundleGetDataPointerForName == nil {
		return nil, symbolCallError("CFBundleGetDataPointerForName", "", _cFBundleGetDataPointerForNameErr)
	}
	return _cFBundleGetDataPointerForName(bundle, symbolName), nil
}

// CFBundleGetDataPointerForName returns a data pointer to a symbol of the given name.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointerForName(_:_:)
func CFBundleGetDataPointerForName(bundle CFBundleRef, symbolName CFStringRef) unsafe.Pointer {
	result, callErr := tryCFBundleGetDataPointerForName(bundle, symbolName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetDataPointersForNames func(bundle CFBundleRef, symbolNames CFArrayRef, stbl unsafe.Pointer)
var _cFBundleGetDataPointersForNamesErr error

func tryCFBundleGetDataPointersForNames(bundle CFBundleRef, symbolNames CFArrayRef, stbl unsafe.Pointer) error {
	if _cFBundleGetDataPointersForNames == nil {
		return symbolCallError("CFBundleGetDataPointersForNames", "", _cFBundleGetDataPointersForNamesErr)
	}
	_cFBundleGetDataPointersForNames(bundle, symbolNames, stbl)
	return nil
}

// CFBundleGetDataPointersForNames returns a C array of data pointer to symbols of the given names.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointersForNames(_:_:_:)
func CFBundleGetDataPointersForNames(bundle CFBundleRef, symbolNames CFArrayRef, stbl unsafe.Pointer) {
	if callErr := tryCFBundleGetDataPointersForNames(bundle, symbolNames, stbl); callErr != nil {
		panic(callErr)
	}
}

var _cFBundleGetDevelopmentRegion func(bundle CFBundleRef) CFStringRef
var _cFBundleGetDevelopmentRegionErr error

func tryCFBundleGetDevelopmentRegion(bundle CFBundleRef) (CFStringRef, error) {
	if _cFBundleGetDevelopmentRegion == nil {
		return 0, symbolCallError("CFBundleGetDevelopmentRegion", "", _cFBundleGetDevelopmentRegionErr)
	}
	return _cFBundleGetDevelopmentRegion(bundle), nil
}

// CFBundleGetDevelopmentRegion returns the bundle’s development region from the bundle’s information property list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDevelopmentRegion(_:)
func CFBundleGetDevelopmentRegion(bundle CFBundleRef) CFStringRef {
	result, callErr := tryCFBundleGetDevelopmentRegion(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetFunctionPointerForName func(bundle CFBundleRef, functionName CFStringRef) unsafe.Pointer
var _cFBundleGetFunctionPointerForNameErr error

func tryCFBundleGetFunctionPointerForName(bundle CFBundleRef, functionName CFStringRef) (unsafe.Pointer, error) {
	if _cFBundleGetFunctionPointerForName == nil {
		return nil, symbolCallError("CFBundleGetFunctionPointerForName", "", _cFBundleGetFunctionPointerForNameErr)
	}
	return _cFBundleGetFunctionPointerForName(bundle, functionName), nil
}

// CFBundleGetFunctionPointerForName returns a pointer to a function in a bundle’s executable code using the function name as the search key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointerForName(_:_:)
func CFBundleGetFunctionPointerForName(bundle CFBundleRef, functionName CFStringRef) unsafe.Pointer {
	result, callErr := tryCFBundleGetFunctionPointerForName(bundle, functionName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetFunctionPointersForNames func(bundle CFBundleRef, functionNames CFArrayRef, ftbl unsafe.Pointer)
var _cFBundleGetFunctionPointersForNamesErr error

func tryCFBundleGetFunctionPointersForNames(bundle CFBundleRef, functionNames CFArrayRef, ftbl unsafe.Pointer) error {
	if _cFBundleGetFunctionPointersForNames == nil {
		return symbolCallError("CFBundleGetFunctionPointersForNames", "", _cFBundleGetFunctionPointersForNamesErr)
	}
	_cFBundleGetFunctionPointersForNames(bundle, functionNames, ftbl)
	return nil
}

// CFBundleGetFunctionPointersForNames constructs a function table containing pointers to all of the functions found in a bundle’s main executable code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointersForNames(_:_:_:)
func CFBundleGetFunctionPointersForNames(bundle CFBundleRef, functionNames CFArrayRef, ftbl unsafe.Pointer) {
	if callErr := tryCFBundleGetFunctionPointersForNames(bundle, functionNames, ftbl); callErr != nil {
		panic(callErr)
	}
}

var _cFBundleGetIdentifier func(bundle CFBundleRef) CFStringRef
var _cFBundleGetIdentifierErr error

func tryCFBundleGetIdentifier(bundle CFBundleRef) (CFStringRef, error) {
	if _cFBundleGetIdentifier == nil {
		return 0, symbolCallError("CFBundleGetIdentifier", "", _cFBundleGetIdentifierErr)
	}
	return _cFBundleGetIdentifier(bundle), nil
}

// CFBundleGetIdentifier returns the bundle identifier from a bundle’s information property list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetIdentifier(_:)
func CFBundleGetIdentifier(bundle CFBundleRef) CFStringRef {
	result, callErr := tryCFBundleGetIdentifier(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetInfoDictionary func(bundle CFBundleRef) CFDictionaryRef
var _cFBundleGetInfoDictionaryErr error

func tryCFBundleGetInfoDictionary(bundle CFBundleRef) (CFDictionaryRef, error) {
	if _cFBundleGetInfoDictionary == nil {
		return 0, symbolCallError("CFBundleGetInfoDictionary", "", _cFBundleGetInfoDictionaryErr)
	}
	return _cFBundleGetInfoDictionary(bundle), nil
}

// CFBundleGetInfoDictionary returns a bundle’s information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetInfoDictionary(_:)
func CFBundleGetInfoDictionary(bundle CFBundleRef) CFDictionaryRef {
	result, callErr := tryCFBundleGetInfoDictionary(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetLocalInfoDictionary func(bundle CFBundleRef) CFDictionaryRef
var _cFBundleGetLocalInfoDictionaryErr error

func tryCFBundleGetLocalInfoDictionary(bundle CFBundleRef) (CFDictionaryRef, error) {
	if _cFBundleGetLocalInfoDictionary == nil {
		return 0, symbolCallError("CFBundleGetLocalInfoDictionary", "", _cFBundleGetLocalInfoDictionaryErr)
	}
	return _cFBundleGetLocalInfoDictionary(bundle), nil
}

// CFBundleGetLocalInfoDictionary returns a bundle’s localized information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetLocalInfoDictionary(_:)
func CFBundleGetLocalInfoDictionary(bundle CFBundleRef) CFDictionaryRef {
	result, callErr := tryCFBundleGetLocalInfoDictionary(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetMainBundle func() CFBundleRef
var _cFBundleGetMainBundleErr error

func tryCFBundleGetMainBundle() (CFBundleRef, error) {
	if _cFBundleGetMainBundle == nil {
		return 0, symbolCallError("CFBundleGetMainBundle", "", _cFBundleGetMainBundleErr)
	}
	return _cFBundleGetMainBundle(), nil
}

// CFBundleGetMainBundle returns an application’s main bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetMainBundle()
func CFBundleGetMainBundle() CFBundleRef {
	result, callErr := tryCFBundleGetMainBundle()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetPackageInfo func(bundle CFBundleRef, packageType *uint32, packageCreator *uint32)
var _cFBundleGetPackageInfoErr error

func tryCFBundleGetPackageInfo(bundle CFBundleRef, packageType *uint32, packageCreator *uint32) error {
	if _cFBundleGetPackageInfo == nil {
		return symbolCallError("CFBundleGetPackageInfo", "", _cFBundleGetPackageInfoErr)
	}
	_cFBundleGetPackageInfo(bundle, packageType, packageCreator)
	return nil
}

// CFBundleGetPackageInfo returns a bundle’s package type and creator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfo(_:_:_:)
func CFBundleGetPackageInfo(bundle CFBundleRef, packageType *uint32, packageCreator *uint32) {
	if callErr := tryCFBundleGetPackageInfo(bundle, packageType, packageCreator); callErr != nil {
		panic(callErr)
	}
}

var _cFBundleGetPackageInfoInDirectory func(url CFURLRef, packageType *uint32, packageCreator *uint32) bool
var _cFBundleGetPackageInfoInDirectoryErr error

func tryCFBundleGetPackageInfoInDirectory(url CFURLRef, packageType *uint32, packageCreator *uint32) (bool, error) {
	if _cFBundleGetPackageInfoInDirectory == nil {
		return false, symbolCallError("CFBundleGetPackageInfoInDirectory", "", _cFBundleGetPackageInfoInDirectoryErr)
	}
	return _cFBundleGetPackageInfoInDirectory(url, packageType, packageCreator), nil
}

// CFBundleGetPackageInfoInDirectory returns a bundle’s package type and creator without having to create a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfoInDirectory(_:_:_:)
func CFBundleGetPackageInfoInDirectory(url CFURLRef, packageType *uint32, packageCreator *uint32) bool {
	result, callErr := tryCFBundleGetPackageInfoInDirectory(url, packageType, packageCreator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetPlugIn func(bundle CFBundleRef) CFPlugInRef
var _cFBundleGetPlugInErr error

func tryCFBundleGetPlugIn(bundle CFBundleRef) (CFPlugInRef, error) {
	if _cFBundleGetPlugIn == nil {
		return 0, symbolCallError("CFBundleGetPlugIn", "", _cFBundleGetPlugInErr)
	}
	return _cFBundleGetPlugIn(bundle), nil
}

// CFBundleGetPlugIn returns a bundle’s plug-in.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPlugIn(_:)
func CFBundleGetPlugIn(bundle CFBundleRef) CFPlugInRef {
	result, callErr := tryCFBundleGetPlugIn(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetTypeID func() uint
var _cFBundleGetTypeIDErr error

func tryCFBundleGetTypeID() (uint, error) {
	if _cFBundleGetTypeID == nil {
		return 0, symbolCallError("CFBundleGetTypeID", "", _cFBundleGetTypeIDErr)
	}
	return _cFBundleGetTypeID(), nil
}

// CFBundleGetTypeID returns the type identifier for the CFBundle opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetTypeID()
func CFBundleGetTypeID() uint {
	result, callErr := tryCFBundleGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetValueForInfoDictionaryKey func(bundle CFBundleRef, key CFStringRef) CFTypeRef
var _cFBundleGetValueForInfoDictionaryKeyErr error

func tryCFBundleGetValueForInfoDictionaryKey(bundle CFBundleRef, key CFStringRef) (CFTypeRef, error) {
	if _cFBundleGetValueForInfoDictionaryKey == nil {
		return 0, symbolCallError("CFBundleGetValueForInfoDictionaryKey", "", _cFBundleGetValueForInfoDictionaryKeyErr)
	}
	return _cFBundleGetValueForInfoDictionaryKey(bundle, key), nil
}

// CFBundleGetValueForInfoDictionaryKey returns a value (localized if possible) from a bundle’s information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetValueForInfoDictionaryKey(_:_:)
func CFBundleGetValueForInfoDictionaryKey(bundle CFBundleRef, key CFStringRef) CFTypeRef {
	result, callErr := tryCFBundleGetValueForInfoDictionaryKey(bundle, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleGetVersionNumber func(bundle CFBundleRef) uint32
var _cFBundleGetVersionNumberErr error

func tryCFBundleGetVersionNumber(bundle CFBundleRef) (uint32, error) {
	if _cFBundleGetVersionNumber == nil {
		return 0, symbolCallError("CFBundleGetVersionNumber", "", _cFBundleGetVersionNumberErr)
	}
	return _cFBundleGetVersionNumber(bundle), nil
}

// CFBundleGetVersionNumber returns a bundle’s version number.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetVersionNumber(_:)
func CFBundleGetVersionNumber(bundle CFBundleRef) uint32 {
	result, callErr := tryCFBundleGetVersionNumber(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleIsArchitectureLoadable func(arch int32) bool
var _cFBundleIsArchitectureLoadableErr error

func tryCFBundleIsArchitectureLoadable(arch int32) (bool, error) {
	if _cFBundleIsArchitectureLoadable == nil {
		return false, symbolCallError("CFBundleIsArchitectureLoadable", "11.0", _cFBundleIsArchitectureLoadableErr)
	}
	return _cFBundleIsArchitectureLoadable(arch), nil
}

// CFBundleIsArchitectureLoadable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsArchitectureLoadable(_:)
func CFBundleIsArchitectureLoadable(arch int32) bool {
	result, callErr := tryCFBundleIsArchitectureLoadable(arch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleIsExecutableLoadable func(bundle CFBundleRef) bool
var _cFBundleIsExecutableLoadableErr error

func tryCFBundleIsExecutableLoadable(bundle CFBundleRef) (bool, error) {
	if _cFBundleIsExecutableLoadable == nil {
		return false, symbolCallError("CFBundleIsExecutableLoadable", "11.0", _cFBundleIsExecutableLoadableErr)
	}
	return _cFBundleIsExecutableLoadable(bundle), nil
}

// CFBundleIsExecutableLoadable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadable(_:)
func CFBundleIsExecutableLoadable(bundle CFBundleRef) bool {
	result, callErr := tryCFBundleIsExecutableLoadable(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleIsExecutableLoadableForURL func(url CFURLRef) bool
var _cFBundleIsExecutableLoadableForURLErr error

func tryCFBundleIsExecutableLoadableForURL(url CFURLRef) (bool, error) {
	if _cFBundleIsExecutableLoadableForURL == nil {
		return false, symbolCallError("CFBundleIsExecutableLoadableForURL", "11.0", _cFBundleIsExecutableLoadableForURLErr)
	}
	return _cFBundleIsExecutableLoadableForURL(url), nil
}

// CFBundleIsExecutableLoadableForURL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadableForURL(_:)
func CFBundleIsExecutableLoadableForURL(url CFURLRef) bool {
	result, callErr := tryCFBundleIsExecutableLoadableForURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleIsExecutableLoaded func(bundle CFBundleRef) bool
var _cFBundleIsExecutableLoadedErr error

func tryCFBundleIsExecutableLoaded(bundle CFBundleRef) (bool, error) {
	if _cFBundleIsExecutableLoaded == nil {
		return false, symbolCallError("CFBundleIsExecutableLoaded", "", _cFBundleIsExecutableLoadedErr)
	}
	return _cFBundleIsExecutableLoaded(bundle), nil
}

// CFBundleIsExecutableLoaded obtains information about the load status for a bundle’s main executable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoaded(_:)
func CFBundleIsExecutableLoaded(bundle CFBundleRef) bool {
	result, callErr := tryCFBundleIsExecutableLoaded(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleLoadExecutable func(bundle CFBundleRef) bool
var _cFBundleLoadExecutableErr error

func tryCFBundleLoadExecutable(bundle CFBundleRef) (bool, error) {
	if _cFBundleLoadExecutable == nil {
		return false, symbolCallError("CFBundleLoadExecutable", "", _cFBundleLoadExecutableErr)
	}
	return _cFBundleLoadExecutable(bundle), nil
}

// CFBundleLoadExecutable loads a bundle’s main executable code into memory and dynamically links it into the running application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutable(_:)
func CFBundleLoadExecutable(bundle CFBundleRef) bool {
	result, callErr := tryCFBundleLoadExecutable(bundle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleLoadExecutableAndReturnError func(bundle CFBundleRef, err *CFErrorRef) bool
var _cFBundleLoadExecutableAndReturnErrorErr error

func tryCFBundleLoadExecutableAndReturnError(bundle CFBundleRef, err *CFErrorRef) (bool, error) {
	if _cFBundleLoadExecutableAndReturnError == nil {
		return false, symbolCallError("CFBundleLoadExecutableAndReturnError", "10.5", _cFBundleLoadExecutableAndReturnErrorErr)
	}
	return _cFBundleLoadExecutableAndReturnError(bundle, err), nil
}

// CFBundleLoadExecutableAndReturnError returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutableAndReturnError(_:_:)
func CFBundleLoadExecutableAndReturnError(bundle CFBundleRef, err *CFErrorRef) bool {
	result, callErr := tryCFBundleLoadExecutableAndReturnError(bundle, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundlePreflightExecutable func(bundle CFBundleRef, err *CFErrorRef) bool
var _cFBundlePreflightExecutableErr error

func tryCFBundlePreflightExecutable(bundle CFBundleRef, err *CFErrorRef) (bool, error) {
	if _cFBundlePreflightExecutable == nil {
		return false, symbolCallError("CFBundlePreflightExecutable", "10.5", _cFBundlePreflightExecutableErr)
	}
	return _cFBundlePreflightExecutable(bundle, err), nil
}

// CFBundlePreflightExecutable returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundlePreflightExecutable(_:_:)
func CFBundlePreflightExecutable(bundle CFBundleRef, err *CFErrorRef) bool {
	result, callErr := tryCFBundlePreflightExecutable(bundle, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFBundleUnloadExecutable func(bundle CFBundleRef)
var _cFBundleUnloadExecutableErr error

func tryCFBundleUnloadExecutable(bundle CFBundleRef) error {
	if _cFBundleUnloadExecutable == nil {
		return symbolCallError("CFBundleUnloadExecutable", "", _cFBundleUnloadExecutableErr)
	}
	_cFBundleUnloadExecutable(bundle)
	return nil
}

// CFBundleUnloadExecutable unloads the main executable for the specified bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleUnloadExecutable(_:)
func CFBundleUnloadExecutable(bundle CFBundleRef) {
	if callErr := tryCFBundleUnloadExecutable(bundle); callErr != nil {
		panic(callErr)
	}
}

var _cFCalendarAddComponents func(calendar CFCalendarRef, at *CFAbsoluteTime, options uint64, componentDesc string) bool
var _cFCalendarAddComponentsErr error

func tryCFCalendarAddComponents(calendar CFCalendarRef, at *CFAbsoluteTime, options uint64, componentDesc string) (bool, error) {
	if _cFCalendarAddComponents == nil {
		return false, symbolCallError("CFCalendarAddComponents", "", _cFCalendarAddComponentsErr)
	}
	return _cFCalendarAddComponents(calendar, at, options, componentDesc), nil
}

// CFCalendarAddComponents computes the absolute time when specified components are added to a given absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarAddComponents
func CFCalendarAddComponents(calendar CFCalendarRef, at *CFAbsoluteTime, options uint64, componentDesc string) bool {
	result, callErr := tryCFCalendarAddComponents(calendar, at, options, componentDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarComposeAbsoluteTime func(calendar CFCalendarRef, at *CFAbsoluteTime, componentDesc string) bool
var _cFCalendarComposeAbsoluteTimeErr error

func tryCFCalendarComposeAbsoluteTime(calendar CFCalendarRef, at *CFAbsoluteTime, componentDesc string) (bool, error) {
	if _cFCalendarComposeAbsoluteTime == nil {
		return false, symbolCallError("CFCalendarComposeAbsoluteTime", "", _cFCalendarComposeAbsoluteTimeErr)
	}
	return _cFCalendarComposeAbsoluteTime(calendar, at, componentDesc), nil
}

// CFCalendarComposeAbsoluteTime computes the absolute time from components in a description string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarComposeAbsoluteTime
func CFCalendarComposeAbsoluteTime(calendar CFCalendarRef, at *CFAbsoluteTime, componentDesc string) bool {
	result, callErr := tryCFCalendarComposeAbsoluteTime(calendar, at, componentDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarCopyCurrent func() CFCalendarRef
var _cFCalendarCopyCurrentErr error

func tryCFCalendarCopyCurrent() (CFCalendarRef, error) {
	if _cFCalendarCopyCurrent == nil {
		return 0, symbolCallError("CFCalendarCopyCurrent", "", _cFCalendarCopyCurrentErr)
	}
	return _cFCalendarCopyCurrent(), nil
}

// CFCalendarCopyCurrent returns a copy of the logical calendar for the current user.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyCurrent()
func CFCalendarCopyCurrent() CFCalendarRef {
	result, callErr := tryCFCalendarCopyCurrent()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarCopyLocale func(calendar CFCalendarRef) CFLocaleRef
var _cFCalendarCopyLocaleErr error

func tryCFCalendarCopyLocale(calendar CFCalendarRef) (CFLocaleRef, error) {
	if _cFCalendarCopyLocale == nil {
		return 0, symbolCallError("CFCalendarCopyLocale", "", _cFCalendarCopyLocaleErr)
	}
	return _cFCalendarCopyLocale(calendar), nil
}

// CFCalendarCopyLocale returns a locale object for a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyLocale(_:)
func CFCalendarCopyLocale(calendar CFCalendarRef) CFLocaleRef {
	result, callErr := tryCFCalendarCopyLocale(calendar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarCopyTimeZone func(calendar CFCalendarRef) CFTimeZoneRef
var _cFCalendarCopyTimeZoneErr error

func tryCFCalendarCopyTimeZone(calendar CFCalendarRef) (CFTimeZoneRef, error) {
	if _cFCalendarCopyTimeZone == nil {
		return 0, symbolCallError("CFCalendarCopyTimeZone", "", _cFCalendarCopyTimeZoneErr)
	}
	return _cFCalendarCopyTimeZone(calendar), nil
}

// CFCalendarCopyTimeZone returns a time zone object for a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyTimeZone(_:)
func CFCalendarCopyTimeZone(calendar CFCalendarRef) CFTimeZoneRef {
	result, callErr := tryCFCalendarCopyTimeZone(calendar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarCreateWithIdentifier func(allocator CFAllocatorRef, identifier CFCalendarIdentifier) CFCalendarRef
var _cFCalendarCreateWithIdentifierErr error

func tryCFCalendarCreateWithIdentifier(allocator CFAllocatorRef, identifier CFCalendarIdentifier) (CFCalendarRef, error) {
	if _cFCalendarCreateWithIdentifier == nil {
		return 0, symbolCallError("CFCalendarCreateWithIdentifier", "", _cFCalendarCreateWithIdentifierErr)
	}
	return _cFCalendarCreateWithIdentifier(allocator, identifier), nil
}

// CFCalendarCreateWithIdentifier returns a calendar object for the calendar identified by a calendar identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCreateWithIdentifier(_:_:)
func CFCalendarCreateWithIdentifier(allocator CFAllocatorRef, identifier CFCalendarIdentifier) CFCalendarRef {
	result, callErr := tryCFCalendarCreateWithIdentifier(allocator, identifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarDecomposeAbsoluteTime func(calendar CFCalendarRef, at CFAbsoluteTime, componentDesc string) bool
var _cFCalendarDecomposeAbsoluteTimeErr error

func tryCFCalendarDecomposeAbsoluteTime(calendar CFCalendarRef, at CFAbsoluteTime, componentDesc string) (bool, error) {
	if _cFCalendarDecomposeAbsoluteTime == nil {
		return false, symbolCallError("CFCalendarDecomposeAbsoluteTime", "", _cFCalendarDecomposeAbsoluteTimeErr)
	}
	return _cFCalendarDecomposeAbsoluteTime(calendar, at, componentDesc), nil
}

// CFCalendarDecomposeAbsoluteTime computes the components which are indicated by the componentDesc description string for the given absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarDecomposeAbsoluteTime
func CFCalendarDecomposeAbsoluteTime(calendar CFCalendarRef, at CFAbsoluteTime, componentDesc string) bool {
	result, callErr := tryCFCalendarDecomposeAbsoluteTime(calendar, at, componentDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetComponentDifference func(calendar CFCalendarRef, startingAT CFAbsoluteTime, resultAT CFAbsoluteTime, options uint64, componentDesc string) bool
var _cFCalendarGetComponentDifferenceErr error

func tryCFCalendarGetComponentDifference(calendar CFCalendarRef, startingAT CFAbsoluteTime, resultAT CFAbsoluteTime, options uint64, componentDesc string) (bool, error) {
	if _cFCalendarGetComponentDifference == nil {
		return false, symbolCallError("CFCalendarGetComponentDifference", "", _cFCalendarGetComponentDifferenceErr)
	}
	return _cFCalendarGetComponentDifference(calendar, startingAT, resultAT, options, componentDesc), nil
}

// CFCalendarGetComponentDifference computes the difference between the two absolute times, in terms of specified calendrical components.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetComponentDifference
func CFCalendarGetComponentDifference(calendar CFCalendarRef, startingAT CFAbsoluteTime, resultAT CFAbsoluteTime, options uint64, componentDesc string) bool {
	result, callErr := tryCFCalendarGetComponentDifference(calendar, startingAT, resultAT, options, componentDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetFirstWeekday func(calendar CFCalendarRef) int
var _cFCalendarGetFirstWeekdayErr error

func tryCFCalendarGetFirstWeekday(calendar CFCalendarRef) (int, error) {
	if _cFCalendarGetFirstWeekday == nil {
		return 0, symbolCallError("CFCalendarGetFirstWeekday", "", _cFCalendarGetFirstWeekdayErr)
	}
	return _cFCalendarGetFirstWeekday(calendar), nil
}

// CFCalendarGetFirstWeekday returns the index of first weekday for a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetFirstWeekday(_:)
func CFCalendarGetFirstWeekday(calendar CFCalendarRef) int {
	result, callErr := tryCFCalendarGetFirstWeekday(calendar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetIdentifier func(calendar CFCalendarRef) CFCalendarIdentifier
var _cFCalendarGetIdentifierErr error

func tryCFCalendarGetIdentifier(calendar CFCalendarRef) (CFCalendarIdentifier, error) {
	if _cFCalendarGetIdentifier == nil {
		return *new(CFCalendarIdentifier), symbolCallError("CFCalendarGetIdentifier", "", _cFCalendarGetIdentifierErr)
	}
	return _cFCalendarGetIdentifier(calendar), nil
}

// CFCalendarGetIdentifier returns the given calendar’s identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetIdentifier(_:)
func CFCalendarGetIdentifier(calendar CFCalendarRef) CFCalendarIdentifier {
	result, callErr := tryCFCalendarGetIdentifier(calendar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetMaximumRangeOfUnit func(calendar CFCalendarRef, unit CFCalendarUnit) CFRange
var _cFCalendarGetMaximumRangeOfUnitErr error

func tryCFCalendarGetMaximumRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit) (CFRange, error) {
	if _cFCalendarGetMaximumRangeOfUnit == nil {
		return CFRange{}, symbolCallError("CFCalendarGetMaximumRangeOfUnit", "", _cFCalendarGetMaximumRangeOfUnitErr)
	}
	return _cFCalendarGetMaximumRangeOfUnit(calendar, unit), nil
}

// CFCalendarGetMaximumRangeOfUnit returns the maximum range limits of the values that a specified unit can take on in a given calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMaximumRangeOfUnit(_:_:)
func CFCalendarGetMaximumRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit) CFRange {
	result, callErr := tryCFCalendarGetMaximumRangeOfUnit(calendar, unit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetMinimumDaysInFirstWeek func(calendar CFCalendarRef) int
var _cFCalendarGetMinimumDaysInFirstWeekErr error

func tryCFCalendarGetMinimumDaysInFirstWeek(calendar CFCalendarRef) (int, error) {
	if _cFCalendarGetMinimumDaysInFirstWeek == nil {
		return 0, symbolCallError("CFCalendarGetMinimumDaysInFirstWeek", "", _cFCalendarGetMinimumDaysInFirstWeekErr)
	}
	return _cFCalendarGetMinimumDaysInFirstWeek(calendar), nil
}

// CFCalendarGetMinimumDaysInFirstWeek returns the minimum number of days in the first week of a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumDaysInFirstWeek(_:)
func CFCalendarGetMinimumDaysInFirstWeek(calendar CFCalendarRef) int {
	result, callErr := tryCFCalendarGetMinimumDaysInFirstWeek(calendar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetMinimumRangeOfUnit func(calendar CFCalendarRef, unit CFCalendarUnit) CFRange
var _cFCalendarGetMinimumRangeOfUnitErr error

func tryCFCalendarGetMinimumRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit) (CFRange, error) {
	if _cFCalendarGetMinimumRangeOfUnit == nil {
		return CFRange{}, symbolCallError("CFCalendarGetMinimumRangeOfUnit", "", _cFCalendarGetMinimumRangeOfUnitErr)
	}
	return _cFCalendarGetMinimumRangeOfUnit(calendar, unit), nil
}

// CFCalendarGetMinimumRangeOfUnit returns the minimum range limits of the values that a specified unit can take on in a given calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumRangeOfUnit(_:_:)
func CFCalendarGetMinimumRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit) CFRange {
	result, callErr := tryCFCalendarGetMinimumRangeOfUnit(calendar, unit)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetOrdinalityOfUnit func(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) int
var _cFCalendarGetOrdinalityOfUnitErr error

func tryCFCalendarGetOrdinalityOfUnit(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) (int, error) {
	if _cFCalendarGetOrdinalityOfUnit == nil {
		return 0, symbolCallError("CFCalendarGetOrdinalityOfUnit", "", _cFCalendarGetOrdinalityOfUnitErr)
	}
	return _cFCalendarGetOrdinalityOfUnit(calendar, smallerUnit, biggerUnit, at), nil
}

// CFCalendarGetOrdinalityOfUnit returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetOrdinalityOfUnit(_:_:_:_:)
func CFCalendarGetOrdinalityOfUnit(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) int {
	result, callErr := tryCFCalendarGetOrdinalityOfUnit(calendar, smallerUnit, biggerUnit, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetRangeOfUnit func(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) CFRange
var _cFCalendarGetRangeOfUnitErr error

func tryCFCalendarGetRangeOfUnit(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) (CFRange, error) {
	if _cFCalendarGetRangeOfUnit == nil {
		return CFRange{}, symbolCallError("CFCalendarGetRangeOfUnit", "", _cFCalendarGetRangeOfUnitErr)
	}
	return _cFCalendarGetRangeOfUnit(calendar, smallerUnit, biggerUnit, at), nil
}

// CFCalendarGetRangeOfUnit returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetRangeOfUnit(_:_:_:_:)
func CFCalendarGetRangeOfUnit(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) CFRange {
	result, callErr := tryCFCalendarGetRangeOfUnit(calendar, smallerUnit, biggerUnit, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetTimeRangeOfUnit func(calendar CFCalendarRef, unit CFCalendarUnit, at CFAbsoluteTime, startp *CFAbsoluteTime, tip *float64) bool
var _cFCalendarGetTimeRangeOfUnitErr error

func tryCFCalendarGetTimeRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit, at CFAbsoluteTime, startp *CFAbsoluteTime, tip *float64) (bool, error) {
	if _cFCalendarGetTimeRangeOfUnit == nil {
		return false, symbolCallError("CFCalendarGetTimeRangeOfUnit", "10.5", _cFCalendarGetTimeRangeOfUnitErr)
	}
	return _cFCalendarGetTimeRangeOfUnit(calendar, unit, at, startp, tip), nil
}

// CFCalendarGetTimeRangeOfUnit returns by reference the start time and duration of a given calendar unit that contains a given absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTimeRangeOfUnit(_:_:_:_:_:)
func CFCalendarGetTimeRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit, at CFAbsoluteTime, startp *CFAbsoluteTime, tip *float64) bool {
	result, callErr := tryCFCalendarGetTimeRangeOfUnit(calendar, unit, at, startp, tip)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarGetTypeID func() uint
var _cFCalendarGetTypeIDErr error

func tryCFCalendarGetTypeID() (uint, error) {
	if _cFCalendarGetTypeID == nil {
		return 0, symbolCallError("CFCalendarGetTypeID", "", _cFCalendarGetTypeIDErr)
	}
	return _cFCalendarGetTypeID(), nil
}

// CFCalendarGetTypeID returns the type identifier for the CFCalendar opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTypeID()
func CFCalendarGetTypeID() uint {
	result, callErr := tryCFCalendarGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCalendarSetFirstWeekday func(calendar CFCalendarRef, wkdy int)
var _cFCalendarSetFirstWeekdayErr error

func tryCFCalendarSetFirstWeekday(calendar CFCalendarRef, wkdy int) error {
	if _cFCalendarSetFirstWeekday == nil {
		return symbolCallError("CFCalendarSetFirstWeekday", "", _cFCalendarSetFirstWeekdayErr)
	}
	_cFCalendarSetFirstWeekday(calendar, wkdy)
	return nil
}

// CFCalendarSetFirstWeekday sets the first weekday for a calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetFirstWeekday(_:_:)
func CFCalendarSetFirstWeekday(calendar CFCalendarRef, wkdy int) {
	if callErr := tryCFCalendarSetFirstWeekday(calendar, wkdy); callErr != nil {
		panic(callErr)
	}
}

var _cFCalendarSetLocale func(calendar CFCalendarRef, locale CFLocaleRef)
var _cFCalendarSetLocaleErr error

func tryCFCalendarSetLocale(calendar CFCalendarRef, locale CFLocaleRef) error {
	if _cFCalendarSetLocale == nil {
		return symbolCallError("CFCalendarSetLocale", "", _cFCalendarSetLocaleErr)
	}
	_cFCalendarSetLocale(calendar, locale)
	return nil
}

// CFCalendarSetLocale sets the locale for a calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetLocale(_:_:)
func CFCalendarSetLocale(calendar CFCalendarRef, locale CFLocaleRef) {
	if callErr := tryCFCalendarSetLocale(calendar, locale); callErr != nil {
		panic(callErr)
	}
}

var _cFCalendarSetMinimumDaysInFirstWeek func(calendar CFCalendarRef, mwd int)
var _cFCalendarSetMinimumDaysInFirstWeekErr error

func tryCFCalendarSetMinimumDaysInFirstWeek(calendar CFCalendarRef, mwd int) error {
	if _cFCalendarSetMinimumDaysInFirstWeek == nil {
		return symbolCallError("CFCalendarSetMinimumDaysInFirstWeek", "", _cFCalendarSetMinimumDaysInFirstWeekErr)
	}
	_cFCalendarSetMinimumDaysInFirstWeek(calendar, mwd)
	return nil
}

// CFCalendarSetMinimumDaysInFirstWeek sets the minimum number of days in the first week of a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetMinimumDaysInFirstWeek(_:_:)
func CFCalendarSetMinimumDaysInFirstWeek(calendar CFCalendarRef, mwd int) {
	if callErr := tryCFCalendarSetMinimumDaysInFirstWeek(calendar, mwd); callErr != nil {
		panic(callErr)
	}
}

var _cFCalendarSetTimeZone func(calendar CFCalendarRef, tz CFTimeZoneRef)
var _cFCalendarSetTimeZoneErr error

func tryCFCalendarSetTimeZone(calendar CFCalendarRef, tz CFTimeZoneRef) error {
	if _cFCalendarSetTimeZone == nil {
		return symbolCallError("CFCalendarSetTimeZone", "", _cFCalendarSetTimeZoneErr)
	}
	_cFCalendarSetTimeZone(calendar, tz)
	return nil
}

// CFCalendarSetTimeZone sets the time zone for a calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetTimeZone(_:_:)
func CFCalendarSetTimeZone(calendar CFCalendarRef, tz CFTimeZoneRef) {
	if callErr := tryCFCalendarSetTimeZone(calendar, tz); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetAddCharactersInRange func(theSet CFMutableCharacterSetRef, theRange CFRange)
var _cFCharacterSetAddCharactersInRangeErr error

func tryCFCharacterSetAddCharactersInRange(theSet CFMutableCharacterSetRef, theRange CFRange) error {
	if _cFCharacterSetAddCharactersInRange == nil {
		return symbolCallError("CFCharacterSetAddCharactersInRange", "", _cFCharacterSetAddCharactersInRangeErr)
	}
	_cFCharacterSetAddCharactersInRange(theSet, theRange)
	return nil
}

// CFCharacterSetAddCharactersInRange adds a given range to a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInRange(_:_:)
func CFCharacterSetAddCharactersInRange(theSet CFMutableCharacterSetRef, theRange CFRange) {
	if callErr := tryCFCharacterSetAddCharactersInRange(theSet, theRange); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetAddCharactersInString func(theSet CFMutableCharacterSetRef, theString CFStringRef)
var _cFCharacterSetAddCharactersInStringErr error

func tryCFCharacterSetAddCharactersInString(theSet CFMutableCharacterSetRef, theString CFStringRef) error {
	if _cFCharacterSetAddCharactersInString == nil {
		return symbolCallError("CFCharacterSetAddCharactersInString", "", _cFCharacterSetAddCharactersInStringErr)
	}
	_cFCharacterSetAddCharactersInString(theSet, theString)
	return nil
}

// CFCharacterSetAddCharactersInString adds the characters in a given string to a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInString(_:_:)
func CFCharacterSetAddCharactersInString(theSet CFMutableCharacterSetRef, theString CFStringRef) {
	if callErr := tryCFCharacterSetAddCharactersInString(theSet, theString); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetCreateBitmapRepresentation func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFDataRef
var _cFCharacterSetCreateBitmapRepresentationErr error

func tryCFCharacterSetCreateBitmapRepresentation(alloc CFAllocatorRef, theSet CFCharacterSetRef) (CFDataRef, error) {
	if _cFCharacterSetCreateBitmapRepresentation == nil {
		return 0, symbolCallError("CFCharacterSetCreateBitmapRepresentation", "", _cFCharacterSetCreateBitmapRepresentationErr)
	}
	return _cFCharacterSetCreateBitmapRepresentation(alloc, theSet), nil
}

// CFCharacterSetCreateBitmapRepresentation creates a new immutable data with the bitmap representation from the given character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateBitmapRepresentation(_:_:)
func CFCharacterSetCreateBitmapRepresentation(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFDataRef {
	result, callErr := tryCFCharacterSetCreateBitmapRepresentation(alloc, theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateCopy func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef
var _cFCharacterSetCreateCopyErr error

func tryCFCharacterSetCreateCopy(alloc CFAllocatorRef, theSet CFCharacterSetRef) (CFCharacterSetRef, error) {
	if _cFCharacterSetCreateCopy == nil {
		return 0, symbolCallError("CFCharacterSetCreateCopy", "", _cFCharacterSetCreateCopyErr)
	}
	return _cFCharacterSetCreateCopy(alloc, theSet), nil
}

// CFCharacterSetCreateCopy creates a new character set with the values from a given character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateCopy(_:_:)
func CFCharacterSetCreateCopy(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateCopy(alloc, theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateInvertedSet func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef
var _cFCharacterSetCreateInvertedSetErr error

func tryCFCharacterSetCreateInvertedSet(alloc CFAllocatorRef, theSet CFCharacterSetRef) (CFCharacterSetRef, error) {
	if _cFCharacterSetCreateInvertedSet == nil {
		return 0, symbolCallError("CFCharacterSetCreateInvertedSet", "", _cFCharacterSetCreateInvertedSetErr)
	}
	return _cFCharacterSetCreateInvertedSet(alloc, theSet), nil
}

// CFCharacterSetCreateInvertedSet creates a new immutable character set that is the invert of the specified character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateInvertedSet(_:_:)
func CFCharacterSetCreateInvertedSet(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateInvertedSet(alloc, theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateMutable func(alloc CFAllocatorRef) CFMutableCharacterSetRef
var _cFCharacterSetCreateMutableErr error

func tryCFCharacterSetCreateMutable(alloc CFAllocatorRef) (CFMutableCharacterSetRef, error) {
	if _cFCharacterSetCreateMutable == nil {
		return 0, symbolCallError("CFCharacterSetCreateMutable", "", _cFCharacterSetCreateMutableErr)
	}
	return _cFCharacterSetCreateMutable(alloc), nil
}

// CFCharacterSetCreateMutable creates a new empty mutable character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutable(_:)
func CFCharacterSetCreateMutable(alloc CFAllocatorRef) CFMutableCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateMutable(alloc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateMutableCopy func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFMutableCharacterSetRef
var _cFCharacterSetCreateMutableCopyErr error

func tryCFCharacterSetCreateMutableCopy(alloc CFAllocatorRef, theSet CFCharacterSetRef) (CFMutableCharacterSetRef, error) {
	if _cFCharacterSetCreateMutableCopy == nil {
		return 0, symbolCallError("CFCharacterSetCreateMutableCopy", "", _cFCharacterSetCreateMutableCopyErr)
	}
	return _cFCharacterSetCreateMutableCopy(alloc, theSet), nil
}

// CFCharacterSetCreateMutableCopy creates a new mutable character set with the values from another character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutableCopy(_:_:)
func CFCharacterSetCreateMutableCopy(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFMutableCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateMutableCopy(alloc, theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateWithBitmapRepresentation func(alloc CFAllocatorRef, theData CFDataRef) CFCharacterSetRef
var _cFCharacterSetCreateWithBitmapRepresentationErr error

func tryCFCharacterSetCreateWithBitmapRepresentation(alloc CFAllocatorRef, theData CFDataRef) (CFCharacterSetRef, error) {
	if _cFCharacterSetCreateWithBitmapRepresentation == nil {
		return 0, symbolCallError("CFCharacterSetCreateWithBitmapRepresentation", "", _cFCharacterSetCreateWithBitmapRepresentationErr)
	}
	return _cFCharacterSetCreateWithBitmapRepresentation(alloc, theData), nil
}

// CFCharacterSetCreateWithBitmapRepresentation creates a new immutable character set with the bitmap representation specified by given data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithBitmapRepresentation(_:_:)
func CFCharacterSetCreateWithBitmapRepresentation(alloc CFAllocatorRef, theData CFDataRef) CFCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateWithBitmapRepresentation(alloc, theData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateWithCharactersInRange func(alloc CFAllocatorRef, theRange CFRange) CFCharacterSetRef
var _cFCharacterSetCreateWithCharactersInRangeErr error

func tryCFCharacterSetCreateWithCharactersInRange(alloc CFAllocatorRef, theRange CFRange) (CFCharacterSetRef, error) {
	if _cFCharacterSetCreateWithCharactersInRange == nil {
		return 0, symbolCallError("CFCharacterSetCreateWithCharactersInRange", "", _cFCharacterSetCreateWithCharactersInRangeErr)
	}
	return _cFCharacterSetCreateWithCharactersInRange(alloc, theRange), nil
}

// CFCharacterSetCreateWithCharactersInRange creates a new character set with the values from the given range of Unicode characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInRange(_:_:)
func CFCharacterSetCreateWithCharactersInRange(alloc CFAllocatorRef, theRange CFRange) CFCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateWithCharactersInRange(alloc, theRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetCreateWithCharactersInString func(alloc CFAllocatorRef, theString CFStringRef) CFCharacterSetRef
var _cFCharacterSetCreateWithCharactersInStringErr error

func tryCFCharacterSetCreateWithCharactersInString(alloc CFAllocatorRef, theString CFStringRef) (CFCharacterSetRef, error) {
	if _cFCharacterSetCreateWithCharactersInString == nil {
		return 0, symbolCallError("CFCharacterSetCreateWithCharactersInString", "", _cFCharacterSetCreateWithCharactersInStringErr)
	}
	return _cFCharacterSetCreateWithCharactersInString(alloc, theString), nil
}

// CFCharacterSetCreateWithCharactersInString creates a new character set with the values in the given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInString(_:_:)
func CFCharacterSetCreateWithCharactersInString(alloc CFAllocatorRef, theString CFStringRef) CFCharacterSetRef {
	result, callErr := tryCFCharacterSetCreateWithCharactersInString(alloc, theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetGetPredefined func(theSetIdentifier CFCharacterSetPredefinedSet) CFCharacterSetRef
var _cFCharacterSetGetPredefinedErr error

func tryCFCharacterSetGetPredefined(theSetIdentifier CFCharacterSetPredefinedSet) (CFCharacterSetRef, error) {
	if _cFCharacterSetGetPredefined == nil {
		return 0, symbolCallError("CFCharacterSetGetPredefined", "", _cFCharacterSetGetPredefinedErr)
	}
	return _cFCharacterSetGetPredefined(theSetIdentifier), nil
}

// CFCharacterSetGetPredefined returns a predefined character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetPredefined(_:)
func CFCharacterSetGetPredefined(theSetIdentifier CFCharacterSetPredefinedSet) CFCharacterSetRef {
	result, callErr := tryCFCharacterSetGetPredefined(theSetIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetGetTypeID func() uint
var _cFCharacterSetGetTypeIDErr error

func tryCFCharacterSetGetTypeID() (uint, error) {
	if _cFCharacterSetGetTypeID == nil {
		return 0, symbolCallError("CFCharacterSetGetTypeID", "", _cFCharacterSetGetTypeIDErr)
	}
	return _cFCharacterSetGetTypeID(), nil
}

// CFCharacterSetGetTypeID returns the type identifier of the CFCharacterSet opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetTypeID()
func CFCharacterSetGetTypeID() uint {
	result, callErr := tryCFCharacterSetGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetHasMemberInPlane func(theSet CFCharacterSetRef, thePlane int) bool
var _cFCharacterSetHasMemberInPlaneErr error

func tryCFCharacterSetHasMemberInPlane(theSet CFCharacterSetRef, thePlane int) (bool, error) {
	if _cFCharacterSetHasMemberInPlane == nil {
		return false, symbolCallError("CFCharacterSetHasMemberInPlane", "", _cFCharacterSetHasMemberInPlaneErr)
	}
	return _cFCharacterSetHasMemberInPlane(theSet, thePlane), nil
}

// CFCharacterSetHasMemberInPlane reports whether or not a character set contains at least one member character in the specified plane.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetHasMemberInPlane(_:_:)
func CFCharacterSetHasMemberInPlane(theSet CFCharacterSetRef, thePlane int) bool {
	result, callErr := tryCFCharacterSetHasMemberInPlane(theSet, thePlane)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetIntersect func(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef)
var _cFCharacterSetIntersectErr error

func tryCFCharacterSetIntersect(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef) error {
	if _cFCharacterSetIntersect == nil {
		return symbolCallError("CFCharacterSetIntersect", "", _cFCharacterSetIntersectErr)
	}
	_cFCharacterSetIntersect(theSet, theOtherSet)
	return nil
}

// CFCharacterSetIntersect forms an intersection of two character sets.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIntersect(_:_:)
func CFCharacterSetIntersect(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef) {
	if callErr := tryCFCharacterSetIntersect(theSet, theOtherSet); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetInvert func(theSet CFMutableCharacterSetRef)
var _cFCharacterSetInvertErr error

func tryCFCharacterSetInvert(theSet CFMutableCharacterSetRef) error {
	if _cFCharacterSetInvert == nil {
		return symbolCallError("CFCharacterSetInvert", "", _cFCharacterSetInvertErr)
	}
	_cFCharacterSetInvert(theSet)
	return nil
}

// CFCharacterSetInvert inverts the content of a given character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetInvert(_:)
func CFCharacterSetInvert(theSet CFMutableCharacterSetRef) {
	if callErr := tryCFCharacterSetInvert(theSet); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetIsCharacterMember func(theSet CFCharacterSetRef, theChar uint16) bool
var _cFCharacterSetIsCharacterMemberErr error

func tryCFCharacterSetIsCharacterMember(theSet CFCharacterSetRef, theChar uint16) (bool, error) {
	if _cFCharacterSetIsCharacterMember == nil {
		return false, symbolCallError("CFCharacterSetIsCharacterMember", "", _cFCharacterSetIsCharacterMemberErr)
	}
	return _cFCharacterSetIsCharacterMember(theSet, theChar), nil
}

// CFCharacterSetIsCharacterMember reports whether or not a given Unicode character is in a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsCharacterMember(_:_:)
func CFCharacterSetIsCharacterMember(theSet CFCharacterSetRef, theChar uint16) bool {
	result, callErr := tryCFCharacterSetIsCharacterMember(theSet, theChar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetIsLongCharacterMember func(theSet CFCharacterSetRef, theChar uint32) bool
var _cFCharacterSetIsLongCharacterMemberErr error

func tryCFCharacterSetIsLongCharacterMember(theSet CFCharacterSetRef, theChar uint32) (bool, error) {
	if _cFCharacterSetIsLongCharacterMember == nil {
		return false, symbolCallError("CFCharacterSetIsLongCharacterMember", "", _cFCharacterSetIsLongCharacterMemberErr)
	}
	return _cFCharacterSetIsLongCharacterMember(theSet, theChar), nil
}

// CFCharacterSetIsLongCharacterMember reports whether or not a given UTF-32 character is in a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsLongCharacterMember(_:_:)
func CFCharacterSetIsLongCharacterMember(theSet CFCharacterSetRef, theChar uint32) bool {
	result, callErr := tryCFCharacterSetIsLongCharacterMember(theSet, theChar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetIsSupersetOfSet func(theSet CFCharacterSetRef, theOtherset CFCharacterSetRef) bool
var _cFCharacterSetIsSupersetOfSetErr error

func tryCFCharacterSetIsSupersetOfSet(theSet CFCharacterSetRef, theOtherset CFCharacterSetRef) (bool, error) {
	if _cFCharacterSetIsSupersetOfSet == nil {
		return false, symbolCallError("CFCharacterSetIsSupersetOfSet", "", _cFCharacterSetIsSupersetOfSetErr)
	}
	return _cFCharacterSetIsSupersetOfSet(theSet, theOtherset), nil
}

// CFCharacterSetIsSupersetOfSet reports whether or not a character set is a superset of another set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsSupersetOfSet(_:_:)
func CFCharacterSetIsSupersetOfSet(theSet CFCharacterSetRef, theOtherset CFCharacterSetRef) bool {
	result, callErr := tryCFCharacterSetIsSupersetOfSet(theSet, theOtherset)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCharacterSetRemoveCharactersInRange func(theSet CFMutableCharacterSetRef, theRange CFRange)
var _cFCharacterSetRemoveCharactersInRangeErr error

func tryCFCharacterSetRemoveCharactersInRange(theSet CFMutableCharacterSetRef, theRange CFRange) error {
	if _cFCharacterSetRemoveCharactersInRange == nil {
		return symbolCallError("CFCharacterSetRemoveCharactersInRange", "", _cFCharacterSetRemoveCharactersInRangeErr)
	}
	_cFCharacterSetRemoveCharactersInRange(theSet, theRange)
	return nil
}

// CFCharacterSetRemoveCharactersInRange removes a given range of Unicode characters from a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInRange(_:_:)
func CFCharacterSetRemoveCharactersInRange(theSet CFMutableCharacterSetRef, theRange CFRange) {
	if callErr := tryCFCharacterSetRemoveCharactersInRange(theSet, theRange); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetRemoveCharactersInString func(theSet CFMutableCharacterSetRef, theString CFStringRef)
var _cFCharacterSetRemoveCharactersInStringErr error

func tryCFCharacterSetRemoveCharactersInString(theSet CFMutableCharacterSetRef, theString CFStringRef) error {
	if _cFCharacterSetRemoveCharactersInString == nil {
		return symbolCallError("CFCharacterSetRemoveCharactersInString", "", _cFCharacterSetRemoveCharactersInStringErr)
	}
	_cFCharacterSetRemoveCharactersInString(theSet, theString)
	return nil
}

// CFCharacterSetRemoveCharactersInString removes the characters in a given string from a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInString(_:_:)
func CFCharacterSetRemoveCharactersInString(theSet CFMutableCharacterSetRef, theString CFStringRef) {
	if callErr := tryCFCharacterSetRemoveCharactersInString(theSet, theString); callErr != nil {
		panic(callErr)
	}
}

var _cFCharacterSetUnion func(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef)
var _cFCharacterSetUnionErr error

func tryCFCharacterSetUnion(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef) error {
	if _cFCharacterSetUnion == nil {
		return symbolCallError("CFCharacterSetUnion", "", _cFCharacterSetUnionErr)
	}
	_cFCharacterSetUnion(theSet, theOtherSet)
	return nil
}

// CFCharacterSetUnion forms the union of two character sets.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetUnion(_:_:)
func CFCharacterSetUnion(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef) {
	if callErr := tryCFCharacterSetUnion(theSet, theOtherSet); callErr != nil {
		panic(callErr)
	}
}

var _cFCopyDescription func(cf CFTypeRef) CFStringRef
var _cFCopyDescriptionErr error

func tryCFCopyDescription(cf CFTypeRef) (CFStringRef, error) {
	if _cFCopyDescription == nil {
		return 0, symbolCallError("CFCopyDescription", "", _cFCopyDescriptionErr)
	}
	return _cFCopyDescription(cf), nil
}

// CFCopyDescription returns a textual description of a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCopyDescription(_:)
func CFCopyDescription(cf CFTypeRef) CFStringRef {
	result, callErr := tryCFCopyDescription(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFCopyTypeIDDescription func(type_id uint) CFStringRef
var _cFCopyTypeIDDescriptionErr error

func tryCFCopyTypeIDDescription(type_id uint) (CFStringRef, error) {
	if _cFCopyTypeIDDescription == nil {
		return 0, symbolCallError("CFCopyTypeIDDescription", "", _cFCopyTypeIDDescriptionErr)
	}
	return _cFCopyTypeIDDescription(type_id), nil
}

// CFCopyTypeIDDescription returns a textual description of a Core Foundation type, as identified by its type ID, which can be used when debugging.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCopyTypeIDDescription(_:)
func CFCopyTypeIDDescription(type_id uint) CFStringRef {
	result, callErr := tryCFCopyTypeIDDescription(type_id)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataAppendBytes func(theData CFMutableDataRef, bytes *uint8, length int)
var _cFDataAppendBytesErr error

func tryCFDataAppendBytes(theData CFMutableDataRef, bytes *uint8, length int) error {
	if _cFDataAppendBytes == nil {
		return symbolCallError("CFDataAppendBytes", "", _cFDataAppendBytesErr)
	}
	_cFDataAppendBytes(theData, bytes, length)
	return nil
}

// CFDataAppendBytes appends the bytes from a byte buffer to the contents of a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataAppendBytes(_:_:_:)
func CFDataAppendBytes(theData CFMutableDataRef, bytes *uint8, length int) {
	if callErr := tryCFDataAppendBytes(theData, bytes, length); callErr != nil {
		panic(callErr)
	}
}

var _cFDataCreate func(allocator CFAllocatorRef, bytes *uint8, length int) CFDataRef
var _cFDataCreateErr error

func tryCFDataCreate(allocator CFAllocatorRef, bytes *uint8, length int) (CFDataRef, error) {
	if _cFDataCreate == nil {
		return 0, symbolCallError("CFDataCreate", "", _cFDataCreateErr)
	}
	return _cFDataCreate(allocator, bytes, length), nil
}

// CFDataCreate creates an immutable CFData object using data copied from a specified byte buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreate(_:_:_:)
func CFDataCreate(allocator CFAllocatorRef, bytes *uint8, length int) CFDataRef {
	result, callErr := tryCFDataCreate(allocator, bytes, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataCreateCopy func(allocator CFAllocatorRef, theData CFDataRef) CFDataRef
var _cFDataCreateCopyErr error

func tryCFDataCreateCopy(allocator CFAllocatorRef, theData CFDataRef) (CFDataRef, error) {
	if _cFDataCreateCopy == nil {
		return 0, symbolCallError("CFDataCreateCopy", "", _cFDataCreateCopyErr)
	}
	return _cFDataCreateCopy(allocator, theData), nil
}

// CFDataCreateCopy creates an immutable copy of a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateCopy(_:_:)
func CFDataCreateCopy(allocator CFAllocatorRef, theData CFDataRef) CFDataRef {
	result, callErr := tryCFDataCreateCopy(allocator, theData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataCreateMutable func(allocator CFAllocatorRef, capacity int) CFMutableDataRef
var _cFDataCreateMutableErr error

func tryCFDataCreateMutable(allocator CFAllocatorRef, capacity int) (CFMutableDataRef, error) {
	if _cFDataCreateMutable == nil {
		return 0, symbolCallError("CFDataCreateMutable", "", _cFDataCreateMutableErr)
	}
	return _cFDataCreateMutable(allocator, capacity), nil
}

// CFDataCreateMutable creates an empty CFMutableData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutable(_:_:)
func CFDataCreateMutable(allocator CFAllocatorRef, capacity int) CFMutableDataRef {
	result, callErr := tryCFDataCreateMutable(allocator, capacity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theData CFDataRef) CFMutableDataRef
var _cFDataCreateMutableCopyErr error

func tryCFDataCreateMutableCopy(allocator CFAllocatorRef, capacity int, theData CFDataRef) (CFMutableDataRef, error) {
	if _cFDataCreateMutableCopy == nil {
		return 0, symbolCallError("CFDataCreateMutableCopy", "", _cFDataCreateMutableCopyErr)
	}
	return _cFDataCreateMutableCopy(allocator, capacity, theData), nil
}

// CFDataCreateMutableCopy creates a CFMutableData object by copying another CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutableCopy(_:_:_:)
func CFDataCreateMutableCopy(allocator CFAllocatorRef, capacity int, theData CFDataRef) CFMutableDataRef {
	result, callErr := tryCFDataCreateMutableCopy(allocator, capacity, theData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataCreateWithBytesNoCopy func(allocator CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFDataRef
var _cFDataCreateWithBytesNoCopyErr error

func tryCFDataCreateWithBytesNoCopy(allocator CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) (CFDataRef, error) {
	if _cFDataCreateWithBytesNoCopy == nil {
		return 0, symbolCallError("CFDataCreateWithBytesNoCopy", "", _cFDataCreateWithBytesNoCopyErr)
	}
	return _cFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator), nil
}

// CFDataCreateWithBytesNoCopy creates an immutable CFData object from an external (client-owned) byte buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateWithBytesNoCopy(_:_:_:_:)
func CFDataCreateWithBytesNoCopy(allocator CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFDataRef {
	result, callErr := tryCFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataDeleteBytes func(theData CFMutableDataRef, range_ CFRange)
var _cFDataDeleteBytesErr error

func tryCFDataDeleteBytes(theData CFMutableDataRef, range_ CFRange) error {
	if _cFDataDeleteBytes == nil {
		return symbolCallError("CFDataDeleteBytes", "", _cFDataDeleteBytesErr)
	}
	_cFDataDeleteBytes(theData, range_)
	return nil
}

// CFDataDeleteBytes deletes the bytes in a CFMutableData object within a specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataDeleteBytes(_:_:)
func CFDataDeleteBytes(theData CFMutableDataRef, range_ CFRange) {
	if callErr := tryCFDataDeleteBytes(theData, range_); callErr != nil {
		panic(callErr)
	}
}

var _cFDataFind func(theData CFDataRef, dataToFind CFDataRef, searchRange CFRange, compareOptions CFDataSearchFlags) CFRange
var _cFDataFindErr error

func tryCFDataFind(theData CFDataRef, dataToFind CFDataRef, searchRange CFRange, compareOptions CFDataSearchFlags) (CFRange, error) {
	if _cFDataFind == nil {
		return CFRange{}, symbolCallError("CFDataFind", "10.6", _cFDataFindErr)
	}
	return _cFDataFind(theData, dataToFind, searchRange, compareOptions), nil
}

// CFDataFind finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataFind(_:_:_:_:)
func CFDataFind(theData CFDataRef, dataToFind CFDataRef, searchRange CFRange, compareOptions CFDataSearchFlags) CFRange {
	result, callErr := tryCFDataFind(theData, dataToFind, searchRange, compareOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataGetBytePtr func(theData CFDataRef) *uint8
var _cFDataGetBytePtrErr error

func tryCFDataGetBytePtr(theData CFDataRef) (*uint8, error) {
	if _cFDataGetBytePtr == nil {
		return nil, symbolCallError("CFDataGetBytePtr", "", _cFDataGetBytePtrErr)
	}
	return _cFDataGetBytePtr(theData), nil
}

// CFDataGetBytePtr returns a read-only pointer to the bytes of a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytePtr(_:)
func CFDataGetBytePtr(theData CFDataRef) *uint8 {
	result, callErr := tryCFDataGetBytePtr(theData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataGetBytes func(theData CFDataRef, range_ CFRange, buffer *uint8)
var _cFDataGetBytesErr error

func tryCFDataGetBytes(theData CFDataRef, range_ CFRange, buffer *uint8) error {
	if _cFDataGetBytes == nil {
		return symbolCallError("CFDataGetBytes", "", _cFDataGetBytesErr)
	}
	_cFDataGetBytes(theData, range_, buffer)
	return nil
}

// CFDataGetBytes copies the byte contents of a CFData object to an external buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytes(_:_:_:)
func CFDataGetBytes(theData CFDataRef, range_ CFRange, buffer *uint8) {
	if callErr := tryCFDataGetBytes(theData, range_, buffer); callErr != nil {
		panic(callErr)
	}
}

var _cFDataGetLength func(theData CFDataRef) int
var _cFDataGetLengthErr error

func tryCFDataGetLength(theData CFDataRef) (int, error) {
	if _cFDataGetLength == nil {
		return 0, symbolCallError("CFDataGetLength", "", _cFDataGetLengthErr)
	}
	return _cFDataGetLength(theData), nil
}

// CFDataGetLength returns the number of bytes contained by a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetLength(_:)
func CFDataGetLength(theData CFDataRef) int {
	result, callErr := tryCFDataGetLength(theData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataGetMutableBytePtr func(theData CFMutableDataRef) *uint8
var _cFDataGetMutableBytePtrErr error

func tryCFDataGetMutableBytePtr(theData CFMutableDataRef) (*uint8, error) {
	if _cFDataGetMutableBytePtr == nil {
		return nil, symbolCallError("CFDataGetMutableBytePtr", "", _cFDataGetMutableBytePtrErr)
	}
	return _cFDataGetMutableBytePtr(theData), nil
}

// CFDataGetMutableBytePtr returns a pointer to a mutable byte buffer of a CFMutableData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetMutableBytePtr(_:)
func CFDataGetMutableBytePtr(theData CFMutableDataRef) *uint8 {
	result, callErr := tryCFDataGetMutableBytePtr(theData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataGetTypeID func() uint
var _cFDataGetTypeIDErr error

func tryCFDataGetTypeID() (uint, error) {
	if _cFDataGetTypeID == nil {
		return 0, symbolCallError("CFDataGetTypeID", "", _cFDataGetTypeIDErr)
	}
	return _cFDataGetTypeID(), nil
}

// CFDataGetTypeID returns the type identifier for the CFData opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetTypeID()
func CFDataGetTypeID() uint {
	result, callErr := tryCFDataGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDataIncreaseLength func(theData CFMutableDataRef, extraLength int)
var _cFDataIncreaseLengthErr error

func tryCFDataIncreaseLength(theData CFMutableDataRef, extraLength int) error {
	if _cFDataIncreaseLength == nil {
		return symbolCallError("CFDataIncreaseLength", "", _cFDataIncreaseLengthErr)
	}
	_cFDataIncreaseLength(theData, extraLength)
	return nil
}

// CFDataIncreaseLength increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataIncreaseLength(_:_:)
func CFDataIncreaseLength(theData CFMutableDataRef, extraLength int) {
	if callErr := tryCFDataIncreaseLength(theData, extraLength); callErr != nil {
		panic(callErr)
	}
}

var _cFDataReplaceBytes func(theData CFMutableDataRef, range_ CFRange, newBytes *uint8, newLength int)
var _cFDataReplaceBytesErr error

func tryCFDataReplaceBytes(theData CFMutableDataRef, range_ CFRange, newBytes *uint8, newLength int) error {
	if _cFDataReplaceBytes == nil {
		return symbolCallError("CFDataReplaceBytes", "", _cFDataReplaceBytesErr)
	}
	_cFDataReplaceBytes(theData, range_, newBytes, newLength)
	return nil
}

// CFDataReplaceBytes replaces those bytes in a CFMutableData object that fall within a specified range with other bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataReplaceBytes(_:_:_:_:)
func CFDataReplaceBytes(theData CFMutableDataRef, range_ CFRange, newBytes *uint8, newLength int) {
	if callErr := tryCFDataReplaceBytes(theData, range_, newBytes, newLength); callErr != nil {
		panic(callErr)
	}
}

var _cFDataSetLength func(theData CFMutableDataRef, length int)
var _cFDataSetLengthErr error

func tryCFDataSetLength(theData CFMutableDataRef, length int) error {
	if _cFDataSetLength == nil {
		return symbolCallError("CFDataSetLength", "", _cFDataSetLengthErr)
	}
	_cFDataSetLength(theData, length)
	return nil
}

// CFDataSetLength resets the length of a CFMutableData object’s internal byte buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataSetLength(_:_:)
func CFDataSetLength(theData CFMutableDataRef, length int) {
	if callErr := tryCFDataSetLength(theData, length); callErr != nil {
		panic(callErr)
	}
}

var _cFDateCompare func(theDate CFDateRef, otherDate CFDateRef, context unsafe.Pointer) CFComparisonResult
var _cFDateCompareErr error

func tryCFDateCompare(theDate CFDateRef, otherDate CFDateRef, context unsafe.Pointer) (CFComparisonResult, error) {
	if _cFDateCompare == nil {
		return *new(CFComparisonResult), symbolCallError("CFDateCompare", "", _cFDateCompareErr)
	}
	return _cFDateCompare(theDate, otherDate, context), nil
}

// CFDateCompare compares two [CFDate] objects and returns a comparison result.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateCompare(_:_:_:)
func CFDateCompare(theDate CFDateRef, otherDate CFDateRef, context unsafe.Pointer) CFComparisonResult {
	result, callErr := tryCFDateCompare(theDate, otherDate, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateCreate func(allocator CFAllocatorRef, at CFAbsoluteTime) CFDateRef
var _cFDateCreateErr error

func tryCFDateCreate(allocator CFAllocatorRef, at CFAbsoluteTime) (CFDateRef, error) {
	if _cFDateCreate == nil {
		return 0, symbolCallError("CFDateCreate", "", _cFDateCreateErr)
	}
	return _cFDateCreate(allocator, at), nil
}

// CFDateCreate creates a [CFDate] object given an absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateCreate(_:_:)
func CFDateCreate(allocator CFAllocatorRef, at CFAbsoluteTime) CFDateRef {
	result, callErr := tryCFDateCreate(allocator, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCopyProperty func(formatter CFDateFormatterRef, key CFDateFormatterKey) CFTypeRef
var _cFDateFormatterCopyPropertyErr error

func tryCFDateFormatterCopyProperty(formatter CFDateFormatterRef, key CFDateFormatterKey) (CFTypeRef, error) {
	if _cFDateFormatterCopyProperty == nil {
		return 0, symbolCallError("CFDateFormatterCopyProperty", "", _cFDateFormatterCopyPropertyErr)
	}
	return _cFDateFormatterCopyProperty(formatter, key), nil
}

// CFDateFormatterCopyProperty returns a copy of a date formatter’s value for a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCopyProperty(_:_:)
func CFDateFormatterCopyProperty(formatter CFDateFormatterRef, key CFDateFormatterKey) CFTypeRef {
	result, callErr := tryCFDateFormatterCopyProperty(formatter, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCreate func(allocator CFAllocatorRef, locale CFLocaleRef, dateStyle CFDateFormatterStyle, timeStyle CFDateFormatterStyle) CFDateFormatterRef
var _cFDateFormatterCreateErr error

func tryCFDateFormatterCreate(allocator CFAllocatorRef, locale CFLocaleRef, dateStyle CFDateFormatterStyle, timeStyle CFDateFormatterStyle) (CFDateFormatterRef, error) {
	if _cFDateFormatterCreate == nil {
		return 0, symbolCallError("CFDateFormatterCreate", "", _cFDateFormatterCreateErr)
	}
	return _cFDateFormatterCreate(allocator, locale, dateStyle, timeStyle), nil
}

// CFDateFormatterCreate creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreate(_:_:_:_:)
func CFDateFormatterCreate(allocator CFAllocatorRef, locale CFLocaleRef, dateStyle CFDateFormatterStyle, timeStyle CFDateFormatterStyle) CFDateFormatterRef {
	result, callErr := tryCFDateFormatterCreate(allocator, locale, dateStyle, timeStyle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCreateDateFormatFromTemplate func(allocator CFAllocatorRef, tmplate CFStringRef, options uint64, locale CFLocaleRef) CFStringRef
var _cFDateFormatterCreateDateFormatFromTemplateErr error

func tryCFDateFormatterCreateDateFormatFromTemplate(allocator CFAllocatorRef, tmplate CFStringRef, options uint64, locale CFLocaleRef) (CFStringRef, error) {
	if _cFDateFormatterCreateDateFormatFromTemplate == nil {
		return 0, symbolCallError("CFDateFormatterCreateDateFormatFromTemplate", "10.6", _cFDateFormatterCreateDateFormatFromTemplateErr)
	}
	return _cFDateFormatterCreateDateFormatFromTemplate(allocator, tmplate, options, locale), nil
}

// CFDateFormatterCreateDateFormatFromTemplate returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFormatFromTemplate(_:_:_:_:)
func CFDateFormatterCreateDateFormatFromTemplate(allocator CFAllocatorRef, tmplate CFStringRef, options uint64, locale CFLocaleRef) CFStringRef {
	result, callErr := tryCFDateFormatterCreateDateFormatFromTemplate(allocator, tmplate, options, locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCreateDateFromString func(allocator CFAllocatorRef, formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange) CFDateRef
var _cFDateFormatterCreateDateFromStringErr error

func tryCFDateFormatterCreateDateFromString(allocator CFAllocatorRef, formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange) (CFDateRef, error) {
	if _cFDateFormatterCreateDateFromString == nil {
		return 0, symbolCallError("CFDateFormatterCreateDateFromString", "", _cFDateFormatterCreateDateFromStringErr)
	}
	return _cFDateFormatterCreateDateFromString(allocator, formatter, string_, rangep), nil
}

// CFDateFormatterCreateDateFromString returns a date object representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFromString(_:_:_:_:)
func CFDateFormatterCreateDateFromString(allocator CFAllocatorRef, formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange) CFDateRef {
	result, callErr := tryCFDateFormatterCreateDateFromString(allocator, formatter, string_, rangep)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCreateISO8601Formatter func(allocator CFAllocatorRef, formatOptions CFISO8601DateFormatOptions) CFDateFormatterRef
var _cFDateFormatterCreateISO8601FormatterErr error

func tryCFDateFormatterCreateISO8601Formatter(allocator CFAllocatorRef, formatOptions CFISO8601DateFormatOptions) (CFDateFormatterRef, error) {
	if _cFDateFormatterCreateISO8601Formatter == nil {
		return 0, symbolCallError("CFDateFormatterCreateISO8601Formatter", "10.12", _cFDateFormatterCreateISO8601FormatterErr)
	}
	return _cFDateFormatterCreateISO8601Formatter(allocator, formatOptions), nil
}

// CFDateFormatterCreateISO8601Formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateISO8601Formatter(_:_:)
func CFDateFormatterCreateISO8601Formatter(allocator CFAllocatorRef, formatOptions CFISO8601DateFormatOptions) CFDateFormatterRef {
	result, callErr := tryCFDateFormatterCreateISO8601Formatter(allocator, formatOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCreateStringWithAbsoluteTime func(allocator CFAllocatorRef, formatter CFDateFormatterRef, at CFAbsoluteTime) CFStringRef
var _cFDateFormatterCreateStringWithAbsoluteTimeErr error

func tryCFDateFormatterCreateStringWithAbsoluteTime(allocator CFAllocatorRef, formatter CFDateFormatterRef, at CFAbsoluteTime) (CFStringRef, error) {
	if _cFDateFormatterCreateStringWithAbsoluteTime == nil {
		return 0, symbolCallError("CFDateFormatterCreateStringWithAbsoluteTime", "", _cFDateFormatterCreateStringWithAbsoluteTimeErr)
	}
	return _cFDateFormatterCreateStringWithAbsoluteTime(allocator, formatter, at), nil
}

// CFDateFormatterCreateStringWithAbsoluteTime returns a string representation of the given absolute time using the specified date formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithAbsoluteTime(_:_:_:)
func CFDateFormatterCreateStringWithAbsoluteTime(allocator CFAllocatorRef, formatter CFDateFormatterRef, at CFAbsoluteTime) CFStringRef {
	result, callErr := tryCFDateFormatterCreateStringWithAbsoluteTime(allocator, formatter, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterCreateStringWithDate func(allocator CFAllocatorRef, formatter CFDateFormatterRef, date CFDateRef) CFStringRef
var _cFDateFormatterCreateStringWithDateErr error

func tryCFDateFormatterCreateStringWithDate(allocator CFAllocatorRef, formatter CFDateFormatterRef, date CFDateRef) (CFStringRef, error) {
	if _cFDateFormatterCreateStringWithDate == nil {
		return 0, symbolCallError("CFDateFormatterCreateStringWithDate", "", _cFDateFormatterCreateStringWithDateErr)
	}
	return _cFDateFormatterCreateStringWithDate(allocator, formatter, date), nil
}

// CFDateFormatterCreateStringWithDate returns a string representation of the given date using the specified date formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithDate(_:_:_:)
func CFDateFormatterCreateStringWithDate(allocator CFAllocatorRef, formatter CFDateFormatterRef, date CFDateRef) CFStringRef {
	result, callErr := tryCFDateFormatterCreateStringWithDate(allocator, formatter, date)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterGetAbsoluteTimeFromString func(formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange, atp *CFAbsoluteTime) bool
var _cFDateFormatterGetAbsoluteTimeFromStringErr error

func tryCFDateFormatterGetAbsoluteTimeFromString(formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange, atp *CFAbsoluteTime) (bool, error) {
	if _cFDateFormatterGetAbsoluteTimeFromString == nil {
		return false, symbolCallError("CFDateFormatterGetAbsoluteTimeFromString", "", _cFDateFormatterGetAbsoluteTimeFromStringErr)
	}
	return _cFDateFormatterGetAbsoluteTimeFromString(formatter, string_, rangep, atp), nil
}

// CFDateFormatterGetAbsoluteTimeFromString returns an absolute time object representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetAbsoluteTimeFromString(_:_:_:_:)
func CFDateFormatterGetAbsoluteTimeFromString(formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange, atp *CFAbsoluteTime) bool {
	result, callErr := tryCFDateFormatterGetAbsoluteTimeFromString(formatter, string_, rangep, atp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterGetDateStyle func(formatter CFDateFormatterRef) CFDateFormatterStyle
var _cFDateFormatterGetDateStyleErr error

func tryCFDateFormatterGetDateStyle(formatter CFDateFormatterRef) (CFDateFormatterStyle, error) {
	if _cFDateFormatterGetDateStyle == nil {
		return *new(CFDateFormatterStyle), symbolCallError("CFDateFormatterGetDateStyle", "", _cFDateFormatterGetDateStyleErr)
	}
	return _cFDateFormatterGetDateStyle(formatter), nil
}

// CFDateFormatterGetDateStyle returns the date style used to create the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetDateStyle(_:)
func CFDateFormatterGetDateStyle(formatter CFDateFormatterRef) CFDateFormatterStyle {
	result, callErr := tryCFDateFormatterGetDateStyle(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterGetFormat func(formatter CFDateFormatterRef) CFStringRef
var _cFDateFormatterGetFormatErr error

func tryCFDateFormatterGetFormat(formatter CFDateFormatterRef) (CFStringRef, error) {
	if _cFDateFormatterGetFormat == nil {
		return 0, symbolCallError("CFDateFormatterGetFormat", "", _cFDateFormatterGetFormatErr)
	}
	return _cFDateFormatterGetFormat(formatter), nil
}

// CFDateFormatterGetFormat returns a format string for the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetFormat(_:)
func CFDateFormatterGetFormat(formatter CFDateFormatterRef) CFStringRef {
	result, callErr := tryCFDateFormatterGetFormat(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterGetLocale func(formatter CFDateFormatterRef) CFLocaleRef
var _cFDateFormatterGetLocaleErr error

func tryCFDateFormatterGetLocale(formatter CFDateFormatterRef) (CFLocaleRef, error) {
	if _cFDateFormatterGetLocale == nil {
		return 0, symbolCallError("CFDateFormatterGetLocale", "", _cFDateFormatterGetLocaleErr)
	}
	return _cFDateFormatterGetLocale(formatter), nil
}

// CFDateFormatterGetLocale returns the locale object used to create the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetLocale(_:)
func CFDateFormatterGetLocale(formatter CFDateFormatterRef) CFLocaleRef {
	result, callErr := tryCFDateFormatterGetLocale(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterGetTimeStyle func(formatter CFDateFormatterRef) CFDateFormatterStyle
var _cFDateFormatterGetTimeStyleErr error

func tryCFDateFormatterGetTimeStyle(formatter CFDateFormatterRef) (CFDateFormatterStyle, error) {
	if _cFDateFormatterGetTimeStyle == nil {
		return *new(CFDateFormatterStyle), symbolCallError("CFDateFormatterGetTimeStyle", "", _cFDateFormatterGetTimeStyleErr)
	}
	return _cFDateFormatterGetTimeStyle(formatter), nil
}

// CFDateFormatterGetTimeStyle returns the time style used to create the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTimeStyle(_:)
func CFDateFormatterGetTimeStyle(formatter CFDateFormatterRef) CFDateFormatterStyle {
	result, callErr := tryCFDateFormatterGetTimeStyle(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterGetTypeID func() uint
var _cFDateFormatterGetTypeIDErr error

func tryCFDateFormatterGetTypeID() (uint, error) {
	if _cFDateFormatterGetTypeID == nil {
		return 0, symbolCallError("CFDateFormatterGetTypeID", "", _cFDateFormatterGetTypeIDErr)
	}
	return _cFDateFormatterGetTypeID(), nil
}

// CFDateFormatterGetTypeID returns the type identifier for CFDateFormatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTypeID()
func CFDateFormatterGetTypeID() uint {
	result, callErr := tryCFDateFormatterGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateFormatterSetFormat func(formatter CFDateFormatterRef, formatString CFStringRef)
var _cFDateFormatterSetFormatErr error

func tryCFDateFormatterSetFormat(formatter CFDateFormatterRef, formatString CFStringRef) error {
	if _cFDateFormatterSetFormat == nil {
		return symbolCallError("CFDateFormatterSetFormat", "", _cFDateFormatterSetFormatErr)
	}
	_cFDateFormatterSetFormat(formatter, formatString)
	return nil
}

// CFDateFormatterSetFormat sets the format string of the given date formatter to the specified value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetFormat(_:_:)
func CFDateFormatterSetFormat(formatter CFDateFormatterRef, formatString CFStringRef) {
	if callErr := tryCFDateFormatterSetFormat(formatter, formatString); callErr != nil {
		panic(callErr)
	}
}

var _cFDateFormatterSetProperty func(formatter CFDateFormatterRef, key CFStringRef, value CFTypeRef)
var _cFDateFormatterSetPropertyErr error

func tryCFDateFormatterSetProperty(formatter CFDateFormatterRef, key CFStringRef, value CFTypeRef) error {
	if _cFDateFormatterSetProperty == nil {
		return symbolCallError("CFDateFormatterSetProperty", "", _cFDateFormatterSetPropertyErr)
	}
	_cFDateFormatterSetProperty(formatter, key, value)
	return nil
}

// CFDateFormatterSetProperty sets a date formatter property using a key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetProperty(_:_:_:)
func CFDateFormatterSetProperty(formatter CFDateFormatterRef, key CFStringRef, value CFTypeRef) {
	if callErr := tryCFDateFormatterSetProperty(formatter, key, value); callErr != nil {
		panic(callErr)
	}
}

var _cFDateGetAbsoluteTime func(theDate CFDateRef) CFAbsoluteTime
var _cFDateGetAbsoluteTimeErr error

func tryCFDateGetAbsoluteTime(theDate CFDateRef) (CFAbsoluteTime, error) {
	if _cFDateGetAbsoluteTime == nil {
		return *new(CFAbsoluteTime), symbolCallError("CFDateGetAbsoluteTime", "", _cFDateGetAbsoluteTimeErr)
	}
	return _cFDateGetAbsoluteTime(theDate), nil
}

// CFDateGetAbsoluteTime returns a [CFDate] object’s absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateGetAbsoluteTime(_:)
func CFDateGetAbsoluteTime(theDate CFDateRef) CFAbsoluteTime {
	result, callErr := tryCFDateGetAbsoluteTime(theDate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateGetTimeIntervalSinceDate func(theDate CFDateRef, otherDate CFDateRef) float64
var _cFDateGetTimeIntervalSinceDateErr error

func tryCFDateGetTimeIntervalSinceDate(theDate CFDateRef, otherDate CFDateRef) (float64, error) {
	if _cFDateGetTimeIntervalSinceDate == nil {
		return 0.0, symbolCallError("CFDateGetTimeIntervalSinceDate", "", _cFDateGetTimeIntervalSinceDateErr)
	}
	return _cFDateGetTimeIntervalSinceDate(theDate, otherDate), nil
}

// CFDateGetTimeIntervalSinceDate returns the number of elapsed seconds between the given [CFDate] objects.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTimeIntervalSinceDate(_:_:)
func CFDateGetTimeIntervalSinceDate(theDate CFDateRef, otherDate CFDateRef) float64 {
	result, callErr := tryCFDateGetTimeIntervalSinceDate(theDate, otherDate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDateGetTypeID func() uint
var _cFDateGetTypeIDErr error

func tryCFDateGetTypeID() (uint, error) {
	if _cFDateGetTypeID == nil {
		return 0, symbolCallError("CFDateGetTypeID", "", _cFDateGetTypeIDErr)
	}
	return _cFDateGetTypeID(), nil
}

// CFDateGetTypeID returns the type identifier for the [CFDate] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTypeID()
func CFDateGetTypeID() uint {
	result, callErr := tryCFDateGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryAddValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer)
var _cFDictionaryAddValueErr error

func tryCFDictionaryAddValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) error {
	if _cFDictionaryAddValue == nil {
		return symbolCallError("CFDictionaryAddValue", "", _cFDictionaryAddValueErr)
	}
	_cFDictionaryAddValue(theDict, key, value)
	return nil
}

// CFDictionaryAddValue adds a key-value pair to a dictionary if the specified key is not already present.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryAddValue(_:_:_:)
func CFDictionaryAddValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) {
	if callErr := tryCFDictionaryAddValue(theDict, key, value); callErr != nil {
		panic(callErr)
	}
}

var _cFDictionaryApplyFunction func(theDict CFDictionaryRef, applier CFDictionaryApplierFunction, context unsafe.Pointer)
var _cFDictionaryApplyFunctionErr error

func tryCFDictionaryApplyFunction(theDict CFDictionaryRef, applier CFDictionaryApplierFunction, context unsafe.Pointer) error {
	if _cFDictionaryApplyFunction == nil {
		return symbolCallError("CFDictionaryApplyFunction", "", _cFDictionaryApplyFunctionErr)
	}
	_cFDictionaryApplyFunction(theDict, applier, context)
	return nil
}

// CFDictionaryApplyFunction calls a function once for each key-value pair in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplyFunction(_:_:_:)
func CFDictionaryApplyFunction(theDict CFDictionaryRef, applier CFDictionaryApplierFunction, context unsafe.Pointer) {
	if callErr := tryCFDictionaryApplyFunction(theDict, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cFDictionaryContainsKey func(theDict CFDictionaryRef, key unsafe.Pointer) bool
var _cFDictionaryContainsKeyErr error

func tryCFDictionaryContainsKey(theDict CFDictionaryRef, key unsafe.Pointer) (bool, error) {
	if _cFDictionaryContainsKey == nil {
		return false, symbolCallError("CFDictionaryContainsKey", "", _cFDictionaryContainsKeyErr)
	}
	return _cFDictionaryContainsKey(theDict, key), nil
}

// CFDictionaryContainsKey returns a Boolean value that indicates whether a given key is in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsKey(_:_:)
func CFDictionaryContainsKey(theDict CFDictionaryRef, key unsafe.Pointer) bool {
	result, callErr := tryCFDictionaryContainsKey(theDict, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryContainsValue func(theDict CFDictionaryRef, value unsafe.Pointer) bool
var _cFDictionaryContainsValueErr error

func tryCFDictionaryContainsValue(theDict CFDictionaryRef, value unsafe.Pointer) (bool, error) {
	if _cFDictionaryContainsValue == nil {
		return false, symbolCallError("CFDictionaryContainsValue", "", _cFDictionaryContainsValueErr)
	}
	return _cFDictionaryContainsValue(theDict, value), nil
}

// CFDictionaryContainsValue returns a Boolean value that indicates whether a given value is in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsValue(_:_:)
func CFDictionaryContainsValue(theDict CFDictionaryRef, value unsafe.Pointer) bool {
	result, callErr := tryCFDictionaryContainsValue(theDict, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryCreate func(allocator CFAllocatorRef, keys unsafe.Pointer, values unsafe.Pointer, numValues int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFDictionaryRef
var _cFDictionaryCreateErr error

func tryCFDictionaryCreate(allocator CFAllocatorRef, keys unsafe.Pointer, values unsafe.Pointer, numValues int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) (CFDictionaryRef, error) {
	if _cFDictionaryCreate == nil {
		return 0, symbolCallError("CFDictionaryCreate", "", _cFDictionaryCreateErr)
	}
	return _cFDictionaryCreate(allocator, keys, values, numValues, keyCallBacks, valueCallBacks), nil
}

// CFDictionaryCreate creates an immutable dictionary containing the specified key-value pairs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreate(_:_:_:_:_:_:)
func CFDictionaryCreate(allocator CFAllocatorRef, keys unsafe.Pointer, values unsafe.Pointer, numValues int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFDictionaryRef {
	result, callErr := tryCFDictionaryCreate(allocator, keys, values, numValues, keyCallBacks, valueCallBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryCreateCopy func(allocator CFAllocatorRef, theDict CFDictionaryRef) CFDictionaryRef
var _cFDictionaryCreateCopyErr error

func tryCFDictionaryCreateCopy(allocator CFAllocatorRef, theDict CFDictionaryRef) (CFDictionaryRef, error) {
	if _cFDictionaryCreateCopy == nil {
		return 0, symbolCallError("CFDictionaryCreateCopy", "", _cFDictionaryCreateCopyErr)
	}
	return _cFDictionaryCreateCopy(allocator, theDict), nil
}

// CFDictionaryCreateCopy creates and returns a new immutable dictionary with the key-value pairs of another dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateCopy(_:_:)
func CFDictionaryCreateCopy(allocator CFAllocatorRef, theDict CFDictionaryRef) CFDictionaryRef {
	result, callErr := tryCFDictionaryCreateCopy(allocator, theDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryCreateMutable func(allocator CFAllocatorRef, capacity int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFMutableDictionaryRef
var _cFDictionaryCreateMutableErr error

func tryCFDictionaryCreateMutable(allocator CFAllocatorRef, capacity int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) (CFMutableDictionaryRef, error) {
	if _cFDictionaryCreateMutable == nil {
		return 0, symbolCallError("CFDictionaryCreateMutable", "", _cFDictionaryCreateMutableErr)
	}
	return _cFDictionaryCreateMutable(allocator, capacity, keyCallBacks, valueCallBacks), nil
}

// CFDictionaryCreateMutable creates a new mutable dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutable(_:_:_:_:)
func CFDictionaryCreateMutable(allocator CFAllocatorRef, capacity int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFMutableDictionaryRef {
	result, callErr := tryCFDictionaryCreateMutable(allocator, capacity, keyCallBacks, valueCallBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theDict CFDictionaryRef) CFMutableDictionaryRef
var _cFDictionaryCreateMutableCopyErr error

func tryCFDictionaryCreateMutableCopy(allocator CFAllocatorRef, capacity int, theDict CFDictionaryRef) (CFMutableDictionaryRef, error) {
	if _cFDictionaryCreateMutableCopy == nil {
		return 0, symbolCallError("CFDictionaryCreateMutableCopy", "", _cFDictionaryCreateMutableCopyErr)
	}
	return _cFDictionaryCreateMutableCopy(allocator, capacity, theDict), nil
}

// CFDictionaryCreateMutableCopy creates a new mutable dictionary with the key-value pairs from another dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutableCopy(_:_:_:)
func CFDictionaryCreateMutableCopy(allocator CFAllocatorRef, capacity int, theDict CFDictionaryRef) CFMutableDictionaryRef {
	result, callErr := tryCFDictionaryCreateMutableCopy(allocator, capacity, theDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryGetCount func(theDict CFDictionaryRef) int
var _cFDictionaryGetCountErr error

func tryCFDictionaryGetCount(theDict CFDictionaryRef) (int, error) {
	if _cFDictionaryGetCount == nil {
		return 0, symbolCallError("CFDictionaryGetCount", "", _cFDictionaryGetCountErr)
	}
	return _cFDictionaryGetCount(theDict), nil
}

// CFDictionaryGetCount returns the number of key-value pairs in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCount(_:)
func CFDictionaryGetCount(theDict CFDictionaryRef) int {
	result, callErr := tryCFDictionaryGetCount(theDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryGetCountOfKey func(theDict CFDictionaryRef, key unsafe.Pointer) int
var _cFDictionaryGetCountOfKeyErr error

func tryCFDictionaryGetCountOfKey(theDict CFDictionaryRef, key unsafe.Pointer) (int, error) {
	if _cFDictionaryGetCountOfKey == nil {
		return 0, symbolCallError("CFDictionaryGetCountOfKey", "", _cFDictionaryGetCountOfKeyErr)
	}
	return _cFDictionaryGetCountOfKey(theDict, key), nil
}

// CFDictionaryGetCountOfKey returns the number of times a key occurs in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfKey(_:_:)
func CFDictionaryGetCountOfKey(theDict CFDictionaryRef, key unsafe.Pointer) int {
	result, callErr := tryCFDictionaryGetCountOfKey(theDict, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryGetCountOfValue func(theDict CFDictionaryRef, value unsafe.Pointer) int
var _cFDictionaryGetCountOfValueErr error

func tryCFDictionaryGetCountOfValue(theDict CFDictionaryRef, value unsafe.Pointer) (int, error) {
	if _cFDictionaryGetCountOfValue == nil {
		return 0, symbolCallError("CFDictionaryGetCountOfValue", "", _cFDictionaryGetCountOfValueErr)
	}
	return _cFDictionaryGetCountOfValue(theDict, value), nil
}

// CFDictionaryGetCountOfValue counts the number of times a given value occurs in the dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfValue(_:_:)
func CFDictionaryGetCountOfValue(theDict CFDictionaryRef, value unsafe.Pointer) int {
	result, callErr := tryCFDictionaryGetCountOfValue(theDict, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryGetKeysAndValues func(theDict CFDictionaryRef, keys unsafe.Pointer, values unsafe.Pointer)
var _cFDictionaryGetKeysAndValuesErr error

func tryCFDictionaryGetKeysAndValues(theDict CFDictionaryRef, keys unsafe.Pointer, values unsafe.Pointer) error {
	if _cFDictionaryGetKeysAndValues == nil {
		return symbolCallError("CFDictionaryGetKeysAndValues", "", _cFDictionaryGetKeysAndValuesErr)
	}
	_cFDictionaryGetKeysAndValues(theDict, keys, values)
	return nil
}

// CFDictionaryGetKeysAndValues fills two buffers with the keys and values from a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetKeysAndValues(_:_:_:)
func CFDictionaryGetKeysAndValues(theDict CFDictionaryRef, keys unsafe.Pointer, values unsafe.Pointer) {
	if callErr := tryCFDictionaryGetKeysAndValues(theDict, keys, values); callErr != nil {
		panic(callErr)
	}
}

var _cFDictionaryGetTypeID func() uint
var _cFDictionaryGetTypeIDErr error

func tryCFDictionaryGetTypeID() (uint, error) {
	if _cFDictionaryGetTypeID == nil {
		return 0, symbolCallError("CFDictionaryGetTypeID", "", _cFDictionaryGetTypeIDErr)
	}
	return _cFDictionaryGetTypeID(), nil
}

// CFDictionaryGetTypeID returns the type identifier for the CFDictionary opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetTypeID()
func CFDictionaryGetTypeID() uint {
	result, callErr := tryCFDictionaryGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryGetValue func(theDict CFDictionaryRef, key unsafe.Pointer) unsafe.Pointer
var _cFDictionaryGetValueErr error

func tryCFDictionaryGetValue(theDict CFDictionaryRef, key unsafe.Pointer) (unsafe.Pointer, error) {
	if _cFDictionaryGetValue == nil {
		return nil, symbolCallError("CFDictionaryGetValue", "", _cFDictionaryGetValueErr)
	}
	return _cFDictionaryGetValue(theDict, key), nil
}

// CFDictionaryGetValue returns the value associated with a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValue(_:_:)
func CFDictionaryGetValue(theDict CFDictionaryRef, key unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryCFDictionaryGetValue(theDict, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryGetValueIfPresent func(theDict CFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool
var _cFDictionaryGetValueIfPresentErr error

func tryCFDictionaryGetValueIfPresent(theDict CFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) (bool, error) {
	if _cFDictionaryGetValueIfPresent == nil {
		return false, symbolCallError("CFDictionaryGetValueIfPresent", "", _cFDictionaryGetValueIfPresentErr)
	}
	return _cFDictionaryGetValueIfPresent(theDict, key, value), nil
}

// CFDictionaryGetValueIfPresent returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValueIfPresent(_:_:_:)
func CFDictionaryGetValueIfPresent(theDict CFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	result, callErr := tryCFDictionaryGetValueIfPresent(theDict, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFDictionaryRemoveAllValues func(theDict CFMutableDictionaryRef)
var _cFDictionaryRemoveAllValuesErr error

func tryCFDictionaryRemoveAllValues(theDict CFMutableDictionaryRef) error {
	if _cFDictionaryRemoveAllValues == nil {
		return symbolCallError("CFDictionaryRemoveAllValues", "", _cFDictionaryRemoveAllValuesErr)
	}
	_cFDictionaryRemoveAllValues(theDict)
	return nil
}

// CFDictionaryRemoveAllValues removes all the key-value pairs from a dictionary, making it empty.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveAllValues(_:)
func CFDictionaryRemoveAllValues(theDict CFMutableDictionaryRef) {
	if callErr := tryCFDictionaryRemoveAllValues(theDict); callErr != nil {
		panic(callErr)
	}
}

var _cFDictionaryRemoveValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer)
var _cFDictionaryRemoveValueErr error

func tryCFDictionaryRemoveValue(theDict CFMutableDictionaryRef, key unsafe.Pointer) error {
	if _cFDictionaryRemoveValue == nil {
		return symbolCallError("CFDictionaryRemoveValue", "", _cFDictionaryRemoveValueErr)
	}
	_cFDictionaryRemoveValue(theDict, key)
	return nil
}

// CFDictionaryRemoveValue removes a key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveValue(_:_:)
func CFDictionaryRemoveValue(theDict CFMutableDictionaryRef, key unsafe.Pointer) {
	if callErr := tryCFDictionaryRemoveValue(theDict, key); callErr != nil {
		panic(callErr)
	}
}

var _cFDictionaryReplaceValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer)
var _cFDictionaryReplaceValueErr error

func tryCFDictionaryReplaceValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) error {
	if _cFDictionaryReplaceValue == nil {
		return symbolCallError("CFDictionaryReplaceValue", "", _cFDictionaryReplaceValueErr)
	}
	_cFDictionaryReplaceValue(theDict, key, value)
	return nil
}

// CFDictionaryReplaceValue replaces a value corresponding to a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryReplaceValue(_:_:_:)
func CFDictionaryReplaceValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) {
	if callErr := tryCFDictionaryReplaceValue(theDict, key, value); callErr != nil {
		panic(callErr)
	}
}

var _cFDictionarySetValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer)
var _cFDictionarySetValueErr error

func tryCFDictionarySetValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) error {
	if _cFDictionarySetValue == nil {
		return symbolCallError("CFDictionarySetValue", "", _cFDictionarySetValueErr)
	}
	_cFDictionarySetValue(theDict, key, value)
	return nil
}

// CFDictionarySetValue sets the value corresponding to a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionarySetValue(_:_:_:)
func CFDictionarySetValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) {
	if callErr := tryCFDictionarySetValue(theDict, key, value); callErr != nil {
		panic(callErr)
	}
}

var _cFEqual func(cf1 CFTypeRef, cf2 CFTypeRef) bool
var _cFEqualErr error

func tryCFEqual(cf1 CFTypeRef, cf2 CFTypeRef) (bool, error) {
	if _cFEqual == nil {
		return false, symbolCallError("CFEqual", "", _cFEqualErr)
	}
	return _cFEqual(cf1, cf2), nil
}

// CFEqual determines whether two Core Foundation objects are considered equal.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFEqual(_:_:)
func CFEqual(cf1 CFTypeRef, cf2 CFTypeRef) bool {
	result, callErr := tryCFEqual(cf1, cf2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorCopyDescription func(err CFErrorRef) CFStringRef
var _cFErrorCopyDescriptionErr error

func tryCFErrorCopyDescription(err CFErrorRef) (CFStringRef, error) {
	if _cFErrorCopyDescription == nil {
		return 0, symbolCallError("CFErrorCopyDescription", "10.5", _cFErrorCopyDescriptionErr)
	}
	return _cFErrorCopyDescription(err), nil
}

// CFErrorCopyDescription returns a human-presentable description for a given error.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyDescription(_:)
func CFErrorCopyDescription(err CFErrorRef) CFStringRef {
	result, callErr := tryCFErrorCopyDescription(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorCopyFailureReason func(err CFErrorRef) CFStringRef
var _cFErrorCopyFailureReasonErr error

func tryCFErrorCopyFailureReason(err CFErrorRef) (CFStringRef, error) {
	if _cFErrorCopyFailureReason == nil {
		return 0, symbolCallError("CFErrorCopyFailureReason", "10.5", _cFErrorCopyFailureReasonErr)
	}
	return _cFErrorCopyFailureReason(err), nil
}

// CFErrorCopyFailureReason returns a human-presentable failure reason for a given error.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyFailureReason(_:)
func CFErrorCopyFailureReason(err CFErrorRef) CFStringRef {
	result, callErr := tryCFErrorCopyFailureReason(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorCopyRecoverySuggestion func(err CFErrorRef) CFStringRef
var _cFErrorCopyRecoverySuggestionErr error

func tryCFErrorCopyRecoverySuggestion(err CFErrorRef) (CFStringRef, error) {
	if _cFErrorCopyRecoverySuggestion == nil {
		return 0, symbolCallError("CFErrorCopyRecoverySuggestion", "10.5", _cFErrorCopyRecoverySuggestionErr)
	}
	return _cFErrorCopyRecoverySuggestion(err), nil
}

// CFErrorCopyRecoverySuggestion returns a human presentable recovery suggestion for a given error.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyRecoverySuggestion(_:)
func CFErrorCopyRecoverySuggestion(err CFErrorRef) CFStringRef {
	result, callErr := tryCFErrorCopyRecoverySuggestion(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorCopyUserInfo func(err CFErrorRef) CFDictionaryRef
var _cFErrorCopyUserInfoErr error

func tryCFErrorCopyUserInfo(err CFErrorRef) (CFDictionaryRef, error) {
	if _cFErrorCopyUserInfo == nil {
		return 0, symbolCallError("CFErrorCopyUserInfo", "10.5", _cFErrorCopyUserInfoErr)
	}
	return _cFErrorCopyUserInfo(err), nil
}

// CFErrorCopyUserInfo returns the user info dictionary for a given CFError.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyUserInfo(_:)
func CFErrorCopyUserInfo(err CFErrorRef) CFDictionaryRef {
	result, callErr := tryCFErrorCopyUserInfo(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorCreate func(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfo CFDictionaryRef) CFErrorRef
var _cFErrorCreateErr error

func tryCFErrorCreate(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfo CFDictionaryRef) (CFErrorRef, error) {
	if _cFErrorCreate == nil {
		return 0, symbolCallError("CFErrorCreate", "10.5", _cFErrorCreateErr)
	}
	return _cFErrorCreate(allocator, domain, code, userInfo), nil
}

// CFErrorCreate creates a new CFError object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreate(_:_:_:_:)
func CFErrorCreate(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfo CFDictionaryRef) CFErrorRef {
	result, callErr := tryCFErrorCreate(allocator, domain, code, userInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorCreateWithUserInfoKeysAndValues func(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues int) CFErrorRef
var _cFErrorCreateWithUserInfoKeysAndValuesErr error

func tryCFErrorCreateWithUserInfoKeysAndValues(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues int) (CFErrorRef, error) {
	if _cFErrorCreateWithUserInfoKeysAndValues == nil {
		return 0, symbolCallError("CFErrorCreateWithUserInfoKeysAndValues", "10.5", _cFErrorCreateWithUserInfoKeysAndValuesErr)
	}
	return _cFErrorCreateWithUserInfoKeysAndValues(allocator, domain, code, userInfoKeys, userInfoValues, numUserInfoValues), nil
}

// CFErrorCreateWithUserInfoKeysAndValues creates a new CFError object using given keys and values to create the user info dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreateWithUserInfoKeysAndValues(_:_:_:_:_:_:)
func CFErrorCreateWithUserInfoKeysAndValues(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues int) CFErrorRef {
	result, callErr := tryCFErrorCreateWithUserInfoKeysAndValues(allocator, domain, code, userInfoKeys, userInfoValues, numUserInfoValues)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorGetCode func(err CFErrorRef) int
var _cFErrorGetCodeErr error

func tryCFErrorGetCode(err CFErrorRef) (int, error) {
	if _cFErrorGetCode == nil {
		return 0, symbolCallError("CFErrorGetCode", "10.5", _cFErrorGetCodeErr)
	}
	return _cFErrorGetCode(err), nil
}

// CFErrorGetCode returns the error code for a given CFError.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetCode(_:)
func CFErrorGetCode(err CFErrorRef) int {
	result, callErr := tryCFErrorGetCode(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorGetDomain func(err CFErrorRef) CFErrorDomain
var _cFErrorGetDomainErr error

func tryCFErrorGetDomain(err CFErrorRef) (CFErrorDomain, error) {
	if _cFErrorGetDomain == nil {
		return *new(CFErrorDomain), symbolCallError("CFErrorGetDomain", "10.5", _cFErrorGetDomainErr)
	}
	return _cFErrorGetDomain(err), nil
}

// CFErrorGetDomain returns the error domain for a given CFError.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetDomain(_:)
func CFErrorGetDomain(err CFErrorRef) CFErrorDomain {
	result, callErr := tryCFErrorGetDomain(err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFErrorGetTypeID func() uint
var _cFErrorGetTypeIDErr error

func tryCFErrorGetTypeID() (uint, error) {
	if _cFErrorGetTypeID == nil {
		return 0, symbolCallError("CFErrorGetTypeID", "10.5", _cFErrorGetTypeIDErr)
	}
	return _cFErrorGetTypeID(), nil
}

// CFErrorGetTypeID returns the type identifier for the CFError opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetTypeID()
func CFErrorGetTypeID() uint {
	result, callErr := tryCFErrorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileDescriptorCreate func(allocator CFAllocatorRef, fd CFFileDescriptorNativeDescriptor, closeOnInvalidate bool, callout CFFileDescriptorCallBack, context *CFFileDescriptorContext) CFFileDescriptorRef
var _cFFileDescriptorCreateErr error

func tryCFFileDescriptorCreate(allocator CFAllocatorRef, fd CFFileDescriptorNativeDescriptor, closeOnInvalidate bool, callout CFFileDescriptorCallBack, context *CFFileDescriptorContext) (CFFileDescriptorRef, error) {
	if _cFFileDescriptorCreate == nil {
		return 0, symbolCallError("CFFileDescriptorCreate", "10.5", _cFFileDescriptorCreateErr)
	}
	return _cFFileDescriptorCreate(allocator, fd, closeOnInvalidate, callout, context), nil
}

// CFFileDescriptorCreate creates a new CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreate(_:_:_:_:_:)
func CFFileDescriptorCreate(allocator CFAllocatorRef, fd CFFileDescriptorNativeDescriptor, closeOnInvalidate bool, callout CFFileDescriptorCallBack, context *CFFileDescriptorContext) CFFileDescriptorRef {
	result, callErr := tryCFFileDescriptorCreate(allocator, fd, closeOnInvalidate, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileDescriptorCreateRunLoopSource func(allocator CFAllocatorRef, f CFFileDescriptorRef, order int) CFRunLoopSourceRef
var _cFFileDescriptorCreateRunLoopSourceErr error

func tryCFFileDescriptorCreateRunLoopSource(allocator CFAllocatorRef, f CFFileDescriptorRef, order int) (CFRunLoopSourceRef, error) {
	if _cFFileDescriptorCreateRunLoopSource == nil {
		return 0, symbolCallError("CFFileDescriptorCreateRunLoopSource", "10.5", _cFFileDescriptorCreateRunLoopSourceErr)
	}
	return _cFFileDescriptorCreateRunLoopSource(allocator, f, order), nil
}

// CFFileDescriptorCreateRunLoopSource creates a new runloop source for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreateRunLoopSource(_:_:_:)
func CFFileDescriptorCreateRunLoopSource(allocator CFAllocatorRef, f CFFileDescriptorRef, order int) CFRunLoopSourceRef {
	result, callErr := tryCFFileDescriptorCreateRunLoopSource(allocator, f, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileDescriptorDisableCallBacks func(f CFFileDescriptorRef, callBackTypes uint64)
var _cFFileDescriptorDisableCallBacksErr error

func tryCFFileDescriptorDisableCallBacks(f CFFileDescriptorRef, callBackTypes uint64) error {
	if _cFFileDescriptorDisableCallBacks == nil {
		return symbolCallError("CFFileDescriptorDisableCallBacks", "10.5", _cFFileDescriptorDisableCallBacksErr)
	}
	_cFFileDescriptorDisableCallBacks(f, callBackTypes)
	return nil
}

// CFFileDescriptorDisableCallBacks disables callbacks for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorDisableCallBacks(_:_:)
func CFFileDescriptorDisableCallBacks(f CFFileDescriptorRef, callBackTypes uint64) {
	if callErr := tryCFFileDescriptorDisableCallBacks(f, callBackTypes); callErr != nil {
		panic(callErr)
	}
}

var _cFFileDescriptorEnableCallBacks func(f CFFileDescriptorRef, callBackTypes uint64)
var _cFFileDescriptorEnableCallBacksErr error

func tryCFFileDescriptorEnableCallBacks(f CFFileDescriptorRef, callBackTypes uint64) error {
	if _cFFileDescriptorEnableCallBacks == nil {
		return symbolCallError("CFFileDescriptorEnableCallBacks", "10.5", _cFFileDescriptorEnableCallBacksErr)
	}
	_cFFileDescriptorEnableCallBacks(f, callBackTypes)
	return nil
}

// CFFileDescriptorEnableCallBacks enables callbacks for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorEnableCallBacks(_:_:)
func CFFileDescriptorEnableCallBacks(f CFFileDescriptorRef, callBackTypes uint64) {
	if callErr := tryCFFileDescriptorEnableCallBacks(f, callBackTypes); callErr != nil {
		panic(callErr)
	}
}

var _cFFileDescriptorGetContext func(f CFFileDescriptorRef, context *CFFileDescriptorContext)
var _cFFileDescriptorGetContextErr error

func tryCFFileDescriptorGetContext(f CFFileDescriptorRef, context *CFFileDescriptorContext) error {
	if _cFFileDescriptorGetContext == nil {
		return symbolCallError("CFFileDescriptorGetContext", "10.5", _cFFileDescriptorGetContextErr)
	}
	_cFFileDescriptorGetContext(f, context)
	return nil
}

// CFFileDescriptorGetContext gets the context for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetContext(_:_:)
func CFFileDescriptorGetContext(f CFFileDescriptorRef, context *CFFileDescriptorContext) {
	if callErr := tryCFFileDescriptorGetContext(f, context); callErr != nil {
		panic(callErr)
	}
}

var _cFFileDescriptorGetNativeDescriptor func(f CFFileDescriptorRef) CFFileDescriptorNativeDescriptor
var _cFFileDescriptorGetNativeDescriptorErr error

func tryCFFileDescriptorGetNativeDescriptor(f CFFileDescriptorRef) (CFFileDescriptorNativeDescriptor, error) {
	if _cFFileDescriptorGetNativeDescriptor == nil {
		return *new(CFFileDescriptorNativeDescriptor), symbolCallError("CFFileDescriptorGetNativeDescriptor", "10.5", _cFFileDescriptorGetNativeDescriptorErr)
	}
	return _cFFileDescriptorGetNativeDescriptor(f), nil
}

// CFFileDescriptorGetNativeDescriptor returns the native file descriptor for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetNativeDescriptor(_:)
func CFFileDescriptorGetNativeDescriptor(f CFFileDescriptorRef) CFFileDescriptorNativeDescriptor {
	result, callErr := tryCFFileDescriptorGetNativeDescriptor(f)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileDescriptorGetTypeID func() uint
var _cFFileDescriptorGetTypeIDErr error

func tryCFFileDescriptorGetTypeID() (uint, error) {
	if _cFFileDescriptorGetTypeID == nil {
		return 0, symbolCallError("CFFileDescriptorGetTypeID", "10.5", _cFFileDescriptorGetTypeIDErr)
	}
	return _cFFileDescriptorGetTypeID(), nil
}

// CFFileDescriptorGetTypeID returns the type identifier for the CFFileDescriptor opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetTypeID()
func CFFileDescriptorGetTypeID() uint {
	result, callErr := tryCFFileDescriptorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileDescriptorInvalidate func(f CFFileDescriptorRef)
var _cFFileDescriptorInvalidateErr error

func tryCFFileDescriptorInvalidate(f CFFileDescriptorRef) error {
	if _cFFileDescriptorInvalidate == nil {
		return symbolCallError("CFFileDescriptorInvalidate", "10.5", _cFFileDescriptorInvalidateErr)
	}
	_cFFileDescriptorInvalidate(f)
	return nil
}

// CFFileDescriptorInvalidate invalidates a CFFileDescriptor object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorInvalidate(_:)
func CFFileDescriptorInvalidate(f CFFileDescriptorRef) {
	if callErr := tryCFFileDescriptorInvalidate(f); callErr != nil {
		panic(callErr)
	}
}

var _cFFileDescriptorIsValid func(f CFFileDescriptorRef) bool
var _cFFileDescriptorIsValidErr error

func tryCFFileDescriptorIsValid(f CFFileDescriptorRef) (bool, error) {
	if _cFFileDescriptorIsValid == nil {
		return false, symbolCallError("CFFileDescriptorIsValid", "10.5", _cFFileDescriptorIsValidErr)
	}
	return _cFFileDescriptorIsValid(f), nil
}

// CFFileDescriptorIsValid returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorIsValid(_:)
func CFFileDescriptorIsValid(f CFFileDescriptorRef) bool {
	result, callErr := tryCFFileDescriptorIsValid(f)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityClearProperties func(fileSec CFFileSecurityRef, clearPropertyMask CFFileSecurityClearOptions) bool
var _cFFileSecurityClearPropertiesErr error

func tryCFFileSecurityClearProperties(fileSec CFFileSecurityRef, clearPropertyMask CFFileSecurityClearOptions) (bool, error) {
	if _cFFileSecurityClearProperties == nil {
		return false, symbolCallError("CFFileSecurityClearProperties", "10.8", _cFFileSecurityClearPropertiesErr)
	}
	return _cFFileSecurityClearProperties(fileSec, clearPropertyMask), nil
}

// CFFileSecurityClearProperties clears properties from a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearProperties(_:_:)
func CFFileSecurityClearProperties(fileSec CFFileSecurityRef, clearPropertyMask CFFileSecurityClearOptions) bool {
	result, callErr := tryCFFileSecurityClearProperties(fileSec, clearPropertyMask)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityCopyAccessControlList func(fileSec CFFileSecurityRef, accessControlList uintptr) bool
var _cFFileSecurityCopyAccessControlListErr error

func tryCFFileSecurityCopyAccessControlList(fileSec CFFileSecurityRef, accessControlList uintptr) (bool, error) {
	if _cFFileSecurityCopyAccessControlList == nil {
		return false, symbolCallError("CFFileSecurityCopyAccessControlList", "10.7", _cFFileSecurityCopyAccessControlListErr)
	}
	return _cFFileSecurityCopyAccessControlList(fileSec, accessControlList), nil
}

// CFFileSecurityCopyAccessControlList copies the access control list associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyAccessControlList(_:_:)
func CFFileSecurityCopyAccessControlList(fileSec CFFileSecurityRef, accessControlList uintptr) bool {
	result, callErr := tryCFFileSecurityCopyAccessControlList(fileSec, accessControlList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityCopyGroupUUID func(fileSec CFFileSecurityRef, groupUUID *CFUUIDRef) bool
var _cFFileSecurityCopyGroupUUIDErr error

func tryCFFileSecurityCopyGroupUUID(fileSec CFFileSecurityRef, groupUUID *CFUUIDRef) (bool, error) {
	if _cFFileSecurityCopyGroupUUID == nil {
		return false, symbolCallError("CFFileSecurityCopyGroupUUID", "10.7", _cFFileSecurityCopyGroupUUIDErr)
	}
	return _cFFileSecurityCopyGroupUUID(fileSec, groupUUID), nil
}

// CFFileSecurityCopyGroupUUID copies the group UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyGroupUUID(_:_:)
func CFFileSecurityCopyGroupUUID(fileSec CFFileSecurityRef, groupUUID *CFUUIDRef) bool {
	result, callErr := tryCFFileSecurityCopyGroupUUID(fileSec, groupUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityCopyOwnerUUID func(fileSec CFFileSecurityRef, ownerUUID *CFUUIDRef) bool
var _cFFileSecurityCopyOwnerUUIDErr error

func tryCFFileSecurityCopyOwnerUUID(fileSec CFFileSecurityRef, ownerUUID *CFUUIDRef) (bool, error) {
	if _cFFileSecurityCopyOwnerUUID == nil {
		return false, symbolCallError("CFFileSecurityCopyOwnerUUID", "10.7", _cFFileSecurityCopyOwnerUUIDErr)
	}
	return _cFFileSecurityCopyOwnerUUID(fileSec, ownerUUID), nil
}

// CFFileSecurityCopyOwnerUUID copies the owner UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyOwnerUUID(_:_:)
func CFFileSecurityCopyOwnerUUID(fileSec CFFileSecurityRef, ownerUUID *CFUUIDRef) bool {
	result, callErr := tryCFFileSecurityCopyOwnerUUID(fileSec, ownerUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityCreate func(allocator CFAllocatorRef) CFFileSecurityRef
var _cFFileSecurityCreateErr error

func tryCFFileSecurityCreate(allocator CFAllocatorRef) (CFFileSecurityRef, error) {
	if _cFFileSecurityCreate == nil {
		return 0, symbolCallError("CFFileSecurityCreate", "10.7", _cFFileSecurityCreateErr)
	}
	return _cFFileSecurityCreate(allocator), nil
}

// CFFileSecurityCreate creates a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreate(_:)
func CFFileSecurityCreate(allocator CFAllocatorRef) CFFileSecurityRef {
	result, callErr := tryCFFileSecurityCreate(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityCreateCopy func(allocator CFAllocatorRef, fileSec CFFileSecurityRef) CFFileSecurityRef
var _cFFileSecurityCreateCopyErr error

func tryCFFileSecurityCreateCopy(allocator CFAllocatorRef, fileSec CFFileSecurityRef) (CFFileSecurityRef, error) {
	if _cFFileSecurityCreateCopy == nil {
		return 0, symbolCallError("CFFileSecurityCreateCopy", "10.7", _cFFileSecurityCreateCopyErr)
	}
	return _cFFileSecurityCreateCopy(allocator, fileSec), nil
}

// CFFileSecurityCreateCopy creates a copy of a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreateCopy(_:_:)
func CFFileSecurityCreateCopy(allocator CFAllocatorRef, fileSec CFFileSecurityRef) CFFileSecurityRef {
	result, callErr := tryCFFileSecurityCreateCopy(allocator, fileSec)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityGetGroup func(fileSec CFFileSecurityRef, group *uint32) bool
var _cFFileSecurityGetGroupErr error

func tryCFFileSecurityGetGroup(fileSec CFFileSecurityRef, group *uint32) (bool, error) {
	if _cFFileSecurityGetGroup == nil {
		return false, symbolCallError("CFFileSecurityGetGroup", "10.7", _cFFileSecurityGetGroupErr)
	}
	return _cFFileSecurityGetGroup(fileSec, group), nil
}

// CFFileSecurityGetGroup gets the group ID associated with a [CFFileSecurityRef] object
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetGroup(_:_:)
func CFFileSecurityGetGroup(fileSec CFFileSecurityRef, group *uint32) bool {
	result, callErr := tryCFFileSecurityGetGroup(fileSec, group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityGetMode func(fileSec CFFileSecurityRef, mode *uint16) bool
var _cFFileSecurityGetModeErr error

func tryCFFileSecurityGetMode(fileSec CFFileSecurityRef, mode *uint16) (bool, error) {
	if _cFFileSecurityGetMode == nil {
		return false, symbolCallError("CFFileSecurityGetMode", "10.7", _cFFileSecurityGetModeErr)
	}
	return _cFFileSecurityGetMode(fileSec, mode), nil
}

// CFFileSecurityGetMode gets the file mode associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetMode(_:_:)
func CFFileSecurityGetMode(fileSec CFFileSecurityRef, mode *uint16) bool {
	result, callErr := tryCFFileSecurityGetMode(fileSec, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityGetOwner func(fileSec CFFileSecurityRef, owner *uint32) bool
var _cFFileSecurityGetOwnerErr error

func tryCFFileSecurityGetOwner(fileSec CFFileSecurityRef, owner *uint32) (bool, error) {
	if _cFFileSecurityGetOwner == nil {
		return false, symbolCallError("CFFileSecurityGetOwner", "10.7", _cFFileSecurityGetOwnerErr)
	}
	return _cFFileSecurityGetOwner(fileSec, owner), nil
}

// CFFileSecurityGetOwner gets the owner ID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetOwner(_:_:)
func CFFileSecurityGetOwner(fileSec CFFileSecurityRef, owner *uint32) bool {
	result, callErr := tryCFFileSecurityGetOwner(fileSec, owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecurityGetTypeID func() uint
var _cFFileSecurityGetTypeIDErr error

func tryCFFileSecurityGetTypeID() (uint, error) {
	if _cFFileSecurityGetTypeID == nil {
		return 0, symbolCallError("CFFileSecurityGetTypeID", "10.7", _cFFileSecurityGetTypeIDErr)
	}
	return _cFFileSecurityGetTypeID(), nil
}

// CFFileSecurityGetTypeID returns the type identifier for the [CFFileSecurityRef] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetTypeID()
func CFFileSecurityGetTypeID() uint {
	result, callErr := tryCFFileSecurityGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecuritySetAccessControlList func(fileSec CFFileSecurityRef, accessControlList unsafe.Pointer) bool
var _cFFileSecuritySetAccessControlListErr error

func tryCFFileSecuritySetAccessControlList(fileSec CFFileSecurityRef, accessControlList unsafe.Pointer) (bool, error) {
	if _cFFileSecuritySetAccessControlList == nil {
		return false, symbolCallError("CFFileSecuritySetAccessControlList", "10.7", _cFFileSecuritySetAccessControlListErr)
	}
	return _cFFileSecuritySetAccessControlList(fileSec, accessControlList), nil
}

// CFFileSecuritySetAccessControlList sets the access control list associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetAccessControlList(_:_:)
func CFFileSecuritySetAccessControlList(fileSec CFFileSecurityRef, accessControlList unsafe.Pointer) bool {
	result, callErr := tryCFFileSecuritySetAccessControlList(fileSec, accessControlList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecuritySetGroup func(fileSec CFFileSecurityRef, group uint32) bool
var _cFFileSecuritySetGroupErr error

func tryCFFileSecuritySetGroup(fileSec CFFileSecurityRef, group uint32) (bool, error) {
	if _cFFileSecuritySetGroup == nil {
		return false, symbolCallError("CFFileSecuritySetGroup", "10.7", _cFFileSecuritySetGroupErr)
	}
	return _cFFileSecuritySetGroup(fileSec, group), nil
}

// CFFileSecuritySetGroup sets the group ID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroup(_:_:)
func CFFileSecuritySetGroup(fileSec CFFileSecurityRef, group uint32) bool {
	result, callErr := tryCFFileSecuritySetGroup(fileSec, group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecuritySetGroupUUID func(fileSec CFFileSecurityRef, groupUUID CFUUIDRef) bool
var _cFFileSecuritySetGroupUUIDErr error

func tryCFFileSecuritySetGroupUUID(fileSec CFFileSecurityRef, groupUUID CFUUIDRef) (bool, error) {
	if _cFFileSecuritySetGroupUUID == nil {
		return false, symbolCallError("CFFileSecuritySetGroupUUID", "10.7", _cFFileSecuritySetGroupUUIDErr)
	}
	return _cFFileSecuritySetGroupUUID(fileSec, groupUUID), nil
}

// CFFileSecuritySetGroupUUID sets the group UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroupUUID(_:_:)
func CFFileSecuritySetGroupUUID(fileSec CFFileSecurityRef, groupUUID CFUUIDRef) bool {
	result, callErr := tryCFFileSecuritySetGroupUUID(fileSec, groupUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecuritySetMode func(fileSec CFFileSecurityRef, mode uint16) bool
var _cFFileSecuritySetModeErr error

func tryCFFileSecuritySetMode(fileSec CFFileSecurityRef, mode uint16) (bool, error) {
	if _cFFileSecuritySetMode == nil {
		return false, symbolCallError("CFFileSecuritySetMode", "10.7", _cFFileSecuritySetModeErr)
	}
	return _cFFileSecuritySetMode(fileSec, mode), nil
}

// CFFileSecuritySetMode sets the file mode associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetMode(_:_:)
func CFFileSecuritySetMode(fileSec CFFileSecurityRef, mode uint16) bool {
	result, callErr := tryCFFileSecuritySetMode(fileSec, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecuritySetOwner func(fileSec CFFileSecurityRef, owner uint32) bool
var _cFFileSecuritySetOwnerErr error

func tryCFFileSecuritySetOwner(fileSec CFFileSecurityRef, owner uint32) (bool, error) {
	if _cFFileSecuritySetOwner == nil {
		return false, symbolCallError("CFFileSecuritySetOwner", "10.7", _cFFileSecuritySetOwnerErr)
	}
	return _cFFileSecuritySetOwner(fileSec, owner), nil
}

// CFFileSecuritySetOwner sets the owner ID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwner(_:_:)
func CFFileSecuritySetOwner(fileSec CFFileSecurityRef, owner uint32) bool {
	result, callErr := tryCFFileSecuritySetOwner(fileSec, owner)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFFileSecuritySetOwnerUUID func(fileSec CFFileSecurityRef, ownerUUID CFUUIDRef) bool
var _cFFileSecuritySetOwnerUUIDErr error

func tryCFFileSecuritySetOwnerUUID(fileSec CFFileSecurityRef, ownerUUID CFUUIDRef) (bool, error) {
	if _cFFileSecuritySetOwnerUUID == nil {
		return false, symbolCallError("CFFileSecuritySetOwnerUUID", "10.7", _cFFileSecuritySetOwnerUUIDErr)
	}
	return _cFFileSecuritySetOwnerUUID(fileSec, ownerUUID), nil
}

// CFFileSecuritySetOwnerUUID sets the owner UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwnerUUID(_:_:)
func CFFileSecuritySetOwnerUUID(fileSec CFFileSecurityRef, ownerUUID CFUUIDRef) bool {
	result, callErr := tryCFFileSecuritySetOwnerUUID(fileSec, ownerUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFGetAllocator func(cf CFTypeRef) CFAllocatorRef
var _cFGetAllocatorErr error

func tryCFGetAllocator(cf CFTypeRef) (CFAllocatorRef, error) {
	if _cFGetAllocator == nil {
		return 0, symbolCallError("CFGetAllocator", "", _cFGetAllocatorErr)
	}
	return _cFGetAllocator(cf), nil
}

// CFGetAllocator returns the allocator used to allocate a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFGetAllocator(_:)
func CFGetAllocator(cf CFTypeRef) CFAllocatorRef {
	result, callErr := tryCFGetAllocator(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFGetRetainCount func(cf CFTypeRef) int
var _cFGetRetainCountErr error

func tryCFGetRetainCount(cf CFTypeRef) (int, error) {
	if _cFGetRetainCount == nil {
		return 0, symbolCallError("CFGetRetainCount", "", _cFGetRetainCountErr)
	}
	return _cFGetRetainCount(cf), nil
}

// CFGetRetainCount returns the reference count of a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFGetRetainCount(_:)
func CFGetRetainCount(cf CFTypeRef) int {
	result, callErr := tryCFGetRetainCount(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFGetTypeID func(cf CFTypeRef) uint
var _cFGetTypeIDErr error

func tryCFGetTypeID(cf CFTypeRef) (uint, error) {
	if _cFGetTypeID == nil {
		return 0, symbolCallError("CFGetTypeID", "", _cFGetTypeIDErr)
	}
	return _cFGetTypeID(cf), nil
}

// CFGetTypeID returns the unique identifier of an opaque type to which a Core Foundation object belongs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFGetTypeID(_:)
func CFGetTypeID(cf CFTypeRef) uint {
	result, callErr := tryCFGetTypeID(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFHash func(cf CFTypeRef) uint
var _cFHashErr error

func tryCFHash(cf CFTypeRef) (uint, error) {
	if _cFHash == nil {
		return 0, symbolCallError("CFHash", "", _cFHashErr)
	}
	return _cFHash(cf), nil
}

// CFHash returns a code that can be used to identify an object in a hashing structure.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFHash(_:)
func CFHash(cf CFTypeRef) uint {
	result, callErr := tryCFHash(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyAvailableLocaleIdentifiers func() CFArrayRef
var _cFLocaleCopyAvailableLocaleIdentifiersErr error

func tryCFLocaleCopyAvailableLocaleIdentifiers() (CFArrayRef, error) {
	if _cFLocaleCopyAvailableLocaleIdentifiers == nil {
		return 0, symbolCallError("CFLocaleCopyAvailableLocaleIdentifiers", "", _cFLocaleCopyAvailableLocaleIdentifiersErr)
	}
	return _cFLocaleCopyAvailableLocaleIdentifiers(), nil
}

// CFLocaleCopyAvailableLocaleIdentifiers returns an array of CFString objects that represents all locales for which locale data is available.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyAvailableLocaleIdentifiers()
func CFLocaleCopyAvailableLocaleIdentifiers() CFArrayRef {
	result, callErr := tryCFLocaleCopyAvailableLocaleIdentifiers()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyCommonISOCurrencyCodes func() CFArrayRef
var _cFLocaleCopyCommonISOCurrencyCodesErr error

func tryCFLocaleCopyCommonISOCurrencyCodes() (CFArrayRef, error) {
	if _cFLocaleCopyCommonISOCurrencyCodes == nil {
		return 0, symbolCallError("CFLocaleCopyCommonISOCurrencyCodes", "10.5", _cFLocaleCopyCommonISOCurrencyCodesErr)
	}
	return _cFLocaleCopyCommonISOCurrencyCodes(), nil
}

// CFLocaleCopyCommonISOCurrencyCodes returns an array of strings that represents ISO currency codes for currencies in common use.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCommonISOCurrencyCodes()
func CFLocaleCopyCommonISOCurrencyCodes() CFArrayRef {
	result, callErr := tryCFLocaleCopyCommonISOCurrencyCodes()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyCurrent func() CFLocaleRef
var _cFLocaleCopyCurrentErr error

func tryCFLocaleCopyCurrent() (CFLocaleRef, error) {
	if _cFLocaleCopyCurrent == nil {
		return 0, symbolCallError("CFLocaleCopyCurrent", "", _cFLocaleCopyCurrentErr)
	}
	return _cFLocaleCopyCurrent(), nil
}

// CFLocaleCopyCurrent returns a copy of the logical locale for the current user.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCurrent()
func CFLocaleCopyCurrent() CFLocaleRef {
	result, callErr := tryCFLocaleCopyCurrent()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyDisplayNameForPropertyValue func(displayLocale CFLocaleRef, key CFLocaleKey, value CFStringRef) CFStringRef
var _cFLocaleCopyDisplayNameForPropertyValueErr error

func tryCFLocaleCopyDisplayNameForPropertyValue(displayLocale CFLocaleRef, key CFLocaleKey, value CFStringRef) (CFStringRef, error) {
	if _cFLocaleCopyDisplayNameForPropertyValue == nil {
		return 0, symbolCallError("CFLocaleCopyDisplayNameForPropertyValue", "", _cFLocaleCopyDisplayNameForPropertyValueErr)
	}
	return _cFLocaleCopyDisplayNameForPropertyValue(displayLocale, key, value), nil
}

// CFLocaleCopyDisplayNameForPropertyValue returns the display name for the given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyDisplayNameForPropertyValue(_:_:_:)
func CFLocaleCopyDisplayNameForPropertyValue(displayLocale CFLocaleRef, key CFLocaleKey, value CFStringRef) CFStringRef {
	result, callErr := tryCFLocaleCopyDisplayNameForPropertyValue(displayLocale, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyISOCountryCodes func() CFArrayRef
var _cFLocaleCopyISOCountryCodesErr error

func tryCFLocaleCopyISOCountryCodes() (CFArrayRef, error) {
	if _cFLocaleCopyISOCountryCodes == nil {
		return 0, symbolCallError("CFLocaleCopyISOCountryCodes", "", _cFLocaleCopyISOCountryCodesErr)
	}
	return _cFLocaleCopyISOCountryCodes(), nil
}

// CFLocaleCopyISOCountryCodes returns an array of CFString objects that represents all known legal ISO country codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCountryCodes()
func CFLocaleCopyISOCountryCodes() CFArrayRef {
	result, callErr := tryCFLocaleCopyISOCountryCodes()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyISOCurrencyCodes func() CFArrayRef
var _cFLocaleCopyISOCurrencyCodesErr error

func tryCFLocaleCopyISOCurrencyCodes() (CFArrayRef, error) {
	if _cFLocaleCopyISOCurrencyCodes == nil {
		return 0, symbolCallError("CFLocaleCopyISOCurrencyCodes", "", _cFLocaleCopyISOCurrencyCodesErr)
	}
	return _cFLocaleCopyISOCurrencyCodes(), nil
}

// CFLocaleCopyISOCurrencyCodes returns an array of CFString objects that represents all known legal ISO currency codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCurrencyCodes()
func CFLocaleCopyISOCurrencyCodes() CFArrayRef {
	result, callErr := tryCFLocaleCopyISOCurrencyCodes()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyISOLanguageCodes func() CFArrayRef
var _cFLocaleCopyISOLanguageCodesErr error

func tryCFLocaleCopyISOLanguageCodes() (CFArrayRef, error) {
	if _cFLocaleCopyISOLanguageCodes == nil {
		return 0, symbolCallError("CFLocaleCopyISOLanguageCodes", "", _cFLocaleCopyISOLanguageCodesErr)
	}
	return _cFLocaleCopyISOLanguageCodes(), nil
}

// CFLocaleCopyISOLanguageCodes returns an array of CFString objects that represents all known legal ISO language codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOLanguageCodes()
func CFLocaleCopyISOLanguageCodes() CFArrayRef {
	result, callErr := tryCFLocaleCopyISOLanguageCodes()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCopyPreferredLanguages func() CFArrayRef
var _cFLocaleCopyPreferredLanguagesErr error

func tryCFLocaleCopyPreferredLanguages() (CFArrayRef, error) {
	if _cFLocaleCopyPreferredLanguages == nil {
		return 0, symbolCallError("CFLocaleCopyPreferredLanguages", "10.5", _cFLocaleCopyPreferredLanguagesErr)
	}
	return _cFLocaleCopyPreferredLanguages(), nil
}

// CFLocaleCopyPreferredLanguages returns the array of canonicalized language IDs that the user prefers.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyPreferredLanguages()
func CFLocaleCopyPreferredLanguages() CFArrayRef {
	result, callErr := tryCFLocaleCopyPreferredLanguages()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreate func(allocator CFAllocatorRef, localeIdentifier CFLocaleIdentifier) CFLocaleRef
var _cFLocaleCreateErr error

func tryCFLocaleCreate(allocator CFAllocatorRef, localeIdentifier CFLocaleIdentifier) (CFLocaleRef, error) {
	if _cFLocaleCreate == nil {
		return 0, symbolCallError("CFLocaleCreate", "", _cFLocaleCreateErr)
	}
	return _cFLocaleCreate(allocator, localeIdentifier), nil
}

// CFLocaleCreate creates a locale for the given arbitrary locale identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreate(_:_:)
func CFLocaleCreate(allocator CFAllocatorRef, localeIdentifier CFLocaleIdentifier) CFLocaleRef {
	result, callErr := tryCFLocaleCreate(allocator, localeIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateCanonicalLanguageIdentifierFromString func(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier
var _cFLocaleCreateCanonicalLanguageIdentifierFromStringErr error

func tryCFLocaleCreateCanonicalLanguageIdentifierFromString(allocator CFAllocatorRef, localeIdentifier CFStringRef) (CFLocaleIdentifier, error) {
	if _cFLocaleCreateCanonicalLanguageIdentifierFromString == nil {
		return *new(CFLocaleIdentifier), symbolCallError("CFLocaleCreateCanonicalLanguageIdentifierFromString", "", _cFLocaleCreateCanonicalLanguageIdentifierFromStringErr)
	}
	return _cFLocaleCreateCanonicalLanguageIdentifierFromString(allocator, localeIdentifier), nil
}

// CFLocaleCreateCanonicalLanguageIdentifierFromString returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLanguageIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier {
	result, callErr := tryCFLocaleCreateCanonicalLanguageIdentifierFromString(allocator, localeIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes func(allocator CFAllocatorRef, lcode unsafe.Pointer, rcode unsafe.Pointer) CFLocaleIdentifier
var _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodesErr error

func tryCFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator CFAllocatorRef, lcode unsafe.Pointer, rcode unsafe.Pointer) (CFLocaleIdentifier, error) {
	if _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes == nil {
		return *new(CFLocaleIdentifier), symbolCallError("CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes", "", _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodesErr)
	}
	return _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator, lcode, rcode), nil
}

// CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes returns a canonical locale identifier from given language and region codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(_:_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator CFAllocatorRef, lcode unsafe.Pointer, rcode unsafe.Pointer) CFLocaleIdentifier {
	result, callErr := tryCFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator, lcode, rcode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateCanonicalLocaleIdentifierFromString func(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier
var _cFLocaleCreateCanonicalLocaleIdentifierFromStringErr error

func tryCFLocaleCreateCanonicalLocaleIdentifierFromString(allocator CFAllocatorRef, localeIdentifier CFStringRef) (CFLocaleIdentifier, error) {
	if _cFLocaleCreateCanonicalLocaleIdentifierFromString == nil {
		return *new(CFLocaleIdentifier), symbolCallError("CFLocaleCreateCanonicalLocaleIdentifierFromString", "", _cFLocaleCreateCanonicalLocaleIdentifierFromStringErr)
	}
	return _cFLocaleCreateCanonicalLocaleIdentifierFromString(allocator, localeIdentifier), nil
}

// CFLocaleCreateCanonicalLocaleIdentifierFromString returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier {
	result, callErr := tryCFLocaleCreateCanonicalLocaleIdentifierFromString(allocator, localeIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateComponentsFromLocaleIdentifier func(allocator CFAllocatorRef, localeID CFLocaleIdentifier) CFDictionaryRef
var _cFLocaleCreateComponentsFromLocaleIdentifierErr error

func tryCFLocaleCreateComponentsFromLocaleIdentifier(allocator CFAllocatorRef, localeID CFLocaleIdentifier) (CFDictionaryRef, error) {
	if _cFLocaleCreateComponentsFromLocaleIdentifier == nil {
		return 0, symbolCallError("CFLocaleCreateComponentsFromLocaleIdentifier", "", _cFLocaleCreateComponentsFromLocaleIdentifierErr)
	}
	return _cFLocaleCreateComponentsFromLocaleIdentifier(allocator, localeID), nil
}

// CFLocaleCreateComponentsFromLocaleIdentifier returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateComponentsFromLocaleIdentifier(_:_:)
func CFLocaleCreateComponentsFromLocaleIdentifier(allocator CFAllocatorRef, localeID CFLocaleIdentifier) CFDictionaryRef {
	result, callErr := tryCFLocaleCreateComponentsFromLocaleIdentifier(allocator, localeID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateCopy func(allocator CFAllocatorRef, locale CFLocaleRef) CFLocaleRef
var _cFLocaleCreateCopyErr error

func tryCFLocaleCreateCopy(allocator CFAllocatorRef, locale CFLocaleRef) (CFLocaleRef, error) {
	if _cFLocaleCreateCopy == nil {
		return 0, symbolCallError("CFLocaleCreateCopy", "", _cFLocaleCreateCopyErr)
	}
	return _cFLocaleCreateCopy(allocator, locale), nil
}

// CFLocaleCreateCopy returns a copy of a locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCopy(_:_:)
func CFLocaleCreateCopy(allocator CFAllocatorRef, locale CFLocaleRef) CFLocaleRef {
	result, callErr := tryCFLocaleCreateCopy(allocator, locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateLocaleIdentifierFromComponents func(allocator CFAllocatorRef, dictionary CFDictionaryRef) CFLocaleIdentifier
var _cFLocaleCreateLocaleIdentifierFromComponentsErr error

func tryCFLocaleCreateLocaleIdentifierFromComponents(allocator CFAllocatorRef, dictionary CFDictionaryRef) (CFLocaleIdentifier, error) {
	if _cFLocaleCreateLocaleIdentifierFromComponents == nil {
		return *new(CFLocaleIdentifier), symbolCallError("CFLocaleCreateLocaleIdentifierFromComponents", "", _cFLocaleCreateLocaleIdentifierFromComponentsErr)
	}
	return _cFLocaleCreateLocaleIdentifierFromComponents(allocator, dictionary), nil
}

// CFLocaleCreateLocaleIdentifierFromComponents returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromComponents(_:_:)
func CFLocaleCreateLocaleIdentifierFromComponents(allocator CFAllocatorRef, dictionary CFDictionaryRef) CFLocaleIdentifier {
	result, callErr := tryCFLocaleCreateLocaleIdentifierFromComponents(allocator, dictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode func(allocator CFAllocatorRef, lcid uint32) CFLocaleIdentifier
var _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCodeErr error

func tryCFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator CFAllocatorRef, lcid uint32) (CFLocaleIdentifier, error) {
	if _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode == nil {
		return *new(CFLocaleIdentifier), symbolCallError("CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode", "10.6", _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCodeErr)
	}
	return _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator, lcid), nil
}

// CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode returns a locale identifier from a Windows locale code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(_:_:)
func CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator CFAllocatorRef, lcid uint32) CFLocaleIdentifier {
	result, callErr := tryCFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator, lcid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetIdentifier func(locale CFLocaleRef) CFLocaleIdentifier
var _cFLocaleGetIdentifierErr error

func tryCFLocaleGetIdentifier(locale CFLocaleRef) (CFLocaleIdentifier, error) {
	if _cFLocaleGetIdentifier == nil {
		return *new(CFLocaleIdentifier), symbolCallError("CFLocaleGetIdentifier", "", _cFLocaleGetIdentifierErr)
	}
	return _cFLocaleGetIdentifier(locale), nil
}

// CFLocaleGetIdentifier returns the given locale’s identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetIdentifier(_:)
func CFLocaleGetIdentifier(locale CFLocaleRef) CFLocaleIdentifier {
	result, callErr := tryCFLocaleGetIdentifier(locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetLanguageCharacterDirection func(isoLangCode CFStringRef) CFLocaleLanguageDirection
var _cFLocaleGetLanguageCharacterDirectionErr error

func tryCFLocaleGetLanguageCharacterDirection(isoLangCode CFStringRef) (CFLocaleLanguageDirection, error) {
	if _cFLocaleGetLanguageCharacterDirection == nil {
		return *new(CFLocaleLanguageDirection), symbolCallError("CFLocaleGetLanguageCharacterDirection", "10.6", _cFLocaleGetLanguageCharacterDirectionErr)
	}
	return _cFLocaleGetLanguageCharacterDirection(isoLangCode), nil
}

// CFLocaleGetLanguageCharacterDirection returns the character direction for the specified ISO language code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageCharacterDirection(_:)
func CFLocaleGetLanguageCharacterDirection(isoLangCode CFStringRef) CFLocaleLanguageDirection {
	result, callErr := tryCFLocaleGetLanguageCharacterDirection(isoLangCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetLanguageLineDirection func(isoLangCode CFStringRef) CFLocaleLanguageDirection
var _cFLocaleGetLanguageLineDirectionErr error

func tryCFLocaleGetLanguageLineDirection(isoLangCode CFStringRef) (CFLocaleLanguageDirection, error) {
	if _cFLocaleGetLanguageLineDirection == nil {
		return *new(CFLocaleLanguageDirection), symbolCallError("CFLocaleGetLanguageLineDirection", "10.6", _cFLocaleGetLanguageLineDirectionErr)
	}
	return _cFLocaleGetLanguageLineDirection(isoLangCode), nil
}

// CFLocaleGetLanguageLineDirection returns the line direction for the specified ISO language code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageLineDirection(_:)
func CFLocaleGetLanguageLineDirection(isoLangCode CFStringRef) CFLocaleLanguageDirection {
	result, callErr := tryCFLocaleGetLanguageLineDirection(isoLangCode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetSystem func() CFLocaleRef
var _cFLocaleGetSystemErr error

func tryCFLocaleGetSystem() (CFLocaleRef, error) {
	if _cFLocaleGetSystem == nil {
		return 0, symbolCallError("CFLocaleGetSystem", "", _cFLocaleGetSystemErr)
	}
	return _cFLocaleGetSystem(), nil
}

// CFLocaleGetSystem returns the root, canonical locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetSystem()
func CFLocaleGetSystem() CFLocaleRef {
	result, callErr := tryCFLocaleGetSystem()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetTypeID func() uint
var _cFLocaleGetTypeIDErr error

func tryCFLocaleGetTypeID() (uint, error) {
	if _cFLocaleGetTypeID == nil {
		return 0, symbolCallError("CFLocaleGetTypeID", "", _cFLocaleGetTypeIDErr)
	}
	return _cFLocaleGetTypeID(), nil
}

// CFLocaleGetTypeID returns the type identifier for the CFLocale opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetTypeID()
func CFLocaleGetTypeID() uint {
	result, callErr := tryCFLocaleGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetValue func(locale CFLocaleRef, key CFLocaleKey) CFTypeRef
var _cFLocaleGetValueErr error

func tryCFLocaleGetValue(locale CFLocaleRef, key CFLocaleKey) (CFTypeRef, error) {
	if _cFLocaleGetValue == nil {
		return 0, symbolCallError("CFLocaleGetValue", "", _cFLocaleGetValueErr)
	}
	return _cFLocaleGetValue(locale, key), nil
}

// CFLocaleGetValue returns the corresponding value for the given key of a locale’s key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetValue(_:_:)
func CFLocaleGetValue(locale CFLocaleRef, key CFLocaleKey) CFTypeRef {
	result, callErr := tryCFLocaleGetValue(locale, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier func(localeIdentifier CFLocaleIdentifier) uint32
var _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifierErr error

func tryCFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier CFLocaleIdentifier) (uint32, error) {
	if _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier == nil {
		return 0, symbolCallError("CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier", "10.6", _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifierErr)
	}
	return _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier), nil
}

// CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier returns a Windows locale code from the locale identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(_:)
func CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier CFLocaleIdentifier) uint32 {
	result, callErr := tryCFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortCreate func(allocator CFAllocatorRef, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef
var _cFMachPortCreateErr error

func tryCFMachPortCreate(allocator CFAllocatorRef, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) (CFMachPortRef, error) {
	if _cFMachPortCreate == nil {
		return 0, symbolCallError("CFMachPortCreate", "", _cFMachPortCreateErr)
	}
	return _cFMachPortCreate(allocator, callout, context, shouldFreeInfo), nil
}

// CFMachPortCreate creates a CFMachPort object with a new Mach port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreate(_:_:_:_:)
func CFMachPortCreate(allocator CFAllocatorRef, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef {
	result, callErr := tryCFMachPortCreate(allocator, callout, context, shouldFreeInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortCreateRunLoopSource func(allocator CFAllocatorRef, port CFMachPortRef, order int) CFRunLoopSourceRef
var _cFMachPortCreateRunLoopSourceErr error

func tryCFMachPortCreateRunLoopSource(allocator CFAllocatorRef, port CFMachPortRef, order int) (CFRunLoopSourceRef, error) {
	if _cFMachPortCreateRunLoopSource == nil {
		return 0, symbolCallError("CFMachPortCreateRunLoopSource", "", _cFMachPortCreateRunLoopSourceErr)
	}
	return _cFMachPortCreateRunLoopSource(allocator, port, order), nil
}

// CFMachPortCreateRunLoopSource creates a CFRunLoopSource object for a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateRunLoopSource(_:_:_:)
func CFMachPortCreateRunLoopSource(allocator CFAllocatorRef, port CFMachPortRef, order int) CFRunLoopSourceRef {
	result, callErr := tryCFMachPortCreateRunLoopSource(allocator, port, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortCreateWithPort func(allocator CFAllocatorRef, portNum uint32, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef
var _cFMachPortCreateWithPortErr error

func tryCFMachPortCreateWithPort(allocator CFAllocatorRef, portNum uint32, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) (CFMachPortRef, error) {
	if _cFMachPortCreateWithPort == nil {
		return 0, symbolCallError("CFMachPortCreateWithPort", "", _cFMachPortCreateWithPortErr)
	}
	return _cFMachPortCreateWithPort(allocator, portNum, callout, context, shouldFreeInfo), nil
}

// CFMachPortCreateWithPort creates a CFMachPort object for a pre-existing native Mach port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateWithPort(_:_:_:_:_:)
func CFMachPortCreateWithPort(allocator CFAllocatorRef, portNum uint32, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef {
	result, callErr := tryCFMachPortCreateWithPort(allocator, portNum, callout, context, shouldFreeInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortGetContext func(port CFMachPortRef, context *CFMachPortContext)
var _cFMachPortGetContextErr error

func tryCFMachPortGetContext(port CFMachPortRef, context *CFMachPortContext) error {
	if _cFMachPortGetContext == nil {
		return symbolCallError("CFMachPortGetContext", "", _cFMachPortGetContextErr)
	}
	_cFMachPortGetContext(port, context)
	return nil
}

// CFMachPortGetContext returns the context information for a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetContext(_:_:)
func CFMachPortGetContext(port CFMachPortRef, context *CFMachPortContext) {
	if callErr := tryCFMachPortGetContext(port, context); callErr != nil {
		panic(callErr)
	}
}

var _cFMachPortGetInvalidationCallBack func(port CFMachPortRef) CFMachPortInvalidationCallBack
var _cFMachPortGetInvalidationCallBackErr error

func tryCFMachPortGetInvalidationCallBack(port CFMachPortRef) (CFMachPortInvalidationCallBack, error) {
	if _cFMachPortGetInvalidationCallBack == nil {
		return *new(CFMachPortInvalidationCallBack), symbolCallError("CFMachPortGetInvalidationCallBack", "", _cFMachPortGetInvalidationCallBackErr)
	}
	return _cFMachPortGetInvalidationCallBack(port), nil
}

// CFMachPortGetInvalidationCallBack returns the invalidation callback function for a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetInvalidationCallBack(_:)
func CFMachPortGetInvalidationCallBack(port CFMachPortRef) CFMachPortInvalidationCallBack {
	result, callErr := tryCFMachPortGetInvalidationCallBack(port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortGetPort func(port CFMachPortRef) uint32
var _cFMachPortGetPortErr error

func tryCFMachPortGetPort(port CFMachPortRef) (uint32, error) {
	if _cFMachPortGetPort == nil {
		return 0, symbolCallError("CFMachPortGetPort", "", _cFMachPortGetPortErr)
	}
	return _cFMachPortGetPort(port), nil
}

// CFMachPortGetPort returns the native Mach port represented by a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetPort(_:)
func CFMachPortGetPort(port CFMachPortRef) uint32 {
	result, callErr := tryCFMachPortGetPort(port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortGetTypeID func() uint
var _cFMachPortGetTypeIDErr error

func tryCFMachPortGetTypeID() (uint, error) {
	if _cFMachPortGetTypeID == nil {
		return 0, symbolCallError("CFMachPortGetTypeID", "", _cFMachPortGetTypeIDErr)
	}
	return _cFMachPortGetTypeID(), nil
}

// CFMachPortGetTypeID returns the type identifier for the CFMachPort opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetTypeID()
func CFMachPortGetTypeID() uint {
	result, callErr := tryCFMachPortGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortInvalidate func(port CFMachPortRef)
var _cFMachPortInvalidateErr error

func tryCFMachPortInvalidate(port CFMachPortRef) error {
	if _cFMachPortInvalidate == nil {
		return symbolCallError("CFMachPortInvalidate", "", _cFMachPortInvalidateErr)
	}
	_cFMachPortInvalidate(port)
	return nil
}

// CFMachPortInvalidate invalidates a CFMachPort object, stopping it from receiving any more messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortInvalidate(_:)
func CFMachPortInvalidate(port CFMachPortRef) {
	if callErr := tryCFMachPortInvalidate(port); callErr != nil {
		panic(callErr)
	}
}

var _cFMachPortIsValid func(port CFMachPortRef) bool
var _cFMachPortIsValidErr error

func tryCFMachPortIsValid(port CFMachPortRef) (bool, error) {
	if _cFMachPortIsValid == nil {
		return false, symbolCallError("CFMachPortIsValid", "", _cFMachPortIsValidErr)
	}
	return _cFMachPortIsValid(port), nil
}

// CFMachPortIsValid returns a Boolean value that indicates whether a CFMachPort object is valid and able to receive messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortIsValid(_:)
func CFMachPortIsValid(port CFMachPortRef) bool {
	result, callErr := tryCFMachPortIsValid(port)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMachPortSetInvalidationCallBack func(port CFMachPortRef, callout CFMachPortInvalidationCallBack)
var _cFMachPortSetInvalidationCallBackErr error

func tryCFMachPortSetInvalidationCallBack(port CFMachPortRef, callout CFMachPortInvalidationCallBack) error {
	if _cFMachPortSetInvalidationCallBack == nil {
		return symbolCallError("CFMachPortSetInvalidationCallBack", "", _cFMachPortSetInvalidationCallBackErr)
	}
	_cFMachPortSetInvalidationCallBack(port, callout)
	return nil
}

// CFMachPortSetInvalidationCallBack sets the callback function invoked when a CFMachPort object is invalidated.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortSetInvalidationCallBack(_:_:)
func CFMachPortSetInvalidationCallBack(port CFMachPortRef, callout CFMachPortInvalidationCallBack) {
	if callErr := tryCFMachPortSetInvalidationCallBack(port, callout); callErr != nil {
		panic(callErr)
	}
}

var _cFMakeCollectable func(cf CFTypeRef) CFTypeRef
var _cFMakeCollectableErr error

func tryCFMakeCollectable(cf CFTypeRef) (CFTypeRef, error) {
	if _cFMakeCollectable == nil {
		return 0, symbolCallError("CFMakeCollectable", "", _cFMakeCollectableErr)
	}
	return _cFMakeCollectable(cf), nil
}

// CFMakeCollectable makes a newly-allocated Core Foundation object eligible for garbage collection.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMakeCollectable
func CFMakeCollectable(cf CFTypeRef) CFTypeRef {
	result, callErr := tryCFMakeCollectable(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortCreateLocal func(allocator CFAllocatorRef, name CFStringRef, callout CFMessagePortCallBack, context *CFMessagePortContext, shouldFreeInfo *bool) CFMessagePortRef
var _cFMessagePortCreateLocalErr error

func tryCFMessagePortCreateLocal(allocator CFAllocatorRef, name CFStringRef, callout CFMessagePortCallBack, context *CFMessagePortContext, shouldFreeInfo *bool) (CFMessagePortRef, error) {
	if _cFMessagePortCreateLocal == nil {
		return 0, symbolCallError("CFMessagePortCreateLocal", "", _cFMessagePortCreateLocalErr)
	}
	return _cFMessagePortCreateLocal(allocator, name, callout, context, shouldFreeInfo), nil
}

// CFMessagePortCreateLocal returns a local CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateLocal(_:_:_:_:_:)
func CFMessagePortCreateLocal(allocator CFAllocatorRef, name CFStringRef, callout CFMessagePortCallBack, context *CFMessagePortContext, shouldFreeInfo *bool) CFMessagePortRef {
	result, callErr := tryCFMessagePortCreateLocal(allocator, name, callout, context, shouldFreeInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortCreateRemote func(allocator CFAllocatorRef, name CFStringRef) CFMessagePortRef
var _cFMessagePortCreateRemoteErr error

func tryCFMessagePortCreateRemote(allocator CFAllocatorRef, name CFStringRef) (CFMessagePortRef, error) {
	if _cFMessagePortCreateRemote == nil {
		return 0, symbolCallError("CFMessagePortCreateRemote", "", _cFMessagePortCreateRemoteErr)
	}
	return _cFMessagePortCreateRemote(allocator, name), nil
}

// CFMessagePortCreateRemote returns a CFMessagePort object connected to a remote port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRemote(_:_:)
func CFMessagePortCreateRemote(allocator CFAllocatorRef, name CFStringRef) CFMessagePortRef {
	result, callErr := tryCFMessagePortCreateRemote(allocator, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortCreateRunLoopSource func(allocator CFAllocatorRef, local CFMessagePortRef, order int) CFRunLoopSourceRef
var _cFMessagePortCreateRunLoopSourceErr error

func tryCFMessagePortCreateRunLoopSource(allocator CFAllocatorRef, local CFMessagePortRef, order int) (CFRunLoopSourceRef, error) {
	if _cFMessagePortCreateRunLoopSource == nil {
		return 0, symbolCallError("CFMessagePortCreateRunLoopSource", "", _cFMessagePortCreateRunLoopSourceErr)
	}
	return _cFMessagePortCreateRunLoopSource(allocator, local, order), nil
}

// CFMessagePortCreateRunLoopSource creates a CFRunLoopSource object for a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRunLoopSource(_:_:_:)
func CFMessagePortCreateRunLoopSource(allocator CFAllocatorRef, local CFMessagePortRef, order int) CFRunLoopSourceRef {
	result, callErr := tryCFMessagePortCreateRunLoopSource(allocator, local, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortGetContext func(ms CFMessagePortRef, context *CFMessagePortContext)
var _cFMessagePortGetContextErr error

func tryCFMessagePortGetContext(ms CFMessagePortRef, context *CFMessagePortContext) error {
	if _cFMessagePortGetContext == nil {
		return symbolCallError("CFMessagePortGetContext", "", _cFMessagePortGetContextErr)
	}
	_cFMessagePortGetContext(ms, context)
	return nil
}

// CFMessagePortGetContext returns the context information for a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetContext(_:_:)
func CFMessagePortGetContext(ms CFMessagePortRef, context *CFMessagePortContext) {
	if callErr := tryCFMessagePortGetContext(ms, context); callErr != nil {
		panic(callErr)
	}
}

var _cFMessagePortGetInvalidationCallBack func(ms CFMessagePortRef) CFMessagePortInvalidationCallBack
var _cFMessagePortGetInvalidationCallBackErr error

func tryCFMessagePortGetInvalidationCallBack(ms CFMessagePortRef) (CFMessagePortInvalidationCallBack, error) {
	if _cFMessagePortGetInvalidationCallBack == nil {
		return *new(CFMessagePortInvalidationCallBack), symbolCallError("CFMessagePortGetInvalidationCallBack", "", _cFMessagePortGetInvalidationCallBackErr)
	}
	return _cFMessagePortGetInvalidationCallBack(ms), nil
}

// CFMessagePortGetInvalidationCallBack returns the invalidation callback function for a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetInvalidationCallBack(_:)
func CFMessagePortGetInvalidationCallBack(ms CFMessagePortRef) CFMessagePortInvalidationCallBack {
	result, callErr := tryCFMessagePortGetInvalidationCallBack(ms)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortGetName func(ms CFMessagePortRef) CFStringRef
var _cFMessagePortGetNameErr error

func tryCFMessagePortGetName(ms CFMessagePortRef) (CFStringRef, error) {
	if _cFMessagePortGetName == nil {
		return 0, symbolCallError("CFMessagePortGetName", "", _cFMessagePortGetNameErr)
	}
	return _cFMessagePortGetName(ms), nil
}

// CFMessagePortGetName returns the name with which a CFMessagePort object is registered.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetName(_:)
func CFMessagePortGetName(ms CFMessagePortRef) CFStringRef {
	result, callErr := tryCFMessagePortGetName(ms)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortGetTypeID func() uint
var _cFMessagePortGetTypeIDErr error

func tryCFMessagePortGetTypeID() (uint, error) {
	if _cFMessagePortGetTypeID == nil {
		return 0, symbolCallError("CFMessagePortGetTypeID", "", _cFMessagePortGetTypeIDErr)
	}
	return _cFMessagePortGetTypeID(), nil
}

// CFMessagePortGetTypeID returns the type identifier for the CFMessagePort opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetTypeID()
func CFMessagePortGetTypeID() uint {
	result, callErr := tryCFMessagePortGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortInvalidate func(ms CFMessagePortRef)
var _cFMessagePortInvalidateErr error

func tryCFMessagePortInvalidate(ms CFMessagePortRef) error {
	if _cFMessagePortInvalidate == nil {
		return symbolCallError("CFMessagePortInvalidate", "", _cFMessagePortInvalidateErr)
	}
	_cFMessagePortInvalidate(ms)
	return nil
}

// CFMessagePortInvalidate invalidates a CFMessagePort object, stopping it from receiving or sending any more messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortInvalidate(_:)
func CFMessagePortInvalidate(ms CFMessagePortRef) {
	if callErr := tryCFMessagePortInvalidate(ms); callErr != nil {
		panic(callErr)
	}
}

var _cFMessagePortIsRemote func(ms CFMessagePortRef) bool
var _cFMessagePortIsRemoteErr error

func tryCFMessagePortIsRemote(ms CFMessagePortRef) (bool, error) {
	if _cFMessagePortIsRemote == nil {
		return false, symbolCallError("CFMessagePortIsRemote", "", _cFMessagePortIsRemoteErr)
	}
	return _cFMessagePortIsRemote(ms), nil
}

// CFMessagePortIsRemote returns a Boolean value that indicates whether a CFMessagePort object represents a remote port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsRemote(_:)
func CFMessagePortIsRemote(ms CFMessagePortRef) bool {
	result, callErr := tryCFMessagePortIsRemote(ms)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortIsValid func(ms CFMessagePortRef) bool
var _cFMessagePortIsValidErr error

func tryCFMessagePortIsValid(ms CFMessagePortRef) (bool, error) {
	if _cFMessagePortIsValid == nil {
		return false, symbolCallError("CFMessagePortIsValid", "", _cFMessagePortIsValidErr)
	}
	return _cFMessagePortIsValid(ms), nil
}

// CFMessagePortIsValid returns a Boolean value that indicates whether a CFMessagePort object is valid and able to send or receive messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsValid(_:)
func CFMessagePortIsValid(ms CFMessagePortRef) bool {
	result, callErr := tryCFMessagePortIsValid(ms)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortSendRequest func(remote CFMessagePortRef, msgid int32, data CFDataRef, sendTimeout float64, rcvTimeout float64, replyMode CFStringRef, returnData *CFDataRef) int32
var _cFMessagePortSendRequestErr error

func tryCFMessagePortSendRequest(remote CFMessagePortRef, msgid int32, data CFDataRef, sendTimeout float64, rcvTimeout float64, replyMode CFStringRef, returnData *CFDataRef) (int32, error) {
	if _cFMessagePortSendRequest == nil {
		return 0, symbolCallError("CFMessagePortSendRequest", "", _cFMessagePortSendRequestErr)
	}
	return _cFMessagePortSendRequest(remote, msgid, data, sendTimeout, rcvTimeout, replyMode, returnData), nil
}

// CFMessagePortSendRequest sends a message to a remote CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSendRequest(_:_:_:_:_:_:_:)
func CFMessagePortSendRequest(remote CFMessagePortRef, msgid int32, data CFDataRef, sendTimeout float64, rcvTimeout float64, replyMode CFStringRef, returnData *CFDataRef) int32 {
	result, callErr := tryCFMessagePortSendRequest(remote, msgid, data, sendTimeout, rcvTimeout, replyMode, returnData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFMessagePortSetDispatchQueue func(ms CFMessagePortRef, queue uintptr)
var _cFMessagePortSetDispatchQueueErr error

func tryCFMessagePortSetDispatchQueue(ms CFMessagePortRef, queue dispatch.Queue) error {
	if _cFMessagePortSetDispatchQueue == nil {
		return symbolCallError("CFMessagePortSetDispatchQueue", "10.6", _cFMessagePortSetDispatchQueueErr)
	}
	_cFMessagePortSetDispatchQueue(ms, uintptr(queue.Handle()))
	return nil
}

// CFMessagePortSetDispatchQueue schedules callbacks for the specified message port on the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetDispatchQueue(_:_:)
func CFMessagePortSetDispatchQueue(ms CFMessagePortRef, queue dispatch.Queue) {
	if callErr := tryCFMessagePortSetDispatchQueue(ms, queue); callErr != nil {
		panic(callErr)
	}
}

var _cFMessagePortSetInvalidationCallBack func(ms CFMessagePortRef, callout CFMessagePortInvalidationCallBack)
var _cFMessagePortSetInvalidationCallBackErr error

func tryCFMessagePortSetInvalidationCallBack(ms CFMessagePortRef, callout CFMessagePortInvalidationCallBack) error {
	if _cFMessagePortSetInvalidationCallBack == nil {
		return symbolCallError("CFMessagePortSetInvalidationCallBack", "", _cFMessagePortSetInvalidationCallBackErr)
	}
	_cFMessagePortSetInvalidationCallBack(ms, callout)
	return nil
}

// CFMessagePortSetInvalidationCallBack sets the callback function invoked when a CFMessagePort object is invalidated.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetInvalidationCallBack(_:_:)
func CFMessagePortSetInvalidationCallBack(ms CFMessagePortRef, callout CFMessagePortInvalidationCallBack) {
	if callErr := tryCFMessagePortSetInvalidationCallBack(ms, callout); callErr != nil {
		panic(callErr)
	}
}

var _cFMessagePortSetName func(ms CFMessagePortRef, newName CFStringRef) bool
var _cFMessagePortSetNameErr error

func tryCFMessagePortSetName(ms CFMessagePortRef, newName CFStringRef) (bool, error) {
	if _cFMessagePortSetName == nil {
		return false, symbolCallError("CFMessagePortSetName", "", _cFMessagePortSetNameErr)
	}
	return _cFMessagePortSetName(ms, newName), nil
}

// CFMessagePortSetName sets the name of a local CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetName(_:_:)
func CFMessagePortSetName(ms CFMessagePortRef, newName CFStringRef) bool {
	result, callErr := tryCFMessagePortSetName(ms, newName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNotificationCenterAddObserver func(center CFNotificationCenterRef, observer unsafe.Pointer, callBack CFNotificationCallback, name CFStringRef, object unsafe.Pointer, suspensionBehavior CFNotificationSuspensionBehavior)
var _cFNotificationCenterAddObserverErr error

func tryCFNotificationCenterAddObserver(center CFNotificationCenterRef, observer unsafe.Pointer, callBack CFNotificationCallback, name CFStringRef, object unsafe.Pointer, suspensionBehavior CFNotificationSuspensionBehavior) error {
	if _cFNotificationCenterAddObserver == nil {
		return symbolCallError("CFNotificationCenterAddObserver", "", _cFNotificationCenterAddObserverErr)
	}
	_cFNotificationCenterAddObserver(center, observer, callBack, name, object, suspensionBehavior)
	return nil
}

// CFNotificationCenterAddObserver registers an observer to receive notifications.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterAddObserver(_:_:_:_:_:_:)
func CFNotificationCenterAddObserver(center CFNotificationCenterRef, observer unsafe.Pointer, callBack CFNotificationCallback, name CFStringRef, object unsafe.Pointer, suspensionBehavior CFNotificationSuspensionBehavior) {
	if callErr := tryCFNotificationCenterAddObserver(center, observer, callBack, name, object, suspensionBehavior); callErr != nil {
		panic(callErr)
	}
}

var _cFNotificationCenterGetDarwinNotifyCenter func() CFNotificationCenterRef
var _cFNotificationCenterGetDarwinNotifyCenterErr error

func tryCFNotificationCenterGetDarwinNotifyCenter() (CFNotificationCenterRef, error) {
	if _cFNotificationCenterGetDarwinNotifyCenter == nil {
		return 0, symbolCallError("CFNotificationCenterGetDarwinNotifyCenter", "", _cFNotificationCenterGetDarwinNotifyCenterErr)
	}
	return _cFNotificationCenterGetDarwinNotifyCenter(), nil
}

// CFNotificationCenterGetDarwinNotifyCenter returns the application’s Darwin notification center.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDarwinNotifyCenter()
func CFNotificationCenterGetDarwinNotifyCenter() CFNotificationCenterRef {
	result, callErr := tryCFNotificationCenterGetDarwinNotifyCenter()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNotificationCenterGetDistributedCenter func() CFNotificationCenterRef
var _cFNotificationCenterGetDistributedCenterErr error

func tryCFNotificationCenterGetDistributedCenter() (CFNotificationCenterRef, error) {
	if _cFNotificationCenterGetDistributedCenter == nil {
		return 0, symbolCallError("CFNotificationCenterGetDistributedCenter", "", _cFNotificationCenterGetDistributedCenterErr)
	}
	return _cFNotificationCenterGetDistributedCenter(), nil
}

// CFNotificationCenterGetDistributedCenter returns the application’s distributed notification center.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDistributedCenter()
func CFNotificationCenterGetDistributedCenter() CFNotificationCenterRef {
	result, callErr := tryCFNotificationCenterGetDistributedCenter()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNotificationCenterGetLocalCenter func() CFNotificationCenterRef
var _cFNotificationCenterGetLocalCenterErr error

func tryCFNotificationCenterGetLocalCenter() (CFNotificationCenterRef, error) {
	if _cFNotificationCenterGetLocalCenter == nil {
		return 0, symbolCallError("CFNotificationCenterGetLocalCenter", "", _cFNotificationCenterGetLocalCenterErr)
	}
	return _cFNotificationCenterGetLocalCenter(), nil
}

// CFNotificationCenterGetLocalCenter returns the application’s local notification center.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetLocalCenter()
func CFNotificationCenterGetLocalCenter() CFNotificationCenterRef {
	result, callErr := tryCFNotificationCenterGetLocalCenter()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNotificationCenterGetTypeID func() uint
var _cFNotificationCenterGetTypeIDErr error

func tryCFNotificationCenterGetTypeID() (uint, error) {
	if _cFNotificationCenterGetTypeID == nil {
		return 0, symbolCallError("CFNotificationCenterGetTypeID", "", _cFNotificationCenterGetTypeIDErr)
	}
	return _cFNotificationCenterGetTypeID(), nil
}

// CFNotificationCenterGetTypeID returns the type identifier for the CFNotificationCenter opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetTypeID()
func CFNotificationCenterGetTypeID() uint {
	result, callErr := tryCFNotificationCenterGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNotificationCenterPostNotification func(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, deliverImmediately bool)
var _cFNotificationCenterPostNotificationErr error

func tryCFNotificationCenterPostNotification(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, deliverImmediately bool) error {
	if _cFNotificationCenterPostNotification == nil {
		return symbolCallError("CFNotificationCenterPostNotification", "", _cFNotificationCenterPostNotificationErr)
	}
	_cFNotificationCenterPostNotification(center, name, object, userInfo, deliverImmediately)
	return nil
}

// CFNotificationCenterPostNotification posts a notification for an object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotification(_:_:_:_:_:)
func CFNotificationCenterPostNotification(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, deliverImmediately bool) {
	if callErr := tryCFNotificationCenterPostNotification(center, name, object, userInfo, deliverImmediately); callErr != nil {
		panic(callErr)
	}
}

var _cFNotificationCenterPostNotificationWithOptions func(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, options uint64)
var _cFNotificationCenterPostNotificationWithOptionsErr error

func tryCFNotificationCenterPostNotificationWithOptions(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, options uint64) error {
	if _cFNotificationCenterPostNotificationWithOptions == nil {
		return symbolCallError("CFNotificationCenterPostNotificationWithOptions", "", _cFNotificationCenterPostNotificationWithOptionsErr)
	}
	_cFNotificationCenterPostNotificationWithOptions(center, name, object, userInfo, options)
	return nil
}

// CFNotificationCenterPostNotificationWithOptions posts a notification for an object using specified options.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotificationWithOptions(_:_:_:_:_:)
func CFNotificationCenterPostNotificationWithOptions(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, options uint64) {
	if callErr := tryCFNotificationCenterPostNotificationWithOptions(center, name, object, userInfo, options); callErr != nil {
		panic(callErr)
	}
}

var _cFNotificationCenterRemoveEveryObserver func(center CFNotificationCenterRef, observer unsafe.Pointer)
var _cFNotificationCenterRemoveEveryObserverErr error

func tryCFNotificationCenterRemoveEveryObserver(center CFNotificationCenterRef, observer unsafe.Pointer) error {
	if _cFNotificationCenterRemoveEveryObserver == nil {
		return symbolCallError("CFNotificationCenterRemoveEveryObserver", "", _cFNotificationCenterRemoveEveryObserverErr)
	}
	_cFNotificationCenterRemoveEveryObserver(center, observer)
	return nil
}

// CFNotificationCenterRemoveEveryObserver stops an observer from receiving any notifications from any object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveEveryObserver(_:_:)
func CFNotificationCenterRemoveEveryObserver(center CFNotificationCenterRef, observer unsafe.Pointer) {
	if callErr := tryCFNotificationCenterRemoveEveryObserver(center, observer); callErr != nil {
		panic(callErr)
	}
}

var _cFNotificationCenterRemoveObserver func(center CFNotificationCenterRef, observer unsafe.Pointer, name CFNotificationName, object unsafe.Pointer)
var _cFNotificationCenterRemoveObserverErr error

func tryCFNotificationCenterRemoveObserver(center CFNotificationCenterRef, observer unsafe.Pointer, name CFNotificationName, object unsafe.Pointer) error {
	if _cFNotificationCenterRemoveObserver == nil {
		return symbolCallError("CFNotificationCenterRemoveObserver", "", _cFNotificationCenterRemoveObserverErr)
	}
	_cFNotificationCenterRemoveObserver(center, observer, name, object)
	return nil
}

// CFNotificationCenterRemoveObserver stops an observer from receiving certain notifications.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveObserver(_:_:_:_:)
func CFNotificationCenterRemoveObserver(center CFNotificationCenterRef, observer unsafe.Pointer, name CFNotificationName, object unsafe.Pointer) {
	if callErr := tryCFNotificationCenterRemoveObserver(center, observer, name, object); callErr != nil {
		panic(callErr)
	}
}

var _cFNullGetTypeID func() uint
var _cFNullGetTypeIDErr error

func tryCFNullGetTypeID() (uint, error) {
	if _cFNullGetTypeID == nil {
		return 0, symbolCallError("CFNullGetTypeID", "", _cFNullGetTypeIDErr)
	}
	return _cFNullGetTypeID(), nil
}

// CFNullGetTypeID returns the type identifier for the CFNull opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNullGetTypeID()
func CFNullGetTypeID() uint {
	result, callErr := tryCFNullGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberCompare func(number CFNumberRef, otherNumber CFNumberRef, context unsafe.Pointer) CFComparisonResult
var _cFNumberCompareErr error

func tryCFNumberCompare(number CFNumberRef, otherNumber CFNumberRef, context unsafe.Pointer) (CFComparisonResult, error) {
	if _cFNumberCompare == nil {
		return *new(CFComparisonResult), symbolCallError("CFNumberCompare", "", _cFNumberCompareErr)
	}
	return _cFNumberCompare(number, otherNumber, context), nil
}

// CFNumberCompare compares two CFNumber objects and returns a comparison result.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberCompare(_:_:_:)
func CFNumberCompare(number CFNumberRef, otherNumber CFNumberRef, context unsafe.Pointer) CFComparisonResult {
	result, callErr := tryCFNumberCompare(number, otherNumber, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberCreate func(allocator CFAllocatorRef, theType CFNumberType, valuePtr unsafe.Pointer) CFNumberRef
var _cFNumberCreateErr error

func tryCFNumberCreate(allocator CFAllocatorRef, theType CFNumberType, valuePtr unsafe.Pointer) (CFNumberRef, error) {
	if _cFNumberCreate == nil {
		return 0, symbolCallError("CFNumberCreate", "", _cFNumberCreateErr)
	}
	return _cFNumberCreate(allocator, theType, valuePtr), nil
}

// CFNumberCreate creates a CFNumber object using a specified value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberCreate(_:_:_:)
func CFNumberCreate(allocator CFAllocatorRef, theType CFNumberType, valuePtr unsafe.Pointer) CFNumberRef {
	result, callErr := tryCFNumberCreate(allocator, theType, valuePtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterCopyProperty func(formatter CFNumberFormatterRef, key CFNumberFormatterKey) CFTypeRef
var _cFNumberFormatterCopyPropertyErr error

func tryCFNumberFormatterCopyProperty(formatter CFNumberFormatterRef, key CFNumberFormatterKey) (CFTypeRef, error) {
	if _cFNumberFormatterCopyProperty == nil {
		return 0, symbolCallError("CFNumberFormatterCopyProperty", "", _cFNumberFormatterCopyPropertyErr)
	}
	return _cFNumberFormatterCopyProperty(formatter, key), nil
}

// CFNumberFormatterCopyProperty returns a copy of a number formatter’s value for a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCopyProperty(_:_:)
func CFNumberFormatterCopyProperty(formatter CFNumberFormatterRef, key CFNumberFormatterKey) CFTypeRef {
	result, callErr := tryCFNumberFormatterCopyProperty(formatter, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterCreate func(allocator CFAllocatorRef, locale CFLocaleRef, style CFNumberFormatterStyle) CFNumberFormatterRef
var _cFNumberFormatterCreateErr error

func tryCFNumberFormatterCreate(allocator CFAllocatorRef, locale CFLocaleRef, style CFNumberFormatterStyle) (CFNumberFormatterRef, error) {
	if _cFNumberFormatterCreate == nil {
		return 0, symbolCallError("CFNumberFormatterCreate", "", _cFNumberFormatterCreateErr)
	}
	return _cFNumberFormatterCreate(allocator, locale, style), nil
}

// CFNumberFormatterCreate creates a new CFNumberFormatter object, localized to the given locale, which will format numbers to the given style.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreate(_:_:_:)
func CFNumberFormatterCreate(allocator CFAllocatorRef, locale CFLocaleRef, style CFNumberFormatterStyle) CFNumberFormatterRef {
	result, callErr := tryCFNumberFormatterCreate(allocator, locale, style)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterCreateNumberFromString func(allocator CFAllocatorRef, formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, options uint64) CFNumberRef
var _cFNumberFormatterCreateNumberFromStringErr error

func tryCFNumberFormatterCreateNumberFromString(allocator CFAllocatorRef, formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, options uint64) (CFNumberRef, error) {
	if _cFNumberFormatterCreateNumberFromString == nil {
		return 0, symbolCallError("CFNumberFormatterCreateNumberFromString", "", _cFNumberFormatterCreateNumberFromStringErr)
	}
	return _cFNumberFormatterCreateNumberFromString(allocator, formatter, string_, rangep, options), nil
}

// CFNumberFormatterCreateNumberFromString returns a number object representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateNumberFromString(_:_:_:_:_:)
func CFNumberFormatterCreateNumberFromString(allocator CFAllocatorRef, formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, options uint64) CFNumberRef {
	result, callErr := tryCFNumberFormatterCreateNumberFromString(allocator, formatter, string_, rangep, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterCreateStringWithNumber func(allocator CFAllocatorRef, formatter CFNumberFormatterRef, number CFNumberRef) CFStringRef
var _cFNumberFormatterCreateStringWithNumberErr error

func tryCFNumberFormatterCreateStringWithNumber(allocator CFAllocatorRef, formatter CFNumberFormatterRef, number CFNumberRef) (CFStringRef, error) {
	if _cFNumberFormatterCreateStringWithNumber == nil {
		return 0, symbolCallError("CFNumberFormatterCreateStringWithNumber", "", _cFNumberFormatterCreateStringWithNumberErr)
	}
	return _cFNumberFormatterCreateStringWithNumber(allocator, formatter, number), nil
}

// CFNumberFormatterCreateStringWithNumber returns a string representation of the given number using the specified number formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithNumber(_:_:_:)
func CFNumberFormatterCreateStringWithNumber(allocator CFAllocatorRef, formatter CFNumberFormatterRef, number CFNumberRef) CFStringRef {
	result, callErr := tryCFNumberFormatterCreateStringWithNumber(allocator, formatter, number)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterCreateStringWithValue func(allocator CFAllocatorRef, formatter CFNumberFormatterRef, numberType CFNumberType, valuePtr unsafe.Pointer) CFStringRef
var _cFNumberFormatterCreateStringWithValueErr error

func tryCFNumberFormatterCreateStringWithValue(allocator CFAllocatorRef, formatter CFNumberFormatterRef, numberType CFNumberType, valuePtr unsafe.Pointer) (CFStringRef, error) {
	if _cFNumberFormatterCreateStringWithValue == nil {
		return 0, symbolCallError("CFNumberFormatterCreateStringWithValue", "", _cFNumberFormatterCreateStringWithValueErr)
	}
	return _cFNumberFormatterCreateStringWithValue(allocator, formatter, numberType, valuePtr), nil
}

// CFNumberFormatterCreateStringWithValue returns a string representation of the given number or value using the specified number formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithValue(_:_:_:_:)
func CFNumberFormatterCreateStringWithValue(allocator CFAllocatorRef, formatter CFNumberFormatterRef, numberType CFNumberType, valuePtr unsafe.Pointer) CFStringRef {
	result, callErr := tryCFNumberFormatterCreateStringWithValue(allocator, formatter, numberType, valuePtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterGetDecimalInfoForCurrencyCode func(currencyCode CFStringRef, defaultFractionDigits *int32, roundingIncrement *float64) bool
var _cFNumberFormatterGetDecimalInfoForCurrencyCodeErr error

func tryCFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode CFStringRef, defaultFractionDigits *int32, roundingIncrement []float64) (bool, error) {
	if _cFNumberFormatterGetDecimalInfoForCurrencyCode == nil {
		return false, symbolCallError("CFNumberFormatterGetDecimalInfoForCurrencyCode", "", _cFNumberFormatterGetDecimalInfoForCurrencyCodeErr)
	}
	return _cFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode, defaultFractionDigits, unsafe.SliceData(roundingIncrement)), nil
}

// CFNumberFormatterGetDecimalInfoForCurrencyCode returns the number of fraction digits that should be displayed, and the rounding increment, for a given currency.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetDecimalInfoForCurrencyCode(_:_:_:)
func CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode CFStringRef, defaultFractionDigits *int32, roundingIncrement []float64) bool {
	result, callErr := tryCFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode, defaultFractionDigits, roundingIncrement)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterGetFormat func(formatter CFNumberFormatterRef) CFStringRef
var _cFNumberFormatterGetFormatErr error

func tryCFNumberFormatterGetFormat(formatter CFNumberFormatterRef) (CFStringRef, error) {
	if _cFNumberFormatterGetFormat == nil {
		return 0, symbolCallError("CFNumberFormatterGetFormat", "", _cFNumberFormatterGetFormatErr)
	}
	return _cFNumberFormatterGetFormat(formatter), nil
}

// CFNumberFormatterGetFormat returns a format string for the given number formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetFormat(_:)
func CFNumberFormatterGetFormat(formatter CFNumberFormatterRef) CFStringRef {
	result, callErr := tryCFNumberFormatterGetFormat(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterGetLocale func(formatter CFNumberFormatterRef) CFLocaleRef
var _cFNumberFormatterGetLocaleErr error

func tryCFNumberFormatterGetLocale(formatter CFNumberFormatterRef) (CFLocaleRef, error) {
	if _cFNumberFormatterGetLocale == nil {
		return 0, symbolCallError("CFNumberFormatterGetLocale", "", _cFNumberFormatterGetLocaleErr)
	}
	return _cFNumberFormatterGetLocale(formatter), nil
}

// CFNumberFormatterGetLocale returns the locale object used to create the given number formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetLocale(_:)
func CFNumberFormatterGetLocale(formatter CFNumberFormatterRef) CFLocaleRef {
	result, callErr := tryCFNumberFormatterGetLocale(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterGetStyle func(formatter CFNumberFormatterRef) CFNumberFormatterStyle
var _cFNumberFormatterGetStyleErr error

func tryCFNumberFormatterGetStyle(formatter CFNumberFormatterRef) (CFNumberFormatterStyle, error) {
	if _cFNumberFormatterGetStyle == nil {
		return *new(CFNumberFormatterStyle), symbolCallError("CFNumberFormatterGetStyle", "", _cFNumberFormatterGetStyleErr)
	}
	return _cFNumberFormatterGetStyle(formatter), nil
}

// CFNumberFormatterGetStyle returns the number style used to create the given number formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetStyle(_:)
func CFNumberFormatterGetStyle(formatter CFNumberFormatterRef) CFNumberFormatterStyle {
	result, callErr := tryCFNumberFormatterGetStyle(formatter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterGetTypeID func() uint
var _cFNumberFormatterGetTypeIDErr error

func tryCFNumberFormatterGetTypeID() (uint, error) {
	if _cFNumberFormatterGetTypeID == nil {
		return 0, symbolCallError("CFNumberFormatterGetTypeID", "", _cFNumberFormatterGetTypeIDErr)
	}
	return _cFNumberFormatterGetTypeID(), nil
}

// CFNumberFormatterGetTypeID returns the type identifier for the [CFNumberFormatter] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetTypeID()
func CFNumberFormatterGetTypeID() uint {
	result, callErr := tryCFNumberFormatterGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterGetValueFromString func(formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, numberType CFNumberType, valuePtr unsafe.Pointer) bool
var _cFNumberFormatterGetValueFromStringErr error

func tryCFNumberFormatterGetValueFromString(formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, numberType CFNumberType, valuePtr unsafe.Pointer) (bool, error) {
	if _cFNumberFormatterGetValueFromString == nil {
		return false, symbolCallError("CFNumberFormatterGetValueFromString", "", _cFNumberFormatterGetValueFromStringErr)
	}
	return _cFNumberFormatterGetValueFromString(formatter, string_, rangep, numberType, valuePtr), nil
}

// CFNumberFormatterGetValueFromString returns a number or value representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetValueFromString(_:_:_:_:_:)
func CFNumberFormatterGetValueFromString(formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, numberType CFNumberType, valuePtr unsafe.Pointer) bool {
	result, callErr := tryCFNumberFormatterGetValueFromString(formatter, string_, rangep, numberType, valuePtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberFormatterSetFormat func(formatter CFNumberFormatterRef, formatString CFStringRef)
var _cFNumberFormatterSetFormatErr error

func tryCFNumberFormatterSetFormat(formatter CFNumberFormatterRef, formatString CFStringRef) error {
	if _cFNumberFormatterSetFormat == nil {
		return symbolCallError("CFNumberFormatterSetFormat", "", _cFNumberFormatterSetFormatErr)
	}
	_cFNumberFormatterSetFormat(formatter, formatString)
	return nil
}

// CFNumberFormatterSetFormat sets the format string of a number formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetFormat(_:_:)
func CFNumberFormatterSetFormat(formatter CFNumberFormatterRef, formatString CFStringRef) {
	if callErr := tryCFNumberFormatterSetFormat(formatter, formatString); callErr != nil {
		panic(callErr)
	}
}

var _cFNumberFormatterSetProperty func(formatter CFNumberFormatterRef, key CFNumberFormatterKey, value CFTypeRef)
var _cFNumberFormatterSetPropertyErr error

func tryCFNumberFormatterSetProperty(formatter CFNumberFormatterRef, key CFNumberFormatterKey, value CFTypeRef) error {
	if _cFNumberFormatterSetProperty == nil {
		return symbolCallError("CFNumberFormatterSetProperty", "", _cFNumberFormatterSetPropertyErr)
	}
	_cFNumberFormatterSetProperty(formatter, key, value)
	return nil
}

// CFNumberFormatterSetProperty sets a number formatter property using a key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetProperty(_:_:_:)
func CFNumberFormatterSetProperty(formatter CFNumberFormatterRef, key CFNumberFormatterKey, value CFTypeRef) {
	if callErr := tryCFNumberFormatterSetProperty(formatter, key, value); callErr != nil {
		panic(callErr)
	}
}

var _cFNumberGetByteSize func(number CFNumberRef) int
var _cFNumberGetByteSizeErr error

func tryCFNumberGetByteSize(number CFNumberRef) (int, error) {
	if _cFNumberGetByteSize == nil {
		return 0, symbolCallError("CFNumberGetByteSize", "", _cFNumberGetByteSizeErr)
	}
	return _cFNumberGetByteSize(number), nil
}

// CFNumberGetByteSize returns the number of bytes used by a CFNumber object to store its value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetByteSize(_:)
func CFNumberGetByteSize(number CFNumberRef) int {
	result, callErr := tryCFNumberGetByteSize(number)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberGetType func(number CFNumberRef) CFNumberType
var _cFNumberGetTypeErr error

func tryCFNumberGetType(number CFNumberRef) (CFNumberType, error) {
	if _cFNumberGetType == nil {
		return *new(CFNumberType), symbolCallError("CFNumberGetType", "", _cFNumberGetTypeErr)
	}
	return _cFNumberGetType(number), nil
}

// CFNumberGetType returns the type used by a CFNumber object to store its value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetType(_:)
func CFNumberGetType(number CFNumberRef) CFNumberType {
	result, callErr := tryCFNumberGetType(number)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberGetTypeID func() uint
var _cFNumberGetTypeIDErr error

func tryCFNumberGetTypeID() (uint, error) {
	if _cFNumberGetTypeID == nil {
		return 0, symbolCallError("CFNumberGetTypeID", "", _cFNumberGetTypeIDErr)
	}
	return _cFNumberGetTypeID(), nil
}

// CFNumberGetTypeID returns the type identifier for the CFNumber opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetTypeID()
func CFNumberGetTypeID() uint {
	result, callErr := tryCFNumberGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberGetValue func(number CFNumberRef, theType CFNumberType, valuePtr unsafe.Pointer) bool
var _cFNumberGetValueErr error

func tryCFNumberGetValue(number CFNumberRef, theType CFNumberType, valuePtr unsafe.Pointer) (bool, error) {
	if _cFNumberGetValue == nil {
		return false, symbolCallError("CFNumberGetValue", "", _cFNumberGetValueErr)
	}
	return _cFNumberGetValue(number, theType, valuePtr), nil
}

// CFNumberGetValue obtains the value of a CFNumber object cast to a specified type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetValue(_:_:_:)
func CFNumberGetValue(number CFNumberRef, theType CFNumberType, valuePtr unsafe.Pointer) bool {
	result, callErr := tryCFNumberGetValue(number, theType, valuePtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFNumberIsFloatType func(number CFNumberRef) bool
var _cFNumberIsFloatTypeErr error

func tryCFNumberIsFloatType(number CFNumberRef) (bool, error) {
	if _cFNumberIsFloatType == nil {
		return false, symbolCallError("CFNumberIsFloatType", "", _cFNumberIsFloatTypeErr)
	}
	return _cFNumberIsFloatType(number), nil
}

// CFNumberIsFloatType determines whether a CFNumber object contains a value stored as one of the defined floating point types.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberIsFloatType(_:)
func CFNumberIsFloatType(number CFNumberRef) bool {
	result, callErr := tryCFNumberIsFloatType(number)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInAddInstanceForFactory func(factoryID CFUUIDRef)
var _cFPlugInAddInstanceForFactoryErr error

func tryCFPlugInAddInstanceForFactory(factoryID CFUUIDRef) error {
	if _cFPlugInAddInstanceForFactory == nil {
		return symbolCallError("CFPlugInAddInstanceForFactory", "", _cFPlugInAddInstanceForFactoryErr)
	}
	_cFPlugInAddInstanceForFactory(factoryID)
	return nil
}

// CFPlugInAddInstanceForFactory registers a new instance of a type with [CFPlugIn].
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInAddInstanceForFactory(_:)
func CFPlugInAddInstanceForFactory(factoryID CFUUIDRef) {
	if callErr := tryCFPlugInAddInstanceForFactory(factoryID); callErr != nil {
		panic(callErr)
	}
}

var _cFPlugInCreate func(allocator CFAllocatorRef, plugInURL CFURLRef) CFPlugInRef
var _cFPlugInCreateErr error

func tryCFPlugInCreate(allocator CFAllocatorRef, plugInURL CFURLRef) (CFPlugInRef, error) {
	if _cFPlugInCreate == nil {
		return 0, symbolCallError("CFPlugInCreate", "", _cFPlugInCreateErr)
	}
	return _cFPlugInCreate(allocator, plugInURL), nil
}

// CFPlugInCreate creates a CFPlugIn given its URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInCreate(_:_:)
func CFPlugInCreate(allocator CFAllocatorRef, plugInURL CFURLRef) CFPlugInRef {
	result, callErr := tryCFPlugInCreate(allocator, plugInURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInFindFactoriesForPlugInType func(typeUUID CFUUIDRef) CFArrayRef
var _cFPlugInFindFactoriesForPlugInTypeErr error

func tryCFPlugInFindFactoriesForPlugInType(typeUUID CFUUIDRef) (CFArrayRef, error) {
	if _cFPlugInFindFactoriesForPlugInType == nil {
		return 0, symbolCallError("CFPlugInFindFactoriesForPlugInType", "", _cFPlugInFindFactoriesForPlugInTypeErr)
	}
	return _cFPlugInFindFactoriesForPlugInType(typeUUID), nil
}

// CFPlugInFindFactoriesForPlugInType searches all registered plug-ins for factory functions capable of creating an instance of the given type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInType(_:)
func CFPlugInFindFactoriesForPlugInType(typeUUID CFUUIDRef) CFArrayRef {
	result, callErr := tryCFPlugInFindFactoriesForPlugInType(typeUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInFindFactoriesForPlugInTypeInPlugIn func(typeUUID CFUUIDRef, plugIn CFPlugInRef) CFArrayRef
var _cFPlugInFindFactoriesForPlugInTypeInPlugInErr error

func tryCFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID CFUUIDRef, plugIn CFPlugInRef) (CFArrayRef, error) {
	if _cFPlugInFindFactoriesForPlugInTypeInPlugIn == nil {
		return 0, symbolCallError("CFPlugInFindFactoriesForPlugInTypeInPlugIn", "", _cFPlugInFindFactoriesForPlugInTypeInPlugInErr)
	}
	return _cFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID, plugIn), nil
}

// CFPlugInFindFactoriesForPlugInTypeInPlugIn searches the given plug-in for factory functions capable of creating an instance of the given type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInTypeInPlugIn(_:_:)
func CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID CFUUIDRef, plugIn CFPlugInRef) CFArrayRef {
	result, callErr := tryCFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID, plugIn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInGetBundle func(plugIn CFPlugInRef) CFBundleRef
var _cFPlugInGetBundleErr error

func tryCFPlugInGetBundle(plugIn CFPlugInRef) (CFBundleRef, error) {
	if _cFPlugInGetBundle == nil {
		return 0, symbolCallError("CFPlugInGetBundle", "", _cFPlugInGetBundleErr)
	}
	return _cFPlugInGetBundle(plugIn), nil
}

// CFPlugInGetBundle returns a plug-in’s bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetBundle(_:)
func CFPlugInGetBundle(plugIn CFPlugInRef) CFBundleRef {
	result, callErr := tryCFPlugInGetBundle(plugIn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInGetTypeID func() uint
var _cFPlugInGetTypeIDErr error

func tryCFPlugInGetTypeID() (uint, error) {
	if _cFPlugInGetTypeID == nil {
		return 0, symbolCallError("CFPlugInGetTypeID", "", _cFPlugInGetTypeIDErr)
	}
	return _cFPlugInGetTypeID(), nil
}

// CFPlugInGetTypeID returns the type identifier for the [CFPlugIn] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetTypeID()
func CFPlugInGetTypeID() uint {
	result, callErr := tryCFPlugInGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInInstanceCreate func(allocator CFAllocatorRef, factoryUUID CFUUIDRef, typeUUID CFUUIDRef) unsafe.Pointer
var _cFPlugInInstanceCreateErr error

func tryCFPlugInInstanceCreate(allocator CFAllocatorRef, factoryUUID CFUUIDRef, typeUUID CFUUIDRef) (unsafe.Pointer, error) {
	if _cFPlugInInstanceCreate == nil {
		return nil, symbolCallError("CFPlugInInstanceCreate", "", _cFPlugInInstanceCreateErr)
	}
	return _cFPlugInInstanceCreate(allocator, factoryUUID, typeUUID), nil
}

// CFPlugInInstanceCreate creates a [CFPlugIn] instance of a given type using a given factory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreate(_:_:_:)
func CFPlugInInstanceCreate(allocator CFAllocatorRef, factoryUUID CFUUIDRef, typeUUID CFUUIDRef) unsafe.Pointer {
	result, callErr := tryCFPlugInInstanceCreate(allocator, factoryUUID, typeUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInInstanceCreateWithInstanceDataSize func(allocator CFAllocatorRef, instanceDataSize int, deallocateInstanceFunction CFPlugInInstanceDeallocateInstanceDataFunction, factoryName CFStringRef, getInterfaceFunction CFPlugInInstanceGetInterfaceFunction) CFPlugInInstanceRef
var _cFPlugInInstanceCreateWithInstanceDataSizeErr error

func tryCFPlugInInstanceCreateWithInstanceDataSize(allocator CFAllocatorRef, instanceDataSize int, deallocateInstanceFunction CFPlugInInstanceDeallocateInstanceDataFunction, factoryName CFStringRef, getInterfaceFunction CFPlugInInstanceGetInterfaceFunction) (CFPlugInInstanceRef, error) {
	if _cFPlugInInstanceCreateWithInstanceDataSize == nil {
		return 0, symbolCallError("CFPlugInInstanceCreateWithInstanceDataSize", "", _cFPlugInInstanceCreateWithInstanceDataSizeErr)
	}
	return _cFPlugInInstanceCreateWithInstanceDataSize(allocator, instanceDataSize, deallocateInstanceFunction, factoryName, getInterfaceFunction), nil
}

// CFPlugInInstanceCreateWithInstanceDataSize not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreateWithInstanceDataSize(_:_:_:_:_:)
func CFPlugInInstanceCreateWithInstanceDataSize(allocator CFAllocatorRef, instanceDataSize int, deallocateInstanceFunction CFPlugInInstanceDeallocateInstanceDataFunction, factoryName CFStringRef, getInterfaceFunction CFPlugInInstanceGetInterfaceFunction) CFPlugInInstanceRef {
	result, callErr := tryCFPlugInInstanceCreateWithInstanceDataSize(allocator, instanceDataSize, deallocateInstanceFunction, factoryName, getInterfaceFunction)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInInstanceGetFactoryName func(instance CFPlugInInstanceRef) CFStringRef
var _cFPlugInInstanceGetFactoryNameErr error

func tryCFPlugInInstanceGetFactoryName(instance CFPlugInInstanceRef) (CFStringRef, error) {
	if _cFPlugInInstanceGetFactoryName == nil {
		return 0, symbolCallError("CFPlugInInstanceGetFactoryName", "", _cFPlugInInstanceGetFactoryNameErr)
	}
	return _cFPlugInInstanceGetFactoryName(instance), nil
}

// CFPlugInInstanceGetFactoryName not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetFactoryName(_:)
func CFPlugInInstanceGetFactoryName(instance CFPlugInInstanceRef) CFStringRef {
	result, callErr := tryCFPlugInInstanceGetFactoryName(instance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInInstanceGetInstanceData func(instance CFPlugInInstanceRef) unsafe.Pointer
var _cFPlugInInstanceGetInstanceDataErr error

func tryCFPlugInInstanceGetInstanceData(instance CFPlugInInstanceRef) (unsafe.Pointer, error) {
	if _cFPlugInInstanceGetInstanceData == nil {
		return nil, symbolCallError("CFPlugInInstanceGetInstanceData", "", _cFPlugInInstanceGetInstanceDataErr)
	}
	return _cFPlugInInstanceGetInstanceData(instance), nil
}

// CFPlugInInstanceGetInstanceData not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInstanceData(_:)
func CFPlugInInstanceGetInstanceData(instance CFPlugInInstanceRef) unsafe.Pointer {
	result, callErr := tryCFPlugInInstanceGetInstanceData(instance)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInInstanceGetInterfaceFunctionTable func(instance CFPlugInInstanceRef, interfaceName CFStringRef, ftbl unsafe.Pointer) bool
var _cFPlugInInstanceGetInterfaceFunctionTableErr error

func tryCFPlugInInstanceGetInterfaceFunctionTable(instance CFPlugInInstanceRef, interfaceName CFStringRef, ftbl unsafe.Pointer) (bool, error) {
	if _cFPlugInInstanceGetInterfaceFunctionTable == nil {
		return false, symbolCallError("CFPlugInInstanceGetInterfaceFunctionTable", "", _cFPlugInInstanceGetInterfaceFunctionTableErr)
	}
	return _cFPlugInInstanceGetInterfaceFunctionTable(instance, interfaceName, ftbl), nil
}

// CFPlugInInstanceGetInterfaceFunctionTable not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInterfaceFunctionTable(_:_:_:)
func CFPlugInInstanceGetInterfaceFunctionTable(instance CFPlugInInstanceRef, interfaceName CFStringRef, ftbl unsafe.Pointer) bool {
	result, callErr := tryCFPlugInInstanceGetInterfaceFunctionTable(instance, interfaceName, ftbl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInInstanceGetTypeID func() uint
var _cFPlugInInstanceGetTypeIDErr error

func tryCFPlugInInstanceGetTypeID() (uint, error) {
	if _cFPlugInInstanceGetTypeID == nil {
		return 0, symbolCallError("CFPlugInInstanceGetTypeID", "", _cFPlugInInstanceGetTypeIDErr)
	}
	return _cFPlugInInstanceGetTypeID(), nil
}

// CFPlugInInstanceGetTypeID not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetTypeID()
func CFPlugInInstanceGetTypeID() uint {
	result, callErr := tryCFPlugInInstanceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInIsLoadOnDemand func(plugIn CFPlugInRef) bool
var _cFPlugInIsLoadOnDemandErr error

func tryCFPlugInIsLoadOnDemand(plugIn CFPlugInRef) (bool, error) {
	if _cFPlugInIsLoadOnDemand == nil {
		return false, symbolCallError("CFPlugInIsLoadOnDemand", "", _cFPlugInIsLoadOnDemandErr)
	}
	return _cFPlugInIsLoadOnDemand(plugIn), nil
}

// CFPlugInIsLoadOnDemand determines whether or not a plug-in is loaded on demand.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInIsLoadOnDemand(_:)
func CFPlugInIsLoadOnDemand(plugIn CFPlugInRef) bool {
	result, callErr := tryCFPlugInIsLoadOnDemand(plugIn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInRegisterFactoryFunction func(factoryUUID CFUUIDRef, func_ CFPlugInFactoryFunction) bool
var _cFPlugInRegisterFactoryFunctionErr error

func tryCFPlugInRegisterFactoryFunction(factoryUUID CFUUIDRef, func_ CFPlugInFactoryFunction) (bool, error) {
	if _cFPlugInRegisterFactoryFunction == nil {
		return false, symbolCallError("CFPlugInRegisterFactoryFunction", "", _cFPlugInRegisterFactoryFunctionErr)
	}
	return _cFPlugInRegisterFactoryFunction(factoryUUID, func_), nil
}

// CFPlugInRegisterFactoryFunction registers a factory function and its UUID with a [CFPlugIn] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunction(_:_:)
func CFPlugInRegisterFactoryFunction(factoryUUID CFUUIDRef, func_ CFPlugInFactoryFunction) bool {
	result, callErr := tryCFPlugInRegisterFactoryFunction(factoryUUID, func_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInRegisterFactoryFunctionByName func(factoryUUID CFUUIDRef, plugIn CFPlugInRef, functionName CFStringRef) bool
var _cFPlugInRegisterFactoryFunctionByNameErr error

func tryCFPlugInRegisterFactoryFunctionByName(factoryUUID CFUUIDRef, plugIn CFPlugInRef, functionName CFStringRef) (bool, error) {
	if _cFPlugInRegisterFactoryFunctionByName == nil {
		return false, symbolCallError("CFPlugInRegisterFactoryFunctionByName", "", _cFPlugInRegisterFactoryFunctionByNameErr)
	}
	return _cFPlugInRegisterFactoryFunctionByName(factoryUUID, plugIn, functionName), nil
}

// CFPlugInRegisterFactoryFunctionByName registers a factory function with a [CFPlugIn] object using the function’s name instead of its UUID.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunctionByName(_:_:_:)
func CFPlugInRegisterFactoryFunctionByName(factoryUUID CFUUIDRef, plugIn CFPlugInRef, functionName CFStringRef) bool {
	result, callErr := tryCFPlugInRegisterFactoryFunctionByName(factoryUUID, plugIn, functionName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInRegisterPlugInType func(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool
var _cFPlugInRegisterPlugInTypeErr error

func tryCFPlugInRegisterPlugInType(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) (bool, error) {
	if _cFPlugInRegisterPlugInType == nil {
		return false, symbolCallError("CFPlugInRegisterPlugInType", "", _cFPlugInRegisterPlugInTypeErr)
	}
	return _cFPlugInRegisterPlugInType(factoryUUID, typeUUID), nil
}

// CFPlugInRegisterPlugInType registers a type and its corresponding factory function with a [CFPlugIn] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterPlugInType(_:_:)
func CFPlugInRegisterPlugInType(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool {
	result, callErr := tryCFPlugInRegisterPlugInType(factoryUUID, typeUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInRemoveInstanceForFactory func(factoryID CFUUIDRef)
var _cFPlugInRemoveInstanceForFactoryErr error

func tryCFPlugInRemoveInstanceForFactory(factoryID CFUUIDRef) error {
	if _cFPlugInRemoveInstanceForFactory == nil {
		return symbolCallError("CFPlugInRemoveInstanceForFactory", "", _cFPlugInRemoveInstanceForFactoryErr)
	}
	_cFPlugInRemoveInstanceForFactory(factoryID)
	return nil
}

// CFPlugInRemoveInstanceForFactory unregisters an instance of a type with [CFPlugIn].
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRemoveInstanceForFactory(_:)
func CFPlugInRemoveInstanceForFactory(factoryID CFUUIDRef) {
	if callErr := tryCFPlugInRemoveInstanceForFactory(factoryID); callErr != nil {
		panic(callErr)
	}
}

var _cFPlugInSetLoadOnDemand func(plugIn CFPlugInRef, flag bool)
var _cFPlugInSetLoadOnDemandErr error

func tryCFPlugInSetLoadOnDemand(plugIn CFPlugInRef, flag bool) error {
	if _cFPlugInSetLoadOnDemand == nil {
		return symbolCallError("CFPlugInSetLoadOnDemand", "", _cFPlugInSetLoadOnDemandErr)
	}
	_cFPlugInSetLoadOnDemand(plugIn, flag)
	return nil
}

// CFPlugInSetLoadOnDemand enables or disables load on demand for plug-ins that do dynamic registration (only when a client requests an instance of a supported type).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInSetLoadOnDemand(_:_:)
func CFPlugInSetLoadOnDemand(plugIn CFPlugInRef, flag bool) {
	if callErr := tryCFPlugInSetLoadOnDemand(plugIn, flag); callErr != nil {
		panic(callErr)
	}
}

var _cFPlugInUnregisterFactory func(factoryUUID CFUUIDRef) bool
var _cFPlugInUnregisterFactoryErr error

func tryCFPlugInUnregisterFactory(factoryUUID CFUUIDRef) (bool, error) {
	if _cFPlugInUnregisterFactory == nil {
		return false, symbolCallError("CFPlugInUnregisterFactory", "", _cFPlugInUnregisterFactoryErr)
	}
	return _cFPlugInUnregisterFactory(factoryUUID), nil
}

// CFPlugInUnregisterFactory removes the given function from a plug-in’s list of registered factory functions.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterFactory(_:)
func CFPlugInUnregisterFactory(factoryUUID CFUUIDRef) bool {
	result, callErr := tryCFPlugInUnregisterFactory(factoryUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPlugInUnregisterPlugInType func(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool
var _cFPlugInUnregisterPlugInTypeErr error

func tryCFPlugInUnregisterPlugInType(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) (bool, error) {
	if _cFPlugInUnregisterPlugInType == nil {
		return false, symbolCallError("CFPlugInUnregisterPlugInType", "", _cFPlugInUnregisterPlugInTypeErr)
	}
	return _cFPlugInUnregisterPlugInType(factoryUUID, typeUUID), nil
}

// CFPlugInUnregisterPlugInType removes the given type from a plug-in’s list of registered types.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterPlugInType(_:_:)
func CFPlugInUnregisterPlugInType(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool {
	result, callErr := tryCFPlugInUnregisterPlugInType(factoryUUID, typeUUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesAddSuitePreferencesToApp func(applicationID CFStringRef, suiteID CFStringRef)
var _cFPreferencesAddSuitePreferencesToAppErr error

func tryCFPreferencesAddSuitePreferencesToApp(applicationID CFStringRef, suiteID CFStringRef) error {
	if _cFPreferencesAddSuitePreferencesToApp == nil {
		return symbolCallError("CFPreferencesAddSuitePreferencesToApp", "", _cFPreferencesAddSuitePreferencesToAppErr)
	}
	_cFPreferencesAddSuitePreferencesToApp(applicationID, suiteID)
	return nil
}

// CFPreferencesAddSuitePreferencesToApp adds suite preferences to an application’s preference search chain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAddSuitePreferencesToApp(_:_:)
func CFPreferencesAddSuitePreferencesToApp(applicationID CFStringRef, suiteID CFStringRef) {
	if callErr := tryCFPreferencesAddSuitePreferencesToApp(applicationID, suiteID); callErr != nil {
		panic(callErr)
	}
}

var _cFPreferencesAppSynchronize func(applicationID CFStringRef) bool
var _cFPreferencesAppSynchronizeErr error

func tryCFPreferencesAppSynchronize(applicationID CFStringRef) (bool, error) {
	if _cFPreferencesAppSynchronize == nil {
		return false, symbolCallError("CFPreferencesAppSynchronize", "", _cFPreferencesAppSynchronizeErr)
	}
	return _cFPreferencesAppSynchronize(applicationID), nil
}

// CFPreferencesAppSynchronize writes to permanent storage all pending changes to the preference data for the application, and reads the latest preference data from permanent storage.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppSynchronize(_:)
func CFPreferencesAppSynchronize(applicationID CFStringRef) bool {
	result, callErr := tryCFPreferencesAppSynchronize(applicationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesAppValueIsForced func(key CFStringRef, applicationID CFStringRef) bool
var _cFPreferencesAppValueIsForcedErr error

func tryCFPreferencesAppValueIsForced(key CFStringRef, applicationID CFStringRef) (bool, error) {
	if _cFPreferencesAppValueIsForced == nil {
		return false, symbolCallError("CFPreferencesAppValueIsForced", "", _cFPreferencesAppValueIsForcedErr)
	}
	return _cFPreferencesAppValueIsForced(key, applicationID), nil
}

// CFPreferencesAppValueIsForced determines whether or not a given key has been imposed on the user.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppValueIsForced(_:_:)
func CFPreferencesAppValueIsForced(key CFStringRef, applicationID CFStringRef) bool {
	result, callErr := tryCFPreferencesAppValueIsForced(key, applicationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesCopyAppValue func(key CFStringRef, applicationID CFStringRef) CFPropertyListRef
var _cFPreferencesCopyAppValueErr error

func tryCFPreferencesCopyAppValue(key CFStringRef, applicationID CFStringRef) (CFPropertyListRef, error) {
	if _cFPreferencesCopyAppValue == nil {
		return 0, symbolCallError("CFPreferencesCopyAppValue", "", _cFPreferencesCopyAppValueErr)
	}
	return _cFPreferencesCopyAppValue(key, applicationID), nil
}

// CFPreferencesCopyAppValue obtains a preference value for the specified key and application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyAppValue(_:_:)
func CFPreferencesCopyAppValue(key CFStringRef, applicationID CFStringRef) CFPropertyListRef {
	result, callErr := tryCFPreferencesCopyAppValue(key, applicationID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesCopyKeyList func(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFArrayRef
var _cFPreferencesCopyKeyListErr error

func tryCFPreferencesCopyKeyList(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) (CFArrayRef, error) {
	if _cFPreferencesCopyKeyList == nil {
		return 0, symbolCallError("CFPreferencesCopyKeyList", "", _cFPreferencesCopyKeyListErr)
	}
	return _cFPreferencesCopyKeyList(applicationID, userName, hostName), nil
}

// CFPreferencesCopyKeyList constructs and returns the list of all keys set in the specified domain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyKeyList(_:_:_:)
func CFPreferencesCopyKeyList(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFArrayRef {
	result, callErr := tryCFPreferencesCopyKeyList(applicationID, userName, hostName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesCopyMultiple func(keysToFetch CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFDictionaryRef
var _cFPreferencesCopyMultipleErr error

func tryCFPreferencesCopyMultiple(keysToFetch CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) (CFDictionaryRef, error) {
	if _cFPreferencesCopyMultiple == nil {
		return 0, symbolCallError("CFPreferencesCopyMultiple", "", _cFPreferencesCopyMultipleErr)
	}
	return _cFPreferencesCopyMultiple(keysToFetch, applicationID, userName, hostName), nil
}

// CFPreferencesCopyMultiple returns a dictionary containing preference values for multiple keys.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyMultiple(_:_:_:_:)
func CFPreferencesCopyMultiple(keysToFetch CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFDictionaryRef {
	result, callErr := tryCFPreferencesCopyMultiple(keysToFetch, applicationID, userName, hostName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesCopyValue func(key CFStringRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFPropertyListRef
var _cFPreferencesCopyValueErr error

func tryCFPreferencesCopyValue(key CFStringRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) (CFPropertyListRef, error) {
	if _cFPreferencesCopyValue == nil {
		return 0, symbolCallError("CFPreferencesCopyValue", "", _cFPreferencesCopyValueErr)
	}
	return _cFPreferencesCopyValue(key, applicationID, userName, hostName), nil
}

// CFPreferencesCopyValue returns a preference value for a given domain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyValue(_:_:_:_:)
func CFPreferencesCopyValue(key CFStringRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFPropertyListRef {
	result, callErr := tryCFPreferencesCopyValue(key, applicationID, userName, hostName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesGetAppBooleanValue func(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) bool
var _cFPreferencesGetAppBooleanValueErr error

func tryCFPreferencesGetAppBooleanValue(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) (bool, error) {
	if _cFPreferencesGetAppBooleanValue == nil {
		return false, symbolCallError("CFPreferencesGetAppBooleanValue", "", _cFPreferencesGetAppBooleanValueErr)
	}
	return _cFPreferencesGetAppBooleanValue(key, applicationID, keyExistsAndHasValidFormat), nil
}

// CFPreferencesGetAppBooleanValue convenience function that directly obtains a Boolean preference value for the specified key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppBooleanValue(_:_:_:)
func CFPreferencesGetAppBooleanValue(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) bool {
	result, callErr := tryCFPreferencesGetAppBooleanValue(key, applicationID, keyExistsAndHasValidFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesGetAppIntegerValue func(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) int
var _cFPreferencesGetAppIntegerValueErr error

func tryCFPreferencesGetAppIntegerValue(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) (int, error) {
	if _cFPreferencesGetAppIntegerValue == nil {
		return 0, symbolCallError("CFPreferencesGetAppIntegerValue", "", _cFPreferencesGetAppIntegerValueErr)
	}
	return _cFPreferencesGetAppIntegerValue(key, applicationID, keyExistsAndHasValidFormat), nil
}

// CFPreferencesGetAppIntegerValue convenience function that directly obtains an integer preference value for the specified key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppIntegerValue(_:_:_:)
func CFPreferencesGetAppIntegerValue(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) int {
	result, callErr := tryCFPreferencesGetAppIntegerValue(key, applicationID, keyExistsAndHasValidFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPreferencesRemoveSuitePreferencesFromApp func(applicationID CFStringRef, suiteID CFStringRef)
var _cFPreferencesRemoveSuitePreferencesFromAppErr error

func tryCFPreferencesRemoveSuitePreferencesFromApp(applicationID CFStringRef, suiteID CFStringRef) error {
	if _cFPreferencesRemoveSuitePreferencesFromApp == nil {
		return symbolCallError("CFPreferencesRemoveSuitePreferencesFromApp", "", _cFPreferencesRemoveSuitePreferencesFromAppErr)
	}
	_cFPreferencesRemoveSuitePreferencesFromApp(applicationID, suiteID)
	return nil
}

// CFPreferencesRemoveSuitePreferencesFromApp removes suite preferences from an application’s search chain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesRemoveSuitePreferencesFromApp(_:_:)
func CFPreferencesRemoveSuitePreferencesFromApp(applicationID CFStringRef, suiteID CFStringRef) {
	if callErr := tryCFPreferencesRemoveSuitePreferencesFromApp(applicationID, suiteID); callErr != nil {
		panic(callErr)
	}
}

var _cFPreferencesSetAppValue func(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef)
var _cFPreferencesSetAppValueErr error

func tryCFPreferencesSetAppValue(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef) error {
	if _cFPreferencesSetAppValue == nil {
		return symbolCallError("CFPreferencesSetAppValue", "", _cFPreferencesSetAppValueErr)
	}
	_cFPreferencesSetAppValue(key, value, applicationID)
	return nil
}

// CFPreferencesSetAppValue adds, modifies, or removes a preference.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetAppValue(_:_:_:)
func CFPreferencesSetAppValue(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef) {
	if callErr := tryCFPreferencesSetAppValue(key, value, applicationID); callErr != nil {
		panic(callErr)
	}
}

var _cFPreferencesSetMultiple func(keysToSet CFDictionaryRef, keysToRemove CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef)
var _cFPreferencesSetMultipleErr error

func tryCFPreferencesSetMultiple(keysToSet CFDictionaryRef, keysToRemove CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) error {
	if _cFPreferencesSetMultiple == nil {
		return symbolCallError("CFPreferencesSetMultiple", "", _cFPreferencesSetMultipleErr)
	}
	_cFPreferencesSetMultiple(keysToSet, keysToRemove, applicationID, userName, hostName)
	return nil
}

// CFPreferencesSetMultiple convenience function that allows you to set and remove multiple preference values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetMultiple(_:_:_:_:_:)
func CFPreferencesSetMultiple(keysToSet CFDictionaryRef, keysToRemove CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) {
	if callErr := tryCFPreferencesSetMultiple(keysToSet, keysToRemove, applicationID, userName, hostName); callErr != nil {
		panic(callErr)
	}
}

var _cFPreferencesSetValue func(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef)
var _cFPreferencesSetValueErr error

func tryCFPreferencesSetValue(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) error {
	if _cFPreferencesSetValue == nil {
		return symbolCallError("CFPreferencesSetValue", "", _cFPreferencesSetValueErr)
	}
	_cFPreferencesSetValue(key, value, applicationID, userName, hostName)
	return nil
}

// CFPreferencesSetValue adds, modifies, or removes a preference value for the specified domain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetValue(_:_:_:_:_:)
func CFPreferencesSetValue(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) {
	if callErr := tryCFPreferencesSetValue(key, value, applicationID, userName, hostName); callErr != nil {
		panic(callErr)
	}
}

var _cFPreferencesSynchronize func(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) bool
var _cFPreferencesSynchronizeErr error

func tryCFPreferencesSynchronize(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) (bool, error) {
	if _cFPreferencesSynchronize == nil {
		return false, symbolCallError("CFPreferencesSynchronize", "", _cFPreferencesSynchronizeErr)
	}
	return _cFPreferencesSynchronize(applicationID, userName, hostName), nil
}

// CFPreferencesSynchronize for the specified domain, writes all pending changes to preference data to permanent storage, and reads latest preference data from permanent storage.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSynchronize(_:_:_:)
func CFPreferencesSynchronize(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) bool {
	result, callErr := tryCFPreferencesSynchronize(applicationID, userName, hostName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPropertyListCreateData func(allocator CFAllocatorRef, propertyList CFPropertyListRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) CFDataRef
var _cFPropertyListCreateDataErr error

func tryCFPropertyListCreateData(allocator CFAllocatorRef, propertyList CFPropertyListRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) (CFDataRef, error) {
	if _cFPropertyListCreateData == nil {
		return 0, symbolCallError("CFPropertyListCreateData", "10.6", _cFPropertyListCreateDataErr)
	}
	return _cFPropertyListCreateData(allocator, propertyList, format, options, err), nil
}

// CFPropertyListCreateData returns a CFData object containing a serialized representation of a given property list in a specified format.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateData(_:_:_:_:_:)
func CFPropertyListCreateData(allocator CFAllocatorRef, propertyList CFPropertyListRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) CFDataRef {
	result, callErr := tryCFPropertyListCreateData(allocator, propertyList, format, options, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPropertyListCreateDeepCopy func(allocator CFAllocatorRef, propertyList CFPropertyListRef, mutabilityOption uint64) CFPropertyListRef
var _cFPropertyListCreateDeepCopyErr error

func tryCFPropertyListCreateDeepCopy(allocator CFAllocatorRef, propertyList CFPropertyListRef, mutabilityOption uint64) (CFPropertyListRef, error) {
	if _cFPropertyListCreateDeepCopy == nil {
		return 0, symbolCallError("CFPropertyListCreateDeepCopy", "", _cFPropertyListCreateDeepCopyErr)
	}
	return _cFPropertyListCreateDeepCopy(allocator, propertyList, mutabilityOption), nil
}

// CFPropertyListCreateDeepCopy recursively creates a copy of a given property list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateDeepCopy(_:_:_:)
func CFPropertyListCreateDeepCopy(allocator CFAllocatorRef, propertyList CFPropertyListRef, mutabilityOption uint64) CFPropertyListRef {
	result, callErr := tryCFPropertyListCreateDeepCopy(allocator, propertyList, mutabilityOption)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPropertyListCreateWithData func(allocator CFAllocatorRef, data CFDataRef, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef
var _cFPropertyListCreateWithDataErr error

func tryCFPropertyListCreateWithData(allocator CFAllocatorRef, data CFDataRef, options uint64, format *CFPropertyListFormat, err *CFErrorRef) (CFPropertyListRef, error) {
	if _cFPropertyListCreateWithData == nil {
		return 0, symbolCallError("CFPropertyListCreateWithData", "10.6", _cFPropertyListCreateWithDataErr)
	}
	return _cFPropertyListCreateWithData(allocator, data, options, format, err), nil
}

// CFPropertyListCreateWithData creates a property list from a given CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithData(_:_:_:_:_:)
func CFPropertyListCreateWithData(allocator CFAllocatorRef, data CFDataRef, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef {
	result, callErr := tryCFPropertyListCreateWithData(allocator, data, options, format, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPropertyListCreateWithStream func(allocator CFAllocatorRef, stream CFReadStreamRef, streamLength int, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef
var _cFPropertyListCreateWithStreamErr error

func tryCFPropertyListCreateWithStream(allocator CFAllocatorRef, stream CFReadStreamRef, streamLength int, options uint64, format *CFPropertyListFormat, err *CFErrorRef) (CFPropertyListRef, error) {
	if _cFPropertyListCreateWithStream == nil {
		return 0, symbolCallError("CFPropertyListCreateWithStream", "10.6", _cFPropertyListCreateWithStreamErr)
	}
	return _cFPropertyListCreateWithStream(allocator, stream, streamLength, options, format, err), nil
}

// CFPropertyListCreateWithStream create and return a property list with a CFReadStream input.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithStream(_:_:_:_:_:_:)
func CFPropertyListCreateWithStream(allocator CFAllocatorRef, stream CFReadStreamRef, streamLength int, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef {
	result, callErr := tryCFPropertyListCreateWithStream(allocator, stream, streamLength, options, format, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPropertyListIsValid func(plist CFPropertyListRef, format CFPropertyListFormat) bool
var _cFPropertyListIsValidErr error

func tryCFPropertyListIsValid(plist CFPropertyListRef, format CFPropertyListFormat) (bool, error) {
	if _cFPropertyListIsValid == nil {
		return false, symbolCallError("CFPropertyListIsValid", "", _cFPropertyListIsValidErr)
	}
	return _cFPropertyListIsValid(plist, format), nil
}

// CFPropertyListIsValid determines if a property list is valid.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListIsValid(_:_:)
func CFPropertyListIsValid(plist CFPropertyListRef, format CFPropertyListFormat) bool {
	result, callErr := tryCFPropertyListIsValid(plist, format)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFPropertyListWrite func(propertyList CFPropertyListRef, stream CFWriteStreamRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) int
var _cFPropertyListWriteErr error

func tryCFPropertyListWrite(propertyList CFPropertyListRef, stream CFWriteStreamRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) (int, error) {
	if _cFPropertyListWrite == nil {
		return 0, symbolCallError("CFPropertyListWrite", "10.6", _cFPropertyListWriteErr)
	}
	return _cFPropertyListWrite(propertyList, stream, format, options, err), nil
}

// CFPropertyListWrite write the bytes of a serialized property list out to a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListWrite(_:_:_:_:_:)
func CFPropertyListWrite(propertyList CFPropertyListRef, stream CFWriteStreamRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) int {
	result, callErr := tryCFPropertyListWrite(propertyList, stream, format, options, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamClose func(stream CFReadStreamRef)
var _cFReadStreamCloseErr error

func tryCFReadStreamClose(stream CFReadStreamRef) error {
	if _cFReadStreamClose == nil {
		return symbolCallError("CFReadStreamClose", "", _cFReadStreamCloseErr)
	}
	_cFReadStreamClose(stream)
	return nil
}

// CFReadStreamClose closes a readable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamClose(_:)
func CFReadStreamClose(stream CFReadStreamRef) {
	if callErr := tryCFReadStreamClose(stream); callErr != nil {
		panic(callErr)
	}
}

var _cFReadStreamCopyDispatchQueue func(stream CFReadStreamRef) uintptr
var _cFReadStreamCopyDispatchQueueErr error

func tryCFReadStreamCopyDispatchQueue(stream CFReadStreamRef) (dispatch.Queue, error) {
	if _cFReadStreamCopyDispatchQueue == nil {
		return dispatch.QueueFromHandle(0), symbolCallError("CFReadStreamCopyDispatchQueue", "10.9", _cFReadStreamCopyDispatchQueueErr)
	}
	return dispatch.QueueFromHandle(_cFReadStreamCopyDispatchQueue(stream)), nil
}

// CFReadStreamCopyDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyDispatchQueue(_:)
func CFReadStreamCopyDispatchQueue(stream CFReadStreamRef) dispatch.Queue {
	result, callErr := tryCFReadStreamCopyDispatchQueue(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamCopyError func(stream CFReadStreamRef) CFErrorRef
var _cFReadStreamCopyErrorErr error

func tryCFReadStreamCopyError(stream CFReadStreamRef) (CFErrorRef, error) {
	if _cFReadStreamCopyError == nil {
		return 0, symbolCallError("CFReadStreamCopyError", "10.5", _cFReadStreamCopyErrorErr)
	}
	return _cFReadStreamCopyError(stream), nil
}

// CFReadStreamCopyError returns the error associated with a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyError(_:)
func CFReadStreamCopyError(stream CFReadStreamRef) CFErrorRef {
	result, callErr := tryCFReadStreamCopyError(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamCopyProperty func(stream CFReadStreamRef, propertyName CFStreamPropertyKey) CFTypeRef
var _cFReadStreamCopyPropertyErr error

func tryCFReadStreamCopyProperty(stream CFReadStreamRef, propertyName CFStreamPropertyKey) (CFTypeRef, error) {
	if _cFReadStreamCopyProperty == nil {
		return 0, symbolCallError("CFReadStreamCopyProperty", "", _cFReadStreamCopyPropertyErr)
	}
	return _cFReadStreamCopyProperty(stream, propertyName), nil
}

// CFReadStreamCopyProperty returns the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyProperty(_:_:)
func CFReadStreamCopyProperty(stream CFReadStreamRef, propertyName CFStreamPropertyKey) CFTypeRef {
	result, callErr := tryCFReadStreamCopyProperty(stream, propertyName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamCreateWithBytesNoCopy func(alloc CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFReadStreamRef
var _cFReadStreamCreateWithBytesNoCopyErr error

func tryCFReadStreamCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) (CFReadStreamRef, error) {
	if _cFReadStreamCreateWithBytesNoCopy == nil {
		return 0, symbolCallError("CFReadStreamCreateWithBytesNoCopy", "", _cFReadStreamCreateWithBytesNoCopyErr)
	}
	return _cFReadStreamCreateWithBytesNoCopy(alloc, bytes, length, bytesDeallocator), nil
}

// CFReadStreamCreateWithBytesNoCopy creates a readable stream for a block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithBytesNoCopy(_:_:_:_:)
func CFReadStreamCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFReadStreamRef {
	result, callErr := tryCFReadStreamCreateWithBytesNoCopy(alloc, bytes, length, bytesDeallocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamCreateWithFile func(alloc CFAllocatorRef, fileURL CFURLRef) CFReadStreamRef
var _cFReadStreamCreateWithFileErr error

func tryCFReadStreamCreateWithFile(alloc CFAllocatorRef, fileURL CFURLRef) (CFReadStreamRef, error) {
	if _cFReadStreamCreateWithFile == nil {
		return 0, symbolCallError("CFReadStreamCreateWithFile", "", _cFReadStreamCreateWithFileErr)
	}
	return _cFReadStreamCreateWithFile(alloc, fileURL), nil
}

// CFReadStreamCreateWithFile creates a readable stream for a file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithFile(_:_:)
func CFReadStreamCreateWithFile(alloc CFAllocatorRef, fileURL CFURLRef) CFReadStreamRef {
	result, callErr := tryCFReadStreamCreateWithFile(alloc, fileURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamGetBuffer func(stream CFReadStreamRef, maxBytesToRead int, numBytesRead *int) *uint8
var _cFReadStreamGetBufferErr error

func tryCFReadStreamGetBuffer(stream CFReadStreamRef, maxBytesToRead int, numBytesRead *int) (*uint8, error) {
	if _cFReadStreamGetBuffer == nil {
		return nil, symbolCallError("CFReadStreamGetBuffer", "", _cFReadStreamGetBufferErr)
	}
	return _cFReadStreamGetBuffer(stream, maxBytesToRead, numBytesRead), nil
}

// CFReadStreamGetBuffer returns a pointer to a stream’s internal buffer of unread data, if possible.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetBuffer(_:_:_:)
func CFReadStreamGetBuffer(stream CFReadStreamRef, maxBytesToRead int, numBytesRead *int) *uint8 {
	result, callErr := tryCFReadStreamGetBuffer(stream, maxBytesToRead, numBytesRead)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamGetError func(stream CFReadStreamRef) CFStreamError
var _cFReadStreamGetErrorErr error

func tryCFReadStreamGetError(stream CFReadStreamRef) (CFStreamError, error) {
	if _cFReadStreamGetError == nil {
		return CFStreamError{}, symbolCallError("CFReadStreamGetError", "", _cFReadStreamGetErrorErr)
	}
	return _cFReadStreamGetError(stream), nil
}

// CFReadStreamGetError returns the error status of a stream.
//
// Deprecated: Use [CFReadStreamCopyError(_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFReadStreamCopyError(_:)>) instead.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetError(_:)
func CFReadStreamGetError(stream CFReadStreamRef) CFStreamError {
	result, callErr := tryCFReadStreamGetError(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamGetStatus func(stream CFReadStreamRef) CFStreamStatus
var _cFReadStreamGetStatusErr error

func tryCFReadStreamGetStatus(stream CFReadStreamRef) (CFStreamStatus, error) {
	if _cFReadStreamGetStatus == nil {
		return *new(CFStreamStatus), symbolCallError("CFReadStreamGetStatus", "", _cFReadStreamGetStatusErr)
	}
	return _cFReadStreamGetStatus(stream), nil
}

// CFReadStreamGetStatus returns the current state of a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetStatus(_:)
func CFReadStreamGetStatus(stream CFReadStreamRef) CFStreamStatus {
	result, callErr := tryCFReadStreamGetStatus(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamGetTypeID func() uint
var _cFReadStreamGetTypeIDErr error

func tryCFReadStreamGetTypeID() (uint, error) {
	if _cFReadStreamGetTypeID == nil {
		return 0, symbolCallError("CFReadStreamGetTypeID", "", _cFReadStreamGetTypeIDErr)
	}
	return _cFReadStreamGetTypeID(), nil
}

// CFReadStreamGetTypeID returns the type identifier the [CFReadStream] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetTypeID()
func CFReadStreamGetTypeID() uint {
	result, callErr := tryCFReadStreamGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamHasBytesAvailable func(stream CFReadStreamRef) bool
var _cFReadStreamHasBytesAvailableErr error

func tryCFReadStreamHasBytesAvailable(stream CFReadStreamRef) (bool, error) {
	if _cFReadStreamHasBytesAvailable == nil {
		return false, symbolCallError("CFReadStreamHasBytesAvailable", "", _cFReadStreamHasBytesAvailableErr)
	}
	return _cFReadStreamHasBytesAvailable(stream), nil
}

// CFReadStreamHasBytesAvailable returns a Boolean value that indicates whether a readable stream has data that can be read without blocking.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamHasBytesAvailable(_:)
func CFReadStreamHasBytesAvailable(stream CFReadStreamRef) bool {
	result, callErr := tryCFReadStreamHasBytesAvailable(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamOpen func(stream CFReadStreamRef) bool
var _cFReadStreamOpenErr error

func tryCFReadStreamOpen(stream CFReadStreamRef) (bool, error) {
	if _cFReadStreamOpen == nil {
		return false, symbolCallError("CFReadStreamOpen", "", _cFReadStreamOpenErr)
	}
	return _cFReadStreamOpen(stream), nil
}

// CFReadStreamOpen opens a stream for reading.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamOpen(_:)
func CFReadStreamOpen(stream CFReadStreamRef) bool {
	result, callErr := tryCFReadStreamOpen(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamRead func(stream CFReadStreamRef, buffer *uint8, bufferLength int) int
var _cFReadStreamReadErr error

func tryCFReadStreamRead(stream CFReadStreamRef, buffer *uint8, bufferLength int) (int, error) {
	if _cFReadStreamRead == nil {
		return 0, symbolCallError("CFReadStreamRead", "", _cFReadStreamReadErr)
	}
	return _cFReadStreamRead(stream, buffer, bufferLength), nil
}

// CFReadStreamRead reads data from a readable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamRead(_:_:_:)
func CFReadStreamRead(stream CFReadStreamRef, buffer *uint8, bufferLength int) int {
	result, callErr := tryCFReadStreamRead(stream, buffer, bufferLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamScheduleWithRunLoop func(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)
var _cFReadStreamScheduleWithRunLoopErr error

func tryCFReadStreamScheduleWithRunLoop(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) error {
	if _cFReadStreamScheduleWithRunLoop == nil {
		return symbolCallError("CFReadStreamScheduleWithRunLoop", "", _cFReadStreamScheduleWithRunLoopErr)
	}
	_cFReadStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
	return nil
}

// CFReadStreamScheduleWithRunLoop schedules a stream into a run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamScheduleWithRunLoop(_:_:_:)
func CFReadStreamScheduleWithRunLoop(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) {
	if callErr := tryCFReadStreamScheduleWithRunLoop(stream, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _cFReadStreamSetClient func(stream CFReadStreamRef, streamEvents uint64, clientCB CFReadStreamClientCallBack, clientContext *CFStreamClientContext) bool
var _cFReadStreamSetClientErr error

func tryCFReadStreamSetClient(stream CFReadStreamRef, streamEvents uint64, clientCB CFReadStreamClientCallBack, clientContext *CFStreamClientContext) (bool, error) {
	if _cFReadStreamSetClient == nil {
		return false, symbolCallError("CFReadStreamSetClient", "", _cFReadStreamSetClientErr)
	}
	return _cFReadStreamSetClient(stream, streamEvents, clientCB, clientContext), nil
}

// CFReadStreamSetClient assigns a client to a stream, which receives callbacks when certain events occur.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetClient(_:_:_:_:)
func CFReadStreamSetClient(stream CFReadStreamRef, streamEvents uint64, clientCB CFReadStreamClientCallBack, clientContext *CFStreamClientContext) bool {
	result, callErr := tryCFReadStreamSetClient(stream, streamEvents, clientCB, clientContext)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamSetDispatchQueue func(stream CFReadStreamRef, q uintptr)
var _cFReadStreamSetDispatchQueueErr error

func tryCFReadStreamSetDispatchQueue(stream CFReadStreamRef, q dispatch.Queue) error {
	if _cFReadStreamSetDispatchQueue == nil {
		return symbolCallError("CFReadStreamSetDispatchQueue", "10.9", _cFReadStreamSetDispatchQueueErr)
	}
	_cFReadStreamSetDispatchQueue(stream, uintptr(q.Handle()))
	return nil
}

// CFReadStreamSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetDispatchQueue(_:_:)
func CFReadStreamSetDispatchQueue(stream CFReadStreamRef, q dispatch.Queue) {
	if callErr := tryCFReadStreamSetDispatchQueue(stream, q); callErr != nil {
		panic(callErr)
	}
}

var _cFReadStreamSetProperty func(stream CFReadStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool
var _cFReadStreamSetPropertyErr error

func tryCFReadStreamSetProperty(stream CFReadStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) (bool, error) {
	if _cFReadStreamSetProperty == nil {
		return false, symbolCallError("CFReadStreamSetProperty", "", _cFReadStreamSetPropertyErr)
	}
	return _cFReadStreamSetProperty(stream, propertyName, propertyValue), nil
}

// CFReadStreamSetProperty sets the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetProperty(_:_:_:)
func CFReadStreamSetProperty(stream CFReadStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool {
	result, callErr := tryCFReadStreamSetProperty(stream, propertyName, propertyValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFReadStreamUnscheduleFromRunLoop func(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)
var _cFReadStreamUnscheduleFromRunLoopErr error

func tryCFReadStreamUnscheduleFromRunLoop(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) error {
	if _cFReadStreamUnscheduleFromRunLoop == nil {
		return symbolCallError("CFReadStreamUnscheduleFromRunLoop", "", _cFReadStreamUnscheduleFromRunLoopErr)
	}
	_cFReadStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
	return nil
}

// CFReadStreamUnscheduleFromRunLoop removes a read stream from a given run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamUnscheduleFromRunLoop(_:_:_:)
func CFReadStreamUnscheduleFromRunLoop(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) {
	if callErr := tryCFReadStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _cFRelease func(cf CFTypeRef)
var _cFReleaseErr error

func tryCFRelease(cf CFTypeRef) error {
	if _cFRelease == nil {
		return symbolCallError("CFRelease", "", _cFReleaseErr)
	}
	_cFRelease(cf)
	return nil
}

// CFRelease releases a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRelease
func CFRelease(cf CFTypeRef) {
	if callErr := tryCFRelease(cf); callErr != nil {
		panic(callErr)
	}
}

var _cFRetain func(cf CFTypeRef) CFTypeRef
var _cFRetainErr error

func tryCFRetain(cf CFTypeRef) (CFTypeRef, error) {
	if _cFRetain == nil {
		return 0, symbolCallError("CFRetain", "", _cFRetainErr)
	}
	return _cFRetain(cf), nil
}

// CFRetain retains a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRetain
func CFRetain(cf CFTypeRef) CFTypeRef {
	result, callErr := tryCFRetain(cf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopAddCommonMode func(rl CFRunLoopRef, mode CFRunLoopMode)
var _cFRunLoopAddCommonModeErr error

func tryCFRunLoopAddCommonMode(rl CFRunLoopRef, mode CFRunLoopMode) error {
	if _cFRunLoopAddCommonMode == nil {
		return symbolCallError("CFRunLoopAddCommonMode", "", _cFRunLoopAddCommonModeErr)
	}
	_cFRunLoopAddCommonMode(rl, mode)
	return nil
}

// CFRunLoopAddCommonMode adds a mode to the set of run loop common modes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddCommonMode(_:_:)
func CFRunLoopAddCommonMode(rl CFRunLoopRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopAddCommonMode(rl, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopAddObserver func(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode)
var _cFRunLoopAddObserverErr error

func tryCFRunLoopAddObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) error {
	if _cFRunLoopAddObserver == nil {
		return symbolCallError("CFRunLoopAddObserver", "", _cFRunLoopAddObserverErr)
	}
	_cFRunLoopAddObserver(rl, observer, mode)
	return nil
}

// CFRunLoopAddObserver adds a CFRunLoopObserver object to a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddObserver(_:_:_:)
func CFRunLoopAddObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopAddObserver(rl, observer, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopAddSource func(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode)
var _cFRunLoopAddSourceErr error

func tryCFRunLoopAddSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) error {
	if _cFRunLoopAddSource == nil {
		return symbolCallError("CFRunLoopAddSource", "", _cFRunLoopAddSourceErr)
	}
	_cFRunLoopAddSource(rl, source, mode)
	return nil
}

// CFRunLoopAddSource adds a CFRunLoopSource object to a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddSource(_:_:_:)
func CFRunLoopAddSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopAddSource(rl, source, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopAddTimer func(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode)
var _cFRunLoopAddTimerErr error

func tryCFRunLoopAddTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) error {
	if _cFRunLoopAddTimer == nil {
		return symbolCallError("CFRunLoopAddTimer", "", _cFRunLoopAddTimerErr)
	}
	_cFRunLoopAddTimer(rl, timer, mode)
	return nil
}

// CFRunLoopAddTimer adds a CFRunLoopTimer object to a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddTimer(_:_:_:)
func CFRunLoopAddTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopAddTimer(rl, timer, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopContainsObserver func(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) bool
var _cFRunLoopContainsObserverErr error

func tryCFRunLoopContainsObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) (bool, error) {
	if _cFRunLoopContainsObserver == nil {
		return false, symbolCallError("CFRunLoopContainsObserver", "", _cFRunLoopContainsObserverErr)
	}
	return _cFRunLoopContainsObserver(rl, observer, mode), nil
}

// CFRunLoopContainsObserver returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopObserver object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsObserver(_:_:_:)
func CFRunLoopContainsObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) bool {
	result, callErr := tryCFRunLoopContainsObserver(rl, observer, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopContainsSource func(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) bool
var _cFRunLoopContainsSourceErr error

func tryCFRunLoopContainsSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) (bool, error) {
	if _cFRunLoopContainsSource == nil {
		return false, symbolCallError("CFRunLoopContainsSource", "", _cFRunLoopContainsSourceErr)
	}
	return _cFRunLoopContainsSource(rl, source, mode), nil
}

// CFRunLoopContainsSource returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsSource(_:_:_:)
func CFRunLoopContainsSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) bool {
	result, callErr := tryCFRunLoopContainsSource(rl, source, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopContainsTimer func(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) bool
var _cFRunLoopContainsTimerErr error

func tryCFRunLoopContainsTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) (bool, error) {
	if _cFRunLoopContainsTimer == nil {
		return false, symbolCallError("CFRunLoopContainsTimer", "", _cFRunLoopContainsTimerErr)
	}
	return _cFRunLoopContainsTimer(rl, timer, mode), nil
}

// CFRunLoopContainsTimer returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsTimer(_:_:_:)
func CFRunLoopContainsTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) bool {
	result, callErr := tryCFRunLoopContainsTimer(rl, timer, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopCopyAllModes func(rl CFRunLoopRef) CFArrayRef
var _cFRunLoopCopyAllModesErr error

func tryCFRunLoopCopyAllModes(rl CFRunLoopRef) (CFArrayRef, error) {
	if _cFRunLoopCopyAllModes == nil {
		return 0, symbolCallError("CFRunLoopCopyAllModes", "", _cFRunLoopCopyAllModesErr)
	}
	return _cFRunLoopCopyAllModes(rl), nil
}

// CFRunLoopCopyAllModes returns an array that contains all the defined modes for a CFRunLoop object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyAllModes(_:)
func CFRunLoopCopyAllModes(rl CFRunLoopRef) CFArrayRef {
	result, callErr := tryCFRunLoopCopyAllModes(rl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopCopyCurrentMode func(rl CFRunLoopRef) CFRunLoopMode
var _cFRunLoopCopyCurrentModeErr error

func tryCFRunLoopCopyCurrentMode(rl CFRunLoopRef) (CFRunLoopMode, error) {
	if _cFRunLoopCopyCurrentMode == nil {
		return *new(CFRunLoopMode), symbolCallError("CFRunLoopCopyCurrentMode", "", _cFRunLoopCopyCurrentModeErr)
	}
	return _cFRunLoopCopyCurrentMode(rl), nil
}

// CFRunLoopCopyCurrentMode returns the name of the mode in which a given run loop is currently running.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyCurrentMode(_:)
func CFRunLoopCopyCurrentMode(rl CFRunLoopRef) CFRunLoopMode {
	result, callErr := tryCFRunLoopCopyCurrentMode(rl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopGetCurrent func() CFRunLoopRef
var _cFRunLoopGetCurrentErr error

func tryCFRunLoopGetCurrent() (CFRunLoopRef, error) {
	if _cFRunLoopGetCurrent == nil {
		return 0, symbolCallError("CFRunLoopGetCurrent", "", _cFRunLoopGetCurrentErr)
	}
	return _cFRunLoopGetCurrent(), nil
}

// CFRunLoopGetCurrent returns the CFRunLoop object for the current thread.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetCurrent()
func CFRunLoopGetCurrent() CFRunLoopRef {
	result, callErr := tryCFRunLoopGetCurrent()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopGetMain func() CFRunLoopRef
var _cFRunLoopGetMainErr error

func tryCFRunLoopGetMain() (CFRunLoopRef, error) {
	if _cFRunLoopGetMain == nil {
		return 0, symbolCallError("CFRunLoopGetMain", "", _cFRunLoopGetMainErr)
	}
	return _cFRunLoopGetMain(), nil
}

// CFRunLoopGetMain returns the main CFRunLoop object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetMain()
func CFRunLoopGetMain() CFRunLoopRef {
	result, callErr := tryCFRunLoopGetMain()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopGetNextTimerFireDate func(rl CFRunLoopRef, mode CFRunLoopMode) CFAbsoluteTime
var _cFRunLoopGetNextTimerFireDateErr error

func tryCFRunLoopGetNextTimerFireDate(rl CFRunLoopRef, mode CFRunLoopMode) (CFAbsoluteTime, error) {
	if _cFRunLoopGetNextTimerFireDate == nil {
		return *new(CFAbsoluteTime), symbolCallError("CFRunLoopGetNextTimerFireDate", "", _cFRunLoopGetNextTimerFireDateErr)
	}
	return _cFRunLoopGetNextTimerFireDate(rl, mode), nil
}

// CFRunLoopGetNextTimerFireDate returns the time at which the next timer will fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetNextTimerFireDate(_:_:)
func CFRunLoopGetNextTimerFireDate(rl CFRunLoopRef, mode CFRunLoopMode) CFAbsoluteTime {
	result, callErr := tryCFRunLoopGetNextTimerFireDate(rl, mode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopGetTypeID func() uint
var _cFRunLoopGetTypeIDErr error

func tryCFRunLoopGetTypeID() (uint, error) {
	if _cFRunLoopGetTypeID == nil {
		return 0, symbolCallError("CFRunLoopGetTypeID", "", _cFRunLoopGetTypeIDErr)
	}
	return _cFRunLoopGetTypeID(), nil
}

// CFRunLoopGetTypeID returns the type identifier for the CFRunLoop opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetTypeID()
func CFRunLoopGetTypeID() uint {
	result, callErr := tryCFRunLoopGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopIsWaiting func(rl CFRunLoopRef) bool
var _cFRunLoopIsWaitingErr error

func tryCFRunLoopIsWaiting(rl CFRunLoopRef) (bool, error) {
	if _cFRunLoopIsWaiting == nil {
		return false, symbolCallError("CFRunLoopIsWaiting", "", _cFRunLoopIsWaitingErr)
	}
	return _cFRunLoopIsWaiting(rl), nil
}

// CFRunLoopIsWaiting returns a Boolean value that indicates whether the run loop is waiting for an event.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopIsWaiting(_:)
func CFRunLoopIsWaiting(rl CFRunLoopRef) bool {
	result, callErr := tryCFRunLoopIsWaiting(rl)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverCreate func(allocator CFAllocatorRef, activities uint64, repeats bool, order int, callout CFRunLoopObserverCallBack, context *CFRunLoopObserverContext) CFRunLoopObserverRef
var _cFRunLoopObserverCreateErr error

func tryCFRunLoopObserverCreate(allocator CFAllocatorRef, activities uint64, repeats bool, order int, callout CFRunLoopObserverCallBack, context *CFRunLoopObserverContext) (CFRunLoopObserverRef, error) {
	if _cFRunLoopObserverCreate == nil {
		return 0, symbolCallError("CFRunLoopObserverCreate", "", _cFRunLoopObserverCreateErr)
	}
	return _cFRunLoopObserverCreate(allocator, activities, repeats, order, callout, context), nil
}

// CFRunLoopObserverCreate creates a CFRunLoopObserver object with a function callback.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreate(_:_:_:_:_:_:)
func CFRunLoopObserverCreate(allocator CFAllocatorRef, activities uint64, repeats bool, order int, callout CFRunLoopObserverCallBack, context *CFRunLoopObserverContext) CFRunLoopObserverRef {
	result, callErr := tryCFRunLoopObserverCreate(allocator, activities, repeats, order, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverCreateWithHandler func(allocator CFAllocatorRef, activities uint64, repeats bool, order int) CFRunLoopObserverRef
var _cFRunLoopObserverCreateWithHandlerErr error

func tryCFRunLoopObserverCreateWithHandler(allocator CFAllocatorRef, activities uint64, repeats bool, order int) (CFRunLoopObserverRef, error) {
	if _cFRunLoopObserverCreateWithHandler == nil {
		return 0, symbolCallError("CFRunLoopObserverCreateWithHandler", "10.7", _cFRunLoopObserverCreateWithHandlerErr)
	}
	return _cFRunLoopObserverCreateWithHandler(allocator, activities, repeats, order), nil
}

// CFRunLoopObserverCreateWithHandler creates a CFRunLoopObserver object with a block-based handler.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreateWithHandler(_:_:_:_:_:)
func CFRunLoopObserverCreateWithHandler(allocator CFAllocatorRef, activities uint64, repeats bool, order int) CFRunLoopObserverRef {
	result, callErr := tryCFRunLoopObserverCreateWithHandler(allocator, activities, repeats, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverDoesRepeat func(observer CFRunLoopObserverRef) bool
var _cFRunLoopObserverDoesRepeatErr error

func tryCFRunLoopObserverDoesRepeat(observer CFRunLoopObserverRef) (bool, error) {
	if _cFRunLoopObserverDoesRepeat == nil {
		return false, symbolCallError("CFRunLoopObserverDoesRepeat", "", _cFRunLoopObserverDoesRepeatErr)
	}
	return _cFRunLoopObserverDoesRepeat(observer), nil
}

// CFRunLoopObserverDoesRepeat returns a Boolean value that indicates whether a CFRunLoopObserver repeats.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverDoesRepeat(_:)
func CFRunLoopObserverDoesRepeat(observer CFRunLoopObserverRef) bool {
	result, callErr := tryCFRunLoopObserverDoesRepeat(observer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverGetActivities func(observer CFRunLoopObserverRef) uint64
var _cFRunLoopObserverGetActivitiesErr error

func tryCFRunLoopObserverGetActivities(observer CFRunLoopObserverRef) (uint64, error) {
	if _cFRunLoopObserverGetActivities == nil {
		return 0, symbolCallError("CFRunLoopObserverGetActivities", "", _cFRunLoopObserverGetActivitiesErr)
	}
	return _cFRunLoopObserverGetActivities(observer), nil
}

// CFRunLoopObserverGetActivities returns the run loop stages during which an observer runs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetActivities(_:)
func CFRunLoopObserverGetActivities(observer CFRunLoopObserverRef) uint64 {
	result, callErr := tryCFRunLoopObserverGetActivities(observer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverGetContext func(observer CFRunLoopObserverRef, context *CFRunLoopObserverContext)
var _cFRunLoopObserverGetContextErr error

func tryCFRunLoopObserverGetContext(observer CFRunLoopObserverRef, context *CFRunLoopObserverContext) error {
	if _cFRunLoopObserverGetContext == nil {
		return symbolCallError("CFRunLoopObserverGetContext", "", _cFRunLoopObserverGetContextErr)
	}
	_cFRunLoopObserverGetContext(observer, context)
	return nil
}

// CFRunLoopObserverGetContext returns the context information for a CFRunLoopObserver object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetContext(_:_:)
func CFRunLoopObserverGetContext(observer CFRunLoopObserverRef, context *CFRunLoopObserverContext) {
	if callErr := tryCFRunLoopObserverGetContext(observer, context); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopObserverGetOrder func(observer CFRunLoopObserverRef) int
var _cFRunLoopObserverGetOrderErr error

func tryCFRunLoopObserverGetOrder(observer CFRunLoopObserverRef) (int, error) {
	if _cFRunLoopObserverGetOrder == nil {
		return 0, symbolCallError("CFRunLoopObserverGetOrder", "", _cFRunLoopObserverGetOrderErr)
	}
	return _cFRunLoopObserverGetOrder(observer), nil
}

// CFRunLoopObserverGetOrder returns the ordering parameter for a CFRunLoopObserver object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetOrder(_:)
func CFRunLoopObserverGetOrder(observer CFRunLoopObserverRef) int {
	result, callErr := tryCFRunLoopObserverGetOrder(observer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverGetTypeID func() uint
var _cFRunLoopObserverGetTypeIDErr error

func tryCFRunLoopObserverGetTypeID() (uint, error) {
	if _cFRunLoopObserverGetTypeID == nil {
		return 0, symbolCallError("CFRunLoopObserverGetTypeID", "", _cFRunLoopObserverGetTypeIDErr)
	}
	return _cFRunLoopObserverGetTypeID(), nil
}

// CFRunLoopObserverGetTypeID returns the type identifier for the CFRunLoopObserver opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetTypeID()
func CFRunLoopObserverGetTypeID() uint {
	result, callErr := tryCFRunLoopObserverGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopObserverInvalidate func(observer CFRunLoopObserverRef)
var _cFRunLoopObserverInvalidateErr error

func tryCFRunLoopObserverInvalidate(observer CFRunLoopObserverRef) error {
	if _cFRunLoopObserverInvalidate == nil {
		return symbolCallError("CFRunLoopObserverInvalidate", "", _cFRunLoopObserverInvalidateErr)
	}
	_cFRunLoopObserverInvalidate(observer)
	return nil
}

// CFRunLoopObserverInvalidate invalidates a CFRunLoopObserver object, stopping it from ever firing again.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverInvalidate(_:)
func CFRunLoopObserverInvalidate(observer CFRunLoopObserverRef) {
	if callErr := tryCFRunLoopObserverInvalidate(observer); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopObserverIsValid func(observer CFRunLoopObserverRef) bool
var _cFRunLoopObserverIsValidErr error

func tryCFRunLoopObserverIsValid(observer CFRunLoopObserverRef) (bool, error) {
	if _cFRunLoopObserverIsValid == nil {
		return false, symbolCallError("CFRunLoopObserverIsValid", "", _cFRunLoopObserverIsValidErr)
	}
	return _cFRunLoopObserverIsValid(observer), nil
}

// CFRunLoopObserverIsValid returns a Boolean value that indicates whether a CFRunLoopObserver object is valid and able to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverIsValid(_:)
func CFRunLoopObserverIsValid(observer CFRunLoopObserverRef) bool {
	result, callErr := tryCFRunLoopObserverIsValid(observer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopPerformBlock func(rl CFRunLoopRef, mode CFTypeRef)
var _cFRunLoopPerformBlockErr error

func tryCFRunLoopPerformBlock(rl CFRunLoopRef, mode CFTypeRef) error {
	if _cFRunLoopPerformBlock == nil {
		return symbolCallError("CFRunLoopPerformBlock", "10.6", _cFRunLoopPerformBlockErr)
	}
	_cFRunLoopPerformBlock(rl, mode)
	return nil
}

// CFRunLoopPerformBlock enqueues a block object on a given runloop to be executed as the runloop cycles in specified modes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopPerformBlock(_:_:_:)
func CFRunLoopPerformBlock(rl CFRunLoopRef, mode CFTypeRef) {
	if callErr := tryCFRunLoopPerformBlock(rl, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopRemoveObserver func(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode)
var _cFRunLoopRemoveObserverErr error

func tryCFRunLoopRemoveObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) error {
	if _cFRunLoopRemoveObserver == nil {
		return symbolCallError("CFRunLoopRemoveObserver", "", _cFRunLoopRemoveObserverErr)
	}
	_cFRunLoopRemoveObserver(rl, observer, mode)
	return nil
}

// CFRunLoopRemoveObserver removes a CFRunLoopObserver object from a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveObserver(_:_:_:)
func CFRunLoopRemoveObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopRemoveObserver(rl, observer, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopRemoveSource func(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode)
var _cFRunLoopRemoveSourceErr error

func tryCFRunLoopRemoveSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) error {
	if _cFRunLoopRemoveSource == nil {
		return symbolCallError("CFRunLoopRemoveSource", "", _cFRunLoopRemoveSourceErr)
	}
	_cFRunLoopRemoveSource(rl, source, mode)
	return nil
}

// CFRunLoopRemoveSource removes a CFRunLoopSource object from a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveSource(_:_:_:)
func CFRunLoopRemoveSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopRemoveSource(rl, source, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopRemoveTimer func(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode)
var _cFRunLoopRemoveTimerErr error

func tryCFRunLoopRemoveTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) error {
	if _cFRunLoopRemoveTimer == nil {
		return symbolCallError("CFRunLoopRemoveTimer", "", _cFRunLoopRemoveTimerErr)
	}
	_cFRunLoopRemoveTimer(rl, timer, mode)
	return nil
}

// CFRunLoopRemoveTimer removes a CFRunLoopTimer object from a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveTimer(_:_:_:)
func CFRunLoopRemoveTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) {
	if callErr := tryCFRunLoopRemoveTimer(rl, timer, mode); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopRun func()
var _cFRunLoopRunErr error

func tryCFRunLoopRun() error {
	if _cFRunLoopRun == nil {
		return symbolCallError("CFRunLoopRun", "", _cFRunLoopRunErr)
	}
	_cFRunLoopRun()
	return nil
}

// CFRunLoopRun runs the current thread’s CFRunLoop object in its default mode indefinitely.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRun()
func CFRunLoopRun() {
	if callErr := tryCFRunLoopRun(); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopRunInMode func(mode CFRunLoopMode, seconds float64, returnAfterSourceHandled bool) CFRunLoopRunResult
var _cFRunLoopRunInModeErr error

func tryCFRunLoopRunInMode(mode CFRunLoopMode, seconds float64, returnAfterSourceHandled bool) (CFRunLoopRunResult, error) {
	if _cFRunLoopRunInMode == nil {
		return *new(CFRunLoopRunResult), symbolCallError("CFRunLoopRunInMode", "", _cFRunLoopRunInModeErr)
	}
	return _cFRunLoopRunInMode(mode, seconds, returnAfterSourceHandled), nil
}

// CFRunLoopRunInMode runs the current thread’s CFRunLoop object in a particular mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunInMode(_:_:_:)
func CFRunLoopRunInMode(mode CFRunLoopMode, seconds float64, returnAfterSourceHandled bool) CFRunLoopRunResult {
	result, callErr := tryCFRunLoopRunInMode(mode, seconds, returnAfterSourceHandled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopSourceCreate func(allocator CFAllocatorRef, order int, context *CFRunLoopSourceContext) CFRunLoopSourceRef
var _cFRunLoopSourceCreateErr error

func tryCFRunLoopSourceCreate(allocator CFAllocatorRef, order int, context *CFRunLoopSourceContext) (CFRunLoopSourceRef, error) {
	if _cFRunLoopSourceCreate == nil {
		return 0, symbolCallError("CFRunLoopSourceCreate", "", _cFRunLoopSourceCreateErr)
	}
	return _cFRunLoopSourceCreate(allocator, order, context), nil
}

// CFRunLoopSourceCreate creates a CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceCreate(_:_:_:)
func CFRunLoopSourceCreate(allocator CFAllocatorRef, order int, context *CFRunLoopSourceContext) CFRunLoopSourceRef {
	result, callErr := tryCFRunLoopSourceCreate(allocator, order, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopSourceGetContext func(source CFRunLoopSourceRef, context *CFRunLoopSourceContext)
var _cFRunLoopSourceGetContextErr error

func tryCFRunLoopSourceGetContext(source CFRunLoopSourceRef, context *CFRunLoopSourceContext) error {
	if _cFRunLoopSourceGetContext == nil {
		return symbolCallError("CFRunLoopSourceGetContext", "", _cFRunLoopSourceGetContextErr)
	}
	_cFRunLoopSourceGetContext(source, context)
	return nil
}

// CFRunLoopSourceGetContext returns the context information for a CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetContext(_:_:)
func CFRunLoopSourceGetContext(source CFRunLoopSourceRef, context *CFRunLoopSourceContext) {
	if callErr := tryCFRunLoopSourceGetContext(source, context); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopSourceGetOrder func(source CFRunLoopSourceRef) int
var _cFRunLoopSourceGetOrderErr error

func tryCFRunLoopSourceGetOrder(source CFRunLoopSourceRef) (int, error) {
	if _cFRunLoopSourceGetOrder == nil {
		return 0, symbolCallError("CFRunLoopSourceGetOrder", "", _cFRunLoopSourceGetOrderErr)
	}
	return _cFRunLoopSourceGetOrder(source), nil
}

// CFRunLoopSourceGetOrder returns the ordering parameter for a CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetOrder(_:)
func CFRunLoopSourceGetOrder(source CFRunLoopSourceRef) int {
	result, callErr := tryCFRunLoopSourceGetOrder(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopSourceGetTypeID func() uint
var _cFRunLoopSourceGetTypeIDErr error

func tryCFRunLoopSourceGetTypeID() (uint, error) {
	if _cFRunLoopSourceGetTypeID == nil {
		return 0, symbolCallError("CFRunLoopSourceGetTypeID", "", _cFRunLoopSourceGetTypeIDErr)
	}
	return _cFRunLoopSourceGetTypeID(), nil
}

// CFRunLoopSourceGetTypeID returns the type identifier of the CFRunLoopSource opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetTypeID()
func CFRunLoopSourceGetTypeID() uint {
	result, callErr := tryCFRunLoopSourceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopSourceInvalidate func(source CFRunLoopSourceRef)
var _cFRunLoopSourceInvalidateErr error

func tryCFRunLoopSourceInvalidate(source CFRunLoopSourceRef) error {
	if _cFRunLoopSourceInvalidate == nil {
		return symbolCallError("CFRunLoopSourceInvalidate", "", _cFRunLoopSourceInvalidateErr)
	}
	_cFRunLoopSourceInvalidate(source)
	return nil
}

// CFRunLoopSourceInvalidate invalidates a CFRunLoopSource object, stopping it from ever firing again.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceInvalidate(_:)
func CFRunLoopSourceInvalidate(source CFRunLoopSourceRef) {
	if callErr := tryCFRunLoopSourceInvalidate(source); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopSourceIsValid func(source CFRunLoopSourceRef) bool
var _cFRunLoopSourceIsValidErr error

func tryCFRunLoopSourceIsValid(source CFRunLoopSourceRef) (bool, error) {
	if _cFRunLoopSourceIsValid == nil {
		return false, symbolCallError("CFRunLoopSourceIsValid", "", _cFRunLoopSourceIsValidErr)
	}
	return _cFRunLoopSourceIsValid(source), nil
}

// CFRunLoopSourceIsValid returns a Boolean value that indicates whether a CFRunLoopSource object is valid and able to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceIsValid(_:)
func CFRunLoopSourceIsValid(source CFRunLoopSourceRef) bool {
	result, callErr := tryCFRunLoopSourceIsValid(source)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopSourceSignal func(source CFRunLoopSourceRef)
var _cFRunLoopSourceSignalErr error

func tryCFRunLoopSourceSignal(source CFRunLoopSourceRef) error {
	if _cFRunLoopSourceSignal == nil {
		return symbolCallError("CFRunLoopSourceSignal", "", _cFRunLoopSourceSignalErr)
	}
	_cFRunLoopSourceSignal(source)
	return nil
}

// CFRunLoopSourceSignal signals a CFRunLoopSource object, marking it as ready to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceSignal(_:)
func CFRunLoopSourceSignal(source CFRunLoopSourceRef) {
	if callErr := tryCFRunLoopSourceSignal(source); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopStop func(rl CFRunLoopRef)
var _cFRunLoopStopErr error

func tryCFRunLoopStop(rl CFRunLoopRef) error {
	if _cFRunLoopStop == nil {
		return symbolCallError("CFRunLoopStop", "", _cFRunLoopStopErr)
	}
	_cFRunLoopStop(rl)
	return nil
}

// CFRunLoopStop forces a CFRunLoop object to stop running.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopStop(_:)
func CFRunLoopStop(rl CFRunLoopRef) {
	if callErr := tryCFRunLoopStop(rl); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopTimerCreate func(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int, callout CFRunLoopTimerCallBack, context *CFRunLoopTimerContext) CFRunLoopTimerRef
var _cFRunLoopTimerCreateErr error

func tryCFRunLoopTimerCreate(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int, callout CFRunLoopTimerCallBack, context *CFRunLoopTimerContext) (CFRunLoopTimerRef, error) {
	if _cFRunLoopTimerCreate == nil {
		return 0, symbolCallError("CFRunLoopTimerCreate", "", _cFRunLoopTimerCreateErr)
	}
	return _cFRunLoopTimerCreate(allocator, fireDate, interval, flags, order, callout, context), nil
}

// CFRunLoopTimerCreate creates a new CFRunLoopTimer object with a function callback.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreate(_:_:_:_:_:_:_:)
func CFRunLoopTimerCreate(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int, callout CFRunLoopTimerCallBack, context *CFRunLoopTimerContext) CFRunLoopTimerRef {
	result, callErr := tryCFRunLoopTimerCreate(allocator, fireDate, interval, flags, order, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerCreateWithHandler func(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int) CFRunLoopTimerRef
var _cFRunLoopTimerCreateWithHandlerErr error

func tryCFRunLoopTimerCreateWithHandler(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int) (CFRunLoopTimerRef, error) {
	if _cFRunLoopTimerCreateWithHandler == nil {
		return 0, symbolCallError("CFRunLoopTimerCreateWithHandler", "10.7", _cFRunLoopTimerCreateWithHandlerErr)
	}
	return _cFRunLoopTimerCreateWithHandler(allocator, fireDate, interval, flags, order), nil
}

// CFRunLoopTimerCreateWithHandler creates a new CFRunLoopTimer object with a block-based handler.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreateWithHandler(_:_:_:_:_:_:)
func CFRunLoopTimerCreateWithHandler(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int) CFRunLoopTimerRef {
	result, callErr := tryCFRunLoopTimerCreateWithHandler(allocator, fireDate, interval, flags, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerDoesRepeat func(timer CFRunLoopTimerRef) bool
var _cFRunLoopTimerDoesRepeatErr error

func tryCFRunLoopTimerDoesRepeat(timer CFRunLoopTimerRef) (bool, error) {
	if _cFRunLoopTimerDoesRepeat == nil {
		return false, symbolCallError("CFRunLoopTimerDoesRepeat", "", _cFRunLoopTimerDoesRepeatErr)
	}
	return _cFRunLoopTimerDoesRepeat(timer), nil
}

// CFRunLoopTimerDoesRepeat returns a Boolean value that indicates whether a CFRunLoopTimer object repeats.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerDoesRepeat(_:)
func CFRunLoopTimerDoesRepeat(timer CFRunLoopTimerRef) bool {
	result, callErr := tryCFRunLoopTimerDoesRepeat(timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerGetContext func(timer CFRunLoopTimerRef, context *CFRunLoopTimerContext)
var _cFRunLoopTimerGetContextErr error

func tryCFRunLoopTimerGetContext(timer CFRunLoopTimerRef, context *CFRunLoopTimerContext) error {
	if _cFRunLoopTimerGetContext == nil {
		return symbolCallError("CFRunLoopTimerGetContext", "", _cFRunLoopTimerGetContextErr)
	}
	_cFRunLoopTimerGetContext(timer, context)
	return nil
}

// CFRunLoopTimerGetContext returns the context information for a CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetContext(_:_:)
func CFRunLoopTimerGetContext(timer CFRunLoopTimerRef, context *CFRunLoopTimerContext) {
	if callErr := tryCFRunLoopTimerGetContext(timer, context); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopTimerGetInterval func(timer CFRunLoopTimerRef) float64
var _cFRunLoopTimerGetIntervalErr error

func tryCFRunLoopTimerGetInterval(timer CFRunLoopTimerRef) (float64, error) {
	if _cFRunLoopTimerGetInterval == nil {
		return 0.0, symbolCallError("CFRunLoopTimerGetInterval", "", _cFRunLoopTimerGetIntervalErr)
	}
	return _cFRunLoopTimerGetInterval(timer), nil
}

// CFRunLoopTimerGetInterval returns the firing interval of a repeating CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetInterval(_:)
func CFRunLoopTimerGetInterval(timer CFRunLoopTimerRef) float64 {
	result, callErr := tryCFRunLoopTimerGetInterval(timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerGetNextFireDate func(timer CFRunLoopTimerRef) CFAbsoluteTime
var _cFRunLoopTimerGetNextFireDateErr error

func tryCFRunLoopTimerGetNextFireDate(timer CFRunLoopTimerRef) (CFAbsoluteTime, error) {
	if _cFRunLoopTimerGetNextFireDate == nil {
		return *new(CFAbsoluteTime), symbolCallError("CFRunLoopTimerGetNextFireDate", "", _cFRunLoopTimerGetNextFireDateErr)
	}
	return _cFRunLoopTimerGetNextFireDate(timer), nil
}

// CFRunLoopTimerGetNextFireDate returns the next firing time for a CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetNextFireDate(_:)
func CFRunLoopTimerGetNextFireDate(timer CFRunLoopTimerRef) CFAbsoluteTime {
	result, callErr := tryCFRunLoopTimerGetNextFireDate(timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerGetOrder func(timer CFRunLoopTimerRef) int
var _cFRunLoopTimerGetOrderErr error

func tryCFRunLoopTimerGetOrder(timer CFRunLoopTimerRef) (int, error) {
	if _cFRunLoopTimerGetOrder == nil {
		return 0, symbolCallError("CFRunLoopTimerGetOrder", "", _cFRunLoopTimerGetOrderErr)
	}
	return _cFRunLoopTimerGetOrder(timer), nil
}

// CFRunLoopTimerGetOrder returns the ordering parameter for a CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetOrder(_:)
func CFRunLoopTimerGetOrder(timer CFRunLoopTimerRef) int {
	result, callErr := tryCFRunLoopTimerGetOrder(timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerGetTolerance func(timer CFRunLoopTimerRef) float64
var _cFRunLoopTimerGetToleranceErr error

func tryCFRunLoopTimerGetTolerance(timer CFRunLoopTimerRef) (float64, error) {
	if _cFRunLoopTimerGetTolerance == nil {
		return 0.0, symbolCallError("CFRunLoopTimerGetTolerance", "10.9", _cFRunLoopTimerGetToleranceErr)
	}
	return _cFRunLoopTimerGetTolerance(timer), nil
}

// CFRunLoopTimerGetTolerance.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTolerance(_:)
func CFRunLoopTimerGetTolerance(timer CFRunLoopTimerRef) float64 {
	result, callErr := tryCFRunLoopTimerGetTolerance(timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerGetTypeID func() uint
var _cFRunLoopTimerGetTypeIDErr error

func tryCFRunLoopTimerGetTypeID() (uint, error) {
	if _cFRunLoopTimerGetTypeID == nil {
		return 0, symbolCallError("CFRunLoopTimerGetTypeID", "", _cFRunLoopTimerGetTypeIDErr)
	}
	return _cFRunLoopTimerGetTypeID(), nil
}

// CFRunLoopTimerGetTypeID returns the type identifier of the CFRunLoopTimer opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTypeID()
func CFRunLoopTimerGetTypeID() uint {
	result, callErr := tryCFRunLoopTimerGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerInvalidate func(timer CFRunLoopTimerRef)
var _cFRunLoopTimerInvalidateErr error

func tryCFRunLoopTimerInvalidate(timer CFRunLoopTimerRef) error {
	if _cFRunLoopTimerInvalidate == nil {
		return symbolCallError("CFRunLoopTimerInvalidate", "", _cFRunLoopTimerInvalidateErr)
	}
	_cFRunLoopTimerInvalidate(timer)
	return nil
}

// CFRunLoopTimerInvalidate invalidates a CFRunLoopTimer object, stopping it from ever firing again.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerInvalidate(_:)
func CFRunLoopTimerInvalidate(timer CFRunLoopTimerRef) {
	if callErr := tryCFRunLoopTimerInvalidate(timer); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopTimerIsValid func(timer CFRunLoopTimerRef) bool
var _cFRunLoopTimerIsValidErr error

func tryCFRunLoopTimerIsValid(timer CFRunLoopTimerRef) (bool, error) {
	if _cFRunLoopTimerIsValid == nil {
		return false, symbolCallError("CFRunLoopTimerIsValid", "", _cFRunLoopTimerIsValidErr)
	}
	return _cFRunLoopTimerIsValid(timer), nil
}

// CFRunLoopTimerIsValid returns a Boolean value that indicates whether a CFRunLoopTimer object is valid and able to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerIsValid(_:)
func CFRunLoopTimerIsValid(timer CFRunLoopTimerRef) bool {
	result, callErr := tryCFRunLoopTimerIsValid(timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFRunLoopTimerSetNextFireDate func(timer CFRunLoopTimerRef, fireDate CFAbsoluteTime)
var _cFRunLoopTimerSetNextFireDateErr error

func tryCFRunLoopTimerSetNextFireDate(timer CFRunLoopTimerRef, fireDate CFAbsoluteTime) error {
	if _cFRunLoopTimerSetNextFireDate == nil {
		return symbolCallError("CFRunLoopTimerSetNextFireDate", "", _cFRunLoopTimerSetNextFireDateErr)
	}
	_cFRunLoopTimerSetNextFireDate(timer, fireDate)
	return nil
}

// CFRunLoopTimerSetNextFireDate sets the next firing date for a CFRunLoopTimer object .
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetNextFireDate(_:_:)
func CFRunLoopTimerSetNextFireDate(timer CFRunLoopTimerRef, fireDate CFAbsoluteTime) {
	if callErr := tryCFRunLoopTimerSetNextFireDate(timer, fireDate); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopTimerSetTolerance func(timer CFRunLoopTimerRef, tolerance float64)
var _cFRunLoopTimerSetToleranceErr error

func tryCFRunLoopTimerSetTolerance(timer CFRunLoopTimerRef, tolerance float64) error {
	if _cFRunLoopTimerSetTolerance == nil {
		return symbolCallError("CFRunLoopTimerSetTolerance", "10.9", _cFRunLoopTimerSetToleranceErr)
	}
	_cFRunLoopTimerSetTolerance(timer, tolerance)
	return nil
}

// CFRunLoopTimerSetTolerance.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetTolerance(_:_:)
func CFRunLoopTimerSetTolerance(timer CFRunLoopTimerRef, tolerance float64) {
	if callErr := tryCFRunLoopTimerSetTolerance(timer, tolerance); callErr != nil {
		panic(callErr)
	}
}

var _cFRunLoopWakeUp func(rl CFRunLoopRef)
var _cFRunLoopWakeUpErr error

func tryCFRunLoopWakeUp(rl CFRunLoopRef) error {
	if _cFRunLoopWakeUp == nil {
		return symbolCallError("CFRunLoopWakeUp", "", _cFRunLoopWakeUpErr)
	}
	_cFRunLoopWakeUp(rl)
	return nil
}

// CFRunLoopWakeUp wakes a waiting CFRunLoop object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopWakeUp(_:)
func CFRunLoopWakeUp(rl CFRunLoopRef) {
	if callErr := tryCFRunLoopWakeUp(rl); callErr != nil {
		panic(callErr)
	}
}

var _cFSetAddValue func(theSet CFMutableSetRef, value unsafe.Pointer)
var _cFSetAddValueErr error

func tryCFSetAddValue(theSet CFMutableSetRef, value unsafe.Pointer) error {
	if _cFSetAddValue == nil {
		return symbolCallError("CFSetAddValue", "", _cFSetAddValueErr)
	}
	_cFSetAddValue(theSet, value)
	return nil
}

// CFSetAddValue adds a value to a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetAddValue(_:_:)
func CFSetAddValue(theSet CFMutableSetRef, value unsafe.Pointer) {
	if callErr := tryCFSetAddValue(theSet, value); callErr != nil {
		panic(callErr)
	}
}

var _cFSetApplyFunction func(theSet CFSetRef, applier CFSetApplierFunction, context unsafe.Pointer)
var _cFSetApplyFunctionErr error

func tryCFSetApplyFunction(theSet CFSetRef, applier CFSetApplierFunction, context unsafe.Pointer) error {
	if _cFSetApplyFunction == nil {
		return symbolCallError("CFSetApplyFunction", "", _cFSetApplyFunctionErr)
	}
	_cFSetApplyFunction(theSet, applier, context)
	return nil
}

// CFSetApplyFunction calls a function once for each value in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetApplyFunction(_:_:_:)
func CFSetApplyFunction(theSet CFSetRef, applier CFSetApplierFunction, context unsafe.Pointer) {
	if callErr := tryCFSetApplyFunction(theSet, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cFSetContainsValue func(theSet CFSetRef, value unsafe.Pointer) bool
var _cFSetContainsValueErr error

func tryCFSetContainsValue(theSet CFSetRef, value unsafe.Pointer) (bool, error) {
	if _cFSetContainsValue == nil {
		return false, symbolCallError("CFSetContainsValue", "", _cFSetContainsValueErr)
	}
	return _cFSetContainsValue(theSet, value), nil
}

// CFSetContainsValue returns a Boolean that indicates whether a set contains a given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetContainsValue(_:_:)
func CFSetContainsValue(theSet CFSetRef, value unsafe.Pointer) bool {
	result, callErr := tryCFSetContainsValue(theSet, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetCreate func(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFSetCallBacks) CFSetRef
var _cFSetCreateErr error

func tryCFSetCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFSetCallBacks) (CFSetRef, error) {
	if _cFSetCreate == nil {
		return 0, symbolCallError("CFSetCreate", "", _cFSetCreateErr)
	}
	return _cFSetCreate(allocator, values, numValues, callBacks), nil
}

// CFSetCreate creates an immutable CFSet object containing supplied values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreate(_:_:_:_:)
func CFSetCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFSetCallBacks) CFSetRef {
	result, callErr := tryCFSetCreate(allocator, values, numValues, callBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetCreateCopy func(allocator CFAllocatorRef, theSet CFSetRef) CFSetRef
var _cFSetCreateCopyErr error

func tryCFSetCreateCopy(allocator CFAllocatorRef, theSet CFSetRef) (CFSetRef, error) {
	if _cFSetCreateCopy == nil {
		return 0, symbolCallError("CFSetCreateCopy", "", _cFSetCreateCopyErr)
	}
	return _cFSetCreateCopy(allocator, theSet), nil
}

// CFSetCreateCopy creates an immutable set containing the values of an existing set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateCopy(_:_:)
func CFSetCreateCopy(allocator CFAllocatorRef, theSet CFSetRef) CFSetRef {
	result, callErr := tryCFSetCreateCopy(allocator, theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetCreateMutable func(allocator CFAllocatorRef, capacity int, callBacks *CFSetCallBacks) CFMutableSetRef
var _cFSetCreateMutableErr error

func tryCFSetCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFSetCallBacks) (CFMutableSetRef, error) {
	if _cFSetCreateMutable == nil {
		return 0, symbolCallError("CFSetCreateMutable", "", _cFSetCreateMutableErr)
	}
	return _cFSetCreateMutable(allocator, capacity, callBacks), nil
}

// CFSetCreateMutable creates an empty CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutable(_:_:_:)
func CFSetCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFSetCallBacks) CFMutableSetRef {
	result, callErr := tryCFSetCreateMutable(allocator, capacity, callBacks)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theSet CFSetRef) CFMutableSetRef
var _cFSetCreateMutableCopyErr error

func tryCFSetCreateMutableCopy(allocator CFAllocatorRef, capacity int, theSet CFSetRef) (CFMutableSetRef, error) {
	if _cFSetCreateMutableCopy == nil {
		return 0, symbolCallError("CFSetCreateMutableCopy", "", _cFSetCreateMutableCopyErr)
	}
	return _cFSetCreateMutableCopy(allocator, capacity, theSet), nil
}

// CFSetCreateMutableCopy creates a new mutable set with the values from another set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutableCopy(_:_:_:)
func CFSetCreateMutableCopy(allocator CFAllocatorRef, capacity int, theSet CFSetRef) CFMutableSetRef {
	result, callErr := tryCFSetCreateMutableCopy(allocator, capacity, theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetGetCount func(theSet CFSetRef) int
var _cFSetGetCountErr error

func tryCFSetGetCount(theSet CFSetRef) (int, error) {
	if _cFSetGetCount == nil {
		return 0, symbolCallError("CFSetGetCount", "", _cFSetGetCountErr)
	}
	return _cFSetGetCount(theSet), nil
}

// CFSetGetCount returns the number of values currently in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCount(_:)
func CFSetGetCount(theSet CFSetRef) int {
	result, callErr := tryCFSetGetCount(theSet)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetGetCountOfValue func(theSet CFSetRef, value unsafe.Pointer) int
var _cFSetGetCountOfValueErr error

func tryCFSetGetCountOfValue(theSet CFSetRef, value unsafe.Pointer) (int, error) {
	if _cFSetGetCountOfValue == nil {
		return 0, symbolCallError("CFSetGetCountOfValue", "", _cFSetGetCountOfValueErr)
	}
	return _cFSetGetCountOfValue(theSet, value), nil
}

// CFSetGetCountOfValue returns the number of values in a set that match a given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCountOfValue(_:_:)
func CFSetGetCountOfValue(theSet CFSetRef, value unsafe.Pointer) int {
	result, callErr := tryCFSetGetCountOfValue(theSet, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetGetTypeID func() uint
var _cFSetGetTypeIDErr error

func tryCFSetGetTypeID() (uint, error) {
	if _cFSetGetTypeID == nil {
		return 0, symbolCallError("CFSetGetTypeID", "", _cFSetGetTypeIDErr)
	}
	return _cFSetGetTypeID(), nil
}

// CFSetGetTypeID returns the type identifier for the CFSet type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetTypeID()
func CFSetGetTypeID() uint {
	result, callErr := tryCFSetGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetGetValue func(theSet CFSetRef, value unsafe.Pointer) unsafe.Pointer
var _cFSetGetValueErr error

func tryCFSetGetValue(theSet CFSetRef, value unsafe.Pointer) (unsafe.Pointer, error) {
	if _cFSetGetValue == nil {
		return nil, symbolCallError("CFSetGetValue", "", _cFSetGetValueErr)
	}
	return _cFSetGetValue(theSet, value), nil
}

// CFSetGetValue obtains a specified value from a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValue(_:_:)
func CFSetGetValue(theSet CFSetRef, value unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryCFSetGetValue(theSet, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetGetValueIfPresent func(theSet CFSetRef, candidate unsafe.Pointer, value unsafe.Pointer) bool
var _cFSetGetValueIfPresentErr error

func tryCFSetGetValueIfPresent(theSet CFSetRef, candidate unsafe.Pointer, value unsafe.Pointer) (bool, error) {
	if _cFSetGetValueIfPresent == nil {
		return false, symbolCallError("CFSetGetValueIfPresent", "", _cFSetGetValueIfPresentErr)
	}
	return _cFSetGetValueIfPresent(theSet, candidate, value), nil
}

// CFSetGetValueIfPresent reports whether or not a value is in a set, and if it exists returns the value indirectly.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValueIfPresent(_:_:_:)
func CFSetGetValueIfPresent(theSet CFSetRef, candidate unsafe.Pointer, value unsafe.Pointer) bool {
	result, callErr := tryCFSetGetValueIfPresent(theSet, candidate, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSetGetValues func(theSet CFSetRef, values unsafe.Pointer)
var _cFSetGetValuesErr error

func tryCFSetGetValues(theSet CFSetRef, values unsafe.Pointer) error {
	if _cFSetGetValues == nil {
		return symbolCallError("CFSetGetValues", "", _cFSetGetValuesErr)
	}
	_cFSetGetValues(theSet, values)
	return nil
}

// CFSetGetValues obtains all values in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValues(_:_:)
func CFSetGetValues(theSet CFSetRef, values unsafe.Pointer) {
	if callErr := tryCFSetGetValues(theSet, values); callErr != nil {
		panic(callErr)
	}
}

var _cFSetRemoveAllValues func(theSet CFMutableSetRef)
var _cFSetRemoveAllValuesErr error

func tryCFSetRemoveAllValues(theSet CFMutableSetRef) error {
	if _cFSetRemoveAllValues == nil {
		return symbolCallError("CFSetRemoveAllValues", "", _cFSetRemoveAllValuesErr)
	}
	_cFSetRemoveAllValues(theSet)
	return nil
}

// CFSetRemoveAllValues removes all values from a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveAllValues(_:)
func CFSetRemoveAllValues(theSet CFMutableSetRef) {
	if callErr := tryCFSetRemoveAllValues(theSet); callErr != nil {
		panic(callErr)
	}
}

var _cFSetRemoveValue func(theSet CFMutableSetRef, value unsafe.Pointer)
var _cFSetRemoveValueErr error

func tryCFSetRemoveValue(theSet CFMutableSetRef, value unsafe.Pointer) error {
	if _cFSetRemoveValue == nil {
		return symbolCallError("CFSetRemoveValue", "", _cFSetRemoveValueErr)
	}
	_cFSetRemoveValue(theSet, value)
	return nil
}

// CFSetRemoveValue removes a value from a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveValue(_:_:)
func CFSetRemoveValue(theSet CFMutableSetRef, value unsafe.Pointer) {
	if callErr := tryCFSetRemoveValue(theSet, value); callErr != nil {
		panic(callErr)
	}
}

var _cFSetReplaceValue func(theSet CFMutableSetRef, value unsafe.Pointer)
var _cFSetReplaceValueErr error

func tryCFSetReplaceValue(theSet CFMutableSetRef, value unsafe.Pointer) error {
	if _cFSetReplaceValue == nil {
		return symbolCallError("CFSetReplaceValue", "", _cFSetReplaceValueErr)
	}
	_cFSetReplaceValue(theSet, value)
	return nil
}

// CFSetReplaceValue replaces a value in a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetReplaceValue(_:_:)
func CFSetReplaceValue(theSet CFMutableSetRef, value unsafe.Pointer) {
	if callErr := tryCFSetReplaceValue(theSet, value); callErr != nil {
		panic(callErr)
	}
}

var _cFSetSetValue func(theSet CFMutableSetRef, value unsafe.Pointer)
var _cFSetSetValueErr error

func tryCFSetSetValue(theSet CFMutableSetRef, value unsafe.Pointer) error {
	if _cFSetSetValue == nil {
		return symbolCallError("CFSetSetValue", "", _cFSetSetValueErr)
	}
	_cFSetSetValue(theSet, value)
	return nil
}

// CFSetSetValue sets a value in a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetSetValue(_:_:)
func CFSetSetValue(theSet CFMutableSetRef, value unsafe.Pointer) {
	if callErr := tryCFSetSetValue(theSet, value); callErr != nil {
		panic(callErr)
	}
}

var _cFShow func(obj CFTypeRef)
var _cFShowErr error

func tryCFShow(obj CFTypeRef) error {
	if _cFShow == nil {
		return symbolCallError("CFShow", "", _cFShowErr)
	}
	_cFShow(obj)
	return nil
}

// CFShow prints a description of a Core Foundation object to stderr.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFShow(_:)
func CFShow(obj CFTypeRef) {
	if callErr := tryCFShow(obj); callErr != nil {
		panic(callErr)
	}
}

var _cFShowStr func(str CFStringRef)
var _cFShowStrErr error

func tryCFShowStr(str CFStringRef) error {
	if _cFShowStr == nil {
		return symbolCallError("CFShowStr", "", _cFShowStrErr)
	}
	_cFShowStr(str)
	return nil
}

// CFShowStr prints the attributes of a string during debugging.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFShowStr(_:)
func CFShowStr(str CFStringRef) {
	if callErr := tryCFShowStr(str); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketConnectToAddress func(s CFSocketRef, address CFDataRef, timeout float64) CFSocketError
var _cFSocketConnectToAddressErr error

func tryCFSocketConnectToAddress(s CFSocketRef, address CFDataRef, timeout float64) (CFSocketError, error) {
	if _cFSocketConnectToAddress == nil {
		return *new(CFSocketError), symbolCallError("CFSocketConnectToAddress", "", _cFSocketConnectToAddressErr)
	}
	return _cFSocketConnectToAddress(s, address, timeout), nil
}

// CFSocketConnectToAddress opens a connection to a remote socket.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketConnectToAddress(_:_:_:)
func CFSocketConnectToAddress(s CFSocketRef, address CFDataRef, timeout float64) CFSocketError {
	result, callErr := tryCFSocketConnectToAddress(s, address, timeout)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCopyAddress func(s CFSocketRef) CFDataRef
var _cFSocketCopyAddressErr error

func tryCFSocketCopyAddress(s CFSocketRef) (CFDataRef, error) {
	if _cFSocketCopyAddress == nil {
		return 0, symbolCallError("CFSocketCopyAddress", "", _cFSocketCopyAddressErr)
	}
	return _cFSocketCopyAddress(s), nil
}

// CFSocketCopyAddress returns the local address of a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyAddress(_:)
func CFSocketCopyAddress(s CFSocketRef) CFDataRef {
	result, callErr := tryCFSocketCopyAddress(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCopyPeerAddress func(s CFSocketRef) CFDataRef
var _cFSocketCopyPeerAddressErr error

func tryCFSocketCopyPeerAddress(s CFSocketRef) (CFDataRef, error) {
	if _cFSocketCopyPeerAddress == nil {
		return 0, symbolCallError("CFSocketCopyPeerAddress", "", _cFSocketCopyPeerAddressErr)
	}
	return _cFSocketCopyPeerAddress(s), nil
}

// CFSocketCopyPeerAddress returns the remote address to which a CFSocket object is connected.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyPeerAddress(_:)
func CFSocketCopyPeerAddress(s CFSocketRef) CFDataRef {
	result, callErr := tryCFSocketCopyPeerAddress(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCopyRegisteredSocketSignature func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature, nameServerAddress *CFDataRef) CFSocketError
var _cFSocketCopyRegisteredSocketSignatureErr error

func tryCFSocketCopyRegisteredSocketSignature(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature, nameServerAddress *CFDataRef) (CFSocketError, error) {
	if _cFSocketCopyRegisteredSocketSignature == nil {
		return *new(CFSocketError), symbolCallError("CFSocketCopyRegisteredSocketSignature", "", _cFSocketCopyRegisteredSocketSignatureErr)
	}
	return _cFSocketCopyRegisteredSocketSignature(nameServerSignature, timeout, name, signature, nameServerAddress), nil
}

// CFSocketCopyRegisteredSocketSignature returns a socket signature registered with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredSocketSignature(_:_:_:_:_:)
func CFSocketCopyRegisteredSocketSignature(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature, nameServerAddress *CFDataRef) CFSocketError {
	result, callErr := tryCFSocketCopyRegisteredSocketSignature(nameServerSignature, timeout, name, signature, nameServerAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCopyRegisteredValue func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value *CFPropertyListRef, nameServerAddress *CFDataRef) CFSocketError
var _cFSocketCopyRegisteredValueErr error

func tryCFSocketCopyRegisteredValue(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value *CFPropertyListRef, nameServerAddress *CFDataRef) (CFSocketError, error) {
	if _cFSocketCopyRegisteredValue == nil {
		return *new(CFSocketError), symbolCallError("CFSocketCopyRegisteredValue", "", _cFSocketCopyRegisteredValueErr)
	}
	return _cFSocketCopyRegisteredValue(nameServerSignature, timeout, name, value, nameServerAddress), nil
}

// CFSocketCopyRegisteredValue returns a value registered with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredValue(_:_:_:_:_:)
func CFSocketCopyRegisteredValue(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value *CFPropertyListRef, nameServerAddress *CFDataRef) CFSocketError {
	result, callErr := tryCFSocketCopyRegisteredValue(nameServerSignature, timeout, name, value, nameServerAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCreate func(allocator CFAllocatorRef, protocolFamily int32, socketType int32, protocol_ int32, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef
var _cFSocketCreateErr error

func tryCFSocketCreate(allocator CFAllocatorRef, protocolFamily int32, socketType int32, protocol_ int32, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) (CFSocketRef, error) {
	if _cFSocketCreate == nil {
		return 0, symbolCallError("CFSocketCreate", "", _cFSocketCreateErr)
	}
	return _cFSocketCreate(allocator, protocolFamily, socketType, protocol_, callBackTypes, callout, context), nil
}

// CFSocketCreate creates a CFSocket object of a specified protocol and type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreate(_:_:_:_:_:_:_:)
func CFSocketCreate(allocator CFAllocatorRef, protocolFamily int32, socketType int32, protocol_ int32, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef {
	result, callErr := tryCFSocketCreate(allocator, protocolFamily, socketType, protocol_, callBackTypes, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCreateConnectedToSocketSignature func(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext, timeout float64) CFSocketRef
var _cFSocketCreateConnectedToSocketSignatureErr error

func tryCFSocketCreateConnectedToSocketSignature(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext, timeout float64) (CFSocketRef, error) {
	if _cFSocketCreateConnectedToSocketSignature == nil {
		return 0, symbolCallError("CFSocketCreateConnectedToSocketSignature", "", _cFSocketCreateConnectedToSocketSignatureErr)
	}
	return _cFSocketCreateConnectedToSocketSignature(allocator, signature, callBackTypes, callout, context, timeout), nil
}

// CFSocketCreateConnectedToSocketSignature creates a CFSocket object and opens a connection to a remote socket.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateConnectedToSocketSignature(_:_:_:_:_:_:)
func CFSocketCreateConnectedToSocketSignature(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext, timeout float64) CFSocketRef {
	result, callErr := tryCFSocketCreateConnectedToSocketSignature(allocator, signature, callBackTypes, callout, context, timeout)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCreateRunLoopSource func(allocator CFAllocatorRef, s CFSocketRef, order int) CFRunLoopSourceRef
var _cFSocketCreateRunLoopSourceErr error

func tryCFSocketCreateRunLoopSource(allocator CFAllocatorRef, s CFSocketRef, order int) (CFRunLoopSourceRef, error) {
	if _cFSocketCreateRunLoopSource == nil {
		return 0, symbolCallError("CFSocketCreateRunLoopSource", "", _cFSocketCreateRunLoopSourceErr)
	}
	return _cFSocketCreateRunLoopSource(allocator, s, order), nil
}

// CFSocketCreateRunLoopSource creates a CFRunLoopSource object for a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateRunLoopSource(_:_:_:)
func CFSocketCreateRunLoopSource(allocator CFAllocatorRef, s CFSocketRef, order int) CFRunLoopSourceRef {
	result, callErr := tryCFSocketCreateRunLoopSource(allocator, s, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCreateWithNative func(allocator CFAllocatorRef, sock CFSocketNativeHandle, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef
var _cFSocketCreateWithNativeErr error

func tryCFSocketCreateWithNative(allocator CFAllocatorRef, sock CFSocketNativeHandle, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) (CFSocketRef, error) {
	if _cFSocketCreateWithNative == nil {
		return 0, symbolCallError("CFSocketCreateWithNative", "", _cFSocketCreateWithNativeErr)
	}
	return _cFSocketCreateWithNative(allocator, sock, callBackTypes, callout, context), nil
}

// CFSocketCreateWithNative creates a CFSocket object for a pre-existing native socket.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithNative(_:_:_:_:_:)
func CFSocketCreateWithNative(allocator CFAllocatorRef, sock CFSocketNativeHandle, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef {
	result, callErr := tryCFSocketCreateWithNative(allocator, sock, callBackTypes, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketCreateWithSocketSignature func(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef
var _cFSocketCreateWithSocketSignatureErr error

func tryCFSocketCreateWithSocketSignature(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) (CFSocketRef, error) {
	if _cFSocketCreateWithSocketSignature == nil {
		return 0, symbolCallError("CFSocketCreateWithSocketSignature", "", _cFSocketCreateWithSocketSignatureErr)
	}
	return _cFSocketCreateWithSocketSignature(allocator, signature, callBackTypes, callout, context), nil
}

// CFSocketCreateWithSocketSignature creates a CFSocket object using information from a CFSocketSignature structure.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithSocketSignature(_:_:_:_:_:)
func CFSocketCreateWithSocketSignature(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef {
	result, callErr := tryCFSocketCreateWithSocketSignature(allocator, signature, callBackTypes, callout, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketDisableCallBacks func(s CFSocketRef, callBackTypes uint64)
var _cFSocketDisableCallBacksErr error

func tryCFSocketDisableCallBacks(s CFSocketRef, callBackTypes uint64) error {
	if _cFSocketDisableCallBacks == nil {
		return symbolCallError("CFSocketDisableCallBacks", "", _cFSocketDisableCallBacksErr)
	}
	_cFSocketDisableCallBacks(s, callBackTypes)
	return nil
}

// CFSocketDisableCallBacks disables the callback function of a CFSocket object for certain types of socket activity.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketDisableCallBacks(_:_:)
func CFSocketDisableCallBacks(s CFSocketRef, callBackTypes uint64) {
	if callErr := tryCFSocketDisableCallBacks(s, callBackTypes); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketEnableCallBacks func(s CFSocketRef, callBackTypes uint64)
var _cFSocketEnableCallBacksErr error

func tryCFSocketEnableCallBacks(s CFSocketRef, callBackTypes uint64) error {
	if _cFSocketEnableCallBacks == nil {
		return symbolCallError("CFSocketEnableCallBacks", "", _cFSocketEnableCallBacksErr)
	}
	_cFSocketEnableCallBacks(s, callBackTypes)
	return nil
}

// CFSocketEnableCallBacks enables the callback function of a CFSocket object for certain types of socket activity.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketEnableCallBacks(_:_:)
func CFSocketEnableCallBacks(s CFSocketRef, callBackTypes uint64) {
	if callErr := tryCFSocketEnableCallBacks(s, callBackTypes); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketGetContext func(s CFSocketRef, context *CFSocketContext)
var _cFSocketGetContextErr error

func tryCFSocketGetContext(s CFSocketRef, context *CFSocketContext) error {
	if _cFSocketGetContext == nil {
		return symbolCallError("CFSocketGetContext", "", _cFSocketGetContextErr)
	}
	_cFSocketGetContext(s, context)
	return nil
}

// CFSocketGetContext returns the context information for a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetContext(_:_:)
func CFSocketGetContext(s CFSocketRef, context *CFSocketContext) {
	if callErr := tryCFSocketGetContext(s, context); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketGetDefaultNameRegistryPortNumber func() uint16
var _cFSocketGetDefaultNameRegistryPortNumberErr error

func tryCFSocketGetDefaultNameRegistryPortNumber() (uint16, error) {
	if _cFSocketGetDefaultNameRegistryPortNumber == nil {
		return 0, symbolCallError("CFSocketGetDefaultNameRegistryPortNumber", "", _cFSocketGetDefaultNameRegistryPortNumberErr)
	}
	return _cFSocketGetDefaultNameRegistryPortNumber(), nil
}

// CFSocketGetDefaultNameRegistryPortNumber returns the default port number with which to connect to a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetDefaultNameRegistryPortNumber()
func CFSocketGetDefaultNameRegistryPortNumber() uint16 {
	result, callErr := tryCFSocketGetDefaultNameRegistryPortNumber()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketGetNative func(s CFSocketRef) CFSocketNativeHandle
var _cFSocketGetNativeErr error

func tryCFSocketGetNative(s CFSocketRef) (CFSocketNativeHandle, error) {
	if _cFSocketGetNative == nil {
		return *new(CFSocketNativeHandle), symbolCallError("CFSocketGetNative", "", _cFSocketGetNativeErr)
	}
	return _cFSocketGetNative(s), nil
}

// CFSocketGetNative returns the native socket associated with a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetNative(_:)
func CFSocketGetNative(s CFSocketRef) CFSocketNativeHandle {
	result, callErr := tryCFSocketGetNative(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketGetSocketFlags func(s CFSocketRef) uint64
var _cFSocketGetSocketFlagsErr error

func tryCFSocketGetSocketFlags(s CFSocketRef) (uint64, error) {
	if _cFSocketGetSocketFlags == nil {
		return 0, symbolCallError("CFSocketGetSocketFlags", "", _cFSocketGetSocketFlagsErr)
	}
	return _cFSocketGetSocketFlags(s), nil
}

// CFSocketGetSocketFlags returns flags that control certain behaviors of a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetSocketFlags(_:)
func CFSocketGetSocketFlags(s CFSocketRef) uint64 {
	result, callErr := tryCFSocketGetSocketFlags(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketGetTypeID func() uint
var _cFSocketGetTypeIDErr error

func tryCFSocketGetTypeID() (uint, error) {
	if _cFSocketGetTypeID == nil {
		return 0, symbolCallError("CFSocketGetTypeID", "", _cFSocketGetTypeIDErr)
	}
	return _cFSocketGetTypeID(), nil
}

// CFSocketGetTypeID returns the type identifier for the CFSocket opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetTypeID()
func CFSocketGetTypeID() uint {
	result, callErr := tryCFSocketGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketInvalidate func(s CFSocketRef)
var _cFSocketInvalidateErr error

func tryCFSocketInvalidate(s CFSocketRef) error {
	if _cFSocketInvalidate == nil {
		return symbolCallError("CFSocketInvalidate", "", _cFSocketInvalidateErr)
	}
	_cFSocketInvalidate(s)
	return nil
}

// CFSocketInvalidate invalidates a CFSocket object, stopping it from sending or receiving any more messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketInvalidate(_:)
func CFSocketInvalidate(s CFSocketRef) {
	if callErr := tryCFSocketInvalidate(s); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketIsValid func(s CFSocketRef) bool
var _cFSocketIsValidErr error

func tryCFSocketIsValid(s CFSocketRef) (bool, error) {
	if _cFSocketIsValid == nil {
		return false, symbolCallError("CFSocketIsValid", "", _cFSocketIsValidErr)
	}
	return _cFSocketIsValid(s), nil
}

// CFSocketIsValid returns a Boolean value that indicates whether a CFSocket object is valid and able to send or receive messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketIsValid(_:)
func CFSocketIsValid(s CFSocketRef) bool {
	result, callErr := tryCFSocketIsValid(s)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketRegisterSocketSignature func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature) CFSocketError
var _cFSocketRegisterSocketSignatureErr error

func tryCFSocketRegisterSocketSignature(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature) (CFSocketError, error) {
	if _cFSocketRegisterSocketSignature == nil {
		return *new(CFSocketError), symbolCallError("CFSocketRegisterSocketSignature", "", _cFSocketRegisterSocketSignatureErr)
	}
	return _cFSocketRegisterSocketSignature(nameServerSignature, timeout, name, signature), nil
}

// CFSocketRegisterSocketSignature registers a socket signature with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterSocketSignature(_:_:_:_:)
func CFSocketRegisterSocketSignature(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature) CFSocketError {
	result, callErr := tryCFSocketRegisterSocketSignature(nameServerSignature, timeout, name, signature)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketRegisterValue func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value CFPropertyListRef) CFSocketError
var _cFSocketRegisterValueErr error

func tryCFSocketRegisterValue(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value CFPropertyListRef) (CFSocketError, error) {
	if _cFSocketRegisterValue == nil {
		return *new(CFSocketError), symbolCallError("CFSocketRegisterValue", "", _cFSocketRegisterValueErr)
	}
	return _cFSocketRegisterValue(nameServerSignature, timeout, name, value), nil
}

// CFSocketRegisterValue registers a property-list value with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterValue(_:_:_:_:)
func CFSocketRegisterValue(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value CFPropertyListRef) CFSocketError {
	result, callErr := tryCFSocketRegisterValue(nameServerSignature, timeout, name, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketSendData func(s CFSocketRef, address CFDataRef, data CFDataRef, timeout float64) CFSocketError
var _cFSocketSendDataErr error

func tryCFSocketSendData(s CFSocketRef, address CFDataRef, data CFDataRef, timeout float64) (CFSocketError, error) {
	if _cFSocketSendData == nil {
		return *new(CFSocketError), symbolCallError("CFSocketSendData", "", _cFSocketSendDataErr)
	}
	return _cFSocketSendData(s, address, data, timeout), nil
}

// CFSocketSendData sends data over a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSendData(_:_:_:_:)
func CFSocketSendData(s CFSocketRef, address CFDataRef, data CFDataRef, timeout float64) CFSocketError {
	result, callErr := tryCFSocketSendData(s, address, data, timeout)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketSetAddress func(s CFSocketRef, address CFDataRef) CFSocketError
var _cFSocketSetAddressErr error

func tryCFSocketSetAddress(s CFSocketRef, address CFDataRef) (CFSocketError, error) {
	if _cFSocketSetAddress == nil {
		return *new(CFSocketError), symbolCallError("CFSocketSetAddress", "", _cFSocketSetAddressErr)
	}
	return _cFSocketSetAddress(s, address), nil
}

// CFSocketSetAddress binds a local address to a CFSocket object and configures it for listening.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetAddress(_:_:)
func CFSocketSetAddress(s CFSocketRef, address CFDataRef) CFSocketError {
	result, callErr := tryCFSocketSetAddress(s, address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFSocketSetDefaultNameRegistryPortNumber func(port uint16)
var _cFSocketSetDefaultNameRegistryPortNumberErr error

func tryCFSocketSetDefaultNameRegistryPortNumber(port uint16) error {
	if _cFSocketSetDefaultNameRegistryPortNumber == nil {
		return symbolCallError("CFSocketSetDefaultNameRegistryPortNumber", "", _cFSocketSetDefaultNameRegistryPortNumberErr)
	}
	_cFSocketSetDefaultNameRegistryPortNumber(port)
	return nil
}

// CFSocketSetDefaultNameRegistryPortNumber sets the default port number with which to connect to a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetDefaultNameRegistryPortNumber(_:)
func CFSocketSetDefaultNameRegistryPortNumber(port uint16) {
	if callErr := tryCFSocketSetDefaultNameRegistryPortNumber(port); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketSetSocketFlags func(s CFSocketRef, flags uint64)
var _cFSocketSetSocketFlagsErr error

func tryCFSocketSetSocketFlags(s CFSocketRef, flags uint64) error {
	if _cFSocketSetSocketFlags == nil {
		return symbolCallError("CFSocketSetSocketFlags", "", _cFSocketSetSocketFlagsErr)
	}
	_cFSocketSetSocketFlags(s, flags)
	return nil
}

// CFSocketSetSocketFlags sets flags that control certain behaviors of a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetSocketFlags(_:_:)
func CFSocketSetSocketFlags(s CFSocketRef, flags uint64) {
	if callErr := tryCFSocketSetSocketFlags(s, flags); callErr != nil {
		panic(callErr)
	}
}

var _cFSocketUnregister func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef) CFSocketError
var _cFSocketUnregisterErr error

func tryCFSocketUnregister(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef) (CFSocketError, error) {
	if _cFSocketUnregister == nil {
		return *new(CFSocketError), symbolCallError("CFSocketUnregister", "", _cFSocketUnregisterErr)
	}
	return _cFSocketUnregister(nameServerSignature, timeout, name), nil
}

// CFSocketUnregister unregisters a value or socket signature with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketUnregister(_:_:_:)
func CFSocketUnregister(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef) CFSocketError {
	result, callErr := tryCFSocketUnregister(nameServerSignature, timeout, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStreamCreateBoundPair func(alloc CFAllocatorRef, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef, transferBufferSize int)
var _cFStreamCreateBoundPairErr error

func tryCFStreamCreateBoundPair(alloc CFAllocatorRef, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef, transferBufferSize int) error {
	if _cFStreamCreateBoundPair == nil {
		return symbolCallError("CFStreamCreateBoundPair", "", _cFStreamCreateBoundPairErr)
	}
	_cFStreamCreateBoundPair(alloc, readStream, writeStream, transferBufferSize)
	return nil
}

// CFStreamCreateBoundPair creates a bound pair of read and write streams.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreateBoundPair(_:_:_:_:)
func CFStreamCreateBoundPair(alloc CFAllocatorRef, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef, transferBufferSize int) {
	if callErr := tryCFStreamCreateBoundPair(alloc, readStream, writeStream, transferBufferSize); callErr != nil {
		panic(callErr)
	}
}

var _cFStreamCreatePairWithPeerSocketSignature func(alloc CFAllocatorRef, signature *CFSocketSignature, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef)
var _cFStreamCreatePairWithPeerSocketSignatureErr error

func tryCFStreamCreatePairWithPeerSocketSignature(alloc CFAllocatorRef, signature *CFSocketSignature, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) error {
	if _cFStreamCreatePairWithPeerSocketSignature == nil {
		return symbolCallError("CFStreamCreatePairWithPeerSocketSignature", "10.1", _cFStreamCreatePairWithPeerSocketSignatureErr)
	}
	_cFStreamCreatePairWithPeerSocketSignature(alloc, signature, readStream, writeStream)
	return nil
}

// CFStreamCreatePairWithPeerSocketSignature creates readable and writable streams connected to a socket.
//
// Deprecated: Deprecated since macOS 26.4. Use nw_connection_t in Network framework instead
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithPeerSocketSignature(_:_:_:_:)
func CFStreamCreatePairWithPeerSocketSignature(alloc CFAllocatorRef, signature *CFSocketSignature, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) {
	if callErr := tryCFStreamCreatePairWithPeerSocketSignature(alloc, signature, readStream, writeStream); callErr != nil {
		panic(callErr)
	}
}

var _cFStreamCreatePairWithSocket func(alloc CFAllocatorRef, sock CFSocketNativeHandle, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef)
var _cFStreamCreatePairWithSocketErr error

func tryCFStreamCreatePairWithSocket(alloc CFAllocatorRef, sock CFSocketNativeHandle, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) error {
	if _cFStreamCreatePairWithSocket == nil {
		return symbolCallError("CFStreamCreatePairWithSocket", "10.1", _cFStreamCreatePairWithSocketErr)
	}
	_cFStreamCreatePairWithSocket(alloc, sock, readStream, writeStream)
	return nil
}

// CFStreamCreatePairWithSocket creates readable and writable streams connected to a socket.
//
// Deprecated: Deprecated since macOS 26.4. Use nw_connection_t in Network framework instead
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocket(_:_:_:_:)
func CFStreamCreatePairWithSocket(alloc CFAllocatorRef, sock CFSocketNativeHandle, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) {
	if callErr := tryCFStreamCreatePairWithSocket(alloc, sock, readStream, writeStream); callErr != nil {
		panic(callErr)
	}
}

var _cFStreamCreatePairWithSocketToHost func(alloc CFAllocatorRef, host CFStringRef, port uint32, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef)
var _cFStreamCreatePairWithSocketToHostErr error

func tryCFStreamCreatePairWithSocketToHost(alloc CFAllocatorRef, host CFStringRef, port uint32, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) error {
	if _cFStreamCreatePairWithSocketToHost == nil {
		return symbolCallError("CFStreamCreatePairWithSocketToHost", "10.1", _cFStreamCreatePairWithSocketToHostErr)
	}
	_cFStreamCreatePairWithSocketToHost(alloc, host, port, readStream, writeStream)
	return nil
}

// CFStreamCreatePairWithSocketToHost creates readable and writable streams connected to a TCP/IP port of a particular host.
//
// Deprecated: Deprecated since macOS 26.4. Use nw_connection_t in Network framework instead
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocketToHost(_:_:_:_:_:)
func CFStreamCreatePairWithSocketToHost(alloc CFAllocatorRef, host CFStringRef, port uint32, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) {
	if callErr := tryCFStreamCreatePairWithSocketToHost(alloc, host, port, readStream, writeStream); callErr != nil {
		panic(callErr)
	}
}

var _cFStringAppend func(theString CFMutableStringRef, appendedString CFStringRef)
var _cFStringAppendErr error

func tryCFStringAppend(theString CFMutableStringRef, appendedString CFStringRef) error {
	if _cFStringAppend == nil {
		return symbolCallError("CFStringAppend", "", _cFStringAppendErr)
	}
	_cFStringAppend(theString, appendedString)
	return nil
}

// CFStringAppend appends the characters of a string to those of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppend(_:_:)
func CFStringAppend(theString CFMutableStringRef, appendedString CFStringRef) {
	if callErr := tryCFStringAppend(theString, appendedString); callErr != nil {
		panic(callErr)
	}
}

var _cFStringAppendCString func(theString CFMutableStringRef, cStr string, encoding uint32)
var _cFStringAppendCStringErr error

func tryCFStringAppendCString(theString CFMutableStringRef, cStr string, encoding uint32) error {
	if _cFStringAppendCString == nil {
		return symbolCallError("CFStringAppendCString", "", _cFStringAppendCStringErr)
	}
	_cFStringAppendCString(theString, cStr, encoding)
	return nil
}

// CFStringAppendCString appends a C string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCString(_:_:_:)
func CFStringAppendCString(theString CFMutableStringRef, cStr string, encoding uint32) {
	if callErr := tryCFStringAppendCString(theString, cStr, encoding); callErr != nil {
		panic(callErr)
	}
}

var _cFStringAppendCharacters func(theString CFMutableStringRef, chars *uint16, numChars int)
var _cFStringAppendCharactersErr error

func tryCFStringAppendCharacters(theString CFMutableStringRef, chars *uint16, numChars int) error {
	if _cFStringAppendCharacters == nil {
		return symbolCallError("CFStringAppendCharacters", "", _cFStringAppendCharactersErr)
	}
	_cFStringAppendCharacters(theString, chars, numChars)
	return nil
}

// CFStringAppendCharacters appends a buffer of Unicode characters to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCharacters(_:_:_:)
func CFStringAppendCharacters(theString CFMutableStringRef, chars *uint16, numChars int) {
	if callErr := tryCFStringAppendCharacters(theString, chars, numChars); callErr != nil {
		panic(callErr)
	}
}

var _cFStringAppendFormat func(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef)
var _cFStringAppendFormatErr error

func tryCFStringAppendFormat(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef) error {
	if _cFStringAppendFormat == nil {
		return symbolCallError("CFStringAppendFormat", "", _cFStringAppendFormatErr)
	}
	_cFStringAppendFormat(theString, formatOptions, format)
	return nil
}

// CFStringAppendFormat appends a formatted string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormat
func CFStringAppendFormat(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef) {
	if callErr := tryCFStringAppendFormat(theString, formatOptions, format); callErr != nil {
		panic(callErr)
	}
}

var _cFStringAppendFormatAndArguments func(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef, arguments kernel.Va_list)
var _cFStringAppendFormatAndArgumentsErr error

func tryCFStringAppendFormatAndArguments(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef, arguments kernel.Va_list) error {
	if _cFStringAppendFormatAndArguments == nil {
		return symbolCallError("CFStringAppendFormatAndArguments", "", _cFStringAppendFormatAndArgumentsErr)
	}
	_cFStringAppendFormatAndArguments(theString, formatOptions, format, arguments)
	return nil
}

// CFStringAppendFormatAndArguments appends a formatted string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormatAndArguments(_:_:_:_:)
func CFStringAppendFormatAndArguments(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef, arguments kernel.Va_list) {
	if callErr := tryCFStringAppendFormatAndArguments(theString, formatOptions, format, arguments); callErr != nil {
		panic(callErr)
	}
}

var _cFStringAppendPascalString func(theString CFMutableStringRef, pStr unsafe.Pointer, encoding uint32)
var _cFStringAppendPascalStringErr error

func tryCFStringAppendPascalString(theString CFMutableStringRef, pStr unsafe.Pointer, encoding uint32) error {
	if _cFStringAppendPascalString == nil {
		return symbolCallError("CFStringAppendPascalString", "", _cFStringAppendPascalStringErr)
	}
	_cFStringAppendPascalString(theString, pStr, encoding)
	return nil
}

// CFStringAppendPascalString appends a Pascal string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendPascalString(_:_:_:)
func CFStringAppendPascalString(theString CFMutableStringRef, pStr unsafe.Pointer, encoding uint32) {
	if callErr := tryCFStringAppendPascalString(theString, pStr, encoding); callErr != nil {
		panic(callErr)
	}
}

var _cFStringCapitalize func(theString CFMutableStringRef, locale CFLocaleRef)
var _cFStringCapitalizeErr error

func tryCFStringCapitalize(theString CFMutableStringRef, locale CFLocaleRef) error {
	if _cFStringCapitalize == nil {
		return symbolCallError("CFStringCapitalize", "", _cFStringCapitalizeErr)
	}
	_cFStringCapitalize(theString, locale)
	return nil
}

// CFStringCapitalize changes the first character in each word of a string to uppercase (if it is a lowercase alphabetical character).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCapitalize(_:_:)
func CFStringCapitalize(theString CFMutableStringRef, locale CFLocaleRef) {
	if callErr := tryCFStringCapitalize(theString, locale); callErr != nil {
		panic(callErr)
	}
}

var _cFStringCompare func(theString1 CFStringRef, theString2 CFStringRef, compareOptions CFStringCompareFlags) CFComparisonResult
var _cFStringCompareErr error

func tryCFStringCompare(theString1 CFStringRef, theString2 CFStringRef, compareOptions CFStringCompareFlags) (CFComparisonResult, error) {
	if _cFStringCompare == nil {
		return *new(CFComparisonResult), symbolCallError("CFStringCompare", "", _cFStringCompareErr)
	}
	return _cFStringCompare(theString1, theString2, compareOptions), nil
}

// CFStringCompare compares one string with another string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompare(_:_:_:)
func CFStringCompare(theString1 CFStringRef, theString2 CFStringRef, compareOptions CFStringCompareFlags) CFComparisonResult {
	result, callErr := tryCFStringCompare(theString1, theString2, compareOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCompareWithOptions func(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags) CFComparisonResult
var _cFStringCompareWithOptionsErr error

func tryCFStringCompareWithOptions(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags) (CFComparisonResult, error) {
	if _cFStringCompareWithOptions == nil {
		return *new(CFComparisonResult), symbolCallError("CFStringCompareWithOptions", "", _cFStringCompareWithOptionsErr)
	}
	return _cFStringCompareWithOptions(theString1, theString2, rangeToCompare, compareOptions), nil
}

// CFStringCompareWithOptions compares a range of the characters in one string with that of another string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptions(_:_:_:_:)
func CFStringCompareWithOptions(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags) CFComparisonResult {
	result, callErr := tryCFStringCompareWithOptions(theString1, theString2, rangeToCompare, compareOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCompareWithOptionsAndLocale func(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags, locale CFLocaleRef) CFComparisonResult
var _cFStringCompareWithOptionsAndLocaleErr error

func tryCFStringCompareWithOptionsAndLocale(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags, locale CFLocaleRef) (CFComparisonResult, error) {
	if _cFStringCompareWithOptionsAndLocale == nil {
		return *new(CFComparisonResult), symbolCallError("CFStringCompareWithOptionsAndLocale", "10.5", _cFStringCompareWithOptionsAndLocaleErr)
	}
	return _cFStringCompareWithOptionsAndLocale(theString1, theString2, rangeToCompare, compareOptions, locale), nil
}

// CFStringCompareWithOptionsAndLocale compares a range of the characters in one string with another string using a given locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptionsAndLocale(_:_:_:_:_:)
func CFStringCompareWithOptionsAndLocale(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags, locale CFLocaleRef) CFComparisonResult {
	result, callErr := tryCFStringCompareWithOptionsAndLocale(theString1, theString2, rangeToCompare, compareOptions, locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringConvertEncodingToIANACharSetName func(encoding uint32) CFStringRef
var _cFStringConvertEncodingToIANACharSetNameErr error

func tryCFStringConvertEncodingToIANACharSetName(encoding uint32) (CFStringRef, error) {
	if _cFStringConvertEncodingToIANACharSetName == nil {
		return 0, symbolCallError("CFStringConvertEncodingToIANACharSetName", "", _cFStringConvertEncodingToIANACharSetNameErr)
	}
	return _cFStringConvertEncodingToIANACharSetName(encoding), nil
}

// CFStringConvertEncodingToIANACharSetName returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToIANACharSetName(_:)
func CFStringConvertEncodingToIANACharSetName(encoding uint32) CFStringRef {
	result, callErr := tryCFStringConvertEncodingToIANACharSetName(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringConvertEncodingToNSStringEncoding func(encoding uint32) uint
var _cFStringConvertEncodingToNSStringEncodingErr error

func tryCFStringConvertEncodingToNSStringEncoding(encoding uint32) (uint, error) {
	if _cFStringConvertEncodingToNSStringEncoding == nil {
		return 0, symbolCallError("CFStringConvertEncodingToNSStringEncoding", "", _cFStringConvertEncodingToNSStringEncodingErr)
	}
	return _cFStringConvertEncodingToNSStringEncoding(encoding), nil
}

// CFStringConvertEncodingToNSStringEncoding returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
func CFStringConvertEncodingToNSStringEncoding(encoding uint32) uint {
	result, callErr := tryCFStringConvertEncodingToNSStringEncoding(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringConvertEncodingToWindowsCodepage func(encoding uint32) uint32
var _cFStringConvertEncodingToWindowsCodepageErr error

func tryCFStringConvertEncodingToWindowsCodepage(encoding uint32) (uint32, error) {
	if _cFStringConvertEncodingToWindowsCodepage == nil {
		return 0, symbolCallError("CFStringConvertEncodingToWindowsCodepage", "", _cFStringConvertEncodingToWindowsCodepageErr)
	}
	return _cFStringConvertEncodingToWindowsCodepage(encoding), nil
}

// CFStringConvertEncodingToWindowsCodepage returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToWindowsCodepage(_:)
func CFStringConvertEncodingToWindowsCodepage(encoding uint32) uint32 {
	result, callErr := tryCFStringConvertEncodingToWindowsCodepage(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringConvertIANACharSetNameToEncoding func(theString CFStringRef) uint32
var _cFStringConvertIANACharSetNameToEncodingErr error

func tryCFStringConvertIANACharSetNameToEncoding(theString CFStringRef) (uint32, error) {
	if _cFStringConvertIANACharSetNameToEncoding == nil {
		return 0, symbolCallError("CFStringConvertIANACharSetNameToEncoding", "", _cFStringConvertIANACharSetNameToEncodingErr)
	}
	return _cFStringConvertIANACharSetNameToEncoding(theString), nil
}

// CFStringConvertIANACharSetNameToEncoding returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertIANACharSetNameToEncoding(_:)
func CFStringConvertIANACharSetNameToEncoding(theString CFStringRef) uint32 {
	result, callErr := tryCFStringConvertIANACharSetNameToEncoding(theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringConvertNSStringEncodingToEncoding func(encoding uint) uint32
var _cFStringConvertNSStringEncodingToEncodingErr error

func tryCFStringConvertNSStringEncodingToEncoding(encoding uint) (uint32, error) {
	if _cFStringConvertNSStringEncodingToEncoding == nil {
		return 0, symbolCallError("CFStringConvertNSStringEncodingToEncoding", "", _cFStringConvertNSStringEncodingToEncodingErr)
	}
	return _cFStringConvertNSStringEncodingToEncoding(encoding), nil
}

// CFStringConvertNSStringEncodingToEncoding returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertNSStringEncodingToEncoding(_:)
func CFStringConvertNSStringEncodingToEncoding(encoding uint) uint32 {
	result, callErr := tryCFStringConvertNSStringEncodingToEncoding(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringConvertWindowsCodepageToEncoding func(codepage uint32) uint32
var _cFStringConvertWindowsCodepageToEncodingErr error

func tryCFStringConvertWindowsCodepageToEncoding(codepage uint32) (uint32, error) {
	if _cFStringConvertWindowsCodepageToEncoding == nil {
		return 0, symbolCallError("CFStringConvertWindowsCodepageToEncoding", "", _cFStringConvertWindowsCodepageToEncodingErr)
	}
	return _cFStringConvertWindowsCodepageToEncoding(codepage), nil
}

// CFStringConvertWindowsCodepageToEncoding returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertWindowsCodepageToEncoding(_:)
func CFStringConvertWindowsCodepageToEncoding(codepage uint32) uint32 {
	result, callErr := tryCFStringConvertWindowsCodepageToEncoding(codepage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateArrayBySeparatingStrings func(alloc CFAllocatorRef, theString CFStringRef, separatorString CFStringRef) CFArrayRef
var _cFStringCreateArrayBySeparatingStringsErr error

func tryCFStringCreateArrayBySeparatingStrings(alloc CFAllocatorRef, theString CFStringRef, separatorString CFStringRef) (CFArrayRef, error) {
	if _cFStringCreateArrayBySeparatingStrings == nil {
		return 0, symbolCallError("CFStringCreateArrayBySeparatingStrings", "", _cFStringCreateArrayBySeparatingStringsErr)
	}
	return _cFStringCreateArrayBySeparatingStrings(alloc, theString, separatorString), nil
}

// CFStringCreateArrayBySeparatingStrings creates an array of CFString objects from a single CFString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayBySeparatingStrings(_:_:_:)
func CFStringCreateArrayBySeparatingStrings(alloc CFAllocatorRef, theString CFStringRef, separatorString CFStringRef) CFArrayRef {
	result, callErr := tryCFStringCreateArrayBySeparatingStrings(alloc, theString, separatorString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateArrayWithFindResults func(alloc CFAllocatorRef, theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) CFArrayRef
var _cFStringCreateArrayWithFindResultsErr error

func tryCFStringCreateArrayWithFindResults(alloc CFAllocatorRef, theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) (CFArrayRef, error) {
	if _cFStringCreateArrayWithFindResults == nil {
		return 0, symbolCallError("CFStringCreateArrayWithFindResults", "", _cFStringCreateArrayWithFindResultsErr)
	}
	return _cFStringCreateArrayWithFindResults(alloc, theString, stringToFind, rangeToSearch, compareOptions), nil
}

// CFStringCreateArrayWithFindResults searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayWithFindResults(_:_:_:_:_:)
func CFStringCreateArrayWithFindResults(alloc CFAllocatorRef, theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) CFArrayRef {
	result, callErr := tryCFStringCreateArrayWithFindResults(alloc, theString, stringToFind, rangeToSearch, compareOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateByCombiningStrings func(alloc CFAllocatorRef, theArray CFArrayRef, separatorString CFStringRef) CFStringRef
var _cFStringCreateByCombiningStringsErr error

func tryCFStringCreateByCombiningStrings(alloc CFAllocatorRef, theArray CFArrayRef, separatorString CFStringRef) (CFStringRef, error) {
	if _cFStringCreateByCombiningStrings == nil {
		return 0, symbolCallError("CFStringCreateByCombiningStrings", "", _cFStringCreateByCombiningStringsErr)
	}
	return _cFStringCreateByCombiningStrings(alloc, theArray, separatorString), nil
}

// CFStringCreateByCombiningStrings creates a single string from the individual CFString objects that comprise the elements of an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateByCombiningStrings(_:_:_:)
func CFStringCreateByCombiningStrings(alloc CFAllocatorRef, theArray CFArrayRef, separatorString CFStringRef) CFStringRef {
	result, callErr := tryCFStringCreateByCombiningStrings(alloc, theArray, separatorString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateCopy func(alloc CFAllocatorRef, theString CFStringRef) CFStringRef
var _cFStringCreateCopyErr error

func tryCFStringCreateCopy(alloc CFAllocatorRef, theString CFStringRef) (CFStringRef, error) {
	if _cFStringCreateCopy == nil {
		return 0, symbolCallError("CFStringCreateCopy", "", _cFStringCreateCopyErr)
	}
	return _cFStringCreateCopy(alloc, theString), nil
}

// CFStringCreateCopy creates an immutable copy of a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateCopy(_:_:)
func CFStringCreateCopy(alloc CFAllocatorRef, theString CFStringRef) CFStringRef {
	result, callErr := tryCFStringCreateCopy(alloc, theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateExternalRepresentation func(alloc CFAllocatorRef, theString CFStringRef, encoding uint32, lossByte uint8) CFDataRef
var _cFStringCreateExternalRepresentationErr error

func tryCFStringCreateExternalRepresentation(alloc CFAllocatorRef, theString CFStringRef, encoding uint32, lossByte uint8) (CFDataRef, error) {
	if _cFStringCreateExternalRepresentation == nil {
		return 0, symbolCallError("CFStringCreateExternalRepresentation", "", _cFStringCreateExternalRepresentationErr)
	}
	return _cFStringCreateExternalRepresentation(alloc, theString, encoding, lossByte), nil
}

// CFStringCreateExternalRepresentation creates an “external representation” of a CFString object, that is, a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateExternalRepresentation(_:_:_:_:)
func CFStringCreateExternalRepresentation(alloc CFAllocatorRef, theString CFStringRef, encoding uint32, lossByte uint8) CFDataRef {
	result, callErr := tryCFStringCreateExternalRepresentation(alloc, theString, encoding, lossByte)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateFromExternalRepresentation func(alloc CFAllocatorRef, data CFDataRef, encoding uint32) CFStringRef
var _cFStringCreateFromExternalRepresentationErr error

func tryCFStringCreateFromExternalRepresentation(alloc CFAllocatorRef, data CFDataRef, encoding uint32) (CFStringRef, error) {
	if _cFStringCreateFromExternalRepresentation == nil {
		return 0, symbolCallError("CFStringCreateFromExternalRepresentation", "", _cFStringCreateFromExternalRepresentationErr)
	}
	return _cFStringCreateFromExternalRepresentation(alloc, data, encoding), nil
}

// CFStringCreateFromExternalRepresentation creates a string from its “external representation.”
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateFromExternalRepresentation(_:_:_:)
func CFStringCreateFromExternalRepresentation(alloc CFAllocatorRef, data CFDataRef, encoding uint32) CFStringRef {
	result, callErr := tryCFStringCreateFromExternalRepresentation(alloc, data, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateMutable func(alloc CFAllocatorRef, maxLength int) CFMutableStringRef
var _cFStringCreateMutableErr error

func tryCFStringCreateMutable(alloc CFAllocatorRef, maxLength int) (CFMutableStringRef, error) {
	if _cFStringCreateMutable == nil {
		return 0, symbolCallError("CFStringCreateMutable", "", _cFStringCreateMutableErr)
	}
	return _cFStringCreateMutable(alloc, maxLength), nil
}

// CFStringCreateMutable creates an empty CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutable(_:_:)
func CFStringCreateMutable(alloc CFAllocatorRef, maxLength int) CFMutableStringRef {
	result, callErr := tryCFStringCreateMutable(alloc, maxLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateMutableCopy func(alloc CFAllocatorRef, maxLength int, theString CFStringRef) CFMutableStringRef
var _cFStringCreateMutableCopyErr error

func tryCFStringCreateMutableCopy(alloc CFAllocatorRef, maxLength int, theString CFStringRef) (CFMutableStringRef, error) {
	if _cFStringCreateMutableCopy == nil {
		return 0, symbolCallError("CFStringCreateMutableCopy", "", _cFStringCreateMutableCopyErr)
	}
	return _cFStringCreateMutableCopy(alloc, maxLength, theString), nil
}

// CFStringCreateMutableCopy creates a mutable copy of a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableCopy(_:_:_:)
func CFStringCreateMutableCopy(alloc CFAllocatorRef, maxLength int, theString CFStringRef) CFMutableStringRef {
	result, callErr := tryCFStringCreateMutableCopy(alloc, maxLength, theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateMutableWithExternalCharactersNoCopy func(alloc CFAllocatorRef, chars *uint16, numChars int, capacity int, externalCharactersAllocator CFAllocatorRef) CFMutableStringRef
var _cFStringCreateMutableWithExternalCharactersNoCopyErr error

func tryCFStringCreateMutableWithExternalCharactersNoCopy(alloc CFAllocatorRef, chars *uint16, numChars int, capacity int, externalCharactersAllocator CFAllocatorRef) (CFMutableStringRef, error) {
	if _cFStringCreateMutableWithExternalCharactersNoCopy == nil {
		return 0, symbolCallError("CFStringCreateMutableWithExternalCharactersNoCopy", "", _cFStringCreateMutableWithExternalCharactersNoCopyErr)
	}
	return _cFStringCreateMutableWithExternalCharactersNoCopy(alloc, chars, numChars, capacity, externalCharactersAllocator), nil
}

// CFStringCreateMutableWithExternalCharactersNoCopy creates a CFMutableString object whose Unicode character buffer is controlled externally.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableWithExternalCharactersNoCopy(_:_:_:_:_:)
func CFStringCreateMutableWithExternalCharactersNoCopy(alloc CFAllocatorRef, chars *uint16, numChars int, capacity int, externalCharactersAllocator CFAllocatorRef) CFMutableStringRef {
	result, callErr := tryCFStringCreateMutableWithExternalCharactersNoCopy(alloc, chars, numChars, capacity, externalCharactersAllocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateStringWithValidatedFormat func(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, errorPtr *CFErrorRef) CFStringRef
var _cFStringCreateStringWithValidatedFormatErr error

func tryCFStringCreateStringWithValidatedFormat(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, errorPtr *CFErrorRef) (CFStringRef, error) {
	if _cFStringCreateStringWithValidatedFormat == nil {
		return 0, symbolCallError("CFStringCreateStringWithValidatedFormat", "13.0", _cFStringCreateStringWithValidatedFormatErr)
	}
	return _cFStringCreateStringWithValidatedFormat(alloc, formatOptions, validFormatSpecifiers, format, errorPtr), nil
}

// CFStringCreateStringWithValidatedFormat.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormat
func CFStringCreateStringWithValidatedFormat(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, errorPtr *CFErrorRef) CFStringRef {
	result, callErr := tryCFStringCreateStringWithValidatedFormat(alloc, formatOptions, validFormatSpecifiers, format, errorPtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateStringWithValidatedFormatAndArguments func(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, arguments kernel.Va_list, errorPtr *CFErrorRef) CFStringRef
var _cFStringCreateStringWithValidatedFormatAndArgumentsErr error

func tryCFStringCreateStringWithValidatedFormatAndArguments(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, arguments kernel.Va_list, errorPtr *CFErrorRef) (CFStringRef, error) {
	if _cFStringCreateStringWithValidatedFormatAndArguments == nil {
		return 0, symbolCallError("CFStringCreateStringWithValidatedFormatAndArguments", "13.0", _cFStringCreateStringWithValidatedFormatAndArgumentsErr)
	}
	return _cFStringCreateStringWithValidatedFormatAndArguments(alloc, formatOptions, validFormatSpecifiers, format, arguments, errorPtr), nil
}

// CFStringCreateStringWithValidatedFormatAndArguments.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormatAndArguments
func CFStringCreateStringWithValidatedFormatAndArguments(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, arguments kernel.Va_list, errorPtr *CFErrorRef) CFStringRef {
	result, callErr := tryCFStringCreateStringWithValidatedFormatAndArguments(alloc, formatOptions, validFormatSpecifiers, format, arguments, errorPtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithBytes func(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool) CFStringRef
var _cFStringCreateWithBytesErr error

func tryCFStringCreateWithBytes(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool) (CFStringRef, error) {
	if _cFStringCreateWithBytes == nil {
		return 0, symbolCallError("CFStringCreateWithBytes", "", _cFStringCreateWithBytesErr)
	}
	return _cFStringCreateWithBytes(alloc, bytes, numBytes, encoding, isExternalRepresentation), nil
}

// CFStringCreateWithBytes creates a string from a buffer containing characters in a specified encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytes(_:_:_:_:_:)
func CFStringCreateWithBytes(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool) CFStringRef {
	result, callErr := tryCFStringCreateWithBytes(alloc, bytes, numBytes, encoding, isExternalRepresentation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithBytesNoCopy func(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool, contentsDeallocator CFAllocatorRef) CFStringRef
var _cFStringCreateWithBytesNoCopyErr error

func tryCFStringCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool, contentsDeallocator CFAllocatorRef) (CFStringRef, error) {
	if _cFStringCreateWithBytesNoCopy == nil {
		return 0, symbolCallError("CFStringCreateWithBytesNoCopy", "", _cFStringCreateWithBytesNoCopyErr)
	}
	return _cFStringCreateWithBytesNoCopy(alloc, bytes, numBytes, encoding, isExternalRepresentation, contentsDeallocator), nil
}

// CFStringCreateWithBytesNoCopy creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytesNoCopy(_:_:_:_:_:_:)
func CFStringCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool, contentsDeallocator CFAllocatorRef) CFStringRef {
	result, callErr := tryCFStringCreateWithBytesNoCopy(alloc, bytes, numBytes, encoding, isExternalRepresentation, contentsDeallocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithCString func(alloc CFAllocatorRef, cStr string, encoding uint32) CFStringRef
var _cFStringCreateWithCStringErr error

func tryCFStringCreateWithCString(alloc CFAllocatorRef, cStr string, encoding uint32) (CFStringRef, error) {
	if _cFStringCreateWithCString == nil {
		return 0, symbolCallError("CFStringCreateWithCString", "", _cFStringCreateWithCStringErr)
	}
	return _cFStringCreateWithCString(alloc, cStr, encoding), nil
}

// CFStringCreateWithCString creates an immutable string from a C string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCString(_:_:_:)
func CFStringCreateWithCString(alloc CFAllocatorRef, cStr string, encoding uint32) CFStringRef {
	result, callErr := tryCFStringCreateWithCString(alloc, cStr, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithCStringNoCopy func(alloc CFAllocatorRef, cStr string, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef
var _cFStringCreateWithCStringNoCopyErr error

func tryCFStringCreateWithCStringNoCopy(alloc CFAllocatorRef, cStr string, encoding uint32, contentsDeallocator CFAllocatorRef) (CFStringRef, error) {
	if _cFStringCreateWithCStringNoCopy == nil {
		return 0, symbolCallError("CFStringCreateWithCStringNoCopy", "", _cFStringCreateWithCStringNoCopyErr)
	}
	return _cFStringCreateWithCStringNoCopy(alloc, cStr, encoding, contentsDeallocator), nil
}

// CFStringCreateWithCStringNoCopy creates a CFString object from an external C string buffer that might serve as the backing store for the object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCStringNoCopy(_:_:_:_:)
func CFStringCreateWithCStringNoCopy(alloc CFAllocatorRef, cStr string, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef {
	result, callErr := tryCFStringCreateWithCStringNoCopy(alloc, cStr, encoding, contentsDeallocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithCharacters func(alloc CFAllocatorRef, chars *uint16, numChars int) CFStringRef
var _cFStringCreateWithCharactersErr error

func tryCFStringCreateWithCharacters(alloc CFAllocatorRef, chars *uint16, numChars int) (CFStringRef, error) {
	if _cFStringCreateWithCharacters == nil {
		return 0, symbolCallError("CFStringCreateWithCharacters", "", _cFStringCreateWithCharactersErr)
	}
	return _cFStringCreateWithCharacters(alloc, chars, numChars), nil
}

// CFStringCreateWithCharacters creates a string from a buffer of Unicode characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharacters(_:_:_:)
func CFStringCreateWithCharacters(alloc CFAllocatorRef, chars *uint16, numChars int) CFStringRef {
	result, callErr := tryCFStringCreateWithCharacters(alloc, chars, numChars)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithCharactersNoCopy func(alloc CFAllocatorRef, chars *uint16, numChars int, contentsDeallocator CFAllocatorRef) CFStringRef
var _cFStringCreateWithCharactersNoCopyErr error

func tryCFStringCreateWithCharactersNoCopy(alloc CFAllocatorRef, chars *uint16, numChars int, contentsDeallocator CFAllocatorRef) (CFStringRef, error) {
	if _cFStringCreateWithCharactersNoCopy == nil {
		return 0, symbolCallError("CFStringCreateWithCharactersNoCopy", "", _cFStringCreateWithCharactersNoCopyErr)
	}
	return _cFStringCreateWithCharactersNoCopy(alloc, chars, numChars, contentsDeallocator), nil
}

// CFStringCreateWithCharactersNoCopy creates a string from a buffer of Unicode characters that might serve as the backing store for the object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharactersNoCopy(_:_:_:_:)
func CFStringCreateWithCharactersNoCopy(alloc CFAllocatorRef, chars *uint16, numChars int, contentsDeallocator CFAllocatorRef) CFStringRef {
	result, callErr := tryCFStringCreateWithCharactersNoCopy(alloc, chars, numChars, contentsDeallocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithFileSystemRepresentation func(alloc CFAllocatorRef, buffer *byte) CFStringRef
var _cFStringCreateWithFileSystemRepresentationErr error

func tryCFStringCreateWithFileSystemRepresentation(alloc CFAllocatorRef, buffer *byte) (CFStringRef, error) {
	if _cFStringCreateWithFileSystemRepresentation == nil {
		return 0, symbolCallError("CFStringCreateWithFileSystemRepresentation", "", _cFStringCreateWithFileSystemRepresentationErr)
	}
	return _cFStringCreateWithFileSystemRepresentation(alloc, buffer), nil
}

// CFStringCreateWithFileSystemRepresentation creates a CFString from a zero-terminated POSIX file system representation.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFileSystemRepresentation(_:_:)
func CFStringCreateWithFileSystemRepresentation(alloc CFAllocatorRef, buffer *byte) CFStringRef {
	result, callErr := tryCFStringCreateWithFileSystemRepresentation(alloc, buffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithFormat func(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef) CFStringRef
var _cFStringCreateWithFormatErr error

func tryCFStringCreateWithFormat(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef) (CFStringRef, error) {
	if _cFStringCreateWithFormat == nil {
		return 0, symbolCallError("CFStringCreateWithFormat", "", _cFStringCreateWithFormatErr)
	}
	return _cFStringCreateWithFormat(alloc, formatOptions, format), nil
}

// CFStringCreateWithFormat creates an immutable string from a formatted string and a variable number of arguments.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormat
func CFStringCreateWithFormat(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef) CFStringRef {
	result, callErr := tryCFStringCreateWithFormat(alloc, formatOptions, format)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithFormatAndArguments func(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef, arguments kernel.Va_list) CFStringRef
var _cFStringCreateWithFormatAndArgumentsErr error

func tryCFStringCreateWithFormatAndArguments(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef, arguments kernel.Va_list) (CFStringRef, error) {
	if _cFStringCreateWithFormatAndArguments == nil {
		return 0, symbolCallError("CFStringCreateWithFormatAndArguments", "", _cFStringCreateWithFormatAndArgumentsErr)
	}
	return _cFStringCreateWithFormatAndArguments(alloc, formatOptions, format, arguments), nil
}

// CFStringCreateWithFormatAndArguments creates an immutable string from a formatted string and a variable number of arguments (specified in a parameter of type `va_list`).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormatAndArguments(_:_:_:_:)
func CFStringCreateWithFormatAndArguments(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef, arguments kernel.Va_list) CFStringRef {
	result, callErr := tryCFStringCreateWithFormatAndArguments(alloc, formatOptions, format, arguments)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithPascalString func(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32) CFStringRef
var _cFStringCreateWithPascalStringErr error

func tryCFStringCreateWithPascalString(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32) (CFStringRef, error) {
	if _cFStringCreateWithPascalString == nil {
		return 0, symbolCallError("CFStringCreateWithPascalString", "", _cFStringCreateWithPascalStringErr)
	}
	return _cFStringCreateWithPascalString(alloc, pStr, encoding), nil
}

// CFStringCreateWithPascalString creates an immutable CFString object from a Pascal string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalString(_:_:_:)
func CFStringCreateWithPascalString(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32) CFStringRef {
	result, callErr := tryCFStringCreateWithPascalString(alloc, pStr, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithPascalStringNoCopy func(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef
var _cFStringCreateWithPascalStringNoCopyErr error

func tryCFStringCreateWithPascalStringNoCopy(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32, contentsDeallocator CFAllocatorRef) (CFStringRef, error) {
	if _cFStringCreateWithPascalStringNoCopy == nil {
		return 0, symbolCallError("CFStringCreateWithPascalStringNoCopy", "", _cFStringCreateWithPascalStringNoCopyErr)
	}
	return _cFStringCreateWithPascalStringNoCopy(alloc, pStr, encoding, contentsDeallocator), nil
}

// CFStringCreateWithPascalStringNoCopy creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalStringNoCopy(_:_:_:_:)
func CFStringCreateWithPascalStringNoCopy(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef {
	result, callErr := tryCFStringCreateWithPascalStringNoCopy(alloc, pStr, encoding, contentsDeallocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringCreateWithSubstring func(alloc CFAllocatorRef, str CFStringRef, range_ CFRange) CFStringRef
var _cFStringCreateWithSubstringErr error

func tryCFStringCreateWithSubstring(alloc CFAllocatorRef, str CFStringRef, range_ CFRange) (CFStringRef, error) {
	if _cFStringCreateWithSubstring == nil {
		return 0, symbolCallError("CFStringCreateWithSubstring", "", _cFStringCreateWithSubstringErr)
	}
	return _cFStringCreateWithSubstring(alloc, str, range_), nil
}

// CFStringCreateWithSubstring creates an immutable string from a segment (substring) of an existing string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithSubstring(_:_:_:)
func CFStringCreateWithSubstring(alloc CFAllocatorRef, str CFStringRef, range_ CFRange) CFStringRef {
	result, callErr := tryCFStringCreateWithSubstring(alloc, str, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringDelete func(theString CFMutableStringRef, range_ CFRange)
var _cFStringDeleteErr error

func tryCFStringDelete(theString CFMutableStringRef, range_ CFRange) error {
	if _cFStringDelete == nil {
		return symbolCallError("CFStringDelete", "", _cFStringDeleteErr)
	}
	_cFStringDelete(theString, range_)
	return nil
}

// CFStringDelete deletes a range of characters in a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringDelete(_:_:)
func CFStringDelete(theString CFMutableStringRef, range_ CFRange) {
	if callErr := tryCFStringDelete(theString, range_); callErr != nil {
		panic(callErr)
	}
}

var _cFStringFind func(theString CFStringRef, stringToFind CFStringRef, compareOptions CFStringCompareFlags) CFRange
var _cFStringFindErr error

func tryCFStringFind(theString CFStringRef, stringToFind CFStringRef, compareOptions CFStringCompareFlags) (CFRange, error) {
	if _cFStringFind == nil {
		return CFRange{}, symbolCallError("CFStringFind", "", _cFStringFindErr)
	}
	return _cFStringFind(theString, stringToFind, compareOptions), nil
}

// CFStringFind searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFind(_:_:_:)
func CFStringFind(theString CFStringRef, stringToFind CFStringRef, compareOptions CFStringCompareFlags) CFRange {
	result, callErr := tryCFStringFind(theString, stringToFind, compareOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringFindAndReplace func(theString CFMutableStringRef, stringToFind CFStringRef, replacementString CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) int
var _cFStringFindAndReplaceErr error

func tryCFStringFindAndReplace(theString CFMutableStringRef, stringToFind CFStringRef, replacementString CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) (int, error) {
	if _cFStringFindAndReplace == nil {
		return 0, symbolCallError("CFStringFindAndReplace", "", _cFStringFindAndReplaceErr)
	}
	return _cFStringFindAndReplace(theString, stringToFind, replacementString, rangeToSearch, compareOptions), nil
}

// CFStringFindAndReplace replaces all occurrences of a substring within a given range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindAndReplace(_:_:_:_:_:)
func CFStringFindAndReplace(theString CFMutableStringRef, stringToFind CFStringRef, replacementString CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) int {
	result, callErr := tryCFStringFindAndReplace(theString, stringToFind, replacementString, rangeToSearch, compareOptions)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringFindCharacterFromSet func(theString CFStringRef, theSet CFCharacterSetRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool
var _cFStringFindCharacterFromSetErr error

func tryCFStringFindCharacterFromSet(theString CFStringRef, theSet CFCharacterSetRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) (bool, error) {
	if _cFStringFindCharacterFromSet == nil {
		return false, symbolCallError("CFStringFindCharacterFromSet", "", _cFStringFindCharacterFromSetErr)
	}
	return _cFStringFindCharacterFromSet(theString, theSet, rangeToSearch, searchOptions, result), nil
}

// CFStringFindCharacterFromSet query the range of the first character contained in the specified character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindCharacterFromSet(_:_:_:_:_:)
func CFStringFindCharacterFromSet(theString CFStringRef, theSet CFCharacterSetRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool {
	result0, callErr := tryCFStringFindCharacterFromSet(theString, theSet, rangeToSearch, searchOptions, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _cFStringFindWithOptions func(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool
var _cFStringFindWithOptionsErr error

func tryCFStringFindWithOptions(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) (bool, error) {
	if _cFStringFindWithOptions == nil {
		return false, symbolCallError("CFStringFindWithOptions", "", _cFStringFindWithOptionsErr)
	}
	return _cFStringFindWithOptions(theString, stringToFind, rangeToSearch, searchOptions, result), nil
}

// CFStringFindWithOptions searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptions(_:_:_:_:_:)
func CFStringFindWithOptions(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool {
	result0, callErr := tryCFStringFindWithOptions(theString, stringToFind, rangeToSearch, searchOptions, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _cFStringFindWithOptionsAndLocale func(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, locale CFLocaleRef, result *CFRange) bool
var _cFStringFindWithOptionsAndLocaleErr error

func tryCFStringFindWithOptionsAndLocale(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, locale CFLocaleRef, result *CFRange) (bool, error) {
	if _cFStringFindWithOptionsAndLocale == nil {
		return false, symbolCallError("CFStringFindWithOptionsAndLocale", "10.5", _cFStringFindWithOptionsAndLocaleErr)
	}
	return _cFStringFindWithOptionsAndLocale(theString, stringToFind, rangeToSearch, searchOptions, locale, result), nil
}

// CFStringFindWithOptionsAndLocale returns a Boolean value that indicates whether a given string was found in a given source string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptionsAndLocale(_:_:_:_:_:_:)
func CFStringFindWithOptionsAndLocale(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, locale CFLocaleRef, result *CFRange) bool {
	result0, callErr := tryCFStringFindWithOptionsAndLocale(theString, stringToFind, rangeToSearch, searchOptions, locale, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _cFStringFold func(theString CFMutableStringRef, theFlags CFStringCompareFlags, theLocale CFLocaleRef)
var _cFStringFoldErr error

func tryCFStringFold(theString CFMutableStringRef, theFlags CFStringCompareFlags, theLocale CFLocaleRef) error {
	if _cFStringFold == nil {
		return symbolCallError("CFStringFold", "10.5", _cFStringFoldErr)
	}
	_cFStringFold(theString, theFlags, theLocale)
	return nil
}

// CFStringFold folds a given string into the form specified by optional flags.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFold(_:_:_:)
func CFStringFold(theString CFMutableStringRef, theFlags CFStringCompareFlags, theLocale CFLocaleRef) {
	if callErr := tryCFStringFold(theString, theFlags, theLocale); callErr != nil {
		panic(callErr)
	}
}

var _cFStringGetBytes func(theString CFStringRef, range_ CFRange, encoding uint32, lossByte uint8, isExternalRepresentation bool, buffer *uint8, maxBufLen int, usedBufLen *int) int
var _cFStringGetBytesErr error

func tryCFStringGetBytes(theString CFStringRef, range_ CFRange, encoding uint32, lossByte uint8, isExternalRepresentation bool, buffer *uint8, maxBufLen int, usedBufLen *int) (int, error) {
	if _cFStringGetBytes == nil {
		return 0, symbolCallError("CFStringGetBytes", "", _cFStringGetBytesErr)
	}
	return _cFStringGetBytes(theString, range_, encoding, lossByte, isExternalRepresentation, buffer, maxBufLen, usedBufLen), nil
}

// CFStringGetBytes fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetBytes(_:_:_:_:_:_:_:_:)
func CFStringGetBytes(theString CFStringRef, range_ CFRange, encoding uint32, lossByte uint8, isExternalRepresentation bool, buffer *uint8, maxBufLen int, usedBufLen *int) int {
	result, callErr := tryCFStringGetBytes(theString, range_, encoding, lossByte, isExternalRepresentation, buffer, maxBufLen, usedBufLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetCString func(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool
var _cFStringGetCStringErr error

func tryCFStringGetCString(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) (bool, error) {
	if _cFStringGetCString == nil {
		return false, symbolCallError("CFStringGetCString", "", _cFStringGetCStringErr)
	}
	return _cFStringGetCString(theString, buffer, bufferSize, encoding), nil
}

// CFStringGetCString copies the character contents of a string to a local C string buffer after converting the characters to a given encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCString(_:_:_:_:)
func CFStringGetCString(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool {
	result, callErr := tryCFStringGetCString(theString, buffer, bufferSize, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetCStringPtr func(theString CFStringRef, encoding uint32) *byte
var _cFStringGetCStringPtrErr error

func tryCFStringGetCStringPtr(theString CFStringRef, encoding uint32) (*byte, error) {
	if _cFStringGetCStringPtr == nil {
		return nil, symbolCallError("CFStringGetCStringPtr", "", _cFStringGetCStringPtrErr)
	}
	return _cFStringGetCStringPtr(theString, encoding), nil
}

// CFStringGetCStringPtr quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCStringPtr(_:_:)
func CFStringGetCStringPtr(theString CFStringRef, encoding uint32) *byte {
	result, callErr := tryCFStringGetCStringPtr(theString, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetCharacterAtIndex func(theString CFStringRef, idx int) uint16
var _cFStringGetCharacterAtIndexErr error

func tryCFStringGetCharacterAtIndex(theString CFStringRef, idx int) (uint16, error) {
	if _cFStringGetCharacterAtIndex == nil {
		return 0, symbolCallError("CFStringGetCharacterAtIndex", "", _cFStringGetCharacterAtIndexErr)
	}
	return _cFStringGetCharacterAtIndex(theString, idx), nil
}

// CFStringGetCharacterAtIndex returns the Unicode character at a specified location in a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacterAtIndex(_:_:)
func CFStringGetCharacterAtIndex(theString CFStringRef, idx int) uint16 {
	result, callErr := tryCFStringGetCharacterAtIndex(theString, idx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetCharacters func(theString CFStringRef, range_ CFRange, buffer *uint16)
var _cFStringGetCharactersErr error

func tryCFStringGetCharacters(theString CFStringRef, range_ CFRange, buffer *uint16) error {
	if _cFStringGetCharacters == nil {
		return symbolCallError("CFStringGetCharacters", "", _cFStringGetCharactersErr)
	}
	_cFStringGetCharacters(theString, range_, buffer)
	return nil
}

// CFStringGetCharacters copies a range of the Unicode characters from a string to a user-provided buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacters(_:_:_:)
func CFStringGetCharacters(theString CFStringRef, range_ CFRange, buffer *uint16) {
	if callErr := tryCFStringGetCharacters(theString, range_, buffer); callErr != nil {
		panic(callErr)
	}
}

var _cFStringGetCharactersPtr func(theString CFStringRef) *uint16
var _cFStringGetCharactersPtrErr error

func tryCFStringGetCharactersPtr(theString CFStringRef) (*uint16, error) {
	if _cFStringGetCharactersPtr == nil {
		return nil, symbolCallError("CFStringGetCharactersPtr", "", _cFStringGetCharactersPtrErr)
	}
	return _cFStringGetCharactersPtr(theString), nil
}

// CFStringGetCharactersPtr quickly obtains a pointer to the contents of a string as a buffer of Unicode characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharactersPtr(_:)
func CFStringGetCharactersPtr(theString CFStringRef) *uint16 {
	result, callErr := tryCFStringGetCharactersPtr(theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetDoubleValue func(str CFStringRef) float64
var _cFStringGetDoubleValueErr error

func tryCFStringGetDoubleValue(str CFStringRef) (float64, error) {
	if _cFStringGetDoubleValue == nil {
		return 0.0, symbolCallError("CFStringGetDoubleValue", "", _cFStringGetDoubleValueErr)
	}
	return _cFStringGetDoubleValue(str), nil
}

// CFStringGetDoubleValue returns the primary `double` value represented by a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetDoubleValue(_:)
func CFStringGetDoubleValue(str CFStringRef) float64 {
	result, callErr := tryCFStringGetDoubleValue(str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetFastestEncoding func(theString CFStringRef) uint32
var _cFStringGetFastestEncodingErr error

func tryCFStringGetFastestEncoding(theString CFStringRef) (uint32, error) {
	if _cFStringGetFastestEncoding == nil {
		return 0, symbolCallError("CFStringGetFastestEncoding", "", _cFStringGetFastestEncodingErr)
	}
	return _cFStringGetFastestEncoding(theString), nil
}

// CFStringGetFastestEncoding returns for a CFString object the character encoding that requires the least conversion time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFastestEncoding(_:)
func CFStringGetFastestEncoding(theString CFStringRef) uint32 {
	result, callErr := tryCFStringGetFastestEncoding(theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetFileSystemRepresentation func(string_ CFStringRef, buffer *byte, maxBufLen int) bool
var _cFStringGetFileSystemRepresentationErr error

func tryCFStringGetFileSystemRepresentation(string_ CFStringRef, buffer *byte, maxBufLen int) (bool, error) {
	if _cFStringGetFileSystemRepresentation == nil {
		return false, symbolCallError("CFStringGetFileSystemRepresentation", "", _cFStringGetFileSystemRepresentationErr)
	}
	return _cFStringGetFileSystemRepresentation(string_, buffer, maxBufLen), nil
}

// CFStringGetFileSystemRepresentation extracts the contents of a string as a [NULL]-terminated 8-bit string appropriate for passing to POSIX APIs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFileSystemRepresentation(_:_:_:)
func CFStringGetFileSystemRepresentation(string_ CFStringRef, buffer *byte, maxBufLen int) bool {
	result, callErr := tryCFStringGetFileSystemRepresentation(string_, buffer, maxBufLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetHyphenationLocationBeforeIndex func(string_ CFStringRef, location int, limitRange CFRange, options uint64, locale CFLocaleRef, character *uint32) int
var _cFStringGetHyphenationLocationBeforeIndexErr error

func tryCFStringGetHyphenationLocationBeforeIndex(string_ CFStringRef, location int, limitRange CFRange, options uint64, locale CFLocaleRef, character *uint32) (int, error) {
	if _cFStringGetHyphenationLocationBeforeIndex == nil {
		return 0, symbolCallError("CFStringGetHyphenationLocationBeforeIndex", "10.7", _cFStringGetHyphenationLocationBeforeIndexErr)
	}
	return _cFStringGetHyphenationLocationBeforeIndex(string_, location, limitRange, options, locale, character), nil
}

// CFStringGetHyphenationLocationBeforeIndex retrieve the first potential hyphenation location found before the specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetHyphenationLocationBeforeIndex(_:_:_:_:_:_:)
func CFStringGetHyphenationLocationBeforeIndex(string_ CFStringRef, location int, limitRange CFRange, options uint64, locale CFLocaleRef, character *uint32) int {
	result, callErr := tryCFStringGetHyphenationLocationBeforeIndex(string_, location, limitRange, options, locale, character)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetIntValue func(str CFStringRef) int32
var _cFStringGetIntValueErr error

func tryCFStringGetIntValue(str CFStringRef) (int32, error) {
	if _cFStringGetIntValue == nil {
		return 0, symbolCallError("CFStringGetIntValue", "", _cFStringGetIntValueErr)
	}
	return _cFStringGetIntValue(str), nil
}

// CFStringGetIntValue returns the integer value represented by a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetIntValue(_:)
func CFStringGetIntValue(str CFStringRef) int32 {
	result, callErr := tryCFStringGetIntValue(str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetLength func(theString CFStringRef) int
var _cFStringGetLengthErr error

func tryCFStringGetLength(theString CFStringRef) (int, error) {
	if _cFStringGetLength == nil {
		return 0, symbolCallError("CFStringGetLength", "", _cFStringGetLengthErr)
	}
	return _cFStringGetLength(theString), nil
}

// CFStringGetLength returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLength(_:)
func CFStringGetLength(theString CFStringRef) int {
	result, callErr := tryCFStringGetLength(theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetLineBounds func(theString CFStringRef, range_ CFRange, lineBeginIndex *int, lineEndIndex *int, contentsEndIndex *int)
var _cFStringGetLineBoundsErr error

func tryCFStringGetLineBounds(theString CFStringRef, range_ CFRange, lineBeginIndex *int, lineEndIndex *int, contentsEndIndex *int) error {
	if _cFStringGetLineBounds == nil {
		return symbolCallError("CFStringGetLineBounds", "", _cFStringGetLineBoundsErr)
	}
	_cFStringGetLineBounds(theString, range_, lineBeginIndex, lineEndIndex, contentsEndIndex)
	return nil
}

// CFStringGetLineBounds given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLineBounds(_:_:_:_:_:)
func CFStringGetLineBounds(theString CFStringRef, range_ CFRange, lineBeginIndex *int, lineEndIndex *int, contentsEndIndex *int) {
	if callErr := tryCFStringGetLineBounds(theString, range_, lineBeginIndex, lineEndIndex, contentsEndIndex); callErr != nil {
		panic(callErr)
	}
}

var _cFStringGetListOfAvailableEncodings func() *uint32
var _cFStringGetListOfAvailableEncodingsErr error

func tryCFStringGetListOfAvailableEncodings() (*uint32, error) {
	if _cFStringGetListOfAvailableEncodings == nil {
		return nil, symbolCallError("CFStringGetListOfAvailableEncodings", "", _cFStringGetListOfAvailableEncodingsErr)
	}
	return _cFStringGetListOfAvailableEncodings(), nil
}

// CFStringGetListOfAvailableEncodings returns a pointer to a list of string encodings supported by the current system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetListOfAvailableEncodings()
func CFStringGetListOfAvailableEncodings() *uint32 {
	result, callErr := tryCFStringGetListOfAvailableEncodings()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetMaximumSizeForEncoding func(length int, encoding uint32) int
var _cFStringGetMaximumSizeForEncodingErr error

func tryCFStringGetMaximumSizeForEncoding(length int, encoding uint32) (int, error) {
	if _cFStringGetMaximumSizeForEncoding == nil {
		return 0, symbolCallError("CFStringGetMaximumSizeForEncoding", "", _cFStringGetMaximumSizeForEncodingErr)
	}
	return _cFStringGetMaximumSizeForEncoding(length, encoding), nil
}

// CFStringGetMaximumSizeForEncoding returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeForEncoding(_:_:)
func CFStringGetMaximumSizeForEncoding(length int, encoding uint32) int {
	result, callErr := tryCFStringGetMaximumSizeForEncoding(length, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetMaximumSizeOfFileSystemRepresentation func(string_ CFStringRef) int
var _cFStringGetMaximumSizeOfFileSystemRepresentationErr error

func tryCFStringGetMaximumSizeOfFileSystemRepresentation(string_ CFStringRef) (int, error) {
	if _cFStringGetMaximumSizeOfFileSystemRepresentation == nil {
		return 0, symbolCallError("CFStringGetMaximumSizeOfFileSystemRepresentation", "", _cFStringGetMaximumSizeOfFileSystemRepresentationErr)
	}
	return _cFStringGetMaximumSizeOfFileSystemRepresentation(string_), nil
}

// CFStringGetMaximumSizeOfFileSystemRepresentation determines the upper bound on the number of bytes required to hold the file system representation of the string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeOfFileSystemRepresentation(_:)
func CFStringGetMaximumSizeOfFileSystemRepresentation(string_ CFStringRef) int {
	result, callErr := tryCFStringGetMaximumSizeOfFileSystemRepresentation(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetMostCompatibleMacStringEncoding func(encoding uint32) uint32
var _cFStringGetMostCompatibleMacStringEncodingErr error

func tryCFStringGetMostCompatibleMacStringEncoding(encoding uint32) (uint32, error) {
	if _cFStringGetMostCompatibleMacStringEncoding == nil {
		return 0, symbolCallError("CFStringGetMostCompatibleMacStringEncoding", "", _cFStringGetMostCompatibleMacStringEncodingErr)
	}
	return _cFStringGetMostCompatibleMacStringEncoding(encoding), nil
}

// CFStringGetMostCompatibleMacStringEncoding returns the most compatible Mac OS script value for the given input encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMostCompatibleMacStringEncoding(_:)
func CFStringGetMostCompatibleMacStringEncoding(encoding uint32) uint32 {
	result, callErr := tryCFStringGetMostCompatibleMacStringEncoding(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetNameOfEncoding func(encoding uint32) CFStringRef
var _cFStringGetNameOfEncodingErr error

func tryCFStringGetNameOfEncoding(encoding uint32) (CFStringRef, error) {
	if _cFStringGetNameOfEncoding == nil {
		return 0, symbolCallError("CFStringGetNameOfEncoding", "", _cFStringGetNameOfEncodingErr)
	}
	return _cFStringGetNameOfEncoding(encoding), nil
}

// CFStringGetNameOfEncoding returns the canonical name of a specified string encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetNameOfEncoding(_:)
func CFStringGetNameOfEncoding(encoding uint32) CFStringRef {
	result, callErr := tryCFStringGetNameOfEncoding(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetParagraphBounds func(string_ CFStringRef, range_ CFRange, parBeginIndex *int, parEndIndex *int, contentsEndIndex *int)
var _cFStringGetParagraphBoundsErr error

func tryCFStringGetParagraphBounds(string_ CFStringRef, range_ CFRange, parBeginIndex *int, parEndIndex *int, contentsEndIndex *int) error {
	if _cFStringGetParagraphBounds == nil {
		return symbolCallError("CFStringGetParagraphBounds", "10.5", _cFStringGetParagraphBoundsErr)
	}
	_cFStringGetParagraphBounds(string_, range_, parBeginIndex, parEndIndex, contentsEndIndex)
	return nil
}

// CFStringGetParagraphBounds given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetParagraphBounds(_:_:_:_:_:)
func CFStringGetParagraphBounds(string_ CFStringRef, range_ CFRange, parBeginIndex *int, parEndIndex *int, contentsEndIndex *int) {
	if callErr := tryCFStringGetParagraphBounds(string_, range_, parBeginIndex, parEndIndex, contentsEndIndex); callErr != nil {
		panic(callErr)
	}
}

var _cFStringGetPascalString func(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool
var _cFStringGetPascalStringErr error

func tryCFStringGetPascalString(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) (bool, error) {
	if _cFStringGetPascalString == nil {
		return false, symbolCallError("CFStringGetPascalString", "", _cFStringGetPascalStringErr)
	}
	return _cFStringGetPascalString(theString, buffer, bufferSize, encoding), nil
}

// CFStringGetPascalString copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalString(_:_:_:_:)
func CFStringGetPascalString(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool {
	result, callErr := tryCFStringGetPascalString(theString, buffer, bufferSize, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetPascalStringPtr func(theString CFStringRef, encoding uint32) *byte
var _cFStringGetPascalStringPtrErr error

func tryCFStringGetPascalStringPtr(theString CFStringRef, encoding uint32) (*byte, error) {
	if _cFStringGetPascalStringPtr == nil {
		return nil, symbolCallError("CFStringGetPascalStringPtr", "", _cFStringGetPascalStringPtrErr)
	}
	return _cFStringGetPascalStringPtr(theString, encoding), nil
}

// CFStringGetPascalStringPtr quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalStringPtr(_:_:)
func CFStringGetPascalStringPtr(theString CFStringRef, encoding uint32) *byte {
	result, callErr := tryCFStringGetPascalStringPtr(theString, encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetRangeOfComposedCharactersAtIndex func(theString CFStringRef, theIndex int) CFRange
var _cFStringGetRangeOfComposedCharactersAtIndexErr error

func tryCFStringGetRangeOfComposedCharactersAtIndex(theString CFStringRef, theIndex int) (CFRange, error) {
	if _cFStringGetRangeOfComposedCharactersAtIndex == nil {
		return CFRange{}, symbolCallError("CFStringGetRangeOfComposedCharactersAtIndex", "", _cFStringGetRangeOfComposedCharactersAtIndexErr)
	}
	return _cFStringGetRangeOfComposedCharactersAtIndex(theString, theIndex), nil
}

// CFStringGetRangeOfComposedCharactersAtIndex returns the range of the composed character sequence at a specified index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetRangeOfComposedCharactersAtIndex(_:_:)
func CFStringGetRangeOfComposedCharactersAtIndex(theString CFStringRef, theIndex int) CFRange {
	result, callErr := tryCFStringGetRangeOfComposedCharactersAtIndex(theString, theIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetSmallestEncoding func(theString CFStringRef) uint32
var _cFStringGetSmallestEncodingErr error

func tryCFStringGetSmallestEncoding(theString CFStringRef) (uint32, error) {
	if _cFStringGetSmallestEncoding == nil {
		return 0, symbolCallError("CFStringGetSmallestEncoding", "", _cFStringGetSmallestEncodingErr)
	}
	return _cFStringGetSmallestEncoding(theString), nil
}

// CFStringGetSmallestEncoding returns the smallest encoding on the current system for the character contents of a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSmallestEncoding(_:)
func CFStringGetSmallestEncoding(theString CFStringRef) uint32 {
	result, callErr := tryCFStringGetSmallestEncoding(theString)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetSystemEncoding func() uint32
var _cFStringGetSystemEncodingErr error

func tryCFStringGetSystemEncoding() (uint32, error) {
	if _cFStringGetSystemEncoding == nil {
		return 0, symbolCallError("CFStringGetSystemEncoding", "", _cFStringGetSystemEncodingErr)
	}
	return _cFStringGetSystemEncoding(), nil
}

// CFStringGetSystemEncoding returns the default encoding used by the operating system when it creates strings.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSystemEncoding()
func CFStringGetSystemEncoding() uint32 {
	result, callErr := tryCFStringGetSystemEncoding()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringGetTypeID func() uint
var _cFStringGetTypeIDErr error

func tryCFStringGetTypeID() (uint, error) {
	if _cFStringGetTypeID == nil {
		return 0, symbolCallError("CFStringGetTypeID", "", _cFStringGetTypeIDErr)
	}
	return _cFStringGetTypeID(), nil
}

// CFStringGetTypeID returns the type identifier for the CFString opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetTypeID()
func CFStringGetTypeID() uint {
	result, callErr := tryCFStringGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringHasPrefix func(theString CFStringRef, prefix CFStringRef) bool
var _cFStringHasPrefixErr error

func tryCFStringHasPrefix(theString CFStringRef, prefix CFStringRef) (bool, error) {
	if _cFStringHasPrefix == nil {
		return false, symbolCallError("CFStringHasPrefix", "", _cFStringHasPrefixErr)
	}
	return _cFStringHasPrefix(theString, prefix), nil
}

// CFStringHasPrefix determines if the character data of a string begin with a specified sequence of characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringHasPrefix(_:_:)
func CFStringHasPrefix(theString CFStringRef, prefix CFStringRef) bool {
	result, callErr := tryCFStringHasPrefix(theString, prefix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringHasSuffix func(theString CFStringRef, suffix CFStringRef) bool
var _cFStringHasSuffixErr error

func tryCFStringHasSuffix(theString CFStringRef, suffix CFStringRef) (bool, error) {
	if _cFStringHasSuffix == nil {
		return false, symbolCallError("CFStringHasSuffix", "", _cFStringHasSuffixErr)
	}
	return _cFStringHasSuffix(theString, suffix), nil
}

// CFStringHasSuffix determines if a string ends with a specified sequence of characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringHasSuffix(_:_:)
func CFStringHasSuffix(theString CFStringRef, suffix CFStringRef) bool {
	result, callErr := tryCFStringHasSuffix(theString, suffix)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringInsert func(str CFMutableStringRef, idx int, insertedStr CFStringRef)
var _cFStringInsertErr error

func tryCFStringInsert(str CFMutableStringRef, idx int, insertedStr CFStringRef) error {
	if _cFStringInsert == nil {
		return symbolCallError("CFStringInsert", "", _cFStringInsertErr)
	}
	_cFStringInsert(str, idx, insertedStr)
	return nil
}

// CFStringInsert inserts a string at a specified location in the character buffer of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringInsert(_:_:_:)
func CFStringInsert(str CFMutableStringRef, idx int, insertedStr CFStringRef) {
	if callErr := tryCFStringInsert(str, idx, insertedStr); callErr != nil {
		panic(callErr)
	}
}

var _cFStringIsEncodingAvailable func(encoding uint32) bool
var _cFStringIsEncodingAvailableErr error

func tryCFStringIsEncodingAvailable(encoding uint32) (bool, error) {
	if _cFStringIsEncodingAvailable == nil {
		return false, symbolCallError("CFStringIsEncodingAvailable", "", _cFStringIsEncodingAvailableErr)
	}
	return _cFStringIsEncodingAvailable(encoding), nil
}

// CFStringIsEncodingAvailable determines whether a given Core Foundation string encoding is available on the current system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringIsEncodingAvailable(_:)
func CFStringIsEncodingAvailable(encoding uint32) bool {
	result, callErr := tryCFStringIsEncodingAvailable(encoding)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringIsHyphenationAvailableForLocale func(locale CFLocaleRef) bool
var _cFStringIsHyphenationAvailableForLocaleErr error

func tryCFStringIsHyphenationAvailableForLocale(locale CFLocaleRef) (bool, error) {
	if _cFStringIsHyphenationAvailableForLocale == nil {
		return false, symbolCallError("CFStringIsHyphenationAvailableForLocale", "10.7", _cFStringIsHyphenationAvailableForLocaleErr)
	}
	return _cFStringIsHyphenationAvailableForLocale(locale), nil
}

// CFStringIsHyphenationAvailableForLocale returns a Boolean value that indicates whether hyphenation data is available.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringIsHyphenationAvailableForLocale(_:)
func CFStringIsHyphenationAvailableForLocale(locale CFLocaleRef) bool {
	result, callErr := tryCFStringIsHyphenationAvailableForLocale(locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringLowercase func(theString CFMutableStringRef, locale CFLocaleRef)
var _cFStringLowercaseErr error

func tryCFStringLowercase(theString CFMutableStringRef, locale CFLocaleRef) error {
	if _cFStringLowercase == nil {
		return symbolCallError("CFStringLowercase", "", _cFStringLowercaseErr)
	}
	_cFStringLowercase(theString, locale)
	return nil
}

// CFStringLowercase changes all uppercase alphabetical characters in a CFMutableString to lowercase.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringLowercase(_:_:)
func CFStringLowercase(theString CFMutableStringRef, locale CFLocaleRef) {
	if callErr := tryCFStringLowercase(theString, locale); callErr != nil {
		panic(callErr)
	}
}

var _cFStringNormalize func(theString CFMutableStringRef, theForm CFStringNormalizationForm)
var _cFStringNormalizeErr error

func tryCFStringNormalize(theString CFMutableStringRef, theForm CFStringNormalizationForm) error {
	if _cFStringNormalize == nil {
		return symbolCallError("CFStringNormalize", "", _cFStringNormalizeErr)
	}
	_cFStringNormalize(theString, theForm)
	return nil
}

// CFStringNormalize normalizes the string into the specified form as described in Unicode Technical Report #15.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalize(_:_:)
func CFStringNormalize(theString CFMutableStringRef, theForm CFStringNormalizationForm) {
	if callErr := tryCFStringNormalize(theString, theForm); callErr != nil {
		panic(callErr)
	}
}

var _cFStringPad func(theString CFMutableStringRef, padString CFStringRef, length int, indexIntoPad int)
var _cFStringPadErr error

func tryCFStringPad(theString CFMutableStringRef, padString CFStringRef, length int, indexIntoPad int) error {
	if _cFStringPad == nil {
		return symbolCallError("CFStringPad", "", _cFStringPadErr)
	}
	_cFStringPad(theString, padString, length, indexIntoPad)
	return nil
}

// CFStringPad enlarges a string, padding it with specified characters, or truncates the string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringPad(_:_:_:_:)
func CFStringPad(theString CFMutableStringRef, padString CFStringRef, length int, indexIntoPad int) {
	if callErr := tryCFStringPad(theString, padString, length, indexIntoPad); callErr != nil {
		panic(callErr)
	}
}

var _cFStringReplace func(theString CFMutableStringRef, range_ CFRange, replacement CFStringRef)
var _cFStringReplaceErr error

func tryCFStringReplace(theString CFMutableStringRef, range_ CFRange, replacement CFStringRef) error {
	if _cFStringReplace == nil {
		return symbolCallError("CFStringReplace", "", _cFStringReplaceErr)
	}
	_cFStringReplace(theString, range_, replacement)
	return nil
}

// CFStringReplace replaces part of the character contents of a CFMutableString object with another string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringReplace(_:_:_:)
func CFStringReplace(theString CFMutableStringRef, range_ CFRange, replacement CFStringRef) {
	if callErr := tryCFStringReplace(theString, range_, replacement); callErr != nil {
		panic(callErr)
	}
}

var _cFStringReplaceAll func(theString CFMutableStringRef, replacement CFStringRef)
var _cFStringReplaceAllErr error

func tryCFStringReplaceAll(theString CFMutableStringRef, replacement CFStringRef) error {
	if _cFStringReplaceAll == nil {
		return symbolCallError("CFStringReplaceAll", "", _cFStringReplaceAllErr)
	}
	_cFStringReplaceAll(theString, replacement)
	return nil
}

// CFStringReplaceAll replaces all characters of a CFMutableString object with other characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringReplaceAll(_:_:)
func CFStringReplaceAll(theString CFMutableStringRef, replacement CFStringRef) {
	if callErr := tryCFStringReplaceAll(theString, replacement); callErr != nil {
		panic(callErr)
	}
}

var _cFStringSetExternalCharactersNoCopy func(theString CFMutableStringRef, chars *uint16, length int, capacity int)
var _cFStringSetExternalCharactersNoCopyErr error

func tryCFStringSetExternalCharactersNoCopy(theString CFMutableStringRef, chars *uint16, length int, capacity int) error {
	if _cFStringSetExternalCharactersNoCopy == nil {
		return symbolCallError("CFStringSetExternalCharactersNoCopy", "", _cFStringSetExternalCharactersNoCopyErr)
	}
	_cFStringSetExternalCharactersNoCopy(theString, chars, length, capacity)
	return nil
}

// CFStringSetExternalCharactersNoCopy notifies a CFMutableString object that its external backing store of Unicode characters has changed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringSetExternalCharactersNoCopy(_:_:_:_:)
func CFStringSetExternalCharactersNoCopy(theString CFMutableStringRef, chars *uint16, length int, capacity int) {
	if callErr := tryCFStringSetExternalCharactersNoCopy(theString, chars, length, capacity); callErr != nil {
		panic(callErr)
	}
}

var _cFStringTokenizerAdvanceToNextToken func(tokenizer CFStringTokenizerRef) CFStringTokenizerTokenType
var _cFStringTokenizerAdvanceToNextTokenErr error

func tryCFStringTokenizerAdvanceToNextToken(tokenizer CFStringTokenizerRef) (CFStringTokenizerTokenType, error) {
	if _cFStringTokenizerAdvanceToNextToken == nil {
		return *new(CFStringTokenizerTokenType), symbolCallError("CFStringTokenizerAdvanceToNextToken", "10.5", _cFStringTokenizerAdvanceToNextTokenErr)
	}
	return _cFStringTokenizerAdvanceToNextToken(tokenizer), nil
}

// CFStringTokenizerAdvanceToNextToken advances the tokenizer to the next token and sets that as the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerAdvanceToNextToken(_:)
func CFStringTokenizerAdvanceToNextToken(tokenizer CFStringTokenizerRef) CFStringTokenizerTokenType {
	result, callErr := tryCFStringTokenizerAdvanceToNextToken(tokenizer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerCopyBestStringLanguage func(string_ CFStringRef, range_ CFRange) CFStringRef
var _cFStringTokenizerCopyBestStringLanguageErr error

func tryCFStringTokenizerCopyBestStringLanguage(string_ CFStringRef, range_ CFRange) (CFStringRef, error) {
	if _cFStringTokenizerCopyBestStringLanguage == nil {
		return 0, symbolCallError("CFStringTokenizerCopyBestStringLanguage", "10.5", _cFStringTokenizerCopyBestStringLanguageErr)
	}
	return _cFStringTokenizerCopyBestStringLanguage(string_, range_), nil
}

// CFStringTokenizerCopyBestStringLanguage guesses a language of a given string and returns the guess as a BCP 47 string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyBestStringLanguage(_:_:)
func CFStringTokenizerCopyBestStringLanguage(string_ CFStringRef, range_ CFRange) CFStringRef {
	result, callErr := tryCFStringTokenizerCopyBestStringLanguage(string_, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerCopyCurrentTokenAttribute func(tokenizer CFStringTokenizerRef, attribute uint64) CFTypeRef
var _cFStringTokenizerCopyCurrentTokenAttributeErr error

func tryCFStringTokenizerCopyCurrentTokenAttribute(tokenizer CFStringTokenizerRef, attribute uint64) (CFTypeRef, error) {
	if _cFStringTokenizerCopyCurrentTokenAttribute == nil {
		return 0, symbolCallError("CFStringTokenizerCopyCurrentTokenAttribute", "10.5", _cFStringTokenizerCopyCurrentTokenAttributeErr)
	}
	return _cFStringTokenizerCopyCurrentTokenAttribute(tokenizer, attribute), nil
}

// CFStringTokenizerCopyCurrentTokenAttribute returns a given attribute of the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyCurrentTokenAttribute(_:_:)
func CFStringTokenizerCopyCurrentTokenAttribute(tokenizer CFStringTokenizerRef, attribute uint64) CFTypeRef {
	result, callErr := tryCFStringTokenizerCopyCurrentTokenAttribute(tokenizer, attribute)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerCreate func(alloc CFAllocatorRef, string_ CFStringRef, range_ CFRange, options uint64, locale CFLocaleRef) CFStringTokenizerRef
var _cFStringTokenizerCreateErr error

func tryCFStringTokenizerCreate(alloc CFAllocatorRef, string_ CFStringRef, range_ CFRange, options uint64, locale CFLocaleRef) (CFStringTokenizerRef, error) {
	if _cFStringTokenizerCreate == nil {
		return 0, symbolCallError("CFStringTokenizerCreate", "10.5", _cFStringTokenizerCreateErr)
	}
	return _cFStringTokenizerCreate(alloc, string_, range_, options, locale), nil
}

// CFStringTokenizerCreate returns a tokenizer for a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCreate(_:_:_:_:_:)
func CFStringTokenizerCreate(alloc CFAllocatorRef, string_ CFStringRef, range_ CFRange, options uint64, locale CFLocaleRef) CFStringTokenizerRef {
	result, callErr := tryCFStringTokenizerCreate(alloc, string_, range_, options, locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerGetCurrentSubTokens func(tokenizer CFStringTokenizerRef, ranges *CFRange, maxRangeLength int, derivedSubTokens CFMutableArrayRef) int
var _cFStringTokenizerGetCurrentSubTokensErr error

func tryCFStringTokenizerGetCurrentSubTokens(tokenizer CFStringTokenizerRef, ranges *CFRange, maxRangeLength int, derivedSubTokens CFMutableArrayRef) (int, error) {
	if _cFStringTokenizerGetCurrentSubTokens == nil {
		return 0, symbolCallError("CFStringTokenizerGetCurrentSubTokens", "10.5", _cFStringTokenizerGetCurrentSubTokensErr)
	}
	return _cFStringTokenizerGetCurrentSubTokens(tokenizer, ranges, maxRangeLength, derivedSubTokens), nil
}

// CFStringTokenizerGetCurrentSubTokens retrieves the subtokens or derived subtokens contained in the compound token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentSubTokens(_:_:_:_:)
func CFStringTokenizerGetCurrentSubTokens(tokenizer CFStringTokenizerRef, ranges *CFRange, maxRangeLength int, derivedSubTokens CFMutableArrayRef) int {
	result, callErr := tryCFStringTokenizerGetCurrentSubTokens(tokenizer, ranges, maxRangeLength, derivedSubTokens)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerGetCurrentTokenRange func(tokenizer CFStringTokenizerRef) CFRange
var _cFStringTokenizerGetCurrentTokenRangeErr error

func tryCFStringTokenizerGetCurrentTokenRange(tokenizer CFStringTokenizerRef) (CFRange, error) {
	if _cFStringTokenizerGetCurrentTokenRange == nil {
		return CFRange{}, symbolCallError("CFStringTokenizerGetCurrentTokenRange", "10.5", _cFStringTokenizerGetCurrentTokenRangeErr)
	}
	return _cFStringTokenizerGetCurrentTokenRange(tokenizer), nil
}

// CFStringTokenizerGetCurrentTokenRange returns the range of the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentTokenRange(_:)
func CFStringTokenizerGetCurrentTokenRange(tokenizer CFStringTokenizerRef) CFRange {
	result, callErr := tryCFStringTokenizerGetCurrentTokenRange(tokenizer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerGetTypeID func() uint
var _cFStringTokenizerGetTypeIDErr error

func tryCFStringTokenizerGetTypeID() (uint, error) {
	if _cFStringTokenizerGetTypeID == nil {
		return 0, symbolCallError("CFStringTokenizerGetTypeID", "10.5", _cFStringTokenizerGetTypeIDErr)
	}
	return _cFStringTokenizerGetTypeID(), nil
}

// CFStringTokenizerGetTypeID returns the type ID for CFStringTokenizer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetTypeID()
func CFStringTokenizerGetTypeID() uint {
	result, callErr := tryCFStringTokenizerGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerGoToTokenAtIndex func(tokenizer CFStringTokenizerRef, index int) CFStringTokenizerTokenType
var _cFStringTokenizerGoToTokenAtIndexErr error

func tryCFStringTokenizerGoToTokenAtIndex(tokenizer CFStringTokenizerRef, index int) (CFStringTokenizerTokenType, error) {
	if _cFStringTokenizerGoToTokenAtIndex == nil {
		return *new(CFStringTokenizerTokenType), symbolCallError("CFStringTokenizerGoToTokenAtIndex", "10.5", _cFStringTokenizerGoToTokenAtIndexErr)
	}
	return _cFStringTokenizerGoToTokenAtIndex(tokenizer, index), nil
}

// CFStringTokenizerGoToTokenAtIndex finds a token that includes the character at a given index, and set it as the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGoToTokenAtIndex(_:_:)
func CFStringTokenizerGoToTokenAtIndex(tokenizer CFStringTokenizerRef, index int) CFStringTokenizerTokenType {
	result, callErr := tryCFStringTokenizerGoToTokenAtIndex(tokenizer, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTokenizerSetString func(tokenizer CFStringTokenizerRef, string_ CFStringRef, range_ CFRange)
var _cFStringTokenizerSetStringErr error

func tryCFStringTokenizerSetString(tokenizer CFStringTokenizerRef, string_ CFStringRef, range_ CFRange) error {
	if _cFStringTokenizerSetString == nil {
		return symbolCallError("CFStringTokenizerSetString", "10.5", _cFStringTokenizerSetStringErr)
	}
	_cFStringTokenizerSetString(tokenizer, string_, range_)
	return nil
}

// CFStringTokenizerSetString sets the string for a tokenizer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerSetString(_:_:_:)
func CFStringTokenizerSetString(tokenizer CFStringTokenizerRef, string_ CFStringRef, range_ CFRange) {
	if callErr := tryCFStringTokenizerSetString(tokenizer, string_, range_); callErr != nil {
		panic(callErr)
	}
}

var _cFStringTransform func(string_ CFMutableStringRef, range_ *CFRange, transform CFStringRef, reverse bool) bool
var _cFStringTransformErr error

func tryCFStringTransform(string_ CFMutableStringRef, range_ *CFRange, transform CFStringRef, reverse bool) (bool, error) {
	if _cFStringTransform == nil {
		return false, symbolCallError("CFStringTransform", "", _cFStringTransformErr)
	}
	return _cFStringTransform(string_, range_, transform, reverse), nil
}

// CFStringTransform perform in-place transliteration on a mutable string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTransform(_:_:_:_:)
func CFStringTransform(string_ CFMutableStringRef, range_ *CFRange, transform CFStringRef, reverse bool) bool {
	result, callErr := tryCFStringTransform(string_, range_, transform, reverse)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFStringTrim func(theString CFMutableStringRef, trimString CFStringRef)
var _cFStringTrimErr error

func tryCFStringTrim(theString CFMutableStringRef, trimString CFStringRef) error {
	if _cFStringTrim == nil {
		return symbolCallError("CFStringTrim", "", _cFStringTrimErr)
	}
	_cFStringTrim(theString, trimString)
	return nil
}

// CFStringTrim trims a specified substring from the beginning and end of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTrim(_:_:)
func CFStringTrim(theString CFMutableStringRef, trimString CFStringRef) {
	if callErr := tryCFStringTrim(theString, trimString); callErr != nil {
		panic(callErr)
	}
}

var _cFStringTrimWhitespace func(theString CFMutableStringRef)
var _cFStringTrimWhitespaceErr error

func tryCFStringTrimWhitespace(theString CFMutableStringRef) error {
	if _cFStringTrimWhitespace == nil {
		return symbolCallError("CFStringTrimWhitespace", "", _cFStringTrimWhitespaceErr)
	}
	_cFStringTrimWhitespace(theString)
	return nil
}

// CFStringTrimWhitespace trims whitespace from the beginning and end of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTrimWhitespace(_:)
func CFStringTrimWhitespace(theString CFMutableStringRef) {
	if callErr := tryCFStringTrimWhitespace(theString); callErr != nil {
		panic(callErr)
	}
}

var _cFStringUppercase func(theString CFMutableStringRef, locale CFLocaleRef)
var _cFStringUppercaseErr error

func tryCFStringUppercase(theString CFMutableStringRef, locale CFLocaleRef) error {
	if _cFStringUppercase == nil {
		return symbolCallError("CFStringUppercase", "", _cFStringUppercaseErr)
	}
	_cFStringUppercase(theString, locale)
	return nil
}

// CFStringUppercase changes all lowercase alphabetical characters in a CFMutableString object to uppercase.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringUppercase(_:_:)
func CFStringUppercase(theString CFMutableStringRef, locale CFLocaleRef) {
	if callErr := tryCFStringUppercase(theString, locale); callErr != nil {
		panic(callErr)
	}
}

var _cFTimeZoneCopyAbbreviation func(tz CFTimeZoneRef, at CFAbsoluteTime) CFStringRef
var _cFTimeZoneCopyAbbreviationErr error

func tryCFTimeZoneCopyAbbreviation(tz CFTimeZoneRef, at CFAbsoluteTime) (CFStringRef, error) {
	if _cFTimeZoneCopyAbbreviation == nil {
		return 0, symbolCallError("CFTimeZoneCopyAbbreviation", "", _cFTimeZoneCopyAbbreviationErr)
	}
	return _cFTimeZoneCopyAbbreviation(tz, at), nil
}

// CFTimeZoneCopyAbbreviation returns the abbreviation of a time zone at a specified date.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviation(_:_:)
func CFTimeZoneCopyAbbreviation(tz CFTimeZoneRef, at CFAbsoluteTime) CFStringRef {
	result, callErr := tryCFTimeZoneCopyAbbreviation(tz, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCopyAbbreviationDictionary func() CFDictionaryRef
var _cFTimeZoneCopyAbbreviationDictionaryErr error

func tryCFTimeZoneCopyAbbreviationDictionary() (CFDictionaryRef, error) {
	if _cFTimeZoneCopyAbbreviationDictionary == nil {
		return 0, symbolCallError("CFTimeZoneCopyAbbreviationDictionary", "", _cFTimeZoneCopyAbbreviationDictionaryErr)
	}
	return _cFTimeZoneCopyAbbreviationDictionary(), nil
}

// CFTimeZoneCopyAbbreviationDictionary returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviationDictionary()
func CFTimeZoneCopyAbbreviationDictionary() CFDictionaryRef {
	result, callErr := tryCFTimeZoneCopyAbbreviationDictionary()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCopyDefault func() CFTimeZoneRef
var _cFTimeZoneCopyDefaultErr error

func tryCFTimeZoneCopyDefault() (CFTimeZoneRef, error) {
	if _cFTimeZoneCopyDefault == nil {
		return 0, symbolCallError("CFTimeZoneCopyDefault", "", _cFTimeZoneCopyDefaultErr)
	}
	return _cFTimeZoneCopyDefault(), nil
}

// CFTimeZoneCopyDefault returns the default time zone set for your application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyDefault()
func CFTimeZoneCopyDefault() CFTimeZoneRef {
	result, callErr := tryCFTimeZoneCopyDefault()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCopyKnownNames func() CFArrayRef
var _cFTimeZoneCopyKnownNamesErr error

func tryCFTimeZoneCopyKnownNames() (CFArrayRef, error) {
	if _cFTimeZoneCopyKnownNames == nil {
		return 0, symbolCallError("CFTimeZoneCopyKnownNames", "", _cFTimeZoneCopyKnownNamesErr)
	}
	return _cFTimeZoneCopyKnownNames(), nil
}

// CFTimeZoneCopyKnownNames returns an array of strings containing the names of all the time zones known to the system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyKnownNames()
func CFTimeZoneCopyKnownNames() CFArrayRef {
	result, callErr := tryCFTimeZoneCopyKnownNames()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCopyLocalizedName func(tz CFTimeZoneRef, style CFTimeZoneNameStyle, locale CFLocaleRef) CFStringRef
var _cFTimeZoneCopyLocalizedNameErr error

func tryCFTimeZoneCopyLocalizedName(tz CFTimeZoneRef, style CFTimeZoneNameStyle, locale CFLocaleRef) (CFStringRef, error) {
	if _cFTimeZoneCopyLocalizedName == nil {
		return 0, symbolCallError("CFTimeZoneCopyLocalizedName", "10.5", _cFTimeZoneCopyLocalizedNameErr)
	}
	return _cFTimeZoneCopyLocalizedName(tz, style, locale), nil
}

// CFTimeZoneCopyLocalizedName returns the localized name of a given time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyLocalizedName(_:_:_:)
func CFTimeZoneCopyLocalizedName(tz CFTimeZoneRef, style CFTimeZoneNameStyle, locale CFLocaleRef) CFStringRef {
	result, callErr := tryCFTimeZoneCopyLocalizedName(tz, style, locale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCopySystem func() CFTimeZoneRef
var _cFTimeZoneCopySystemErr error

func tryCFTimeZoneCopySystem() (CFTimeZoneRef, error) {
	if _cFTimeZoneCopySystem == nil {
		return 0, symbolCallError("CFTimeZoneCopySystem", "", _cFTimeZoneCopySystemErr)
	}
	return _cFTimeZoneCopySystem(), nil
}

// CFTimeZoneCopySystem returns the time zone currently used by the system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopySystem()
func CFTimeZoneCopySystem() CFTimeZoneRef {
	result, callErr := tryCFTimeZoneCopySystem()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCreate func(allocator CFAllocatorRef, name CFStringRef, data CFDataRef) CFTimeZoneRef
var _cFTimeZoneCreateErr error

func tryCFTimeZoneCreate(allocator CFAllocatorRef, name CFStringRef, data CFDataRef) (CFTimeZoneRef, error) {
	if _cFTimeZoneCreate == nil {
		return 0, symbolCallError("CFTimeZoneCreate", "", _cFTimeZoneCreateErr)
	}
	return _cFTimeZoneCreate(allocator, name, data), nil
}

// CFTimeZoneCreate creates a time zone with a given name and data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreate(_:_:_:)
func CFTimeZoneCreate(allocator CFAllocatorRef, name CFStringRef, data CFDataRef) CFTimeZoneRef {
	result, callErr := tryCFTimeZoneCreate(allocator, name, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCreateWithName func(allocator CFAllocatorRef, name CFStringRef, tryAbbrev bool) CFTimeZoneRef
var _cFTimeZoneCreateWithNameErr error

func tryCFTimeZoneCreateWithName(allocator CFAllocatorRef, name CFStringRef, tryAbbrev bool) (CFTimeZoneRef, error) {
	if _cFTimeZoneCreateWithName == nil {
		return 0, symbolCallError("CFTimeZoneCreateWithName", "", _cFTimeZoneCreateWithNameErr)
	}
	return _cFTimeZoneCreateWithName(allocator, name, tryAbbrev), nil
}

// CFTimeZoneCreateWithName returns the time zone object identified by a given name or abbreviation.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithName(_:_:_:)
func CFTimeZoneCreateWithName(allocator CFAllocatorRef, name CFStringRef, tryAbbrev bool) CFTimeZoneRef {
	result, callErr := tryCFTimeZoneCreateWithName(allocator, name, tryAbbrev)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneCreateWithTimeIntervalFromGMT func(allocator CFAllocatorRef, ti float64) CFTimeZoneRef
var _cFTimeZoneCreateWithTimeIntervalFromGMTErr error

func tryCFTimeZoneCreateWithTimeIntervalFromGMT(allocator CFAllocatorRef, ti float64) (CFTimeZoneRef, error) {
	if _cFTimeZoneCreateWithTimeIntervalFromGMT == nil {
		return 0, symbolCallError("CFTimeZoneCreateWithTimeIntervalFromGMT", "", _cFTimeZoneCreateWithTimeIntervalFromGMTErr)
	}
	return _cFTimeZoneCreateWithTimeIntervalFromGMT(allocator, ti), nil
}

// CFTimeZoneCreateWithTimeIntervalFromGMT returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithTimeIntervalFromGMT(_:_:)
func CFTimeZoneCreateWithTimeIntervalFromGMT(allocator CFAllocatorRef, ti float64) CFTimeZoneRef {
	result, callErr := tryCFTimeZoneCreateWithTimeIntervalFromGMT(allocator, ti)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneGetData func(tz CFTimeZoneRef) CFDataRef
var _cFTimeZoneGetDataErr error

func tryCFTimeZoneGetData(tz CFTimeZoneRef) (CFDataRef, error) {
	if _cFTimeZoneGetData == nil {
		return 0, symbolCallError("CFTimeZoneGetData", "", _cFTimeZoneGetDataErr)
	}
	return _cFTimeZoneGetData(tz), nil
}

// CFTimeZoneGetData returns the data that stores the information used by a time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetData(_:)
func CFTimeZoneGetData(tz CFTimeZoneRef) CFDataRef {
	result, callErr := tryCFTimeZoneGetData(tz)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneGetDaylightSavingTimeOffset func(tz CFTimeZoneRef, at CFAbsoluteTime) float64
var _cFTimeZoneGetDaylightSavingTimeOffsetErr error

func tryCFTimeZoneGetDaylightSavingTimeOffset(tz CFTimeZoneRef, at CFAbsoluteTime) (float64, error) {
	if _cFTimeZoneGetDaylightSavingTimeOffset == nil {
		return 0.0, symbolCallError("CFTimeZoneGetDaylightSavingTimeOffset", "10.5", _cFTimeZoneGetDaylightSavingTimeOffsetErr)
	}
	return _cFTimeZoneGetDaylightSavingTimeOffset(tz, at), nil
}

// CFTimeZoneGetDaylightSavingTimeOffset returns the daylight saving time offset for a time zone at a given time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetDaylightSavingTimeOffset(_:_:)
func CFTimeZoneGetDaylightSavingTimeOffset(tz CFTimeZoneRef, at CFAbsoluteTime) float64 {
	result, callErr := tryCFTimeZoneGetDaylightSavingTimeOffset(tz, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneGetName func(tz CFTimeZoneRef) CFStringRef
var _cFTimeZoneGetNameErr error

func tryCFTimeZoneGetName(tz CFTimeZoneRef) (CFStringRef, error) {
	if _cFTimeZoneGetName == nil {
		return 0, symbolCallError("CFTimeZoneGetName", "", _cFTimeZoneGetNameErr)
	}
	return _cFTimeZoneGetName(tz), nil
}

// CFTimeZoneGetName returns the geopolitical region name that identifies a given time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetName(_:)
func CFTimeZoneGetName(tz CFTimeZoneRef) CFStringRef {
	result, callErr := tryCFTimeZoneGetName(tz)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneGetNextDaylightSavingTimeTransition func(tz CFTimeZoneRef, at CFAbsoluteTime) CFAbsoluteTime
var _cFTimeZoneGetNextDaylightSavingTimeTransitionErr error

func tryCFTimeZoneGetNextDaylightSavingTimeTransition(tz CFTimeZoneRef, at CFAbsoluteTime) (CFAbsoluteTime, error) {
	if _cFTimeZoneGetNextDaylightSavingTimeTransition == nil {
		return *new(CFAbsoluteTime), symbolCallError("CFTimeZoneGetNextDaylightSavingTimeTransition", "10.5", _cFTimeZoneGetNextDaylightSavingTimeTransitionErr)
	}
	return _cFTimeZoneGetNextDaylightSavingTimeTransition(tz, at), nil
}

// CFTimeZoneGetNextDaylightSavingTimeTransition returns the time in a given time zone of the next daylight saving time transition after a given time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetNextDaylightSavingTimeTransition(_:_:)
func CFTimeZoneGetNextDaylightSavingTimeTransition(tz CFTimeZoneRef, at CFAbsoluteTime) CFAbsoluteTime {
	result, callErr := tryCFTimeZoneGetNextDaylightSavingTimeTransition(tz, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneGetSecondsFromGMT func(tz CFTimeZoneRef, at CFAbsoluteTime) float64
var _cFTimeZoneGetSecondsFromGMTErr error

func tryCFTimeZoneGetSecondsFromGMT(tz CFTimeZoneRef, at CFAbsoluteTime) (float64, error) {
	if _cFTimeZoneGetSecondsFromGMT == nil {
		return 0.0, symbolCallError("CFTimeZoneGetSecondsFromGMT", "", _cFTimeZoneGetSecondsFromGMTErr)
	}
	return _cFTimeZoneGetSecondsFromGMT(tz, at), nil
}

// CFTimeZoneGetSecondsFromGMT returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetSecondsFromGMT(_:_:)
func CFTimeZoneGetSecondsFromGMT(tz CFTimeZoneRef, at CFAbsoluteTime) float64 {
	result, callErr := tryCFTimeZoneGetSecondsFromGMT(tz, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneGetTypeID func() uint
var _cFTimeZoneGetTypeIDErr error

func tryCFTimeZoneGetTypeID() (uint, error) {
	if _cFTimeZoneGetTypeID == nil {
		return 0, symbolCallError("CFTimeZoneGetTypeID", "", _cFTimeZoneGetTypeIDErr)
	}
	return _cFTimeZoneGetTypeID(), nil
}

// CFTimeZoneGetTypeID returns the type identifier for the CFTimeZone opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetTypeID()
func CFTimeZoneGetTypeID() uint {
	result, callErr := tryCFTimeZoneGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneIsDaylightSavingTime func(tz CFTimeZoneRef, at CFAbsoluteTime) bool
var _cFTimeZoneIsDaylightSavingTimeErr error

func tryCFTimeZoneIsDaylightSavingTime(tz CFTimeZoneRef, at CFAbsoluteTime) (bool, error) {
	if _cFTimeZoneIsDaylightSavingTime == nil {
		return false, symbolCallError("CFTimeZoneIsDaylightSavingTime", "", _cFTimeZoneIsDaylightSavingTimeErr)
	}
	return _cFTimeZoneIsDaylightSavingTime(tz, at), nil
}

// CFTimeZoneIsDaylightSavingTime returns whether or not a time zone is in daylight savings time at a specified date.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneIsDaylightSavingTime(_:_:)
func CFTimeZoneIsDaylightSavingTime(tz CFTimeZoneRef, at CFAbsoluteTime) bool {
	result, callErr := tryCFTimeZoneIsDaylightSavingTime(tz, at)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTimeZoneResetSystem func()
var _cFTimeZoneResetSystemErr error

func tryCFTimeZoneResetSystem() error {
	if _cFTimeZoneResetSystem == nil {
		return symbolCallError("CFTimeZoneResetSystem", "", _cFTimeZoneResetSystemErr)
	}
	_cFTimeZoneResetSystem()
	return nil
}

// CFTimeZoneResetSystem clears the previously determined system time zone, if any.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneResetSystem()
func CFTimeZoneResetSystem() {
	if callErr := tryCFTimeZoneResetSystem(); callErr != nil {
		panic(callErr)
	}
}

var _cFTimeZoneSetAbbreviationDictionary func(dict CFDictionaryRef)
var _cFTimeZoneSetAbbreviationDictionaryErr error

func tryCFTimeZoneSetAbbreviationDictionary(dict CFDictionaryRef) error {
	if _cFTimeZoneSetAbbreviationDictionary == nil {
		return symbolCallError("CFTimeZoneSetAbbreviationDictionary", "", _cFTimeZoneSetAbbreviationDictionaryErr)
	}
	_cFTimeZoneSetAbbreviationDictionary(dict)
	return nil
}

// CFTimeZoneSetAbbreviationDictionary sets the abbreviation dictionary to a given dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetAbbreviationDictionary(_:)
func CFTimeZoneSetAbbreviationDictionary(dict CFDictionaryRef) {
	if callErr := tryCFTimeZoneSetAbbreviationDictionary(dict); callErr != nil {
		panic(callErr)
	}
}

var _cFTimeZoneSetDefault func(tz CFTimeZoneRef)
var _cFTimeZoneSetDefaultErr error

func tryCFTimeZoneSetDefault(tz CFTimeZoneRef) error {
	if _cFTimeZoneSetDefault == nil {
		return symbolCallError("CFTimeZoneSetDefault", "", _cFTimeZoneSetDefaultErr)
	}
	_cFTimeZoneSetDefault(tz)
	return nil
}

// CFTimeZoneSetDefault sets the default time zone for your application the given time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetDefault(_:)
func CFTimeZoneSetDefault(tz CFTimeZoneRef) {
	if callErr := tryCFTimeZoneSetDefault(tz); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeAppendChild func(tree CFTreeRef, newChild CFTreeRef)
var _cFTreeAppendChildErr error

func tryCFTreeAppendChild(tree CFTreeRef, newChild CFTreeRef) error {
	if _cFTreeAppendChild == nil {
		return symbolCallError("CFTreeAppendChild", "", _cFTreeAppendChildErr)
	}
	_cFTreeAppendChild(tree, newChild)
	return nil
}

// CFTreeAppendChild adds a new child to a tree as the last in its list of children.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeAppendChild(_:_:)
func CFTreeAppendChild(tree CFTreeRef, newChild CFTreeRef) {
	if callErr := tryCFTreeAppendChild(tree, newChild); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeApplyFunctionToChildren func(tree CFTreeRef, applier CFTreeApplierFunction, context unsafe.Pointer)
var _cFTreeApplyFunctionToChildrenErr error

func tryCFTreeApplyFunctionToChildren(tree CFTreeRef, applier CFTreeApplierFunction, context unsafe.Pointer) error {
	if _cFTreeApplyFunctionToChildren == nil {
		return symbolCallError("CFTreeApplyFunctionToChildren", "", _cFTreeApplyFunctionToChildrenErr)
	}
	_cFTreeApplyFunctionToChildren(tree, applier, context)
	return nil
}

// CFTreeApplyFunctionToChildren calls a function once for each immediate child of a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeApplyFunctionToChildren(_:_:_:)
func CFTreeApplyFunctionToChildren(tree CFTreeRef, applier CFTreeApplierFunction, context unsafe.Pointer) {
	if callErr := tryCFTreeApplyFunctionToChildren(tree, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeCreate func(allocator CFAllocatorRef, context *CFTreeContext) CFTreeRef
var _cFTreeCreateErr error

func tryCFTreeCreate(allocator CFAllocatorRef, context *CFTreeContext) (CFTreeRef, error) {
	if _cFTreeCreate == nil {
		return 0, symbolCallError("CFTreeCreate", "", _cFTreeCreateErr)
	}
	return _cFTreeCreate(allocator, context), nil
}

// CFTreeCreate creates a new CFTree object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeCreate(_:_:)
func CFTreeCreate(allocator CFAllocatorRef, context *CFTreeContext) CFTreeRef {
	result, callErr := tryCFTreeCreate(allocator, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeFindRoot func(tree CFTreeRef) CFTreeRef
var _cFTreeFindRootErr error

func tryCFTreeFindRoot(tree CFTreeRef) (CFTreeRef, error) {
	if _cFTreeFindRoot == nil {
		return 0, symbolCallError("CFTreeFindRoot", "", _cFTreeFindRootErr)
	}
	return _cFTreeFindRoot(tree), nil
}

// CFTreeFindRoot returns the root tree of a given tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeFindRoot(_:)
func CFTreeFindRoot(tree CFTreeRef) CFTreeRef {
	result, callErr := tryCFTreeFindRoot(tree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeGetChildAtIndex func(tree CFTreeRef, idx int) CFTreeRef
var _cFTreeGetChildAtIndexErr error

func tryCFTreeGetChildAtIndex(tree CFTreeRef, idx int) (CFTreeRef, error) {
	if _cFTreeGetChildAtIndex == nil {
		return 0, symbolCallError("CFTreeGetChildAtIndex", "", _cFTreeGetChildAtIndexErr)
	}
	return _cFTreeGetChildAtIndex(tree, idx), nil
}

// CFTreeGetChildAtIndex returns the child of a tree at the specified index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildAtIndex(_:_:)
func CFTreeGetChildAtIndex(tree CFTreeRef, idx int) CFTreeRef {
	result, callErr := tryCFTreeGetChildAtIndex(tree, idx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeGetChildCount func(tree CFTreeRef) int
var _cFTreeGetChildCountErr error

func tryCFTreeGetChildCount(tree CFTreeRef) (int, error) {
	if _cFTreeGetChildCount == nil {
		return 0, symbolCallError("CFTreeGetChildCount", "", _cFTreeGetChildCountErr)
	}
	return _cFTreeGetChildCount(tree), nil
}

// CFTreeGetChildCount returns the number of children in a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildCount(_:)
func CFTreeGetChildCount(tree CFTreeRef) int {
	result, callErr := tryCFTreeGetChildCount(tree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeGetChildren func(tree CFTreeRef, children *CFTreeRef)
var _cFTreeGetChildrenErr error

func tryCFTreeGetChildren(tree CFTreeRef, children *CFTreeRef) error {
	if _cFTreeGetChildren == nil {
		return symbolCallError("CFTreeGetChildren", "", _cFTreeGetChildrenErr)
	}
	_cFTreeGetChildren(tree, children)
	return nil
}

// CFTreeGetChildren fills a buffer with children from the tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildren(_:_:)
func CFTreeGetChildren(tree CFTreeRef, children *CFTreeRef) {
	if callErr := tryCFTreeGetChildren(tree, children); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeGetContext func(tree CFTreeRef, context *CFTreeContext)
var _cFTreeGetContextErr error

func tryCFTreeGetContext(tree CFTreeRef, context *CFTreeContext) error {
	if _cFTreeGetContext == nil {
		return symbolCallError("CFTreeGetContext", "", _cFTreeGetContextErr)
	}
	_cFTreeGetContext(tree, context)
	return nil
}

// CFTreeGetContext returns the context of the specified tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetContext(_:_:)
func CFTreeGetContext(tree CFTreeRef, context *CFTreeContext) {
	if callErr := tryCFTreeGetContext(tree, context); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeGetFirstChild func(tree CFTreeRef) CFTreeRef
var _cFTreeGetFirstChildErr error

func tryCFTreeGetFirstChild(tree CFTreeRef) (CFTreeRef, error) {
	if _cFTreeGetFirstChild == nil {
		return 0, symbolCallError("CFTreeGetFirstChild", "", _cFTreeGetFirstChildErr)
	}
	return _cFTreeGetFirstChild(tree), nil
}

// CFTreeGetFirstChild returns the first child of a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetFirstChild(_:)
func CFTreeGetFirstChild(tree CFTreeRef) CFTreeRef {
	result, callErr := tryCFTreeGetFirstChild(tree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeGetNextSibling func(tree CFTreeRef) CFTreeRef
var _cFTreeGetNextSiblingErr error

func tryCFTreeGetNextSibling(tree CFTreeRef) (CFTreeRef, error) {
	if _cFTreeGetNextSibling == nil {
		return 0, symbolCallError("CFTreeGetNextSibling", "", _cFTreeGetNextSiblingErr)
	}
	return _cFTreeGetNextSibling(tree), nil
}

// CFTreeGetNextSibling returns the next sibling, adjacent to a given tree, in the parent’s children list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetNextSibling(_:)
func CFTreeGetNextSibling(tree CFTreeRef) CFTreeRef {
	result, callErr := tryCFTreeGetNextSibling(tree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeGetParent func(tree CFTreeRef) CFTreeRef
var _cFTreeGetParentErr error

func tryCFTreeGetParent(tree CFTreeRef) (CFTreeRef, error) {
	if _cFTreeGetParent == nil {
		return 0, symbolCallError("CFTreeGetParent", "", _cFTreeGetParentErr)
	}
	return _cFTreeGetParent(tree), nil
}

// CFTreeGetParent returns the parent of a given tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetParent(_:)
func CFTreeGetParent(tree CFTreeRef) CFTreeRef {
	result, callErr := tryCFTreeGetParent(tree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeGetTypeID func() uint
var _cFTreeGetTypeIDErr error

func tryCFTreeGetTypeID() (uint, error) {
	if _cFTreeGetTypeID == nil {
		return 0, symbolCallError("CFTreeGetTypeID", "", _cFTreeGetTypeIDErr)
	}
	return _cFTreeGetTypeID(), nil
}

// CFTreeGetTypeID returns the type identifier of the CFTree opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetTypeID()
func CFTreeGetTypeID() uint {
	result, callErr := tryCFTreeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFTreeInsertSibling func(tree CFTreeRef, newSibling CFTreeRef)
var _cFTreeInsertSiblingErr error

func tryCFTreeInsertSibling(tree CFTreeRef, newSibling CFTreeRef) error {
	if _cFTreeInsertSibling == nil {
		return symbolCallError("CFTreeInsertSibling", "", _cFTreeInsertSiblingErr)
	}
	_cFTreeInsertSibling(tree, newSibling)
	return nil
}

// CFTreeInsertSibling inserts a new sibling after a given tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeInsertSibling(_:_:)
func CFTreeInsertSibling(tree CFTreeRef, newSibling CFTreeRef) {
	if callErr := tryCFTreeInsertSibling(tree, newSibling); callErr != nil {
		panic(callErr)
	}
}

var _cFTreePrependChild func(tree CFTreeRef, newChild CFTreeRef)
var _cFTreePrependChildErr error

func tryCFTreePrependChild(tree CFTreeRef, newChild CFTreeRef) error {
	if _cFTreePrependChild == nil {
		return symbolCallError("CFTreePrependChild", "", _cFTreePrependChildErr)
	}
	_cFTreePrependChild(tree, newChild)
	return nil
}

// CFTreePrependChild adds a new child to the specified tree as the first in its list of children.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreePrependChild(_:_:)
func CFTreePrependChild(tree CFTreeRef, newChild CFTreeRef) {
	if callErr := tryCFTreePrependChild(tree, newChild); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeRemove func(tree CFTreeRef)
var _cFTreeRemoveErr error

func tryCFTreeRemove(tree CFTreeRef) error {
	if _cFTreeRemove == nil {
		return symbolCallError("CFTreeRemove", "", _cFTreeRemoveErr)
	}
	_cFTreeRemove(tree)
	return nil
}

// CFTreeRemove removes a tree from its parent.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemove(_:)
func CFTreeRemove(tree CFTreeRef) {
	if callErr := tryCFTreeRemove(tree); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeRemoveAllChildren func(tree CFTreeRef)
var _cFTreeRemoveAllChildrenErr error

func tryCFTreeRemoveAllChildren(tree CFTreeRef) error {
	if _cFTreeRemoveAllChildren == nil {
		return symbolCallError("CFTreeRemoveAllChildren", "", _cFTreeRemoveAllChildrenErr)
	}
	_cFTreeRemoveAllChildren(tree)
	return nil
}

// CFTreeRemoveAllChildren removes all the children of a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemoveAllChildren(_:)
func CFTreeRemoveAllChildren(tree CFTreeRef) {
	if callErr := tryCFTreeRemoveAllChildren(tree); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeSetContext func(tree CFTreeRef, context *CFTreeContext)
var _cFTreeSetContextErr error

func tryCFTreeSetContext(tree CFTreeRef, context *CFTreeContext) error {
	if _cFTreeSetContext == nil {
		return symbolCallError("CFTreeSetContext", "", _cFTreeSetContextErr)
	}
	_cFTreeSetContext(tree, context)
	return nil
}

// CFTreeSetContext replaces the context of a tree by releasing the old information pointer and retaining the new one.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeSetContext(_:_:)
func CFTreeSetContext(tree CFTreeRef, context *CFTreeContext) {
	if callErr := tryCFTreeSetContext(tree, context); callErr != nil {
		panic(callErr)
	}
}

var _cFTreeSortChildren func(tree CFTreeRef, comparator CFComparatorFunction, context unsafe.Pointer)
var _cFTreeSortChildrenErr error

func tryCFTreeSortChildren(tree CFTreeRef, comparator CFComparatorFunction, context unsafe.Pointer) error {
	if _cFTreeSortChildren == nil {
		return symbolCallError("CFTreeSortChildren", "", _cFTreeSortChildrenErr)
	}
	_cFTreeSortChildren(tree, comparator, context)
	return nil
}

// CFTreeSortChildren sorts the immediate children of a tree using a specified comparator function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeSortChildren(_:_:_:)
func CFTreeSortChildren(tree CFTreeRef, comparator CFComparatorFunction, context unsafe.Pointer) {
	if callErr := tryCFTreeSortChildren(tree, comparator, context); callErr != nil {
		panic(callErr)
	}
}

var _cFURLCanBeDecomposed func(anURL CFURLRef) bool
var _cFURLCanBeDecomposedErr error

func tryCFURLCanBeDecomposed(anURL CFURLRef) (bool, error) {
	if _cFURLCanBeDecomposed == nil {
		return false, symbolCallError("CFURLCanBeDecomposed", "", _cFURLCanBeDecomposedErr)
	}
	return _cFURLCanBeDecomposed(anURL), nil
}

// CFURLCanBeDecomposed determines if the given URL conforms to RFC 1808 and therefore can be decomposed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCanBeDecomposed(_:)
func CFURLCanBeDecomposed(anURL CFURLRef) bool {
	result, callErr := tryCFURLCanBeDecomposed(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLClearResourcePropertyCache func(url CFURLRef)
var _cFURLClearResourcePropertyCacheErr error

func tryCFURLClearResourcePropertyCache(url CFURLRef) error {
	if _cFURLClearResourcePropertyCache == nil {
		return symbolCallError("CFURLClearResourcePropertyCache", "10.6", _cFURLClearResourcePropertyCacheErr)
	}
	_cFURLClearResourcePropertyCache(url)
	return nil
}

// CFURLClearResourcePropertyCache removes all cached resource values and temporary resource values from the URL object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCache(_:)
func CFURLClearResourcePropertyCache(url CFURLRef) {
	if callErr := tryCFURLClearResourcePropertyCache(url); callErr != nil {
		panic(callErr)
	}
}

var _cFURLClearResourcePropertyCacheForKey func(url CFURLRef, key CFStringRef)
var _cFURLClearResourcePropertyCacheForKeyErr error

func tryCFURLClearResourcePropertyCacheForKey(url CFURLRef, key CFStringRef) error {
	if _cFURLClearResourcePropertyCacheForKey == nil {
		return symbolCallError("CFURLClearResourcePropertyCacheForKey", "10.6", _cFURLClearResourcePropertyCacheForKeyErr)
	}
	_cFURLClearResourcePropertyCacheForKey(url, key)
	return nil
}

// CFURLClearResourcePropertyCacheForKey removes the cached resource value identified by a given key from the URL object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCacheForKey(_:_:)
func CFURLClearResourcePropertyCacheForKey(url CFURLRef, key CFStringRef) {
	if callErr := tryCFURLClearResourcePropertyCacheForKey(url, key); callErr != nil {
		panic(callErr)
	}
}

var _cFURLCopyAbsoluteURL func(relativeURL CFURLRef) CFURLRef
var _cFURLCopyAbsoluteURLErr error

func tryCFURLCopyAbsoluteURL(relativeURL CFURLRef) (CFURLRef, error) {
	if _cFURLCopyAbsoluteURL == nil {
		return 0, symbolCallError("CFURLCopyAbsoluteURL", "", _cFURLCopyAbsoluteURLErr)
	}
	return _cFURLCopyAbsoluteURL(relativeURL), nil
}

// CFURLCopyAbsoluteURL creates a new [CFURL] object by resolving the relative portion of a URL against its base.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyAbsoluteURL(_:)
func CFURLCopyAbsoluteURL(relativeURL CFURLRef) CFURLRef {
	result, callErr := tryCFURLCopyAbsoluteURL(relativeURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyFileSystemPath func(anURL CFURLRef, pathStyle CFURLPathStyle) CFStringRef
var _cFURLCopyFileSystemPathErr error

func tryCFURLCopyFileSystemPath(anURL CFURLRef, pathStyle CFURLPathStyle) (CFStringRef, error) {
	if _cFURLCopyFileSystemPath == nil {
		return 0, symbolCallError("CFURLCopyFileSystemPath", "", _cFURLCopyFileSystemPathErr)
	}
	return _cFURLCopyFileSystemPath(anURL, pathStyle), nil
}

// CFURLCopyFileSystemPath returns the path portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFileSystemPath(_:_:)
func CFURLCopyFileSystemPath(anURL CFURLRef, pathStyle CFURLPathStyle) CFStringRef {
	result, callErr := tryCFURLCopyFileSystemPath(anURL, pathStyle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyFragment func(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef
var _cFURLCopyFragmentErr error

func tryCFURLCopyFragment(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) (CFStringRef, error) {
	if _cFURLCopyFragment == nil {
		return 0, symbolCallError("CFURLCopyFragment", "", _cFURLCopyFragmentErr)
	}
	return _cFURLCopyFragment(anURL, charactersToLeaveEscaped), nil
}

// CFURLCopyFragment returns the fragment from a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFragment(_:_:)
func CFURLCopyFragment(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef {
	result, callErr := tryCFURLCopyFragment(anURL, charactersToLeaveEscaped)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyHostName func(anURL CFURLRef) CFStringRef
var _cFURLCopyHostNameErr error

func tryCFURLCopyHostName(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyHostName == nil {
		return 0, symbolCallError("CFURLCopyHostName", "", _cFURLCopyHostNameErr)
	}
	return _cFURLCopyHostName(anURL), nil
}

// CFURLCopyHostName returns the host name of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyHostName(_:)
func CFURLCopyHostName(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyHostName(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyLastPathComponent func(url CFURLRef) CFStringRef
var _cFURLCopyLastPathComponentErr error

func tryCFURLCopyLastPathComponent(url CFURLRef) (CFStringRef, error) {
	if _cFURLCopyLastPathComponent == nil {
		return 0, symbolCallError("CFURLCopyLastPathComponent", "", _cFURLCopyLastPathComponentErr)
	}
	return _cFURLCopyLastPathComponent(url), nil
}

// CFURLCopyLastPathComponent returns the last path component of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyLastPathComponent(_:)
func CFURLCopyLastPathComponent(url CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyLastPathComponent(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyNetLocation func(anURL CFURLRef) CFStringRef
var _cFURLCopyNetLocationErr error

func tryCFURLCopyNetLocation(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyNetLocation == nil {
		return 0, symbolCallError("CFURLCopyNetLocation", "", _cFURLCopyNetLocationErr)
	}
	return _cFURLCopyNetLocation(anURL), nil
}

// CFURLCopyNetLocation returns the net location portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyNetLocation(_:)
func CFURLCopyNetLocation(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyNetLocation(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyPassword func(anURL CFURLRef) CFStringRef
var _cFURLCopyPasswordErr error

func tryCFURLCopyPassword(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyPassword == nil {
		return 0, symbolCallError("CFURLCopyPassword", "", _cFURLCopyPasswordErr)
	}
	return _cFURLCopyPassword(anURL), nil
}

// CFURLCopyPassword returns the password of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPassword(_:)
func CFURLCopyPassword(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyPassword(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyPath func(anURL CFURLRef) CFStringRef
var _cFURLCopyPathErr error

func tryCFURLCopyPath(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyPath == nil {
		return 0, symbolCallError("CFURLCopyPath", "", _cFURLCopyPathErr)
	}
	return _cFURLCopyPath(anURL), nil
}

// CFURLCopyPath returns the path portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPath(_:)
func CFURLCopyPath(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyPath(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyPathExtension func(url CFURLRef) CFStringRef
var _cFURLCopyPathExtensionErr error

func tryCFURLCopyPathExtension(url CFURLRef) (CFStringRef, error) {
	if _cFURLCopyPathExtension == nil {
		return 0, symbolCallError("CFURLCopyPathExtension", "", _cFURLCopyPathExtensionErr)
	}
	return _cFURLCopyPathExtension(url), nil
}

// CFURLCopyPathExtension returns the path extension of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPathExtension(_:)
func CFURLCopyPathExtension(url CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyPathExtension(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyQueryString func(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef
var _cFURLCopyQueryStringErr error

func tryCFURLCopyQueryString(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) (CFStringRef, error) {
	if _cFURLCopyQueryString == nil {
		return 0, symbolCallError("CFURLCopyQueryString", "", _cFURLCopyQueryStringErr)
	}
	return _cFURLCopyQueryString(anURL, charactersToLeaveEscaped), nil
}

// CFURLCopyQueryString returns the query string of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyQueryString(_:_:)
func CFURLCopyQueryString(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef {
	result, callErr := tryCFURLCopyQueryString(anURL, charactersToLeaveEscaped)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyResourcePropertiesForKeys func(url CFURLRef, keys CFArrayRef, err *CFErrorRef) CFDictionaryRef
var _cFURLCopyResourcePropertiesForKeysErr error

func tryCFURLCopyResourcePropertiesForKeys(url CFURLRef, keys CFArrayRef, err *CFErrorRef) (CFDictionaryRef, error) {
	if _cFURLCopyResourcePropertiesForKeys == nil {
		return 0, symbolCallError("CFURLCopyResourcePropertiesForKeys", "10.6", _cFURLCopyResourcePropertiesForKeysErr)
	}
	return _cFURLCopyResourcePropertiesForKeys(url, keys, err), nil
}

// CFURLCopyResourcePropertiesForKeys returns the resource values for the properties identified by specified array of keys.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertiesForKeys(_:_:_:)
func CFURLCopyResourcePropertiesForKeys(url CFURLRef, keys CFArrayRef, err *CFErrorRef) CFDictionaryRef {
	result, callErr := tryCFURLCopyResourcePropertiesForKeys(url, keys, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyResourcePropertyForKey func(url CFURLRef, key CFStringRef, propertyValueTypeRefPtr unsafe.Pointer, err *CFErrorRef) bool
var _cFURLCopyResourcePropertyForKeyErr error

func tryCFURLCopyResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValueTypeRefPtr unsafe.Pointer, err *CFErrorRef) (bool, error) {
	if _cFURLCopyResourcePropertyForKey == nil {
		return false, symbolCallError("CFURLCopyResourcePropertyForKey", "10.6", _cFURLCopyResourcePropertyForKeyErr)
	}
	return _cFURLCopyResourcePropertyForKey(url, key, propertyValueTypeRefPtr, err), nil
}

// CFURLCopyResourcePropertyForKey returns the value of a given resource property of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertyForKey(_:_:_:_:)
func CFURLCopyResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValueTypeRefPtr unsafe.Pointer, err *CFErrorRef) bool {
	result, callErr := tryCFURLCopyResourcePropertyForKey(url, key, propertyValueTypeRefPtr, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyResourceSpecifier func(anURL CFURLRef) CFStringRef
var _cFURLCopyResourceSpecifierErr error

func tryCFURLCopyResourceSpecifier(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyResourceSpecifier == nil {
		return 0, symbolCallError("CFURLCopyResourceSpecifier", "", _cFURLCopyResourceSpecifierErr)
	}
	return _cFURLCopyResourceSpecifier(anURL), nil
}

// CFURLCopyResourceSpecifier returns any additional resource specifiers after the path.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourceSpecifier(_:)
func CFURLCopyResourceSpecifier(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyResourceSpecifier(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyScheme func(anURL CFURLRef) CFStringRef
var _cFURLCopySchemeErr error

func tryCFURLCopyScheme(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyScheme == nil {
		return 0, symbolCallError("CFURLCopyScheme", "", _cFURLCopySchemeErr)
	}
	return _cFURLCopyScheme(anURL), nil
}

// CFURLCopyScheme returns the scheme portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyScheme(_:)
func CFURLCopyScheme(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyScheme(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyStrictPath func(anURL CFURLRef, isAbsolute *bool) CFStringRef
var _cFURLCopyStrictPathErr error

func tryCFURLCopyStrictPath(anURL CFURLRef, isAbsolute *bool) (CFStringRef, error) {
	if _cFURLCopyStrictPath == nil {
		return 0, symbolCallError("CFURLCopyStrictPath", "", _cFURLCopyStrictPathErr)
	}
	return _cFURLCopyStrictPath(anURL, isAbsolute), nil
}

// CFURLCopyStrictPath returns the path portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyStrictPath(_:_:)
func CFURLCopyStrictPath(anURL CFURLRef, isAbsolute *bool) CFStringRef {
	result, callErr := tryCFURLCopyStrictPath(anURL, isAbsolute)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCopyUserName func(anURL CFURLRef) CFStringRef
var _cFURLCopyUserNameErr error

func tryCFURLCopyUserName(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLCopyUserName == nil {
		return 0, symbolCallError("CFURLCopyUserName", "", _cFURLCopyUserNameErr)
	}
	return _cFURLCopyUserName(anURL), nil
}

// CFURLCopyUserName returns the user name from a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyUserName(_:)
func CFURLCopyUserName(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLCopyUserName(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateAbsoluteURLWithBytes func(alloc CFAllocatorRef, relativeURLBytes *uint8, length int, encoding uint32, baseURL CFURLRef, useCompatibilityMode bool) CFURLRef
var _cFURLCreateAbsoluteURLWithBytesErr error

func tryCFURLCreateAbsoluteURLWithBytes(alloc CFAllocatorRef, relativeURLBytes *uint8, length int, encoding uint32, baseURL CFURLRef, useCompatibilityMode bool) (CFURLRef, error) {
	if _cFURLCreateAbsoluteURLWithBytes == nil {
		return 0, symbolCallError("CFURLCreateAbsoluteURLWithBytes", "", _cFURLCreateAbsoluteURLWithBytesErr)
	}
	return _cFURLCreateAbsoluteURLWithBytes(alloc, relativeURLBytes, length, encoding, baseURL, useCompatibilityMode), nil
}

// CFURLCreateAbsoluteURLWithBytes creates a new [CFURL] object by resolving the relative portion of a URL, specified as bytes, against its given base URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateAbsoluteURLWithBytes(_:_:_:_:_:_:)
func CFURLCreateAbsoluteURLWithBytes(alloc CFAllocatorRef, relativeURLBytes *uint8, length int, encoding uint32, baseURL CFURLRef, useCompatibilityMode bool) CFURLRef {
	result, callErr := tryCFURLCreateAbsoluteURLWithBytes(alloc, relativeURLBytes, length, encoding, baseURL, useCompatibilityMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateBookmarkData func(allocator CFAllocatorRef, url CFURLRef, options CFURLBookmarkCreationOptions, resourcePropertiesToInclude CFArrayRef, relativeToURL CFURLRef, err *CFErrorRef) CFDataRef
var _cFURLCreateBookmarkDataErr error

func tryCFURLCreateBookmarkData(allocator CFAllocatorRef, url CFURLRef, options CFURLBookmarkCreationOptions, resourcePropertiesToInclude CFArrayRef, relativeToURL CFURLRef, err *CFErrorRef) (CFDataRef, error) {
	if _cFURLCreateBookmarkData == nil {
		return 0, symbolCallError("CFURLCreateBookmarkData", "10.6", _cFURLCreateBookmarkDataErr)
	}
	return _cFURLCreateBookmarkData(allocator, url, options, resourcePropertiesToInclude, relativeToURL, err), nil
}

// CFURLCreateBookmarkData returns bookmark data for a URL, created with specified options and resource values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkData(_:_:_:_:_:_:)
func CFURLCreateBookmarkData(allocator CFAllocatorRef, url CFURLRef, options CFURLBookmarkCreationOptions, resourcePropertiesToInclude CFArrayRef, relativeToURL CFURLRef, err *CFErrorRef) CFDataRef {
	result, callErr := tryCFURLCreateBookmarkData(allocator, url, options, resourcePropertiesToInclude, relativeToURL, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateBookmarkDataFromFile func(allocator CFAllocatorRef, fileURL CFURLRef, errorRef *CFErrorRef) CFDataRef
var _cFURLCreateBookmarkDataFromFileErr error

func tryCFURLCreateBookmarkDataFromFile(allocator CFAllocatorRef, fileURL CFURLRef, errorRef *CFErrorRef) (CFDataRef, error) {
	if _cFURLCreateBookmarkDataFromFile == nil {
		return 0, symbolCallError("CFURLCreateBookmarkDataFromFile", "10.6", _cFURLCreateBookmarkDataFromFileErr)
	}
	return _cFURLCreateBookmarkDataFromFile(allocator, fileURL, errorRef), nil
}

// CFURLCreateBookmarkDataFromFile initializes and returns bookmark data derived from a file pointed to by a specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromFile(_:_:_:)
func CFURLCreateBookmarkDataFromFile(allocator CFAllocatorRef, fileURL CFURLRef, errorRef *CFErrorRef) CFDataRef {
	result, callErr := tryCFURLCreateBookmarkDataFromFile(allocator, fileURL, errorRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateByResolvingBookmarkData func(allocator CFAllocatorRef, bookmark CFDataRef, options CFURLBookmarkResolutionOptions, relativeToURL CFURLRef, resourcePropertiesToInclude CFArrayRef, isStale *bool, err *CFErrorRef) CFURLRef
var _cFURLCreateByResolvingBookmarkDataErr error

func tryCFURLCreateByResolvingBookmarkData(allocator CFAllocatorRef, bookmark CFDataRef, options CFURLBookmarkResolutionOptions, relativeToURL CFURLRef, resourcePropertiesToInclude CFArrayRef, isStale *bool, err *CFErrorRef) (CFURLRef, error) {
	if _cFURLCreateByResolvingBookmarkData == nil {
		return 0, symbolCallError("CFURLCreateByResolvingBookmarkData", "10.6", _cFURLCreateByResolvingBookmarkDataErr)
	}
	return _cFURLCreateByResolvingBookmarkData(allocator, bookmark, options, relativeToURL, resourcePropertiesToInclude, isStale, err), nil
}

// CFURLCreateByResolvingBookmarkData returns a new URL made by resolving bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateByResolvingBookmarkData(_:_:_:_:_:_:_:)
func CFURLCreateByResolvingBookmarkData(allocator CFAllocatorRef, bookmark CFDataRef, options CFURLBookmarkResolutionOptions, relativeToURL CFURLRef, resourcePropertiesToInclude CFArrayRef, isStale *bool, err *CFErrorRef) CFURLRef {
	result, callErr := tryCFURLCreateByResolvingBookmarkData(allocator, bookmark, options, relativeToURL, resourcePropertiesToInclude, isStale, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateCopyAppendingPathComponent func(allocator CFAllocatorRef, url CFURLRef, pathComponent CFStringRef, isDirectory bool) CFURLRef
var _cFURLCreateCopyAppendingPathComponentErr error

func tryCFURLCreateCopyAppendingPathComponent(allocator CFAllocatorRef, url CFURLRef, pathComponent CFStringRef, isDirectory bool) (CFURLRef, error) {
	if _cFURLCreateCopyAppendingPathComponent == nil {
		return 0, symbolCallError("CFURLCreateCopyAppendingPathComponent", "", _cFURLCreateCopyAppendingPathComponentErr)
	}
	return _cFURLCreateCopyAppendingPathComponent(allocator, url, pathComponent, isDirectory), nil
}

// CFURLCreateCopyAppendingPathComponent creates a copy of a given URL and appends a path component.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathComponent(_:_:_:_:)
func CFURLCreateCopyAppendingPathComponent(allocator CFAllocatorRef, url CFURLRef, pathComponent CFStringRef, isDirectory bool) CFURLRef {
	result, callErr := tryCFURLCreateCopyAppendingPathComponent(allocator, url, pathComponent, isDirectory)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateCopyAppendingPathExtension func(allocator CFAllocatorRef, url CFURLRef, extension CFStringRef) CFURLRef
var _cFURLCreateCopyAppendingPathExtensionErr error

func tryCFURLCreateCopyAppendingPathExtension(allocator CFAllocatorRef, url CFURLRef, extension CFStringRef) (CFURLRef, error) {
	if _cFURLCreateCopyAppendingPathExtension == nil {
		return 0, symbolCallError("CFURLCreateCopyAppendingPathExtension", "", _cFURLCreateCopyAppendingPathExtensionErr)
	}
	return _cFURLCreateCopyAppendingPathExtension(allocator, url, extension), nil
}

// CFURLCreateCopyAppendingPathExtension creates a copy of a given URL and appends a path extension.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathExtension(_:_:_:)
func CFURLCreateCopyAppendingPathExtension(allocator CFAllocatorRef, url CFURLRef, extension CFStringRef) CFURLRef {
	result, callErr := tryCFURLCreateCopyAppendingPathExtension(allocator, url, extension)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateCopyDeletingLastPathComponent func(allocator CFAllocatorRef, url CFURLRef) CFURLRef
var _cFURLCreateCopyDeletingLastPathComponentErr error

func tryCFURLCreateCopyDeletingLastPathComponent(allocator CFAllocatorRef, url CFURLRef) (CFURLRef, error) {
	if _cFURLCreateCopyDeletingLastPathComponent == nil {
		return 0, symbolCallError("CFURLCreateCopyDeletingLastPathComponent", "", _cFURLCreateCopyDeletingLastPathComponentErr)
	}
	return _cFURLCreateCopyDeletingLastPathComponent(allocator, url), nil
}

// CFURLCreateCopyDeletingLastPathComponent creates a copy of a given URL with the last path component deleted.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingLastPathComponent(_:_:)
func CFURLCreateCopyDeletingLastPathComponent(allocator CFAllocatorRef, url CFURLRef) CFURLRef {
	result, callErr := tryCFURLCreateCopyDeletingLastPathComponent(allocator, url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateCopyDeletingPathExtension func(allocator CFAllocatorRef, url CFURLRef) CFURLRef
var _cFURLCreateCopyDeletingPathExtensionErr error

func tryCFURLCreateCopyDeletingPathExtension(allocator CFAllocatorRef, url CFURLRef) (CFURLRef, error) {
	if _cFURLCreateCopyDeletingPathExtension == nil {
		return 0, symbolCallError("CFURLCreateCopyDeletingPathExtension", "", _cFURLCreateCopyDeletingPathExtensionErr)
	}
	return _cFURLCreateCopyDeletingPathExtension(allocator, url), nil
}

// CFURLCreateCopyDeletingPathExtension creates a copy of a given URL with its last path extension removed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingPathExtension(_:_:)
func CFURLCreateCopyDeletingPathExtension(allocator CFAllocatorRef, url CFURLRef) CFURLRef {
	result, callErr := tryCFURLCreateCopyDeletingPathExtension(allocator, url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateData func(allocator CFAllocatorRef, url CFURLRef, encoding uint32, escapeWhitespace bool) CFDataRef
var _cFURLCreateDataErr error

func tryCFURLCreateData(allocator CFAllocatorRef, url CFURLRef, encoding uint32, escapeWhitespace bool) (CFDataRef, error) {
	if _cFURLCreateData == nil {
		return 0, symbolCallError("CFURLCreateData", "", _cFURLCreateDataErr)
	}
	return _cFURLCreateData(allocator, url, encoding, escapeWhitespace), nil
}

// CFURLCreateData creates a [CFData] object containing the content of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateData(_:_:_:_:)
func CFURLCreateData(allocator CFAllocatorRef, url CFURLRef, encoding uint32, escapeWhitespace bool) CFDataRef {
	result, callErr := tryCFURLCreateData(allocator, url, encoding, escapeWhitespace)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateFilePathURL func(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef
var _cFURLCreateFilePathURLErr error

func tryCFURLCreateFilePathURL(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) (CFURLRef, error) {
	if _cFURLCreateFilePathURL == nil {
		return 0, symbolCallError("CFURLCreateFilePathURL", "10.6", _cFURLCreateFilePathURLErr)
	}
	return _cFURLCreateFilePathURL(allocator, url, err), nil
}

// CFURLCreateFilePathURL returns a new file path URL that refers to the same resource as a specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFilePathURL(_:_:_:)
func CFURLCreateFilePathURL(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef {
	result, callErr := tryCFURLCreateFilePathURL(allocator, url, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateFileReferenceURL func(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef
var _cFURLCreateFileReferenceURLErr error

func tryCFURLCreateFileReferenceURL(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) (CFURLRef, error) {
	if _cFURLCreateFileReferenceURL == nil {
		return 0, symbolCallError("CFURLCreateFileReferenceURL", "10.6", _cFURLCreateFileReferenceURLErr)
	}
	return _cFURLCreateFileReferenceURL(allocator, url, err), nil
}

// CFURLCreateFileReferenceURL returns a new file reference URL that points to the same resource as a specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFileReferenceURL(_:_:_:)
func CFURLCreateFileReferenceURL(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef {
	result, callErr := tryCFURLCreateFileReferenceURL(allocator, url, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateFromFileSystemRepresentation func(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool) CFURLRef
var _cFURLCreateFromFileSystemRepresentationErr error

func tryCFURLCreateFromFileSystemRepresentation(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool) (CFURLRef, error) {
	if _cFURLCreateFromFileSystemRepresentation == nil {
		return 0, symbolCallError("CFURLCreateFromFileSystemRepresentation", "", _cFURLCreateFromFileSystemRepresentationErr)
	}
	return _cFURLCreateFromFileSystemRepresentation(allocator, buffer, bufLen, isDirectory), nil
}

// CFURLCreateFromFileSystemRepresentation creates a new [CFURL] object for a file system entity using the native representation.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentation(_:_:_:_:)
func CFURLCreateFromFileSystemRepresentation(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool) CFURLRef {
	result, callErr := tryCFURLCreateFromFileSystemRepresentation(allocator, buffer, bufLen, isDirectory)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateFromFileSystemRepresentationRelativeToBase func(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool, baseURL CFURLRef) CFURLRef
var _cFURLCreateFromFileSystemRepresentationRelativeToBaseErr error

func tryCFURLCreateFromFileSystemRepresentationRelativeToBase(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool, baseURL CFURLRef) (CFURLRef, error) {
	if _cFURLCreateFromFileSystemRepresentationRelativeToBase == nil {
		return 0, symbolCallError("CFURLCreateFromFileSystemRepresentationRelativeToBase", "", _cFURLCreateFromFileSystemRepresentationRelativeToBaseErr)
	}
	return _cFURLCreateFromFileSystemRepresentationRelativeToBase(allocator, buffer, bufLen, isDirectory, baseURL), nil
}

// CFURLCreateFromFileSystemRepresentationRelativeToBase creates a [CFURL] object from a native character string path relative to a base URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentationRelativeToBase(_:_:_:_:_:)
func CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool, baseURL CFURLRef) CFURLRef {
	result, callErr := tryCFURLCreateFromFileSystemRepresentationRelativeToBase(allocator, buffer, bufLen, isDirectory, baseURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateResourcePropertiesForKeysFromBookmarkData func(allocator CFAllocatorRef, resourcePropertiesToReturn CFArrayRef, bookmark CFDataRef) CFDictionaryRef
var _cFURLCreateResourcePropertiesForKeysFromBookmarkDataErr error

func tryCFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator CFAllocatorRef, resourcePropertiesToReturn CFArrayRef, bookmark CFDataRef) (CFDictionaryRef, error) {
	if _cFURLCreateResourcePropertiesForKeysFromBookmarkData == nil {
		return 0, symbolCallError("CFURLCreateResourcePropertiesForKeysFromBookmarkData", "10.6", _cFURLCreateResourcePropertiesForKeysFromBookmarkDataErr)
	}
	return _cFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator, resourcePropertiesToReturn, bookmark), nil
}

// CFURLCreateResourcePropertiesForKeysFromBookmarkData returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertiesForKeysFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator CFAllocatorRef, resourcePropertiesToReturn CFArrayRef, bookmark CFDataRef) CFDictionaryRef {
	result, callErr := tryCFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator, resourcePropertiesToReturn, bookmark)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateResourcePropertyForKeyFromBookmarkData func(allocator CFAllocatorRef, resourcePropertyKey CFStringRef, bookmark CFDataRef) CFTypeRef
var _cFURLCreateResourcePropertyForKeyFromBookmarkDataErr error

func tryCFURLCreateResourcePropertyForKeyFromBookmarkData(allocator CFAllocatorRef, resourcePropertyKey CFStringRef, bookmark CFDataRef) (CFTypeRef, error) {
	if _cFURLCreateResourcePropertyForKeyFromBookmarkData == nil {
		return 0, symbolCallError("CFURLCreateResourcePropertyForKeyFromBookmarkData", "10.6", _cFURLCreateResourcePropertyForKeyFromBookmarkDataErr)
	}
	return _cFURLCreateResourcePropertyForKeyFromBookmarkData(allocator, resourcePropertyKey, bookmark), nil
}

// CFURLCreateResourcePropertyForKeyFromBookmarkData returns the value of a resource property from specified bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertyForKeyFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator CFAllocatorRef, resourcePropertyKey CFStringRef, bookmark CFDataRef) CFTypeRef {
	result, callErr := tryCFURLCreateResourcePropertyForKeyFromBookmarkData(allocator, resourcePropertyKey, bookmark)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateStringByReplacingPercentEscapes func(allocator CFAllocatorRef, originalString CFStringRef, charactersToLeaveEscaped CFStringRef) CFStringRef
var _cFURLCreateStringByReplacingPercentEscapesErr error

func tryCFURLCreateStringByReplacingPercentEscapes(allocator CFAllocatorRef, originalString CFStringRef, charactersToLeaveEscaped CFStringRef) (CFStringRef, error) {
	if _cFURLCreateStringByReplacingPercentEscapes == nil {
		return 0, symbolCallError("CFURLCreateStringByReplacingPercentEscapes", "", _cFURLCreateStringByReplacingPercentEscapesErr)
	}
	return _cFURLCreateStringByReplacingPercentEscapes(allocator, originalString, charactersToLeaveEscaped), nil
}

// CFURLCreateStringByReplacingPercentEscapes creates a new string by replacing any percent escape sequences with their character equivalent.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapes(_:_:_:)
func CFURLCreateStringByReplacingPercentEscapes(allocator CFAllocatorRef, originalString CFStringRef, charactersToLeaveEscaped CFStringRef) CFStringRef {
	result, callErr := tryCFURLCreateStringByReplacingPercentEscapes(allocator, originalString, charactersToLeaveEscaped)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateWithBytes func(allocator CFAllocatorRef, URLBytes *uint8, length int, encoding uint32, baseURL CFURLRef) CFURLRef
var _cFURLCreateWithBytesErr error

func tryCFURLCreateWithBytes(allocator CFAllocatorRef, URLBytes *uint8, length int, encoding uint32, baseURL CFURLRef) (CFURLRef, error) {
	if _cFURLCreateWithBytes == nil {
		return 0, symbolCallError("CFURLCreateWithBytes", "", _cFURLCreateWithBytesErr)
	}
	return _cFURLCreateWithBytes(allocator, URLBytes, length, encoding, baseURL), nil
}

// CFURLCreateWithBytes creates a [CFURL] object using a given character bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithBytes(_:_:_:_:_:)
func CFURLCreateWithBytes(allocator CFAllocatorRef, URLBytes *uint8, length int, encoding uint32, baseURL CFURLRef) CFURLRef {
	result, callErr := tryCFURLCreateWithBytes(allocator, URLBytes, length, encoding, baseURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateWithFileSystemPath func(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool) CFURLRef
var _cFURLCreateWithFileSystemPathErr error

func tryCFURLCreateWithFileSystemPath(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool) (CFURLRef, error) {
	if _cFURLCreateWithFileSystemPath == nil {
		return 0, symbolCallError("CFURLCreateWithFileSystemPath", "", _cFURLCreateWithFileSystemPathErr)
	}
	return _cFURLCreateWithFileSystemPath(allocator, filePath, pathStyle, isDirectory), nil
}

// CFURLCreateWithFileSystemPath creates a [CFURL] object using a local file system path string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPath(_:_:_:_:)
func CFURLCreateWithFileSystemPath(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool) CFURLRef {
	result, callErr := tryCFURLCreateWithFileSystemPath(allocator, filePath, pathStyle, isDirectory)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateWithFileSystemPathRelativeToBase func(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool, baseURL CFURLRef) CFURLRef
var _cFURLCreateWithFileSystemPathRelativeToBaseErr error

func tryCFURLCreateWithFileSystemPathRelativeToBase(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool, baseURL CFURLRef) (CFURLRef, error) {
	if _cFURLCreateWithFileSystemPathRelativeToBase == nil {
		return 0, symbolCallError("CFURLCreateWithFileSystemPathRelativeToBase", "", _cFURLCreateWithFileSystemPathRelativeToBaseErr)
	}
	return _cFURLCreateWithFileSystemPathRelativeToBase(allocator, filePath, pathStyle, isDirectory, baseURL), nil
}

// CFURLCreateWithFileSystemPathRelativeToBase creates a [CFURL] object using a local file system path string relative to a base URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPathRelativeToBase(_:_:_:_:_:)
func CFURLCreateWithFileSystemPathRelativeToBase(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool, baseURL CFURLRef) CFURLRef {
	result, callErr := tryCFURLCreateWithFileSystemPathRelativeToBase(allocator, filePath, pathStyle, isDirectory, baseURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLCreateWithString func(allocator CFAllocatorRef, URLString CFStringRef, baseURL CFURLRef) CFURLRef
var _cFURLCreateWithStringErr error

func tryCFURLCreateWithString(allocator CFAllocatorRef, URLString CFStringRef, baseURL CFURLRef) (CFURLRef, error) {
	if _cFURLCreateWithString == nil {
		return 0, symbolCallError("CFURLCreateWithString", "", _cFURLCreateWithStringErr)
	}
	return _cFURLCreateWithString(allocator, URLString, baseURL), nil
}

// CFURLCreateWithString creates a [CFURL] object using a given [CFString] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithString(_:_:_:)
func CFURLCreateWithString(allocator CFAllocatorRef, URLString CFStringRef, baseURL CFURLRef) CFURLRef {
	result, callErr := tryCFURLCreateWithString(allocator, URLString, baseURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLEnumeratorCreateForDirectoryURL func(alloc CFAllocatorRef, directoryURL CFURLRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef
var _cFURLEnumeratorCreateForDirectoryURLErr error

func tryCFURLEnumeratorCreateForDirectoryURL(alloc CFAllocatorRef, directoryURL CFURLRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) (CFURLEnumeratorRef, error) {
	if _cFURLEnumeratorCreateForDirectoryURL == nil {
		return 0, symbolCallError("CFURLEnumeratorCreateForDirectoryURL", "10.6", _cFURLEnumeratorCreateForDirectoryURLErr)
	}
	return _cFURLEnumeratorCreateForDirectoryURL(alloc, directoryURL, option, propertyKeys), nil
}

// CFURLEnumeratorCreateForDirectoryURL creates and returns a directory enumerator with provided enumerator behavior options and properties to be prefetched.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForDirectoryURL(_:_:_:_:)
func CFURLEnumeratorCreateForDirectoryURL(alloc CFAllocatorRef, directoryURL CFURLRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef {
	result, callErr := tryCFURLEnumeratorCreateForDirectoryURL(alloc, directoryURL, option, propertyKeys)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLEnumeratorCreateForMountedVolumes func(alloc CFAllocatorRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef
var _cFURLEnumeratorCreateForMountedVolumesErr error

func tryCFURLEnumeratorCreateForMountedVolumes(alloc CFAllocatorRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) (CFURLEnumeratorRef, error) {
	if _cFURLEnumeratorCreateForMountedVolumes == nil {
		return 0, symbolCallError("CFURLEnumeratorCreateForMountedVolumes", "10.6", _cFURLEnumeratorCreateForMountedVolumesErr)
	}
	return _cFURLEnumeratorCreateForMountedVolumes(alloc, option, propertyKeys), nil
}

// CFURLEnumeratorCreateForMountedVolumes creates and returns a volume enumerator with provided enumerator behavior options and properties to be prefetched.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForMountedVolumes(_:_:_:)
func CFURLEnumeratorCreateForMountedVolumes(alloc CFAllocatorRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef {
	result, callErr := tryCFURLEnumeratorCreateForMountedVolumes(alloc, option, propertyKeys)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLEnumeratorGetDescendentLevel func(enumerator CFURLEnumeratorRef) int
var _cFURLEnumeratorGetDescendentLevelErr error

func tryCFURLEnumeratorGetDescendentLevel(enumerator CFURLEnumeratorRef) (int, error) {
	if _cFURLEnumeratorGetDescendentLevel == nil {
		return 0, symbolCallError("CFURLEnumeratorGetDescendentLevel", "10.6", _cFURLEnumeratorGetDescendentLevelErr)
	}
	return _cFURLEnumeratorGetDescendentLevel(enumerator), nil
}

// CFURLEnumeratorGetDescendentLevel returns the number of levels a recursive directory enumerator has descended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetDescendentLevel(_:)
func CFURLEnumeratorGetDescendentLevel(enumerator CFURLEnumeratorRef) int {
	result, callErr := tryCFURLEnumeratorGetDescendentLevel(enumerator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLEnumeratorGetNextURL func(enumerator CFURLEnumeratorRef, url *CFURLRef, err *CFErrorRef) CFURLEnumeratorResult
var _cFURLEnumeratorGetNextURLErr error

func tryCFURLEnumeratorGetNextURL(enumerator CFURLEnumeratorRef, url *CFURLRef, err *CFErrorRef) (CFURLEnumeratorResult, error) {
	if _cFURLEnumeratorGetNextURL == nil {
		return *new(CFURLEnumeratorResult), symbolCallError("CFURLEnumeratorGetNextURL", "10.6", _cFURLEnumeratorGetNextURLErr)
	}
	return _cFURLEnumeratorGetNextURL(enumerator, url, err), nil
}

// CFURLEnumeratorGetNextURL advances an enumerator to the next URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetNextURL(_:_:_:)
func CFURLEnumeratorGetNextURL(enumerator CFURLEnumeratorRef, url *CFURLRef, err *CFErrorRef) CFURLEnumeratorResult {
	result, callErr := tryCFURLEnumeratorGetNextURL(enumerator, url, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLEnumeratorGetTypeID func() uint
var _cFURLEnumeratorGetTypeIDErr error

func tryCFURLEnumeratorGetTypeID() (uint, error) {
	if _cFURLEnumeratorGetTypeID == nil {
		return 0, symbolCallError("CFURLEnumeratorGetTypeID", "10.6", _cFURLEnumeratorGetTypeIDErr)
	}
	return _cFURLEnumeratorGetTypeID(), nil
}

// CFURLEnumeratorGetTypeID returns the opaque type identifier for the CFURLEnumerator opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetTypeID()
func CFURLEnumeratorGetTypeID() uint {
	result, callErr := tryCFURLEnumeratorGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLEnumeratorSkipDescendents func(enumerator CFURLEnumeratorRef)
var _cFURLEnumeratorSkipDescendentsErr error

func tryCFURLEnumeratorSkipDescendents(enumerator CFURLEnumeratorRef) error {
	if _cFURLEnumeratorSkipDescendents == nil {
		return symbolCallError("CFURLEnumeratorSkipDescendents", "10.6", _cFURLEnumeratorSkipDescendentsErr)
	}
	_cFURLEnumeratorSkipDescendents(enumerator)
	return nil
}

// CFURLEnumeratorSkipDescendents tells a recursive enumerator not to descend into the directory at the URL that was returned by the most recent call to the CFURLEnumeratorGetNextURL(_:_:_:) function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorSkipDescendents(_:)
func CFURLEnumeratorSkipDescendents(enumerator CFURLEnumeratorRef) {
	if callErr := tryCFURLEnumeratorSkipDescendents(enumerator); callErr != nil {
		panic(callErr)
	}
}

var _cFURLGetBaseURL func(anURL CFURLRef) CFURLRef
var _cFURLGetBaseURLErr error

func tryCFURLGetBaseURL(anURL CFURLRef) (CFURLRef, error) {
	if _cFURLGetBaseURL == nil {
		return 0, symbolCallError("CFURLGetBaseURL", "", _cFURLGetBaseURLErr)
	}
	return _cFURLGetBaseURL(anURL), nil
}

// CFURLGetBaseURL returns the base URL of a given URL if it exists.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBaseURL(_:)
func CFURLGetBaseURL(anURL CFURLRef) CFURLRef {
	result, callErr := tryCFURLGetBaseURL(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLGetByteRangeForComponent func(url CFURLRef, component CFURLComponentType, rangeIncludingSeparators *CFRange) CFRange
var _cFURLGetByteRangeForComponentErr error

func tryCFURLGetByteRangeForComponent(url CFURLRef, component CFURLComponentType, rangeIncludingSeparators *CFRange) (CFRange, error) {
	if _cFURLGetByteRangeForComponent == nil {
		return CFRange{}, symbolCallError("CFURLGetByteRangeForComponent", "", _cFURLGetByteRangeForComponentErr)
	}
	return _cFURLGetByteRangeForComponent(url, component, rangeIncludingSeparators), nil
}

// CFURLGetByteRangeForComponent returns the range of the specified component in the bytes of a URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetByteRangeForComponent(_:_:_:)
func CFURLGetByteRangeForComponent(url CFURLRef, component CFURLComponentType, rangeIncludingSeparators *CFRange) CFRange {
	result, callErr := tryCFURLGetByteRangeForComponent(url, component, rangeIncludingSeparators)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLGetBytes func(url CFURLRef, buffer *uint8, bufferLength int) int
var _cFURLGetBytesErr error

func tryCFURLGetBytes(url CFURLRef, buffer *uint8, bufferLength int) (int, error) {
	if _cFURLGetBytes == nil {
		return 0, symbolCallError("CFURLGetBytes", "", _cFURLGetBytesErr)
	}
	return _cFURLGetBytes(url, buffer, bufferLength), nil
}

// CFURLGetBytes returns by reference the byte representation of a URL object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBytes(_:_:_:)
func CFURLGetBytes(url CFURLRef, buffer *uint8, bufferLength int) int {
	result, callErr := tryCFURLGetBytes(url, buffer, bufferLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLGetFileSystemRepresentation func(url CFURLRef, resolveAgainstBase bool, buffer *uint8, maxBufLen int) bool
var _cFURLGetFileSystemRepresentationErr error

func tryCFURLGetFileSystemRepresentation(url CFURLRef, resolveAgainstBase bool, buffer *uint8, maxBufLen int) (bool, error) {
	if _cFURLGetFileSystemRepresentation == nil {
		return false, symbolCallError("CFURLGetFileSystemRepresentation", "", _cFURLGetFileSystemRepresentationErr)
	}
	return _cFURLGetFileSystemRepresentation(url, resolveAgainstBase, buffer, maxBufLen), nil
}

// CFURLGetFileSystemRepresentation fills a buffer with the file system’s native string representation of a given URL’s path.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFileSystemRepresentation(_:_:_:_:)
func CFURLGetFileSystemRepresentation(url CFURLRef, resolveAgainstBase bool, buffer *uint8, maxBufLen int) bool {
	result, callErr := tryCFURLGetFileSystemRepresentation(url, resolveAgainstBase, buffer, maxBufLen)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLGetPortNumber func(anURL CFURLRef) int32
var _cFURLGetPortNumberErr error

func tryCFURLGetPortNumber(anURL CFURLRef) (int32, error) {
	if _cFURLGetPortNumber == nil {
		return 0, symbolCallError("CFURLGetPortNumber", "", _cFURLGetPortNumberErr)
	}
	return _cFURLGetPortNumber(anURL), nil
}

// CFURLGetPortNumber returns the port number from a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetPortNumber(_:)
func CFURLGetPortNumber(anURL CFURLRef) int32 {
	result, callErr := tryCFURLGetPortNumber(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLGetString func(anURL CFURLRef) CFStringRef
var _cFURLGetStringErr error

func tryCFURLGetString(anURL CFURLRef) (CFStringRef, error) {
	if _cFURLGetString == nil {
		return 0, symbolCallError("CFURLGetString", "", _cFURLGetStringErr)
	}
	return _cFURLGetString(anURL), nil
}

// CFURLGetString returns the URL as a [CFString] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetString(_:)
func CFURLGetString(anURL CFURLRef) CFStringRef {
	result, callErr := tryCFURLGetString(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLGetTypeID func() uint
var _cFURLGetTypeIDErr error

func tryCFURLGetTypeID() (uint, error) {
	if _cFURLGetTypeID == nil {
		return 0, symbolCallError("CFURLGetTypeID", "", _cFURLGetTypeIDErr)
	}
	return _cFURLGetTypeID(), nil
}

// CFURLGetTypeID returns the type identifier for the [CFURL] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetTypeID()
func CFURLGetTypeID() uint {
	result, callErr := tryCFURLGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLHasDirectoryPath func(anURL CFURLRef) bool
var _cFURLHasDirectoryPathErr error

func tryCFURLHasDirectoryPath(anURL CFURLRef) (bool, error) {
	if _cFURLHasDirectoryPath == nil {
		return false, symbolCallError("CFURLHasDirectoryPath", "", _cFURLHasDirectoryPathErr)
	}
	return _cFURLHasDirectoryPath(anURL), nil
}

// CFURLHasDirectoryPath determines if a given URL’s path represents a directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLHasDirectoryPath(_:)
func CFURLHasDirectoryPath(anURL CFURLRef) bool {
	result, callErr := tryCFURLHasDirectoryPath(anURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLIsFileReferenceURL func(url CFURLRef) bool
var _cFURLIsFileReferenceURLErr error

func tryCFURLIsFileReferenceURL(url CFURLRef) (bool, error) {
	if _cFURLIsFileReferenceURL == nil {
		return false, symbolCallError("CFURLIsFileReferenceURL", "10.9", _cFURLIsFileReferenceURLErr)
	}
	return _cFURLIsFileReferenceURL(url), nil
}

// CFURLIsFileReferenceURL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLIsFileReferenceURL(_:)
func CFURLIsFileReferenceURL(url CFURLRef) bool {
	result, callErr := tryCFURLIsFileReferenceURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLResourceIsReachable func(url CFURLRef, err *CFErrorRef) bool
var _cFURLResourceIsReachableErr error

func tryCFURLResourceIsReachable(url CFURLRef, err *CFErrorRef) (bool, error) {
	if _cFURLResourceIsReachable == nil {
		return false, symbolCallError("CFURLResourceIsReachable", "10.6", _cFURLResourceIsReachableErr)
	}
	return _cFURLResourceIsReachable(url, err), nil
}

// CFURLResourceIsReachable returns whether the resource pointed to by a file URL can be reached.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLResourceIsReachable(_:_:)
func CFURLResourceIsReachable(url CFURLRef, err *CFErrorRef) bool {
	result, callErr := tryCFURLResourceIsReachable(url, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLSetResourcePropertiesForKeys func(url CFURLRef, keyedPropertyValues CFDictionaryRef, err *CFErrorRef) bool
var _cFURLSetResourcePropertiesForKeysErr error

func tryCFURLSetResourcePropertiesForKeys(url CFURLRef, keyedPropertyValues CFDictionaryRef, err *CFErrorRef) (bool, error) {
	if _cFURLSetResourcePropertiesForKeys == nil {
		return false, symbolCallError("CFURLSetResourcePropertiesForKeys", "10.6", _cFURLSetResourcePropertiesForKeysErr)
	}
	return _cFURLSetResourcePropertiesForKeys(url, keyedPropertyValues, err), nil
}

// CFURLSetResourcePropertiesForKeys sets the URL’s resource properties for a given set of keys to a given set of values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertiesForKeys(_:_:_:)
func CFURLSetResourcePropertiesForKeys(url CFURLRef, keyedPropertyValues CFDictionaryRef, err *CFErrorRef) bool {
	result, callErr := tryCFURLSetResourcePropertiesForKeys(url, keyedPropertyValues, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLSetResourcePropertyForKey func(url CFURLRef, key CFStringRef, propertyValue CFTypeRef, err *CFErrorRef) bool
var _cFURLSetResourcePropertyForKeyErr error

func tryCFURLSetResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValue CFTypeRef, err *CFErrorRef) (bool, error) {
	if _cFURLSetResourcePropertyForKey == nil {
		return false, symbolCallError("CFURLSetResourcePropertyForKey", "10.6", _cFURLSetResourcePropertyForKeyErr)
	}
	return _cFURLSetResourcePropertyForKey(url, key, propertyValue, err), nil
}

// CFURLSetResourcePropertyForKey sets the URL’s resource property for a given key to a given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertyForKey(_:_:_:_:)
func CFURLSetResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValue CFTypeRef, err *CFErrorRef) bool {
	result, callErr := tryCFURLSetResourcePropertyForKey(url, key, propertyValue, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLSetTemporaryResourcePropertyForKey func(url CFURLRef, key CFStringRef, propertyValue CFTypeRef)
var _cFURLSetTemporaryResourcePropertyForKeyErr error

func tryCFURLSetTemporaryResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValue CFTypeRef) error {
	if _cFURLSetTemporaryResourcePropertyForKey == nil {
		return symbolCallError("CFURLSetTemporaryResourcePropertyForKey", "10.6", _cFURLSetTemporaryResourcePropertyForKeyErr)
	}
	_cFURLSetTemporaryResourcePropertyForKey(url, key, propertyValue)
	return nil
}

// CFURLSetTemporaryResourcePropertyForKey sets a temporary resource value on the URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLSetTemporaryResourcePropertyForKey(_:_:_:)
func CFURLSetTemporaryResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValue CFTypeRef) {
	if callErr := tryCFURLSetTemporaryResourcePropertyForKey(url, key, propertyValue); callErr != nil {
		panic(callErr)
	}
}

var _cFURLStartAccessingSecurityScopedResource func(url CFURLRef) bool
var _cFURLStartAccessingSecurityScopedResourceErr error

func tryCFURLStartAccessingSecurityScopedResource(url CFURLRef) (bool, error) {
	if _cFURLStartAccessingSecurityScopedResource == nil {
		return false, symbolCallError("CFURLStartAccessingSecurityScopedResource", "10.7", _cFURLStartAccessingSecurityScopedResourceErr)
	}
	return _cFURLStartAccessingSecurityScopedResource(url), nil
}

// CFURLStartAccessingSecurityScopedResource in an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
func CFURLStartAccessingSecurityScopedResource(url CFURLRef) bool {
	result, callErr := tryCFURLStartAccessingSecurityScopedResource(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFURLStopAccessingSecurityScopedResource func(url CFURLRef)
var _cFURLStopAccessingSecurityScopedResourceErr error

func tryCFURLStopAccessingSecurityScopedResource(url CFURLRef) error {
	if _cFURLStopAccessingSecurityScopedResource == nil {
		return symbolCallError("CFURLStopAccessingSecurityScopedResource", "10.7", _cFURLStopAccessingSecurityScopedResourceErr)
	}
	_cFURLStopAccessingSecurityScopedResource(url)
	return nil
}

// CFURLStopAccessingSecurityScopedResource in an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
func CFURLStopAccessingSecurityScopedResource(url CFURLRef) {
	if callErr := tryCFURLStopAccessingSecurityScopedResource(url); callErr != nil {
		panic(callErr)
	}
}

var _cFURLWriteBookmarkDataToFile func(bookmarkRef CFDataRef, fileURL CFURLRef, options CFURLBookmarkFileCreationOptions, errorRef *CFErrorRef) bool
var _cFURLWriteBookmarkDataToFileErr error

func tryCFURLWriteBookmarkDataToFile(bookmarkRef CFDataRef, fileURL CFURLRef, options CFURLBookmarkFileCreationOptions, errorRef *CFErrorRef) (bool, error) {
	if _cFURLWriteBookmarkDataToFile == nil {
		return false, symbolCallError("CFURLWriteBookmarkDataToFile", "10.6", _cFURLWriteBookmarkDataToFileErr)
	}
	return _cFURLWriteBookmarkDataToFile(bookmarkRef, fileURL, options, errorRef), nil
}

// CFURLWriteBookmarkDataToFile creates an alias file on disk at a specified location with specified bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteBookmarkDataToFile(_:_:_:_:)
func CFURLWriteBookmarkDataToFile(bookmarkRef CFDataRef, fileURL CFURLRef, options CFURLBookmarkFileCreationOptions, errorRef *CFErrorRef) bool {
	result, callErr := tryCFURLWriteBookmarkDataToFile(bookmarkRef, fileURL, options, errorRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDCreate func(alloc CFAllocatorRef) CFUUIDRef
var _cFUUIDCreateErr error

func tryCFUUIDCreate(alloc CFAllocatorRef) (CFUUIDRef, error) {
	if _cFUUIDCreate == nil {
		return 0, symbolCallError("CFUUIDCreate", "", _cFUUIDCreateErr)
	}
	return _cFUUIDCreate(alloc), nil
}

// CFUUIDCreate creates a Universally Unique Identifier (UUID) object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreate(_:)
func CFUUIDCreate(alloc CFAllocatorRef) CFUUIDRef {
	result, callErr := tryCFUUIDCreate(alloc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDCreateFromString func(alloc CFAllocatorRef, uuidStr CFStringRef) CFUUIDRef
var _cFUUIDCreateFromStringErr error

func tryCFUUIDCreateFromString(alloc CFAllocatorRef, uuidStr CFStringRef) (CFUUIDRef, error) {
	if _cFUUIDCreateFromString == nil {
		return 0, symbolCallError("CFUUIDCreateFromString", "", _cFUUIDCreateFromStringErr)
	}
	return _cFUUIDCreateFromString(alloc, uuidStr), nil
}

// CFUUIDCreateFromString creates a CFUUID object for a specified string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromString(_:_:)
func CFUUIDCreateFromString(alloc CFAllocatorRef, uuidStr CFStringRef) CFUUIDRef {
	result, callErr := tryCFUUIDCreateFromString(alloc, uuidStr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDCreateFromUUIDBytes func(alloc CFAllocatorRef, bytes CFUUIDBytes) CFUUIDRef
var _cFUUIDCreateFromUUIDBytesErr error

func tryCFUUIDCreateFromUUIDBytes(alloc CFAllocatorRef, bytes CFUUIDBytes) (CFUUIDRef, error) {
	if _cFUUIDCreateFromUUIDBytes == nil {
		return 0, symbolCallError("CFUUIDCreateFromUUIDBytes", "", _cFUUIDCreateFromUUIDBytesErr)
	}
	return _cFUUIDCreateFromUUIDBytes(alloc, bytes), nil
}

// CFUUIDCreateFromUUIDBytes creates a CFUUID object from raw UUID bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromUUIDBytes(_:_:)
func CFUUIDCreateFromUUIDBytes(alloc CFAllocatorRef, bytes CFUUIDBytes) CFUUIDRef {
	result, callErr := tryCFUUIDCreateFromUUIDBytes(alloc, bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDCreateString func(alloc CFAllocatorRef, uuid CFUUIDRef) CFStringRef
var _cFUUIDCreateStringErr error

func tryCFUUIDCreateString(alloc CFAllocatorRef, uuid CFUUIDRef) (CFStringRef, error) {
	if _cFUUIDCreateString == nil {
		return 0, symbolCallError("CFUUIDCreateString", "", _cFUUIDCreateStringErr)
	}
	return _cFUUIDCreateString(alloc, uuid), nil
}

// CFUUIDCreateString returns the string representation of a specified CFUUID object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateString(_:_:)
func CFUUIDCreateString(alloc CFAllocatorRef, uuid CFUUIDRef) CFStringRef {
	result, callErr := tryCFUUIDCreateString(alloc, uuid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDCreateWithBytes func(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef
var _cFUUIDCreateWithBytesErr error

func tryCFUUIDCreateWithBytes(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) (CFUUIDRef, error) {
	if _cFUUIDCreateWithBytes == nil {
		return 0, symbolCallError("CFUUIDCreateWithBytes", "", _cFUUIDCreateWithBytesErr)
	}
	return _cFUUIDCreateWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15), nil
}

// CFUUIDCreateWithBytes creates a CFUUID object from raw UUID bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDCreateWithBytes(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef {
	result, callErr := tryCFUUIDCreateWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDGetConstantUUIDWithBytes func(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef
var _cFUUIDGetConstantUUIDWithBytesErr error

func tryCFUUIDGetConstantUUIDWithBytes(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) (CFUUIDRef, error) {
	if _cFUUIDGetConstantUUIDWithBytes == nil {
		return 0, symbolCallError("CFUUIDGetConstantUUIDWithBytes", "", _cFUUIDGetConstantUUIDWithBytesErr)
	}
	return _cFUUIDGetConstantUUIDWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15), nil
}

// CFUUIDGetConstantUUIDWithBytes returns a CFUUID object from raw UUID bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetConstantUUIDWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDGetConstantUUIDWithBytes(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef {
	result, callErr := tryCFUUIDGetConstantUUIDWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDGetTypeID func() uint
var _cFUUIDGetTypeIDErr error

func tryCFUUIDGetTypeID() (uint, error) {
	if _cFUUIDGetTypeID == nil {
		return 0, symbolCallError("CFUUIDGetTypeID", "", _cFUUIDGetTypeIDErr)
	}
	return _cFUUIDGetTypeID(), nil
}

// CFUUIDGetTypeID returns the type identifier for all CFUUID objects.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetTypeID()
func CFUUIDGetTypeID() uint {
	result, callErr := tryCFUUIDGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUUIDGetUUIDBytes func(uuid CFUUIDRef) CFUUIDBytes
var _cFUUIDGetUUIDBytesErr error

func tryCFUUIDGetUUIDBytes(uuid CFUUIDRef) (CFUUIDBytes, error) {
	if _cFUUIDGetUUIDBytes == nil {
		return CFUUIDBytes{}, symbolCallError("CFUUIDGetUUIDBytes", "", _cFUUIDGetUUIDBytesErr)
	}
	return _cFUUIDGetUUIDBytes(uuid), nil
}

// CFUUIDGetUUIDBytes returns the value of a UUID object as raw bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetUUIDBytes(_:)
func CFUUIDGetUUIDBytes(uuid CFUUIDRef) CFUUIDBytes {
	result, callErr := tryCFUUIDGetUUIDBytes(uuid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationCancel func(userNotification CFUserNotificationRef) int32
var _cFUserNotificationCancelErr error

func tryCFUserNotificationCancel(userNotification CFUserNotificationRef) (int32, error) {
	if _cFUserNotificationCancel == nil {
		return 0, symbolCallError("CFUserNotificationCancel", "10.0", _cFUserNotificationCancelErr)
	}
	return _cFUserNotificationCancel(userNotification), nil
}

// CFUserNotificationCancel cancels a user notification dialog.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCancel(_:)
func CFUserNotificationCancel(userNotification CFUserNotificationRef) int32 {
	result, callErr := tryCFUserNotificationCancel(userNotification)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationCreate func(allocator CFAllocatorRef, timeout float64, flags uint64, err *int32, dictionary CFDictionaryRef) CFUserNotificationRef
var _cFUserNotificationCreateErr error

func tryCFUserNotificationCreate(allocator CFAllocatorRef, timeout float64, flags uint64, err *int32, dictionary CFDictionaryRef) (CFUserNotificationRef, error) {
	if _cFUserNotificationCreate == nil {
		return 0, symbolCallError("CFUserNotificationCreate", "10.0", _cFUserNotificationCreateErr)
	}
	return _cFUserNotificationCreate(allocator, timeout, flags, err, dictionary), nil
}

// CFUserNotificationCreate creates a CFUserNotification object and displays its notification dialog on screen.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreate(_:_:_:_:_:)
func CFUserNotificationCreate(allocator CFAllocatorRef, timeout float64, flags uint64, err *int32, dictionary CFDictionaryRef) CFUserNotificationRef {
	result, callErr := tryCFUserNotificationCreate(allocator, timeout, flags, err, dictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationCreateRunLoopSource func(allocator CFAllocatorRef, userNotification CFUserNotificationRef, callout CFUserNotificationCallBack, order int) CFRunLoopSourceRef
var _cFUserNotificationCreateRunLoopSourceErr error

func tryCFUserNotificationCreateRunLoopSource(allocator CFAllocatorRef, userNotification CFUserNotificationRef, callout CFUserNotificationCallBack, order int) (CFRunLoopSourceRef, error) {
	if _cFUserNotificationCreateRunLoopSource == nil {
		return 0, symbolCallError("CFUserNotificationCreateRunLoopSource", "10.0", _cFUserNotificationCreateRunLoopSourceErr)
	}
	return _cFUserNotificationCreateRunLoopSource(allocator, userNotification, callout, order), nil
}

// CFUserNotificationCreateRunLoopSource creates a run loop source for a user notification.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreateRunLoopSource(_:_:_:_:)
func CFUserNotificationCreateRunLoopSource(allocator CFAllocatorRef, userNotification CFUserNotificationRef, callout CFUserNotificationCallBack, order int) CFRunLoopSourceRef {
	result, callErr := tryCFUserNotificationCreateRunLoopSource(allocator, userNotification, callout, order)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationDisplayAlert func(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef, alternateButtonTitle CFStringRef, otherButtonTitle CFStringRef, responseFlags *uint64) int32
var _cFUserNotificationDisplayAlertErr error

func tryCFUserNotificationDisplayAlert(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef, alternateButtonTitle CFStringRef, otherButtonTitle CFStringRef, responseFlags *uint64) (int32, error) {
	if _cFUserNotificationDisplayAlert == nil {
		return 0, symbolCallError("CFUserNotificationDisplayAlert", "10.0", _cFUserNotificationDisplayAlertErr)
	}
	return _cFUserNotificationDisplayAlert(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle, alternateButtonTitle, otherButtonTitle, responseFlags), nil
}

// CFUserNotificationDisplayAlert displays a user notification dialog and waits for a user response.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayAlert(_:_:_:_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayAlert(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef, alternateButtonTitle CFStringRef, otherButtonTitle CFStringRef, responseFlags *uint64) int32 {
	result, callErr := tryCFUserNotificationDisplayAlert(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle, alternateButtonTitle, otherButtonTitle, responseFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationDisplayNotice func(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef) int32
var _cFUserNotificationDisplayNoticeErr error

func tryCFUserNotificationDisplayNotice(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef) (int32, error) {
	if _cFUserNotificationDisplayNotice == nil {
		return 0, symbolCallError("CFUserNotificationDisplayNotice", "10.0", _cFUserNotificationDisplayNoticeErr)
	}
	return _cFUserNotificationDisplayNotice(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle), nil
}

// CFUserNotificationDisplayNotice displays a user notification dialog that does not need a user response.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayNotice(_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayNotice(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef) int32 {
	result, callErr := tryCFUserNotificationDisplayNotice(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationGetResponseDictionary func(userNotification CFUserNotificationRef) CFDictionaryRef
var _cFUserNotificationGetResponseDictionaryErr error

func tryCFUserNotificationGetResponseDictionary(userNotification CFUserNotificationRef) (CFDictionaryRef, error) {
	if _cFUserNotificationGetResponseDictionary == nil {
		return 0, symbolCallError("CFUserNotificationGetResponseDictionary", "10.0", _cFUserNotificationGetResponseDictionaryErr)
	}
	return _cFUserNotificationGetResponseDictionary(userNotification), nil
}

// CFUserNotificationGetResponseDictionary returns the dictionary containing all the text field values from a dismissed notification dialog.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseDictionary(_:)
func CFUserNotificationGetResponseDictionary(userNotification CFUserNotificationRef) CFDictionaryRef {
	result, callErr := tryCFUserNotificationGetResponseDictionary(userNotification)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationGetResponseValue func(userNotification CFUserNotificationRef, key CFStringRef, idx int) CFStringRef
var _cFUserNotificationGetResponseValueErr error

func tryCFUserNotificationGetResponseValue(userNotification CFUserNotificationRef, key CFStringRef, idx int) (CFStringRef, error) {
	if _cFUserNotificationGetResponseValue == nil {
		return 0, symbolCallError("CFUserNotificationGetResponseValue", "10.0", _cFUserNotificationGetResponseValueErr)
	}
	return _cFUserNotificationGetResponseValue(userNotification, key, idx), nil
}

// CFUserNotificationGetResponseValue extracts the values of the text fields from a dismissed notification dialog.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseValue(_:_:_:)
func CFUserNotificationGetResponseValue(userNotification CFUserNotificationRef, key CFStringRef, idx int) CFStringRef {
	result, callErr := tryCFUserNotificationGetResponseValue(userNotification, key, idx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationGetTypeID func() uint
var _cFUserNotificationGetTypeIDErr error

func tryCFUserNotificationGetTypeID() (uint, error) {
	if _cFUserNotificationGetTypeID == nil {
		return 0, symbolCallError("CFUserNotificationGetTypeID", "10.0", _cFUserNotificationGetTypeIDErr)
	}
	return _cFUserNotificationGetTypeID(), nil
}

// CFUserNotificationGetTypeID returns the type identifier for the [CFUserNotification] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetTypeID()
func CFUserNotificationGetTypeID() uint {
	result, callErr := tryCFUserNotificationGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationReceiveResponse func(userNotification CFUserNotificationRef, timeout float64, responseFlags *uint64) int32
var _cFUserNotificationReceiveResponseErr error

func tryCFUserNotificationReceiveResponse(userNotification CFUserNotificationRef, timeout float64, responseFlags *uint64) (int32, error) {
	if _cFUserNotificationReceiveResponse == nil {
		return 0, symbolCallError("CFUserNotificationReceiveResponse", "10.0", _cFUserNotificationReceiveResponseErr)
	}
	return _cFUserNotificationReceiveResponse(userNotification, timeout, responseFlags), nil
}

// CFUserNotificationReceiveResponse waits for the user to respond to a notification or for the notification to time out.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationReceiveResponse(_:_:_:)
func CFUserNotificationReceiveResponse(userNotification CFUserNotificationRef, timeout float64, responseFlags *uint64) int32 {
	result, callErr := tryCFUserNotificationReceiveResponse(userNotification, timeout, responseFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFUserNotificationUpdate func(userNotification CFUserNotificationRef, timeout float64, flags uint64, dictionary CFDictionaryRef) int32
var _cFUserNotificationUpdateErr error

func tryCFUserNotificationUpdate(userNotification CFUserNotificationRef, timeout float64, flags uint64, dictionary CFDictionaryRef) (int32, error) {
	if _cFUserNotificationUpdate == nil {
		return 0, symbolCallError("CFUserNotificationUpdate", "10.0", _cFUserNotificationUpdateErr)
	}
	return _cFUserNotificationUpdate(userNotification, timeout, flags, dictionary), nil
}

// CFUserNotificationUpdate updates a displayed user notification dialog with new user interface information.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationUpdate(_:_:_:_:)
func CFUserNotificationUpdate(userNotification CFUserNotificationRef, timeout float64, flags uint64, dictionary CFDictionaryRef) int32 {
	result, callErr := tryCFUserNotificationUpdate(userNotification, timeout, flags, dictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamCanAcceptBytes func(stream CFWriteStreamRef) bool
var _cFWriteStreamCanAcceptBytesErr error

func tryCFWriteStreamCanAcceptBytes(stream CFWriteStreamRef) (bool, error) {
	if _cFWriteStreamCanAcceptBytes == nil {
		return false, symbolCallError("CFWriteStreamCanAcceptBytes", "", _cFWriteStreamCanAcceptBytesErr)
	}
	return _cFWriteStreamCanAcceptBytes(stream), nil
}

// CFWriteStreamCanAcceptBytes returns whether a writable stream can accept new data without blocking.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCanAcceptBytes(_:)
func CFWriteStreamCanAcceptBytes(stream CFWriteStreamRef) bool {
	result, callErr := tryCFWriteStreamCanAcceptBytes(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamClose func(stream CFWriteStreamRef)
var _cFWriteStreamCloseErr error

func tryCFWriteStreamClose(stream CFWriteStreamRef) error {
	if _cFWriteStreamClose == nil {
		return symbolCallError("CFWriteStreamClose", "", _cFWriteStreamCloseErr)
	}
	_cFWriteStreamClose(stream)
	return nil
}

// CFWriteStreamClose closes a writable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClose(_:)
func CFWriteStreamClose(stream CFWriteStreamRef) {
	if callErr := tryCFWriteStreamClose(stream); callErr != nil {
		panic(callErr)
	}
}

var _cFWriteStreamCopyDispatchQueue func(stream CFWriteStreamRef) uintptr
var _cFWriteStreamCopyDispatchQueueErr error

func tryCFWriteStreamCopyDispatchQueue(stream CFWriteStreamRef) (dispatch.Queue, error) {
	if _cFWriteStreamCopyDispatchQueue == nil {
		return dispatch.QueueFromHandle(0), symbolCallError("CFWriteStreamCopyDispatchQueue", "10.9", _cFWriteStreamCopyDispatchQueueErr)
	}
	return dispatch.QueueFromHandle(_cFWriteStreamCopyDispatchQueue(stream)), nil
}

// CFWriteStreamCopyDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyDispatchQueue(_:)
func CFWriteStreamCopyDispatchQueue(stream CFWriteStreamRef) dispatch.Queue {
	result, callErr := tryCFWriteStreamCopyDispatchQueue(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamCopyError func(stream CFWriteStreamRef) CFErrorRef
var _cFWriteStreamCopyErrorErr error

func tryCFWriteStreamCopyError(stream CFWriteStreamRef) (CFErrorRef, error) {
	if _cFWriteStreamCopyError == nil {
		return 0, symbolCallError("CFWriteStreamCopyError", "10.5", _cFWriteStreamCopyErrorErr)
	}
	return _cFWriteStreamCopyError(stream), nil
}

// CFWriteStreamCopyError returns the error associated with a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyError(_:)
func CFWriteStreamCopyError(stream CFWriteStreamRef) CFErrorRef {
	result, callErr := tryCFWriteStreamCopyError(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamCopyProperty func(stream CFWriteStreamRef, propertyName CFStreamPropertyKey) CFTypeRef
var _cFWriteStreamCopyPropertyErr error

func tryCFWriteStreamCopyProperty(stream CFWriteStreamRef, propertyName CFStreamPropertyKey) (CFTypeRef, error) {
	if _cFWriteStreamCopyProperty == nil {
		return 0, symbolCallError("CFWriteStreamCopyProperty", "", _cFWriteStreamCopyPropertyErr)
	}
	return _cFWriteStreamCopyProperty(stream, propertyName), nil
}

// CFWriteStreamCopyProperty returns the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyProperty(_:_:)
func CFWriteStreamCopyProperty(stream CFWriteStreamRef, propertyName CFStreamPropertyKey) CFTypeRef {
	result, callErr := tryCFWriteStreamCopyProperty(stream, propertyName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamCreateWithAllocatedBuffers func(alloc CFAllocatorRef, bufferAllocator CFAllocatorRef) CFWriteStreamRef
var _cFWriteStreamCreateWithAllocatedBuffersErr error

func tryCFWriteStreamCreateWithAllocatedBuffers(alloc CFAllocatorRef, bufferAllocator CFAllocatorRef) (CFWriteStreamRef, error) {
	if _cFWriteStreamCreateWithAllocatedBuffers == nil {
		return 0, symbolCallError("CFWriteStreamCreateWithAllocatedBuffers", "", _cFWriteStreamCreateWithAllocatedBuffersErr)
	}
	return _cFWriteStreamCreateWithAllocatedBuffers(alloc, bufferAllocator), nil
}

// CFWriteStreamCreateWithAllocatedBuffers creates a writable stream for a growable block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithAllocatedBuffers(_:_:)
func CFWriteStreamCreateWithAllocatedBuffers(alloc CFAllocatorRef, bufferAllocator CFAllocatorRef) CFWriteStreamRef {
	result, callErr := tryCFWriteStreamCreateWithAllocatedBuffers(alloc, bufferAllocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamCreateWithBuffer func(alloc CFAllocatorRef, buffer *uint8, bufferCapacity int) CFWriteStreamRef
var _cFWriteStreamCreateWithBufferErr error

func tryCFWriteStreamCreateWithBuffer(alloc CFAllocatorRef, buffer *uint8, bufferCapacity int) (CFWriteStreamRef, error) {
	if _cFWriteStreamCreateWithBuffer == nil {
		return 0, symbolCallError("CFWriteStreamCreateWithBuffer", "", _cFWriteStreamCreateWithBufferErr)
	}
	return _cFWriteStreamCreateWithBuffer(alloc, buffer, bufferCapacity), nil
}

// CFWriteStreamCreateWithBuffer creates a writable stream for a fixed-size block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithBuffer(_:_:_:)
func CFWriteStreamCreateWithBuffer(alloc CFAllocatorRef, buffer *uint8, bufferCapacity int) CFWriteStreamRef {
	result, callErr := tryCFWriteStreamCreateWithBuffer(alloc, buffer, bufferCapacity)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamCreateWithFile func(alloc CFAllocatorRef, fileURL CFURLRef) CFWriteStreamRef
var _cFWriteStreamCreateWithFileErr error

func tryCFWriteStreamCreateWithFile(alloc CFAllocatorRef, fileURL CFURLRef) (CFWriteStreamRef, error) {
	if _cFWriteStreamCreateWithFile == nil {
		return 0, symbolCallError("CFWriteStreamCreateWithFile", "", _cFWriteStreamCreateWithFileErr)
	}
	return _cFWriteStreamCreateWithFile(alloc, fileURL), nil
}

// CFWriteStreamCreateWithFile creates a writable stream for a file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithFile(_:_:)
func CFWriteStreamCreateWithFile(alloc CFAllocatorRef, fileURL CFURLRef) CFWriteStreamRef {
	result, callErr := tryCFWriteStreamCreateWithFile(alloc, fileURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamGetError func(stream CFWriteStreamRef) CFStreamError
var _cFWriteStreamGetErrorErr error

func tryCFWriteStreamGetError(stream CFWriteStreamRef) (CFStreamError, error) {
	if _cFWriteStreamGetError == nil {
		return CFStreamError{}, symbolCallError("CFWriteStreamGetError", "", _cFWriteStreamGetErrorErr)
	}
	return _cFWriteStreamGetError(stream), nil
}

// CFWriteStreamGetError returns the error status of a stream.
//
// Deprecated: Use [CFWriteStreamCopyError(_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFWriteStreamCopyError(_:)>) instead.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetError(_:)
func CFWriteStreamGetError(stream CFWriteStreamRef) CFStreamError {
	result, callErr := tryCFWriteStreamGetError(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamGetStatus func(stream CFWriteStreamRef) CFStreamStatus
var _cFWriteStreamGetStatusErr error

func tryCFWriteStreamGetStatus(stream CFWriteStreamRef) (CFStreamStatus, error) {
	if _cFWriteStreamGetStatus == nil {
		return *new(CFStreamStatus), symbolCallError("CFWriteStreamGetStatus", "", _cFWriteStreamGetStatusErr)
	}
	return _cFWriteStreamGetStatus(stream), nil
}

// CFWriteStreamGetStatus returns the current state of a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetStatus(_:)
func CFWriteStreamGetStatus(stream CFWriteStreamRef) CFStreamStatus {
	result, callErr := tryCFWriteStreamGetStatus(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamGetTypeID func() uint
var _cFWriteStreamGetTypeIDErr error

func tryCFWriteStreamGetTypeID() (uint, error) {
	if _cFWriteStreamGetTypeID == nil {
		return 0, symbolCallError("CFWriteStreamGetTypeID", "", _cFWriteStreamGetTypeIDErr)
	}
	return _cFWriteStreamGetTypeID(), nil
}

// CFWriteStreamGetTypeID returns the type identifier of all CFWriteStream objects.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetTypeID()
func CFWriteStreamGetTypeID() uint {
	result, callErr := tryCFWriteStreamGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamOpen func(stream CFWriteStreamRef) bool
var _cFWriteStreamOpenErr error

func tryCFWriteStreamOpen(stream CFWriteStreamRef) (bool, error) {
	if _cFWriteStreamOpen == nil {
		return false, symbolCallError("CFWriteStreamOpen", "", _cFWriteStreamOpenErr)
	}
	return _cFWriteStreamOpen(stream), nil
}

// CFWriteStreamOpen opens a stream for writing.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamOpen(_:)
func CFWriteStreamOpen(stream CFWriteStreamRef) bool {
	result, callErr := tryCFWriteStreamOpen(stream)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamScheduleWithRunLoop func(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)
var _cFWriteStreamScheduleWithRunLoopErr error

func tryCFWriteStreamScheduleWithRunLoop(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) error {
	if _cFWriteStreamScheduleWithRunLoop == nil {
		return symbolCallError("CFWriteStreamScheduleWithRunLoop", "", _cFWriteStreamScheduleWithRunLoopErr)
	}
	_cFWriteStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
	return nil
}

// CFWriteStreamScheduleWithRunLoop schedules a stream into a run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamScheduleWithRunLoop(_:_:_:)
func CFWriteStreamScheduleWithRunLoop(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) {
	if callErr := tryCFWriteStreamScheduleWithRunLoop(stream, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _cFWriteStreamSetClient func(stream CFWriteStreamRef, streamEvents uint64, clientCB CFWriteStreamClientCallBack, clientContext *CFStreamClientContext) bool
var _cFWriteStreamSetClientErr error

func tryCFWriteStreamSetClient(stream CFWriteStreamRef, streamEvents uint64, clientCB CFWriteStreamClientCallBack, clientContext *CFStreamClientContext) (bool, error) {
	if _cFWriteStreamSetClient == nil {
		return false, symbolCallError("CFWriteStreamSetClient", "", _cFWriteStreamSetClientErr)
	}
	return _cFWriteStreamSetClient(stream, streamEvents, clientCB, clientContext), nil
}

// CFWriteStreamSetClient assigns a client to a stream, which receives callbacks when certain events occur.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetClient(_:_:_:_:)
func CFWriteStreamSetClient(stream CFWriteStreamRef, streamEvents uint64, clientCB CFWriteStreamClientCallBack, clientContext *CFStreamClientContext) bool {
	result, callErr := tryCFWriteStreamSetClient(stream, streamEvents, clientCB, clientContext)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamSetDispatchQueue func(stream CFWriteStreamRef, q uintptr)
var _cFWriteStreamSetDispatchQueueErr error

func tryCFWriteStreamSetDispatchQueue(stream CFWriteStreamRef, q dispatch.Queue) error {
	if _cFWriteStreamSetDispatchQueue == nil {
		return symbolCallError("CFWriteStreamSetDispatchQueue", "10.9", _cFWriteStreamSetDispatchQueueErr)
	}
	_cFWriteStreamSetDispatchQueue(stream, uintptr(q.Handle()))
	return nil
}

// CFWriteStreamSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetDispatchQueue(_:_:)
func CFWriteStreamSetDispatchQueue(stream CFWriteStreamRef, q dispatch.Queue) {
	if callErr := tryCFWriteStreamSetDispatchQueue(stream, q); callErr != nil {
		panic(callErr)
	}
}

var _cFWriteStreamSetProperty func(stream CFWriteStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool
var _cFWriteStreamSetPropertyErr error

func tryCFWriteStreamSetProperty(stream CFWriteStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) (bool, error) {
	if _cFWriteStreamSetProperty == nil {
		return false, symbolCallError("CFWriteStreamSetProperty", "", _cFWriteStreamSetPropertyErr)
	}
	return _cFWriteStreamSetProperty(stream, propertyName, propertyValue), nil
}

// CFWriteStreamSetProperty sets the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetProperty(_:_:_:)
func CFWriteStreamSetProperty(stream CFWriteStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool {
	result, callErr := tryCFWriteStreamSetProperty(stream, propertyName, propertyValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFWriteStreamUnscheduleFromRunLoop func(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)
var _cFWriteStreamUnscheduleFromRunLoopErr error

func tryCFWriteStreamUnscheduleFromRunLoop(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) error {
	if _cFWriteStreamUnscheduleFromRunLoop == nil {
		return symbolCallError("CFWriteStreamUnscheduleFromRunLoop", "", _cFWriteStreamUnscheduleFromRunLoopErr)
	}
	_cFWriteStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
	return nil
}

// CFWriteStreamUnscheduleFromRunLoop removes a stream from a particular run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamUnscheduleFromRunLoop(_:_:_:)
func CFWriteStreamUnscheduleFromRunLoop(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) {
	if callErr := tryCFWriteStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _cFWriteStreamWrite func(stream CFWriteStreamRef, buffer *uint8, bufferLength int) int
var _cFWriteStreamWriteErr error

func tryCFWriteStreamWrite(stream CFWriteStreamRef, buffer *uint8, bufferLength int) (int, error) {
	if _cFWriteStreamWrite == nil {
		return 0, symbolCallError("CFWriteStreamWrite", "", _cFWriteStreamWriteErr)
	}
	return _cFWriteStreamWrite(stream, buffer, bufferLength), nil
}

// CFWriteStreamWrite writes data to a writable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamWrite(_:_:_:)
func CFWriteStreamWrite(stream CFWriteStreamRef, buffer *uint8, bufferLength int) int {
	result, callErr := tryCFWriteStreamWrite(stream, buffer, bufferLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLCreateStringByEscapingEntities func(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef
var _cFXMLCreateStringByEscapingEntitiesErr error

func tryCFXMLCreateStringByEscapingEntities(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) (CFStringRef, error) {
	if _cFXMLCreateStringByEscapingEntities == nil {
		return 0, symbolCallError("CFXMLCreateStringByEscapingEntities", "", _cFXMLCreateStringByEscapingEntitiesErr)
	}
	return _cFXMLCreateStringByEscapingEntities(allocator, string_, entitiesDictionary), nil
}

// CFXMLCreateStringByEscapingEntities given a CFString object containing XML source with unescaped entities, returns a string with specified XML entities escaped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByEscapingEntities(_:_:_:)
func CFXMLCreateStringByEscapingEntities(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef {
	result, callErr := tryCFXMLCreateStringByEscapingEntities(allocator, string_, entitiesDictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLCreateStringByUnescapingEntities func(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef
var _cFXMLCreateStringByUnescapingEntitiesErr error

func tryCFXMLCreateStringByUnescapingEntities(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) (CFStringRef, error) {
	if _cFXMLCreateStringByUnescapingEntities == nil {
		return 0, symbolCallError("CFXMLCreateStringByUnescapingEntities", "", _cFXMLCreateStringByUnescapingEntitiesErr)
	}
	return _cFXMLCreateStringByUnescapingEntities(allocator, string_, entitiesDictionary), nil
}

// CFXMLCreateStringByUnescapingEntities given a CFString object containing XML source with escaped entities, returns a string with specified XML entities unescaped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByUnescapingEntities(_:_:_:)
func CFXMLCreateStringByUnescapingEntities(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef {
	result, callErr := tryCFXMLCreateStringByUnescapingEntities(allocator, string_, entitiesDictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeCreate func(alloc CFAllocatorRef, xmlType uint32, dataString CFStringRef, additionalInfoPtr unsafe.Pointer, version int) unsafe.Pointer
var _cFXMLNodeCreateErr error

func tryCFXMLNodeCreate(alloc CFAllocatorRef, xmlType uint32, dataString CFStringRef, additionalInfoPtr unsafe.Pointer, version int) (unsafe.Pointer, error) {
	if _cFXMLNodeCreate == nil {
		return nil, symbolCallError("CFXMLNodeCreate", "10.0", _cFXMLNodeCreateErr)
	}
	return _cFXMLNodeCreate(alloc, xmlType, dataString, additionalInfoPtr, version), nil
}

// CFXMLNodeCreate creates a new CFXMLNode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreate
func CFXMLNodeCreate(alloc CFAllocatorRef, xmlType uint32, dataString CFStringRef, additionalInfoPtr unsafe.Pointer, version int) unsafe.Pointer {
	result, callErr := tryCFXMLNodeCreate(alloc, xmlType, dataString, additionalInfoPtr, version)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeCreateCopy func(alloc CFAllocatorRef, origNode unsafe.Pointer) unsafe.Pointer
var _cFXMLNodeCreateCopyErr error

func tryCFXMLNodeCreateCopy(alloc CFAllocatorRef, origNode unsafe.Pointer) (unsafe.Pointer, error) {
	if _cFXMLNodeCreateCopy == nil {
		return nil, symbolCallError("CFXMLNodeCreateCopy", "10.0", _cFXMLNodeCreateCopyErr)
	}
	return _cFXMLNodeCreateCopy(alloc, origNode), nil
}

// CFXMLNodeCreateCopy creates a copy of a CFXMLNode object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreateCopy
func CFXMLNodeCreateCopy(alloc CFAllocatorRef, origNode unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryCFXMLNodeCreateCopy(alloc, origNode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeGetInfoPtr func(node unsafe.Pointer) unsafe.Pointer
var _cFXMLNodeGetInfoPtrErr error

func tryCFXMLNodeGetInfoPtr(node unsafe.Pointer) (unsafe.Pointer, error) {
	if _cFXMLNodeGetInfoPtr == nil {
		return nil, symbolCallError("CFXMLNodeGetInfoPtr", "10.0", _cFXMLNodeGetInfoPtrErr)
	}
	return _cFXMLNodeGetInfoPtr(node), nil
}

// CFXMLNodeGetInfoPtr returns the additional information pointer of a CFXMLNode object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetInfoPtr
func CFXMLNodeGetInfoPtr(node unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryCFXMLNodeGetInfoPtr(node)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeGetString func(node unsafe.Pointer) CFStringRef
var _cFXMLNodeGetStringErr error

func tryCFXMLNodeGetString(node unsafe.Pointer) (CFStringRef, error) {
	if _cFXMLNodeGetString == nil {
		return 0, symbolCallError("CFXMLNodeGetString", "10.0", _cFXMLNodeGetStringErr)
	}
	return _cFXMLNodeGetString(node), nil
}

// CFXMLNodeGetString returns the data string from a CFXMLNode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetString
func CFXMLNodeGetString(node unsafe.Pointer) CFStringRef {
	result, callErr := tryCFXMLNodeGetString(node)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeGetTypeCode func(node unsafe.Pointer) uint32
var _cFXMLNodeGetTypeCodeErr error

func tryCFXMLNodeGetTypeCode(node unsafe.Pointer) (uint32, error) {
	if _cFXMLNodeGetTypeCode == nil {
		return 0, symbolCallError("CFXMLNodeGetTypeCode", "10.0", _cFXMLNodeGetTypeCodeErr)
	}
	return _cFXMLNodeGetTypeCode(node), nil
}

// CFXMLNodeGetTypeCode returns the XML structure type code for a CFXMLNode object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeCode
func CFXMLNodeGetTypeCode(node unsafe.Pointer) uint32 {
	result, callErr := tryCFXMLNodeGetTypeCode(node)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeGetTypeID func() uint
var _cFXMLNodeGetTypeIDErr error

func tryCFXMLNodeGetTypeID() (uint, error) {
	if _cFXMLNodeGetTypeID == nil {
		return 0, symbolCallError("CFXMLNodeGetTypeID", "10.0", _cFXMLNodeGetTypeIDErr)
	}
	return _cFXMLNodeGetTypeID(), nil
}

// CFXMLNodeGetTypeID returns the type identifier code for the CFXMLNode opaque type.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeID
func CFXMLNodeGetTypeID() uint {
	result, callErr := tryCFXMLNodeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLNodeGetVersion func(node unsafe.Pointer) int
var _cFXMLNodeGetVersionErr error

func tryCFXMLNodeGetVersion(node unsafe.Pointer) (int, error) {
	if _cFXMLNodeGetVersion == nil {
		return 0, symbolCallError("CFXMLNodeGetVersion", "10.0", _cFXMLNodeGetVersionErr)
	}
	return _cFXMLNodeGetVersion(node), nil
}

// CFXMLNodeGetVersion returns the version number for a CFXMLNode object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetVersion
func CFXMLNodeGetVersion(node unsafe.Pointer) int {
	result, callErr := tryCFXMLNodeGetVersion(node)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserAbort func(parser CFXMLParserRef, errorCode CFXMLParserStatusCode, errorDescription CFStringRef)
var _cFXMLParserAbortErr error

func tryCFXMLParserAbort(parser CFXMLParserRef, errorCode CFXMLParserStatusCode, errorDescription CFStringRef) error {
	if _cFXMLParserAbort == nil {
		return symbolCallError("CFXMLParserAbort", "10.0", _cFXMLParserAbortErr)
	}
	_cFXMLParserAbort(parser, errorCode, errorDescription)
	return nil
}

// CFXMLParserAbort causes a parser to abort with the given error code and description.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserAbort
func CFXMLParserAbort(parser CFXMLParserRef, errorCode CFXMLParserStatusCode, errorDescription CFStringRef) {
	if callErr := tryCFXMLParserAbort(parser, errorCode, errorDescription); callErr != nil {
		panic(callErr)
	}
}

var _cFXMLParserCopyErrorDescription func(parser CFXMLParserRef) CFStringRef
var _cFXMLParserCopyErrorDescriptionErr error

func tryCFXMLParserCopyErrorDescription(parser CFXMLParserRef) (CFStringRef, error) {
	if _cFXMLParserCopyErrorDescription == nil {
		return 0, symbolCallError("CFXMLParserCopyErrorDescription", "10.0", _cFXMLParserCopyErrorDescriptionErr)
	}
	return _cFXMLParserCopyErrorDescription(parser), nil
}

// CFXMLParserCopyErrorDescription returns the user-readable description of the current error condition.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCopyErrorDescription
func CFXMLParserCopyErrorDescription(parser CFXMLParserRef) CFStringRef {
	result, callErr := tryCFXMLParserCopyErrorDescription(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserCreate func(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef
var _cFXMLParserCreateErr error

func tryCFXMLParserCreate(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) (CFXMLParserRef, error) {
	if _cFXMLParserCreate == nil {
		return 0, symbolCallError("CFXMLParserCreate", "10.0", _cFXMLParserCreateErr)
	}
	return _cFXMLParserCreate(allocator, xmlData, dataSource, parseOptions, versionOfNodes, callBacks, context), nil
}

// CFXMLParserCreate creates a new XML parser for the specified XML data.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreate
func CFXMLParserCreate(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef {
	result, callErr := tryCFXMLParserCreate(allocator, xmlData, dataSource, parseOptions, versionOfNodes, callBacks, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserCreateWithDataFromURL func(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef
var _cFXMLParserCreateWithDataFromURLErr error

func tryCFXMLParserCreateWithDataFromURL(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) (CFXMLParserRef, error) {
	if _cFXMLParserCreateWithDataFromURL == nil {
		return 0, symbolCallError("CFXMLParserCreateWithDataFromURL", "10.0", _cFXMLParserCreateWithDataFromURLErr)
	}
	return _cFXMLParserCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes, callBacks, context), nil
}

// CFXMLParserCreateWithDataFromURL creates a new XML parser for the specified XML data at the specified URL.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreateWithDataFromURL
func CFXMLParserCreateWithDataFromURL(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef {
	result, callErr := tryCFXMLParserCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes, callBacks, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserGetCallBacks func(parser CFXMLParserRef, callBacks *CFXMLParserCallBacks)
var _cFXMLParserGetCallBacksErr error

func tryCFXMLParserGetCallBacks(parser CFXMLParserRef, callBacks *CFXMLParserCallBacks) error {
	if _cFXMLParserGetCallBacks == nil {
		return symbolCallError("CFXMLParserGetCallBacks", "10.0", _cFXMLParserGetCallBacksErr)
	}
	_cFXMLParserGetCallBacks(parser, callBacks)
	return nil
}

// CFXMLParserGetCallBacks returns the callbacks associated with an XML parser when it was created.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetCallBacks
func CFXMLParserGetCallBacks(parser CFXMLParserRef, callBacks *CFXMLParserCallBacks) {
	if callErr := tryCFXMLParserGetCallBacks(parser, callBacks); callErr != nil {
		panic(callErr)
	}
}

var _cFXMLParserGetContext func(parser CFXMLParserRef, context *CFXMLParserContext)
var _cFXMLParserGetContextErr error

func tryCFXMLParserGetContext(parser CFXMLParserRef, context *CFXMLParserContext) error {
	if _cFXMLParserGetContext == nil {
		return symbolCallError("CFXMLParserGetContext", "10.0", _cFXMLParserGetContextErr)
	}
	_cFXMLParserGetContext(parser, context)
	return nil
}

// CFXMLParserGetContext returns the context for an XML parser.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetContext
func CFXMLParserGetContext(parser CFXMLParserRef, context *CFXMLParserContext) {
	if callErr := tryCFXMLParserGetContext(parser, context); callErr != nil {
		panic(callErr)
	}
}

var _cFXMLParserGetDocument func(parser CFXMLParserRef) unsafe.Pointer
var _cFXMLParserGetDocumentErr error

func tryCFXMLParserGetDocument(parser CFXMLParserRef) (unsafe.Pointer, error) {
	if _cFXMLParserGetDocument == nil {
		return nil, symbolCallError("CFXMLParserGetDocument", "10.0", _cFXMLParserGetDocumentErr)
	}
	return _cFXMLParserGetDocument(parser), nil
}

// CFXMLParserGetDocument returns the top-most object returned by the create XML structure callback.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetDocument
func CFXMLParserGetDocument(parser CFXMLParserRef) unsafe.Pointer {
	result, callErr := tryCFXMLParserGetDocument(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserGetLineNumber func(parser CFXMLParserRef) int
var _cFXMLParserGetLineNumberErr error

func tryCFXMLParserGetLineNumber(parser CFXMLParserRef) (int, error) {
	if _cFXMLParserGetLineNumber == nil {
		return 0, symbolCallError("CFXMLParserGetLineNumber", "10.0", _cFXMLParserGetLineNumberErr)
	}
	return _cFXMLParserGetLineNumber(parser), nil
}

// CFXMLParserGetLineNumber returns the line number of the current parse location.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLineNumber
func CFXMLParserGetLineNumber(parser CFXMLParserRef) int {
	result, callErr := tryCFXMLParserGetLineNumber(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserGetLocation func(parser CFXMLParserRef) int
var _cFXMLParserGetLocationErr error

func tryCFXMLParserGetLocation(parser CFXMLParserRef) (int, error) {
	if _cFXMLParserGetLocation == nil {
		return 0, symbolCallError("CFXMLParserGetLocation", "10.0", _cFXMLParserGetLocationErr)
	}
	return _cFXMLParserGetLocation(parser), nil
}

// CFXMLParserGetLocation returns the character index of the current parse location.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLocation
func CFXMLParserGetLocation(parser CFXMLParserRef) int {
	result, callErr := tryCFXMLParserGetLocation(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserGetSourceURL func(parser CFXMLParserRef) CFURLRef
var _cFXMLParserGetSourceURLErr error

func tryCFXMLParserGetSourceURL(parser CFXMLParserRef) (CFURLRef, error) {
	if _cFXMLParserGetSourceURL == nil {
		return 0, symbolCallError("CFXMLParserGetSourceURL", "10.0", _cFXMLParserGetSourceURLErr)
	}
	return _cFXMLParserGetSourceURL(parser), nil
}

// CFXMLParserGetSourceURL returns the URL for the XML data being parsed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetSourceURL
func CFXMLParserGetSourceURL(parser CFXMLParserRef) CFURLRef {
	result, callErr := tryCFXMLParserGetSourceURL(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserGetStatusCode func(parser CFXMLParserRef) CFXMLParserStatusCode
var _cFXMLParserGetStatusCodeErr error

func tryCFXMLParserGetStatusCode(parser CFXMLParserRef) (CFXMLParserStatusCode, error) {
	if _cFXMLParserGetStatusCode == nil {
		return *new(CFXMLParserStatusCode), symbolCallError("CFXMLParserGetStatusCode", "10.0", _cFXMLParserGetStatusCodeErr)
	}
	return _cFXMLParserGetStatusCode(parser), nil
}

// CFXMLParserGetStatusCode returns a numeric code indicating the current status of the parser.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetStatusCode
func CFXMLParserGetStatusCode(parser CFXMLParserRef) CFXMLParserStatusCode {
	result, callErr := tryCFXMLParserGetStatusCode(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserGetTypeID func() uint
var _cFXMLParserGetTypeIDErr error

func tryCFXMLParserGetTypeID() (uint, error) {
	if _cFXMLParserGetTypeID == nil {
		return 0, symbolCallError("CFXMLParserGetTypeID", "10.0", _cFXMLParserGetTypeIDErr)
	}
	return _cFXMLParserGetTypeID(), nil
}

// CFXMLParserGetTypeID returns the type identifier for the CFXMLParser opaque type.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetTypeID
func CFXMLParserGetTypeID() uint {
	result, callErr := tryCFXMLParserGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLParserParse func(parser CFXMLParserRef) bool
var _cFXMLParserParseErr error

func tryCFXMLParserParse(parser CFXMLParserRef) (bool, error) {
	if _cFXMLParserParse == nil {
		return false, symbolCallError("CFXMLParserParse", "10.0", _cFXMLParserParseErr)
	}
	return _cFXMLParserParse(parser), nil
}

// CFXMLParserParse begins a parse of the XML data that was associated with the parser when it was created.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserParse
func CFXMLParserParse(parser CFXMLParserRef) bool {
	result, callErr := tryCFXMLParserParse(parser)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLTreeCreateFromData func(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef
var _cFXMLTreeCreateFromDataErr error

func tryCFXMLTreeCreateFromData(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) (CFXMLTreeRef, error) {
	if _cFXMLTreeCreateFromData == nil {
		return 0, symbolCallError("CFXMLTreeCreateFromData", "10.0", _cFXMLTreeCreateFromDataErr)
	}
	return _cFXMLTreeCreateFromData(allocator, xmlData, dataSource, parseOptions, versionOfNodes), nil
}

// CFXMLTreeCreateFromData parses the given XML data and returns the resulting CFXMLTree object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromData
func CFXMLTreeCreateFromData(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef {
	result, callErr := tryCFXMLTreeCreateFromData(allocator, xmlData, dataSource, parseOptions, versionOfNodes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLTreeCreateFromDataWithError func(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, errorDict *CFDictionaryRef) CFXMLTreeRef
var _cFXMLTreeCreateFromDataWithErrorErr error

func tryCFXMLTreeCreateFromDataWithError(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, errorDict *CFDictionaryRef) (CFXMLTreeRef, error) {
	if _cFXMLTreeCreateFromDataWithError == nil {
		return 0, symbolCallError("CFXMLTreeCreateFromDataWithError", "10.0", _cFXMLTreeCreateFromDataWithErrorErr)
	}
	return _cFXMLTreeCreateFromDataWithError(allocator, xmlData, dataSource, parseOptions, versionOfNodes, errorDict), nil
}

// CFXMLTreeCreateFromDataWithError parses the given XML data and returns the resulting CFXMLTree object and any error information.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromDataWithError
func CFXMLTreeCreateFromDataWithError(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, errorDict *CFDictionaryRef) CFXMLTreeRef {
	result, callErr := tryCFXMLTreeCreateFromDataWithError(allocator, xmlData, dataSource, parseOptions, versionOfNodes, errorDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLTreeCreateWithDataFromURL func(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef
var _cFXMLTreeCreateWithDataFromURLErr error

func tryCFXMLTreeCreateWithDataFromURL(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) (CFXMLTreeRef, error) {
	if _cFXMLTreeCreateWithDataFromURL == nil {
		return 0, symbolCallError("CFXMLTreeCreateWithDataFromURL", "10.0", _cFXMLTreeCreateWithDataFromURLErr)
	}
	return _cFXMLTreeCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes), nil
}

// CFXMLTreeCreateWithDataFromURL creates a new CFXMLTree object by loading the data to be parsed directly from a data source.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithDataFromURL
func CFXMLTreeCreateWithDataFromURL(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef {
	result, callErr := tryCFXMLTreeCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLTreeCreateWithNode func(allocator CFAllocatorRef, node unsafe.Pointer) CFXMLTreeRef
var _cFXMLTreeCreateWithNodeErr error

func tryCFXMLTreeCreateWithNode(allocator CFAllocatorRef, node unsafe.Pointer) (CFXMLTreeRef, error) {
	if _cFXMLTreeCreateWithNode == nil {
		return 0, symbolCallError("CFXMLTreeCreateWithNode", "10.0", _cFXMLTreeCreateWithNodeErr)
	}
	return _cFXMLTreeCreateWithNode(allocator, node), nil
}

// CFXMLTreeCreateWithNode creates a childless, parentless CFXMLTree object node for a CFXMLNode object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithNode
func CFXMLTreeCreateWithNode(allocator CFAllocatorRef, node unsafe.Pointer) CFXMLTreeRef {
	result, callErr := tryCFXMLTreeCreateWithNode(allocator, node)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLTreeCreateXMLData func(allocator CFAllocatorRef, xmlTree CFXMLTreeRef) CFDataRef
var _cFXMLTreeCreateXMLDataErr error

func tryCFXMLTreeCreateXMLData(allocator CFAllocatorRef, xmlTree CFXMLTreeRef) (CFDataRef, error) {
	if _cFXMLTreeCreateXMLData == nil {
		return 0, symbolCallError("CFXMLTreeCreateXMLData", "10.0", _cFXMLTreeCreateXMLDataErr)
	}
	return _cFXMLTreeCreateXMLData(allocator, xmlTree), nil
}

// CFXMLTreeCreateXMLData generates an XML document from a CFXMLTree object which is ready to be written to permanent storage.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateXMLData
func CFXMLTreeCreateXMLData(allocator CFAllocatorRef, xmlTree CFXMLTreeRef) CFDataRef {
	result, callErr := tryCFXMLTreeCreateXMLData(allocator, xmlTree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cFXMLTreeGetNode func(xmlTree CFXMLTreeRef) unsafe.Pointer
var _cFXMLTreeGetNodeErr error

func tryCFXMLTreeGetNode(xmlTree CFXMLTreeRef) (unsafe.Pointer, error) {
	if _cFXMLTreeGetNode == nil {
		return nil, symbolCallError("CFXMLTreeGetNode", "10.0", _cFXMLTreeGetNodeErr)
	}
	return _cFXMLTreeGetNode(xmlTree), nil
}

// CFXMLTreeGetNode returns the node of a CFXMLTree object.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeGetNode
func CFXMLTreeGetNode(xmlTree CFXMLTreeRef) unsafe.Pointer {
	result, callErr := tryCFXMLTreeGetNode(xmlTree)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cFAbsoluteTimeGetCurrent, &_cFAbsoluteTimeGetCurrentErr, frameworkHandle, "CFAbsoluteTimeGetCurrent", "")
	registerFunc(&_cFAllocatorAllocate, &_cFAllocatorAllocateErr, frameworkHandle, "CFAllocatorAllocate", "")
	registerFunc(&_cFAllocatorAllocateBytes, &_cFAllocatorAllocateBytesErr, frameworkHandle, "CFAllocatorAllocateBytes", "15.0")
	registerFunc(&_cFAllocatorAllocateTyped, &_cFAllocatorAllocateTypedErr, frameworkHandle, "CFAllocatorAllocateTyped", "15.0")
	registerFunc(&_cFAllocatorCreate, &_cFAllocatorCreateErr, frameworkHandle, "CFAllocatorCreate", "")
	registerFunc(&_cFAllocatorCreateWithZone, &_cFAllocatorCreateWithZoneErr, frameworkHandle, "CFAllocatorCreateWithZone", "")
	registerFunc(&_cFAllocatorDeallocate, &_cFAllocatorDeallocateErr, frameworkHandle, "CFAllocatorDeallocate", "")
	registerFunc(&_cFAllocatorGetContext, &_cFAllocatorGetContextErr, frameworkHandle, "CFAllocatorGetContext", "")
	registerFunc(&_cFAllocatorGetDefault, &_cFAllocatorGetDefaultErr, frameworkHandle, "CFAllocatorGetDefault", "")
	registerFunc(&_cFAllocatorGetPreferredSizeForSize, &_cFAllocatorGetPreferredSizeForSizeErr, frameworkHandle, "CFAllocatorGetPreferredSizeForSize", "")
	registerFunc(&_cFAllocatorGetTypeID, &_cFAllocatorGetTypeIDErr, frameworkHandle, "CFAllocatorGetTypeID", "")
	registerFunc(&_cFAllocatorReallocate, &_cFAllocatorReallocateErr, frameworkHandle, "CFAllocatorReallocate", "")
	registerFunc(&_cFAllocatorReallocateBytes, &_cFAllocatorReallocateBytesErr, frameworkHandle, "CFAllocatorReallocateBytes", "15.0")
	registerFunc(&_cFAllocatorReallocateTyped, &_cFAllocatorReallocateTypedErr, frameworkHandle, "CFAllocatorReallocateTyped", "15.0")
	registerFunc(&_cFAllocatorSetDefault, &_cFAllocatorSetDefaultErr, frameworkHandle, "CFAllocatorSetDefault", "")
	registerFunc(&_cFArrayAppendArray, &_cFArrayAppendArrayErr, frameworkHandle, "CFArrayAppendArray", "")
	registerFunc(&_cFArrayAppendValue, &_cFArrayAppendValueErr, frameworkHandle, "CFArrayAppendValue", "")
	registerFunc(&_cFArrayApplyFunction, &_cFArrayApplyFunctionErr, frameworkHandle, "CFArrayApplyFunction", "")
	registerFunc(&_cFArrayBSearchValues, &_cFArrayBSearchValuesErr, frameworkHandle, "CFArrayBSearchValues", "")
	registerFunc(&_cFArrayContainsValue, &_cFArrayContainsValueErr, frameworkHandle, "CFArrayContainsValue", "")
	registerFunc(&_cFArrayCreate, &_cFArrayCreateErr, frameworkHandle, "CFArrayCreate", "")
	registerFunc(&_cFArrayCreateCopy, &_cFArrayCreateCopyErr, frameworkHandle, "CFArrayCreateCopy", "")
	registerFunc(&_cFArrayCreateMutable, &_cFArrayCreateMutableErr, frameworkHandle, "CFArrayCreateMutable", "")
	registerFunc(&_cFArrayCreateMutableCopy, &_cFArrayCreateMutableCopyErr, frameworkHandle, "CFArrayCreateMutableCopy", "")
	registerFunc(&_cFArrayExchangeValuesAtIndices, &_cFArrayExchangeValuesAtIndicesErr, frameworkHandle, "CFArrayExchangeValuesAtIndices", "")
	registerFunc(&_cFArrayGetCount, &_cFArrayGetCountErr, frameworkHandle, "CFArrayGetCount", "")
	registerFunc(&_cFArrayGetCountOfValue, &_cFArrayGetCountOfValueErr, frameworkHandle, "CFArrayGetCountOfValue", "")
	registerFunc(&_cFArrayGetFirstIndexOfValue, &_cFArrayGetFirstIndexOfValueErr, frameworkHandle, "CFArrayGetFirstIndexOfValue", "")
	registerFunc(&_cFArrayGetLastIndexOfValue, &_cFArrayGetLastIndexOfValueErr, frameworkHandle, "CFArrayGetLastIndexOfValue", "")
	registerFunc(&_cFArrayGetTypeID, &_cFArrayGetTypeIDErr, frameworkHandle, "CFArrayGetTypeID", "")
	registerFunc(&_cFArrayGetValueAtIndex, &_cFArrayGetValueAtIndexErr, frameworkHandle, "CFArrayGetValueAtIndex", "")
	registerFunc(&_cFArrayGetValues, &_cFArrayGetValuesErr, frameworkHandle, "CFArrayGetValues", "")
	registerFunc(&_cFArrayInsertValueAtIndex, &_cFArrayInsertValueAtIndexErr, frameworkHandle, "CFArrayInsertValueAtIndex", "")
	registerFunc(&_cFArrayRemoveAllValues, &_cFArrayRemoveAllValuesErr, frameworkHandle, "CFArrayRemoveAllValues", "")
	registerFunc(&_cFArrayRemoveValueAtIndex, &_cFArrayRemoveValueAtIndexErr, frameworkHandle, "CFArrayRemoveValueAtIndex", "")
	registerFunc(&_cFArrayReplaceValues, &_cFArrayReplaceValuesErr, frameworkHandle, "CFArrayReplaceValues", "")
	registerFunc(&_cFArraySetValueAtIndex, &_cFArraySetValueAtIndexErr, frameworkHandle, "CFArraySetValueAtIndex", "")
	registerFunc(&_cFArraySortValues, &_cFArraySortValuesErr, frameworkHandle, "CFArraySortValues", "")
	registerFunc(&_cFAttributedStringBeginEditing, &_cFAttributedStringBeginEditingErr, frameworkHandle, "CFAttributedStringBeginEditing", "")
	registerFunc(&_cFAttributedStringCreate, &_cFAttributedStringCreateErr, frameworkHandle, "CFAttributedStringCreate", "")
	registerFunc(&_cFAttributedStringCreateCopy, &_cFAttributedStringCreateCopyErr, frameworkHandle, "CFAttributedStringCreateCopy", "")
	registerFunc(&_cFAttributedStringCreateMutable, &_cFAttributedStringCreateMutableErr, frameworkHandle, "CFAttributedStringCreateMutable", "")
	registerFunc(&_cFAttributedStringCreateMutableCopy, &_cFAttributedStringCreateMutableCopyErr, frameworkHandle, "CFAttributedStringCreateMutableCopy", "")
	registerFunc(&_cFAttributedStringCreateWithSubstring, &_cFAttributedStringCreateWithSubstringErr, frameworkHandle, "CFAttributedStringCreateWithSubstring", "")
	registerFunc(&_cFAttributedStringEndEditing, &_cFAttributedStringEndEditingErr, frameworkHandle, "CFAttributedStringEndEditing", "")
	registerFunc(&_cFAttributedStringGetAttribute, &_cFAttributedStringGetAttributeErr, frameworkHandle, "CFAttributedStringGetAttribute", "")
	registerFunc(&_cFAttributedStringGetAttributeAndLongestEffectiveRange, &_cFAttributedStringGetAttributeAndLongestEffectiveRangeErr, frameworkHandle, "CFAttributedStringGetAttributeAndLongestEffectiveRange", "")
	registerFunc(&_cFAttributedStringGetAttributes, &_cFAttributedStringGetAttributesErr, frameworkHandle, "CFAttributedStringGetAttributes", "")
	registerFunc(&_cFAttributedStringGetAttributesAndLongestEffectiveRange, &_cFAttributedStringGetAttributesAndLongestEffectiveRangeErr, frameworkHandle, "CFAttributedStringGetAttributesAndLongestEffectiveRange", "")
	registerFunc(&_cFAttributedStringGetBidiLevelsAndResolvedDirections, &_cFAttributedStringGetBidiLevelsAndResolvedDirectionsErr, frameworkHandle, "CFAttributedStringGetBidiLevelsAndResolvedDirections", "")
	registerFunc(&_cFAttributedStringGetLength, &_cFAttributedStringGetLengthErr, frameworkHandle, "CFAttributedStringGetLength", "")
	registerFunc(&_cFAttributedStringGetMutableString, &_cFAttributedStringGetMutableStringErr, frameworkHandle, "CFAttributedStringGetMutableString", "")
	registerFunc(&_cFAttributedStringGetStatisticalWritingDirections, &_cFAttributedStringGetStatisticalWritingDirectionsErr, frameworkHandle, "CFAttributedStringGetStatisticalWritingDirections", "26.0")
	registerFunc(&_cFAttributedStringGetString, &_cFAttributedStringGetStringErr, frameworkHandle, "CFAttributedStringGetString", "")
	registerFunc(&_cFAttributedStringGetTypeID, &_cFAttributedStringGetTypeIDErr, frameworkHandle, "CFAttributedStringGetTypeID", "")
	registerFunc(&_cFAttributedStringRemoveAttribute, &_cFAttributedStringRemoveAttributeErr, frameworkHandle, "CFAttributedStringRemoveAttribute", "")
	registerFunc(&_cFAttributedStringReplaceAttributedString, &_cFAttributedStringReplaceAttributedStringErr, frameworkHandle, "CFAttributedStringReplaceAttributedString", "")
	registerFunc(&_cFAttributedStringReplaceString, &_cFAttributedStringReplaceStringErr, frameworkHandle, "CFAttributedStringReplaceString", "")
	registerFunc(&_cFAttributedStringSetAttribute, &_cFAttributedStringSetAttributeErr, frameworkHandle, "CFAttributedStringSetAttribute", "")
	registerFunc(&_cFAttributedStringSetAttributes, &_cFAttributedStringSetAttributesErr, frameworkHandle, "CFAttributedStringSetAttributes", "")
	registerFunc(&_cFAutorelease, &_cFAutoreleaseErr, frameworkHandle, "CFAutorelease", "10.9")
	registerFunc(&_cFBagAddValue, &_cFBagAddValueErr, frameworkHandle, "CFBagAddValue", "")
	registerFunc(&_cFBagApplyFunction, &_cFBagApplyFunctionErr, frameworkHandle, "CFBagApplyFunction", "")
	registerFunc(&_cFBagContainsValue, &_cFBagContainsValueErr, frameworkHandle, "CFBagContainsValue", "")
	registerFunc(&_cFBagCreate, &_cFBagCreateErr, frameworkHandle, "CFBagCreate", "")
	registerFunc(&_cFBagCreateCopy, &_cFBagCreateCopyErr, frameworkHandle, "CFBagCreateCopy", "")
	registerFunc(&_cFBagCreateMutable, &_cFBagCreateMutableErr, frameworkHandle, "CFBagCreateMutable", "")
	registerFunc(&_cFBagCreateMutableCopy, &_cFBagCreateMutableCopyErr, frameworkHandle, "CFBagCreateMutableCopy", "")
	registerFunc(&_cFBagGetCount, &_cFBagGetCountErr, frameworkHandle, "CFBagGetCount", "")
	registerFunc(&_cFBagGetCountOfValue, &_cFBagGetCountOfValueErr, frameworkHandle, "CFBagGetCountOfValue", "")
	registerFunc(&_cFBagGetTypeID, &_cFBagGetTypeIDErr, frameworkHandle, "CFBagGetTypeID", "")
	registerFunc(&_cFBagGetValue, &_cFBagGetValueErr, frameworkHandle, "CFBagGetValue", "")
	registerFunc(&_cFBagGetValueIfPresent, &_cFBagGetValueIfPresentErr, frameworkHandle, "CFBagGetValueIfPresent", "")
	registerFunc(&_cFBagGetValues, &_cFBagGetValuesErr, frameworkHandle, "CFBagGetValues", "")
	registerFunc(&_cFBagRemoveAllValues, &_cFBagRemoveAllValuesErr, frameworkHandle, "CFBagRemoveAllValues", "")
	registerFunc(&_cFBagRemoveValue, &_cFBagRemoveValueErr, frameworkHandle, "CFBagRemoveValue", "")
	registerFunc(&_cFBagReplaceValue, &_cFBagReplaceValueErr, frameworkHandle, "CFBagReplaceValue", "")
	registerFunc(&_cFBagSetValue, &_cFBagSetValueErr, frameworkHandle, "CFBagSetValue", "")
	registerFunc(&_cFBinaryHeapAddValue, &_cFBinaryHeapAddValueErr, frameworkHandle, "CFBinaryHeapAddValue", "")
	registerFunc(&_cFBinaryHeapApplyFunction, &_cFBinaryHeapApplyFunctionErr, frameworkHandle, "CFBinaryHeapApplyFunction", "")
	registerFunc(&_cFBinaryHeapContainsValue, &_cFBinaryHeapContainsValueErr, frameworkHandle, "CFBinaryHeapContainsValue", "")
	registerFunc(&_cFBinaryHeapCreate, &_cFBinaryHeapCreateErr, frameworkHandle, "CFBinaryHeapCreate", "")
	registerFunc(&_cFBinaryHeapCreateCopy, &_cFBinaryHeapCreateCopyErr, frameworkHandle, "CFBinaryHeapCreateCopy", "")
	registerFunc(&_cFBinaryHeapGetCount, &_cFBinaryHeapGetCountErr, frameworkHandle, "CFBinaryHeapGetCount", "")
	registerFunc(&_cFBinaryHeapGetCountOfValue, &_cFBinaryHeapGetCountOfValueErr, frameworkHandle, "CFBinaryHeapGetCountOfValue", "")
	registerFunc(&_cFBinaryHeapGetMinimum, &_cFBinaryHeapGetMinimumErr, frameworkHandle, "CFBinaryHeapGetMinimum", "")
	registerFunc(&_cFBinaryHeapGetMinimumIfPresent, &_cFBinaryHeapGetMinimumIfPresentErr, frameworkHandle, "CFBinaryHeapGetMinimumIfPresent", "")
	registerFunc(&_cFBinaryHeapGetTypeID, &_cFBinaryHeapGetTypeIDErr, frameworkHandle, "CFBinaryHeapGetTypeID", "")
	registerFunc(&_cFBinaryHeapGetValues, &_cFBinaryHeapGetValuesErr, frameworkHandle, "CFBinaryHeapGetValues", "")
	registerFunc(&_cFBinaryHeapRemoveAllValues, &_cFBinaryHeapRemoveAllValuesErr, frameworkHandle, "CFBinaryHeapRemoveAllValues", "")
	registerFunc(&_cFBinaryHeapRemoveMinimumValue, &_cFBinaryHeapRemoveMinimumValueErr, frameworkHandle, "CFBinaryHeapRemoveMinimumValue", "")
	registerFunc(&_cFBitVectorContainsBit, &_cFBitVectorContainsBitErr, frameworkHandle, "CFBitVectorContainsBit", "")
	registerFunc(&_cFBitVectorCreate, &_cFBitVectorCreateErr, frameworkHandle, "CFBitVectorCreate", "")
	registerFunc(&_cFBitVectorCreateCopy, &_cFBitVectorCreateCopyErr, frameworkHandle, "CFBitVectorCreateCopy", "")
	registerFunc(&_cFBitVectorCreateMutable, &_cFBitVectorCreateMutableErr, frameworkHandle, "CFBitVectorCreateMutable", "")
	registerFunc(&_cFBitVectorCreateMutableCopy, &_cFBitVectorCreateMutableCopyErr, frameworkHandle, "CFBitVectorCreateMutableCopy", "")
	registerFunc(&_cFBitVectorFlipBitAtIndex, &_cFBitVectorFlipBitAtIndexErr, frameworkHandle, "CFBitVectorFlipBitAtIndex", "")
	registerFunc(&_cFBitVectorFlipBits, &_cFBitVectorFlipBitsErr, frameworkHandle, "CFBitVectorFlipBits", "")
	registerFunc(&_cFBitVectorGetBitAtIndex, &_cFBitVectorGetBitAtIndexErr, frameworkHandle, "CFBitVectorGetBitAtIndex", "")
	registerFunc(&_cFBitVectorGetBits, &_cFBitVectorGetBitsErr, frameworkHandle, "CFBitVectorGetBits", "")
	registerFunc(&_cFBitVectorGetCount, &_cFBitVectorGetCountErr, frameworkHandle, "CFBitVectorGetCount", "")
	registerFunc(&_cFBitVectorGetCountOfBit, &_cFBitVectorGetCountOfBitErr, frameworkHandle, "CFBitVectorGetCountOfBit", "")
	registerFunc(&_cFBitVectorGetFirstIndexOfBit, &_cFBitVectorGetFirstIndexOfBitErr, frameworkHandle, "CFBitVectorGetFirstIndexOfBit", "")
	registerFunc(&_cFBitVectorGetLastIndexOfBit, &_cFBitVectorGetLastIndexOfBitErr, frameworkHandle, "CFBitVectorGetLastIndexOfBit", "")
	registerFunc(&_cFBitVectorGetTypeID, &_cFBitVectorGetTypeIDErr, frameworkHandle, "CFBitVectorGetTypeID", "")
	registerFunc(&_cFBitVectorSetAllBits, &_cFBitVectorSetAllBitsErr, frameworkHandle, "CFBitVectorSetAllBits", "")
	registerFunc(&_cFBitVectorSetBitAtIndex, &_cFBitVectorSetBitAtIndexErr, frameworkHandle, "CFBitVectorSetBitAtIndex", "")
	registerFunc(&_cFBitVectorSetBits, &_cFBitVectorSetBitsErr, frameworkHandle, "CFBitVectorSetBits", "")
	registerFunc(&_cFBitVectorSetCount, &_cFBitVectorSetCountErr, frameworkHandle, "CFBitVectorSetCount", "")
	registerFunc(&_cFBooleanGetTypeID, &_cFBooleanGetTypeIDErr, frameworkHandle, "CFBooleanGetTypeID", "")
	registerFunc(&_cFBooleanGetValue, &_cFBooleanGetValueErr, frameworkHandle, "CFBooleanGetValue", "")
	registerFunc(&_cFBundleCopyAuxiliaryExecutableURL, &_cFBundleCopyAuxiliaryExecutableURLErr, frameworkHandle, "CFBundleCopyAuxiliaryExecutableURL", "")
	registerFunc(&_cFBundleCopyBuiltInPlugInsURL, &_cFBundleCopyBuiltInPlugInsURLErr, frameworkHandle, "CFBundleCopyBuiltInPlugInsURL", "")
	registerFunc(&_cFBundleCopyBundleLocalizations, &_cFBundleCopyBundleLocalizationsErr, frameworkHandle, "CFBundleCopyBundleLocalizations", "")
	registerFunc(&_cFBundleCopyBundleURL, &_cFBundleCopyBundleURLErr, frameworkHandle, "CFBundleCopyBundleURL", "")
	registerFunc(&_cFBundleCopyExecutableArchitectures, &_cFBundleCopyExecutableArchitecturesErr, frameworkHandle, "CFBundleCopyExecutableArchitectures", "10.5")
	registerFunc(&_cFBundleCopyExecutableArchitecturesForURL, &_cFBundleCopyExecutableArchitecturesForURLErr, frameworkHandle, "CFBundleCopyExecutableArchitecturesForURL", "10.5")
	registerFunc(&_cFBundleCopyExecutableURL, &_cFBundleCopyExecutableURLErr, frameworkHandle, "CFBundleCopyExecutableURL", "")
	registerFunc(&_cFBundleCopyInfoDictionaryForURL, &_cFBundleCopyInfoDictionaryForURLErr, frameworkHandle, "CFBundleCopyInfoDictionaryForURL", "")
	registerFunc(&_cFBundleCopyInfoDictionaryInDirectory, &_cFBundleCopyInfoDictionaryInDirectoryErr, frameworkHandle, "CFBundleCopyInfoDictionaryInDirectory", "")
	registerFunc(&_cFBundleCopyLocalizationsForPreferences, &_cFBundleCopyLocalizationsForPreferencesErr, frameworkHandle, "CFBundleCopyLocalizationsForPreferences", "")
	registerFunc(&_cFBundleCopyLocalizationsForURL, &_cFBundleCopyLocalizationsForURLErr, frameworkHandle, "CFBundleCopyLocalizationsForURL", "")
	registerFunc(&_cFBundleCopyLocalizedString, &_cFBundleCopyLocalizedStringErr, frameworkHandle, "CFBundleCopyLocalizedString", "")
	registerFunc(&_cFBundleCopyLocalizedStringForLocalizations, &_cFBundleCopyLocalizedStringForLocalizationsErr, frameworkHandle, "CFBundleCopyLocalizedStringForLocalizations", "15.4")
	registerFunc(&_cFBundleCopyPreferredLocalizationsFromArray, &_cFBundleCopyPreferredLocalizationsFromArrayErr, frameworkHandle, "CFBundleCopyPreferredLocalizationsFromArray", "")
	registerFunc(&_cFBundleCopyPrivateFrameworksURL, &_cFBundleCopyPrivateFrameworksURLErr, frameworkHandle, "CFBundleCopyPrivateFrameworksURL", "")
	registerFunc(&_cFBundleCopyResourceURL, &_cFBundleCopyResourceURLErr, frameworkHandle, "CFBundleCopyResourceURL", "")
	registerFunc(&_cFBundleCopyResourceURLForLocalization, &_cFBundleCopyResourceURLForLocalizationErr, frameworkHandle, "CFBundleCopyResourceURLForLocalization", "")
	registerFunc(&_cFBundleCopyResourceURLInDirectory, &_cFBundleCopyResourceURLInDirectoryErr, frameworkHandle, "CFBundleCopyResourceURLInDirectory", "")
	registerFunc(&_cFBundleCopyResourceURLsOfType, &_cFBundleCopyResourceURLsOfTypeErr, frameworkHandle, "CFBundleCopyResourceURLsOfType", "")
	registerFunc(&_cFBundleCopyResourceURLsOfTypeForLocalization, &_cFBundleCopyResourceURLsOfTypeForLocalizationErr, frameworkHandle, "CFBundleCopyResourceURLsOfTypeForLocalization", "")
	registerFunc(&_cFBundleCopyResourceURLsOfTypeInDirectory, &_cFBundleCopyResourceURLsOfTypeInDirectoryErr, frameworkHandle, "CFBundleCopyResourceURLsOfTypeInDirectory", "")
	registerFunc(&_cFBundleCopyResourcesDirectoryURL, &_cFBundleCopyResourcesDirectoryURLErr, frameworkHandle, "CFBundleCopyResourcesDirectoryURL", "")
	registerFunc(&_cFBundleCopySharedFrameworksURL, &_cFBundleCopySharedFrameworksURLErr, frameworkHandle, "CFBundleCopySharedFrameworksURL", "")
	registerFunc(&_cFBundleCopySharedSupportURL, &_cFBundleCopySharedSupportURLErr, frameworkHandle, "CFBundleCopySharedSupportURL", "")
	registerFunc(&_cFBundleCopySupportFilesDirectoryURL, &_cFBundleCopySupportFilesDirectoryURLErr, frameworkHandle, "CFBundleCopySupportFilesDirectoryURL", "")
	registerFunc(&_cFBundleCreate, &_cFBundleCreateErr, frameworkHandle, "CFBundleCreate", "")
	registerFunc(&_cFBundleCreateBundlesFromDirectory, &_cFBundleCreateBundlesFromDirectoryErr, frameworkHandle, "CFBundleCreateBundlesFromDirectory", "")
	registerFunc(&_cFBundleGetAllBundles, &_cFBundleGetAllBundlesErr, frameworkHandle, "CFBundleGetAllBundles", "")
	registerFunc(&_cFBundleGetBundleWithIdentifier, &_cFBundleGetBundleWithIdentifierErr, frameworkHandle, "CFBundleGetBundleWithIdentifier", "")
	registerFunc(&_cFBundleGetDataPointerForName, &_cFBundleGetDataPointerForNameErr, frameworkHandle, "CFBundleGetDataPointerForName", "")
	registerFunc(&_cFBundleGetDataPointersForNames, &_cFBundleGetDataPointersForNamesErr, frameworkHandle, "CFBundleGetDataPointersForNames", "")
	registerFunc(&_cFBundleGetDevelopmentRegion, &_cFBundleGetDevelopmentRegionErr, frameworkHandle, "CFBundleGetDevelopmentRegion", "")
	registerFunc(&_cFBundleGetFunctionPointerForName, &_cFBundleGetFunctionPointerForNameErr, frameworkHandle, "CFBundleGetFunctionPointerForName", "")
	registerFunc(&_cFBundleGetFunctionPointersForNames, &_cFBundleGetFunctionPointersForNamesErr, frameworkHandle, "CFBundleGetFunctionPointersForNames", "")
	registerFunc(&_cFBundleGetIdentifier, &_cFBundleGetIdentifierErr, frameworkHandle, "CFBundleGetIdentifier", "")
	registerFunc(&_cFBundleGetInfoDictionary, &_cFBundleGetInfoDictionaryErr, frameworkHandle, "CFBundleGetInfoDictionary", "")
	registerFunc(&_cFBundleGetLocalInfoDictionary, &_cFBundleGetLocalInfoDictionaryErr, frameworkHandle, "CFBundleGetLocalInfoDictionary", "")
	registerFunc(&_cFBundleGetMainBundle, &_cFBundleGetMainBundleErr, frameworkHandle, "CFBundleGetMainBundle", "")
	registerFunc(&_cFBundleGetPackageInfo, &_cFBundleGetPackageInfoErr, frameworkHandle, "CFBundleGetPackageInfo", "")
	registerFunc(&_cFBundleGetPackageInfoInDirectory, &_cFBundleGetPackageInfoInDirectoryErr, frameworkHandle, "CFBundleGetPackageInfoInDirectory", "")
	registerFunc(&_cFBundleGetPlugIn, &_cFBundleGetPlugInErr, frameworkHandle, "CFBundleGetPlugIn", "")
	registerFunc(&_cFBundleGetTypeID, &_cFBundleGetTypeIDErr, frameworkHandle, "CFBundleGetTypeID", "")
	registerFunc(&_cFBundleGetValueForInfoDictionaryKey, &_cFBundleGetValueForInfoDictionaryKeyErr, frameworkHandle, "CFBundleGetValueForInfoDictionaryKey", "")
	registerFunc(&_cFBundleGetVersionNumber, &_cFBundleGetVersionNumberErr, frameworkHandle, "CFBundleGetVersionNumber", "")
	registerFunc(&_cFBundleIsArchitectureLoadable, &_cFBundleIsArchitectureLoadableErr, frameworkHandle, "CFBundleIsArchitectureLoadable", "11.0")
	registerFunc(&_cFBundleIsExecutableLoadable, &_cFBundleIsExecutableLoadableErr, frameworkHandle, "CFBundleIsExecutableLoadable", "11.0")
	registerFunc(&_cFBundleIsExecutableLoadableForURL, &_cFBundleIsExecutableLoadableForURLErr, frameworkHandle, "CFBundleIsExecutableLoadableForURL", "11.0")
	registerFunc(&_cFBundleIsExecutableLoaded, &_cFBundleIsExecutableLoadedErr, frameworkHandle, "CFBundleIsExecutableLoaded", "")
	registerFunc(&_cFBundleLoadExecutable, &_cFBundleLoadExecutableErr, frameworkHandle, "CFBundleLoadExecutable", "")
	registerFunc(&_cFBundleLoadExecutableAndReturnError, &_cFBundleLoadExecutableAndReturnErrorErr, frameworkHandle, "CFBundleLoadExecutableAndReturnError", "10.5")
	registerFunc(&_cFBundlePreflightExecutable, &_cFBundlePreflightExecutableErr, frameworkHandle, "CFBundlePreflightExecutable", "10.5")
	registerFunc(&_cFBundleUnloadExecutable, &_cFBundleUnloadExecutableErr, frameworkHandle, "CFBundleUnloadExecutable", "")
	registerFunc(&_cFCalendarAddComponents, &_cFCalendarAddComponentsErr, frameworkHandle, "CFCalendarAddComponents", "")
	registerFunc(&_cFCalendarComposeAbsoluteTime, &_cFCalendarComposeAbsoluteTimeErr, frameworkHandle, "CFCalendarComposeAbsoluteTime", "")
	registerFunc(&_cFCalendarCopyCurrent, &_cFCalendarCopyCurrentErr, frameworkHandle, "CFCalendarCopyCurrent", "")
	registerFunc(&_cFCalendarCopyLocale, &_cFCalendarCopyLocaleErr, frameworkHandle, "CFCalendarCopyLocale", "")
	registerFunc(&_cFCalendarCopyTimeZone, &_cFCalendarCopyTimeZoneErr, frameworkHandle, "CFCalendarCopyTimeZone", "")
	registerFunc(&_cFCalendarCreateWithIdentifier, &_cFCalendarCreateWithIdentifierErr, frameworkHandle, "CFCalendarCreateWithIdentifier", "")
	registerFunc(&_cFCalendarDecomposeAbsoluteTime, &_cFCalendarDecomposeAbsoluteTimeErr, frameworkHandle, "CFCalendarDecomposeAbsoluteTime", "")
	registerFunc(&_cFCalendarGetComponentDifference, &_cFCalendarGetComponentDifferenceErr, frameworkHandle, "CFCalendarGetComponentDifference", "")
	registerFunc(&_cFCalendarGetFirstWeekday, &_cFCalendarGetFirstWeekdayErr, frameworkHandle, "CFCalendarGetFirstWeekday", "")
	registerFunc(&_cFCalendarGetIdentifier, &_cFCalendarGetIdentifierErr, frameworkHandle, "CFCalendarGetIdentifier", "")
	registerFunc(&_cFCalendarGetMaximumRangeOfUnit, &_cFCalendarGetMaximumRangeOfUnitErr, frameworkHandle, "CFCalendarGetMaximumRangeOfUnit", "")
	registerFunc(&_cFCalendarGetMinimumDaysInFirstWeek, &_cFCalendarGetMinimumDaysInFirstWeekErr, frameworkHandle, "CFCalendarGetMinimumDaysInFirstWeek", "")
	registerFunc(&_cFCalendarGetMinimumRangeOfUnit, &_cFCalendarGetMinimumRangeOfUnitErr, frameworkHandle, "CFCalendarGetMinimumRangeOfUnit", "")
	registerFunc(&_cFCalendarGetOrdinalityOfUnit, &_cFCalendarGetOrdinalityOfUnitErr, frameworkHandle, "CFCalendarGetOrdinalityOfUnit", "")
	registerFunc(&_cFCalendarGetRangeOfUnit, &_cFCalendarGetRangeOfUnitErr, frameworkHandle, "CFCalendarGetRangeOfUnit", "")
	registerFunc(&_cFCalendarGetTimeRangeOfUnit, &_cFCalendarGetTimeRangeOfUnitErr, frameworkHandle, "CFCalendarGetTimeRangeOfUnit", "10.5")
	registerFunc(&_cFCalendarGetTypeID, &_cFCalendarGetTypeIDErr, frameworkHandle, "CFCalendarGetTypeID", "")
	registerFunc(&_cFCalendarSetFirstWeekday, &_cFCalendarSetFirstWeekdayErr, frameworkHandle, "CFCalendarSetFirstWeekday", "")
	registerFunc(&_cFCalendarSetLocale, &_cFCalendarSetLocaleErr, frameworkHandle, "CFCalendarSetLocale", "")
	registerFunc(&_cFCalendarSetMinimumDaysInFirstWeek, &_cFCalendarSetMinimumDaysInFirstWeekErr, frameworkHandle, "CFCalendarSetMinimumDaysInFirstWeek", "")
	registerFunc(&_cFCalendarSetTimeZone, &_cFCalendarSetTimeZoneErr, frameworkHandle, "CFCalendarSetTimeZone", "")
	registerFunc(&_cFCharacterSetAddCharactersInRange, &_cFCharacterSetAddCharactersInRangeErr, frameworkHandle, "CFCharacterSetAddCharactersInRange", "")
	registerFunc(&_cFCharacterSetAddCharactersInString, &_cFCharacterSetAddCharactersInStringErr, frameworkHandle, "CFCharacterSetAddCharactersInString", "")
	registerFunc(&_cFCharacterSetCreateBitmapRepresentation, &_cFCharacterSetCreateBitmapRepresentationErr, frameworkHandle, "CFCharacterSetCreateBitmapRepresentation", "")
	registerFunc(&_cFCharacterSetCreateCopy, &_cFCharacterSetCreateCopyErr, frameworkHandle, "CFCharacterSetCreateCopy", "")
	registerFunc(&_cFCharacterSetCreateInvertedSet, &_cFCharacterSetCreateInvertedSetErr, frameworkHandle, "CFCharacterSetCreateInvertedSet", "")
	registerFunc(&_cFCharacterSetCreateMutable, &_cFCharacterSetCreateMutableErr, frameworkHandle, "CFCharacterSetCreateMutable", "")
	registerFunc(&_cFCharacterSetCreateMutableCopy, &_cFCharacterSetCreateMutableCopyErr, frameworkHandle, "CFCharacterSetCreateMutableCopy", "")
	registerFunc(&_cFCharacterSetCreateWithBitmapRepresentation, &_cFCharacterSetCreateWithBitmapRepresentationErr, frameworkHandle, "CFCharacterSetCreateWithBitmapRepresentation", "")
	registerFunc(&_cFCharacterSetCreateWithCharactersInRange, &_cFCharacterSetCreateWithCharactersInRangeErr, frameworkHandle, "CFCharacterSetCreateWithCharactersInRange", "")
	registerFunc(&_cFCharacterSetCreateWithCharactersInString, &_cFCharacterSetCreateWithCharactersInStringErr, frameworkHandle, "CFCharacterSetCreateWithCharactersInString", "")
	registerFunc(&_cFCharacterSetGetPredefined, &_cFCharacterSetGetPredefinedErr, frameworkHandle, "CFCharacterSetGetPredefined", "")
	registerFunc(&_cFCharacterSetGetTypeID, &_cFCharacterSetGetTypeIDErr, frameworkHandle, "CFCharacterSetGetTypeID", "")
	registerFunc(&_cFCharacterSetHasMemberInPlane, &_cFCharacterSetHasMemberInPlaneErr, frameworkHandle, "CFCharacterSetHasMemberInPlane", "")
	registerFunc(&_cFCharacterSetIntersect, &_cFCharacterSetIntersectErr, frameworkHandle, "CFCharacterSetIntersect", "")
	registerFunc(&_cFCharacterSetInvert, &_cFCharacterSetInvertErr, frameworkHandle, "CFCharacterSetInvert", "")
	registerFunc(&_cFCharacterSetIsCharacterMember, &_cFCharacterSetIsCharacterMemberErr, frameworkHandle, "CFCharacterSetIsCharacterMember", "")
	registerFunc(&_cFCharacterSetIsLongCharacterMember, &_cFCharacterSetIsLongCharacterMemberErr, frameworkHandle, "CFCharacterSetIsLongCharacterMember", "")
	registerFunc(&_cFCharacterSetIsSupersetOfSet, &_cFCharacterSetIsSupersetOfSetErr, frameworkHandle, "CFCharacterSetIsSupersetOfSet", "")
	registerFunc(&_cFCharacterSetRemoveCharactersInRange, &_cFCharacterSetRemoveCharactersInRangeErr, frameworkHandle, "CFCharacterSetRemoveCharactersInRange", "")
	registerFunc(&_cFCharacterSetRemoveCharactersInString, &_cFCharacterSetRemoveCharactersInStringErr, frameworkHandle, "CFCharacterSetRemoveCharactersInString", "")
	registerFunc(&_cFCharacterSetUnion, &_cFCharacterSetUnionErr, frameworkHandle, "CFCharacterSetUnion", "")
	registerFunc(&_cFCopyDescription, &_cFCopyDescriptionErr, frameworkHandle, "CFCopyDescription", "")
	registerFunc(&_cFCopyTypeIDDescription, &_cFCopyTypeIDDescriptionErr, frameworkHandle, "CFCopyTypeIDDescription", "")
	registerFunc(&_cFDataAppendBytes, &_cFDataAppendBytesErr, frameworkHandle, "CFDataAppendBytes", "")
	registerFunc(&_cFDataCreate, &_cFDataCreateErr, frameworkHandle, "CFDataCreate", "")
	registerFunc(&_cFDataCreateCopy, &_cFDataCreateCopyErr, frameworkHandle, "CFDataCreateCopy", "")
	registerFunc(&_cFDataCreateMutable, &_cFDataCreateMutableErr, frameworkHandle, "CFDataCreateMutable", "")
	registerFunc(&_cFDataCreateMutableCopy, &_cFDataCreateMutableCopyErr, frameworkHandle, "CFDataCreateMutableCopy", "")
	registerFunc(&_cFDataCreateWithBytesNoCopy, &_cFDataCreateWithBytesNoCopyErr, frameworkHandle, "CFDataCreateWithBytesNoCopy", "")
	registerFunc(&_cFDataDeleteBytes, &_cFDataDeleteBytesErr, frameworkHandle, "CFDataDeleteBytes", "")
	registerFunc(&_cFDataFind, &_cFDataFindErr, frameworkHandle, "CFDataFind", "10.6")
	registerFunc(&_cFDataGetBytePtr, &_cFDataGetBytePtrErr, frameworkHandle, "CFDataGetBytePtr", "")
	registerFunc(&_cFDataGetBytes, &_cFDataGetBytesErr, frameworkHandle, "CFDataGetBytes", "")
	registerFunc(&_cFDataGetLength, &_cFDataGetLengthErr, frameworkHandle, "CFDataGetLength", "")
	registerFunc(&_cFDataGetMutableBytePtr, &_cFDataGetMutableBytePtrErr, frameworkHandle, "CFDataGetMutableBytePtr", "")
	registerFunc(&_cFDataGetTypeID, &_cFDataGetTypeIDErr, frameworkHandle, "CFDataGetTypeID", "")
	registerFunc(&_cFDataIncreaseLength, &_cFDataIncreaseLengthErr, frameworkHandle, "CFDataIncreaseLength", "")
	registerFunc(&_cFDataReplaceBytes, &_cFDataReplaceBytesErr, frameworkHandle, "CFDataReplaceBytes", "")
	registerFunc(&_cFDataSetLength, &_cFDataSetLengthErr, frameworkHandle, "CFDataSetLength", "")
	registerFunc(&_cFDateCompare, &_cFDateCompareErr, frameworkHandle, "CFDateCompare", "")
	registerFunc(&_cFDateCreate, &_cFDateCreateErr, frameworkHandle, "CFDateCreate", "")
	registerFunc(&_cFDateFormatterCopyProperty, &_cFDateFormatterCopyPropertyErr, frameworkHandle, "CFDateFormatterCopyProperty", "")
	registerFunc(&_cFDateFormatterCreate, &_cFDateFormatterCreateErr, frameworkHandle, "CFDateFormatterCreate", "")
	registerFunc(&_cFDateFormatterCreateDateFormatFromTemplate, &_cFDateFormatterCreateDateFormatFromTemplateErr, frameworkHandle, "CFDateFormatterCreateDateFormatFromTemplate", "10.6")
	registerFunc(&_cFDateFormatterCreateDateFromString, &_cFDateFormatterCreateDateFromStringErr, frameworkHandle, "CFDateFormatterCreateDateFromString", "")
	registerFunc(&_cFDateFormatterCreateISO8601Formatter, &_cFDateFormatterCreateISO8601FormatterErr, frameworkHandle, "CFDateFormatterCreateISO8601Formatter", "10.12")
	registerFunc(&_cFDateFormatterCreateStringWithAbsoluteTime, &_cFDateFormatterCreateStringWithAbsoluteTimeErr, frameworkHandle, "CFDateFormatterCreateStringWithAbsoluteTime", "")
	registerFunc(&_cFDateFormatterCreateStringWithDate, &_cFDateFormatterCreateStringWithDateErr, frameworkHandle, "CFDateFormatterCreateStringWithDate", "")
	registerFunc(&_cFDateFormatterGetAbsoluteTimeFromString, &_cFDateFormatterGetAbsoluteTimeFromStringErr, frameworkHandle, "CFDateFormatterGetAbsoluteTimeFromString", "")
	registerFunc(&_cFDateFormatterGetDateStyle, &_cFDateFormatterGetDateStyleErr, frameworkHandle, "CFDateFormatterGetDateStyle", "")
	registerFunc(&_cFDateFormatterGetFormat, &_cFDateFormatterGetFormatErr, frameworkHandle, "CFDateFormatterGetFormat", "")
	registerFunc(&_cFDateFormatterGetLocale, &_cFDateFormatterGetLocaleErr, frameworkHandle, "CFDateFormatterGetLocale", "")
	registerFunc(&_cFDateFormatterGetTimeStyle, &_cFDateFormatterGetTimeStyleErr, frameworkHandle, "CFDateFormatterGetTimeStyle", "")
	registerFunc(&_cFDateFormatterGetTypeID, &_cFDateFormatterGetTypeIDErr, frameworkHandle, "CFDateFormatterGetTypeID", "")
	registerFunc(&_cFDateFormatterSetFormat, &_cFDateFormatterSetFormatErr, frameworkHandle, "CFDateFormatterSetFormat", "")
	registerFunc(&_cFDateFormatterSetProperty, &_cFDateFormatterSetPropertyErr, frameworkHandle, "CFDateFormatterSetProperty", "")
	registerFunc(&_cFDateGetAbsoluteTime, &_cFDateGetAbsoluteTimeErr, frameworkHandle, "CFDateGetAbsoluteTime", "")
	registerFunc(&_cFDateGetTimeIntervalSinceDate, &_cFDateGetTimeIntervalSinceDateErr, frameworkHandle, "CFDateGetTimeIntervalSinceDate", "")
	registerFunc(&_cFDateGetTypeID, &_cFDateGetTypeIDErr, frameworkHandle, "CFDateGetTypeID", "")
	registerFunc(&_cFDictionaryAddValue, &_cFDictionaryAddValueErr, frameworkHandle, "CFDictionaryAddValue", "")
	registerFunc(&_cFDictionaryApplyFunction, &_cFDictionaryApplyFunctionErr, frameworkHandle, "CFDictionaryApplyFunction", "")
	registerFunc(&_cFDictionaryContainsKey, &_cFDictionaryContainsKeyErr, frameworkHandle, "CFDictionaryContainsKey", "")
	registerFunc(&_cFDictionaryContainsValue, &_cFDictionaryContainsValueErr, frameworkHandle, "CFDictionaryContainsValue", "")
	registerFunc(&_cFDictionaryCreate, &_cFDictionaryCreateErr, frameworkHandle, "CFDictionaryCreate", "")
	registerFunc(&_cFDictionaryCreateCopy, &_cFDictionaryCreateCopyErr, frameworkHandle, "CFDictionaryCreateCopy", "")
	registerFunc(&_cFDictionaryCreateMutable, &_cFDictionaryCreateMutableErr, frameworkHandle, "CFDictionaryCreateMutable", "")
	registerFunc(&_cFDictionaryCreateMutableCopy, &_cFDictionaryCreateMutableCopyErr, frameworkHandle, "CFDictionaryCreateMutableCopy", "")
	registerFunc(&_cFDictionaryGetCount, &_cFDictionaryGetCountErr, frameworkHandle, "CFDictionaryGetCount", "")
	registerFunc(&_cFDictionaryGetCountOfKey, &_cFDictionaryGetCountOfKeyErr, frameworkHandle, "CFDictionaryGetCountOfKey", "")
	registerFunc(&_cFDictionaryGetCountOfValue, &_cFDictionaryGetCountOfValueErr, frameworkHandle, "CFDictionaryGetCountOfValue", "")
	registerFunc(&_cFDictionaryGetKeysAndValues, &_cFDictionaryGetKeysAndValuesErr, frameworkHandle, "CFDictionaryGetKeysAndValues", "")
	registerFunc(&_cFDictionaryGetTypeID, &_cFDictionaryGetTypeIDErr, frameworkHandle, "CFDictionaryGetTypeID", "")
	registerFunc(&_cFDictionaryGetValue, &_cFDictionaryGetValueErr, frameworkHandle, "CFDictionaryGetValue", "")
	registerFunc(&_cFDictionaryGetValueIfPresent, &_cFDictionaryGetValueIfPresentErr, frameworkHandle, "CFDictionaryGetValueIfPresent", "")
	registerFunc(&_cFDictionaryRemoveAllValues, &_cFDictionaryRemoveAllValuesErr, frameworkHandle, "CFDictionaryRemoveAllValues", "")
	registerFunc(&_cFDictionaryRemoveValue, &_cFDictionaryRemoveValueErr, frameworkHandle, "CFDictionaryRemoveValue", "")
	registerFunc(&_cFDictionaryReplaceValue, &_cFDictionaryReplaceValueErr, frameworkHandle, "CFDictionaryReplaceValue", "")
	registerFunc(&_cFDictionarySetValue, &_cFDictionarySetValueErr, frameworkHandle, "CFDictionarySetValue", "")
	registerFunc(&_cFEqual, &_cFEqualErr, frameworkHandle, "CFEqual", "")
	registerFunc(&_cFErrorCopyDescription, &_cFErrorCopyDescriptionErr, frameworkHandle, "CFErrorCopyDescription", "10.5")
	registerFunc(&_cFErrorCopyFailureReason, &_cFErrorCopyFailureReasonErr, frameworkHandle, "CFErrorCopyFailureReason", "10.5")
	registerFunc(&_cFErrorCopyRecoverySuggestion, &_cFErrorCopyRecoverySuggestionErr, frameworkHandle, "CFErrorCopyRecoverySuggestion", "10.5")
	registerFunc(&_cFErrorCopyUserInfo, &_cFErrorCopyUserInfoErr, frameworkHandle, "CFErrorCopyUserInfo", "10.5")
	registerFunc(&_cFErrorCreate, &_cFErrorCreateErr, frameworkHandle, "CFErrorCreate", "10.5")
	registerFunc(&_cFErrorCreateWithUserInfoKeysAndValues, &_cFErrorCreateWithUserInfoKeysAndValuesErr, frameworkHandle, "CFErrorCreateWithUserInfoKeysAndValues", "10.5")
	registerFunc(&_cFErrorGetCode, &_cFErrorGetCodeErr, frameworkHandle, "CFErrorGetCode", "10.5")
	registerFunc(&_cFErrorGetDomain, &_cFErrorGetDomainErr, frameworkHandle, "CFErrorGetDomain", "10.5")
	registerFunc(&_cFErrorGetTypeID, &_cFErrorGetTypeIDErr, frameworkHandle, "CFErrorGetTypeID", "10.5")
	registerFunc(&_cFFileDescriptorCreate, &_cFFileDescriptorCreateErr, frameworkHandle, "CFFileDescriptorCreate", "10.5")
	registerFunc(&_cFFileDescriptorCreateRunLoopSource, &_cFFileDescriptorCreateRunLoopSourceErr, frameworkHandle, "CFFileDescriptorCreateRunLoopSource", "10.5")
	registerFunc(&_cFFileDescriptorDisableCallBacks, &_cFFileDescriptorDisableCallBacksErr, frameworkHandle, "CFFileDescriptorDisableCallBacks", "10.5")
	registerFunc(&_cFFileDescriptorEnableCallBacks, &_cFFileDescriptorEnableCallBacksErr, frameworkHandle, "CFFileDescriptorEnableCallBacks", "10.5")
	registerFunc(&_cFFileDescriptorGetContext, &_cFFileDescriptorGetContextErr, frameworkHandle, "CFFileDescriptorGetContext", "10.5")
	registerFunc(&_cFFileDescriptorGetNativeDescriptor, &_cFFileDescriptorGetNativeDescriptorErr, frameworkHandle, "CFFileDescriptorGetNativeDescriptor", "10.5")
	registerFunc(&_cFFileDescriptorGetTypeID, &_cFFileDescriptorGetTypeIDErr, frameworkHandle, "CFFileDescriptorGetTypeID", "10.5")
	registerFunc(&_cFFileDescriptorInvalidate, &_cFFileDescriptorInvalidateErr, frameworkHandle, "CFFileDescriptorInvalidate", "10.5")
	registerFunc(&_cFFileDescriptorIsValid, &_cFFileDescriptorIsValidErr, frameworkHandle, "CFFileDescriptorIsValid", "10.5")
	registerFunc(&_cFFileSecurityClearProperties, &_cFFileSecurityClearPropertiesErr, frameworkHandle, "CFFileSecurityClearProperties", "10.8")
	registerFunc(&_cFFileSecurityCopyAccessControlList, &_cFFileSecurityCopyAccessControlListErr, frameworkHandle, "CFFileSecurityCopyAccessControlList", "10.7")
	registerFunc(&_cFFileSecurityCopyGroupUUID, &_cFFileSecurityCopyGroupUUIDErr, frameworkHandle, "CFFileSecurityCopyGroupUUID", "10.7")
	registerFunc(&_cFFileSecurityCopyOwnerUUID, &_cFFileSecurityCopyOwnerUUIDErr, frameworkHandle, "CFFileSecurityCopyOwnerUUID", "10.7")
	registerFunc(&_cFFileSecurityCreate, &_cFFileSecurityCreateErr, frameworkHandle, "CFFileSecurityCreate", "10.7")
	registerFunc(&_cFFileSecurityCreateCopy, &_cFFileSecurityCreateCopyErr, frameworkHandle, "CFFileSecurityCreateCopy", "10.7")
	registerFunc(&_cFFileSecurityGetGroup, &_cFFileSecurityGetGroupErr, frameworkHandle, "CFFileSecurityGetGroup", "10.7")
	registerFunc(&_cFFileSecurityGetMode, &_cFFileSecurityGetModeErr, frameworkHandle, "CFFileSecurityGetMode", "10.7")
	registerFunc(&_cFFileSecurityGetOwner, &_cFFileSecurityGetOwnerErr, frameworkHandle, "CFFileSecurityGetOwner", "10.7")
	registerFunc(&_cFFileSecurityGetTypeID, &_cFFileSecurityGetTypeIDErr, frameworkHandle, "CFFileSecurityGetTypeID", "10.7")
	registerFunc(&_cFFileSecuritySetAccessControlList, &_cFFileSecuritySetAccessControlListErr, frameworkHandle, "CFFileSecuritySetAccessControlList", "10.7")
	registerFunc(&_cFFileSecuritySetGroup, &_cFFileSecuritySetGroupErr, frameworkHandle, "CFFileSecuritySetGroup", "10.7")
	registerFunc(&_cFFileSecuritySetGroupUUID, &_cFFileSecuritySetGroupUUIDErr, frameworkHandle, "CFFileSecuritySetGroupUUID", "10.7")
	registerFunc(&_cFFileSecuritySetMode, &_cFFileSecuritySetModeErr, frameworkHandle, "CFFileSecuritySetMode", "10.7")
	registerFunc(&_cFFileSecuritySetOwner, &_cFFileSecuritySetOwnerErr, frameworkHandle, "CFFileSecuritySetOwner", "10.7")
	registerFunc(&_cFFileSecuritySetOwnerUUID, &_cFFileSecuritySetOwnerUUIDErr, frameworkHandle, "CFFileSecuritySetOwnerUUID", "10.7")
	registerFunc(&_cFGetAllocator, &_cFGetAllocatorErr, frameworkHandle, "CFGetAllocator", "")
	registerFunc(&_cFGetRetainCount, &_cFGetRetainCountErr, frameworkHandle, "CFGetRetainCount", "")
	registerFunc(&_cFGetTypeID, &_cFGetTypeIDErr, frameworkHandle, "CFGetTypeID", "")
	registerFunc(&_cFHash, &_cFHashErr, frameworkHandle, "CFHash", "")
	registerFunc(&_cFLocaleCopyAvailableLocaleIdentifiers, &_cFLocaleCopyAvailableLocaleIdentifiersErr, frameworkHandle, "CFLocaleCopyAvailableLocaleIdentifiers", "")
	registerFunc(&_cFLocaleCopyCommonISOCurrencyCodes, &_cFLocaleCopyCommonISOCurrencyCodesErr, frameworkHandle, "CFLocaleCopyCommonISOCurrencyCodes", "10.5")
	registerFunc(&_cFLocaleCopyCurrent, &_cFLocaleCopyCurrentErr, frameworkHandle, "CFLocaleCopyCurrent", "")
	registerFunc(&_cFLocaleCopyDisplayNameForPropertyValue, &_cFLocaleCopyDisplayNameForPropertyValueErr, frameworkHandle, "CFLocaleCopyDisplayNameForPropertyValue", "")
	registerFunc(&_cFLocaleCopyISOCountryCodes, &_cFLocaleCopyISOCountryCodesErr, frameworkHandle, "CFLocaleCopyISOCountryCodes", "")
	registerFunc(&_cFLocaleCopyISOCurrencyCodes, &_cFLocaleCopyISOCurrencyCodesErr, frameworkHandle, "CFLocaleCopyISOCurrencyCodes", "")
	registerFunc(&_cFLocaleCopyISOLanguageCodes, &_cFLocaleCopyISOLanguageCodesErr, frameworkHandle, "CFLocaleCopyISOLanguageCodes", "")
	registerFunc(&_cFLocaleCopyPreferredLanguages, &_cFLocaleCopyPreferredLanguagesErr, frameworkHandle, "CFLocaleCopyPreferredLanguages", "10.5")
	registerFunc(&_cFLocaleCreate, &_cFLocaleCreateErr, frameworkHandle, "CFLocaleCreate", "")
	registerFunc(&_cFLocaleCreateCanonicalLanguageIdentifierFromString, &_cFLocaleCreateCanonicalLanguageIdentifierFromStringErr, frameworkHandle, "CFLocaleCreateCanonicalLanguageIdentifierFromString", "")
	registerFunc(&_cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes, &_cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodesErr, frameworkHandle, "CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes", "")
	registerFunc(&_cFLocaleCreateCanonicalLocaleIdentifierFromString, &_cFLocaleCreateCanonicalLocaleIdentifierFromStringErr, frameworkHandle, "CFLocaleCreateCanonicalLocaleIdentifierFromString", "")
	registerFunc(&_cFLocaleCreateComponentsFromLocaleIdentifier, &_cFLocaleCreateComponentsFromLocaleIdentifierErr, frameworkHandle, "CFLocaleCreateComponentsFromLocaleIdentifier", "")
	registerFunc(&_cFLocaleCreateCopy, &_cFLocaleCreateCopyErr, frameworkHandle, "CFLocaleCreateCopy", "")
	registerFunc(&_cFLocaleCreateLocaleIdentifierFromComponents, &_cFLocaleCreateLocaleIdentifierFromComponentsErr, frameworkHandle, "CFLocaleCreateLocaleIdentifierFromComponents", "")
	registerFunc(&_cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode, &_cFLocaleCreateLocaleIdentifierFromWindowsLocaleCodeErr, frameworkHandle, "CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode", "10.6")
	registerFunc(&_cFLocaleGetIdentifier, &_cFLocaleGetIdentifierErr, frameworkHandle, "CFLocaleGetIdentifier", "")
	registerFunc(&_cFLocaleGetLanguageCharacterDirection, &_cFLocaleGetLanguageCharacterDirectionErr, frameworkHandle, "CFLocaleGetLanguageCharacterDirection", "10.6")
	registerFunc(&_cFLocaleGetLanguageLineDirection, &_cFLocaleGetLanguageLineDirectionErr, frameworkHandle, "CFLocaleGetLanguageLineDirection", "10.6")
	registerFunc(&_cFLocaleGetSystem, &_cFLocaleGetSystemErr, frameworkHandle, "CFLocaleGetSystem", "")
	registerFunc(&_cFLocaleGetTypeID, &_cFLocaleGetTypeIDErr, frameworkHandle, "CFLocaleGetTypeID", "")
	registerFunc(&_cFLocaleGetValue, &_cFLocaleGetValueErr, frameworkHandle, "CFLocaleGetValue", "")
	registerFunc(&_cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier, &_cFLocaleGetWindowsLocaleCodeFromLocaleIdentifierErr, frameworkHandle, "CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier", "10.6")
	registerFunc(&_cFMachPortCreate, &_cFMachPortCreateErr, frameworkHandle, "CFMachPortCreate", "")
	registerFunc(&_cFMachPortCreateRunLoopSource, &_cFMachPortCreateRunLoopSourceErr, frameworkHandle, "CFMachPortCreateRunLoopSource", "")
	registerFunc(&_cFMachPortCreateWithPort, &_cFMachPortCreateWithPortErr, frameworkHandle, "CFMachPortCreateWithPort", "")
	registerFunc(&_cFMachPortGetContext, &_cFMachPortGetContextErr, frameworkHandle, "CFMachPortGetContext", "")
	registerFunc(&_cFMachPortGetInvalidationCallBack, &_cFMachPortGetInvalidationCallBackErr, frameworkHandle, "CFMachPortGetInvalidationCallBack", "")
	registerFunc(&_cFMachPortGetPort, &_cFMachPortGetPortErr, frameworkHandle, "CFMachPortGetPort", "")
	registerFunc(&_cFMachPortGetTypeID, &_cFMachPortGetTypeIDErr, frameworkHandle, "CFMachPortGetTypeID", "")
	registerFunc(&_cFMachPortInvalidate, &_cFMachPortInvalidateErr, frameworkHandle, "CFMachPortInvalidate", "")
	registerFunc(&_cFMachPortIsValid, &_cFMachPortIsValidErr, frameworkHandle, "CFMachPortIsValid", "")
	registerFunc(&_cFMachPortSetInvalidationCallBack, &_cFMachPortSetInvalidationCallBackErr, frameworkHandle, "CFMachPortSetInvalidationCallBack", "")
	registerFunc(&_cFMakeCollectable, &_cFMakeCollectableErr, frameworkHandle, "CFMakeCollectable", "")
	registerFunc(&_cFMessagePortCreateLocal, &_cFMessagePortCreateLocalErr, frameworkHandle, "CFMessagePortCreateLocal", "")
	registerFunc(&_cFMessagePortCreateRemote, &_cFMessagePortCreateRemoteErr, frameworkHandle, "CFMessagePortCreateRemote", "")
	registerFunc(&_cFMessagePortCreateRunLoopSource, &_cFMessagePortCreateRunLoopSourceErr, frameworkHandle, "CFMessagePortCreateRunLoopSource", "")
	registerFunc(&_cFMessagePortGetContext, &_cFMessagePortGetContextErr, frameworkHandle, "CFMessagePortGetContext", "")
	registerFunc(&_cFMessagePortGetInvalidationCallBack, &_cFMessagePortGetInvalidationCallBackErr, frameworkHandle, "CFMessagePortGetInvalidationCallBack", "")
	registerFunc(&_cFMessagePortGetName, &_cFMessagePortGetNameErr, frameworkHandle, "CFMessagePortGetName", "")
	registerFunc(&_cFMessagePortGetTypeID, &_cFMessagePortGetTypeIDErr, frameworkHandle, "CFMessagePortGetTypeID", "")
	registerFunc(&_cFMessagePortInvalidate, &_cFMessagePortInvalidateErr, frameworkHandle, "CFMessagePortInvalidate", "")
	registerFunc(&_cFMessagePortIsRemote, &_cFMessagePortIsRemoteErr, frameworkHandle, "CFMessagePortIsRemote", "")
	registerFunc(&_cFMessagePortIsValid, &_cFMessagePortIsValidErr, frameworkHandle, "CFMessagePortIsValid", "")
	registerFunc(&_cFMessagePortSendRequest, &_cFMessagePortSendRequestErr, frameworkHandle, "CFMessagePortSendRequest", "")
	registerFunc(&_cFMessagePortSetDispatchQueue, &_cFMessagePortSetDispatchQueueErr, frameworkHandle, "CFMessagePortSetDispatchQueue", "10.6")
	registerFunc(&_cFMessagePortSetInvalidationCallBack, &_cFMessagePortSetInvalidationCallBackErr, frameworkHandle, "CFMessagePortSetInvalidationCallBack", "")
	registerFunc(&_cFMessagePortSetName, &_cFMessagePortSetNameErr, frameworkHandle, "CFMessagePortSetName", "")
	registerFunc(&_cFNotificationCenterAddObserver, &_cFNotificationCenterAddObserverErr, frameworkHandle, "CFNotificationCenterAddObserver", "")
	registerFunc(&_cFNotificationCenterGetDarwinNotifyCenter, &_cFNotificationCenterGetDarwinNotifyCenterErr, frameworkHandle, "CFNotificationCenterGetDarwinNotifyCenter", "")
	registerFunc(&_cFNotificationCenterGetDistributedCenter, &_cFNotificationCenterGetDistributedCenterErr, frameworkHandle, "CFNotificationCenterGetDistributedCenter", "")
	registerFunc(&_cFNotificationCenterGetLocalCenter, &_cFNotificationCenterGetLocalCenterErr, frameworkHandle, "CFNotificationCenterGetLocalCenter", "")
	registerFunc(&_cFNotificationCenterGetTypeID, &_cFNotificationCenterGetTypeIDErr, frameworkHandle, "CFNotificationCenterGetTypeID", "")
	registerFunc(&_cFNotificationCenterPostNotification, &_cFNotificationCenterPostNotificationErr, frameworkHandle, "CFNotificationCenterPostNotification", "")
	registerFunc(&_cFNotificationCenterPostNotificationWithOptions, &_cFNotificationCenterPostNotificationWithOptionsErr, frameworkHandle, "CFNotificationCenterPostNotificationWithOptions", "")
	registerFunc(&_cFNotificationCenterRemoveEveryObserver, &_cFNotificationCenterRemoveEveryObserverErr, frameworkHandle, "CFNotificationCenterRemoveEveryObserver", "")
	registerFunc(&_cFNotificationCenterRemoveObserver, &_cFNotificationCenterRemoveObserverErr, frameworkHandle, "CFNotificationCenterRemoveObserver", "")
	registerFunc(&_cFNullGetTypeID, &_cFNullGetTypeIDErr, frameworkHandle, "CFNullGetTypeID", "")
	registerFunc(&_cFNumberCompare, &_cFNumberCompareErr, frameworkHandle, "CFNumberCompare", "")
	registerFunc(&_cFNumberCreate, &_cFNumberCreateErr, frameworkHandle, "CFNumberCreate", "")
	registerFunc(&_cFNumberFormatterCopyProperty, &_cFNumberFormatterCopyPropertyErr, frameworkHandle, "CFNumberFormatterCopyProperty", "")
	registerFunc(&_cFNumberFormatterCreate, &_cFNumberFormatterCreateErr, frameworkHandle, "CFNumberFormatterCreate", "")
	registerFunc(&_cFNumberFormatterCreateNumberFromString, &_cFNumberFormatterCreateNumberFromStringErr, frameworkHandle, "CFNumberFormatterCreateNumberFromString", "")
	registerFunc(&_cFNumberFormatterCreateStringWithNumber, &_cFNumberFormatterCreateStringWithNumberErr, frameworkHandle, "CFNumberFormatterCreateStringWithNumber", "")
	registerFunc(&_cFNumberFormatterCreateStringWithValue, &_cFNumberFormatterCreateStringWithValueErr, frameworkHandle, "CFNumberFormatterCreateStringWithValue", "")
	registerFunc(&_cFNumberFormatterGetDecimalInfoForCurrencyCode, &_cFNumberFormatterGetDecimalInfoForCurrencyCodeErr, frameworkHandle, "CFNumberFormatterGetDecimalInfoForCurrencyCode", "")
	registerFunc(&_cFNumberFormatterGetFormat, &_cFNumberFormatterGetFormatErr, frameworkHandle, "CFNumberFormatterGetFormat", "")
	registerFunc(&_cFNumberFormatterGetLocale, &_cFNumberFormatterGetLocaleErr, frameworkHandle, "CFNumberFormatterGetLocale", "")
	registerFunc(&_cFNumberFormatterGetStyle, &_cFNumberFormatterGetStyleErr, frameworkHandle, "CFNumberFormatterGetStyle", "")
	registerFunc(&_cFNumberFormatterGetTypeID, &_cFNumberFormatterGetTypeIDErr, frameworkHandle, "CFNumberFormatterGetTypeID", "")
	registerFunc(&_cFNumberFormatterGetValueFromString, &_cFNumberFormatterGetValueFromStringErr, frameworkHandle, "CFNumberFormatterGetValueFromString", "")
	registerFunc(&_cFNumberFormatterSetFormat, &_cFNumberFormatterSetFormatErr, frameworkHandle, "CFNumberFormatterSetFormat", "")
	registerFunc(&_cFNumberFormatterSetProperty, &_cFNumberFormatterSetPropertyErr, frameworkHandle, "CFNumberFormatterSetProperty", "")
	registerFunc(&_cFNumberGetByteSize, &_cFNumberGetByteSizeErr, frameworkHandle, "CFNumberGetByteSize", "")
	registerFunc(&_cFNumberGetType, &_cFNumberGetTypeErr, frameworkHandle, "CFNumberGetType", "")
	registerFunc(&_cFNumberGetTypeID, &_cFNumberGetTypeIDErr, frameworkHandle, "CFNumberGetTypeID", "")
	registerFunc(&_cFNumberGetValue, &_cFNumberGetValueErr, frameworkHandle, "CFNumberGetValue", "")
	registerFunc(&_cFNumberIsFloatType, &_cFNumberIsFloatTypeErr, frameworkHandle, "CFNumberIsFloatType", "")
	registerFunc(&_cFPlugInAddInstanceForFactory, &_cFPlugInAddInstanceForFactoryErr, frameworkHandle, "CFPlugInAddInstanceForFactory", "")
	registerFunc(&_cFPlugInCreate, &_cFPlugInCreateErr, frameworkHandle, "CFPlugInCreate", "")
	registerFunc(&_cFPlugInFindFactoriesForPlugInType, &_cFPlugInFindFactoriesForPlugInTypeErr, frameworkHandle, "CFPlugInFindFactoriesForPlugInType", "")
	registerFunc(&_cFPlugInFindFactoriesForPlugInTypeInPlugIn, &_cFPlugInFindFactoriesForPlugInTypeInPlugInErr, frameworkHandle, "CFPlugInFindFactoriesForPlugInTypeInPlugIn", "")
	registerFunc(&_cFPlugInGetBundle, &_cFPlugInGetBundleErr, frameworkHandle, "CFPlugInGetBundle", "")
	registerFunc(&_cFPlugInGetTypeID, &_cFPlugInGetTypeIDErr, frameworkHandle, "CFPlugInGetTypeID", "")
	registerFunc(&_cFPlugInInstanceCreate, &_cFPlugInInstanceCreateErr, frameworkHandle, "CFPlugInInstanceCreate", "")
	registerFunc(&_cFPlugInInstanceCreateWithInstanceDataSize, &_cFPlugInInstanceCreateWithInstanceDataSizeErr, frameworkHandle, "CFPlugInInstanceCreateWithInstanceDataSize", "")
	registerFunc(&_cFPlugInInstanceGetFactoryName, &_cFPlugInInstanceGetFactoryNameErr, frameworkHandle, "CFPlugInInstanceGetFactoryName", "")
	registerFunc(&_cFPlugInInstanceGetInstanceData, &_cFPlugInInstanceGetInstanceDataErr, frameworkHandle, "CFPlugInInstanceGetInstanceData", "")
	registerFunc(&_cFPlugInInstanceGetInterfaceFunctionTable, &_cFPlugInInstanceGetInterfaceFunctionTableErr, frameworkHandle, "CFPlugInInstanceGetInterfaceFunctionTable", "")
	registerFunc(&_cFPlugInInstanceGetTypeID, &_cFPlugInInstanceGetTypeIDErr, frameworkHandle, "CFPlugInInstanceGetTypeID", "")
	registerFunc(&_cFPlugInIsLoadOnDemand, &_cFPlugInIsLoadOnDemandErr, frameworkHandle, "CFPlugInIsLoadOnDemand", "")
	registerFunc(&_cFPlugInRegisterFactoryFunction, &_cFPlugInRegisterFactoryFunctionErr, frameworkHandle, "CFPlugInRegisterFactoryFunction", "")
	registerFunc(&_cFPlugInRegisterFactoryFunctionByName, &_cFPlugInRegisterFactoryFunctionByNameErr, frameworkHandle, "CFPlugInRegisterFactoryFunctionByName", "")
	registerFunc(&_cFPlugInRegisterPlugInType, &_cFPlugInRegisterPlugInTypeErr, frameworkHandle, "CFPlugInRegisterPlugInType", "")
	registerFunc(&_cFPlugInRemoveInstanceForFactory, &_cFPlugInRemoveInstanceForFactoryErr, frameworkHandle, "CFPlugInRemoveInstanceForFactory", "")
	registerFunc(&_cFPlugInSetLoadOnDemand, &_cFPlugInSetLoadOnDemandErr, frameworkHandle, "CFPlugInSetLoadOnDemand", "")
	registerFunc(&_cFPlugInUnregisterFactory, &_cFPlugInUnregisterFactoryErr, frameworkHandle, "CFPlugInUnregisterFactory", "")
	registerFunc(&_cFPlugInUnregisterPlugInType, &_cFPlugInUnregisterPlugInTypeErr, frameworkHandle, "CFPlugInUnregisterPlugInType", "")
	registerFunc(&_cFPreferencesAddSuitePreferencesToApp, &_cFPreferencesAddSuitePreferencesToAppErr, frameworkHandle, "CFPreferencesAddSuitePreferencesToApp", "")
	registerFunc(&_cFPreferencesAppSynchronize, &_cFPreferencesAppSynchronizeErr, frameworkHandle, "CFPreferencesAppSynchronize", "")
	registerFunc(&_cFPreferencesAppValueIsForced, &_cFPreferencesAppValueIsForcedErr, frameworkHandle, "CFPreferencesAppValueIsForced", "")
	registerFunc(&_cFPreferencesCopyAppValue, &_cFPreferencesCopyAppValueErr, frameworkHandle, "CFPreferencesCopyAppValue", "")
	registerFunc(&_cFPreferencesCopyKeyList, &_cFPreferencesCopyKeyListErr, frameworkHandle, "CFPreferencesCopyKeyList", "")
	registerFunc(&_cFPreferencesCopyMultiple, &_cFPreferencesCopyMultipleErr, frameworkHandle, "CFPreferencesCopyMultiple", "")
	registerFunc(&_cFPreferencesCopyValue, &_cFPreferencesCopyValueErr, frameworkHandle, "CFPreferencesCopyValue", "")
	registerFunc(&_cFPreferencesGetAppBooleanValue, &_cFPreferencesGetAppBooleanValueErr, frameworkHandle, "CFPreferencesGetAppBooleanValue", "")
	registerFunc(&_cFPreferencesGetAppIntegerValue, &_cFPreferencesGetAppIntegerValueErr, frameworkHandle, "CFPreferencesGetAppIntegerValue", "")
	registerFunc(&_cFPreferencesRemoveSuitePreferencesFromApp, &_cFPreferencesRemoveSuitePreferencesFromAppErr, frameworkHandle, "CFPreferencesRemoveSuitePreferencesFromApp", "")
	registerFunc(&_cFPreferencesSetAppValue, &_cFPreferencesSetAppValueErr, frameworkHandle, "CFPreferencesSetAppValue", "")
	registerFunc(&_cFPreferencesSetMultiple, &_cFPreferencesSetMultipleErr, frameworkHandle, "CFPreferencesSetMultiple", "")
	registerFunc(&_cFPreferencesSetValue, &_cFPreferencesSetValueErr, frameworkHandle, "CFPreferencesSetValue", "")
	registerFunc(&_cFPreferencesSynchronize, &_cFPreferencesSynchronizeErr, frameworkHandle, "CFPreferencesSynchronize", "")
	registerFunc(&_cFPropertyListCreateData, &_cFPropertyListCreateDataErr, frameworkHandle, "CFPropertyListCreateData", "10.6")
	registerFunc(&_cFPropertyListCreateDeepCopy, &_cFPropertyListCreateDeepCopyErr, frameworkHandle, "CFPropertyListCreateDeepCopy", "")
	registerFunc(&_cFPropertyListCreateWithData, &_cFPropertyListCreateWithDataErr, frameworkHandle, "CFPropertyListCreateWithData", "10.6")
	registerFunc(&_cFPropertyListCreateWithStream, &_cFPropertyListCreateWithStreamErr, frameworkHandle, "CFPropertyListCreateWithStream", "10.6")
	registerFunc(&_cFPropertyListIsValid, &_cFPropertyListIsValidErr, frameworkHandle, "CFPropertyListIsValid", "")
	registerFunc(&_cFPropertyListWrite, &_cFPropertyListWriteErr, frameworkHandle, "CFPropertyListWrite", "10.6")
	registerFunc(&_cFReadStreamClose, &_cFReadStreamCloseErr, frameworkHandle, "CFReadStreamClose", "")
	registerFunc(&_cFReadStreamCopyDispatchQueue, &_cFReadStreamCopyDispatchQueueErr, frameworkHandle, "CFReadStreamCopyDispatchQueue", "10.9")
	registerFunc(&_cFReadStreamCopyError, &_cFReadStreamCopyErrorErr, frameworkHandle, "CFReadStreamCopyError", "10.5")
	registerFunc(&_cFReadStreamCopyProperty, &_cFReadStreamCopyPropertyErr, frameworkHandle, "CFReadStreamCopyProperty", "")
	registerFunc(&_cFReadStreamCreateWithBytesNoCopy, &_cFReadStreamCreateWithBytesNoCopyErr, frameworkHandle, "CFReadStreamCreateWithBytesNoCopy", "")
	registerFunc(&_cFReadStreamCreateWithFile, &_cFReadStreamCreateWithFileErr, frameworkHandle, "CFReadStreamCreateWithFile", "")
	registerFunc(&_cFReadStreamGetBuffer, &_cFReadStreamGetBufferErr, frameworkHandle, "CFReadStreamGetBuffer", "")
	registerFunc(&_cFReadStreamGetError, &_cFReadStreamGetErrorErr, frameworkHandle, "CFReadStreamGetError", "")
	registerFunc(&_cFReadStreamGetStatus, &_cFReadStreamGetStatusErr, frameworkHandle, "CFReadStreamGetStatus", "")
	registerFunc(&_cFReadStreamGetTypeID, &_cFReadStreamGetTypeIDErr, frameworkHandle, "CFReadStreamGetTypeID", "")
	registerFunc(&_cFReadStreamHasBytesAvailable, &_cFReadStreamHasBytesAvailableErr, frameworkHandle, "CFReadStreamHasBytesAvailable", "")
	registerFunc(&_cFReadStreamOpen, &_cFReadStreamOpenErr, frameworkHandle, "CFReadStreamOpen", "")
	registerFunc(&_cFReadStreamRead, &_cFReadStreamReadErr, frameworkHandle, "CFReadStreamRead", "")
	registerFunc(&_cFReadStreamScheduleWithRunLoop, &_cFReadStreamScheduleWithRunLoopErr, frameworkHandle, "CFReadStreamScheduleWithRunLoop", "")
	registerFunc(&_cFReadStreamSetClient, &_cFReadStreamSetClientErr, frameworkHandle, "CFReadStreamSetClient", "")
	registerFunc(&_cFReadStreamSetDispatchQueue, &_cFReadStreamSetDispatchQueueErr, frameworkHandle, "CFReadStreamSetDispatchQueue", "10.9")
	registerFunc(&_cFReadStreamSetProperty, &_cFReadStreamSetPropertyErr, frameworkHandle, "CFReadStreamSetProperty", "")
	registerFunc(&_cFReadStreamUnscheduleFromRunLoop, &_cFReadStreamUnscheduleFromRunLoopErr, frameworkHandle, "CFReadStreamUnscheduleFromRunLoop", "")
	registerFunc(&_cFRelease, &_cFReleaseErr, frameworkHandle, "CFRelease", "")
	registerFunc(&_cFRetain, &_cFRetainErr, frameworkHandle, "CFRetain", "")
	registerFunc(&_cFRunLoopAddCommonMode, &_cFRunLoopAddCommonModeErr, frameworkHandle, "CFRunLoopAddCommonMode", "")
	registerFunc(&_cFRunLoopAddObserver, &_cFRunLoopAddObserverErr, frameworkHandle, "CFRunLoopAddObserver", "")
	registerFunc(&_cFRunLoopAddSource, &_cFRunLoopAddSourceErr, frameworkHandle, "CFRunLoopAddSource", "")
	registerFunc(&_cFRunLoopAddTimer, &_cFRunLoopAddTimerErr, frameworkHandle, "CFRunLoopAddTimer", "")
	registerFunc(&_cFRunLoopContainsObserver, &_cFRunLoopContainsObserverErr, frameworkHandle, "CFRunLoopContainsObserver", "")
	registerFunc(&_cFRunLoopContainsSource, &_cFRunLoopContainsSourceErr, frameworkHandle, "CFRunLoopContainsSource", "")
	registerFunc(&_cFRunLoopContainsTimer, &_cFRunLoopContainsTimerErr, frameworkHandle, "CFRunLoopContainsTimer", "")
	registerFunc(&_cFRunLoopCopyAllModes, &_cFRunLoopCopyAllModesErr, frameworkHandle, "CFRunLoopCopyAllModes", "")
	registerFunc(&_cFRunLoopCopyCurrentMode, &_cFRunLoopCopyCurrentModeErr, frameworkHandle, "CFRunLoopCopyCurrentMode", "")
	registerFunc(&_cFRunLoopGetCurrent, &_cFRunLoopGetCurrentErr, frameworkHandle, "CFRunLoopGetCurrent", "")
	registerFunc(&_cFRunLoopGetMain, &_cFRunLoopGetMainErr, frameworkHandle, "CFRunLoopGetMain", "")
	registerFunc(&_cFRunLoopGetNextTimerFireDate, &_cFRunLoopGetNextTimerFireDateErr, frameworkHandle, "CFRunLoopGetNextTimerFireDate", "")
	registerFunc(&_cFRunLoopGetTypeID, &_cFRunLoopGetTypeIDErr, frameworkHandle, "CFRunLoopGetTypeID", "")
	registerFunc(&_cFRunLoopIsWaiting, &_cFRunLoopIsWaitingErr, frameworkHandle, "CFRunLoopIsWaiting", "")
	registerFunc(&_cFRunLoopObserverCreate, &_cFRunLoopObserverCreateErr, frameworkHandle, "CFRunLoopObserverCreate", "")
	registerFunc(&_cFRunLoopObserverCreateWithHandler, &_cFRunLoopObserverCreateWithHandlerErr, frameworkHandle, "CFRunLoopObserverCreateWithHandler", "10.7")
	registerFunc(&_cFRunLoopObserverDoesRepeat, &_cFRunLoopObserverDoesRepeatErr, frameworkHandle, "CFRunLoopObserverDoesRepeat", "")
	registerFunc(&_cFRunLoopObserverGetActivities, &_cFRunLoopObserverGetActivitiesErr, frameworkHandle, "CFRunLoopObserverGetActivities", "")
	registerFunc(&_cFRunLoopObserverGetContext, &_cFRunLoopObserverGetContextErr, frameworkHandle, "CFRunLoopObserverGetContext", "")
	registerFunc(&_cFRunLoopObserverGetOrder, &_cFRunLoopObserverGetOrderErr, frameworkHandle, "CFRunLoopObserverGetOrder", "")
	registerFunc(&_cFRunLoopObserverGetTypeID, &_cFRunLoopObserverGetTypeIDErr, frameworkHandle, "CFRunLoopObserverGetTypeID", "")
	registerFunc(&_cFRunLoopObserverInvalidate, &_cFRunLoopObserverInvalidateErr, frameworkHandle, "CFRunLoopObserverInvalidate", "")
	registerFunc(&_cFRunLoopObserverIsValid, &_cFRunLoopObserverIsValidErr, frameworkHandle, "CFRunLoopObserverIsValid", "")
	registerFunc(&_cFRunLoopPerformBlock, &_cFRunLoopPerformBlockErr, frameworkHandle, "CFRunLoopPerformBlock", "10.6")
	registerFunc(&_cFRunLoopRemoveObserver, &_cFRunLoopRemoveObserverErr, frameworkHandle, "CFRunLoopRemoveObserver", "")
	registerFunc(&_cFRunLoopRemoveSource, &_cFRunLoopRemoveSourceErr, frameworkHandle, "CFRunLoopRemoveSource", "")
	registerFunc(&_cFRunLoopRemoveTimer, &_cFRunLoopRemoveTimerErr, frameworkHandle, "CFRunLoopRemoveTimer", "")
	registerFunc(&_cFRunLoopRun, &_cFRunLoopRunErr, frameworkHandle, "CFRunLoopRun", "")
	registerFunc(&_cFRunLoopRunInMode, &_cFRunLoopRunInModeErr, frameworkHandle, "CFRunLoopRunInMode", "")
	registerFunc(&_cFRunLoopSourceCreate, &_cFRunLoopSourceCreateErr, frameworkHandle, "CFRunLoopSourceCreate", "")
	registerFunc(&_cFRunLoopSourceGetContext, &_cFRunLoopSourceGetContextErr, frameworkHandle, "CFRunLoopSourceGetContext", "")
	registerFunc(&_cFRunLoopSourceGetOrder, &_cFRunLoopSourceGetOrderErr, frameworkHandle, "CFRunLoopSourceGetOrder", "")
	registerFunc(&_cFRunLoopSourceGetTypeID, &_cFRunLoopSourceGetTypeIDErr, frameworkHandle, "CFRunLoopSourceGetTypeID", "")
	registerFunc(&_cFRunLoopSourceInvalidate, &_cFRunLoopSourceInvalidateErr, frameworkHandle, "CFRunLoopSourceInvalidate", "")
	registerFunc(&_cFRunLoopSourceIsValid, &_cFRunLoopSourceIsValidErr, frameworkHandle, "CFRunLoopSourceIsValid", "")
	registerFunc(&_cFRunLoopSourceSignal, &_cFRunLoopSourceSignalErr, frameworkHandle, "CFRunLoopSourceSignal", "")
	registerFunc(&_cFRunLoopStop, &_cFRunLoopStopErr, frameworkHandle, "CFRunLoopStop", "")
	registerFunc(&_cFRunLoopTimerCreate, &_cFRunLoopTimerCreateErr, frameworkHandle, "CFRunLoopTimerCreate", "")
	registerFunc(&_cFRunLoopTimerCreateWithHandler, &_cFRunLoopTimerCreateWithHandlerErr, frameworkHandle, "CFRunLoopTimerCreateWithHandler", "10.7")
	registerFunc(&_cFRunLoopTimerDoesRepeat, &_cFRunLoopTimerDoesRepeatErr, frameworkHandle, "CFRunLoopTimerDoesRepeat", "")
	registerFunc(&_cFRunLoopTimerGetContext, &_cFRunLoopTimerGetContextErr, frameworkHandle, "CFRunLoopTimerGetContext", "")
	registerFunc(&_cFRunLoopTimerGetInterval, &_cFRunLoopTimerGetIntervalErr, frameworkHandle, "CFRunLoopTimerGetInterval", "")
	registerFunc(&_cFRunLoopTimerGetNextFireDate, &_cFRunLoopTimerGetNextFireDateErr, frameworkHandle, "CFRunLoopTimerGetNextFireDate", "")
	registerFunc(&_cFRunLoopTimerGetOrder, &_cFRunLoopTimerGetOrderErr, frameworkHandle, "CFRunLoopTimerGetOrder", "")
	registerFunc(&_cFRunLoopTimerGetTolerance, &_cFRunLoopTimerGetToleranceErr, frameworkHandle, "CFRunLoopTimerGetTolerance", "10.9")
	registerFunc(&_cFRunLoopTimerGetTypeID, &_cFRunLoopTimerGetTypeIDErr, frameworkHandle, "CFRunLoopTimerGetTypeID", "")
	registerFunc(&_cFRunLoopTimerInvalidate, &_cFRunLoopTimerInvalidateErr, frameworkHandle, "CFRunLoopTimerInvalidate", "")
	registerFunc(&_cFRunLoopTimerIsValid, &_cFRunLoopTimerIsValidErr, frameworkHandle, "CFRunLoopTimerIsValid", "")
	registerFunc(&_cFRunLoopTimerSetNextFireDate, &_cFRunLoopTimerSetNextFireDateErr, frameworkHandle, "CFRunLoopTimerSetNextFireDate", "")
	registerFunc(&_cFRunLoopTimerSetTolerance, &_cFRunLoopTimerSetToleranceErr, frameworkHandle, "CFRunLoopTimerSetTolerance", "10.9")
	registerFunc(&_cFRunLoopWakeUp, &_cFRunLoopWakeUpErr, frameworkHandle, "CFRunLoopWakeUp", "")
	registerFunc(&_cFSetAddValue, &_cFSetAddValueErr, frameworkHandle, "CFSetAddValue", "")
	registerFunc(&_cFSetApplyFunction, &_cFSetApplyFunctionErr, frameworkHandle, "CFSetApplyFunction", "")
	registerFunc(&_cFSetContainsValue, &_cFSetContainsValueErr, frameworkHandle, "CFSetContainsValue", "")
	registerFunc(&_cFSetCreate, &_cFSetCreateErr, frameworkHandle, "CFSetCreate", "")
	registerFunc(&_cFSetCreateCopy, &_cFSetCreateCopyErr, frameworkHandle, "CFSetCreateCopy", "")
	registerFunc(&_cFSetCreateMutable, &_cFSetCreateMutableErr, frameworkHandle, "CFSetCreateMutable", "")
	registerFunc(&_cFSetCreateMutableCopy, &_cFSetCreateMutableCopyErr, frameworkHandle, "CFSetCreateMutableCopy", "")
	registerFunc(&_cFSetGetCount, &_cFSetGetCountErr, frameworkHandle, "CFSetGetCount", "")
	registerFunc(&_cFSetGetCountOfValue, &_cFSetGetCountOfValueErr, frameworkHandle, "CFSetGetCountOfValue", "")
	registerFunc(&_cFSetGetTypeID, &_cFSetGetTypeIDErr, frameworkHandle, "CFSetGetTypeID", "")
	registerFunc(&_cFSetGetValue, &_cFSetGetValueErr, frameworkHandle, "CFSetGetValue", "")
	registerFunc(&_cFSetGetValueIfPresent, &_cFSetGetValueIfPresentErr, frameworkHandle, "CFSetGetValueIfPresent", "")
	registerFunc(&_cFSetGetValues, &_cFSetGetValuesErr, frameworkHandle, "CFSetGetValues", "")
	registerFunc(&_cFSetRemoveAllValues, &_cFSetRemoveAllValuesErr, frameworkHandle, "CFSetRemoveAllValues", "")
	registerFunc(&_cFSetRemoveValue, &_cFSetRemoveValueErr, frameworkHandle, "CFSetRemoveValue", "")
	registerFunc(&_cFSetReplaceValue, &_cFSetReplaceValueErr, frameworkHandle, "CFSetReplaceValue", "")
	registerFunc(&_cFSetSetValue, &_cFSetSetValueErr, frameworkHandle, "CFSetSetValue", "")
	registerFunc(&_cFShow, &_cFShowErr, frameworkHandle, "CFShow", "")
	registerFunc(&_cFShowStr, &_cFShowStrErr, frameworkHandle, "CFShowStr", "")
	registerFunc(&_cFSocketConnectToAddress, &_cFSocketConnectToAddressErr, frameworkHandle, "CFSocketConnectToAddress", "")
	registerFunc(&_cFSocketCopyAddress, &_cFSocketCopyAddressErr, frameworkHandle, "CFSocketCopyAddress", "")
	registerFunc(&_cFSocketCopyPeerAddress, &_cFSocketCopyPeerAddressErr, frameworkHandle, "CFSocketCopyPeerAddress", "")
	registerFunc(&_cFSocketCopyRegisteredSocketSignature, &_cFSocketCopyRegisteredSocketSignatureErr, frameworkHandle, "CFSocketCopyRegisteredSocketSignature", "")
	registerFunc(&_cFSocketCopyRegisteredValue, &_cFSocketCopyRegisteredValueErr, frameworkHandle, "CFSocketCopyRegisteredValue", "")
	registerFunc(&_cFSocketCreate, &_cFSocketCreateErr, frameworkHandle, "CFSocketCreate", "")
	registerFunc(&_cFSocketCreateConnectedToSocketSignature, &_cFSocketCreateConnectedToSocketSignatureErr, frameworkHandle, "CFSocketCreateConnectedToSocketSignature", "")
	registerFunc(&_cFSocketCreateRunLoopSource, &_cFSocketCreateRunLoopSourceErr, frameworkHandle, "CFSocketCreateRunLoopSource", "")
	registerFunc(&_cFSocketCreateWithNative, &_cFSocketCreateWithNativeErr, frameworkHandle, "CFSocketCreateWithNative", "")
	registerFunc(&_cFSocketCreateWithSocketSignature, &_cFSocketCreateWithSocketSignatureErr, frameworkHandle, "CFSocketCreateWithSocketSignature", "")
	registerFunc(&_cFSocketDisableCallBacks, &_cFSocketDisableCallBacksErr, frameworkHandle, "CFSocketDisableCallBacks", "")
	registerFunc(&_cFSocketEnableCallBacks, &_cFSocketEnableCallBacksErr, frameworkHandle, "CFSocketEnableCallBacks", "")
	registerFunc(&_cFSocketGetContext, &_cFSocketGetContextErr, frameworkHandle, "CFSocketGetContext", "")
	registerFunc(&_cFSocketGetDefaultNameRegistryPortNumber, &_cFSocketGetDefaultNameRegistryPortNumberErr, frameworkHandle, "CFSocketGetDefaultNameRegistryPortNumber", "")
	registerFunc(&_cFSocketGetNative, &_cFSocketGetNativeErr, frameworkHandle, "CFSocketGetNative", "")
	registerFunc(&_cFSocketGetSocketFlags, &_cFSocketGetSocketFlagsErr, frameworkHandle, "CFSocketGetSocketFlags", "")
	registerFunc(&_cFSocketGetTypeID, &_cFSocketGetTypeIDErr, frameworkHandle, "CFSocketGetTypeID", "")
	registerFunc(&_cFSocketInvalidate, &_cFSocketInvalidateErr, frameworkHandle, "CFSocketInvalidate", "")
	registerFunc(&_cFSocketIsValid, &_cFSocketIsValidErr, frameworkHandle, "CFSocketIsValid", "")
	registerFunc(&_cFSocketRegisterSocketSignature, &_cFSocketRegisterSocketSignatureErr, frameworkHandle, "CFSocketRegisterSocketSignature", "")
	registerFunc(&_cFSocketRegisterValue, &_cFSocketRegisterValueErr, frameworkHandle, "CFSocketRegisterValue", "")
	registerFunc(&_cFSocketSendData, &_cFSocketSendDataErr, frameworkHandle, "CFSocketSendData", "")
	registerFunc(&_cFSocketSetAddress, &_cFSocketSetAddressErr, frameworkHandle, "CFSocketSetAddress", "")
	registerFunc(&_cFSocketSetDefaultNameRegistryPortNumber, &_cFSocketSetDefaultNameRegistryPortNumberErr, frameworkHandle, "CFSocketSetDefaultNameRegistryPortNumber", "")
	registerFunc(&_cFSocketSetSocketFlags, &_cFSocketSetSocketFlagsErr, frameworkHandle, "CFSocketSetSocketFlags", "")
	registerFunc(&_cFSocketUnregister, &_cFSocketUnregisterErr, frameworkHandle, "CFSocketUnregister", "")
	registerFunc(&_cFStreamCreateBoundPair, &_cFStreamCreateBoundPairErr, frameworkHandle, "CFStreamCreateBoundPair", "")
	registerFunc(&_cFStreamCreatePairWithPeerSocketSignature, &_cFStreamCreatePairWithPeerSocketSignatureErr, frameworkHandle, "CFStreamCreatePairWithPeerSocketSignature", "10.1")
	registerFunc(&_cFStreamCreatePairWithSocket, &_cFStreamCreatePairWithSocketErr, frameworkHandle, "CFStreamCreatePairWithSocket", "10.1")
	registerFunc(&_cFStreamCreatePairWithSocketToHost, &_cFStreamCreatePairWithSocketToHostErr, frameworkHandle, "CFStreamCreatePairWithSocketToHost", "10.1")
	registerFunc(&_cFStringAppend, &_cFStringAppendErr, frameworkHandle, "CFStringAppend", "")
	registerFunc(&_cFStringAppendCString, &_cFStringAppendCStringErr, frameworkHandle, "CFStringAppendCString", "")
	registerFunc(&_cFStringAppendCharacters, &_cFStringAppendCharactersErr, frameworkHandle, "CFStringAppendCharacters", "")
	registerFunc(&_cFStringAppendFormat, &_cFStringAppendFormatErr, frameworkHandle, "CFStringAppendFormat", "")
	registerFunc(&_cFStringAppendFormatAndArguments, &_cFStringAppendFormatAndArgumentsErr, frameworkHandle, "CFStringAppendFormatAndArguments", "")
	registerFunc(&_cFStringAppendPascalString, &_cFStringAppendPascalStringErr, frameworkHandle, "CFStringAppendPascalString", "")
	registerFunc(&_cFStringCapitalize, &_cFStringCapitalizeErr, frameworkHandle, "CFStringCapitalize", "")
	registerFunc(&_cFStringCompare, &_cFStringCompareErr, frameworkHandle, "CFStringCompare", "")
	registerFunc(&_cFStringCompareWithOptions, &_cFStringCompareWithOptionsErr, frameworkHandle, "CFStringCompareWithOptions", "")
	registerFunc(&_cFStringCompareWithOptionsAndLocale, &_cFStringCompareWithOptionsAndLocaleErr, frameworkHandle, "CFStringCompareWithOptionsAndLocale", "10.5")
	registerFunc(&_cFStringConvertEncodingToIANACharSetName, &_cFStringConvertEncodingToIANACharSetNameErr, frameworkHandle, "CFStringConvertEncodingToIANACharSetName", "")
	registerFunc(&_cFStringConvertEncodingToNSStringEncoding, &_cFStringConvertEncodingToNSStringEncodingErr, frameworkHandle, "CFStringConvertEncodingToNSStringEncoding", "")
	registerFunc(&_cFStringConvertEncodingToWindowsCodepage, &_cFStringConvertEncodingToWindowsCodepageErr, frameworkHandle, "CFStringConvertEncodingToWindowsCodepage", "")
	registerFunc(&_cFStringConvertIANACharSetNameToEncoding, &_cFStringConvertIANACharSetNameToEncodingErr, frameworkHandle, "CFStringConvertIANACharSetNameToEncoding", "")
	registerFunc(&_cFStringConvertNSStringEncodingToEncoding, &_cFStringConvertNSStringEncodingToEncodingErr, frameworkHandle, "CFStringConvertNSStringEncodingToEncoding", "")
	registerFunc(&_cFStringConvertWindowsCodepageToEncoding, &_cFStringConvertWindowsCodepageToEncodingErr, frameworkHandle, "CFStringConvertWindowsCodepageToEncoding", "")
	registerFunc(&_cFStringCreateArrayBySeparatingStrings, &_cFStringCreateArrayBySeparatingStringsErr, frameworkHandle, "CFStringCreateArrayBySeparatingStrings", "")
	registerFunc(&_cFStringCreateArrayWithFindResults, &_cFStringCreateArrayWithFindResultsErr, frameworkHandle, "CFStringCreateArrayWithFindResults", "")
	registerFunc(&_cFStringCreateByCombiningStrings, &_cFStringCreateByCombiningStringsErr, frameworkHandle, "CFStringCreateByCombiningStrings", "")
	registerFunc(&_cFStringCreateCopy, &_cFStringCreateCopyErr, frameworkHandle, "CFStringCreateCopy", "")
	registerFunc(&_cFStringCreateExternalRepresentation, &_cFStringCreateExternalRepresentationErr, frameworkHandle, "CFStringCreateExternalRepresentation", "")
	registerFunc(&_cFStringCreateFromExternalRepresentation, &_cFStringCreateFromExternalRepresentationErr, frameworkHandle, "CFStringCreateFromExternalRepresentation", "")
	registerFunc(&_cFStringCreateMutable, &_cFStringCreateMutableErr, frameworkHandle, "CFStringCreateMutable", "")
	registerFunc(&_cFStringCreateMutableCopy, &_cFStringCreateMutableCopyErr, frameworkHandle, "CFStringCreateMutableCopy", "")
	registerFunc(&_cFStringCreateMutableWithExternalCharactersNoCopy, &_cFStringCreateMutableWithExternalCharactersNoCopyErr, frameworkHandle, "CFStringCreateMutableWithExternalCharactersNoCopy", "")
	registerFunc(&_cFStringCreateStringWithValidatedFormat, &_cFStringCreateStringWithValidatedFormatErr, frameworkHandle, "CFStringCreateStringWithValidatedFormat", "13.0")
	registerFunc(&_cFStringCreateStringWithValidatedFormatAndArguments, &_cFStringCreateStringWithValidatedFormatAndArgumentsErr, frameworkHandle, "CFStringCreateStringWithValidatedFormatAndArguments", "13.0")
	registerFunc(&_cFStringCreateWithBytes, &_cFStringCreateWithBytesErr, frameworkHandle, "CFStringCreateWithBytes", "")
	registerFunc(&_cFStringCreateWithBytesNoCopy, &_cFStringCreateWithBytesNoCopyErr, frameworkHandle, "CFStringCreateWithBytesNoCopy", "")
	registerFunc(&_cFStringCreateWithCString, &_cFStringCreateWithCStringErr, frameworkHandle, "CFStringCreateWithCString", "")
	registerFunc(&_cFStringCreateWithCStringNoCopy, &_cFStringCreateWithCStringNoCopyErr, frameworkHandle, "CFStringCreateWithCStringNoCopy", "")
	registerFunc(&_cFStringCreateWithCharacters, &_cFStringCreateWithCharactersErr, frameworkHandle, "CFStringCreateWithCharacters", "")
	registerFunc(&_cFStringCreateWithCharactersNoCopy, &_cFStringCreateWithCharactersNoCopyErr, frameworkHandle, "CFStringCreateWithCharactersNoCopy", "")
	registerFunc(&_cFStringCreateWithFileSystemRepresentation, &_cFStringCreateWithFileSystemRepresentationErr, frameworkHandle, "CFStringCreateWithFileSystemRepresentation", "")
	registerFunc(&_cFStringCreateWithFormat, &_cFStringCreateWithFormatErr, frameworkHandle, "CFStringCreateWithFormat", "")
	registerFunc(&_cFStringCreateWithFormatAndArguments, &_cFStringCreateWithFormatAndArgumentsErr, frameworkHandle, "CFStringCreateWithFormatAndArguments", "")
	registerFunc(&_cFStringCreateWithPascalString, &_cFStringCreateWithPascalStringErr, frameworkHandle, "CFStringCreateWithPascalString", "")
	registerFunc(&_cFStringCreateWithPascalStringNoCopy, &_cFStringCreateWithPascalStringNoCopyErr, frameworkHandle, "CFStringCreateWithPascalStringNoCopy", "")
	registerFunc(&_cFStringCreateWithSubstring, &_cFStringCreateWithSubstringErr, frameworkHandle, "CFStringCreateWithSubstring", "")
	registerFunc(&_cFStringDelete, &_cFStringDeleteErr, frameworkHandle, "CFStringDelete", "")
	registerFunc(&_cFStringFind, &_cFStringFindErr, frameworkHandle, "CFStringFind", "")
	registerFunc(&_cFStringFindAndReplace, &_cFStringFindAndReplaceErr, frameworkHandle, "CFStringFindAndReplace", "")
	registerFunc(&_cFStringFindCharacterFromSet, &_cFStringFindCharacterFromSetErr, frameworkHandle, "CFStringFindCharacterFromSet", "")
	registerFunc(&_cFStringFindWithOptions, &_cFStringFindWithOptionsErr, frameworkHandle, "CFStringFindWithOptions", "")
	registerFunc(&_cFStringFindWithOptionsAndLocale, &_cFStringFindWithOptionsAndLocaleErr, frameworkHandle, "CFStringFindWithOptionsAndLocale", "10.5")
	registerFunc(&_cFStringFold, &_cFStringFoldErr, frameworkHandle, "CFStringFold", "10.5")
	registerFunc(&_cFStringGetBytes, &_cFStringGetBytesErr, frameworkHandle, "CFStringGetBytes", "")
	registerFunc(&_cFStringGetCString, &_cFStringGetCStringErr, frameworkHandle, "CFStringGetCString", "")
	registerFunc(&_cFStringGetCStringPtr, &_cFStringGetCStringPtrErr, frameworkHandle, "CFStringGetCStringPtr", "")
	registerFunc(&_cFStringGetCharacterAtIndex, &_cFStringGetCharacterAtIndexErr, frameworkHandle, "CFStringGetCharacterAtIndex", "")
	registerFunc(&_cFStringGetCharacters, &_cFStringGetCharactersErr, frameworkHandle, "CFStringGetCharacters", "")
	registerFunc(&_cFStringGetCharactersPtr, &_cFStringGetCharactersPtrErr, frameworkHandle, "CFStringGetCharactersPtr", "")
	registerFunc(&_cFStringGetDoubleValue, &_cFStringGetDoubleValueErr, frameworkHandle, "CFStringGetDoubleValue", "")
	registerFunc(&_cFStringGetFastestEncoding, &_cFStringGetFastestEncodingErr, frameworkHandle, "CFStringGetFastestEncoding", "")
	registerFunc(&_cFStringGetFileSystemRepresentation, &_cFStringGetFileSystemRepresentationErr, frameworkHandle, "CFStringGetFileSystemRepresentation", "")
	registerFunc(&_cFStringGetHyphenationLocationBeforeIndex, &_cFStringGetHyphenationLocationBeforeIndexErr, frameworkHandle, "CFStringGetHyphenationLocationBeforeIndex", "10.7")
	registerFunc(&_cFStringGetIntValue, &_cFStringGetIntValueErr, frameworkHandle, "CFStringGetIntValue", "")
	registerFunc(&_cFStringGetLength, &_cFStringGetLengthErr, frameworkHandle, "CFStringGetLength", "")
	registerFunc(&_cFStringGetLineBounds, &_cFStringGetLineBoundsErr, frameworkHandle, "CFStringGetLineBounds", "")
	registerFunc(&_cFStringGetListOfAvailableEncodings, &_cFStringGetListOfAvailableEncodingsErr, frameworkHandle, "CFStringGetListOfAvailableEncodings", "")
	registerFunc(&_cFStringGetMaximumSizeForEncoding, &_cFStringGetMaximumSizeForEncodingErr, frameworkHandle, "CFStringGetMaximumSizeForEncoding", "")
	registerFunc(&_cFStringGetMaximumSizeOfFileSystemRepresentation, &_cFStringGetMaximumSizeOfFileSystemRepresentationErr, frameworkHandle, "CFStringGetMaximumSizeOfFileSystemRepresentation", "")
	registerFunc(&_cFStringGetMostCompatibleMacStringEncoding, &_cFStringGetMostCompatibleMacStringEncodingErr, frameworkHandle, "CFStringGetMostCompatibleMacStringEncoding", "")
	registerFunc(&_cFStringGetNameOfEncoding, &_cFStringGetNameOfEncodingErr, frameworkHandle, "CFStringGetNameOfEncoding", "")
	registerFunc(&_cFStringGetParagraphBounds, &_cFStringGetParagraphBoundsErr, frameworkHandle, "CFStringGetParagraphBounds", "10.5")
	registerFunc(&_cFStringGetPascalString, &_cFStringGetPascalStringErr, frameworkHandle, "CFStringGetPascalString", "")
	registerFunc(&_cFStringGetPascalStringPtr, &_cFStringGetPascalStringPtrErr, frameworkHandle, "CFStringGetPascalStringPtr", "")
	registerFunc(&_cFStringGetRangeOfComposedCharactersAtIndex, &_cFStringGetRangeOfComposedCharactersAtIndexErr, frameworkHandle, "CFStringGetRangeOfComposedCharactersAtIndex", "")
	registerFunc(&_cFStringGetSmallestEncoding, &_cFStringGetSmallestEncodingErr, frameworkHandle, "CFStringGetSmallestEncoding", "")
	registerFunc(&_cFStringGetSystemEncoding, &_cFStringGetSystemEncodingErr, frameworkHandle, "CFStringGetSystemEncoding", "")
	registerFunc(&_cFStringGetTypeID, &_cFStringGetTypeIDErr, frameworkHandle, "CFStringGetTypeID", "")
	registerFunc(&_cFStringHasPrefix, &_cFStringHasPrefixErr, frameworkHandle, "CFStringHasPrefix", "")
	registerFunc(&_cFStringHasSuffix, &_cFStringHasSuffixErr, frameworkHandle, "CFStringHasSuffix", "")
	registerFunc(&_cFStringInsert, &_cFStringInsertErr, frameworkHandle, "CFStringInsert", "")
	registerFunc(&_cFStringIsEncodingAvailable, &_cFStringIsEncodingAvailableErr, frameworkHandle, "CFStringIsEncodingAvailable", "")
	registerFunc(&_cFStringIsHyphenationAvailableForLocale, &_cFStringIsHyphenationAvailableForLocaleErr, frameworkHandle, "CFStringIsHyphenationAvailableForLocale", "10.7")
	registerFunc(&_cFStringLowercase, &_cFStringLowercaseErr, frameworkHandle, "CFStringLowercase", "")
	registerFunc(&_cFStringNormalize, &_cFStringNormalizeErr, frameworkHandle, "CFStringNormalize", "")
	registerFunc(&_cFStringPad, &_cFStringPadErr, frameworkHandle, "CFStringPad", "")
	registerFunc(&_cFStringReplace, &_cFStringReplaceErr, frameworkHandle, "CFStringReplace", "")
	registerFunc(&_cFStringReplaceAll, &_cFStringReplaceAllErr, frameworkHandle, "CFStringReplaceAll", "")
	registerFunc(&_cFStringSetExternalCharactersNoCopy, &_cFStringSetExternalCharactersNoCopyErr, frameworkHandle, "CFStringSetExternalCharactersNoCopy", "")
	registerFunc(&_cFStringTokenizerAdvanceToNextToken, &_cFStringTokenizerAdvanceToNextTokenErr, frameworkHandle, "CFStringTokenizerAdvanceToNextToken", "10.5")
	registerFunc(&_cFStringTokenizerCopyBestStringLanguage, &_cFStringTokenizerCopyBestStringLanguageErr, frameworkHandle, "CFStringTokenizerCopyBestStringLanguage", "10.5")
	registerFunc(&_cFStringTokenizerCopyCurrentTokenAttribute, &_cFStringTokenizerCopyCurrentTokenAttributeErr, frameworkHandle, "CFStringTokenizerCopyCurrentTokenAttribute", "10.5")
	registerFunc(&_cFStringTokenizerCreate, &_cFStringTokenizerCreateErr, frameworkHandle, "CFStringTokenizerCreate", "10.5")
	registerFunc(&_cFStringTokenizerGetCurrentSubTokens, &_cFStringTokenizerGetCurrentSubTokensErr, frameworkHandle, "CFStringTokenizerGetCurrentSubTokens", "10.5")
	registerFunc(&_cFStringTokenizerGetCurrentTokenRange, &_cFStringTokenizerGetCurrentTokenRangeErr, frameworkHandle, "CFStringTokenizerGetCurrentTokenRange", "10.5")
	registerFunc(&_cFStringTokenizerGetTypeID, &_cFStringTokenizerGetTypeIDErr, frameworkHandle, "CFStringTokenizerGetTypeID", "10.5")
	registerFunc(&_cFStringTokenizerGoToTokenAtIndex, &_cFStringTokenizerGoToTokenAtIndexErr, frameworkHandle, "CFStringTokenizerGoToTokenAtIndex", "10.5")
	registerFunc(&_cFStringTokenizerSetString, &_cFStringTokenizerSetStringErr, frameworkHandle, "CFStringTokenizerSetString", "10.5")
	registerFunc(&_cFStringTransform, &_cFStringTransformErr, frameworkHandle, "CFStringTransform", "")
	registerFunc(&_cFStringTrim, &_cFStringTrimErr, frameworkHandle, "CFStringTrim", "")
	registerFunc(&_cFStringTrimWhitespace, &_cFStringTrimWhitespaceErr, frameworkHandle, "CFStringTrimWhitespace", "")
	registerFunc(&_cFStringUppercase, &_cFStringUppercaseErr, frameworkHandle, "CFStringUppercase", "")
	registerFunc(&_cFTimeZoneCopyAbbreviation, &_cFTimeZoneCopyAbbreviationErr, frameworkHandle, "CFTimeZoneCopyAbbreviation", "")
	registerFunc(&_cFTimeZoneCopyAbbreviationDictionary, &_cFTimeZoneCopyAbbreviationDictionaryErr, frameworkHandle, "CFTimeZoneCopyAbbreviationDictionary", "")
	registerFunc(&_cFTimeZoneCopyDefault, &_cFTimeZoneCopyDefaultErr, frameworkHandle, "CFTimeZoneCopyDefault", "")
	registerFunc(&_cFTimeZoneCopyKnownNames, &_cFTimeZoneCopyKnownNamesErr, frameworkHandle, "CFTimeZoneCopyKnownNames", "")
	registerFunc(&_cFTimeZoneCopyLocalizedName, &_cFTimeZoneCopyLocalizedNameErr, frameworkHandle, "CFTimeZoneCopyLocalizedName", "10.5")
	registerFunc(&_cFTimeZoneCopySystem, &_cFTimeZoneCopySystemErr, frameworkHandle, "CFTimeZoneCopySystem", "")
	registerFunc(&_cFTimeZoneCreate, &_cFTimeZoneCreateErr, frameworkHandle, "CFTimeZoneCreate", "")
	registerFunc(&_cFTimeZoneCreateWithName, &_cFTimeZoneCreateWithNameErr, frameworkHandle, "CFTimeZoneCreateWithName", "")
	registerFunc(&_cFTimeZoneCreateWithTimeIntervalFromGMT, &_cFTimeZoneCreateWithTimeIntervalFromGMTErr, frameworkHandle, "CFTimeZoneCreateWithTimeIntervalFromGMT", "")
	registerFunc(&_cFTimeZoneGetData, &_cFTimeZoneGetDataErr, frameworkHandle, "CFTimeZoneGetData", "")
	registerFunc(&_cFTimeZoneGetDaylightSavingTimeOffset, &_cFTimeZoneGetDaylightSavingTimeOffsetErr, frameworkHandle, "CFTimeZoneGetDaylightSavingTimeOffset", "10.5")
	registerFunc(&_cFTimeZoneGetName, &_cFTimeZoneGetNameErr, frameworkHandle, "CFTimeZoneGetName", "")
	registerFunc(&_cFTimeZoneGetNextDaylightSavingTimeTransition, &_cFTimeZoneGetNextDaylightSavingTimeTransitionErr, frameworkHandle, "CFTimeZoneGetNextDaylightSavingTimeTransition", "10.5")
	registerFunc(&_cFTimeZoneGetSecondsFromGMT, &_cFTimeZoneGetSecondsFromGMTErr, frameworkHandle, "CFTimeZoneGetSecondsFromGMT", "")
	registerFunc(&_cFTimeZoneGetTypeID, &_cFTimeZoneGetTypeIDErr, frameworkHandle, "CFTimeZoneGetTypeID", "")
	registerFunc(&_cFTimeZoneIsDaylightSavingTime, &_cFTimeZoneIsDaylightSavingTimeErr, frameworkHandle, "CFTimeZoneIsDaylightSavingTime", "")
	registerFunc(&_cFTimeZoneResetSystem, &_cFTimeZoneResetSystemErr, frameworkHandle, "CFTimeZoneResetSystem", "")
	registerFunc(&_cFTimeZoneSetAbbreviationDictionary, &_cFTimeZoneSetAbbreviationDictionaryErr, frameworkHandle, "CFTimeZoneSetAbbreviationDictionary", "")
	registerFunc(&_cFTimeZoneSetDefault, &_cFTimeZoneSetDefaultErr, frameworkHandle, "CFTimeZoneSetDefault", "")
	registerFunc(&_cFTreeAppendChild, &_cFTreeAppendChildErr, frameworkHandle, "CFTreeAppendChild", "")
	registerFunc(&_cFTreeApplyFunctionToChildren, &_cFTreeApplyFunctionToChildrenErr, frameworkHandle, "CFTreeApplyFunctionToChildren", "")
	registerFunc(&_cFTreeCreate, &_cFTreeCreateErr, frameworkHandle, "CFTreeCreate", "")
	registerFunc(&_cFTreeFindRoot, &_cFTreeFindRootErr, frameworkHandle, "CFTreeFindRoot", "")
	registerFunc(&_cFTreeGetChildAtIndex, &_cFTreeGetChildAtIndexErr, frameworkHandle, "CFTreeGetChildAtIndex", "")
	registerFunc(&_cFTreeGetChildCount, &_cFTreeGetChildCountErr, frameworkHandle, "CFTreeGetChildCount", "")
	registerFunc(&_cFTreeGetChildren, &_cFTreeGetChildrenErr, frameworkHandle, "CFTreeGetChildren", "")
	registerFunc(&_cFTreeGetContext, &_cFTreeGetContextErr, frameworkHandle, "CFTreeGetContext", "")
	registerFunc(&_cFTreeGetFirstChild, &_cFTreeGetFirstChildErr, frameworkHandle, "CFTreeGetFirstChild", "")
	registerFunc(&_cFTreeGetNextSibling, &_cFTreeGetNextSiblingErr, frameworkHandle, "CFTreeGetNextSibling", "")
	registerFunc(&_cFTreeGetParent, &_cFTreeGetParentErr, frameworkHandle, "CFTreeGetParent", "")
	registerFunc(&_cFTreeGetTypeID, &_cFTreeGetTypeIDErr, frameworkHandle, "CFTreeGetTypeID", "")
	registerFunc(&_cFTreeInsertSibling, &_cFTreeInsertSiblingErr, frameworkHandle, "CFTreeInsertSibling", "")
	registerFunc(&_cFTreePrependChild, &_cFTreePrependChildErr, frameworkHandle, "CFTreePrependChild", "")
	registerFunc(&_cFTreeRemove, &_cFTreeRemoveErr, frameworkHandle, "CFTreeRemove", "")
	registerFunc(&_cFTreeRemoveAllChildren, &_cFTreeRemoveAllChildrenErr, frameworkHandle, "CFTreeRemoveAllChildren", "")
	registerFunc(&_cFTreeSetContext, &_cFTreeSetContextErr, frameworkHandle, "CFTreeSetContext", "")
	registerFunc(&_cFTreeSortChildren, &_cFTreeSortChildrenErr, frameworkHandle, "CFTreeSortChildren", "")
	registerFunc(&_cFURLCanBeDecomposed, &_cFURLCanBeDecomposedErr, frameworkHandle, "CFURLCanBeDecomposed", "")
	registerFunc(&_cFURLClearResourcePropertyCache, &_cFURLClearResourcePropertyCacheErr, frameworkHandle, "CFURLClearResourcePropertyCache", "10.6")
	registerFunc(&_cFURLClearResourcePropertyCacheForKey, &_cFURLClearResourcePropertyCacheForKeyErr, frameworkHandle, "CFURLClearResourcePropertyCacheForKey", "10.6")
	registerFunc(&_cFURLCopyAbsoluteURL, &_cFURLCopyAbsoluteURLErr, frameworkHandle, "CFURLCopyAbsoluteURL", "")
	registerFunc(&_cFURLCopyFileSystemPath, &_cFURLCopyFileSystemPathErr, frameworkHandle, "CFURLCopyFileSystemPath", "")
	registerFunc(&_cFURLCopyFragment, &_cFURLCopyFragmentErr, frameworkHandle, "CFURLCopyFragment", "")
	registerFunc(&_cFURLCopyHostName, &_cFURLCopyHostNameErr, frameworkHandle, "CFURLCopyHostName", "")
	registerFunc(&_cFURLCopyLastPathComponent, &_cFURLCopyLastPathComponentErr, frameworkHandle, "CFURLCopyLastPathComponent", "")
	registerFunc(&_cFURLCopyNetLocation, &_cFURLCopyNetLocationErr, frameworkHandle, "CFURLCopyNetLocation", "")
	registerFunc(&_cFURLCopyPassword, &_cFURLCopyPasswordErr, frameworkHandle, "CFURLCopyPassword", "")
	registerFunc(&_cFURLCopyPath, &_cFURLCopyPathErr, frameworkHandle, "CFURLCopyPath", "")
	registerFunc(&_cFURLCopyPathExtension, &_cFURLCopyPathExtensionErr, frameworkHandle, "CFURLCopyPathExtension", "")
	registerFunc(&_cFURLCopyQueryString, &_cFURLCopyQueryStringErr, frameworkHandle, "CFURLCopyQueryString", "")
	registerFunc(&_cFURLCopyResourcePropertiesForKeys, &_cFURLCopyResourcePropertiesForKeysErr, frameworkHandle, "CFURLCopyResourcePropertiesForKeys", "10.6")
	registerFunc(&_cFURLCopyResourcePropertyForKey, &_cFURLCopyResourcePropertyForKeyErr, frameworkHandle, "CFURLCopyResourcePropertyForKey", "10.6")
	registerFunc(&_cFURLCopyResourceSpecifier, &_cFURLCopyResourceSpecifierErr, frameworkHandle, "CFURLCopyResourceSpecifier", "")
	registerFunc(&_cFURLCopyScheme, &_cFURLCopySchemeErr, frameworkHandle, "CFURLCopyScheme", "")
	registerFunc(&_cFURLCopyStrictPath, &_cFURLCopyStrictPathErr, frameworkHandle, "CFURLCopyStrictPath", "")
	registerFunc(&_cFURLCopyUserName, &_cFURLCopyUserNameErr, frameworkHandle, "CFURLCopyUserName", "")
	registerFunc(&_cFURLCreateAbsoluteURLWithBytes, &_cFURLCreateAbsoluteURLWithBytesErr, frameworkHandle, "CFURLCreateAbsoluteURLWithBytes", "")
	registerFunc(&_cFURLCreateBookmarkData, &_cFURLCreateBookmarkDataErr, frameworkHandle, "CFURLCreateBookmarkData", "10.6")
	registerFunc(&_cFURLCreateBookmarkDataFromFile, &_cFURLCreateBookmarkDataFromFileErr, frameworkHandle, "CFURLCreateBookmarkDataFromFile", "10.6")
	registerFunc(&_cFURLCreateByResolvingBookmarkData, &_cFURLCreateByResolvingBookmarkDataErr, frameworkHandle, "CFURLCreateByResolvingBookmarkData", "10.6")
	registerFunc(&_cFURLCreateCopyAppendingPathComponent, &_cFURLCreateCopyAppendingPathComponentErr, frameworkHandle, "CFURLCreateCopyAppendingPathComponent", "")
	registerFunc(&_cFURLCreateCopyAppendingPathExtension, &_cFURLCreateCopyAppendingPathExtensionErr, frameworkHandle, "CFURLCreateCopyAppendingPathExtension", "")
	registerFunc(&_cFURLCreateCopyDeletingLastPathComponent, &_cFURLCreateCopyDeletingLastPathComponentErr, frameworkHandle, "CFURLCreateCopyDeletingLastPathComponent", "")
	registerFunc(&_cFURLCreateCopyDeletingPathExtension, &_cFURLCreateCopyDeletingPathExtensionErr, frameworkHandle, "CFURLCreateCopyDeletingPathExtension", "")
	registerFunc(&_cFURLCreateData, &_cFURLCreateDataErr, frameworkHandle, "CFURLCreateData", "")
	registerFunc(&_cFURLCreateFilePathURL, &_cFURLCreateFilePathURLErr, frameworkHandle, "CFURLCreateFilePathURL", "10.6")
	registerFunc(&_cFURLCreateFileReferenceURL, &_cFURLCreateFileReferenceURLErr, frameworkHandle, "CFURLCreateFileReferenceURL", "10.6")
	registerFunc(&_cFURLCreateFromFileSystemRepresentation, &_cFURLCreateFromFileSystemRepresentationErr, frameworkHandle, "CFURLCreateFromFileSystemRepresentation", "")
	registerFunc(&_cFURLCreateFromFileSystemRepresentationRelativeToBase, &_cFURLCreateFromFileSystemRepresentationRelativeToBaseErr, frameworkHandle, "CFURLCreateFromFileSystemRepresentationRelativeToBase", "")
	registerFunc(&_cFURLCreateResourcePropertiesForKeysFromBookmarkData, &_cFURLCreateResourcePropertiesForKeysFromBookmarkDataErr, frameworkHandle, "CFURLCreateResourcePropertiesForKeysFromBookmarkData", "10.6")
	registerFunc(&_cFURLCreateResourcePropertyForKeyFromBookmarkData, &_cFURLCreateResourcePropertyForKeyFromBookmarkDataErr, frameworkHandle, "CFURLCreateResourcePropertyForKeyFromBookmarkData", "10.6")
	registerFunc(&_cFURLCreateStringByReplacingPercentEscapes, &_cFURLCreateStringByReplacingPercentEscapesErr, frameworkHandle, "CFURLCreateStringByReplacingPercentEscapes", "")
	registerFunc(&_cFURLCreateWithBytes, &_cFURLCreateWithBytesErr, frameworkHandle, "CFURLCreateWithBytes", "")
	registerFunc(&_cFURLCreateWithFileSystemPath, &_cFURLCreateWithFileSystemPathErr, frameworkHandle, "CFURLCreateWithFileSystemPath", "")
	registerFunc(&_cFURLCreateWithFileSystemPathRelativeToBase, &_cFURLCreateWithFileSystemPathRelativeToBaseErr, frameworkHandle, "CFURLCreateWithFileSystemPathRelativeToBase", "")
	registerFunc(&_cFURLCreateWithString, &_cFURLCreateWithStringErr, frameworkHandle, "CFURLCreateWithString", "")
	registerFunc(&_cFURLEnumeratorCreateForDirectoryURL, &_cFURLEnumeratorCreateForDirectoryURLErr, frameworkHandle, "CFURLEnumeratorCreateForDirectoryURL", "10.6")
	registerFunc(&_cFURLEnumeratorCreateForMountedVolumes, &_cFURLEnumeratorCreateForMountedVolumesErr, frameworkHandle, "CFURLEnumeratorCreateForMountedVolumes", "10.6")
	registerFunc(&_cFURLEnumeratorGetDescendentLevel, &_cFURLEnumeratorGetDescendentLevelErr, frameworkHandle, "CFURLEnumeratorGetDescendentLevel", "10.6")
	registerFunc(&_cFURLEnumeratorGetNextURL, &_cFURLEnumeratorGetNextURLErr, frameworkHandle, "CFURLEnumeratorGetNextURL", "10.6")
	registerFunc(&_cFURLEnumeratorGetTypeID, &_cFURLEnumeratorGetTypeIDErr, frameworkHandle, "CFURLEnumeratorGetTypeID", "10.6")
	registerFunc(&_cFURLEnumeratorSkipDescendents, &_cFURLEnumeratorSkipDescendentsErr, frameworkHandle, "CFURLEnumeratorSkipDescendents", "10.6")
	registerFunc(&_cFURLGetBaseURL, &_cFURLGetBaseURLErr, frameworkHandle, "CFURLGetBaseURL", "")
	registerFunc(&_cFURLGetByteRangeForComponent, &_cFURLGetByteRangeForComponentErr, frameworkHandle, "CFURLGetByteRangeForComponent", "")
	registerFunc(&_cFURLGetBytes, &_cFURLGetBytesErr, frameworkHandle, "CFURLGetBytes", "")
	registerFunc(&_cFURLGetFileSystemRepresentation, &_cFURLGetFileSystemRepresentationErr, frameworkHandle, "CFURLGetFileSystemRepresentation", "")
	registerFunc(&_cFURLGetPortNumber, &_cFURLGetPortNumberErr, frameworkHandle, "CFURLGetPortNumber", "")
	registerFunc(&_cFURLGetString, &_cFURLGetStringErr, frameworkHandle, "CFURLGetString", "")
	registerFunc(&_cFURLGetTypeID, &_cFURLGetTypeIDErr, frameworkHandle, "CFURLGetTypeID", "")
	registerFunc(&_cFURLHasDirectoryPath, &_cFURLHasDirectoryPathErr, frameworkHandle, "CFURLHasDirectoryPath", "")
	registerFunc(&_cFURLIsFileReferenceURL, &_cFURLIsFileReferenceURLErr, frameworkHandle, "CFURLIsFileReferenceURL", "10.9")
	registerFunc(&_cFURLResourceIsReachable, &_cFURLResourceIsReachableErr, frameworkHandle, "CFURLResourceIsReachable", "10.6")
	registerFunc(&_cFURLSetResourcePropertiesForKeys, &_cFURLSetResourcePropertiesForKeysErr, frameworkHandle, "CFURLSetResourcePropertiesForKeys", "10.6")
	registerFunc(&_cFURLSetResourcePropertyForKey, &_cFURLSetResourcePropertyForKeyErr, frameworkHandle, "CFURLSetResourcePropertyForKey", "10.6")
	registerFunc(&_cFURLSetTemporaryResourcePropertyForKey, &_cFURLSetTemporaryResourcePropertyForKeyErr, frameworkHandle, "CFURLSetTemporaryResourcePropertyForKey", "10.6")
	registerFunc(&_cFURLStartAccessingSecurityScopedResource, &_cFURLStartAccessingSecurityScopedResourceErr, frameworkHandle, "CFURLStartAccessingSecurityScopedResource", "10.7")
	registerFunc(&_cFURLStopAccessingSecurityScopedResource, &_cFURLStopAccessingSecurityScopedResourceErr, frameworkHandle, "CFURLStopAccessingSecurityScopedResource", "10.7")
	registerFunc(&_cFURLWriteBookmarkDataToFile, &_cFURLWriteBookmarkDataToFileErr, frameworkHandle, "CFURLWriteBookmarkDataToFile", "10.6")
	registerFunc(&_cFUUIDCreate, &_cFUUIDCreateErr, frameworkHandle, "CFUUIDCreate", "")
	registerFunc(&_cFUUIDCreateFromString, &_cFUUIDCreateFromStringErr, frameworkHandle, "CFUUIDCreateFromString", "")
	registerFunc(&_cFUUIDCreateFromUUIDBytes, &_cFUUIDCreateFromUUIDBytesErr, frameworkHandle, "CFUUIDCreateFromUUIDBytes", "")
	registerFunc(&_cFUUIDCreateString, &_cFUUIDCreateStringErr, frameworkHandle, "CFUUIDCreateString", "")
	registerFunc(&_cFUUIDCreateWithBytes, &_cFUUIDCreateWithBytesErr, frameworkHandle, "CFUUIDCreateWithBytes", "")
	registerFunc(&_cFUUIDGetConstantUUIDWithBytes, &_cFUUIDGetConstantUUIDWithBytesErr, frameworkHandle, "CFUUIDGetConstantUUIDWithBytes", "")
	registerFunc(&_cFUUIDGetTypeID, &_cFUUIDGetTypeIDErr, frameworkHandle, "CFUUIDGetTypeID", "")
	registerFunc(&_cFUUIDGetUUIDBytes, &_cFUUIDGetUUIDBytesErr, frameworkHandle, "CFUUIDGetUUIDBytes", "")
	registerFunc(&_cFUserNotificationCancel, &_cFUserNotificationCancelErr, frameworkHandle, "CFUserNotificationCancel", "10.0")
	registerFunc(&_cFUserNotificationCreate, &_cFUserNotificationCreateErr, frameworkHandle, "CFUserNotificationCreate", "10.0")
	registerFunc(&_cFUserNotificationCreateRunLoopSource, &_cFUserNotificationCreateRunLoopSourceErr, frameworkHandle, "CFUserNotificationCreateRunLoopSource", "10.0")
	registerFunc(&_cFUserNotificationDisplayAlert, &_cFUserNotificationDisplayAlertErr, frameworkHandle, "CFUserNotificationDisplayAlert", "10.0")
	registerFunc(&_cFUserNotificationDisplayNotice, &_cFUserNotificationDisplayNoticeErr, frameworkHandle, "CFUserNotificationDisplayNotice", "10.0")
	registerFunc(&_cFUserNotificationGetResponseDictionary, &_cFUserNotificationGetResponseDictionaryErr, frameworkHandle, "CFUserNotificationGetResponseDictionary", "10.0")
	registerFunc(&_cFUserNotificationGetResponseValue, &_cFUserNotificationGetResponseValueErr, frameworkHandle, "CFUserNotificationGetResponseValue", "10.0")
	registerFunc(&_cFUserNotificationGetTypeID, &_cFUserNotificationGetTypeIDErr, frameworkHandle, "CFUserNotificationGetTypeID", "10.0")
	registerFunc(&_cFUserNotificationReceiveResponse, &_cFUserNotificationReceiveResponseErr, frameworkHandle, "CFUserNotificationReceiveResponse", "10.0")
	registerFunc(&_cFUserNotificationUpdate, &_cFUserNotificationUpdateErr, frameworkHandle, "CFUserNotificationUpdate", "10.0")
	registerFunc(&_cFWriteStreamCanAcceptBytes, &_cFWriteStreamCanAcceptBytesErr, frameworkHandle, "CFWriteStreamCanAcceptBytes", "")
	registerFunc(&_cFWriteStreamClose, &_cFWriteStreamCloseErr, frameworkHandle, "CFWriteStreamClose", "")
	registerFunc(&_cFWriteStreamCopyDispatchQueue, &_cFWriteStreamCopyDispatchQueueErr, frameworkHandle, "CFWriteStreamCopyDispatchQueue", "10.9")
	registerFunc(&_cFWriteStreamCopyError, &_cFWriteStreamCopyErrorErr, frameworkHandle, "CFWriteStreamCopyError", "10.5")
	registerFunc(&_cFWriteStreamCopyProperty, &_cFWriteStreamCopyPropertyErr, frameworkHandle, "CFWriteStreamCopyProperty", "")
	registerFunc(&_cFWriteStreamCreateWithAllocatedBuffers, &_cFWriteStreamCreateWithAllocatedBuffersErr, frameworkHandle, "CFWriteStreamCreateWithAllocatedBuffers", "")
	registerFunc(&_cFWriteStreamCreateWithBuffer, &_cFWriteStreamCreateWithBufferErr, frameworkHandle, "CFWriteStreamCreateWithBuffer", "")
	registerFunc(&_cFWriteStreamCreateWithFile, &_cFWriteStreamCreateWithFileErr, frameworkHandle, "CFWriteStreamCreateWithFile", "")
	registerFunc(&_cFWriteStreamGetError, &_cFWriteStreamGetErrorErr, frameworkHandle, "CFWriteStreamGetError", "")
	registerFunc(&_cFWriteStreamGetStatus, &_cFWriteStreamGetStatusErr, frameworkHandle, "CFWriteStreamGetStatus", "")
	registerFunc(&_cFWriteStreamGetTypeID, &_cFWriteStreamGetTypeIDErr, frameworkHandle, "CFWriteStreamGetTypeID", "")
	registerFunc(&_cFWriteStreamOpen, &_cFWriteStreamOpenErr, frameworkHandle, "CFWriteStreamOpen", "")
	registerFunc(&_cFWriteStreamScheduleWithRunLoop, &_cFWriteStreamScheduleWithRunLoopErr, frameworkHandle, "CFWriteStreamScheduleWithRunLoop", "")
	registerFunc(&_cFWriteStreamSetClient, &_cFWriteStreamSetClientErr, frameworkHandle, "CFWriteStreamSetClient", "")
	registerFunc(&_cFWriteStreamSetDispatchQueue, &_cFWriteStreamSetDispatchQueueErr, frameworkHandle, "CFWriteStreamSetDispatchQueue", "10.9")
	registerFunc(&_cFWriteStreamSetProperty, &_cFWriteStreamSetPropertyErr, frameworkHandle, "CFWriteStreamSetProperty", "")
	registerFunc(&_cFWriteStreamUnscheduleFromRunLoop, &_cFWriteStreamUnscheduleFromRunLoopErr, frameworkHandle, "CFWriteStreamUnscheduleFromRunLoop", "")
	registerFunc(&_cFWriteStreamWrite, &_cFWriteStreamWriteErr, frameworkHandle, "CFWriteStreamWrite", "")
	registerFunc(&_cFXMLCreateStringByEscapingEntities, &_cFXMLCreateStringByEscapingEntitiesErr, frameworkHandle, "CFXMLCreateStringByEscapingEntities", "")
	registerFunc(&_cFXMLCreateStringByUnescapingEntities, &_cFXMLCreateStringByUnescapingEntitiesErr, frameworkHandle, "CFXMLCreateStringByUnescapingEntities", "")
	registerFunc(&_cFXMLNodeCreate, &_cFXMLNodeCreateErr, frameworkHandle, "CFXMLNodeCreate", "10.0")
	registerFunc(&_cFXMLNodeCreateCopy, &_cFXMLNodeCreateCopyErr, frameworkHandle, "CFXMLNodeCreateCopy", "10.0")
	registerFunc(&_cFXMLNodeGetInfoPtr, &_cFXMLNodeGetInfoPtrErr, frameworkHandle, "CFXMLNodeGetInfoPtr", "10.0")
	registerFunc(&_cFXMLNodeGetString, &_cFXMLNodeGetStringErr, frameworkHandle, "CFXMLNodeGetString", "10.0")
	registerFunc(&_cFXMLNodeGetTypeCode, &_cFXMLNodeGetTypeCodeErr, frameworkHandle, "CFXMLNodeGetTypeCode", "10.0")
	registerFunc(&_cFXMLNodeGetTypeID, &_cFXMLNodeGetTypeIDErr, frameworkHandle, "CFXMLNodeGetTypeID", "10.0")
	registerFunc(&_cFXMLNodeGetVersion, &_cFXMLNodeGetVersionErr, frameworkHandle, "CFXMLNodeGetVersion", "10.0")
	registerFunc(&_cFXMLParserAbort, &_cFXMLParserAbortErr, frameworkHandle, "CFXMLParserAbort", "10.0")
	registerFunc(&_cFXMLParserCopyErrorDescription, &_cFXMLParserCopyErrorDescriptionErr, frameworkHandle, "CFXMLParserCopyErrorDescription", "10.0")
	registerFunc(&_cFXMLParserCreate, &_cFXMLParserCreateErr, frameworkHandle, "CFXMLParserCreate", "10.0")
	registerFunc(&_cFXMLParserCreateWithDataFromURL, &_cFXMLParserCreateWithDataFromURLErr, frameworkHandle, "CFXMLParserCreateWithDataFromURL", "10.0")
	registerFunc(&_cFXMLParserGetCallBacks, &_cFXMLParserGetCallBacksErr, frameworkHandle, "CFXMLParserGetCallBacks", "10.0")
	registerFunc(&_cFXMLParserGetContext, &_cFXMLParserGetContextErr, frameworkHandle, "CFXMLParserGetContext", "10.0")
	registerFunc(&_cFXMLParserGetDocument, &_cFXMLParserGetDocumentErr, frameworkHandle, "CFXMLParserGetDocument", "10.0")
	registerFunc(&_cFXMLParserGetLineNumber, &_cFXMLParserGetLineNumberErr, frameworkHandle, "CFXMLParserGetLineNumber", "10.0")
	registerFunc(&_cFXMLParserGetLocation, &_cFXMLParserGetLocationErr, frameworkHandle, "CFXMLParserGetLocation", "10.0")
	registerFunc(&_cFXMLParserGetSourceURL, &_cFXMLParserGetSourceURLErr, frameworkHandle, "CFXMLParserGetSourceURL", "10.0")
	registerFunc(&_cFXMLParserGetStatusCode, &_cFXMLParserGetStatusCodeErr, frameworkHandle, "CFXMLParserGetStatusCode", "10.0")
	registerFunc(&_cFXMLParserGetTypeID, &_cFXMLParserGetTypeIDErr, frameworkHandle, "CFXMLParserGetTypeID", "10.0")
	registerFunc(&_cFXMLParserParse, &_cFXMLParserParseErr, frameworkHandle, "CFXMLParserParse", "10.0")
	registerFunc(&_cFXMLTreeCreateFromData, &_cFXMLTreeCreateFromDataErr, frameworkHandle, "CFXMLTreeCreateFromData", "10.0")
	registerFunc(&_cFXMLTreeCreateFromDataWithError, &_cFXMLTreeCreateFromDataWithErrorErr, frameworkHandle, "CFXMLTreeCreateFromDataWithError", "10.0")
	registerFunc(&_cFXMLTreeCreateWithDataFromURL, &_cFXMLTreeCreateWithDataFromURLErr, frameworkHandle, "CFXMLTreeCreateWithDataFromURL", "10.0")
	registerFunc(&_cFXMLTreeCreateWithNode, &_cFXMLTreeCreateWithNodeErr, frameworkHandle, "CFXMLTreeCreateWithNode", "10.0")
	registerFunc(&_cFXMLTreeCreateXMLData, &_cFXMLTreeCreateXMLDataErr, frameworkHandle, "CFXMLTreeCreateXMLData", "10.0")
	registerFunc(&_cFXMLTreeGetNode, &_cFXMLTreeGetNodeErr, frameworkHandle, "CFXMLTreeGetNode", "10.0")
}
