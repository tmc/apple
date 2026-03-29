// Code generated from Apple documentation for coreservices. DO NOT EDIT.

package coreservices

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/diskarbitration"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/security"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: coreservices: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: coreservices: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: coreservices: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _aEBuildAppleEvent func(arg0 uint32, arg1 uint32, arg2 uint32, arg3 corefoundation.CGSize, arg4 int16, arg5 int32, arg6 AEDesc, arg7 AEBuildError, arg8 int8) int32

// AEBuildAppleEvent constructs an entire Apple event in a single call.
//
// See: https://developer.apple.com/documentation/coreservices/1573757-aebuildappleevent
func AEBuildAppleEvent(arg0 uint32, arg1 uint32, arg2 uint32, arg3 corefoundation.CGSize, arg4 int16, arg5 int32, arg6 AEDesc, arg7 AEBuildError, arg8 int8) int32 {
	if _aEBuildAppleEvent == nil {
		panic("coreservices: symbol AEBuildAppleEvent not loaded")
	}
	return _aEBuildAppleEvent(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _aEBuildDesc func(arg0 AEDesc, arg1 AEBuildError, arg2 int8) int32

// AEBuildDesc provides a facility for compiling AEBuild descriptor strings into Apple event descriptors ([AEDesc]).
//
// See: https://developer.apple.com/documentation/coreservices/1573758-aebuilddesc
func AEBuildDesc(arg0 AEDesc, arg1 AEBuildError, arg2 int8) int32 {
	if _aEBuildDesc == nil {
		panic("coreservices: symbol AEBuildDesc not loaded")
	}
	return _aEBuildDesc(arg0, arg1, arg2)
}

var _aEBuildParameters func(arg0 AEDesc, arg1 AEBuildError, arg2 int8) int32

// AEBuildParameters adds additional parameters or attributes to an existing Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/1573755-aebuildparameters
func AEBuildParameters(arg0 AEDesc, arg1 AEBuildError, arg2 int8) int32 {
	if _aEBuildParameters == nil {
		panic("coreservices: symbol AEBuildParameters not loaded")
	}
	return _aEBuildParameters(arg0, arg1, arg2)
}

var _aECallObjectAccessor func(desiredClass uint32, containerToken unsafe.Pointer, containerClass uint32, keyForm uint32, keyData unsafe.Pointer, token unsafe.Pointer) int16

// AECallObjectAccessor invokes the appropriate object accessor function for a specific desired type and container type.
//
// See: https://developer.apple.com/documentation/coreservices/1447059-aecallobjectaccessor
func AECallObjectAccessor(desiredClass uint32, containerToken unsafe.Pointer, containerClass uint32, keyForm uint32, keyData unsafe.Pointer, token unsafe.Pointer) int16 {
	if _aECallObjectAccessor == nil {
		panic("coreservices: symbol AECallObjectAccessor not loaded")
	}
	return _aECallObjectAccessor(desiredClass, containerToken, containerClass, keyForm, keyData, token)
}

var _aECheckIsRecord func(theDesc unsafe.Pointer) bool

// AECheckIsRecord determines whether a descriptor is truly an [AERecord].
//
// See: https://developer.apple.com/documentation/coreservices/1444011-aecheckisrecord
func AECheckIsRecord(theDesc unsafe.Pointer) bool {
	if _aECheckIsRecord == nil {
		panic("coreservices: symbol AECheckIsRecord not loaded")
	}
	return _aECheckIsRecord(theDesc)
}

var _aECoerceDesc func(theAEDesc unsafe.Pointer, toType uint32, result unsafe.Pointer) int16

// AECoerceDesc coerces the data in a descriptor to another descriptor type and creates a descriptor containing the newly coerced data.
//
// See: https://developer.apple.com/documentation/coreservices/1446519-aecoercedesc
func AECoerceDesc(theAEDesc unsafe.Pointer, toType uint32, result unsafe.Pointer) int16 {
	if _aECoerceDesc == nil {
		panic("coreservices: symbol AECoerceDesc not loaded")
	}
	return _aECoerceDesc(theAEDesc, toType, result)
}

var _aECoercePtr func(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, toType uint32, result unsafe.Pointer) int16

// AECoercePtr coerces data to a desired descriptor type and creates a descriptor containing the newly coerced data.
//
// See: https://developer.apple.com/documentation/coreservices/1441846-aecoerceptr
func AECoercePtr(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, toType uint32, result unsafe.Pointer) int16 {
	if _aECoercePtr == nil {
		panic("coreservices: symbol AECoercePtr not loaded")
	}
	return _aECoercePtr(typeCode, dataPtr, dataSize, toType, result)
}

var _aECompareDesc func(desc1 unsafe.Pointer, desc2 unsafe.Pointer, resultP unsafe.Pointer) int32

// AECompareDesc.
//
// See: https://developer.apple.com/documentation/coreservices/1448782-aecomparedesc
func AECompareDesc(desc1 unsafe.Pointer, desc2 unsafe.Pointer, resultP unsafe.Pointer) int32 {
	if _aECompareDesc == nil {
		panic("coreservices: symbol AECompareDesc not loaded")
	}
	return _aECompareDesc(desc1, desc2, resultP)
}

var _aECountItems func(theAEDescList unsafe.Pointer, theCount unsafe.Pointer) int16

// AECountItems counts the number of descriptors in a descriptor list.
//
// See: https://developer.apple.com/documentation/coreservices/1449533-aecountitems
func AECountItems(theAEDescList unsafe.Pointer, theCount unsafe.Pointer) int16 {
	if _aECountItems == nil {
		panic("coreservices: symbol AECountItems not loaded")
	}
	return _aECountItems(theAEDescList, theCount)
}

var _aECreateAppleEvent func(theAEEventClass uint32, theAEEventID uint32, target unsafe.Pointer, returnID AEReturnID, transactionID AETransactionID, result *AEDesc) int16

// AECreateAppleEvent creates an Apple event with several important attributes but no parameters.
//
// See: https://developer.apple.com/documentation/coreservices/1448525-aecreateappleevent
func AECreateAppleEvent(theAEEventClass uint32, theAEEventID uint32, target unsafe.Pointer, returnID AEReturnID, transactionID AETransactionID, result *AEDesc) int16 {
	if _aECreateAppleEvent == nil {
		panic("coreservices: symbol AECreateAppleEvent not loaded")
	}
	return _aECreateAppleEvent(theAEEventClass, theAEEventID, target, returnID, transactionID, result)
}

var _aECreateDesc func(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, result unsafe.Pointer) int16

// AECreateDesc creates a new descriptor that incorporates the specified data.
//
// See: https://developer.apple.com/documentation/coreservices/1448535-aecreatedesc
func AECreateDesc(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, result unsafe.Pointer) int16 {
	if _aECreateDesc == nil {
		panic("coreservices: symbol AECreateDesc not loaded")
	}
	return _aECreateDesc(typeCode, dataPtr, dataSize, result)
}

var _aECreateDescFromExternalPtr func(descriptorType uint32, dataPtr unsafe.Pointer, dataLength corefoundation.CGSize, disposeCallback AEDisposeExternalUPP, disposeRefcon uintptr, theDesc unsafe.Pointer) int32

// AECreateDescFromExternalPtr creates a new descriptor that uses a memory buffer supplied by the caller.
//
// See: https://developer.apple.com/documentation/coreservices/1446239-aecreatedescfromexternalptr
func AECreateDescFromExternalPtr(descriptorType uint32, dataPtr unsafe.Pointer, dataLength corefoundation.CGSize, disposeCallback AEDisposeExternalUPP, disposeRefcon uintptr, theDesc unsafe.Pointer) int32 {
	if _aECreateDescFromExternalPtr == nil {
		panic("coreservices: symbol AECreateDescFromExternalPtr not loaded")
	}
	return _aECreateDescFromExternalPtr(descriptorType, dataPtr, dataLength, disposeCallback, disposeRefcon, theDesc)
}

var _aECreateList func(factoringPtr unsafe.Pointer, factoredSize corefoundation.CGSize, isRecord unsafe.Pointer, resultList unsafe.Pointer) int16

// AECreateList creates an empty descriptor list or Apple event record.
//
// See: https://developer.apple.com/documentation/coreservices/1448643-aecreatelist
func AECreateList(factoringPtr unsafe.Pointer, factoredSize corefoundation.CGSize, isRecord unsafe.Pointer, resultList unsafe.Pointer) int16 {
	if _aECreateList == nil {
		panic("coreservices: symbol AECreateList not loaded")
	}
	return _aECreateList(factoringPtr, factoredSize, isRecord, resultList)
}

var _aECreateRemoteProcessResolver func(allocator corefoundation.CFAllocatorRef, url corefoundation.CFURLRef) AERemoteProcessResolverRef

// AECreateRemoteProcessResolver creates an object for resolving a list of remote processes.
//
// See: https://developer.apple.com/documentation/coreservices/1445692-aecreateremoteprocessresolver
func AECreateRemoteProcessResolver(allocator corefoundation.CFAllocatorRef, url corefoundation.CFURLRef) AERemoteProcessResolverRef {
	if _aECreateRemoteProcessResolver == nil {
		panic("coreservices: symbol AECreateRemoteProcessResolver not loaded")
	}
	return _aECreateRemoteProcessResolver(allocator, url)
}

var _aEDecodeMessage func(header unsafe.Pointer, event *AEDesc, reply *AEDesc) int32

// AEDecodeMessage decodes a Mach message and converts it into an Apple event and its related reply.
//
// See: https://developer.apple.com/documentation/coreservices/1447827-aedecodemessage
func AEDecodeMessage(header unsafe.Pointer, event *AEDesc, reply *AEDesc) int32 {
	if _aEDecodeMessage == nil {
		panic("coreservices: symbol AEDecodeMessage not loaded")
	}
	return _aEDecodeMessage(header, event, reply)
}

var _aEDeleteItem func(theAEDescList unsafe.Pointer, index unsafe.Pointer) int16

// AEDeleteItem deletes a descriptor from a descriptor list, causing all subsequent descriptors to move up one place.
//
// See: https://developer.apple.com/documentation/coreservices/1447164-aedeleteitem
func AEDeleteItem(theAEDescList unsafe.Pointer, index unsafe.Pointer) int16 {
	if _aEDeleteItem == nil {
		panic("coreservices: symbol AEDeleteItem not loaded")
	}
	return _aEDeleteItem(theAEDescList, index)
}

var _aEDeleteParam func(theAppleEvent *AEDesc, theAEKeyword uint32) int16

// AEDeleteParam deletes a keyword-specified parameter from an Apple event record.
//
// See: https://developer.apple.com/documentation/coreservices/1444338-aedeleteparam
func AEDeleteParam(theAppleEvent *AEDesc, theAEKeyword uint32) int16 {
	if _aEDeleteParam == nil {
		panic("coreservices: symbol AEDeleteParam not loaded")
	}
	return _aEDeleteParam(theAppleEvent, theAEKeyword)
}

var _aEDeterminePermissionToAutomateTarget func(target unsafe.Pointer, theAEEventClass uint32, theAEEventID uint32, askUserIfNeeded unsafe.Pointer) int32

// AEDeterminePermissionToAutomateTarget.
//
// See: https://developer.apple.com/documentation/coreservices/3025784-aedeterminepermissiontoautomatet
func AEDeterminePermissionToAutomateTarget(target unsafe.Pointer, theAEEventClass uint32, theAEEventID uint32, askUserIfNeeded unsafe.Pointer) int32 {
	if _aEDeterminePermissionToAutomateTarget == nil {
		panic("coreservices: symbol AEDeterminePermissionToAutomateTarget not loaded")
	}
	return _aEDeterminePermissionToAutomateTarget(target, theAEEventClass, theAEEventID, askUserIfNeeded)
}

var _aEDisposeDesc func(theAEDesc unsafe.Pointer) int16

// AEDisposeDesc deallocates the memory used by a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1444208-aedisposedesc
func AEDisposeDesc(theAEDesc unsafe.Pointer) int16 {
	if _aEDisposeDesc == nil {
		panic("coreservices: symbol AEDisposeDesc not loaded")
	}
	return _aEDisposeDesc(theAEDesc)
}

var _aEDisposeRemoteProcessResolver func(ref AERemoteProcessResolverRef)

// AEDisposeRemoteProcessResolver disposes of an [AERemoteProcessResolverRef].
//
// See: https://developer.apple.com/documentation/coreservices/1442572-aedisposeremoteprocessresolver
func AEDisposeRemoteProcessResolver(ref AERemoteProcessResolverRef) {
	if _aEDisposeRemoteProcessResolver == nil {
		panic("coreservices: symbol AEDisposeRemoteProcessResolver not loaded")
	}
	_aEDisposeRemoteProcessResolver(ref)
}

var _aEDisposeToken func(theToken unsafe.Pointer) int16

// AEDisposeToken deallocates the memory used by a token.
//
// See: https://developer.apple.com/documentation/coreservices/1446783-aedisposetoken
func AEDisposeToken(theToken unsafe.Pointer) int16 {
	if _aEDisposeToken == nil {
		panic("coreservices: symbol AEDisposeToken not loaded")
	}
	return _aEDisposeToken(theToken)
}

var _aEDuplicateDesc func(theAEDesc unsafe.Pointer, result unsafe.Pointer) int16

// AEDuplicateDesc creates a copy of a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1442661-aeduplicatedesc
func AEDuplicateDesc(theAEDesc unsafe.Pointer, result unsafe.Pointer) int16 {
	if _aEDuplicateDesc == nil {
		panic("coreservices: symbol AEDuplicateDesc not loaded")
	}
	return _aEDuplicateDesc(theAEDesc, result)
}

var _aEFlattenDesc func(theAEDesc unsafe.Pointer, buffer coreimage.Ptr, bufferSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int32

// AEFlattenDesc flattens the specified descriptor and stores the data in the supplied buffer.
//
// See: https://developer.apple.com/documentation/coreservices/1441808-aeflattendesc
func AEFlattenDesc(theAEDesc unsafe.Pointer, buffer coreimage.Ptr, bufferSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int32 {
	if _aEFlattenDesc == nil {
		panic("coreservices: symbol AEFlattenDesc not loaded")
	}
	return _aEFlattenDesc(theAEDesc, buffer, bufferSize, actualSize)
}

var _aEGetArray func(theAEDescList unsafe.Pointer, arrayType AEArrayType, arrayPtr AEArrayDataPointer, maximumSize corefoundation.CGSize, itemType *uint32, itemSize *corefoundation.CGSize, itemCount unsafe.Pointer) int16

// AEGetArray extracts data from an Apple event array created with the [AEPutArray] function and stores it as a standard array of fixed size items in the specified buffer.
//
// See: https://developer.apple.com/documentation/coreservices/1445720-aegetarray
func AEGetArray(theAEDescList unsafe.Pointer, arrayType AEArrayType, arrayPtr AEArrayDataPointer, maximumSize corefoundation.CGSize, itemType *uint32, itemSize *corefoundation.CGSize, itemCount unsafe.Pointer) int16 {
	if _aEGetArray == nil {
		panic("coreservices: symbol AEGetArray not loaded")
	}
	return _aEGetArray(theAEDescList, arrayType, arrayPtr, maximumSize, itemType, itemSize, itemCount)
}

var _aEGetAttributeDesc func(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, result unsafe.Pointer) int16

// AEGetAttributeDesc gets a copy of the descriptor for a specified Apple event attribute from an Apple event; typically used when your application needs to pass the descriptor on to another function.
//
// See: https://developer.apple.com/documentation/coreservices/1450314-aegetattributedesc
func AEGetAttributeDesc(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, result unsafe.Pointer) int16 {
	if _aEGetAttributeDesc == nil {
		panic("coreservices: symbol AEGetAttributeDesc not loaded")
	}
	return _aEGetAttributeDesc(theAppleEvent, theAEKeyword, desiredType, result)
}

var _aEGetAttributePtr func(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, typeCode *uint32, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int16

// AEGetAttributePtr gets a copy of the data for a specified Apple event attribute from an Apple event; typically used when your application needs to work with the data directly.
//
// See: https://developer.apple.com/documentation/coreservices/1445109-aegetattributeptr
func AEGetAttributePtr(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, typeCode *uint32, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int16 {
	if _aEGetAttributePtr == nil {
		panic("coreservices: symbol AEGetAttributePtr not loaded")
	}
	return _aEGetAttributePtr(theAppleEvent, theAEKeyword, desiredType, typeCode, dataPtr, maximumSize, actualSize)
}

var _aEGetCoercionHandler func(fromType uint32, toType uint32, handler unsafe.Pointer, handlerRefcon *uintptr, fromTypeIsDesc unsafe.Pointer, isSysHandler unsafe.Pointer) int16

// AEGetCoercionHandler gets the coercion handler for a specified descriptor type.
//
// See: https://developer.apple.com/documentation/coreservices/1445348-aegetcoercionhandler
func AEGetCoercionHandler(fromType uint32, toType uint32, handler unsafe.Pointer, handlerRefcon *uintptr, fromTypeIsDesc unsafe.Pointer, isSysHandler unsafe.Pointer) int16 {
	if _aEGetCoercionHandler == nil {
		panic("coreservices: symbol AEGetCoercionHandler not loaded")
	}
	return _aEGetCoercionHandler(fromType, toType, handler, handlerRefcon, fromTypeIsDesc, isSysHandler)
}

var _aEGetDescData func(theAEDesc unsafe.Pointer, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize) int16

// AEGetDescData gets the data from the specified descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1444427-aegetdescdata
func AEGetDescData(theAEDesc unsafe.Pointer, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize) int16 {
	if _aEGetDescData == nil {
		panic("coreservices: symbol AEGetDescData not loaded")
	}
	return _aEGetDescData(theAEDesc, dataPtr, maximumSize)
}

var _aEGetDescDataRange func(dataDesc unsafe.Pointer, buffer unsafe.Pointer, offset corefoundation.CGSize, length corefoundation.CGSize) int32

// AEGetDescDataRange retrieves a specified series of bytes from the specified descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1446560-aegetdescdatarange
func AEGetDescDataRange(dataDesc unsafe.Pointer, buffer unsafe.Pointer, offset corefoundation.CGSize, length corefoundation.CGSize) int32 {
	if _aEGetDescDataRange == nil {
		panic("coreservices: symbol AEGetDescDataRange not loaded")
	}
	return _aEGetDescDataRange(dataDesc, buffer, offset, length)
}

var _aEGetDescDataSize func(theAEDesc unsafe.Pointer) corefoundation.CGSize

// AEGetDescDataSize gets the size, in bytes, of the data in the specified descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1450119-aegetdescdatasize
func AEGetDescDataSize(theAEDesc unsafe.Pointer) corefoundation.CGSize {
	if _aEGetDescDataSize == nil {
		panic("coreservices: symbol AEGetDescDataSize not loaded")
	}
	return _aEGetDescDataSize(theAEDesc)
}

var _aEGetEventHandler func(theAEEventClass uint32, theAEEventID uint32, handler unsafe.Pointer, handlerRefcon *uintptr, isSysHandler unsafe.Pointer) int16

// AEGetEventHandler gets an event handler from an Apple event dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1445631-aegeteventhandler
func AEGetEventHandler(theAEEventClass uint32, theAEEventID uint32, handler unsafe.Pointer, handlerRefcon *uintptr, isSysHandler unsafe.Pointer) int16 {
	if _aEGetEventHandler == nil {
		panic("coreservices: symbol AEGetEventHandler not loaded")
	}
	return _aEGetEventHandler(theAEEventClass, theAEEventID, handler, handlerRefcon, isSysHandler)
}

var _aEGetNthDesc func(theAEDescList unsafe.Pointer, index unsafe.Pointer, desiredType uint32, theAEKeyword unsafe.Pointer, result unsafe.Pointer) int16

// AEGetNthDesc copies a descriptor from a specified position in a descriptor list into a specified descriptor; typically used when your application needs to pass the extracted data to another function as a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1448326-aegetnthdesc
func AEGetNthDesc(theAEDescList unsafe.Pointer, index unsafe.Pointer, desiredType uint32, theAEKeyword unsafe.Pointer, result unsafe.Pointer) int16 {
	if _aEGetNthDesc == nil {
		panic("coreservices: symbol AEGetNthDesc not loaded")
	}
	return _aEGetNthDesc(theAEDescList, index, desiredType, theAEKeyword, result)
}

var _aEGetNthPtr func(theAEDescList unsafe.Pointer, index unsafe.Pointer, desiredType uint32, theAEKeyword unsafe.Pointer, typeCode *uint32, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int16

// AEGetNthPtr gets a copy of the data from a descriptor at a specified position in a descriptor list; typically used when your application needs to work with the extracted data directly.
//
// See: https://developer.apple.com/documentation/coreservices/1447539-aegetnthptr
func AEGetNthPtr(theAEDescList unsafe.Pointer, index unsafe.Pointer, desiredType uint32, theAEKeyword unsafe.Pointer, typeCode *uint32, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int16 {
	if _aEGetNthPtr == nil {
		panic("coreservices: symbol AEGetNthPtr not loaded")
	}
	return _aEGetNthPtr(theAEDescList, index, desiredType, theAEKeyword, typeCode, dataPtr, maximumSize, actualSize)
}

var _aEGetObjectAccessor func(desiredClass uint32, containerType uint32, accessor unsafe.Pointer, accessorRefcon *uintptr, isSysHandler unsafe.Pointer) int16

// AEGetObjectAccessor gets an object accessor function from an object accessor dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1449054-aegetobjectaccessor
func AEGetObjectAccessor(desiredClass uint32, containerType uint32, accessor unsafe.Pointer, accessorRefcon *uintptr, isSysHandler unsafe.Pointer) int16 {
	if _aEGetObjectAccessor == nil {
		panic("coreservices: symbol AEGetObjectAccessor not loaded")
	}
	return _aEGetObjectAccessor(desiredClass, containerType, accessor, accessorRefcon, isSysHandler)
}

var _aEGetParamDesc func(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, result unsafe.Pointer) int16

// AEGetParamDesc gets a copy of the descriptor for a keyword-specified Apple event parameter from an Apple event or an Apple event record.
//
// See: https://developer.apple.com/documentation/coreservices/1449233-aegetparamdesc
func AEGetParamDesc(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, result unsafe.Pointer) int16 {
	if _aEGetParamDesc == nil {
		panic("coreservices: symbol AEGetParamDesc not loaded")
	}
	return _aEGetParamDesc(theAppleEvent, theAEKeyword, desiredType, result)
}

var _aEGetParamPtr func(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, actualType *uint32, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int16

// AEGetParamPtr gets a copy of the data for a specified Apple event parameter from an Apple event or an Apple event record.
//
// See: https://developer.apple.com/documentation/coreservices/1444069-aegetparamptr
func AEGetParamPtr(theAppleEvent *AEDesc, theAEKeyword uint32, desiredType uint32, actualType *uint32, dataPtr unsafe.Pointer, maximumSize corefoundation.CGSize, actualSize *corefoundation.CGSize) int16 {
	if _aEGetParamPtr == nil {
		panic("coreservices: symbol AEGetParamPtr not loaded")
	}
	return _aEGetParamPtr(theAppleEvent, theAEKeyword, desiredType, actualType, dataPtr, maximumSize, actualSize)
}

var _aEGetRegisteredMachPort func() uint32

// AEGetRegisteredMachPort returns the Mach port (in the form of a `mach_port_t`) that was registered with the bootstrap server for this process.
//
// See: https://developer.apple.com/documentation/coreservices/1449736-aegetregisteredmachport
func AEGetRegisteredMachPort() uint32 {
	if _aEGetRegisteredMachPort == nil {
		panic("coreservices: symbol AEGetRegisteredMachPort not loaded")
	}
	return _aEGetRegisteredMachPort()
}

var _aEGetSpecialHandler func(functionClass uint32, handler unsafe.Pointer, isSysHandler unsafe.Pointer) int16

// AEGetSpecialHandler gets a specified handler from a special handler dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1444274-aegetspecialhandler
func AEGetSpecialHandler(functionClass uint32, handler unsafe.Pointer, isSysHandler unsafe.Pointer) int16 {
	if _aEGetSpecialHandler == nil {
		panic("coreservices: symbol AEGetSpecialHandler not loaded")
	}
	return _aEGetSpecialHandler(functionClass, handler, isSysHandler)
}

var _aEInitializeDesc func(desc unsafe.Pointer)

// AEInitializeDesc initializes a new descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1446047-aeinitializedesc
func AEInitializeDesc(desc unsafe.Pointer) {
	if _aEInitializeDesc == nil {
		panic("coreservices: symbol AEInitializeDesc not loaded")
	}
	_aEInitializeDesc(desc)
}

var _aEInstallCoercionHandler func(fromType uint32, toType uint32, handler AECoercionHandlerUPP, handlerRefcon uintptr, fromTypeIsDesc unsafe.Pointer, isSysHandler unsafe.Pointer) int16

// AEInstallCoercionHandler installs a coercion handler in either the application or system coercion handler dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1445548-aeinstallcoercionhandler
func AEInstallCoercionHandler(fromType uint32, toType uint32, handler AECoercionHandlerUPP, handlerRefcon uintptr, fromTypeIsDesc unsafe.Pointer, isSysHandler unsafe.Pointer) int16 {
	if _aEInstallCoercionHandler == nil {
		panic("coreservices: symbol AEInstallCoercionHandler not loaded")
	}
	return _aEInstallCoercionHandler(fromType, toType, handler, handlerRefcon, fromTypeIsDesc, isSysHandler)
}

var _aEInstallEventHandler func(theAEEventClass uint32, theAEEventID uint32, handler AEEventHandlerUPP, handlerRefcon uintptr, isSysHandler unsafe.Pointer) int16

// AEInstallEventHandler adds an entry for an event handler to an Apple event dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1448596-aeinstalleventhandler
func AEInstallEventHandler(theAEEventClass uint32, theAEEventID uint32, handler AEEventHandlerUPP, handlerRefcon uintptr, isSysHandler unsafe.Pointer) int16 {
	if _aEInstallEventHandler == nil {
		panic("coreservices: symbol AEInstallEventHandler not loaded")
	}
	return _aEInstallEventHandler(theAEEventClass, theAEEventID, handler, handlerRefcon, isSysHandler)
}

var _aEInstallObjectAccessor func(desiredClass uint32, containerType uint32, theAccessor OSLAccessorUPP, accessorRefcon uintptr, isSysHandler unsafe.Pointer) int16

// AEInstallObjectAccessor adds or replaces an entry for an object accessor function to an object accessor dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1447905-aeinstallobjectaccessor
func AEInstallObjectAccessor(desiredClass uint32, containerType uint32, theAccessor OSLAccessorUPP, accessorRefcon uintptr, isSysHandler unsafe.Pointer) int16 {
	if _aEInstallObjectAccessor == nil {
		panic("coreservices: symbol AEInstallObjectAccessor not loaded")
	}
	return _aEInstallObjectAccessor(desiredClass, containerType, theAccessor, accessorRefcon, isSysHandler)
}

var _aEInstallSpecialHandler func(functionClass uint32, handler AEEventHandlerUPP, isSysHandler unsafe.Pointer) int16

// AEInstallSpecialHandler installs a callback function in a special handler dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1445532-aeinstallspecialhandler
func AEInstallSpecialHandler(functionClass uint32, handler AEEventHandlerUPP, isSysHandler unsafe.Pointer) int16 {
	if _aEInstallSpecialHandler == nil {
		panic("coreservices: symbol AEInstallSpecialHandler not loaded")
	}
	return _aEInstallSpecialHandler(functionClass, handler, isSysHandler)
}

var _aEManagerInfo func(keyWord uint32, result unsafe.Pointer) int16

// AEManagerInfo provides information about the version of the Apple Event Manager currently available or the number of processes that are currently recording Apple events.
//
// See: https://developer.apple.com/documentation/coreservices/1449373-aemanagerinfo
func AEManagerInfo(keyWord uint32, result unsafe.Pointer) int16 {
	if _aEManagerInfo == nil {
		panic("coreservices: symbol AEManagerInfo not loaded")
	}
	return _aEManagerInfo(keyWord, result)
}

var _aEObjectInit func() int16

// AEObjectInit initializes the Object Support Library.
//
// See: https://developer.apple.com/documentation/coreservices/1447372-aeobjectinit
func AEObjectInit() int16 {
	if _aEObjectInit == nil {
		panic("coreservices: symbol AEObjectInit not loaded")
	}
	return _aEObjectInit()
}

var _aEPrintDescToHandle func(desc unsafe.Pointer, result uintptr) int32

// AEPrintDescToHandle provides a pretty printer facility for displaying the contents of Apple event descriptors.
//
// See: https://developer.apple.com/documentation/coreservices/1445158-aeprintdesctohandle
func AEPrintDescToHandle(desc unsafe.Pointer, result uintptr) int32 {
	if _aEPrintDescToHandle == nil {
		panic("coreservices: symbol AEPrintDescToHandle not loaded")
	}
	return _aEPrintDescToHandle(desc, result)
}

var _aEProcessMessage func(header unsafe.Pointer) int32

// AEProcessMessage decodes and dispatches a low level Mach message event to an event handler, including packaging and returning the reply to the sender.
//
// See: https://developer.apple.com/documentation/coreservices/1444387-aeprocessmessage
func AEProcessMessage(header unsafe.Pointer) int32 {
	if _aEProcessMessage == nil {
		panic("coreservices: symbol AEProcessMessage not loaded")
	}
	return _aEProcessMessage(header)
}

var _aEPutArray func(theAEDescList unsafe.Pointer, arrayType AEArrayType, arrayPtr unsafe.Pointer, itemType uint32, itemSize corefoundation.CGSize, itemCount unsafe.Pointer) int16

// AEPutArray inserts the data for an Apple event array into a descriptor list, replacing any previous descriptors in the list.
//
// See: https://developer.apple.com/documentation/coreservices/1442535-aeputarray
func AEPutArray(theAEDescList unsafe.Pointer, arrayType AEArrayType, arrayPtr unsafe.Pointer, itemType uint32, itemSize corefoundation.CGSize, itemCount unsafe.Pointer) int16 {
	if _aEPutArray == nil {
		panic("coreservices: symbol AEPutArray not loaded")
	}
	return _aEPutArray(theAEDescList, arrayType, arrayPtr, itemType, itemSize, itemCount)
}

var _aEPutAttributeDesc func(theAppleEvent *AEDesc, theAEKeyword uint32, theAEDesc unsafe.Pointer) int16

// AEPutAttributeDesc adds a descriptor and a keyword to an Apple event as an attribute.
//
// See: https://developer.apple.com/documentation/coreservices/1441790-aeputattributedesc
func AEPutAttributeDesc(theAppleEvent *AEDesc, theAEKeyword uint32, theAEDesc unsafe.Pointer) int16 {
	if _aEPutAttributeDesc == nil {
		panic("coreservices: symbol AEPutAttributeDesc not loaded")
	}
	return _aEPutAttributeDesc(theAppleEvent, theAEKeyword, theAEDesc)
}

var _aEPutAttributePtr func(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize) int16

// AEPutAttributePtr adds a pointer to data, a descriptor type, and a keyword to an Apple event as an attribute.
//
// See: https://developer.apple.com/documentation/coreservices/1445940-aeputattributeptr
func AEPutAttributePtr(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize) int16 {
	if _aEPutAttributePtr == nil {
		panic("coreservices: symbol AEPutAttributePtr not loaded")
	}
	return _aEPutAttributePtr(theAppleEvent, theAEKeyword, typeCode, dataPtr, dataSize)
}

var _aEPutDesc func(theAEDescList unsafe.Pointer, index unsafe.Pointer, theAEDesc unsafe.Pointer) int16

// AEPutDesc adds a descriptor to any descriptor list, possibly replacing an existing descriptor in the list.
//
// See: https://developer.apple.com/documentation/coreservices/1450093-aeputdesc
func AEPutDesc(theAEDescList unsafe.Pointer, index unsafe.Pointer, theAEDesc unsafe.Pointer) int16 {
	if _aEPutDesc == nil {
		panic("coreservices: symbol AEPutDesc not loaded")
	}
	return _aEPutDesc(theAEDescList, index, theAEDesc)
}

var _aEPutParamDesc func(theAppleEvent *AEDesc, theAEKeyword uint32, theAEDesc unsafe.Pointer) int16

// AEPutParamDesc inserts a descriptor and a keyword into an Apple event or Apple event record as an Apple event parameter.
//
// See: https://developer.apple.com/documentation/coreservices/1447576-aeputparamdesc
func AEPutParamDesc(theAppleEvent *AEDesc, theAEKeyword uint32, theAEDesc unsafe.Pointer) int16 {
	if _aEPutParamDesc == nil {
		panic("coreservices: symbol AEPutParamDesc not loaded")
	}
	return _aEPutParamDesc(theAppleEvent, theAEKeyword, theAEDesc)
}

var _aEPutParamPtr func(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize) int16

// AEPutParamPtr inserts data, a descriptor type, and a keyword into an Apple event or Apple event record as an Apple event parameter.
//
// See: https://developer.apple.com/documentation/coreservices/1449263-aeputparamptr
func AEPutParamPtr(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize) int16 {
	if _aEPutParamPtr == nil {
		panic("coreservices: symbol AEPutParamPtr not loaded")
	}
	return _aEPutParamPtr(theAppleEvent, theAEKeyword, typeCode, dataPtr, dataSize)
}

var _aEPutPtr func(theAEDescList unsafe.Pointer, index unsafe.Pointer, typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize) int16

// AEPutPtr inserts data specified in a buffer into a descriptor list as a descriptor, possibly replacing an existing descriptor in the list.
//
// See: https://developer.apple.com/documentation/coreservices/1445287-aeputptr
func AEPutPtr(theAEDescList unsafe.Pointer, index unsafe.Pointer, typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize) int16 {
	if _aEPutPtr == nil {
		panic("coreservices: symbol AEPutPtr not loaded")
	}
	return _aEPutPtr(theAEDescList, index, typeCode, dataPtr, dataSize)
}

var _aERemoteProcessResolverGetProcesses func(ref AERemoteProcessResolverRef, outError unsafe.Pointer) corefoundation.CFArrayRef

// AERemoteProcessResolverGetProcesses returns an array of objects containing information about processes running on a remote machine.
//
// See: https://developer.apple.com/documentation/coreservices/1444456-aeremoteprocessresolvergetproces
func AERemoteProcessResolverGetProcesses(ref AERemoteProcessResolverRef, outError unsafe.Pointer) corefoundation.CFArrayRef {
	if _aERemoteProcessResolverGetProcesses == nil {
		panic("coreservices: symbol AERemoteProcessResolverGetProcesses not loaded")
	}
	return _aERemoteProcessResolverGetProcesses(ref, outError)
}

var _aERemoteProcessResolverScheduleWithRunLoop func(ref AERemoteProcessResolverRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef, callback AERemoteProcessResolverCallback, ctx unsafe.Pointer)

// AERemoteProcessResolverScheduleWithRunLoop schedules a resolver for execution on a given run loop in a given mode.
//
// See: https://developer.apple.com/documentation/coreservices/1447259-aeremoteprocessresolverschedulew
func AERemoteProcessResolverScheduleWithRunLoop(ref AERemoteProcessResolverRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef, callback AERemoteProcessResolverCallback, ctx unsafe.Pointer) {
	if _aERemoteProcessResolverScheduleWithRunLoop == nil {
		panic("coreservices: symbol AERemoteProcessResolverScheduleWithRunLoop not loaded")
	}
	_aERemoteProcessResolverScheduleWithRunLoop(ref, runLoop, runLoopMode, callback, ctx)
}

var _aERemoveCoercionHandler func(fromType uint32, toType uint32, handler AECoercionHandlerUPP, isSysHandler unsafe.Pointer) int16

// AERemoveCoercionHandler removes a coercion handler from a coercion handler dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1441907-aeremovecoercionhandler
func AERemoveCoercionHandler(fromType uint32, toType uint32, handler AECoercionHandlerUPP, isSysHandler unsafe.Pointer) int16 {
	if _aERemoveCoercionHandler == nil {
		panic("coreservices: symbol AERemoveCoercionHandler not loaded")
	}
	return _aERemoveCoercionHandler(fromType, toType, handler, isSysHandler)
}

var _aERemoveEventHandler func(theAEEventClass uint32, theAEEventID uint32, handler AEEventHandlerUPP, isSysHandler unsafe.Pointer) int16

// AERemoveEventHandler removes an event handler entry from an Apple event dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1445239-aeremoveeventhandler
func AERemoveEventHandler(theAEEventClass uint32, theAEEventID uint32, handler AEEventHandlerUPP, isSysHandler unsafe.Pointer) int16 {
	if _aERemoveEventHandler == nil {
		panic("coreservices: symbol AERemoveEventHandler not loaded")
	}
	return _aERemoveEventHandler(theAEEventClass, theAEEventID, handler, isSysHandler)
}

var _aERemoveObjectAccessor func(desiredClass uint32, containerType uint32, theAccessor OSLAccessorUPP, isSysHandler unsafe.Pointer) int16

// AERemoveObjectAccessor removes an object accessor function from an object accessor dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1442552-aeremoveobjectaccessor
func AERemoveObjectAccessor(desiredClass uint32, containerType uint32, theAccessor OSLAccessorUPP, isSysHandler unsafe.Pointer) int16 {
	if _aERemoveObjectAccessor == nil {
		panic("coreservices: symbol AERemoveObjectAccessor not loaded")
	}
	return _aERemoveObjectAccessor(desiredClass, containerType, theAccessor, isSysHandler)
}

var _aERemoveSpecialHandler func(functionClass uint32, handler AEEventHandlerUPP, isSysHandler unsafe.Pointer) int16

// AERemoveSpecialHandler removes a handler from a special handler dispatch table.
//
// See: https://developer.apple.com/documentation/coreservices/1447960-aeremovespecialhandler
func AERemoveSpecialHandler(functionClass uint32, handler AEEventHandlerUPP, isSysHandler unsafe.Pointer) int16 {
	if _aERemoveSpecialHandler == nil {
		panic("coreservices: symbol AERemoveSpecialHandler not loaded")
	}
	return _aERemoveSpecialHandler(functionClass, handler, isSysHandler)
}

var _aEReplaceDescData func(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, theAEDesc unsafe.Pointer) int16

// AEReplaceDescData copies the specified data into the specified descriptor, replacing any previous data.
//
// See: https://developer.apple.com/documentation/coreservices/1446695-aereplacedescdata
func AEReplaceDescData(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, theAEDesc unsafe.Pointer) int16 {
	if _aEReplaceDescData == nil {
		panic("coreservices: symbol AEReplaceDescData not loaded")
	}
	return _aEReplaceDescData(typeCode, dataPtr, dataSize, theAEDesc)
}

var _aEResolve func(objectSpecifier unsafe.Pointer, callbackFlags unsafe.Pointer, theToken unsafe.Pointer) int16

// AEResolve resolves an object specifier.
//
// See: https://developer.apple.com/documentation/coreservices/1449720-aeresolve
func AEResolve(objectSpecifier unsafe.Pointer, callbackFlags unsafe.Pointer, theToken unsafe.Pointer) int16 {
	if _aEResolve == nil {
		panic("coreservices: symbol AEResolve not loaded")
	}
	return _aEResolve(objectSpecifier, callbackFlags, theToken)
}

var _aESendMessage func(event *AEDesc, reply *AEDesc, sendMode AESendMode, timeOutInTicks unsafe.Pointer) int32

// AESendMessage sends an AppleEvent to a target process without some of the overhead required by [AESend].
//
// See: https://developer.apple.com/documentation/coreservices/1442994-aesendmessage
func AESendMessage(event *AEDesc, reply *AEDesc, sendMode AESendMode, timeOutInTicks unsafe.Pointer) int32 {
	if _aESendMessage == nil {
		panic("coreservices: symbol AESendMessage not loaded")
	}
	return _aESendMessage(event, reply, sendMode, timeOutInTicks)
}

var _aESetObjectCallbacks func(myCompareProc OSLCompareUPP, myCountProc OSLCountUPP, myDisposeTokenProc OSLDisposeTokenUPP, myGetMarkTokenProc OSLGetMarkTokenUPP, myMarkProc OSLMarkUPP, myAdjustMarksProc OSLAdjustMarksUPP, myGetErrDescProcPtr OSLGetErrDescUPP) int16

// AESetObjectCallbacks specifies the object callback functions for your application.
//
// See: https://developer.apple.com/documentation/coreservices/1447756-aesetobjectcallbacks
func AESetObjectCallbacks(myCompareProc OSLCompareUPP, myCountProc OSLCountUPP, myDisposeTokenProc OSLDisposeTokenUPP, myGetMarkTokenProc OSLGetMarkTokenUPP, myMarkProc OSLMarkUPP, myAdjustMarksProc OSLAdjustMarksUPP, myGetErrDescProcPtr OSLGetErrDescUPP) int16 {
	if _aESetObjectCallbacks == nil {
		panic("coreservices: symbol AESetObjectCallbacks not loaded")
	}
	return _aESetObjectCallbacks(myCompareProc, myCountProc, myDisposeTokenProc, myGetMarkTokenProc, myMarkProc, myAdjustMarksProc, myGetErrDescProcPtr)
}

var _aESizeOfAttribute func(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode *uint32, dataSize *corefoundation.CGSize) int16

// AESizeOfAttribute gets the size and descriptor type of an Apple event attribute from a descriptor of type [AppleEvent].
//
// See: https://developer.apple.com/documentation/coreservices/1445764-aesizeofattribute
func AESizeOfAttribute(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode *uint32, dataSize *corefoundation.CGSize) int16 {
	if _aESizeOfAttribute == nil {
		panic("coreservices: symbol AESizeOfAttribute not loaded")
	}
	return _aESizeOfAttribute(theAppleEvent, theAEKeyword, typeCode, dataSize)
}

var _aESizeOfFlattenedDesc func(theAEDesc unsafe.Pointer) corefoundation.CGSize

// AESizeOfFlattenedDesc returns the amount of buffer space needed to store the descriptor after flattening it.
//
// See: https://developer.apple.com/documentation/coreservices/1447305-aesizeofflatteneddesc
func AESizeOfFlattenedDesc(theAEDesc unsafe.Pointer) corefoundation.CGSize {
	if _aESizeOfFlattenedDesc == nil {
		panic("coreservices: symbol AESizeOfFlattenedDesc not loaded")
	}
	return _aESizeOfFlattenedDesc(theAEDesc)
}

var _aESizeOfNthItem func(theAEDescList unsafe.Pointer, index unsafe.Pointer, typeCode *uint32, dataSize *corefoundation.CGSize) int16

// AESizeOfNthItem gets the data size and descriptor type of the descriptor at a specified position in a descriptor list.
//
// See: https://developer.apple.com/documentation/coreservices/1447307-aesizeofnthitem
func AESizeOfNthItem(theAEDescList unsafe.Pointer, index unsafe.Pointer, typeCode *uint32, dataSize *corefoundation.CGSize) int16 {
	if _aESizeOfNthItem == nil {
		panic("coreservices: symbol AESizeOfNthItem not loaded")
	}
	return _aESizeOfNthItem(theAEDescList, index, typeCode, dataSize)
}

var _aESizeOfParam func(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode *uint32, dataSize *corefoundation.CGSize) int16

// AESizeOfParam gets the size and descriptor type of an Apple event parameter from a descriptor of type [AERecord] or [AppleEvent].
//
// See: https://developer.apple.com/documentation/coreservices/1449998-aesizeofparam
func AESizeOfParam(theAppleEvent *AEDesc, theAEKeyword uint32, typeCode *uint32, dataSize *corefoundation.CGSize) int16 {
	if _aESizeOfParam == nil {
		panic("coreservices: symbol AESizeOfParam not loaded")
	}
	return _aESizeOfParam(theAppleEvent, theAEKeyword, typeCode, dataSize)
}

var _aEStreamClose func(ref AEStreamRef, desc unsafe.Pointer) int32

// AEStreamClose closes and deallocates an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1449821-aestreamclose
func AEStreamClose(ref AEStreamRef, desc unsafe.Pointer) int32 {
	if _aEStreamClose == nil {
		panic("coreservices: symbol AEStreamClose not loaded")
	}
	return _aEStreamClose(ref, desc)
}

var _aEStreamCloseDesc func(ref AEStreamRef) int32

// AEStreamCloseDesc marks the end of a descriptor in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1449272-aestreamclosedesc
func AEStreamCloseDesc(ref AEStreamRef) int32 {
	if _aEStreamCloseDesc == nil {
		panic("coreservices: symbol AEStreamCloseDesc not loaded")
	}
	return _aEStreamCloseDesc(ref)
}

var _aEStreamCloseList func(ref AEStreamRef) int32

// AEStreamCloseList marks the end of a list of descriptors in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1448185-aestreamcloselist
func AEStreamCloseList(ref AEStreamRef) int32 {
	if _aEStreamCloseList == nil {
		panic("coreservices: symbol AEStreamCloseList not loaded")
	}
	return _aEStreamCloseList(ref)
}

var _aEStreamCloseRecord func(ref AEStreamRef) int32

// AEStreamCloseRecord marks the end of a record in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1449522-aestreamcloserecord
func AEStreamCloseRecord(ref AEStreamRef) int32 {
	if _aEStreamCloseRecord == nil {
		panic("coreservices: symbol AEStreamCloseRecord not loaded")
	}
	return _aEStreamCloseRecord(ref)
}

var _aEStreamCreateEvent func(clazz uint32, id uint32, targetType uint32, targetData unsafe.Pointer, targetLength corefoundation.CGSize, returnID unsafe.Pointer, transactionID unsafe.Pointer) AEStreamRef

// AEStreamCreateEvent creates a new Apple event and opens a stream for writing data to it.
//
// See: https://developer.apple.com/documentation/coreservices/1446562-aestreamcreateevent
func AEStreamCreateEvent(clazz uint32, id uint32, targetType uint32, targetData unsafe.Pointer, targetLength corefoundation.CGSize, returnID unsafe.Pointer, transactionID unsafe.Pointer) AEStreamRef {
	if _aEStreamCreateEvent == nil {
		panic("coreservices: symbol AEStreamCreateEvent not loaded")
	}
	return _aEStreamCreateEvent(clazz, id, targetType, targetData, targetLength, returnID, transactionID)
}

var _aEStreamOpen func() AEStreamRef

// AEStreamOpen opens a new [AEStreamRef] for use in building a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1447732-aestreamopen
func AEStreamOpen() AEStreamRef {
	if _aEStreamOpen == nil {
		panic("coreservices: symbol AEStreamOpen not loaded")
	}
	return _aEStreamOpen()
}

var _aEStreamOpenDesc func(ref AEStreamRef, newType uint32) int32

// AEStreamOpenDesc marks the beginning of a descriptor in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1446544-aestreamopendesc
func AEStreamOpenDesc(ref AEStreamRef, newType uint32) int32 {
	if _aEStreamOpenDesc == nil {
		panic("coreservices: symbol AEStreamOpenDesc not loaded")
	}
	return _aEStreamOpenDesc(ref, newType)
}

var _aEStreamOpenEvent func(event *AEDesc) AEStreamRef

// AEStreamOpenEvent opens a stream for an existing Apple event.
//
// See: https://developer.apple.com/documentation/coreservices/1445366-aestreamopenevent
func AEStreamOpenEvent(event *AEDesc) AEStreamRef {
	if _aEStreamOpenEvent == nil {
		panic("coreservices: symbol AEStreamOpenEvent not loaded")
	}
	return _aEStreamOpenEvent(event)
}

var _aEStreamOpenKeyDesc func(ref AEStreamRef, key uint32, newType uint32) int32

// AEStreamOpenKeyDesc marks the beginning of a key descriptor in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1442897-aestreamopenkeydesc
func AEStreamOpenKeyDesc(ref AEStreamRef, key uint32, newType uint32) int32 {
	if _aEStreamOpenKeyDesc == nil {
		panic("coreservices: symbol AEStreamOpenKeyDesc not loaded")
	}
	return _aEStreamOpenKeyDesc(ref, key, newType)
}

var _aEStreamOpenList func(ref AEStreamRef) int32

// AEStreamOpenList marks the beginning of a descriptor list in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1448594-aestreamopenlist
func AEStreamOpenList(ref AEStreamRef) int32 {
	if _aEStreamOpenList == nil {
		panic("coreservices: symbol AEStreamOpenList not loaded")
	}
	return _aEStreamOpenList(ref)
}

var _aEStreamOpenRecord func(ref AEStreamRef, newType uint32) int32

// AEStreamOpenRecord marks the beginning of an Apple event record in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1447141-aestreamopenrecord
func AEStreamOpenRecord(ref AEStreamRef, newType uint32) int32 {
	if _aEStreamOpenRecord == nil {
		panic("coreservices: symbol AEStreamOpenRecord not loaded")
	}
	return _aEStreamOpenRecord(ref, newType)
}

var _aEStreamOptionalParam func(ref AEStreamRef, key uint32) int32

// AEStreamOptionalParam designates a parameter in an Apple event as optional.
//
// See: https://developer.apple.com/documentation/coreservices/1444481-aestreamoptionalparam
func AEStreamOptionalParam(ref AEStreamRef, key uint32) int32 {
	if _aEStreamOptionalParam == nil {
		panic("coreservices: symbol AEStreamOptionalParam not loaded")
	}
	return _aEStreamOptionalParam(ref, key)
}

var _aEStreamSetRecordType func(ref AEStreamRef, newType uint32) int32

// AEStreamSetRecordType sets the type of the most recently created record in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1447704-aestreamsetrecordtype
func AEStreamSetRecordType(ref AEStreamRef, newType uint32) int32 {
	if _aEStreamSetRecordType == nil {
		panic("coreservices: symbol AEStreamSetRecordType not loaded")
	}
	return _aEStreamSetRecordType(ref, newType)
}

var _aEStreamWriteAEDesc func(ref AEStreamRef, desc unsafe.Pointer) int32

// AEStreamWriteAEDesc copies an existing descriptor into an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1448487-aestreamwriteaedesc
func AEStreamWriteAEDesc(ref AEStreamRef, desc unsafe.Pointer) int32 {
	if _aEStreamWriteAEDesc == nil {
		panic("coreservices: symbol AEStreamWriteAEDesc not loaded")
	}
	return _aEStreamWriteAEDesc(ref, desc)
}

var _aEStreamWriteData func(ref AEStreamRef, data unsafe.Pointer, length corefoundation.CGSize) int32

// AEStreamWriteData appends data to the current descriptor in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1442610-aestreamwritedata
func AEStreamWriteData(ref AEStreamRef, data unsafe.Pointer, length corefoundation.CGSize) int32 {
	if _aEStreamWriteData == nil {
		panic("coreservices: symbol AEStreamWriteData not loaded")
	}
	return _aEStreamWriteData(ref, data, length)
}

var _aEStreamWriteDesc func(ref AEStreamRef, newType uint32, data unsafe.Pointer, length corefoundation.CGSize) int32

// AEStreamWriteDesc appends the data for a complete descriptor to an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1450387-aestreamwritedesc
func AEStreamWriteDesc(ref AEStreamRef, newType uint32, data unsafe.Pointer, length corefoundation.CGSize) int32 {
	if _aEStreamWriteDesc == nil {
		panic("coreservices: symbol AEStreamWriteDesc not loaded")
	}
	return _aEStreamWriteDesc(ref, newType, data, length)
}

var _aEStreamWriteKey func(ref AEStreamRef, key uint32) int32

// AEStreamWriteKey marks the beginning of a keyword/descriptor pair for a descriptor in an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1448750-aestreamwritekey
func AEStreamWriteKey(ref AEStreamRef, key uint32) int32 {
	if _aEStreamWriteKey == nil {
		panic("coreservices: symbol AEStreamWriteKey not loaded")
	}
	return _aEStreamWriteKey(ref, key)
}

var _aEStreamWriteKeyDesc func(ref AEStreamRef, key uint32, newType uint32, data unsafe.Pointer, length corefoundation.CGSize) int32

// AEStreamWriteKeyDesc writes a complete keyword/descriptor pair to an [AEStreamRef].
//
// See: https://developer.apple.com/documentation/coreservices/1442568-aestreamwritekeydesc
func AEStreamWriteKeyDesc(ref AEStreamRef, key uint32, newType uint32, data unsafe.Pointer, length corefoundation.CGSize) int32 {
	if _aEStreamWriteKeyDesc == nil {
		panic("coreservices: symbol AEStreamWriteKeyDesc not loaded")
	}
	return _aEStreamWriteKeyDesc(ref, key, newType, data, length)
}

var _aEUnflattenDesc func(buffer unsafe.Pointer, result unsafe.Pointer) int32

// AEUnflattenDesc unflattens the data in the passed buffer and creates a descriptor from it.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/coreservices/1448997-aeunflattendesc
func AEUnflattenDesc(buffer unsafe.Pointer, result unsafe.Pointer) int32 {
	if _aEUnflattenDesc == nil {
		panic("coreservices: symbol AEUnflattenDesc not loaded")
	}
	return _aEUnflattenDesc(buffer, result)
}

var _aEUnflattenDescFromBytes func(buffer unsafe.Pointer, bufferLen unsafe.Pointer, result unsafe.Pointer) int32

// AEUnflattenDescFromBytes.
//
// See: https://developer.apple.com/documentation/coreservices/3553279-aeunflattendescfrombytes
func AEUnflattenDescFromBytes(buffer unsafe.Pointer, bufferLen unsafe.Pointer, result unsafe.Pointer) int32 {
	if _aEUnflattenDescFromBytes == nil {
		panic("coreservices: symbol AEUnflattenDescFromBytes not loaded")
	}
	return _aEUnflattenDescFromBytes(buffer, bufferLen, result)
}

var _absoluteDeltaToDuration func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// AbsoluteDeltaToDuration.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501245-absolutedeltatoduration
func AbsoluteDeltaToDuration(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _absoluteDeltaToDuration == nil {
		panic("coreservices: symbol AbsoluteDeltaToDuration not loaded")
	}
	return _absoluteDeltaToDuration(arg0, arg1)
}

var _absoluteDeltaToNanoseconds func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) Nanoseconds

// AbsoluteDeltaToNanoseconds.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501268-absolutedeltatonanoseconds
func AbsoluteDeltaToNanoseconds(arg0 unsafe.Pointer, arg1 unsafe.Pointer) Nanoseconds {
	if _absoluteDeltaToNanoseconds == nil {
		panic("coreservices: symbol AbsoluteDeltaToNanoseconds not loaded")
	}
	return _absoluteDeltaToNanoseconds(arg0, arg1)
}

var _absoluteToDuration func(arg0 unsafe.Pointer) unsafe.Pointer

// AbsoluteToDuration.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501266-absolutetoduration
func AbsoluteToDuration(arg0 unsafe.Pointer) unsafe.Pointer {
	if _absoluteToDuration == nil {
		panic("coreservices: symbol AbsoluteToDuration not loaded")
	}
	return _absoluteToDuration(arg0)
}

var _absoluteToNanoseconds func(arg0 unsafe.Pointer) Nanoseconds

// AbsoluteToNanoseconds.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501246-absolutetonanoseconds
func AbsoluteToNanoseconds(arg0 unsafe.Pointer) Nanoseconds {
	if _absoluteToNanoseconds == nil {
		panic("coreservices: symbol AbsoluteToNanoseconds not loaded")
	}
	return _absoluteToNanoseconds(arg0)
}

var _acquireIconRef func(theIconRef uintptr) int16

// AcquireIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1441852-acquireiconref
func AcquireIconRef(theIconRef uintptr) int16 {
	if _acquireIconRef == nil {
		panic("coreservices: symbol AcquireIconRef not loaded")
	}
	return _acquireIconRef(theIconRef)
}

var _addAbsoluteToAbsolute func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// AddAbsoluteToAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501267-addabsolutetoabsolute
func AddAbsoluteToAbsolute(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _addAbsoluteToAbsolute == nil {
		panic("coreservices: symbol AddAbsoluteToAbsolute not loaded")
	}
	return _addAbsoluteToAbsolute(arg0, arg1)
}

var _addAtomic func(arg0 int32, arg1 int32) int32

// AddAtomic.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490591-addatomic
func AddAtomic(arg0 int32, arg1 int32) int32 {
	if _addAtomic == nil {
		panic("coreservices: symbol AddAtomic not loaded")
	}
	return _addAtomic(arg0, arg1)
}

var _addAtomic16 func(arg0 int32, arg1 int16) int16

// AddAtomic16.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490589-addatomic16
func AddAtomic16(arg0 int32, arg1 int16) int16 {
	if _addAtomic16 == nil {
		panic("coreservices: symbol AddAtomic16 not loaded")
	}
	return _addAtomic16(arg0, arg1)
}

var _addAtomic8 func(arg0 int32, arg1 int8) int8

// AddAtomic8.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490594-addatomic8
func AddAtomic8(arg0 int32, arg1 int8) int8 {
	if _addAtomic8 == nil {
		panic("coreservices: symbol AddAtomic8 not loaded")
	}
	return _addAtomic8(arg0, arg1)
}

var _addCollectionItem func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32) int16

// AddCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551404-addcollectionitem
func AddCollectionItem(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32) int16 {
	if _addCollectionItem == nil {
		panic("coreservices: symbol AddCollectionItem not loaded")
	}
	return _addCollectionItem(arg0, arg1, arg2, arg3)
}

var _addCollectionItemHdl func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 uintptr) int16

// AddCollectionItemHdl.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551458-addcollectionitemhdl
func AddCollectionItemHdl(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 uintptr) int16 {
	if _addCollectionItemHdl == nil {
		panic("coreservices: symbol AddCollectionItemHdl not loaded")
	}
	return _addCollectionItemHdl(arg0, arg1, arg2, arg3)
}

var _addDurationToAbsolute func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// AddDurationToAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501249-adddurationtoabsolute
func AddDurationToAbsolute(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _addDurationToAbsolute == nil {
		panic("coreservices: symbol AddDurationToAbsolute not loaded")
	}
	return _addDurationToAbsolute(arg0, arg1)
}

var _addFolderDescriptor func(arg0 FolderType, arg1 FolderDescFlags, arg2 FolderClass, arg3 FolderLocation, arg4 uint32, arg5 uint32, arg6 unsafe.Pointer, arg7 bool) int16

// AddFolderDescriptor.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389301-addfolderdescriptor
func AddFolderDescriptor(arg0 FolderType, arg1 FolderDescFlags, arg2 FolderClass, arg3 FolderLocation, arg4 uint32, arg5 uint32, arg6 unsafe.Pointer, arg7 bool) int16 {
	if _addFolderDescriptor == nil {
		panic("coreservices: symbol AddFolderDescriptor not loaded")
	}
	return _addFolderDescriptor(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _addNanosecondsToAbsolute func(arg0 Nanoseconds, arg1 unsafe.Pointer) unsafe.Pointer

// AddNanosecondsToAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501264-addnanosecondstoabsolute
func AddNanosecondsToAbsolute(arg0 Nanoseconds, arg1 unsafe.Pointer) unsafe.Pointer {
	if _addNanosecondsToAbsolute == nil {
		panic("coreservices: symbol AddNanosecondsToAbsolute not loaded")
	}
	return _addNanosecondsToAbsolute(arg0, arg1)
}

var _addResource func(arg0 uintptr, arg1 uint32, arg2 ResID, arg3 unsafe.Pointer)

// AddResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529276-addresource
func AddResource(arg0 uintptr, arg1 uint32, arg2 ResID, arg3 unsafe.Pointer) {
	if _addResource == nil {
		panic("coreservices: symbol AddResource not loaded")
	}
	_addResource(arg0, arg1, arg2, arg3)
}

var _batteryCount func() int16

// BatteryCount.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427413-batterycount
func BatteryCount() int16 {
	if _batteryCount == nil {
		panic("coreservices: symbol BatteryCount not loaded")
	}
	return _batteryCount()
}

var _bitAnd func(arg0 int, arg1 int) int

// BitAnd.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494613-bitand
func BitAnd(arg0 int, arg1 int) int {
	if _bitAnd == nil {
		panic("coreservices: symbol BitAnd not loaded")
	}
	return _bitAnd(arg0, arg1)
}

var _bitAndAtomic func(arg0 uint32, arg1 uint32) uint32

// BitAndAtomic.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490596-bitandatomic
func BitAndAtomic(arg0 uint32, arg1 uint32) uint32 {
	if _bitAndAtomic == nil {
		panic("coreservices: symbol BitAndAtomic not loaded")
	}
	return _bitAndAtomic(arg0, arg1)
}

var _bitAndAtomic16 func(arg0 uint32, arg1 uint16) uint16

// BitAndAtomic16.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490599-bitandatomic16
func BitAndAtomic16(arg0 uint32, arg1 uint16) uint16 {
	if _bitAndAtomic16 == nil {
		panic("coreservices: symbol BitAndAtomic16 not loaded")
	}
	return _bitAndAtomic16(arg0, arg1)
}

var _bitAndAtomic8 func(arg0 uint32, arg1 uint8) uint8

// BitAndAtomic8.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490584-bitandatomic8
func BitAndAtomic8(arg0 uint32, arg1 uint8) uint8 {
	if _bitAndAtomic8 == nil {
		panic("coreservices: symbol BitAndAtomic8 not loaded")
	}
	return _bitAndAtomic8(arg0, arg1)
}

var _bitClr func(arg0 int)

// BitClr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494616-bitclr
func BitClr(arg0 int) {
	if _bitClr == nil {
		panic("coreservices: symbol BitClr not loaded")
	}
	_bitClr(arg0)
}

var _bitNot func(arg0 int) int

// BitNot.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494611-bitnot
func BitNot(arg0 int) int {
	if _bitNot == nil {
		panic("coreservices: symbol BitNot not loaded")
	}
	return _bitNot(arg0)
}

var _bitOr func(arg0 int, arg1 int) int

// BitOr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494619-bitor
func BitOr(arg0 int, arg1 int) int {
	if _bitOr == nil {
		panic("coreservices: symbol BitOr not loaded")
	}
	return _bitOr(arg0, arg1)
}

var _bitOrAtomic func(arg0 uint32, arg1 uint32) uint32

// BitOrAtomic.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490597-bitoratomic
func BitOrAtomic(arg0 uint32, arg1 uint32) uint32 {
	if _bitOrAtomic == nil {
		panic("coreservices: symbol BitOrAtomic not loaded")
	}
	return _bitOrAtomic(arg0, arg1)
}

var _bitOrAtomic16 func(arg0 uint32, arg1 uint16) uint16

// BitOrAtomic16.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490587-bitoratomic16
func BitOrAtomic16(arg0 uint32, arg1 uint16) uint16 {
	if _bitOrAtomic16 == nil {
		panic("coreservices: symbol BitOrAtomic16 not loaded")
	}
	return _bitOrAtomic16(arg0, arg1)
}

var _bitOrAtomic8 func(arg0 uint32, arg1 uint8) uint8

// BitOrAtomic8.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490581-bitoratomic8
func BitOrAtomic8(arg0 uint32, arg1 uint8) uint8 {
	if _bitOrAtomic8 == nil {
		panic("coreservices: symbol BitOrAtomic8 not loaded")
	}
	return _bitOrAtomic8(arg0, arg1)
}

var _bitSet func(arg0 int)

// BitSet.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494618-bitset
func BitSet(arg0 int) {
	if _bitSet == nil {
		panic("coreservices: symbol BitSet not loaded")
	}
	_bitSet(arg0)
}

var _bitShift func(arg0 int, arg1 int16) int

// BitShift.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494615-bitshift
func BitShift(arg0 int, arg1 int16) int {
	if _bitShift == nil {
		panic("coreservices: symbol BitShift not loaded")
	}
	return _bitShift(arg0, arg1)
}

var _bitTst func(arg0 int) bool

// BitTst.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494606-bittst
func BitTst(arg0 int) bool {
	if _bitTst == nil {
		panic("coreservices: symbol BitTst not loaded")
	}
	return _bitTst(arg0)
}

var _bitXor func(arg0 int, arg1 int) int

// BitXor.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1494620-bitxor
func BitXor(arg0 int, arg1 int) int {
	if _bitXor == nil {
		panic("coreservices: symbol BitXor not loaded")
	}
	return _bitXor(arg0, arg1)
}

var _bitXorAtomic func(arg0 uint32, arg1 uint32) uint32

// BitXorAtomic.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490592-bitxoratomic
func BitXorAtomic(arg0 uint32, arg1 uint32) uint32 {
	if _bitXorAtomic == nil {
		panic("coreservices: symbol BitXorAtomic not loaded")
	}
	return _bitXorAtomic(arg0, arg1)
}

var _bitXorAtomic16 func(arg0 uint32, arg1 uint16) uint16

// BitXorAtomic16.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490572-bitxoratomic16
func BitXorAtomic16(arg0 uint32, arg1 uint16) uint16 {
	if _bitXorAtomic16 == nil {
		panic("coreservices: symbol BitXorAtomic16 not loaded")
	}
	return _bitXorAtomic16(arg0, arg1)
}

var _bitXorAtomic8 func(arg0 uint32, arg1 uint8) uint8

// BitXorAtomic8.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490577-bitxoratomic8
func BitXorAtomic8(arg0 uint32, arg1 uint8) uint8 {
	if _bitXorAtomic8 == nil {
		panic("coreservices: symbol BitXorAtomic8 not loaded")
	}
	return _bitXorAtomic8(arg0, arg1)
}

var _cSBackupIsItemExcluded func(item corefoundation.CFURLRef, excludeByPath unsafe.Pointer) bool

// CSBackupIsItemExcluded returns a Boolean value indicating whether an item is currently excluded from the backup.
//
// See: https://developer.apple.com/documentation/coreservices/1443602-csbackupisitemexcluded
func CSBackupIsItemExcluded(item corefoundation.CFURLRef, excludeByPath unsafe.Pointer) bool {
	if _cSBackupIsItemExcluded == nil {
		panic("coreservices: symbol CSBackupIsItemExcluded not loaded")
	}
	return _cSBackupIsItemExcluded(item, excludeByPath)
}

var _cSBackupSetItemExcluded func(item corefoundation.CFURLRef, exclude unsafe.Pointer, excludeByPath unsafe.Pointer) int32

// CSBackupSetItemExcluded includes or excludes an item from the backup.
//
// See: https://developer.apple.com/documentation/coreservices/1445043-csbackupsetitemexcluded
func CSBackupSetItemExcluded(item corefoundation.CFURLRef, exclude unsafe.Pointer, excludeByPath unsafe.Pointer) int32 {
	if _cSBackupSetItemExcluded == nil {
		panic("coreservices: symbol CSBackupSetItemExcluded not loaded")
	}
	return _cSBackupSetItemExcluded(item, exclude, excludeByPath)
}

var _cSCopyMachineName func() corefoundation.CFStringRef

// CSCopyMachineName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533316-cscopymachinename
func CSCopyMachineName() corefoundation.CFStringRef {
	if _cSCopyMachineName == nil {
		panic("coreservices: symbol CSCopyMachineName not loaded")
	}
	return _cSCopyMachineName()
}

var _cSCopyUserName func(arg0 bool) corefoundation.CFStringRef

// CSCopyUserName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533406-cscopyusername
func CSCopyUserName(arg0 bool) corefoundation.CFStringRef {
	if _cSCopyUserName == nil {
		panic("coreservices: symbol CSCopyUserName not loaded")
	}
	return _cSCopyUserName(arg0)
}

var _cSDiskSpaceCancelRecovery func(operationUUID corefoundation.CFUUID)

// CSDiskSpaceCancelRecovery.
//
// See: https://developer.apple.com/documentation/coreservices/1448335-csdiskspacecancelrecovery
func CSDiskSpaceCancelRecovery(operationUUID corefoundation.CFUUID) {
	if _cSDiskSpaceCancelRecovery == nil {
		panic("coreservices: symbol CSDiskSpaceCancelRecovery not loaded")
	}
	_cSDiskSpaceCancelRecovery(operationUUID)
}

var _cSDiskSpaceGetRecoveryEstimate func(volumeURL corefoundation.CFURLRef) uint64

// CSDiskSpaceGetRecoveryEstimate.
//
// See: https://developer.apple.com/documentation/coreservices/1448439-csdiskspacegetrecoveryestimate
func CSDiskSpaceGetRecoveryEstimate(volumeURL corefoundation.CFURLRef) uint64 {
	if _cSDiskSpaceGetRecoveryEstimate == nil {
		panic("coreservices: symbol CSDiskSpaceGetRecoveryEstimate not loaded")
	}
	return _cSDiskSpaceGetRecoveryEstimate(volumeURL)
}

var _cSDiskSpaceStartRecovery func(volumeURL corefoundation.CFURLRef, bytesNeeded uint64, options uint32, outOperationUUID *corefoundation.CFUUID, callbackQueue uintptr, callback unsafe.Pointer)

// CSDiskSpaceStartRecovery.
//
// See: https://developer.apple.com/documentation/coreservices/1447968-csdiskspacestartrecovery
func CSDiskSpaceStartRecovery(volumeURL corefoundation.CFURLRef, bytesNeeded uint64, options uint32, outOperationUUID *corefoundation.CFUUID, callbackQueue dispatch.Queue, callback unsafe.Pointer) {
	if _cSDiskSpaceStartRecovery == nil {
		panic("coreservices: symbol CSDiskSpaceStartRecovery not loaded")
	}
	_cSDiskSpaceStartRecovery(volumeURL, bytesNeeded, options, outOperationUUID, uintptr(callbackQueue.Handle()), callback)
}

var _cSGetComponentsThreadMode func() CSComponentsThreadMode

// CSGetComponentsThreadMode indicates whether using thread-unsafe components is allowed in the current thread.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516409-csgetcomponentsthreadmode
func CSGetComponentsThreadMode() CSComponentsThreadMode {
	if _cSGetComponentsThreadMode == nil {
		panic("coreservices: symbol CSGetComponentsThreadMode not loaded")
	}
	return _cSGetComponentsThreadMode()
}

var _cSGetDefaultIdentityAuthority func() CSIdentityAuthorityRef

// CSGetDefaultIdentityAuthority.
//
// See: https://developer.apple.com/documentation/coreservices/1448826-csgetdefaultidentityauthority
func CSGetDefaultIdentityAuthority() CSIdentityAuthorityRef {
	if _cSGetDefaultIdentityAuthority == nil {
		panic("coreservices: symbol CSGetDefaultIdentityAuthority not loaded")
	}
	return _cSGetDefaultIdentityAuthority()
}

var _cSGetLocalIdentityAuthority func() CSIdentityAuthorityRef

// CSGetLocalIdentityAuthority.
//
// See: https://developer.apple.com/documentation/coreservices/1444814-csgetlocalidentityauthority
func CSGetLocalIdentityAuthority() CSIdentityAuthorityRef {
	if _cSGetLocalIdentityAuthority == nil {
		panic("coreservices: symbol CSGetLocalIdentityAuthority not loaded")
	}
	return _cSGetLocalIdentityAuthority()
}

var _cSGetManagedIdentityAuthority func() CSIdentityAuthorityRef

// CSGetManagedIdentityAuthority.
//
// See: https://developer.apple.com/documentation/coreservices/1446750-csgetmanagedidentityauthority
func CSGetManagedIdentityAuthority() CSIdentityAuthorityRef {
	if _cSGetManagedIdentityAuthority == nil {
		panic("coreservices: symbol CSGetManagedIdentityAuthority not loaded")
	}
	return _cSGetManagedIdentityAuthority()
}

var _cSIdentityAddAlias func(identity CSIdentityRef, alias corefoundation.CFStringRef)

// CSIdentityAddAlias.
//
// See: https://developer.apple.com/documentation/coreservices/1443062-csidentityaddalias
func CSIdentityAddAlias(identity CSIdentityRef, alias corefoundation.CFStringRef) {
	if _cSIdentityAddAlias == nil {
		panic("coreservices: symbol CSIdentityAddAlias not loaded")
	}
	_cSIdentityAddAlias(identity, alias)
}

var _cSIdentityAddMember func(group CSIdentityRef, member CSIdentityRef)

// CSIdentityAddMember.
//
// See: https://developer.apple.com/documentation/coreservices/1447513-csidentityaddmember
func CSIdentityAddMember(group CSIdentityRef, member CSIdentityRef) {
	if _cSIdentityAddMember == nil {
		panic("coreservices: symbol CSIdentityAddMember not loaded")
	}
	_cSIdentityAddMember(group, member)
}

var _cSIdentityAuthenticateUsingPassword func(user CSIdentityRef, password corefoundation.CFStringRef) bool

// CSIdentityAuthenticateUsingPassword.
//
// See: https://developer.apple.com/documentation/coreservices/1449855-csidentityauthenticateusingpassw
func CSIdentityAuthenticateUsingPassword(user CSIdentityRef, password corefoundation.CFStringRef) bool {
	if _cSIdentityAuthenticateUsingPassword == nil {
		panic("coreservices: symbol CSIdentityAuthenticateUsingPassword not loaded")
	}
	return _cSIdentityAuthenticateUsingPassword(user, password)
}

var _cSIdentityAuthorityCopyLocalizedName func(authority CSIdentityAuthorityRef) corefoundation.CFStringRef

// CSIdentityAuthorityCopyLocalizedName.
//
// See: https://developer.apple.com/documentation/coreservices/1448990-csidentityauthoritycopylocalized
func CSIdentityAuthorityCopyLocalizedName(authority CSIdentityAuthorityRef) corefoundation.CFStringRef {
	if _cSIdentityAuthorityCopyLocalizedName == nil {
		panic("coreservices: symbol CSIdentityAuthorityCopyLocalizedName not loaded")
	}
	return _cSIdentityAuthorityCopyLocalizedName(authority)
}

var _cSIdentityAuthorityGetTypeID func() uint

// CSIdentityAuthorityGetTypeID.
//
// See: https://developer.apple.com/documentation/coreservices/1449199-csidentityauthoritygettypeid
func CSIdentityAuthorityGetTypeID() uint {
	if _cSIdentityAuthorityGetTypeID == nil {
		panic("coreservices: symbol CSIdentityAuthorityGetTypeID not loaded")
	}
	return _cSIdentityAuthorityGetTypeID()
}

var _cSIdentityCommit func(identity CSIdentityRef, authorization security.AuthorizationRef, err *corefoundation.CFErrorRef) bool

// CSIdentityCommit.
//
// See: https://developer.apple.com/documentation/coreservices/1449575-csidentitycommit
func CSIdentityCommit(identity CSIdentityRef, authorization security.AuthorizationRef, err *corefoundation.CFErrorRef) bool {
	if _cSIdentityCommit == nil {
		panic("coreservices: symbol CSIdentityCommit not loaded")
	}
	return _cSIdentityCommit(identity, authorization, err)
}

var _cSIdentityCommitAsynchronously func(identity CSIdentityRef, clientContext unsafe.Pointer, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef, authorization security.AuthorizationRef) bool

// CSIdentityCommitAsynchronously.
//
// See: https://developer.apple.com/documentation/coreservices/1447936-csidentitycommitasynchronously
func CSIdentityCommitAsynchronously(identity CSIdentityRef, clientContext unsafe.Pointer, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef, authorization security.AuthorizationRef) bool {
	if _cSIdentityCommitAsynchronously == nil {
		panic("coreservices: symbol CSIdentityCommitAsynchronously not loaded")
	}
	return _cSIdentityCommitAsynchronously(identity, clientContext, runLoop, runLoopMode, authorization)
}

var _cSIdentityCreate func(allocator corefoundation.CFAllocatorRef, identityClass int32, fullName corefoundation.CFStringRef, posixName corefoundation.CFStringRef, flags uint32, authority CSIdentityAuthorityRef) CSIdentityRef

// CSIdentityCreate.
//
// See: https://developer.apple.com/documentation/coreservices/1447616-csidentitycreate
func CSIdentityCreate(allocator corefoundation.CFAllocatorRef, identityClass int32, fullName corefoundation.CFStringRef, posixName corefoundation.CFStringRef, flags uint32, authority CSIdentityAuthorityRef) CSIdentityRef {
	if _cSIdentityCreate == nil {
		panic("coreservices: symbol CSIdentityCreate not loaded")
	}
	return _cSIdentityCreate(allocator, identityClass, fullName, posixName, flags, authority)
}

var _cSIdentityCreateCopy func(allocator corefoundation.CFAllocatorRef, identity CSIdentityRef) CSIdentityRef

// CSIdentityCreateCopy.
//
// See: https://developer.apple.com/documentation/coreservices/1443553-csidentitycreatecopy
func CSIdentityCreateCopy(allocator corefoundation.CFAllocatorRef, identity CSIdentityRef) CSIdentityRef {
	if _cSIdentityCreateCopy == nil {
		panic("coreservices: symbol CSIdentityCreateCopy not loaded")
	}
	return _cSIdentityCreateCopy(allocator, identity)
}

var _cSIdentityCreateGroupMembershipQuery func(allocator corefoundation.CFAllocatorRef, group CSIdentityRef) CSIdentityQueryRef

// CSIdentityCreateGroupMembershipQuery.
//
// See: https://developer.apple.com/documentation/coreservices/1448605-csidentitycreategroupmembershipq
func CSIdentityCreateGroupMembershipQuery(allocator corefoundation.CFAllocatorRef, group CSIdentityRef) CSIdentityQueryRef {
	if _cSIdentityCreateGroupMembershipQuery == nil {
		panic("coreservices: symbol CSIdentityCreateGroupMembershipQuery not loaded")
	}
	return _cSIdentityCreateGroupMembershipQuery(allocator, group)
}

var _cSIdentityCreatePersistentReference func(allocator corefoundation.CFAllocatorRef, identity CSIdentityRef) corefoundation.CFDataRef

// CSIdentityCreatePersistentReference.
//
// See: https://developer.apple.com/documentation/coreservices/1444037-csidentitycreatepersistentrefere
func CSIdentityCreatePersistentReference(allocator corefoundation.CFAllocatorRef, identity CSIdentityRef) corefoundation.CFDataRef {
	if _cSIdentityCreatePersistentReference == nil {
		panic("coreservices: symbol CSIdentityCreatePersistentReference not loaded")
	}
	return _cSIdentityCreatePersistentReference(allocator, identity)
}

var _cSIdentityDelete func(identity CSIdentityRef)

// CSIdentityDelete.
//
// See: https://developer.apple.com/documentation/coreservices/1442616-csidentitydelete
func CSIdentityDelete(identity CSIdentityRef) {
	if _cSIdentityDelete == nil {
		panic("coreservices: symbol CSIdentityDelete not loaded")
	}
	_cSIdentityDelete(identity)
}

var _cSIdentityGetAliases func(identity CSIdentityRef) corefoundation.CFArrayRef

// CSIdentityGetAliases.
//
// See: https://developer.apple.com/documentation/coreservices/1446455-csidentitygetaliases
func CSIdentityGetAliases(identity CSIdentityRef) corefoundation.CFArrayRef {
	if _cSIdentityGetAliases == nil {
		panic("coreservices: symbol CSIdentityGetAliases not loaded")
	}
	return _cSIdentityGetAliases(identity)
}

var _cSIdentityGetAuthority func(identity CSIdentityRef) CSIdentityAuthorityRef

// CSIdentityGetAuthority.
//
// See: https://developer.apple.com/documentation/coreservices/1446828-csidentitygetauthority
func CSIdentityGetAuthority(identity CSIdentityRef) CSIdentityAuthorityRef {
	if _cSIdentityGetAuthority == nil {
		panic("coreservices: symbol CSIdentityGetAuthority not loaded")
	}
	return _cSIdentityGetAuthority(identity)
}

var _cSIdentityGetCertificate func(user CSIdentityRef) security.SecCertificateRef

// CSIdentityGetCertificate.
//
// See: https://developer.apple.com/documentation/coreservices/1448582-csidentitygetcertificate
func CSIdentityGetCertificate(user CSIdentityRef) security.SecCertificateRef {
	if _cSIdentityGetCertificate == nil {
		panic("coreservices: symbol CSIdentityGetCertificate not loaded")
	}
	return _cSIdentityGetCertificate(user)
}

var _cSIdentityGetClass func(identity CSIdentityRef) int32

// CSIdentityGetClass.
//
// See: https://developer.apple.com/documentation/coreservices/1447194-csidentitygetclass
func CSIdentityGetClass(identity CSIdentityRef) int32 {
	if _cSIdentityGetClass == nil {
		panic("coreservices: symbol CSIdentityGetClass not loaded")
	}
	return _cSIdentityGetClass(identity)
}

var _cSIdentityGetEmailAddress func(identity CSIdentityRef) corefoundation.CFStringRef

// CSIdentityGetEmailAddress.
//
// See: https://developer.apple.com/documentation/coreservices/1446211-csidentitygetemailaddress
func CSIdentityGetEmailAddress(identity CSIdentityRef) corefoundation.CFStringRef {
	if _cSIdentityGetEmailAddress == nil {
		panic("coreservices: symbol CSIdentityGetEmailAddress not loaded")
	}
	return _cSIdentityGetEmailAddress(identity)
}

var _cSIdentityGetFullName func(identity CSIdentityRef) corefoundation.CFStringRef

// CSIdentityGetFullName.
//
// See: https://developer.apple.com/documentation/coreservices/1447315-csidentitygetfullname
func CSIdentityGetFullName(identity CSIdentityRef) corefoundation.CFStringRef {
	if _cSIdentityGetFullName == nil {
		panic("coreservices: symbol CSIdentityGetFullName not loaded")
	}
	return _cSIdentityGetFullName(identity)
}

var _cSIdentityGetImageData func(identity CSIdentityRef) corefoundation.CFDataRef

// CSIdentityGetImageData.
//
// See: https://developer.apple.com/documentation/coreservices/1444544-csidentitygetimagedata
func CSIdentityGetImageData(identity CSIdentityRef) corefoundation.CFDataRef {
	if _cSIdentityGetImageData == nil {
		panic("coreservices: symbol CSIdentityGetImageData not loaded")
	}
	return _cSIdentityGetImageData(identity)
}

var _cSIdentityGetImageDataType func(identity CSIdentityRef) corefoundation.CFStringRef

// CSIdentityGetImageDataType.
//
// See: https://developer.apple.com/documentation/coreservices/1447478-csidentitygetimagedatatype
func CSIdentityGetImageDataType(identity CSIdentityRef) corefoundation.CFStringRef {
	if _cSIdentityGetImageDataType == nil {
		panic("coreservices: symbol CSIdentityGetImageDataType not loaded")
	}
	return _cSIdentityGetImageDataType(identity)
}

var _cSIdentityGetImageURL func(identity CSIdentityRef) corefoundation.CFURLRef

// CSIdentityGetImageURL.
//
// See: https://developer.apple.com/documentation/coreservices/1446099-csidentitygetimageurl
func CSIdentityGetImageURL(identity CSIdentityRef) corefoundation.CFURLRef {
	if _cSIdentityGetImageURL == nil {
		panic("coreservices: symbol CSIdentityGetImageURL not loaded")
	}
	return _cSIdentityGetImageURL(identity)
}

var _cSIdentityGetPosixID func(identity CSIdentityRef) unsafe.Pointer

// CSIdentityGetPosixID.
//
// See: https://developer.apple.com/documentation/coreservices/1443230-csidentitygetposixid
func CSIdentityGetPosixID(identity CSIdentityRef) unsafe.Pointer {
	if _cSIdentityGetPosixID == nil {
		panic("coreservices: symbol CSIdentityGetPosixID not loaded")
	}
	return _cSIdentityGetPosixID(identity)
}

var _cSIdentityGetPosixName func(identity CSIdentityRef) corefoundation.CFStringRef

// CSIdentityGetPosixName.
//
// See: https://developer.apple.com/documentation/coreservices/1447210-csidentitygetposixname
func CSIdentityGetPosixName(identity CSIdentityRef) corefoundation.CFStringRef {
	if _cSIdentityGetPosixName == nil {
		panic("coreservices: symbol CSIdentityGetPosixName not loaded")
	}
	return _cSIdentityGetPosixName(identity)
}

var _cSIdentityGetTypeID func() uint

// CSIdentityGetTypeID.
//
// See: https://developer.apple.com/documentation/coreservices/1444732-csidentitygettypeid
func CSIdentityGetTypeID() uint {
	if _cSIdentityGetTypeID == nil {
		panic("coreservices: symbol CSIdentityGetTypeID not loaded")
	}
	return _cSIdentityGetTypeID()
}

var _cSIdentityGetUUID func(identity CSIdentityRef) corefoundation.CFUUID

// CSIdentityGetUUID.
//
// See: https://developer.apple.com/documentation/coreservices/1447987-csidentitygetuuid
func CSIdentityGetUUID(identity CSIdentityRef) corefoundation.CFUUID {
	if _cSIdentityGetUUID == nil {
		panic("coreservices: symbol CSIdentityGetUUID not loaded")
	}
	return _cSIdentityGetUUID(identity)
}

var _cSIdentityIsCommitting func(identity CSIdentityRef) bool

// CSIdentityIsCommitting.
//
// See: https://developer.apple.com/documentation/coreservices/1449734-csidentityiscommitting
func CSIdentityIsCommitting(identity CSIdentityRef) bool {
	if _cSIdentityIsCommitting == nil {
		panic("coreservices: symbol CSIdentityIsCommitting not loaded")
	}
	return _cSIdentityIsCommitting(identity)
}

var _cSIdentityIsEnabled func(user CSIdentityRef) bool

// CSIdentityIsEnabled.
//
// See: https://developer.apple.com/documentation/coreservices/1443379-csidentityisenabled
func CSIdentityIsEnabled(user CSIdentityRef) bool {
	if _cSIdentityIsEnabled == nil {
		panic("coreservices: symbol CSIdentityIsEnabled not loaded")
	}
	return _cSIdentityIsEnabled(user)
}

var _cSIdentityIsHidden func(identity CSIdentityRef) bool

// CSIdentityIsHidden.
//
// See: https://developer.apple.com/documentation/coreservices/1449476-csidentityishidden
func CSIdentityIsHidden(identity CSIdentityRef) bool {
	if _cSIdentityIsHidden == nil {
		panic("coreservices: symbol CSIdentityIsHidden not loaded")
	}
	return _cSIdentityIsHidden(identity)
}

var _cSIdentityIsMemberOfGroup func(identity CSIdentityRef, group CSIdentityRef) bool

// CSIdentityIsMemberOfGroup.
//
// See: https://developer.apple.com/documentation/coreservices/1449237-csidentityismemberofgroup
func CSIdentityIsMemberOfGroup(identity CSIdentityRef, group CSIdentityRef) bool {
	if _cSIdentityIsMemberOfGroup == nil {
		panic("coreservices: symbol CSIdentityIsMemberOfGroup not loaded")
	}
	return _cSIdentityIsMemberOfGroup(identity, group)
}

var _cSIdentityQueryCopyResults func(query CSIdentityQueryRef) corefoundation.CFArrayRef

// CSIdentityQueryCopyResults.
//
// See: https://developer.apple.com/documentation/coreservices/1429035-csidentityquerycopyresults
func CSIdentityQueryCopyResults(query CSIdentityQueryRef) corefoundation.CFArrayRef {
	if _cSIdentityQueryCopyResults == nil {
		panic("coreservices: symbol CSIdentityQueryCopyResults not loaded")
	}
	return _cSIdentityQueryCopyResults(query)
}

var _cSIdentityQueryCreate func(allocator corefoundation.CFAllocatorRef, identityClass int32, authority CSIdentityAuthorityRef) CSIdentityQueryRef

// CSIdentityQueryCreate.
//
// See: https://developer.apple.com/documentation/coreservices/1429003-csidentityquerycreate
func CSIdentityQueryCreate(allocator corefoundation.CFAllocatorRef, identityClass int32, authority CSIdentityAuthorityRef) CSIdentityQueryRef {
	if _cSIdentityQueryCreate == nil {
		panic("coreservices: symbol CSIdentityQueryCreate not loaded")
	}
	return _cSIdentityQueryCreate(allocator, identityClass, authority)
}

var _cSIdentityQueryCreateForCurrentUser func(allocator corefoundation.CFAllocatorRef) CSIdentityQueryRef

// CSIdentityQueryCreateForCurrentUser.
//
// See: https://developer.apple.com/documentation/coreservices/1429037-csidentityquerycreateforcurrentu
func CSIdentityQueryCreateForCurrentUser(allocator corefoundation.CFAllocatorRef) CSIdentityQueryRef {
	if _cSIdentityQueryCreateForCurrentUser == nil {
		panic("coreservices: symbol CSIdentityQueryCreateForCurrentUser not loaded")
	}
	return _cSIdentityQueryCreateForCurrentUser(allocator)
}

var _cSIdentityQueryCreateForName func(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, comparisonMethod CSIdentityQueryStringComparisonMethod, identityClass int32, authority CSIdentityAuthorityRef) CSIdentityQueryRef

// CSIdentityQueryCreateForName.
//
// See: https://developer.apple.com/documentation/coreservices/1428997-csidentityquerycreateforname
func CSIdentityQueryCreateForName(allocator corefoundation.CFAllocatorRef, name corefoundation.CFStringRef, comparisonMethod CSIdentityQueryStringComparisonMethod, identityClass int32, authority CSIdentityAuthorityRef) CSIdentityQueryRef {
	if _cSIdentityQueryCreateForName == nil {
		panic("coreservices: symbol CSIdentityQueryCreateForName not loaded")
	}
	return _cSIdentityQueryCreateForName(allocator, name, comparisonMethod, identityClass, authority)
}

var _cSIdentityQueryCreateForPersistentReference func(allocator corefoundation.CFAllocatorRef, referenceData corefoundation.CFDataRef) CSIdentityQueryRef

// CSIdentityQueryCreateForPersistentReference.
//
// See: https://developer.apple.com/documentation/coreservices/1428991-csidentityquerycreateforpersiste
func CSIdentityQueryCreateForPersistentReference(allocator corefoundation.CFAllocatorRef, referenceData corefoundation.CFDataRef) CSIdentityQueryRef {
	if _cSIdentityQueryCreateForPersistentReference == nil {
		panic("coreservices: symbol CSIdentityQueryCreateForPersistentReference not loaded")
	}
	return _cSIdentityQueryCreateForPersistentReference(allocator, referenceData)
}

var _cSIdentityQueryCreateForPosixID func(allocator corefoundation.CFAllocatorRef, posixID unsafe.Pointer, identityClass int32, authority CSIdentityAuthorityRef) CSIdentityQueryRef

// CSIdentityQueryCreateForPosixID.
//
// See: https://developer.apple.com/documentation/coreservices/1428990-csidentityquerycreateforposixid
func CSIdentityQueryCreateForPosixID(allocator corefoundation.CFAllocatorRef, posixID unsafe.Pointer, identityClass int32, authority CSIdentityAuthorityRef) CSIdentityQueryRef {
	if _cSIdentityQueryCreateForPosixID == nil {
		panic("coreservices: symbol CSIdentityQueryCreateForPosixID not loaded")
	}
	return _cSIdentityQueryCreateForPosixID(allocator, posixID, identityClass, authority)
}

var _cSIdentityQueryCreateForUUID func(allocator corefoundation.CFAllocatorRef, uuid corefoundation.CFUUID, authority CSIdentityAuthorityRef) CSIdentityQueryRef

// CSIdentityQueryCreateForUUID.
//
// See: https://developer.apple.com/documentation/coreservices/1429007-csidentityquerycreateforuuid
func CSIdentityQueryCreateForUUID(allocator corefoundation.CFAllocatorRef, uuid corefoundation.CFUUID, authority CSIdentityAuthorityRef) CSIdentityQueryRef {
	if _cSIdentityQueryCreateForUUID == nil {
		panic("coreservices: symbol CSIdentityQueryCreateForUUID not loaded")
	}
	return _cSIdentityQueryCreateForUUID(allocator, uuid, authority)
}

var _cSIdentityQueryExecute func(query CSIdentityQueryRef, flags uint32, err *corefoundation.CFErrorRef) bool

// CSIdentityQueryExecute.
//
// See: https://developer.apple.com/documentation/coreservices/1429041-csidentityqueryexecute
func CSIdentityQueryExecute(query CSIdentityQueryRef, flags uint32, err *corefoundation.CFErrorRef) bool {
	if _cSIdentityQueryExecute == nil {
		panic("coreservices: symbol CSIdentityQueryExecute not loaded")
	}
	return _cSIdentityQueryExecute(query, flags, err)
}

var _cSIdentityQueryExecuteAsynchronously func(query CSIdentityQueryRef, flags uint32, clientContext unsafe.Pointer, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool

// CSIdentityQueryExecuteAsynchronously.
//
// See: https://developer.apple.com/documentation/coreservices/1429011-csidentityqueryexecuteasynchrono
func CSIdentityQueryExecuteAsynchronously(query CSIdentityQueryRef, flags uint32, clientContext unsafe.Pointer, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) bool {
	if _cSIdentityQueryExecuteAsynchronously == nil {
		panic("coreservices: symbol CSIdentityQueryExecuteAsynchronously not loaded")
	}
	return _cSIdentityQueryExecuteAsynchronously(query, flags, clientContext, runLoop, runLoopMode)
}

var _cSIdentityQueryGetTypeID func() uint

// CSIdentityQueryGetTypeID.
//
// See: https://developer.apple.com/documentation/coreservices/1429012-csidentityquerygettypeid
func CSIdentityQueryGetTypeID() uint {
	if _cSIdentityQueryGetTypeID == nil {
		panic("coreservices: symbol CSIdentityQueryGetTypeID not loaded")
	}
	return _cSIdentityQueryGetTypeID()
}

var _cSIdentityQueryStop func(query CSIdentityQueryRef)

// CSIdentityQueryStop.
//
// See: https://developer.apple.com/documentation/coreservices/1429047-csidentityquerystop
func CSIdentityQueryStop(query CSIdentityQueryRef) {
	if _cSIdentityQueryStop == nil {
		panic("coreservices: symbol CSIdentityQueryStop not loaded")
	}
	_cSIdentityQueryStop(query)
}

var _cSIdentityRemoveAlias func(identity CSIdentityRef, alias corefoundation.CFStringRef)

// CSIdentityRemoveAlias.
//
// See: https://developer.apple.com/documentation/coreservices/1442114-csidentityremovealias
func CSIdentityRemoveAlias(identity CSIdentityRef, alias corefoundation.CFStringRef) {
	if _cSIdentityRemoveAlias == nil {
		panic("coreservices: symbol CSIdentityRemoveAlias not loaded")
	}
	_cSIdentityRemoveAlias(identity, alias)
}

var _cSIdentityRemoveClient func(identity CSIdentityRef)

// CSIdentityRemoveClient.
//
// See: https://developer.apple.com/documentation/coreservices/1448933-csidentityremoveclient
func CSIdentityRemoveClient(identity CSIdentityRef) {
	if _cSIdentityRemoveClient == nil {
		panic("coreservices: symbol CSIdentityRemoveClient not loaded")
	}
	_cSIdentityRemoveClient(identity)
}

var _cSIdentityRemoveMember func(group CSIdentityRef, member CSIdentityRef)

// CSIdentityRemoveMember.
//
// See: https://developer.apple.com/documentation/coreservices/1448796-csidentityremovemember
func CSIdentityRemoveMember(group CSIdentityRef, member CSIdentityRef) {
	if _cSIdentityRemoveMember == nil {
		panic("coreservices: symbol CSIdentityRemoveMember not loaded")
	}
	_cSIdentityRemoveMember(group, member)
}

var _cSIdentitySetCertificate func(user CSIdentityRef, certificate security.SecCertificateRef)

// CSIdentitySetCertificate.
//
// See: https://developer.apple.com/documentation/coreservices/1447691-csidentitysetcertificate
func CSIdentitySetCertificate(user CSIdentityRef, certificate security.SecCertificateRef) {
	if _cSIdentitySetCertificate == nil {
		panic("coreservices: symbol CSIdentitySetCertificate not loaded")
	}
	_cSIdentitySetCertificate(user, certificate)
}

var _cSIdentitySetEmailAddress func(identity CSIdentityRef, emailAddress corefoundation.CFStringRef)

// CSIdentitySetEmailAddress.
//
// See: https://developer.apple.com/documentation/coreservices/1443235-csidentitysetemailaddress
func CSIdentitySetEmailAddress(identity CSIdentityRef, emailAddress corefoundation.CFStringRef) {
	if _cSIdentitySetEmailAddress == nil {
		panic("coreservices: symbol CSIdentitySetEmailAddress not loaded")
	}
	_cSIdentitySetEmailAddress(identity, emailAddress)
}

var _cSIdentitySetFullName func(identity CSIdentityRef, fullName corefoundation.CFStringRef)

// CSIdentitySetFullName.
//
// See: https://developer.apple.com/documentation/coreservices/1446623-csidentitysetfullname
func CSIdentitySetFullName(identity CSIdentityRef, fullName corefoundation.CFStringRef) {
	if _cSIdentitySetFullName == nil {
		panic("coreservices: symbol CSIdentitySetFullName not loaded")
	}
	_cSIdentitySetFullName(identity, fullName)
}

var _cSIdentitySetImageData func(identity CSIdentityRef, imageData corefoundation.CFDataRef, imageDataType corefoundation.CFStringRef)

// CSIdentitySetImageData.
//
// See: https://developer.apple.com/documentation/coreservices/1441866-csidentitysetimagedata
func CSIdentitySetImageData(identity CSIdentityRef, imageData corefoundation.CFDataRef, imageDataType corefoundation.CFStringRef) {
	if _cSIdentitySetImageData == nil {
		panic("coreservices: symbol CSIdentitySetImageData not loaded")
	}
	_cSIdentitySetImageData(identity, imageData, imageDataType)
}

var _cSIdentitySetImageURL func(identity CSIdentityRef, url corefoundation.CFURLRef)

// CSIdentitySetImageURL.
//
// See: https://developer.apple.com/documentation/coreservices/1446717-csidentitysetimageurl
func CSIdentitySetImageURL(identity CSIdentityRef, url corefoundation.CFURLRef) {
	if _cSIdentitySetImageURL == nil {
		panic("coreservices: symbol CSIdentitySetImageURL not loaded")
	}
	_cSIdentitySetImageURL(identity, url)
}

var _cSIdentitySetIsEnabled func(user CSIdentityRef, isEnabled unsafe.Pointer)

// CSIdentitySetIsEnabled.
//
// See: https://developer.apple.com/documentation/coreservices/1443028-csidentitysetisenabled
func CSIdentitySetIsEnabled(user CSIdentityRef, isEnabled unsafe.Pointer) {
	if _cSIdentitySetIsEnabled == nil {
		panic("coreservices: symbol CSIdentitySetIsEnabled not loaded")
	}
	_cSIdentitySetIsEnabled(user, isEnabled)
}

var _cSIdentitySetPassword func(user CSIdentityRef, password corefoundation.CFStringRef)

// CSIdentitySetPassword.
//
// See: https://developer.apple.com/documentation/coreservices/1443568-csidentitysetpassword
func CSIdentitySetPassword(user CSIdentityRef, password corefoundation.CFStringRef) {
	if _cSIdentitySetPassword == nil {
		panic("coreservices: symbol CSIdentitySetPassword not loaded")
	}
	_cSIdentitySetPassword(user, password)
}

var _cSSetComponentsThreadMode func(arg0 CSComponentsThreadMode)

// CSSetComponentsThreadMode sets whether or not using thread-unsafe components is allowed in the current thread.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516422-cssetcomponentsthreadmode
func CSSetComponentsThreadMode(arg0 CSComponentsThreadMode) {
	if _cSSetComponentsThreadMode == nil {
		panic("coreservices: symbol CSSetComponentsThreadMode not loaded")
	}
	_cSSetComponentsThreadMode(arg0)
}

var _callComponentCanDo func(arg0 ComponentInstance, arg1 int16) ComponentResult

// CallComponentCanDo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516321-callcomponentcando
func CallComponentCanDo(arg0 ComponentInstance, arg1 int16) ComponentResult {
	if _callComponentCanDo == nil {
		panic("coreservices: symbol CallComponentCanDo not loaded")
	}
	return _callComponentCanDo(arg0, arg1)
}

var _callComponentClose func(arg0 ComponentInstance, arg1 ComponentInstance) ComponentResult

// CallComponentClose.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516604-callcomponentclose
func CallComponentClose(arg0 ComponentInstance, arg1 ComponentInstance) ComponentResult {
	if _callComponentClose == nil {
		panic("coreservices: symbol CallComponentClose not loaded")
	}
	return _callComponentClose(arg0, arg1)
}

var _callComponentDispatch func(arg0 ComponentParameters) ComponentResult

// CallComponentDispatch.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516547-callcomponentdispatch
func CallComponentDispatch(arg0 ComponentParameters) ComponentResult {
	if _callComponentDispatch == nil {
		panic("coreservices: symbol CallComponentDispatch not loaded")
	}
	return _callComponentDispatch(arg0)
}

var _callComponentFunction func(arg0 ComponentParameters, arg1 ComponentFunctionUPP) ComponentResult

// CallComponentFunction invokes the specified function of your component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516603-callcomponentfunction
func CallComponentFunction(arg0 ComponentParameters, arg1 ComponentFunctionUPP) ComponentResult {
	if _callComponentFunction == nil {
		panic("coreservices: symbol CallComponentFunction not loaded")
	}
	return _callComponentFunction(arg0, arg1)
}

var _callComponentFunctionWithStorage func(arg0 uintptr, arg1 ComponentParameters, arg2 ComponentFunctionUPP) ComponentResult

// CallComponentFunctionWithStorage invokes the specified function of your component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516610-callcomponentfunctionwithstorage
func CallComponentFunctionWithStorage(arg0 uintptr, arg1 ComponentParameters, arg2 ComponentFunctionUPP) ComponentResult {
	if _callComponentFunctionWithStorage == nil {
		panic("coreservices: symbol CallComponentFunctionWithStorage not loaded")
	}
	return _callComponentFunctionWithStorage(arg0, arg1, arg2)
}

var _callComponentFunctionWithStorageProcInfo func(arg0 uintptr, arg1 ComponentParameters, arg2 unsafe.Pointer, arg3 ProcInfoType) ComponentResult

// CallComponentFunctionWithStorageProcInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516344-callcomponentfunctionwithstorage
func CallComponentFunctionWithStorageProcInfo(arg0 uintptr, arg1 ComponentParameters, arg2 unsafe.Pointer, arg3 ProcInfoType) ComponentResult {
	if _callComponentFunctionWithStorageProcInfo == nil {
		panic("coreservices: symbol CallComponentFunctionWithStorageProcInfo not loaded")
	}
	return _callComponentFunctionWithStorageProcInfo(arg0, arg1, arg2, arg3)
}

var _callComponentGetMPWorkFunction func(arg0 ComponentInstance, arg1 ComponentMPWorkFunctionUPP) ComponentResult

// CallComponentGetMPWorkFunction.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516346-callcomponentgetmpworkfunction
func CallComponentGetMPWorkFunction(arg0 ComponentInstance, arg1 ComponentMPWorkFunctionUPP) ComponentResult {
	if _callComponentGetMPWorkFunction == nil {
		panic("coreservices: symbol CallComponentGetMPWorkFunction not loaded")
	}
	return _callComponentGetMPWorkFunction(arg0, arg1)
}

var _callComponentGetPublicResource func(arg0 ComponentInstance, arg1 uint32, arg2 int16, arg3 uintptr) ComponentResult

// CallComponentGetPublicResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516354-callcomponentgetpublicresource
func CallComponentGetPublicResource(arg0 ComponentInstance, arg1 uint32, arg2 int16, arg3 uintptr) ComponentResult {
	if _callComponentGetPublicResource == nil {
		panic("coreservices: symbol CallComponentGetPublicResource not loaded")
	}
	return _callComponentGetPublicResource(arg0, arg1, arg2, arg3)
}

var _callComponentOpen func(arg0 ComponentInstance, arg1 ComponentInstance) ComponentResult

// CallComponentOpen.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516529-callcomponentopen
func CallComponentOpen(arg0 ComponentInstance, arg1 ComponentInstance) ComponentResult {
	if _callComponentOpen == nil {
		panic("coreservices: symbol CallComponentOpen not loaded")
	}
	return _callComponentOpen(arg0, arg1)
}

var _callComponentRegister func(arg0 ComponentInstance) ComponentResult

// CallComponentRegister.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516385-callcomponentregister
func CallComponentRegister(arg0 ComponentInstance) ComponentResult {
	if _callComponentRegister == nil {
		panic("coreservices: symbol CallComponentRegister not loaded")
	}
	return _callComponentRegister(arg0)
}

var _callComponentTarget func(arg0 ComponentInstance, arg1 ComponentInstance) ComponentResult

// CallComponentTarget.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516493-callcomponenttarget
func CallComponentTarget(arg0 ComponentInstance, arg1 ComponentInstance) ComponentResult {
	if _callComponentTarget == nil {
		panic("coreservices: symbol CallComponentTarget not loaded")
	}
	return _callComponentTarget(arg0, arg1)
}

var _callComponentUnregister func(arg0 ComponentInstance) ComponentResult

// CallComponentUnregister.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516337-callcomponentunregister
func CallComponentUnregister(arg0 ComponentInstance) ComponentResult {
	if _callComponentUnregister == nil {
		panic("coreservices: symbol CallComponentUnregister not loaded")
	}
	return _callComponentUnregister(arg0)
}

var _callComponentVersion func(arg0 ComponentInstance) ComponentResult

// CallComponentVersion.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516489-callcomponentversion
func CallComponentVersion(arg0 ComponentInstance) ComponentResult {
	if _callComponentVersion == nil {
		panic("coreservices: symbol CallComponentVersion not loaded")
	}
	return _callComponentVersion(arg0)
}

var _captureComponent func(arg0 Component, arg1 Component) Component

// CaptureComponent allows your component to capture another component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516357-capturecomponent
func CaptureComponent(arg0 Component, arg1 Component) Component {
	if _captureComponent == nil {
		panic("coreservices: symbol CaptureComponent not loaded")
	}
	return _captureComponent(arg0, arg1)
}

var _changeTextToUnicodeInfo func(arg0 TextToUnicodeInfo, arg1 ConstUnicodeMappingPtr) int32

// ChangeTextToUnicodeInfo changes the mapping information for the specified Unicodeconverter object used to convert text to Unicode to the new mappingyou provide.
//
// See: https://developer.apple.com/documentation/coreservices/1433487-changetexttounicodeinfo
func ChangeTextToUnicodeInfo(arg0 TextToUnicodeInfo, arg1 ConstUnicodeMappingPtr) int32 {
	if _changeTextToUnicodeInfo == nil {
		panic("coreservices: symbol ChangeTextToUnicodeInfo not loaded")
	}
	return _changeTextToUnicodeInfo(arg0, arg1)
}

var _changeUnicodeToTextInfo func(arg0 UnicodeToTextInfo, arg1 ConstUnicodeMappingPtr) int32

// ChangeUnicodeToTextInfo changes the mapping information contained in the specifiedUnicode converter object used to convert Unicode text to a non-Unicodeencoding.
//
// See: https://developer.apple.com/documentation/coreservices/1433509-changeunicodetotextinfo
func ChangeUnicodeToTextInfo(arg0 UnicodeToTextInfo, arg1 ConstUnicodeMappingPtr) int32 {
	if _changeUnicodeToTextInfo == nil {
		panic("coreservices: symbol ChangeUnicodeToTextInfo not loaded")
	}
	return _changeUnicodeToTextInfo(arg0, arg1)
}

var _changedResource func(arg0 uintptr)

// ChangedResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529318-changedresource
func ChangedResource(arg0 uintptr) {
	if _changedResource == nil {
		panic("coreservices: symbol ChangedResource not loaded")
	}
	_changedResource(arg0)
}

var _cloneCollection func(arg0 Collection) Collection

// CloneCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551333-clonecollection
func CloneCollection(arg0 Collection) Collection {
	if _cloneCollection == nil {
		panic("coreservices: symbol CloneCollection not loaded")
	}
	return _cloneCollection(arg0)
}

var _closeComponent func(arg0 ComponentInstance) int16

// CloseComponent terminates your application’s connection to a component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516436-closecomponent
func CloseComponent(arg0 ComponentInstance) int16 {
	if _closeComponent == nil {
		panic("coreservices: symbol CloseComponent not loaded")
	}
	return _closeComponent(arg0)
}

var _closeComponentResFile func(arg0 ResFileRefNum) int16

// CloseComponentResFile closes the resource file that your component opened previously with the [OpenComponentResFile] function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516319-closecomponentresfile
func CloseComponentResFile(arg0 ResFileRefNum) int16 {
	if _closeComponentResFile == nil {
		panic("coreservices: symbol CloseComponentResFile not loaded")
	}
	return _closeComponentResFile(arg0)
}

var _closeResFile func(arg0 ResFileRefNum)

// CloseResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529362-closeresfile
func CloseResFile(arg0 ResFileRefNum) {
	if _closeResFile == nil {
		panic("coreservices: symbol CloseResFile not loaded")
	}
	_closeResFile(arg0)
}

var _collectionTagExists func(arg0 Collection, arg1 CollectionTag) bool

// CollectionTagExists.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551428-collectiontagexists
func CollectionTagExists(arg0 Collection, arg1 CollectionTag) bool {
	if _collectionTagExists == nil {
		panic("coreservices: symbol CollectionTagExists not loaded")
	}
	return _collectionTagExists(arg0, arg1)
}

var _compareAndSwap func(arg0 uint32, arg1 uint32, arg2 uint32) bool

// CompareAndSwap.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490575-compareandswap
func CompareAndSwap(arg0 uint32, arg1 uint32, arg2 uint32) bool {
	if _compareAndSwap == nil {
		panic("coreservices: symbol CompareAndSwap not loaded")
	}
	return _compareAndSwap(arg0, arg1, arg2)
}

var _compositeIconRef func(backgroundIconRef uintptr, foregroundIconRef uintptr, compositeIconRef uintptr) int16

// CompositeIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1450541-compositeiconref
func CompositeIconRef(backgroundIconRef uintptr, foregroundIconRef uintptr, compositeIconRef uintptr) int16 {
	if _compositeIconRef == nil {
		panic("coreservices: symbol CompositeIconRef not loaded")
	}
	return _compositeIconRef(backgroundIconRef, foregroundIconRef, compositeIconRef)
}

var _convertFromPStringToUnicode func(arg0 TextToUnicodeInfo, arg1 unsafe.Pointer, arg2 uintptr, arg3 uintptr, arg4 uint16) int32

// ConvertFromPStringToUnicode converts a Pascal string in a Mac OS text encoding toa Unicode string.
//
// See: https://developer.apple.com/documentation/coreservices/1433483-convertfrompstringtounicode
func ConvertFromPStringToUnicode(arg0 TextToUnicodeInfo, arg1 unsafe.Pointer, arg2 uintptr, arg3 uintptr, arg4 uint16) int32 {
	if _convertFromPStringToUnicode == nil {
		panic("coreservices: symbol ConvertFromPStringToUnicode not loaded")
	}
	return _convertFromPStringToUnicode(arg0, arg1, arg2, arg3, arg4)
}

var _convertFromTextToUnicode func(arg0 TextToUnicodeInfo, arg1 uintptr, arg2 unsafe.Pointer, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 uint16) int32

// ConvertFromTextToUnicode converts a string from any encoding to Unicode.
//
// See: https://developer.apple.com/documentation/coreservices/1433517-convertfromtexttounicode
func ConvertFromTextToUnicode(arg0 TextToUnicodeInfo, arg1 uintptr, arg2 unsafe.Pointer, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 uint16) int32 {
	if _convertFromTextToUnicode == nil {
		panic("coreservices: symbol ConvertFromTextToUnicode not loaded")
	}
	return _convertFromTextToUnicode(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
}

var _convertFromUnicodeToPString func(arg0 UnicodeToTextInfo, arg1 uintptr, arg2 uint16, arg3 unsafe.Pointer) int32

// ConvertFromUnicodeToPString converts a Unicode string to Pascal in a Mac OS text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1433581-convertfromunicodetopstring
func ConvertFromUnicodeToPString(arg0 UnicodeToTextInfo, arg1 uintptr, arg2 uint16, arg3 unsafe.Pointer) int32 {
	if _convertFromUnicodeToPString == nil {
		panic("coreservices: symbol ConvertFromUnicodeToPString not loaded")
	}
	return _convertFromUnicodeToPString(arg0, arg1, arg2, arg3)
}

var _convertFromUnicodeToScriptCodeRun func(arg0 UnicodeToTextRunInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 unsafe.Pointer, arg12 uintptr, arg13 uintptr, arg14 ScriptCodeRun) int32

// ConvertFromUnicodeToScriptCodeRun converts a string from Unicode to one or more scripts.
//
// See: https://developer.apple.com/documentation/coreservices/1433662-convertfromunicodetoscriptcoderu
func ConvertFromUnicodeToScriptCodeRun(arg0 UnicodeToTextRunInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 unsafe.Pointer, arg12 uintptr, arg13 uintptr, arg14 ScriptCodeRun) int32 {
	if _convertFromUnicodeToScriptCodeRun == nil {
		panic("coreservices: symbol ConvertFromUnicodeToScriptCodeRun not loaded")
	}
	return _convertFromUnicodeToScriptCodeRun(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14)
}

var _convertFromUnicodeToText func(arg0 UnicodeToTextInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 unsafe.Pointer) int32

// ConvertFromUnicodeToText converts a Unicode text string to the destination encodingyou specify.
//
// See: https://developer.apple.com/documentation/coreservices/1433542-convertfromunicodetotext
func ConvertFromUnicodeToText(arg0 UnicodeToTextInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 unsafe.Pointer) int32 {
	if _convertFromUnicodeToText == nil {
		panic("coreservices: symbol ConvertFromUnicodeToText not loaded")
	}
	return _convertFromUnicodeToText(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
}

var _convertFromUnicodeToTextRun func(arg0 UnicodeToTextRunInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 unsafe.Pointer, arg12 uintptr, arg13 uintptr, arg14 TextEncodingRun) int32

// ConvertFromUnicodeToTextRun converts a string from Unicode to one or more encodings.
//
// See: https://developer.apple.com/documentation/coreservices/1433511-convertfromunicodetotextrun
func ConvertFromUnicodeToTextRun(arg0 UnicodeToTextRunInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr, arg9 uintptr, arg10 uintptr, arg11 unsafe.Pointer, arg12 uintptr, arg13 uintptr, arg14 TextEncodingRun) int32 {
	if _convertFromUnicodeToTextRun == nil {
		panic("coreservices: symbol ConvertFromUnicodeToTextRun not loaded")
	}
	return _convertFromUnicodeToTextRun(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14)
}

var _copyCollection func(arg0 Collection, arg1 Collection) Collection

// CopyCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551383-copycollection
func CopyCollection(arg0 Collection, arg1 Collection) Collection {
	if _copyCollection == nil {
		panic("coreservices: symbol CopyCollection not loaded")
	}
	return _copyCollection(arg0, arg1)
}

var _coreEndianFlipData func(arg0 uint32, arg1 uint32, arg2 int16, arg3 uintptr, arg4 bool) int32

// CoreEndianFlipData calls the flipper callback associated with the specifieddata type.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1575610-coreendianflipdata
func CoreEndianFlipData(arg0 uint32, arg1 uint32, arg2 int16, arg3 uintptr, arg4 bool) int32 {
	if _coreEndianFlipData == nil {
		panic("coreservices: symbol CoreEndianFlipData not loaded")
	}
	return _coreEndianFlipData(arg0, arg1, arg2, arg3, arg4)
}

var _coreEndianGetFlipper func(arg0 uint32, arg1 uint32, arg2 unsafe.Pointer) int32

// CoreEndianGetFlipper obtains the flipper callback that is installed for thespecified data type.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1575564-coreendiangetflipper
func CoreEndianGetFlipper(arg0 uint32, arg1 uint32, arg2 unsafe.Pointer) int32 {
	if _coreEndianGetFlipper == nil {
		panic("coreservices: symbol CoreEndianGetFlipper not loaded")
	}
	return _coreEndianGetFlipper(arg0, arg1, arg2)
}

var _coreEndianInstallFlipper func(arg0 uint32, arg1 uint32, arg2 unsafe.Pointer) int32

// CoreEndianInstallFlipper installs a flipper callback for the specified data type.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1575602-coreendianinstallflipper
func CoreEndianInstallFlipper(arg0 uint32, arg1 uint32, arg2 unsafe.Pointer) int32 {
	if _coreEndianInstallFlipper == nil {
		panic("coreservices: symbol CoreEndianInstallFlipper not loaded")
	}
	return _coreEndianInstallFlipper(arg0, arg1, arg2)
}

var _count1Resources func(arg0 uint32) ResourceCount

// Count1Resources.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529312-count1resources
func Count1Resources(arg0 uint32) ResourceCount {
	if _count1Resources == nil {
		panic("coreservices: symbol Count1Resources not loaded")
	}
	return _count1Resources(arg0)
}

var _count1Types func() ResourceCount

// Count1Types.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529244-count1types
func Count1Types() ResourceCount {
	if _count1Types == nil {
		panic("coreservices: symbol Count1Types not loaded")
	}
	return _count1Types()
}

var _countCollectionItems func(arg0 Collection) int32

// CountCollectionItems.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551438-countcollectionitems
func CountCollectionItems(arg0 Collection) int32 {
	if _countCollectionItems == nil {
		panic("coreservices: symbol CountCollectionItems not loaded")
	}
	return _countCollectionItems(arg0)
}

var _countCollectionOwners func(arg0 Collection) int32

// CountCollectionOwners.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551353-countcollectionowners
func CountCollectionOwners(arg0 Collection) int32 {
	if _countCollectionOwners == nil {
		panic("coreservices: symbol CountCollectionOwners not loaded")
	}
	return _countCollectionOwners(arg0)
}

var _countCollectionTags func(arg0 Collection) int32

// CountCollectionTags.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551378-countcollectiontags
func CountCollectionTags(arg0 Collection) int32 {
	if _countCollectionTags == nil {
		panic("coreservices: symbol CountCollectionTags not loaded")
	}
	return _countCollectionTags(arg0)
}

var _countComponentInstances func(arg0 Component) int

// CountComponentInstances determines the number of open connections being managed by a specified component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516394-countcomponentinstances
func CountComponentInstances(arg0 Component) int {
	if _countComponentInstances == nil {
		panic("coreservices: symbol CountComponentInstances not loaded")
	}
	return _countComponentInstances(arg0)
}

var _countComponents func(arg0 ComponentDescription) int

// CountComponents returns the number of registered components that meet the selection criteria specified by your application.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516515-countcomponents
func CountComponents(arg0 ComponentDescription) int {
	if _countComponents == nil {
		panic("coreservices: symbol CountComponents not loaded")
	}
	return _countComponents(arg0)
}

var _countResources func(arg0 uint32) ResourceCount

// CountResources.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529322-countresources
func CountResources(arg0 uint32) ResourceCount {
	if _countResources == nil {
		panic("coreservices: symbol CountResources not loaded")
	}
	return _countResources(arg0)
}

var _countTaggedCollectionItems func(arg0 Collection, arg1 CollectionTag) int32

// CountTaggedCollectionItems.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551379-counttaggedcollectionitems
func CountTaggedCollectionItems(arg0 Collection, arg1 CollectionTag) int32 {
	if _countTaggedCollectionItems == nil {
		panic("coreservices: symbol CountTaggedCollectionItems not loaded")
	}
	return _countTaggedCollectionItems(arg0, arg1)
}

var _countTypes func() ResourceCount

// CountTypes.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529346-counttypes
func CountTypes() ResourceCount {
	if _countTypes == nil {
		panic("coreservices: symbol CountTypes not loaded")
	}
	return _countTypes()
}

var _countUnicodeMappings func(arg0 uintptr, arg1 ConstUnicodeMappingPtr, arg2 uintptr) int32

// CountUnicodeMappings counts available mappings that meet the specified matchingcriteria.
//
// See: https://developer.apple.com/documentation/coreservices/1433665-countunicodemappings
func CountUnicodeMappings(arg0 uintptr, arg1 ConstUnicodeMappingPtr, arg2 uintptr) int32 {
	if _countUnicodeMappings == nil {
		panic("coreservices: symbol CountUnicodeMappings not loaded")
	}
	return _countUnicodeMappings(arg0, arg1, arg2)
}

var _createCompDescriptor func(comparisonOperator uint32, operand1 unsafe.Pointer, operand2 unsafe.Pointer, disposeInputs unsafe.Pointer, theDescriptor unsafe.Pointer) int16

// CreateCompDescriptor creates a comparison descriptor that specifies how to compare one or more Apple event objects with either another Apple event object or a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1449155-createcompdescriptor
func CreateCompDescriptor(comparisonOperator uint32, operand1 unsafe.Pointer, operand2 unsafe.Pointer, disposeInputs unsafe.Pointer, theDescriptor unsafe.Pointer) int16 {
	if _createCompDescriptor == nil {
		panic("coreservices: symbol CreateCompDescriptor not loaded")
	}
	return _createCompDescriptor(comparisonOperator, operand1, operand2, disposeInputs, theDescriptor)
}

var _createLogicalDescriptor func(theLogicalTerms unsafe.Pointer, theLogicOperator uint32, disposeInputs unsafe.Pointer, theDescriptor unsafe.Pointer) int16

// CreateLogicalDescriptor creates a logical descriptor that specifies a logical operator and one or more logical terms for the Apple Event Manager to evaluate.
//
// See: https://developer.apple.com/documentation/coreservices/1445212-createlogicaldescriptor
func CreateLogicalDescriptor(theLogicalTerms unsafe.Pointer, theLogicOperator uint32, disposeInputs unsafe.Pointer, theDescriptor unsafe.Pointer) int16 {
	if _createLogicalDescriptor == nil {
		panic("coreservices: symbol CreateLogicalDescriptor not loaded")
	}
	return _createLogicalDescriptor(theLogicalTerms, theLogicOperator, disposeInputs, theDescriptor)
}

var _createObjSpecifier func(desiredClass uint32, theContainer unsafe.Pointer, keyForm uint32, keyData unsafe.Pointer, disposeInputs unsafe.Pointer, objSpecifier unsafe.Pointer) int16

// CreateObjSpecifier assembles an object specifier that identifies one or more Apple event objects, from other descriptors.
//
// See: https://developer.apple.com/documentation/coreservices/1450244-createobjspecifier
func CreateObjSpecifier(desiredClass uint32, theContainer unsafe.Pointer, keyForm uint32, keyData unsafe.Pointer, disposeInputs unsafe.Pointer, objSpecifier unsafe.Pointer) int16 {
	if _createObjSpecifier == nil {
		panic("coreservices: symbol CreateObjSpecifier not loaded")
	}
	return _createObjSpecifier(desiredClass, theContainer, keyForm, keyData, disposeInputs, objSpecifier)
}

var _createOffsetDescriptor func(theOffset unsafe.Pointer, theDescriptor unsafe.Pointer) int16

// CreateOffsetDescriptor creates an offset descriptor that specifies the position of an element in relation to the beginning or end of its container.
//
// See: https://developer.apple.com/documentation/coreservices/1444957-createoffsetdescriptor
func CreateOffsetDescriptor(theOffset unsafe.Pointer, theDescriptor unsafe.Pointer) int16 {
	if _createOffsetDescriptor == nil {
		panic("coreservices: symbol CreateOffsetDescriptor not loaded")
	}
	return _createOffsetDescriptor(theOffset, theDescriptor)
}

var _createRangeDescriptor func(rangeStart unsafe.Pointer, rangeStop unsafe.Pointer, disposeInputs unsafe.Pointer, theDescriptor unsafe.Pointer) int16

// CreateRangeDescriptor creates a range descriptor that specifies a series of consecutive elements in the same container.
//
// See: https://developer.apple.com/documentation/coreservices/1444087-createrangedescriptor
func CreateRangeDescriptor(rangeStart unsafe.Pointer, rangeStop unsafe.Pointer, disposeInputs unsafe.Pointer, theDescriptor unsafe.Pointer) int16 {
	if _createRangeDescriptor == nil {
		panic("coreservices: symbol CreateRangeDescriptor not loaded")
	}
	return _createRangeDescriptor(rangeStart, rangeStop, disposeInputs, theDescriptor)
}

var _createTextEncoding func(arg0 TextEncodingBase, arg1 TextEncodingVariant, arg2 TextEncodingFormat) TextEncoding

// CreateTextEncoding creates and returns a text encoding specification.
//
// See: https://developer.apple.com/documentation/coreservices/1399639-createtextencoding
func CreateTextEncoding(arg0 TextEncodingBase, arg1 TextEncodingVariant, arg2 TextEncodingFormat) TextEncoding {
	if _createTextEncoding == nil {
		panic("coreservices: symbol CreateTextEncoding not loaded")
	}
	return _createTextEncoding(arg0, arg1, arg2)
}

var _createTextToUnicodeInfo func(arg0 ConstUnicodeMappingPtr, arg1 TextToUnicodeInfo) int32

// CreateTextToUnicodeInfo creates and returns a Unicode converter object containinginformation required for converting strings from a non-Unicode encodingto Unicode.
//
// See: https://developer.apple.com/documentation/coreservices/1433598-createtexttounicodeinfo
func CreateTextToUnicodeInfo(arg0 ConstUnicodeMappingPtr, arg1 TextToUnicodeInfo) int32 {
	if _createTextToUnicodeInfo == nil {
		panic("coreservices: symbol CreateTextToUnicodeInfo not loaded")
	}
	return _createTextToUnicodeInfo(arg0, arg1)
}

var _createTextToUnicodeInfoByEncoding func(arg0 TextEncoding, arg1 TextToUnicodeInfo) int32

// CreateTextToUnicodeInfoByEncoding based on the given text encoding specification, createsand returns a Unicode converter object containing information requiredfor converting strings from the specified non-Unicode encoding toUnicode.
//
// See: https://developer.apple.com/documentation/coreservices/1433560-createtexttounicodeinfobyencodin
func CreateTextToUnicodeInfoByEncoding(arg0 TextEncoding, arg1 TextToUnicodeInfo) int32 {
	if _createTextToUnicodeInfoByEncoding == nil {
		panic("coreservices: symbol CreateTextToUnicodeInfoByEncoding not loaded")
	}
	return _createTextToUnicodeInfoByEncoding(arg0, arg1)
}

var _createThreadPool func(arg0 ThreadStyle, arg1 int16, arg2 corefoundation.CGSize) int16

// CreateThreadPool.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574200-createthreadpool
func CreateThreadPool(arg0 ThreadStyle, arg1 int16, arg2 corefoundation.CGSize) int16 {
	if _createThreadPool == nil {
		panic("coreservices: symbol CreateThreadPool not loaded")
	}
	return _createThreadPool(arg0, arg1, arg2)
}

var _createUnicodeToTextInfo func(arg0 ConstUnicodeMappingPtr, arg1 UnicodeToTextInfo) int32

// CreateUnicodeToTextInfo creates and returns a Unicode converter object containinginformation required for converting strings from Unicode to a non-Unicodeencoding.
//
// See: https://developer.apple.com/documentation/coreservices/1433522-createunicodetotextinfo
func CreateUnicodeToTextInfo(arg0 ConstUnicodeMappingPtr, arg1 UnicodeToTextInfo) int32 {
	if _createUnicodeToTextInfo == nil {
		panic("coreservices: symbol CreateUnicodeToTextInfo not loaded")
	}
	return _createUnicodeToTextInfo(arg0, arg1)
}

var _createUnicodeToTextInfoByEncoding func(arg0 TextEncoding, arg1 UnicodeToTextInfo) int32

// CreateUnicodeToTextInfoByEncoding based on the given text encoding specification for theconverted text, creates and returns a Unicode converter object containinginformation required for converting strings from Unicode to thespecified non-Unicode encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1433550-createunicodetotextinfobyencodin
func CreateUnicodeToTextInfoByEncoding(arg0 TextEncoding, arg1 UnicodeToTextInfo) int32 {
	if _createUnicodeToTextInfoByEncoding == nil {
		panic("coreservices: symbol CreateUnicodeToTextInfoByEncoding not loaded")
	}
	return _createUnicodeToTextInfoByEncoding(arg0, arg1)
}

var _createUnicodeToTextRunInfo func(arg0 uintptr, arg1 UnicodeMapping, arg2 UnicodeToTextRunInfo) int32

// CreateUnicodeToTextRunInfo creates and returns a Unicode converter object containingthe information required for converting a Unicode text string tostrings in one or more non-Unicode encodings.
//
// See: https://developer.apple.com/documentation/coreservices/1433632-createunicodetotextruninfo
func CreateUnicodeToTextRunInfo(arg0 uintptr, arg1 UnicodeMapping, arg2 UnicodeToTextRunInfo) int32 {
	if _createUnicodeToTextRunInfo == nil {
		panic("coreservices: symbol CreateUnicodeToTextRunInfo not loaded")
	}
	return _createUnicodeToTextRunInfo(arg0, arg1, arg2)
}

var _createUnicodeToTextRunInfoByEncoding func(arg0 uintptr, arg1 TextEncoding, arg2 UnicodeToTextRunInfo) int32

// CreateUnicodeToTextRunInfoByEncoding based on the given text encoding specifications for theconverted text runs, creates and returns a Unicode converter objectcontaining information required for converting strings from Unicodeto one or more specified non-Unicode encodings.
//
// See: https://developer.apple.com/documentation/coreservices/1433651-createunicodetotextruninfobyenco
func CreateUnicodeToTextRunInfoByEncoding(arg0 uintptr, arg1 TextEncoding, arg2 UnicodeToTextRunInfo) int32 {
	if _createUnicodeToTextRunInfoByEncoding == nil {
		panic("coreservices: symbol CreateUnicodeToTextRunInfoByEncoding not loaded")
	}
	return _createUnicodeToTextRunInfoByEncoding(arg0, arg1, arg2)
}

var _createUnicodeToTextRunInfoByScriptCode func(arg0 uintptr, arg1 unsafe.Pointer, arg2 UnicodeToTextRunInfo) int32

// CreateUnicodeToTextRunInfoByScriptCode based on the given script codes for the converted textruns, creates and returns a Unicode converter object containinginformation required for converting strings from Unicode to oneor more specified non-Unicode encodings.
//
// See: https://developer.apple.com/documentation/coreservices/1433657-createunicodetotextruninfobyscri
func CreateUnicodeToTextRunInfoByScriptCode(arg0 uintptr, arg1 unsafe.Pointer, arg2 UnicodeToTextRunInfo) int32 {
	if _createUnicodeToTextRunInfoByScriptCode == nil {
		panic("coreservices: symbol CreateUnicodeToTextRunInfoByScriptCode not loaded")
	}
	return _createUnicodeToTextRunInfoByScriptCode(arg0, arg1, arg2)
}

var _curResFile func() ResFileRefNum

// CurResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529286-curresfile
func CurResFile() ResFileRefNum {
	if _curResFile == nil {
		panic("coreservices: symbol CurResFile not loaded")
	}
	return _curResFile()
}

var _currentProcessorSpeed func() int16

// CurrentProcessorSpeed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427463-currentprocessorspeed
func CurrentProcessorSpeed() int16 {
	if _currentProcessorSpeed == nil {
		panic("coreservices: symbol CurrentProcessorSpeed not loaded")
	}
	return _currentProcessorSpeed()
}

var _dCSCopyTextDefinition func(dictionary DCSDictionaryRef, textString corefoundation.CFStringRef, range_ unsafe.Pointer) corefoundation.CFStringRef

// DCSCopyTextDefinition returns the definition associated with the provided text range.
//
// See: https://developer.apple.com/documentation/coreservices/1446842-dcscopytextdefinition
func DCSCopyTextDefinition(dictionary DCSDictionaryRef, textString corefoundation.CFStringRef, range_ unsafe.Pointer) corefoundation.CFStringRef {
	if _dCSCopyTextDefinition == nil {
		panic("coreservices: symbol DCSCopyTextDefinition not loaded")
	}
	return _dCSCopyTextDefinition(dictionary, textString, range_)
}

var _dCSGetTermRangeInString func(dictionary DCSDictionaryRef, textString corefoundation.CFStringRef, offset int) corefoundation.CFRange

// DCSGetTermRangeInString determines the range of the longest word or phrase with respect to an offset.
//
// See: https://developer.apple.com/documentation/coreservices/1450556-dcsgettermrangeinstring
func DCSGetTermRangeInString(dictionary DCSDictionaryRef, textString corefoundation.CFStringRef, offset int) corefoundation.CFRange {
	if _dCSGetTermRangeInString == nil {
		panic("coreservices: symbol DCSGetTermRangeInString not loaded")
	}
	return _dCSGetTermRangeInString(dictionary, textString, offset)
}

var _debugAssert func(arg0 uint32, arg1 uint32, arg2 int8, arg3 int8, arg4 int8, arg5 int8, arg6 int)

// DebugAssert.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517794-debugassert
func DebugAssert(arg0 uint32, arg1 uint32, arg2 int8, arg3 int8, arg4 int8, arg5 int8, arg6 int) {
	if _debugAssert == nil {
		panic("coreservices: symbol DebugAssert not loaded")
	}
	_debugAssert(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _decrementAtomic func(arg0 int32) int32

// DecrementAtomic.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490603-decrementatomic
func DecrementAtomic(arg0 int32) int32 {
	if _decrementAtomic == nil {
		panic("coreservices: symbol DecrementAtomic not loaded")
	}
	return _decrementAtomic(arg0)
}

var _decrementAtomic16 func(arg0 int16) int16

// DecrementAtomic16.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490579-decrementatomic16
func DecrementAtomic16(arg0 int16) int16 {
	if _decrementAtomic16 == nil {
		panic("coreservices: symbol DecrementAtomic16 not loaded")
	}
	return _decrementAtomic16(arg0)
}

var _decrementAtomic8 func(arg0 int8) int8

// DecrementAtomic8.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490570-decrementatomic8
func DecrementAtomic8(arg0 int8) int8 {
	if _decrementAtomic8 == nil {
		panic("coreservices: symbol DecrementAtomic8 not loaded")
	}
	return _decrementAtomic8(arg0)
}

var _delay func(arg0 uint, arg1 uint)

// Delay.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533332-delay
func Delay(arg0 uint, arg1 uint) {
	if _delay == nil {
		panic("coreservices: symbol Delay not loaded")
	}
	_delay(arg0, arg1)
}

var _delegateComponentCall func(arg0 ComponentParameters, arg1 ComponentInstance) ComponentResult

// DelegateComponentCall allows your component to pass on a request to a specified component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516373-delegatecomponentcall
func DelegateComponentCall(arg0 ComponentParameters, arg1 ComponentInstance) ComponentResult {
	if _delegateComponentCall == nil {
		panic("coreservices: symbol DelegateComponentCall not loaded")
	}
	return _delegateComponentCall(arg0, arg1)
}

var _deleteGestaltValue func(arg0 uint32) int16

// DeleteGestaltValue deletes a [Gestalt] selector code so that it is no longer recognized by [Gestalt].
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1472699-deletegestaltvalue
func DeleteGestaltValue(arg0 uint32) int16 {
	if _deleteGestaltValue == nil {
		panic("coreservices: symbol DeleteGestaltValue not loaded")
	}
	return _deleteGestaltValue(arg0)
}

var _dequeue func(arg0 QElemPtr, arg1 QHdrPtr) int16

// Dequeue.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533381-dequeue
func Dequeue(arg0 QElemPtr, arg1 QHdrPtr) int16 {
	if _dequeue == nil {
		panic("coreservices: symbol Dequeue not loaded")
	}
	return _dequeue(arg0, arg1)
}

var _detachResource func(arg0 uintptr)

// DetachResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529238-detachresource
func DetachResource(arg0 uintptr) {
	if _detachResource == nil {
		panic("coreservices: symbol DetachResource not loaded")
	}
	_detachResource(arg0)
}

var _detachResourceFile func(arg0 ResFileRefNum) int16

// DetachResourceFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529246-detachresourcefile
func DetachResourceFile(arg0 ResFileRefNum) int16 {
	if _detachResourceFile == nil {
		panic("coreservices: symbol DetachResourceFile not loaded")
	}
	return _detachResourceFile(arg0)
}

var _determineIfPathIsEnclosedByFolder func(arg0 uintptr, arg1 uint32, arg2 uint8, arg3 bool, arg4 bool) int16

// DetermineIfPathIsEnclosedByFolder.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389166-determineifpathisenclosedbyfolde
func DetermineIfPathIsEnclosedByFolder(arg0 uintptr, arg1 uint32, arg2 uint8, arg3 bool, arg4 bool) int16 {
	if _determineIfPathIsEnclosedByFolder == nil {
		panic("coreservices: symbol DetermineIfPathIsEnclosedByFolder not loaded")
	}
	return _determineIfPathIsEnclosedByFolder(arg0, arg1, arg2, arg3, arg4)
}

var _disposeAECoerceDescUPP func(userUPP AECoerceDescUPP)

// DisposeAECoerceDescUPP disposes of a universal procedure pointer to a function that coerces data stored in a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1448721-disposeaecoercedescupp
func DisposeAECoerceDescUPP(userUPP AECoerceDescUPP) {
	if _disposeAECoerceDescUPP == nil {
		panic("coreservices: symbol DisposeAECoerceDescUPP not loaded")
	}
	_disposeAECoerceDescUPP(userUPP)
}

var _disposeAECoercePtrUPP func(userUPP AECoercePtrUPP)

// DisposeAECoercePtrUPP disposes of a universal procedure pointer to a function that coerces data stored in a buffer.
//
// See: https://developer.apple.com/documentation/coreservices/1450664-disposeaecoerceptrupp
func DisposeAECoercePtrUPP(userUPP AECoercePtrUPP) {
	if _disposeAECoercePtrUPP == nil {
		panic("coreservices: symbol DisposeAECoercePtrUPP not loaded")
	}
	_disposeAECoercePtrUPP(userUPP)
}

var _disposeAEDisposeExternalUPP func(userUPP AEDisposeExternalUPP)

// DisposeAEDisposeExternalUPP disposes of a universal procedure pointer to a function that disposes of data supplied to the [AECreateDescFromExternalPtr] function.
//
// See: https://developer.apple.com/documentation/coreservices/1447284-disposeaedisposeexternalupp
func DisposeAEDisposeExternalUPP(userUPP AEDisposeExternalUPP) {
	if _disposeAEDisposeExternalUPP == nil {
		panic("coreservices: symbol DisposeAEDisposeExternalUPP not loaded")
	}
	_disposeAEDisposeExternalUPP(userUPP)
}

var _disposeAEEventHandlerUPP func(userUPP AEEventHandlerUPP)

// DisposeAEEventHandlerUPP disposes of a universal procedure pointer to an event handler function.
//
// See: https://developer.apple.com/documentation/coreservices/1442066-disposeaeeventhandlerupp
func DisposeAEEventHandlerUPP(userUPP AEEventHandlerUPP) {
	if _disposeAEEventHandlerUPP == nil {
		panic("coreservices: symbol DisposeAEEventHandlerUPP not loaded")
	}
	_disposeAEEventHandlerUPP(userUPP)
}

var _disposeCollection func(arg0 Collection)

// DisposeCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551429-disposecollection
func DisposeCollection(arg0 Collection) {
	if _disposeCollection == nil {
		panic("coreservices: symbol DisposeCollection not loaded")
	}
	_disposeCollection(arg0)
}

var _disposeCollectionExceptionUPP func(arg0 CollectionExceptionUPP)

// DisposeCollectionExceptionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551385-disposecollectionexceptionupp
func DisposeCollectionExceptionUPP(arg0 CollectionExceptionUPP) {
	if _disposeCollectionExceptionUPP == nil {
		panic("coreservices: symbol DisposeCollectionExceptionUPP not loaded")
	}
	_disposeCollectionExceptionUPP(arg0)
}

var _disposeCollectionFlattenUPP func(arg0 CollectionFlattenUPP)

// DisposeCollectionFlattenUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551454-disposecollectionflattenupp
func DisposeCollectionFlattenUPP(arg0 CollectionFlattenUPP) {
	if _disposeCollectionFlattenUPP == nil {
		panic("coreservices: symbol DisposeCollectionFlattenUPP not loaded")
	}
	_disposeCollectionFlattenUPP(arg0)
}

var _disposeComponentFunctionUPP func(arg0 ComponentFunctionUPP)

// DisposeComponentFunctionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516369-disposecomponentfunctionupp
func DisposeComponentFunctionUPP(arg0 ComponentFunctionUPP) {
	if _disposeComponentFunctionUPP == nil {
		panic("coreservices: symbol DisposeComponentFunctionUPP not loaded")
	}
	_disposeComponentFunctionUPP(arg0)
}

var _disposeComponentMPWorkFunctionUPP func(arg0 ComponentMPWorkFunctionUPP)

// DisposeComponentMPWorkFunctionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516328-disposecomponentmpworkfunctionup
func DisposeComponentMPWorkFunctionUPP(arg0 ComponentMPWorkFunctionUPP) {
	if _disposeComponentMPWorkFunctionUPP == nil {
		panic("coreservices: symbol DisposeComponentMPWorkFunctionUPP not loaded")
	}
	_disposeComponentMPWorkFunctionUPP(arg0)
}

var _disposeComponentRoutineUPP func(arg0 ComponentRoutineUPP)

// DisposeComponentRoutineUPP disposes of the universal procedure pointer (UPP) to a component routine callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516334-disposecomponentroutineupp
func DisposeComponentRoutineUPP(arg0 ComponentRoutineUPP) {
	if _disposeComponentRoutineUPP == nil {
		panic("coreservices: symbol DisposeComponentRoutineUPP not loaded")
	}
	_disposeComponentRoutineUPP(arg0)
}

var _disposeDebugAssertOutputHandlerUPP func(arg0 DebugAssertOutputHandlerUPP)

// DisposeDebugAssertOutputHandlerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517746-disposedebugassertoutputhandleru
func DisposeDebugAssertOutputHandlerUPP(arg0 DebugAssertOutputHandlerUPP) {
	if _disposeDebugAssertOutputHandlerUPP == nil {
		panic("coreservices: symbol DisposeDebugAssertOutputHandlerUPP not loaded")
	}
	_disposeDebugAssertOutputHandlerUPP(arg0)
}

var _disposeDebugComponent func(arg0 uint32) int32

// DisposeDebugComponent.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517752-disposedebugcomponent
func DisposeDebugComponent(arg0 uint32) int32 {
	if _disposeDebugComponent == nil {
		panic("coreservices: symbol DisposeDebugComponent not loaded")
	}
	return _disposeDebugComponent(arg0)
}

var _disposeDebugComponentCallbackUPP func(arg0 DebugComponentCallbackUPP)

// DisposeDebugComponentCallbackUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517724-disposedebugcomponentcallbackupp
func DisposeDebugComponentCallbackUPP(arg0 DebugComponentCallbackUPP) {
	if _disposeDebugComponentCallbackUPP == nil {
		panic("coreservices: symbol DisposeDebugComponentCallbackUPP not loaded")
	}
	_disposeDebugComponentCallbackUPP(arg0)
}

var _disposeDebuggerDisposeThreadUPP func(arg0 DebuggerDisposeThreadUPP)

// DisposeDebuggerDisposeThreadUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574277-disposedebuggerdisposethreadupp
func DisposeDebuggerDisposeThreadUPP(arg0 DebuggerDisposeThreadUPP) {
	if _disposeDebuggerDisposeThreadUPP == nil {
		panic("coreservices: symbol DisposeDebuggerDisposeThreadUPP not loaded")
	}
	_disposeDebuggerDisposeThreadUPP(arg0)
}

var _disposeDebuggerNewThreadUPP func(arg0 DebuggerNewThreadUPP)

// DisposeDebuggerNewThreadUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574218-disposedebuggernewthreadupp
func DisposeDebuggerNewThreadUPP(arg0 DebuggerNewThreadUPP) {
	if _disposeDebuggerNewThreadUPP == nil {
		panic("coreservices: symbol DisposeDebuggerNewThreadUPP not loaded")
	}
	_disposeDebuggerNewThreadUPP(arg0)
}

var _disposeDebuggerThreadSchedulerUPP func(arg0 DebuggerThreadSchedulerUPP)

// DisposeDebuggerThreadSchedulerUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574286-disposedebuggerthreadschedulerup
func DisposeDebuggerThreadSchedulerUPP(arg0 DebuggerThreadSchedulerUPP) {
	if _disposeDebuggerThreadSchedulerUPP == nil {
		panic("coreservices: symbol DisposeDebuggerThreadSchedulerUPP not loaded")
	}
	_disposeDebuggerThreadSchedulerUPP(arg0)
}

var _disposeDeferredTaskUPP func(arg0 DeferredTaskUPP)

// DisposeDeferredTaskUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533363-disposedeferredtaskupp
func DisposeDeferredTaskUPP(arg0 DeferredTaskUPP) {
	if _disposeDeferredTaskUPP == nil {
		panic("coreservices: symbol DisposeDeferredTaskUPP not loaded")
	}
	_disposeDeferredTaskUPP(arg0)
}

var _disposeExceptionHandlerUPP func(arg0 ExceptionHandlerUPP)

// DisposeExceptionHandlerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1500674-disposeexceptionhandlerupp
func DisposeExceptionHandlerUPP(arg0 ExceptionHandlerUPP) {
	if _disposeExceptionHandlerUPP == nil {
		panic("coreservices: symbol DisposeExceptionHandlerUPP not loaded")
	}
	_disposeExceptionHandlerUPP(arg0)
}

var _disposeFNSubscriptionUPP func(arg0 FNSubscriptionUPP)

// DisposeFNSubscriptionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566772-disposefnsubscriptionupp
func DisposeFNSubscriptionUPP(arg0 FNSubscriptionUPP) {
	if _disposeFNSubscriptionUPP == nil {
		panic("coreservices: symbol DisposeFNSubscriptionUPP not loaded")
	}
	_disposeFNSubscriptionUPP(arg0)
}

var _disposeFSVolumeEjectUPP func(arg0 FSVolumeEjectUPP)

// DisposeFSVolumeEjectUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566754-disposefsvolumeejectupp
func DisposeFSVolumeEjectUPP(arg0 FSVolumeEjectUPP) {
	if _disposeFSVolumeEjectUPP == nil {
		panic("coreservices: symbol DisposeFSVolumeEjectUPP not loaded")
	}
	_disposeFSVolumeEjectUPP(arg0)
}

var _disposeFSVolumeMountUPP func(arg0 FSVolumeMountUPP)

// DisposeFSVolumeMountUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565735-disposefsvolumemountupp
func DisposeFSVolumeMountUPP(arg0 FSVolumeMountUPP) {
	if _disposeFSVolumeMountUPP == nil {
		panic("coreservices: symbol DisposeFSVolumeMountUPP not loaded")
	}
	_disposeFSVolumeMountUPP(arg0)
}

var _disposeFSVolumeUnmountUPP func(arg0 FSVolumeUnmountUPP)

// DisposeFSVolumeUnmountUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565476-disposefsvolumeunmountupp
func DisposeFSVolumeUnmountUPP(arg0 FSVolumeUnmountUPP) {
	if _disposeFSVolumeUnmountUPP == nil {
		panic("coreservices: symbol DisposeFSVolumeUnmountUPP not loaded")
	}
	_disposeFSVolumeUnmountUPP(arg0)
}

var _disposeFolderManagerNotificationUPP func(arg0 FolderManagerNotificationUPP)

// DisposeFolderManagerNotificationUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389143-disposefoldermanagernotification
func DisposeFolderManagerNotificationUPP(arg0 FolderManagerNotificationUPP) {
	if _disposeFolderManagerNotificationUPP == nil {
		panic("coreservices: symbol DisposeFolderManagerNotificationUPP not loaded")
	}
	_disposeFolderManagerNotificationUPP(arg0)
}

var _disposeGetMissingComponentResourceUPP func(arg0 GetMissingComponentResourceUPP)

// DisposeGetMissingComponentResourceUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516414-disposegetmissingcomponentresour
func DisposeGetMissingComponentResourceUPP(arg0 GetMissingComponentResourceUPP) {
	if _disposeGetMissingComponentResourceUPP == nil {
		panic("coreservices: symbol DisposeGetMissingComponentResourceUPP not loaded")
	}
	_disposeGetMissingComponentResourceUPP(arg0)
}

var _disposeHandle func(arg0 uintptr)

// DisposeHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506451-disposehandle
func DisposeHandle(arg0 uintptr) {
	if _disposeHandle == nil {
		panic("coreservices: symbol DisposeHandle not loaded")
	}
	_disposeHandle(arg0)
}

var _disposeIOCompletionUPP func(arg0 IOCompletionUPP)

// DisposeIOCompletionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566947-disposeiocompletionupp
func DisposeIOCompletionUPP(arg0 IOCompletionUPP) {
	if _disposeIOCompletionUPP == nil {
		panic("coreservices: symbol DisposeIOCompletionUPP not loaded")
	}
	_disposeIOCompletionUPP(arg0)
}

var _disposeIndexToUCStringUPP func(userUPP IndexToUCStringUPP)

// DisposeIndexToUCStringUPP.
//
// See: https://developer.apple.com/documentation/coreservices/1390390-disposeindextoucstringupp
func DisposeIndexToUCStringUPP(userUPP IndexToUCStringUPP) {
	if _disposeIndexToUCStringUPP == nil {
		panic("coreservices: symbol DisposeIndexToUCStringUPP not loaded")
	}
	_disposeIndexToUCStringUPP(userUPP)
}

var _disposeKCCallbackUPP func(arg0 KCCallbackUPP)

// DisposeKCCallbackUPP.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563023-disposekccallbackupp
func DisposeKCCallbackUPP(arg0 KCCallbackUPP) {
	if _disposeKCCallbackUPP == nil {
		panic("coreservices: symbol DisposeKCCallbackUPP not loaded")
	}
	_disposeKCCallbackUPP(arg0)
}

var _disposeOSLAccessorUPP func(userUPP OSLAccessorUPP)

// DisposeOSLAccessorUPP disposes of a universal procedure pointer to an object accessor function.
//
// See: https://developer.apple.com/documentation/coreservices/1444684-disposeoslaccessorupp
func DisposeOSLAccessorUPP(userUPP OSLAccessorUPP) {
	if _disposeOSLAccessorUPP == nil {
		panic("coreservices: symbol DisposeOSLAccessorUPP not loaded")
	}
	_disposeOSLAccessorUPP(userUPP)
}

var _disposeOSLAdjustMarksUPP func(userUPP OSLAdjustMarksUPP)

// DisposeOSLAdjustMarksUPP disposes of a universal procedure pointer to an object callback adjust marks function.
//
// See: https://developer.apple.com/documentation/coreservices/1443940-disposeosladjustmarksupp
func DisposeOSLAdjustMarksUPP(userUPP OSLAdjustMarksUPP) {
	if _disposeOSLAdjustMarksUPP == nil {
		panic("coreservices: symbol DisposeOSLAdjustMarksUPP not loaded")
	}
	_disposeOSLAdjustMarksUPP(userUPP)
}

var _disposeOSLCompareUPP func(userUPP OSLCompareUPP)

// DisposeOSLCompareUPP disposes of a universal procedure pointer to an object callback comparison function.
//
// See: https://developer.apple.com/documentation/coreservices/1448398-disposeoslcompareupp
func DisposeOSLCompareUPP(userUPP OSLCompareUPP) {
	if _disposeOSLCompareUPP == nil {
		panic("coreservices: symbol DisposeOSLCompareUPP not loaded")
	}
	_disposeOSLCompareUPP(userUPP)
}

var _disposeOSLCountUPP func(userUPP OSLCountUPP)

// DisposeOSLCountUPP disposes of a universal procedure pointer to an object callback count function.
//
// See: https://developer.apple.com/documentation/coreservices/1443984-disposeoslcountupp
func DisposeOSLCountUPP(userUPP OSLCountUPP) {
	if _disposeOSLCountUPP == nil {
		panic("coreservices: symbol DisposeOSLCountUPP not loaded")
	}
	_disposeOSLCountUPP(userUPP)
}

var _disposeOSLDisposeTokenUPP func(userUPP OSLDisposeTokenUPP)

// DisposeOSLDisposeTokenUPP disposes of a universal procedure pointer to an object callback dispose token function.
//
// See: https://developer.apple.com/documentation/coreservices/1442670-disposeosldisposetokenupp
func DisposeOSLDisposeTokenUPP(userUPP OSLDisposeTokenUPP) {
	if _disposeOSLDisposeTokenUPP == nil {
		panic("coreservices: symbol DisposeOSLDisposeTokenUPP not loaded")
	}
	_disposeOSLDisposeTokenUPP(userUPP)
}

var _disposeOSLGetErrDescUPP func(userUPP OSLGetErrDescUPP)

// DisposeOSLGetErrDescUPP disposes of a universal procedure pointer to an object callback get error descriptor function.
//
// See: https://developer.apple.com/documentation/coreservices/1446061-disposeoslgeterrdescupp
func DisposeOSLGetErrDescUPP(userUPP OSLGetErrDescUPP) {
	if _disposeOSLGetErrDescUPP == nil {
		panic("coreservices: symbol DisposeOSLGetErrDescUPP not loaded")
	}
	_disposeOSLGetErrDescUPP(userUPP)
}

var _disposeOSLGetMarkTokenUPP func(userUPP OSLGetMarkTokenUPP)

// DisposeOSLGetMarkTokenUPP disposes of a universal procedure pointer to an object callback get mark function.
//
// See: https://developer.apple.com/documentation/coreservices/1442377-disposeoslgetmarktokenupp
func DisposeOSLGetMarkTokenUPP(userUPP OSLGetMarkTokenUPP) {
	if _disposeOSLGetMarkTokenUPP == nil {
		panic("coreservices: symbol DisposeOSLGetMarkTokenUPP not loaded")
	}
	_disposeOSLGetMarkTokenUPP(userUPP)
}

var _disposeOSLMarkUPP func(userUPP OSLMarkUPP)

// DisposeOSLMarkUPP disposes of a universal procedure pointer to an object callback mark function.
//
// See: https://developer.apple.com/documentation/coreservices/1449253-disposeoslmarkupp
func DisposeOSLMarkUPP(userUPP OSLMarkUPP) {
	if _disposeOSLMarkUPP == nil {
		panic("coreservices: symbol DisposeOSLMarkUPP not loaded")
	}
	_disposeOSLMarkUPP(userUPP)
}

var _disposePtr func(arg0 coreimage.Ptr)

// DisposePtr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506427-disposeptr
func DisposePtr(arg0 coreimage.Ptr) {
	if _disposePtr == nil {
		panic("coreservices: symbol DisposePtr not loaded")
	}
	_disposePtr(arg0)
}

var _disposeResErrUPP func(arg0 ResErrUPP)

// DisposeResErrUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529332-disposereserrupp
func DisposeResErrUPP(arg0 ResErrUPP) {
	if _disposeResErrUPP == nil {
		panic("coreservices: symbol DisposeResErrUPP not loaded")
	}
	_disposeResErrUPP(arg0)
}

var _disposeSelectorFunctionUPP func(arg0 SelectorFunctionUPP)

// DisposeSelectorFunctionUPP disposes of a universal procedure pointer to a selector callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1472149-disposeselectorfunctionupp
func DisposeSelectorFunctionUPP(arg0 SelectorFunctionUPP) {
	if _disposeSelectorFunctionUPP == nil {
		panic("coreservices: symbol DisposeSelectorFunctionUPP not loaded")
	}
	_disposeSelectorFunctionUPP(arg0)
}

var _disposeSleepQUPP func(arg0 SleepQUPP)

// DisposeSleepQUPP.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/coreservices/1427426-disposesleepqupp
func DisposeSleepQUPP(arg0 SleepQUPP) {
	if _disposeSleepQUPP == nil {
		panic("coreservices: symbol DisposeSleepQUPP not loaded")
	}
	_disposeSleepQUPP(arg0)
}

var _disposeTextToUnicodeInfo func(arg0 TextToUnicodeInfo) int32

// DisposeTextToUnicodeInfo releases the memory allocated for the specified Unicodeconverter object.
//
// See: https://developer.apple.com/documentation/coreservices/1433669-disposetexttounicodeinfo
func DisposeTextToUnicodeInfo(arg0 TextToUnicodeInfo) int32 {
	if _disposeTextToUnicodeInfo == nil {
		panic("coreservices: symbol DisposeTextToUnicodeInfo not loaded")
	}
	return _disposeTextToUnicodeInfo(arg0)
}

var _disposeThread func(arg0 ThreadID, arg1 bool) int16

// DisposeThread.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574219-disposethread
func DisposeThread(arg0 ThreadID, arg1 bool) int16 {
	if _disposeThread == nil {
		panic("coreservices: symbol DisposeThread not loaded")
	}
	return _disposeThread(arg0, arg1)
}

var _disposeThreadEntryUPP func(arg0 ThreadEntryUPP)

// DisposeThreadEntryUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574275-disposethreadentryupp
func DisposeThreadEntryUPP(arg0 ThreadEntryUPP) {
	if _disposeThreadEntryUPP == nil {
		panic("coreservices: symbol DisposeThreadEntryUPP not loaded")
	}
	_disposeThreadEntryUPP(arg0)
}

var _disposeThreadSchedulerUPP func(arg0 ThreadSchedulerUPP)

// DisposeThreadSchedulerUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574280-disposethreadschedulerupp
func DisposeThreadSchedulerUPP(arg0 ThreadSchedulerUPP) {
	if _disposeThreadSchedulerUPP == nil {
		panic("coreservices: symbol DisposeThreadSchedulerUPP not loaded")
	}
	_disposeThreadSchedulerUPP(arg0)
}

var _disposeThreadSwitchUPP func(arg0 ThreadSwitchUPP)

// DisposeThreadSwitchUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574264-disposethreadswitchupp
func DisposeThreadSwitchUPP(arg0 ThreadSwitchUPP) {
	if _disposeThreadSwitchUPP == nil {
		panic("coreservices: symbol DisposeThreadSwitchUPP not loaded")
	}
	_disposeThreadSwitchUPP(arg0)
}

var _disposeThreadTerminationUPP func(arg0 ThreadTerminationUPP)

// DisposeThreadTerminationUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574233-disposethreadterminationupp
func DisposeThreadTerminationUPP(arg0 ThreadTerminationUPP) {
	if _disposeThreadTerminationUPP == nil {
		panic("coreservices: symbol DisposeThreadTerminationUPP not loaded")
	}
	_disposeThreadTerminationUPP(arg0)
}

var _disposeTimerUPP func(arg0 TimerUPP)

// DisposeTimerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1550797-disposetimerupp
func DisposeTimerUPP(arg0 TimerUPP) {
	if _disposeTimerUPP == nil {
		panic("coreservices: symbol DisposeTimerUPP not loaded")
	}
	_disposeTimerUPP(arg0)
}

var _disposeUnicodeToTextFallbackUPP func(arg0 UnicodeToTextFallbackUPP)

// DisposeUnicodeToTextFallbackUPP disposes of a a new universal procedure pointer (UPP)to a Unicode-to-text fallback callback.
//
// See: https://developer.apple.com/documentation/coreservices/1433648-disposeunicodetotextfallbackupp
func DisposeUnicodeToTextFallbackUPP(arg0 UnicodeToTextFallbackUPP) {
	if _disposeUnicodeToTextFallbackUPP == nil {
		panic("coreservices: symbol DisposeUnicodeToTextFallbackUPP not loaded")
	}
	_disposeUnicodeToTextFallbackUPP(arg0)
}

var _disposeUnicodeToTextInfo func(arg0 UnicodeToTextInfo) int32

// DisposeUnicodeToTextInfo releases the memory allocated for the specified Unicodeconverter object.
//
// See: https://developer.apple.com/documentation/coreservices/1433564-disposeunicodetotextinfo
func DisposeUnicodeToTextInfo(arg0 UnicodeToTextInfo) int32 {
	if _disposeUnicodeToTextInfo == nil {
		panic("coreservices: symbol DisposeUnicodeToTextInfo not loaded")
	}
	return _disposeUnicodeToTextInfo(arg0)
}

var _disposeUnicodeToTextRunInfo func(arg0 UnicodeToTextRunInfo) int32

// DisposeUnicodeToTextRunInfo releases the memory allocated for the specified Unicodeconverter object.
//
// See: https://developer.apple.com/documentation/coreservices/1433595-disposeunicodetotextruninfo
func DisposeUnicodeToTextRunInfo(arg0 UnicodeToTextRunInfo) int32 {
	if _disposeUnicodeToTextRunInfo == nil {
		panic("coreservices: symbol DisposeUnicodeToTextRunInfo not loaded")
	}
	return _disposeUnicodeToTextRunInfo(arg0)
}

var _durationToAbsolute func(arg0 unsafe.Pointer) unsafe.Pointer

// DurationToAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501247-durationtoabsolute
func DurationToAbsolute(arg0 unsafe.Pointer) unsafe.Pointer {
	if _durationToAbsolute == nil {
		panic("coreservices: symbol DurationToAbsolute not loaded")
	}
	return _durationToAbsolute(arg0)
}

var _durationToNanoseconds func(arg0 unsafe.Pointer) Nanoseconds

// DurationToNanoseconds.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501254-durationtonanoseconds
func DurationToNanoseconds(arg0 unsafe.Pointer) Nanoseconds {
	if _durationToNanoseconds == nil {
		panic("coreservices: symbol DurationToNanoseconds not loaded")
	}
	return _durationToNanoseconds(arg0)
}

var _emptyCollection func(arg0 Collection)

// EmptyCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551422-emptycollection
func EmptyCollection(arg0 Collection) {
	if _emptyCollection == nil {
		panic("coreservices: symbol EmptyCollection not loaded")
	}
	_emptyCollection(arg0)
}

var _emptyHandle func(arg0 uintptr)

// EmptyHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506345-emptyhandle
func EmptyHandle(arg0 uintptr) {
	if _emptyHandle == nil {
		panic("coreservices: symbol EmptyHandle not loaded")
	}
	_emptyHandle(arg0)
}

var _enqueue func(arg0 QElemPtr, arg1 QHdrPtr)

// Enqueue.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533340-enqueue
func Enqueue(arg0 QElemPtr, arg1 QHdrPtr) {
	if _enqueue == nil {
		panic("coreservices: symbol Enqueue not loaded")
	}
	_enqueue(arg0, arg1)
}

var _fNGetDirectoryForSubscription func(arg0 FNSubscriptionRef, arg1 uintptr) int32

// FNGetDirectoryForSubscription.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566615-fngetdirectoryforsubscription
func FNGetDirectoryForSubscription(arg0 FNSubscriptionRef, arg1 uintptr) int32 {
	if _fNGetDirectoryForSubscription == nil {
		panic("coreservices: symbol FNGetDirectoryForSubscription not loaded")
	}
	return _fNGetDirectoryForSubscription(arg0, arg1)
}

var _fNNotify func(arg0 uintptr, arg1 FNMessage, arg2 uintptr) int32

// FNNotify.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565421-fnnotify
func FNNotify(arg0 uintptr, arg1 FNMessage, arg2 uintptr) int32 {
	if _fNNotify == nil {
		panic("coreservices: symbol FNNotify not loaded")
	}
	return _fNNotify(arg0, arg1, arg2)
}

var _fNNotifyAll func(arg0 FNMessage, arg1 uintptr) int32

// FNNotifyAll.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565760-fnnotifyall
func FNNotifyAll(arg0 FNMessage, arg1 uintptr) int32 {
	if _fNNotifyAll == nil {
		panic("coreservices: symbol FNNotifyAll not loaded")
	}
	return _fNNotifyAll(arg0, arg1)
}

var _fNNotifyByPath func(arg0 uint8, arg1 FNMessage, arg2 uintptr) int32

// FNNotifyByPath.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566346-fnnotifybypath
func FNNotifyByPath(arg0 uint8, arg1 FNMessage, arg2 uintptr) int32 {
	if _fNNotifyByPath == nil {
		panic("coreservices: symbol FNNotifyByPath not loaded")
	}
	return _fNNotifyByPath(arg0, arg1, arg2)
}

var _fNSubscribe func(arg0 uintptr, arg1 FNSubscriptionUPP, arg2 uintptr, arg3 FNSubscriptionRef) int32

// FNSubscribe.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565373-fnsubscribe
func FNSubscribe(arg0 uintptr, arg1 FNSubscriptionUPP, arg2 uintptr, arg3 FNSubscriptionRef) int32 {
	if _fNSubscribe == nil {
		panic("coreservices: symbol FNSubscribe not loaded")
	}
	return _fNSubscribe(arg0, arg1, arg2, arg3)
}

var _fNSubscribeByPath func(arg0 uint8, arg1 FNSubscriptionUPP, arg2 uintptr, arg3 FNSubscriptionRef) int32

// FNSubscribeByPath.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566843-fnsubscribebypath
func FNSubscribeByPath(arg0 uint8, arg1 FNSubscriptionUPP, arg2 uintptr, arg3 FNSubscriptionRef) int32 {
	if _fNSubscribeByPath == nil {
		panic("coreservices: symbol FNSubscribeByPath not loaded")
	}
	return _fNSubscribeByPath(arg0, arg1, arg2, arg3)
}

var _fNUnsubscribe func(arg0 FNSubscriptionRef) int32

// FNUnsubscribe.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565232-fnunsubscribe
func FNUnsubscribe(arg0 FNSubscriptionRef) int32 {
	if _fNUnsubscribe == nil {
		panic("coreservices: symbol FNUnsubscribe not loaded")
	}
	return _fNUnsubscribe(arg0)
}

var _fSAllocateFork func(arg0 FSIORefNum, arg1 FSAllocationFlags, arg2 uint16, arg3 int64, arg4 uint64, arg5 uint64) int16

// FSAllocateFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566175-fsallocatefork
func FSAllocateFork(arg0 FSIORefNum, arg1 FSAllocationFlags, arg2 uint16, arg3 int64, arg4 uint64, arg5 uint64) int16 {
	if _fSAllocateFork == nil {
		panic("coreservices: symbol FSAllocateFork not loaded")
	}
	return _fSAllocateFork(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSCancelVolumeOperation func(arg0 FSVolumeOperation) int32

// FSCancelVolumeOperation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565396-fscancelvolumeoperation
func FSCancelVolumeOperation(arg0 FSVolumeOperation) int32 {
	if _fSCancelVolumeOperation == nil {
		panic("coreservices: symbol FSCancelVolumeOperation not loaded")
	}
	return _fSCancelVolumeOperation(arg0)
}

var _fSCatalogSearch func(arg0 FSIterator, arg1 FSSearchParams, arg2 uintptr, arg3 uintptr, arg4 bool, arg5 FSCatalogInfoBitmap, arg6 FSCatalogInfo, arg7 uintptr, arg8 FSSpecPtr, arg9 unsafe.Pointer) int16

// FSCatalogSearch.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565862-fscatalogsearch
func FSCatalogSearch(arg0 FSIterator, arg1 FSSearchParams, arg2 uintptr, arg3 uintptr, arg4 bool, arg5 FSCatalogInfoBitmap, arg6 FSCatalogInfo, arg7 uintptr, arg8 FSSpecPtr, arg9 unsafe.Pointer) int16 {
	if _fSCatalogSearch == nil {
		panic("coreservices: symbol FSCatalogSearch not loaded")
	}
	return _fSCatalogSearch(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

var _fSCloseFork func(arg0 FSIORefNum) int16

// FSCloseFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566900-fsclosefork
func FSCloseFork(arg0 FSIORefNum) int16 {
	if _fSCloseFork == nil {
		panic("coreservices: symbol FSCloseFork not loaded")
	}
	return _fSCloseFork(arg0)
}

var _fSCloseIterator func(arg0 FSIterator) int16

// FSCloseIterator.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566681-fscloseiterator
func FSCloseIterator(arg0 FSIterator) int16 {
	if _fSCloseIterator == nil {
		panic("coreservices: symbol FSCloseIterator not loaded")
	}
	return _fSCloseIterator(arg0)
}

var _fSCompareFSRefs func(arg0 uintptr, arg1 uintptr) int16

// FSCompareFSRefs.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565571-fscomparefsrefs
func FSCompareFSRefs(arg0 uintptr, arg1 uintptr) int16 {
	if _fSCompareFSRefs == nil {
		panic("coreservices: symbol FSCompareFSRefs not loaded")
	}
	return _fSCompareFSRefs(arg0, arg1)
}

var _fSCopyDADiskForVolume func(arg0 uintptr, arg1 diskarbitration.DADiskRef) int32

// FSCopyDADiskForVolume.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565813-fscopydadiskforvolume
func FSCopyDADiskForVolume(arg0 uintptr, arg1 diskarbitration.DADiskRef) int32 {
	if _fSCopyDADiskForVolume == nil {
		panic("coreservices: symbol FSCopyDADiskForVolume not loaded")
	}
	return _fSCopyDADiskForVolume(arg0, arg1)
}

var _fSCopyDiskIDForVolume func(arg0 uintptr, arg1 corefoundation.CFStringRef) int32

// FSCopyDiskIDForVolume.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565655-fscopydiskidforvolume
func FSCopyDiskIDForVolume(arg0 uintptr, arg1 corefoundation.CFStringRef) int32 {
	if _fSCopyDiskIDForVolume == nil {
		panic("coreservices: symbol FSCopyDiskIDForVolume not loaded")
	}
	return _fSCopyDiskIDForVolume(arg0, arg1)
}

var _fSCopyObjectAsync func(arg0 FSFileOperationRef, arg1 uintptr, arg2 uintptr, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32

// FSCopyObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566285-fscopyobjectasync
func FSCopyObjectAsync(arg0 FSFileOperationRef, arg1 uintptr, arg2 uintptr, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32 {
	if _fSCopyObjectAsync == nil {
		panic("coreservices: symbol FSCopyObjectAsync not loaded")
	}
	return _fSCopyObjectAsync(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _fSCopyObjectSync func(arg0 uintptr, arg1 uintptr, arg2 corefoundation.CFStringRef, arg3 uintptr, arg4 uintptr) int32

// FSCopyObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565258-fscopyobjectsync
func FSCopyObjectSync(arg0 uintptr, arg1 uintptr, arg2 corefoundation.CFStringRef, arg3 uintptr, arg4 uintptr) int32 {
	if _fSCopyObjectSync == nil {
		panic("coreservices: symbol FSCopyObjectSync not loaded")
	}
	return _fSCopyObjectSync(arg0, arg1, arg2, arg3, arg4)
}

var _fSCopyURLForVolume func(arg0 uintptr, arg1 corefoundation.CFURLRef) int32

// FSCopyURLForVolume.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566159-fscopyurlforvolume
func FSCopyURLForVolume(arg0 uintptr, arg1 corefoundation.CFURLRef) int32 {
	if _fSCopyURLForVolume == nil {
		panic("coreservices: symbol FSCopyURLForVolume not loaded")
	}
	return _fSCopyURLForVolume(arg0, arg1)
}

var _fSCreateDirectoryUnicode func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uintptr, arg6 FSSpecPtr, arg7 uint32) int16

// FSCreateDirectoryUnicode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565979-fscreatedirectoryunicode
func FSCreateDirectoryUnicode(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uintptr, arg6 FSSpecPtr, arg7 uint32) int16 {
	if _fSCreateDirectoryUnicode == nil {
		panic("coreservices: symbol FSCreateDirectoryUnicode not loaded")
	}
	return _fSCreateDirectoryUnicode(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _fSCreateFileAndOpenForkUnicode func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uint, arg6 uint16, arg7 int8, arg8 FSIORefNum, arg9 uintptr) int32

// FSCreateFileAndOpenForkUnicode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565699-fscreatefileandopenforkunicode
func FSCreateFileAndOpenForkUnicode(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uint, arg6 uint16, arg7 int8, arg8 FSIORefNum, arg9 uintptr) int32 {
	if _fSCreateFileAndOpenForkUnicode == nil {
		panic("coreservices: symbol FSCreateFileAndOpenForkUnicode not loaded")
	}
	return _fSCreateFileAndOpenForkUnicode(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

var _fSCreateFileUnicode func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uintptr, arg6 FSSpecPtr) int16

// FSCreateFileUnicode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565163-fscreatefileunicode
func FSCreateFileUnicode(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uintptr, arg6 FSSpecPtr) int16 {
	if _fSCreateFileUnicode == nil {
		panic("coreservices: symbol FSCreateFileUnicode not loaded")
	}
	return _fSCreateFileUnicode(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _fSCreateFork func(arg0 uintptr, arg1 uint, arg2 uint16) int16

// FSCreateFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565554-fscreatefork
func FSCreateFork(arg0 uintptr, arg1 uint, arg2 uint16) int16 {
	if _fSCreateFork == nil {
		panic("coreservices: symbol FSCreateFork not loaded")
	}
	return _fSCreateFork(arg0, arg1, arg2)
}

var _fSCreateResFile func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uintptr, arg6 FSSpecPtr)

// FSCreateResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529336-fscreateresfile
func FSCreateResFile(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uintptr, arg6 FSSpecPtr) {
	if _fSCreateResFile == nil {
		panic("coreservices: symbol FSCreateResFile not loaded")
	}
	_fSCreateResFile(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _fSCreateResourceFile func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uint, arg6 uint16, arg7 uintptr, arg8 FSSpecPtr) int16

// FSCreateResourceFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529242-fscreateresourcefile
func FSCreateResourceFile(arg0 uintptr, arg1 uint, arg2 uint16, arg3 FSCatalogInfoBitmap, arg4 FSCatalogInfo, arg5 uint, arg6 uint16, arg7 uintptr, arg8 FSSpecPtr) int16 {
	if _fSCreateResourceFile == nil {
		panic("coreservices: symbol FSCreateResourceFile not loaded")
	}
	return _fSCreateResourceFile(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _fSCreateResourceFork func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 uint32) int16

// FSCreateResourceFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529360-fscreateresourcefork
func FSCreateResourceFork(arg0 uintptr, arg1 uint, arg2 uint16, arg3 uint32) int16 {
	if _fSCreateResourceFork == nil {
		panic("coreservices: symbol FSCreateResourceFork not loaded")
	}
	return _fSCreateResourceFork(arg0, arg1, arg2, arg3)
}

var _fSCreateStringFromHFSUniStr func(arg0 corefoundation.CFAllocatorRef, arg1 unsafe.Pointer) corefoundation.CFStringRef

// FSCreateStringFromHFSUniStr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566957-fscreatestringfromhfsunistr
func FSCreateStringFromHFSUniStr(arg0 corefoundation.CFAllocatorRef, arg1 unsafe.Pointer) corefoundation.CFStringRef {
	if _fSCreateStringFromHFSUniStr == nil {
		panic("coreservices: symbol FSCreateStringFromHFSUniStr not loaded")
	}
	return _fSCreateStringFromHFSUniStr(arg0, arg1)
}

var _fSCreateVolumeOperation func(arg0 FSVolumeOperation) int32

// FSCreateVolumeOperation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566671-fscreatevolumeoperation
func FSCreateVolumeOperation(arg0 FSVolumeOperation) int32 {
	if _fSCreateVolumeOperation == nil {
		panic("coreservices: symbol FSCreateVolumeOperation not loaded")
	}
	return _fSCreateVolumeOperation(arg0)
}

var _fSDeleteFork func(arg0 uintptr, arg1 uint, arg2 uint16) int16

// FSDeleteFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565370-fsdeletefork
func FSDeleteFork(arg0 uintptr, arg1 uint, arg2 uint16) int16 {
	if _fSDeleteFork == nil {
		panic("coreservices: symbol FSDeleteFork not loaded")
	}
	return _fSDeleteFork(arg0, arg1, arg2)
}

var _fSDeleteObject func(arg0 uintptr) int16

// FSDeleteObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566089-fsdeleteobject
func FSDeleteObject(arg0 uintptr) int16 {
	if _fSDeleteObject == nil {
		panic("coreservices: symbol FSDeleteObject not loaded")
	}
	return _fSDeleteObject(arg0)
}

var _fSDetermineIfRefIsEnclosedByFolder func(arg0 uintptr, arg1 uint32, arg2 uintptr, arg3 bool) int16

// FSDetermineIfRefIsEnclosedByFolder.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389228-fsdetermineifrefisenclosedbyfold
func FSDetermineIfRefIsEnclosedByFolder(arg0 uintptr, arg1 uint32, arg2 uintptr, arg3 bool) int16 {
	if _fSDetermineIfRefIsEnclosedByFolder == nil {
		panic("coreservices: symbol FSDetermineIfRefIsEnclosedByFolder not loaded")
	}
	return _fSDetermineIfRefIsEnclosedByFolder(arg0, arg1, arg2, arg3)
}

var _fSDisposeVolumeOperation func(arg0 FSVolumeOperation) int32

// FSDisposeVolumeOperation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565382-fsdisposevolumeoperation
func FSDisposeVolumeOperation(arg0 FSVolumeOperation) int32 {
	if _fSDisposeVolumeOperation == nil {
		panic("coreservices: symbol FSDisposeVolumeOperation not loaded")
	}
	return _fSDisposeVolumeOperation(arg0)
}

var _fSEjectVolumeAsync func(arg0 uintptr, arg1 uintptr, arg2 FSVolumeOperation, arg3 FSVolumeEjectUPP, arg4 corefoundation.CFRunLoopRef, arg5 corefoundation.CFStringRef) int32

// FSEjectVolumeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566919-fsejectvolumeasync
func FSEjectVolumeAsync(arg0 uintptr, arg1 uintptr, arg2 FSVolumeOperation, arg3 FSVolumeEjectUPP, arg4 corefoundation.CFRunLoopRef, arg5 corefoundation.CFStringRef) int32 {
	if _fSEjectVolumeAsync == nil {
		panic("coreservices: symbol FSEjectVolumeAsync not loaded")
	}
	return _fSEjectVolumeAsync(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSEjectVolumeSync func(arg0 uintptr, arg1 uintptr, arg2 int32) int32

// FSEjectVolumeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566031-fsejectvolumesync
func FSEjectVolumeSync(arg0 uintptr, arg1 uintptr, arg2 int32) int32 {
	if _fSEjectVolumeSync == nil {
		panic("coreservices: symbol FSEjectVolumeSync not loaded")
	}
	return _fSEjectVolumeSync(arg0, arg1, arg2)
}

var _fSEventStreamCopyDescription func(streamRef unsafe.Pointer) corefoundation.CFStringRef

// FSEventStreamCopyDescription.
//
// See: https://developer.apple.com/documentation/coreservices/1442676-fseventstreamcopydescription
func FSEventStreamCopyDescription(streamRef unsafe.Pointer) corefoundation.CFStringRef {
	if _fSEventStreamCopyDescription == nil {
		panic("coreservices: symbol FSEventStreamCopyDescription not loaded")
	}
	return _fSEventStreamCopyDescription(streamRef)
}

var _fSEventStreamCopyPathsBeingWatched func(streamRef unsafe.Pointer) corefoundation.CFArrayRef

// FSEventStreamCopyPathsBeingWatched.
//
// See: https://developer.apple.com/documentation/coreservices/1447917-fseventstreamcopypathsbeingwatch
func FSEventStreamCopyPathsBeingWatched(streamRef unsafe.Pointer) corefoundation.CFArrayRef {
	if _fSEventStreamCopyPathsBeingWatched == nil {
		panic("coreservices: symbol FSEventStreamCopyPathsBeingWatched not loaded")
	}
	return _fSEventStreamCopyPathsBeingWatched(streamRef)
}

var _fSEventStreamCreate func(allocator corefoundation.CFAllocatorRef, callback unsafe.Pointer, context unsafe.Pointer, pathsToWatch corefoundation.CFArrayRef, sinceWhen uint64, latency float64, flags uint32) unsafe.Pointer

// FSEventStreamCreate.
//
// See: https://developer.apple.com/documentation/coreservices/1443980-fseventstreamcreate
func FSEventStreamCreate(allocator corefoundation.CFAllocatorRef, callback unsafe.Pointer, context unsafe.Pointer, pathsToWatch corefoundation.CFArrayRef, sinceWhen uint64, latency float64, flags uint32) unsafe.Pointer {
	if _fSEventStreamCreate == nil {
		panic("coreservices: symbol FSEventStreamCreate not loaded")
	}
	return _fSEventStreamCreate(allocator, callback, context, pathsToWatch, sinceWhen, latency, flags)
}

var _fSEventStreamCreateRelativeToDevice func(allocator corefoundation.CFAllocatorRef, callback unsafe.Pointer, context unsafe.Pointer, deviceToWatch int32, pathsToWatchRelativeToDevice corefoundation.CFArrayRef, sinceWhen uint64, latency float64, flags uint32) unsafe.Pointer

// FSEventStreamCreateRelativeToDevice.
//
// See: https://developer.apple.com/documentation/coreservices/1447341-fseventstreamcreaterelativetodev
func FSEventStreamCreateRelativeToDevice(allocator corefoundation.CFAllocatorRef, callback unsafe.Pointer, context unsafe.Pointer, deviceToWatch int32, pathsToWatchRelativeToDevice corefoundation.CFArrayRef, sinceWhen uint64, latency float64, flags uint32) unsafe.Pointer {
	if _fSEventStreamCreateRelativeToDevice == nil {
		panic("coreservices: symbol FSEventStreamCreateRelativeToDevice not loaded")
	}
	return _fSEventStreamCreateRelativeToDevice(allocator, callback, context, deviceToWatch, pathsToWatchRelativeToDevice, sinceWhen, latency, flags)
}

var _fSEventStreamFlushAsync func(streamRef unsafe.Pointer) uint64

// FSEventStreamFlushAsync.
//
// See: https://developer.apple.com/documentation/coreservices/1441727-fseventstreamflushasync
func FSEventStreamFlushAsync(streamRef unsafe.Pointer) uint64 {
	if _fSEventStreamFlushAsync == nil {
		panic("coreservices: symbol FSEventStreamFlushAsync not loaded")
	}
	return _fSEventStreamFlushAsync(streamRef)
}

var _fSEventStreamFlushSync func(streamRef unsafe.Pointer)

// FSEventStreamFlushSync.
//
// See: https://developer.apple.com/documentation/coreservices/1445629-fseventstreamflushsync
func FSEventStreamFlushSync(streamRef unsafe.Pointer) {
	if _fSEventStreamFlushSync == nil {
		panic("coreservices: symbol FSEventStreamFlushSync not loaded")
	}
	_fSEventStreamFlushSync(streamRef)
}

var _fSEventStreamGetDeviceBeingWatched func(streamRef unsafe.Pointer) int32

// FSEventStreamGetDeviceBeingWatched.
//
// See: https://developer.apple.com/documentation/coreservices/1449675-fseventstreamgetdevicebeingwatch
func FSEventStreamGetDeviceBeingWatched(streamRef unsafe.Pointer) int32 {
	if _fSEventStreamGetDeviceBeingWatched == nil {
		panic("coreservices: symbol FSEventStreamGetDeviceBeingWatched not loaded")
	}
	return _fSEventStreamGetDeviceBeingWatched(streamRef)
}

var _fSEventStreamGetLatestEventId func(streamRef unsafe.Pointer) uint64

// FSEventStreamGetLatestEventId.
//
// See: https://developer.apple.com/documentation/coreservices/1446030-fseventstreamgetlatesteventid
func FSEventStreamGetLatestEventId(streamRef unsafe.Pointer) uint64 {
	if _fSEventStreamGetLatestEventId == nil {
		panic("coreservices: symbol FSEventStreamGetLatestEventId not loaded")
	}
	return _fSEventStreamGetLatestEventId(streamRef)
}

var _fSEventStreamInvalidate func(streamRef unsafe.Pointer)

// FSEventStreamInvalidate.
//
// See: https://developer.apple.com/documentation/coreservices/1446990-fseventstreaminvalidate
func FSEventStreamInvalidate(streamRef unsafe.Pointer) {
	if _fSEventStreamInvalidate == nil {
		panic("coreservices: symbol FSEventStreamInvalidate not loaded")
	}
	_fSEventStreamInvalidate(streamRef)
}

var _fSEventStreamRelease func(streamRef unsafe.Pointer)

// FSEventStreamRelease.
//
// See: https://developer.apple.com/documentation/coreservices/1445989-fseventstreamrelease
func FSEventStreamRelease(streamRef unsafe.Pointer) {
	if _fSEventStreamRelease == nil {
		panic("coreservices: symbol FSEventStreamRelease not loaded")
	}
	_fSEventStreamRelease(streamRef)
}

var _fSEventStreamRetain func(streamRef unsafe.Pointer)

// FSEventStreamRetain.
//
// See: https://developer.apple.com/documentation/coreservices/1444986-fseventstreamretain
func FSEventStreamRetain(streamRef unsafe.Pointer) {
	if _fSEventStreamRetain == nil {
		panic("coreservices: symbol FSEventStreamRetain not loaded")
	}
	_fSEventStreamRetain(streamRef)
}

var _fSEventStreamSetDispatchQueue func(streamRef unsafe.Pointer, q uintptr)

// FSEventStreamSetDispatchQueue schedules the stream on the specified dispatch queue.
//
// See: https://developer.apple.com/documentation/coreservices/1444164-fseventstreamsetdispatchqueue
func FSEventStreamSetDispatchQueue(streamRef unsafe.Pointer, q dispatch.Queue) {
	if _fSEventStreamSetDispatchQueue == nil {
		panic("coreservices: symbol FSEventStreamSetDispatchQueue not loaded")
	}
	_fSEventStreamSetDispatchQueue(streamRef, uintptr(q.Handle()))
}

var _fSEventStreamSetExclusionPaths func(streamRef unsafe.Pointer, pathsToExclude corefoundation.CFArrayRef) bool

// FSEventStreamSetExclusionPaths.
//
// See: https://developer.apple.com/documentation/coreservices/1444666-fseventstreamsetexclusionpaths
func FSEventStreamSetExclusionPaths(streamRef unsafe.Pointer, pathsToExclude corefoundation.CFArrayRef) bool {
	if _fSEventStreamSetExclusionPaths == nil {
		panic("coreservices: symbol FSEventStreamSetExclusionPaths not loaded")
	}
	return _fSEventStreamSetExclusionPaths(streamRef, pathsToExclude)
}

var _fSEventStreamShow func(streamRef unsafe.Pointer)

// FSEventStreamShow.
//
// See: https://developer.apple.com/documentation/coreservices/1444302-fseventstreamshow
func FSEventStreamShow(streamRef unsafe.Pointer) {
	if _fSEventStreamShow == nil {
		panic("coreservices: symbol FSEventStreamShow not loaded")
	}
	_fSEventStreamShow(streamRef)
}

var _fSEventStreamStart func(streamRef unsafe.Pointer) bool

// FSEventStreamStart.
//
// See: https://developer.apple.com/documentation/coreservices/1448000-fseventstreamstart
func FSEventStreamStart(streamRef unsafe.Pointer) bool {
	if _fSEventStreamStart == nil {
		panic("coreservices: symbol FSEventStreamStart not loaded")
	}
	return _fSEventStreamStart(streamRef)
}

var _fSEventStreamStop func(streamRef unsafe.Pointer)

// FSEventStreamStop.
//
// See: https://developer.apple.com/documentation/coreservices/1447673-fseventstreamstop
func FSEventStreamStop(streamRef unsafe.Pointer) {
	if _fSEventStreamStop == nil {
		panic("coreservices: symbol FSEventStreamStop not loaded")
	}
	_fSEventStreamStop(streamRef)
}

var _fSEventsCopyUUIDForDevice func(dev int32) corefoundation.CFUUID

// FSEventsCopyUUIDForDevice.
//
// See: https://developer.apple.com/documentation/coreservices/1444453-fseventscopyuuidfordevice
func FSEventsCopyUUIDForDevice(dev int32) corefoundation.CFUUID {
	if _fSEventsCopyUUIDForDevice == nil {
		panic("coreservices: symbol FSEventsCopyUUIDForDevice not loaded")
	}
	return _fSEventsCopyUUIDForDevice(dev)
}

var _fSEventsGetCurrentEventId func() uint64

// FSEventsGetCurrentEventId.
//
// See: https://developer.apple.com/documentation/coreservices/1442917-fseventsgetcurrenteventid
func FSEventsGetCurrentEventId() uint64 {
	if _fSEventsGetCurrentEventId == nil {
		panic("coreservices: symbol FSEventsGetCurrentEventId not loaded")
	}
	return _fSEventsGetCurrentEventId()
}

var _fSEventsGetLastEventIdForDeviceBeforeTime func(dev int32, time corefoundation.CFAbsoluteTime) uint64

// FSEventsGetLastEventIdForDeviceBeforeTime.
//
// See: https://developer.apple.com/documentation/coreservices/1449772-fseventsgetlasteventidfordeviceb
func FSEventsGetLastEventIdForDeviceBeforeTime(dev int32, time corefoundation.CFAbsoluteTime) uint64 {
	if _fSEventsGetLastEventIdForDeviceBeforeTime == nil {
		panic("coreservices: symbol FSEventsGetLastEventIdForDeviceBeforeTime not loaded")
	}
	return _fSEventsGetLastEventIdForDeviceBeforeTime(dev, time)
}

var _fSEventsPurgeEventsForDeviceUpToEventId func(dev int32, eventId uint64) bool

// FSEventsPurgeEventsForDeviceUpToEventId.
//
// See: https://developer.apple.com/documentation/coreservices/1447985-fseventspurgeeventsfordeviceupto
func FSEventsPurgeEventsForDeviceUpToEventId(dev int32, eventId uint64) bool {
	if _fSEventsPurgeEventsForDeviceUpToEventId == nil {
		panic("coreservices: symbol FSEventsPurgeEventsForDeviceUpToEventId not loaded")
	}
	return _fSEventsPurgeEventsForDeviceUpToEventId(dev, eventId)
}

var _fSFileOperationCancel func(arg0 FSFileOperationRef) int32

// FSFileOperationCancel.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565866-fsfileoperationcancel
func FSFileOperationCancel(arg0 FSFileOperationRef) int32 {
	if _fSFileOperationCancel == nil {
		panic("coreservices: symbol FSFileOperationCancel not loaded")
	}
	return _fSFileOperationCancel(arg0)
}

var _fSFileOperationCopyStatus func(arg0 FSFileOperationRef, arg1 uintptr, arg2 FSFileOperationStage, arg3 int32, arg4 corefoundation.CFDictionaryRef) int32

// FSFileOperationCopyStatus.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566279-fsfileoperationcopystatus
func FSFileOperationCopyStatus(arg0 FSFileOperationRef, arg1 uintptr, arg2 FSFileOperationStage, arg3 int32, arg4 corefoundation.CFDictionaryRef) int32 {
	if _fSFileOperationCopyStatus == nil {
		panic("coreservices: symbol FSFileOperationCopyStatus not loaded")
	}
	return _fSFileOperationCopyStatus(arg0, arg1, arg2, arg3, arg4)
}

var _fSFileOperationCreate func(arg0 corefoundation.CFAllocatorRef) FSFileOperationRef

// FSFileOperationCreate.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566666-fsfileoperationcreate
func FSFileOperationCreate(arg0 corefoundation.CFAllocatorRef) FSFileOperationRef {
	if _fSFileOperationCreate == nil {
		panic("coreservices: symbol FSFileOperationCreate not loaded")
	}
	return _fSFileOperationCreate(arg0)
}

var _fSFileOperationGetTypeID func() uint

// FSFileOperationGetTypeID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565673-fsfileoperationgettypeid
func FSFileOperationGetTypeID() uint {
	if _fSFileOperationGetTypeID == nil {
		panic("coreservices: symbol FSFileOperationGetTypeID not loaded")
	}
	return _fSFileOperationGetTypeID()
}

var _fSFileOperationScheduleWithRunLoop func(arg0 FSFileOperationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) int32

// FSFileOperationScheduleWithRunLoop.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565845-fsfileoperationschedulewithrunlo
func FSFileOperationScheduleWithRunLoop(arg0 FSFileOperationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) int32 {
	if _fSFileOperationScheduleWithRunLoop == nil {
		panic("coreservices: symbol FSFileOperationScheduleWithRunLoop not loaded")
	}
	return _fSFileOperationScheduleWithRunLoop(arg0, arg1, arg2)
}

var _fSFileOperationUnscheduleFromRunLoop func(arg0 FSFileOperationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) int32

// FSFileOperationUnscheduleFromRunLoop.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565483-fsfileoperationunschedulefromrun
func FSFileOperationUnscheduleFromRunLoop(arg0 FSFileOperationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) int32 {
	if _fSFileOperationUnscheduleFromRunLoop == nil {
		panic("coreservices: symbol FSFileOperationUnscheduleFromRunLoop not loaded")
	}
	return _fSFileOperationUnscheduleFromRunLoop(arg0, arg1, arg2)
}

var _fSFileSecurityCopyAccessControlList func(arg0 FSFileSecurityRef, arg1 unsafe.Pointer) int32

// FSFileSecurityCopyAccessControlList.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565603-fsfilesecuritycopyaccesscontroll
func FSFileSecurityCopyAccessControlList(arg0 FSFileSecurityRef, arg1 unsafe.Pointer) int32 {
	if _fSFileSecurityCopyAccessControlList == nil {
		panic("coreservices: symbol FSFileSecurityCopyAccessControlList not loaded")
	}
	return _fSFileSecurityCopyAccessControlList(arg0, arg1)
}

var _fSFileSecurityCreate func(arg0 corefoundation.CFAllocatorRef) FSFileSecurityRef

// FSFileSecurityCreate.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565742-fsfilesecuritycreate
func FSFileSecurityCreate(arg0 corefoundation.CFAllocatorRef) FSFileSecurityRef {
	if _fSFileSecurityCreate == nil {
		panic("coreservices: symbol FSFileSecurityCreate not loaded")
	}
	return _fSFileSecurityCreate(arg0)
}

var _fSFileSecurityCreateWithFSPermissionInfo func(arg0 corefoundation.CFAllocatorRef, arg1 FSPermissionInfo) FSFileSecurityRef

// FSFileSecurityCreateWithFSPermissionInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566954-fsfilesecuritycreatewithfspermis
func FSFileSecurityCreateWithFSPermissionInfo(arg0 corefoundation.CFAllocatorRef, arg1 FSPermissionInfo) FSFileSecurityRef {
	if _fSFileSecurityCreateWithFSPermissionInfo == nil {
		panic("coreservices: symbol FSFileSecurityCreateWithFSPermissionInfo not loaded")
	}
	return _fSFileSecurityCreateWithFSPermissionInfo(arg0, arg1)
}

var _fSFileSecurityGetGroup func(arg0 FSFileSecurityRef, arg1 uint32) int32

// FSFileSecurityGetGroup.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565811-fsfilesecuritygetgroup
func FSFileSecurityGetGroup(arg0 FSFileSecurityRef, arg1 uint32) int32 {
	if _fSFileSecurityGetGroup == nil {
		panic("coreservices: symbol FSFileSecurityGetGroup not loaded")
	}
	return _fSFileSecurityGetGroup(arg0, arg1)
}

var _fSFileSecurityGetGroupUUID func(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32

// FSFileSecurityGetGroupUUID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566799-fsfilesecuritygetgroupuuid
func FSFileSecurityGetGroupUUID(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32 {
	if _fSFileSecurityGetGroupUUID == nil {
		panic("coreservices: symbol FSFileSecurityGetGroupUUID not loaded")
	}
	return _fSFileSecurityGetGroupUUID(arg0, arg1)
}

var _fSFileSecurityGetMode func(arg0 FSFileSecurityRef, arg1 uint16) int32

// FSFileSecurityGetMode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565106-fsfilesecuritygetmode
func FSFileSecurityGetMode(arg0 FSFileSecurityRef, arg1 uint16) int32 {
	if _fSFileSecurityGetMode == nil {
		panic("coreservices: symbol FSFileSecurityGetMode not loaded")
	}
	return _fSFileSecurityGetMode(arg0, arg1)
}

var _fSFileSecurityGetOwner func(arg0 FSFileSecurityRef, arg1 uint32) int32

// FSFileSecurityGetOwner.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566659-fsfilesecuritygetowner
func FSFileSecurityGetOwner(arg0 FSFileSecurityRef, arg1 uint32) int32 {
	if _fSFileSecurityGetOwner == nil {
		panic("coreservices: symbol FSFileSecurityGetOwner not loaded")
	}
	return _fSFileSecurityGetOwner(arg0, arg1)
}

var _fSFileSecurityGetOwnerUUID func(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32

// FSFileSecurityGetOwnerUUID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566694-fsfilesecuritygetowneruuid
func FSFileSecurityGetOwnerUUID(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32 {
	if _fSFileSecurityGetOwnerUUID == nil {
		panic("coreservices: symbol FSFileSecurityGetOwnerUUID not loaded")
	}
	return _fSFileSecurityGetOwnerUUID(arg0, arg1)
}

var _fSFileSecurityGetTypeID func() uint

// FSFileSecurityGetTypeID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565576-fsfilesecuritygettypeid
func FSFileSecurityGetTypeID() uint {
	if _fSFileSecurityGetTypeID == nil {
		panic("coreservices: symbol FSFileSecurityGetTypeID not loaded")
	}
	return _fSFileSecurityGetTypeID()
}

var _fSFileSecurityRefCreateCopy func(arg0 corefoundation.CFAllocatorRef, arg1 FSFileSecurityRef) FSFileSecurityRef

// FSFileSecurityRefCreateCopy.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566223-fsfilesecurityrefcreatecopy
func FSFileSecurityRefCreateCopy(arg0 corefoundation.CFAllocatorRef, arg1 FSFileSecurityRef) FSFileSecurityRef {
	if _fSFileSecurityRefCreateCopy == nil {
		panic("coreservices: symbol FSFileSecurityRefCreateCopy not loaded")
	}
	return _fSFileSecurityRefCreateCopy(arg0, arg1)
}

var _fSFileSecuritySetAccessControlList func(arg0 FSFileSecurityRef, arg1 unsafe.Pointer) int32

// FSFileSecuritySetAccessControlList.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566029-fsfilesecuritysetaccesscontrolli
func FSFileSecuritySetAccessControlList(arg0 FSFileSecurityRef, arg1 unsafe.Pointer) int32 {
	if _fSFileSecuritySetAccessControlList == nil {
		panic("coreservices: symbol FSFileSecuritySetAccessControlList not loaded")
	}
	return _fSFileSecuritySetAccessControlList(arg0, arg1)
}

var _fSFileSecuritySetGroup func(arg0 FSFileSecurityRef, arg1 uint32) int32

// FSFileSecuritySetGroup.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565296-fsfilesecuritysetgroup
func FSFileSecuritySetGroup(arg0 FSFileSecurityRef, arg1 uint32) int32 {
	if _fSFileSecuritySetGroup == nil {
		panic("coreservices: symbol FSFileSecuritySetGroup not loaded")
	}
	return _fSFileSecuritySetGroup(arg0, arg1)
}

var _fSFileSecuritySetGroupUUID func(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32

// FSFileSecuritySetGroupUUID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566445-fsfilesecuritysetgroupuuid
func FSFileSecuritySetGroupUUID(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32 {
	if _fSFileSecuritySetGroupUUID == nil {
		panic("coreservices: symbol FSFileSecuritySetGroupUUID not loaded")
	}
	return _fSFileSecuritySetGroupUUID(arg0, arg1)
}

var _fSFileSecuritySetMode func(arg0 FSFileSecurityRef, arg1 uint16) int32

// FSFileSecuritySetMode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565105-fsfilesecuritysetmode
func FSFileSecuritySetMode(arg0 FSFileSecurityRef, arg1 uint16) int32 {
	if _fSFileSecuritySetMode == nil {
		panic("coreservices: symbol FSFileSecuritySetMode not loaded")
	}
	return _fSFileSecuritySetMode(arg0, arg1)
}

var _fSFileSecuritySetOwner func(arg0 FSFileSecurityRef, arg1 uint32) int32

// FSFileSecuritySetOwner.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566946-fsfilesecuritysetowner
func FSFileSecuritySetOwner(arg0 FSFileSecurityRef, arg1 uint32) int32 {
	if _fSFileSecuritySetOwner == nil {
		panic("coreservices: symbol FSFileSecuritySetOwner not loaded")
	}
	return _fSFileSecuritySetOwner(arg0, arg1)
}

var _fSFileSecuritySetOwnerUUID func(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32

// FSFileSecuritySetOwnerUUID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565207-fsfilesecuritysetowneruuid
func FSFileSecuritySetOwnerUUID(arg0 FSFileSecurityRef, arg1 corefoundation.CFUUIDBytes) int32 {
	if _fSFileSecuritySetOwnerUUID == nil {
		panic("coreservices: symbol FSFileSecuritySetOwnerUUID not loaded")
	}
	return _fSFileSecuritySetOwnerUUID(arg0, arg1)
}

var _fSFindFolder func(arg0 uintptr, arg1 uint32, arg2 bool, arg3 uintptr) int16

// FSFindFolder.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389059-fsfindfolder
func FSFindFolder(arg0 uintptr, arg1 uint32, arg2 bool, arg3 uintptr) int16 {
	if _fSFindFolder == nil {
		panic("coreservices: symbol FSFindFolder not loaded")
	}
	return _fSFindFolder(arg0, arg1, arg2, arg3)
}

var _fSFlushFork func(arg0 FSIORefNum) int16

// FSFlushFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566414-fsflushfork
func FSFlushFork(arg0 FSIORefNum) int16 {
	if _fSFlushFork == nil {
		panic("coreservices: symbol FSFlushFork not loaded")
	}
	return _fSFlushFork(arg0)
}

var _fSFlushVolume func(arg0 uintptr) int32

// FSFlushVolume.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565506-fsflushvolume
func FSFlushVolume(arg0 uintptr) int32 {
	if _fSFlushVolume == nil {
		panic("coreservices: symbol FSFlushVolume not loaded")
	}
	return _fSFlushVolume(arg0)
}

var _fSGetAsyncEjectStatus func(arg0 FSVolumeOperation, arg1 FSEjectStatus, arg2 int32, arg3 uintptr, arg4 int32) int32

// FSGetAsyncEjectStatus.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565993-fsgetasyncejectstatus
func FSGetAsyncEjectStatus(arg0 FSVolumeOperation, arg1 FSEjectStatus, arg2 int32, arg3 uintptr, arg4 int32) int32 {
	if _fSGetAsyncEjectStatus == nil {
		panic("coreservices: symbol FSGetAsyncEjectStatus not loaded")
	}
	return _fSGetAsyncEjectStatus(arg0, arg1, arg2, arg3, arg4)
}

var _fSGetAsyncMountStatus func(arg0 FSVolumeOperation, arg1 FSMountStatus, arg2 int32, arg3 uintptr) int32

// FSGetAsyncMountStatus.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566519-fsgetasyncmountstatus
func FSGetAsyncMountStatus(arg0 FSVolumeOperation, arg1 FSMountStatus, arg2 int32, arg3 uintptr) int32 {
	if _fSGetAsyncMountStatus == nil {
		panic("coreservices: symbol FSGetAsyncMountStatus not loaded")
	}
	return _fSGetAsyncMountStatus(arg0, arg1, arg2, arg3)
}

var _fSGetAsyncUnmountStatus func(arg0 FSVolumeOperation, arg1 FSUnmountStatus, arg2 int32, arg3 uintptr, arg4 int32) int32

// FSGetAsyncUnmountStatus.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565808-fsgetasyncunmountstatus
func FSGetAsyncUnmountStatus(arg0 FSVolumeOperation, arg1 FSUnmountStatus, arg2 int32, arg3 uintptr, arg4 int32) int32 {
	if _fSGetAsyncUnmountStatus == nil {
		panic("coreservices: symbol FSGetAsyncUnmountStatus not loaded")
	}
	return _fSGetAsyncUnmountStatus(arg0, arg1, arg2, arg3, arg4)
}

var _fSGetDataForkName func(arg0 unsafe.Pointer) int16

// FSGetDataForkName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566760-fsgetdataforkname
func FSGetDataForkName(arg0 unsafe.Pointer) int16 {
	if _fSGetDataForkName == nil {
		panic("coreservices: symbol FSGetDataForkName not loaded")
	}
	return _fSGetDataForkName(arg0)
}

var _fSGetForkCBInfo func(arg0 FSIORefNum, arg1 uintptr, arg2 int16, arg3 FSIORefNum, arg4 FSForkInfo, arg5 uintptr, arg6 unsafe.Pointer) int16

// FSGetForkCBInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565345-fsgetforkcbinfo
func FSGetForkCBInfo(arg0 FSIORefNum, arg1 uintptr, arg2 int16, arg3 FSIORefNum, arg4 FSForkInfo, arg5 uintptr, arg6 unsafe.Pointer) int16 {
	if _fSGetForkCBInfo == nil {
		panic("coreservices: symbol FSGetForkCBInfo not loaded")
	}
	return _fSGetForkCBInfo(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _fSGetForkPosition func(arg0 FSIORefNum, arg1 int64) int16

// FSGetForkPosition.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565089-fsgetforkposition
func FSGetForkPosition(arg0 FSIORefNum, arg1 int64) int16 {
	if _fSGetForkPosition == nil {
		panic("coreservices: symbol FSGetForkPosition not loaded")
	}
	return _fSGetForkPosition(arg0, arg1)
}

var _fSGetForkSize func(arg0 FSIORefNum, arg1 int64) int16

// FSGetForkSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565455-fsgetforksize
func FSGetForkSize(arg0 FSIORefNum, arg1 int64) int16 {
	if _fSGetForkSize == nil {
		panic("coreservices: symbol FSGetForkSize not loaded")
	}
	return _fSGetForkSize(arg0, arg1)
}

var _fSGetHFSUniStrFromString func(arg0 corefoundation.CFStringRef, arg1 unsafe.Pointer) int32

// FSGetHFSUniStrFromString.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565877-fsgethfsunistrfromstring
func FSGetHFSUniStrFromString(arg0 corefoundation.CFStringRef, arg1 unsafe.Pointer) int32 {
	if _fSGetHFSUniStrFromString == nil {
		panic("coreservices: symbol FSGetHFSUniStrFromString not loaded")
	}
	return _fSGetHFSUniStrFromString(arg0, arg1)
}

var _fSGetResourceForkName func(arg0 unsafe.Pointer) int16

// FSGetResourceForkName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566158-fsgetresourceforkname
func FSGetResourceForkName(arg0 unsafe.Pointer) int16 {
	if _fSGetResourceForkName == nil {
		panic("coreservices: symbol FSGetResourceForkName not loaded")
	}
	return _fSGetResourceForkName(arg0)
}

var _fSGetTemporaryDirectoryForReplaceObject func(arg0 uintptr, arg1 uintptr, arg2 uintptr) int32

// FSGetTemporaryDirectoryForReplaceObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566436-fsgettemporarydirectoryforreplac
func FSGetTemporaryDirectoryForReplaceObject(arg0 uintptr, arg1 uintptr, arg2 uintptr) int32 {
	if _fSGetTemporaryDirectoryForReplaceObject == nil {
		panic("coreservices: symbol FSGetTemporaryDirectoryForReplaceObject not loaded")
	}
	return _fSGetTemporaryDirectoryForReplaceObject(arg0, arg1, arg2)
}

var _fSGetVolumeForDADisk func(arg0 diskarbitration.DADiskRef, arg1 uintptr) int32

// FSGetVolumeForDADisk.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566612-fsgetvolumefordadisk
func FSGetVolumeForDADisk(arg0 diskarbitration.DADiskRef, arg1 uintptr) int32 {
	if _fSGetVolumeForDADisk == nil {
		panic("coreservices: symbol FSGetVolumeForDADisk not loaded")
	}
	return _fSGetVolumeForDADisk(arg0, arg1)
}

var _fSGetVolumeForDiskID func(arg0 corefoundation.CFStringRef, arg1 uintptr) int32

// FSGetVolumeForDiskID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565529-fsgetvolumefordiskid
func FSGetVolumeForDiskID(arg0 corefoundation.CFStringRef, arg1 uintptr) int32 {
	if _fSGetVolumeForDiskID == nil {
		panic("coreservices: symbol FSGetVolumeForDiskID not loaded")
	}
	return _fSGetVolumeForDiskID(arg0, arg1)
}

var _fSGetVolumeInfo func(arg0 uintptr, arg1 uintptr, arg2 uintptr, arg3 FSVolumeInfoBitmap, arg4 FSVolumeInfo, arg5 unsafe.Pointer, arg6 uintptr) int16

// FSGetVolumeInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566350-fsgetvolumeinfo
func FSGetVolumeInfo(arg0 uintptr, arg1 uintptr, arg2 uintptr, arg3 FSVolumeInfoBitmap, arg4 FSVolumeInfo, arg5 unsafe.Pointer, arg6 uintptr) int16 {
	if _fSGetVolumeInfo == nil {
		panic("coreservices: symbol FSGetVolumeInfo not loaded")
	}
	return _fSGetVolumeInfo(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _fSGetVolumeMountInfo func(arg0 uintptr, arg1 unsafe.Pointer, arg2 uintptr, arg3 uintptr) int32

// FSGetVolumeMountInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565587-fsgetvolumemountinfo
func FSGetVolumeMountInfo(arg0 uintptr, arg1 unsafe.Pointer, arg2 uintptr, arg3 uintptr) int32 {
	if _fSGetVolumeMountInfo == nil {
		panic("coreservices: symbol FSGetVolumeMountInfo not loaded")
	}
	return _fSGetVolumeMountInfo(arg0, arg1, arg2, arg3)
}

var _fSGetVolumeMountInfoSize func(arg0 uintptr, arg1 uintptr) int32

// FSGetVolumeMountInfoSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566473-fsgetvolumemountinfosize
func FSGetVolumeMountInfoSize(arg0 uintptr, arg1 uintptr) int32 {
	if _fSGetVolumeMountInfoSize == nil {
		panic("coreservices: symbol FSGetVolumeMountInfoSize not loaded")
	}
	return _fSGetVolumeMountInfoSize(arg0, arg1)
}

var _fSGetVolumeParms func(arg0 uintptr, arg1 GetVolParmsInfoBuffer, arg2 uintptr) int32

// FSGetVolumeParms.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565147-fsgetvolumeparms
func FSGetVolumeParms(arg0 uintptr, arg1 GetVolParmsInfoBuffer, arg2 uintptr) int32 {
	if _fSGetVolumeParms == nil {
		panic("coreservices: symbol FSGetVolumeParms not loaded")
	}
	return _fSGetVolumeParms(arg0, arg1, arg2)
}

var _fSIsFSRefValid func(arg0 uintptr) bool

// FSIsFSRefValid.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565952-fsisfsrefvalid
func FSIsFSRefValid(arg0 uintptr) bool {
	if _fSIsFSRefValid == nil {
		panic("coreservices: symbol FSIsFSRefValid not loaded")
	}
	return _fSIsFSRefValid(arg0)
}

var _fSIterateForks func(arg0 uintptr, arg1 CatPositionRec, arg2 unsafe.Pointer, arg3 int64, arg4 uint64) int16

// FSIterateForks.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565757-fsiterateforks
func FSIterateForks(arg0 uintptr, arg1 CatPositionRec, arg2 unsafe.Pointer, arg3 int64, arg4 uint64) int16 {
	if _fSIterateForks == nil {
		panic("coreservices: symbol FSIterateForks not loaded")
	}
	return _fSIterateForks(arg0, arg1, arg2, arg3, arg4)
}

var _fSLockRange func(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uint64, arg4 uint64) int32

// FSLockRange.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566371-fslockrange
func FSLockRange(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uint64, arg4 uint64) int32 {
	if _fSLockRange == nil {
		panic("coreservices: symbol FSLockRange not loaded")
	}
	return _fSLockRange(arg0, arg1, arg2, arg3, arg4)
}

var _fSMakeFSRefUnicode func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 TextEncoding, arg4 uintptr) int16

// FSMakeFSRefUnicode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565210-fsmakefsrefunicode
func FSMakeFSRefUnicode(arg0 uintptr, arg1 uint, arg2 uint16, arg3 TextEncoding, arg4 uintptr) int16 {
	if _fSMakeFSRefUnicode == nil {
		panic("coreservices: symbol FSMakeFSRefUnicode not loaded")
	}
	return _fSMakeFSRefUnicode(arg0, arg1, arg2, arg3, arg4)
}

var _fSMountLocalVolumeAsync func(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFURLRef, arg2 FSVolumeOperation, arg3 uintptr, arg4 FSVolumeMountUPP, arg5 corefoundation.CFRunLoopRef, arg6 corefoundation.CFStringRef) int32

// FSMountLocalVolumeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565530-fsmountlocalvolumeasync
func FSMountLocalVolumeAsync(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFURLRef, arg2 FSVolumeOperation, arg3 uintptr, arg4 FSVolumeMountUPP, arg5 corefoundation.CFRunLoopRef, arg6 corefoundation.CFStringRef) int32 {
	if _fSMountLocalVolumeAsync == nil {
		panic("coreservices: symbol FSMountLocalVolumeAsync not loaded")
	}
	return _fSMountLocalVolumeAsync(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _fSMountLocalVolumeSync func(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFURLRef, arg2 uintptr, arg3 uintptr) int32

// FSMountLocalVolumeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566302-fsmountlocalvolumesync
func FSMountLocalVolumeSync(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFURLRef, arg2 uintptr, arg3 uintptr) int32 {
	if _fSMountLocalVolumeSync == nil {
		panic("coreservices: symbol FSMountLocalVolumeSync not loaded")
	}
	return _fSMountLocalVolumeSync(arg0, arg1, arg2, arg3)
}

var _fSMountServerVolumeAsync func(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 FSVolumeOperation, arg5 uintptr, arg6 FSVolumeMountUPP, arg7 corefoundation.CFRunLoopRef, arg8 corefoundation.CFStringRef) int32

// FSMountServerVolumeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565318-fsmountservervolumeasync
func FSMountServerVolumeAsync(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 FSVolumeOperation, arg5 uintptr, arg6 FSVolumeMountUPP, arg7 corefoundation.CFRunLoopRef, arg8 corefoundation.CFStringRef) int32 {
	if _fSMountServerVolumeAsync == nil {
		panic("coreservices: symbol FSMountServerVolumeAsync not loaded")
	}
	return _fSMountServerVolumeAsync(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _fSMountServerVolumeSync func(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 uintptr) int32

// FSMountServerVolumeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565166-fsmountservervolumesync
func FSMountServerVolumeSync(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 uintptr) int32 {
	if _fSMountServerVolumeSync == nil {
		panic("coreservices: symbol FSMountServerVolumeSync not loaded")
	}
	return _fSMountServerVolumeSync(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSMoveObject func(arg0 uintptr, arg1 uintptr, arg2 uintptr) int16

// FSMoveObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566291-fsmoveobject
func FSMoveObject(arg0 uintptr, arg1 uintptr, arg2 uintptr) int16 {
	if _fSMoveObject == nil {
		panic("coreservices: symbol FSMoveObject not loaded")
	}
	return _fSMoveObject(arg0, arg1, arg2)
}

var _fSMoveObjectAsync func(arg0 FSFileOperationRef, arg1 uintptr, arg2 uintptr, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32

// FSMoveObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566762-fsmoveobjectasync
func FSMoveObjectAsync(arg0 FSFileOperationRef, arg1 uintptr, arg2 uintptr, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32 {
	if _fSMoveObjectAsync == nil {
		panic("coreservices: symbol FSMoveObjectAsync not loaded")
	}
	return _fSMoveObjectAsync(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _fSMoveObjectSync func(arg0 uintptr, arg1 uintptr, arg2 corefoundation.CFStringRef, arg3 uintptr, arg4 uintptr) int32

// FSMoveObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566525-fsmoveobjectsync
func FSMoveObjectSync(arg0 uintptr, arg1 uintptr, arg2 corefoundation.CFStringRef, arg3 uintptr, arg4 uintptr) int32 {
	if _fSMoveObjectSync == nil {
		panic("coreservices: symbol FSMoveObjectSync not loaded")
	}
	return _fSMoveObjectSync(arg0, arg1, arg2, arg3, arg4)
}

var _fSMoveObjectToTrashAsync func(arg0 FSFileOperationRef, arg1 uintptr, arg2 uintptr, arg3 unsafe.Pointer, arg4 float64, arg5 FSFileOperationClientContext) int32

// FSMoveObjectToTrashAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565854-fsmoveobjecttotrashasync
func FSMoveObjectToTrashAsync(arg0 FSFileOperationRef, arg1 uintptr, arg2 uintptr, arg3 unsafe.Pointer, arg4 float64, arg5 FSFileOperationClientContext) int32 {
	if _fSMoveObjectToTrashAsync == nil {
		panic("coreservices: symbol FSMoveObjectToTrashAsync not loaded")
	}
	return _fSMoveObjectToTrashAsync(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSMoveObjectToTrashSync func(arg0 uintptr, arg1 uintptr, arg2 uintptr) int32

// FSMoveObjectToTrashSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566651-fsmoveobjecttotrashsync
func FSMoveObjectToTrashSync(arg0 uintptr, arg1 uintptr, arg2 uintptr) int32 {
	if _fSMoveObjectToTrashSync == nil {
		panic("coreservices: symbol FSMoveObjectToTrashSync not loaded")
	}
	return _fSMoveObjectToTrashSync(arg0, arg1, arg2)
}

var _fSOpenFork func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 int8, arg4 FSIORefNum) int16

// FSOpenFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565689-fsopenfork
func FSOpenFork(arg0 uintptr, arg1 uint, arg2 uint16, arg3 int8, arg4 FSIORefNum) int16 {
	if _fSOpenFork == nil {
		panic("coreservices: symbol FSOpenFork not loaded")
	}
	return _fSOpenFork(arg0, arg1, arg2, arg3, arg4)
}

var _fSOpenIterator func(arg0 uintptr, arg1 FSIteratorFlags, arg2 FSIterator) int16

// FSOpenIterator.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565368-fsopeniterator
func FSOpenIterator(arg0 uintptr, arg1 FSIteratorFlags, arg2 FSIterator) int16 {
	if _fSOpenIterator == nil {
		panic("coreservices: symbol FSOpenIterator not loaded")
	}
	return _fSOpenIterator(arg0, arg1, arg2)
}

var _fSOpenOrphanResFile func(arg0 uintptr, arg1 unsafe.Pointer, arg2 ResFileRefNum) int16

// FSOpenOrphanResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529349-fsopenorphanresfile
func FSOpenOrphanResFile(arg0 uintptr, arg1 unsafe.Pointer, arg2 ResFileRefNum) int16 {
	if _fSOpenOrphanResFile == nil {
		panic("coreservices: symbol FSOpenOrphanResFile not loaded")
	}
	return _fSOpenOrphanResFile(arg0, arg1, arg2)
}

var _fSOpenResFile func(arg0 uintptr, arg1 int8) ResFileRefNum

// FSOpenResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529232-fsopenresfile
func FSOpenResFile(arg0 uintptr, arg1 int8) ResFileRefNum {
	if _fSOpenResFile == nil {
		panic("coreservices: symbol FSOpenResFile not loaded")
	}
	return _fSOpenResFile(arg0, arg1)
}

var _fSOpenResourceFile func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 int8, arg4 ResFileRefNum) int16

// FSOpenResourceFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529373-fsopenresourcefile
func FSOpenResourceFile(arg0 uintptr, arg1 uint, arg2 uint16, arg3 int8, arg4 ResFileRefNum) int16 {
	if _fSOpenResourceFile == nil {
		panic("coreservices: symbol FSOpenResourceFile not loaded")
	}
	return _fSOpenResourceFile(arg0, arg1, arg2, arg3, arg4)
}

var _fSPathCopyObjectAsync func(arg0 FSFileOperationRef, arg1 int8, arg2 int8, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32

// FSPathCopyObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566056-fspathcopyobjectasync
func FSPathCopyObjectAsync(arg0 FSFileOperationRef, arg1 int8, arg2 int8, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32 {
	if _fSPathCopyObjectAsync == nil {
		panic("coreservices: symbol FSPathCopyObjectAsync not loaded")
	}
	return _fSPathCopyObjectAsync(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _fSPathCopyObjectSync func(arg0 int8, arg1 int8, arg2 corefoundation.CFStringRef, arg3 int8, arg4 uintptr) int32

// FSPathCopyObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565384-fspathcopyobjectsync
func FSPathCopyObjectSync(arg0 int8, arg1 int8, arg2 corefoundation.CFStringRef, arg3 int8, arg4 uintptr) int32 {
	if _fSPathCopyObjectSync == nil {
		panic("coreservices: symbol FSPathCopyObjectSync not loaded")
	}
	return _fSPathCopyObjectSync(arg0, arg1, arg2, arg3, arg4)
}

var _fSPathFileOperationCopyStatus func(arg0 FSFileOperationRef, arg1 int8, arg2 FSFileOperationStage, arg3 int32, arg4 corefoundation.CFDictionaryRef) int32

// FSPathFileOperationCopyStatus.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566335-fspathfileoperationcopystatus
func FSPathFileOperationCopyStatus(arg0 FSFileOperationRef, arg1 int8, arg2 FSFileOperationStage, arg3 int32, arg4 corefoundation.CFDictionaryRef) int32 {
	if _fSPathFileOperationCopyStatus == nil {
		panic("coreservices: symbol FSPathFileOperationCopyStatus not loaded")
	}
	return _fSPathFileOperationCopyStatus(arg0, arg1, arg2, arg3, arg4)
}

var _fSPathGetTemporaryDirectoryForReplaceObject func(arg0 int8, arg1 int8, arg2 uint32, arg3 uintptr) int32

// FSPathGetTemporaryDirectoryForReplaceObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566163-fspathgettemporarydirectoryforre
func FSPathGetTemporaryDirectoryForReplaceObject(arg0 int8, arg1 int8, arg2 uint32, arg3 uintptr) int32 {
	if _fSPathGetTemporaryDirectoryForReplaceObject == nil {
		panic("coreservices: symbol FSPathGetTemporaryDirectoryForReplaceObject not loaded")
	}
	return _fSPathGetTemporaryDirectoryForReplaceObject(arg0, arg1, arg2, arg3)
}

var _fSPathMakeRef func(arg0 uint8, arg1 uintptr, arg2 bool) int32

// FSPathMakeRef.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565195-fspathmakeref
func FSPathMakeRef(arg0 uint8, arg1 uintptr, arg2 bool) int32 {
	if _fSPathMakeRef == nil {
		panic("coreservices: symbol FSPathMakeRef not loaded")
	}
	return _fSPathMakeRef(arg0, arg1, arg2)
}

var _fSPathMakeRefWithOptions func(arg0 uint8, arg1 uintptr, arg2 uintptr, arg3 bool) int32

// FSPathMakeRefWithOptions.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566339-fspathmakerefwithoptions
func FSPathMakeRefWithOptions(arg0 uint8, arg1 uintptr, arg2 uintptr, arg3 bool) int32 {
	if _fSPathMakeRefWithOptions == nil {
		panic("coreservices: symbol FSPathMakeRefWithOptions not loaded")
	}
	return _fSPathMakeRefWithOptions(arg0, arg1, arg2, arg3)
}

var _fSPathMoveObjectAsync func(arg0 FSFileOperationRef, arg1 int8, arg2 int8, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32

// FSPathMoveObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566169-fspathmoveobjectasync
func FSPathMoveObjectAsync(arg0 FSFileOperationRef, arg1 int8, arg2 int8, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 unsafe.Pointer, arg6 float64, arg7 FSFileOperationClientContext) int32 {
	if _fSPathMoveObjectAsync == nil {
		panic("coreservices: symbol FSPathMoveObjectAsync not loaded")
	}
	return _fSPathMoveObjectAsync(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _fSPathMoveObjectSync func(arg0 int8, arg1 int8, arg2 corefoundation.CFStringRef, arg3 int8, arg4 uintptr) int32

// FSPathMoveObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566022-fspathmoveobjectsync
func FSPathMoveObjectSync(arg0 int8, arg1 int8, arg2 corefoundation.CFStringRef, arg3 int8, arg4 uintptr) int32 {
	if _fSPathMoveObjectSync == nil {
		panic("coreservices: symbol FSPathMoveObjectSync not loaded")
	}
	return _fSPathMoveObjectSync(arg0, arg1, arg2, arg3, arg4)
}

var _fSPathMoveObjectToTrashAsync func(arg0 FSFileOperationRef, arg1 int8, arg2 uintptr, arg3 unsafe.Pointer, arg4 float64, arg5 FSFileOperationClientContext) int32

// FSPathMoveObjectToTrashAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565270-fspathmoveobjecttotrashasync
func FSPathMoveObjectToTrashAsync(arg0 FSFileOperationRef, arg1 int8, arg2 uintptr, arg3 unsafe.Pointer, arg4 float64, arg5 FSFileOperationClientContext) int32 {
	if _fSPathMoveObjectToTrashAsync == nil {
		panic("coreservices: symbol FSPathMoveObjectToTrashAsync not loaded")
	}
	return _fSPathMoveObjectToTrashAsync(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSPathMoveObjectToTrashSync func(arg0 int8, arg1 int8, arg2 uintptr) int32

// FSPathMoveObjectToTrashSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566818-fspathmoveobjecttotrashsync
func FSPathMoveObjectToTrashSync(arg0 int8, arg1 int8, arg2 uintptr) int32 {
	if _fSPathMoveObjectToTrashSync == nil {
		panic("coreservices: symbol FSPathMoveObjectToTrashSync not loaded")
	}
	return _fSPathMoveObjectToTrashSync(arg0, arg1, arg2)
}

var _fSPathReplaceObject func(arg0 int8, arg1 int8, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 int8, arg5 uintptr) int32

// FSPathReplaceObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566463-fspathreplaceobject
func FSPathReplaceObject(arg0 int8, arg1 int8, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 int8, arg5 uintptr) int32 {
	if _fSPathReplaceObject == nil {
		panic("coreservices: symbol FSPathReplaceObject not loaded")
	}
	return _fSPathReplaceObject(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSReadFork func(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uintptr, arg4 uintptr) int16

// FSReadFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565997-fsreadfork
func FSReadFork(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uintptr, arg4 uintptr) int16 {
	if _fSReadFork == nil {
		panic("coreservices: symbol FSReadFork not loaded")
	}
	return _fSReadFork(arg0, arg1, arg2, arg3, arg4)
}

var _fSRefMakePath func(arg0 uintptr, arg1 uint8, arg2 uint32) int32

// FSRefMakePath.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565635-fsrefmakepath
func FSRefMakePath(arg0 uintptr, arg1 uint8, arg2 uint32) int32 {
	if _fSRefMakePath == nil {
		panic("coreservices: symbol FSRefMakePath not loaded")
	}
	return _fSRefMakePath(arg0, arg1, arg2)
}

var _fSRenameUnicode func(arg0 uintptr, arg1 uint, arg2 uint16, arg3 TextEncoding, arg4 uintptr) int16

// FSRenameUnicode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565792-fsrenameunicode
func FSRenameUnicode(arg0 uintptr, arg1 uint, arg2 uint16, arg3 TextEncoding, arg4 uintptr) int16 {
	if _fSRenameUnicode == nil {
		panic("coreservices: symbol FSRenameUnicode not loaded")
	}
	return _fSRenameUnicode(arg0, arg1, arg2, arg3, arg4)
}

var _fSReplaceObject func(arg0 uintptr, arg1 uintptr, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 uintptr, arg6 uintptr) int32

// FSReplaceObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566736-fsreplaceobject
func FSReplaceObject(arg0 uintptr, arg1 uintptr, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 uintptr, arg5 uintptr, arg6 uintptr) int32 {
	if _fSReplaceObject == nil {
		panic("coreservices: symbol FSReplaceObject not loaded")
	}
	return _fSReplaceObject(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _fSResolveNodeID func(arg0 uintptr, arg1 uint32, arg2 FSRefPtr) int32

// FSResolveNodeID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565623-fsresolvenodeid
func FSResolveNodeID(arg0 uintptr, arg1 uint32, arg2 FSRefPtr) int32 {
	if _fSResolveNodeID == nil {
		panic("coreservices: symbol FSResolveNodeID not loaded")
	}
	return _fSResolveNodeID(arg0, arg1, arg2)
}

var _fSResourceFileAlreadyOpen func(arg0 uintptr, arg1 bool, arg2 ResFileRefNum) bool

// FSResourceFileAlreadyOpen.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529331-fsresourcefilealreadyopen
func FSResourceFileAlreadyOpen(arg0 uintptr, arg1 bool, arg2 ResFileRefNum) bool {
	if _fSResourceFileAlreadyOpen == nil {
		panic("coreservices: symbol FSResourceFileAlreadyOpen not loaded")
	}
	return _fSResourceFileAlreadyOpen(arg0, arg1, arg2)
}

var _fSSetCatalogInfo func(arg0 uintptr, arg1 FSCatalogInfoBitmap, arg2 FSCatalogInfo) int16

// FSSetCatalogInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566580-fssetcataloginfo
func FSSetCatalogInfo(arg0 uintptr, arg1 FSCatalogInfoBitmap, arg2 FSCatalogInfo) int16 {
	if _fSSetCatalogInfo == nil {
		panic("coreservices: symbol FSSetCatalogInfo not loaded")
	}
	return _fSSetCatalogInfo(arg0, arg1, arg2)
}

var _fSSetForkPosition func(arg0 FSIORefNum, arg1 uint16, arg2 int64) int16

// FSSetForkPosition.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566508-fssetforkposition
func FSSetForkPosition(arg0 FSIORefNum, arg1 uint16, arg2 int64) int16 {
	if _fSSetForkPosition == nil {
		panic("coreservices: symbol FSSetForkPosition not loaded")
	}
	return _fSSetForkPosition(arg0, arg1, arg2)
}

var _fSSetForkSize func(arg0 FSIORefNum, arg1 uint16, arg2 int64) int16

// FSSetForkSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566185-fssetforksize
func FSSetForkSize(arg0 FSIORefNum, arg1 uint16, arg2 int64) int16 {
	if _fSSetForkSize == nil {
		panic("coreservices: symbol FSSetForkSize not loaded")
	}
	return _fSSetForkSize(arg0, arg1, arg2)
}

var _fSSetVolumeInfo func(arg0 uintptr, arg1 FSVolumeInfoBitmap, arg2 FSVolumeInfo) int16

// FSSetVolumeInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565574-fssetvolumeinfo
func FSSetVolumeInfo(arg0 uintptr, arg1 FSVolumeInfoBitmap, arg2 FSVolumeInfo) int16 {
	if _fSSetVolumeInfo == nil {
		panic("coreservices: symbol FSSetVolumeInfo not loaded")
	}
	return _fSSetVolumeInfo(arg0, arg1, arg2)
}

var _fSUnlinkObject func(arg0 uintptr) int16

// FSUnlinkObject.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565924-fsunlinkobject
func FSUnlinkObject(arg0 uintptr) int16 {
	if _fSUnlinkObject == nil {
		panic("coreservices: symbol FSUnlinkObject not loaded")
	}
	return _fSUnlinkObject(arg0)
}

var _fSUnlockRange func(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uint64, arg4 uint64) int32

// FSUnlockRange.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565302-fsunlockrange
func FSUnlockRange(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uint64, arg4 uint64) int32 {
	if _fSUnlockRange == nil {
		panic("coreservices: symbol FSUnlockRange not loaded")
	}
	return _fSUnlockRange(arg0, arg1, arg2, arg3, arg4)
}

var _fSUnmountVolumeAsync func(arg0 uintptr, arg1 uintptr, arg2 FSVolumeOperation, arg3 FSVolumeUnmountUPP, arg4 corefoundation.CFRunLoopRef, arg5 corefoundation.CFStringRef) int32

// FSUnmountVolumeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566587-fsunmountvolumeasync
func FSUnmountVolumeAsync(arg0 uintptr, arg1 uintptr, arg2 FSVolumeOperation, arg3 FSVolumeUnmountUPP, arg4 corefoundation.CFRunLoopRef, arg5 corefoundation.CFStringRef) int32 {
	if _fSUnmountVolumeAsync == nil {
		panic("coreservices: symbol FSUnmountVolumeAsync not loaded")
	}
	return _fSUnmountVolumeAsync(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _fSUnmountVolumeSync func(arg0 uintptr, arg1 uintptr, arg2 int32) int32

// FSUnmountVolumeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566764-fsunmountvolumesync
func FSUnmountVolumeSync(arg0 uintptr, arg1 uintptr, arg2 int32) int32 {
	if _fSUnmountVolumeSync == nil {
		panic("coreservices: symbol FSUnmountVolumeSync not loaded")
	}
	return _fSUnmountVolumeSync(arg0, arg1, arg2)
}

var _fSVolumeMount func(arg0 unsafe.Pointer, arg1 uintptr) int32

// FSVolumeMount.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566510-fsvolumemount
func FSVolumeMount(arg0 unsafe.Pointer, arg1 uintptr) int32 {
	if _fSVolumeMount == nil {
		panic("coreservices: symbol FSVolumeMount not loaded")
	}
	return _fSVolumeMount(arg0, arg1)
}

var _fSWriteFork func(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uintptr, arg4 uintptr) int16

// FSWriteFork.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565526-fswritefork
func FSWriteFork(arg0 FSIORefNum, arg1 uint16, arg2 int64, arg3 uintptr, arg4 uintptr) int16 {
	if _fSWriteFork == nil {
		panic("coreservices: symbol FSWriteFork not loaded")
	}
	return _fSWriteFork(arg0, arg1, arg2, arg3, arg4)
}

var _findFolder func(arg0 uintptr, arg1 uint32, arg2 bool, arg3 uintptr, arg4 int32) int16

// FindFolder.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389175-findfolder
func FindFolder(arg0 uintptr, arg1 uint32, arg2 bool, arg3 uintptr, arg4 int32) int16 {
	if _findFolder == nil {
		panic("coreservices: symbol FindFolder not loaded")
	}
	return _findFolder(arg0, arg1, arg2, arg3, arg4)
}

var _findNextComponent func(arg0 Component, arg1 ComponentDescription) Component

// FindNextComponent returns the component identifier for the next registered component that meets the selection criteria specified by your application.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516552-findnextcomponent
func FindNextComponent(arg0 Component, arg1 ComponentDescription) Component {
	if _findNextComponent == nil {
		panic("coreservices: symbol FindNextComponent not loaded")
	}
	return _findNextComponent(arg0, arg1)
}

var _fix2Frac func(arg0 uintptr) unsafe.Pointer

// Fix2Frac.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409219-fix2frac
func Fix2Frac(arg0 uintptr) unsafe.Pointer {
	if _fix2Frac == nil {
		panic("coreservices: symbol Fix2Frac not loaded")
	}
	return _fix2Frac(arg0)
}

var _fix2Long func(arg0 uintptr) int32

// Fix2Long.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409214-fix2long
func Fix2Long(arg0 uintptr) int32 {
	if _fix2Long == nil {
		panic("coreservices: symbol Fix2Long not loaded")
	}
	return _fix2Long(arg0)
}

var _fix2X func(arg0 uintptr) float64

// Fix2X.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409225-fix2x
func Fix2X(arg0 uintptr) float64 {
	if _fix2X == nil {
		panic("coreservices: symbol Fix2X not loaded")
	}
	return _fix2X(arg0)
}

var _fixATan2 func(arg0 int32, arg1 int32) objectivec.IObject

// FixATan2.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409269-fixatan2
func FixATan2(arg0 int32, arg1 int32) objectivec.IObject {
	if _fixATan2 == nil {
		panic("coreservices: symbol FixATan2 not loaded")
	}
	return _fixATan2(arg0, arg1)
}

var _fixDiv func(arg0 uintptr, arg1 uintptr) objectivec.IObject

// FixDiv.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409285-fixdiv
func FixDiv(arg0 uintptr, arg1 uintptr) objectivec.IObject {
	if _fixDiv == nil {
		panic("coreservices: symbol FixDiv not loaded")
	}
	return _fixDiv(arg0, arg1)
}

var _fixMul func(arg0 uintptr, arg1 uintptr) objectivec.IObject

// FixMul.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409241-fixmul
func FixMul(arg0 uintptr, arg1 uintptr) objectivec.IObject {
	if _fixMul == nil {
		panic("coreservices: symbol FixMul not loaded")
	}
	return _fixMul(arg0, arg1)
}

var _fixRatio func(arg0 int16, arg1 int16) objectivec.IObject

// FixRatio.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409283-fixratio
func FixRatio(arg0 int16, arg1 int16) objectivec.IObject {
	if _fixRatio == nil {
		panic("coreservices: symbol FixRatio not loaded")
	}
	return _fixRatio(arg0, arg1)
}

var _fixRound func(arg0 uintptr) int16

// FixRound.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409274-fixround
func FixRound(arg0 uintptr) int16 {
	if _fixRound == nil {
		panic("coreservices: symbol FixRound not loaded")
	}
	return _fixRound(arg0)
}

var _flattenCollection func(arg0 Collection, arg1 CollectionFlattenUPP) int16

// FlattenCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551324-flattencollection
func FlattenCollection(arg0 Collection, arg1 CollectionFlattenUPP) int16 {
	if _flattenCollection == nil {
		panic("coreservices: symbol FlattenCollection not loaded")
	}
	return _flattenCollection(arg0, arg1)
}

var _flattenCollectionToHdl func(arg0 Collection, arg1 uintptr) int16

// FlattenCollectionToHdl.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551407-flattencollectiontohdl
func FlattenCollectionToHdl(arg0 Collection, arg1 uintptr) int16 {
	if _flattenCollectionToHdl == nil {
		panic("coreservices: symbol FlattenCollectionToHdl not loaded")
	}
	return _flattenCollectionToHdl(arg0, arg1)
}

var _flattenPartialCollection func(arg0 Collection, arg1 CollectionFlattenUPP, arg2 int32, arg3 int32) int16

// FlattenPartialCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551435-flattenpartialcollection
func FlattenPartialCollection(arg0 Collection, arg1 CollectionFlattenUPP, arg2 int32, arg3 int32) int16 {
	if _flattenPartialCollection == nil {
		panic("coreservices: symbol FlattenPartialCollection not loaded")
	}
	return _flattenPartialCollection(arg0, arg1, arg2, arg3)
}

var _frac2Fix func(arg0 unsafe.Pointer) objectivec.IObject

// Frac2Fix.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409286-frac2fix
func Frac2Fix(arg0 unsafe.Pointer) objectivec.IObject {
	if _frac2Fix == nil {
		panic("coreservices: symbol Frac2Fix not loaded")
	}
	return _frac2Fix(arg0)
}

var _frac2X func(arg0 unsafe.Pointer) float64

// Frac2X.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409216-frac2x
func Frac2X(arg0 unsafe.Pointer) float64 {
	if _frac2X == nil {
		panic("coreservices: symbol Frac2X not loaded")
	}
	return _frac2X(arg0)
}

var _fracCos func(arg0 uintptr) unsafe.Pointer

// FracCos.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409257-fraccos
func FracCos(arg0 uintptr) unsafe.Pointer {
	if _fracCos == nil {
		panic("coreservices: symbol FracCos not loaded")
	}
	return _fracCos(arg0)
}

var _fracDiv func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// FracDiv.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409280-fracdiv
func FracDiv(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _fracDiv == nil {
		panic("coreservices: symbol FracDiv not loaded")
	}
	return _fracDiv(arg0, arg1)
}

var _fracMul func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// FracMul.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409259-fracmul
func FracMul(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _fracMul == nil {
		panic("coreservices: symbol FracMul not loaded")
	}
	return _fracMul(arg0, arg1)
}

var _fracSin func(arg0 uintptr) unsafe.Pointer

// FracSin.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409243-fracsin
func FracSin(arg0 uintptr) unsafe.Pointer {
	if _fracSin == nil {
		panic("coreservices: symbol FracSin not loaded")
	}
	return _fracSin(arg0)
}

var _fracSqrt func(arg0 unsafe.Pointer) unsafe.Pointer

// FracSqrt.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409284-fracsqrt
func FracSqrt(arg0 unsafe.Pointer) unsafe.Pointer {
	if _fracSqrt == nil {
		panic("coreservices: symbol FracSqrt not loaded")
	}
	return _fracSqrt(arg0)
}

var _gestaltFunc func(arg0 uint32, arg1 int32) int16

// GestaltFunc obtains information about the operating environment.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1471624-gestalt
func GestaltFunc(arg0 uint32, arg1 int32) int16 {
	if _gestaltFunc == nil {
		panic("coreservices: symbol Gestalt not loaded")
	}
	return _gestaltFunc(arg0, arg1)
}

var _get1IndResource func(arg0 uint32, arg1 ResourceIndex) objectivec.IObject

// Get1IndResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529284-get1indresource
func Get1IndResource(arg0 uint32, arg1 ResourceIndex) objectivec.IObject {
	if _get1IndResource == nil {
		panic("coreservices: symbol Get1IndResource not loaded")
	}
	return _get1IndResource(arg0, arg1)
}

var _get1IndType func(arg0 uint32, arg1 ResourceIndex)

// Get1IndType.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529368-get1indtype
func Get1IndType(arg0 uint32, arg1 ResourceIndex) {
	if _get1IndType == nil {
		panic("coreservices: symbol Get1IndType not loaded")
	}
	_get1IndType(arg0, arg1)
}

var _get1NamedResource func(arg0 uint32, arg1 unsafe.Pointer) objectivec.IObject

// Get1NamedResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529253-get1namedresource
func Get1NamedResource(arg0 uint32, arg1 unsafe.Pointer) objectivec.IObject {
	if _get1NamedResource == nil {
		panic("coreservices: symbol Get1NamedResource not loaded")
	}
	return _get1NamedResource(arg0, arg1)
}

var _get1Resource func(arg0 uint32, arg1 ResID) objectivec.IObject

// Get1Resource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529324-get1resource
func Get1Resource(arg0 uint32, arg1 ResID) objectivec.IObject {
	if _get1Resource == nil {
		panic("coreservices: symbol Get1Resource not loaded")
	}
	return _get1Resource(arg0, arg1)
}

var _getCPUSpeed func() int

// GetCPUSpeed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427401-getcpuspeed
func GetCPUSpeed() int {
	if _getCPUSpeed == nil {
		panic("coreservices: symbol GetCPUSpeed not loaded")
	}
	return _getCPUSpeed()
}

var _getCollectionDefaultAttributes func(arg0 Collection) int32

// GetCollectionDefaultAttributes.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551445-getcollectiondefaultattributes
func GetCollectionDefaultAttributes(arg0 Collection) int32 {
	if _getCollectionDefaultAttributes == nil {
		panic("coreservices: symbol GetCollectionDefaultAttributes not loaded")
	}
	return _getCollectionDefaultAttributes(arg0)
}

var _getCollectionExceptionProc func(arg0 Collection) CollectionExceptionUPP

// GetCollectionExceptionProc.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551420-getcollectionexceptionproc
func GetCollectionExceptionProc(arg0 Collection) CollectionExceptionUPP {
	if _getCollectionExceptionProc == nil {
		panic("coreservices: symbol GetCollectionExceptionProc not loaded")
	}
	return _getCollectionExceptionProc(arg0)
}

var _getCollectionItem func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32) int16

// GetCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551357-getcollectionitem
func GetCollectionItem(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32) int16 {
	if _getCollectionItem == nil {
		panic("coreservices: symbol GetCollectionItem not loaded")
	}
	return _getCollectionItem(arg0, arg1, arg2, arg3)
}

var _getCollectionItemHdl func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 uintptr) int16

// GetCollectionItemHdl.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551423-getcollectionitemhdl
func GetCollectionItemHdl(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 uintptr) int16 {
	if _getCollectionItemHdl == nil {
		panic("coreservices: symbol GetCollectionItemHdl not loaded")
	}
	return _getCollectionItemHdl(arg0, arg1, arg2, arg3)
}

var _getCollectionItemInfo func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32, arg4 int32, arg5 int32) int16

// GetCollectionItemInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551399-getcollectioniteminfo
func GetCollectionItemInfo(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32, arg4 int32, arg5 int32) int16 {
	if _getCollectionItemInfo == nil {
		panic("coreservices: symbol GetCollectionItemInfo not loaded")
	}
	return _getCollectionItemInfo(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _getCollectionRetainCount func(arg0 Collection) objectivec.IObject

// GetCollectionRetainCount.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551403-getcollectionretaincount
func GetCollectionRetainCount(arg0 Collection) objectivec.IObject {
	if _getCollectionRetainCount == nil {
		panic("coreservices: symbol GetCollectionRetainCount not loaded")
	}
	return _getCollectionRetainCount(arg0)
}

var _getComponentIndString func(arg0 Component, arg1 unsafe.Pointer, arg2 int16, arg3 int16) int16

// GetComponentIndString.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516616-getcomponentindstring
func GetComponentIndString(arg0 Component, arg1 unsafe.Pointer, arg2 int16, arg3 int16) int16 {
	if _getComponentIndString == nil {
		panic("coreservices: symbol GetComponentIndString not loaded")
	}
	return _getComponentIndString(arg0, arg1, arg2, arg3)
}

var _getComponentInfo func(arg0 Component, arg1 ComponentDescription, arg2 uintptr, arg3 uintptr, arg4 uintptr) int16

// GetComponentInfo returns to your application the registration information for a component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516438-getcomponentinfo
func GetComponentInfo(arg0 Component, arg1 ComponentDescription, arg2 uintptr, arg3 uintptr, arg4 uintptr) int16 {
	if _getComponentInfo == nil {
		panic("coreservices: symbol GetComponentInfo not loaded")
	}
	return _getComponentInfo(arg0, arg1, arg2, arg3, arg4)
}

var _getComponentInstanceError func(arg0 ComponentInstance) int16

// GetComponentInstanceError returns to your application the last error generated by a specific connection to a component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516618-getcomponentinstanceerror
func GetComponentInstanceError(arg0 ComponentInstance) int16 {
	if _getComponentInstanceError == nil {
		panic("coreservices: symbol GetComponentInstanceError not loaded")
	}
	return _getComponentInstanceError(arg0)
}

var _getComponentInstanceStorage func(arg0 ComponentInstance) objectivec.IObject

// GetComponentInstanceStorage allows your component to retrieve a handle to the memory associated with a connection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516517-getcomponentinstancestorage
func GetComponentInstanceStorage(arg0 ComponentInstance) objectivec.IObject {
	if _getComponentInstanceStorage == nil {
		panic("coreservices: symbol GetComponentInstanceStorage not loaded")
	}
	return _getComponentInstanceStorage(arg0)
}

var _getComponentListModSeed func() int32

// GetComponentListModSeed allows your application to determine if the list of registered components has changed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516399-getcomponentlistmodseed
func GetComponentListModSeed() int32 {
	if _getComponentListModSeed == nil {
		panic("coreservices: symbol GetComponentListModSeed not loaded")
	}
	return _getComponentListModSeed()
}

var _getComponentPublicIndString func(arg0 Component, arg1 unsafe.Pointer, arg2 int16, arg3 int16) int16

// GetComponentPublicIndString.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516516-getcomponentpublicindstring
func GetComponentPublicIndString(arg0 Component, arg1 unsafe.Pointer, arg2 int16, arg3 int16) int16 {
	if _getComponentPublicIndString == nil {
		panic("coreservices: symbol GetComponentPublicIndString not loaded")
	}
	return _getComponentPublicIndString(arg0, arg1, arg2, arg3)
}

var _getComponentPublicResource func(arg0 Component, arg1 uint32, arg2 int16, arg3 uintptr) int16

// GetComponentPublicResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516336-getcomponentpublicresource
func GetComponentPublicResource(arg0 Component, arg1 uint32, arg2 int16, arg3 uintptr) int16 {
	if _getComponentPublicResource == nil {
		panic("coreservices: symbol GetComponentPublicResource not loaded")
	}
	return _getComponentPublicResource(arg0, arg1, arg2, arg3)
}

var _getComponentPublicResourceList func(arg0 uint32, arg1 int16, arg2 int32, arg3 ComponentDescription, arg4 GetMissingComponentResourceUPP) int16

// GetComponentPublicResourceList.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516460-getcomponentpublicresourcelist
func GetComponentPublicResourceList(arg0 uint32, arg1 int16, arg2 int32, arg3 ComponentDescription, arg4 GetMissingComponentResourceUPP) int16 {
	if _getComponentPublicResourceList == nil {
		panic("coreservices: symbol GetComponentPublicResourceList not loaded")
	}
	return _getComponentPublicResourceList(arg0, arg1, arg2, arg3, arg4)
}

var _getComponentRefcon func(arg0 Component) int

// GetComponentRefcon retrieves the value of the reference constant for your component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516331-getcomponentrefcon
func GetComponentRefcon(arg0 Component) int {
	if _getComponentRefcon == nil {
		panic("coreservices: symbol GetComponentRefcon not loaded")
	}
	return _getComponentRefcon(arg0)
}

var _getComponentResource func(arg0 Component, arg1 uint32, arg2 int16, arg3 uintptr) int16

// GetComponentResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516387-getcomponentresource
func GetComponentResource(arg0 Component, arg1 uint32, arg2 int16, arg3 uintptr) int16 {
	if _getComponentResource == nil {
		panic("coreservices: symbol GetComponentResource not loaded")
	}
	return _getComponentResource(arg0, arg1, arg2, arg3)
}

var _getComponentTypeModSeed func(arg0 uint32) int32

// GetComponentTypeModSeed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516653-getcomponenttypemodseed
func GetComponentTypeModSeed(arg0 uint32) int32 {
	if _getComponentTypeModSeed == nil {
		panic("coreservices: symbol GetComponentTypeModSeed not loaded")
	}
	return _getComponentTypeModSeed(arg0)
}

var _getCurrentThread func(arg0 ThreadID) int16

// GetCurrentThread.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574216-getcurrentthread
func GetCurrentThread(arg0 ThreadID) int16 {
	if _getCurrentThread == nil {
		panic("coreservices: symbol GetCurrentThread not loaded")
	}
	return _getCurrentThread(arg0)
}

var _getCustomIconsEnabled func(vRefNum unsafe.Pointer, customIconsEnabled unsafe.Pointer) int16

// GetCustomIconsEnabled.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1442255-getcustomiconsenabled
func GetCustomIconsEnabled(vRefNum unsafe.Pointer, customIconsEnabled unsafe.Pointer) int16 {
	if _getCustomIconsEnabled == nil {
		panic("coreservices: symbol GetCustomIconsEnabled not loaded")
	}
	return _getCustomIconsEnabled(vRefNum, customIconsEnabled)
}

var _getDebugComponentInfo func(arg0 uint32, arg1 uint32, arg2 unsafe.Pointer) int32

// GetDebugComponentInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517776-getdebugcomponentinfo
func GetDebugComponentInfo(arg0 uint32, arg1 uint32, arg2 unsafe.Pointer) int32 {
	if _getDebugComponentInfo == nil {
		panic("coreservices: symbol GetDebugComponentInfo not loaded")
	}
	return _getDebugComponentInfo(arg0, arg1, arg2)
}

var _getDebugOptionInfo func(arg0 uint32, arg1 uint32, arg2 int32, arg3 unsafe.Pointer, arg4 bool) int32

// GetDebugOptionInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517775-getdebugoptioninfo
func GetDebugOptionInfo(arg0 uint32, arg1 uint32, arg2 int32, arg3 unsafe.Pointer, arg4 bool) int32 {
	if _getDebugOptionInfo == nil {
		panic("coreservices: symbol GetDebugOptionInfo not loaded")
	}
	return _getDebugOptionInfo(arg0, arg1, arg2, arg3, arg4)
}

var _getDefaultThreadStackSize func(arg0 ThreadStyle, arg1 corefoundation.CGSize) int16

// GetDefaultThreadStackSize.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574231-getdefaultthreadstacksize
func GetDefaultThreadStackSize(arg0 ThreadStyle, arg1 corefoundation.CGSize) int16 {
	if _getDefaultThreadStackSize == nil {
		panic("coreservices: symbol GetDefaultThreadStackSize not loaded")
	}
	return _getDefaultThreadStackSize(arg0, arg1)
}

var _getFolderNameUnicode func(arg0 uintptr, arg1 uint32, arg2 uintptr, arg3 unsafe.Pointer) int32

// GetFolderNameUnicode.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389375-getfoldernameunicode
func GetFolderNameUnicode(arg0 uintptr, arg1 uint32, arg2 uintptr, arg3 unsafe.Pointer) int32 {
	if _getFolderNameUnicode == nil {
		panic("coreservices: symbol GetFolderNameUnicode not loaded")
	}
	return _getFolderNameUnicode(arg0, arg1, arg2, arg3)
}

var _getFolderTypes func(arg0 uint32, arg1 uint32, arg2 FolderType) int16

// GetFolderTypes.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389428-getfoldertypes
func GetFolderTypes(arg0 uint32, arg1 uint32, arg2 FolderType) int16 {
	if _getFolderTypes == nil {
		panic("coreservices: symbol GetFolderTypes not loaded")
	}
	return _getFolderTypes(arg0, arg1, arg2)
}

var _getHandleSize func(arg0 uintptr) corefoundation.CGSize

// GetHandleSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506398-gethandlesize
func GetHandleSize(arg0 uintptr) corefoundation.CGSize {
	if _getHandleSize == nil {
		panic("coreservices: symbol GetHandleSize not loaded")
	}
	return _getHandleSize(arg0)
}

var _getIconRef func(vRefNum unsafe.Pointer, creator uint32, iconType uint32, theIconRef uintptr) int16

// GetIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1442776-geticonref
func GetIconRef(vRefNum unsafe.Pointer, creator uint32, iconType uint32, theIconRef uintptr) int16 {
	if _getIconRef == nil {
		panic("coreservices: symbol GetIconRef not loaded")
	}
	return _getIconRef(vRefNum, creator, iconType, theIconRef)
}

var _getIconRefFromComponent func(arg0 Component, arg1 uintptr) int32

// GetIconRefFromComponent.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1447113-geticonreffromcomponent
func GetIconRefFromComponent(arg0 Component, arg1 uintptr) int32 {
	if _getIconRefFromComponent == nil {
		panic("coreservices: symbol GetIconRefFromComponent not loaded")
	}
	return _getIconRefFromComponent(arg0, arg1)
}

var _getIconRefFromFileInfo func(inRef unsafe.Pointer, inFileNameLength unsafe.Pointer, inFileName *uint16, inWhichInfo FSCatalogInfoBitmap, inCatalogInfo unsafe.Pointer, inUsageFlags IconServicesUsageFlags, outIconRef uintptr, outLabel *uintptr) int32

// GetIconRefFromFileInfo.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/coreservices/1447966-geticonreffromfileinfo
func GetIconRefFromFileInfo(inRef unsafe.Pointer, inFileNameLength unsafe.Pointer, inFileName *uint16, inWhichInfo FSCatalogInfoBitmap, inCatalogInfo unsafe.Pointer, inUsageFlags IconServicesUsageFlags, outIconRef uintptr, outLabel *uintptr) int32 {
	if _getIconRefFromFileInfo == nil {
		panic("coreservices: symbol GetIconRefFromFileInfo not loaded")
	}
	return _getIconRefFromFileInfo(inRef, inFileNameLength, inFileName, inWhichInfo, inCatalogInfo, inUsageFlags, outIconRef, outLabel)
}

var _getIconRefFromFolder func(vRefNum unsafe.Pointer, parentFolderID unsafe.Pointer, folderID unsafe.Pointer, attributes int8, accessPrivileges int8, theIconRef uintptr) int16

// GetIconRefFromFolder.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1441712-geticonreffromfolder
func GetIconRefFromFolder(vRefNum unsafe.Pointer, parentFolderID unsafe.Pointer, folderID unsafe.Pointer, attributes int8, accessPrivileges int8, theIconRef uintptr) int16 {
	if _getIconRefFromFolder == nil {
		panic("coreservices: symbol GetIconRefFromFolder not loaded")
	}
	return _getIconRefFromFolder(vRefNum, parentFolderID, folderID, attributes, accessPrivileges, theIconRef)
}

var _getIconRefFromIconFamilyPtr func(inIconFamilyPtr unsafe.Pointer, inSize corefoundation.CGSize, outIconRef uintptr) int32

// GetIconRefFromIconFamilyPtr.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1443251-geticonreffromiconfamilyptr
func GetIconRefFromIconFamilyPtr(inIconFamilyPtr unsafe.Pointer, inSize corefoundation.CGSize, outIconRef uintptr) int32 {
	if _getIconRefFromIconFamilyPtr == nil {
		panic("coreservices: symbol GetIconRefFromIconFamilyPtr not loaded")
	}
	return _getIconRefFromIconFamilyPtr(inIconFamilyPtr, inSize, outIconRef)
}

var _getIconRefFromTypeInfo func(inCreator uint32, inType uint32, inExtension corefoundation.CFStringRef, inMIMEType corefoundation.CFStringRef, inUsageFlags IconServicesUsageFlags, outIconRef uintptr) int16

// GetIconRefFromTypeInfo.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1445758-geticonreffromtypeinfo
func GetIconRefFromTypeInfo(inCreator uint32, inType uint32, inExtension corefoundation.CFStringRef, inMIMEType corefoundation.CFStringRef, inUsageFlags IconServicesUsageFlags, outIconRef uintptr) int16 {
	if _getIconRefFromTypeInfo == nil {
		panic("coreservices: symbol GetIconRefFromTypeInfo not loaded")
	}
	return _getIconRefFromTypeInfo(inCreator, inType, inExtension, inMIMEType, inUsageFlags, outIconRef)
}

var _getIconRefOwners func(theIconRef uintptr, owners *uint16) int16

// GetIconRefOwners.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1447221-geticonrefowners
func GetIconRefOwners(theIconRef uintptr, owners *uint16) int16 {
	if _getIconRefOwners == nil {
		panic("coreservices: symbol GetIconRefOwners not loaded")
	}
	return _getIconRefOwners(theIconRef, owners)
}

var _getIndResource func(arg0 uint32, arg1 ResourceIndex) objectivec.IObject

// GetIndResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529227-getindresource
func GetIndResource(arg0 uint32, arg1 ResourceIndex) objectivec.IObject {
	if _getIndResource == nil {
		panic("coreservices: symbol GetIndResource not loaded")
	}
	return _getIndResource(arg0, arg1)
}

var _getIndType func(arg0 uint32, arg1 ResourceIndex)

// GetIndType.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529367-getindtype
func GetIndType(arg0 uint32, arg1 ResourceIndex) {
	if _getIndType == nil {
		panic("coreservices: symbol GetIndType not loaded")
	}
	_getIndType(arg0, arg1)
}

var _getIndexedCollectionItem func(arg0 Collection, arg1 int32, arg2 int32) int16

// GetIndexedCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551437-getindexedcollectionitem
func GetIndexedCollectionItem(arg0 Collection, arg1 int32, arg2 int32) int16 {
	if _getIndexedCollectionItem == nil {
		panic("coreservices: symbol GetIndexedCollectionItem not loaded")
	}
	return _getIndexedCollectionItem(arg0, arg1, arg2)
}

var _getIndexedCollectionItemHdl func(arg0 Collection, arg1 int32, arg2 uintptr) int16

// GetIndexedCollectionItemHdl.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551443-getindexedcollectionitemhdl
func GetIndexedCollectionItemHdl(arg0 Collection, arg1 int32, arg2 uintptr) int16 {
	if _getIndexedCollectionItemHdl == nil {
		panic("coreservices: symbol GetIndexedCollectionItemHdl not loaded")
	}
	return _getIndexedCollectionItemHdl(arg0, arg1, arg2)
}

var _getIndexedCollectionItemInfo func(arg0 Collection, arg1 int32, arg2 CollectionTag, arg3 int32, arg4 int32, arg5 int32) int16

// GetIndexedCollectionItemInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551449-getindexedcollectioniteminfo
func GetIndexedCollectionItemInfo(arg0 Collection, arg1 int32, arg2 CollectionTag, arg3 int32, arg4 int32, arg5 int32) int16 {
	if _getIndexedCollectionItemInfo == nil {
		panic("coreservices: symbol GetIndexedCollectionItemInfo not loaded")
	}
	return _getIndexedCollectionItemInfo(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _getIndexedCollectionTag func(arg0 Collection, arg1 int32, arg2 CollectionTag) int16

// GetIndexedCollectionTag.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551355-getindexedcollectiontag
func GetIndexedCollectionTag(arg0 Collection, arg1 int32, arg2 CollectionTag) int16 {
	if _getIndexedCollectionTag == nil {
		panic("coreservices: symbol GetIndexedCollectionTag not loaded")
	}
	return _getIndexedCollectionTag(arg0, arg1, arg2)
}

var _getMacOSStatusCommentString func(arg0 int32) *byte

// GetMacOSStatusCommentString.
//
// See: https://developer.apple.com/documentation/coreservices/1517787-getmacosstatuscommentstring
func GetMacOSStatusCommentString(arg0 int32) *byte {
	if _getMacOSStatusCommentString == nil {
		panic("coreservices: symbol GetMacOSStatusCommentString not loaded")
	}
	return _getMacOSStatusCommentString(arg0)
}

var _getMacOSStatusErrorString func(arg0 int32) *byte

// GetMacOSStatusErrorString.
//
// See: https://developer.apple.com/documentation/coreservices/1517786-getmacosstatuserrorstring
func GetMacOSStatusErrorString(arg0 int32) *byte {
	if _getMacOSStatusErrorString == nil {
		panic("coreservices: symbol GetMacOSStatusErrorString not loaded")
	}
	return _getMacOSStatusErrorString(arg0)
}

var _getMaxResourceSize func(arg0 uintptr) int

// GetMaxResourceSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529298-getmaxresourcesize
func GetMaxResourceSize(arg0 uintptr) int {
	if _getMaxResourceSize == nil {
		panic("coreservices: symbol GetMaxResourceSize not loaded")
	}
	return _getMaxResourceSize(arg0)
}

var _getNamedResource func(arg0 uint32, arg1 unsafe.Pointer) objectivec.IObject

// GetNamedResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529296-getnamedresource
func GetNamedResource(arg0 uint32, arg1 unsafe.Pointer) objectivec.IObject {
	if _getNamedResource == nil {
		panic("coreservices: symbol GetNamedResource not loaded")
	}
	return _getNamedResource(arg0, arg1)
}

var _getNewCollection func(arg0 int16) Collection

// GetNewCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551366-getnewcollection
func GetNewCollection(arg0 int16) Collection {
	if _getNewCollection == nil {
		panic("coreservices: symbol GetNewCollection not loaded")
	}
	return _getNewCollection(arg0)
}

var _getNextFOND func(arg0 uintptr) objectivec.IObject

// GetNextFOND.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529354-getnextfond
func GetNextFOND(arg0 uintptr) objectivec.IObject {
	if _getNextFOND == nil {
		panic("coreservices: symbol GetNextFOND not loaded")
	}
	return _getNextFOND(arg0)
}

var _getNextResourceFile func(arg0 ResFileRefNum, arg1 ResFileRefNum) int16

// GetNextResourceFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529310-getnextresourcefile
func GetNextResourceFile(arg0 ResFileRefNum, arg1 ResFileRefNum) int16 {
	if _getNextResourceFile == nil {
		panic("coreservices: symbol GetNextResourceFile not loaded")
	}
	return _getNextResourceFile(arg0, arg1)
}

var _getPtrSize func(arg0 coreimage.Ptr) corefoundation.CGSize

// GetPtrSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506465-getptrsize
func GetPtrSize(arg0 coreimage.Ptr) corefoundation.CGSize {
	if _getPtrSize == nil {
		panic("coreservices: symbol GetPtrSize not loaded")
	}
	return _getPtrSize(arg0)
}

var _getResAttrs func(arg0 uintptr) ResAttributes

// GetResAttrs.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529365-getresattrs
func GetResAttrs(arg0 uintptr) ResAttributes {
	if _getResAttrs == nil {
		panic("coreservices: symbol GetResAttrs not loaded")
	}
	return _getResAttrs(arg0)
}

var _getResFileAttrs func(arg0 ResFileRefNum) ResFileAttributes

// GetResFileAttrs.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529231-getresfileattrs
func GetResFileAttrs(arg0 ResFileRefNum) ResFileAttributes {
	if _getResFileAttrs == nil {
		panic("coreservices: symbol GetResFileAttrs not loaded")
	}
	return _getResFileAttrs(arg0)
}

var _getResInfo func(arg0 uintptr, arg1 ResID, arg2 uint32, arg3 unsafe.Pointer)

// GetResInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529369-getresinfo
func GetResInfo(arg0 uintptr, arg1 ResID, arg2 uint32, arg3 unsafe.Pointer) {
	if _getResInfo == nil {
		panic("coreservices: symbol GetResInfo not loaded")
	}
	_getResInfo(arg0, arg1, arg2, arg3)
}

var _getResource func(arg0 uint32, arg1 ResID) objectivec.IObject

// GetResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529302-getresource
func GetResource(arg0 uint32, arg1 ResID) objectivec.IObject {
	if _getResource == nil {
		panic("coreservices: symbol GetResource not loaded")
	}
	return _getResource(arg0, arg1)
}

var _getResourceSizeOnDisk func(arg0 uintptr) int

// GetResourceSizeOnDisk.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529348-getresourcesizeondisk
func GetResourceSizeOnDisk(arg0 uintptr) int {
	if _getResourceSizeOnDisk == nil {
		panic("coreservices: symbol GetResourceSizeOnDisk not loaded")
	}
	return _getResourceSizeOnDisk(arg0)
}

var _getScriptInfoFromTextEncoding func(arg0 TextEncoding, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32

// GetScriptInfoFromTextEncoding.
//
// See: https://developer.apple.com/documentation/coreservices/1399675-getscriptinfofromtextencoding
func GetScriptInfoFromTextEncoding(arg0 TextEncoding, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32 {
	if _getScriptInfoFromTextEncoding == nil {
		panic("coreservices: symbol GetScriptInfoFromTextEncoding not loaded")
	}
	return _getScriptInfoFromTextEncoding(arg0, arg1, arg2)
}

var _getScriptManagerVariable func(arg0 int16) int

// GetScriptManagerVariable.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/coreservices/1485022-getscriptmanagervariable
func GetScriptManagerVariable(arg0 int16) int {
	if _getScriptManagerVariable == nil {
		panic("coreservices: symbol GetScriptManagerVariable not loaded")
	}
	return _getScriptManagerVariable(arg0)
}

var _getTaggedCollectionItem func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32) int16

// GetTaggedCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551359-gettaggedcollectionitem
func GetTaggedCollectionItem(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32) int16 {
	if _getTaggedCollectionItem == nil {
		panic("coreservices: symbol GetTaggedCollectionItem not loaded")
	}
	return _getTaggedCollectionItem(arg0, arg1, arg2, arg3)
}

var _getTaggedCollectionItemInfo func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32, arg4 int32, arg5 int32, arg6 int32) int16

// GetTaggedCollectionItemInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551370-gettaggedcollectioniteminfo
func GetTaggedCollectionItemInfo(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32, arg4 int32, arg5 int32, arg6 int32) int16 {
	if _getTaggedCollectionItemInfo == nil {
		panic("coreservices: symbol GetTaggedCollectionItemInfo not loaded")
	}
	return _getTaggedCollectionItemInfo(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _getTextEncodingBase func(arg0 TextEncoding) TextEncodingBase

// GetTextEncodingBase returns the base encoding of the specified text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1399792-gettextencodingbase
func GetTextEncodingBase(arg0 TextEncoding) TextEncodingBase {
	if _getTextEncodingBase == nil {
		panic("coreservices: symbol GetTextEncodingBase not loaded")
	}
	return _getTextEncodingBase(arg0)
}

var _getTextEncodingFormat func(arg0 TextEncoding) TextEncodingFormat

// GetTextEncodingFormat returns the format value of the specified text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1400318-gettextencodingformat
func GetTextEncodingFormat(arg0 TextEncoding) TextEncodingFormat {
	if _getTextEncodingFormat == nil {
		panic("coreservices: symbol GetTextEncodingFormat not loaded")
	}
	return _getTextEncodingFormat(arg0)
}

var _getTextEncodingFromScriptInfo func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 TextEncoding) int32

// GetTextEncodingFromScriptInfo.
//
// See: https://developer.apple.com/documentation/coreservices/1399689-gettextencodingfromscriptinfo
func GetTextEncodingFromScriptInfo(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 TextEncoding) int32 {
	if _getTextEncodingFromScriptInfo == nil {
		panic("coreservices: symbol GetTextEncodingFromScriptInfo not loaded")
	}
	return _getTextEncodingFromScriptInfo(arg0, arg1, arg2, arg3)
}

var _getTextEncodingName func(arg0 TextEncoding, arg1 TextEncodingNameSelector, arg2 unsafe.Pointer, arg3 TextEncoding, arg4 uintptr, arg5 uintptr, arg6 unsafe.Pointer, arg7 TextEncoding, arg8 TextPtr) int32

// GetTextEncodingName returns the localized name for a specified text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1399645-gettextencodingname
func GetTextEncodingName(arg0 TextEncoding, arg1 TextEncodingNameSelector, arg2 unsafe.Pointer, arg3 TextEncoding, arg4 uintptr, arg5 uintptr, arg6 unsafe.Pointer, arg7 TextEncoding, arg8 TextPtr) int32 {
	if _getTextEncodingName == nil {
		panic("coreservices: symbol GetTextEncodingName not loaded")
	}
	return _getTextEncodingName(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _getTextEncodingVariant func(arg0 TextEncoding) TextEncodingVariant

// GetTextEncodingVariant returns the variant from the specified text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1400250-gettextencodingvariant
func GetTextEncodingVariant(arg0 TextEncoding) TextEncodingVariant {
	if _getTextEncodingVariant == nil {
		panic("coreservices: symbol GetTextEncodingVariant not loaded")
	}
	return _getTextEncodingVariant(arg0)
}

var _getThreadCurrentTaskRef func(arg0 ThreadTaskRef) int16

// GetThreadCurrentTaskRef.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574236-getthreadcurrenttaskref
func GetThreadCurrentTaskRef(arg0 ThreadTaskRef) int16 {
	if _getThreadCurrentTaskRef == nil {
		panic("coreservices: symbol GetThreadCurrentTaskRef not loaded")
	}
	return _getThreadCurrentTaskRef(arg0)
}

var _getThreadState func(arg0 ThreadID, arg1 ThreadState) int16

// GetThreadState.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574250-getthreadstate
func GetThreadState(arg0 ThreadID, arg1 ThreadState) int16 {
	if _getThreadState == nil {
		panic("coreservices: symbol GetThreadState not loaded")
	}
	return _getThreadState(arg0, arg1)
}

var _getThreadStateGivenTaskRef func(arg0 ThreadTaskRef, arg1 ThreadID, arg2 ThreadState) int16

// GetThreadStateGivenTaskRef.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574234-getthreadstategiventaskref
func GetThreadStateGivenTaskRef(arg0 ThreadTaskRef, arg1 ThreadID, arg2 ThreadState) int16 {
	if _getThreadStateGivenTaskRef == nil {
		panic("coreservices: symbol GetThreadStateGivenTaskRef not loaded")
	}
	return _getThreadStateGivenTaskRef(arg0, arg1, arg2)
}

var _getTopResourceFile func(arg0 ResFileRefNum) int16

// GetTopResourceFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529307-gettopresourcefile
func GetTopResourceFile(arg0 ResFileRefNum) int16 {
	if _getTopResourceFile == nil {
		panic("coreservices: symbol GetTopResourceFile not loaded")
	}
	return _getTopResourceFile(arg0)
}

var _hClrRBit func(arg0 uintptr)

// HClrRBit.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506459-hclrrbit
func HClrRBit(arg0 uintptr) {
	if _hClrRBit == nil {
		panic("coreservices: symbol HClrRBit not loaded")
	}
	_hClrRBit(arg0)
}

var _hGetState func(arg0 uintptr) int8

// HGetState.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506392-hgetstate
func HGetState(arg0 uintptr) int8 {
	if _hGetState == nil {
		panic("coreservices: symbol HGetState not loaded")
	}
	return _hGetState(arg0)
}

var _hLock func(arg0 uintptr)

// HLock.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506234-hlock
func HLock(arg0 uintptr) {
	if _hLock == nil {
		panic("coreservices: symbol HLock not loaded")
	}
	_hLock(arg0)
}

var _hLockHi func(arg0 uintptr)

// HLockHi.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506354-hlockhi
func HLockHi(arg0 uintptr) {
	if _hLockHi == nil {
		panic("coreservices: symbol HLockHi not loaded")
	}
	_hLockHi(arg0)
}

var _hSetRBit func(arg0 uintptr)

// HSetRBit.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506324-hsetrbit
func HSetRBit(arg0 uintptr) {
	if _hSetRBit == nil {
		panic("coreservices: symbol HSetRBit not loaded")
	}
	_hSetRBit(arg0)
}

var _hSetState func(arg0 uintptr, arg1 int8)

// HSetState.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506435-hsetstate
func HSetState(arg0 uintptr, arg1 int8) {
	if _hSetState == nil {
		panic("coreservices: symbol HSetState not loaded")
	}
	_hSetState(arg0, arg1)
}

var _hUnlock func(arg0 uintptr)

// HUnlock.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506251-hunlock
func HUnlock(arg0 uintptr) {
	if _hUnlock == nil {
		panic("coreservices: symbol HUnlock not loaded")
	}
	_hUnlock(arg0)
}

var _handAndHand func(arg0 uintptr, arg1 uintptr) int16

// HandAndHand.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506243-handandhand
func HandAndHand(arg0 uintptr, arg1 uintptr) int16 {
	if _handAndHand == nil {
		panic("coreservices: symbol HandAndHand not loaded")
	}
	return _handAndHand(arg0, arg1)
}

var _handToHand func(arg0 uintptr) int16

// HandToHand.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506272-handtohand
func HandToHand(arg0 uintptr) int16 {
	if _handToHand == nil {
		panic("coreservices: symbol HandToHand not loaded")
	}
	return _handToHand(arg0)
}

var _homeResFile func(arg0 uintptr) ResFileRefNum

// HomeResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529352-homeresfile
func HomeResFile(arg0 uintptr) ResFileRefNum {
	if _homeResFile == nil {
		panic("coreservices: symbol HomeResFile not loaded")
	}
	return _homeResFile(arg0)
}

var _identifyFolder func(arg0 uintptr, arg1 int32, arg2 FolderType) int16

// IdentifyFolder.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389047-identifyfolder
func IdentifyFolder(arg0 uintptr, arg1 int32, arg2 FolderType) int16 {
	if _identifyFolder == nil {
		panic("coreservices: symbol IdentifyFolder not loaded")
	}
	return _identifyFolder(arg0, arg1, arg2)
}

var _incrementAtomic func(arg0 int32) int32

// IncrementAtomic.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490602-incrementatomic
func IncrementAtomic(arg0 int32) int32 {
	if _incrementAtomic == nil {
		panic("coreservices: symbol IncrementAtomic not loaded")
	}
	return _incrementAtomic(arg0)
}

var _incrementAtomic16 func(arg0 int16) int16

// IncrementAtomic16.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490600-incrementatomic16
func IncrementAtomic16(arg0 int16) int16 {
	if _incrementAtomic16 == nil {
		panic("coreservices: symbol IncrementAtomic16 not loaded")
	}
	return _incrementAtomic16(arg0)
}

var _incrementAtomic8 func(arg0 int8) int8

// IncrementAtomic8.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490585-incrementatomic8
func IncrementAtomic8(arg0 int8) int8 {
	if _incrementAtomic8 == nil {
		panic("coreservices: symbol IncrementAtomic8 not loaded")
	}
	return _incrementAtomic8(arg0)
}

var _insTime func(arg0 QElemPtr)

// InsTime.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550782-instime
func InsTime(arg0 QElemPtr) {
	if _insTime == nil {
		panic("coreservices: symbol InsTime not loaded")
	}
	_insTime(arg0)
}

var _insXTime func(arg0 QElemPtr)

// InsXTime.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550792-insxtime
func InsXTime(arg0 QElemPtr) {
	if _insXTime == nil {
		panic("coreservices: symbol InsXTime not loaded")
	}
	_insXTime(arg0)
}

var _insertResourceFile func(arg0 ResFileRefNum, arg1 RsrcChainLocation) int16

// InsertResourceFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529327-insertresourcefile
func InsertResourceFile(arg0 ResFileRefNum, arg1 RsrcChainLocation) int16 {
	if _insertResourceFile == nil {
		panic("coreservices: symbol InsertResourceFile not loaded")
	}
	return _insertResourceFile(arg0, arg1)
}

var _installDebugAssertOutputHandler func(arg0 DebugAssertOutputHandlerUPP)

// InstallDebugAssertOutputHandler.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517758-installdebugassertoutputhandler
func InstallDebugAssertOutputHandler(arg0 DebugAssertOutputHandlerUPP) {
	if _installDebugAssertOutputHandler == nil {
		panic("coreservices: symbol InstallDebugAssertOutputHandler not loaded")
	}
	_installDebugAssertOutputHandler(arg0)
}

var _installExceptionHandler func(arg0 ExceptionHandlerTPP) ExceptionHandlerTPP

// InstallExceptionHandler.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1500730-installexceptionhandler
func InstallExceptionHandler(arg0 ExceptionHandlerTPP) ExceptionHandlerTPP {
	if _installExceptionHandler == nil {
		panic("coreservices: symbol InstallExceptionHandler not loaded")
	}
	return _installExceptionHandler(arg0)
}

var _installTimeTask func(arg0 QElemPtr) int16

// InstallTimeTask.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550791-installtimetask
func InstallTimeTask(arg0 QElemPtr) int16 {
	if _installTimeTask == nil {
		panic("coreservices: symbol InstallTimeTask not loaded")
	}
	return _installTimeTask(arg0)
}

var _installXTimeTask func(arg0 QElemPtr) int16

// InstallXTimeTask.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550795-installxtimetask
func InstallXTimeTask(arg0 QElemPtr) int16 {
	if _installXTimeTask == nil {
		panic("coreservices: symbol InstallXTimeTask not loaded")
	}
	return _installXTimeTask(arg0)
}

var _invalidateFolderDescriptorCache func(arg0 uintptr, arg1 int32) int16

// InvalidateFolderDescriptorCache.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389097-invalidatefolderdescriptorcache
func InvalidateFolderDescriptorCache(arg0 uintptr, arg1 int32) int16 {
	if _invalidateFolderDescriptorCache == nil {
		panic("coreservices: symbol InvalidateFolderDescriptorCache not loaded")
	}
	return _invalidateFolderDescriptorCache(arg0, arg1)
}

var _invokeAECoerceDescUPP func(fromDesc unsafe.Pointer, toType uint32, handlerRefcon uintptr, toDesc unsafe.Pointer, userUPP AECoerceDescUPP) int16

// InvokeAECoerceDescUPP calls a universal procedure pointer to a function that coerces data stored in a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1445450-invokeaecoercedescupp
func InvokeAECoerceDescUPP(fromDesc unsafe.Pointer, toType uint32, handlerRefcon uintptr, toDesc unsafe.Pointer, userUPP AECoerceDescUPP) int16 {
	if _invokeAECoerceDescUPP == nil {
		panic("coreservices: symbol InvokeAECoerceDescUPP not loaded")
	}
	return _invokeAECoerceDescUPP(fromDesc, toType, handlerRefcon, toDesc, userUPP)
}

var _invokeAECoercePtrUPP func(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, toType uint32, handlerRefcon uintptr, result unsafe.Pointer, userUPP AECoercePtrUPP) int16

// InvokeAECoercePtrUPP calls a universal procedure pointer to a function that coerces data stored in a buffer.
//
// See: https://developer.apple.com/documentation/coreservices/1447079-invokeaecoerceptrupp
func InvokeAECoercePtrUPP(typeCode uint32, dataPtr unsafe.Pointer, dataSize corefoundation.CGSize, toType uint32, handlerRefcon uintptr, result unsafe.Pointer, userUPP AECoercePtrUPP) int16 {
	if _invokeAECoercePtrUPP == nil {
		panic("coreservices: symbol InvokeAECoercePtrUPP not loaded")
	}
	return _invokeAECoercePtrUPP(typeCode, dataPtr, dataSize, toType, handlerRefcon, result, userUPP)
}

var _invokeAEDisposeExternalUPP func(dataPtr unsafe.Pointer, dataLength corefoundation.CGSize, refcon uintptr, userUPP AEDisposeExternalUPP)

// InvokeAEDisposeExternalUPP calls a dispose external universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1441717-invokeaedisposeexternalupp
func InvokeAEDisposeExternalUPP(dataPtr unsafe.Pointer, dataLength corefoundation.CGSize, refcon uintptr, userUPP AEDisposeExternalUPP) {
	if _invokeAEDisposeExternalUPP == nil {
		panic("coreservices: symbol InvokeAEDisposeExternalUPP not loaded")
	}
	_invokeAEDisposeExternalUPP(dataPtr, dataLength, refcon, userUPP)
}

var _invokeAEEventHandlerUPP func(theAppleEvent *AEDesc, reply *AEDesc, handlerRefcon uintptr, userUPP AEEventHandlerUPP) int16

// InvokeAEEventHandlerUPP calls an event handler universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1446585-invokeaeeventhandlerupp
func InvokeAEEventHandlerUPP(theAppleEvent *AEDesc, reply *AEDesc, handlerRefcon uintptr, userUPP AEEventHandlerUPP) int16 {
	if _invokeAEEventHandlerUPP == nil {
		panic("coreservices: symbol InvokeAEEventHandlerUPP not loaded")
	}
	return _invokeAEEventHandlerUPP(theAppleEvent, reply, handlerRefcon, userUPP)
}

var _invokeCollectionExceptionUPP func(arg0 Collection, arg1 int16, arg2 CollectionExceptionUPP) int16

// InvokeCollectionExceptionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551347-invokecollectionexceptionupp
func InvokeCollectionExceptionUPP(arg0 Collection, arg1 int16, arg2 CollectionExceptionUPP) int16 {
	if _invokeCollectionExceptionUPP == nil {
		panic("coreservices: symbol InvokeCollectionExceptionUPP not loaded")
	}
	return _invokeCollectionExceptionUPP(arg0, arg1, arg2)
}

var _invokeCollectionFlattenUPP func(arg0 int32, arg1 CollectionFlattenUPP) int16

// InvokeCollectionFlattenUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551380-invokecollectionflattenupp
func InvokeCollectionFlattenUPP(arg0 int32, arg1 CollectionFlattenUPP) int16 {
	if _invokeCollectionFlattenUPP == nil {
		panic("coreservices: symbol InvokeCollectionFlattenUPP not loaded")
	}
	return _invokeCollectionFlattenUPP(arg0, arg1)
}

var _invokeComponentMPWorkFunctionUPP func(arg0 ComponentMPWorkFunctionHeaderRecordPtr, arg1 ComponentMPWorkFunctionUPP) ComponentResult

// InvokeComponentMPWorkFunctionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516356-invokecomponentmpworkfunctionupp
func InvokeComponentMPWorkFunctionUPP(arg0 ComponentMPWorkFunctionHeaderRecordPtr, arg1 ComponentMPWorkFunctionUPP) ComponentResult {
	if _invokeComponentMPWorkFunctionUPP == nil {
		panic("coreservices: symbol InvokeComponentMPWorkFunctionUPP not loaded")
	}
	return _invokeComponentMPWorkFunctionUPP(arg0, arg1)
}

var _invokeComponentRoutineUPP func(arg0 ComponentParameters, arg1 uintptr, arg2 ComponentRoutineUPP) ComponentResult

// InvokeComponentRoutineUPP calls your component routine callback function
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516654-invokecomponentroutineupp
func InvokeComponentRoutineUPP(arg0 ComponentParameters, arg1 uintptr, arg2 ComponentRoutineUPP) ComponentResult {
	if _invokeComponentRoutineUPP == nil {
		panic("coreservices: symbol InvokeComponentRoutineUPP not loaded")
	}
	return _invokeComponentRoutineUPP(arg0, arg1, arg2)
}

var _invokeDebugAssertOutputHandlerUPP func(arg0 uint32, arg1 uint32, arg2 int8, arg3 int8, arg4 int8, arg5 int8, arg6 int, arg7 unsafe.Pointer, arg8 DebugAssertOutputHandlerUPP)

// InvokeDebugAssertOutputHandlerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517712-invokedebugassertoutputhandlerup
func InvokeDebugAssertOutputHandlerUPP(arg0 uint32, arg1 uint32, arg2 int8, arg3 int8, arg4 int8, arg5 int8, arg6 int, arg7 unsafe.Pointer, arg8 DebugAssertOutputHandlerUPP) {
	if _invokeDebugAssertOutputHandlerUPP == nil {
		panic("coreservices: symbol InvokeDebugAssertOutputHandlerUPP not loaded")
	}
	_invokeDebugAssertOutputHandlerUPP(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _invokeDebugComponentCallbackUPP func(arg0 int32, arg1 uint32, arg2 bool, arg3 DebugComponentCallbackUPP)

// InvokeDebugComponentCallbackUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517770-invokedebugcomponentcallbackupp
func InvokeDebugComponentCallbackUPP(arg0 int32, arg1 uint32, arg2 bool, arg3 DebugComponentCallbackUPP) {
	if _invokeDebugComponentCallbackUPP == nil {
		panic("coreservices: symbol InvokeDebugComponentCallbackUPP not loaded")
	}
	_invokeDebugComponentCallbackUPP(arg0, arg1, arg2, arg3)
}

var _invokeDebuggerDisposeThreadUPP func(arg0 ThreadID, arg1 DebuggerDisposeThreadUPP)

// InvokeDebuggerDisposeThreadUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574308-invokedebuggerdisposethreadupp
func InvokeDebuggerDisposeThreadUPP(arg0 ThreadID, arg1 DebuggerDisposeThreadUPP) {
	if _invokeDebuggerDisposeThreadUPP == nil {
		panic("coreservices: symbol InvokeDebuggerDisposeThreadUPP not loaded")
	}
	_invokeDebuggerDisposeThreadUPP(arg0, arg1)
}

var _invokeDebuggerNewThreadUPP func(arg0 ThreadID, arg1 DebuggerNewThreadUPP)

// InvokeDebuggerNewThreadUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574207-invokedebuggernewthreadupp
func InvokeDebuggerNewThreadUPP(arg0 ThreadID, arg1 DebuggerNewThreadUPP) {
	if _invokeDebuggerNewThreadUPP == nil {
		panic("coreservices: symbol InvokeDebuggerNewThreadUPP not loaded")
	}
	_invokeDebuggerNewThreadUPP(arg0, arg1)
}

var _invokeDebuggerThreadSchedulerUPP func(arg0 SchedulerInfoRecPtr, arg1 DebuggerThreadSchedulerUPP) ThreadID

// InvokeDebuggerThreadSchedulerUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574211-invokedebuggerthreadschedulerupp
func InvokeDebuggerThreadSchedulerUPP(arg0 SchedulerInfoRecPtr, arg1 DebuggerThreadSchedulerUPP) ThreadID {
	if _invokeDebuggerThreadSchedulerUPP == nil {
		panic("coreservices: symbol InvokeDebuggerThreadSchedulerUPP not loaded")
	}
	return _invokeDebuggerThreadSchedulerUPP(arg0, arg1)
}

var _invokeDeferredTaskUPP func(arg0 int, arg1 DeferredTaskUPP)

// InvokeDeferredTaskUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533349-invokedeferredtaskupp
func InvokeDeferredTaskUPP(arg0 int, arg1 DeferredTaskUPP) {
	if _invokeDeferredTaskUPP == nil {
		panic("coreservices: symbol InvokeDeferredTaskUPP not loaded")
	}
	_invokeDeferredTaskUPP(arg0, arg1)
}

var _invokeExceptionHandlerUPP func(arg0 ExceptionInformation, arg1 ExceptionHandlerUPP) int32

// InvokeExceptionHandlerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1500606-invokeexceptionhandlerupp
func InvokeExceptionHandlerUPP(arg0 ExceptionInformation, arg1 ExceptionHandlerUPP) int32 {
	if _invokeExceptionHandlerUPP == nil {
		panic("coreservices: symbol InvokeExceptionHandlerUPP not loaded")
	}
	return _invokeExceptionHandlerUPP(arg0, arg1)
}

var _invokeFNSubscriptionUPP func(arg0 FNMessage, arg1 uintptr, arg2 FNSubscriptionRef, arg3 FNSubscriptionUPP)

// InvokeFNSubscriptionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565242-invokefnsubscriptionupp
func InvokeFNSubscriptionUPP(arg0 FNMessage, arg1 uintptr, arg2 FNSubscriptionRef, arg3 FNSubscriptionUPP) {
	if _invokeFNSubscriptionUPP == nil {
		panic("coreservices: symbol InvokeFNSubscriptionUPP not loaded")
	}
	_invokeFNSubscriptionUPP(arg0, arg1, arg2, arg3)
}

var _invokeFSVolumeEjectUPP func(arg0 FSVolumeOperation, arg1 int32, arg2 uintptr, arg3 int32, arg4 FSVolumeEjectUPP)

// InvokeFSVolumeEjectUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565652-invokefsvolumeejectupp
func InvokeFSVolumeEjectUPP(arg0 FSVolumeOperation, arg1 int32, arg2 uintptr, arg3 int32, arg4 FSVolumeEjectUPP) {
	if _invokeFSVolumeEjectUPP == nil {
		panic("coreservices: symbol InvokeFSVolumeEjectUPP not loaded")
	}
	_invokeFSVolumeEjectUPP(arg0, arg1, arg2, arg3, arg4)
}

var _invokeFSVolumeMountUPP func(arg0 FSVolumeOperation, arg1 int32, arg2 uintptr, arg3 FSVolumeMountUPP)

// InvokeFSVolumeMountUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566167-invokefsvolumemountupp
func InvokeFSVolumeMountUPP(arg0 FSVolumeOperation, arg1 int32, arg2 uintptr, arg3 FSVolumeMountUPP) {
	if _invokeFSVolumeMountUPP == nil {
		panic("coreservices: symbol InvokeFSVolumeMountUPP not loaded")
	}
	_invokeFSVolumeMountUPP(arg0, arg1, arg2, arg3)
}

var _invokeFSVolumeUnmountUPP func(arg0 FSVolumeOperation, arg1 int32, arg2 uintptr, arg3 int32, arg4 FSVolumeUnmountUPP)

// InvokeFSVolumeUnmountUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565842-invokefsvolumeunmountupp
func InvokeFSVolumeUnmountUPP(arg0 FSVolumeOperation, arg1 int32, arg2 uintptr, arg3 int32, arg4 FSVolumeUnmountUPP) {
	if _invokeFSVolumeUnmountUPP == nil {
		panic("coreservices: symbol InvokeFSVolumeUnmountUPP not loaded")
	}
	_invokeFSVolumeUnmountUPP(arg0, arg1, arg2, arg3, arg4)
}

var _invokeFolderManagerNotificationUPP func(arg0 uint32, arg1 FolderManagerNotificationUPP) int32

// InvokeFolderManagerNotificationUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389104-invokefoldermanagernotificationu
func InvokeFolderManagerNotificationUPP(arg0 uint32, arg1 FolderManagerNotificationUPP) int32 {
	if _invokeFolderManagerNotificationUPP == nil {
		panic("coreservices: symbol InvokeFolderManagerNotificationUPP not loaded")
	}
	return _invokeFolderManagerNotificationUPP(arg0, arg1)
}

var _invokeGetMissingComponentResourceUPP func(arg0 Component, arg1 uint32, arg2 int16, arg3 uintptr, arg4 GetMissingComponentResourceUPP) int16

// InvokeGetMissingComponentResourceUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516491-invokegetmissingcomponentresourc
func InvokeGetMissingComponentResourceUPP(arg0 Component, arg1 uint32, arg2 int16, arg3 uintptr, arg4 GetMissingComponentResourceUPP) int16 {
	if _invokeGetMissingComponentResourceUPP == nil {
		panic("coreservices: symbol InvokeGetMissingComponentResourceUPP not loaded")
	}
	return _invokeGetMissingComponentResourceUPP(arg0, arg1, arg2, arg3, arg4)
}

var _invokeIOCompletionUPP func(arg0 ParmBlkPtr, arg1 IOCompletionUPP)

// InvokeIOCompletionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565303-invokeiocompletionupp
func InvokeIOCompletionUPP(arg0 ParmBlkPtr, arg1 IOCompletionUPP) {
	if _invokeIOCompletionUPP == nil {
		panic("coreservices: symbol InvokeIOCompletionUPP not loaded")
	}
	_invokeIOCompletionUPP(arg0, arg1)
}

var _invokeIndexToUCStringUPP func(index uint32, listDataPtr unsafe.Pointer, refcon uintptr, outString *corefoundation.CFStringRef, tsOptions *UCTypeSelectOptions, userUPP IndexToUCStringUPP) bool

// InvokeIndexToUCStringUPP.
//
// See: https://developer.apple.com/documentation/coreservices/1390660-invokeindextoucstringupp
func InvokeIndexToUCStringUPP(index uint32, listDataPtr unsafe.Pointer, refcon uintptr, outString *corefoundation.CFStringRef, tsOptions *UCTypeSelectOptions, userUPP IndexToUCStringUPP) bool {
	if _invokeIndexToUCStringUPP == nil {
		panic("coreservices: symbol InvokeIndexToUCStringUPP not loaded")
	}
	return _invokeIndexToUCStringUPP(index, listDataPtr, refcon, outString, tsOptions, userUPP)
}

var _invokeKCCallbackUPP func(arg0 KCEvent, arg1 KCCallbackInfo, arg2 KCCallbackUPP) int32

// InvokeKCCallbackUPP.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563117-invokekccallbackupp
func InvokeKCCallbackUPP(arg0 KCEvent, arg1 KCCallbackInfo, arg2 KCCallbackUPP) int32 {
	if _invokeKCCallbackUPP == nil {
		panic("coreservices: symbol InvokeKCCallbackUPP not loaded")
	}
	return _invokeKCCallbackUPP(arg0, arg1, arg2)
}

var _invokeOSLAccessorUPP func(desiredClass uint32, container unsafe.Pointer, containerClass uint32, form uint32, selectionData unsafe.Pointer, value unsafe.Pointer, accessorRefcon uintptr, userUPP OSLAccessorUPP) int16

// InvokeOSLAccessorUPP calls an object accessor universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1448978-invokeoslaccessorupp
func InvokeOSLAccessorUPP(desiredClass uint32, container unsafe.Pointer, containerClass uint32, form uint32, selectionData unsafe.Pointer, value unsafe.Pointer, accessorRefcon uintptr, userUPP OSLAccessorUPP) int16 {
	if _invokeOSLAccessorUPP == nil {
		panic("coreservices: symbol InvokeOSLAccessorUPP not loaded")
	}
	return _invokeOSLAccessorUPP(desiredClass, container, containerClass, form, selectionData, value, accessorRefcon, userUPP)
}

var _invokeOSLAdjustMarksUPP func(newStart unsafe.Pointer, newStop unsafe.Pointer, markToken unsafe.Pointer, userUPP OSLAdjustMarksUPP) int16

// InvokeOSLAdjustMarksUPP calls an object callback adjust marks universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1448506-invokeosladjustmarksupp
func InvokeOSLAdjustMarksUPP(newStart unsafe.Pointer, newStop unsafe.Pointer, markToken unsafe.Pointer, userUPP OSLAdjustMarksUPP) int16 {
	if _invokeOSLAdjustMarksUPP == nil {
		panic("coreservices: symbol InvokeOSLAdjustMarksUPP not loaded")
	}
	return _invokeOSLAdjustMarksUPP(newStart, newStop, markToken, userUPP)
}

var _invokeOSLCompareUPP func(oper uint32, obj1 unsafe.Pointer, obj2 unsafe.Pointer, result unsafe.Pointer, userUPP OSLCompareUPP) int16

// InvokeOSLCompareUPP calls an object callback comparison universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1443110-invokeoslcompareupp
func InvokeOSLCompareUPP(oper uint32, obj1 unsafe.Pointer, obj2 unsafe.Pointer, result unsafe.Pointer, userUPP OSLCompareUPP) int16 {
	if _invokeOSLCompareUPP == nil {
		panic("coreservices: symbol InvokeOSLCompareUPP not loaded")
	}
	return _invokeOSLCompareUPP(oper, obj1, obj2, result, userUPP)
}

var _invokeOSLCountUPP func(desiredType uint32, containerClass uint32, container unsafe.Pointer, result unsafe.Pointer, userUPP OSLCountUPP) int16

// InvokeOSLCountUPP calls an object callback count universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1448030-invokeoslcountupp
func InvokeOSLCountUPP(desiredType uint32, containerClass uint32, container unsafe.Pointer, result unsafe.Pointer, userUPP OSLCountUPP) int16 {
	if _invokeOSLCountUPP == nil {
		panic("coreservices: symbol InvokeOSLCountUPP not loaded")
	}
	return _invokeOSLCountUPP(desiredType, containerClass, container, result, userUPP)
}

var _invokeOSLDisposeTokenUPP func(unneededToken unsafe.Pointer, userUPP OSLDisposeTokenUPP) int16

// InvokeOSLDisposeTokenUPP calls an object callback dispose token universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1443963-invokeosldisposetokenupp
func InvokeOSLDisposeTokenUPP(unneededToken unsafe.Pointer, userUPP OSLDisposeTokenUPP) int16 {
	if _invokeOSLDisposeTokenUPP == nil {
		panic("coreservices: symbol InvokeOSLDisposeTokenUPP not loaded")
	}
	return _invokeOSLDisposeTokenUPP(unneededToken, userUPP)
}

var _invokeOSLGetErrDescUPP func(appDescPtr unsafe.Pointer, userUPP OSLGetErrDescUPP) int16

// InvokeOSLGetErrDescUPP calls an object callback get error descriptor universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1448420-invokeoslgeterrdescupp
func InvokeOSLGetErrDescUPP(appDescPtr unsafe.Pointer, userUPP OSLGetErrDescUPP) int16 {
	if _invokeOSLGetErrDescUPP == nil {
		panic("coreservices: symbol InvokeOSLGetErrDescUPP not loaded")
	}
	return _invokeOSLGetErrDescUPP(appDescPtr, userUPP)
}

var _invokeOSLGetMarkTokenUPP func(dContainerToken unsafe.Pointer, containerClass uint32, result unsafe.Pointer, userUPP OSLGetMarkTokenUPP) int16

// InvokeOSLGetMarkTokenUPP calls an object callback get mark universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1441894-invokeoslgetmarktokenupp
func InvokeOSLGetMarkTokenUPP(dContainerToken unsafe.Pointer, containerClass uint32, result unsafe.Pointer, userUPP OSLGetMarkTokenUPP) int16 {
	if _invokeOSLGetMarkTokenUPP == nil {
		panic("coreservices: symbol InvokeOSLGetMarkTokenUPP not loaded")
	}
	return _invokeOSLGetMarkTokenUPP(dContainerToken, containerClass, result, userUPP)
}

var _invokeOSLMarkUPP func(dToken unsafe.Pointer, markToken unsafe.Pointer, index unsafe.Pointer, userUPP OSLMarkUPP) int16

// InvokeOSLMarkUPP calls an object callback mark universal procedure pointer.
//
// See: https://developer.apple.com/documentation/coreservices/1447444-invokeoslmarkupp
func InvokeOSLMarkUPP(dToken unsafe.Pointer, markToken unsafe.Pointer, index unsafe.Pointer, userUPP OSLMarkUPP) int16 {
	if _invokeOSLMarkUPP == nil {
		panic("coreservices: symbol InvokeOSLMarkUPP not loaded")
	}
	return _invokeOSLMarkUPP(dToken, markToken, index, userUPP)
}

var _invokeResErrUPP func(arg0 int16, arg1 ResErrUPP)

// InvokeResErrUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529311-invokereserrupp
func InvokeResErrUPP(arg0 int16, arg1 ResErrUPP) {
	if _invokeResErrUPP == nil {
		panic("coreservices: symbol InvokeResErrUPP not loaded")
	}
	_invokeResErrUPP(arg0, arg1)
}

var _invokeSelectorFunctionUPP func(arg0 uint32, arg1 int32, arg2 SelectorFunctionUPP) int16

// InvokeSelectorFunctionUPP invokes a selector callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1473043-invokeselectorfunctionupp
func InvokeSelectorFunctionUPP(arg0 uint32, arg1 int32, arg2 SelectorFunctionUPP) int16 {
	if _invokeSelectorFunctionUPP == nil {
		panic("coreservices: symbol InvokeSelectorFunctionUPP not loaded")
	}
	return _invokeSelectorFunctionUPP(arg0, arg1, arg2)
}

var _invokeSleepQUPP func(arg0 int, arg1 SleepQRecPtr, arg2 SleepQUPP) int

// InvokeSleepQUPP.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/coreservices/1427428-invokesleepqupp
func InvokeSleepQUPP(arg0 int, arg1 SleepQRecPtr, arg2 SleepQUPP) int {
	if _invokeSleepQUPP == nil {
		panic("coreservices: symbol InvokeSleepQUPP not loaded")
	}
	return _invokeSleepQUPP(arg0, arg1, arg2)
}

var _invokeThreadEntryUPP func(arg0 ThreadEntryUPP) VoidPtr

// InvokeThreadEntryUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574226-invokethreadentryupp
func InvokeThreadEntryUPP(arg0 ThreadEntryUPP) VoidPtr {
	if _invokeThreadEntryUPP == nil {
		panic("coreservices: symbol InvokeThreadEntryUPP not loaded")
	}
	return _invokeThreadEntryUPP(arg0)
}

var _invokeThreadSchedulerUPP func(arg0 SchedulerInfoRecPtr, arg1 ThreadSchedulerUPP) ThreadID

// InvokeThreadSchedulerUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574302-invokethreadschedulerupp
func InvokeThreadSchedulerUPP(arg0 SchedulerInfoRecPtr, arg1 ThreadSchedulerUPP) ThreadID {
	if _invokeThreadSchedulerUPP == nil {
		panic("coreservices: symbol InvokeThreadSchedulerUPP not loaded")
	}
	return _invokeThreadSchedulerUPP(arg0, arg1)
}

var _invokeThreadSwitchUPP func(arg0 ThreadID, arg1 ThreadSwitchUPP)

// InvokeThreadSwitchUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574281-invokethreadswitchupp
func InvokeThreadSwitchUPP(arg0 ThreadID, arg1 ThreadSwitchUPP) {
	if _invokeThreadSwitchUPP == nil {
		panic("coreservices: symbol InvokeThreadSwitchUPP not loaded")
	}
	_invokeThreadSwitchUPP(arg0, arg1)
}

var _invokeThreadTerminationUPP func(arg0 ThreadID, arg1 ThreadTerminationUPP)

// InvokeThreadTerminationUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574290-invokethreadterminationupp
func InvokeThreadTerminationUPP(arg0 ThreadID, arg1 ThreadTerminationUPP) {
	if _invokeThreadTerminationUPP == nil {
		panic("coreservices: symbol InvokeThreadTerminationUPP not loaded")
	}
	_invokeThreadTerminationUPP(arg0, arg1)
}

var _invokeTimerUPP func(arg0 TMTaskPtr, arg1 TimerUPP)

// InvokeTimerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1550780-invoketimerupp
func InvokeTimerUPP(arg0 TMTaskPtr, arg1 TimerUPP) {
	if _invokeTimerUPP == nil {
		panic("coreservices: symbol InvokeTimerUPP not loaded")
	}
	_invokeTimerUPP(arg0, arg1)
}

var _invokeUnicodeToTextFallbackUPP func(arg0 uint16, arg1 uintptr, arg2 uintptr, arg3 TextPtr, arg4 uintptr, arg5 uintptr, arg6 unsafe.Pointer, arg7 ConstUnicodeMappingPtr, arg8 UnicodeToTextFallbackUPP) int32

// InvokeUnicodeToTextFallbackUPP calls your Unicode-to-text fallback callback.
//
// See: https://developer.apple.com/documentation/coreservices/1433599-invokeunicodetotextfallbackupp
func InvokeUnicodeToTextFallbackUPP(arg0 uint16, arg1 uintptr, arg2 uintptr, arg3 TextPtr, arg4 uintptr, arg5 uintptr, arg6 unsafe.Pointer, arg7 ConstUnicodeMappingPtr, arg8 UnicodeToTextFallbackUPP) int32 {
	if _invokeUnicodeToTextFallbackUPP == nil {
		panic("coreservices: symbol InvokeUnicodeToTextFallbackUPP not loaded")
	}
	return _invokeUnicodeToTextFallbackUPP(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _isDataAvailableInIconRef func(inIconKind uint32, inIconRef uintptr) bool

// IsDataAvailableInIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1446627-isdataavailableiniconref
func IsDataAvailableInIconRef(inIconKind uint32, inIconRef uintptr) bool {
	if _isDataAvailableInIconRef == nil {
		panic("coreservices: symbol IsDataAvailableInIconRef not loaded")
	}
	return _isDataAvailableInIconRef(inIconKind, inIconRef)
}

var _isHandleValid func(arg0 uintptr) bool

// IsHandleValid.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506347-ishandlevalid
func IsHandleValid(arg0 uintptr) bool {
	if _isHandleValid == nil {
		panic("coreservices: symbol IsHandleValid not loaded")
	}
	return _isHandleValid(arg0)
}

var _isHeapValid func() bool

// IsHeapValid.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506349-isheapvalid
func IsHeapValid() bool {
	if _isHeapValid == nil {
		panic("coreservices: symbol IsHeapValid not loaded")
	}
	return _isHeapValid()
}

var _isIconRefComposite func(compositeIconRef uintptr, backgroundIconRef uintptr, foregroundIconRef uintptr) int16

// IsIconRefComposite.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1446300-isiconrefcomposite
func IsIconRefComposite(compositeIconRef uintptr, backgroundIconRef uintptr, foregroundIconRef uintptr) int16 {
	if _isIconRefComposite == nil {
		panic("coreservices: symbol IsIconRefComposite not loaded")
	}
	return _isIconRefComposite(compositeIconRef, backgroundIconRef, foregroundIconRef)
}

var _isMetric func() bool

// IsMetric.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1533364-ismetric
func IsMetric() bool {
	if _isMetric == nil {
		panic("coreservices: symbol IsMetric not loaded")
	}
	return _isMetric()
}

var _isPointerValid func(arg0 coreimage.Ptr) bool

// IsPointerValid.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506360-ispointervalid
func IsPointerValid(arg0 coreimage.Ptr) bool {
	if _isPointerValid == nil {
		panic("coreservices: symbol IsPointerValid not loaded")
	}
	return _isPointerValid(arg0)
}

var _isValidIconRef func(theIconRef uintptr) bool

// IsValidIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1450233-isvalidiconref
func IsValidIconRef(theIconRef uintptr) bool {
	if _isValidIconRef == nil {
		panic("coreservices: symbol IsValidIconRef not loaded")
	}
	return _isValidIconRef(theIconRef)
}

var _kCAddCallback func(arg0 KCCallbackUPP, arg1 KCEventMask) int32

// KCAddCallback.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563150-kcaddcallback
func KCAddCallback(arg0 KCCallbackUPP, arg1 KCEventMask) int32 {
	if _kCAddCallback == nil {
		panic("coreservices: symbol KCAddCallback not loaded")
	}
	return _kCAddCallback(arg0, arg1)
}

var _kCCopyItem func(arg0 KCItemRef, arg1 KCRef, arg2 KCItemRef) int32

// KCCopyItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1562984-kccopyitem
func KCCopyItem(arg0 KCItemRef, arg1 KCRef, arg2 KCItemRef) int32 {
	if _kCCopyItem == nil {
		panic("coreservices: symbol KCCopyItem not loaded")
	}
	return _kCCopyItem(arg0, arg1, arg2)
}

var _kCCountKeychains func() uint16

// KCCountKeychains.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563038-kccountkeychains
func KCCountKeychains() uint16 {
	if _kCCountKeychains == nil {
		panic("coreservices: symbol KCCountKeychains not loaded")
	}
	return _kCCountKeychains()
}

var _kCDeleteItem func(arg0 KCItemRef) int32

// KCDeleteItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563195-kcdeleteitem
func KCDeleteItem(arg0 KCItemRef) int32 {
	if _kCDeleteItem == nil {
		panic("coreservices: symbol KCDeleteItem not loaded")
	}
	return _kCDeleteItem(arg0)
}

var _kCFindAppleSharePassword func(arg0 unsafe.Pointer, arg1 string, arg2 string, arg3 string, arg4 string, arg5 uint32, arg6 uint32, arg7 KCItemRef) int32

// KCFindAppleSharePassword.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563124-kcfindapplesharepassword
func KCFindAppleSharePassword(arg0 unsafe.Pointer, arg1 string, arg2 string, arg3 string, arg4 string, arg5 uint32, arg6 uint32, arg7 KCItemRef) int32 {
	if _kCFindAppleSharePassword == nil {
		panic("coreservices: symbol KCFindAppleSharePassword not loaded")
	}
	return _kCFindAppleSharePassword(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _kCFindFirstItem func(arg0 KCRef, arg1 KCAttributeList, arg2 KCSearchRef, arg3 KCItemRef) int32

// KCFindFirstItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563050-kcfindfirstitem
func KCFindFirstItem(arg0 KCRef, arg1 KCAttributeList, arg2 KCSearchRef, arg3 KCItemRef) int32 {
	if _kCFindFirstItem == nil {
		panic("coreservices: symbol KCFindFirstItem not loaded")
	}
	return _kCFindFirstItem(arg0, arg1, arg2, arg3)
}

var _kCFindGenericPassword func(arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 KCItemRef) int32

// KCFindGenericPassword.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563056-kcfindgenericpassword
func KCFindGenericPassword(arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 KCItemRef) int32 {
	if _kCFindGenericPassword == nil {
		panic("coreservices: symbol KCFindGenericPassword not loaded")
	}
	return _kCFindGenericPassword(arg0, arg1, arg2, arg3, arg4)
}

var _kCFindInternetPassword func(arg0 string, arg1 string, arg2 string, arg3 uint16, arg4 uint32, arg5 uint32, arg6 uint32, arg7 uint32, arg8 KCItemRef) int32

// KCFindInternetPassword.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563101-kcfindinternetpassword
func KCFindInternetPassword(arg0 string, arg1 string, arg2 string, arg3 uint16, arg4 uint32, arg5 uint32, arg6 uint32, arg7 uint32, arg8 KCItemRef) int32 {
	if _kCFindInternetPassword == nil {
		panic("coreservices: symbol KCFindInternetPassword not loaded")
	}
	return _kCFindInternetPassword(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _kCFindInternetPasswordWithPath func(arg0 string, arg1 string, arg2 string, arg3 string, arg4 uint16, arg5 uint32, arg6 uint32, arg7 uint32, arg8 uint32, arg9 KCItemRef) int32

// KCFindInternetPasswordWithPath.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563059-kcfindinternetpasswordwithpath
func KCFindInternetPasswordWithPath(arg0 string, arg1 string, arg2 string, arg3 string, arg4 uint16, arg5 uint32, arg6 uint32, arg7 uint32, arg8 uint32, arg9 KCItemRef) int32 {
	if _kCFindInternetPasswordWithPath == nil {
		panic("coreservices: symbol KCFindInternetPasswordWithPath not loaded")
	}
	return _kCFindInternetPasswordWithPath(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

var _kCFindNextItem func(arg0 KCSearchRef, arg1 KCItemRef) int32

// KCFindNextItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563007-kcfindnextitem
func KCFindNextItem(arg0 KCSearchRef, arg1 KCItemRef) int32 {
	if _kCFindNextItem == nil {
		panic("coreservices: symbol KCFindNextItem not loaded")
	}
	return _kCFindNextItem(arg0, arg1)
}

var _kCGetAttribute func(arg0 KCItemRef, arg1 KCAttribute, arg2 uint32) int32

// KCGetAttribute.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563176-kcgetattribute
func KCGetAttribute(arg0 KCItemRef, arg1 KCAttribute, arg2 uint32) int32 {
	if _kCGetAttribute == nil {
		panic("coreservices: symbol KCGetAttribute not loaded")
	}
	return _kCGetAttribute(arg0, arg1, arg2)
}

var _kCGetData func(arg0 KCItemRef, arg1 uint32, arg2 uint32) int32

// KCGetData.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563092-kcgetdata
func KCGetData(arg0 KCItemRef, arg1 uint32, arg2 uint32) int32 {
	if _kCGetData == nil {
		panic("coreservices: symbol KCGetData not loaded")
	}
	return _kCGetData(arg0, arg1, arg2)
}

var _kCGetDefaultKeychain func(arg0 KCRef) int32

// KCGetDefaultKeychain.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563054-kcgetdefaultkeychain
func KCGetDefaultKeychain(arg0 KCRef) int32 {
	if _kCGetDefaultKeychain == nil {
		panic("coreservices: symbol KCGetDefaultKeychain not loaded")
	}
	return _kCGetDefaultKeychain(arg0)
}

var _kCGetIndKeychain func(arg0 uint16, arg1 KCRef) int32

// KCGetIndKeychain.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563047-kcgetindkeychain
func KCGetIndKeychain(arg0 uint16, arg1 KCRef) int32 {
	if _kCGetIndKeychain == nil {
		panic("coreservices: symbol KCGetIndKeychain not loaded")
	}
	return _kCGetIndKeychain(arg0, arg1)
}

var _kCGetKeychain func(arg0 KCItemRef, arg1 KCRef) int32

// KCGetKeychain.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563140-kcgetkeychain
func KCGetKeychain(arg0 KCItemRef, arg1 KCRef) int32 {
	if _kCGetKeychain == nil {
		panic("coreservices: symbol KCGetKeychain not loaded")
	}
	return _kCGetKeychain(arg0, arg1)
}

var _kCGetKeychainManagerVersion func(arg0 uint32) int32

// KCGetKeychainManagerVersion.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1562979-kcgetkeychainmanagerversion
func KCGetKeychainManagerVersion(arg0 uint32) int32 {
	if _kCGetKeychainManagerVersion == nil {
		panic("coreservices: symbol KCGetKeychainManagerVersion not loaded")
	}
	return _kCGetKeychainManagerVersion(arg0)
}

var _kCGetKeychainName func(arg0 KCRef, arg1 string) int32

// KCGetKeychainName.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563052-kcgetkeychainname
func KCGetKeychainName(arg0 KCRef, arg1 string) int32 {
	if _kCGetKeychainName == nil {
		panic("coreservices: symbol KCGetKeychainName not loaded")
	}
	return _kCGetKeychainName(arg0, arg1)
}

var _kCGetStatus func(arg0 KCRef, arg1 uint32) int32

// KCGetStatus.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563057-kcgetstatus
func KCGetStatus(arg0 KCRef, arg1 uint32) int32 {
	if _kCGetStatus == nil {
		panic("coreservices: symbol KCGetStatus not loaded")
	}
	return _kCGetStatus(arg0, arg1)
}

var _kCIsInteractionAllowed func() bool

// KCIsInteractionAllowed.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563065-kcisinteractionallowed
func KCIsInteractionAllowed() bool {
	if _kCIsInteractionAllowed == nil {
		panic("coreservices: symbol KCIsInteractionAllowed not loaded")
	}
	return _kCIsInteractionAllowed()
}

var _kCLock func(arg0 KCRef) int32

// KCLock.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563012-kclock
func KCLock(arg0 KCRef) int32 {
	if _kCLock == nil {
		panic("coreservices: symbol KCLock not loaded")
	}
	return _kCLock(arg0)
}

var _kCMakeAliasFromKCRef func(arg0 KCRef, arg1 AliasHandle) int32

// KCMakeAliasFromKCRef.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563163-kcmakealiasfromkcref
func KCMakeAliasFromKCRef(arg0 KCRef, arg1 AliasHandle) int32 {
	if _kCMakeAliasFromKCRef == nil {
		panic("coreservices: symbol KCMakeAliasFromKCRef not loaded")
	}
	return _kCMakeAliasFromKCRef(arg0, arg1)
}

var _kCMakeKCRefFromAlias func(arg0 AliasHandle, arg1 KCRef) int32

// KCMakeKCRefFromAlias.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563083-kcmakekcreffromalias
func KCMakeKCRefFromAlias(arg0 AliasHandle, arg1 KCRef) int32 {
	if _kCMakeKCRefFromAlias == nil {
		panic("coreservices: symbol KCMakeKCRefFromAlias not loaded")
	}
	return _kCMakeKCRefFromAlias(arg0, arg1)
}

var _kCMakeKCRefFromFSRef func(arg0 uintptr, arg1 KCRef) int32

// KCMakeKCRefFromFSRef.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563116-kcmakekcreffromfsref
func KCMakeKCRefFromFSRef(arg0 uintptr, arg1 KCRef) int32 {
	if _kCMakeKCRefFromFSRef == nil {
		panic("coreservices: symbol KCMakeKCRefFromFSRef not loaded")
	}
	return _kCMakeKCRefFromFSRef(arg0, arg1)
}

var _kCNewItem func(arg0 KCItemClass, arg1 uint32, arg2 uint32, arg3 KCItemRef) int32

// KCNewItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563010-kcnewitem
func KCNewItem(arg0 KCItemClass, arg1 uint32, arg2 uint32, arg3 KCItemRef) int32 {
	if _kCNewItem == nil {
		panic("coreservices: symbol KCNewItem not loaded")
	}
	return _kCNewItem(arg0, arg1, arg2, arg3)
}

var _kCReleaseItem func(arg0 KCItemRef) int32

// KCReleaseItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563161-kcreleaseitem
func KCReleaseItem(arg0 KCItemRef) int32 {
	if _kCReleaseItem == nil {
		panic("coreservices: symbol KCReleaseItem not loaded")
	}
	return _kCReleaseItem(arg0)
}

var _kCReleaseKeychain func(arg0 KCRef) int32

// KCReleaseKeychain.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563130-kcreleasekeychain
func KCReleaseKeychain(arg0 KCRef) int32 {
	if _kCReleaseKeychain == nil {
		panic("coreservices: symbol KCReleaseKeychain not loaded")
	}
	return _kCReleaseKeychain(arg0)
}

var _kCReleaseSearch func(arg0 KCSearchRef) int32

// KCReleaseSearch.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1562982-kcreleasesearch
func KCReleaseSearch(arg0 KCSearchRef) int32 {
	if _kCReleaseSearch == nil {
		panic("coreservices: symbol KCReleaseSearch not loaded")
	}
	return _kCReleaseSearch(arg0)
}

var _kCRemoveCallback func(arg0 KCCallbackUPP) int32

// KCRemoveCallback.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563066-kcremovecallback
func KCRemoveCallback(arg0 KCCallbackUPP) int32 {
	if _kCRemoveCallback == nil {
		panic("coreservices: symbol KCRemoveCallback not loaded")
	}
	return _kCRemoveCallback(arg0)
}

var _kCSetAttribute func(arg0 KCItemRef, arg1 KCAttribute) int32

// KCSetAttribute.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563089-kcsetattribute
func KCSetAttribute(arg0 KCItemRef, arg1 KCAttribute) int32 {
	if _kCSetAttribute == nil {
		panic("coreservices: symbol KCSetAttribute not loaded")
	}
	return _kCSetAttribute(arg0, arg1)
}

var _kCSetData func(arg0 KCItemRef, arg1 uint32) int32

// KCSetData.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563018-kcsetdata
func KCSetData(arg0 KCItemRef, arg1 uint32) int32 {
	if _kCSetData == nil {
		panic("coreservices: symbol KCSetData not loaded")
	}
	return _kCSetData(arg0, arg1)
}

var _kCSetDefaultKeychain func(arg0 KCRef) int32

// KCSetDefaultKeychain.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563084-kcsetdefaultkeychain
func KCSetDefaultKeychain(arg0 KCRef) int32 {
	if _kCSetDefaultKeychain == nil {
		panic("coreservices: symbol KCSetDefaultKeychain not loaded")
	}
	return _kCSetDefaultKeychain(arg0)
}

var _kCSetInteractionAllowed func(arg0 bool) int32

// KCSetInteractionAllowed.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563079-kcsetinteractionallowed
func KCSetInteractionAllowed(arg0 bool) int32 {
	if _kCSetInteractionAllowed == nil {
		panic("coreservices: symbol KCSetInteractionAllowed not loaded")
	}
	return _kCSetInteractionAllowed(arg0)
}

var _kCUpdateItem func(arg0 KCItemRef) int32

// KCUpdateItem.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563138-kcupdateitem
func KCUpdateItem(arg0 KCItemRef) int32 {
	if _kCUpdateItem == nil {
		panic("coreservices: symbol KCUpdateItem not loaded")
	}
	return _kCUpdateItem(arg0)
}

var _lMGetApFontID func() int16

// LMGetApFontID.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1564978-lmgetapfontid
func LMGetApFontID() int16 {
	if _lMGetApFontID == nil {
		panic("coreservices: symbol LMGetApFontID not loaded")
	}
	return _lMGetApFontID()
}

var _lMGetBootDrive func() int16

// LMGetBootDrive.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1564982-lmgetbootdrive
func LMGetBootDrive() int16 {
	if _lMGetBootDrive == nil {
		panic("coreservices: symbol LMGetBootDrive not loaded")
	}
	return _lMGetBootDrive()
}

var _lMGetIntlSpec func() coreimage.Ptr

// LMGetIntlSpec.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565043-lmgetintlspec
func LMGetIntlSpec() coreimage.Ptr {
	if _lMGetIntlSpec == nil {
		panic("coreservices: symbol LMGetIntlSpec not loaded")
	}
	return _lMGetIntlSpec()
}

var _lMGetMemErr func() int16

// LMGetMemErr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506460-lmgetmemerr
func LMGetMemErr() int16 {
	if _lMGetMemErr == nil {
		panic("coreservices: symbol LMGetMemErr not loaded")
	}
	return _lMGetMemErr()
}

var _lMGetResErr func() int16

// LMGetResErr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565008-lmgetreserr
func LMGetResErr() int16 {
	if _lMGetResErr == nil {
		panic("coreservices: symbol LMGetResErr not loaded")
	}
	return _lMGetResErr()
}

var _lMGetResLoad func() uint8

// LMGetResLoad.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565056-lmgetresload
func LMGetResLoad() uint8 {
	if _lMGetResLoad == nil {
		panic("coreservices: symbol LMGetResLoad not loaded")
	}
	return _lMGetResLoad()
}

var _lMGetSysFontSize func() int16

// LMGetSysFontSize.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1565057-lmgetsysfontsize
func LMGetSysFontSize() int16 {
	if _lMGetSysFontSize == nil {
		panic("coreservices: symbol LMGetSysFontSize not loaded")
	}
	return _lMGetSysFontSize()
}

var _lMGetSysMap func() int16

// LMGetSysMap.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565032-lmgetsysmap
func LMGetSysMap() int16 {
	if _lMGetSysMap == nil {
		panic("coreservices: symbol LMGetSysMap not loaded")
	}
	return _lMGetSysMap()
}

var _lMGetTmpResLoad func() uint8

// LMGetTmpResLoad.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1564999-lmgettmpresload
func LMGetTmpResLoad() uint8 {
	if _lMGetTmpResLoad == nil {
		panic("coreservices: symbol LMGetTmpResLoad not loaded")
	}
	return _lMGetTmpResLoad()
}

var _lMSetApFontID func(arg0 int16)

// LMSetApFontID.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1565011-lmsetapfontid
func LMSetApFontID(arg0 int16) {
	if _lMSetApFontID == nil {
		panic("coreservices: symbol LMSetApFontID not loaded")
	}
	_lMSetApFontID(arg0)
}

var _lMSetBootDrive func(arg0 int16)

// LMSetBootDrive.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1564973-lmsetbootdrive
func LMSetBootDrive(arg0 int16) {
	if _lMSetBootDrive == nil {
		panic("coreservices: symbol LMSetBootDrive not loaded")
	}
	_lMSetBootDrive(arg0)
}

var _lMSetIntlSpec func(arg0 coreimage.Ptr)

// LMSetIntlSpec.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565016-lmsetintlspec
func LMSetIntlSpec(arg0 coreimage.Ptr) {
	if _lMSetIntlSpec == nil {
		panic("coreservices: symbol LMSetIntlSpec not loaded")
	}
	_lMSetIntlSpec(arg0)
}

var _lMSetMemErr func(arg0 int16)

// LMSetMemErr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506279-lmsetmemerr
func LMSetMemErr(arg0 int16) {
	if _lMSetMemErr == nil {
		panic("coreservices: symbol LMSetMemErr not loaded")
	}
	_lMSetMemErr(arg0)
}

var _lMSetResErr func(arg0 int16)

// LMSetResErr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565000-lmsetreserr
func LMSetResErr(arg0 int16) {
	if _lMSetResErr == nil {
		panic("coreservices: symbol LMSetResErr not loaded")
	}
	_lMSetResErr(arg0)
}

var _lMSetResLoad func(arg0 uint8)

// LMSetResLoad.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565052-lmsetresload
func LMSetResLoad(arg0 uint8) {
	if _lMSetResLoad == nil {
		panic("coreservices: symbol LMSetResLoad not loaded")
	}
	_lMSetResLoad(arg0)
}

var _lMSetSysFontFam func(arg0 int16)

// LMSetSysFontFam.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1565044-lmsetsysfontfam
func LMSetSysFontFam(arg0 int16) {
	if _lMSetSysFontFam == nil {
		panic("coreservices: symbol LMSetSysFontFam not loaded")
	}
	_lMSetSysFontFam(arg0)
}

var _lMSetSysFontSize func(arg0 int16)

// LMSetSysFontSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565048-lmsetsysfontsize
func LMSetSysFontSize(arg0 int16) {
	if _lMSetSysFontSize == nil {
		panic("coreservices: symbol LMSetSysFontSize not loaded")
	}
	_lMSetSysFontSize(arg0)
}

var _lMSetSysMap func(arg0 int16)

// LMSetSysMap.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565018-lmsetsysmap
func LMSetSysMap(arg0 int16) {
	if _lMSetSysMap == nil {
		panic("coreservices: symbol LMSetSysMap not loaded")
	}
	_lMSetSysMap(arg0)
}

var _lMSetTmpResLoad func(arg0 uint8)

// LMSetTmpResLoad.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565025-lmsettmpresload
func LMSetTmpResLoad(arg0 uint8) {
	if _lMSetTmpResLoad == nil {
		panic("coreservices: symbol LMSetTmpResLoad not loaded")
	}
	_lMSetTmpResLoad(arg0)
}

var _lSCanURLAcceptURL func(inItemURL corefoundation.CFURLRef, inTargetURL corefoundation.CFURLRef, inRoleMask LSRolesMask, inFlags LSAcceptanceFlags, outAcceptsItem unsafe.Pointer) int32

// LSCanURLAcceptURL tests whether an app can accept (open) an item for a URL.
//
// See: https://developer.apple.com/documentation/coreservices/1441854-lscanurlaccepturl
func LSCanURLAcceptURL(inItemURL corefoundation.CFURLRef, inTargetURL corefoundation.CFURLRef, inRoleMask LSRolesMask, inFlags LSAcceptanceFlags, outAcceptsItem unsafe.Pointer) int32 {
	if _lSCanURLAcceptURL == nil {
		panic("coreservices: symbol LSCanURLAcceptURL not loaded")
	}
	return _lSCanURLAcceptURL(inItemURL, inTargetURL, inRoleMask, inFlags, outAcceptsItem)
}

var _lSCopyAllHandlersForURLScheme func(inURLScheme corefoundation.CFStringRef) corefoundation.CFArrayRef

// LSCopyAllHandlersForURLScheme locates app bundle identifiers for apps capable of handling the specified URL scheme.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1443240-lscopyallhandlersforurlscheme
func LSCopyAllHandlersForURLScheme(inURLScheme corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _lSCopyAllHandlersForURLScheme == nil {
		panic("coreservices: symbol LSCopyAllHandlersForURLScheme not loaded")
	}
	return _lSCopyAllHandlersForURLScheme(inURLScheme)
}

var _lSCopyAllRoleHandlersForContentType func(inContentType corefoundation.CFStringRef, inRole LSRolesMask) corefoundation.CFArrayRef

// LSCopyAllRoleHandlersForContentType locates an array of bundle identifiers for apps capable of handling a specified content type with the specified roles.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1448020-lscopyallrolehandlersforcontentt
func LSCopyAllRoleHandlersForContentType(inContentType corefoundation.CFStringRef, inRole LSRolesMask) corefoundation.CFArrayRef {
	if _lSCopyAllRoleHandlersForContentType == nil {
		panic("coreservices: symbol LSCopyAllRoleHandlersForContentType not loaded")
	}
	return _lSCopyAllRoleHandlersForContentType(inContentType, inRole)
}

var _lSCopyApplicationURLsForBundleIdentifier func(inBundleIdentifier corefoundation.CFStringRef, outError *corefoundation.CFErrorRef) corefoundation.CFArrayRef

// LSCopyApplicationURLsForBundleIdentifier locates all URLs for apps that correspond to the specified bundle identifier.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1449290-lscopyapplicationurlsforbundleid
func LSCopyApplicationURLsForBundleIdentifier(inBundleIdentifier corefoundation.CFStringRef, outError *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	if _lSCopyApplicationURLsForBundleIdentifier == nil {
		panic("coreservices: symbol LSCopyApplicationURLsForBundleIdentifier not loaded")
	}
	return _lSCopyApplicationURLsForBundleIdentifier(inBundleIdentifier, outError)
}

var _lSCopyApplicationURLsForURL func(inURL corefoundation.CFURLRef, inRoleMask LSRolesMask) corefoundation.CFArrayRef

// LSCopyApplicationURLsForURL locates all known apps suitable for opening an item for the specified URL.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1445148-lscopyapplicationurlsforurl
func LSCopyApplicationURLsForURL(inURL corefoundation.CFURLRef, inRoleMask LSRolesMask) corefoundation.CFArrayRef {
	if _lSCopyApplicationURLsForURL == nil {
		panic("coreservices: symbol LSCopyApplicationURLsForURL not loaded")
	}
	return _lSCopyApplicationURLsForURL(inURL, inRoleMask)
}

var _lSCopyDefaultApplicationURLForContentType func(inContentType corefoundation.CFStringRef, inRoleMask LSRolesMask, outError *corefoundation.CFErrorRef) corefoundation.CFURLRef

// LSCopyDefaultApplicationURLForContentType returns the app that opens a content type.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1447734-lscopydefaultapplicationurlforco
func LSCopyDefaultApplicationURLForContentType(inContentType corefoundation.CFStringRef, inRoleMask LSRolesMask, outError *corefoundation.CFErrorRef) corefoundation.CFURLRef {
	if _lSCopyDefaultApplicationURLForContentType == nil {
		panic("coreservices: symbol LSCopyDefaultApplicationURLForContentType not loaded")
	}
	return _lSCopyDefaultApplicationURLForContentType(inContentType, inRoleMask, outError)
}

var _lSCopyDefaultApplicationURLForURL func(inURL corefoundation.CFURLRef, inRoleMask LSRolesMask, outError *corefoundation.CFErrorRef) corefoundation.CFURLRef

// LSCopyDefaultApplicationURLForURL returns the app that opens an item.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1448824-lscopydefaultapplicationurlforur
func LSCopyDefaultApplicationURLForURL(inURL corefoundation.CFURLRef, inRoleMask LSRolesMask, outError *corefoundation.CFErrorRef) corefoundation.CFURLRef {
	if _lSCopyDefaultApplicationURLForURL == nil {
		panic("coreservices: symbol LSCopyDefaultApplicationURLForURL not loaded")
	}
	return _lSCopyDefaultApplicationURLForURL(inURL, inRoleMask, outError)
}

var _lSCopyDefaultHandlerForURLScheme func(inURLScheme corefoundation.CFStringRef) corefoundation.CFStringRef

// LSCopyDefaultHandlerForURLScheme returns the bundle identifier of the user’s preferred default handler for the specified URL scheme.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1441725-lscopydefaulthandlerforurlscheme
func LSCopyDefaultHandlerForURLScheme(inURLScheme corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _lSCopyDefaultHandlerForURLScheme == nil {
		panic("coreservices: symbol LSCopyDefaultHandlerForURLScheme not loaded")
	}
	return _lSCopyDefaultHandlerForURLScheme(inURLScheme)
}

var _lSCopyDefaultRoleHandlerForContentType func(inContentType corefoundation.CFStringRef, inRole LSRolesMask) corefoundation.CFStringRef

// LSCopyDefaultRoleHandlerForContentType returns the bundle identifier of the user’s preferred default handler for the specified content type with the specified role.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1449868-lscopydefaultrolehandlerforconte
func LSCopyDefaultRoleHandlerForContentType(inContentType corefoundation.CFStringRef, inRole LSRolesMask) corefoundation.CFStringRef {
	if _lSCopyDefaultRoleHandlerForContentType == nil {
		panic("coreservices: symbol LSCopyDefaultRoleHandlerForContentType not loaded")
	}
	return _lSCopyDefaultRoleHandlerForContentType(inContentType, inRole)
}

var _lSCopyDisplayNameForURL func(inURL corefoundation.CFURLRef, outDisplayName *corefoundation.CFStringRef) int32

// LSCopyDisplayNameForURL obtains the display name for an item with a URL.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1446850-lscopydisplaynameforurl
func LSCopyDisplayNameForURL(inURL corefoundation.CFURLRef, outDisplayName *corefoundation.CFStringRef) int32 {
	if _lSCopyDisplayNameForURL == nil {
		panic("coreservices: symbol LSCopyDisplayNameForURL not loaded")
	}
	return _lSCopyDisplayNameForURL(inURL, outDisplayName)
}

var _lSCopyItemInfoForURL func(inURL corefoundation.CFURLRef, inWhichInfo LSRequestedInfo, outItemInfo unsafe.Pointer) int32

// LSCopyItemInfoForURL obtains requested information about an item with a URL.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1445685-lscopyiteminfoforurl
func LSCopyItemInfoForURL(inURL corefoundation.CFURLRef, inWhichInfo LSRequestedInfo, outItemInfo unsafe.Pointer) int32 {
	if _lSCopyItemInfoForURL == nil {
		panic("coreservices: symbol LSCopyItemInfoForURL not loaded")
	}
	return _lSCopyItemInfoForURL(inURL, inWhichInfo, outItemInfo)
}

var _lSCopyKindStringForURL func(inURL corefoundation.CFURLRef, outKindString *corefoundation.CFStringRef) int32

// LSCopyKindStringForURL obtains the kind string for an item with a URL.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1447481-lscopykindstringforurl
func LSCopyKindStringForURL(inURL corefoundation.CFURLRef, outKindString *corefoundation.CFStringRef) int32 {
	if _lSCopyKindStringForURL == nil {
		panic("coreservices: symbol LSCopyKindStringForURL not loaded")
	}
	return _lSCopyKindStringForURL(inURL, outKindString)
}

var _lSGetExtensionInfo func(inNameLen unsafe.Pointer, inNameBuffer *uint16, outExtStartIndex unsafe.Pointer) int32

// LSGetExtensionInfo obtains the starting index of the extension within a filename.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1446043-lsgetextensioninfo
func LSGetExtensionInfo(inNameLen unsafe.Pointer, inNameBuffer *uint16, outExtStartIndex unsafe.Pointer) int32 {
	if _lSGetExtensionInfo == nil {
		panic("coreservices: symbol LSGetExtensionInfo not loaded")
	}
	return _lSGetExtensionInfo(inNameLen, inNameBuffer, outExtStartIndex)
}

var _lSGetHandlerOptionsForContentType func(inContentType corefoundation.CFStringRef) LSHandlerOptions

// LSGetHandlerOptionsForContentType gets the handler options for the specified content type.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1445296-lsgethandleroptionsforcontenttyp
func LSGetHandlerOptionsForContentType(inContentType corefoundation.CFStringRef) LSHandlerOptions {
	if _lSGetHandlerOptionsForContentType == nil {
		panic("coreservices: symbol LSGetHandlerOptionsForContentType not loaded")
	}
	return _lSGetHandlerOptionsForContentType(inContentType)
}

var _lSOpenApplication func(appParams unsafe.Pointer, outPSN unsafe.Pointer) int32

// LSOpenApplication launches the specified app.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1447930-lsopenapplication
func LSOpenApplication(appParams unsafe.Pointer, outPSN unsafe.Pointer) int32 {
	if _lSOpenApplication == nil {
		panic("coreservices: symbol LSOpenApplication not loaded")
	}
	return _lSOpenApplication(appParams, outPSN)
}

var _lSOpenCFURLRef func(inURL corefoundation.CFURLRef, outLaunchedURL *corefoundation.CFURLRef) int32

// LSOpenCFURLRef opens an item for a URL in the default manner in its preferred app.
//
// See: https://developer.apple.com/documentation/coreservices/1442850-lsopencfurlref
func LSOpenCFURLRef(inURL corefoundation.CFURLRef, outLaunchedURL *corefoundation.CFURLRef) int32 {
	if _lSOpenCFURLRef == nil {
		panic("coreservices: symbol LSOpenCFURLRef not loaded")
	}
	return _lSOpenCFURLRef(inURL, outLaunchedURL)
}

var _lSOpenFSRef func(inRef unsafe.Pointer, outLaunchedRef unsafe.Pointer) int32

// LSOpenFSRef opens an item with a file-system reference in the default manner in its preferred app.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1445663-lsopenfsref
func LSOpenFSRef(inRef unsafe.Pointer, outLaunchedRef unsafe.Pointer) int32 {
	if _lSOpenFSRef == nil {
		panic("coreservices: symbol LSOpenFSRef not loaded")
	}
	return _lSOpenFSRef(inRef, outLaunchedRef)
}

var _lSOpenFromRefSpec func(inLaunchSpec unsafe.Pointer, outLaunchedRef unsafe.Pointer) int32

// LSOpenFromRefSpec opens one or more items with a file-system reference in either their preferred apps or a designated app.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1444466-lsopenfromrefspec
func LSOpenFromRefSpec(inLaunchSpec unsafe.Pointer, outLaunchedRef unsafe.Pointer) int32 {
	if _lSOpenFromRefSpec == nil {
		panic("coreservices: symbol LSOpenFromRefSpec not loaded")
	}
	return _lSOpenFromRefSpec(inLaunchSpec, outLaunchedRef)
}

var _lSOpenFromURLSpec func(inLaunchSpec unsafe.Pointer, outLaunchedURL *corefoundation.CFURLRef) int32

// LSOpenFromURLSpec opens one or more items for a URL in the preferred apps or a designated app.
//
// See: https://developer.apple.com/documentation/coreservices/1441986-lsopenfromurlspec
func LSOpenFromURLSpec(inLaunchSpec unsafe.Pointer, outLaunchedURL *corefoundation.CFURLRef) int32 {
	if _lSOpenFromURLSpec == nil {
		panic("coreservices: symbol LSOpenFromURLSpec not loaded")
	}
	return _lSOpenFromURLSpec(inLaunchSpec, outLaunchedURL)
}

var _lSOpenItemsWithRole func(inItems unsafe.Pointer, inItemCount int, inRole LSRolesMask, inAEParam unsafe.Pointer, inAppParams unsafe.Pointer, outPSNs unsafe.Pointer, inMaxPSNCount int) int32

// LSOpenItemsWithRole opens items with an array of file-system references with a specified role.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1449783-lsopenitemswithrole
func LSOpenItemsWithRole(inItems unsafe.Pointer, inItemCount int, inRole LSRolesMask, inAEParam unsafe.Pointer, inAppParams unsafe.Pointer, outPSNs unsafe.Pointer, inMaxPSNCount int) int32 {
	if _lSOpenItemsWithRole == nil {
		panic("coreservices: symbol LSOpenItemsWithRole not loaded")
	}
	return _lSOpenItemsWithRole(inItems, inItemCount, inRole, inAEParam, inAppParams, outPSNs, inMaxPSNCount)
}

var _lSOpenURLsWithRole func(inURLs corefoundation.CFArrayRef, inRole LSRolesMask, inAEParam unsafe.Pointer, inAppParams unsafe.Pointer, outPSNs unsafe.Pointer, inMaxPSNCount int) int32

// LSOpenURLsWithRole opens one or more URLs with the specified roles.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1448184-lsopenurlswithrole
func LSOpenURLsWithRole(inURLs corefoundation.CFArrayRef, inRole LSRolesMask, inAEParam unsafe.Pointer, inAppParams unsafe.Pointer, outPSNs unsafe.Pointer, inMaxPSNCount int) int32 {
	if _lSOpenURLsWithRole == nil {
		panic("coreservices: symbol LSOpenURLsWithRole not loaded")
	}
	return _lSOpenURLsWithRole(inURLs, inRole, inAEParam, inAppParams, outPSNs, inMaxPSNCount)
}

var _lSRegisterURL func(inURL corefoundation.CFURLRef, inUpdate unsafe.Pointer) int32

// LSRegisterURL registers an app, using a URL, in the Launch Services database.
//
// See: https://developer.apple.com/documentation/coreservices/1446350-lsregisterurl
func LSRegisterURL(inURL corefoundation.CFURLRef, inUpdate unsafe.Pointer) int32 {
	if _lSRegisterURL == nil {
		panic("coreservices: symbol LSRegisterURL not loaded")
	}
	return _lSRegisterURL(inURL, inUpdate)
}

var _lSSetDefaultHandlerForURLScheme func(inURLScheme corefoundation.CFStringRef, inHandlerBundleID corefoundation.CFStringRef) int32

// LSSetDefaultHandlerForURLScheme sets the user’s preferred default handler for the specified URL scheme.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1447760-lssetdefaulthandlerforurlscheme
func LSSetDefaultHandlerForURLScheme(inURLScheme corefoundation.CFStringRef, inHandlerBundleID corefoundation.CFStringRef) int32 {
	if _lSSetDefaultHandlerForURLScheme == nil {
		panic("coreservices: symbol LSSetDefaultHandlerForURLScheme not loaded")
	}
	return _lSSetDefaultHandlerForURLScheme(inURLScheme, inHandlerBundleID)
}

var _lSSetDefaultRoleHandlerForContentType func(inContentType corefoundation.CFStringRef, inRole LSRolesMask, inHandlerBundleID corefoundation.CFStringRef) int32

// LSSetDefaultRoleHandlerForContentType sets the user’s preferred default handler for the specified content type in the specified roles.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1444955-lssetdefaultrolehandlerforconten
func LSSetDefaultRoleHandlerForContentType(inContentType corefoundation.CFStringRef, inRole LSRolesMask, inHandlerBundleID corefoundation.CFStringRef) int32 {
	if _lSSetDefaultRoleHandlerForContentType == nil {
		panic("coreservices: symbol LSSetDefaultRoleHandlerForContentType not loaded")
	}
	return _lSSetDefaultRoleHandlerForContentType(inContentType, inRole, inHandlerBundleID)
}

var _lSSetExtensionHiddenForURL func(inURL corefoundation.CFURLRef, inHide unsafe.Pointer) int32

// LSSetExtensionHiddenForURL specifies whether to show or hide the filename extension for an item with a URL.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1443948-lssetextensionhiddenforurl
func LSSetExtensionHiddenForURL(inURL corefoundation.CFURLRef, inHide unsafe.Pointer) int32 {
	if _lSSetExtensionHiddenForURL == nil {
		panic("coreservices: symbol LSSetExtensionHiddenForURL not loaded")
	}
	return _lSSetExtensionHiddenForURL(inURL, inHide)
}

var _lSSetHandlerOptionsForContentType func(inContentType corefoundation.CFStringRef, inOptions LSHandlerOptions) int32

// LSSetHandlerOptionsForContentType sets the handler option for the specified content type.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1447588-lssethandleroptionsforcontenttyp
func LSSetHandlerOptionsForContentType(inContentType corefoundation.CFStringRef, inOptions LSHandlerOptions) int32 {
	if _lSSetHandlerOptionsForContentType == nil {
		panic("coreservices: symbol LSSetHandlerOptionsForContentType not loaded")
	}
	return _lSSetHandlerOptionsForContentType(inContentType, inOptions)
}

var _lSSetItemAttribute func(inItem unsafe.Pointer, inRoles LSRolesMask, inAttributeName corefoundation.CFStringRef, inValue corefoundation.CFTypeRef) int32

// LSSetItemAttribute.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1446733-lssetitemattribute
func LSSetItemAttribute(inItem unsafe.Pointer, inRoles LSRolesMask, inAttributeName corefoundation.CFStringRef, inValue corefoundation.CFTypeRef) int32 {
	if _lSSetItemAttribute == nil {
		panic("coreservices: symbol LSSetItemAttribute not loaded")
	}
	return _lSSetItemAttribute(inItem, inRoles, inAttributeName, inValue)
}

var _lSSharedFileListAddObserver func(inList LSSharedFileListRef, inRunloop corefoundation.CFRunLoopRef, inRunloopMode corefoundation.CFStringRef, callback LSSharedFileListChangedProcPtr, context unsafe.Pointer)

// LSSharedFileListAddObserver.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1445770-lssharedfilelistaddobserver
func LSSharedFileListAddObserver(inList LSSharedFileListRef, inRunloop corefoundation.CFRunLoopRef, inRunloopMode corefoundation.CFStringRef, callback LSSharedFileListChangedProcPtr, context unsafe.Pointer) {
	if _lSSharedFileListAddObserver == nil {
		panic("coreservices: symbol LSSharedFileListAddObserver not loaded")
	}
	_lSSharedFileListAddObserver(inList, inRunloop, inRunloopMode, callback, context)
}

var _lSSharedFileListCopyProperty func(inList LSSharedFileListRef, inPropertyName corefoundation.CFStringRef) corefoundation.CFTypeRef

// LSSharedFileListCopyProperty.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1444588-lssharedfilelistcopyproperty
func LSSharedFileListCopyProperty(inList LSSharedFileListRef, inPropertyName corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _lSSharedFileListCopyProperty == nil {
		panic("coreservices: symbol LSSharedFileListCopyProperty not loaded")
	}
	return _lSSharedFileListCopyProperty(inList, inPropertyName)
}

var _lSSharedFileListCopySnapshot func(inList LSSharedFileListRef, outSnapshotSeed *uint32) corefoundation.CFArrayRef

// LSSharedFileListCopySnapshot.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1448112-lssharedfilelistcopysnapshot
func LSSharedFileListCopySnapshot(inList LSSharedFileListRef, outSnapshotSeed *uint32) corefoundation.CFArrayRef {
	if _lSSharedFileListCopySnapshot == nil {
		panic("coreservices: symbol LSSharedFileListCopySnapshot not loaded")
	}
	return _lSSharedFileListCopySnapshot(inList, outSnapshotSeed)
}

var _lSSharedFileListCreate func(inAllocator corefoundation.CFAllocatorRef, inListType corefoundation.CFStringRef, listOptions corefoundation.CFTypeRef) LSSharedFileListRef

// LSSharedFileListCreate.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1443926-lssharedfilelistcreate
func LSSharedFileListCreate(inAllocator corefoundation.CFAllocatorRef, inListType corefoundation.CFStringRef, listOptions corefoundation.CFTypeRef) LSSharedFileListRef {
	if _lSSharedFileListCreate == nil {
		panic("coreservices: symbol LSSharedFileListCreate not loaded")
	}
	return _lSSharedFileListCreate(inAllocator, inListType, listOptions)
}

var _lSSharedFileListGetSeedValue func(inList LSSharedFileListRef) uint32

// LSSharedFileListGetSeedValue.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1444885-lssharedfilelistgetseedvalue
func LSSharedFileListGetSeedValue(inList LSSharedFileListRef) uint32 {
	if _lSSharedFileListGetSeedValue == nil {
		panic("coreservices: symbol LSSharedFileListGetSeedValue not loaded")
	}
	return _lSSharedFileListGetSeedValue(inList)
}

var _lSSharedFileListGetTypeID func() uint

// LSSharedFileListGetTypeID.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1450618-lssharedfilelistgettypeid
func LSSharedFileListGetTypeID() uint {
	if _lSSharedFileListGetTypeID == nil {
		panic("coreservices: symbol LSSharedFileListGetTypeID not loaded")
	}
	return _lSSharedFileListGetTypeID()
}

var _lSSharedFileListInsertItemFSRef func(inList LSSharedFileListRef, insertAfterThisItem LSSharedFileListItemRef, inDisplayName corefoundation.CFStringRef, inIconRef uintptr, inFSRef unsafe.Pointer, inPropertiesToSet corefoundation.CFDictionaryRef, inPropertiesToClear corefoundation.CFArrayRef) LSSharedFileListItemRef

// LSSharedFileListInsertItemFSRef.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1449884-lssharedfilelistinsertitemfsref
func LSSharedFileListInsertItemFSRef(inList LSSharedFileListRef, insertAfterThisItem LSSharedFileListItemRef, inDisplayName corefoundation.CFStringRef, inIconRef uintptr, inFSRef unsafe.Pointer, inPropertiesToSet corefoundation.CFDictionaryRef, inPropertiesToClear corefoundation.CFArrayRef) LSSharedFileListItemRef {
	if _lSSharedFileListInsertItemFSRef == nil {
		panic("coreservices: symbol LSSharedFileListInsertItemFSRef not loaded")
	}
	return _lSSharedFileListInsertItemFSRef(inList, insertAfterThisItem, inDisplayName, inIconRef, inFSRef, inPropertiesToSet, inPropertiesToClear)
}

var _lSSharedFileListInsertItemURL func(inList LSSharedFileListRef, insertAfterThisItem LSSharedFileListItemRef, inDisplayName corefoundation.CFStringRef, inIconRef uintptr, inURL corefoundation.CFURLRef, inPropertiesToSet corefoundation.CFDictionaryRef, inPropertiesToClear corefoundation.CFArrayRef) LSSharedFileListItemRef

// LSSharedFileListInsertItemURL.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1444471-lssharedfilelistinsertitemurl
func LSSharedFileListInsertItemURL(inList LSSharedFileListRef, insertAfterThisItem LSSharedFileListItemRef, inDisplayName corefoundation.CFStringRef, inIconRef uintptr, inURL corefoundation.CFURLRef, inPropertiesToSet corefoundation.CFDictionaryRef, inPropertiesToClear corefoundation.CFArrayRef) LSSharedFileListItemRef {
	if _lSSharedFileListInsertItemURL == nil {
		panic("coreservices: symbol LSSharedFileListInsertItemURL not loaded")
	}
	return _lSSharedFileListInsertItemURL(inList, insertAfterThisItem, inDisplayName, inIconRef, inURL, inPropertiesToSet, inPropertiesToClear)
}

var _lSSharedFileListItemCopyDisplayName func(inItem LSSharedFileListItemRef) corefoundation.CFStringRef

// LSSharedFileListItemCopyDisplayName.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1449716-lssharedfilelistitemcopydisplayn
func LSSharedFileListItemCopyDisplayName(inItem LSSharedFileListItemRef) corefoundation.CFStringRef {
	if _lSSharedFileListItemCopyDisplayName == nil {
		panic("coreservices: symbol LSSharedFileListItemCopyDisplayName not loaded")
	}
	return _lSSharedFileListItemCopyDisplayName(inItem)
}

var _lSSharedFileListItemCopyIconRef func(inItem LSSharedFileListItemRef) objectivec.IObject

// LSSharedFileListItemCopyIconRef.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1442889-lssharedfilelistitemcopyiconref
func LSSharedFileListItemCopyIconRef(inItem LSSharedFileListItemRef) objectivec.IObject {
	if _lSSharedFileListItemCopyIconRef == nil {
		panic("coreservices: symbol LSSharedFileListItemCopyIconRef not loaded")
	}
	return _lSSharedFileListItemCopyIconRef(inItem)
}

var _lSSharedFileListItemCopyProperty func(inItem LSSharedFileListItemRef, inPropertyName corefoundation.CFStringRef) corefoundation.CFTypeRef

// LSSharedFileListItemCopyProperty.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1445074-lssharedfilelistitemcopyproperty
func LSSharedFileListItemCopyProperty(inItem LSSharedFileListItemRef, inPropertyName corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _lSSharedFileListItemCopyProperty == nil {
		panic("coreservices: symbol LSSharedFileListItemCopyProperty not loaded")
	}
	return _lSSharedFileListItemCopyProperty(inItem, inPropertyName)
}

var _lSSharedFileListItemCopyResolvedURL func(inItem LSSharedFileListItemRef, inFlags LSSharedFileListResolutionFlags, outError *corefoundation.CFErrorRef) corefoundation.CFURLRef

// LSSharedFileListItemCopyResolvedURL.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1449882-lssharedfilelistitemcopyresolved
func LSSharedFileListItemCopyResolvedURL(inItem LSSharedFileListItemRef, inFlags LSSharedFileListResolutionFlags, outError *corefoundation.CFErrorRef) corefoundation.CFURLRef {
	if _lSSharedFileListItemCopyResolvedURL == nil {
		panic("coreservices: symbol LSSharedFileListItemCopyResolvedURL not loaded")
	}
	return _lSSharedFileListItemCopyResolvedURL(inItem, inFlags, outError)
}

var _lSSharedFileListItemGetID func(inItem LSSharedFileListItemRef) uint32

// LSSharedFileListItemGetID.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1443305-lssharedfilelistitemgetid
func LSSharedFileListItemGetID(inItem LSSharedFileListItemRef) uint32 {
	if _lSSharedFileListItemGetID == nil {
		panic("coreservices: symbol LSSharedFileListItemGetID not loaded")
	}
	return _lSSharedFileListItemGetID(inItem)
}

var _lSSharedFileListItemGetTypeID func() uint

// LSSharedFileListItemGetTypeID.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1447138-lssharedfilelistitemgettypeid
func LSSharedFileListItemGetTypeID() uint {
	if _lSSharedFileListItemGetTypeID == nil {
		panic("coreservices: symbol LSSharedFileListItemGetTypeID not loaded")
	}
	return _lSSharedFileListItemGetTypeID()
}

var _lSSharedFileListItemMove func(inList LSSharedFileListRef, inItem LSSharedFileListItemRef, inMoveAfterItem LSSharedFileListItemRef) int32

// LSSharedFileListItemMove.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1444348-lssharedfilelistitemmove
func LSSharedFileListItemMove(inList LSSharedFileListRef, inItem LSSharedFileListItemRef, inMoveAfterItem LSSharedFileListItemRef) int32 {
	if _lSSharedFileListItemMove == nil {
		panic("coreservices: symbol LSSharedFileListItemMove not loaded")
	}
	return _lSSharedFileListItemMove(inList, inItem, inMoveAfterItem)
}

var _lSSharedFileListItemRemove func(inList LSSharedFileListRef, inItem LSSharedFileListItemRef) int32

// LSSharedFileListItemRemove.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1442025-lssharedfilelistitemremove
func LSSharedFileListItemRemove(inList LSSharedFileListRef, inItem LSSharedFileListItemRef) int32 {
	if _lSSharedFileListItemRemove == nil {
		panic("coreservices: symbol LSSharedFileListItemRemove not loaded")
	}
	return _lSSharedFileListItemRemove(inList, inItem)
}

var _lSSharedFileListItemResolve func(inItem LSSharedFileListItemRef, inFlags LSSharedFileListResolutionFlags, outURL *corefoundation.CFURLRef, outRef unsafe.Pointer) int32

// LSSharedFileListItemResolve.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/coreservices/1447347-lssharedfilelistitemresolve
func LSSharedFileListItemResolve(inItem LSSharedFileListItemRef, inFlags LSSharedFileListResolutionFlags, outURL *corefoundation.CFURLRef, outRef unsafe.Pointer) int32 {
	if _lSSharedFileListItemResolve == nil {
		panic("coreservices: symbol LSSharedFileListItemResolve not loaded")
	}
	return _lSSharedFileListItemResolve(inItem, inFlags, outURL, outRef)
}

var _lSSharedFileListItemSetProperty func(inItem LSSharedFileListItemRef, inPropertyName corefoundation.CFStringRef, inPropertyData corefoundation.CFTypeRef) int32

// LSSharedFileListItemSetProperty.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1445766-lssharedfilelistitemsetproperty
func LSSharedFileListItemSetProperty(inItem LSSharedFileListItemRef, inPropertyName corefoundation.CFStringRef, inPropertyData corefoundation.CFTypeRef) int32 {
	if _lSSharedFileListItemSetProperty == nil {
		panic("coreservices: symbol LSSharedFileListItemSetProperty not loaded")
	}
	return _lSSharedFileListItemSetProperty(inItem, inPropertyName, inPropertyData)
}

var _lSSharedFileListRemoveAllItems func(inList LSSharedFileListRef) int32

// LSSharedFileListRemoveAllItems.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1446389-lssharedfilelistremoveallitems
func LSSharedFileListRemoveAllItems(inList LSSharedFileListRef) int32 {
	if _lSSharedFileListRemoveAllItems == nil {
		panic("coreservices: symbol LSSharedFileListRemoveAllItems not loaded")
	}
	return _lSSharedFileListRemoveAllItems(inList)
}

var _lSSharedFileListRemoveObserver func(inList LSSharedFileListRef, inRunloop corefoundation.CFRunLoopRef, inRunloopMode corefoundation.CFStringRef, callback LSSharedFileListChangedProcPtr, context unsafe.Pointer)

// LSSharedFileListRemoveObserver.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1443404-lssharedfilelistremoveobserver
func LSSharedFileListRemoveObserver(inList LSSharedFileListRef, inRunloop corefoundation.CFRunLoopRef, inRunloopMode corefoundation.CFStringRef, callback LSSharedFileListChangedProcPtr, context unsafe.Pointer) {
	if _lSSharedFileListRemoveObserver == nil {
		panic("coreservices: symbol LSSharedFileListRemoveObserver not loaded")
	}
	_lSSharedFileListRemoveObserver(inList, inRunloop, inRunloopMode, callback, context)
}

var _lSSharedFileListSetAuthorization func(inList LSSharedFileListRef, inAuthorization security.AuthorizationRef) int32

// LSSharedFileListSetAuthorization.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1446834-lssharedfilelistsetauthorization
func LSSharedFileListSetAuthorization(inList LSSharedFileListRef, inAuthorization security.AuthorizationRef) int32 {
	if _lSSharedFileListSetAuthorization == nil {
		panic("coreservices: symbol LSSharedFileListSetAuthorization not loaded")
	}
	return _lSSharedFileListSetAuthorization(inList, inAuthorization)
}

var _lSSharedFileListSetProperty func(inList LSSharedFileListRef, inPropertyName corefoundation.CFStringRef, inPropertyData corefoundation.CFTypeRef) int32

// LSSharedFileListSetProperty.
//
// Deprecated: Deprecated since macOS 10.11.
//
// See: https://developer.apple.com/documentation/coreservices/1448857-lssharedfilelistsetproperty
func LSSharedFileListSetProperty(inList LSSharedFileListRef, inPropertyName corefoundation.CFStringRef, inPropertyData corefoundation.CFTypeRef) int32 {
	if _lSSharedFileListSetProperty == nil {
		panic("coreservices: symbol LSSharedFileListSetProperty not loaded")
	}
	return _lSSharedFileListSetProperty(inList, inPropertyName, inPropertyData)
}

var _loadResource func(arg0 uintptr)

// LoadResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529245-loadresource
func LoadResource(arg0 uintptr) {
	if _loadResource == nil {
		panic("coreservices: symbol LoadResource not loaded")
	}
	_loadResource(arg0)
}

var _localeCountNames func(arg0 LocaleRef, arg1 LocaleOperationVariant, arg2 LocaleNameMask, arg3 uintptr) int32

// LocaleCountNames.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1580255-localecountnames
func LocaleCountNames(arg0 LocaleRef, arg1 LocaleOperationVariant, arg2 LocaleNameMask, arg3 uintptr) int32 {
	if _localeCountNames == nil {
		panic("coreservices: symbol LocaleCountNames not loaded")
	}
	return _localeCountNames(arg0, arg1, arg2, arg3)
}

var _localeGetIndName func(arg0 LocaleRef, arg1 LocaleOperationVariant, arg2 LocaleNameMask, arg3 uintptr, arg4 uint, arg5 uint, arg6 uint16, arg7 LocaleRef) int32

// LocaleGetIndName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1580265-localegetindname
func LocaleGetIndName(arg0 LocaleRef, arg1 LocaleOperationVariant, arg2 LocaleNameMask, arg3 uintptr, arg4 uint, arg5 uint, arg6 uint16, arg7 LocaleRef) int32 {
	if _localeGetIndName == nil {
		panic("coreservices: symbol LocaleGetIndName not loaded")
	}
	return _localeGetIndName(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _localeGetName func(arg0 LocaleRef, arg1 LocaleOperationVariant, arg2 LocaleNameMask, arg3 LocaleRef, arg4 uint, arg5 uint, arg6 uint16) int32

// LocaleGetName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1580248-localegetname
func LocaleGetName(arg0 LocaleRef, arg1 LocaleOperationVariant, arg2 LocaleNameMask, arg3 LocaleRef, arg4 uint, arg5 uint, arg6 uint16) int32 {
	if _localeGetName == nil {
		panic("coreservices: symbol LocaleGetName not loaded")
	}
	return _localeGetName(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _localeOperationCountLocales func(arg0 LocaleOperationClass, arg1 uintptr) int32

// LocaleOperationCountLocales.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1580264-localeoperationcountlocales
func LocaleOperationCountLocales(arg0 LocaleOperationClass, arg1 uintptr) int32 {
	if _localeOperationCountLocales == nil {
		panic("coreservices: symbol LocaleOperationCountLocales not loaded")
	}
	return _localeOperationCountLocales(arg0, arg1)
}

var _localeOperationCountNames func(arg0 LocaleOperationClass, arg1 uintptr) int32

// LocaleOperationCountNames.
//
// See: https://developer.apple.com/documentation/coreservices/1580257-localeoperationcountnames
func LocaleOperationCountNames(arg0 LocaleOperationClass, arg1 uintptr) int32 {
	if _localeOperationCountNames == nil {
		panic("coreservices: symbol LocaleOperationCountNames not loaded")
	}
	return _localeOperationCountNames(arg0, arg1)
}

var _localeOperationGetIndName func(arg0 LocaleOperationClass, arg1 uintptr, arg2 uint, arg3 uint, arg4 uint16, arg5 LocaleRef) int32

// LocaleOperationGetIndName.
//
// See: https://developer.apple.com/documentation/coreservices/1580260-localeoperationgetindname
func LocaleOperationGetIndName(arg0 LocaleOperationClass, arg1 uintptr, arg2 uint, arg3 uint, arg4 uint16, arg5 LocaleRef) int32 {
	if _localeOperationGetIndName == nil {
		panic("coreservices: symbol LocaleOperationGetIndName not loaded")
	}
	return _localeOperationGetIndName(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _localeOperationGetLocales func(arg0 LocaleOperationClass, arg1 uintptr, arg2 uintptr, arg3 LocaleAndVariant) int32

// LocaleOperationGetLocales.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1580279-localeoperationgetlocales
func LocaleOperationGetLocales(arg0 LocaleOperationClass, arg1 uintptr, arg2 uintptr, arg3 LocaleAndVariant) int32 {
	if _localeOperationGetLocales == nil {
		panic("coreservices: symbol LocaleOperationGetLocales not loaded")
	}
	return _localeOperationGetLocales(arg0, arg1, arg2, arg3)
}

var _localeOperationGetName func(arg0 LocaleOperationClass, arg1 LocaleRef, arg2 uint, arg3 uint, arg4 uint16) int32

// LocaleOperationGetName.
//
// See: https://developer.apple.com/documentation/coreservices/1580269-localeoperationgetname
func LocaleOperationGetName(arg0 LocaleOperationClass, arg1 LocaleRef, arg2 uint, arg3 uint, arg4 uint16) int32 {
	if _localeOperationGetName == nil {
		panic("coreservices: symbol LocaleOperationGetName not loaded")
	}
	return _localeOperationGetName(arg0, arg1, arg2, arg3, arg4)
}

var _localeRefFromLangOrRegionCode func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 LocaleRef) int32

// LocaleRefFromLangOrRegionCode.
//
// See: https://developer.apple.com/documentation/coreservices/1580276-localereffromlangorregioncode
func LocaleRefFromLangOrRegionCode(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 LocaleRef) int32 {
	if _localeRefFromLangOrRegionCode == nil {
		panic("coreservices: symbol LocaleRefFromLangOrRegionCode not loaded")
	}
	return _localeRefFromLangOrRegionCode(arg0, arg1, arg2)
}

var _localeRefFromLocaleString func(arg0 int8, arg1 LocaleRef) int32

// LocaleRefFromLocaleString.
//
// See: https://developer.apple.com/documentation/coreservices/1580261-localereffromlocalestring
func LocaleRefFromLocaleString(arg0 int8, arg1 LocaleRef) int32 {
	if _localeRefFromLocaleString == nil {
		panic("coreservices: symbol LocaleRefFromLocaleString not loaded")
	}
	return _localeRefFromLocaleString(arg0, arg1)
}

var _localeRefGetPartString func(arg0 LocaleRef, arg1 LocalePartMask, arg2 uintptr, arg3 int8) int32

// LocaleRefGetPartString.
//
// See: https://developer.apple.com/documentation/coreservices/1580273-localerefgetpartstring
func LocaleRefGetPartString(arg0 LocaleRef, arg1 LocalePartMask, arg2 uintptr, arg3 int8) int32 {
	if _localeRefGetPartString == nil {
		panic("coreservices: symbol LocaleRefGetPartString not loaded")
	}
	return _localeRefGetPartString(arg0, arg1, arg2, arg3)
}

var _localeStringToLangAndRegionCodes func(arg0 int8, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32

// LocaleStringToLangAndRegionCodes.
//
// See: https://developer.apple.com/documentation/coreservices/1580277-localestringtolangandregioncodes
func LocaleStringToLangAndRegionCodes(arg0 int8, arg1 unsafe.Pointer, arg2 unsafe.Pointer) int32 {
	if _localeStringToLangAndRegionCodes == nil {
		panic("coreservices: symbol LocaleStringToLangAndRegionCodes not loaded")
	}
	return _localeStringToLangAndRegionCodes(arg0, arg1, arg2)
}

var _long2Fix func(arg0 int32) objectivec.IObject

// Long2Fix.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409235-long2fix
func Long2Fix(arg0 int32) objectivec.IObject {
	if _long2Fix == nil {
		panic("coreservices: symbol Long2Fix not loaded")
	}
	return _long2Fix(arg0)
}

var _longDoubleToSInt64 func(arg0 unsafe.Pointer) int64

// LongDoubleToSInt64.
//
// See: https://developer.apple.com/documentation/coreservices/1536234-longdoubletosint64
func LongDoubleToSInt64(arg0 unsafe.Pointer) int64 {
	if _longDoubleToSInt64 == nil {
		panic("coreservices: symbol LongDoubleToSInt64 not loaded")
	}
	return _longDoubleToSInt64(arg0)
}

var _longDoubleToUInt64 func(arg0 unsafe.Pointer) uint64

// LongDoubleToUInt64.
//
// See: https://developer.apple.com/documentation/coreservices/1536290-longdoubletouint64
func LongDoubleToUInt64(arg0 unsafe.Pointer) uint64 {
	if _longDoubleToUInt64 == nil {
		panic("coreservices: symbol LongDoubleToUInt64 not loaded")
	}
	return _longDoubleToUInt64(arg0)
}

var _mDCopyLabelKinds func() corefoundation.CFArrayRef

// MDCopyLabelKinds.
//
// See: https://developer.apple.com/documentation/coreservices/1442887-mdcopylabelkinds
func MDCopyLabelKinds() corefoundation.CFArrayRef {
	if _mDCopyLabelKinds == nil {
		panic("coreservices: symbol MDCopyLabelKinds not loaded")
	}
	return _mDCopyLabelKinds()
}

var _mDCopyLabelWithUUID func(labelUUID corefoundation.CFUUID) MDLabelRef

// MDCopyLabelWithUUID.
//
// See: https://developer.apple.com/documentation/coreservices/1447030-mdcopylabelwithuuid
func MDCopyLabelWithUUID(labelUUID corefoundation.CFUUID) MDLabelRef {
	if _mDCopyLabelWithUUID == nil {
		panic("coreservices: symbol MDCopyLabelWithUUID not loaded")
	}
	return _mDCopyLabelWithUUID(labelUUID)
}

var _mDCopyLabelsMatchingExpression func(simpleQueryString corefoundation.CFStringRef) corefoundation.CFArrayRef

// MDCopyLabelsMatchingExpression.
//
// See: https://developer.apple.com/documentation/coreservices/1448237-mdcopylabelsmatchingexpression
func MDCopyLabelsMatchingExpression(simpleQueryString corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _mDCopyLabelsMatchingExpression == nil {
		panic("coreservices: symbol MDCopyLabelsMatchingExpression not loaded")
	}
	return _mDCopyLabelsMatchingExpression(simpleQueryString)
}

var _mDCopyLabelsWithKind func(kind corefoundation.CFStringRef) corefoundation.CFArrayRef

// MDCopyLabelsWithKind.
//
// See: https://developer.apple.com/documentation/coreservices/1444230-mdcopylabelswithkind
func MDCopyLabelsWithKind(kind corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _mDCopyLabelsWithKind == nil {
		panic("coreservices: symbol MDCopyLabelsWithKind not loaded")
	}
	return _mDCopyLabelsWithKind(kind)
}

var _mDItemCopyAttribute func(item MDItemRef, name corefoundation.CFStringRef) corefoundation.CFTypeRef

// MDItemCopyAttribute returns the value of the specified attribute in the metadata item.
//
// See: https://developer.apple.com/documentation/coreservices/1427080-mditemcopyattribute
func MDItemCopyAttribute(item MDItemRef, name corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _mDItemCopyAttribute == nil {
		panic("coreservices: symbol MDItemCopyAttribute not loaded")
	}
	return _mDItemCopyAttribute(item, name)
}

var _mDItemCopyAttributeList func(arg0 MDItemRef) corefoundation.CFDictionaryRef

// MDItemCopyAttributeList returns the values of the specified attributes in the metadata item.
//
// See: https://developer.apple.com/documentation/coreservices/1427028-mditemcopyattributelist
func MDItemCopyAttributeList(arg0 MDItemRef) corefoundation.CFDictionaryRef {
	if _mDItemCopyAttributeList == nil {
		panic("coreservices: symbol MDItemCopyAttributeList not loaded")
	}
	return _mDItemCopyAttributeList(arg0)
}

var _mDItemCopyAttributeNames func(item MDItemRef) corefoundation.CFArrayRef

// MDItemCopyAttributeNames returns an array containing the attribute names existing in the metadata item.
//
// See: https://developer.apple.com/documentation/coreservices/1427066-mditemcopyattributenames
func MDItemCopyAttributeNames(item MDItemRef) corefoundation.CFArrayRef {
	if _mDItemCopyAttributeNames == nil {
		panic("coreservices: symbol MDItemCopyAttributeNames not loaded")
	}
	return _mDItemCopyAttributeNames(item)
}

var _mDItemCopyAttributes func(item MDItemRef, names corefoundation.CFArrayRef) corefoundation.CFDictionaryRef

// MDItemCopyAttributes returns the values of the specified attributes in the metadata item.
//
// See: https://developer.apple.com/documentation/coreservices/1426980-mditemcopyattributes
func MDItemCopyAttributes(item MDItemRef, names corefoundation.CFArrayRef) corefoundation.CFDictionaryRef {
	if _mDItemCopyAttributes == nil {
		panic("coreservices: symbol MDItemCopyAttributes not loaded")
	}
	return _mDItemCopyAttributes(item, names)
}

var _mDItemCopyLabels func(item MDItemRef) corefoundation.CFArrayRef

// MDItemCopyLabels.
//
// See: https://developer.apple.com/documentation/coreservices/1442606-mditemcopylabels
func MDItemCopyLabels(item MDItemRef) corefoundation.CFArrayRef {
	if _mDItemCopyLabels == nil {
		panic("coreservices: symbol MDItemCopyLabels not loaded")
	}
	return _mDItemCopyLabels(item)
}

var _mDItemCreate func(allocator corefoundation.CFAllocatorRef, path corefoundation.CFStringRef) MDItemRef

// MDItemCreate creates an MDItem object for a file at the specified path.
//
// See: https://developer.apple.com/documentation/coreservices/1426917-mditemcreate
func MDItemCreate(allocator corefoundation.CFAllocatorRef, path corefoundation.CFStringRef) MDItemRef {
	if _mDItemCreate == nil {
		panic("coreservices: symbol MDItemCreate not loaded")
	}
	return _mDItemCreate(allocator, path)
}

var _mDItemCreateWithURL func(allocator corefoundation.CFAllocatorRef, url corefoundation.CFURLRef) MDItemRef

// MDItemCreateWithURL creates an MDItem object for a file at the specified file URL.
//
// See: https://developer.apple.com/documentation/coreservices/1427034-mditemcreatewithurl
func MDItemCreateWithURL(allocator corefoundation.CFAllocatorRef, url corefoundation.CFURLRef) MDItemRef {
	if _mDItemCreateWithURL == nil {
		panic("coreservices: symbol MDItemCreateWithURL not loaded")
	}
	return _mDItemCreateWithURL(allocator, url)
}

var _mDItemGetCacheFileDescriptors func(items corefoundation.CFArrayRef, completionHandler corefoundation.CFArrayRef) unsafe.Pointer

// MDItemGetCacheFileDescriptors.
//
// See: https://developer.apple.com/documentation/coreservices/4485578-mditemgetcachefiledescriptors
func MDItemGetCacheFileDescriptors(items corefoundation.CFArrayRef, completionHandler corefoundation.CFArrayRef) unsafe.Pointer {
	if _mDItemGetCacheFileDescriptors == nil {
		panic("coreservices: symbol MDItemGetCacheFileDescriptors not loaded")
	}
	return _mDItemGetCacheFileDescriptors(items, completionHandler)
}

var _mDItemGetTypeID func() uint

// MDItemGetTypeID returns the type identifier of all MDItem instances.
//
// See: https://developer.apple.com/documentation/coreservices/1427168-mditemgettypeid
func MDItemGetTypeID() uint {
	if _mDItemGetTypeID == nil {
		panic("coreservices: symbol MDItemGetTypeID not loaded")
	}
	return _mDItemGetTypeID()
}

var _mDItemRemoveLabel func(item MDItemRef, label MDLabelRef) bool

// MDItemRemoveLabel.
//
// See: https://developer.apple.com/documentation/coreservices/1446067-mditemremovelabel
func MDItemRemoveLabel(item MDItemRef, label MDLabelRef) bool {
	if _mDItemRemoveLabel == nil {
		panic("coreservices: symbol MDItemRemoveLabel not loaded")
	}
	return _mDItemRemoveLabel(item, label)
}

var _mDItemSetLabel func(item MDItemRef, label MDLabelRef) bool

// MDItemSetLabel.
//
// See: https://developer.apple.com/documentation/coreservices/1442559-mditemsetlabel
func MDItemSetLabel(item MDItemRef, label MDLabelRef) bool {
	if _mDItemSetLabel == nil {
		panic("coreservices: symbol MDItemSetLabel not loaded")
	}
	return _mDItemSetLabel(item, label)
}

var _mDItemsCopyAttributes func(items corefoundation.CFArrayRef, names corefoundation.CFArrayRef) corefoundation.CFArrayRef

// MDItemsCopyAttributes.
//
// See: https://developer.apple.com/documentation/coreservices/1426975-mditemscopyattributes
func MDItemsCopyAttributes(items corefoundation.CFArrayRef, names corefoundation.CFArrayRef) corefoundation.CFArrayRef {
	if _mDItemsCopyAttributes == nil {
		panic("coreservices: symbol MDItemsCopyAttributes not loaded")
	}
	return _mDItemsCopyAttributes(items, names)
}

var _mDItemsCreateWithURLs func(allocator corefoundation.CFAllocatorRef, urls corefoundation.CFArrayRef) corefoundation.CFArrayRef

// MDItemsCreateWithURLs.
//
// See: https://developer.apple.com/documentation/coreservices/1427086-mditemscreatewithurls
func MDItemsCreateWithURLs(allocator corefoundation.CFAllocatorRef, urls corefoundation.CFArrayRef) corefoundation.CFArrayRef {
	if _mDItemsCreateWithURLs == nil {
		panic("coreservices: symbol MDItemsCreateWithURLs not loaded")
	}
	return _mDItemsCreateWithURLs(allocator, urls)
}

var _mDLabelCopyAttribute func(label MDLabelRef, name corefoundation.CFStringRef) corefoundation.CFTypeRef

// MDLabelCopyAttribute.
//
// See: https://developer.apple.com/documentation/coreservices/1445456-mdlabelcopyattribute
func MDLabelCopyAttribute(label MDLabelRef, name corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _mDLabelCopyAttribute == nil {
		panic("coreservices: symbol MDLabelCopyAttribute not loaded")
	}
	return _mDLabelCopyAttribute(label, name)
}

var _mDLabelCopyAttributeName func(label MDLabelRef) corefoundation.CFStringRef

// MDLabelCopyAttributeName.
//
// See: https://developer.apple.com/documentation/coreservices/1445522-mdlabelcopyattributename
func MDLabelCopyAttributeName(label MDLabelRef) corefoundation.CFStringRef {
	if _mDLabelCopyAttributeName == nil {
		panic("coreservices: symbol MDLabelCopyAttributeName not loaded")
	}
	return _mDLabelCopyAttributeName(label)
}

var _mDLabelCreate func(allocator corefoundation.CFAllocatorRef, displayName corefoundation.CFStringRef, kind corefoundation.CFStringRef, domain unsafe.Pointer) MDLabelRef

// MDLabelCreate.
//
// See: https://developer.apple.com/documentation/coreservices/1442614-mdlabelcreate
func MDLabelCreate(allocator corefoundation.CFAllocatorRef, displayName corefoundation.CFStringRef, kind corefoundation.CFStringRef, domain unsafe.Pointer) MDLabelRef {
	if _mDLabelCreate == nil {
		panic("coreservices: symbol MDLabelCreate not loaded")
	}
	return _mDLabelCreate(allocator, displayName, kind, domain)
}

var _mDLabelDelete func(label MDLabelRef) bool

// MDLabelDelete.
//
// See: https://developer.apple.com/documentation/coreservices/1449203-mdlabeldelete
func MDLabelDelete(label MDLabelRef) bool {
	if _mDLabelDelete == nil {
		panic("coreservices: symbol MDLabelDelete not loaded")
	}
	return _mDLabelDelete(label)
}

var _mDLabelGetTypeID func() uint

// MDLabelGetTypeID.
//
// See: https://developer.apple.com/documentation/coreservices/1446579-mdlabelgettypeid
func MDLabelGetTypeID() uint {
	if _mDLabelGetTypeID == nil {
		panic("coreservices: symbol MDLabelGetTypeID not loaded")
	}
	return _mDLabelGetTypeID()
}

var _mDLabelSetAttributes func(label MDLabelRef, attrs corefoundation.CFDictionaryRef) bool

// MDLabelSetAttributes.
//
// See: https://developer.apple.com/documentation/coreservices/1449005-mdlabelsetattributes
func MDLabelSetAttributes(label MDLabelRef, attrs corefoundation.CFDictionaryRef) bool {
	if _mDLabelSetAttributes == nil {
		panic("coreservices: symbol MDLabelSetAttributes not loaded")
	}
	return _mDLabelSetAttributes(label, attrs)
}

var _mDQueryCopyQueryString func(query MDQueryRef) corefoundation.CFStringRef

// MDQueryCopyQueryString returns the query string of the query.
//
// See: https://developer.apple.com/documentation/coreservices/1413004-mdquerycopyquerystring
func MDQueryCopyQueryString(query MDQueryRef) corefoundation.CFStringRef {
	if _mDQueryCopyQueryString == nil {
		panic("coreservices: symbol MDQueryCopyQueryString not loaded")
	}
	return _mDQueryCopyQueryString(query)
}

var _mDQueryCopySortingAttributes func(query MDQueryRef) corefoundation.CFArrayRef

// MDQueryCopySortingAttributes returns the list of attribute names used to sort the results.
//
// See: https://developer.apple.com/documentation/coreservices/1413059-mdquerycopysortingattributes
func MDQueryCopySortingAttributes(query MDQueryRef) corefoundation.CFArrayRef {
	if _mDQueryCopySortingAttributes == nil {
		panic("coreservices: symbol MDQueryCopySortingAttributes not loaded")
	}
	return _mDQueryCopySortingAttributes(query)
}

var _mDQueryCopyValueListAttributes func(query MDQueryRef) corefoundation.CFArrayRef

// MDQueryCopyValueListAttributes returns the list of attribute names for which values are being collected by the query.
//
// See: https://developer.apple.com/documentation/coreservices/1413071-mdquerycopyvaluelistattributes
func MDQueryCopyValueListAttributes(query MDQueryRef) corefoundation.CFArrayRef {
	if _mDQueryCopyValueListAttributes == nil {
		panic("coreservices: symbol MDQueryCopyValueListAttributes not loaded")
	}
	return _mDQueryCopyValueListAttributes(query)
}

var _mDQueryCopyValuesOfAttribute func(query MDQueryRef, name corefoundation.CFStringRef) corefoundation.CFArrayRef

// MDQueryCopyValuesOfAttribute returns the list of values from the results of the query for the specified attribute.
//
// See: https://developer.apple.com/documentation/coreservices/1413105-mdquerycopyvaluesofattribute
func MDQueryCopyValuesOfAttribute(query MDQueryRef, name corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _mDQueryCopyValuesOfAttribute == nil {
		panic("coreservices: symbol MDQueryCopyValuesOfAttribute not loaded")
	}
	return _mDQueryCopyValuesOfAttribute(query, name)
}

var _mDQueryCreate func(allocator corefoundation.CFAllocatorRef, queryString corefoundation.CFStringRef, valueListAttrs corefoundation.CFArrayRef, sortingAttrs corefoundation.CFArrayRef) MDQueryRef

// MDQueryCreate creates a new query instance.
//
// See: https://developer.apple.com/documentation/coreservices/1413029-mdquerycreate
func MDQueryCreate(allocator corefoundation.CFAllocatorRef, queryString corefoundation.CFStringRef, valueListAttrs corefoundation.CFArrayRef, sortingAttrs corefoundation.CFArrayRef) MDQueryRef {
	if _mDQueryCreate == nil {
		panic("coreservices: symbol MDQueryCreate not loaded")
	}
	return _mDQueryCreate(allocator, queryString, valueListAttrs, sortingAttrs)
}

var _mDQueryCreateForItems func(allocator corefoundation.CFAllocatorRef, queryString corefoundation.CFStringRef, valueListAttrs corefoundation.CFArrayRef, sortingAttrs corefoundation.CFArrayRef, items corefoundation.CFArrayRef) MDQueryRef

// MDQueryCreateForItems.
//
// See: https://developer.apple.com/documentation/coreservices/1413031-mdquerycreateforitems
func MDQueryCreateForItems(allocator corefoundation.CFAllocatorRef, queryString corefoundation.CFStringRef, valueListAttrs corefoundation.CFArrayRef, sortingAttrs corefoundation.CFArrayRef, items corefoundation.CFArrayRef) MDQueryRef {
	if _mDQueryCreateForItems == nil {
		panic("coreservices: symbol MDQueryCreateForItems not loaded")
	}
	return _mDQueryCreateForItems(allocator, queryString, valueListAttrs, sortingAttrs, items)
}

var _mDQueryCreateSubset func(allocator corefoundation.CFAllocatorRef, query MDQueryRef, queryString corefoundation.CFStringRef, valueListAttrs corefoundation.CFArrayRef, sortingAttrs corefoundation.CFArrayRef) MDQueryRef

// MDQueryCreateSubset creates a new query that is a subset of the specified parentquery.
//
// See: https://developer.apple.com/documentation/coreservices/1413027-mdquerycreatesubset
func MDQueryCreateSubset(allocator corefoundation.CFAllocatorRef, query MDQueryRef, queryString corefoundation.CFStringRef, valueListAttrs corefoundation.CFArrayRef, sortingAttrs corefoundation.CFArrayRef) MDQueryRef {
	if _mDQueryCreateSubset == nil {
		panic("coreservices: symbol MDQueryCreateSubset not loaded")
	}
	return _mDQueryCreateSubset(allocator, query, queryString, valueListAttrs, sortingAttrs)
}

var _mDQueryDisableUpdates func(query MDQueryRef)

// MDQueryDisableUpdates disables updates to the query result list.
//
// See: https://developer.apple.com/documentation/coreservices/1413041-mdquerydisableupdates
func MDQueryDisableUpdates(query MDQueryRef) {
	if _mDQueryDisableUpdates == nil {
		panic("coreservices: symbol MDQueryDisableUpdates not loaded")
	}
	_mDQueryDisableUpdates(query)
}

var _mDQueryEnableUpdates func(query MDQueryRef)

// MDQueryEnableUpdates enables updates to the query result list.
//
// See: https://developer.apple.com/documentation/coreservices/1413066-mdqueryenableupdates
func MDQueryEnableUpdates(query MDQueryRef) {
	if _mDQueryEnableUpdates == nil {
		panic("coreservices: symbol MDQueryEnableUpdates not loaded")
	}
	_mDQueryEnableUpdates(query)
}

var _mDQueryExecute func(query MDQueryRef, optionFlags uint64) bool

// MDQueryExecute run the query, and populate the query with the results.
//
// See: https://developer.apple.com/documentation/coreservices/1413099-mdqueryexecute
func MDQueryExecute(query MDQueryRef, optionFlags uint64) bool {
	if _mDQueryExecute == nil {
		panic("coreservices: symbol MDQueryExecute not loaded")
	}
	return _mDQueryExecute(query, optionFlags)
}

var _mDQueryGetAttributeValueOfResultAtIndex func(query MDQueryRef, name corefoundation.CFStringRef, idx int) unsafe.Pointer

// MDQueryGetAttributeValueOfResultAtIndex returns the value of the named attribute for the result at the given index.
//
// See: https://developer.apple.com/documentation/coreservices/1413046-mdquerygetattributevalueofresult
func MDQueryGetAttributeValueOfResultAtIndex(query MDQueryRef, name corefoundation.CFStringRef, idx int) unsafe.Pointer {
	if _mDQueryGetAttributeValueOfResultAtIndex == nil {
		panic("coreservices: symbol MDQueryGetAttributeValueOfResultAtIndex not loaded")
	}
	return _mDQueryGetAttributeValueOfResultAtIndex(query, name, idx)
}

var _mDQueryGetBatchingParameters func(query MDQueryRef) MDQueryBatchingParams

// MDQueryGetBatchingParameters returns the current parameters that control the batching of progress notifications.
//
// See: https://developer.apple.com/documentation/coreservices/1413006-mdquerygetbatchingparameters
func MDQueryGetBatchingParameters(query MDQueryRef) MDQueryBatchingParams {
	if _mDQueryGetBatchingParameters == nil {
		panic("coreservices: symbol MDQueryGetBatchingParameters not loaded")
	}
	return _mDQueryGetBatchingParameters(query)
}

var _mDQueryGetCountOfResultsWithAttributeValue func(query MDQueryRef, name corefoundation.CFStringRef, value corefoundation.CFTypeRef) int

// MDQueryGetCountOfResultsWithAttributeValue returns the number of results which have the given attribute and attribute value.
//
// See: https://developer.apple.com/documentation/coreservices/1413009-mdquerygetcountofresultswithattr
func MDQueryGetCountOfResultsWithAttributeValue(query MDQueryRef, name corefoundation.CFStringRef, value corefoundation.CFTypeRef) int {
	if _mDQueryGetCountOfResultsWithAttributeValue == nil {
		panic("coreservices: symbol MDQueryGetCountOfResultsWithAttributeValue not loaded")
	}
	return _mDQueryGetCountOfResultsWithAttributeValue(query, name, value)
}

var _mDQueryGetIndexOfResult func(query MDQueryRef, result unsafe.Pointer) int

// MDQueryGetIndexOfResult returns the current index of the given result.
//
// See: https://developer.apple.com/documentation/coreservices/1413093-mdquerygetindexofresult
func MDQueryGetIndexOfResult(query MDQueryRef, result unsafe.Pointer) int {
	if _mDQueryGetIndexOfResult == nil {
		panic("coreservices: symbol MDQueryGetIndexOfResult not loaded")
	}
	return _mDQueryGetIndexOfResult(query, result)
}

var _mDQueryGetResultAtIndex func(query MDQueryRef, idx int) unsafe.Pointer

// MDQueryGetResultAtIndex returns the current result at the given index.
//
// See: https://developer.apple.com/documentation/coreservices/1413055-mdquerygetresultatindex
func MDQueryGetResultAtIndex(query MDQueryRef, idx int) unsafe.Pointer {
	if _mDQueryGetResultAtIndex == nil {
		panic("coreservices: symbol MDQueryGetResultAtIndex not loaded")
	}
	return _mDQueryGetResultAtIndex(query, idx)
}

var _mDQueryGetResultCount func(query MDQueryRef) int

// MDQueryGetResultCount returns the number of results currently collected by the query.
//
// See: https://developer.apple.com/documentation/coreservices/1413008-mdquerygetresultcount
func MDQueryGetResultCount(query MDQueryRef) int {
	if _mDQueryGetResultCount == nil {
		panic("coreservices: symbol MDQueryGetResultCount not loaded")
	}
	return _mDQueryGetResultCount(query)
}

var _mDQueryGetSortOptionFlagsForAttribute func(query MDQueryRef, fieldName corefoundation.CFStringRef) uint32

// MDQueryGetSortOptionFlagsForAttribute.
//
// See: https://developer.apple.com/documentation/coreservices/1413013-mdquerygetsortoptionflagsforattr
func MDQueryGetSortOptionFlagsForAttribute(query MDQueryRef, fieldName corefoundation.CFStringRef) uint32 {
	if _mDQueryGetSortOptionFlagsForAttribute == nil {
		panic("coreservices: symbol MDQueryGetSortOptionFlagsForAttribute not loaded")
	}
	return _mDQueryGetSortOptionFlagsForAttribute(query, fieldName)
}

var _mDQueryGetTypeID func() uint

// MDQueryGetTypeID returns the type identifier of all MDQuery instances
//
// See: https://developer.apple.com/documentation/coreservices/1413037-mdquerygettypeid
func MDQueryGetTypeID() uint {
	if _mDQueryGetTypeID == nil {
		panic("coreservices: symbol MDQueryGetTypeID not loaded")
	}
	return _mDQueryGetTypeID()
}

var _mDQueryIsGatheringComplete func(query MDQueryRef) bool

// MDQueryIsGatheringComplete returns true if the first phase of a query, the initial result gathering, has finished.
//
// See: https://developer.apple.com/documentation/coreservices/1413032-mdqueryisgatheringcomplete
func MDQueryIsGatheringComplete(query MDQueryRef) bool {
	if _mDQueryIsGatheringComplete == nil {
		panic("coreservices: symbol MDQueryIsGatheringComplete not loaded")
	}
	return _mDQueryIsGatheringComplete(query)
}

var _mDQuerySetBatchingParameters func(query MDQueryRef, params unsafe.Pointer)

// MDQuerySetBatchingParameters set the query batching parameters.
//
// See: https://developer.apple.com/documentation/coreservices/1413103-mdquerysetbatchingparameters
func MDQuerySetBatchingParameters(query MDQueryRef, params unsafe.Pointer) {
	if _mDQuerySetBatchingParameters == nil {
		panic("coreservices: symbol MDQuerySetBatchingParameters not loaded")
	}
	_mDQuerySetBatchingParameters(query, params)
}

var _mDQuerySetCreateResultFunction func(query MDQueryRef, func_ MDQueryCreateResultFunction, context unsafe.Pointer, cb unsafe.Pointer)

// MDQuerySetCreateResultFunction sets the function used to create the result objects of the MDQuery.
//
// See: https://developer.apple.com/documentation/coreservices/1413064-mdquerysetcreateresultfunction
func MDQuerySetCreateResultFunction(query MDQueryRef, func_ MDQueryCreateResultFunction, context unsafe.Pointer, cb unsafe.Pointer) {
	if _mDQuerySetCreateResultFunction == nil {
		panic("coreservices: symbol MDQuerySetCreateResultFunction not loaded")
	}
	_mDQuerySetCreateResultFunction(query, func_, context, cb)
}

var _mDQuerySetCreateValueFunction func(query MDQueryRef, func_ MDQueryCreateValueFunction, context unsafe.Pointer, cb unsafe.Pointer)

// MDQuerySetCreateValueFunction sets the function used to create the value objects of the MDQuery.
//
// See: https://developer.apple.com/documentation/coreservices/1413017-mdquerysetcreatevaluefunction
func MDQuerySetCreateValueFunction(query MDQueryRef, func_ MDQueryCreateValueFunction, context unsafe.Pointer, cb unsafe.Pointer) {
	if _mDQuerySetCreateValueFunction == nil {
		panic("coreservices: symbol MDQuerySetCreateValueFunction not loaded")
	}
	_mDQuerySetCreateValueFunction(query, func_, context, cb)
}

var _mDQuerySetDispatchQueue func(query MDQueryRef, queue uintptr)

// MDQuerySetDispatchQueue sets the dispatch queue on which query results will be delivered by MDQueryExecute.
//
// See: https://developer.apple.com/documentation/coreservices/1413019-mdquerysetdispatchqueue
func MDQuerySetDispatchQueue(query MDQueryRef, queue dispatch.Queue) {
	if _mDQuerySetDispatchQueue == nil {
		panic("coreservices: symbol MDQuerySetDispatchQueue not loaded")
	}
	_mDQuerySetDispatchQueue(query, uintptr(queue.Handle()))
}

var _mDQuerySetMaxCount func(query MDQueryRef, size int)

// MDQuerySetMaxCount sets the maximum number of results returned.
//
// See: https://developer.apple.com/documentation/coreservices/1413085-mdquerysetmaxcount
func MDQuerySetMaxCount(query MDQueryRef, size int) {
	if _mDQuerySetMaxCount == nil {
		panic("coreservices: symbol MDQuerySetMaxCount not loaded")
	}
	_mDQuerySetMaxCount(query, size)
}

var _mDQuerySetSearchScope func(query MDQueryRef, scopeDirectories corefoundation.CFArrayRef, scopeOptions uintptr)

// MDQuerySetSearchScope sets the search scope for a query instance.
//
// See: https://developer.apple.com/documentation/coreservices/1413048-mdquerysetsearchscope
func MDQuerySetSearchScope(query MDQueryRef, scopeDirectories corefoundation.CFArrayRef, scopeOptions uintptr) {
	if _mDQuerySetSearchScope == nil {
		panic("coreservices: symbol MDQuerySetSearchScope not loaded")
	}
	_mDQuerySetSearchScope(query, scopeDirectories, scopeOptions)
}

var _mDQuerySetSortComparator func(query MDQueryRef, comparator MDQuerySortComparatorFunction, context unsafe.Pointer)

// MDQuerySetSortComparator sets the function used to sort the results of an MDQuery.
//
// See: https://developer.apple.com/documentation/coreservices/1413087-mdquerysetsortcomparator
func MDQuerySetSortComparator(query MDQueryRef, comparator MDQuerySortComparatorFunction, context unsafe.Pointer) {
	if _mDQuerySetSortComparator == nil {
		panic("coreservices: symbol MDQuerySetSortComparator not loaded")
	}
	_mDQuerySetSortComparator(query, comparator, context)
}

var _mDQuerySetSortComparatorBlock func(query MDQueryRef, comparator *corefoundation.CFTypeRef, arg2 *corefoundation.CFTypeRef) corefoundation.CFComparisonResult

// MDQuerySetSortComparatorBlock sets the block used to sort the results of an MDQuery.
//
// See: https://developer.apple.com/documentation/coreservices/1413021-mdquerysetsortcomparatorblock
func MDQuerySetSortComparatorBlock(query MDQueryRef, comparator *corefoundation.CFTypeRef, arg2 *corefoundation.CFTypeRef) corefoundation.CFComparisonResult {
	if _mDQuerySetSortComparatorBlock == nil {
		panic("coreservices: symbol MDQuerySetSortComparatorBlock not loaded")
	}
	return _mDQuerySetSortComparatorBlock(query, comparator, arg2)
}

var _mDQuerySetSortOptionFlagsForAttribute func(query MDQueryRef, fieldName corefoundation.CFStringRef, flags uint32) bool

// MDQuerySetSortOptionFlagsForAttribute.
//
// See: https://developer.apple.com/documentation/coreservices/1413075-mdquerysetsortoptionflagsforattr
func MDQuerySetSortOptionFlagsForAttribute(query MDQueryRef, fieldName corefoundation.CFStringRef, flags uint32) bool {
	if _mDQuerySetSortOptionFlagsForAttribute == nil {
		panic("coreservices: symbol MDQuerySetSortOptionFlagsForAttribute not loaded")
	}
	return _mDQuerySetSortOptionFlagsForAttribute(query, fieldName, flags)
}

var _mDQuerySetSortOrder func(query MDQueryRef, sortingAttrs corefoundation.CFArrayRef) bool

// MDQuerySetSortOrder.
//
// See: https://developer.apple.com/documentation/coreservices/1413096-mdquerysetsortorder
func MDQuerySetSortOrder(query MDQueryRef, sortingAttrs corefoundation.CFArrayRef) bool {
	if _mDQuerySetSortOrder == nil {
		panic("coreservices: symbol MDQuerySetSortOrder not loaded")
	}
	return _mDQuerySetSortOrder(query, sortingAttrs)
}

var _mDQueryStop func(query MDQueryRef)

// MDQueryStop stops the query from generating more results.
//
// See: https://developer.apple.com/documentation/coreservices/1413077-mdquerystop
func MDQueryStop(query MDQueryRef) {
	if _mDQueryStop == nil {
		panic("coreservices: symbol MDQueryStop not loaded")
	}
	_mDQueryStop(query)
}

var _mDSchemaCopyAllAttributes func() corefoundation.CFArrayRef

// MDSchemaCopyAllAttributes returns an array containing all the metadata attributesdefined in the schema.
//
// See: https://developer.apple.com/documentation/coreservices/1445665-mdschemacopyallattributes
func MDSchemaCopyAllAttributes() corefoundation.CFArrayRef {
	if _mDSchemaCopyAllAttributes == nil {
		panic("coreservices: symbol MDSchemaCopyAllAttributes not loaded")
	}
	return _mDSchemaCopyAllAttributes()
}

var _mDSchemaCopyAttributesForContentType func(contentTypeUTI corefoundation.CFStringRef) corefoundation.CFDictionaryRef

// MDSchemaCopyAttributesForContentType returns a dictionary containing the metadata attributesfor the specified UTI type.
//
// See: https://developer.apple.com/documentation/coreservices/1444459-mdschemacopyattributesforcontent
func MDSchemaCopyAttributesForContentType(contentTypeUTI corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	if _mDSchemaCopyAttributesForContentType == nil {
		panic("coreservices: symbol MDSchemaCopyAttributesForContentType not loaded")
	}
	return _mDSchemaCopyAttributesForContentType(contentTypeUTI)
}

var _mDSchemaCopyDisplayDescriptionForAttribute func(name corefoundation.CFStringRef) corefoundation.CFStringRef

// MDSchemaCopyDisplayDescriptionForAttribute returns the localized description of a metadata attributekey.
//
// See: https://developer.apple.com/documentation/coreservices/1442582-mdschemacopydisplaydescriptionfo
func MDSchemaCopyDisplayDescriptionForAttribute(name corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _mDSchemaCopyDisplayDescriptionForAttribute == nil {
		panic("coreservices: symbol MDSchemaCopyDisplayDescriptionForAttribute not loaded")
	}
	return _mDSchemaCopyDisplayDescriptionForAttribute(name)
}

var _mDSchemaCopyDisplayNameForAttribute func(name corefoundation.CFStringRef) corefoundation.CFStringRef

// MDSchemaCopyDisplayNameForAttribute returns the localized display name of a metadata attributekey.
//
// See: https://developer.apple.com/documentation/coreservices/1450203-mdschemacopydisplaynameforattrib
func MDSchemaCopyDisplayNameForAttribute(name corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _mDSchemaCopyDisplayNameForAttribute == nil {
		panic("coreservices: symbol MDSchemaCopyDisplayNameForAttribute not loaded")
	}
	return _mDSchemaCopyDisplayNameForAttribute(name)
}

var _mDSchemaCopyMetaAttributesForAttribute func(name corefoundation.CFStringRef) corefoundation.CFDictionaryRef

// MDSchemaCopyMetaAttributesForAttribute returns a dictionary describing the values for the specifiedmetadata attribute key.
//
// See: https://developer.apple.com/documentation/coreservices/1450052-mdschemacopymetaattributesforatt
func MDSchemaCopyMetaAttributesForAttribute(name corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	if _mDSchemaCopyMetaAttributesForAttribute == nil {
		panic("coreservices: symbol MDSchemaCopyMetaAttributesForAttribute not loaded")
	}
	return _mDSchemaCopyMetaAttributesForAttribute(name)
}

var _mPAllocateAligned func(arg0 uintptr, arg1 uint8, arg2 uintptr) unsafe.Pointer

// MPAllocateAligned allocates a nonrelocatable memory block.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585774-mpallocatealigned
func MPAllocateAligned(arg0 uintptr, arg1 uint8, arg2 uintptr) unsafe.Pointer {
	if _mPAllocateAligned == nil {
		panic("coreservices: symbol MPAllocateAligned not loaded")
	}
	return _mPAllocateAligned(arg0, arg1, arg2)
}

var _mPAllocateTaskStorageIndex func(arg0 TaskStorageIndex) int32

// MPAllocateTaskStorageIndex returns an index number to access per-task storage.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585719-mpallocatetaskstorageindex
func MPAllocateTaskStorageIndex(arg0 TaskStorageIndex) int32 {
	if _mPAllocateTaskStorageIndex == nil {
		panic("coreservices: symbol MPAllocateTaskStorageIndex not loaded")
	}
	return _mPAllocateTaskStorageIndex(arg0)
}

var _mPArmTimer func(arg0 MPTimerID, arg1 unsafe.Pointer, arg2 uintptr) int32

// MPArmTimer arms the timer to expire at a given time.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585612-mparmtimer
func MPArmTimer(arg0 MPTimerID, arg1 unsafe.Pointer, arg2 uintptr) int32 {
	if _mPArmTimer == nil {
		panic("coreservices: symbol MPArmTimer not loaded")
	}
	return _mPArmTimer(arg0, arg1, arg2)
}

var _mPBlockClear func(arg0 unsafe.Pointer, arg1 uintptr)

// MPBlockClear clears a block of memory.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585642-mpblockclear
func MPBlockClear(arg0 unsafe.Pointer, arg1 uintptr) {
	if _mPBlockClear == nil {
		panic("coreservices: symbol MPBlockClear not loaded")
	}
	_mPBlockClear(arg0, arg1)
}

var _mPBlockCopy func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uintptr)

// MPBlockCopy copies a block of memory.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585707-mpblockcopy
func MPBlockCopy(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 uintptr) {
	if _mPBlockCopy == nil {
		panic("coreservices: symbol MPBlockCopy not loaded")
	}
	_mPBlockCopy(arg0, arg1, arg2)
}

var _mPCancelTimer func(arg0 MPTimerID, arg1 unsafe.Pointer) int32

// MPCancelTimer cancels an armed timer.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585745-mpcanceltimer
func MPCancelTimer(arg0 MPTimerID, arg1 unsafe.Pointer) int32 {
	if _mPCancelTimer == nil {
		panic("coreservices: symbol MPCancelTimer not loaded")
	}
	return _mPCancelTimer(arg0, arg1)
}

var _mPCauseNotification func(arg0 MPNotificationID) int32

// MPCauseNotification signals a kernel notification.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585754-mpcausenotification
func MPCauseNotification(arg0 MPNotificationID) int32 {
	if _mPCauseNotification == nil {
		panic("coreservices: symbol MPCauseNotification not loaded")
	}
	return _mPCauseNotification(arg0)
}

var _mPCreateCriticalRegion func(arg0 MPCriticalRegionID) int32

// MPCreateCriticalRegion creates a critical region object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585663-mpcreatecriticalregion
func MPCreateCriticalRegion(arg0 MPCriticalRegionID) int32 {
	if _mPCreateCriticalRegion == nil {
		panic("coreservices: symbol MPCreateCriticalRegion not loaded")
	}
	return _mPCreateCriticalRegion(arg0)
}

var _mPCreateEvent func(arg0 MPEventID) int32

// MPCreateEvent creates an event group.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585702-mpcreateevent
func MPCreateEvent(arg0 MPEventID) int32 {
	if _mPCreateEvent == nil {
		panic("coreservices: symbol MPCreateEvent not loaded")
	}
	return _mPCreateEvent(arg0)
}

var _mPCreateNotification func(arg0 MPNotificationID) int32

// MPCreateNotification creates a kernel notification
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585723-mpcreatenotification
func MPCreateNotification(arg0 MPNotificationID) int32 {
	if _mPCreateNotification == nil {
		panic("coreservices: symbol MPCreateNotification not loaded")
	}
	return _mPCreateNotification(arg0)
}

var _mPCreateQueue func(arg0 MPQueueID) int32

// MPCreateQueue creates a message queue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585694-mpcreatequeue
func MPCreateQueue(arg0 MPQueueID) int32 {
	if _mPCreateQueue == nil {
		panic("coreservices: symbol MPCreateQueue not loaded")
	}
	return _mPCreateQueue(arg0)
}

var _mPCreateSemaphore func(arg0 MPSemaphoreCount, arg1 MPSemaphoreCount, arg2 MPSemaphoreID) int32

// MPCreateSemaphore creates a semaphore.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585569-mpcreatesemaphore
func MPCreateSemaphore(arg0 MPSemaphoreCount, arg1 MPSemaphoreCount, arg2 MPSemaphoreID) int32 {
	if _mPCreateSemaphore == nil {
		panic("coreservices: symbol MPCreateSemaphore not loaded")
	}
	return _mPCreateSemaphore(arg0, arg1, arg2)
}

var _mPCreateTask func(arg0 unsafe.Pointer, arg1 uintptr, arg2 MPQueueID, arg3 MPTaskOptions, arg4 MPTaskID) int32

// MPCreateTask creates a preemptive task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585779-mpcreatetask
func MPCreateTask(arg0 unsafe.Pointer, arg1 uintptr, arg2 MPQueueID, arg3 MPTaskOptions, arg4 MPTaskID) int32 {
	if _mPCreateTask == nil {
		panic("coreservices: symbol MPCreateTask not loaded")
	}
	return _mPCreateTask(arg0, arg1, arg2, arg3, arg4)
}

var _mPCreateTimer func(arg0 MPTimerID) int32

// MPCreateTimer creates a timer.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585748-mpcreatetimer
func MPCreateTimer(arg0 MPTimerID) int32 {
	if _mPCreateTimer == nil {
		panic("coreservices: symbol MPCreateTimer not loaded")
	}
	return _mPCreateTimer(arg0)
}

var _mPCurrentTaskID func() MPTaskID

// MPCurrentTaskID obtains the task ID of the currently-executing preemptive task
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585673-mpcurrenttaskid
func MPCurrentTaskID() MPTaskID {
	if _mPCurrentTaskID == nil {
		panic("coreservices: symbol MPCurrentTaskID not loaded")
	}
	return _mPCurrentTaskID()
}

var _mPDeallocateTaskStorageIndex func(arg0 TaskStorageIndex) int32

// MPDeallocateTaskStorageIndex frees an index number used to access per-task storage
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585649-mpdeallocatetaskstorageindex
func MPDeallocateTaskStorageIndex(arg0 TaskStorageIndex) int32 {
	if _mPDeallocateTaskStorageIndex == nil {
		panic("coreservices: symbol MPDeallocateTaskStorageIndex not loaded")
	}
	return _mPDeallocateTaskStorageIndex(arg0)
}

var _mPDelayUntil func(arg0 unsafe.Pointer) int32

// MPDelayUntil blocks the calling task until a specified time.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585647-mpdelayuntil
func MPDelayUntil(arg0 unsafe.Pointer) int32 {
	if _mPDelayUntil == nil {
		panic("coreservices: symbol MPDelayUntil not loaded")
	}
	return _mPDelayUntil(arg0)
}

var _mPDeleteCriticalRegion func(arg0 MPCriticalRegionID) int32

// MPDeleteCriticalRegion removes the specified critical region object.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585704-mpdeletecriticalregion
func MPDeleteCriticalRegion(arg0 MPCriticalRegionID) int32 {
	if _mPDeleteCriticalRegion == nil {
		panic("coreservices: symbol MPDeleteCriticalRegion not loaded")
	}
	return _mPDeleteCriticalRegion(arg0)
}

var _mPDeleteEvent func(arg0 MPEventID) int32

// MPDeleteEvent removes an event group.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585691-mpdeleteevent
func MPDeleteEvent(arg0 MPEventID) int32 {
	if _mPDeleteEvent == nil {
		panic("coreservices: symbol MPDeleteEvent not loaded")
	}
	return _mPDeleteEvent(arg0)
}

var _mPDeleteNotification func(arg0 MPNotificationID) int32

// MPDeleteNotification removes a kernel notification.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585659-mpdeletenotification
func MPDeleteNotification(arg0 MPNotificationID) int32 {
	if _mPDeleteNotification == nil {
		panic("coreservices: symbol MPDeleteNotification not loaded")
	}
	return _mPDeleteNotification(arg0)
}

var _mPDeleteQueue func(arg0 MPQueueID) int32

// MPDeleteQueue deletes a message queue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585571-mpdeletequeue
func MPDeleteQueue(arg0 MPQueueID) int32 {
	if _mPDeleteQueue == nil {
		panic("coreservices: symbol MPDeleteQueue not loaded")
	}
	return _mPDeleteQueue(arg0)
}

var _mPDeleteSemaphore func(arg0 MPSemaphoreID) int32

// MPDeleteSemaphore removes a semaphore.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585586-mpdeletesemaphore
func MPDeleteSemaphore(arg0 MPSemaphoreID) int32 {
	if _mPDeleteSemaphore == nil {
		panic("coreservices: symbol MPDeleteSemaphore not loaded")
	}
	return _mPDeleteSemaphore(arg0)
}

var _mPDeleteTimer func(arg0 MPTimerID) int32

// MPDeleteTimer removes a timer.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585761-mpdeletetimer
func MPDeleteTimer(arg0 MPTimerID) int32 {
	if _mPDeleteTimer == nil {
		panic("coreservices: symbol MPDeleteTimer not loaded")
	}
	return _mPDeleteTimer(arg0)
}

var _mPDisposeTaskException func(arg0 MPTaskID, arg1 uintptr) int32

// MPDisposeTaskException removes a task exception.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585607-mpdisposetaskexception
func MPDisposeTaskException(arg0 MPTaskID, arg1 uintptr) int32 {
	if _mPDisposeTaskException == nil {
		panic("coreservices: symbol MPDisposeTaskException not loaded")
	}
	return _mPDisposeTaskException(arg0, arg1)
}

var _mPEnterCriticalRegion func(arg0 MPCriticalRegionID, arg1 unsafe.Pointer) int32

// MPEnterCriticalRegion attempts to enter a critical region.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585622-mpentercriticalregion
func MPEnterCriticalRegion(arg0 MPCriticalRegionID, arg1 unsafe.Pointer) int32 {
	if _mPEnterCriticalRegion == nil {
		panic("coreservices: symbol MPEnterCriticalRegion not loaded")
	}
	return _mPEnterCriticalRegion(arg0, arg1)
}

var _mPExit func(arg0 int32)

// MPExit allows a task to terminate itself
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585705-mpexit
func MPExit(arg0 int32) {
	if _mPExit == nil {
		panic("coreservices: symbol MPExit not loaded")
	}
	_mPExit(arg0)
}

var _mPExitCriticalRegion func(arg0 MPCriticalRegionID) int32

// MPExitCriticalRegion exits a critical region.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585758-mpexitcriticalregion
func MPExitCriticalRegion(arg0 MPCriticalRegionID) int32 {
	if _mPExitCriticalRegion == nil {
		panic("coreservices: symbol MPExitCriticalRegion not loaded")
	}
	return _mPExitCriticalRegion(arg0)
}

var _mPExtractTaskState func(arg0 MPTaskID, arg1 MPTaskStateKind) int32

// MPExtractTaskState extracts state information from a suspended task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585718-mpextracttaskstate
func MPExtractTaskState(arg0 MPTaskID, arg1 MPTaskStateKind) int32 {
	if _mPExtractTaskState == nil {
		panic("coreservices: symbol MPExtractTaskState not loaded")
	}
	return _mPExtractTaskState(arg0, arg1)
}

var _mPFree func(arg0 unsafe.Pointer)

// MPFree frees memory allocated by [MPAllocateAligned].
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585676-mpfree
func MPFree(arg0 unsafe.Pointer) {
	if _mPFree == nil {
		panic("coreservices: symbol MPFree not loaded")
	}
	_mPFree(arg0)
}

var _mPGetAllocatedBlockSize func(arg0 unsafe.Pointer) objectivec.IObject

// MPGetAllocatedBlockSize returns the size of a memory block.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585717-mpgetallocatedblocksize
func MPGetAllocatedBlockSize(arg0 unsafe.Pointer) objectivec.IObject {
	if _mPGetAllocatedBlockSize == nil {
		panic("coreservices: symbol MPGetAllocatedBlockSize not loaded")
	}
	return _mPGetAllocatedBlockSize(arg0)
}

var _mPGetNextCpuID func(arg0 MPCoherenceID, arg1 MPCpuID) int32

// MPGetNextCpuID obtains the next CPU ID in the list of physical processors of the specified memory coherence group.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1508189-mpgetnextcpuid
func MPGetNextCpuID(arg0 MPCoherenceID, arg1 MPCpuID) int32 {
	if _mPGetNextCpuID == nil {
		panic("coreservices: symbol MPGetNextCpuID not loaded")
	}
	return _mPGetNextCpuID(arg0, arg1)
}

var _mPGetNextTaskID func(arg0 MPProcessID, arg1 MPTaskID) int32

// MPGetNextTaskID obtains the next task ID in the list of available tasks.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1508163-mpgetnexttaskid
func MPGetNextTaskID(arg0 MPProcessID, arg1 MPTaskID) int32 {
	if _mPGetNextTaskID == nil {
		panic("coreservices: symbol MPGetNextTaskID not loaded")
	}
	return _mPGetNextTaskID(arg0, arg1)
}

var _mPGetTaskStorageValue func(arg0 TaskStorageIndex) TaskStorageValue

// MPGetTaskStorageValue gets the storage value stored at a specified index number.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585589-mpgettaskstoragevalue
func MPGetTaskStorageValue(arg0 TaskStorageIndex) TaskStorageValue {
	if _mPGetTaskStorageValue == nil {
		panic("coreservices: symbol MPGetTaskStorageValue not loaded")
	}
	return _mPGetTaskStorageValue(arg0)
}

var _mPModifyNotification func(arg0 MPNotificationID, arg1 MPOpaqueID) int32

// MPModifyNotification adds a simple notification to a kernel notification.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585780-mpmodifynotification
func MPModifyNotification(arg0 MPNotificationID, arg1 MPOpaqueID) int32 {
	if _mPModifyNotification == nil {
		panic("coreservices: symbol MPModifyNotification not loaded")
	}
	return _mPModifyNotification(arg0, arg1)
}

var _mPModifyNotificationParameters func(arg0 MPNotificationID, arg1 MPOpaqueIDClass) int32

// MPModifyNotificationParameters.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585668-mpmodifynotificationparameters
func MPModifyNotificationParameters(arg0 MPNotificationID, arg1 MPOpaqueIDClass) int32 {
	if _mPModifyNotificationParameters == nil {
		panic("coreservices: symbol MPModifyNotificationParameters not loaded")
	}
	return _mPModifyNotificationParameters(arg0, arg1)
}

var _mPNotifyQueue func(arg0 MPQueueID) int32

// MPNotifyQueue sends a message to the specified message queue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585699-mpnotifyqueue
func MPNotifyQueue(arg0 MPQueueID) int32 {
	if _mPNotifyQueue == nil {
		panic("coreservices: symbol MPNotifyQueue not loaded")
	}
	return _mPNotifyQueue(arg0)
}

var _mPProcessors func() objectivec.IObject

// MPProcessors returns the number of processors on the host computer.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585778-mpprocessors
func MPProcessors() objectivec.IObject {
	if _mPProcessors == nil {
		panic("coreservices: symbol MPProcessors not loaded")
	}
	return _mPProcessors()
}

var _mPProcessorsScheduled func() objectivec.IObject

// MPProcessorsScheduled returns the number of active processors available on the host computer.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585777-mpprocessorsscheduled
func MPProcessorsScheduled() objectivec.IObject {
	if _mPProcessorsScheduled == nil {
		panic("coreservices: symbol MPProcessorsScheduled not loaded")
	}
	return _mPProcessorsScheduled()
}

var _mPRegisterDebugger func(arg0 MPQueueID, arg1 MPDebuggerLevel) int32

// MPRegisterDebugger registers a debugger.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585573-mpregisterdebugger
func MPRegisterDebugger(arg0 MPQueueID, arg1 MPDebuggerLevel) int32 {
	if _mPRegisterDebugger == nil {
		panic("coreservices: symbol MPRegisterDebugger not loaded")
	}
	return _mPRegisterDebugger(arg0, arg1)
}

var _mPRemoteCall func(arg0 unsafe.Pointer, arg1 MPRemoteContext) unsafe.Pointer

// MPRemoteCall calls a non-reentrant function and blocks the current task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585652-mpremotecall
func MPRemoteCall(arg0 unsafe.Pointer, arg1 MPRemoteContext) unsafe.Pointer {
	if _mPRemoteCall == nil {
		panic("coreservices: symbol MPRemoteCall not loaded")
	}
	return _mPRemoteCall(arg0, arg1)
}

var _mPRemoteCallCFM func(arg0 unsafe.Pointer, arg1 MPRemoteContext) unsafe.Pointer

// MPRemoteCallCFM calls a non-reentrant function and blocks the current task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585757-mpremotecallcfm
func MPRemoteCallCFM(arg0 unsafe.Pointer, arg1 MPRemoteContext) unsafe.Pointer {
	if _mPRemoteCallCFM == nil {
		panic("coreservices: symbol MPRemoteCallCFM not loaded")
	}
	return _mPRemoteCallCFM(arg0, arg1)
}

var _mPSetEvent func(arg0 MPEventID, arg1 MPEventFlags) int32

// MPSetEvent merges event flags into a specified event group.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585752-mpsetevent
func MPSetEvent(arg0 MPEventID, arg1 MPEventFlags) int32 {
	if _mPSetEvent == nil {
		panic("coreservices: symbol MPSetEvent not loaded")
	}
	return _mPSetEvent(arg0, arg1)
}

var _mPSetExceptionHandler func(arg0 MPTaskID, arg1 MPQueueID) int32

// MPSetExceptionHandler sets an exception handler for a task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585759-mpsetexceptionhandler
func MPSetExceptionHandler(arg0 MPTaskID, arg1 MPQueueID) int32 {
	if _mPSetExceptionHandler == nil {
		panic("coreservices: symbol MPSetExceptionHandler not loaded")
	}
	return _mPSetExceptionHandler(arg0, arg1)
}

var _mPSetQueueReserve func(arg0 MPQueueID, arg1 uintptr) int32

// MPSetQueueReserve reserves space for messages on a specified message queue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585671-mpsetqueuereserve
func MPSetQueueReserve(arg0 MPQueueID, arg1 uintptr) int32 {
	if _mPSetQueueReserve == nil {
		panic("coreservices: symbol MPSetQueueReserve not loaded")
	}
	return _mPSetQueueReserve(arg0, arg1)
}

var _mPSetTaskState func(arg0 MPTaskID, arg1 MPTaskStateKind) int32

// MPSetTaskState sets state information for a suspended task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585601-mpsettaskstate
func MPSetTaskState(arg0 MPTaskID, arg1 MPTaskStateKind) int32 {
	if _mPSetTaskState == nil {
		panic("coreservices: symbol MPSetTaskState not loaded")
	}
	return _mPSetTaskState(arg0, arg1)
}

var _mPSetTaskStorageValue func(arg0 TaskStorageIndex, arg1 TaskStorageValue) int32

// MPSetTaskStorageValue sets the storage value for a given index number.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585626-mpsettaskstoragevalue
func MPSetTaskStorageValue(arg0 TaskStorageIndex, arg1 TaskStorageValue) int32 {
	if _mPSetTaskStorageValue == nil {
		panic("coreservices: symbol MPSetTaskStorageValue not loaded")
	}
	return _mPSetTaskStorageValue(arg0, arg1)
}

var _mPSetTaskType func(arg0 MPTaskID, arg1 uint32) int32

// MPSetTaskType sets the type of the task.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1585695-mpsettasktype
func MPSetTaskType(arg0 MPTaskID, arg1 uint32) int32 {
	if _mPSetTaskType == nil {
		panic("coreservices: symbol MPSetTaskType not loaded")
	}
	return _mPSetTaskType(arg0, arg1)
}

var _mPSetTaskWeight func(arg0 MPTaskID, arg1 MPTaskWeight) int32

// MPSetTaskWeight assigns a relative weight to a task, indicating how much processor time it should receive compared to other available tasks.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585665-mpsettaskweight
func MPSetTaskWeight(arg0 MPTaskID, arg1 MPTaskWeight) int32 {
	if _mPSetTaskWeight == nil {
		panic("coreservices: symbol MPSetTaskWeight not loaded")
	}
	return _mPSetTaskWeight(arg0, arg1)
}

var _mPSetTimerNotify func(arg0 MPTimerID, arg1 MPOpaqueID) int32

// MPSetTimerNotify sets the notification information associated with a timer.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585726-mpsettimernotify
func MPSetTimerNotify(arg0 MPTimerID, arg1 MPOpaqueID) int32 {
	if _mPSetTimerNotify == nil {
		panic("coreservices: symbol MPSetTimerNotify not loaded")
	}
	return _mPSetTimerNotify(arg0, arg1)
}

var _mPSignalSemaphore func(arg0 MPSemaphoreID) int32

// MPSignalSemaphore signals a semaphore.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585713-mpsignalsemaphore
func MPSignalSemaphore(arg0 MPSemaphoreID) int32 {
	if _mPSignalSemaphore == nil {
		panic("coreservices: symbol MPSignalSemaphore not loaded")
	}
	return _mPSignalSemaphore(arg0)
}

var _mPTaskIsPreemptive func(arg0 MPTaskID) bool

// MPTaskIsPreemptive determines whether a task is preemptively scheduled.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585681-mptaskispreemptive
func MPTaskIsPreemptive(arg0 MPTaskID) bool {
	if _mPTaskIsPreemptive == nil {
		panic("coreservices: symbol MPTaskIsPreemptive not loaded")
	}
	return _mPTaskIsPreemptive(arg0)
}

var _mPTerminateTask func(arg0 MPTaskID, arg1 int32) int32

// MPTerminateTask terminates an existing task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585769-mpterminatetask
func MPTerminateTask(arg0 MPTaskID, arg1 int32) int32 {
	if _mPTerminateTask == nil {
		panic("coreservices: symbol MPTerminateTask not loaded")
	}
	return _mPTerminateTask(arg0, arg1)
}

var _mPUnregisterDebugger func(arg0 MPQueueID) int32

// MPUnregisterDebugger unregisters a debugger.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585598-mpunregisterdebugger
func MPUnregisterDebugger(arg0 MPQueueID) int32 {
	if _mPUnregisterDebugger == nil {
		panic("coreservices: symbol MPUnregisterDebugger not loaded")
	}
	return _mPUnregisterDebugger(arg0)
}

var _mPWaitForEvent func(arg0 MPEventID, arg1 MPEventFlags, arg2 unsafe.Pointer) int32

// MPWaitForEvent retrieves event flags from a specified event group.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585656-mpwaitforevent
func MPWaitForEvent(arg0 MPEventID, arg1 MPEventFlags, arg2 unsafe.Pointer) int32 {
	if _mPWaitForEvent == nil {
		panic("coreservices: symbol MPWaitForEvent not loaded")
	}
	return _mPWaitForEvent(arg0, arg1, arg2)
}

var _mPWaitOnQueue func(arg0 MPQueueID, arg1 unsafe.Pointer) int32

// MPWaitOnQueue obtains a message from a specified message queue.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585762-mpwaitonqueue
func MPWaitOnQueue(arg0 MPQueueID, arg1 unsafe.Pointer) int32 {
	if _mPWaitOnQueue == nil {
		panic("coreservices: symbol MPWaitOnQueue not loaded")
	}
	return _mPWaitOnQueue(arg0, arg1)
}

var _mPWaitOnSemaphore func(arg0 MPSemaphoreID, arg1 unsafe.Pointer) int32

// MPWaitOnSemaphore waits on a semaphore
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585722-mpwaitonsemaphore
func MPWaitOnSemaphore(arg0 MPSemaphoreID, arg1 unsafe.Pointer) int32 {
	if _mPWaitOnSemaphore == nil {
		panic("coreservices: symbol MPWaitOnSemaphore not loaded")
	}
	return _mPWaitOnSemaphore(arg0, arg1)
}

var _mPYield func()

// MPYield allows a task to yield the processor to another task.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1585732-mpyield
func MPYield() {
	if _mPYield == nil {
		panic("coreservices: symbol MPYield not loaded")
	}
	_mPYield()
}

var _maximumProcessorSpeed func() int16

// MaximumProcessorSpeed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427409-maximumprocessorspeed
func MaximumProcessorSpeed() int16 {
	if _maximumProcessorSpeed == nil {
		panic("coreservices: symbol MaximumProcessorSpeed not loaded")
	}
	return _maximumProcessorSpeed()
}

var _memError func() int16

// MemError.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506254-memerror
func MemError() int16 {
	if _memError == nil {
		panic("coreservices: symbol MemError not loaded")
	}
	return _memError()
}

var _microseconds func(arg0 uint64)

// Microseconds.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1550773-microseconds
func Microseconds(arg0 uint64) {
	if _microseconds == nil {
		panic("coreservices: symbol Microseconds not loaded")
	}
	_microseconds(arg0)
}

var _minimumProcessorSpeed func() int16

// MinimumProcessorSpeed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427399-minimumprocessorspeed
func MinimumProcessorSpeed() int16 {
	if _minimumProcessorSpeed == nil {
		panic("coreservices: symbol MinimumProcessorSpeed not loaded")
	}
	return _minimumProcessorSpeed()
}

var _munger func(arg0 uintptr, arg1 int, arg2 int, arg3 int) int

// Munger.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1588638-munger
func Munger(arg0 uintptr, arg1 int, arg2 int, arg3 int) int {
	if _munger == nil {
		panic("coreservices: symbol Munger not loaded")
	}
	return _munger(arg0, arg1, arg2, arg3)
}

var _nanosecondsToAbsolute func(arg0 Nanoseconds) unsafe.Pointer

// NanosecondsToAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501255-nanosecondstoabsolute
func NanosecondsToAbsolute(arg0 Nanoseconds) unsafe.Pointer {
	if _nanosecondsToAbsolute == nil {
		panic("coreservices: symbol NanosecondsToAbsolute not loaded")
	}
	return _nanosecondsToAbsolute(arg0)
}

var _nanosecondsToDuration func(arg0 Nanoseconds) unsafe.Pointer

// NanosecondsToDuration.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501257-nanosecondstoduration
func NanosecondsToDuration(arg0 Nanoseconds) unsafe.Pointer {
	if _nanosecondsToDuration == nil {
		panic("coreservices: symbol NanosecondsToDuration not loaded")
	}
	return _nanosecondsToDuration(arg0)
}

var _nearestMacTextEncodings func(arg0 TextEncoding, arg1 TextEncoding, arg2 TextEncoding) int32

// NearestMacTextEncodings obtains the best and alternate Mac text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1399736-nearestmactextencodings
func NearestMacTextEncodings(arg0 TextEncoding, arg1 TextEncoding, arg2 TextEncoding) int32 {
	if _nearestMacTextEncodings == nil {
		panic("coreservices: symbol NearestMacTextEncodings not loaded")
	}
	return _nearestMacTextEncodings(arg0, arg1, arg2)
}

var _newAECoerceDescUPP func(userRoutine AECoerceDescProcPtr) AECoerceDescUPP

// NewAECoerceDescUPP creates a new universal procedure pointer to a function that coerces data stored in a descriptor.
//
// See: https://developer.apple.com/documentation/coreservices/1445885-newaecoercedescupp
func NewAECoerceDescUPP(userRoutine AECoerceDescProcPtr) AECoerceDescUPP {
	if _newAECoerceDescUPP == nil {
		panic("coreservices: symbol NewAECoerceDescUPP not loaded")
	}
	return _newAECoerceDescUPP(userRoutine)
}

var _newAECoercePtrUPP func(userRoutine AECoercePtrProcPtr) AECoercePtrUPP

// NewAECoercePtrUPP creates a new universal procedure pointer to a function that coerces data stored in a buffer.
//
// See: https://developer.apple.com/documentation/coreservices/1449962-newaecoerceptrupp
func NewAECoercePtrUPP(userRoutine AECoercePtrProcPtr) AECoercePtrUPP {
	if _newAECoercePtrUPP == nil {
		panic("coreservices: symbol NewAECoercePtrUPP not loaded")
	}
	return _newAECoercePtrUPP(userRoutine)
}

var _newAEDisposeExternalUPP func(userRoutine AEDisposeExternalProcPtr) AEDisposeExternalUPP

// NewAEDisposeExternalUPP creates a new universal procedure pointer to a function that disposes of data stored in a buffer.
//
// See: https://developer.apple.com/documentation/coreservices/1447774-newaedisposeexternalupp
func NewAEDisposeExternalUPP(userRoutine AEDisposeExternalProcPtr) AEDisposeExternalUPP {
	if _newAEDisposeExternalUPP == nil {
		panic("coreservices: symbol NewAEDisposeExternalUPP not loaded")
	}
	return _newAEDisposeExternalUPP(userRoutine)
}

var _newAEEventHandlerUPP func(userRoutine AEEventHandlerProcPtr) AEEventHandlerUPP

// NewAEEventHandlerUPP creates a new universal procedure pointer to an event handler function.
//
// See: https://developer.apple.com/documentation/coreservices/1446862-newaeeventhandlerupp
func NewAEEventHandlerUPP(userRoutine AEEventHandlerProcPtr) AEEventHandlerUPP {
	if _newAEEventHandlerUPP == nil {
		panic("coreservices: symbol NewAEEventHandlerUPP not loaded")
	}
	return _newAEEventHandlerUPP(userRoutine)
}

var _newCollection func() Collection

// NewCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551416-newcollection
func NewCollection() Collection {
	if _newCollection == nil {
		panic("coreservices: symbol NewCollection not loaded")
	}
	return _newCollection()
}

var _newCollectionExceptionUPP func(arg0 unsafe.Pointer) CollectionExceptionUPP

// NewCollectionExceptionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551448-newcollectionexceptionupp
func NewCollectionExceptionUPP(arg0 unsafe.Pointer) CollectionExceptionUPP {
	if _newCollectionExceptionUPP == nil {
		panic("coreservices: symbol NewCollectionExceptionUPP not loaded")
	}
	return _newCollectionExceptionUPP(arg0)
}

var _newCollectionFlattenUPP func(arg0 unsafe.Pointer) CollectionFlattenUPP

// NewCollectionFlattenUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551360-newcollectionflattenupp
func NewCollectionFlattenUPP(arg0 unsafe.Pointer) CollectionFlattenUPP {
	if _newCollectionFlattenUPP == nil {
		panic("coreservices: symbol NewCollectionFlattenUPP not loaded")
	}
	return _newCollectionFlattenUPP(arg0)
}

var _newComponentFunctionUPP func(arg0 unsafe.Pointer, arg1 ProcInfoType) ComponentFunctionUPP

// NewComponentFunctionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516583-newcomponentfunctionupp
func NewComponentFunctionUPP(arg0 unsafe.Pointer, arg1 ProcInfoType) ComponentFunctionUPP {
	if _newComponentFunctionUPP == nil {
		panic("coreservices: symbol NewComponentFunctionUPP not loaded")
	}
	return _newComponentFunctionUPP(arg0, arg1)
}

var _newComponentMPWorkFunctionUPP func(arg0 unsafe.Pointer) ComponentMPWorkFunctionUPP

// NewComponentMPWorkFunctionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516448-newcomponentmpworkfunctionupp
func NewComponentMPWorkFunctionUPP(arg0 unsafe.Pointer) ComponentMPWorkFunctionUPP {
	if _newComponentMPWorkFunctionUPP == nil {
		panic("coreservices: symbol NewComponentMPWorkFunctionUPP not loaded")
	}
	return _newComponentMPWorkFunctionUPP(arg0)
}

var _newComponentRoutineUPP func(arg0 unsafe.Pointer) ComponentRoutineUPP

// NewComponentRoutineUPP creates a new universal procedure pointer (UPP) to a component routine callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516579-newcomponentroutineupp
func NewComponentRoutineUPP(arg0 unsafe.Pointer) ComponentRoutineUPP {
	if _newComponentRoutineUPP == nil {
		panic("coreservices: symbol NewComponentRoutineUPP not loaded")
	}
	return _newComponentRoutineUPP(arg0)
}

var _newDebugAssertOutputHandlerUPP func(arg0 unsafe.Pointer) DebugAssertOutputHandlerUPP

// NewDebugAssertOutputHandlerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517735-newdebugassertoutputhandlerupp
func NewDebugAssertOutputHandlerUPP(arg0 unsafe.Pointer) DebugAssertOutputHandlerUPP {
	if _newDebugAssertOutputHandlerUPP == nil {
		panic("coreservices: symbol NewDebugAssertOutputHandlerUPP not loaded")
	}
	return _newDebugAssertOutputHandlerUPP(arg0)
}

var _newDebugComponent func(arg0 uint32, arg1 unsafe.Pointer, arg2 DebugComponentCallbackUPP) int32

// NewDebugComponent.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517721-newdebugcomponent
func NewDebugComponent(arg0 uint32, arg1 unsafe.Pointer, arg2 DebugComponentCallbackUPP) int32 {
	if _newDebugComponent == nil {
		panic("coreservices: symbol NewDebugComponent not loaded")
	}
	return _newDebugComponent(arg0, arg1, arg2)
}

var _newDebugComponentCallbackUPP func(arg0 unsafe.Pointer) DebugComponentCallbackUPP

// NewDebugComponentCallbackUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517764-newdebugcomponentcallbackupp
func NewDebugComponentCallbackUPP(arg0 unsafe.Pointer) DebugComponentCallbackUPP {
	if _newDebugComponentCallbackUPP == nil {
		panic("coreservices: symbol NewDebugComponentCallbackUPP not loaded")
	}
	return _newDebugComponentCallbackUPP(arg0)
}

var _newDebugOption func(arg0 uint32, arg1 int32, arg2 unsafe.Pointer) int32

// NewDebugOption.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517718-newdebugoption
func NewDebugOption(arg0 uint32, arg1 int32, arg2 unsafe.Pointer) int32 {
	if _newDebugOption == nil {
		panic("coreservices: symbol NewDebugOption not loaded")
	}
	return _newDebugOption(arg0, arg1, arg2)
}

var _newDebuggerDisposeThreadUPP func(arg0 unsafe.Pointer) DebuggerDisposeThreadUPP

// NewDebuggerDisposeThreadUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574285-newdebuggerdisposethreadupp
func NewDebuggerDisposeThreadUPP(arg0 unsafe.Pointer) DebuggerDisposeThreadUPP {
	if _newDebuggerDisposeThreadUPP == nil {
		panic("coreservices: symbol NewDebuggerDisposeThreadUPP not loaded")
	}
	return _newDebuggerDisposeThreadUPP(arg0)
}

var _newDebuggerNewThreadUPP func(arg0 unsafe.Pointer) DebuggerNewThreadUPP

// NewDebuggerNewThreadUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574208-newdebuggernewthreadupp
func NewDebuggerNewThreadUPP(arg0 unsafe.Pointer) DebuggerNewThreadUPP {
	if _newDebuggerNewThreadUPP == nil {
		panic("coreservices: symbol NewDebuggerNewThreadUPP not loaded")
	}
	return _newDebuggerNewThreadUPP(arg0)
}

var _newDebuggerThreadSchedulerUPP func(arg0 unsafe.Pointer) DebuggerThreadSchedulerUPP

// NewDebuggerThreadSchedulerUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574284-newdebuggerthreadschedulerupp
func NewDebuggerThreadSchedulerUPP(arg0 unsafe.Pointer) DebuggerThreadSchedulerUPP {
	if _newDebuggerThreadSchedulerUPP == nil {
		panic("coreservices: symbol NewDebuggerThreadSchedulerUPP not loaded")
	}
	return _newDebuggerThreadSchedulerUPP(arg0)
}

var _newDeferredTaskUPP func(arg0 unsafe.Pointer) DeferredTaskUPP

// NewDeferredTaskUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533385-newdeferredtaskupp
func NewDeferredTaskUPP(arg0 unsafe.Pointer) DeferredTaskUPP {
	if _newDeferredTaskUPP == nil {
		panic("coreservices: symbol NewDeferredTaskUPP not loaded")
	}
	return _newDeferredTaskUPP(arg0)
}

var _newEmptyHandle func() objectivec.IObject

// NewEmptyHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506288-newemptyhandle
func NewEmptyHandle() objectivec.IObject {
	if _newEmptyHandle == nil {
		panic("coreservices: symbol NewEmptyHandle not loaded")
	}
	return _newEmptyHandle()
}

var _newExceptionHandlerUPP func(arg0 unsafe.Pointer) ExceptionHandlerUPP

// NewExceptionHandlerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1500655-newexceptionhandlerupp
func NewExceptionHandlerUPP(arg0 unsafe.Pointer) ExceptionHandlerUPP {
	if _newExceptionHandlerUPP == nil {
		panic("coreservices: symbol NewExceptionHandlerUPP not loaded")
	}
	return _newExceptionHandlerUPP(arg0)
}

var _newFNSubscriptionUPP func(arg0 unsafe.Pointer) FNSubscriptionUPP

// NewFNSubscriptionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566945-newfnsubscriptionupp
func NewFNSubscriptionUPP(arg0 unsafe.Pointer) FNSubscriptionUPP {
	if _newFNSubscriptionUPP == nil {
		panic("coreservices: symbol NewFNSubscriptionUPP not loaded")
	}
	return _newFNSubscriptionUPP(arg0)
}

var _newFSVolumeEjectUPP func(arg0 unsafe.Pointer) FSVolumeEjectUPP

// NewFSVolumeEjectUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566367-newfsvolumeejectupp
func NewFSVolumeEjectUPP(arg0 unsafe.Pointer) FSVolumeEjectUPP {
	if _newFSVolumeEjectUPP == nil {
		panic("coreservices: symbol NewFSVolumeEjectUPP not loaded")
	}
	return _newFSVolumeEjectUPP(arg0)
}

var _newFSVolumeMountUPP func(arg0 unsafe.Pointer) FSVolumeMountUPP

// NewFSVolumeMountUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565615-newfsvolumemountupp
func NewFSVolumeMountUPP(arg0 unsafe.Pointer) FSVolumeMountUPP {
	if _newFSVolumeMountUPP == nil {
		panic("coreservices: symbol NewFSVolumeMountUPP not loaded")
	}
	return _newFSVolumeMountUPP(arg0)
}

var _newFSVolumeUnmountUPP func(arg0 unsafe.Pointer) FSVolumeUnmountUPP

// NewFSVolumeUnmountUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565539-newfsvolumeunmountupp
func NewFSVolumeUnmountUPP(arg0 unsafe.Pointer) FSVolumeUnmountUPP {
	if _newFSVolumeUnmountUPP == nil {
		panic("coreservices: symbol NewFSVolumeUnmountUPP not loaded")
	}
	return _newFSVolumeUnmountUPP(arg0)
}

var _newFolderManagerNotificationUPP func(arg0 unsafe.Pointer) FolderManagerNotificationUPP

// NewFolderManagerNotificationUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389484-newfoldermanagernotificationupp
func NewFolderManagerNotificationUPP(arg0 unsafe.Pointer) FolderManagerNotificationUPP {
	if _newFolderManagerNotificationUPP == nil {
		panic("coreservices: symbol NewFolderManagerNotificationUPP not loaded")
	}
	return _newFolderManagerNotificationUPP(arg0)
}

var _newGestaltValue func(arg0 uint32, arg1 int32) int16

// NewGestaltValue installs a new [Gestalt] selector code and a value that [Gestalt] returns for that selector.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1472055-newgestaltvalue
func NewGestaltValue(arg0 uint32, arg1 int32) int16 {
	if _newGestaltValue == nil {
		panic("coreservices: symbol NewGestaltValue not loaded")
	}
	return _newGestaltValue(arg0, arg1)
}

var _newGetMissingComponentResourceUPP func(arg0 unsafe.Pointer) GetMissingComponentResourceUPP

// NewGetMissingComponentResourceUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516619-newgetmissingcomponentresourceup
func NewGetMissingComponentResourceUPP(arg0 unsafe.Pointer) GetMissingComponentResourceUPP {
	if _newGetMissingComponentResourceUPP == nil {
		panic("coreservices: symbol NewGetMissingComponentResourceUPP not loaded")
	}
	return _newGetMissingComponentResourceUPP(arg0)
}

var _newHandle func(arg0 corefoundation.CGSize) objectivec.IObject

// NewHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506244-newhandle
func NewHandle(arg0 corefoundation.CGSize) objectivec.IObject {
	if _newHandle == nil {
		panic("coreservices: symbol NewHandle not loaded")
	}
	return _newHandle(arg0)
}

var _newHandleClear func(arg0 corefoundation.CGSize) objectivec.IObject

// NewHandleClear.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506471-newhandleclear
func NewHandleClear(arg0 corefoundation.CGSize) objectivec.IObject {
	if _newHandleClear == nil {
		panic("coreservices: symbol NewHandleClear not loaded")
	}
	return _newHandleClear(arg0)
}

var _newIOCompletionUPP func(arg0 unsafe.Pointer) IOCompletionUPP

// NewIOCompletionUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565869-newiocompletionupp
func NewIOCompletionUPP(arg0 unsafe.Pointer) IOCompletionUPP {
	if _newIOCompletionUPP == nil {
		panic("coreservices: symbol NewIOCompletionUPP not loaded")
	}
	return _newIOCompletionUPP(arg0)
}

var _newIndexToUCStringUPP func(userRoutine IndexToUCStringProcPtr) IndexToUCStringUPP

// NewIndexToUCStringUPP.
//
// See: https://developer.apple.com/documentation/coreservices/1390384-newindextoucstringupp
func NewIndexToUCStringUPP(userRoutine IndexToUCStringProcPtr) IndexToUCStringUPP {
	if _newIndexToUCStringUPP == nil {
		panic("coreservices: symbol NewIndexToUCStringUPP not loaded")
	}
	return _newIndexToUCStringUPP(userRoutine)
}

var _newKCCallbackUPP func(arg0 unsafe.Pointer) KCCallbackUPP

// NewKCCallbackUPP.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1562977-newkccallbackupp
func NewKCCallbackUPP(arg0 unsafe.Pointer) KCCallbackUPP {
	if _newKCCallbackUPP == nil {
		panic("coreservices: symbol NewKCCallbackUPP not loaded")
	}
	return _newKCCallbackUPP(arg0)
}

var _newOSLAccessorUPP func(userRoutine OSLAccessorProcPtr) OSLAccessorUPP

// NewOSLAccessorUPP creates a new universal procedure pointer to an object accessor function.
//
// See: https://developer.apple.com/documentation/coreservices/1449584-newoslaccessorupp
func NewOSLAccessorUPP(userRoutine OSLAccessorProcPtr) OSLAccessorUPP {
	if _newOSLAccessorUPP == nil {
		panic("coreservices: symbol NewOSLAccessorUPP not loaded")
	}
	return _newOSLAccessorUPP(userRoutine)
}

var _newOSLAdjustMarksUPP func(userRoutine OSLAdjustMarksProcPtr) OSLAdjustMarksUPP

// NewOSLAdjustMarksUPP creates a new universal procedure pointer to an object callback adjust marks function.
//
// See: https://developer.apple.com/documentation/coreservices/1443347-newosladjustmarksupp
func NewOSLAdjustMarksUPP(userRoutine OSLAdjustMarksProcPtr) OSLAdjustMarksUPP {
	if _newOSLAdjustMarksUPP == nil {
		panic("coreservices: symbol NewOSLAdjustMarksUPP not loaded")
	}
	return _newOSLAdjustMarksUPP(userRoutine)
}

var _newOSLCompareUPP func(userRoutine OSLCompareProcPtr) OSLCompareUPP

// NewOSLCompareUPP creates a new universal procedure pointer to an object callback comparison function.
//
// See: https://developer.apple.com/documentation/coreservices/1444603-newoslcompareupp
func NewOSLCompareUPP(userRoutine OSLCompareProcPtr) OSLCompareUPP {
	if _newOSLCompareUPP == nil {
		panic("coreservices: symbol NewOSLCompareUPP not loaded")
	}
	return _newOSLCompareUPP(userRoutine)
}

var _newOSLCountUPP func(userRoutine OSLCountProcPtr) OSLCountUPP

// NewOSLCountUPP creates a new universal procedure pointer to an object callback count function.
//
// See: https://developer.apple.com/documentation/coreservices/1448156-newoslcountupp
func NewOSLCountUPP(userRoutine OSLCountProcPtr) OSLCountUPP {
	if _newOSLCountUPP == nil {
		panic("coreservices: symbol NewOSLCountUPP not loaded")
	}
	return _newOSLCountUPP(userRoutine)
}

var _newOSLDisposeTokenUPP func(userRoutine OSLDisposeTokenProcPtr) OSLDisposeTokenUPP

// NewOSLDisposeTokenUPP creates a new universal procedure pointer to an object callback dispose token function.
//
// See: https://developer.apple.com/documentation/coreservices/1450027-newosldisposetokenupp
func NewOSLDisposeTokenUPP(userRoutine OSLDisposeTokenProcPtr) OSLDisposeTokenUPP {
	if _newOSLDisposeTokenUPP == nil {
		panic("coreservices: symbol NewOSLDisposeTokenUPP not loaded")
	}
	return _newOSLDisposeTokenUPP(userRoutine)
}

var _newOSLGetErrDescUPP func(userRoutine OSLGetErrDescProcPtr) OSLGetErrDescUPP

// NewOSLGetErrDescUPP creates a new universal procedure pointer to an object callback get error descriptor function.
//
// See: https://developer.apple.com/documentation/coreservices/1447934-newoslgeterrdescupp
func NewOSLGetErrDescUPP(userRoutine OSLGetErrDescProcPtr) OSLGetErrDescUPP {
	if _newOSLGetErrDescUPP == nil {
		panic("coreservices: symbol NewOSLGetErrDescUPP not loaded")
	}
	return _newOSLGetErrDescUPP(userRoutine)
}

var _newOSLGetMarkTokenUPP func(userRoutine OSLGetMarkTokenProcPtr) OSLGetMarkTokenUPP

// NewOSLGetMarkTokenUPP creates a new universal procedure pointer to an object callback get mark function.
//
// See: https://developer.apple.com/documentation/coreservices/1445166-newoslgetmarktokenupp
func NewOSLGetMarkTokenUPP(userRoutine OSLGetMarkTokenProcPtr) OSLGetMarkTokenUPP {
	if _newOSLGetMarkTokenUPP == nil {
		panic("coreservices: symbol NewOSLGetMarkTokenUPP not loaded")
	}
	return _newOSLGetMarkTokenUPP(userRoutine)
}

var _newOSLMarkUPP func(userRoutine OSLMarkProcPtr) OSLMarkUPP

// NewOSLMarkUPP creates a new universal procedure pointer to an object callback mark function.
//
// See: https://developer.apple.com/documentation/coreservices/1446942-newoslmarkupp
func NewOSLMarkUPP(userRoutine OSLMarkProcPtr) OSLMarkUPP {
	if _newOSLMarkUPP == nil {
		panic("coreservices: symbol NewOSLMarkUPP not loaded")
	}
	return _newOSLMarkUPP(userRoutine)
}

var _newPtr func(arg0 corefoundation.CGSize) coreimage.Ptr

// NewPtr.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506463-newptr
func NewPtr(arg0 corefoundation.CGSize) coreimage.Ptr {
	if _newPtr == nil {
		panic("coreservices: symbol NewPtr not loaded")
	}
	return _newPtr(arg0)
}

var _newPtrClear func(arg0 corefoundation.CGSize) coreimage.Ptr

// NewPtrClear.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506425-newptrclear
func NewPtrClear(arg0 corefoundation.CGSize) coreimage.Ptr {
	if _newPtrClear == nil {
		panic("coreservices: symbol NewPtrClear not loaded")
	}
	return _newPtrClear(arg0)
}

var _newResErrUPP func(arg0 unsafe.Pointer) ResErrUPP

// NewResErrUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529285-newreserrupp
func NewResErrUPP(arg0 unsafe.Pointer) ResErrUPP {
	if _newResErrUPP == nil {
		panic("coreservices: symbol NewResErrUPP not loaded")
	}
	return _newResErrUPP(arg0)
}

var _newSelectorFunctionUPP func(arg0 unsafe.Pointer) SelectorFunctionUPP

// NewSelectorFunctionUPP creates a universal procedure pointer (UPP) to a selector callback function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1472304-newselectorfunctionupp
func NewSelectorFunctionUPP(arg0 unsafe.Pointer) SelectorFunctionUPP {
	if _newSelectorFunctionUPP == nil {
		panic("coreservices: symbol NewSelectorFunctionUPP not loaded")
	}
	return _newSelectorFunctionUPP(arg0)
}

var _newSleepQUPP func(arg0 unsafe.Pointer) SleepQUPP

// NewSleepQUPP.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/coreservices/1427411-newsleepqupp
func NewSleepQUPP(arg0 unsafe.Pointer) SleepQUPP {
	if _newSleepQUPP == nil {
		panic("coreservices: symbol NewSleepQUPP not loaded")
	}
	return _newSleepQUPP(arg0)
}

var _newThread func(arg0 ThreadStyle, arg1 ThreadEntryTPP, arg2 corefoundation.CGSize, arg3 ThreadOptions, arg4 ThreadID) int16

// NewThread.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574248-newthread
func NewThread(arg0 ThreadStyle, arg1 ThreadEntryTPP, arg2 corefoundation.CGSize, arg3 ThreadOptions, arg4 ThreadID) int16 {
	if _newThread == nil {
		panic("coreservices: symbol NewThread not loaded")
	}
	return _newThread(arg0, arg1, arg2, arg3, arg4)
}

var _newThreadEntryUPP func(arg0 unsafe.Pointer) ThreadEntryUPP

// NewThreadEntryUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574224-newthreadentryupp
func NewThreadEntryUPP(arg0 unsafe.Pointer) ThreadEntryUPP {
	if _newThreadEntryUPP == nil {
		panic("coreservices: symbol NewThreadEntryUPP not loaded")
	}
	return _newThreadEntryUPP(arg0)
}

var _newThreadSchedulerUPP func(arg0 unsafe.Pointer) ThreadSchedulerUPP

// NewThreadSchedulerUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574293-newthreadschedulerupp
func NewThreadSchedulerUPP(arg0 unsafe.Pointer) ThreadSchedulerUPP {
	if _newThreadSchedulerUPP == nil {
		panic("coreservices: symbol NewThreadSchedulerUPP not loaded")
	}
	return _newThreadSchedulerUPP(arg0)
}

var _newThreadSwitchUPP func(arg0 unsafe.Pointer) ThreadSwitchUPP

// NewThreadSwitchUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574243-newthreadswitchupp
func NewThreadSwitchUPP(arg0 unsafe.Pointer) ThreadSwitchUPP {
	if _newThreadSwitchUPP == nil {
		panic("coreservices: symbol NewThreadSwitchUPP not loaded")
	}
	return _newThreadSwitchUPP(arg0)
}

var _newThreadTerminationUPP func(arg0 unsafe.Pointer) ThreadTerminationUPP

// NewThreadTerminationUPP.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574221-newthreadterminationupp
func NewThreadTerminationUPP(arg0 unsafe.Pointer) ThreadTerminationUPP {
	if _newThreadTerminationUPP == nil {
		panic("coreservices: symbol NewThreadTerminationUPP not loaded")
	}
	return _newThreadTerminationUPP(arg0)
}

var _newTimerUPP func(arg0 unsafe.Pointer) TimerUPP

// NewTimerUPP.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1550787-newtimerupp
func NewTimerUPP(arg0 unsafe.Pointer) TimerUPP {
	if _newTimerUPP == nil {
		panic("coreservices: symbol NewTimerUPP not loaded")
	}
	return _newTimerUPP(arg0)
}

var _newUnicodeToTextFallbackUPP func(arg0 unsafe.Pointer) UnicodeToTextFallbackUPP

// NewUnicodeToTextFallbackUPP creates a new universal procedure pointer (UPP) to a Unicode-to-textfallback callback.
//
// See: https://developer.apple.com/documentation/coreservices/1433556-newunicodetotextfallbackupp
func NewUnicodeToTextFallbackUPP(arg0 unsafe.Pointer) UnicodeToTextFallbackUPP {
	if _newUnicodeToTextFallbackUPP == nil {
		panic("coreservices: symbol NewUnicodeToTextFallbackUPP not loaded")
	}
	return _newUnicodeToTextFallbackUPP(arg0)
}

var _openAComponent func(arg0 Component, arg1 ComponentInstance) int16

// OpenAComponent.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516558-openacomponent
func OpenAComponent(arg0 Component, arg1 ComponentInstance) int16 {
	if _openAComponent == nil {
		panic("coreservices: symbol OpenAComponent not loaded")
	}
	return _openAComponent(arg0, arg1)
}

var _openAComponentResFile func(arg0 Component, arg1 ResFileRefNum) int16

// OpenAComponentResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516643-openacomponentresfile
func OpenAComponentResFile(arg0 Component, arg1 ResFileRefNum) int16 {
	if _openAComponentResFile == nil {
		panic("coreservices: symbol OpenAComponentResFile not loaded")
	}
	return _openAComponentResFile(arg0, arg1)
}

var _openADefaultComponent func(arg0 uint32, arg1 uint32, arg2 ComponentInstance) int16

// OpenADefaultComponent.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516360-openadefaultcomponent
func OpenADefaultComponent(arg0 uint32, arg1 uint32, arg2 ComponentInstance) int16 {
	if _openADefaultComponent == nil {
		panic("coreservices: symbol OpenADefaultComponent not loaded")
	}
	return _openADefaultComponent(arg0, arg1, arg2)
}

var _openComponent func(arg0 Component) ComponentInstance

// OpenComponent opens a connection to the component with the component identifier specified by your application.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516607-opencomponent
func OpenComponent(arg0 Component) ComponentInstance {
	if _openComponent == nil {
		panic("coreservices: symbol OpenComponent not loaded")
	}
	return _openComponent(arg0)
}

var _openComponentResFile func(arg0 Component) ResFileRefNum

// OpenComponentResFile allows your component to gain access to its resource file.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516550-opencomponentresfile
func OpenComponentResFile(arg0 Component) ResFileRefNum {
	if _openComponentResFile == nil {
		panic("coreservices: symbol OpenComponentResFile not loaded")
	}
	return _openComponentResFile(arg0)
}

var _openDefaultComponent func(arg0 uint32, arg1 uint32) ComponentInstance

// OpenDefaultComponent opens a connection to a registered component of the component type and subtype specified by your application.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516523-opendefaultcomponent
func OpenDefaultComponent(arg0 uint32, arg1 uint32) ComponentInstance {
	if _openDefaultComponent == nil {
		panic("coreservices: symbol OpenDefaultComponent not loaded")
	}
	return _openDefaultComponent(arg0, arg1)
}

var _overrideIconRef func(oldIconRef uintptr, newIconRef uintptr) int16

// OverrideIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1445253-overrideiconref
func OverrideIconRef(oldIconRef uintptr, newIconRef uintptr) int16 {
	if _overrideIconRef == nil {
		panic("coreservices: symbol OverrideIconRef not loaded")
	}
	return _overrideIconRef(oldIconRef, newIconRef)
}

var _pBAllocateForkAsync func(arg0 FSForkIOParam)

// PBAllocateForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566083-pballocateforkasync
func PBAllocateForkAsync(arg0 FSForkIOParam) {
	if _pBAllocateForkAsync == nil {
		panic("coreservices: symbol PBAllocateForkAsync not loaded")
	}
	_pBAllocateForkAsync(arg0)
}

var _pBAllocateForkSync func(arg0 FSForkIOParam) int16

// PBAllocateForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565523-pballocateforksync
func PBAllocateForkSync(arg0 FSForkIOParam) int16 {
	if _pBAllocateForkSync == nil {
		panic("coreservices: symbol PBAllocateForkSync not loaded")
	}
	return _pBAllocateForkSync(arg0)
}

var _pBCatalogSearchAsync func(arg0 FSCatalogBulkParam)

// PBCatalogSearchAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566401-pbcatalogsearchasync
func PBCatalogSearchAsync(arg0 FSCatalogBulkParam) {
	if _pBCatalogSearchAsync == nil {
		panic("coreservices: symbol PBCatalogSearchAsync not loaded")
	}
	_pBCatalogSearchAsync(arg0)
}

var _pBCatalogSearchSync func(arg0 FSCatalogBulkParam) int16

// PBCatalogSearchSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565713-pbcatalogsearchsync
func PBCatalogSearchSync(arg0 FSCatalogBulkParam) int16 {
	if _pBCatalogSearchSync == nil {
		panic("coreservices: symbol PBCatalogSearchSync not loaded")
	}
	return _pBCatalogSearchSync(arg0)
}

var _pBCloseForkAsync func(arg0 FSForkIOParam)

// PBCloseForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565189-pbcloseforkasync
func PBCloseForkAsync(arg0 FSForkIOParam) {
	if _pBCloseForkAsync == nil {
		panic("coreservices: symbol PBCloseForkAsync not loaded")
	}
	_pBCloseForkAsync(arg0)
}

var _pBCloseForkSync func(arg0 FSForkIOParam) int16

// PBCloseForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566173-pbcloseforksync
func PBCloseForkSync(arg0 FSForkIOParam) int16 {
	if _pBCloseForkSync == nil {
		panic("coreservices: symbol PBCloseForkSync not loaded")
	}
	return _pBCloseForkSync(arg0)
}

var _pBCloseIteratorAsync func(arg0 FSCatalogBulkParam)

// PBCloseIteratorAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566055-pbcloseiteratorasync
func PBCloseIteratorAsync(arg0 FSCatalogBulkParam) {
	if _pBCloseIteratorAsync == nil {
		panic("coreservices: symbol PBCloseIteratorAsync not loaded")
	}
	_pBCloseIteratorAsync(arg0)
}

var _pBCloseIteratorSync func(arg0 FSCatalogBulkParam) int16

// PBCloseIteratorSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566452-pbcloseiteratorsync
func PBCloseIteratorSync(arg0 FSCatalogBulkParam) int16 {
	if _pBCloseIteratorSync == nil {
		panic("coreservices: symbol PBCloseIteratorSync not loaded")
	}
	return _pBCloseIteratorSync(arg0)
}

var _pBCompareFSRefsAsync func(arg0 FSRefParam)

// PBCompareFSRefsAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566717-pbcomparefsrefsasync
func PBCompareFSRefsAsync(arg0 FSRefParam) {
	if _pBCompareFSRefsAsync == nil {
		panic("coreservices: symbol PBCompareFSRefsAsync not loaded")
	}
	_pBCompareFSRefsAsync(arg0)
}

var _pBCompareFSRefsSync func(arg0 FSRefParam) int16

// PBCompareFSRefsSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566817-pbcomparefsrefssync
func PBCompareFSRefsSync(arg0 FSRefParam) int16 {
	if _pBCompareFSRefsSync == nil {
		panic("coreservices: symbol PBCompareFSRefsSync not loaded")
	}
	return _pBCompareFSRefsSync(arg0)
}

var _pBCreateDirectoryUnicodeAsync func(arg0 FSRefParam)

// PBCreateDirectoryUnicodeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565107-pbcreatedirectoryunicodeasync
func PBCreateDirectoryUnicodeAsync(arg0 FSRefParam) {
	if _pBCreateDirectoryUnicodeAsync == nil {
		panic("coreservices: symbol PBCreateDirectoryUnicodeAsync not loaded")
	}
	_pBCreateDirectoryUnicodeAsync(arg0)
}

var _pBCreateDirectoryUnicodeSync func(arg0 FSRefParam) int16

// PBCreateDirectoryUnicodeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565315-pbcreatedirectoryunicodesync
func PBCreateDirectoryUnicodeSync(arg0 FSRefParam) int16 {
	if _pBCreateDirectoryUnicodeSync == nil {
		panic("coreservices: symbol PBCreateDirectoryUnicodeSync not loaded")
	}
	return _pBCreateDirectoryUnicodeSync(arg0)
}

var _pBCreateFileAndOpenForkUnicodeAsync func(arg0 FSRefForkIOParamPtr)

// PBCreateFileAndOpenForkUnicodeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566593-pbcreatefileandopenforkunicodeas
func PBCreateFileAndOpenForkUnicodeAsync(arg0 FSRefForkIOParamPtr) {
	if _pBCreateFileAndOpenForkUnicodeAsync == nil {
		panic("coreservices: symbol PBCreateFileAndOpenForkUnicodeAsync not loaded")
	}
	_pBCreateFileAndOpenForkUnicodeAsync(arg0)
}

var _pBCreateFileAndOpenForkUnicodeSync func(arg0 FSRefForkIOParamPtr) int32

// PBCreateFileAndOpenForkUnicodeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566046-pbcreatefileandopenforkunicodesy
func PBCreateFileAndOpenForkUnicodeSync(arg0 FSRefForkIOParamPtr) int32 {
	if _pBCreateFileAndOpenForkUnicodeSync == nil {
		panic("coreservices: symbol PBCreateFileAndOpenForkUnicodeSync not loaded")
	}
	return _pBCreateFileAndOpenForkUnicodeSync(arg0)
}

var _pBCreateFileUnicodeAsync func(arg0 FSRefParam)

// PBCreateFileUnicodeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565280-pbcreatefileunicodeasync
func PBCreateFileUnicodeAsync(arg0 FSRefParam) {
	if _pBCreateFileUnicodeAsync == nil {
		panic("coreservices: symbol PBCreateFileUnicodeAsync not loaded")
	}
	_pBCreateFileUnicodeAsync(arg0)
}

var _pBCreateFileUnicodeSync func(arg0 FSRefParam) int16

// PBCreateFileUnicodeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566896-pbcreatefileunicodesync
func PBCreateFileUnicodeSync(arg0 FSRefParam) int16 {
	if _pBCreateFileUnicodeSync == nil {
		panic("coreservices: symbol PBCreateFileUnicodeSync not loaded")
	}
	return _pBCreateFileUnicodeSync(arg0)
}

var _pBCreateForkAsync func(arg0 FSForkIOParam)

// PBCreateForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565676-pbcreateforkasync
func PBCreateForkAsync(arg0 FSForkIOParam) {
	if _pBCreateForkAsync == nil {
		panic("coreservices: symbol PBCreateForkAsync not loaded")
	}
	_pBCreateForkAsync(arg0)
}

var _pBCreateForkSync func(arg0 FSForkIOParam) int16

// PBCreateForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566697-pbcreateforksync
func PBCreateForkSync(arg0 FSForkIOParam) int16 {
	if _pBCreateForkSync == nil {
		panic("coreservices: symbol PBCreateForkSync not loaded")
	}
	return _pBCreateForkSync(arg0)
}

var _pBDeleteForkAsync func(arg0 FSForkIOParam)

// PBDeleteForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566752-pbdeleteforkasync
func PBDeleteForkAsync(arg0 FSForkIOParam) {
	if _pBDeleteForkAsync == nil {
		panic("coreservices: symbol PBDeleteForkAsync not loaded")
	}
	_pBDeleteForkAsync(arg0)
}

var _pBDeleteForkSync func(arg0 FSForkIOParam) int16

// PBDeleteForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566822-pbdeleteforksync
func PBDeleteForkSync(arg0 FSForkIOParam) int16 {
	if _pBDeleteForkSync == nil {
		panic("coreservices: symbol PBDeleteForkSync not loaded")
	}
	return _pBDeleteForkSync(arg0)
}

var _pBDeleteObjectAsync func(arg0 FSRefParam)

// PBDeleteObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566325-pbdeleteobjectasync
func PBDeleteObjectAsync(arg0 FSRefParam) {
	if _pBDeleteObjectAsync == nil {
		panic("coreservices: symbol PBDeleteObjectAsync not loaded")
	}
	_pBDeleteObjectAsync(arg0)
}

var _pBDeleteObjectSync func(arg0 FSRefParam) int16

// PBDeleteObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565966-pbdeleteobjectsync
func PBDeleteObjectSync(arg0 FSRefParam) int16 {
	if _pBDeleteObjectSync == nil {
		panic("coreservices: symbol PBDeleteObjectSync not loaded")
	}
	return _pBDeleteObjectSync(arg0)
}

var _pBExchangeObjectsAsync func(arg0 FSRefParam)

// PBExchangeObjectsAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566021-pbexchangeobjectsasync
func PBExchangeObjectsAsync(arg0 FSRefParam) {
	if _pBExchangeObjectsAsync == nil {
		panic("coreservices: symbol PBExchangeObjectsAsync not loaded")
	}
	_pBExchangeObjectsAsync(arg0)
}

var _pBExchangeObjectsSync func(arg0 FSRefParam) int16

// PBExchangeObjectsSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566627-pbexchangeobjectssync
func PBExchangeObjectsSync(arg0 FSRefParam) int16 {
	if _pBExchangeObjectsSync == nil {
		panic("coreservices: symbol PBExchangeObjectsSync not loaded")
	}
	return _pBExchangeObjectsSync(arg0)
}

var _pBFSCopyFileAsync func(arg0 FSRefParamPtr) int32

// PBFSCopyFileAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565810-pbfscopyfileasync
func PBFSCopyFileAsync(arg0 FSRefParamPtr) int32 {
	if _pBFSCopyFileAsync == nil {
		panic("coreservices: symbol PBFSCopyFileAsync not loaded")
	}
	return _pBFSCopyFileAsync(arg0)
}

var _pBFSCopyFileSync func(arg0 FSRefParamPtr) int32

// PBFSCopyFileSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566297-pbfscopyfilesync
func PBFSCopyFileSync(arg0 FSRefParamPtr) int32 {
	if _pBFSCopyFileSync == nil {
		panic("coreservices: symbol PBFSCopyFileSync not loaded")
	}
	return _pBFSCopyFileSync(arg0)
}

var _pBFSResolveNodeIDAsync func(arg0 FSRefParamPtr) int32

// PBFSResolveNodeIDAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566562-pbfsresolvenodeidasync
func PBFSResolveNodeIDAsync(arg0 FSRefParamPtr) int32 {
	if _pBFSResolveNodeIDAsync == nil {
		panic("coreservices: symbol PBFSResolveNodeIDAsync not loaded")
	}
	return _pBFSResolveNodeIDAsync(arg0)
}

var _pBFSResolveNodeIDSync func(arg0 FSRefParamPtr) int32

// PBFSResolveNodeIDSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565133-pbfsresolvenodeidsync
func PBFSResolveNodeIDSync(arg0 FSRefParamPtr) int32 {
	if _pBFSResolveNodeIDSync == nil {
		panic("coreservices: symbol PBFSResolveNodeIDSync not loaded")
	}
	return _pBFSResolveNodeIDSync(arg0)
}

var _pBFlushForkAsync func(arg0 FSForkIOParam)

// PBFlushForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565669-pbflushforkasync
func PBFlushForkAsync(arg0 FSForkIOParam) {
	if _pBFlushForkAsync == nil {
		panic("coreservices: symbol PBFlushForkAsync not loaded")
	}
	_pBFlushForkAsync(arg0)
}

var _pBFlushForkSync func(arg0 FSForkIOParam) int16

// PBFlushForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565853-pbflushforksync
func PBFlushForkSync(arg0 FSForkIOParam) int16 {
	if _pBFlushForkSync == nil {
		panic("coreservices: symbol PBFlushForkSync not loaded")
	}
	return _pBFlushForkSync(arg0)
}

var _pBFlushVolumeAsync func(arg0 FSRefParamPtr) int32

// PBFlushVolumeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566072-pbflushvolumeasync
func PBFlushVolumeAsync(arg0 FSRefParamPtr) int32 {
	if _pBFlushVolumeAsync == nil {
		panic("coreservices: symbol PBFlushVolumeAsync not loaded")
	}
	return _pBFlushVolumeAsync(arg0)
}

var _pBFlushVolumeSync func(arg0 FSRefParamPtr) int32

// PBFlushVolumeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565550-pbflushvolumesync
func PBFlushVolumeSync(arg0 FSRefParamPtr) int32 {
	if _pBFlushVolumeSync == nil {
		panic("coreservices: symbol PBFlushVolumeSync not loaded")
	}
	return _pBFlushVolumeSync(arg0)
}

var _pBGetCatalogInfoAsync func(arg0 FSRefParam)

// PBGetCatalogInfoAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565934-pbgetcataloginfoasync
func PBGetCatalogInfoAsync(arg0 FSRefParam) {
	if _pBGetCatalogInfoAsync == nil {
		panic("coreservices: symbol PBGetCatalogInfoAsync not loaded")
	}
	_pBGetCatalogInfoAsync(arg0)
}

var _pBGetCatalogInfoBulkAsync func(arg0 FSCatalogBulkParam)

// PBGetCatalogInfoBulkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566399-pbgetcataloginfobulkasync
func PBGetCatalogInfoBulkAsync(arg0 FSCatalogBulkParam) {
	if _pBGetCatalogInfoBulkAsync == nil {
		panic("coreservices: symbol PBGetCatalogInfoBulkAsync not loaded")
	}
	_pBGetCatalogInfoBulkAsync(arg0)
}

var _pBGetCatalogInfoBulkSync func(arg0 FSCatalogBulkParam) int16

// PBGetCatalogInfoBulkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565707-pbgetcataloginfobulksync
func PBGetCatalogInfoBulkSync(arg0 FSCatalogBulkParam) int16 {
	if _pBGetCatalogInfoBulkSync == nil {
		panic("coreservices: symbol PBGetCatalogInfoBulkSync not loaded")
	}
	return _pBGetCatalogInfoBulkSync(arg0)
}

var _pBGetCatalogInfoSync func(arg0 FSRefParam) int16

// PBGetCatalogInfoSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566328-pbgetcataloginfosync
func PBGetCatalogInfoSync(arg0 FSRefParam) int16 {
	if _pBGetCatalogInfoSync == nil {
		panic("coreservices: symbol PBGetCatalogInfoSync not loaded")
	}
	return _pBGetCatalogInfoSync(arg0)
}

var _pBGetForkCBInfoAsync func(arg0 FSForkCBInfoParam)

// PBGetForkCBInfoAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566027-pbgetforkcbinfoasync
func PBGetForkCBInfoAsync(arg0 FSForkCBInfoParam) {
	if _pBGetForkCBInfoAsync == nil {
		panic("coreservices: symbol PBGetForkCBInfoAsync not loaded")
	}
	_pBGetForkCBInfoAsync(arg0)
}

var _pBGetForkCBInfoSync func(arg0 FSForkCBInfoParam) int16

// PBGetForkCBInfoSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566258-pbgetforkcbinfosync
func PBGetForkCBInfoSync(arg0 FSForkCBInfoParam) int16 {
	if _pBGetForkCBInfoSync == nil {
		panic("coreservices: symbol PBGetForkCBInfoSync not loaded")
	}
	return _pBGetForkCBInfoSync(arg0)
}

var _pBGetForkPositionAsync func(arg0 FSForkIOParam)

// PBGetForkPositionAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566418-pbgetforkpositionasync
func PBGetForkPositionAsync(arg0 FSForkIOParam) {
	if _pBGetForkPositionAsync == nil {
		panic("coreservices: symbol PBGetForkPositionAsync not loaded")
	}
	_pBGetForkPositionAsync(arg0)
}

var _pBGetForkPositionSync func(arg0 FSForkIOParam) int16

// PBGetForkPositionSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565594-pbgetforkpositionsync
func PBGetForkPositionSync(arg0 FSForkIOParam) int16 {
	if _pBGetForkPositionSync == nil {
		panic("coreservices: symbol PBGetForkPositionSync not loaded")
	}
	return _pBGetForkPositionSync(arg0)
}

var _pBGetForkSizeAsync func(arg0 FSForkIOParam)

// PBGetForkSizeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565453-pbgetforksizeasync
func PBGetForkSizeAsync(arg0 FSForkIOParam) {
	if _pBGetForkSizeAsync == nil {
		panic("coreservices: symbol PBGetForkSizeAsync not loaded")
	}
	_pBGetForkSizeAsync(arg0)
}

var _pBGetForkSizeSync func(arg0 FSForkIOParam) int16

// PBGetForkSizeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566956-pbgetforksizesync
func PBGetForkSizeSync(arg0 FSForkIOParam) int16 {
	if _pBGetForkSizeSync == nil {
		panic("coreservices: symbol PBGetForkSizeSync not loaded")
	}
	return _pBGetForkSizeSync(arg0)
}

var _pBGetVolumeInfoAsync func(arg0 FSVolumeInfoParam)

// PBGetVolumeInfoAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565244-pbgetvolumeinfoasync
func PBGetVolumeInfoAsync(arg0 FSVolumeInfoParam) {
	if _pBGetVolumeInfoAsync == nil {
		panic("coreservices: symbol PBGetVolumeInfoAsync not loaded")
	}
	_pBGetVolumeInfoAsync(arg0)
}

var _pBGetVolumeInfoSync func(arg0 FSVolumeInfoParam) int16

// PBGetVolumeInfoSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566892-pbgetvolumeinfosync
func PBGetVolumeInfoSync(arg0 FSVolumeInfoParam) int16 {
	if _pBGetVolumeInfoSync == nil {
		panic("coreservices: symbol PBGetVolumeInfoSync not loaded")
	}
	return _pBGetVolumeInfoSync(arg0)
}

var _pBIterateForksAsync func(arg0 FSForkIOParam)

// PBIterateForksAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566251-pbiterateforksasync
func PBIterateForksAsync(arg0 FSForkIOParam) {
	if _pBIterateForksAsync == nil {
		panic("coreservices: symbol PBIterateForksAsync not loaded")
	}
	_pBIterateForksAsync(arg0)
}

var _pBIterateForksSync func(arg0 FSForkIOParam) int16

// PBIterateForksSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566093-pbiterateforkssync
func PBIterateForksSync(arg0 FSForkIOParam) int16 {
	if _pBIterateForksSync == nil {
		panic("coreservices: symbol PBIterateForksSync not loaded")
	}
	return _pBIterateForksSync(arg0)
}

var _pBMakeFSRefUnicodeAsync func(arg0 FSRefParam)

// PBMakeFSRefUnicodeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566268-pbmakefsrefunicodeasync
func PBMakeFSRefUnicodeAsync(arg0 FSRefParam) {
	if _pBMakeFSRefUnicodeAsync == nil {
		panic("coreservices: symbol PBMakeFSRefUnicodeAsync not loaded")
	}
	_pBMakeFSRefUnicodeAsync(arg0)
}

var _pBMakeFSRefUnicodeSync func(arg0 FSRefParam) int16

// PBMakeFSRefUnicodeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566925-pbmakefsrefunicodesync
func PBMakeFSRefUnicodeSync(arg0 FSRefParam) int16 {
	if _pBMakeFSRefUnicodeSync == nil {
		panic("coreservices: symbol PBMakeFSRefUnicodeSync not loaded")
	}
	return _pBMakeFSRefUnicodeSync(arg0)
}

var _pBMoveObjectAsync func(arg0 FSRefParam)

// PBMoveObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566270-pbmoveobjectasync
func PBMoveObjectAsync(arg0 FSRefParam) {
	if _pBMoveObjectAsync == nil {
		panic("coreservices: symbol PBMoveObjectAsync not loaded")
	}
	_pBMoveObjectAsync(arg0)
}

var _pBMoveObjectSync func(arg0 FSRefParam) int16

// PBMoveObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565935-pbmoveobjectsync
func PBMoveObjectSync(arg0 FSRefParam) int16 {
	if _pBMoveObjectSync == nil {
		panic("coreservices: symbol PBMoveObjectSync not loaded")
	}
	return _pBMoveObjectSync(arg0)
}

var _pBOpenForkAsync func(arg0 FSForkIOParam)

// PBOpenForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566059-pbopenforkasync
func PBOpenForkAsync(arg0 FSForkIOParam) {
	if _pBOpenForkAsync == nil {
		panic("coreservices: symbol PBOpenForkAsync not loaded")
	}
	_pBOpenForkAsync(arg0)
}

var _pBOpenForkSync func(arg0 FSForkIOParam) int16

// PBOpenForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566802-pbopenforksync
func PBOpenForkSync(arg0 FSForkIOParam) int16 {
	if _pBOpenForkSync == nil {
		panic("coreservices: symbol PBOpenForkSync not loaded")
	}
	return _pBOpenForkSync(arg0)
}

var _pBOpenIteratorAsync func(arg0 FSCatalogBulkParam)

// PBOpenIteratorAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566686-pbopeniteratorasync
func PBOpenIteratorAsync(arg0 FSCatalogBulkParam) {
	if _pBOpenIteratorAsync == nil {
		panic("coreservices: symbol PBOpenIteratorAsync not loaded")
	}
	_pBOpenIteratorAsync(arg0)
}

var _pBOpenIteratorSync func(arg0 FSCatalogBulkParam) int16

// PBOpenIteratorSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566865-pbopeniteratorsync
func PBOpenIteratorSync(arg0 FSCatalogBulkParam) int16 {
	if _pBOpenIteratorSync == nil {
		panic("coreservices: symbol PBOpenIteratorSync not loaded")
	}
	return _pBOpenIteratorSync(arg0)
}

var _pBReadForkAsync func(arg0 FSForkIOParam)

// PBReadForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566703-pbreadforkasync
func PBReadForkAsync(arg0 FSForkIOParam) {
	if _pBReadForkAsync == nil {
		panic("coreservices: symbol PBReadForkAsync not loaded")
	}
	_pBReadForkAsync(arg0)
}

var _pBReadForkSync func(arg0 FSForkIOParam) int16

// PBReadForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566739-pbreadforksync
func PBReadForkSync(arg0 FSForkIOParam) int16 {
	if _pBReadForkSync == nil {
		panic("coreservices: symbol PBReadForkSync not loaded")
	}
	return _pBReadForkSync(arg0)
}

var _pBRenameUnicodeAsync func(arg0 FSRefParam)

// PBRenameUnicodeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566334-pbrenameunicodeasync
func PBRenameUnicodeAsync(arg0 FSRefParam) {
	if _pBRenameUnicodeAsync == nil {
		panic("coreservices: symbol PBRenameUnicodeAsync not loaded")
	}
	_pBRenameUnicodeAsync(arg0)
}

var _pBRenameUnicodeSync func(arg0 FSRefParam) int16

// PBRenameUnicodeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566879-pbrenameunicodesync
func PBRenameUnicodeSync(arg0 FSRefParam) int16 {
	if _pBRenameUnicodeSync == nil {
		panic("coreservices: symbol PBRenameUnicodeSync not loaded")
	}
	return _pBRenameUnicodeSync(arg0)
}

var _pBSetCatalogInfoAsync func(arg0 FSRefParam)

// PBSetCatalogInfoAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566340-pbsetcataloginfoasync
func PBSetCatalogInfoAsync(arg0 FSRefParam) {
	if _pBSetCatalogInfoAsync == nil {
		panic("coreservices: symbol PBSetCatalogInfoAsync not loaded")
	}
	_pBSetCatalogInfoAsync(arg0)
}

var _pBSetCatalogInfoSync func(arg0 FSRefParam) int16

// PBSetCatalogInfoSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565271-pbsetcataloginfosync
func PBSetCatalogInfoSync(arg0 FSRefParam) int16 {
	if _pBSetCatalogInfoSync == nil {
		panic("coreservices: symbol PBSetCatalogInfoSync not loaded")
	}
	return _pBSetCatalogInfoSync(arg0)
}

var _pBSetForkPositionAsync func(arg0 FSForkIOParam)

// PBSetForkPositionAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565887-pbsetforkpositionasync
func PBSetForkPositionAsync(arg0 FSForkIOParam) {
	if _pBSetForkPositionAsync == nil {
		panic("coreservices: symbol PBSetForkPositionAsync not loaded")
	}
	_pBSetForkPositionAsync(arg0)
}

var _pBSetForkPositionSync func(arg0 FSForkIOParam) int16

// PBSetForkPositionSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566035-pbsetforkpositionsync
func PBSetForkPositionSync(arg0 FSForkIOParam) int16 {
	if _pBSetForkPositionSync == nil {
		panic("coreservices: symbol PBSetForkPositionSync not loaded")
	}
	return _pBSetForkPositionSync(arg0)
}

var _pBSetForkSizeAsync func(arg0 FSForkIOParam)

// PBSetForkSizeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566552-pbsetforksizeasync
func PBSetForkSizeAsync(arg0 FSForkIOParam) {
	if _pBSetForkSizeAsync == nil {
		panic("coreservices: symbol PBSetForkSizeAsync not loaded")
	}
	_pBSetForkSizeAsync(arg0)
}

var _pBSetForkSizeSync func(arg0 FSForkIOParam) int16

// PBSetForkSizeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566944-pbsetforksizesync
func PBSetForkSizeSync(arg0 FSForkIOParam) int16 {
	if _pBSetForkSizeSync == nil {
		panic("coreservices: symbol PBSetForkSizeSync not loaded")
	}
	return _pBSetForkSizeSync(arg0)
}

var _pBSetVolumeInfoAsync func(arg0 FSVolumeInfoParam)

// PBSetVolumeInfoAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566541-pbsetvolumeinfoasync
func PBSetVolumeInfoAsync(arg0 FSVolumeInfoParam) {
	if _pBSetVolumeInfoAsync == nil {
		panic("coreservices: symbol PBSetVolumeInfoAsync not loaded")
	}
	_pBSetVolumeInfoAsync(arg0)
}

var _pBSetVolumeInfoSync func(arg0 FSVolumeInfoParam) int16

// PBSetVolumeInfoSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566427-pbsetvolumeinfosync
func PBSetVolumeInfoSync(arg0 FSVolumeInfoParam) int16 {
	if _pBSetVolumeInfoSync == nil {
		panic("coreservices: symbol PBSetVolumeInfoSync not loaded")
	}
	return _pBSetVolumeInfoSync(arg0)
}

var _pBUnlinkObjectAsync func(arg0 FSRefParam)

// PBUnlinkObjectAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565840-pbunlinkobjectasync
func PBUnlinkObjectAsync(arg0 FSRefParam) {
	if _pBUnlinkObjectAsync == nil {
		panic("coreservices: symbol PBUnlinkObjectAsync not loaded")
	}
	_pBUnlinkObjectAsync(arg0)
}

var _pBUnlinkObjectSync func(arg0 FSRefParam) int16

// PBUnlinkObjectSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566516-pbunlinkobjectsync
func PBUnlinkObjectSync(arg0 FSRefParam) int16 {
	if _pBUnlinkObjectSync == nil {
		panic("coreservices: symbol PBUnlinkObjectSync not loaded")
	}
	return _pBUnlinkObjectSync(arg0)
}

var _pBWriteForkAsync func(arg0 FSForkIOParam)

// PBWriteForkAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566491-pbwriteforkasync
func PBWriteForkAsync(arg0 FSForkIOParam) {
	if _pBWriteForkAsync == nil {
		panic("coreservices: symbol PBWriteForkAsync not loaded")
	}
	_pBWriteForkAsync(arg0)
}

var _pBWriteForkSync func(arg0 FSForkIOParam) int16

// PBWriteForkSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565632-pbwriteforksync
func PBWriteForkSync(arg0 FSForkIOParam) int16 {
	if _pBWriteForkSync == nil {
		panic("coreservices: symbol PBWriteForkSync not loaded")
	}
	return _pBWriteForkSync(arg0)
}

var _pBXLockRangeAsync func(arg0 FSRangeLockParamPtr) int32

// PBXLockRangeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565710-pbxlockrangeasync
func PBXLockRangeAsync(arg0 FSRangeLockParamPtr) int32 {
	if _pBXLockRangeAsync == nil {
		panic("coreservices: symbol PBXLockRangeAsync not loaded")
	}
	return _pBXLockRangeAsync(arg0)
}

var _pBXLockRangeSync func(arg0 FSRangeLockParamPtr) int32

// PBXLockRangeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1566800-pbxlockrangesync
func PBXLockRangeSync(arg0 FSRangeLockParamPtr) int32 {
	if _pBXLockRangeSync == nil {
		panic("coreservices: symbol PBXLockRangeSync not loaded")
	}
	return _pBXLockRangeSync(arg0)
}

var _pBXUnlockRangeAsync func(arg0 FSRangeLockParamPtr) int32

// PBXUnlockRangeAsync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565257-pbxunlockrangeasync
func PBXUnlockRangeAsync(arg0 FSRangeLockParamPtr) int32 {
	if _pBXUnlockRangeAsync == nil {
		panic("coreservices: symbol PBXUnlockRangeAsync not loaded")
	}
	return _pBXUnlockRangeAsync(arg0)
}

var _pBXUnlockRangeSync func(arg0 FSRangeLockParamPtr) int32

// PBXUnlockRangeSync.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1565709-pbxunlockrangesync
func PBXUnlockRangeSync(arg0 FSRangeLockParamPtr) int32 {
	if _pBXUnlockRangeSync == nil {
		panic("coreservices: symbol PBXUnlockRangeSync not loaded")
	}
	return _pBXUnlockRangeSync(arg0)
}

var _pLpos func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16

// PLpos.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585952-plpos
func PLpos(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16 {
	if _pLpos == nil {
		panic("coreservices: symbol PLpos not loaded")
	}
	return _pLpos(arg0, arg1)
}

var _pLstrcat func(arg0 string, arg1 unsafe.Pointer) *byte

// PLstrcat.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585941-plstrcat
func PLstrcat(arg0 string, arg1 unsafe.Pointer) *byte {
	if _pLstrcat == nil {
		panic("coreservices: symbol PLstrcat not loaded")
	}
	return _pLstrcat(arg0, arg1)
}

var _pLstrchr func(arg0 unsafe.Pointer, arg1 int16) coreimage.Ptr

// PLstrchr.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585945-plstrchr
func PLstrchr(arg0 unsafe.Pointer, arg1 int16) coreimage.Ptr {
	if _pLstrchr == nil {
		panic("coreservices: symbol PLstrchr not loaded")
	}
	return _pLstrchr(arg0, arg1)
}

var _pLstrcmp func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16

// PLstrcmp.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585947-plstrcmp
func PLstrcmp(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16 {
	if _pLstrcmp == nil {
		panic("coreservices: symbol PLstrcmp not loaded")
	}
	return _pLstrcmp(arg0, arg1)
}

var _pLstrcpy func(arg0 string, arg1 unsafe.Pointer) *byte

// PLstrcpy.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585940-plstrcpy
func PLstrcpy(arg0 string, arg1 unsafe.Pointer) *byte {
	if _pLstrcpy == nil {
		panic("coreservices: symbol PLstrcpy not loaded")
	}
	return _pLstrcpy(arg0, arg1)
}

var _pLstrlen func(arg0 unsafe.Pointer) int16

// PLstrlen.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585944-plstrlen
func PLstrlen(arg0 unsafe.Pointer) int16 {
	if _pLstrlen == nil {
		panic("coreservices: symbol PLstrlen not loaded")
	}
	return _pLstrlen(arg0)
}

var _pLstrncat func(arg0 string, arg1 unsafe.Pointer, arg2 int16) *byte

// PLstrncat.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585949-plstrncat
func PLstrncat(arg0 string, arg1 unsafe.Pointer, arg2 int16) *byte {
	if _pLstrncat == nil {
		panic("coreservices: symbol PLstrncat not loaded")
	}
	return _pLstrncat(arg0, arg1, arg2)
}

var _pLstrncmp func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 int16) int16

// PLstrncmp.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585951-plstrncmp
func PLstrncmp(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 int16) int16 {
	if _pLstrncmp == nil {
		panic("coreservices: symbol PLstrncmp not loaded")
	}
	return _pLstrncmp(arg0, arg1, arg2)
}

var _pLstrncpy func(arg0 string, arg1 unsafe.Pointer, arg2 int16) *byte

// PLstrncpy.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585946-plstrncpy
func PLstrncpy(arg0 string, arg1 unsafe.Pointer, arg2 int16) *byte {
	if _pLstrncpy == nil {
		panic("coreservices: symbol PLstrncpy not loaded")
	}
	return _pLstrncpy(arg0, arg1, arg2)
}

var _pLstrpbrk func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) coreimage.Ptr

// PLstrpbrk.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585948-plstrpbrk
func PLstrpbrk(arg0 unsafe.Pointer, arg1 unsafe.Pointer) coreimage.Ptr {
	if _pLstrpbrk == nil {
		panic("coreservices: symbol PLstrpbrk not loaded")
	}
	return _pLstrpbrk(arg0, arg1)
}

var _pLstrrchr func(arg0 unsafe.Pointer, arg1 int16) coreimage.Ptr

// PLstrrchr.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585943-plstrrchr
func PLstrrchr(arg0 unsafe.Pointer, arg1 int16) coreimage.Ptr {
	if _pLstrrchr == nil {
		panic("coreservices: symbol PLstrrchr not loaded")
	}
	return _pLstrrchr(arg0, arg1)
}

var _pLstrspn func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16

// PLstrspn.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585950-plstrspn
func PLstrspn(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16 {
	if _pLstrspn == nil {
		panic("coreservices: symbol PLstrspn not loaded")
	}
	return _pLstrspn(arg0, arg1)
}

var _pLstrstr func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) coreimage.Ptr

// PLstrstr.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1585942-plstrstr
func PLstrstr(arg0 unsafe.Pointer, arg1 unsafe.Pointer) coreimage.Ptr {
	if _pLstrstr == nil {
		panic("coreservices: symbol PLstrstr not loaded")
	}
	return _pLstrstr(arg0, arg1)
}

var _primeTime func(arg0 QElemPtr, arg1 int)

// PrimeTime.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550774-primetime
func PrimeTime(arg0 QElemPtr, arg1 int) {
	if _primeTime == nil {
		panic("coreservices: symbol PrimeTime not loaded")
	}
	_primeTime(arg0, arg1)
}

var _primeTimeTask func(arg0 QElemPtr, arg1 int) int16

// PrimeTimeTask.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550783-primetimetask
func PrimeTimeTask(arg0 QElemPtr, arg1 int) int16 {
	if _primeTimeTask == nil {
		panic("coreservices: symbol PrimeTimeTask not loaded")
	}
	return _primeTimeTask(arg0, arg1)
}

var _ptrAndHand func(arg0 uintptr, arg1 int) int16

// PtrAndHand.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506365-ptrandhand
func PtrAndHand(arg0 uintptr, arg1 int) int16 {
	if _ptrAndHand == nil {
		panic("coreservices: symbol PtrAndHand not loaded")
	}
	return _ptrAndHand(arg0, arg1)
}

var _ptrToHand func(arg0 uintptr, arg1 int) int16

// PtrToHand.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506370-ptrtohand
func PtrToHand(arg0 uintptr, arg1 int) int16 {
	if _ptrToHand == nil {
		panic("coreservices: symbol PtrToHand not loaded")
	}
	return _ptrToHand(arg0, arg1)
}

var _ptrToXHand func(arg0 uintptr, arg1 int) int16

// PtrToXHand.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506246-ptrtoxhand
func PtrToXHand(arg0 uintptr, arg1 int) int16 {
	if _ptrToXHand == nil {
		panic("coreservices: symbol PtrToXHand not loaded")
	}
	return _ptrToXHand(arg0, arg1)
}

var _purgeCollection func(arg0 Collection, arg1 int32, arg2 int32)

// PurgeCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551434-purgecollection
func PurgeCollection(arg0 Collection, arg1 int32, arg2 int32) {
	if _purgeCollection == nil {
		panic("coreservices: symbol PurgeCollection not loaded")
	}
	_purgeCollection(arg0, arg1, arg2)
}

var _purgeCollectionTag func(arg0 Collection, arg1 CollectionTag)

// PurgeCollectionTag.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551327-purgecollectiontag
func PurgeCollectionTag(arg0 Collection, arg1 CollectionTag) {
	if _purgeCollectionTag == nil {
		panic("coreservices: symbol PurgeCollectionTag not loaded")
	}
	_purgeCollectionTag(arg0, arg1)
}

var _queryUnicodeMappings func(arg0 uintptr, arg1 ConstUnicodeMappingPtr, arg2 uintptr, arg3 uintptr, arg4 UnicodeMapping) int32

// QueryUnicodeMappings returns a list of the conversion mappings available onthe system that meet specified matching criteria and returns thenumber of mappings found.
//
// See: https://developer.apple.com/documentation/coreservices/1433630-queryunicodemappings
func QueryUnicodeMappings(arg0 uintptr, arg1 ConstUnicodeMappingPtr, arg2 uintptr, arg3 uintptr, arg4 UnicodeMapping) int32 {
	if _queryUnicodeMappings == nil {
		panic("coreservices: symbol QueryUnicodeMappings not loaded")
	}
	return _queryUnicodeMappings(arg0, arg1, arg2, arg3, arg4)
}

var _readIconFromFSRef func(ref unsafe.Pointer, iconFamily *IconFamilyHandle) int32

// ReadIconFromFSRef.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/coreservices/1444939-readiconfromfsref
func ReadIconFromFSRef(ref unsafe.Pointer, iconFamily *IconFamilyHandle) int32 {
	if _readIconFromFSRef == nil {
		panic("coreservices: symbol ReadIconFromFSRef not loaded")
	}
	return _readIconFromFSRef(ref, iconFamily)
}

var _readLocation func(arg0 MachineLocation)

// ReadLocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533377-readlocation
func ReadLocation(arg0 MachineLocation) {
	if _readLocation == nil {
		panic("coreservices: symbol ReadLocation not loaded")
	}
	_readLocation(arg0)
}

var _readPartialResource func(arg0 uintptr, arg1 int, arg2 int)

// ReadPartialResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529248-readpartialresource
func ReadPartialResource(arg0 uintptr, arg1 int, arg2 int) {
	if _readPartialResource == nil {
		panic("coreservices: symbol ReadPartialResource not loaded")
	}
	_readPartialResource(arg0, arg1, arg2)
}

var _reallocateHandle func(arg0 uintptr, arg1 corefoundation.CGSize)

// ReallocateHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506393-reallocatehandle
func ReallocateHandle(arg0 uintptr, arg1 corefoundation.CGSize) {
	if _reallocateHandle == nil {
		panic("coreservices: symbol ReallocateHandle not loaded")
	}
	_reallocateHandle(arg0, arg1)
}

var _recoverHandle func(arg0 coreimage.Ptr) objectivec.IObject

// RecoverHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506310-recoverhandle
func RecoverHandle(arg0 coreimage.Ptr) objectivec.IObject {
	if _recoverHandle == nil {
		panic("coreservices: symbol RecoverHandle not loaded")
	}
	return _recoverHandle(arg0)
}

var _registerComponent func(arg0 ComponentDescription, arg1 ComponentRoutineUPP, arg2 int16, arg3 uintptr, arg4 uintptr, arg5 uintptr) Component

// RegisterComponent registers a component stored in memory.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516537-registercomponent
func RegisterComponent(arg0 ComponentDescription, arg1 ComponentRoutineUPP, arg2 int16, arg3 uintptr, arg4 uintptr, arg5 uintptr) Component {
	if _registerComponent == nil {
		panic("coreservices: symbol RegisterComponent not loaded")
	}
	return _registerComponent(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _registerComponentFileRef func(arg0 uintptr, arg1 int16) int16

// RegisterComponentFileRef.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516564-registercomponentfileref
func RegisterComponentFileRef(arg0 uintptr, arg1 int16) int16 {
	if _registerComponentFileRef == nil {
		panic("coreservices: symbol RegisterComponentFileRef not loaded")
	}
	return _registerComponentFileRef(arg0, arg1)
}

var _registerComponentFileRefEntries func(arg0 uintptr, arg1 int16, arg2 ComponentDescription, arg3 uint32) int16

// RegisterComponentFileRefEntries.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516395-registercomponentfilerefentries
func RegisterComponentFileRefEntries(arg0 uintptr, arg1 int16, arg2 ComponentDescription, arg3 uint32) int16 {
	if _registerComponentFileRefEntries == nil {
		panic("coreservices: symbol RegisterComponentFileRefEntries not loaded")
	}
	return _registerComponentFileRefEntries(arg0, arg1, arg2, arg3)
}

var _registerComponentResource func(arg0 ComponentResourceHandle, arg1 int16) Component

// RegisterComponentResource registers a component stored in a resource file.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516594-registercomponentresource
func RegisterComponentResource(arg0 ComponentResourceHandle, arg1 int16) Component {
	if _registerComponentResource == nil {
		panic("coreservices: symbol RegisterComponentResource not loaded")
	}
	return _registerComponentResource(arg0, arg1)
}

var _registerComponentResourceFile func(arg0 int16, arg1 int16) int32

// RegisterComponentResourceFile registers all component resources in the given resource file.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516511-registercomponentresourcefile
func RegisterComponentResourceFile(arg0 int16, arg1 int16) int32 {
	if _registerComponentResourceFile == nil {
		panic("coreservices: symbol RegisterComponentResourceFile not loaded")
	}
	return _registerComponentResourceFile(arg0, arg1)
}

var _registerIconRefFromFSRef func(creator uint32, iconType uint32, iconFile unsafe.Pointer, theIconRef uintptr) int32

// RegisterIconRefFromFSRef.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/coreservices/1446795-registericonreffromfsref
func RegisterIconRefFromFSRef(creator uint32, iconType uint32, iconFile unsafe.Pointer, theIconRef uintptr) int32 {
	if _registerIconRefFromFSRef == nil {
		panic("coreservices: symbol RegisterIconRefFromFSRef not loaded")
	}
	return _registerIconRefFromFSRef(creator, iconType, iconFile, theIconRef)
}

var _registerIconRefFromIconFamily func(creator uint32, iconType uint32, iconFamily IconFamilyHandle, theIconRef uintptr) int16

// RegisterIconRefFromIconFamily.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1443918-registericonreffromiconfamily
func RegisterIconRefFromIconFamily(creator uint32, iconType uint32, iconFamily IconFamilyHandle, theIconRef uintptr) int16 {
	if _registerIconRefFromIconFamily == nil {
		panic("coreservices: symbol RegisterIconRefFromIconFamily not loaded")
	}
	return _registerIconRefFromIconFamily(creator, iconType, iconFamily, theIconRef)
}

var _releaseCollection func(arg0 Collection) int32

// ReleaseCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551418-releasecollection
func ReleaseCollection(arg0 Collection) int32 {
	if _releaseCollection == nil {
		panic("coreservices: symbol ReleaseCollection not loaded")
	}
	return _releaseCollection(arg0)
}

var _releaseFolder func(arg0 uintptr, arg1 uint32) int16

// ReleaseFolder.
//
// Deprecated: Deprecated since macOS 10.3.
//
// See: https://developer.apple.com/documentation/coreservices/1389109-releasefolder
func ReleaseFolder(arg0 uintptr, arg1 uint32) int16 {
	if _releaseFolder == nil {
		panic("coreservices: symbol ReleaseFolder not loaded")
	}
	return _releaseFolder(arg0, arg1)
}

var _releaseIconRef func(theIconRef uintptr) int16

// ReleaseIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1443504-releaseiconref
func ReleaseIconRef(theIconRef uintptr) int16 {
	if _releaseIconRef == nil {
		panic("coreservices: symbol ReleaseIconRef not loaded")
	}
	return _releaseIconRef(theIconRef)
}

var _releaseResource func(arg0 uintptr)

// ReleaseResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529259-releaseresource
func ReleaseResource(arg0 uintptr) {
	if _releaseResource == nil {
		panic("coreservices: symbol ReleaseResource not loaded")
	}
	_releaseResource(arg0)
}

var _removeCollectionItem func(arg0 Collection, arg1 CollectionTag, arg2 int32) int16

// RemoveCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551326-removecollectionitem
func RemoveCollectionItem(arg0 Collection, arg1 CollectionTag, arg2 int32) int16 {
	if _removeCollectionItem == nil {
		panic("coreservices: symbol RemoveCollectionItem not loaded")
	}
	return _removeCollectionItem(arg0, arg1, arg2)
}

var _removeFolderDescriptor func(arg0 FolderType) int16

// RemoveFolderDescriptor.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1389531-removefolderdescriptor
func RemoveFolderDescriptor(arg0 FolderType) int16 {
	if _removeFolderDescriptor == nil {
		panic("coreservices: symbol RemoveFolderDescriptor not loaded")
	}
	return _removeFolderDescriptor(arg0)
}

var _removeIconRefOverride func(theIconRef uintptr) int16

// RemoveIconRefOverride.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1445832-removeiconrefoverride
func RemoveIconRefOverride(theIconRef uintptr) int16 {
	if _removeIconRefOverride == nil {
		panic("coreservices: symbol RemoveIconRefOverride not loaded")
	}
	return _removeIconRefOverride(theIconRef)
}

var _removeIndexedCollectionItem func(arg0 Collection, arg1 int32) int16

// RemoveIndexedCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551414-removeindexedcollectionitem
func RemoveIndexedCollectionItem(arg0 Collection, arg1 int32) int16 {
	if _removeIndexedCollectionItem == nil {
		panic("coreservices: symbol RemoveIndexedCollectionItem not loaded")
	}
	return _removeIndexedCollectionItem(arg0, arg1)
}

var _removeResource func(arg0 uintptr)

// RemoveResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529272-removeresource
func RemoveResource(arg0 uintptr) {
	if _removeResource == nil {
		panic("coreservices: symbol RemoveResource not loaded")
	}
	_removeResource(arg0)
}

var _removeTimeTask func(arg0 QElemPtr) int16

// RemoveTimeTask.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550772-removetimetask
func RemoveTimeTask(arg0 QElemPtr) int16 {
	if _removeTimeTask == nil {
		panic("coreservices: symbol RemoveTimeTask not loaded")
	}
	return _removeTimeTask(arg0)
}

var _replaceGestaltValue func(arg0 uint32, arg1 int32) int16

// ReplaceGestaltValue replaces the value that the function [Gestalt] returns for a specified selector code with the value provided to the function.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1472000-replacegestaltvalue
func ReplaceGestaltValue(arg0 uint32, arg1 int32) int16 {
	if _replaceGestaltValue == nil {
		panic("coreservices: symbol ReplaceGestaltValue not loaded")
	}
	return _replaceGestaltValue(arg0, arg1)
}

var _replaceIndexedCollectionItem func(arg0 Collection, arg1 int32, arg2 int32) int16

// ReplaceIndexedCollectionItem.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551424-replaceindexedcollectionitem
func ReplaceIndexedCollectionItem(arg0 Collection, arg1 int32, arg2 int32) int16 {
	if _replaceIndexedCollectionItem == nil {
		panic("coreservices: symbol ReplaceIndexedCollectionItem not loaded")
	}
	return _replaceIndexedCollectionItem(arg0, arg1, arg2)
}

var _replaceIndexedCollectionItemHdl func(arg0 Collection, arg1 int32, arg2 uintptr) int16

// ReplaceIndexedCollectionItemHdl.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551345-replaceindexedcollectionitemhdl
func ReplaceIndexedCollectionItemHdl(arg0 Collection, arg1 int32, arg2 uintptr) int16 {
	if _replaceIndexedCollectionItemHdl == nil {
		panic("coreservices: symbol ReplaceIndexedCollectionItemHdl not loaded")
	}
	return _replaceIndexedCollectionItemHdl(arg0, arg1, arg2)
}

var _resError func() int16

// ResError.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529333-reserror
func ResError() int16 {
	if _resError == nil {
		panic("coreservices: symbol ResError not loaded")
	}
	return _resError()
}

var _resetTextToUnicodeInfo func(arg0 TextToUnicodeInfo) int32

// ResetTextToUnicodeInfo reinitializes all state information kept by the contextobjects.
//
// See: https://developer.apple.com/documentation/coreservices/1433526-resettexttounicodeinfo
func ResetTextToUnicodeInfo(arg0 TextToUnicodeInfo) int32 {
	if _resetTextToUnicodeInfo == nil {
		panic("coreservices: symbol ResetTextToUnicodeInfo not loaded")
	}
	return _resetTextToUnicodeInfo(arg0)
}

var _resetUnicodeToTextInfo func(arg0 UnicodeToTextInfo) int32

// ResetUnicodeToTextInfo reinitializes all state information kept by a Unicodeconverter object.
//
// See: https://developer.apple.com/documentation/coreservices/1433647-resetunicodetotextinfo
func ResetUnicodeToTextInfo(arg0 UnicodeToTextInfo) int32 {
	if _resetUnicodeToTextInfo == nil {
		panic("coreservices: symbol ResetUnicodeToTextInfo not loaded")
	}
	return _resetUnicodeToTextInfo(arg0)
}

var _resetUnicodeToTextRunInfo func(arg0 UnicodeToTextRunInfo) int32

// ResetUnicodeToTextRunInfo reinitializes all state information kept by the contextobjects in TextRun conversions.
//
// See: https://developer.apple.com/documentation/coreservices/1433495-resetunicodetotextruninfo
func ResetUnicodeToTextRunInfo(arg0 UnicodeToTextRunInfo) int32 {
	if _resetUnicodeToTextRunInfo == nil {
		panic("coreservices: symbol ResetUnicodeToTextRunInfo not loaded")
	}
	return _resetUnicodeToTextRunInfo(arg0)
}

var _resolveComponentAlias func(arg0 Component) Component

// ResolveComponentAlias.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516484-resolvecomponentalias
func ResolveComponentAlias(arg0 Component) Component {
	if _resolveComponentAlias == nil {
		panic("coreservices: symbol ResolveComponentAlias not loaded")
	}
	return _resolveComponentAlias(arg0)
}

var _resolveDefaultTextEncoding func(arg0 TextEncoding) TextEncoding

// ResolveDefaultTextEncoding returns a text encoding specification in which any meta-valueshave been resolved to real values.
//
// See: https://developer.apple.com/documentation/coreservices/1400111-resolvedefaulttextencoding
func ResolveDefaultTextEncoding(arg0 TextEncoding) TextEncoding {
	if _resolveDefaultTextEncoding == nil {
		panic("coreservices: symbol ResolveDefaultTextEncoding not loaded")
	}
	return _resolveDefaultTextEncoding(arg0)
}

var _retainCollection func(arg0 Collection) int32

// RetainCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551332-retaincollection
func RetainCollection(arg0 Collection) int32 {
	if _retainCollection == nil {
		panic("coreservices: symbol RetainCollection not loaded")
	}
	return _retainCollection(arg0)
}

var _revertTextEncodingToScriptInfo func(arg0 TextEncoding, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 unsafe.Pointer) int32

// RevertTextEncodingToScriptInfo converts the given Mac OS text encoding specificationto the corresponding script code and, if possible, language codeand font name.
//
// See: https://developer.apple.com/documentation/coreservices/1400023-reverttextencodingtoscriptinfo
func RevertTextEncodingToScriptInfo(arg0 TextEncoding, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 unsafe.Pointer) int32 {
	if _revertTextEncodingToScriptInfo == nil {
		panic("coreservices: symbol RevertTextEncodingToScriptInfo not loaded")
	}
	return _revertTextEncodingToScriptInfo(arg0, arg1, arg2, arg3)
}

var _rmvTime func(arg0 QElemPtr)

// RmvTime.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1550775-rmvtime
func RmvTime(arg0 QElemPtr) {
	if _rmvTime == nil {
		panic("coreservices: symbol RmvTime not loaded")
	}
	_rmvTime(arg0)
}

var _s32Set func(arg0 int64) int32

// S32Set.
//
// See: https://developer.apple.com/documentation/coreservices/1536232-s32set
func S32Set(arg0 int64) int32 {
	if _s32Set == nil {
		panic("coreservices: symbol S32Set not loaded")
	}
	return _s32Set(arg0)
}

var _s64Absolute func(arg0 int64) int64

// S64Absolute.
//
// See: https://developer.apple.com/documentation/coreservices/4357548-s64absolute
func S64Absolute(arg0 int64) int64 {
	if _s64Absolute == nil {
		panic("coreservices: symbol S64Absolute not loaded")
	}
	return _s64Absolute(arg0)
}

var _s64Add func(arg0 int64, arg1 int64) int64

// S64Add.
//
// See: https://developer.apple.com/documentation/coreservices/1536277-s64add
func S64Add(arg0 int64, arg1 int64) int64 {
	if _s64Add == nil {
		panic("coreservices: symbol S64Add not loaded")
	}
	return _s64Add(arg0, arg1)
}

var _s64And func(arg0 int64, arg1 int64) bool

// S64And.
//
// See: https://developer.apple.com/documentation/coreservices/1536260-s64and
func S64And(arg0 int64, arg1 int64) bool {
	if _s64And == nil {
		panic("coreservices: symbol S64And not loaded")
	}
	return _s64And(arg0, arg1)
}

var _s64BitwiseAnd func(arg0 int64, arg1 int64) int64

// S64BitwiseAnd.
//
// See: https://developer.apple.com/documentation/coreservices/1536315-s64bitwiseand
func S64BitwiseAnd(arg0 int64, arg1 int64) int64 {
	if _s64BitwiseAnd == nil {
		panic("coreservices: symbol S64BitwiseAnd not loaded")
	}
	return _s64BitwiseAnd(arg0, arg1)
}

var _s64BitwiseEor func(arg0 int64, arg1 int64) int64

// S64BitwiseEor.
//
// See: https://developer.apple.com/documentation/coreservices/1536312-s64bitwiseeor
func S64BitwiseEor(arg0 int64, arg1 int64) int64 {
	if _s64BitwiseEor == nil {
		panic("coreservices: symbol S64BitwiseEor not loaded")
	}
	return _s64BitwiseEor(arg0, arg1)
}

var _s64BitwiseNot func(arg0 int64) int64

// S64BitwiseNot.
//
// See: https://developer.apple.com/documentation/coreservices/1536333-s64bitwisenot
func S64BitwiseNot(arg0 int64) int64 {
	if _s64BitwiseNot == nil {
		panic("coreservices: symbol S64BitwiseNot not loaded")
	}
	return _s64BitwiseNot(arg0)
}

var _s64BitwiseOr func(arg0 int64, arg1 int64) int64

// S64BitwiseOr.
//
// See: https://developer.apple.com/documentation/coreservices/1536285-s64bitwiseor
func S64BitwiseOr(arg0 int64, arg1 int64) int64 {
	if _s64BitwiseOr == nil {
		panic("coreservices: symbol S64BitwiseOr not loaded")
	}
	return _s64BitwiseOr(arg0, arg1)
}

var _s64Compare func(arg0 int64, arg1 int64) int32

// S64Compare.
//
// See: https://developer.apple.com/documentation/coreservices/1536359-s64compare
func S64Compare(arg0 int64, arg1 int64) int32 {
	if _s64Compare == nil {
		panic("coreservices: symbol S64Compare not loaded")
	}
	return _s64Compare(arg0, arg1)
}

var _s64Div func(arg0 int64, arg1 int64) int64

// S64Div.
//
// See: https://developer.apple.com/documentation/coreservices/1536247-s64div
func S64Div(arg0 int64, arg1 int64) int64 {
	if _s64Div == nil {
		panic("coreservices: symbol S64Div not loaded")
	}
	return _s64Div(arg0, arg1)
}

var _s64Divide func(arg0 int64, arg1 int64, arg2 int64) int64

// S64Divide.
//
// See: https://developer.apple.com/documentation/coreservices/1536237-s64divide
func S64Divide(arg0 int64, arg1 int64, arg2 int64) int64 {
	if _s64Divide == nil {
		panic("coreservices: symbol S64Divide not loaded")
	}
	return _s64Divide(arg0, arg1, arg2)
}

var _s64Eor func(arg0 int64, arg1 int64) bool

// S64Eor.
//
// See: https://developer.apple.com/documentation/coreservices/1536245-s64eor
func S64Eor(arg0 int64, arg1 int64) bool {
	if _s64Eor == nil {
		panic("coreservices: symbol S64Eor not loaded")
	}
	return _s64Eor(arg0, arg1)
}

var _s64Max func() int64

// S64Max.
//
// See: https://developer.apple.com/documentation/coreservices/1536355-s64max
func S64Max() int64 {
	if _s64Max == nil {
		panic("coreservices: symbol S64Max not loaded")
	}
	return _s64Max()
}

var _s64Min func() int64

// S64Min.
//
// See: https://developer.apple.com/documentation/coreservices/1536244-s64min
func S64Min() int64 {
	if _s64Min == nil {
		panic("coreservices: symbol S64Min not loaded")
	}
	return _s64Min()
}

var _s64Mod func(arg0 int64, arg1 int64) int64

// S64Mod.
//
// See: https://developer.apple.com/documentation/coreservices/1536242-s64mod
func S64Mod(arg0 int64, arg1 int64) int64 {
	if _s64Mod == nil {
		panic("coreservices: symbol S64Mod not loaded")
	}
	return _s64Mod(arg0, arg1)
}

var _s64Multiply func(arg0 int64, arg1 int64) int64

// S64Multiply.
//
// See: https://developer.apple.com/documentation/coreservices/1536274-s64multiply
func S64Multiply(arg0 int64, arg1 int64) int64 {
	if _s64Multiply == nil {
		panic("coreservices: symbol S64Multiply not loaded")
	}
	return _s64Multiply(arg0, arg1)
}

var _s64Negate func(arg0 int64) int64

// S64Negate.
//
// See: https://developer.apple.com/documentation/coreservices/1536337-s64negate
func S64Negate(arg0 int64) int64 {
	if _s64Negate == nil {
		panic("coreservices: symbol S64Negate not loaded")
	}
	return _s64Negate(arg0)
}

var _s64Not func(arg0 int64) bool

// S64Not.
//
// See: https://developer.apple.com/documentation/coreservices/1536319-s64not
func S64Not(arg0 int64) bool {
	if _s64Not == nil {
		panic("coreservices: symbol S64Not not loaded")
	}
	return _s64Not(arg0)
}

var _s64Or func(arg0 int64, arg1 int64) bool

// S64Or.
//
// See: https://developer.apple.com/documentation/coreservices/1536264-s64or
func S64Or(arg0 int64, arg1 int64) bool {
	if _s64Or == nil {
		panic("coreservices: symbol S64Or not loaded")
	}
	return _s64Or(arg0, arg1)
}

var _s64Set func(arg0 int32) int64

// S64Set.
//
// See: https://developer.apple.com/documentation/coreservices/1536243-s64set
func S64Set(arg0 int32) int64 {
	if _s64Set == nil {
		panic("coreservices: symbol S64Set not loaded")
	}
	return _s64Set(arg0)
}

var _s64SetU func(arg0 uint32) int64

// S64SetU.
//
// See: https://developer.apple.com/documentation/coreservices/1536330-s64setu
func S64SetU(arg0 uint32) int64 {
	if _s64SetU == nil {
		panic("coreservices: symbol S64SetU not loaded")
	}
	return _s64SetU(arg0)
}

var _s64ShiftLeft func(arg0 int64, arg1 uint32) int64

// S64ShiftLeft.
//
// See: https://developer.apple.com/documentation/coreservices/1536299-s64shiftleft
func S64ShiftLeft(arg0 int64, arg1 uint32) int64 {
	if _s64ShiftLeft == nil {
		panic("coreservices: symbol S64ShiftLeft not loaded")
	}
	return _s64ShiftLeft(arg0, arg1)
}

var _s64ShiftRight func(arg0 int64, arg1 uint32) int64

// S64ShiftRight.
//
// See: https://developer.apple.com/documentation/coreservices/1536323-s64shiftright
func S64ShiftRight(arg0 int64, arg1 uint32) int64 {
	if _s64ShiftRight == nil {
		panic("coreservices: symbol S64ShiftRight not loaded")
	}
	return _s64ShiftRight(arg0, arg1)
}

var _s64Subtract func(arg0 int64, arg1 int64) int64

// S64Subtract.
//
// See: https://developer.apple.com/documentation/coreservices/1536289-s64subtract
func S64Subtract(arg0 int64, arg1 int64) int64 {
	if _s64Subtract == nil {
		panic("coreservices: symbol S64Subtract not loaded")
	}
	return _s64Subtract(arg0, arg1)
}

var _sInt64ToLongDouble func(arg0 int64) unsafe.Pointer

// SInt64ToLongDouble.
//
// See: https://developer.apple.com/documentation/coreservices/1536253-sint64tolongdouble
func SInt64ToLongDouble(arg0 int64) unsafe.Pointer {
	if _sInt64ToLongDouble == nil {
		panic("coreservices: symbol SInt64ToLongDouble not loaded")
	}
	return _sInt64ToLongDouble(arg0)
}

var _sInt64ToUInt64 func(arg0 int64) uint64

// SInt64ToUInt64.
//
// See: https://developer.apple.com/documentation/coreservices/1536320-sint64touint64
func SInt64ToUInt64(arg0 int64) uint64 {
	if _sInt64ToUInt64 == nil {
		panic("coreservices: symbol SInt64ToUInt64 not loaded")
	}
	return _sInt64ToUInt64(arg0)
}

var _sInt64ToWide func(arg0 int64) unsafe.Pointer

// SInt64ToWide.
//
// See: https://developer.apple.com/documentation/coreservices/1536229-sint64towide
func SInt64ToWide(arg0 int64) unsafe.Pointer {
	if _sInt64ToWide == nil {
		panic("coreservices: symbol SInt64ToWide not loaded")
	}
	return _sInt64ToWide(arg0)
}

var _sKDocumentCopyURL func(inDocument SKDocumentRef) corefoundation.CFURLRef

// SKDocumentCopyURL builds a CFURL object from a document URL object (of type SKDocument).
//
// See: https://developer.apple.com/documentation/coreservices/1449624-skdocumentcopyurl
func SKDocumentCopyURL(inDocument SKDocumentRef) corefoundation.CFURLRef {
	if _sKDocumentCopyURL == nil {
		panic("coreservices: symbol SKDocumentCopyURL not loaded")
	}
	return _sKDocumentCopyURL(inDocument)
}

var _sKDocumentCreate func(inScheme corefoundation.CFStringRef, inParent SKDocumentRef, inName corefoundation.CFStringRef) SKDocumentRef

// SKDocumentCreate creates a document URL object (of type SKDocument) based on a scheme, parent, and name.
//
// See: https://developer.apple.com/documentation/coreservices/1443212-skdocumentcreate
func SKDocumentCreate(inScheme corefoundation.CFStringRef, inParent SKDocumentRef, inName corefoundation.CFStringRef) SKDocumentRef {
	if _sKDocumentCreate == nil {
		panic("coreservices: symbol SKDocumentCreate not loaded")
	}
	return _sKDocumentCreate(inScheme, inParent, inName)
}

var _sKDocumentCreateWithURL func(inURL corefoundation.CFURLRef) SKDocumentRef

// SKDocumentCreateWithURL creates a document URL object (of type SKDocument) from a CFURL object.
//
// See: https://developer.apple.com/documentation/coreservices/1442564-skdocumentcreatewithurl
func SKDocumentCreateWithURL(inURL corefoundation.CFURLRef) SKDocumentRef {
	if _sKDocumentCreateWithURL == nil {
		panic("coreservices: symbol SKDocumentCreateWithURL not loaded")
	}
	return _sKDocumentCreateWithURL(inURL)
}

var _sKDocumentGetName func(inDocument SKDocumentRef) corefoundation.CFStringRef

// SKDocumentGetName gets the name of a document URL object (of type SKDocument).
//
// See: https://developer.apple.com/documentation/coreservices/1442657-skdocumentgetname
func SKDocumentGetName(inDocument SKDocumentRef) corefoundation.CFStringRef {
	if _sKDocumentGetName == nil {
		panic("coreservices: symbol SKDocumentGetName not loaded")
	}
	return _sKDocumentGetName(inDocument)
}

var _sKDocumentGetParent func(inDocument SKDocumentRef) SKDocumentRef

// SKDocumentGetParent gets the parent of a document URL object (of type SKDocument).
//
// See: https://developer.apple.com/documentation/coreservices/1444449-skdocumentgetparent
func SKDocumentGetParent(inDocument SKDocumentRef) SKDocumentRef {
	if _sKDocumentGetParent == nil {
		panic("coreservices: symbol SKDocumentGetParent not loaded")
	}
	return _sKDocumentGetParent(inDocument)
}

var _sKDocumentGetSchemeName func(inDocument SKDocumentRef) corefoundation.CFStringRef

// SKDocumentGetSchemeName gets the scheme name for a document URL object (of type SKDocument).
//
// See: https://developer.apple.com/documentation/coreservices/1448262-skdocumentgetschemename
func SKDocumentGetSchemeName(inDocument SKDocumentRef) corefoundation.CFStringRef {
	if _sKDocumentGetSchemeName == nil {
		panic("coreservices: symbol SKDocumentGetSchemeName not loaded")
	}
	return _sKDocumentGetSchemeName(inDocument)
}

var _sKDocumentGetTypeID func() uint

// SKDocumentGetTypeID gets the type identifier for Search Kit document URL objects.
//
// See: https://developer.apple.com/documentation/coreservices/1448891-skdocumentgettypeid
func SKDocumentGetTypeID() uint {
	if _sKDocumentGetTypeID == nil {
		panic("coreservices: symbol SKDocumentGetTypeID not loaded")
	}
	return _sKDocumentGetTypeID()
}

var _sKIndexAddDocument func(inIndex SKIndexRef, inDocument SKDocumentRef, inMIMETypeHint corefoundation.CFStringRef, inCanReplace unsafe.Pointer) bool

// SKIndexAddDocument adds location information for a file-based document, and the document’s textual content, to an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444897-skindexadddocument
func SKIndexAddDocument(inIndex SKIndexRef, inDocument SKDocumentRef, inMIMETypeHint corefoundation.CFStringRef, inCanReplace unsafe.Pointer) bool {
	if _sKIndexAddDocument == nil {
		panic("coreservices: symbol SKIndexAddDocument not loaded")
	}
	return _sKIndexAddDocument(inIndex, inDocument, inMIMETypeHint, inCanReplace)
}

var _sKIndexAddDocumentWithText func(inIndex SKIndexRef, inDocument SKDocumentRef, inDocumentText corefoundation.CFStringRef, inCanReplace unsafe.Pointer) bool

// SKIndexAddDocumentWithText adds a document URL (SKDocument) object, and the associated document’s textual content, to an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444518-skindexadddocumentwithtext
func SKIndexAddDocumentWithText(inIndex SKIndexRef, inDocument SKDocumentRef, inDocumentText corefoundation.CFStringRef, inCanReplace unsafe.Pointer) bool {
	if _sKIndexAddDocumentWithText == nil {
		panic("coreservices: symbol SKIndexAddDocumentWithText not loaded")
	}
	return _sKIndexAddDocumentWithText(inIndex, inDocument, inDocumentText, inCanReplace)
}

var _sKIndexClose func(inIndex SKIndexRef)

// SKIndexClose closes an index.
//
// See: https://developer.apple.com/documentation/coreservices/1442401-skindexclose
func SKIndexClose(inIndex SKIndexRef) {
	if _sKIndexClose == nil {
		panic("coreservices: symbol SKIndexClose not loaded")
	}
	_sKIndexClose(inIndex)
}

var _sKIndexCompact func(inIndex SKIndexRef) bool

// SKIndexCompact invokes all pending updates associated with an index, compacts the index if compaction is needed, and commits all changes to backing store.
//
// See: https://developer.apple.com/documentation/coreservices/1443628-skindexcompact
func SKIndexCompact(inIndex SKIndexRef) bool {
	if _sKIndexCompact == nil {
		panic("coreservices: symbol SKIndexCompact not loaded")
	}
	return _sKIndexCompact(inIndex)
}

var _sKIndexCopyDocumentForDocumentID func(inIndex SKIndexRef, inDocumentID SKDocumentID) SKDocumentRef

// SKIndexCopyDocumentForDocumentID obtains a document URL object (of type SKDocument) from an index.
//
// See: https://developer.apple.com/documentation/coreservices/1442760-skindexcopydocumentfordocumentid
func SKIndexCopyDocumentForDocumentID(inIndex SKIndexRef, inDocumentID SKDocumentID) SKDocumentRef {
	if _sKIndexCopyDocumentForDocumentID == nil {
		panic("coreservices: symbol SKIndexCopyDocumentForDocumentID not loaded")
	}
	return _sKIndexCopyDocumentForDocumentID(inIndex, inDocumentID)
}

var _sKIndexCopyDocumentIDArrayForTermID func(inIndex SKIndexRef, inTermID int) corefoundation.CFArrayRef

// SKIndexCopyDocumentIDArrayForTermID obtains document IDs for documents that contain a given term.
//
// See: https://developer.apple.com/documentation/coreservices/1448003-skindexcopydocumentidarrayforter
func SKIndexCopyDocumentIDArrayForTermID(inIndex SKIndexRef, inTermID int) corefoundation.CFArrayRef {
	if _sKIndexCopyDocumentIDArrayForTermID == nil {
		panic("coreservices: symbol SKIndexCopyDocumentIDArrayForTermID not loaded")
	}
	return _sKIndexCopyDocumentIDArrayForTermID(inIndex, inTermID)
}

var _sKIndexCopyDocumentProperties func(inIndex SKIndexRef, inDocument SKDocumentRef) corefoundation.CFDictionaryRef

// SKIndexCopyDocumentProperties obtains the application-defined properties of an indexed document.
//
// See: https://developer.apple.com/documentation/coreservices/1449500-skindexcopydocumentproperties
func SKIndexCopyDocumentProperties(inIndex SKIndexRef, inDocument SKDocumentRef) corefoundation.CFDictionaryRef {
	if _sKIndexCopyDocumentProperties == nil {
		panic("coreservices: symbol SKIndexCopyDocumentProperties not loaded")
	}
	return _sKIndexCopyDocumentProperties(inIndex, inDocument)
}

var _sKIndexCopyDocumentRefsForDocumentIDs func(inIndex SKIndexRef, inCount int, inDocumentIDsArray *SKDocumentID, outDocumentRefsArray *SKDocumentRef)

// SKIndexCopyDocumentRefsForDocumentIDs gets document URL objects (of type SKDocument) based on document IDs.
//
// See: https://developer.apple.com/documentation/coreservices/1445305-skindexcopydocumentrefsfordocume
func SKIndexCopyDocumentRefsForDocumentIDs(inIndex SKIndexRef, inCount int, inDocumentIDsArray *SKDocumentID, outDocumentRefsArray *SKDocumentRef) {
	if _sKIndexCopyDocumentRefsForDocumentIDs == nil {
		panic("coreservices: symbol SKIndexCopyDocumentRefsForDocumentIDs not loaded")
	}
	_sKIndexCopyDocumentRefsForDocumentIDs(inIndex, inCount, inDocumentIDsArray, outDocumentRefsArray)
}

var _sKIndexCopyDocumentURLsForDocumentIDs func(inIndex SKIndexRef, inCount int, inDocumentIDsArray *SKDocumentID, outDocumentURLsArray *corefoundation.CFURLRef)

// SKIndexCopyDocumentURLsForDocumentIDs gets document URLs based on document IDs.
//
// See: https://developer.apple.com/documentation/coreservices/1443501-skindexcopydocumenturlsfordocume
func SKIndexCopyDocumentURLsForDocumentIDs(inIndex SKIndexRef, inCount int, inDocumentIDsArray *SKDocumentID, outDocumentURLsArray *corefoundation.CFURLRef) {
	if _sKIndexCopyDocumentURLsForDocumentIDs == nil {
		panic("coreservices: symbol SKIndexCopyDocumentURLsForDocumentIDs not loaded")
	}
	_sKIndexCopyDocumentURLsForDocumentIDs(inIndex, inCount, inDocumentIDsArray, outDocumentURLsArray)
}

var _sKIndexCopyInfoForDocumentIDs func(inIndex SKIndexRef, inCount int, inDocumentIDsArray *SKDocumentID, outNamesArray *corefoundation.CFStringRef, outParentIDsArray *SKDocumentID)

// SKIndexCopyInfoForDocumentIDs gets document names and parent IDs based on document IDs.
//
// See: https://developer.apple.com/documentation/coreservices/1445499-skindexcopyinfofordocumentids
func SKIndexCopyInfoForDocumentIDs(inIndex SKIndexRef, inCount int, inDocumentIDsArray *SKDocumentID, outNamesArray *corefoundation.CFStringRef, outParentIDsArray *SKDocumentID) {
	if _sKIndexCopyInfoForDocumentIDs == nil {
		panic("coreservices: symbol SKIndexCopyInfoForDocumentIDs not loaded")
	}
	_sKIndexCopyInfoForDocumentIDs(inIndex, inCount, inDocumentIDsArray, outNamesArray, outParentIDsArray)
}

var _sKIndexCopyTermIDArrayForDocumentID func(inIndex SKIndexRef, inDocumentID SKDocumentID) corefoundation.CFArrayRef

// SKIndexCopyTermIDArrayForDocumentID obtains the IDs for the terms of an indexed document.
//
// See: https://developer.apple.com/documentation/coreservices/1446868-skindexcopytermidarrayfordocumen
func SKIndexCopyTermIDArrayForDocumentID(inIndex SKIndexRef, inDocumentID SKDocumentID) corefoundation.CFArrayRef {
	if _sKIndexCopyTermIDArrayForDocumentID == nil {
		panic("coreservices: symbol SKIndexCopyTermIDArrayForDocumentID not loaded")
	}
	return _sKIndexCopyTermIDArrayForDocumentID(inIndex, inDocumentID)
}

var _sKIndexCopyTermStringForTermID func(inIndex SKIndexRef, inTermID int) corefoundation.CFStringRef

// SKIndexCopyTermStringForTermID obtains a term, specified by ID, from an index.
//
// See: https://developer.apple.com/documentation/coreservices/1442802-skindexcopytermstringfortermid
func SKIndexCopyTermStringForTermID(inIndex SKIndexRef, inTermID int) corefoundation.CFStringRef {
	if _sKIndexCopyTermStringForTermID == nil {
		panic("coreservices: symbol SKIndexCopyTermStringForTermID not loaded")
	}
	return _sKIndexCopyTermStringForTermID(inIndex, inTermID)
}

var _sKIndexCreateWithMutableData func(inData corefoundation.CFMutableDataRef, inIndexName corefoundation.CFStringRef, inIndexType SKIndexType, inAnalysisProperties corefoundation.CFDictionaryRef) SKIndexRef

// SKIndexCreateWithMutableData creates a named index stored in a [CFMutableDataRef] object.
//
// See: https://developer.apple.com/documentation/coreservices/1447500-skindexcreatewithmutabledata
func SKIndexCreateWithMutableData(inData corefoundation.CFMutableDataRef, inIndexName corefoundation.CFStringRef, inIndexType SKIndexType, inAnalysisProperties corefoundation.CFDictionaryRef) SKIndexRef {
	if _sKIndexCreateWithMutableData == nil {
		panic("coreservices: symbol SKIndexCreateWithMutableData not loaded")
	}
	return _sKIndexCreateWithMutableData(inData, inIndexName, inIndexType, inAnalysisProperties)
}

var _sKIndexCreateWithURL func(inURL corefoundation.CFURLRef, inIndexName corefoundation.CFStringRef, inIndexType SKIndexType, inAnalysisProperties corefoundation.CFDictionaryRef) SKIndexRef

// SKIndexCreateWithURL creates a named index in a file whose location is specified with a CFURL object.
//
// See: https://developer.apple.com/documentation/coreservices/1446111-skindexcreatewithurl
func SKIndexCreateWithURL(inURL corefoundation.CFURLRef, inIndexName corefoundation.CFStringRef, inIndexType SKIndexType, inAnalysisProperties corefoundation.CFDictionaryRef) SKIndexRef {
	if _sKIndexCreateWithURL == nil {
		panic("coreservices: symbol SKIndexCreateWithURL not loaded")
	}
	return _sKIndexCreateWithURL(inURL, inIndexName, inIndexType, inAnalysisProperties)
}

var _sKIndexDocumentIteratorCopyNext func(inIterator SKIndexDocumentIteratorRef) SKDocumentRef

// SKIndexDocumentIteratorCopyNext obtains the next document URL object (of type SKDocument) from an index using a document iterator.
//
// See: https://developer.apple.com/documentation/coreservices/1442212-skindexdocumentiteratorcopynext
func SKIndexDocumentIteratorCopyNext(inIterator SKIndexDocumentIteratorRef) SKDocumentRef {
	if _sKIndexDocumentIteratorCopyNext == nil {
		panic("coreservices: symbol SKIndexDocumentIteratorCopyNext not loaded")
	}
	return _sKIndexDocumentIteratorCopyNext(inIterator)
}

var _sKIndexDocumentIteratorCreate func(inIndex SKIndexRef, inParentDocument SKDocumentRef) SKIndexDocumentIteratorRef

// SKIndexDocumentIteratorCreate creates an index-based iterator for document URL objects (of type SKDocument) owned by a parent document URL object.
//
// See: https://developer.apple.com/documentation/coreservices/1446189-skindexdocumentiteratorcreate
func SKIndexDocumentIteratorCreate(inIndex SKIndexRef, inParentDocument SKDocumentRef) SKIndexDocumentIteratorRef {
	if _sKIndexDocumentIteratorCreate == nil {
		panic("coreservices: symbol SKIndexDocumentIteratorCreate not loaded")
	}
	return _sKIndexDocumentIteratorCreate(inIndex, inParentDocument)
}

var _sKIndexDocumentIteratorGetTypeID func() uint

// SKIndexDocumentIteratorGetTypeID gets the type identifier for Search Kit document iterators.
//
// See: https://developer.apple.com/documentation/coreservices/1443022-skindexdocumentiteratorgettypeid
func SKIndexDocumentIteratorGetTypeID() uint {
	if _sKIndexDocumentIteratorGetTypeID == nil {
		panic("coreservices: symbol SKIndexDocumentIteratorGetTypeID not loaded")
	}
	return _sKIndexDocumentIteratorGetTypeID()
}

var _sKIndexFlush func(inIndex SKIndexRef) bool

// SKIndexFlush invokes all pending updates associated with an index and commits them to backing store.
//
// See: https://developer.apple.com/documentation/coreservices/1450667-skindexflush
func SKIndexFlush(inIndex SKIndexRef) bool {
	if _sKIndexFlush == nil {
		panic("coreservices: symbol SKIndexFlush not loaded")
	}
	return _sKIndexFlush(inIndex)
}

var _sKIndexGetAnalysisProperties func(inIndex SKIndexRef) corefoundation.CFDictionaryRef

// SKIndexGetAnalysisProperties gets the text analysis properties of an index.
//
// See: https://developer.apple.com/documentation/coreservices/1443461-skindexgetanalysisproperties
func SKIndexGetAnalysisProperties(inIndex SKIndexRef) corefoundation.CFDictionaryRef {
	if _sKIndexGetAnalysisProperties == nil {
		panic("coreservices: symbol SKIndexGetAnalysisProperties not loaded")
	}
	return _sKIndexGetAnalysisProperties(inIndex)
}

var _sKIndexGetDocumentCount func(inIndex SKIndexRef) int

// SKIndexGetDocumentCount gets the total number of documents represented in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1449093-skindexgetdocumentcount
func SKIndexGetDocumentCount(inIndex SKIndexRef) int {
	if _sKIndexGetDocumentCount == nil {
		panic("coreservices: symbol SKIndexGetDocumentCount not loaded")
	}
	return _sKIndexGetDocumentCount(inIndex)
}

var _sKIndexGetDocumentID func(inIndex SKIndexRef, inDocument SKDocumentRef) SKDocumentID

// SKIndexGetDocumentID gets the ID of a document URL object (of type SKDocument) in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444437-skindexgetdocumentid
func SKIndexGetDocumentID(inIndex SKIndexRef, inDocument SKDocumentRef) SKDocumentID {
	if _sKIndexGetDocumentID == nil {
		panic("coreservices: symbol SKIndexGetDocumentID not loaded")
	}
	return _sKIndexGetDocumentID(inIndex, inDocument)
}

var _sKIndexGetDocumentState func(inIndex SKIndexRef, inDocument SKDocumentRef) SKDocumentIndexState

// SKIndexGetDocumentState gets the current indexing state of a document URL object (of type SKDocument) in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1443396-skindexgetdocumentstate
func SKIndexGetDocumentState(inIndex SKIndexRef, inDocument SKDocumentRef) SKDocumentIndexState {
	if _sKIndexGetDocumentState == nil {
		panic("coreservices: symbol SKIndexGetDocumentState not loaded")
	}
	return _sKIndexGetDocumentState(inIndex, inDocument)
}

var _sKIndexGetDocumentTermCount func(inIndex SKIndexRef, inDocumentID SKDocumentID) int

// SKIndexGetDocumentTermCount gets the number of terms for a document in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1448341-skindexgetdocumenttermcount
func SKIndexGetDocumentTermCount(inIndex SKIndexRef, inDocumentID SKDocumentID) int {
	if _sKIndexGetDocumentTermCount == nil {
		panic("coreservices: symbol SKIndexGetDocumentTermCount not loaded")
	}
	return _sKIndexGetDocumentTermCount(inIndex, inDocumentID)
}

var _sKIndexGetDocumentTermFrequency func(inIndex SKIndexRef, inDocumentID SKDocumentID, inTermID int) int

// SKIndexGetDocumentTermFrequency gets the number of occurrences of a term in a document.
//
// See: https://developer.apple.com/documentation/coreservices/1447537-skindexgetdocumenttermfrequency
func SKIndexGetDocumentTermFrequency(inIndex SKIndexRef, inDocumentID SKDocumentID, inTermID int) int {
	if _sKIndexGetDocumentTermFrequency == nil {
		panic("coreservices: symbol SKIndexGetDocumentTermFrequency not loaded")
	}
	return _sKIndexGetDocumentTermFrequency(inIndex, inDocumentID, inTermID)
}

var _sKIndexGetIndexType func(inIndex SKIndexRef) SKIndexType

// SKIndexGetIndexType gets the category of an index.
//
// See: https://developer.apple.com/documentation/coreservices/1442236-skindexgetindextype
func SKIndexGetIndexType(inIndex SKIndexRef) SKIndexType {
	if _sKIndexGetIndexType == nil {
		panic("coreservices: symbol SKIndexGetIndexType not loaded")
	}
	return _sKIndexGetIndexType(inIndex)
}

var _sKIndexGetMaximumBytesBeforeFlush func(inIndex SKIndexRef) int

// SKIndexGetMaximumBytesBeforeFlush not recommended.
//
// See: https://developer.apple.com/documentation/coreservices/1445329-skindexgetmaximumbytesbeforeflus
func SKIndexGetMaximumBytesBeforeFlush(inIndex SKIndexRef) int {
	if _sKIndexGetMaximumBytesBeforeFlush == nil {
		panic("coreservices: symbol SKIndexGetMaximumBytesBeforeFlush not loaded")
	}
	return _sKIndexGetMaximumBytesBeforeFlush(inIndex)
}

var _sKIndexGetMaximumDocumentID func(inIndex SKIndexRef) SKDocumentID

// SKIndexGetMaximumDocumentID gets the highest-numbered document ID in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444628-skindexgetmaximumdocumentid
func SKIndexGetMaximumDocumentID(inIndex SKIndexRef) SKDocumentID {
	if _sKIndexGetMaximumDocumentID == nil {
		panic("coreservices: symbol SKIndexGetMaximumDocumentID not loaded")
	}
	return _sKIndexGetMaximumDocumentID(inIndex)
}

var _sKIndexGetMaximumTermID func(inIndex SKIndexRef) int

// SKIndexGetMaximumTermID gets the highest-numbered term ID in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444278-skindexgetmaximumtermid
func SKIndexGetMaximumTermID(inIndex SKIndexRef) int {
	if _sKIndexGetMaximumTermID == nil {
		panic("coreservices: symbol SKIndexGetMaximumTermID not loaded")
	}
	return _sKIndexGetMaximumTermID(inIndex)
}

var _sKIndexGetTermDocumentCount func(inIndex SKIndexRef, inTermID int) int

// SKIndexGetTermDocumentCount gets the number of documents containing a given term represented in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444015-skindexgettermdocumentcount
func SKIndexGetTermDocumentCount(inIndex SKIndexRef, inTermID int) int {
	if _sKIndexGetTermDocumentCount == nil {
		panic("coreservices: symbol SKIndexGetTermDocumentCount not loaded")
	}
	return _sKIndexGetTermDocumentCount(inIndex, inTermID)
}

var _sKIndexGetTermIDForTermString func(inIndex SKIndexRef, inTermString corefoundation.CFStringRef) int

// SKIndexGetTermIDForTermString gets the ID for a term in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1448558-skindexgettermidfortermstring
func SKIndexGetTermIDForTermString(inIndex SKIndexRef, inTermString corefoundation.CFStringRef) int {
	if _sKIndexGetTermIDForTermString == nil {
		panic("coreservices: symbol SKIndexGetTermIDForTermString not loaded")
	}
	return _sKIndexGetTermIDForTermString(inIndex, inTermString)
}

var _sKIndexGetTypeID func() uint

// SKIndexGetTypeID gets the type identifier for Search Kit indexes.
//
// See: https://developer.apple.com/documentation/coreservices/1450223-skindexgettypeid
func SKIndexGetTypeID() uint {
	if _sKIndexGetTypeID == nil {
		panic("coreservices: symbol SKIndexGetTypeID not loaded")
	}
	return _sKIndexGetTypeID()
}

var _sKIndexMoveDocument func(inIndex SKIndexRef, inDocument SKDocumentRef, inNewParent SKDocumentRef) bool

// SKIndexMoveDocument changes the parent of a document URL object (of type SKDocument) in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1449899-skindexmovedocument
func SKIndexMoveDocument(inIndex SKIndexRef, inDocument SKDocumentRef, inNewParent SKDocumentRef) bool {
	if _sKIndexMoveDocument == nil {
		panic("coreservices: symbol SKIndexMoveDocument not loaded")
	}
	return _sKIndexMoveDocument(inIndex, inDocument, inNewParent)
}

var _sKIndexOpenWithData func(inData corefoundation.CFDataRef, inIndexName corefoundation.CFStringRef) SKIndexRef

// SKIndexOpenWithData opens an existing, named index for searching only.
//
// See: https://developer.apple.com/documentation/coreservices/1446398-skindexopenwithdata
func SKIndexOpenWithData(inData corefoundation.CFDataRef, inIndexName corefoundation.CFStringRef) SKIndexRef {
	if _sKIndexOpenWithData == nil {
		panic("coreservices: symbol SKIndexOpenWithData not loaded")
	}
	return _sKIndexOpenWithData(inData, inIndexName)
}

var _sKIndexOpenWithMutableData func(inData corefoundation.CFMutableDataRef, inIndexName corefoundation.CFStringRef) SKIndexRef

// SKIndexOpenWithMutableData opens an existing, named index for searching and updating.
//
// See: https://developer.apple.com/documentation/coreservices/1444201-skindexopenwithmutabledata
func SKIndexOpenWithMutableData(inData corefoundation.CFMutableDataRef, inIndexName corefoundation.CFStringRef) SKIndexRef {
	if _sKIndexOpenWithMutableData == nil {
		panic("coreservices: symbol SKIndexOpenWithMutableData not loaded")
	}
	return _sKIndexOpenWithMutableData(inData, inIndexName)
}

var _sKIndexOpenWithURL func(inURL corefoundation.CFURLRef, inIndexName corefoundation.CFStringRef, inWriteAccess unsafe.Pointer) SKIndexRef

// SKIndexOpenWithURL opens an existing, named index stored in a file whose location is specified with a CFURL object.
//
// See: https://developer.apple.com/documentation/coreservices/1449017-skindexopenwithurl
func SKIndexOpenWithURL(inURL corefoundation.CFURLRef, inIndexName corefoundation.CFStringRef, inWriteAccess unsafe.Pointer) SKIndexRef {
	if _sKIndexOpenWithURL == nil {
		panic("coreservices: symbol SKIndexOpenWithURL not loaded")
	}
	return _sKIndexOpenWithURL(inURL, inIndexName, inWriteAccess)
}

var _sKIndexRemoveDocument func(inIndex SKIndexRef, inDocument SKDocumentRef) bool

// SKIndexRemoveDocument removes a document URL object (of type SKDocument) and its children, if any, from an index.
//
// See: https://developer.apple.com/documentation/coreservices/1444375-skindexremovedocument
func SKIndexRemoveDocument(inIndex SKIndexRef, inDocument SKDocumentRef) bool {
	if _sKIndexRemoveDocument == nil {
		panic("coreservices: symbol SKIndexRemoveDocument not loaded")
	}
	return _sKIndexRemoveDocument(inIndex, inDocument)
}

var _sKIndexRenameDocument func(inIndex SKIndexRef, inDocument SKDocumentRef, inNewName corefoundation.CFStringRef) bool

// SKIndexRenameDocument changes the name of a document URL object (of type SKDocument) in an index.
//
// See: https://developer.apple.com/documentation/coreservices/1448935-skindexrenamedocument
func SKIndexRenameDocument(inIndex SKIndexRef, inDocument SKDocumentRef, inNewName corefoundation.CFStringRef) bool {
	if _sKIndexRenameDocument == nil {
		panic("coreservices: symbol SKIndexRenameDocument not loaded")
	}
	return _sKIndexRenameDocument(inIndex, inDocument, inNewName)
}

var _sKIndexSetDocumentProperties func(inIndex SKIndexRef, inDocument SKDocumentRef, inProperties corefoundation.CFDictionaryRef)

// SKIndexSetDocumentProperties sets the application-defined properties of a document URL object (of type SKDocument).
//
// See: https://developer.apple.com/documentation/coreservices/1444576-skindexsetdocumentproperties
func SKIndexSetDocumentProperties(inIndex SKIndexRef, inDocument SKDocumentRef, inProperties corefoundation.CFDictionaryRef) {
	if _sKIndexSetDocumentProperties == nil {
		panic("coreservices: symbol SKIndexSetDocumentProperties not loaded")
	}
	_sKIndexSetDocumentProperties(inIndex, inDocument, inProperties)
}

var _sKIndexSetMaximumBytesBeforeFlush func(inIndex SKIndexRef, inBytesForUpdate int)

// SKIndexSetMaximumBytesBeforeFlush not recommended.
//
// See: https://developer.apple.com/documentation/coreservices/1448696-skindexsetmaximumbytesbeforeflus
func SKIndexSetMaximumBytesBeforeFlush(inIndex SKIndexRef, inBytesForUpdate int) {
	if _sKIndexSetMaximumBytesBeforeFlush == nil {
		panic("coreservices: symbol SKIndexSetMaximumBytesBeforeFlush not loaded")
	}
	_sKIndexSetMaximumBytesBeforeFlush(inIndex, inBytesForUpdate)
}

var _sKLoadDefaultExtractorPlugIns func()

// SKLoadDefaultExtractorPlugIns tells Search Kit to use the Spotlight metadata importers.
//
// See: https://developer.apple.com/documentation/coreservices/1447859-skloaddefaultextractorplugins
func SKLoadDefaultExtractorPlugIns() {
	if _sKLoadDefaultExtractorPlugIns == nil {
		panic("coreservices: symbol SKLoadDefaultExtractorPlugIns not loaded")
	}
	_sKLoadDefaultExtractorPlugIns()
}

var _sKSearchCancel func(inSearch SKSearchRef)

// SKSearchCancel cancels an asynchronous search request.
//
// See: https://developer.apple.com/documentation/coreservices/1442083-sksearchcancel
func SKSearchCancel(inSearch SKSearchRef) {
	if _sKSearchCancel == nil {
		panic("coreservices: symbol SKSearchCancel not loaded")
	}
	_sKSearchCancel(inSearch)
}

var _sKSearchCreate func(inIndex SKIndexRef, inQuery corefoundation.CFStringRef, inSearchOptions SKSearchOptions) SKSearchRef

// SKSearchCreate creates an asynchronous search object for querying an index, and initiates search.
//
// See: https://developer.apple.com/documentation/coreservices/1443079-sksearchcreate
func SKSearchCreate(inIndex SKIndexRef, inQuery corefoundation.CFStringRef, inSearchOptions SKSearchOptions) SKSearchRef {
	if _sKSearchCreate == nil {
		panic("coreservices: symbol SKSearchCreate not loaded")
	}
	return _sKSearchCreate(inIndex, inQuery, inSearchOptions)
}

var _sKSearchFindMatches func(inSearch SKSearchRef, inMaximumCount int, outDocumentIDsArray *SKDocumentID, outScoresArray unsafe.Pointer, maximumTime float64, outFoundCount *int) bool

// SKSearchFindMatches extracts search result information from a search object.
//
// See: https://developer.apple.com/documentation/coreservices/1448608-sksearchfindmatches
func SKSearchFindMatches(inSearch SKSearchRef, inMaximumCount int, outDocumentIDsArray *SKDocumentID, outScoresArray unsafe.Pointer, maximumTime float64, outFoundCount *int) bool {
	if _sKSearchFindMatches == nil {
		panic("coreservices: symbol SKSearchFindMatches not loaded")
	}
	return _sKSearchFindMatches(inSearch, inMaximumCount, outDocumentIDsArray, outScoresArray, maximumTime, outFoundCount)
}

var _sKSearchGetTypeID func() uint

// SKSearchGetTypeID gets the type identifier for Search Kit search objects.
//
// See: https://developer.apple.com/documentation/coreservices/1448621-sksearchgettypeid
func SKSearchGetTypeID() uint {
	if _sKSearchGetTypeID == nil {
		panic("coreservices: symbol SKSearchGetTypeID not loaded")
	}
	return _sKSearchGetTypeID()
}

var _sKSearchGroupGetTypeID func() uint

// SKSearchGroupGetTypeID deprecated.
//
// Deprecated: Deprecated since macOS 10.4.
//
// See: https://developer.apple.com/documentation/coreservices/1448637-sksearchgroupgettypeid
func SKSearchGroupGetTypeID() uint {
	if _sKSearchGroupGetTypeID == nil {
		panic("coreservices: symbol SKSearchGroupGetTypeID not loaded")
	}
	return _sKSearchGroupGetTypeID()
}

var _sKSummaryCopyParagraphAtIndex func(summary SKSummaryRef, i int) corefoundation.CFStringRef

// SKSummaryCopyParagraphAtIndex gets a specified paragraph from the text in a summarization object.
//
// See: https://developer.apple.com/documentation/coreservices/1445711-sksummarycopyparagraphatindex
func SKSummaryCopyParagraphAtIndex(summary SKSummaryRef, i int) corefoundation.CFStringRef {
	if _sKSummaryCopyParagraphAtIndex == nil {
		panic("coreservices: symbol SKSummaryCopyParagraphAtIndex not loaded")
	}
	return _sKSummaryCopyParagraphAtIndex(summary, i)
}

var _sKSummaryCopyParagraphSummaryString func(summary SKSummaryRef, numParagraphs int) corefoundation.CFStringRef

// SKSummaryCopyParagraphSummaryString gets a text string consisting of a summary with, at most, the requested number of paragraphs.
//
// See: https://developer.apple.com/documentation/coreservices/1449746-sksummarycopyparagraphsummarystr
func SKSummaryCopyParagraphSummaryString(summary SKSummaryRef, numParagraphs int) corefoundation.CFStringRef {
	if _sKSummaryCopyParagraphSummaryString == nil {
		panic("coreservices: symbol SKSummaryCopyParagraphSummaryString not loaded")
	}
	return _sKSummaryCopyParagraphSummaryString(summary, numParagraphs)
}

var _sKSummaryCopySentenceAtIndex func(summary SKSummaryRef, i int) corefoundation.CFStringRef

// SKSummaryCopySentenceAtIndex gets a specified sentence from the text in a summarization object.
//
// See: https://developer.apple.com/documentation/coreservices/1450287-sksummarycopysentenceatindex
func SKSummaryCopySentenceAtIndex(summary SKSummaryRef, i int) corefoundation.CFStringRef {
	if _sKSummaryCopySentenceAtIndex == nil {
		panic("coreservices: symbol SKSummaryCopySentenceAtIndex not loaded")
	}
	return _sKSummaryCopySentenceAtIndex(summary, i)
}

var _sKSummaryCopySentenceSummaryString func(summary SKSummaryRef, numSentences int) corefoundation.CFStringRef

// SKSummaryCopySentenceSummaryString gets a text string consisting of a summary with, at most, the requested number of sentences.
//
// See: https://developer.apple.com/documentation/coreservices/1449700-sksummarycopysentencesummarystri
func SKSummaryCopySentenceSummaryString(summary SKSummaryRef, numSentences int) corefoundation.CFStringRef {
	if _sKSummaryCopySentenceSummaryString == nil {
		panic("coreservices: symbol SKSummaryCopySentenceSummaryString not loaded")
	}
	return _sKSummaryCopySentenceSummaryString(summary, numSentences)
}

var _sKSummaryCreateWithString func(inString corefoundation.CFStringRef) SKSummaryRef

// SKSummaryCreateWithString creates a summary object based on a text string.
//
// See: https://developer.apple.com/documentation/coreservices/1446229-sksummarycreatewithstring
func SKSummaryCreateWithString(inString corefoundation.CFStringRef) SKSummaryRef {
	if _sKSummaryCreateWithString == nil {
		panic("coreservices: symbol SKSummaryCreateWithString not loaded")
	}
	return _sKSummaryCreateWithString(inString)
}

var _sKSummaryGetParagraphCount func(summary SKSummaryRef) int

// SKSummaryGetParagraphCount gets the number of paragraphs in a summarization object.
//
// See: https://developer.apple.com/documentation/coreservices/1449304-sksummarygetparagraphcount
func SKSummaryGetParagraphCount(summary SKSummaryRef) int {
	if _sKSummaryGetParagraphCount == nil {
		panic("coreservices: symbol SKSummaryGetParagraphCount not loaded")
	}
	return _sKSummaryGetParagraphCount(summary)
}

var _sKSummaryGetParagraphSummaryInfo func(summary SKSummaryRef, numParagraphsInSummary int, outRankOrderOfParagraphs *int, outParagraphIndexOfParagraphs *int) int

// SKSummaryGetParagraphSummaryInfo gets detailed information about a body of text for constructing a custom paragraph-based summary string.
//
// See: https://developer.apple.com/documentation/coreservices/1447517-sksummarygetparagraphsummaryinfo
func SKSummaryGetParagraphSummaryInfo(summary SKSummaryRef, numParagraphsInSummary int, outRankOrderOfParagraphs *int, outParagraphIndexOfParagraphs *int) int {
	if _sKSummaryGetParagraphSummaryInfo == nil {
		panic("coreservices: symbol SKSummaryGetParagraphSummaryInfo not loaded")
	}
	return _sKSummaryGetParagraphSummaryInfo(summary, numParagraphsInSummary, outRankOrderOfParagraphs, outParagraphIndexOfParagraphs)
}

var _sKSummaryGetSentenceCount func(summary SKSummaryRef) int

// SKSummaryGetSentenceCount gets the number of sentences in a summarization object.
//
// See: https://developer.apple.com/documentation/coreservices/1450009-sksummarygetsentencecount
func SKSummaryGetSentenceCount(summary SKSummaryRef) int {
	if _sKSummaryGetSentenceCount == nil {
		panic("coreservices: symbol SKSummaryGetSentenceCount not loaded")
	}
	return _sKSummaryGetSentenceCount(summary)
}

var _sKSummaryGetSentenceSummaryInfo func(summary SKSummaryRef, numSentencesInSummary int, outRankOrderOfSentences *int, outSentenceIndexOfSentences *int, outParagraphIndexOfSentences *int) int

// SKSummaryGetSentenceSummaryInfo gets detailed information about a body of text for constructing a custom sentence-based summary string.
//
// See: https://developer.apple.com/documentation/coreservices/1444767-sksummarygetsentencesummaryinfo
func SKSummaryGetSentenceSummaryInfo(summary SKSummaryRef, numSentencesInSummary int, outRankOrderOfSentences *int, outSentenceIndexOfSentences *int, outParagraphIndexOfSentences *int) int {
	if _sKSummaryGetSentenceSummaryInfo == nil {
		panic("coreservices: symbol SKSummaryGetSentenceSummaryInfo not loaded")
	}
	return _sKSummaryGetSentenceSummaryInfo(summary, numSentencesInSummary, outRankOrderOfSentences, outSentenceIndexOfSentences, outParagraphIndexOfSentences)
}

var _sKSummaryGetTypeID func() uint

// SKSummaryGetTypeID gets the type identifier for Search Kit summarization objects.
//
// See: https://developer.apple.com/documentation/coreservices/1444796-sksummarygettypeid
func SKSummaryGetTypeID() uint {
	if _sKSummaryGetTypeID == nil {
		panic("coreservices: symbol SKSummaryGetTypeID not loaded")
	}
	return _sKSummaryGetTypeID()
}

var _setCollectionDefaultAttributes func(arg0 Collection, arg1 int32, arg2 int32)

// SetCollectionDefaultAttributes.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551442-setcollectiondefaultattributes
func SetCollectionDefaultAttributes(arg0 Collection, arg1 int32, arg2 int32) {
	if _setCollectionDefaultAttributes == nil {
		panic("coreservices: symbol SetCollectionDefaultAttributes not loaded")
	}
	_setCollectionDefaultAttributes(arg0, arg1, arg2)
}

var _setCollectionExceptionProc func(arg0 Collection, arg1 CollectionExceptionUPP)

// SetCollectionExceptionProc.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551417-setcollectionexceptionproc
func SetCollectionExceptionProc(arg0 Collection, arg1 CollectionExceptionUPP) {
	if _setCollectionExceptionProc == nil {
		panic("coreservices: symbol SetCollectionExceptionProc not loaded")
	}
	_setCollectionExceptionProc(arg0, arg1)
}

var _setCollectionItemInfo func(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32, arg4 int32) int16

// SetCollectionItemInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551372-setcollectioniteminfo
func SetCollectionItemInfo(arg0 Collection, arg1 CollectionTag, arg2 int32, arg3 int32, arg4 int32) int16 {
	if _setCollectionItemInfo == nil {
		panic("coreservices: symbol SetCollectionItemInfo not loaded")
	}
	return _setCollectionItemInfo(arg0, arg1, arg2, arg3, arg4)
}

var _setComponentInstanceError func(arg0 ComponentInstance, arg1 int16)

// SetComponentInstanceError passes error information to the Component Manager which sets the current error value for the appropriate connection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516663-setcomponentinstanceerror
func SetComponentInstanceError(arg0 ComponentInstance, arg1 int16) {
	if _setComponentInstanceError == nil {
		panic("coreservices: symbol SetComponentInstanceError not loaded")
	}
	_setComponentInstanceError(arg0, arg1)
}

var _setComponentInstanceStorage func(arg0 ComponentInstance, arg1 uintptr)

// SetComponentInstanceStorage allows your component to associate memory with a connection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516556-setcomponentinstancestorage
func SetComponentInstanceStorage(arg0 ComponentInstance, arg1 uintptr) {
	if _setComponentInstanceStorage == nil {
		panic("coreservices: symbol SetComponentInstanceStorage not loaded")
	}
	_setComponentInstanceStorage(arg0, arg1)
}

var _setComponentRefcon func(arg0 Component, arg1 int)

// SetComponentRefcon sets the reference constant for your component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516562-setcomponentrefcon
func SetComponentRefcon(arg0 Component, arg1 int) {
	if _setComponentRefcon == nil {
		panic("coreservices: symbol SetComponentRefcon not loaded")
	}
	_setComponentRefcon(arg0, arg1)
}

var _setCustomIconsEnabled func(vRefNum unsafe.Pointer, enableCustomIcons unsafe.Pointer) int16

// SetCustomIconsEnabled.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1449302-setcustomiconsenabled
func SetCustomIconsEnabled(vRefNum unsafe.Pointer, enableCustomIcons unsafe.Pointer) int16 {
	if _setCustomIconsEnabled == nil {
		panic("coreservices: symbol SetCustomIconsEnabled not loaded")
	}
	return _setCustomIconsEnabled(vRefNum, enableCustomIcons)
}

var _setDebugOptionValue func(arg0 uint32, arg1 int32, arg2 bool) int32

// SetDebugOptionValue.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517709-setdebugoptionvalue
func SetDebugOptionValue(arg0 uint32, arg1 int32, arg2 bool) int32 {
	if _setDebugOptionValue == nil {
		panic("coreservices: symbol SetDebugOptionValue not loaded")
	}
	return _setDebugOptionValue(arg0, arg1, arg2)
}

var _setDebuggerNotificationProcs func(arg0 DebuggerNewThreadTPP, arg1 DebuggerDisposeThreadTPP, arg2 DebuggerThreadSchedulerTPP) int16

// SetDebuggerNotificationProcs.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574202-setdebuggernotificationprocs
func SetDebuggerNotificationProcs(arg0 DebuggerNewThreadTPP, arg1 DebuggerDisposeThreadTPP, arg2 DebuggerThreadSchedulerTPP) int16 {
	if _setDebuggerNotificationProcs == nil {
		panic("coreservices: symbol SetDebuggerNotificationProcs not loaded")
	}
	return _setDebuggerNotificationProcs(arg0, arg1, arg2)
}

var _setDefaultComponent func(arg0 Component, arg1 int16) int16

// SetDefaultComponent changes the search order for registered components.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516566-setdefaultcomponent
func SetDefaultComponent(arg0 Component, arg1 int16) int16 {
	if _setDefaultComponent == nil {
		panic("coreservices: symbol SetDefaultComponent not loaded")
	}
	return _setDefaultComponent(arg0, arg1)
}

var _setFallbackUnicodeToText func(arg0 UnicodeToTextInfo, arg1 UnicodeToTextFallbackUPP, arg2 uintptr, arg3 unsafe.Pointer) int32

// SetFallbackUnicodeToText specifies a fallback handler to be used for convertinga Unicode text segment to another encoding when the Unicode Convertercannot convert the text using the mapping table specified by theUnicode converter object.
//
// See: https://developer.apple.com/documentation/coreservices/1433614-setfallbackunicodetotext
func SetFallbackUnicodeToText(arg0 UnicodeToTextInfo, arg1 UnicodeToTextFallbackUPP, arg2 uintptr, arg3 unsafe.Pointer) int32 {
	if _setFallbackUnicodeToText == nil {
		panic("coreservices: symbol SetFallbackUnicodeToText not loaded")
	}
	return _setFallbackUnicodeToText(arg0, arg1, arg2, arg3)
}

var _setFallbackUnicodeToTextRun func(arg0 UnicodeToTextRunInfo, arg1 UnicodeToTextFallbackUPP, arg2 uintptr, arg3 unsafe.Pointer) int32

// SetFallbackUnicodeToTextRun specifies a fallback handler to be used for convertinga Unicode text segment to another encoding when the Unicode Convertercannot convert the text using the mapping table specified by a Unicodeconverter object.
//
// See: https://developer.apple.com/documentation/coreservices/1433644-setfallbackunicodetotextrun
func SetFallbackUnicodeToTextRun(arg0 UnicodeToTextRunInfo, arg1 UnicodeToTextFallbackUPP, arg2 uintptr, arg3 unsafe.Pointer) int32 {
	if _setFallbackUnicodeToTextRun == nil {
		panic("coreservices: symbol SetFallbackUnicodeToTextRun not loaded")
	}
	return _setFallbackUnicodeToTextRun(arg0, arg1, arg2, arg3)
}

var _setGestaltValue func(arg0 uint32, arg1 int32) int16

// SetGestaltValue sets the value the function [Gestalt] will return for a specified selector code, installing the selector if it was not already installed.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1470991-setgestaltvalue
func SetGestaltValue(arg0 uint32, arg1 int32) int16 {
	if _setGestaltValue == nil {
		panic("coreservices: symbol SetGestaltValue not loaded")
	}
	return _setGestaltValue(arg0, arg1)
}

var _setHandleSize func(arg0 uintptr, arg1 corefoundation.CGSize)

// SetHandleSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506241-sethandlesize
func SetHandleSize(arg0 uintptr, arg1 corefoundation.CGSize) {
	if _setHandleSize == nil {
		panic("coreservices: symbol SetHandleSize not loaded")
	}
	_setHandleSize(arg0, arg1)
}

var _setIndexedCollectionItemInfo func(arg0 Collection, arg1 int32, arg2 int32, arg3 int32) int16

// SetIndexedCollectionItemInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551406-setindexedcollectioniteminfo
func SetIndexedCollectionItemInfo(arg0 Collection, arg1 int32, arg2 int32, arg3 int32) int16 {
	if _setIndexedCollectionItemInfo == nil {
		panic("coreservices: symbol SetIndexedCollectionItemInfo not loaded")
	}
	return _setIndexedCollectionItemInfo(arg0, arg1, arg2, arg3)
}

var _setPtrSize func(arg0 coreimage.Ptr, arg1 corefoundation.CGSize)

// SetPtrSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506428-setptrsize
func SetPtrSize(arg0 coreimage.Ptr, arg1 corefoundation.CGSize) {
	if _setPtrSize == nil {
		panic("coreservices: symbol SetPtrSize not loaded")
	}
	_setPtrSize(arg0, arg1)
}

var _setResAttrs func(arg0 uintptr, arg1 ResAttributes)

// SetResAttrs.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529240-setresattrs
func SetResAttrs(arg0 uintptr, arg1 ResAttributes) {
	if _setResAttrs == nil {
		panic("coreservices: symbol SetResAttrs not loaded")
	}
	_setResAttrs(arg0, arg1)
}

var _setResFileAttrs func(arg0 ResFileRefNum, arg1 ResFileAttributes)

// SetResFileAttrs.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529257-setresfileattrs
func SetResFileAttrs(arg0 ResFileRefNum, arg1 ResFileAttributes) {
	if _setResFileAttrs == nil {
		panic("coreservices: symbol SetResFileAttrs not loaded")
	}
	_setResFileAttrs(arg0, arg1)
}

var _setResInfo func(arg0 uintptr, arg1 ResID, arg2 unsafe.Pointer)

// SetResInfo.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529247-setresinfo
func SetResInfo(arg0 uintptr, arg1 ResID, arg2 unsafe.Pointer) {
	if _setResInfo == nil {
		panic("coreservices: symbol SetResInfo not loaded")
	}
	_setResInfo(arg0, arg1, arg2)
}

var _setResLoad func(arg0 bool)

// SetResLoad.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529294-setresload
func SetResLoad(arg0 bool) {
	if _setResLoad == nil {
		panic("coreservices: symbol SetResLoad not loaded")
	}
	_setResLoad(arg0)
}

var _setResPurge func(arg0 bool)

// SetResPurge.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529316-setrespurge
func SetResPurge(arg0 bool) {
	if _setResPurge == nil {
		panic("coreservices: symbol SetResPurge not loaded")
	}
	_setResPurge(arg0)
}

var _setResourceSize func(arg0 uintptr, arg1 int)

// SetResourceSize.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529277-setresourcesize
func SetResourceSize(arg0 uintptr, arg1 int) {
	if _setResourceSize == nil {
		panic("coreservices: symbol SetResourceSize not loaded")
	}
	_setResourceSize(arg0, arg1)
}

var _setScriptManagerVariable func(arg0 int16, arg1 int) int16

// SetScriptManagerVariable.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/coreservices/1483931-setscriptmanagervariable
func SetScriptManagerVariable(arg0 int16, arg1 int) int16 {
	if _setScriptManagerVariable == nil {
		panic("coreservices: symbol SetScriptManagerVariable not loaded")
	}
	return _setScriptManagerVariable(arg0, arg1)
}

var _setThreadReadyGivenTaskRef func(arg0 ThreadTaskRef, arg1 ThreadID) int16

// SetThreadReadyGivenTaskRef.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574279-setthreadreadygiventaskref
func SetThreadReadyGivenTaskRef(arg0 ThreadTaskRef, arg1 ThreadID) int16 {
	if _setThreadReadyGivenTaskRef == nil {
		panic("coreservices: symbol SetThreadReadyGivenTaskRef not loaded")
	}
	return _setThreadReadyGivenTaskRef(arg0, arg1)
}

var _setThreadScheduler func(arg0 ThreadSchedulerTPP) int16

// SetThreadScheduler.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574195-setthreadscheduler
func SetThreadScheduler(arg0 ThreadSchedulerTPP) int16 {
	if _setThreadScheduler == nil {
		panic("coreservices: symbol SetThreadScheduler not loaded")
	}
	return _setThreadScheduler(arg0)
}

var _setThreadState func(arg0 ThreadID, arg1 ThreadState, arg2 ThreadID) int16

// SetThreadState.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574212-setthreadstate
func SetThreadState(arg0 ThreadID, arg1 ThreadState, arg2 ThreadID) int16 {
	if _setThreadState == nil {
		panic("coreservices: symbol SetThreadState not loaded")
	}
	return _setThreadState(arg0, arg1, arg2)
}

var _setThreadStateEndCritical func(arg0 ThreadID, arg1 ThreadState, arg2 ThreadID) int16

// SetThreadStateEndCritical.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574289-setthreadstateendcritical
func SetThreadStateEndCritical(arg0 ThreadID, arg1 ThreadState, arg2 ThreadID) int16 {
	if _setThreadStateEndCritical == nil {
		panic("coreservices: symbol SetThreadStateEndCritical not loaded")
	}
	return _setThreadStateEndCritical(arg0, arg1, arg2)
}

var _setThreadSwitcher func(arg0 ThreadID, arg1 ThreadSwitchTPP, arg2 bool) int16

// SetThreadSwitcher.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574270-setthreadswitcher
func SetThreadSwitcher(arg0 ThreadID, arg1 ThreadSwitchTPP, arg2 bool) int16 {
	if _setThreadSwitcher == nil {
		panic("coreservices: symbol SetThreadSwitcher not loaded")
	}
	return _setThreadSwitcher(arg0, arg1, arg2)
}

var _setThreadTerminator func(arg0 ThreadID, arg1 ThreadTerminationTPP) int16

// SetThreadTerminator.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574205-setthreadterminator
func SetThreadTerminator(arg0 ThreadID, arg1 ThreadTerminationTPP) int16 {
	if _setThreadTerminator == nil {
		panic("coreservices: symbol SetThreadTerminator not loaded")
	}
	return _setThreadTerminator(arg0, arg1)
}

var _sleepQInstall func(arg0 SleepQRecPtr)

// SleepQInstall.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427477-sleepqinstall
func SleepQInstall(arg0 SleepQRecPtr) {
	if _sleepQInstall == nil {
		panic("coreservices: symbol SleepQInstall not loaded")
	}
	_sleepQInstall(arg0)
}

var _sleepQRemove func(arg0 SleepQRecPtr)

// SleepQRemove.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427457-sleepqremove
func SleepQRemove(arg0 SleepQRecPtr) {
	if _sleepQRemove == nil {
		panic("coreservices: symbol SleepQRemove not loaded")
	}
	_sleepQRemove(arg0)
}

var _subAbsoluteFromAbsolute func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// SubAbsoluteFromAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501265-subabsolutefromabsolute
func SubAbsoluteFromAbsolute(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _subAbsoluteFromAbsolute == nil {
		panic("coreservices: symbol SubAbsoluteFromAbsolute not loaded")
	}
	return _subAbsoluteFromAbsolute(arg0, arg1)
}

var _subDurationFromAbsolute func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer

// SubDurationFromAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501261-subdurationfromabsolute
func SubDurationFromAbsolute(arg0 unsafe.Pointer, arg1 unsafe.Pointer) unsafe.Pointer {
	if _subDurationFromAbsolute == nil {
		panic("coreservices: symbol SubDurationFromAbsolute not loaded")
	}
	return _subDurationFromAbsolute(arg0, arg1)
}

var _subNanosecondsFromAbsolute func(arg0 Nanoseconds, arg1 unsafe.Pointer) unsafe.Pointer

// SubNanosecondsFromAbsolute.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501240-subnanosecondsfromabsolute
func SubNanosecondsFromAbsolute(arg0 Nanoseconds, arg1 unsafe.Pointer) unsafe.Pointer {
	if _subNanosecondsFromAbsolute == nil {
		panic("coreservices: symbol SubNanosecondsFromAbsolute not loaded")
	}
	return _subNanosecondsFromAbsolute(arg0, arg1)
}

var _sysError func(arg0 int16)

// SysError.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1559937-syserror
func SysError(arg0 int16) {
	if _sysError == nil {
		panic("coreservices: symbol SysError not loaded")
	}
	_sysError(arg0)
}

var _tECClearConverterContextInfo func(arg0 TECObjectRef) int32

// TECClearConverterContextInfo resets a converter object to its initial state so youcan reuse it.
//
// See: https://developer.apple.com/documentation/coreservices/1571844-tecclearconvertercontextinfo
func TECClearConverterContextInfo(arg0 TECObjectRef) int32 {
	if _tECClearConverterContextInfo == nil {
		panic("coreservices: symbol TECClearConverterContextInfo not loaded")
	}
	return _tECClearConverterContextInfo(arg0)
}

var _tECClearSnifferContextInfo func(arg0 TECSnifferObjectRef) int32

// TECClearSnifferContextInfo resets a sniffer object to its initial settings so youcan reuse it.
//
// See: https://developer.apple.com/documentation/coreservices/1571850-tecclearsniffercontextinfo
func TECClearSnifferContextInfo(arg0 TECSnifferObjectRef) int32 {
	if _tECClearSnifferContextInfo == nil {
		panic("coreservices: symbol TECClearSnifferContextInfo not loaded")
	}
	return _tECClearSnifferContextInfo(arg0)
}

var _tECConvertText func(arg0 TECObjectRef, arg1 ConstTextPtr, arg2 uintptr, arg3 uintptr, arg4 TextPtr, arg5 uintptr, arg6 uintptr) int32

// TECConvertText converts a stream of text from a source encoding to adestination encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1571824-tecconverttext
func TECConvertText(arg0 TECObjectRef, arg1 ConstTextPtr, arg2 uintptr, arg3 uintptr, arg4 TextPtr, arg5 uintptr, arg6 uintptr) int32 {
	if _tECConvertText == nil {
		panic("coreservices: symbol TECConvertText not loaded")
	}
	return _tECConvertText(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _tECConvertTextToMultipleEncodings func(arg0 TECObjectRef, arg1 ConstTextPtr, arg2 uintptr, arg3 uintptr, arg4 TextPtr, arg5 uintptr, arg6 uintptr, arg7 TextEncodingRun, arg8 uintptr, arg9 uintptr) int32

// TECConvertTextToMultipleEncodings converts text in the source encoding to runs of text inmultiple destination encodings.
//
// See: https://developer.apple.com/documentation/coreservices/1571849-tecconverttexttomultipleencoding
func TECConvertTextToMultipleEncodings(arg0 TECObjectRef, arg1 ConstTextPtr, arg2 uintptr, arg3 uintptr, arg4 TextPtr, arg5 uintptr, arg6 uintptr, arg7 TextEncodingRun, arg8 uintptr, arg9 uintptr) int32 {
	if _tECConvertTextToMultipleEncodings == nil {
		panic("coreservices: symbol TECConvertTextToMultipleEncodings not loaded")
	}
	return _tECConvertTextToMultipleEncodings(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

var _tECCopyTextEncodingInternetNameAndMIB func(arg0 TextEncoding, arg1 TECInternetNameUsageMask, arg2 corefoundation.CFStringRef, arg3 int32) int32

// TECCopyTextEncodingInternetNameAndMIB.
//
// See: https://developer.apple.com/documentation/coreservices/1571835-teccopytextencodinginternetnamea
func TECCopyTextEncodingInternetNameAndMIB(arg0 TextEncoding, arg1 TECInternetNameUsageMask, arg2 corefoundation.CFStringRef, arg3 int32) int32 {
	if _tECCopyTextEncodingInternetNameAndMIB == nil {
		panic("coreservices: symbol TECCopyTextEncodingInternetNameAndMIB not loaded")
	}
	return _tECCopyTextEncodingInternetNameAndMIB(arg0, arg1, arg2, arg3)
}

var _tECCountAvailableSniffers func(arg0 uintptr) int32

// TECCountAvailableSniffers counts and returns the number of sniffers available inall installed plug-ins.
//
// See: https://developer.apple.com/documentation/coreservices/1571795-teccountavailablesniffers
func TECCountAvailableSniffers(arg0 uintptr) int32 {
	if _tECCountAvailableSniffers == nil {
		panic("coreservices: symbol TECCountAvailableSniffers not loaded")
	}
	return _tECCountAvailableSniffers(arg0)
}

var _tECCountAvailableTextEncodings func(arg0 uintptr) int32

// TECCountAvailableTextEncodings counts and returns the number of text encodings currentlyconfigured in the Text Encoding Converter.
//
// See: https://developer.apple.com/documentation/coreservices/1571853-teccountavailabletextencodings
func TECCountAvailableTextEncodings(arg0 uintptr) int32 {
	if _tECCountAvailableTextEncodings == nil {
		panic("coreservices: symbol TECCountAvailableTextEncodings not loaded")
	}
	return _tECCountAvailableTextEncodings(arg0)
}

var _tECCountDestinationTextEncodings func(arg0 TextEncoding, arg1 uintptr) int32

// TECCountDestinationTextEncodings counts and returns the number of destination encodingsto which a specified source encoding can be converted in one step.
//
// See: https://developer.apple.com/documentation/coreservices/1571805-teccountdestinationtextencodings
func TECCountDestinationTextEncodings(arg0 TextEncoding, arg1 uintptr) int32 {
	if _tECCountDestinationTextEncodings == nil {
		panic("coreservices: symbol TECCountDestinationTextEncodings not loaded")
	}
	return _tECCountDestinationTextEncodings(arg0, arg1)
}

var _tECCountDirectTextEncodingConversions func(arg0 uintptr) int32

// TECCountDirectTextEncodingConversions counts and returns the number of direct conversions currentlyconfigured in the Text Encoding Converter.
//
// See: https://developer.apple.com/documentation/coreservices/1571822-teccountdirecttextencodingconver
func TECCountDirectTextEncodingConversions(arg0 uintptr) int32 {
	if _tECCountDirectTextEncodingConversions == nil {
		panic("coreservices: symbol TECCountDirectTextEncodingConversions not loaded")
	}
	return _tECCountDirectTextEncodingConversions(arg0)
}

var _tECCountMailTextEncodings func(arg0 unsafe.Pointer, arg1 uintptr) int32

// TECCountMailTextEncodings counts and returns the number of currently supported e-mailencodings for a specified region.
//
// See: https://developer.apple.com/documentation/coreservices/1571826-teccountmailtextencodings
func TECCountMailTextEncodings(arg0 unsafe.Pointer, arg1 uintptr) int32 {
	if _tECCountMailTextEncodings == nil {
		panic("coreservices: symbol TECCountMailTextEncodings not loaded")
	}
	return _tECCountMailTextEncodings(arg0, arg1)
}

var _tECCountSubTextEncodings func(arg0 TextEncoding, arg1 uintptr) int32

// TECCountSubTextEncodings counts and returns the number of subencodings a text encodingsupports.
//
// See: https://developer.apple.com/documentation/coreservices/1571820-teccountsubtextencodings
func TECCountSubTextEncodings(arg0 TextEncoding, arg1 uintptr) int32 {
	if _tECCountSubTextEncodings == nil {
		panic("coreservices: symbol TECCountSubTextEncodings not loaded")
	}
	return _tECCountSubTextEncodings(arg0, arg1)
}

var _tECCountWebTextEncodings func(arg0 unsafe.Pointer, arg1 uintptr) int32

// TECCountWebTextEncodings counts and returns the number of currently supported textencodings for a region code.
//
// See: https://developer.apple.com/documentation/coreservices/1571837-teccountwebtextencodings
func TECCountWebTextEncodings(arg0 unsafe.Pointer, arg1 uintptr) int32 {
	if _tECCountWebTextEncodings == nil {
		panic("coreservices: symbol TECCountWebTextEncodings not loaded")
	}
	return _tECCountWebTextEncodings(arg0, arg1)
}

var _tECCreateConverter func(arg0 TECObjectRef, arg1 TextEncoding, arg2 TextEncoding) int32

// TECCreateConverter determines a conversion path for a source and destinationencoding, then creates a text encoding converter object and returnsa pointer to it.
//
// See: https://developer.apple.com/documentation/coreservices/1571815-teccreateconverter
func TECCreateConverter(arg0 TECObjectRef, arg1 TextEncoding, arg2 TextEncoding) int32 {
	if _tECCreateConverter == nil {
		panic("coreservices: symbol TECCreateConverter not loaded")
	}
	return _tECCreateConverter(arg0, arg1, arg2)
}

var _tECCreateConverterFromPath func(arg0 TECObjectRef, arg1 TextEncoding, arg2 uintptr) int32

// TECCreateConverterFromPath creates a converter object for a specific conversion path—froma source encoding through intermediate encodings to a destinationencoding—and returns a pointer to it.
//
// See: https://developer.apple.com/documentation/coreservices/1571829-teccreateconverterfrompath
func TECCreateConverterFromPath(arg0 TECObjectRef, arg1 TextEncoding, arg2 uintptr) int32 {
	if _tECCreateConverterFromPath == nil {
		panic("coreservices: symbol TECCreateConverterFromPath not loaded")
	}
	return _tECCreateConverterFromPath(arg0, arg1, arg2)
}

var _tECCreateOneToManyConverter func(arg0 TECObjectRef, arg1 TextEncoding, arg2 uintptr, arg3 TextEncoding) int32

// TECCreateOneToManyConverter determines a conversion path for the source encoding anddestinations encodings you specify, creates a text encoding converterobject, and returns a reference to it.
//
// See: https://developer.apple.com/documentation/coreservices/1571794-teccreateonetomanyconverter
func TECCreateOneToManyConverter(arg0 TECObjectRef, arg1 TextEncoding, arg2 uintptr, arg3 TextEncoding) int32 {
	if _tECCreateOneToManyConverter == nil {
		panic("coreservices: symbol TECCreateOneToManyConverter not loaded")
	}
	return _tECCreateOneToManyConverter(arg0, arg1, arg2, arg3)
}

var _tECCreateSniffer func(arg0 TECSnifferObjectRef, arg1 TextEncoding, arg2 uintptr) int32

// TECCreateSniffer creates a sniffer object and returns a reference to it.
//
// See: https://developer.apple.com/documentation/coreservices/1571832-teccreatesniffer
func TECCreateSniffer(arg0 TECSnifferObjectRef, arg1 TextEncoding, arg2 uintptr) int32 {
	if _tECCreateSniffer == nil {
		panic("coreservices: symbol TECCreateSniffer not loaded")
	}
	return _tECCreateSniffer(arg0, arg1, arg2)
}

var _tECDisposeConverter func(arg0 TECObjectRef) int32

// TECDisposeConverter disposes of a converter object.
//
// See: https://developer.apple.com/documentation/coreservices/1571839-tecdisposeconverter
func TECDisposeConverter(arg0 TECObjectRef) int32 {
	if _tECDisposeConverter == nil {
		panic("coreservices: symbol TECDisposeConverter not loaded")
	}
	return _tECDisposeConverter(arg0)
}

var _tECDisposeSniffer func(arg0 TECSnifferObjectRef) int32

// TECDisposeSniffer disposes of a sniffer object.
//
// See: https://developer.apple.com/documentation/coreservices/1571854-tecdisposesniffer
func TECDisposeSniffer(arg0 TECSnifferObjectRef) int32 {
	if _tECDisposeSniffer == nil {
		panic("coreservices: symbol TECDisposeSniffer not loaded")
	}
	return _tECDisposeSniffer(arg0)
}

var _tECFlushMultipleEncodings func(arg0 TECObjectRef, arg1 TextPtr, arg2 uintptr, arg3 uintptr, arg4 TextEncodingRun, arg5 uintptr, arg6 uintptr) int32

// TECFlushMultipleEncodings flushes out any encodings that may be stored in a converterobject’s temporary buffers and shifts encodings back to theirdefault state, if any.
//
// See: https://developer.apple.com/documentation/coreservices/1571842-tecflushmultipleencodings
func TECFlushMultipleEncodings(arg0 TECObjectRef, arg1 TextPtr, arg2 uintptr, arg3 uintptr, arg4 TextEncodingRun, arg5 uintptr, arg6 uintptr) int32 {
	if _tECFlushMultipleEncodings == nil {
		panic("coreservices: symbol TECFlushMultipleEncodings not loaded")
	}
	return _tECFlushMultipleEncodings(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _tECFlushText func(arg0 TECObjectRef, arg1 TextPtr, arg2 uintptr, arg3 uintptr) int32

// TECFlushText flushes out any data in a converter object’s temporarybuffers and resets the converter object.
//
// See: https://developer.apple.com/documentation/coreservices/1571848-tecflushtext
func TECFlushText(arg0 TECObjectRef, arg1 TextPtr, arg2 uintptr, arg3 uintptr) int32 {
	if _tECFlushText == nil {
		panic("coreservices: symbol TECFlushText not loaded")
	}
	return _tECFlushText(arg0, arg1, arg2, arg3)
}

var _tECGetAvailableSniffers func(arg0 TextEncoding, arg1 uintptr, arg2 uintptr) int32

// TECGetAvailableSniffers returns the list of sniffers available in all installedplug-ins.
//
// See: https://developer.apple.com/documentation/coreservices/1571841-tecgetavailablesniffers
func TECGetAvailableSniffers(arg0 TextEncoding, arg1 uintptr, arg2 uintptr) int32 {
	if _tECGetAvailableSniffers == nil {
		panic("coreservices: symbol TECGetAvailableSniffers not loaded")
	}
	return _tECGetAvailableSniffers(arg0, arg1, arg2)
}

var _tECGetAvailableTextEncodings func(arg0 TextEncoding, arg1 uintptr, arg2 uintptr) int32

// TECGetAvailableTextEncodings returns the text encoding specifications currently configuredin the Text Encoding Converter.
//
// See: https://developer.apple.com/documentation/coreservices/1571819-tecgetavailabletextencodings
func TECGetAvailableTextEncodings(arg0 TextEncoding, arg1 uintptr, arg2 uintptr) int32 {
	if _tECGetAvailableTextEncodings == nil {
		panic("coreservices: symbol TECGetAvailableTextEncodings not loaded")
	}
	return _tECGetAvailableTextEncodings(arg0, arg1, arg2)
}

var _tECGetDestinationTextEncodings func(arg0 TextEncoding, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32

// TECGetDestinationTextEncodings returns the encoding specifications for all the destinationtext encodings to which the Text Encoding Converter can directlyconvert the specified source encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1571808-tecgetdestinationtextencodings
func TECGetDestinationTextEncodings(arg0 TextEncoding, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32 {
	if _tECGetDestinationTextEncodings == nil {
		panic("coreservices: symbol TECGetDestinationTextEncodings not loaded")
	}
	return _tECGetDestinationTextEncodings(arg0, arg1, arg2, arg3)
}

var _tECGetDirectTextEncodingConversions func(arg0 TECConversionInfo, arg1 uintptr, arg2 uintptr) int32

// TECGetDirectTextEncodingConversions returns the types of direct conversions currently configuredin the Text Encoding Converter.
//
// See: https://developer.apple.com/documentation/coreservices/1571834-tecgetdirecttextencodingconversi
func TECGetDirectTextEncodingConversions(arg0 TECConversionInfo, arg1 uintptr, arg2 uintptr) int32 {
	if _tECGetDirectTextEncodingConversions == nil {
		panic("coreservices: symbol TECGetDirectTextEncodingConversions not loaded")
	}
	return _tECGetDirectTextEncodingConversions(arg0, arg1, arg2)
}

var _tECGetEncodingList func(arg0 TECObjectRef, arg1 uintptr, arg2 uintptr) int32

// TECGetEncodingList gets the list of destination encodings from a converterobject.
//
// See: https://developer.apple.com/documentation/coreservices/1571830-tecgetencodinglist
func TECGetEncodingList(arg0 TECObjectRef, arg1 uintptr, arg2 uintptr) int32 {
	if _tECGetEncodingList == nil {
		panic("coreservices: symbol TECGetEncodingList not loaded")
	}
	return _tECGetEncodingList(arg0, arg1, arg2)
}

var _tECGetInfo func(arg0 TECInfoHandle) int32

// TECGetInfo allocates a converter information structure of type [TECInfo] inthe application heap using [NewHandle],fills it out, and returns a handle.
//
// See: https://developer.apple.com/documentation/coreservices/1400430-tecgetinfo
func TECGetInfo(arg0 TECInfoHandle) int32 {
	if _tECGetInfo == nil {
		panic("coreservices: symbol TECGetInfo not loaded")
	}
	return _tECGetInfo(arg0)
}

var _tECGetMailTextEncodings func(arg0 unsafe.Pointer, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32

// TECGetMailTextEncodings returns the currently supported mail encoding specificationsfor a region code.
//
// See: https://developer.apple.com/documentation/coreservices/1571845-tecgetmailtextencodings
func TECGetMailTextEncodings(arg0 unsafe.Pointer, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32 {
	if _tECGetMailTextEncodings == nil {
		panic("coreservices: symbol TECGetMailTextEncodings not loaded")
	}
	return _tECGetMailTextEncodings(arg0, arg1, arg2, arg3)
}

var _tECGetSubTextEncodings func(arg0 TextEncoding, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32

// TECGetSubTextEncodings returns the text encoding specifications for the subencodingsthe encoding scheme supports.
//
// See: https://developer.apple.com/documentation/coreservices/1571796-tecgetsubtextencodings
func TECGetSubTextEncodings(arg0 TextEncoding, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32 {
	if _tECGetSubTextEncodings == nil {
		panic("coreservices: symbol TECGetSubTextEncodings not loaded")
	}
	return _tECGetSubTextEncodings(arg0, arg1, arg2, arg3)
}

var _tECGetTextEncodingFromInternetName func(arg0 TextEncoding, arg1 unsafe.Pointer) int32

// TECGetTextEncodingFromInternetName returns the Mac OS text encoding specification that correspondsto an Internet encoding name.
//
// See: https://developer.apple.com/documentation/coreservices/1571825-tecgettextencodingfrominternetna
func TECGetTextEncodingFromInternetName(arg0 TextEncoding, arg1 unsafe.Pointer) int32 {
	if _tECGetTextEncodingFromInternetName == nil {
		panic("coreservices: symbol TECGetTextEncodingFromInternetName not loaded")
	}
	return _tECGetTextEncodingFromInternetName(arg0, arg1)
}

var _tECGetTextEncodingFromInternetNameOrMIB func(arg0 TextEncoding, arg1 TECInternetNameUsageMask, arg2 corefoundation.CFStringRef, arg3 int32) int32

// TECGetTextEncodingFromInternetNameOrMIB.
//
// See: https://developer.apple.com/documentation/coreservices/1571816-tecgettextencodingfrominternetna
func TECGetTextEncodingFromInternetNameOrMIB(arg0 TextEncoding, arg1 TECInternetNameUsageMask, arg2 corefoundation.CFStringRef, arg3 int32) int32 {
	if _tECGetTextEncodingFromInternetNameOrMIB == nil {
		panic("coreservices: symbol TECGetTextEncodingFromInternetNameOrMIB not loaded")
	}
	return _tECGetTextEncodingFromInternetNameOrMIB(arg0, arg1, arg2, arg3)
}

var _tECGetTextEncodingInternetName func(arg0 TextEncoding, arg1 unsafe.Pointer) int32

// TECGetTextEncodingInternetName returns the Internet encoding name that corresponds toa Mac OS text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1571851-tecgettextencodinginternetname
func TECGetTextEncodingInternetName(arg0 TextEncoding, arg1 unsafe.Pointer) int32 {
	if _tECGetTextEncodingInternetName == nil {
		panic("coreservices: symbol TECGetTextEncodingInternetName not loaded")
	}
	return _tECGetTextEncodingInternetName(arg0, arg1)
}

var _tECGetWebTextEncodings func(arg0 unsafe.Pointer, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32

// TECGetWebTextEncodings returns the currently supported text encoding specificationsfor a region code.
//
// See: https://developer.apple.com/documentation/coreservices/1571798-tecgetwebtextencodings
func TECGetWebTextEncodings(arg0 unsafe.Pointer, arg1 TextEncoding, arg2 uintptr, arg3 uintptr) int32 {
	if _tECGetWebTextEncodings == nil {
		panic("coreservices: symbol TECGetWebTextEncodings not loaded")
	}
	return _tECGetWebTextEncodings(arg0, arg1, arg2, arg3)
}

var _tECSetBasicOptions func(arg0 TECObjectRef, arg1 uintptr) int32

// TECSetBasicOptions.
//
// See: https://developer.apple.com/documentation/coreservices/1571847-tecsetbasicoptions
func TECSetBasicOptions(arg0 TECObjectRef, arg1 uintptr) int32 {
	if _tECSetBasicOptions == nil {
		panic("coreservices: symbol TECSetBasicOptions not loaded")
	}
	return _tECSetBasicOptions(arg0, arg1)
}

var _tECSniffTextEncoding func(arg0 TECSnifferObjectRef, arg1 ConstTextPtr, arg2 uintptr, arg3 TextEncoding, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr) int32

// TECSniffTextEncoding analyzes a text stream and returns the probable encodingsin a ranked list, based on an array of possible encodings you supply.It also returns the number of errors and features for each encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1571836-tecsnifftextencoding
func TECSniffTextEncoding(arg0 TECSnifferObjectRef, arg1 ConstTextPtr, arg2 uintptr, arg3 TextEncoding, arg4 uintptr, arg5 uintptr, arg6 uintptr, arg7 uintptr, arg8 uintptr) int32 {
	if _tECSniffTextEncoding == nil {
		panic("coreservices: symbol TECSniffTextEncoding not loaded")
	}
	return _tECSniffTextEncoding(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _taskLevel func() uint32

// TaskLevel.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1517734-tasklevel
func TaskLevel() uint32 {
	if _taskLevel == nil {
		panic("coreservices: symbol TaskLevel not loaded")
	}
	return _taskLevel()
}

var _tempNewHandle func(arg0 corefoundation.CGSize, arg1 int16) objectivec.IObject

// TempNewHandle.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1506376-tempnewhandle
func TempNewHandle(arg0 corefoundation.CGSize, arg1 int16) objectivec.IObject {
	if _tempNewHandle == nil {
		panic("coreservices: symbol TempNewHandle not loaded")
	}
	return _tempNewHandle(arg0, arg1)
}

var _testAndClear func(arg0 uint32, arg1 uint8) bool

// TestAndClear.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490573-testandclear
func TestAndClear(arg0 uint32, arg1 uint8) bool {
	if _testAndClear == nil {
		panic("coreservices: symbol TestAndClear not loaded")
	}
	return _testAndClear(arg0, arg1)
}

var _testAndSet func(arg0 uint32, arg1 uint8) bool

// TestAndSet.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1490582-testandset
func TestAndSet(arg0 uint32, arg1 uint8) bool {
	if _testAndSet == nil {
		panic("coreservices: symbol TestAndSet not loaded")
	}
	return _testAndSet(arg0, arg1)
}

var _threadBeginCritical func() int16

// ThreadBeginCritical.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574258-threadbegincritical
func ThreadBeginCritical() int16 {
	if _threadBeginCritical == nil {
		panic("coreservices: symbol ThreadBeginCritical not loaded")
	}
	return _threadBeginCritical()
}

var _threadCurrentStackSpace func(arg0 ThreadID, arg1 uintptr) int16

// ThreadCurrentStackSpace.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574269-threadcurrentstackspace
func ThreadCurrentStackSpace(arg0 ThreadID, arg1 uintptr) int16 {
	if _threadCurrentStackSpace == nil {
		panic("coreservices: symbol ThreadCurrentStackSpace not loaded")
	}
	return _threadCurrentStackSpace(arg0, arg1)
}

var _threadEndCritical func() int16

// ThreadEndCritical.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574230-threadendcritical
func ThreadEndCritical() int16 {
	if _threadEndCritical == nil {
		panic("coreservices: symbol ThreadEndCritical not loaded")
	}
	return _threadEndCritical()
}

var _tickCount func() uint32

// TickCount.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1533365-tickcount
func TickCount() uint32 {
	if _tickCount == nil {
		panic("coreservices: symbol TickCount not loaded")
	}
	return _tickCount()
}

var _truncateForTextToUnicode func(arg0 ConstTextToUnicodeInfo, arg1 uintptr, arg2 unsafe.Pointer, arg3 uintptr, arg4 uintptr) int32

// TruncateForTextToUnicode identifies where your application can safely break a multibytestring to be converted to Unicode so that the string is not brokenin the middle of a multibyte character.
//
// See: https://developer.apple.com/documentation/coreservices/1433518-truncatefortexttounicode
func TruncateForTextToUnicode(arg0 ConstTextToUnicodeInfo, arg1 uintptr, arg2 unsafe.Pointer, arg3 uintptr, arg4 uintptr) int32 {
	if _truncateForTextToUnicode == nil {
		panic("coreservices: symbol TruncateForTextToUnicode not loaded")
	}
	return _truncateForTextToUnicode(arg0, arg1, arg2, arg3, arg4)
}

var _truncateForUnicodeToText func(arg0 ConstUnicodeToTextInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr) int32

// TruncateForUnicodeToText identifies where your application can safely break a Unicodestring to be converted to any encoding so that the string is brokenin a way that preserves the text element integrity.
//
// See: https://developer.apple.com/documentation/coreservices/1433649-truncateforunicodetotext
func TruncateForUnicodeToText(arg0 ConstUnicodeToTextInfo, arg1 uintptr, arg2 uint16, arg3 uintptr, arg4 uintptr, arg5 uintptr) int32 {
	if _truncateForUnicodeToText == nil {
		panic("coreservices: symbol TruncateForUnicodeToText not loaded")
	}
	return _truncateForUnicodeToText(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _u32SetU func(arg0 uint64) uint32

// U32SetU.
//
// See: https://developer.apple.com/documentation/coreservices/1536258-u32setu
func U32SetU(arg0 uint64) uint32 {
	if _u32SetU == nil {
		panic("coreservices: symbol U32SetU not loaded")
	}
	return _u32SetU(arg0)
}

var _u64Add func(arg0 uint64, arg1 uint64) uint64

// U64Add.
//
// See: https://developer.apple.com/documentation/coreservices/1536306-u64add
func U64Add(arg0 uint64, arg1 uint64) uint64 {
	if _u64Add == nil {
		panic("coreservices: symbol U64Add not loaded")
	}
	return _u64Add(arg0, arg1)
}

var _u64And func(arg0 uint64, arg1 uint64) bool

// U64And.
//
// See: https://developer.apple.com/documentation/coreservices/1536257-u64and
func U64And(arg0 uint64, arg1 uint64) bool {
	if _u64And == nil {
		panic("coreservices: symbol U64And not loaded")
	}
	return _u64And(arg0, arg1)
}

var _u64BitwiseAnd func(arg0 uint64, arg1 uint64) uint64

// U64BitwiseAnd.
//
// See: https://developer.apple.com/documentation/coreservices/1536227-u64bitwiseand
func U64BitwiseAnd(arg0 uint64, arg1 uint64) uint64 {
	if _u64BitwiseAnd == nil {
		panic("coreservices: symbol U64BitwiseAnd not loaded")
	}
	return _u64BitwiseAnd(arg0, arg1)
}

var _u64BitwiseEor func(arg0 uint64, arg1 uint64) uint64

// U64BitwiseEor.
//
// See: https://developer.apple.com/documentation/coreservices/1536230-u64bitwiseeor
func U64BitwiseEor(arg0 uint64, arg1 uint64) uint64 {
	if _u64BitwiseEor == nil {
		panic("coreservices: symbol U64BitwiseEor not loaded")
	}
	return _u64BitwiseEor(arg0, arg1)
}

var _u64BitwiseNot func(arg0 uint64) uint64

// U64BitwiseNot.
//
// See: https://developer.apple.com/documentation/coreservices/1536304-u64bitwisenot
func U64BitwiseNot(arg0 uint64) uint64 {
	if _u64BitwiseNot == nil {
		panic("coreservices: symbol U64BitwiseNot not loaded")
	}
	return _u64BitwiseNot(arg0)
}

var _u64BitwiseOr func(arg0 uint64, arg1 uint64) uint64

// U64BitwiseOr.
//
// See: https://developer.apple.com/documentation/coreservices/1536365-u64bitwiseor
func U64BitwiseOr(arg0 uint64, arg1 uint64) uint64 {
	if _u64BitwiseOr == nil {
		panic("coreservices: symbol U64BitwiseOr not loaded")
	}
	return _u64BitwiseOr(arg0, arg1)
}

var _u64Compare func(arg0 uint64, arg1 uint64) int32

// U64Compare.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1536343-u64compare
func U64Compare(arg0 uint64, arg1 uint64) int32 {
	if _u64Compare == nil {
		panic("coreservices: symbol U64Compare not loaded")
	}
	return _u64Compare(arg0, arg1)
}

var _u64Div func(arg0 uint64, arg1 uint64) uint64

// U64Div.
//
// See: https://developer.apple.com/documentation/coreservices/1536271-u64div
func U64Div(arg0 uint64, arg1 uint64) uint64 {
	if _u64Div == nil {
		panic("coreservices: symbol U64Div not loaded")
	}
	return _u64Div(arg0, arg1)
}

var _u64Divide func(arg0 uint64, arg1 uint64, arg2 uint64) uint64

// U64Divide.
//
// See: https://developer.apple.com/documentation/coreservices/1536326-u64divide
func U64Divide(arg0 uint64, arg1 uint64, arg2 uint64) uint64 {
	if _u64Divide == nil {
		panic("coreservices: symbol U64Divide not loaded")
	}
	return _u64Divide(arg0, arg1, arg2)
}

var _u64Eor func(arg0 uint64, arg1 uint64) bool

// U64Eor.
//
// See: https://developer.apple.com/documentation/coreservices/1536351-u64eor
func U64Eor(arg0 uint64, arg1 uint64) bool {
	if _u64Eor == nil {
		panic("coreservices: symbol U64Eor not loaded")
	}
	return _u64Eor(arg0, arg1)
}

var _u64Max func() uint64

// U64Max.
//
// See: https://developer.apple.com/documentation/coreservices/1536321-u64max
func U64Max() uint64 {
	if _u64Max == nil {
		panic("coreservices: symbol U64Max not loaded")
	}
	return _u64Max()
}

var _u64Mod func(arg0 uint64, arg1 uint64) uint64

// U64Mod.
//
// See: https://developer.apple.com/documentation/coreservices/1536225-u64mod
func U64Mod(arg0 uint64, arg1 uint64) uint64 {
	if _u64Mod == nil {
		panic("coreservices: symbol U64Mod not loaded")
	}
	return _u64Mod(arg0, arg1)
}

var _u64Multiply func(arg0 uint64, arg1 uint64) uint64

// U64Multiply.
//
// See: https://developer.apple.com/documentation/coreservices/1536288-u64multiply
func U64Multiply(arg0 uint64, arg1 uint64) uint64 {
	if _u64Multiply == nil {
		panic("coreservices: symbol U64Multiply not loaded")
	}
	return _u64Multiply(arg0, arg1)
}

var _u64Not func(arg0 uint64) bool

// U64Not.
//
// See: https://developer.apple.com/documentation/coreservices/1536276-u64not
func U64Not(arg0 uint64) bool {
	if _u64Not == nil {
		panic("coreservices: symbol U64Not not loaded")
	}
	return _u64Not(arg0)
}

var _u64Or func(arg0 uint64, arg1 uint64) bool

// U64Or.
//
// See: https://developer.apple.com/documentation/coreservices/1536336-u64or
func U64Or(arg0 uint64, arg1 uint64) bool {
	if _u64Or == nil {
		panic("coreservices: symbol U64Or not loaded")
	}
	return _u64Or(arg0, arg1)
}

var _u64Set func(arg0 int32) uint64

// U64Set.
//
// See: https://developer.apple.com/documentation/coreservices/1536364-u64set
func U64Set(arg0 int32) uint64 {
	if _u64Set == nil {
		panic("coreservices: symbol U64Set not loaded")
	}
	return _u64Set(arg0)
}

var _u64SetU func(arg0 uint32) uint64

// U64SetU.
//
// See: https://developer.apple.com/documentation/coreservices/1536282-u64setu
func U64SetU(arg0 uint32) uint64 {
	if _u64SetU == nil {
		panic("coreservices: symbol U64SetU not loaded")
	}
	return _u64SetU(arg0)
}

var _u64ShiftLeft func(arg0 uint64, arg1 uint32) uint64

// U64ShiftLeft.
//
// See: https://developer.apple.com/documentation/coreservices/1536248-u64shiftleft
func U64ShiftLeft(arg0 uint64, arg1 uint32) uint64 {
	if _u64ShiftLeft == nil {
		panic("coreservices: symbol U64ShiftLeft not loaded")
	}
	return _u64ShiftLeft(arg0, arg1)
}

var _u64ShiftRight func(arg0 uint64, arg1 uint32) uint64

// U64ShiftRight.
//
// See: https://developer.apple.com/documentation/coreservices/1536363-u64shiftright
func U64ShiftRight(arg0 uint64, arg1 uint32) uint64 {
	if _u64ShiftRight == nil {
		panic("coreservices: symbol U64ShiftRight not loaded")
	}
	return _u64ShiftRight(arg0, arg1)
}

var _u64Subtract func(arg0 uint64, arg1 uint64) uint64

// U64Subtract.
//
// See: https://developer.apple.com/documentation/coreservices/1536348-u64subtract
func U64Subtract(arg0 uint64, arg1 uint64) uint64 {
	if _u64Subtract == nil {
		panic("coreservices: symbol U64Subtract not loaded")
	}
	return _u64Subtract(arg0, arg1)
}

var _uCCompareCollationKeys func(key1Ptr *UCCollationValue, key1Length unsafe.Pointer, key2Ptr *UCCollationValue, key2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32

// UCCompareCollationKeys uses collation keys to compare Unicode strings.
//
// See: https://developer.apple.com/documentation/coreservices/1390378-uccomparecollationkeys
func UCCompareCollationKeys(key1Ptr *UCCollationValue, key1Length unsafe.Pointer, key2Ptr *UCCollationValue, key2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32 {
	if _uCCompareCollationKeys == nil {
		panic("coreservices: symbol UCCompareCollationKeys not loaded")
	}
	return _uCCompareCollationKeys(key1Ptr, key1Length, key2Ptr, key2Length, equivalent, order)
}

var _uCCompareText func(collatorRef CollatorRef, text1Ptr *uint16, text1Length unsafe.Pointer, text2Ptr *uint16, text2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32

// UCCompareText uses locale-specific collation information to compare Unicode strings.
//
// See: https://developer.apple.com/documentation/coreservices/1390642-uccomparetext
func UCCompareText(collatorRef CollatorRef, text1Ptr *uint16, text1Length unsafe.Pointer, text2Ptr *uint16, text2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32 {
	if _uCCompareText == nil {
		panic("coreservices: symbol UCCompareText not loaded")
	}
	return _uCCompareText(collatorRef, text1Ptr, text1Length, text2Ptr, text2Length, equivalent, order)
}

var _uCCompareTextDefault func(options UCCollateOptions, text1Ptr *uint16, text1Length unsafe.Pointer, text2Ptr *uint16, text2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32

// UCCompareTextDefault uses the default system locale to compare Unicode strings.
//
// See: https://developer.apple.com/documentation/coreservices/1390472-uccomparetextdefault
func UCCompareTextDefault(options UCCollateOptions, text1Ptr *uint16, text1Length unsafe.Pointer, text2Ptr *uint16, text2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32 {
	if _uCCompareTextDefault == nil {
		panic("coreservices: symbol UCCompareTextDefault not loaded")
	}
	return _uCCompareTextDefault(options, text1Ptr, text1Length, text2Ptr, text2Length, equivalent, order)
}

var _uCCompareTextNoLocale func(options UCCollateOptions, text1Ptr *uint16, text1Length unsafe.Pointer, text2Ptr *uint16, text2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32

// UCCompareTextNoLocale uses a fixed, locale-insensitive order to compare Unicode strings.
//
// See: https://developer.apple.com/documentation/coreservices/1390513-uccomparetextnolocale
func UCCompareTextNoLocale(options UCCollateOptions, text1Ptr *uint16, text1Length unsafe.Pointer, text2Ptr *uint16, text2Length unsafe.Pointer, equivalent unsafe.Pointer, order *uintptr) int32 {
	if _uCCompareTextNoLocale == nil {
		panic("coreservices: symbol UCCompareTextNoLocale not loaded")
	}
	return _uCCompareTextNoLocale(options, text1Ptr, text1Length, text2Ptr, text2Length, equivalent, order)
}

var _uCConvertCFAbsoluteTimeToLongDateTime func(arg0 corefoundation.CFAbsoluteTime, arg1 LongDateTime) int32

// UCConvertCFAbsoluteTimeToLongDateTime.
//
// See: https://developer.apple.com/documentation/coreservices/1551577-ucconvertcfabsolutetimetolongdat
func UCConvertCFAbsoluteTimeToLongDateTime(arg0 corefoundation.CFAbsoluteTime, arg1 LongDateTime) int32 {
	if _uCConvertCFAbsoluteTimeToLongDateTime == nil {
		panic("coreservices: symbol UCConvertCFAbsoluteTimeToLongDateTime not loaded")
	}
	return _uCConvertCFAbsoluteTimeToLongDateTime(arg0, arg1)
}

var _uCConvertCFAbsoluteTimeToSeconds func(arg0 corefoundation.CFAbsoluteTime, arg1 uint32) int32

// UCConvertCFAbsoluteTimeToSeconds.
//
// See: https://developer.apple.com/documentation/coreservices/1551534-ucconvertcfabsolutetimetoseconds
func UCConvertCFAbsoluteTimeToSeconds(arg0 corefoundation.CFAbsoluteTime, arg1 uint32) int32 {
	if _uCConvertCFAbsoluteTimeToSeconds == nil {
		panic("coreservices: symbol UCConvertCFAbsoluteTimeToSeconds not loaded")
	}
	return _uCConvertCFAbsoluteTimeToSeconds(arg0, arg1)
}

var _uCConvertCFAbsoluteTimeToUTCDateTime func(arg0 corefoundation.CFAbsoluteTime, arg1 UTCDateTime) int32

// UCConvertCFAbsoluteTimeToUTCDateTime.
//
// See: https://developer.apple.com/documentation/coreservices/1551529-ucconvertcfabsolutetimetoutcdate
func UCConvertCFAbsoluteTimeToUTCDateTime(arg0 corefoundation.CFAbsoluteTime, arg1 UTCDateTime) int32 {
	if _uCConvertCFAbsoluteTimeToUTCDateTime == nil {
		panic("coreservices: symbol UCConvertCFAbsoluteTimeToUTCDateTime not loaded")
	}
	return _uCConvertCFAbsoluteTimeToUTCDateTime(arg0, arg1)
}

var _uCConvertLongDateTimeToCFAbsoluteTime func(arg0 LongDateTime, arg1 corefoundation.CFAbsoluteTime) int32

// UCConvertLongDateTimeToCFAbsoluteTime.
//
// See: https://developer.apple.com/documentation/coreservices/1551619-ucconvertlongdatetimetocfabsolut
func UCConvertLongDateTimeToCFAbsoluteTime(arg0 LongDateTime, arg1 corefoundation.CFAbsoluteTime) int32 {
	if _uCConvertLongDateTimeToCFAbsoluteTime == nil {
		panic("coreservices: symbol UCConvertLongDateTimeToCFAbsoluteTime not loaded")
	}
	return _uCConvertLongDateTimeToCFAbsoluteTime(arg0, arg1)
}

var _uCConvertSecondsToCFAbsoluteTime func(arg0 uint32, arg1 corefoundation.CFAbsoluteTime) int32

// UCConvertSecondsToCFAbsoluteTime.
//
// See: https://developer.apple.com/documentation/coreservices/1551512-ucconvertsecondstocfabsolutetime
func UCConvertSecondsToCFAbsoluteTime(arg0 uint32, arg1 corefoundation.CFAbsoluteTime) int32 {
	if _uCConvertSecondsToCFAbsoluteTime == nil {
		panic("coreservices: symbol UCConvertSecondsToCFAbsoluteTime not loaded")
	}
	return _uCConvertSecondsToCFAbsoluteTime(arg0, arg1)
}

var _uCConvertUTCDateTimeToCFAbsoluteTime func(arg0 UTCDateTime, arg1 corefoundation.CFAbsoluteTime) int32

// UCConvertUTCDateTimeToCFAbsoluteTime.
//
// See: https://developer.apple.com/documentation/coreservices/1551556-ucconvertutcdatetimetocfabsolute
func UCConvertUTCDateTimeToCFAbsoluteTime(arg0 UTCDateTime, arg1 corefoundation.CFAbsoluteTime) int32 {
	if _uCConvertUTCDateTimeToCFAbsoluteTime == nil {
		panic("coreservices: symbol UCConvertUTCDateTimeToCFAbsoluteTime not loaded")
	}
	return _uCConvertUTCDateTimeToCFAbsoluteTime(arg0, arg1)
}

var _uCCreateCollator func(locale LocaleRef, opVariant LocaleOperationVariant, options UCCollateOptions, collatorRef *CollatorRef) int32

// UCCreateCollator creates an object encapsulating locale and collation information, for the purpose of performing Unicode string comparison.
//
// See: https://developer.apple.com/documentation/coreservices/1390403-uccreatecollator
func UCCreateCollator(locale LocaleRef, opVariant LocaleOperationVariant, options UCCollateOptions, collatorRef *CollatorRef) int32 {
	if _uCCreateCollator == nil {
		panic("coreservices: symbol UCCreateCollator not loaded")
	}
	return _uCCreateCollator(locale, opVariant, options, collatorRef)
}

var _uCDisposeCollator func(collatorRef *CollatorRef) int32

// UCDisposeCollator disposes a collator object.
//
// See: https://developer.apple.com/documentation/coreservices/1390435-ucdisposecollator
func UCDisposeCollator(collatorRef *CollatorRef) int32 {
	if _uCDisposeCollator == nil {
		panic("coreservices: symbol UCDisposeCollator not loaded")
	}
	return _uCDisposeCollator(collatorRef)
}

var _uCGetCharProperty func(arg0 uint16, arg1 uint, arg2 UCCharPropertyType, arg3 UCCharPropertyValue) int32

// UCGetCharProperty obtains the value associated with a property type forthe specified [UniChar] characters.
//
// See: https://developer.apple.com/documentation/coreservices/1400224-ucgetcharproperty
func UCGetCharProperty(arg0 uint16, arg1 uint, arg2 UCCharPropertyType, arg3 UCCharPropertyValue) int32 {
	if _uCGetCharProperty == nil {
		panic("coreservices: symbol UCGetCharProperty not loaded")
	}
	return _uCGetCharProperty(arg0, arg1, arg2, arg3)
}

var _uCGetCollationKey func(collatorRef CollatorRef, textPtr *uint16, textLength unsafe.Pointer, maxKeySize unsafe.Pointer, actualKeySize unsafe.Pointer, collationKey *UCCollationValue) int32

// UCGetCollationKey uses locale-specific collation information to generate a collation key for a Unicode string.
//
// See: https://developer.apple.com/documentation/coreservices/1390468-ucgetcollationkey
func UCGetCollationKey(collatorRef CollatorRef, textPtr *uint16, textLength unsafe.Pointer, maxKeySize unsafe.Pointer, actualKeySize unsafe.Pointer, collationKey *UCCollationValue) int32 {
	if _uCGetCollationKey == nil {
		panic("coreservices: symbol UCGetCollationKey not loaded")
	}
	return _uCGetCollationKey(collatorRef, textPtr, textLength, maxKeySize, actualKeySize, collationKey)
}

var _uCGetUnicodeScalarValueForSurrogatePair func(arg0 uint16, arg1 uint16) unsafe.Pointer

// UCGetUnicodeScalarValueForSurrogatePair.
//
// See: https://developer.apple.com/documentation/coreservices/1399615-ucgetunicodescalarvalueforsurrog
func UCGetUnicodeScalarValueForSurrogatePair(arg0 uint16, arg1 uint16) unsafe.Pointer {
	if _uCGetUnicodeScalarValueForSurrogatePair == nil {
		panic("coreservices: symbol UCGetUnicodeScalarValueForSurrogatePair not loaded")
	}
	return _uCGetUnicodeScalarValueForSurrogatePair(arg0, arg1)
}

var _uCIsSurrogateHighCharacter func(arg0 uint16) bool

// UCIsSurrogateHighCharacter.
//
// See: https://developer.apple.com/documentation/coreservices/1400432-ucissurrogatehighcharacter
func UCIsSurrogateHighCharacter(arg0 uint16) bool {
	if _uCIsSurrogateHighCharacter == nil {
		panic("coreservices: symbol UCIsSurrogateHighCharacter not loaded")
	}
	return _uCIsSurrogateHighCharacter(arg0)
}

var _uCIsSurrogateLowCharacter func(arg0 uint16) bool

// UCIsSurrogateLowCharacter.
//
// See: https://developer.apple.com/documentation/coreservices/1400456-ucissurrogatelowcharacter
func UCIsSurrogateLowCharacter(arg0 uint16) bool {
	if _uCIsSurrogateLowCharacter == nil {
		panic("coreservices: symbol UCIsSurrogateLowCharacter not loaded")
	}
	return _uCIsSurrogateLowCharacter(arg0)
}

var _uCKeyTranslate func(keyLayoutPtr unsafe.Pointer, virtualKeyCode uint16, keyAction uint16, modifierKeyState uint32, keyboardType uint32, keyTranslateOptions uintptr, deadKeyState *uint32, maxStringLength unsafe.Pointer, actualStringLength unsafe.Pointer, unicodeString *uint16) int32

// UCKeyTranslate converts a combination of a virtual key code, a modifier key state, and a dead-key state into a string of one or more Unicode characters.
//
// See: https://developer.apple.com/documentation/coreservices/1390584-uckeytranslate
func UCKeyTranslate(keyLayoutPtr unsafe.Pointer, virtualKeyCode uint16, keyAction uint16, modifierKeyState uint32, keyboardType uint32, keyTranslateOptions uintptr, deadKeyState *uint32, maxStringLength unsafe.Pointer, actualStringLength unsafe.Pointer, unicodeString *uint16) int32 {
	if _uCKeyTranslate == nil {
		panic("coreservices: symbol UCKeyTranslate not loaded")
	}
	return _uCKeyTranslate(keyLayoutPtr, virtualKeyCode, keyAction, modifierKeyState, keyboardType, keyTranslateOptions, deadKeyState, maxStringLength, actualStringLength, unicodeString)
}

var _uCTypeSelectAddKeyToSelector func(inRef UCTypeSelectRef, inText corefoundation.CFStringRef, inEventTime unsafe.Pointer, updateFlag unsafe.Pointer) int32

// UCTypeSelectAddKeyToSelector.
//
// See: https://developer.apple.com/documentation/coreservices/1390517-uctypeselectaddkeytoselector
func UCTypeSelectAddKeyToSelector(inRef UCTypeSelectRef, inText corefoundation.CFStringRef, inEventTime unsafe.Pointer, updateFlag unsafe.Pointer) int32 {
	if _uCTypeSelectAddKeyToSelector == nil {
		panic("coreservices: symbol UCTypeSelectAddKeyToSelector not loaded")
	}
	return _uCTypeSelectAddKeyToSelector(inRef, inText, inEventTime, updateFlag)
}

var _uCTypeSelectCompare func(ref UCTypeSelectRef, inText corefoundation.CFStringRef, result *UCTypeSelectCompareResult) int32

// UCTypeSelectCompare.
//
// See: https://developer.apple.com/documentation/coreservices/1390474-uctypeselectcompare
func UCTypeSelectCompare(ref UCTypeSelectRef, inText corefoundation.CFStringRef, result *UCTypeSelectCompareResult) int32 {
	if _uCTypeSelectCompare == nil {
		panic("coreservices: symbol UCTypeSelectCompare not loaded")
	}
	return _uCTypeSelectCompare(ref, inText, result)
}

var _uCTypeSelectCreateSelector func(locale LocaleRef, opVariant LocaleOperationVariant, options UCCollateOptions, newSelector *UCTypeSelectRef) int32

// UCTypeSelectCreateSelector.
//
// See: https://developer.apple.com/documentation/coreservices/1390445-uctypeselectcreateselector
func UCTypeSelectCreateSelector(locale LocaleRef, opVariant LocaleOperationVariant, options UCCollateOptions, newSelector *UCTypeSelectRef) int32 {
	if _uCTypeSelectCreateSelector == nil {
		panic("coreservices: symbol UCTypeSelectCreateSelector not loaded")
	}
	return _uCTypeSelectCreateSelector(locale, opVariant, options, newSelector)
}

var _uCTypeSelectFindItem func(ref UCTypeSelectRef, listSize uint32, listDataPtr unsafe.Pointer, refcon uintptr, userUPP IndexToUCStringUPP, closestItem *uint32) int32

// UCTypeSelectFindItem.
//
// See: https://developer.apple.com/documentation/coreservices/1390368-uctypeselectfinditem
func UCTypeSelectFindItem(ref UCTypeSelectRef, listSize uint32, listDataPtr unsafe.Pointer, refcon uintptr, userUPP IndexToUCStringUPP, closestItem *uint32) int32 {
	if _uCTypeSelectFindItem == nil {
		panic("coreservices: symbol UCTypeSelectFindItem not loaded")
	}
	return _uCTypeSelectFindItem(ref, listSize, listDataPtr, refcon, userUPP, closestItem)
}

var _uCTypeSelectFlushSelectorData func(ref UCTypeSelectRef) int32

// UCTypeSelectFlushSelectorData.
//
// See: https://developer.apple.com/documentation/coreservices/1390367-uctypeselectflushselectordata
func UCTypeSelectFlushSelectorData(ref UCTypeSelectRef) int32 {
	if _uCTypeSelectFlushSelectorData == nil {
		panic("coreservices: symbol UCTypeSelectFlushSelectorData not loaded")
	}
	return _uCTypeSelectFlushSelectorData(ref)
}

var _uCTypeSelectReleaseSelector func(ref *UCTypeSelectRef) int32

// UCTypeSelectReleaseSelector.
//
// See: https://developer.apple.com/documentation/coreservices/1390644-uctypeselectreleaseselector
func UCTypeSelectReleaseSelector(ref *UCTypeSelectRef) int32 {
	if _uCTypeSelectReleaseSelector == nil {
		panic("coreservices: symbol UCTypeSelectReleaseSelector not loaded")
	}
	return _uCTypeSelectReleaseSelector(ref)
}

var _uCTypeSelectWalkList func(ref UCTypeSelectRef, currSelect corefoundation.CFStringRef, direction UCTSWalkDirection, listSize uint32, listDataPtr unsafe.Pointer, refcon uintptr, userUPP IndexToUCStringUPP, closestItem *uint32) int32

// UCTypeSelectWalkList.
//
// See: https://developer.apple.com/documentation/coreservices/1390442-uctypeselectwalklist
func UCTypeSelectWalkList(ref UCTypeSelectRef, currSelect corefoundation.CFStringRef, direction UCTSWalkDirection, listSize uint32, listDataPtr unsafe.Pointer, refcon uintptr, userUPP IndexToUCStringUPP, closestItem *uint32) int32 {
	if _uCTypeSelectWalkList == nil {
		panic("coreservices: symbol UCTypeSelectWalkList not loaded")
	}
	return _uCTypeSelectWalkList(ref, currSelect, direction, listSize, listDataPtr, refcon, userUPP, closestItem)
}

var _uCTypeSelectWouldResetBuffer func(inRef UCTypeSelectRef, inText corefoundation.CFStringRef, inEventTime unsafe.Pointer) bool

// UCTypeSelectWouldResetBuffer.
//
// See: https://developer.apple.com/documentation/coreservices/1390538-uctypeselectwouldresetbuffer
func UCTypeSelectWouldResetBuffer(inRef UCTypeSelectRef, inText corefoundation.CFStringRef, inEventTime unsafe.Pointer) bool {
	if _uCTypeSelectWouldResetBuffer == nil {
		panic("coreservices: symbol UCTypeSelectWouldResetBuffer not loaded")
	}
	return _uCTypeSelectWouldResetBuffer(inRef, inText, inEventTime)
}

var _uInt64ToLongDouble func(arg0 uint64) unsafe.Pointer

// UInt64ToLongDouble.
//
// See: https://developer.apple.com/documentation/coreservices/1536335-uint64tolongdouble
func UInt64ToLongDouble(arg0 uint64) unsafe.Pointer {
	if _uInt64ToLongDouble == nil {
		panic("coreservices: symbol UInt64ToLongDouble not loaded")
	}
	return _uInt64ToLongDouble(arg0)
}

var _uInt64ToSInt64 func(arg0 uint64) int64

// UInt64ToSInt64.
//
// See: https://developer.apple.com/documentation/coreservices/1536266-uint64tosint64
func UInt64ToSInt64(arg0 uint64) int64 {
	if _uInt64ToSInt64 == nil {
		panic("coreservices: symbol UInt64ToSInt64 not loaded")
	}
	return _uInt64ToSInt64(arg0)
}

var _uInt64ToUnsignedWide func(arg0 uint64) uint64

// UInt64ToUnsignedWide.
//
// See: https://developer.apple.com/documentation/coreservices/1536362-uint64tounsignedwide
func UInt64ToUnsignedWide(arg0 uint64) uint64 {
	if _uInt64ToUnsignedWide == nil {
		panic("coreservices: symbol UInt64ToUnsignedWide not loaded")
	}
	return _uInt64ToUnsignedWide(arg0)
}

var _uTCreateStringForOSType func(inOSType uint32) corefoundation.CFStringRef

// UTCreateStringForOSType encodes an [OSType] into a string suitable for use as a tag argument.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1442804-utcreatestringforostype
func UTCreateStringForOSType(inOSType uint32) corefoundation.CFStringRef {
	if _uTCreateStringForOSType == nil {
		panic("coreservices: symbol UTCreateStringForOSType not loaded")
	}
	return _uTCreateStringForOSType(inOSType)
}

var _uTGetOSTypeFromString func(inString corefoundation.CFStringRef) uint32

// UTGetOSTypeFromString decodes a tag string into an OSType.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1450472-utgetostypefromstring
func UTGetOSTypeFromString(inString corefoundation.CFStringRef) uint32 {
	if _uTGetOSTypeFromString == nil {
		panic("coreservices: symbol UTGetOSTypeFromString not loaded")
	}
	return _uTGetOSTypeFromString(inString)
}

var _uTTypeConformsTo func(inUTI corefoundation.CFStringRef, inConformsToUTI corefoundation.CFStringRef) bool

// UTTypeConformsTo returns whether a uniform type identifier conforms to another uniform type identifier.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1444079-uttypeconformsto
func UTTypeConformsTo(inUTI corefoundation.CFStringRef, inConformsToUTI corefoundation.CFStringRef) bool {
	if _uTTypeConformsTo == nil {
		panic("coreservices: symbol UTTypeConformsTo not loaded")
	}
	return _uTTypeConformsTo(inUTI, inConformsToUTI)
}

var _uTTypeCopyAllTagsWithClass func(inUTI corefoundation.CFStringRef, inTagClass corefoundation.CFStringRef) corefoundation.CFArrayRef

// UTTypeCopyAllTagsWithClass.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1448473-uttypecopyalltagswithclass
func UTTypeCopyAllTagsWithClass(inUTI corefoundation.CFStringRef, inTagClass corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _uTTypeCopyAllTagsWithClass == nil {
		panic("coreservices: symbol UTTypeCopyAllTagsWithClass not loaded")
	}
	return _uTTypeCopyAllTagsWithClass(inUTI, inTagClass)
}

var _uTTypeCopyDeclaration func(inUTI corefoundation.CFStringRef) corefoundation.CFDictionaryRef

// UTTypeCopyDeclaration returns a uniform type’s declaration.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1442505-uttypecopydeclaration
func UTTypeCopyDeclaration(inUTI corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	if _uTTypeCopyDeclaration == nil {
		panic("coreservices: symbol UTTypeCopyDeclaration not loaded")
	}
	return _uTTypeCopyDeclaration(inUTI)
}

var _uTTypeCopyDeclaringBundleURL func(inUTI corefoundation.CFStringRef) corefoundation.CFURLRef

// UTTypeCopyDeclaringBundleURL returns the location of a bundle containing the declaration for a type.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/coreservices/1447781-uttypecopydeclaringbundleurl
func UTTypeCopyDeclaringBundleURL(inUTI corefoundation.CFStringRef) corefoundation.CFURLRef {
	if _uTTypeCopyDeclaringBundleURL == nil {
		panic("coreservices: symbol UTTypeCopyDeclaringBundleURL not loaded")
	}
	return _uTTypeCopyDeclaringBundleURL(inUTI)
}

var _uTTypeCopyDescription func(inUTI corefoundation.CFStringRef) corefoundation.CFStringRef

// UTTypeCopyDescription returns the localized, user-readable type description string associated with a uniform type identifier.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1448514-uttypecopydescription
func UTTypeCopyDescription(inUTI corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _uTTypeCopyDescription == nil {
		panic("coreservices: symbol UTTypeCopyDescription not loaded")
	}
	return _uTTypeCopyDescription(inUTI)
}

var _uTTypeCopyPreferredTagWithClass func(inUTI corefoundation.CFStringRef, inTagClass corefoundation.CFStringRef) corefoundation.CFStringRef

// UTTypeCopyPreferredTagWithClass translates a uniform type identifier to a list of tags in a different type classification method.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1442744-uttypecopypreferredtagwithclass
func UTTypeCopyPreferredTagWithClass(inUTI corefoundation.CFStringRef, inTagClass corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _uTTypeCopyPreferredTagWithClass == nil {
		panic("coreservices: symbol UTTypeCopyPreferredTagWithClass not loaded")
	}
	return _uTTypeCopyPreferredTagWithClass(inUTI, inTagClass)
}

var _uTTypeCreateAllIdentifiersForTag func(inTagClass corefoundation.CFStringRef, inTag corefoundation.CFStringRef, inConformingToUTI corefoundation.CFStringRef) corefoundation.CFArrayRef

// UTTypeCreateAllIdentifiersForTag creates an array of all uniform type identifiers for the type indicated by the specified tag.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1447261-uttypecreateallidentifiersfortag
func UTTypeCreateAllIdentifiersForTag(inTagClass corefoundation.CFStringRef, inTag corefoundation.CFStringRef, inConformingToUTI corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _uTTypeCreateAllIdentifiersForTag == nil {
		panic("coreservices: symbol UTTypeCreateAllIdentifiersForTag not loaded")
	}
	return _uTTypeCreateAllIdentifiersForTag(inTagClass, inTag, inConformingToUTI)
}

var _uTTypeCreatePreferredIdentifierForTag func(inTagClass corefoundation.CFStringRef, inTag corefoundation.CFStringRef, inConformingToUTI corefoundation.CFStringRef) corefoundation.CFStringRef

// UTTypeCreatePreferredIdentifierForTag creates a uniform type identifier for the type indicated by the specified tag.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1448939-uttypecreatepreferredidentifierf
func UTTypeCreatePreferredIdentifierForTag(inTagClass corefoundation.CFStringRef, inTag corefoundation.CFStringRef, inConformingToUTI corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _uTTypeCreatePreferredIdentifierForTag == nil {
		panic("coreservices: symbol UTTypeCreatePreferredIdentifierForTag not loaded")
	}
	return _uTTypeCreatePreferredIdentifierForTag(inTagClass, inTag, inConformingToUTI)
}

var _uTTypeEqual func(inUTI1 corefoundation.CFStringRef, inUTI2 corefoundation.CFStringRef) bool

// UTTypeEqual returns whether two uniform type identifiers are equal.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1447783-uttypeequal
func UTTypeEqual(inUTI1 corefoundation.CFStringRef, inUTI2 corefoundation.CFStringRef) bool {
	if _uTTypeEqual == nil {
		panic("coreservices: symbol UTTypeEqual not loaded")
	}
	return _uTTypeEqual(inUTI1, inUTI2)
}

var _uTTypeIsDeclared func(inUTI corefoundation.CFStringRef) bool

// UTTypeIsDeclared.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1450352-uttypeisdeclared
func UTTypeIsDeclared(inUTI corefoundation.CFStringRef) bool {
	if _uTTypeIsDeclared == nil {
		panic("coreservices: symbol UTTypeIsDeclared not loaded")
	}
	return _uTTypeIsDeclared(inUTI)
}

var _uTTypeIsDynamic func(inUTI corefoundation.CFStringRef) bool

// UTTypeIsDynamic.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/coreservices/1442980-uttypeisdynamic
func UTTypeIsDynamic(inUTI corefoundation.CFStringRef) bool {
	if _uTTypeIsDynamic == nil {
		panic("coreservices: symbol UTTypeIsDynamic not loaded")
	}
	return _uTTypeIsDynamic(inUTI)
}

var _uncaptureComponent func(arg0 Component) int16

// UncaptureComponent allows your component to uncapture a previously captured component.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516646-uncapturecomponent
func UncaptureComponent(arg0 Component) int16 {
	if _uncaptureComponent == nil {
		panic("coreservices: symbol UncaptureComponent not loaded")
	}
	return _uncaptureComponent(arg0)
}

var _unflattenCollection func(arg0 Collection, arg1 CollectionFlattenUPP) int16

// UnflattenCollection.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551348-unflattencollection
func UnflattenCollection(arg0 Collection, arg1 CollectionFlattenUPP) int16 {
	if _unflattenCollection == nil {
		panic("coreservices: symbol UnflattenCollection not loaded")
	}
	return _unflattenCollection(arg0, arg1)
}

var _unflattenCollectionFromHdl func(arg0 Collection, arg1 uintptr) int16

// UnflattenCollectionFromHdl.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1551400-unflattencollectionfromhdl
func UnflattenCollectionFromHdl(arg0 Collection, arg1 uintptr) int16 {
	if _unflattenCollectionFromHdl == nil {
		panic("coreservices: symbol UnflattenCollectionFromHdl not loaded")
	}
	return _unflattenCollectionFromHdl(arg0, arg1)
}

var _unique1ID func(arg0 uint32) ResID

// Unique1ID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529314-unique1id
func Unique1ID(arg0 uint32) ResID {
	if _unique1ID == nil {
		panic("coreservices: symbol Unique1ID not loaded")
	}
	return _unique1ID(arg0)
}

var _uniqueID func(arg0 uint32) ResID

// UniqueID.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529255-uniqueid
func UniqueID(arg0 uint32) ResID {
	if _uniqueID == nil {
		panic("coreservices: symbol UniqueID not loaded")
	}
	return _uniqueID(arg0)
}

var _unregisterComponent func(arg0 Component) int16

// UnregisterComponent removes a component from the Component Manager’s registration list.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1516645-unregistercomponent
func UnregisterComponent(arg0 Component) int16 {
	if _unregisterComponent == nil {
		panic("coreservices: symbol UnregisterComponent not loaded")
	}
	return _unregisterComponent(arg0)
}

var _unregisterIconRef func(creator uint32, iconType uint32) int16

// UnregisterIconRef.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/coreservices/1444660-unregistericonref
func UnregisterIconRef(creator uint32, iconType uint32) int16 {
	if _unregisterIconRef == nil {
		panic("coreservices: symbol UnregisterIconRef not loaded")
	}
	return _unregisterIconRef(creator, iconType)
}

var _unsignedFixedMulDiv func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer) unsafe.Pointer

// UnsignedFixedMulDiv.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409276-unsignedfixedmuldiv
func UnsignedFixedMulDiv(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer) unsafe.Pointer {
	if _unsignedFixedMulDiv == nil {
		panic("coreservices: symbol UnsignedFixedMulDiv not loaded")
	}
	return _unsignedFixedMulDiv(arg0, arg1, arg2)
}

var _unsignedWideToUInt64 func(arg0 uint64) uint64

// UnsignedWideToUInt64.
//
// See: https://developer.apple.com/documentation/coreservices/1536357-unsignedwidetouint64
func UnsignedWideToUInt64(arg0 uint64) uint64 {
	if _unsignedWideToUInt64 == nil {
		panic("coreservices: symbol UnsignedWideToUInt64 not loaded")
	}
	return _unsignedWideToUInt64(arg0)
}

var _upTime func() unsafe.Pointer

// UpTime.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1501237-uptime
func UpTime() unsafe.Pointer {
	if _upTime == nil {
		panic("coreservices: symbol UpTime not loaded")
	}
	return _upTime()
}

var _updateIconRef func(theIconRef uintptr) int16

// UpdateIconRef.
//
// Deprecated: Deprecated since macOS 10.15.
//
// See: https://developer.apple.com/documentation/coreservices/1445921-updateiconref
func UpdateIconRef(theIconRef uintptr) int16 {
	if _updateIconRef == nil {
		panic("coreservices: symbol UpdateIconRef not loaded")
	}
	return _updateIconRef(theIconRef)
}

var _updateResFile func(arg0 ResFileRefNum)

// UpdateResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529379-updateresfile
func UpdateResFile(arg0 ResFileRefNum) {
	if _updateResFile == nil {
		panic("coreservices: symbol UpdateResFile not loaded")
	}
	_updateResFile(arg0)
}

var _updateSystemActivity func(arg0 uint8) int16

// UpdateSystemActivity.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1427505-updatesystemactivity
func UpdateSystemActivity(arg0 uint8) int16 {
	if _updateSystemActivity == nil {
		panic("coreservices: symbol UpdateSystemActivity not loaded")
	}
	return _updateSystemActivity(arg0)
}

var _upgradeScriptInfoToTextEncoding func(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 unsafe.Pointer, arg4 TextEncoding) int32

// UpgradeScriptInfoToTextEncoding converts any combination of a Mac OS script code, a languagecode, a region code, and a font name to a text encoding.
//
// See: https://developer.apple.com/documentation/coreservices/1399625-upgradescriptinfototextencoding
func UpgradeScriptInfoToTextEncoding(arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 unsafe.Pointer, arg4 TextEncoding) int32 {
	if _upgradeScriptInfoToTextEncoding == nil {
		panic("coreservices: symbol UpgradeScriptInfoToTextEncoding not loaded")
	}
	return _upgradeScriptInfoToTextEncoding(arg0, arg1, arg2, arg3, arg4)
}

var _useResFile func(arg0 ResFileRefNum)

// UseResFile.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529268-useresfile
func UseResFile(arg0 ResFileRefNum) {
	if _useResFile == nil {
		panic("coreservices: symbol UseResFile not loaded")
	}
	_useResFile(arg0)
}

var _wSGetCFTypeIDFromWSTypeID func(arg0 WSTypeID) uint

// WSGetCFTypeIDFromWSTypeID gets the CFType associated with a given WSType
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1508451-wsgetcftypeidfromwstypeid
func WSGetCFTypeIDFromWSTypeID(arg0 WSTypeID) uint {
	if _wSGetCFTypeIDFromWSTypeID == nil {
		panic("coreservices: symbol WSGetCFTypeIDFromWSTypeID not loaded")
	}
	return _wSGetCFTypeIDFromWSTypeID(arg0)
}

var _wSGetWSTypeIDFromCFType func(arg0 corefoundation.CFTypeRef) WSTypeID

// WSGetWSTypeIDFromCFType returns the [WSTypeID] associated with a given [CFTypeRef].
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1508442-wsgetwstypeidfromcftype
func WSGetWSTypeIDFromCFType(arg0 corefoundation.CFTypeRef) WSTypeID {
	if _wSGetWSTypeIDFromCFType == nil {
		panic("coreservices: symbol WSGetWSTypeIDFromCFType not loaded")
	}
	return _wSGetWSTypeIDFromCFType(arg0)
}

var _wSMethodInvocationAddDeserializationOverride func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 unsafe.Pointer, arg4 WSClientContext)

// WSMethodInvocationAddDeserializationOverride specifies a callback to be made when parsing the XML in a method response.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576436-wsmethodinvocationadddeserializa
func WSMethodInvocationAddDeserializationOverride(arg0 WSMethodInvocationRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 unsafe.Pointer, arg4 WSClientContext) {
	if _wSMethodInvocationAddDeserializationOverride == nil {
		panic("coreservices: symbol WSMethodInvocationAddDeserializationOverride not loaded")
	}
	_wSMethodInvocationAddDeserializationOverride(arg0, arg1, arg2, arg3, arg4)
}

var _wSMethodInvocationAddSerializationOverride func(arg0 WSMethodInvocationRef, arg1 uint, arg2 unsafe.Pointer, arg3 WSClientContext)

// WSMethodInvocationAddSerializationOverride specifies a callback to be made when creating the XML for an method invocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576435-wsmethodinvocationaddserializati
func WSMethodInvocationAddSerializationOverride(arg0 WSMethodInvocationRef, arg1 uint, arg2 unsafe.Pointer, arg3 WSClientContext) {
	if _wSMethodInvocationAddSerializationOverride == nil {
		panic("coreservices: symbol WSMethodInvocationAddSerializationOverride not loaded")
	}
	_wSMethodInvocationAddSerializationOverride(arg0, arg1, arg2, arg3)
}

var _wSMethodInvocationCopyParameters func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFArrayRef) corefoundation.CFDictionaryRef

// WSMethodInvocationCopyParameters creates a copy of the parameters dictionary and sets the order in an array.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576399-wsmethodinvocationcopyparameters
func WSMethodInvocationCopyParameters(arg0 WSMethodInvocationRef, arg1 corefoundation.CFArrayRef) corefoundation.CFDictionaryRef {
	if _wSMethodInvocationCopyParameters == nil {
		panic("coreservices: symbol WSMethodInvocationCopyParameters not loaded")
	}
	return _wSMethodInvocationCopyParameters(arg0, arg1)
}

var _wSMethodInvocationCopyProperty func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// WSMethodInvocationCopyProperty creates a copy of a named property of the invocation reference.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576432-wsmethodinvocationcopyproperty
func WSMethodInvocationCopyProperty(arg0 WSMethodInvocationRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _wSMethodInvocationCopyProperty == nil {
		panic("coreservices: symbol WSMethodInvocationCopyProperty not loaded")
	}
	return _wSMethodInvocationCopyProperty(arg0, arg1)
}

var _wSMethodInvocationCopySerialization func(arg0 WSMethodInvocationRef) corefoundation.CFDataRef

// WSMethodInvocationCopySerialization creates an XML serialization of a method invocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576441-wsmethodinvocationcopyserializat
func WSMethodInvocationCopySerialization(arg0 WSMethodInvocationRef) corefoundation.CFDataRef {
	if _wSMethodInvocationCopySerialization == nil {
		panic("coreservices: symbol WSMethodInvocationCopySerialization not loaded")
	}
	return _wSMethodInvocationCopySerialization(arg0)
}

var _wSMethodInvocationCreate func(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef) WSMethodInvocationRef

// WSMethodInvocationCreate creates a reference to a method invocation, containing the URL of the service, the operation name, and the protocol.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576406-wsmethodinvocationcreate
func WSMethodInvocationCreate(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef) WSMethodInvocationRef {
	if _wSMethodInvocationCreate == nil {
		panic("coreservices: symbol WSMethodInvocationCreate not loaded")
	}
	return _wSMethodInvocationCreate(arg0, arg1, arg2)
}

var _wSMethodInvocationCreateFromSerialization func(arg0 corefoundation.CFDataRef) WSMethodInvocationRef

// WSMethodInvocationCreateFromSerialization creates a method invocation object from an XML serialization.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576400-wsmethodinvocationcreatefromseri
func WSMethodInvocationCreateFromSerialization(arg0 corefoundation.CFDataRef) WSMethodInvocationRef {
	if _wSMethodInvocationCreateFromSerialization == nil {
		panic("coreservices: symbol WSMethodInvocationCreateFromSerialization not loaded")
	}
	return _wSMethodInvocationCreateFromSerialization(arg0)
}

var _wSMethodInvocationGetTypeID func() uint

// WSMethodInvocationGetTypeID returns the type ID of the current method invocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576434-wsmethodinvocationgettypeid
func WSMethodInvocationGetTypeID() uint {
	if _wSMethodInvocationGetTypeID == nil {
		panic("coreservices: symbol WSMethodInvocationGetTypeID not loaded")
	}
	return _wSMethodInvocationGetTypeID()
}

var _wSMethodInvocationInvoke func(arg0 WSMethodInvocationRef) corefoundation.CFDictionaryRef

// WSMethodInvocationInvoke invokes an web services operation synchronously.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576428-wsmethodinvocationinvoke
func WSMethodInvocationInvoke(arg0 WSMethodInvocationRef) corefoundation.CFDictionaryRef {
	if _wSMethodInvocationInvoke == nil {
		panic("coreservices: symbol WSMethodInvocationInvoke not loaded")
	}
	return _wSMethodInvocationInvoke(arg0)
}

var _wSMethodInvocationScheduleWithRunLoop func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// WSMethodInvocationScheduleWithRunLoop schedule a method invocation for asynchronous execution on a run loop.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576408-wsmethodinvocationschedulewithru
func WSMethodInvocationScheduleWithRunLoop(arg0 WSMethodInvocationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _wSMethodInvocationScheduleWithRunLoop == nil {
		panic("coreservices: symbol WSMethodInvocationScheduleWithRunLoop not loaded")
	}
	_wSMethodInvocationScheduleWithRunLoop(arg0, arg1, arg2)
}

var _wSMethodInvocationSetCallBack func(arg0 WSMethodInvocationRef, arg1 unsafe.Pointer, arg2 WSClientContext)

// WSMethodInvocationSetCallBack set a callback to handle the response to an asynchronous method invocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576439-wsmethodinvocationsetcallback
func WSMethodInvocationSetCallBack(arg0 WSMethodInvocationRef, arg1 unsafe.Pointer, arg2 WSClientContext) {
	if _wSMethodInvocationSetCallBack == nil {
		panic("coreservices: symbol WSMethodInvocationSetCallBack not loaded")
	}
	_wSMethodInvocationSetCallBack(arg0, arg1, arg2)
}

var _wSMethodInvocationSetParameters func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFDictionaryRef, arg2 corefoundation.CFArrayRef)

// WSMethodInvocationSetParameters set the parameter names, types, and order for a method invocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576420-wsmethodinvocationsetparameters
func WSMethodInvocationSetParameters(arg0 WSMethodInvocationRef, arg1 corefoundation.CFDictionaryRef, arg2 corefoundation.CFArrayRef) {
	if _wSMethodInvocationSetParameters == nil {
		panic("coreservices: symbol WSMethodInvocationSetParameters not loaded")
	}
	_wSMethodInvocationSetParameters(arg0, arg1, arg2)
}

var _wSMethodInvocationSetProperty func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef)

// WSMethodInvocationSetProperty sets a named property of the method invocation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576407-wsmethodinvocationsetproperty
func WSMethodInvocationSetProperty(arg0 WSMethodInvocationRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) {
	if _wSMethodInvocationSetProperty == nil {
		panic("coreservices: symbol WSMethodInvocationSetProperty not loaded")
	}
	_wSMethodInvocationSetProperty(arg0, arg1, arg2)
}

var _wSMethodInvocationUnscheduleFromRunLoop func(arg0 WSMethodInvocationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// WSMethodInvocationUnscheduleFromRunLoop unschedules a method invocation from a run loop.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576409-wsmethodinvocationunschedulefrom
func WSMethodInvocationUnscheduleFromRunLoop(arg0 WSMethodInvocationRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _wSMethodInvocationUnscheduleFromRunLoop == nil {
		panic("coreservices: symbol WSMethodInvocationUnscheduleFromRunLoop not loaded")
	}
	_wSMethodInvocationUnscheduleFromRunLoop(arg0, arg1, arg2)
}

var _wSMethodResultIsFault func(arg0 corefoundation.CFDictionaryRef) bool

// WSMethodResultIsFault tests a method result dictionary for a fault condition.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1576401-wsmethodresultisfault
func WSMethodResultIsFault(arg0 corefoundation.CFDictionaryRef) bool {
	if _wSMethodResultIsFault == nil {
		panic("coreservices: symbol WSMethodResultIsFault not loaded")
	}
	return _wSMethodResultIsFault(arg0)
}

var _wSProtocolHandlerCopyFaultDocument func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFDictionaryRef, arg2 corefoundation.CFDictionaryRef) corefoundation.CFDataRef

// WSProtocolHandlerCopyFaultDocument creates a Fault XML response for a given WSProtocolHandler and fault details dictionary.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423449-wsprotocolhandlercopyfaultdocume
func WSProtocolHandlerCopyFaultDocument(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFDictionaryRef, arg2 corefoundation.CFDictionaryRef) corefoundation.CFDataRef {
	if _wSProtocolHandlerCopyFaultDocument == nil {
		panic("coreservices: symbol WSProtocolHandlerCopyFaultDocument not loaded")
	}
	return _wSProtocolHandlerCopyFaultDocument(arg0, arg1, arg2)
}

var _wSProtocolHandlerCopyProperty func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// WSProtocolHandlerCopyProperty returns a copy of a property from a protocol handler reference.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423436-wsprotocolhandlercopyproperty
func WSProtocolHandlerCopyProperty(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _wSProtocolHandlerCopyProperty == nil {
		panic("coreservices: symbol WSProtocolHandlerCopyProperty not loaded")
	}
	return _wSProtocolHandlerCopyProperty(arg0, arg1)
}

var _wSProtocolHandlerCopyReplyDictionary func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFDataRef) corefoundation.CFDictionaryRef

// WSProtocolHandlerCopyReplyDictionary parses an incoming XML document as if it were the reply of a method.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423441-wsprotocolhandlercopyreplydictio
func WSProtocolHandlerCopyReplyDictionary(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFDataRef) corefoundation.CFDictionaryRef {
	if _wSProtocolHandlerCopyReplyDictionary == nil {
		panic("coreservices: symbol WSProtocolHandlerCopyReplyDictionary not loaded")
	}
	return _wSProtocolHandlerCopyReplyDictionary(arg0, arg1, arg2)
}

var _wSProtocolHandlerCopyReplyDocument func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFDictionaryRef, arg2 corefoundation.CFTypeRef) corefoundation.CFDataRef

// WSProtocolHandlerCopyReplyDocument creates a Reply XML document for a given WS ProtocolHandler and context dictionary.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423440-wsprotocolhandlercopyreplydocume
func WSProtocolHandlerCopyReplyDocument(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFDictionaryRef, arg2 corefoundation.CFTypeRef) corefoundation.CFDataRef {
	if _wSProtocolHandlerCopyReplyDocument == nil {
		panic("coreservices: symbol WSProtocolHandlerCopyReplyDocument not loaded")
	}
	return _wSProtocolHandlerCopyReplyDocument(arg0, arg1, arg2)
}

var _wSProtocolHandlerCopyRequestDictionary func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFDataRef) corefoundation.CFDictionaryRef

// WSProtocolHandlerCopyRequestDictionary parses an incoming XML document for the method name and parameters.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423451-wsprotocolhandlercopyrequestdict
func WSProtocolHandlerCopyRequestDictionary(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFDataRef) corefoundation.CFDictionaryRef {
	if _wSProtocolHandlerCopyRequestDictionary == nil {
		panic("coreservices: symbol WSProtocolHandlerCopyRequestDictionary not loaded")
	}
	return _wSProtocolHandlerCopyRequestDictionary(arg0, arg1)
}

var _wSProtocolHandlerCopyRequestDocument func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFDictionaryRef, arg3 corefoundation.CFArrayRef, arg4 corefoundation.CFDictionaryRef) corefoundation.CFDataRef

// WSProtocolHandlerCopyRequestDocument creates an XML request for a given [WSProtocolHandler] and parameter list.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423437-wsprotocolhandlercopyrequestdocu
func WSProtocolHandlerCopyRequestDocument(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFDictionaryRef, arg3 corefoundation.CFArrayRef, arg4 corefoundation.CFDictionaryRef) corefoundation.CFDataRef {
	if _wSProtocolHandlerCopyRequestDocument == nil {
		panic("coreservices: symbol WSProtocolHandlerCopyRequestDocument not loaded")
	}
	return _wSProtocolHandlerCopyRequestDocument(arg0, arg1, arg2, arg3, arg4)
}

var _wSProtocolHandlerCreate func(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFStringRef) WSProtocolHandlerRef

// WSProtocolHandlerCreate creates a [WSProtocolHandlerRef] for use in translating an XML document.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423434-wsprotocolhandlercreate
func WSProtocolHandlerCreate(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFStringRef) WSProtocolHandlerRef {
	if _wSProtocolHandlerCreate == nil {
		panic("coreservices: symbol WSProtocolHandlerCreate not loaded")
	}
	return _wSProtocolHandlerCreate(arg0, arg1)
}

var _wSProtocolHandlerGetTypeID func() uint

// WSProtocolHandlerGetTypeID returns a [CFTypeID] for the current [WSProtocolHandlerRef].
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423455-wsprotocolhandlergettypeid
func WSProtocolHandlerGetTypeID() uint {
	if _wSProtocolHandlerGetTypeID == nil {
		panic("coreservices: symbol WSProtocolHandlerGetTypeID not loaded")
	}
	return _wSProtocolHandlerGetTypeID()
}

var _wSProtocolHandlerSetDeserializationOverride func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 unsafe.Pointer, arg4 WSClientContext)

// WSProtocolHandlerSetDeserializationOverride specifies a callback to be made when parsing an XML method response.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423444-wsprotocolhandlersetdeserializat
func WSProtocolHandlerSetDeserializationOverride(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 unsafe.Pointer, arg4 WSClientContext) {
	if _wSProtocolHandlerSetDeserializationOverride == nil {
		panic("coreservices: symbol WSProtocolHandlerSetDeserializationOverride not loaded")
	}
	_wSProtocolHandlerSetDeserializationOverride(arg0, arg1, arg2, arg3, arg4)
}

var _wSProtocolHandlerSetProperty func(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef)

// WSProtocolHandlerSetProperty sets a property in a specified protocol handler.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423442-wsprotocolhandlersetproperty
func WSProtocolHandlerSetProperty(arg0 WSProtocolHandlerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) {
	if _wSProtocolHandlerSetProperty == nil {
		panic("coreservices: symbol WSProtocolHandlerSetProperty not loaded")
	}
	_wSProtocolHandlerSetProperty(arg0, arg1, arg2)
}

var _wSProtocolHandlerSetSerializationOverride func(arg0 WSProtocolHandlerRef, arg1 uint, arg2 unsafe.Pointer, arg3 WSClientContext)

// WSProtocolHandlerSetSerializationOverride specifies a callback which will be called to produce the XML that represents the serialization of a given type ref.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1423438-wsprotocolhandlersetserializatio
func WSProtocolHandlerSetSerializationOverride(arg0 WSProtocolHandlerRef, arg1 uint, arg2 unsafe.Pointer, arg3 WSClientContext) {
	if _wSProtocolHandlerSetSerializationOverride == nil {
		panic("coreservices: symbol WSProtocolHandlerSetSerializationOverride not loaded")
	}
	_wSProtocolHandlerSetSerializationOverride(arg0, arg1, arg2, arg3)
}

var _wideAdd func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) objectivec.IObject

// WideAdd.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409245-wideadd
func WideAdd(arg0 unsafe.Pointer, arg1 unsafe.Pointer) objectivec.IObject {
	if _wideAdd == nil {
		panic("coreservices: symbol WideAdd not loaded")
	}
	return _wideAdd(arg0, arg1)
}

var _wideBitShift func(arg0 unsafe.Pointer, arg1 int32) objectivec.IObject

// WideBitShift.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409247-widebitshift
func WideBitShift(arg0 unsafe.Pointer, arg1 int32) objectivec.IObject {
	if _wideBitShift == nil {
		panic("coreservices: symbol WideBitShift not loaded")
	}
	return _wideBitShift(arg0, arg1)
}

var _wideCompare func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16

// WideCompare.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409264-widecompare
func WideCompare(arg0 unsafe.Pointer, arg1 unsafe.Pointer) int16 {
	if _wideCompare == nil {
		panic("coreservices: symbol WideCompare not loaded")
	}
	return _wideCompare(arg0, arg1)
}

var _wideDivide func(arg0 unsafe.Pointer, arg1 int32, arg2 int32) int32

// WideDivide.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409233-widedivide
func WideDivide(arg0 unsafe.Pointer, arg1 int32, arg2 int32) int32 {
	if _wideDivide == nil {
		panic("coreservices: symbol WideDivide not loaded")
	}
	return _wideDivide(arg0, arg1, arg2)
}

var _wideMultiply func(arg0 int32, arg1 int32, arg2 unsafe.Pointer) objectivec.IObject

// WideMultiply.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409273-widemultiply
func WideMultiply(arg0 int32, arg1 int32, arg2 unsafe.Pointer) objectivec.IObject {
	if _wideMultiply == nil {
		panic("coreservices: symbol WideMultiply not loaded")
	}
	return _wideMultiply(arg0, arg1, arg2)
}

var _wideNegate func(arg0 unsafe.Pointer) objectivec.IObject

// WideNegate.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409221-widenegate
func WideNegate(arg0 unsafe.Pointer) objectivec.IObject {
	if _wideNegate == nil {
		panic("coreservices: symbol WideNegate not loaded")
	}
	return _wideNegate(arg0)
}

var _wideShift func(arg0 unsafe.Pointer, arg1 int32) objectivec.IObject

// WideShift.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409239-wideshift
func WideShift(arg0 unsafe.Pointer, arg1 int32) objectivec.IObject {
	if _wideShift == nil {
		panic("coreservices: symbol WideShift not loaded")
	}
	return _wideShift(arg0, arg1)
}

var _wideSquareRoot func(arg0 unsafe.Pointer) uint32

// WideSquareRoot.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409227-widesquareroot
func WideSquareRoot(arg0 unsafe.Pointer) uint32 {
	if _wideSquareRoot == nil {
		panic("coreservices: symbol WideSquareRoot not loaded")
	}
	return _wideSquareRoot(arg0)
}

var _wideSubtract func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) objectivec.IObject

// WideSubtract.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409210-widesubtract
func WideSubtract(arg0 unsafe.Pointer, arg1 unsafe.Pointer) objectivec.IObject {
	if _wideSubtract == nil {
		panic("coreservices: symbol WideSubtract not loaded")
	}
	return _wideSubtract(arg0, arg1)
}

var _wideToSInt64 func(arg0 unsafe.Pointer) int64

// WideToSInt64.
//
// See: https://developer.apple.com/documentation/coreservices/1536356-widetosint64
func WideToSInt64(arg0 unsafe.Pointer) int64 {
	if _wideToSInt64 == nil {
		panic("coreservices: symbol WideToSInt64 not loaded")
	}
	return _wideToSInt64(arg0)
}

var _wideWideDivide func(arg0 unsafe.Pointer, arg1 int32, arg2 int32) objectivec.IObject

// WideWideDivide.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409265-widewidedivide
func WideWideDivide(arg0 unsafe.Pointer, arg1 int32, arg2 int32) objectivec.IObject {
	if _wideWideDivide == nil {
		panic("coreservices: symbol WideWideDivide not loaded")
	}
	return _wideWideDivide(arg0, arg1, arg2)
}

var _writePartialResource func(arg0 uintptr, arg1 int, arg2 int)

// WritePartialResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529305-writepartialresource
func WritePartialResource(arg0 uintptr, arg1 int, arg2 int) {
	if _writePartialResource == nil {
		panic("coreservices: symbol WritePartialResource not loaded")
	}
	_writePartialResource(arg0, arg1, arg2)
}

var _writeResource func(arg0 uintptr)

// WriteResource.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1529291-writeresource
func WriteResource(arg0 uintptr) {
	if _writeResource == nil {
		panic("coreservices: symbol WriteResource not loaded")
	}
	_writeResource(arg0)
}

var _x2Fix func(arg0 float64) objectivec.IObject

// X2Fix.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409237-x2fix
func X2Fix(arg0 float64) objectivec.IObject {
	if _x2Fix == nil {
		panic("coreservices: symbol X2Fix not loaded")
	}
	return _x2Fix(arg0)
}

var _x2Frac func(arg0 float64) unsafe.Pointer

// X2Frac.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1409261-x2frac
func X2Frac(arg0 float64) unsafe.Pointer {
	if _x2Frac == nil {
		panic("coreservices: symbol X2Frac not loaded")
	}
	return _x2Frac(arg0)
}

var _yieldToAnyThread func() int16

// YieldToAnyThread.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574238-yieldtoanythread
func YieldToAnyThread() int16 {
	if _yieldToAnyThread == nil {
		panic("coreservices: symbol YieldToAnyThread not loaded")
	}
	return _yieldToAnyThread()
}

var _yieldToThread func(arg0 ThreadID) int16

// YieldToThread.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/coreservices/1574240-yieldtothread
func YieldToThread(arg0 ThreadID) int16 {
	if _yieldToThread == nil {
		panic("coreservices: symbol YieldToThread not loaded")
	}
	return _yieldToThread(arg0)
}

var _annuity func(arg0 float64, arg1 float64) float64

// Annuity.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492633-annuity
func Annuity(arg0 float64, arg1 float64) float64 {
	if _annuity == nil {
		panic("coreservices: symbol annuity not loaded")
	}
	return _annuity(arg0, arg1)
}

var _compound func(arg0 float64, arg1 float64) float64

// Compound.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492661-compound
func Compound(arg0 float64, arg1 float64) float64 {
	if _compound == nil {
		panic("coreservices: symbol compound not loaded")
	}
	return _compound(arg0, arg1)
}

var _dec2f func(arg0 Decimal) float32

// Dec2f.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492641-dec2f
func Dec2f(arg0 Decimal) float32 {
	if _dec2f == nil {
		panic("coreservices: symbol dec2f not loaded")
	}
	return _dec2f(arg0)
}

var _dec2l func(arg0 Decimal) int

// Dec2l.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492660-dec2l
func Dec2l(arg0 Decimal) int {
	if _dec2l == nil {
		panic("coreservices: symbol dec2l not loaded")
	}
	return _dec2l(arg0)
}

var _dec2num func(arg0 Decimal) unsafe.Pointer

// Dec2num.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492667-dec2num
func Dec2num(arg0 Decimal) unsafe.Pointer {
	if _dec2num == nil {
		panic("coreservices: symbol dec2num not loaded")
	}
	return _dec2num(arg0)
}

var _dec2numl func(arg0 Decimal) unsafe.Pointer

// Dec2numl.
//
// See: https://developer.apple.com/documentation/coreservices/1492631-dec2numl
func Dec2numl(arg0 Decimal) unsafe.Pointer {
	if _dec2numl == nil {
		panic("coreservices: symbol dec2numl not loaded")
	}
	return _dec2numl(arg0)
}

var _dec2s func(arg0 Decimal) int16

// Dec2s.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492647-dec2s
func Dec2s(arg0 Decimal) int16 {
	if _dec2s == nil {
		panic("coreservices: symbol dec2s not loaded")
	}
	return _dec2s(arg0)
}

var _dec2str func(arg0 Decform, arg1 Decimal, arg2 int8)

// Dec2str.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492629-dec2str
func Dec2str(arg0 Decform, arg1 Decimal, arg2 int8) {
	if _dec2str == nil {
		panic("coreservices: symbol dec2str not loaded")
	}
	_dec2str(arg0, arg1, arg2)
}

var _dtox80 func(arg0 float64, arg1 unsafe.Pointer)

// Dtox80.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492664-dtox80
func Dtox80(arg0 float64, arg1 unsafe.Pointer) {
	if _dtox80 == nil {
		panic("coreservices: symbol dtox80 not loaded")
	}
	_dtox80(arg0, arg1)
}

var _kcfindapplesharepassword func(arg0 unsafe.Pointer, arg1 int8, arg2 int8, arg3 int8, arg4 int8, arg5 uint32, arg6 uint32, arg7 KCItemRef) int32

// Kcfindapplesharepassword.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563164-kcfindapplesharepassword
func Kcfindapplesharepassword(arg0 unsafe.Pointer, arg1 int8, arg2 int8, arg3 int8, arg4 int8, arg5 uint32, arg6 uint32, arg7 KCItemRef) int32 {
	if _kcfindapplesharepassword == nil {
		panic("coreservices: symbol kcfindapplesharepassword not loaded")
	}
	return _kcfindapplesharepassword(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _kcfindgenericpassword func(arg0 int8, arg1 int8, arg2 uint32, arg3 uint32, arg4 KCItemRef) int32

// Kcfindgenericpassword.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1562994-kcfindgenericpassword
func Kcfindgenericpassword(arg0 int8, arg1 int8, arg2 uint32, arg3 uint32, arg4 KCItemRef) int32 {
	if _kcfindgenericpassword == nil {
		panic("coreservices: symbol kcfindgenericpassword not loaded")
	}
	return _kcfindgenericpassword(arg0, arg1, arg2, arg3, arg4)
}

var _kcfindinternetpassword func(arg0 int8, arg1 int8, arg2 int8, arg3 uint16, arg4 uint32, arg5 uint32, arg6 uint32, arg7 uint32, arg8 KCItemRef) int32

// Kcfindinternetpassword.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563036-kcfindinternetpassword
func Kcfindinternetpassword(arg0 int8, arg1 int8, arg2 int8, arg3 uint16, arg4 uint32, arg5 uint32, arg6 uint32, arg7 uint32, arg8 KCItemRef) int32 {
	if _kcfindinternetpassword == nil {
		panic("coreservices: symbol kcfindinternetpassword not loaded")
	}
	return _kcfindinternetpassword(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

var _kcfindinternetpasswordwithpath func(arg0 int8, arg1 int8, arg2 int8, arg3 int8, arg4 uint16, arg5 uint32, arg6 uint32, arg7 uint32, arg8 uint32, arg9 KCItemRef) int32

// Kcfindinternetpasswordwithpath.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563135-kcfindinternetpasswordwithpath
func Kcfindinternetpasswordwithpath(arg0 int8, arg1 int8, arg2 int8, arg3 int8, arg4 uint16, arg5 uint32, arg6 uint32, arg7 uint32, arg8 uint32, arg9 KCItemRef) int32 {
	if _kcfindinternetpasswordwithpath == nil {
		panic("coreservices: symbol kcfindinternetpasswordwithpath not loaded")
	}
	return _kcfindinternetpasswordwithpath(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

var _kcgetkeychainname func(arg0 KCRef, arg1 int8) int32

// Kcgetkeychainname.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/coreservices/1563005-kcgetkeychainname
func Kcgetkeychainname(arg0 KCRef, arg1 int8) int32 {
	if _kcgetkeychainname == nil {
		panic("coreservices: symbol kcgetkeychainname not loaded")
	}
	return _kcgetkeychainname(arg0, arg1)
}

var _ldtox80 func(arg0 unsafe.Pointer, arg1 unsafe.Pointer)

// Ldtox80.
//
// See: https://developer.apple.com/documentation/coreservices/1492662-ldtox80
func Ldtox80(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	if _ldtox80 == nil {
		panic("coreservices: symbol ldtox80 not loaded")
	}
	_ldtox80(arg0, arg1)
}

var _num2dec func(arg0 Decform, arg1 unsafe.Pointer, arg2 Decimal)

// Num2dec.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492649-num2dec
func Num2dec(arg0 Decform, arg1 unsafe.Pointer, arg2 Decimal) {
	if _num2dec == nil {
		panic("coreservices: symbol num2dec not loaded")
	}
	_num2dec(arg0, arg1, arg2)
}

var _num2decl func(arg0 Decform, arg1 unsafe.Pointer, arg2 Decimal)

// Num2decl.
//
// See: https://developer.apple.com/documentation/coreservices/1492642-num2decl
func Num2decl(arg0 Decform, arg1 unsafe.Pointer, arg2 Decimal) {
	if _num2decl == nil {
		panic("coreservices: symbol num2decl not loaded")
	}
	_num2decl(arg0, arg1, arg2)
}

var _numtostring func(arg0 int, arg1 int8)

// Numtostring.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1579613-numtostring
func Numtostring(arg0 int, arg1 int8) {
	if _numtostring == nil {
		panic("coreservices: symbol numtostring not loaded")
	}
	_numtostring(arg0, arg1)
}

var _randomx func(arg0 unsafe.Pointer) unsafe.Pointer

// Randomx.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492636-randomx
func Randomx(arg0 unsafe.Pointer) unsafe.Pointer {
	if _randomx == nil {
		panic("coreservices: symbol randomx not loaded")
	}
	return _randomx(arg0)
}

var _relation func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) Relop

// Relation.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492646-relation
func Relation(arg0 unsafe.Pointer, arg1 unsafe.Pointer) Relop {
	if _relation == nil {
		panic("coreservices: symbol relation not loaded")
	}
	return _relation(arg0, arg1)
}

var _relationl func(arg0 unsafe.Pointer, arg1 unsafe.Pointer) Relop

// Relationl.
//
// See: https://developer.apple.com/documentation/coreservices/1492627-relationl
func Relationl(arg0 unsafe.Pointer, arg1 unsafe.Pointer) Relop {
	if _relationl == nil {
		panic("coreservices: symbol relationl not loaded")
	}
	return _relationl(arg0, arg1)
}

var _str2dec func(arg0 int8, arg1 int16, arg2 Decimal, arg3 int16)

// Str2dec.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492635-str2dec
func Str2dec(arg0 int8, arg1 int16, arg2 Decimal, arg3 int16) {
	if _str2dec == nil {
		panic("coreservices: symbol str2dec not loaded")
	}
	_str2dec(arg0, arg1, arg2, arg3)
}

var _vAEBuildAppleEvent func(theClass uint32, theID uint32, addressType uint32, addressData unsafe.Pointer, addressLength corefoundation.CGSize, returnID unsafe.Pointer, transactionID unsafe.Pointer, resultEvt *AEDesc, err unsafe.Pointer, paramsFmt string, args unsafe.Pointer) int32

// VAEBuildAppleEvent allows you to encapsulate calls to [AEBuildAppleEvent] in a wrapper routine.
//
// See: https://developer.apple.com/documentation/coreservices/1441729-vaebuildappleevent
func VAEBuildAppleEvent(theClass uint32, theID uint32, addressType uint32, addressData unsafe.Pointer, addressLength corefoundation.CGSize, returnID unsafe.Pointer, transactionID unsafe.Pointer, resultEvt *AEDesc, err unsafe.Pointer, paramsFmt string, args unsafe.Pointer) int32 {
	if _vAEBuildAppleEvent == nil {
		panic("coreservices: symbol vAEBuildAppleEvent not loaded")
	}
	return _vAEBuildAppleEvent(theClass, theID, addressType, addressData, addressLength, returnID, transactionID, resultEvt, err, paramsFmt, args)
}

var _vAEBuildDesc func(dst unsafe.Pointer, err unsafe.Pointer, src string, args unsafe.Pointer) int32

// VAEBuildDesc allows you to encapsulate calls to [AEBuildDesc] in your own wrapper routines.
//
// See: https://developer.apple.com/documentation/coreservices/1446775-vaebuilddesc
func VAEBuildDesc(dst unsafe.Pointer, err unsafe.Pointer, src string, args unsafe.Pointer) int32 {
	if _vAEBuildDesc == nil {
		panic("coreservices: symbol vAEBuildDesc not loaded")
	}
	return _vAEBuildDesc(dst, err, src, args)
}

var _vAEBuildParameters func(event *AEDesc, err unsafe.Pointer, format string, args unsafe.Pointer) int32

// VAEBuildParameters allows you to encapsulate calls to [AEBuildParameters] in your own `stdarg`-style wrapper routines, using techniques similar to those allowed by vsprintf.
//
// See: https://developer.apple.com/documentation/coreservices/1448040-vaebuildparameters
func VAEBuildParameters(event *AEDesc, err unsafe.Pointer, format string, args unsafe.Pointer) int32 {
	if _vAEBuildParameters == nil {
		panic("coreservices: symbol vAEBuildParameters not loaded")
	}
	return _vAEBuildParameters(event, err, format, args)
}

var _x80tod func(arg0 unsafe.Pointer) float64

// X80tod.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/coreservices/1492658-x80tod
func X80tod(arg0 unsafe.Pointer) float64 {
	if _x80tod == nil {
		panic("coreservices: symbol x80tod not loaded")
	}
	return _x80tod(arg0)
}

var _x80told func(arg0 unsafe.Pointer, arg1 unsafe.Pointer)

// X80told.
//
// See: https://developer.apple.com/documentation/coreservices/1492670-x80told
func X80told(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	if _x80told == nil {
		panic("coreservices: symbol x80told not loaded")
	}
	_x80told(arg0, arg1)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_aEBuildAppleEvent, frameworkHandle, "AEBuildAppleEvent")
		registerFunc(&_aEBuildDesc, frameworkHandle, "AEBuildDesc")
		registerFunc(&_aEBuildParameters, frameworkHandle, "AEBuildParameters")
		registerFunc(&_aECallObjectAccessor, frameworkHandle, "AECallObjectAccessor")
		registerFunc(&_aECheckIsRecord, frameworkHandle, "AECheckIsRecord")
		registerFunc(&_aECoerceDesc, frameworkHandle, "AECoerceDesc")
		registerFunc(&_aECoercePtr, frameworkHandle, "AECoercePtr")
		registerFunc(&_aECompareDesc, frameworkHandle, "AECompareDesc")
		registerFunc(&_aECountItems, frameworkHandle, "AECountItems")
		registerFunc(&_aECreateAppleEvent, frameworkHandle, "AECreateAppleEvent")
		registerFunc(&_aECreateDesc, frameworkHandle, "AECreateDesc")
		registerFunc(&_aECreateDescFromExternalPtr, frameworkHandle, "AECreateDescFromExternalPtr")
		registerFunc(&_aECreateList, frameworkHandle, "AECreateList")
		registerFunc(&_aECreateRemoteProcessResolver, frameworkHandle, "AECreateRemoteProcessResolver")
		registerFunc(&_aEDecodeMessage, frameworkHandle, "AEDecodeMessage")
		registerFunc(&_aEDeleteItem, frameworkHandle, "AEDeleteItem")
		registerFunc(&_aEDeleteParam, frameworkHandle, "AEDeleteParam")
		registerFunc(&_aEDeterminePermissionToAutomateTarget, frameworkHandle, "AEDeterminePermissionToAutomateTarget")
		registerFunc(&_aEDisposeDesc, frameworkHandle, "AEDisposeDesc")
		registerFunc(&_aEDisposeRemoteProcessResolver, frameworkHandle, "AEDisposeRemoteProcessResolver")
		registerFunc(&_aEDisposeToken, frameworkHandle, "AEDisposeToken")
		registerFunc(&_aEDuplicateDesc, frameworkHandle, "AEDuplicateDesc")
		registerFunc(&_aEFlattenDesc, frameworkHandle, "AEFlattenDesc")
		registerFunc(&_aEGetArray, frameworkHandle, "AEGetArray")
		registerFunc(&_aEGetAttributeDesc, frameworkHandle, "AEGetAttributeDesc")
		registerFunc(&_aEGetAttributePtr, frameworkHandle, "AEGetAttributePtr")
		registerFunc(&_aEGetCoercionHandler, frameworkHandle, "AEGetCoercionHandler")
		registerFunc(&_aEGetDescData, frameworkHandle, "AEGetDescData")
		registerFunc(&_aEGetDescDataRange, frameworkHandle, "AEGetDescDataRange")
		registerFunc(&_aEGetDescDataSize, frameworkHandle, "AEGetDescDataSize")
		registerFunc(&_aEGetEventHandler, frameworkHandle, "AEGetEventHandler")
		registerFunc(&_aEGetNthDesc, frameworkHandle, "AEGetNthDesc")
		registerFunc(&_aEGetNthPtr, frameworkHandle, "AEGetNthPtr")
		registerFunc(&_aEGetObjectAccessor, frameworkHandle, "AEGetObjectAccessor")
		registerFunc(&_aEGetParamDesc, frameworkHandle, "AEGetParamDesc")
		registerFunc(&_aEGetParamPtr, frameworkHandle, "AEGetParamPtr")
		registerFunc(&_aEGetRegisteredMachPort, frameworkHandle, "AEGetRegisteredMachPort")
		registerFunc(&_aEGetSpecialHandler, frameworkHandle, "AEGetSpecialHandler")
		registerFunc(&_aEInitializeDesc, frameworkHandle, "AEInitializeDesc")
		registerFunc(&_aEInstallCoercionHandler, frameworkHandle, "AEInstallCoercionHandler")
		registerFunc(&_aEInstallEventHandler, frameworkHandle, "AEInstallEventHandler")
		registerFunc(&_aEInstallObjectAccessor, frameworkHandle, "AEInstallObjectAccessor")
		registerFunc(&_aEInstallSpecialHandler, frameworkHandle, "AEInstallSpecialHandler")
		registerFunc(&_aEManagerInfo, frameworkHandle, "AEManagerInfo")
		registerFunc(&_aEObjectInit, frameworkHandle, "AEObjectInit")
		registerFunc(&_aEPrintDescToHandle, frameworkHandle, "AEPrintDescToHandle")
		registerFunc(&_aEProcessMessage, frameworkHandle, "AEProcessMessage")
		registerFunc(&_aEPutArray, frameworkHandle, "AEPutArray")
		registerFunc(&_aEPutAttributeDesc, frameworkHandle, "AEPutAttributeDesc")
		registerFunc(&_aEPutAttributePtr, frameworkHandle, "AEPutAttributePtr")
		registerFunc(&_aEPutDesc, frameworkHandle, "AEPutDesc")
		registerFunc(&_aEPutParamDesc, frameworkHandle, "AEPutParamDesc")
		registerFunc(&_aEPutParamPtr, frameworkHandle, "AEPutParamPtr")
		registerFunc(&_aEPutPtr, frameworkHandle, "AEPutPtr")
		registerFunc(&_aERemoteProcessResolverGetProcesses, frameworkHandle, "AERemoteProcessResolverGetProcesses")
		registerFunc(&_aERemoteProcessResolverScheduleWithRunLoop, frameworkHandle, "AERemoteProcessResolverScheduleWithRunLoop")
		registerFunc(&_aERemoveCoercionHandler, frameworkHandle, "AERemoveCoercionHandler")
		registerFunc(&_aERemoveEventHandler, frameworkHandle, "AERemoveEventHandler")
		registerFunc(&_aERemoveObjectAccessor, frameworkHandle, "AERemoveObjectAccessor")
		registerFunc(&_aERemoveSpecialHandler, frameworkHandle, "AERemoveSpecialHandler")
		registerFunc(&_aEReplaceDescData, frameworkHandle, "AEReplaceDescData")
		registerFunc(&_aEResolve, frameworkHandle, "AEResolve")
		registerFunc(&_aESendMessage, frameworkHandle, "AESendMessage")
		registerFunc(&_aESetObjectCallbacks, frameworkHandle, "AESetObjectCallbacks")
		registerFunc(&_aESizeOfAttribute, frameworkHandle, "AESizeOfAttribute")
		registerFunc(&_aESizeOfFlattenedDesc, frameworkHandle, "AESizeOfFlattenedDesc")
		registerFunc(&_aESizeOfNthItem, frameworkHandle, "AESizeOfNthItem")
		registerFunc(&_aESizeOfParam, frameworkHandle, "AESizeOfParam")
		registerFunc(&_aEStreamClose, frameworkHandle, "AEStreamClose")
		registerFunc(&_aEStreamCloseDesc, frameworkHandle, "AEStreamCloseDesc")
		registerFunc(&_aEStreamCloseList, frameworkHandle, "AEStreamCloseList")
		registerFunc(&_aEStreamCloseRecord, frameworkHandle, "AEStreamCloseRecord")
		registerFunc(&_aEStreamCreateEvent, frameworkHandle, "AEStreamCreateEvent")
		registerFunc(&_aEStreamOpen, frameworkHandle, "AEStreamOpen")
		registerFunc(&_aEStreamOpenDesc, frameworkHandle, "AEStreamOpenDesc")
		registerFunc(&_aEStreamOpenEvent, frameworkHandle, "AEStreamOpenEvent")
		registerFunc(&_aEStreamOpenKeyDesc, frameworkHandle, "AEStreamOpenKeyDesc")
		registerFunc(&_aEStreamOpenList, frameworkHandle, "AEStreamOpenList")
		registerFunc(&_aEStreamOpenRecord, frameworkHandle, "AEStreamOpenRecord")
		registerFunc(&_aEStreamOptionalParam, frameworkHandle, "AEStreamOptionalParam")
		registerFunc(&_aEStreamSetRecordType, frameworkHandle, "AEStreamSetRecordType")
		registerFunc(&_aEStreamWriteAEDesc, frameworkHandle, "AEStreamWriteAEDesc")
		registerFunc(&_aEStreamWriteData, frameworkHandle, "AEStreamWriteData")
		registerFunc(&_aEStreamWriteDesc, frameworkHandle, "AEStreamWriteDesc")
		registerFunc(&_aEStreamWriteKey, frameworkHandle, "AEStreamWriteKey")
		registerFunc(&_aEStreamWriteKeyDesc, frameworkHandle, "AEStreamWriteKeyDesc")
		registerFunc(&_aEUnflattenDesc, frameworkHandle, "AEUnflattenDesc")
		registerFunc(&_aEUnflattenDescFromBytes, frameworkHandle, "AEUnflattenDescFromBytes")
		registerFunc(&_absoluteDeltaToDuration, frameworkHandle, "AbsoluteDeltaToDuration")
		registerFunc(&_absoluteDeltaToNanoseconds, frameworkHandle, "AbsoluteDeltaToNanoseconds")
		registerFunc(&_absoluteToDuration, frameworkHandle, "AbsoluteToDuration")
		registerFunc(&_absoluteToNanoseconds, frameworkHandle, "AbsoluteToNanoseconds")
		registerFunc(&_acquireIconRef, frameworkHandle, "AcquireIconRef")
		registerFunc(&_addAbsoluteToAbsolute, frameworkHandle, "AddAbsoluteToAbsolute")
		registerFunc(&_addAtomic, frameworkHandle, "AddAtomic")
		registerFunc(&_addAtomic16, frameworkHandle, "AddAtomic16")
		registerFunc(&_addAtomic8, frameworkHandle, "AddAtomic8")
		registerFunc(&_addCollectionItem, frameworkHandle, "AddCollectionItem")
		registerFunc(&_addCollectionItemHdl, frameworkHandle, "AddCollectionItemHdl")
		registerFunc(&_addDurationToAbsolute, frameworkHandle, "AddDurationToAbsolute")
		registerFunc(&_addFolderDescriptor, frameworkHandle, "AddFolderDescriptor")
		registerFunc(&_addNanosecondsToAbsolute, frameworkHandle, "AddNanosecondsToAbsolute")
		registerFunc(&_addResource, frameworkHandle, "AddResource")
		registerFunc(&_batteryCount, frameworkHandle, "BatteryCount")
		registerFunc(&_bitAnd, frameworkHandle, "BitAnd")
		registerFunc(&_bitAndAtomic, frameworkHandle, "BitAndAtomic")
		registerFunc(&_bitAndAtomic16, frameworkHandle, "BitAndAtomic16")
		registerFunc(&_bitAndAtomic8, frameworkHandle, "BitAndAtomic8")
		registerFunc(&_bitClr, frameworkHandle, "BitClr")
		registerFunc(&_bitNot, frameworkHandle, "BitNot")
		registerFunc(&_bitOr, frameworkHandle, "BitOr")
		registerFunc(&_bitOrAtomic, frameworkHandle, "BitOrAtomic")
		registerFunc(&_bitOrAtomic16, frameworkHandle, "BitOrAtomic16")
		registerFunc(&_bitOrAtomic8, frameworkHandle, "BitOrAtomic8")
		registerFunc(&_bitSet, frameworkHandle, "BitSet")
		registerFunc(&_bitShift, frameworkHandle, "BitShift")
		registerFunc(&_bitTst, frameworkHandle, "BitTst")
		registerFunc(&_bitXor, frameworkHandle, "BitXor")
		registerFunc(&_bitXorAtomic, frameworkHandle, "BitXorAtomic")
		registerFunc(&_bitXorAtomic16, frameworkHandle, "BitXorAtomic16")
		registerFunc(&_bitXorAtomic8, frameworkHandle, "BitXorAtomic8")
		registerFunc(&_cSBackupIsItemExcluded, frameworkHandle, "CSBackupIsItemExcluded")
		registerFunc(&_cSBackupSetItemExcluded, frameworkHandle, "CSBackupSetItemExcluded")
		registerFunc(&_cSCopyMachineName, frameworkHandle, "CSCopyMachineName")
		registerFunc(&_cSCopyUserName, frameworkHandle, "CSCopyUserName")
		registerFunc(&_cSDiskSpaceCancelRecovery, frameworkHandle, "CSDiskSpaceCancelRecovery")
		registerFunc(&_cSDiskSpaceGetRecoveryEstimate, frameworkHandle, "CSDiskSpaceGetRecoveryEstimate")
		registerFunc(&_cSDiskSpaceStartRecovery, frameworkHandle, "CSDiskSpaceStartRecovery")
		registerFunc(&_cSGetComponentsThreadMode, frameworkHandle, "CSGetComponentsThreadMode")
		registerFunc(&_cSGetDefaultIdentityAuthority, frameworkHandle, "CSGetDefaultIdentityAuthority")
		registerFunc(&_cSGetLocalIdentityAuthority, frameworkHandle, "CSGetLocalIdentityAuthority")
		registerFunc(&_cSGetManagedIdentityAuthority, frameworkHandle, "CSGetManagedIdentityAuthority")
		registerFunc(&_cSIdentityAddAlias, frameworkHandle, "CSIdentityAddAlias")
		registerFunc(&_cSIdentityAddMember, frameworkHandle, "CSIdentityAddMember")
		registerFunc(&_cSIdentityAuthenticateUsingPassword, frameworkHandle, "CSIdentityAuthenticateUsingPassword")
		registerFunc(&_cSIdentityAuthorityCopyLocalizedName, frameworkHandle, "CSIdentityAuthorityCopyLocalizedName")
		registerFunc(&_cSIdentityAuthorityGetTypeID, frameworkHandle, "CSIdentityAuthorityGetTypeID")
		registerFunc(&_cSIdentityCommit, frameworkHandle, "CSIdentityCommit")
		registerFunc(&_cSIdentityCommitAsynchronously, frameworkHandle, "CSIdentityCommitAsynchronously")
		registerFunc(&_cSIdentityCreate, frameworkHandle, "CSIdentityCreate")
		registerFunc(&_cSIdentityCreateCopy, frameworkHandle, "CSIdentityCreateCopy")
		registerFunc(&_cSIdentityCreateGroupMembershipQuery, frameworkHandle, "CSIdentityCreateGroupMembershipQuery")
		registerFunc(&_cSIdentityCreatePersistentReference, frameworkHandle, "CSIdentityCreatePersistentReference")
		registerFunc(&_cSIdentityDelete, frameworkHandle, "CSIdentityDelete")
		registerFunc(&_cSIdentityGetAliases, frameworkHandle, "CSIdentityGetAliases")
		registerFunc(&_cSIdentityGetAuthority, frameworkHandle, "CSIdentityGetAuthority")
		registerFunc(&_cSIdentityGetCertificate, frameworkHandle, "CSIdentityGetCertificate")
		registerFunc(&_cSIdentityGetClass, frameworkHandle, "CSIdentityGetClass")
		registerFunc(&_cSIdentityGetEmailAddress, frameworkHandle, "CSIdentityGetEmailAddress")
		registerFunc(&_cSIdentityGetFullName, frameworkHandle, "CSIdentityGetFullName")
		registerFunc(&_cSIdentityGetImageData, frameworkHandle, "CSIdentityGetImageData")
		registerFunc(&_cSIdentityGetImageDataType, frameworkHandle, "CSIdentityGetImageDataType")
		registerFunc(&_cSIdentityGetImageURL, frameworkHandle, "CSIdentityGetImageURL")
		registerFunc(&_cSIdentityGetPosixID, frameworkHandle, "CSIdentityGetPosixID")
		registerFunc(&_cSIdentityGetPosixName, frameworkHandle, "CSIdentityGetPosixName")
		registerFunc(&_cSIdentityGetTypeID, frameworkHandle, "CSIdentityGetTypeID")
		registerFunc(&_cSIdentityGetUUID, frameworkHandle, "CSIdentityGetUUID")
		registerFunc(&_cSIdentityIsCommitting, frameworkHandle, "CSIdentityIsCommitting")
		registerFunc(&_cSIdentityIsEnabled, frameworkHandle, "CSIdentityIsEnabled")
		registerFunc(&_cSIdentityIsHidden, frameworkHandle, "CSIdentityIsHidden")
		registerFunc(&_cSIdentityIsMemberOfGroup, frameworkHandle, "CSIdentityIsMemberOfGroup")
		registerFunc(&_cSIdentityQueryCopyResults, frameworkHandle, "CSIdentityQueryCopyResults")
		registerFunc(&_cSIdentityQueryCreate, frameworkHandle, "CSIdentityQueryCreate")
		registerFunc(&_cSIdentityQueryCreateForCurrentUser, frameworkHandle, "CSIdentityQueryCreateForCurrentUser")
		registerFunc(&_cSIdentityQueryCreateForName, frameworkHandle, "CSIdentityQueryCreateForName")
		registerFunc(&_cSIdentityQueryCreateForPersistentReference, frameworkHandle, "CSIdentityQueryCreateForPersistentReference")
		registerFunc(&_cSIdentityQueryCreateForPosixID, frameworkHandle, "CSIdentityQueryCreateForPosixID")
		registerFunc(&_cSIdentityQueryCreateForUUID, frameworkHandle, "CSIdentityQueryCreateForUUID")
		registerFunc(&_cSIdentityQueryExecute, frameworkHandle, "CSIdentityQueryExecute")
		registerFunc(&_cSIdentityQueryExecuteAsynchronously, frameworkHandle, "CSIdentityQueryExecuteAsynchronously")
		registerFunc(&_cSIdentityQueryGetTypeID, frameworkHandle, "CSIdentityQueryGetTypeID")
		registerFunc(&_cSIdentityQueryStop, frameworkHandle, "CSIdentityQueryStop")
		registerFunc(&_cSIdentityRemoveAlias, frameworkHandle, "CSIdentityRemoveAlias")
		registerFunc(&_cSIdentityRemoveClient, frameworkHandle, "CSIdentityRemoveClient")
		registerFunc(&_cSIdentityRemoveMember, frameworkHandle, "CSIdentityRemoveMember")
		registerFunc(&_cSIdentitySetCertificate, frameworkHandle, "CSIdentitySetCertificate")
		registerFunc(&_cSIdentitySetEmailAddress, frameworkHandle, "CSIdentitySetEmailAddress")
		registerFunc(&_cSIdentitySetFullName, frameworkHandle, "CSIdentitySetFullName")
		registerFunc(&_cSIdentitySetImageData, frameworkHandle, "CSIdentitySetImageData")
		registerFunc(&_cSIdentitySetImageURL, frameworkHandle, "CSIdentitySetImageURL")
		registerFunc(&_cSIdentitySetIsEnabled, frameworkHandle, "CSIdentitySetIsEnabled")
		registerFunc(&_cSIdentitySetPassword, frameworkHandle, "CSIdentitySetPassword")
		registerFunc(&_cSSetComponentsThreadMode, frameworkHandle, "CSSetComponentsThreadMode")
		registerFunc(&_callComponentCanDo, frameworkHandle, "CallComponentCanDo")
		registerFunc(&_callComponentClose, frameworkHandle, "CallComponentClose")
		registerFunc(&_callComponentDispatch, frameworkHandle, "CallComponentDispatch")
		registerFunc(&_callComponentFunction, frameworkHandle, "CallComponentFunction")
		registerFunc(&_callComponentFunctionWithStorage, frameworkHandle, "CallComponentFunctionWithStorage")
		registerFunc(&_callComponentFunctionWithStorageProcInfo, frameworkHandle, "CallComponentFunctionWithStorageProcInfo")
		registerFunc(&_callComponentGetMPWorkFunction, frameworkHandle, "CallComponentGetMPWorkFunction")
		registerFunc(&_callComponentGetPublicResource, frameworkHandle, "CallComponentGetPublicResource")
		registerFunc(&_callComponentOpen, frameworkHandle, "CallComponentOpen")
		registerFunc(&_callComponentRegister, frameworkHandle, "CallComponentRegister")
		registerFunc(&_callComponentTarget, frameworkHandle, "CallComponentTarget")
		registerFunc(&_callComponentUnregister, frameworkHandle, "CallComponentUnregister")
		registerFunc(&_callComponentVersion, frameworkHandle, "CallComponentVersion")
		registerFunc(&_captureComponent, frameworkHandle, "CaptureComponent")
		registerFunc(&_changeTextToUnicodeInfo, frameworkHandle, "ChangeTextToUnicodeInfo")
		registerFunc(&_changeUnicodeToTextInfo, frameworkHandle, "ChangeUnicodeToTextInfo")
		registerFunc(&_changedResource, frameworkHandle, "ChangedResource")
		registerFunc(&_cloneCollection, frameworkHandle, "CloneCollection")
		registerFunc(&_closeComponent, frameworkHandle, "CloseComponent")
		registerFunc(&_closeComponentResFile, frameworkHandle, "CloseComponentResFile")
		registerFunc(&_closeResFile, frameworkHandle, "CloseResFile")
		registerFunc(&_collectionTagExists, frameworkHandle, "CollectionTagExists")
		registerFunc(&_compareAndSwap, frameworkHandle, "CompareAndSwap")
		registerFunc(&_compositeIconRef, frameworkHandle, "CompositeIconRef")
		registerFunc(&_convertFromPStringToUnicode, frameworkHandle, "ConvertFromPStringToUnicode")
		registerFunc(&_convertFromTextToUnicode, frameworkHandle, "ConvertFromTextToUnicode")
		registerFunc(&_convertFromUnicodeToPString, frameworkHandle, "ConvertFromUnicodeToPString")
		registerFunc(&_convertFromUnicodeToScriptCodeRun, frameworkHandle, "ConvertFromUnicodeToScriptCodeRun")
		registerFunc(&_convertFromUnicodeToText, frameworkHandle, "ConvertFromUnicodeToText")
		registerFunc(&_convertFromUnicodeToTextRun, frameworkHandle, "ConvertFromUnicodeToTextRun")
		registerFunc(&_copyCollection, frameworkHandle, "CopyCollection")
		registerFunc(&_coreEndianFlipData, frameworkHandle, "CoreEndianFlipData")
		registerFunc(&_coreEndianGetFlipper, frameworkHandle, "CoreEndianGetFlipper")
		registerFunc(&_coreEndianInstallFlipper, frameworkHandle, "CoreEndianInstallFlipper")
		registerFunc(&_count1Resources, frameworkHandle, "Count1Resources")
		registerFunc(&_count1Types, frameworkHandle, "Count1Types")
		registerFunc(&_countCollectionItems, frameworkHandle, "CountCollectionItems")
		registerFunc(&_countCollectionOwners, frameworkHandle, "CountCollectionOwners")
		registerFunc(&_countCollectionTags, frameworkHandle, "CountCollectionTags")
		registerFunc(&_countComponentInstances, frameworkHandle, "CountComponentInstances")
		registerFunc(&_countComponents, frameworkHandle, "CountComponents")
		registerFunc(&_countResources, frameworkHandle, "CountResources")
		registerFunc(&_countTaggedCollectionItems, frameworkHandle, "CountTaggedCollectionItems")
		registerFunc(&_countTypes, frameworkHandle, "CountTypes")
		registerFunc(&_countUnicodeMappings, frameworkHandle, "CountUnicodeMappings")
		registerFunc(&_createCompDescriptor, frameworkHandle, "CreateCompDescriptor")
		registerFunc(&_createLogicalDescriptor, frameworkHandle, "CreateLogicalDescriptor")
		registerFunc(&_createObjSpecifier, frameworkHandle, "CreateObjSpecifier")
		registerFunc(&_createOffsetDescriptor, frameworkHandle, "CreateOffsetDescriptor")
		registerFunc(&_createRangeDescriptor, frameworkHandle, "CreateRangeDescriptor")
		registerFunc(&_createTextEncoding, frameworkHandle, "CreateTextEncoding")
		registerFunc(&_createTextToUnicodeInfo, frameworkHandle, "CreateTextToUnicodeInfo")
		registerFunc(&_createTextToUnicodeInfoByEncoding, frameworkHandle, "CreateTextToUnicodeInfoByEncoding")
		registerFunc(&_createThreadPool, frameworkHandle, "CreateThreadPool")
		registerFunc(&_createUnicodeToTextInfo, frameworkHandle, "CreateUnicodeToTextInfo")
		registerFunc(&_createUnicodeToTextInfoByEncoding, frameworkHandle, "CreateUnicodeToTextInfoByEncoding")
		registerFunc(&_createUnicodeToTextRunInfo, frameworkHandle, "CreateUnicodeToTextRunInfo")
		registerFunc(&_createUnicodeToTextRunInfoByEncoding, frameworkHandle, "CreateUnicodeToTextRunInfoByEncoding")
		registerFunc(&_createUnicodeToTextRunInfoByScriptCode, frameworkHandle, "CreateUnicodeToTextRunInfoByScriptCode")
		registerFunc(&_curResFile, frameworkHandle, "CurResFile")
		registerFunc(&_currentProcessorSpeed, frameworkHandle, "CurrentProcessorSpeed")
		registerFunc(&_dCSCopyTextDefinition, frameworkHandle, "DCSCopyTextDefinition")
		registerFunc(&_dCSGetTermRangeInString, frameworkHandle, "DCSGetTermRangeInString")
		registerFunc(&_debugAssert, frameworkHandle, "DebugAssert")
		registerFunc(&_decrementAtomic, frameworkHandle, "DecrementAtomic")
		registerFunc(&_decrementAtomic16, frameworkHandle, "DecrementAtomic16")
		registerFunc(&_decrementAtomic8, frameworkHandle, "DecrementAtomic8")
		registerFunc(&_delay, frameworkHandle, "Delay")
		registerFunc(&_delegateComponentCall, frameworkHandle, "DelegateComponentCall")
		registerFunc(&_deleteGestaltValue, frameworkHandle, "DeleteGestaltValue")
		registerFunc(&_dequeue, frameworkHandle, "Dequeue")
		registerFunc(&_detachResource, frameworkHandle, "DetachResource")
		registerFunc(&_detachResourceFile, frameworkHandle, "DetachResourceFile")
		registerFunc(&_determineIfPathIsEnclosedByFolder, frameworkHandle, "DetermineIfPathIsEnclosedByFolder")
		registerFunc(&_disposeAECoerceDescUPP, frameworkHandle, "DisposeAECoerceDescUPP")
		registerFunc(&_disposeAECoercePtrUPP, frameworkHandle, "DisposeAECoercePtrUPP")
		registerFunc(&_disposeAEDisposeExternalUPP, frameworkHandle, "DisposeAEDisposeExternalUPP")
		registerFunc(&_disposeAEEventHandlerUPP, frameworkHandle, "DisposeAEEventHandlerUPP")
		registerFunc(&_disposeCollection, frameworkHandle, "DisposeCollection")
		registerFunc(&_disposeCollectionExceptionUPP, frameworkHandle, "DisposeCollectionExceptionUPP")
		registerFunc(&_disposeCollectionFlattenUPP, frameworkHandle, "DisposeCollectionFlattenUPP")
		registerFunc(&_disposeComponentFunctionUPP, frameworkHandle, "DisposeComponentFunctionUPP")
		registerFunc(&_disposeComponentMPWorkFunctionUPP, frameworkHandle, "DisposeComponentMPWorkFunctionUPP")
		registerFunc(&_disposeComponentRoutineUPP, frameworkHandle, "DisposeComponentRoutineUPP")
		registerFunc(&_disposeDebugAssertOutputHandlerUPP, frameworkHandle, "DisposeDebugAssertOutputHandlerUPP")
		registerFunc(&_disposeDebugComponent, frameworkHandle, "DisposeDebugComponent")
		registerFunc(&_disposeDebugComponentCallbackUPP, frameworkHandle, "DisposeDebugComponentCallbackUPP")
		registerFunc(&_disposeDebuggerDisposeThreadUPP, frameworkHandle, "DisposeDebuggerDisposeThreadUPP")
		registerFunc(&_disposeDebuggerNewThreadUPP, frameworkHandle, "DisposeDebuggerNewThreadUPP")
		registerFunc(&_disposeDebuggerThreadSchedulerUPP, frameworkHandle, "DisposeDebuggerThreadSchedulerUPP")
		registerFunc(&_disposeDeferredTaskUPP, frameworkHandle, "DisposeDeferredTaskUPP")
		registerFunc(&_disposeExceptionHandlerUPP, frameworkHandle, "DisposeExceptionHandlerUPP")
		registerFunc(&_disposeFNSubscriptionUPP, frameworkHandle, "DisposeFNSubscriptionUPP")
		registerFunc(&_disposeFSVolumeEjectUPP, frameworkHandle, "DisposeFSVolumeEjectUPP")
		registerFunc(&_disposeFSVolumeMountUPP, frameworkHandle, "DisposeFSVolumeMountUPP")
		registerFunc(&_disposeFSVolumeUnmountUPP, frameworkHandle, "DisposeFSVolumeUnmountUPP")
		registerFunc(&_disposeFolderManagerNotificationUPP, frameworkHandle, "DisposeFolderManagerNotificationUPP")
		registerFunc(&_disposeGetMissingComponentResourceUPP, frameworkHandle, "DisposeGetMissingComponentResourceUPP")
		registerFunc(&_disposeHandle, frameworkHandle, "DisposeHandle")
		registerFunc(&_disposeIOCompletionUPP, frameworkHandle, "DisposeIOCompletionUPP")
		registerFunc(&_disposeIndexToUCStringUPP, frameworkHandle, "DisposeIndexToUCStringUPP")
		registerFunc(&_disposeKCCallbackUPP, frameworkHandle, "DisposeKCCallbackUPP")
		registerFunc(&_disposeOSLAccessorUPP, frameworkHandle, "DisposeOSLAccessorUPP")
		registerFunc(&_disposeOSLAdjustMarksUPP, frameworkHandle, "DisposeOSLAdjustMarksUPP")
		registerFunc(&_disposeOSLCompareUPP, frameworkHandle, "DisposeOSLCompareUPP")
		registerFunc(&_disposeOSLCountUPP, frameworkHandle, "DisposeOSLCountUPP")
		registerFunc(&_disposeOSLDisposeTokenUPP, frameworkHandle, "DisposeOSLDisposeTokenUPP")
		registerFunc(&_disposeOSLGetErrDescUPP, frameworkHandle, "DisposeOSLGetErrDescUPP")
		registerFunc(&_disposeOSLGetMarkTokenUPP, frameworkHandle, "DisposeOSLGetMarkTokenUPP")
		registerFunc(&_disposeOSLMarkUPP, frameworkHandle, "DisposeOSLMarkUPP")
		registerFunc(&_disposePtr, frameworkHandle, "DisposePtr")
		registerFunc(&_disposeResErrUPP, frameworkHandle, "DisposeResErrUPP")
		registerFunc(&_disposeSelectorFunctionUPP, frameworkHandle, "DisposeSelectorFunctionUPP")
		registerFunc(&_disposeSleepQUPP, frameworkHandle, "DisposeSleepQUPP")
		registerFunc(&_disposeTextToUnicodeInfo, frameworkHandle, "DisposeTextToUnicodeInfo")
		registerFunc(&_disposeThread, frameworkHandle, "DisposeThread")
		registerFunc(&_disposeThreadEntryUPP, frameworkHandle, "DisposeThreadEntryUPP")
		registerFunc(&_disposeThreadSchedulerUPP, frameworkHandle, "DisposeThreadSchedulerUPP")
		registerFunc(&_disposeThreadSwitchUPP, frameworkHandle, "DisposeThreadSwitchUPP")
		registerFunc(&_disposeThreadTerminationUPP, frameworkHandle, "DisposeThreadTerminationUPP")
		registerFunc(&_disposeTimerUPP, frameworkHandle, "DisposeTimerUPP")
		registerFunc(&_disposeUnicodeToTextFallbackUPP, frameworkHandle, "DisposeUnicodeToTextFallbackUPP")
		registerFunc(&_disposeUnicodeToTextInfo, frameworkHandle, "DisposeUnicodeToTextInfo")
		registerFunc(&_disposeUnicodeToTextRunInfo, frameworkHandle, "DisposeUnicodeToTextRunInfo")
		registerFunc(&_durationToAbsolute, frameworkHandle, "DurationToAbsolute")
		registerFunc(&_durationToNanoseconds, frameworkHandle, "DurationToNanoseconds")
		registerFunc(&_emptyCollection, frameworkHandle, "EmptyCollection")
		registerFunc(&_emptyHandle, frameworkHandle, "EmptyHandle")
		registerFunc(&_enqueue, frameworkHandle, "Enqueue")
		registerFunc(&_fNGetDirectoryForSubscription, frameworkHandle, "FNGetDirectoryForSubscription")
		registerFunc(&_fNNotify, frameworkHandle, "FNNotify")
		registerFunc(&_fNNotifyAll, frameworkHandle, "FNNotifyAll")
		registerFunc(&_fNNotifyByPath, frameworkHandle, "FNNotifyByPath")
		registerFunc(&_fNSubscribe, frameworkHandle, "FNSubscribe")
		registerFunc(&_fNSubscribeByPath, frameworkHandle, "FNSubscribeByPath")
		registerFunc(&_fNUnsubscribe, frameworkHandle, "FNUnsubscribe")
		registerFunc(&_fSAllocateFork, frameworkHandle, "FSAllocateFork")
		registerFunc(&_fSCancelVolumeOperation, frameworkHandle, "FSCancelVolumeOperation")
		registerFunc(&_fSCatalogSearch, frameworkHandle, "FSCatalogSearch")
		registerFunc(&_fSCloseFork, frameworkHandle, "FSCloseFork")
		registerFunc(&_fSCloseIterator, frameworkHandle, "FSCloseIterator")
		registerFunc(&_fSCompareFSRefs, frameworkHandle, "FSCompareFSRefs")
		registerFunc(&_fSCopyDADiskForVolume, frameworkHandle, "FSCopyDADiskForVolume")
		registerFunc(&_fSCopyDiskIDForVolume, frameworkHandle, "FSCopyDiskIDForVolume")
		registerFunc(&_fSCopyObjectAsync, frameworkHandle, "FSCopyObjectAsync")
		registerFunc(&_fSCopyObjectSync, frameworkHandle, "FSCopyObjectSync")
		registerFunc(&_fSCopyURLForVolume, frameworkHandle, "FSCopyURLForVolume")
		registerFunc(&_fSCreateDirectoryUnicode, frameworkHandle, "FSCreateDirectoryUnicode")
		registerFunc(&_fSCreateFileAndOpenForkUnicode, frameworkHandle, "FSCreateFileAndOpenForkUnicode")
		registerFunc(&_fSCreateFileUnicode, frameworkHandle, "FSCreateFileUnicode")
		registerFunc(&_fSCreateFork, frameworkHandle, "FSCreateFork")
		registerFunc(&_fSCreateResFile, frameworkHandle, "FSCreateResFile")
		registerFunc(&_fSCreateResourceFile, frameworkHandle, "FSCreateResourceFile")
		registerFunc(&_fSCreateResourceFork, frameworkHandle, "FSCreateResourceFork")
		registerFunc(&_fSCreateStringFromHFSUniStr, frameworkHandle, "FSCreateStringFromHFSUniStr")
		registerFunc(&_fSCreateVolumeOperation, frameworkHandle, "FSCreateVolumeOperation")
		registerFunc(&_fSDeleteFork, frameworkHandle, "FSDeleteFork")
		registerFunc(&_fSDeleteObject, frameworkHandle, "FSDeleteObject")
		registerFunc(&_fSDetermineIfRefIsEnclosedByFolder, frameworkHandle, "FSDetermineIfRefIsEnclosedByFolder")
		registerFunc(&_fSDisposeVolumeOperation, frameworkHandle, "FSDisposeVolumeOperation")
		registerFunc(&_fSEjectVolumeAsync, frameworkHandle, "FSEjectVolumeAsync")
		registerFunc(&_fSEjectVolumeSync, frameworkHandle, "FSEjectVolumeSync")
		registerFunc(&_fSEventStreamCopyDescription, frameworkHandle, "FSEventStreamCopyDescription")
		registerFunc(&_fSEventStreamCopyPathsBeingWatched, frameworkHandle, "FSEventStreamCopyPathsBeingWatched")
		registerFunc(&_fSEventStreamCreate, frameworkHandle, "FSEventStreamCreate")
		registerFunc(&_fSEventStreamCreateRelativeToDevice, frameworkHandle, "FSEventStreamCreateRelativeToDevice")
		registerFunc(&_fSEventStreamFlushAsync, frameworkHandle, "FSEventStreamFlushAsync")
		registerFunc(&_fSEventStreamFlushSync, frameworkHandle, "FSEventStreamFlushSync")
		registerFunc(&_fSEventStreamGetDeviceBeingWatched, frameworkHandle, "FSEventStreamGetDeviceBeingWatched")
		registerFunc(&_fSEventStreamGetLatestEventId, frameworkHandle, "FSEventStreamGetLatestEventId")
		registerFunc(&_fSEventStreamInvalidate, frameworkHandle, "FSEventStreamInvalidate")
		registerFunc(&_fSEventStreamRelease, frameworkHandle, "FSEventStreamRelease")
		registerFunc(&_fSEventStreamRetain, frameworkHandle, "FSEventStreamRetain")
		registerFunc(&_fSEventStreamSetDispatchQueue, frameworkHandle, "FSEventStreamSetDispatchQueue")
		registerFunc(&_fSEventStreamSetExclusionPaths, frameworkHandle, "FSEventStreamSetExclusionPaths")
		registerFunc(&_fSEventStreamShow, frameworkHandle, "FSEventStreamShow")
		registerFunc(&_fSEventStreamStart, frameworkHandle, "FSEventStreamStart")
		registerFunc(&_fSEventStreamStop, frameworkHandle, "FSEventStreamStop")
		registerFunc(&_fSEventsCopyUUIDForDevice, frameworkHandle, "FSEventsCopyUUIDForDevice")
		registerFunc(&_fSEventsGetCurrentEventId, frameworkHandle, "FSEventsGetCurrentEventId")
		registerFunc(&_fSEventsGetLastEventIdForDeviceBeforeTime, frameworkHandle, "FSEventsGetLastEventIdForDeviceBeforeTime")
		registerFunc(&_fSEventsPurgeEventsForDeviceUpToEventId, frameworkHandle, "FSEventsPurgeEventsForDeviceUpToEventId")
		registerFunc(&_fSFileOperationCancel, frameworkHandle, "FSFileOperationCancel")
		registerFunc(&_fSFileOperationCopyStatus, frameworkHandle, "FSFileOperationCopyStatus")
		registerFunc(&_fSFileOperationCreate, frameworkHandle, "FSFileOperationCreate")
		registerFunc(&_fSFileOperationGetTypeID, frameworkHandle, "FSFileOperationGetTypeID")
		registerFunc(&_fSFileOperationScheduleWithRunLoop, frameworkHandle, "FSFileOperationScheduleWithRunLoop")
		registerFunc(&_fSFileOperationUnscheduleFromRunLoop, frameworkHandle, "FSFileOperationUnscheduleFromRunLoop")
		registerFunc(&_fSFileSecurityCopyAccessControlList, frameworkHandle, "FSFileSecurityCopyAccessControlList")
		registerFunc(&_fSFileSecurityCreate, frameworkHandle, "FSFileSecurityCreate")
		registerFunc(&_fSFileSecurityCreateWithFSPermissionInfo, frameworkHandle, "FSFileSecurityCreateWithFSPermissionInfo")
		registerFunc(&_fSFileSecurityGetGroup, frameworkHandle, "FSFileSecurityGetGroup")
		registerFunc(&_fSFileSecurityGetGroupUUID, frameworkHandle, "FSFileSecurityGetGroupUUID")
		registerFunc(&_fSFileSecurityGetMode, frameworkHandle, "FSFileSecurityGetMode")
		registerFunc(&_fSFileSecurityGetOwner, frameworkHandle, "FSFileSecurityGetOwner")
		registerFunc(&_fSFileSecurityGetOwnerUUID, frameworkHandle, "FSFileSecurityGetOwnerUUID")
		registerFunc(&_fSFileSecurityGetTypeID, frameworkHandle, "FSFileSecurityGetTypeID")
		registerFunc(&_fSFileSecurityRefCreateCopy, frameworkHandle, "FSFileSecurityRefCreateCopy")
		registerFunc(&_fSFileSecuritySetAccessControlList, frameworkHandle, "FSFileSecuritySetAccessControlList")
		registerFunc(&_fSFileSecuritySetGroup, frameworkHandle, "FSFileSecuritySetGroup")
		registerFunc(&_fSFileSecuritySetGroupUUID, frameworkHandle, "FSFileSecuritySetGroupUUID")
		registerFunc(&_fSFileSecuritySetMode, frameworkHandle, "FSFileSecuritySetMode")
		registerFunc(&_fSFileSecuritySetOwner, frameworkHandle, "FSFileSecuritySetOwner")
		registerFunc(&_fSFileSecuritySetOwnerUUID, frameworkHandle, "FSFileSecuritySetOwnerUUID")
		registerFunc(&_fSFindFolder, frameworkHandle, "FSFindFolder")
		registerFunc(&_fSFlushFork, frameworkHandle, "FSFlushFork")
		registerFunc(&_fSFlushVolume, frameworkHandle, "FSFlushVolume")
		registerFunc(&_fSGetAsyncEjectStatus, frameworkHandle, "FSGetAsyncEjectStatus")
		registerFunc(&_fSGetAsyncMountStatus, frameworkHandle, "FSGetAsyncMountStatus")
		registerFunc(&_fSGetAsyncUnmountStatus, frameworkHandle, "FSGetAsyncUnmountStatus")
		registerFunc(&_fSGetDataForkName, frameworkHandle, "FSGetDataForkName")
		registerFunc(&_fSGetForkCBInfo, frameworkHandle, "FSGetForkCBInfo")
		registerFunc(&_fSGetForkPosition, frameworkHandle, "FSGetForkPosition")
		registerFunc(&_fSGetForkSize, frameworkHandle, "FSGetForkSize")
		registerFunc(&_fSGetHFSUniStrFromString, frameworkHandle, "FSGetHFSUniStrFromString")
		registerFunc(&_fSGetResourceForkName, frameworkHandle, "FSGetResourceForkName")
		registerFunc(&_fSGetTemporaryDirectoryForReplaceObject, frameworkHandle, "FSGetTemporaryDirectoryForReplaceObject")
		registerFunc(&_fSGetVolumeForDADisk, frameworkHandle, "FSGetVolumeForDADisk")
		registerFunc(&_fSGetVolumeForDiskID, frameworkHandle, "FSGetVolumeForDiskID")
		registerFunc(&_fSGetVolumeInfo, frameworkHandle, "FSGetVolumeInfo")
		registerFunc(&_fSGetVolumeMountInfo, frameworkHandle, "FSGetVolumeMountInfo")
		registerFunc(&_fSGetVolumeMountInfoSize, frameworkHandle, "FSGetVolumeMountInfoSize")
		registerFunc(&_fSGetVolumeParms, frameworkHandle, "FSGetVolumeParms")
		registerFunc(&_fSIsFSRefValid, frameworkHandle, "FSIsFSRefValid")
		registerFunc(&_fSIterateForks, frameworkHandle, "FSIterateForks")
		registerFunc(&_fSLockRange, frameworkHandle, "FSLockRange")
		registerFunc(&_fSMakeFSRefUnicode, frameworkHandle, "FSMakeFSRefUnicode")
		registerFunc(&_fSMountLocalVolumeAsync, frameworkHandle, "FSMountLocalVolumeAsync")
		registerFunc(&_fSMountLocalVolumeSync, frameworkHandle, "FSMountLocalVolumeSync")
		registerFunc(&_fSMountServerVolumeAsync, frameworkHandle, "FSMountServerVolumeAsync")
		registerFunc(&_fSMountServerVolumeSync, frameworkHandle, "FSMountServerVolumeSync")
		registerFunc(&_fSMoveObject, frameworkHandle, "FSMoveObject")
		registerFunc(&_fSMoveObjectAsync, frameworkHandle, "FSMoveObjectAsync")
		registerFunc(&_fSMoveObjectSync, frameworkHandle, "FSMoveObjectSync")
		registerFunc(&_fSMoveObjectToTrashAsync, frameworkHandle, "FSMoveObjectToTrashAsync")
		registerFunc(&_fSMoveObjectToTrashSync, frameworkHandle, "FSMoveObjectToTrashSync")
		registerFunc(&_fSOpenFork, frameworkHandle, "FSOpenFork")
		registerFunc(&_fSOpenIterator, frameworkHandle, "FSOpenIterator")
		registerFunc(&_fSOpenOrphanResFile, frameworkHandle, "FSOpenOrphanResFile")
		registerFunc(&_fSOpenResFile, frameworkHandle, "FSOpenResFile")
		registerFunc(&_fSOpenResourceFile, frameworkHandle, "FSOpenResourceFile")
		registerFunc(&_fSPathCopyObjectAsync, frameworkHandle, "FSPathCopyObjectAsync")
		registerFunc(&_fSPathCopyObjectSync, frameworkHandle, "FSPathCopyObjectSync")
		registerFunc(&_fSPathFileOperationCopyStatus, frameworkHandle, "FSPathFileOperationCopyStatus")
		registerFunc(&_fSPathGetTemporaryDirectoryForReplaceObject, frameworkHandle, "FSPathGetTemporaryDirectoryForReplaceObject")
		registerFunc(&_fSPathMakeRef, frameworkHandle, "FSPathMakeRef")
		registerFunc(&_fSPathMakeRefWithOptions, frameworkHandle, "FSPathMakeRefWithOptions")
		registerFunc(&_fSPathMoveObjectAsync, frameworkHandle, "FSPathMoveObjectAsync")
		registerFunc(&_fSPathMoveObjectSync, frameworkHandle, "FSPathMoveObjectSync")
		registerFunc(&_fSPathMoveObjectToTrashAsync, frameworkHandle, "FSPathMoveObjectToTrashAsync")
		registerFunc(&_fSPathMoveObjectToTrashSync, frameworkHandle, "FSPathMoveObjectToTrashSync")
		registerFunc(&_fSPathReplaceObject, frameworkHandle, "FSPathReplaceObject")
		registerFunc(&_fSReadFork, frameworkHandle, "FSReadFork")
		registerFunc(&_fSRefMakePath, frameworkHandle, "FSRefMakePath")
		registerFunc(&_fSRenameUnicode, frameworkHandle, "FSRenameUnicode")
		registerFunc(&_fSReplaceObject, frameworkHandle, "FSReplaceObject")
		registerFunc(&_fSResolveNodeID, frameworkHandle, "FSResolveNodeID")
		registerFunc(&_fSResourceFileAlreadyOpen, frameworkHandle, "FSResourceFileAlreadyOpen")
		registerFunc(&_fSSetCatalogInfo, frameworkHandle, "FSSetCatalogInfo")
		registerFunc(&_fSSetForkPosition, frameworkHandle, "FSSetForkPosition")
		registerFunc(&_fSSetForkSize, frameworkHandle, "FSSetForkSize")
		registerFunc(&_fSSetVolumeInfo, frameworkHandle, "FSSetVolumeInfo")
		registerFunc(&_fSUnlinkObject, frameworkHandle, "FSUnlinkObject")
		registerFunc(&_fSUnlockRange, frameworkHandle, "FSUnlockRange")
		registerFunc(&_fSUnmountVolumeAsync, frameworkHandle, "FSUnmountVolumeAsync")
		registerFunc(&_fSUnmountVolumeSync, frameworkHandle, "FSUnmountVolumeSync")
		registerFunc(&_fSVolumeMount, frameworkHandle, "FSVolumeMount")
		registerFunc(&_fSWriteFork, frameworkHandle, "FSWriteFork")
		registerFunc(&_findFolder, frameworkHandle, "FindFolder")
		registerFunc(&_findNextComponent, frameworkHandle, "FindNextComponent")
		registerFunc(&_fix2Frac, frameworkHandle, "Fix2Frac")
		registerFunc(&_fix2Long, frameworkHandle, "Fix2Long")
		registerFunc(&_fix2X, frameworkHandle, "Fix2X")
		registerFunc(&_fixATan2, frameworkHandle, "FixATan2")
		registerFunc(&_fixDiv, frameworkHandle, "FixDiv")
		registerFunc(&_fixMul, frameworkHandle, "FixMul")
		registerFunc(&_fixRatio, frameworkHandle, "FixRatio")
		registerFunc(&_fixRound, frameworkHandle, "FixRound")
		registerFunc(&_flattenCollection, frameworkHandle, "FlattenCollection")
		registerFunc(&_flattenCollectionToHdl, frameworkHandle, "FlattenCollectionToHdl")
		registerFunc(&_flattenPartialCollection, frameworkHandle, "FlattenPartialCollection")
		registerFunc(&_frac2Fix, frameworkHandle, "Frac2Fix")
		registerFunc(&_frac2X, frameworkHandle, "Frac2X")
		registerFunc(&_fracCos, frameworkHandle, "FracCos")
		registerFunc(&_fracDiv, frameworkHandle, "FracDiv")
		registerFunc(&_fracMul, frameworkHandle, "FracMul")
		registerFunc(&_fracSin, frameworkHandle, "FracSin")
		registerFunc(&_fracSqrt, frameworkHandle, "FracSqrt")
		registerFunc(&_gestaltFunc, frameworkHandle, "Gestalt")
		registerFunc(&_get1IndResource, frameworkHandle, "Get1IndResource")
		registerFunc(&_get1IndType, frameworkHandle, "Get1IndType")
		registerFunc(&_get1NamedResource, frameworkHandle, "Get1NamedResource")
		registerFunc(&_get1Resource, frameworkHandle, "Get1Resource")
		registerFunc(&_getCPUSpeed, frameworkHandle, "GetCPUSpeed")
		registerFunc(&_getCollectionDefaultAttributes, frameworkHandle, "GetCollectionDefaultAttributes")
		registerFunc(&_getCollectionExceptionProc, frameworkHandle, "GetCollectionExceptionProc")
		registerFunc(&_getCollectionItem, frameworkHandle, "GetCollectionItem")
		registerFunc(&_getCollectionItemHdl, frameworkHandle, "GetCollectionItemHdl")
		registerFunc(&_getCollectionItemInfo, frameworkHandle, "GetCollectionItemInfo")
		registerFunc(&_getCollectionRetainCount, frameworkHandle, "GetCollectionRetainCount")
		registerFunc(&_getComponentIndString, frameworkHandle, "GetComponentIndString")
		registerFunc(&_getComponentInfo, frameworkHandle, "GetComponentInfo")
		registerFunc(&_getComponentInstanceError, frameworkHandle, "GetComponentInstanceError")
		registerFunc(&_getComponentInstanceStorage, frameworkHandle, "GetComponentInstanceStorage")
		registerFunc(&_getComponentListModSeed, frameworkHandle, "GetComponentListModSeed")
		registerFunc(&_getComponentPublicIndString, frameworkHandle, "GetComponentPublicIndString")
		registerFunc(&_getComponentPublicResource, frameworkHandle, "GetComponentPublicResource")
		registerFunc(&_getComponentPublicResourceList, frameworkHandle, "GetComponentPublicResourceList")
		registerFunc(&_getComponentRefcon, frameworkHandle, "GetComponentRefcon")
		registerFunc(&_getComponentResource, frameworkHandle, "GetComponentResource")
		registerFunc(&_getComponentTypeModSeed, frameworkHandle, "GetComponentTypeModSeed")
		registerFunc(&_getCurrentThread, frameworkHandle, "GetCurrentThread")
		registerFunc(&_getCustomIconsEnabled, frameworkHandle, "GetCustomIconsEnabled")
		registerFunc(&_getDebugComponentInfo, frameworkHandle, "GetDebugComponentInfo")
		registerFunc(&_getDebugOptionInfo, frameworkHandle, "GetDebugOptionInfo")
		registerFunc(&_getDefaultThreadStackSize, frameworkHandle, "GetDefaultThreadStackSize")
		registerFunc(&_getFolderNameUnicode, frameworkHandle, "GetFolderNameUnicode")
		registerFunc(&_getFolderTypes, frameworkHandle, "GetFolderTypes")
		registerFunc(&_getHandleSize, frameworkHandle, "GetHandleSize")
		registerFunc(&_getIconRef, frameworkHandle, "GetIconRef")
		registerFunc(&_getIconRefFromComponent, frameworkHandle, "GetIconRefFromComponent")
		registerFunc(&_getIconRefFromFileInfo, frameworkHandle, "GetIconRefFromFileInfo")
		registerFunc(&_getIconRefFromFolder, frameworkHandle, "GetIconRefFromFolder")
		registerFunc(&_getIconRefFromIconFamilyPtr, frameworkHandle, "GetIconRefFromIconFamilyPtr")
		registerFunc(&_getIconRefFromTypeInfo, frameworkHandle, "GetIconRefFromTypeInfo")
		registerFunc(&_getIconRefOwners, frameworkHandle, "GetIconRefOwners")
		registerFunc(&_getIndResource, frameworkHandle, "GetIndResource")
		registerFunc(&_getIndType, frameworkHandle, "GetIndType")
		registerFunc(&_getIndexedCollectionItem, frameworkHandle, "GetIndexedCollectionItem")
		registerFunc(&_getIndexedCollectionItemHdl, frameworkHandle, "GetIndexedCollectionItemHdl")
		registerFunc(&_getIndexedCollectionItemInfo, frameworkHandle, "GetIndexedCollectionItemInfo")
		registerFunc(&_getIndexedCollectionTag, frameworkHandle, "GetIndexedCollectionTag")
		registerFunc(&_getMacOSStatusCommentString, frameworkHandle, "GetMacOSStatusCommentString")
		registerFunc(&_getMacOSStatusErrorString, frameworkHandle, "GetMacOSStatusErrorString")
		registerFunc(&_getMaxResourceSize, frameworkHandle, "GetMaxResourceSize")
		registerFunc(&_getNamedResource, frameworkHandle, "GetNamedResource")
		registerFunc(&_getNewCollection, frameworkHandle, "GetNewCollection")
		registerFunc(&_getNextFOND, frameworkHandle, "GetNextFOND")
		registerFunc(&_getNextResourceFile, frameworkHandle, "GetNextResourceFile")
		registerFunc(&_getPtrSize, frameworkHandle, "GetPtrSize")
		registerFunc(&_getResAttrs, frameworkHandle, "GetResAttrs")
		registerFunc(&_getResFileAttrs, frameworkHandle, "GetResFileAttrs")
		registerFunc(&_getResInfo, frameworkHandle, "GetResInfo")
		registerFunc(&_getResource, frameworkHandle, "GetResource")
		registerFunc(&_getResourceSizeOnDisk, frameworkHandle, "GetResourceSizeOnDisk")
		registerFunc(&_getScriptInfoFromTextEncoding, frameworkHandle, "GetScriptInfoFromTextEncoding")
		registerFunc(&_getScriptManagerVariable, frameworkHandle, "GetScriptManagerVariable")
		registerFunc(&_getTaggedCollectionItem, frameworkHandle, "GetTaggedCollectionItem")
		registerFunc(&_getTaggedCollectionItemInfo, frameworkHandle, "GetTaggedCollectionItemInfo")
		registerFunc(&_getTextEncodingBase, frameworkHandle, "GetTextEncodingBase")
		registerFunc(&_getTextEncodingFormat, frameworkHandle, "GetTextEncodingFormat")
		registerFunc(&_getTextEncodingFromScriptInfo, frameworkHandle, "GetTextEncodingFromScriptInfo")
		registerFunc(&_getTextEncodingName, frameworkHandle, "GetTextEncodingName")
		registerFunc(&_getTextEncodingVariant, frameworkHandle, "GetTextEncodingVariant")
		registerFunc(&_getThreadCurrentTaskRef, frameworkHandle, "GetThreadCurrentTaskRef")
		registerFunc(&_getThreadState, frameworkHandle, "GetThreadState")
		registerFunc(&_getThreadStateGivenTaskRef, frameworkHandle, "GetThreadStateGivenTaskRef")
		registerFunc(&_getTopResourceFile, frameworkHandle, "GetTopResourceFile")
		registerFunc(&_hClrRBit, frameworkHandle, "HClrRBit")
		registerFunc(&_hGetState, frameworkHandle, "HGetState")
		registerFunc(&_hLock, frameworkHandle, "HLock")
		registerFunc(&_hLockHi, frameworkHandle, "HLockHi")
		registerFunc(&_hSetRBit, frameworkHandle, "HSetRBit")
		registerFunc(&_hSetState, frameworkHandle, "HSetState")
		registerFunc(&_hUnlock, frameworkHandle, "HUnlock")
		registerFunc(&_handAndHand, frameworkHandle, "HandAndHand")
		registerFunc(&_handToHand, frameworkHandle, "HandToHand")
		registerFunc(&_homeResFile, frameworkHandle, "HomeResFile")
		registerFunc(&_identifyFolder, frameworkHandle, "IdentifyFolder")
		registerFunc(&_incrementAtomic, frameworkHandle, "IncrementAtomic")
		registerFunc(&_incrementAtomic16, frameworkHandle, "IncrementAtomic16")
		registerFunc(&_incrementAtomic8, frameworkHandle, "IncrementAtomic8")
		registerFunc(&_insTime, frameworkHandle, "InsTime")
		registerFunc(&_insXTime, frameworkHandle, "InsXTime")
		registerFunc(&_insertResourceFile, frameworkHandle, "InsertResourceFile")
		registerFunc(&_installDebugAssertOutputHandler, frameworkHandle, "InstallDebugAssertOutputHandler")
		registerFunc(&_installExceptionHandler, frameworkHandle, "InstallExceptionHandler")
		registerFunc(&_installTimeTask, frameworkHandle, "InstallTimeTask")
		registerFunc(&_installXTimeTask, frameworkHandle, "InstallXTimeTask")
		registerFunc(&_invalidateFolderDescriptorCache, frameworkHandle, "InvalidateFolderDescriptorCache")
		registerFunc(&_invokeAECoerceDescUPP, frameworkHandle, "InvokeAECoerceDescUPP")
		registerFunc(&_invokeAECoercePtrUPP, frameworkHandle, "InvokeAECoercePtrUPP")
		registerFunc(&_invokeAEDisposeExternalUPP, frameworkHandle, "InvokeAEDisposeExternalUPP")
		registerFunc(&_invokeAEEventHandlerUPP, frameworkHandle, "InvokeAEEventHandlerUPP")
		registerFunc(&_invokeCollectionExceptionUPP, frameworkHandle, "InvokeCollectionExceptionUPP")
		registerFunc(&_invokeCollectionFlattenUPP, frameworkHandle, "InvokeCollectionFlattenUPP")
		registerFunc(&_invokeComponentMPWorkFunctionUPP, frameworkHandle, "InvokeComponentMPWorkFunctionUPP")
		registerFunc(&_invokeComponentRoutineUPP, frameworkHandle, "InvokeComponentRoutineUPP")
		registerFunc(&_invokeDebugAssertOutputHandlerUPP, frameworkHandle, "InvokeDebugAssertOutputHandlerUPP")
		registerFunc(&_invokeDebugComponentCallbackUPP, frameworkHandle, "InvokeDebugComponentCallbackUPP")
		registerFunc(&_invokeDebuggerDisposeThreadUPP, frameworkHandle, "InvokeDebuggerDisposeThreadUPP")
		registerFunc(&_invokeDebuggerNewThreadUPP, frameworkHandle, "InvokeDebuggerNewThreadUPP")
		registerFunc(&_invokeDebuggerThreadSchedulerUPP, frameworkHandle, "InvokeDebuggerThreadSchedulerUPP")
		registerFunc(&_invokeDeferredTaskUPP, frameworkHandle, "InvokeDeferredTaskUPP")
		registerFunc(&_invokeExceptionHandlerUPP, frameworkHandle, "InvokeExceptionHandlerUPP")
		registerFunc(&_invokeFNSubscriptionUPP, frameworkHandle, "InvokeFNSubscriptionUPP")
		registerFunc(&_invokeFSVolumeEjectUPP, frameworkHandle, "InvokeFSVolumeEjectUPP")
		registerFunc(&_invokeFSVolumeMountUPP, frameworkHandle, "InvokeFSVolumeMountUPP")
		registerFunc(&_invokeFSVolumeUnmountUPP, frameworkHandle, "InvokeFSVolumeUnmountUPP")
		registerFunc(&_invokeFolderManagerNotificationUPP, frameworkHandle, "InvokeFolderManagerNotificationUPP")
		registerFunc(&_invokeGetMissingComponentResourceUPP, frameworkHandle, "InvokeGetMissingComponentResourceUPP")
		registerFunc(&_invokeIOCompletionUPP, frameworkHandle, "InvokeIOCompletionUPP")
		registerFunc(&_invokeIndexToUCStringUPP, frameworkHandle, "InvokeIndexToUCStringUPP")
		registerFunc(&_invokeKCCallbackUPP, frameworkHandle, "InvokeKCCallbackUPP")
		registerFunc(&_invokeOSLAccessorUPP, frameworkHandle, "InvokeOSLAccessorUPP")
		registerFunc(&_invokeOSLAdjustMarksUPP, frameworkHandle, "InvokeOSLAdjustMarksUPP")
		registerFunc(&_invokeOSLCompareUPP, frameworkHandle, "InvokeOSLCompareUPP")
		registerFunc(&_invokeOSLCountUPP, frameworkHandle, "InvokeOSLCountUPP")
		registerFunc(&_invokeOSLDisposeTokenUPP, frameworkHandle, "InvokeOSLDisposeTokenUPP")
		registerFunc(&_invokeOSLGetErrDescUPP, frameworkHandle, "InvokeOSLGetErrDescUPP")
		registerFunc(&_invokeOSLGetMarkTokenUPP, frameworkHandle, "InvokeOSLGetMarkTokenUPP")
		registerFunc(&_invokeOSLMarkUPP, frameworkHandle, "InvokeOSLMarkUPP")
		registerFunc(&_invokeResErrUPP, frameworkHandle, "InvokeResErrUPP")
		registerFunc(&_invokeSelectorFunctionUPP, frameworkHandle, "InvokeSelectorFunctionUPP")
		registerFunc(&_invokeSleepQUPP, frameworkHandle, "InvokeSleepQUPP")
		registerFunc(&_invokeThreadEntryUPP, frameworkHandle, "InvokeThreadEntryUPP")
		registerFunc(&_invokeThreadSchedulerUPP, frameworkHandle, "InvokeThreadSchedulerUPP")
		registerFunc(&_invokeThreadSwitchUPP, frameworkHandle, "InvokeThreadSwitchUPP")
		registerFunc(&_invokeThreadTerminationUPP, frameworkHandle, "InvokeThreadTerminationUPP")
		registerFunc(&_invokeTimerUPP, frameworkHandle, "InvokeTimerUPP")
		registerFunc(&_invokeUnicodeToTextFallbackUPP, frameworkHandle, "InvokeUnicodeToTextFallbackUPP")
		registerFunc(&_isDataAvailableInIconRef, frameworkHandle, "IsDataAvailableInIconRef")
		registerFunc(&_isHandleValid, frameworkHandle, "IsHandleValid")
		registerFunc(&_isHeapValid, frameworkHandle, "IsHeapValid")
		registerFunc(&_isIconRefComposite, frameworkHandle, "IsIconRefComposite")
		registerFunc(&_isMetric, frameworkHandle, "IsMetric")
		registerFunc(&_isPointerValid, frameworkHandle, "IsPointerValid")
		registerFunc(&_isValidIconRef, frameworkHandle, "IsValidIconRef")
		registerFunc(&_kCAddCallback, frameworkHandle, "KCAddCallback")
		registerFunc(&_kCCopyItem, frameworkHandle, "KCCopyItem")
		registerFunc(&_kCCountKeychains, frameworkHandle, "KCCountKeychains")
		registerFunc(&_kCDeleteItem, frameworkHandle, "KCDeleteItem")
		registerFunc(&_kCFindAppleSharePassword, frameworkHandle, "KCFindAppleSharePassword")
		registerFunc(&_kCFindFirstItem, frameworkHandle, "KCFindFirstItem")
		registerFunc(&_kCFindGenericPassword, frameworkHandle, "KCFindGenericPassword")
		registerFunc(&_kCFindInternetPassword, frameworkHandle, "KCFindInternetPassword")
		registerFunc(&_kCFindInternetPasswordWithPath, frameworkHandle, "KCFindInternetPasswordWithPath")
		registerFunc(&_kCFindNextItem, frameworkHandle, "KCFindNextItem")
		registerFunc(&_kCGetAttribute, frameworkHandle, "KCGetAttribute")
		registerFunc(&_kCGetData, frameworkHandle, "KCGetData")
		registerFunc(&_kCGetDefaultKeychain, frameworkHandle, "KCGetDefaultKeychain")
		registerFunc(&_kCGetIndKeychain, frameworkHandle, "KCGetIndKeychain")
		registerFunc(&_kCGetKeychain, frameworkHandle, "KCGetKeychain")
		registerFunc(&_kCGetKeychainManagerVersion, frameworkHandle, "KCGetKeychainManagerVersion")
		registerFunc(&_kCGetKeychainName, frameworkHandle, "KCGetKeychainName")
		registerFunc(&_kCGetStatus, frameworkHandle, "KCGetStatus")
		registerFunc(&_kCIsInteractionAllowed, frameworkHandle, "KCIsInteractionAllowed")
		registerFunc(&_kCLock, frameworkHandle, "KCLock")
		registerFunc(&_kCMakeAliasFromKCRef, frameworkHandle, "KCMakeAliasFromKCRef")
		registerFunc(&_kCMakeKCRefFromAlias, frameworkHandle, "KCMakeKCRefFromAlias")
		registerFunc(&_kCMakeKCRefFromFSRef, frameworkHandle, "KCMakeKCRefFromFSRef")
		registerFunc(&_kCNewItem, frameworkHandle, "KCNewItem")
		registerFunc(&_kCReleaseItem, frameworkHandle, "KCReleaseItem")
		registerFunc(&_kCReleaseKeychain, frameworkHandle, "KCReleaseKeychain")
		registerFunc(&_kCReleaseSearch, frameworkHandle, "KCReleaseSearch")
		registerFunc(&_kCRemoveCallback, frameworkHandle, "KCRemoveCallback")
		registerFunc(&_kCSetAttribute, frameworkHandle, "KCSetAttribute")
		registerFunc(&_kCSetData, frameworkHandle, "KCSetData")
		registerFunc(&_kCSetDefaultKeychain, frameworkHandle, "KCSetDefaultKeychain")
		registerFunc(&_kCSetInteractionAllowed, frameworkHandle, "KCSetInteractionAllowed")
		registerFunc(&_kCUpdateItem, frameworkHandle, "KCUpdateItem")
		registerFunc(&_lMGetApFontID, frameworkHandle, "LMGetApFontID")
		registerFunc(&_lMGetBootDrive, frameworkHandle, "LMGetBootDrive")
		registerFunc(&_lMGetIntlSpec, frameworkHandle, "LMGetIntlSpec")
		registerFunc(&_lMGetMemErr, frameworkHandle, "LMGetMemErr")
		registerFunc(&_lMGetResErr, frameworkHandle, "LMGetResErr")
		registerFunc(&_lMGetResLoad, frameworkHandle, "LMGetResLoad")
		registerFunc(&_lMGetSysFontSize, frameworkHandle, "LMGetSysFontSize")
		registerFunc(&_lMGetSysMap, frameworkHandle, "LMGetSysMap")
		registerFunc(&_lMGetTmpResLoad, frameworkHandle, "LMGetTmpResLoad")
		registerFunc(&_lMSetApFontID, frameworkHandle, "LMSetApFontID")
		registerFunc(&_lMSetBootDrive, frameworkHandle, "LMSetBootDrive")
		registerFunc(&_lMSetIntlSpec, frameworkHandle, "LMSetIntlSpec")
		registerFunc(&_lMSetMemErr, frameworkHandle, "LMSetMemErr")
		registerFunc(&_lMSetResErr, frameworkHandle, "LMSetResErr")
		registerFunc(&_lMSetResLoad, frameworkHandle, "LMSetResLoad")
		registerFunc(&_lMSetSysFontFam, frameworkHandle, "LMSetSysFontFam")
		registerFunc(&_lMSetSysFontSize, frameworkHandle, "LMSetSysFontSize")
		registerFunc(&_lMSetSysMap, frameworkHandle, "LMSetSysMap")
		registerFunc(&_lMSetTmpResLoad, frameworkHandle, "LMSetTmpResLoad")
		registerFunc(&_lSCanURLAcceptURL, frameworkHandle, "LSCanURLAcceptURL")
		registerFunc(&_lSCopyAllHandlersForURLScheme, frameworkHandle, "LSCopyAllHandlersForURLScheme")
		registerFunc(&_lSCopyAllRoleHandlersForContentType, frameworkHandle, "LSCopyAllRoleHandlersForContentType")
		registerFunc(&_lSCopyApplicationURLsForBundleIdentifier, frameworkHandle, "LSCopyApplicationURLsForBundleIdentifier")
		registerFunc(&_lSCopyApplicationURLsForURL, frameworkHandle, "LSCopyApplicationURLsForURL")
		registerFunc(&_lSCopyDefaultApplicationURLForContentType, frameworkHandle, "LSCopyDefaultApplicationURLForContentType")
		registerFunc(&_lSCopyDefaultApplicationURLForURL, frameworkHandle, "LSCopyDefaultApplicationURLForURL")
		registerFunc(&_lSCopyDefaultHandlerForURLScheme, frameworkHandle, "LSCopyDefaultHandlerForURLScheme")
		registerFunc(&_lSCopyDefaultRoleHandlerForContentType, frameworkHandle, "LSCopyDefaultRoleHandlerForContentType")
		registerFunc(&_lSCopyDisplayNameForURL, frameworkHandle, "LSCopyDisplayNameForURL")
		registerFunc(&_lSCopyItemInfoForURL, frameworkHandle, "LSCopyItemInfoForURL")
		registerFunc(&_lSCopyKindStringForURL, frameworkHandle, "LSCopyKindStringForURL")
		registerFunc(&_lSGetExtensionInfo, frameworkHandle, "LSGetExtensionInfo")
		registerFunc(&_lSGetHandlerOptionsForContentType, frameworkHandle, "LSGetHandlerOptionsForContentType")
		registerFunc(&_lSOpenApplication, frameworkHandle, "LSOpenApplication")
		registerFunc(&_lSOpenCFURLRef, frameworkHandle, "LSOpenCFURLRef")
		registerFunc(&_lSOpenFSRef, frameworkHandle, "LSOpenFSRef")
		registerFunc(&_lSOpenFromRefSpec, frameworkHandle, "LSOpenFromRefSpec")
		registerFunc(&_lSOpenFromURLSpec, frameworkHandle, "LSOpenFromURLSpec")
		registerFunc(&_lSOpenItemsWithRole, frameworkHandle, "LSOpenItemsWithRole")
		registerFunc(&_lSOpenURLsWithRole, frameworkHandle, "LSOpenURLsWithRole")
		registerFunc(&_lSRegisterURL, frameworkHandle, "LSRegisterURL")
		registerFunc(&_lSSetDefaultHandlerForURLScheme, frameworkHandle, "LSSetDefaultHandlerForURLScheme")
		registerFunc(&_lSSetDefaultRoleHandlerForContentType, frameworkHandle, "LSSetDefaultRoleHandlerForContentType")
		registerFunc(&_lSSetExtensionHiddenForURL, frameworkHandle, "LSSetExtensionHiddenForURL")
		registerFunc(&_lSSetHandlerOptionsForContentType, frameworkHandle, "LSSetHandlerOptionsForContentType")
		registerFunc(&_lSSetItemAttribute, frameworkHandle, "LSSetItemAttribute")
		registerFunc(&_lSSharedFileListAddObserver, frameworkHandle, "LSSharedFileListAddObserver")
		registerFunc(&_lSSharedFileListCopyProperty, frameworkHandle, "LSSharedFileListCopyProperty")
		registerFunc(&_lSSharedFileListCopySnapshot, frameworkHandle, "LSSharedFileListCopySnapshot")
		registerFunc(&_lSSharedFileListCreate, frameworkHandle, "LSSharedFileListCreate")
		registerFunc(&_lSSharedFileListGetSeedValue, frameworkHandle, "LSSharedFileListGetSeedValue")
		registerFunc(&_lSSharedFileListGetTypeID, frameworkHandle, "LSSharedFileListGetTypeID")
		registerFunc(&_lSSharedFileListInsertItemFSRef, frameworkHandle, "LSSharedFileListInsertItemFSRef")
		registerFunc(&_lSSharedFileListInsertItemURL, frameworkHandle, "LSSharedFileListInsertItemURL")
		registerFunc(&_lSSharedFileListItemCopyDisplayName, frameworkHandle, "LSSharedFileListItemCopyDisplayName")
		registerFunc(&_lSSharedFileListItemCopyIconRef, frameworkHandle, "LSSharedFileListItemCopyIconRef")
		registerFunc(&_lSSharedFileListItemCopyProperty, frameworkHandle, "LSSharedFileListItemCopyProperty")
		registerFunc(&_lSSharedFileListItemCopyResolvedURL, frameworkHandle, "LSSharedFileListItemCopyResolvedURL")
		registerFunc(&_lSSharedFileListItemGetID, frameworkHandle, "LSSharedFileListItemGetID")
		registerFunc(&_lSSharedFileListItemGetTypeID, frameworkHandle, "LSSharedFileListItemGetTypeID")
		registerFunc(&_lSSharedFileListItemMove, frameworkHandle, "LSSharedFileListItemMove")
		registerFunc(&_lSSharedFileListItemRemove, frameworkHandle, "LSSharedFileListItemRemove")
		registerFunc(&_lSSharedFileListItemResolve, frameworkHandle, "LSSharedFileListItemResolve")
		registerFunc(&_lSSharedFileListItemSetProperty, frameworkHandle, "LSSharedFileListItemSetProperty")
		registerFunc(&_lSSharedFileListRemoveAllItems, frameworkHandle, "LSSharedFileListRemoveAllItems")
		registerFunc(&_lSSharedFileListRemoveObserver, frameworkHandle, "LSSharedFileListRemoveObserver")
		registerFunc(&_lSSharedFileListSetAuthorization, frameworkHandle, "LSSharedFileListSetAuthorization")
		registerFunc(&_lSSharedFileListSetProperty, frameworkHandle, "LSSharedFileListSetProperty")
		registerFunc(&_loadResource, frameworkHandle, "LoadResource")
		registerFunc(&_localeCountNames, frameworkHandle, "LocaleCountNames")
		registerFunc(&_localeGetIndName, frameworkHandle, "LocaleGetIndName")
		registerFunc(&_localeGetName, frameworkHandle, "LocaleGetName")
		registerFunc(&_localeOperationCountLocales, frameworkHandle, "LocaleOperationCountLocales")
		registerFunc(&_localeOperationCountNames, frameworkHandle, "LocaleOperationCountNames")
		registerFunc(&_localeOperationGetIndName, frameworkHandle, "LocaleOperationGetIndName")
		registerFunc(&_localeOperationGetLocales, frameworkHandle, "LocaleOperationGetLocales")
		registerFunc(&_localeOperationGetName, frameworkHandle, "LocaleOperationGetName")
		registerFunc(&_localeRefFromLangOrRegionCode, frameworkHandle, "LocaleRefFromLangOrRegionCode")
		registerFunc(&_localeRefFromLocaleString, frameworkHandle, "LocaleRefFromLocaleString")
		registerFunc(&_localeRefGetPartString, frameworkHandle, "LocaleRefGetPartString")
		registerFunc(&_localeStringToLangAndRegionCodes, frameworkHandle, "LocaleStringToLangAndRegionCodes")
		registerFunc(&_long2Fix, frameworkHandle, "Long2Fix")
		registerFunc(&_longDoubleToSInt64, frameworkHandle, "LongDoubleToSInt64")
		registerFunc(&_longDoubleToUInt64, frameworkHandle, "LongDoubleToUInt64")
		registerFunc(&_mDCopyLabelKinds, frameworkHandle, "MDCopyLabelKinds")
		registerFunc(&_mDCopyLabelWithUUID, frameworkHandle, "MDCopyLabelWithUUID")
		registerFunc(&_mDCopyLabelsMatchingExpression, frameworkHandle, "MDCopyLabelsMatchingExpression")
		registerFunc(&_mDCopyLabelsWithKind, frameworkHandle, "MDCopyLabelsWithKind")
		registerFunc(&_mDItemCopyAttribute, frameworkHandle, "MDItemCopyAttribute")
		registerFunc(&_mDItemCopyAttributeList, frameworkHandle, "MDItemCopyAttributeList")
		registerFunc(&_mDItemCopyAttributeNames, frameworkHandle, "MDItemCopyAttributeNames")
		registerFunc(&_mDItemCopyAttributes, frameworkHandle, "MDItemCopyAttributes")
		registerFunc(&_mDItemCopyLabels, frameworkHandle, "MDItemCopyLabels")
		registerFunc(&_mDItemCreate, frameworkHandle, "MDItemCreate")
		registerFunc(&_mDItemCreateWithURL, frameworkHandle, "MDItemCreateWithURL")
		registerFunc(&_mDItemGetCacheFileDescriptors, frameworkHandle, "MDItemGetCacheFileDescriptors")
		registerFunc(&_mDItemGetTypeID, frameworkHandle, "MDItemGetTypeID")
		registerFunc(&_mDItemRemoveLabel, frameworkHandle, "MDItemRemoveLabel")
		registerFunc(&_mDItemSetLabel, frameworkHandle, "MDItemSetLabel")
		registerFunc(&_mDItemsCopyAttributes, frameworkHandle, "MDItemsCopyAttributes")
		registerFunc(&_mDItemsCreateWithURLs, frameworkHandle, "MDItemsCreateWithURLs")
		registerFunc(&_mDLabelCopyAttribute, frameworkHandle, "MDLabelCopyAttribute")
		registerFunc(&_mDLabelCopyAttributeName, frameworkHandle, "MDLabelCopyAttributeName")
		registerFunc(&_mDLabelCreate, frameworkHandle, "MDLabelCreate")
		registerFunc(&_mDLabelDelete, frameworkHandle, "MDLabelDelete")
		registerFunc(&_mDLabelGetTypeID, frameworkHandle, "MDLabelGetTypeID")
		registerFunc(&_mDLabelSetAttributes, frameworkHandle, "MDLabelSetAttributes")
		registerFunc(&_mDQueryCopyQueryString, frameworkHandle, "MDQueryCopyQueryString")
		registerFunc(&_mDQueryCopySortingAttributes, frameworkHandle, "MDQueryCopySortingAttributes")
		registerFunc(&_mDQueryCopyValueListAttributes, frameworkHandle, "MDQueryCopyValueListAttributes")
		registerFunc(&_mDQueryCopyValuesOfAttribute, frameworkHandle, "MDQueryCopyValuesOfAttribute")
		registerFunc(&_mDQueryCreate, frameworkHandle, "MDQueryCreate")
		registerFunc(&_mDQueryCreateForItems, frameworkHandle, "MDQueryCreateForItems")
		registerFunc(&_mDQueryCreateSubset, frameworkHandle, "MDQueryCreateSubset")
		registerFunc(&_mDQueryDisableUpdates, frameworkHandle, "MDQueryDisableUpdates")
		registerFunc(&_mDQueryEnableUpdates, frameworkHandle, "MDQueryEnableUpdates")
		registerFunc(&_mDQueryExecute, frameworkHandle, "MDQueryExecute")
		registerFunc(&_mDQueryGetAttributeValueOfResultAtIndex, frameworkHandle, "MDQueryGetAttributeValueOfResultAtIndex")
		registerFunc(&_mDQueryGetBatchingParameters, frameworkHandle, "MDQueryGetBatchingParameters")
		registerFunc(&_mDQueryGetCountOfResultsWithAttributeValue, frameworkHandle, "MDQueryGetCountOfResultsWithAttributeValue")
		registerFunc(&_mDQueryGetIndexOfResult, frameworkHandle, "MDQueryGetIndexOfResult")
		registerFunc(&_mDQueryGetResultAtIndex, frameworkHandle, "MDQueryGetResultAtIndex")
		registerFunc(&_mDQueryGetResultCount, frameworkHandle, "MDQueryGetResultCount")
		registerFunc(&_mDQueryGetSortOptionFlagsForAttribute, frameworkHandle, "MDQueryGetSortOptionFlagsForAttribute")
		registerFunc(&_mDQueryGetTypeID, frameworkHandle, "MDQueryGetTypeID")
		registerFunc(&_mDQueryIsGatheringComplete, frameworkHandle, "MDQueryIsGatheringComplete")
		registerFunc(&_mDQuerySetBatchingParameters, frameworkHandle, "MDQuerySetBatchingParameters")
		registerFunc(&_mDQuerySetCreateResultFunction, frameworkHandle, "MDQuerySetCreateResultFunction")
		registerFunc(&_mDQuerySetCreateValueFunction, frameworkHandle, "MDQuerySetCreateValueFunction")
		registerFunc(&_mDQuerySetDispatchQueue, frameworkHandle, "MDQuerySetDispatchQueue")
		registerFunc(&_mDQuerySetMaxCount, frameworkHandle, "MDQuerySetMaxCount")
		registerFunc(&_mDQuerySetSearchScope, frameworkHandle, "MDQuerySetSearchScope")
		registerFunc(&_mDQuerySetSortComparator, frameworkHandle, "MDQuerySetSortComparator")
		registerFunc(&_mDQuerySetSortComparatorBlock, frameworkHandle, "MDQuerySetSortComparatorBlock")
		registerFunc(&_mDQuerySetSortOptionFlagsForAttribute, frameworkHandle, "MDQuerySetSortOptionFlagsForAttribute")
		registerFunc(&_mDQuerySetSortOrder, frameworkHandle, "MDQuerySetSortOrder")
		registerFunc(&_mDQueryStop, frameworkHandle, "MDQueryStop")
		registerFunc(&_mDSchemaCopyAllAttributes, frameworkHandle, "MDSchemaCopyAllAttributes")
		registerFunc(&_mDSchemaCopyAttributesForContentType, frameworkHandle, "MDSchemaCopyAttributesForContentType")
		registerFunc(&_mDSchemaCopyDisplayDescriptionForAttribute, frameworkHandle, "MDSchemaCopyDisplayDescriptionForAttribute")
		registerFunc(&_mDSchemaCopyDisplayNameForAttribute, frameworkHandle, "MDSchemaCopyDisplayNameForAttribute")
		registerFunc(&_mDSchemaCopyMetaAttributesForAttribute, frameworkHandle, "MDSchemaCopyMetaAttributesForAttribute")
		registerFunc(&_mPAllocateAligned, frameworkHandle, "MPAllocateAligned")
		registerFunc(&_mPAllocateTaskStorageIndex, frameworkHandle, "MPAllocateTaskStorageIndex")
		registerFunc(&_mPArmTimer, frameworkHandle, "MPArmTimer")
		registerFunc(&_mPBlockClear, frameworkHandle, "MPBlockClear")
		registerFunc(&_mPBlockCopy, frameworkHandle, "MPBlockCopy")
		registerFunc(&_mPCancelTimer, frameworkHandle, "MPCancelTimer")
		registerFunc(&_mPCauseNotification, frameworkHandle, "MPCauseNotification")
		registerFunc(&_mPCreateCriticalRegion, frameworkHandle, "MPCreateCriticalRegion")
		registerFunc(&_mPCreateEvent, frameworkHandle, "MPCreateEvent")
		registerFunc(&_mPCreateNotification, frameworkHandle, "MPCreateNotification")
		registerFunc(&_mPCreateQueue, frameworkHandle, "MPCreateQueue")
		registerFunc(&_mPCreateSemaphore, frameworkHandle, "MPCreateSemaphore")
		registerFunc(&_mPCreateTask, frameworkHandle, "MPCreateTask")
		registerFunc(&_mPCreateTimer, frameworkHandle, "MPCreateTimer")
		registerFunc(&_mPCurrentTaskID, frameworkHandle, "MPCurrentTaskID")
		registerFunc(&_mPDeallocateTaskStorageIndex, frameworkHandle, "MPDeallocateTaskStorageIndex")
		registerFunc(&_mPDelayUntil, frameworkHandle, "MPDelayUntil")
		registerFunc(&_mPDeleteCriticalRegion, frameworkHandle, "MPDeleteCriticalRegion")
		registerFunc(&_mPDeleteEvent, frameworkHandle, "MPDeleteEvent")
		registerFunc(&_mPDeleteNotification, frameworkHandle, "MPDeleteNotification")
		registerFunc(&_mPDeleteQueue, frameworkHandle, "MPDeleteQueue")
		registerFunc(&_mPDeleteSemaphore, frameworkHandle, "MPDeleteSemaphore")
		registerFunc(&_mPDeleteTimer, frameworkHandle, "MPDeleteTimer")
		registerFunc(&_mPDisposeTaskException, frameworkHandle, "MPDisposeTaskException")
		registerFunc(&_mPEnterCriticalRegion, frameworkHandle, "MPEnterCriticalRegion")
		registerFunc(&_mPExit, frameworkHandle, "MPExit")
		registerFunc(&_mPExitCriticalRegion, frameworkHandle, "MPExitCriticalRegion")
		registerFunc(&_mPExtractTaskState, frameworkHandle, "MPExtractTaskState")
		registerFunc(&_mPFree, frameworkHandle, "MPFree")
		registerFunc(&_mPGetAllocatedBlockSize, frameworkHandle, "MPGetAllocatedBlockSize")
		registerFunc(&_mPGetNextCpuID, frameworkHandle, "MPGetNextCpuID")
		registerFunc(&_mPGetNextTaskID, frameworkHandle, "MPGetNextTaskID")
		registerFunc(&_mPGetTaskStorageValue, frameworkHandle, "MPGetTaskStorageValue")
		registerFunc(&_mPModifyNotification, frameworkHandle, "MPModifyNotification")
		registerFunc(&_mPModifyNotificationParameters, frameworkHandle, "MPModifyNotificationParameters")
		registerFunc(&_mPNotifyQueue, frameworkHandle, "MPNotifyQueue")
		registerFunc(&_mPProcessors, frameworkHandle, "MPProcessors")
		registerFunc(&_mPProcessorsScheduled, frameworkHandle, "MPProcessorsScheduled")
		registerFunc(&_mPRegisterDebugger, frameworkHandle, "MPRegisterDebugger")
		registerFunc(&_mPRemoteCall, frameworkHandle, "MPRemoteCall")
		registerFunc(&_mPRemoteCallCFM, frameworkHandle, "MPRemoteCallCFM")
		registerFunc(&_mPSetEvent, frameworkHandle, "MPSetEvent")
		registerFunc(&_mPSetExceptionHandler, frameworkHandle, "MPSetExceptionHandler")
		registerFunc(&_mPSetQueueReserve, frameworkHandle, "MPSetQueueReserve")
		registerFunc(&_mPSetTaskState, frameworkHandle, "MPSetTaskState")
		registerFunc(&_mPSetTaskStorageValue, frameworkHandle, "MPSetTaskStorageValue")
		registerFunc(&_mPSetTaskType, frameworkHandle, "MPSetTaskType")
		registerFunc(&_mPSetTaskWeight, frameworkHandle, "MPSetTaskWeight")
		registerFunc(&_mPSetTimerNotify, frameworkHandle, "MPSetTimerNotify")
		registerFunc(&_mPSignalSemaphore, frameworkHandle, "MPSignalSemaphore")
		registerFunc(&_mPTaskIsPreemptive, frameworkHandle, "MPTaskIsPreemptive")
		registerFunc(&_mPTerminateTask, frameworkHandle, "MPTerminateTask")
		registerFunc(&_mPUnregisterDebugger, frameworkHandle, "MPUnregisterDebugger")
		registerFunc(&_mPWaitForEvent, frameworkHandle, "MPWaitForEvent")
		registerFunc(&_mPWaitOnQueue, frameworkHandle, "MPWaitOnQueue")
		registerFunc(&_mPWaitOnSemaphore, frameworkHandle, "MPWaitOnSemaphore")
		registerFunc(&_mPYield, frameworkHandle, "MPYield")
		registerFunc(&_maximumProcessorSpeed, frameworkHandle, "MaximumProcessorSpeed")
		registerFunc(&_memError, frameworkHandle, "MemError")
		registerFunc(&_microseconds, frameworkHandle, "Microseconds")
		registerFunc(&_minimumProcessorSpeed, frameworkHandle, "MinimumProcessorSpeed")
		registerFunc(&_munger, frameworkHandle, "Munger")
		registerFunc(&_nanosecondsToAbsolute, frameworkHandle, "NanosecondsToAbsolute")
		registerFunc(&_nanosecondsToDuration, frameworkHandle, "NanosecondsToDuration")
		registerFunc(&_nearestMacTextEncodings, frameworkHandle, "NearestMacTextEncodings")
		registerFunc(&_newAECoerceDescUPP, frameworkHandle, "NewAECoerceDescUPP")
		registerFunc(&_newAECoercePtrUPP, frameworkHandle, "NewAECoercePtrUPP")
		registerFunc(&_newAEDisposeExternalUPP, frameworkHandle, "NewAEDisposeExternalUPP")
		registerFunc(&_newAEEventHandlerUPP, frameworkHandle, "NewAEEventHandlerUPP")
		registerFunc(&_newCollection, frameworkHandle, "NewCollection")
		registerFunc(&_newCollectionExceptionUPP, frameworkHandle, "NewCollectionExceptionUPP")
		registerFunc(&_newCollectionFlattenUPP, frameworkHandle, "NewCollectionFlattenUPP")
		registerFunc(&_newComponentFunctionUPP, frameworkHandle, "NewComponentFunctionUPP")
		registerFunc(&_newComponentMPWorkFunctionUPP, frameworkHandle, "NewComponentMPWorkFunctionUPP")
		registerFunc(&_newComponentRoutineUPP, frameworkHandle, "NewComponentRoutineUPP")
		registerFunc(&_newDebugAssertOutputHandlerUPP, frameworkHandle, "NewDebugAssertOutputHandlerUPP")
		registerFunc(&_newDebugComponent, frameworkHandle, "NewDebugComponent")
		registerFunc(&_newDebugComponentCallbackUPP, frameworkHandle, "NewDebugComponentCallbackUPP")
		registerFunc(&_newDebugOption, frameworkHandle, "NewDebugOption")
		registerFunc(&_newDebuggerDisposeThreadUPP, frameworkHandle, "NewDebuggerDisposeThreadUPP")
		registerFunc(&_newDebuggerNewThreadUPP, frameworkHandle, "NewDebuggerNewThreadUPP")
		registerFunc(&_newDebuggerThreadSchedulerUPP, frameworkHandle, "NewDebuggerThreadSchedulerUPP")
		registerFunc(&_newDeferredTaskUPP, frameworkHandle, "NewDeferredTaskUPP")
		registerFunc(&_newEmptyHandle, frameworkHandle, "NewEmptyHandle")
		registerFunc(&_newExceptionHandlerUPP, frameworkHandle, "NewExceptionHandlerUPP")
		registerFunc(&_newFNSubscriptionUPP, frameworkHandle, "NewFNSubscriptionUPP")
		registerFunc(&_newFSVolumeEjectUPP, frameworkHandle, "NewFSVolumeEjectUPP")
		registerFunc(&_newFSVolumeMountUPP, frameworkHandle, "NewFSVolumeMountUPP")
		registerFunc(&_newFSVolumeUnmountUPP, frameworkHandle, "NewFSVolumeUnmountUPP")
		registerFunc(&_newFolderManagerNotificationUPP, frameworkHandle, "NewFolderManagerNotificationUPP")
		registerFunc(&_newGestaltValue, frameworkHandle, "NewGestaltValue")
		registerFunc(&_newGetMissingComponentResourceUPP, frameworkHandle, "NewGetMissingComponentResourceUPP")
		registerFunc(&_newHandle, frameworkHandle, "NewHandle")
		registerFunc(&_newHandleClear, frameworkHandle, "NewHandleClear")
		registerFunc(&_newIOCompletionUPP, frameworkHandle, "NewIOCompletionUPP")
		registerFunc(&_newIndexToUCStringUPP, frameworkHandle, "NewIndexToUCStringUPP")
		registerFunc(&_newKCCallbackUPP, frameworkHandle, "NewKCCallbackUPP")
		registerFunc(&_newOSLAccessorUPP, frameworkHandle, "NewOSLAccessorUPP")
		registerFunc(&_newOSLAdjustMarksUPP, frameworkHandle, "NewOSLAdjustMarksUPP")
		registerFunc(&_newOSLCompareUPP, frameworkHandle, "NewOSLCompareUPP")
		registerFunc(&_newOSLCountUPP, frameworkHandle, "NewOSLCountUPP")
		registerFunc(&_newOSLDisposeTokenUPP, frameworkHandle, "NewOSLDisposeTokenUPP")
		registerFunc(&_newOSLGetErrDescUPP, frameworkHandle, "NewOSLGetErrDescUPP")
		registerFunc(&_newOSLGetMarkTokenUPP, frameworkHandle, "NewOSLGetMarkTokenUPP")
		registerFunc(&_newOSLMarkUPP, frameworkHandle, "NewOSLMarkUPP")
		registerFunc(&_newPtr, frameworkHandle, "NewPtr")
		registerFunc(&_newPtrClear, frameworkHandle, "NewPtrClear")
		registerFunc(&_newResErrUPP, frameworkHandle, "NewResErrUPP")
		registerFunc(&_newSelectorFunctionUPP, frameworkHandle, "NewSelectorFunctionUPP")
		registerFunc(&_newSleepQUPP, frameworkHandle, "NewSleepQUPP")
		registerFunc(&_newThread, frameworkHandle, "NewThread")
		registerFunc(&_newThreadEntryUPP, frameworkHandle, "NewThreadEntryUPP")
		registerFunc(&_newThreadSchedulerUPP, frameworkHandle, "NewThreadSchedulerUPP")
		registerFunc(&_newThreadSwitchUPP, frameworkHandle, "NewThreadSwitchUPP")
		registerFunc(&_newThreadTerminationUPP, frameworkHandle, "NewThreadTerminationUPP")
		registerFunc(&_newTimerUPP, frameworkHandle, "NewTimerUPP")
		registerFunc(&_newUnicodeToTextFallbackUPP, frameworkHandle, "NewUnicodeToTextFallbackUPP")
		registerFunc(&_openAComponent, frameworkHandle, "OpenAComponent")
		registerFunc(&_openAComponentResFile, frameworkHandle, "OpenAComponentResFile")
		registerFunc(&_openADefaultComponent, frameworkHandle, "OpenADefaultComponent")
		registerFunc(&_openComponent, frameworkHandle, "OpenComponent")
		registerFunc(&_openComponentResFile, frameworkHandle, "OpenComponentResFile")
		registerFunc(&_openDefaultComponent, frameworkHandle, "OpenDefaultComponent")
		registerFunc(&_overrideIconRef, frameworkHandle, "OverrideIconRef")
		registerFunc(&_pBAllocateForkAsync, frameworkHandle, "PBAllocateForkAsync")
		registerFunc(&_pBAllocateForkSync, frameworkHandle, "PBAllocateForkSync")
		registerFunc(&_pBCatalogSearchAsync, frameworkHandle, "PBCatalogSearchAsync")
		registerFunc(&_pBCatalogSearchSync, frameworkHandle, "PBCatalogSearchSync")
		registerFunc(&_pBCloseForkAsync, frameworkHandle, "PBCloseForkAsync")
		registerFunc(&_pBCloseForkSync, frameworkHandle, "PBCloseForkSync")
		registerFunc(&_pBCloseIteratorAsync, frameworkHandle, "PBCloseIteratorAsync")
		registerFunc(&_pBCloseIteratorSync, frameworkHandle, "PBCloseIteratorSync")
		registerFunc(&_pBCompareFSRefsAsync, frameworkHandle, "PBCompareFSRefsAsync")
		registerFunc(&_pBCompareFSRefsSync, frameworkHandle, "PBCompareFSRefsSync")
		registerFunc(&_pBCreateDirectoryUnicodeAsync, frameworkHandle, "PBCreateDirectoryUnicodeAsync")
		registerFunc(&_pBCreateDirectoryUnicodeSync, frameworkHandle, "PBCreateDirectoryUnicodeSync")
		registerFunc(&_pBCreateFileAndOpenForkUnicodeAsync, frameworkHandle, "PBCreateFileAndOpenForkUnicodeAsync")
		registerFunc(&_pBCreateFileAndOpenForkUnicodeSync, frameworkHandle, "PBCreateFileAndOpenForkUnicodeSync")
		registerFunc(&_pBCreateFileUnicodeAsync, frameworkHandle, "PBCreateFileUnicodeAsync")
		registerFunc(&_pBCreateFileUnicodeSync, frameworkHandle, "PBCreateFileUnicodeSync")
		registerFunc(&_pBCreateForkAsync, frameworkHandle, "PBCreateForkAsync")
		registerFunc(&_pBCreateForkSync, frameworkHandle, "PBCreateForkSync")
		registerFunc(&_pBDeleteForkAsync, frameworkHandle, "PBDeleteForkAsync")
		registerFunc(&_pBDeleteForkSync, frameworkHandle, "PBDeleteForkSync")
		registerFunc(&_pBDeleteObjectAsync, frameworkHandle, "PBDeleteObjectAsync")
		registerFunc(&_pBDeleteObjectSync, frameworkHandle, "PBDeleteObjectSync")
		registerFunc(&_pBExchangeObjectsAsync, frameworkHandle, "PBExchangeObjectsAsync")
		registerFunc(&_pBExchangeObjectsSync, frameworkHandle, "PBExchangeObjectsSync")
		registerFunc(&_pBFSCopyFileAsync, frameworkHandle, "PBFSCopyFileAsync")
		registerFunc(&_pBFSCopyFileSync, frameworkHandle, "PBFSCopyFileSync")
		registerFunc(&_pBFSResolveNodeIDAsync, frameworkHandle, "PBFSResolveNodeIDAsync")
		registerFunc(&_pBFSResolveNodeIDSync, frameworkHandle, "PBFSResolveNodeIDSync")
		registerFunc(&_pBFlushForkAsync, frameworkHandle, "PBFlushForkAsync")
		registerFunc(&_pBFlushForkSync, frameworkHandle, "PBFlushForkSync")
		registerFunc(&_pBFlushVolumeAsync, frameworkHandle, "PBFlushVolumeAsync")
		registerFunc(&_pBFlushVolumeSync, frameworkHandle, "PBFlushVolumeSync")
		registerFunc(&_pBGetCatalogInfoAsync, frameworkHandle, "PBGetCatalogInfoAsync")
		registerFunc(&_pBGetCatalogInfoBulkAsync, frameworkHandle, "PBGetCatalogInfoBulkAsync")
		registerFunc(&_pBGetCatalogInfoBulkSync, frameworkHandle, "PBGetCatalogInfoBulkSync")
		registerFunc(&_pBGetCatalogInfoSync, frameworkHandle, "PBGetCatalogInfoSync")
		registerFunc(&_pBGetForkCBInfoAsync, frameworkHandle, "PBGetForkCBInfoAsync")
		registerFunc(&_pBGetForkCBInfoSync, frameworkHandle, "PBGetForkCBInfoSync")
		registerFunc(&_pBGetForkPositionAsync, frameworkHandle, "PBGetForkPositionAsync")
		registerFunc(&_pBGetForkPositionSync, frameworkHandle, "PBGetForkPositionSync")
		registerFunc(&_pBGetForkSizeAsync, frameworkHandle, "PBGetForkSizeAsync")
		registerFunc(&_pBGetForkSizeSync, frameworkHandle, "PBGetForkSizeSync")
		registerFunc(&_pBGetVolumeInfoAsync, frameworkHandle, "PBGetVolumeInfoAsync")
		registerFunc(&_pBGetVolumeInfoSync, frameworkHandle, "PBGetVolumeInfoSync")
		registerFunc(&_pBIterateForksAsync, frameworkHandle, "PBIterateForksAsync")
		registerFunc(&_pBIterateForksSync, frameworkHandle, "PBIterateForksSync")
		registerFunc(&_pBMakeFSRefUnicodeAsync, frameworkHandle, "PBMakeFSRefUnicodeAsync")
		registerFunc(&_pBMakeFSRefUnicodeSync, frameworkHandle, "PBMakeFSRefUnicodeSync")
		registerFunc(&_pBMoveObjectAsync, frameworkHandle, "PBMoveObjectAsync")
		registerFunc(&_pBMoveObjectSync, frameworkHandle, "PBMoveObjectSync")
		registerFunc(&_pBOpenForkAsync, frameworkHandle, "PBOpenForkAsync")
		registerFunc(&_pBOpenForkSync, frameworkHandle, "PBOpenForkSync")
		registerFunc(&_pBOpenIteratorAsync, frameworkHandle, "PBOpenIteratorAsync")
		registerFunc(&_pBOpenIteratorSync, frameworkHandle, "PBOpenIteratorSync")
		registerFunc(&_pBReadForkAsync, frameworkHandle, "PBReadForkAsync")
		registerFunc(&_pBReadForkSync, frameworkHandle, "PBReadForkSync")
		registerFunc(&_pBRenameUnicodeAsync, frameworkHandle, "PBRenameUnicodeAsync")
		registerFunc(&_pBRenameUnicodeSync, frameworkHandle, "PBRenameUnicodeSync")
		registerFunc(&_pBSetCatalogInfoAsync, frameworkHandle, "PBSetCatalogInfoAsync")
		registerFunc(&_pBSetCatalogInfoSync, frameworkHandle, "PBSetCatalogInfoSync")
		registerFunc(&_pBSetForkPositionAsync, frameworkHandle, "PBSetForkPositionAsync")
		registerFunc(&_pBSetForkPositionSync, frameworkHandle, "PBSetForkPositionSync")
		registerFunc(&_pBSetForkSizeAsync, frameworkHandle, "PBSetForkSizeAsync")
		registerFunc(&_pBSetForkSizeSync, frameworkHandle, "PBSetForkSizeSync")
		registerFunc(&_pBSetVolumeInfoAsync, frameworkHandle, "PBSetVolumeInfoAsync")
		registerFunc(&_pBSetVolumeInfoSync, frameworkHandle, "PBSetVolumeInfoSync")
		registerFunc(&_pBUnlinkObjectAsync, frameworkHandle, "PBUnlinkObjectAsync")
		registerFunc(&_pBUnlinkObjectSync, frameworkHandle, "PBUnlinkObjectSync")
		registerFunc(&_pBWriteForkAsync, frameworkHandle, "PBWriteForkAsync")
		registerFunc(&_pBWriteForkSync, frameworkHandle, "PBWriteForkSync")
		registerFunc(&_pBXLockRangeAsync, frameworkHandle, "PBXLockRangeAsync")
		registerFunc(&_pBXLockRangeSync, frameworkHandle, "PBXLockRangeSync")
		registerFunc(&_pBXUnlockRangeAsync, frameworkHandle, "PBXUnlockRangeAsync")
		registerFunc(&_pBXUnlockRangeSync, frameworkHandle, "PBXUnlockRangeSync")
		registerFunc(&_pLpos, frameworkHandle, "PLpos")
		registerFunc(&_pLstrcat, frameworkHandle, "PLstrcat")
		registerFunc(&_pLstrchr, frameworkHandle, "PLstrchr")
		registerFunc(&_pLstrcmp, frameworkHandle, "PLstrcmp")
		registerFunc(&_pLstrcpy, frameworkHandle, "PLstrcpy")
		registerFunc(&_pLstrlen, frameworkHandle, "PLstrlen")
		registerFunc(&_pLstrncat, frameworkHandle, "PLstrncat")
		registerFunc(&_pLstrncmp, frameworkHandle, "PLstrncmp")
		registerFunc(&_pLstrncpy, frameworkHandle, "PLstrncpy")
		registerFunc(&_pLstrpbrk, frameworkHandle, "PLstrpbrk")
		registerFunc(&_pLstrrchr, frameworkHandle, "PLstrrchr")
		registerFunc(&_pLstrspn, frameworkHandle, "PLstrspn")
		registerFunc(&_pLstrstr, frameworkHandle, "PLstrstr")
		registerFunc(&_primeTime, frameworkHandle, "PrimeTime")
		registerFunc(&_primeTimeTask, frameworkHandle, "PrimeTimeTask")
		registerFunc(&_ptrAndHand, frameworkHandle, "PtrAndHand")
		registerFunc(&_ptrToHand, frameworkHandle, "PtrToHand")
		registerFunc(&_ptrToXHand, frameworkHandle, "PtrToXHand")
		registerFunc(&_purgeCollection, frameworkHandle, "PurgeCollection")
		registerFunc(&_purgeCollectionTag, frameworkHandle, "PurgeCollectionTag")
		registerFunc(&_queryUnicodeMappings, frameworkHandle, "QueryUnicodeMappings")
		registerFunc(&_readIconFromFSRef, frameworkHandle, "ReadIconFromFSRef")
		registerFunc(&_readLocation, frameworkHandle, "ReadLocation")
		registerFunc(&_readPartialResource, frameworkHandle, "ReadPartialResource")
		registerFunc(&_reallocateHandle, frameworkHandle, "ReallocateHandle")
		registerFunc(&_recoverHandle, frameworkHandle, "RecoverHandle")
		registerFunc(&_registerComponent, frameworkHandle, "RegisterComponent")
		registerFunc(&_registerComponentFileRef, frameworkHandle, "RegisterComponentFileRef")
		registerFunc(&_registerComponentFileRefEntries, frameworkHandle, "RegisterComponentFileRefEntries")
		registerFunc(&_registerComponentResource, frameworkHandle, "RegisterComponentResource")
		registerFunc(&_registerComponentResourceFile, frameworkHandle, "RegisterComponentResourceFile")
		registerFunc(&_registerIconRefFromFSRef, frameworkHandle, "RegisterIconRefFromFSRef")
		registerFunc(&_registerIconRefFromIconFamily, frameworkHandle, "RegisterIconRefFromIconFamily")
		registerFunc(&_releaseCollection, frameworkHandle, "ReleaseCollection")
		registerFunc(&_releaseFolder, frameworkHandle, "ReleaseFolder")
		registerFunc(&_releaseIconRef, frameworkHandle, "ReleaseIconRef")
		registerFunc(&_releaseResource, frameworkHandle, "ReleaseResource")
		registerFunc(&_removeCollectionItem, frameworkHandle, "RemoveCollectionItem")
		registerFunc(&_removeFolderDescriptor, frameworkHandle, "RemoveFolderDescriptor")
		registerFunc(&_removeIconRefOverride, frameworkHandle, "RemoveIconRefOverride")
		registerFunc(&_removeIndexedCollectionItem, frameworkHandle, "RemoveIndexedCollectionItem")
		registerFunc(&_removeResource, frameworkHandle, "RemoveResource")
		registerFunc(&_removeTimeTask, frameworkHandle, "RemoveTimeTask")
		registerFunc(&_replaceGestaltValue, frameworkHandle, "ReplaceGestaltValue")
		registerFunc(&_replaceIndexedCollectionItem, frameworkHandle, "ReplaceIndexedCollectionItem")
		registerFunc(&_replaceIndexedCollectionItemHdl, frameworkHandle, "ReplaceIndexedCollectionItemHdl")
		registerFunc(&_resError, frameworkHandle, "ResError")
		registerFunc(&_resetTextToUnicodeInfo, frameworkHandle, "ResetTextToUnicodeInfo")
		registerFunc(&_resetUnicodeToTextInfo, frameworkHandle, "ResetUnicodeToTextInfo")
		registerFunc(&_resetUnicodeToTextRunInfo, frameworkHandle, "ResetUnicodeToTextRunInfo")
		registerFunc(&_resolveComponentAlias, frameworkHandle, "ResolveComponentAlias")
		registerFunc(&_resolveDefaultTextEncoding, frameworkHandle, "ResolveDefaultTextEncoding")
		registerFunc(&_retainCollection, frameworkHandle, "RetainCollection")
		registerFunc(&_revertTextEncodingToScriptInfo, frameworkHandle, "RevertTextEncodingToScriptInfo")
		registerFunc(&_rmvTime, frameworkHandle, "RmvTime")
		registerFunc(&_s32Set, frameworkHandle, "S32Set")
		registerFunc(&_s64Absolute, frameworkHandle, "S64Absolute")
		registerFunc(&_s64Add, frameworkHandle, "S64Add")
		registerFunc(&_s64And, frameworkHandle, "S64And")
		registerFunc(&_s64BitwiseAnd, frameworkHandle, "S64BitwiseAnd")
		registerFunc(&_s64BitwiseEor, frameworkHandle, "S64BitwiseEor")
		registerFunc(&_s64BitwiseNot, frameworkHandle, "S64BitwiseNot")
		registerFunc(&_s64BitwiseOr, frameworkHandle, "S64BitwiseOr")
		registerFunc(&_s64Compare, frameworkHandle, "S64Compare")
		registerFunc(&_s64Div, frameworkHandle, "S64Div")
		registerFunc(&_s64Divide, frameworkHandle, "S64Divide")
		registerFunc(&_s64Eor, frameworkHandle, "S64Eor")
		registerFunc(&_s64Max, frameworkHandle, "S64Max")
		registerFunc(&_s64Min, frameworkHandle, "S64Min")
		registerFunc(&_s64Mod, frameworkHandle, "S64Mod")
		registerFunc(&_s64Multiply, frameworkHandle, "S64Multiply")
		registerFunc(&_s64Negate, frameworkHandle, "S64Negate")
		registerFunc(&_s64Not, frameworkHandle, "S64Not")
		registerFunc(&_s64Or, frameworkHandle, "S64Or")
		registerFunc(&_s64Set, frameworkHandle, "S64Set")
		registerFunc(&_s64SetU, frameworkHandle, "S64SetU")
		registerFunc(&_s64ShiftLeft, frameworkHandle, "S64ShiftLeft")
		registerFunc(&_s64ShiftRight, frameworkHandle, "S64ShiftRight")
		registerFunc(&_s64Subtract, frameworkHandle, "S64Subtract")
		registerFunc(&_sInt64ToLongDouble, frameworkHandle, "SInt64ToLongDouble")
		registerFunc(&_sInt64ToUInt64, frameworkHandle, "SInt64ToUInt64")
		registerFunc(&_sInt64ToWide, frameworkHandle, "SInt64ToWide")
		registerFunc(&_sKDocumentCopyURL, frameworkHandle, "SKDocumentCopyURL")
		registerFunc(&_sKDocumentCreate, frameworkHandle, "SKDocumentCreate")
		registerFunc(&_sKDocumentCreateWithURL, frameworkHandle, "SKDocumentCreateWithURL")
		registerFunc(&_sKDocumentGetName, frameworkHandle, "SKDocumentGetName")
		registerFunc(&_sKDocumentGetParent, frameworkHandle, "SKDocumentGetParent")
		registerFunc(&_sKDocumentGetSchemeName, frameworkHandle, "SKDocumentGetSchemeName")
		registerFunc(&_sKDocumentGetTypeID, frameworkHandle, "SKDocumentGetTypeID")
		registerFunc(&_sKIndexAddDocument, frameworkHandle, "SKIndexAddDocument")
		registerFunc(&_sKIndexAddDocumentWithText, frameworkHandle, "SKIndexAddDocumentWithText")
		registerFunc(&_sKIndexClose, frameworkHandle, "SKIndexClose")
		registerFunc(&_sKIndexCompact, frameworkHandle, "SKIndexCompact")
		registerFunc(&_sKIndexCopyDocumentForDocumentID, frameworkHandle, "SKIndexCopyDocumentForDocumentID")
		registerFunc(&_sKIndexCopyDocumentIDArrayForTermID, frameworkHandle, "SKIndexCopyDocumentIDArrayForTermID")
		registerFunc(&_sKIndexCopyDocumentProperties, frameworkHandle, "SKIndexCopyDocumentProperties")
		registerFunc(&_sKIndexCopyDocumentRefsForDocumentIDs, frameworkHandle, "SKIndexCopyDocumentRefsForDocumentIDs")
		registerFunc(&_sKIndexCopyDocumentURLsForDocumentIDs, frameworkHandle, "SKIndexCopyDocumentURLsForDocumentIDs")
		registerFunc(&_sKIndexCopyInfoForDocumentIDs, frameworkHandle, "SKIndexCopyInfoForDocumentIDs")
		registerFunc(&_sKIndexCopyTermIDArrayForDocumentID, frameworkHandle, "SKIndexCopyTermIDArrayForDocumentID")
		registerFunc(&_sKIndexCopyTermStringForTermID, frameworkHandle, "SKIndexCopyTermStringForTermID")
		registerFunc(&_sKIndexCreateWithMutableData, frameworkHandle, "SKIndexCreateWithMutableData")
		registerFunc(&_sKIndexCreateWithURL, frameworkHandle, "SKIndexCreateWithURL")
		registerFunc(&_sKIndexDocumentIteratorCopyNext, frameworkHandle, "SKIndexDocumentIteratorCopyNext")
		registerFunc(&_sKIndexDocumentIteratorCreate, frameworkHandle, "SKIndexDocumentIteratorCreate")
		registerFunc(&_sKIndexDocumentIteratorGetTypeID, frameworkHandle, "SKIndexDocumentIteratorGetTypeID")
		registerFunc(&_sKIndexFlush, frameworkHandle, "SKIndexFlush")
		registerFunc(&_sKIndexGetAnalysisProperties, frameworkHandle, "SKIndexGetAnalysisProperties")
		registerFunc(&_sKIndexGetDocumentCount, frameworkHandle, "SKIndexGetDocumentCount")
		registerFunc(&_sKIndexGetDocumentID, frameworkHandle, "SKIndexGetDocumentID")
		registerFunc(&_sKIndexGetDocumentState, frameworkHandle, "SKIndexGetDocumentState")
		registerFunc(&_sKIndexGetDocumentTermCount, frameworkHandle, "SKIndexGetDocumentTermCount")
		registerFunc(&_sKIndexGetDocumentTermFrequency, frameworkHandle, "SKIndexGetDocumentTermFrequency")
		registerFunc(&_sKIndexGetIndexType, frameworkHandle, "SKIndexGetIndexType")
		registerFunc(&_sKIndexGetMaximumBytesBeforeFlush, frameworkHandle, "SKIndexGetMaximumBytesBeforeFlush")
		registerFunc(&_sKIndexGetMaximumDocumentID, frameworkHandle, "SKIndexGetMaximumDocumentID")
		registerFunc(&_sKIndexGetMaximumTermID, frameworkHandle, "SKIndexGetMaximumTermID")
		registerFunc(&_sKIndexGetTermDocumentCount, frameworkHandle, "SKIndexGetTermDocumentCount")
		registerFunc(&_sKIndexGetTermIDForTermString, frameworkHandle, "SKIndexGetTermIDForTermString")
		registerFunc(&_sKIndexGetTypeID, frameworkHandle, "SKIndexGetTypeID")
		registerFunc(&_sKIndexMoveDocument, frameworkHandle, "SKIndexMoveDocument")
		registerFunc(&_sKIndexOpenWithData, frameworkHandle, "SKIndexOpenWithData")
		registerFunc(&_sKIndexOpenWithMutableData, frameworkHandle, "SKIndexOpenWithMutableData")
		registerFunc(&_sKIndexOpenWithURL, frameworkHandle, "SKIndexOpenWithURL")
		registerFunc(&_sKIndexRemoveDocument, frameworkHandle, "SKIndexRemoveDocument")
		registerFunc(&_sKIndexRenameDocument, frameworkHandle, "SKIndexRenameDocument")
		registerFunc(&_sKIndexSetDocumentProperties, frameworkHandle, "SKIndexSetDocumentProperties")
		registerFunc(&_sKIndexSetMaximumBytesBeforeFlush, frameworkHandle, "SKIndexSetMaximumBytesBeforeFlush")
		registerFunc(&_sKLoadDefaultExtractorPlugIns, frameworkHandle, "SKLoadDefaultExtractorPlugIns")
		registerFunc(&_sKSearchCancel, frameworkHandle, "SKSearchCancel")
		registerFunc(&_sKSearchCreate, frameworkHandle, "SKSearchCreate")
		registerFunc(&_sKSearchFindMatches, frameworkHandle, "SKSearchFindMatches")
		registerFunc(&_sKSearchGetTypeID, frameworkHandle, "SKSearchGetTypeID")
		registerFunc(&_sKSearchGroupGetTypeID, frameworkHandle, "SKSearchGroupGetTypeID")
		registerFunc(&_sKSummaryCopyParagraphAtIndex, frameworkHandle, "SKSummaryCopyParagraphAtIndex")
		registerFunc(&_sKSummaryCopyParagraphSummaryString, frameworkHandle, "SKSummaryCopyParagraphSummaryString")
		registerFunc(&_sKSummaryCopySentenceAtIndex, frameworkHandle, "SKSummaryCopySentenceAtIndex")
		registerFunc(&_sKSummaryCopySentenceSummaryString, frameworkHandle, "SKSummaryCopySentenceSummaryString")
		registerFunc(&_sKSummaryCreateWithString, frameworkHandle, "SKSummaryCreateWithString")
		registerFunc(&_sKSummaryGetParagraphCount, frameworkHandle, "SKSummaryGetParagraphCount")
		registerFunc(&_sKSummaryGetParagraphSummaryInfo, frameworkHandle, "SKSummaryGetParagraphSummaryInfo")
		registerFunc(&_sKSummaryGetSentenceCount, frameworkHandle, "SKSummaryGetSentenceCount")
		registerFunc(&_sKSummaryGetSentenceSummaryInfo, frameworkHandle, "SKSummaryGetSentenceSummaryInfo")
		registerFunc(&_sKSummaryGetTypeID, frameworkHandle, "SKSummaryGetTypeID")
		registerFunc(&_setCollectionDefaultAttributes, frameworkHandle, "SetCollectionDefaultAttributes")
		registerFunc(&_setCollectionExceptionProc, frameworkHandle, "SetCollectionExceptionProc")
		registerFunc(&_setCollectionItemInfo, frameworkHandle, "SetCollectionItemInfo")
		registerFunc(&_setComponentInstanceError, frameworkHandle, "SetComponentInstanceError")
		registerFunc(&_setComponentInstanceStorage, frameworkHandle, "SetComponentInstanceStorage")
		registerFunc(&_setComponentRefcon, frameworkHandle, "SetComponentRefcon")
		registerFunc(&_setCustomIconsEnabled, frameworkHandle, "SetCustomIconsEnabled")
		registerFunc(&_setDebugOptionValue, frameworkHandle, "SetDebugOptionValue")
		registerFunc(&_setDebuggerNotificationProcs, frameworkHandle, "SetDebuggerNotificationProcs")
		registerFunc(&_setDefaultComponent, frameworkHandle, "SetDefaultComponent")
		registerFunc(&_setFallbackUnicodeToText, frameworkHandle, "SetFallbackUnicodeToText")
		registerFunc(&_setFallbackUnicodeToTextRun, frameworkHandle, "SetFallbackUnicodeToTextRun")
		registerFunc(&_setGestaltValue, frameworkHandle, "SetGestaltValue")
		registerFunc(&_setHandleSize, frameworkHandle, "SetHandleSize")
		registerFunc(&_setIndexedCollectionItemInfo, frameworkHandle, "SetIndexedCollectionItemInfo")
		registerFunc(&_setPtrSize, frameworkHandle, "SetPtrSize")
		registerFunc(&_setResAttrs, frameworkHandle, "SetResAttrs")
		registerFunc(&_setResFileAttrs, frameworkHandle, "SetResFileAttrs")
		registerFunc(&_setResInfo, frameworkHandle, "SetResInfo")
		registerFunc(&_setResLoad, frameworkHandle, "SetResLoad")
		registerFunc(&_setResPurge, frameworkHandle, "SetResPurge")
		registerFunc(&_setResourceSize, frameworkHandle, "SetResourceSize")
		registerFunc(&_setScriptManagerVariable, frameworkHandle, "SetScriptManagerVariable")
		registerFunc(&_setThreadReadyGivenTaskRef, frameworkHandle, "SetThreadReadyGivenTaskRef")
		registerFunc(&_setThreadScheduler, frameworkHandle, "SetThreadScheduler")
		registerFunc(&_setThreadState, frameworkHandle, "SetThreadState")
		registerFunc(&_setThreadStateEndCritical, frameworkHandle, "SetThreadStateEndCritical")
		registerFunc(&_setThreadSwitcher, frameworkHandle, "SetThreadSwitcher")
		registerFunc(&_setThreadTerminator, frameworkHandle, "SetThreadTerminator")
		registerFunc(&_sleepQInstall, frameworkHandle, "SleepQInstall")
		registerFunc(&_sleepQRemove, frameworkHandle, "SleepQRemove")
		registerFunc(&_subAbsoluteFromAbsolute, frameworkHandle, "SubAbsoluteFromAbsolute")
		registerFunc(&_subDurationFromAbsolute, frameworkHandle, "SubDurationFromAbsolute")
		registerFunc(&_subNanosecondsFromAbsolute, frameworkHandle, "SubNanosecondsFromAbsolute")
		registerFunc(&_sysError, frameworkHandle, "SysError")
		registerFunc(&_tECClearConverterContextInfo, frameworkHandle, "TECClearConverterContextInfo")
		registerFunc(&_tECClearSnifferContextInfo, frameworkHandle, "TECClearSnifferContextInfo")
		registerFunc(&_tECConvertText, frameworkHandle, "TECConvertText")
		registerFunc(&_tECConvertTextToMultipleEncodings, frameworkHandle, "TECConvertTextToMultipleEncodings")
		registerFunc(&_tECCopyTextEncodingInternetNameAndMIB, frameworkHandle, "TECCopyTextEncodingInternetNameAndMIB")
		registerFunc(&_tECCountAvailableSniffers, frameworkHandle, "TECCountAvailableSniffers")
		registerFunc(&_tECCountAvailableTextEncodings, frameworkHandle, "TECCountAvailableTextEncodings")
		registerFunc(&_tECCountDestinationTextEncodings, frameworkHandle, "TECCountDestinationTextEncodings")
		registerFunc(&_tECCountDirectTextEncodingConversions, frameworkHandle, "TECCountDirectTextEncodingConversions")
		registerFunc(&_tECCountMailTextEncodings, frameworkHandle, "TECCountMailTextEncodings")
		registerFunc(&_tECCountSubTextEncodings, frameworkHandle, "TECCountSubTextEncodings")
		registerFunc(&_tECCountWebTextEncodings, frameworkHandle, "TECCountWebTextEncodings")
		registerFunc(&_tECCreateConverter, frameworkHandle, "TECCreateConverter")
		registerFunc(&_tECCreateConverterFromPath, frameworkHandle, "TECCreateConverterFromPath")
		registerFunc(&_tECCreateOneToManyConverter, frameworkHandle, "TECCreateOneToManyConverter")
		registerFunc(&_tECCreateSniffer, frameworkHandle, "TECCreateSniffer")
		registerFunc(&_tECDisposeConverter, frameworkHandle, "TECDisposeConverter")
		registerFunc(&_tECDisposeSniffer, frameworkHandle, "TECDisposeSniffer")
		registerFunc(&_tECFlushMultipleEncodings, frameworkHandle, "TECFlushMultipleEncodings")
		registerFunc(&_tECFlushText, frameworkHandle, "TECFlushText")
		registerFunc(&_tECGetAvailableSniffers, frameworkHandle, "TECGetAvailableSniffers")
		registerFunc(&_tECGetAvailableTextEncodings, frameworkHandle, "TECGetAvailableTextEncodings")
		registerFunc(&_tECGetDestinationTextEncodings, frameworkHandle, "TECGetDestinationTextEncodings")
		registerFunc(&_tECGetDirectTextEncodingConversions, frameworkHandle, "TECGetDirectTextEncodingConversions")
		registerFunc(&_tECGetEncodingList, frameworkHandle, "TECGetEncodingList")
		registerFunc(&_tECGetInfo, frameworkHandle, "TECGetInfo")
		registerFunc(&_tECGetMailTextEncodings, frameworkHandle, "TECGetMailTextEncodings")
		registerFunc(&_tECGetSubTextEncodings, frameworkHandle, "TECGetSubTextEncodings")
		registerFunc(&_tECGetTextEncodingFromInternetName, frameworkHandle, "TECGetTextEncodingFromInternetName")
		registerFunc(&_tECGetTextEncodingFromInternetNameOrMIB, frameworkHandle, "TECGetTextEncodingFromInternetNameOrMIB")
		registerFunc(&_tECGetTextEncodingInternetName, frameworkHandle, "TECGetTextEncodingInternetName")
		registerFunc(&_tECGetWebTextEncodings, frameworkHandle, "TECGetWebTextEncodings")
		registerFunc(&_tECSetBasicOptions, frameworkHandle, "TECSetBasicOptions")
		registerFunc(&_tECSniffTextEncoding, frameworkHandle, "TECSniffTextEncoding")
		registerFunc(&_taskLevel, frameworkHandle, "TaskLevel")
		registerFunc(&_tempNewHandle, frameworkHandle, "TempNewHandle")
		registerFunc(&_testAndClear, frameworkHandle, "TestAndClear")
		registerFunc(&_testAndSet, frameworkHandle, "TestAndSet")
		registerFunc(&_threadBeginCritical, frameworkHandle, "ThreadBeginCritical")
		registerFunc(&_threadCurrentStackSpace, frameworkHandle, "ThreadCurrentStackSpace")
		registerFunc(&_threadEndCritical, frameworkHandle, "ThreadEndCritical")
		registerFunc(&_tickCount, frameworkHandle, "TickCount")
		registerFunc(&_truncateForTextToUnicode, frameworkHandle, "TruncateForTextToUnicode")
		registerFunc(&_truncateForUnicodeToText, frameworkHandle, "TruncateForUnicodeToText")
		registerFunc(&_u32SetU, frameworkHandle, "U32SetU")
		registerFunc(&_u64Add, frameworkHandle, "U64Add")
		registerFunc(&_u64And, frameworkHandle, "U64And")
		registerFunc(&_u64BitwiseAnd, frameworkHandle, "U64BitwiseAnd")
		registerFunc(&_u64BitwiseEor, frameworkHandle, "U64BitwiseEor")
		registerFunc(&_u64BitwiseNot, frameworkHandle, "U64BitwiseNot")
		registerFunc(&_u64BitwiseOr, frameworkHandle, "U64BitwiseOr")
		registerFunc(&_u64Compare, frameworkHandle, "U64Compare")
		registerFunc(&_u64Div, frameworkHandle, "U64Div")
		registerFunc(&_u64Divide, frameworkHandle, "U64Divide")
		registerFunc(&_u64Eor, frameworkHandle, "U64Eor")
		registerFunc(&_u64Max, frameworkHandle, "U64Max")
		registerFunc(&_u64Mod, frameworkHandle, "U64Mod")
		registerFunc(&_u64Multiply, frameworkHandle, "U64Multiply")
		registerFunc(&_u64Not, frameworkHandle, "U64Not")
		registerFunc(&_u64Or, frameworkHandle, "U64Or")
		registerFunc(&_u64Set, frameworkHandle, "U64Set")
		registerFunc(&_u64SetU, frameworkHandle, "U64SetU")
		registerFunc(&_u64ShiftLeft, frameworkHandle, "U64ShiftLeft")
		registerFunc(&_u64ShiftRight, frameworkHandle, "U64ShiftRight")
		registerFunc(&_u64Subtract, frameworkHandle, "U64Subtract")
		registerFunc(&_uCCompareCollationKeys, frameworkHandle, "UCCompareCollationKeys")
		registerFunc(&_uCCompareText, frameworkHandle, "UCCompareText")
		registerFunc(&_uCCompareTextDefault, frameworkHandle, "UCCompareTextDefault")
		registerFunc(&_uCCompareTextNoLocale, frameworkHandle, "UCCompareTextNoLocale")
		registerFunc(&_uCConvertCFAbsoluteTimeToLongDateTime, frameworkHandle, "UCConvertCFAbsoluteTimeToLongDateTime")
		registerFunc(&_uCConvertCFAbsoluteTimeToSeconds, frameworkHandle, "UCConvertCFAbsoluteTimeToSeconds")
		registerFunc(&_uCConvertCFAbsoluteTimeToUTCDateTime, frameworkHandle, "UCConvertCFAbsoluteTimeToUTCDateTime")
		registerFunc(&_uCConvertLongDateTimeToCFAbsoluteTime, frameworkHandle, "UCConvertLongDateTimeToCFAbsoluteTime")
		registerFunc(&_uCConvertSecondsToCFAbsoluteTime, frameworkHandle, "UCConvertSecondsToCFAbsoluteTime")
		registerFunc(&_uCConvertUTCDateTimeToCFAbsoluteTime, frameworkHandle, "UCConvertUTCDateTimeToCFAbsoluteTime")
		registerFunc(&_uCCreateCollator, frameworkHandle, "UCCreateCollator")
		registerFunc(&_uCDisposeCollator, frameworkHandle, "UCDisposeCollator")
		registerFunc(&_uCGetCharProperty, frameworkHandle, "UCGetCharProperty")
		registerFunc(&_uCGetCollationKey, frameworkHandle, "UCGetCollationKey")
		registerFunc(&_uCGetUnicodeScalarValueForSurrogatePair, frameworkHandle, "UCGetUnicodeScalarValueForSurrogatePair")
		registerFunc(&_uCIsSurrogateHighCharacter, frameworkHandle, "UCIsSurrogateHighCharacter")
		registerFunc(&_uCIsSurrogateLowCharacter, frameworkHandle, "UCIsSurrogateLowCharacter")
		registerFunc(&_uCKeyTranslate, frameworkHandle, "UCKeyTranslate")
		registerFunc(&_uCTypeSelectAddKeyToSelector, frameworkHandle, "UCTypeSelectAddKeyToSelector")
		registerFunc(&_uCTypeSelectCompare, frameworkHandle, "UCTypeSelectCompare")
		registerFunc(&_uCTypeSelectCreateSelector, frameworkHandle, "UCTypeSelectCreateSelector")
		registerFunc(&_uCTypeSelectFindItem, frameworkHandle, "UCTypeSelectFindItem")
		registerFunc(&_uCTypeSelectFlushSelectorData, frameworkHandle, "UCTypeSelectFlushSelectorData")
		registerFunc(&_uCTypeSelectReleaseSelector, frameworkHandle, "UCTypeSelectReleaseSelector")
		registerFunc(&_uCTypeSelectWalkList, frameworkHandle, "UCTypeSelectWalkList")
		registerFunc(&_uCTypeSelectWouldResetBuffer, frameworkHandle, "UCTypeSelectWouldResetBuffer")
		registerFunc(&_uInt64ToLongDouble, frameworkHandle, "UInt64ToLongDouble")
		registerFunc(&_uInt64ToSInt64, frameworkHandle, "UInt64ToSInt64")
		registerFunc(&_uInt64ToUnsignedWide, frameworkHandle, "UInt64ToUnsignedWide")
		registerFunc(&_uTCreateStringForOSType, frameworkHandle, "UTCreateStringForOSType")
		registerFunc(&_uTGetOSTypeFromString, frameworkHandle, "UTGetOSTypeFromString")
		registerFunc(&_uTTypeConformsTo, frameworkHandle, "UTTypeConformsTo")
		registerFunc(&_uTTypeCopyAllTagsWithClass, frameworkHandle, "UTTypeCopyAllTagsWithClass")
		registerFunc(&_uTTypeCopyDeclaration, frameworkHandle, "UTTypeCopyDeclaration")
		registerFunc(&_uTTypeCopyDeclaringBundleURL, frameworkHandle, "UTTypeCopyDeclaringBundleURL")
		registerFunc(&_uTTypeCopyDescription, frameworkHandle, "UTTypeCopyDescription")
		registerFunc(&_uTTypeCopyPreferredTagWithClass, frameworkHandle, "UTTypeCopyPreferredTagWithClass")
		registerFunc(&_uTTypeCreateAllIdentifiersForTag, frameworkHandle, "UTTypeCreateAllIdentifiersForTag")
		registerFunc(&_uTTypeCreatePreferredIdentifierForTag, frameworkHandle, "UTTypeCreatePreferredIdentifierForTag")
		registerFunc(&_uTTypeEqual, frameworkHandle, "UTTypeEqual")
		registerFunc(&_uTTypeIsDeclared, frameworkHandle, "UTTypeIsDeclared")
		registerFunc(&_uTTypeIsDynamic, frameworkHandle, "UTTypeIsDynamic")
		registerFunc(&_uncaptureComponent, frameworkHandle, "UncaptureComponent")
		registerFunc(&_unflattenCollection, frameworkHandle, "UnflattenCollection")
		registerFunc(&_unflattenCollectionFromHdl, frameworkHandle, "UnflattenCollectionFromHdl")
		registerFunc(&_unique1ID, frameworkHandle, "Unique1ID")
		registerFunc(&_uniqueID, frameworkHandle, "UniqueID")
		registerFunc(&_unregisterComponent, frameworkHandle, "UnregisterComponent")
		registerFunc(&_unregisterIconRef, frameworkHandle, "UnregisterIconRef")
		registerFunc(&_unsignedFixedMulDiv, frameworkHandle, "UnsignedFixedMulDiv")
		registerFunc(&_unsignedWideToUInt64, frameworkHandle, "UnsignedWideToUInt64")
		registerFunc(&_upTime, frameworkHandle, "UpTime")
		registerFunc(&_updateIconRef, frameworkHandle, "UpdateIconRef")
		registerFunc(&_updateResFile, frameworkHandle, "UpdateResFile")
		registerFunc(&_updateSystemActivity, frameworkHandle, "UpdateSystemActivity")
		registerFunc(&_upgradeScriptInfoToTextEncoding, frameworkHandle, "UpgradeScriptInfoToTextEncoding")
		registerFunc(&_useResFile, frameworkHandle, "UseResFile")
		registerFunc(&_wSGetCFTypeIDFromWSTypeID, frameworkHandle, "WSGetCFTypeIDFromWSTypeID")
		registerFunc(&_wSGetWSTypeIDFromCFType, frameworkHandle, "WSGetWSTypeIDFromCFType")
		registerFunc(&_wSMethodInvocationAddDeserializationOverride, frameworkHandle, "WSMethodInvocationAddDeserializationOverride")
		registerFunc(&_wSMethodInvocationAddSerializationOverride, frameworkHandle, "WSMethodInvocationAddSerializationOverride")
		registerFunc(&_wSMethodInvocationCopyParameters, frameworkHandle, "WSMethodInvocationCopyParameters")
		registerFunc(&_wSMethodInvocationCopyProperty, frameworkHandle, "WSMethodInvocationCopyProperty")
		registerFunc(&_wSMethodInvocationCopySerialization, frameworkHandle, "WSMethodInvocationCopySerialization")
		registerFunc(&_wSMethodInvocationCreate, frameworkHandle, "WSMethodInvocationCreate")
		registerFunc(&_wSMethodInvocationCreateFromSerialization, frameworkHandle, "WSMethodInvocationCreateFromSerialization")
		registerFunc(&_wSMethodInvocationGetTypeID, frameworkHandle, "WSMethodInvocationGetTypeID")
		registerFunc(&_wSMethodInvocationInvoke, frameworkHandle, "WSMethodInvocationInvoke")
		registerFunc(&_wSMethodInvocationScheduleWithRunLoop, frameworkHandle, "WSMethodInvocationScheduleWithRunLoop")
		registerFunc(&_wSMethodInvocationSetCallBack, frameworkHandle, "WSMethodInvocationSetCallBack")
		registerFunc(&_wSMethodInvocationSetParameters, frameworkHandle, "WSMethodInvocationSetParameters")
		registerFunc(&_wSMethodInvocationSetProperty, frameworkHandle, "WSMethodInvocationSetProperty")
		registerFunc(&_wSMethodInvocationUnscheduleFromRunLoop, frameworkHandle, "WSMethodInvocationUnscheduleFromRunLoop")
		registerFunc(&_wSMethodResultIsFault, frameworkHandle, "WSMethodResultIsFault")
		registerFunc(&_wSProtocolHandlerCopyFaultDocument, frameworkHandle, "WSProtocolHandlerCopyFaultDocument")
		registerFunc(&_wSProtocolHandlerCopyProperty, frameworkHandle, "WSProtocolHandlerCopyProperty")
		registerFunc(&_wSProtocolHandlerCopyReplyDictionary, frameworkHandle, "WSProtocolHandlerCopyReplyDictionary")
		registerFunc(&_wSProtocolHandlerCopyReplyDocument, frameworkHandle, "WSProtocolHandlerCopyReplyDocument")
		registerFunc(&_wSProtocolHandlerCopyRequestDictionary, frameworkHandle, "WSProtocolHandlerCopyRequestDictionary")
		registerFunc(&_wSProtocolHandlerCopyRequestDocument, frameworkHandle, "WSProtocolHandlerCopyRequestDocument")
		registerFunc(&_wSProtocolHandlerCreate, frameworkHandle, "WSProtocolHandlerCreate")
		registerFunc(&_wSProtocolHandlerGetTypeID, frameworkHandle, "WSProtocolHandlerGetTypeID")
		registerFunc(&_wSProtocolHandlerSetDeserializationOverride, frameworkHandle, "WSProtocolHandlerSetDeserializationOverride")
		registerFunc(&_wSProtocolHandlerSetProperty, frameworkHandle, "WSProtocolHandlerSetProperty")
		registerFunc(&_wSProtocolHandlerSetSerializationOverride, frameworkHandle, "WSProtocolHandlerSetSerializationOverride")
		registerFunc(&_wideAdd, frameworkHandle, "WideAdd")
		registerFunc(&_wideBitShift, frameworkHandle, "WideBitShift")
		registerFunc(&_wideCompare, frameworkHandle, "WideCompare")
		registerFunc(&_wideDivide, frameworkHandle, "WideDivide")
		registerFunc(&_wideMultiply, frameworkHandle, "WideMultiply")
		registerFunc(&_wideNegate, frameworkHandle, "WideNegate")
		registerFunc(&_wideShift, frameworkHandle, "WideShift")
		registerFunc(&_wideSquareRoot, frameworkHandle, "WideSquareRoot")
		registerFunc(&_wideSubtract, frameworkHandle, "WideSubtract")
		registerFunc(&_wideToSInt64, frameworkHandle, "WideToSInt64")
		registerFunc(&_wideWideDivide, frameworkHandle, "WideWideDivide")
		registerFunc(&_writePartialResource, frameworkHandle, "WritePartialResource")
		registerFunc(&_writeResource, frameworkHandle, "WriteResource")
		registerFunc(&_x2Fix, frameworkHandle, "X2Fix")
		registerFunc(&_x2Frac, frameworkHandle, "X2Frac")
		registerFunc(&_yieldToAnyThread, frameworkHandle, "YieldToAnyThread")
		registerFunc(&_yieldToThread, frameworkHandle, "YieldToThread")
		registerFunc(&_annuity, frameworkHandle, "annuity")
		registerFunc(&_compound, frameworkHandle, "compound")
		registerFunc(&_dec2f, frameworkHandle, "dec2f")
		registerFunc(&_dec2l, frameworkHandle, "dec2l")
		registerFunc(&_dec2num, frameworkHandle, "dec2num")
		registerFunc(&_dec2numl, frameworkHandle, "dec2numl")
		registerFunc(&_dec2s, frameworkHandle, "dec2s")
		registerFunc(&_dec2str, frameworkHandle, "dec2str")
		registerFunc(&_dtox80, frameworkHandle, "dtox80")
		registerFunc(&_kcfindapplesharepassword, frameworkHandle, "kcfindapplesharepassword")
		registerFunc(&_kcfindgenericpassword, frameworkHandle, "kcfindgenericpassword")
		registerFunc(&_kcfindinternetpassword, frameworkHandle, "kcfindinternetpassword")
		registerFunc(&_kcfindinternetpasswordwithpath, frameworkHandle, "kcfindinternetpasswordwithpath")
		registerFunc(&_kcgetkeychainname, frameworkHandle, "kcgetkeychainname")
		registerFunc(&_ldtox80, frameworkHandle, "ldtox80")
		registerFunc(&_num2dec, frameworkHandle, "num2dec")
		registerFunc(&_num2decl, frameworkHandle, "num2decl")
		registerFunc(&_numtostring, frameworkHandle, "numtostring")
		registerFunc(&_randomx, frameworkHandle, "randomx")
		registerFunc(&_relation, frameworkHandle, "relation")
		registerFunc(&_relationl, frameworkHandle, "relationl")
		registerFunc(&_str2dec, frameworkHandle, "str2dec")
		registerFunc(&_vAEBuildAppleEvent, frameworkHandle, "vAEBuildAppleEvent")
		registerFunc(&_vAEBuildDesc, frameworkHandle, "vAEBuildDesc")
		registerFunc(&_vAEBuildParameters, frameworkHandle, "vAEBuildParameters")
		registerFunc(&_x80tod, frameworkHandle, "x80tod")
		registerFunc(&_x80told, frameworkHandle, "x80told")
	}

