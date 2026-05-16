// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreservices"
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
		return fmt.Sprintf("ApplicationServices: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("ApplicationServices: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("ApplicationServices: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("ApplicationServices: register symbol %s: %v", name, r)
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

var _aXAPIEnabled func() bool
var _aXAPIEnabledErr error

func tryAXAPIEnabled() (bool, error) {
	if _aXAPIEnabled == nil {
		return false, symbolCallError("AXAPIEnabled", "10.0", _aXAPIEnabledErr)
	}
	return _aXAPIEnabled(), nil
}

// AXAPIEnabled returns whether the accessibility API is enabled.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1462072-axapienabled
func AXAPIEnabled() bool {
	result, callErr := tryAXAPIEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXIsProcessTrusted func() bool
var _aXIsProcessTrustedErr error

func tryAXIsProcessTrusted() (bool, error) {
	if _aXIsProcessTrusted == nil {
		return false, symbolCallError("AXIsProcessTrusted", "10.4", _aXIsProcessTrustedErr)
	}
	return _aXIsProcessTrusted(), nil
}

// AXIsProcessTrusted returns whether the current process is a trusted accessibility client.
//
// See: https://developer.apple.com/documentation/applicationservices/1460720-axisprocesstrusted
func AXIsProcessTrusted() bool {
	result, callErr := tryAXIsProcessTrusted()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXIsProcessTrustedWithOptions func(options corefoundation.CFDictionaryRef) bool
var _aXIsProcessTrustedWithOptionsErr error

func tryAXIsProcessTrustedWithOptions(options corefoundation.CFDictionaryRef) (bool, error) {
	if _aXIsProcessTrustedWithOptions == nil {
		return false, symbolCallError("AXIsProcessTrustedWithOptions", "10.9", _aXIsProcessTrustedWithOptionsErr)
	}
	return _aXIsProcessTrustedWithOptions(options), nil
}

// AXIsProcessTrustedWithOptions returns whether the current process is a trusted accessibility client.
//
// See: https://developer.apple.com/documentation/applicationservices/1459186-axisprocesstrustedwithoptions
func AXIsProcessTrustedWithOptions(options corefoundation.CFDictionaryRef) bool {
	result, callErr := tryAXIsProcessTrustedWithOptions(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXMakeProcessTrusted func(arg0 corefoundation.CFStringRef) AXError
var _aXMakeProcessTrustedErr error

func tryAXMakeProcessTrusted(arg0 corefoundation.CFStringRef) (AXError, error) {
	if _aXMakeProcessTrusted == nil {
		return *new(AXError), symbolCallError("AXMakeProcessTrusted", "10.4", _aXMakeProcessTrustedErr)
	}
	return _aXMakeProcessTrusted(arg0), nil
}

// AXMakeProcessTrusted attempts to make the process represented by the specified path a trusted accessibility client.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1462083-axmakeprocesstrusted
func AXMakeProcessTrusted(arg0 corefoundation.CFStringRef) AXError {
	result, callErr := tryAXMakeProcessTrusted(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXObserverAddNotification func(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef, refcon uintptr) AXError
var _aXObserverAddNotificationErr error

func tryAXObserverAddNotification(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef, refcon uintptr) (AXError, error) {
	if _aXObserverAddNotification == nil {
		return *new(AXError), symbolCallError("AXObserverAddNotification", "10.2", _aXObserverAddNotificationErr)
	}
	return _aXObserverAddNotification(observer, element, notification, refcon), nil
}

// AXObserverAddNotification registers the specified observer to receive notifications from the specified accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462089-axobserveraddnotification
func AXObserverAddNotification(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef, refcon uintptr) AXError {
	result, callErr := tryAXObserverAddNotification(observer, element, notification, refcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXObserverCreate func(application int32, callback unsafe.Pointer, outObserver *AXObserverRef) AXError
var _aXObserverCreateErr error

func tryAXObserverCreate(application int32, callback AXObserverCallback, outObserver *AXObserverRef) (AXError, error) {
	if _aXObserverCreate == nil {
		return *new(AXError), symbolCallError("AXObserverCreate", "10.2", _aXObserverCreateErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 AXObserverRef, blockArg1 AXUIElementRef, blockArg2 corefoundation.CFStringRef, blockArg3 unsafe.Pointer) {
		callback(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _aXObserverCreate(application, _block0, outObserver), nil
}

// AXObserverCreate creates a new observer that can receive notifications from the specified application.
//
// See: https://developer.apple.com/documentation/applicationservices/1460133-axobservercreate
func AXObserverCreate(application int32, callback AXObserverCallback, outObserver *AXObserverRef) AXError {
	result, callErr := tryAXObserverCreate(application, callback, outObserver)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXObserverCreateWithInfoCallback func(application int32, callback unsafe.Pointer, outObserver *AXObserverRef) AXError
var _aXObserverCreateWithInfoCallbackErr error

func tryAXObserverCreateWithInfoCallback(application int32, callback AXObserverCallbackWithInfo, outObserver *AXObserverRef) (AXError, error) {
	if _aXObserverCreateWithInfoCallback == nil {
		return *new(AXError), symbolCallError("AXObserverCreateWithInfoCallback", "10.9", _aXObserverCreateWithInfoCallbackErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 AXObserverRef, blockArg1 AXUIElementRef, blockArg2 corefoundation.CFStringRef, blockArg3 corefoundation.CFDictionaryRef, blockArg4 unsafe.Pointer) {
		callback(blockArg0, blockArg1, blockArg2, blockArg3, blockArg4)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _aXObserverCreateWithInfoCallback(application, _block0, outObserver), nil
}

// AXObserverCreateWithInfoCallback creates a new observer that can receive notifications with an information dictionary from the specified application.
//
// See: https://developer.apple.com/documentation/applicationservices/1460610-axobservercreatewithinfocallback
func AXObserverCreateWithInfoCallback(application int32, callback AXObserverCallbackWithInfo, outObserver *AXObserverRef) AXError {
	result, callErr := tryAXObserverCreateWithInfoCallback(application, callback, outObserver)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXObserverGetRunLoopSource func(observer AXObserverRef) corefoundation.CFRunLoopSourceRef
var _aXObserverGetRunLoopSourceErr error

func tryAXObserverGetRunLoopSource(observer AXObserverRef) (corefoundation.CFRunLoopSourceRef, error) {
	if _aXObserverGetRunLoopSource == nil {
		return 0, symbolCallError("AXObserverGetRunLoopSource", "10.2", _aXObserverGetRunLoopSourceErr)
	}
	return _aXObserverGetRunLoopSource(observer), nil
}

// AXObserverGetRunLoopSource returns the observer's run loop source.
//
// See: https://developer.apple.com/documentation/applicationservices/1459139-axobservergetrunloopsource
func AXObserverGetRunLoopSource(observer AXObserverRef) corefoundation.CFRunLoopSourceRef {
	result, callErr := tryAXObserverGetRunLoopSource(observer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXObserverGetTypeID func() uint
var _aXObserverGetTypeIDErr error

func tryAXObserverGetTypeID() (uint, error) {
	if _aXObserverGetTypeID == nil {
		return 0, symbolCallError("AXObserverGetTypeID", "10.2", _aXObserverGetTypeIDErr)
	}
	return _aXObserverGetTypeID(), nil
}

// AXObserverGetTypeID returns the unique type identifier for the AXObserverRef type.
//
// See: https://developer.apple.com/documentation/applicationservices/1461244-axobservergettypeid
func AXObserverGetTypeID() uint {
	result, callErr := tryAXObserverGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXObserverRemoveNotification func(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef) AXError
var _aXObserverRemoveNotificationErr error

func tryAXObserverRemoveNotification(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef) (AXError, error) {
	if _aXObserverRemoveNotification == nil {
		return *new(AXError), symbolCallError("AXObserverRemoveNotification", "10.2", _aXObserverRemoveNotificationErr)
	}
	return _aXObserverRemoveNotification(observer, element, notification), nil
}

// AXObserverRemoveNotification removes the specified notification from the list of notifications the observer wants to receive from the accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462066-axobserverremovenotification
func AXObserverRemoveNotification(observer AXObserverRef, element AXUIElementRef, notification corefoundation.CFStringRef) AXError {
	result, callErr := tryAXObserverRemoveNotification(observer, element, notification)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerCreate func(allocator corefoundation.CFAllocatorRef, bytes *byte, length int) AXTextMarkerRef
var _aXTextMarkerCreateErr error

func tryAXTextMarkerCreate(allocator corefoundation.CFAllocatorRef, bytes []byte, length int) (AXTextMarkerRef, error) {
	if _aXTextMarkerCreate == nil {
		return 0, symbolCallError("AXTextMarkerCreate", "12.0", _aXTextMarkerCreateErr)
	}
	return _aXTextMarkerCreate(allocator, unsafe.SliceData(bytes), length), nil
}

// AXTextMarkerCreate.
//
// See: https://developer.apple.com/documentation/applicationservices/3882823-axtextmarkercreate
func AXTextMarkerCreate(allocator corefoundation.CFAllocatorRef, bytes []byte, length int) AXTextMarkerRef {
	result, callErr := tryAXTextMarkerCreate(allocator, bytes, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerGetBytePtr func(theTextMarker AXTextMarkerRef) uint8
var _aXTextMarkerGetBytePtrErr error

func tryAXTextMarkerGetBytePtr(theTextMarker AXTextMarkerRef) (uint8, error) {
	if _aXTextMarkerGetBytePtr == nil {
		return 0, symbolCallError("AXTextMarkerGetBytePtr", "12.0", _aXTextMarkerGetBytePtrErr)
	}
	return _aXTextMarkerGetBytePtr(theTextMarker), nil
}

// AXTextMarkerGetBytePtr.
//
// See: https://developer.apple.com/documentation/applicationservices/3882824-axtextmarkergetbyteptr
func AXTextMarkerGetBytePtr(theTextMarker AXTextMarkerRef) uint8 {
	result, callErr := tryAXTextMarkerGetBytePtr(theTextMarker)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerGetLength func(marker AXTextMarkerRef) int
var _aXTextMarkerGetLengthErr error

func tryAXTextMarkerGetLength(marker AXTextMarkerRef) (int, error) {
	if _aXTextMarkerGetLength == nil {
		return 0, symbolCallError("AXTextMarkerGetLength", "12.0", _aXTextMarkerGetLengthErr)
	}
	return _aXTextMarkerGetLength(marker), nil
}

// AXTextMarkerGetLength.
//
// See: https://developer.apple.com/documentation/applicationservices/3882825-axtextmarkergetlength
func AXTextMarkerGetLength(marker AXTextMarkerRef) int {
	result, callErr := tryAXTextMarkerGetLength(marker)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerGetTypeID func() uint
var _aXTextMarkerGetTypeIDErr error

func tryAXTextMarkerGetTypeID() (uint, error) {
	if _aXTextMarkerGetTypeID == nil {
		return 0, symbolCallError("AXTextMarkerGetTypeID", "12.0", _aXTextMarkerGetTypeIDErr)
	}
	return _aXTextMarkerGetTypeID(), nil
}

// AXTextMarkerGetTypeID.
//
// See: https://developer.apple.com/documentation/applicationservices/3882826-axtextmarkergettypeid
func AXTextMarkerGetTypeID() uint {
	result, callErr := tryAXTextMarkerGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerRangeCopyEndMarker func(textMarkerRange AXTextMarkerRangeRef) AXTextMarkerRef
var _aXTextMarkerRangeCopyEndMarkerErr error

func tryAXTextMarkerRangeCopyEndMarker(textMarkerRange AXTextMarkerRangeRef) (AXTextMarkerRef, error) {
	if _aXTextMarkerRangeCopyEndMarker == nil {
		return 0, symbolCallError("AXTextMarkerRangeCopyEndMarker", "12.0", _aXTextMarkerRangeCopyEndMarkerErr)
	}
	return _aXTextMarkerRangeCopyEndMarker(textMarkerRange), nil
}

// AXTextMarkerRangeCopyEndMarker.
//
// See: https://developer.apple.com/documentation/applicationservices/3882827-axtextmarkerrangecopyendmarker
func AXTextMarkerRangeCopyEndMarker(textMarkerRange AXTextMarkerRangeRef) AXTextMarkerRef {
	result, callErr := tryAXTextMarkerRangeCopyEndMarker(textMarkerRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerRangeCopyStartMarker func(textMarkerRange AXTextMarkerRangeRef) AXTextMarkerRef
var _aXTextMarkerRangeCopyStartMarkerErr error

func tryAXTextMarkerRangeCopyStartMarker(textMarkerRange AXTextMarkerRangeRef) (AXTextMarkerRef, error) {
	if _aXTextMarkerRangeCopyStartMarker == nil {
		return 0, symbolCallError("AXTextMarkerRangeCopyStartMarker", "12.0", _aXTextMarkerRangeCopyStartMarkerErr)
	}
	return _aXTextMarkerRangeCopyStartMarker(textMarkerRange), nil
}

// AXTextMarkerRangeCopyStartMarker.
//
// See: https://developer.apple.com/documentation/applicationservices/3882828-axtextmarkerrangecopystartmarker
func AXTextMarkerRangeCopyStartMarker(textMarkerRange AXTextMarkerRangeRef) AXTextMarkerRef {
	result, callErr := tryAXTextMarkerRangeCopyStartMarker(textMarkerRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerRangeCreate func(allocator corefoundation.CFAllocatorRef, startMarker AXTextMarkerRef, endMarker AXTextMarkerRef) AXTextMarkerRangeRef
var _aXTextMarkerRangeCreateErr error

func tryAXTextMarkerRangeCreate(allocator corefoundation.CFAllocatorRef, startMarker AXTextMarkerRef, endMarker AXTextMarkerRef) (AXTextMarkerRangeRef, error) {
	if _aXTextMarkerRangeCreate == nil {
		return 0, symbolCallError("AXTextMarkerRangeCreate", "12.0", _aXTextMarkerRangeCreateErr)
	}
	return _aXTextMarkerRangeCreate(allocator, startMarker, endMarker), nil
}

// AXTextMarkerRangeCreate.
//
// See: https://developer.apple.com/documentation/applicationservices/3882829-axtextmarkerrangecreate
func AXTextMarkerRangeCreate(allocator corefoundation.CFAllocatorRef, startMarker AXTextMarkerRef, endMarker AXTextMarkerRef) AXTextMarkerRangeRef {
	result, callErr := tryAXTextMarkerRangeCreate(allocator, startMarker, endMarker)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerRangeCreateWithBytes func(allocator corefoundation.CFAllocatorRef, startMarkerBytes *byte, startMarkerLength int, endMarkerBytes *byte, endMarkerLength int) AXTextMarkerRangeRef
var _aXTextMarkerRangeCreateWithBytesErr error

func tryAXTextMarkerRangeCreateWithBytes(allocator corefoundation.CFAllocatorRef, startMarkerBytes []byte, startMarkerLength int, endMarkerBytes []byte, endMarkerLength int) (AXTextMarkerRangeRef, error) {
	if _aXTextMarkerRangeCreateWithBytes == nil {
		return 0, symbolCallError("AXTextMarkerRangeCreateWithBytes", "12.0", _aXTextMarkerRangeCreateWithBytesErr)
	}
	return _aXTextMarkerRangeCreateWithBytes(allocator, unsafe.SliceData(startMarkerBytes), startMarkerLength, unsafe.SliceData(endMarkerBytes), endMarkerLength), nil
}

// AXTextMarkerRangeCreateWithBytes.
//
// See: https://developer.apple.com/documentation/applicationservices/3882830-axtextmarkerrangecreatewithbytes
func AXTextMarkerRangeCreateWithBytes(allocator corefoundation.CFAllocatorRef, startMarkerBytes []byte, startMarkerLength int, endMarkerBytes []byte, endMarkerLength int) AXTextMarkerRangeRef {
	result, callErr := tryAXTextMarkerRangeCreateWithBytes(allocator, startMarkerBytes, startMarkerLength, endMarkerBytes, endMarkerLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXTextMarkerRangeGetTypeID func() uint
var _aXTextMarkerRangeGetTypeIDErr error

func tryAXTextMarkerRangeGetTypeID() (uint, error) {
	if _aXTextMarkerRangeGetTypeID == nil {
		return 0, symbolCallError("AXTextMarkerRangeGetTypeID", "12.0", _aXTextMarkerRangeGetTypeIDErr)
	}
	return _aXTextMarkerRangeGetTypeID(), nil
}

// AXTextMarkerRangeGetTypeID.
//
// See: https://developer.apple.com/documentation/applicationservices/3882831-axtextmarkerrangegettypeid
func AXTextMarkerRangeGetTypeID() uint {
	result, callErr := tryAXTextMarkerRangeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyActionDescription func(element AXUIElementRef, action corefoundation.CFStringRef, description *corefoundation.CFStringRef) AXError
var _aXUIElementCopyActionDescriptionErr error

func tryAXUIElementCopyActionDescription(element AXUIElementRef, action corefoundation.CFStringRef, description *corefoundation.CFStringRef) (AXError, error) {
	if _aXUIElementCopyActionDescription == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyActionDescription", "10.2", _aXUIElementCopyActionDescriptionErr)
	}
	return _aXUIElementCopyActionDescription(element, action, description), nil
}

// AXUIElementCopyActionDescription returns a localized description of the specified accessibility object's action.
//
// See: https://developer.apple.com/documentation/applicationservices/1462075-axuielementcopyactiondescription
func AXUIElementCopyActionDescription(element AXUIElementRef, action corefoundation.CFStringRef, description *corefoundation.CFStringRef) AXError {
	result, callErr := tryAXUIElementCopyActionDescription(element, action, description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyActionNames func(element AXUIElementRef, names *corefoundation.CFArrayRef) AXError
var _aXUIElementCopyActionNamesErr error

func tryAXUIElementCopyActionNames(element AXUIElementRef, names *corefoundation.CFArrayRef) (AXError, error) {
	if _aXUIElementCopyActionNames == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyActionNames", "10.2", _aXUIElementCopyActionNamesErr)
	}
	return _aXUIElementCopyActionNames(element, names), nil
}

// AXUIElementCopyActionNames returns a list of all the actions the specified accessibility object can perform.
//
// See: https://developer.apple.com/documentation/applicationservices/1462053-axuielementcopyactionnames
func AXUIElementCopyActionNames(element AXUIElementRef, names *corefoundation.CFArrayRef) AXError {
	result, callErr := tryAXUIElementCopyActionNames(element, names)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyAttributeNames func(element AXUIElementRef, names *corefoundation.CFArrayRef) AXError
var _aXUIElementCopyAttributeNamesErr error

func tryAXUIElementCopyAttributeNames(element AXUIElementRef, names *corefoundation.CFArrayRef) (AXError, error) {
	if _aXUIElementCopyAttributeNames == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyAttributeNames", "10.2", _aXUIElementCopyAttributeNamesErr)
	}
	return _aXUIElementCopyAttributeNames(element, names), nil
}

// AXUIElementCopyAttributeNames returns a list of all the attributes supported by the specified accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/1459475-axuielementcopyattributenames
func AXUIElementCopyAttributeNames(element AXUIElementRef, names *corefoundation.CFArrayRef) AXError {
	result, callErr := tryAXUIElementCopyAttributeNames(element, names)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyAttributeValue func(element AXUIElementRef, attribute corefoundation.CFStringRef, value *corefoundation.CFTypeRef) AXError
var _aXUIElementCopyAttributeValueErr error

func tryAXUIElementCopyAttributeValue(element AXUIElementRef, attribute corefoundation.CFStringRef, value *corefoundation.CFTypeRef) (AXError, error) {
	if _aXUIElementCopyAttributeValue == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyAttributeValue", "10.2", _aXUIElementCopyAttributeValueErr)
	}
	return _aXUIElementCopyAttributeValue(element, attribute, value), nil
}

// AXUIElementCopyAttributeValue returns the value of an accessibility object's attribute.
//
// See: https://developer.apple.com/documentation/applicationservices/1462085-axuielementcopyattributevalue
func AXUIElementCopyAttributeValue(element AXUIElementRef, attribute corefoundation.CFStringRef, value *corefoundation.CFTypeRef) AXError {
	result, callErr := tryAXUIElementCopyAttributeValue(element, attribute, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyAttributeValues func(element AXUIElementRef, attribute corefoundation.CFStringRef, index int, maxValues int, values *corefoundation.CFArrayRef) AXError
var _aXUIElementCopyAttributeValuesErr error

func tryAXUIElementCopyAttributeValues(element AXUIElementRef, attribute corefoundation.CFStringRef, index int, maxValues int, values *corefoundation.CFArrayRef) (AXError, error) {
	if _aXUIElementCopyAttributeValues == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyAttributeValues", "10.2", _aXUIElementCopyAttributeValuesErr)
	}
	return _aXUIElementCopyAttributeValues(element, attribute, index, maxValues, values), nil
}

// AXUIElementCopyAttributeValues returns an array of attribute values for the accessibility object's attribute, starting at the specified index.
//
// See: https://developer.apple.com/documentation/applicationservices/1462060-axuielementcopyattributevalues
func AXUIElementCopyAttributeValues(element AXUIElementRef, attribute corefoundation.CFStringRef, index int, maxValues int, values *corefoundation.CFArrayRef) AXError {
	result, callErr := tryAXUIElementCopyAttributeValues(element, attribute, index, maxValues, values)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyElementAtPosition func(application AXUIElementRef, x unsafe.Pointer, y unsafe.Pointer, element *AXUIElementRef) AXError
var _aXUIElementCopyElementAtPositionErr error

func tryAXUIElementCopyElementAtPosition(application AXUIElementRef, x unsafe.Pointer, y unsafe.Pointer, element *AXUIElementRef) (AXError, error) {
	if _aXUIElementCopyElementAtPosition == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyElementAtPosition", "10.2", _aXUIElementCopyElementAtPositionErr)
	}
	return _aXUIElementCopyElementAtPosition(application, x, y, element), nil
}

// AXUIElementCopyElementAtPosition returns the accessibility object at the specified position in top-left relative screen coordinates.
//
// See: https://developer.apple.com/documentation/applicationservices/1462077-axuielementcopyelementatposition
func AXUIElementCopyElementAtPosition(application AXUIElementRef, x unsafe.Pointer, y unsafe.Pointer, element *AXUIElementRef) AXError {
	result, callErr := tryAXUIElementCopyElementAtPosition(application, x, y, element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyMultipleAttributeValues func(element AXUIElementRef, attributes corefoundation.CFArrayRef, options AXCopyMultipleAttributeOptions, values *corefoundation.CFArrayRef) AXError
var _aXUIElementCopyMultipleAttributeValuesErr error

func tryAXUIElementCopyMultipleAttributeValues(element AXUIElementRef, attributes corefoundation.CFArrayRef, options AXCopyMultipleAttributeOptions, values *corefoundation.CFArrayRef) (AXError, error) {
	if _aXUIElementCopyMultipleAttributeValues == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyMultipleAttributeValues", "10.4", _aXUIElementCopyMultipleAttributeValuesErr)
	}
	return _aXUIElementCopyMultipleAttributeValues(element, attributes, options, values), nil
}

// AXUIElementCopyMultipleAttributeValues returns the values of multiple attributes in the accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462051-axuielementcopymultipleattribute
func AXUIElementCopyMultipleAttributeValues(element AXUIElementRef, attributes corefoundation.CFArrayRef, options AXCopyMultipleAttributeOptions, values *corefoundation.CFArrayRef) AXError {
	result, callErr := tryAXUIElementCopyMultipleAttributeValues(element, attributes, options, values)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyParameterizedAttributeNames func(element AXUIElementRef, names *corefoundation.CFArrayRef) AXError
var _aXUIElementCopyParameterizedAttributeNamesErr error

func tryAXUIElementCopyParameterizedAttributeNames(element AXUIElementRef, names *corefoundation.CFArrayRef) (AXError, error) {
	if _aXUIElementCopyParameterizedAttributeNames == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyParameterizedAttributeNames", "10.3", _aXUIElementCopyParameterizedAttributeNamesErr)
	}
	return _aXUIElementCopyParameterizedAttributeNames(element, names), nil
}

// AXUIElementCopyParameterizedAttributeNames returns a list of all the parameterized attributes supported by the specified accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/1458783-axuielementcopyparameterizedattr
func AXUIElementCopyParameterizedAttributeNames(element AXUIElementRef, names *corefoundation.CFArrayRef) AXError {
	result, callErr := tryAXUIElementCopyParameterizedAttributeNames(element, names)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCopyParameterizedAttributeValue func(element AXUIElementRef, parameterizedAttribute corefoundation.CFStringRef, parameter corefoundation.CFTypeRef, result *corefoundation.CFTypeRef) AXError
var _aXUIElementCopyParameterizedAttributeValueErr error

func tryAXUIElementCopyParameterizedAttributeValue(element AXUIElementRef, parameterizedAttribute corefoundation.CFStringRef, parameter corefoundation.CFTypeRef, result *corefoundation.CFTypeRef) (AXError, error) {
	if _aXUIElementCopyParameterizedAttributeValue == nil {
		return *new(AXError), symbolCallError("AXUIElementCopyParameterizedAttributeValue", "10.3", _aXUIElementCopyParameterizedAttributeValueErr)
	}
	return _aXUIElementCopyParameterizedAttributeValue(element, parameterizedAttribute, parameter, result), nil
}

// AXUIElementCopyParameterizedAttributeValue returns the value of an accessibility object's parameterized attribute.
//
// See: https://developer.apple.com/documentation/applicationservices/1461203-axuielementcopyparameterizedattr
func AXUIElementCopyParameterizedAttributeValue(element AXUIElementRef, parameterizedAttribute corefoundation.CFStringRef, parameter corefoundation.CFTypeRef, result *corefoundation.CFTypeRef) AXError {
	result0, callErr := tryAXUIElementCopyParameterizedAttributeValue(element, parameterizedAttribute, parameter, result)
	if callErr != nil {
		panic(callErr)
	}
	return result0
}

var _aXUIElementCreateApplication func(pid int32) AXUIElementRef
var _aXUIElementCreateApplicationErr error

func tryAXUIElementCreateApplication(pid int32) (AXUIElementRef, error) {
	if _aXUIElementCreateApplication == nil {
		return 0, symbolCallError("AXUIElementCreateApplication", "10.2", _aXUIElementCreateApplicationErr)
	}
	return _aXUIElementCreateApplication(pid), nil
}

// AXUIElementCreateApplication creates and returns the top-level accessibility object for the application with the specified process ID.
//
// See: https://developer.apple.com/documentation/applicationservices/1459374-axuielementcreateapplication
func AXUIElementCreateApplication(pid int32) AXUIElementRef {
	result, callErr := tryAXUIElementCreateApplication(pid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementCreateSystemWide func() AXUIElementRef
var _aXUIElementCreateSystemWideErr error

func tryAXUIElementCreateSystemWide() (AXUIElementRef, error) {
	if _aXUIElementCreateSystemWide == nil {
		return 0, symbolCallError("AXUIElementCreateSystemWide", "10.2", _aXUIElementCreateSystemWideErr)
	}
	return _aXUIElementCreateSystemWide(), nil
}

// AXUIElementCreateSystemWide returns an accessibility object that provides access to system attributes.
//
// See: https://developer.apple.com/documentation/applicationservices/1462095-axuielementcreatesystemwide
func AXUIElementCreateSystemWide() AXUIElementRef {
	result, callErr := tryAXUIElementCreateSystemWide()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementGetAttributeValueCount func(element AXUIElementRef, attribute corefoundation.CFStringRef, count *int) AXError
var _aXUIElementGetAttributeValueCountErr error

func tryAXUIElementGetAttributeValueCount(element AXUIElementRef, attribute corefoundation.CFStringRef, count *int) (AXError, error) {
	if _aXUIElementGetAttributeValueCount == nil {
		return *new(AXError), symbolCallError("AXUIElementGetAttributeValueCount", "10.2", _aXUIElementGetAttributeValueCountErr)
	}
	return _aXUIElementGetAttributeValueCount(element, attribute, count), nil
}

// AXUIElementGetAttributeValueCount returns the count of the array of an accessibility object's attribute value.
//
// See: https://developer.apple.com/documentation/applicationservices/1459066-axuielementgetattributevaluecoun
func AXUIElementGetAttributeValueCount(element AXUIElementRef, attribute corefoundation.CFStringRef, count *int) AXError {
	result, callErr := tryAXUIElementGetAttributeValueCount(element, attribute, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementGetPid func(element AXUIElementRef, pid *int32) AXError
var _aXUIElementGetPidErr error

func tryAXUIElementGetPid(element AXUIElementRef, pid *int32) (AXError, error) {
	if _aXUIElementGetPid == nil {
		return *new(AXError), symbolCallError("AXUIElementGetPid", "10.2", _aXUIElementGetPidErr)
	}
	return _aXUIElementGetPid(element, pid), nil
}

// AXUIElementGetPid returns the process ID associated with the specified accessibility object.
//
// See: https://developer.apple.com/documentation/applicationservices/1460337-axuielementgetpid
func AXUIElementGetPid(element AXUIElementRef, pid *int32) AXError {
	result, callErr := tryAXUIElementGetPid(element, pid)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementGetTypeID func() uint
var _aXUIElementGetTypeIDErr error

func tryAXUIElementGetTypeID() (uint, error) {
	if _aXUIElementGetTypeID == nil {
		return 0, symbolCallError("AXUIElementGetTypeID", "10.2", _aXUIElementGetTypeIDErr)
	}
	return _aXUIElementGetTypeID(), nil
}

// AXUIElementGetTypeID returns the unique type identifier for the AXUIElementRef type.
//
// See: https://developer.apple.com/documentation/applicationservices/1460085-axuielementgettypeid
func AXUIElementGetTypeID() uint {
	result, callErr := tryAXUIElementGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementIsAttributeSettable func(element AXUIElementRef, attribute corefoundation.CFStringRef, settable unsafe.Pointer) AXError
var _aXUIElementIsAttributeSettableErr error

func tryAXUIElementIsAttributeSettable(element AXUIElementRef, attribute corefoundation.CFStringRef, settable unsafe.Pointer) (AXError, error) {
	if _aXUIElementIsAttributeSettable == nil {
		return *new(AXError), symbolCallError("AXUIElementIsAttributeSettable", "10.2", _aXUIElementIsAttributeSettableErr)
	}
	return _aXUIElementIsAttributeSettable(element, attribute, settable), nil
}

// AXUIElementIsAttributeSettable returns whether the specified accessibility object's attribute can be modified.
//
// See: https://developer.apple.com/documentation/applicationservices/1459972-axuielementisattributesettable
func AXUIElementIsAttributeSettable(element AXUIElementRef, attribute corefoundation.CFStringRef, settable unsafe.Pointer) AXError {
	result, callErr := tryAXUIElementIsAttributeSettable(element, attribute, settable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementPerformAction func(element AXUIElementRef, action corefoundation.CFStringRef) AXError
var _aXUIElementPerformActionErr error

func tryAXUIElementPerformAction(element AXUIElementRef, action corefoundation.CFStringRef) (AXError, error) {
	if _aXUIElementPerformAction == nil {
		return *new(AXError), symbolCallError("AXUIElementPerformAction", "10.2", _aXUIElementPerformActionErr)
	}
	return _aXUIElementPerformAction(element, action), nil
}

// AXUIElementPerformAction requests that the specified accessibility object perform the specified action.
//
// See: https://developer.apple.com/documentation/applicationservices/1462091-axuielementperformaction
func AXUIElementPerformAction(element AXUIElementRef, action corefoundation.CFStringRef) AXError {
	result, callErr := tryAXUIElementPerformAction(element, action)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementPostKeyboardEvent func(arg0 AXUIElementRef, arg1 uint16, arg2 uint16, arg3 bool) AXError
var _aXUIElementPostKeyboardEventErr error

func tryAXUIElementPostKeyboardEvent(arg0 AXUIElementRef, arg1 uint16, arg2 uint16, arg3 bool) (AXError, error) {
	if _aXUIElementPostKeyboardEvent == nil {
		return *new(AXError), symbolCallError("AXUIElementPostKeyboardEvent", "10.0", _aXUIElementPostKeyboardEventErr)
	}
	return _aXUIElementPostKeyboardEvent(arg0, arg1, arg2, arg3), nil
}

// AXUIElementPostKeyboardEvent posts keys to the specified application.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1462057-axuielementpostkeyboardevent
func AXUIElementPostKeyboardEvent(arg0 AXUIElementRef, arg1 uint16, arg2 uint16, arg3 bool) AXError {
	result, callErr := tryAXUIElementPostKeyboardEvent(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementSetAttributeValue func(element AXUIElementRef, attribute corefoundation.CFStringRef, value corefoundation.CFTypeRef) AXError
var _aXUIElementSetAttributeValueErr error

func tryAXUIElementSetAttributeValue(element AXUIElementRef, attribute corefoundation.CFStringRef, value corefoundation.CFTypeRef) (AXError, error) {
	if _aXUIElementSetAttributeValue == nil {
		return *new(AXError), symbolCallError("AXUIElementSetAttributeValue", "10.2", _aXUIElementSetAttributeValueErr)
	}
	return _aXUIElementSetAttributeValue(element, attribute, value), nil
}

// AXUIElementSetAttributeValue sets the accessibility object's attribute to the specified value.
//
// See: https://developer.apple.com/documentation/applicationservices/1460434-axuielementsetattributevalue
func AXUIElementSetAttributeValue(element AXUIElementRef, attribute corefoundation.CFStringRef, value corefoundation.CFTypeRef) AXError {
	result, callErr := tryAXUIElementSetAttributeValue(element, attribute, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXUIElementSetMessagingTimeout func(element AXUIElementRef, timeoutInSeconds unsafe.Pointer) AXError
var _aXUIElementSetMessagingTimeoutErr error

func tryAXUIElementSetMessagingTimeout(element AXUIElementRef, timeoutInSeconds unsafe.Pointer) (AXError, error) {
	if _aXUIElementSetMessagingTimeout == nil {
		return *new(AXError), symbolCallError("AXUIElementSetMessagingTimeout", "10.4", _aXUIElementSetMessagingTimeoutErr)
	}
	return _aXUIElementSetMessagingTimeout(element, timeoutInSeconds), nil
}

// AXUIElementSetMessagingTimeout sets the timeout value used in the accessibility API.
//
// See: https://developer.apple.com/documentation/applicationservices/1459345-axuielementsetmessagingtimeout
func AXUIElementSetMessagingTimeout(element AXUIElementRef, timeoutInSeconds unsafe.Pointer) AXError {
	result, callErr := tryAXUIElementSetMessagingTimeout(element, timeoutInSeconds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXValueCreate func(theType AXValueType, valuePtr unsafe.Pointer) AXValueRef
var _aXValueCreateErr error

func tryAXValueCreate(theType AXValueType, valuePtr unsafe.Pointer) (AXValueRef, error) {
	if _aXValueCreate == nil {
		return 0, symbolCallError("AXValueCreate", "10.2", _aXValueCreateErr)
	}
	return _aXValueCreate(theType, valuePtr), nil
}

// AXValueCreate.
//
// See: https://developer.apple.com/documentation/applicationservices/1459351-axvaluecreate
func AXValueCreate(theType AXValueType, valuePtr unsafe.Pointer) AXValueRef {
	result, callErr := tryAXValueCreate(theType, valuePtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXValueGetType func(value AXValueRef) AXValueType
var _aXValueGetTypeErr error

func tryAXValueGetType(value AXValueRef) (AXValueType, error) {
	if _aXValueGetType == nil {
		return *new(AXValueType), symbolCallError("AXValueGetType", "10.2", _aXValueGetTypeErr)
	}
	return _aXValueGetType(value), nil
}

// AXValueGetType.
//
// See: https://developer.apple.com/documentation/applicationservices/1460911-axvaluegettype
func AXValueGetType(value AXValueRef) AXValueType {
	result, callErr := tryAXValueGetType(value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXValueGetTypeID func() uint
var _aXValueGetTypeIDErr error

func tryAXValueGetTypeID() (uint, error) {
	if _aXValueGetTypeID == nil {
		return 0, symbolCallError("AXValueGetTypeID", "10.3", _aXValueGetTypeIDErr)
	}
	return _aXValueGetTypeID(), nil
}

// AXValueGetTypeID.
//
// See: https://developer.apple.com/documentation/applicationservices/1460780-axvaluegettypeid
func AXValueGetTypeID() uint {
	result, callErr := tryAXValueGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aXValueGetValue func(value AXValueRef, theType AXValueType, valuePtr unsafe.Pointer) bool
var _aXValueGetValueErr error

func tryAXValueGetValue(value AXValueRef, theType AXValueType, valuePtr unsafe.Pointer) (bool, error) {
	if _aXValueGetValue == nil {
		return false, symbolCallError("AXValueGetValue", "10.2", _aXValueGetValueErr)
	}
	return _aXValueGetValue(value, theType, valuePtr), nil
}

// AXValueGetValue.
//
// See: https://developer.apple.com/documentation/applicationservices/1462933-axvaluegetvalue
func AXValueGetValue(value AXValueRef, theType AXValueType, valuePtr unsafe.Pointer) bool {
	result, callErr := tryAXValueGetValue(value, theType, valuePtr)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _continueSpeech func(chan_ *SpeechChannelRecord) int16
var _continueSpeechErr error

func tryContinueSpeech(chan_ *SpeechChannelRecord) (int16, error) {
	if _continueSpeech == nil {
		return 0, symbolCallError("ContinueSpeech", "10.0", _continueSpeechErr)
	}
	return _continueSpeech(chan_), nil
}

// ContinueSpeech resumes speech paused by the [PauseSpeechAt] function.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1462728-continuespeech
func ContinueSpeech(chan_ *SpeechChannelRecord) int16 {
	result, callErr := tryContinueSpeech(chan_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copyPhonemesFromText func(chan_ *SpeechChannelRecord, text corefoundation.CFStringRef, phonemes *corefoundation.CFStringRef) int16
var _copyPhonemesFromTextErr error

func tryCopyPhonemesFromText(chan_ *SpeechChannelRecord, text corefoundation.CFStringRef, phonemes *corefoundation.CFStringRef) (int16, error) {
	if _copyPhonemesFromText == nil {
		return 0, symbolCallError("CopyPhonemesFromText", "10.5", _copyPhonemesFromTextErr)
	}
	return _copyPhonemesFromText(chan_, text, phonemes), nil
}

// CopyPhonemesFromText converts the specified text string into its equivalent phonemic representation.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1460918-copyphonemesfromtext
func CopyPhonemesFromText(chan_ *SpeechChannelRecord, text corefoundation.CFStringRef, phonemes *corefoundation.CFStringRef) int16 {
	result, callErr := tryCopyPhonemesFromText(chan_, text, phonemes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copyProcessName func(arg0 *ProcessSerialNumber, arg1 corefoundation.CFStringRef) int32
var _copyProcessNameErr error

func tryCopyProcessName(arg0 *ProcessSerialNumber, arg1 corefoundation.CFStringRef) (int32, error) {
	if _copyProcessName == nil {
		return 0, symbolCallError("CopyProcessName", "10.0", _copyProcessNameErr)
	}
	return _copyProcessName(arg0, arg1), nil
}

// CopyProcessName.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501067-copyprocessname
func CopyProcessName(arg0 *ProcessSerialNumber, arg1 corefoundation.CFStringRef) int32 {
	result, callErr := tryCopyProcessName(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _copySpeechProperty func(chan_ *SpeechChannelRecord, property corefoundation.CFStringRef, object *corefoundation.CFTypeRef) int16
var _copySpeechPropertyErr error

func tryCopySpeechProperty(chan_ *SpeechChannelRecord, property corefoundation.CFStringRef, object *corefoundation.CFTypeRef) (int16, error) {
	if _copySpeechProperty == nil {
		return 0, symbolCallError("CopySpeechProperty", "10.5", _copySpeechPropertyErr)
	}
	return _copySpeechProperty(chan_, property, object), nil
}

// CopySpeechProperty gets the value associated with the specified property of a speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1459075-copyspeechproperty
func CopySpeechProperty(chan_ *SpeechChannelRecord, property corefoundation.CFStringRef, object *corefoundation.CFTypeRef) int16 {
	result, callErr := tryCopySpeechProperty(chan_, property, object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _countVoices func(numVoices *uintptr) int16
var _countVoicesErr error

func tryCountVoices(numVoices *uintptr) (int16, error) {
	if _countVoices == nil {
		return 0, symbolCallError("CountVoices", "10.0", _countVoicesErr)
	}
	return _countVoices(numVoices), nil
}

// CountVoices determines how many voices are available.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1459947-countvoices
func CountVoices(numVoices *uintptr) int16 {
	result, callErr := tryCountVoices(numVoices)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _disposeIconActionUPP func(userUPP IconActionUPP)
var _disposeIconActionUPPErr error

func tryDisposeIconActionUPP(userUPP IconActionUPP) error {
	if _disposeIconActionUPP == nil {
		return symbolCallError("DisposeIconActionUPP", "10.0", _disposeIconActionUPPErr)
	}
	_disposeIconActionUPP(userUPP)
	return nil
}

// DisposeIconActionUPP.
//
// See: https://developer.apple.com/documentation/applicationservices/1461028-disposeiconactionupp
func DisposeIconActionUPP(userUPP IconActionUPP) {
	if callErr := tryDisposeIconActionUPP(userUPP); callErr != nil {
		panic(callErr)
	}
}

var _disposeIconGetterUPP func(userUPP IconGetterUPP)
var _disposeIconGetterUPPErr error

func tryDisposeIconGetterUPP(userUPP IconGetterUPP) error {
	if _disposeIconGetterUPP == nil {
		return symbolCallError("DisposeIconGetterUPP", "10.0", _disposeIconGetterUPPErr)
	}
	_disposeIconGetterUPP(userUPP)
	return nil
}

// DisposeIconGetterUPP.
//
// See: https://developer.apple.com/documentation/applicationservices/1461061-disposeicongetterupp
func DisposeIconGetterUPP(userUPP IconGetterUPP) {
	if callErr := tryDisposeIconGetterUPP(userUPP); callErr != nil {
		panic(callErr)
	}
}

var _disposeSpeechChannel func(chan_ *SpeechChannelRecord) int16
var _disposeSpeechChannelErr error

func tryDisposeSpeechChannel(chan_ *SpeechChannelRecord) (int16, error) {
	if _disposeSpeechChannel == nil {
		return 0, symbolCallError("DisposeSpeechChannel", "10.0", _disposeSpeechChannelErr)
	}
	return _disposeSpeechChannel(chan_), nil
}

// DisposeSpeechChannel disposes of an existing speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1462081-disposespeechchannel
func DisposeSpeechChannel(chan_ *SpeechChannelRecord) int16 {
	result, callErr := tryDisposeSpeechChannel(chan_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _disposeSpeechDoneUPP func(arg0 SpeechDoneUPP)
var _disposeSpeechDoneUPPErr error

func tryDisposeSpeechDoneUPP(arg0 SpeechDoneUPP) error {
	if _disposeSpeechDoneUPP == nil {
		return symbolCallError("DisposeSpeechDoneUPP", "10.0", _disposeSpeechDoneUPPErr)
	}
	_disposeSpeechDoneUPP(arg0)
	return nil
}

// DisposeSpeechDoneUPP disposes of a universal procedure pointer (UPP) to a speech-donecallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552237-disposespeechdoneupp
func DisposeSpeechDoneUPP(arg0 SpeechDoneUPP) {
	if callErr := tryDisposeSpeechDoneUPP(arg0); callErr != nil {
		panic(callErr)
	}
}

var _disposeSpeechErrorUPP func(arg0 SpeechErrorUPP)
var _disposeSpeechErrorUPPErr error

func tryDisposeSpeechErrorUPP(arg0 SpeechErrorUPP) error {
	if _disposeSpeechErrorUPP == nil {
		return symbolCallError("DisposeSpeechErrorUPP", "10.0", _disposeSpeechErrorUPPErr)
	}
	_disposeSpeechErrorUPP(arg0)
	return nil
}

// DisposeSpeechErrorUPP disposes of a universal procedure pointer (UPP) to anerror callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552245-disposespeecherrorupp
func DisposeSpeechErrorUPP(arg0 SpeechErrorUPP) {
	if callErr := tryDisposeSpeechErrorUPP(arg0); callErr != nil {
		panic(callErr)
	}
}

var _disposeSpeechPhonemeUPP func(arg0 SpeechPhonemeUPP)
var _disposeSpeechPhonemeUPPErr error

func tryDisposeSpeechPhonemeUPP(arg0 SpeechPhonemeUPP) error {
	if _disposeSpeechPhonemeUPP == nil {
		return symbolCallError("DisposeSpeechPhonemeUPP", "10.0", _disposeSpeechPhonemeUPPErr)
	}
	_disposeSpeechPhonemeUPP(arg0)
	return nil
}

// DisposeSpeechPhonemeUPP disposes of a universal procedure pointer (UPP) to a phonemecallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552226-disposespeechphonemeupp
func DisposeSpeechPhonemeUPP(arg0 SpeechPhonemeUPP) {
	if callErr := tryDisposeSpeechPhonemeUPP(arg0); callErr != nil {
		panic(callErr)
	}
}

var _disposeSpeechSyncUPP func(arg0 SpeechSyncUPP)
var _disposeSpeechSyncUPPErr error

func tryDisposeSpeechSyncUPP(arg0 SpeechSyncUPP) error {
	if _disposeSpeechSyncUPP == nil {
		return symbolCallError("DisposeSpeechSyncUPP", "10.0", _disposeSpeechSyncUPPErr)
	}
	_disposeSpeechSyncUPP(arg0)
	return nil
}

// DisposeSpeechSyncUPP disposes of a universal procedure pointer (UPP) to a synchronizationcallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552219-disposespeechsyncupp
func DisposeSpeechSyncUPP(arg0 SpeechSyncUPP) {
	if callErr := tryDisposeSpeechSyncUPP(arg0); callErr != nil {
		panic(callErr)
	}
}

var _disposeSpeechTextDoneUPP func(arg0 SpeechTextDoneUPP)
var _disposeSpeechTextDoneUPPErr error

func tryDisposeSpeechTextDoneUPP(arg0 SpeechTextDoneUPP) error {
	if _disposeSpeechTextDoneUPP == nil {
		return symbolCallError("DisposeSpeechTextDoneUPP", "10.0", _disposeSpeechTextDoneUPPErr)
	}
	_disposeSpeechTextDoneUPP(arg0)
	return nil
}

// DisposeSpeechTextDoneUPP disposes of a universal procedure pointer (UPP) to a text-donecallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552229-disposespeechtextdoneupp
func DisposeSpeechTextDoneUPP(arg0 SpeechTextDoneUPP) {
	if callErr := tryDisposeSpeechTextDoneUPP(arg0); callErr != nil {
		panic(callErr)
	}
}

var _disposeSpeechWordUPP func(arg0 SpeechWordUPP)
var _disposeSpeechWordUPPErr error

func tryDisposeSpeechWordUPP(arg0 SpeechWordUPP) error {
	if _disposeSpeechWordUPP == nil {
		return symbolCallError("DisposeSpeechWordUPP", "10.0", _disposeSpeechWordUPPErr)
	}
	_disposeSpeechWordUPP(arg0)
	return nil
}

// DisposeSpeechWordUPP disposes of a universal procedure pointer (UPP) to a wordcallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552222-disposespeechwordupp
func DisposeSpeechWordUPP(arg0 SpeechWordUPP) {
	if callErr := tryDisposeSpeechWordUPP(arg0); callErr != nil {
		panic(callErr)
	}
}

var _exitToShell func()
var _exitToShellErr error

func tryExitToShell() error {
	if _exitToShell == nil {
		return symbolCallError("ExitToShell", "10.0", _exitToShellErr)
	}
	_exitToShell()
	return nil
}

// ExitToShell.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1500985-exittoshell
func ExitToShell() {
	if callErr := tryExitToShell(); callErr != nil {
		panic(callErr)
	}
}

var _getCurrentProcess func(arg0 *ProcessSerialNumber) int16
var _getCurrentProcessErr error

func tryGetCurrentProcess(arg0 *ProcessSerialNumber) (int16, error) {
	if _getCurrentProcess == nil {
		return 0, symbolCallError("GetCurrentProcess", "10.0", _getCurrentProcessErr)
	}
	return _getCurrentProcess(arg0), nil
}

// GetCurrentProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501115-getcurrentprocess
func GetCurrentProcess(arg0 *ProcessSerialNumber) int16 {
	result, callErr := tryGetCurrentProcess(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getFrontProcess func(arg0 *ProcessSerialNumber) int16
var _getFrontProcessErr error

func tryGetFrontProcess(arg0 *ProcessSerialNumber) (int16, error) {
	if _getFrontProcess == nil {
		return 0, symbolCallError("GetFrontProcess", "10.0", _getFrontProcessErr)
	}
	return _getFrontProcess(arg0), nil
}

// GetFrontProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501050-getfrontprocess
func GetFrontProcess(arg0 *ProcessSerialNumber) int16 {
	result, callErr := tryGetFrontProcess(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getIconFamilyData func(iconFamily uintptr, iconType uint32, h unsafe.Pointer) int16
var _getIconFamilyDataErr error

func tryGetIconFamilyData(iconFamily uintptr, iconType uint32, h unsafe.Pointer) (int16, error) {
	if _getIconFamilyData == nil {
		return 0, symbolCallError("GetIconFamilyData", "10.0", _getIconFamilyDataErr)
	}
	return _getIconFamilyData(iconFamily, iconType, h), nil
}

// GetIconFamilyData.
//
// See: https://developer.apple.com/documentation/applicationservices/1462743-geticonfamilydata
func GetIconFamilyData(iconFamily uintptr, iconType uint32, h unsafe.Pointer) int16 {
	result, callErr := tryGetIconFamilyData(iconFamily, iconType, h)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getIconRefVariant func(inIconRef unsafe.Pointer, inVariant uint32, outTransform *IconTransformType) unsafe.Pointer
var _getIconRefVariantErr error

func tryGetIconRefVariant(inIconRef unsafe.Pointer, inVariant uint32, outTransform *IconTransformType) (unsafe.Pointer, error) {
	if _getIconRefVariant == nil {
		return nil, symbolCallError("GetIconRefVariant", "10.0", _getIconRefVariantErr)
	}
	return _getIconRefVariant(inIconRef, inVariant, outTransform), nil
}

// GetIconRefVariant.
//
// See: https://developer.apple.com/documentation/applicationservices/1463088-geticonrefvariant
func GetIconRefVariant(inIconRef unsafe.Pointer, inVariant uint32, outTransform *IconTransformType) unsafe.Pointer {
	result, callErr := tryGetIconRefVariant(inIconRef, inVariant, outTransform)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getIndVoice func(index unsafe.Pointer, voice unsafe.Pointer) int16
var _getIndVoiceErr error

func tryGetIndVoice(index unsafe.Pointer, voice unsafe.Pointer) (int16, error) {
	if _getIndVoice == nil {
		return 0, symbolCallError("GetIndVoice", "10.0", _getIndVoiceErr)
	}
	return _getIndVoice(index, voice), nil
}

// GetIndVoice gets a voice specification structure for a voice bypassing an index to the [GetIndVoice] function.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1464595-getindvoice
func GetIndVoice(index unsafe.Pointer, voice unsafe.Pointer) int16 {
	result, callErr := tryGetIndVoice(index, voice)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getNextProcess func(arg0 *ProcessSerialNumber) int16
var _getNextProcessErr error

func tryGetNextProcess(arg0 *ProcessSerialNumber) (int16, error) {
	if _getNextProcess == nil {
		return 0, symbolCallError("GetNextProcess", "10.0", _getNextProcessErr)
	}
	return _getNextProcess(arg0), nil
}

// GetNextProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501061-getnextprocess
func GetNextProcess(arg0 *ProcessSerialNumber) int16 {
	result, callErr := tryGetNextProcess(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getProcessBundleLocation func(arg0 *ProcessSerialNumber, arg1 unsafe.Pointer) int32
var _getProcessBundleLocationErr error

func tryGetProcessBundleLocation(arg0 *ProcessSerialNumber, arg1 unsafe.Pointer) (int32, error) {
	if _getProcessBundleLocation == nil {
		return 0, symbolCallError("GetProcessBundleLocation", "10.0", _getProcessBundleLocationErr)
	}
	return _getProcessBundleLocation(arg0, arg1), nil
}

// GetProcessBundleLocation.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501092-getprocessbundlelocation
func GetProcessBundleLocation(arg0 *ProcessSerialNumber, arg1 unsafe.Pointer) int32 {
	result, callErr := tryGetProcessBundleLocation(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getProcessForPID func(arg0 int32, arg1 *ProcessSerialNumber) int32
var _getProcessForPIDErr error

func tryGetProcessForPID(arg0 int32, arg1 *ProcessSerialNumber) (int32, error) {
	if _getProcessForPID == nil {
		return 0, symbolCallError("GetProcessForPID", "10.0", _getProcessForPIDErr)
	}
	return _getProcessForPID(arg0, arg1), nil
}

// GetProcessForPID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501069-getprocessforpid
func GetProcessForPID(arg0 int32, arg1 *ProcessSerialNumber) int32 {
	result, callErr := tryGetProcessForPID(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getProcessInformation func(arg0 *ProcessSerialNumber, arg1 ProcessInfoRec) int16
var _getProcessInformationErr error

func tryGetProcessInformation(arg0 *ProcessSerialNumber, arg1 ProcessInfoRec) (int16, error) {
	if _getProcessInformation == nil {
		return 0, symbolCallError("GetProcessInformation", "10.0", _getProcessInformationErr)
	}
	return _getProcessInformation(arg0, arg1), nil
}

// GetProcessInformation.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501011-getprocessinformation
func GetProcessInformation(arg0 *ProcessSerialNumber, arg1 ProcessInfoRec) int16 {
	result, callErr := tryGetProcessInformation(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getProcessPID func(arg0 *ProcessSerialNumber, arg1 int32) int32
var _getProcessPIDErr error

func tryGetProcessPID(arg0 *ProcessSerialNumber, arg1 int32) (int32, error) {
	if _getProcessPID == nil {
		return 0, symbolCallError("GetProcessPID", "10.0", _getProcessPIDErr)
	}
	return _getProcessPID(arg0, arg1), nil
}

// GetProcessPID.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1500992-getprocesspid
func GetProcessPID(arg0 *ProcessSerialNumber, arg1 int32) int32 {
	result, callErr := tryGetProcessPID(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getSpeechInfo func(arg0 *SpeechChannelRecord, arg1 uint32) int16
var _getSpeechInfoErr error

func tryGetSpeechInfo(arg0 *SpeechChannelRecord, arg1 uint32) (int16, error) {
	if _getSpeechInfo == nil {
		return 0, symbolCallError("GetSpeechInfo", "10.0", _getSpeechInfoErr)
	}
	return _getSpeechInfo(arg0, arg1), nil
}

// GetSpeechInfo gets information about a designated speech channel.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552220-getspeechinfo
func GetSpeechInfo(arg0 *SpeechChannelRecord, arg1 uint32) int16 {
	result, callErr := tryGetSpeechInfo(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getSpeechPitch func(chan_ *SpeechChannelRecord, pitch *int32) int16
var _getSpeechPitchErr error

func tryGetSpeechPitch(chan_ *SpeechChannelRecord, pitch *int32) (int16, error) {
	if _getSpeechPitch == nil {
		return 0, symbolCallError("GetSpeechPitch", "10.0", _getSpeechPitchErr)
	}
	return _getSpeechPitch(chan_, pitch), nil
}

// GetSpeechPitch gets a speech channel’s current speech pitch.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1464774-getspeechpitch
func GetSpeechPitch(chan_ *SpeechChannelRecord, pitch *int32) int16 {
	result, callErr := tryGetSpeechPitch(chan_, pitch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getSpeechRate func(chan_ *SpeechChannelRecord, rate *int32) int16
var _getSpeechRateErr error

func tryGetSpeechRate(chan_ *SpeechChannelRecord, rate *int32) (int16, error) {
	if _getSpeechRate == nil {
		return 0, symbolCallError("GetSpeechRate", "10.0", _getSpeechRateErr)
	}
	return _getSpeechRate(chan_, rate), nil
}

// GetSpeechRate gets a speech channel’s current speech rate.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1460797-getspeechrate
func GetSpeechRate(chan_ *SpeechChannelRecord, rate *int32) int16 {
	result, callErr := tryGetSpeechRate(chan_, rate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getVoiceDescription func(voice unsafe.Pointer, info unsafe.Pointer, infoLength unsafe.Pointer) int16
var _getVoiceDescriptionErr error

func tryGetVoiceDescription(voice unsafe.Pointer, info unsafe.Pointer, infoLength unsafe.Pointer) (int16, error) {
	if _getVoiceDescription == nil {
		return 0, symbolCallError("GetVoiceDescription", "10.0", _getVoiceDescriptionErr)
	}
	return _getVoiceDescription(voice, info, infoLength), nil
}

// GetVoiceDescription gets a description of a voice by using the [GetVoiceDescription] function.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1463940-getvoicedescription
func GetVoiceDescription(voice unsafe.Pointer, info unsafe.Pointer, infoLength unsafe.Pointer) int16 {
	result, callErr := tryGetVoiceDescription(voice, info, infoLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _getVoiceInfo func(voice unsafe.Pointer, selector uint32, voiceInfo unsafe.Pointer) int16
var _getVoiceInfoErr error

func tryGetVoiceInfo(voice unsafe.Pointer, selector uint32, voiceInfo unsafe.Pointer) (int16, error) {
	if _getVoiceInfo == nil {
		return 0, symbolCallError("GetVoiceInfo", "10.0", _getVoiceInfoErr)
	}
	return _getVoiceInfo(voice, selector, voiceInfo), nil
}

// GetVoiceInfo gets the same information about a voice that the [GetVoiceDescription] function provides,or to determine in which file and resource a voice is stored.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1461410-getvoiceinfo
func GetVoiceInfo(voice unsafe.Pointer, selector uint32, voiceInfo unsafe.Pointer) int16 {
	result, callErr := tryGetVoiceInfo(voice, selector, voiceInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeContainsPoint func(inShape HIShapeRef, inPoint unsafe.Pointer) bool
var _hIShapeContainsPointErr error

func tryHIShapeContainsPoint(inShape HIShapeRef, inPoint unsafe.Pointer) (bool, error) {
	if _hIShapeContainsPoint == nil {
		return false, symbolCallError("HIShapeContainsPoint", "10.2", _hIShapeContainsPointErr)
	}
	return _hIShapeContainsPoint(inShape, inPoint), nil
}

// HIShapeContainsPoint.
//
// See: https://developer.apple.com/documentation/applicationservices/1464704-hishapecontainspoint
func HIShapeContainsPoint(inShape HIShapeRef, inPoint unsafe.Pointer) bool {
	result, callErr := tryHIShapeContainsPoint(inShape, inPoint)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateCopy func(inShape HIShapeRef) HIShapeRef
var _hIShapeCreateCopyErr error

func tryHIShapeCreateCopy(inShape HIShapeRef) (HIShapeRef, error) {
	if _hIShapeCreateCopy == nil {
		return 0, symbolCallError("HIShapeCreateCopy", "10.2", _hIShapeCreateCopyErr)
	}
	return _hIShapeCreateCopy(inShape), nil
}

// HIShapeCreateCopy.
//
// See: https://developer.apple.com/documentation/applicationservices/1463131-hishapecreatecopy
func HIShapeCreateCopy(inShape HIShapeRef) HIShapeRef {
	result, callErr := tryHIShapeCreateCopy(inShape)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateDifference func(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef
var _hIShapeCreateDifferenceErr error

func tryHIShapeCreateDifference(inShape1 HIShapeRef, inShape2 HIShapeRef) (HIShapeRef, error) {
	if _hIShapeCreateDifference == nil {
		return 0, symbolCallError("HIShapeCreateDifference", "10.2", _hIShapeCreateDifferenceErr)
	}
	return _hIShapeCreateDifference(inShape1, inShape2), nil
}

// HIShapeCreateDifference.
//
// See: https://developer.apple.com/documentation/applicationservices/1460901-hishapecreatedifference
func HIShapeCreateDifference(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef {
	result, callErr := tryHIShapeCreateDifference(inShape1, inShape2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateEmpty func() HIShapeRef
var _hIShapeCreateEmptyErr error

func tryHIShapeCreateEmpty() (HIShapeRef, error) {
	if _hIShapeCreateEmpty == nil {
		return 0, symbolCallError("HIShapeCreateEmpty", "10.4", _hIShapeCreateEmptyErr)
	}
	return _hIShapeCreateEmpty(), nil
}

// HIShapeCreateEmpty.
//
// See: https://developer.apple.com/documentation/applicationservices/1462651-hishapecreateempty
func HIShapeCreateEmpty() HIShapeRef {
	result, callErr := tryHIShapeCreateEmpty()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateIntersection func(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef
var _hIShapeCreateIntersectionErr error

func tryHIShapeCreateIntersection(inShape1 HIShapeRef, inShape2 HIShapeRef) (HIShapeRef, error) {
	if _hIShapeCreateIntersection == nil {
		return 0, symbolCallError("HIShapeCreateIntersection", "10.2", _hIShapeCreateIntersectionErr)
	}
	return _hIShapeCreateIntersection(inShape1, inShape2), nil
}

// HIShapeCreateIntersection.
//
// See: https://developer.apple.com/documentation/applicationservices/1464400-hishapecreateintersection
func HIShapeCreateIntersection(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef {
	result, callErr := tryHIShapeCreateIntersection(inShape1, inShape2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateMutable func() HIMutableShapeRef
var _hIShapeCreateMutableErr error

func tryHIShapeCreateMutable() (HIMutableShapeRef, error) {
	if _hIShapeCreateMutable == nil {
		return 0, symbolCallError("HIShapeCreateMutable", "10.2", _hIShapeCreateMutableErr)
	}
	return _hIShapeCreateMutable(), nil
}

// HIShapeCreateMutable.
//
// See: https://developer.apple.com/documentation/applicationservices/1459565-hishapecreatemutable
func HIShapeCreateMutable() HIMutableShapeRef {
	result, callErr := tryHIShapeCreateMutable()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateMutableCopy func(inOrig HIShapeRef) HIMutableShapeRef
var _hIShapeCreateMutableCopyErr error

func tryHIShapeCreateMutableCopy(inOrig HIShapeRef) (HIMutableShapeRef, error) {
	if _hIShapeCreateMutableCopy == nil {
		return 0, symbolCallError("HIShapeCreateMutableCopy", "10.2", _hIShapeCreateMutableCopyErr)
	}
	return _hIShapeCreateMutableCopy(inOrig), nil
}

// HIShapeCreateMutableCopy.
//
// See: https://developer.apple.com/documentation/applicationservices/1463298-hishapecreatemutablecopy
func HIShapeCreateMutableCopy(inOrig HIShapeRef) HIMutableShapeRef {
	result, callErr := tryHIShapeCreateMutableCopy(inOrig)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateMutableWithRect func(inRect unsafe.Pointer) HIMutableShapeRef
var _hIShapeCreateMutableWithRectErr error

func tryHIShapeCreateMutableWithRect(inRect unsafe.Pointer) (HIMutableShapeRef, error) {
	if _hIShapeCreateMutableWithRect == nil {
		return 0, symbolCallError("HIShapeCreateMutableWithRect", "10.5", _hIShapeCreateMutableWithRectErr)
	}
	return _hIShapeCreateMutableWithRect(inRect), nil
}

// HIShapeCreateMutableWithRect.
//
// See: https://developer.apple.com/documentation/applicationservices/1459532-hishapecreatemutablewithrect
func HIShapeCreateMutableWithRect(inRect unsafe.Pointer) HIMutableShapeRef {
	result, callErr := tryHIShapeCreateMutableWithRect(inRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateUnion func(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef
var _hIShapeCreateUnionErr error

func tryHIShapeCreateUnion(inShape1 HIShapeRef, inShape2 HIShapeRef) (HIShapeRef, error) {
	if _hIShapeCreateUnion == nil {
		return 0, symbolCallError("HIShapeCreateUnion", "10.2", _hIShapeCreateUnionErr)
	}
	return _hIShapeCreateUnion(inShape1, inShape2), nil
}

// HIShapeCreateUnion.
//
// See: https://developer.apple.com/documentation/applicationservices/1460112-hishapecreateunion
func HIShapeCreateUnion(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef {
	result, callErr := tryHIShapeCreateUnion(inShape1, inShape2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateWithQDRgn func(inRgn unsafe.Pointer) HIShapeRef
var _hIShapeCreateWithQDRgnErr error

func tryHIShapeCreateWithQDRgn(inRgn unsafe.Pointer) (HIShapeRef, error) {
	if _hIShapeCreateWithQDRgn == nil {
		return 0, symbolCallError("HIShapeCreateWithQDRgn", "10.2", _hIShapeCreateWithQDRgnErr)
	}
	return _hIShapeCreateWithQDRgn(inRgn), nil
}

// HIShapeCreateWithQDRgn.
//
// See: https://developer.apple.com/documentation/applicationservices/1464296-hishapecreatewithqdrgn
func HIShapeCreateWithQDRgn(inRgn unsafe.Pointer) HIShapeRef {
	result, callErr := tryHIShapeCreateWithQDRgn(inRgn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateWithRect func(inRect unsafe.Pointer) HIShapeRef
var _hIShapeCreateWithRectErr error

func tryHIShapeCreateWithRect(inRect unsafe.Pointer) (HIShapeRef, error) {
	if _hIShapeCreateWithRect == nil {
		return 0, symbolCallError("HIShapeCreateWithRect", "10.2", _hIShapeCreateWithRectErr)
	}
	return _hIShapeCreateWithRect(inRect), nil
}

// HIShapeCreateWithRect.
//
// See: https://developer.apple.com/documentation/applicationservices/1460650-hishapecreatewithrect
func HIShapeCreateWithRect(inRect unsafe.Pointer) HIShapeRef {
	result, callErr := tryHIShapeCreateWithRect(inRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeCreateXor func(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef
var _hIShapeCreateXorErr error

func tryHIShapeCreateXor(inShape1 HIShapeRef, inShape2 HIShapeRef) (HIShapeRef, error) {
	if _hIShapeCreateXor == nil {
		return 0, symbolCallError("HIShapeCreateXor", "10.5", _hIShapeCreateXorErr)
	}
	return _hIShapeCreateXor(inShape1, inShape2), nil
}

// HIShapeCreateXor.
//
// See: https://developer.apple.com/documentation/applicationservices/1459148-hishapecreatexor
func HIShapeCreateXor(inShape1 HIShapeRef, inShape2 HIShapeRef) HIShapeRef {
	result, callErr := tryHIShapeCreateXor(inShape1, inShape2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeDifference func(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32
var _hIShapeDifferenceErr error

func tryHIShapeDifference(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) (int32, error) {
	if _hIShapeDifference == nil {
		return 0, symbolCallError("HIShapeDifference", "10.2", _hIShapeDifferenceErr)
	}
	return _hIShapeDifference(inShape1, inShape2, outResult), nil
}

// HIShapeDifference.
//
// See: https://developer.apple.com/documentation/applicationservices/1458876-hishapedifference
func HIShapeDifference(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32 {
	result, callErr := tryHIShapeDifference(inShape1, inShape2, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeEnumerate func(inShape HIShapeRef, inOptions OptionBits, inProc unsafe.Pointer, inRefcon unsafe.Pointer) int32
var _hIShapeEnumerateErr error

func tryHIShapeEnumerate(inShape HIShapeRef, inOptions OptionBits, inProc HIShapeEnumerateProcPtr, inRefcon unsafe.Pointer) (int32, error) {
	if _hIShapeEnumerate == nil {
		return 0, symbolCallError("HIShapeEnumerate", "10.5", _hIShapeEnumerateErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 HIShapeRef, blockArg2 objc.ID, blockArg3 unsafe.Pointer) int32 {
		return inProc(blockArg0, blockArg1, objectivec.ObjectFromID(blockArg2), blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _hIShapeEnumerate(inShape, inOptions, _block0, inRefcon), nil
}

// HIShapeEnumerate.
//
// See: https://developer.apple.com/documentation/applicationservices/1459161-hishapeenumerate
func HIShapeEnumerate(inShape HIShapeRef, inOptions OptionBits, inProc HIShapeEnumerateProcPtr, inRefcon unsafe.Pointer) int32 {
	result, callErr := tryHIShapeEnumerate(inShape, inOptions, inProc, inRefcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeGetAsQDRgn func(inShape HIShapeRef, outRgn unsafe.Pointer) int32
var _hIShapeGetAsQDRgnErr error

func tryHIShapeGetAsQDRgn(inShape HIShapeRef, outRgn unsafe.Pointer) (int32, error) {
	if _hIShapeGetAsQDRgn == nil {
		return 0, symbolCallError("HIShapeGetAsQDRgn", "10.2", _hIShapeGetAsQDRgnErr)
	}
	return _hIShapeGetAsQDRgn(inShape, outRgn), nil
}

// HIShapeGetAsQDRgn.
//
// See: https://developer.apple.com/documentation/applicationservices/1464369-hishapegetasqdrgn
func HIShapeGetAsQDRgn(inShape HIShapeRef, outRgn unsafe.Pointer) int32 {
	result, callErr := tryHIShapeGetAsQDRgn(inShape, outRgn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeGetBounds func(inShape HIShapeRef, outRect unsafe.Pointer) *corefoundation.CGRect
var _hIShapeGetBoundsErr error

func tryHIShapeGetBounds(inShape HIShapeRef, outRect unsafe.Pointer) (*corefoundation.CGRect, error) {
	if _hIShapeGetBounds == nil {
		return nil, symbolCallError("HIShapeGetBounds", "10.2", _hIShapeGetBoundsErr)
	}
	return _hIShapeGetBounds(inShape, outRect), nil
}

// HIShapeGetBounds.
//
// See: https://developer.apple.com/documentation/applicationservices/1460255-hishapegetbounds
func HIShapeGetBounds(inShape HIShapeRef, outRect unsafe.Pointer) *corefoundation.CGRect {
	result, callErr := tryHIShapeGetBounds(inShape, outRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeGetTypeID func() uint
var _hIShapeGetTypeIDErr error

func tryHIShapeGetTypeID() (uint, error) {
	if _hIShapeGetTypeID == nil {
		return 0, symbolCallError("HIShapeGetTypeID", "10.2", _hIShapeGetTypeIDErr)
	}
	return _hIShapeGetTypeID(), nil
}

// HIShapeGetTypeID.
//
// See: https://developer.apple.com/documentation/applicationservices/1463371-hishapegettypeid
func HIShapeGetTypeID() uint {
	result, callErr := tryHIShapeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeInset func(inShape HIMutableShapeRef, inDX float64, inDY float64) int32
var _hIShapeInsetErr error

func tryHIShapeInset(inShape HIMutableShapeRef, inDX float64, inDY float64) (int32, error) {
	if _hIShapeInset == nil {
		return 0, symbolCallError("HIShapeInset", "10.5", _hIShapeInsetErr)
	}
	return _hIShapeInset(inShape, inDX, inDY), nil
}

// HIShapeInset.
//
// See: https://developer.apple.com/documentation/applicationservices/1463337-hishapeinset
func HIShapeInset(inShape HIMutableShapeRef, inDX float64, inDY float64) int32 {
	result, callErr := tryHIShapeInset(inShape, inDX, inDY)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeIntersect func(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32
var _hIShapeIntersectErr error

func tryHIShapeIntersect(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) (int32, error) {
	if _hIShapeIntersect == nil {
		return 0, symbolCallError("HIShapeIntersect", "10.2", _hIShapeIntersectErr)
	}
	return _hIShapeIntersect(inShape1, inShape2, outResult), nil
}

// HIShapeIntersect.
//
// See: https://developer.apple.com/documentation/applicationservices/1462645-hishapeintersect
func HIShapeIntersect(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32 {
	result, callErr := tryHIShapeIntersect(inShape1, inShape2, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeIntersectsRect func(inShape HIShapeRef, inRect unsafe.Pointer) bool
var _hIShapeIntersectsRectErr error

func tryHIShapeIntersectsRect(inShape HIShapeRef, inRect unsafe.Pointer) (bool, error) {
	if _hIShapeIntersectsRect == nil {
		return false, symbolCallError("HIShapeIntersectsRect", "10.4", _hIShapeIntersectsRectErr)
	}
	return _hIShapeIntersectsRect(inShape, inRect), nil
}

// HIShapeIntersectsRect.
//
// See: https://developer.apple.com/documentation/applicationservices/1459614-hishapeintersectsrect
func HIShapeIntersectsRect(inShape HIShapeRef, inRect unsafe.Pointer) bool {
	result, callErr := tryHIShapeIntersectsRect(inShape, inRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeIsEmpty func(inShape HIShapeRef) bool
var _hIShapeIsEmptyErr error

func tryHIShapeIsEmpty(inShape HIShapeRef) (bool, error) {
	if _hIShapeIsEmpty == nil {
		return false, symbolCallError("HIShapeIsEmpty", "10.2", _hIShapeIsEmptyErr)
	}
	return _hIShapeIsEmpty(inShape), nil
}

// HIShapeIsEmpty.
//
// See: https://developer.apple.com/documentation/applicationservices/1461878-hishapeisempty
func HIShapeIsEmpty(inShape HIShapeRef) bool {
	result, callErr := tryHIShapeIsEmpty(inShape)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeIsRectangular func(inShape HIShapeRef) bool
var _hIShapeIsRectangularErr error

func tryHIShapeIsRectangular(inShape HIShapeRef) (bool, error) {
	if _hIShapeIsRectangular == nil {
		return false, symbolCallError("HIShapeIsRectangular", "10.2", _hIShapeIsRectangularErr)
	}
	return _hIShapeIsRectangular(inShape), nil
}

// HIShapeIsRectangular.
//
// See: https://developer.apple.com/documentation/applicationservices/1461292-hishapeisrectangular
func HIShapeIsRectangular(inShape HIShapeRef) bool {
	result, callErr := tryHIShapeIsRectangular(inShape)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeOffset func(inShape HIMutableShapeRef, inDX float64, inDY float64) int32
var _hIShapeOffsetErr error

func tryHIShapeOffset(inShape HIMutableShapeRef, inDX float64, inDY float64) (int32, error) {
	if _hIShapeOffset == nil {
		return 0, symbolCallError("HIShapeOffset", "10.2", _hIShapeOffsetErr)
	}
	return _hIShapeOffset(inShape, inDX, inDY), nil
}

// HIShapeOffset.
//
// See: https://developer.apple.com/documentation/applicationservices/1461775-hishapeoffset
func HIShapeOffset(inShape HIMutableShapeRef, inDX float64, inDY float64) int32 {
	result, callErr := tryHIShapeOffset(inShape, inDX, inDY)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeReplacePathInCGContext func(inShape HIShapeRef, inContext coregraphics.CGContextRef) int32
var _hIShapeReplacePathInCGContextErr error

func tryHIShapeReplacePathInCGContext(inShape HIShapeRef, inContext coregraphics.CGContextRef) (int32, error) {
	if _hIShapeReplacePathInCGContext == nil {
		return 0, symbolCallError("HIShapeReplacePathInCGContext", "10.2", _hIShapeReplacePathInCGContextErr)
	}
	return _hIShapeReplacePathInCGContext(inShape, inContext), nil
}

// HIShapeReplacePathInCGContext.
//
// See: https://developer.apple.com/documentation/applicationservices/1460747-hishapereplacepathincgcontext
func HIShapeReplacePathInCGContext(inShape HIShapeRef, inContext coregraphics.CGContextRef) int32 {
	result, callErr := tryHIShapeReplacePathInCGContext(inShape, inContext)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeSetEmpty func(inShape HIMutableShapeRef) int32
var _hIShapeSetEmptyErr error

func tryHIShapeSetEmpty(inShape HIMutableShapeRef) (int32, error) {
	if _hIShapeSetEmpty == nil {
		return 0, symbolCallError("HIShapeSetEmpty", "10.2", _hIShapeSetEmptyErr)
	}
	return _hIShapeSetEmpty(inShape), nil
}

// HIShapeSetEmpty.
//
// See: https://developer.apple.com/documentation/applicationservices/1461259-hishapesetempty
func HIShapeSetEmpty(inShape HIMutableShapeRef) int32 {
	result, callErr := tryHIShapeSetEmpty(inShape)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeSetWithShape func(inDestShape HIMutableShapeRef, inSrcShape HIShapeRef) int32
var _hIShapeSetWithShapeErr error

func tryHIShapeSetWithShape(inDestShape HIMutableShapeRef, inSrcShape HIShapeRef) (int32, error) {
	if _hIShapeSetWithShape == nil {
		return 0, symbolCallError("HIShapeSetWithShape", "10.5", _hIShapeSetWithShapeErr)
	}
	return _hIShapeSetWithShape(inDestShape, inSrcShape), nil
}

// HIShapeSetWithShape.
//
// See: https://developer.apple.com/documentation/applicationservices/1462473-hishapesetwithshape
func HIShapeSetWithShape(inDestShape HIMutableShapeRef, inSrcShape HIShapeRef) int32 {
	result, callErr := tryHIShapeSetWithShape(inDestShape, inSrcShape)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeUnion func(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32
var _hIShapeUnionErr error

func tryHIShapeUnion(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) (int32, error) {
	if _hIShapeUnion == nil {
		return 0, symbolCallError("HIShapeUnion", "10.2", _hIShapeUnionErr)
	}
	return _hIShapeUnion(inShape1, inShape2, outResult), nil
}

// HIShapeUnion.
//
// See: https://developer.apple.com/documentation/applicationservices/1459542-hishapeunion
func HIShapeUnion(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32 {
	result, callErr := tryHIShapeUnion(inShape1, inShape2, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeUnionWithRect func(inShape HIMutableShapeRef, inRect unsafe.Pointer) int32
var _hIShapeUnionWithRectErr error

func tryHIShapeUnionWithRect(inShape HIMutableShapeRef, inRect unsafe.Pointer) (int32, error) {
	if _hIShapeUnionWithRect == nil {
		return 0, symbolCallError("HIShapeUnionWithRect", "10.5", _hIShapeUnionWithRectErr)
	}
	return _hIShapeUnionWithRect(inShape, inRect), nil
}

// HIShapeUnionWithRect.
//
// See: https://developer.apple.com/documentation/applicationservices/1462757-hishapeunionwithrect
func HIShapeUnionWithRect(inShape HIMutableShapeRef, inRect unsafe.Pointer) int32 {
	result, callErr := tryHIShapeUnionWithRect(inShape, inRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _hIShapeXor func(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32
var _hIShapeXorErr error

func tryHIShapeXor(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) (int32, error) {
	if _hIShapeXor == nil {
		return 0, symbolCallError("HIShapeXor", "10.5", _hIShapeXorErr)
	}
	return _hIShapeXor(inShape1, inShape2, outResult), nil
}

// HIShapeXor.
//
// See: https://developer.apple.com/documentation/applicationservices/1461294-hishapexor
func HIShapeXor(inShape1 HIShapeRef, inShape2 HIShapeRef, outResult HIMutableShapeRef) int32 {
	result, callErr := tryHIShapeXor(inShape1, inShape2, outResult)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCAddMapEntry func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICMapEntry) int32
var _iCAddMapEntryErr error

func tryICAddMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICMapEntry) (int32, error) {
	if _iCAddMapEntry == nil {
		return 0, symbolCallError("ICAddMapEntry", "10.0", _iCAddMapEntryErr)
	}
	return _iCAddMapEntry(arg0, arg1, arg2), nil
}

// ICAddMapEntry.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578495-icaddmapentry
func ICAddMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICMapEntry) int32 {
	result, callErr := tryICAddMapEntry(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCAddProfile func(arg0 ICInstance, arg1 ICProfileID, arg2 ICProfileID) int32
var _iCAddProfileErr error

func tryICAddProfile(arg0 ICInstance, arg1 ICProfileID, arg2 ICProfileID) (int32, error) {
	if _iCAddProfile == nil {
		return 0, symbolCallError("ICAddProfile", "10.0", _iCAddProfileErr)
	}
	return _iCAddProfile(arg0, arg1, arg2), nil
}

// ICAddProfile.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578498-icaddprofile
func ICAddProfile(arg0 ICInstance, arg1 ICProfileID, arg2 ICProfileID) int32 {
	result, callErr := tryICAddProfile(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCBegin func(arg0 ICInstance, arg1 ICPerm) int32
var _iCBeginErr error

func tryICBegin(arg0 ICInstance, arg1 ICPerm) (int32, error) {
	if _iCBegin == nil {
		return 0, symbolCallError("ICBegin", "10.0", _iCBeginErr)
	}
	return _iCBegin(arg0, arg1), nil
}

// ICBegin.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578517-icbegin
func ICBegin(arg0 ICInstance, arg1 ICPerm) int32 {
	result, callErr := tryICBegin(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCCountMapEntries func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int) int32
var _iCCountMapEntriesErr error

func tryICCountMapEntries(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int) (int32, error) {
	if _iCCountMapEntries == nil {
		return 0, symbolCallError("ICCountMapEntries", "10.0", _iCCountMapEntriesErr)
	}
	return _iCCountMapEntries(arg0, arg1, arg2), nil
}

// ICCountMapEntries.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578509-iccountmapentries
func ICCountMapEntries(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int) int32 {
	result, callErr := tryICCountMapEntries(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCCountPref func(arg0 ICInstance, arg1 int) int32
var _iCCountPrefErr error

func tryICCountPref(arg0 ICInstance, arg1 int) (int32, error) {
	if _iCCountPref == nil {
		return 0, symbolCallError("ICCountPref", "10.0", _iCCountPrefErr)
	}
	return _iCCountPref(arg0, arg1), nil
}

// ICCountPref.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578542-iccountpref
func ICCountPref(arg0 ICInstance, arg1 int) int32 {
	result, callErr := tryICCountPref(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCCountProfiles func(arg0 ICInstance, arg1 int) int32
var _iCCountProfilesErr error

func tryICCountProfiles(arg0 ICInstance, arg1 int) (int32, error) {
	if _iCCountProfiles == nil {
		return 0, symbolCallError("ICCountProfiles", "10.0", _iCCountProfilesErr)
	}
	return _iCCountProfiles(arg0, arg1), nil
}

// ICCountProfiles.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578494-iccountprofiles
func ICCountProfiles(arg0 ICInstance, arg1 int) int32 {
	result, callErr := tryICCountProfiles(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCCreateGURLEvent func(arg0 ICInstance, arg1 uint32, arg2 unsafe.Pointer, arg3 coreservices.AEDesc) int32
var _iCCreateGURLEventErr error

func tryICCreateGURLEvent(arg0 ICInstance, arg1 uint32, arg2 unsafe.Pointer, arg3 coreservices.AEDesc) (int32, error) {
	if _iCCreateGURLEvent == nil {
		return 0, symbolCallError("ICCreateGURLEvent", "10.0", _iCCreateGURLEventErr)
	}
	return _iCCreateGURLEvent(arg0, arg1, arg2, arg3), nil
}

// ICCreateGURLEvent.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578532-iccreategurlevent
func ICCreateGURLEvent(arg0 ICInstance, arg1 uint32, arg2 unsafe.Pointer, arg3 coreservices.AEDesc) int32 {
	result, callErr := tryICCreateGURLEvent(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCDeleteMapEntry func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int) int32
var _iCDeleteMapEntryErr error

func tryICDeleteMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int) (int32, error) {
	if _iCDeleteMapEntry == nil {
		return 0, symbolCallError("ICDeleteMapEntry", "10.0", _iCDeleteMapEntryErr)
	}
	return _iCDeleteMapEntry(arg0, arg1, arg2), nil
}

// ICDeleteMapEntry.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578488-icdeletemapentry
func ICDeleteMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int) int32 {
	result, callErr := tryICDeleteMapEntry(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCDeletePref func(arg0 ICInstance, arg1 unsafe.Pointer) int32
var _iCDeletePrefErr error

func tryICDeletePref(arg0 ICInstance, arg1 unsafe.Pointer) (int32, error) {
	if _iCDeletePref == nil {
		return 0, symbolCallError("ICDeletePref", "10.0", _iCDeletePrefErr)
	}
	return _iCDeletePref(arg0, arg1), nil
}

// ICDeletePref.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578534-icdeletepref
func ICDeletePref(arg0 ICInstance, arg1 unsafe.Pointer) int32 {
	result, callErr := tryICDeletePref(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCDeleteProfile func(arg0 ICInstance, arg1 ICProfileID) int32
var _iCDeleteProfileErr error

func tryICDeleteProfile(arg0 ICInstance, arg1 ICProfileID) (int32, error) {
	if _iCDeleteProfile == nil {
		return 0, symbolCallError("ICDeleteProfile", "10.0", _iCDeleteProfileErr)
	}
	return _iCDeleteProfile(arg0, arg1), nil
}

// ICDeleteProfile.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578483-icdeleteprofile
func ICDeleteProfile(arg0 ICInstance, arg1 ICProfileID) int32 {
	result, callErr := tryICDeleteProfile(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCEditPreferences func(arg0 ICInstance, arg1 unsafe.Pointer) int32
var _iCEditPreferencesErr error

func tryICEditPreferences(arg0 ICInstance, arg1 unsafe.Pointer) (int32, error) {
	if _iCEditPreferences == nil {
		return 0, symbolCallError("ICEditPreferences", "10.0", _iCEditPreferencesErr)
	}
	return _iCEditPreferences(arg0, arg1), nil
}

// ICEditPreferences.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578512-iceditpreferences
func ICEditPreferences(arg0 ICInstance, arg1 unsafe.Pointer) int32 {
	result, callErr := tryICEditPreferences(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCEnd func(arg0 ICInstance) int32
var _iCEndErr error

func tryICEnd(arg0 ICInstance) (int32, error) {
	if _iCEnd == nil {
		return 0, symbolCallError("ICEnd", "10.0", _iCEndErr)
	}
	return _iCEnd(arg0), nil
}

// ICEnd.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578544-icend
func ICEnd(arg0 ICInstance) int32 {
	result, callErr := tryICEnd(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCFindPrefHandle func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) int32
var _iCFindPrefHandleErr error

func tryICFindPrefHandle(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) (int32, error) {
	if _iCFindPrefHandle == nil {
		return 0, symbolCallError("ICFindPrefHandle", "10.0", _iCFindPrefHandleErr)
	}
	return _iCFindPrefHandle(arg0, arg1, arg2, arg3), nil
}

// ICFindPrefHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578526-icfindprefhandle
func ICFindPrefHandle(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) int32 {
	result, callErr := tryICFindPrefHandle(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetConfigName func(arg0 ICInstance, arg1 bool, arg2 unsafe.Pointer) int32
var _iCGetConfigNameErr error

func tryICGetConfigName(arg0 ICInstance, arg1 bool, arg2 unsafe.Pointer) (int32, error) {
	if _iCGetConfigName == nil {
		return 0, symbolCallError("ICGetConfigName", "10.0", _iCGetConfigNameErr)
	}
	return _iCGetConfigName(arg0, arg1, arg2), nil
}

// ICGetConfigName.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578511-icgetconfigname
func ICGetConfigName(arg0 ICInstance, arg1 bool, arg2 unsafe.Pointer) int32 {
	result, callErr := tryICGetConfigName(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetCurrentProfile func(arg0 ICInstance, arg1 ICProfileID) int32
var _iCGetCurrentProfileErr error

func tryICGetCurrentProfile(arg0 ICInstance, arg1 ICProfileID) (int32, error) {
	if _iCGetCurrentProfile == nil {
		return 0, symbolCallError("ICGetCurrentProfile", "10.0", _iCGetCurrentProfileErr)
	}
	return _iCGetCurrentProfile(arg0, arg1), nil
}

// ICGetCurrentProfile.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578505-icgetcurrentprofile
func ICGetCurrentProfile(arg0 ICInstance, arg1 ICProfileID) int32 {
	result, callErr := tryICGetCurrentProfile(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetDefaultPref func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32
var _iCGetDefaultPrefErr error

func tryICGetDefaultPref(arg0 ICInstance, arg1 unsafe.Pointer, arg2 unsafe.Pointer) (int32, error) {
	if _iCGetDefaultPref == nil {
		return 0, symbolCallError("ICGetDefaultPref", "10.0", _iCGetDefaultPrefErr)
	}
	return _iCGetDefaultPref(arg0, arg1, arg2), nil
}

// ICGetDefaultPref.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578528-icgetdefaultpref
func ICGetDefaultPref(arg0 ICInstance, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32 {
	result, callErr := tryICGetDefaultPref(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetIndMapEntry func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 ICMapEntry) int32
var _iCGetIndMapEntryErr error

func tryICGetIndMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 ICMapEntry) (int32, error) {
	if _iCGetIndMapEntry == nil {
		return 0, symbolCallError("ICGetIndMapEntry", "10.0", _iCGetIndMapEntryErr)
	}
	return _iCGetIndMapEntry(arg0, arg1, arg2, arg3, arg4), nil
}

// ICGetIndMapEntry.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578493-icgetindmapentry
func ICGetIndMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 ICMapEntry) int32 {
	result, callErr := tryICGetIndMapEntry(arg0, arg1, arg2, arg3, arg4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetIndPref func(arg0 ICInstance, arg1 int, arg2 unsafe.Pointer) int32
var _iCGetIndPrefErr error

func tryICGetIndPref(arg0 ICInstance, arg1 int, arg2 unsafe.Pointer) (int32, error) {
	if _iCGetIndPref == nil {
		return 0, symbolCallError("ICGetIndPref", "10.0", _iCGetIndPrefErr)
	}
	return _iCGetIndPref(arg0, arg1, arg2), nil
}

// ICGetIndPref.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578503-icgetindpref
func ICGetIndPref(arg0 ICInstance, arg1 int, arg2 unsafe.Pointer) int32 {
	result, callErr := tryICGetIndPref(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetIndProfile func(arg0 ICInstance, arg1 int, arg2 ICProfileID) int32
var _iCGetIndProfileErr error

func tryICGetIndProfile(arg0 ICInstance, arg1 int, arg2 ICProfileID) (int32, error) {
	if _iCGetIndProfile == nil {
		return 0, symbolCallError("ICGetIndProfile", "10.0", _iCGetIndProfileErr)
	}
	return _iCGetIndProfile(arg0, arg1, arg2), nil
}

// ICGetIndProfile.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578519-icgetindprofile
func ICGetIndProfile(arg0 ICInstance, arg1 int, arg2 ICProfileID) int32 {
	result, callErr := tryICGetIndProfile(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetMapEntry func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 ICMapEntry) int32
var _iCGetMapEntryErr error

func tryICGetMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 ICMapEntry) (int32, error) {
	if _iCGetMapEntry == nil {
		return 0, symbolCallError("ICGetMapEntry", "10.0", _iCGetMapEntryErr)
	}
	return _iCGetMapEntry(arg0, arg1, arg2, arg3), nil
}

// ICGetMapEntry.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578510-icgetmapentry
func ICGetMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 ICMapEntry) int32 {
	result, callErr := tryICGetMapEntry(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetPerm func(arg0 ICInstance, arg1 ICPerm) int32
var _iCGetPermErr error

func tryICGetPerm(arg0 ICInstance, arg1 ICPerm) (int32, error) {
	if _iCGetPerm == nil {
		return 0, symbolCallError("ICGetPerm", "10.0", _iCGetPermErr)
	}
	return _iCGetPerm(arg0, arg1), nil
}

// ICGetPerm.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578492-icgetperm
func ICGetPerm(arg0 ICInstance, arg1 ICPerm) int32 {
	result, callErr := tryICGetPerm(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetPref func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 int) int32
var _iCGetPrefErr error

func tryICGetPref(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 int) (int32, error) {
	if _iCGetPref == nil {
		return 0, symbolCallError("ICGetPref", "10.0", _iCGetPrefErr)
	}
	return _iCGetPref(arg0, arg1, arg2, arg3), nil
}

// ICGetPref.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578522-icgetpref
func ICGetPref(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 int) int32 {
	result, callErr := tryICGetPref(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetPrefHandle func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) int32
var _iCGetPrefHandleErr error

func tryICGetPrefHandle(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) (int32, error) {
	if _iCGetPrefHandle == nil {
		return 0, symbolCallError("ICGetPrefHandle", "10.0", _iCGetPrefHandleErr)
	}
	return _iCGetPrefHandle(arg0, arg1, arg2, arg3), nil
}

// ICGetPrefHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578489-icgetprefhandle
func ICGetPrefHandle(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) int32 {
	result, callErr := tryICGetPrefHandle(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetProfileName func(arg0 ICInstance, arg1 ICProfileID, arg2 unsafe.Pointer) int32
var _iCGetProfileNameErr error

func tryICGetProfileName(arg0 ICInstance, arg1 ICProfileID, arg2 unsafe.Pointer) (int32, error) {
	if _iCGetProfileName == nil {
		return 0, symbolCallError("ICGetProfileName", "10.0", _iCGetProfileNameErr)
	}
	return _iCGetProfileName(arg0, arg1, arg2), nil
}

// ICGetProfileName.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578523-icgetprofilename
func ICGetProfileName(arg0 ICInstance, arg1 ICProfileID, arg2 unsafe.Pointer) int32 {
	result, callErr := tryICGetProfileName(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetSeed func(arg0 ICInstance, arg1 int) int32
var _iCGetSeedErr error

func tryICGetSeed(arg0 ICInstance, arg1 int) (int32, error) {
	if _iCGetSeed == nil {
		return 0, symbolCallError("ICGetSeed", "10.0", _iCGetSeedErr)
	}
	return _iCGetSeed(arg0, arg1), nil
}

// ICGetSeed.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578514-icgetseed
func ICGetSeed(arg0 ICInstance, arg1 int) int32 {
	result, callErr := tryICGetSeed(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCGetVersion func(arg0 ICInstance, arg1 int, arg2 uint32) int32
var _iCGetVersionErr error

func tryICGetVersion(arg0 ICInstance, arg1 int, arg2 uint32) (int32, error) {
	if _iCGetVersion == nil {
		return 0, symbolCallError("ICGetVersion", "10.0", _iCGetVersionErr)
	}
	return _iCGetVersion(arg0, arg1, arg2), nil
}

// ICGetVersion.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578536-icgetversion
func ICGetVersion(arg0 ICInstance, arg1 int, arg2 uint32) int32 {
	result, callErr := tryICGetVersion(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCLaunchURL func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 int) int32
var _iCLaunchURLErr error

func tryICLaunchURL(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 int) (int32, error) {
	if _iCLaunchURL == nil {
		return 0, symbolCallError("ICLaunchURL", "10.0", _iCLaunchURLErr)
	}
	return _iCLaunchURL(arg0, arg1, arg2, arg3, arg4), nil
}

// ICLaunchURL.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578504-iclaunchurl
func ICLaunchURL(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 int) int32 {
	result, callErr := tryICLaunchURL(arg0, arg1, arg2, arg3, arg4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCMapEntriesFilename func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 ICMapEntry) int32
var _iCMapEntriesFilenameErr error

func tryICMapEntriesFilename(arg0 ICInstance, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 ICMapEntry) (int32, error) {
	if _iCMapEntriesFilename == nil {
		return 0, symbolCallError("ICMapEntriesFilename", "10.0", _iCMapEntriesFilenameErr)
	}
	return _iCMapEntriesFilename(arg0, arg1, arg2, arg3), nil
}

// ICMapEntriesFilename.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578539-icmapentriesfilename
func ICMapEntriesFilename(arg0 ICInstance, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 ICMapEntry) int32 {
	result, callErr := tryICMapEntriesFilename(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCMapEntriesTypeCreator func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32, arg4 unsafe.Pointer, arg5 ICMapEntry) int32
var _iCMapEntriesTypeCreatorErr error

func tryICMapEntriesTypeCreator(arg0 ICInstance, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32, arg4 unsafe.Pointer, arg5 ICMapEntry) (int32, error) {
	if _iCMapEntriesTypeCreator == nil {
		return 0, symbolCallError("ICMapEntriesTypeCreator", "10.0", _iCMapEntriesTypeCreatorErr)
	}
	return _iCMapEntriesTypeCreator(arg0, arg1, arg2, arg3, arg4, arg5), nil
}

// ICMapEntriesTypeCreator.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578486-icmapentriestypecreator
func ICMapEntriesTypeCreator(arg0 ICInstance, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32, arg4 unsafe.Pointer, arg5 ICMapEntry) int32 {
	result, callErr := tryICMapEntriesTypeCreator(arg0, arg1, arg2, arg3, arg4, arg5)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCMapFilename func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICMapEntry) int32
var _iCMapFilenameErr error

func tryICMapFilename(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICMapEntry) (int32, error) {
	if _iCMapFilename == nil {
		return 0, symbolCallError("ICMapFilename", "10.0", _iCMapFilenameErr)
	}
	return _iCMapFilename(arg0, arg1, arg2), nil
}

// ICMapFilename.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578490-icmapfilename
func ICMapFilename(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICMapEntry) int32 {
	result, callErr := tryICMapFilename(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCMapTypeCreator func(arg0 ICInstance, arg1 uint32, arg2 uint32, arg3 unsafe.Pointer, arg4 ICMapEntry) int32
var _iCMapTypeCreatorErr error

func tryICMapTypeCreator(arg0 ICInstance, arg1 uint32, arg2 uint32, arg3 unsafe.Pointer, arg4 ICMapEntry) (int32, error) {
	if _iCMapTypeCreator == nil {
		return 0, symbolCallError("ICMapTypeCreator", "10.0", _iCMapTypeCreatorErr)
	}
	return _iCMapTypeCreator(arg0, arg1, arg2, arg3, arg4), nil
}

// ICMapTypeCreator.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578520-icmaptypecreator
func ICMapTypeCreator(arg0 ICInstance, arg1 uint32, arg2 uint32, arg3 unsafe.Pointer, arg4 ICMapEntry) int32 {
	result, callErr := tryICMapTypeCreator(arg0, arg1, arg2, arg3, arg4)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCParseURL func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 int, arg5 unsafe.Pointer) int32
var _iCParseURLErr error

func tryICParseURL(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 int, arg5 unsafe.Pointer) (int32, error) {
	if _iCParseURL == nil {
		return 0, symbolCallError("ICParseURL", "10.0", _iCParseURLErr)
	}
	return _iCParseURL(arg0, arg1, arg2, arg3, arg4, arg5), nil
}

// ICParseURL.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578515-icparseurl
func ICParseURL(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 int, arg4 int, arg5 unsafe.Pointer) int32 {
	result, callErr := tryICParseURL(arg0, arg1, arg2, arg3, arg4, arg5)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCSendGURLEvent func(arg0 ICInstance, arg1 coreservices.AEDesc) int32
var _iCSendGURLEventErr error

func tryICSendGURLEvent(arg0 ICInstance, arg1 coreservices.AEDesc) (int32, error) {
	if _iCSendGURLEvent == nil {
		return 0, symbolCallError("ICSendGURLEvent", "10.0", _iCSendGURLEventErr)
	}
	return _iCSendGURLEvent(arg0, arg1), nil
}

// ICSendGURLEvent.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578487-icsendgurlevent
func ICSendGURLEvent(arg0 ICInstance, arg1 coreservices.AEDesc) int32 {
	result, callErr := tryICSendGURLEvent(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCSetCurrentProfile func(arg0 ICInstance, arg1 ICProfileID) int32
var _iCSetCurrentProfileErr error

func tryICSetCurrentProfile(arg0 ICInstance, arg1 ICProfileID) (int32, error) {
	if _iCSetCurrentProfile == nil {
		return 0, symbolCallError("ICSetCurrentProfile", "10.0", _iCSetCurrentProfileErr)
	}
	return _iCSetCurrentProfile(arg0, arg1), nil
}

// ICSetCurrentProfile.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578485-icsetcurrentprofile
func ICSetCurrentProfile(arg0 ICInstance, arg1 ICProfileID) int32 {
	result, callErr := tryICSetCurrentProfile(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCSetMapEntry func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 ICMapEntry) int32
var _iCSetMapEntryErr error

func tryICSetMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 ICMapEntry) (int32, error) {
	if _iCSetMapEntry == nil {
		return 0, symbolCallError("ICSetMapEntry", "10.0", _iCSetMapEntryErr)
	}
	return _iCSetMapEntry(arg0, arg1, arg2, arg3), nil
}

// ICSetMapEntry.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578540-icsetmapentry
func ICSetMapEntry(arg0 ICInstance, arg1 unsafe.Pointer, arg2 int, arg3 ICMapEntry) int32 {
	result, callErr := tryICSetMapEntry(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCSetPref func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 int) int32
var _iCSetPrefErr error

func tryICSetPref(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 int) (int32, error) {
	if _iCSetPref == nil {
		return 0, symbolCallError("ICSetPref", "10.0", _iCSetPrefErr)
	}
	return _iCSetPref(arg0, arg1, arg2, arg3), nil
}

// ICSetPref.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578533-icsetpref
func ICSetPref(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 int) int32 {
	result, callErr := tryICSetPref(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCSetPrefHandle func(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) int32
var _iCSetPrefHandleErr error

func tryICSetPrefHandle(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) (int32, error) {
	if _iCSetPrefHandle == nil {
		return 0, symbolCallError("ICSetPrefHandle", "10.0", _iCSetPrefHandleErr)
	}
	return _iCSetPrefHandle(arg0, arg1, arg2, arg3), nil
}

// ICSetPrefHandle.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578545-icsetprefhandle
func ICSetPrefHandle(arg0 ICInstance, arg1 unsafe.Pointer, arg2 ICAttr, arg3 unsafe.Pointer) int32 {
	result, callErr := tryICSetPrefHandle(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCSetProfileName func(arg0 ICInstance, arg1 ICProfileID, arg2 unsafe.Pointer) int32
var _iCSetProfileNameErr error

func tryICSetProfileName(arg0 ICInstance, arg1 ICProfileID, arg2 unsafe.Pointer) (int32, error) {
	if _iCSetProfileName == nil {
		return 0, symbolCallError("ICSetProfileName", "10.0", _iCSetProfileNameErr)
	}
	return _iCSetProfileName(arg0, arg1, arg2), nil
}

// ICSetProfileName.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578501-icsetprofilename
func ICSetProfileName(arg0 ICInstance, arg1 ICProfileID, arg2 unsafe.Pointer) int32 {
	result, callErr := tryICSetProfileName(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCStart func(arg0 ICInstance, arg1 uint32) int32
var _iCStartErr error

func tryICStart(arg0 ICInstance, arg1 uint32) (int32, error) {
	if _iCStart == nil {
		return 0, symbolCallError("ICStart", "10.0", _iCStartErr)
	}
	return _iCStart(arg0, arg1), nil
}

// ICStart.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578513-icstart
func ICStart(arg0 ICInstance, arg1 uint32) int32 {
	result, callErr := tryICStart(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iCStop func(arg0 ICInstance) int32
var _iCStopErr error

func tryICStop(arg0 ICInstance) (int32, error) {
	if _iCStop == nil {
		return 0, symbolCallError("ICStop", "10.0", _iCStopErr)
	}
	return _iCStop(arg0), nil
}

// ICStop.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/applicationservices/1578518-icstop
func ICStop(arg0 ICInstance) int32 {
	result, callErr := tryICStop(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iconRefContainsCGPoint func(testPt unsafe.Pointer, iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) bool
var _iconRefContainsCGPointErr error

func tryIconRefContainsCGPoint(testPt unsafe.Pointer, iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) (bool, error) {
	if _iconRefContainsCGPoint == nil {
		return false, symbolCallError("IconRefContainsCGPoint", "10.5", _iconRefContainsCGPointErr)
	}
	return _iconRefContainsCGPoint(testPt, iconRect, align, iconServicesUsageFlags, theIconRef), nil
}

// IconRefContainsCGPoint.
//
// See: https://developer.apple.com/documentation/applicationservices/1461049-iconrefcontainscgpoint
func IconRefContainsCGPoint(testPt unsafe.Pointer, iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) bool {
	result, callErr := tryIconRefContainsCGPoint(testPt, iconRect, align, iconServicesUsageFlags, theIconRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iconRefIntersectsCGRect func(testRect unsafe.Pointer, iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) bool
var _iconRefIntersectsCGRectErr error

func tryIconRefIntersectsCGRect(testRect unsafe.Pointer, iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) (bool, error) {
	if _iconRefIntersectsCGRect == nil {
		return false, symbolCallError("IconRefIntersectsCGRect", "10.5", _iconRefIntersectsCGRectErr)
	}
	return _iconRefIntersectsCGRect(testRect, iconRect, align, iconServicesUsageFlags, theIconRef), nil
}

// IconRefIntersectsCGRect.
//
// See: https://developer.apple.com/documentation/applicationservices/1462553-iconrefintersectscgrect
func IconRefIntersectsCGRect(testRect unsafe.Pointer, iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) bool {
	result, callErr := tryIconRefIntersectsCGRect(testRect, iconRect, align, iconServicesUsageFlags, theIconRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iconRefToHIShape func(iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) HIShapeRef
var _iconRefToHIShapeErr error

func tryIconRefToHIShape(iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) (HIShapeRef, error) {
	if _iconRefToHIShape == nil {
		return 0, symbolCallError("IconRefToHIShape", "10.5", _iconRefToHIShapeErr)
	}
	return _iconRefToHIShape(iconRect, align, iconServicesUsageFlags, theIconRef), nil
}

// IconRefToHIShape.
//
// See: https://developer.apple.com/documentation/applicationservices/1464005-iconreftohishape
func IconRefToHIShape(iconRect unsafe.Pointer, align IconAlignmentType, iconServicesUsageFlags uintptr, theIconRef unsafe.Pointer) HIShapeRef {
	result, callErr := tryIconRefToHIShape(iconRect, align, iconServicesUsageFlags, theIconRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _iconRefToIconFamily func(theIconRef unsafe.Pointer, whichIcons IconSelectorValue, iconFamily uintptr) int16
var _iconRefToIconFamilyErr error

func tryIconRefToIconFamily(theIconRef unsafe.Pointer, whichIcons IconSelectorValue, iconFamily uintptr) (int16, error) {
	if _iconRefToIconFamily == nil {
		return 0, symbolCallError("IconRefToIconFamily", "10.0", _iconRefToIconFamilyErr)
	}
	return _iconRefToIconFamily(theIconRef, whichIcons, iconFamily), nil
}

// IconRefToIconFamily.
//
// See: https://developer.apple.com/documentation/applicationservices/1459977-iconreftoiconfamily
func IconRefToIconFamily(theIconRef unsafe.Pointer, whichIcons IconSelectorValue, iconFamily uintptr) int16 {
	result, callErr := tryIconRefToIconFamily(theIconRef, whichIcons, iconFamily)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _invokeIconActionUPP func(theType uint32, theIcon unsafe.Pointer, yourDataPtr unsafe.Pointer, userUPP IconActionUPP) int16
var _invokeIconActionUPPErr error

func tryInvokeIconActionUPP(theType uint32, theIcon unsafe.Pointer, yourDataPtr unsafe.Pointer, userUPP IconActionUPP) (int16, error) {
	if _invokeIconActionUPP == nil {
		return 0, symbolCallError("InvokeIconActionUPP", "10.0", _invokeIconActionUPPErr)
	}
	return _invokeIconActionUPP(theType, theIcon, yourDataPtr, userUPP), nil
}

// InvokeIconActionUPP.
//
// See: https://developer.apple.com/documentation/applicationservices/1464116-invokeiconactionupp
func InvokeIconActionUPP(theType uint32, theIcon unsafe.Pointer, yourDataPtr unsafe.Pointer, userUPP IconActionUPP) int16 {
	result, callErr := tryInvokeIconActionUPP(theType, theIcon, yourDataPtr, userUPP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _invokeIconGetterUPP func(theType uint32, yourDataPtr unsafe.Pointer, userUPP IconGetterUPP) unsafe.Pointer
var _invokeIconGetterUPPErr error

func tryInvokeIconGetterUPP(theType uint32, yourDataPtr unsafe.Pointer, userUPP IconGetterUPP) (unsafe.Pointer, error) {
	if _invokeIconGetterUPP == nil {
		return nil, symbolCallError("InvokeIconGetterUPP", "10.0", _invokeIconGetterUPPErr)
	}
	return _invokeIconGetterUPP(theType, yourDataPtr, userUPP), nil
}

// InvokeIconGetterUPP.
//
// See: https://developer.apple.com/documentation/applicationservices/1460976-invokeicongetterupp
func InvokeIconGetterUPP(theType uint32, yourDataPtr unsafe.Pointer, userUPP IconGetterUPP) unsafe.Pointer {
	result, callErr := tryInvokeIconGetterUPP(theType, yourDataPtr, userUPP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _invokeSpeechDoneUPP func(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 SpeechDoneUPP)
var _invokeSpeechDoneUPPErr error

func tryInvokeSpeechDoneUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 SpeechDoneUPP) error {
	if _invokeSpeechDoneUPP == nil {
		return symbolCallError("InvokeSpeechDoneUPP", "10.0", _invokeSpeechDoneUPPErr)
	}
	_invokeSpeechDoneUPP(arg0, arg1, arg2)
	return nil
}

// InvokeSpeechDoneUPP invokes your speech-done callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552215-invokespeechdoneupp
func InvokeSpeechDoneUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 SpeechDoneUPP) {
	if callErr := tryInvokeSpeechDoneUPP(arg0, arg1, arg2); callErr != nil {
		panic(callErr)
	}
}

var _invokeSpeechErrorUPP func(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 int16, arg3 int, arg4 SpeechErrorUPP)
var _invokeSpeechErrorUPPErr error

func tryInvokeSpeechErrorUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 int16, arg3 int, arg4 SpeechErrorUPP) error {
	if _invokeSpeechErrorUPP == nil {
		return symbolCallError("InvokeSpeechErrorUPP", "10.0", _invokeSpeechErrorUPPErr)
	}
	_invokeSpeechErrorUPP(arg0, arg1, arg2, arg3, arg4)
	return nil
}

// InvokeSpeechErrorUPP invokes your error callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552214-invokespeecherrorupp
func InvokeSpeechErrorUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 int16, arg3 int, arg4 SpeechErrorUPP) {
	if callErr := tryInvokeSpeechErrorUPP(arg0, arg1, arg2, arg3, arg4); callErr != nil {
		panic(callErr)
	}
}

var _invokeSpeechPhonemeUPP func(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 int16, arg3 SpeechPhonemeUPP)
var _invokeSpeechPhonemeUPPErr error

func tryInvokeSpeechPhonemeUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 int16, arg3 SpeechPhonemeUPP) error {
	if _invokeSpeechPhonemeUPP == nil {
		return symbolCallError("InvokeSpeechPhonemeUPP", "10.0", _invokeSpeechPhonemeUPPErr)
	}
	_invokeSpeechPhonemeUPP(arg0, arg1, arg2, arg3)
	return nil
}

// InvokeSpeechPhonemeUPP invokes your phoneme callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552234-invokespeechphonemeupp
func InvokeSpeechPhonemeUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 int16, arg3 SpeechPhonemeUPP) {
	if callErr := tryInvokeSpeechPhonemeUPP(arg0, arg1, arg2, arg3); callErr != nil {
		panic(callErr)
	}
}

var _invokeSpeechSyncUPP func(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 uint32, arg3 SpeechSyncUPP)
var _invokeSpeechSyncUPPErr error

func tryInvokeSpeechSyncUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 uint32, arg3 SpeechSyncUPP) error {
	if _invokeSpeechSyncUPP == nil {
		return symbolCallError("InvokeSpeechSyncUPP", "10.0", _invokeSpeechSyncUPPErr)
	}
	_invokeSpeechSyncUPP(arg0, arg1, arg2, arg3)
	return nil
}

// InvokeSpeechSyncUPP invokes your synchronization callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552243-invokespeechsyncupp
func InvokeSpeechSyncUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 uint32, arg3 SpeechSyncUPP) {
	if callErr := tryInvokeSpeechSyncUPP(arg0, arg1, arg2, arg3); callErr != nil {
		panic(callErr)
	}
}

var _invokeSpeechTextDoneUPP func(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 unsafe.Pointer, arg3 uint, arg4 int32, arg5 SpeechTextDoneUPP)
var _invokeSpeechTextDoneUPPErr error

func tryInvokeSpeechTextDoneUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 unsafe.Pointer, arg3 uint, arg4 int32, arg5 SpeechTextDoneUPP) error {
	if _invokeSpeechTextDoneUPP == nil {
		return symbolCallError("InvokeSpeechTextDoneUPP", "10.0", _invokeSpeechTextDoneUPPErr)
	}
	_invokeSpeechTextDoneUPP(arg0, arg1, arg2, arg3, arg4, arg5)
	return nil
}

// InvokeSpeechTextDoneUPP invokes your text-done callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552249-invokespeechtextdoneupp
func InvokeSpeechTextDoneUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 unsafe.Pointer, arg3 uint, arg4 int32, arg5 SpeechTextDoneUPP) {
	if callErr := tryInvokeSpeechTextDoneUPP(arg0, arg1, arg2, arg3, arg4, arg5); callErr != nil {
		panic(callErr)
	}
}

var _invokeSpeechWordUPP func(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 uint, arg3 uint16, arg4 SpeechWordUPP)
var _invokeSpeechWordUPPErr error

func tryInvokeSpeechWordUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 uint, arg3 uint16, arg4 SpeechWordUPP) error {
	if _invokeSpeechWordUPP == nil {
		return symbolCallError("InvokeSpeechWordUPP", "10.0", _invokeSpeechWordUPPErr)
	}
	_invokeSpeechWordUPP(arg0, arg1, arg2, arg3, arg4)
	return nil
}

// InvokeSpeechWordUPP invokes your word callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552227-invokespeechwordupp
func InvokeSpeechWordUPP(arg0 *SpeechChannelRecord, arg1 uintptr, arg2 uint, arg3 uint16, arg4 SpeechWordUPP) {
	if callErr := tryInvokeSpeechWordUPP(arg0, arg1, arg2, arg3, arg4); callErr != nil {
		panic(callErr)
	}
}

var _isIconRefMaskEmpty func(iconRef unsafe.Pointer) bool
var _isIconRefMaskEmptyErr error

func tryIsIconRefMaskEmpty(iconRef unsafe.Pointer) (bool, error) {
	if _isIconRefMaskEmpty == nil {
		return false, symbolCallError("IsIconRefMaskEmpty", "10.0", _isIconRefMaskEmptyErr)
	}
	return _isIconRefMaskEmpty(iconRef), nil
}

// IsIconRefMaskEmpty.
//
// See: https://developer.apple.com/documentation/applicationservices/1464419-isiconrefmaskempty
func IsIconRefMaskEmpty(iconRef unsafe.Pointer) bool {
	result, callErr := tryIsIconRefMaskEmpty(iconRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _isProcessVisible func(arg0 *ProcessSerialNumber) bool
var _isProcessVisibleErr error

func tryIsProcessVisible(arg0 *ProcessSerialNumber) (bool, error) {
	if _isProcessVisible == nil {
		return false, symbolCallError("IsProcessVisible", "10.1", _isProcessVisibleErr)
	}
	return _isProcessVisible(arg0), nil
}

// IsProcessVisible.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501035-isprocessvisible
func IsProcessVisible(arg0 *ProcessSerialNumber) bool {
	result, callErr := tryIsProcessVisible(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _killProcess func(arg0 *ProcessSerialNumber) int16
var _killProcessErr error

func tryKillProcess(arg0 *ProcessSerialNumber) (int16, error) {
	if _killProcess == nil {
		return 0, symbolCallError("KillProcess", "10.2", _killProcessErr)
	}
	return _killProcess(arg0), nil
}

// KillProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501110-killprocess
func KillProcess(arg0 *ProcessSerialNumber) int16 {
	result, callErr := tryKillProcess(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _launchApplication func(arg0 LaunchPBPtr) int16
var _launchApplicationErr error

func tryLaunchApplication(arg0 LaunchPBPtr) (int16, error) {
	if _launchApplication == nil {
		return 0, symbolCallError("LaunchApplication", "10.0", _launchApplicationErr)
	}
	return _launchApplication(arg0), nil
}

// LaunchApplication.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501089-launchapplication
func LaunchApplication(arg0 LaunchPBPtr) int16 {
	result, callErr := tryLaunchApplication(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _makeVoiceSpec func(creator uint32, id uint32, voice unsafe.Pointer) int16
var _makeVoiceSpecErr error

func tryMakeVoiceSpec(creator uint32, id uint32, voice unsafe.Pointer) (int16, error) {
	if _makeVoiceSpec == nil {
		return 0, symbolCallError("MakeVoiceSpec", "10.0", _makeVoiceSpecErr)
	}
	return _makeVoiceSpec(creator, id, voice), nil
}

// MakeVoiceSpec sets the fields of a voice specification structure.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1461446-makevoicespec
func MakeVoiceSpec(creator uint32, id uint32, voice unsafe.Pointer) int16 {
	result, callErr := tryMakeVoiceSpec(creator, id, voice)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newIconActionUPP func(userRoutine unsafe.Pointer) IconActionUPP
var _newIconActionUPPErr error

func tryNewIconActionUPP(userRoutine IconActionProcPtr) (IconActionUPP, error) {
	if _newIconActionUPP == nil {
		return *new(IconActionUPP), symbolCallError("NewIconActionUPP", "10.0", _newIconActionUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 uint32, blockArg1 objc.ID, blockArg2 unsafe.Pointer) int16 {
		return userRoutine(blockArg0, objectivec.ObjectFromID(blockArg1), blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newIconActionUPP(_block0), nil
}

// NewIconActionUPP.
//
// See: https://developer.apple.com/documentation/applicationservices/1462738-newiconactionupp
func NewIconActionUPP(userRoutine IconActionProcPtr) IconActionUPP {
	result, callErr := tryNewIconActionUPP(userRoutine)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newIconGetterUPP func(userRoutine unsafe.Pointer) IconGetterUPP
var _newIconGetterUPPErr error

func tryNewIconGetterUPP(userRoutine IconGetterProcPtr) (IconGetterUPP, error) {
	if _newIconGetterUPP == nil {
		return *new(IconGetterUPP), symbolCallError("NewIconGetterUPP", "10.0", _newIconGetterUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 uint32, blockArg1 unsafe.Pointer) kernel.Ptr {
		return userRoutine(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newIconGetterUPP(_block0), nil
}

// NewIconGetterUPP.
//
// See: https://developer.apple.com/documentation/applicationservices/1458777-newicongetterupp
func NewIconGetterUPP(userRoutine IconGetterProcPtr) IconGetterUPP {
	result, callErr := tryNewIconGetterUPP(userRoutine)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechChannel func(voice unsafe.Pointer, chan_ **SpeechChannelRecord) int16
var _newSpeechChannelErr error

func tryNewSpeechChannel(voice unsafe.Pointer, chan_ **SpeechChannelRecord) (int16, error) {
	if _newSpeechChannel == nil {
		return 0, symbolCallError("NewSpeechChannel", "10.0", _newSpeechChannelErr)
	}
	return _newSpeechChannel(voice, chan_), nil
}

// NewSpeechChannel creates a new speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1461367-newspeechchannel
func NewSpeechChannel(voice unsafe.Pointer, chan_ **SpeechChannelRecord) int16 {
	result, callErr := tryNewSpeechChannel(voice, chan_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechDoneUPP func(arg0 unsafe.Pointer) SpeechDoneUPP
var _newSpeechDoneUPPErr error

func tryNewSpeechDoneUPP(arg0 SpeechDoneProcPtr) (SpeechDoneUPP, error) {
	if _newSpeechDoneUPP == nil {
		return *new(SpeechDoneUPP), symbolCallError("NewSpeechDoneUPP", "10.0", _newSpeechDoneUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *SpeechChannelRecord, blockArg1 uintptr) { arg0(blockArg0, blockArg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newSpeechDoneUPP(_block0), nil
}

// NewSpeechDoneUPP creates a new universal procedure pointer (UPP) to a speech-donecallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552218-newspeechdoneupp
func NewSpeechDoneUPP(arg0 SpeechDoneProcPtr) SpeechDoneUPP {
	result, callErr := tryNewSpeechDoneUPP(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechErrorUPP func(arg0 unsafe.Pointer) SpeechErrorUPP
var _newSpeechErrorUPPErr error

func tryNewSpeechErrorUPP(arg0 SpeechErrorProcPtr) (SpeechErrorUPP, error) {
	if _newSpeechErrorUPP == nil {
		return *new(SpeechErrorUPP), symbolCallError("NewSpeechErrorUPP", "10.0", _newSpeechErrorUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *SpeechChannelRecord, blockArg1 uintptr, blockArg2 int16, blockArg3 int) {
		arg0(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newSpeechErrorUPP(_block0), nil
}

// NewSpeechErrorUPP creates a new universal procedure pointer to an errorcallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552224-newspeecherrorupp
func NewSpeechErrorUPP(arg0 SpeechErrorProcPtr) SpeechErrorUPP {
	result, callErr := tryNewSpeechErrorUPP(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechPhonemeUPP func(arg0 unsafe.Pointer) SpeechPhonemeUPP
var _newSpeechPhonemeUPPErr error

func tryNewSpeechPhonemeUPP(arg0 SpeechPhonemeProcPtr) (SpeechPhonemeUPP, error) {
	if _newSpeechPhonemeUPP == nil {
		return *new(SpeechPhonemeUPP), symbolCallError("NewSpeechPhonemeUPP", "10.0", _newSpeechPhonemeUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *SpeechChannelRecord, blockArg1 uintptr, blockArg2 int16) {
		arg0(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newSpeechPhonemeUPP(_block0), nil
}

// NewSpeechPhonemeUPP disposes of a universal procedure pointer (UPP) to a phonemecallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552225-newspeechphonemeupp
func NewSpeechPhonemeUPP(arg0 SpeechPhonemeProcPtr) SpeechPhonemeUPP {
	result, callErr := tryNewSpeechPhonemeUPP(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechSyncUPP func(arg0 unsafe.Pointer) SpeechSyncUPP
var _newSpeechSyncUPPErr error

func tryNewSpeechSyncUPP(arg0 SpeechSyncProcPtr) (SpeechSyncUPP, error) {
	if _newSpeechSyncUPP == nil {
		return *new(SpeechSyncUPP), symbolCallError("NewSpeechSyncUPP", "10.0", _newSpeechSyncUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *SpeechChannelRecord, blockArg1 uintptr, blockArg2 uint32) {
		arg0(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newSpeechSyncUPP(_block0), nil
}

// NewSpeechSyncUPP creates a new universal procedure pointer (UPP) to a synchronizationcallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552244-newspeechsyncupp
func NewSpeechSyncUPP(arg0 SpeechSyncProcPtr) SpeechSyncUPP {
	result, callErr := tryNewSpeechSyncUPP(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechTextDoneUPP func(arg0 unsafe.Pointer) SpeechTextDoneUPP
var _newSpeechTextDoneUPPErr error

func tryNewSpeechTextDoneUPP(arg0 SpeechTextDoneProcPtr) (SpeechTextDoneUPP, error) {
	if _newSpeechTextDoneUPP == nil {
		return *new(SpeechTextDoneUPP), symbolCallError("NewSpeechTextDoneUPP", "10.0", _newSpeechTextDoneUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *SpeechChannelRecord, blockArg1 uintptr, blockArg2 objc.ID, blockArg3 objc.ID, blockArg4 objc.ID) {
		arg0(blockArg0, blockArg1, objectivec.ObjectFromID(blockArg2), objectivec.ObjectFromID(blockArg3), objectivec.ObjectFromID(blockArg4))
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newSpeechTextDoneUPP(_block0), nil
}

// NewSpeechTextDoneUPP creates a new universal procedure pointer (UPP) to a text-donecallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552247-newspeechtextdoneupp
func NewSpeechTextDoneUPP(arg0 SpeechTextDoneProcPtr) SpeechTextDoneUPP {
	result, callErr := tryNewSpeechTextDoneUPP(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _newSpeechWordUPP func(arg0 unsafe.Pointer) SpeechWordUPP
var _newSpeechWordUPPErr error

func tryNewSpeechWordUPP(arg0 SpeechWordProcPtr) (SpeechWordUPP, error) {
	if _newSpeechWordUPP == nil {
		return *new(SpeechWordUPP), symbolCallError("NewSpeechWordUPP", "10.0", _newSpeechWordUPPErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 *SpeechChannelRecord, blockArg1 uintptr, blockArg2 uint, blockArg3 uint16) {
		arg0(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _newSpeechWordUPP(_block0), nil
}

// NewSpeechWordUPP creates a new universal procedure pointer (UPP) to a wordcallback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552230-newspeechwordupp
func NewSpeechWordUPP(arg0 SpeechWordProcPtr) SpeechWordUPP {
	result, callErr := tryNewSpeechWordUPP(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCGImageCreateWithEPSDataProvider func(epsDataProvider coregraphics.CGDataProviderRef, epsPreview coregraphics.CGImageRef) coregraphics.CGImageRef
var _pMCGImageCreateWithEPSDataProviderErr error

func tryPMCGImageCreateWithEPSDataProvider(epsDataProvider coregraphics.CGDataProviderRef, epsPreview coregraphics.CGImageRef) (coregraphics.CGImageRef, error) {
	if _pMCGImageCreateWithEPSDataProvider == nil {
		return 0, symbolCallError("PMCGImageCreateWithEPSDataProvider", "10.1", _pMCGImageCreateWithEPSDataProviderErr)
	}
	return _pMCGImageCreateWithEPSDataProvider(epsDataProvider, epsPreview), nil
}

// PMCGImageCreateWithEPSDataProvider creates an image that references both the PostScript contents of EPS data and a preview (proxy) image for the data.
//
// See: https://developer.apple.com/documentation/applicationservices/1462361-pmcgimagecreatewithepsdataprovid
func PMCGImageCreateWithEPSDataProvider(epsDataProvider coregraphics.CGDataProviderRef, epsPreview coregraphics.CGImageRef) coregraphics.CGImageRef {
	result, callErr := tryPMCGImageCreateWithEPSDataProvider(epsDataProvider, epsPreview)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCopyAvailablePPDs func(domain PMPPDDomain, ppds *corefoundation.CFArrayRef) int32
var _pMCopyAvailablePPDsErr error

func tryPMCopyAvailablePPDs(domain PMPPDDomain, ppds *corefoundation.CFArrayRef) (int32, error) {
	if _pMCopyAvailablePPDs == nil {
		return 0, symbolCallError("PMCopyAvailablePPDs", "10.3", _pMCopyAvailablePPDsErr)
	}
	return _pMCopyAvailablePPDs(domain, ppds), nil
}

// PMCopyAvailablePPDs obtains the list of PostScript printer description (PPD) files in a PPD domain.
//
// See: https://developer.apple.com/documentation/applicationservices/1464170-pmcopyavailableppds
func PMCopyAvailablePPDs(domain PMPPDDomain, ppds *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMCopyAvailablePPDs(domain, ppds)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCopyLocalizedPPD func(ppd corefoundation.CFURLRef, localizedPPD *corefoundation.CFURLRef) int32
var _pMCopyLocalizedPPDErr error

func tryPMCopyLocalizedPPD(ppd corefoundation.CFURLRef, localizedPPD *corefoundation.CFURLRef) (int32, error) {
	if _pMCopyLocalizedPPD == nil {
		return 0, symbolCallError("PMCopyLocalizedPPD", "10.3", _pMCopyLocalizedPPDErr)
	}
	return _pMCopyLocalizedPPD(ppd, localizedPPD), nil
}

// PMCopyLocalizedPPD obtains a localized PostScript printer description (PPD) file.
//
// See: https://developer.apple.com/documentation/applicationservices/1459690-pmcopylocalizedppd
func PMCopyLocalizedPPD(ppd corefoundation.CFURLRef, localizedPPD *corefoundation.CFURLRef) int32 {
	result, callErr := tryPMCopyLocalizedPPD(ppd, localizedPPD)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCopyPPDData func(ppd corefoundation.CFURLRef, data *corefoundation.CFDataRef) int32
var _pMCopyPPDDataErr error

func tryPMCopyPPDData(ppd corefoundation.CFURLRef, data *corefoundation.CFDataRef) (int32, error) {
	if _pMCopyPPDData == nil {
		return 0, symbolCallError("PMCopyPPDData", "10.3", _pMCopyPPDDataErr)
	}
	return _pMCopyPPDData(ppd, data), nil
}

// PMCopyPPDData obtains the uncompressed PPD data for a PostScript printer description (PPD) file.
//
// See: https://developer.apple.com/documentation/applicationservices/1460345-pmcopyppddata
func PMCopyPPDData(ppd corefoundation.CFURLRef, data *corefoundation.CFDataRef) int32 {
	result, callErr := tryPMCopyPPDData(ppd, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCopyPageFormat func(formatSrc uintptr, formatDest uintptr) int32
var _pMCopyPageFormatErr error

func tryPMCopyPageFormat(formatSrc uintptr, formatDest uintptr) (int32, error) {
	if _pMCopyPageFormat == nil {
		return 0, symbolCallError("PMCopyPageFormat", "10.0", _pMCopyPageFormatErr)
	}
	return _pMCopyPageFormat(formatSrc, formatDest), nil
}

// PMCopyPageFormat copies the settings from one page format object into another.
//
// See: https://developer.apple.com/documentation/applicationservices/1464669-pmcopypageformat
func PMCopyPageFormat(formatSrc uintptr, formatDest uintptr) int32 {
	result, callErr := tryPMCopyPageFormat(formatSrc, formatDest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCopyPrintSettings func(settingSrc uintptr, settingDest uintptr) int32
var _pMCopyPrintSettingsErr error

func tryPMCopyPrintSettings(settingSrc uintptr, settingDest uintptr) (int32, error) {
	if _pMCopyPrintSettings == nil {
		return 0, symbolCallError("PMCopyPrintSettings", "10.0", _pMCopyPrintSettingsErr)
	}
	return _pMCopyPrintSettings(settingSrc, settingDest), nil
}

// PMCopyPrintSettings copies the settings from one print settings object into another.
//
// See: https://developer.apple.com/documentation/applicationservices/1462491-pmcopyprintsettings
func PMCopyPrintSettings(settingSrc uintptr, settingDest uintptr) int32 {
	result, callErr := tryPMCopyPrintSettings(settingSrc, settingDest)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCreateGenericPrinter func(printer *uintptr) int32
var _pMCreateGenericPrinterErr error

func tryPMCreateGenericPrinter(printer *uintptr) (int32, error) {
	if _pMCreateGenericPrinter == nil {
		return 0, symbolCallError("PMCreateGenericPrinter", "10.5", _pMCreateGenericPrinterErr)
	}
	return _pMCreateGenericPrinter(printer), nil
}

// PMCreateGenericPrinter creates a generic printer object.
//
// See: https://developer.apple.com/documentation/applicationservices/1461960-pmcreategenericprinter
func PMCreateGenericPrinter(printer *uintptr) int32 {
	result, callErr := tryPMCreateGenericPrinter(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCreatePageFormat func(pageFormat *uintptr) int32
var _pMCreatePageFormatErr error

func tryPMCreatePageFormat(pageFormat *uintptr) (int32, error) {
	if _pMCreatePageFormat == nil {
		return 0, symbolCallError("PMCreatePageFormat", "10.0", _pMCreatePageFormatErr)
	}
	return _pMCreatePageFormat(pageFormat), nil
}

// PMCreatePageFormat creates a new page format object.
//
// See: https://developer.apple.com/documentation/applicationservices/1459485-pmcreatepageformat
func PMCreatePageFormat(pageFormat *uintptr) int32 {
	result, callErr := tryPMCreatePageFormat(pageFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCreatePageFormatWithPMPaper func(pageFormat *uintptr, paper uintptr) int32
var _pMCreatePageFormatWithPMPaperErr error

func tryPMCreatePageFormatWithPMPaper(pageFormat *uintptr, paper uintptr) (int32, error) {
	if _pMCreatePageFormatWithPMPaper == nil {
		return 0, symbolCallError("PMCreatePageFormatWithPMPaper", "10.3", _pMCreatePageFormatWithPMPaperErr)
	}
	return _pMCreatePageFormatWithPMPaper(pageFormat, paper), nil
}

// PMCreatePageFormatWithPMPaper creates a page format object with a specified paper.
//
// See: https://developer.apple.com/documentation/applicationservices/1459274-pmcreatepageformatwithpmpaper
func PMCreatePageFormatWithPMPaper(pageFormat *uintptr, paper uintptr) int32 {
	result, callErr := tryPMCreatePageFormatWithPMPaper(pageFormat, paper)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCreatePrintSettings func(printSettings *uintptr) int32
var _pMCreatePrintSettingsErr error

func tryPMCreatePrintSettings(printSettings *uintptr) (int32, error) {
	if _pMCreatePrintSettings == nil {
		return 0, symbolCallError("PMCreatePrintSettings", "10.0", _pMCreatePrintSettingsErr)
	}
	return _pMCreatePrintSettings(printSettings), nil
}

// PMCreatePrintSettings creates a new print settings object.
//
// See: https://developer.apple.com/documentation/applicationservices/1463239-pmcreateprintsettings
func PMCreatePrintSettings(printSettings *uintptr) int32 {
	result, callErr := tryPMCreatePrintSettings(printSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMCreateSession func(printSession *PMPrintSession) int32
var _pMCreateSessionErr error

func tryPMCreateSession(printSession *PMPrintSession) (int32, error) {
	if _pMCreateSession == nil {
		return 0, symbolCallError("PMCreateSession", "10.0", _pMCreateSessionErr)
	}
	return _pMCreateSession(printSession), nil
}

// PMCreateSession creates and initializes a printing session object and creates a context for printing operations.
//
// See: https://developer.apple.com/documentation/applicationservices/1463247-pmcreatesession
func PMCreateSession(printSession *PMPrintSession) int32 {
	result, callErr := tryPMCreateSession(printSession)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetAdjustedPageRect func(pageFormat uintptr, pageRect unsafe.Pointer) int32
var _pMGetAdjustedPageRectErr error

func tryPMGetAdjustedPageRect(pageFormat uintptr, pageRect unsafe.Pointer) (int32, error) {
	if _pMGetAdjustedPageRect == nil {
		return 0, symbolCallError("PMGetAdjustedPageRect", "10.0", _pMGetAdjustedPageRectErr)
	}
	return _pMGetAdjustedPageRect(pageFormat, pageRect), nil
}

// PMGetAdjustedPageRect obtains the imageable area or page rectangle, taking into account orientation, application drawing resolution, and scaling settings.
//
// See: https://developer.apple.com/documentation/applicationservices/1461543-pmgetadjustedpagerect
func PMGetAdjustedPageRect(pageFormat uintptr, pageRect unsafe.Pointer) int32 {
	result, callErr := tryPMGetAdjustedPageRect(pageFormat, pageRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetAdjustedPaperRect func(pageFormat uintptr, paperRect unsafe.Pointer) int32
var _pMGetAdjustedPaperRectErr error

func tryPMGetAdjustedPaperRect(pageFormat uintptr, paperRect unsafe.Pointer) (int32, error) {
	if _pMGetAdjustedPaperRect == nil {
		return 0, symbolCallError("PMGetAdjustedPaperRect", "10.0", _pMGetAdjustedPaperRectErr)
	}
	return _pMGetAdjustedPaperRect(pageFormat, paperRect), nil
}

// PMGetAdjustedPaperRect obtains the rectangle defining the paper size, taking into account orientation, application drawing resolution, and scaling settings.
//
// See: https://developer.apple.com/documentation/applicationservices/1459167-pmgetadjustedpaperrect
func PMGetAdjustedPaperRect(pageFormat uintptr, paperRect unsafe.Pointer) int32 {
	result, callErr := tryPMGetAdjustedPaperRect(pageFormat, paperRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetCollate func(printSettings uintptr, collate unsafe.Pointer) int32
var _pMGetCollateErr error

func tryPMGetCollate(printSettings uintptr, collate unsafe.Pointer) (int32, error) {
	if _pMGetCollate == nil {
		return 0, symbolCallError("PMGetCollate", "10.2", _pMGetCollateErr)
	}
	return _pMGetCollate(printSettings, collate), nil
}

// PMGetCollate obtains a Boolean value that indicates whether the job collate option is selected.
//
// See: https://developer.apple.com/documentation/applicationservices/1464492-pmgetcollate
func PMGetCollate(printSettings uintptr, collate unsafe.Pointer) int32 {
	result, callErr := tryPMGetCollate(printSettings, collate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetCopies func(printSettings uintptr, copies *uint32) int32
var _pMGetCopiesErr error

func tryPMGetCopies(printSettings uintptr, copies *uint32) (int32, error) {
	if _pMGetCopies == nil {
		return 0, symbolCallError("PMGetCopies", "10.0", _pMGetCopiesErr)
	}
	return _pMGetCopies(printSettings, copies), nil
}

// PMGetCopies obtains the number of copies that the user requests to be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1464480-pmgetcopies
func PMGetCopies(printSettings uintptr, copies *uint32) int32 {
	result, callErr := tryPMGetCopies(printSettings, copies)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetDuplex func(printSettings uintptr, duplexSetting *PMDuplexMode) int32
var _pMGetDuplexErr error

func tryPMGetDuplex(printSettings uintptr, duplexSetting *PMDuplexMode) (int32, error) {
	if _pMGetDuplex == nil {
		return 0, symbolCallError("PMGetDuplex", "10.4", _pMGetDuplexErr)
	}
	return _pMGetDuplex(printSettings, duplexSetting), nil
}

// PMGetDuplex obtains the selected duplex mode.
//
// See: https://developer.apple.com/documentation/applicationservices/1458921-pmgetduplex
func PMGetDuplex(printSettings uintptr, duplexSetting *PMDuplexMode) int32 {
	result, callErr := tryPMGetDuplex(printSettings, duplexSetting)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetFirstPage func(printSettings uintptr, first *uint32) int32
var _pMGetFirstPageErr error

func tryPMGetFirstPage(printSettings uintptr, first *uint32) (int32, error) {
	if _pMGetFirstPage == nil {
		return 0, symbolCallError("PMGetFirstPage", "10.0", _pMGetFirstPageErr)
	}
	return _pMGetFirstPage(printSettings, first), nil
}

// PMGetFirstPage obtains the number of the first page to be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1460271-pmgetfirstpage
func PMGetFirstPage(printSettings uintptr, first *uint32) int32 {
	result, callErr := tryPMGetFirstPage(printSettings, first)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetLastPage func(printSettings uintptr, last *uint32) int32
var _pMGetLastPageErr error

func tryPMGetLastPage(printSettings uintptr, last *uint32) (int32, error) {
	if _pMGetLastPage == nil {
		return 0, symbolCallError("PMGetLastPage", "10.0", _pMGetLastPageErr)
	}
	return _pMGetLastPage(printSettings, last), nil
}

// PMGetLastPage obtains the number of the last page to be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1462747-pmgetlastpage
func PMGetLastPage(printSettings uintptr, last *uint32) int32 {
	result, callErr := tryPMGetLastPage(printSettings, last)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetOrientation func(pageFormat uintptr, orientation *PMOrientation) int32
var _pMGetOrientationErr error

func tryPMGetOrientation(pageFormat uintptr, orientation *PMOrientation) (int32, error) {
	if _pMGetOrientation == nil {
		return 0, symbolCallError("PMGetOrientation", "10.0", _pMGetOrientationErr)
	}
	return _pMGetOrientation(pageFormat, orientation), nil
}

// PMGetOrientation obtains the current setting for page orientation.
//
// See: https://developer.apple.com/documentation/applicationservices/1459144-pmgetorientation
func PMGetOrientation(pageFormat uintptr, orientation *PMOrientation) int32 {
	result, callErr := tryPMGetOrientation(pageFormat, orientation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetPageFormatExtendedData func(pageFormat uintptr, dataID uint32, size *uint32, extendedData unsafe.Pointer) int32
var _pMGetPageFormatExtendedDataErr error

func tryPMGetPageFormatExtendedData(pageFormat uintptr, dataID uint32, size *uint32, extendedData unsafe.Pointer) (int32, error) {
	if _pMGetPageFormatExtendedData == nil {
		return 0, symbolCallError("PMGetPageFormatExtendedData", "10.0", _pMGetPageFormatExtendedDataErr)
	}
	return _pMGetPageFormatExtendedData(pageFormat, dataID, size, extendedData), nil
}

// PMGetPageFormatExtendedData obtains extended page format data previously stored by your application.
//
// See: https://developer.apple.com/documentation/applicationservices/1464455-pmgetpageformatextendeddata
func PMGetPageFormatExtendedData(pageFormat uintptr, dataID uint32, size *uint32, extendedData unsafe.Pointer) int32 {
	result, callErr := tryPMGetPageFormatExtendedData(pageFormat, dataID, size, extendedData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetPageFormatPaper func(format uintptr, paper *uintptr) int32
var _pMGetPageFormatPaperErr error

func tryPMGetPageFormatPaper(format uintptr, paper *uintptr) (int32, error) {
	if _pMGetPageFormatPaper == nil {
		return 0, symbolCallError("PMGetPageFormatPaper", "10.3", _pMGetPageFormatPaperErr)
	}
	return _pMGetPageFormatPaper(format, paper), nil
}

// PMGetPageFormatPaper obtains the paper associated with a page format object.
//
// See: https://developer.apple.com/documentation/applicationservices/1461319-pmgetpageformatpaper
func PMGetPageFormatPaper(format uintptr, paper *uintptr) int32 {
	result, callErr := tryPMGetPageFormatPaper(format, paper)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetPageRange func(printSettings uintptr, minPage *uint32, maxPage *uint32) int32
var _pMGetPageRangeErr error

func tryPMGetPageRange(printSettings uintptr, minPage *uint32, maxPage *uint32) (int32, error) {
	if _pMGetPageRange == nil {
		return 0, symbolCallError("PMGetPageRange", "10.0", _pMGetPageRangeErr)
	}
	return _pMGetPageRange(printSettings, minPage, maxPage), nil
}

// PMGetPageRange obtains the valid range of pages that can be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1459324-pmgetpagerange
func PMGetPageRange(printSettings uintptr, minPage *uint32, maxPage *uint32) int32 {
	result, callErr := tryPMGetPageRange(printSettings, minPage, maxPage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetScale func(pageFormat uintptr, scale unsafe.Pointer) int32
var _pMGetScaleErr error

func tryPMGetScale(pageFormat uintptr, scale unsafe.Pointer) (int32, error) {
	if _pMGetScale == nil {
		return 0, symbolCallError("PMGetScale", "10.0", _pMGetScaleErr)
	}
	return _pMGetScale(pageFormat, scale), nil
}

// PMGetScale obtains the scaling factor currently applied to the page and paper rectangles.
//
// See: https://developer.apple.com/documentation/applicationservices/1458796-pmgetscale
func PMGetScale(pageFormat uintptr, scale unsafe.Pointer) int32 {
	result, callErr := tryPMGetScale(pageFormat, scale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetUnadjustedPageRect func(pageFormat uintptr, pageRect unsafe.Pointer) int32
var _pMGetUnadjustedPageRectErr error

func tryPMGetUnadjustedPageRect(pageFormat uintptr, pageRect unsafe.Pointer) (int32, error) {
	if _pMGetUnadjustedPageRect == nil {
		return 0, symbolCallError("PMGetUnadjustedPageRect", "10.0", _pMGetUnadjustedPageRectErr)
	}
	return _pMGetUnadjustedPageRect(pageFormat, pageRect), nil
}

// PMGetUnadjustedPageRect obtains the imageable area or page rectangle, unaffected by orientation, resolution, or scaling.
//
// See: https://developer.apple.com/documentation/applicationservices/1462944-pmgetunadjustedpagerect
func PMGetUnadjustedPageRect(pageFormat uintptr, pageRect unsafe.Pointer) int32 {
	result, callErr := tryPMGetUnadjustedPageRect(pageFormat, pageRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMGetUnadjustedPaperRect func(pageFormat uintptr, paperRect unsafe.Pointer) int32
var _pMGetUnadjustedPaperRectErr error

func tryPMGetUnadjustedPaperRect(pageFormat uintptr, paperRect unsafe.Pointer) (int32, error) {
	if _pMGetUnadjustedPaperRect == nil {
		return 0, symbolCallError("PMGetUnadjustedPaperRect", "10.0", _pMGetUnadjustedPaperRectErr)
	}
	return _pMGetUnadjustedPaperRect(pageFormat, paperRect), nil
}

// PMGetUnadjustedPaperRect obtains the paper rectangle, unaffected by rotation, resolution, or scaling.
//
// See: https://developer.apple.com/documentation/applicationservices/1462939-pmgetunadjustedpaperrect
func PMGetUnadjustedPaperRect(pageFormat uintptr, paperRect unsafe.Pointer) int32 {
	result, callErr := tryPMGetUnadjustedPaperRect(pageFormat, paperRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPageFormatCreateDataRepresentation func(pageFormat uintptr, data *corefoundation.CFDataRef, format PMDataFormat) int32
var _pMPageFormatCreateDataRepresentationErr error

func tryPMPageFormatCreateDataRepresentation(pageFormat uintptr, data *corefoundation.CFDataRef, format PMDataFormat) (int32, error) {
	if _pMPageFormatCreateDataRepresentation == nil {
		return 0, symbolCallError("PMPageFormatCreateDataRepresentation", "10.5", _pMPageFormatCreateDataRepresentationErr)
	}
	return _pMPageFormatCreateDataRepresentation(pageFormat, data, format), nil
}

// PMPageFormatCreateDataRepresentation creates a data representation of a page format object.
//
// See: https://developer.apple.com/documentation/applicationservices/1464227-pmpageformatcreatedatarepresenta
func PMPageFormatCreateDataRepresentation(pageFormat uintptr, data *corefoundation.CFDataRef, format PMDataFormat) int32 {
	result, callErr := tryPMPageFormatCreateDataRepresentation(pageFormat, data, format)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPageFormatCreateWithDataRepresentation func(data corefoundation.CFDataRef, pageFormat *uintptr) int32
var _pMPageFormatCreateWithDataRepresentationErr error

func tryPMPageFormatCreateWithDataRepresentation(data corefoundation.CFDataRef, pageFormat *uintptr) (int32, error) {
	if _pMPageFormatCreateWithDataRepresentation == nil {
		return 0, symbolCallError("PMPageFormatCreateWithDataRepresentation", "10.5", _pMPageFormatCreateWithDataRepresentationErr)
	}
	return _pMPageFormatCreateWithDataRepresentation(data, pageFormat), nil
}

// PMPageFormatCreateWithDataRepresentation creates a page format object from a data representation.
//
// See: https://developer.apple.com/documentation/applicationservices/1462876-pmpageformatcreatewithdatarepres
func PMPageFormatCreateWithDataRepresentation(data corefoundation.CFDataRef, pageFormat *uintptr) int32 {
	result, callErr := tryPMPageFormatCreateWithDataRepresentation(data, pageFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPageFormatGetPrinterID func(pageFormat uintptr, printerID *corefoundation.CFStringRef) int32
var _pMPageFormatGetPrinterIDErr error

func tryPMPageFormatGetPrinterID(pageFormat uintptr, printerID *corefoundation.CFStringRef) (int32, error) {
	if _pMPageFormatGetPrinterID == nil {
		return 0, symbolCallError("PMPageFormatGetPrinterID", "10.5", _pMPageFormatGetPrinterIDErr)
	}
	return _pMPageFormatGetPrinterID(pageFormat, printerID), nil
}

// PMPageFormatGetPrinterID obtains the identifier of the formatting printer for a page format object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462961-pmpageformatgetprinterid
func PMPageFormatGetPrinterID(pageFormat uintptr, printerID *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPageFormatGetPrinterID(pageFormat, printerID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperCreateCustom func(printer uintptr, id corefoundation.CFStringRef, name corefoundation.CFStringRef, width unsafe.Pointer, height unsafe.Pointer, margins *PMPaperMargins, paperP *uintptr) int32
var _pMPaperCreateCustomErr error

func tryPMPaperCreateCustom(printer uintptr, id corefoundation.CFStringRef, name corefoundation.CFStringRef, width unsafe.Pointer, height unsafe.Pointer, margins *PMPaperMargins, paperP *uintptr) (int32, error) {
	if _pMPaperCreateCustom == nil {
		return 0, symbolCallError("PMPaperCreateCustom", "10.5", _pMPaperCreateCustomErr)
	}
	return _pMPaperCreateCustom(printer, id, name, width, height, margins, paperP), nil
}

// PMPaperCreateCustom creates a custom paper object.
//
// See: https://developer.apple.com/documentation/applicationservices/1459322-pmpapercreatecustom
func PMPaperCreateCustom(printer uintptr, id corefoundation.CFStringRef, name corefoundation.CFStringRef, width unsafe.Pointer, height unsafe.Pointer, margins *PMPaperMargins, paperP *uintptr) int32 {
	result, callErr := tryPMPaperCreateCustom(printer, id, name, width, height, margins, paperP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperCreateLocalizedName func(paper uintptr, printer uintptr, paperName *corefoundation.CFStringRef) int32
var _pMPaperCreateLocalizedNameErr error

func tryPMPaperCreateLocalizedName(paper uintptr, printer uintptr, paperName *corefoundation.CFStringRef) (int32, error) {
	if _pMPaperCreateLocalizedName == nil {
		return 0, symbolCallError("PMPaperCreateLocalizedName", "10.5", _pMPaperCreateLocalizedNameErr)
	}
	return _pMPaperCreateLocalizedName(paper, printer, paperName), nil
}

// PMPaperCreateLocalizedName obtains the localized name for a given paper.
//
// See: https://developer.apple.com/documentation/applicationservices/1460981-pmpapercreatelocalizedname
func PMPaperCreateLocalizedName(paper uintptr, printer uintptr, paperName *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPaperCreateLocalizedName(paper, printer, paperName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperGetHeight func(paper uintptr, paperHeight unsafe.Pointer) int32
var _pMPaperGetHeightErr error

func tryPMPaperGetHeight(paper uintptr, paperHeight unsafe.Pointer) (int32, error) {
	if _pMPaperGetHeight == nil {
		return 0, symbolCallError("PMPaperGetHeight", "10.3", _pMPaperGetHeightErr)
	}
	return _pMPaperGetHeight(paper, paperHeight), nil
}

// PMPaperGetHeight obtains the height of the sheet of paper represented by a paper object.
//
// See: https://developer.apple.com/documentation/applicationservices/1460389-pmpapergetheight
func PMPaperGetHeight(paper uintptr, paperHeight unsafe.Pointer) int32 {
	result, callErr := tryPMPaperGetHeight(paper, paperHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperGetID func(paper uintptr, paperID *corefoundation.CFStringRef) int32
var _pMPaperGetIDErr error

func tryPMPaperGetID(paper uintptr, paperID *corefoundation.CFStringRef) (int32, error) {
	if _pMPaperGetID == nil {
		return 0, symbolCallError("PMPaperGetID", "10.3", _pMPaperGetIDErr)
	}
	return _pMPaperGetID(paper, paperID), nil
}

// PMPaperGetID obtains the identifier of a paper object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462910-pmpapergetid
func PMPaperGetID(paper uintptr, paperID *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPaperGetID(paper, paperID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperGetMargins func(paper uintptr, paperMargins *PMPaperMargins) int32
var _pMPaperGetMarginsErr error

func tryPMPaperGetMargins(paper uintptr, paperMargins *PMPaperMargins) (int32, error) {
	if _pMPaperGetMargins == nil {
		return 0, symbolCallError("PMPaperGetMargins", "10.3", _pMPaperGetMarginsErr)
	}
	return _pMPaperGetMargins(paper, paperMargins), nil
}

// PMPaperGetMargins obtains the margins describing the unprintable area of the sheet represented by a paper object.
//
// See: https://developer.apple.com/documentation/applicationservices/1461994-pmpapergetmargins
func PMPaperGetMargins(paper uintptr, paperMargins *PMPaperMargins) int32 {
	result, callErr := tryPMPaperGetMargins(paper, paperMargins)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperGetPPDPaperName func(paper uintptr, paperName *corefoundation.CFStringRef) int32
var _pMPaperGetPPDPaperNameErr error

func tryPMPaperGetPPDPaperName(paper uintptr, paperName *corefoundation.CFStringRef) (int32, error) {
	if _pMPaperGetPPDPaperName == nil {
		return 0, symbolCallError("PMPaperGetPPDPaperName", "10.5", _pMPaperGetPPDPaperNameErr)
	}
	return _pMPaperGetPPDPaperName(paper, paperName), nil
}

// PMPaperGetPPDPaperName obtains the PPD paper name for a given paper.
//
// See: https://developer.apple.com/documentation/applicationservices/1461039-pmpapergetppdpapername
func PMPaperGetPPDPaperName(paper uintptr, paperName *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPaperGetPPDPaperName(paper, paperName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperGetPrinterID func(paper uintptr, printerID *corefoundation.CFStringRef) int32
var _pMPaperGetPrinterIDErr error

func tryPMPaperGetPrinterID(paper uintptr, printerID *corefoundation.CFStringRef) (int32, error) {
	if _pMPaperGetPrinterID == nil {
		return 0, symbolCallError("PMPaperGetPrinterID", "10.5", _pMPaperGetPrinterIDErr)
	}
	return _pMPaperGetPrinterID(paper, printerID), nil
}

// PMPaperGetPrinterID obtains the printer ID of the printer to which a given paper corresponds.
//
// See: https://developer.apple.com/documentation/applicationservices/1461737-pmpapergetprinterid
func PMPaperGetPrinterID(paper uintptr, printerID *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPaperGetPrinterID(paper, printerID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperGetWidth func(paper uintptr, paperWidth unsafe.Pointer) int32
var _pMPaperGetWidthErr error

func tryPMPaperGetWidth(paper uintptr, paperWidth unsafe.Pointer) (int32, error) {
	if _pMPaperGetWidth == nil {
		return 0, symbolCallError("PMPaperGetWidth", "10.3", _pMPaperGetWidthErr)
	}
	return _pMPaperGetWidth(paper, paperWidth), nil
}

// PMPaperGetWidth obtains the width of the sheet of paper represented by a paper object.
//
// See: https://developer.apple.com/documentation/applicationservices/1459209-pmpapergetwidth
func PMPaperGetWidth(paper uintptr, paperWidth unsafe.Pointer) int32 {
	result, callErr := tryPMPaperGetWidth(paper, paperWidth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPaperIsCustom func(paper uintptr) bool
var _pMPaperIsCustomErr error

func tryPMPaperIsCustom(paper uintptr) (bool, error) {
	if _pMPaperIsCustom == nil {
		return false, symbolCallError("PMPaperIsCustom", "10.5", _pMPaperIsCustomErr)
	}
	return _pMPaperIsCustom(paper), nil
}

// PMPaperIsCustom returns a Boolean value indicating whether a specified paper is a custom paper.
//
// See: https://developer.apple.com/documentation/applicationservices/1459526-pmpaperiscustom
func PMPaperIsCustom(paper uintptr) bool {
	result, callErr := tryPMPaperIsCustom(paper)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPresetCopyName func(preset PMPreset, name *corefoundation.CFStringRef) int32
var _pMPresetCopyNameErr error

func tryPMPresetCopyName(preset PMPreset, name *corefoundation.CFStringRef) (int32, error) {
	if _pMPresetCopyName == nil {
		return 0, symbolCallError("PMPresetCopyName", "10.3", _pMPresetCopyNameErr)
	}
	return _pMPresetCopyName(preset, name), nil
}

// PMPresetCopyName obtains the localized name for a preset.
//
// See: https://developer.apple.com/documentation/applicationservices/1460343-pmpresetcopyname
func PMPresetCopyName(preset PMPreset, name *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPresetCopyName(preset, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPresetCreatePrintSettings func(preset PMPreset, session PMPrintSession, printSettings *uintptr) int32
var _pMPresetCreatePrintSettingsErr error

func tryPMPresetCreatePrintSettings(preset PMPreset, session PMPrintSession, printSettings *uintptr) (int32, error) {
	if _pMPresetCreatePrintSettings == nil {
		return 0, symbolCallError("PMPresetCreatePrintSettings", "10.3", _pMPresetCreatePrintSettingsErr)
	}
	return _pMPresetCreatePrintSettings(preset, session, printSettings), nil
}

// PMPresetCreatePrintSettings creates a print settings object with settings that correspond to a preset.
//
// See: https://developer.apple.com/documentation/applicationservices/1463414-pmpresetcreateprintsettings
func PMPresetCreatePrintSettings(preset PMPreset, session PMPrintSession, printSettings *uintptr) int32 {
	result, callErr := tryPMPresetCreatePrintSettings(preset, session, printSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPresetGetAttributes func(preset PMPreset, attributes *corefoundation.CFDictionaryRef) int32
var _pMPresetGetAttributesErr error

func tryPMPresetGetAttributes(preset PMPreset, attributes *corefoundation.CFDictionaryRef) (int32, error) {
	if _pMPresetGetAttributes == nil {
		return 0, symbolCallError("PMPresetGetAttributes", "10.3", _pMPresetGetAttributesErr)
	}
	return _pMPresetGetAttributes(preset, attributes), nil
}

// PMPresetGetAttributes obtains the attributes of a preset.
//
// See: https://developer.apple.com/documentation/applicationservices/1459042-pmpresetgetattributes
func PMPresetGetAttributes(preset PMPreset, attributes *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryPMPresetGetAttributes(preset, attributes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsCopyAsDictionary func(printSettings uintptr, settingsDictionary *corefoundation.CFDictionaryRef) int32
var _pMPrintSettingsCopyAsDictionaryErr error

func tryPMPrintSettingsCopyAsDictionary(printSettings uintptr, settingsDictionary *corefoundation.CFDictionaryRef) (int32, error) {
	if _pMPrintSettingsCopyAsDictionary == nil {
		return 0, symbolCallError("PMPrintSettingsCopyAsDictionary", "10.5", _pMPrintSettingsCopyAsDictionaryErr)
	}
	return _pMPrintSettingsCopyAsDictionary(printSettings, settingsDictionary), nil
}

// PMPrintSettingsCopyAsDictionary creates a dictionary that contains the settings in a print settings object.
//
// See: https://developer.apple.com/documentation/applicationservices/1459088-pmprintsettingscopyasdictionary
func PMPrintSettingsCopyAsDictionary(printSettings uintptr, settingsDictionary *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryPMPrintSettingsCopyAsDictionary(printSettings, settingsDictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsCopyKeys func(printSettings uintptr, settingsKeys *corefoundation.CFArrayRef) int32
var _pMPrintSettingsCopyKeysErr error

func tryPMPrintSettingsCopyKeys(printSettings uintptr, settingsKeys *corefoundation.CFArrayRef) (int32, error) {
	if _pMPrintSettingsCopyKeys == nil {
		return 0, symbolCallError("PMPrintSettingsCopyKeys", "10.5", _pMPrintSettingsCopyKeysErr)
	}
	return _pMPrintSettingsCopyKeys(printSettings, settingsKeys), nil
}

// PMPrintSettingsCopyKeys obtains the keys for items in a print settings object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462730-pmprintsettingscopykeys
func PMPrintSettingsCopyKeys(printSettings uintptr, settingsKeys *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMPrintSettingsCopyKeys(printSettings, settingsKeys)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsCreateDataRepresentation func(printSettings uintptr, data *corefoundation.CFDataRef, format PMDataFormat) int32
var _pMPrintSettingsCreateDataRepresentationErr error

func tryPMPrintSettingsCreateDataRepresentation(printSettings uintptr, data *corefoundation.CFDataRef, format PMDataFormat) (int32, error) {
	if _pMPrintSettingsCreateDataRepresentation == nil {
		return 0, symbolCallError("PMPrintSettingsCreateDataRepresentation", "10.5", _pMPrintSettingsCreateDataRepresentationErr)
	}
	return _pMPrintSettingsCreateDataRepresentation(printSettings, data, format), nil
}

// PMPrintSettingsCreateDataRepresentation creates a data representation of a print settings object.
//
// See: https://developer.apple.com/documentation/applicationservices/1464570-pmprintsettingscreatedatareprese
func PMPrintSettingsCreateDataRepresentation(printSettings uintptr, data *corefoundation.CFDataRef, format PMDataFormat) int32 {
	result, callErr := tryPMPrintSettingsCreateDataRepresentation(printSettings, data, format)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsCreateWithDataRepresentation func(data corefoundation.CFDataRef, printSettings *uintptr) int32
var _pMPrintSettingsCreateWithDataRepresentationErr error

func tryPMPrintSettingsCreateWithDataRepresentation(data corefoundation.CFDataRef, printSettings *uintptr) (int32, error) {
	if _pMPrintSettingsCreateWithDataRepresentation == nil {
		return 0, symbolCallError("PMPrintSettingsCreateWithDataRepresentation", "10.5", _pMPrintSettingsCreateWithDataRepresentationErr)
	}
	return _pMPrintSettingsCreateWithDataRepresentation(data, printSettings), nil
}

// PMPrintSettingsCreateWithDataRepresentation creates a print settings object from a data representation.
//
// See: https://developer.apple.com/documentation/applicationservices/1462203-pmprintsettingscreatewithdatarep
func PMPrintSettingsCreateWithDataRepresentation(data corefoundation.CFDataRef, printSettings *uintptr) int32 {
	result, callErr := tryPMPrintSettingsCreateWithDataRepresentation(data, printSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsGetJobName func(printSettings uintptr, name *corefoundation.CFStringRef) int32
var _pMPrintSettingsGetJobNameErr error

func tryPMPrintSettingsGetJobName(printSettings uintptr, name *corefoundation.CFStringRef) (int32, error) {
	if _pMPrintSettingsGetJobName == nil {
		return 0, symbolCallError("PMPrintSettingsGetJobName", "10.4", _pMPrintSettingsGetJobNameErr)
	}
	return _pMPrintSettingsGetJobName(printSettings, name), nil
}

// PMPrintSettingsGetJobName obtains the name of a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/1459233-pmprintsettingsgetjobname
func PMPrintSettingsGetJobName(printSettings uintptr, name *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPrintSettingsGetJobName(printSettings, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsGetValue func(printSettings uintptr, key corefoundation.CFStringRef, value *corefoundation.CFTypeRef) int32
var _pMPrintSettingsGetValueErr error

func tryPMPrintSettingsGetValue(printSettings uintptr, key corefoundation.CFStringRef, value *corefoundation.CFTypeRef) (int32, error) {
	if _pMPrintSettingsGetValue == nil {
		return 0, symbolCallError("PMPrintSettingsGetValue", "10.4", _pMPrintSettingsGetValueErr)
	}
	return _pMPrintSettingsGetValue(printSettings, key, value), nil
}

// PMPrintSettingsGetValue obtains the value of a setting in a print settings object.
//
// See: https://developer.apple.com/documentation/applicationservices/1460602-pmprintsettingsgetvalue
func PMPrintSettingsGetValue(printSettings uintptr, key corefoundation.CFStringRef, value *corefoundation.CFTypeRef) int32 {
	result, callErr := tryPMPrintSettingsGetValue(printSettings, key, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsSetJobName func(printSettings uintptr, name corefoundation.CFStringRef) int32
var _pMPrintSettingsSetJobNameErr error

func tryPMPrintSettingsSetJobName(printSettings uintptr, name corefoundation.CFStringRef) (int32, error) {
	if _pMPrintSettingsSetJobName == nil {
		return 0, symbolCallError("PMPrintSettingsSetJobName", "10.4", _pMPrintSettingsSetJobNameErr)
	}
	return _pMPrintSettingsSetJobName(printSettings, name), nil
}

// PMPrintSettingsSetJobName specifies the name of a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/1460149-pmprintsettingssetjobname
func PMPrintSettingsSetJobName(printSettings uintptr, name corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPrintSettingsSetJobName(printSettings, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsSetValue func(printSettings uintptr, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, locked unsafe.Pointer) int32
var _pMPrintSettingsSetValueErr error

func tryPMPrintSettingsSetValue(printSettings uintptr, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, locked unsafe.Pointer) (int32, error) {
	if _pMPrintSettingsSetValue == nil {
		return 0, symbolCallError("PMPrintSettingsSetValue", "10.4", _pMPrintSettingsSetValueErr)
	}
	return _pMPrintSettingsSetValue(printSettings, key, value, locked), nil
}

// PMPrintSettingsSetValue stores the value of a setting in a print settings object.
//
// See: https://developer.apple.com/documentation/applicationservices/1461697-pmprintsettingssetvalue
func PMPrintSettingsSetValue(printSettings uintptr, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, locked unsafe.Pointer) int32 {
	result, callErr := tryPMPrintSettingsSetValue(printSettings, key, value, locked)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsToOptions func(settings uintptr, options string) int32
var _pMPrintSettingsToOptionsErr error

func tryPMPrintSettingsToOptions(settings uintptr, options string) (int32, error) {
	if _pMPrintSettingsToOptions == nil {
		return 0, symbolCallError("PMPrintSettingsToOptions", "10.3", _pMPrintSettingsToOptionsErr)
	}
	return _pMPrintSettingsToOptions(settings, options), nil
}

// PMPrintSettingsToOptions converts print settings into a CUPS options string.
//
// See: https://developer.apple.com/documentation/applicationservices/1459069-pmprintsettingstooptions
func PMPrintSettingsToOptions(settings uintptr, options string) int32 {
	result, callErr := tryPMPrintSettingsToOptions(settings, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrintSettingsToOptionsWithPrinterAndPageFormat func(settings uintptr, printer uintptr, pageFormat uintptr, options string) int32
var _pMPrintSettingsToOptionsWithPrinterAndPageFormatErr error

func tryPMPrintSettingsToOptionsWithPrinterAndPageFormat(settings uintptr, printer uintptr, pageFormat uintptr, options string) (int32, error) {
	if _pMPrintSettingsToOptionsWithPrinterAndPageFormat == nil {
		return 0, symbolCallError("PMPrintSettingsToOptionsWithPrinterAndPageFormat", "10.5", _pMPrintSettingsToOptionsWithPrinterAndPageFormatErr)
	}
	return _pMPrintSettingsToOptionsWithPrinterAndPageFormat(settings, printer, pageFormat, options), nil
}

// PMPrintSettingsToOptionsWithPrinterAndPageFormat converts print settings and page format data into a CUPS options string for a specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459435-pmprintsettingstooptionswithprin
func PMPrintSettingsToOptionsWithPrinterAndPageFormat(settings uintptr, printer uintptr, pageFormat uintptr, options string) int32 {
	result, callErr := tryPMPrintSettingsToOptionsWithPrinterAndPageFormat(settings, printer, pageFormat, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterCopyDescriptionURL func(printer uintptr, descriptionType corefoundation.CFStringRef, fileURL *corefoundation.CFURLRef) int32
var _pMPrinterCopyDescriptionURLErr error

func tryPMPrinterCopyDescriptionURL(printer uintptr, descriptionType corefoundation.CFStringRef, fileURL *corefoundation.CFURLRef) (int32, error) {
	if _pMPrinterCopyDescriptionURL == nil {
		return 0, symbolCallError("PMPrinterCopyDescriptionURL", "10.4", _pMPrinterCopyDescriptionURLErr)
	}
	return _pMPrinterCopyDescriptionURL(printer, descriptionType, fileURL), nil
}

// PMPrinterCopyDescriptionURL obtains the URL of the description file for a given printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459187-pmprintercopydescriptionurl
func PMPrinterCopyDescriptionURL(printer uintptr, descriptionType corefoundation.CFStringRef, fileURL *corefoundation.CFURLRef) int32 {
	result, callErr := tryPMPrinterCopyDescriptionURL(printer, descriptionType, fileURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterCopyDeviceURI func(printer uintptr, deviceURI *corefoundation.CFURLRef) int32
var _pMPrinterCopyDeviceURIErr error

func tryPMPrinterCopyDeviceURI(printer uintptr, deviceURI *corefoundation.CFURLRef) (int32, error) {
	if _pMPrinterCopyDeviceURI == nil {
		return 0, symbolCallError("PMPrinterCopyDeviceURI", "10.4", _pMPrinterCopyDeviceURIErr)
	}
	return _pMPrinterCopyDeviceURI(printer, deviceURI), nil
}

// PMPrinterCopyDeviceURI obtains the device URI of a given printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1460543-pmprintercopydeviceuri
func PMPrinterCopyDeviceURI(printer uintptr, deviceURI *corefoundation.CFURLRef) int32 {
	result, callErr := tryPMPrinterCopyDeviceURI(printer, deviceURI)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterCopyHostName func(printer uintptr, hostNameP *corefoundation.CFStringRef) int32
var _pMPrinterCopyHostNameErr error

func tryPMPrinterCopyHostName(printer uintptr, hostNameP *corefoundation.CFStringRef) (int32, error) {
	if _pMPrinterCopyHostName == nil {
		return 0, symbolCallError("PMPrinterCopyHostName", "10.3", _pMPrinterCopyHostNameErr)
	}
	return _pMPrinterCopyHostName(printer, hostNameP), nil
}

// PMPrinterCopyHostName obtains the name of the server hosting the print queue for a given printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1462076-pmprintercopyhostname
func PMPrinterCopyHostName(printer uintptr, hostNameP *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPrinterCopyHostName(printer, hostNameP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterCopyPresets func(printer uintptr, presetList *corefoundation.CFArrayRef) int32
var _pMPrinterCopyPresetsErr error

func tryPMPrinterCopyPresets(printer uintptr, presetList *corefoundation.CFArrayRef) (int32, error) {
	if _pMPrinterCopyPresets == nil {
		return 0, symbolCallError("PMPrinterCopyPresets", "10.3", _pMPrinterCopyPresetsErr)
	}
	return _pMPrinterCopyPresets(printer, presetList), nil
}

// PMPrinterCopyPresets obtains a list of print settings presets for a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459117-pmprintercopypresets
func PMPrinterCopyPresets(printer uintptr, presetList *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMPrinterCopyPresets(printer, presetList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterCopyState func(printer uintptr, stateDict *corefoundation.CFDictionaryRef) int32
var _pMPrinterCopyStateErr error

func tryPMPrinterCopyState(printer uintptr, stateDict *corefoundation.CFDictionaryRef) (int32, error) {
	if _pMPrinterCopyState == nil {
		return 0, symbolCallError("PMPrinterCopyState", "10.6", _pMPrinterCopyStateErr)
	}
	return _pMPrinterCopyState(printer, stateDict), nil
}

// PMPrinterCopyState.
//
// See: https://developer.apple.com/documentation/applicationservices/1460381-pmprintercopystate
func PMPrinterCopyState(printer uintptr, stateDict *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryPMPrinterCopyState(printer, stateDict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterCreateFromPrinterID func(printerID corefoundation.CFStringRef) uintptr
var _pMPrinterCreateFromPrinterIDErr error

func tryPMPrinterCreateFromPrinterID(printerID corefoundation.CFStringRef) (uintptr, error) {
	if _pMPrinterCreateFromPrinterID == nil {
		return 0, symbolCallError("PMPrinterCreateFromPrinterID", "10.4", _pMPrinterCreateFromPrinterIDErr)
	}
	return _pMPrinterCreateFromPrinterID(printerID), nil
}

// PMPrinterCreateFromPrinterID creates a printer object from a print queue identifier.
//
// See: https://developer.apple.com/documentation/applicationservices/1461363-pmprintercreatefromprinterid
func PMPrinterCreateFromPrinterID(printerID corefoundation.CFStringRef) uintptr {
	result, callErr := tryPMPrinterCreateFromPrinterID(printerID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetCommInfo func(printer uintptr, supportsControlCharRangeP unsafe.Pointer, supportsEightBitP unsafe.Pointer) int32
var _pMPrinterGetCommInfoErr error

func tryPMPrinterGetCommInfo(printer uintptr, supportsControlCharRangeP unsafe.Pointer, supportsEightBitP unsafe.Pointer) (int32, error) {
	if _pMPrinterGetCommInfo == nil {
		return 0, symbolCallError("PMPrinterGetCommInfo", "10.3", _pMPrinterGetCommInfoErr)
	}
	return _pMPrinterGetCommInfo(printer, supportsControlCharRangeP, supportsEightBitP), nil
}

// PMPrinterGetCommInfo obtains information about the communication channel for a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1461069-pmprintergetcomminfo
func PMPrinterGetCommInfo(printer uintptr, supportsControlCharRangeP unsafe.Pointer, supportsEightBitP unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterGetCommInfo(printer, supportsControlCharRangeP, supportsEightBitP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetDriverCreator func(printer uintptr, creator *uint32) int32
var _pMPrinterGetDriverCreatorErr error

func tryPMPrinterGetDriverCreator(printer uintptr, creator *uint32) (int32, error) {
	if _pMPrinterGetDriverCreator == nil {
		return 0, symbolCallError("PMPrinterGetDriverCreator", "10.0", _pMPrinterGetDriverCreatorErr)
	}
	return _pMPrinterGetDriverCreator(printer, creator), nil
}

// PMPrinterGetDriverCreator obtains the creator of the driver associated with the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459107-pmprintergetdrivercreator
func PMPrinterGetDriverCreator(printer uintptr, creator *uint32) int32 {
	result, callErr := tryPMPrinterGetDriverCreator(printer, creator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetDriverReleaseInfo func(arg0 uintptr, arg1 unsafe.Pointer) int32
var _pMPrinterGetDriverReleaseInfoErr error

func tryPMPrinterGetDriverReleaseInfo(arg0 uintptr, arg1 unsafe.Pointer) (int32, error) {
	if _pMPrinterGetDriverReleaseInfo == nil {
		return 0, symbolCallError("PMPrinterGetDriverReleaseInfo", "10.0", _pMPrinterGetDriverReleaseInfoErr)
	}
	return _pMPrinterGetDriverReleaseInfo(arg0, arg1), nil
}

// PMPrinterGetDriverReleaseInfo obtains version information for the driver associated with the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1464149-pmprintergetdriverreleaseinfo
func PMPrinterGetDriverReleaseInfo(arg0 uintptr, arg1 unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterGetDriverReleaseInfo(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetID func(printer uintptr) corefoundation.CFStringRef
var _pMPrinterGetIDErr error

func tryPMPrinterGetID(printer uintptr) (corefoundation.CFStringRef, error) {
	if _pMPrinterGetID == nil {
		return 0, symbolCallError("PMPrinterGetID", "10.2", _pMPrinterGetIDErr)
	}
	return _pMPrinterGetID(printer), nil
}

// PMPrinterGetID returns the unique identifier of a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459606-pmprintergetid
func PMPrinterGetID(printer uintptr) corefoundation.CFStringRef {
	result, callErr := tryPMPrinterGetID(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetIndexedPrinterResolution func(printer uintptr, index uint32, resolutionP unsafe.Pointer) int32
var _pMPrinterGetIndexedPrinterResolutionErr error

func tryPMPrinterGetIndexedPrinterResolution(printer uintptr, index uint32, resolutionP unsafe.Pointer) (int32, error) {
	if _pMPrinterGetIndexedPrinterResolution == nil {
		return 0, symbolCallError("PMPrinterGetIndexedPrinterResolution", "10.0", _pMPrinterGetIndexedPrinterResolutionErr)
	}
	return _pMPrinterGetIndexedPrinterResolution(printer, index, resolutionP), nil
}

// PMPrinterGetIndexedPrinterResolution obtains a resolution setting based on an index into the range of settings supported by the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1464490-pmprintergetindexedprinterresolu
func PMPrinterGetIndexedPrinterResolution(printer uintptr, index uint32, resolutionP unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterGetIndexedPrinterResolution(printer, index, resolutionP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetLanguageInfo func(arg0 uintptr, arg1 PMLanguageInfo) int32
var _pMPrinterGetLanguageInfoErr error

func tryPMPrinterGetLanguageInfo(arg0 uintptr, arg1 PMLanguageInfo) (int32, error) {
	if _pMPrinterGetLanguageInfo == nil {
		return 0, symbolCallError("PMPrinterGetLanguageInfo", "10.0", _pMPrinterGetLanguageInfoErr)
	}
	return _pMPrinterGetLanguageInfo(arg0, arg1), nil
}

// PMPrinterGetLanguageInfo obtains information about the imaging language for the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1458956-pmprintergetlanguageinfo
func PMPrinterGetLanguageInfo(arg0 uintptr, arg1 PMLanguageInfo) int32 {
	result, callErr := tryPMPrinterGetLanguageInfo(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetLocation func(printer uintptr) corefoundation.CFStringRef
var _pMPrinterGetLocationErr error

func tryPMPrinterGetLocation(printer uintptr) (corefoundation.CFStringRef, error) {
	if _pMPrinterGetLocation == nil {
		return 0, symbolCallError("PMPrinterGetLocation", "10.2", _pMPrinterGetLocationErr)
	}
	return _pMPrinterGetLocation(printer), nil
}

// PMPrinterGetLocation returns the location of a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1461467-pmprintergetlocation
func PMPrinterGetLocation(printer uintptr) corefoundation.CFStringRef {
	result, callErr := tryPMPrinterGetLocation(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetMakeAndModelName func(printer uintptr, makeAndModel *corefoundation.CFStringRef) int32
var _pMPrinterGetMakeAndModelNameErr error

func tryPMPrinterGetMakeAndModelName(printer uintptr, makeAndModel *corefoundation.CFStringRef) (int32, error) {
	if _pMPrinterGetMakeAndModelName == nil {
		return 0, symbolCallError("PMPrinterGetMakeAndModelName", "10.2", _pMPrinterGetMakeAndModelNameErr)
	}
	return _pMPrinterGetMakeAndModelName(printer, makeAndModel), nil
}

// PMPrinterGetMakeAndModelName obtains the manufacturer and model name of the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1463347-pmprintergetmakeandmodelname
func PMPrinterGetMakeAndModelName(printer uintptr, makeAndModel *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMPrinterGetMakeAndModelName(printer, makeAndModel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetMimeTypes func(printer uintptr, settings uintptr, mimeTypes *corefoundation.CFArrayRef) int32
var _pMPrinterGetMimeTypesErr error

func tryPMPrinterGetMimeTypes(printer uintptr, settings uintptr, mimeTypes *corefoundation.CFArrayRef) (int32, error) {
	if _pMPrinterGetMimeTypes == nil {
		return 0, symbolCallError("PMPrinterGetMimeTypes", "10.3", _pMPrinterGetMimeTypesErr)
	}
	return _pMPrinterGetMimeTypes(printer, settings, mimeTypes), nil
}

// PMPrinterGetMimeTypes obtains a list of MIME content types supported by a printer using the specified print settings.
//
// See: https://developer.apple.com/documentation/applicationservices/1460125-pmprintergetmimetypes
func PMPrinterGetMimeTypes(printer uintptr, settings uintptr, mimeTypes *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMPrinterGetMimeTypes(printer, settings, mimeTypes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetName func(printer uintptr) corefoundation.CFStringRef
var _pMPrinterGetNameErr error

func tryPMPrinterGetName(printer uintptr) (corefoundation.CFStringRef, error) {
	if _pMPrinterGetName == nil {
		return 0, symbolCallError("PMPrinterGetName", "10.2", _pMPrinterGetNameErr)
	}
	return _pMPrinterGetName(printer), nil
}

// PMPrinterGetName returns the human-readable name of a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459018-pmprintergetname
func PMPrinterGetName(printer uintptr) corefoundation.CFStringRef {
	result, callErr := tryPMPrinterGetName(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetOutputResolution func(printer uintptr, printSettings uintptr, resolutionP unsafe.Pointer) int32
var _pMPrinterGetOutputResolutionErr error

func tryPMPrinterGetOutputResolution(printer uintptr, printSettings uintptr, resolutionP unsafe.Pointer) (int32, error) {
	if _pMPrinterGetOutputResolution == nil {
		return 0, symbolCallError("PMPrinterGetOutputResolution", "10.5", _pMPrinterGetOutputResolutionErr)
	}
	return _pMPrinterGetOutputResolution(printer, printSettings, resolutionP), nil
}

// PMPrinterGetOutputResolution obtains the printer hardware output resolution for the specified print settings.
//
// See: https://developer.apple.com/documentation/applicationservices/1459076-pmprintergetoutputresolution
func PMPrinterGetOutputResolution(printer uintptr, printSettings uintptr, resolutionP unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterGetOutputResolution(printer, printSettings, resolutionP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetPaperList func(printer uintptr, paperList *corefoundation.CFArrayRef) int32
var _pMPrinterGetPaperListErr error

func tryPMPrinterGetPaperList(printer uintptr, paperList *corefoundation.CFArrayRef) (int32, error) {
	if _pMPrinterGetPaperList == nil {
		return 0, symbolCallError("PMPrinterGetPaperList", "10.3", _pMPrinterGetPaperListErr)
	}
	return _pMPrinterGetPaperList(printer, paperList), nil
}

// PMPrinterGetPaperList obtains the list of papers available for a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1460088-pmprintergetpaperlist
func PMPrinterGetPaperList(printer uintptr, paperList *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMPrinterGetPaperList(printer, paperList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetPrinterResolutionCount func(printer uintptr, countP *uint32) int32
var _pMPrinterGetPrinterResolutionCountErr error

func tryPMPrinterGetPrinterResolutionCount(printer uintptr, countP *uint32) (int32, error) {
	if _pMPrinterGetPrinterResolutionCount == nil {
		return 0, symbolCallError("PMPrinterGetPrinterResolutionCount", "10.0", _pMPrinterGetPrinterResolutionCountErr)
	}
	return _pMPrinterGetPrinterResolutionCount(printer, countP), nil
}

// PMPrinterGetPrinterResolutionCount obtains the number of resolution settings supported by the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1462004-pmprintergetprinterresolutioncou
func PMPrinterGetPrinterResolutionCount(printer uintptr, countP *uint32) int32 {
	result, callErr := tryPMPrinterGetPrinterResolutionCount(printer, countP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterGetState func(printer uintptr, state *PMPrinterState) int32
var _pMPrinterGetStateErr error

func tryPMPrinterGetState(printer uintptr, state *PMPrinterState) (int32, error) {
	if _pMPrinterGetState == nil {
		return 0, symbolCallError("PMPrinterGetState", "10.2", _pMPrinterGetStateErr)
	}
	return _pMPrinterGetState(printer, state), nil
}

// PMPrinterGetState obtains the current state of the print queue for a printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1462954-pmprintergetstate
func PMPrinterGetState(printer uintptr, state *PMPrinterState) int32 {
	result, callErr := tryPMPrinterGetState(printer, state)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterIsDefault func(printer uintptr) bool
var _pMPrinterIsDefaultErr error

func tryPMPrinterIsDefault(printer uintptr) (bool, error) {
	if _pMPrinterIsDefault == nil {
		return false, symbolCallError("PMPrinterIsDefault", "10.2", _pMPrinterIsDefaultErr)
	}
	return _pMPrinterIsDefault(printer), nil
}

// PMPrinterIsDefault returns a Boolean value indicating whether a printer is the default printer for the current user.
//
// See: https://developer.apple.com/documentation/applicationservices/1459030-pmprinterisdefault
func PMPrinterIsDefault(printer uintptr) bool {
	result, callErr := tryPMPrinterIsDefault(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterIsFavorite func(printer uintptr) bool
var _pMPrinterIsFavoriteErr error

func tryPMPrinterIsFavorite(printer uintptr) (bool, error) {
	if _pMPrinterIsFavorite == nil {
		return false, symbolCallError("PMPrinterIsFavorite", "10.2", _pMPrinterIsFavoriteErr)
	}
	return _pMPrinterIsFavorite(printer), nil
}

// PMPrinterIsFavorite returns a Boolean value indicating whether a printer is in the user’s list of favorite printers.
//
// See: https://developer.apple.com/documentation/applicationservices/1462074-pmprinterisfavorite
func PMPrinterIsFavorite(printer uintptr) bool {
	result, callErr := tryPMPrinterIsFavorite(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterIsPostScriptCapable func(printer uintptr) bool
var _pMPrinterIsPostScriptCapableErr error

func tryPMPrinterIsPostScriptCapable(printer uintptr) (bool, error) {
	if _pMPrinterIsPostScriptCapable == nil {
		return false, symbolCallError("PMPrinterIsPostScriptCapable", "10.2", _pMPrinterIsPostScriptCapableErr)
	}
	return _pMPrinterIsPostScriptCapable(printer), nil
}

// PMPrinterIsPostScriptCapable returns a Boolean value indicating whether a printer is PostScript capable.
//
// See: https://developer.apple.com/documentation/applicationservices/1464168-pmprinterispostscriptcapable
func PMPrinterIsPostScriptCapable(printer uintptr) bool {
	result, callErr := tryPMPrinterIsPostScriptCapable(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterIsPostScriptPrinter func(printer uintptr, isPSPrinter unsafe.Pointer) int32
var _pMPrinterIsPostScriptPrinterErr error

func tryPMPrinterIsPostScriptPrinter(printer uintptr, isPSPrinter unsafe.Pointer) (int32, error) {
	if _pMPrinterIsPostScriptPrinter == nil {
		return 0, symbolCallError("PMPrinterIsPostScriptPrinter", "10.5", _pMPrinterIsPostScriptPrinterErr)
	}
	return _pMPrinterIsPostScriptPrinter(printer, isPSPrinter), nil
}

// PMPrinterIsPostScriptPrinter determines whether a printer is a PostScript printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1462257-pmprinterispostscriptprinter
func PMPrinterIsPostScriptPrinter(printer uintptr, isPSPrinter unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterIsPostScriptPrinter(printer, isPSPrinter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterIsRemote func(printer uintptr, isRemoteP unsafe.Pointer) int32
var _pMPrinterIsRemoteErr error

func tryPMPrinterIsRemote(printer uintptr, isRemoteP unsafe.Pointer) (int32, error) {
	if _pMPrinterIsRemote == nil {
		return 0, symbolCallError("PMPrinterIsRemote", "10.3", _pMPrinterIsRemoteErr)
	}
	return _pMPrinterIsRemote(printer, isRemoteP), nil
}

// PMPrinterIsRemote indicates whether a printer is hosted by a remote print server.
//
// See: https://developer.apple.com/documentation/applicationservices/1461377-pmprinterisremote
func PMPrinterIsRemote(printer uintptr, isRemoteP unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterIsRemote(printer, isRemoteP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterPrintWithFile func(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, fileURL corefoundation.CFURLRef) int32
var _pMPrinterPrintWithFileErr error

func tryPMPrinterPrintWithFile(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, fileURL corefoundation.CFURLRef) (int32, error) {
	if _pMPrinterPrintWithFile == nil {
		return 0, symbolCallError("PMPrinterPrintWithFile", "10.3", _pMPrinterPrintWithFileErr)
	}
	return _pMPrinterPrintWithFile(printer, settings, format, mimeType, fileURL), nil
}

// PMPrinterPrintWithFile submits a print job to a specified printer using a file that contains print data.
//
// See: https://developer.apple.com/documentation/applicationservices/1464600-pmprinterprintwithfile
func PMPrinterPrintWithFile(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, fileURL corefoundation.CFURLRef) int32 {
	result, callErr := tryPMPrinterPrintWithFile(printer, settings, format, mimeType, fileURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterPrintWithProvider func(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, provider coregraphics.CGDataProviderRef) int32
var _pMPrinterPrintWithProviderErr error

func tryPMPrinterPrintWithProvider(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, provider coregraphics.CGDataProviderRef) (int32, error) {
	if _pMPrinterPrintWithProvider == nil {
		return 0, symbolCallError("PMPrinterPrintWithProvider", "10.3", _pMPrinterPrintWithProviderErr)
	}
	return _pMPrinterPrintWithProvider(printer, settings, format, mimeType, provider), nil
}

// PMPrinterPrintWithProvider submits a print job to a specified printer using a Quartz data provider to obtain the print data.
//
// See: https://developer.apple.com/documentation/applicationservices/1461110-pmprinterprintwithprovider
func PMPrinterPrintWithProvider(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, provider coregraphics.CGDataProviderRef) int32 {
	result, callErr := tryPMPrinterPrintWithProvider(printer, settings, format, mimeType, provider)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterSendCommand func(printer uintptr, commandString corefoundation.CFStringRef, jobTitle corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) int32
var _pMPrinterSendCommandErr error

func tryPMPrinterSendCommand(printer uintptr, commandString corefoundation.CFStringRef, jobTitle corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) (int32, error) {
	if _pMPrinterSendCommand == nil {
		return 0, symbolCallError("PMPrinterSendCommand", "10.6", _pMPrinterSendCommandErr)
	}
	return _pMPrinterSendCommand(printer, commandString, jobTitle, options), nil
}

// PMPrinterSendCommand.
//
// See: https://developer.apple.com/documentation/applicationservices/1463872-pmprintersendcommand
func PMPrinterSendCommand(printer uintptr, commandString corefoundation.CFStringRef, jobTitle corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryPMPrinterSendCommand(printer, commandString, jobTitle, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterSetDefault func(printer uintptr) int32
var _pMPrinterSetDefaultErr error

func tryPMPrinterSetDefault(printer uintptr) (int32, error) {
	if _pMPrinterSetDefault == nil {
		return 0, symbolCallError("PMPrinterSetDefault", "10.5", _pMPrinterSetDefaultErr)
	}
	return _pMPrinterSetDefault(printer), nil
}

// PMPrinterSetDefault sets the default printer for the current user.
//
// See: https://developer.apple.com/documentation/applicationservices/1461118-pmprintersetdefault
func PMPrinterSetDefault(printer uintptr) int32 {
	result, callErr := tryPMPrinterSetDefault(printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterSetOutputResolution func(printer uintptr, printSettings uintptr, resolutionP unsafe.Pointer) int32
var _pMPrinterSetOutputResolutionErr error

func tryPMPrinterSetOutputResolution(printer uintptr, printSettings uintptr, resolutionP unsafe.Pointer) (int32, error) {
	if _pMPrinterSetOutputResolution == nil {
		return 0, symbolCallError("PMPrinterSetOutputResolution", "10.5", _pMPrinterSetOutputResolutionErr)
	}
	return _pMPrinterSetOutputResolution(printer, printSettings, resolutionP), nil
}

// PMPrinterSetOutputResolution sets the print settings to reflect the specified printer hardware output resolution.
//
// See: https://developer.apple.com/documentation/applicationservices/1459931-pmprintersetoutputresolution
func PMPrinterSetOutputResolution(printer uintptr, printSettings uintptr, resolutionP unsafe.Pointer) int32 {
	result, callErr := tryPMPrinterSetOutputResolution(printer, printSettings, resolutionP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMPrinterWritePostScriptToURL func(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, sourceFileURL corefoundation.CFURLRef, destinationFileURL corefoundation.CFURLRef) int32
var _pMPrinterWritePostScriptToURLErr error

func tryPMPrinterWritePostScriptToURL(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, sourceFileURL corefoundation.CFURLRef, destinationFileURL corefoundation.CFURLRef) (int32, error) {
	if _pMPrinterWritePostScriptToURL == nil {
		return 0, symbolCallError("PMPrinterWritePostScriptToURL", "10.5", _pMPrinterWritePostScriptToURLErr)
	}
	return _pMPrinterWritePostScriptToURL(printer, settings, format, mimeType, sourceFileURL, destinationFileURL), nil
}

// PMPrinterWritePostScriptToURL converts an input file of the specified MIME type to printer-ready PostScript for a destination printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459729-pmprinterwritepostscripttourl
func PMPrinterWritePostScriptToURL(printer uintptr, settings uintptr, format uintptr, mimeType corefoundation.CFStringRef, sourceFileURL corefoundation.CFURLRef, destinationFileURL corefoundation.CFURLRef) int32 {
	result, callErr := tryPMPrinterWritePostScriptToURL(printer, settings, format, mimeType, sourceFileURL, destinationFileURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMRelease func(object PMObject) int32
var _pMReleaseErr error

func tryPMRelease(object PMObject) (int32, error) {
	if _pMRelease == nil {
		return 0, symbolCallError("PMRelease", "10.0", _pMReleaseErr)
	}
	return _pMRelease(object), nil
}

// PMRelease releases a printing object by decrementing its reference count.
//
// See: https://developer.apple.com/documentation/applicationservices/1461402-pmrelease
func PMRelease(object PMObject) int32 {
	result, callErr := tryPMRelease(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMRetain func(object PMObject) int32
var _pMRetainErr error

func tryPMRetain(object PMObject) (int32, error) {
	if _pMRetain == nil {
		return 0, symbolCallError("PMRetain", "10.0", _pMRetainErr)
	}
	return _pMRetain(object), nil
}

// PMRetain retains a printing object by incrementing its reference count.
//
// See: https://developer.apple.com/documentation/applicationservices/1460190-pmretain
func PMRetain(object PMObject) int32 {
	result, callErr := tryPMRetain(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMServerCreatePrinterList func(server PMServer, printerList *corefoundation.CFArrayRef) int32
var _pMServerCreatePrinterListErr error

func tryPMServerCreatePrinterList(server PMServer, printerList *corefoundation.CFArrayRef) (int32, error) {
	if _pMServerCreatePrinterList == nil {
		return 0, symbolCallError("PMServerCreatePrinterList", "10.2", _pMServerCreatePrinterListErr)
	}
	return _pMServerCreatePrinterList(server, printerList), nil
}

// PMServerCreatePrinterList creates a list of printers available to a print server.
//
// See: https://developer.apple.com/documentation/applicationservices/1459953-pmservercreateprinterlist
func PMServerCreatePrinterList(server PMServer, printerList *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMServerCreatePrinterList(server, printerList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMServerLaunchPrinterBrowser func(server PMServer, options corefoundation.CFDictionaryRef) int32
var _pMServerLaunchPrinterBrowserErr error

func tryPMServerLaunchPrinterBrowser(server PMServer, options corefoundation.CFDictionaryRef) (int32, error) {
	if _pMServerLaunchPrinterBrowser == nil {
		return 0, symbolCallError("PMServerLaunchPrinterBrowser", "10.5", _pMServerLaunchPrinterBrowserErr)
	}
	return _pMServerLaunchPrinterBrowser(server, options), nil
}

// PMServerLaunchPrinterBrowser launches the printer browser to browse the printers available for a print server.
//
// See: https://developer.apple.com/documentation/applicationservices/1460175-pmserverlaunchprinterbrowser
func PMServerLaunchPrinterBrowser(server PMServer, options corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryPMServerLaunchPrinterBrowser(server, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionBeginCGDocumentNoDialog func(printSession PMPrintSession, printSettings uintptr, pageFormat uintptr) int32
var _pMSessionBeginCGDocumentNoDialogErr error

func tryPMSessionBeginCGDocumentNoDialog(printSession PMPrintSession, printSettings uintptr, pageFormat uintptr) (int32, error) {
	if _pMSessionBeginCGDocumentNoDialog == nil {
		return 0, symbolCallError("PMSessionBeginCGDocumentNoDialog", "10.4", _pMSessionBeginCGDocumentNoDialogErr)
	}
	return _pMSessionBeginCGDocumentNoDialog(printSession, printSettings, pageFormat), nil
}

// PMSessionBeginCGDocumentNoDialog begins a print job that draws into a Quartz graphics context and suppresses the printing status dialog.
//
// See: https://developer.apple.com/documentation/applicationservices/1460101-pmsessionbegincgdocumentnodialog
func PMSessionBeginCGDocumentNoDialog(printSession PMPrintSession, printSettings uintptr, pageFormat uintptr) int32 {
	result, callErr := tryPMSessionBeginCGDocumentNoDialog(printSession, printSettings, pageFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionBeginPageNoDialog func(printSession PMPrintSession, pageFormat uintptr, pageFrame unsafe.Pointer) int32
var _pMSessionBeginPageNoDialogErr error

func tryPMSessionBeginPageNoDialog(printSession PMPrintSession, pageFormat uintptr, pageFrame unsafe.Pointer) (int32, error) {
	if _pMSessionBeginPageNoDialog == nil {
		return 0, symbolCallError("PMSessionBeginPageNoDialog", "10.2", _pMSessionBeginPageNoDialogErr)
	}
	return _pMSessionBeginPageNoDialog(printSession, pageFormat, pageFrame), nil
}

// PMSessionBeginPageNoDialog starts a new page for printing in the specified printing session and suppresses the printing status dialog.
//
// See: https://developer.apple.com/documentation/applicationservices/1463416-pmsessionbeginpagenodialog
func PMSessionBeginPageNoDialog(printSession PMPrintSession, pageFormat uintptr, pageFrame unsafe.Pointer) int32 {
	result, callErr := tryPMSessionBeginPageNoDialog(printSession, pageFormat, pageFrame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionCopyDestinationFormat func(printSession PMPrintSession, printSettings uintptr, destFormatP *corefoundation.CFStringRef) int32
var _pMSessionCopyDestinationFormatErr error

func tryPMSessionCopyDestinationFormat(printSession PMPrintSession, printSettings uintptr, destFormatP *corefoundation.CFStringRef) (int32, error) {
	if _pMSessionCopyDestinationFormat == nil {
		return 0, symbolCallError("PMSessionCopyDestinationFormat", "10.1", _pMSessionCopyDestinationFormatErr)
	}
	return _pMSessionCopyDestinationFormat(printSession, printSettings, destFormatP), nil
}

// PMSessionCopyDestinationFormat obtains the destination format for a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/1464266-pmsessioncopydestinationformat
func PMSessionCopyDestinationFormat(printSession PMPrintSession, printSettings uintptr, destFormatP *corefoundation.CFStringRef) int32 {
	result, callErr := tryPMSessionCopyDestinationFormat(printSession, printSettings, destFormatP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionCopyDestinationLocation func(printSession PMPrintSession, printSettings uintptr, destLocationP *corefoundation.CFURLRef) int32
var _pMSessionCopyDestinationLocationErr error

func tryPMSessionCopyDestinationLocation(printSession PMPrintSession, printSettings uintptr, destLocationP *corefoundation.CFURLRef) (int32, error) {
	if _pMSessionCopyDestinationLocation == nil {
		return 0, symbolCallError("PMSessionCopyDestinationLocation", "10.1", _pMSessionCopyDestinationLocationErr)
	}
	return _pMSessionCopyDestinationLocation(printSession, printSettings, destLocationP), nil
}

// PMSessionCopyDestinationLocation obtains a destination location for a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/1462967-pmsessioncopydestinationlocation
func PMSessionCopyDestinationLocation(printSession PMPrintSession, printSettings uintptr, destLocationP *corefoundation.CFURLRef) int32 {
	result, callErr := tryPMSessionCopyDestinationLocation(printSession, printSettings, destLocationP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionCopyOutputFormatList func(printSession PMPrintSession, destType PMDestinationType, documentFormatP *corefoundation.CFArrayRef) int32
var _pMSessionCopyOutputFormatListErr error

func tryPMSessionCopyOutputFormatList(printSession PMPrintSession, destType PMDestinationType, documentFormatP *corefoundation.CFArrayRef) (int32, error) {
	if _pMSessionCopyOutputFormatList == nil {
		return 0, symbolCallError("PMSessionCopyOutputFormatList", "10.1", _pMSessionCopyOutputFormatListErr)
	}
	return _pMSessionCopyOutputFormatList(printSession, destType, documentFormatP), nil
}

// PMSessionCopyOutputFormatList obtains an array of destination formats supported by the current print destination.
//
// See: https://developer.apple.com/documentation/applicationservices/1461332-pmsessioncopyoutputformatlist
func PMSessionCopyOutputFormatList(printSession PMPrintSession, destType PMDestinationType, documentFormatP *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMSessionCopyOutputFormatList(printSession, destType, documentFormatP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionCreatePageFormatList func(printSession PMPrintSession, printer uintptr, pageFormatList *corefoundation.CFArrayRef) int32
var _pMSessionCreatePageFormatListErr error

func tryPMSessionCreatePageFormatList(printSession PMPrintSession, printer uintptr, pageFormatList *corefoundation.CFArrayRef) (int32, error) {
	if _pMSessionCreatePageFormatList == nil {
		return 0, symbolCallError("PMSessionCreatePageFormatList", "10.1", _pMSessionCreatePageFormatListErr)
	}
	return _pMSessionCreatePageFormatList(printSession, printer, pageFormatList), nil
}

// PMSessionCreatePageFormatList obtains a list of page format objects, each of which describes a paper size available on the specified printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1463985-pmsessioncreatepageformatlist
func PMSessionCreatePageFormatList(printSession PMPrintSession, printer uintptr, pageFormatList *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMSessionCreatePageFormatList(printSession, printer, pageFormatList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionCreatePrinterList func(printSession PMPrintSession, printerList *corefoundation.CFArrayRef, currentIndex *int, currentPrinter *uintptr) int32
var _pMSessionCreatePrinterListErr error

func tryPMSessionCreatePrinterList(printSession PMPrintSession, printerList *corefoundation.CFArrayRef, currentIndex *int, currentPrinter *uintptr) (int32, error) {
	if _pMSessionCreatePrinterList == nil {
		return 0, symbolCallError("PMSessionCreatePrinterList", "10.1", _pMSessionCreatePrinterListErr)
	}
	return _pMSessionCreatePrinterList(printSession, printerList, currentIndex, currentPrinter), nil
}

// PMSessionCreatePrinterList creates a list of printers available in the specified printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1460119-pmsessioncreateprinterlist
func PMSessionCreatePrinterList(printSession PMPrintSession, printerList *corefoundation.CFArrayRef, currentIndex *int, currentPrinter *uintptr) int32 {
	result, callErr := tryPMSessionCreatePrinterList(printSession, printerList, currentIndex, currentPrinter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionDefaultPageFormat func(printSession PMPrintSession, pageFormat uintptr) int32
var _pMSessionDefaultPageFormatErr error

func tryPMSessionDefaultPageFormat(printSession PMPrintSession, pageFormat uintptr) (int32, error) {
	if _pMSessionDefaultPageFormat == nil {
		return 0, symbolCallError("PMSessionDefaultPageFormat", "10.0", _pMSessionDefaultPageFormatErr)
	}
	return _pMSessionDefaultPageFormat(printSession, pageFormat), nil
}

// PMSessionDefaultPageFormat assigns default parameter values to a page format object used in the specified printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1462217-pmsessiondefaultpageformat
func PMSessionDefaultPageFormat(printSession PMPrintSession, pageFormat uintptr) int32 {
	result, callErr := tryPMSessionDefaultPageFormat(printSession, pageFormat)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionDefaultPrintSettings func(printSession PMPrintSession, printSettings uintptr) int32
var _pMSessionDefaultPrintSettingsErr error

func tryPMSessionDefaultPrintSettings(printSession PMPrintSession, printSettings uintptr) (int32, error) {
	if _pMSessionDefaultPrintSettings == nil {
		return 0, symbolCallError("PMSessionDefaultPrintSettings", "10.0", _pMSessionDefaultPrintSettingsErr)
	}
	return _pMSessionDefaultPrintSettings(printSession, printSettings), nil
}

// PMSessionDefaultPrintSettings assigns default parameter values to a print settings object for the specified printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1460138-pmsessiondefaultprintsettings
func PMSessionDefaultPrintSettings(printSession PMPrintSession, printSettings uintptr) int32 {
	result, callErr := tryPMSessionDefaultPrintSettings(printSession, printSettings)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionEndDocumentNoDialog func(printSession PMPrintSession) int32
var _pMSessionEndDocumentNoDialogErr error

func tryPMSessionEndDocumentNoDialog(printSession PMPrintSession) (int32, error) {
	if _pMSessionEndDocumentNoDialog == nil {
		return 0, symbolCallError("PMSessionEndDocumentNoDialog", "10.2", _pMSessionEndDocumentNoDialogErr)
	}
	return _pMSessionEndDocumentNoDialog(printSession), nil
}

// PMSessionEndDocumentNoDialog ends a print job started by calling the function PMSessionBeginCGDocumentNoDialog(_:_:_:) or PMSessionBeginDocumentNoDialog.
//
// See: https://developer.apple.com/documentation/applicationservices/1464527-pmsessionenddocumentnodialog
func PMSessionEndDocumentNoDialog(printSession PMPrintSession) int32 {
	result, callErr := tryPMSessionEndDocumentNoDialog(printSession)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionEndPageNoDialog func(printSession PMPrintSession) int32
var _pMSessionEndPageNoDialogErr error

func tryPMSessionEndPageNoDialog(printSession PMPrintSession) (int32, error) {
	if _pMSessionEndPageNoDialog == nil {
		return 0, symbolCallError("PMSessionEndPageNoDialog", "10.2", _pMSessionEndPageNoDialogErr)
	}
	return _pMSessionEndPageNoDialog(printSession), nil
}

// PMSessionEndPageNoDialog indicates the end of drawing the current page for the specified printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1462014-pmsessionendpagenodialog
func PMSessionEndPageNoDialog(printSession PMPrintSession) int32 {
	result, callErr := tryPMSessionEndPageNoDialog(printSession)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionError func(printSession PMPrintSession) int32
var _pMSessionErrorErr error

func tryPMSessionError(printSession PMPrintSession) (int32, error) {
	if _pMSessionError == nil {
		return 0, symbolCallError("PMSessionError", "10.0", _pMSessionErrorErr)
	}
	return _pMSessionError(printSession), nil
}

// PMSessionError obtains the result code for any error returned by the printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1460003-pmsessionerror
func PMSessionError(printSession PMPrintSession) int32 {
	result, callErr := tryPMSessionError(printSession)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionGetCGGraphicsContext func(printSession PMPrintSession, context *coregraphics.CGContextRef) int32
var _pMSessionGetCGGraphicsContextErr error

func tryPMSessionGetCGGraphicsContext(printSession PMPrintSession, context *coregraphics.CGContextRef) (int32, error) {
	if _pMSessionGetCGGraphicsContext == nil {
		return 0, symbolCallError("PMSessionGetCGGraphicsContext", "10.4", _pMSessionGetCGGraphicsContextErr)
	}
	return _pMSessionGetCGGraphicsContext(printSession, context), nil
}

// PMSessionGetCGGraphicsContext obtains the Quartz graphics context for the current page in a printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1461952-pmsessiongetcggraphicscontext
func PMSessionGetCGGraphicsContext(printSession PMPrintSession, context *coregraphics.CGContextRef) int32 {
	result, callErr := tryPMSessionGetCGGraphicsContext(printSession, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionGetCurrentPrinter func(printSession PMPrintSession, currentPrinter *uintptr) int32
var _pMSessionGetCurrentPrinterErr error

func tryPMSessionGetCurrentPrinter(printSession PMPrintSession, currentPrinter *uintptr) (int32, error) {
	if _pMSessionGetCurrentPrinter == nil {
		return 0, symbolCallError("PMSessionGetCurrentPrinter", "10.0", _pMSessionGetCurrentPrinterErr)
	}
	return _pMSessionGetCurrentPrinter(printSession, currentPrinter), nil
}

// PMSessionGetCurrentPrinter obtains the current printer associated with a printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1458998-pmsessiongetcurrentprinter
func PMSessionGetCurrentPrinter(printSession PMPrintSession, currentPrinter *uintptr) int32 {
	result, callErr := tryPMSessionGetCurrentPrinter(printSession, currentPrinter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionGetDataFromSession func(printSession PMPrintSession, key corefoundation.CFStringRef, data *corefoundation.CFTypeRef) int32
var _pMSessionGetDataFromSessionErr error

func tryPMSessionGetDataFromSession(printSession PMPrintSession, key corefoundation.CFStringRef, data *corefoundation.CFTypeRef) (int32, error) {
	if _pMSessionGetDataFromSession == nil {
		return 0, symbolCallError("PMSessionGetDataFromSession", "10.0", _pMSessionGetDataFromSessionErr)
	}
	return _pMSessionGetDataFromSession(printSession, key, data), nil
}

// PMSessionGetDataFromSession obtains application-specific data previously stored in a printing session object.
//
// See: https://developer.apple.com/documentation/applicationservices/1462964-pmsessiongetdatafromsession
func PMSessionGetDataFromSession(printSession PMPrintSession, key corefoundation.CFStringRef, data *corefoundation.CFTypeRef) int32 {
	result, callErr := tryPMSessionGetDataFromSession(printSession, key, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionGetDestinationType func(printSession PMPrintSession, printSettings uintptr, destTypeP *PMDestinationType) int32
var _pMSessionGetDestinationTypeErr error

func tryPMSessionGetDestinationType(printSession PMPrintSession, printSettings uintptr, destTypeP *PMDestinationType) (int32, error) {
	if _pMSessionGetDestinationType == nil {
		return 0, symbolCallError("PMSessionGetDestinationType", "10.1", _pMSessionGetDestinationTypeErr)
	}
	return _pMSessionGetDestinationType(printSession, printSettings, destTypeP), nil
}

// PMSessionGetDestinationType obtains the output destination for a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/1461071-pmsessiongetdestinationtype
func PMSessionGetDestinationType(printSession PMPrintSession, printSettings uintptr, destTypeP *PMDestinationType) int32 {
	result, callErr := tryPMSessionGetDestinationType(printSession, printSettings, destTypeP)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionSetCurrentPMPrinter func(session PMPrintSession, printer uintptr) int32
var _pMSessionSetCurrentPMPrinterErr error

func tryPMSessionSetCurrentPMPrinter(session PMPrintSession, printer uintptr) (int32, error) {
	if _pMSessionSetCurrentPMPrinter == nil {
		return 0, symbolCallError("PMSessionSetCurrentPMPrinter", "10.3", _pMSessionSetCurrentPMPrinterErr)
	}
	return _pMSessionSetCurrentPMPrinter(session, printer), nil
}

// PMSessionSetCurrentPMPrinter changes the current printer for a printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1461096-pmsessionsetcurrentpmprinter
func PMSessionSetCurrentPMPrinter(session PMPrintSession, printer uintptr) int32 {
	result, callErr := tryPMSessionSetCurrentPMPrinter(session, printer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionSetDataInSession func(printSession PMPrintSession, key corefoundation.CFStringRef, data corefoundation.CFTypeRef) int32
var _pMSessionSetDataInSessionErr error

func tryPMSessionSetDataInSession(printSession PMPrintSession, key corefoundation.CFStringRef, data corefoundation.CFTypeRef) (int32, error) {
	if _pMSessionSetDataInSession == nil {
		return 0, symbolCallError("PMSessionSetDataInSession", "10.0", _pMSessionSetDataInSessionErr)
	}
	return _pMSessionSetDataInSession(printSession, key, data), nil
}

// PMSessionSetDataInSession stores your application-specific data in a printing session object.
//
// See: https://developer.apple.com/documentation/applicationservices/1461902-pmsessionsetdatainsession
func PMSessionSetDataInSession(printSession PMPrintSession, key corefoundation.CFStringRef, data corefoundation.CFTypeRef) int32 {
	result, callErr := tryPMSessionSetDataInSession(printSession, key, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionSetDestination func(printSession PMPrintSession, printSettings uintptr, destType PMDestinationType, destFormat corefoundation.CFStringRef, destLocation corefoundation.CFURLRef) int32
var _pMSessionSetDestinationErr error

func tryPMSessionSetDestination(printSession PMPrintSession, printSettings uintptr, destType PMDestinationType, destFormat corefoundation.CFStringRef, destLocation corefoundation.CFURLRef) (int32, error) {
	if _pMSessionSetDestination == nil {
		return 0, symbolCallError("PMSessionSetDestination", "10.1", _pMSessionSetDestinationErr)
	}
	return _pMSessionSetDestination(printSession, printSettings, destType, destFormat, destLocation), nil
}

// PMSessionSetDestination sets the destination location, format, and type for a print job.
//
// See: https://developer.apple.com/documentation/applicationservices/1459855-pmsessionsetdestination
func PMSessionSetDestination(printSession PMPrintSession, printSettings uintptr, destType PMDestinationType, destFormat corefoundation.CFStringRef, destLocation corefoundation.CFURLRef) int32 {
	result, callErr := tryPMSessionSetDestination(printSession, printSettings, destType, destFormat, destLocation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionSetError func(printSession PMPrintSession, printError int32) int32
var _pMSessionSetErrorErr error

func tryPMSessionSetError(printSession PMPrintSession, printError int32) (int32, error) {
	if _pMSessionSetError == nil {
		return 0, symbolCallError("PMSessionSetError", "10.0", _pMSessionSetErrorErr)
	}
	return _pMSessionSetError(printSession, printError), nil
}

// PMSessionSetError sets the value of the current result code for the specified printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1460216-pmsessionseterror
func PMSessionSetError(printSession PMPrintSession, printError int32) int32 {
	result, callErr := tryPMSessionSetError(printSession, printError)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionValidatePageFormat func(printSession PMPrintSession, pageFormat uintptr, changed unsafe.Pointer) int32
var _pMSessionValidatePageFormatErr error

func tryPMSessionValidatePageFormat(printSession PMPrintSession, pageFormat uintptr, changed unsafe.Pointer) (int32, error) {
	if _pMSessionValidatePageFormat == nil {
		return 0, symbolCallError("PMSessionValidatePageFormat", "10.0", _pMSessionValidatePageFormatErr)
	}
	return _pMSessionValidatePageFormat(printSession, pageFormat, changed), nil
}

// PMSessionValidatePageFormat updates the values in a page format object and validates them against the current formatting printer.
//
// See: https://developer.apple.com/documentation/applicationservices/1459090-pmsessionvalidatepageformat
func PMSessionValidatePageFormat(printSession PMPrintSession, pageFormat uintptr, changed unsafe.Pointer) int32 {
	result, callErr := tryPMSessionValidatePageFormat(printSession, pageFormat, changed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSessionValidatePrintSettings func(printSession PMPrintSession, printSettings uintptr, changed unsafe.Pointer) int32
var _pMSessionValidatePrintSettingsErr error

func tryPMSessionValidatePrintSettings(printSession PMPrintSession, printSettings uintptr, changed unsafe.Pointer) (int32, error) {
	if _pMSessionValidatePrintSettings == nil {
		return 0, symbolCallError("PMSessionValidatePrintSettings", "10.0", _pMSessionValidatePrintSettingsErr)
	}
	return _pMSessionValidatePrintSettings(printSession, printSettings, changed), nil
}

// PMSessionValidatePrintSettings validates a print settings object within the context of the specified printing session.
//
// See: https://developer.apple.com/documentation/applicationservices/1458994-pmsessionvalidateprintsettings
func PMSessionValidatePrintSettings(printSession PMPrintSession, printSettings uintptr, changed unsafe.Pointer) int32 {
	result, callErr := tryPMSessionValidatePrintSettings(printSession, printSettings, changed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetCollate func(printSettings uintptr, collate unsafe.Pointer) int32
var _pMSetCollateErr error

func tryPMSetCollate(printSettings uintptr, collate unsafe.Pointer) (int32, error) {
	if _pMSetCollate == nil {
		return 0, symbolCallError("PMSetCollate", "10.2", _pMSetCollateErr)
	}
	return _pMSetCollate(printSettings, collate), nil
}

// PMSetCollate specifies whether the job collate option is selected.
//
// See: https://developer.apple.com/documentation/applicationservices/1463223-pmsetcollate
func PMSetCollate(printSettings uintptr, collate unsafe.Pointer) int32 {
	result, callErr := tryPMSetCollate(printSettings, collate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetCopies func(printSettings uintptr, copies uint32, lock unsafe.Pointer) int32
var _pMSetCopiesErr error

func tryPMSetCopies(printSettings uintptr, copies uint32, lock unsafe.Pointer) (int32, error) {
	if _pMSetCopies == nil {
		return 0, symbolCallError("PMSetCopies", "10.0", _pMSetCopiesErr)
	}
	return _pMSetCopies(printSettings, copies, lock), nil
}

// PMSetCopies sets the initial value for the number of copies to be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1463804-pmsetcopies
func PMSetCopies(printSettings uintptr, copies uint32, lock unsafe.Pointer) int32 {
	result, callErr := tryPMSetCopies(printSettings, copies, lock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetDuplex func(printSettings uintptr, duplexSetting PMDuplexMode) int32
var _pMSetDuplexErr error

func tryPMSetDuplex(printSettings uintptr, duplexSetting PMDuplexMode) (int32, error) {
	if _pMSetDuplex == nil {
		return 0, symbolCallError("PMSetDuplex", "10.4", _pMSetDuplexErr)
	}
	return _pMSetDuplex(printSettings, duplexSetting), nil
}

// PMSetDuplex sets the duplex mode.
//
// See: https://developer.apple.com/documentation/applicationservices/1462000-pmsetduplex
func PMSetDuplex(printSettings uintptr, duplexSetting PMDuplexMode) int32 {
	result, callErr := tryPMSetDuplex(printSettings, duplexSetting)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetFirstPage func(printSettings uintptr, first uint32, lock unsafe.Pointer) int32
var _pMSetFirstPageErr error

func tryPMSetFirstPage(printSettings uintptr, first uint32, lock unsafe.Pointer) (int32, error) {
	if _pMSetFirstPage == nil {
		return 0, symbolCallError("PMSetFirstPage", "10.0", _pMSetFirstPageErr)
	}
	return _pMSetFirstPage(printSettings, first, lock), nil
}

// PMSetFirstPage sets the default page number of the first page to be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1461519-pmsetfirstpage
func PMSetFirstPage(printSettings uintptr, first uint32, lock unsafe.Pointer) int32 {
	result, callErr := tryPMSetFirstPage(printSettings, first, lock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetLastPage func(printSettings uintptr, last uint32, lock unsafe.Pointer) int32
var _pMSetLastPageErr error

func tryPMSetLastPage(printSettings uintptr, last uint32, lock unsafe.Pointer) (int32, error) {
	if _pMSetLastPage == nil {
		return 0, symbolCallError("PMSetLastPage", "10.0", _pMSetLastPageErr)
	}
	return _pMSetLastPage(printSettings, last, lock), nil
}

// PMSetLastPage sets the page number of the last page to be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1463595-pmsetlastpage
func PMSetLastPage(printSettings uintptr, last uint32, lock unsafe.Pointer) int32 {
	result, callErr := tryPMSetLastPage(printSettings, last, lock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetOrientation func(pageFormat uintptr, orientation PMOrientation, lock unsafe.Pointer) int32
var _pMSetOrientationErr error

func tryPMSetOrientation(pageFormat uintptr, orientation PMOrientation, lock unsafe.Pointer) (int32, error) {
	if _pMSetOrientation == nil {
		return 0, symbolCallError("PMSetOrientation", "10.0", _pMSetOrientationErr)
	}
	return _pMSetOrientation(pageFormat, orientation, lock), nil
}

// PMSetOrientation sets the page orientation for printing.
//
// See: https://developer.apple.com/documentation/applicationservices/1459016-pmsetorientation
func PMSetOrientation(pageFormat uintptr, orientation PMOrientation, lock unsafe.Pointer) int32 {
	result, callErr := tryPMSetOrientation(pageFormat, orientation, lock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetPageFormatExtendedData func(pageFormat uintptr, dataID uint32, size uint32, extendedData unsafe.Pointer) int32
var _pMSetPageFormatExtendedDataErr error

func tryPMSetPageFormatExtendedData(pageFormat uintptr, dataID uint32, size uint32, extendedData unsafe.Pointer) (int32, error) {
	if _pMSetPageFormatExtendedData == nil {
		return 0, symbolCallError("PMSetPageFormatExtendedData", "10.0", _pMSetPageFormatExtendedDataErr)
	}
	return _pMSetPageFormatExtendedData(pageFormat, dataID, size, extendedData), nil
}

// PMSetPageFormatExtendedData stores your application-specific data in a page format object.
//
// See: https://developer.apple.com/documentation/applicationservices/1463464-pmsetpageformatextendeddata
func PMSetPageFormatExtendedData(pageFormat uintptr, dataID uint32, size uint32, extendedData unsafe.Pointer) int32 {
	result, callErr := tryPMSetPageFormatExtendedData(pageFormat, dataID, size, extendedData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetPageRange func(printSettings uintptr, minPage uint32, maxPage uint32) int32
var _pMSetPageRangeErr error

func tryPMSetPageRange(printSettings uintptr, minPage uint32, maxPage uint32) (int32, error) {
	if _pMSetPageRange == nil {
		return 0, symbolCallError("PMSetPageRange", "10.0", _pMSetPageRangeErr)
	}
	return _pMSetPageRange(printSettings, minPage, maxPage), nil
}

// PMSetPageRange sets the valid range of pages that can be printed.
//
// See: https://developer.apple.com/documentation/applicationservices/1462294-pmsetpagerange
func PMSetPageRange(printSettings uintptr, minPage uint32, maxPage uint32) int32 {
	result, callErr := tryPMSetPageRange(printSettings, minPage, maxPage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMSetScale func(pageFormat uintptr, scale unsafe.Pointer) int32
var _pMSetScaleErr error

func tryPMSetScale(pageFormat uintptr, scale unsafe.Pointer) (int32, error) {
	if _pMSetScale == nil {
		return 0, symbolCallError("PMSetScale", "10.0", _pMSetScaleErr)
	}
	return _pMSetScale(pageFormat, scale), nil
}

// PMSetScale sets the scaling factor for the page and paper rectangles.
//
// See: https://developer.apple.com/documentation/applicationservices/1463343-pmsetscale
func PMSetScale(pageFormat uintptr, scale unsafe.Pointer) int32 {
	result, callErr := tryPMSetScale(pageFormat, scale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMWorkflowCopyItems func(workflowItems *corefoundation.CFArrayRef) int32
var _pMWorkflowCopyItemsErr error

func tryPMWorkflowCopyItems(workflowItems *corefoundation.CFArrayRef) (int32, error) {
	if _pMWorkflowCopyItems == nil {
		return 0, symbolCallError("PMWorkflowCopyItems", "10.3", _pMWorkflowCopyItemsErr)
	}
	return _pMWorkflowCopyItems(workflowItems), nil
}

// PMWorkflowCopyItems obtains an array of the available PDF workflow items.
//
// See: https://developer.apple.com/documentation/applicationservices/1459914-pmworkflowcopyitems
func PMWorkflowCopyItems(workflowItems *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPMWorkflowCopyItems(workflowItems)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMWorkflowSubmitPDFWithOptions func(workflowItem corefoundation.CFURLRef, title corefoundation.CFStringRef, options string, pdfFile corefoundation.CFURLRef) int32
var _pMWorkflowSubmitPDFWithOptionsErr error

func tryPMWorkflowSubmitPDFWithOptions(workflowItem corefoundation.CFURLRef, title corefoundation.CFStringRef, options string, pdfFile corefoundation.CFURLRef) (int32, error) {
	if _pMWorkflowSubmitPDFWithOptions == nil {
		return 0, symbolCallError("PMWorkflowSubmitPDFWithOptions", "10.3", _pMWorkflowSubmitPDFWithOptionsErr)
	}
	return _pMWorkflowSubmitPDFWithOptions(workflowItem, title, options, pdfFile), nil
}

// PMWorkflowSubmitPDFWithOptions submits a PDF file for workflow processing using the specified CUPS options string.
//
// See: https://developer.apple.com/documentation/applicationservices/1463747-pmworkflowsubmitpdfwithoptions
func PMWorkflowSubmitPDFWithOptions(workflowItem corefoundation.CFURLRef, title corefoundation.CFStringRef, options string, pdfFile corefoundation.CFURLRef) int32 {
	result, callErr := tryPMWorkflowSubmitPDFWithOptions(workflowItem, title, options, pdfFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pMWorkflowSubmitPDFWithSettings func(workflowItem corefoundation.CFURLRef, settings uintptr, pdfFile corefoundation.CFURLRef) int32
var _pMWorkflowSubmitPDFWithSettingsErr error

func tryPMWorkflowSubmitPDFWithSettings(workflowItem corefoundation.CFURLRef, settings uintptr, pdfFile corefoundation.CFURLRef) (int32, error) {
	if _pMWorkflowSubmitPDFWithSettings == nil {
		return 0, symbolCallError("PMWorkflowSubmitPDFWithSettings", "10.3", _pMWorkflowSubmitPDFWithSettingsErr)
	}
	return _pMWorkflowSubmitPDFWithSettings(workflowItem, settings, pdfFile), nil
}

// PMWorkflowSubmitPDFWithSettings submits a PDF file for workflow processing using the specified print settings.
//
// See: https://developer.apple.com/documentation/applicationservices/1458874-pmworkflowsubmitpdfwithsettings
func PMWorkflowSubmitPDFWithSettings(workflowItem corefoundation.CFURLRef, settings uintptr, pdfFile corefoundation.CFURLRef) int32 {
	result, callErr := tryPMWorkflowSubmitPDFWithSettings(workflowItem, settings, pdfFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardClear func(inPasteboard PasteboardRef) int32
var _pasteboardClearErr error

func tryPasteboardClear(inPasteboard PasteboardRef) (int32, error) {
	if _pasteboardClear == nil {
		return 0, symbolCallError("PasteboardClear", "10.3", _pasteboardClearErr)
	}
	return _pasteboardClear(inPasteboard), nil
}

// PasteboardClear.
//
// See: https://developer.apple.com/documentation/applicationservices/1460800-pasteboardclear
func PasteboardClear(inPasteboard PasteboardRef) int32 {
	result, callErr := tryPasteboardClear(inPasteboard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardCopyItemFlavorData func(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, outData *corefoundation.CFDataRef) int32
var _pasteboardCopyItemFlavorDataErr error

func tryPasteboardCopyItemFlavorData(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, outData *corefoundation.CFDataRef) (int32, error) {
	if _pasteboardCopyItemFlavorData == nil {
		return 0, symbolCallError("PasteboardCopyItemFlavorData", "10.3", _pasteboardCopyItemFlavorDataErr)
	}
	return _pasteboardCopyItemFlavorData(inPasteboard, inItem, inFlavorType, outData), nil
}

// PasteboardCopyItemFlavorData.
//
// See: https://developer.apple.com/documentation/applicationservices/1458917-pasteboardcopyitemflavordata
func PasteboardCopyItemFlavorData(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, outData *corefoundation.CFDataRef) int32 {
	result, callErr := tryPasteboardCopyItemFlavorData(inPasteboard, inItem, inFlavorType, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardCopyItemFlavors func(inPasteboard PasteboardRef, inItem PasteboardItemID, outFlavorTypes *corefoundation.CFArrayRef) int32
var _pasteboardCopyItemFlavorsErr error

func tryPasteboardCopyItemFlavors(inPasteboard PasteboardRef, inItem PasteboardItemID, outFlavorTypes *corefoundation.CFArrayRef) (int32, error) {
	if _pasteboardCopyItemFlavors == nil {
		return 0, symbolCallError("PasteboardCopyItemFlavors", "10.3", _pasteboardCopyItemFlavorsErr)
	}
	return _pasteboardCopyItemFlavors(inPasteboard, inItem, outFlavorTypes), nil
}

// PasteboardCopyItemFlavors.
//
// See: https://developer.apple.com/documentation/applicationservices/1460005-pasteboardcopyitemflavors
func PasteboardCopyItemFlavors(inPasteboard PasteboardRef, inItem PasteboardItemID, outFlavorTypes *corefoundation.CFArrayRef) int32 {
	result, callErr := tryPasteboardCopyItemFlavors(inPasteboard, inItem, outFlavorTypes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardCopyName func(inPasteboard PasteboardRef, outName *corefoundation.CFStringRef) int32
var _pasteboardCopyNameErr error

func tryPasteboardCopyName(inPasteboard PasteboardRef, outName *corefoundation.CFStringRef) (int32, error) {
	if _pasteboardCopyName == nil {
		return 0, symbolCallError("PasteboardCopyName", "10.4", _pasteboardCopyNameErr)
	}
	return _pasteboardCopyName(inPasteboard, outName), nil
}

// PasteboardCopyName.
//
// See: https://developer.apple.com/documentation/applicationservices/1459455-pasteboardcopyname
func PasteboardCopyName(inPasteboard PasteboardRef, outName *corefoundation.CFStringRef) int32 {
	result, callErr := tryPasteboardCopyName(inPasteboard, outName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardCopyPasteLocation func(inPasteboard PasteboardRef, outPasteLocation *corefoundation.CFURLRef) int32
var _pasteboardCopyPasteLocationErr error

func tryPasteboardCopyPasteLocation(inPasteboard PasteboardRef, outPasteLocation *corefoundation.CFURLRef) (int32, error) {
	if _pasteboardCopyPasteLocation == nil {
		return 0, symbolCallError("PasteboardCopyPasteLocation", "10.3", _pasteboardCopyPasteLocationErr)
	}
	return _pasteboardCopyPasteLocation(inPasteboard, outPasteLocation), nil
}

// PasteboardCopyPasteLocation.
//
// See: https://developer.apple.com/documentation/applicationservices/1462546-pasteboardcopypastelocation
func PasteboardCopyPasteLocation(inPasteboard PasteboardRef, outPasteLocation *corefoundation.CFURLRef) int32 {
	result, callErr := tryPasteboardCopyPasteLocation(inPasteboard, outPasteLocation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardCreate func(inName corefoundation.CFStringRef, outPasteboard *PasteboardRef) int32
var _pasteboardCreateErr error

func tryPasteboardCreate(inName corefoundation.CFStringRef, outPasteboard *PasteboardRef) (int32, error) {
	if _pasteboardCreate == nil {
		return 0, symbolCallError("PasteboardCreate", "10.3", _pasteboardCreateErr)
	}
	return _pasteboardCreate(inName, outPasteboard), nil
}

// PasteboardCreate.
//
// See: https://developer.apple.com/documentation/applicationservices/1461248-pasteboardcreate
func PasteboardCreate(inName corefoundation.CFStringRef, outPasteboard *PasteboardRef) int32 {
	result, callErr := tryPasteboardCreate(inName, outPasteboard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardGetItemCount func(inPasteboard PasteboardRef, outItemCount unsafe.Pointer) int32
var _pasteboardGetItemCountErr error

func tryPasteboardGetItemCount(inPasteboard PasteboardRef, outItemCount unsafe.Pointer) (int32, error) {
	if _pasteboardGetItemCount == nil {
		return 0, symbolCallError("PasteboardGetItemCount", "10.3", _pasteboardGetItemCountErr)
	}
	return _pasteboardGetItemCount(inPasteboard, outItemCount), nil
}

// PasteboardGetItemCount.
//
// See: https://developer.apple.com/documentation/applicationservices/1459551-pasteboardgetitemcount
func PasteboardGetItemCount(inPasteboard PasteboardRef, outItemCount unsafe.Pointer) int32 {
	result, callErr := tryPasteboardGetItemCount(inPasteboard, outItemCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardGetItemFlavorFlags func(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, outFlags *PasteboardFlavorFlags) int32
var _pasteboardGetItemFlavorFlagsErr error

func tryPasteboardGetItemFlavorFlags(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, outFlags *PasteboardFlavorFlags) (int32, error) {
	if _pasteboardGetItemFlavorFlags == nil {
		return 0, symbolCallError("PasteboardGetItemFlavorFlags", "10.3", _pasteboardGetItemFlavorFlagsErr)
	}
	return _pasteboardGetItemFlavorFlags(inPasteboard, inItem, inFlavorType, outFlags), nil
}

// PasteboardGetItemFlavorFlags.
//
// See: https://developer.apple.com/documentation/applicationservices/1459353-pasteboardgetitemflavorflags
func PasteboardGetItemFlavorFlags(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, outFlags *PasteboardFlavorFlags) int32 {
	result, callErr := tryPasteboardGetItemFlavorFlags(inPasteboard, inItem, inFlavorType, outFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardGetItemIdentifier func(inPasteboard PasteboardRef, inIndex int, outItem *PasteboardItemID) int32
var _pasteboardGetItemIdentifierErr error

func tryPasteboardGetItemIdentifier(inPasteboard PasteboardRef, inIndex int, outItem *PasteboardItemID) (int32, error) {
	if _pasteboardGetItemIdentifier == nil {
		return 0, symbolCallError("PasteboardGetItemIdentifier", "10.3", _pasteboardGetItemIdentifierErr)
	}
	return _pasteboardGetItemIdentifier(inPasteboard, inIndex, outItem), nil
}

// PasteboardGetItemIdentifier.
//
// See: https://developer.apple.com/documentation/applicationservices/1463412-pasteboardgetitemidentifier
func PasteboardGetItemIdentifier(inPasteboard PasteboardRef, inIndex int, outItem *PasteboardItemID) int32 {
	result, callErr := tryPasteboardGetItemIdentifier(inPasteboard, inIndex, outItem)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardGetTypeID func() uint
var _pasteboardGetTypeIDErr error

func tryPasteboardGetTypeID() (uint, error) {
	if _pasteboardGetTypeID == nil {
		return 0, symbolCallError("PasteboardGetTypeID", "10.3", _pasteboardGetTypeIDErr)
	}
	return _pasteboardGetTypeID(), nil
}

// PasteboardGetTypeID.
//
// See: https://developer.apple.com/documentation/applicationservices/1463386-pasteboardgettypeid
func PasteboardGetTypeID() uint {
	result, callErr := tryPasteboardGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardPutItemFlavor func(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, inData corefoundation.CFDataRef, inFlags PasteboardFlavorFlags) int32
var _pasteboardPutItemFlavorErr error

func tryPasteboardPutItemFlavor(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, inData corefoundation.CFDataRef, inFlags PasteboardFlavorFlags) (int32, error) {
	if _pasteboardPutItemFlavor == nil {
		return 0, symbolCallError("PasteboardPutItemFlavor", "10.3", _pasteboardPutItemFlavorErr)
	}
	return _pasteboardPutItemFlavor(inPasteboard, inItem, inFlavorType, inData, inFlags), nil
}

// PasteboardPutItemFlavor.
//
// See: https://developer.apple.com/documentation/applicationservices/1463184-pasteboardputitemflavor
func PasteboardPutItemFlavor(inPasteboard PasteboardRef, inItem PasteboardItemID, inFlavorType corefoundation.CFStringRef, inData corefoundation.CFDataRef, inFlags PasteboardFlavorFlags) int32 {
	result, callErr := tryPasteboardPutItemFlavor(inPasteboard, inItem, inFlavorType, inData, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardResolvePromises func(inPasteboard PasteboardRef) int32
var _pasteboardResolvePromisesErr error

func tryPasteboardResolvePromises(inPasteboard PasteboardRef) (int32, error) {
	if _pasteboardResolvePromises == nil {
		return 0, symbolCallError("PasteboardResolvePromises", "10.3", _pasteboardResolvePromisesErr)
	}
	return _pasteboardResolvePromises(inPasteboard), nil
}

// PasteboardResolvePromises.
//
// See: https://developer.apple.com/documentation/applicationservices/1460816-pasteboardresolvepromises
func PasteboardResolvePromises(inPasteboard PasteboardRef) int32 {
	result, callErr := tryPasteboardResolvePromises(inPasteboard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardSetPasteLocation func(inPasteboard PasteboardRef, inPasteLocation corefoundation.CFURLRef) int32
var _pasteboardSetPasteLocationErr error

func tryPasteboardSetPasteLocation(inPasteboard PasteboardRef, inPasteLocation corefoundation.CFURLRef) (int32, error) {
	if _pasteboardSetPasteLocation == nil {
		return 0, symbolCallError("PasteboardSetPasteLocation", "10.3", _pasteboardSetPasteLocationErr)
	}
	return _pasteboardSetPasteLocation(inPasteboard, inPasteLocation), nil
}

// PasteboardSetPasteLocation.
//
// See: https://developer.apple.com/documentation/applicationservices/1460572-pasteboardsetpastelocation
func PasteboardSetPasteLocation(inPasteboard PasteboardRef, inPasteLocation corefoundation.CFURLRef) int32 {
	result, callErr := tryPasteboardSetPasteLocation(inPasteboard, inPasteLocation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardSetPromiseKeeper func(inPasteboard PasteboardRef, inPromiseKeeper unsafe.Pointer, inContext unsafe.Pointer) int32
var _pasteboardSetPromiseKeeperErr error

func tryPasteboardSetPromiseKeeper(inPasteboard PasteboardRef, inPromiseKeeper PasteboardPromiseKeeperProcPtr, inContext unsafe.Pointer) (int32, error) {
	if _pasteboardSetPromiseKeeper == nil {
		return 0, symbolCallError("PasteboardSetPromiseKeeper", "10.3", _pasteboardSetPromiseKeeperErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 PasteboardRef, blockArg1 PasteboardItemID, blockArg2 corefoundation.CFStringRef, blockArg3 unsafe.Pointer) int32 {
		return inPromiseKeeper(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _pasteboardSetPromiseKeeper(inPasteboard, _block0, inContext), nil
}

// PasteboardSetPromiseKeeper.
//
// See: https://developer.apple.com/documentation/applicationservices/1463604-pasteboardsetpromisekeeper
func PasteboardSetPromiseKeeper(inPasteboard PasteboardRef, inPromiseKeeper PasteboardPromiseKeeperProcPtr, inContext unsafe.Pointer) int32 {
	result, callErr := tryPasteboardSetPromiseKeeper(inPasteboard, inPromiseKeeper, inContext)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pasteboardSynchronize func(inPasteboard PasteboardRef) PasteboardSyncFlags
var _pasteboardSynchronizeErr error

func tryPasteboardSynchronize(inPasteboard PasteboardRef) (PasteboardSyncFlags, error) {
	if _pasteboardSynchronize == nil {
		return *new(PasteboardSyncFlags), symbolCallError("PasteboardSynchronize", "10.3", _pasteboardSynchronizeErr)
	}
	return _pasteboardSynchronize(inPasteboard), nil
}

// PasteboardSynchronize.
//
// See: https://developer.apple.com/documentation/applicationservices/1459590-pasteboardsynchronize
func PasteboardSynchronize(inPasteboard PasteboardRef) PasteboardSyncFlags {
	result, callErr := tryPasteboardSynchronize(inPasteboard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _pauseSpeechAt func(chan_ *SpeechChannelRecord, whereToPause unsafe.Pointer) int16
var _pauseSpeechAtErr error

func tryPauseSpeechAt(chan_ *SpeechChannelRecord, whereToPause unsafe.Pointer) (int16, error) {
	if _pauseSpeechAt == nil {
		return 0, symbolCallError("PauseSpeechAt", "10.0", _pauseSpeechAtErr)
	}
	return _pauseSpeechAt(chan_, whereToPause), nil
}

// PauseSpeechAt pauses speech on a speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1461174-pausespeechat
func PauseSpeechAt(chan_ *SpeechChannelRecord, whereToPause unsafe.Pointer) int16 {
	result, callErr := tryPauseSpeechAt(chan_, whereToPause)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _plotIconRefInContext func(inContext coregraphics.CGContextRef, inRect unsafe.Pointer, inAlign IconAlignmentType, inTransform IconTransformType, inLabelColor unsafe.Pointer, inFlags PlotIconRefFlags, inIconRef unsafe.Pointer) int32
var _plotIconRefInContextErr error

func tryPlotIconRefInContext(inContext coregraphics.CGContextRef, inRect unsafe.Pointer, inAlign IconAlignmentType, inTransform IconTransformType, inLabelColor unsafe.Pointer, inFlags PlotIconRefFlags, inIconRef unsafe.Pointer) (int32, error) {
	if _plotIconRefInContext == nil {
		return 0, symbolCallError("PlotIconRefInContext", "10.1", _plotIconRefInContextErr)
	}
	return _plotIconRefInContext(inContext, inRect, inAlign, inTransform, inLabelColor, inFlags, inIconRef), nil
}

// PlotIconRefInContext.
//
// See: https://developer.apple.com/documentation/applicationservices/1463721-ploticonrefincontext
func PlotIconRefInContext(inContext coregraphics.CGContextRef, inRect unsafe.Pointer, inAlign IconAlignmentType, inTransform IconTransformType, inLabelColor unsafe.Pointer, inFlags PlotIconRefFlags, inIconRef unsafe.Pointer) int32 {
	result, callErr := tryPlotIconRefInContext(inContext, inRect, inAlign, inTransform, inLabelColor, inFlags, inIconRef)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _processInformationCopyDictionary func(arg0 *ProcessSerialNumber, arg1 uint32) corefoundation.CFDictionaryRef
var _processInformationCopyDictionaryErr error

func tryProcessInformationCopyDictionary(arg0 *ProcessSerialNumber, arg1 uint32) (corefoundation.CFDictionaryRef, error) {
	if _processInformationCopyDictionary == nil {
		return 0, symbolCallError("ProcessInformationCopyDictionary", "10.2", _processInformationCopyDictionaryErr)
	}
	return _processInformationCopyDictionary(arg0, arg1), nil
}

// ProcessInformationCopyDictionary.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501104-processinformationcopydictionary
func ProcessInformationCopyDictionary(arg0 *ProcessSerialNumber, arg1 uint32) corefoundation.CFDictionaryRef {
	result, callErr := tryProcessInformationCopyDictionary(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sameProcess func(arg0 *ProcessSerialNumber, arg1 *ProcessSerialNumber, arg2 bool) int16
var _sameProcessErr error

func trySameProcess(arg0 *ProcessSerialNumber, arg1 *ProcessSerialNumber, arg2 bool) (int16, error) {
	if _sameProcess == nil {
		return 0, symbolCallError("SameProcess", "10.0", _sameProcessErr)
	}
	return _sameProcess(arg0, arg1, arg2), nil
}

// SameProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501087-sameprocess
func SameProcess(arg0 *ProcessSerialNumber, arg1 *ProcessSerialNumber, arg2 bool) int16 {
	result, callErr := trySameProcess(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setFrontProcess func(arg0 *ProcessSerialNumber) int16
var _setFrontProcessErr error

func trySetFrontProcess(arg0 *ProcessSerialNumber) (int16, error) {
	if _setFrontProcess == nil {
		return 0, symbolCallError("SetFrontProcess", "10.0", _setFrontProcessErr)
	}
	return _setFrontProcess(arg0), nil
}

// SetFrontProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501042-setfrontprocess
func SetFrontProcess(arg0 *ProcessSerialNumber) int16 {
	result, callErr := trySetFrontProcess(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setFrontProcessWithOptions func(arg0 *ProcessSerialNumber, arg1 OptionBits) int32
var _setFrontProcessWithOptionsErr error

func trySetFrontProcessWithOptions(arg0 *ProcessSerialNumber, arg1 OptionBits) (int32, error) {
	if _setFrontProcessWithOptions == nil {
		return 0, symbolCallError("SetFrontProcessWithOptions", "10.2", _setFrontProcessWithOptionsErr)
	}
	return _setFrontProcessWithOptions(arg0, arg1), nil
}

// SetFrontProcessWithOptions.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501003-setfrontprocesswithoptions
func SetFrontProcessWithOptions(arg0 *ProcessSerialNumber, arg1 OptionBits) int32 {
	result, callErr := trySetFrontProcessWithOptions(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setIconFamilyData func(iconFamily uintptr, iconType uint32, h unsafe.Pointer) int16
var _setIconFamilyDataErr error

func trySetIconFamilyData(iconFamily uintptr, iconType uint32, h unsafe.Pointer) (int16, error) {
	if _setIconFamilyData == nil {
		return 0, symbolCallError("SetIconFamilyData", "10.0", _setIconFamilyDataErr)
	}
	return _setIconFamilyData(iconFamily, iconType, h), nil
}

// SetIconFamilyData.
//
// See: https://developer.apple.com/documentation/applicationservices/1462050-seticonfamilydata
func SetIconFamilyData(iconFamily uintptr, iconType uint32, h unsafe.Pointer) int16 {
	result, callErr := trySetIconFamilyData(iconFamily, iconType, h)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setSpeechInfo func(arg0 *SpeechChannelRecord, arg1 uint32) int16
var _setSpeechInfoErr error

func trySetSpeechInfo(arg0 *SpeechChannelRecord, arg1 uint32) (int16, error) {
	if _setSpeechInfo == nil {
		return 0, symbolCallError("SetSpeechInfo", "10.0", _setSpeechInfoErr)
	}
	return _setSpeechInfo(arg0, arg1), nil
}

// SetSpeechInfo changes a setting of a particular speech channel.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552223-setspeechinfo
func SetSpeechInfo(arg0 *SpeechChannelRecord, arg1 uint32) int16 {
	result, callErr := trySetSpeechInfo(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setSpeechPitch func(chan_ *SpeechChannelRecord, pitch int32) int16
var _setSpeechPitchErr error

func trySetSpeechPitch(chan_ *SpeechChannelRecord, pitch int32) (int16, error) {
	if _setSpeechPitch == nil {
		return 0, symbolCallError("SetSpeechPitch", "10.0", _setSpeechPitchErr)
	}
	return _setSpeechPitch(chan_, pitch), nil
}

// SetSpeechPitch sets the speech pitch on a designated speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1462674-setspeechpitch
func SetSpeechPitch(chan_ *SpeechChannelRecord, pitch int32) int16 {
	result, callErr := trySetSpeechPitch(chan_, pitch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setSpeechProperty func(chan_ *SpeechChannelRecord, property corefoundation.CFStringRef, object corefoundation.CFTypeRef) int16
var _setSpeechPropertyErr error

func trySetSpeechProperty(chan_ *SpeechChannelRecord, property corefoundation.CFStringRef, object corefoundation.CFTypeRef) (int16, error) {
	if _setSpeechProperty == nil {
		return 0, symbolCallError("SetSpeechProperty", "10.5", _setSpeechPropertyErr)
	}
	return _setSpeechProperty(chan_, property, object), nil
}

// SetSpeechProperty sets the value of the specified speech-channel property.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1459256-setspeechproperty
func SetSpeechProperty(chan_ *SpeechChannelRecord, property corefoundation.CFStringRef, object corefoundation.CFTypeRef) int16 {
	result, callErr := trySetSpeechProperty(chan_, property, object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _setSpeechRate func(chan_ *SpeechChannelRecord, rate int32) int16
var _setSpeechRateErr error

func trySetSpeechRate(chan_ *SpeechChannelRecord, rate int32) (int16, error) {
	if _setSpeechRate == nil {
		return 0, symbolCallError("SetSpeechRate", "10.0", _setSpeechRateErr)
	}
	return _setSpeechRate(chan_, rate), nil
}

// SetSpeechRate sets the speech rate of a designated speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1459896-setspeechrate
func SetSpeechRate(chan_ *SpeechChannelRecord, rate int32) int16 {
	result, callErr := trySetSpeechRate(chan_, rate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _showHideProcess func(arg0 *ProcessSerialNumber, arg1 bool) int16
var _showHideProcessErr error

func tryShowHideProcess(arg0 *ProcessSerialNumber, arg1 bool) (int16, error) {
	if _showHideProcess == nil {
		return 0, symbolCallError("ShowHideProcess", "10.1", _showHideProcessErr)
	}
	return _showHideProcess(arg0, arg1), nil
}

// ShowHideProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501053-showhideprocess
func ShowHideProcess(arg0 *ProcessSerialNumber, arg1 bool) int16 {
	result, callErr := tryShowHideProcess(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speakBuffer func(arg0 *SpeechChannelRecord, arg1 uint, arg2 int32) int16
var _speakBufferErr error

func trySpeakBuffer(arg0 *SpeechChannelRecord, arg1 uint, arg2 int32) (int16, error) {
	if _speakBuffer == nil {
		return 0, symbolCallError("SpeakBuffer", "10.0", _speakBufferErr)
	}
	return _speakBuffer(arg0, arg1, arg2), nil
}

// SpeakBuffer speaks a buffer of text, using certain flags to controlspeech behavior.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552252-speakbuffer
func SpeakBuffer(arg0 *SpeechChannelRecord, arg1 uint, arg2 int32) int16 {
	result, callErr := trySpeakBuffer(arg0, arg1, arg2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speakCFString func(chan_ *SpeechChannelRecord, aString corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) int16
var _speakCFStringErr error

func trySpeakCFString(chan_ *SpeechChannelRecord, aString corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) (int16, error) {
	if _speakCFString == nil {
		return 0, symbolCallError("SpeakCFString", "10.5", _speakCFStringErr)
	}
	return _speakCFString(chan_, aString, options), nil
}

// SpeakCFString begins speaking a string represented as a [CFString] object.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1461621-speakcfstring
func SpeakCFString(chan_ *SpeechChannelRecord, aString corefoundation.CFStringRef, options corefoundation.CFDictionaryRef) int16 {
	result, callErr := trySpeakCFString(chan_, aString, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speakString func(arg0 unsafe.Pointer) int16
var _speakStringErr error

func trySpeakString(arg0 unsafe.Pointer) (int16, error) {
	if _speakString == nil {
		return 0, symbolCallError("SpeakString", "10.0", _speakStringErr)
	}
	return _speakString(arg0), nil
}

// SpeakString begins speaking a text string.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552250-speakstring
func SpeakString(arg0 unsafe.Pointer) int16 {
	result, callErr := trySpeakString(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speakText func(arg0 *SpeechChannelRecord, arg1 uint) int16
var _speakTextErr error

func trySpeakText(arg0 *SpeechChannelRecord, arg1 uint) (int16, error) {
	if _speakText == nil {
		return 0, symbolCallError("SpeakText", "10.0", _speakTextErr)
	}
	return _speakText(arg0, arg1), nil
}

// SpeakText begins speaking a buffer of text.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552236-speaktext
func SpeakText(arg0 *SpeechChannelRecord, arg1 uint) int16 {
	result, callErr := trySpeakText(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speechBusy func() unsafe.Pointer
var _speechBusyErr error

func trySpeechBusy() (unsafe.Pointer, error) {
	if _speechBusy == nil {
		return nil, symbolCallError("SpeechBusy", "10.0", _speechBusyErr)
	}
	return _speechBusy(), nil
}

// SpeechBusy determines whether any channels of speech are currentlysynthesizing speech.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1464581-speechbusy
func SpeechBusy() unsafe.Pointer {
	result, callErr := trySpeechBusy()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speechBusySystemWide func() unsafe.Pointer
var _speechBusySystemWideErr error

func trySpeechBusySystemWide() (unsafe.Pointer, error) {
	if _speechBusySystemWide == nil {
		return nil, symbolCallError("SpeechBusySystemWide", "10.0", _speechBusySystemWideErr)
	}
	return _speechBusySystemWide(), nil
}

// SpeechBusySystemWide determines if any speech is currently being synthesizedin your application or elsewhere on the computer.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1460113-speechbusysystemwide
func SpeechBusySystemWide() unsafe.Pointer {
	result, callErr := trySpeechBusySystemWide()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speechManagerVersion func() uint32
var _speechManagerVersionErr error

func trySpeechManagerVersion() (uint32, error) {
	if _speechManagerVersion == nil {
		return 0, symbolCallError("SpeechManagerVersion", "10.0", _speechManagerVersionErr)
	}
	return _speechManagerVersion(), nil
}

// SpeechManagerVersion determines the current version of the Speech SynthesisManager installed in the system.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1462334-speechmanagerversion
func SpeechManagerVersion() uint32 {
	result, callErr := trySpeechManagerVersion()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speechSynthesisRegisterModuleURL func(url corefoundation.CFURLRef) int16
var _speechSynthesisRegisterModuleURLErr error

func trySpeechSynthesisRegisterModuleURL(url corefoundation.CFURLRef) (int16, error) {
	if _speechSynthesisRegisterModuleURL == nil {
		return 0, symbolCallError("SpeechSynthesisRegisterModuleURL", "10.6", _speechSynthesisRegisterModuleURLErr)
	}
	return _speechSynthesisRegisterModuleURL(url), nil
}

// SpeechSynthesisRegisterModuleURL registers and makes available a speech synthesizer or voice.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1459624-speechsynthesisregistermoduleurl
func SpeechSynthesisRegisterModuleURL(url corefoundation.CFURLRef) int16 {
	result, callErr := trySpeechSynthesisRegisterModuleURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _speechSynthesisUnregisterModuleURL func(url corefoundation.CFURLRef) int16
var _speechSynthesisUnregisterModuleURLErr error

func trySpeechSynthesisUnregisterModuleURL(url corefoundation.CFURLRef) (int16, error) {
	if _speechSynthesisUnregisterModuleURL == nil {
		return 0, symbolCallError("SpeechSynthesisUnregisterModuleURL", "10.6", _speechSynthesisUnregisterModuleURLErr)
	}
	return _speechSynthesisUnregisterModuleURL(url), nil
}

// SpeechSynthesisUnregisterModuleURL unregisters a registered speech synthesizer or voice.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1462511-speechsynthesisunregistermoduleu
func SpeechSynthesisUnregisterModuleURL(url corefoundation.CFURLRef) int16 {
	result, callErr := trySpeechSynthesisUnregisterModuleURL(url)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _stopSpeech func(chan_ *SpeechChannelRecord) int16
var _stopSpeechErr error

func tryStopSpeech(chan_ *SpeechChannelRecord) (int16, error) {
	if _stopSpeech == nil {
		return 0, symbolCallError("StopSpeech", "10.0", _stopSpeechErr)
	}
	return _stopSpeech(chan_), nil
}

// StopSpeech terminates speech immediately on the specified channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1462745-stopspeech
func StopSpeech(chan_ *SpeechChannelRecord) int16 {
	result, callErr := tryStopSpeech(chan_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _stopSpeechAt func(chan_ *SpeechChannelRecord, whereToStop unsafe.Pointer) int16
var _stopSpeechAtErr error

func tryStopSpeechAt(chan_ *SpeechChannelRecord, whereToStop unsafe.Pointer) (int16, error) {
	if _stopSpeechAt == nil {
		return 0, symbolCallError("StopSpeechAt", "10.0", _stopSpeechAtErr)
	}
	return _stopSpeechAt(chan_, whereToStop), nil
}

// StopSpeechAt terminates speech delivery on a specified channel eitherimmediately or at the end of the current word or sentence.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1459780-stopspeechat
func StopSpeechAt(chan_ *SpeechChannelRecord, whereToStop unsafe.Pointer) int16 {
	result, callErr := tryStopSpeechAt(chan_, whereToStop)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _textToPhonemes func(arg0 *SpeechChannelRecord, arg1 uint, arg2 unsafe.Pointer, arg3 int) int16
var _textToPhonemesErr error

func tryTextToPhonemes(arg0 *SpeechChannelRecord, arg1 uint, arg2 unsafe.Pointer, arg3 int) (int16, error) {
	if _textToPhonemes == nil {
		return 0, symbolCallError("TextToPhonemes", "10.0", _textToPhonemesErr)
	}
	return _textToPhonemes(arg0, arg1, arg2, arg3), nil
}

// TextToPhonemes converts a buffer of textual data into phonemic data.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552235-texttophonemes
func TextToPhonemes(arg0 *SpeechChannelRecord, arg1 uint, arg2 unsafe.Pointer, arg3 int) int16 {
	result, callErr := tryTextToPhonemes(arg0, arg1, arg2, arg3)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _transformProcessType func(psn unsafe.Pointer, transformState ProcessApplicationTransformState) int32
var _transformProcessTypeErr error

func tryTransformProcessType(psn unsafe.Pointer, transformState ProcessApplicationTransformState) (int32, error) {
	if _transformProcessType == nil {
		return 0, symbolCallError("TransformProcessType", "10.3", _transformProcessTypeErr)
	}
	return _transformProcessType(psn, transformState), nil
}

// TransformProcessType.
//
// See: https://developer.apple.com/documentation/applicationservices/1462420-transformprocesstype
func TransformProcessType(psn unsafe.Pointer, transformState ProcessApplicationTransformState) int32 {
	result, callErr := tryTransformProcessType(psn, transformState)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationCopyDestinationType func(inTranslation TranslationRef, outDestinationType *corefoundation.CFStringRef) int32
var _translationCopyDestinationTypeErr error

func tryTranslationCopyDestinationType(inTranslation TranslationRef, outDestinationType *corefoundation.CFStringRef) (int32, error) {
	if _translationCopyDestinationType == nil {
		return 0, symbolCallError("TranslationCopyDestinationType", "10.3", _translationCopyDestinationTypeErr)
	}
	return _translationCopyDestinationType(inTranslation, outDestinationType), nil
}

// TranslationCopyDestinationType.
//
// See: https://developer.apple.com/documentation/applicationservices/1459620-translationcopydestinationtype
func TranslationCopyDestinationType(inTranslation TranslationRef, outDestinationType *corefoundation.CFStringRef) int32 {
	result, callErr := tryTranslationCopyDestinationType(inTranslation, outDestinationType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationCopySourceType func(inTranslation TranslationRef, outSourceType *corefoundation.CFStringRef) int32
var _translationCopySourceTypeErr error

func tryTranslationCopySourceType(inTranslation TranslationRef, outSourceType *corefoundation.CFStringRef) (int32, error) {
	if _translationCopySourceType == nil {
		return 0, symbolCallError("TranslationCopySourceType", "10.3", _translationCopySourceTypeErr)
	}
	return _translationCopySourceType(inTranslation, outSourceType), nil
}

// TranslationCopySourceType.
//
// See: https://developer.apple.com/documentation/applicationservices/1459344-translationcopysourcetype
func TranslationCopySourceType(inTranslation TranslationRef, outSourceType *corefoundation.CFStringRef) int32 {
	result, callErr := tryTranslationCopySourceType(inTranslation, outSourceType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationCreate func(inSourceType corefoundation.CFStringRef, inDestinationType corefoundation.CFStringRef, inTranslationFlags TranslationFlags, outTranslation *TranslationRef) int32
var _translationCreateErr error

func tryTranslationCreate(inSourceType corefoundation.CFStringRef, inDestinationType corefoundation.CFStringRef, inTranslationFlags TranslationFlags, outTranslation *TranslationRef) (int32, error) {
	if _translationCreate == nil {
		return 0, symbolCallError("TranslationCreate", "10.3", _translationCreateErr)
	}
	return _translationCreate(inSourceType, inDestinationType, inTranslationFlags, outTranslation), nil
}

// TranslationCreate.
//
// See: https://developer.apple.com/documentation/applicationservices/1459231-translationcreate
func TranslationCreate(inSourceType corefoundation.CFStringRef, inDestinationType corefoundation.CFStringRef, inTranslationFlags TranslationFlags, outTranslation *TranslationRef) int32 {
	result, callErr := tryTranslationCreate(inSourceType, inDestinationType, inTranslationFlags, outTranslation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationCreateWithSourceArray func(inSourceTypes corefoundation.CFArrayRef, inTranslationFlags TranslationFlags, outDestinationTypes *corefoundation.CFArrayRef, outTranslations *corefoundation.CFDictionaryRef) int32
var _translationCreateWithSourceArrayErr error

func tryTranslationCreateWithSourceArray(inSourceTypes corefoundation.CFArrayRef, inTranslationFlags TranslationFlags, outDestinationTypes *corefoundation.CFArrayRef, outTranslations *corefoundation.CFDictionaryRef) (int32, error) {
	if _translationCreateWithSourceArray == nil {
		return 0, symbolCallError("TranslationCreateWithSourceArray", "10.3", _translationCreateWithSourceArrayErr)
	}
	return _translationCreateWithSourceArray(inSourceTypes, inTranslationFlags, outDestinationTypes, outTranslations), nil
}

// TranslationCreateWithSourceArray.
//
// See: https://developer.apple.com/documentation/applicationservices/1464306-translationcreatewithsourcearray
func TranslationCreateWithSourceArray(inSourceTypes corefoundation.CFArrayRef, inTranslationFlags TranslationFlags, outDestinationTypes *corefoundation.CFArrayRef, outTranslations *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryTranslationCreateWithSourceArray(inSourceTypes, inTranslationFlags, outDestinationTypes, outTranslations)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationGetTranslationFlags func(inTranslation TranslationRef, outTranslationFlags *TranslationFlags) int32
var _translationGetTranslationFlagsErr error

func tryTranslationGetTranslationFlags(inTranslation TranslationRef, outTranslationFlags *TranslationFlags) (int32, error) {
	if _translationGetTranslationFlags == nil {
		return 0, symbolCallError("TranslationGetTranslationFlags", "10.3", _translationGetTranslationFlagsErr)
	}
	return _translationGetTranslationFlags(inTranslation, outTranslationFlags), nil
}

// TranslationGetTranslationFlags.
//
// See: https://developer.apple.com/documentation/applicationservices/1459307-translationgettranslationflags
func TranslationGetTranslationFlags(inTranslation TranslationRef, outTranslationFlags *TranslationFlags) int32 {
	result, callErr := tryTranslationGetTranslationFlags(inTranslation, outTranslationFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationGetTypeID func() uint
var _translationGetTypeIDErr error

func tryTranslationGetTypeID() (uint, error) {
	if _translationGetTypeID == nil {
		return 0, symbolCallError("TranslationGetTypeID", "10.3", _translationGetTypeIDErr)
	}
	return _translationGetTypeID(), nil
}

// TranslationGetTypeID.
//
// See: https://developer.apple.com/documentation/applicationservices/1463809-translationgettypeid
func TranslationGetTypeID() uint {
	result, callErr := tryTranslationGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationPerformForData func(inTranslation TranslationRef, inSourceData corefoundation.CFDataRef, outDestinationData *corefoundation.CFDataRef) int32
var _translationPerformForDataErr error

func tryTranslationPerformForData(inTranslation TranslationRef, inSourceData corefoundation.CFDataRef, outDestinationData *corefoundation.CFDataRef) (int32, error) {
	if _translationPerformForData == nil {
		return 0, symbolCallError("TranslationPerformForData", "10.3", _translationPerformForDataErr)
	}
	return _translationPerformForData(inTranslation, inSourceData, outDestinationData), nil
}

// TranslationPerformForData.
//
// See: https://developer.apple.com/documentation/applicationservices/1460828-translationperformfordata
func TranslationPerformForData(inTranslation TranslationRef, inSourceData corefoundation.CFDataRef, outDestinationData *corefoundation.CFDataRef) int32 {
	result, callErr := tryTranslationPerformForData(inTranslation, inSourceData, outDestinationData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationPerformForFile func(inTranslation TranslationRef, inSourceFile unsafe.Pointer, inDestinationDirectory unsafe.Pointer, inDestinationName corefoundation.CFStringRef, outTranslatedFile unsafe.Pointer) int32
var _translationPerformForFileErr error

func tryTranslationPerformForFile(inTranslation TranslationRef, inSourceFile unsafe.Pointer, inDestinationDirectory unsafe.Pointer, inDestinationName corefoundation.CFStringRef, outTranslatedFile unsafe.Pointer) (int32, error) {
	if _translationPerformForFile == nil {
		return 0, symbolCallError("TranslationPerformForFile", "10.3", _translationPerformForFileErr)
	}
	return _translationPerformForFile(inTranslation, inSourceFile, inDestinationDirectory, inDestinationName, outTranslatedFile), nil
}

// TranslationPerformForFile.
//
// See: https://developer.apple.com/documentation/applicationservices/1464541-translationperformforfile
func TranslationPerformForFile(inTranslation TranslationRef, inSourceFile unsafe.Pointer, inDestinationDirectory unsafe.Pointer, inDestinationName corefoundation.CFStringRef, outTranslatedFile unsafe.Pointer) int32 {
	result, callErr := tryTranslationPerformForFile(inTranslation, inSourceFile, inDestinationDirectory, inDestinationName, outTranslatedFile)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _translationPerformForURL func(inTranslation TranslationRef, inSourceURL corefoundation.CFURLRef, inDestinationURL corefoundation.CFURLRef, outTranslatedURL *corefoundation.CFURLRef) int32
var _translationPerformForURLErr error

func tryTranslationPerformForURL(inTranslation TranslationRef, inSourceURL corefoundation.CFURLRef, inDestinationURL corefoundation.CFURLRef, outTranslatedURL *corefoundation.CFURLRef) (int32, error) {
	if _translationPerformForURL == nil {
		return 0, symbolCallError("TranslationPerformForURL", "10.3", _translationPerformForURLErr)
	}
	return _translationPerformForURL(inTranslation, inSourceURL, inDestinationURL, outTranslatedURL), nil
}

// TranslationPerformForURL.
//
// See: https://developer.apple.com/documentation/applicationservices/1460118-translationperformforurl
func TranslationPerformForURL(inTranslation TranslationRef, inSourceURL corefoundation.CFURLRef, inDestinationURL corefoundation.CFURLRef, outTranslatedURL *corefoundation.CFURLRef) int32 {
	result, callErr := tryTranslationPerformForURL(inTranslation, inSourceURL, inDestinationURL, outTranslatedURL)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _uAZoomChangeFocus func(inRect unsafe.Pointer, inHighlightRect unsafe.Pointer, inType UAZoomChangeFocusType) int32
var _uAZoomChangeFocusErr error

func tryUAZoomChangeFocus(inRect unsafe.Pointer, inHighlightRect unsafe.Pointer, inType UAZoomChangeFocusType) (int32, error) {
	if _uAZoomChangeFocus == nil {
		return 0, symbolCallError("UAZoomChangeFocus", "10.4", _uAZoomChangeFocusErr)
	}
	return _uAZoomChangeFocus(inRect, inHighlightRect, inType), nil
}

// UAZoomChangeFocus tells the Universal Access zoom feature where it should focus.
//
// See: https://developer.apple.com/documentation/applicationservices/1458830-uazoomchangefocus
func UAZoomChangeFocus(inRect unsafe.Pointer, inHighlightRect unsafe.Pointer, inType UAZoomChangeFocusType) int32 {
	result, callErr := tryUAZoomChangeFocus(inRect, inHighlightRect, inType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _uAZoomEnabled func() bool
var _uAZoomEnabledErr error

func tryUAZoomEnabled() (bool, error) {
	if _uAZoomEnabled == nil {
		return false, symbolCallError("UAZoomEnabled", "10.4", _uAZoomEnabledErr)
	}
	return _uAZoomEnabled(), nil
}

// UAZoomEnabled determines if the Universal Access zoom feature is enabled.
//
// See: https://developer.apple.com/documentation/applicationservices/1462288-uazoomenabled
func UAZoomEnabled() bool {
	result, callErr := tryUAZoomEnabled()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _useDictionary func(arg0 *SpeechChannelRecord, arg1 unsafe.Pointer) int16
var _useDictionaryErr error

func tryUseDictionary(arg0 *SpeechChannelRecord, arg1 unsafe.Pointer) (int16, error) {
	if _useDictionary == nil {
		return 0, symbolCallError("UseDictionary", "10.0", _useDictionaryErr)
	}
	return _useDictionary(arg0, arg1), nil
}

// UseDictionary installs the designated dictionary into a speech channel.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/applicationservices/1552255-usedictionary
func UseDictionary(arg0 *SpeechChannelRecord, arg1 unsafe.Pointer) int16 {
	result, callErr := tryUseDictionary(arg0, arg1)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _useSpeechDictionary func(chan_ *SpeechChannelRecord, speechDictionary corefoundation.CFDictionaryRef) int16
var _useSpeechDictionaryErr error

func tryUseSpeechDictionary(chan_ *SpeechChannelRecord, speechDictionary corefoundation.CFDictionaryRef) (int16, error) {
	if _useSpeechDictionary == nil {
		return 0, symbolCallError("UseSpeechDictionary", "10.5", _useSpeechDictionaryErr)
	}
	return _useSpeechDictionary(chan_, speechDictionary), nil
}

// UseSpeechDictionary registers a speech dictionary with a speech channel.
//
// Deprecated: Deprecated since macOS 13.0.
//
// See: https://developer.apple.com/documentation/applicationservices/1463688-usespeechdictionary
func UseSpeechDictionary(chan_ *SpeechChannelRecord, speechDictionary corefoundation.CFDictionaryRef) int16 {
	result, callErr := tryUseSpeechDictionary(chan_, speechDictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _wakeUpProcess func(arg0 *ProcessSerialNumber) int16
var _wakeUpProcessErr error

func tryWakeUpProcess(arg0 *ProcessSerialNumber) (int16, error) {
	if _wakeUpProcess == nil {
		return 0, symbolCallError("WakeUpProcess", "10.0", _wakeUpProcessErr)
	}
	return _wakeUpProcess(arg0), nil
}

// WakeUpProcess.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/applicationservices/1501091-wakeupprocess
func WakeUpProcess(arg0 *ProcessSerialNumber) int16 {
	result, callErr := tryWakeUpProcess(arg0)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_aXAPIEnabled, &_aXAPIEnabledErr, frameworkHandle, "AXAPIEnabled", "10.0")
	registerFunc(&_aXIsProcessTrusted, &_aXIsProcessTrustedErr, frameworkHandle, "AXIsProcessTrusted", "10.4")
	registerFunc(&_aXIsProcessTrustedWithOptions, &_aXIsProcessTrustedWithOptionsErr, frameworkHandle, "AXIsProcessTrustedWithOptions", "10.9")
	registerFunc(&_aXMakeProcessTrusted, &_aXMakeProcessTrustedErr, frameworkHandle, "AXMakeProcessTrusted", "10.4")
	registerFunc(&_aXObserverAddNotification, &_aXObserverAddNotificationErr, frameworkHandle, "AXObserverAddNotification", "10.2")
	registerFunc(&_aXObserverCreate, &_aXObserverCreateErr, frameworkHandle, "AXObserverCreate", "10.2")
	registerFunc(&_aXObserverCreateWithInfoCallback, &_aXObserverCreateWithInfoCallbackErr, frameworkHandle, "AXObserverCreateWithInfoCallback", "10.9")
	registerFunc(&_aXObserverGetRunLoopSource, &_aXObserverGetRunLoopSourceErr, frameworkHandle, "AXObserverGetRunLoopSource", "10.2")
	registerFunc(&_aXObserverGetTypeID, &_aXObserverGetTypeIDErr, frameworkHandle, "AXObserverGetTypeID", "10.2")
	registerFunc(&_aXObserverRemoveNotification, &_aXObserverRemoveNotificationErr, frameworkHandle, "AXObserverRemoveNotification", "10.2")
	registerFunc(&_aXTextMarkerCreate, &_aXTextMarkerCreateErr, frameworkHandle, "AXTextMarkerCreate", "12.0")
	registerFunc(&_aXTextMarkerGetBytePtr, &_aXTextMarkerGetBytePtrErr, frameworkHandle, "AXTextMarkerGetBytePtr", "12.0")
	registerFunc(&_aXTextMarkerGetLength, &_aXTextMarkerGetLengthErr, frameworkHandle, "AXTextMarkerGetLength", "12.0")
	registerFunc(&_aXTextMarkerGetTypeID, &_aXTextMarkerGetTypeIDErr, frameworkHandle, "AXTextMarkerGetTypeID", "12.0")
	registerFunc(&_aXTextMarkerRangeCopyEndMarker, &_aXTextMarkerRangeCopyEndMarkerErr, frameworkHandle, "AXTextMarkerRangeCopyEndMarker", "12.0")
	registerFunc(&_aXTextMarkerRangeCopyStartMarker, &_aXTextMarkerRangeCopyStartMarkerErr, frameworkHandle, "AXTextMarkerRangeCopyStartMarker", "12.0")
	registerFunc(&_aXTextMarkerRangeCreate, &_aXTextMarkerRangeCreateErr, frameworkHandle, "AXTextMarkerRangeCreate", "12.0")
	registerFunc(&_aXTextMarkerRangeCreateWithBytes, &_aXTextMarkerRangeCreateWithBytesErr, frameworkHandle, "AXTextMarkerRangeCreateWithBytes", "12.0")
	registerFunc(&_aXTextMarkerRangeGetTypeID, &_aXTextMarkerRangeGetTypeIDErr, frameworkHandle, "AXTextMarkerRangeGetTypeID", "12.0")
	registerFunc(&_aXUIElementCopyActionDescription, &_aXUIElementCopyActionDescriptionErr, frameworkHandle, "AXUIElementCopyActionDescription", "10.2")
	registerFunc(&_aXUIElementCopyActionNames, &_aXUIElementCopyActionNamesErr, frameworkHandle, "AXUIElementCopyActionNames", "10.2")
	registerFunc(&_aXUIElementCopyAttributeNames, &_aXUIElementCopyAttributeNamesErr, frameworkHandle, "AXUIElementCopyAttributeNames", "10.2")
	registerFunc(&_aXUIElementCopyAttributeValue, &_aXUIElementCopyAttributeValueErr, frameworkHandle, "AXUIElementCopyAttributeValue", "10.2")
	registerFunc(&_aXUIElementCopyAttributeValues, &_aXUIElementCopyAttributeValuesErr, frameworkHandle, "AXUIElementCopyAttributeValues", "10.2")
	registerFunc(&_aXUIElementCopyElementAtPosition, &_aXUIElementCopyElementAtPositionErr, frameworkHandle, "AXUIElementCopyElementAtPosition", "10.2")
	registerFunc(&_aXUIElementCopyMultipleAttributeValues, &_aXUIElementCopyMultipleAttributeValuesErr, frameworkHandle, "AXUIElementCopyMultipleAttributeValues", "10.4")
	registerFunc(&_aXUIElementCopyParameterizedAttributeNames, &_aXUIElementCopyParameterizedAttributeNamesErr, frameworkHandle, "AXUIElementCopyParameterizedAttributeNames", "10.3")
	registerFunc(&_aXUIElementCopyParameterizedAttributeValue, &_aXUIElementCopyParameterizedAttributeValueErr, frameworkHandle, "AXUIElementCopyParameterizedAttributeValue", "10.3")
	registerFunc(&_aXUIElementCreateApplication, &_aXUIElementCreateApplicationErr, frameworkHandle, "AXUIElementCreateApplication", "10.2")
	registerFunc(&_aXUIElementCreateSystemWide, &_aXUIElementCreateSystemWideErr, frameworkHandle, "AXUIElementCreateSystemWide", "10.2")
	registerFunc(&_aXUIElementGetAttributeValueCount, &_aXUIElementGetAttributeValueCountErr, frameworkHandle, "AXUIElementGetAttributeValueCount", "10.2")
	registerFunc(&_aXUIElementGetPid, &_aXUIElementGetPidErr, frameworkHandle, "AXUIElementGetPid", "10.2")
	registerFunc(&_aXUIElementGetTypeID, &_aXUIElementGetTypeIDErr, frameworkHandle, "AXUIElementGetTypeID", "10.2")
	registerFunc(&_aXUIElementIsAttributeSettable, &_aXUIElementIsAttributeSettableErr, frameworkHandle, "AXUIElementIsAttributeSettable", "10.2")
	registerFunc(&_aXUIElementPerformAction, &_aXUIElementPerformActionErr, frameworkHandle, "AXUIElementPerformAction", "10.2")
	registerFunc(&_aXUIElementPostKeyboardEvent, &_aXUIElementPostKeyboardEventErr, frameworkHandle, "AXUIElementPostKeyboardEvent", "10.0")
	registerFunc(&_aXUIElementSetAttributeValue, &_aXUIElementSetAttributeValueErr, frameworkHandle, "AXUIElementSetAttributeValue", "10.2")
	registerFunc(&_aXUIElementSetMessagingTimeout, &_aXUIElementSetMessagingTimeoutErr, frameworkHandle, "AXUIElementSetMessagingTimeout", "10.4")
	registerFunc(&_aXValueCreate, &_aXValueCreateErr, frameworkHandle, "AXValueCreate", "10.2")
	registerFunc(&_aXValueGetType, &_aXValueGetTypeErr, frameworkHandle, "AXValueGetType", "10.2")
	registerFunc(&_aXValueGetTypeID, &_aXValueGetTypeIDErr, frameworkHandle, "AXValueGetTypeID", "10.3")
	registerFunc(&_aXValueGetValue, &_aXValueGetValueErr, frameworkHandle, "AXValueGetValue", "10.2")
	registerFunc(&_continueSpeech, &_continueSpeechErr, frameworkHandle, "ContinueSpeech", "10.0")
	registerFunc(&_copyPhonemesFromText, &_copyPhonemesFromTextErr, frameworkHandle, "CopyPhonemesFromText", "10.5")
	registerFunc(&_copyProcessName, &_copyProcessNameErr, frameworkHandle, "CopyProcessName", "10.0")
	registerFunc(&_copySpeechProperty, &_copySpeechPropertyErr, frameworkHandle, "CopySpeechProperty", "10.5")
	registerFunc(&_countVoices, &_countVoicesErr, frameworkHandle, "CountVoices", "10.0")
	registerFunc(&_disposeIconActionUPP, &_disposeIconActionUPPErr, frameworkHandle, "DisposeIconActionUPP", "10.0")
	registerFunc(&_disposeIconGetterUPP, &_disposeIconGetterUPPErr, frameworkHandle, "DisposeIconGetterUPP", "10.0")
	registerFunc(&_disposeSpeechChannel, &_disposeSpeechChannelErr, frameworkHandle, "DisposeSpeechChannel", "10.0")
	registerFunc(&_disposeSpeechDoneUPP, &_disposeSpeechDoneUPPErr, frameworkHandle, "DisposeSpeechDoneUPP", "10.0")
	registerFunc(&_disposeSpeechErrorUPP, &_disposeSpeechErrorUPPErr, frameworkHandle, "DisposeSpeechErrorUPP", "10.0")
	registerFunc(&_disposeSpeechPhonemeUPP, &_disposeSpeechPhonemeUPPErr, frameworkHandle, "DisposeSpeechPhonemeUPP", "10.0")
	registerFunc(&_disposeSpeechSyncUPP, &_disposeSpeechSyncUPPErr, frameworkHandle, "DisposeSpeechSyncUPP", "10.0")
	registerFunc(&_disposeSpeechTextDoneUPP, &_disposeSpeechTextDoneUPPErr, frameworkHandle, "DisposeSpeechTextDoneUPP", "10.0")
	registerFunc(&_disposeSpeechWordUPP, &_disposeSpeechWordUPPErr, frameworkHandle, "DisposeSpeechWordUPP", "10.0")
	registerFunc(&_exitToShell, &_exitToShellErr, frameworkHandle, "ExitToShell", "10.0")
	registerFunc(&_getCurrentProcess, &_getCurrentProcessErr, frameworkHandle, "GetCurrentProcess", "10.0")
	registerFunc(&_getFrontProcess, &_getFrontProcessErr, frameworkHandle, "GetFrontProcess", "10.0")
	registerFunc(&_getIconFamilyData, &_getIconFamilyDataErr, frameworkHandle, "GetIconFamilyData", "10.0")
	registerFunc(&_getIconRefVariant, &_getIconRefVariantErr, frameworkHandle, "GetIconRefVariant", "10.0")
	registerFunc(&_getIndVoice, &_getIndVoiceErr, frameworkHandle, "GetIndVoice", "10.0")
	registerFunc(&_getNextProcess, &_getNextProcessErr, frameworkHandle, "GetNextProcess", "10.0")
	registerFunc(&_getProcessBundleLocation, &_getProcessBundleLocationErr, frameworkHandle, "GetProcessBundleLocation", "10.0")
	registerFunc(&_getProcessForPID, &_getProcessForPIDErr, frameworkHandle, "GetProcessForPID", "10.0")
	registerFunc(&_getProcessInformation, &_getProcessInformationErr, frameworkHandle, "GetProcessInformation", "10.0")
	registerFunc(&_getProcessPID, &_getProcessPIDErr, frameworkHandle, "GetProcessPID", "10.0")
	registerFunc(&_getSpeechInfo, &_getSpeechInfoErr, frameworkHandle, "GetSpeechInfo", "10.0")
	registerFunc(&_getSpeechPitch, &_getSpeechPitchErr, frameworkHandle, "GetSpeechPitch", "10.0")
	registerFunc(&_getSpeechRate, &_getSpeechRateErr, frameworkHandle, "GetSpeechRate", "10.0")
	registerFunc(&_getVoiceDescription, &_getVoiceDescriptionErr, frameworkHandle, "GetVoiceDescription", "10.0")
	registerFunc(&_getVoiceInfo, &_getVoiceInfoErr, frameworkHandle, "GetVoiceInfo", "10.0")
	registerFunc(&_hIShapeContainsPoint, &_hIShapeContainsPointErr, frameworkHandle, "HIShapeContainsPoint", "10.2")
	registerFunc(&_hIShapeCreateCopy, &_hIShapeCreateCopyErr, frameworkHandle, "HIShapeCreateCopy", "10.2")
	registerFunc(&_hIShapeCreateDifference, &_hIShapeCreateDifferenceErr, frameworkHandle, "HIShapeCreateDifference", "10.2")
	registerFunc(&_hIShapeCreateEmpty, &_hIShapeCreateEmptyErr, frameworkHandle, "HIShapeCreateEmpty", "10.4")
	registerFunc(&_hIShapeCreateIntersection, &_hIShapeCreateIntersectionErr, frameworkHandle, "HIShapeCreateIntersection", "10.2")
	registerFunc(&_hIShapeCreateMutable, &_hIShapeCreateMutableErr, frameworkHandle, "HIShapeCreateMutable", "10.2")
	registerFunc(&_hIShapeCreateMutableCopy, &_hIShapeCreateMutableCopyErr, frameworkHandle, "HIShapeCreateMutableCopy", "10.2")
	registerFunc(&_hIShapeCreateMutableWithRect, &_hIShapeCreateMutableWithRectErr, frameworkHandle, "HIShapeCreateMutableWithRect", "10.5")
	registerFunc(&_hIShapeCreateUnion, &_hIShapeCreateUnionErr, frameworkHandle, "HIShapeCreateUnion", "10.2")
	registerFunc(&_hIShapeCreateWithQDRgn, &_hIShapeCreateWithQDRgnErr, frameworkHandle, "HIShapeCreateWithQDRgn", "10.2")
	registerFunc(&_hIShapeCreateWithRect, &_hIShapeCreateWithRectErr, frameworkHandle, "HIShapeCreateWithRect", "10.2")
	registerFunc(&_hIShapeCreateXor, &_hIShapeCreateXorErr, frameworkHandle, "HIShapeCreateXor", "10.5")
	registerFunc(&_hIShapeDifference, &_hIShapeDifferenceErr, frameworkHandle, "HIShapeDifference", "10.2")
	registerFunc(&_hIShapeEnumerate, &_hIShapeEnumerateErr, frameworkHandle, "HIShapeEnumerate", "10.5")
	registerFunc(&_hIShapeGetAsQDRgn, &_hIShapeGetAsQDRgnErr, frameworkHandle, "HIShapeGetAsQDRgn", "10.2")
	registerFunc(&_hIShapeGetBounds, &_hIShapeGetBoundsErr, frameworkHandle, "HIShapeGetBounds", "10.2")
	registerFunc(&_hIShapeGetTypeID, &_hIShapeGetTypeIDErr, frameworkHandle, "HIShapeGetTypeID", "10.2")
	registerFunc(&_hIShapeInset, &_hIShapeInsetErr, frameworkHandle, "HIShapeInset", "10.5")
	registerFunc(&_hIShapeIntersect, &_hIShapeIntersectErr, frameworkHandle, "HIShapeIntersect", "10.2")
	registerFunc(&_hIShapeIntersectsRect, &_hIShapeIntersectsRectErr, frameworkHandle, "HIShapeIntersectsRect", "10.4")
	registerFunc(&_hIShapeIsEmpty, &_hIShapeIsEmptyErr, frameworkHandle, "HIShapeIsEmpty", "10.2")
	registerFunc(&_hIShapeIsRectangular, &_hIShapeIsRectangularErr, frameworkHandle, "HIShapeIsRectangular", "10.2")
	registerFunc(&_hIShapeOffset, &_hIShapeOffsetErr, frameworkHandle, "HIShapeOffset", "10.2")
	registerFunc(&_hIShapeReplacePathInCGContext, &_hIShapeReplacePathInCGContextErr, frameworkHandle, "HIShapeReplacePathInCGContext", "10.2")
	registerFunc(&_hIShapeSetEmpty, &_hIShapeSetEmptyErr, frameworkHandle, "HIShapeSetEmpty", "10.2")
	registerFunc(&_hIShapeSetWithShape, &_hIShapeSetWithShapeErr, frameworkHandle, "HIShapeSetWithShape", "10.5")
	registerFunc(&_hIShapeUnion, &_hIShapeUnionErr, frameworkHandle, "HIShapeUnion", "10.2")
	registerFunc(&_hIShapeUnionWithRect, &_hIShapeUnionWithRectErr, frameworkHandle, "HIShapeUnionWithRect", "10.5")
	registerFunc(&_hIShapeXor, &_hIShapeXorErr, frameworkHandle, "HIShapeXor", "10.5")
	registerFunc(&_iCAddMapEntry, &_iCAddMapEntryErr, frameworkHandle, "ICAddMapEntry", "10.0")
	registerFunc(&_iCAddProfile, &_iCAddProfileErr, frameworkHandle, "ICAddProfile", "10.0")
	registerFunc(&_iCBegin, &_iCBeginErr, frameworkHandle, "ICBegin", "10.0")
	registerFunc(&_iCCountMapEntries, &_iCCountMapEntriesErr, frameworkHandle, "ICCountMapEntries", "10.0")
	registerFunc(&_iCCountPref, &_iCCountPrefErr, frameworkHandle, "ICCountPref", "10.0")
	registerFunc(&_iCCountProfiles, &_iCCountProfilesErr, frameworkHandle, "ICCountProfiles", "10.0")
	registerFunc(&_iCCreateGURLEvent, &_iCCreateGURLEventErr, frameworkHandle, "ICCreateGURLEvent", "10.0")
	registerFunc(&_iCDeleteMapEntry, &_iCDeleteMapEntryErr, frameworkHandle, "ICDeleteMapEntry", "10.0")
	registerFunc(&_iCDeletePref, &_iCDeletePrefErr, frameworkHandle, "ICDeletePref", "10.0")
	registerFunc(&_iCDeleteProfile, &_iCDeleteProfileErr, frameworkHandle, "ICDeleteProfile", "10.0")
	registerFunc(&_iCEditPreferences, &_iCEditPreferencesErr, frameworkHandle, "ICEditPreferences", "10.0")
	registerFunc(&_iCEnd, &_iCEndErr, frameworkHandle, "ICEnd", "10.0")
	registerFunc(&_iCFindPrefHandle, &_iCFindPrefHandleErr, frameworkHandle, "ICFindPrefHandle", "10.0")
	registerFunc(&_iCGetConfigName, &_iCGetConfigNameErr, frameworkHandle, "ICGetConfigName", "10.0")
	registerFunc(&_iCGetCurrentProfile, &_iCGetCurrentProfileErr, frameworkHandle, "ICGetCurrentProfile", "10.0")
	registerFunc(&_iCGetDefaultPref, &_iCGetDefaultPrefErr, frameworkHandle, "ICGetDefaultPref", "10.0")
	registerFunc(&_iCGetIndMapEntry, &_iCGetIndMapEntryErr, frameworkHandle, "ICGetIndMapEntry", "10.0")
	registerFunc(&_iCGetIndPref, &_iCGetIndPrefErr, frameworkHandle, "ICGetIndPref", "10.0")
	registerFunc(&_iCGetIndProfile, &_iCGetIndProfileErr, frameworkHandle, "ICGetIndProfile", "10.0")
	registerFunc(&_iCGetMapEntry, &_iCGetMapEntryErr, frameworkHandle, "ICGetMapEntry", "10.0")
	registerFunc(&_iCGetPerm, &_iCGetPermErr, frameworkHandle, "ICGetPerm", "10.0")
	registerFunc(&_iCGetPref, &_iCGetPrefErr, frameworkHandle, "ICGetPref", "10.0")
	registerFunc(&_iCGetPrefHandle, &_iCGetPrefHandleErr, frameworkHandle, "ICGetPrefHandle", "10.0")
	registerFunc(&_iCGetProfileName, &_iCGetProfileNameErr, frameworkHandle, "ICGetProfileName", "10.0")
	registerFunc(&_iCGetSeed, &_iCGetSeedErr, frameworkHandle, "ICGetSeed", "10.0")
	registerFunc(&_iCGetVersion, &_iCGetVersionErr, frameworkHandle, "ICGetVersion", "10.0")
	registerFunc(&_iCLaunchURL, &_iCLaunchURLErr, frameworkHandle, "ICLaunchURL", "10.0")
	registerFunc(&_iCMapEntriesFilename, &_iCMapEntriesFilenameErr, frameworkHandle, "ICMapEntriesFilename", "10.0")
	registerFunc(&_iCMapEntriesTypeCreator, &_iCMapEntriesTypeCreatorErr, frameworkHandle, "ICMapEntriesTypeCreator", "10.0")
	registerFunc(&_iCMapFilename, &_iCMapFilenameErr, frameworkHandle, "ICMapFilename", "10.0")
	registerFunc(&_iCMapTypeCreator, &_iCMapTypeCreatorErr, frameworkHandle, "ICMapTypeCreator", "10.0")
	registerFunc(&_iCParseURL, &_iCParseURLErr, frameworkHandle, "ICParseURL", "10.0")
	registerFunc(&_iCSendGURLEvent, &_iCSendGURLEventErr, frameworkHandle, "ICSendGURLEvent", "10.0")
	registerFunc(&_iCSetCurrentProfile, &_iCSetCurrentProfileErr, frameworkHandle, "ICSetCurrentProfile", "10.0")
	registerFunc(&_iCSetMapEntry, &_iCSetMapEntryErr, frameworkHandle, "ICSetMapEntry", "10.0")
	registerFunc(&_iCSetPref, &_iCSetPrefErr, frameworkHandle, "ICSetPref", "10.0")
	registerFunc(&_iCSetPrefHandle, &_iCSetPrefHandleErr, frameworkHandle, "ICSetPrefHandle", "10.0")
	registerFunc(&_iCSetProfileName, &_iCSetProfileNameErr, frameworkHandle, "ICSetProfileName", "10.0")
	registerFunc(&_iCStart, &_iCStartErr, frameworkHandle, "ICStart", "10.0")
	registerFunc(&_iCStop, &_iCStopErr, frameworkHandle, "ICStop", "10.0")
	registerFunc(&_iconRefContainsCGPoint, &_iconRefContainsCGPointErr, frameworkHandle, "IconRefContainsCGPoint", "10.5")
	registerFunc(&_iconRefIntersectsCGRect, &_iconRefIntersectsCGRectErr, frameworkHandle, "IconRefIntersectsCGRect", "10.5")
	registerFunc(&_iconRefToHIShape, &_iconRefToHIShapeErr, frameworkHandle, "IconRefToHIShape", "10.5")
	registerFunc(&_iconRefToIconFamily, &_iconRefToIconFamilyErr, frameworkHandle, "IconRefToIconFamily", "10.0")
	registerFunc(&_invokeIconActionUPP, &_invokeIconActionUPPErr, frameworkHandle, "InvokeIconActionUPP", "10.0")
	registerFunc(&_invokeIconGetterUPP, &_invokeIconGetterUPPErr, frameworkHandle, "InvokeIconGetterUPP", "10.0")
	registerFunc(&_invokeSpeechDoneUPP, &_invokeSpeechDoneUPPErr, frameworkHandle, "InvokeSpeechDoneUPP", "10.0")
	registerFunc(&_invokeSpeechErrorUPP, &_invokeSpeechErrorUPPErr, frameworkHandle, "InvokeSpeechErrorUPP", "10.0")
	registerFunc(&_invokeSpeechPhonemeUPP, &_invokeSpeechPhonemeUPPErr, frameworkHandle, "InvokeSpeechPhonemeUPP", "10.0")
	registerFunc(&_invokeSpeechSyncUPP, &_invokeSpeechSyncUPPErr, frameworkHandle, "InvokeSpeechSyncUPP", "10.0")
	registerFunc(&_invokeSpeechTextDoneUPP, &_invokeSpeechTextDoneUPPErr, frameworkHandle, "InvokeSpeechTextDoneUPP", "10.0")
	registerFunc(&_invokeSpeechWordUPP, &_invokeSpeechWordUPPErr, frameworkHandle, "InvokeSpeechWordUPP", "10.0")
	registerFunc(&_isIconRefMaskEmpty, &_isIconRefMaskEmptyErr, frameworkHandle, "IsIconRefMaskEmpty", "10.0")
	registerFunc(&_isProcessVisible, &_isProcessVisibleErr, frameworkHandle, "IsProcessVisible", "10.1")
	registerFunc(&_killProcess, &_killProcessErr, frameworkHandle, "KillProcess", "10.2")
	registerFunc(&_launchApplication, &_launchApplicationErr, frameworkHandle, "LaunchApplication", "10.0")
	registerFunc(&_makeVoiceSpec, &_makeVoiceSpecErr, frameworkHandle, "MakeVoiceSpec", "10.0")
	registerFunc(&_newIconActionUPP, &_newIconActionUPPErr, frameworkHandle, "NewIconActionUPP", "10.0")
	registerFunc(&_newIconGetterUPP, &_newIconGetterUPPErr, frameworkHandle, "NewIconGetterUPP", "10.0")
	registerFunc(&_newSpeechChannel, &_newSpeechChannelErr, frameworkHandle, "NewSpeechChannel", "10.0")
	registerFunc(&_newSpeechDoneUPP, &_newSpeechDoneUPPErr, frameworkHandle, "NewSpeechDoneUPP", "10.0")
	registerFunc(&_newSpeechErrorUPP, &_newSpeechErrorUPPErr, frameworkHandle, "NewSpeechErrorUPP", "10.0")
	registerFunc(&_newSpeechPhonemeUPP, &_newSpeechPhonemeUPPErr, frameworkHandle, "NewSpeechPhonemeUPP", "10.0")
	registerFunc(&_newSpeechSyncUPP, &_newSpeechSyncUPPErr, frameworkHandle, "NewSpeechSyncUPP", "10.0")
	registerFunc(&_newSpeechTextDoneUPP, &_newSpeechTextDoneUPPErr, frameworkHandle, "NewSpeechTextDoneUPP", "10.0")
	registerFunc(&_newSpeechWordUPP, &_newSpeechWordUPPErr, frameworkHandle, "NewSpeechWordUPP", "10.0")
	registerFunc(&_pMCGImageCreateWithEPSDataProvider, &_pMCGImageCreateWithEPSDataProviderErr, frameworkHandle, "PMCGImageCreateWithEPSDataProvider", "10.1")
	registerFunc(&_pMCopyAvailablePPDs, &_pMCopyAvailablePPDsErr, frameworkHandle, "PMCopyAvailablePPDs", "10.3")
	registerFunc(&_pMCopyLocalizedPPD, &_pMCopyLocalizedPPDErr, frameworkHandle, "PMCopyLocalizedPPD", "10.3")
	registerFunc(&_pMCopyPPDData, &_pMCopyPPDDataErr, frameworkHandle, "PMCopyPPDData", "10.3")
	registerFunc(&_pMCopyPageFormat, &_pMCopyPageFormatErr, frameworkHandle, "PMCopyPageFormat", "10.0")
	registerFunc(&_pMCopyPrintSettings, &_pMCopyPrintSettingsErr, frameworkHandle, "PMCopyPrintSettings", "10.0")
	registerFunc(&_pMCreateGenericPrinter, &_pMCreateGenericPrinterErr, frameworkHandle, "PMCreateGenericPrinter", "10.5")
	registerFunc(&_pMCreatePageFormat, &_pMCreatePageFormatErr, frameworkHandle, "PMCreatePageFormat", "10.0")
	registerFunc(&_pMCreatePageFormatWithPMPaper, &_pMCreatePageFormatWithPMPaperErr, frameworkHandle, "PMCreatePageFormatWithPMPaper", "10.3")
	registerFunc(&_pMCreatePrintSettings, &_pMCreatePrintSettingsErr, frameworkHandle, "PMCreatePrintSettings", "10.0")
	registerFunc(&_pMCreateSession, &_pMCreateSessionErr, frameworkHandle, "PMCreateSession", "10.0")
	registerFunc(&_pMGetAdjustedPageRect, &_pMGetAdjustedPageRectErr, frameworkHandle, "PMGetAdjustedPageRect", "10.0")
	registerFunc(&_pMGetAdjustedPaperRect, &_pMGetAdjustedPaperRectErr, frameworkHandle, "PMGetAdjustedPaperRect", "10.0")
	registerFunc(&_pMGetCollate, &_pMGetCollateErr, frameworkHandle, "PMGetCollate", "10.2")
	registerFunc(&_pMGetCopies, &_pMGetCopiesErr, frameworkHandle, "PMGetCopies", "10.0")
	registerFunc(&_pMGetDuplex, &_pMGetDuplexErr, frameworkHandle, "PMGetDuplex", "10.4")
	registerFunc(&_pMGetFirstPage, &_pMGetFirstPageErr, frameworkHandle, "PMGetFirstPage", "10.0")
	registerFunc(&_pMGetLastPage, &_pMGetLastPageErr, frameworkHandle, "PMGetLastPage", "10.0")
	registerFunc(&_pMGetOrientation, &_pMGetOrientationErr, frameworkHandle, "PMGetOrientation", "10.0")
	registerFunc(&_pMGetPageFormatExtendedData, &_pMGetPageFormatExtendedDataErr, frameworkHandle, "PMGetPageFormatExtendedData", "10.0")
	registerFunc(&_pMGetPageFormatPaper, &_pMGetPageFormatPaperErr, frameworkHandle, "PMGetPageFormatPaper", "10.3")
	registerFunc(&_pMGetPageRange, &_pMGetPageRangeErr, frameworkHandle, "PMGetPageRange", "10.0")
	registerFunc(&_pMGetScale, &_pMGetScaleErr, frameworkHandle, "PMGetScale", "10.0")
	registerFunc(&_pMGetUnadjustedPageRect, &_pMGetUnadjustedPageRectErr, frameworkHandle, "PMGetUnadjustedPageRect", "10.0")
	registerFunc(&_pMGetUnadjustedPaperRect, &_pMGetUnadjustedPaperRectErr, frameworkHandle, "PMGetUnadjustedPaperRect", "10.0")
	registerFunc(&_pMPageFormatCreateDataRepresentation, &_pMPageFormatCreateDataRepresentationErr, frameworkHandle, "PMPageFormatCreateDataRepresentation", "10.5")
	registerFunc(&_pMPageFormatCreateWithDataRepresentation, &_pMPageFormatCreateWithDataRepresentationErr, frameworkHandle, "PMPageFormatCreateWithDataRepresentation", "10.5")
	registerFunc(&_pMPageFormatGetPrinterID, &_pMPageFormatGetPrinterIDErr, frameworkHandle, "PMPageFormatGetPrinterID", "10.5")
	registerFunc(&_pMPaperCreateCustom, &_pMPaperCreateCustomErr, frameworkHandle, "PMPaperCreateCustom", "10.5")
	registerFunc(&_pMPaperCreateLocalizedName, &_pMPaperCreateLocalizedNameErr, frameworkHandle, "PMPaperCreateLocalizedName", "10.5")
	registerFunc(&_pMPaperGetHeight, &_pMPaperGetHeightErr, frameworkHandle, "PMPaperGetHeight", "10.3")
	registerFunc(&_pMPaperGetID, &_pMPaperGetIDErr, frameworkHandle, "PMPaperGetID", "10.3")
	registerFunc(&_pMPaperGetMargins, &_pMPaperGetMarginsErr, frameworkHandle, "PMPaperGetMargins", "10.3")
	registerFunc(&_pMPaperGetPPDPaperName, &_pMPaperGetPPDPaperNameErr, frameworkHandle, "PMPaperGetPPDPaperName", "10.5")
	registerFunc(&_pMPaperGetPrinterID, &_pMPaperGetPrinterIDErr, frameworkHandle, "PMPaperGetPrinterID", "10.5")
	registerFunc(&_pMPaperGetWidth, &_pMPaperGetWidthErr, frameworkHandle, "PMPaperGetWidth", "10.3")
	registerFunc(&_pMPaperIsCustom, &_pMPaperIsCustomErr, frameworkHandle, "PMPaperIsCustom", "10.5")
	registerFunc(&_pMPresetCopyName, &_pMPresetCopyNameErr, frameworkHandle, "PMPresetCopyName", "10.3")
	registerFunc(&_pMPresetCreatePrintSettings, &_pMPresetCreatePrintSettingsErr, frameworkHandle, "PMPresetCreatePrintSettings", "10.3")
	registerFunc(&_pMPresetGetAttributes, &_pMPresetGetAttributesErr, frameworkHandle, "PMPresetGetAttributes", "10.3")
	registerFunc(&_pMPrintSettingsCopyAsDictionary, &_pMPrintSettingsCopyAsDictionaryErr, frameworkHandle, "PMPrintSettingsCopyAsDictionary", "10.5")
	registerFunc(&_pMPrintSettingsCopyKeys, &_pMPrintSettingsCopyKeysErr, frameworkHandle, "PMPrintSettingsCopyKeys", "10.5")
	registerFunc(&_pMPrintSettingsCreateDataRepresentation, &_pMPrintSettingsCreateDataRepresentationErr, frameworkHandle, "PMPrintSettingsCreateDataRepresentation", "10.5")
	registerFunc(&_pMPrintSettingsCreateWithDataRepresentation, &_pMPrintSettingsCreateWithDataRepresentationErr, frameworkHandle, "PMPrintSettingsCreateWithDataRepresentation", "10.5")
	registerFunc(&_pMPrintSettingsGetJobName, &_pMPrintSettingsGetJobNameErr, frameworkHandle, "PMPrintSettingsGetJobName", "10.4")
	registerFunc(&_pMPrintSettingsGetValue, &_pMPrintSettingsGetValueErr, frameworkHandle, "PMPrintSettingsGetValue", "10.4")
	registerFunc(&_pMPrintSettingsSetJobName, &_pMPrintSettingsSetJobNameErr, frameworkHandle, "PMPrintSettingsSetJobName", "10.4")
	registerFunc(&_pMPrintSettingsSetValue, &_pMPrintSettingsSetValueErr, frameworkHandle, "PMPrintSettingsSetValue", "10.4")
	registerFunc(&_pMPrintSettingsToOptions, &_pMPrintSettingsToOptionsErr, frameworkHandle, "PMPrintSettingsToOptions", "10.3")
	registerFunc(&_pMPrintSettingsToOptionsWithPrinterAndPageFormat, &_pMPrintSettingsToOptionsWithPrinterAndPageFormatErr, frameworkHandle, "PMPrintSettingsToOptionsWithPrinterAndPageFormat", "10.5")
	registerFunc(&_pMPrinterCopyDescriptionURL, &_pMPrinterCopyDescriptionURLErr, frameworkHandle, "PMPrinterCopyDescriptionURL", "10.4")
	registerFunc(&_pMPrinterCopyDeviceURI, &_pMPrinterCopyDeviceURIErr, frameworkHandle, "PMPrinterCopyDeviceURI", "10.4")
	registerFunc(&_pMPrinterCopyHostName, &_pMPrinterCopyHostNameErr, frameworkHandle, "PMPrinterCopyHostName", "10.3")
	registerFunc(&_pMPrinterCopyPresets, &_pMPrinterCopyPresetsErr, frameworkHandle, "PMPrinterCopyPresets", "10.3")
	registerFunc(&_pMPrinterCopyState, &_pMPrinterCopyStateErr, frameworkHandle, "PMPrinterCopyState", "10.6")
	registerFunc(&_pMPrinterCreateFromPrinterID, &_pMPrinterCreateFromPrinterIDErr, frameworkHandle, "PMPrinterCreateFromPrinterID", "10.4")
	registerFunc(&_pMPrinterGetCommInfo, &_pMPrinterGetCommInfoErr, frameworkHandle, "PMPrinterGetCommInfo", "10.3")
	registerFunc(&_pMPrinterGetDriverCreator, &_pMPrinterGetDriverCreatorErr, frameworkHandle, "PMPrinterGetDriverCreator", "10.0")
	registerFunc(&_pMPrinterGetDriverReleaseInfo, &_pMPrinterGetDriverReleaseInfoErr, frameworkHandle, "PMPrinterGetDriverReleaseInfo", "10.0")
	registerFunc(&_pMPrinterGetID, &_pMPrinterGetIDErr, frameworkHandle, "PMPrinterGetID", "10.2")
	registerFunc(&_pMPrinterGetIndexedPrinterResolution, &_pMPrinterGetIndexedPrinterResolutionErr, frameworkHandle, "PMPrinterGetIndexedPrinterResolution", "10.0")
	registerFunc(&_pMPrinterGetLanguageInfo, &_pMPrinterGetLanguageInfoErr, frameworkHandle, "PMPrinterGetLanguageInfo", "10.0")
	registerFunc(&_pMPrinterGetLocation, &_pMPrinterGetLocationErr, frameworkHandle, "PMPrinterGetLocation", "10.2")
	registerFunc(&_pMPrinterGetMakeAndModelName, &_pMPrinterGetMakeAndModelNameErr, frameworkHandle, "PMPrinterGetMakeAndModelName", "10.2")
	registerFunc(&_pMPrinterGetMimeTypes, &_pMPrinterGetMimeTypesErr, frameworkHandle, "PMPrinterGetMimeTypes", "10.3")
	registerFunc(&_pMPrinterGetName, &_pMPrinterGetNameErr, frameworkHandle, "PMPrinterGetName", "10.2")
	registerFunc(&_pMPrinterGetOutputResolution, &_pMPrinterGetOutputResolutionErr, frameworkHandle, "PMPrinterGetOutputResolution", "10.5")
	registerFunc(&_pMPrinterGetPaperList, &_pMPrinterGetPaperListErr, frameworkHandle, "PMPrinterGetPaperList", "10.3")
	registerFunc(&_pMPrinterGetPrinterResolutionCount, &_pMPrinterGetPrinterResolutionCountErr, frameworkHandle, "PMPrinterGetPrinterResolutionCount", "10.0")
	registerFunc(&_pMPrinterGetState, &_pMPrinterGetStateErr, frameworkHandle, "PMPrinterGetState", "10.2")
	registerFunc(&_pMPrinterIsDefault, &_pMPrinterIsDefaultErr, frameworkHandle, "PMPrinterIsDefault", "10.2")
	registerFunc(&_pMPrinterIsFavorite, &_pMPrinterIsFavoriteErr, frameworkHandle, "PMPrinterIsFavorite", "10.2")
	registerFunc(&_pMPrinterIsPostScriptCapable, &_pMPrinterIsPostScriptCapableErr, frameworkHandle, "PMPrinterIsPostScriptCapable", "10.2")
	registerFunc(&_pMPrinterIsPostScriptPrinter, &_pMPrinterIsPostScriptPrinterErr, frameworkHandle, "PMPrinterIsPostScriptPrinter", "10.5")
	registerFunc(&_pMPrinterIsRemote, &_pMPrinterIsRemoteErr, frameworkHandle, "PMPrinterIsRemote", "10.3")
	registerFunc(&_pMPrinterPrintWithFile, &_pMPrinterPrintWithFileErr, frameworkHandle, "PMPrinterPrintWithFile", "10.3")
	registerFunc(&_pMPrinterPrintWithProvider, &_pMPrinterPrintWithProviderErr, frameworkHandle, "PMPrinterPrintWithProvider", "10.3")
	registerFunc(&_pMPrinterSendCommand, &_pMPrinterSendCommandErr, frameworkHandle, "PMPrinterSendCommand", "10.6")
	registerFunc(&_pMPrinterSetDefault, &_pMPrinterSetDefaultErr, frameworkHandle, "PMPrinterSetDefault", "10.5")
	registerFunc(&_pMPrinterSetOutputResolution, &_pMPrinterSetOutputResolutionErr, frameworkHandle, "PMPrinterSetOutputResolution", "10.5")
	registerFunc(&_pMPrinterWritePostScriptToURL, &_pMPrinterWritePostScriptToURLErr, frameworkHandle, "PMPrinterWritePostScriptToURL", "10.5")
	registerFunc(&_pMRelease, &_pMReleaseErr, frameworkHandle, "PMRelease", "10.0")
	registerFunc(&_pMRetain, &_pMRetainErr, frameworkHandle, "PMRetain", "10.0")
	registerFunc(&_pMServerCreatePrinterList, &_pMServerCreatePrinterListErr, frameworkHandle, "PMServerCreatePrinterList", "10.2")
	registerFunc(&_pMServerLaunchPrinterBrowser, &_pMServerLaunchPrinterBrowserErr, frameworkHandle, "PMServerLaunchPrinterBrowser", "10.5")
	registerFunc(&_pMSessionBeginCGDocumentNoDialog, &_pMSessionBeginCGDocumentNoDialogErr, frameworkHandle, "PMSessionBeginCGDocumentNoDialog", "10.4")
	registerFunc(&_pMSessionBeginPageNoDialog, &_pMSessionBeginPageNoDialogErr, frameworkHandle, "PMSessionBeginPageNoDialog", "10.2")
	registerFunc(&_pMSessionCopyDestinationFormat, &_pMSessionCopyDestinationFormatErr, frameworkHandle, "PMSessionCopyDestinationFormat", "10.1")
	registerFunc(&_pMSessionCopyDestinationLocation, &_pMSessionCopyDestinationLocationErr, frameworkHandle, "PMSessionCopyDestinationLocation", "10.1")
	registerFunc(&_pMSessionCopyOutputFormatList, &_pMSessionCopyOutputFormatListErr, frameworkHandle, "PMSessionCopyOutputFormatList", "10.1")
	registerFunc(&_pMSessionCreatePageFormatList, &_pMSessionCreatePageFormatListErr, frameworkHandle, "PMSessionCreatePageFormatList", "10.1")
	registerFunc(&_pMSessionCreatePrinterList, &_pMSessionCreatePrinterListErr, frameworkHandle, "PMSessionCreatePrinterList", "10.1")
	registerFunc(&_pMSessionDefaultPageFormat, &_pMSessionDefaultPageFormatErr, frameworkHandle, "PMSessionDefaultPageFormat", "10.0")
	registerFunc(&_pMSessionDefaultPrintSettings, &_pMSessionDefaultPrintSettingsErr, frameworkHandle, "PMSessionDefaultPrintSettings", "10.0")
	registerFunc(&_pMSessionEndDocumentNoDialog, &_pMSessionEndDocumentNoDialogErr, frameworkHandle, "PMSessionEndDocumentNoDialog", "10.2")
	registerFunc(&_pMSessionEndPageNoDialog, &_pMSessionEndPageNoDialogErr, frameworkHandle, "PMSessionEndPageNoDialog", "10.2")
	registerFunc(&_pMSessionError, &_pMSessionErrorErr, frameworkHandle, "PMSessionError", "10.0")
	registerFunc(&_pMSessionGetCGGraphicsContext, &_pMSessionGetCGGraphicsContextErr, frameworkHandle, "PMSessionGetCGGraphicsContext", "10.4")
	registerFunc(&_pMSessionGetCurrentPrinter, &_pMSessionGetCurrentPrinterErr, frameworkHandle, "PMSessionGetCurrentPrinter", "10.0")
	registerFunc(&_pMSessionGetDataFromSession, &_pMSessionGetDataFromSessionErr, frameworkHandle, "PMSessionGetDataFromSession", "10.0")
	registerFunc(&_pMSessionGetDestinationType, &_pMSessionGetDestinationTypeErr, frameworkHandle, "PMSessionGetDestinationType", "10.1")
	registerFunc(&_pMSessionSetCurrentPMPrinter, &_pMSessionSetCurrentPMPrinterErr, frameworkHandle, "PMSessionSetCurrentPMPrinter", "10.3")
	registerFunc(&_pMSessionSetDataInSession, &_pMSessionSetDataInSessionErr, frameworkHandle, "PMSessionSetDataInSession", "10.0")
	registerFunc(&_pMSessionSetDestination, &_pMSessionSetDestinationErr, frameworkHandle, "PMSessionSetDestination", "10.1")
	registerFunc(&_pMSessionSetError, &_pMSessionSetErrorErr, frameworkHandle, "PMSessionSetError", "10.0")
	registerFunc(&_pMSessionValidatePageFormat, &_pMSessionValidatePageFormatErr, frameworkHandle, "PMSessionValidatePageFormat", "10.0")
	registerFunc(&_pMSessionValidatePrintSettings, &_pMSessionValidatePrintSettingsErr, frameworkHandle, "PMSessionValidatePrintSettings", "10.0")
	registerFunc(&_pMSetCollate, &_pMSetCollateErr, frameworkHandle, "PMSetCollate", "10.2")
	registerFunc(&_pMSetCopies, &_pMSetCopiesErr, frameworkHandle, "PMSetCopies", "10.0")
	registerFunc(&_pMSetDuplex, &_pMSetDuplexErr, frameworkHandle, "PMSetDuplex", "10.4")
	registerFunc(&_pMSetFirstPage, &_pMSetFirstPageErr, frameworkHandle, "PMSetFirstPage", "10.0")
	registerFunc(&_pMSetLastPage, &_pMSetLastPageErr, frameworkHandle, "PMSetLastPage", "10.0")
	registerFunc(&_pMSetOrientation, &_pMSetOrientationErr, frameworkHandle, "PMSetOrientation", "10.0")
	registerFunc(&_pMSetPageFormatExtendedData, &_pMSetPageFormatExtendedDataErr, frameworkHandle, "PMSetPageFormatExtendedData", "10.0")
	registerFunc(&_pMSetPageRange, &_pMSetPageRangeErr, frameworkHandle, "PMSetPageRange", "10.0")
	registerFunc(&_pMSetScale, &_pMSetScaleErr, frameworkHandle, "PMSetScale", "10.0")
	registerFunc(&_pMWorkflowCopyItems, &_pMWorkflowCopyItemsErr, frameworkHandle, "PMWorkflowCopyItems", "10.3")
	registerFunc(&_pMWorkflowSubmitPDFWithOptions, &_pMWorkflowSubmitPDFWithOptionsErr, frameworkHandle, "PMWorkflowSubmitPDFWithOptions", "10.3")
	registerFunc(&_pMWorkflowSubmitPDFWithSettings, &_pMWorkflowSubmitPDFWithSettingsErr, frameworkHandle, "PMWorkflowSubmitPDFWithSettings", "10.3")
	registerFunc(&_pasteboardClear, &_pasteboardClearErr, frameworkHandle, "PasteboardClear", "10.3")
	registerFunc(&_pasteboardCopyItemFlavorData, &_pasteboardCopyItemFlavorDataErr, frameworkHandle, "PasteboardCopyItemFlavorData", "10.3")
	registerFunc(&_pasteboardCopyItemFlavors, &_pasteboardCopyItemFlavorsErr, frameworkHandle, "PasteboardCopyItemFlavors", "10.3")
	registerFunc(&_pasteboardCopyName, &_pasteboardCopyNameErr, frameworkHandle, "PasteboardCopyName", "10.4")
	registerFunc(&_pasteboardCopyPasteLocation, &_pasteboardCopyPasteLocationErr, frameworkHandle, "PasteboardCopyPasteLocation", "10.3")
	registerFunc(&_pasteboardCreate, &_pasteboardCreateErr, frameworkHandle, "PasteboardCreate", "10.3")
	registerFunc(&_pasteboardGetItemCount, &_pasteboardGetItemCountErr, frameworkHandle, "PasteboardGetItemCount", "10.3")
	registerFunc(&_pasteboardGetItemFlavorFlags, &_pasteboardGetItemFlavorFlagsErr, frameworkHandle, "PasteboardGetItemFlavorFlags", "10.3")
	registerFunc(&_pasteboardGetItemIdentifier, &_pasteboardGetItemIdentifierErr, frameworkHandle, "PasteboardGetItemIdentifier", "10.3")
	registerFunc(&_pasteboardGetTypeID, &_pasteboardGetTypeIDErr, frameworkHandle, "PasteboardGetTypeID", "10.3")
	registerFunc(&_pasteboardPutItemFlavor, &_pasteboardPutItemFlavorErr, frameworkHandle, "PasteboardPutItemFlavor", "10.3")
	registerFunc(&_pasteboardResolvePromises, &_pasteboardResolvePromisesErr, frameworkHandle, "PasteboardResolvePromises", "10.3")
	registerFunc(&_pasteboardSetPasteLocation, &_pasteboardSetPasteLocationErr, frameworkHandle, "PasteboardSetPasteLocation", "10.3")
	registerFunc(&_pasteboardSetPromiseKeeper, &_pasteboardSetPromiseKeeperErr, frameworkHandle, "PasteboardSetPromiseKeeper", "10.3")
	registerFunc(&_pasteboardSynchronize, &_pasteboardSynchronizeErr, frameworkHandle, "PasteboardSynchronize", "10.3")
	registerFunc(&_pauseSpeechAt, &_pauseSpeechAtErr, frameworkHandle, "PauseSpeechAt", "10.0")
	registerFunc(&_plotIconRefInContext, &_plotIconRefInContextErr, frameworkHandle, "PlotIconRefInContext", "10.1")
	registerFunc(&_processInformationCopyDictionary, &_processInformationCopyDictionaryErr, frameworkHandle, "ProcessInformationCopyDictionary", "10.2")
	registerFunc(&_sameProcess, &_sameProcessErr, frameworkHandle, "SameProcess", "10.0")
	registerFunc(&_setFrontProcess, &_setFrontProcessErr, frameworkHandle, "SetFrontProcess", "10.0")
	registerFunc(&_setFrontProcessWithOptions, &_setFrontProcessWithOptionsErr, frameworkHandle, "SetFrontProcessWithOptions", "10.2")
	registerFunc(&_setIconFamilyData, &_setIconFamilyDataErr, frameworkHandle, "SetIconFamilyData", "10.0")
	registerFunc(&_setSpeechInfo, &_setSpeechInfoErr, frameworkHandle, "SetSpeechInfo", "10.0")
	registerFunc(&_setSpeechPitch, &_setSpeechPitchErr, frameworkHandle, "SetSpeechPitch", "10.0")
	registerFunc(&_setSpeechProperty, &_setSpeechPropertyErr, frameworkHandle, "SetSpeechProperty", "10.5")
	registerFunc(&_setSpeechRate, &_setSpeechRateErr, frameworkHandle, "SetSpeechRate", "10.0")
	registerFunc(&_showHideProcess, &_showHideProcessErr, frameworkHandle, "ShowHideProcess", "10.1")
	registerFunc(&_speakBuffer, &_speakBufferErr, frameworkHandle, "SpeakBuffer", "10.0")
	registerFunc(&_speakCFString, &_speakCFStringErr, frameworkHandle, "SpeakCFString", "10.5")
	registerFunc(&_speakString, &_speakStringErr, frameworkHandle, "SpeakString", "10.0")
	registerFunc(&_speakText, &_speakTextErr, frameworkHandle, "SpeakText", "10.0")
	registerFunc(&_speechBusy, &_speechBusyErr, frameworkHandle, "SpeechBusy", "10.0")
	registerFunc(&_speechBusySystemWide, &_speechBusySystemWideErr, frameworkHandle, "SpeechBusySystemWide", "10.0")
	registerFunc(&_speechManagerVersion, &_speechManagerVersionErr, frameworkHandle, "SpeechManagerVersion", "10.0")
	registerFunc(&_speechSynthesisRegisterModuleURL, &_speechSynthesisRegisterModuleURLErr, frameworkHandle, "SpeechSynthesisRegisterModuleURL", "10.6")
	registerFunc(&_speechSynthesisUnregisterModuleURL, &_speechSynthesisUnregisterModuleURLErr, frameworkHandle, "SpeechSynthesisUnregisterModuleURL", "10.6")
	registerFunc(&_stopSpeech, &_stopSpeechErr, frameworkHandle, "StopSpeech", "10.0")
	registerFunc(&_stopSpeechAt, &_stopSpeechAtErr, frameworkHandle, "StopSpeechAt", "10.0")
	registerFunc(&_textToPhonemes, &_textToPhonemesErr, frameworkHandle, "TextToPhonemes", "10.0")
	registerFunc(&_transformProcessType, &_transformProcessTypeErr, frameworkHandle, "TransformProcessType", "10.3")
	registerFunc(&_translationCopyDestinationType, &_translationCopyDestinationTypeErr, frameworkHandle, "TranslationCopyDestinationType", "10.3")
	registerFunc(&_translationCopySourceType, &_translationCopySourceTypeErr, frameworkHandle, "TranslationCopySourceType", "10.3")
	registerFunc(&_translationCreate, &_translationCreateErr, frameworkHandle, "TranslationCreate", "10.3")
	registerFunc(&_translationCreateWithSourceArray, &_translationCreateWithSourceArrayErr, frameworkHandle, "TranslationCreateWithSourceArray", "10.3")
	registerFunc(&_translationGetTranslationFlags, &_translationGetTranslationFlagsErr, frameworkHandle, "TranslationGetTranslationFlags", "10.3")
	registerFunc(&_translationGetTypeID, &_translationGetTypeIDErr, frameworkHandle, "TranslationGetTypeID", "10.3")
	registerFunc(&_translationPerformForData, &_translationPerformForDataErr, frameworkHandle, "TranslationPerformForData", "10.3")
	registerFunc(&_translationPerformForFile, &_translationPerformForFileErr, frameworkHandle, "TranslationPerformForFile", "10.3")
	registerFunc(&_translationPerformForURL, &_translationPerformForURLErr, frameworkHandle, "TranslationPerformForURL", "10.3")
	registerFunc(&_uAZoomChangeFocus, &_uAZoomChangeFocusErr, frameworkHandle, "UAZoomChangeFocus", "10.4")
	registerFunc(&_uAZoomEnabled, &_uAZoomEnabledErr, frameworkHandle, "UAZoomEnabled", "10.4")
	registerFunc(&_useDictionary, &_useDictionaryErr, frameworkHandle, "UseDictionary", "10.0")
	registerFunc(&_useSpeechDictionary, &_useSpeechDictionaryErr, frameworkHandle, "UseSpeechDictionary", "10.5")
	registerFunc(&_wakeUpProcess, &_wakeUpProcessErr, frameworkHandle, "WakeUpProcess", "10.0")
}
