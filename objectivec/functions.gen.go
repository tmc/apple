// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
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
		return fmt.Sprintf("ObjectiveC: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("ObjectiveC: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("ObjectiveC: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("ObjectiveC: register symbol %s: %v", name, r)
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

var _nXCompareHashTables func(table1 *NXHashTable, table2 *NXHashTable) bool
var _nXCompareHashTablesErr error

func tryNXCompareHashTables(table1 *NXHashTable, table2 *NXHashTable) (bool, error) {
	if _nXCompareHashTables == nil {
		return false, symbolCallError("NXCompareHashTables", "10.0", _nXCompareHashTablesErr)
	}
	return _nXCompareHashTables(table1, table2), nil
}

// NXCompareHashTables.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCompareHashTables
func NXCompareHashTables(table1 *NXHashTable, table2 *NXHashTable) bool {
	result, callErr := tryNXCompareHashTables(table1, table2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXCopyHashTable func(table *NXHashTable) *NXHashTable
var _nXCopyHashTableErr error

func tryNXCopyHashTable(table *NXHashTable) (*NXHashTable, error) {
	if _nXCopyHashTable == nil {
		return nil, symbolCallError("NXCopyHashTable", "10.0", _nXCopyHashTableErr)
	}
	return _nXCopyHashTable(table), nil
}

// NXCopyHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCopyHashTable
func NXCopyHashTable(table *NXHashTable) *NXHashTable {
	result, callErr := tryNXCopyHashTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXCountHashTable func(table *NXHashTable) uint
var _nXCountHashTableErr error

func tryNXCountHashTable(table *NXHashTable) (uint, error) {
	if _nXCountHashTable == nil {
		return 0, symbolCallError("NXCountHashTable", "10.0", _nXCountHashTableErr)
	}
	return _nXCountHashTable(table), nil
}

// NXCountHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCountHashTable
func NXCountHashTable(table *NXHashTable) uint {
	result, callErr := tryNXCountHashTable(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXCreateHashTable func(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer) *NXHashTable
var _nXCreateHashTableErr error

func tryNXCreateHashTable(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer) (*NXHashTable, error) {
	if _nXCreateHashTable == nil {
		return nil, symbolCallError("NXCreateHashTable", "10.0", _nXCreateHashTableErr)
	}
	return _nXCreateHashTable(prototype, capacity, info), nil
}

// NXCreateHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTable
func NXCreateHashTable(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer) *NXHashTable {
	result, callErr := tryNXCreateHashTable(prototype, capacity, info)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXCreateHashTableFromZone func(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer, zone unsafe.Pointer) *NXHashTable
var _nXCreateHashTableFromZoneErr error

func tryNXCreateHashTableFromZone(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer, zone unsafe.Pointer) (*NXHashTable, error) {
	if _nXCreateHashTableFromZone == nil {
		return nil, symbolCallError("NXCreateHashTableFromZone", "10.0", _nXCreateHashTableFromZoneErr)
	}
	return _nXCreateHashTableFromZone(prototype, capacity, info, zone), nil
}

// NXCreateHashTableFromZone.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXCreateHashTableFromZone
func NXCreateHashTableFromZone(prototype NXHashTablePrototype, capacity uint, info unsafe.Pointer, zone unsafe.Pointer) *NXHashTable {
	result, callErr := tryNXCreateHashTableFromZone(prototype, capacity, info, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXEmptyHashTable func(table *NXHashTable)
var _nXEmptyHashTableErr error

func tryNXEmptyHashTable(table *NXHashTable) error {
	if _nXEmptyHashTable == nil {
		return symbolCallError("NXEmptyHashTable", "10.0", _nXEmptyHashTableErr)
	}
	_nXEmptyHashTable(table)
	return nil
}

// NXEmptyHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXEmptyHashTable
func NXEmptyHashTable(table *NXHashTable) {
	if callErr := tryNXEmptyHashTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nXFreeHashTable func(table *NXHashTable)
var _nXFreeHashTableErr error

func tryNXFreeHashTable(table *NXHashTable) error {
	if _nXFreeHashTable == nil {
		return symbolCallError("NXFreeHashTable", "10.0", _nXFreeHashTableErr)
	}
	_nXFreeHashTable(table)
	return nil
}

// NXFreeHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXFreeHashTable
func NXFreeHashTable(table *NXHashTable) {
	if callErr := tryNXFreeHashTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nXHashGet func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer
var _nXHashGetErr error

func tryNXHashGet(table *NXHashTable, data unsafe.Pointer) (unsafe.Pointer, error) {
	if _nXHashGet == nil {
		return nil, symbolCallError("NXHashGet", "10.0", _nXHashGetErr)
	}
	return _nXHashGet(table, data), nil
}

// NXHashGet.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashGet
func NXHashGet(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNXHashGet(table, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXHashInsert func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer
var _nXHashInsertErr error

func tryNXHashInsert(table *NXHashTable, data unsafe.Pointer) (unsafe.Pointer, error) {
	if _nXHashInsert == nil {
		return nil, symbolCallError("NXHashInsert", "10.0", _nXHashInsertErr)
	}
	return _nXHashInsert(table, data), nil
}

// NXHashInsert.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashInsert
func NXHashInsert(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNXHashInsert(table, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXHashInsertIfAbsent func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer
var _nXHashInsertIfAbsentErr error

func tryNXHashInsertIfAbsent(table *NXHashTable, data unsafe.Pointer) (unsafe.Pointer, error) {
	if _nXHashInsertIfAbsent == nil {
		return nil, symbolCallError("NXHashInsertIfAbsent", "10.0", _nXHashInsertIfAbsentErr)
	}
	return _nXHashInsertIfAbsent(table, data), nil
}

// NXHashInsertIfAbsent.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashInsertIfAbsent
func NXHashInsertIfAbsent(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNXHashInsertIfAbsent(table, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXHashMember func(table *NXHashTable, data unsafe.Pointer) int
var _nXHashMemberErr error

func tryNXHashMember(table *NXHashTable, data unsafe.Pointer) (int, error) {
	if _nXHashMember == nil {
		return 0, symbolCallError("NXHashMember", "10.0", _nXHashMemberErr)
	}
	return _nXHashMember(table, data), nil
}

// NXHashMember.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashMember
func NXHashMember(table *NXHashTable, data unsafe.Pointer) int {
	result, callErr := tryNXHashMember(table, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXHashRemove func(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer
var _nXHashRemoveErr error

func tryNXHashRemove(table *NXHashTable, data unsafe.Pointer) (unsafe.Pointer, error) {
	if _nXHashRemove == nil {
		return nil, symbolCallError("NXHashRemove", "10.0", _nXHashRemoveErr)
	}
	return _nXHashRemove(table, data), nil
}

// NXHashRemove.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXHashRemove
func NXHashRemove(table *NXHashTable, data unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryNXHashRemove(table, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXInitHashState func(table *NXHashTable) NXHashState
var _nXInitHashStateErr error

func tryNXInitHashState(table *NXHashTable) (NXHashState, error) {
	if _nXInitHashState == nil {
		return NXHashState{}, symbolCallError("NXInitHashState", "10.0", _nXInitHashStateErr)
	}
	return _nXInitHashState(table), nil
}

// NXInitHashState.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXInitHashState
func NXInitHashState(table *NXHashTable) NXHashState {
	result, callErr := tryNXInitHashState(table)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXNextHashState func(table *NXHashTable, state *NXHashState, data unsafe.Pointer) int
var _nXNextHashStateErr error

func tryNXNextHashState(table *NXHashTable, state *NXHashState, data unsafe.Pointer) (int, error) {
	if _nXNextHashState == nil {
		return 0, symbolCallError("NXNextHashState", "10.0", _nXNextHashStateErr)
	}
	return _nXNextHashState(table, state, data), nil
}

// NXNextHashState.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXNextHashState
func NXNextHashState(table *NXHashTable, state *NXHashState, data unsafe.Pointer) int {
	result, callErr := tryNXNextHashState(table, state, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXNoEffectFree func(info unsafe.Pointer, data unsafe.Pointer)
var _nXNoEffectFreeErr error

func tryNXNoEffectFree(info unsafe.Pointer, data unsafe.Pointer) error {
	if _nXNoEffectFree == nil {
		return symbolCallError("NXNoEffectFree", "10.0", _nXNoEffectFreeErr)
	}
	_nXNoEffectFree(info, data)
	return nil
}

// NXNoEffectFree.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXNoEffectFree
func NXNoEffectFree(info unsafe.Pointer, data unsafe.Pointer) {
	if callErr := tryNXNoEffectFree(info, data); callErr != nil {
		panic(callErr)
	}
}

var _nXPtrHash func(info unsafe.Pointer, data unsafe.Pointer) uintptr
var _nXPtrHashErr error

func tryNXPtrHash(info unsafe.Pointer, data unsafe.Pointer) (uintptr, error) {
	if _nXPtrHash == nil {
		return 0, symbolCallError("NXPtrHash", "10.0", _nXPtrHashErr)
	}
	return _nXPtrHash(info, data), nil
}

// NXPtrHash.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXPtrHash
func NXPtrHash(info unsafe.Pointer, data unsafe.Pointer) uintptr {
	result, callErr := tryNXPtrHash(info, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXPtrIsEqual func(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int
var _nXPtrIsEqualErr error

func tryNXPtrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) (int, error) {
	if _nXPtrIsEqual == nil {
		return 0, symbolCallError("NXPtrIsEqual", "10.0", _nXPtrIsEqualErr)
	}
	return _nXPtrIsEqual(info, data1, data2), nil
}

// NXPtrIsEqual.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXPtrIsEqual
func NXPtrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	result, callErr := tryNXPtrIsEqual(info, data1, data2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXReallyFree func(info unsafe.Pointer, data unsafe.Pointer)
var _nXReallyFreeErr error

func tryNXReallyFree(info unsafe.Pointer, data unsafe.Pointer) error {
	if _nXReallyFree == nil {
		return symbolCallError("NXReallyFree", "10.0", _nXReallyFreeErr)
	}
	_nXReallyFree(info, data)
	return nil
}

// NXReallyFree.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXReallyFree
func NXReallyFree(info unsafe.Pointer, data unsafe.Pointer) {
	if callErr := tryNXReallyFree(info, data); callErr != nil {
		panic(callErr)
	}
}

var _nXResetHashTable func(table *NXHashTable)
var _nXResetHashTableErr error

func tryNXResetHashTable(table *NXHashTable) error {
	if _nXResetHashTable == nil {
		return symbolCallError("NXResetHashTable", "10.0", _nXResetHashTableErr)
	}
	_nXResetHashTable(table)
	return nil
}

// NXResetHashTable.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXResetHashTable
func NXResetHashTable(table *NXHashTable) {
	if callErr := tryNXResetHashTable(table); callErr != nil {
		panic(callErr)
	}
}

var _nXStrHash func(info unsafe.Pointer, data unsafe.Pointer) uintptr
var _nXStrHashErr error

func tryNXStrHash(info unsafe.Pointer, data unsafe.Pointer) (uintptr, error) {
	if _nXStrHash == nil {
		return 0, symbolCallError("NXStrHash", "10.0", _nXStrHashErr)
	}
	return _nXStrHash(info, data), nil
}

// NXStrHash.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXStrHash
func NXStrHash(info unsafe.Pointer, data unsafe.Pointer) uintptr {
	result, callErr := tryNXStrHash(info, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nXStrIsEqual func(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int
var _nXStrIsEqualErr error

func tryNXStrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) (int, error) {
	if _nXStrIsEqual == nil {
		return 0, symbolCallError("NXStrIsEqual", "10.0", _nXStrIsEqualErr)
	}
	return _nXStrIsEqual(info, data1, data2), nil
}

// NXStrIsEqual.
//
// Deprecated: Deprecated since macOS 10.1.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NXStrIsEqual
func NXStrIsEqual(info unsafe.Pointer, data1 unsafe.Pointer, data2 unsafe.Pointer) int {
	result, callErr := tryNXStrIsEqual(info, data1, data2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_addIvar func(cls objc.Class, name string, size uintptr, alignment uint8, types string) bool
var _class_addIvarErr error

func tryClass_addIvar(cls objc.Class, name string, size uintptr, alignment uint8, types string) (bool, error) {
	if _class_addIvar == nil {
		return false, symbolCallError("class_addIvar", "10.5", _class_addIvarErr)
	}
	return _class_addIvar(cls, name, size, alignment, types), nil
}

// Class_addIvar adds a new instance variable to a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addIvar(_:_:_:_:_:)
func Class_addIvar(cls objc.Class, name string, size uintptr, alignment uint8, types string) bool {
	result, callErr := tryClass_addIvar(cls, name, size, alignment, types)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_addMethod func(cls objc.Class, name SEL, imp IMP, types string) bool
var _class_addMethodErr error

func tryClass_addMethod(cls objc.Class, name SEL, imp IMP, types string) (bool, error) {
	if _class_addMethod == nil {
		return false, symbolCallError("class_addMethod", "10.5", _class_addMethodErr)
	}
	return _class_addMethod(cls, name, imp, types), nil
}

// Class_addMethod adds a new method to a class with a given name and implementation.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addMethod(_:_:_:_:)
func Class_addMethod(cls objc.Class, name SEL, imp IMP, types string) bool {
	result, callErr := tryClass_addMethod(cls, name, imp, types)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_addProperty func(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) bool
var _class_addPropertyErr error

func tryClass_addProperty(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) (bool, error) {
	if _class_addProperty == nil {
		return false, symbolCallError("class_addProperty", "10.7", _class_addPropertyErr)
	}
	return _class_addProperty(cls, name, attributes, attributeCount), nil
}

// Class_addProperty adds a property to a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addProperty(_:_:_:_:)
func Class_addProperty(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) bool {
	result, callErr := tryClass_addProperty(cls, name, attributes, attributeCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_addProtocol func(cls objc.Class, protocol_ **Protocol) bool
var _class_addProtocolErr error

func tryClass_addProtocol(cls objc.Class, protocol_ **Protocol) (bool, error) {
	if _class_addProtocol == nil {
		return false, symbolCallError("class_addProtocol", "10.5", _class_addProtocolErr)
	}
	return _class_addProtocol(cls, protocol_), nil
}

// Class_addProtocol adds a protocol to a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_addProtocol(_:_:)
func Class_addProtocol(cls objc.Class, protocol_ **Protocol) bool {
	result, callErr := tryClass_addProtocol(cls, protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_conformsToProtocol func(cls objc.Class, protocol_ **Protocol) bool
var _class_conformsToProtocolErr error

func tryClass_conformsToProtocol(cls objc.Class, protocol_ **Protocol) (bool, error) {
	if _class_conformsToProtocol == nil {
		return false, symbolCallError("class_conformsToProtocol", "10.5", _class_conformsToProtocolErr)
	}
	return _class_conformsToProtocol(cls, protocol_), nil
}

// Class_conformsToProtocol returns a Boolean value that indicates whether a class conforms to a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_conformsToProtocol(_:_:)
func Class_conformsToProtocol(cls objc.Class, protocol_ **Protocol) bool {
	result, callErr := tryClass_conformsToProtocol(cls, protocol_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_copyIvarList func(cls objc.Class, outCount *uint) Ivar
var _class_copyIvarListErr error

func tryClass_copyIvarList(cls objc.Class, outCount *uint) (Ivar, error) {
	if _class_copyIvarList == nil {
		return *new(Ivar), symbolCallError("class_copyIvarList", "10.5", _class_copyIvarListErr)
	}
	return _class_copyIvarList(cls, outCount), nil
}

// Class_copyIvarList describes the instance variables declared by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyIvarList(_:_:)
func Class_copyIvarList(cls objc.Class, outCount *uint) Ivar {
	result, callErr := tryClass_copyIvarList(cls, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_copyMethodList func(cls objc.Class, outCount *uint) Method
var _class_copyMethodListErr error

func tryClass_copyMethodList(cls objc.Class, outCount *uint) (Method, error) {
	if _class_copyMethodList == nil {
		return *new(Method), symbolCallError("class_copyMethodList", "10.5", _class_copyMethodListErr)
	}
	return _class_copyMethodList(cls, outCount), nil
}

// Class_copyMethodList describes the instance methods implemented by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyMethodList(_:_:)
func Class_copyMethodList(cls objc.Class, outCount *uint) Method {
	result, callErr := tryClass_copyMethodList(cls, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_copyPropertyList func(cls objc.Class, outCount *uint) *Objc_property_t
var _class_copyPropertyListErr error

func tryClass_copyPropertyList(cls objc.Class, outCount *uint) (*Objc_property_t, error) {
	if _class_copyPropertyList == nil {
		return nil, symbolCallError("class_copyPropertyList", "10.5", _class_copyPropertyListErr)
	}
	return _class_copyPropertyList(cls, outCount), nil
}

// Class_copyPropertyList describes the properties declared by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyPropertyList(_:_:)
func Class_copyPropertyList(cls objc.Class, outCount *uint) *Objc_property_t {
	result, callErr := tryClass_copyPropertyList(cls, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_copyProtocolList func(cls objc.Class, outCount *uint) **Protocol
var _class_copyProtocolListErr error

func tryClass_copyProtocolList(cls objc.Class, outCount *uint) (**Protocol, error) {
	if _class_copyProtocolList == nil {
		return nil, symbolCallError("class_copyProtocolList", "10.5", _class_copyProtocolListErr)
	}
	return _class_copyProtocolList(cls, outCount), nil
}

// Class_copyProtocolList describes the protocols adopted by a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_copyProtocolList(_:_:)
func Class_copyProtocolList(cls objc.Class, outCount *uint) **Protocol {
	result, callErr := tryClass_copyProtocolList(cls, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_createInstance func(cls objc.Class, extraBytes uintptr) Object
var _class_createInstanceErr error

func tryClass_createInstance(cls objc.Class, extraBytes uintptr) (Object, error) {
	if _class_createInstance == nil {
		return Object{}, symbolCallError("class_createInstance", "10.0", _class_createInstanceErr)
	}
	return _class_createInstance(cls, extraBytes), nil
}

// Class_createInstance creates an instance of a class, allocating memory for the class in the default malloc memory zone.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_createInstance(_:_:)
func Class_createInstance(cls objc.Class, extraBytes uintptr) Object {
	result, callErr := tryClass_createInstance(cls, extraBytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_createInstanceFromZone func(arg0 objc.Class, idxIvars uintptr, zone unsafe.Pointer) Object
var _class_createInstanceFromZoneErr error

func tryClass_createInstanceFromZone(arg0 objc.Class, idxIvars uintptr, zone unsafe.Pointer) (Object, error) {
	if _class_createInstanceFromZone == nil {
		return Object{}, symbolCallError("class_createInstanceFromZone", "10.0", _class_createInstanceFromZoneErr)
	}
	return _class_createInstanceFromZone(arg0, idxIvars, zone), nil
}

// Class_createInstanceFromZone.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_createInstanceFromZone
func Class_createInstanceFromZone(arg0 objc.Class, idxIvars uintptr, zone unsafe.Pointer) Object {
	result, callErr := tryClass_createInstanceFromZone(arg0, idxIvars, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getClassMethod func(cls objc.Class, name SEL) Method
var _class_getClassMethodErr error

func tryClass_getClassMethod(cls objc.Class, name SEL) (Method, error) {
	if _class_getClassMethod == nil {
		return *new(Method), symbolCallError("class_getClassMethod", "10.0", _class_getClassMethodErr)
	}
	return _class_getClassMethod(cls, name), nil
}

// Class_getClassMethod returns a pointer to the data structure describing a given class method for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getClassMethod(_:_:)
func Class_getClassMethod(cls objc.Class, name SEL) Method {
	result, callErr := tryClass_getClassMethod(cls, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getClassVariable func(cls objc.Class, name string) Ivar
var _class_getClassVariableErr error

func tryClass_getClassVariable(cls objc.Class, name string) (Ivar, error) {
	if _class_getClassVariable == nil {
		return *new(Ivar), symbolCallError("class_getClassVariable", "10.5", _class_getClassVariableErr)
	}
	return _class_getClassVariable(cls, name), nil
}

// Class_getClassVariable returns the [Ivar] for a specified class variable of a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getClassVariable(_:_:)
func Class_getClassVariable(cls objc.Class, name string) Ivar {
	result, callErr := tryClass_getClassVariable(cls, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getImageName func(cls objc.Class) *byte
var _class_getImageNameErr error

func tryClass_getImageName(cls objc.Class) (*byte, error) {
	if _class_getImageName == nil {
		return nil, symbolCallError("class_getImageName", "10.5", _class_getImageNameErr)
	}
	return _class_getImageName(cls), nil
}

// Class_getImageName returns the name of the dynamic library a class originated from.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getImageName(_:)
func Class_getImageName(cls objc.Class) *byte {
	result, callErr := tryClass_getImageName(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getInstanceMethod func(cls objc.Class, name SEL) Method
var _class_getInstanceMethodErr error

func tryClass_getInstanceMethod(cls objc.Class, name SEL) (Method, error) {
	if _class_getInstanceMethod == nil {
		return *new(Method), symbolCallError("class_getInstanceMethod", "10.0", _class_getInstanceMethodErr)
	}
	return _class_getInstanceMethod(cls, name), nil
}

// Class_getInstanceMethod returns a specified instance method for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceMethod(_:_:)
func Class_getInstanceMethod(cls objc.Class, name SEL) Method {
	result, callErr := tryClass_getInstanceMethod(cls, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getInstanceSize func(cls objc.Class) uintptr
var _class_getInstanceSizeErr error

func tryClass_getInstanceSize(cls objc.Class) (uintptr, error) {
	if _class_getInstanceSize == nil {
		return 0, symbolCallError("class_getInstanceSize", "10.5", _class_getInstanceSizeErr)
	}
	return _class_getInstanceSize(cls), nil
}

// Class_getInstanceSize returns the size of instances of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceSize(_:)
func Class_getInstanceSize(cls objc.Class) uintptr {
	result, callErr := tryClass_getInstanceSize(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getInstanceVariable func(cls objc.Class, name string) Ivar
var _class_getInstanceVariableErr error

func tryClass_getInstanceVariable(cls objc.Class, name string) (Ivar, error) {
	if _class_getInstanceVariable == nil {
		return *new(Ivar), symbolCallError("class_getInstanceVariable", "10.0", _class_getInstanceVariableErr)
	}
	return _class_getInstanceVariable(cls, name), nil
}

// Class_getInstanceVariable returns the [Ivar] for a specified instance variable of a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getInstanceVariable(_:_:)
func Class_getInstanceVariable(cls objc.Class, name string) Ivar {
	result, callErr := tryClass_getInstanceVariable(cls, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getIvarLayout func(cls objc.Class) *uint8
var _class_getIvarLayoutErr error

func tryClass_getIvarLayout(cls objc.Class) (*uint8, error) {
	if _class_getIvarLayout == nil {
		return nil, symbolCallError("class_getIvarLayout", "10.5", _class_getIvarLayoutErr)
	}
	return _class_getIvarLayout(cls), nil
}

// Class_getIvarLayout returns a description of the [Ivar] layout for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getIvarLayout(_:)
func Class_getIvarLayout(cls objc.Class) *uint8 {
	result, callErr := tryClass_getIvarLayout(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getMethodImplementation func(cls objc.Class, name SEL) IMP
var _class_getMethodImplementationErr error

func tryClass_getMethodImplementation(cls objc.Class, name SEL) (IMP, error) {
	if _class_getMethodImplementation == nil {
		return nil, symbolCallError("class_getMethodImplementation", "10.5", _class_getMethodImplementationErr)
	}
	return _class_getMethodImplementation(cls, name), nil
}

// Class_getMethodImplementation returns the function pointer that would be called if a particular message were sent to an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getMethodImplementation(_:_:)
func Class_getMethodImplementation(cls objc.Class, name SEL) IMP {
	result, callErr := tryClass_getMethodImplementation(cls, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getName func(cls objc.Class) *byte
var _class_getNameErr error

func tryClass_getName(cls objc.Class) (*byte, error) {
	if _class_getName == nil {
		return nil, symbolCallError("class_getName", "10.5", _class_getNameErr)
	}
	return _class_getName(cls), nil
}

// Class_getName returns the name of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getName(_:)
func Class_getName(cls objc.Class) *byte {
	result, callErr := tryClass_getName(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getProperty func(cls objc.Class, name string) unsafe.Pointer
var _class_getPropertyErr error

func tryClass_getProperty(cls objc.Class, name string) (unsafe.Pointer, error) {
	if _class_getProperty == nil {
		return nil, symbolCallError("class_getProperty", "10.5", _class_getPropertyErr)
	}
	return _class_getProperty(cls, name), nil
}

// Class_getProperty returns a property with a given name of a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getProperty(_:_:)
func Class_getProperty(cls objc.Class, name string) unsafe.Pointer {
	result, callErr := tryClass_getProperty(cls, name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getSuperclass func(cls objc.Class) objc.Class
var _class_getSuperclassErr error

func tryClass_getSuperclass(cls objc.Class) (objc.Class, error) {
	if _class_getSuperclass == nil {
		return 0, symbolCallError("class_getSuperclass", "10.5", _class_getSuperclassErr)
	}
	return _class_getSuperclass(cls), nil
}

// Class_getSuperclass returns the superclass of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getSuperclass(_:)
func Class_getSuperclass(cls objc.Class) objc.Class {
	result, callErr := tryClass_getSuperclass(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getVersion func(cls objc.Class) int
var _class_getVersionErr error

func tryClass_getVersion(cls objc.Class) (int, error) {
	if _class_getVersion == nil {
		return 0, symbolCallError("class_getVersion", "10.0", _class_getVersionErr)
	}
	return _class_getVersion(cls), nil
}

// Class_getVersion returns the version number of a class definition.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getVersion(_:)
func Class_getVersion(cls objc.Class) int {
	result, callErr := tryClass_getVersion(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_getWeakIvarLayout func(cls objc.Class) *uint8
var _class_getWeakIvarLayoutErr error

func tryClass_getWeakIvarLayout(cls objc.Class) (*uint8, error) {
	if _class_getWeakIvarLayout == nil {
		return nil, symbolCallError("class_getWeakIvarLayout", "10.5", _class_getWeakIvarLayoutErr)
	}
	return _class_getWeakIvarLayout(cls), nil
}

// Class_getWeakIvarLayout returns a description of the layout of weak [Ivar]s for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_getWeakIvarLayout(_:)
func Class_getWeakIvarLayout(cls objc.Class) *uint8 {
	result, callErr := tryClass_getWeakIvarLayout(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_isMetaClass func(cls objc.Class) bool
var _class_isMetaClassErr error

func tryClass_isMetaClass(cls objc.Class) (bool, error) {
	if _class_isMetaClass == nil {
		return false, symbolCallError("class_isMetaClass", "10.5", _class_isMetaClassErr)
	}
	return _class_isMetaClass(cls), nil
}

// Class_isMetaClass returns a Boolean value that indicates whether a class object is a metaclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_isMetaClass(_:)
func Class_isMetaClass(cls objc.Class) bool {
	result, callErr := tryClass_isMetaClass(cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_replaceMethod func(cls objc.Class, name SEL, imp IMP, types string) IMP
var _class_replaceMethodErr error

func tryClass_replaceMethod(cls objc.Class, name SEL, imp IMP, types string) (IMP, error) {
	if _class_replaceMethod == nil {
		return nil, symbolCallError("class_replaceMethod", "10.5", _class_replaceMethodErr)
	}
	return _class_replaceMethod(cls, name, imp, types), nil
}

// Class_replaceMethod replaces the implementation of a method for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_replaceMethod(_:_:_:_:)
func Class_replaceMethod(cls objc.Class, name SEL, imp IMP, types string) IMP {
	result, callErr := tryClass_replaceMethod(cls, name, imp, types)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_replaceProperty func(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint)
var _class_replacePropertyErr error

func tryClass_replaceProperty(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) error {
	if _class_replaceProperty == nil {
		return symbolCallError("class_replaceProperty", "10.7", _class_replacePropertyErr)
	}
	_class_replaceProperty(cls, name, attributes, attributeCount)
	return nil
}

// Class_replaceProperty replace a property of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_replaceProperty(_:_:_:_:)
func Class_replaceProperty(cls objc.Class, name string, attributes *Objc_property_attribute_t, attributeCount uint) {
	if callErr := tryClass_replaceProperty(cls, name, attributes, attributeCount); callErr != nil {
		panic(callErr)
	}
}

var _class_respondsToSelector func(cls objc.Class, sel SEL) bool
var _class_respondsToSelectorErr error

func tryClass_respondsToSelector(cls objc.Class, sel SEL) (bool, error) {
	if _class_respondsToSelector == nil {
		return false, symbolCallError("class_respondsToSelector", "10.5", _class_respondsToSelectorErr)
	}
	return _class_respondsToSelector(cls, sel), nil
}

// Class_respondsToSelector returns a Boolean value that indicates whether instances of a class respond to a particular selector.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_respondsToSelector(_:_:)
func Class_respondsToSelector(cls objc.Class, sel SEL) bool {
	result, callErr := tryClass_respondsToSelector(cls, sel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _class_setIvarLayout func(cls objc.Class, layout *uint8)
var _class_setIvarLayoutErr error

func tryClass_setIvarLayout(cls objc.Class, layout *uint8) error {
	if _class_setIvarLayout == nil {
		return symbolCallError("class_setIvarLayout", "10.5", _class_setIvarLayoutErr)
	}
	_class_setIvarLayout(cls, layout)
	return nil
}

// Class_setIvarLayout sets the [Ivar] layout for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_setIvarLayout(_:_:)
func Class_setIvarLayout(cls objc.Class, layout *uint8) {
	if callErr := tryClass_setIvarLayout(cls, layout); callErr != nil {
		panic(callErr)
	}
}

var _class_setVersion func(cls objc.Class, version int)
var _class_setVersionErr error

func tryClass_setVersion(cls objc.Class, version int) error {
	if _class_setVersion == nil {
		return symbolCallError("class_setVersion", "10.0", _class_setVersionErr)
	}
	_class_setVersion(cls, version)
	return nil
}

// Class_setVersion sets the version number of a class definition.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_setVersion(_:_:)
func Class_setVersion(cls objc.Class, version int) {
	if callErr := tryClass_setVersion(cls, version); callErr != nil {
		panic(callErr)
	}
}

var _class_setWeakIvarLayout func(cls objc.Class, layout *uint8)
var _class_setWeakIvarLayoutErr error

func tryClass_setWeakIvarLayout(cls objc.Class, layout *uint8) error {
	if _class_setWeakIvarLayout == nil {
		return symbolCallError("class_setWeakIvarLayout", "10.5", _class_setWeakIvarLayoutErr)
	}
	_class_setWeakIvarLayout(cls, layout)
	return nil
}

// Class_setWeakIvarLayout sets the layout for weak [Ivar]s for a given class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/class_setWeakIvarLayout(_:_:)
func Class_setWeakIvarLayout(cls objc.Class, layout *uint8) {
	if callErr := tryClass_setWeakIvarLayout(cls, layout); callErr != nil {
		panic(callErr)
	}
}

var _imp_getBlock func(anImp IMP) Object
var _imp_getBlockErr error

func tryImp_getBlock(anImp IMP) (Object, error) {
	if _imp_getBlock == nil {
		return Object{}, symbolCallError("imp_getBlock", "10.7", _imp_getBlockErr)
	}
	return _imp_getBlock(anImp), nil
}

// Imp_getBlock returns the block associated with an [IMP] that was created using imp_implementationWithBlock(_:).
//
// See: https://developer.apple.com/documentation/ObjectiveC/imp_getBlock(_:)
func Imp_getBlock(anImp IMP) Object {
	result, callErr := tryImp_getBlock(anImp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _imp_implementationWithBlock func(block Object) IMP
var _imp_implementationWithBlockErr error

func tryImp_implementationWithBlock(block Object) (IMP, error) {
	if _imp_implementationWithBlock == nil {
		return nil, symbolCallError("imp_implementationWithBlock", "10.7", _imp_implementationWithBlockErr)
	}
	return _imp_implementationWithBlock(block), nil
}

// Imp_implementationWithBlock creates a pointer to a function that calls the specified block when the method is called.
//
// See: https://developer.apple.com/documentation/ObjectiveC/imp_implementationWithBlock(_:)
func Imp_implementationWithBlock(block Object) IMP {
	result, callErr := tryImp_implementationWithBlock(block)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _imp_removeBlock func(anImp IMP) bool
var _imp_removeBlockErr error

func tryImp_removeBlock(anImp IMP) (bool, error) {
	if _imp_removeBlock == nil {
		return false, symbolCallError("imp_removeBlock", "10.7", _imp_removeBlockErr)
	}
	return _imp_removeBlock(anImp), nil
}

// Imp_removeBlock disassociates a block from an [IMP] that was created using imp_implementationWithBlock(_:), and releases the copy of the block that was created.
//
// See: https://developer.apple.com/documentation/ObjectiveC/imp_removeBlock(_:)
func Imp_removeBlock(anImp IMP) bool {
	result, callErr := tryImp_removeBlock(anImp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ivar_getName func(v Ivar) *byte
var _ivar_getNameErr error

func tryIvar_getName(v Ivar) (*byte, error) {
	if _ivar_getName == nil {
		return nil, symbolCallError("ivar_getName", "10.5", _ivar_getNameErr)
	}
	return _ivar_getName(v), nil
}

// Ivar_getName returns the name of an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/ivar_getName(_:)
func Ivar_getName(v Ivar) *byte {
	result, callErr := tryIvar_getName(v)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ivar_getOffset func(v Ivar) int
var _ivar_getOffsetErr error

func tryIvar_getOffset(v Ivar) (int, error) {
	if _ivar_getOffset == nil {
		return 0, symbolCallError("ivar_getOffset", "10.5", _ivar_getOffsetErr)
	}
	return _ivar_getOffset(v), nil
}

// Ivar_getOffset returns the offset of an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/ivar_getOffset(_:)
func Ivar_getOffset(v Ivar) int {
	result, callErr := tryIvar_getOffset(v)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _ivar_getTypeEncoding func(v Ivar) *byte
var _ivar_getTypeEncodingErr error

func tryIvar_getTypeEncoding(v Ivar) (*byte, error) {
	if _ivar_getTypeEncoding == nil {
		return nil, symbolCallError("ivar_getTypeEncoding", "10.5", _ivar_getTypeEncodingErr)
	}
	return _ivar_getTypeEncoding(v), nil
}

// Ivar_getTypeEncoding returns the type string of an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/ivar_getTypeEncoding(_:)
func Ivar_getTypeEncoding(v Ivar) *byte {
	result, callErr := tryIvar_getTypeEncoding(v)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_copyArgumentType func(m Method, index uint) *byte
var _method_copyArgumentTypeErr error

func tryMethod_copyArgumentType(m Method, index uint) (*byte, error) {
	if _method_copyArgumentType == nil {
		return nil, symbolCallError("method_copyArgumentType", "10.5", _method_copyArgumentTypeErr)
	}
	return _method_copyArgumentType(m, index), nil
}

// Method_copyArgumentType returns a string describing a single parameter type of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_copyArgumentType(_:_:)
func Method_copyArgumentType(m Method, index uint) *byte {
	result, callErr := tryMethod_copyArgumentType(m, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_copyReturnType func(m Method) *byte
var _method_copyReturnTypeErr error

func tryMethod_copyReturnType(m Method) (*byte, error) {
	if _method_copyReturnType == nil {
		return nil, symbolCallError("method_copyReturnType", "10.5", _method_copyReturnTypeErr)
	}
	return _method_copyReturnType(m), nil
}

// Method_copyReturnType returns a string describing a method’s return type.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_copyReturnType(_:)
func Method_copyReturnType(m Method) *byte {
	result, callErr := tryMethod_copyReturnType(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_exchangeImplementations func(m1 Method, m2 Method)
var _method_exchangeImplementationsErr error

func tryMethod_exchangeImplementations(m1 Method, m2 Method) error {
	if _method_exchangeImplementations == nil {
		return symbolCallError("method_exchangeImplementations", "10.5", _method_exchangeImplementationsErr)
	}
	_method_exchangeImplementations(m1, m2)
	return nil
}

// Method_exchangeImplementations exchanges the implementations of two methods.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_exchangeImplementations(_:_:)
func Method_exchangeImplementations(m1 Method, m2 Method) {
	if callErr := tryMethod_exchangeImplementations(m1, m2); callErr != nil {
		panic(callErr)
	}
}

var _method_getArgumentType func(m Method, index uint, dst *byte, dst_len uintptr)
var _method_getArgumentTypeErr error

func tryMethod_getArgumentType(m Method, index uint, dst *byte, dst_len uintptr) error {
	if _method_getArgumentType == nil {
		return symbolCallError("method_getArgumentType", "10.5", _method_getArgumentTypeErr)
	}
	_method_getArgumentType(m, index, dst, dst_len)
	return nil
}

// Method_getArgumentType returns by reference a string describing a single parameter type of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getArgumentType(_:_:_:_:)
func Method_getArgumentType(m Method, index uint, dst *byte, dst_len uintptr) {
	if callErr := tryMethod_getArgumentType(m, index, dst, dst_len); callErr != nil {
		panic(callErr)
	}
}

var _method_getDescription func(m Method) *Objc_method_description
var _method_getDescriptionErr error

func tryMethod_getDescription(m Method) (*Objc_method_description, error) {
	if _method_getDescription == nil {
		return nil, symbolCallError("method_getDescription", "10.5", _method_getDescriptionErr)
	}
	return _method_getDescription(m), nil
}

// Method_getDescription returns a method description structure for a specified method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getDescription(_:)
func Method_getDescription(m Method) *Objc_method_description {
	result, callErr := tryMethod_getDescription(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_getImplementation func(m Method) IMP
var _method_getImplementationErr error

func tryMethod_getImplementation(m Method) (IMP, error) {
	if _method_getImplementation == nil {
		return nil, symbolCallError("method_getImplementation", "10.5", _method_getImplementationErr)
	}
	return _method_getImplementation(m), nil
}

// Method_getImplementation returns the implementation of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getImplementation(_:)
func Method_getImplementation(m Method) IMP {
	result, callErr := tryMethod_getImplementation(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_getName func(m Method) SEL
var _method_getNameErr error

func tryMethod_getName(m Method) (SEL, error) {
	if _method_getName == nil {
		return *new(SEL), symbolCallError("method_getName", "10.5", _method_getNameErr)
	}
	return _method_getName(m), nil
}

// Method_getName returns the name of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getName(_:)
func Method_getName(m Method) SEL {
	result, callErr := tryMethod_getName(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_getNumberOfArguments func(m Method) uint
var _method_getNumberOfArgumentsErr error

func tryMethod_getNumberOfArguments(m Method) (uint, error) {
	if _method_getNumberOfArguments == nil {
		return 0, symbolCallError("method_getNumberOfArguments", "10.0", _method_getNumberOfArgumentsErr)
	}
	return _method_getNumberOfArguments(m), nil
}

// Method_getNumberOfArguments returns the number of arguments accepted by a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getNumberOfArguments(_:)
func Method_getNumberOfArguments(m Method) uint {
	result, callErr := tryMethod_getNumberOfArguments(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_getReturnType func(m Method, dst *byte, dst_len uintptr)
var _method_getReturnTypeErr error

func tryMethod_getReturnType(m Method, dst *byte, dst_len uintptr) error {
	if _method_getReturnType == nil {
		return symbolCallError("method_getReturnType", "10.5", _method_getReturnTypeErr)
	}
	_method_getReturnType(m, dst, dst_len)
	return nil
}

// Method_getReturnType returns by reference a string describing a method’s return type.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getReturnType(_:_:_:)
func Method_getReturnType(m Method, dst *byte, dst_len uintptr) {
	if callErr := tryMethod_getReturnType(m, dst, dst_len); callErr != nil {
		panic(callErr)
	}
}

var _method_getTypeEncoding func(m Method) *byte
var _method_getTypeEncodingErr error

func tryMethod_getTypeEncoding(m Method) (*byte, error) {
	if _method_getTypeEncoding == nil {
		return nil, symbolCallError("method_getTypeEncoding", "10.5", _method_getTypeEncodingErr)
	}
	return _method_getTypeEncoding(m), nil
}

// Method_getTypeEncoding returns a string describing a method’s parameter and return types.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_getTypeEncoding(_:)
func Method_getTypeEncoding(m Method) *byte {
	result, callErr := tryMethod_getTypeEncoding(m)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _method_invoke func()
var _method_invokeErr error

func tryMethod_invoke() error {
	if _method_invoke == nil {
		return symbolCallError("method_invoke", "10.5", _method_invokeErr)
	}
	_method_invoke()
	return nil
}

// Method_invoke calls the implementation of a specified method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_invoke
func Method_invoke() {
	if callErr := tryMethod_invoke(); callErr != nil {
		panic(callErr)
	}
}

var _method_setImplementation func(m Method, imp IMP) IMP
var _method_setImplementationErr error

func tryMethod_setImplementation(m Method, imp IMP) (IMP, error) {
	if _method_setImplementation == nil {
		return nil, symbolCallError("method_setImplementation", "10.5", _method_setImplementationErr)
	}
	return _method_setImplementation(m, imp), nil
}

// Method_setImplementation sets the implementation of a method.
//
// See: https://developer.apple.com/documentation/ObjectiveC/method_setImplementation(_:_:)
func Method_setImplementation(m Method, imp IMP) IMP {
	result, callErr := tryMethod_setImplementation(m, imp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_addExceptionHandler func(fn unsafe.Pointer, context unsafe.Pointer) uintptr
var _objc_addExceptionHandlerErr error

func tryObjc_addExceptionHandler(fn unsafe.Pointer, context unsafe.Pointer) (uintptr, error) {
	if _objc_addExceptionHandler == nil {
		return 0, symbolCallError("objc_addExceptionHandler", "10.5", _objc_addExceptionHandlerErr)
	}
	return _objc_addExceptionHandler(fn, context), nil
}

// Objc_addExceptionHandler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_addExceptionHandler(_:_:)
func Objc_addExceptionHandler(fn unsafe.Pointer, context unsafe.Pointer) uintptr {
	result, callErr := tryObjc_addExceptionHandler(fn, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_addLoadImageFunc func(func_ unsafe.Pointer)
var _objc_addLoadImageFuncErr error

func tryObjc_addLoadImageFunc(func_ unsafe.Pointer) error {
	if _objc_addLoadImageFunc == nil {
		return symbolCallError("objc_addLoadImageFunc", "10.15", _objc_addLoadImageFuncErr)
	}
	_objc_addLoadImageFunc(func_)
	return nil
}

// Objc_addLoadImageFunc.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_addLoadImageFunc(_:)
func Objc_addLoadImageFunc(func_ unsafe.Pointer) {
	if callErr := tryObjc_addLoadImageFunc(func_); callErr != nil {
		panic(callErr)
	}
}

var _objc_allocateClassPair func(superclass objc.Class, name string, extraBytes uintptr) objc.Class
var _objc_allocateClassPairErr error

func tryObjc_allocateClassPair(superclass objc.Class, name string, extraBytes uintptr) (objc.Class, error) {
	if _objc_allocateClassPair == nil {
		return 0, symbolCallError("objc_allocateClassPair", "10.5", _objc_allocateClassPairErr)
	}
	return _objc_allocateClassPair(superclass, name, extraBytes), nil
}

// Objc_allocateClassPair creates a new class and metaclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_allocateClassPair(_:_:_:)
func Objc_allocateClassPair(superclass objc.Class, name string, extraBytes uintptr) objc.Class {
	result, callErr := tryObjc_allocateClassPair(superclass, name, extraBytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_allocateProtocol func(name string) **Protocol
var _objc_allocateProtocolErr error

func tryObjc_allocateProtocol(name string) (**Protocol, error) {
	if _objc_allocateProtocol == nil {
		return nil, symbolCallError("objc_allocateProtocol", "10.7", _objc_allocateProtocolErr)
	}
	return _objc_allocateProtocol(name), nil
}

// Objc_allocateProtocol creates a new protocol instance.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_allocateProtocol(_:)
func Objc_allocateProtocol(name string) **Protocol {
	result, callErr := tryObjc_allocateProtocol(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_begin_catch func(exc_buf unsafe.Pointer) Object
var _objc_begin_catchErr error

func tryObjc_begin_catch(exc_buf unsafe.Pointer) (Object, error) {
	if _objc_begin_catch == nil {
		return Object{}, symbolCallError("objc_begin_catch", "10.5", _objc_begin_catchErr)
	}
	return _objc_begin_catch(exc_buf), nil
}

// Objc_begin_catch.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_begin_catch(_:)
func Objc_begin_catch(exc_buf unsafe.Pointer) Object {
	result, callErr := tryObjc_begin_catch(exc_buf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_constructInstance func(cls objc.Class, bytes unsafe.Pointer) Object
var _objc_constructInstanceErr error

func tryObjc_constructInstance(cls objc.Class, bytes unsafe.Pointer) (Object, error) {
	if _objc_constructInstance == nil {
		return Object{}, symbolCallError("objc_constructInstance", "10.6", _objc_constructInstanceErr)
	}
	return _objc_constructInstance(cls, bytes), nil
}

// Objc_constructInstance creates an instance of a class at the specified location.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_constructInstance
func Objc_constructInstance(cls objc.Class, bytes unsafe.Pointer) Object {
	result, callErr := tryObjc_constructInstance(cls, bytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_copyClassList func(outCount *uint) objc.Class
var _objc_copyClassListErr error

func tryObjc_copyClassList(outCount *uint) (objc.Class, error) {
	if _objc_copyClassList == nil {
		return 0, symbolCallError("objc_copyClassList", "10.7", _objc_copyClassListErr)
	}
	return _objc_copyClassList(outCount), nil
}

// Objc_copyClassList creates and returns a list of pointers to all registered class definitions.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassList(_:)
func Objc_copyClassList(outCount *uint) objc.Class {
	result, callErr := tryObjc_copyClassList(outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_copyClassNamesForImage func(image string, outCount *uint) *byte
var _objc_copyClassNamesForImageErr error

func tryObjc_copyClassNamesForImage(image string, outCount *uint) (*byte, error) {
	if _objc_copyClassNamesForImage == nil {
		return nil, symbolCallError("objc_copyClassNamesForImage", "10.5", _objc_copyClassNamesForImageErr)
	}
	return _objc_copyClassNamesForImage(image, outCount), nil
}

// Objc_copyClassNamesForImage returns the names of all the classes within a specified library or framework.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyClassNamesForImage(_:_:)
func Objc_copyClassNamesForImage(image string, outCount *uint) *byte {
	result, callErr := tryObjc_copyClassNamesForImage(image, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_copyImageNames func(outCount *uint) *byte
var _objc_copyImageNamesErr error

func tryObjc_copyImageNames(outCount *uint) (*byte, error) {
	if _objc_copyImageNames == nil {
		return nil, symbolCallError("objc_copyImageNames", "10.5", _objc_copyImageNamesErr)
	}
	return _objc_copyImageNames(outCount), nil
}

// Objc_copyImageNames returns the names of all the loaded Objective-C frameworks and dynamic libraries.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyImageNames(_:)
func Objc_copyImageNames(outCount *uint) *byte {
	result, callErr := tryObjc_copyImageNames(outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_copyProtocolList func(outCount *uint) **Protocol
var _objc_copyProtocolListErr error

func tryObjc_copyProtocolList(outCount *uint) (**Protocol, error) {
	if _objc_copyProtocolList == nil {
		return nil, symbolCallError("objc_copyProtocolList", "10.5", _objc_copyProtocolListErr)
	}
	return _objc_copyProtocolList(outCount), nil
}

// Objc_copyProtocolList returns an array of all the protocols known to the runtime.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_copyProtocolList(_:)
func Objc_copyProtocolList(outCount *uint) **Protocol {
	result, callErr := tryObjc_copyProtocolList(outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_destructInstance func(obj Object) unsafe.Pointer
var _objc_destructInstanceErr error

func tryObjc_destructInstance(obj Object) (unsafe.Pointer, error) {
	if _objc_destructInstance == nil {
		return nil, symbolCallError("objc_destructInstance", "10.6", _objc_destructInstanceErr)
	}
	return _objc_destructInstance(obj), nil
}

// Objc_destructInstance destroys an instance of a class without freeing memory and removes any of its associated references.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_destructInstance
func Objc_destructInstance(obj Object) unsafe.Pointer {
	result, callErr := tryObjc_destructInstance(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_disposeClassPair func(cls objc.Class)
var _objc_disposeClassPairErr error

func tryObjc_disposeClassPair(cls objc.Class) error {
	if _objc_disposeClassPair == nil {
		return symbolCallError("objc_disposeClassPair", "10.5", _objc_disposeClassPairErr)
	}
	_objc_disposeClassPair(cls)
	return nil
}

// Objc_disposeClassPair destroys a class and its associated metaclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_disposeClassPair(_:)
func Objc_disposeClassPair(cls objc.Class) {
	if callErr := tryObjc_disposeClassPair(cls); callErr != nil {
		panic(callErr)
	}
}

var _objc_duplicateClass func(original objc.Class, name string, extraBytes uintptr) objc.Class
var _objc_duplicateClassErr error

func tryObjc_duplicateClass(original objc.Class, name string, extraBytes uintptr) (objc.Class, error) {
	if _objc_duplicateClass == nil {
		return 0, symbolCallError("objc_duplicateClass", "10.5", _objc_duplicateClassErr)
	}
	return _objc_duplicateClass(original, name, extraBytes), nil
}

// Objc_duplicateClass used by Foundation’s Key-Value Observing.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_duplicateClass(_:_:_:)
func Objc_duplicateClass(original objc.Class, name string, extraBytes uintptr) objc.Class {
	result, callErr := tryObjc_duplicateClass(original, name, extraBytes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_end_catch func()
var _objc_end_catchErr error

func tryObjc_end_catch() error {
	if _objc_end_catch == nil {
		return symbolCallError("objc_end_catch", "10.5", _objc_end_catchErr)
	}
	_objc_end_catch()
	return nil
}

// Objc_end_catch.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_end_catch()
func Objc_end_catch() {
	if callErr := tryObjc_end_catch(); callErr != nil {
		panic(callErr)
	}
}

var _objc_enumerateClasses func(image unsafe.Pointer, namePrefix string, conformingTo **Protocol, subclassing objc.Class)
var _objc_enumerateClassesErr error

func tryObjc_enumerateClasses(image unsafe.Pointer, namePrefix string, conformingTo **Protocol, subclassing objc.Class) error {
	if _objc_enumerateClasses == nil {
		return symbolCallError("objc_enumerateClasses", "13.0", _objc_enumerateClassesErr)
	}
	_objc_enumerateClasses(image, namePrefix, conformingTo, subclassing)
	return nil
}

// Objc_enumerateClasses.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_enumerateClasses
func Objc_enumerateClasses(image unsafe.Pointer, namePrefix string, conformingTo **Protocol, subclassing objc.Class) {
	if callErr := tryObjc_enumerateClasses(image, namePrefix, conformingTo, subclassing); callErr != nil {
		panic(callErr)
	}
}

var _objc_enumerationMutation func(obj Object)
var _objc_enumerationMutationErr error

func tryObjc_enumerationMutation(obj Object) error {
	if _objc_enumerationMutation == nil {
		return symbolCallError("objc_enumerationMutation", "10.5", _objc_enumerationMutationErr)
	}
	_objc_enumerationMutation(obj)
	return nil
}

// Objc_enumerationMutation inserted by the compiler when a mutation is detected during a foreach iteration.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_enumerationMutation(_:)
func Objc_enumerationMutation(obj Object) {
	if callErr := tryObjc_enumerationMutation(obj); callErr != nil {
		panic(callErr)
	}
}

var _objc_exception_rethrow func()
var _objc_exception_rethrowErr error

func tryObjc_exception_rethrow() error {
	if _objc_exception_rethrow == nil {
		return symbolCallError("objc_exception_rethrow", "10.5", _objc_exception_rethrowErr)
	}
	_objc_exception_rethrow()
	return nil
}

// Objc_exception_rethrow.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_rethrow()
func Objc_exception_rethrow() {
	if callErr := tryObjc_exception_rethrow(); callErr != nil {
		panic(callErr)
	}
}

var _objc_exception_throw func(exception Object)
var _objc_exception_throwErr error

func tryObjc_exception_throw(exception Object) error {
	if _objc_exception_throw == nil {
		return symbolCallError("objc_exception_throw", "10.5", _objc_exception_throwErr)
	}
	_objc_exception_throw(exception)
	return nil
}

// Objc_exception_throw throw a runtime exception.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_throw(_:)
func Objc_exception_throw(exception Object) {
	if callErr := tryObjc_exception_throw(exception); callErr != nil {
		panic(callErr)
	}
}

var _objc_getAssociatedObject func(object Object, key unsafe.Pointer) Object
var _objc_getAssociatedObjectErr error

func tryObjc_getAssociatedObject(object Object, key unsafe.Pointer) (Object, error) {
	if _objc_getAssociatedObject == nil {
		return Object{}, symbolCallError("objc_getAssociatedObject", "10.6", _objc_getAssociatedObjectErr)
	}
	return _objc_getAssociatedObject(object, key), nil
}

// Objc_getAssociatedObject returns the value associated with a given object for a given key.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getAssociatedObject(_:_:)
func Objc_getAssociatedObject(object Object, key unsafe.Pointer) Object {
	result, callErr := tryObjc_getAssociatedObject(object, key)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_getClass func(name string) Object
var _objc_getClassErr error

func tryObjc_getClass(name string) (Object, error) {
	if _objc_getClass == nil {
		return Object{}, symbolCallError("objc_getClass", "10.0", _objc_getClassErr)
	}
	return _objc_getClass(name), nil
}

// Objc_getClass returns the class definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getClass(_:)
func Objc_getClass(name string) Object {
	result, callErr := tryObjc_getClass(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_getClassList func(buffer objc.Class, bufferCount int) int
var _objc_getClassListErr error

func tryObjc_getClassList(buffer objc.Class, bufferCount int) (int, error) {
	if _objc_getClassList == nil {
		return 0, symbolCallError("objc_getClassList", "10.0", _objc_getClassListErr)
	}
	return _objc_getClassList(buffer, bufferCount), nil
}

// Objc_getClassList obtains the list of registered class definitions.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getClassList(_:_:)
func Objc_getClassList(buffer objc.Class, bufferCount int) int {
	result, callErr := tryObjc_getClassList(buffer, bufferCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_getFutureClass func(name string) objc.Class
var _objc_getFutureClassErr error

func tryObjc_getFutureClass(name string) (objc.Class, error) {
	if _objc_getFutureClass == nil {
		return 0, symbolCallError("objc_getFutureClass", "10.5", _objc_getFutureClassErr)
	}
	return _objc_getFutureClass(name), nil
}

// Objc_getFutureClass used by CoreFoundation’s toll-free bridging.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getFutureClass
func Objc_getFutureClass(name string) objc.Class {
	result, callErr := tryObjc_getFutureClass(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_getMetaClass func(name string) Object
var _objc_getMetaClassErr error

func tryObjc_getMetaClass(name string) (Object, error) {
	if _objc_getMetaClass == nil {
		return Object{}, symbolCallError("objc_getMetaClass", "10.0", _objc_getMetaClassErr)
	}
	return _objc_getMetaClass(name), nil
}

// Objc_getMetaClass returns the metaclass definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getMetaClass(_:)
func Objc_getMetaClass(name string) Object {
	result, callErr := tryObjc_getMetaClass(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_getProtocol func(name string) **Protocol
var _objc_getProtocolErr error

func tryObjc_getProtocol(name string) (**Protocol, error) {
	if _objc_getProtocol == nil {
		return nil, symbolCallError("objc_getProtocol", "10.5", _objc_getProtocolErr)
	}
	return _objc_getProtocol(name), nil
}

// Objc_getProtocol returns a specified protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getProtocol(_:)
func Objc_getProtocol(name string) **Protocol {
	result, callErr := tryObjc_getProtocol(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_getRequiredClass func(name string) objc.Class
var _objc_getRequiredClassErr error

func tryObjc_getRequiredClass(name string) (objc.Class, error) {
	if _objc_getRequiredClass == nil {
		return 0, symbolCallError("objc_getRequiredClass", "10.0", _objc_getRequiredClassErr)
	}
	return _objc_getRequiredClass(name), nil
}

// Objc_getRequiredClass returns the class definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_getRequiredClass(_:)
func Objc_getRequiredClass(name string) objc.Class {
	result, callErr := tryObjc_getRequiredClass(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_loadWeak func(location uintptr) Object
var _objc_loadWeakErr error

func tryObjc_loadWeak(location uintptr) (Object, error) {
	if _objc_loadWeak == nil {
		return Object{}, symbolCallError("objc_loadWeak", "10.7", _objc_loadWeakErr)
	}
	return _objc_loadWeak(location), nil
}

// Objc_loadWeak loads the object referenced by a weak pointer and returns it.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_loadWeak(_:)
func Objc_loadWeak(location uintptr) Object {
	result, callErr := tryObjc_loadWeak(location)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_lookUpClass func(name string) objc.Class
var _objc_lookUpClassErr error

func tryObjc_lookUpClass(name string) (objc.Class, error) {
	if _objc_lookUpClass == nil {
		return 0, symbolCallError("objc_lookUpClass", "10.0", _objc_lookUpClassErr)
	}
	return _objc_lookUpClass(name), nil
}

// Objc_lookUpClass returns the class definition of a specified class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_lookUpClass(_:)
func Objc_lookUpClass(name string) objc.Class {
	result, callErr := tryObjc_lookUpClass(name)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_msgSend func()
var _objc_msgSendErr error

func tryObjc_msgSend() error {
	if _objc_msgSend == nil {
		return symbolCallError("objc_msgSend", "10.0", _objc_msgSendErr)
	}
	_objc_msgSend()
	return nil
}

// Objc_msgSend sends a message with a simple return value to an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_msgSend
func Objc_msgSend() {
	if callErr := tryObjc_msgSend(); callErr != nil {
		panic(callErr)
	}
}

var _objc_msgSendSuper func()
var _objc_msgSendSuperErr error

func tryObjc_msgSendSuper() error {
	if _objc_msgSendSuper == nil {
		return symbolCallError("objc_msgSendSuper", "10.0", _objc_msgSendSuperErr)
	}
	_objc_msgSendSuper()
	return nil
}

// Objc_msgSendSuper sends a message with a simple return value to the superclass of an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_msgSendSuper
func Objc_msgSendSuper() {
	if callErr := tryObjc_msgSendSuper(); callErr != nil {
		panic(callErr)
	}
}

var _objc_registerClassPair func(cls objc.Class)
var _objc_registerClassPairErr error

func tryObjc_registerClassPair(cls objc.Class) error {
	if _objc_registerClassPair == nil {
		return symbolCallError("objc_registerClassPair", "10.5", _objc_registerClassPairErr)
	}
	_objc_registerClassPair(cls)
	return nil
}

// Objc_registerClassPair registers a class that was allocated using objc_allocateClassPair(_:_:_:).
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_registerClassPair(_:)
func Objc_registerClassPair(cls objc.Class) {
	if callErr := tryObjc_registerClassPair(cls); callErr != nil {
		panic(callErr)
	}
}

var _objc_registerProtocol func(proto **Protocol)
var _objc_registerProtocolErr error

func tryObjc_registerProtocol(proto **Protocol) error {
	if _objc_registerProtocol == nil {
		return symbolCallError("objc_registerProtocol", "10.7", _objc_registerProtocolErr)
	}
	_objc_registerProtocol(proto)
	return nil
}

// Objc_registerProtocol registers a newly created protocol with the Objective-C runtime.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_registerProtocol(_:)
func Objc_registerProtocol(proto **Protocol) {
	if callErr := tryObjc_registerProtocol(proto); callErr != nil {
		panic(callErr)
	}
}

var _objc_removeAssociatedObjects func(object Object)
var _objc_removeAssociatedObjectsErr error

func tryObjc_removeAssociatedObjects(object Object) error {
	if _objc_removeAssociatedObjects == nil {
		return symbolCallError("objc_removeAssociatedObjects", "10.6", _objc_removeAssociatedObjectsErr)
	}
	_objc_removeAssociatedObjects(object)
	return nil
}

// Objc_removeAssociatedObjects removes all associations for a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_removeAssociatedObjects(_:)
func Objc_removeAssociatedObjects(object Object) {
	if callErr := tryObjc_removeAssociatedObjects(object); callErr != nil {
		panic(callErr)
	}
}

var _objc_removeExceptionHandler func(token uintptr)
var _objc_removeExceptionHandlerErr error

func tryObjc_removeExceptionHandler(token uintptr) error {
	if _objc_removeExceptionHandler == nil {
		return symbolCallError("objc_removeExceptionHandler", "10.5", _objc_removeExceptionHandlerErr)
	}
	_objc_removeExceptionHandler(token)
	return nil
}

// Objc_removeExceptionHandler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_removeExceptionHandler(_:)
func Objc_removeExceptionHandler(token uintptr) {
	if callErr := tryObjc_removeExceptionHandler(token); callErr != nil {
		panic(callErr)
	}
}

var _objc_setAssociatedObject func(object Object, key unsafe.Pointer, value Object, policy Objc_AssociationPolicy)
var _objc_setAssociatedObjectErr error

func tryObjc_setAssociatedObject(object Object, key unsafe.Pointer, value Object, policy Objc_AssociationPolicy) error {
	if _objc_setAssociatedObject == nil {
		return symbolCallError("objc_setAssociatedObject", "10.6", _objc_setAssociatedObjectErr)
	}
	_objc_setAssociatedObject(object, key, value, policy)
	return nil
}

// Objc_setAssociatedObject sets an associated value for a given object using a given key and association policy.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setAssociatedObject(_:_:_:_:)
func Objc_setAssociatedObject(object Object, key unsafe.Pointer, value Object, policy Objc_AssociationPolicy) {
	if callErr := tryObjc_setAssociatedObject(object, key, value, policy); callErr != nil {
		panic(callErr)
	}
}

var _objc_setEnumerationMutationHandler func(handler func(Object))
var _objc_setEnumerationMutationHandlerErr error

func tryObjc_setEnumerationMutationHandler(handler func(Object)) error {
	if _objc_setEnumerationMutationHandler == nil {
		return symbolCallError("objc_setEnumerationMutationHandler", "10.5", _objc_setEnumerationMutationHandlerErr)
	}
	_objc_setEnumerationMutationHandler(handler)
	return nil
}

// Objc_setEnumerationMutationHandler sets the current mutation handler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setEnumerationMutationHandler(_:)
func Objc_setEnumerationMutationHandler(handler func(Object)) {
	if callErr := tryObjc_setEnumerationMutationHandler(handler); callErr != nil {
		panic(callErr)
	}
}

var _objc_setExceptionMatcher func(fn unsafe.Pointer) unsafe.Pointer
var _objc_setExceptionMatcherErr error

func tryObjc_setExceptionMatcher(fn unsafe.Pointer) (unsafe.Pointer, error) {
	if _objc_setExceptionMatcher == nil {
		return nil, symbolCallError("objc_setExceptionMatcher", "10.5", _objc_setExceptionMatcherErr)
	}
	return _objc_setExceptionMatcher(fn), nil
}

// Objc_setExceptionMatcher.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionMatcher(_:)
func Objc_setExceptionMatcher(fn unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryObjc_setExceptionMatcher(fn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_setExceptionPreprocessor func(fn unsafe.Pointer) unsafe.Pointer
var _objc_setExceptionPreprocessorErr error

func tryObjc_setExceptionPreprocessor(fn unsafe.Pointer) (unsafe.Pointer, error) {
	if _objc_setExceptionPreprocessor == nil {
		return nil, symbolCallError("objc_setExceptionPreprocessor", "10.5", _objc_setExceptionPreprocessorErr)
	}
	return _objc_setExceptionPreprocessor(fn), nil
}

// Objc_setExceptionPreprocessor.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setExceptionPreprocessor(_:)
func Objc_setExceptionPreprocessor(fn unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryObjc_setExceptionPreprocessor(fn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_setForwardHandler func(fwd unsafe.Pointer, fwd_stret unsafe.Pointer)
var _objc_setForwardHandlerErr error

func tryObjc_setForwardHandler(fwd unsafe.Pointer, fwd_stret unsafe.Pointer) error {
	if _objc_setForwardHandler == nil {
		return symbolCallError("objc_setForwardHandler", "10.5", _objc_setForwardHandlerErr)
	}
	_objc_setForwardHandler(fwd, fwd_stret)
	return nil
}

// Objc_setForwardHandler set the function to be called by objc_msgForward.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setForwardHandler(_:_:)
func Objc_setForwardHandler(fwd unsafe.Pointer, fwd_stret unsafe.Pointer) {
	if callErr := tryObjc_setForwardHandler(fwd, fwd_stret); callErr != nil {
		panic(callErr)
	}
}

var _objc_setHook_getClass func(newValue unsafe.Pointer, outOldValue *Objc_hook_getClass)
var _objc_setHook_getClassErr error

func tryObjc_setHook_getClass(newValue unsafe.Pointer, outOldValue *Objc_hook_getClass) error {
	if _objc_setHook_getClass == nil {
		return symbolCallError("objc_setHook_getClass", "10.14.4", _objc_setHook_getClassErr)
	}
	_objc_setHook_getClass(newValue, outOldValue)
	return nil
}

// Objc_setHook_getClass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getClass(_:_:)
func Objc_setHook_getClass(newValue unsafe.Pointer, outOldValue *Objc_hook_getClass) {
	if callErr := tryObjc_setHook_getClass(newValue, outOldValue); callErr != nil {
		panic(callErr)
	}
}

var _objc_setHook_getImageName func(newValue unsafe.Pointer, outOldValue *Objc_hook_getImageName)
var _objc_setHook_getImageNameErr error

func tryObjc_setHook_getImageName(newValue unsafe.Pointer, outOldValue *Objc_hook_getImageName) error {
	if _objc_setHook_getImageName == nil {
		return symbolCallError("objc_setHook_getImageName", "10.14", _objc_setHook_getImageNameErr)
	}
	_objc_setHook_getImageName(newValue, outOldValue)
	return nil
}

// Objc_setHook_getImageName.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_getImageName(_:_:)
func Objc_setHook_getImageName(newValue unsafe.Pointer, outOldValue *Objc_hook_getImageName) {
	if callErr := tryObjc_setHook_getImageName(newValue, outOldValue); callErr != nil {
		panic(callErr)
	}
}

var _objc_setHook_lazyClassNamer func(newValue unsafe.Pointer, oldOutValue *Objc_hook_lazyClassNamer)
var _objc_setHook_lazyClassNamerErr error

func tryObjc_setHook_lazyClassNamer(newValue unsafe.Pointer, oldOutValue *Objc_hook_lazyClassNamer) error {
	if _objc_setHook_lazyClassNamer == nil {
		return symbolCallError("objc_setHook_lazyClassNamer", "11.0", _objc_setHook_lazyClassNamerErr)
	}
	_objc_setHook_lazyClassNamer(newValue, oldOutValue)
	return nil
}

// Objc_setHook_lazyClassNamer.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setHook_lazyClassNamer(_:_:)
func Objc_setHook_lazyClassNamer(newValue unsafe.Pointer, oldOutValue *Objc_hook_lazyClassNamer) {
	if callErr := tryObjc_setHook_lazyClassNamer(newValue, oldOutValue); callErr != nil {
		panic(callErr)
	}
}

var _objc_setUncaughtExceptionHandler func(fn unsafe.Pointer) unsafe.Pointer
var _objc_setUncaughtExceptionHandlerErr error

func tryObjc_setUncaughtExceptionHandler(fn unsafe.Pointer) (unsafe.Pointer, error) {
	if _objc_setUncaughtExceptionHandler == nil {
		return nil, symbolCallError("objc_setUncaughtExceptionHandler", "10.5", _objc_setUncaughtExceptionHandlerErr)
	}
	return _objc_setUncaughtExceptionHandler(fn), nil
}

// Objc_setUncaughtExceptionHandler.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_setUncaughtExceptionHandler(_:)
func Objc_setUncaughtExceptionHandler(fn unsafe.Pointer) unsafe.Pointer {
	result, callErr := tryObjc_setUncaughtExceptionHandler(fn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_storeWeak func(location uintptr, obj Object) Object
var _objc_storeWeakErr error

func tryObjc_storeWeak(location uintptr, obj Object) (Object, error) {
	if _objc_storeWeak == nil {
		return Object{}, symbolCallError("objc_storeWeak", "10.7", _objc_storeWeakErr)
	}
	return _objc_storeWeak(location, obj), nil
}

// Objc_storeWeak stores a new value in a `__weak` variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_storeWeak(_:_:)
func Objc_storeWeak(location uintptr, obj Object) Object {
	result, callErr := tryObjc_storeWeak(location, obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_sync_enter func(obj Object) int
var _objc_sync_enterErr error

func tryObjc_sync_enter(obj Object) (int, error) {
	if _objc_sync_enter == nil {
		return 0, symbolCallError("objc_sync_enter", "10.3", _objc_sync_enterErr)
	}
	return _objc_sync_enter(obj), nil
}

// Objc_sync_enter begin synchronizing on ‘obj’.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_sync_enter
func Objc_sync_enter(obj Object) int {
	result, callErr := tryObjc_sync_enter(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_sync_exit func(obj Object) int
var _objc_sync_exitErr error

func tryObjc_sync_exit(obj Object) (int, error) {
	if _objc_sync_exit == nil {
		return 0, symbolCallError("objc_sync_exit", "10.3", _objc_sync_exitErr)
	}
	return _objc_sync_exit(obj), nil
}

// Objc_sync_exit end synchronizing on ‘obj’.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_sync_exit
func Objc_sync_exit(obj Object) int {
	result, callErr := tryObjc_sync_exit(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _objc_terminate func()
var _objc_terminateErr error

func tryObjc_terminate() error {
	if _objc_terminate == nil {
		return symbolCallError("objc_terminate", "10.8", _objc_terminateErr)
	}
	_objc_terminate()
	return nil
}

// Objc_terminate.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_terminate()
func Objc_terminate() {
	if callErr := tryObjc_terminate(); callErr != nil {
		panic(callErr)
	}
}

var _object_copy func(obj Object, size uintptr) Object
var _object_copyErr error

func tryObject_copy(obj Object, size uintptr) (Object, error) {
	if _object_copy == nil {
		return Object{}, symbolCallError("object_copy", "10.0", _object_copyErr)
	}
	return _object_copy(obj, size), nil
}

// Object_copy returns a copy of a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_copy
func Object_copy(obj Object, size uintptr) Object {
	result, callErr := tryObject_copy(obj, size)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_copyFromZone func(anObject Object, nBytes uintptr, zone unsafe.Pointer) Object
var _object_copyFromZoneErr error

func tryObject_copyFromZone(anObject Object, nBytes uintptr, zone unsafe.Pointer) (Object, error) {
	if _object_copyFromZone == nil {
		return Object{}, symbolCallError("object_copyFromZone", "10.0", _object_copyFromZoneErr)
	}
	return _object_copyFromZone(anObject, nBytes, zone), nil
}

// Object_copyFromZone.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_copyFromZone
func Object_copyFromZone(anObject Object, nBytes uintptr, zone unsafe.Pointer) Object {
	result, callErr := tryObject_copyFromZone(anObject, nBytes, zone)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_dispose func(obj Object) Object
var _object_disposeErr error

func tryObject_dispose(obj Object) (Object, error) {
	if _object_dispose == nil {
		return Object{}, symbolCallError("object_dispose", "10.0", _object_disposeErr)
	}
	return _object_dispose(obj), nil
}

// Object_dispose frees the memory occupied by a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_dispose
func Object_dispose(obj Object) Object {
	result, callErr := tryObject_dispose(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_getClass func(obj Object) objc.Class
var _object_getClassErr error

func tryObject_getClass(obj Object) (objc.Class, error) {
	if _object_getClass == nil {
		return 0, symbolCallError("object_getClass", "10.5", _object_getClassErr)
	}
	return _object_getClass(obj), nil
}

// Object_getClass returns the class of an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getClass(_:)
func Object_getClass(obj Object) objc.Class {
	result, callErr := tryObject_getClass(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_getClassName func(obj Object) *byte
var _object_getClassNameErr error

func tryObject_getClassName(obj Object) (*byte, error) {
	if _object_getClassName == nil {
		return nil, symbolCallError("object_getClassName", "10.0", _object_getClassNameErr)
	}
	return _object_getClassName(obj), nil
}

// Object_getClassName returns the class name of a given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getClassName(_:)
func Object_getClassName(obj Object) *byte {
	result, callErr := tryObject_getClassName(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_getIndexedIvars func(obj Object) unsafe.Pointer
var _object_getIndexedIvarsErr error

func tryObject_getIndexedIvars(obj Object) (unsafe.Pointer, error) {
	if _object_getIndexedIvars == nil {
		return nil, symbolCallError("object_getIndexedIvars", "10.0", _object_getIndexedIvarsErr)
	}
	return _object_getIndexedIvars(obj), nil
}

// Object_getIndexedIvars returns a pointer to any extra bytes allocated with a instance given object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getIndexedIvars(_:)
func Object_getIndexedIvars(obj Object) unsafe.Pointer {
	result, callErr := tryObject_getIndexedIvars(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_getInstanceVariable func(obj Object, name string, outValue unsafe.Pointer) Ivar
var _object_getInstanceVariableErr error

func tryObject_getInstanceVariable(obj Object, name string, outValue unsafe.Pointer) (Ivar, error) {
	if _object_getInstanceVariable == nil {
		return *new(Ivar), symbolCallError("object_getInstanceVariable", "10.0", _object_getInstanceVariableErr)
	}
	return _object_getInstanceVariable(obj, name, outValue), nil
}

// Object_getInstanceVariable obtains the value of an instance variable of a class instance.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getInstanceVariable
func Object_getInstanceVariable(obj Object, name string, outValue unsafe.Pointer) Ivar {
	result, callErr := tryObject_getInstanceVariable(obj, name, outValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_getIvar func(obj Object, ivar Ivar) Object
var _object_getIvarErr error

func tryObject_getIvar(obj Object, ivar Ivar) (Object, error) {
	if _object_getIvar == nil {
		return Object{}, symbolCallError("object_getIvar", "10.5", _object_getIvarErr)
	}
	return _object_getIvar(obj, ivar), nil
}

// Object_getIvar reads the value of an instance variable in an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_getIvar(_:_:)
func Object_getIvar(obj Object, ivar Ivar) Object {
	result, callErr := tryObject_getIvar(obj, ivar)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_isClass func(obj Object) bool
var _object_isClassErr error

func tryObject_isClass(obj Object) (bool, error) {
	if _object_isClass == nil {
		return false, symbolCallError("object_isClass", "10.10", _object_isClassErr)
	}
	return _object_isClass(obj), nil
}

// Object_isClass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_isClass(_:)
func Object_isClass(obj Object) bool {
	result, callErr := tryObject_isClass(obj)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_setClass func(obj Object, cls objc.Class) objc.Class
var _object_setClassErr error

func tryObject_setClass(obj Object, cls objc.Class) (objc.Class, error) {
	if _object_setClass == nil {
		return 0, symbolCallError("object_setClass", "10.5", _object_setClassErr)
	}
	return _object_setClass(obj, cls), nil
}

// Object_setClass sets the class of an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setClass(_:_:)
func Object_setClass(obj Object, cls objc.Class) objc.Class {
	result, callErr := tryObject_setClass(obj, cls)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_setInstanceVariable func(obj Object, name string, value unsafe.Pointer) Ivar
var _object_setInstanceVariableErr error

func tryObject_setInstanceVariable(obj Object, name string, value unsafe.Pointer) (Ivar, error) {
	if _object_setInstanceVariable == nil {
		return *new(Ivar), symbolCallError("object_setInstanceVariable", "10.0", _object_setInstanceVariableErr)
	}
	return _object_setInstanceVariable(obj, name, value), nil
}

// Object_setInstanceVariable changes the value of an instance variable of a class instance.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariable
func Object_setInstanceVariable(obj Object, name string, value unsafe.Pointer) Ivar {
	result, callErr := tryObject_setInstanceVariable(obj, name, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_setInstanceVariableWithStrongDefault func(obj Object, name string, value unsafe.Pointer) Ivar
var _object_setInstanceVariableWithStrongDefaultErr error

func tryObject_setInstanceVariableWithStrongDefault(obj Object, name string, value unsafe.Pointer) (Ivar, error) {
	if _object_setInstanceVariableWithStrongDefault == nil {
		return *new(Ivar), symbolCallError("object_setInstanceVariableWithStrongDefault", "10.12", _object_setInstanceVariableWithStrongDefaultErr)
	}
	return _object_setInstanceVariableWithStrongDefault(obj, name, value), nil
}

// Object_setInstanceVariableWithStrongDefault.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setInstanceVariableWithStrongDefault
func Object_setInstanceVariableWithStrongDefault(obj Object, name string, value unsafe.Pointer) Ivar {
	result, callErr := tryObject_setInstanceVariableWithStrongDefault(obj, name, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _object_setIvar func(obj Object, ivar Ivar, value Object)
var _object_setIvarErr error

func tryObject_setIvar(obj Object, ivar Ivar, value Object) error {
	if _object_setIvar == nil {
		return symbolCallError("object_setIvar", "10.5", _object_setIvarErr)
	}
	_object_setIvar(obj, ivar, value)
	return nil
}

// Object_setIvar sets the value of an instance variable in an object.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setIvar(_:_:_:)
func Object_setIvar(obj Object, ivar Ivar, value Object) {
	if callErr := tryObject_setIvar(obj, ivar, value); callErr != nil {
		panic(callErr)
	}
}

var _object_setIvarWithStrongDefault func(obj Object, ivar Ivar, value Object)
var _object_setIvarWithStrongDefaultErr error

func tryObject_setIvarWithStrongDefault(obj Object, ivar Ivar, value Object) error {
	if _object_setIvarWithStrongDefault == nil {
		return symbolCallError("object_setIvarWithStrongDefault", "10.12", _object_setIvarWithStrongDefaultErr)
	}
	_object_setIvarWithStrongDefault(obj, ivar, value)
	return nil
}

// Object_setIvarWithStrongDefault.
//
// See: https://developer.apple.com/documentation/ObjectiveC/object_setIvarWithStrongDefault(_:_:_:)
func Object_setIvarWithStrongDefault(obj Object, ivar Ivar, value Object) {
	if callErr := tryObject_setIvarWithStrongDefault(obj, ivar, value); callErr != nil {
		panic(callErr)
	}
}

var _property_copyAttributeList func(property unsafe.Pointer, outCount *uint) *Objc_property_attribute_t
var _property_copyAttributeListErr error

func tryProperty_copyAttributeList(property unsafe.Pointer, outCount *uint) (*Objc_property_attribute_t, error) {
	if _property_copyAttributeList == nil {
		return nil, symbolCallError("property_copyAttributeList", "10.7", _property_copyAttributeListErr)
	}
	return _property_copyAttributeList(property, outCount), nil
}

// Property_copyAttributeList returns an array of property attributes for a given property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeList(_:_:)
func Property_copyAttributeList(property unsafe.Pointer, outCount *uint) *Objc_property_attribute_t {
	result, callErr := tryProperty_copyAttributeList(property, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _property_copyAttributeValue func(property unsafe.Pointer, attributeName string) *byte
var _property_copyAttributeValueErr error

func tryProperty_copyAttributeValue(property unsafe.Pointer, attributeName string) (*byte, error) {
	if _property_copyAttributeValue == nil {
		return nil, symbolCallError("property_copyAttributeValue", "10.7", _property_copyAttributeValueErr)
	}
	return _property_copyAttributeValue(property, attributeName), nil
}

// Property_copyAttributeValue returns the value of a property attribute given the attribute name.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_copyAttributeValue(_:_:)
func Property_copyAttributeValue(property unsafe.Pointer, attributeName string) *byte {
	result, callErr := tryProperty_copyAttributeValue(property, attributeName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _property_getAttributes func(property unsafe.Pointer) *byte
var _property_getAttributesErr error

func tryProperty_getAttributes(property unsafe.Pointer) (*byte, error) {
	if _property_getAttributes == nil {
		return nil, symbolCallError("property_getAttributes", "10.5", _property_getAttributesErr)
	}
	return _property_getAttributes(property), nil
}

// Property_getAttributes returns the attribute string of a property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_getAttributes(_:)
func Property_getAttributes(property unsafe.Pointer) *byte {
	result, callErr := tryProperty_getAttributes(property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _property_getName func(property unsafe.Pointer) *byte
var _property_getNameErr error

func tryProperty_getName(property unsafe.Pointer) (*byte, error) {
	if _property_getName == nil {
		return nil, symbolCallError("property_getName", "10.5", _property_getNameErr)
	}
	return _property_getName(property), nil
}

// Property_getName returns the name of a property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/property_getName(_:)
func Property_getName(property unsafe.Pointer) *byte {
	result, callErr := tryProperty_getName(property)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_addMethodDescription func(proto **Protocol, name SEL, types string, isRequiredMethod bool, isInstanceMethod bool)
var _protocol_addMethodDescriptionErr error

func tryProtocol_addMethodDescription(proto **Protocol, name SEL, types string, isRequiredMethod bool, isInstanceMethod bool) error {
	if _protocol_addMethodDescription == nil {
		return symbolCallError("protocol_addMethodDescription", "10.7", _protocol_addMethodDescriptionErr)
	}
	_protocol_addMethodDescription(proto, name, types, isRequiredMethod, isInstanceMethod)
	return nil
}

// Protocol_addMethodDescription adds a method to a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_addMethodDescription(_:_:_:_:_:)
func Protocol_addMethodDescription(proto **Protocol, name SEL, types string, isRequiredMethod bool, isInstanceMethod bool) {
	if callErr := tryProtocol_addMethodDescription(proto, name, types, isRequiredMethod, isInstanceMethod); callErr != nil {
		panic(callErr)
	}
}

var _protocol_addProperty func(proto **Protocol, name string, attributes *Objc_property_attribute_t, attributeCount uint, isRequiredProperty bool, isInstanceProperty bool)
var _protocol_addPropertyErr error

func tryProtocol_addProperty(proto **Protocol, name string, attributes *Objc_property_attribute_t, attributeCount uint, isRequiredProperty bool, isInstanceProperty bool) error {
	if _protocol_addProperty == nil {
		return symbolCallError("protocol_addProperty", "10.7", _protocol_addPropertyErr)
	}
	_protocol_addProperty(proto, name, attributes, attributeCount, isRequiredProperty, isInstanceProperty)
	return nil
}

// Protocol_addProperty adds a property to a protocol that is under construction.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_addProperty(_:_:_:_:_:_:)
func Protocol_addProperty(proto **Protocol, name string, attributes *Objc_property_attribute_t, attributeCount uint, isRequiredProperty bool, isInstanceProperty bool) {
	if callErr := tryProtocol_addProperty(proto, name, attributes, attributeCount, isRequiredProperty, isInstanceProperty); callErr != nil {
		panic(callErr)
	}
}

var _protocol_addProtocol func(proto **Protocol, addition **Protocol)
var _protocol_addProtocolErr error

func tryProtocol_addProtocol(proto **Protocol, addition **Protocol) error {
	if _protocol_addProtocol == nil {
		return symbolCallError("protocol_addProtocol", "10.7", _protocol_addProtocolErr)
	}
	_protocol_addProtocol(proto, addition)
	return nil
}

// Protocol_addProtocol adds a registered protocol to another protocol that is under construction.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_addProtocol(_:_:)
func Protocol_addProtocol(proto **Protocol, addition **Protocol) {
	if callErr := tryProtocol_addProtocol(proto, addition); callErr != nil {
		panic(callErr)
	}
}

var _protocol_conformsToProtocol func(proto **Protocol, other **Protocol) bool
var _protocol_conformsToProtocolErr error

func tryProtocol_conformsToProtocol(proto **Protocol, other **Protocol) (bool, error) {
	if _protocol_conformsToProtocol == nil {
		return false, symbolCallError("protocol_conformsToProtocol", "10.5", _protocol_conformsToProtocolErr)
	}
	return _protocol_conformsToProtocol(proto, other), nil
}

// Protocol_conformsToProtocol returns a Boolean value that indicates whether one protocol conforms to another protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_conformsToProtocol(_:_:)
func Protocol_conformsToProtocol(proto **Protocol, other **Protocol) bool {
	result, callErr := tryProtocol_conformsToProtocol(proto, other)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_copyMethodDescriptionList func(proto **Protocol, isRequiredMethod bool, isInstanceMethod bool, outCount *uint) *Objc_method_description
var _protocol_copyMethodDescriptionListErr error

func tryProtocol_copyMethodDescriptionList(proto **Protocol, isRequiredMethod bool, isInstanceMethod bool, outCount *uint) (*Objc_method_description, error) {
	if _protocol_copyMethodDescriptionList == nil {
		return nil, symbolCallError("protocol_copyMethodDescriptionList", "10.5", _protocol_copyMethodDescriptionListErr)
	}
	return _protocol_copyMethodDescriptionList(proto, isRequiredMethod, isInstanceMethod, outCount), nil
}

// Protocol_copyMethodDescriptionList returns an array of method descriptions of methods meeting a given specification for a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyMethodDescriptionList(_:_:_:_:)
func Protocol_copyMethodDescriptionList(proto **Protocol, isRequiredMethod bool, isInstanceMethod bool, outCount *uint) *Objc_method_description {
	result, callErr := tryProtocol_copyMethodDescriptionList(proto, isRequiredMethod, isInstanceMethod, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_copyPropertyList func(proto **Protocol, outCount *uint) *Objc_property_t
var _protocol_copyPropertyListErr error

func tryProtocol_copyPropertyList(proto **Protocol, outCount *uint) (*Objc_property_t, error) {
	if _protocol_copyPropertyList == nil {
		return nil, symbolCallError("protocol_copyPropertyList", "10.5", _protocol_copyPropertyListErr)
	}
	return _protocol_copyPropertyList(proto, outCount), nil
}

// Protocol_copyPropertyList returns an array of the properties declared by a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList(_:_:)
func Protocol_copyPropertyList(proto **Protocol, outCount *uint) *Objc_property_t {
	result, callErr := tryProtocol_copyPropertyList(proto, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_copyPropertyList2 func(proto **Protocol, outCount *uint, isRequiredProperty bool, isInstanceProperty bool) *Objc_property_t
var _protocol_copyPropertyList2Err error

func tryProtocol_copyPropertyList2(proto **Protocol, outCount *uint, isRequiredProperty bool, isInstanceProperty bool) (*Objc_property_t, error) {
	if _protocol_copyPropertyList2 == nil {
		return nil, symbolCallError("protocol_copyPropertyList2", "10.12", _protocol_copyPropertyList2Err)
	}
	return _protocol_copyPropertyList2(proto, outCount, isRequiredProperty, isInstanceProperty), nil
}

// Protocol_copyPropertyList2.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyPropertyList2(_:_:_:_:)
func Protocol_copyPropertyList2(proto **Protocol, outCount *uint, isRequiredProperty bool, isInstanceProperty bool) *Objc_property_t {
	result, callErr := tryProtocol_copyPropertyList2(proto, outCount, isRequiredProperty, isInstanceProperty)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_copyProtocolList func(proto **Protocol, outCount *uint) **Protocol
var _protocol_copyProtocolListErr error

func tryProtocol_copyProtocolList(proto **Protocol, outCount *uint) (**Protocol, error) {
	if _protocol_copyProtocolList == nil {
		return nil, symbolCallError("protocol_copyProtocolList", "10.5", _protocol_copyProtocolListErr)
	}
	return _protocol_copyProtocolList(proto, outCount), nil
}

// Protocol_copyProtocolList returns an array of the protocols adopted by a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_copyProtocolList(_:_:)
func Protocol_copyProtocolList(proto **Protocol, outCount *uint) **Protocol {
	result, callErr := tryProtocol_copyProtocolList(proto, outCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_getMethodDescription func(proto **Protocol, aSel SEL, isRequiredMethod bool, isInstanceMethod bool) unsafe.Pointer
var _protocol_getMethodDescriptionErr error

func tryProtocol_getMethodDescription(proto **Protocol, aSel SEL, isRequiredMethod bool, isInstanceMethod bool) (unsafe.Pointer, error) {
	if _protocol_getMethodDescription == nil {
		return nil, symbolCallError("protocol_getMethodDescription", "10.5", _protocol_getMethodDescriptionErr)
	}
	return _protocol_getMethodDescription(proto, aSel, isRequiredMethod, isInstanceMethod), nil
}

// Protocol_getMethodDescription returns a method description structure for a specified method of a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_getMethodDescription(_:_:_:_:)
func Protocol_getMethodDescription(proto **Protocol, aSel SEL, isRequiredMethod bool, isInstanceMethod bool) unsafe.Pointer {
	result, callErr := tryProtocol_getMethodDescription(proto, aSel, isRequiredMethod, isInstanceMethod)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_getName func(proto **Protocol) *byte
var _protocol_getNameErr error

func tryProtocol_getName(proto **Protocol) (*byte, error) {
	if _protocol_getName == nil {
		return nil, symbolCallError("protocol_getName", "10.5", _protocol_getNameErr)
	}
	return _protocol_getName(proto), nil
}

// Protocol_getName returns the name of a protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_getName(_:)
func Protocol_getName(proto **Protocol) *byte {
	result, callErr := tryProtocol_getName(proto)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_getProperty func(proto **Protocol, name string, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer
var _protocol_getPropertyErr error

func tryProtocol_getProperty(proto **Protocol, name string, isRequiredProperty bool, isInstanceProperty bool) (unsafe.Pointer, error) {
	if _protocol_getProperty == nil {
		return nil, symbolCallError("protocol_getProperty", "10.5", _protocol_getPropertyErr)
	}
	return _protocol_getProperty(proto, name, isRequiredProperty, isInstanceProperty), nil
}

// Protocol_getProperty returns the specified property of a given protocol.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_getProperty(_:_:_:_:)
func Protocol_getProperty(proto **Protocol, name string, isRequiredProperty bool, isInstanceProperty bool) unsafe.Pointer {
	result, callErr := tryProtocol_getProperty(proto, name, isRequiredProperty, isInstanceProperty)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _protocol_isEqual func(proto **Protocol, other **Protocol) bool
var _protocol_isEqualErr error

func tryProtocol_isEqual(proto **Protocol, other **Protocol) (bool, error) {
	if _protocol_isEqual == nil {
		return false, symbolCallError("protocol_isEqual", "10.5", _protocol_isEqualErr)
	}
	return _protocol_isEqual(proto, other), nil
}

// Protocol_isEqual returns a Boolean value that indicates whether two protocols are equal.
//
// See: https://developer.apple.com/documentation/ObjectiveC/protocol_isEqual(_:_:)
func Protocol_isEqual(proto **Protocol, other **Protocol) bool {
	result, callErr := tryProtocol_isEqual(proto, other)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sel_getName func(sel SEL) *byte
var _sel_getNameErr error

func trySel_getName(sel SEL) (*byte, error) {
	if _sel_getName == nil {
		return nil, symbolCallError("sel_getName", "10.0", _sel_getNameErr)
	}
	return _sel_getName(sel), nil
}

// Sel_getName returns the name of the method specified by a given selector.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_getName(_:)
func Sel_getName(sel SEL) *byte {
	result, callErr := trySel_getName(sel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sel_getUid func(str string) SEL
var _sel_getUidErr error

func trySel_getUid(str string) (SEL, error) {
	if _sel_getUid == nil {
		return *new(SEL), symbolCallError("sel_getUid", "10.0", _sel_getUidErr)
	}
	return _sel_getUid(str), nil
}

// Sel_getUid registers a method name with the Objective-C runtime system.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_getUid(_:)
func Sel_getUid(str string) SEL {
	result, callErr := trySel_getUid(str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sel_isEqual func(lhs SEL, rhs SEL) bool
var _sel_isEqualErr error

func trySel_isEqual(lhs SEL, rhs SEL) (bool, error) {
	if _sel_isEqual == nil {
		return false, symbolCallError("sel_isEqual", "10.5", _sel_isEqualErr)
	}
	return _sel_isEqual(lhs, rhs), nil
}

// Sel_isEqual returns a Boolean value that indicates whether two selectors are equal.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_isEqual(_:_:)
func Sel_isEqual(lhs SEL, rhs SEL) bool {
	result, callErr := trySel_isEqual(lhs, rhs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sel_isMapped func(sel SEL) bool
var _sel_isMappedErr error

func trySel_isMapped(sel SEL) (bool, error) {
	if _sel_isMapped == nil {
		return false, symbolCallError("sel_isMapped", "10.0", _sel_isMappedErr)
	}
	return _sel_isMapped(sel), nil
}

// Sel_isMapped identifies a selector as being valid or invalid.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_isMapped(_:)
func Sel_isMapped(sel SEL) bool {
	result, callErr := trySel_isMapped(sel)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _sel_registerName func(str string) SEL
var _sel_registerNameErr error

func trySel_registerName(str string) (SEL, error) {
	if _sel_registerName == nil {
		return *new(SEL), symbolCallError("sel_registerName", "10.0", _sel_registerNameErr)
	}
	return _sel_registerName(str), nil
}

// Sel_registerName registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value.
//
// See: https://developer.apple.com/documentation/ObjectiveC/sel_registerName(_:)
func Sel_registerName(str string) SEL {
	result, callErr := trySel_registerName(str)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_nXCompareHashTables, &_nXCompareHashTablesErr, frameworkHandle, "NXCompareHashTables", "10.0")
	registerFunc(&_nXCopyHashTable, &_nXCopyHashTableErr, frameworkHandle, "NXCopyHashTable", "10.0")
	registerFunc(&_nXCountHashTable, &_nXCountHashTableErr, frameworkHandle, "NXCountHashTable", "10.0")
	registerFunc(&_nXCreateHashTable, &_nXCreateHashTableErr, frameworkHandle, "NXCreateHashTable", "10.0")
	registerFunc(&_nXCreateHashTableFromZone, &_nXCreateHashTableFromZoneErr, frameworkHandle, "NXCreateHashTableFromZone", "10.0")
	registerFunc(&_nXEmptyHashTable, &_nXEmptyHashTableErr, frameworkHandle, "NXEmptyHashTable", "10.0")
	registerFunc(&_nXFreeHashTable, &_nXFreeHashTableErr, frameworkHandle, "NXFreeHashTable", "10.0")
	registerFunc(&_nXHashGet, &_nXHashGetErr, frameworkHandle, "NXHashGet", "10.0")
	registerFunc(&_nXHashInsert, &_nXHashInsertErr, frameworkHandle, "NXHashInsert", "10.0")
	registerFunc(&_nXHashInsertIfAbsent, &_nXHashInsertIfAbsentErr, frameworkHandle, "NXHashInsertIfAbsent", "10.0")
	registerFunc(&_nXHashMember, &_nXHashMemberErr, frameworkHandle, "NXHashMember", "10.0")
	registerFunc(&_nXHashRemove, &_nXHashRemoveErr, frameworkHandle, "NXHashRemove", "10.0")
	registerFunc(&_nXInitHashState, &_nXInitHashStateErr, frameworkHandle, "NXInitHashState", "10.0")
	registerFunc(&_nXNextHashState, &_nXNextHashStateErr, frameworkHandle, "NXNextHashState", "10.0")
	registerFunc(&_nXNoEffectFree, &_nXNoEffectFreeErr, frameworkHandle, "NXNoEffectFree", "10.0")
	registerFunc(&_nXPtrHash, &_nXPtrHashErr, frameworkHandle, "NXPtrHash", "10.0")
	registerFunc(&_nXPtrIsEqual, &_nXPtrIsEqualErr, frameworkHandle, "NXPtrIsEqual", "10.0")
	registerFunc(&_nXReallyFree, &_nXReallyFreeErr, frameworkHandle, "NXReallyFree", "10.0")
	registerFunc(&_nXResetHashTable, &_nXResetHashTableErr, frameworkHandle, "NXResetHashTable", "10.0")
	registerFunc(&_nXStrHash, &_nXStrHashErr, frameworkHandle, "NXStrHash", "10.0")
	registerFunc(&_nXStrIsEqual, &_nXStrIsEqualErr, frameworkHandle, "NXStrIsEqual", "10.0")
	registerFunc(&_class_addIvar, &_class_addIvarErr, frameworkHandle, "class_addIvar", "10.5")
	registerFunc(&_class_addMethod, &_class_addMethodErr, frameworkHandle, "class_addMethod", "10.5")
	registerFunc(&_class_addProperty, &_class_addPropertyErr, frameworkHandle, "class_addProperty", "10.7")
	registerFunc(&_class_addProtocol, &_class_addProtocolErr, frameworkHandle, "class_addProtocol", "10.5")
	registerFunc(&_class_conformsToProtocol, &_class_conformsToProtocolErr, frameworkHandle, "class_conformsToProtocol", "10.5")
	registerFunc(&_class_copyIvarList, &_class_copyIvarListErr, frameworkHandle, "class_copyIvarList", "10.5")
	registerFunc(&_class_copyMethodList, &_class_copyMethodListErr, frameworkHandle, "class_copyMethodList", "10.5")
	registerFunc(&_class_copyPropertyList, &_class_copyPropertyListErr, frameworkHandle, "class_copyPropertyList", "10.5")
	registerFunc(&_class_copyProtocolList, &_class_copyProtocolListErr, frameworkHandle, "class_copyProtocolList", "10.5")
	registerFunc(&_class_createInstance, &_class_createInstanceErr, frameworkHandle, "class_createInstance", "10.0")
	registerFunc(&_class_createInstanceFromZone, &_class_createInstanceFromZoneErr, frameworkHandle, "class_createInstanceFromZone", "10.0")
	registerFunc(&_class_getClassMethod, &_class_getClassMethodErr, frameworkHandle, "class_getClassMethod", "10.0")
	registerFunc(&_class_getClassVariable, &_class_getClassVariableErr, frameworkHandle, "class_getClassVariable", "10.5")
	registerFunc(&_class_getImageName, &_class_getImageNameErr, frameworkHandle, "class_getImageName", "10.5")
	registerFunc(&_class_getInstanceMethod, &_class_getInstanceMethodErr, frameworkHandle, "class_getInstanceMethod", "10.0")
	registerFunc(&_class_getInstanceSize, &_class_getInstanceSizeErr, frameworkHandle, "class_getInstanceSize", "10.5")
	registerFunc(&_class_getInstanceVariable, &_class_getInstanceVariableErr, frameworkHandle, "class_getInstanceVariable", "10.0")
	registerFunc(&_class_getIvarLayout, &_class_getIvarLayoutErr, frameworkHandle, "class_getIvarLayout", "10.5")
	registerFunc(&_class_getMethodImplementation, &_class_getMethodImplementationErr, frameworkHandle, "class_getMethodImplementation", "10.5")
	registerFunc(&_class_getName, &_class_getNameErr, frameworkHandle, "class_getName", "10.5")
	registerFunc(&_class_getProperty, &_class_getPropertyErr, frameworkHandle, "class_getProperty", "10.5")
	registerFunc(&_class_getSuperclass, &_class_getSuperclassErr, frameworkHandle, "class_getSuperclass", "10.5")
	registerFunc(&_class_getVersion, &_class_getVersionErr, frameworkHandle, "class_getVersion", "10.0")
	registerFunc(&_class_getWeakIvarLayout, &_class_getWeakIvarLayoutErr, frameworkHandle, "class_getWeakIvarLayout", "10.5")
	registerFunc(&_class_isMetaClass, &_class_isMetaClassErr, frameworkHandle, "class_isMetaClass", "10.5")
	registerFunc(&_class_replaceMethod, &_class_replaceMethodErr, frameworkHandle, "class_replaceMethod", "10.5")
	registerFunc(&_class_replaceProperty, &_class_replacePropertyErr, frameworkHandle, "class_replaceProperty", "10.7")
	registerFunc(&_class_respondsToSelector, &_class_respondsToSelectorErr, frameworkHandle, "class_respondsToSelector", "10.5")
	registerFunc(&_class_setIvarLayout, &_class_setIvarLayoutErr, frameworkHandle, "class_setIvarLayout", "10.5")
	registerFunc(&_class_setVersion, &_class_setVersionErr, frameworkHandle, "class_setVersion", "10.0")
	registerFunc(&_class_setWeakIvarLayout, &_class_setWeakIvarLayoutErr, frameworkHandle, "class_setWeakIvarLayout", "10.5")
	registerFunc(&_imp_getBlock, &_imp_getBlockErr, frameworkHandle, "imp_getBlock", "10.7")
	registerFunc(&_imp_implementationWithBlock, &_imp_implementationWithBlockErr, frameworkHandle, "imp_implementationWithBlock", "10.7")
	registerFunc(&_imp_removeBlock, &_imp_removeBlockErr, frameworkHandle, "imp_removeBlock", "10.7")
	registerFunc(&_ivar_getName, &_ivar_getNameErr, frameworkHandle, "ivar_getName", "10.5")
	registerFunc(&_ivar_getOffset, &_ivar_getOffsetErr, frameworkHandle, "ivar_getOffset", "10.5")
	registerFunc(&_ivar_getTypeEncoding, &_ivar_getTypeEncodingErr, frameworkHandle, "ivar_getTypeEncoding", "10.5")
	registerFunc(&_method_copyArgumentType, &_method_copyArgumentTypeErr, frameworkHandle, "method_copyArgumentType", "10.5")
	registerFunc(&_method_copyReturnType, &_method_copyReturnTypeErr, frameworkHandle, "method_copyReturnType", "10.5")
	registerFunc(&_method_exchangeImplementations, &_method_exchangeImplementationsErr, frameworkHandle, "method_exchangeImplementations", "10.5")
	registerFunc(&_method_getArgumentType, &_method_getArgumentTypeErr, frameworkHandle, "method_getArgumentType", "10.5")
	registerFunc(&_method_getDescription, &_method_getDescriptionErr, frameworkHandle, "method_getDescription", "10.5")
	registerFunc(&_method_getImplementation, &_method_getImplementationErr, frameworkHandle, "method_getImplementation", "10.5")
	registerFunc(&_method_getName, &_method_getNameErr, frameworkHandle, "method_getName", "10.5")
	registerFunc(&_method_getNumberOfArguments, &_method_getNumberOfArgumentsErr, frameworkHandle, "method_getNumberOfArguments", "10.0")
	registerFunc(&_method_getReturnType, &_method_getReturnTypeErr, frameworkHandle, "method_getReturnType", "10.5")
	registerFunc(&_method_getTypeEncoding, &_method_getTypeEncodingErr, frameworkHandle, "method_getTypeEncoding", "10.5")
	registerFunc(&_method_invoke, &_method_invokeErr, frameworkHandle, "method_invoke", "10.5")
	registerFunc(&_method_setImplementation, &_method_setImplementationErr, frameworkHandle, "method_setImplementation", "10.5")
	registerFunc(&_objc_addExceptionHandler, &_objc_addExceptionHandlerErr, frameworkHandle, "objc_addExceptionHandler", "10.5")
	registerFunc(&_objc_addLoadImageFunc, &_objc_addLoadImageFuncErr, frameworkHandle, "objc_addLoadImageFunc", "10.15")
	registerFunc(&_objc_allocateClassPair, &_objc_allocateClassPairErr, frameworkHandle, "objc_allocateClassPair", "10.5")
	registerFunc(&_objc_allocateProtocol, &_objc_allocateProtocolErr, frameworkHandle, "objc_allocateProtocol", "10.7")
	registerFunc(&_objc_begin_catch, &_objc_begin_catchErr, frameworkHandle, "objc_begin_catch", "10.5")
	registerFunc(&_objc_constructInstance, &_objc_constructInstanceErr, frameworkHandle, "objc_constructInstance", "10.6")
	registerFunc(&_objc_copyClassList, &_objc_copyClassListErr, frameworkHandle, "objc_copyClassList", "10.7")
	registerFunc(&_objc_copyClassNamesForImage, &_objc_copyClassNamesForImageErr, frameworkHandle, "objc_copyClassNamesForImage", "10.5")
	registerFunc(&_objc_copyImageNames, &_objc_copyImageNamesErr, frameworkHandle, "objc_copyImageNames", "10.5")
	registerFunc(&_objc_copyProtocolList, &_objc_copyProtocolListErr, frameworkHandle, "objc_copyProtocolList", "10.5")
	registerFunc(&_objc_destructInstance, &_objc_destructInstanceErr, frameworkHandle, "objc_destructInstance", "10.6")
	registerFunc(&_objc_disposeClassPair, &_objc_disposeClassPairErr, frameworkHandle, "objc_disposeClassPair", "10.5")
	registerFunc(&_objc_duplicateClass, &_objc_duplicateClassErr, frameworkHandle, "objc_duplicateClass", "10.5")
	registerFunc(&_objc_end_catch, &_objc_end_catchErr, frameworkHandle, "objc_end_catch", "10.5")
	registerFunc(&_objc_enumerateClasses, &_objc_enumerateClassesErr, frameworkHandle, "objc_enumerateClasses", "13.0")
	registerFunc(&_objc_enumerationMutation, &_objc_enumerationMutationErr, frameworkHandle, "objc_enumerationMutation", "10.5")
	registerFunc(&_objc_exception_rethrow, &_objc_exception_rethrowErr, frameworkHandle, "objc_exception_rethrow", "10.5")
	registerFunc(&_objc_exception_throw, &_objc_exception_throwErr, frameworkHandle, "objc_exception_throw", "10.5")
	registerFunc(&_objc_getAssociatedObject, &_objc_getAssociatedObjectErr, frameworkHandle, "objc_getAssociatedObject", "10.6")
	registerFunc(&_objc_getClass, &_objc_getClassErr, frameworkHandle, "objc_getClass", "10.0")
	registerFunc(&_objc_getClassList, &_objc_getClassListErr, frameworkHandle, "objc_getClassList", "10.0")
	registerFunc(&_objc_getFutureClass, &_objc_getFutureClassErr, frameworkHandle, "objc_getFutureClass", "10.5")
	registerFunc(&_objc_getMetaClass, &_objc_getMetaClassErr, frameworkHandle, "objc_getMetaClass", "10.0")
	registerFunc(&_objc_getProtocol, &_objc_getProtocolErr, frameworkHandle, "objc_getProtocol", "10.5")
	registerFunc(&_objc_getRequiredClass, &_objc_getRequiredClassErr, frameworkHandle, "objc_getRequiredClass", "10.0")
	registerFunc(&_objc_loadWeak, &_objc_loadWeakErr, frameworkHandle, "objc_loadWeak", "10.7")
	registerFunc(&_objc_lookUpClass, &_objc_lookUpClassErr, frameworkHandle, "objc_lookUpClass", "10.0")
	registerFunc(&_objc_msgSend, &_objc_msgSendErr, frameworkHandle, "objc_msgSend", "10.0")
	registerFunc(&_objc_msgSendSuper, &_objc_msgSendSuperErr, frameworkHandle, "objc_msgSendSuper", "10.0")
	registerFunc(&_objc_registerClassPair, &_objc_registerClassPairErr, frameworkHandle, "objc_registerClassPair", "10.5")
	registerFunc(&_objc_registerProtocol, &_objc_registerProtocolErr, frameworkHandle, "objc_registerProtocol", "10.7")
	registerFunc(&_objc_removeAssociatedObjects, &_objc_removeAssociatedObjectsErr, frameworkHandle, "objc_removeAssociatedObjects", "10.6")
	registerFunc(&_objc_removeExceptionHandler, &_objc_removeExceptionHandlerErr, frameworkHandle, "objc_removeExceptionHandler", "10.5")
	registerFunc(&_objc_setAssociatedObject, &_objc_setAssociatedObjectErr, frameworkHandle, "objc_setAssociatedObject", "10.6")
	registerFunc(&_objc_setEnumerationMutationHandler, &_objc_setEnumerationMutationHandlerErr, frameworkHandle, "objc_setEnumerationMutationHandler", "10.5")
	registerFunc(&_objc_setExceptionMatcher, &_objc_setExceptionMatcherErr, frameworkHandle, "objc_setExceptionMatcher", "10.5")
	registerFunc(&_objc_setExceptionPreprocessor, &_objc_setExceptionPreprocessorErr, frameworkHandle, "objc_setExceptionPreprocessor", "10.5")
	registerFunc(&_objc_setForwardHandler, &_objc_setForwardHandlerErr, frameworkHandle, "objc_setForwardHandler", "10.5")
	registerFunc(&_objc_setHook_getClass, &_objc_setHook_getClassErr, frameworkHandle, "objc_setHook_getClass", "10.14.4")
	registerFunc(&_objc_setHook_getImageName, &_objc_setHook_getImageNameErr, frameworkHandle, "objc_setHook_getImageName", "10.14")
	registerFunc(&_objc_setHook_lazyClassNamer, &_objc_setHook_lazyClassNamerErr, frameworkHandle, "objc_setHook_lazyClassNamer", "11.0")
	registerFunc(&_objc_setUncaughtExceptionHandler, &_objc_setUncaughtExceptionHandlerErr, frameworkHandle, "objc_setUncaughtExceptionHandler", "10.5")
	registerFunc(&_objc_storeWeak, &_objc_storeWeakErr, frameworkHandle, "objc_storeWeak", "10.7")
	registerFunc(&_objc_sync_enter, &_objc_sync_enterErr, frameworkHandle, "objc_sync_enter", "10.3")
	registerFunc(&_objc_sync_exit, &_objc_sync_exitErr, frameworkHandle, "objc_sync_exit", "10.3")
	registerFunc(&_objc_terminate, &_objc_terminateErr, frameworkHandle, "objc_terminate", "10.8")
	registerFunc(&_object_copy, &_object_copyErr, frameworkHandle, "object_copy", "10.0")
	registerFunc(&_object_copyFromZone, &_object_copyFromZoneErr, frameworkHandle, "object_copyFromZone", "10.0")
	registerFunc(&_object_dispose, &_object_disposeErr, frameworkHandle, "object_dispose", "10.0")
	registerFunc(&_object_getClass, &_object_getClassErr, frameworkHandle, "object_getClass", "10.5")
	registerFunc(&_object_getClassName, &_object_getClassNameErr, frameworkHandle, "object_getClassName", "10.0")
	registerFunc(&_object_getIndexedIvars, &_object_getIndexedIvarsErr, frameworkHandle, "object_getIndexedIvars", "10.0")
	registerFunc(&_object_getInstanceVariable, &_object_getInstanceVariableErr, frameworkHandle, "object_getInstanceVariable", "10.0")
	registerFunc(&_object_getIvar, &_object_getIvarErr, frameworkHandle, "object_getIvar", "10.5")
	registerFunc(&_object_isClass, &_object_isClassErr, frameworkHandle, "object_isClass", "10.10")
	registerFunc(&_object_setClass, &_object_setClassErr, frameworkHandle, "object_setClass", "10.5")
	registerFunc(&_object_setInstanceVariable, &_object_setInstanceVariableErr, frameworkHandle, "object_setInstanceVariable", "10.0")
	registerFunc(&_object_setInstanceVariableWithStrongDefault, &_object_setInstanceVariableWithStrongDefaultErr, frameworkHandle, "object_setInstanceVariableWithStrongDefault", "10.12")
	registerFunc(&_object_setIvar, &_object_setIvarErr, frameworkHandle, "object_setIvar", "10.5")
	registerFunc(&_object_setIvarWithStrongDefault, &_object_setIvarWithStrongDefaultErr, frameworkHandle, "object_setIvarWithStrongDefault", "10.12")
	registerFunc(&_property_copyAttributeList, &_property_copyAttributeListErr, frameworkHandle, "property_copyAttributeList", "10.7")
	registerFunc(&_property_copyAttributeValue, &_property_copyAttributeValueErr, frameworkHandle, "property_copyAttributeValue", "10.7")
	registerFunc(&_property_getAttributes, &_property_getAttributesErr, frameworkHandle, "property_getAttributes", "10.5")
	registerFunc(&_property_getName, &_property_getNameErr, frameworkHandle, "property_getName", "10.5")
	registerFunc(&_protocol_addMethodDescription, &_protocol_addMethodDescriptionErr, frameworkHandle, "protocol_addMethodDescription", "10.7")
	registerFunc(&_protocol_addProperty, &_protocol_addPropertyErr, frameworkHandle, "protocol_addProperty", "10.7")
	registerFunc(&_protocol_addProtocol, &_protocol_addProtocolErr, frameworkHandle, "protocol_addProtocol", "10.7")
	registerFunc(&_protocol_conformsToProtocol, &_protocol_conformsToProtocolErr, frameworkHandle, "protocol_conformsToProtocol", "10.5")
	registerFunc(&_protocol_copyMethodDescriptionList, &_protocol_copyMethodDescriptionListErr, frameworkHandle, "protocol_copyMethodDescriptionList", "10.5")
	registerFunc(&_protocol_copyPropertyList, &_protocol_copyPropertyListErr, frameworkHandle, "protocol_copyPropertyList", "10.5")
	registerFunc(&_protocol_copyPropertyList2, &_protocol_copyPropertyList2Err, frameworkHandle, "protocol_copyPropertyList2", "10.12")
	registerFunc(&_protocol_copyProtocolList, &_protocol_copyProtocolListErr, frameworkHandle, "protocol_copyProtocolList", "10.5")
	registerFunc(&_protocol_getMethodDescription, &_protocol_getMethodDescriptionErr, frameworkHandle, "protocol_getMethodDescription", "10.5")
	registerFunc(&_protocol_getName, &_protocol_getNameErr, frameworkHandle, "protocol_getName", "10.5")
	registerFunc(&_protocol_getProperty, &_protocol_getPropertyErr, frameworkHandle, "protocol_getProperty", "10.5")
	registerFunc(&_protocol_isEqual, &_protocol_isEqualErr, frameworkHandle, "protocol_isEqual", "10.5")
	registerFunc(&_sel_getName, &_sel_getNameErr, frameworkHandle, "sel_getName", "10.0")
	registerFunc(&_sel_getUid, &_sel_getUidErr, frameworkHandle, "sel_getUid", "10.0")
	registerFunc(&_sel_isEqual, &_sel_isEqualErr, frameworkHandle, "sel_isEqual", "10.5")
	registerFunc(&_sel_isMapped, &_sel_isMappedErr, frameworkHandle, "sel_isMapped", "10.0")
	registerFunc(&_sel_registerName, &_sel_registerNameErr, frameworkHandle, "sel_registerName", "10.0")
}
