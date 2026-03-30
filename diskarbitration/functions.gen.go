// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

package diskarbitration

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: DiskArbitration: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: DiskArbitration: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: DiskArbitration: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _dAApprovalSessionCreate func(allocator corefoundation.CFAllocatorRef) DAApprovalSessionRef

// DAApprovalSessionCreate.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionCreate
func DAApprovalSessionCreate(allocator corefoundation.CFAllocatorRef) DAApprovalSessionRef {
	if _dAApprovalSessionCreate == nil {
		panic("DiskArbitration: symbol DAApprovalSessionCreate not loaded")
	}
	return _dAApprovalSessionCreate(allocator)
}

var _dAApprovalSessionGetTypeID func() uint

// DAApprovalSessionGetTypeID.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionGetTypeID
func DAApprovalSessionGetTypeID() uint {
	if _dAApprovalSessionGetTypeID == nil {
		panic("DiskArbitration: symbol DAApprovalSessionGetTypeID not loaded")
	}
	return _dAApprovalSessionGetTypeID()
}

var _dAApprovalSessionScheduleWithRunLoop func(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)

// DAApprovalSessionScheduleWithRunLoop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionScheduleWithRunLoop
func DAApprovalSessionScheduleWithRunLoop(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if _dAApprovalSessionScheduleWithRunLoop == nil {
		panic("DiskArbitration: symbol DAApprovalSessionScheduleWithRunLoop not loaded")
	}
	_dAApprovalSessionScheduleWithRunLoop(session, runLoop, runLoopMode)
}

var _dAApprovalSessionUnscheduleFromRunLoop func(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)

// DAApprovalSessionUnscheduleFromRunLoop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAApprovalSessionUnscheduleFromRunLoop
func DAApprovalSessionUnscheduleFromRunLoop(session DAApprovalSessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if _dAApprovalSessionUnscheduleFromRunLoop == nil {
		panic("DiskArbitration: symbol DAApprovalSessionUnscheduleFromRunLoop not loaded")
	}
	_dAApprovalSessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
}

var _dADiskClaim func(disk DADiskRef, options DADiskClaimOptions, release DADiskClaimReleaseCallback, releaseContext unsafe.Pointer, callback DADiskClaimCallback, callbackContext unsafe.Pointer)

// DADiskClaim claims the specified disk object for exclusive use.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskClaim(_:_:_:_:_:_:)
func DADiskClaim(disk DADiskRef, options DADiskClaimOptions, release DADiskClaimReleaseCallback, releaseContext unsafe.Pointer, callback DADiskClaimCallback, callbackContext unsafe.Pointer) {
	if _dADiskClaim == nil {
		panic("DiskArbitration: symbol DADiskClaim not loaded")
	}
	_dADiskClaim(disk, options, release, releaseContext, callback, callbackContext)
}

var _dADiskCopyDescription func(disk DADiskRef) corefoundation.CFDictionaryRef

// DADiskCopyDescription obtains the Disk Arbitration description of the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyDescription(_:)
func DADiskCopyDescription(disk DADiskRef) corefoundation.CFDictionaryRef {
	if _dADiskCopyDescription == nil {
		panic("DiskArbitration: symbol DADiskCopyDescription not loaded")
	}
	return _dADiskCopyDescription(disk)
}

var _dADiskCopyIOMedia func(disk DADiskRef) uintptr

// DADiskCopyIOMedia obtains the I/O Kit media object for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyIOMedia(_:)
func DADiskCopyIOMedia(disk DADiskRef) uintptr {
	if _dADiskCopyIOMedia == nil {
		panic("DiskArbitration: symbol DADiskCopyIOMedia not loaded")
	}
	return _dADiskCopyIOMedia(disk)
}

var _dADiskCopyWholeDisk func(disk DADiskRef) DADiskRef

// DADiskCopyWholeDisk obtain the associated whole disk object for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCopyWholeDisk(_:)
func DADiskCopyWholeDisk(disk DADiskRef) DADiskRef {
	if _dADiskCopyWholeDisk == nil {
		panic("DiskArbitration: symbol DADiskCopyWholeDisk not loaded")
	}
	return _dADiskCopyWholeDisk(disk)
}

