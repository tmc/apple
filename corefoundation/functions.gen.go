// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/dispatch"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreFoundation: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreFoundation: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreFoundation: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _cFAbsoluteTimeGetCurrent func() CFAbsoluteTime

// CFAbsoluteTimeGetCurrent returns the current system absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetCurrent()
func CFAbsoluteTimeGetCurrent() CFAbsoluteTime {
	if _cFAbsoluteTimeGetCurrent == nil {
		panic("CoreFoundation: symbol CFAbsoluteTimeGetCurrent not loaded")
	}
	return _cFAbsoluteTimeGetCurrent()
}


var _cFAllocatorAllocate func(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer

// CFAllocatorAllocate allocates memory using the specified allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocate(_:_:_:)
func CFAllocatorAllocate(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer {
	if _cFAllocatorAllocate == nil {
		panic("CoreFoundation: symbol CFAllocatorAllocate not loaded")
	}
	return _cFAllocatorAllocate(allocator, size, hint)
}


var _cFAllocatorAllocateBytes func(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer

// CFAllocatorAllocateBytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateBytes(_:_:_:)
func CFAllocatorAllocateBytes(allocator CFAllocatorRef, size int, hint uint64) unsafe.Pointer {
	if _cFAllocatorAllocateBytes == nil {
		panic("CoreFoundation: symbol CFAllocatorAllocateBytes not loaded")
	}
	return _cFAllocatorAllocateBytes(allocator, size, hint)
}


var _cFAllocatorAllocateTyped func(allocator CFAllocatorRef, size int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer

// CFAllocatorAllocateTyped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateTyped(_:_:_:_:)
func CFAllocatorAllocateTyped(allocator CFAllocatorRef, size int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer {
	if _cFAllocatorAllocateTyped == nil {
		panic("CoreFoundation: symbol CFAllocatorAllocateTyped not loaded")
	}
	return _cFAllocatorAllocateTyped(allocator, size, descriptor, hint)
}


var _cFAllocatorCreate func(allocator CFAllocatorRef, context *CFAllocatorContext) CFAllocatorRef

// CFAllocatorCreate creates an allocator object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreate(_:_:)
func CFAllocatorCreate(allocator CFAllocatorRef, context *CFAllocatorContext) CFAllocatorRef {
	if _cFAllocatorCreate == nil {
		panic("CoreFoundation: symbol CFAllocatorCreate not loaded")
	}
	return _cFAllocatorCreate(allocator, context)
}



var _cFAllocatorDeallocate func(allocator CFAllocatorRef, ptr unsafe.Pointer) 

// CFAllocatorDeallocate deallocates a block of memory with a given allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocate(_:_:)
func CFAllocatorDeallocate(allocator CFAllocatorRef, ptr unsafe.Pointer)  {
	if _cFAllocatorDeallocate == nil {
		panic("CoreFoundation: symbol CFAllocatorDeallocate not loaded")
	}
	_cFAllocatorDeallocate(allocator, ptr)
}


var _cFAllocatorGetContext func(allocator CFAllocatorRef, context *CFAllocatorContext) 

// CFAllocatorGetContext obtains the context of the specified allocator or of the default allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetContext(_:_:)
func CFAllocatorGetContext(allocator CFAllocatorRef, context *CFAllocatorContext)  {
	if _cFAllocatorGetContext == nil {
		panic("CoreFoundation: symbol CFAllocatorGetContext not loaded")
	}
	_cFAllocatorGetContext(allocator, context)
}


var _cFAllocatorGetDefault func() CFAllocatorRef

// CFAllocatorGetDefault gets the default allocator object for the current thread.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetDefault()
func CFAllocatorGetDefault() CFAllocatorRef {
	if _cFAllocatorGetDefault == nil {
		panic("CoreFoundation: symbol CFAllocatorGetDefault not loaded")
	}
	return _cFAllocatorGetDefault()
}


var _cFAllocatorGetPreferredSizeForSize func(allocator CFAllocatorRef, size int, hint uint64) int

// CFAllocatorGetPreferredSizeForSize obtains the number of bytes likely to be allocated upon a specific request.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetPreferredSizeForSize(_:_:_:)
func CFAllocatorGetPreferredSizeForSize(allocator CFAllocatorRef, size int, hint uint64) int {
	if _cFAllocatorGetPreferredSizeForSize == nil {
		panic("CoreFoundation: symbol CFAllocatorGetPreferredSizeForSize not loaded")
	}
	return _cFAllocatorGetPreferredSizeForSize(allocator, size, hint)
}


var _cFAllocatorGetTypeID func() uint

// CFAllocatorGetTypeID returns the type identifier for the CFAllocator opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetTypeID()
func CFAllocatorGetTypeID() uint {
	if _cFAllocatorGetTypeID == nil {
		panic("CoreFoundation: symbol CFAllocatorGetTypeID not loaded")
	}
	return _cFAllocatorGetTypeID()
}


var _cFAllocatorReallocate func(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer

// CFAllocatorReallocate reallocates memory using the specified allocator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocate(_:_:_:_:)
func CFAllocatorReallocate(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer {
	if _cFAllocatorReallocate == nil {
		panic("CoreFoundation: symbol CFAllocatorReallocate not loaded")
	}
	return _cFAllocatorReallocate(allocator, ptr, newsize, hint)
}


var _cFAllocatorReallocateBytes func(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer

// CFAllocatorReallocateBytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateBytes(_:_:_:_:)
func CFAllocatorReallocateBytes(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, hint uint64) unsafe.Pointer {
	if _cFAllocatorReallocateBytes == nil {
		panic("CoreFoundation: symbol CFAllocatorReallocateBytes not loaded")
	}
	return _cFAllocatorReallocateBytes(allocator, ptr, newsize, hint)
}


var _cFAllocatorReallocateTyped func(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer

// CFAllocatorReallocateTyped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateTyped(_:_:_:_:_:)
func CFAllocatorReallocateTyped(allocator CFAllocatorRef, ptr unsafe.Pointer, newsize int, descriptor CFAllocatorTypeID, hint uint64) unsafe.Pointer {
	if _cFAllocatorReallocateTyped == nil {
		panic("CoreFoundation: symbol CFAllocatorReallocateTyped not loaded")
	}
	return _cFAllocatorReallocateTyped(allocator, ptr, newsize, descriptor, hint)
}


var _cFAllocatorSetDefault func(allocator CFAllocatorRef) 

// CFAllocatorSetDefault sets the given allocator as the default for the current thread.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorSetDefault(_:)
func CFAllocatorSetDefault(allocator CFAllocatorRef)  {
	if _cFAllocatorSetDefault == nil {
		panic("CoreFoundation: symbol CFAllocatorSetDefault not loaded")
	}
	_cFAllocatorSetDefault(allocator)
}


var _cFArrayAppendArray func(theArray CFMutableArrayRef, otherArray CFArrayRef, otherRange CFRange) 

// CFArrayAppendArray adds the values from one array to another array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendArray(_:_:_:)
func CFArrayAppendArray(theArray CFMutableArrayRef, otherArray CFArrayRef, otherRange CFRange)  {
	if _cFArrayAppendArray == nil {
		panic("CoreFoundation: symbol CFArrayAppendArray not loaded")
	}
	_cFArrayAppendArray(theArray, otherArray, otherRange)
}


var _cFArrayAppendValue func(theArray CFMutableArrayRef, value unsafe.Pointer) 

// CFArrayAppendValue adds a value to an array giving it the new largest index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendValue(_:_:)
func CFArrayAppendValue(theArray CFMutableArrayRef, value unsafe.Pointer)  {
	if _cFArrayAppendValue == nil {
		panic("CoreFoundation: symbol CFArrayAppendValue not loaded")
	}
	_cFArrayAppendValue(theArray, value)
}


var _cFArrayApplyFunction func(theArray CFArrayRef, range_ CFRange, applier CFArrayApplierFunction, context unsafe.Pointer) 

// CFArrayApplyFunction calls a function once for each element in range in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplyFunction(_:_:_:_:)
func CFArrayApplyFunction(theArray CFArrayRef, range_ CFRange, applier CFArrayApplierFunction, context unsafe.Pointer)  {
	if _cFArrayApplyFunction == nil {
		panic("CoreFoundation: symbol CFArrayApplyFunction not loaded")
	}
	_cFArrayApplyFunction(theArray, range_, applier, context)
}


var _cFArrayBSearchValues func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer, comparator CFComparatorFunction, context unsafe.Pointer) int

// CFArrayBSearchValues searches an array for a value using a binary search algorithm.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayBSearchValues(_:_:_:_:_:)
func CFArrayBSearchValues(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer, comparator CFComparatorFunction, context unsafe.Pointer) int {
	if _cFArrayBSearchValues == nil {
		panic("CoreFoundation: symbol CFArrayBSearchValues not loaded")
	}
	return _cFArrayBSearchValues(theArray, range_, value, comparator, context)
}


var _cFArrayContainsValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) bool

// CFArrayContainsValue reports whether or not a value is in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayContainsValue(_:_:_:)
func CFArrayContainsValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) bool {
	if _cFArrayContainsValue == nil {
		panic("CoreFoundation: symbol CFArrayContainsValue not loaded")
	}
	return _cFArrayContainsValue(theArray, range_, value)
}


var _cFArrayCreate func(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFArrayCallBacks) CFArrayRef

// CFArrayCreate creates a new immutable array with the given values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreate(_:_:_:_:)
func CFArrayCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFArrayCallBacks) CFArrayRef {
	if _cFArrayCreate == nil {
		panic("CoreFoundation: symbol CFArrayCreate not loaded")
	}
	return _cFArrayCreate(allocator, values, numValues, callBacks)
}


var _cFArrayCreateCopy func(allocator CFAllocatorRef, theArray CFArrayRef) CFArrayRef

// CFArrayCreateCopy creates a new immutable array with the values from another array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateCopy(_:_:)
func CFArrayCreateCopy(allocator CFAllocatorRef, theArray CFArrayRef) CFArrayRef {
	if _cFArrayCreateCopy == nil {
		panic("CoreFoundation: symbol CFArrayCreateCopy not loaded")
	}
	return _cFArrayCreateCopy(allocator, theArray)
}


var _cFArrayCreateMutable func(allocator CFAllocatorRef, capacity int, callBacks *CFArrayCallBacks) CFMutableArrayRef

// CFArrayCreateMutable creates a new empty mutable array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutable(_:_:_:)
func CFArrayCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFArrayCallBacks) CFMutableArrayRef {
	if _cFArrayCreateMutable == nil {
		panic("CoreFoundation: symbol CFArrayCreateMutable not loaded")
	}
	return _cFArrayCreateMutable(allocator, capacity, callBacks)
}


var _cFArrayCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theArray CFArrayRef) CFMutableArrayRef

// CFArrayCreateMutableCopy creates a new mutable array with the values from another array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutableCopy(_:_:_:)
func CFArrayCreateMutableCopy(allocator CFAllocatorRef, capacity int, theArray CFArrayRef) CFMutableArrayRef {
	if _cFArrayCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFArrayCreateMutableCopy not loaded")
	}
	return _cFArrayCreateMutableCopy(allocator, capacity, theArray)
}


var _cFArrayExchangeValuesAtIndices func(theArray CFMutableArrayRef, idx1 int, idx2 int) 

// CFArrayExchangeValuesAtIndices exchanges the values at two indices of an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayExchangeValuesAtIndices(_:_:_:)
func CFArrayExchangeValuesAtIndices(theArray CFMutableArrayRef, idx1 int, idx2 int)  {
	if _cFArrayExchangeValuesAtIndices == nil {
		panic("CoreFoundation: symbol CFArrayExchangeValuesAtIndices not loaded")
	}
	_cFArrayExchangeValuesAtIndices(theArray, idx1, idx2)
}


var _cFArrayGetCount func(theArray CFArrayRef) int

// CFArrayGetCount returns the number of values currently in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCount(_:)
func CFArrayGetCount(theArray CFArrayRef) int {
	if _cFArrayGetCount == nil {
		panic("CoreFoundation: symbol CFArrayGetCount not loaded")
	}
	return _cFArrayGetCount(theArray)
}


var _cFArrayGetCountOfValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int

// CFArrayGetCountOfValue counts the number of times a given value occurs in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCountOfValue(_:_:_:)
func CFArrayGetCountOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int {
	if _cFArrayGetCountOfValue == nil {
		panic("CoreFoundation: symbol CFArrayGetCountOfValue not loaded")
	}
	return _cFArrayGetCountOfValue(theArray, range_, value)
}


var _cFArrayGetFirstIndexOfValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int

// CFArrayGetFirstIndexOfValue searches an array forward for a value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetFirstIndexOfValue(_:_:_:)
func CFArrayGetFirstIndexOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int {
	if _cFArrayGetFirstIndexOfValue == nil {
		panic("CoreFoundation: symbol CFArrayGetFirstIndexOfValue not loaded")
	}
	return _cFArrayGetFirstIndexOfValue(theArray, range_, value)
}


var _cFArrayGetLastIndexOfValue func(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int

// CFArrayGetLastIndexOfValue searches an array backward for a value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetLastIndexOfValue(_:_:_:)
func CFArrayGetLastIndexOfValue(theArray CFArrayRef, range_ CFRange, value unsafe.Pointer) int {
	if _cFArrayGetLastIndexOfValue == nil {
		panic("CoreFoundation: symbol CFArrayGetLastIndexOfValue not loaded")
	}
	return _cFArrayGetLastIndexOfValue(theArray, range_, value)
}


var _cFArrayGetTypeID func() uint

// CFArrayGetTypeID returns the type identifier for the CFArray opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetTypeID()
func CFArrayGetTypeID() uint {
	if _cFArrayGetTypeID == nil {
		panic("CoreFoundation: symbol CFArrayGetTypeID not loaded")
	}
	return _cFArrayGetTypeID()
}


var _cFArrayGetValueAtIndex func(theArray CFArrayRef, idx int) unsafe.Pointer

// CFArrayGetValueAtIndex retrieves a value at a given index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValueAtIndex(_:_:)
func CFArrayGetValueAtIndex(theArray CFArrayRef, idx int) unsafe.Pointer {
	if _cFArrayGetValueAtIndex == nil {
		panic("CoreFoundation: symbol CFArrayGetValueAtIndex not loaded")
	}
	return _cFArrayGetValueAtIndex(theArray, idx)
}


var _cFArrayGetValues func(theArray CFArrayRef, range_ CFRange, values unsafe.Pointer) 

// CFArrayGetValues fills a buffer with values from an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValues(_:_:_:)
func CFArrayGetValues(theArray CFArrayRef, range_ CFRange, values unsafe.Pointer)  {
	if _cFArrayGetValues == nil {
		panic("CoreFoundation: symbol CFArrayGetValues not loaded")
	}
	_cFArrayGetValues(theArray, range_, values)
}


var _cFArrayInsertValueAtIndex func(theArray CFMutableArrayRef, idx int, value unsafe.Pointer) 

// CFArrayInsertValueAtIndex inserts a value into an array at a given index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayInsertValueAtIndex(_:_:_:)
func CFArrayInsertValueAtIndex(theArray CFMutableArrayRef, idx int, value unsafe.Pointer)  {
	if _cFArrayInsertValueAtIndex == nil {
		panic("CoreFoundation: symbol CFArrayInsertValueAtIndex not loaded")
	}
	_cFArrayInsertValueAtIndex(theArray, idx, value)
}


var _cFArrayRemoveAllValues func(theArray CFMutableArrayRef) 

// CFArrayRemoveAllValues removes all the values from an array, making it empty.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveAllValues(_:)
func CFArrayRemoveAllValues(theArray CFMutableArrayRef)  {
	if _cFArrayRemoveAllValues == nil {
		panic("CoreFoundation: symbol CFArrayRemoveAllValues not loaded")
	}
	_cFArrayRemoveAllValues(theArray)
}


var _cFArrayRemoveValueAtIndex func(theArray CFMutableArrayRef, idx int) 

// CFArrayRemoveValueAtIndex removes the value at a given index from an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveValueAtIndex(_:_:)
func CFArrayRemoveValueAtIndex(theArray CFMutableArrayRef, idx int)  {
	if _cFArrayRemoveValueAtIndex == nil {
		panic("CoreFoundation: symbol CFArrayRemoveValueAtIndex not loaded")
	}
	_cFArrayRemoveValueAtIndex(theArray, idx)
}


var _cFArrayReplaceValues func(theArray CFMutableArrayRef, range_ CFRange, newValues unsafe.Pointer, newCount int) 

// CFArrayReplaceValues replaces a range of values in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArrayReplaceValues(_:_:_:_:)
func CFArrayReplaceValues(theArray CFMutableArrayRef, range_ CFRange, newValues unsafe.Pointer, newCount int)  {
	if _cFArrayReplaceValues == nil {
		panic("CoreFoundation: symbol CFArrayReplaceValues not loaded")
	}
	_cFArrayReplaceValues(theArray, range_, newValues, newCount)
}


var _cFArraySetValueAtIndex func(theArray CFMutableArrayRef, idx int, value unsafe.Pointer) 

// CFArraySetValueAtIndex changes the value at a given index in an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArraySetValueAtIndex(_:_:_:)
func CFArraySetValueAtIndex(theArray CFMutableArrayRef, idx int, value unsafe.Pointer)  {
	if _cFArraySetValueAtIndex == nil {
		panic("CoreFoundation: symbol CFArraySetValueAtIndex not loaded")
	}
	_cFArraySetValueAtIndex(theArray, idx, value)
}


var _cFArraySortValues func(theArray CFMutableArrayRef, range_ CFRange, comparator CFComparatorFunction, context unsafe.Pointer) 

// CFArraySortValues sorts the values in an array using a given comparison function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFArraySortValues(_:_:_:_:)
func CFArraySortValues(theArray CFMutableArrayRef, range_ CFRange, comparator CFComparatorFunction, context unsafe.Pointer)  {
	if _cFArraySortValues == nil {
		panic("CoreFoundation: symbol CFArraySortValues not loaded")
	}
	_cFArraySortValues(theArray, range_, comparator, context)
}


var _cFAttributedStringBeginEditing func(aStr CFMutableAttributedStringRef) 

// CFAttributedStringBeginEditing defers internal consistency-checking and coalescing for a mutable attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringBeginEditing(_:)
func CFAttributedStringBeginEditing(aStr CFMutableAttributedStringRef)  {
	if _cFAttributedStringBeginEditing == nil {
		panic("CoreFoundation: symbol CFAttributedStringBeginEditing not loaded")
	}
	_cFAttributedStringBeginEditing(aStr)
}


var _cFAttributedStringCreate func(alloc CFAllocatorRef, str CFStringRef, attributes CFDictionaryRef) CFAttributedStringRef

// CFAttributedStringCreate creates an attributed string with specified string and attributes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreate(_:_:_:)
func CFAttributedStringCreate(alloc CFAllocatorRef, str CFStringRef, attributes CFDictionaryRef) CFAttributedStringRef {
	if _cFAttributedStringCreate == nil {
		panic("CoreFoundation: symbol CFAttributedStringCreate not loaded")
	}
	return _cFAttributedStringCreate(alloc, str, attributes)
}


var _cFAttributedStringCreateCopy func(alloc CFAllocatorRef, aStr CFAttributedStringRef) CFAttributedStringRef

// CFAttributedStringCreateCopy creates an immutable copy of an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateCopy(_:_:)
func CFAttributedStringCreateCopy(alloc CFAllocatorRef, aStr CFAttributedStringRef) CFAttributedStringRef {
	if _cFAttributedStringCreateCopy == nil {
		panic("CoreFoundation: symbol CFAttributedStringCreateCopy not loaded")
	}
	return _cFAttributedStringCreateCopy(alloc, aStr)
}


var _cFAttributedStringCreateMutable func(alloc CFAllocatorRef, maxLength int) CFMutableAttributedStringRef

// CFAttributedStringCreateMutable creates a mutable attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutable(_:_:)
func CFAttributedStringCreateMutable(alloc CFAllocatorRef, maxLength int) CFMutableAttributedStringRef {
	if _cFAttributedStringCreateMutable == nil {
		panic("CoreFoundation: symbol CFAttributedStringCreateMutable not loaded")
	}
	return _cFAttributedStringCreateMutable(alloc, maxLength)
}


var _cFAttributedStringCreateMutableCopy func(alloc CFAllocatorRef, maxLength int, aStr CFAttributedStringRef) CFMutableAttributedStringRef

// CFAttributedStringCreateMutableCopy creates a mutable copy of an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutableCopy(_:_:_:)
func CFAttributedStringCreateMutableCopy(alloc CFAllocatorRef, maxLength int, aStr CFAttributedStringRef) CFMutableAttributedStringRef {
	if _cFAttributedStringCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFAttributedStringCreateMutableCopy not loaded")
	}
	return _cFAttributedStringCreateMutableCopy(alloc, maxLength, aStr)
}


var _cFAttributedStringCreateWithSubstring func(alloc CFAllocatorRef, aStr CFAttributedStringRef, range_ CFRange) CFAttributedStringRef

// CFAttributedStringCreateWithSubstring creates a sub-attributed string from the specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateWithSubstring(_:_:_:)
func CFAttributedStringCreateWithSubstring(alloc CFAllocatorRef, aStr CFAttributedStringRef, range_ CFRange) CFAttributedStringRef {
	if _cFAttributedStringCreateWithSubstring == nil {
		panic("CoreFoundation: symbol CFAttributedStringCreateWithSubstring not loaded")
	}
	return _cFAttributedStringCreateWithSubstring(alloc, aStr, range_)
}


var _cFAttributedStringEndEditing func(aStr CFMutableAttributedStringRef) 

// CFAttributedStringEndEditing re-enables internal consistency-checking and coalescing for a mutable attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringEndEditing(_:)
func CFAttributedStringEndEditing(aStr CFMutableAttributedStringRef)  {
	if _cFAttributedStringEndEditing == nil {
		panic("CoreFoundation: symbol CFAttributedStringEndEditing not loaded")
	}
	_cFAttributedStringEndEditing(aStr)
}


var _cFAttributedStringGetAttribute func(aStr CFAttributedStringRef, loc int, attrName CFStringRef, effectiveRange *CFRange) CFTypeRef

// CFAttributedStringGetAttribute returns the value of a given attribute of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttribute(_:_:_:_:)
func CFAttributedStringGetAttribute(aStr CFAttributedStringRef, loc int, attrName CFStringRef, effectiveRange *CFRange) CFTypeRef {
	if _cFAttributedStringGetAttribute == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetAttribute not loaded")
	}
	return _cFAttributedStringGetAttribute(aStr, loc, attrName, effectiveRange)
}


var _cFAttributedStringGetAttributeAndLongestEffectiveRange func(aStr CFAttributedStringRef, loc int, attrName CFStringRef, inRange CFRange, longestEffectiveRange *CFRange) CFTypeRef

// CFAttributedStringGetAttributeAndLongestEffectiveRange returns the value of a given attribute of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributeAndLongestEffectiveRange(_:_:_:_:_:)
func CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr CFAttributedStringRef, loc int, attrName CFStringRef, inRange CFRange, longestEffectiveRange *CFRange) CFTypeRef {
	if _cFAttributedStringGetAttributeAndLongestEffectiveRange == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetAttributeAndLongestEffectiveRange not loaded")
	}
	return _cFAttributedStringGetAttributeAndLongestEffectiveRange(aStr, loc, attrName, inRange, longestEffectiveRange)
}


var _cFAttributedStringGetAttributes func(aStr CFAttributedStringRef, loc int, effectiveRange *CFRange) CFDictionaryRef

// CFAttributedStringGetAttributes returns the attributes of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributes(_:_:_:)
func CFAttributedStringGetAttributes(aStr CFAttributedStringRef, loc int, effectiveRange *CFRange) CFDictionaryRef {
	if _cFAttributedStringGetAttributes == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetAttributes not loaded")
	}
	return _cFAttributedStringGetAttributes(aStr, loc, effectiveRange)
}


var _cFAttributedStringGetAttributesAndLongestEffectiveRange func(aStr CFAttributedStringRef, loc int, inRange CFRange, longestEffectiveRange *CFRange) CFDictionaryRef

// CFAttributedStringGetAttributesAndLongestEffectiveRange returns the attributes of an attributed string at a specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributesAndLongestEffectiveRange(_:_:_:_:)
func CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr CFAttributedStringRef, loc int, inRange CFRange, longestEffectiveRange *CFRange) CFDictionaryRef {
	if _cFAttributedStringGetAttributesAndLongestEffectiveRange == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetAttributesAndLongestEffectiveRange not loaded")
	}
	return _cFAttributedStringGetAttributesAndLongestEffectiveRange(aStr, loc, inRange, longestEffectiveRange)
}


var _cFAttributedStringGetBidiLevelsAndResolvedDirections func(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool

// CFAttributedStringGetBidiLevelsAndResolvedDirections.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetBidiLevelsAndResolvedDirections(_:_:_:_:_:)
func CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool {
	if _cFAttributedStringGetBidiLevelsAndResolvedDirections == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetBidiLevelsAndResolvedDirections not loaded")
	}
	return _cFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
}


var _cFAttributedStringGetLength func(aStr CFAttributedStringRef) int

// CFAttributedStringGetLength returns the length of the attributed string in characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetLength(_:)
func CFAttributedStringGetLength(aStr CFAttributedStringRef) int {
	if _cFAttributedStringGetLength == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetLength not loaded")
	}
	return _cFAttributedStringGetLength(aStr)
}


var _cFAttributedStringGetMutableString func(aStr CFMutableAttributedStringRef) CFMutableStringRef

// CFAttributedStringGetMutableString gets as a mutable string the string for an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetMutableString(_:)
func CFAttributedStringGetMutableString(aStr CFMutableAttributedStringRef) CFMutableStringRef {
	if _cFAttributedStringGetMutableString == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetMutableString not loaded")
	}
	return _cFAttributedStringGetMutableString(aStr)
}


var _cFAttributedStringGetStatisticalWritingDirections func(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool

// CFAttributedStringGetStatisticalWritingDirections.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetStatisticalWritingDirections(_:_:_:_:_:)
func CFAttributedStringGetStatisticalWritingDirections(attributedString CFAttributedStringRef, range_ CFRange, baseDirection int8, bidiLevels *uint8, baseDirections *uint8) bool {
	if _cFAttributedStringGetStatisticalWritingDirections == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetStatisticalWritingDirections not loaded")
	}
	return _cFAttributedStringGetStatisticalWritingDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
}


var _cFAttributedStringGetString func(aStr CFAttributedStringRef) CFStringRef

// CFAttributedStringGetString returns the string for an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetString(_:)
func CFAttributedStringGetString(aStr CFAttributedStringRef) CFStringRef {
	if _cFAttributedStringGetString == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetString not loaded")
	}
	return _cFAttributedStringGetString(aStr)
}


var _cFAttributedStringGetTypeID func() uint

// CFAttributedStringGetTypeID returns the type identifier for the CFAttributedString opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetTypeID()
func CFAttributedStringGetTypeID() uint {
	if _cFAttributedStringGetTypeID == nil {
		panic("CoreFoundation: symbol CFAttributedStringGetTypeID not loaded")
	}
	return _cFAttributedStringGetTypeID()
}


var _cFAttributedStringRemoveAttribute func(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef) 

// CFAttributedStringRemoveAttribute removes the value of a single attribute over a specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringRemoveAttribute(_:_:_:)
func CFAttributedStringRemoveAttribute(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef)  {
	if _cFAttributedStringRemoveAttribute == nil {
		panic("CoreFoundation: symbol CFAttributedStringRemoveAttribute not loaded")
	}
	_cFAttributedStringRemoveAttribute(aStr, range_, attrName)
}


var _cFAttributedStringReplaceAttributedString func(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFAttributedStringRef) 

// CFAttributedStringReplaceAttributedString replaces the attributed substring over a range with another attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceAttributedString(_:_:_:)
func CFAttributedStringReplaceAttributedString(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFAttributedStringRef)  {
	if _cFAttributedStringReplaceAttributedString == nil {
		panic("CoreFoundation: symbol CFAttributedStringReplaceAttributedString not loaded")
	}
	_cFAttributedStringReplaceAttributedString(aStr, range_, replacement)
}


var _cFAttributedStringReplaceString func(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFStringRef) 

// CFAttributedStringReplaceString modifies the string of an attributed string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceString(_:_:_:)
func CFAttributedStringReplaceString(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFStringRef)  {
	if _cFAttributedStringReplaceString == nil {
		panic("CoreFoundation: symbol CFAttributedStringReplaceString not loaded")
	}
	_cFAttributedStringReplaceString(aStr, range_, replacement)
}


var _cFAttributedStringSetAttribute func(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef, value CFTypeRef) 

// CFAttributedStringSetAttribute sets the value of a single attribute over the specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttribute(_:_:_:_:)
func CFAttributedStringSetAttribute(aStr CFMutableAttributedStringRef, range_ CFRange, attrName CFStringRef, value CFTypeRef)  {
	if _cFAttributedStringSetAttribute == nil {
		panic("CoreFoundation: symbol CFAttributedStringSetAttribute not loaded")
	}
	_cFAttributedStringSetAttribute(aStr, range_, attrName, value)
}


var _cFAttributedStringSetAttributes func(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFDictionaryRef, clearOtherAttributes bool) 

// CFAttributedStringSetAttributes sets the value of attributes of a mutable attributed string over a specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttributes(_:_:_:_:)
func CFAttributedStringSetAttributes(aStr CFMutableAttributedStringRef, range_ CFRange, replacement CFDictionaryRef, clearOtherAttributes bool)  {
	if _cFAttributedStringSetAttributes == nil {
		panic("CoreFoundation: symbol CFAttributedStringSetAttributes not loaded")
	}
	_cFAttributedStringSetAttributes(aStr, range_, replacement, clearOtherAttributes)
}


var _cFAutorelease func(arg CFTypeRef) CFTypeRef

// CFAutorelease.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFAutorelease
func CFAutorelease(arg CFTypeRef) CFTypeRef {
	if _cFAutorelease == nil {
		panic("CoreFoundation: symbol CFAutorelease not loaded")
	}
	return _cFAutorelease(arg)
}


var _cFBagAddValue func(theBag CFMutableBagRef, value unsafe.Pointer) 

// CFBagAddValue adds a value to a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagAddValue(_:_:)
func CFBagAddValue(theBag CFMutableBagRef, value unsafe.Pointer)  {
	if _cFBagAddValue == nil {
		panic("CoreFoundation: symbol CFBagAddValue not loaded")
	}
	_cFBagAddValue(theBag, value)
}


var _cFBagApplyFunction func(theBag CFBagRef, applier CFBagApplierFunction, context unsafe.Pointer) 

// CFBagApplyFunction calls a function once for each value in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagApplyFunction(_:_:_:)
func CFBagApplyFunction(theBag CFBagRef, applier CFBagApplierFunction, context unsafe.Pointer)  {
	if _cFBagApplyFunction == nil {
		panic("CoreFoundation: symbol CFBagApplyFunction not loaded")
	}
	_cFBagApplyFunction(theBag, applier, context)
}


var _cFBagContainsValue func(theBag CFBagRef, value unsafe.Pointer) bool

// CFBagContainsValue reports whether or not a value is in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagContainsValue(_:_:)
func CFBagContainsValue(theBag CFBagRef, value unsafe.Pointer) bool {
	if _cFBagContainsValue == nil {
		panic("CoreFoundation: symbol CFBagContainsValue not loaded")
	}
	return _cFBagContainsValue(theBag, value)
}


var _cFBagCreate func(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFBagCallBacks) CFBagRef

// CFBagCreate creates an immutable bag containing specified values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreate(_:_:_:_:)
func CFBagCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFBagCallBacks) CFBagRef {
	if _cFBagCreate == nil {
		panic("CoreFoundation: symbol CFBagCreate not loaded")
	}
	return _cFBagCreate(allocator, values, numValues, callBacks)
}


var _cFBagCreateCopy func(allocator CFAllocatorRef, theBag CFBagRef) CFBagRef

// CFBagCreateCopy creates an immutable bag with the values of another bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateCopy(_:_:)
func CFBagCreateCopy(allocator CFAllocatorRef, theBag CFBagRef) CFBagRef {
	if _cFBagCreateCopy == nil {
		panic("CoreFoundation: symbol CFBagCreateCopy not loaded")
	}
	return _cFBagCreateCopy(allocator, theBag)
}


var _cFBagCreateMutable func(allocator CFAllocatorRef, capacity int, callBacks *CFBagCallBacks) CFMutableBagRef

// CFBagCreateMutable creates a new empty mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutable(_:_:_:)
func CFBagCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFBagCallBacks) CFMutableBagRef {
	if _cFBagCreateMutable == nil {
		panic("CoreFoundation: symbol CFBagCreateMutable not loaded")
	}
	return _cFBagCreateMutable(allocator, capacity, callBacks)
}


var _cFBagCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theBag CFBagRef) CFMutableBagRef

// CFBagCreateMutableCopy creates a new mutable bag with the values from another bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutableCopy(_:_:_:)
func CFBagCreateMutableCopy(allocator CFAllocatorRef, capacity int, theBag CFBagRef) CFMutableBagRef {
	if _cFBagCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFBagCreateMutableCopy not loaded")
	}
	return _cFBagCreateMutableCopy(allocator, capacity, theBag)
}


var _cFBagGetCount func(theBag CFBagRef) int

// CFBagGetCount returns the number of values currently in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCount(_:)
func CFBagGetCount(theBag CFBagRef) int {
	if _cFBagGetCount == nil {
		panic("CoreFoundation: symbol CFBagGetCount not loaded")
	}
	return _cFBagGetCount(theBag)
}


var _cFBagGetCountOfValue func(theBag CFBagRef, value unsafe.Pointer) int

// CFBagGetCountOfValue returns the number of times a value occurs in a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCountOfValue(_:_:)
func CFBagGetCountOfValue(theBag CFBagRef, value unsafe.Pointer) int {
	if _cFBagGetCountOfValue == nil {
		panic("CoreFoundation: symbol CFBagGetCountOfValue not loaded")
	}
	return _cFBagGetCountOfValue(theBag, value)
}


var _cFBagGetTypeID func() uint

// CFBagGetTypeID returns the type identifier for the CFBag opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetTypeID()
func CFBagGetTypeID() uint {
	if _cFBagGetTypeID == nil {
		panic("CoreFoundation: symbol CFBagGetTypeID not loaded")
	}
	return _cFBagGetTypeID()
}


var _cFBagGetValue func(theBag CFBagRef, value unsafe.Pointer) unsafe.Pointer

// CFBagGetValue returns a requested value from a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValue(_:_:)
func CFBagGetValue(theBag CFBagRef, value unsafe.Pointer) unsafe.Pointer {
	if _cFBagGetValue == nil {
		panic("CoreFoundation: symbol CFBagGetValue not loaded")
	}
	return _cFBagGetValue(theBag, value)
}


var _cFBagGetValueIfPresent func(theBag CFBagRef, candidate unsafe.Pointer, value unsafe.Pointer) bool

// CFBagGetValueIfPresent reports whether or not a value is in a bag, and returns that value indirectly if it exists.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValueIfPresent(_:_:_:)
func CFBagGetValueIfPresent(theBag CFBagRef, candidate unsafe.Pointer, value unsafe.Pointer) bool {
	if _cFBagGetValueIfPresent == nil {
		panic("CoreFoundation: symbol CFBagGetValueIfPresent not loaded")
	}
	return _cFBagGetValueIfPresent(theBag, candidate, value)
}


var _cFBagGetValues func(theBag CFBagRef, values unsafe.Pointer) 

// CFBagGetValues fills a buffer with values from a bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValues(_:_:)
func CFBagGetValues(theBag CFBagRef, values unsafe.Pointer)  {
	if _cFBagGetValues == nil {
		panic("CoreFoundation: symbol CFBagGetValues not loaded")
	}
	_cFBagGetValues(theBag, values)
}


var _cFBagRemoveAllValues func(theBag CFMutableBagRef) 

// CFBagRemoveAllValues removes all values from a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveAllValues(_:)
func CFBagRemoveAllValues(theBag CFMutableBagRef)  {
	if _cFBagRemoveAllValues == nil {
		panic("CoreFoundation: symbol CFBagRemoveAllValues not loaded")
	}
	_cFBagRemoveAllValues(theBag)
}


var _cFBagRemoveValue func(theBag CFMutableBagRef, value unsafe.Pointer) 

// CFBagRemoveValue removes a value from a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveValue(_:_:)
func CFBagRemoveValue(theBag CFMutableBagRef, value unsafe.Pointer)  {
	if _cFBagRemoveValue == nil {
		panic("CoreFoundation: symbol CFBagRemoveValue not loaded")
	}
	_cFBagRemoveValue(theBag, value)
}


var _cFBagReplaceValue func(theBag CFMutableBagRef, value unsafe.Pointer) 

// CFBagReplaceValue replaces a value in a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagReplaceValue(_:_:)
func CFBagReplaceValue(theBag CFMutableBagRef, value unsafe.Pointer)  {
	if _cFBagReplaceValue == nil {
		panic("CoreFoundation: symbol CFBagReplaceValue not loaded")
	}
	_cFBagReplaceValue(theBag, value)
}


var _cFBagSetValue func(theBag CFMutableBagRef, value unsafe.Pointer) 

// CFBagSetValue sets a value in a mutable bag.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBagSetValue(_:_:)
func CFBagSetValue(theBag CFMutableBagRef, value unsafe.Pointer)  {
	if _cFBagSetValue == nil {
		panic("CoreFoundation: symbol CFBagSetValue not loaded")
	}
	_cFBagSetValue(theBag, value)
}


var _cFBinaryHeapAddValue func(heap CFBinaryHeapRef, value unsafe.Pointer) 

// CFBinaryHeapAddValue adds a value to a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapAddValue(_:_:)
func CFBinaryHeapAddValue(heap CFBinaryHeapRef, value unsafe.Pointer)  {
	if _cFBinaryHeapAddValue == nil {
		panic("CoreFoundation: symbol CFBinaryHeapAddValue not loaded")
	}
	_cFBinaryHeapAddValue(heap, value)
}


var _cFBinaryHeapApplyFunction func(heap CFBinaryHeapRef, applier CFBinaryHeapApplierFunction, context unsafe.Pointer) 

// CFBinaryHeapApplyFunction iteratively applies a function to all the values in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplyFunction(_:_:_:)
func CFBinaryHeapApplyFunction(heap CFBinaryHeapRef, applier CFBinaryHeapApplierFunction, context unsafe.Pointer)  {
	if _cFBinaryHeapApplyFunction == nil {
		panic("CoreFoundation: symbol CFBinaryHeapApplyFunction not loaded")
	}
	_cFBinaryHeapApplyFunction(heap, applier, context)
}


var _cFBinaryHeapContainsValue func(heap CFBinaryHeapRef, value unsafe.Pointer) bool

// CFBinaryHeapContainsValue returns whether a given value is in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapContainsValue(_:_:)
func CFBinaryHeapContainsValue(heap CFBinaryHeapRef, value unsafe.Pointer) bool {
	if _cFBinaryHeapContainsValue == nil {
		panic("CoreFoundation: symbol CFBinaryHeapContainsValue not loaded")
	}
	return _cFBinaryHeapContainsValue(heap, value)
}


var _cFBinaryHeapCreate func(allocator CFAllocatorRef, capacity int, callBacks *CFBinaryHeapCallBacks, compareContext *CFBinaryHeapCompareContext) CFBinaryHeapRef

// CFBinaryHeapCreate creates a new mutable or fixed-mutable binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreate(_:_:_:_:)
func CFBinaryHeapCreate(allocator CFAllocatorRef, capacity int, callBacks *CFBinaryHeapCallBacks, compareContext *CFBinaryHeapCompareContext) CFBinaryHeapRef {
	if _cFBinaryHeapCreate == nil {
		panic("CoreFoundation: symbol CFBinaryHeapCreate not loaded")
	}
	return _cFBinaryHeapCreate(allocator, capacity, callBacks, compareContext)
}


var _cFBinaryHeapCreateCopy func(allocator CFAllocatorRef, capacity int, heap CFBinaryHeapRef) CFBinaryHeapRef

// CFBinaryHeapCreateCopy creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreateCopy(_:_:_:)
func CFBinaryHeapCreateCopy(allocator CFAllocatorRef, capacity int, heap CFBinaryHeapRef) CFBinaryHeapRef {
	if _cFBinaryHeapCreateCopy == nil {
		panic("CoreFoundation: symbol CFBinaryHeapCreateCopy not loaded")
	}
	return _cFBinaryHeapCreateCopy(allocator, capacity, heap)
}


var _cFBinaryHeapGetCount func(heap CFBinaryHeapRef) int

// CFBinaryHeapGetCount returns the number of values currently in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCount(_:)
func CFBinaryHeapGetCount(heap CFBinaryHeapRef) int {
	if _cFBinaryHeapGetCount == nil {
		panic("CoreFoundation: symbol CFBinaryHeapGetCount not loaded")
	}
	return _cFBinaryHeapGetCount(heap)
}


var _cFBinaryHeapGetCountOfValue func(heap CFBinaryHeapRef, value unsafe.Pointer) int

// CFBinaryHeapGetCountOfValue counts the number of times a given value occurs in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCountOfValue(_:_:)
func CFBinaryHeapGetCountOfValue(heap CFBinaryHeapRef, value unsafe.Pointer) int {
	if _cFBinaryHeapGetCountOfValue == nil {
		panic("CoreFoundation: symbol CFBinaryHeapGetCountOfValue not loaded")
	}
	return _cFBinaryHeapGetCountOfValue(heap, value)
}


var _cFBinaryHeapGetMinimum func(heap CFBinaryHeapRef) unsafe.Pointer

// CFBinaryHeapGetMinimum returns the minimum value in a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimum(_:)
func CFBinaryHeapGetMinimum(heap CFBinaryHeapRef) unsafe.Pointer {
	if _cFBinaryHeapGetMinimum == nil {
		panic("CoreFoundation: symbol CFBinaryHeapGetMinimum not loaded")
	}
	return _cFBinaryHeapGetMinimum(heap)
}


var _cFBinaryHeapGetMinimumIfPresent func(heap CFBinaryHeapRef, value unsafe.Pointer) bool

// CFBinaryHeapGetMinimumIfPresent returns the minimum value in a binary heap, if present.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimumIfPresent(_:_:)
func CFBinaryHeapGetMinimumIfPresent(heap CFBinaryHeapRef, value unsafe.Pointer) bool {
	if _cFBinaryHeapGetMinimumIfPresent == nil {
		panic("CoreFoundation: symbol CFBinaryHeapGetMinimumIfPresent not loaded")
	}
	return _cFBinaryHeapGetMinimumIfPresent(heap, value)
}


var _cFBinaryHeapGetTypeID func() uint

// CFBinaryHeapGetTypeID returns the type identifier of the [CFBinaryHeap] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetTypeID()
func CFBinaryHeapGetTypeID() uint {
	if _cFBinaryHeapGetTypeID == nil {
		panic("CoreFoundation: symbol CFBinaryHeapGetTypeID not loaded")
	}
	return _cFBinaryHeapGetTypeID()
}


var _cFBinaryHeapGetValues func(heap CFBinaryHeapRef, values unsafe.Pointer) 

// CFBinaryHeapGetValues copies all the values from a binary heap into a sorted C array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetValues(_:_:)
func CFBinaryHeapGetValues(heap CFBinaryHeapRef, values unsafe.Pointer)  {
	if _cFBinaryHeapGetValues == nil {
		panic("CoreFoundation: symbol CFBinaryHeapGetValues not loaded")
	}
	_cFBinaryHeapGetValues(heap, values)
}


var _cFBinaryHeapRemoveAllValues func(heap CFBinaryHeapRef) 

// CFBinaryHeapRemoveAllValues removes all values from a binary heap, making it empty.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveAllValues(_:)
func CFBinaryHeapRemoveAllValues(heap CFBinaryHeapRef)  {
	if _cFBinaryHeapRemoveAllValues == nil {
		panic("CoreFoundation: symbol CFBinaryHeapRemoveAllValues not loaded")
	}
	_cFBinaryHeapRemoveAllValues(heap)
}


var _cFBinaryHeapRemoveMinimumValue func(heap CFBinaryHeapRef) 

// CFBinaryHeapRemoveMinimumValue removes the minimum value from a binary heap.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveMinimumValue(_:)
func CFBinaryHeapRemoveMinimumValue(heap CFBinaryHeapRef)  {
	if _cFBinaryHeapRemoveMinimumValue == nil {
		panic("CoreFoundation: symbol CFBinaryHeapRemoveMinimumValue not loaded")
	}
	_cFBinaryHeapRemoveMinimumValue(heap)
}


var _cFBitVectorContainsBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) bool

// CFBitVectorContainsBit returns whether a bit vector contains a particular bit value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorContainsBit(_:_:_:)
func CFBitVectorContainsBit(bv CFBitVectorRef, range_ CFRange, value CFBit) bool {
	if _cFBitVectorContainsBit == nil {
		panic("CoreFoundation: symbol CFBitVectorContainsBit not loaded")
	}
	return _cFBitVectorContainsBit(bv, range_, value)
}


var _cFBitVectorCreate func(allocator CFAllocatorRef, bytes *uint8, numBits int) CFBitVectorRef

// CFBitVectorCreate creates an immutable bit vector from a block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreate(_:_:_:)
func CFBitVectorCreate(allocator CFAllocatorRef, bytes *uint8, numBits int) CFBitVectorRef {
	if _cFBitVectorCreate == nil {
		panic("CoreFoundation: symbol CFBitVectorCreate not loaded")
	}
	return _cFBitVectorCreate(allocator, bytes, numBits)
}


var _cFBitVectorCreateCopy func(allocator CFAllocatorRef, bv CFBitVectorRef) CFBitVectorRef

// CFBitVectorCreateCopy creates an immutable bit vector that is a copy of another bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateCopy(_:_:)
func CFBitVectorCreateCopy(allocator CFAllocatorRef, bv CFBitVectorRef) CFBitVectorRef {
	if _cFBitVectorCreateCopy == nil {
		panic("CoreFoundation: symbol CFBitVectorCreateCopy not loaded")
	}
	return _cFBitVectorCreateCopy(allocator, bv)
}


var _cFBitVectorCreateMutable func(allocator CFAllocatorRef, capacity int) CFMutableBitVectorRef

// CFBitVectorCreateMutable creates a mutable bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutable(_:_:)
func CFBitVectorCreateMutable(allocator CFAllocatorRef, capacity int) CFMutableBitVectorRef {
	if _cFBitVectorCreateMutable == nil {
		panic("CoreFoundation: symbol CFBitVectorCreateMutable not loaded")
	}
	return _cFBitVectorCreateMutable(allocator, capacity)
}


var _cFBitVectorCreateMutableCopy func(allocator CFAllocatorRef, capacity int, bv CFBitVectorRef) CFMutableBitVectorRef

// CFBitVectorCreateMutableCopy creates a new mutable bit vector from a pre-existing bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutableCopy(_:_:_:)
func CFBitVectorCreateMutableCopy(allocator CFAllocatorRef, capacity int, bv CFBitVectorRef) CFMutableBitVectorRef {
	if _cFBitVectorCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFBitVectorCreateMutableCopy not loaded")
	}
	return _cFBitVectorCreateMutableCopy(allocator, capacity, bv)
}


var _cFBitVectorFlipBitAtIndex func(bv CFMutableBitVectorRef, idx int) 

// CFBitVectorFlipBitAtIndex flips a bit value in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBitAtIndex(_:_:)
func CFBitVectorFlipBitAtIndex(bv CFMutableBitVectorRef, idx int)  {
	if _cFBitVectorFlipBitAtIndex == nil {
		panic("CoreFoundation: symbol CFBitVectorFlipBitAtIndex not loaded")
	}
	_cFBitVectorFlipBitAtIndex(bv, idx)
}


var _cFBitVectorFlipBits func(bv CFMutableBitVectorRef, range_ CFRange) 

// CFBitVectorFlipBits flips a range of bit values in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBits(_:_:)
func CFBitVectorFlipBits(bv CFMutableBitVectorRef, range_ CFRange)  {
	if _cFBitVectorFlipBits == nil {
		panic("CoreFoundation: symbol CFBitVectorFlipBits not loaded")
	}
	_cFBitVectorFlipBits(bv, range_)
}


var _cFBitVectorGetBitAtIndex func(bv CFBitVectorRef, idx int) CFBit

// CFBitVectorGetBitAtIndex returns the bit value at a given index in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBitAtIndex(_:_:)
func CFBitVectorGetBitAtIndex(bv CFBitVectorRef, idx int) CFBit {
	if _cFBitVectorGetBitAtIndex == nil {
		panic("CoreFoundation: symbol CFBitVectorGetBitAtIndex not loaded")
	}
	return _cFBitVectorGetBitAtIndex(bv, idx)
}


var _cFBitVectorGetBits func(bv CFBitVectorRef, range_ CFRange, bytes *uint8) 

// CFBitVectorGetBits returns the bit values in a range of indices in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBits(_:_:_:)
func CFBitVectorGetBits(bv CFBitVectorRef, range_ CFRange, bytes *uint8)  {
	if _cFBitVectorGetBits == nil {
		panic("CoreFoundation: symbol CFBitVectorGetBits not loaded")
	}
	_cFBitVectorGetBits(bv, range_, bytes)
}


var _cFBitVectorGetCount func(bv CFBitVectorRef) int

// CFBitVectorGetCount returns the number of bit values in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCount(_:)
func CFBitVectorGetCount(bv CFBitVectorRef) int {
	if _cFBitVectorGetCount == nil {
		panic("CoreFoundation: symbol CFBitVectorGetCount not loaded")
	}
	return _cFBitVectorGetCount(bv)
}


var _cFBitVectorGetCountOfBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) int

// CFBitVectorGetCountOfBit counts the number of times a certain bit value occurs within a range of bits in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCountOfBit(_:_:_:)
func CFBitVectorGetCountOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) int {
	if _cFBitVectorGetCountOfBit == nil {
		panic("CoreFoundation: symbol CFBitVectorGetCountOfBit not loaded")
	}
	return _cFBitVectorGetCountOfBit(bv, range_, value)
}


var _cFBitVectorGetFirstIndexOfBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) int

// CFBitVectorGetFirstIndexOfBit locates the first occurrence of a certain bit value within a range of bits in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetFirstIndexOfBit(_:_:_:)
func CFBitVectorGetFirstIndexOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) int {
	if _cFBitVectorGetFirstIndexOfBit == nil {
		panic("CoreFoundation: symbol CFBitVectorGetFirstIndexOfBit not loaded")
	}
	return _cFBitVectorGetFirstIndexOfBit(bv, range_, value)
}


var _cFBitVectorGetLastIndexOfBit func(bv CFBitVectorRef, range_ CFRange, value CFBit) int

// CFBitVectorGetLastIndexOfBit locates the last occurrence of a certain bit value within a range of bits in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetLastIndexOfBit(_:_:_:)
func CFBitVectorGetLastIndexOfBit(bv CFBitVectorRef, range_ CFRange, value CFBit) int {
	if _cFBitVectorGetLastIndexOfBit == nil {
		panic("CoreFoundation: symbol CFBitVectorGetLastIndexOfBit not loaded")
	}
	return _cFBitVectorGetLastIndexOfBit(bv, range_, value)
}


var _cFBitVectorGetTypeID func() uint

// CFBitVectorGetTypeID returns the type identifier for the CFBitVector opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetTypeID()
func CFBitVectorGetTypeID() uint {
	if _cFBitVectorGetTypeID == nil {
		panic("CoreFoundation: symbol CFBitVectorGetTypeID not loaded")
	}
	return _cFBitVectorGetTypeID()
}


var _cFBitVectorSetAllBits func(bv CFMutableBitVectorRef, value CFBit) 

// CFBitVectorSetAllBits sets all bits in a bit vector to a particular value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetAllBits(_:_:)
func CFBitVectorSetAllBits(bv CFMutableBitVectorRef, value CFBit)  {
	if _cFBitVectorSetAllBits == nil {
		panic("CoreFoundation: symbol CFBitVectorSetAllBits not loaded")
	}
	_cFBitVectorSetAllBits(bv, value)
}


var _cFBitVectorSetBitAtIndex func(bv CFMutableBitVectorRef, idx int, value CFBit) 

// CFBitVectorSetBitAtIndex sets the value of a particular bit in a bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBitAtIndex(_:_:_:)
func CFBitVectorSetBitAtIndex(bv CFMutableBitVectorRef, idx int, value CFBit)  {
	if _cFBitVectorSetBitAtIndex == nil {
		panic("CoreFoundation: symbol CFBitVectorSetBitAtIndex not loaded")
	}
	_cFBitVectorSetBitAtIndex(bv, idx, value)
}


var _cFBitVectorSetBits func(bv CFMutableBitVectorRef, range_ CFRange, value CFBit) 

// CFBitVectorSetBits sets a range of bits in a bit vector to a particular value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBits(_:_:_:)
func CFBitVectorSetBits(bv CFMutableBitVectorRef, range_ CFRange, value CFBit)  {
	if _cFBitVectorSetBits == nil {
		panic("CoreFoundation: symbol CFBitVectorSetBits not loaded")
	}
	_cFBitVectorSetBits(bv, range_, value)
}


var _cFBitVectorSetCount func(bv CFMutableBitVectorRef, count int) 

// CFBitVectorSetCount changes the size of a mutable bit vector.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetCount(_:_:)
func CFBitVectorSetCount(bv CFMutableBitVectorRef, count int)  {
	if _cFBitVectorSetCount == nil {
		panic("CoreFoundation: symbol CFBitVectorSetCount not loaded")
	}
	_cFBitVectorSetCount(bv, count)
}


var _cFBooleanGetTypeID func() uint

// CFBooleanGetTypeID returns the Core Foundation type identifier for the CFBoolean opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetTypeID()
func CFBooleanGetTypeID() uint {
	if _cFBooleanGetTypeID == nil {
		panic("CoreFoundation: symbol CFBooleanGetTypeID not loaded")
	}
	return _cFBooleanGetTypeID()
}


var _cFBooleanGetValue func(boolean CFBooleanRef) bool

// CFBooleanGetValue returns the value of a CFBoolean object as a standard C type [Boolean].
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetValue(_:)
func CFBooleanGetValue(boolean CFBooleanRef) bool {
	if _cFBooleanGetValue == nil {
		panic("CoreFoundation: symbol CFBooleanGetValue not loaded")
	}
	return _cFBooleanGetValue(boolean)
}


var _cFBundleCopyAuxiliaryExecutableURL func(bundle CFBundleRef, executableName CFStringRef) CFURLRef

// CFBundleCopyAuxiliaryExecutableURL returns the location of a bundle’s auxiliary executable code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyAuxiliaryExecutableURL(_:_:)
func CFBundleCopyAuxiliaryExecutableURL(bundle CFBundleRef, executableName CFStringRef) CFURLRef {
	if _cFBundleCopyAuxiliaryExecutableURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyAuxiliaryExecutableURL not loaded")
	}
	return _cFBundleCopyAuxiliaryExecutableURL(bundle, executableName)
}


var _cFBundleCopyBuiltInPlugInsURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopyBuiltInPlugInsURL returns the location of a bundle’s built in plug-in.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBuiltInPlugInsURL(_:)
func CFBundleCopyBuiltInPlugInsURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopyBuiltInPlugInsURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyBuiltInPlugInsURL not loaded")
	}
	return _cFBundleCopyBuiltInPlugInsURL(bundle)
}


var _cFBundleCopyBundleLocalizations func(bundle CFBundleRef) CFArrayRef

// CFBundleCopyBundleLocalizations returns an array containing a bundle’s localizations.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleLocalizations(_:)
func CFBundleCopyBundleLocalizations(bundle CFBundleRef) CFArrayRef {
	if _cFBundleCopyBundleLocalizations == nil {
		panic("CoreFoundation: symbol CFBundleCopyBundleLocalizations not loaded")
	}
	return _cFBundleCopyBundleLocalizations(bundle)
}


var _cFBundleCopyBundleURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopyBundleURL returns the location of a bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleURL(_:)
func CFBundleCopyBundleURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopyBundleURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyBundleURL not loaded")
	}
	return _cFBundleCopyBundleURL(bundle)
}


var _cFBundleCopyExecutableArchitectures func(bundle CFBundleRef) CFArrayRef

// CFBundleCopyExecutableArchitectures returns an array of CFNumbers representing the architectures a given bundle provides.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitectures(_:)
func CFBundleCopyExecutableArchitectures(bundle CFBundleRef) CFArrayRef {
	if _cFBundleCopyExecutableArchitectures == nil {
		panic("CoreFoundation: symbol CFBundleCopyExecutableArchitectures not loaded")
	}
	return _cFBundleCopyExecutableArchitectures(bundle)
}


var _cFBundleCopyExecutableArchitecturesForURL func(url CFURLRef) CFArrayRef

// CFBundleCopyExecutableArchitecturesForURL returns an array of CFNumbers representing the architectures a given URL provides.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitecturesForURL(_:)
func CFBundleCopyExecutableArchitecturesForURL(url CFURLRef) CFArrayRef {
	if _cFBundleCopyExecutableArchitecturesForURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyExecutableArchitecturesForURL not loaded")
	}
	return _cFBundleCopyExecutableArchitecturesForURL(url)
}


var _cFBundleCopyExecutableURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopyExecutableURL returns the location of a bundle’s main executable code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableURL(_:)
func CFBundleCopyExecutableURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopyExecutableURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyExecutableURL not loaded")
	}
	return _cFBundleCopyExecutableURL(bundle)
}


var _cFBundleCopyInfoDictionaryForURL func(url CFURLRef) CFDictionaryRef

// CFBundleCopyInfoDictionaryForURL returns the information dictionary for a given URL location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryForURL(_:)
func CFBundleCopyInfoDictionaryForURL(url CFURLRef) CFDictionaryRef {
	if _cFBundleCopyInfoDictionaryForURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyInfoDictionaryForURL not loaded")
	}
	return _cFBundleCopyInfoDictionaryForURL(url)
}


var _cFBundleCopyInfoDictionaryInDirectory func(bundleURL CFURLRef) CFDictionaryRef

// CFBundleCopyInfoDictionaryInDirectory returns a bundle’s information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryInDirectory(_:)
func CFBundleCopyInfoDictionaryInDirectory(bundleURL CFURLRef) CFDictionaryRef {
	if _cFBundleCopyInfoDictionaryInDirectory == nil {
		panic("CoreFoundation: symbol CFBundleCopyInfoDictionaryInDirectory not loaded")
	}
	return _cFBundleCopyInfoDictionaryInDirectory(bundleURL)
}


var _cFBundleCopyLocalizationsForPreferences func(locArray CFArrayRef, prefArray CFArrayRef) CFArrayRef

// CFBundleCopyLocalizationsForPreferences given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForPreferences(_:_:)
func CFBundleCopyLocalizationsForPreferences(locArray CFArrayRef, prefArray CFArrayRef) CFArrayRef {
	if _cFBundleCopyLocalizationsForPreferences == nil {
		panic("CoreFoundation: symbol CFBundleCopyLocalizationsForPreferences not loaded")
	}
	return _cFBundleCopyLocalizationsForPreferences(locArray, prefArray)
}


var _cFBundleCopyLocalizationsForURL func(url CFURLRef) CFArrayRef

// CFBundleCopyLocalizationsForURL returns an array containing the localizations for a bundle or executable at a particular location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForURL(_:)
func CFBundleCopyLocalizationsForURL(url CFURLRef) CFArrayRef {
	if _cFBundleCopyLocalizationsForURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyLocalizationsForURL not loaded")
	}
	return _cFBundleCopyLocalizationsForURL(url)
}


var _cFBundleCopyLocalizedString func(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef) CFStringRef

// CFBundleCopyLocalizedString returns a localized string from a bundle’s strings file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedString(_:_:_:_:)
func CFBundleCopyLocalizedString(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef) CFStringRef {
	if _cFBundleCopyLocalizedString == nil {
		panic("CoreFoundation: symbol CFBundleCopyLocalizedString not loaded")
	}
	return _cFBundleCopyLocalizedString(bundle, key, value, tableName)
}


var _cFBundleCopyLocalizedStringForLocalizations func(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef, localizations CFArrayRef) CFStringRef

// CFBundleCopyLocalizedStringForLocalizations returns a localized string from a bundle’s strings file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedStringForLocalizations(_:_:_:_:_:)
func CFBundleCopyLocalizedStringForLocalizations(bundle CFBundleRef, key CFStringRef, value CFStringRef, tableName CFStringRef, localizations CFArrayRef) CFStringRef {
	if _cFBundleCopyLocalizedStringForLocalizations == nil {
		panic("CoreFoundation: symbol CFBundleCopyLocalizedStringForLocalizations not loaded")
	}
	return _cFBundleCopyLocalizedStringForLocalizations(bundle, key, value, tableName, localizations)
}


var _cFBundleCopyPreferredLocalizationsFromArray func(locArray CFArrayRef) CFArrayRef

// CFBundleCopyPreferredLocalizationsFromArray given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPreferredLocalizationsFromArray(_:)
func CFBundleCopyPreferredLocalizationsFromArray(locArray CFArrayRef) CFArrayRef {
	if _cFBundleCopyPreferredLocalizationsFromArray == nil {
		panic("CoreFoundation: symbol CFBundleCopyPreferredLocalizationsFromArray not loaded")
	}
	return _cFBundleCopyPreferredLocalizationsFromArray(locArray)
}


var _cFBundleCopyPrivateFrameworksURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopyPrivateFrameworksURL returns the location of a bundle’s private Frameworks directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPrivateFrameworksURL(_:)
func CFBundleCopyPrivateFrameworksURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopyPrivateFrameworksURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyPrivateFrameworksURL not loaded")
	}
	return _cFBundleCopyPrivateFrameworksURL(bundle)
}


var _cFBundleCopyResourceURL func(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef

// CFBundleCopyResourceURL returns the location of a resource contained in the specified bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURL(_:_:_:_:)
func CFBundleCopyResourceURL(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef {
	if _cFBundleCopyResourceURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourceURL not loaded")
	}
	return _cFBundleCopyResourceURL(bundle, resourceName, resourceType, subDirName)
}


var _cFBundleCopyResourceURLForLocalization func(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFURLRef

// CFBundleCopyResourceURLForLocalization returns the location of a localized resource in a bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLForLocalization(_:_:_:_:_:)
func CFBundleCopyResourceURLForLocalization(bundle CFBundleRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFURLRef {
	if _cFBundleCopyResourceURLForLocalization == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourceURLForLocalization not loaded")
	}
	return _cFBundleCopyResourceURLForLocalization(bundle, resourceName, resourceType, subDirName, localizationName)
}


var _cFBundleCopyResourceURLInDirectory func(bundleURL CFURLRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef

// CFBundleCopyResourceURLInDirectory returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLInDirectory(_:_:_:_:)
func CFBundleCopyResourceURLInDirectory(bundleURL CFURLRef, resourceName CFStringRef, resourceType CFStringRef, subDirName CFStringRef) CFURLRef {
	if _cFBundleCopyResourceURLInDirectory == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourceURLInDirectory not loaded")
	}
	return _cFBundleCopyResourceURLInDirectory(bundleURL, resourceName, resourceType, subDirName)
}


var _cFBundleCopyResourceURLsOfType func(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef

// CFBundleCopyResourceURLsOfType assembles an array of URLs specifying all of the resources of the specified type found in a bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfType(_:_:_:)
func CFBundleCopyResourceURLsOfType(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef {
	if _cFBundleCopyResourceURLsOfType == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourceURLsOfType not loaded")
	}
	return _cFBundleCopyResourceURLsOfType(bundle, resourceType, subDirName)
}


var _cFBundleCopyResourceURLsOfTypeForLocalization func(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFArrayRef

// CFBundleCopyResourceURLsOfTypeForLocalization returns an array containing copies of the URL locations for a specified bundle, resource, and localization name.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeForLocalization(_:_:_:_:)
func CFBundleCopyResourceURLsOfTypeForLocalization(bundle CFBundleRef, resourceType CFStringRef, subDirName CFStringRef, localizationName CFStringRef) CFArrayRef {
	if _cFBundleCopyResourceURLsOfTypeForLocalization == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourceURLsOfTypeForLocalization not loaded")
	}
	return _cFBundleCopyResourceURLsOfTypeForLocalization(bundle, resourceType, subDirName, localizationName)
}


var _cFBundleCopyResourceURLsOfTypeInDirectory func(bundleURL CFURLRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef

// CFBundleCopyResourceURLsOfTypeInDirectory returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeInDirectory(_:_:_:)
func CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL CFURLRef, resourceType CFStringRef, subDirName CFStringRef) CFArrayRef {
	if _cFBundleCopyResourceURLsOfTypeInDirectory == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourceURLsOfTypeInDirectory not loaded")
	}
	return _cFBundleCopyResourceURLsOfTypeInDirectory(bundleURL, resourceType, subDirName)
}


var _cFBundleCopyResourcesDirectoryURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopyResourcesDirectoryURL returns the location of a bundle’s Resources directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourcesDirectoryURL(_:)
func CFBundleCopyResourcesDirectoryURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopyResourcesDirectoryURL == nil {
		panic("CoreFoundation: symbol CFBundleCopyResourcesDirectoryURL not loaded")
	}
	return _cFBundleCopyResourcesDirectoryURL(bundle)
}


var _cFBundleCopySharedFrameworksURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopySharedFrameworksURL returns the location of a bundle’s shared frameworks directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedFrameworksURL(_:)
func CFBundleCopySharedFrameworksURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopySharedFrameworksURL == nil {
		panic("CoreFoundation: symbol CFBundleCopySharedFrameworksURL not loaded")
	}
	return _cFBundleCopySharedFrameworksURL(bundle)
}


var _cFBundleCopySharedSupportURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopySharedSupportURL returns the location of a bundle’s shared support files directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedSupportURL(_:)
func CFBundleCopySharedSupportURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopySharedSupportURL == nil {
		panic("CoreFoundation: symbol CFBundleCopySharedSupportURL not loaded")
	}
	return _cFBundleCopySharedSupportURL(bundle)
}


var _cFBundleCopySupportFilesDirectoryURL func(bundle CFBundleRef) CFURLRef

// CFBundleCopySupportFilesDirectoryURL returns the location of the bundle’s support files directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySupportFilesDirectoryURL(_:)
func CFBundleCopySupportFilesDirectoryURL(bundle CFBundleRef) CFURLRef {
	if _cFBundleCopySupportFilesDirectoryURL == nil {
		panic("CoreFoundation: symbol CFBundleCopySupportFilesDirectoryURL not loaded")
	}
	return _cFBundleCopySupportFilesDirectoryURL(bundle)
}


var _cFBundleCreate func(allocator CFAllocatorRef, bundleURL CFURLRef) CFBundleRef

// CFBundleCreate creates a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreate(_:_:)
func CFBundleCreate(allocator CFAllocatorRef, bundleURL CFURLRef) CFBundleRef {
	if _cFBundleCreate == nil {
		panic("CoreFoundation: symbol CFBundleCreate not loaded")
	}
	return _cFBundleCreate(allocator, bundleURL)
}


var _cFBundleCreateBundlesFromDirectory func(allocator CFAllocatorRef, directoryURL CFURLRef, bundleType CFStringRef) CFArrayRef

// CFBundleCreateBundlesFromDirectory searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreateBundlesFromDirectory(_:_:_:)
func CFBundleCreateBundlesFromDirectory(allocator CFAllocatorRef, directoryURL CFURLRef, bundleType CFStringRef) CFArrayRef {
	if _cFBundleCreateBundlesFromDirectory == nil {
		panic("CoreFoundation: symbol CFBundleCreateBundlesFromDirectory not loaded")
	}
	return _cFBundleCreateBundlesFromDirectory(allocator, directoryURL, bundleType)
}


var _cFBundleGetAllBundles func() CFArrayRef

// CFBundleGetAllBundles returns an array containing all of the bundles currently open in the application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetAllBundles()
func CFBundleGetAllBundles() CFArrayRef {
	if _cFBundleGetAllBundles == nil {
		panic("CoreFoundation: symbol CFBundleGetAllBundles not loaded")
	}
	return _cFBundleGetAllBundles()
}


var _cFBundleGetBundleWithIdentifier func(bundleID CFStringRef) CFBundleRef

// CFBundleGetBundleWithIdentifier locate a bundle given its program-defined identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetBundleWithIdentifier(_:)
func CFBundleGetBundleWithIdentifier(bundleID CFStringRef) CFBundleRef {
	if _cFBundleGetBundleWithIdentifier == nil {
		panic("CoreFoundation: symbol CFBundleGetBundleWithIdentifier not loaded")
	}
	return _cFBundleGetBundleWithIdentifier(bundleID)
}


var _cFBundleGetDataPointerForName func(bundle CFBundleRef, symbolName CFStringRef) unsafe.Pointer

// CFBundleGetDataPointerForName returns a data pointer to a symbol of the given name.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointerForName(_:_:)
func CFBundleGetDataPointerForName(bundle CFBundleRef, symbolName CFStringRef) unsafe.Pointer {
	if _cFBundleGetDataPointerForName == nil {
		panic("CoreFoundation: symbol CFBundleGetDataPointerForName not loaded")
	}
	return _cFBundleGetDataPointerForName(bundle, symbolName)
}


var _cFBundleGetDataPointersForNames func(bundle CFBundleRef, symbolNames CFArrayRef, stbl unsafe.Pointer) 

// CFBundleGetDataPointersForNames returns a C array of data pointer to symbols of the given names.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointersForNames(_:_:_:)
func CFBundleGetDataPointersForNames(bundle CFBundleRef, symbolNames CFArrayRef, stbl unsafe.Pointer)  {
	if _cFBundleGetDataPointersForNames == nil {
		panic("CoreFoundation: symbol CFBundleGetDataPointersForNames not loaded")
	}
	_cFBundleGetDataPointersForNames(bundle, symbolNames, stbl)
}


var _cFBundleGetDevelopmentRegion func(bundle CFBundleRef) CFStringRef

// CFBundleGetDevelopmentRegion returns the bundle’s development region from the bundle’s information property list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDevelopmentRegion(_:)
func CFBundleGetDevelopmentRegion(bundle CFBundleRef) CFStringRef {
	if _cFBundleGetDevelopmentRegion == nil {
		panic("CoreFoundation: symbol CFBundleGetDevelopmentRegion not loaded")
	}
	return _cFBundleGetDevelopmentRegion(bundle)
}


var _cFBundleGetFunctionPointerForName func(bundle CFBundleRef, functionName CFStringRef) unsafe.Pointer

// CFBundleGetFunctionPointerForName returns a pointer to a function in a bundle’s executable code using the function name as the search key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointerForName(_:_:)
func CFBundleGetFunctionPointerForName(bundle CFBundleRef, functionName CFStringRef) unsafe.Pointer {
	if _cFBundleGetFunctionPointerForName == nil {
		panic("CoreFoundation: symbol CFBundleGetFunctionPointerForName not loaded")
	}
	return _cFBundleGetFunctionPointerForName(bundle, functionName)
}


var _cFBundleGetFunctionPointersForNames func(bundle CFBundleRef, functionNames CFArrayRef, ftbl unsafe.Pointer) 

// CFBundleGetFunctionPointersForNames constructs a function table containing pointers to all of the functions found in a bundle’s main executable code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointersForNames(_:_:_:)
func CFBundleGetFunctionPointersForNames(bundle CFBundleRef, functionNames CFArrayRef, ftbl unsafe.Pointer)  {
	if _cFBundleGetFunctionPointersForNames == nil {
		panic("CoreFoundation: symbol CFBundleGetFunctionPointersForNames not loaded")
	}
	_cFBundleGetFunctionPointersForNames(bundle, functionNames, ftbl)
}


var _cFBundleGetIdentifier func(bundle CFBundleRef) CFStringRef

// CFBundleGetIdentifier returns the bundle identifier from a bundle’s information property list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetIdentifier(_:)
func CFBundleGetIdentifier(bundle CFBundleRef) CFStringRef {
	if _cFBundleGetIdentifier == nil {
		panic("CoreFoundation: symbol CFBundleGetIdentifier not loaded")
	}
	return _cFBundleGetIdentifier(bundle)
}


var _cFBundleGetInfoDictionary func(bundle CFBundleRef) CFDictionaryRef

// CFBundleGetInfoDictionary returns a bundle’s information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetInfoDictionary(_:)
func CFBundleGetInfoDictionary(bundle CFBundleRef) CFDictionaryRef {
	if _cFBundleGetInfoDictionary == nil {
		panic("CoreFoundation: symbol CFBundleGetInfoDictionary not loaded")
	}
	return _cFBundleGetInfoDictionary(bundle)
}


var _cFBundleGetLocalInfoDictionary func(bundle CFBundleRef) CFDictionaryRef

// CFBundleGetLocalInfoDictionary returns a bundle’s localized information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetLocalInfoDictionary(_:)
func CFBundleGetLocalInfoDictionary(bundle CFBundleRef) CFDictionaryRef {
	if _cFBundleGetLocalInfoDictionary == nil {
		panic("CoreFoundation: symbol CFBundleGetLocalInfoDictionary not loaded")
	}
	return _cFBundleGetLocalInfoDictionary(bundle)
}


var _cFBundleGetMainBundle func() CFBundleRef

// CFBundleGetMainBundle returns an application’s main bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetMainBundle()
func CFBundleGetMainBundle() CFBundleRef {
	if _cFBundleGetMainBundle == nil {
		panic("CoreFoundation: symbol CFBundleGetMainBundle not loaded")
	}
	return _cFBundleGetMainBundle()
}


var _cFBundleGetPackageInfo func(bundle CFBundleRef, packageType *uint32, packageCreator *uint32) 

// CFBundleGetPackageInfo returns a bundle’s package type and creator.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfo(_:_:_:)
func CFBundleGetPackageInfo(bundle CFBundleRef, packageType *uint32, packageCreator *uint32)  {
	if _cFBundleGetPackageInfo == nil {
		panic("CoreFoundation: symbol CFBundleGetPackageInfo not loaded")
	}
	_cFBundleGetPackageInfo(bundle, packageType, packageCreator)
}


var _cFBundleGetPackageInfoInDirectory func(url CFURLRef, packageType *uint32, packageCreator *uint32) bool

// CFBundleGetPackageInfoInDirectory returns a bundle’s package type and creator without having to create a CFBundle object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfoInDirectory(_:_:_:)
func CFBundleGetPackageInfoInDirectory(url CFURLRef, packageType *uint32, packageCreator *uint32) bool {
	if _cFBundleGetPackageInfoInDirectory == nil {
		panic("CoreFoundation: symbol CFBundleGetPackageInfoInDirectory not loaded")
	}
	return _cFBundleGetPackageInfoInDirectory(url, packageType, packageCreator)
}


var _cFBundleGetPlugIn func(bundle CFBundleRef) CFPlugInRef

// CFBundleGetPlugIn returns a bundle’s plug-in.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPlugIn(_:)
func CFBundleGetPlugIn(bundle CFBundleRef) CFPlugInRef {
	if _cFBundleGetPlugIn == nil {
		panic("CoreFoundation: symbol CFBundleGetPlugIn not loaded")
	}
	return _cFBundleGetPlugIn(bundle)
}


var _cFBundleGetTypeID func() uint

// CFBundleGetTypeID returns the type identifier for the CFBundle opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetTypeID()
func CFBundleGetTypeID() uint {
	if _cFBundleGetTypeID == nil {
		panic("CoreFoundation: symbol CFBundleGetTypeID not loaded")
	}
	return _cFBundleGetTypeID()
}


var _cFBundleGetValueForInfoDictionaryKey func(bundle CFBundleRef, key CFStringRef) CFTypeRef

// CFBundleGetValueForInfoDictionaryKey returns a value (localized if possible) from a bundle’s information dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetValueForInfoDictionaryKey(_:_:)
func CFBundleGetValueForInfoDictionaryKey(bundle CFBundleRef, key CFStringRef) CFTypeRef {
	if _cFBundleGetValueForInfoDictionaryKey == nil {
		panic("CoreFoundation: symbol CFBundleGetValueForInfoDictionaryKey not loaded")
	}
	return _cFBundleGetValueForInfoDictionaryKey(bundle, key)
}


var _cFBundleGetVersionNumber func(bundle CFBundleRef) uint32

// CFBundleGetVersionNumber returns a bundle’s version number.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetVersionNumber(_:)
func CFBundleGetVersionNumber(bundle CFBundleRef) uint32 {
	if _cFBundleGetVersionNumber == nil {
		panic("CoreFoundation: symbol CFBundleGetVersionNumber not loaded")
	}
	return _cFBundleGetVersionNumber(bundle)
}


var _cFBundleIsArchitectureLoadable func(arch int32) bool

// CFBundleIsArchitectureLoadable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsArchitectureLoadable(_:)
func CFBundleIsArchitectureLoadable(arch int32) bool {
	if _cFBundleIsArchitectureLoadable == nil {
		panic("CoreFoundation: symbol CFBundleIsArchitectureLoadable not loaded")
	}
	return _cFBundleIsArchitectureLoadable(arch)
}


var _cFBundleIsExecutableLoadable func(bundle CFBundleRef) bool

// CFBundleIsExecutableLoadable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadable(_:)
func CFBundleIsExecutableLoadable(bundle CFBundleRef) bool {
	if _cFBundleIsExecutableLoadable == nil {
		panic("CoreFoundation: symbol CFBundleIsExecutableLoadable not loaded")
	}
	return _cFBundleIsExecutableLoadable(bundle)
}


var _cFBundleIsExecutableLoadableForURL func(url CFURLRef) bool

// CFBundleIsExecutableLoadableForURL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadableForURL(_:)
func CFBundleIsExecutableLoadableForURL(url CFURLRef) bool {
	if _cFBundleIsExecutableLoadableForURL == nil {
		panic("CoreFoundation: symbol CFBundleIsExecutableLoadableForURL not loaded")
	}
	return _cFBundleIsExecutableLoadableForURL(url)
}


var _cFBundleIsExecutableLoaded func(bundle CFBundleRef) bool

// CFBundleIsExecutableLoaded obtains information about the load status for a bundle’s main executable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoaded(_:)
func CFBundleIsExecutableLoaded(bundle CFBundleRef) bool {
	if _cFBundleIsExecutableLoaded == nil {
		panic("CoreFoundation: symbol CFBundleIsExecutableLoaded not loaded")
	}
	return _cFBundleIsExecutableLoaded(bundle)
}


var _cFBundleLoadExecutable func(bundle CFBundleRef) bool

// CFBundleLoadExecutable loads a bundle’s main executable code into memory and dynamically links it into the running application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutable(_:)
func CFBundleLoadExecutable(bundle CFBundleRef) bool {
	if _cFBundleLoadExecutable == nil {
		panic("CoreFoundation: symbol CFBundleLoadExecutable not loaded")
	}
	return _cFBundleLoadExecutable(bundle)
}


var _cFBundleLoadExecutableAndReturnError func(bundle CFBundleRef, err *CFErrorRef) bool

// CFBundleLoadExecutableAndReturnError returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutableAndReturnError(_:_:)
func CFBundleLoadExecutableAndReturnError(bundle CFBundleRef, err *CFErrorRef) bool {
	if _cFBundleLoadExecutableAndReturnError == nil {
		panic("CoreFoundation: symbol CFBundleLoadExecutableAndReturnError not loaded")
	}
	return _cFBundleLoadExecutableAndReturnError(bundle, err)
}


var _cFBundlePreflightExecutable func(bundle CFBundleRef, err *CFErrorRef) bool

// CFBundlePreflightExecutable returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundlePreflightExecutable(_:_:)
func CFBundlePreflightExecutable(bundle CFBundleRef, err *CFErrorRef) bool {
	if _cFBundlePreflightExecutable == nil {
		panic("CoreFoundation: symbol CFBundlePreflightExecutable not loaded")
	}
	return _cFBundlePreflightExecutable(bundle, err)
}


var _cFBundleUnloadExecutable func(bundle CFBundleRef) 

// CFBundleUnloadExecutable unloads the main executable for the specified bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFBundleUnloadExecutable(_:)
func CFBundleUnloadExecutable(bundle CFBundleRef)  {
	if _cFBundleUnloadExecutable == nil {
		panic("CoreFoundation: symbol CFBundleUnloadExecutable not loaded")
	}
	_cFBundleUnloadExecutable(bundle)
}



var _cFCalendarAddComponents func(calendar CFCalendarRef, at *CFAbsoluteTime, options uint64, componentDesc *byte) bool

// CFCalendarAddComponents computes the absolute time when specified components are added to a given absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarAddComponents
func CFCalendarAddComponents(calendar CFCalendarRef, at *CFAbsoluteTime, options uint64, componentDesc *byte) bool {
	if _cFCalendarAddComponents == nil {
		panic("CoreFoundation: symbol CFCalendarAddComponents not loaded")
	}
	return _cFCalendarAddComponents(calendar, at, options, componentDesc)
}


var _cFCalendarComposeAbsoluteTime func(calendar CFCalendarRef, at *CFAbsoluteTime, componentDesc *byte) bool

// CFCalendarComposeAbsoluteTime computes the absolute time from components in a description string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarComposeAbsoluteTime
func CFCalendarComposeAbsoluteTime(calendar CFCalendarRef, at *CFAbsoluteTime, componentDesc *byte) bool {
	if _cFCalendarComposeAbsoluteTime == nil {
		panic("CoreFoundation: symbol CFCalendarComposeAbsoluteTime not loaded")
	}
	return _cFCalendarComposeAbsoluteTime(calendar, at, componentDesc)
}


var _cFCalendarCopyCurrent func() CFCalendarRef

// CFCalendarCopyCurrent returns a copy of the logical calendar for the current user.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyCurrent()
func CFCalendarCopyCurrent() CFCalendarRef {
	if _cFCalendarCopyCurrent == nil {
		panic("CoreFoundation: symbol CFCalendarCopyCurrent not loaded")
	}
	return _cFCalendarCopyCurrent()
}


var _cFCalendarCopyLocale func(calendar CFCalendarRef) CFLocaleRef

// CFCalendarCopyLocale returns a locale object for a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyLocale(_:)
func CFCalendarCopyLocale(calendar CFCalendarRef) CFLocaleRef {
	if _cFCalendarCopyLocale == nil {
		panic("CoreFoundation: symbol CFCalendarCopyLocale not loaded")
	}
	return _cFCalendarCopyLocale(calendar)
}


var _cFCalendarCopyTimeZone func(calendar CFCalendarRef) CFTimeZoneRef

// CFCalendarCopyTimeZone returns a time zone object for a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyTimeZone(_:)
func CFCalendarCopyTimeZone(calendar CFCalendarRef) CFTimeZoneRef {
	if _cFCalendarCopyTimeZone == nil {
		panic("CoreFoundation: symbol CFCalendarCopyTimeZone not loaded")
	}
	return _cFCalendarCopyTimeZone(calendar)
}


var _cFCalendarCreateWithIdentifier func(allocator CFAllocatorRef, identifier CFCalendarIdentifier) CFCalendarRef

// CFCalendarCreateWithIdentifier returns a calendar object for the calendar identified by a calendar identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCreateWithIdentifier(_:_:)
func CFCalendarCreateWithIdentifier(allocator CFAllocatorRef, identifier CFCalendarIdentifier) CFCalendarRef {
	if _cFCalendarCreateWithIdentifier == nil {
		panic("CoreFoundation: symbol CFCalendarCreateWithIdentifier not loaded")
	}
	return _cFCalendarCreateWithIdentifier(allocator, identifier)
}


var _cFCalendarDecomposeAbsoluteTime func(calendar CFCalendarRef, at CFAbsoluteTime, componentDesc *byte) bool

// CFCalendarDecomposeAbsoluteTime computes the components which are indicated by the componentDesc description string for the given absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarDecomposeAbsoluteTime
func CFCalendarDecomposeAbsoluteTime(calendar CFCalendarRef, at CFAbsoluteTime, componentDesc *byte) bool {
	if _cFCalendarDecomposeAbsoluteTime == nil {
		panic("CoreFoundation: symbol CFCalendarDecomposeAbsoluteTime not loaded")
	}
	return _cFCalendarDecomposeAbsoluteTime(calendar, at, componentDesc)
}


var _cFCalendarGetComponentDifference func(calendar CFCalendarRef, startingAT CFAbsoluteTime, resultAT CFAbsoluteTime, options uint64, componentDesc *byte) bool

// CFCalendarGetComponentDifference computes the difference between the two absolute times, in terms of specified calendrical components.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetComponentDifference
func CFCalendarGetComponentDifference(calendar CFCalendarRef, startingAT CFAbsoluteTime, resultAT CFAbsoluteTime, options uint64, componentDesc *byte) bool {
	if _cFCalendarGetComponentDifference == nil {
		panic("CoreFoundation: symbol CFCalendarGetComponentDifference not loaded")
	}
	return _cFCalendarGetComponentDifference(calendar, startingAT, resultAT, options, componentDesc)
}


var _cFCalendarGetFirstWeekday func(calendar CFCalendarRef) int

// CFCalendarGetFirstWeekday returns the index of first weekday for a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetFirstWeekday(_:)
func CFCalendarGetFirstWeekday(calendar CFCalendarRef) int {
	if _cFCalendarGetFirstWeekday == nil {
		panic("CoreFoundation: symbol CFCalendarGetFirstWeekday not loaded")
	}
	return _cFCalendarGetFirstWeekday(calendar)
}


var _cFCalendarGetIdentifier func(calendar CFCalendarRef) CFCalendarIdentifier

// CFCalendarGetIdentifier returns the given calendar’s identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetIdentifier(_:)
func CFCalendarGetIdentifier(calendar CFCalendarRef) CFCalendarIdentifier {
	if _cFCalendarGetIdentifier == nil {
		panic("CoreFoundation: symbol CFCalendarGetIdentifier not loaded")
	}
	return _cFCalendarGetIdentifier(calendar)
}


var _cFCalendarGetMaximumRangeOfUnit func(calendar CFCalendarRef, unit CFCalendarUnit) CFRange

// CFCalendarGetMaximumRangeOfUnit returns the maximum range limits of the values that a specified unit can take on in a given calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMaximumRangeOfUnit(_:_:)
func CFCalendarGetMaximumRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit) CFRange {
	if _cFCalendarGetMaximumRangeOfUnit == nil {
		panic("CoreFoundation: symbol CFCalendarGetMaximumRangeOfUnit not loaded")
	}
	return _cFCalendarGetMaximumRangeOfUnit(calendar, unit)
}


var _cFCalendarGetMinimumDaysInFirstWeek func(calendar CFCalendarRef) int

// CFCalendarGetMinimumDaysInFirstWeek returns the minimum number of days in the first week of a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumDaysInFirstWeek(_:)
func CFCalendarGetMinimumDaysInFirstWeek(calendar CFCalendarRef) int {
	if _cFCalendarGetMinimumDaysInFirstWeek == nil {
		panic("CoreFoundation: symbol CFCalendarGetMinimumDaysInFirstWeek not loaded")
	}
	return _cFCalendarGetMinimumDaysInFirstWeek(calendar)
}


var _cFCalendarGetMinimumRangeOfUnit func(calendar CFCalendarRef, unit CFCalendarUnit) CFRange

// CFCalendarGetMinimumRangeOfUnit returns the minimum range limits of the values that a specified unit can take on in a given calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumRangeOfUnit(_:_:)
func CFCalendarGetMinimumRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit) CFRange {
	if _cFCalendarGetMinimumRangeOfUnit == nil {
		panic("CoreFoundation: symbol CFCalendarGetMinimumRangeOfUnit not loaded")
	}
	return _cFCalendarGetMinimumRangeOfUnit(calendar, unit)
}


var _cFCalendarGetOrdinalityOfUnit func(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) int

// CFCalendarGetOrdinalityOfUnit returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetOrdinalityOfUnit(_:_:_:_:)
func CFCalendarGetOrdinalityOfUnit(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) int {
	if _cFCalendarGetOrdinalityOfUnit == nil {
		panic("CoreFoundation: symbol CFCalendarGetOrdinalityOfUnit not loaded")
	}
	return _cFCalendarGetOrdinalityOfUnit(calendar, smallerUnit, biggerUnit, at)
}


var _cFCalendarGetRangeOfUnit func(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) CFRange

// CFCalendarGetRangeOfUnit returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetRangeOfUnit(_:_:_:_:)
func CFCalendarGetRangeOfUnit(calendar CFCalendarRef, smallerUnit CFCalendarUnit, biggerUnit CFCalendarUnit, at CFAbsoluteTime) CFRange {
	if _cFCalendarGetRangeOfUnit == nil {
		panic("CoreFoundation: symbol CFCalendarGetRangeOfUnit not loaded")
	}
	return _cFCalendarGetRangeOfUnit(calendar, smallerUnit, biggerUnit, at)
}


var _cFCalendarGetTimeRangeOfUnit func(calendar CFCalendarRef, unit CFCalendarUnit, at CFAbsoluteTime, startp *CFAbsoluteTime, tip *float64) bool

// CFCalendarGetTimeRangeOfUnit returns by reference the start time and duration of a given calendar unit that contains a given absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTimeRangeOfUnit(_:_:_:_:_:)
func CFCalendarGetTimeRangeOfUnit(calendar CFCalendarRef, unit CFCalendarUnit, at CFAbsoluteTime, startp *CFAbsoluteTime, tip *float64) bool {
	if _cFCalendarGetTimeRangeOfUnit == nil {
		panic("CoreFoundation: symbol CFCalendarGetTimeRangeOfUnit not loaded")
	}
	return _cFCalendarGetTimeRangeOfUnit(calendar, unit, at, startp, tip)
}


var _cFCalendarGetTypeID func() uint

// CFCalendarGetTypeID returns the type identifier for the CFCalendar opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTypeID()
func CFCalendarGetTypeID() uint {
	if _cFCalendarGetTypeID == nil {
		panic("CoreFoundation: symbol CFCalendarGetTypeID not loaded")
	}
	return _cFCalendarGetTypeID()
}


var _cFCalendarSetFirstWeekday func(calendar CFCalendarRef, wkdy int) 

// CFCalendarSetFirstWeekday sets the first weekday for a calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetFirstWeekday(_:_:)
func CFCalendarSetFirstWeekday(calendar CFCalendarRef, wkdy int)  {
	if _cFCalendarSetFirstWeekday == nil {
		panic("CoreFoundation: symbol CFCalendarSetFirstWeekday not loaded")
	}
	_cFCalendarSetFirstWeekday(calendar, wkdy)
}


var _cFCalendarSetLocale func(calendar CFCalendarRef, locale CFLocaleRef) 

// CFCalendarSetLocale sets the locale for a calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetLocale(_:_:)
func CFCalendarSetLocale(calendar CFCalendarRef, locale CFLocaleRef)  {
	if _cFCalendarSetLocale == nil {
		panic("CoreFoundation: symbol CFCalendarSetLocale not loaded")
	}
	_cFCalendarSetLocale(calendar, locale)
}


var _cFCalendarSetMinimumDaysInFirstWeek func(calendar CFCalendarRef, mwd int) 

// CFCalendarSetMinimumDaysInFirstWeek sets the minimum number of days in the first week of a specified calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetMinimumDaysInFirstWeek(_:_:)
func CFCalendarSetMinimumDaysInFirstWeek(calendar CFCalendarRef, mwd int)  {
	if _cFCalendarSetMinimumDaysInFirstWeek == nil {
		panic("CoreFoundation: symbol CFCalendarSetMinimumDaysInFirstWeek not loaded")
	}
	_cFCalendarSetMinimumDaysInFirstWeek(calendar, mwd)
}


var _cFCalendarSetTimeZone func(calendar CFCalendarRef, tz CFTimeZoneRef) 

// CFCalendarSetTimeZone sets the time zone for a calendar.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetTimeZone(_:_:)
func CFCalendarSetTimeZone(calendar CFCalendarRef, tz CFTimeZoneRef)  {
	if _cFCalendarSetTimeZone == nil {
		panic("CoreFoundation: symbol CFCalendarSetTimeZone not loaded")
	}
	_cFCalendarSetTimeZone(calendar, tz)
}


var _cFCharacterSetAddCharactersInRange func(theSet CFMutableCharacterSetRef, theRange CFRange) 

// CFCharacterSetAddCharactersInRange adds a given range to a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInRange(_:_:)
func CFCharacterSetAddCharactersInRange(theSet CFMutableCharacterSetRef, theRange CFRange)  {
	if _cFCharacterSetAddCharactersInRange == nil {
		panic("CoreFoundation: symbol CFCharacterSetAddCharactersInRange not loaded")
	}
	_cFCharacterSetAddCharactersInRange(theSet, theRange)
}


var _cFCharacterSetAddCharactersInString func(theSet CFMutableCharacterSetRef, theString CFStringRef) 

// CFCharacterSetAddCharactersInString adds the characters in a given string to a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInString(_:_:)
func CFCharacterSetAddCharactersInString(theSet CFMutableCharacterSetRef, theString CFStringRef)  {
	if _cFCharacterSetAddCharactersInString == nil {
		panic("CoreFoundation: symbol CFCharacterSetAddCharactersInString not loaded")
	}
	_cFCharacterSetAddCharactersInString(theSet, theString)
}


var _cFCharacterSetCreateBitmapRepresentation func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFDataRef

// CFCharacterSetCreateBitmapRepresentation creates a new immutable data with the bitmap representation from the given character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateBitmapRepresentation(_:_:)
func CFCharacterSetCreateBitmapRepresentation(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFDataRef {
	if _cFCharacterSetCreateBitmapRepresentation == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateBitmapRepresentation not loaded")
	}
	return _cFCharacterSetCreateBitmapRepresentation(alloc, theSet)
}


var _cFCharacterSetCreateCopy func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef

// CFCharacterSetCreateCopy creates a new character set with the values from a given character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateCopy(_:_:)
func CFCharacterSetCreateCopy(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef {
	if _cFCharacterSetCreateCopy == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateCopy not loaded")
	}
	return _cFCharacterSetCreateCopy(alloc, theSet)
}


var _cFCharacterSetCreateInvertedSet func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef

// CFCharacterSetCreateInvertedSet creates a new immutable character set that is the invert of the specified character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateInvertedSet(_:_:)
func CFCharacterSetCreateInvertedSet(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFCharacterSetRef {
	if _cFCharacterSetCreateInvertedSet == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateInvertedSet not loaded")
	}
	return _cFCharacterSetCreateInvertedSet(alloc, theSet)
}


var _cFCharacterSetCreateMutable func(alloc CFAllocatorRef) CFMutableCharacterSetRef

// CFCharacterSetCreateMutable creates a new empty mutable character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutable(_:)
func CFCharacterSetCreateMutable(alloc CFAllocatorRef) CFMutableCharacterSetRef {
	if _cFCharacterSetCreateMutable == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateMutable not loaded")
	}
	return _cFCharacterSetCreateMutable(alloc)
}


var _cFCharacterSetCreateMutableCopy func(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFMutableCharacterSetRef

// CFCharacterSetCreateMutableCopy creates a new mutable character set with the values from another character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutableCopy(_:_:)
func CFCharacterSetCreateMutableCopy(alloc CFAllocatorRef, theSet CFCharacterSetRef) CFMutableCharacterSetRef {
	if _cFCharacterSetCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateMutableCopy not loaded")
	}
	return _cFCharacterSetCreateMutableCopy(alloc, theSet)
}


var _cFCharacterSetCreateWithBitmapRepresentation func(alloc CFAllocatorRef, theData CFDataRef) CFCharacterSetRef

// CFCharacterSetCreateWithBitmapRepresentation creates a new immutable character set with the bitmap representation specified by given data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithBitmapRepresentation(_:_:)
func CFCharacterSetCreateWithBitmapRepresentation(alloc CFAllocatorRef, theData CFDataRef) CFCharacterSetRef {
	if _cFCharacterSetCreateWithBitmapRepresentation == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateWithBitmapRepresentation not loaded")
	}
	return _cFCharacterSetCreateWithBitmapRepresentation(alloc, theData)
}


var _cFCharacterSetCreateWithCharactersInRange func(alloc CFAllocatorRef, theRange CFRange) CFCharacterSetRef

// CFCharacterSetCreateWithCharactersInRange creates a new character set with the values from the given range of Unicode characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInRange(_:_:)
func CFCharacterSetCreateWithCharactersInRange(alloc CFAllocatorRef, theRange CFRange) CFCharacterSetRef {
	if _cFCharacterSetCreateWithCharactersInRange == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateWithCharactersInRange not loaded")
	}
	return _cFCharacterSetCreateWithCharactersInRange(alloc, theRange)
}


var _cFCharacterSetCreateWithCharactersInString func(alloc CFAllocatorRef, theString CFStringRef) CFCharacterSetRef

// CFCharacterSetCreateWithCharactersInString creates a new character set with the values in the given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInString(_:_:)
func CFCharacterSetCreateWithCharactersInString(alloc CFAllocatorRef, theString CFStringRef) CFCharacterSetRef {
	if _cFCharacterSetCreateWithCharactersInString == nil {
		panic("CoreFoundation: symbol CFCharacterSetCreateWithCharactersInString not loaded")
	}
	return _cFCharacterSetCreateWithCharactersInString(alloc, theString)
}


var _cFCharacterSetGetPredefined func(theSetIdentifier CFCharacterSetPredefinedSet) CFCharacterSetRef

// CFCharacterSetGetPredefined returns a predefined character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetPredefined(_:)
func CFCharacterSetGetPredefined(theSetIdentifier CFCharacterSetPredefinedSet) CFCharacterSetRef {
	if _cFCharacterSetGetPredefined == nil {
		panic("CoreFoundation: symbol CFCharacterSetGetPredefined not loaded")
	}
	return _cFCharacterSetGetPredefined(theSetIdentifier)
}


var _cFCharacterSetGetTypeID func() uint

// CFCharacterSetGetTypeID returns the type identifier of the CFCharacterSet opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetTypeID()
func CFCharacterSetGetTypeID() uint {
	if _cFCharacterSetGetTypeID == nil {
		panic("CoreFoundation: symbol CFCharacterSetGetTypeID not loaded")
	}
	return _cFCharacterSetGetTypeID()
}


var _cFCharacterSetHasMemberInPlane func(theSet CFCharacterSetRef, thePlane int) bool

// CFCharacterSetHasMemberInPlane reports whether or not a character set contains at least one member character in the specified plane.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetHasMemberInPlane(_:_:)
func CFCharacterSetHasMemberInPlane(theSet CFCharacterSetRef, thePlane int) bool {
	if _cFCharacterSetHasMemberInPlane == nil {
		panic("CoreFoundation: symbol CFCharacterSetHasMemberInPlane not loaded")
	}
	return _cFCharacterSetHasMemberInPlane(theSet, thePlane)
}


var _cFCharacterSetIntersect func(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef) 

// CFCharacterSetIntersect forms an intersection of two character sets.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIntersect(_:_:)
func CFCharacterSetIntersect(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef)  {
	if _cFCharacterSetIntersect == nil {
		panic("CoreFoundation: symbol CFCharacterSetIntersect not loaded")
	}
	_cFCharacterSetIntersect(theSet, theOtherSet)
}


var _cFCharacterSetInvert func(theSet CFMutableCharacterSetRef) 

// CFCharacterSetInvert inverts the content of a given character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetInvert(_:)
func CFCharacterSetInvert(theSet CFMutableCharacterSetRef)  {
	if _cFCharacterSetInvert == nil {
		panic("CoreFoundation: symbol CFCharacterSetInvert not loaded")
	}
	_cFCharacterSetInvert(theSet)
}


var _cFCharacterSetIsCharacterMember func(theSet CFCharacterSetRef, theChar uint16) bool

// CFCharacterSetIsCharacterMember reports whether or not a given Unicode character is in a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsCharacterMember(_:_:)
func CFCharacterSetIsCharacterMember(theSet CFCharacterSetRef, theChar uint16) bool {
	if _cFCharacterSetIsCharacterMember == nil {
		panic("CoreFoundation: symbol CFCharacterSetIsCharacterMember not loaded")
	}
	return _cFCharacterSetIsCharacterMember(theSet, theChar)
}


var _cFCharacterSetIsLongCharacterMember func(theSet CFCharacterSetRef, theChar uint32) bool

// CFCharacterSetIsLongCharacterMember reports whether or not a given UTF-32 character is in a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsLongCharacterMember(_:_:)
func CFCharacterSetIsLongCharacterMember(theSet CFCharacterSetRef, theChar uint32) bool {
	if _cFCharacterSetIsLongCharacterMember == nil {
		panic("CoreFoundation: symbol CFCharacterSetIsLongCharacterMember not loaded")
	}
	return _cFCharacterSetIsLongCharacterMember(theSet, theChar)
}


var _cFCharacterSetIsSupersetOfSet func(theSet CFCharacterSetRef, theOtherset CFCharacterSetRef) bool

// CFCharacterSetIsSupersetOfSet reports whether or not a character set is a superset of another set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsSupersetOfSet(_:_:)
func CFCharacterSetIsSupersetOfSet(theSet CFCharacterSetRef, theOtherset CFCharacterSetRef) bool {
	if _cFCharacterSetIsSupersetOfSet == nil {
		panic("CoreFoundation: symbol CFCharacterSetIsSupersetOfSet not loaded")
	}
	return _cFCharacterSetIsSupersetOfSet(theSet, theOtherset)
}


var _cFCharacterSetRemoveCharactersInRange func(theSet CFMutableCharacterSetRef, theRange CFRange) 

// CFCharacterSetRemoveCharactersInRange removes a given range of Unicode characters from a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInRange(_:_:)
func CFCharacterSetRemoveCharactersInRange(theSet CFMutableCharacterSetRef, theRange CFRange)  {
	if _cFCharacterSetRemoveCharactersInRange == nil {
		panic("CoreFoundation: symbol CFCharacterSetRemoveCharactersInRange not loaded")
	}
	_cFCharacterSetRemoveCharactersInRange(theSet, theRange)
}


var _cFCharacterSetRemoveCharactersInString func(theSet CFMutableCharacterSetRef, theString CFStringRef) 

// CFCharacterSetRemoveCharactersInString removes the characters in a given string from a character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInString(_:_:)
func CFCharacterSetRemoveCharactersInString(theSet CFMutableCharacterSetRef, theString CFStringRef)  {
	if _cFCharacterSetRemoveCharactersInString == nil {
		panic("CoreFoundation: symbol CFCharacterSetRemoveCharactersInString not loaded")
	}
	_cFCharacterSetRemoveCharactersInString(theSet, theString)
}


var _cFCharacterSetUnion func(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef) 

// CFCharacterSetUnion forms the union of two character sets.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetUnion(_:_:)
func CFCharacterSetUnion(theSet CFMutableCharacterSetRef, theOtherSet CFCharacterSetRef)  {
	if _cFCharacterSetUnion == nil {
		panic("CoreFoundation: symbol CFCharacterSetUnion not loaded")
	}
	_cFCharacterSetUnion(theSet, theOtherSet)
}










var _cFCopyDescription func(cf CFTypeRef) CFStringRef

// CFCopyDescription returns a textual description of a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCopyDescription(_:)
func CFCopyDescription(cf CFTypeRef) CFStringRef {
	if _cFCopyDescription == nil {
		panic("CoreFoundation: symbol CFCopyDescription not loaded")
	}
	return _cFCopyDescription(cf)
}


var _cFCopyTypeIDDescription func(type_id uint) CFStringRef

// CFCopyTypeIDDescription returns a textual description of a Core Foundation type, as identified by its type ID, which can be used when debugging.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFCopyTypeIDDescription(_:)
func CFCopyTypeIDDescription(type_id uint) CFStringRef {
	if _cFCopyTypeIDDescription == nil {
		panic("CoreFoundation: symbol CFCopyTypeIDDescription not loaded")
	}
	return _cFCopyTypeIDDescription(type_id)
}


var _cFDataAppendBytes func(theData CFMutableDataRef, bytes *uint8, length int) 

// CFDataAppendBytes appends the bytes from a byte buffer to the contents of a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataAppendBytes(_:_:_:)
func CFDataAppendBytes(theData CFMutableDataRef, bytes *uint8, length int)  {
	if _cFDataAppendBytes == nil {
		panic("CoreFoundation: symbol CFDataAppendBytes not loaded")
	}
	_cFDataAppendBytes(theData, bytes, length)
}


var _cFDataCreate func(allocator CFAllocatorRef, bytes *uint8, length int) CFDataRef

// CFDataCreate creates an immutable CFData object using data copied from a specified byte buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreate(_:_:_:)
func CFDataCreate(allocator CFAllocatorRef, bytes *uint8, length int) CFDataRef {
	if _cFDataCreate == nil {
		panic("CoreFoundation: symbol CFDataCreate not loaded")
	}
	return _cFDataCreate(allocator, bytes, length)
}


var _cFDataCreateCopy func(allocator CFAllocatorRef, theData CFDataRef) CFDataRef

// CFDataCreateCopy creates an immutable copy of a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateCopy(_:_:)
func CFDataCreateCopy(allocator CFAllocatorRef, theData CFDataRef) CFDataRef {
	if _cFDataCreateCopy == nil {
		panic("CoreFoundation: symbol CFDataCreateCopy not loaded")
	}
	return _cFDataCreateCopy(allocator, theData)
}


var _cFDataCreateMutable func(allocator CFAllocatorRef, capacity int) CFMutableDataRef

// CFDataCreateMutable creates an empty CFMutableData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutable(_:_:)
func CFDataCreateMutable(allocator CFAllocatorRef, capacity int) CFMutableDataRef {
	if _cFDataCreateMutable == nil {
		panic("CoreFoundation: symbol CFDataCreateMutable not loaded")
	}
	return _cFDataCreateMutable(allocator, capacity)
}


var _cFDataCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theData CFDataRef) CFMutableDataRef

// CFDataCreateMutableCopy creates a CFMutableData object by copying another CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutableCopy(_:_:_:)
func CFDataCreateMutableCopy(allocator CFAllocatorRef, capacity int, theData CFDataRef) CFMutableDataRef {
	if _cFDataCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFDataCreateMutableCopy not loaded")
	}
	return _cFDataCreateMutableCopy(allocator, capacity, theData)
}


var _cFDataCreateWithBytesNoCopy func(allocator CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFDataRef

// CFDataCreateWithBytesNoCopy creates an immutable CFData object from an external (client-owned) byte buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateWithBytesNoCopy(_:_:_:_:)
func CFDataCreateWithBytesNoCopy(allocator CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFDataRef {
	if _cFDataCreateWithBytesNoCopy == nil {
		panic("CoreFoundation: symbol CFDataCreateWithBytesNoCopy not loaded")
	}
	return _cFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator)
}


var _cFDataDeleteBytes func(theData CFMutableDataRef, range_ CFRange) 

// CFDataDeleteBytes deletes the bytes in a CFMutableData object within a specified range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataDeleteBytes(_:_:)
func CFDataDeleteBytes(theData CFMutableDataRef, range_ CFRange)  {
	if _cFDataDeleteBytes == nil {
		panic("CoreFoundation: symbol CFDataDeleteBytes not loaded")
	}
	_cFDataDeleteBytes(theData, range_)
}


var _cFDataFind func(theData CFDataRef, dataToFind CFDataRef, searchRange CFRange, compareOptions CFDataSearchFlags) CFRange

// CFDataFind finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataFind(_:_:_:_:)
func CFDataFind(theData CFDataRef, dataToFind CFDataRef, searchRange CFRange, compareOptions CFDataSearchFlags) CFRange {
	if _cFDataFind == nil {
		panic("CoreFoundation: symbol CFDataFind not loaded")
	}
	return _cFDataFind(theData, dataToFind, searchRange, compareOptions)
}


var _cFDataGetBytePtr func(theData CFDataRef) *uint8

// CFDataGetBytePtr returns a read-only pointer to the bytes of a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytePtr(_:)
func CFDataGetBytePtr(theData CFDataRef) *uint8 {
	if _cFDataGetBytePtr == nil {
		panic("CoreFoundation: symbol CFDataGetBytePtr not loaded")
	}
	return _cFDataGetBytePtr(theData)
}


var _cFDataGetBytes func(theData CFDataRef, range_ CFRange, buffer *uint8) 

// CFDataGetBytes copies the byte contents of a CFData object to an external buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytes(_:_:_:)
func CFDataGetBytes(theData CFDataRef, range_ CFRange, buffer *uint8)  {
	if _cFDataGetBytes == nil {
		panic("CoreFoundation: symbol CFDataGetBytes not loaded")
	}
	_cFDataGetBytes(theData, range_, buffer)
}


var _cFDataGetLength func(theData CFDataRef) int

// CFDataGetLength returns the number of bytes contained by a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetLength(_:)
func CFDataGetLength(theData CFDataRef) int {
	if _cFDataGetLength == nil {
		panic("CoreFoundation: symbol CFDataGetLength not loaded")
	}
	return _cFDataGetLength(theData)
}


var _cFDataGetMutableBytePtr func(theData CFMutableDataRef) *uint8

// CFDataGetMutableBytePtr returns a pointer to a mutable byte buffer of a CFMutableData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetMutableBytePtr(_:)
func CFDataGetMutableBytePtr(theData CFMutableDataRef) *uint8 {
	if _cFDataGetMutableBytePtr == nil {
		panic("CoreFoundation: symbol CFDataGetMutableBytePtr not loaded")
	}
	return _cFDataGetMutableBytePtr(theData)
}


var _cFDataGetTypeID func() uint

// CFDataGetTypeID returns the type identifier for the CFData opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataGetTypeID()
func CFDataGetTypeID() uint {
	if _cFDataGetTypeID == nil {
		panic("CoreFoundation: symbol CFDataGetTypeID not loaded")
	}
	return _cFDataGetTypeID()
}


var _cFDataIncreaseLength func(theData CFMutableDataRef, extraLength int) 

// CFDataIncreaseLength increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataIncreaseLength(_:_:)
func CFDataIncreaseLength(theData CFMutableDataRef, extraLength int)  {
	if _cFDataIncreaseLength == nil {
		panic("CoreFoundation: symbol CFDataIncreaseLength not loaded")
	}
	_cFDataIncreaseLength(theData, extraLength)
}


var _cFDataReplaceBytes func(theData CFMutableDataRef, range_ CFRange, newBytes *uint8, newLength int) 

// CFDataReplaceBytes replaces those bytes in a CFMutableData object that fall within a specified range with other bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataReplaceBytes(_:_:_:_:)
func CFDataReplaceBytes(theData CFMutableDataRef, range_ CFRange, newBytes *uint8, newLength int)  {
	if _cFDataReplaceBytes == nil {
		panic("CoreFoundation: symbol CFDataReplaceBytes not loaded")
	}
	_cFDataReplaceBytes(theData, range_, newBytes, newLength)
}


var _cFDataSetLength func(theData CFMutableDataRef, length int) 

// CFDataSetLength resets the length of a CFMutableData object’s internal byte buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDataSetLength(_:_:)
func CFDataSetLength(theData CFMutableDataRef, length int)  {
	if _cFDataSetLength == nil {
		panic("CoreFoundation: symbol CFDataSetLength not loaded")
	}
	_cFDataSetLength(theData, length)
}


var _cFDateCompare func(theDate CFDateRef, otherDate CFDateRef, context unsafe.Pointer) CFComparisonResult

// CFDateCompare compares two [CFDate] objects and returns a comparison result.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateCompare(_:_:_:)
func CFDateCompare(theDate CFDateRef, otherDate CFDateRef, context unsafe.Pointer) CFComparisonResult {
	if _cFDateCompare == nil {
		panic("CoreFoundation: symbol CFDateCompare not loaded")
	}
	return _cFDateCompare(theDate, otherDate, context)
}


var _cFDateCreate func(allocator CFAllocatorRef, at CFAbsoluteTime) CFDateRef

// CFDateCreate creates a [CFDate] object given an absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateCreate(_:_:)
func CFDateCreate(allocator CFAllocatorRef, at CFAbsoluteTime) CFDateRef {
	if _cFDateCreate == nil {
		panic("CoreFoundation: symbol CFDateCreate not loaded")
	}
	return _cFDateCreate(allocator, at)
}


var _cFDateFormatterCopyProperty func(formatter CFDateFormatterRef, key CFDateFormatterKey) CFTypeRef

// CFDateFormatterCopyProperty returns a copy of a date formatter’s value for a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCopyProperty(_:_:)
func CFDateFormatterCopyProperty(formatter CFDateFormatterRef, key CFDateFormatterKey) CFTypeRef {
	if _cFDateFormatterCopyProperty == nil {
		panic("CoreFoundation: symbol CFDateFormatterCopyProperty not loaded")
	}
	return _cFDateFormatterCopyProperty(formatter, key)
}


var _cFDateFormatterCreate func(allocator CFAllocatorRef, locale CFLocaleRef, dateStyle CFDateFormatterStyle, timeStyle CFDateFormatterStyle) CFDateFormatterRef

// CFDateFormatterCreate creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreate(_:_:_:_:)
func CFDateFormatterCreate(allocator CFAllocatorRef, locale CFLocaleRef, dateStyle CFDateFormatterStyle, timeStyle CFDateFormatterStyle) CFDateFormatterRef {
	if _cFDateFormatterCreate == nil {
		panic("CoreFoundation: symbol CFDateFormatterCreate not loaded")
	}
	return _cFDateFormatterCreate(allocator, locale, dateStyle, timeStyle)
}


var _cFDateFormatterCreateDateFormatFromTemplate func(allocator CFAllocatorRef, tmplate CFStringRef, options uint64, locale CFLocaleRef) CFStringRef

// CFDateFormatterCreateDateFormatFromTemplate returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFormatFromTemplate(_:_:_:_:)
func CFDateFormatterCreateDateFormatFromTemplate(allocator CFAllocatorRef, tmplate CFStringRef, options uint64, locale CFLocaleRef) CFStringRef {
	if _cFDateFormatterCreateDateFormatFromTemplate == nil {
		panic("CoreFoundation: symbol CFDateFormatterCreateDateFormatFromTemplate not loaded")
	}
	return _cFDateFormatterCreateDateFormatFromTemplate(allocator, tmplate, options, locale)
}


var _cFDateFormatterCreateDateFromString func(allocator CFAllocatorRef, formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange) CFDateRef

// CFDateFormatterCreateDateFromString returns a date object representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFromString(_:_:_:_:)
func CFDateFormatterCreateDateFromString(allocator CFAllocatorRef, formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange) CFDateRef {
	if _cFDateFormatterCreateDateFromString == nil {
		panic("CoreFoundation: symbol CFDateFormatterCreateDateFromString not loaded")
	}
	return _cFDateFormatterCreateDateFromString(allocator, formatter, string_, rangep)
}


var _cFDateFormatterCreateISO8601Formatter func(allocator CFAllocatorRef, formatOptions CFISO8601DateFormatOptions) CFDateFormatterRef

// CFDateFormatterCreateISO8601Formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateISO8601Formatter(_:_:)
func CFDateFormatterCreateISO8601Formatter(allocator CFAllocatorRef, formatOptions CFISO8601DateFormatOptions) CFDateFormatterRef {
	if _cFDateFormatterCreateISO8601Formatter == nil {
		panic("CoreFoundation: symbol CFDateFormatterCreateISO8601Formatter not loaded")
	}
	return _cFDateFormatterCreateISO8601Formatter(allocator, formatOptions)
}


var _cFDateFormatterCreateStringWithAbsoluteTime func(allocator CFAllocatorRef, formatter CFDateFormatterRef, at CFAbsoluteTime) CFStringRef

// CFDateFormatterCreateStringWithAbsoluteTime returns a string representation of the given absolute time using the specified date formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithAbsoluteTime(_:_:_:)
func CFDateFormatterCreateStringWithAbsoluteTime(allocator CFAllocatorRef, formatter CFDateFormatterRef, at CFAbsoluteTime) CFStringRef {
	if _cFDateFormatterCreateStringWithAbsoluteTime == nil {
		panic("CoreFoundation: symbol CFDateFormatterCreateStringWithAbsoluteTime not loaded")
	}
	return _cFDateFormatterCreateStringWithAbsoluteTime(allocator, formatter, at)
}


var _cFDateFormatterCreateStringWithDate func(allocator CFAllocatorRef, formatter CFDateFormatterRef, date CFDateRef) CFStringRef

// CFDateFormatterCreateStringWithDate returns a string representation of the given date using the specified date formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithDate(_:_:_:)
func CFDateFormatterCreateStringWithDate(allocator CFAllocatorRef, formatter CFDateFormatterRef, date CFDateRef) CFStringRef {
	if _cFDateFormatterCreateStringWithDate == nil {
		panic("CoreFoundation: symbol CFDateFormatterCreateStringWithDate not loaded")
	}
	return _cFDateFormatterCreateStringWithDate(allocator, formatter, date)
}


var _cFDateFormatterGetAbsoluteTimeFromString func(formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange, atp *CFAbsoluteTime) bool

// CFDateFormatterGetAbsoluteTimeFromString returns an absolute time object representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetAbsoluteTimeFromString(_:_:_:_:)
func CFDateFormatterGetAbsoluteTimeFromString(formatter CFDateFormatterRef, string_ CFStringRef, rangep *CFRange, atp *CFAbsoluteTime) bool {
	if _cFDateFormatterGetAbsoluteTimeFromString == nil {
		panic("CoreFoundation: symbol CFDateFormatterGetAbsoluteTimeFromString not loaded")
	}
	return _cFDateFormatterGetAbsoluteTimeFromString(formatter, string_, rangep, atp)
}


var _cFDateFormatterGetDateStyle func(formatter CFDateFormatterRef) CFDateFormatterStyle

// CFDateFormatterGetDateStyle returns the date style used to create the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetDateStyle(_:)
func CFDateFormatterGetDateStyle(formatter CFDateFormatterRef) CFDateFormatterStyle {
	if _cFDateFormatterGetDateStyle == nil {
		panic("CoreFoundation: symbol CFDateFormatterGetDateStyle not loaded")
	}
	return _cFDateFormatterGetDateStyle(formatter)
}


var _cFDateFormatterGetFormat func(formatter CFDateFormatterRef) CFStringRef

// CFDateFormatterGetFormat returns a format string for the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetFormat(_:)
func CFDateFormatterGetFormat(formatter CFDateFormatterRef) CFStringRef {
	if _cFDateFormatterGetFormat == nil {
		panic("CoreFoundation: symbol CFDateFormatterGetFormat not loaded")
	}
	return _cFDateFormatterGetFormat(formatter)
}


var _cFDateFormatterGetLocale func(formatter CFDateFormatterRef) CFLocaleRef

// CFDateFormatterGetLocale returns the locale object used to create the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetLocale(_:)
func CFDateFormatterGetLocale(formatter CFDateFormatterRef) CFLocaleRef {
	if _cFDateFormatterGetLocale == nil {
		panic("CoreFoundation: symbol CFDateFormatterGetLocale not loaded")
	}
	return _cFDateFormatterGetLocale(formatter)
}


var _cFDateFormatterGetTimeStyle func(formatter CFDateFormatterRef) CFDateFormatterStyle

// CFDateFormatterGetTimeStyle returns the time style used to create the given date formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTimeStyle(_:)
func CFDateFormatterGetTimeStyle(formatter CFDateFormatterRef) CFDateFormatterStyle {
	if _cFDateFormatterGetTimeStyle == nil {
		panic("CoreFoundation: symbol CFDateFormatterGetTimeStyle not loaded")
	}
	return _cFDateFormatterGetTimeStyle(formatter)
}


var _cFDateFormatterGetTypeID func() uint

// CFDateFormatterGetTypeID returns the type identifier for CFDateFormatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTypeID()
func CFDateFormatterGetTypeID() uint {
	if _cFDateFormatterGetTypeID == nil {
		panic("CoreFoundation: symbol CFDateFormatterGetTypeID not loaded")
	}
	return _cFDateFormatterGetTypeID()
}


var _cFDateFormatterSetFormat func(formatter CFDateFormatterRef, formatString CFStringRef) 

// CFDateFormatterSetFormat sets the format string of the given date formatter to the specified value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetFormat(_:_:)
func CFDateFormatterSetFormat(formatter CFDateFormatterRef, formatString CFStringRef)  {
	if _cFDateFormatterSetFormat == nil {
		panic("CoreFoundation: symbol CFDateFormatterSetFormat not loaded")
	}
	_cFDateFormatterSetFormat(formatter, formatString)
}


var _cFDateFormatterSetProperty func(formatter CFDateFormatterRef, key CFStringRef, value CFTypeRef) 

// CFDateFormatterSetProperty sets a date formatter property using a key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetProperty(_:_:_:)
func CFDateFormatterSetProperty(formatter CFDateFormatterRef, key CFStringRef, value CFTypeRef)  {
	if _cFDateFormatterSetProperty == nil {
		panic("CoreFoundation: symbol CFDateFormatterSetProperty not loaded")
	}
	_cFDateFormatterSetProperty(formatter, key, value)
}


var _cFDateGetAbsoluteTime func(theDate CFDateRef) CFAbsoluteTime

// CFDateGetAbsoluteTime returns a [CFDate] object’s absolute time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateGetAbsoluteTime(_:)
func CFDateGetAbsoluteTime(theDate CFDateRef) CFAbsoluteTime {
	if _cFDateGetAbsoluteTime == nil {
		panic("CoreFoundation: symbol CFDateGetAbsoluteTime not loaded")
	}
	return _cFDateGetAbsoluteTime(theDate)
}


var _cFDateGetTimeIntervalSinceDate func(theDate CFDateRef, otherDate CFDateRef) float64

// CFDateGetTimeIntervalSinceDate returns the number of elapsed seconds between the given [CFDate] objects.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTimeIntervalSinceDate(_:_:)
func CFDateGetTimeIntervalSinceDate(theDate CFDateRef, otherDate CFDateRef) float64 {
	if _cFDateGetTimeIntervalSinceDate == nil {
		panic("CoreFoundation: symbol CFDateGetTimeIntervalSinceDate not loaded")
	}
	return _cFDateGetTimeIntervalSinceDate(theDate, otherDate)
}


var _cFDateGetTypeID func() uint

// CFDateGetTypeID returns the type identifier for the [CFDate] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTypeID()
func CFDateGetTypeID() uint {
	if _cFDateGetTypeID == nil {
		panic("CoreFoundation: symbol CFDateGetTypeID not loaded")
	}
	return _cFDateGetTypeID()
}


var _cFDictionaryAddValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) 

// CFDictionaryAddValue adds a key-value pair to a dictionary if the specified key is not already present.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryAddValue(_:_:_:)
func CFDictionaryAddValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer)  {
	if _cFDictionaryAddValue == nil {
		panic("CoreFoundation: symbol CFDictionaryAddValue not loaded")
	}
	_cFDictionaryAddValue(theDict, key, value)
}


var _cFDictionaryApplyFunction func(theDict CFDictionaryRef, applier CFDictionaryApplierFunction, context unsafe.Pointer) 

// CFDictionaryApplyFunction calls a function once for each key-value pair in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplyFunction(_:_:_:)
func CFDictionaryApplyFunction(theDict CFDictionaryRef, applier CFDictionaryApplierFunction, context unsafe.Pointer)  {
	if _cFDictionaryApplyFunction == nil {
		panic("CoreFoundation: symbol CFDictionaryApplyFunction not loaded")
	}
	_cFDictionaryApplyFunction(theDict, applier, context)
}


var _cFDictionaryContainsKey func(theDict CFDictionaryRef, key unsafe.Pointer) bool

// CFDictionaryContainsKey returns a Boolean value that indicates whether a given key is in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsKey(_:_:)
func CFDictionaryContainsKey(theDict CFDictionaryRef, key unsafe.Pointer) bool {
	if _cFDictionaryContainsKey == nil {
		panic("CoreFoundation: symbol CFDictionaryContainsKey not loaded")
	}
	return _cFDictionaryContainsKey(theDict, key)
}


var _cFDictionaryContainsValue func(theDict CFDictionaryRef, value unsafe.Pointer) bool

// CFDictionaryContainsValue returns a Boolean value that indicates whether a given value is in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsValue(_:_:)
func CFDictionaryContainsValue(theDict CFDictionaryRef, value unsafe.Pointer) bool {
	if _cFDictionaryContainsValue == nil {
		panic("CoreFoundation: symbol CFDictionaryContainsValue not loaded")
	}
	return _cFDictionaryContainsValue(theDict, value)
}


var _cFDictionaryCreate func(allocator CFAllocatorRef, keys unsafe.Pointer, values unsafe.Pointer, numValues int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFDictionaryRef

// CFDictionaryCreate creates an immutable dictionary containing the specified key-value pairs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreate(_:_:_:_:_:_:)
func CFDictionaryCreate(allocator CFAllocatorRef, keys unsafe.Pointer, values unsafe.Pointer, numValues int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFDictionaryRef {
	if _cFDictionaryCreate == nil {
		panic("CoreFoundation: symbol CFDictionaryCreate not loaded")
	}
	return _cFDictionaryCreate(allocator, keys, values, numValues, keyCallBacks, valueCallBacks)
}


var _cFDictionaryCreateCopy func(allocator CFAllocatorRef, theDict CFDictionaryRef) CFDictionaryRef

// CFDictionaryCreateCopy creates and returns a new immutable dictionary with the key-value pairs of another dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateCopy(_:_:)
func CFDictionaryCreateCopy(allocator CFAllocatorRef, theDict CFDictionaryRef) CFDictionaryRef {
	if _cFDictionaryCreateCopy == nil {
		panic("CoreFoundation: symbol CFDictionaryCreateCopy not loaded")
	}
	return _cFDictionaryCreateCopy(allocator, theDict)
}


var _cFDictionaryCreateMutable func(allocator CFAllocatorRef, capacity int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFMutableDictionaryRef

// CFDictionaryCreateMutable creates a new mutable dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutable(_:_:_:_:)
func CFDictionaryCreateMutable(allocator CFAllocatorRef, capacity int, keyCallBacks *CFDictionaryKeyCallBacks, valueCallBacks *CFDictionaryValueCallBacks) CFMutableDictionaryRef {
	if _cFDictionaryCreateMutable == nil {
		panic("CoreFoundation: symbol CFDictionaryCreateMutable not loaded")
	}
	return _cFDictionaryCreateMutable(allocator, capacity, keyCallBacks, valueCallBacks)
}


var _cFDictionaryCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theDict CFDictionaryRef) CFMutableDictionaryRef

// CFDictionaryCreateMutableCopy creates a new mutable dictionary with the key-value pairs from another dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutableCopy(_:_:_:)
func CFDictionaryCreateMutableCopy(allocator CFAllocatorRef, capacity int, theDict CFDictionaryRef) CFMutableDictionaryRef {
	if _cFDictionaryCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFDictionaryCreateMutableCopy not loaded")
	}
	return _cFDictionaryCreateMutableCopy(allocator, capacity, theDict)
}


var _cFDictionaryGetCount func(theDict CFDictionaryRef) int

// CFDictionaryGetCount returns the number of key-value pairs in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCount(_:)
func CFDictionaryGetCount(theDict CFDictionaryRef) int {
	if _cFDictionaryGetCount == nil {
		panic("CoreFoundation: symbol CFDictionaryGetCount not loaded")
	}
	return _cFDictionaryGetCount(theDict)
}


var _cFDictionaryGetCountOfKey func(theDict CFDictionaryRef, key unsafe.Pointer) int

// CFDictionaryGetCountOfKey returns the number of times a key occurs in a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfKey(_:_:)
func CFDictionaryGetCountOfKey(theDict CFDictionaryRef, key unsafe.Pointer) int {
	if _cFDictionaryGetCountOfKey == nil {
		panic("CoreFoundation: symbol CFDictionaryGetCountOfKey not loaded")
	}
	return _cFDictionaryGetCountOfKey(theDict, key)
}


var _cFDictionaryGetCountOfValue func(theDict CFDictionaryRef, value unsafe.Pointer) int

// CFDictionaryGetCountOfValue counts the number of times a given value occurs in the dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfValue(_:_:)
func CFDictionaryGetCountOfValue(theDict CFDictionaryRef, value unsafe.Pointer) int {
	if _cFDictionaryGetCountOfValue == nil {
		panic("CoreFoundation: symbol CFDictionaryGetCountOfValue not loaded")
	}
	return _cFDictionaryGetCountOfValue(theDict, value)
}


var _cFDictionaryGetKeysAndValues func(theDict CFDictionaryRef, keys unsafe.Pointer, values unsafe.Pointer) 

// CFDictionaryGetKeysAndValues fills two buffers with the keys and values from a dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetKeysAndValues(_:_:_:)
func CFDictionaryGetKeysAndValues(theDict CFDictionaryRef, keys unsafe.Pointer, values unsafe.Pointer)  {
	if _cFDictionaryGetKeysAndValues == nil {
		panic("CoreFoundation: symbol CFDictionaryGetKeysAndValues not loaded")
	}
	_cFDictionaryGetKeysAndValues(theDict, keys, values)
}


var _cFDictionaryGetTypeID func() uint

// CFDictionaryGetTypeID returns the type identifier for the CFDictionary opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetTypeID()
func CFDictionaryGetTypeID() uint {
	if _cFDictionaryGetTypeID == nil {
		panic("CoreFoundation: symbol CFDictionaryGetTypeID not loaded")
	}
	return _cFDictionaryGetTypeID()
}


var _cFDictionaryGetValue func(theDict CFDictionaryRef, key unsafe.Pointer) unsafe.Pointer

// CFDictionaryGetValue returns the value associated with a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValue(_:_:)
func CFDictionaryGetValue(theDict CFDictionaryRef, key unsafe.Pointer) unsafe.Pointer {
	if _cFDictionaryGetValue == nil {
		panic("CoreFoundation: symbol CFDictionaryGetValue not loaded")
	}
	return _cFDictionaryGetValue(theDict, key)
}


var _cFDictionaryGetValueIfPresent func(theDict CFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool

// CFDictionaryGetValueIfPresent returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValueIfPresent(_:_:_:)
func CFDictionaryGetValueIfPresent(theDict CFDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) bool {
	if _cFDictionaryGetValueIfPresent == nil {
		panic("CoreFoundation: symbol CFDictionaryGetValueIfPresent not loaded")
	}
	return _cFDictionaryGetValueIfPresent(theDict, key, value)
}


var _cFDictionaryRemoveAllValues func(theDict CFMutableDictionaryRef) 

// CFDictionaryRemoveAllValues removes all the key-value pairs from a dictionary, making it empty.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveAllValues(_:)
func CFDictionaryRemoveAllValues(theDict CFMutableDictionaryRef)  {
	if _cFDictionaryRemoveAllValues == nil {
		panic("CoreFoundation: symbol CFDictionaryRemoveAllValues not loaded")
	}
	_cFDictionaryRemoveAllValues(theDict)
}


var _cFDictionaryRemoveValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer) 

// CFDictionaryRemoveValue removes a key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveValue(_:_:)
func CFDictionaryRemoveValue(theDict CFMutableDictionaryRef, key unsafe.Pointer)  {
	if _cFDictionaryRemoveValue == nil {
		panic("CoreFoundation: symbol CFDictionaryRemoveValue not loaded")
	}
	_cFDictionaryRemoveValue(theDict, key)
}


var _cFDictionaryReplaceValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) 

// CFDictionaryReplaceValue replaces a value corresponding to a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryReplaceValue(_:_:_:)
func CFDictionaryReplaceValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer)  {
	if _cFDictionaryReplaceValue == nil {
		panic("CoreFoundation: symbol CFDictionaryReplaceValue not loaded")
	}
	_cFDictionaryReplaceValue(theDict, key, value)
}


var _cFDictionarySetValue func(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) 

// CFDictionarySetValue sets the value corresponding to a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFDictionarySetValue(_:_:_:)
func CFDictionarySetValue(theDict CFMutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer)  {
	if _cFDictionarySetValue == nil {
		panic("CoreFoundation: symbol CFDictionarySetValue not loaded")
	}
	_cFDictionarySetValue(theDict, key, value)
}


var _cFEqual func(cf1 CFTypeRef, cf2 CFTypeRef) bool

// CFEqual determines whether two Core Foundation objects are considered equal.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFEqual(_:_:)
func CFEqual(cf1 CFTypeRef, cf2 CFTypeRef) bool {
	if _cFEqual == nil {
		panic("CoreFoundation: symbol CFEqual not loaded")
	}
	return _cFEqual(cf1, cf2)
}


var _cFErrorCopyDescription func(err CFErrorRef) CFStringRef

// CFErrorCopyDescription returns a human-presentable description for a given error.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyDescription(_:)
func CFErrorCopyDescription(err CFErrorRef) CFStringRef {
	if _cFErrorCopyDescription == nil {
		panic("CoreFoundation: symbol CFErrorCopyDescription not loaded")
	}
	return _cFErrorCopyDescription(err)
}


var _cFErrorCopyFailureReason func(err CFErrorRef) CFStringRef

// CFErrorCopyFailureReason returns a human-presentable failure reason for a given error.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyFailureReason(_:)
func CFErrorCopyFailureReason(err CFErrorRef) CFStringRef {
	if _cFErrorCopyFailureReason == nil {
		panic("CoreFoundation: symbol CFErrorCopyFailureReason not loaded")
	}
	return _cFErrorCopyFailureReason(err)
}


var _cFErrorCopyRecoverySuggestion func(err CFErrorRef) CFStringRef

// CFErrorCopyRecoverySuggestion returns a human presentable recovery suggestion for a given error.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyRecoverySuggestion(_:)
func CFErrorCopyRecoverySuggestion(err CFErrorRef) CFStringRef {
	if _cFErrorCopyRecoverySuggestion == nil {
		panic("CoreFoundation: symbol CFErrorCopyRecoverySuggestion not loaded")
	}
	return _cFErrorCopyRecoverySuggestion(err)
}


var _cFErrorCopyUserInfo func(err CFErrorRef) CFDictionaryRef

// CFErrorCopyUserInfo returns the user info dictionary for a given CFError.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyUserInfo(_:)
func CFErrorCopyUserInfo(err CFErrorRef) CFDictionaryRef {
	if _cFErrorCopyUserInfo == nil {
		panic("CoreFoundation: symbol CFErrorCopyUserInfo not loaded")
	}
	return _cFErrorCopyUserInfo(err)
}


var _cFErrorCreate func(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfo CFDictionaryRef) CFErrorRef

// CFErrorCreate creates a new CFError object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreate(_:_:_:_:)
func CFErrorCreate(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfo CFDictionaryRef) CFErrorRef {
	if _cFErrorCreate == nil {
		panic("CoreFoundation: symbol CFErrorCreate not loaded")
	}
	return _cFErrorCreate(allocator, domain, code, userInfo)
}


var _cFErrorCreateWithUserInfoKeysAndValues func(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues int) CFErrorRef

// CFErrorCreateWithUserInfoKeysAndValues creates a new CFError object using given keys and values to create the user info dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreateWithUserInfoKeysAndValues(_:_:_:_:_:_:)
func CFErrorCreateWithUserInfoKeysAndValues(allocator CFAllocatorRef, domain CFErrorDomain, code int, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues int) CFErrorRef {
	if _cFErrorCreateWithUserInfoKeysAndValues == nil {
		panic("CoreFoundation: symbol CFErrorCreateWithUserInfoKeysAndValues not loaded")
	}
	return _cFErrorCreateWithUserInfoKeysAndValues(allocator, domain, code, userInfoKeys, userInfoValues, numUserInfoValues)
}


var _cFErrorGetCode func(err CFErrorRef) int

// CFErrorGetCode returns the error code for a given CFError.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetCode(_:)
func CFErrorGetCode(err CFErrorRef) int {
	if _cFErrorGetCode == nil {
		panic("CoreFoundation: symbol CFErrorGetCode not loaded")
	}
	return _cFErrorGetCode(err)
}


var _cFErrorGetDomain func(err CFErrorRef) CFErrorDomain

// CFErrorGetDomain returns the error domain for a given CFError.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetDomain(_:)
func CFErrorGetDomain(err CFErrorRef) CFErrorDomain {
	if _cFErrorGetDomain == nil {
		panic("CoreFoundation: symbol CFErrorGetDomain not loaded")
	}
	return _cFErrorGetDomain(err)
}


var _cFErrorGetTypeID func() uint

// CFErrorGetTypeID returns the type identifier for the CFError opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetTypeID()
func CFErrorGetTypeID() uint {
	if _cFErrorGetTypeID == nil {
		panic("CoreFoundation: symbol CFErrorGetTypeID not loaded")
	}
	return _cFErrorGetTypeID()
}


var _cFFileDescriptorCreate func(allocator CFAllocatorRef, fd CFFileDescriptorNativeDescriptor, closeOnInvalidate bool, callout CFFileDescriptorCallBack, context *CFFileDescriptorContext) CFFileDescriptorRef

// CFFileDescriptorCreate creates a new CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreate(_:_:_:_:_:)
func CFFileDescriptorCreate(allocator CFAllocatorRef, fd CFFileDescriptorNativeDescriptor, closeOnInvalidate bool, callout CFFileDescriptorCallBack, context *CFFileDescriptorContext) CFFileDescriptorRef {
	if _cFFileDescriptorCreate == nil {
		panic("CoreFoundation: symbol CFFileDescriptorCreate not loaded")
	}
	return _cFFileDescriptorCreate(allocator, fd, closeOnInvalidate, callout, context)
}


var _cFFileDescriptorCreateRunLoopSource func(allocator CFAllocatorRef, f CFFileDescriptorRef, order int) CFRunLoopSourceRef

// CFFileDescriptorCreateRunLoopSource creates a new runloop source for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreateRunLoopSource(_:_:_:)
func CFFileDescriptorCreateRunLoopSource(allocator CFAllocatorRef, f CFFileDescriptorRef, order int) CFRunLoopSourceRef {
	if _cFFileDescriptorCreateRunLoopSource == nil {
		panic("CoreFoundation: symbol CFFileDescriptorCreateRunLoopSource not loaded")
	}
	return _cFFileDescriptorCreateRunLoopSource(allocator, f, order)
}


var _cFFileDescriptorDisableCallBacks func(f CFFileDescriptorRef, callBackTypes uint64) 

// CFFileDescriptorDisableCallBacks disables callbacks for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorDisableCallBacks(_:_:)
func CFFileDescriptorDisableCallBacks(f CFFileDescriptorRef, callBackTypes uint64)  {
	if _cFFileDescriptorDisableCallBacks == nil {
		panic("CoreFoundation: symbol CFFileDescriptorDisableCallBacks not loaded")
	}
	_cFFileDescriptorDisableCallBacks(f, callBackTypes)
}


var _cFFileDescriptorEnableCallBacks func(f CFFileDescriptorRef, callBackTypes uint64) 

// CFFileDescriptorEnableCallBacks enables callbacks for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorEnableCallBacks(_:_:)
func CFFileDescriptorEnableCallBacks(f CFFileDescriptorRef, callBackTypes uint64)  {
	if _cFFileDescriptorEnableCallBacks == nil {
		panic("CoreFoundation: symbol CFFileDescriptorEnableCallBacks not loaded")
	}
	_cFFileDescriptorEnableCallBacks(f, callBackTypes)
}


var _cFFileDescriptorGetContext func(f CFFileDescriptorRef, context *CFFileDescriptorContext) 

// CFFileDescriptorGetContext gets the context for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetContext(_:_:)
func CFFileDescriptorGetContext(f CFFileDescriptorRef, context *CFFileDescriptorContext)  {
	if _cFFileDescriptorGetContext == nil {
		panic("CoreFoundation: symbol CFFileDescriptorGetContext not loaded")
	}
	_cFFileDescriptorGetContext(f, context)
}


var _cFFileDescriptorGetNativeDescriptor func(f CFFileDescriptorRef) CFFileDescriptorNativeDescriptor

// CFFileDescriptorGetNativeDescriptor returns the native file descriptor for a given CFFileDescriptor.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetNativeDescriptor(_:)
func CFFileDescriptorGetNativeDescriptor(f CFFileDescriptorRef) CFFileDescriptorNativeDescriptor {
	if _cFFileDescriptorGetNativeDescriptor == nil {
		panic("CoreFoundation: symbol CFFileDescriptorGetNativeDescriptor not loaded")
	}
	return _cFFileDescriptorGetNativeDescriptor(f)
}


var _cFFileDescriptorGetTypeID func() uint

// CFFileDescriptorGetTypeID returns the type identifier for the CFFileDescriptor opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetTypeID()
func CFFileDescriptorGetTypeID() uint {
	if _cFFileDescriptorGetTypeID == nil {
		panic("CoreFoundation: symbol CFFileDescriptorGetTypeID not loaded")
	}
	return _cFFileDescriptorGetTypeID()
}


var _cFFileDescriptorInvalidate func(f CFFileDescriptorRef) 

// CFFileDescriptorInvalidate invalidates a CFFileDescriptor object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorInvalidate(_:)
func CFFileDescriptorInvalidate(f CFFileDescriptorRef)  {
	if _cFFileDescriptorInvalidate == nil {
		panic("CoreFoundation: symbol CFFileDescriptorInvalidate not loaded")
	}
	_cFFileDescriptorInvalidate(f)
}


var _cFFileDescriptorIsValid func(f CFFileDescriptorRef) bool

// CFFileDescriptorIsValid returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorIsValid(_:)
func CFFileDescriptorIsValid(f CFFileDescriptorRef) bool {
	if _cFFileDescriptorIsValid == nil {
		panic("CoreFoundation: symbol CFFileDescriptorIsValid not loaded")
	}
	return _cFFileDescriptorIsValid(f)
}


var _cFFileSecurityClearProperties func(fileSec CFFileSecurityRef, clearPropertyMask CFFileSecurityClearOptions) bool

// CFFileSecurityClearProperties clears properties from a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearProperties(_:_:)
func CFFileSecurityClearProperties(fileSec CFFileSecurityRef, clearPropertyMask CFFileSecurityClearOptions) bool {
	if _cFFileSecurityClearProperties == nil {
		panic("CoreFoundation: symbol CFFileSecurityClearProperties not loaded")
	}
	return _cFFileSecurityClearProperties(fileSec, clearPropertyMask)
}



var _cFFileSecurityCopyGroupUUID func(fileSec CFFileSecurityRef, groupUUID *CFUUIDRef) bool

// CFFileSecurityCopyGroupUUID copies the group UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyGroupUUID(_:_:)
func CFFileSecurityCopyGroupUUID(fileSec CFFileSecurityRef, groupUUID *CFUUIDRef) bool {
	if _cFFileSecurityCopyGroupUUID == nil {
		panic("CoreFoundation: symbol CFFileSecurityCopyGroupUUID not loaded")
	}
	return _cFFileSecurityCopyGroupUUID(fileSec, groupUUID)
}


var _cFFileSecurityCopyOwnerUUID func(fileSec CFFileSecurityRef, ownerUUID *CFUUIDRef) bool

// CFFileSecurityCopyOwnerUUID copies the owner UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyOwnerUUID(_:_:)
func CFFileSecurityCopyOwnerUUID(fileSec CFFileSecurityRef, ownerUUID *CFUUIDRef) bool {
	if _cFFileSecurityCopyOwnerUUID == nil {
		panic("CoreFoundation: symbol CFFileSecurityCopyOwnerUUID not loaded")
	}
	return _cFFileSecurityCopyOwnerUUID(fileSec, ownerUUID)
}


var _cFFileSecurityCreate func(allocator CFAllocatorRef) CFFileSecurityRef

// CFFileSecurityCreate creates a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreate(_:)
func CFFileSecurityCreate(allocator CFAllocatorRef) CFFileSecurityRef {
	if _cFFileSecurityCreate == nil {
		panic("CoreFoundation: symbol CFFileSecurityCreate not loaded")
	}
	return _cFFileSecurityCreate(allocator)
}


var _cFFileSecurityCreateCopy func(allocator CFAllocatorRef, fileSec CFFileSecurityRef) CFFileSecurityRef

// CFFileSecurityCreateCopy creates a copy of a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreateCopy(_:_:)
func CFFileSecurityCreateCopy(allocator CFAllocatorRef, fileSec CFFileSecurityRef) CFFileSecurityRef {
	if _cFFileSecurityCreateCopy == nil {
		panic("CoreFoundation: symbol CFFileSecurityCreateCopy not loaded")
	}
	return _cFFileSecurityCreateCopy(allocator, fileSec)
}


var _cFFileSecurityGetGroup func(fileSec CFFileSecurityRef, group *uint32) bool

// CFFileSecurityGetGroup gets the group ID associated with a [CFFileSecurityRef] object
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetGroup(_:_:)
func CFFileSecurityGetGroup(fileSec CFFileSecurityRef, group *uint32) bool {
	if _cFFileSecurityGetGroup == nil {
		panic("CoreFoundation: symbol CFFileSecurityGetGroup not loaded")
	}
	return _cFFileSecurityGetGroup(fileSec, group)
}


var _cFFileSecurityGetMode func(fileSec CFFileSecurityRef, mode *uint16) bool

// CFFileSecurityGetMode gets the file mode associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetMode(_:_:)
func CFFileSecurityGetMode(fileSec CFFileSecurityRef, mode *uint16) bool {
	if _cFFileSecurityGetMode == nil {
		panic("CoreFoundation: symbol CFFileSecurityGetMode not loaded")
	}
	return _cFFileSecurityGetMode(fileSec, mode)
}


var _cFFileSecurityGetOwner func(fileSec CFFileSecurityRef, owner *uint32) bool

// CFFileSecurityGetOwner gets the owner ID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetOwner(_:_:)
func CFFileSecurityGetOwner(fileSec CFFileSecurityRef, owner *uint32) bool {
	if _cFFileSecurityGetOwner == nil {
		panic("CoreFoundation: symbol CFFileSecurityGetOwner not loaded")
	}
	return _cFFileSecurityGetOwner(fileSec, owner)
}


var _cFFileSecurityGetTypeID func() uint

// CFFileSecurityGetTypeID returns the type identifier for the [CFFileSecurityRef] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetTypeID()
func CFFileSecurityGetTypeID() uint {
	if _cFFileSecurityGetTypeID == nil {
		panic("CoreFoundation: symbol CFFileSecurityGetTypeID not loaded")
	}
	return _cFFileSecurityGetTypeID()
}


var _cFFileSecuritySetAccessControlList func(fileSec CFFileSecurityRef, accessControlList unsafe.Pointer) bool

// CFFileSecuritySetAccessControlList sets the access control list associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetAccessControlList(_:_:)
func CFFileSecuritySetAccessControlList(fileSec CFFileSecurityRef, accessControlList unsafe.Pointer) bool {
	if _cFFileSecuritySetAccessControlList == nil {
		panic("CoreFoundation: symbol CFFileSecuritySetAccessControlList not loaded")
	}
	return _cFFileSecuritySetAccessControlList(fileSec, accessControlList)
}


var _cFFileSecuritySetGroup func(fileSec CFFileSecurityRef, group uint32) bool

// CFFileSecuritySetGroup sets the group ID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroup(_:_:)
func CFFileSecuritySetGroup(fileSec CFFileSecurityRef, group uint32) bool {
	if _cFFileSecuritySetGroup == nil {
		panic("CoreFoundation: symbol CFFileSecuritySetGroup not loaded")
	}
	return _cFFileSecuritySetGroup(fileSec, group)
}


var _cFFileSecuritySetGroupUUID func(fileSec CFFileSecurityRef, groupUUID CFUUIDRef) bool

// CFFileSecuritySetGroupUUID sets the group UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroupUUID(_:_:)
func CFFileSecuritySetGroupUUID(fileSec CFFileSecurityRef, groupUUID CFUUIDRef) bool {
	if _cFFileSecuritySetGroupUUID == nil {
		panic("CoreFoundation: symbol CFFileSecuritySetGroupUUID not loaded")
	}
	return _cFFileSecuritySetGroupUUID(fileSec, groupUUID)
}


var _cFFileSecuritySetMode func(fileSec CFFileSecurityRef, mode uint16) bool

// CFFileSecuritySetMode sets the file mode associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetMode(_:_:)
func CFFileSecuritySetMode(fileSec CFFileSecurityRef, mode uint16) bool {
	if _cFFileSecuritySetMode == nil {
		panic("CoreFoundation: symbol CFFileSecuritySetMode not loaded")
	}
	return _cFFileSecuritySetMode(fileSec, mode)
}


var _cFFileSecuritySetOwner func(fileSec CFFileSecurityRef, owner uint32) bool

// CFFileSecuritySetOwner sets the owner ID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwner(_:_:)
func CFFileSecuritySetOwner(fileSec CFFileSecurityRef, owner uint32) bool {
	if _cFFileSecuritySetOwner == nil {
		panic("CoreFoundation: symbol CFFileSecuritySetOwner not loaded")
	}
	return _cFFileSecuritySetOwner(fileSec, owner)
}


var _cFFileSecuritySetOwnerUUID func(fileSec CFFileSecurityRef, ownerUUID CFUUIDRef) bool

// CFFileSecuritySetOwnerUUID sets the owner UUID associated with a [CFFileSecurityRef] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwnerUUID(_:_:)
func CFFileSecuritySetOwnerUUID(fileSec CFFileSecurityRef, ownerUUID CFUUIDRef) bool {
	if _cFFileSecuritySetOwnerUUID == nil {
		panic("CoreFoundation: symbol CFFileSecuritySetOwnerUUID not loaded")
	}
	return _cFFileSecuritySetOwnerUUID(fileSec, ownerUUID)
}


var _cFGetAllocator func(cf CFTypeRef) CFAllocatorRef

// CFGetAllocator returns the allocator used to allocate a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFGetAllocator(_:)
func CFGetAllocator(cf CFTypeRef) CFAllocatorRef {
	if _cFGetAllocator == nil {
		panic("CoreFoundation: symbol CFGetAllocator not loaded")
	}
	return _cFGetAllocator(cf)
}


var _cFGetRetainCount func(cf CFTypeRef) int

// CFGetRetainCount returns the reference count of a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFGetRetainCount(_:)
func CFGetRetainCount(cf CFTypeRef) int {
	if _cFGetRetainCount == nil {
		panic("CoreFoundation: symbol CFGetRetainCount not loaded")
	}
	return _cFGetRetainCount(cf)
}


var _cFGetTypeID func(cf CFTypeRef) uint

// CFGetTypeID returns the unique identifier of an opaque type to which a Core Foundation object belongs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFGetTypeID(_:)
func CFGetTypeID(cf CFTypeRef) uint {
	if _cFGetTypeID == nil {
		panic("CoreFoundation: symbol CFGetTypeID not loaded")
	}
	return _cFGetTypeID(cf)
}


var _cFHash func(cf CFTypeRef) uint

// CFHash returns a code that can be used to identify an object in a hashing structure.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFHash(_:)
func CFHash(cf CFTypeRef) uint {
	if _cFHash == nil {
		panic("CoreFoundation: symbol CFHash not loaded")
	}
	return _cFHash(cf)
}


var _cFLocaleCopyAvailableLocaleIdentifiers func() CFArrayRef

// CFLocaleCopyAvailableLocaleIdentifiers returns an array of CFString objects that represents all locales for which locale data is available.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyAvailableLocaleIdentifiers()
func CFLocaleCopyAvailableLocaleIdentifiers() CFArrayRef {
	if _cFLocaleCopyAvailableLocaleIdentifiers == nil {
		panic("CoreFoundation: symbol CFLocaleCopyAvailableLocaleIdentifiers not loaded")
	}
	return _cFLocaleCopyAvailableLocaleIdentifiers()
}


var _cFLocaleCopyCommonISOCurrencyCodes func() CFArrayRef

// CFLocaleCopyCommonISOCurrencyCodes returns an array of strings that represents ISO currency codes for currencies in common use.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCommonISOCurrencyCodes()
func CFLocaleCopyCommonISOCurrencyCodes() CFArrayRef {
	if _cFLocaleCopyCommonISOCurrencyCodes == nil {
		panic("CoreFoundation: symbol CFLocaleCopyCommonISOCurrencyCodes not loaded")
	}
	return _cFLocaleCopyCommonISOCurrencyCodes()
}


var _cFLocaleCopyCurrent func() CFLocaleRef

// CFLocaleCopyCurrent returns a copy of the logical locale for the current user.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCurrent()
func CFLocaleCopyCurrent() CFLocaleRef {
	if _cFLocaleCopyCurrent == nil {
		panic("CoreFoundation: symbol CFLocaleCopyCurrent not loaded")
	}
	return _cFLocaleCopyCurrent()
}


var _cFLocaleCopyDisplayNameForPropertyValue func(displayLocale CFLocaleRef, key CFLocaleKey, value CFStringRef) CFStringRef

// CFLocaleCopyDisplayNameForPropertyValue returns the display name for the given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyDisplayNameForPropertyValue(_:_:_:)
func CFLocaleCopyDisplayNameForPropertyValue(displayLocale CFLocaleRef, key CFLocaleKey, value CFStringRef) CFStringRef {
	if _cFLocaleCopyDisplayNameForPropertyValue == nil {
		panic("CoreFoundation: symbol CFLocaleCopyDisplayNameForPropertyValue not loaded")
	}
	return _cFLocaleCopyDisplayNameForPropertyValue(displayLocale, key, value)
}


var _cFLocaleCopyISOCountryCodes func() CFArrayRef

// CFLocaleCopyISOCountryCodes returns an array of CFString objects that represents all known legal ISO country codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCountryCodes()
func CFLocaleCopyISOCountryCodes() CFArrayRef {
	if _cFLocaleCopyISOCountryCodes == nil {
		panic("CoreFoundation: symbol CFLocaleCopyISOCountryCodes not loaded")
	}
	return _cFLocaleCopyISOCountryCodes()
}


var _cFLocaleCopyISOCurrencyCodes func() CFArrayRef

// CFLocaleCopyISOCurrencyCodes returns an array of CFString objects that represents all known legal ISO currency codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCurrencyCodes()
func CFLocaleCopyISOCurrencyCodes() CFArrayRef {
	if _cFLocaleCopyISOCurrencyCodes == nil {
		panic("CoreFoundation: symbol CFLocaleCopyISOCurrencyCodes not loaded")
	}
	return _cFLocaleCopyISOCurrencyCodes()
}


var _cFLocaleCopyISOLanguageCodes func() CFArrayRef

// CFLocaleCopyISOLanguageCodes returns an array of CFString objects that represents all known legal ISO language codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOLanguageCodes()
func CFLocaleCopyISOLanguageCodes() CFArrayRef {
	if _cFLocaleCopyISOLanguageCodes == nil {
		panic("CoreFoundation: symbol CFLocaleCopyISOLanguageCodes not loaded")
	}
	return _cFLocaleCopyISOLanguageCodes()
}


var _cFLocaleCopyPreferredLanguages func() CFArrayRef

// CFLocaleCopyPreferredLanguages returns the array of canonicalized language IDs that the user prefers.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyPreferredLanguages()
func CFLocaleCopyPreferredLanguages() CFArrayRef {
	if _cFLocaleCopyPreferredLanguages == nil {
		panic("CoreFoundation: symbol CFLocaleCopyPreferredLanguages not loaded")
	}
	return _cFLocaleCopyPreferredLanguages()
}


var _cFLocaleCreate func(allocator CFAllocatorRef, localeIdentifier CFLocaleIdentifier) CFLocaleRef

// CFLocaleCreate creates a locale for the given arbitrary locale identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreate(_:_:)
func CFLocaleCreate(allocator CFAllocatorRef, localeIdentifier CFLocaleIdentifier) CFLocaleRef {
	if _cFLocaleCreate == nil {
		panic("CoreFoundation: symbol CFLocaleCreate not loaded")
	}
	return _cFLocaleCreate(allocator, localeIdentifier)
}


var _cFLocaleCreateCanonicalLanguageIdentifierFromString func(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier

// CFLocaleCreateCanonicalLanguageIdentifierFromString returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLanguageIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier {
	if _cFLocaleCreateCanonicalLanguageIdentifierFromString == nil {
		panic("CoreFoundation: symbol CFLocaleCreateCanonicalLanguageIdentifierFromString not loaded")
	}
	return _cFLocaleCreateCanonicalLanguageIdentifierFromString(allocator, localeIdentifier)
}


var _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes func(allocator CFAllocatorRef, lcode unsafe.Pointer, rcode unsafe.Pointer) CFLocaleIdentifier

// CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes returns a canonical locale identifier from given language and region codes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(_:_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator CFAllocatorRef, lcode unsafe.Pointer, rcode unsafe.Pointer) CFLocaleIdentifier {
	if _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes == nil {
		panic("CoreFoundation: symbol CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes not loaded")
	}
	return _cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator, lcode, rcode)
}


var _cFLocaleCreateCanonicalLocaleIdentifierFromString func(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier

// CFLocaleCreateCanonicalLocaleIdentifierFromString returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator CFAllocatorRef, localeIdentifier CFStringRef) CFLocaleIdentifier {
	if _cFLocaleCreateCanonicalLocaleIdentifierFromString == nil {
		panic("CoreFoundation: symbol CFLocaleCreateCanonicalLocaleIdentifierFromString not loaded")
	}
	return _cFLocaleCreateCanonicalLocaleIdentifierFromString(allocator, localeIdentifier)
}


var _cFLocaleCreateComponentsFromLocaleIdentifier func(allocator CFAllocatorRef, localeID CFLocaleIdentifier) CFDictionaryRef

// CFLocaleCreateComponentsFromLocaleIdentifier returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateComponentsFromLocaleIdentifier(_:_:)
func CFLocaleCreateComponentsFromLocaleIdentifier(allocator CFAllocatorRef, localeID CFLocaleIdentifier) CFDictionaryRef {
	if _cFLocaleCreateComponentsFromLocaleIdentifier == nil {
		panic("CoreFoundation: symbol CFLocaleCreateComponentsFromLocaleIdentifier not loaded")
	}
	return _cFLocaleCreateComponentsFromLocaleIdentifier(allocator, localeID)
}


var _cFLocaleCreateCopy func(allocator CFAllocatorRef, locale CFLocaleRef) CFLocaleRef

// CFLocaleCreateCopy returns a copy of a locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCopy(_:_:)
func CFLocaleCreateCopy(allocator CFAllocatorRef, locale CFLocaleRef) CFLocaleRef {
	if _cFLocaleCreateCopy == nil {
		panic("CoreFoundation: symbol CFLocaleCreateCopy not loaded")
	}
	return _cFLocaleCreateCopy(allocator, locale)
}


var _cFLocaleCreateLocaleIdentifierFromComponents func(allocator CFAllocatorRef, dictionary CFDictionaryRef) CFLocaleIdentifier

// CFLocaleCreateLocaleIdentifierFromComponents returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromComponents(_:_:)
func CFLocaleCreateLocaleIdentifierFromComponents(allocator CFAllocatorRef, dictionary CFDictionaryRef) CFLocaleIdentifier {
	if _cFLocaleCreateLocaleIdentifierFromComponents == nil {
		panic("CoreFoundation: symbol CFLocaleCreateLocaleIdentifierFromComponents not loaded")
	}
	return _cFLocaleCreateLocaleIdentifierFromComponents(allocator, dictionary)
}


var _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode func(allocator CFAllocatorRef, lcid uint32) CFLocaleIdentifier

// CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode returns a locale identifier from a Windows locale code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(_:_:)
func CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator CFAllocatorRef, lcid uint32) CFLocaleIdentifier {
	if _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode == nil {
		panic("CoreFoundation: symbol CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode not loaded")
	}
	return _cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator, lcid)
}


var _cFLocaleGetIdentifier func(locale CFLocaleRef) CFLocaleIdentifier

// CFLocaleGetIdentifier returns the given locale’s identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetIdentifier(_:)
func CFLocaleGetIdentifier(locale CFLocaleRef) CFLocaleIdentifier {
	if _cFLocaleGetIdentifier == nil {
		panic("CoreFoundation: symbol CFLocaleGetIdentifier not loaded")
	}
	return _cFLocaleGetIdentifier(locale)
}


var _cFLocaleGetLanguageCharacterDirection func(isoLangCode CFStringRef) CFLocaleLanguageDirection

// CFLocaleGetLanguageCharacterDirection returns the character direction for the specified ISO language code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageCharacterDirection(_:)
func CFLocaleGetLanguageCharacterDirection(isoLangCode CFStringRef) CFLocaleLanguageDirection {
	if _cFLocaleGetLanguageCharacterDirection == nil {
		panic("CoreFoundation: symbol CFLocaleGetLanguageCharacterDirection not loaded")
	}
	return _cFLocaleGetLanguageCharacterDirection(isoLangCode)
}


var _cFLocaleGetLanguageLineDirection func(isoLangCode CFStringRef) CFLocaleLanguageDirection

// CFLocaleGetLanguageLineDirection returns the line direction for the specified ISO language code.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageLineDirection(_:)
func CFLocaleGetLanguageLineDirection(isoLangCode CFStringRef) CFLocaleLanguageDirection {
	if _cFLocaleGetLanguageLineDirection == nil {
		panic("CoreFoundation: symbol CFLocaleGetLanguageLineDirection not loaded")
	}
	return _cFLocaleGetLanguageLineDirection(isoLangCode)
}


var _cFLocaleGetSystem func() CFLocaleRef

// CFLocaleGetSystem returns the root, canonical locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetSystem()
func CFLocaleGetSystem() CFLocaleRef {
	if _cFLocaleGetSystem == nil {
		panic("CoreFoundation: symbol CFLocaleGetSystem not loaded")
	}
	return _cFLocaleGetSystem()
}


var _cFLocaleGetTypeID func() uint

// CFLocaleGetTypeID returns the type identifier for the CFLocale opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetTypeID()
func CFLocaleGetTypeID() uint {
	if _cFLocaleGetTypeID == nil {
		panic("CoreFoundation: symbol CFLocaleGetTypeID not loaded")
	}
	return _cFLocaleGetTypeID()
}


var _cFLocaleGetValue func(locale CFLocaleRef, key CFLocaleKey) CFTypeRef

// CFLocaleGetValue returns the corresponding value for the given key of a locale’s key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetValue(_:_:)
func CFLocaleGetValue(locale CFLocaleRef, key CFLocaleKey) CFTypeRef {
	if _cFLocaleGetValue == nil {
		panic("CoreFoundation: symbol CFLocaleGetValue not loaded")
	}
	return _cFLocaleGetValue(locale, key)
}


var _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier func(localeIdentifier CFLocaleIdentifier) uint32

// CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier returns a Windows locale code from the locale identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(_:)
func CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier CFLocaleIdentifier) uint32 {
	if _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier == nil {
		panic("CoreFoundation: symbol CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier not loaded")
	}
	return _cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
}


var _cFMachPortCreate func(allocator CFAllocatorRef, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef

// CFMachPortCreate creates a CFMachPort object with a new Mach port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreate(_:_:_:_:)
func CFMachPortCreate(allocator CFAllocatorRef, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef {
	if _cFMachPortCreate == nil {
		panic("CoreFoundation: symbol CFMachPortCreate not loaded")
	}
	return _cFMachPortCreate(allocator, callout, context, shouldFreeInfo)
}


var _cFMachPortCreateRunLoopSource func(allocator CFAllocatorRef, port CFMachPortRef, order int) CFRunLoopSourceRef

// CFMachPortCreateRunLoopSource creates a CFRunLoopSource object for a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateRunLoopSource(_:_:_:)
func CFMachPortCreateRunLoopSource(allocator CFAllocatorRef, port CFMachPortRef, order int) CFRunLoopSourceRef {
	if _cFMachPortCreateRunLoopSource == nil {
		panic("CoreFoundation: symbol CFMachPortCreateRunLoopSource not loaded")
	}
	return _cFMachPortCreateRunLoopSource(allocator, port, order)
}


var _cFMachPortCreateWithPort func(allocator CFAllocatorRef, portNum uint32, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef

// CFMachPortCreateWithPort creates a CFMachPort object for a pre-existing native Mach port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateWithPort(_:_:_:_:_:)
func CFMachPortCreateWithPort(allocator CFAllocatorRef, portNum uint32, callout CFMachPortCallBack, context *CFMachPortContext, shouldFreeInfo *bool) CFMachPortRef {
	if _cFMachPortCreateWithPort == nil {
		panic("CoreFoundation: symbol CFMachPortCreateWithPort not loaded")
	}
	return _cFMachPortCreateWithPort(allocator, portNum, callout, context, shouldFreeInfo)
}


var _cFMachPortGetContext func(port CFMachPortRef, context *CFMachPortContext) 

// CFMachPortGetContext returns the context information for a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetContext(_:_:)
func CFMachPortGetContext(port CFMachPortRef, context *CFMachPortContext)  {
	if _cFMachPortGetContext == nil {
		panic("CoreFoundation: symbol CFMachPortGetContext not loaded")
	}
	_cFMachPortGetContext(port, context)
}


var _cFMachPortGetInvalidationCallBack func(port CFMachPortRef) CFMachPortInvalidationCallBack

// CFMachPortGetInvalidationCallBack returns the invalidation callback function for a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetInvalidationCallBack(_:)
func CFMachPortGetInvalidationCallBack(port CFMachPortRef) CFMachPortInvalidationCallBack {
	if _cFMachPortGetInvalidationCallBack == nil {
		panic("CoreFoundation: symbol CFMachPortGetInvalidationCallBack not loaded")
	}
	return _cFMachPortGetInvalidationCallBack(port)
}


var _cFMachPortGetPort func(port CFMachPortRef) uint32

// CFMachPortGetPort returns the native Mach port represented by a CFMachPort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetPort(_:)
func CFMachPortGetPort(port CFMachPortRef) uint32 {
	if _cFMachPortGetPort == nil {
		panic("CoreFoundation: symbol CFMachPortGetPort not loaded")
	}
	return _cFMachPortGetPort(port)
}


var _cFMachPortGetTypeID func() uint

// CFMachPortGetTypeID returns the type identifier for the CFMachPort opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetTypeID()
func CFMachPortGetTypeID() uint {
	if _cFMachPortGetTypeID == nil {
		panic("CoreFoundation: symbol CFMachPortGetTypeID not loaded")
	}
	return _cFMachPortGetTypeID()
}


var _cFMachPortInvalidate func(port CFMachPortRef) 

// CFMachPortInvalidate invalidates a CFMachPort object, stopping it from receiving any more messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortInvalidate(_:)
func CFMachPortInvalidate(port CFMachPortRef)  {
	if _cFMachPortInvalidate == nil {
		panic("CoreFoundation: symbol CFMachPortInvalidate not loaded")
	}
	_cFMachPortInvalidate(port)
}


var _cFMachPortIsValid func(port CFMachPortRef) bool

// CFMachPortIsValid returns a Boolean value that indicates whether a CFMachPort object is valid and able to receive messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortIsValid(_:)
func CFMachPortIsValid(port CFMachPortRef) bool {
	if _cFMachPortIsValid == nil {
		panic("CoreFoundation: symbol CFMachPortIsValid not loaded")
	}
	return _cFMachPortIsValid(port)
}


var _cFMachPortSetInvalidationCallBack func(port CFMachPortRef, callout CFMachPortInvalidationCallBack) 

// CFMachPortSetInvalidationCallBack sets the callback function invoked when a CFMachPort object is invalidated.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMachPortSetInvalidationCallBack(_:_:)
func CFMachPortSetInvalidationCallBack(port CFMachPortRef, callout CFMachPortInvalidationCallBack)  {
	if _cFMachPortSetInvalidationCallBack == nil {
		panic("CoreFoundation: symbol CFMachPortSetInvalidationCallBack not loaded")
	}
	_cFMachPortSetInvalidationCallBack(port, callout)
}


var _cFMakeCollectable func(cf CFTypeRef) CFTypeRef

// CFMakeCollectable makes a newly-allocated Core Foundation object eligible for garbage collection.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMakeCollectable
func CFMakeCollectable(cf CFTypeRef) CFTypeRef {
	if _cFMakeCollectable == nil {
		panic("CoreFoundation: symbol CFMakeCollectable not loaded")
	}
	return _cFMakeCollectable(cf)
}


var _cFMessagePortCreateLocal func(allocator CFAllocatorRef, name CFStringRef, callout CFMessagePortCallBack, context *CFMessagePortContext, shouldFreeInfo *bool) CFMessagePortRef

// CFMessagePortCreateLocal returns a local CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateLocal(_:_:_:_:_:)
func CFMessagePortCreateLocal(allocator CFAllocatorRef, name CFStringRef, callout CFMessagePortCallBack, context *CFMessagePortContext, shouldFreeInfo *bool) CFMessagePortRef {
	if _cFMessagePortCreateLocal == nil {
		panic("CoreFoundation: symbol CFMessagePortCreateLocal not loaded")
	}
	return _cFMessagePortCreateLocal(allocator, name, callout, context, shouldFreeInfo)
}


var _cFMessagePortCreateRemote func(allocator CFAllocatorRef, name CFStringRef) CFMessagePortRef

// CFMessagePortCreateRemote returns a CFMessagePort object connected to a remote port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRemote(_:_:)
func CFMessagePortCreateRemote(allocator CFAllocatorRef, name CFStringRef) CFMessagePortRef {
	if _cFMessagePortCreateRemote == nil {
		panic("CoreFoundation: symbol CFMessagePortCreateRemote not loaded")
	}
	return _cFMessagePortCreateRemote(allocator, name)
}


var _cFMessagePortCreateRunLoopSource func(allocator CFAllocatorRef, local CFMessagePortRef, order int) CFRunLoopSourceRef

// CFMessagePortCreateRunLoopSource creates a CFRunLoopSource object for a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRunLoopSource(_:_:_:)
func CFMessagePortCreateRunLoopSource(allocator CFAllocatorRef, local CFMessagePortRef, order int) CFRunLoopSourceRef {
	if _cFMessagePortCreateRunLoopSource == nil {
		panic("CoreFoundation: symbol CFMessagePortCreateRunLoopSource not loaded")
	}
	return _cFMessagePortCreateRunLoopSource(allocator, local, order)
}


var _cFMessagePortGetContext func(ms CFMessagePortRef, context *CFMessagePortContext) 

// CFMessagePortGetContext returns the context information for a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetContext(_:_:)
func CFMessagePortGetContext(ms CFMessagePortRef, context *CFMessagePortContext)  {
	if _cFMessagePortGetContext == nil {
		panic("CoreFoundation: symbol CFMessagePortGetContext not loaded")
	}
	_cFMessagePortGetContext(ms, context)
}


var _cFMessagePortGetInvalidationCallBack func(ms CFMessagePortRef) CFMessagePortInvalidationCallBack

// CFMessagePortGetInvalidationCallBack returns the invalidation callback function for a CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetInvalidationCallBack(_:)
func CFMessagePortGetInvalidationCallBack(ms CFMessagePortRef) CFMessagePortInvalidationCallBack {
	if _cFMessagePortGetInvalidationCallBack == nil {
		panic("CoreFoundation: symbol CFMessagePortGetInvalidationCallBack not loaded")
	}
	return _cFMessagePortGetInvalidationCallBack(ms)
}


var _cFMessagePortGetName func(ms CFMessagePortRef) CFStringRef

// CFMessagePortGetName returns the name with which a CFMessagePort object is registered.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetName(_:)
func CFMessagePortGetName(ms CFMessagePortRef) CFStringRef {
	if _cFMessagePortGetName == nil {
		panic("CoreFoundation: symbol CFMessagePortGetName not loaded")
	}
	return _cFMessagePortGetName(ms)
}


var _cFMessagePortGetTypeID func() uint

// CFMessagePortGetTypeID returns the type identifier for the CFMessagePort opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetTypeID()
func CFMessagePortGetTypeID() uint {
	if _cFMessagePortGetTypeID == nil {
		panic("CoreFoundation: symbol CFMessagePortGetTypeID not loaded")
	}
	return _cFMessagePortGetTypeID()
}


var _cFMessagePortInvalidate func(ms CFMessagePortRef) 

// CFMessagePortInvalidate invalidates a CFMessagePort object, stopping it from receiving or sending any more messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortInvalidate(_:)
func CFMessagePortInvalidate(ms CFMessagePortRef)  {
	if _cFMessagePortInvalidate == nil {
		panic("CoreFoundation: symbol CFMessagePortInvalidate not loaded")
	}
	_cFMessagePortInvalidate(ms)
}


var _cFMessagePortIsRemote func(ms CFMessagePortRef) bool

// CFMessagePortIsRemote returns a Boolean value that indicates whether a CFMessagePort object represents a remote port.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsRemote(_:)
func CFMessagePortIsRemote(ms CFMessagePortRef) bool {
	if _cFMessagePortIsRemote == nil {
		panic("CoreFoundation: symbol CFMessagePortIsRemote not loaded")
	}
	return _cFMessagePortIsRemote(ms)
}


var _cFMessagePortIsValid func(ms CFMessagePortRef) bool

// CFMessagePortIsValid returns a Boolean value that indicates whether a CFMessagePort object is valid and able to send or receive messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsValid(_:)
func CFMessagePortIsValid(ms CFMessagePortRef) bool {
	if _cFMessagePortIsValid == nil {
		panic("CoreFoundation: symbol CFMessagePortIsValid not loaded")
	}
	return _cFMessagePortIsValid(ms)
}


var _cFMessagePortSendRequest func(remote CFMessagePortRef, msgid int32, data CFDataRef, sendTimeout float64, rcvTimeout float64, replyMode CFStringRef, returnData *CFDataRef) int32

// CFMessagePortSendRequest sends a message to a remote CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSendRequest(_:_:_:_:_:_:_:)
func CFMessagePortSendRequest(remote CFMessagePortRef, msgid int32, data CFDataRef, sendTimeout float64, rcvTimeout float64, replyMode CFStringRef, returnData *CFDataRef) int32 {
	if _cFMessagePortSendRequest == nil {
		panic("CoreFoundation: symbol CFMessagePortSendRequest not loaded")
	}
	return _cFMessagePortSendRequest(remote, msgid, data, sendTimeout, rcvTimeout, replyMode, returnData)
}


var _cFMessagePortSetDispatchQueue func(ms CFMessagePortRef, queue uintptr) 

// CFMessagePortSetDispatchQueue schedules callbacks for the specified message port on the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetDispatchQueue(_:_:)
func CFMessagePortSetDispatchQueue(ms CFMessagePortRef, queue dispatch.Queue)  {
	if _cFMessagePortSetDispatchQueue == nil {
		panic("CoreFoundation: symbol CFMessagePortSetDispatchQueue not loaded")
	}
	_cFMessagePortSetDispatchQueue(ms, uintptr(queue.Handle()))
}


var _cFMessagePortSetInvalidationCallBack func(ms CFMessagePortRef, callout CFMessagePortInvalidationCallBack) 

// CFMessagePortSetInvalidationCallBack sets the callback function invoked when a CFMessagePort object is invalidated.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetInvalidationCallBack(_:_:)
func CFMessagePortSetInvalidationCallBack(ms CFMessagePortRef, callout CFMessagePortInvalidationCallBack)  {
	if _cFMessagePortSetInvalidationCallBack == nil {
		panic("CoreFoundation: symbol CFMessagePortSetInvalidationCallBack not loaded")
	}
	_cFMessagePortSetInvalidationCallBack(ms, callout)
}


var _cFMessagePortSetName func(ms CFMessagePortRef, newName CFStringRef) bool

// CFMessagePortSetName sets the name of a local CFMessagePort object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetName(_:_:)
func CFMessagePortSetName(ms CFMessagePortRef, newName CFStringRef) bool {
	if _cFMessagePortSetName == nil {
		panic("CoreFoundation: symbol CFMessagePortSetName not loaded")
	}
	return _cFMessagePortSetName(ms, newName)
}


var _cFNotificationCenterAddObserver func(center CFNotificationCenterRef, observer unsafe.Pointer, callBack CFNotificationCallback, name CFStringRef, object unsafe.Pointer, suspensionBehavior CFNotificationSuspensionBehavior) 

// CFNotificationCenterAddObserver registers an observer to receive notifications.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterAddObserver(_:_:_:_:_:_:)
func CFNotificationCenterAddObserver(center CFNotificationCenterRef, observer unsafe.Pointer, callBack CFNotificationCallback, name CFStringRef, object unsafe.Pointer, suspensionBehavior CFNotificationSuspensionBehavior)  {
	if _cFNotificationCenterAddObserver == nil {
		panic("CoreFoundation: symbol CFNotificationCenterAddObserver not loaded")
	}
	_cFNotificationCenterAddObserver(center, observer, callBack, name, object, suspensionBehavior)
}


var _cFNotificationCenterGetDarwinNotifyCenter func() CFNotificationCenterRef

// CFNotificationCenterGetDarwinNotifyCenter returns the application’s Darwin notification center.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDarwinNotifyCenter()
func CFNotificationCenterGetDarwinNotifyCenter() CFNotificationCenterRef {
	if _cFNotificationCenterGetDarwinNotifyCenter == nil {
		panic("CoreFoundation: symbol CFNotificationCenterGetDarwinNotifyCenter not loaded")
	}
	return _cFNotificationCenterGetDarwinNotifyCenter()
}


var _cFNotificationCenterGetDistributedCenter func() CFNotificationCenterRef

// CFNotificationCenterGetDistributedCenter returns the application’s distributed notification center.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDistributedCenter()
func CFNotificationCenterGetDistributedCenter() CFNotificationCenterRef {
	if _cFNotificationCenterGetDistributedCenter == nil {
		panic("CoreFoundation: symbol CFNotificationCenterGetDistributedCenter not loaded")
	}
	return _cFNotificationCenterGetDistributedCenter()
}


var _cFNotificationCenterGetLocalCenter func() CFNotificationCenterRef

// CFNotificationCenterGetLocalCenter returns the application’s local notification center.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetLocalCenter()
func CFNotificationCenterGetLocalCenter() CFNotificationCenterRef {
	if _cFNotificationCenterGetLocalCenter == nil {
		panic("CoreFoundation: symbol CFNotificationCenterGetLocalCenter not loaded")
	}
	return _cFNotificationCenterGetLocalCenter()
}


var _cFNotificationCenterGetTypeID func() uint

// CFNotificationCenterGetTypeID returns the type identifier for the CFNotificationCenter opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetTypeID()
func CFNotificationCenterGetTypeID() uint {
	if _cFNotificationCenterGetTypeID == nil {
		panic("CoreFoundation: symbol CFNotificationCenterGetTypeID not loaded")
	}
	return _cFNotificationCenterGetTypeID()
}


var _cFNotificationCenterPostNotification func(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, deliverImmediately bool) 

// CFNotificationCenterPostNotification posts a notification for an object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotification(_:_:_:_:_:)
func CFNotificationCenterPostNotification(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, deliverImmediately bool)  {
	if _cFNotificationCenterPostNotification == nil {
		panic("CoreFoundation: symbol CFNotificationCenterPostNotification not loaded")
	}
	_cFNotificationCenterPostNotification(center, name, object, userInfo, deliverImmediately)
}


var _cFNotificationCenterPostNotificationWithOptions func(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, options uint64) 

// CFNotificationCenterPostNotificationWithOptions posts a notification for an object using specified options.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotificationWithOptions(_:_:_:_:_:)
func CFNotificationCenterPostNotificationWithOptions(center CFNotificationCenterRef, name CFNotificationName, object unsafe.Pointer, userInfo CFDictionaryRef, options uint64)  {
	if _cFNotificationCenterPostNotificationWithOptions == nil {
		panic("CoreFoundation: symbol CFNotificationCenterPostNotificationWithOptions not loaded")
	}
	_cFNotificationCenterPostNotificationWithOptions(center, name, object, userInfo, options)
}


var _cFNotificationCenterRemoveEveryObserver func(center CFNotificationCenterRef, observer unsafe.Pointer) 

// CFNotificationCenterRemoveEveryObserver stops an observer from receiving any notifications from any object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveEveryObserver(_:_:)
func CFNotificationCenterRemoveEveryObserver(center CFNotificationCenterRef, observer unsafe.Pointer)  {
	if _cFNotificationCenterRemoveEveryObserver == nil {
		panic("CoreFoundation: symbol CFNotificationCenterRemoveEveryObserver not loaded")
	}
	_cFNotificationCenterRemoveEveryObserver(center, observer)
}


var _cFNotificationCenterRemoveObserver func(center CFNotificationCenterRef, observer unsafe.Pointer, name CFNotificationName, object unsafe.Pointer) 

// CFNotificationCenterRemoveObserver stops an observer from receiving certain notifications.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveObserver(_:_:_:_:)
func CFNotificationCenterRemoveObserver(center CFNotificationCenterRef, observer unsafe.Pointer, name CFNotificationName, object unsafe.Pointer)  {
	if _cFNotificationCenterRemoveObserver == nil {
		panic("CoreFoundation: symbol CFNotificationCenterRemoveObserver not loaded")
	}
	_cFNotificationCenterRemoveObserver(center, observer, name, object)
}


var _cFNullGetTypeID func() uint

// CFNullGetTypeID returns the type identifier for the CFNull opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNullGetTypeID()
func CFNullGetTypeID() uint {
	if _cFNullGetTypeID == nil {
		panic("CoreFoundation: symbol CFNullGetTypeID not loaded")
	}
	return _cFNullGetTypeID()
}


var _cFNumberCompare func(number CFNumberRef, otherNumber CFNumberRef, context unsafe.Pointer) CFComparisonResult

// CFNumberCompare compares two CFNumber objects and returns a comparison result.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberCompare(_:_:_:)
func CFNumberCompare(number CFNumberRef, otherNumber CFNumberRef, context unsafe.Pointer) CFComparisonResult {
	if _cFNumberCompare == nil {
		panic("CoreFoundation: symbol CFNumberCompare not loaded")
	}
	return _cFNumberCompare(number, otherNumber, context)
}


var _cFNumberCreate func(allocator CFAllocatorRef, theType CFNumberType, valuePtr unsafe.Pointer) CFNumberRef

// CFNumberCreate creates a CFNumber object using a specified value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberCreate(_:_:_:)
func CFNumberCreate(allocator CFAllocatorRef, theType CFNumberType, valuePtr unsafe.Pointer) CFNumberRef {
	if _cFNumberCreate == nil {
		panic("CoreFoundation: symbol CFNumberCreate not loaded")
	}
	return _cFNumberCreate(allocator, theType, valuePtr)
}


var _cFNumberFormatterCopyProperty func(formatter CFNumberFormatterRef, key CFNumberFormatterKey) CFTypeRef

// CFNumberFormatterCopyProperty returns a copy of a number formatter’s value for a given key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCopyProperty(_:_:)
func CFNumberFormatterCopyProperty(formatter CFNumberFormatterRef, key CFNumberFormatterKey) CFTypeRef {
	if _cFNumberFormatterCopyProperty == nil {
		panic("CoreFoundation: symbol CFNumberFormatterCopyProperty not loaded")
	}
	return _cFNumberFormatterCopyProperty(formatter, key)
}


var _cFNumberFormatterCreate func(allocator CFAllocatorRef, locale CFLocaleRef, style CFNumberFormatterStyle) CFNumberFormatterRef

// CFNumberFormatterCreate creates a new CFNumberFormatter object, localized to the given locale, which will format numbers to the given style.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreate(_:_:_:)
func CFNumberFormatterCreate(allocator CFAllocatorRef, locale CFLocaleRef, style CFNumberFormatterStyle) CFNumberFormatterRef {
	if _cFNumberFormatterCreate == nil {
		panic("CoreFoundation: symbol CFNumberFormatterCreate not loaded")
	}
	return _cFNumberFormatterCreate(allocator, locale, style)
}


var _cFNumberFormatterCreateNumberFromString func(allocator CFAllocatorRef, formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, options uint64) CFNumberRef

// CFNumberFormatterCreateNumberFromString returns a number object representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateNumberFromString(_:_:_:_:_:)
func CFNumberFormatterCreateNumberFromString(allocator CFAllocatorRef, formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, options uint64) CFNumberRef {
	if _cFNumberFormatterCreateNumberFromString == nil {
		panic("CoreFoundation: symbol CFNumberFormatterCreateNumberFromString not loaded")
	}
	return _cFNumberFormatterCreateNumberFromString(allocator, formatter, string_, rangep, options)
}


var _cFNumberFormatterCreateStringWithNumber func(allocator CFAllocatorRef, formatter CFNumberFormatterRef, number CFNumberRef) CFStringRef

// CFNumberFormatterCreateStringWithNumber returns a string representation of the given number using the specified number formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithNumber(_:_:_:)
func CFNumberFormatterCreateStringWithNumber(allocator CFAllocatorRef, formatter CFNumberFormatterRef, number CFNumberRef) CFStringRef {
	if _cFNumberFormatterCreateStringWithNumber == nil {
		panic("CoreFoundation: symbol CFNumberFormatterCreateStringWithNumber not loaded")
	}
	return _cFNumberFormatterCreateStringWithNumber(allocator, formatter, number)
}


var _cFNumberFormatterCreateStringWithValue func(allocator CFAllocatorRef, formatter CFNumberFormatterRef, numberType CFNumberType, valuePtr unsafe.Pointer) CFStringRef

// CFNumberFormatterCreateStringWithValue returns a string representation of the given number or value using the specified number formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithValue(_:_:_:_:)
func CFNumberFormatterCreateStringWithValue(allocator CFAllocatorRef, formatter CFNumberFormatterRef, numberType CFNumberType, valuePtr unsafe.Pointer) CFStringRef {
	if _cFNumberFormatterCreateStringWithValue == nil {
		panic("CoreFoundation: symbol CFNumberFormatterCreateStringWithValue not loaded")
	}
	return _cFNumberFormatterCreateStringWithValue(allocator, formatter, numberType, valuePtr)
}


var _cFNumberFormatterGetDecimalInfoForCurrencyCode func(currencyCode CFStringRef, defaultFractionDigits *int32, roundingIncrement *float64) bool

// CFNumberFormatterGetDecimalInfoForCurrencyCode returns the number of fraction digits that should be displayed, and the rounding increment, for a given currency.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetDecimalInfoForCurrencyCode(_:_:_:)
func CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode CFStringRef, defaultFractionDigits *int32, roundingIncrement []float64) bool {
	if _cFNumberFormatterGetDecimalInfoForCurrencyCode == nil {
		panic("CoreFoundation: symbol CFNumberFormatterGetDecimalInfoForCurrencyCode not loaded")
	}
	return _cFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode, defaultFractionDigits, unsafe.SliceData(roundingIncrement))
}


var _cFNumberFormatterGetFormat func(formatter CFNumberFormatterRef) CFStringRef

// CFNumberFormatterGetFormat returns a format string for the given number formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetFormat(_:)
func CFNumberFormatterGetFormat(formatter CFNumberFormatterRef) CFStringRef {
	if _cFNumberFormatterGetFormat == nil {
		panic("CoreFoundation: symbol CFNumberFormatterGetFormat not loaded")
	}
	return _cFNumberFormatterGetFormat(formatter)
}


var _cFNumberFormatterGetLocale func(formatter CFNumberFormatterRef) CFLocaleRef

// CFNumberFormatterGetLocale returns the locale object used to create the given number formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetLocale(_:)
func CFNumberFormatterGetLocale(formatter CFNumberFormatterRef) CFLocaleRef {
	if _cFNumberFormatterGetLocale == nil {
		panic("CoreFoundation: symbol CFNumberFormatterGetLocale not loaded")
	}
	return _cFNumberFormatterGetLocale(formatter)
}


var _cFNumberFormatterGetStyle func(formatter CFNumberFormatterRef) CFNumberFormatterStyle

// CFNumberFormatterGetStyle returns the number style used to create the given number formatter object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetStyle(_:)
func CFNumberFormatterGetStyle(formatter CFNumberFormatterRef) CFNumberFormatterStyle {
	if _cFNumberFormatterGetStyle == nil {
		panic("CoreFoundation: symbol CFNumberFormatterGetStyle not loaded")
	}
	return _cFNumberFormatterGetStyle(formatter)
}


var _cFNumberFormatterGetTypeID func() uint

// CFNumberFormatterGetTypeID returns the type identifier for the [CFNumberFormatter] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetTypeID()
func CFNumberFormatterGetTypeID() uint {
	if _cFNumberFormatterGetTypeID == nil {
		panic("CoreFoundation: symbol CFNumberFormatterGetTypeID not loaded")
	}
	return _cFNumberFormatterGetTypeID()
}


var _cFNumberFormatterGetValueFromString func(formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, numberType CFNumberType, valuePtr unsafe.Pointer) bool

// CFNumberFormatterGetValueFromString returns a number or value representing a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetValueFromString(_:_:_:_:_:)
func CFNumberFormatterGetValueFromString(formatter CFNumberFormatterRef, string_ CFStringRef, rangep *CFRange, numberType CFNumberType, valuePtr unsafe.Pointer) bool {
	if _cFNumberFormatterGetValueFromString == nil {
		panic("CoreFoundation: symbol CFNumberFormatterGetValueFromString not loaded")
	}
	return _cFNumberFormatterGetValueFromString(formatter, string_, rangep, numberType, valuePtr)
}


var _cFNumberFormatterSetFormat func(formatter CFNumberFormatterRef, formatString CFStringRef) 

// CFNumberFormatterSetFormat sets the format string of a number formatter.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetFormat(_:_:)
func CFNumberFormatterSetFormat(formatter CFNumberFormatterRef, formatString CFStringRef)  {
	if _cFNumberFormatterSetFormat == nil {
		panic("CoreFoundation: symbol CFNumberFormatterSetFormat not loaded")
	}
	_cFNumberFormatterSetFormat(formatter, formatString)
}


var _cFNumberFormatterSetProperty func(formatter CFNumberFormatterRef, key CFNumberFormatterKey, value CFTypeRef) 

// CFNumberFormatterSetProperty sets a number formatter property using a key-value pair.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetProperty(_:_:_:)
func CFNumberFormatterSetProperty(formatter CFNumberFormatterRef, key CFNumberFormatterKey, value CFTypeRef)  {
	if _cFNumberFormatterSetProperty == nil {
		panic("CoreFoundation: symbol CFNumberFormatterSetProperty not loaded")
	}
	_cFNumberFormatterSetProperty(formatter, key, value)
}


var _cFNumberGetByteSize func(number CFNumberRef) int

// CFNumberGetByteSize returns the number of bytes used by a CFNumber object to store its value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetByteSize(_:)
func CFNumberGetByteSize(number CFNumberRef) int {
	if _cFNumberGetByteSize == nil {
		panic("CoreFoundation: symbol CFNumberGetByteSize not loaded")
	}
	return _cFNumberGetByteSize(number)
}


var _cFNumberGetType func(number CFNumberRef) CFNumberType

// CFNumberGetType returns the type used by a CFNumber object to store its value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetType(_:)
func CFNumberGetType(number CFNumberRef) CFNumberType {
	if _cFNumberGetType == nil {
		panic("CoreFoundation: symbol CFNumberGetType not loaded")
	}
	return _cFNumberGetType(number)
}


var _cFNumberGetTypeID func() uint

// CFNumberGetTypeID returns the type identifier for the CFNumber opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetTypeID()
func CFNumberGetTypeID() uint {
	if _cFNumberGetTypeID == nil {
		panic("CoreFoundation: symbol CFNumberGetTypeID not loaded")
	}
	return _cFNumberGetTypeID()
}


var _cFNumberGetValue func(number CFNumberRef, theType CFNumberType, valuePtr unsafe.Pointer) bool

// CFNumberGetValue obtains the value of a CFNumber object cast to a specified type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetValue(_:_:_:)
func CFNumberGetValue(number CFNumberRef, theType CFNumberType, valuePtr unsafe.Pointer) bool {
	if _cFNumberGetValue == nil {
		panic("CoreFoundation: symbol CFNumberGetValue not loaded")
	}
	return _cFNumberGetValue(number, theType, valuePtr)
}


var _cFNumberIsFloatType func(number CFNumberRef) bool

// CFNumberIsFloatType determines whether a CFNumber object contains a value stored as one of the defined floating point types.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberIsFloatType(_:)
func CFNumberIsFloatType(number CFNumberRef) bool {
	if _cFNumberIsFloatType == nil {
		panic("CoreFoundation: symbol CFNumberIsFloatType not loaded")
	}
	return _cFNumberIsFloatType(number)
}


var _cFPlugInAddInstanceForFactory func(factoryID CFUUIDRef) 

// CFPlugInAddInstanceForFactory registers a new instance of a type with [CFPlugIn].
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInAddInstanceForFactory(_:)
func CFPlugInAddInstanceForFactory(factoryID CFUUIDRef)  {
	if _cFPlugInAddInstanceForFactory == nil {
		panic("CoreFoundation: symbol CFPlugInAddInstanceForFactory not loaded")
	}
	_cFPlugInAddInstanceForFactory(factoryID)
}


var _cFPlugInCreate func(allocator CFAllocatorRef, plugInURL CFURLRef) CFPlugInRef

// CFPlugInCreate creates a CFPlugIn given its URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInCreate(_:_:)
func CFPlugInCreate(allocator CFAllocatorRef, plugInURL CFURLRef) CFPlugInRef {
	if _cFPlugInCreate == nil {
		panic("CoreFoundation: symbol CFPlugInCreate not loaded")
	}
	return _cFPlugInCreate(allocator, plugInURL)
}


var _cFPlugInFindFactoriesForPlugInType func(typeUUID CFUUIDRef) CFArrayRef

// CFPlugInFindFactoriesForPlugInType searches all registered plug-ins for factory functions capable of creating an instance of the given type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInType(_:)
func CFPlugInFindFactoriesForPlugInType(typeUUID CFUUIDRef) CFArrayRef {
	if _cFPlugInFindFactoriesForPlugInType == nil {
		panic("CoreFoundation: symbol CFPlugInFindFactoriesForPlugInType not loaded")
	}
	return _cFPlugInFindFactoriesForPlugInType(typeUUID)
}


var _cFPlugInFindFactoriesForPlugInTypeInPlugIn func(typeUUID CFUUIDRef, plugIn CFPlugInRef) CFArrayRef

// CFPlugInFindFactoriesForPlugInTypeInPlugIn searches the given plug-in for factory functions capable of creating an instance of the given type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInTypeInPlugIn(_:_:)
func CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID CFUUIDRef, plugIn CFPlugInRef) CFArrayRef {
	if _cFPlugInFindFactoriesForPlugInTypeInPlugIn == nil {
		panic("CoreFoundation: symbol CFPlugInFindFactoriesForPlugInTypeInPlugIn not loaded")
	}
	return _cFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID, plugIn)
}


var _cFPlugInGetBundle func(plugIn CFPlugInRef) CFBundleRef

// CFPlugInGetBundle returns a plug-in’s bundle.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetBundle(_:)
func CFPlugInGetBundle(plugIn CFPlugInRef) CFBundleRef {
	if _cFPlugInGetBundle == nil {
		panic("CoreFoundation: symbol CFPlugInGetBundle not loaded")
	}
	return _cFPlugInGetBundle(plugIn)
}


var _cFPlugInGetTypeID func() uint

// CFPlugInGetTypeID returns the type identifier for the [CFPlugIn] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetTypeID()
func CFPlugInGetTypeID() uint {
	if _cFPlugInGetTypeID == nil {
		panic("CoreFoundation: symbol CFPlugInGetTypeID not loaded")
	}
	return _cFPlugInGetTypeID()
}


var _cFPlugInInstanceCreate func(allocator CFAllocatorRef, factoryUUID CFUUIDRef, typeUUID CFUUIDRef) unsafe.Pointer

// CFPlugInInstanceCreate creates a [CFPlugIn] instance of a given type using a given factory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreate(_:_:_:)
func CFPlugInInstanceCreate(allocator CFAllocatorRef, factoryUUID CFUUIDRef, typeUUID CFUUIDRef) unsafe.Pointer {
	if _cFPlugInInstanceCreate == nil {
		panic("CoreFoundation: symbol CFPlugInInstanceCreate not loaded")
	}
	return _cFPlugInInstanceCreate(allocator, factoryUUID, typeUUID)
}


var _cFPlugInInstanceCreateWithInstanceDataSize func(allocator CFAllocatorRef, instanceDataSize int, deallocateInstanceFunction CFPlugInInstanceDeallocateInstanceDataFunction, factoryName CFStringRef, getInterfaceFunction CFPlugInInstanceGetInterfaceFunction) CFPlugInInstanceRef

// CFPlugInInstanceCreateWithInstanceDataSize not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreateWithInstanceDataSize(_:_:_:_:_:)
func CFPlugInInstanceCreateWithInstanceDataSize(allocator CFAllocatorRef, instanceDataSize int, deallocateInstanceFunction CFPlugInInstanceDeallocateInstanceDataFunction, factoryName CFStringRef, getInterfaceFunction CFPlugInInstanceGetInterfaceFunction) CFPlugInInstanceRef {
	if _cFPlugInInstanceCreateWithInstanceDataSize == nil {
		panic("CoreFoundation: symbol CFPlugInInstanceCreateWithInstanceDataSize not loaded")
	}
	return _cFPlugInInstanceCreateWithInstanceDataSize(allocator, instanceDataSize, deallocateInstanceFunction, factoryName, getInterfaceFunction)
}


var _cFPlugInInstanceGetFactoryName func(instance CFPlugInInstanceRef) CFStringRef

// CFPlugInInstanceGetFactoryName not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetFactoryName(_:)
func CFPlugInInstanceGetFactoryName(instance CFPlugInInstanceRef) CFStringRef {
	if _cFPlugInInstanceGetFactoryName == nil {
		panic("CoreFoundation: symbol CFPlugInInstanceGetFactoryName not loaded")
	}
	return _cFPlugInInstanceGetFactoryName(instance)
}


var _cFPlugInInstanceGetInstanceData func(instance CFPlugInInstanceRef) unsafe.Pointer

// CFPlugInInstanceGetInstanceData not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInstanceData(_:)
func CFPlugInInstanceGetInstanceData(instance CFPlugInInstanceRef) unsafe.Pointer {
	if _cFPlugInInstanceGetInstanceData == nil {
		panic("CoreFoundation: symbol CFPlugInInstanceGetInstanceData not loaded")
	}
	return _cFPlugInInstanceGetInstanceData(instance)
}


var _cFPlugInInstanceGetInterfaceFunctionTable func(instance CFPlugInInstanceRef, interfaceName CFStringRef, ftbl unsafe.Pointer) bool

// CFPlugInInstanceGetInterfaceFunctionTable not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInterfaceFunctionTable(_:_:_:)
func CFPlugInInstanceGetInterfaceFunctionTable(instance CFPlugInInstanceRef, interfaceName CFStringRef, ftbl unsafe.Pointer) bool {
	if _cFPlugInInstanceGetInterfaceFunctionTable == nil {
		panic("CoreFoundation: symbol CFPlugInInstanceGetInterfaceFunctionTable not loaded")
	}
	return _cFPlugInInstanceGetInterfaceFunctionTable(instance, interfaceName, ftbl)
}


var _cFPlugInInstanceGetTypeID func() uint

// CFPlugInInstanceGetTypeID not recommended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetTypeID()
func CFPlugInInstanceGetTypeID() uint {
	if _cFPlugInInstanceGetTypeID == nil {
		panic("CoreFoundation: symbol CFPlugInInstanceGetTypeID not loaded")
	}
	return _cFPlugInInstanceGetTypeID()
}


var _cFPlugInIsLoadOnDemand func(plugIn CFPlugInRef) bool

// CFPlugInIsLoadOnDemand determines whether or not a plug-in is loaded on demand.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInIsLoadOnDemand(_:)
func CFPlugInIsLoadOnDemand(plugIn CFPlugInRef) bool {
	if _cFPlugInIsLoadOnDemand == nil {
		panic("CoreFoundation: symbol CFPlugInIsLoadOnDemand not loaded")
	}
	return _cFPlugInIsLoadOnDemand(plugIn)
}


var _cFPlugInRegisterFactoryFunction func(factoryUUID CFUUIDRef, func_ CFPlugInFactoryFunction) bool

// CFPlugInRegisterFactoryFunction registers a factory function and its UUID with a [CFPlugIn] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunction(_:_:)
func CFPlugInRegisterFactoryFunction(factoryUUID CFUUIDRef, func_ CFPlugInFactoryFunction) bool {
	if _cFPlugInRegisterFactoryFunction == nil {
		panic("CoreFoundation: symbol CFPlugInRegisterFactoryFunction not loaded")
	}
	return _cFPlugInRegisterFactoryFunction(factoryUUID, func_)
}


var _cFPlugInRegisterFactoryFunctionByName func(factoryUUID CFUUIDRef, plugIn CFPlugInRef, functionName CFStringRef) bool

// CFPlugInRegisterFactoryFunctionByName registers a factory function with a [CFPlugIn] object using the function’s name instead of its UUID.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunctionByName(_:_:_:)
func CFPlugInRegisterFactoryFunctionByName(factoryUUID CFUUIDRef, plugIn CFPlugInRef, functionName CFStringRef) bool {
	if _cFPlugInRegisterFactoryFunctionByName == nil {
		panic("CoreFoundation: symbol CFPlugInRegisterFactoryFunctionByName not loaded")
	}
	return _cFPlugInRegisterFactoryFunctionByName(factoryUUID, plugIn, functionName)
}


var _cFPlugInRegisterPlugInType func(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool

// CFPlugInRegisterPlugInType registers a type and its corresponding factory function with a [CFPlugIn] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterPlugInType(_:_:)
func CFPlugInRegisterPlugInType(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool {
	if _cFPlugInRegisterPlugInType == nil {
		panic("CoreFoundation: symbol CFPlugInRegisterPlugInType not loaded")
	}
	return _cFPlugInRegisterPlugInType(factoryUUID, typeUUID)
}


var _cFPlugInRemoveInstanceForFactory func(factoryID CFUUIDRef) 

// CFPlugInRemoveInstanceForFactory unregisters an instance of a type with [CFPlugIn].
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRemoveInstanceForFactory(_:)
func CFPlugInRemoveInstanceForFactory(factoryID CFUUIDRef)  {
	if _cFPlugInRemoveInstanceForFactory == nil {
		panic("CoreFoundation: symbol CFPlugInRemoveInstanceForFactory not loaded")
	}
	_cFPlugInRemoveInstanceForFactory(factoryID)
}


var _cFPlugInSetLoadOnDemand func(plugIn CFPlugInRef, flag bool) 

// CFPlugInSetLoadOnDemand enables or disables load on demand for plug-ins that do dynamic registration (only when a client requests an instance of a supported type).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInSetLoadOnDemand(_:_:)
func CFPlugInSetLoadOnDemand(plugIn CFPlugInRef, flag bool)  {
	if _cFPlugInSetLoadOnDemand == nil {
		panic("CoreFoundation: symbol CFPlugInSetLoadOnDemand not loaded")
	}
	_cFPlugInSetLoadOnDemand(plugIn, flag)
}


var _cFPlugInUnregisterFactory func(factoryUUID CFUUIDRef) bool

// CFPlugInUnregisterFactory removes the given function from a plug-in’s list of registered factory functions.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterFactory(_:)
func CFPlugInUnregisterFactory(factoryUUID CFUUIDRef) bool {
	if _cFPlugInUnregisterFactory == nil {
		panic("CoreFoundation: symbol CFPlugInUnregisterFactory not loaded")
	}
	return _cFPlugInUnregisterFactory(factoryUUID)
}


var _cFPlugInUnregisterPlugInType func(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool

// CFPlugInUnregisterPlugInType removes the given type from a plug-in’s list of registered types.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterPlugInType(_:_:)
func CFPlugInUnregisterPlugInType(factoryUUID CFUUIDRef, typeUUID CFUUIDRef) bool {
	if _cFPlugInUnregisterPlugInType == nil {
		panic("CoreFoundation: symbol CFPlugInUnregisterPlugInType not loaded")
	}
	return _cFPlugInUnregisterPlugInType(factoryUUID, typeUUID)
}


var _cFPreferencesAddSuitePreferencesToApp func(applicationID CFStringRef, suiteID CFStringRef) 

// CFPreferencesAddSuitePreferencesToApp adds suite preferences to an application’s preference search chain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAddSuitePreferencesToApp(_:_:)
func CFPreferencesAddSuitePreferencesToApp(applicationID CFStringRef, suiteID CFStringRef)  {
	if _cFPreferencesAddSuitePreferencesToApp == nil {
		panic("CoreFoundation: symbol CFPreferencesAddSuitePreferencesToApp not loaded")
	}
	_cFPreferencesAddSuitePreferencesToApp(applicationID, suiteID)
}


var _cFPreferencesAppSynchronize func(applicationID CFStringRef) bool

// CFPreferencesAppSynchronize writes to permanent storage all pending changes to the preference data for the application, and reads the latest preference data from permanent storage.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppSynchronize(_:)
func CFPreferencesAppSynchronize(applicationID CFStringRef) bool {
	if _cFPreferencesAppSynchronize == nil {
		panic("CoreFoundation: symbol CFPreferencesAppSynchronize not loaded")
	}
	return _cFPreferencesAppSynchronize(applicationID)
}


var _cFPreferencesAppValueIsForced func(key CFStringRef, applicationID CFStringRef) bool

// CFPreferencesAppValueIsForced determines whether or not a given key has been imposed on the user.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppValueIsForced(_:_:)
func CFPreferencesAppValueIsForced(key CFStringRef, applicationID CFStringRef) bool {
	if _cFPreferencesAppValueIsForced == nil {
		panic("CoreFoundation: symbol CFPreferencesAppValueIsForced not loaded")
	}
	return _cFPreferencesAppValueIsForced(key, applicationID)
}


var _cFPreferencesCopyAppValue func(key CFStringRef, applicationID CFStringRef) CFPropertyListRef

// CFPreferencesCopyAppValue obtains a preference value for the specified key and application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyAppValue(_:_:)
func CFPreferencesCopyAppValue(key CFStringRef, applicationID CFStringRef) CFPropertyListRef {
	if _cFPreferencesCopyAppValue == nil {
		panic("CoreFoundation: symbol CFPreferencesCopyAppValue not loaded")
	}
	return _cFPreferencesCopyAppValue(key, applicationID)
}


var _cFPreferencesCopyKeyList func(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFArrayRef

// CFPreferencesCopyKeyList constructs and returns the list of all keys set in the specified domain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyKeyList(_:_:_:)
func CFPreferencesCopyKeyList(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFArrayRef {
	if _cFPreferencesCopyKeyList == nil {
		panic("CoreFoundation: symbol CFPreferencesCopyKeyList not loaded")
	}
	return _cFPreferencesCopyKeyList(applicationID, userName, hostName)
}


var _cFPreferencesCopyMultiple func(keysToFetch CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFDictionaryRef

// CFPreferencesCopyMultiple returns a dictionary containing preference values for multiple keys.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyMultiple(_:_:_:_:)
func CFPreferencesCopyMultiple(keysToFetch CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFDictionaryRef {
	if _cFPreferencesCopyMultiple == nil {
		panic("CoreFoundation: symbol CFPreferencesCopyMultiple not loaded")
	}
	return _cFPreferencesCopyMultiple(keysToFetch, applicationID, userName, hostName)
}


var _cFPreferencesCopyValue func(key CFStringRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFPropertyListRef

// CFPreferencesCopyValue returns a preference value for a given domain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyValue(_:_:_:_:)
func CFPreferencesCopyValue(key CFStringRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) CFPropertyListRef {
	if _cFPreferencesCopyValue == nil {
		panic("CoreFoundation: symbol CFPreferencesCopyValue not loaded")
	}
	return _cFPreferencesCopyValue(key, applicationID, userName, hostName)
}


var _cFPreferencesGetAppBooleanValue func(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) bool

// CFPreferencesGetAppBooleanValue convenience function that directly obtains a Boolean preference value for the specified key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppBooleanValue(_:_:_:)
func CFPreferencesGetAppBooleanValue(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) bool {
	if _cFPreferencesGetAppBooleanValue == nil {
		panic("CoreFoundation: symbol CFPreferencesGetAppBooleanValue not loaded")
	}
	return _cFPreferencesGetAppBooleanValue(key, applicationID, keyExistsAndHasValidFormat)
}


var _cFPreferencesGetAppIntegerValue func(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) int

// CFPreferencesGetAppIntegerValue convenience function that directly obtains an integer preference value for the specified key.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppIntegerValue(_:_:_:)
func CFPreferencesGetAppIntegerValue(key CFStringRef, applicationID CFStringRef, keyExistsAndHasValidFormat *bool) int {
	if _cFPreferencesGetAppIntegerValue == nil {
		panic("CoreFoundation: symbol CFPreferencesGetAppIntegerValue not loaded")
	}
	return _cFPreferencesGetAppIntegerValue(key, applicationID, keyExistsAndHasValidFormat)
}


var _cFPreferencesRemoveSuitePreferencesFromApp func(applicationID CFStringRef, suiteID CFStringRef) 

// CFPreferencesRemoveSuitePreferencesFromApp removes suite preferences from an application’s search chain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesRemoveSuitePreferencesFromApp(_:_:)
func CFPreferencesRemoveSuitePreferencesFromApp(applicationID CFStringRef, suiteID CFStringRef)  {
	if _cFPreferencesRemoveSuitePreferencesFromApp == nil {
		panic("CoreFoundation: symbol CFPreferencesRemoveSuitePreferencesFromApp not loaded")
	}
	_cFPreferencesRemoveSuitePreferencesFromApp(applicationID, suiteID)
}


var _cFPreferencesSetAppValue func(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef) 

// CFPreferencesSetAppValue adds, modifies, or removes a preference.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetAppValue(_:_:_:)
func CFPreferencesSetAppValue(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef)  {
	if _cFPreferencesSetAppValue == nil {
		panic("CoreFoundation: symbol CFPreferencesSetAppValue not loaded")
	}
	_cFPreferencesSetAppValue(key, value, applicationID)
}


var _cFPreferencesSetMultiple func(keysToSet CFDictionaryRef, keysToRemove CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) 

// CFPreferencesSetMultiple convenience function that allows you to set and remove multiple preference values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetMultiple(_:_:_:_:_:)
func CFPreferencesSetMultiple(keysToSet CFDictionaryRef, keysToRemove CFArrayRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef)  {
	if _cFPreferencesSetMultiple == nil {
		panic("CoreFoundation: symbol CFPreferencesSetMultiple not loaded")
	}
	_cFPreferencesSetMultiple(keysToSet, keysToRemove, applicationID, userName, hostName)
}


var _cFPreferencesSetValue func(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) 

// CFPreferencesSetValue adds, modifies, or removes a preference value for the specified domain.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetValue(_:_:_:_:_:)
func CFPreferencesSetValue(key CFStringRef, value CFPropertyListRef, applicationID CFStringRef, userName CFStringRef, hostName CFStringRef)  {
	if _cFPreferencesSetValue == nil {
		panic("CoreFoundation: symbol CFPreferencesSetValue not loaded")
	}
	_cFPreferencesSetValue(key, value, applicationID, userName, hostName)
}


var _cFPreferencesSynchronize func(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) bool

// CFPreferencesSynchronize for the specified domain, writes all pending changes to preference data to permanent storage, and reads latest preference data from permanent storage.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSynchronize(_:_:_:)
func CFPreferencesSynchronize(applicationID CFStringRef, userName CFStringRef, hostName CFStringRef) bool {
	if _cFPreferencesSynchronize == nil {
		panic("CoreFoundation: symbol CFPreferencesSynchronize not loaded")
	}
	return _cFPreferencesSynchronize(applicationID, userName, hostName)
}


var _cFPropertyListCreateData func(allocator CFAllocatorRef, propertyList CFPropertyListRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) CFDataRef

// CFPropertyListCreateData returns a CFData object containing a serialized representation of a given property list in a specified format.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateData(_:_:_:_:_:)
func CFPropertyListCreateData(allocator CFAllocatorRef, propertyList CFPropertyListRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) CFDataRef {
	if _cFPropertyListCreateData == nil {
		panic("CoreFoundation: symbol CFPropertyListCreateData not loaded")
	}
	return _cFPropertyListCreateData(allocator, propertyList, format, options, err)
}


var _cFPropertyListCreateDeepCopy func(allocator CFAllocatorRef, propertyList CFPropertyListRef, mutabilityOption uint64) CFPropertyListRef

// CFPropertyListCreateDeepCopy recursively creates a copy of a given property list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateDeepCopy(_:_:_:)
func CFPropertyListCreateDeepCopy(allocator CFAllocatorRef, propertyList CFPropertyListRef, mutabilityOption uint64) CFPropertyListRef {
	if _cFPropertyListCreateDeepCopy == nil {
		panic("CoreFoundation: symbol CFPropertyListCreateDeepCopy not loaded")
	}
	return _cFPropertyListCreateDeepCopy(allocator, propertyList, mutabilityOption)
}


var _cFPropertyListCreateWithData func(allocator CFAllocatorRef, data CFDataRef, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef

// CFPropertyListCreateWithData creates a property list from a given CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithData(_:_:_:_:_:)
func CFPropertyListCreateWithData(allocator CFAllocatorRef, data CFDataRef, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef {
	if _cFPropertyListCreateWithData == nil {
		panic("CoreFoundation: symbol CFPropertyListCreateWithData not loaded")
	}
	return _cFPropertyListCreateWithData(allocator, data, options, format, err)
}


var _cFPropertyListCreateWithStream func(allocator CFAllocatorRef, stream CFReadStreamRef, streamLength int, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef

// CFPropertyListCreateWithStream create and return a property list with a CFReadStream input.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithStream(_:_:_:_:_:_:)
func CFPropertyListCreateWithStream(allocator CFAllocatorRef, stream CFReadStreamRef, streamLength int, options uint64, format *CFPropertyListFormat, err *CFErrorRef) CFPropertyListRef {
	if _cFPropertyListCreateWithStream == nil {
		panic("CoreFoundation: symbol CFPropertyListCreateWithStream not loaded")
	}
	return _cFPropertyListCreateWithStream(allocator, stream, streamLength, options, format, err)
}


var _cFPropertyListIsValid func(plist CFPropertyListRef, format CFPropertyListFormat) bool

// CFPropertyListIsValid determines if a property list is valid.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListIsValid(_:_:)
func CFPropertyListIsValid(plist CFPropertyListRef, format CFPropertyListFormat) bool {
	if _cFPropertyListIsValid == nil {
		panic("CoreFoundation: symbol CFPropertyListIsValid not loaded")
	}
	return _cFPropertyListIsValid(plist, format)
}


var _cFPropertyListWrite func(propertyList CFPropertyListRef, stream CFWriteStreamRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) int

// CFPropertyListWrite write the bytes of a serialized property list out to a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListWrite(_:_:_:_:_:)
func CFPropertyListWrite(propertyList CFPropertyListRef, stream CFWriteStreamRef, format CFPropertyListFormat, options uint64, err *CFErrorRef) int {
	if _cFPropertyListWrite == nil {
		panic("CoreFoundation: symbol CFPropertyListWrite not loaded")
	}
	return _cFPropertyListWrite(propertyList, stream, format, options, err)
}



var _cFReadStreamClose func(stream CFReadStreamRef) 

// CFReadStreamClose closes a readable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamClose(_:)
func CFReadStreamClose(stream CFReadStreamRef)  {
	if _cFReadStreamClose == nil {
		panic("CoreFoundation: symbol CFReadStreamClose not loaded")
	}
	_cFReadStreamClose(stream)
}


var _cFReadStreamCopyDispatchQueue func(stream CFReadStreamRef) uintptr

// CFReadStreamCopyDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyDispatchQueue(_:)
func CFReadStreamCopyDispatchQueue(stream CFReadStreamRef) dispatch.Queue {
	if _cFReadStreamCopyDispatchQueue == nil {
		panic("CoreFoundation: symbol CFReadStreamCopyDispatchQueue not loaded")
	}
	return dispatch.QueueFromHandle(_cFReadStreamCopyDispatchQueue(stream))
}


var _cFReadStreamCopyError func(stream CFReadStreamRef) CFErrorRef

// CFReadStreamCopyError returns the error associated with a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyError(_:)
func CFReadStreamCopyError(stream CFReadStreamRef) CFErrorRef {
	if _cFReadStreamCopyError == nil {
		panic("CoreFoundation: symbol CFReadStreamCopyError not loaded")
	}
	return _cFReadStreamCopyError(stream)
}


var _cFReadStreamCopyProperty func(stream CFReadStreamRef, propertyName CFStreamPropertyKey) CFTypeRef

// CFReadStreamCopyProperty returns the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyProperty(_:_:)
func CFReadStreamCopyProperty(stream CFReadStreamRef, propertyName CFStreamPropertyKey) CFTypeRef {
	if _cFReadStreamCopyProperty == nil {
		panic("CoreFoundation: symbol CFReadStreamCopyProperty not loaded")
	}
	return _cFReadStreamCopyProperty(stream, propertyName)
}


var _cFReadStreamCreateWithBytesNoCopy func(alloc CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFReadStreamRef

// CFReadStreamCreateWithBytesNoCopy creates a readable stream for a block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithBytesNoCopy(_:_:_:_:)
func CFReadStreamCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes *uint8, length int, bytesDeallocator CFAllocatorRef) CFReadStreamRef {
	if _cFReadStreamCreateWithBytesNoCopy == nil {
		panic("CoreFoundation: symbol CFReadStreamCreateWithBytesNoCopy not loaded")
	}
	return _cFReadStreamCreateWithBytesNoCopy(alloc, bytes, length, bytesDeallocator)
}


var _cFReadStreamCreateWithFile func(alloc CFAllocatorRef, fileURL CFURLRef) CFReadStreamRef

// CFReadStreamCreateWithFile creates a readable stream for a file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithFile(_:_:)
func CFReadStreamCreateWithFile(alloc CFAllocatorRef, fileURL CFURLRef) CFReadStreamRef {
	if _cFReadStreamCreateWithFile == nil {
		panic("CoreFoundation: symbol CFReadStreamCreateWithFile not loaded")
	}
	return _cFReadStreamCreateWithFile(alloc, fileURL)
}


var _cFReadStreamGetBuffer func(stream CFReadStreamRef, maxBytesToRead int, numBytesRead *int) *uint8

// CFReadStreamGetBuffer returns a pointer to a stream’s internal buffer of unread data, if possible.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetBuffer(_:_:_:)
func CFReadStreamGetBuffer(stream CFReadStreamRef, maxBytesToRead int, numBytesRead *int) *uint8 {
	if _cFReadStreamGetBuffer == nil {
		panic("CoreFoundation: symbol CFReadStreamGetBuffer not loaded")
	}
	return _cFReadStreamGetBuffer(stream, maxBytesToRead, numBytesRead)
}


var _cFReadStreamGetError func(stream CFReadStreamRef) CFStreamError

// CFReadStreamGetError returns the error status of a stream.
//
// Deprecated: Use [CFReadStreamCopyError(_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFReadStreamCopyError(_:)>) instead.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetError(_:)
func CFReadStreamGetError(stream CFReadStreamRef) CFStreamError {
	if _cFReadStreamGetError == nil {
		panic("CoreFoundation: symbol CFReadStreamGetError not loaded")
	}
	return _cFReadStreamGetError(stream)
}


var _cFReadStreamGetStatus func(stream CFReadStreamRef) CFStreamStatus

// CFReadStreamGetStatus returns the current state of a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetStatus(_:)
func CFReadStreamGetStatus(stream CFReadStreamRef) CFStreamStatus {
	if _cFReadStreamGetStatus == nil {
		panic("CoreFoundation: symbol CFReadStreamGetStatus not loaded")
	}
	return _cFReadStreamGetStatus(stream)
}


var _cFReadStreamGetTypeID func() uint

// CFReadStreamGetTypeID returns the type identifier the [CFReadStream] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetTypeID()
func CFReadStreamGetTypeID() uint {
	if _cFReadStreamGetTypeID == nil {
		panic("CoreFoundation: symbol CFReadStreamGetTypeID not loaded")
	}
	return _cFReadStreamGetTypeID()
}


var _cFReadStreamHasBytesAvailable func(stream CFReadStreamRef) bool

// CFReadStreamHasBytesAvailable returns a Boolean value that indicates whether a readable stream has data that can be read without blocking.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamHasBytesAvailable(_:)
func CFReadStreamHasBytesAvailable(stream CFReadStreamRef) bool {
	if _cFReadStreamHasBytesAvailable == nil {
		panic("CoreFoundation: symbol CFReadStreamHasBytesAvailable not loaded")
	}
	return _cFReadStreamHasBytesAvailable(stream)
}


var _cFReadStreamOpen func(stream CFReadStreamRef) bool

// CFReadStreamOpen opens a stream for reading.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamOpen(_:)
func CFReadStreamOpen(stream CFReadStreamRef) bool {
	if _cFReadStreamOpen == nil {
		panic("CoreFoundation: symbol CFReadStreamOpen not loaded")
	}
	return _cFReadStreamOpen(stream)
}


var _cFReadStreamRead func(stream CFReadStreamRef, buffer *uint8, bufferLength int) int

// CFReadStreamRead reads data from a readable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamRead(_:_:_:)
func CFReadStreamRead(stream CFReadStreamRef, buffer *uint8, bufferLength int) int {
	if _cFReadStreamRead == nil {
		panic("CoreFoundation: symbol CFReadStreamRead not loaded")
	}
	return _cFReadStreamRead(stream, buffer, bufferLength)
}


var _cFReadStreamScheduleWithRunLoop func(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) 

// CFReadStreamScheduleWithRunLoop schedules a stream into a run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamScheduleWithRunLoop(_:_:_:)
func CFReadStreamScheduleWithRunLoop(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)  {
	if _cFReadStreamScheduleWithRunLoop == nil {
		panic("CoreFoundation: symbol CFReadStreamScheduleWithRunLoop not loaded")
	}
	_cFReadStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
}


var _cFReadStreamSetClient func(stream CFReadStreamRef, streamEvents uint64, clientCB CFReadStreamClientCallBack, clientContext *CFStreamClientContext) bool

// CFReadStreamSetClient assigns a client to a stream, which receives callbacks when certain events occur.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetClient(_:_:_:_:)
func CFReadStreamSetClient(stream CFReadStreamRef, streamEvents uint64, clientCB CFReadStreamClientCallBack, clientContext *CFStreamClientContext) bool {
	if _cFReadStreamSetClient == nil {
		panic("CoreFoundation: symbol CFReadStreamSetClient not loaded")
	}
	return _cFReadStreamSetClient(stream, streamEvents, clientCB, clientContext)
}


var _cFReadStreamSetDispatchQueue func(stream CFReadStreamRef, q uintptr) 

// CFReadStreamSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetDispatchQueue(_:_:)
func CFReadStreamSetDispatchQueue(stream CFReadStreamRef, q dispatch.Queue)  {
	if _cFReadStreamSetDispatchQueue == nil {
		panic("CoreFoundation: symbol CFReadStreamSetDispatchQueue not loaded")
	}
	_cFReadStreamSetDispatchQueue(stream, uintptr(q.Handle()))
}


var _cFReadStreamSetProperty func(stream CFReadStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool

// CFReadStreamSetProperty sets the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetProperty(_:_:_:)
func CFReadStreamSetProperty(stream CFReadStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool {
	if _cFReadStreamSetProperty == nil {
		panic("CoreFoundation: symbol CFReadStreamSetProperty not loaded")
	}
	return _cFReadStreamSetProperty(stream, propertyName, propertyValue)
}


var _cFReadStreamUnscheduleFromRunLoop func(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) 

// CFReadStreamUnscheduleFromRunLoop removes a read stream from a given run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamUnscheduleFromRunLoop(_:_:_:)
func CFReadStreamUnscheduleFromRunLoop(stream CFReadStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)  {
	if _cFReadStreamUnscheduleFromRunLoop == nil {
		panic("CoreFoundation: symbol CFReadStreamUnscheduleFromRunLoop not loaded")
	}
	_cFReadStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
}


var _cFRelease func(cf CFTypeRef) 

// CFRelease releases a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRelease
func CFRelease(cf CFTypeRef)  {
	if _cFRelease == nil {
		panic("CoreFoundation: symbol CFRelease not loaded")
	}
	_cFRelease(cf)
}


var _cFRetain func(cf CFTypeRef) CFTypeRef

// CFRetain retains a Core Foundation object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRetain
func CFRetain(cf CFTypeRef) CFTypeRef {
	if _cFRetain == nil {
		panic("CoreFoundation: symbol CFRetain not loaded")
	}
	return _cFRetain(cf)
}


var _cFRunLoopAddCommonMode func(rl CFRunLoopRef, mode CFRunLoopMode) 

// CFRunLoopAddCommonMode adds a mode to the set of run loop common modes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddCommonMode(_:_:)
func CFRunLoopAddCommonMode(rl CFRunLoopRef, mode CFRunLoopMode)  {
	if _cFRunLoopAddCommonMode == nil {
		panic("CoreFoundation: symbol CFRunLoopAddCommonMode not loaded")
	}
	_cFRunLoopAddCommonMode(rl, mode)
}


var _cFRunLoopAddObserver func(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) 

// CFRunLoopAddObserver adds a CFRunLoopObserver object to a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddObserver(_:_:_:)
func CFRunLoopAddObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode)  {
	if _cFRunLoopAddObserver == nil {
		panic("CoreFoundation: symbol CFRunLoopAddObserver not loaded")
	}
	_cFRunLoopAddObserver(rl, observer, mode)
}


var _cFRunLoopAddSource func(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) 

// CFRunLoopAddSource adds a CFRunLoopSource object to a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddSource(_:_:_:)
func CFRunLoopAddSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode)  {
	if _cFRunLoopAddSource == nil {
		panic("CoreFoundation: symbol CFRunLoopAddSource not loaded")
	}
	_cFRunLoopAddSource(rl, source, mode)
}


var _cFRunLoopAddTimer func(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) 

// CFRunLoopAddTimer adds a CFRunLoopTimer object to a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddTimer(_:_:_:)
func CFRunLoopAddTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode)  {
	if _cFRunLoopAddTimer == nil {
		panic("CoreFoundation: symbol CFRunLoopAddTimer not loaded")
	}
	_cFRunLoopAddTimer(rl, timer, mode)
}


var _cFRunLoopContainsObserver func(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) bool

// CFRunLoopContainsObserver returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopObserver object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsObserver(_:_:_:)
func CFRunLoopContainsObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) bool {
	if _cFRunLoopContainsObserver == nil {
		panic("CoreFoundation: symbol CFRunLoopContainsObserver not loaded")
	}
	return _cFRunLoopContainsObserver(rl, observer, mode)
}


var _cFRunLoopContainsSource func(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) bool

// CFRunLoopContainsSource returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsSource(_:_:_:)
func CFRunLoopContainsSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) bool {
	if _cFRunLoopContainsSource == nil {
		panic("CoreFoundation: symbol CFRunLoopContainsSource not loaded")
	}
	return _cFRunLoopContainsSource(rl, source, mode)
}


var _cFRunLoopContainsTimer func(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) bool

// CFRunLoopContainsTimer returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsTimer(_:_:_:)
func CFRunLoopContainsTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) bool {
	if _cFRunLoopContainsTimer == nil {
		panic("CoreFoundation: symbol CFRunLoopContainsTimer not loaded")
	}
	return _cFRunLoopContainsTimer(rl, timer, mode)
}


var _cFRunLoopCopyAllModes func(rl CFRunLoopRef) CFArrayRef

// CFRunLoopCopyAllModes returns an array that contains all the defined modes for a CFRunLoop object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyAllModes(_:)
func CFRunLoopCopyAllModes(rl CFRunLoopRef) CFArrayRef {
	if _cFRunLoopCopyAllModes == nil {
		panic("CoreFoundation: symbol CFRunLoopCopyAllModes not loaded")
	}
	return _cFRunLoopCopyAllModes(rl)
}


var _cFRunLoopCopyCurrentMode func(rl CFRunLoopRef) CFRunLoopMode

// CFRunLoopCopyCurrentMode returns the name of the mode in which a given run loop is currently running.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyCurrentMode(_:)
func CFRunLoopCopyCurrentMode(rl CFRunLoopRef) CFRunLoopMode {
	if _cFRunLoopCopyCurrentMode == nil {
		panic("CoreFoundation: symbol CFRunLoopCopyCurrentMode not loaded")
	}
	return _cFRunLoopCopyCurrentMode(rl)
}


var _cFRunLoopGetCurrent func() CFRunLoopRef

// CFRunLoopGetCurrent returns the CFRunLoop object for the current thread.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetCurrent()
func CFRunLoopGetCurrent() CFRunLoopRef {
	if _cFRunLoopGetCurrent == nil {
		panic("CoreFoundation: symbol CFRunLoopGetCurrent not loaded")
	}
	return _cFRunLoopGetCurrent()
}


var _cFRunLoopGetMain func() CFRunLoopRef

// CFRunLoopGetMain returns the main CFRunLoop object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetMain()
func CFRunLoopGetMain() CFRunLoopRef {
	if _cFRunLoopGetMain == nil {
		panic("CoreFoundation: symbol CFRunLoopGetMain not loaded")
	}
	return _cFRunLoopGetMain()
}


var _cFRunLoopGetNextTimerFireDate func(rl CFRunLoopRef, mode CFRunLoopMode) CFAbsoluteTime

// CFRunLoopGetNextTimerFireDate returns the time at which the next timer will fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetNextTimerFireDate(_:_:)
func CFRunLoopGetNextTimerFireDate(rl CFRunLoopRef, mode CFRunLoopMode) CFAbsoluteTime {
	if _cFRunLoopGetNextTimerFireDate == nil {
		panic("CoreFoundation: symbol CFRunLoopGetNextTimerFireDate not loaded")
	}
	return _cFRunLoopGetNextTimerFireDate(rl, mode)
}


var _cFRunLoopGetTypeID func() uint

// CFRunLoopGetTypeID returns the type identifier for the CFRunLoop opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetTypeID()
func CFRunLoopGetTypeID() uint {
	if _cFRunLoopGetTypeID == nil {
		panic("CoreFoundation: symbol CFRunLoopGetTypeID not loaded")
	}
	return _cFRunLoopGetTypeID()
}


var _cFRunLoopIsWaiting func(rl CFRunLoopRef) bool

// CFRunLoopIsWaiting returns a Boolean value that indicates whether the run loop is waiting for an event.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopIsWaiting(_:)
func CFRunLoopIsWaiting(rl CFRunLoopRef) bool {
	if _cFRunLoopIsWaiting == nil {
		panic("CoreFoundation: symbol CFRunLoopIsWaiting not loaded")
	}
	return _cFRunLoopIsWaiting(rl)
}


var _cFRunLoopObserverCreate func(allocator CFAllocatorRef, activities uint64, repeats bool, order int, callout CFRunLoopObserverCallBack, context *CFRunLoopObserverContext) CFRunLoopObserverRef

// CFRunLoopObserverCreate creates a CFRunLoopObserver object with a function callback.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreate(_:_:_:_:_:_:)
func CFRunLoopObserverCreate(allocator CFAllocatorRef, activities uint64, repeats bool, order int, callout CFRunLoopObserverCallBack, context *CFRunLoopObserverContext) CFRunLoopObserverRef {
	if _cFRunLoopObserverCreate == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverCreate not loaded")
	}
	return _cFRunLoopObserverCreate(allocator, activities, repeats, order, callout, context)
}


var _cFRunLoopObserverCreateWithHandler func(allocator CFAllocatorRef, activities uint64, repeats bool, order int) CFRunLoopObserverRef

// CFRunLoopObserverCreateWithHandler creates a CFRunLoopObserver object with a block-based handler.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreateWithHandler(_:_:_:_:_:)
func CFRunLoopObserverCreateWithHandler(allocator CFAllocatorRef, activities uint64, repeats bool, order int) CFRunLoopObserverRef {
	if _cFRunLoopObserverCreateWithHandler == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverCreateWithHandler not loaded")
	}
	return _cFRunLoopObserverCreateWithHandler(allocator, activities, repeats, order)
}


var _cFRunLoopObserverDoesRepeat func(observer CFRunLoopObserverRef) bool

// CFRunLoopObserverDoesRepeat returns a Boolean value that indicates whether a CFRunLoopObserver repeats.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverDoesRepeat(_:)
func CFRunLoopObserverDoesRepeat(observer CFRunLoopObserverRef) bool {
	if _cFRunLoopObserverDoesRepeat == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverDoesRepeat not loaded")
	}
	return _cFRunLoopObserverDoesRepeat(observer)
}


var _cFRunLoopObserverGetActivities func(observer CFRunLoopObserverRef) uint64

// CFRunLoopObserverGetActivities returns the run loop stages during which an observer runs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetActivities(_:)
func CFRunLoopObserverGetActivities(observer CFRunLoopObserverRef) uint64 {
	if _cFRunLoopObserverGetActivities == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverGetActivities not loaded")
	}
	return _cFRunLoopObserverGetActivities(observer)
}


var _cFRunLoopObserverGetContext func(observer CFRunLoopObserverRef, context *CFRunLoopObserverContext) 

// CFRunLoopObserverGetContext returns the context information for a CFRunLoopObserver object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetContext(_:_:)
func CFRunLoopObserverGetContext(observer CFRunLoopObserverRef, context *CFRunLoopObserverContext)  {
	if _cFRunLoopObserverGetContext == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverGetContext not loaded")
	}
	_cFRunLoopObserverGetContext(observer, context)
}


var _cFRunLoopObserverGetOrder func(observer CFRunLoopObserverRef) int

// CFRunLoopObserverGetOrder returns the ordering parameter for a CFRunLoopObserver object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetOrder(_:)
func CFRunLoopObserverGetOrder(observer CFRunLoopObserverRef) int {
	if _cFRunLoopObserverGetOrder == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverGetOrder not loaded")
	}
	return _cFRunLoopObserverGetOrder(observer)
}


var _cFRunLoopObserverGetTypeID func() uint

// CFRunLoopObserverGetTypeID returns the type identifier for the CFRunLoopObserver opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetTypeID()
func CFRunLoopObserverGetTypeID() uint {
	if _cFRunLoopObserverGetTypeID == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverGetTypeID not loaded")
	}
	return _cFRunLoopObserverGetTypeID()
}


var _cFRunLoopObserverInvalidate func(observer CFRunLoopObserverRef) 

// CFRunLoopObserverInvalidate invalidates a CFRunLoopObserver object, stopping it from ever firing again.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverInvalidate(_:)
func CFRunLoopObserverInvalidate(observer CFRunLoopObserverRef)  {
	if _cFRunLoopObserverInvalidate == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverInvalidate not loaded")
	}
	_cFRunLoopObserverInvalidate(observer)
}


var _cFRunLoopObserverIsValid func(observer CFRunLoopObserverRef) bool

// CFRunLoopObserverIsValid returns a Boolean value that indicates whether a CFRunLoopObserver object is valid and able to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverIsValid(_:)
func CFRunLoopObserverIsValid(observer CFRunLoopObserverRef) bool {
	if _cFRunLoopObserverIsValid == nil {
		panic("CoreFoundation: symbol CFRunLoopObserverIsValid not loaded")
	}
	return _cFRunLoopObserverIsValid(observer)
}


var _cFRunLoopPerformBlock func(rl CFRunLoopRef, mode CFTypeRef) 

// CFRunLoopPerformBlock enqueues a block object on a given runloop to be executed as the runloop cycles in specified modes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopPerformBlock(_:_:_:)
func CFRunLoopPerformBlock(rl CFRunLoopRef, mode CFTypeRef)  {
	if _cFRunLoopPerformBlock == nil {
		panic("CoreFoundation: symbol CFRunLoopPerformBlock not loaded")
	}
	_cFRunLoopPerformBlock(rl, mode)
}


var _cFRunLoopRemoveObserver func(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode) 

// CFRunLoopRemoveObserver removes a CFRunLoopObserver object from a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveObserver(_:_:_:)
func CFRunLoopRemoveObserver(rl CFRunLoopRef, observer CFRunLoopObserverRef, mode CFRunLoopMode)  {
	if _cFRunLoopRemoveObserver == nil {
		panic("CoreFoundation: symbol CFRunLoopRemoveObserver not loaded")
	}
	_cFRunLoopRemoveObserver(rl, observer, mode)
}


var _cFRunLoopRemoveSource func(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode) 

// CFRunLoopRemoveSource removes a CFRunLoopSource object from a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveSource(_:_:_:)
func CFRunLoopRemoveSource(rl CFRunLoopRef, source CFRunLoopSourceRef, mode CFRunLoopMode)  {
	if _cFRunLoopRemoveSource == nil {
		panic("CoreFoundation: symbol CFRunLoopRemoveSource not loaded")
	}
	_cFRunLoopRemoveSource(rl, source, mode)
}


var _cFRunLoopRemoveTimer func(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode) 

// CFRunLoopRemoveTimer removes a CFRunLoopTimer object from a run loop mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveTimer(_:_:_:)
func CFRunLoopRemoveTimer(rl CFRunLoopRef, timer CFRunLoopTimerRef, mode CFRunLoopMode)  {
	if _cFRunLoopRemoveTimer == nil {
		panic("CoreFoundation: symbol CFRunLoopRemoveTimer not loaded")
	}
	_cFRunLoopRemoveTimer(rl, timer, mode)
}


var _cFRunLoopRun func() 

// CFRunLoopRun runs the current thread’s CFRunLoop object in its default mode indefinitely.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRun()
func CFRunLoopRun()  {
	if _cFRunLoopRun == nil {
		panic("CoreFoundation: symbol CFRunLoopRun not loaded")
	}
	_cFRunLoopRun()
}


var _cFRunLoopRunInMode func(mode CFRunLoopMode, seconds float64, returnAfterSourceHandled bool) CFRunLoopRunResult

// CFRunLoopRunInMode runs the current thread’s CFRunLoop object in a particular mode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunInMode(_:_:_:)
func CFRunLoopRunInMode(mode CFRunLoopMode, seconds float64, returnAfterSourceHandled bool) CFRunLoopRunResult {
	if _cFRunLoopRunInMode == nil {
		panic("CoreFoundation: symbol CFRunLoopRunInMode not loaded")
	}
	return _cFRunLoopRunInMode(mode, seconds, returnAfterSourceHandled)
}


var _cFRunLoopSourceCreate func(allocator CFAllocatorRef, order int, context *CFRunLoopSourceContext) CFRunLoopSourceRef

// CFRunLoopSourceCreate creates a CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceCreate(_:_:_:)
func CFRunLoopSourceCreate(allocator CFAllocatorRef, order int, context *CFRunLoopSourceContext) CFRunLoopSourceRef {
	if _cFRunLoopSourceCreate == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceCreate not loaded")
	}
	return _cFRunLoopSourceCreate(allocator, order, context)
}


var _cFRunLoopSourceGetContext func(source CFRunLoopSourceRef, context *CFRunLoopSourceContext) 

// CFRunLoopSourceGetContext returns the context information for a CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetContext(_:_:)
func CFRunLoopSourceGetContext(source CFRunLoopSourceRef, context *CFRunLoopSourceContext)  {
	if _cFRunLoopSourceGetContext == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceGetContext not loaded")
	}
	_cFRunLoopSourceGetContext(source, context)
}


var _cFRunLoopSourceGetOrder func(source CFRunLoopSourceRef) int

// CFRunLoopSourceGetOrder returns the ordering parameter for a CFRunLoopSource object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetOrder(_:)
func CFRunLoopSourceGetOrder(source CFRunLoopSourceRef) int {
	if _cFRunLoopSourceGetOrder == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceGetOrder not loaded")
	}
	return _cFRunLoopSourceGetOrder(source)
}


var _cFRunLoopSourceGetTypeID func() uint

// CFRunLoopSourceGetTypeID returns the type identifier of the CFRunLoopSource opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetTypeID()
func CFRunLoopSourceGetTypeID() uint {
	if _cFRunLoopSourceGetTypeID == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceGetTypeID not loaded")
	}
	return _cFRunLoopSourceGetTypeID()
}


var _cFRunLoopSourceInvalidate func(source CFRunLoopSourceRef) 

// CFRunLoopSourceInvalidate invalidates a CFRunLoopSource object, stopping it from ever firing again.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceInvalidate(_:)
func CFRunLoopSourceInvalidate(source CFRunLoopSourceRef)  {
	if _cFRunLoopSourceInvalidate == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceInvalidate not loaded")
	}
	_cFRunLoopSourceInvalidate(source)
}


var _cFRunLoopSourceIsValid func(source CFRunLoopSourceRef) bool

// CFRunLoopSourceIsValid returns a Boolean value that indicates whether a CFRunLoopSource object is valid and able to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceIsValid(_:)
func CFRunLoopSourceIsValid(source CFRunLoopSourceRef) bool {
	if _cFRunLoopSourceIsValid == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceIsValid not loaded")
	}
	return _cFRunLoopSourceIsValid(source)
}


var _cFRunLoopSourceSignal func(source CFRunLoopSourceRef) 

// CFRunLoopSourceSignal signals a CFRunLoopSource object, marking it as ready to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceSignal(_:)
func CFRunLoopSourceSignal(source CFRunLoopSourceRef)  {
	if _cFRunLoopSourceSignal == nil {
		panic("CoreFoundation: symbol CFRunLoopSourceSignal not loaded")
	}
	_cFRunLoopSourceSignal(source)
}


var _cFRunLoopStop func(rl CFRunLoopRef) 

// CFRunLoopStop forces a CFRunLoop object to stop running.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopStop(_:)
func CFRunLoopStop(rl CFRunLoopRef)  {
	if _cFRunLoopStop == nil {
		panic("CoreFoundation: symbol CFRunLoopStop not loaded")
	}
	_cFRunLoopStop(rl)
}


var _cFRunLoopTimerCreate func(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int, callout CFRunLoopTimerCallBack, context *CFRunLoopTimerContext) CFRunLoopTimerRef

// CFRunLoopTimerCreate creates a new CFRunLoopTimer object with a function callback.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreate(_:_:_:_:_:_:_:)
func CFRunLoopTimerCreate(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int, callout CFRunLoopTimerCallBack, context *CFRunLoopTimerContext) CFRunLoopTimerRef {
	if _cFRunLoopTimerCreate == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerCreate not loaded")
	}
	return _cFRunLoopTimerCreate(allocator, fireDate, interval, flags, order, callout, context)
}


var _cFRunLoopTimerCreateWithHandler func(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int) CFRunLoopTimerRef

// CFRunLoopTimerCreateWithHandler creates a new CFRunLoopTimer object with a block-based handler.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreateWithHandler(_:_:_:_:_:_:)
func CFRunLoopTimerCreateWithHandler(allocator CFAllocatorRef, fireDate CFAbsoluteTime, interval float64, flags uint64, order int) CFRunLoopTimerRef {
	if _cFRunLoopTimerCreateWithHandler == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerCreateWithHandler not loaded")
	}
	return _cFRunLoopTimerCreateWithHandler(allocator, fireDate, interval, flags, order)
}


var _cFRunLoopTimerDoesRepeat func(timer CFRunLoopTimerRef) bool

// CFRunLoopTimerDoesRepeat returns a Boolean value that indicates whether a CFRunLoopTimer object repeats.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerDoesRepeat(_:)
func CFRunLoopTimerDoesRepeat(timer CFRunLoopTimerRef) bool {
	if _cFRunLoopTimerDoesRepeat == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerDoesRepeat not loaded")
	}
	return _cFRunLoopTimerDoesRepeat(timer)
}


var _cFRunLoopTimerGetContext func(timer CFRunLoopTimerRef, context *CFRunLoopTimerContext) 

// CFRunLoopTimerGetContext returns the context information for a CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetContext(_:_:)
func CFRunLoopTimerGetContext(timer CFRunLoopTimerRef, context *CFRunLoopTimerContext)  {
	if _cFRunLoopTimerGetContext == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerGetContext not loaded")
	}
	_cFRunLoopTimerGetContext(timer, context)
}


var _cFRunLoopTimerGetInterval func(timer CFRunLoopTimerRef) float64

// CFRunLoopTimerGetInterval returns the firing interval of a repeating CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetInterval(_:)
func CFRunLoopTimerGetInterval(timer CFRunLoopTimerRef) float64 {
	if _cFRunLoopTimerGetInterval == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerGetInterval not loaded")
	}
	return _cFRunLoopTimerGetInterval(timer)
}


var _cFRunLoopTimerGetNextFireDate func(timer CFRunLoopTimerRef) CFAbsoluteTime

// CFRunLoopTimerGetNextFireDate returns the next firing time for a CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetNextFireDate(_:)
func CFRunLoopTimerGetNextFireDate(timer CFRunLoopTimerRef) CFAbsoluteTime {
	if _cFRunLoopTimerGetNextFireDate == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerGetNextFireDate not loaded")
	}
	return _cFRunLoopTimerGetNextFireDate(timer)
}


var _cFRunLoopTimerGetOrder func(timer CFRunLoopTimerRef) int

// CFRunLoopTimerGetOrder returns the ordering parameter for a CFRunLoopTimer object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetOrder(_:)
func CFRunLoopTimerGetOrder(timer CFRunLoopTimerRef) int {
	if _cFRunLoopTimerGetOrder == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerGetOrder not loaded")
	}
	return _cFRunLoopTimerGetOrder(timer)
}


var _cFRunLoopTimerGetTolerance func(timer CFRunLoopTimerRef) float64

// CFRunLoopTimerGetTolerance.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTolerance(_:)
func CFRunLoopTimerGetTolerance(timer CFRunLoopTimerRef) float64 {
	if _cFRunLoopTimerGetTolerance == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerGetTolerance not loaded")
	}
	return _cFRunLoopTimerGetTolerance(timer)
}


var _cFRunLoopTimerGetTypeID func() uint

// CFRunLoopTimerGetTypeID returns the type identifier of the CFRunLoopTimer opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTypeID()
func CFRunLoopTimerGetTypeID() uint {
	if _cFRunLoopTimerGetTypeID == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerGetTypeID not loaded")
	}
	return _cFRunLoopTimerGetTypeID()
}


var _cFRunLoopTimerInvalidate func(timer CFRunLoopTimerRef) 

// CFRunLoopTimerInvalidate invalidates a CFRunLoopTimer object, stopping it from ever firing again.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerInvalidate(_:)
func CFRunLoopTimerInvalidate(timer CFRunLoopTimerRef)  {
	if _cFRunLoopTimerInvalidate == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerInvalidate not loaded")
	}
	_cFRunLoopTimerInvalidate(timer)
}


var _cFRunLoopTimerIsValid func(timer CFRunLoopTimerRef) bool

// CFRunLoopTimerIsValid returns a Boolean value that indicates whether a CFRunLoopTimer object is valid and able to fire.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerIsValid(_:)
func CFRunLoopTimerIsValid(timer CFRunLoopTimerRef) bool {
	if _cFRunLoopTimerIsValid == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerIsValid not loaded")
	}
	return _cFRunLoopTimerIsValid(timer)
}


var _cFRunLoopTimerSetNextFireDate func(timer CFRunLoopTimerRef, fireDate CFAbsoluteTime) 

// CFRunLoopTimerSetNextFireDate sets the next firing date for a CFRunLoopTimer object .
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetNextFireDate(_:_:)
func CFRunLoopTimerSetNextFireDate(timer CFRunLoopTimerRef, fireDate CFAbsoluteTime)  {
	if _cFRunLoopTimerSetNextFireDate == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerSetNextFireDate not loaded")
	}
	_cFRunLoopTimerSetNextFireDate(timer, fireDate)
}


var _cFRunLoopTimerSetTolerance func(timer CFRunLoopTimerRef, tolerance float64) 

// CFRunLoopTimerSetTolerance.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetTolerance(_:_:)
func CFRunLoopTimerSetTolerance(timer CFRunLoopTimerRef, tolerance float64)  {
	if _cFRunLoopTimerSetTolerance == nil {
		panic("CoreFoundation: symbol CFRunLoopTimerSetTolerance not loaded")
	}
	_cFRunLoopTimerSetTolerance(timer, tolerance)
}


var _cFRunLoopWakeUp func(rl CFRunLoopRef) 

// CFRunLoopWakeUp wakes a waiting CFRunLoop object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopWakeUp(_:)
func CFRunLoopWakeUp(rl CFRunLoopRef)  {
	if _cFRunLoopWakeUp == nil {
		panic("CoreFoundation: symbol CFRunLoopWakeUp not loaded")
	}
	_cFRunLoopWakeUp(rl)
}


var _cFSetAddValue func(theSet CFMutableSetRef, value unsafe.Pointer) 

// CFSetAddValue adds a value to a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetAddValue(_:_:)
func CFSetAddValue(theSet CFMutableSetRef, value unsafe.Pointer)  {
	if _cFSetAddValue == nil {
		panic("CoreFoundation: symbol CFSetAddValue not loaded")
	}
	_cFSetAddValue(theSet, value)
}


var _cFSetApplyFunction func(theSet CFSetRef, applier CFSetApplierFunction, context unsafe.Pointer) 

// CFSetApplyFunction calls a function once for each value in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetApplyFunction(_:_:_:)
func CFSetApplyFunction(theSet CFSetRef, applier CFSetApplierFunction, context unsafe.Pointer)  {
	if _cFSetApplyFunction == nil {
		panic("CoreFoundation: symbol CFSetApplyFunction not loaded")
	}
	_cFSetApplyFunction(theSet, applier, context)
}


var _cFSetContainsValue func(theSet CFSetRef, value unsafe.Pointer) bool

// CFSetContainsValue returns a Boolean that indicates whether a set contains a given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetContainsValue(_:_:)
func CFSetContainsValue(theSet CFSetRef, value unsafe.Pointer) bool {
	if _cFSetContainsValue == nil {
		panic("CoreFoundation: symbol CFSetContainsValue not loaded")
	}
	return _cFSetContainsValue(theSet, value)
}


var _cFSetCreate func(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFSetCallBacks) CFSetRef

// CFSetCreate creates an immutable CFSet object containing supplied values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreate(_:_:_:_:)
func CFSetCreate(allocator CFAllocatorRef, values unsafe.Pointer, numValues int, callBacks *CFSetCallBacks) CFSetRef {
	if _cFSetCreate == nil {
		panic("CoreFoundation: symbol CFSetCreate not loaded")
	}
	return _cFSetCreate(allocator, values, numValues, callBacks)
}


var _cFSetCreateCopy func(allocator CFAllocatorRef, theSet CFSetRef) CFSetRef

// CFSetCreateCopy creates an immutable set containing the values of an existing set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateCopy(_:_:)
func CFSetCreateCopy(allocator CFAllocatorRef, theSet CFSetRef) CFSetRef {
	if _cFSetCreateCopy == nil {
		panic("CoreFoundation: symbol CFSetCreateCopy not loaded")
	}
	return _cFSetCreateCopy(allocator, theSet)
}


var _cFSetCreateMutable func(allocator CFAllocatorRef, capacity int, callBacks *CFSetCallBacks) CFMutableSetRef

// CFSetCreateMutable creates an empty CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutable(_:_:_:)
func CFSetCreateMutable(allocator CFAllocatorRef, capacity int, callBacks *CFSetCallBacks) CFMutableSetRef {
	if _cFSetCreateMutable == nil {
		panic("CoreFoundation: symbol CFSetCreateMutable not loaded")
	}
	return _cFSetCreateMutable(allocator, capacity, callBacks)
}


var _cFSetCreateMutableCopy func(allocator CFAllocatorRef, capacity int, theSet CFSetRef) CFMutableSetRef

// CFSetCreateMutableCopy creates a new mutable set with the values from another set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutableCopy(_:_:_:)
func CFSetCreateMutableCopy(allocator CFAllocatorRef, capacity int, theSet CFSetRef) CFMutableSetRef {
	if _cFSetCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFSetCreateMutableCopy not loaded")
	}
	return _cFSetCreateMutableCopy(allocator, capacity, theSet)
}


var _cFSetGetCount func(theSet CFSetRef) int

// CFSetGetCount returns the number of values currently in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCount(_:)
func CFSetGetCount(theSet CFSetRef) int {
	if _cFSetGetCount == nil {
		panic("CoreFoundation: symbol CFSetGetCount not loaded")
	}
	return _cFSetGetCount(theSet)
}


var _cFSetGetCountOfValue func(theSet CFSetRef, value unsafe.Pointer) int

// CFSetGetCountOfValue returns the number of values in a set that match a given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCountOfValue(_:_:)
func CFSetGetCountOfValue(theSet CFSetRef, value unsafe.Pointer) int {
	if _cFSetGetCountOfValue == nil {
		panic("CoreFoundation: symbol CFSetGetCountOfValue not loaded")
	}
	return _cFSetGetCountOfValue(theSet, value)
}


var _cFSetGetTypeID func() uint

// CFSetGetTypeID returns the type identifier for the CFSet type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetTypeID()
func CFSetGetTypeID() uint {
	if _cFSetGetTypeID == nil {
		panic("CoreFoundation: symbol CFSetGetTypeID not loaded")
	}
	return _cFSetGetTypeID()
}


var _cFSetGetValue func(theSet CFSetRef, value unsafe.Pointer) unsafe.Pointer

// CFSetGetValue obtains a specified value from a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValue(_:_:)
func CFSetGetValue(theSet CFSetRef, value unsafe.Pointer) unsafe.Pointer {
	if _cFSetGetValue == nil {
		panic("CoreFoundation: symbol CFSetGetValue not loaded")
	}
	return _cFSetGetValue(theSet, value)
}


var _cFSetGetValueIfPresent func(theSet CFSetRef, candidate unsafe.Pointer, value unsafe.Pointer) bool

// CFSetGetValueIfPresent reports whether or not a value is in a set, and if it exists returns the value indirectly.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValueIfPresent(_:_:_:)
func CFSetGetValueIfPresent(theSet CFSetRef, candidate unsafe.Pointer, value unsafe.Pointer) bool {
	if _cFSetGetValueIfPresent == nil {
		panic("CoreFoundation: symbol CFSetGetValueIfPresent not loaded")
	}
	return _cFSetGetValueIfPresent(theSet, candidate, value)
}


var _cFSetGetValues func(theSet CFSetRef, values unsafe.Pointer) 

// CFSetGetValues obtains all values in a set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValues(_:_:)
func CFSetGetValues(theSet CFSetRef, values unsafe.Pointer)  {
	if _cFSetGetValues == nil {
		panic("CoreFoundation: symbol CFSetGetValues not loaded")
	}
	_cFSetGetValues(theSet, values)
}


var _cFSetRemoveAllValues func(theSet CFMutableSetRef) 

// CFSetRemoveAllValues removes all values from a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveAllValues(_:)
func CFSetRemoveAllValues(theSet CFMutableSetRef)  {
	if _cFSetRemoveAllValues == nil {
		panic("CoreFoundation: symbol CFSetRemoveAllValues not loaded")
	}
	_cFSetRemoveAllValues(theSet)
}


var _cFSetRemoveValue func(theSet CFMutableSetRef, value unsafe.Pointer) 

// CFSetRemoveValue removes a value from a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveValue(_:_:)
func CFSetRemoveValue(theSet CFMutableSetRef, value unsafe.Pointer)  {
	if _cFSetRemoveValue == nil {
		panic("CoreFoundation: symbol CFSetRemoveValue not loaded")
	}
	_cFSetRemoveValue(theSet, value)
}


var _cFSetReplaceValue func(theSet CFMutableSetRef, value unsafe.Pointer) 

// CFSetReplaceValue replaces a value in a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetReplaceValue(_:_:)
func CFSetReplaceValue(theSet CFMutableSetRef, value unsafe.Pointer)  {
	if _cFSetReplaceValue == nil {
		panic("CoreFoundation: symbol CFSetReplaceValue not loaded")
	}
	_cFSetReplaceValue(theSet, value)
}


var _cFSetSetValue func(theSet CFMutableSetRef, value unsafe.Pointer) 

// CFSetSetValue sets a value in a CFMutableSet object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSetSetValue(_:_:)
func CFSetSetValue(theSet CFMutableSetRef, value unsafe.Pointer)  {
	if _cFSetSetValue == nil {
		panic("CoreFoundation: symbol CFSetSetValue not loaded")
	}
	_cFSetSetValue(theSet, value)
}


var _cFShow func(obj CFTypeRef) 

// CFShow prints a description of a Core Foundation object to stderr.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFShow(_:)
func CFShow(obj CFTypeRef)  {
	if _cFShow == nil {
		panic("CoreFoundation: symbol CFShow not loaded")
	}
	_cFShow(obj)
}


var _cFShowStr func(str CFStringRef) 

// CFShowStr prints the attributes of a string during debugging.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFShowStr(_:)
func CFShowStr(str CFStringRef)  {
	if _cFShowStr == nil {
		panic("CoreFoundation: symbol CFShowStr not loaded")
	}
	_cFShowStr(str)
}


var _cFSocketConnectToAddress func(s CFSocketRef, address CFDataRef, timeout float64) CFSocketError

// CFSocketConnectToAddress opens a connection to a remote socket.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketConnectToAddress(_:_:_:)
func CFSocketConnectToAddress(s CFSocketRef, address CFDataRef, timeout float64) CFSocketError {
	if _cFSocketConnectToAddress == nil {
		panic("CoreFoundation: symbol CFSocketConnectToAddress not loaded")
	}
	return _cFSocketConnectToAddress(s, address, timeout)
}


var _cFSocketCopyAddress func(s CFSocketRef) CFDataRef

// CFSocketCopyAddress returns the local address of a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyAddress(_:)
func CFSocketCopyAddress(s CFSocketRef) CFDataRef {
	if _cFSocketCopyAddress == nil {
		panic("CoreFoundation: symbol CFSocketCopyAddress not loaded")
	}
	return _cFSocketCopyAddress(s)
}


var _cFSocketCopyPeerAddress func(s CFSocketRef) CFDataRef

// CFSocketCopyPeerAddress returns the remote address to which a CFSocket object is connected.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyPeerAddress(_:)
func CFSocketCopyPeerAddress(s CFSocketRef) CFDataRef {
	if _cFSocketCopyPeerAddress == nil {
		panic("CoreFoundation: symbol CFSocketCopyPeerAddress not loaded")
	}
	return _cFSocketCopyPeerAddress(s)
}


var _cFSocketCopyRegisteredSocketSignature func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature, nameServerAddress *CFDataRef) CFSocketError

// CFSocketCopyRegisteredSocketSignature returns a socket signature registered with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredSocketSignature(_:_:_:_:_:)
func CFSocketCopyRegisteredSocketSignature(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature, nameServerAddress *CFDataRef) CFSocketError {
	if _cFSocketCopyRegisteredSocketSignature == nil {
		panic("CoreFoundation: symbol CFSocketCopyRegisteredSocketSignature not loaded")
	}
	return _cFSocketCopyRegisteredSocketSignature(nameServerSignature, timeout, name, signature, nameServerAddress)
}


var _cFSocketCopyRegisteredValue func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value *CFPropertyListRef, nameServerAddress *CFDataRef) CFSocketError

// CFSocketCopyRegisteredValue returns a value registered with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredValue(_:_:_:_:_:)
func CFSocketCopyRegisteredValue(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value *CFPropertyListRef, nameServerAddress *CFDataRef) CFSocketError {
	if _cFSocketCopyRegisteredValue == nil {
		panic("CoreFoundation: symbol CFSocketCopyRegisteredValue not loaded")
	}
	return _cFSocketCopyRegisteredValue(nameServerSignature, timeout, name, value, nameServerAddress)
}


var _cFSocketCreate func(allocator CFAllocatorRef, protocolFamily int32, socketType int32, protocol_ int32, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef

// CFSocketCreate creates a CFSocket object of a specified protocol and type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreate(_:_:_:_:_:_:_:)
func CFSocketCreate(allocator CFAllocatorRef, protocolFamily int32, socketType int32, protocol_ int32, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef {
	if _cFSocketCreate == nil {
		panic("CoreFoundation: symbol CFSocketCreate not loaded")
	}
	return _cFSocketCreate(allocator, protocolFamily, socketType, protocol_, callBackTypes, callout, context)
}


var _cFSocketCreateConnectedToSocketSignature func(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext, timeout float64) CFSocketRef

// CFSocketCreateConnectedToSocketSignature creates a CFSocket object and opens a connection to a remote socket.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateConnectedToSocketSignature(_:_:_:_:_:_:)
func CFSocketCreateConnectedToSocketSignature(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext, timeout float64) CFSocketRef {
	if _cFSocketCreateConnectedToSocketSignature == nil {
		panic("CoreFoundation: symbol CFSocketCreateConnectedToSocketSignature not loaded")
	}
	return _cFSocketCreateConnectedToSocketSignature(allocator, signature, callBackTypes, callout, context, timeout)
}


var _cFSocketCreateRunLoopSource func(allocator CFAllocatorRef, s CFSocketRef, order int) CFRunLoopSourceRef

// CFSocketCreateRunLoopSource creates a CFRunLoopSource object for a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateRunLoopSource(_:_:_:)
func CFSocketCreateRunLoopSource(allocator CFAllocatorRef, s CFSocketRef, order int) CFRunLoopSourceRef {
	if _cFSocketCreateRunLoopSource == nil {
		panic("CoreFoundation: symbol CFSocketCreateRunLoopSource not loaded")
	}
	return _cFSocketCreateRunLoopSource(allocator, s, order)
}


var _cFSocketCreateWithNative func(allocator CFAllocatorRef, sock CFSocketNativeHandle, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef

// CFSocketCreateWithNative creates a CFSocket object for a pre-existing native socket.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithNative(_:_:_:_:_:)
func CFSocketCreateWithNative(allocator CFAllocatorRef, sock CFSocketNativeHandle, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef {
	if _cFSocketCreateWithNative == nil {
		panic("CoreFoundation: symbol CFSocketCreateWithNative not loaded")
	}
	return _cFSocketCreateWithNative(allocator, sock, callBackTypes, callout, context)
}


var _cFSocketCreateWithSocketSignature func(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef

// CFSocketCreateWithSocketSignature creates a CFSocket object using information from a CFSocketSignature structure.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithSocketSignature(_:_:_:_:_:)
func CFSocketCreateWithSocketSignature(allocator CFAllocatorRef, signature *CFSocketSignature, callBackTypes uint64, callout CFSocketCallBack, context *CFSocketContext) CFSocketRef {
	if _cFSocketCreateWithSocketSignature == nil {
		panic("CoreFoundation: symbol CFSocketCreateWithSocketSignature not loaded")
	}
	return _cFSocketCreateWithSocketSignature(allocator, signature, callBackTypes, callout, context)
}


var _cFSocketDisableCallBacks func(s CFSocketRef, callBackTypes uint64) 

// CFSocketDisableCallBacks disables the callback function of a CFSocket object for certain types of socket activity.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketDisableCallBacks(_:_:)
func CFSocketDisableCallBacks(s CFSocketRef, callBackTypes uint64)  {
	if _cFSocketDisableCallBacks == nil {
		panic("CoreFoundation: symbol CFSocketDisableCallBacks not loaded")
	}
	_cFSocketDisableCallBacks(s, callBackTypes)
}


var _cFSocketEnableCallBacks func(s CFSocketRef, callBackTypes uint64) 

// CFSocketEnableCallBacks enables the callback function of a CFSocket object for certain types of socket activity.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketEnableCallBacks(_:_:)
func CFSocketEnableCallBacks(s CFSocketRef, callBackTypes uint64)  {
	if _cFSocketEnableCallBacks == nil {
		panic("CoreFoundation: symbol CFSocketEnableCallBacks not loaded")
	}
	_cFSocketEnableCallBacks(s, callBackTypes)
}


var _cFSocketGetContext func(s CFSocketRef, context *CFSocketContext) 

// CFSocketGetContext returns the context information for a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetContext(_:_:)
func CFSocketGetContext(s CFSocketRef, context *CFSocketContext)  {
	if _cFSocketGetContext == nil {
		panic("CoreFoundation: symbol CFSocketGetContext not loaded")
	}
	_cFSocketGetContext(s, context)
}


var _cFSocketGetDefaultNameRegistryPortNumber func() uint16

// CFSocketGetDefaultNameRegistryPortNumber returns the default port number with which to connect to a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetDefaultNameRegistryPortNumber()
func CFSocketGetDefaultNameRegistryPortNumber() uint16 {
	if _cFSocketGetDefaultNameRegistryPortNumber == nil {
		panic("CoreFoundation: symbol CFSocketGetDefaultNameRegistryPortNumber not loaded")
	}
	return _cFSocketGetDefaultNameRegistryPortNumber()
}


var _cFSocketGetNative func(s CFSocketRef) CFSocketNativeHandle

// CFSocketGetNative returns the native socket associated with a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetNative(_:)
func CFSocketGetNative(s CFSocketRef) CFSocketNativeHandle {
	if _cFSocketGetNative == nil {
		panic("CoreFoundation: symbol CFSocketGetNative not loaded")
	}
	return _cFSocketGetNative(s)
}


var _cFSocketGetSocketFlags func(s CFSocketRef) uint64

// CFSocketGetSocketFlags returns flags that control certain behaviors of a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetSocketFlags(_:)
func CFSocketGetSocketFlags(s CFSocketRef) uint64 {
	if _cFSocketGetSocketFlags == nil {
		panic("CoreFoundation: symbol CFSocketGetSocketFlags not loaded")
	}
	return _cFSocketGetSocketFlags(s)
}


var _cFSocketGetTypeID func() uint

// CFSocketGetTypeID returns the type identifier for the CFSocket opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetTypeID()
func CFSocketGetTypeID() uint {
	if _cFSocketGetTypeID == nil {
		panic("CoreFoundation: symbol CFSocketGetTypeID not loaded")
	}
	return _cFSocketGetTypeID()
}


var _cFSocketInvalidate func(s CFSocketRef) 

// CFSocketInvalidate invalidates a CFSocket object, stopping it from sending or receiving any more messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketInvalidate(_:)
func CFSocketInvalidate(s CFSocketRef)  {
	if _cFSocketInvalidate == nil {
		panic("CoreFoundation: symbol CFSocketInvalidate not loaded")
	}
	_cFSocketInvalidate(s)
}


var _cFSocketIsValid func(s CFSocketRef) bool

// CFSocketIsValid returns a Boolean value that indicates whether a CFSocket object is valid and able to send or receive messages.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketIsValid(_:)
func CFSocketIsValid(s CFSocketRef) bool {
	if _cFSocketIsValid == nil {
		panic("CoreFoundation: symbol CFSocketIsValid not loaded")
	}
	return _cFSocketIsValid(s)
}


var _cFSocketRegisterSocketSignature func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature) CFSocketError

// CFSocketRegisterSocketSignature registers a socket signature with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterSocketSignature(_:_:_:_:)
func CFSocketRegisterSocketSignature(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, signature *CFSocketSignature) CFSocketError {
	if _cFSocketRegisterSocketSignature == nil {
		panic("CoreFoundation: symbol CFSocketRegisterSocketSignature not loaded")
	}
	return _cFSocketRegisterSocketSignature(nameServerSignature, timeout, name, signature)
}


var _cFSocketRegisterValue func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value CFPropertyListRef) CFSocketError

// CFSocketRegisterValue registers a property-list value with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterValue(_:_:_:_:)
func CFSocketRegisterValue(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef, value CFPropertyListRef) CFSocketError {
	if _cFSocketRegisterValue == nil {
		panic("CoreFoundation: symbol CFSocketRegisterValue not loaded")
	}
	return _cFSocketRegisterValue(nameServerSignature, timeout, name, value)
}


var _cFSocketSendData func(s CFSocketRef, address CFDataRef, data CFDataRef, timeout float64) CFSocketError

// CFSocketSendData sends data over a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSendData(_:_:_:_:)
func CFSocketSendData(s CFSocketRef, address CFDataRef, data CFDataRef, timeout float64) CFSocketError {
	if _cFSocketSendData == nil {
		panic("CoreFoundation: symbol CFSocketSendData not loaded")
	}
	return _cFSocketSendData(s, address, data, timeout)
}


var _cFSocketSetAddress func(s CFSocketRef, address CFDataRef) CFSocketError

// CFSocketSetAddress binds a local address to a CFSocket object and configures it for listening.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetAddress(_:_:)
func CFSocketSetAddress(s CFSocketRef, address CFDataRef) CFSocketError {
	if _cFSocketSetAddress == nil {
		panic("CoreFoundation: symbol CFSocketSetAddress not loaded")
	}
	return _cFSocketSetAddress(s, address)
}


var _cFSocketSetDefaultNameRegistryPortNumber func(port uint16) 

// CFSocketSetDefaultNameRegistryPortNumber sets the default port number with which to connect to a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetDefaultNameRegistryPortNumber(_:)
func CFSocketSetDefaultNameRegistryPortNumber(port uint16)  {
	if _cFSocketSetDefaultNameRegistryPortNumber == nil {
		panic("CoreFoundation: symbol CFSocketSetDefaultNameRegistryPortNumber not loaded")
	}
	_cFSocketSetDefaultNameRegistryPortNumber(port)
}


var _cFSocketSetSocketFlags func(s CFSocketRef, flags uint64) 

// CFSocketSetSocketFlags sets flags that control certain behaviors of a CFSocket object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetSocketFlags(_:_:)
func CFSocketSetSocketFlags(s CFSocketRef, flags uint64)  {
	if _cFSocketSetSocketFlags == nil {
		panic("CoreFoundation: symbol CFSocketSetSocketFlags not loaded")
	}
	_cFSocketSetSocketFlags(s, flags)
}


var _cFSocketUnregister func(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef) CFSocketError

// CFSocketUnregister unregisters a value or socket signature with a CFSocket name server.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketUnregister(_:_:_:)
func CFSocketUnregister(nameServerSignature *CFSocketSignature, timeout float64, name CFStringRef) CFSocketError {
	if _cFSocketUnregister == nil {
		panic("CoreFoundation: symbol CFSocketUnregister not loaded")
	}
	return _cFSocketUnregister(nameServerSignature, timeout, name)
}


var _cFStreamCreateBoundPair func(alloc CFAllocatorRef, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef, transferBufferSize int) 

// CFStreamCreateBoundPair creates a bound pair of read and write streams.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreateBoundPair(_:_:_:_:)
func CFStreamCreateBoundPair(alloc CFAllocatorRef, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef, transferBufferSize int)  {
	if _cFStreamCreateBoundPair == nil {
		panic("CoreFoundation: symbol CFStreamCreateBoundPair not loaded")
	}
	_cFStreamCreateBoundPair(alloc, readStream, writeStream, transferBufferSize)
}


var _cFStreamCreatePairWithPeerSocketSignature func(alloc CFAllocatorRef, signature *CFSocketSignature, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) 

// CFStreamCreatePairWithPeerSocketSignature creates readable and writable streams connected to a socket.
//
// Deprecated: Deprecated since macOS 26.2. Use nw_connection_t in Network framework instead
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithPeerSocketSignature(_:_:_:_:)
func CFStreamCreatePairWithPeerSocketSignature(alloc CFAllocatorRef, signature *CFSocketSignature, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef)  {
	if _cFStreamCreatePairWithPeerSocketSignature == nil {
		panic("CoreFoundation: symbol CFStreamCreatePairWithPeerSocketSignature not loaded")
	}
	_cFStreamCreatePairWithPeerSocketSignature(alloc, signature, readStream, writeStream)
}


var _cFStreamCreatePairWithSocket func(alloc CFAllocatorRef, sock CFSocketNativeHandle, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) 

// CFStreamCreatePairWithSocket creates readable and writable streams connected to a socket.
//
// Deprecated: Deprecated since macOS 26.2. Use nw_connection_t in Network framework instead
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocket(_:_:_:_:)
func CFStreamCreatePairWithSocket(alloc CFAllocatorRef, sock CFSocketNativeHandle, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef)  {
	if _cFStreamCreatePairWithSocket == nil {
		panic("CoreFoundation: symbol CFStreamCreatePairWithSocket not loaded")
	}
	_cFStreamCreatePairWithSocket(alloc, sock, readStream, writeStream)
}


var _cFStreamCreatePairWithSocketToHost func(alloc CFAllocatorRef, host CFStringRef, port uint32, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef) 

// CFStreamCreatePairWithSocketToHost creates readable and writable streams connected to a TCP/IP port of a particular host.
//
// Deprecated: Deprecated since macOS 26.2. Use nw_connection_t in Network framework instead
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocketToHost(_:_:_:_:_:)
func CFStreamCreatePairWithSocketToHost(alloc CFAllocatorRef, host CFStringRef, port uint32, readStream *CFReadStreamRef, writeStream *CFWriteStreamRef)  {
	if _cFStreamCreatePairWithSocketToHost == nil {
		panic("CoreFoundation: symbol CFStreamCreatePairWithSocketToHost not loaded")
	}
	_cFStreamCreatePairWithSocketToHost(alloc, host, port, readStream, writeStream)
}


var _cFStringAppend func(theString CFMutableStringRef, appendedString CFStringRef) 

// CFStringAppend appends the characters of a string to those of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppend(_:_:)
func CFStringAppend(theString CFMutableStringRef, appendedString CFStringRef)  {
	if _cFStringAppend == nil {
		panic("CoreFoundation: symbol CFStringAppend not loaded")
	}
	_cFStringAppend(theString, appendedString)
}


var _cFStringAppendCString func(theString CFMutableStringRef, cStr *byte, encoding uint32) 

// CFStringAppendCString appends a C string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCString(_:_:_:)
func CFStringAppendCString(theString CFMutableStringRef, cStr *byte, encoding uint32)  {
	if _cFStringAppendCString == nil {
		panic("CoreFoundation: symbol CFStringAppendCString not loaded")
	}
	_cFStringAppendCString(theString, cStr, encoding)
}


var _cFStringAppendCharacters func(theString CFMutableStringRef, chars *uint16, numChars int) 

// CFStringAppendCharacters appends a buffer of Unicode characters to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCharacters(_:_:_:)
func CFStringAppendCharacters(theString CFMutableStringRef, chars *uint16, numChars int)  {
	if _cFStringAppendCharacters == nil {
		panic("CoreFoundation: symbol CFStringAppendCharacters not loaded")
	}
	_cFStringAppendCharacters(theString, chars, numChars)
}


var _cFStringAppendFormat func(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef) 

// CFStringAppendFormat appends a formatted string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormat
func CFStringAppendFormat(theString CFMutableStringRef, formatOptions CFDictionaryRef, format CFStringRef)  {
	if _cFStringAppendFormat == nil {
		panic("CoreFoundation: symbol CFStringAppendFormat not loaded")
	}
	_cFStringAppendFormat(theString, formatOptions, format)
}



var _cFStringAppendPascalString func(theString CFMutableStringRef, pStr unsafe.Pointer, encoding uint32) 

// CFStringAppendPascalString appends a Pascal string to the character contents of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendPascalString(_:_:_:)
func CFStringAppendPascalString(theString CFMutableStringRef, pStr unsafe.Pointer, encoding uint32)  {
	if _cFStringAppendPascalString == nil {
		panic("CoreFoundation: symbol CFStringAppendPascalString not loaded")
	}
	_cFStringAppendPascalString(theString, pStr, encoding)
}


var _cFStringCapitalize func(theString CFMutableStringRef, locale CFLocaleRef) 

// CFStringCapitalize changes the first character in each word of a string to uppercase (if it is a lowercase alphabetical character).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCapitalize(_:_:)
func CFStringCapitalize(theString CFMutableStringRef, locale CFLocaleRef)  {
	if _cFStringCapitalize == nil {
		panic("CoreFoundation: symbol CFStringCapitalize not loaded")
	}
	_cFStringCapitalize(theString, locale)
}


var _cFStringCompare func(theString1 CFStringRef, theString2 CFStringRef, compareOptions CFStringCompareFlags) CFComparisonResult

// CFStringCompare compares one string with another string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompare(_:_:_:)
func CFStringCompare(theString1 CFStringRef, theString2 CFStringRef, compareOptions CFStringCompareFlags) CFComparisonResult {
	if _cFStringCompare == nil {
		panic("CoreFoundation: symbol CFStringCompare not loaded")
	}
	return _cFStringCompare(theString1, theString2, compareOptions)
}


var _cFStringCompareWithOptions func(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags) CFComparisonResult

// CFStringCompareWithOptions compares a range of the characters in one string with that of another string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptions(_:_:_:_:)
func CFStringCompareWithOptions(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags) CFComparisonResult {
	if _cFStringCompareWithOptions == nil {
		panic("CoreFoundation: symbol CFStringCompareWithOptions not loaded")
	}
	return _cFStringCompareWithOptions(theString1, theString2, rangeToCompare, compareOptions)
}


var _cFStringCompareWithOptionsAndLocale func(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags, locale CFLocaleRef) CFComparisonResult

// CFStringCompareWithOptionsAndLocale compares a range of the characters in one string with another string using a given locale.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptionsAndLocale(_:_:_:_:_:)
func CFStringCompareWithOptionsAndLocale(theString1 CFStringRef, theString2 CFStringRef, rangeToCompare CFRange, compareOptions CFStringCompareFlags, locale CFLocaleRef) CFComparisonResult {
	if _cFStringCompareWithOptionsAndLocale == nil {
		panic("CoreFoundation: symbol CFStringCompareWithOptionsAndLocale not loaded")
	}
	return _cFStringCompareWithOptionsAndLocale(theString1, theString2, rangeToCompare, compareOptions, locale)
}


var _cFStringConvertEncodingToIANACharSetName func(encoding uint32) CFStringRef

// CFStringConvertEncodingToIANACharSetName returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToIANACharSetName(_:)
func CFStringConvertEncodingToIANACharSetName(encoding uint32) CFStringRef {
	if _cFStringConvertEncodingToIANACharSetName == nil {
		panic("CoreFoundation: symbol CFStringConvertEncodingToIANACharSetName not loaded")
	}
	return _cFStringConvertEncodingToIANACharSetName(encoding)
}


var _cFStringConvertEncodingToNSStringEncoding func(encoding uint32) uint

// CFStringConvertEncodingToNSStringEncoding returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
func CFStringConvertEncodingToNSStringEncoding(encoding uint32) uint {
	if _cFStringConvertEncodingToNSStringEncoding == nil {
		panic("CoreFoundation: symbol CFStringConvertEncodingToNSStringEncoding not loaded")
	}
	return _cFStringConvertEncodingToNSStringEncoding(encoding)
}


var _cFStringConvertEncodingToWindowsCodepage func(encoding uint32) uint32

// CFStringConvertEncodingToWindowsCodepage returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToWindowsCodepage(_:)
func CFStringConvertEncodingToWindowsCodepage(encoding uint32) uint32 {
	if _cFStringConvertEncodingToWindowsCodepage == nil {
		panic("CoreFoundation: symbol CFStringConvertEncodingToWindowsCodepage not loaded")
	}
	return _cFStringConvertEncodingToWindowsCodepage(encoding)
}


var _cFStringConvertIANACharSetNameToEncoding func(theString CFStringRef) uint32

// CFStringConvertIANACharSetNameToEncoding returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertIANACharSetNameToEncoding(_:)
func CFStringConvertIANACharSetNameToEncoding(theString CFStringRef) uint32 {
	if _cFStringConvertIANACharSetNameToEncoding == nil {
		panic("CoreFoundation: symbol CFStringConvertIANACharSetNameToEncoding not loaded")
	}
	return _cFStringConvertIANACharSetNameToEncoding(theString)
}


var _cFStringConvertNSStringEncodingToEncoding func(encoding uint) uint32

// CFStringConvertNSStringEncodingToEncoding returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertNSStringEncodingToEncoding(_:)
func CFStringConvertNSStringEncodingToEncoding(encoding uint) uint32 {
	if _cFStringConvertNSStringEncodingToEncoding == nil {
		panic("CoreFoundation: symbol CFStringConvertNSStringEncodingToEncoding not loaded")
	}
	return _cFStringConvertNSStringEncodingToEncoding(encoding)
}


var _cFStringConvertWindowsCodepageToEncoding func(codepage uint32) uint32

// CFStringConvertWindowsCodepageToEncoding returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertWindowsCodepageToEncoding(_:)
func CFStringConvertWindowsCodepageToEncoding(codepage uint32) uint32 {
	if _cFStringConvertWindowsCodepageToEncoding == nil {
		panic("CoreFoundation: symbol CFStringConvertWindowsCodepageToEncoding not loaded")
	}
	return _cFStringConvertWindowsCodepageToEncoding(codepage)
}


var _cFStringCreateArrayBySeparatingStrings func(alloc CFAllocatorRef, theString CFStringRef, separatorString CFStringRef) CFArrayRef

// CFStringCreateArrayBySeparatingStrings creates an array of CFString objects from a single CFString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayBySeparatingStrings(_:_:_:)
func CFStringCreateArrayBySeparatingStrings(alloc CFAllocatorRef, theString CFStringRef, separatorString CFStringRef) CFArrayRef {
	if _cFStringCreateArrayBySeparatingStrings == nil {
		panic("CoreFoundation: symbol CFStringCreateArrayBySeparatingStrings not loaded")
	}
	return _cFStringCreateArrayBySeparatingStrings(alloc, theString, separatorString)
}


var _cFStringCreateArrayWithFindResults func(alloc CFAllocatorRef, theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) CFArrayRef

// CFStringCreateArrayWithFindResults searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayWithFindResults(_:_:_:_:_:)
func CFStringCreateArrayWithFindResults(alloc CFAllocatorRef, theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) CFArrayRef {
	if _cFStringCreateArrayWithFindResults == nil {
		panic("CoreFoundation: symbol CFStringCreateArrayWithFindResults not loaded")
	}
	return _cFStringCreateArrayWithFindResults(alloc, theString, stringToFind, rangeToSearch, compareOptions)
}


var _cFStringCreateByCombiningStrings func(alloc CFAllocatorRef, theArray CFArrayRef, separatorString CFStringRef) CFStringRef

// CFStringCreateByCombiningStrings creates a single string from the individual CFString objects that comprise the elements of an array.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateByCombiningStrings(_:_:_:)
func CFStringCreateByCombiningStrings(alloc CFAllocatorRef, theArray CFArrayRef, separatorString CFStringRef) CFStringRef {
	if _cFStringCreateByCombiningStrings == nil {
		panic("CoreFoundation: symbol CFStringCreateByCombiningStrings not loaded")
	}
	return _cFStringCreateByCombiningStrings(alloc, theArray, separatorString)
}


var _cFStringCreateCopy func(alloc CFAllocatorRef, theString CFStringRef) CFStringRef

// CFStringCreateCopy creates an immutable copy of a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateCopy(_:_:)
func CFStringCreateCopy(alloc CFAllocatorRef, theString CFStringRef) CFStringRef {
	if _cFStringCreateCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateCopy not loaded")
	}
	return _cFStringCreateCopy(alloc, theString)
}


var _cFStringCreateExternalRepresentation func(alloc CFAllocatorRef, theString CFStringRef, encoding uint32, lossByte uint8) CFDataRef

// CFStringCreateExternalRepresentation creates an “external representation” of a CFString object, that is, a CFData object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateExternalRepresentation(_:_:_:_:)
func CFStringCreateExternalRepresentation(alloc CFAllocatorRef, theString CFStringRef, encoding uint32, lossByte uint8) CFDataRef {
	if _cFStringCreateExternalRepresentation == nil {
		panic("CoreFoundation: symbol CFStringCreateExternalRepresentation not loaded")
	}
	return _cFStringCreateExternalRepresentation(alloc, theString, encoding, lossByte)
}


var _cFStringCreateFromExternalRepresentation func(alloc CFAllocatorRef, data CFDataRef, encoding uint32) CFStringRef

// CFStringCreateFromExternalRepresentation creates a string from its “external representation.”
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateFromExternalRepresentation(_:_:_:)
func CFStringCreateFromExternalRepresentation(alloc CFAllocatorRef, data CFDataRef, encoding uint32) CFStringRef {
	if _cFStringCreateFromExternalRepresentation == nil {
		panic("CoreFoundation: symbol CFStringCreateFromExternalRepresentation not loaded")
	}
	return _cFStringCreateFromExternalRepresentation(alloc, data, encoding)
}


var _cFStringCreateMutable func(alloc CFAllocatorRef, maxLength int) CFMutableStringRef

// CFStringCreateMutable creates an empty CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutable(_:_:)
func CFStringCreateMutable(alloc CFAllocatorRef, maxLength int) CFMutableStringRef {
	if _cFStringCreateMutable == nil {
		panic("CoreFoundation: symbol CFStringCreateMutable not loaded")
	}
	return _cFStringCreateMutable(alloc, maxLength)
}


var _cFStringCreateMutableCopy func(alloc CFAllocatorRef, maxLength int, theString CFStringRef) CFMutableStringRef

// CFStringCreateMutableCopy creates a mutable copy of a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableCopy(_:_:_:)
func CFStringCreateMutableCopy(alloc CFAllocatorRef, maxLength int, theString CFStringRef) CFMutableStringRef {
	if _cFStringCreateMutableCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateMutableCopy not loaded")
	}
	return _cFStringCreateMutableCopy(alloc, maxLength, theString)
}


var _cFStringCreateMutableWithExternalCharactersNoCopy func(alloc CFAllocatorRef, chars *uint16, numChars int, capacity int, externalCharactersAllocator CFAllocatorRef) CFMutableStringRef

// CFStringCreateMutableWithExternalCharactersNoCopy creates a CFMutableString object whose Unicode character buffer is controlled externally.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableWithExternalCharactersNoCopy(_:_:_:_:_:)
func CFStringCreateMutableWithExternalCharactersNoCopy(alloc CFAllocatorRef, chars *uint16, numChars int, capacity int, externalCharactersAllocator CFAllocatorRef) CFMutableStringRef {
	if _cFStringCreateMutableWithExternalCharactersNoCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateMutableWithExternalCharactersNoCopy not loaded")
	}
	return _cFStringCreateMutableWithExternalCharactersNoCopy(alloc, chars, numChars, capacity, externalCharactersAllocator)
}


var _cFStringCreateStringWithValidatedFormat func(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, errorPtr *CFErrorRef) CFStringRef

// CFStringCreateStringWithValidatedFormat.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormat
func CFStringCreateStringWithValidatedFormat(alloc CFAllocatorRef, formatOptions CFDictionaryRef, validFormatSpecifiers CFStringRef, format CFStringRef, errorPtr *CFErrorRef) CFStringRef {
	if _cFStringCreateStringWithValidatedFormat == nil {
		panic("CoreFoundation: symbol CFStringCreateStringWithValidatedFormat not loaded")
	}
	return _cFStringCreateStringWithValidatedFormat(alloc, formatOptions, validFormatSpecifiers, format, errorPtr)
}



var _cFStringCreateWithBytes func(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool) CFStringRef

// CFStringCreateWithBytes creates a string from a buffer containing characters in a specified encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytes(_:_:_:_:_:)
func CFStringCreateWithBytes(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool) CFStringRef {
	if _cFStringCreateWithBytes == nil {
		panic("CoreFoundation: symbol CFStringCreateWithBytes not loaded")
	}
	return _cFStringCreateWithBytes(alloc, bytes, numBytes, encoding, isExternalRepresentation)
}


var _cFStringCreateWithBytesNoCopy func(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool, contentsDeallocator CFAllocatorRef) CFStringRef

// CFStringCreateWithBytesNoCopy creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytesNoCopy(_:_:_:_:_:_:)
func CFStringCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes *uint8, numBytes int, encoding uint32, isExternalRepresentation bool, contentsDeallocator CFAllocatorRef) CFStringRef {
	if _cFStringCreateWithBytesNoCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateWithBytesNoCopy not loaded")
	}
	return _cFStringCreateWithBytesNoCopy(alloc, bytes, numBytes, encoding, isExternalRepresentation, contentsDeallocator)
}


var _cFStringCreateWithCString func(alloc CFAllocatorRef, cStr *byte, encoding uint32) CFStringRef

// CFStringCreateWithCString creates an immutable string from a C string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCString(_:_:_:)
func CFStringCreateWithCString(alloc CFAllocatorRef, cStr *byte, encoding uint32) CFStringRef {
	if _cFStringCreateWithCString == nil {
		panic("CoreFoundation: symbol CFStringCreateWithCString not loaded")
	}
	return _cFStringCreateWithCString(alloc, cStr, encoding)
}


var _cFStringCreateWithCStringNoCopy func(alloc CFAllocatorRef, cStr *byte, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef

// CFStringCreateWithCStringNoCopy creates a CFString object from an external C string buffer that might serve as the backing store for the object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCStringNoCopy(_:_:_:_:)
func CFStringCreateWithCStringNoCopy(alloc CFAllocatorRef, cStr *byte, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef {
	if _cFStringCreateWithCStringNoCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateWithCStringNoCopy not loaded")
	}
	return _cFStringCreateWithCStringNoCopy(alloc, cStr, encoding, contentsDeallocator)
}


var _cFStringCreateWithCharacters func(alloc CFAllocatorRef, chars *uint16, numChars int) CFStringRef

// CFStringCreateWithCharacters creates a string from a buffer of Unicode characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharacters(_:_:_:)
func CFStringCreateWithCharacters(alloc CFAllocatorRef, chars *uint16, numChars int) CFStringRef {
	if _cFStringCreateWithCharacters == nil {
		panic("CoreFoundation: symbol CFStringCreateWithCharacters not loaded")
	}
	return _cFStringCreateWithCharacters(alloc, chars, numChars)
}


var _cFStringCreateWithCharactersNoCopy func(alloc CFAllocatorRef, chars *uint16, numChars int, contentsDeallocator CFAllocatorRef) CFStringRef

// CFStringCreateWithCharactersNoCopy creates a string from a buffer of Unicode characters that might serve as the backing store for the object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharactersNoCopy(_:_:_:_:)
func CFStringCreateWithCharactersNoCopy(alloc CFAllocatorRef, chars *uint16, numChars int, contentsDeallocator CFAllocatorRef) CFStringRef {
	if _cFStringCreateWithCharactersNoCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateWithCharactersNoCopy not loaded")
	}
	return _cFStringCreateWithCharactersNoCopy(alloc, chars, numChars, contentsDeallocator)
}


var _cFStringCreateWithFileSystemRepresentation func(alloc CFAllocatorRef, buffer *byte) CFStringRef

// CFStringCreateWithFileSystemRepresentation creates a CFString from a zero-terminated POSIX file system representation.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFileSystemRepresentation(_:_:)
func CFStringCreateWithFileSystemRepresentation(alloc CFAllocatorRef, buffer *byte) CFStringRef {
	if _cFStringCreateWithFileSystemRepresentation == nil {
		panic("CoreFoundation: symbol CFStringCreateWithFileSystemRepresentation not loaded")
	}
	return _cFStringCreateWithFileSystemRepresentation(alloc, buffer)
}


var _cFStringCreateWithFormat func(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef) CFStringRef

// CFStringCreateWithFormat creates an immutable string from a formatted string and a variable number of arguments.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormat
func CFStringCreateWithFormat(alloc CFAllocatorRef, formatOptions CFDictionaryRef, format CFStringRef) CFStringRef {
	if _cFStringCreateWithFormat == nil {
		panic("CoreFoundation: symbol CFStringCreateWithFormat not loaded")
	}
	return _cFStringCreateWithFormat(alloc, formatOptions, format)
}



var _cFStringCreateWithPascalString func(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32) CFStringRef

// CFStringCreateWithPascalString creates an immutable CFString object from a Pascal string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalString(_:_:_:)
func CFStringCreateWithPascalString(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32) CFStringRef {
	if _cFStringCreateWithPascalString == nil {
		panic("CoreFoundation: symbol CFStringCreateWithPascalString not loaded")
	}
	return _cFStringCreateWithPascalString(alloc, pStr, encoding)
}


var _cFStringCreateWithPascalStringNoCopy func(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef

// CFStringCreateWithPascalStringNoCopy creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalStringNoCopy(_:_:_:_:)
func CFStringCreateWithPascalStringNoCopy(alloc CFAllocatorRef, pStr unsafe.Pointer, encoding uint32, contentsDeallocator CFAllocatorRef) CFStringRef {
	if _cFStringCreateWithPascalStringNoCopy == nil {
		panic("CoreFoundation: symbol CFStringCreateWithPascalStringNoCopy not loaded")
	}
	return _cFStringCreateWithPascalStringNoCopy(alloc, pStr, encoding, contentsDeallocator)
}


var _cFStringCreateWithSubstring func(alloc CFAllocatorRef, str CFStringRef, range_ CFRange) CFStringRef

// CFStringCreateWithSubstring creates an immutable string from a segment (substring) of an existing string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithSubstring(_:_:_:)
func CFStringCreateWithSubstring(alloc CFAllocatorRef, str CFStringRef, range_ CFRange) CFStringRef {
	if _cFStringCreateWithSubstring == nil {
		panic("CoreFoundation: symbol CFStringCreateWithSubstring not loaded")
	}
	return _cFStringCreateWithSubstring(alloc, str, range_)
}


var _cFStringDelete func(theString CFMutableStringRef, range_ CFRange) 

// CFStringDelete deletes a range of characters in a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringDelete(_:_:)
func CFStringDelete(theString CFMutableStringRef, range_ CFRange)  {
	if _cFStringDelete == nil {
		panic("CoreFoundation: symbol CFStringDelete not loaded")
	}
	_cFStringDelete(theString, range_)
}


var _cFStringFind func(theString CFStringRef, stringToFind CFStringRef, compareOptions CFStringCompareFlags) CFRange

// CFStringFind searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFind(_:_:_:)
func CFStringFind(theString CFStringRef, stringToFind CFStringRef, compareOptions CFStringCompareFlags) CFRange {
	if _cFStringFind == nil {
		panic("CoreFoundation: symbol CFStringFind not loaded")
	}
	return _cFStringFind(theString, stringToFind, compareOptions)
}


var _cFStringFindAndReplace func(theString CFMutableStringRef, stringToFind CFStringRef, replacementString CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) int

// CFStringFindAndReplace replaces all occurrences of a substring within a given range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindAndReplace(_:_:_:_:_:)
func CFStringFindAndReplace(theString CFMutableStringRef, stringToFind CFStringRef, replacementString CFStringRef, rangeToSearch CFRange, compareOptions CFStringCompareFlags) int {
	if _cFStringFindAndReplace == nil {
		panic("CoreFoundation: symbol CFStringFindAndReplace not loaded")
	}
	return _cFStringFindAndReplace(theString, stringToFind, replacementString, rangeToSearch, compareOptions)
}


var _cFStringFindCharacterFromSet func(theString CFStringRef, theSet CFCharacterSetRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool

// CFStringFindCharacterFromSet query the range of the first character contained in the specified character set.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindCharacterFromSet(_:_:_:_:_:)
func CFStringFindCharacterFromSet(theString CFStringRef, theSet CFCharacterSetRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool {
	if _cFStringFindCharacterFromSet == nil {
		panic("CoreFoundation: symbol CFStringFindCharacterFromSet not loaded")
	}
	return _cFStringFindCharacterFromSet(theString, theSet, rangeToSearch, searchOptions, result)
}


var _cFStringFindWithOptions func(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool

// CFStringFindWithOptions searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptions(_:_:_:_:_:)
func CFStringFindWithOptions(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, result *CFRange) bool {
	if _cFStringFindWithOptions == nil {
		panic("CoreFoundation: symbol CFStringFindWithOptions not loaded")
	}
	return _cFStringFindWithOptions(theString, stringToFind, rangeToSearch, searchOptions, result)
}


var _cFStringFindWithOptionsAndLocale func(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, locale CFLocaleRef, result *CFRange) bool

// CFStringFindWithOptionsAndLocale returns a Boolean value that indicates whether a given string was found in a given source string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptionsAndLocale(_:_:_:_:_:_:)
func CFStringFindWithOptionsAndLocale(theString CFStringRef, stringToFind CFStringRef, rangeToSearch CFRange, searchOptions CFStringCompareFlags, locale CFLocaleRef, result *CFRange) bool {
	if _cFStringFindWithOptionsAndLocale == nil {
		panic("CoreFoundation: symbol CFStringFindWithOptionsAndLocale not loaded")
	}
	return _cFStringFindWithOptionsAndLocale(theString, stringToFind, rangeToSearch, searchOptions, locale, result)
}


var _cFStringFold func(theString CFMutableStringRef, theFlags CFStringCompareFlags, theLocale CFLocaleRef) 

// CFStringFold folds a given string into the form specified by optional flags.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringFold(_:_:_:)
func CFStringFold(theString CFMutableStringRef, theFlags CFStringCompareFlags, theLocale CFLocaleRef)  {
	if _cFStringFold == nil {
		panic("CoreFoundation: symbol CFStringFold not loaded")
	}
	_cFStringFold(theString, theFlags, theLocale)
}


var _cFStringGetBytes func(theString CFStringRef, range_ CFRange, encoding uint32, lossByte uint8, isExternalRepresentation bool, buffer *uint8, maxBufLen int, usedBufLen *int) int

// CFStringGetBytes fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetBytes(_:_:_:_:_:_:_:_:)
func CFStringGetBytes(theString CFStringRef, range_ CFRange, encoding uint32, lossByte uint8, isExternalRepresentation bool, buffer *uint8, maxBufLen int, usedBufLen *int) int {
	if _cFStringGetBytes == nil {
		panic("CoreFoundation: symbol CFStringGetBytes not loaded")
	}
	return _cFStringGetBytes(theString, range_, encoding, lossByte, isExternalRepresentation, buffer, maxBufLen, usedBufLen)
}


var _cFStringGetCString func(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool

// CFStringGetCString copies the character contents of a string to a local C string buffer after converting the characters to a given encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCString(_:_:_:_:)
func CFStringGetCString(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool {
	if _cFStringGetCString == nil {
		panic("CoreFoundation: symbol CFStringGetCString not loaded")
	}
	return _cFStringGetCString(theString, buffer, bufferSize, encoding)
}


var _cFStringGetCStringPtr func(theString CFStringRef, encoding uint32) *byte

// CFStringGetCStringPtr quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCStringPtr(_:_:)
func CFStringGetCStringPtr(theString CFStringRef, encoding uint32) *byte {
	if _cFStringGetCStringPtr == nil {
		panic("CoreFoundation: symbol CFStringGetCStringPtr not loaded")
	}
	return _cFStringGetCStringPtr(theString, encoding)
}


var _cFStringGetCharacterAtIndex func(theString CFStringRef, idx int) uint16

// CFStringGetCharacterAtIndex returns the Unicode character at a specified location in a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacterAtIndex(_:_:)
func CFStringGetCharacterAtIndex(theString CFStringRef, idx int) uint16 {
	if _cFStringGetCharacterAtIndex == nil {
		panic("CoreFoundation: symbol CFStringGetCharacterAtIndex not loaded")
	}
	return _cFStringGetCharacterAtIndex(theString, idx)
}



var _cFStringGetCharacters func(theString CFStringRef, range_ CFRange, buffer *uint16) 

// CFStringGetCharacters copies a range of the Unicode characters from a string to a user-provided buffer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacters(_:_:_:)
func CFStringGetCharacters(theString CFStringRef, range_ CFRange, buffer *uint16)  {
	if _cFStringGetCharacters == nil {
		panic("CoreFoundation: symbol CFStringGetCharacters not loaded")
	}
	_cFStringGetCharacters(theString, range_, buffer)
}


var _cFStringGetCharactersPtr func(theString CFStringRef) *uint16

// CFStringGetCharactersPtr quickly obtains a pointer to the contents of a string as a buffer of Unicode characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharactersPtr(_:)
func CFStringGetCharactersPtr(theString CFStringRef) *uint16 {
	if _cFStringGetCharactersPtr == nil {
		panic("CoreFoundation: symbol CFStringGetCharactersPtr not loaded")
	}
	return _cFStringGetCharactersPtr(theString)
}


var _cFStringGetDoubleValue func(str CFStringRef) float64

// CFStringGetDoubleValue returns the primary `double` value represented by a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetDoubleValue(_:)
func CFStringGetDoubleValue(str CFStringRef) float64 {
	if _cFStringGetDoubleValue == nil {
		panic("CoreFoundation: symbol CFStringGetDoubleValue not loaded")
	}
	return _cFStringGetDoubleValue(str)
}


var _cFStringGetFastestEncoding func(theString CFStringRef) uint32

// CFStringGetFastestEncoding returns for a CFString object the character encoding that requires the least conversion time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFastestEncoding(_:)
func CFStringGetFastestEncoding(theString CFStringRef) uint32 {
	if _cFStringGetFastestEncoding == nil {
		panic("CoreFoundation: symbol CFStringGetFastestEncoding not loaded")
	}
	return _cFStringGetFastestEncoding(theString)
}


var _cFStringGetFileSystemRepresentation func(string_ CFStringRef, buffer *byte, maxBufLen int) bool

// CFStringGetFileSystemRepresentation extracts the contents of a string as a [NULL]-terminated 8-bit string appropriate for passing to POSIX APIs.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFileSystemRepresentation(_:_:_:)
func CFStringGetFileSystemRepresentation(string_ CFStringRef, buffer *byte, maxBufLen int) bool {
	if _cFStringGetFileSystemRepresentation == nil {
		panic("CoreFoundation: symbol CFStringGetFileSystemRepresentation not loaded")
	}
	return _cFStringGetFileSystemRepresentation(string_, buffer, maxBufLen)
}


var _cFStringGetHyphenationLocationBeforeIndex func(string_ CFStringRef, location int, limitRange CFRange, options uint64, locale CFLocaleRef, character *uint32) int

// CFStringGetHyphenationLocationBeforeIndex retrieve the first potential hyphenation location found before the specified location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetHyphenationLocationBeforeIndex(_:_:_:_:_:_:)
func CFStringGetHyphenationLocationBeforeIndex(string_ CFStringRef, location int, limitRange CFRange, options uint64, locale CFLocaleRef, character *uint32) int {
	if _cFStringGetHyphenationLocationBeforeIndex == nil {
		panic("CoreFoundation: symbol CFStringGetHyphenationLocationBeforeIndex not loaded")
	}
	return _cFStringGetHyphenationLocationBeforeIndex(string_, location, limitRange, options, locale, character)
}


var _cFStringGetIntValue func(str CFStringRef) int32

// CFStringGetIntValue returns the integer value represented by a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetIntValue(_:)
func CFStringGetIntValue(str CFStringRef) int32 {
	if _cFStringGetIntValue == nil {
		panic("CoreFoundation: symbol CFStringGetIntValue not loaded")
	}
	return _cFStringGetIntValue(str)
}


var _cFStringGetLength func(theString CFStringRef) int

// CFStringGetLength returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLength(_:)
func CFStringGetLength(theString CFStringRef) int {
	if _cFStringGetLength == nil {
		panic("CoreFoundation: symbol CFStringGetLength not loaded")
	}
	return _cFStringGetLength(theString)
}


var _cFStringGetLineBounds func(theString CFStringRef, range_ CFRange, lineBeginIndex *int, lineEndIndex *int, contentsEndIndex *int) 

// CFStringGetLineBounds given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLineBounds(_:_:_:_:_:)
func CFStringGetLineBounds(theString CFStringRef, range_ CFRange, lineBeginIndex *int, lineEndIndex *int, contentsEndIndex *int)  {
	if _cFStringGetLineBounds == nil {
		panic("CoreFoundation: symbol CFStringGetLineBounds not loaded")
	}
	_cFStringGetLineBounds(theString, range_, lineBeginIndex, lineEndIndex, contentsEndIndex)
}


var _cFStringGetListOfAvailableEncodings func() *uint32

// CFStringGetListOfAvailableEncodings returns a pointer to a list of string encodings supported by the current system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetListOfAvailableEncodings()
func CFStringGetListOfAvailableEncodings() *uint32 {
	if _cFStringGetListOfAvailableEncodings == nil {
		panic("CoreFoundation: symbol CFStringGetListOfAvailableEncodings not loaded")
	}
	return _cFStringGetListOfAvailableEncodings()
}



var _cFStringGetMaximumSizeForEncoding func(length int, encoding uint32) int

// CFStringGetMaximumSizeForEncoding returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeForEncoding(_:_:)
func CFStringGetMaximumSizeForEncoding(length int, encoding uint32) int {
	if _cFStringGetMaximumSizeForEncoding == nil {
		panic("CoreFoundation: symbol CFStringGetMaximumSizeForEncoding not loaded")
	}
	return _cFStringGetMaximumSizeForEncoding(length, encoding)
}


var _cFStringGetMaximumSizeOfFileSystemRepresentation func(string_ CFStringRef) int

// CFStringGetMaximumSizeOfFileSystemRepresentation determines the upper bound on the number of bytes required to hold the file system representation of the string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeOfFileSystemRepresentation(_:)
func CFStringGetMaximumSizeOfFileSystemRepresentation(string_ CFStringRef) int {
	if _cFStringGetMaximumSizeOfFileSystemRepresentation == nil {
		panic("CoreFoundation: symbol CFStringGetMaximumSizeOfFileSystemRepresentation not loaded")
	}
	return _cFStringGetMaximumSizeOfFileSystemRepresentation(string_)
}


var _cFStringGetMostCompatibleMacStringEncoding func(encoding uint32) uint32

// CFStringGetMostCompatibleMacStringEncoding returns the most compatible Mac OS script value for the given input encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMostCompatibleMacStringEncoding(_:)
func CFStringGetMostCompatibleMacStringEncoding(encoding uint32) uint32 {
	if _cFStringGetMostCompatibleMacStringEncoding == nil {
		panic("CoreFoundation: symbol CFStringGetMostCompatibleMacStringEncoding not loaded")
	}
	return _cFStringGetMostCompatibleMacStringEncoding(encoding)
}


var _cFStringGetNameOfEncoding func(encoding uint32) CFStringRef

// CFStringGetNameOfEncoding returns the canonical name of a specified string encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetNameOfEncoding(_:)
func CFStringGetNameOfEncoding(encoding uint32) CFStringRef {
	if _cFStringGetNameOfEncoding == nil {
		panic("CoreFoundation: symbol CFStringGetNameOfEncoding not loaded")
	}
	return _cFStringGetNameOfEncoding(encoding)
}


var _cFStringGetParagraphBounds func(string_ CFStringRef, range_ CFRange, parBeginIndex *int, parEndIndex *int, contentsEndIndex *int) 

// CFStringGetParagraphBounds given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetParagraphBounds(_:_:_:_:_:)
func CFStringGetParagraphBounds(string_ CFStringRef, range_ CFRange, parBeginIndex *int, parEndIndex *int, contentsEndIndex *int)  {
	if _cFStringGetParagraphBounds == nil {
		panic("CoreFoundation: symbol CFStringGetParagraphBounds not loaded")
	}
	_cFStringGetParagraphBounds(string_, range_, parBeginIndex, parEndIndex, contentsEndIndex)
}


var _cFStringGetPascalString func(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool

// CFStringGetPascalString copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalString(_:_:_:_:)
func CFStringGetPascalString(theString CFStringRef, buffer *byte, bufferSize int, encoding uint32) bool {
	if _cFStringGetPascalString == nil {
		panic("CoreFoundation: symbol CFStringGetPascalString not loaded")
	}
	return _cFStringGetPascalString(theString, buffer, bufferSize, encoding)
}


var _cFStringGetPascalStringPtr func(theString CFStringRef, encoding uint32) *byte

// CFStringGetPascalStringPtr quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalStringPtr(_:_:)
func CFStringGetPascalStringPtr(theString CFStringRef, encoding uint32) *byte {
	if _cFStringGetPascalStringPtr == nil {
		panic("CoreFoundation: symbol CFStringGetPascalStringPtr not loaded")
	}
	return _cFStringGetPascalStringPtr(theString, encoding)
}


var _cFStringGetRangeOfComposedCharactersAtIndex func(theString CFStringRef, theIndex int) CFRange

// CFStringGetRangeOfComposedCharactersAtIndex returns the range of the composed character sequence at a specified index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetRangeOfComposedCharactersAtIndex(_:_:)
func CFStringGetRangeOfComposedCharactersAtIndex(theString CFStringRef, theIndex int) CFRange {
	if _cFStringGetRangeOfComposedCharactersAtIndex == nil {
		panic("CoreFoundation: symbol CFStringGetRangeOfComposedCharactersAtIndex not loaded")
	}
	return _cFStringGetRangeOfComposedCharactersAtIndex(theString, theIndex)
}


var _cFStringGetSmallestEncoding func(theString CFStringRef) uint32

// CFStringGetSmallestEncoding returns the smallest encoding on the current system for the character contents of a string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSmallestEncoding(_:)
func CFStringGetSmallestEncoding(theString CFStringRef) uint32 {
	if _cFStringGetSmallestEncoding == nil {
		panic("CoreFoundation: symbol CFStringGetSmallestEncoding not loaded")
	}
	return _cFStringGetSmallestEncoding(theString)
}



var _cFStringGetSystemEncoding func() uint32

// CFStringGetSystemEncoding returns the default encoding used by the operating system when it creates strings.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSystemEncoding()
func CFStringGetSystemEncoding() uint32 {
	if _cFStringGetSystemEncoding == nil {
		panic("CoreFoundation: symbol CFStringGetSystemEncoding not loaded")
	}
	return _cFStringGetSystemEncoding()
}


var _cFStringGetTypeID func() uint

// CFStringGetTypeID returns the type identifier for the CFString opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringGetTypeID()
func CFStringGetTypeID() uint {
	if _cFStringGetTypeID == nil {
		panic("CoreFoundation: symbol CFStringGetTypeID not loaded")
	}
	return _cFStringGetTypeID()
}


var _cFStringHasPrefix func(theString CFStringRef, prefix CFStringRef) bool

// CFStringHasPrefix determines if the character data of a string begin with a specified sequence of characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringHasPrefix(_:_:)
func CFStringHasPrefix(theString CFStringRef, prefix CFStringRef) bool {
	if _cFStringHasPrefix == nil {
		panic("CoreFoundation: symbol CFStringHasPrefix not loaded")
	}
	return _cFStringHasPrefix(theString, prefix)
}


var _cFStringHasSuffix func(theString CFStringRef, suffix CFStringRef) bool

// CFStringHasSuffix determines if a string ends with a specified sequence of characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringHasSuffix(_:_:)
func CFStringHasSuffix(theString CFStringRef, suffix CFStringRef) bool {
	if _cFStringHasSuffix == nil {
		panic("CoreFoundation: symbol CFStringHasSuffix not loaded")
	}
	return _cFStringHasSuffix(theString, suffix)
}



var _cFStringInsert func(str CFMutableStringRef, idx int, insertedStr CFStringRef) 

// CFStringInsert inserts a string at a specified location in the character buffer of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringInsert(_:_:_:)
func CFStringInsert(str CFMutableStringRef, idx int, insertedStr CFStringRef)  {
	if _cFStringInsert == nil {
		panic("CoreFoundation: symbol CFStringInsert not loaded")
	}
	_cFStringInsert(str, idx, insertedStr)
}


var _cFStringIsEncodingAvailable func(encoding uint32) bool

// CFStringIsEncodingAvailable determines whether a given Core Foundation string encoding is available on the current system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringIsEncodingAvailable(_:)
func CFStringIsEncodingAvailable(encoding uint32) bool {
	if _cFStringIsEncodingAvailable == nil {
		panic("CoreFoundation: symbol CFStringIsEncodingAvailable not loaded")
	}
	return _cFStringIsEncodingAvailable(encoding)
}


var _cFStringIsHyphenationAvailableForLocale func(locale CFLocaleRef) bool

// CFStringIsHyphenationAvailableForLocale returns a Boolean value that indicates whether hyphenation data is available.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringIsHyphenationAvailableForLocale(_:)
func CFStringIsHyphenationAvailableForLocale(locale CFLocaleRef) bool {
	if _cFStringIsHyphenationAvailableForLocale == nil {
		panic("CoreFoundation: symbol CFStringIsHyphenationAvailableForLocale not loaded")
	}
	return _cFStringIsHyphenationAvailableForLocale(locale)
}




var _cFStringLowercase func(theString CFMutableStringRef, locale CFLocaleRef) 

// CFStringLowercase changes all uppercase alphabetical characters in a CFMutableString to lowercase.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringLowercase(_:_:)
func CFStringLowercase(theString CFMutableStringRef, locale CFLocaleRef)  {
	if _cFStringLowercase == nil {
		panic("CoreFoundation: symbol CFStringLowercase not loaded")
	}
	_cFStringLowercase(theString, locale)
}


var _cFStringNormalize func(theString CFMutableStringRef, theForm CFStringNormalizationForm) 

// CFStringNormalize normalizes the string into the specified form as described in Unicode Technical Report #15.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalize(_:_:)
func CFStringNormalize(theString CFMutableStringRef, theForm CFStringNormalizationForm)  {
	if _cFStringNormalize == nil {
		panic("CoreFoundation: symbol CFStringNormalize not loaded")
	}
	_cFStringNormalize(theString, theForm)
}


var _cFStringPad func(theString CFMutableStringRef, padString CFStringRef, length int, indexIntoPad int) 

// CFStringPad enlarges a string, padding it with specified characters, or truncates the string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringPad(_:_:_:_:)
func CFStringPad(theString CFMutableStringRef, padString CFStringRef, length int, indexIntoPad int)  {
	if _cFStringPad == nil {
		panic("CoreFoundation: symbol CFStringPad not loaded")
	}
	_cFStringPad(theString, padString, length, indexIntoPad)
}


var _cFStringReplace func(theString CFMutableStringRef, range_ CFRange, replacement CFStringRef) 

// CFStringReplace replaces part of the character contents of a CFMutableString object with another string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringReplace(_:_:_:)
func CFStringReplace(theString CFMutableStringRef, range_ CFRange, replacement CFStringRef)  {
	if _cFStringReplace == nil {
		panic("CoreFoundation: symbol CFStringReplace not loaded")
	}
	_cFStringReplace(theString, range_, replacement)
}


var _cFStringReplaceAll func(theString CFMutableStringRef, replacement CFStringRef) 

// CFStringReplaceAll replaces all characters of a CFMutableString object with other characters.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringReplaceAll(_:_:)
func CFStringReplaceAll(theString CFMutableStringRef, replacement CFStringRef)  {
	if _cFStringReplaceAll == nil {
		panic("CoreFoundation: symbol CFStringReplaceAll not loaded")
	}
	_cFStringReplaceAll(theString, replacement)
}


var _cFStringSetExternalCharactersNoCopy func(theString CFMutableStringRef, chars *uint16, length int, capacity int) 

// CFStringSetExternalCharactersNoCopy notifies a CFMutableString object that its external backing store of Unicode characters has changed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringSetExternalCharactersNoCopy(_:_:_:_:)
func CFStringSetExternalCharactersNoCopy(theString CFMutableStringRef, chars *uint16, length int, capacity int)  {
	if _cFStringSetExternalCharactersNoCopy == nil {
		panic("CoreFoundation: symbol CFStringSetExternalCharactersNoCopy not loaded")
	}
	_cFStringSetExternalCharactersNoCopy(theString, chars, length, capacity)
}


var _cFStringTokenizerAdvanceToNextToken func(tokenizer CFStringTokenizerRef) CFStringTokenizerTokenType

// CFStringTokenizerAdvanceToNextToken advances the tokenizer to the next token and sets that as the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerAdvanceToNextToken(_:)
func CFStringTokenizerAdvanceToNextToken(tokenizer CFStringTokenizerRef) CFStringTokenizerTokenType {
	if _cFStringTokenizerAdvanceToNextToken == nil {
		panic("CoreFoundation: symbol CFStringTokenizerAdvanceToNextToken not loaded")
	}
	return _cFStringTokenizerAdvanceToNextToken(tokenizer)
}


var _cFStringTokenizerCopyBestStringLanguage func(string_ CFStringRef, range_ CFRange) CFStringRef

// CFStringTokenizerCopyBestStringLanguage guesses a language of a given string and returns the guess as a BCP 47 string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyBestStringLanguage(_:_:)
func CFStringTokenizerCopyBestStringLanguage(string_ CFStringRef, range_ CFRange) CFStringRef {
	if _cFStringTokenizerCopyBestStringLanguage == nil {
		panic("CoreFoundation: symbol CFStringTokenizerCopyBestStringLanguage not loaded")
	}
	return _cFStringTokenizerCopyBestStringLanguage(string_, range_)
}


var _cFStringTokenizerCopyCurrentTokenAttribute func(tokenizer CFStringTokenizerRef, attribute uint64) CFTypeRef

// CFStringTokenizerCopyCurrentTokenAttribute returns a given attribute of the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyCurrentTokenAttribute(_:_:)
func CFStringTokenizerCopyCurrentTokenAttribute(tokenizer CFStringTokenizerRef, attribute uint64) CFTypeRef {
	if _cFStringTokenizerCopyCurrentTokenAttribute == nil {
		panic("CoreFoundation: symbol CFStringTokenizerCopyCurrentTokenAttribute not loaded")
	}
	return _cFStringTokenizerCopyCurrentTokenAttribute(tokenizer, attribute)
}


var _cFStringTokenizerCreate func(alloc CFAllocatorRef, string_ CFStringRef, range_ CFRange, options uint64, locale CFLocaleRef) CFStringTokenizerRef

// CFStringTokenizerCreate returns a tokenizer for a given string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCreate(_:_:_:_:_:)
func CFStringTokenizerCreate(alloc CFAllocatorRef, string_ CFStringRef, range_ CFRange, options uint64, locale CFLocaleRef) CFStringTokenizerRef {
	if _cFStringTokenizerCreate == nil {
		panic("CoreFoundation: symbol CFStringTokenizerCreate not loaded")
	}
	return _cFStringTokenizerCreate(alloc, string_, range_, options, locale)
}


var _cFStringTokenizerGetCurrentSubTokens func(tokenizer CFStringTokenizerRef, ranges *CFRange, maxRangeLength int, derivedSubTokens CFMutableArrayRef) int

// CFStringTokenizerGetCurrentSubTokens retrieves the subtokens or derived subtokens contained in the compound token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentSubTokens(_:_:_:_:)
func CFStringTokenizerGetCurrentSubTokens(tokenizer CFStringTokenizerRef, ranges *CFRange, maxRangeLength int, derivedSubTokens CFMutableArrayRef) int {
	if _cFStringTokenizerGetCurrentSubTokens == nil {
		panic("CoreFoundation: symbol CFStringTokenizerGetCurrentSubTokens not loaded")
	}
	return _cFStringTokenizerGetCurrentSubTokens(tokenizer, ranges, maxRangeLength, derivedSubTokens)
}


var _cFStringTokenizerGetCurrentTokenRange func(tokenizer CFStringTokenizerRef) CFRange

// CFStringTokenizerGetCurrentTokenRange returns the range of the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentTokenRange(_:)
func CFStringTokenizerGetCurrentTokenRange(tokenizer CFStringTokenizerRef) CFRange {
	if _cFStringTokenizerGetCurrentTokenRange == nil {
		panic("CoreFoundation: symbol CFStringTokenizerGetCurrentTokenRange not loaded")
	}
	return _cFStringTokenizerGetCurrentTokenRange(tokenizer)
}


var _cFStringTokenizerGetTypeID func() uint

// CFStringTokenizerGetTypeID returns the type ID for CFStringTokenizer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetTypeID()
func CFStringTokenizerGetTypeID() uint {
	if _cFStringTokenizerGetTypeID == nil {
		panic("CoreFoundation: symbol CFStringTokenizerGetTypeID not loaded")
	}
	return _cFStringTokenizerGetTypeID()
}


var _cFStringTokenizerGoToTokenAtIndex func(tokenizer CFStringTokenizerRef, index int) CFStringTokenizerTokenType

// CFStringTokenizerGoToTokenAtIndex finds a token that includes the character at a given index, and set it as the current token.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGoToTokenAtIndex(_:_:)
func CFStringTokenizerGoToTokenAtIndex(tokenizer CFStringTokenizerRef, index int) CFStringTokenizerTokenType {
	if _cFStringTokenizerGoToTokenAtIndex == nil {
		panic("CoreFoundation: symbol CFStringTokenizerGoToTokenAtIndex not loaded")
	}
	return _cFStringTokenizerGoToTokenAtIndex(tokenizer, index)
}


var _cFStringTokenizerSetString func(tokenizer CFStringTokenizerRef, string_ CFStringRef, range_ CFRange) 

// CFStringTokenizerSetString sets the string for a tokenizer.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerSetString(_:_:_:)
func CFStringTokenizerSetString(tokenizer CFStringTokenizerRef, string_ CFStringRef, range_ CFRange)  {
	if _cFStringTokenizerSetString == nil {
		panic("CoreFoundation: symbol CFStringTokenizerSetString not loaded")
	}
	_cFStringTokenizerSetString(tokenizer, string_, range_)
}


var _cFStringTransform func(string_ CFMutableStringRef, range_ *CFRange, transform CFStringRef, reverse bool) bool

// CFStringTransform perform in-place transliteration on a mutable string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTransform(_:_:_:_:)
func CFStringTransform(string_ CFMutableStringRef, range_ *CFRange, transform CFStringRef, reverse bool) bool {
	if _cFStringTransform == nil {
		panic("CoreFoundation: symbol CFStringTransform not loaded")
	}
	return _cFStringTransform(string_, range_, transform, reverse)
}


var _cFStringTrim func(theString CFMutableStringRef, trimString CFStringRef) 

// CFStringTrim trims a specified substring from the beginning and end of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTrim(_:_:)
func CFStringTrim(theString CFMutableStringRef, trimString CFStringRef)  {
	if _cFStringTrim == nil {
		panic("CoreFoundation: symbol CFStringTrim not loaded")
	}
	_cFStringTrim(theString, trimString)
}


var _cFStringTrimWhitespace func(theString CFMutableStringRef) 

// CFStringTrimWhitespace trims whitespace from the beginning and end of a CFMutableString object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTrimWhitespace(_:)
func CFStringTrimWhitespace(theString CFMutableStringRef)  {
	if _cFStringTrimWhitespace == nil {
		panic("CoreFoundation: symbol CFStringTrimWhitespace not loaded")
	}
	_cFStringTrimWhitespace(theString)
}


var _cFStringUppercase func(theString CFMutableStringRef, locale CFLocaleRef) 

// CFStringUppercase changes all lowercase alphabetical characters in a CFMutableString object to uppercase.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFStringUppercase(_:_:)
func CFStringUppercase(theString CFMutableStringRef, locale CFLocaleRef)  {
	if _cFStringUppercase == nil {
		panic("CoreFoundation: symbol CFStringUppercase not loaded")
	}
	_cFStringUppercase(theString, locale)
}

















var _cFTimeZoneCopyAbbreviation func(tz CFTimeZoneRef, at CFAbsoluteTime) CFStringRef

// CFTimeZoneCopyAbbreviation returns the abbreviation of a time zone at a specified date.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviation(_:_:)
func CFTimeZoneCopyAbbreviation(tz CFTimeZoneRef, at CFAbsoluteTime) CFStringRef {
	if _cFTimeZoneCopyAbbreviation == nil {
		panic("CoreFoundation: symbol CFTimeZoneCopyAbbreviation not loaded")
	}
	return _cFTimeZoneCopyAbbreviation(tz, at)
}


var _cFTimeZoneCopyAbbreviationDictionary func() CFDictionaryRef

// CFTimeZoneCopyAbbreviationDictionary returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviationDictionary()
func CFTimeZoneCopyAbbreviationDictionary() CFDictionaryRef {
	if _cFTimeZoneCopyAbbreviationDictionary == nil {
		panic("CoreFoundation: symbol CFTimeZoneCopyAbbreviationDictionary not loaded")
	}
	return _cFTimeZoneCopyAbbreviationDictionary()
}


var _cFTimeZoneCopyDefault func() CFTimeZoneRef

// CFTimeZoneCopyDefault returns the default time zone set for your application.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyDefault()
func CFTimeZoneCopyDefault() CFTimeZoneRef {
	if _cFTimeZoneCopyDefault == nil {
		panic("CoreFoundation: symbol CFTimeZoneCopyDefault not loaded")
	}
	return _cFTimeZoneCopyDefault()
}


var _cFTimeZoneCopyKnownNames func() CFArrayRef

// CFTimeZoneCopyKnownNames returns an array of strings containing the names of all the time zones known to the system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyKnownNames()
func CFTimeZoneCopyKnownNames() CFArrayRef {
	if _cFTimeZoneCopyKnownNames == nil {
		panic("CoreFoundation: symbol CFTimeZoneCopyKnownNames not loaded")
	}
	return _cFTimeZoneCopyKnownNames()
}


var _cFTimeZoneCopyLocalizedName func(tz CFTimeZoneRef, style CFTimeZoneNameStyle, locale CFLocaleRef) CFStringRef

// CFTimeZoneCopyLocalizedName returns the localized name of a given time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyLocalizedName(_:_:_:)
func CFTimeZoneCopyLocalizedName(tz CFTimeZoneRef, style CFTimeZoneNameStyle, locale CFLocaleRef) CFStringRef {
	if _cFTimeZoneCopyLocalizedName == nil {
		panic("CoreFoundation: symbol CFTimeZoneCopyLocalizedName not loaded")
	}
	return _cFTimeZoneCopyLocalizedName(tz, style, locale)
}


var _cFTimeZoneCopySystem func() CFTimeZoneRef

// CFTimeZoneCopySystem returns the time zone currently used by the system.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopySystem()
func CFTimeZoneCopySystem() CFTimeZoneRef {
	if _cFTimeZoneCopySystem == nil {
		panic("CoreFoundation: symbol CFTimeZoneCopySystem not loaded")
	}
	return _cFTimeZoneCopySystem()
}


var _cFTimeZoneCreate func(allocator CFAllocatorRef, name CFStringRef, data CFDataRef) CFTimeZoneRef

// CFTimeZoneCreate creates a time zone with a given name and data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreate(_:_:_:)
func CFTimeZoneCreate(allocator CFAllocatorRef, name CFStringRef, data CFDataRef) CFTimeZoneRef {
	if _cFTimeZoneCreate == nil {
		panic("CoreFoundation: symbol CFTimeZoneCreate not loaded")
	}
	return _cFTimeZoneCreate(allocator, name, data)
}


var _cFTimeZoneCreateWithName func(allocator CFAllocatorRef, name CFStringRef, tryAbbrev bool) CFTimeZoneRef

// CFTimeZoneCreateWithName returns the time zone object identified by a given name or abbreviation.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithName(_:_:_:)
func CFTimeZoneCreateWithName(allocator CFAllocatorRef, name CFStringRef, tryAbbrev bool) CFTimeZoneRef {
	if _cFTimeZoneCreateWithName == nil {
		panic("CoreFoundation: symbol CFTimeZoneCreateWithName not loaded")
	}
	return _cFTimeZoneCreateWithName(allocator, name, tryAbbrev)
}


var _cFTimeZoneCreateWithTimeIntervalFromGMT func(allocator CFAllocatorRef, ti float64) CFTimeZoneRef

// CFTimeZoneCreateWithTimeIntervalFromGMT returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT).
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithTimeIntervalFromGMT(_:_:)
func CFTimeZoneCreateWithTimeIntervalFromGMT(allocator CFAllocatorRef, ti float64) CFTimeZoneRef {
	if _cFTimeZoneCreateWithTimeIntervalFromGMT == nil {
		panic("CoreFoundation: symbol CFTimeZoneCreateWithTimeIntervalFromGMT not loaded")
	}
	return _cFTimeZoneCreateWithTimeIntervalFromGMT(allocator, ti)
}


var _cFTimeZoneGetData func(tz CFTimeZoneRef) CFDataRef

// CFTimeZoneGetData returns the data that stores the information used by a time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetData(_:)
func CFTimeZoneGetData(tz CFTimeZoneRef) CFDataRef {
	if _cFTimeZoneGetData == nil {
		panic("CoreFoundation: symbol CFTimeZoneGetData not loaded")
	}
	return _cFTimeZoneGetData(tz)
}


var _cFTimeZoneGetDaylightSavingTimeOffset func(tz CFTimeZoneRef, at CFAbsoluteTime) float64

// CFTimeZoneGetDaylightSavingTimeOffset returns the daylight saving time offset for a time zone at a given time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetDaylightSavingTimeOffset(_:_:)
func CFTimeZoneGetDaylightSavingTimeOffset(tz CFTimeZoneRef, at CFAbsoluteTime) float64 {
	if _cFTimeZoneGetDaylightSavingTimeOffset == nil {
		panic("CoreFoundation: symbol CFTimeZoneGetDaylightSavingTimeOffset not loaded")
	}
	return _cFTimeZoneGetDaylightSavingTimeOffset(tz, at)
}


var _cFTimeZoneGetName func(tz CFTimeZoneRef) CFStringRef

// CFTimeZoneGetName returns the geopolitical region name that identifies a given time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetName(_:)
func CFTimeZoneGetName(tz CFTimeZoneRef) CFStringRef {
	if _cFTimeZoneGetName == nil {
		panic("CoreFoundation: symbol CFTimeZoneGetName not loaded")
	}
	return _cFTimeZoneGetName(tz)
}


var _cFTimeZoneGetNextDaylightSavingTimeTransition func(tz CFTimeZoneRef, at CFAbsoluteTime) CFAbsoluteTime

// CFTimeZoneGetNextDaylightSavingTimeTransition returns the time in a given time zone of the next daylight saving time transition after a given time.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetNextDaylightSavingTimeTransition(_:_:)
func CFTimeZoneGetNextDaylightSavingTimeTransition(tz CFTimeZoneRef, at CFAbsoluteTime) CFAbsoluteTime {
	if _cFTimeZoneGetNextDaylightSavingTimeTransition == nil {
		panic("CoreFoundation: symbol CFTimeZoneGetNextDaylightSavingTimeTransition not loaded")
	}
	return _cFTimeZoneGetNextDaylightSavingTimeTransition(tz, at)
}


var _cFTimeZoneGetSecondsFromGMT func(tz CFTimeZoneRef, at CFAbsoluteTime) float64

// CFTimeZoneGetSecondsFromGMT returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetSecondsFromGMT(_:_:)
func CFTimeZoneGetSecondsFromGMT(tz CFTimeZoneRef, at CFAbsoluteTime) float64 {
	if _cFTimeZoneGetSecondsFromGMT == nil {
		panic("CoreFoundation: symbol CFTimeZoneGetSecondsFromGMT not loaded")
	}
	return _cFTimeZoneGetSecondsFromGMT(tz, at)
}


var _cFTimeZoneGetTypeID func() uint

// CFTimeZoneGetTypeID returns the type identifier for the CFTimeZone opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetTypeID()
func CFTimeZoneGetTypeID() uint {
	if _cFTimeZoneGetTypeID == nil {
		panic("CoreFoundation: symbol CFTimeZoneGetTypeID not loaded")
	}
	return _cFTimeZoneGetTypeID()
}


var _cFTimeZoneIsDaylightSavingTime func(tz CFTimeZoneRef, at CFAbsoluteTime) bool

// CFTimeZoneIsDaylightSavingTime returns whether or not a time zone is in daylight savings time at a specified date.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneIsDaylightSavingTime(_:_:)
func CFTimeZoneIsDaylightSavingTime(tz CFTimeZoneRef, at CFAbsoluteTime) bool {
	if _cFTimeZoneIsDaylightSavingTime == nil {
		panic("CoreFoundation: symbol CFTimeZoneIsDaylightSavingTime not loaded")
	}
	return _cFTimeZoneIsDaylightSavingTime(tz, at)
}


var _cFTimeZoneResetSystem func() 

// CFTimeZoneResetSystem clears the previously determined system time zone, if any.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneResetSystem()
func CFTimeZoneResetSystem()  {
	if _cFTimeZoneResetSystem == nil {
		panic("CoreFoundation: symbol CFTimeZoneResetSystem not loaded")
	}
	_cFTimeZoneResetSystem()
}


var _cFTimeZoneSetAbbreviationDictionary func(dict CFDictionaryRef) 

// CFTimeZoneSetAbbreviationDictionary sets the abbreviation dictionary to a given dictionary.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetAbbreviationDictionary(_:)
func CFTimeZoneSetAbbreviationDictionary(dict CFDictionaryRef)  {
	if _cFTimeZoneSetAbbreviationDictionary == nil {
		panic("CoreFoundation: symbol CFTimeZoneSetAbbreviationDictionary not loaded")
	}
	_cFTimeZoneSetAbbreviationDictionary(dict)
}


var _cFTimeZoneSetDefault func(tz CFTimeZoneRef) 

// CFTimeZoneSetDefault sets the default time zone for your application the given time zone.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetDefault(_:)
func CFTimeZoneSetDefault(tz CFTimeZoneRef)  {
	if _cFTimeZoneSetDefault == nil {
		panic("CoreFoundation: symbol CFTimeZoneSetDefault not loaded")
	}
	_cFTimeZoneSetDefault(tz)
}


var _cFTreeAppendChild func(tree CFTreeRef, newChild CFTreeRef) 

// CFTreeAppendChild adds a new child to a tree as the last in its list of children.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeAppendChild(_:_:)
func CFTreeAppendChild(tree CFTreeRef, newChild CFTreeRef)  {
	if _cFTreeAppendChild == nil {
		panic("CoreFoundation: symbol CFTreeAppendChild not loaded")
	}
	_cFTreeAppendChild(tree, newChild)
}


var _cFTreeApplyFunctionToChildren func(tree CFTreeRef, applier CFTreeApplierFunction, context unsafe.Pointer) 

// CFTreeApplyFunctionToChildren calls a function once for each immediate child of a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeApplyFunctionToChildren(_:_:_:)
func CFTreeApplyFunctionToChildren(tree CFTreeRef, applier CFTreeApplierFunction, context unsafe.Pointer)  {
	if _cFTreeApplyFunctionToChildren == nil {
		panic("CoreFoundation: symbol CFTreeApplyFunctionToChildren not loaded")
	}
	_cFTreeApplyFunctionToChildren(tree, applier, context)
}


var _cFTreeCreate func(allocator CFAllocatorRef, context *CFTreeContext) CFTreeRef

// CFTreeCreate creates a new CFTree object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeCreate(_:_:)
func CFTreeCreate(allocator CFAllocatorRef, context *CFTreeContext) CFTreeRef {
	if _cFTreeCreate == nil {
		panic("CoreFoundation: symbol CFTreeCreate not loaded")
	}
	return _cFTreeCreate(allocator, context)
}


var _cFTreeFindRoot func(tree CFTreeRef) CFTreeRef

// CFTreeFindRoot returns the root tree of a given tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeFindRoot(_:)
func CFTreeFindRoot(tree CFTreeRef) CFTreeRef {
	if _cFTreeFindRoot == nil {
		panic("CoreFoundation: symbol CFTreeFindRoot not loaded")
	}
	return _cFTreeFindRoot(tree)
}


var _cFTreeGetChildAtIndex func(tree CFTreeRef, idx int) CFTreeRef

// CFTreeGetChildAtIndex returns the child of a tree at the specified index.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildAtIndex(_:_:)
func CFTreeGetChildAtIndex(tree CFTreeRef, idx int) CFTreeRef {
	if _cFTreeGetChildAtIndex == nil {
		panic("CoreFoundation: symbol CFTreeGetChildAtIndex not loaded")
	}
	return _cFTreeGetChildAtIndex(tree, idx)
}


var _cFTreeGetChildCount func(tree CFTreeRef) int

// CFTreeGetChildCount returns the number of children in a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildCount(_:)
func CFTreeGetChildCount(tree CFTreeRef) int {
	if _cFTreeGetChildCount == nil {
		panic("CoreFoundation: symbol CFTreeGetChildCount not loaded")
	}
	return _cFTreeGetChildCount(tree)
}


var _cFTreeGetChildren func(tree CFTreeRef, children *CFTreeRef) 

// CFTreeGetChildren fills a buffer with children from the tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildren(_:_:)
func CFTreeGetChildren(tree CFTreeRef, children *CFTreeRef)  {
	if _cFTreeGetChildren == nil {
		panic("CoreFoundation: symbol CFTreeGetChildren not loaded")
	}
	_cFTreeGetChildren(tree, children)
}


var _cFTreeGetContext func(tree CFTreeRef, context *CFTreeContext) 

// CFTreeGetContext returns the context of the specified tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetContext(_:_:)
func CFTreeGetContext(tree CFTreeRef, context *CFTreeContext)  {
	if _cFTreeGetContext == nil {
		panic("CoreFoundation: symbol CFTreeGetContext not loaded")
	}
	_cFTreeGetContext(tree, context)
}


var _cFTreeGetFirstChild func(tree CFTreeRef) CFTreeRef

// CFTreeGetFirstChild returns the first child of a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetFirstChild(_:)
func CFTreeGetFirstChild(tree CFTreeRef) CFTreeRef {
	if _cFTreeGetFirstChild == nil {
		panic("CoreFoundation: symbol CFTreeGetFirstChild not loaded")
	}
	return _cFTreeGetFirstChild(tree)
}


var _cFTreeGetNextSibling func(tree CFTreeRef) CFTreeRef

// CFTreeGetNextSibling returns the next sibling, adjacent to a given tree, in the parent’s children list.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetNextSibling(_:)
func CFTreeGetNextSibling(tree CFTreeRef) CFTreeRef {
	if _cFTreeGetNextSibling == nil {
		panic("CoreFoundation: symbol CFTreeGetNextSibling not loaded")
	}
	return _cFTreeGetNextSibling(tree)
}


var _cFTreeGetParent func(tree CFTreeRef) CFTreeRef

// CFTreeGetParent returns the parent of a given tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetParent(_:)
func CFTreeGetParent(tree CFTreeRef) CFTreeRef {
	if _cFTreeGetParent == nil {
		panic("CoreFoundation: symbol CFTreeGetParent not loaded")
	}
	return _cFTreeGetParent(tree)
}


var _cFTreeGetTypeID func() uint

// CFTreeGetTypeID returns the type identifier of the CFTree opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetTypeID()
func CFTreeGetTypeID() uint {
	if _cFTreeGetTypeID == nil {
		panic("CoreFoundation: symbol CFTreeGetTypeID not loaded")
	}
	return _cFTreeGetTypeID()
}


var _cFTreeInsertSibling func(tree CFTreeRef, newSibling CFTreeRef) 

// CFTreeInsertSibling inserts a new sibling after a given tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeInsertSibling(_:_:)
func CFTreeInsertSibling(tree CFTreeRef, newSibling CFTreeRef)  {
	if _cFTreeInsertSibling == nil {
		panic("CoreFoundation: symbol CFTreeInsertSibling not loaded")
	}
	_cFTreeInsertSibling(tree, newSibling)
}


var _cFTreePrependChild func(tree CFTreeRef, newChild CFTreeRef) 

// CFTreePrependChild adds a new child to the specified tree as the first in its list of children.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreePrependChild(_:_:)
func CFTreePrependChild(tree CFTreeRef, newChild CFTreeRef)  {
	if _cFTreePrependChild == nil {
		panic("CoreFoundation: symbol CFTreePrependChild not loaded")
	}
	_cFTreePrependChild(tree, newChild)
}


var _cFTreeRemove func(tree CFTreeRef) 

// CFTreeRemove removes a tree from its parent.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemove(_:)
func CFTreeRemove(tree CFTreeRef)  {
	if _cFTreeRemove == nil {
		panic("CoreFoundation: symbol CFTreeRemove not loaded")
	}
	_cFTreeRemove(tree)
}


var _cFTreeRemoveAllChildren func(tree CFTreeRef) 

// CFTreeRemoveAllChildren removes all the children of a tree.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemoveAllChildren(_:)
func CFTreeRemoveAllChildren(tree CFTreeRef)  {
	if _cFTreeRemoveAllChildren == nil {
		panic("CoreFoundation: symbol CFTreeRemoveAllChildren not loaded")
	}
	_cFTreeRemoveAllChildren(tree)
}


var _cFTreeSetContext func(tree CFTreeRef, context *CFTreeContext) 

// CFTreeSetContext replaces the context of a tree by releasing the old information pointer and retaining the new one.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeSetContext(_:_:)
func CFTreeSetContext(tree CFTreeRef, context *CFTreeContext)  {
	if _cFTreeSetContext == nil {
		panic("CoreFoundation: symbol CFTreeSetContext not loaded")
	}
	_cFTreeSetContext(tree, context)
}


var _cFTreeSortChildren func(tree CFTreeRef, comparator CFComparatorFunction, context unsafe.Pointer) 

// CFTreeSortChildren sorts the immediate children of a tree using a specified comparator function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFTreeSortChildren(_:_:_:)
func CFTreeSortChildren(tree CFTreeRef, comparator CFComparatorFunction, context unsafe.Pointer)  {
	if _cFTreeSortChildren == nil {
		panic("CoreFoundation: symbol CFTreeSortChildren not loaded")
	}
	_cFTreeSortChildren(tree, comparator, context)
}


var _cFURLCanBeDecomposed func(anURL CFURLRef) bool

// CFURLCanBeDecomposed determines if the given URL conforms to RFC 1808 and therefore can be decomposed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCanBeDecomposed(_:)
func CFURLCanBeDecomposed(anURL CFURLRef) bool {
	if _cFURLCanBeDecomposed == nil {
		panic("CoreFoundation: symbol CFURLCanBeDecomposed not loaded")
	}
	return _cFURLCanBeDecomposed(anURL)
}


var _cFURLClearResourcePropertyCache func(url CFURLRef) 

// CFURLClearResourcePropertyCache removes all cached resource values and temporary resource values from the URL object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCache(_:)
func CFURLClearResourcePropertyCache(url CFURLRef)  {
	if _cFURLClearResourcePropertyCache == nil {
		panic("CoreFoundation: symbol CFURLClearResourcePropertyCache not loaded")
	}
	_cFURLClearResourcePropertyCache(url)
}


var _cFURLClearResourcePropertyCacheForKey func(url CFURLRef, key CFStringRef) 

// CFURLClearResourcePropertyCacheForKey removes the cached resource value identified by a given key from the URL object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCacheForKey(_:_:)
func CFURLClearResourcePropertyCacheForKey(url CFURLRef, key CFStringRef)  {
	if _cFURLClearResourcePropertyCacheForKey == nil {
		panic("CoreFoundation: symbol CFURLClearResourcePropertyCacheForKey not loaded")
	}
	_cFURLClearResourcePropertyCacheForKey(url, key)
}


var _cFURLCopyAbsoluteURL func(relativeURL CFURLRef) CFURLRef

// CFURLCopyAbsoluteURL creates a new [CFURL] object by resolving the relative portion of a URL against its base.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyAbsoluteURL(_:)
func CFURLCopyAbsoluteURL(relativeURL CFURLRef) CFURLRef {
	if _cFURLCopyAbsoluteURL == nil {
		panic("CoreFoundation: symbol CFURLCopyAbsoluteURL not loaded")
	}
	return _cFURLCopyAbsoluteURL(relativeURL)
}


var _cFURLCopyFileSystemPath func(anURL CFURLRef, pathStyle CFURLPathStyle) CFStringRef

// CFURLCopyFileSystemPath returns the path portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFileSystemPath(_:_:)
func CFURLCopyFileSystemPath(anURL CFURLRef, pathStyle CFURLPathStyle) CFStringRef {
	if _cFURLCopyFileSystemPath == nil {
		panic("CoreFoundation: symbol CFURLCopyFileSystemPath not loaded")
	}
	return _cFURLCopyFileSystemPath(anURL, pathStyle)
}


var _cFURLCopyFragment func(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef

// CFURLCopyFragment returns the fragment from a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFragment(_:_:)
func CFURLCopyFragment(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef {
	if _cFURLCopyFragment == nil {
		panic("CoreFoundation: symbol CFURLCopyFragment not loaded")
	}
	return _cFURLCopyFragment(anURL, charactersToLeaveEscaped)
}


var _cFURLCopyHostName func(anURL CFURLRef) CFStringRef

// CFURLCopyHostName returns the host name of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyHostName(_:)
func CFURLCopyHostName(anURL CFURLRef) CFStringRef {
	if _cFURLCopyHostName == nil {
		panic("CoreFoundation: symbol CFURLCopyHostName not loaded")
	}
	return _cFURLCopyHostName(anURL)
}


var _cFURLCopyLastPathComponent func(url CFURLRef) CFStringRef

// CFURLCopyLastPathComponent returns the last path component of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyLastPathComponent(_:)
func CFURLCopyLastPathComponent(url CFURLRef) CFStringRef {
	if _cFURLCopyLastPathComponent == nil {
		panic("CoreFoundation: symbol CFURLCopyLastPathComponent not loaded")
	}
	return _cFURLCopyLastPathComponent(url)
}


var _cFURLCopyNetLocation func(anURL CFURLRef) CFStringRef

// CFURLCopyNetLocation returns the net location portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyNetLocation(_:)
func CFURLCopyNetLocation(anURL CFURLRef) CFStringRef {
	if _cFURLCopyNetLocation == nil {
		panic("CoreFoundation: symbol CFURLCopyNetLocation not loaded")
	}
	return _cFURLCopyNetLocation(anURL)
}


var _cFURLCopyPassword func(anURL CFURLRef) CFStringRef

// CFURLCopyPassword returns the password of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPassword(_:)
func CFURLCopyPassword(anURL CFURLRef) CFStringRef {
	if _cFURLCopyPassword == nil {
		panic("CoreFoundation: symbol CFURLCopyPassword not loaded")
	}
	return _cFURLCopyPassword(anURL)
}


var _cFURLCopyPath func(anURL CFURLRef) CFStringRef

// CFURLCopyPath returns the path portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPath(_:)
func CFURLCopyPath(anURL CFURLRef) CFStringRef {
	if _cFURLCopyPath == nil {
		panic("CoreFoundation: symbol CFURLCopyPath not loaded")
	}
	return _cFURLCopyPath(anURL)
}


var _cFURLCopyPathExtension func(url CFURLRef) CFStringRef

// CFURLCopyPathExtension returns the path extension of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPathExtension(_:)
func CFURLCopyPathExtension(url CFURLRef) CFStringRef {
	if _cFURLCopyPathExtension == nil {
		panic("CoreFoundation: symbol CFURLCopyPathExtension not loaded")
	}
	return _cFURLCopyPathExtension(url)
}


var _cFURLCopyQueryString func(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef

// CFURLCopyQueryString returns the query string of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyQueryString(_:_:)
func CFURLCopyQueryString(anURL CFURLRef, charactersToLeaveEscaped CFStringRef) CFStringRef {
	if _cFURLCopyQueryString == nil {
		panic("CoreFoundation: symbol CFURLCopyQueryString not loaded")
	}
	return _cFURLCopyQueryString(anURL, charactersToLeaveEscaped)
}


var _cFURLCopyResourcePropertiesForKeys func(url CFURLRef, keys CFArrayRef, err *CFErrorRef) CFDictionaryRef

// CFURLCopyResourcePropertiesForKeys returns the resource values for the properties identified by specified array of keys.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertiesForKeys(_:_:_:)
func CFURLCopyResourcePropertiesForKeys(url CFURLRef, keys CFArrayRef, err *CFErrorRef) CFDictionaryRef {
	if _cFURLCopyResourcePropertiesForKeys == nil {
		panic("CoreFoundation: symbol CFURLCopyResourcePropertiesForKeys not loaded")
	}
	return _cFURLCopyResourcePropertiesForKeys(url, keys, err)
}


var _cFURLCopyResourcePropertyForKey func(url CFURLRef, key CFStringRef, propertyValueTypeRefPtr unsafe.Pointer, err *CFErrorRef) bool

// CFURLCopyResourcePropertyForKey returns the value of a given resource property of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertyForKey(_:_:_:_:)
func CFURLCopyResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValueTypeRefPtr unsafe.Pointer, err *CFErrorRef) bool {
	if _cFURLCopyResourcePropertyForKey == nil {
		panic("CoreFoundation: symbol CFURLCopyResourcePropertyForKey not loaded")
	}
	return _cFURLCopyResourcePropertyForKey(url, key, propertyValueTypeRefPtr, err)
}


var _cFURLCopyResourceSpecifier func(anURL CFURLRef) CFStringRef

// CFURLCopyResourceSpecifier returns any additional resource specifiers after the path.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourceSpecifier(_:)
func CFURLCopyResourceSpecifier(anURL CFURLRef) CFStringRef {
	if _cFURLCopyResourceSpecifier == nil {
		panic("CoreFoundation: symbol CFURLCopyResourceSpecifier not loaded")
	}
	return _cFURLCopyResourceSpecifier(anURL)
}


var _cFURLCopyScheme func(anURL CFURLRef) CFStringRef

// CFURLCopyScheme returns the scheme portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyScheme(_:)
func CFURLCopyScheme(anURL CFURLRef) CFStringRef {
	if _cFURLCopyScheme == nil {
		panic("CoreFoundation: symbol CFURLCopyScheme not loaded")
	}
	return _cFURLCopyScheme(anURL)
}


var _cFURLCopyStrictPath func(anURL CFURLRef, isAbsolute *bool) CFStringRef

// CFURLCopyStrictPath returns the path portion of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyStrictPath(_:_:)
func CFURLCopyStrictPath(anURL CFURLRef, isAbsolute *bool) CFStringRef {
	if _cFURLCopyStrictPath == nil {
		panic("CoreFoundation: symbol CFURLCopyStrictPath not loaded")
	}
	return _cFURLCopyStrictPath(anURL, isAbsolute)
}


var _cFURLCopyUserName func(anURL CFURLRef) CFStringRef

// CFURLCopyUserName returns the user name from a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyUserName(_:)
func CFURLCopyUserName(anURL CFURLRef) CFStringRef {
	if _cFURLCopyUserName == nil {
		panic("CoreFoundation: symbol CFURLCopyUserName not loaded")
	}
	return _cFURLCopyUserName(anURL)
}


var _cFURLCreateAbsoluteURLWithBytes func(alloc CFAllocatorRef, relativeURLBytes *uint8, length int, encoding uint32, baseURL CFURLRef, useCompatibilityMode bool) CFURLRef

// CFURLCreateAbsoluteURLWithBytes creates a new [CFURL] object by resolving the relative portion of a URL, specified as bytes, against its given base URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateAbsoluteURLWithBytes(_:_:_:_:_:_:)
func CFURLCreateAbsoluteURLWithBytes(alloc CFAllocatorRef, relativeURLBytes *uint8, length int, encoding uint32, baseURL CFURLRef, useCompatibilityMode bool) CFURLRef {
	if _cFURLCreateAbsoluteURLWithBytes == nil {
		panic("CoreFoundation: symbol CFURLCreateAbsoluteURLWithBytes not loaded")
	}
	return _cFURLCreateAbsoluteURLWithBytes(alloc, relativeURLBytes, length, encoding, baseURL, useCompatibilityMode)
}


var _cFURLCreateBookmarkData func(allocator CFAllocatorRef, url CFURLRef, options CFURLBookmarkCreationOptions, resourcePropertiesToInclude CFArrayRef, relativeToURL CFURLRef, err *CFErrorRef) CFDataRef

// CFURLCreateBookmarkData returns bookmark data for a URL, created with specified options and resource values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkData(_:_:_:_:_:_:)
func CFURLCreateBookmarkData(allocator CFAllocatorRef, url CFURLRef, options CFURLBookmarkCreationOptions, resourcePropertiesToInclude CFArrayRef, relativeToURL CFURLRef, err *CFErrorRef) CFDataRef {
	if _cFURLCreateBookmarkData == nil {
		panic("CoreFoundation: symbol CFURLCreateBookmarkData not loaded")
	}
	return _cFURLCreateBookmarkData(allocator, url, options, resourcePropertiesToInclude, relativeToURL, err)
}


var _cFURLCreateBookmarkDataFromFile func(allocator CFAllocatorRef, fileURL CFURLRef, errorRef *CFErrorRef) CFDataRef

// CFURLCreateBookmarkDataFromFile initializes and returns bookmark data derived from a file pointed to by a specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromFile(_:_:_:)
func CFURLCreateBookmarkDataFromFile(allocator CFAllocatorRef, fileURL CFURLRef, errorRef *CFErrorRef) CFDataRef {
	if _cFURLCreateBookmarkDataFromFile == nil {
		panic("CoreFoundation: symbol CFURLCreateBookmarkDataFromFile not loaded")
	}
	return _cFURLCreateBookmarkDataFromFile(allocator, fileURL, errorRef)
}


var _cFURLCreateByResolvingBookmarkData func(allocator CFAllocatorRef, bookmark CFDataRef, options CFURLBookmarkResolutionOptions, relativeToURL CFURLRef, resourcePropertiesToInclude CFArrayRef, isStale *bool, err *CFErrorRef) CFURLRef

// CFURLCreateByResolvingBookmarkData returns a new URL made by resolving bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateByResolvingBookmarkData(_:_:_:_:_:_:_:)
func CFURLCreateByResolvingBookmarkData(allocator CFAllocatorRef, bookmark CFDataRef, options CFURLBookmarkResolutionOptions, relativeToURL CFURLRef, resourcePropertiesToInclude CFArrayRef, isStale *bool, err *CFErrorRef) CFURLRef {
	if _cFURLCreateByResolvingBookmarkData == nil {
		panic("CoreFoundation: symbol CFURLCreateByResolvingBookmarkData not loaded")
	}
	return _cFURLCreateByResolvingBookmarkData(allocator, bookmark, options, relativeToURL, resourcePropertiesToInclude, isStale, err)
}


var _cFURLCreateCopyAppendingPathComponent func(allocator CFAllocatorRef, url CFURLRef, pathComponent CFStringRef, isDirectory bool) CFURLRef

// CFURLCreateCopyAppendingPathComponent creates a copy of a given URL and appends a path component.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathComponent(_:_:_:_:)
func CFURLCreateCopyAppendingPathComponent(allocator CFAllocatorRef, url CFURLRef, pathComponent CFStringRef, isDirectory bool) CFURLRef {
	if _cFURLCreateCopyAppendingPathComponent == nil {
		panic("CoreFoundation: symbol CFURLCreateCopyAppendingPathComponent not loaded")
	}
	return _cFURLCreateCopyAppendingPathComponent(allocator, url, pathComponent, isDirectory)
}


var _cFURLCreateCopyAppendingPathExtension func(allocator CFAllocatorRef, url CFURLRef, extension CFStringRef) CFURLRef

// CFURLCreateCopyAppendingPathExtension creates a copy of a given URL and appends a path extension.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathExtension(_:_:_:)
func CFURLCreateCopyAppendingPathExtension(allocator CFAllocatorRef, url CFURLRef, extension CFStringRef) CFURLRef {
	if _cFURLCreateCopyAppendingPathExtension == nil {
		panic("CoreFoundation: symbol CFURLCreateCopyAppendingPathExtension not loaded")
	}
	return _cFURLCreateCopyAppendingPathExtension(allocator, url, extension)
}


var _cFURLCreateCopyDeletingLastPathComponent func(allocator CFAllocatorRef, url CFURLRef) CFURLRef

// CFURLCreateCopyDeletingLastPathComponent creates a copy of a given URL with the last path component deleted.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingLastPathComponent(_:_:)
func CFURLCreateCopyDeletingLastPathComponent(allocator CFAllocatorRef, url CFURLRef) CFURLRef {
	if _cFURLCreateCopyDeletingLastPathComponent == nil {
		panic("CoreFoundation: symbol CFURLCreateCopyDeletingLastPathComponent not loaded")
	}
	return _cFURLCreateCopyDeletingLastPathComponent(allocator, url)
}


var _cFURLCreateCopyDeletingPathExtension func(allocator CFAllocatorRef, url CFURLRef) CFURLRef

// CFURLCreateCopyDeletingPathExtension creates a copy of a given URL with its last path extension removed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingPathExtension(_:_:)
func CFURLCreateCopyDeletingPathExtension(allocator CFAllocatorRef, url CFURLRef) CFURLRef {
	if _cFURLCreateCopyDeletingPathExtension == nil {
		panic("CoreFoundation: symbol CFURLCreateCopyDeletingPathExtension not loaded")
	}
	return _cFURLCreateCopyDeletingPathExtension(allocator, url)
}


var _cFURLCreateData func(allocator CFAllocatorRef, url CFURLRef, encoding uint32, escapeWhitespace bool) CFDataRef

// CFURLCreateData creates a [CFData] object containing the content of a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateData(_:_:_:_:)
func CFURLCreateData(allocator CFAllocatorRef, url CFURLRef, encoding uint32, escapeWhitespace bool) CFDataRef {
	if _cFURLCreateData == nil {
		panic("CoreFoundation: symbol CFURLCreateData not loaded")
	}
	return _cFURLCreateData(allocator, url, encoding, escapeWhitespace)
}


var _cFURLCreateFilePathURL func(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef

// CFURLCreateFilePathURL returns a new file path URL that refers to the same resource as a specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFilePathURL(_:_:_:)
func CFURLCreateFilePathURL(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef {
	if _cFURLCreateFilePathURL == nil {
		panic("CoreFoundation: symbol CFURLCreateFilePathURL not loaded")
	}
	return _cFURLCreateFilePathURL(allocator, url, err)
}


var _cFURLCreateFileReferenceURL func(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef

// CFURLCreateFileReferenceURL returns a new file reference URL that points to the same resource as a specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFileReferenceURL(_:_:_:)
func CFURLCreateFileReferenceURL(allocator CFAllocatorRef, url CFURLRef, err *CFErrorRef) CFURLRef {
	if _cFURLCreateFileReferenceURL == nil {
		panic("CoreFoundation: symbol CFURLCreateFileReferenceURL not loaded")
	}
	return _cFURLCreateFileReferenceURL(allocator, url, err)
}


var _cFURLCreateFromFileSystemRepresentation func(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool) CFURLRef

// CFURLCreateFromFileSystemRepresentation creates a new [CFURL] object for a file system entity using the native representation.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentation(_:_:_:_:)
func CFURLCreateFromFileSystemRepresentation(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool) CFURLRef {
	if _cFURLCreateFromFileSystemRepresentation == nil {
		panic("CoreFoundation: symbol CFURLCreateFromFileSystemRepresentation not loaded")
	}
	return _cFURLCreateFromFileSystemRepresentation(allocator, buffer, bufLen, isDirectory)
}


var _cFURLCreateFromFileSystemRepresentationRelativeToBase func(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool, baseURL CFURLRef) CFURLRef

// CFURLCreateFromFileSystemRepresentationRelativeToBase creates a [CFURL] object from a native character string path relative to a base URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentationRelativeToBase(_:_:_:_:_:)
func CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator CFAllocatorRef, buffer *uint8, bufLen int, isDirectory bool, baseURL CFURLRef) CFURLRef {
	if _cFURLCreateFromFileSystemRepresentationRelativeToBase == nil {
		panic("CoreFoundation: symbol CFURLCreateFromFileSystemRepresentationRelativeToBase not loaded")
	}
	return _cFURLCreateFromFileSystemRepresentationRelativeToBase(allocator, buffer, bufLen, isDirectory, baseURL)
}


var _cFURLCreateResourcePropertiesForKeysFromBookmarkData func(allocator CFAllocatorRef, resourcePropertiesToReturn CFArrayRef, bookmark CFDataRef) CFDictionaryRef

// CFURLCreateResourcePropertiesForKeysFromBookmarkData returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertiesForKeysFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator CFAllocatorRef, resourcePropertiesToReturn CFArrayRef, bookmark CFDataRef) CFDictionaryRef {
	if _cFURLCreateResourcePropertiesForKeysFromBookmarkData == nil {
		panic("CoreFoundation: symbol CFURLCreateResourcePropertiesForKeysFromBookmarkData not loaded")
	}
	return _cFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator, resourcePropertiesToReturn, bookmark)
}


var _cFURLCreateResourcePropertyForKeyFromBookmarkData func(allocator CFAllocatorRef, resourcePropertyKey CFStringRef, bookmark CFDataRef) CFTypeRef

// CFURLCreateResourcePropertyForKeyFromBookmarkData returns the value of a resource property from specified bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertyForKeyFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator CFAllocatorRef, resourcePropertyKey CFStringRef, bookmark CFDataRef) CFTypeRef {
	if _cFURLCreateResourcePropertyForKeyFromBookmarkData == nil {
		panic("CoreFoundation: symbol CFURLCreateResourcePropertyForKeyFromBookmarkData not loaded")
	}
	return _cFURLCreateResourcePropertyForKeyFromBookmarkData(allocator, resourcePropertyKey, bookmark)
}


var _cFURLCreateStringByReplacingPercentEscapes func(allocator CFAllocatorRef, originalString CFStringRef, charactersToLeaveEscaped CFStringRef) CFStringRef

// CFURLCreateStringByReplacingPercentEscapes creates a new string by replacing any percent escape sequences with their character equivalent.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapes(_:_:_:)
func CFURLCreateStringByReplacingPercentEscapes(allocator CFAllocatorRef, originalString CFStringRef, charactersToLeaveEscaped CFStringRef) CFStringRef {
	if _cFURLCreateStringByReplacingPercentEscapes == nil {
		panic("CoreFoundation: symbol CFURLCreateStringByReplacingPercentEscapes not loaded")
	}
	return _cFURLCreateStringByReplacingPercentEscapes(allocator, originalString, charactersToLeaveEscaped)
}


var _cFURLCreateWithBytes func(allocator CFAllocatorRef, URLBytes *uint8, length int, encoding uint32, baseURL CFURLRef) CFURLRef

// CFURLCreateWithBytes creates a [CFURL] object using a given character bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithBytes(_:_:_:_:_:)
func CFURLCreateWithBytes(allocator CFAllocatorRef, URLBytes *uint8, length int, encoding uint32, baseURL CFURLRef) CFURLRef {
	if _cFURLCreateWithBytes == nil {
		panic("CoreFoundation: symbol CFURLCreateWithBytes not loaded")
	}
	return _cFURLCreateWithBytes(allocator, URLBytes, length, encoding, baseURL)
}


var _cFURLCreateWithFileSystemPath func(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool) CFURLRef

// CFURLCreateWithFileSystemPath creates a [CFURL] object using a local file system path string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPath(_:_:_:_:)
func CFURLCreateWithFileSystemPath(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool) CFURLRef {
	if _cFURLCreateWithFileSystemPath == nil {
		panic("CoreFoundation: symbol CFURLCreateWithFileSystemPath not loaded")
	}
	return _cFURLCreateWithFileSystemPath(allocator, filePath, pathStyle, isDirectory)
}


var _cFURLCreateWithFileSystemPathRelativeToBase func(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool, baseURL CFURLRef) CFURLRef

// CFURLCreateWithFileSystemPathRelativeToBase creates a [CFURL] object using a local file system path string relative to a base URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPathRelativeToBase(_:_:_:_:_:)
func CFURLCreateWithFileSystemPathRelativeToBase(allocator CFAllocatorRef, filePath CFStringRef, pathStyle CFURLPathStyle, isDirectory bool, baseURL CFURLRef) CFURLRef {
	if _cFURLCreateWithFileSystemPathRelativeToBase == nil {
		panic("CoreFoundation: symbol CFURLCreateWithFileSystemPathRelativeToBase not loaded")
	}
	return _cFURLCreateWithFileSystemPathRelativeToBase(allocator, filePath, pathStyle, isDirectory, baseURL)
}


var _cFURLCreateWithString func(allocator CFAllocatorRef, URLString CFStringRef, baseURL CFURLRef) CFURLRef

// CFURLCreateWithString creates a [CFURL] object using a given [CFString] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithString(_:_:_:)
func CFURLCreateWithString(allocator CFAllocatorRef, URLString CFStringRef, baseURL CFURLRef) CFURLRef {
	if _cFURLCreateWithString == nil {
		panic("CoreFoundation: symbol CFURLCreateWithString not loaded")
	}
	return _cFURLCreateWithString(allocator, URLString, baseURL)
}


var _cFURLEnumeratorCreateForDirectoryURL func(alloc CFAllocatorRef, directoryURL CFURLRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef

// CFURLEnumeratorCreateForDirectoryURL creates and returns a directory enumerator with provided enumerator behavior options and properties to be prefetched.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForDirectoryURL(_:_:_:_:)
func CFURLEnumeratorCreateForDirectoryURL(alloc CFAllocatorRef, directoryURL CFURLRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef {
	if _cFURLEnumeratorCreateForDirectoryURL == nil {
		panic("CoreFoundation: symbol CFURLEnumeratorCreateForDirectoryURL not loaded")
	}
	return _cFURLEnumeratorCreateForDirectoryURL(alloc, directoryURL, option, propertyKeys)
}


var _cFURLEnumeratorCreateForMountedVolumes func(alloc CFAllocatorRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef

// CFURLEnumeratorCreateForMountedVolumes creates and returns a volume enumerator with provided enumerator behavior options and properties to be prefetched.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForMountedVolumes(_:_:_:)
func CFURLEnumeratorCreateForMountedVolumes(alloc CFAllocatorRef, option CFURLEnumeratorOptions, propertyKeys CFArrayRef) CFURLEnumeratorRef {
	if _cFURLEnumeratorCreateForMountedVolumes == nil {
		panic("CoreFoundation: symbol CFURLEnumeratorCreateForMountedVolumes not loaded")
	}
	return _cFURLEnumeratorCreateForMountedVolumes(alloc, option, propertyKeys)
}


var _cFURLEnumeratorGetDescendentLevel func(enumerator CFURLEnumeratorRef) int

// CFURLEnumeratorGetDescendentLevel returns the number of levels a recursive directory enumerator has descended.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetDescendentLevel(_:)
func CFURLEnumeratorGetDescendentLevel(enumerator CFURLEnumeratorRef) int {
	if _cFURLEnumeratorGetDescendentLevel == nil {
		panic("CoreFoundation: symbol CFURLEnumeratorGetDescendentLevel not loaded")
	}
	return _cFURLEnumeratorGetDescendentLevel(enumerator)
}


var _cFURLEnumeratorGetNextURL func(enumerator CFURLEnumeratorRef, url *CFURLRef, err *CFErrorRef) CFURLEnumeratorResult

// CFURLEnumeratorGetNextURL advances an enumerator to the next URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetNextURL(_:_:_:)
func CFURLEnumeratorGetNextURL(enumerator CFURLEnumeratorRef, url *CFURLRef, err *CFErrorRef) CFURLEnumeratorResult {
	if _cFURLEnumeratorGetNextURL == nil {
		panic("CoreFoundation: symbol CFURLEnumeratorGetNextURL not loaded")
	}
	return _cFURLEnumeratorGetNextURL(enumerator, url, err)
}


var _cFURLEnumeratorGetTypeID func() uint

// CFURLEnumeratorGetTypeID returns the opaque type identifier for the CFURLEnumerator opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetTypeID()
func CFURLEnumeratorGetTypeID() uint {
	if _cFURLEnumeratorGetTypeID == nil {
		panic("CoreFoundation: symbol CFURLEnumeratorGetTypeID not loaded")
	}
	return _cFURLEnumeratorGetTypeID()
}


var _cFURLEnumeratorSkipDescendents func(enumerator CFURLEnumeratorRef) 

// CFURLEnumeratorSkipDescendents tells a recursive enumerator not to descend into the directory at the URL that was returned by the most recent call to the CFURLEnumeratorGetNextURL(_:_:_:) function.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorSkipDescendents(_:)
func CFURLEnumeratorSkipDescendents(enumerator CFURLEnumeratorRef)  {
	if _cFURLEnumeratorSkipDescendents == nil {
		panic("CoreFoundation: symbol CFURLEnumeratorSkipDescendents not loaded")
	}
	_cFURLEnumeratorSkipDescendents(enumerator)
}


var _cFURLGetBaseURL func(anURL CFURLRef) CFURLRef

// CFURLGetBaseURL returns the base URL of a given URL if it exists.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBaseURL(_:)
func CFURLGetBaseURL(anURL CFURLRef) CFURLRef {
	if _cFURLGetBaseURL == nil {
		panic("CoreFoundation: symbol CFURLGetBaseURL not loaded")
	}
	return _cFURLGetBaseURL(anURL)
}


var _cFURLGetByteRangeForComponent func(url CFURLRef, component CFURLComponentType, rangeIncludingSeparators *CFRange) CFRange

// CFURLGetByteRangeForComponent returns the range of the specified component in the bytes of a URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetByteRangeForComponent(_:_:_:)
func CFURLGetByteRangeForComponent(url CFURLRef, component CFURLComponentType, rangeIncludingSeparators *CFRange) CFRange {
	if _cFURLGetByteRangeForComponent == nil {
		panic("CoreFoundation: symbol CFURLGetByteRangeForComponent not loaded")
	}
	return _cFURLGetByteRangeForComponent(url, component, rangeIncludingSeparators)
}


var _cFURLGetBytes func(url CFURLRef, buffer *uint8, bufferLength int) int

// CFURLGetBytes returns by reference the byte representation of a URL object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBytes(_:_:_:)
func CFURLGetBytes(url CFURLRef, buffer *uint8, bufferLength int) int {
	if _cFURLGetBytes == nil {
		panic("CoreFoundation: symbol CFURLGetBytes not loaded")
	}
	return _cFURLGetBytes(url, buffer, bufferLength)
}


var _cFURLGetFileSystemRepresentation func(url CFURLRef, resolveAgainstBase bool, buffer *uint8, maxBufLen int) bool

// CFURLGetFileSystemRepresentation fills a buffer with the file system’s native string representation of a given URL’s path.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFileSystemRepresentation(_:_:_:_:)
func CFURLGetFileSystemRepresentation(url CFURLRef, resolveAgainstBase bool, buffer *uint8, maxBufLen int) bool {
	if _cFURLGetFileSystemRepresentation == nil {
		panic("CoreFoundation: symbol CFURLGetFileSystemRepresentation not loaded")
	}
	return _cFURLGetFileSystemRepresentation(url, resolveAgainstBase, buffer, maxBufLen)
}


var _cFURLGetPortNumber func(anURL CFURLRef) int32

// CFURLGetPortNumber returns the port number from a given URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetPortNumber(_:)
func CFURLGetPortNumber(anURL CFURLRef) int32 {
	if _cFURLGetPortNumber == nil {
		panic("CoreFoundation: symbol CFURLGetPortNumber not loaded")
	}
	return _cFURLGetPortNumber(anURL)
}


var _cFURLGetString func(anURL CFURLRef) CFStringRef

// CFURLGetString returns the URL as a [CFString] object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetString(_:)
func CFURLGetString(anURL CFURLRef) CFStringRef {
	if _cFURLGetString == nil {
		panic("CoreFoundation: symbol CFURLGetString not loaded")
	}
	return _cFURLGetString(anURL)
}


var _cFURLGetTypeID func() uint

// CFURLGetTypeID returns the type identifier for the [CFURL] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLGetTypeID()
func CFURLGetTypeID() uint {
	if _cFURLGetTypeID == nil {
		panic("CoreFoundation: symbol CFURLGetTypeID not loaded")
	}
	return _cFURLGetTypeID()
}


var _cFURLHasDirectoryPath func(anURL CFURLRef) bool

// CFURLHasDirectoryPath determines if a given URL’s path represents a directory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLHasDirectoryPath(_:)
func CFURLHasDirectoryPath(anURL CFURLRef) bool {
	if _cFURLHasDirectoryPath == nil {
		panic("CoreFoundation: symbol CFURLHasDirectoryPath not loaded")
	}
	return _cFURLHasDirectoryPath(anURL)
}


var _cFURLIsFileReferenceURL func(url CFURLRef) bool

// CFURLIsFileReferenceURL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLIsFileReferenceURL(_:)
func CFURLIsFileReferenceURL(url CFURLRef) bool {
	if _cFURLIsFileReferenceURL == nil {
		panic("CoreFoundation: symbol CFURLIsFileReferenceURL not loaded")
	}
	return _cFURLIsFileReferenceURL(url)
}


var _cFURLResourceIsReachable func(url CFURLRef, err *CFErrorRef) bool

// CFURLResourceIsReachable returns whether the resource pointed to by a file URL can be reached.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLResourceIsReachable(_:_:)
func CFURLResourceIsReachable(url CFURLRef, err *CFErrorRef) bool {
	if _cFURLResourceIsReachable == nil {
		panic("CoreFoundation: symbol CFURLResourceIsReachable not loaded")
	}
	return _cFURLResourceIsReachable(url, err)
}


var _cFURLSetResourcePropertiesForKeys func(url CFURLRef, keyedPropertyValues CFDictionaryRef, err *CFErrorRef) bool

// CFURLSetResourcePropertiesForKeys sets the URL’s resource properties for a given set of keys to a given set of values.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertiesForKeys(_:_:_:)
func CFURLSetResourcePropertiesForKeys(url CFURLRef, keyedPropertyValues CFDictionaryRef, err *CFErrorRef) bool {
	if _cFURLSetResourcePropertiesForKeys == nil {
		panic("CoreFoundation: symbol CFURLSetResourcePropertiesForKeys not loaded")
	}
	return _cFURLSetResourcePropertiesForKeys(url, keyedPropertyValues, err)
}


var _cFURLSetResourcePropertyForKey func(url CFURLRef, key CFStringRef, propertyValue CFTypeRef, err *CFErrorRef) bool

// CFURLSetResourcePropertyForKey sets the URL’s resource property for a given key to a given value.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertyForKey(_:_:_:_:)
func CFURLSetResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValue CFTypeRef, err *CFErrorRef) bool {
	if _cFURLSetResourcePropertyForKey == nil {
		panic("CoreFoundation: symbol CFURLSetResourcePropertyForKey not loaded")
	}
	return _cFURLSetResourcePropertyForKey(url, key, propertyValue, err)
}


var _cFURLSetTemporaryResourcePropertyForKey func(url CFURLRef, key CFStringRef, propertyValue CFTypeRef) 

// CFURLSetTemporaryResourcePropertyForKey sets a temporary resource value on the URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLSetTemporaryResourcePropertyForKey(_:_:_:)
func CFURLSetTemporaryResourcePropertyForKey(url CFURLRef, key CFStringRef, propertyValue CFTypeRef)  {
	if _cFURLSetTemporaryResourcePropertyForKey == nil {
		panic("CoreFoundation: symbol CFURLSetTemporaryResourcePropertyForKey not loaded")
	}
	_cFURLSetTemporaryResourcePropertyForKey(url, key, propertyValue)
}


var _cFURLStartAccessingSecurityScopedResource func(url CFURLRef) bool

// CFURLStartAccessingSecurityScopedResource in an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
func CFURLStartAccessingSecurityScopedResource(url CFURLRef) bool {
	if _cFURLStartAccessingSecurityScopedResource == nil {
		panic("CoreFoundation: symbol CFURLStartAccessingSecurityScopedResource not loaded")
	}
	return _cFURLStartAccessingSecurityScopedResource(url)
}


var _cFURLStopAccessingSecurityScopedResource func(url CFURLRef) 

// CFURLStopAccessingSecurityScopedResource in an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
func CFURLStopAccessingSecurityScopedResource(url CFURLRef)  {
	if _cFURLStopAccessingSecurityScopedResource == nil {
		panic("CoreFoundation: symbol CFURLStopAccessingSecurityScopedResource not loaded")
	}
	_cFURLStopAccessingSecurityScopedResource(url)
}


var _cFURLWriteBookmarkDataToFile func(bookmarkRef CFDataRef, fileURL CFURLRef, options CFURLBookmarkFileCreationOptions, errorRef *CFErrorRef) bool

// CFURLWriteBookmarkDataToFile creates an alias file on disk at a specified location with specified bookmark data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteBookmarkDataToFile(_:_:_:_:)
func CFURLWriteBookmarkDataToFile(bookmarkRef CFDataRef, fileURL CFURLRef, options CFURLBookmarkFileCreationOptions, errorRef *CFErrorRef) bool {
	if _cFURLWriteBookmarkDataToFile == nil {
		panic("CoreFoundation: symbol CFURLWriteBookmarkDataToFile not loaded")
	}
	return _cFURLWriteBookmarkDataToFile(bookmarkRef, fileURL, options, errorRef)
}


var _cFUUIDCreate func(alloc CFAllocatorRef) CFUUIDRef

// CFUUIDCreate creates a Universally Unique Identifier (UUID) object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreate(_:)
func CFUUIDCreate(alloc CFAllocatorRef) CFUUIDRef {
	if _cFUUIDCreate == nil {
		panic("CoreFoundation: symbol CFUUIDCreate not loaded")
	}
	return _cFUUIDCreate(alloc)
}


var _cFUUIDCreateFromString func(alloc CFAllocatorRef, uuidStr CFStringRef) CFUUIDRef

// CFUUIDCreateFromString creates a CFUUID object for a specified string.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromString(_:_:)
func CFUUIDCreateFromString(alloc CFAllocatorRef, uuidStr CFStringRef) CFUUIDRef {
	if _cFUUIDCreateFromString == nil {
		panic("CoreFoundation: symbol CFUUIDCreateFromString not loaded")
	}
	return _cFUUIDCreateFromString(alloc, uuidStr)
}


var _cFUUIDCreateFromUUIDBytes func(alloc CFAllocatorRef, bytes CFUUIDBytes) CFUUIDRef

// CFUUIDCreateFromUUIDBytes creates a CFUUID object from raw UUID bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromUUIDBytes(_:_:)
func CFUUIDCreateFromUUIDBytes(alloc CFAllocatorRef, bytes CFUUIDBytes) CFUUIDRef {
	if _cFUUIDCreateFromUUIDBytes == nil {
		panic("CoreFoundation: symbol CFUUIDCreateFromUUIDBytes not loaded")
	}
	return _cFUUIDCreateFromUUIDBytes(alloc, bytes)
}


var _cFUUIDCreateString func(alloc CFAllocatorRef, uuid CFUUIDRef) CFStringRef

// CFUUIDCreateString returns the string representation of a specified CFUUID object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateString(_:_:)
func CFUUIDCreateString(alloc CFAllocatorRef, uuid CFUUIDRef) CFStringRef {
	if _cFUUIDCreateString == nil {
		panic("CoreFoundation: symbol CFUUIDCreateString not loaded")
	}
	return _cFUUIDCreateString(alloc, uuid)
}


var _cFUUIDCreateWithBytes func(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef

// CFUUIDCreateWithBytes creates a CFUUID object from raw UUID bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDCreateWithBytes(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef {
	if _cFUUIDCreateWithBytes == nil {
		panic("CoreFoundation: symbol CFUUIDCreateWithBytes not loaded")
	}
	return _cFUUIDCreateWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
}


var _cFUUIDGetConstantUUIDWithBytes func(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef

// CFUUIDGetConstantUUIDWithBytes returns a CFUUID object from raw UUID bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetConstantUUIDWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDGetConstantUUIDWithBytes(alloc CFAllocatorRef, byte0 uint8, byte1 uint8, byte2 uint8, byte3 uint8, byte4 uint8, byte5 uint8, byte6 uint8, byte7 uint8, byte8 uint8, byte9 uint8, byte10 uint8, byte11 uint8, byte12 uint8, byte13 uint8, byte14 uint8, byte15 uint8) CFUUIDRef {
	if _cFUUIDGetConstantUUIDWithBytes == nil {
		panic("CoreFoundation: symbol CFUUIDGetConstantUUIDWithBytes not loaded")
	}
	return _cFUUIDGetConstantUUIDWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
}


var _cFUUIDGetTypeID func() uint

// CFUUIDGetTypeID returns the type identifier for all CFUUID objects.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetTypeID()
func CFUUIDGetTypeID() uint {
	if _cFUUIDGetTypeID == nil {
		panic("CoreFoundation: symbol CFUUIDGetTypeID not loaded")
	}
	return _cFUUIDGetTypeID()
}


var _cFUUIDGetUUIDBytes func(uuid CFUUIDRef) CFUUIDBytes

// CFUUIDGetUUIDBytes returns the value of a UUID object as raw bytes.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetUUIDBytes(_:)
func CFUUIDGetUUIDBytes(uuid CFUUIDRef) CFUUIDBytes {
	if _cFUUIDGetUUIDBytes == nil {
		panic("CoreFoundation: symbol CFUUIDGetUUIDBytes not loaded")
	}
	return _cFUUIDGetUUIDBytes(uuid)
}


var _cFUserNotificationCancel func(userNotification CFUserNotificationRef) int32

// CFUserNotificationCancel cancels a user notification dialog.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCancel(_:)
func CFUserNotificationCancel(userNotification CFUserNotificationRef) int32 {
	if _cFUserNotificationCancel == nil {
		panic("CoreFoundation: symbol CFUserNotificationCancel not loaded")
	}
	return _cFUserNotificationCancel(userNotification)
}



var _cFUserNotificationCreate func(allocator CFAllocatorRef, timeout float64, flags uint64, err *int32, dictionary CFDictionaryRef) CFUserNotificationRef

// CFUserNotificationCreate creates a CFUserNotification object and displays its notification dialog on screen.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreate(_:_:_:_:_:)
func CFUserNotificationCreate(allocator CFAllocatorRef, timeout float64, flags uint64, err *int32, dictionary CFDictionaryRef) CFUserNotificationRef {
	if _cFUserNotificationCreate == nil {
		panic("CoreFoundation: symbol CFUserNotificationCreate not loaded")
	}
	return _cFUserNotificationCreate(allocator, timeout, flags, err, dictionary)
}


var _cFUserNotificationCreateRunLoopSource func(allocator CFAllocatorRef, userNotification CFUserNotificationRef, callout CFUserNotificationCallBack, order int) CFRunLoopSourceRef

// CFUserNotificationCreateRunLoopSource creates a run loop source for a user notification.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreateRunLoopSource(_:_:_:_:)
func CFUserNotificationCreateRunLoopSource(allocator CFAllocatorRef, userNotification CFUserNotificationRef, callout CFUserNotificationCallBack, order int) CFRunLoopSourceRef {
	if _cFUserNotificationCreateRunLoopSource == nil {
		panic("CoreFoundation: symbol CFUserNotificationCreateRunLoopSource not loaded")
	}
	return _cFUserNotificationCreateRunLoopSource(allocator, userNotification, callout, order)
}


var _cFUserNotificationDisplayAlert func(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef, alternateButtonTitle CFStringRef, otherButtonTitle CFStringRef, responseFlags *uint64) int32

// CFUserNotificationDisplayAlert displays a user notification dialog and waits for a user response.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayAlert(_:_:_:_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayAlert(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef, alternateButtonTitle CFStringRef, otherButtonTitle CFStringRef, responseFlags *uint64) int32 {
	if _cFUserNotificationDisplayAlert == nil {
		panic("CoreFoundation: symbol CFUserNotificationDisplayAlert not loaded")
	}
	return _cFUserNotificationDisplayAlert(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle, alternateButtonTitle, otherButtonTitle, responseFlags)
}


var _cFUserNotificationDisplayNotice func(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef) int32

// CFUserNotificationDisplayNotice displays a user notification dialog that does not need a user response.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayNotice(_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayNotice(timeout float64, flags uint64, iconURL CFURLRef, soundURL CFURLRef, localizationURL CFURLRef, alertHeader CFStringRef, alertMessage CFStringRef, defaultButtonTitle CFStringRef) int32 {
	if _cFUserNotificationDisplayNotice == nil {
		panic("CoreFoundation: symbol CFUserNotificationDisplayNotice not loaded")
	}
	return _cFUserNotificationDisplayNotice(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle)
}


var _cFUserNotificationGetResponseDictionary func(userNotification CFUserNotificationRef) CFDictionaryRef

// CFUserNotificationGetResponseDictionary returns the dictionary containing all the text field values from a dismissed notification dialog.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseDictionary(_:)
func CFUserNotificationGetResponseDictionary(userNotification CFUserNotificationRef) CFDictionaryRef {
	if _cFUserNotificationGetResponseDictionary == nil {
		panic("CoreFoundation: symbol CFUserNotificationGetResponseDictionary not loaded")
	}
	return _cFUserNotificationGetResponseDictionary(userNotification)
}


var _cFUserNotificationGetResponseValue func(userNotification CFUserNotificationRef, key CFStringRef, idx int) CFStringRef

// CFUserNotificationGetResponseValue extracts the values of the text fields from a dismissed notification dialog.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseValue(_:_:_:)
func CFUserNotificationGetResponseValue(userNotification CFUserNotificationRef, key CFStringRef, idx int) CFStringRef {
	if _cFUserNotificationGetResponseValue == nil {
		panic("CoreFoundation: symbol CFUserNotificationGetResponseValue not loaded")
	}
	return _cFUserNotificationGetResponseValue(userNotification, key, idx)
}


var _cFUserNotificationGetTypeID func() uint

// CFUserNotificationGetTypeID returns the type identifier for the [CFUserNotification] opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetTypeID()
func CFUserNotificationGetTypeID() uint {
	if _cFUserNotificationGetTypeID == nil {
		panic("CoreFoundation: symbol CFUserNotificationGetTypeID not loaded")
	}
	return _cFUserNotificationGetTypeID()
}



var _cFUserNotificationReceiveResponse func(userNotification CFUserNotificationRef, timeout float64, responseFlags *uint64) int32

// CFUserNotificationReceiveResponse waits for the user to respond to a notification or for the notification to time out.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationReceiveResponse(_:_:_:)
func CFUserNotificationReceiveResponse(userNotification CFUserNotificationRef, timeout float64, responseFlags *uint64) int32 {
	if _cFUserNotificationReceiveResponse == nil {
		panic("CoreFoundation: symbol CFUserNotificationReceiveResponse not loaded")
	}
	return _cFUserNotificationReceiveResponse(userNotification, timeout, responseFlags)
}



var _cFUserNotificationUpdate func(userNotification CFUserNotificationRef, timeout float64, flags uint64, dictionary CFDictionaryRef) int32

// CFUserNotificationUpdate updates a displayed user notification dialog with new user interface information.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationUpdate(_:_:_:_:)
func CFUserNotificationUpdate(userNotification CFUserNotificationRef, timeout float64, flags uint64, dictionary CFDictionaryRef) int32 {
	if _cFUserNotificationUpdate == nil {
		panic("CoreFoundation: symbol CFUserNotificationUpdate not loaded")
	}
	return _cFUserNotificationUpdate(userNotification, timeout, flags, dictionary)
}


var _cFWriteStreamCanAcceptBytes func(stream CFWriteStreamRef) bool

// CFWriteStreamCanAcceptBytes returns whether a writable stream can accept new data without blocking.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCanAcceptBytes(_:)
func CFWriteStreamCanAcceptBytes(stream CFWriteStreamRef) bool {
	if _cFWriteStreamCanAcceptBytes == nil {
		panic("CoreFoundation: symbol CFWriteStreamCanAcceptBytes not loaded")
	}
	return _cFWriteStreamCanAcceptBytes(stream)
}


var _cFWriteStreamClose func(stream CFWriteStreamRef) 

// CFWriteStreamClose closes a writable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClose(_:)
func CFWriteStreamClose(stream CFWriteStreamRef)  {
	if _cFWriteStreamClose == nil {
		panic("CoreFoundation: symbol CFWriteStreamClose not loaded")
	}
	_cFWriteStreamClose(stream)
}


var _cFWriteStreamCopyDispatchQueue func(stream CFWriteStreamRef) uintptr

// CFWriteStreamCopyDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyDispatchQueue(_:)
func CFWriteStreamCopyDispatchQueue(stream CFWriteStreamRef) dispatch.Queue {
	if _cFWriteStreamCopyDispatchQueue == nil {
		panic("CoreFoundation: symbol CFWriteStreamCopyDispatchQueue not loaded")
	}
	return dispatch.QueueFromHandle(_cFWriteStreamCopyDispatchQueue(stream))
}


var _cFWriteStreamCopyError func(stream CFWriteStreamRef) CFErrorRef

// CFWriteStreamCopyError returns the error associated with a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyError(_:)
func CFWriteStreamCopyError(stream CFWriteStreamRef) CFErrorRef {
	if _cFWriteStreamCopyError == nil {
		panic("CoreFoundation: symbol CFWriteStreamCopyError not loaded")
	}
	return _cFWriteStreamCopyError(stream)
}


var _cFWriteStreamCopyProperty func(stream CFWriteStreamRef, propertyName CFStreamPropertyKey) CFTypeRef

// CFWriteStreamCopyProperty returns the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyProperty(_:_:)
func CFWriteStreamCopyProperty(stream CFWriteStreamRef, propertyName CFStreamPropertyKey) CFTypeRef {
	if _cFWriteStreamCopyProperty == nil {
		panic("CoreFoundation: symbol CFWriteStreamCopyProperty not loaded")
	}
	return _cFWriteStreamCopyProperty(stream, propertyName)
}


var _cFWriteStreamCreateWithAllocatedBuffers func(alloc CFAllocatorRef, bufferAllocator CFAllocatorRef) CFWriteStreamRef

// CFWriteStreamCreateWithAllocatedBuffers creates a writable stream for a growable block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithAllocatedBuffers(_:_:)
func CFWriteStreamCreateWithAllocatedBuffers(alloc CFAllocatorRef, bufferAllocator CFAllocatorRef) CFWriteStreamRef {
	if _cFWriteStreamCreateWithAllocatedBuffers == nil {
		panic("CoreFoundation: symbol CFWriteStreamCreateWithAllocatedBuffers not loaded")
	}
	return _cFWriteStreamCreateWithAllocatedBuffers(alloc, bufferAllocator)
}


var _cFWriteStreamCreateWithBuffer func(alloc CFAllocatorRef, buffer *uint8, bufferCapacity int) CFWriteStreamRef

// CFWriteStreamCreateWithBuffer creates a writable stream for a fixed-size block of memory.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithBuffer(_:_:_:)
func CFWriteStreamCreateWithBuffer(alloc CFAllocatorRef, buffer *uint8, bufferCapacity int) CFWriteStreamRef {
	if _cFWriteStreamCreateWithBuffer == nil {
		panic("CoreFoundation: symbol CFWriteStreamCreateWithBuffer not loaded")
	}
	return _cFWriteStreamCreateWithBuffer(alloc, buffer, bufferCapacity)
}


var _cFWriteStreamCreateWithFile func(alloc CFAllocatorRef, fileURL CFURLRef) CFWriteStreamRef

// CFWriteStreamCreateWithFile creates a writable stream for a file.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithFile(_:_:)
func CFWriteStreamCreateWithFile(alloc CFAllocatorRef, fileURL CFURLRef) CFWriteStreamRef {
	if _cFWriteStreamCreateWithFile == nil {
		panic("CoreFoundation: symbol CFWriteStreamCreateWithFile not loaded")
	}
	return _cFWriteStreamCreateWithFile(alloc, fileURL)
}


var _cFWriteStreamGetError func(stream CFWriteStreamRef) CFStreamError

// CFWriteStreamGetError returns the error status of a stream.
//
// Deprecated: Use [CFWriteStreamCopyError(_:)](<doc://com.apple.corefoundation/documentation/CoreFoundation/CFWriteStreamCopyError(_:)>) instead.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetError(_:)
func CFWriteStreamGetError(stream CFWriteStreamRef) CFStreamError {
	if _cFWriteStreamGetError == nil {
		panic("CoreFoundation: symbol CFWriteStreamGetError not loaded")
	}
	return _cFWriteStreamGetError(stream)
}


var _cFWriteStreamGetStatus func(stream CFWriteStreamRef) CFStreamStatus

// CFWriteStreamGetStatus returns the current state of a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetStatus(_:)
func CFWriteStreamGetStatus(stream CFWriteStreamRef) CFStreamStatus {
	if _cFWriteStreamGetStatus == nil {
		panic("CoreFoundation: symbol CFWriteStreamGetStatus not loaded")
	}
	return _cFWriteStreamGetStatus(stream)
}


var _cFWriteStreamGetTypeID func() uint

// CFWriteStreamGetTypeID returns the type identifier of all CFWriteStream objects.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetTypeID()
func CFWriteStreamGetTypeID() uint {
	if _cFWriteStreamGetTypeID == nil {
		panic("CoreFoundation: symbol CFWriteStreamGetTypeID not loaded")
	}
	return _cFWriteStreamGetTypeID()
}


var _cFWriteStreamOpen func(stream CFWriteStreamRef) bool

// CFWriteStreamOpen opens a stream for writing.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamOpen(_:)
func CFWriteStreamOpen(stream CFWriteStreamRef) bool {
	if _cFWriteStreamOpen == nil {
		panic("CoreFoundation: symbol CFWriteStreamOpen not loaded")
	}
	return _cFWriteStreamOpen(stream)
}


var _cFWriteStreamScheduleWithRunLoop func(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) 

// CFWriteStreamScheduleWithRunLoop schedules a stream into a run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamScheduleWithRunLoop(_:_:_:)
func CFWriteStreamScheduleWithRunLoop(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)  {
	if _cFWriteStreamScheduleWithRunLoop == nil {
		panic("CoreFoundation: symbol CFWriteStreamScheduleWithRunLoop not loaded")
	}
	_cFWriteStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
}


var _cFWriteStreamSetClient func(stream CFWriteStreamRef, streamEvents uint64, clientCB CFWriteStreamClientCallBack, clientContext *CFStreamClientContext) bool

// CFWriteStreamSetClient assigns a client to a stream, which receives callbacks when certain events occur.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetClient(_:_:_:_:)
func CFWriteStreamSetClient(stream CFWriteStreamRef, streamEvents uint64, clientCB CFWriteStreamClientCallBack, clientContext *CFStreamClientContext) bool {
	if _cFWriteStreamSetClient == nil {
		panic("CoreFoundation: symbol CFWriteStreamSetClient not loaded")
	}
	return _cFWriteStreamSetClient(stream, streamEvents, clientCB, clientContext)
}


var _cFWriteStreamSetDispatchQueue func(stream CFWriteStreamRef, q uintptr) 

// CFWriteStreamSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetDispatchQueue(_:_:)
func CFWriteStreamSetDispatchQueue(stream CFWriteStreamRef, q dispatch.Queue)  {
	if _cFWriteStreamSetDispatchQueue == nil {
		panic("CoreFoundation: symbol CFWriteStreamSetDispatchQueue not loaded")
	}
	_cFWriteStreamSetDispatchQueue(stream, uintptr(q.Handle()))
}


var _cFWriteStreamSetProperty func(stream CFWriteStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool

// CFWriteStreamSetProperty sets the value of a property for a stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetProperty(_:_:_:)
func CFWriteStreamSetProperty(stream CFWriteStreamRef, propertyName CFStreamPropertyKey, propertyValue CFTypeRef) bool {
	if _cFWriteStreamSetProperty == nil {
		panic("CoreFoundation: symbol CFWriteStreamSetProperty not loaded")
	}
	return _cFWriteStreamSetProperty(stream, propertyName, propertyValue)
}


var _cFWriteStreamUnscheduleFromRunLoop func(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode) 

// CFWriteStreamUnscheduleFromRunLoop removes a stream from a particular run loop.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamUnscheduleFromRunLoop(_:_:_:)
func CFWriteStreamUnscheduleFromRunLoop(stream CFWriteStreamRef, runLoop CFRunLoopRef, runLoopMode CFRunLoopMode)  {
	if _cFWriteStreamUnscheduleFromRunLoop == nil {
		panic("CoreFoundation: symbol CFWriteStreamUnscheduleFromRunLoop not loaded")
	}
	_cFWriteStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
}


var _cFWriteStreamWrite func(stream CFWriteStreamRef, buffer *uint8, bufferLength int) int

// CFWriteStreamWrite writes data to a writable stream.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamWrite(_:_:_:)
func CFWriteStreamWrite(stream CFWriteStreamRef, buffer *uint8, bufferLength int) int {
	if _cFWriteStreamWrite == nil {
		panic("CoreFoundation: symbol CFWriteStreamWrite not loaded")
	}
	return _cFWriteStreamWrite(stream, buffer, bufferLength)
}


var _cFXMLCreateStringByEscapingEntities func(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef

// CFXMLCreateStringByEscapingEntities given a CFString object containing XML source with unescaped entities, returns a string with specified XML entities escaped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByEscapingEntities(_:_:_:)
func CFXMLCreateStringByEscapingEntities(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef {
	if _cFXMLCreateStringByEscapingEntities == nil {
		panic("CoreFoundation: symbol CFXMLCreateStringByEscapingEntities not loaded")
	}
	return _cFXMLCreateStringByEscapingEntities(allocator, string_, entitiesDictionary)
}


var _cFXMLCreateStringByUnescapingEntities func(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef

// CFXMLCreateStringByUnescapingEntities given a CFString object containing XML source with escaped entities, returns a string with specified XML entities unescaped.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByUnescapingEntities(_:_:_:)
func CFXMLCreateStringByUnescapingEntities(allocator CFAllocatorRef, string_ CFStringRef, entitiesDictionary CFDictionaryRef) CFStringRef {
	if _cFXMLCreateStringByUnescapingEntities == nil {
		panic("CoreFoundation: symbol CFXMLCreateStringByUnescapingEntities not loaded")
	}
	return _cFXMLCreateStringByUnescapingEntities(allocator, string_, entitiesDictionary)
}


var _cFXMLNodeCreate func(alloc CFAllocatorRef, xmlType uint32, dataString CFStringRef, additionalInfoPtr unsafe.Pointer, version int) unsafe.Pointer

// CFXMLNodeCreate creates a new CFXMLNode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreate
func CFXMLNodeCreate(alloc CFAllocatorRef, xmlType uint32, dataString CFStringRef, additionalInfoPtr unsafe.Pointer, version int) unsafe.Pointer {
	if _cFXMLNodeCreate == nil {
		panic("CoreFoundation: symbol CFXMLNodeCreate not loaded")
	}
	return _cFXMLNodeCreate(alloc, xmlType, dataString, additionalInfoPtr, version)
}


var _cFXMLNodeCreateCopy func(alloc CFAllocatorRef, origNode unsafe.Pointer) unsafe.Pointer

// CFXMLNodeCreateCopy creates a copy of a CFXMLNode object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreateCopy
func CFXMLNodeCreateCopy(alloc CFAllocatorRef, origNode unsafe.Pointer) unsafe.Pointer {
	if _cFXMLNodeCreateCopy == nil {
		panic("CoreFoundation: symbol CFXMLNodeCreateCopy not loaded")
	}
	return _cFXMLNodeCreateCopy(alloc, origNode)
}


var _cFXMLNodeGetInfoPtr func(node unsafe.Pointer) unsafe.Pointer

// CFXMLNodeGetInfoPtr returns the additional information pointer of a CFXMLNode object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetInfoPtr
func CFXMLNodeGetInfoPtr(node unsafe.Pointer) unsafe.Pointer {
	if _cFXMLNodeGetInfoPtr == nil {
		panic("CoreFoundation: symbol CFXMLNodeGetInfoPtr not loaded")
	}
	return _cFXMLNodeGetInfoPtr(node)
}


var _cFXMLNodeGetString func(node unsafe.Pointer) CFStringRef

// CFXMLNodeGetString returns the data string from a CFXMLNode.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetString
func CFXMLNodeGetString(node unsafe.Pointer) CFStringRef {
	if _cFXMLNodeGetString == nil {
		panic("CoreFoundation: symbol CFXMLNodeGetString not loaded")
	}
	return _cFXMLNodeGetString(node)
}


var _cFXMLNodeGetTypeCode func(node unsafe.Pointer) uint32

// CFXMLNodeGetTypeCode returns the XML structure type code for a CFXMLNode object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeCode
func CFXMLNodeGetTypeCode(node unsafe.Pointer) uint32 {
	if _cFXMLNodeGetTypeCode == nil {
		panic("CoreFoundation: symbol CFXMLNodeGetTypeCode not loaded")
	}
	return _cFXMLNodeGetTypeCode(node)
}


var _cFXMLNodeGetTypeID func() uint

// CFXMLNodeGetTypeID returns the type identifier code for the CFXMLNode opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeID
func CFXMLNodeGetTypeID() uint {
	if _cFXMLNodeGetTypeID == nil {
		panic("CoreFoundation: symbol CFXMLNodeGetTypeID not loaded")
	}
	return _cFXMLNodeGetTypeID()
}


var _cFXMLNodeGetVersion func(node unsafe.Pointer) int

// CFXMLNodeGetVersion returns the version number for a CFXMLNode object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetVersion
func CFXMLNodeGetVersion(node unsafe.Pointer) int {
	if _cFXMLNodeGetVersion == nil {
		panic("CoreFoundation: symbol CFXMLNodeGetVersion not loaded")
	}
	return _cFXMLNodeGetVersion(node)
}


var _cFXMLParserAbort func(parser CFXMLParserRef, errorCode CFXMLParserStatusCode, errorDescription CFStringRef) 

// CFXMLParserAbort causes a parser to abort with the given error code and description.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserAbort
func CFXMLParserAbort(parser CFXMLParserRef, errorCode CFXMLParserStatusCode, errorDescription CFStringRef)  {
	if _cFXMLParserAbort == nil {
		panic("CoreFoundation: symbol CFXMLParserAbort not loaded")
	}
	_cFXMLParserAbort(parser, errorCode, errorDescription)
}


var _cFXMLParserCopyErrorDescription func(parser CFXMLParserRef) CFStringRef

// CFXMLParserCopyErrorDescription returns the user-readable description of the current error condition.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCopyErrorDescription
func CFXMLParserCopyErrorDescription(parser CFXMLParserRef) CFStringRef {
	if _cFXMLParserCopyErrorDescription == nil {
		panic("CoreFoundation: symbol CFXMLParserCopyErrorDescription not loaded")
	}
	return _cFXMLParserCopyErrorDescription(parser)
}


var _cFXMLParserCreate func(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef

// CFXMLParserCreate creates a new XML parser for the specified XML data.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreate
func CFXMLParserCreate(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef {
	if _cFXMLParserCreate == nil {
		panic("CoreFoundation: symbol CFXMLParserCreate not loaded")
	}
	return _cFXMLParserCreate(allocator, xmlData, dataSource, parseOptions, versionOfNodes, callBacks, context)
}


var _cFXMLParserCreateWithDataFromURL func(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef

// CFXMLParserCreateWithDataFromURL creates a new XML parser for the specified XML data at the specified URL.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreateWithDataFromURL
func CFXMLParserCreateWithDataFromURL(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, callBacks *CFXMLParserCallBacks, context *CFXMLParserContext) CFXMLParserRef {
	if _cFXMLParserCreateWithDataFromURL == nil {
		panic("CoreFoundation: symbol CFXMLParserCreateWithDataFromURL not loaded")
	}
	return _cFXMLParserCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes, callBacks, context)
}


var _cFXMLParserGetCallBacks func(parser CFXMLParserRef, callBacks *CFXMLParserCallBacks) 

// CFXMLParserGetCallBacks returns the callbacks associated with an XML parser when it was created.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetCallBacks
func CFXMLParserGetCallBacks(parser CFXMLParserRef, callBacks *CFXMLParserCallBacks)  {
	if _cFXMLParserGetCallBacks == nil {
		panic("CoreFoundation: symbol CFXMLParserGetCallBacks not loaded")
	}
	_cFXMLParserGetCallBacks(parser, callBacks)
}


var _cFXMLParserGetContext func(parser CFXMLParserRef, context *CFXMLParserContext) 

// CFXMLParserGetContext returns the context for an XML parser.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetContext
func CFXMLParserGetContext(parser CFXMLParserRef, context *CFXMLParserContext)  {
	if _cFXMLParserGetContext == nil {
		panic("CoreFoundation: symbol CFXMLParserGetContext not loaded")
	}
	_cFXMLParserGetContext(parser, context)
}


var _cFXMLParserGetDocument func(parser CFXMLParserRef) unsafe.Pointer

// CFXMLParserGetDocument returns the top-most object returned by the create XML structure callback.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetDocument
func CFXMLParserGetDocument(parser CFXMLParserRef) unsafe.Pointer {
	if _cFXMLParserGetDocument == nil {
		panic("CoreFoundation: symbol CFXMLParserGetDocument not loaded")
	}
	return _cFXMLParserGetDocument(parser)
}


var _cFXMLParserGetLineNumber func(parser CFXMLParserRef) int

// CFXMLParserGetLineNumber returns the line number of the current parse location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLineNumber
func CFXMLParserGetLineNumber(parser CFXMLParserRef) int {
	if _cFXMLParserGetLineNumber == nil {
		panic("CoreFoundation: symbol CFXMLParserGetLineNumber not loaded")
	}
	return _cFXMLParserGetLineNumber(parser)
}


var _cFXMLParserGetLocation func(parser CFXMLParserRef) int

// CFXMLParserGetLocation returns the character index of the current parse location.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLocation
func CFXMLParserGetLocation(parser CFXMLParserRef) int {
	if _cFXMLParserGetLocation == nil {
		panic("CoreFoundation: symbol CFXMLParserGetLocation not loaded")
	}
	return _cFXMLParserGetLocation(parser)
}


var _cFXMLParserGetSourceURL func(parser CFXMLParserRef) CFURLRef

// CFXMLParserGetSourceURL returns the URL for the XML data being parsed.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetSourceURL
func CFXMLParserGetSourceURL(parser CFXMLParserRef) CFURLRef {
	if _cFXMLParserGetSourceURL == nil {
		panic("CoreFoundation: symbol CFXMLParserGetSourceURL not loaded")
	}
	return _cFXMLParserGetSourceURL(parser)
}


var _cFXMLParserGetStatusCode func(parser CFXMLParserRef) CFXMLParserStatusCode

// CFXMLParserGetStatusCode returns a numeric code indicating the current status of the parser.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetStatusCode
func CFXMLParserGetStatusCode(parser CFXMLParserRef) CFXMLParserStatusCode {
	if _cFXMLParserGetStatusCode == nil {
		panic("CoreFoundation: symbol CFXMLParserGetStatusCode not loaded")
	}
	return _cFXMLParserGetStatusCode(parser)
}


var _cFXMLParserGetTypeID func() uint

// CFXMLParserGetTypeID returns the type identifier for the CFXMLParser opaque type.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetTypeID
func CFXMLParserGetTypeID() uint {
	if _cFXMLParserGetTypeID == nil {
		panic("CoreFoundation: symbol CFXMLParserGetTypeID not loaded")
	}
	return _cFXMLParserGetTypeID()
}


var _cFXMLParserParse func(parser CFXMLParserRef) bool

// CFXMLParserParse begins a parse of the XML data that was associated with the parser when it was created.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserParse
func CFXMLParserParse(parser CFXMLParserRef) bool {
	if _cFXMLParserParse == nil {
		panic("CoreFoundation: symbol CFXMLParserParse not loaded")
	}
	return _cFXMLParserParse(parser)
}


var _cFXMLTreeCreateFromData func(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef

// CFXMLTreeCreateFromData parses the given XML data and returns the resulting CFXMLTree object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromData
func CFXMLTreeCreateFromData(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef {
	if _cFXMLTreeCreateFromData == nil {
		panic("CoreFoundation: symbol CFXMLTreeCreateFromData not loaded")
	}
	return _cFXMLTreeCreateFromData(allocator, xmlData, dataSource, parseOptions, versionOfNodes)
}


var _cFXMLTreeCreateFromDataWithError func(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, errorDict *CFDictionaryRef) CFXMLTreeRef

// CFXMLTreeCreateFromDataWithError parses the given XML data and returns the resulting CFXMLTree object and any error information.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromDataWithError
func CFXMLTreeCreateFromDataWithError(allocator CFAllocatorRef, xmlData CFDataRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int, errorDict *CFDictionaryRef) CFXMLTreeRef {
	if _cFXMLTreeCreateFromDataWithError == nil {
		panic("CoreFoundation: symbol CFXMLTreeCreateFromDataWithError not loaded")
	}
	return _cFXMLTreeCreateFromDataWithError(allocator, xmlData, dataSource, parseOptions, versionOfNodes, errorDict)
}


var _cFXMLTreeCreateWithDataFromURL func(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef

// CFXMLTreeCreateWithDataFromURL creates a new CFXMLTree object by loading the data to be parsed directly from a data source.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithDataFromURL
func CFXMLTreeCreateWithDataFromURL(allocator CFAllocatorRef, dataSource CFURLRef, parseOptions uint64, versionOfNodes int) CFXMLTreeRef {
	if _cFXMLTreeCreateWithDataFromURL == nil {
		panic("CoreFoundation: symbol CFXMLTreeCreateWithDataFromURL not loaded")
	}
	return _cFXMLTreeCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes)
}


var _cFXMLTreeCreateWithNode func(allocator CFAllocatorRef, node unsafe.Pointer) CFXMLTreeRef

// CFXMLTreeCreateWithNode creates a childless, parentless CFXMLTree object node for a CFXMLNode object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithNode
func CFXMLTreeCreateWithNode(allocator CFAllocatorRef, node unsafe.Pointer) CFXMLTreeRef {
	if _cFXMLTreeCreateWithNode == nil {
		panic("CoreFoundation: symbol CFXMLTreeCreateWithNode not loaded")
	}
	return _cFXMLTreeCreateWithNode(allocator, node)
}


var _cFXMLTreeCreateXMLData func(allocator CFAllocatorRef, xmlTree CFXMLTreeRef) CFDataRef

// CFXMLTreeCreateXMLData generates an XML document from a CFXMLTree object which is ready to be written to permanent storage.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateXMLData
func CFXMLTreeCreateXMLData(allocator CFAllocatorRef, xmlTree CFXMLTreeRef) CFDataRef {
	if _cFXMLTreeCreateXMLData == nil {
		panic("CoreFoundation: symbol CFXMLTreeCreateXMLData not loaded")
	}
	return _cFXMLTreeCreateXMLData(allocator, xmlTree)
}


var _cFXMLTreeGetNode func(xmlTree CFXMLTreeRef) unsafe.Pointer

// CFXMLTreeGetNode returns the node of a CFXMLTree object.
//
// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeGetNode
func CFXMLTreeGetNode(xmlTree CFXMLTreeRef) unsafe.Pointer {
	if _cFXMLTreeGetNode == nil {
		panic("CoreFoundation: symbol CFXMLTreeGetNode not loaded")
	}
	return _cFXMLTreeGetNode(xmlTree)
}



func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_cFAbsoluteTimeGetCurrent, frameworkHandle, "CFAbsoluteTimeGetCurrent")
		registerFunc(&_cFAllocatorAllocate, frameworkHandle, "CFAllocatorAllocate")
		registerFunc(&_cFAllocatorAllocateBytes, frameworkHandle, "CFAllocatorAllocateBytes")
		registerFunc(&_cFAllocatorAllocateTyped, frameworkHandle, "CFAllocatorAllocateTyped")
		registerFunc(&_cFAllocatorCreate, frameworkHandle, "CFAllocatorCreate")
		registerFunc(&_cFAllocatorDeallocate, frameworkHandle, "CFAllocatorDeallocate")
		registerFunc(&_cFAllocatorGetContext, frameworkHandle, "CFAllocatorGetContext")
		registerFunc(&_cFAllocatorGetDefault, frameworkHandle, "CFAllocatorGetDefault")
		registerFunc(&_cFAllocatorGetPreferredSizeForSize, frameworkHandle, "CFAllocatorGetPreferredSizeForSize")
		registerFunc(&_cFAllocatorGetTypeID, frameworkHandle, "CFAllocatorGetTypeID")
		registerFunc(&_cFAllocatorReallocate, frameworkHandle, "CFAllocatorReallocate")
		registerFunc(&_cFAllocatorReallocateBytes, frameworkHandle, "CFAllocatorReallocateBytes")
		registerFunc(&_cFAllocatorReallocateTyped, frameworkHandle, "CFAllocatorReallocateTyped")
		registerFunc(&_cFAllocatorSetDefault, frameworkHandle, "CFAllocatorSetDefault")
		registerFunc(&_cFArrayAppendArray, frameworkHandle, "CFArrayAppendArray")
		registerFunc(&_cFArrayAppendValue, frameworkHandle, "CFArrayAppendValue")
		registerFunc(&_cFArrayApplyFunction, frameworkHandle, "CFArrayApplyFunction")
		registerFunc(&_cFArrayBSearchValues, frameworkHandle, "CFArrayBSearchValues")
		registerFunc(&_cFArrayContainsValue, frameworkHandle, "CFArrayContainsValue")
		registerFunc(&_cFArrayCreate, frameworkHandle, "CFArrayCreate")
		registerFunc(&_cFArrayCreateCopy, frameworkHandle, "CFArrayCreateCopy")
		registerFunc(&_cFArrayCreateMutable, frameworkHandle, "CFArrayCreateMutable")
		registerFunc(&_cFArrayCreateMutableCopy, frameworkHandle, "CFArrayCreateMutableCopy")
		registerFunc(&_cFArrayExchangeValuesAtIndices, frameworkHandle, "CFArrayExchangeValuesAtIndices")
		registerFunc(&_cFArrayGetCount, frameworkHandle, "CFArrayGetCount")
		registerFunc(&_cFArrayGetCountOfValue, frameworkHandle, "CFArrayGetCountOfValue")
		registerFunc(&_cFArrayGetFirstIndexOfValue, frameworkHandle, "CFArrayGetFirstIndexOfValue")
		registerFunc(&_cFArrayGetLastIndexOfValue, frameworkHandle, "CFArrayGetLastIndexOfValue")
		registerFunc(&_cFArrayGetTypeID, frameworkHandle, "CFArrayGetTypeID")
		registerFunc(&_cFArrayGetValueAtIndex, frameworkHandle, "CFArrayGetValueAtIndex")
		registerFunc(&_cFArrayGetValues, frameworkHandle, "CFArrayGetValues")
		registerFunc(&_cFArrayInsertValueAtIndex, frameworkHandle, "CFArrayInsertValueAtIndex")
		registerFunc(&_cFArrayRemoveAllValues, frameworkHandle, "CFArrayRemoveAllValues")
		registerFunc(&_cFArrayRemoveValueAtIndex, frameworkHandle, "CFArrayRemoveValueAtIndex")
		registerFunc(&_cFArrayReplaceValues, frameworkHandle, "CFArrayReplaceValues")
		registerFunc(&_cFArraySetValueAtIndex, frameworkHandle, "CFArraySetValueAtIndex")
		registerFunc(&_cFArraySortValues, frameworkHandle, "CFArraySortValues")
		registerFunc(&_cFAttributedStringBeginEditing, frameworkHandle, "CFAttributedStringBeginEditing")
		registerFunc(&_cFAttributedStringCreate, frameworkHandle, "CFAttributedStringCreate")
		registerFunc(&_cFAttributedStringCreateCopy, frameworkHandle, "CFAttributedStringCreateCopy")
		registerFunc(&_cFAttributedStringCreateMutable, frameworkHandle, "CFAttributedStringCreateMutable")
		registerFunc(&_cFAttributedStringCreateMutableCopy, frameworkHandle, "CFAttributedStringCreateMutableCopy")
		registerFunc(&_cFAttributedStringCreateWithSubstring, frameworkHandle, "CFAttributedStringCreateWithSubstring")
		registerFunc(&_cFAttributedStringEndEditing, frameworkHandle, "CFAttributedStringEndEditing")
		registerFunc(&_cFAttributedStringGetAttribute, frameworkHandle, "CFAttributedStringGetAttribute")
		registerFunc(&_cFAttributedStringGetAttributeAndLongestEffectiveRange, frameworkHandle, "CFAttributedStringGetAttributeAndLongestEffectiveRange")
		registerFunc(&_cFAttributedStringGetAttributes, frameworkHandle, "CFAttributedStringGetAttributes")
		registerFunc(&_cFAttributedStringGetAttributesAndLongestEffectiveRange, frameworkHandle, "CFAttributedStringGetAttributesAndLongestEffectiveRange")
		registerFunc(&_cFAttributedStringGetBidiLevelsAndResolvedDirections, frameworkHandle, "CFAttributedStringGetBidiLevelsAndResolvedDirections")
		registerFunc(&_cFAttributedStringGetLength, frameworkHandle, "CFAttributedStringGetLength")
		registerFunc(&_cFAttributedStringGetMutableString, frameworkHandle, "CFAttributedStringGetMutableString")
		registerFunc(&_cFAttributedStringGetStatisticalWritingDirections, frameworkHandle, "CFAttributedStringGetStatisticalWritingDirections")
		registerFunc(&_cFAttributedStringGetString, frameworkHandle, "CFAttributedStringGetString")
		registerFunc(&_cFAttributedStringGetTypeID, frameworkHandle, "CFAttributedStringGetTypeID")
		registerFunc(&_cFAttributedStringRemoveAttribute, frameworkHandle, "CFAttributedStringRemoveAttribute")
		registerFunc(&_cFAttributedStringReplaceAttributedString, frameworkHandle, "CFAttributedStringReplaceAttributedString")
		registerFunc(&_cFAttributedStringReplaceString, frameworkHandle, "CFAttributedStringReplaceString")
		registerFunc(&_cFAttributedStringSetAttribute, frameworkHandle, "CFAttributedStringSetAttribute")
		registerFunc(&_cFAttributedStringSetAttributes, frameworkHandle, "CFAttributedStringSetAttributes")
		registerFunc(&_cFAutorelease, frameworkHandle, "CFAutorelease")
		registerFunc(&_cFBagAddValue, frameworkHandle, "CFBagAddValue")
		registerFunc(&_cFBagApplyFunction, frameworkHandle, "CFBagApplyFunction")
		registerFunc(&_cFBagContainsValue, frameworkHandle, "CFBagContainsValue")
		registerFunc(&_cFBagCreate, frameworkHandle, "CFBagCreate")
		registerFunc(&_cFBagCreateCopy, frameworkHandle, "CFBagCreateCopy")
		registerFunc(&_cFBagCreateMutable, frameworkHandle, "CFBagCreateMutable")
		registerFunc(&_cFBagCreateMutableCopy, frameworkHandle, "CFBagCreateMutableCopy")
		registerFunc(&_cFBagGetCount, frameworkHandle, "CFBagGetCount")
		registerFunc(&_cFBagGetCountOfValue, frameworkHandle, "CFBagGetCountOfValue")
		registerFunc(&_cFBagGetTypeID, frameworkHandle, "CFBagGetTypeID")
		registerFunc(&_cFBagGetValue, frameworkHandle, "CFBagGetValue")
		registerFunc(&_cFBagGetValueIfPresent, frameworkHandle, "CFBagGetValueIfPresent")
		registerFunc(&_cFBagGetValues, frameworkHandle, "CFBagGetValues")
		registerFunc(&_cFBagRemoveAllValues, frameworkHandle, "CFBagRemoveAllValues")
		registerFunc(&_cFBagRemoveValue, frameworkHandle, "CFBagRemoveValue")
		registerFunc(&_cFBagReplaceValue, frameworkHandle, "CFBagReplaceValue")
		registerFunc(&_cFBagSetValue, frameworkHandle, "CFBagSetValue")
		registerFunc(&_cFBinaryHeapAddValue, frameworkHandle, "CFBinaryHeapAddValue")
		registerFunc(&_cFBinaryHeapApplyFunction, frameworkHandle, "CFBinaryHeapApplyFunction")
		registerFunc(&_cFBinaryHeapContainsValue, frameworkHandle, "CFBinaryHeapContainsValue")
		registerFunc(&_cFBinaryHeapCreate, frameworkHandle, "CFBinaryHeapCreate")
		registerFunc(&_cFBinaryHeapCreateCopy, frameworkHandle, "CFBinaryHeapCreateCopy")
		registerFunc(&_cFBinaryHeapGetCount, frameworkHandle, "CFBinaryHeapGetCount")
		registerFunc(&_cFBinaryHeapGetCountOfValue, frameworkHandle, "CFBinaryHeapGetCountOfValue")
		registerFunc(&_cFBinaryHeapGetMinimum, frameworkHandle, "CFBinaryHeapGetMinimum")
		registerFunc(&_cFBinaryHeapGetMinimumIfPresent, frameworkHandle, "CFBinaryHeapGetMinimumIfPresent")
		registerFunc(&_cFBinaryHeapGetTypeID, frameworkHandle, "CFBinaryHeapGetTypeID")
		registerFunc(&_cFBinaryHeapGetValues, frameworkHandle, "CFBinaryHeapGetValues")
		registerFunc(&_cFBinaryHeapRemoveAllValues, frameworkHandle, "CFBinaryHeapRemoveAllValues")
		registerFunc(&_cFBinaryHeapRemoveMinimumValue, frameworkHandle, "CFBinaryHeapRemoveMinimumValue")
		registerFunc(&_cFBitVectorContainsBit, frameworkHandle, "CFBitVectorContainsBit")
		registerFunc(&_cFBitVectorCreate, frameworkHandle, "CFBitVectorCreate")
		registerFunc(&_cFBitVectorCreateCopy, frameworkHandle, "CFBitVectorCreateCopy")
		registerFunc(&_cFBitVectorCreateMutable, frameworkHandle, "CFBitVectorCreateMutable")
		registerFunc(&_cFBitVectorCreateMutableCopy, frameworkHandle, "CFBitVectorCreateMutableCopy")
		registerFunc(&_cFBitVectorFlipBitAtIndex, frameworkHandle, "CFBitVectorFlipBitAtIndex")
		registerFunc(&_cFBitVectorFlipBits, frameworkHandle, "CFBitVectorFlipBits")
		registerFunc(&_cFBitVectorGetBitAtIndex, frameworkHandle, "CFBitVectorGetBitAtIndex")
		registerFunc(&_cFBitVectorGetBits, frameworkHandle, "CFBitVectorGetBits")
		registerFunc(&_cFBitVectorGetCount, frameworkHandle, "CFBitVectorGetCount")
		registerFunc(&_cFBitVectorGetCountOfBit, frameworkHandle, "CFBitVectorGetCountOfBit")
		registerFunc(&_cFBitVectorGetFirstIndexOfBit, frameworkHandle, "CFBitVectorGetFirstIndexOfBit")
		registerFunc(&_cFBitVectorGetLastIndexOfBit, frameworkHandle, "CFBitVectorGetLastIndexOfBit")
		registerFunc(&_cFBitVectorGetTypeID, frameworkHandle, "CFBitVectorGetTypeID")
		registerFunc(&_cFBitVectorSetAllBits, frameworkHandle, "CFBitVectorSetAllBits")
		registerFunc(&_cFBitVectorSetBitAtIndex, frameworkHandle, "CFBitVectorSetBitAtIndex")
		registerFunc(&_cFBitVectorSetBits, frameworkHandle, "CFBitVectorSetBits")
		registerFunc(&_cFBitVectorSetCount, frameworkHandle, "CFBitVectorSetCount")
		registerFunc(&_cFBooleanGetTypeID, frameworkHandle, "CFBooleanGetTypeID")
		registerFunc(&_cFBooleanGetValue, frameworkHandle, "CFBooleanGetValue")
		registerFunc(&_cFBundleCopyAuxiliaryExecutableURL, frameworkHandle, "CFBundleCopyAuxiliaryExecutableURL")
		registerFunc(&_cFBundleCopyBuiltInPlugInsURL, frameworkHandle, "CFBundleCopyBuiltInPlugInsURL")
		registerFunc(&_cFBundleCopyBundleLocalizations, frameworkHandle, "CFBundleCopyBundleLocalizations")
		registerFunc(&_cFBundleCopyBundleURL, frameworkHandle, "CFBundleCopyBundleURL")
		registerFunc(&_cFBundleCopyExecutableArchitectures, frameworkHandle, "CFBundleCopyExecutableArchitectures")
		registerFunc(&_cFBundleCopyExecutableArchitecturesForURL, frameworkHandle, "CFBundleCopyExecutableArchitecturesForURL")
		registerFunc(&_cFBundleCopyExecutableURL, frameworkHandle, "CFBundleCopyExecutableURL")
		registerFunc(&_cFBundleCopyInfoDictionaryForURL, frameworkHandle, "CFBundleCopyInfoDictionaryForURL")
		registerFunc(&_cFBundleCopyInfoDictionaryInDirectory, frameworkHandle, "CFBundleCopyInfoDictionaryInDirectory")
		registerFunc(&_cFBundleCopyLocalizationsForPreferences, frameworkHandle, "CFBundleCopyLocalizationsForPreferences")
		registerFunc(&_cFBundleCopyLocalizationsForURL, frameworkHandle, "CFBundleCopyLocalizationsForURL")
		registerFunc(&_cFBundleCopyLocalizedString, frameworkHandle, "CFBundleCopyLocalizedString")
		registerFunc(&_cFBundleCopyLocalizedStringForLocalizations, frameworkHandle, "CFBundleCopyLocalizedStringForLocalizations")
		registerFunc(&_cFBundleCopyPreferredLocalizationsFromArray, frameworkHandle, "CFBundleCopyPreferredLocalizationsFromArray")
		registerFunc(&_cFBundleCopyPrivateFrameworksURL, frameworkHandle, "CFBundleCopyPrivateFrameworksURL")
		registerFunc(&_cFBundleCopyResourceURL, frameworkHandle, "CFBundleCopyResourceURL")
		registerFunc(&_cFBundleCopyResourceURLForLocalization, frameworkHandle, "CFBundleCopyResourceURLForLocalization")
		registerFunc(&_cFBundleCopyResourceURLInDirectory, frameworkHandle, "CFBundleCopyResourceURLInDirectory")
		registerFunc(&_cFBundleCopyResourceURLsOfType, frameworkHandle, "CFBundleCopyResourceURLsOfType")
		registerFunc(&_cFBundleCopyResourceURLsOfTypeForLocalization, frameworkHandle, "CFBundleCopyResourceURLsOfTypeForLocalization")
		registerFunc(&_cFBundleCopyResourceURLsOfTypeInDirectory, frameworkHandle, "CFBundleCopyResourceURLsOfTypeInDirectory")
		registerFunc(&_cFBundleCopyResourcesDirectoryURL, frameworkHandle, "CFBundleCopyResourcesDirectoryURL")
		registerFunc(&_cFBundleCopySharedFrameworksURL, frameworkHandle, "CFBundleCopySharedFrameworksURL")
		registerFunc(&_cFBundleCopySharedSupportURL, frameworkHandle, "CFBundleCopySharedSupportURL")
		registerFunc(&_cFBundleCopySupportFilesDirectoryURL, frameworkHandle, "CFBundleCopySupportFilesDirectoryURL")
		registerFunc(&_cFBundleCreate, frameworkHandle, "CFBundleCreate")
		registerFunc(&_cFBundleCreateBundlesFromDirectory, frameworkHandle, "CFBundleCreateBundlesFromDirectory")
		registerFunc(&_cFBundleGetAllBundles, frameworkHandle, "CFBundleGetAllBundles")
		registerFunc(&_cFBundleGetBundleWithIdentifier, frameworkHandle, "CFBundleGetBundleWithIdentifier")
		registerFunc(&_cFBundleGetDataPointerForName, frameworkHandle, "CFBundleGetDataPointerForName")
		registerFunc(&_cFBundleGetDataPointersForNames, frameworkHandle, "CFBundleGetDataPointersForNames")
		registerFunc(&_cFBundleGetDevelopmentRegion, frameworkHandle, "CFBundleGetDevelopmentRegion")
		registerFunc(&_cFBundleGetFunctionPointerForName, frameworkHandle, "CFBundleGetFunctionPointerForName")
		registerFunc(&_cFBundleGetFunctionPointersForNames, frameworkHandle, "CFBundleGetFunctionPointersForNames")
		registerFunc(&_cFBundleGetIdentifier, frameworkHandle, "CFBundleGetIdentifier")
		registerFunc(&_cFBundleGetInfoDictionary, frameworkHandle, "CFBundleGetInfoDictionary")
		registerFunc(&_cFBundleGetLocalInfoDictionary, frameworkHandle, "CFBundleGetLocalInfoDictionary")
		registerFunc(&_cFBundleGetMainBundle, frameworkHandle, "CFBundleGetMainBundle")
		registerFunc(&_cFBundleGetPackageInfo, frameworkHandle, "CFBundleGetPackageInfo")
		registerFunc(&_cFBundleGetPackageInfoInDirectory, frameworkHandle, "CFBundleGetPackageInfoInDirectory")
		registerFunc(&_cFBundleGetPlugIn, frameworkHandle, "CFBundleGetPlugIn")
		registerFunc(&_cFBundleGetTypeID, frameworkHandle, "CFBundleGetTypeID")
		registerFunc(&_cFBundleGetValueForInfoDictionaryKey, frameworkHandle, "CFBundleGetValueForInfoDictionaryKey")
		registerFunc(&_cFBundleGetVersionNumber, frameworkHandle, "CFBundleGetVersionNumber")
		registerFunc(&_cFBundleIsArchitectureLoadable, frameworkHandle, "CFBundleIsArchitectureLoadable")
		registerFunc(&_cFBundleIsExecutableLoadable, frameworkHandle, "CFBundleIsExecutableLoadable")
		registerFunc(&_cFBundleIsExecutableLoadableForURL, frameworkHandle, "CFBundleIsExecutableLoadableForURL")
		registerFunc(&_cFBundleIsExecutableLoaded, frameworkHandle, "CFBundleIsExecutableLoaded")
		registerFunc(&_cFBundleLoadExecutable, frameworkHandle, "CFBundleLoadExecutable")
		registerFunc(&_cFBundleLoadExecutableAndReturnError, frameworkHandle, "CFBundleLoadExecutableAndReturnError")
		registerFunc(&_cFBundlePreflightExecutable, frameworkHandle, "CFBundlePreflightExecutable")
		registerFunc(&_cFBundleUnloadExecutable, frameworkHandle, "CFBundleUnloadExecutable")
		registerFunc(&_cFCalendarAddComponents, frameworkHandle, "CFCalendarAddComponents")
		registerFunc(&_cFCalendarComposeAbsoluteTime, frameworkHandle, "CFCalendarComposeAbsoluteTime")
		registerFunc(&_cFCalendarCopyCurrent, frameworkHandle, "CFCalendarCopyCurrent")
		registerFunc(&_cFCalendarCopyLocale, frameworkHandle, "CFCalendarCopyLocale")
		registerFunc(&_cFCalendarCopyTimeZone, frameworkHandle, "CFCalendarCopyTimeZone")
		registerFunc(&_cFCalendarCreateWithIdentifier, frameworkHandle, "CFCalendarCreateWithIdentifier")
		registerFunc(&_cFCalendarDecomposeAbsoluteTime, frameworkHandle, "CFCalendarDecomposeAbsoluteTime")
		registerFunc(&_cFCalendarGetComponentDifference, frameworkHandle, "CFCalendarGetComponentDifference")
		registerFunc(&_cFCalendarGetFirstWeekday, frameworkHandle, "CFCalendarGetFirstWeekday")
		registerFunc(&_cFCalendarGetIdentifier, frameworkHandle, "CFCalendarGetIdentifier")
		registerFunc(&_cFCalendarGetMaximumRangeOfUnit, frameworkHandle, "CFCalendarGetMaximumRangeOfUnit")
		registerFunc(&_cFCalendarGetMinimumDaysInFirstWeek, frameworkHandle, "CFCalendarGetMinimumDaysInFirstWeek")
		registerFunc(&_cFCalendarGetMinimumRangeOfUnit, frameworkHandle, "CFCalendarGetMinimumRangeOfUnit")
		registerFunc(&_cFCalendarGetOrdinalityOfUnit, frameworkHandle, "CFCalendarGetOrdinalityOfUnit")
		registerFunc(&_cFCalendarGetRangeOfUnit, frameworkHandle, "CFCalendarGetRangeOfUnit")
		registerFunc(&_cFCalendarGetTimeRangeOfUnit, frameworkHandle, "CFCalendarGetTimeRangeOfUnit")
		registerFunc(&_cFCalendarGetTypeID, frameworkHandle, "CFCalendarGetTypeID")
		registerFunc(&_cFCalendarSetFirstWeekday, frameworkHandle, "CFCalendarSetFirstWeekday")
		registerFunc(&_cFCalendarSetLocale, frameworkHandle, "CFCalendarSetLocale")
		registerFunc(&_cFCalendarSetMinimumDaysInFirstWeek, frameworkHandle, "CFCalendarSetMinimumDaysInFirstWeek")
		registerFunc(&_cFCalendarSetTimeZone, frameworkHandle, "CFCalendarSetTimeZone")
		registerFunc(&_cFCharacterSetAddCharactersInRange, frameworkHandle, "CFCharacterSetAddCharactersInRange")
		registerFunc(&_cFCharacterSetAddCharactersInString, frameworkHandle, "CFCharacterSetAddCharactersInString")
		registerFunc(&_cFCharacterSetCreateBitmapRepresentation, frameworkHandle, "CFCharacterSetCreateBitmapRepresentation")
		registerFunc(&_cFCharacterSetCreateCopy, frameworkHandle, "CFCharacterSetCreateCopy")
		registerFunc(&_cFCharacterSetCreateInvertedSet, frameworkHandle, "CFCharacterSetCreateInvertedSet")
		registerFunc(&_cFCharacterSetCreateMutable, frameworkHandle, "CFCharacterSetCreateMutable")
		registerFunc(&_cFCharacterSetCreateMutableCopy, frameworkHandle, "CFCharacterSetCreateMutableCopy")
		registerFunc(&_cFCharacterSetCreateWithBitmapRepresentation, frameworkHandle, "CFCharacterSetCreateWithBitmapRepresentation")
		registerFunc(&_cFCharacterSetCreateWithCharactersInRange, frameworkHandle, "CFCharacterSetCreateWithCharactersInRange")
		registerFunc(&_cFCharacterSetCreateWithCharactersInString, frameworkHandle, "CFCharacterSetCreateWithCharactersInString")
		registerFunc(&_cFCharacterSetGetPredefined, frameworkHandle, "CFCharacterSetGetPredefined")
		registerFunc(&_cFCharacterSetGetTypeID, frameworkHandle, "CFCharacterSetGetTypeID")
		registerFunc(&_cFCharacterSetHasMemberInPlane, frameworkHandle, "CFCharacterSetHasMemberInPlane")
		registerFunc(&_cFCharacterSetIntersect, frameworkHandle, "CFCharacterSetIntersect")
		registerFunc(&_cFCharacterSetInvert, frameworkHandle, "CFCharacterSetInvert")
		registerFunc(&_cFCharacterSetIsCharacterMember, frameworkHandle, "CFCharacterSetIsCharacterMember")
		registerFunc(&_cFCharacterSetIsLongCharacterMember, frameworkHandle, "CFCharacterSetIsLongCharacterMember")
		registerFunc(&_cFCharacterSetIsSupersetOfSet, frameworkHandle, "CFCharacterSetIsSupersetOfSet")
		registerFunc(&_cFCharacterSetRemoveCharactersInRange, frameworkHandle, "CFCharacterSetRemoveCharactersInRange")
		registerFunc(&_cFCharacterSetRemoveCharactersInString, frameworkHandle, "CFCharacterSetRemoveCharactersInString")
		registerFunc(&_cFCharacterSetUnion, frameworkHandle, "CFCharacterSetUnion")
		registerFunc(&_cFCopyDescription, frameworkHandle, "CFCopyDescription")
		registerFunc(&_cFCopyTypeIDDescription, frameworkHandle, "CFCopyTypeIDDescription")
		registerFunc(&_cFDataAppendBytes, frameworkHandle, "CFDataAppendBytes")
		registerFunc(&_cFDataCreate, frameworkHandle, "CFDataCreate")
		registerFunc(&_cFDataCreateCopy, frameworkHandle, "CFDataCreateCopy")
		registerFunc(&_cFDataCreateMutable, frameworkHandle, "CFDataCreateMutable")
		registerFunc(&_cFDataCreateMutableCopy, frameworkHandle, "CFDataCreateMutableCopy")
		registerFunc(&_cFDataCreateWithBytesNoCopy, frameworkHandle, "CFDataCreateWithBytesNoCopy")
		registerFunc(&_cFDataDeleteBytes, frameworkHandle, "CFDataDeleteBytes")
		registerFunc(&_cFDataFind, frameworkHandle, "CFDataFind")
		registerFunc(&_cFDataGetBytePtr, frameworkHandle, "CFDataGetBytePtr")
		registerFunc(&_cFDataGetBytes, frameworkHandle, "CFDataGetBytes")
		registerFunc(&_cFDataGetLength, frameworkHandle, "CFDataGetLength")
		registerFunc(&_cFDataGetMutableBytePtr, frameworkHandle, "CFDataGetMutableBytePtr")
		registerFunc(&_cFDataGetTypeID, frameworkHandle, "CFDataGetTypeID")
		registerFunc(&_cFDataIncreaseLength, frameworkHandle, "CFDataIncreaseLength")
		registerFunc(&_cFDataReplaceBytes, frameworkHandle, "CFDataReplaceBytes")
		registerFunc(&_cFDataSetLength, frameworkHandle, "CFDataSetLength")
		registerFunc(&_cFDateCompare, frameworkHandle, "CFDateCompare")
		registerFunc(&_cFDateCreate, frameworkHandle, "CFDateCreate")
		registerFunc(&_cFDateFormatterCopyProperty, frameworkHandle, "CFDateFormatterCopyProperty")
		registerFunc(&_cFDateFormatterCreate, frameworkHandle, "CFDateFormatterCreate")
		registerFunc(&_cFDateFormatterCreateDateFormatFromTemplate, frameworkHandle, "CFDateFormatterCreateDateFormatFromTemplate")
		registerFunc(&_cFDateFormatterCreateDateFromString, frameworkHandle, "CFDateFormatterCreateDateFromString")
		registerFunc(&_cFDateFormatterCreateISO8601Formatter, frameworkHandle, "CFDateFormatterCreateISO8601Formatter")
		registerFunc(&_cFDateFormatterCreateStringWithAbsoluteTime, frameworkHandle, "CFDateFormatterCreateStringWithAbsoluteTime")
		registerFunc(&_cFDateFormatterCreateStringWithDate, frameworkHandle, "CFDateFormatterCreateStringWithDate")
		registerFunc(&_cFDateFormatterGetAbsoluteTimeFromString, frameworkHandle, "CFDateFormatterGetAbsoluteTimeFromString")
		registerFunc(&_cFDateFormatterGetDateStyle, frameworkHandle, "CFDateFormatterGetDateStyle")
		registerFunc(&_cFDateFormatterGetFormat, frameworkHandle, "CFDateFormatterGetFormat")
		registerFunc(&_cFDateFormatterGetLocale, frameworkHandle, "CFDateFormatterGetLocale")
		registerFunc(&_cFDateFormatterGetTimeStyle, frameworkHandle, "CFDateFormatterGetTimeStyle")
		registerFunc(&_cFDateFormatterGetTypeID, frameworkHandle, "CFDateFormatterGetTypeID")
		registerFunc(&_cFDateFormatterSetFormat, frameworkHandle, "CFDateFormatterSetFormat")
		registerFunc(&_cFDateFormatterSetProperty, frameworkHandle, "CFDateFormatterSetProperty")
		registerFunc(&_cFDateGetAbsoluteTime, frameworkHandle, "CFDateGetAbsoluteTime")
		registerFunc(&_cFDateGetTimeIntervalSinceDate, frameworkHandle, "CFDateGetTimeIntervalSinceDate")
		registerFunc(&_cFDateGetTypeID, frameworkHandle, "CFDateGetTypeID")
		registerFunc(&_cFDictionaryAddValue, frameworkHandle, "CFDictionaryAddValue")
		registerFunc(&_cFDictionaryApplyFunction, frameworkHandle, "CFDictionaryApplyFunction")
		registerFunc(&_cFDictionaryContainsKey, frameworkHandle, "CFDictionaryContainsKey")
		registerFunc(&_cFDictionaryContainsValue, frameworkHandle, "CFDictionaryContainsValue")
		registerFunc(&_cFDictionaryCreate, frameworkHandle, "CFDictionaryCreate")
		registerFunc(&_cFDictionaryCreateCopy, frameworkHandle, "CFDictionaryCreateCopy")
		registerFunc(&_cFDictionaryCreateMutable, frameworkHandle, "CFDictionaryCreateMutable")
		registerFunc(&_cFDictionaryCreateMutableCopy, frameworkHandle, "CFDictionaryCreateMutableCopy")
		registerFunc(&_cFDictionaryGetCount, frameworkHandle, "CFDictionaryGetCount")
		registerFunc(&_cFDictionaryGetCountOfKey, frameworkHandle, "CFDictionaryGetCountOfKey")
		registerFunc(&_cFDictionaryGetCountOfValue, frameworkHandle, "CFDictionaryGetCountOfValue")
		registerFunc(&_cFDictionaryGetKeysAndValues, frameworkHandle, "CFDictionaryGetKeysAndValues")
		registerFunc(&_cFDictionaryGetTypeID, frameworkHandle, "CFDictionaryGetTypeID")
		registerFunc(&_cFDictionaryGetValue, frameworkHandle, "CFDictionaryGetValue")
		registerFunc(&_cFDictionaryGetValueIfPresent, frameworkHandle, "CFDictionaryGetValueIfPresent")
		registerFunc(&_cFDictionaryRemoveAllValues, frameworkHandle, "CFDictionaryRemoveAllValues")
		registerFunc(&_cFDictionaryRemoveValue, frameworkHandle, "CFDictionaryRemoveValue")
		registerFunc(&_cFDictionaryReplaceValue, frameworkHandle, "CFDictionaryReplaceValue")
		registerFunc(&_cFDictionarySetValue, frameworkHandle, "CFDictionarySetValue")
		registerFunc(&_cFEqual, frameworkHandle, "CFEqual")
		registerFunc(&_cFErrorCopyDescription, frameworkHandle, "CFErrorCopyDescription")
		registerFunc(&_cFErrorCopyFailureReason, frameworkHandle, "CFErrorCopyFailureReason")
		registerFunc(&_cFErrorCopyRecoverySuggestion, frameworkHandle, "CFErrorCopyRecoverySuggestion")
		registerFunc(&_cFErrorCopyUserInfo, frameworkHandle, "CFErrorCopyUserInfo")
		registerFunc(&_cFErrorCreate, frameworkHandle, "CFErrorCreate")
		registerFunc(&_cFErrorCreateWithUserInfoKeysAndValues, frameworkHandle, "CFErrorCreateWithUserInfoKeysAndValues")
		registerFunc(&_cFErrorGetCode, frameworkHandle, "CFErrorGetCode")
		registerFunc(&_cFErrorGetDomain, frameworkHandle, "CFErrorGetDomain")
		registerFunc(&_cFErrorGetTypeID, frameworkHandle, "CFErrorGetTypeID")
		registerFunc(&_cFFileDescriptorCreate, frameworkHandle, "CFFileDescriptorCreate")
		registerFunc(&_cFFileDescriptorCreateRunLoopSource, frameworkHandle, "CFFileDescriptorCreateRunLoopSource")
		registerFunc(&_cFFileDescriptorDisableCallBacks, frameworkHandle, "CFFileDescriptorDisableCallBacks")
		registerFunc(&_cFFileDescriptorEnableCallBacks, frameworkHandle, "CFFileDescriptorEnableCallBacks")
		registerFunc(&_cFFileDescriptorGetContext, frameworkHandle, "CFFileDescriptorGetContext")
		registerFunc(&_cFFileDescriptorGetNativeDescriptor, frameworkHandle, "CFFileDescriptorGetNativeDescriptor")
		registerFunc(&_cFFileDescriptorGetTypeID, frameworkHandle, "CFFileDescriptorGetTypeID")
		registerFunc(&_cFFileDescriptorInvalidate, frameworkHandle, "CFFileDescriptorInvalidate")
		registerFunc(&_cFFileDescriptorIsValid, frameworkHandle, "CFFileDescriptorIsValid")
		registerFunc(&_cFFileSecurityClearProperties, frameworkHandle, "CFFileSecurityClearProperties")
		registerFunc(&_cFFileSecurityCopyGroupUUID, frameworkHandle, "CFFileSecurityCopyGroupUUID")
		registerFunc(&_cFFileSecurityCopyOwnerUUID, frameworkHandle, "CFFileSecurityCopyOwnerUUID")
		registerFunc(&_cFFileSecurityCreate, frameworkHandle, "CFFileSecurityCreate")
		registerFunc(&_cFFileSecurityCreateCopy, frameworkHandle, "CFFileSecurityCreateCopy")
		registerFunc(&_cFFileSecurityGetGroup, frameworkHandle, "CFFileSecurityGetGroup")
		registerFunc(&_cFFileSecurityGetMode, frameworkHandle, "CFFileSecurityGetMode")
		registerFunc(&_cFFileSecurityGetOwner, frameworkHandle, "CFFileSecurityGetOwner")
		registerFunc(&_cFFileSecurityGetTypeID, frameworkHandle, "CFFileSecurityGetTypeID")
		registerFunc(&_cFFileSecuritySetAccessControlList, frameworkHandle, "CFFileSecuritySetAccessControlList")
		registerFunc(&_cFFileSecuritySetGroup, frameworkHandle, "CFFileSecuritySetGroup")
		registerFunc(&_cFFileSecuritySetGroupUUID, frameworkHandle, "CFFileSecuritySetGroupUUID")
		registerFunc(&_cFFileSecuritySetMode, frameworkHandle, "CFFileSecuritySetMode")
		registerFunc(&_cFFileSecuritySetOwner, frameworkHandle, "CFFileSecuritySetOwner")
		registerFunc(&_cFFileSecuritySetOwnerUUID, frameworkHandle, "CFFileSecuritySetOwnerUUID")
		registerFunc(&_cFGetAllocator, frameworkHandle, "CFGetAllocator")
		registerFunc(&_cFGetRetainCount, frameworkHandle, "CFGetRetainCount")
		registerFunc(&_cFGetTypeID, frameworkHandle, "CFGetTypeID")
		registerFunc(&_cFHash, frameworkHandle, "CFHash")
		registerFunc(&_cFLocaleCopyAvailableLocaleIdentifiers, frameworkHandle, "CFLocaleCopyAvailableLocaleIdentifiers")
		registerFunc(&_cFLocaleCopyCommonISOCurrencyCodes, frameworkHandle, "CFLocaleCopyCommonISOCurrencyCodes")
		registerFunc(&_cFLocaleCopyCurrent, frameworkHandle, "CFLocaleCopyCurrent")
		registerFunc(&_cFLocaleCopyDisplayNameForPropertyValue, frameworkHandle, "CFLocaleCopyDisplayNameForPropertyValue")
		registerFunc(&_cFLocaleCopyISOCountryCodes, frameworkHandle, "CFLocaleCopyISOCountryCodes")
		registerFunc(&_cFLocaleCopyISOCurrencyCodes, frameworkHandle, "CFLocaleCopyISOCurrencyCodes")
		registerFunc(&_cFLocaleCopyISOLanguageCodes, frameworkHandle, "CFLocaleCopyISOLanguageCodes")
		registerFunc(&_cFLocaleCopyPreferredLanguages, frameworkHandle, "CFLocaleCopyPreferredLanguages")
		registerFunc(&_cFLocaleCreate, frameworkHandle, "CFLocaleCreate")
		registerFunc(&_cFLocaleCreateCanonicalLanguageIdentifierFromString, frameworkHandle, "CFLocaleCreateCanonicalLanguageIdentifierFromString")
		registerFunc(&_cFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes, frameworkHandle, "CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes")
		registerFunc(&_cFLocaleCreateCanonicalLocaleIdentifierFromString, frameworkHandle, "CFLocaleCreateCanonicalLocaleIdentifierFromString")
		registerFunc(&_cFLocaleCreateComponentsFromLocaleIdentifier, frameworkHandle, "CFLocaleCreateComponentsFromLocaleIdentifier")
		registerFunc(&_cFLocaleCreateCopy, frameworkHandle, "CFLocaleCreateCopy")
		registerFunc(&_cFLocaleCreateLocaleIdentifierFromComponents, frameworkHandle, "CFLocaleCreateLocaleIdentifierFromComponents")
		registerFunc(&_cFLocaleCreateLocaleIdentifierFromWindowsLocaleCode, frameworkHandle, "CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode")
		registerFunc(&_cFLocaleGetIdentifier, frameworkHandle, "CFLocaleGetIdentifier")
		registerFunc(&_cFLocaleGetLanguageCharacterDirection, frameworkHandle, "CFLocaleGetLanguageCharacterDirection")
		registerFunc(&_cFLocaleGetLanguageLineDirection, frameworkHandle, "CFLocaleGetLanguageLineDirection")
		registerFunc(&_cFLocaleGetSystem, frameworkHandle, "CFLocaleGetSystem")
		registerFunc(&_cFLocaleGetTypeID, frameworkHandle, "CFLocaleGetTypeID")
		registerFunc(&_cFLocaleGetValue, frameworkHandle, "CFLocaleGetValue")
		registerFunc(&_cFLocaleGetWindowsLocaleCodeFromLocaleIdentifier, frameworkHandle, "CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier")
		registerFunc(&_cFMachPortCreate, frameworkHandle, "CFMachPortCreate")
		registerFunc(&_cFMachPortCreateRunLoopSource, frameworkHandle, "CFMachPortCreateRunLoopSource")
		registerFunc(&_cFMachPortCreateWithPort, frameworkHandle, "CFMachPortCreateWithPort")
		registerFunc(&_cFMachPortGetContext, frameworkHandle, "CFMachPortGetContext")
		registerFunc(&_cFMachPortGetInvalidationCallBack, frameworkHandle, "CFMachPortGetInvalidationCallBack")
		registerFunc(&_cFMachPortGetPort, frameworkHandle, "CFMachPortGetPort")
		registerFunc(&_cFMachPortGetTypeID, frameworkHandle, "CFMachPortGetTypeID")
		registerFunc(&_cFMachPortInvalidate, frameworkHandle, "CFMachPortInvalidate")
		registerFunc(&_cFMachPortIsValid, frameworkHandle, "CFMachPortIsValid")
		registerFunc(&_cFMachPortSetInvalidationCallBack, frameworkHandle, "CFMachPortSetInvalidationCallBack")
		registerFunc(&_cFMakeCollectable, frameworkHandle, "CFMakeCollectable")
		registerFunc(&_cFMessagePortCreateLocal, frameworkHandle, "CFMessagePortCreateLocal")
		registerFunc(&_cFMessagePortCreateRemote, frameworkHandle, "CFMessagePortCreateRemote")
		registerFunc(&_cFMessagePortCreateRunLoopSource, frameworkHandle, "CFMessagePortCreateRunLoopSource")
		registerFunc(&_cFMessagePortGetContext, frameworkHandle, "CFMessagePortGetContext")
		registerFunc(&_cFMessagePortGetInvalidationCallBack, frameworkHandle, "CFMessagePortGetInvalidationCallBack")
		registerFunc(&_cFMessagePortGetName, frameworkHandle, "CFMessagePortGetName")
		registerFunc(&_cFMessagePortGetTypeID, frameworkHandle, "CFMessagePortGetTypeID")
		registerFunc(&_cFMessagePortInvalidate, frameworkHandle, "CFMessagePortInvalidate")
		registerFunc(&_cFMessagePortIsRemote, frameworkHandle, "CFMessagePortIsRemote")
		registerFunc(&_cFMessagePortIsValid, frameworkHandle, "CFMessagePortIsValid")
		registerFunc(&_cFMessagePortSendRequest, frameworkHandle, "CFMessagePortSendRequest")
		registerFunc(&_cFMessagePortSetDispatchQueue, frameworkHandle, "CFMessagePortSetDispatchQueue")
		registerFunc(&_cFMessagePortSetInvalidationCallBack, frameworkHandle, "CFMessagePortSetInvalidationCallBack")
		registerFunc(&_cFMessagePortSetName, frameworkHandle, "CFMessagePortSetName")
		registerFunc(&_cFNotificationCenterAddObserver, frameworkHandle, "CFNotificationCenterAddObserver")
		registerFunc(&_cFNotificationCenterGetDarwinNotifyCenter, frameworkHandle, "CFNotificationCenterGetDarwinNotifyCenter")
		registerFunc(&_cFNotificationCenterGetDistributedCenter, frameworkHandle, "CFNotificationCenterGetDistributedCenter")
		registerFunc(&_cFNotificationCenterGetLocalCenter, frameworkHandle, "CFNotificationCenterGetLocalCenter")
		registerFunc(&_cFNotificationCenterGetTypeID, frameworkHandle, "CFNotificationCenterGetTypeID")
		registerFunc(&_cFNotificationCenterPostNotification, frameworkHandle, "CFNotificationCenterPostNotification")
		registerFunc(&_cFNotificationCenterPostNotificationWithOptions, frameworkHandle, "CFNotificationCenterPostNotificationWithOptions")
		registerFunc(&_cFNotificationCenterRemoveEveryObserver, frameworkHandle, "CFNotificationCenterRemoveEveryObserver")
		registerFunc(&_cFNotificationCenterRemoveObserver, frameworkHandle, "CFNotificationCenterRemoveObserver")
		registerFunc(&_cFNullGetTypeID, frameworkHandle, "CFNullGetTypeID")
		registerFunc(&_cFNumberCompare, frameworkHandle, "CFNumberCompare")
		registerFunc(&_cFNumberCreate, frameworkHandle, "CFNumberCreate")
		registerFunc(&_cFNumberFormatterCopyProperty, frameworkHandle, "CFNumberFormatterCopyProperty")
		registerFunc(&_cFNumberFormatterCreate, frameworkHandle, "CFNumberFormatterCreate")
		registerFunc(&_cFNumberFormatterCreateNumberFromString, frameworkHandle, "CFNumberFormatterCreateNumberFromString")
		registerFunc(&_cFNumberFormatterCreateStringWithNumber, frameworkHandle, "CFNumberFormatterCreateStringWithNumber")
		registerFunc(&_cFNumberFormatterCreateStringWithValue, frameworkHandle, "CFNumberFormatterCreateStringWithValue")
		registerFunc(&_cFNumberFormatterGetDecimalInfoForCurrencyCode, frameworkHandle, "CFNumberFormatterGetDecimalInfoForCurrencyCode")
		registerFunc(&_cFNumberFormatterGetFormat, frameworkHandle, "CFNumberFormatterGetFormat")
		registerFunc(&_cFNumberFormatterGetLocale, frameworkHandle, "CFNumberFormatterGetLocale")
		registerFunc(&_cFNumberFormatterGetStyle, frameworkHandle, "CFNumberFormatterGetStyle")
		registerFunc(&_cFNumberFormatterGetTypeID, frameworkHandle, "CFNumberFormatterGetTypeID")
		registerFunc(&_cFNumberFormatterGetValueFromString, frameworkHandle, "CFNumberFormatterGetValueFromString")
		registerFunc(&_cFNumberFormatterSetFormat, frameworkHandle, "CFNumberFormatterSetFormat")
		registerFunc(&_cFNumberFormatterSetProperty, frameworkHandle, "CFNumberFormatterSetProperty")
		registerFunc(&_cFNumberGetByteSize, frameworkHandle, "CFNumberGetByteSize")
		registerFunc(&_cFNumberGetType, frameworkHandle, "CFNumberGetType")
		registerFunc(&_cFNumberGetTypeID, frameworkHandle, "CFNumberGetTypeID")
		registerFunc(&_cFNumberGetValue, frameworkHandle, "CFNumberGetValue")
		registerFunc(&_cFNumberIsFloatType, frameworkHandle, "CFNumberIsFloatType")
		registerFunc(&_cFPlugInAddInstanceForFactory, frameworkHandle, "CFPlugInAddInstanceForFactory")
		registerFunc(&_cFPlugInCreate, frameworkHandle, "CFPlugInCreate")
		registerFunc(&_cFPlugInFindFactoriesForPlugInType, frameworkHandle, "CFPlugInFindFactoriesForPlugInType")
		registerFunc(&_cFPlugInFindFactoriesForPlugInTypeInPlugIn, frameworkHandle, "CFPlugInFindFactoriesForPlugInTypeInPlugIn")
		registerFunc(&_cFPlugInGetBundle, frameworkHandle, "CFPlugInGetBundle")
		registerFunc(&_cFPlugInGetTypeID, frameworkHandle, "CFPlugInGetTypeID")
		registerFunc(&_cFPlugInInstanceCreate, frameworkHandle, "CFPlugInInstanceCreate")
		registerFunc(&_cFPlugInInstanceCreateWithInstanceDataSize, frameworkHandle, "CFPlugInInstanceCreateWithInstanceDataSize")
		registerFunc(&_cFPlugInInstanceGetFactoryName, frameworkHandle, "CFPlugInInstanceGetFactoryName")
		registerFunc(&_cFPlugInInstanceGetInstanceData, frameworkHandle, "CFPlugInInstanceGetInstanceData")
		registerFunc(&_cFPlugInInstanceGetInterfaceFunctionTable, frameworkHandle, "CFPlugInInstanceGetInterfaceFunctionTable")
		registerFunc(&_cFPlugInInstanceGetTypeID, frameworkHandle, "CFPlugInInstanceGetTypeID")
		registerFunc(&_cFPlugInIsLoadOnDemand, frameworkHandle, "CFPlugInIsLoadOnDemand")
		registerFunc(&_cFPlugInRegisterFactoryFunction, frameworkHandle, "CFPlugInRegisterFactoryFunction")
		registerFunc(&_cFPlugInRegisterFactoryFunctionByName, frameworkHandle, "CFPlugInRegisterFactoryFunctionByName")
		registerFunc(&_cFPlugInRegisterPlugInType, frameworkHandle, "CFPlugInRegisterPlugInType")
		registerFunc(&_cFPlugInRemoveInstanceForFactory, frameworkHandle, "CFPlugInRemoveInstanceForFactory")
		registerFunc(&_cFPlugInSetLoadOnDemand, frameworkHandle, "CFPlugInSetLoadOnDemand")
		registerFunc(&_cFPlugInUnregisterFactory, frameworkHandle, "CFPlugInUnregisterFactory")
		registerFunc(&_cFPlugInUnregisterPlugInType, frameworkHandle, "CFPlugInUnregisterPlugInType")
		registerFunc(&_cFPreferencesAddSuitePreferencesToApp, frameworkHandle, "CFPreferencesAddSuitePreferencesToApp")
		registerFunc(&_cFPreferencesAppSynchronize, frameworkHandle, "CFPreferencesAppSynchronize")
		registerFunc(&_cFPreferencesAppValueIsForced, frameworkHandle, "CFPreferencesAppValueIsForced")
		registerFunc(&_cFPreferencesCopyAppValue, frameworkHandle, "CFPreferencesCopyAppValue")
		registerFunc(&_cFPreferencesCopyKeyList, frameworkHandle, "CFPreferencesCopyKeyList")
		registerFunc(&_cFPreferencesCopyMultiple, frameworkHandle, "CFPreferencesCopyMultiple")
		registerFunc(&_cFPreferencesCopyValue, frameworkHandle, "CFPreferencesCopyValue")
		registerFunc(&_cFPreferencesGetAppBooleanValue, frameworkHandle, "CFPreferencesGetAppBooleanValue")
		registerFunc(&_cFPreferencesGetAppIntegerValue, frameworkHandle, "CFPreferencesGetAppIntegerValue")
		registerFunc(&_cFPreferencesRemoveSuitePreferencesFromApp, frameworkHandle, "CFPreferencesRemoveSuitePreferencesFromApp")
		registerFunc(&_cFPreferencesSetAppValue, frameworkHandle, "CFPreferencesSetAppValue")
		registerFunc(&_cFPreferencesSetMultiple, frameworkHandle, "CFPreferencesSetMultiple")
		registerFunc(&_cFPreferencesSetValue, frameworkHandle, "CFPreferencesSetValue")
		registerFunc(&_cFPreferencesSynchronize, frameworkHandle, "CFPreferencesSynchronize")
		registerFunc(&_cFPropertyListCreateData, frameworkHandle, "CFPropertyListCreateData")
		registerFunc(&_cFPropertyListCreateDeepCopy, frameworkHandle, "CFPropertyListCreateDeepCopy")
		registerFunc(&_cFPropertyListCreateWithData, frameworkHandle, "CFPropertyListCreateWithData")
		registerFunc(&_cFPropertyListCreateWithStream, frameworkHandle, "CFPropertyListCreateWithStream")
		registerFunc(&_cFPropertyListIsValid, frameworkHandle, "CFPropertyListIsValid")
		registerFunc(&_cFPropertyListWrite, frameworkHandle, "CFPropertyListWrite")
		registerFunc(&_cFReadStreamClose, frameworkHandle, "CFReadStreamClose")
		registerFunc(&_cFReadStreamCopyDispatchQueue, frameworkHandle, "CFReadStreamCopyDispatchQueue")
		registerFunc(&_cFReadStreamCopyError, frameworkHandle, "CFReadStreamCopyError")
		registerFunc(&_cFReadStreamCopyProperty, frameworkHandle, "CFReadStreamCopyProperty")
		registerFunc(&_cFReadStreamCreateWithBytesNoCopy, frameworkHandle, "CFReadStreamCreateWithBytesNoCopy")
		registerFunc(&_cFReadStreamCreateWithFile, frameworkHandle, "CFReadStreamCreateWithFile")
		registerFunc(&_cFReadStreamGetBuffer, frameworkHandle, "CFReadStreamGetBuffer")
		registerFunc(&_cFReadStreamGetError, frameworkHandle, "CFReadStreamGetError")
		registerFunc(&_cFReadStreamGetStatus, frameworkHandle, "CFReadStreamGetStatus")
		registerFunc(&_cFReadStreamGetTypeID, frameworkHandle, "CFReadStreamGetTypeID")
		registerFunc(&_cFReadStreamHasBytesAvailable, frameworkHandle, "CFReadStreamHasBytesAvailable")
		registerFunc(&_cFReadStreamOpen, frameworkHandle, "CFReadStreamOpen")
		registerFunc(&_cFReadStreamRead, frameworkHandle, "CFReadStreamRead")
		registerFunc(&_cFReadStreamScheduleWithRunLoop, frameworkHandle, "CFReadStreamScheduleWithRunLoop")
		registerFunc(&_cFReadStreamSetClient, frameworkHandle, "CFReadStreamSetClient")
		registerFunc(&_cFReadStreamSetDispatchQueue, frameworkHandle, "CFReadStreamSetDispatchQueue")
		registerFunc(&_cFReadStreamSetProperty, frameworkHandle, "CFReadStreamSetProperty")
		registerFunc(&_cFReadStreamUnscheduleFromRunLoop, frameworkHandle, "CFReadStreamUnscheduleFromRunLoop")
		registerFunc(&_cFRelease, frameworkHandle, "CFRelease")
		registerFunc(&_cFRetain, frameworkHandle, "CFRetain")
		registerFunc(&_cFRunLoopAddCommonMode, frameworkHandle, "CFRunLoopAddCommonMode")
		registerFunc(&_cFRunLoopAddObserver, frameworkHandle, "CFRunLoopAddObserver")
		registerFunc(&_cFRunLoopAddSource, frameworkHandle, "CFRunLoopAddSource")
		registerFunc(&_cFRunLoopAddTimer, frameworkHandle, "CFRunLoopAddTimer")
		registerFunc(&_cFRunLoopContainsObserver, frameworkHandle, "CFRunLoopContainsObserver")
		registerFunc(&_cFRunLoopContainsSource, frameworkHandle, "CFRunLoopContainsSource")
		registerFunc(&_cFRunLoopContainsTimer, frameworkHandle, "CFRunLoopContainsTimer")
		registerFunc(&_cFRunLoopCopyAllModes, frameworkHandle, "CFRunLoopCopyAllModes")
		registerFunc(&_cFRunLoopCopyCurrentMode, frameworkHandle, "CFRunLoopCopyCurrentMode")
		registerFunc(&_cFRunLoopGetCurrent, frameworkHandle, "CFRunLoopGetCurrent")
		registerFunc(&_cFRunLoopGetMain, frameworkHandle, "CFRunLoopGetMain")
		registerFunc(&_cFRunLoopGetNextTimerFireDate, frameworkHandle, "CFRunLoopGetNextTimerFireDate")
		registerFunc(&_cFRunLoopGetTypeID, frameworkHandle, "CFRunLoopGetTypeID")
		registerFunc(&_cFRunLoopIsWaiting, frameworkHandle, "CFRunLoopIsWaiting")
		registerFunc(&_cFRunLoopObserverCreate, frameworkHandle, "CFRunLoopObserverCreate")
		registerFunc(&_cFRunLoopObserverCreateWithHandler, frameworkHandle, "CFRunLoopObserverCreateWithHandler")
		registerFunc(&_cFRunLoopObserverDoesRepeat, frameworkHandle, "CFRunLoopObserverDoesRepeat")
		registerFunc(&_cFRunLoopObserverGetActivities, frameworkHandle, "CFRunLoopObserverGetActivities")
		registerFunc(&_cFRunLoopObserverGetContext, frameworkHandle, "CFRunLoopObserverGetContext")
		registerFunc(&_cFRunLoopObserverGetOrder, frameworkHandle, "CFRunLoopObserverGetOrder")
		registerFunc(&_cFRunLoopObserverGetTypeID, frameworkHandle, "CFRunLoopObserverGetTypeID")
		registerFunc(&_cFRunLoopObserverInvalidate, frameworkHandle, "CFRunLoopObserverInvalidate")
		registerFunc(&_cFRunLoopObserverIsValid, frameworkHandle, "CFRunLoopObserverIsValid")
		registerFunc(&_cFRunLoopPerformBlock, frameworkHandle, "CFRunLoopPerformBlock")
		registerFunc(&_cFRunLoopRemoveObserver, frameworkHandle, "CFRunLoopRemoveObserver")
		registerFunc(&_cFRunLoopRemoveSource, frameworkHandle, "CFRunLoopRemoveSource")
		registerFunc(&_cFRunLoopRemoveTimer, frameworkHandle, "CFRunLoopRemoveTimer")
		registerFunc(&_cFRunLoopRun, frameworkHandle, "CFRunLoopRun")
		registerFunc(&_cFRunLoopRunInMode, frameworkHandle, "CFRunLoopRunInMode")
		registerFunc(&_cFRunLoopSourceCreate, frameworkHandle, "CFRunLoopSourceCreate")
		registerFunc(&_cFRunLoopSourceGetContext, frameworkHandle, "CFRunLoopSourceGetContext")
		registerFunc(&_cFRunLoopSourceGetOrder, frameworkHandle, "CFRunLoopSourceGetOrder")
		registerFunc(&_cFRunLoopSourceGetTypeID, frameworkHandle, "CFRunLoopSourceGetTypeID")
		registerFunc(&_cFRunLoopSourceInvalidate, frameworkHandle, "CFRunLoopSourceInvalidate")
		registerFunc(&_cFRunLoopSourceIsValid, frameworkHandle, "CFRunLoopSourceIsValid")
		registerFunc(&_cFRunLoopSourceSignal, frameworkHandle, "CFRunLoopSourceSignal")
		registerFunc(&_cFRunLoopStop, frameworkHandle, "CFRunLoopStop")
		registerFunc(&_cFRunLoopTimerCreate, frameworkHandle, "CFRunLoopTimerCreate")
		registerFunc(&_cFRunLoopTimerCreateWithHandler, frameworkHandle, "CFRunLoopTimerCreateWithHandler")
		registerFunc(&_cFRunLoopTimerDoesRepeat, frameworkHandle, "CFRunLoopTimerDoesRepeat")
		registerFunc(&_cFRunLoopTimerGetContext, frameworkHandle, "CFRunLoopTimerGetContext")
		registerFunc(&_cFRunLoopTimerGetInterval, frameworkHandle, "CFRunLoopTimerGetInterval")
		registerFunc(&_cFRunLoopTimerGetNextFireDate, frameworkHandle, "CFRunLoopTimerGetNextFireDate")
		registerFunc(&_cFRunLoopTimerGetOrder, frameworkHandle, "CFRunLoopTimerGetOrder")
		registerFunc(&_cFRunLoopTimerGetTolerance, frameworkHandle, "CFRunLoopTimerGetTolerance")
		registerFunc(&_cFRunLoopTimerGetTypeID, frameworkHandle, "CFRunLoopTimerGetTypeID")
		registerFunc(&_cFRunLoopTimerInvalidate, frameworkHandle, "CFRunLoopTimerInvalidate")
		registerFunc(&_cFRunLoopTimerIsValid, frameworkHandle, "CFRunLoopTimerIsValid")
		registerFunc(&_cFRunLoopTimerSetNextFireDate, frameworkHandle, "CFRunLoopTimerSetNextFireDate")
		registerFunc(&_cFRunLoopTimerSetTolerance, frameworkHandle, "CFRunLoopTimerSetTolerance")
		registerFunc(&_cFRunLoopWakeUp, frameworkHandle, "CFRunLoopWakeUp")
		registerFunc(&_cFSetAddValue, frameworkHandle, "CFSetAddValue")
		registerFunc(&_cFSetApplyFunction, frameworkHandle, "CFSetApplyFunction")
		registerFunc(&_cFSetContainsValue, frameworkHandle, "CFSetContainsValue")
		registerFunc(&_cFSetCreate, frameworkHandle, "CFSetCreate")
		registerFunc(&_cFSetCreateCopy, frameworkHandle, "CFSetCreateCopy")
		registerFunc(&_cFSetCreateMutable, frameworkHandle, "CFSetCreateMutable")
		registerFunc(&_cFSetCreateMutableCopy, frameworkHandle, "CFSetCreateMutableCopy")
		registerFunc(&_cFSetGetCount, frameworkHandle, "CFSetGetCount")
		registerFunc(&_cFSetGetCountOfValue, frameworkHandle, "CFSetGetCountOfValue")
		registerFunc(&_cFSetGetTypeID, frameworkHandle, "CFSetGetTypeID")
		registerFunc(&_cFSetGetValue, frameworkHandle, "CFSetGetValue")
		registerFunc(&_cFSetGetValueIfPresent, frameworkHandle, "CFSetGetValueIfPresent")
		registerFunc(&_cFSetGetValues, frameworkHandle, "CFSetGetValues")
		registerFunc(&_cFSetRemoveAllValues, frameworkHandle, "CFSetRemoveAllValues")
		registerFunc(&_cFSetRemoveValue, frameworkHandle, "CFSetRemoveValue")
		registerFunc(&_cFSetReplaceValue, frameworkHandle, "CFSetReplaceValue")
		registerFunc(&_cFSetSetValue, frameworkHandle, "CFSetSetValue")
		registerFunc(&_cFShow, frameworkHandle, "CFShow")
		registerFunc(&_cFShowStr, frameworkHandle, "CFShowStr")
		registerFunc(&_cFSocketConnectToAddress, frameworkHandle, "CFSocketConnectToAddress")
		registerFunc(&_cFSocketCopyAddress, frameworkHandle, "CFSocketCopyAddress")
		registerFunc(&_cFSocketCopyPeerAddress, frameworkHandle, "CFSocketCopyPeerAddress")
		registerFunc(&_cFSocketCopyRegisteredSocketSignature, frameworkHandle, "CFSocketCopyRegisteredSocketSignature")
		registerFunc(&_cFSocketCopyRegisteredValue, frameworkHandle, "CFSocketCopyRegisteredValue")
		registerFunc(&_cFSocketCreate, frameworkHandle, "CFSocketCreate")
		registerFunc(&_cFSocketCreateConnectedToSocketSignature, frameworkHandle, "CFSocketCreateConnectedToSocketSignature")
		registerFunc(&_cFSocketCreateRunLoopSource, frameworkHandle, "CFSocketCreateRunLoopSource")
		registerFunc(&_cFSocketCreateWithNative, frameworkHandle, "CFSocketCreateWithNative")
		registerFunc(&_cFSocketCreateWithSocketSignature, frameworkHandle, "CFSocketCreateWithSocketSignature")
		registerFunc(&_cFSocketDisableCallBacks, frameworkHandle, "CFSocketDisableCallBacks")
		registerFunc(&_cFSocketEnableCallBacks, frameworkHandle, "CFSocketEnableCallBacks")
		registerFunc(&_cFSocketGetContext, frameworkHandle, "CFSocketGetContext")
		registerFunc(&_cFSocketGetDefaultNameRegistryPortNumber, frameworkHandle, "CFSocketGetDefaultNameRegistryPortNumber")
		registerFunc(&_cFSocketGetNative, frameworkHandle, "CFSocketGetNative")
		registerFunc(&_cFSocketGetSocketFlags, frameworkHandle, "CFSocketGetSocketFlags")
		registerFunc(&_cFSocketGetTypeID, frameworkHandle, "CFSocketGetTypeID")
		registerFunc(&_cFSocketInvalidate, frameworkHandle, "CFSocketInvalidate")
		registerFunc(&_cFSocketIsValid, frameworkHandle, "CFSocketIsValid")
		registerFunc(&_cFSocketRegisterSocketSignature, frameworkHandle, "CFSocketRegisterSocketSignature")
		registerFunc(&_cFSocketRegisterValue, frameworkHandle, "CFSocketRegisterValue")
		registerFunc(&_cFSocketSendData, frameworkHandle, "CFSocketSendData")
		registerFunc(&_cFSocketSetAddress, frameworkHandle, "CFSocketSetAddress")
		registerFunc(&_cFSocketSetDefaultNameRegistryPortNumber, frameworkHandle, "CFSocketSetDefaultNameRegistryPortNumber")
		registerFunc(&_cFSocketSetSocketFlags, frameworkHandle, "CFSocketSetSocketFlags")
		registerFunc(&_cFSocketUnregister, frameworkHandle, "CFSocketUnregister")
		registerFunc(&_cFStreamCreateBoundPair, frameworkHandle, "CFStreamCreateBoundPair")
		registerFunc(&_cFStreamCreatePairWithPeerSocketSignature, frameworkHandle, "CFStreamCreatePairWithPeerSocketSignature")
		registerFunc(&_cFStreamCreatePairWithSocket, frameworkHandle, "CFStreamCreatePairWithSocket")
		registerFunc(&_cFStreamCreatePairWithSocketToHost, frameworkHandle, "CFStreamCreatePairWithSocketToHost")
		registerFunc(&_cFStringAppend, frameworkHandle, "CFStringAppend")
		registerFunc(&_cFStringAppendCString, frameworkHandle, "CFStringAppendCString")
		registerFunc(&_cFStringAppendCharacters, frameworkHandle, "CFStringAppendCharacters")
		registerFunc(&_cFStringAppendFormat, frameworkHandle, "CFStringAppendFormat")
		registerFunc(&_cFStringAppendPascalString, frameworkHandle, "CFStringAppendPascalString")
		registerFunc(&_cFStringCapitalize, frameworkHandle, "CFStringCapitalize")
		registerFunc(&_cFStringCompare, frameworkHandle, "CFStringCompare")
		registerFunc(&_cFStringCompareWithOptions, frameworkHandle, "CFStringCompareWithOptions")
		registerFunc(&_cFStringCompareWithOptionsAndLocale, frameworkHandle, "CFStringCompareWithOptionsAndLocale")
		registerFunc(&_cFStringConvertEncodingToIANACharSetName, frameworkHandle, "CFStringConvertEncodingToIANACharSetName")
		registerFunc(&_cFStringConvertEncodingToNSStringEncoding, frameworkHandle, "CFStringConvertEncodingToNSStringEncoding")
		registerFunc(&_cFStringConvertEncodingToWindowsCodepage, frameworkHandle, "CFStringConvertEncodingToWindowsCodepage")
		registerFunc(&_cFStringConvertIANACharSetNameToEncoding, frameworkHandle, "CFStringConvertIANACharSetNameToEncoding")
		registerFunc(&_cFStringConvertNSStringEncodingToEncoding, frameworkHandle, "CFStringConvertNSStringEncodingToEncoding")
		registerFunc(&_cFStringConvertWindowsCodepageToEncoding, frameworkHandle, "CFStringConvertWindowsCodepageToEncoding")
		registerFunc(&_cFStringCreateArrayBySeparatingStrings, frameworkHandle, "CFStringCreateArrayBySeparatingStrings")
		registerFunc(&_cFStringCreateArrayWithFindResults, frameworkHandle, "CFStringCreateArrayWithFindResults")
		registerFunc(&_cFStringCreateByCombiningStrings, frameworkHandle, "CFStringCreateByCombiningStrings")
		registerFunc(&_cFStringCreateCopy, frameworkHandle, "CFStringCreateCopy")
		registerFunc(&_cFStringCreateExternalRepresentation, frameworkHandle, "CFStringCreateExternalRepresentation")
		registerFunc(&_cFStringCreateFromExternalRepresentation, frameworkHandle, "CFStringCreateFromExternalRepresentation")
		registerFunc(&_cFStringCreateMutable, frameworkHandle, "CFStringCreateMutable")
		registerFunc(&_cFStringCreateMutableCopy, frameworkHandle, "CFStringCreateMutableCopy")
		registerFunc(&_cFStringCreateMutableWithExternalCharactersNoCopy, frameworkHandle, "CFStringCreateMutableWithExternalCharactersNoCopy")
		registerFunc(&_cFStringCreateStringWithValidatedFormat, frameworkHandle, "CFStringCreateStringWithValidatedFormat")
		registerFunc(&_cFStringCreateWithBytes, frameworkHandle, "CFStringCreateWithBytes")
		registerFunc(&_cFStringCreateWithBytesNoCopy, frameworkHandle, "CFStringCreateWithBytesNoCopy")
		registerFunc(&_cFStringCreateWithCString, frameworkHandle, "CFStringCreateWithCString")
		registerFunc(&_cFStringCreateWithCStringNoCopy, frameworkHandle, "CFStringCreateWithCStringNoCopy")
		registerFunc(&_cFStringCreateWithCharacters, frameworkHandle, "CFStringCreateWithCharacters")
		registerFunc(&_cFStringCreateWithCharactersNoCopy, frameworkHandle, "CFStringCreateWithCharactersNoCopy")
		registerFunc(&_cFStringCreateWithFileSystemRepresentation, frameworkHandle, "CFStringCreateWithFileSystemRepresentation")
		registerFunc(&_cFStringCreateWithFormat, frameworkHandle, "CFStringCreateWithFormat")
		registerFunc(&_cFStringCreateWithPascalString, frameworkHandle, "CFStringCreateWithPascalString")
		registerFunc(&_cFStringCreateWithPascalStringNoCopy, frameworkHandle, "CFStringCreateWithPascalStringNoCopy")
		registerFunc(&_cFStringCreateWithSubstring, frameworkHandle, "CFStringCreateWithSubstring")
		registerFunc(&_cFStringDelete, frameworkHandle, "CFStringDelete")
		registerFunc(&_cFStringFind, frameworkHandle, "CFStringFind")
		registerFunc(&_cFStringFindAndReplace, frameworkHandle, "CFStringFindAndReplace")
		registerFunc(&_cFStringFindCharacterFromSet, frameworkHandle, "CFStringFindCharacterFromSet")
		registerFunc(&_cFStringFindWithOptions, frameworkHandle, "CFStringFindWithOptions")
		registerFunc(&_cFStringFindWithOptionsAndLocale, frameworkHandle, "CFStringFindWithOptionsAndLocale")
		registerFunc(&_cFStringFold, frameworkHandle, "CFStringFold")
		registerFunc(&_cFStringGetBytes, frameworkHandle, "CFStringGetBytes")
		registerFunc(&_cFStringGetCString, frameworkHandle, "CFStringGetCString")
		registerFunc(&_cFStringGetCStringPtr, frameworkHandle, "CFStringGetCStringPtr")
		registerFunc(&_cFStringGetCharacterAtIndex, frameworkHandle, "CFStringGetCharacterAtIndex")
		registerFunc(&_cFStringGetCharacters, frameworkHandle, "CFStringGetCharacters")
		registerFunc(&_cFStringGetCharactersPtr, frameworkHandle, "CFStringGetCharactersPtr")
		registerFunc(&_cFStringGetDoubleValue, frameworkHandle, "CFStringGetDoubleValue")
		registerFunc(&_cFStringGetFastestEncoding, frameworkHandle, "CFStringGetFastestEncoding")
		registerFunc(&_cFStringGetFileSystemRepresentation, frameworkHandle, "CFStringGetFileSystemRepresentation")
		registerFunc(&_cFStringGetHyphenationLocationBeforeIndex, frameworkHandle, "CFStringGetHyphenationLocationBeforeIndex")
		registerFunc(&_cFStringGetIntValue, frameworkHandle, "CFStringGetIntValue")
		registerFunc(&_cFStringGetLength, frameworkHandle, "CFStringGetLength")
		registerFunc(&_cFStringGetLineBounds, frameworkHandle, "CFStringGetLineBounds")
		registerFunc(&_cFStringGetListOfAvailableEncodings, frameworkHandle, "CFStringGetListOfAvailableEncodings")
		registerFunc(&_cFStringGetMaximumSizeForEncoding, frameworkHandle, "CFStringGetMaximumSizeForEncoding")
		registerFunc(&_cFStringGetMaximumSizeOfFileSystemRepresentation, frameworkHandle, "CFStringGetMaximumSizeOfFileSystemRepresentation")
		registerFunc(&_cFStringGetMostCompatibleMacStringEncoding, frameworkHandle, "CFStringGetMostCompatibleMacStringEncoding")
		registerFunc(&_cFStringGetNameOfEncoding, frameworkHandle, "CFStringGetNameOfEncoding")
		registerFunc(&_cFStringGetParagraphBounds, frameworkHandle, "CFStringGetParagraphBounds")
		registerFunc(&_cFStringGetPascalString, frameworkHandle, "CFStringGetPascalString")
		registerFunc(&_cFStringGetPascalStringPtr, frameworkHandle, "CFStringGetPascalStringPtr")
		registerFunc(&_cFStringGetRangeOfComposedCharactersAtIndex, frameworkHandle, "CFStringGetRangeOfComposedCharactersAtIndex")
		registerFunc(&_cFStringGetSmallestEncoding, frameworkHandle, "CFStringGetSmallestEncoding")
		registerFunc(&_cFStringGetSystemEncoding, frameworkHandle, "CFStringGetSystemEncoding")
		registerFunc(&_cFStringGetTypeID, frameworkHandle, "CFStringGetTypeID")
		registerFunc(&_cFStringHasPrefix, frameworkHandle, "CFStringHasPrefix")
		registerFunc(&_cFStringHasSuffix, frameworkHandle, "CFStringHasSuffix")
		registerFunc(&_cFStringInsert, frameworkHandle, "CFStringInsert")
		registerFunc(&_cFStringIsEncodingAvailable, frameworkHandle, "CFStringIsEncodingAvailable")
		registerFunc(&_cFStringIsHyphenationAvailableForLocale, frameworkHandle, "CFStringIsHyphenationAvailableForLocale")
		registerFunc(&_cFStringLowercase, frameworkHandle, "CFStringLowercase")
		registerFunc(&_cFStringNormalize, frameworkHandle, "CFStringNormalize")
		registerFunc(&_cFStringPad, frameworkHandle, "CFStringPad")
		registerFunc(&_cFStringReplace, frameworkHandle, "CFStringReplace")
		registerFunc(&_cFStringReplaceAll, frameworkHandle, "CFStringReplaceAll")
		registerFunc(&_cFStringSetExternalCharactersNoCopy, frameworkHandle, "CFStringSetExternalCharactersNoCopy")
		registerFunc(&_cFStringTokenizerAdvanceToNextToken, frameworkHandle, "CFStringTokenizerAdvanceToNextToken")
		registerFunc(&_cFStringTokenizerCopyBestStringLanguage, frameworkHandle, "CFStringTokenizerCopyBestStringLanguage")
		registerFunc(&_cFStringTokenizerCopyCurrentTokenAttribute, frameworkHandle, "CFStringTokenizerCopyCurrentTokenAttribute")
		registerFunc(&_cFStringTokenizerCreate, frameworkHandle, "CFStringTokenizerCreate")
		registerFunc(&_cFStringTokenizerGetCurrentSubTokens, frameworkHandle, "CFStringTokenizerGetCurrentSubTokens")
		registerFunc(&_cFStringTokenizerGetCurrentTokenRange, frameworkHandle, "CFStringTokenizerGetCurrentTokenRange")
		registerFunc(&_cFStringTokenizerGetTypeID, frameworkHandle, "CFStringTokenizerGetTypeID")
		registerFunc(&_cFStringTokenizerGoToTokenAtIndex, frameworkHandle, "CFStringTokenizerGoToTokenAtIndex")
		registerFunc(&_cFStringTokenizerSetString, frameworkHandle, "CFStringTokenizerSetString")
		registerFunc(&_cFStringTransform, frameworkHandle, "CFStringTransform")
		registerFunc(&_cFStringTrim, frameworkHandle, "CFStringTrim")
		registerFunc(&_cFStringTrimWhitespace, frameworkHandle, "CFStringTrimWhitespace")
		registerFunc(&_cFStringUppercase, frameworkHandle, "CFStringUppercase")
		registerFunc(&_cFTimeZoneCopyAbbreviation, frameworkHandle, "CFTimeZoneCopyAbbreviation")
		registerFunc(&_cFTimeZoneCopyAbbreviationDictionary, frameworkHandle, "CFTimeZoneCopyAbbreviationDictionary")
		registerFunc(&_cFTimeZoneCopyDefault, frameworkHandle, "CFTimeZoneCopyDefault")
		registerFunc(&_cFTimeZoneCopyKnownNames, frameworkHandle, "CFTimeZoneCopyKnownNames")
		registerFunc(&_cFTimeZoneCopyLocalizedName, frameworkHandle, "CFTimeZoneCopyLocalizedName")
		registerFunc(&_cFTimeZoneCopySystem, frameworkHandle, "CFTimeZoneCopySystem")
		registerFunc(&_cFTimeZoneCreate, frameworkHandle, "CFTimeZoneCreate")
		registerFunc(&_cFTimeZoneCreateWithName, frameworkHandle, "CFTimeZoneCreateWithName")
		registerFunc(&_cFTimeZoneCreateWithTimeIntervalFromGMT, frameworkHandle, "CFTimeZoneCreateWithTimeIntervalFromGMT")
		registerFunc(&_cFTimeZoneGetData, frameworkHandle, "CFTimeZoneGetData")
		registerFunc(&_cFTimeZoneGetDaylightSavingTimeOffset, frameworkHandle, "CFTimeZoneGetDaylightSavingTimeOffset")
		registerFunc(&_cFTimeZoneGetName, frameworkHandle, "CFTimeZoneGetName")
		registerFunc(&_cFTimeZoneGetNextDaylightSavingTimeTransition, frameworkHandle, "CFTimeZoneGetNextDaylightSavingTimeTransition")
		registerFunc(&_cFTimeZoneGetSecondsFromGMT, frameworkHandle, "CFTimeZoneGetSecondsFromGMT")
		registerFunc(&_cFTimeZoneGetTypeID, frameworkHandle, "CFTimeZoneGetTypeID")
		registerFunc(&_cFTimeZoneIsDaylightSavingTime, frameworkHandle, "CFTimeZoneIsDaylightSavingTime")
		registerFunc(&_cFTimeZoneResetSystem, frameworkHandle, "CFTimeZoneResetSystem")
		registerFunc(&_cFTimeZoneSetAbbreviationDictionary, frameworkHandle, "CFTimeZoneSetAbbreviationDictionary")
		registerFunc(&_cFTimeZoneSetDefault, frameworkHandle, "CFTimeZoneSetDefault")
		registerFunc(&_cFTreeAppendChild, frameworkHandle, "CFTreeAppendChild")
		registerFunc(&_cFTreeApplyFunctionToChildren, frameworkHandle, "CFTreeApplyFunctionToChildren")
		registerFunc(&_cFTreeCreate, frameworkHandle, "CFTreeCreate")
		registerFunc(&_cFTreeFindRoot, frameworkHandle, "CFTreeFindRoot")
		registerFunc(&_cFTreeGetChildAtIndex, frameworkHandle, "CFTreeGetChildAtIndex")
		registerFunc(&_cFTreeGetChildCount, frameworkHandle, "CFTreeGetChildCount")
		registerFunc(&_cFTreeGetChildren, frameworkHandle, "CFTreeGetChildren")
		registerFunc(&_cFTreeGetContext, frameworkHandle, "CFTreeGetContext")
		registerFunc(&_cFTreeGetFirstChild, frameworkHandle, "CFTreeGetFirstChild")
		registerFunc(&_cFTreeGetNextSibling, frameworkHandle, "CFTreeGetNextSibling")
		registerFunc(&_cFTreeGetParent, frameworkHandle, "CFTreeGetParent")
		registerFunc(&_cFTreeGetTypeID, frameworkHandle, "CFTreeGetTypeID")
		registerFunc(&_cFTreeInsertSibling, frameworkHandle, "CFTreeInsertSibling")
		registerFunc(&_cFTreePrependChild, frameworkHandle, "CFTreePrependChild")
		registerFunc(&_cFTreeRemove, frameworkHandle, "CFTreeRemove")
		registerFunc(&_cFTreeRemoveAllChildren, frameworkHandle, "CFTreeRemoveAllChildren")
		registerFunc(&_cFTreeSetContext, frameworkHandle, "CFTreeSetContext")
		registerFunc(&_cFTreeSortChildren, frameworkHandle, "CFTreeSortChildren")
		registerFunc(&_cFURLCanBeDecomposed, frameworkHandle, "CFURLCanBeDecomposed")
		registerFunc(&_cFURLClearResourcePropertyCache, frameworkHandle, "CFURLClearResourcePropertyCache")
		registerFunc(&_cFURLClearResourcePropertyCacheForKey, frameworkHandle, "CFURLClearResourcePropertyCacheForKey")
		registerFunc(&_cFURLCopyAbsoluteURL, frameworkHandle, "CFURLCopyAbsoluteURL")
		registerFunc(&_cFURLCopyFileSystemPath, frameworkHandle, "CFURLCopyFileSystemPath")
		registerFunc(&_cFURLCopyFragment, frameworkHandle, "CFURLCopyFragment")
		registerFunc(&_cFURLCopyHostName, frameworkHandle, "CFURLCopyHostName")
		registerFunc(&_cFURLCopyLastPathComponent, frameworkHandle, "CFURLCopyLastPathComponent")
		registerFunc(&_cFURLCopyNetLocation, frameworkHandle, "CFURLCopyNetLocation")
		registerFunc(&_cFURLCopyPassword, frameworkHandle, "CFURLCopyPassword")
		registerFunc(&_cFURLCopyPath, frameworkHandle, "CFURLCopyPath")
		registerFunc(&_cFURLCopyPathExtension, frameworkHandle, "CFURLCopyPathExtension")
		registerFunc(&_cFURLCopyQueryString, frameworkHandle, "CFURLCopyQueryString")
		registerFunc(&_cFURLCopyResourcePropertiesForKeys, frameworkHandle, "CFURLCopyResourcePropertiesForKeys")
		registerFunc(&_cFURLCopyResourcePropertyForKey, frameworkHandle, "CFURLCopyResourcePropertyForKey")
		registerFunc(&_cFURLCopyResourceSpecifier, frameworkHandle, "CFURLCopyResourceSpecifier")
		registerFunc(&_cFURLCopyScheme, frameworkHandle, "CFURLCopyScheme")
		registerFunc(&_cFURLCopyStrictPath, frameworkHandle, "CFURLCopyStrictPath")
		registerFunc(&_cFURLCopyUserName, frameworkHandle, "CFURLCopyUserName")
		registerFunc(&_cFURLCreateAbsoluteURLWithBytes, frameworkHandle, "CFURLCreateAbsoluteURLWithBytes")
		registerFunc(&_cFURLCreateBookmarkData, frameworkHandle, "CFURLCreateBookmarkData")
		registerFunc(&_cFURLCreateBookmarkDataFromFile, frameworkHandle, "CFURLCreateBookmarkDataFromFile")
		registerFunc(&_cFURLCreateByResolvingBookmarkData, frameworkHandle, "CFURLCreateByResolvingBookmarkData")
		registerFunc(&_cFURLCreateCopyAppendingPathComponent, frameworkHandle, "CFURLCreateCopyAppendingPathComponent")
		registerFunc(&_cFURLCreateCopyAppendingPathExtension, frameworkHandle, "CFURLCreateCopyAppendingPathExtension")
		registerFunc(&_cFURLCreateCopyDeletingLastPathComponent, frameworkHandle, "CFURLCreateCopyDeletingLastPathComponent")
		registerFunc(&_cFURLCreateCopyDeletingPathExtension, frameworkHandle, "CFURLCreateCopyDeletingPathExtension")
		registerFunc(&_cFURLCreateData, frameworkHandle, "CFURLCreateData")
		registerFunc(&_cFURLCreateFilePathURL, frameworkHandle, "CFURLCreateFilePathURL")
		registerFunc(&_cFURLCreateFileReferenceURL, frameworkHandle, "CFURLCreateFileReferenceURL")
		registerFunc(&_cFURLCreateFromFileSystemRepresentation, frameworkHandle, "CFURLCreateFromFileSystemRepresentation")
		registerFunc(&_cFURLCreateFromFileSystemRepresentationRelativeToBase, frameworkHandle, "CFURLCreateFromFileSystemRepresentationRelativeToBase")
		registerFunc(&_cFURLCreateResourcePropertiesForKeysFromBookmarkData, frameworkHandle, "CFURLCreateResourcePropertiesForKeysFromBookmarkData")
		registerFunc(&_cFURLCreateResourcePropertyForKeyFromBookmarkData, frameworkHandle, "CFURLCreateResourcePropertyForKeyFromBookmarkData")
		registerFunc(&_cFURLCreateStringByReplacingPercentEscapes, frameworkHandle, "CFURLCreateStringByReplacingPercentEscapes")
		registerFunc(&_cFURLCreateWithBytes, frameworkHandle, "CFURLCreateWithBytes")
		registerFunc(&_cFURLCreateWithFileSystemPath, frameworkHandle, "CFURLCreateWithFileSystemPath")
		registerFunc(&_cFURLCreateWithFileSystemPathRelativeToBase, frameworkHandle, "CFURLCreateWithFileSystemPathRelativeToBase")
		registerFunc(&_cFURLCreateWithString, frameworkHandle, "CFURLCreateWithString")
		registerFunc(&_cFURLEnumeratorCreateForDirectoryURL, frameworkHandle, "CFURLEnumeratorCreateForDirectoryURL")
		registerFunc(&_cFURLEnumeratorCreateForMountedVolumes, frameworkHandle, "CFURLEnumeratorCreateForMountedVolumes")
		registerFunc(&_cFURLEnumeratorGetDescendentLevel, frameworkHandle, "CFURLEnumeratorGetDescendentLevel")
		registerFunc(&_cFURLEnumeratorGetNextURL, frameworkHandle, "CFURLEnumeratorGetNextURL")
		registerFunc(&_cFURLEnumeratorGetTypeID, frameworkHandle, "CFURLEnumeratorGetTypeID")
		registerFunc(&_cFURLEnumeratorSkipDescendents, frameworkHandle, "CFURLEnumeratorSkipDescendents")
		registerFunc(&_cFURLGetBaseURL, frameworkHandle, "CFURLGetBaseURL")
		registerFunc(&_cFURLGetByteRangeForComponent, frameworkHandle, "CFURLGetByteRangeForComponent")
		registerFunc(&_cFURLGetBytes, frameworkHandle, "CFURLGetBytes")
		registerFunc(&_cFURLGetFileSystemRepresentation, frameworkHandle, "CFURLGetFileSystemRepresentation")
		registerFunc(&_cFURLGetPortNumber, frameworkHandle, "CFURLGetPortNumber")
		registerFunc(&_cFURLGetString, frameworkHandle, "CFURLGetString")
		registerFunc(&_cFURLGetTypeID, frameworkHandle, "CFURLGetTypeID")
		registerFunc(&_cFURLHasDirectoryPath, frameworkHandle, "CFURLHasDirectoryPath")
		registerFunc(&_cFURLIsFileReferenceURL, frameworkHandle, "CFURLIsFileReferenceURL")
		registerFunc(&_cFURLResourceIsReachable, frameworkHandle, "CFURLResourceIsReachable")
		registerFunc(&_cFURLSetResourcePropertiesForKeys, frameworkHandle, "CFURLSetResourcePropertiesForKeys")
		registerFunc(&_cFURLSetResourcePropertyForKey, frameworkHandle, "CFURLSetResourcePropertyForKey")
		registerFunc(&_cFURLSetTemporaryResourcePropertyForKey, frameworkHandle, "CFURLSetTemporaryResourcePropertyForKey")
		registerFunc(&_cFURLStartAccessingSecurityScopedResource, frameworkHandle, "CFURLStartAccessingSecurityScopedResource")
		registerFunc(&_cFURLStopAccessingSecurityScopedResource, frameworkHandle, "CFURLStopAccessingSecurityScopedResource")
		registerFunc(&_cFURLWriteBookmarkDataToFile, frameworkHandle, "CFURLWriteBookmarkDataToFile")
		registerFunc(&_cFUUIDCreate, frameworkHandle, "CFUUIDCreate")
		registerFunc(&_cFUUIDCreateFromString, frameworkHandle, "CFUUIDCreateFromString")
		registerFunc(&_cFUUIDCreateFromUUIDBytes, frameworkHandle, "CFUUIDCreateFromUUIDBytes")
		registerFunc(&_cFUUIDCreateString, frameworkHandle, "CFUUIDCreateString")
		registerFunc(&_cFUUIDCreateWithBytes, frameworkHandle, "CFUUIDCreateWithBytes")
		registerFunc(&_cFUUIDGetConstantUUIDWithBytes, frameworkHandle, "CFUUIDGetConstantUUIDWithBytes")
		registerFunc(&_cFUUIDGetTypeID, frameworkHandle, "CFUUIDGetTypeID")
		registerFunc(&_cFUUIDGetUUIDBytes, frameworkHandle, "CFUUIDGetUUIDBytes")
		registerFunc(&_cFUserNotificationCancel, frameworkHandle, "CFUserNotificationCancel")
		registerFunc(&_cFUserNotificationCreate, frameworkHandle, "CFUserNotificationCreate")
		registerFunc(&_cFUserNotificationCreateRunLoopSource, frameworkHandle, "CFUserNotificationCreateRunLoopSource")
		registerFunc(&_cFUserNotificationDisplayAlert, frameworkHandle, "CFUserNotificationDisplayAlert")
		registerFunc(&_cFUserNotificationDisplayNotice, frameworkHandle, "CFUserNotificationDisplayNotice")
		registerFunc(&_cFUserNotificationGetResponseDictionary, frameworkHandle, "CFUserNotificationGetResponseDictionary")
		registerFunc(&_cFUserNotificationGetResponseValue, frameworkHandle, "CFUserNotificationGetResponseValue")
		registerFunc(&_cFUserNotificationGetTypeID, frameworkHandle, "CFUserNotificationGetTypeID")
		registerFunc(&_cFUserNotificationReceiveResponse, frameworkHandle, "CFUserNotificationReceiveResponse")
		registerFunc(&_cFUserNotificationUpdate, frameworkHandle, "CFUserNotificationUpdate")
		registerFunc(&_cFWriteStreamCanAcceptBytes, frameworkHandle, "CFWriteStreamCanAcceptBytes")
		registerFunc(&_cFWriteStreamClose, frameworkHandle, "CFWriteStreamClose")
		registerFunc(&_cFWriteStreamCopyDispatchQueue, frameworkHandle, "CFWriteStreamCopyDispatchQueue")
		registerFunc(&_cFWriteStreamCopyError, frameworkHandle, "CFWriteStreamCopyError")
		registerFunc(&_cFWriteStreamCopyProperty, frameworkHandle, "CFWriteStreamCopyProperty")
		registerFunc(&_cFWriteStreamCreateWithAllocatedBuffers, frameworkHandle, "CFWriteStreamCreateWithAllocatedBuffers")
		registerFunc(&_cFWriteStreamCreateWithBuffer, frameworkHandle, "CFWriteStreamCreateWithBuffer")
		registerFunc(&_cFWriteStreamCreateWithFile, frameworkHandle, "CFWriteStreamCreateWithFile")
		registerFunc(&_cFWriteStreamGetError, frameworkHandle, "CFWriteStreamGetError")
		registerFunc(&_cFWriteStreamGetStatus, frameworkHandle, "CFWriteStreamGetStatus")
		registerFunc(&_cFWriteStreamGetTypeID, frameworkHandle, "CFWriteStreamGetTypeID")
		registerFunc(&_cFWriteStreamOpen, frameworkHandle, "CFWriteStreamOpen")
		registerFunc(&_cFWriteStreamScheduleWithRunLoop, frameworkHandle, "CFWriteStreamScheduleWithRunLoop")
		registerFunc(&_cFWriteStreamSetClient, frameworkHandle, "CFWriteStreamSetClient")
		registerFunc(&_cFWriteStreamSetDispatchQueue, frameworkHandle, "CFWriteStreamSetDispatchQueue")
		registerFunc(&_cFWriteStreamSetProperty, frameworkHandle, "CFWriteStreamSetProperty")
		registerFunc(&_cFWriteStreamUnscheduleFromRunLoop, frameworkHandle, "CFWriteStreamUnscheduleFromRunLoop")
		registerFunc(&_cFWriteStreamWrite, frameworkHandle, "CFWriteStreamWrite")
		registerFunc(&_cFXMLCreateStringByEscapingEntities, frameworkHandle, "CFXMLCreateStringByEscapingEntities")
		registerFunc(&_cFXMLCreateStringByUnescapingEntities, frameworkHandle, "CFXMLCreateStringByUnescapingEntities")
		registerFunc(&_cFXMLNodeCreate, frameworkHandle, "CFXMLNodeCreate")
		registerFunc(&_cFXMLNodeCreateCopy, frameworkHandle, "CFXMLNodeCreateCopy")
		registerFunc(&_cFXMLNodeGetInfoPtr, frameworkHandle, "CFXMLNodeGetInfoPtr")
		registerFunc(&_cFXMLNodeGetString, frameworkHandle, "CFXMLNodeGetString")
		registerFunc(&_cFXMLNodeGetTypeCode, frameworkHandle, "CFXMLNodeGetTypeCode")
		registerFunc(&_cFXMLNodeGetTypeID, frameworkHandle, "CFXMLNodeGetTypeID")
		registerFunc(&_cFXMLNodeGetVersion, frameworkHandle, "CFXMLNodeGetVersion")
		registerFunc(&_cFXMLParserAbort, frameworkHandle, "CFXMLParserAbort")
		registerFunc(&_cFXMLParserCopyErrorDescription, frameworkHandle, "CFXMLParserCopyErrorDescription")
		registerFunc(&_cFXMLParserCreate, frameworkHandle, "CFXMLParserCreate")
		registerFunc(&_cFXMLParserCreateWithDataFromURL, frameworkHandle, "CFXMLParserCreateWithDataFromURL")
		registerFunc(&_cFXMLParserGetCallBacks, frameworkHandle, "CFXMLParserGetCallBacks")
		registerFunc(&_cFXMLParserGetContext, frameworkHandle, "CFXMLParserGetContext")
		registerFunc(&_cFXMLParserGetDocument, frameworkHandle, "CFXMLParserGetDocument")
		registerFunc(&_cFXMLParserGetLineNumber, frameworkHandle, "CFXMLParserGetLineNumber")
		registerFunc(&_cFXMLParserGetLocation, frameworkHandle, "CFXMLParserGetLocation")
		registerFunc(&_cFXMLParserGetSourceURL, frameworkHandle, "CFXMLParserGetSourceURL")
		registerFunc(&_cFXMLParserGetStatusCode, frameworkHandle, "CFXMLParserGetStatusCode")
		registerFunc(&_cFXMLParserGetTypeID, frameworkHandle, "CFXMLParserGetTypeID")
		registerFunc(&_cFXMLParserParse, frameworkHandle, "CFXMLParserParse")
		registerFunc(&_cFXMLTreeCreateFromData, frameworkHandle, "CFXMLTreeCreateFromData")
		registerFunc(&_cFXMLTreeCreateFromDataWithError, frameworkHandle, "CFXMLTreeCreateFromDataWithError")
		registerFunc(&_cFXMLTreeCreateWithDataFromURL, frameworkHandle, "CFXMLTreeCreateWithDataFromURL")
		registerFunc(&_cFXMLTreeCreateWithNode, frameworkHandle, "CFXMLTreeCreateWithNode")
		registerFunc(&_cFXMLTreeCreateXMLData, frameworkHandle, "CFXMLTreeCreateXMLData")
		registerFunc(&_cFXMLTreeGetNode, frameworkHandle, "CFXMLTreeGetNode")
	}


