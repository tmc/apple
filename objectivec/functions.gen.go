// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ObjectiveC: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: ObjectiveC: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ObjectiveC: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _nXCompareHashTables func(table1 *NXHashTable, table2 *NXHashTable) bool

// NXCompareHashTables.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCompareHashTables
func NXCompareHashTables(table1 *NXHashTable, table2 *NXHashTable) bool {
	if _nXCompareHashTables == nil {
		panic("ObjectiveC: symbol NXCompareHashTables not loaded")
	}
	return _nXCompareHashTables(table1, table2)
}

var _nXCopyHashTable func(table *NXHashTable) *NXHashTable

// NXCopyHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCopyHashTable
func NXCopyHashTable(table *NXHashTable) *NXHashTable {
	if _nXCopyHashTable == nil {
		panic("ObjectiveC: symbol NXCopyHashTable not loaded")
	}
	return _nXCopyHashTable(table)
}

var _nXCountHashTable func(table *NXHashTable) uint

// NXCountHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCountHashTable
func NXCountHashTable(table *NXHashTable) uint {
	if _nXCountHashTable == nil {
		panic("ObjectiveC: symbol NXCountHashTable not loaded")
	}
	return _nXCountHashTable(table)
}

var _nXCreateHashTable func(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer) *NXHashTable

// NXCreateHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTable
func NXCreateHashTable(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer) *NXHashTable {
	if _nXCreateHashTable == nil {
		panic("ObjectiveC: symbol NXCreateHashTable not loaded")
	}
	return _nXCreateHashTable(prototype, capacity, info)
}

var _nXCreateHashTableFromZone func(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer, zone unsafe.Pointer) *NXHashTable

// NXCreateHashTableFromZone.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTableFromZone
func NXCreateHashTableFromZone(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer, zone unsafe.Pointer) *NXHashTable {
	if _nXCreateHashTableFromZone == nil {
		panic("ObjectiveC: symbol NXCreateHashTableFromZone not loaded")
	}
	return _nXCreateHashTableFromZone(prototype, capacity, info, zone)
}

var _nXEmptyHashTable func(table *NXHashTable)

// NXEmptyHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXEmptyHashTable
func NXEmptyHashTable(table *NXHashTable) {
	if _nXEmptyHashTable == nil {
		panic("ObjectiveC: symbol NXEmptyHashTable not loaded")
	}
	_nXEmptyHashTable(table)
}

var _nXFreeHashTable func(table *NXHashTable)

// NXFreeHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXFreeHashTable
func NXFreeHashTable(table *NXHashTable) {
	if _nXFreeHashTable == nil {
		panic("ObjectiveC: symbol NXFreeHashTable not loaded")
	}
	_nXFreeHashTable(table)
}

var _nXHashGet func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer

// NXHashGet.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashGet
func NXHashGet(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	if _nXHashGet == nil {
		panic("ObjectiveC: symbol NXHashGet not loaded")
	}
	return _nXHashGet(table, data)
}

var _nXHashInsert func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer

// NXHashInsert.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashInsert
func NXHashInsert(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	if _nXHashInsert == nil {
		panic("ObjectiveC: symbol NXHashInsert not loaded")
	}
	return _nXHashInsert(table, data)
}

var _nXHashInsertIfAbsent func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer

// NXHashInsertIfAbsent.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashInsertIfAbsent
func NXHashInsertIfAbsent(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	if _nXHashInsertIfAbsent == nil {
		panic("ObjectiveC: symbol NXHashInsertIfAbsent not loaded")
	}
	return _nXHashInsertIfAbsent(table, data)
}

var _nXHashMember func(table *NXHashTable, data unsafe.Pointer) int

// NXHashMember.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashMember
func NXHashMember(table *NXHashTable, data unsafe.Pointer) int {
	if _nXHashMember == nil {
		panic("ObjectiveC: symbol NXHashMember not loaded")
	}
	return _nXHashMember(table, data)
}

var _nXHashRemove func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer

// NXHashRemove.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashRemove
func NXHashRemove(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	if _nXHashRemove == nil {
		panic("ObjectiveC: symbol NXHashRemove not loaded")
	}
	return _nXHashRemove(table, data)
}

var _nXInitHashState func(table *NXHashTable) NXHashState

// NXInitHashState.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXInitHashState
func NXInitHashState(table *NXHashTable) NXHashState {
	if _nXInitHashState == nil {
		panic("ObjectiveC: symbol NXInitHashState not loaded")
	}
	return _nXInitHashState(table)
}

var _nXNextHashState func(table *NXHashTable, state *NXHashState, data unsafe.Pointer) int

// NXNextHashState.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXNextHashState
func NXNextHashState(table *NXHashTable, state *NXHashState, data unsafe.Pointer) int {
	if _nXNextHashState == nil {
		panic("ObjectiveC: symbol NXNextHashState not loaded")
	}
	return _nXNextHashState(table, state, data)
}

var _nXNoEffectFree func(info unsafe.Pointer, data unsafe.Pointer)

// NXNoEffectFree.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXNoEffectFree
func NXNoEffectFree(info unsafe.Pointer, data unsafe.Pointer) {
	if _nXNoEffectFree == nil {
		panic("ObjectiveC: symbol NXNoEffectFree not loaded")
	}
	_nXNoEffectFree(info, data)
}

var _nXPtrHash func(info unsafe.Pointer, data unsafe.Pointer) uintptr

// NXPtrHash.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXPtrHash
func NXPtrHash(info unsafe.Pointer, data unsafe.Pointer) uintptr {
	if _nXPtrHash == nil {
		panic("ObjectiveC: symbol NXPtrHash not loaded")
	}
	return _nXPtrHash(info, data)
}

var _nXPtrIsEqual func(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int

// NXPtrIsEqual.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXPtrIsEqual
func NXPtrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	if _nXPtrIsEqual == nil {
		panic("ObjectiveC: symbol NXPtrIsEqual not loaded")
	}
	return _nXPtrIsEqual(info, data1, data2)
}

var _nXReallyFree func(info unsafe.Pointer, data unsafe.Pointer)

// NXReallyFree.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXReallyFree
func NXReallyFree(info unsafe.Pointer, data unsafe.Pointer) {
	if _nXReallyFree == nil {
		panic("ObjectiveC: symbol NXReallyFree not loaded")
	}
	_nXReallyFree(info, data)
}

var _nXResetHashTable func(table *NXHashTable)

// NXResetHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXResetHashTable
func NXResetHashTable(table *NXHashTable) {
	if _nXResetHashTable == nil {
		panic("ObjectiveC: symbol NXResetHashTable not loaded")
	}
	_nXResetHashTable(table)
}

var _nXStrHash func(info unsafe.Pointer, data unsafe.Pointer) uintptr

// NXStrHash.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXStrHash
func NXStrHash(info unsafe.Pointer, data unsafe.Pointer) uintptr {
	if _nXStrHash == nil {
		panic("ObjectiveC: symbol NXStrHash not loaded")
	}
	return _nXStrHash(info, data)
}

var _nXStrIsEqual func(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int

// NXStrIsEqual.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXStrIsEqual
func NXStrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	if _nXStrIsEqual == nil {
		panic("ObjectiveC: symbol NXStrIsEqual not loaded")
	}
	return _nXStrIsEqual(info, data1, data2)
}

var _class_addIvar func(cls objc.Class, name string, size uintptr, alignment uint8, types string) bool

// Class_addIvar adds a new instance variable to a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addIvar(_:_:_:_:_:)
func Class_addIvar(cls objc.Class, name string, size uintptr, alignment uint8, types string) bool {
	if _class_addIvar == nil {
		panic("ObjectiveC: symbol class_addIvar not loaded")
	}
	return _class_addIvar(cls, name, size, alignment, types)
}

var _class_addMethod func(cls objc.Class, name SEL, imp IMP, types string) bool

// Class_addMethod adds a new method to a class with a given name and implementation.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addMethod(_:_:_:_:)
func Class_addMethod(cls objc.Class, name SEL, imp IMP, types string) bool {
	if _class_addMethod == nil {
		panic("ObjectiveC: symbol class_addMethod not loaded")
	}
	return _class_addMethod(cls, name, imp, types)
}

var _class_addProperty func(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) bool

// Class_addProperty adds a property to a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addProperty(_:_:_:_:)
func Class_addProperty(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) bool {
	if _class_addProperty == nil {
		panic("ObjectiveC: symbol class_addProperty not loaded")
	}
	return _class_addProperty(cls, name, attributes, attributeCount)
}

var _class_addProtocol func(cls objc.Class, protocol_ **Protocol) bool

// Class_addProtocol adds a protocol to a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addProtocol(_:_:)
func Class_addProtocol(cls objc.Class, protocol_ **Protocol) bool {
	if _class_addProtocol == nil {
		panic("ObjectiveC: symbol class_addProtocol not loaded")
	}
	return _class_addProtocol(cls, protocol_)
}

var _class_conformsToProtocol func(cls objc.Class, protocol_ **Protocol) bool

// Class_conformsToProtocol returns a Boolean value that indicates whether a class conforms to a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_conformsToProtocol(_:_:)
func Class_conformsToProtocol(cls objc.Class, protocol_ **Protocol) bool {
	if _class_conformsToProtocol == nil {
		panic("ObjectiveC: symbol class_conformsToProtocol not loaded")
	}
	return _class_conformsToProtocol(cls, protocol_)
}