var _dADiskCreateFromBSDName func(allocator corefoundation.CFAllocatorRef, session DASessionRef, name string) DADiskRef

// DADiskCreateFromBSDName creates a new disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromBSDName(_:_:_:)
func DADiskCreateFromBSDName(allocator corefoundation.CFAllocatorRef, session DASessionRef, name string) DADiskRef {
	if _dADiskCreateFromBSDName == nil {
		panic("DiskArbitration: symbol DADiskCreateFromBSDName not loaded")
	}
	return _dADiskCreateFromBSDName(allocator, session, name)
}

var _dADiskCreateFromIOMedia func(allocator corefoundation.CFAllocatorRef, session DASessionRef, media uintptr) DADiskRef

// DADiskCreateFromIOMedia creates a new disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromIOMedia(_:_:_:)
func DADiskCreateFromIOMedia(allocator corefoundation.CFAllocatorRef, session DASessionRef, media uintptr) DADiskRef {
	if _dADiskCreateFromIOMedia == nil {
		panic("DiskArbitration: symbol DADiskCreateFromIOMedia not loaded")
	}
	return _dADiskCreateFromIOMedia(allocator, session, media)
}

var _dADiskCreateFromVolumePath func(allocator corefoundation.CFAllocatorRef, session DASessionRef, path corefoundation.CFURLRef) DADiskRef

// DADiskCreateFromVolumePath creates a new disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskCreateFromVolumePath(_:_:_:)
func DADiskCreateFromVolumePath(allocator corefoundation.CFAllocatorRef, session DASessionRef, path corefoundation.CFURLRef) DADiskRef {
	if _dADiskCreateFromVolumePath == nil {
		panic("DiskArbitration: symbol DADiskCreateFromVolumePath not loaded")
	}
	return _dADiskCreateFromVolumePath(allocator, session, path)
}

var _dADiskEject func(disk DADiskRef, options uint, callback DADiskEjectCallback, context unsafe.Pointer)

// DADiskEject ejects the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskEject(_:_:_:_:)
func DADiskEject(disk DADiskRef, options uint, callback DADiskEjectCallback, context unsafe.Pointer) {
	if _dADiskEject == nil {
		panic("DiskArbitration: symbol DADiskEject not loaded")
	}
	_dADiskEject(disk, options, callback, context)
}

var _dADiskGetBSDName func(disk DADiskRef) *byte

// DADiskGetBSDName obtains the BSD device name for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskGetBSDName(_:)
func DADiskGetBSDName(disk DADiskRef) *byte {
	if _dADiskGetBSDName == nil {
		panic("DiskArbitration: symbol DADiskGetBSDName not loaded")
	}
	return _dADiskGetBSDName(disk)
}

var _dADiskGetOptions func(disk DADiskRef) DADiskOptions

// DADiskGetOptions obtains the options for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskGetOptions(_:)
func DADiskGetOptions(disk DADiskRef) DADiskOptions {
	if _dADiskGetOptions == nil {
		panic("DiskArbitration: symbol DADiskGetOptions not loaded")
	}
	return _dADiskGetOptions(disk)
}

var _dADiskGetTypeID func() uint

// DADiskGetTypeID returns the type identifier of all DADisk instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskGetTypeID()
func DADiskGetTypeID() uint {
	if _dADiskGetTypeID == nil {
		panic("DiskArbitration: symbol DADiskGetTypeID not loaded")
	}
	return _dADiskGetTypeID()
}

var _dADiskIsClaimed func(disk DADiskRef) bool

// DADiskIsClaimed reports whether or not the disk is claimed.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskIsClaimed(_:)
func DADiskIsClaimed(disk DADiskRef) bool {
	if _dADiskIsClaimed == nil {
		panic("DiskArbitration: symbol DADiskIsClaimed not loaded")
	}
	return _dADiskIsClaimed(disk)
}

var _dADiskMount func(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer)

