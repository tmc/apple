// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
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
		return fmt.Sprintf("DiskArbitration: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("DiskArbitration: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("DiskArbitration: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("DiskArbitration: register symbol %s: %v", name, r)
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

var _dAApprovalSessionCreate func(allocator corefoundation.CFAllocatorRef) DAApprovalSessionRef
var _dAApprovalSessionCreateErr error

func tryDAApprovalSessionCreate(allocator corefoundation.CFAllocatorRef) (DAApprovalSessionRef, error) {
	if _dAApprovalSessionCreate == nil {
		return 0, symbolCallError("DAApprovalSessionCreate", "10.4", _dAApprovalSessionCreateErr)
	}
	return _dAApprovalSessionCreate(allocator), nil
}

// DAApprovalSessionCreate.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionCreate
func DAApprovalSessionCreate(allocator corefoundation.CFAllocatorRef) DAApprovalSessionRef {
	result, callErr := tryDAApprovalSessionCreate(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dAApprovalSessionGetTypeID func() uint
var _dAApprovalSessionGetTypeIDErr error

func tryDAApprovalSessionGetTypeID() (uint, error) {
	if _dAApprovalSessionGetTypeID == nil {
		return 0, symbolCallError("DAApprovalSessionGetTypeID", "10.4", _dAApprovalSessionGetTypeIDErr)
	}
	return _dAApprovalSessionGetTypeID(), nil
}

// DAApprovalSessionGetTypeID.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionGetTypeID
func DAApprovalSessionGetTypeID() uint {
	result, callErr := tryDAApprovalSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dAApprovalSessionScheduleWithRunLoop func(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _dAApprovalSessionScheduleWithRunLoopErr error

func tryDAApprovalSessionScheduleWithRunLoop(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _dAApprovalSessionScheduleWithRunLoop == nil {
		return symbolCallError("DAApprovalSessionScheduleWithRunLoop", "10.4", _dAApprovalSessionScheduleWithRunLoopErr)
	}
	_dAApprovalSessionScheduleWithRunLoop(session, runLoop, runLoopMode)
	return nil
}

// DAApprovalSessionScheduleWithRunLoop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionScheduleWithRunLoop
func DAApprovalSessionScheduleWithRunLoop(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryDAApprovalSessionScheduleWithRunLoop(session, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _dAApprovalSessionUnscheduleFromRunLoop func(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _dAApprovalSessionUnscheduleFromRunLoopErr error

func tryDAApprovalSessionUnscheduleFromRunLoop(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _dAApprovalSessionUnscheduleFromRunLoop == nil {
		return symbolCallError("DAApprovalSessionUnscheduleFromRunLoop", "10.4", _dAApprovalSessionUnscheduleFromRunLoopErr)
	}
	_dAApprovalSessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
	return nil
}

// DAApprovalSessionUnscheduleFromRunLoop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionUnscheduleFromRunLoop
func DAApprovalSessionUnscheduleFromRunLoop(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryDAApprovalSessionUnscheduleFromRunLoop(session, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _dADiskClaim func(disk DADiskRef, options DADiskClaimOptions, release DADiskClaimReleaseCallback, releaseContext unsafe.Pointer, callback DADiskClaimCallback, callbackContext unsafe.Pointer)
var _dADiskClaimErr error

func tryDADiskClaim(disk DADiskRef, options DADiskClaimOptions, release DADiskClaimReleaseCallback, releaseContext unsafe.Pointer, callback DADiskClaimCallback, callbackContext unsafe.Pointer) error {
	if _dADiskClaim == nil {
		return symbolCallError("DADiskClaim", "10.4", _dADiskClaimErr)
	}
	_dADiskClaim(disk, options, release, releaseContext, callback, callbackContext)
	return nil
}

// DADiskClaim claims the specified disk object for exclusive use.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskClaim(_:_:_:_:_:_:)
func DADiskClaim(disk DADiskRef, options DADiskClaimOptions, release DADiskClaimReleaseCallback, releaseContext unsafe.Pointer, callback DADiskClaimCallback, callbackContext unsafe.Pointer) {
	if callErr := tryDADiskClaim(disk, options, release, releaseContext, callback, callbackContext); callErr != nil {
		panic(callErr)
	}
}

var _dADiskCopyDescription func(disk DADiskRef) corefoundation.CFDictionaryRef
var _dADiskCopyDescriptionErr error

func tryDADiskCopyDescription(disk DADiskRef) (corefoundation.CFDictionaryRef, error) {
	if _dADiskCopyDescription == nil {
		return 0, symbolCallError("DADiskCopyDescription", "10.4", _dADiskCopyDescriptionErr)
	}
	return _dADiskCopyDescription(disk), nil
}

// DADiskCopyDescription obtains the Disk Arbitration description of the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyDescription(_:)
func DADiskCopyDescription(disk DADiskRef) corefoundation.CFDictionaryRef {
	result, callErr := tryDADiskCopyDescription(disk)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskCopyIOMedia func(disk DADiskRef) uintptr
var _dADiskCopyIOMediaErr error

func tryDADiskCopyIOMedia(disk DADiskRef) (uintptr, error) {
	if _dADiskCopyIOMedia == nil {
		return 0, symbolCallError("DADiskCopyIOMedia", "10.4", _dADiskCopyIOMediaErr)
	}
	return _dADiskCopyIOMedia(disk), nil
}

// DADiskCopyIOMedia obtains the I/O Kit media object for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyIOMedia(_:)
func DADiskCopyIOMedia(disk DADiskRef) uintptr {
	result, callErr := tryDADiskCopyIOMedia(disk)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskCopyWholeDisk func(disk DADiskRef) DADiskRef
var _dADiskCopyWholeDiskErr error

func tryDADiskCopyWholeDisk(disk DADiskRef) (DADiskRef, error) {
	if _dADiskCopyWholeDisk == nil {
		return 0, symbolCallError("DADiskCopyWholeDisk", "10.4", _dADiskCopyWholeDiskErr)
	}
	return _dADiskCopyWholeDisk(disk), nil
}

// DADiskCopyWholeDisk obtain the associated whole disk object for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyWholeDisk(_:)
func DADiskCopyWholeDisk(disk DADiskRef) DADiskRef {
	result, callErr := tryDADiskCopyWholeDisk(disk)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskCreateFromBSDName func(allocator corefoundation.CFAllocatorRef, session DASessionRef, name string) DADiskRef
var _dADiskCreateFromBSDNameErr error

func tryDADiskCreateFromBSDName(allocator corefoundation.CFAllocatorRef, session DASessionRef, name string) (DADiskRef, error) {
	if _dADiskCreateFromBSDName == nil {
		return 0, symbolCallError("DADiskCreateFromBSDName", "10.4", _dADiskCreateFromBSDNameErr)
	}
	return _dADiskCreateFromBSDName(allocator, session, name), nil
}

// DADiskCreateFromBSDName creates a new disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromBSDName(_:_:_:)
func DADiskCreateFromBSDName(allocator corefoundation.CFAllocatorRef, session DASessionRef, name string) DADiskRef {
	result, callErr := tryDADiskCreateFromBSDName(allocator, session, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskCreateFromIOMedia func(allocator corefoundation.CFAllocatorRef, session DASessionRef, media uintptr) DADiskRef
var _dADiskCreateFromIOMediaErr error

func tryDADiskCreateFromIOMedia(allocator corefoundation.CFAllocatorRef, session DASessionRef, media uintptr) (DADiskRef, error) {
	if _dADiskCreateFromIOMedia == nil {
		return 0, symbolCallError("DADiskCreateFromIOMedia", "10.4", _dADiskCreateFromIOMediaErr)
	}
	return _dADiskCreateFromIOMedia(allocator, session, media), nil
}

// DADiskCreateFromIOMedia creates a new disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromIOMedia(_:_:_:)
func DADiskCreateFromIOMedia(allocator corefoundation.CFAllocatorRef, session DASessionRef, media uintptr) DADiskRef {
	result, callErr := tryDADiskCreateFromIOMedia(allocator, session, media)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskCreateFromVolumePath func(allocator corefoundation.CFAllocatorRef, session DASessionRef, path corefoundation.CFURLRef) DADiskRef
var _dADiskCreateFromVolumePathErr error

func tryDADiskCreateFromVolumePath(allocator corefoundation.CFAllocatorRef, session DASessionRef, path corefoundation.CFURLRef) (DADiskRef, error) {
	if _dADiskCreateFromVolumePath == nil {
		return 0, symbolCallError("DADiskCreateFromVolumePath", "10.7", _dADiskCreateFromVolumePathErr)
	}
	return _dADiskCreateFromVolumePath(allocator, session, path), nil
}

// DADiskCreateFromVolumePath creates a new disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromVolumePath(_:_:_:)
func DADiskCreateFromVolumePath(allocator corefoundation.CFAllocatorRef, session DASessionRef, path corefoundation.CFURLRef) DADiskRef {
	result, callErr := tryDADiskCreateFromVolumePath(allocator, session, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskEject func(disk DADiskRef, options uint, callback DADiskEjectCallback, context unsafe.Pointer)
var _dADiskEjectErr error

func tryDADiskEject(disk DADiskRef, options uint, callback DADiskEjectCallback, context unsafe.Pointer) error {
	if _dADiskEject == nil {
		return symbolCallError("DADiskEject", "10.4", _dADiskEjectErr)
	}
	_dADiskEject(disk, options, callback, context)
	return nil
}

// DADiskEject ejects the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskEject(_:_:_:_:)
func DADiskEject(disk DADiskRef, options uint, callback DADiskEjectCallback, context unsafe.Pointer) {
	if callErr := tryDADiskEject(disk, options, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dADiskGetBSDName func(disk DADiskRef) *byte
var _dADiskGetBSDNameErr error

func tryDADiskGetBSDName(disk DADiskRef) (*byte, error) {
	if _dADiskGetBSDName == nil {
		return nil, symbolCallError("DADiskGetBSDName", "10.4", _dADiskGetBSDNameErr)
	}
	return _dADiskGetBSDName(disk), nil
}

// DADiskGetBSDName obtains the BSD device name for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskGetBSDName(_:)
func DADiskGetBSDName(disk DADiskRef) *byte {
	result, callErr := tryDADiskGetBSDName(disk)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskGetOptions func(disk DADiskRef) DADiskOptions
var _dADiskGetOptionsErr error

func tryDADiskGetOptions(disk DADiskRef) (DADiskOptions, error) {
	if _dADiskGetOptions == nil {
		return *new(DADiskOptions), symbolCallError("DADiskGetOptions", "10.4", _dADiskGetOptionsErr)
	}
	return _dADiskGetOptions(disk), nil
}

// DADiskGetOptions obtains the options for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskGetOptions(_:)
func DADiskGetOptions(disk DADiskRef) DADiskOptions {
	result, callErr := tryDADiskGetOptions(disk)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskGetTypeID func() uint
var _dADiskGetTypeIDErr error

func tryDADiskGetTypeID() (uint, error) {
	if _dADiskGetTypeID == nil {
		return 0, symbolCallError("DADiskGetTypeID", "10.4", _dADiskGetTypeIDErr)
	}
	return _dADiskGetTypeID(), nil
}

// DADiskGetTypeID returns the type identifier of all DADisk instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskGetTypeID()
func DADiskGetTypeID() uint {
	result, callErr := tryDADiskGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskIsClaimed func(disk DADiskRef) bool
var _dADiskIsClaimedErr error

func tryDADiskIsClaimed(disk DADiskRef) (bool, error) {
	if _dADiskIsClaimed == nil {
		return false, symbolCallError("DADiskIsClaimed", "10.4", _dADiskIsClaimedErr)
	}
	return _dADiskIsClaimed(disk), nil
}

// DADiskIsClaimed reports whether or not the disk is claimed.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskIsClaimed(_:)
func DADiskIsClaimed(disk DADiskRef) bool {
	result, callErr := tryDADiskIsClaimed(disk)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskMount func(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer)
var _dADiskMountErr error

func tryDADiskMount(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer) error {
	if _dADiskMount == nil {
		return symbolCallError("DADiskMount", "10.4", _dADiskMountErr)
	}
	_dADiskMount(disk, path, options, callback, context)
	return nil
}

// DADiskMount mounts the volume at the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMount(_:_:_:_:_:)
func DADiskMount(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer) {
	if callErr := tryDADiskMount(disk, path, options, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dADiskMountWithArguments func(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer, arguments corefoundation.CFStringRef)
var _dADiskMountWithArgumentsErr error

func tryDADiskMountWithArguments(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer, arguments corefoundation.CFStringRef) error {
	if _dADiskMountWithArguments == nil {
		return symbolCallError("DADiskMountWithArguments", "10.4", _dADiskMountWithArgumentsErr)
	}
	_dADiskMountWithArguments(disk, path, options, callback, context, arguments)
	return nil
}

// DADiskMountWithArguments mounts the volume at the specified disk object, with the specified mount options.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMountWithArguments(_:_:_:_:_:_:)
func DADiskMountWithArguments(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer, arguments corefoundation.CFStringRef) {
	if callErr := tryDADiskMountWithArguments(disk, path, options, callback, context, arguments); callErr != nil {
		panic(callErr)
	}
}

var _dADiskRename func(disk DADiskRef, name corefoundation.CFStringRef, options DADiskRenameOptions, callback DADiskRenameCallback, context unsafe.Pointer)
var _dADiskRenameErr error

func tryDADiskRename(disk DADiskRef, name corefoundation.CFStringRef, options DADiskRenameOptions, callback DADiskRenameCallback, context unsafe.Pointer) error {
	if _dADiskRename == nil {
		return symbolCallError("DADiskRename", "10.4", _dADiskRenameErr)
	}
	_dADiskRename(disk, name, options, callback, context)
	return nil
}

// DADiskRename renames the volume at the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskRename(_:_:_:_:_:)
func DADiskRename(disk DADiskRef, name corefoundation.CFStringRef, options DADiskRenameOptions, callback DADiskRenameCallback, context unsafe.Pointer) {
	if callErr := tryDADiskRename(disk, name, options, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dADiskSetOptions func(disk DADiskRef, options DADiskOptions, value bool) DAReturn
var _dADiskSetOptionsErr error

func tryDADiskSetOptions(disk DADiskRef, options DADiskOptions, value bool) (DAReturn, error) {
	if _dADiskSetOptions == nil {
		return *new(DAReturn), symbolCallError("DADiskSetOptions", "10.4", _dADiskSetOptionsErr)
	}
	return _dADiskSetOptions(disk, options, value), nil
}

// DADiskSetOptions sets the options for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskSetOptions(_:_:_:)
func DADiskSetOptions(disk DADiskRef, options DADiskOptions, value bool) DAReturn {
	result, callErr := tryDADiskSetOptions(disk, options, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADiskUnclaim func(disk DADiskRef)
var _dADiskUnclaimErr error

func tryDADiskUnclaim(disk DADiskRef) error {
	if _dADiskUnclaim == nil {
		return symbolCallError("DADiskUnclaim", "10.4", _dADiskUnclaimErr)
	}
	_dADiskUnclaim(disk)
	return nil
}

// DADiskUnclaim unclaims the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnclaim(_:)
func DADiskUnclaim(disk DADiskRef) {
	if callErr := tryDADiskUnclaim(disk); callErr != nil {
		panic(callErr)
	}
}

var _dADiskUnmount func(disk DADiskRef, options DADiskUnmountOptions, callback DADiskUnmountCallback, context unsafe.Pointer)
var _dADiskUnmountErr error

func tryDADiskUnmount(disk DADiskRef, options DADiskUnmountOptions, callback DADiskUnmountCallback, context unsafe.Pointer) error {
	if _dADiskUnmount == nil {
		return symbolCallError("DADiskUnmount", "10.4", _dADiskUnmountErr)
	}
	_dADiskUnmount(disk, options, callback, context)
	return nil
}

// DADiskUnmount unmounts the volume at the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmount(_:_:_:_:)
func DADiskUnmount(disk DADiskRef, options DADiskUnmountOptions, callback DADiskUnmountCallback, context unsafe.Pointer) {
	if callErr := tryDADiskUnmount(disk, options, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dADissenterCreate func(allocator corefoundation.CFAllocatorRef, status DAReturn, string_ corefoundation.CFStringRef) DADissenterRef
var _dADissenterCreateErr error

func tryDADissenterCreate(allocator corefoundation.CFAllocatorRef, status DAReturn, string_ corefoundation.CFStringRef) (DADissenterRef, error) {
	if _dADissenterCreate == nil {
		return 0, symbolCallError("DADissenterCreate", "10.4", _dADissenterCreateErr)
	}
	return _dADissenterCreate(allocator, status, string_), nil
}

// DADissenterCreate creates a new dissenter object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenterCreate(_:_:_:)
func DADissenterCreate(allocator corefoundation.CFAllocatorRef, status DAReturn, string_ corefoundation.CFStringRef) DADissenterRef {
	result, callErr := tryDADissenterCreate(allocator, status, string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADissenterGetStatus func(dissenter DADissenterRef) DAReturn
var _dADissenterGetStatusErr error

func tryDADissenterGetStatus(dissenter DADissenterRef) (DAReturn, error) {
	if _dADissenterGetStatus == nil {
		return *new(DAReturn), symbolCallError("DADissenterGetStatus", "10.4", _dADissenterGetStatusErr)
	}
	return _dADissenterGetStatus(dissenter), nil
}

// DADissenterGetStatus obtains the return code.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatus(_:)
func DADissenterGetStatus(dissenter DADissenterRef) DAReturn {
	result, callErr := tryDADissenterGetStatus(dissenter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dADissenterGetStatusString func(dissenter DADissenterRef) corefoundation.CFStringRef
var _dADissenterGetStatusStringErr error

func tryDADissenterGetStatusString(dissenter DADissenterRef) (corefoundation.CFStringRef, error) {
	if _dADissenterGetStatusString == nil {
		return 0, symbolCallError("DADissenterGetStatusString", "10.4", _dADissenterGetStatusStringErr)
	}
	return _dADissenterGetStatusString(dissenter), nil
}

// DADissenterGetStatusString obtains the return code string.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatusString(_:)
func DADissenterGetStatusString(dissenter DADissenterRef) corefoundation.CFStringRef {
	result, callErr := tryDADissenterGetStatusString(dissenter)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dARegisterDiskAppearedCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskAppearedCallback, context unsafe.Pointer)
var _dARegisterDiskAppearedCallbackErr error

func tryDARegisterDiskAppearedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskAppearedCallback, context unsafe.Pointer) error {
	if _dARegisterDiskAppearedCallback == nil {
		return symbolCallError("DARegisterDiskAppearedCallback", "10.4", _dARegisterDiskAppearedCallbackErr)
	}
	_dARegisterDiskAppearedCallback(session, match, callback, context)
	return nil
}

// DARegisterDiskAppearedCallback registers a callback function to be called whenever a disk has appeared.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskAppearedCallback(_:_:_:_:)
func DARegisterDiskAppearedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskAppearedCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskAppearedCallback(session, match, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dARegisterDiskDescriptionChangedCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, watch corefoundation.CFArrayRef, callback DADiskDescriptionChangedCallback, context unsafe.Pointer)
var _dARegisterDiskDescriptionChangedCallbackErr error

func tryDARegisterDiskDescriptionChangedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, watch corefoundation.CFArrayRef, callback DADiskDescriptionChangedCallback, context unsafe.Pointer) error {
	if _dARegisterDiskDescriptionChangedCallback == nil {
		return symbolCallError("DARegisterDiskDescriptionChangedCallback", "10.4", _dARegisterDiskDescriptionChangedCallbackErr)
	}
	_dARegisterDiskDescriptionChangedCallback(session, match, watch, callback, context)
	return nil
}

// DARegisterDiskDescriptionChangedCallback registers a callback function to be called whenever a disk description has changed.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDescriptionChangedCallback(_:_:_:_:_:)
func DARegisterDiskDescriptionChangedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, watch corefoundation.CFArrayRef, callback DADiskDescriptionChangedCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskDescriptionChangedCallback(session, match, watch, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dARegisterDiskDisappearedCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskDisappearedCallback, context unsafe.Pointer)
var _dARegisterDiskDisappearedCallbackErr error

func tryDARegisterDiskDisappearedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskDisappearedCallback, context unsafe.Pointer) error {
	if _dARegisterDiskDisappearedCallback == nil {
		return symbolCallError("DARegisterDiskDisappearedCallback", "10.4", _dARegisterDiskDisappearedCallbackErr)
	}
	_dARegisterDiskDisappearedCallback(session, match, callback, context)
	return nil
}

// DARegisterDiskDisappearedCallback registers a callback function to be called whenever a disk has disappeared.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDisappearedCallback(_:_:_:_:)
func DARegisterDiskDisappearedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskDisappearedCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskDisappearedCallback(session, match, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dARegisterDiskEjectApprovalCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskEjectApprovalCallback, context unsafe.Pointer)
var _dARegisterDiskEjectApprovalCallbackErr error

func tryDARegisterDiskEjectApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskEjectApprovalCallback, context unsafe.Pointer) error {
	if _dARegisterDiskEjectApprovalCallback == nil {
		return symbolCallError("DARegisterDiskEjectApprovalCallback", "10.4", _dARegisterDiskEjectApprovalCallbackErr)
	}
	_dARegisterDiskEjectApprovalCallback(session, match, callback, context)
	return nil
}

// DARegisterDiskEjectApprovalCallback registers a callback function to be called whenever a volume is to be ejected.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskEjectApprovalCallback(_:_:_:_:)
func DARegisterDiskEjectApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskEjectApprovalCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskEjectApprovalCallback(session, match, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dARegisterDiskMountApprovalCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskMountApprovalCallback, context unsafe.Pointer)
var _dARegisterDiskMountApprovalCallbackErr error

func tryDARegisterDiskMountApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskMountApprovalCallback, context unsafe.Pointer) error {
	if _dARegisterDiskMountApprovalCallback == nil {
		return symbolCallError("DARegisterDiskMountApprovalCallback", "10.4", _dARegisterDiskMountApprovalCallbackErr)
	}
	_dARegisterDiskMountApprovalCallback(session, match, callback, context)
	return nil
}

// DARegisterDiskMountApprovalCallback registers a callback function to be called whenever a volume is to be mounted.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskMountApprovalCallback(_:_:_:_:)
func DARegisterDiskMountApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskMountApprovalCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskMountApprovalCallback(session, match, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dARegisterDiskPeekCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, order int, callback DADiskPeekCallback, context unsafe.Pointer)
var _dARegisterDiskPeekCallbackErr error

func tryDARegisterDiskPeekCallback(session DASessionRef, match corefoundation.CFDictionaryRef, order int, callback DADiskPeekCallback, context unsafe.Pointer) error {
	if _dARegisterDiskPeekCallback == nil {
		return symbolCallError("DARegisterDiskPeekCallback", "10.4", _dARegisterDiskPeekCallbackErr)
	}
	_dARegisterDiskPeekCallback(session, match, order, callback, context)
	return nil
}

// DARegisterDiskPeekCallback registers a callback function to be called whenever a disk has been probed.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskPeekCallback(_:_:_:_:_:)
func DARegisterDiskPeekCallback(session DASessionRef, match corefoundation.CFDictionaryRef, order int, callback DADiskPeekCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskPeekCallback(session, match, order, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dARegisterDiskUnmountApprovalCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskUnmountApprovalCallback, context unsafe.Pointer)
var _dARegisterDiskUnmountApprovalCallbackErr error

func tryDARegisterDiskUnmountApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskUnmountApprovalCallback, context unsafe.Pointer) error {
	if _dARegisterDiskUnmountApprovalCallback == nil {
		return symbolCallError("DARegisterDiskUnmountApprovalCallback", "10.4", _dARegisterDiskUnmountApprovalCallbackErr)
	}
	_dARegisterDiskUnmountApprovalCallback(session, match, callback, context)
	return nil
}

// DARegisterDiskUnmountApprovalCallback registers a callback function to be called whenever a volume is to be unmounted.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskUnmountApprovalCallback(_:_:_:_:)
func DARegisterDiskUnmountApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskUnmountApprovalCallback, context unsafe.Pointer) {
	if callErr := tryDARegisterDiskUnmountApprovalCallback(session, match, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dASessionCreate func(allocator corefoundation.CFAllocatorRef) DASessionRef
var _dASessionCreateErr error

func tryDASessionCreate(allocator corefoundation.CFAllocatorRef) (DASessionRef, error) {
	if _dASessionCreate == nil {
		return 0, symbolCallError("DASessionCreate", "10.4", _dASessionCreateErr)
	}
	return _dASessionCreate(allocator), nil
}

// DASessionCreate creates a new session.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionCreate(_:)
func DASessionCreate(allocator corefoundation.CFAllocatorRef) DASessionRef {
	result, callErr := tryDASessionCreate(allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dASessionGetTypeID func() uint
var _dASessionGetTypeIDErr error

func tryDASessionGetTypeID() (uint, error) {
	if _dASessionGetTypeID == nil {
		return 0, symbolCallError("DASessionGetTypeID", "10.4", _dASessionGetTypeIDErr)
	}
	return _dASessionGetTypeID(), nil
}

// DASessionGetTypeID returns the type identifier of all DASession instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionGetTypeID()
func DASessionGetTypeID() uint {
	result, callErr := tryDASessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _dASessionScheduleWithRunLoop func(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _dASessionScheduleWithRunLoopErr error

func tryDASessionScheduleWithRunLoop(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _dASessionScheduleWithRunLoop == nil {
		return symbolCallError("DASessionScheduleWithRunLoop", "10.4", _dASessionScheduleWithRunLoopErr)
	}
	_dASessionScheduleWithRunLoop(session, runLoop, runLoopMode)
	return nil
}

// DASessionScheduleWithRunLoop schedules the session on a run loop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionScheduleWithRunLoop(_:_:_:)
func DASessionScheduleWithRunLoop(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryDASessionScheduleWithRunLoop(session, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _dASessionSetDispatchQueue func(session DASessionRef, queue uintptr)
var _dASessionSetDispatchQueueErr error

func tryDASessionSetDispatchQueue(session DASessionRef, queue dispatch.Queue) error {
	if _dASessionSetDispatchQueue == nil {
		return symbolCallError("DASessionSetDispatchQueue", "10.7", _dASessionSetDispatchQueueErr)
	}
	_dASessionSetDispatchQueue(session, uintptr(queue.Handle()))
	return nil
}

// DASessionSetDispatchQueue schedules the session on a dispatch queue.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionSetDispatchQueue(_:_:)
func DASessionSetDispatchQueue(session DASessionRef, queue dispatch.Queue) {
	if callErr := tryDASessionSetDispatchQueue(session, queue); callErr != nil {
		panic(callErr)
	}
}

var _dASessionUnscheduleFromRunLoop func(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _dASessionUnscheduleFromRunLoopErr error

func tryDASessionUnscheduleFromRunLoop(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _dASessionUnscheduleFromRunLoop == nil {
		return symbolCallError("DASessionUnscheduleFromRunLoop", "10.4", _dASessionUnscheduleFromRunLoopErr)
	}
	_dASessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
	return nil
}

// DASessionUnscheduleFromRunLoop unschedules the session from a run loop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionUnscheduleFromRunLoop(_:_:_:)
func DASessionUnscheduleFromRunLoop(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryDASessionUnscheduleFromRunLoop(session, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _dAUnregisterApprovalCallback func(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer)
var _dAUnregisterApprovalCallbackErr error

func tryDAUnregisterApprovalCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _dAUnregisterApprovalCallback == nil {
		return symbolCallError("DAUnregisterApprovalCallback", "10.4", _dAUnregisterApprovalCallbackErr)
	}
	_dAUnregisterApprovalCallback(session, callback, context)
	return nil
}

// DAUnregisterApprovalCallback unregisters a registered callback function.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterApprovalCallback
func DAUnregisterApprovalCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryDAUnregisterApprovalCallback(session, callback, context); callErr != nil {
		panic(callErr)
	}
}

var _dAUnregisterCallback func(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer)
var _dAUnregisterCallbackErr error

func tryDAUnregisterCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) error {
	if _dAUnregisterCallback == nil {
		return symbolCallError("DAUnregisterCallback", "10.4", _dAUnregisterCallbackErr)
	}
	_dAUnregisterCallback(session, callback, context)
	return nil
}

// DAUnregisterCallback unregisters a registered callback function.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterCallback(_:_:_:)
func DAUnregisterCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if callErr := tryDAUnregisterCallback(session, callback, context); callErr != nil {
		panic(callErr)
	}
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_dAApprovalSessionCreate, &_dAApprovalSessionCreateErr, frameworkHandle, "DAApprovalSessionCreate", "10.4")
	registerFunc(&_dAApprovalSessionGetTypeID, &_dAApprovalSessionGetTypeIDErr, frameworkHandle, "DAApprovalSessionGetTypeID", "10.4")
	registerFunc(&_dAApprovalSessionScheduleWithRunLoop, &_dAApprovalSessionScheduleWithRunLoopErr, frameworkHandle, "DAApprovalSessionScheduleWithRunLoop", "10.4")
	registerFunc(&_dAApprovalSessionUnscheduleFromRunLoop, &_dAApprovalSessionUnscheduleFromRunLoopErr, frameworkHandle, "DAApprovalSessionUnscheduleFromRunLoop", "10.4")
	registerFunc(&_dADiskClaim, &_dADiskClaimErr, frameworkHandle, "DADiskClaim", "10.4")
	registerFunc(&_dADiskCopyDescription, &_dADiskCopyDescriptionErr, frameworkHandle, "DADiskCopyDescription", "10.4")
	registerFunc(&_dADiskCopyIOMedia, &_dADiskCopyIOMediaErr, frameworkHandle, "DADiskCopyIOMedia", "10.4")
	registerFunc(&_dADiskCopyWholeDisk, &_dADiskCopyWholeDiskErr, frameworkHandle, "DADiskCopyWholeDisk", "10.4")
	registerFunc(&_dADiskCreateFromBSDName, &_dADiskCreateFromBSDNameErr, frameworkHandle, "DADiskCreateFromBSDName", "10.4")
	registerFunc(&_dADiskCreateFromIOMedia, &_dADiskCreateFromIOMediaErr, frameworkHandle, "DADiskCreateFromIOMedia", "10.4")
	registerFunc(&_dADiskCreateFromVolumePath, &_dADiskCreateFromVolumePathErr, frameworkHandle, "DADiskCreateFromVolumePath", "10.7")
	registerFunc(&_dADiskEject, &_dADiskEjectErr, frameworkHandle, "DADiskEject", "10.4")
	registerFunc(&_dADiskGetBSDName, &_dADiskGetBSDNameErr, frameworkHandle, "DADiskGetBSDName", "10.4")
	registerFunc(&_dADiskGetOptions, &_dADiskGetOptionsErr, frameworkHandle, "DADiskGetOptions", "10.4")
	registerFunc(&_dADiskGetTypeID, &_dADiskGetTypeIDErr, frameworkHandle, "DADiskGetTypeID", "10.4")
	registerFunc(&_dADiskIsClaimed, &_dADiskIsClaimedErr, frameworkHandle, "DADiskIsClaimed", "10.4")
	registerFunc(&_dADiskMount, &_dADiskMountErr, frameworkHandle, "DADiskMount", "10.4")
	registerFunc(&_dADiskMountWithArguments, &_dADiskMountWithArgumentsErr, frameworkHandle, "DADiskMountWithArguments", "10.4")
	registerFunc(&_dADiskRename, &_dADiskRenameErr, frameworkHandle, "DADiskRename", "10.4")
	registerFunc(&_dADiskSetOptions, &_dADiskSetOptionsErr, frameworkHandle, "DADiskSetOptions", "10.4")
	registerFunc(&_dADiskUnclaim, &_dADiskUnclaimErr, frameworkHandle, "DADiskUnclaim", "10.4")
	registerFunc(&_dADiskUnmount, &_dADiskUnmountErr, frameworkHandle, "DADiskUnmount", "10.4")
	registerFunc(&_dADissenterCreate, &_dADissenterCreateErr, frameworkHandle, "DADissenterCreate", "10.4")
	registerFunc(&_dADissenterGetStatus, &_dADissenterGetStatusErr, frameworkHandle, "DADissenterGetStatus", "10.4")
	registerFunc(&_dADissenterGetStatusString, &_dADissenterGetStatusStringErr, frameworkHandle, "DADissenterGetStatusString", "10.4")
	registerFunc(&_dARegisterDiskAppearedCallback, &_dARegisterDiskAppearedCallbackErr, frameworkHandle, "DARegisterDiskAppearedCallback", "10.4")
	registerFunc(&_dARegisterDiskDescriptionChangedCallback, &_dARegisterDiskDescriptionChangedCallbackErr, frameworkHandle, "DARegisterDiskDescriptionChangedCallback", "10.4")
	registerFunc(&_dARegisterDiskDisappearedCallback, &_dARegisterDiskDisappearedCallbackErr, frameworkHandle, "DARegisterDiskDisappearedCallback", "10.4")
	registerFunc(&_dARegisterDiskEjectApprovalCallback, &_dARegisterDiskEjectApprovalCallbackErr, frameworkHandle, "DARegisterDiskEjectApprovalCallback", "10.4")
	registerFunc(&_dARegisterDiskMountApprovalCallback, &_dARegisterDiskMountApprovalCallbackErr, frameworkHandle, "DARegisterDiskMountApprovalCallback", "10.4")
	registerFunc(&_dARegisterDiskPeekCallback, &_dARegisterDiskPeekCallbackErr, frameworkHandle, "DARegisterDiskPeekCallback", "10.4")
	registerFunc(&_dARegisterDiskUnmountApprovalCallback, &_dARegisterDiskUnmountApprovalCallbackErr, frameworkHandle, "DARegisterDiskUnmountApprovalCallback", "10.4")
	registerFunc(&_dASessionCreate, &_dASessionCreateErr, frameworkHandle, "DASessionCreate", "10.4")
	registerFunc(&_dASessionGetTypeID, &_dASessionGetTypeIDErr, frameworkHandle, "DASessionGetTypeID", "10.4")
	registerFunc(&_dASessionScheduleWithRunLoop, &_dASessionScheduleWithRunLoopErr, frameworkHandle, "DASessionScheduleWithRunLoop", "10.4")
	registerFunc(&_dASessionSetDispatchQueue, &_dASessionSetDispatchQueueErr, frameworkHandle, "DASessionSetDispatchQueue", "10.7")
	registerFunc(&_dASessionUnscheduleFromRunLoop, &_dASessionUnscheduleFromRunLoopErr, frameworkHandle, "DASessionUnscheduleFromRunLoop", "10.4")
	registerFunc(&_dAUnregisterApprovalCallback, &_dAUnregisterApprovalCallbackErr, frameworkHandle, "DAUnregisterApprovalCallback", "10.4")
	registerFunc(&_dAUnregisterCallback, &_dAUnregisterCallbackErr, frameworkHandle, "DAUnregisterCallback", "10.4")
}