var _class_copyIvarList func(cls objc.Class, outCount *uint) Ivar

// Class_copyIvarList describes the instance variables declared by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyIvarList(_:_:)
func Class_copyIvarList(cls objc.Class, outCount *uint) Ivar {
	if _class_copyIvarList == nil {
		panic("ObjectiveC: symbol class_copyIvarList not loaded")
	}
	return _class_copyIvarList(cls, outCount)
}

var _class_copyMethodList func(cls objc.Class, outCount *uint) Method

// Class_copyMethodList describes the instance methods implemented by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyMethodList(_:_:)
func Class_copyMethodList(cls objc.Class, outCount *uint) Method {
	if _class_copyMethodList == nil {
		panic("ObjectiveC: symbol class_copyMethodList not loaded")
	}
	return _class_copyMethodList(cls, outCount)
}

var _class_copyPropertyList func(cls objc.Class, outCount *uint) *Objc_property_t

// Class_copyPropertyList describes the properties declared by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyPropertyList(_:_:)
func Class_copyPropertyList(cls objc.Class, outCount *uint) *Objc_property_t {
	if _class_copyPropertyList == nil {
		panic("ObjectiveC: symbol class_copyPropertyList not loaded")
	}
	return _class_copyPropertyList(cls, outCount)
}

var _class_copyProtocolList func(cls objc.Class, outCount *uint) **Protocol

// Class_copyProtocolList describes the protocols adopted by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyProtocolList(_:_:)
func Class_copyProtocolList(cls objc.Class, outCount *uint) **Protocol {
	if _class_copyProtocolList == nil {
		panic("ObjectiveC: symbol class_copyProtocolList not loaded")
	}
	return _class_copyProtocolList(cls, outCount)
}

var _class_createInstance func(cls objc.Class, extraBytes uintptr) Object

// Class_createInstance creates an instance of a class, allocating memory for the class in the default malloc memory zone.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_createInstance(_:_:)
func Class_createInstance(cls objc.Class, extraBytes uintptr) Object {
	if _class_createInstance == nil {
		panic("ObjectiveC: symbol class_createInstance not loaded")
	}
	return _class_createInstance(cls, extraBytes)
}

var _class_createInstanceFromZone func(arg0 objc.Class, idxIvars uintptr, zone unsafe.Pointer) Object

// Class_createInstanceFromZone.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_createInstanceFromZone
func Class_createInstanceFromZone(arg0 objc.Class, idxIvars uintptr, zone unsafe.Pointer) Object {
	if _class_createInstanceFromZone == nil {
		panic("ObjectiveC: symbol class_createInstanceFromZone not loaded")
	}
	return _class_createInstanceFromZone(arg0, idxIvars, zone)
}

var _class_getClassMethod func(cls objc.Class, name SEL) Method

// Class_getClassMethod returns a pointer to the data structure describing a given class method for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getClassMethod(_:_:)
func Class_getClassMethod(cls objc.Class, name SEL) Method {
	if _class_getClassMethod == nil {
		panic("ObjectiveC: symbol class_getClassMethod not loaded")
	}
	return _class_getClassMethod(cls, name)
}

var _class_getClassVariable func(cls objc.Class, name string) Ivar

// Class_getClassVariable returns the [Ivar] for a specified class variable of a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getClassVariable(_:_:)
func Class_getClassVariable(cls objc.Class, name string) Ivar {
	if _class_getClassVariable == nil {
		panic("ObjectiveC: symbol class_getClassVariable not loaded")
	}
	return _class_getClassVariable(cls, name)
}

var _class_getImageName func(cls objc.Class) *byte

// Class_getImageName returns the name of the dynamic library a class originated from.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getImageName(_:)
func Class_getImageName(cls objc.Class) *byte {
	if _class_getImageName == nil {
		panic("ObjectiveC: symbol class_getImageName not loaded")
	}
	return _class_getImageName(cls)
}

var _class_getInstanceMethod func(cls objc.Class, name SEL) Method

// Class_getInstanceMethod returns a specified instance method for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceMethod(_:_:)
func Class_getInstanceMethod(cls objc.Class, name SEL) Method {
	if _class_getInstanceMethod == nil {
		panic("ObjectiveC: symbol class_getInstanceMethod not loaded")
	}
	return _class_getInstanceMethod(cls, name)
}

var _class_getInstanceSize func(cls objc.Class) uintptr

// Class_getInstanceSize returns the size of instances of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceSize(_:)
func Class_getInstanceSize(cls objc.Class) uintptr {
	if _class_getInstanceSize == nil {
		panic("ObjectiveC: symbol class_getInstanceSize not loaded")
	}
	return _class_getInstanceSize(cls)
}

var _class_getInstanceVariable func(cls objc.Class, name string) Ivar

// Class_getInstanceVariable returns the [Ivar] for a specified instance variable of a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceVariable(_:_:)
func Class_getInstanceVariable(cls objc.Class, name string) Ivar {
	if _class_getInstanceVariable == nil {
		panic("ObjectiveC: symbol class_getInstanceVariable not loaded")
	}
	return _class_getInstanceVariable(cls, name)
}

var _class_getIvarLayout func(cls objc.Class) *uint8

// Class_getIvarLayout returns a description of the [Ivar] layout for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getIvarLayout(_:)
func Class_getIvarLayout(cls objc.Class) *uint8 {
	if _class_getIvarLayout == nil {
		panic("ObjectiveC: symbol class_getIvarLayout not loaded")
	}
	return _class_getIvarLayout(cls)
}

var _class_getMethodImplementation func(cls objc.Class, name SEL) IMP

// Class_getMethodImplementation returns the function pointer that would be called if a particular message were sent to an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getMethodImplementation(_:_:)
func Class_getMethodImplementation(cls objc.Class, name SEL) IMP {
	if _class_getMethodImplementation == nil {
		panic("ObjectiveC: symbol class_getMethodImplementation not loaded")
	}
	return _class_getMethodImplementation(cls, name)
}

var _class_getName func(cls objc.Class) *byte

// Class_getName returns the name of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getName(_:)
func Class_getName(cls objc.Class) *byte {
	if _class_getName == nil {
		panic("ObjectiveC: symbol class_getName not loaded")
	}
	return _class_getName(cls)
}

var _class_getProperty func(cls objc.Class, name string) unsafe.Pointer

// Class_getProperty returns a property with a given name of a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getProperty(_:_:)
func Class_getProperty(cls objc.Class, name string) unsafe.Pointer {
	if _class_getProperty == nil {
		panic("ObjectiveC: symbol class_getProperty not loaded")
	}
	return _class_getProperty(cls, name)
}

var _class_getSuperclass func(cls objc.Class) objc.Class

// Class_getSuperclass returns the superclass of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getSuperclass(_:)
func Class_getSuperclass(cls objc.Class) objc.Class {
	if _class_getSuperclass == nil {
		panic("ObjectiveC: symbol class_getSuperclass not loaded")
	}
	return _class_getSuperclass(cls)
}

var _class_getVersion func(cls objc.Class) int

// Class_getVersion returns the version number of a class definition.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getVersion(_:)
func Class_getVersion(cls objc.Class) int {
	if _class_getVersion == nil {
		panic("ObjectiveC: symbol class_getVersion not loaded")
	}
	return _class_getVersion(cls)
}

var _class_getWeakIvarLayout func(cls objc.Class) *uint8

// Class_getWeakIvarLayout returns a description of the layout of weak [Ivar]s for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getWeakIvarLayout(_:)
func Class_getWeakIvarLayout(cls objc.Class) *uint8 {
	if _class_getWeakIvarLayout == nil {
		panic("ObjectiveC: symbol class_getWeakIvarLayout not loaded")
	}
	return _class_getWeakIvarLayout(cls)
}

var _class_isMetaClass func(cls objc.Class) bool

// Class_isMetaClass returns a Boolean value that indicates whether a class object is a metaclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_isMetaClass(_:)
func Class_isMetaClass(cls objc.Class) bool {
	if _class_isMetaClass == nil {
		panic("ObjectiveC: symbol class_isMetaClass not loaded")
	}
	return _class_isMetaClass(cls)
}

var _class_replaceMethod func(cls objc.Class, name SEL, imp IMP, types string) IMP

// Class_replaceMethod replaces the implementation of a method for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_replaceMethod(_:_:_:_:)
func Class_replaceMethod(cls objc.Class, name SEL, imp IMP, types string) IMP {
	if _class_replaceMethod == nil {
		panic("ObjectiveC: symbol class_replaceMethod not loaded")
	}
	return _class_replaceMethod(cls, name, imp, types)
}

var _class_replaceProperty func(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint)

// Class_replaceProperty replace a property of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_replaceProperty(_:_:_:_:)
func Class_replaceProperty(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) {
	if _class_replaceProperty == nil {
		panic("ObjectiveC: symbol class_replaceProperty not loaded")
	}
	_class_replaceProperty(cls, name, attributes, attributeCount)
}

var _class_respondsToSelector func(cls objc.Class, sel SEL) bool

// Class_respondsToSelector returns a Boolean value that indicates whether instances of a class respond to a particular selector.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_respondsToSelector(_:_:)
func Class_respondsToSelector(cls objc.Class, sel SEL) bool {
	if _class_respondsToSelector == nil {
		panic("ObjectiveC: symbol class_respondsToSelector not loaded")
	}
	return _class_respondsToSelector(cls, sel)
}