// DADiskMount mounts the volume at the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMount(_:_:_:_:_:)
func DADiskMount(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer) {
	if _dADiskMount == nil {
		panic("DiskArbitration: symbol DADiskMount not loaded")
	}
	_dADiskMount(disk, path, options, callback, context)
}

var _dADiskMountWithArguments func(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer, arguments corefoundation.CFStringRef)

// DADiskMountWithArguments mounts the volume at the specified disk object, with the specified mount options.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskMountWithArguments(_:_:_:_:_:_:)
func DADiskMountWithArguments(disk DADiskRef, path corefoundation.CFURLRef, options DADiskMountOptions, callback DADiskMountCallback, context unsafe.Pointer, arguments corefoundation.CFStringRef) {
	if _dADiskMountWithArguments == nil {
		panic("DiskArbitration: symbol DADiskMountWithArguments not loaded")
	}
	_dADiskMountWithArguments(disk, path, options, callback, context, arguments)
}

var _dADiskRename func(disk DADiskRef, name corefoundation.CFStringRef, options DADiskRenameOptions, callback DADiskRenameCallback, context unsafe.Pointer)

// DADiskRename renames the volume at the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskRename(_:_:_:_:_:)
func DADiskRename(disk DADiskRef, name corefoundation.CFStringRef, options DADiskRenameOptions, callback DADiskRenameCallback, context unsafe.Pointer) {
	if _dADiskRename == nil {
		panic("DiskArbitration: symbol DADiskRename not loaded")
	}
	_dADiskRename(disk, name, options, callback, context)
}

var _dADiskSetOptions func(disk DADiskRef, options DADiskOptions, value bool) DAReturn

// DADiskSetOptions sets the options for the specified disk.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskSetOptions(_:_:_:)
func DADiskSetOptions(disk DADiskRef, options DADiskOptions, value bool) DAReturn {
	if _dADiskSetOptions == nil {
		panic("DiskArbitration: symbol DADiskSetOptions not loaded")
	}
	return _dADiskSetOptions(disk, options, value)
}

var _dADiskUnclaim func(disk DADiskRef)

// DADiskUnclaim unclaims the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnclaim(_:)
func DADiskUnclaim(disk DADiskRef) {
	if _dADiskUnclaim == nil {
		panic("DiskArbitration: symbol DADiskUnclaim not loaded")
	}
	_dADiskUnclaim(disk)
}

var _dADiskUnmount func(disk DADiskRef, options DADiskUnmountOptions, callback DADiskUnmountCallback, context unsafe.Pointer)

// DADiskUnmount unmounts the volume at the specified disk object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADiskUnmount(_:_:_:_:)
func DADiskUnmount(disk DADiskRef, options DADiskUnmountOptions, callback DADiskUnmountCallback, context unsafe.Pointer) {
	if _dADiskUnmount == nil {
		panic("DiskArbitration: symbol DADiskUnmount not loaded")
	}
	_dADiskUnmount(disk, options, callback, context)
}

var _dADissenterCreate func(allocator corefoundation.CFAllocatorRef, status DAReturn, string_ corefoundation.CFStringRef) DADissenterRef

// DADissenterCreate creates a new dissenter object.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenterCreate(_:_:_:)
func DADissenterCreate(allocator corefoundation.CFAllocatorRef, status DAReturn, string_ corefoundation.CFStringRef) DADissenterRef {
	if _dADissenterCreate == nil {
		panic("DiskArbitration: symbol DADissenterCreate not loaded")
	}
	return _dADissenterCreate(allocator, status, string_)
}

var _dADissenterGetStatus func(dissenter DADissenterRef) DAReturn

// DADissenterGetStatus obtains the return code.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatus(_:)
func DADissenterGetStatus(dissenter DADissenterRef) DAReturn {
	if _dADissenterGetStatus == nil {
		panic("DiskArbitration: symbol DADissenterGetStatus not loaded")
	}
	return _dADissenterGetStatus(dissenter)
}

var _dADissenterGetStatusString func(dissenter DADissenterRef) corefoundation.CFStringRef

