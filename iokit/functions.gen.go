// Code generated from Apple documentation for iokit. DO NOT EDIT.

package iokit

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/kernel"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: iokit: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: iokit: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: iokit: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _cDConvertLBAToMSF func(arg0 uint32) CDMSF

// CDConvertLBAToMSF.
//
// See: https://developer.apple.com/documentation/iokit/1584342-cdconvertlbatomsf
func CDConvertLBAToMSF(arg0 uint32) CDMSF {
	if _cDConvertLBAToMSF == nil {
		panic("iokit: symbol CDConvertLBAToMSF not loaded")
	}
	return _cDConvertLBAToMSF(arg0)
}

var _cDConvertMSFToClippedLBA func(arg0 CDMSF) uint32

// CDConvertMSFToClippedLBA.
//
// See: https://developer.apple.com/documentation/iokit/1584220-cdconvertmsftoclippedlba
func CDConvertMSFToClippedLBA(arg0 CDMSF) uint32 {
	if _cDConvertMSFToClippedLBA == nil {
		panic("iokit: symbol CDConvertMSFToClippedLBA not loaded")
	}
	return _cDConvertMSFToClippedLBA(arg0)
}

var _cDConvertMSFToLBA func(arg0 CDMSF) uint32

// CDConvertMSFToLBA.
//
// See: https://developer.apple.com/documentation/iokit/1584286-cdconvertmsftolba
func CDConvertMSFToLBA(arg0 CDMSF) uint32 {
	if _cDConvertMSFToLBA == nil {
		panic("iokit: symbol CDConvertMSFToLBA not loaded")
	}
	return _cDConvertMSFToLBA(arg0)
}

var _cDConvertTrackNumberToMSF func(arg0 uint8, arg1 CDTOC) CDMSF

// CDConvertTrackNumberToMSF.
//
// See: https://developer.apple.com/documentation/iokit/1584298-cdconverttracknumbertomsf
func CDConvertTrackNumberToMSF(arg0 uint8, arg1 CDTOC) CDMSF {
	if _cDConvertTrackNumberToMSF == nil {
		panic("iokit: symbol CDConvertTrackNumberToMSF not loaded")
	}
	return _cDConvertTrackNumberToMSF(arg0, arg1)
}

var _cDTOCGetDescriptorCount func(arg0 CDTOC) uint32

// CDTOCGetDescriptorCount.
//
// See: https://developer.apple.com/documentation/iokit/1584327-cdtocgetdescriptorcount
func CDTOCGetDescriptorCount(arg0 CDTOC) uint32 {
	if _cDTOCGetDescriptorCount == nil {
		panic("iokit: symbol CDTOCGetDescriptorCount not loaded")
	}
	return _cDTOCGetDescriptorCount(arg0)
}

var _iOAccelFindAccelerator func(arg0 uintptr, arg1 uintptr, arg2 uint32) int

// IOAccelFindAccelerator.
//
// See: https://developer.apple.com/documentation/iokit/1502405-ioaccelfindaccelerator
func IOAccelFindAccelerator(arg0 uintptr, arg1 uintptr, arg2 uint32) int {
	if _iOAccelFindAccelerator == nil {
		panic("iokit: symbol IOAccelFindAccelerator not loaded")
	}
	return _iOAccelFindAccelerator(arg0, arg1, arg2)
}

var _iOAllowPowerChange func(arg0 uintptr, arg1 int) int

// IOAllowPowerChange the caller acknowledges notification of a power state change on a device it has registered for notifications for via IORegisterForSystemPower or IORegisterApp.
//
// See: https://developer.apple.com/documentation/iokit/1557064-ioallowpowerchange
func IOAllowPowerChange(arg0 uintptr, arg1 int) int {
	if _iOAllowPowerChange == nil {
		panic("iokit: symbol IOAllowPowerChange not loaded")
	}
	return _iOAllowPowerChange(arg0, arg1)
}

var _iOBSDNameMatching func(mainPort uint32, options uint32, bsdName string) corefoundation.CFMutableDictionary

// IOBSDNameMatching create a matching dictionary that specifies an IOService match based on BSD device name.
//
// See: https://developer.apple.com/documentation/iokit/1514486-iobsdnamematching
func IOBSDNameMatching(mainPort uint32, options uint32, bsdName string) corefoundation.CFMutableDictionary {
	if _iOBSDNameMatching == nil {
		panic("iokit: symbol IOBSDNameMatching not loaded")
	}
	return _iOBSDNameMatching(mainPort, options, bsdName)
}

var _iOCFSerialize func(object corefoundation.CFTypeRef, options uint64) corefoundation.CFDataRef

// IOCFSerialize.
//
// See: https://developer.apple.com/documentation/iokit/1403329-iocfserialize
func IOCFSerialize(object corefoundation.CFTypeRef, options uint64) corefoundation.CFDataRef {
	if _iOCFSerialize == nil {
		panic("iokit: symbol IOCFSerialize not loaded")
	}
	return _iOCFSerialize(object, options)
}