var _class_setIvarLayout func(cls objc.Class, layout *uint8)

// Class_setIvarLayout sets the [Ivar] layout for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_setIvarLayout(_:_:)
func Class_setIvarLayout(cls objc.Class, layout *uint8) {
	if _class_setIvarLayout == nil {
		panic("ObjectiveC: symbol class_setIvarLayout not loaded")
	}
	_class_setIvarLayout(cls, layout)
}

var _class_setVersion func(cls objc.Class, version int)

// Class_setVersion sets the version number of a class definition.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_setVersion(_:_:)
func Class_setVersion(cls objc.Class, version int) {
	if _class_setVersion == nil {
		panic("ObjectiveC: symbol class_setVersion not loaded")
	}
	_class_setVersion(cls, version)
}

var _class_setWeakIvarLayout func(cls objc.Class, layout *uint8)

// Class_setWeakIvarLayout sets the layout for weak [Ivar]s for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_setWeakIvarLayout(_:_:)
func Class_setWeakIvarLayout(cls objc.Class, layout *uint8) {
	if _class_setWeakIvarLayout == nil {
		panic("ObjectiveC: symbol class_setWeakIvarLayout not loaded")
	}
	_class_setWeakIvarLayout(cls, layout)
}

var _imp_getBlock func(anImp IMP) Object

// Imp_getBlock returns the block associated with an [IMP] that was created using imp_implementationWithBlock(_:).
//
// See: https://developer.apple.com/documentation/ObjectiveC/imp_getBlock(_:)
func Imp_getBlock(anImp IMP) Object {
	if _imp_getBlock == nil {
		panic("ObjectiveC: symbol imp_getBlock not loaded")
	}
	return _imp_getBlock(anImp)
}

var _imp_implementationWithBlock func(block Object) IMP

// Imp_implementationWithBlock creates a pointer to a function that calls the specified block when the method is called.
//
// See: https://developer.apple.com/documentation/ObjectiveC/imp_implementationWithBlock(_:)
func Imp_implementationWithBlock(block Object) IMP {
	if _imp_implementationWithBlock == nil {
		panic("ObjectiveC: symbol imp_implementationWithBlock not loaded")
	}
	return _imp_implementationWithBlock(block)
}

var _imp_removeBlock func(anImp IMP) bool

// Imp_removeBlock disassociates a block from an [IMP] that was created using imp_implementationWithBlock(_:), and releases the copy of the block that was created.
//
// See: https://developer.apple.com/documentation/ObjectiveC/imp_removeBlock(_:)
func Imp_removeBlock(anImp IMP) bool {
	if _imp_removeBlock == nil {
		panic("ObjectiveC: symbol imp_removeBlock not loaded")
	}
	return _imp_removeBlock(anImp)
}

var _ivar_getName func(v Ivar) *byte

// Ivar_getName returns the name of an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/ivar_getName(_:)
func Ivar_getName(v Ivar) *byte {
	if _ivar_getName == nil {
		panic("ObjectiveC: symbol ivar_getName not loaded")
	}
	return _ivar_getName(v)
}

var _ivar_getOffset func(v Ivar) int

// Ivar_getOffset returns the offset of an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/ivar_getOffset(_:)
func Ivar_getOffset(v Ivar) int {
	if _ivar_getOffset == nil {
		panic("ObjectiveC: symbol ivar_getOffset not loaded")
	}
	return _ivar_getOffset(v)
}

var _ivar_getTypeEncoding func(v Ivar) *byte

// Ivar_getTypeEncoding returns the type string of an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/ivar_getTypeEncoding(_:)
func Ivar_getTypeEncoding(v Ivar) *byte {
	if _ivar_getTypeEncoding == nil {
		panic("ObjectiveC: symbol ivar_getTypeEncoding not loaded")
	}
	return _ivar_getTypeEncoding(v)
}

var _method_copyArgumentType func(m Method, index uint) *byte

// Method_copyArgumentType returns a string describing a single parameter type of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_copyArgumentType(_:_:)
func Method_copyArgumentType(m Method, index uint) *byte {
	if _method_copyArgumentType == nil {
		panic("ObjectiveC: symbol method_copyArgumentType not loaded")
	}
	return _method_copyArgumentType(m, index)
}

var _method_copyReturnType func(m Method) *byte

// Method_copyReturnType returns a string describing a method’s return type.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_copyReturnType(_:)
func Method_copyReturnType(m Method) *byte {
	if _method_copyReturnType == nil {
		panic("ObjectiveC: symbol method_copyReturnType not loaded")
	}
	return _method_copyReturnType(m)
}

var _method_exchangeImplementations func(m1 Method, m2 Method)

// Method_exchangeImplementations exchanges the implementations of two methods.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_exchangeImplementations(_:_:)
func Method_exchangeImplementations(m1 Method, m2 Method) {
	if _method_exchangeImplementations == nil {
		panic("ObjectiveC: symbol method_exchangeImplementations not loaded")
	}
	_method_exchangeImplementations(m1, m2)
}

var _method_getArgumentType func(m Method, index uint, dst *byte, dst_len uintptr)

// Method_getArgumentType returns by reference a string describing a single parameter type of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getArgumentType(_:_:_:_:)
func Method_getArgumentType(m Method, index uint, dst *byte, dst_len uintptr) {
	if _method_getArgumentType == nil {
		panic("ObjectiveC: symbol method_getArgumentType not loaded")
	}
	_method_getArgumentType(m, index, dst, dst_len)
}

var _method_getDescription func(m Method) *Objc_method_description

// Method_getDescription returns a method description structure for a specified method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getDescription(_:)
func Method_getDescription(m Method) *Objc_method_description {
	if _method_getDescription == nil {
		panic("ObjectiveC: symbol method_getDescription not loaded")
	}
	return _method_getDescription(m)
}

var _method_getImplementation func(m Method) IMP

// Method_getImplementation returns the implementation of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getImplementation(_:)
func Method_getImplementation(m Method) IMP {
	if _method_getImplementation == nil {
		panic("ObjectiveC: symbol method_getImplementation not loaded")
	}
	return _method_getImplementation(m)
}

var _method_getName func(m Method) SEL

// Method_getName returns the name of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getName(_:)
func Method_getName(m Method) SEL {
	if _method_getName == nil {
		panic("ObjectiveC: symbol method_getName not loaded")
	}
	return _method_getName(m)
}

var _method_getNumberOfArguments func(m Method) uint

// Method_getNumberOfArguments returns the number of arguments accepted by a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getNumberOfArguments(_:)
func Method_getNumberOfArguments(m Method) uint {
	if _method_getNumberOfArguments == nil {
		panic("ObjectiveC: symbol method_getNumberOfArguments not loaded")
	}
	return _method_getNumberOfArguments(m)
}

var _method_getReturnType func(m Method, dst *byte, dst_len uintptr)

// Method_getReturnType returns by reference a string describing a method’s return type.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getReturnType(_:_:_:)
func Method_getReturnType(m Method, dst *byte, dst_len uintptr) {
	if _method_getReturnType == nil {
		panic("ObjectiveC: symbol method_getReturnType not loaded")
	}
	_method_getReturnType(m, dst, dst_len)
}

var _method_getTypeEncoding func(m Method) *byte

// Method_getTypeEncoding returns a string describing a method’s parameter and return types.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getTypeEncoding(_:)
func Method_getTypeEncoding(m Method) *byte {
	if _method_getTypeEncoding == nil {
		panic("ObjectiveC: symbol method_getTypeEncoding not loaded")
	}
	return _method_getTypeEncoding(m)
}

var _method_invoke func()

// Method_invoke calls the implementation of a specified method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_invoke
func Method_invoke() {
	if _method_invoke == nil {
		panic("ObjectiveC: symbol method_invoke not loaded")
	}
	_method_invoke()
}

var _method_setImplementation func(m Method, imp IMP) IMP

// Method_setImplementation sets the implementation of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_setImplementation(_:_:)
func Method_setImplementation(m Method, imp IMP) IMP {
	if _method_setImplementation == nil {
		panic("ObjectiveC: symbol method_setImplementation not loaded")
	}
	return _method_setImplementation(m, imp)
}

var _objc_addExceptionHandler func(fn unsafe.Pointer, context unsafe.Pointer) uintptr

// Objc_addExceptionHandler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_addExceptionHandler(_:_:)
func Objc_addExceptionHandler(fn unsafe.Pointer, context unsafe.Pointer) uintptr {
	if _objc_addExceptionHandler == nil {
		panic("ObjectiveC: symbol objc_addExceptionHandler not loaded")
	}
	return _objc_addExceptionHandler(fn, context)
}

var _objc_addLoadImageFunc func(func_ unsafe.Pointer)

// Objc_addLoadImageFunc.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_addLoadImageFunc(_:)
func Objc_addLoadImageFunc(func_ unsafe.Pointer) {
	if _objc_addLoadImageFunc == nil {
		panic("ObjectiveC: symbol objc_addLoadImageFunc not loaded")
	}
	_objc_addLoadImageFunc(func_)
}

var _objc_allocateClassPair func(superclass objc.Class, name string, extraBytes uintptr) objc.Class

// Objc_allocateClassPair creates a new class and metaclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_allocateClassPair(_:_:_:)
func Objc_allocateClassPair(superclass objc.Class, name string, extraBytes uintptr) objc.Class {
	if _objc_allocateClassPair == nil {
		panic("ObjectiveC: symbol objc_allocateClassPair not loaded")
	}
	return _objc_allocateClassPair(superclass, name, extraBytes)
}