// DADissenterGetStatusString obtains the return code string.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DADissenterGetStatusString(_:)
func DADissenterGetStatusString(dissenter DADissenterRef) corefoundation.CFStringRef {
	if _dADissenterGetStatusString == nil {
		panic("DiskArbitration: symbol DADissenterGetStatusString not loaded")
	}
	return _dADissenterGetStatusString(dissenter)
}

var _dARegisterDiskAppearedCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskAppearedCallback, context unsafe.Pointer)

// DARegisterDiskAppearedCallback registers a callback function to be called whenever a disk has appeared.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskAppearedCallback(_:_:_:_:)
func DARegisterDiskAppearedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskAppearedCallback, context unsafe.Pointer) {
	if _dARegisterDiskAppearedCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskAppearedCallback not loaded")
	}
	_dARegisterDiskAppearedCallback(session, match, callback, context)
}

var _dARegisterDiskDescriptionChangedCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, watch corefoundation.CFArrayRef, callback DADiskDescriptionChangedCallback, context unsafe.Pointer)

// DARegisterDiskDescriptionChangedCallback registers a callback function to be called whenever a disk description has changed.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDescriptionChangedCallback(_:_:_:_:_:)
func DARegisterDiskDescriptionChangedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, watch corefoundation.CFArrayRef, callback DADiskDescriptionChangedCallback, context unsafe.Pointer) {
	if _dARegisterDiskDescriptionChangedCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskDescriptionChangedCallback not loaded")
	}
	_dARegisterDiskDescriptionChangedCallback(session, match, watch, callback, context)
}

var _dARegisterDiskDisappearedCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskDisappearedCallback, context unsafe.Pointer)

// DARegisterDiskDisappearedCallback registers a callback function to be called whenever a disk has disappeared.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskDisappearedCallback(_:_:_:_:)
func DARegisterDiskDisappearedCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskDisappearedCallback, context unsafe.Pointer) {
	if _dARegisterDiskDisappearedCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskDisappearedCallback not loaded")
	}
	_dARegisterDiskDisappearedCallback(session, match, callback, context)
}

var _dARegisterDiskEjectApprovalCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskEjectApprovalCallback, context unsafe.Pointer)

// DARegisterDiskEjectApprovalCallback registers a callback function to be called whenever a volume is to be ejected.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskEjectApprovalCallback(_:_:_:_:)
func DARegisterDiskEjectApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskEjectApprovalCallback, context unsafe.Pointer) {
	if _dARegisterDiskEjectApprovalCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskEjectApprovalCallback not loaded")
	}
	_dARegisterDiskEjectApprovalCallback(session, match, callback, context)
}

var _dARegisterDiskMountApprovalCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskMountApprovalCallback, context unsafe.Pointer)

// DARegisterDiskMountApprovalCallback registers a callback function to be called whenever a volume is to be mounted.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskMountApprovalCallback(_:_:_:_:)
func DARegisterDiskMountApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskMountApprovalCallback, context unsafe.Pointer) {
	if _dARegisterDiskMountApprovalCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskMountApprovalCallback not loaded")
	}
	_dARegisterDiskMountApprovalCallback(session, match, callback, context)
}

var _dARegisterDiskPeekCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, order int, callback DADiskPeekCallback, context unsafe.Pointer)

// DARegisterDiskPeekCallback registers a callback function to be called whenever a disk has been probed.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskPeekCallback(_:_:_:_:_:)
func DARegisterDiskPeekCallback(session DASessionRef, match corefoundation.CFDictionaryRef, order int, callback DADiskPeekCallback, context unsafe.Pointer) {
	if _dARegisterDiskPeekCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskPeekCallback not loaded")
	}
	_dARegisterDiskPeekCallback(session, match, order, callback, context)
}

var _dARegisterDiskUnmountApprovalCallback func(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskUnmountApprovalCallback, context unsafe.Pointer)