var _iOCFUnserialize func(buffer string, allocator corefoundation.CFAllocatorRef, options uint64, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOCFUnserialize.
//
// See: https://developer.apple.com/documentation/iokit/1514265-iocfunserialize
func IOCFUnserialize(buffer string, allocator corefoundation.CFAllocatorRef, options uint64, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOCFUnserialize == nil {
		panic("iokit: symbol IOCFUnserialize not loaded")
	}
	return _iOCFUnserialize(buffer, allocator, options, errorString)
}

var _iOCFUnserializeBinary func(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint64, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOCFUnserializeBinary.
//
// See: https://developer.apple.com/documentation/iokit/1514876-iocfunserializebinary
func IOCFUnserializeBinary(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint64, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOCFUnserializeBinary == nil {
		panic("iokit: symbol IOCFUnserializeBinary not loaded")
	}
	return _iOCFUnserializeBinary(buffer, bufferSize, allocator, options, errorString)
}

var _iOCFUnserializeWithSize func(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint64, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOCFUnserializeWithSize.
//
// See: https://developer.apple.com/documentation/iokit/1514745-iocfunserializewithsize
func IOCFUnserializeWithSize(buffer string, bufferSize unsafe.Pointer, allocator corefoundation.CFAllocatorRef, options uint64, errorString *corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOCFUnserializeWithSize == nil {
		panic("iokit: symbol IOCFUnserializeWithSize not loaded")
	}
	return _iOCFUnserializeWithSize(buffer, bufferSize, allocator, options, errorString)
}

var _iOCancelPowerChange func(arg0 uintptr, arg1 int) int

// IOCancelPowerChange the caller denies an idle system sleep power state change.
//
// See: https://developer.apple.com/documentation/iokit/1557115-iocancelpowerchange
func IOCancelPowerChange(arg0 uintptr, arg1 int) int {
	if _iOCancelPowerChange == nil {
		panic("iokit: symbol IOCancelPowerChange not loaded")
	}
	return _iOCancelPowerChange(arg0, arg1)
}

var _iOCatalogueGetData func(mainPort uint32, flag uint32, buffer string, size *uint32) int32

// IOCatalogueGetData.
//
// See: https://developer.apple.com/documentation/iokit/1514233-iocataloguegetdata
func IOCatalogueGetData(mainPort uint32, flag uint32, buffer string, size *uint32) int32 {
	if _iOCatalogueGetData == nil {
		panic("iokit: symbol IOCatalogueGetData not loaded")
	}
	return _iOCatalogueGetData(mainPort, flag, buffer, size)
}

var _iOCatalogueModuleLoaded func(mainPort uint32, name string) int32

// IOCatalogueModuleLoaded.
//
// See: https://developer.apple.com/documentation/iokit/1514886-iocataloguemoduleloaded
func IOCatalogueModuleLoaded(mainPort uint32, name string) int32 {
	if _iOCatalogueModuleLoaded == nil {
		panic("iokit: symbol IOCatalogueModuleLoaded not loaded")
	}
	return _iOCatalogueModuleLoaded(mainPort, name)
}

var _iOCatalogueReset func(mainPort uint32, flag uint32) int32

// IOCatalogueReset.
//
// See: https://developer.apple.com/documentation/iokit/1514702-iocataloguereset
func IOCatalogueReset(mainPort uint32, flag uint32) int32 {
	if _iOCatalogueReset == nil {
		panic("iokit: symbol IOCatalogueReset not loaded")
	}
	return _iOCatalogueReset(mainPort, flag)
}

var _iOCatalogueSendData func(mainPort uint32, flag uint32, buffer string, size uint32) int32

// IOCatalogueSendData.
//
// See: https://developer.apple.com/documentation/iokit/1514405-iocataloguesenddata
func IOCatalogueSendData(mainPort uint32, flag uint32, buffer string, size uint32) int32 {
	if _iOCatalogueSendData == nil {
		panic("iokit: symbol IOCatalogueSendData not loaded")
	}
	return _iOCatalogueSendData(mainPort, flag, buffer, size)
}

var _iOCatalogueTerminate func(mainPort uint32, flag uint32, description string) int32

// IOCatalogueTerminate.
//
// See: https://developer.apple.com/documentation/iokit/1514665-iocatalogueterminate
func IOCatalogueTerminate(mainPort uint32, flag uint32, description string) int32 {
	if _iOCatalogueTerminate == nil {
		panic("iokit: symbol IOCatalogueTerminate not loaded")
	}
	return _iOCatalogueTerminate(mainPort, flag, description)
}

var _iOConnectAddClient func(connect uintptr, client uintptr) int32

// IOConnectAddClient inform a connection of a second connection.
//
// See: https://developer.apple.com/documentation/iokit/1514609-ioconnectaddclient
func IOConnectAddClient(connect uintptr, client uintptr) int32 {
	if _iOConnectAddClient == nil {
		panic("iokit: symbol IOConnectAddClient not loaded")
	}
	return _iOConnectAddClient(connect, client)
}

var _iOConnectAddRef func(connect uintptr) int32

// IOConnectAddRef adds a reference to the connect handle.
//
// See: https://developer.apple.com/documentation/iokit/1514739-ioconnectaddref
func IOConnectAddRef(connect uintptr) int32 {
	if _iOConnectAddRef == nil {
		panic("iokit: symbol IOConnectAddRef not loaded")
	}
	return _iOConnectAddRef(connect)
}

var _iOConnectCallAsyncMethod func(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32

// IOConnectCallAsyncMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514418-ioconnectcallasyncmethod
func IOConnectCallAsyncMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	if _iOConnectCallAsyncMethod == nil {
		panic("iokit: symbol IOConnectCallAsyncMethod not loaded")
	}
	return _iOConnectCallAsyncMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt)
}

var _iOConnectCallAsyncScalarMethod func(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32

// IOConnectCallAsyncScalarMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514884-ioconnectcallasyncscalarmethod
func IOConnectCallAsyncScalarMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32 {
	if _iOConnectCallAsyncScalarMethod == nil {
		panic("iokit: symbol IOConnectCallAsyncScalarMethod not loaded")
	}
	return _iOConnectCallAsyncScalarMethod(connection, selector, wake_port, reference, referenceCnt, input, inputCnt, output, outputCnt)
}

var _iOConnectCallAsyncStructMethod func(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32

// IOConnectCallAsyncStructMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514403-ioconnectcallasyncstructmethod
func IOConnectCallAsyncStructMethod(connection uint32, selector uint32, wake_port uint32, reference *uint64, referenceCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	if _iOConnectCallAsyncStructMethod == nil {
		panic("iokit: symbol IOConnectCallAsyncStructMethod not loaded")
	}
	return _iOConnectCallAsyncStructMethod(connection, selector, wake_port, reference, referenceCnt, inputStruct, inputStructCnt, outputStruct, outputStructCnt)
}

var _iOConnectCallMethod func(connection uint32, selector uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32

// IOConnectCallMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514240-ioconnectcallmethod
func IOConnectCallMethod(connection uint32, selector uint32, input *uint64, inputCnt uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, output *uint64, outputCnt *uint32, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	if _iOConnectCallMethod == nil {
		panic("iokit: symbol IOConnectCallMethod not loaded")
	}
	return _iOConnectCallMethod(connection, selector, input, inputCnt, inputStruct, inputStructCnt, output, outputCnt, outputStruct, outputStructCnt)
}

var _iOConnectCallScalarMethod func(connection uint32, selector uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32

// IOConnectCallScalarMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514793-ioconnectcallscalarmethod
func IOConnectCallScalarMethod(connection uint32, selector uint32, input *uint64, inputCnt uint32, output *uint64, outputCnt *uint32) int32 {
	if _iOConnectCallScalarMethod == nil {
		panic("iokit: symbol IOConnectCallScalarMethod not loaded")
	}
	return _iOConnectCallScalarMethod(connection, selector, input, inputCnt, output, outputCnt)
}

var _iOConnectCallStructMethod func(connection uint32, selector uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32

// IOConnectCallStructMethod.
//
// See: https://developer.apple.com/documentation/iokit/1514274-ioconnectcallstructmethod
func IOConnectCallStructMethod(connection uint32, selector uint32, inputStruct unsafe.Pointer, inputStructCnt unsafe.Pointer, outputStruct unsafe.Pointer, outputStructCnt unsafe.Pointer) int32 {
	if _iOConnectCallStructMethod == nil {
		panic("iokit: symbol IOConnectCallStructMethod not loaded")
	}
	return _iOConnectCallStructMethod(connection, selector, inputStruct, inputStructCnt, outputStruct, outputStructCnt)
}

var _iOConnectGetService func(connect uintptr, service *uintptr) int32

// IOConnectGetService returns the IOService a connect handle was opened on.
//
// See: https://developer.apple.com/documentation/iokit/1514438-ioconnectgetservice
func IOConnectGetService(connect uintptr, service *uintptr) int32 {
	if _iOConnectGetService == nil {
		panic("iokit: symbol IOConnectGetService not loaded")
	}
	return _iOConnectGetService(connect, service)
}

var _iOConnectMapMemory func(connect uintptr, memoryType uint32, intoTask kernel.Task_port_t, atAddress *kernel.Mach_vm_address_t, ofSize *kernel.Mach_vm_size_t, options uint32) int32

// IOConnectMapMemory map hardware or shared memory into the caller's task.
//
// See: https://developer.apple.com/documentation/iokit/1514377-ioconnectmapmemory
func IOConnectMapMemory(connect uintptr, memoryType uint32, intoTask kernel.Task_port_t, atAddress *kernel.Mach_vm_address_t, ofSize *kernel.Mach_vm_size_t, options uint32) int32 {
	if _iOConnectMapMemory == nil {
		panic("iokit: symbol IOConnectMapMemory not loaded")
	}
	return _iOConnectMapMemory(connect, memoryType, intoTask, atAddress, ofSize, options)
}

var _iOConnectMapMemory64 func(connect uintptr, memoryType uint32, intoTask kernel.Task_port_t, atAddress *kernel.Mach_vm_address_t, ofSize *kernel.Mach_vm_size_t, options uint32) int32

// IOConnectMapMemory64 map hardware or shared memory into the caller's task.
//
// See: https://developer.apple.com/documentation/iokit/1514862-ioconnectmapmemory64
func IOConnectMapMemory64(connect uintptr, memoryType uint32, intoTask kernel.Task_port_t, atAddress *kernel.Mach_vm_address_t, ofSize *kernel.Mach_vm_size_t, options uint32) int32 {
	if _iOConnectMapMemory64 == nil {
		panic("iokit: symbol IOConnectMapMemory64 not loaded")
	}
	return _iOConnectMapMemory64(connect, memoryType, intoTask, atAddress, ofSize, options)
}

var _iOConnectRelease func(connect uintptr) int32

// IOConnectRelease remove a reference to the connect handle.
//
// See: https://developer.apple.com/documentation/iokit/1514511-ioconnectrelease
func IOConnectRelease(connect uintptr) int32 {
	if _iOConnectRelease == nil {
		panic("iokit: symbol IOConnectRelease not loaded")
	}
	return _iOConnectRelease(connect)
}

var _iOConnectSetCFProperties func(connect uintptr, properties corefoundation.CFTypeRef) int32

// IOConnectSetCFProperties set CF container based properties on a connection.
//
// See: https://developer.apple.com/documentation/iokit/1514713-ioconnectsetcfproperties
func IOConnectSetCFProperties(connect uintptr, properties corefoundation.CFTypeRef) int32 {
	if _iOConnectSetCFProperties == nil {
		panic("iokit: symbol IOConnectSetCFProperties not loaded")
	}
	return _iOConnectSetCFProperties(connect, properties)
}

var _iOConnectSetCFProperty func(connect uintptr, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32

// IOConnectSetCFProperty set a CF container based property on a connection.
//
// See: https://developer.apple.com/documentation/iokit/1514796-ioconnectsetcfproperty
func IOConnectSetCFProperty(connect uintptr, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32 {
	if _iOConnectSetCFProperty == nil {
		panic("iokit: symbol IOConnectSetCFProperty not loaded")
	}
	return _iOConnectSetCFProperty(connect, propertyName, property)
}

var _iOConnectSetNotificationPort func(connect uintptr, type_ uint32, port uint32, reference unsafe.Pointer) int32

// IOConnectSetNotificationPort set a port to receive family specific notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514541-ioconnectsetnotificationport
func IOConnectSetNotificationPort(connect uintptr, type_ uint32, port uint32, reference unsafe.Pointer) int32 {
	if _iOConnectSetNotificationPort == nil {
		panic("iokit: symbol IOConnectSetNotificationPort not loaded")
	}
	return _iOConnectSetNotificationPort(connect, type_, port, reference)
}

var _iOConnectTrap0 func(connect uintptr, index uint32) int32

// IOConnectTrap0.
//
// See: https://developer.apple.com/documentation/iokit/1514674-ioconnecttrap0
func IOConnectTrap0(connect uintptr, index uint32) int32 {
	if _iOConnectTrap0 == nil {
		panic("iokit: symbol IOConnectTrap0 not loaded")
	}
	return _iOConnectTrap0(connect, index)
}

var _iOConnectTrap1 func(connect uintptr, index uint32, p1 unsafe.Pointer) int32

// IOConnectTrap1.
//
// See: https://developer.apple.com/documentation/iokit/1514816-ioconnecttrap1
func IOConnectTrap1(connect uintptr, index uint32, p1 unsafe.Pointer) int32 {
	if _iOConnectTrap1 == nil {
		panic("iokit: symbol IOConnectTrap1 not loaded")
	}
	return _iOConnectTrap1(connect, index, p1)
}

var _iOConnectTrap2 func(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer) int32

// IOConnectTrap2.
//
// See: https://developer.apple.com/documentation/iokit/1514354-ioconnecttrap2
func IOConnectTrap2(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer) int32 {
	if _iOConnectTrap2 == nil {
		panic("iokit: symbol IOConnectTrap2 not loaded")
	}
	return _iOConnectTrap2(connect, index, p1, p2)
}

var _iOConnectTrap3 func(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) int32

// IOConnectTrap3.
//
// See: https://developer.apple.com/documentation/iokit/1514833-ioconnecttrap3
func IOConnectTrap3(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer) int32 {
	if _iOConnectTrap3 == nil {
		panic("iokit: symbol IOConnectTrap3 not loaded")
	}
	return _iOConnectTrap3(connect, index, p1, p2, p3)
}

var _iOConnectTrap4 func(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer) int32

// IOConnectTrap4.
//
// See: https://developer.apple.com/documentation/iokit/1514410-ioconnecttrap4
func IOConnectTrap4(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer) int32 {
	if _iOConnectTrap4 == nil {
		panic("iokit: symbol IOConnectTrap4 not loaded")
	}
	return _iOConnectTrap4(connect, index, p1, p2, p3, p4)
}

var _iOConnectTrap5 func(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer) int32

// IOConnectTrap5.
//
// See: https://developer.apple.com/documentation/iokit/1514346-ioconnecttrap5
func IOConnectTrap5(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer) int32 {
	if _iOConnectTrap5 == nil {
		panic("iokit: symbol IOConnectTrap5 not loaded")
	}
	return _iOConnectTrap5(connect, index, p1, p2, p3, p4, p5)
}

var _iOConnectTrap6 func(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p6 unsafe.Pointer) int32

// IOConnectTrap6.
//
// See: https://developer.apple.com/documentation/iokit/1514493-ioconnecttrap6
func IOConnectTrap6(connect uintptr, index uint32, p1 unsafe.Pointer, p2 unsafe.Pointer, p3 unsafe.Pointer, p4 unsafe.Pointer, p5 unsafe.Pointer, p6 unsafe.Pointer) int32 {
	if _iOConnectTrap6 == nil {
		panic("iokit: symbol IOConnectTrap6 not loaded")
	}
	return _iOConnectTrap6(connect, index, p1, p2, p3, p4, p5, p6)
}

var _iOConnectUnmapMemory func(connect uintptr, memoryType uint32, fromTask kernel.Task_port_t, atAddress kernel.Mach_vm_address_t) int32

// IOConnectUnmapMemory remove a mapping made with IOConnectMapMemory.
//
// See: https://developer.apple.com/documentation/iokit/1514527-ioconnectunmapmemory
func IOConnectUnmapMemory(connect uintptr, memoryType uint32, fromTask kernel.Task_port_t, atAddress kernel.Mach_vm_address_t) int32 {
	if _iOConnectUnmapMemory == nil {
		panic("iokit: symbol IOConnectUnmapMemory not loaded")
	}
	return _iOConnectUnmapMemory(connect, memoryType, fromTask, atAddress)
}

var _iOConnectUnmapMemory64 func(connect uintptr, memoryType uint32, fromTask kernel.Task_port_t, atAddress kernel.Mach_vm_address_t) int32

// IOConnectUnmapMemory64 remove a mapping made with IOConnectMapMemory64.
//
// See: https://developer.apple.com/documentation/iokit/1514760-ioconnectunmapmemory64
func IOConnectUnmapMemory64(connect uintptr, memoryType uint32, fromTask kernel.Task_port_t, atAddress kernel.Mach_vm_address_t) int32 {
	if _iOConnectUnmapMemory64 == nil {
		panic("iokit: symbol IOConnectUnmapMemory64 not loaded")
	}
	return _iOConnectUnmapMemory64(connect, memoryType, fromTask, atAddress)
}

var _iOCopySystemLoadAdvisoryDetailed func() corefoundation.CFDictionaryRef

// IOCopySystemLoadAdvisoryDetailed indicates how user activity, battery level, and thermal level each contribute to the overall "SystemLoadAdvisory" level.
//
// See: https://developer.apple.com/documentation/iokit/1557099-iocopysystemloadadvisorydetailed
func IOCopySystemLoadAdvisoryDetailed() corefoundation.CFDictionaryRef {
	if _iOCopySystemLoadAdvisoryDetailed == nil {
		panic("iokit: symbol IOCopySystemLoadAdvisoryDetailed not loaded")
	}
	return _iOCopySystemLoadAdvisoryDetailed()
}

var _iOCreatePlugInInterfaceForService func(service uintptr, pluginType corefoundation.CFUUID, interfaceType corefoundation.CFUUID, theInterface *IOCFPlugInInterface, theScore *uintptr) int32

// IOCreatePlugInInterfaceForService.
//
// See: https://developer.apple.com/documentation/iokit/1412429-iocreateplugininterfaceforservic
func IOCreatePlugInInterfaceForService(service uintptr, pluginType corefoundation.CFUUID, interfaceType corefoundation.CFUUID, theInterface *IOCFPlugInInterface, theScore *uintptr) int32 {
	if _iOCreatePlugInInterfaceForService == nil {
		panic("iokit: symbol IOCreatePlugInInterfaceForService not loaded")
	}
	return _iOCreatePlugInInterfaceForService(service, pluginType, interfaceType, theInterface, theScore)
}

var _iOCreateReceivePort func(msgType uint32, recvPort *uint32) int32

// IOCreateReceivePort creates and returns a mach port suitable for receiving IOKit messages of the specified type.
//
// See: https://developer.apple.com/documentation/iokit/1514698-iocreatereceiveport
func IOCreateReceivePort(msgType uint32, recvPort *uint32) int32 {
	if _iOCreateReceivePort == nil {
		panic("iokit: symbol IOCreateReceivePort not loaded")
	}
	return _iOCreateReceivePort(msgType, recvPort)
}

var _iODataQueueAllocateNotificationPort func() uint32

// IODataQueueAllocateNotificationPort allocates and returns a new mach port able to receive data available notifications from an IODataQueue.
//
// See: https://developer.apple.com/documentation/iokit/1514495-iodataqueueallocatenotificationp
func IODataQueueAllocateNotificationPort() uint32 {
	if _iODataQueueAllocateNotificationPort == nil {
		panic("iokit: symbol IODataQueueAllocateNotificationPort not loaded")
	}
	return _iODataQueueAllocateNotificationPort()
}

var _iODataQueueDataAvailable func(dataQueue *IODataQueueMemory) bool

// IODataQueueDataAvailable used to determine if more data is avilable on the queue.
//
// See: https://developer.apple.com/documentation/iokit/1514386-iodataqueuedataavailable
func IODataQueueDataAvailable(dataQueue *IODataQueueMemory) bool {
	if _iODataQueueDataAvailable == nil {
		panic("iokit: symbol IODataQueueDataAvailable not loaded")
	}
	return _iODataQueueDataAvailable(dataQueue)
}

var _iODataQueueDequeue func(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize *uint32) int

// IODataQueueDequeue dequeues the next available entry on the queue and copies it into the given data pointer.
//
// See: https://developer.apple.com/documentation/iokit/1514287-iodataqueuedequeue
func IODataQueueDequeue(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize *uint32) int {
	if _iODataQueueDequeue == nil {
		panic("iokit: symbol IODataQueueDequeue not loaded")
	}
	return _iODataQueueDequeue(dataQueue, data, dataSize)
}

var _iODataQueueEnqueue func(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize uint32) int

// IODataQueueEnqueue enqueues a new entry on the queue.
//
// See: https://developer.apple.com/documentation/iokit/1514482-iodataqueueenqueue
func IODataQueueEnqueue(dataQueue *IODataQueueMemory, data unsafe.Pointer, dataSize uint32) int {
	if _iODataQueueEnqueue == nil {
		panic("iokit: symbol IODataQueueEnqueue not loaded")
	}
	return _iODataQueueEnqueue(dataQueue, data, dataSize)
}

var _iODataQueuePeek func(dataQueue *IODataQueueMemory) IODataQueueEntry

// IODataQueuePeek used to peek at the next entry on the queue.
//
// See: https://developer.apple.com/documentation/iokit/1514649-iodataqueuepeek
func IODataQueuePeek(dataQueue *IODataQueueMemory) IODataQueueEntry {
	if _iODataQueuePeek == nil {
		panic("iokit: symbol IODataQueuePeek not loaded")
	}
	return _iODataQueuePeek(dataQueue)
}

var _iODataQueueSetNotificationPort func(dataQueue *IODataQueueMemory, notifyPort uint32) int

// IODataQueueSetNotificationPort creates a simple mach message targeting the mach port specified in port.
//
// See: https://developer.apple.com/documentation/iokit/1514301-iodataqueuesetnotificationport
func IODataQueueSetNotificationPort(dataQueue *IODataQueueMemory, notifyPort uint32) int {
	if _iODataQueueSetNotificationPort == nil {
		panic("iokit: symbol IODataQueueSetNotificationPort not loaded")
	}
	return _iODataQueueSetNotificationPort(dataQueue, notifyPort)
}

var _iODataQueueWaitForAvailableData func(dataQueue *IODataQueueMemory, notificationPort uint32) int

// IODataQueueWaitForAvailableData wait for an incoming dataAvailable message on the given notifyPort.
//
// See: https://developer.apple.com/documentation/iokit/1514696-iodataqueuewaitforavailabledata
func IODataQueueWaitForAvailableData(dataQueue *IODataQueueMemory, notificationPort uint32) int {
	if _iODataQueueWaitForAvailableData == nil {
		panic("iokit: symbol IODataQueueWaitForAvailableData not loaded")
	}
	return _iODataQueueWaitForAvailableData(dataQueue, notificationPort)
}

var _iODeregisterApp func(arg0 uintptr) int

// IODeregisterApp disconnects the caller from an IOService after receiving power state change notifications from the IOService.
//
// See: https://developer.apple.com/documentation/iokit/1557129-ioderegisterapp
func IODeregisterApp(arg0 uintptr) int {
	if _iODeregisterApp == nil {
		panic("iokit: symbol IODeregisterApp not loaded")
	}
	return _iODeregisterApp(arg0)
}

var _iODeregisterForSystemPower func(arg0 uintptr) int

// IODeregisterForSystemPower disconnects the caller from the Root Power Domain IOService after receiving system power state change notifications.
//
// See: https://developer.apple.com/documentation/iokit/1557132-ioderegisterforsystempower
func IODeregisterForSystemPower(arg0 uintptr) int {
	if _iODeregisterForSystemPower == nil {
		panic("iokit: symbol IODeregisterForSystemPower not loaded")
	}
	return _iODeregisterForSystemPower(arg0)
}

var _iODestroyPlugInInterface func(interface_ *IOCFPlugInInterface) int32

// IODestroyPlugInInterface.
//
// See: https://developer.apple.com/documentation/iokit/1412425-iodestroyplugininterface
func IODestroyPlugInInterface(interface_ *IOCFPlugInInterface) int32 {
	if _iODestroyPlugInInterface == nil {
		panic("iokit: symbol IODestroyPlugInInterface not loaded")
	}
	return _iODestroyPlugInInterface(interface_)
}

var _iODispatchCalloutFromMessage func(unused unsafe.Pointer, msg unsafe.Pointer, reference unsafe.Pointer)

// IODispatchCalloutFromMessage dispatches callback notifications from a mach message.
//
// See: https://developer.apple.com/documentation/iokit/1514775-iodispatchcalloutfrommessage
func IODispatchCalloutFromMessage(unused unsafe.Pointer, msg unsafe.Pointer, reference unsafe.Pointer) {
	if _iODispatchCalloutFromMessage == nil {
		panic("iokit: symbol IODispatchCalloutFromMessage not loaded")
	}
	_iODispatchCalloutFromMessage(unused, msg, reference)
}

var _iODisplayCommitParameters func(arg0 uintptr, arg1 uint32) int

// IODisplayCommitParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574906-iodisplaycommitparameters
func IODisplayCommitParameters(arg0 uintptr, arg1 uint32) int {
	if _iODisplayCommitParameters == nil {
		panic("iokit: symbol IODisplayCommitParameters not loaded")
	}
	return _iODisplayCommitParameters(arg0, arg1)
}

var _iODisplayCopyFloatParameters func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFDictionaryRef) int

// IODisplayCopyFloatParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574891-iodisplaycopyfloatparameters
func IODisplayCopyFloatParameters(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFDictionaryRef) int {
	if _iODisplayCopyFloatParameters == nil {
		panic("iokit: symbol IODisplayCopyFloatParameters not loaded")
	}
	return _iODisplayCopyFloatParameters(arg0, arg1, arg2)
}

var _iODisplayCopyParameters func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFDictionaryRef) int

// IODisplayCopyParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574865-iodisplaycopyparameters
func IODisplayCopyParameters(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFDictionaryRef) int {
	if _iODisplayCopyParameters == nil {
		panic("iokit: symbol IODisplayCopyParameters not loaded")
	}
	return _iODisplayCopyParameters(arg0, arg1, arg2)
}

var _iODisplayCreateInfoDictionary func(arg0 uintptr, arg1 uint32) corefoundation.CFDictionaryRef

// IODisplayCreateInfoDictionary create a CFDictionary with information about display hardware.
//
// See: https://developer.apple.com/documentation/iokit/1574917-iodisplaycreateinfodictionary
func IODisplayCreateInfoDictionary(arg0 uintptr, arg1 uint32) corefoundation.CFDictionaryRef {
	if _iODisplayCreateInfoDictionary == nil {
		panic("iokit: symbol IODisplayCreateInfoDictionary not loaded")
	}
	return _iODisplayCreateInfoDictionary(arg0, arg1)
}

var _iODisplayForFramebuffer func(arg0 uintptr, arg1 uint32) uintptr

// IODisplayForFramebuffer.
//
// See: https://developer.apple.com/documentation/iokit/1574899-iodisplayforframebuffer
func IODisplayForFramebuffer(arg0 uintptr, arg1 uint32) uintptr {
	if _iODisplayForFramebuffer == nil {
		panic("iokit: symbol IODisplayForFramebuffer not loaded")
	}
	return _iODisplayForFramebuffer(arg0, arg1)
}

var _iODisplayGetFloatParameter func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 float32) int

// IODisplayGetFloatParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574900-iodisplaygetfloatparameter
func IODisplayGetFloatParameter(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 float32) int {
	if _iODisplayGetFloatParameter == nil {
		panic("iokit: symbol IODisplayGetFloatParameter not loaded")
	}
	return _iODisplayGetFloatParameter(arg0, arg1, arg2, arg3)
}

var _iODisplayGetIntegerRangeParameter func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 int32, arg4 int32, arg5 int32) int

// IODisplayGetIntegerRangeParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574908-iodisplaygetintegerrangeparamete
func IODisplayGetIntegerRangeParameter(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 int32, arg4 int32, arg5 int32) int {
	if _iODisplayGetIntegerRangeParameter == nil {
		panic("iokit: symbol IODisplayGetIntegerRangeParameter not loaded")
	}
	return _iODisplayGetIntegerRangeParameter(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _iODisplayMatchDictionaries func(arg0 corefoundation.CFDictionaryRef, arg1 corefoundation.CFDictionaryRef, arg2 uint32) int32

// IODisplayMatchDictionaries match two display information dictionaries to see if they are for the same display.
//
// See: https://developer.apple.com/documentation/iokit/1574879-iodisplaymatchdictionaries
func IODisplayMatchDictionaries(arg0 corefoundation.CFDictionaryRef, arg1 corefoundation.CFDictionaryRef, arg2 uint32) int32 {
	if _iODisplayMatchDictionaries == nil {
		panic("iokit: symbol IODisplayMatchDictionaries not loaded")
	}
	return _iODisplayMatchDictionaries(arg0, arg1, arg2)
}

var _iODisplaySetFloatParameter func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 float32) int

// IODisplaySetFloatParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574926-iodisplaysetfloatparameter
func IODisplaySetFloatParameter(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 float32) int {
	if _iODisplaySetFloatParameter == nil {
		panic("iokit: symbol IODisplaySetFloatParameter not loaded")
	}
	return _iODisplaySetFloatParameter(arg0, arg1, arg2, arg3)
}

var _iODisplaySetIntegerParameter func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 int32) int

// IODisplaySetIntegerParameter.
//
// See: https://developer.apple.com/documentation/iokit/1574915-iodisplaysetintegerparameter
func IODisplaySetIntegerParameter(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFStringRef, arg3 int32) int {
	if _iODisplaySetIntegerParameter == nil {
		panic("iokit: symbol IODisplaySetIntegerParameter not loaded")
	}
	return _iODisplaySetIntegerParameter(arg0, arg1, arg2, arg3)
}

var _iODisplaySetParameters func(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFDictionaryRef) int

// IODisplaySetParameters.
//
// See: https://developer.apple.com/documentation/iokit/1574878-iodisplaysetparameters
func IODisplaySetParameters(arg0 uintptr, arg1 uint32, arg2 corefoundation.CFDictionaryRef) int {
	if _iODisplaySetParameters == nil {
		panic("iokit: symbol IODisplaySetParameters not loaded")
	}
	return _iODisplaySetParameters(arg0, arg1, arg2)
}

var _iOFBCopyI2CInterfaceForBus func(arg0 uintptr, arg1 uint32, arg2 uintptr) int

// IOFBCopyI2CInterfaceForBus returns an instance of an I2C bus interface, associated with an IOFramebuffer instance / bus index pair.
//
// See: https://developer.apple.com/documentation/iokit/1410345-iofbcopyi2cinterfaceforbus
func IOFBCopyI2CInterfaceForBus(arg0 uintptr, arg1 uint32, arg2 uintptr) int {
	if _iOFBCopyI2CInterfaceForBus == nil {
		panic("iokit: symbol IOFBCopyI2CInterfaceForBus not loaded")
	}
	return _iOFBCopyI2CInterfaceForBus(arg0, arg1, arg2)
}

var _iOFBGetI2CInterfaceCount func(arg0 uintptr, arg1 uint) int

// IOFBGetI2CInterfaceCount returns a count of I2C interfaces available associated with an IOFramebuffer instance.
//
// See: https://developer.apple.com/documentation/iokit/1410333-iofbgeti2cinterfacecount
func IOFBGetI2CInterfaceCount(arg0 uintptr, arg1 uint) int {
	if _iOFBGetI2CInterfaceCount == nil {
		panic("iokit: symbol IOFBGetI2CInterfaceCount not loaded")
	}
	return _iOFBGetI2CInterfaceCount(arg0, arg1)
}

var _iOFramebufferOpen func(arg0 uintptr, arg1 kernel.Task_port_t, arg2 uint, arg3 uintptr) int32

// IOFramebufferOpen.
//
// See: https://developer.apple.com/documentation/iokit/1574872-ioframebufferopen
func IOFramebufferOpen(arg0 uintptr, arg1 kernel.Task_port_t, arg2 uint, arg3 uintptr) int32 {
	if _iOFramebufferOpen == nil {
		panic("iokit: symbol IOFramebufferOpen not loaded")
	}
	return _iOFramebufferOpen(arg0, arg1, arg2, arg3)
}

var _iOGetSystemLoadAdvisory func() IOSystemLoadAdvisoryLevel

// IOGetSystemLoadAdvisory returns a hint about whether now would be a good time to perform time-insensitive work.
//
// See: https://developer.apple.com/documentation/iokit/1557110-iogetsystemloadadvisory
func IOGetSystemLoadAdvisory() IOSystemLoadAdvisoryLevel {
	if _iOGetSystemLoadAdvisory == nil {
		panic("iokit: symbol IOGetSystemLoadAdvisory not loaded")
	}
	return _iOGetSystemLoadAdvisory()
}

var _iOHIDCheckAccess func(arg0 IOHIDRequestType) IOHIDAccessType

// IOHIDCheckAccess.
//
// See: https://developer.apple.com/documentation/iokit/3181573-iohidcheckaccess
func IOHIDCheckAccess(arg0 IOHIDRequestType) IOHIDAccessType {
	if _iOHIDCheckAccess == nil {
		panic("iokit: symbol IOHIDCheckAccess not loaded")
	}
	return _iOHIDCheckAccess(arg0)
}

var _iOHIDCopyCFTypeParameter func(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) int32

// IOHIDCopyCFTypeParameter.
//
// See: https://developer.apple.com/documentation/iokit/1555415-iohidcopycftypeparameter
func IOHIDCopyCFTypeParameter(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) int32 {
	if _iOHIDCopyCFTypeParameter == nil {
		panic("iokit: symbol IOHIDCopyCFTypeParameter not loaded")
	}
	return _iOHIDCopyCFTypeParameter(arg0, arg1, arg2)
}

var _iOHIDCreateSharedMemory func(arg0 uintptr, arg1 uint) int32

// IOHIDCreateSharedMemory.
//
// See: https://developer.apple.com/documentation/iokit/1555393-iohidcreatesharedmemory
func IOHIDCreateSharedMemory(arg0 uintptr, arg1 uint) int32 {
	if _iOHIDCreateSharedMemory == nil {
		panic("iokit: symbol IOHIDCreateSharedMemory not loaded")
	}
	return _iOHIDCreateSharedMemory(arg0, arg1)
}

var _iOHIDDeviceActivate func(arg0 IOHIDDeviceRef)

// IOHIDDeviceActivate.
//
// See: https://developer.apple.com/documentation/iokit/3042781-iohiddeviceactivate
func IOHIDDeviceActivate(arg0 IOHIDDeviceRef) {
	if _iOHIDDeviceActivate == nil {
		panic("iokit: symbol IOHIDDeviceActivate not loaded")
	}
	_iOHIDDeviceActivate(arg0)
}

var _iOHIDDeviceCancel func(arg0 IOHIDDeviceRef)

// IOHIDDeviceCancel.
//
// See: https://developer.apple.com/documentation/iokit/3042782-iohiddevicecancel
func IOHIDDeviceCancel(arg0 IOHIDDeviceRef) {
	if _iOHIDDeviceCancel == nil {
		panic("iokit: symbol IOHIDDeviceCancel not loaded")
	}
	_iOHIDDeviceCancel(arg0)
}

var _iOHIDDeviceClose func(arg0 IOHIDDeviceRef, arg1 uint32) int

// IOHIDDeviceClose closes communication with a HID device.
//
// See: https://developer.apple.com/documentation/iokit/1588668-iohiddeviceclose
func IOHIDDeviceClose(arg0 IOHIDDeviceRef, arg1 uint32) int {
	if _iOHIDDeviceClose == nil {
		panic("iokit: symbol IOHIDDeviceClose not loaded")
	}
	return _iOHIDDeviceClose(arg0, arg1)
}

var _iOHIDDeviceConformsTo func(arg0 IOHIDDeviceRef, arg1 uint32, arg2 uint32) bool

// IOHIDDeviceConformsTo convenience function that scans the Application Collection elements to see if it conforms to the provided usagePage and usage.
//
// See: https://developer.apple.com/documentation/iokit/1588665-iohiddeviceconformsto
func IOHIDDeviceConformsTo(arg0 IOHIDDeviceRef, arg1 uint32, arg2 uint32) bool {
	if _iOHIDDeviceConformsTo == nil {
		panic("iokit: symbol IOHIDDeviceConformsTo not loaded")
	}
	return _iOHIDDeviceConformsTo(arg0, arg1, arg2)
}

var _iOHIDDeviceCopyMatchingElements func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef, arg2 uint32) corefoundation.CFArrayRef

// IOHIDDeviceCopyMatchingElements obtains HID elements that match the criteria contained in the matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1588671-iohiddevicecopymatchingelements
func IOHIDDeviceCopyMatchingElements(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef, arg2 uint32) corefoundation.CFArrayRef {
	if _iOHIDDeviceCopyMatchingElements == nil {
		panic("iokit: symbol IOHIDDeviceCopyMatchingElements not loaded")
	}
	return _iOHIDDeviceCopyMatchingElements(arg0, arg1, arg2)
}

var _iOHIDDeviceCopyValueMultiple func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFArrayRef, arg2 corefoundation.CFDictionaryRef) int

// IOHIDDeviceCopyValueMultiple copies a values for multiple elements.
//
// See: https://developer.apple.com/documentation/iokit/1588652-iohiddevicecopyvaluemultiple
func IOHIDDeviceCopyValueMultiple(arg0 IOHIDDeviceRef, arg1 corefoundation.CFArrayRef, arg2 corefoundation.CFDictionaryRef) int {
	if _iOHIDDeviceCopyValueMultiple == nil {
		panic("iokit: symbol IOHIDDeviceCopyValueMultiple not loaded")
	}
	return _iOHIDDeviceCopyValueMultiple(arg0, arg1, arg2)
}

var _iOHIDDeviceCopyValueMultipleWithCallback func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFArrayRef, arg2 corefoundation.CFDictionaryRef, arg3 float64, arg4 unsafe.Pointer) int

// IOHIDDeviceCopyValueMultipleWithCallback copies a values for multiple elements and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588655-iohiddevicecopyvaluemultiplewith
func IOHIDDeviceCopyValueMultipleWithCallback(arg0 IOHIDDeviceRef, arg1 corefoundation.CFArrayRef, arg2 corefoundation.CFDictionaryRef, arg3 float64, arg4 unsafe.Pointer) int {
	if _iOHIDDeviceCopyValueMultipleWithCallback == nil {
		panic("iokit: symbol IOHIDDeviceCopyValueMultipleWithCallback not loaded")
	}
	return _iOHIDDeviceCopyValueMultipleWithCallback(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDDeviceCreate func(arg0 corefoundation.CFAllocatorRef, arg1 uintptr) IOHIDDeviceRef

// IOHIDDeviceCreate creates an element from an io_service_t.
//
// See: https://developer.apple.com/documentation/iokit/1588663-iohiddevicecreate
func IOHIDDeviceCreate(arg0 corefoundation.CFAllocatorRef, arg1 uintptr) IOHIDDeviceRef {
	if _iOHIDDeviceCreate == nil {
		panic("iokit: symbol IOHIDDeviceCreate not loaded")
	}
	return _iOHIDDeviceCreate(arg0, arg1)
}

var _iOHIDDeviceGetProperty func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOHIDDeviceGetProperty obtains a property from an IOHIDDevice.
//
// See: https://developer.apple.com/documentation/iokit/1588648-iohiddevicegetproperty
func IOHIDDeviceGetProperty(arg0 IOHIDDeviceRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOHIDDeviceGetProperty == nil {
		panic("iokit: symbol IOHIDDeviceGetProperty not loaded")
	}
	return _iOHIDDeviceGetProperty(arg0, arg1)
}

var _iOHIDDeviceGetReport func(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int) int

// IOHIDDeviceGetReport obtains a report from the device.
//
// See: https://developer.apple.com/documentation/iokit/1588659-iohiddevicegetreport
func IOHIDDeviceGetReport(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int) int {
	if _iOHIDDeviceGetReport == nil {
		panic("iokit: symbol IOHIDDeviceGetReport not loaded")
	}
	return _iOHIDDeviceGetReport(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDDeviceGetReportWithCallback func(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int, arg5 float64, arg6 unsafe.Pointer) int

// IOHIDDeviceGetReportWithCallback obtains a report from the device.
//
// See: https://developer.apple.com/documentation/iokit/1588662-iohiddevicegetreportwithcallback
func IOHIDDeviceGetReportWithCallback(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int, arg5 float64, arg6 unsafe.Pointer) int {
	if _iOHIDDeviceGetReportWithCallback == nil {
		panic("iokit: symbol IOHIDDeviceGetReportWithCallback not loaded")
	}
	return _iOHIDDeviceGetReportWithCallback(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _iOHIDDeviceGetService func(arg0 IOHIDDeviceRef) uintptr

// IOHIDDeviceGetService returns the io_service_t for an IOHIDDevice, if it has one.
//
// See: https://developer.apple.com/documentation/iokit/1588646-iohiddevicegetservice
func IOHIDDeviceGetService(arg0 IOHIDDeviceRef) uintptr {
	if _iOHIDDeviceGetService == nil {
		panic("iokit: symbol IOHIDDeviceGetService not loaded")
	}
	return _iOHIDDeviceGetService(arg0)
}

var _iOHIDDeviceGetTypeID func() uint

// IOHIDDeviceGetTypeID returns the type identifier of all IOHIDDevice instances.
//
// See: https://developer.apple.com/documentation/iokit/1588664-iohiddevicegettypeid
func IOHIDDeviceGetTypeID() uint {
	if _iOHIDDeviceGetTypeID == nil {
		panic("iokit: symbol IOHIDDeviceGetTypeID not loaded")
	}
	return _iOHIDDeviceGetTypeID()
}

var _iOHIDDeviceGetValue func(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef) int

// IOHIDDeviceGetValue gets a value for an element.
//
// See: https://developer.apple.com/documentation/iokit/1588657-iohiddevicegetvalue
func IOHIDDeviceGetValue(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef) int {
	if _iOHIDDeviceGetValue == nil {
		panic("iokit: symbol IOHIDDeviceGetValue not loaded")
	}
	return _iOHIDDeviceGetValue(arg0, arg1, arg2)
}

var _iOHIDDeviceGetValueWithCallback func(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 float64, arg4 unsafe.Pointer) int

// IOHIDDeviceGetValueWithCallback gets a value for an element and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588647-iohiddevicegetvaluewithcallback
func IOHIDDeviceGetValueWithCallback(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 float64, arg4 unsafe.Pointer) int {
	if _iOHIDDeviceGetValueWithCallback == nil {
		panic("iokit: symbol IOHIDDeviceGetValueWithCallback not loaded")
	}
	return _iOHIDDeviceGetValueWithCallback(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDDeviceGetValueWithOptions func(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 uint32) int

// IOHIDDeviceGetValueWithOptions.
//
// See: https://developer.apple.com/documentation/iokit/2937933-iohiddevicegetvaluewithoptions
func IOHIDDeviceGetValueWithOptions(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 uint32) int {
	if _iOHIDDeviceGetValueWithOptions == nil {
		panic("iokit: symbol IOHIDDeviceGetValueWithOptions not loaded")
	}
	return _iOHIDDeviceGetValueWithOptions(arg0, arg1, arg2, arg3)
}

var _iOHIDDeviceOpen func(arg0 IOHIDDeviceRef, arg1 uint32) int

// IOHIDDeviceOpen opens a HID device for communication.
//
// See: https://developer.apple.com/documentation/iokit/1588670-iohiddeviceopen
func IOHIDDeviceOpen(arg0 IOHIDDeviceRef, arg1 uint32) int {
	if _iOHIDDeviceOpen == nil {
		panic("iokit: symbol IOHIDDeviceOpen not loaded")
	}
	return _iOHIDDeviceOpen(arg0, arg1)
}

var _iOHIDDeviceRegisterInputReportCallback func(arg0 IOHIDDeviceRef, arg1 uint8, arg2 int, arg3 unsafe.Pointer)

// IOHIDDeviceRegisterInputReportCallback registers a callback to be used when an input report is issued by the device.
//
// See: https://developer.apple.com/documentation/iokit/1588666-iohiddeviceregisterinputreportca
func IOHIDDeviceRegisterInputReportCallback(arg0 IOHIDDeviceRef, arg1 uint8, arg2 int, arg3 unsafe.Pointer) {
	if _iOHIDDeviceRegisterInputReportCallback == nil {
		panic("iokit: symbol IOHIDDeviceRegisterInputReportCallback not loaded")
	}
	_iOHIDDeviceRegisterInputReportCallback(arg0, arg1, arg2, arg3)
}

var _iOHIDDeviceRegisterInputReportWithTimeStampCallback func(arg0 IOHIDDeviceRef, arg1 uint8, arg2 int, arg3 unsafe.Pointer)

// IOHIDDeviceRegisterInputReportWithTimeStampCallback.
//
// See: https://developer.apple.com/documentation/iokit/1588649-iohiddeviceregisterinputreportwi
func IOHIDDeviceRegisterInputReportWithTimeStampCallback(arg0 IOHIDDeviceRef, arg1 uint8, arg2 int, arg3 unsafe.Pointer) {
	if _iOHIDDeviceRegisterInputReportWithTimeStampCallback == nil {
		panic("iokit: symbol IOHIDDeviceRegisterInputReportWithTimeStampCallback not loaded")
	}
	_iOHIDDeviceRegisterInputReportWithTimeStampCallback(arg0, arg1, arg2, arg3)
}

var _iOHIDDeviceRegisterInputValueCallback func(arg0 IOHIDDeviceRef, arg1 unsafe.Pointer)

// IOHIDDeviceRegisterInputValueCallback registers a callback to be used when an input value is issued by the device.
//
// See: https://developer.apple.com/documentation/iokit/1588672-iohiddeviceregisterinputvaluecal
func IOHIDDeviceRegisterInputValueCallback(arg0 IOHIDDeviceRef, arg1 unsafe.Pointer) {
	if _iOHIDDeviceRegisterInputValueCallback == nil {
		panic("iokit: symbol IOHIDDeviceRegisterInputValueCallback not loaded")
	}
	_iOHIDDeviceRegisterInputValueCallback(arg0, arg1)
}

var _iOHIDDeviceRegisterRemovalCallback func(arg0 IOHIDDeviceRef, arg1 unsafe.Pointer)

// IOHIDDeviceRegisterRemovalCallback registers a callback to be used when a IOHIDDevice is removed.
//
// See: https://developer.apple.com/documentation/iokit/1588673-iohiddeviceregisterremovalcallba
func IOHIDDeviceRegisterRemovalCallback(arg0 IOHIDDeviceRef, arg1 unsafe.Pointer) {
	if _iOHIDDeviceRegisterRemovalCallback == nil {
		panic("iokit: symbol IOHIDDeviceRegisterRemovalCallback not loaded")
	}
	_iOHIDDeviceRegisterRemovalCallback(arg0, arg1)
}

var _iOHIDDeviceScheduleWithRunLoop func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDDeviceScheduleWithRunLoop schedules HID device with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1588660-iohiddeviceschedulewithrunloop
func IOHIDDeviceScheduleWithRunLoop(arg0 IOHIDDeviceRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDDeviceScheduleWithRunLoop == nil {
		panic("iokit: symbol IOHIDDeviceScheduleWithRunLoop not loaded")
	}
	_iOHIDDeviceScheduleWithRunLoop(arg0, arg1, arg2)
}

var _iOHIDDeviceSetCancelHandler func(arg0 IOHIDDeviceRef, arg1 unsafe.Pointer)

// IOHIDDeviceSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3042783-iohiddevicesetcancelhandler
func IOHIDDeviceSetCancelHandler(arg0 IOHIDDeviceRef, arg1 unsafe.Pointer) {
	if _iOHIDDeviceSetCancelHandler == nil {
		panic("iokit: symbol IOHIDDeviceSetCancelHandler not loaded")
	}
	_iOHIDDeviceSetCancelHandler(arg0, arg1)
}

var _iOHIDDeviceSetDispatchQueue func(arg0 IOHIDDeviceRef, arg1 uintptr)

// IOHIDDeviceSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3042784-iohiddevicesetdispatchqueue
func IOHIDDeviceSetDispatchQueue(arg0 IOHIDDeviceRef, arg1 dispatch.Queue) {
	if _iOHIDDeviceSetDispatchQueue == nil {
		panic("iokit: symbol IOHIDDeviceSetDispatchQueue not loaded")
	}
	_iOHIDDeviceSetDispatchQueue(arg0, uintptr(arg1.Handle()))
}

var _iOHIDDeviceSetInputValueMatching func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef)

// IOHIDDeviceSetInputValueMatching sets matching criteria for input values received via IOHIDDeviceRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1588654-iohiddevicesetinputvaluematching
func IOHIDDeviceSetInputValueMatching(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef) {
	if _iOHIDDeviceSetInputValueMatching == nil {
		panic("iokit: symbol IOHIDDeviceSetInputValueMatching not loaded")
	}
	_iOHIDDeviceSetInputValueMatching(arg0, arg1)
}

var _iOHIDDeviceSetInputValueMatchingMultiple func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFArrayRef)

// IOHIDDeviceSetInputValueMatchingMultiple sets multiple matching criteria for input values received via IOHIDDeviceRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1588645-iohiddevicesetinputvaluematching
func IOHIDDeviceSetInputValueMatchingMultiple(arg0 IOHIDDeviceRef, arg1 corefoundation.CFArrayRef) {
	if _iOHIDDeviceSetInputValueMatchingMultiple == nil {
		panic("iokit: symbol IOHIDDeviceSetInputValueMatchingMultiple not loaded")
	}
	_iOHIDDeviceSetInputValueMatchingMultiple(arg0, arg1)
}

var _iOHIDDeviceSetProperty func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool

// IOHIDDeviceSetProperty sets a property for an IOHIDDevice.
//
// See: https://developer.apple.com/documentation/iokit/1588653-iohiddevicesetproperty
func IOHIDDeviceSetProperty(arg0 IOHIDDeviceRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool {
	if _iOHIDDeviceSetProperty == nil {
		panic("iokit: symbol IOHIDDeviceSetProperty not loaded")
	}
	return _iOHIDDeviceSetProperty(arg0, arg1, arg2)
}

var _iOHIDDeviceSetReport func(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int) int

// IOHIDDeviceSetReport sends a report to the device.
//
// See: https://developer.apple.com/documentation/iokit/1588656-iohiddevicesetreport
func IOHIDDeviceSetReport(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int) int {
	if _iOHIDDeviceSetReport == nil {
		panic("iokit: symbol IOHIDDeviceSetReport not loaded")
	}
	return _iOHIDDeviceSetReport(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDDeviceSetReportWithCallback func(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int, arg5 float64, arg6 unsafe.Pointer) int

// IOHIDDeviceSetReportWithCallback sends a report to the device.
//
// See: https://developer.apple.com/documentation/iokit/1588661-iohiddevicesetreportwithcallback
func IOHIDDeviceSetReportWithCallback(arg0 IOHIDDeviceRef, arg1 IOHIDReportType, arg2 int, arg3 uint8, arg4 int, arg5 float64, arg6 unsafe.Pointer) int {
	if _iOHIDDeviceSetReportWithCallback == nil {
		panic("iokit: symbol IOHIDDeviceSetReportWithCallback not loaded")
	}
	return _iOHIDDeviceSetReportWithCallback(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _iOHIDDeviceSetValue func(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef) int

// IOHIDDeviceSetValue sets a value for an element.
//
// See: https://developer.apple.com/documentation/iokit/1588651-iohiddevicesetvalue
func IOHIDDeviceSetValue(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef) int {
	if _iOHIDDeviceSetValue == nil {
		panic("iokit: symbol IOHIDDeviceSetValue not loaded")
	}
	return _iOHIDDeviceSetValue(arg0, arg1, arg2)
}

var _iOHIDDeviceSetValueMultiple func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef) int

// IOHIDDeviceSetValueMultiple sets multiple values for multiple elements.
//
// See: https://developer.apple.com/documentation/iokit/1588669-iohiddevicesetvaluemultiple
func IOHIDDeviceSetValueMultiple(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef) int {
	if _iOHIDDeviceSetValueMultiple == nil {
		panic("iokit: symbol IOHIDDeviceSetValueMultiple not loaded")
	}
	return _iOHIDDeviceSetValueMultiple(arg0, arg1)
}

var _iOHIDDeviceSetValueMultipleWithCallback func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef, arg2 float64, arg3 unsafe.Pointer) int

// IOHIDDeviceSetValueMultipleWithCallback sets multiple values for multiple elements and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588658-iohiddevicesetvaluemultiplewithc
func IOHIDDeviceSetValueMultipleWithCallback(arg0 IOHIDDeviceRef, arg1 corefoundation.CFDictionaryRef, arg2 float64, arg3 unsafe.Pointer) int {
	if _iOHIDDeviceSetValueMultipleWithCallback == nil {
		panic("iokit: symbol IOHIDDeviceSetValueMultipleWithCallback not loaded")
	}
	return _iOHIDDeviceSetValueMultipleWithCallback(arg0, arg1, arg2, arg3)
}

var _iOHIDDeviceSetValueWithCallback func(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 float64, arg4 unsafe.Pointer) int

// IOHIDDeviceSetValueWithCallback sets a value for an element and returns status via a completion callback.
//
// See: https://developer.apple.com/documentation/iokit/1588667-iohiddevicesetvaluewithcallback
func IOHIDDeviceSetValueWithCallback(arg0 IOHIDDeviceRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 float64, arg4 unsafe.Pointer) int {
	if _iOHIDDeviceSetValueWithCallback == nil {
		panic("iokit: symbol IOHIDDeviceSetValueWithCallback not loaded")
	}
	return _iOHIDDeviceSetValueWithCallback(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDDeviceUnscheduleFromRunLoop func(arg0 IOHIDDeviceRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDDeviceUnscheduleFromRunLoop unschedules HID device with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1588650-iohiddeviceunschedulefromrunloop
func IOHIDDeviceUnscheduleFromRunLoop(arg0 IOHIDDeviceRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDDeviceUnscheduleFromRunLoop == nil {
		panic("iokit: symbol IOHIDDeviceUnscheduleFromRunLoop not loaded")
	}
	_iOHIDDeviceUnscheduleFromRunLoop(arg0, arg1, arg2)
}

var _iOHIDElementAttach func(arg0 IOHIDElementRef, arg1 IOHIDElementRef)

// IOHIDElementAttach establish a relationship between one or more elements.
//
// See: https://developer.apple.com/documentation/iokit/1564146-iohidelementattach
func IOHIDElementAttach(arg0 IOHIDElementRef, arg1 IOHIDElementRef) {
	if _iOHIDElementAttach == nil {
		panic("iokit: symbol IOHIDElementAttach not loaded")
	}
	_iOHIDElementAttach(arg0, arg1)
}

var _iOHIDElementCopyAttached func(arg0 IOHIDElementRef) corefoundation.CFArrayRef

// IOHIDElementCopyAttached obtain attached elements.
//
// See: https://developer.apple.com/documentation/iokit/1564123-iohidelementcopyattached
func IOHIDElementCopyAttached(arg0 IOHIDElementRef) corefoundation.CFArrayRef {
	if _iOHIDElementCopyAttached == nil {
		panic("iokit: symbol IOHIDElementCopyAttached not loaded")
	}
	return _iOHIDElementCopyAttached(arg0)
}

var _iOHIDElementCreateWithDictionary func(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFDictionaryRef) IOHIDElementRef

// IOHIDElementCreateWithDictionary creates an element from a dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1564115-iohidelementcreatewithdictionary
func IOHIDElementCreateWithDictionary(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFDictionaryRef) IOHIDElementRef {
	if _iOHIDElementCreateWithDictionary == nil {
		panic("iokit: symbol IOHIDElementCreateWithDictionary not loaded")
	}
	return _iOHIDElementCreateWithDictionary(arg0, arg1)
}

var _iOHIDElementDetach func(arg0 IOHIDElementRef, arg1 IOHIDElementRef)

// IOHIDElementDetach remove a relationship between one or more elements.
//
// See: https://developer.apple.com/documentation/iokit/1564116-iohidelementdetach
func IOHIDElementDetach(arg0 IOHIDElementRef, arg1 IOHIDElementRef) {
	if _iOHIDElementDetach == nil {
		panic("iokit: symbol IOHIDElementDetach not loaded")
	}
	_iOHIDElementDetach(arg0, arg1)
}

var _iOHIDElementGetChildren func(arg0 IOHIDElementRef) corefoundation.CFArrayRef

// IOHIDElementGetChildren returns the children for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564119-iohidelementgetchildren
func IOHIDElementGetChildren(arg0 IOHIDElementRef) corefoundation.CFArrayRef {
	if _iOHIDElementGetChildren == nil {
		panic("iokit: symbol IOHIDElementGetChildren not loaded")
	}
	return _iOHIDElementGetChildren(arg0)
}

var _iOHIDElementGetCollectionType func(arg0 IOHIDElementRef) IOHIDElementCollectionType

// IOHIDElementGetCollectionType retrieves the collection type for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564132-iohidelementgetcollectiontype
func IOHIDElementGetCollectionType(arg0 IOHIDElementRef) IOHIDElementCollectionType {
	if _iOHIDElementGetCollectionType == nil {
		panic("iokit: symbol IOHIDElementGetCollectionType not loaded")
	}
	return _iOHIDElementGetCollectionType(arg0)
}

var _iOHIDElementGetCookie func(arg0 IOHIDElementRef) IOHIDElementCookie

// IOHIDElementGetCookie retrieves the cookie for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564124-iohidelementgetcookie
func IOHIDElementGetCookie(arg0 IOHIDElementRef) IOHIDElementCookie {
	if _iOHIDElementGetCookie == nil {
		panic("iokit: symbol IOHIDElementGetCookie not loaded")
	}
	return _iOHIDElementGetCookie(arg0)
}

var _iOHIDElementGetDevice func(arg0 IOHIDElementRef) IOHIDDeviceRef

// IOHIDElementGetDevice obtain the device associated with the element.
//
// See: https://developer.apple.com/documentation/iokit/1564139-iohidelementgetdevice
func IOHIDElementGetDevice(arg0 IOHIDElementRef) IOHIDDeviceRef {
	if _iOHIDElementGetDevice == nil {
		panic("iokit: symbol IOHIDElementGetDevice not loaded")
	}
	return _iOHIDElementGetDevice(arg0)
}

var _iOHIDElementGetLogicalMax func(arg0 IOHIDElementRef) int

// IOHIDElementGetLogicalMax returns the maximum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564143-iohidelementgetlogicalmax
func IOHIDElementGetLogicalMax(arg0 IOHIDElementRef) int {
	if _iOHIDElementGetLogicalMax == nil {
		panic("iokit: symbol IOHIDElementGetLogicalMax not loaded")
	}
	return _iOHIDElementGetLogicalMax(arg0)
}

var _iOHIDElementGetLogicalMin func(arg0 IOHIDElementRef) int

// IOHIDElementGetLogicalMin returns the minimum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564137-iohidelementgetlogicalmin
func IOHIDElementGetLogicalMin(arg0 IOHIDElementRef) int {
	if _iOHIDElementGetLogicalMin == nil {
		panic("iokit: symbol IOHIDElementGetLogicalMin not loaded")
	}
	return _iOHIDElementGetLogicalMin(arg0)
}

var _iOHIDElementGetName func(arg0 IOHIDElementRef) corefoundation.CFStringRef

// IOHIDElementGetName returns the name for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564117-iohidelementgetname
func IOHIDElementGetName(arg0 IOHIDElementRef) corefoundation.CFStringRef {
	if _iOHIDElementGetName == nil {
		panic("iokit: symbol IOHIDElementGetName not loaded")
	}
	return _iOHIDElementGetName(arg0)
}

var _iOHIDElementGetParent func(arg0 IOHIDElementRef) IOHIDElementRef

// IOHIDElementGetParent returns the parent for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564144-iohidelementgetparent
func IOHIDElementGetParent(arg0 IOHIDElementRef) IOHIDElementRef {
	if _iOHIDElementGetParent == nil {
		panic("iokit: symbol IOHIDElementGetParent not loaded")
	}
	return _iOHIDElementGetParent(arg0)
}

var _iOHIDElementGetPhysicalMax func(arg0 IOHIDElementRef) int

// IOHIDElementGetPhysicalMax returns the scaled maximum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564134-iohidelementgetphysicalmax
func IOHIDElementGetPhysicalMax(arg0 IOHIDElementRef) int {
	if _iOHIDElementGetPhysicalMax == nil {
		panic("iokit: symbol IOHIDElementGetPhysicalMax not loaded")
	}
	return _iOHIDElementGetPhysicalMax(arg0)
}

var _iOHIDElementGetPhysicalMin func(arg0 IOHIDElementRef) int

// IOHIDElementGetPhysicalMin returns the scaled minimum value possible for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564140-iohidelementgetphysicalmin
func IOHIDElementGetPhysicalMin(arg0 IOHIDElementRef) int {
	if _iOHIDElementGetPhysicalMin == nil {
		panic("iokit: symbol IOHIDElementGetPhysicalMin not loaded")
	}
	return _iOHIDElementGetPhysicalMin(arg0)
}

var _iOHIDElementGetProperty func(arg0 IOHIDElementRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOHIDElementGetProperty returns the an element property.
//
// See: https://developer.apple.com/documentation/iokit/1564118-iohidelementgetproperty
func IOHIDElementGetProperty(arg0 IOHIDElementRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOHIDElementGetProperty == nil {
		panic("iokit: symbol IOHIDElementGetProperty not loaded")
	}
	return _iOHIDElementGetProperty(arg0, arg1)
}

var _iOHIDElementGetReportCount func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetReportCount returns the report count for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564142-iohidelementgetreportcount
func IOHIDElementGetReportCount(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetReportCount == nil {
		panic("iokit: symbol IOHIDElementGetReportCount not loaded")
	}
	return _iOHIDElementGetReportCount(arg0)
}

var _iOHIDElementGetReportID func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetReportID returns the report ID for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564122-iohidelementgetreportid
func IOHIDElementGetReportID(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetReportID == nil {
		panic("iokit: symbol IOHIDElementGetReportID not loaded")
	}
	return _iOHIDElementGetReportID(arg0)
}

var _iOHIDElementGetReportSize func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetReportSize returns the report size in bits for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564130-iohidelementgetreportsize
func IOHIDElementGetReportSize(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetReportSize == nil {
		panic("iokit: symbol IOHIDElementGetReportSize not loaded")
	}
	return _iOHIDElementGetReportSize(arg0)
}

var _iOHIDElementGetType func(arg0 IOHIDElementRef) IOHIDElementType

// IOHIDElementGetType retrieves the type for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564135-iohidelementgettype
func IOHIDElementGetType(arg0 IOHIDElementRef) IOHIDElementType {
	if _iOHIDElementGetType == nil {
		panic("iokit: symbol IOHIDElementGetType not loaded")
	}
	return _iOHIDElementGetType(arg0)
}

var _iOHIDElementGetTypeID func() uint

// IOHIDElementGetTypeID returns the type identifier of all IOHIDElement instances.
//
// See: https://developer.apple.com/documentation/iokit/1564120-iohidelementgettypeid
func IOHIDElementGetTypeID() uint {
	if _iOHIDElementGetTypeID == nil {
		panic("iokit: symbol IOHIDElementGetTypeID not loaded")
	}
	return _iOHIDElementGetTypeID()
}

var _iOHIDElementGetUnit func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetUnit returns the unit property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564136-iohidelementgetunit
func IOHIDElementGetUnit(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetUnit == nil {
		panic("iokit: symbol IOHIDElementGetUnit not loaded")
	}
	return _iOHIDElementGetUnit(arg0)
}

var _iOHIDElementGetUnitExponent func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetUnitExponent returns the unit exponenet in base 10 for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564121-iohidelementgetunitexponent
func IOHIDElementGetUnitExponent(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetUnitExponent == nil {
		panic("iokit: symbol IOHIDElementGetUnitExponent not loaded")
	}
	return _iOHIDElementGetUnitExponent(arg0)
}

var _iOHIDElementGetUsage func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetUsage retrieves the usage for an element.
//
// See: https://developer.apple.com/documentation/iokit/1564126-iohidelementgetusage
func IOHIDElementGetUsage(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetUsage == nil {
		panic("iokit: symbol IOHIDElementGetUsage not loaded")
	}
	return _iOHIDElementGetUsage(arg0)
}

var _iOHIDElementGetUsagePage func(arg0 IOHIDElementRef) uint32

// IOHIDElementGetUsagePage retrieves the usage page for an element.
//
// See: https://developer.apple.com/documentation/iokit/1564128-iohidelementgetusagepage
func IOHIDElementGetUsagePage(arg0 IOHIDElementRef) uint32 {
	if _iOHIDElementGetUsagePage == nil {
		panic("iokit: symbol IOHIDElementGetUsagePage not loaded")
	}
	return _iOHIDElementGetUsagePage(arg0)
}

var _iOHIDElementHasNullState func(arg0 IOHIDElementRef) bool

// IOHIDElementHasNullState returns the null state property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564145-iohidelementhasnullstate
func IOHIDElementHasNullState(arg0 IOHIDElementRef) bool {
	if _iOHIDElementHasNullState == nil {
		panic("iokit: symbol IOHIDElementHasNullState not loaded")
	}
	return _iOHIDElementHasNullState(arg0)
}

var _iOHIDElementHasPreferredState func(arg0 IOHIDElementRef) bool

// IOHIDElementHasPreferredState returns the preferred state property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564141-iohidelementhaspreferredstate
func IOHIDElementHasPreferredState(arg0 IOHIDElementRef) bool {
	if _iOHIDElementHasPreferredState == nil {
		panic("iokit: symbol IOHIDElementHasPreferredState not loaded")
	}
	return _iOHIDElementHasPreferredState(arg0)
}

var _iOHIDElementIsArray func(arg0 IOHIDElementRef) bool

// IOHIDElementIsArray returns the array property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564133-iohidelementisarray
func IOHIDElementIsArray(arg0 IOHIDElementRef) bool {
	if _iOHIDElementIsArray == nil {
		panic("iokit: symbol IOHIDElementIsArray not loaded")
	}
	return _iOHIDElementIsArray(arg0)
}

var _iOHIDElementIsNonLinear func(arg0 IOHIDElementRef) bool

// IOHIDElementIsNonLinear returns the linear property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564131-iohidelementisnonlinear
func IOHIDElementIsNonLinear(arg0 IOHIDElementRef) bool {
	if _iOHIDElementIsNonLinear == nil {
		panic("iokit: symbol IOHIDElementIsNonLinear not loaded")
	}
	return _iOHIDElementIsNonLinear(arg0)
}

var _iOHIDElementIsRelative func(arg0 IOHIDElementRef) bool

// IOHIDElementIsRelative returns the relative property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564129-iohidelementisrelative
func IOHIDElementIsRelative(arg0 IOHIDElementRef) bool {
	if _iOHIDElementIsRelative == nil {
		panic("iokit: symbol IOHIDElementIsRelative not loaded")
	}
	return _iOHIDElementIsRelative(arg0)
}

var _iOHIDElementIsVirtual func(arg0 IOHIDElementRef) bool

// IOHIDElementIsVirtual returns the virtual property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564125-iohidelementisvirtual
func IOHIDElementIsVirtual(arg0 IOHIDElementRef) bool {
	if _iOHIDElementIsVirtual == nil {
		panic("iokit: symbol IOHIDElementIsVirtual not loaded")
	}
	return _iOHIDElementIsVirtual(arg0)
}

var _iOHIDElementIsWrapping func(arg0 IOHIDElementRef) bool

// IOHIDElementIsWrapping returns the wrap property for the element.
//
// See: https://developer.apple.com/documentation/iokit/1564127-iohidelementiswrapping
func IOHIDElementIsWrapping(arg0 IOHIDElementRef) bool {
	if _iOHIDElementIsWrapping == nil {
		panic("iokit: symbol IOHIDElementIsWrapping not loaded")
	}
	return _iOHIDElementIsWrapping(arg0)
}

var _iOHIDElementSetProperty func(arg0 IOHIDElementRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool

// IOHIDElementSetProperty sets an element property.
//
// See: https://developer.apple.com/documentation/iokit/1564138-iohidelementsetproperty
func IOHIDElementSetProperty(arg0 IOHIDElementRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool {
	if _iOHIDElementSetProperty == nil {
		panic("iokit: symbol IOHIDElementSetProperty not loaded")
	}
	return _iOHIDElementSetProperty(arg0, arg1, arg2)
}

var _iOHIDEventSystemClientCopyProperty func(arg0 IOHIDEventSystemClientRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOHIDEventSystemClientCopyProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269513-iohideventsystemclientcopyproper
func IOHIDEventSystemClientCopyProperty(arg0 IOHIDEventSystemClientRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOHIDEventSystemClientCopyProperty == nil {
		panic("iokit: symbol IOHIDEventSystemClientCopyProperty not loaded")
	}
	return _iOHIDEventSystemClientCopyProperty(arg0, arg1)
}

var _iOHIDEventSystemClientCopyServices func(arg0 IOHIDEventSystemClientRef) corefoundation.CFArrayRef

// IOHIDEventSystemClientCopyServices.
//
// See: https://developer.apple.com/documentation/iokit/2269511-iohideventsystemclientcopyservic
func IOHIDEventSystemClientCopyServices(arg0 IOHIDEventSystemClientRef) corefoundation.CFArrayRef {
	if _iOHIDEventSystemClientCopyServices == nil {
		panic("iokit: symbol IOHIDEventSystemClientCopyServices not loaded")
	}
	return _iOHIDEventSystemClientCopyServices(arg0)
}

var _iOHIDEventSystemClientCreateSimpleClient func(arg0 corefoundation.CFAllocatorRef) IOHIDEventSystemClientRef

// IOHIDEventSystemClientCreateSimpleClient.
//
// See: https://developer.apple.com/documentation/iokit/2269514-iohideventsystemclientcreatesimp
func IOHIDEventSystemClientCreateSimpleClient(arg0 corefoundation.CFAllocatorRef) IOHIDEventSystemClientRef {
	if _iOHIDEventSystemClientCreateSimpleClient == nil {
		panic("iokit: symbol IOHIDEventSystemClientCreateSimpleClient not loaded")
	}
	return _iOHIDEventSystemClientCreateSimpleClient(arg0)
}

var _iOHIDEventSystemClientGetTypeID func() uint

// IOHIDEventSystemClientGetTypeID.
//
// See: https://developer.apple.com/documentation/iokit/2269512-iohideventsystemclientgettypeid
func IOHIDEventSystemClientGetTypeID() uint {
	if _iOHIDEventSystemClientGetTypeID == nil {
		panic("iokit: symbol IOHIDEventSystemClientGetTypeID not loaded")
	}
	return _iOHIDEventSystemClientGetTypeID()
}

var _iOHIDEventSystemClientSetProperty func(arg0 IOHIDEventSystemClientRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool

// IOHIDEventSystemClientSetProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269517-iohideventsystemclientsetpropert
func IOHIDEventSystemClientSetProperty(arg0 IOHIDEventSystemClientRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool {
	if _iOHIDEventSystemClientSetProperty == nil {
		panic("iokit: symbol IOHIDEventSystemClientSetProperty not loaded")
	}
	return _iOHIDEventSystemClientSetProperty(arg0, arg1, arg2)
}

var _iOHIDGetAccelerationWithKey func(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 float64) int32

// IOHIDGetAccelerationWithKey.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555418-iohidgetaccelerationwithkey
func IOHIDGetAccelerationWithKey(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 float64) int32 {
	if _iOHIDGetAccelerationWithKey == nil {
		panic("iokit: symbol IOHIDGetAccelerationWithKey not loaded")
	}
	return _iOHIDGetAccelerationWithKey(arg0, arg1, arg2)
}

var _iOHIDGetActivityState func(arg0 uintptr, arg1 bool) int32

// IOHIDGetActivityState.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555412-iohidgetactivitystate
func IOHIDGetActivityState(arg0 uintptr, arg1 bool) int32 {
	if _iOHIDGetActivityState == nil {
		panic("iokit: symbol IOHIDGetActivityState not loaded")
	}
	return _iOHIDGetActivityState(arg0, arg1)
}

var _iOHIDGetButtonEventNum func(arg0 uintptr, arg1 NXMouseButton, arg2 int) int32

// IOHIDGetButtonEventNum.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555407-iohidgetbuttoneventnum
func IOHIDGetButtonEventNum(arg0 uintptr, arg1 NXMouseButton, arg2 int) int32 {
	if _iOHIDGetButtonEventNum == nil {
		panic("iokit: symbol IOHIDGetButtonEventNum not loaded")
	}
	return _iOHIDGetButtonEventNum(arg0, arg1, arg2)
}

var _iOHIDGetModifierLockState func(arg0 uintptr, arg1 int, arg2 bool) int32

// IOHIDGetModifierLockState.
//
// See: https://developer.apple.com/documentation/iokit/1555400-iohidgetmodifierlockstate
func IOHIDGetModifierLockState(arg0 uintptr, arg1 int, arg2 bool) int32 {
	if _iOHIDGetModifierLockState == nil {
		panic("iokit: symbol IOHIDGetModifierLockState not loaded")
	}
	return _iOHIDGetModifierLockState(arg0, arg1, arg2)
}

var _iOHIDGetMouseAcceleration func(arg0 uintptr, arg1 float64) int32

// IOHIDGetMouseAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555417-iohidgetmouseacceleration
func IOHIDGetMouseAcceleration(arg0 uintptr, arg1 float64) int32 {
	if _iOHIDGetMouseAcceleration == nil {
		panic("iokit: symbol IOHIDGetMouseAcceleration not loaded")
	}
	return _iOHIDGetMouseAcceleration(arg0, arg1)
}

var _iOHIDGetMouseButtonMode func(arg0 uintptr, arg1 int) int32

// IOHIDGetMouseButtonMode.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555413-iohidgetmousebuttonmode
func IOHIDGetMouseButtonMode(arg0 uintptr, arg1 int) int32 {
	if _iOHIDGetMouseButtonMode == nil {
		panic("iokit: symbol IOHIDGetMouseButtonMode not loaded")
	}
	return _iOHIDGetMouseButtonMode(arg0, arg1)
}

var _iOHIDGetParameter func(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 uint, arg3 uint) int32

// IOHIDGetParameter.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555405-iohidgetparameter
func IOHIDGetParameter(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 uint, arg3 uint) int32 {
	if _iOHIDGetParameter == nil {
		panic("iokit: symbol IOHIDGetParameter not loaded")
	}
	return _iOHIDGetParameter(arg0, arg1, arg2, arg3)
}

var _iOHIDGetScrollAcceleration func(arg0 uintptr, arg1 float64) int32

// IOHIDGetScrollAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555389-iohidgetscrollacceleration
func IOHIDGetScrollAcceleration(arg0 uintptr, arg1 float64) int32 {
	if _iOHIDGetScrollAcceleration == nil {
		panic("iokit: symbol IOHIDGetScrollAcceleration not loaded")
	}
	return _iOHIDGetScrollAcceleration(arg0, arg1)
}

var _iOHIDGetStateForSelector func(arg0 uintptr, arg1 int, arg2 uint32) int32

// IOHIDGetStateForSelector.
//
// See: https://developer.apple.com/documentation/iokit/1555411-iohidgetstateforselector
func IOHIDGetStateForSelector(arg0 uintptr, arg1 int, arg2 uint32) int32 {
	if _iOHIDGetStateForSelector == nil {
		panic("iokit: symbol IOHIDGetStateForSelector not loaded")
	}
	return _iOHIDGetStateForSelector(arg0, arg1, arg2)
}

var _iOHIDManagerActivate func(arg0 IOHIDManagerRef)

// IOHIDManagerActivate.
//
// See: https://developer.apple.com/documentation/iokit/3042786-iohidmanageractivate
func IOHIDManagerActivate(arg0 IOHIDManagerRef) {
	if _iOHIDManagerActivate == nil {
		panic("iokit: symbol IOHIDManagerActivate not loaded")
	}
	_iOHIDManagerActivate(arg0)
}

var _iOHIDManagerCancel func(arg0 IOHIDManagerRef)

// IOHIDManagerCancel.
//
// See: https://developer.apple.com/documentation/iokit/3042787-iohidmanagercancel
func IOHIDManagerCancel(arg0 IOHIDManagerRef) {
	if _iOHIDManagerCancel == nil {
		panic("iokit: symbol IOHIDManagerCancel not loaded")
	}
	_iOHIDManagerCancel(arg0)
}

var _iOHIDManagerClose func(arg0 IOHIDManagerRef, arg1 uint32) int

// IOHIDManagerClose closes the IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438405-iohidmanagerclose
func IOHIDManagerClose(arg0 IOHIDManagerRef, arg1 uint32) int {
	if _iOHIDManagerClose == nil {
		panic("iokit: symbol IOHIDManagerClose not loaded")
	}
	return _iOHIDManagerClose(arg0, arg1)
}

var _iOHIDManagerCopyDevices func(arg0 IOHIDManagerRef) corefoundation.CFSet

// IOHIDManagerCopyDevices obtains currently enumerated devices.
//
// See: https://developer.apple.com/documentation/iokit/1438391-iohidmanagercopydevices
func IOHIDManagerCopyDevices(arg0 IOHIDManagerRef) corefoundation.CFSet {
	if _iOHIDManagerCopyDevices == nil {
		panic("iokit: symbol IOHIDManagerCopyDevices not loaded")
	}
	return _iOHIDManagerCopyDevices(arg0)
}

var _iOHIDManagerCreate func(arg0 corefoundation.CFAllocatorRef, arg1 uint32) IOHIDManagerRef

// IOHIDManagerCreate creates an IOHIDManager object.
//
// See: https://developer.apple.com/documentation/iokit/1438383-iohidmanagercreate
func IOHIDManagerCreate(arg0 corefoundation.CFAllocatorRef, arg1 uint32) IOHIDManagerRef {
	if _iOHIDManagerCreate == nil {
		panic("iokit: symbol IOHIDManagerCreate not loaded")
	}
	return _iOHIDManagerCreate(arg0, arg1)
}

var _iOHIDManagerGetProperty func(arg0 IOHIDManagerRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOHIDManagerGetProperty obtains a property of an IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438403-iohidmanagergetproperty
func IOHIDManagerGetProperty(arg0 IOHIDManagerRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOHIDManagerGetProperty == nil {
		panic("iokit: symbol IOHIDManagerGetProperty not loaded")
	}
	return _iOHIDManagerGetProperty(arg0, arg1)
}

var _iOHIDManagerGetTypeID func() uint

// IOHIDManagerGetTypeID returns the type identifier of all IOHIDManager instances.
//
// See: https://developer.apple.com/documentation/iokit/1438375-iohidmanagergettypeid
func IOHIDManagerGetTypeID() uint {
	if _iOHIDManagerGetTypeID == nil {
		panic("iokit: symbol IOHIDManagerGetTypeID not loaded")
	}
	return _iOHIDManagerGetTypeID()
}

var _iOHIDManagerOpen func(arg0 IOHIDManagerRef, arg1 uint32) int

// IOHIDManagerOpen opens the IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438369-iohidmanageropen
func IOHIDManagerOpen(arg0 IOHIDManagerRef, arg1 uint32) int {
	if _iOHIDManagerOpen == nil {
		panic("iokit: symbol IOHIDManagerOpen not loaded")
	}
	return _iOHIDManagerOpen(arg0, arg1)
}

var _iOHIDManagerRegisterDeviceMatchingCallback func(arg0 IOHIDManagerRef, arg1 unsafe.Pointer)

// IOHIDManagerRegisterDeviceMatchingCallback registers a callback to be used a device is enumerated.
//
// See: https://developer.apple.com/documentation/iokit/1438399-iohidmanagerregisterdevicematchi
func IOHIDManagerRegisterDeviceMatchingCallback(arg0 IOHIDManagerRef, arg1 unsafe.Pointer) {
	if _iOHIDManagerRegisterDeviceMatchingCallback == nil {
		panic("iokit: symbol IOHIDManagerRegisterDeviceMatchingCallback not loaded")
	}
	_iOHIDManagerRegisterDeviceMatchingCallback(arg0, arg1)
}

var _iOHIDManagerRegisterDeviceRemovalCallback func(arg0 IOHIDManagerRef, arg1 unsafe.Pointer)

// IOHIDManagerRegisterDeviceRemovalCallback registers a callback to be used when any enumerated device is removed.
//
// See: https://developer.apple.com/documentation/iokit/1438376-iohidmanagerregisterdeviceremova
func IOHIDManagerRegisterDeviceRemovalCallback(arg0 IOHIDManagerRef, arg1 unsafe.Pointer) {
	if _iOHIDManagerRegisterDeviceRemovalCallback == nil {
		panic("iokit: symbol IOHIDManagerRegisterDeviceRemovalCallback not loaded")
	}
	_iOHIDManagerRegisterDeviceRemovalCallback(arg0, arg1)
}

var _iOHIDManagerRegisterInputReportCallback func(arg0 IOHIDManagerRef, arg1 unsafe.Pointer)

// IOHIDManagerRegisterInputReportCallback registers a callback to be used when an input report is issued by any enumerated device.
//
// See: https://developer.apple.com/documentation/iokit/1438397-iohidmanagerregisterinputreportc
func IOHIDManagerRegisterInputReportCallback(arg0 IOHIDManagerRef, arg1 unsafe.Pointer) {
	if _iOHIDManagerRegisterInputReportCallback == nil {
		panic("iokit: symbol IOHIDManagerRegisterInputReportCallback not loaded")
	}
	_iOHIDManagerRegisterInputReportCallback(arg0, arg1)
}

var _iOHIDManagerRegisterInputReportWithTimeStampCallback func(arg0 IOHIDManagerRef, arg1 unsafe.Pointer)

// IOHIDManagerRegisterInputReportWithTimeStampCallback.
//
// See: https://developer.apple.com/documentation/iokit/3042788-iohidmanagerregisterinputreportw
func IOHIDManagerRegisterInputReportWithTimeStampCallback(arg0 IOHIDManagerRef, arg1 unsafe.Pointer) {
	if _iOHIDManagerRegisterInputReportWithTimeStampCallback == nil {
		panic("iokit: symbol IOHIDManagerRegisterInputReportWithTimeStampCallback not loaded")
	}
	_iOHIDManagerRegisterInputReportWithTimeStampCallback(arg0, arg1)
}

var _iOHIDManagerRegisterInputValueCallback func(arg0 IOHIDManagerRef, arg1 unsafe.Pointer)

// IOHIDManagerRegisterInputValueCallback registers a callback to be used when an input value is issued by any enumerated device.
//
// See: https://developer.apple.com/documentation/iokit/1438367-iohidmanagerregisterinputvalueca
func IOHIDManagerRegisterInputValueCallback(arg0 IOHIDManagerRef, arg1 unsafe.Pointer) {
	if _iOHIDManagerRegisterInputValueCallback == nil {
		panic("iokit: symbol IOHIDManagerRegisterInputValueCallback not loaded")
	}
	_iOHIDManagerRegisterInputValueCallback(arg0, arg1)
}

var _iOHIDManagerSaveToPropertyDomain func(arg0 IOHIDManagerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 uint32)

// IOHIDManagerSaveToPropertyDomain used to write out the current properties to a specific domain.
//
// See: https://developer.apple.com/documentation/iokit/1438395-iohidmanagersavetopropertydomain
func IOHIDManagerSaveToPropertyDomain(arg0 IOHIDManagerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 uint32) {
	if _iOHIDManagerSaveToPropertyDomain == nil {
		panic("iokit: symbol IOHIDManagerSaveToPropertyDomain not loaded")
	}
	_iOHIDManagerSaveToPropertyDomain(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDManagerScheduleWithRunLoop func(arg0 IOHIDManagerRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDManagerScheduleWithRunLoop schedules HID manager with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1438409-iohidmanagerschedulewithrunloop
func IOHIDManagerScheduleWithRunLoop(arg0 IOHIDManagerRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDManagerScheduleWithRunLoop == nil {
		panic("iokit: symbol IOHIDManagerScheduleWithRunLoop not loaded")
	}
	_iOHIDManagerScheduleWithRunLoop(arg0, arg1, arg2)
}

var _iOHIDManagerSetCancelHandler func(arg0 IOHIDManagerRef, arg1 unsafe.Pointer)

// IOHIDManagerSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3042789-iohidmanagersetcancelhandler
func IOHIDManagerSetCancelHandler(arg0 IOHIDManagerRef, arg1 unsafe.Pointer) {
	if _iOHIDManagerSetCancelHandler == nil {
		panic("iokit: symbol IOHIDManagerSetCancelHandler not loaded")
	}
	_iOHIDManagerSetCancelHandler(arg0, arg1)
}

var _iOHIDManagerSetDeviceMatching func(arg0 IOHIDManagerRef, arg1 corefoundation.CFDictionaryRef)

// IOHIDManagerSetDeviceMatching sets matching criteria for device enumeration.
//
// See: https://developer.apple.com/documentation/iokit/1438371-iohidmanagersetdevicematching
func IOHIDManagerSetDeviceMatching(arg0 IOHIDManagerRef, arg1 corefoundation.CFDictionaryRef) {
	if _iOHIDManagerSetDeviceMatching == nil {
		panic("iokit: symbol IOHIDManagerSetDeviceMatching not loaded")
	}
	_iOHIDManagerSetDeviceMatching(arg0, arg1)
}

var _iOHIDManagerSetDeviceMatchingMultiple func(arg0 IOHIDManagerRef, arg1 corefoundation.CFArrayRef)

// IOHIDManagerSetDeviceMatchingMultiple sets multiple matching criteria for device enumeration.
//
// See: https://developer.apple.com/documentation/iokit/1438387-iohidmanagersetdevicematchingmul
func IOHIDManagerSetDeviceMatchingMultiple(arg0 IOHIDManagerRef, arg1 corefoundation.CFArrayRef) {
	if _iOHIDManagerSetDeviceMatchingMultiple == nil {
		panic("iokit: symbol IOHIDManagerSetDeviceMatchingMultiple not loaded")
	}
	_iOHIDManagerSetDeviceMatchingMultiple(arg0, arg1)
}

var _iOHIDManagerSetDispatchQueue func(arg0 IOHIDManagerRef, arg1 uintptr)

// IOHIDManagerSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3042790-iohidmanagersetdispatchqueue
func IOHIDManagerSetDispatchQueue(arg0 IOHIDManagerRef, arg1 dispatch.Queue) {
	if _iOHIDManagerSetDispatchQueue == nil {
		panic("iokit: symbol IOHIDManagerSetDispatchQueue not loaded")
	}
	_iOHIDManagerSetDispatchQueue(arg0, uintptr(arg1.Handle()))
}

var _iOHIDManagerSetInputValueMatching func(arg0 IOHIDManagerRef, arg1 corefoundation.CFDictionaryRef)

// IOHIDManagerSetInputValueMatching sets matching criteria for input values received via IOHIDManagerRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1438389-iohidmanagersetinputvaluematchin
func IOHIDManagerSetInputValueMatching(arg0 IOHIDManagerRef, arg1 corefoundation.CFDictionaryRef) {
	if _iOHIDManagerSetInputValueMatching == nil {
		panic("iokit: symbol IOHIDManagerSetInputValueMatching not loaded")
	}
	_iOHIDManagerSetInputValueMatching(arg0, arg1)
}

var _iOHIDManagerSetInputValueMatchingMultiple func(arg0 IOHIDManagerRef, arg1 corefoundation.CFArrayRef)

// IOHIDManagerSetInputValueMatchingMultiple sets multiple matching criteria for input values received via IOHIDManagerRegisterInputValueCallback.
//
// See: https://developer.apple.com/documentation/iokit/1438379-iohidmanagersetinputvaluematchin
func IOHIDManagerSetInputValueMatchingMultiple(arg0 IOHIDManagerRef, arg1 corefoundation.CFArrayRef) {
	if _iOHIDManagerSetInputValueMatchingMultiple == nil {
		panic("iokit: symbol IOHIDManagerSetInputValueMatchingMultiple not loaded")
	}
	_iOHIDManagerSetInputValueMatchingMultiple(arg0, arg1)
}

var _iOHIDManagerSetProperty func(arg0 IOHIDManagerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool

// IOHIDManagerSetProperty sets a property for an IOHIDManager.
//
// See: https://developer.apple.com/documentation/iokit/1438401-iohidmanagersetproperty
func IOHIDManagerSetProperty(arg0 IOHIDManagerRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool {
	if _iOHIDManagerSetProperty == nil {
		panic("iokit: symbol IOHIDManagerSetProperty not loaded")
	}
	return _iOHIDManagerSetProperty(arg0, arg1, arg2)
}

var _iOHIDManagerUnscheduleFromRunLoop func(arg0 IOHIDManagerRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDManagerUnscheduleFromRunLoop unschedules HID manager with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1438378-iohidmanagerunschedulefromrunloo
func IOHIDManagerUnscheduleFromRunLoop(arg0 IOHIDManagerRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDManagerUnscheduleFromRunLoop == nil {
		panic("iokit: symbol IOHIDManagerUnscheduleFromRunLoop not loaded")
	}
	_iOHIDManagerUnscheduleFromRunLoop(arg0, arg1, arg2)
}

var _iOHIDPostEvent func(arg0 uintptr, arg1 uint32, arg2 IOGPoint, arg3 kernel.NXEventData, arg4 uint32, arg5 uint32, arg6 uint32) int32

// IOHIDPostEvent.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555406-iohidpostevent
func IOHIDPostEvent(arg0 uintptr, arg1 uint32, arg2 IOGPoint, arg3 kernel.NXEventData, arg4 uint32, arg5 uint32, arg6 uint32) int32 {
	if _iOHIDPostEvent == nil {
		panic("iokit: symbol IOHIDPostEvent not loaded")
	}
	return _iOHIDPostEvent(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

var _iOHIDQueueActivate func(arg0 IOHIDQueueRef)

// IOHIDQueueActivate.
//
// See: https://developer.apple.com/documentation/iokit/3042791-iohidqueueactivate
func IOHIDQueueActivate(arg0 IOHIDQueueRef) {
	if _iOHIDQueueActivate == nil {
		panic("iokit: symbol IOHIDQueueActivate not loaded")
	}
	_iOHIDQueueActivate(arg0)
}

var _iOHIDQueueAddElement func(arg0 IOHIDQueueRef, arg1 IOHIDElementRef)

// IOHIDQueueAddElement adds an element to the queue
//
// See: https://developer.apple.com/documentation/iokit/1545835-iohidqueueaddelement
func IOHIDQueueAddElement(arg0 IOHIDQueueRef, arg1 IOHIDElementRef) {
	if _iOHIDQueueAddElement == nil {
		panic("iokit: symbol IOHIDQueueAddElement not loaded")
	}
	_iOHIDQueueAddElement(arg0, arg1)
}

var _iOHIDQueueCancel func(arg0 IOHIDQueueRef)

// IOHIDQueueCancel.
//
// See: https://developer.apple.com/documentation/iokit/3042792-iohidqueuecancel
func IOHIDQueueCancel(arg0 IOHIDQueueRef) {
	if _iOHIDQueueCancel == nil {
		panic("iokit: symbol IOHIDQueueCancel not loaded")
	}
	_iOHIDQueueCancel(arg0)
}

var _iOHIDQueueContainsElement func(arg0 IOHIDQueueRef, arg1 IOHIDElementRef) bool

// IOHIDQueueContainsElement queries the queue to determine if elemement has been added.
//
// See: https://developer.apple.com/documentation/iokit/1545842-iohidqueuecontainselement
func IOHIDQueueContainsElement(arg0 IOHIDQueueRef, arg1 IOHIDElementRef) bool {
	if _iOHIDQueueContainsElement == nil {
		panic("iokit: symbol IOHIDQueueContainsElement not loaded")
	}
	return _iOHIDQueueContainsElement(arg0, arg1)
}

var _iOHIDQueueCopyNextValue func(arg0 IOHIDQueueRef) IOHIDValueRef

// IOHIDQueueCopyNextValue dequeues a retained copy of an element value from the head of an IOHIDQueue.
//
// See: https://developer.apple.com/documentation/iokit/1545844-iohidqueuecopynextvalue
func IOHIDQueueCopyNextValue(arg0 IOHIDQueueRef) IOHIDValueRef {
	if _iOHIDQueueCopyNextValue == nil {
		panic("iokit: symbol IOHIDQueueCopyNextValue not loaded")
	}
	return _iOHIDQueueCopyNextValue(arg0)
}

var _iOHIDQueueCopyNextValueWithTimeout func(arg0 IOHIDQueueRef, arg1 float64) IOHIDValueRef

// IOHIDQueueCopyNextValueWithTimeout dequeues a retained copy of an element value from the head of an IOHIDQueue.
//
// See: https://developer.apple.com/documentation/iokit/1545832-iohidqueuecopynextvaluewithtimeo
func IOHIDQueueCopyNextValueWithTimeout(arg0 IOHIDQueueRef, arg1 float64) IOHIDValueRef {
	if _iOHIDQueueCopyNextValueWithTimeout == nil {
		panic("iokit: symbol IOHIDQueueCopyNextValueWithTimeout not loaded")
	}
	return _iOHIDQueueCopyNextValueWithTimeout(arg0, arg1)
}

var _iOHIDQueueCreate func(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDDeviceRef, arg2 int, arg3 uint32) IOHIDQueueRef

// IOHIDQueueCreate creates an IOHIDQueue object for the specified device.
//
// See: https://developer.apple.com/documentation/iokit/1545840-iohidqueuecreate
func IOHIDQueueCreate(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDDeviceRef, arg2 int, arg3 uint32) IOHIDQueueRef {
	if _iOHIDQueueCreate == nil {
		panic("iokit: symbol IOHIDQueueCreate not loaded")
	}
	return _iOHIDQueueCreate(arg0, arg1, arg2, arg3)
}

var _iOHIDQueueGetDepth func(arg0 IOHIDQueueRef) int

// IOHIDQueueGetDepth obtain the depth of the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545833-iohidqueuegetdepth
func IOHIDQueueGetDepth(arg0 IOHIDQueueRef) int {
	if _iOHIDQueueGetDepth == nil {
		panic("iokit: symbol IOHIDQueueGetDepth not loaded")
	}
	return _iOHIDQueueGetDepth(arg0)
}

var _iOHIDQueueGetDevice func(arg0 IOHIDQueueRef) IOHIDDeviceRef

// IOHIDQueueGetDevice obtain the device associated with the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545839-iohidqueuegetdevice
func IOHIDQueueGetDevice(arg0 IOHIDQueueRef) IOHIDDeviceRef {
	if _iOHIDQueueGetDevice == nil {
		panic("iokit: symbol IOHIDQueueGetDevice not loaded")
	}
	return _iOHIDQueueGetDevice(arg0)
}

var _iOHIDQueueGetTypeID func() uint

// IOHIDQueueGetTypeID returns the type identifier of all IOHIDQueue instances.
//
// See: https://developer.apple.com/documentation/iokit/1545836-iohidqueuegettypeid
func IOHIDQueueGetTypeID() uint {
	if _iOHIDQueueGetTypeID == nil {
		panic("iokit: symbol IOHIDQueueGetTypeID not loaded")
	}
	return _iOHIDQueueGetTypeID()
}

var _iOHIDQueueRegisterValueAvailableCallback func(arg0 IOHIDQueueRef, arg1 unsafe.Pointer)

// IOHIDQueueRegisterValueAvailableCallback sets callback to be used when the queue transitions to non-empty.
//
// See: https://developer.apple.com/documentation/iokit/1545829-iohidqueueregistervalueavailable
func IOHIDQueueRegisterValueAvailableCallback(arg0 IOHIDQueueRef, arg1 unsafe.Pointer) {
	if _iOHIDQueueRegisterValueAvailableCallback == nil {
		panic("iokit: symbol IOHIDQueueRegisterValueAvailableCallback not loaded")
	}
	_iOHIDQueueRegisterValueAvailableCallback(arg0, arg1)
}

var _iOHIDQueueRemoveElement func(arg0 IOHIDQueueRef, arg1 IOHIDElementRef)

// IOHIDQueueRemoveElement removes an element from the queue
//
// See: https://developer.apple.com/documentation/iokit/1545838-iohidqueueremoveelement
func IOHIDQueueRemoveElement(arg0 IOHIDQueueRef, arg1 IOHIDElementRef) {
	if _iOHIDQueueRemoveElement == nil {
		panic("iokit: symbol IOHIDQueueRemoveElement not loaded")
	}
	_iOHIDQueueRemoveElement(arg0, arg1)
}

var _iOHIDQueueScheduleWithRunLoop func(arg0 IOHIDQueueRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDQueueScheduleWithRunLoop schedules queue with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1545841-iohidqueueschedulewithrunloop
func IOHIDQueueScheduleWithRunLoop(arg0 IOHIDQueueRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDQueueScheduleWithRunLoop == nil {
		panic("iokit: symbol IOHIDQueueScheduleWithRunLoop not loaded")
	}
	_iOHIDQueueScheduleWithRunLoop(arg0, arg1, arg2)
}

var _iOHIDQueueSetCancelHandler func(arg0 IOHIDQueueRef, arg1 unsafe.Pointer)

// IOHIDQueueSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3042793-iohidqueuesetcancelhandler
func IOHIDQueueSetCancelHandler(arg0 IOHIDQueueRef, arg1 unsafe.Pointer) {
	if _iOHIDQueueSetCancelHandler == nil {
		panic("iokit: symbol IOHIDQueueSetCancelHandler not loaded")
	}
	_iOHIDQueueSetCancelHandler(arg0, arg1)
}

var _iOHIDQueueSetDepth func(arg0 IOHIDQueueRef, arg1 int)

// IOHIDQueueSetDepth sets the depth of the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545846-iohidqueuesetdepth
func IOHIDQueueSetDepth(arg0 IOHIDQueueRef, arg1 int) {
	if _iOHIDQueueSetDepth == nil {
		panic("iokit: symbol IOHIDQueueSetDepth not loaded")
	}
	_iOHIDQueueSetDepth(arg0, arg1)
}

var _iOHIDQueueSetDispatchQueue func(arg0 IOHIDQueueRef, arg1 uintptr)

// IOHIDQueueSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3042794-iohidqueuesetdispatchqueue
func IOHIDQueueSetDispatchQueue(arg0 IOHIDQueueRef, arg1 dispatch.Queue) {
	if _iOHIDQueueSetDispatchQueue == nil {
		panic("iokit: symbol IOHIDQueueSetDispatchQueue not loaded")
	}
	_iOHIDQueueSetDispatchQueue(arg0, uintptr(arg1.Handle()))
}

var _iOHIDQueueStart func(arg0 IOHIDQueueRef)

// IOHIDQueueStart starts element value delivery to the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545843-iohidqueuestart
func IOHIDQueueStart(arg0 IOHIDQueueRef) {
	if _iOHIDQueueStart == nil {
		panic("iokit: symbol IOHIDQueueStart not loaded")
	}
	_iOHIDQueueStart(arg0)
}

var _iOHIDQueueStop func(arg0 IOHIDQueueRef)

// IOHIDQueueStop stops element value delivery to the queue.
//
// See: https://developer.apple.com/documentation/iokit/1545830-iohidqueuestop
func IOHIDQueueStop(arg0 IOHIDQueueRef) {
	if _iOHIDQueueStop == nil {
		panic("iokit: symbol IOHIDQueueStop not loaded")
	}
	_iOHIDQueueStop(arg0)
}

var _iOHIDQueueUnscheduleFromRunLoop func(arg0 IOHIDQueueRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDQueueUnscheduleFromRunLoop unschedules queue with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1545834-iohidqueueunschedulefromrunloop
func IOHIDQueueUnscheduleFromRunLoop(arg0 IOHIDQueueRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDQueueUnscheduleFromRunLoop == nil {
		panic("iokit: symbol IOHIDQueueUnscheduleFromRunLoop not loaded")
	}
	_iOHIDQueueUnscheduleFromRunLoop(arg0, arg1, arg2)
}

var _iOHIDRegisterVirtualDisplay func(arg0 uintptr, arg1 uint32) int32

// IOHIDRegisterVirtualDisplay.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555416-iohidregistervirtualdisplay
func IOHIDRegisterVirtualDisplay(arg0 uintptr, arg1 uint32) int32 {
	if _iOHIDRegisterVirtualDisplay == nil {
		panic("iokit: symbol IOHIDRegisterVirtualDisplay not loaded")
	}
	return _iOHIDRegisterVirtualDisplay(arg0, arg1)
}

var _iOHIDRequestAccess func(arg0 IOHIDRequestType) bool

// IOHIDRequestAccess.
//
// See: https://developer.apple.com/documentation/iokit/3181574-iohidrequestaccess
func IOHIDRequestAccess(arg0 IOHIDRequestType) bool {
	if _iOHIDRequestAccess == nil {
		panic("iokit: symbol IOHIDRequestAccess not loaded")
	}
	return _iOHIDRequestAccess(arg0)
}

var _iOHIDServiceClientConformsTo func(arg0 IOHIDServiceClientRef, arg1 uint32, arg2 uint32) bool

// IOHIDServiceClientConformsTo.
//
// See: https://developer.apple.com/documentation/iokit/2269428-iohidserviceclientconformsto
func IOHIDServiceClientConformsTo(arg0 IOHIDServiceClientRef, arg1 uint32, arg2 uint32) bool {
	if _iOHIDServiceClientConformsTo == nil {
		panic("iokit: symbol IOHIDServiceClientConformsTo not loaded")
	}
	return _iOHIDServiceClientConformsTo(arg0, arg1, arg2)
}

var _iOHIDServiceClientCopyProperty func(arg0 IOHIDServiceClientRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOHIDServiceClientCopyProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269430-iohidserviceclientcopyproperty
func IOHIDServiceClientCopyProperty(arg0 IOHIDServiceClientRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOHIDServiceClientCopyProperty == nil {
		panic("iokit: symbol IOHIDServiceClientCopyProperty not loaded")
	}
	return _iOHIDServiceClientCopyProperty(arg0, arg1)
}

var _iOHIDServiceClientGetRegistryID func(arg0 IOHIDServiceClientRef) corefoundation.CFTypeRef

// IOHIDServiceClientGetRegistryID.
//
// See: https://developer.apple.com/documentation/iokit/2269426-iohidserviceclientgetregistryid
func IOHIDServiceClientGetRegistryID(arg0 IOHIDServiceClientRef) corefoundation.CFTypeRef {
	if _iOHIDServiceClientGetRegistryID == nil {
		panic("iokit: symbol IOHIDServiceClientGetRegistryID not loaded")
	}
	return _iOHIDServiceClientGetRegistryID(arg0)
}

var _iOHIDServiceClientGetTypeID func() uint

// IOHIDServiceClientGetTypeID.
//
// See: https://developer.apple.com/documentation/iokit/2269431-iohidserviceclientgettypeid
func IOHIDServiceClientGetTypeID() uint {
	if _iOHIDServiceClientGetTypeID == nil {
		panic("iokit: symbol IOHIDServiceClientGetTypeID not loaded")
	}
	return _iOHIDServiceClientGetTypeID()
}

var _iOHIDServiceClientSetProperty func(arg0 IOHIDServiceClientRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool

// IOHIDServiceClientSetProperty.
//
// See: https://developer.apple.com/documentation/iokit/2269429-iohidserviceclientsetproperty
func IOHIDServiceClientSetProperty(arg0 IOHIDServiceClientRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool {
	if _iOHIDServiceClientSetProperty == nil {
		panic("iokit: symbol IOHIDServiceClientSetProperty not loaded")
	}
	return _iOHIDServiceClientSetProperty(arg0, arg1, arg2)
}

var _iOHIDSetAccelerationWithKey func(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 float64) int32

// IOHIDSetAccelerationWithKey.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555398-iohidsetaccelerationwithkey
func IOHIDSetAccelerationWithKey(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 float64) int32 {
	if _iOHIDSetAccelerationWithKey == nil {
		panic("iokit: symbol IOHIDSetAccelerationWithKey not loaded")
	}
	return _iOHIDSetAccelerationWithKey(arg0, arg1, arg2)
}

var _iOHIDSetCFTypeParameter func(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) int32

// IOHIDSetCFTypeParameter.
//
// See: https://developer.apple.com/documentation/iokit/1555395-iohidsetcftypeparameter
func IOHIDSetCFTypeParameter(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) int32 {
	if _iOHIDSetCFTypeParameter == nil {
		panic("iokit: symbol IOHIDSetCFTypeParameter not loaded")
	}
	return _iOHIDSetCFTypeParameter(arg0, arg1, arg2)
}

var _iOHIDSetCursorEnable func(arg0 uintptr, arg1 bool) int32

// IOHIDSetCursorEnable.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555409-iohidsetcursorenable
func IOHIDSetCursorEnable(arg0 uintptr, arg1 bool) int32 {
	if _iOHIDSetCursorEnable == nil {
		panic("iokit: symbol IOHIDSetCursorEnable not loaded")
	}
	return _iOHIDSetCursorEnable(arg0, arg1)
}

var _iOHIDSetEventsEnable func(arg0 uintptr, arg1 bool) int32

// IOHIDSetEventsEnable.
//
// See: https://developer.apple.com/documentation/iokit/1555396-iohidseteventsenable
func IOHIDSetEventsEnable(arg0 uintptr, arg1 bool) int32 {
	if _iOHIDSetEventsEnable == nil {
		panic("iokit: symbol IOHIDSetEventsEnable not loaded")
	}
	return _iOHIDSetEventsEnable(arg0, arg1)
}

var _iOHIDSetModifierLockState func(arg0 uintptr, arg1 int, arg2 bool) int32

// IOHIDSetModifierLockState.
//
// See: https://developer.apple.com/documentation/iokit/1555420-iohidsetmodifierlockstate
func IOHIDSetModifierLockState(arg0 uintptr, arg1 int, arg2 bool) int32 {
	if _iOHIDSetModifierLockState == nil {
		panic("iokit: symbol IOHIDSetModifierLockState not loaded")
	}
	return _iOHIDSetModifierLockState(arg0, arg1, arg2)
}

var _iOHIDSetMouseAcceleration func(arg0 uintptr, arg1 float64) int32

// IOHIDSetMouseAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555390-iohidsetmouseacceleration
func IOHIDSetMouseAcceleration(arg0 uintptr, arg1 float64) int32 {
	if _iOHIDSetMouseAcceleration == nil {
		panic("iokit: symbol IOHIDSetMouseAcceleration not loaded")
	}
	return _iOHIDSetMouseAcceleration(arg0, arg1)
}

var _iOHIDSetMouseButtonMode func(arg0 uintptr, arg1 int) int32

// IOHIDSetMouseButtonMode.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555399-iohidsetmousebuttonmode
func IOHIDSetMouseButtonMode(arg0 uintptr, arg1 int) int32 {
	if _iOHIDSetMouseButtonMode == nil {
		panic("iokit: symbol IOHIDSetMouseButtonMode not loaded")
	}
	return _iOHIDSetMouseButtonMode(arg0, arg1)
}

var _iOHIDSetMouseLocation func(arg0 uintptr, arg1 int, arg2 int) int32

// IOHIDSetMouseLocation.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555402-iohidsetmouselocation
func IOHIDSetMouseLocation(arg0 uintptr, arg1 int, arg2 int) int32 {
	if _iOHIDSetMouseLocation == nil {
		panic("iokit: symbol IOHIDSetMouseLocation not loaded")
	}
	return _iOHIDSetMouseLocation(arg0, arg1, arg2)
}

var _iOHIDSetParameter func(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 uint) int32

// IOHIDSetParameter.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555394-iohidsetparameter
func IOHIDSetParameter(arg0 uintptr, arg1 corefoundation.CFStringRef, arg2 uint) int32 {
	if _iOHIDSetParameter == nil {
		panic("iokit: symbol IOHIDSetParameter not loaded")
	}
	return _iOHIDSetParameter(arg0, arg1, arg2)
}

var _iOHIDSetScrollAcceleration func(arg0 uintptr, arg1 float64) int32

// IOHIDSetScrollAcceleration.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1555391-iohidsetscrollacceleration
func IOHIDSetScrollAcceleration(arg0 uintptr, arg1 float64) int32 {
	if _iOHIDSetScrollAcceleration == nil {
		panic("iokit: symbol IOHIDSetScrollAcceleration not loaded")
	}
	return _iOHIDSetScrollAcceleration(arg0, arg1)
}

var _iOHIDSetStateForSelector func(arg0 uintptr, arg1 int, arg2 uint32) int32

// IOHIDSetStateForSelector.
//
// See: https://developer.apple.com/documentation/iokit/1555397-iohidsetstateforselector
func IOHIDSetStateForSelector(arg0 uintptr, arg1 int, arg2 uint32) int32 {
	if _iOHIDSetStateForSelector == nil {
		panic("iokit: symbol IOHIDSetStateForSelector not loaded")
	}
	return _iOHIDSetStateForSelector(arg0, arg1, arg2)
}

var _iOHIDSetVirtualDisplayBounds func(arg0 uintptr, arg1 uint32, arg2 IOGBounds) int32

// IOHIDSetVirtualDisplayBounds.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555392-iohidsetvirtualdisplaybounds
func IOHIDSetVirtualDisplayBounds(arg0 uintptr, arg1 uint32, arg2 IOGBounds) int32 {
	if _iOHIDSetVirtualDisplayBounds == nil {
		panic("iokit: symbol IOHIDSetVirtualDisplayBounds not loaded")
	}
	return _iOHIDSetVirtualDisplayBounds(arg0, arg1, arg2)
}

var _iOHIDTransactionAddElement func(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef)

// IOHIDTransactionAddElement adds an element to the transaction @disussion To minimize device traffic it is important to add elements that share a common report type and report id.
//
// See: https://developer.apple.com/documentation/iokit/1561679-iohidtransactionaddelement
func IOHIDTransactionAddElement(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef) {
	if _iOHIDTransactionAddElement == nil {
		panic("iokit: symbol IOHIDTransactionAddElement not loaded")
	}
	_iOHIDTransactionAddElement(arg0, arg1)
}

var _iOHIDTransactionClear func(arg0 IOHIDTransactionRef)

// IOHIDTransactionClear clears element transaction values.
//
// See: https://developer.apple.com/documentation/iokit/1561687-iohidtransactionclear
func IOHIDTransactionClear(arg0 IOHIDTransactionRef) {
	if _iOHIDTransactionClear == nil {
		panic("iokit: symbol IOHIDTransactionClear not loaded")
	}
	_iOHIDTransactionClear(arg0)
}

var _iOHIDTransactionCommit func(arg0 IOHIDTransactionRef) int

// IOHIDTransactionCommit synchronously commits element transaction to the device.
//
// See: https://developer.apple.com/documentation/iokit/1561681-iohidtransactioncommit
func IOHIDTransactionCommit(arg0 IOHIDTransactionRef) int {
	if _iOHIDTransactionCommit == nil {
		panic("iokit: symbol IOHIDTransactionCommit not loaded")
	}
	return _iOHIDTransactionCommit(arg0)
}

var _iOHIDTransactionCommitWithCallback func(arg0 IOHIDTransactionRef, arg1 float64, arg2 unsafe.Pointer) int

// IOHIDTransactionCommitWithCallback commits element transaction to the device.
//
// See: https://developer.apple.com/documentation/iokit/1561677-iohidtransactioncommitwithcallba
func IOHIDTransactionCommitWithCallback(arg0 IOHIDTransactionRef, arg1 float64, arg2 unsafe.Pointer) int {
	if _iOHIDTransactionCommitWithCallback == nil {
		panic("iokit: symbol IOHIDTransactionCommitWithCallback not loaded")
	}
	return _iOHIDTransactionCommitWithCallback(arg0, arg1, arg2)
}

var _iOHIDTransactionContainsElement func(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef) bool

// IOHIDTransactionContainsElement queries the transaction to determine if elemement has been added.
//
// See: https://developer.apple.com/documentation/iokit/1561680-iohidtransactioncontainselement
func IOHIDTransactionContainsElement(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef) bool {
	if _iOHIDTransactionContainsElement == nil {
		panic("iokit: symbol IOHIDTransactionContainsElement not loaded")
	}
	return _iOHIDTransactionContainsElement(arg0, arg1)
}

var _iOHIDTransactionCreate func(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDDeviceRef, arg2 IOHIDTransactionDirectionType, arg3 uint32) IOHIDTransactionRef

// IOHIDTransactionCreate creates an IOHIDTransaction object for the specified device.
//
// See: https://developer.apple.com/documentation/iokit/1561689-iohidtransactioncreate
func IOHIDTransactionCreate(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDDeviceRef, arg2 IOHIDTransactionDirectionType, arg3 uint32) IOHIDTransactionRef {
	if _iOHIDTransactionCreate == nil {
		panic("iokit: symbol IOHIDTransactionCreate not loaded")
	}
	return _iOHIDTransactionCreate(arg0, arg1, arg2, arg3)
}

var _iOHIDTransactionGetDevice func(arg0 IOHIDTransactionRef) IOHIDDeviceRef

// IOHIDTransactionGetDevice obtain the device associated with the transaction.
//
// See: https://developer.apple.com/documentation/iokit/1561685-iohidtransactiongetdevice
func IOHIDTransactionGetDevice(arg0 IOHIDTransactionRef) IOHIDDeviceRef {
	if _iOHIDTransactionGetDevice == nil {
		panic("iokit: symbol IOHIDTransactionGetDevice not loaded")
	}
	return _iOHIDTransactionGetDevice(arg0)
}

var _iOHIDTransactionGetDirection func(arg0 IOHIDTransactionRef) IOHIDTransactionDirectionType

// IOHIDTransactionGetDirection obtain the direction of the transaction.
//
// See: https://developer.apple.com/documentation/iokit/1561674-iohidtransactiongetdirection
func IOHIDTransactionGetDirection(arg0 IOHIDTransactionRef) IOHIDTransactionDirectionType {
	if _iOHIDTransactionGetDirection == nil {
		panic("iokit: symbol IOHIDTransactionGetDirection not loaded")
	}
	return _iOHIDTransactionGetDirection(arg0)
}

var _iOHIDTransactionGetTypeID func() uint

// IOHIDTransactionGetTypeID returns the type identifier of all IOHIDTransaction instances.
//
// See: https://developer.apple.com/documentation/iokit/1561678-iohidtransactiongettypeid
func IOHIDTransactionGetTypeID() uint {
	if _iOHIDTransactionGetTypeID == nil {
		panic("iokit: symbol IOHIDTransactionGetTypeID not loaded")
	}
	return _iOHIDTransactionGetTypeID()
}

var _iOHIDTransactionGetValue func(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef, arg2 uint32) IOHIDValueRef

// IOHIDTransactionGetValue obtains the value for a transaction element.
//
// See: https://developer.apple.com/documentation/iokit/1561683-iohidtransactiongetvalue
func IOHIDTransactionGetValue(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef, arg2 uint32) IOHIDValueRef {
	if _iOHIDTransactionGetValue == nil {
		panic("iokit: symbol IOHIDTransactionGetValue not loaded")
	}
	return _iOHIDTransactionGetValue(arg0, arg1, arg2)
}

var _iOHIDTransactionRemoveElement func(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef)

// IOHIDTransactionRemoveElement removes an element to the transaction
//
// See: https://developer.apple.com/documentation/iokit/1561686-iohidtransactionremoveelement
func IOHIDTransactionRemoveElement(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef) {
	if _iOHIDTransactionRemoveElement == nil {
		panic("iokit: symbol IOHIDTransactionRemoveElement not loaded")
	}
	_iOHIDTransactionRemoveElement(arg0, arg1)
}

var _iOHIDTransactionScheduleWithRunLoop func(arg0 IOHIDTransactionRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDTransactionScheduleWithRunLoop schedules transaction with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1561675-iohidtransactionschedulewithrunl
func IOHIDTransactionScheduleWithRunLoop(arg0 IOHIDTransactionRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDTransactionScheduleWithRunLoop == nil {
		panic("iokit: symbol IOHIDTransactionScheduleWithRunLoop not loaded")
	}
	_iOHIDTransactionScheduleWithRunLoop(arg0, arg1, arg2)
}

var _iOHIDTransactionSetDirection func(arg0 IOHIDTransactionRef, arg1 IOHIDTransactionDirectionType)

// IOHIDTransactionSetDirection sets the direction of the transaction @disussion This method is useful for manipulating bi-direction (feature) elements such that you can set or get element values without creating an additional transaction object.
//
// See: https://developer.apple.com/documentation/iokit/1561688-iohidtransactionsetdirection
func IOHIDTransactionSetDirection(arg0 IOHIDTransactionRef, arg1 IOHIDTransactionDirectionType) {
	if _iOHIDTransactionSetDirection == nil {
		panic("iokit: symbol IOHIDTransactionSetDirection not loaded")
	}
	_iOHIDTransactionSetDirection(arg0, arg1)
}

var _iOHIDTransactionSetValue func(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 uint32)

// IOHIDTransactionSetValue sets the value for a transaction element.
//
// See: https://developer.apple.com/documentation/iokit/1561676-iohidtransactionsetvalue
func IOHIDTransactionSetValue(arg0 IOHIDTransactionRef, arg1 IOHIDElementRef, arg2 IOHIDValueRef, arg3 uint32) {
	if _iOHIDTransactionSetValue == nil {
		panic("iokit: symbol IOHIDTransactionSetValue not loaded")
	}
	_iOHIDTransactionSetValue(arg0, arg1, arg2, arg3)
}

var _iOHIDTransactionUnscheduleFromRunLoop func(arg0 IOHIDTransactionRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef)

// IOHIDTransactionUnscheduleFromRunLoop unschedules transaction with run loop.
//
// See: https://developer.apple.com/documentation/iokit/1561682-iohidtransactionunschedulefromru
func IOHIDTransactionUnscheduleFromRunLoop(arg0 IOHIDTransactionRef, arg1 corefoundation.CFRunLoopRef, arg2 corefoundation.CFStringRef) {
	if _iOHIDTransactionUnscheduleFromRunLoop == nil {
		panic("iokit: symbol IOHIDTransactionUnscheduleFromRunLoop not loaded")
	}
	_iOHIDTransactionUnscheduleFromRunLoop(arg0, arg1, arg2)
}

var _iOHIDUnregisterVirtualDisplay func(arg0 uintptr, arg1 uint32) int32

// IOHIDUnregisterVirtualDisplay.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/iokit/1555410-iohidunregistervirtualdisplay
func IOHIDUnregisterVirtualDisplay(arg0 uintptr, arg1 uint32) int32 {
	if _iOHIDUnregisterVirtualDisplay == nil {
		panic("iokit: symbol IOHIDUnregisterVirtualDisplay not loaded")
	}
	return _iOHIDUnregisterVirtualDisplay(arg0, arg1)
}

var _iOHIDUserDeviceActivate func(arg0 IOHIDUserDeviceRef)

// IOHIDUserDeviceActivate.
//
// See: https://developer.apple.com/documentation/iokit/3334949-iohiduserdeviceactivate
func IOHIDUserDeviceActivate(arg0 IOHIDUserDeviceRef) {
	if _iOHIDUserDeviceActivate == nil {
		panic("iokit: symbol IOHIDUserDeviceActivate not loaded")
	}
	_iOHIDUserDeviceActivate(arg0)
}

var _iOHIDUserDeviceCancel func(arg0 IOHIDUserDeviceRef)

// IOHIDUserDeviceCancel.
//
// See: https://developer.apple.com/documentation/iokit/3334950-iohiduserdevicecancel
func IOHIDUserDeviceCancel(arg0 IOHIDUserDeviceRef) {
	if _iOHIDUserDeviceCancel == nil {
		panic("iokit: symbol IOHIDUserDeviceCancel not loaded")
	}
	_iOHIDUserDeviceCancel(arg0)
}

var _iOHIDUserDeviceCopyProperty func(arg0 IOHIDUserDeviceRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef

// IOHIDUserDeviceCopyProperty.
//
// See: https://developer.apple.com/documentation/iokit/3334951-iohiduserdevicecopyproperty
func IOHIDUserDeviceCopyProperty(arg0 IOHIDUserDeviceRef, arg1 corefoundation.CFStringRef) corefoundation.CFTypeRef {
	if _iOHIDUserDeviceCopyProperty == nil {
		panic("iokit: symbol IOHIDUserDeviceCopyProperty not loaded")
	}
	return _iOHIDUserDeviceCopyProperty(arg0, arg1)
}

var _iOHIDUserDeviceCreateWithProperties func(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFDictionaryRef, arg2 uint32) IOHIDUserDeviceRef

// IOHIDUserDeviceCreateWithProperties.
//
// See: https://developer.apple.com/documentation/iokit/3334952-iohiduserdevicecreatewithpropert
func IOHIDUserDeviceCreateWithProperties(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFDictionaryRef, arg2 uint32) IOHIDUserDeviceRef {
	if _iOHIDUserDeviceCreateWithProperties == nil {
		panic("iokit: symbol IOHIDUserDeviceCreateWithProperties not loaded")
	}
	return _iOHIDUserDeviceCreateWithProperties(arg0, arg1, arg2)
}

var _iOHIDUserDeviceGetTypeID func() uint

// IOHIDUserDeviceGetTypeID.
//
// See: https://developer.apple.com/documentation/iokit/3334954-iohiduserdevicegettypeid
func IOHIDUserDeviceGetTypeID() uint {
	if _iOHIDUserDeviceGetTypeID == nil {
		panic("iokit: symbol IOHIDUserDeviceGetTypeID not loaded")
	}
	return _iOHIDUserDeviceGetTypeID()
}

var _iOHIDUserDeviceHandleReportWithTimeStamp func(arg0 IOHIDUserDeviceRef, arg1 uint64, arg2 uint8, arg3 int) int

// IOHIDUserDeviceHandleReportWithTimeStamp.
//
// See: https://developer.apple.com/documentation/iokit/3334955-iohiduserdevicehandlereportwitht
func IOHIDUserDeviceHandleReportWithTimeStamp(arg0 IOHIDUserDeviceRef, arg1 uint64, arg2 uint8, arg3 int) int {
	if _iOHIDUserDeviceHandleReportWithTimeStamp == nil {
		panic("iokit: symbol IOHIDUserDeviceHandleReportWithTimeStamp not loaded")
	}
	return _iOHIDUserDeviceHandleReportWithTimeStamp(arg0, arg1, arg2, arg3)
}

var _iOHIDUserDeviceRegisterGetReportBlock func(arg0 IOHIDUserDeviceRef, arg1 IOHIDUserDeviceGetReportBlock)

// IOHIDUserDeviceRegisterGetReportBlock.
//
// See: https://developer.apple.com/documentation/iokit/3334959-iohiduserdeviceregistergetreport
func IOHIDUserDeviceRegisterGetReportBlock(arg0 IOHIDUserDeviceRef, arg1 IOHIDUserDeviceGetReportBlock) {
	if _iOHIDUserDeviceRegisterGetReportBlock == nil {
		panic("iokit: symbol IOHIDUserDeviceRegisterGetReportBlock not loaded")
	}
	_iOHIDUserDeviceRegisterGetReportBlock(arg0, arg1)
}

var _iOHIDUserDeviceRegisterSetReportBlock func(arg0 IOHIDUserDeviceRef, arg1 IOHIDUserDeviceSetReportBlock)

// IOHIDUserDeviceRegisterSetReportBlock.
//
// See: https://developer.apple.com/documentation/iokit/3334960-iohiduserdeviceregistersetreport
func IOHIDUserDeviceRegisterSetReportBlock(arg0 IOHIDUserDeviceRef, arg1 IOHIDUserDeviceSetReportBlock) {
	if _iOHIDUserDeviceRegisterSetReportBlock == nil {
		panic("iokit: symbol IOHIDUserDeviceRegisterSetReportBlock not loaded")
	}
	_iOHIDUserDeviceRegisterSetReportBlock(arg0, arg1)
}

var _iOHIDUserDeviceSetCancelHandler func(arg0 IOHIDUserDeviceRef, arg1 unsafe.Pointer)

// IOHIDUserDeviceSetCancelHandler.
//
// See: https://developer.apple.com/documentation/iokit/3334961-iohiduserdevicesetcancelhandler
func IOHIDUserDeviceSetCancelHandler(arg0 IOHIDUserDeviceRef, arg1 unsafe.Pointer) {
	if _iOHIDUserDeviceSetCancelHandler == nil {
		panic("iokit: symbol IOHIDUserDeviceSetCancelHandler not loaded")
	}
	_iOHIDUserDeviceSetCancelHandler(arg0, arg1)
}

var _iOHIDUserDeviceSetDispatchQueue func(arg0 IOHIDUserDeviceRef, arg1 uintptr)

// IOHIDUserDeviceSetDispatchQueue.
//
// See: https://developer.apple.com/documentation/iokit/3334962-iohiduserdevicesetdispatchqueue
func IOHIDUserDeviceSetDispatchQueue(arg0 IOHIDUserDeviceRef, arg1 dispatch.Queue) {
	if _iOHIDUserDeviceSetDispatchQueue == nil {
		panic("iokit: symbol IOHIDUserDeviceSetDispatchQueue not loaded")
	}
	_iOHIDUserDeviceSetDispatchQueue(arg0, uintptr(arg1.Handle()))
}

var _iOHIDUserDeviceSetProperty func(arg0 IOHIDUserDeviceRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool

// IOHIDUserDeviceSetProperty.
//
// See: https://developer.apple.com/documentation/iokit/3334963-iohiduserdevicesetproperty
func IOHIDUserDeviceSetProperty(arg0 IOHIDUserDeviceRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) bool {
	if _iOHIDUserDeviceSetProperty == nil {
		panic("iokit: symbol IOHIDUserDeviceSetProperty not loaded")
	}
	return _iOHIDUserDeviceSetProperty(arg0, arg1, arg2)
}

var _iOHIDValueCreateWithBytes func(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDElementRef, arg2 uint64, arg3 uint8, arg4 int) IOHIDValueRef

// IOHIDValueCreateWithBytes creates a new element value using byte data.
//
// See: https://developer.apple.com/documentation/iokit/1433290-iohidvaluecreatewithbytes
func IOHIDValueCreateWithBytes(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDElementRef, arg2 uint64, arg3 uint8, arg4 int) IOHIDValueRef {
	if _iOHIDValueCreateWithBytes == nil {
		panic("iokit: symbol IOHIDValueCreateWithBytes not loaded")
	}
	return _iOHIDValueCreateWithBytes(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDValueCreateWithBytesNoCopy func(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDElementRef, arg2 uint64, arg3 uint8, arg4 int) IOHIDValueRef

// IOHIDValueCreateWithBytesNoCopy creates a new element value using byte data without performing a copy.
//
// See: https://developer.apple.com/documentation/iokit/1433287-iohidvaluecreatewithbytesnocopy
func IOHIDValueCreateWithBytesNoCopy(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDElementRef, arg2 uint64, arg3 uint8, arg4 int) IOHIDValueRef {
	if _iOHIDValueCreateWithBytesNoCopy == nil {
		panic("iokit: symbol IOHIDValueCreateWithBytesNoCopy not loaded")
	}
	return _iOHIDValueCreateWithBytesNoCopy(arg0, arg1, arg2, arg3, arg4)
}

var _iOHIDValueCreateWithIntegerValue func(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDElementRef, arg2 uint64, arg3 int) IOHIDValueRef

// IOHIDValueCreateWithIntegerValue creates a new element value using an integer value.
//
// See: https://developer.apple.com/documentation/iokit/1433294-iohidvaluecreatewithintegervalue
func IOHIDValueCreateWithIntegerValue(arg0 corefoundation.CFAllocatorRef, arg1 IOHIDElementRef, arg2 uint64, arg3 int) IOHIDValueRef {
	if _iOHIDValueCreateWithIntegerValue == nil {
		panic("iokit: symbol IOHIDValueCreateWithIntegerValue not loaded")
	}
	return _iOHIDValueCreateWithIntegerValue(arg0, arg1, arg2, arg3)
}

var _iOHIDValueGetBytePtr func(arg0 IOHIDValueRef) *uint8

// IOHIDValueGetBytePtr returns a byte pointer to the value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433292-iohidvaluegetbyteptr
func IOHIDValueGetBytePtr(arg0 IOHIDValueRef) *uint8 {
	if _iOHIDValueGetBytePtr == nil {
		panic("iokit: symbol IOHIDValueGetBytePtr not loaded")
	}
	return _iOHIDValueGetBytePtr(arg0)
}

var _iOHIDValueGetElement func(arg0 IOHIDValueRef) IOHIDElementRef

// IOHIDValueGetElement returns the element value associated with this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433285-iohidvaluegetelement
func IOHIDValueGetElement(arg0 IOHIDValueRef) IOHIDElementRef {
	if _iOHIDValueGetElement == nil {
		panic("iokit: symbol IOHIDValueGetElement not loaded")
	}
	return _iOHIDValueGetElement(arg0)
}

var _iOHIDValueGetIntegerValue func(arg0 IOHIDValueRef) int

// IOHIDValueGetIntegerValue returns an integer representaion of the value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433289-iohidvaluegetintegervalue
func IOHIDValueGetIntegerValue(arg0 IOHIDValueRef) int {
	if _iOHIDValueGetIntegerValue == nil {
		panic("iokit: symbol IOHIDValueGetIntegerValue not loaded")
	}
	return _iOHIDValueGetIntegerValue(arg0)
}

var _iOHIDValueGetLength func(arg0 IOHIDValueRef) int

// IOHIDValueGetLength returns the size, in bytes, of the value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433291-iohidvaluegetlength
func IOHIDValueGetLength(arg0 IOHIDValueRef) int {
	if _iOHIDValueGetLength == nil {
		panic("iokit: symbol IOHIDValueGetLength not loaded")
	}
	return _iOHIDValueGetLength(arg0)
}

var _iOHIDValueGetScaledValue func(arg0 IOHIDValueRef, arg1 IOHIDValueScaleType) kernel.Double_t

// IOHIDValueGetScaledValue returns an scaled representaion of the value contained in this IOHIDValueRef based on the scale type.
//
// See: https://developer.apple.com/documentation/iokit/1433288-iohidvaluegetscaledvalue
func IOHIDValueGetScaledValue(arg0 IOHIDValueRef, arg1 IOHIDValueScaleType) kernel.Double_t {
	if _iOHIDValueGetScaledValue == nil {
		panic("iokit: symbol IOHIDValueGetScaledValue not loaded")
	}
	return _iOHIDValueGetScaledValue(arg0, arg1)
}

var _iOHIDValueGetTimeStamp func(arg0 IOHIDValueRef) uint64

// IOHIDValueGetTimeStamp returns the timestamp value contained in this IOHIDValueRef.
//
// See: https://developer.apple.com/documentation/iokit/1433286-iohidvaluegettimestamp
func IOHIDValueGetTimeStamp(arg0 IOHIDValueRef) uint64 {
	if _iOHIDValueGetTimeStamp == nil {
		panic("iokit: symbol IOHIDValueGetTimeStamp not loaded")
	}
	return _iOHIDValueGetTimeStamp(arg0)
}

var _iOHIDValueGetTypeID func() uint

// IOHIDValueGetTypeID returns the type identifier of all IOHIDValue instances.
//
// See: https://developer.apple.com/documentation/iokit/1433293-iohidvaluegettypeid
func IOHIDValueGetTypeID() uint {
	if _iOHIDValueGetTypeID == nil {
		panic("iokit: symbol IOHIDValueGetTypeID not loaded")
	}
	return _iOHIDValueGetTypeID()
}

var _iOI2CCopyInterfaceForID func(arg0 corefoundation.CFTypeRef, arg1 uintptr) int

// IOI2CCopyInterfaceForID.
//
// See: https://developer.apple.com/documentation/iokit/1410384-ioi2ccopyinterfaceforid
func IOI2CCopyInterfaceForID(arg0 corefoundation.CFTypeRef, arg1 uintptr) int {
	if _iOI2CCopyInterfaceForID == nil {
		panic("iokit: symbol IOI2CCopyInterfaceForID not loaded")
	}
	return _iOI2CCopyInterfaceForID(arg0, arg1)
}

var _iOI2CInterfaceClose func(arg0 IOI2CConnectRef, arg1 uint32) int

// IOI2CInterfaceClose closes an IOI2CConnectRef.
//
// See: https://developer.apple.com/documentation/iokit/1410390-ioi2cinterfaceclose
func IOI2CInterfaceClose(arg0 IOI2CConnectRef, arg1 uint32) int {
	if _iOI2CInterfaceClose == nil {
		panic("iokit: symbol IOI2CInterfaceClose not loaded")
	}
	return _iOI2CInterfaceClose(arg0, arg1)
}

var _iOI2CInterfaceOpen func(arg0 uintptr, arg1 uint32, arg2 IOI2CConnectRef) int

// IOI2CInterfaceOpen opens an instance of an I2C bus interface, allowing I2C requests to be made.
//
// See: https://developer.apple.com/documentation/iokit/1410388-ioi2cinterfaceopen
func IOI2CInterfaceOpen(arg0 uintptr, arg1 uint32, arg2 IOI2CConnectRef) int {
	if _iOI2CInterfaceOpen == nil {
		panic("iokit: symbol IOI2CInterfaceOpen not loaded")
	}
	return _iOI2CInterfaceOpen(arg0, arg1, arg2)
}

var _iOI2CSendRequest func(arg0 IOI2CConnectRef, arg1 uint32, arg2 IOI2CRequest) int

// IOI2CSendRequest carries out the I2C transaction specified by an IOI2CRequest structure.
//
// See: https://developer.apple.com/documentation/iokit/1410373-ioi2csendrequest
func IOI2CSendRequest(arg0 IOI2CConnectRef, arg1 uint32, arg2 IOI2CRequest) int {
	if _iOI2CSendRequest == nil {
		panic("iokit: symbol IOI2CSendRequest not loaded")
	}
	return _iOI2CSendRequest(arg0, arg1, arg2)
}

var _iOIteratorIsValid func(iterator uintptr) bool

// IOIteratorIsValid checks an iterator is still valid.
//
// See: https://developer.apple.com/documentation/iokit/1514556-ioiteratorisvalid
func IOIteratorIsValid(iterator uintptr) bool {
	if _iOIteratorIsValid == nil {
		panic("iokit: symbol IOIteratorIsValid not loaded")
	}
	return _iOIteratorIsValid(iterator)
}

var _iOIteratorNext func(iterator uintptr) uintptr

// IOIteratorNext returns the next object in an iteration.
//
// See: https://developer.apple.com/documentation/iokit/1514741-ioiteratornext
func IOIteratorNext(iterator uintptr) uintptr {
	if _iOIteratorNext == nil {
		panic("iokit: symbol IOIteratorNext not loaded")
	}
	return _iOIteratorNext(iterator)
}

var _iOIteratorReset func(iterator uintptr)

// IOIteratorReset resets an iteration back to the beginning.
//
// See: https://developer.apple.com/documentation/iokit/1514379-ioiteratorreset
func IOIteratorReset(iterator uintptr) {
	if _iOIteratorReset == nil {
		panic("iokit: symbol IOIteratorReset not loaded")
	}
	_iOIteratorReset(iterator)
}

var _iOKitGetBusyState func(mainPort uint32, busyState *uint32) int32

// IOKitGetBusyState returns the busyState of all IOServices.
//
// See: https://developer.apple.com/documentation/iokit/1514460-iokitgetbusystate
func IOKitGetBusyState(mainPort uint32, busyState *uint32) int32 {
	if _iOKitGetBusyState == nil {
		panic("iokit: symbol IOKitGetBusyState not loaded")
	}
	return _iOKitGetBusyState(mainPort, busyState)
}

var _iOKitWaitQuiet func(mainPort uint32, waitTime *kernel.Mach_timespec_t) int32

// IOKitWaitQuiet wait for a all IOServices' busyState to be zero.
//
// See: https://developer.apple.com/documentation/iokit/1514440-iokitwaitquiet
func IOKitWaitQuiet(mainPort uint32, waitTime *kernel.Mach_timespec_t) int32 {
	if _iOKitWaitQuiet == nil {
		panic("iokit: symbol IOKitWaitQuiet not loaded")
	}
	return _iOKitWaitQuiet(mainPort, waitTime)
}

var _iOMainPort func(bootstrapPort uint32, mainPort *uint32) int32

// IOMainPort.
//
// See: https://developer.apple.com/documentation/iokit/3753260-iomainport
func IOMainPort(bootstrapPort uint32, mainPort *uint32) int32 {
	if _iOMainPort == nil {
		panic("iokit: symbol IOMainPort not loaded")
	}
	return _iOMainPort(bootstrapPort, mainPort)
}

var _iOMasterPort func(bootstrapPort uint32, mainPort *uint32) int32

// IOMasterPort returns the mach port used to initiate communication with IOKit.
//
// Deprecated: Deprecated since macOS 12.0.
//
// See: https://developer.apple.com/documentation/iokit/1514652-iomasterport
func IOMasterPort(bootstrapPort uint32, mainPort *uint32) int32 {
	if _iOMasterPort == nil {
		panic("iokit: symbol IOMasterPort not loaded")
	}
	return _iOMasterPort(bootstrapPort, mainPort)
}

var _iONetworkClose func(arg0 uintptr) int

// IONetworkClose close the connection to an IONetworkInterface object.
//
// See: https://developer.apple.com/documentation/iokit/1572704-ionetworkclose
func IONetworkClose(arg0 uintptr) int {
	if _iONetworkClose == nil {
		panic("iokit: symbol IONetworkClose not loaded")
	}
	return _iONetworkClose(arg0)
}

var _iONetworkGetDataCapacity func(arg0 uintptr, arg1 IONDHandle, arg2 uint32) int

// IONetworkGetDataCapacity get the capacity (in bytes) of a network data object.
//
// See: https://developer.apple.com/documentation/iokit/1572712-ionetworkgetdatacapacity
func IONetworkGetDataCapacity(arg0 uintptr, arg1 IONDHandle, arg2 uint32) int {
	if _iONetworkGetDataCapacity == nil {
		panic("iokit: symbol IONetworkGetDataCapacity not loaded")
	}
	return _iONetworkGetDataCapacity(arg0, arg1, arg2)
}

var _iONetworkGetDataHandle func(arg0 uintptr, arg1 int8, arg2 IONDHandle) int

// IONetworkGetDataHandle get the handle of a network data object with the given name.
//
// See: https://developer.apple.com/documentation/iokit/1572708-ionetworkgetdatahandle
func IONetworkGetDataHandle(arg0 uintptr, arg1 int8, arg2 IONDHandle) int {
	if _iONetworkGetDataHandle == nil {
		panic("iokit: symbol IONetworkGetDataHandle not loaded")
	}
	return _iONetworkGetDataHandle(arg0, arg1, arg2)
}

var _iONetworkGetPacketFiltersMask func(arg0 uintptr, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32) int

// IONetworkGetPacketFiltersMask get the packet filters for a given filter group.
//
// See: https://developer.apple.com/documentation/iokit/1572711-ionetworkgetpacketfiltersmask
func IONetworkGetPacketFiltersMask(arg0 uintptr, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32) int {
	if _iONetworkGetPacketFiltersMask == nil {
		panic("iokit: symbol IONetworkGetPacketFiltersMask not loaded")
	}
	return _iONetworkGetPacketFiltersMask(arg0, arg1, arg2, arg3)
}

var _iONetworkOpen func(arg0 uintptr, arg1 uintptr) int

// IONetworkOpen open a connection to an IONetworkInterface object.
//
// See: https://developer.apple.com/documentation/iokit/1572709-ionetworkopen
func IONetworkOpen(arg0 uintptr, arg1 uintptr) int {
	if _iONetworkOpen == nil {
		panic("iokit: symbol IONetworkOpen not loaded")
	}
	return _iONetworkOpen(arg0, arg1)
}

var _iONetworkReadData func(arg0 uintptr, arg1 IONDHandle, arg2 uint8, arg3 uint32) int

// IONetworkReadData read the buffer of a network data object.
//
// See: https://developer.apple.com/documentation/iokit/1572706-ionetworkreaddata
func IONetworkReadData(arg0 uintptr, arg1 IONDHandle, arg2 uint8, arg3 uint32) int {
	if _iONetworkReadData == nil {
		panic("iokit: symbol IONetworkReadData not loaded")
	}
	return _iONetworkReadData(arg0, arg1, arg2, arg3)
}

var _iONetworkResetData func(arg0 uintptr, arg1 IONDHandle) int

// IONetworkResetData fill the buffer of a network data object with zeroes.
//
// See: https://developer.apple.com/documentation/iokit/1572710-ionetworkresetdata
func IONetworkResetData(arg0 uintptr, arg1 IONDHandle) int {
	if _iONetworkResetData == nil {
		panic("iokit: symbol IONetworkResetData not loaded")
	}
	return _iONetworkResetData(arg0, arg1)
}

var _iONetworkSetPacketFiltersMask func(arg0 uintptr, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32) int

// IONetworkSetPacketFiltersMask set the packet filters for a given filter group.
//
// See: https://developer.apple.com/documentation/iokit/1572703-ionetworksetpacketfiltersmask
func IONetworkSetPacketFiltersMask(arg0 uintptr, arg1 unsafe.Pointer, arg2 uint32, arg3 uint32) int {
	if _iONetworkSetPacketFiltersMask == nil {
		panic("iokit: symbol IONetworkSetPacketFiltersMask not loaded")
	}
	return _iONetworkSetPacketFiltersMask(arg0, arg1, arg2, arg3)
}

var _iONetworkWriteData func(arg0 uintptr, arg1 IONDHandle, arg2 uint8, arg3 uint32) int

// IONetworkWriteData write to the buffer of a network data object.
//
// See: https://developer.apple.com/documentation/iokit/1572707-ionetworkwritedata
func IONetworkWriteData(arg0 uintptr, arg1 IONDHandle, arg2 uint8, arg3 uint32) int {
	if _iONetworkWriteData == nil {
		panic("iokit: symbol IONetworkWriteData not loaded")
	}
	return _iONetworkWriteData(arg0, arg1, arg2, arg3)
}

var _iONotificationPortCreate func(mainPort uint32) IONotificationPortRef

// IONotificationPortCreate creates and returns a notification object for receiving IOKit notifications of new devices or state changes.
//
// See: https://developer.apple.com/documentation/iokit/1514480-ionotificationportcreate
func IONotificationPortCreate(mainPort uint32) IONotificationPortRef {
	if _iONotificationPortCreate == nil {
		panic("iokit: symbol IONotificationPortCreate not loaded")
	}
	return _iONotificationPortCreate(mainPort)
}

var _iONotificationPortDestroy func(notify IONotificationPortRef)

// IONotificationPortDestroy destroys a notification object created with IONotificationPortCreate.
//
// See: https://developer.apple.com/documentation/iokit/1514751-ionotificationportdestroy
func IONotificationPortDestroy(notify IONotificationPortRef) {
	if _iONotificationPortDestroy == nil {
		panic("iokit: symbol IONotificationPortDestroy not loaded")
	}
	_iONotificationPortDestroy(notify)
}

var _iONotificationPortGetMachPort func(notify IONotificationPortRef) uint32

// IONotificationPortGetMachPort returns a mach_port to be used to listen for notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514875-ionotificationportgetmachport
func IONotificationPortGetMachPort(notify IONotificationPortRef) uint32 {
	if _iONotificationPortGetMachPort == nil {
		panic("iokit: symbol IONotificationPortGetMachPort not loaded")
	}
	return _iONotificationPortGetMachPort(notify)
}

var _iONotificationPortGetRunLoopSource func(notify IONotificationPortRef) corefoundation.CFRunLoopSourceRef

// IONotificationPortGetRunLoopSource returns a CFRunLoopSource to be used to listen for notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514599-ionotificationportgetrunloopsour
func IONotificationPortGetRunLoopSource(notify IONotificationPortRef) corefoundation.CFRunLoopSourceRef {
	if _iONotificationPortGetRunLoopSource == nil {
		panic("iokit: symbol IONotificationPortGetRunLoopSource not loaded")
	}
	return _iONotificationPortGetRunLoopSource(notify)
}

var _iONotificationPortSetDispatchQueue func(notify IONotificationPortRef, queue uintptr)

// IONotificationPortSetDispatchQueue sets a dispatch queue to be used to listen for notifications.
//
// See: https://developer.apple.com/documentation/iokit/1514596-ionotificationportsetdispatchque
func IONotificationPortSetDispatchQueue(notify IONotificationPortRef, queue dispatch.Queue) {
	if _iONotificationPortSetDispatchQueue == nil {
		panic("iokit: symbol IONotificationPortSetDispatchQueue not loaded")
	}
	_iONotificationPortSetDispatchQueue(notify, uintptr(queue.Handle()))
}

var _iONotificationPortSetImportanceReceiver func(notify IONotificationPortRef) int32

// IONotificationPortSetImportanceReceiver.
//
// See: https://developer.apple.com/documentation/iokit/2870065-ionotificationportsetimportancer
func IONotificationPortSetImportanceReceiver(notify IONotificationPortRef) int32 {
	if _iONotificationPortSetImportanceReceiver == nil {
		panic("iokit: symbol IONotificationPortSetImportanceReceiver not loaded")
	}
	return _iONotificationPortSetImportanceReceiver(notify)
}

var _iOObjectConformsTo func(object uintptr, className string) bool

// IOObjectConformsTo performs an OSDynamicCast operation on an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514505-ioobjectconformsto
func IOObjectConformsTo(object uintptr, className string) bool {
	if _iOObjectConformsTo == nil {
		panic("iokit: symbol IOObjectConformsTo not loaded")
	}
	return _iOObjectConformsTo(object, className)
}

var _iOObjectCopyBundleIdentifierForClass func(classname corefoundation.CFStringRef) corefoundation.CFStringRef

// IOObjectCopyBundleIdentifierForClass return the bundle identifier of the given class.
//
// See: https://developer.apple.com/documentation/iokit/1514375-ioobjectcopybundleidentifierforc
func IOObjectCopyBundleIdentifierForClass(classname corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _iOObjectCopyBundleIdentifierForClass == nil {
		panic("iokit: symbol IOObjectCopyBundleIdentifierForClass not loaded")
	}
	return _iOObjectCopyBundleIdentifierForClass(classname)
}

var _iOObjectCopyClass func(object uintptr) corefoundation.CFStringRef

// IOObjectCopyClass return the class name of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514781-ioobjectcopyclass
func IOObjectCopyClass(object uintptr) corefoundation.CFStringRef {
	if _iOObjectCopyClass == nil {
		panic("iokit: symbol IOObjectCopyClass not loaded")
	}
	return _iOObjectCopyClass(object)
}

var _iOObjectCopySuperclassForClass func(classname corefoundation.CFStringRef) corefoundation.CFStringRef

// IOObjectCopySuperclassForClass return the superclass name of the given class.
//
// See: https://developer.apple.com/documentation/iokit/1514635-ioobjectcopysuperclassforclass
func IOObjectCopySuperclassForClass(classname corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _iOObjectCopySuperclassForClass == nil {
		panic("iokit: symbol IOObjectCopySuperclassForClass not loaded")
	}
	return _iOObjectCopySuperclassForClass(classname)
}

var _iOObjectGetClass func(object uintptr, className string) int32

// IOObjectGetClass return the class name of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514756-ioobjectgetclass
func IOObjectGetClass(object uintptr, className string) int32 {
	if _iOObjectGetClass == nil {
		panic("iokit: symbol IOObjectGetClass not loaded")
	}
	return _iOObjectGetClass(object, className)
}

var _iOObjectGetKernelRetainCount func(object uintptr) uint32

// IOObjectGetKernelRetainCount returns kernel retain count of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514325-ioobjectgetkernelretaincount
func IOObjectGetKernelRetainCount(object uintptr) uint32 {
	if _iOObjectGetKernelRetainCount == nil {
		panic("iokit: symbol IOObjectGetKernelRetainCount not loaded")
	}
	return _iOObjectGetKernelRetainCount(object)
}

var _iOObjectGetRetainCount func(object uintptr) uint32

// IOObjectGetRetainCount returns kernel retain count of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514824-ioobjectgetretaincount
func IOObjectGetRetainCount(object uintptr) uint32 {
	if _iOObjectGetRetainCount == nil {
		panic("iokit: symbol IOObjectGetRetainCount not loaded")
	}
	return _iOObjectGetRetainCount(object)
}

var _iOObjectGetUserRetainCount func(object uintptr) uint32

// IOObjectGetUserRetainCount returns the retain count for the current process of an IOKit object.
//
// See: https://developer.apple.com/documentation/iokit/1514464-ioobjectgetuserretaincount
func IOObjectGetUserRetainCount(object uintptr) uint32 {
	if _iOObjectGetUserRetainCount == nil {
		panic("iokit: symbol IOObjectGetUserRetainCount not loaded")
	}
	return _iOObjectGetUserRetainCount(object)
}

var _iOObjectIsEqualTo func(object uintptr, anObject uintptr) bool

// IOObjectIsEqualTo checks two object handles to see if they represent the same kernel object.
//
// See: https://developer.apple.com/documentation/iokit/1514563-ioobjectisequalto
func IOObjectIsEqualTo(object uintptr, anObject uintptr) bool {
	if _iOObjectIsEqualTo == nil {
		panic("iokit: symbol IOObjectIsEqualTo not loaded")
	}
	return _iOObjectIsEqualTo(object, anObject)
}

var _iOObjectRelease func(object uintptr) int32

// IOObjectRelease releases an object handle previously returned by IOKitLib.
//
// See: https://developer.apple.com/documentation/iokit/1514627-ioobjectrelease
func IOObjectRelease(object uintptr) int32 {
	if _iOObjectRelease == nil {
		panic("iokit: symbol IOObjectRelease not loaded")
	}
	return _iOObjectRelease(object)
}

var _iOObjectRetain func(object uintptr) int32

// IOObjectRetain retains an object handle previously returned by IOKitLib.
//
// See: https://developer.apple.com/documentation/iokit/1514769-ioobjectretain
func IOObjectRetain(object uintptr) int32 {
	if _iOObjectRetain == nil {
		panic("iokit: symbol IOObjectRetain not loaded")
	}
	return _iOObjectRetain(object)
}

var _iOOpenFirmwarePathMatching func(mainPort uint32, options uint32, path string) corefoundation.CFMutableDictionary

// IOOpenFirmwarePathMatching.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/iokit/1514715-ioopenfirmwarepathmatching
func IOOpenFirmwarePathMatching(mainPort uint32, options uint32, path string) corefoundation.CFMutableDictionary {
	if _iOOpenFirmwarePathMatching == nil {
		panic("iokit: symbol IOOpenFirmwarePathMatching not loaded")
	}
	return _iOOpenFirmwarePathMatching(mainPort, options, path)
}

var _iOPMAssertionCopyProperties func(arg0 IOPMAssertionID) corefoundation.CFDictionaryRef

// IOPMAssertionCopyProperties copies details about an [IOPMAssertion]
//
// See: https://developer.apple.com/documentation/iokit/1557066-iopmassertioncopyproperties
func IOPMAssertionCopyProperties(arg0 IOPMAssertionID) corefoundation.CFDictionaryRef {
	if _iOPMAssertionCopyProperties == nil {
		panic("iokit: symbol IOPMAssertionCopyProperties not loaded")
	}
	return _iOPMAssertionCopyProperties(arg0)
}

var _iOPMAssertionCreate func(arg0 corefoundation.CFStringRef, arg1 IOPMAssertionLevel, arg2 IOPMAssertionID) int

// IOPMAssertionCreate dynamically requests a system behavior from the power management system.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/iokit/1557118-iopmassertioncreate
func IOPMAssertionCreate(arg0 corefoundation.CFStringRef, arg1 IOPMAssertionLevel, arg2 IOPMAssertionID) int {
	if _iOPMAssertionCreate == nil {
		panic("iokit: symbol IOPMAssertionCreate not loaded")
	}
	return _iOPMAssertionCreate(arg0, arg1, arg2)
}

var _iOPMAssertionCreateWithDescription func(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 corefoundation.CFStringRef, arg5 float64, arg6 corefoundation.CFStringRef, arg7 IOPMAssertionID) int

// IOPMAssertionCreateWithDescription.
//
// See: https://developer.apple.com/documentation/iokit/1557078-iopmassertioncreatewithdescripti
func IOPMAssertionCreateWithDescription(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef, arg3 corefoundation.CFStringRef, arg4 corefoundation.CFStringRef, arg5 float64, arg6 corefoundation.CFStringRef, arg7 IOPMAssertionID) int {
	if _iOPMAssertionCreateWithDescription == nil {
		panic("iokit: symbol IOPMAssertionCreateWithDescription not loaded")
	}
	return _iOPMAssertionCreateWithDescription(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

var _iOPMAssertionCreateWithName func(arg0 corefoundation.CFStringRef, arg1 IOPMAssertionLevel, arg2 corefoundation.CFStringRef, arg3 IOPMAssertionID) int

// IOPMAssertionCreateWithName dynamically requests a system behavior from the power management system.
//
// See: https://developer.apple.com/documentation/iokit/1557134-iopmassertioncreatewithname
func IOPMAssertionCreateWithName(arg0 corefoundation.CFStringRef, arg1 IOPMAssertionLevel, arg2 corefoundation.CFStringRef, arg3 IOPMAssertionID) int {
	if _iOPMAssertionCreateWithName == nil {
		panic("iokit: symbol IOPMAssertionCreateWithName not loaded")
	}
	return _iOPMAssertionCreateWithName(arg0, arg1, arg2, arg3)
}

var _iOPMAssertionCreateWithProperties func(arg0 corefoundation.CFDictionaryRef, arg1 IOPMAssertionID) int

// IOPMAssertionCreateWithProperties creates an IOPMAssertion with more flexibility than IOPMAssertionCreateWithDescription.
//
// See: https://developer.apple.com/documentation/iokit/1557082-iopmassertioncreatewithpropertie
func IOPMAssertionCreateWithProperties(arg0 corefoundation.CFDictionaryRef, arg1 IOPMAssertionID) int {
	if _iOPMAssertionCreateWithProperties == nil {
		panic("iokit: symbol IOPMAssertionCreateWithProperties not loaded")
	}
	return _iOPMAssertionCreateWithProperties(arg0, arg1)
}

var _iOPMAssertionDeclareUserActivity func(arg0 corefoundation.CFStringRef, arg1 IOPMUserActiveType, arg2 IOPMAssertionID) int

// IOPMAssertionDeclareUserActivity declares that the user is active on the system.
//
// See: https://developer.apple.com/documentation/iokit/1557127-iopmassertiondeclareuseractivity
func IOPMAssertionDeclareUserActivity(arg0 corefoundation.CFStringRef, arg1 IOPMUserActiveType, arg2 IOPMAssertionID) int {
	if _iOPMAssertionDeclareUserActivity == nil {
		panic("iokit: symbol IOPMAssertionDeclareUserActivity not loaded")
	}
	return _iOPMAssertionDeclareUserActivity(arg0, arg1, arg2)
}

var _iOPMAssertionRelease func(arg0 IOPMAssertionID) int

// IOPMAssertionRelease decrements the assertion's retain count.
//
// See: https://developer.apple.com/documentation/iokit/1557090-iopmassertionrelease
func IOPMAssertionRelease(arg0 IOPMAssertionID) int {
	if _iOPMAssertionRelease == nil {
		panic("iokit: symbol IOPMAssertionRelease not loaded")
	}
	return _iOPMAssertionRelease(arg0)
}

var _iOPMAssertionRetain func(arg0 IOPMAssertionID)

// IOPMAssertionRetain increments the assertion's retain count.
//
// See: https://developer.apple.com/documentation/iokit/1557071-iopmassertionretain
func IOPMAssertionRetain(arg0 IOPMAssertionID) {
	if _iOPMAssertionRetain == nil {
		panic("iokit: symbol IOPMAssertionRetain not loaded")
	}
	_iOPMAssertionRetain(arg0)
}

var _iOPMAssertionSetProperty func(arg0 IOPMAssertionID, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) int

// IOPMAssertionSetProperty sets a property in the assertion.
//
// See: https://developer.apple.com/documentation/iokit/1557107-iopmassertionsetproperty
func IOPMAssertionSetProperty(arg0 IOPMAssertionID, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFTypeRef) int {
	if _iOPMAssertionSetProperty == nil {
		panic("iokit: symbol IOPMAssertionSetProperty not loaded")
	}
	return _iOPMAssertionSetProperty(arg0, arg1, arg2)
}

var _iOPMCancelScheduledPowerEvent func(arg0 corefoundation.CFDateRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef) int

// IOPMCancelScheduledPowerEvent cancel a previously scheduled power event.
//
// See: https://developer.apple.com/documentation/iokit/1557116-iopmcancelscheduledpowerevent
func IOPMCancelScheduledPowerEvent(arg0 corefoundation.CFDateRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef) int {
	if _iOPMCancelScheduledPowerEvent == nil {
		panic("iokit: symbol IOPMCancelScheduledPowerEvent not loaded")
	}
	return _iOPMCancelScheduledPowerEvent(arg0, arg1, arg2)
}

var _iOPMCopyAssertionsByProcess func(arg0 corefoundation.CFDictionaryRef) int

// IOPMCopyAssertionsByProcess returns a dictionary listing all assertions, grouped by their owning process.
//
// See: https://developer.apple.com/documentation/iokit/1557130-iopmcopyassertionsbyprocess
func IOPMCopyAssertionsByProcess(arg0 corefoundation.CFDictionaryRef) int {
	if _iOPMCopyAssertionsByProcess == nil {
		panic("iokit: symbol IOPMCopyAssertionsByProcess not loaded")
	}
	return _iOPMCopyAssertionsByProcess(arg0)
}

var _iOPMCopyAssertionsStatus func(arg0 corefoundation.CFDictionaryRef) int

// IOPMCopyAssertionsStatus returns a list of available assertions and their system-wide levels.
//
// See: https://developer.apple.com/documentation/iokit/1557072-iopmcopyassertionsstatus
func IOPMCopyAssertionsStatus(arg0 corefoundation.CFDictionaryRef) int {
	if _iOPMCopyAssertionsStatus == nil {
		panic("iokit: symbol IOPMCopyAssertionsStatus not loaded")
	}
	return _iOPMCopyAssertionsStatus(arg0)
}

var _iOPMCopyBatteryInfo func(arg0 uint32, arg1 corefoundation.CFArrayRef) int

// IOPMCopyBatteryInfo request raw battery data from the system.
//
// See: https://developer.apple.com/documentation/iokit/1557138-iopmcopybatteryinfo
func IOPMCopyBatteryInfo(arg0 uint32, arg1 corefoundation.CFArrayRef) int {
	if _iOPMCopyBatteryInfo == nil {
		panic("iokit: symbol IOPMCopyBatteryInfo not loaded")
	}
	return _iOPMCopyBatteryInfo(arg0, arg1)
}

var _iOPMCopyCPUPowerStatus func(arg0 corefoundation.CFDictionaryRef) int

// IOPMCopyCPUPowerStatus copy status of all current CPU power levels.
//
// See: https://developer.apple.com/documentation/iokit/1557079-iopmcopycpupowerstatus
func IOPMCopyCPUPowerStatus(arg0 corefoundation.CFDictionaryRef) int {
	if _iOPMCopyCPUPowerStatus == nil {
		panic("iokit: symbol IOPMCopyCPUPowerStatus not loaded")
	}
	return _iOPMCopyCPUPowerStatus(arg0)
}

var _iOPMCopyScheduledPowerEvents func() corefoundation.CFArrayRef

// IOPMCopyScheduledPowerEvents list all scheduled system power events
//
// See: https://developer.apple.com/documentation/iokit/1557109-iopmcopyscheduledpowerevents
func IOPMCopyScheduledPowerEvents() corefoundation.CFArrayRef {
	if _iOPMCopyScheduledPowerEvents == nil {
		panic("iokit: symbol IOPMCopyScheduledPowerEvents not loaded")
	}
	return _iOPMCopyScheduledPowerEvents()
}

var _iOPMDeclareNetworkClientActivity func(arg0 corefoundation.CFStringRef, arg1 IOPMAssertionID) int

// IOPMDeclareNetworkClientActivity.
//
// See: https://developer.apple.com/documentation/iokit/1557137-iopmdeclarenetworkclientactivity
func IOPMDeclareNetworkClientActivity(arg0 corefoundation.CFStringRef, arg1 IOPMAssertionID) int {
	if _iOPMDeclareNetworkClientActivity == nil {
		panic("iokit: symbol IOPMDeclareNetworkClientActivity not loaded")
	}
	return _iOPMDeclareNetworkClientActivity(arg0, arg1)
}

var _iOPMFindPowerManagement func(arg0 uint32) uintptr

// IOPMFindPowerManagement finds the Root Power Domain IOService.
//
// See: https://developer.apple.com/documentation/iokit/1557133-iopmfindpowermanagement
func IOPMFindPowerManagement(arg0 uint32) uintptr {
	if _iOPMFindPowerManagement == nil {
		panic("iokit: symbol IOPMFindPowerManagement not loaded")
	}
	return _iOPMFindPowerManagement(arg0)
}

var _iOPMGetAggressiveness func(arg0 uintptr, arg1 uint, arg2 uint) int

// IOPMGetAggressiveness retrieves the current value of one of the aggressiveness factors in IOKit Power Management.
//
// See: https://developer.apple.com/documentation/iokit/1557117-iopmgetaggressiveness
func IOPMGetAggressiveness(arg0 uintptr, arg1 uint, arg2 uint) int {
	if _iOPMGetAggressiveness == nil {
		panic("iokit: symbol IOPMGetAggressiveness not loaded")
	}
	return _iOPMGetAggressiveness(arg0, arg1, arg2)
}

var _iOPMGetThermalWarningLevel func(arg0 uint32) int

// IOPMGetThermalWarningLevel get thermal warning level of the system.
//
// See: https://developer.apple.com/documentation/iokit/1557103-iopmgetthermalwarninglevel
func IOPMGetThermalWarningLevel(arg0 uint32) int {
	if _iOPMGetThermalWarningLevel == nil {
		panic("iokit: symbol IOPMGetThermalWarningLevel not loaded")
	}
	return _iOPMGetThermalWarningLevel(arg0)
}

var _iOPMSchedulePowerEvent func(arg0 corefoundation.CFDateRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef) int

// IOPMSchedulePowerEvent schedule the machine to wake from sleep, power on, go to sleep, or shutdown.
//
// See: https://developer.apple.com/documentation/iokit/1557076-iopmschedulepowerevent
func IOPMSchedulePowerEvent(arg0 corefoundation.CFDateRef, arg1 corefoundation.CFStringRef, arg2 corefoundation.CFStringRef) int {
	if _iOPMSchedulePowerEvent == nil {
		panic("iokit: symbol IOPMSchedulePowerEvent not loaded")
	}
	return _iOPMSchedulePowerEvent(arg0, arg1, arg2)
}

var _iOPMSetAggressiveness func(arg0 uintptr, arg1 uint, arg2 uint) int

// IOPMSetAggressiveness sets one of the aggressiveness factors in IOKit Power Management.
//
// See: https://developer.apple.com/documentation/iokit/1557098-iopmsetaggressiveness
func IOPMSetAggressiveness(arg0 uintptr, arg1 uint, arg2 uint) int {
	if _iOPMSetAggressiveness == nil {
		panic("iokit: symbol IOPMSetAggressiveness not loaded")
	}
	return _iOPMSetAggressiveness(arg0, arg1, arg2)
}

var _iOPMSleepEnabled func() bool

// IOPMSleepEnabled tells whether the system supports full sleep, or just doze
//
// See: https://developer.apple.com/documentation/iokit/1557074-iopmsleepenabled
func IOPMSleepEnabled() bool {
	if _iOPMSleepEnabled == nil {
		panic("iokit: symbol IOPMSleepEnabled not loaded")
	}
	return _iOPMSleepEnabled()
}

var _iOPMSleepSystem func(arg0 uintptr) int

// IOPMSleepSystem request that the system initiate sleep.
//
// See: https://developer.apple.com/documentation/iokit/1557121-iopmsleepsystem
func IOPMSleepSystem(arg0 uintptr) int {
	if _iOPMSleepSystem == nil {
		panic("iokit: symbol IOPMSleepSystem not loaded")
	}
	return _iOPMSleepSystem(arg0)
}

var _iOPSCopyExternalPowerAdapterDetails func() corefoundation.CFDictionaryRef

// IOPSCopyExternalPowerAdapterDetails returns a CFDictionary that describes the attached (AC) external power adapter (if any external power adapter is attached.
//
// See: https://developer.apple.com/documentation/iokit/1523866-iopscopyexternalpoweradapterdeta
func IOPSCopyExternalPowerAdapterDetails() corefoundation.CFDictionaryRef {
	if _iOPSCopyExternalPowerAdapterDetails == nil {
		panic("iokit: symbol IOPSCopyExternalPowerAdapterDetails not loaded")
	}
	return _iOPSCopyExternalPowerAdapterDetails()
}

var _iOPSCopyPowerSourcesInfo func() corefoundation.CFTypeRef

// IOPSCopyPowerSourcesInfo returns a blob of Power Source information in an opaque CFTypeRef.
//
// See: https://developer.apple.com/documentation/iokit/1523839-iopscopypowersourcesinfo
func IOPSCopyPowerSourcesInfo() corefoundation.CFTypeRef {
	if _iOPSCopyPowerSourcesInfo == nil {
		panic("iokit: symbol IOPSCopyPowerSourcesInfo not loaded")
	}
	return _iOPSCopyPowerSourcesInfo()
}

var _iOPSCopyPowerSourcesList func(arg0 corefoundation.CFTypeRef) corefoundation.CFArrayRef

// IOPSCopyPowerSourcesList returns a CFArray of Power Source handles, each of type CFTypeRef.
//
// See: https://developer.apple.com/documentation/iokit/1523856-iopscopypowersourceslist
func IOPSCopyPowerSourcesList(arg0 corefoundation.CFTypeRef) corefoundation.CFArrayRef {
	if _iOPSCopyPowerSourcesList == nil {
		panic("iokit: symbol IOPSCopyPowerSourcesList not loaded")
	}
	return _iOPSCopyPowerSourcesList(arg0)
}

var _iOPSCreateLimitedPowerNotification func(arg0 unsafe.Pointer) corefoundation.CFRunLoopSourceRef

// IOPSCreateLimitedPowerNotification.
//
// See: https://developer.apple.com/documentation/iokit/1523865-iopscreatelimitedpowernotificati
func IOPSCreateLimitedPowerNotification(arg0 unsafe.Pointer) corefoundation.CFRunLoopSourceRef {
	if _iOPSCreateLimitedPowerNotification == nil {
		panic("iokit: symbol IOPSCreateLimitedPowerNotification not loaded")
	}
	return _iOPSCreateLimitedPowerNotification(arg0)
}

var _iOPSGetBatteryWarningLevel func() IOPSLowBatteryWarningLevel

// IOPSGetBatteryWarningLevel indicates whether the system is at a low battery warning level.
//
// See: https://developer.apple.com/documentation/iokit/1523851-iopsgetbatterywarninglevel
func IOPSGetBatteryWarningLevel() IOPSLowBatteryWarningLevel {
	if _iOPSGetBatteryWarningLevel == nil {
		panic("iokit: symbol IOPSGetBatteryWarningLevel not loaded")
	}
	return _iOPSGetBatteryWarningLevel()
}

var _iOPSGetPowerSourceDescription func(arg0 corefoundation.CFTypeRef, arg1 corefoundation.CFTypeRef) corefoundation.CFDictionaryRef

// IOPSGetPowerSourceDescription returns a CFDictionary with readable information about the specific power source.
//
// See: https://developer.apple.com/documentation/iokit/1523867-iopsgetpowersourcedescription
func IOPSGetPowerSourceDescription(arg0 corefoundation.CFTypeRef, arg1 corefoundation.CFTypeRef) corefoundation.CFDictionaryRef {
	if _iOPSGetPowerSourceDescription == nil {
		panic("iokit: symbol IOPSGetPowerSourceDescription not loaded")
	}
	return _iOPSGetPowerSourceDescription(arg0, arg1)
}

var _iOPSGetProvidingPowerSourceType func(arg0 corefoundation.CFTypeRef) corefoundation.CFStringRef

// IOPSGetProvidingPowerSourceType.
//
// See: https://developer.apple.com/documentation/iokit/1523858-iopsgetprovidingpowersourcetype
func IOPSGetProvidingPowerSourceType(arg0 corefoundation.CFTypeRef) corefoundation.CFStringRef {
	if _iOPSGetProvidingPowerSourceType == nil {
		panic("iokit: symbol IOPSGetProvidingPowerSourceType not loaded")
	}
	return _iOPSGetProvidingPowerSourceType(arg0)
}

var _iOPSGetTimeRemainingEstimate func() float64

// IOPSGetTimeRemainingEstimate returns the estimated minutes remaining until all power sources (battery and/or UPS's) are empty, or returns kIOPSTimeRemainingUnlimited if attached to an unlimited power source.
//
// See: https://developer.apple.com/documentation/iokit/1523835-iopsgettimeremainingestimate
func IOPSGetTimeRemainingEstimate() float64 {
	if _iOPSGetTimeRemainingEstimate == nil {
		panic("iokit: symbol IOPSGetTimeRemainingEstimate not loaded")
	}
	return _iOPSGetTimeRemainingEstimate()
}

var _iOPSNotificationCreateRunLoopSource func(arg0 unsafe.Pointer) corefoundation.CFRunLoopSourceRef

// IOPSNotificationCreateRunLoopSource returns a CFRunLoopSourceRef that notifies the caller when power source information changes.
//
// See: https://developer.apple.com/documentation/iokit/1523868-iopsnotificationcreaterunloopsou
func IOPSNotificationCreateRunLoopSource(arg0 unsafe.Pointer) corefoundation.CFRunLoopSourceRef {
	if _iOPSNotificationCreateRunLoopSource == nil {
		panic("iokit: symbol IOPSNotificationCreateRunLoopSource not loaded")
	}
	return _iOPSNotificationCreateRunLoopSource(arg0)
}

var _iORPCMessageFromMach func(msg unsafe.Pointer, reply unsafe.Pointer) *IORPCMessage

// IORPCMessageFromMach.
//
// See: https://developer.apple.com/documentation/iokit/3325691-iorpcmessagefrommach
func IORPCMessageFromMach(msg unsafe.Pointer, reply unsafe.Pointer) *IORPCMessage {
	if _iORPCMessageFromMach == nil {
		panic("iokit: symbol IORPCMessageFromMach not loaded")
	}
	return _iORPCMessageFromMach(msg, reply)
}

var _iORegisterApp func(arg0 uintptr, arg1 IONotificationPortRef, arg2 IOServiceInterestCallback, arg3 uintptr) uintptr

// IORegisterApp connects the caller to an IOService for the purpose of receiving power state change notifications for the device controlled by the IOService.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/iokit/1557102-ioregisterapp
func IORegisterApp(arg0 uintptr, arg1 IONotificationPortRef, arg2 IOServiceInterestCallback, arg3 uintptr) uintptr {
	if _iORegisterApp == nil {
		panic("iokit: symbol IORegisterApp not loaded")
	}
	return _iORegisterApp(arg0, arg1, arg2, arg3)
}

var _iORegisterForSystemPower func(arg0 IONotificationPortRef, arg1 IOServiceInterestCallback, arg2 uintptr) uintptr

// IORegisterForSystemPower connects the caller to the Root Power Domain IOService for the purpose of receiving sleep & wake notifications for the system.
//
// See: https://developer.apple.com/documentation/iokit/1557114-ioregisterforsystempower
func IORegisterForSystemPower(arg0 IONotificationPortRef, arg1 IOServiceInterestCallback, arg2 uintptr) uintptr {
	if _iORegisterForSystemPower == nil {
		panic("iokit: symbol IORegisterForSystemPower not loaded")
	}
	return _iORegisterForSystemPower(arg0, arg1, arg2)
}

var _iORegistryCreateIterator func(mainPort uint32, plane string, options uint32, iterator *uintptr) int32

// IORegistryCreateIterator create an iterator rooted at the registry root.
//
// See: https://developer.apple.com/documentation/iokit/1514238-ioregistrycreateiterator
func IORegistryCreateIterator(mainPort uint32, plane string, options uint32, iterator *uintptr) int32 {
	if _iORegistryCreateIterator == nil {
		panic("iokit: symbol IORegistryCreateIterator not loaded")
	}
	return _iORegistryCreateIterator(mainPort, plane, options, iterator)
}

var _iORegistryEntryCopyFromPath func(mainPort uint32, path corefoundation.CFStringRef) uintptr

// IORegistryEntryCopyFromPath.
//
// See: https://developer.apple.com/documentation/iokit/1514248-ioregistryentrycopyfrompath
func IORegistryEntryCopyFromPath(mainPort uint32, path corefoundation.CFStringRef) uintptr {
	if _iORegistryEntryCopyFromPath == nil {
		panic("iokit: symbol IORegistryEntryCopyFromPath not loaded")
	}
	return _iORegistryEntryCopyFromPath(mainPort, path)
}

var _iORegistryEntryCopyPath func(entry uintptr, plane string) corefoundation.CFStringRef

// IORegistryEntryCopyPath.
//
// See: https://developer.apple.com/documentation/iokit/1514853-ioregistryentrycopypath
func IORegistryEntryCopyPath(entry uintptr, plane string) corefoundation.CFStringRef {
	if _iORegistryEntryCopyPath == nil {
		panic("iokit: symbol IORegistryEntryCopyPath not loaded")
	}
	return _iORegistryEntryCopyPath(entry, plane)
}

var _iORegistryEntryCreateCFProperties func(entry uintptr, properties *corefoundation.CFMutableDictionary, allocator corefoundation.CFAllocatorRef, options uint32) int32

// IORegistryEntryCreateCFProperties create a CF dictionary representation of a registry entry's property table.
//
// See: https://developer.apple.com/documentation/iokit/1514310-ioregistryentrycreatecfpropertie
func IORegistryEntryCreateCFProperties(entry uintptr, properties *corefoundation.CFMutableDictionary, allocator corefoundation.CFAllocatorRef, options uint32) int32 {
	if _iORegistryEntryCreateCFProperties == nil {
		panic("iokit: symbol IORegistryEntryCreateCFProperties not loaded")
	}
	return _iORegistryEntryCreateCFProperties(entry, properties, allocator, options)
}

var _iORegistryEntryCreateCFProperty func(entry uintptr, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef

// IORegistryEntryCreateCFProperty create a CF representation of a registry entry's property.
//
// See: https://developer.apple.com/documentation/iokit/1514293-ioregistryentrycreatecfproperty
func IORegistryEntryCreateCFProperty(entry uintptr, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef {
	if _iORegistryEntryCreateCFProperty == nil {
		panic("iokit: symbol IORegistryEntryCreateCFProperty not loaded")
	}
	return _iORegistryEntryCreateCFProperty(entry, key, allocator, options)
}

var _iORegistryEntryCreateIterator func(entry uintptr, plane string, options uint32, iterator *uintptr) int32

// IORegistryEntryCreateIterator create an iterator rooted at a given registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514318-ioregistryentrycreateiterator
func IORegistryEntryCreateIterator(entry uintptr, plane string, options uint32, iterator *uintptr) int32 {
	if _iORegistryEntryCreateIterator == nil {
		panic("iokit: symbol IORegistryEntryCreateIterator not loaded")
	}
	return _iORegistryEntryCreateIterator(entry, plane, options, iterator)
}

var _iORegistryEntryFromPath func(mainPort uint32, path string) uintptr

// IORegistryEntryFromPath looks up a registry entry by path.
//
// See: https://developer.apple.com/documentation/iokit/1514802-ioregistryentryfrompath
func IORegistryEntryFromPath(mainPort uint32, path string) uintptr {
	if _iORegistryEntryFromPath == nil {
		panic("iokit: symbol IORegistryEntryFromPath not loaded")
	}
	return _iORegistryEntryFromPath(mainPort, path)
}

var _iORegistryEntryGetChildEntry func(entry uintptr, plane string, child *uintptr) int32

// IORegistryEntryGetChildEntry returns the first child of a registry entry in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514496-ioregistryentrygetchildentry
func IORegistryEntryGetChildEntry(entry uintptr, plane string, child *uintptr) int32 {
	if _iORegistryEntryGetChildEntry == nil {
		panic("iokit: symbol IORegistryEntryGetChildEntry not loaded")
	}
	return _iORegistryEntryGetChildEntry(entry, plane, child)
}

var _iORegistryEntryGetChildIterator func(entry uintptr, plane string, iterator *uintptr) int32

// IORegistryEntryGetChildIterator returns an iterator over a registry entry’s child entries in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514703-ioregistryentrygetchilditerator
func IORegistryEntryGetChildIterator(entry uintptr, plane string, iterator *uintptr) int32 {
	if _iORegistryEntryGetChildIterator == nil {
		panic("iokit: symbol IORegistryEntryGetChildIterator not loaded")
	}
	return _iORegistryEntryGetChildIterator(entry, plane, iterator)
}

var _iORegistryEntryGetLocationInPlane func(entry uintptr, plane string, location string) int32

// IORegistryEntryGetLocationInPlane returns a C-string location assigned to a registry entry, in a specified plane.
//
// See: https://developer.apple.com/documentation/iokit/1514340-ioregistryentrygetlocationinplan
func IORegistryEntryGetLocationInPlane(entry uintptr, plane string, location string) int32 {
	if _iORegistryEntryGetLocationInPlane == nil {
		panic("iokit: symbol IORegistryEntryGetLocationInPlane not loaded")
	}
	return _iORegistryEntryGetLocationInPlane(entry, plane, location)
}

var _iORegistryEntryGetName func(entry uintptr, name string) int32

// IORegistryEntryGetName returns a C-string name assigned to a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514323-ioregistryentrygetname
func IORegistryEntryGetName(entry uintptr, name string) int32 {
	if _iORegistryEntryGetName == nil {
		panic("iokit: symbol IORegistryEntryGetName not loaded")
	}
	return _iORegistryEntryGetName(entry, name)
}

var _iORegistryEntryGetNameInPlane func(entry uintptr, plane string, name string) int32

// IORegistryEntryGetNameInPlane returns a C-string name assigned to a registry entry, in a specified plane.
//
// See: https://developer.apple.com/documentation/iokit/1514475-ioregistryentrygetnameinplane
func IORegistryEntryGetNameInPlane(entry uintptr, plane string, name string) int32 {
	if _iORegistryEntryGetNameInPlane == nil {
		panic("iokit: symbol IORegistryEntryGetNameInPlane not loaded")
	}
	return _iORegistryEntryGetNameInPlane(entry, plane, name)
}

var _iORegistryEntryGetParentEntry func(entry uintptr, plane string, parent *uintptr) int32

// IORegistryEntryGetParentEntry returns the first parent of a registry entry in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514454-ioregistryentrygetparententry
func IORegistryEntryGetParentEntry(entry uintptr, plane string, parent *uintptr) int32 {
	if _iORegistryEntryGetParentEntry == nil {
		panic("iokit: symbol IORegistryEntryGetParentEntry not loaded")
	}
	return _iORegistryEntryGetParentEntry(entry, plane, parent)
}

var _iORegistryEntryGetParentIterator func(entry uintptr, plane string, iterator *uintptr) int32

// IORegistryEntryGetParentIterator returns an iterator over a registry entry’s parent entries in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514366-ioregistryentrygetparentiterator
func IORegistryEntryGetParentIterator(entry uintptr, plane string, iterator *uintptr) int32 {
	if _iORegistryEntryGetParentIterator == nil {
		panic("iokit: symbol IORegistryEntryGetParentIterator not loaded")
	}
	return _iORegistryEntryGetParentIterator(entry, plane, iterator)
}

var _iORegistryEntryGetPath func(entry uintptr, plane string, path string) int32

// IORegistryEntryGetPath create a path for a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514229-ioregistryentrygetpath
func IORegistryEntryGetPath(entry uintptr, plane string, path string) int32 {
	if _iORegistryEntryGetPath == nil {
		panic("iokit: symbol IORegistryEntryGetPath not loaded")
	}
	return _iORegistryEntryGetPath(entry, plane, path)
}

var _iORegistryEntryGetProperty func(entry uintptr, propertyName string, buffer string, size *uint32) int32

// IORegistryEntryGetProperty.
//
// See: https://developer.apple.com/documentation/iokit/1514254-ioregistryentrygetproperty
func IORegistryEntryGetProperty(entry uintptr, propertyName string, buffer string, size *uint32) int32 {
	if _iORegistryEntryGetProperty == nil {
		panic("iokit: symbol IORegistryEntryGetProperty not loaded")
	}
	return _iORegistryEntryGetProperty(entry, propertyName, buffer, size)
}

var _iORegistryEntryGetRegistryEntryID func(entry uintptr, entryID *uint64) int32

// IORegistryEntryGetRegistryEntryID returns an ID for the registry entry that is global to all tasks.
//
// See: https://developer.apple.com/documentation/iokit/1514719-ioregistryentrygetregistryentryi
func IORegistryEntryGetRegistryEntryID(entry uintptr, entryID *uint64) int32 {
	if _iORegistryEntryGetRegistryEntryID == nil {
		panic("iokit: symbol IORegistryEntryGetRegistryEntryID not loaded")
	}
	return _iORegistryEntryGetRegistryEntryID(entry, entryID)
}

var _iORegistryEntryIDMatching func(entryID uint64) corefoundation.CFMutableDictionary

// IORegistryEntryIDMatching create a matching dictionary that specifies an IOService match based on a registry entry ID.
//
// See: https://developer.apple.com/documentation/iokit/1514880-ioregistryentryidmatching
func IORegistryEntryIDMatching(entryID uint64) corefoundation.CFMutableDictionary {
	if _iORegistryEntryIDMatching == nil {
		panic("iokit: symbol IORegistryEntryIDMatching not loaded")
	}
	return _iORegistryEntryIDMatching(entryID)
}

var _iORegistryEntryInPlane func(entry uintptr, plane string) bool

// IORegistryEntryInPlane determines if the registry entry is attached in a plane.
//
// See: https://developer.apple.com/documentation/iokit/1514668-ioregistryentryinplane
func IORegistryEntryInPlane(entry uintptr, plane string) bool {
	if _iORegistryEntryInPlane == nil {
		panic("iokit: symbol IORegistryEntryInPlane not loaded")
	}
	return _iORegistryEntryInPlane(entry, plane)
}

var _iORegistryEntrySearchCFProperty func(entry uintptr, plane string, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef

// IORegistryEntrySearchCFProperty create a CF representation of a registry entry's property.
//
// See: https://developer.apple.com/documentation/iokit/1514537-ioregistryentrysearchcfproperty
func IORegistryEntrySearchCFProperty(entry uintptr, plane string, key corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, options uint32) corefoundation.CFTypeRef {
	if _iORegistryEntrySearchCFProperty == nil {
		panic("iokit: symbol IORegistryEntrySearchCFProperty not loaded")
	}
	return _iORegistryEntrySearchCFProperty(entry, plane, key, allocator, options)
}

var _iORegistryEntrySetCFProperties func(entry uintptr, properties corefoundation.CFTypeRef) int32

// IORegistryEntrySetCFProperties set CF container based properties in a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514414-ioregistryentrysetcfproperties
func IORegistryEntrySetCFProperties(entry uintptr, properties corefoundation.CFTypeRef) int32 {
	if _iORegistryEntrySetCFProperties == nil {
		panic("iokit: symbol IORegistryEntrySetCFProperties not loaded")
	}
	return _iORegistryEntrySetCFProperties(entry, properties)
}

var _iORegistryEntrySetCFProperty func(entry uintptr, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32

// IORegistryEntrySetCFProperty set a CF container based property in a registry entry.
//
// See: https://developer.apple.com/documentation/iokit/1514882-ioregistryentrysetcfproperty
func IORegistryEntrySetCFProperty(entry uintptr, propertyName corefoundation.CFStringRef, property corefoundation.CFTypeRef) int32 {
	if _iORegistryEntrySetCFProperty == nil {
		panic("iokit: symbol IORegistryEntrySetCFProperty not loaded")
	}
	return _iORegistryEntrySetCFProperty(entry, propertyName, property)
}

var _iORegistryGetRootEntry func(mainPort uint32) uintptr

// IORegistryGetRootEntry return a handle to the registry root.
//
// See: https://developer.apple.com/documentation/iokit/1514878-ioregistrygetrootentry
func IORegistryGetRootEntry(mainPort uint32) uintptr {
	if _iORegistryGetRootEntry == nil {
		panic("iokit: symbol IORegistryGetRootEntry not loaded")
	}
	return _iORegistryGetRootEntry(mainPort)
}

var _iORegistryIteratorEnterEntry func(iterator uintptr) int32

// IORegistryIteratorEnterEntry recurse into the current entry in the registry iteration.
//
// See: https://developer.apple.com/documentation/iokit/1514822-ioregistryiteratorenterentry
func IORegistryIteratorEnterEntry(iterator uintptr) int32 {
	if _iORegistryIteratorEnterEntry == nil {
		panic("iokit: symbol IORegistryIteratorEnterEntry not loaded")
	}
	return _iORegistryIteratorEnterEntry(iterator)
}

var _iORegistryIteratorExitEntry func(iterator uintptr) int32

// IORegistryIteratorExitEntry exits a level of recursion, restoring the current entry.
//
// See: https://developer.apple.com/documentation/iokit/1514334-ioregistryiteratorexitentry
func IORegistryIteratorExitEntry(iterator uintptr) int32 {
	if _iORegistryIteratorExitEntry == nil {
		panic("iokit: symbol IORegistryIteratorExitEntry not loaded")
	}
	return _iORegistryIteratorExitEntry(iterator)
}

var _iOServiceAddInterestNotification func(notifyPort IONotificationPortRef, service uintptr, interestType string, callback IOServiceInterestCallback, refCon uintptr, notification *uintptr) int32

// IOServiceAddInterestNotification register for notification of state changes in an IOService.
//
// See: https://developer.apple.com/documentation/iokit/1514866-ioserviceaddinterestnotification
func IOServiceAddInterestNotification(notifyPort IONotificationPortRef, service uintptr, interestType string, callback IOServiceInterestCallback, refCon uintptr, notification *uintptr) int32 {
	if _iOServiceAddInterestNotification == nil {
		panic("iokit: symbol IOServiceAddInterestNotification not loaded")
	}
	return _iOServiceAddInterestNotification(notifyPort, service, interestType, callback, refCon, notification)
}

var _iOServiceAddMatchingNotification func(notifyPort IONotificationPortRef, notificationType string, matching corefoundation.CFDictionaryRef, callback IOServiceMatchingCallback, refCon uintptr, notification *uintptr) int32

// IOServiceAddMatchingNotification look up registered IOService objects that match a matching dictionary, and install a notification request of new IOServices that match.
//
// See: https://developer.apple.com/documentation/iokit/1514362-ioserviceaddmatchingnotification
func IOServiceAddMatchingNotification(notifyPort IONotificationPortRef, notificationType string, matching corefoundation.CFDictionaryRef, callback IOServiceMatchingCallback, refCon uintptr, notification *uintptr) int32 {
	if _iOServiceAddMatchingNotification == nil {
		panic("iokit: symbol IOServiceAddMatchingNotification not loaded")
	}
	return _iOServiceAddMatchingNotification(notifyPort, notificationType, matching, callback, refCon, notification)
}

var _iOServiceAddNotification func(mainPort uint32, notificationType string, matching corefoundation.CFDictionaryRef, wakePort uint32, reference unsafe.Pointer, notification *uintptr) int32

// IOServiceAddNotification.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/iokit/1514382-ioserviceaddnotification
func IOServiceAddNotification(mainPort uint32, notificationType string, matching corefoundation.CFDictionaryRef, wakePort uint32, reference unsafe.Pointer, notification *uintptr) int32 {
	if _iOServiceAddNotification == nil {
		panic("iokit: symbol IOServiceAddNotification not loaded")
	}
	return _iOServiceAddNotification(mainPort, notificationType, matching, wakePort, reference, notification)
}

var _iOServiceAuthorize func(service uintptr, options uint32) int32

// IOServiceAuthorize.
//
// See: https://developer.apple.com/documentation/iokit/1514533-ioserviceauthorize
func IOServiceAuthorize(service uintptr, options uint32) int32 {
	if _iOServiceAuthorize == nil {
		panic("iokit: symbol IOServiceAuthorize not loaded")
	}
	return _iOServiceAuthorize(service, options)
}

var _iOServiceClose func(connect uintptr) int32

// IOServiceClose close a connection to an IOService and destroy the connect handle.
//
// See: https://developer.apple.com/documentation/iokit/1514646-ioserviceclose
func IOServiceClose(connect uintptr) int32 {
	if _iOServiceClose == nil {
		panic("iokit: symbol IOServiceClose not loaded")
	}
	return _iOServiceClose(connect)
}

var _iOServiceGetBusyState func(service uintptr, busyState *uint32) int32

// IOServiceGetBusyState returns the busyState of an IOService.
//
// See: https://developer.apple.com/documentation/iokit/1514607-ioservicegetbusystate
func IOServiceGetBusyState(service uintptr, busyState *uint32) int32 {
	if _iOServiceGetBusyState == nil {
		panic("iokit: symbol IOServiceGetBusyState not loaded")
	}
	return _iOServiceGetBusyState(service, busyState)
}

var _iOServiceGetMatchingService func(mainPort uint32, matching corefoundation.CFDictionaryRef) uintptr

// IOServiceGetMatchingService look up a registered IOService object that matches a matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1514535-ioservicegetmatchingservice
func IOServiceGetMatchingService(mainPort uint32, matching corefoundation.CFDictionaryRef) uintptr {
	if _iOServiceGetMatchingService == nil {
		panic("iokit: symbol IOServiceGetMatchingService not loaded")
	}
	return _iOServiceGetMatchingService(mainPort, matching)
}

var _iOServiceGetMatchingServices func(mainPort uint32, matching corefoundation.CFDictionaryRef, existing *uintptr) int32

// IOServiceGetMatchingServices look up registered IOService objects that match a matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1514494-ioservicegetmatchingservices
func IOServiceGetMatchingServices(mainPort uint32, matching corefoundation.CFDictionaryRef, existing *uintptr) int32 {
	if _iOServiceGetMatchingServices == nil {
		panic("iokit: symbol IOServiceGetMatchingServices not loaded")
	}
	return _iOServiceGetMatchingServices(mainPort, matching, existing)
}

var _iOServiceMatchPropertyTable func(service uintptr, matching corefoundation.CFDictionaryRef, matches *bool) int32

// IOServiceMatchPropertyTable match an IOService objects with matching dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1514685-ioservicematchpropertytable
func IOServiceMatchPropertyTable(service uintptr, matching corefoundation.CFDictionaryRef, matches *bool) int32 {
	if _iOServiceMatchPropertyTable == nil {
		panic("iokit: symbol IOServiceMatchPropertyTable not loaded")
	}
	return _iOServiceMatchPropertyTable(service, matching, matches)
}

var _iOServiceMatching func(name string) corefoundation.CFMutableDictionary

// IOServiceMatching create a matching dictionary that specifies an IOService class match.
//
// See: https://developer.apple.com/documentation/iokit/1514687-ioservicematching
func IOServiceMatching(name string) corefoundation.CFMutableDictionary {
	if _iOServiceMatching == nil {
		panic("iokit: symbol IOServiceMatching not loaded")
	}
	return _iOServiceMatching(name)
}

var _iOServiceNameMatching func(name string) corefoundation.CFMutableDictionary

// IOServiceNameMatching create a matching dictionary that specifies an IOService name match.
//
// See: https://developer.apple.com/documentation/iokit/1514416-ioservicenamematching
func IOServiceNameMatching(name string) corefoundation.CFMutableDictionary {
	if _iOServiceNameMatching == nil {
		panic("iokit: symbol IOServiceNameMatching not loaded")
	}
	return _iOServiceNameMatching(name)
}

var _iOServiceOFPathToBSDName func(mainPort uint32, openFirmwarePath string, bsdName string) int32

// IOServiceOFPathToBSDName.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/iokit/1514661-ioserviceofpathtobsdname
func IOServiceOFPathToBSDName(mainPort uint32, openFirmwarePath string, bsdName string) int32 {
	if _iOServiceOFPathToBSDName == nil {
		panic("iokit: symbol IOServiceOFPathToBSDName not loaded")
	}
	return _iOServiceOFPathToBSDName(mainPort, openFirmwarePath, bsdName)
}

var _iOServiceOpen func(service uintptr, owningTask kernel.Task_port_t, type_ uint32, connect *uintptr) int32

// IOServiceOpen a request to create a connection to an IOService.
//
// See: https://developer.apple.com/documentation/iokit/1514515-ioserviceopen
func IOServiceOpen(service uintptr, owningTask kernel.Task_port_t, type_ uint32, connect *uintptr) int32 {
	if _iOServiceOpen == nil {
		panic("iokit: symbol IOServiceOpen not loaded")
	}
	return _iOServiceOpen(service, owningTask, type_, connect)
}

var _iOServiceOpenAsFileDescriptor func(service uintptr, oflag unsafe.Pointer) unsafe.Pointer

// IOServiceOpenAsFileDescriptor.
//
// See: https://developer.apple.com/documentation/iokit/1514879-ioserviceopenasfiledescriptor
func IOServiceOpenAsFileDescriptor(service uintptr, oflag unsafe.Pointer) unsafe.Pointer {
	if _iOServiceOpenAsFileDescriptor == nil {
		panic("iokit: symbol IOServiceOpenAsFileDescriptor not loaded")
	}
	return _iOServiceOpenAsFileDescriptor(service, oflag)
}

var _iOServiceRequestProbe func(service uintptr, options uint32) int32

// IOServiceRequestProbe a request to rescan a bus for device changes.
//
// See: https://developer.apple.com/documentation/iokit/1514364-ioservicerequestprobe
func IOServiceRequestProbe(service uintptr, options uint32) int32 {
	if _iOServiceRequestProbe == nil {
		panic("iokit: symbol IOServiceRequestProbe not loaded")
	}
	return _iOServiceRequestProbe(service, options)
}

var _iOServiceWaitQuiet func(service uintptr, waitTime *kernel.Mach_timespec_t) int32

// IOServiceWaitQuiet wait for an IOService's busyState to be zero.
//
// See: https://developer.apple.com/documentation/iokit/1514573-ioservicewaitquiet
func IOServiceWaitQuiet(service uintptr, waitTime *kernel.Mach_timespec_t) int32 {
	if _iOServiceWaitQuiet == nil {
		panic("iokit: symbol IOServiceWaitQuiet not loaded")
	}
	return _iOServiceWaitQuiet(service, waitTime)
}

var _iOURLCreateDataAndPropertiesFromResource func(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFDataRef, arg3 corefoundation.CFDictionaryRef, arg4 corefoundation.CFArrayRef, arg5 int32) bool

// IOURLCreateDataAndPropertiesFromResource.
//
// See: https://developer.apple.com/documentation/iokit/1390335-iourlcreatedataandpropertiesfrom
func IOURLCreateDataAndPropertiesFromResource(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFDataRef, arg3 corefoundation.CFDictionaryRef, arg4 corefoundation.CFArrayRef, arg5 int32) bool {
	if _iOURLCreateDataAndPropertiesFromResource == nil {
		panic("iokit: symbol IOURLCreateDataAndPropertiesFromResource not loaded")
	}
	return _iOURLCreateDataAndPropertiesFromResource(arg0, arg1, arg2, arg3, arg4, arg5)
}

var _iOURLCreatePropertyFromResource func(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFStringRef, arg3 int32) corefoundation.CFTypeRef

// IOURLCreatePropertyFromResource.
//
// See: https://developer.apple.com/documentation/iokit/1390322-iourlcreatepropertyfromresource
func IOURLCreatePropertyFromResource(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFURLRef, arg2 corefoundation.CFStringRef, arg3 int32) corefoundation.CFTypeRef {
	if _iOURLCreatePropertyFromResource == nil {
		panic("iokit: symbol IOURLCreatePropertyFromResource not loaded")
	}
	return _iOURLCreatePropertyFromResource(arg0, arg1, arg2, arg3)
}

var _iOURLWriteDataAndPropertiesToResource func(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFDataRef, arg2 corefoundation.CFDictionaryRef, arg3 int32) bool

// IOURLWriteDataAndPropertiesToResource.
//
// See: https://developer.apple.com/documentation/iokit/1390337-iourlwritedataandpropertiestores
func IOURLWriteDataAndPropertiesToResource(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFDataRef, arg2 corefoundation.CFDictionaryRef, arg3 int32) bool {
	if _iOURLWriteDataAndPropertiesToResource == nil {
		panic("iokit: symbol IOURLWriteDataAndPropertiesToResource not loaded")
	}
	return _iOURLWriteDataAndPropertiesToResource(arg0, arg1, arg2, arg3)
}

var _iOVirtualRangeMake func(arg0 IOVirtualAddress, arg1 uint) IOVirtualRange

// IOVirtualRangeMake.
//
// See: https://developer.apple.com/documentation/iokit/1555660-iovirtualrangemake
func IOVirtualRangeMake(arg0 IOVirtualAddress, arg1 uint) IOVirtualRange {
	if _iOVirtualRangeMake == nil {
		panic("iokit: symbol IOVirtualRangeMake not loaded")
	}
	return _iOVirtualRangeMake(arg0, arg1)
}

var _kextManagerCopyLoadedKextInfo func(arg0 corefoundation.CFArrayRef, arg1 corefoundation.CFArrayRef) corefoundation.CFDictionaryRef

// KextManagerCopyLoadedKextInfo returns information about loaded kexts in a dictionary.
//
// See: https://developer.apple.com/documentation/iokit/1562908-kextmanagercopyloadedkextinfo
func KextManagerCopyLoadedKextInfo(arg0 corefoundation.CFArrayRef, arg1 corefoundation.CFArrayRef) corefoundation.CFDictionaryRef {
	if _kextManagerCopyLoadedKextInfo == nil {
		panic("iokit: symbol KextManagerCopyLoadedKextInfo not loaded")
	}
	return _kextManagerCopyLoadedKextInfo(arg0, arg1)
}

var _kextManagerCreateURLForBundleIdentifier func(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFStringRef) corefoundation.CFURLRef

// KextManagerCreateURLForBundleIdentifier create a URL locating a kext with a given bundle identifier.
//
// See: https://developer.apple.com/documentation/iokit/1562905-kextmanagercreateurlforbundleide
func KextManagerCreateURLForBundleIdentifier(arg0 corefoundation.CFAllocatorRef, arg1 corefoundation.CFStringRef) corefoundation.CFURLRef {
	if _kextManagerCreateURLForBundleIdentifier == nil {
		panic("iokit: symbol KextManagerCreateURLForBundleIdentifier not loaded")
	}
	return _kextManagerCreateURLForBundleIdentifier(arg0, arg1)
}

var _kextManagerLoadKextWithIdentifier func(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFArrayRef) unsafe.Pointer

// KextManagerLoadKextWithIdentifier request the kext loading system to load a kext with a given bundle identifier.
//
// See: https://developer.apple.com/documentation/iokit/1562907-kextmanagerloadkextwithidentifie
func KextManagerLoadKextWithIdentifier(arg0 corefoundation.CFStringRef, arg1 corefoundation.CFArrayRef) unsafe.Pointer {
	if _kextManagerLoadKextWithIdentifier == nil {
		panic("iokit: symbol KextManagerLoadKextWithIdentifier not loaded")
	}
	return _kextManagerLoadKextWithIdentifier(arg0, arg1)
}

var _kextManagerLoadKextWithURL func(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFArrayRef) unsafe.Pointer

// KextManagerLoadKextWithURL request the kext loading system to load a kext with a given URL.
//
// See: https://developer.apple.com/documentation/iokit/1562906-kextmanagerloadkextwithurl
func KextManagerLoadKextWithURL(arg0 corefoundation.CFURLRef, arg1 corefoundation.CFArrayRef) unsafe.Pointer {
	if _kextManagerLoadKextWithURL == nil {
		panic("iokit: symbol KextManagerLoadKextWithURL not loaded")
	}
	return _kextManagerLoadKextWithURL(arg0, arg1)
}

var _kextManagerUnloadKextWithIdentifier func(arg0 corefoundation.CFStringRef) unsafe.Pointer

// KextManagerUnloadKextWithIdentifier request the kernel to unload a kext with a given bundle identifier.
//
// See: https://developer.apple.com/documentation/iokit/1562904-kextmanagerunloadkextwithidentif
func KextManagerUnloadKextWithIdentifier(arg0 corefoundation.CFStringRef) unsafe.Pointer {
	if _kextManagerUnloadKextWithIdentifier == nil {
		panic("iokit: symbol KextManagerUnloadKextWithIdentifier not loaded")
	}
	return _kextManagerUnloadKextWithIdentifier(arg0)
}

var _nXClickTime func(arg0 NXEventHandle) float64

// NXClickTime.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574526-nxclicktime
func NXClickTime(arg0 NXEventHandle) float64 {
	if _nXClickTime == nil {
		panic("iokit: symbol NXClickTime not loaded")
	}
	return _nXClickTime(arg0)
}

var _nXCloseEventStatus func(arg0 NXEventHandle)

// NXCloseEventStatus.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574524-nxcloseeventstatus
func NXCloseEventStatus(arg0 NXEventHandle) {
	if _nXCloseEventStatus == nil {
		panic("iokit: symbol NXCloseEventStatus not loaded")
	}
	_nXCloseEventStatus(arg0)
}

var _nXEventSystemInfo func(arg0 NXEventHandle, arg1 int8, arg2 int, arg3 uint) NXEventSystemInfoType

// NXEventSystemInfo.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574527-nxeventsysteminfo
func NXEventSystemInfo(arg0 NXEventHandle, arg1 int8, arg2 int, arg3 uint) NXEventSystemInfoType {
	if _nXEventSystemInfo == nil {
		panic("iokit: symbol NXEventSystemInfo not loaded")
	}
	return _nXEventSystemInfo(arg0, arg1, arg2, arg3)
}

var _nXGetClickSpace func(arg0 NXEventHandle, arg1 NXSize)

// NXGetClickSpace.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574518-nxgetclickspace
func NXGetClickSpace(arg0 NXEventHandle, arg1 NXSize) {
	if _nXGetClickSpace == nil {
		panic("iokit: symbol NXGetClickSpace not loaded")
	}
	_nXGetClickSpace(arg0, arg1)
}

var _nXKeyRepeatInterval func(arg0 NXEventHandle) float64

// NXKeyRepeatInterval.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574515-nxkeyrepeatinterval
func NXKeyRepeatInterval(arg0 NXEventHandle) float64 {
	if _nXKeyRepeatInterval == nil {
		panic("iokit: symbol NXKeyRepeatInterval not loaded")
	}
	return _nXKeyRepeatInterval(arg0)
}

var _nXKeyRepeatThreshold func(arg0 NXEventHandle) float64

// NXKeyRepeatThreshold.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574514-nxkeyrepeatthreshold
func NXKeyRepeatThreshold(arg0 NXEventHandle) float64 {
	if _nXKeyRepeatThreshold == nil {
		panic("iokit: symbol NXKeyRepeatThreshold not loaded")
	}
	return _nXKeyRepeatThreshold(arg0)
}

var _nXOpenEventStatus func() NXEventHandle

// NXOpenEventStatus.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574517-nxopeneventstatus
func NXOpenEventStatus() NXEventHandle {
	if _nXOpenEventStatus == nil {
		panic("iokit: symbol NXOpenEventStatus not loaded")
	}
	return _nXOpenEventStatus()
}

var _nXResetKeyboard func(arg0 NXEventHandle)

// NXResetKeyboard.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574523-nxresetkeyboard
func NXResetKeyboard(arg0 NXEventHandle) {
	if _nXResetKeyboard == nil {
		panic("iokit: symbol NXResetKeyboard not loaded")
	}
	_nXResetKeyboard(arg0)
}

var _nXResetMouse func(arg0 NXEventHandle)

// NXResetMouse.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574520-nxresetmouse
func NXResetMouse(arg0 NXEventHandle) {
	if _nXResetMouse == nil {
		panic("iokit: symbol NXResetMouse not loaded")
	}
	_nXResetMouse(arg0)
}

var _nXSetClickSpace func(arg0 NXEventHandle, arg1 NXSize)

// NXSetClickSpace.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574525-nxsetclickspace
func NXSetClickSpace(arg0 NXEventHandle, arg1 NXSize) {
	if _nXSetClickSpace == nil {
		panic("iokit: symbol NXSetClickSpace not loaded")
	}
	_nXSetClickSpace(arg0, arg1)
}

var _nXSetClickTime func(arg0 NXEventHandle, arg1 float64)

// NXSetClickTime.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574522-nxsetclicktime
func NXSetClickTime(arg0 NXEventHandle, arg1 float64) {
	if _nXSetClickTime == nil {
		panic("iokit: symbol NXSetClickTime not loaded")
	}
	_nXSetClickTime(arg0, arg1)
}

var _nXSetKeyRepeatInterval func(arg0 NXEventHandle, arg1 float64)

// NXSetKeyRepeatInterval.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574519-nxsetkeyrepeatinterval
func NXSetKeyRepeatInterval(arg0 NXEventHandle, arg1 float64) {
	if _nXSetKeyRepeatInterval == nil {
		panic("iokit: symbol NXSetKeyRepeatInterval not loaded")
	}
	_nXSetKeyRepeatInterval(arg0, arg1)
}

var _nXSetKeyRepeatThreshold func(arg0 NXEventHandle, arg1 float64)

// NXSetKeyRepeatThreshold.
//
// Deprecated: Deprecated since macOS 10.12.
//
// See: https://developer.apple.com/documentation/iokit/1574516-nxsetkeyrepeatthreshold
func NXSetKeyRepeatThreshold(arg0 NXEventHandle, arg1 float64) {
	if _nXSetKeyRepeatThreshold == nil {
		panic("iokit: symbol NXSetKeyRepeatThreshold not loaded")
	}
	_nXSetKeyRepeatThreshold(arg0, arg1)
}

var _oSGetNotificationFromMessage func(msg unsafe.Pointer, index uint32, type_ *uint32, reference unsafe.Pointer, content unsafe.Pointer, size *kernel.Vm_size_t) int32

// OSGetNotificationFromMessage.
//
// See: https://developer.apple.com/documentation/iokit/1514263-osgetnotificationfrommessage
func OSGetNotificationFromMessage(msg unsafe.Pointer, index uint32, type_ *uint32, reference unsafe.Pointer, content unsafe.Pointer, size *kernel.Vm_size_t) int32 {
	if _oSGetNotificationFromMessage == nil {
		panic("iokit: symbol OSGetNotificationFromMessage not loaded")
	}
	return _oSGetNotificationFromMessage(msg, index, type_, reference, content, size)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cDConvertLBAToMSF, frameworkHandle, "CDConvertLBAToMSF")
	registerFunc(&_cDConvertMSFToClippedLBA, frameworkHandle, "CDConvertMSFToClippedLBA")
	registerFunc(&_cDConvertMSFToLBA, frameworkHandle, "CDConvertMSFToLBA")
	registerFunc(&_cDConvertTrackNumberToMSF, frameworkHandle, "CDConvertTrackNumberToMSF")
	registerFunc(&_cDTOCGetDescriptorCount, frameworkHandle, "CDTOCGetDescriptorCount")
	registerFunc(&_iOAccelFindAccelerator, frameworkHandle, "IOAccelFindAccelerator")
	registerFunc(&_iOAllowPowerChange, frameworkHandle, "IOAllowPowerChange")
	registerFunc(&_iOBSDNameMatching, frameworkHandle, "IOBSDNameMatching")
	registerFunc(&_iOCFSerialize, frameworkHandle, "IOCFSerialize")
	registerFunc(&_iOCFUnserialize, frameworkHandle, "IOCFUnserialize")
	registerFunc(&_iOCFUnserializeBinary, frameworkHandle, "IOCFUnserializeBinary")
	registerFunc(&_iOCFUnserializeWithSize, frameworkHandle, "IOCFUnserializeWithSize")
	registerFunc(&_iOCancelPowerChange, frameworkHandle, "IOCancelPowerChange")
	registerFunc(&_iOCatalogueGetData, frameworkHandle, "IOCatalogueGetData")
	registerFunc(&_iOCatalogueModuleLoaded, frameworkHandle, "IOCatalogueModuleLoaded")
	registerFunc(&_iOCatalogueReset, frameworkHandle, "IOCatalogueReset")
	registerFunc(&_iOCatalogueSendData, frameworkHandle, "IOCatalogueSendData")
	registerFunc(&_iOCatalogueTerminate, frameworkHandle, "IOCatalogueTerminate")
	registerFunc(&_iOConnectAddClient, frameworkHandle, "IOConnectAddClient")
	registerFunc(&_iOConnectAddRef, frameworkHandle, "IOConnectAddRef")
	registerFunc(&_iOConnectCallAsyncMethod, frameworkHandle, "IOConnectCallAsyncMethod")
	registerFunc(&_iOConnectCallAsyncScalarMethod, frameworkHandle, "IOConnectCallAsyncScalarMethod")
	registerFunc(&_iOConnectCallAsyncStructMethod, frameworkHandle, "IOConnectCallAsyncStructMethod")
	registerFunc(&_iOConnectCallMethod, frameworkHandle, "IOConnectCallMethod")
	registerFunc(&_iOConnectCallScalarMethod, frameworkHandle, "IOConnectCallScalarMethod")
	registerFunc(&_iOConnectCallStructMethod, frameworkHandle, "IOConnectCallStructMethod")
	registerFunc(&_iOConnectGetService, frameworkHandle, "IOConnectGetService")
	registerFunc(&_iOConnectMapMemory, frameworkHandle, "IOConnectMapMemory")
	registerFunc(&_iOConnectMapMemory64, frameworkHandle, "IOConnectMapMemory64")
	registerFunc(&_iOConnectRelease, frameworkHandle, "IOConnectRelease")
	registerFunc(&_iOConnectSetCFProperties, frameworkHandle, "IOConnectSetCFProperties")
	registerFunc(&_iOConnectSetCFProperty, frameworkHandle, "IOConnectSetCFProperty")
	registerFunc(&_iOConnectSetNotificationPort, frameworkHandle, "IOConnectSetNotificationPort")
	registerFunc(&_iOConnectTrap0, frameworkHandle, "IOConnectTrap0")
	registerFunc(&_iOConnectTrap1, frameworkHandle, "IOConnectTrap1")
	registerFunc(&_iOConnectTrap2, frameworkHandle, "IOConnectTrap2")
	registerFunc(&_iOConnectTrap3, frameworkHandle, "IOConnectTrap3")
	registerFunc(&_iOConnectTrap4, frameworkHandle, "IOConnectTrap4")
	registerFunc(&_iOConnectTrap5, frameworkHandle, "IOConnectTrap5")
	registerFunc(&_iOConnectTrap6, frameworkHandle, "IOConnectTrap6")
	registerFunc(&_iOConnectUnmapMemory, frameworkHandle, "IOConnectUnmapMemory")
	registerFunc(&_iOConnectUnmapMemory64, frameworkHandle, "IOConnectUnmapMemory64")
	registerFunc(&_iOCopySystemLoadAdvisoryDetailed, frameworkHandle, "IOCopySystemLoadAdvisoryDetailed")
	registerFunc(&_iOCreatePlugInInterfaceForService, frameworkHandle, "IOCreatePlugInInterfaceForService")
	registerFunc(&_iOCreateReceivePort, frameworkHandle, "IOCreateReceivePort")
	registerFunc(&_iODataQueueAllocateNotificationPort, frameworkHandle, "IODataQueueAllocateNotificationPort")
	registerFunc(&_iODataQueueDataAvailable, frameworkHandle, "IODataQueueDataAvailable")
	registerFunc(&_iODataQueueDequeue, frameworkHandle, "IODataQueueDequeue")
	registerFunc(&_iODataQueueEnqueue, frameworkHandle, "IODataQueueEnqueue")
	registerFunc(&_iODataQueuePeek, frameworkHandle, "IODataQueuePeek")
	registerFunc(&_iODataQueueSetNotificationPort, frameworkHandle, "IODataQueueSetNotificationPort")
	registerFunc(&_iODataQueueWaitForAvailableData, frameworkHandle, "IODataQueueWaitForAvailableData")
	registerFunc(&_iODeregisterApp, frameworkHandle, "IODeregisterApp")
	registerFunc(&_iODeregisterForSystemPower, frameworkHandle, "IODeregisterForSystemPower")
	registerFunc(&_iODestroyPlugInInterface, frameworkHandle, "IODestroyPlugInInterface")
	registerFunc(&_iODispatchCalloutFromMessage, frameworkHandle, "IODispatchCalloutFromMessage")
	registerFunc(&_iODisplayCommitParameters, frameworkHandle, "IODisplayCommitParameters")
	registerFunc(&_iODisplayCopyFloatParameters, frameworkHandle, "IODisplayCopyFloatParameters")
	registerFunc(&_iODisplayCopyParameters, frameworkHandle, "IODisplayCopyParameters")
	registerFunc(&_iODisplayCreateInfoDictionary, frameworkHandle, "IODisplayCreateInfoDictionary")
	registerFunc(&_iODisplayForFramebuffer, frameworkHandle, "IODisplayForFramebuffer")
	registerFunc(&_iODisplayGetFloatParameter, frameworkHandle, "IODisplayGetFloatParameter")
	registerFunc(&_iODisplayGetIntegerRangeParameter, frameworkHandle, "IODisplayGetIntegerRangeParameter")
	registerFunc(&_iODisplayMatchDictionaries, frameworkHandle, "IODisplayMatchDictionaries")
	registerFunc(&_iODisplaySetFloatParameter, frameworkHandle, "IODisplaySetFloatParameter")
	registerFunc(&_iODisplaySetIntegerParameter, frameworkHandle, "IODisplaySetIntegerParameter")
	registerFunc(&_iODisplaySetParameters, frameworkHandle, "IODisplaySetParameters")
	registerFunc(&_iOFBCopyI2CInterfaceForBus, frameworkHandle, "IOFBCopyI2CInterfaceForBus")
	registerFunc(&_iOFBGetI2CInterfaceCount, frameworkHandle, "IOFBGetI2CInterfaceCount")
	registerFunc(&_iOFramebufferOpen, frameworkHandle, "IOFramebufferOpen")
	registerFunc(&_iOGetSystemLoadAdvisory, frameworkHandle, "IOGetSystemLoadAdvisory")
	registerFunc(&_iOHIDCheckAccess, frameworkHandle, "IOHIDCheckAccess")
	registerFunc(&_iOHIDCopyCFTypeParameter, frameworkHandle, "IOHIDCopyCFTypeParameter")
	registerFunc(&_iOHIDCreateSharedMemory, frameworkHandle, "IOHIDCreateSharedMemory")
	registerFunc(&_iOHIDDeviceActivate, frameworkHandle, "IOHIDDeviceActivate")
	registerFunc(&_iOHIDDeviceCancel, frameworkHandle, "IOHIDDeviceCancel")
	registerFunc(&_iOHIDDeviceClose, frameworkHandle, "IOHIDDeviceClose")
	registerFunc(&_iOHIDDeviceConformsTo, frameworkHandle, "IOHIDDeviceConformsTo")
	registerFunc(&_iOHIDDeviceCopyMatchingElements, frameworkHandle, "IOHIDDeviceCopyMatchingElements")
	registerFunc(&_iOHIDDeviceCopyValueMultiple, frameworkHandle, "IOHIDDeviceCopyValueMultiple")
	registerFunc(&_iOHIDDeviceCopyValueMultipleWithCallback, frameworkHandle, "IOHIDDeviceCopyValueMultipleWithCallback")
	registerFunc(&_iOHIDDeviceCreate, frameworkHandle, "IOHIDDeviceCreate")
	registerFunc(&_iOHIDDeviceGetProperty, frameworkHandle, "IOHIDDeviceGetProperty")
	registerFunc(&_iOHIDDeviceGetReport, frameworkHandle, "IOHIDDeviceGetReport")
	registerFunc(&_iOHIDDeviceGetReportWithCallback, frameworkHandle, "IOHIDDeviceGetReportWithCallback")
	registerFunc(&_iOHIDDeviceGetService, frameworkHandle, "IOHIDDeviceGetService")
	registerFunc(&_iOHIDDeviceGetTypeID, frameworkHandle, "IOHIDDeviceGetTypeID")
	registerFunc(&_iOHIDDeviceGetValue, frameworkHandle, "IOHIDDeviceGetValue")
	registerFunc(&_iOHIDDeviceGetValueWithCallback, frameworkHandle, "IOHIDDeviceGetValueWithCallback")
	registerFunc(&_iOHIDDeviceGetValueWithOptions, frameworkHandle, "IOHIDDeviceGetValueWithOptions")
	registerFunc(&_iOHIDDeviceOpen, frameworkHandle, "IOHIDDeviceOpen")
	registerFunc(&_iOHIDDeviceRegisterInputReportCallback, frameworkHandle, "IOHIDDeviceRegisterInputReportCallback")
	registerFunc(&_iOHIDDeviceRegisterInputReportWithTimeStampCallback, frameworkHandle, "IOHIDDeviceRegisterInputReportWithTimeStampCallback")
	registerFunc(&_iOHIDDeviceRegisterInputValueCallback, frameworkHandle, "IOHIDDeviceRegisterInputValueCallback")
	registerFunc(&_iOHIDDeviceRegisterRemovalCallback, frameworkHandle, "IOHIDDeviceRegisterRemovalCallback")
	registerFunc(&_iOHIDDeviceScheduleWithRunLoop, frameworkHandle, "IOHIDDeviceScheduleWithRunLoop")
	registerFunc(&_iOHIDDeviceSetCancelHandler, frameworkHandle, "IOHIDDeviceSetCancelHandler")
	registerFunc(&_iOHIDDeviceSetDispatchQueue, frameworkHandle, "IOHIDDeviceSetDispatchQueue")
	registerFunc(&_iOHIDDeviceSetInputValueMatching, frameworkHandle, "IOHIDDeviceSetInputValueMatching")
	registerFunc(&_iOHIDDeviceSetInputValueMatchingMultiple, frameworkHandle, "IOHIDDeviceSetInputValueMatchingMultiple")
	registerFunc(&_iOHIDDeviceSetProperty, frameworkHandle, "IOHIDDeviceSetProperty")
	registerFunc(&_iOHIDDeviceSetReport, frameworkHandle, "IOHIDDeviceSetReport")
	registerFunc(&_iOHIDDeviceSetReportWithCallback, frameworkHandle, "IOHIDDeviceSetReportWithCallback")
	registerFunc(&_iOHIDDeviceSetValue, frameworkHandle, "IOHIDDeviceSetValue")
	registerFunc(&_iOHIDDeviceSetValueMultiple, frameworkHandle, "IOHIDDeviceSetValueMultiple")
	registerFunc(&_iOHIDDeviceSetValueMultipleWithCallback, frameworkHandle, "IOHIDDeviceSetValueMultipleWithCallback")
	registerFunc(&_iOHIDDeviceSetValueWithCallback, frameworkHandle, "IOHIDDeviceSetValueWithCallback")
	registerFunc(&_iOHIDDeviceUnscheduleFromRunLoop, frameworkHandle, "IOHIDDeviceUnscheduleFromRunLoop")
	registerFunc(&_iOHIDElementAttach, frameworkHandle, "IOHIDElementAttach")
	registerFunc(&_iOHIDElementCopyAttached, frameworkHandle, "IOHIDElementCopyAttached")
	registerFunc(&_iOHIDElementCreateWithDictionary, frameworkHandle, "IOHIDElementCreateWithDictionary")
	registerFunc(&_iOHIDElementDetach, frameworkHandle, "IOHIDElementDetach")
	registerFunc(&_iOHIDElementGetChildren, frameworkHandle, "IOHIDElementGetChildren")
	registerFunc(&_iOHIDElementGetCollectionType, frameworkHandle, "IOHIDElementGetCollectionType")
	registerFunc(&_iOHIDElementGetCookie, frameworkHandle, "IOHIDElementGetCookie")
	registerFunc(&_iOHIDElementGetDevice, frameworkHandle, "IOHIDElementGetDevice")
	registerFunc(&_iOHIDElementGetLogicalMax, frameworkHandle, "IOHIDElementGetLogicalMax")
	registerFunc(&_iOHIDElementGetLogicalMin, frameworkHandle, "IOHIDElementGetLogicalMin")
	registerFunc(&_iOHIDElementGetName, frameworkHandle, "IOHIDElementGetName")
	registerFunc(&_iOHIDElementGetParent, frameworkHandle, "IOHIDElementGetParent")
	registerFunc(&_iOHIDElementGetPhysicalMax, frameworkHandle, "IOHIDElementGetPhysicalMax")
	registerFunc(&_iOHIDElementGetPhysicalMin, frameworkHandle, "IOHIDElementGetPhysicalMin")
	registerFunc(&_iOHIDElementGetProperty, frameworkHandle, "IOHIDElementGetProperty")
	registerFunc(&_iOHIDElementGetReportCount, frameworkHandle, "IOHIDElementGetReportCount")
	registerFunc(&_iOHIDElementGetReportID, frameworkHandle, "IOHIDElementGetReportID")
	registerFunc(&_iOHIDElementGetReportSize, frameworkHandle, "IOHIDElementGetReportSize")
	registerFunc(&_iOHIDElementGetType, frameworkHandle, "IOHIDElementGetType")
	registerFunc(&_iOHIDElementGetTypeID, frameworkHandle, "IOHIDElementGetTypeID")
	registerFunc(&_iOHIDElementGetUnit, frameworkHandle, "IOHIDElementGetUnit")
	registerFunc(&_iOHIDElementGetUnitExponent, frameworkHandle, "IOHIDElementGetUnitExponent")
	registerFunc(&_iOHIDElementGetUsage, frameworkHandle, "IOHIDElementGetUsage")
	registerFunc(&_iOHIDElementGetUsagePage, frameworkHandle, "IOHIDElementGetUsagePage")
	registerFunc(&_iOHIDElementHasNullState, frameworkHandle, "IOHIDElementHasNullState")
	registerFunc(&_iOHIDElementHasPreferredState, frameworkHandle, "IOHIDElementHasPreferredState")
	registerFunc(&_iOHIDElementIsArray, frameworkHandle, "IOHIDElementIsArray")
	registerFunc(&_iOHIDElementIsNonLinear, frameworkHandle, "IOHIDElementIsNonLinear")
	registerFunc(&_iOHIDElementIsRelative, frameworkHandle, "IOHIDElementIsRelative")
	registerFunc(&_iOHIDElementIsVirtual, frameworkHandle, "IOHIDElementIsVirtual")
	registerFunc(&_iOHIDElementIsWrapping, frameworkHandle, "IOHIDElementIsWrapping")
	registerFunc(&_iOHIDElementSetProperty, frameworkHandle, "IOHIDElementSetProperty")
	registerFunc(&_iOHIDEventSystemClientCopyProperty, frameworkHandle, "IOHIDEventSystemClientCopyProperty")
	registerFunc(&_iOHIDEventSystemClientCopyServices, frameworkHandle, "IOHIDEventSystemClientCopyServices")
	registerFunc(&_iOHIDEventSystemClientCreateSimpleClient, frameworkHandle, "IOHIDEventSystemClientCreateSimpleClient")
	registerFunc(&_iOHIDEventSystemClientGetTypeID, frameworkHandle, "IOHIDEventSystemClientGetTypeID")
	registerFunc(&_iOHIDEventSystemClientSetProperty, frameworkHandle, "IOHIDEventSystemClientSetProperty")
	registerFunc(&_iOHIDGetAccelerationWithKey, frameworkHandle, "IOHIDGetAccelerationWithKey")
	registerFunc(&_iOHIDGetActivityState, frameworkHandle, "IOHIDGetActivityState")
	registerFunc(&_iOHIDGetButtonEventNum, frameworkHandle, "IOHIDGetButtonEventNum")
	registerFunc(&_iOHIDGetModifierLockState, frameworkHandle, "IOHIDGetModifierLockState")
	registerFunc(&_iOHIDGetMouseAcceleration, frameworkHandle, "IOHIDGetMouseAcceleration")
	registerFunc(&_iOHIDGetMouseButtonMode, frameworkHandle, "IOHIDGetMouseButtonMode")
	registerFunc(&_iOHIDGetParameter, frameworkHandle, "IOHIDGetParameter")
	registerFunc(&_iOHIDGetScrollAcceleration, frameworkHandle, "IOHIDGetScrollAcceleration")
	registerFunc(&_iOHIDGetStateForSelector, frameworkHandle, "IOHIDGetStateForSelector")
	registerFunc(&_iOHIDManagerActivate, frameworkHandle, "IOHIDManagerActivate")
	registerFunc(&_iOHIDManagerCancel, frameworkHandle, "IOHIDManagerCancel")
	registerFunc(&_iOHIDManagerClose, frameworkHandle, "IOHIDManagerClose")
	registerFunc(&_iOHIDManagerCopyDevices, frameworkHandle, "IOHIDManagerCopyDevices")
	registerFunc(&_iOHIDManagerCreate, frameworkHandle, "IOHIDManagerCreate")
	registerFunc(&_iOHIDManagerGetProperty, frameworkHandle, "IOHIDManagerGetProperty")
	registerFunc(&_iOHIDManagerGetTypeID, frameworkHandle, "IOHIDManagerGetTypeID")
	registerFunc(&_iOHIDManagerOpen, frameworkHandle, "IOHIDManagerOpen")
	registerFunc(&_iOHIDManagerRegisterDeviceMatchingCallback, frameworkHandle, "IOHIDManagerRegisterDeviceMatchingCallback")
	registerFunc(&_iOHIDManagerRegisterDeviceRemovalCallback, frameworkHandle, "IOHIDManagerRegisterDeviceRemovalCallback")
	registerFunc(&_iOHIDManagerRegisterInputReportCallback, frameworkHandle, "IOHIDManagerRegisterInputReportCallback")
	registerFunc(&_iOHIDManagerRegisterInputReportWithTimeStampCallback, frameworkHandle, "IOHIDManagerRegisterInputReportWithTimeStampCallback")
	registerFunc(&_iOHIDManagerRegisterInputValueCallback, frameworkHandle, "IOHIDManagerRegisterInputValueCallback")
	registerFunc(&_iOHIDManagerSaveToPropertyDomain, frameworkHandle, "IOHIDManagerSaveToPropertyDomain")
	registerFunc(&_iOHIDManagerScheduleWithRunLoop, frameworkHandle, "IOHIDManagerScheduleWithRunLoop")
	registerFunc(&_iOHIDManagerSetCancelHandler, frameworkHandle, "IOHIDManagerSetCancelHandler")
	registerFunc(&_iOHIDManagerSetDeviceMatching, frameworkHandle, "IOHIDManagerSetDeviceMatching")
	registerFunc(&_iOHIDManagerSetDeviceMatchingMultiple, frameworkHandle, "IOHIDManagerSetDeviceMatchingMultiple")
	registerFunc(&_iOHIDManagerSetDispatchQueue, frameworkHandle, "IOHIDManagerSetDispatchQueue")
	registerFunc(&_iOHIDManagerSetInputValueMatching, frameworkHandle, "IOHIDManagerSetInputValueMatching")
	registerFunc(&_iOHIDManagerSetInputValueMatchingMultiple, frameworkHandle, "IOHIDManagerSetInputValueMatchingMultiple")
	registerFunc(&_iOHIDManagerSetProperty, frameworkHandle, "IOHIDManagerSetProperty")
	registerFunc(&_iOHIDManagerUnscheduleFromRunLoop, frameworkHandle, "IOHIDManagerUnscheduleFromRunLoop")
	registerFunc(&_iOHIDPostEvent, frameworkHandle, "IOHIDPostEvent")
	registerFunc(&_iOHIDQueueActivate, frameworkHandle, "IOHIDQueueActivate")
	registerFunc(&_iOHIDQueueAddElement, frameworkHandle, "IOHIDQueueAddElement")
	registerFunc(&_iOHIDQueueCancel, frameworkHandle, "IOHIDQueueCancel")
	registerFunc(&_iOHIDQueueContainsElement, frameworkHandle, "IOHIDQueueContainsElement")
	registerFunc(&_iOHIDQueueCopyNextValue, frameworkHandle, "IOHIDQueueCopyNextValue")
	registerFunc(&_iOHIDQueueCopyNextValueWithTimeout, frameworkHandle, "IOHIDQueueCopyNextValueWithTimeout")
	registerFunc(&_iOHIDQueueCreate, frameworkHandle, "IOHIDQueueCreate")
	registerFunc(&_iOHIDQueueGetDepth, frameworkHandle, "IOHIDQueueGetDepth")
	registerFunc(&_iOHIDQueueGetDevice, frameworkHandle, "IOHIDQueueGetDevice")
	registerFunc(&_iOHIDQueueGetTypeID, frameworkHandle, "IOHIDQueueGetTypeID")
	registerFunc(&_iOHIDQueueRegisterValueAvailableCallback, frameworkHandle, "IOHIDQueueRegisterValueAvailableCallback")
	registerFunc(&_iOHIDQueueRemoveElement, frameworkHandle, "IOHIDQueueRemoveElement")
	registerFunc(&_iOHIDQueueScheduleWithRunLoop, frameworkHandle, "IOHIDQueueScheduleWithRunLoop")
	registerFunc(&_iOHIDQueueSetCancelHandler, frameworkHandle, "IOHIDQueueSetCancelHandler")
	registerFunc(&_iOHIDQueueSetDepth, frameworkHandle, "IOHIDQueueSetDepth")
	registerFunc(&_iOHIDQueueSetDispatchQueue, frameworkHandle, "IOHIDQueueSetDispatchQueue")
	registerFunc(&_iOHIDQueueStart, frameworkHandle, "IOHIDQueueStart")
	registerFunc(&_iOHIDQueueStop, frameworkHandle, "IOHIDQueueStop")
	registerFunc(&_iOHIDQueueUnscheduleFromRunLoop, frameworkHandle, "IOHIDQueueUnscheduleFromRunLoop")
	registerFunc(&_iOHIDRegisterVirtualDisplay, frameworkHandle, "IOHIDRegisterVirtualDisplay")
	registerFunc(&_iOHIDRequestAccess, frameworkHandle, "IOHIDRequestAccess")
	registerFunc(&_iOHIDServiceClientConformsTo, frameworkHandle, "IOHIDServiceClientConformsTo")
	registerFunc(&_iOHIDServiceClientCopyProperty, frameworkHandle, "IOHIDServiceClientCopyProperty")
	registerFunc(&_iOHIDServiceClientGetRegistryID, frameworkHandle, "IOHIDServiceClientGetRegistryID")
	registerFunc(&_iOHIDServiceClientGetTypeID, frameworkHandle, "IOHIDServiceClientGetTypeID")
	registerFunc(&_iOHIDServiceClientSetProperty, frameworkHandle, "IOHIDServiceClientSetProperty")
	registerFunc(&_iOHIDSetAccelerationWithKey, frameworkHandle, "IOHIDSetAccelerationWithKey")
	registerFunc(&_iOHIDSetCFTypeParameter, frameworkHandle, "IOHIDSetCFTypeParameter")
	registerFunc(&_iOHIDSetCursorEnable, frameworkHandle, "IOHIDSetCursorEnable")
	registerFunc(&_iOHIDSetEventsEnable, frameworkHandle, "IOHIDSetEventsEnable")
	registerFunc(&_iOHIDSetModifierLockState, frameworkHandle, "IOHIDSetModifierLockState")
	registerFunc(&_iOHIDSetMouseAcceleration, frameworkHandle, "IOHIDSetMouseAcceleration")
	registerFunc(&_iOHIDSetMouseButtonMode, frameworkHandle, "IOHIDSetMouseButtonMode")
	registerFunc(&_iOHIDSetMouseLocation, frameworkHandle, "IOHIDSetMouseLocation")
	registerFunc(&_iOHIDSetParameter, frameworkHandle, "IOHIDSetParameter")
	registerFunc(&_iOHIDSetScrollAcceleration, frameworkHandle, "IOHIDSetScrollAcceleration")
	registerFunc(&_iOHIDSetStateForSelector, frameworkHandle, "IOHIDSetStateForSelector")
	registerFunc(&_iOHIDSetVirtualDisplayBounds, frameworkHandle, "IOHIDSetVirtualDisplayBounds")
	registerFunc(&_iOHIDTransactionAddElement, frameworkHandle, "IOHIDTransactionAddElement")
	registerFunc(&_iOHIDTransactionClear, frameworkHandle, "IOHIDTransactionClear")
	registerFunc(&_iOHIDTransactionCommit, frameworkHandle, "IOHIDTransactionCommit")
	registerFunc(&_iOHIDTransactionCommitWithCallback, frameworkHandle, "IOHIDTransactionCommitWithCallback")
	registerFunc(&_iOHIDTransactionContainsElement, frameworkHandle, "IOHIDTransactionContainsElement")
	registerFunc(&_iOHIDTransactionCreate, frameworkHandle, "IOHIDTransactionCreate")
	registerFunc(&_iOHIDTransactionGetDevice, frameworkHandle, "IOHIDTransactionGetDevice")
	registerFunc(&_iOHIDTransactionGetDirection, frameworkHandle, "IOHIDTransactionGetDirection")
	registerFunc(&_iOHIDTransactionGetTypeID, frameworkHandle, "IOHIDTransactionGetTypeID")
	registerFunc(&_iOHIDTransactionGetValue, frameworkHandle, "IOHIDTransactionGetValue")
	registerFunc(&_iOHIDTransactionRemoveElement, frameworkHandle, "IOHIDTransactionRemoveElement")
	registerFunc(&_iOHIDTransactionScheduleWithRunLoop, frameworkHandle, "IOHIDTransactionScheduleWithRunLoop")
	registerFunc(&_iOHIDTransactionSetDirection, frameworkHandle, "IOHIDTransactionSetDirection")
	registerFunc(&_iOHIDTransactionSetValue, frameworkHandle, "IOHIDTransactionSetValue")
	registerFunc(&_iOHIDTransactionUnscheduleFromRunLoop, frameworkHandle, "IOHIDTransactionUnscheduleFromRunLoop")
	registerFunc(&_iOHIDUnregisterVirtualDisplay, frameworkHandle, "IOHIDUnregisterVirtualDisplay")
	registerFunc(&_iOHIDUserDeviceActivate, frameworkHandle, "IOHIDUserDeviceActivate")
	registerFunc(&_iOHIDUserDeviceCancel, frameworkHandle, "IOHIDUserDeviceCancel")
	registerFunc(&_iOHIDUserDeviceCopyProperty, frameworkHandle, "IOHIDUserDeviceCopyProperty")
	registerFunc(&_iOHIDUserDeviceCreateWithProperties, frameworkHandle, "IOHIDUserDeviceCreateWithProperties")
	registerFunc(&_iOHIDUserDeviceGetTypeID, frameworkHandle, "IOHIDUserDeviceGetTypeID")
	registerFunc(&_iOHIDUserDeviceHandleReportWithTimeStamp, frameworkHandle, "IOHIDUserDeviceHandleReportWithTimeStamp")
	registerFunc(&_iOHIDUserDeviceRegisterGetReportBlock, frameworkHandle, "IOHIDUserDeviceRegisterGetReportBlock")
	registerFunc(&_iOHIDUserDeviceRegisterSetReportBlock, frameworkHandle, "IOHIDUserDeviceRegisterSetReportBlock")
	registerFunc(&_iOHIDUserDeviceSetCancelHandler, frameworkHandle, "IOHIDUserDeviceSetCancelHandler")
	registerFunc(&_iOHIDUserDeviceSetDispatchQueue, frameworkHandle, "IOHIDUserDeviceSetDispatchQueue")
	registerFunc(&_iOHIDUserDeviceSetProperty, frameworkHandle, "IOHIDUserDeviceSetProperty")
	registerFunc(&_iOHIDValueCreateWithBytes, frameworkHandle, "IOHIDValueCreateWithBytes")
	registerFunc(&_iOHIDValueCreateWithBytesNoCopy, frameworkHandle, "IOHIDValueCreateWithBytesNoCopy")
	registerFunc(&_iOHIDValueCreateWithIntegerValue, frameworkHandle, "IOHIDValueCreateWithIntegerValue")
	registerFunc(&_iOHIDValueGetBytePtr, frameworkHandle, "IOHIDValueGetBytePtr")
	registerFunc(&_iOHIDValueGetElement, frameworkHandle, "IOHIDValueGetElement")
	registerFunc(&_iOHIDValueGetIntegerValue, frameworkHandle, "IOHIDValueGetIntegerValue")
	registerFunc(&_iOHIDValueGetLength, frameworkHandle, "IOHIDValueGetLength")
	registerFunc(&_iOHIDValueGetScaledValue, frameworkHandle, "IOHIDValueGetScaledValue")
	registerFunc(&_iOHIDValueGetTimeStamp, frameworkHandle, "IOHIDValueGetTimeStamp")
	registerFunc(&_iOHIDValueGetTypeID, frameworkHandle, "IOHIDValueGetTypeID")
	registerFunc(&_iOI2CCopyInterfaceForID, frameworkHandle, "IOI2CCopyInterfaceForID")
	registerFunc(&_iOI2CInterfaceClose, frameworkHandle, "IOI2CInterfaceClose")
	registerFunc(&_iOI2CInterfaceOpen, frameworkHandle, "IOI2CInterfaceOpen")
	registerFunc(&_iOI2CSendRequest, frameworkHandle, "IOI2CSendRequest")
	registerFunc(&_iOIteratorIsValid, frameworkHandle, "IOIteratorIsValid")
	registerFunc(&_iOIteratorNext, frameworkHandle, "IOIteratorNext")
	registerFunc(&_iOIteratorReset, frameworkHandle, "IOIteratorReset")
	registerFunc(&_iOKitGetBusyState, frameworkHandle, "IOKitGetBusyState")
	registerFunc(&_iOKitWaitQuiet, frameworkHandle, "IOKitWaitQuiet")
	registerFunc(&_iOMainPort, frameworkHandle, "IOMainPort")
	registerFunc(&_iOMasterPort, frameworkHandle, "IOMasterPort")
	registerFunc(&_iONetworkClose, frameworkHandle, "IONetworkClose")
	registerFunc(&_iONetworkGetDataCapacity, frameworkHandle, "IONetworkGetDataCapacity")
	registerFunc(&_iONetworkGetDataHandle, frameworkHandle, "IONetworkGetDataHandle")
	registerFunc(&_iONetworkGetPacketFiltersMask, frameworkHandle, "IONetworkGetPacketFiltersMask")
	registerFunc(&_iONetworkOpen, frameworkHandle, "IONetworkOpen")
	registerFunc(&_iONetworkReadData, frameworkHandle, "IONetworkReadData")
	registerFunc(&_iONetworkResetData, frameworkHandle, "IONetworkResetData")
	registerFunc(&_iONetworkSetPacketFiltersMask, frameworkHandle, "IONetworkSetPacketFiltersMask")
	registerFunc(&_iONetworkWriteData, frameworkHandle, "IONetworkWriteData")
	registerFunc(&_iONotificationPortCreate, frameworkHandle, "IONotificationPortCreate")
	registerFunc(&_iONotificationPortDestroy, frameworkHandle, "IONotificationPortDestroy")
	registerFunc(&_iONotificationPortGetMachPort, frameworkHandle, "IONotificationPortGetMachPort")
	registerFunc(&_iONotificationPortGetRunLoopSource, frameworkHandle, "IONotificationPortGetRunLoopSource")
	registerFunc(&_iONotificationPortSetDispatchQueue, frameworkHandle, "IONotificationPortSetDispatchQueue")
	registerFunc(&_iONotificationPortSetImportanceReceiver, frameworkHandle, "IONotificationPortSetImportanceReceiver")
	registerFunc(&_iOObjectConformsTo, frameworkHandle, "IOObjectConformsTo")
	registerFunc(&_iOObjectCopyBundleIdentifierForClass, frameworkHandle, "IOObjectCopyBundleIdentifierForClass")
	registerFunc(&_iOObjectCopyClass, frameworkHandle, "IOObjectCopyClass")
	registerFunc(&_iOObjectCopySuperclassForClass, frameworkHandle, "IOObjectCopySuperclassForClass")
	registerFunc(&_iOObjectGetClass, frameworkHandle, "IOObjectGetClass")
	registerFunc(&_iOObjectGetKernelRetainCount, frameworkHandle, "IOObjectGetKernelRetainCount")
	registerFunc(&_iOObjectGetRetainCount, frameworkHandle, "IOObjectGetRetainCount")
	registerFunc(&_iOObjectGetUserRetainCount, frameworkHandle, "IOObjectGetUserRetainCount")
	registerFunc(&_iOObjectIsEqualTo, frameworkHandle, "IOObjectIsEqualTo")
	registerFunc(&_iOObjectRelease, frameworkHandle, "IOObjectRelease")
	registerFunc(&_iOObjectRetain, frameworkHandle, "IOObjectRetain")
	registerFunc(&_iOOpenFirmwarePathMatching, frameworkHandle, "IOOpenFirmwarePathMatching")
	registerFunc(&_iOPMAssertionCopyProperties, frameworkHandle, "IOPMAssertionCopyProperties")
	registerFunc(&_iOPMAssertionCreate, frameworkHandle, "IOPMAssertionCreate")
	registerFunc(&_iOPMAssertionCreateWithDescription, frameworkHandle, "IOPMAssertionCreateWithDescription")
	registerFunc(&_iOPMAssertionCreateWithName, frameworkHandle, "IOPMAssertionCreateWithName")
	registerFunc(&_iOPMAssertionCreateWithProperties, frameworkHandle, "IOPMAssertionCreateWithProperties")
	registerFunc(&_iOPMAssertionDeclareUserActivity, frameworkHandle, "IOPMAssertionDeclareUserActivity")
	registerFunc(&_iOPMAssertionRelease, frameworkHandle, "IOPMAssertionRelease")
	registerFunc(&_iOPMAssertionRetain, frameworkHandle, "IOPMAssertionRetain")
	registerFunc(&_iOPMAssertionSetProperty, frameworkHandle, "IOPMAssertionSetProperty")
	registerFunc(&_iOPMCancelScheduledPowerEvent, frameworkHandle, "IOPMCancelScheduledPowerEvent")
	registerFunc(&_iOPMCopyAssertionsByProcess, frameworkHandle, "IOPMCopyAssertionsByProcess")
	registerFunc(&_iOPMCopyAssertionsStatus, frameworkHandle, "IOPMCopyAssertionsStatus")
	registerFunc(&_iOPMCopyBatteryInfo, frameworkHandle, "IOPMCopyBatteryInfo")
	registerFunc(&_iOPMCopyCPUPowerStatus, frameworkHandle, "IOPMCopyCPUPowerStatus")
	registerFunc(&_iOPMCopyScheduledPowerEvents, frameworkHandle, "IOPMCopyScheduledPowerEvents")
	registerFunc(&_iOPMDeclareNetworkClientActivity, frameworkHandle, "IOPMDeclareNetworkClientActivity")
	registerFunc(&_iOPMFindPowerManagement, frameworkHandle, "IOPMFindPowerManagement")
	registerFunc(&_iOPMGetAggressiveness, frameworkHandle, "IOPMGetAggressiveness")
	registerFunc(&_iOPMGetThermalWarningLevel, frameworkHandle, "IOPMGetThermalWarningLevel")
	registerFunc(&_iOPMSchedulePowerEvent, frameworkHandle, "IOPMSchedulePowerEvent")
	registerFunc(&_iOPMSetAggressiveness, frameworkHandle, "IOPMSetAggressiveness")
	registerFunc(&_iOPMSleepEnabled, frameworkHandle, "IOPMSleepEnabled")
	registerFunc(&_iOPMSleepSystem, frameworkHandle, "IOPMSleepSystem")
	registerFunc(&_iOPSCopyExternalPowerAdapterDetails, frameworkHandle, "IOPSCopyExternalPowerAdapterDetails")
	registerFunc(&_iOPSCopyPowerSourcesInfo, frameworkHandle, "IOPSCopyPowerSourcesInfo")
	registerFunc(&_iOPSCopyPowerSourcesList, frameworkHandle, "IOPSCopyPowerSourcesList")
	registerFunc(&_iOPSCreateLimitedPowerNotification, frameworkHandle, "IOPSCreateLimitedPowerNotification")
	registerFunc(&_iOPSGetBatteryWarningLevel, frameworkHandle, "IOPSGetBatteryWarningLevel")
	registerFunc(&_iOPSGetPowerSourceDescription, frameworkHandle, "IOPSGetPowerSourceDescription")
	registerFunc(&_iOPSGetProvidingPowerSourceType, frameworkHandle, "IOPSGetProvidingPowerSourceType")
	registerFunc(&_iOPSGetTimeRemainingEstimate, frameworkHandle, "IOPSGetTimeRemainingEstimate")
	registerFunc(&_iOPSNotificationCreateRunLoopSource, frameworkHandle, "IOPSNotificationCreateRunLoopSource")
	registerFunc(&_iORPCMessageFromMach, frameworkHandle, "IORPCMessageFromMach")
	registerFunc(&_iORegisterApp, frameworkHandle, "IORegisterApp")
	registerFunc(&_iORegisterForSystemPower, frameworkHandle, "IORegisterForSystemPower")
	registerFunc(&_iORegistryCreateIterator, frameworkHandle, "IORegistryCreateIterator")
	registerFunc(&_iORegistryEntryCopyFromPath, frameworkHandle, "IORegistryEntryCopyFromPath")
	registerFunc(&_iORegistryEntryCopyPath, frameworkHandle, "IORegistryEntryCopyPath")
	registerFunc(&_iORegistryEntryCreateCFProperties, frameworkHandle, "IORegistryEntryCreateCFProperties")
	registerFunc(&_iORegistryEntryCreateCFProperty, frameworkHandle, "IORegistryEntryCreateCFProperty")
	registerFunc(&_iORegistryEntryCreateIterator, frameworkHandle, "IORegistryEntryCreateIterator")
	registerFunc(&_iORegistryEntryFromPath, frameworkHandle, "IORegistryEntryFromPath")
	registerFunc(&_iORegistryEntryGetChildEntry, frameworkHandle, "IORegistryEntryGetChildEntry")
	registerFunc(&_iORegistryEntryGetChildIterator, frameworkHandle, "IORegistryEntryGetChildIterator")
	registerFunc(&_iORegistryEntryGetLocationInPlane, frameworkHandle, "IORegistryEntryGetLocationInPlane")
	registerFunc(&_iORegistryEntryGetName, frameworkHandle, "IORegistryEntryGetName")
	registerFunc(&_iORegistryEntryGetNameInPlane, frameworkHandle, "IORegistryEntryGetNameInPlane")
	registerFunc(&_iORegistryEntryGetParentEntry, frameworkHandle, "IORegistryEntryGetParentEntry")
	registerFunc(&_iORegistryEntryGetParentIterator, frameworkHandle, "IORegistryEntryGetParentIterator")
	registerFunc(&_iORegistryEntryGetPath, frameworkHandle, "IORegistryEntryGetPath")
	registerFunc(&_iORegistryEntryGetProperty, frameworkHandle, "IORegistryEntryGetProperty")
	registerFunc(&_iORegistryEntryGetRegistryEntryID, frameworkHandle, "IORegistryEntryGetRegistryEntryID")
	registerFunc(&_iORegistryEntryIDMatching, frameworkHandle, "IORegistryEntryIDMatching")
	registerFunc(&_iORegistryEntryInPlane, frameworkHandle, "IORegistryEntryInPlane")
	registerFunc(&_iORegistryEntrySearchCFProperty, frameworkHandle, "IORegistryEntrySearchCFProperty")
	registerFunc(&_iORegistryEntrySetCFProperties, frameworkHandle, "IORegistryEntrySetCFProperties")
	registerFunc(&_iORegistryEntrySetCFProperty, frameworkHandle, "IORegistryEntrySetCFProperty")
	registerFunc(&_iORegistryGetRootEntry, frameworkHandle, "IORegistryGetRootEntry")
	registerFunc(&_iORegistryIteratorEnterEntry, frameworkHandle, "IORegistryIteratorEnterEntry")
	registerFunc(&_iORegistryIteratorExitEntry, frameworkHandle, "IORegistryIteratorExitEntry")
	registerFunc(&_iOServiceAddInterestNotification, frameworkHandle, "IOServiceAddInterestNotification")
	registerFunc(&_iOServiceAddMatchingNotification, frameworkHandle, "IOServiceAddMatchingNotification")
	registerFunc(&_iOServiceAddNotification, frameworkHandle, "IOServiceAddNotification")
	registerFunc(&_iOServiceAuthorize, frameworkHandle, "IOServiceAuthorize")
	registerFunc(&_iOServiceClose, frameworkHandle, "IOServiceClose")
	registerFunc(&_iOServiceGetBusyState, frameworkHandle, "IOServiceGetBusyState")
	registerFunc(&_iOServiceGetMatchingService, frameworkHandle, "IOServiceGetMatchingService")
	registerFunc(&_iOServiceGetMatchingServices, frameworkHandle, "IOServiceGetMatchingServices")
	registerFunc(&_iOServiceMatchPropertyTable, frameworkHandle, "IOServiceMatchPropertyTable")
	registerFunc(&_iOServiceMatching, frameworkHandle, "IOServiceMatching")
	registerFunc(&_iOServiceNameMatching, frameworkHandle, "IOServiceNameMatching")
	registerFunc(&_iOServiceOFPathToBSDName, frameworkHandle, "IOServiceOFPathToBSDName")
	registerFunc(&_iOServiceOpen, frameworkHandle, "IOServiceOpen")
	registerFunc(&_iOServiceOpenAsFileDescriptor, frameworkHandle, "IOServiceOpenAsFileDescriptor")
	registerFunc(&_iOServiceRequestProbe, frameworkHandle, "IOServiceRequestProbe")
	registerFunc(&_iOServiceWaitQuiet, frameworkHandle, "IOServiceWaitQuiet")
	registerFunc(&_iOURLCreateDataAndPropertiesFromResource, frameworkHandle, "IOURLCreateDataAndPropertiesFromResource")
	registerFunc(&_iOURLCreatePropertyFromResource, frameworkHandle, "IOURLCreatePropertyFromResource")
	registerFunc(&_iOURLWriteDataAndPropertiesToResource, frameworkHandle, "IOURLWriteDataAndPropertiesToResource")
	registerFunc(&_iOVirtualRangeMake, frameworkHandle, "IOVirtualRangeMake")
	registerFunc(&_kextManagerCopyLoadedKextInfo, frameworkHandle, "KextManagerCopyLoadedKextInfo")
	registerFunc(&_kextManagerCreateURLForBundleIdentifier, frameworkHandle, "KextManagerCreateURLForBundleIdentifier")
	registerFunc(&_kextManagerLoadKextWithIdentifier, frameworkHandle, "KextManagerLoadKextWithIdentifier")
	registerFunc(&_kextManagerLoadKextWithURL, frameworkHandle, "KextManagerLoadKextWithURL")
	registerFunc(&_kextManagerUnloadKextWithIdentifier, frameworkHandle, "KextManagerUnloadKextWithIdentifier")
	registerFunc(&_nXClickTime, frameworkHandle, "NXClickTime")
	registerFunc(&_nXCloseEventStatus, frameworkHandle, "NXCloseEventStatus")
	registerFunc(&_nXEventSystemInfo, frameworkHandle, "NXEventSystemInfo")
	registerFunc(&_nXGetClickSpace, frameworkHandle, "NXGetClickSpace")
	registerFunc(&_nXKeyRepeatInterval, frameworkHandle, "NXKeyRepeatInterval")
	registerFunc(&_nXKeyRepeatThreshold, frameworkHandle, "NXKeyRepeatThreshold")
	registerFunc(&_nXOpenEventStatus, frameworkHandle, "NXOpenEventStatus")
	registerFunc(&_nXResetKeyboard, frameworkHandle, "NXResetKeyboard")
	registerFunc(&_nXResetMouse, frameworkHandle, "NXResetMouse")
	registerFunc(&_nXSetClickSpace, frameworkHandle, "NXSetClickSpace")
	registerFunc(&_nXSetClickTime, frameworkHandle, "NXSetClickTime")
	registerFunc(&_nXSetKeyRepeatInterval, frameworkHandle, "NXSetKeyRepeatInterval")
	registerFunc(&_nXSetKeyRepeatThreshold, frameworkHandle, "NXSetKeyRepeatThreshold")
	registerFunc(&_oSGetNotificationFromMessage, frameworkHandle, "OSGetNotificationFromMessage")
}