var _objc_allocateProtocol func(name string) **Protocol

// Objc_allocateProtocol creates a new protocol instance.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_allocateProtocol(_:)
func Objc_allocateProtocol(name string) **Protocol {
	if _objc_allocateProtocol == nil {
		panic("ObjectiveC: symbol objc_allocateProtocol not loaded")
	}
	return _objc_allocateProtocol(name)
}

var _objc_begin_catch func(exc_buf unsafe.Pointer) Object

// Objc_begin_catch.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_begin_catch(_:)
func Objc_begin_catch(exc_buf unsafe.Pointer) Object {
	if _objc_begin_catch == nil {
		panic("ObjectiveC: symbol objc_begin_catch not loaded")
	}
	return _objc_begin_catch(exc_buf)
}

var _objc_constructInstance func(cls objc.Class, bytes unsafe.Pointer) Object

// Objc_constructInstance creates an instance of a class at the specified location.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_constructInstance
func Objc_constructInstance(cls objc.Class, bytes unsafe.Pointer) Object {
	if _objc_constructInstance == nil {
		panic("ObjectiveC: symbol objc_constructInstance not loaded")
	}
	return _objc_constructInstance(cls, bytes)
}

var _objc_copyClassList func(outCount *uint) objc.Class

// Objc_copyClassList creates and returns a list of pointers to all registered class definitions.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassList(_:)
func Objc_copyClassList(outCount *uint) objc.Class {
	if _objc_copyClassList == nil {
		panic("ObjectiveC: symbol objc_copyClassList not loaded")
	}
	return _objc_copyClassList(outCount)
}

var _objc_copyClassNamesForImage func(image string, outCount *uint) *byte

// Objc_copyClassNamesForImage returns the names of all the classes within a specified library or framework.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassNamesForImage(_:_:)
func Objc_copyClassNamesForImage(image string, outCount *uint) *byte {
	if _objc_copyClassNamesForImage == nil {
		panic("ObjectiveC: symbol objc_copyClassNamesForImage not loaded")
	}
	return _objc_copyClassNamesForImage(image, outCount)
}

var _objc_copyImageNames func(outCount *uint) *byte

// Objc_copyImageNames returns the names of all the loaded Objective-C frameworks and dynamic libraries.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyImageNames(_:)
func Objc_copyImageNames(outCount *uint) *byte {
	if _objc_copyImageNames == nil {
		panic("ObjectiveC: symbol objc_copyImageNames not loaded")
	}
	return _objc_copyImageNames(outCount)
}

var _objc_copyProtocolList func(outCount *uint) **Protocol

// Objc_copyProtocolList returns an array of all the protocols known to the runtime.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyProtocolList(_:)
func Objc_copyProtocolList(outCount *uint) **Protocol {
	if _objc_copyProtocolList == nil {
		panic("ObjectiveC: symbol objc_copyProtocolList not loaded")
	}
	return _objc_copyProtocolList(outCount)
}

var _objc_destructInstance func(obj Object) unsafe.Pointer

// Objc_destructInstance destroys an instance of a class without freeing memory and removes any of its associated references.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_destructInstance
func Objc_destructInstance(obj Object) unsafe.Pointer {
	if _objc_destructInstance == nil {
		panic("ObjectiveC: symbol objc_destructInstance not loaded")
	}
	return _objc_destructInstance(obj)
}

var _objc_disposeClassPair func(cls objc.Class)

// Objc_disposeClassPair destroys a class and its associated metaclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_disposeClassPair(_:)
func Objc_disposeClassPair(cls objc.Class) {
	if _objc_disposeClassPair == nil {
		panic("ObjectiveC: symbol objc_disposeClassPair not loaded")
	}
	_objc_disposeClassPair(cls)
}

var _objc_duplicateClass func(original objc.Class, name string, extraBytes uintptr) objc.Class

// Objc_duplicateClass used by Foundation’s Key-Value Observing.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_duplicateClass(_:_:_:)
func Objc_duplicateClass(original objc.Class, name string, extraBytes uintptr) objc.Class {
	if _objc_duplicateClass == nil {
		panic("ObjectiveC: symbol objc_duplicateClass not loaded")
	}
	return _objc_duplicateClass(original, name, extraBytes)
}

var _objc_end_catch func()

// Objc_end_catch.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_end_catch()
func Objc_end_catch() {
	if _objc_end_catch == nil {
		panic("ObjectiveC: symbol objc_end_catch not loaded")
	}
	_objc_end_catch()
}

var _objc_enumerateClasses func(image unsafe.Pointer, namePrefix string, conformingTo **Protocol, subclassing objc.Class)

// Objc_enumerateClasses.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_enumerateClasses
func Objc_enumerateClasses(image unsafe.Pointer, namePrefix string, conformingTo **Protocol, subclassing objc.Class) {
	if _objc_enumerateClasses == nil {
		panic("ObjectiveC: symbol objc_enumerateClasses not loaded")
	}
	_objc_enumerateClasses(image, namePrefix, conformingTo, subclassing)
}

var _objc_enumerationMutation func(obj Object)

// Objc_enumerationMutation inserted by the compiler when a mutation is detected during a foreach iteration.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_enumerationMutation(_:)
func Objc_enumerationMutation(obj Object) {
	if _objc_enumerationMutation == nil {
		panic("ObjectiveC: symbol objc_enumerationMutation not loaded")
	}
	_objc_enumerationMutation(obj)
}

var _objc_exception_rethrow func()

// Objc_exception_rethrow.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_rethrow()
func Objc_exception_rethrow() {
	if _objc_exception_rethrow == nil {
		panic("ObjectiveC: symbol objc_exception_rethrow not loaded")
	}
	_objc_exception_rethrow()
}

var _objc_exception_throw func(exception Object)

// Objc_exception_throw throw a runtime exception.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_throw(_:)
func Objc_exception_throw(exception Object) {
	if _objc_exception_throw == nil {
		panic("ObjectiveC: symbol objc_exception_throw not loaded")
	}
	_objc_exception_throw(exception)
}

var _objc_getAssociatedObject func(object Object, key unsafe.Pointer) Object

// Objc_getAssociatedObject returns the value associated with a given object for a given key.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getAssociatedObject(_:_:)
func Objc_getAssociatedObject(object Object, key unsafe.Pointer) Object {
	if _objc_getAssociatedObject == nil {
		panic("ObjectiveC: symbol objc_getAssociatedObject not loaded")
	}
	return _objc_getAssociatedObject(object, key)
}

var _objc_getClass func(name string) Object

// Objc_getClass returns the class definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getClass(_:)
func Objc_getClass(name string) Object {
	if _objc_getClass == nil {
		panic("ObjectiveC: symbol objc_getClass not loaded")
	}
	return _objc_getClass(name)
}

var _objc_getClassList func(buffer objc.Class, bufferCount int) int

// Objc_getClassList obtains the list of registered class definitions.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getClassList(_:_:)
func Objc_getClassList(buffer objc.Class, bufferCount int) int {
	if _objc_getClassList == nil {
		panic("ObjectiveC: symbol objc_getClassList not loaded")
	}
	return _objc_getClassList(buffer, bufferCount)
}

var _objc_getFutureClass func(name string) objc.Class

// Objc_getFutureClass used by CoreFoundation’s toll-free bridging.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getFutureClass
func Objc_getFutureClass(name string) objc.Class {
	if _objc_getFutureClass == nil {
		panic("ObjectiveC: symbol objc_getFutureClass not loaded")
	}
	return _objc_getFutureClass(name)
}

var _objc_getMetaClass func(name string) Object

// Objc_getMetaClass returns the metaclass definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getMetaClass(_:)
func Objc_getMetaClass(name string) Object {
	if _objc_getMetaClass == nil {
		panic("ObjectiveC: symbol objc_getMetaClass not loaded")
	}
	return _objc_getMetaClass(name)
}

var _objc_getProtocol func(name string) **Protocol

// Objc_getProtocol returns a specified protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getProtocol(_:)
func Objc_getProtocol(name string) **Protocol {
	if _objc_getProtocol == nil {
		panic("ObjectiveC: symbol objc_getProtocol not loaded")
	}
	return _objc_getProtocol(name)
}

var _objc_getRequiredClass func(name string) objc.Class

// Objc_getRequiredClass returns the class definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getRequiredClass(_:)
func Objc_getRequiredClass(name string) objc.Class {
	if _objc_getRequiredClass == nil {
		panic("ObjectiveC: symbol objc_getRequiredClass not loaded")
	}
	return _objc_getRequiredClass(name)
}

var _objc_loadWeak func(location uintptr) Object

// Objc_loadWeak loads the object referenced by a weak pointer and returns it.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_loadWeak(_:)
func Objc_loadWeak(location uintptr) Object {
	if _objc_loadWeak == nil {
		panic("ObjectiveC: symbol objc_loadWeak not loaded")
	}
	return _objc_loadWeak(location)
}

var _objc_lookUpClass func(name string) objc.Class

// Objc_lookUpClass returns the class definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_lookUpClass(_:)
func Objc_lookUpClass(name string) objc.Class {
	if _objc_lookUpClass == nil {
		panic("ObjectiveC: symbol objc_lookUpClass not loaded")
	}
	return _objc_lookUpClass(name)
}