// DARegisterDiskUnmountApprovalCallback registers a callback function to be called whenever a volume is to be unmounted.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DARegisterDiskUnmountApprovalCallback(_:_:_:_:)
func DARegisterDiskUnmountApprovalCallback(session DASessionRef, match corefoundation.CFDictionaryRef, callback DADiskUnmountApprovalCallback, context unsafe.Pointer) {
	if _dARegisterDiskUnmountApprovalCallback == nil {
		panic("DiskArbitration: symbol DARegisterDiskUnmountApprovalCallback not loaded")
	}
	_dARegisterDiskUnmountApprovalCallback(session, match, callback, context)
}

var _dASessionCreate func(allocator corefoundation.CFAllocatorRef) DASessionRef

// DASessionCreate creates a new session.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionCreate(_:)
func DASessionCreate(allocator corefoundation.CFAllocatorRef) DASessionRef {
	if _dASessionCreate == nil {
		panic("DiskArbitration: symbol DASessionCreate not loaded")
	}
	return _dASessionCreate(allocator)
}

var _dASessionGetTypeID func() uint

// DASessionGetTypeID returns the type identifier of all DASession instances.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionGetTypeID()
func DASessionGetTypeID() uint {
	if _dASessionGetTypeID == nil {
		panic("DiskArbitration: symbol DASessionGetTypeID not loaded")
	}
	return _dASessionGetTypeID()
}

var _dASessionScheduleWithRunLoop func(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)

// DASessionScheduleWithRunLoop schedules the session on a run loop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionScheduleWithRunLoop(_:_:_:)
func DASessionScheduleWithRunLoop(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if _dASessionScheduleWithRunLoop == nil {
		panic("DiskArbitration: symbol DASessionScheduleWithRunLoop not loaded")
	}
	_dASessionScheduleWithRunLoop(session, runLoop, runLoopMode)
}

var _dASessionSetDispatchQueue func(session DASessionRef, queue uintptr)

// DASessionSetDispatchQueue schedules the session on a dispatch queue.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionSetDispatchQueue(_:_:)
func DASessionSetDispatchQueue(session DASessionRef, queue dispatch.Queue) {
	if _dASessionSetDispatchQueue == nil {
		panic("DiskArbitration: symbol DASessionSetDispatchQueue not loaded")
	}
	_dASessionSetDispatchQueue(session, uintptr(queue.Handle()))
}

var _dASessionUnscheduleFromRunLoop func(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)

// DASessionUnscheduleFromRunLoop unschedules the session from a run loop.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DASessionUnscheduleFromRunLoop(_:_:_:)
func DASessionUnscheduleFromRunLoop(session DASessionRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if _dASessionUnscheduleFromRunLoop == nil {
		panic("DiskArbitration: symbol DASessionUnscheduleFromRunLoop not loaded")
	}
	_dASessionUnscheduleFromRunLoop(session, runLoop, runLoopMode)
}

var _dAUnregisterApprovalCallback func(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer)

// DAUnregisterApprovalCallback unregisters a registered callback function.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterApprovalCallback
func DAUnregisterApprovalCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if _dAUnregisterApprovalCallback == nil {
		panic("DiskArbitration: symbol DAUnregisterApprovalCallback not loaded")
	}
	_dAUnregisterApprovalCallback(session, callback, context)
}

var _dAUnregisterCallback func(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer)

// DAUnregisterCallback unregisters a registered callback function.
//
// See: https://developer.apple.com/documentation/DiskArbitration/DAUnregisterCallback(_:_:_:)
func DAUnregisterCallback(session DASessionRef, callback unsafe.Pointer, context unsafe.Pointer) {
	if _dAUnregisterCallback == nil {
		panic("DiskArbitration: symbol DAUnregisterCallback not loaded")
	}
	_dAUnregisterCallback(session, callback, context)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_dAApprovalSessionCreate, frameworkHandle, "DAApprovalSessionCreate")
	registerFunc(&_dAApprovalSessionGetTypeID, frameworkHandle, "DAApprovalSessionGetTypeID")
	registerFunc(&_dAApprovalSessionScheduleWithRunLoop, frameworkHandle, "DAApprovalSessionScheduleWithRunLoop")
	registerFunc(&_dAApprovalSessionUnscheduleFromRunLoop, frameworkHandle, "DAApprovalSessionUnscheduleFromRunLoop")
	registerFunc(&_dADiskClaim, frameworkHandle, "DADiskClaim")
	registerFunc(&_dADiskCopyDescription, frameworkHandle, "DADiskCopyDescription")
	registerFunc(&_dADiskCopyIOMedia, frameworkHandle, "DADiskCopyIOMedia")
	registerFunc(&_dADiskCopyWholeDisk, frameworkHandle, "DADiskCopyWholeDisk")
	registerFunc(&_dADiskCreateFromBSDName, frameworkHandle, "DADiskCreateFromBSDName")
	registerFunc(&_dADiskCreateFromIOMedia, frameworkHandle, "DADiskCreateFromIOMedia")
	registerFunc(&_dADiskCreateFromVolumePath, frameworkHandle, "DADiskCreateFromVolumePath")
	registerFunc(&_dADiskEject, frameworkHandle, "DADiskEject")
	registerFunc(&_dADiskGetBSDName, frameworkHandle, "DADiskGetBSDName")
	registerFunc(&_dADiskGetOptions, frameworkHandle, "DADiskGetOptions")
	registerFunc(&_dADiskGetTypeID, frameworkHandle, "DADiskGetTypeID")
	registerFunc(&_dADiskIsClaimed, frameworkHandle, "DADiskIsClaimed")
	registerFunc(&_dADiskMount, frameworkHandle, "DADiskMount")
	registerFunc(&_dADiskMountWithArguments, frameworkHandle, "DADiskMountWithArguments")
	registerFunc(&_dADiskRename, frameworkHandle, "DADiskRename")
	registerFunc(&_dADiskSetOptions, frameworkHandle, "DADiskSetOptions")
	registerFunc(&_dADiskUnclaim, frameworkHandle, "DADiskUnclaim")
	registerFunc(&_dADiskUnmount, frameworkHandle, "DADiskUnmount")
	registerFunc(&_dADissenterCreate, frameworkHandle, "DADissenterCreate")
	registerFunc(&_dADissenterGetStatus, frameworkHandle, "DADissenterGetStatus")
	registerFunc(&_dADissenterGetStatusString, frameworkHandle, "DADissenterGetStatusString")
	registerFunc(&_dARegisterDiskAppearedCallback, frameworkHandle, "DARegisterDiskAppearedCallback")
	registerFunc(&_dARegisterDiskDescriptionChangedCallback, frameworkHandle, "DARegisterDiskDescriptionChangedCallback")
	registerFunc(&_dARegisterDiskDisappearedCallback, frameworkHandle, "DARegisterDiskDisappearedCallback")
	registerFunc(&_dARegisterDiskEjectApprovalCallback, frameworkHandle, "DARegisterDiskEjectApprovalCallback")
	registerFunc(&_dARegisterDiskMountApprovalCallback, frameworkHandle, "DARegisterDiskMountApprovalCallback")
	registerFunc(&_dARegisterDiskPeekCallback, frameworkHandle, "DARegisterDiskPeekCallback")
	registerFunc(&_dARegisterDiskUnmountApprovalCallback, frameworkHandle, "DARegisterDiskUnmountApprovalCallback")
	registerFunc(&_dASessionCreate, frameworkHandle, "DASessionCreate")
	registerFunc(&_dASessionGetTypeID, frameworkHandle, "DASessionGetTypeID")
	registerFunc(&_dASessionScheduleWithRunLoop, frameworkHandle, "DASessionScheduleWithRunLoop")
	registerFunc(&_dASessionSetDispatchQueue, frameworkHandle, "DASessionSetDispatchQueue")
	registerFunc(&_dASessionUnscheduleFromRunLoop, frameworkHandle, "DASessionUnscheduleFromRunLoop")
	registerFunc(&_dAUnregisterApprovalCallback, frameworkHandle, "DAUnregisterApprovalCallback")
	registerFunc(&_dAUnregisterCallback, frameworkHandle, "DAUnregisterCallback")
}