var _objc_msgSend func()

// Objc_msgSend sends a message with a simple return value to an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend
func Objc_msgSend() {
	if _objc_msgSend == nil {
		panic("ObjectiveC: symbol objc_msgSend not loaded")
	}
	_objc_msgSend()
}

var _objc_msgSendSuper func()

// Objc_msgSendSuper sends a message with a simple return value to the superclass of an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_msgSendSuper
func Objc_msgSendSuper() {
	if _objc_msgSendSuper == nil {
		panic("ObjectiveC: symbol objc_msgSendSuper not loaded")
	}
	_objc_msgSendSuper()
}

var _objc_registerClassPair func(cls objc.Class)

// Objc_registerClassPair registers a class that was allocated using objc_allocateClassPair(_:_:_:).
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_registerClassPair(_:)
func Objc_registerClassPair(cls objc.Class) {
	if _objc_registerClassPair == nil {
		panic("ObjectiveC: symbol objc_registerClassPair not loaded")
	}
	_objc_registerClassPair(cls)
}

var _objc_registerProtocol func(proto **Protocol)

// Objc_registerProtocol registers a newly created protocol with the Objective-C runtime.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_registerProtocol(_:)
func Objc_registerProtocol(proto **Protocol) {
	if _objc_registerProtocol == nil {
		panic("ObjectiveC: symbol objc_registerProtocol not loaded")
	}
	_objc_registerProtocol(proto)
}

var _objc_removeAssociatedObjects func(object Object)

// Objc_removeAssociatedObjects removes all associations for a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_removeAssociatedObjects(_:)
func Objc_removeAssociatedObjects(object Object) {
	if _objc_removeAssociatedObjects == nil {
		panic("ObjectiveC: symbol objc_removeAssociatedObjects not loaded")
	}
	_objc_removeAssociatedObjects(object)
}

var _objc_removeExceptionHandler func(token uintptr)

// Objc_removeExceptionHandler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_removeExceptionHandler(_:)
func Objc_removeExceptionHandler(token uintptr) {
	if _objc_removeExceptionHandler == nil {
		panic("ObjectiveC: symbol objc_removeExceptionHandler not loaded")
	}
	_objc_removeExceptionHandler(token)
}

var _objc_setAssociatedObject func(object Object, key unsafe.Pointer, value Object, policy uintptr)

// Objc_setAssociatedObject sets an associated value for a given object using a given key and association policy.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setAssociatedObject(_:_:_:_:)
func Objc_setAssociatedObject(object Object, key unsafe.Pointer, value Object, policy uintptr) {
	if _objc_setAssociatedObject == nil {
		panic("ObjectiveC: symbol objc_setAssociatedObject not loaded")
	}
	_objc_setAssociatedObject(object, key, value, policy)
}

var _objc_setEnumerationMutationHandler func(handler func(Object))

// Objc_setEnumerationMutationHandler sets the current mutation handler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setEnumerationMutationHandler(_:)
func Objc_setEnumerationMutationHandler(handler func(Object)) {
	if _objc_setEnumerationMutationHandler == nil {
		panic("ObjectiveC: symbol objc_setEnumerationMutationHandler not loaded")
	}
	_objc_setEnumerationMutationHandler(handler)
}

var _objc_setExceptionMatcher func(fn unsafe.Pointer) unsafe.Pointer

// Objc_setExceptionMatcher.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionMatcher(_:)
func Objc_setExceptionMatcher(fn unsafe.Pointer) unsafe.Pointer {
	if _objc_setExceptionMatcher == nil {
		panic("ObjectiveC: symbol objc_setExceptionMatcher not loaded")
	}
	return _objc_setExceptionMatcher(fn)
}

var _objc_setExceptionPreprocessor func(fn unsafe.Pointer) unsafe.Pointer

// Objc_setExceptionPreprocessor.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionPreprocessor(_:)
func Objc_setExceptionPreprocessor(fn unsafe.Pointer) unsafe.Pointer {
	if _objc_setExceptionPreprocessor == nil {
		panic("ObjectiveC: symbol objc_setExceptionPreprocessor not loaded")
	}
	return _objc_setExceptionPreprocessor(fn)
}

var _objc_setForwardHandler func(fwd unsafe.Pointer, fwd_stret unsafe.Pointer)

// Objc_setForwardHandler set the function to be called by objc_msgForward.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setForwardHandler(_:_:)
func Objc_setForwardHandler(fwd unsafe.Pointer, fwd_stret unsafe.Pointer) {
	if _objc_setForwardHandler == nil {
		panic("ObjectiveC: symbol objc_setForwardHandler not loaded")
	}
	_objc_setForwardHandler(fwd, fwd_stret)
}

var _objc_setHook_getClass func(newValue unsafe.Pointer, outOldValue *Objc_hook_getClass)

// Objc_setHook_getClass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getClass(_:_:)
func Objc_setHook_getClass(newValue unsafe.Pointer, outOldValue *Objc_hook_getClass) {
	if _objc_setHook_getClass == nil {
		panic("ObjectiveC: symbol objc_setHook_getClass not loaded")
	}
	_objc_setHook_getClass(newValue, outOldValue)
}

var _objc_setHook_getImageName func(newValue unsafe.Pointer, outOldValue *Objc_hook_getImageName)

// Objc_setHook_getImageName.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getImageName(_:_:)
func Objc_setHook_getImageName(newValue unsafe.Pointer, outOldValue *Objc_hook_getImageName) {
	if _objc_setHook_getImageName == nil {
		panic("ObjectiveC: symbol objc_setHook_getImageName not loaded")
	}
	_objc_setHook_getImageName(newValue, outOldValue)
}

var _objc_setHook_lazyClassNamer func(newValue unsafe.Pointer, oldOutValue *Objc_hook_lazyClassNamer)

// Objc_setHook_lazyClassNamer.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_lazyClassNamer(_:_:)
func Objc_setHook_lazyClassNamer(newValue unsafe.Pointer, oldOutValue *Objc_hook_lazyClassNamer) {
	if _objc_setHook_lazyClassNamer == nil {
		panic("ObjectiveC: symbol objc_setHook_lazyClassNamer not loaded")
	}
	_objc_setHook_lazyClassNamer(newValue, oldOutValue)
}

var _objc_setUncaughtExceptionHandler func(fn unsafe.Pointer) unsafe.Pointer

// Objc_setUncaughtExceptionHandler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setUncaughtExceptionHandler(_:)
func Objc_setUncaughtExceptionHandler(fn unsafe.Pointer) unsafe.Pointer {
	if _objc_setUncaughtExceptionHandler == nil {
		panic("ObjectiveC: symbol objc_setUncaughtExceptionHandler not loaded")
	}
	return _objc_setUncaughtExceptionHandler(fn)
}

var _objc_storeWeak func(location uintptr, obj Object) Object

// Objc_storeWeak stores a new value in a `__weak` variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_storeWeak(_:_:)
func Objc_storeWeak(location uintptr, obj Object) Object {
	if _objc_storeWeak == nil {
		panic("ObjectiveC: symbol objc_storeWeak not loaded")
	}
	return _objc_storeWeak(location, obj)
}

var _objc_sync_enter func(obj Object) int

// Objc_sync_enter begin synchronizing on ‘obj’.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_sync_enter
func Objc_sync_enter(obj Object) int {
	if _objc_sync_enter == nil {
		panic("ObjectiveC: symbol objc_sync_enter not loaded")
	}
	return _objc_sync_enter(obj)
}

var _objc_sync_exit func(obj Object) int

// Objc_sync_exit end synchronizing on ‘obj’.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_sync_exit
func Objc_sync_exit(obj Object) int {
	if _objc_sync_exit == nil {
		panic("ObjectiveC: symbol objc_sync_exit not loaded")
	}
	return _objc_sync_exit(obj)
}

var _objc_terminate func()

// Objc_terminate.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_terminate()
func Objc_terminate() {
	if _objc_terminate == nil {
		panic("ObjectiveC: symbol objc_terminate not loaded")
	}
	_objc_terminate()
}

var _object_copy func(obj Object, size uintptr) Object

// Object_copy returns a copy of a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_copy
func Object_copy(obj Object, size uintptr) Object {
	if _object_copy == nil {
		panic("ObjectiveC: symbol object_copy not loaded")
	}
	return _object_copy(obj, size)
}

var _object_copyFromZone func(anObject Object, nBytes uintptr, zone unsafe.Pointer) Object

// Object_copyFromZone.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_copyFromZone
func Object_copyFromZone(anObject Object, nBytes uintptr, zone unsafe.Pointer) Object {
	if _object_copyFromZone == nil {
		panic("ObjectiveC: symbol object_copyFromZone not loaded")
	}
	return _object_copyFromZone(anObject, nBytes, zone)
}

var _object_dispose func(obj Object) Object

// Object_dispose frees the memory occupied by a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_dispose
func Object_dispose(obj Object) Object {
	if _object_dispose == nil {
		panic("ObjectiveC: symbol object_dispose not loaded")
	}
	return _object_dispose(obj)
}

var _object_getClass func(obj Object) objc.Class

// Object_getClass returns the class of an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getClass(_:)
func Object_getClass(obj Object) objc.Class {
	if _object_getClass == nil {
		panic("ObjectiveC: symbol object_getClass not loaded")
	}
	return _object_getClass(obj)
}

var _object_getClassName func(obj Object) *byte

// Object_getClassName returns the class name of a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getClassName(_:)
func Object_getClassName(obj Object) *byte {
	if _object_getClassName == nil {
		panic("ObjectiveC: symbol object_getClassName not loaded")
	}
	return _object_getClassName(obj)
}

var _object_getIndexedIvars func(obj Object) unsafe.Pointer

// Object_getIndexedIvars returns a pointer to any extra bytes allocated with a instance given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getIndexedIvars(_:)
func Object_getIndexedIvars(obj Object) unsafe.Pointer {
	if _object_getIndexedIvars == nil {
		panic("ObjectiveC: symbol object_getIndexedIvars not loaded")
	}
	return _object_getIndexedIvars(obj)
}

var _object_getInstanceVariable func(obj Object, name string, outValue unsafe.Pointer) Ivar

// Object_getInstanceVariable obtains the value of an instance variable of a class instance.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getInstanceVariable
func Object_getInstanceVariable(obj Object, name string, outValue unsafe.Pointer) Ivar {
	if _object_getInstanceVariable == nil {
		panic("ObjectiveC: symbol object_getInstanceVariable not loaded")
	}
	return _object_getInstanceVariable(obj, name, outValue)
}

var _object_getIvar func(obj Object, ivar Ivar) Object

// Object_getIvar reads the value of an instance variable in an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getIvar(_:_:)
func Object_getIvar(obj Object, ivar Ivar) Object {
	if _object_getIvar == nil {
		panic("ObjectiveC: symbol object_getIvar not loaded")
	}
	return _object_getIvar(obj, ivar)
}

var _object_isClass func(obj Object) bool

// Object_isClass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_isClass(_:)
func Object_isClass(obj Object) bool {
	if _object_isClass == nil {
		panic("ObjectiveC: symbol object_isClass not loaded")
	}
	return _object_isClass(obj)
}

var _object_setClass func(obj Object, cls objc.Class) objc.Class

// Object_setClass sets the class of an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setClass(_:_:)
func Object_setClass(obj Object, cls objc.Class) objc.Class {
	if _object_setClass == nil {
		panic("ObjectiveC: symbol object_setClass not loaded")
	}
	return _object_setClass(obj, cls)
}

var _object_setInstanceVariable func(obj Object, name string, value unsafe.Pointer) Ivar

// Object_setInstanceVariable changes the value of an instance variable of a class instance.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariable
func Object_setInstanceVariable(obj Object, name string, value unsafe.Pointer) Ivar {
	if _object_setInstanceVariable == nil {
		panic("ObjectiveC: symbol object_setInstanceVariable not loaded")
	}
	return _object_setInstanceVariable(obj, name, value)
}

var _object_setInstanceVariableWithStrongDefault func(obj Object, name string, value unsafe.Pointer) Ivar

// Object_setInstanceVariableWithStrongDefault.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariableWithStrongDefault
func Object_setInstanceVariableWithStrongDefault(obj Object, name string, value unsafe.Pointer) Ivar {
	if _object_setInstanceVariableWithStrongDefault == nil {
		panic("ObjectiveC: symbol object_setInstanceVariableWithStrongDefault not loaded")
	}
	return _object_setInstanceVariableWithStrongDefault(obj, name, value)
}

var _object_setIvar func(obj Object, ivar Ivar, value Object)

// Object_setIvar sets the value of an instance variable in an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setIvar(_:_:_:)
func Object_setIvar(obj Object, ivar Ivar, value Object) {
	if _object_setIvar == nil {
		panic("ObjectiveC: symbol object_setIvar not loaded")
	}
	_object_setIvar(obj, ivar, value)
}

var _object_setIvarWithStrongDefault func(obj Object, ivar Ivar, value Object)

// Object_setIvarWithStrongDefault.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setIvarWithStrongDefault(_:_:_:)
func Object_setIvarWithStrongDefault(obj Object, ivar Ivar, value Object) {
	if _object_setIvarWithStrongDefault == nil {
		panic("ObjectiveC: symbol object_setIvarWithStrongDefault not loaded")
	}
	_object_setIvarWithStrongDefault(obj, ivar, value)
}

var _property_copyAttributeList func(property unsafe.Pointer, outCount *uint) *Objc_property_attribute_t

// Property_copyAttributeList returns an array of property attributes for a given property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeList(_:_:)
func Property_copyAttributeList(property unsafe.Pointer, outCount *uint) *Objc_property_attribute_t {
	if _property_copyAttributeList == nil {
		panic("ObjectiveC: symbol property_copyAttributeList not loaded")
	}
	return _property_copyAttributeList(property, outCount)
}

var _property_copyAttributeValue func(property unsafe.Pointer, attributeName string) *byte

// Property_copyAttributeValue returns the value of a property attribute given the attribute name.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeValue(_:_:)
func Property_copyAttributeValue(property unsafe.Pointer, attributeName string) *byte {
	if _property_copyAttributeValue == nil {
		panic("ObjectiveC: symbol property_copyAttributeValue not loaded")
	}
	return _property_copyAttributeValue(property, attributeName)
}

var _property_getAttributes func(property unsafe.Pointer) *byte

// Property_getAttributes returns the attribute string of a property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_getAttributes(_:)
func Property_getAttributes(property unsafe.Pointer) *byte {
	if _property_getAttributes == nil {
		panic("ObjectiveC: symbol property_getAttributes not loaded")
	}
	return _property_getAttributes(property)
}

var _property_getName func(property unsafe.Pointer) *byte

// Property_getName returns the name of a property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_getName(_:)
func Property_getName(property unsafe.Pointer) *byte {
	if _property_getName == nil {
		panic("ObjectiveC: symbol property_getName not loaded")
	}
	return _property_getName(property)
}

var _protocol_addMethodDescription func(proto **Protocol, name SEL, types string, isRequiredMethod bool, isInstanceMethod bool)

// Protocol_addMethodDescription adds a method to a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_addMethodDescription(_:_:_:_:_:)
func Protocol_addMethodDescription(proto **Protocol, name SEL, types string, isRequiredMethod bool, isInstanceMethod bool) {
	if _protocol_addMethodDescription == nil {
		panic("ObjectiveC: symbol protocol_addMethodDescription not loaded")
	}
	_protocol_addMethodDescription(proto, name, types, isRequiredMethod, isInstanceMethod)
}

var _protocol_addProperty func(proto **Protocol, name string, attributes *Objc_property_attribute_t, attributeCount uint, isRequiredProperty bool, isInstanceProperty bool)

// Protocol_addProperty adds a property to a protocol that is under construction.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_addProperty(_:_:_:_:_:_:)
func Protocol_addProperty(proto **Protocol, name string, attributes *Objc_property_attribute_t, attributeCount uint, isRequiredProperty bool, isInstanceProperty bool) {
	if _protocol_addProperty == nil {
		panic("ObjectiveC: symbol protocol_addProperty not loaded")
	}
	_protocol_addProperty(proto, name, attributes, attributeCount, isRequiredProperty, isInstanceProperty)
}

var _protocol_addProtocol func(proto **Protocol, addition **Protocol)

// Protocol_addProtocol adds a registered protocol to another protocol that is under construction.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_addProtocol(_:_:)
func Protocol_addProtocol(proto **Protocol, addition **Protocol) {
	if _protocol_addProtocol == nil {
		panic("ObjectiveC: symbol protocol_addProtocol not loaded")
	}
	_protocol_addProtocol(proto, addition)
}

var _protocol_conformsToProtocol func(proto **Protocol, other **Protocol) bool

// Protocol_conformsToProtocol returns a Boolean value that indicates whether one protocol conforms to another protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_conformsToProtocol(_:_:)
func Protocol_conformsToProtocol(proto **Protocol, other **Protocol) bool {
	if _protocol_conformsToProtocol == nil {
		panic("ObjectiveC: symbol protocol_conformsToProtocol not loaded")
	}
	return _protocol_conformsToProtocol(proto, other)
}

var _protocol_copyMethodDescriptionList func(proto **Protocol, isRequiredMethod bool, isInstanceMethod bool, outCount *uint) *Objc_method_description

// Protocol_copyMethodDescriptionList returns an array of method descriptions of methods meeting a given specification for a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyMethodDescriptionList(_:_:_:_:)
func Protocol_copyMethodDescriptionList(proto **Protocol, isRequiredMethod bool, isInstanceMethod bool, outCount *uint) *Objc_method_description {
	if _protocol_copyMethodDescriptionList == nil {
		panic("ObjectiveC: symbol protocol_copyMethodDescriptionList not loaded")
	}
	return _protocol_copyMethodDescriptionList(proto, isRequiredMethod, isInstanceMethod, outCount)
}

var _protocol_copyPropertyList func(proto **Protocol, outCount *uint) *Objc_property_t

// Protocol_copyPropertyList returns an array of the properties declared by a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList(_:_:)
func Protocol_copyPropertyList(proto **Protocol, outCount *uint) *Objc_property_t {
	if _protocol_copyPropertyList == nil {
		panic("ObjectiveC: symbol protocol_copyPropertyList not loaded")
	}
	return _protocol_copyPropertyList(proto, outCount)
}

var _protocol_copyPropertyList2 func(proto **Protocol, outCount *uint, isRequiredProperty bool, isInstanceProperty bool) *Objc_property_t

// Protocol_copyPropertyList2.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList2(_:_:_:_:)
func Protocol_copyPropertyList2(proto **Protocol, outCount *uint, isRequiredProperty bool, isInstanceProperty bool) *Objc_property_t {
	if _protocol_copyPropertyList2 == nil {
		panic("ObjectiveC: symbol protocol_copyPropertyList2 not loaded")
	}
	return _protocol_copyPropertyList2(proto, outCount, isRequiredProperty, isInstanceProperty)
}

var _protocol_copyProtocolList func(proto **Protocol, outCount *uint) **Protocol

// Protocol_copyProtocolList returns an array of the protocols adopted by a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyProtocolList(_:_:)
func Protocol_copyProtocolList(proto **Protocol, outCount *uint) **Protocol {
	if _protocol_copyProtocolList == nil {
		panic("ObjectiveC: symbol protocol_copyProtocolList not loaded")
	}
	return _protocol_copyProtocolList(proto, outCount)
}

var _protocol_getMethodDescription func(proto **Protocol, aSel SEL, isRequiredMethod bool, isInstanceMethod bool) unsafe.Pointer

// Protocol_getMethodDescription returns a method description structure for a specified method of a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_getMethodDescription(_:_:_:_:)
func Protocol_getMethodDescription(proto **Protocol, aSel SEL, isRequiredMethod bool, isInstanceMethod bool) unsafe.Pointer {
	if _protocol_getMethodDescription == nil {
		panic("ObjectiveC: symbol protocol_getMethodDescription not loaded")
	}
	return _protocol_getMethodDescription(proto, aSel, isRequiredMethod, isInstanceMethod)
}

var _protocol_getName func(proto **Protocol) *byte

// Protocol_getName returns the name of a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_getName(_:)
func Protocol_getName(proto **Protocol) *byte {
	if _protocol_getName == nil {
		panic("ObjectiveC: symbol protocol_getName not loaded")
	}
	return _protocol_getName(proto)
}

var _protocol_getProperty func(proto **Protocol, name string, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer

// Protocol_getProperty returns the specified property of a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_getProperty(_:_:_:_:)
func Protocol_getProperty(proto **Protocol, name string, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer {
	if _protocol_getProperty == nil {
		panic("ObjectiveC: symbol protocol_getProperty not loaded")
	}
	return _protocol_getProperty(proto, name, isRequiredProperty, isInstanceProperty)
}

var _protocol_isEqual func(proto **Protocol, other **Protocol) bool

// Protocol_isEqual returns a Boolean value that indicates whether two protocols are equal.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_isEqual(_:_:)
func Protocol_isEqual(proto **Protocol, other **Protocol) bool {
	if _protocol_isEqual == nil {
		panic("ObjectiveC: symbol protocol_isEqual not loaded")
	}
	return _protocol_isEqual(proto, other)
}

var _sel_getName func(sel SEL) *byte

// Sel_getName returns the name of the method specified by a given selector.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_getName(_:)
func Sel_getName(sel SEL) *byte {
	if _sel_getName == nil {
		panic("ObjectiveC: symbol sel_getName not loaded")
	}
	return _sel_getName(sel)
}

var _sel_getUid func(str string) SEL

// Sel_getUid registers a method name with the Objective-C runtime system.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_getUid(_:)
func Sel_getUid(str string) SEL {
	if _sel_getUid == nil {
		panic("ObjectiveC: symbol sel_getUid not loaded")
	}
	return _sel_getUid(str)
}

var _sel_isEqual func(lhs SEL, rhs SEL) bool

// Sel_isEqual returns a Boolean value that indicates whether two selectors are equal.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_isEqual(_:_:)
func Sel_isEqual(lhs SEL, rhs SEL) bool {
	if _sel_isEqual == nil {
		panic("ObjectiveC: symbol sel_isEqual not loaded")
	}
	return _sel_isEqual(lhs, rhs)
}

var _sel_isMapped func(sel SEL) bool

// Sel_isMapped identifies a selector as being valid or invalid.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_isMapped(_:)
func Sel_isMapped(sel SEL) bool {
	if _sel_isMapped == nil {
		panic("ObjectiveC: symbol sel_isMapped not loaded")
	}
	return _sel_isMapped(sel)
}

var _sel_registerName func(str string) SEL

// Sel_registerName registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_registerName(_:)
func Sel_registerName(str string) SEL {
	if _sel_registerName == nil {
		panic("ObjectiveC: symbol sel_registerName not loaded")
	}
	return _sel_registerName(str)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_nXCompareHashTables, frameworkHandle, "NXCompareHashTables")
	registerFunc(&_nXCopyHashTable, frameworkHandle, "NXCopyHashTable")
	registerFunc(&_nXCountHashTable, frameworkHandle, "NXCountHashTable")
	registerFunc(&_nXCreateHashTable, frameworkHandle, "NXCreateHashTable")
	registerFunc(&_nXCreateHashTableFromZone, frameworkHandle, "NXCreateHashTableFromZone")
	registerFunc(&_nXEmptyHashTable, frameworkHandle, "NXEmptyHashTable")
	registerFunc(&_nXFreeHashTable, frameworkHandle, "NXFreeHashTable")
	registerFunc(&_nXHashGet, frameworkHandle, "NXHashGet")
	registerFunc(&_nXHashInsert, frameworkHandle, "NXHashInsert")
	registerFunc(&_nXHashInsertIfAbsent, frameworkHandle, "NXHashInsertIfAbsent")
	registerFunc(&_nXHashMember, frameworkHandle, "NXHashMember")
	registerFunc(&_nXHashRemove, frameworkHandle, "NXHashRemove")
	registerFunc(&_nXInitHashState, frameworkHandle, "NXInitHashState")
	registerFunc(&_nXNextHashState, frameworkHandle, "NXNextHashState")
	registerFunc(&_nXNoEffectFree, frameworkHandle, "NXNoEffectFree")
	registerFunc(&_nXPtrHash, frameworkHandle, "NXPtrHash")
	registerFunc(&_nXPtrIsEqual, frameworkHandle, "NXPtrIsEqual")
	registerFunc(&_nXReallyFree, frameworkHandle, "NXReallyFree")
	registerFunc(&_nXResetHashTable, frameworkHandle, "NXResetHashTable")
	registerFunc(&_nXStrHash, frameworkHandle, "NXStrHash")
	registerFunc(&_nXStrIsEqual, frameworkHandle, "NXStrIsEqual")
	registerFunc(&_class_addIvar, frameworkHandle, "class_addIvar")
	registerFunc(&_class_addMethod, frameworkHandle, "class_addMethod")
	registerFunc(&_class_addProperty, frameworkHandle, "class_addProperty")
	registerFunc(&_class_addProtocol, frameworkHandle, "class_addProtocol")
	registerFunc(&_class_conformsToProtocol, frameworkHandle, "class_conformsToProtocol")
	registerFunc(&_class_copyIvarList, frameworkHandle, "class_copyIvarList")
	registerFunc(&_class_copyMethodList, frameworkHandle, "class_copyMethodList")
	registerFunc(&_class_copyPropertyList, frameworkHandle, "class_copyPropertyList")
	registerFunc(&_class_copyProtocolList, frameworkHandle, "class_copyProtocolList")
	registerFunc(&_class_createInstance, frameworkHandle, "class_createInstance")
	registerFunc(&_class_createInstanceFromZone, frameworkHandle, "class_createInstanceFromZone")
	registerFunc(&_class_getClassMethod, frameworkHandle, "class_getClassMethod")
	registerFunc(&_class_getClassVariable, frameworkHandle, "class_getClassVariable")
	registerFunc(&_class_getImageName, frameworkHandle, "class_getImageName")
	registerFunc(&_class_getInstanceMethod, frameworkHandle, "class_getInstanceMethod")
	registerFunc(&_class_getInstanceSize, frameworkHandle, "class_getInstanceSize")
	registerFunc(&_class_getInstanceVariable, frameworkHandle, "class_getInstanceVariable")
	registerFunc(&_class_getIvarLayout, frameworkHandle, "class_getIvarLayout")
	registerFunc(&_class_getMethodImplementation, frameworkHandle, "class_getMethodImplementation")
	registerFunc(&_class_getName, frameworkHandle, "class_getName")
	registerFunc(&_class_getProperty, frameworkHandle, "class_getProperty")
	registerFunc(&_class_getSuperclass, frameworkHandle, "class_getSuperclass")
	registerFunc(&_class_getVersion, frameworkHandle, "class_getVersion")
	registerFunc(&_class_getWeakIvarLayout, frameworkHandle, "class_getWeakIvarLayout")
	registerFunc(&_class_isMetaClass, frameworkHandle, "class_isMetaClass")
	registerFunc(&_class_replaceMethod, frameworkHandle, "class_replaceMethod")
	registerFunc(&_class_replaceProperty, frameworkHandle, "class_replaceProperty")
	registerFunc(&_class_respondsToSelector, frameworkHandle, "class_respondsToSelector")
	registerFunc(&_class_setIvarLayout, frameworkHandle, "class_setIvarLayout")
	registerFunc(&_class_setVersion, frameworkHandle, "class_setVersion")
	registerFunc(&_class_setWeakIvarLayout, frameworkHandle, "class_setWeakIvarLayout")
	registerFunc(&_imp_getBlock, frameworkHandle, "imp_getBlock")
	registerFunc(&_imp_implementationWithBlock, frameworkHandle, "imp_implementationWithBlock")
	registerFunc(&_imp_removeBlock, frameworkHandle, "imp_removeBlock")
	registerFunc(&_ivar_getName, frameworkHandle, "ivar_getName")
	registerFunc(&_ivar_getOffset, frameworkHandle, "ivar_getOffset")
	registerFunc(&_ivar_getTypeEncoding, frameworkHandle, "ivar_getTypeEncoding")
	registerFunc(&_method_copyArgumentType, frameworkHandle, "method_copyArgumentType")
	registerFunc(&_method_copyReturnType, frameworkHandle, "method_copyReturnType")
	registerFunc(&_method_exchangeImplementations, frameworkHandle, "method_exchangeImplementations")
	registerFunc(&_method_getArgumentType, frameworkHandle, "method_getArgumentType")
	registerFunc(&_method_getDescription, frameworkHandle, "method_getDescription")
	registerFunc(&_method_getImplementation, frameworkHandle, "method_getImplementation")
	registerFunc(&_method_getName, frameworkHandle, "method_getName")
	registerFunc(&_method_getNumberOfArguments, frameworkHandle, "method_getNumberOfArguments")
	registerFunc(&_method_getReturnType, frameworkHandle, "method_getReturnType")
	registerFunc(&_method_getTypeEncoding, frameworkHandle, "method_getTypeEncoding")
	registerFunc(&_method_invoke, frameworkHandle, "method_invoke")
	registerFunc(&_method_setImplementation, frameworkHandle, "method_setImplementation")
	registerFunc(&_objc_addExceptionHandler, frameworkHandle, "objc_addExceptionHandler")
	registerFunc(&_objc_addLoadImageFunc, frameworkHandle, "objc_addLoadImageFunc")
	registerFunc(&_objc_allocateClassPair, frameworkHandle, "objc_allocateClassPair")
	registerFunc(&_objc_allocateProtocol, frameworkHandle, "objc_allocateProtocol")
	registerFunc(&_objc_begin_catch, frameworkHandle, "objc_begin_catch")
	registerFunc(&_objc_constructInstance, frameworkHandle, "objc_constructInstance")
	registerFunc(&_objc_copyClassList, frameworkHandle, "objc_copyClassList")
	registerFunc(&_objc_copyClassNamesForImage, frameworkHandle, "objc_copyClassNamesForImage")
	registerFunc(&_objc_copyImageNames, frameworkHandle, "objc_copyImageNames")
	registerFunc(&_objc_copyProtocolList, frameworkHandle, "objc_copyProtocolList")
	registerFunc(&_objc_destructInstance, frameworkHandle, "objc_destructInstance")
	registerFunc(&_objc_disposeClassPair, frameworkHandle, "objc_disposeClassPair")
	registerFunc(&_objc_duplicateClass, frameworkHandle, "objc_duplicateClass")
	registerFunc(&_objc_end_catch, frameworkHandle, "objc_end_catch")
	registerFunc(&_objc_enumerateClasses, frameworkHandle, "objc_enumerateClasses")
	registerFunc(&_objc_enumerationMutation, frameworkHandle, "objc_enumerationMutation")
	registerFunc(&_objc_exception_rethrow, frameworkHandle, "objc_exception_rethrow")
	registerFunc(&_objc_exception_throw, frameworkHandle, "objc_exception_throw")
	registerFunc(&_objc_getAssociatedObject, frameworkHandle, "objc_getAssociatedObject")
	registerFunc(&_objc_getClass, frameworkHandle, "objc_getClass")
	registerFunc(&_objc_getClassList, frameworkHandle, "objc_getClassList")
	registerFunc(&_objc_getFutureClass, frameworkHandle, "objc_getFutureClass")
	registerFunc(&_objc_getMetaClass, frameworkHandle, "objc_getMetaClass")
	registerFunc(&_objc_getProtocol, frameworkHandle, "objc_getProtocol")
	registerFunc(&_objc_getRequiredClass, frameworkHandle, "objc_getRequiredClass")
	registerFunc(&_objc_loadWeak, frameworkHandle, "objc_loadWeak")
	registerFunc(&_objc_lookUpClass, frameworkHandle, "objc_lookUpClass")
	registerFunc(&_objc_msgSend, frameworkHandle, "objc_msgSend")
	registerFunc(&_objc_msgSendSuper, frameworkHandle, "objc_msgSendSuper")
	registerFunc(&_objc_registerClassPair, frameworkHandle, "objc_registerClassPair")
	registerFunc(&_objc_registerProtocol, frameworkHandle, "objc_registerProtocol")
	registerFunc(&_objc_removeAssociatedObjects, frameworkHandle, "objc_removeAssociatedObjects")
	registerFunc(&_objc_removeExceptionHandler, frameworkHandle, "objc_removeExceptionHandler")
	registerFunc(&_objc_setAssociatedObject, frameworkHandle, "objc_setAssociatedObject")
	registerFunc(&_objc_setEnumerationMutationHandler, frameworkHandle, "objc_setEnumerationMutationHandler")
	registerFunc(&_objc_setExceptionMatcher, frameworkHandle, "objc_setExceptionMatcher")
	registerFunc(&_objc_setExceptionPreprocessor, frameworkHandle, "objc_setExceptionPreprocessor")
	registerFunc(&_objc_setForwardHandler, frameworkHandle, "objc_setForwardHandler")
	registerFunc(&_objc_setHook_getClass, frameworkHandle, "objc_setHook_getClass")
	registerFunc(&_objc_setHook_getImageName, frameworkHandle, "objc_setHook_getImageName")
	registerFunc(&_objc_setHook_lazyClassNamer, frameworkHandle, "objc_setHook_lazyClassNamer")
	registerFunc(&_objc_setUncaughtExceptionHandler, frameworkHandle, "objc_setUncaughtExceptionHandler")
	registerFunc(&_objc_storeWeak, frameworkHandle, "objc_storeWeak")
	registerFunc(&_objc_sync_enter, frameworkHandle, "objc_sync_enter")
	registerFunc(&_objc_sync_exit, frameworkHandle, "objc_sync_exit")
	registerFunc(&_objc_terminate, frameworkHandle, "objc_terminate")
	registerFunc(&_object_copy, frameworkHandle, "object_copy")
	registerFunc(&_object_copyFromZone, frameworkHandle, "object_copyFromZone")
	registerFunc(&_object_dispose, frameworkHandle, "object_dispose")
	registerFunc(&_object_getClass, frameworkHandle, "object_getClass")
	registerFunc(&_object_getClassName, frameworkHandle, "object_getClassName")
	registerFunc(&_object_getIndexedIvars, frameworkHandle, "object_getIndexedIvars")
	registerFunc(&_object_getInstanceVariable, frameworkHandle, "object_getInstanceVariable")
	registerFunc(&_object_getIvar, frameworkHandle, "object_getIvar")
	registerFunc(&_object_isClass, frameworkHandle, "object_isClass")
	registerFunc(&_object_setClass, frameworkHandle, "object_setClass")
	registerFunc(&_object_setInstanceVariable, frameworkHandle, "object_setInstanceVariable")
	registerFunc(&_object_setInstanceVariableWithStrongDefault, frameworkHandle, "object_setInstanceVariableWithStrongDefault")
	registerFunc(&_object_setIvar, frameworkHandle, "object_setIvar")
	registerFunc(&_object_setIvarWithStrongDefault, frameworkHandle, "object_setIvarWithStrongDefault")
	registerFunc(&_property_copyAttributeList, frameworkHandle, "property_copyAttributeList")
	registerFunc(&_property_copyAttributeValue, frameworkHandle, "property_copyAttributeValue")
	registerFunc(&_property_getAttributes, frameworkHandle, "property_getAttributes")
	registerFunc(&_property_getName, frameworkHandle, "property_getName")
	registerFunc(&_protocol_addMethodDescription, frameworkHandle, "protocol_addMethodDescription")
	registerFunc(&_protocol_addProperty, frameworkHandle, "protocol_addProperty")
	registerFunc(&_protocol_addProtocol, frameworkHandle, "protocol_addProtocol")
	registerFunc(&_protocol_conformsToProtocol, frameworkHandle, "protocol_conformsToProtocol")
	registerFunc(&_protocol_copyMethodDescriptionList, frameworkHandle, "protocol_copyMethodDescriptionList")
	registerFunc(&_protocol_copyPropertyList, frameworkHandle, "protocol_copyPropertyList")
	registerFunc(&_protocol_copyPropertyList2, frameworkHandle, "protocol_copyPropertyList2")
	registerFunc(&_protocol_copyProtocolList, frameworkHandle, "protocol_copyProtocolList")
	registerFunc(&_protocol_getMethodDescription, frameworkHandle, "protocol_getMethodDescription")
	registerFunc(&_protocol_getName, frameworkHandle, "protocol_getName")
	registerFunc(&_protocol_getProperty, frameworkHandle, "protocol_getProperty")
	registerFunc(&_protocol_isEqual, frameworkHandle, "protocol_isEqual")
	registerFunc(&_sel_getName, frameworkHandle, "sel_getName")
	registerFunc(&_sel_getUid, frameworkHandle, "sel_getUid")
	registerFunc(&_sel_isEqual, frameworkHandle, "sel_isEqual")
	registerFunc(&_sel_isMapped, frameworkHandle, "sel_isMapped")
	registerFunc(&_sel_registerName, frameworkHandle, "sel_registerName")
}
