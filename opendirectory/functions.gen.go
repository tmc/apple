// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

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
		return fmt.Sprintf("OpenDirectory: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("OpenDirectory: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("OpenDirectory: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("OpenDirectory: register symbol %s: %v", name, r)
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

var _oDContextGetTypeID func() uint
var _oDContextGetTypeIDErr error

func tryODContextGetTypeID() (uint, error) {
	if _oDContextGetTypeID == nil {
		return 0, symbolCallError("ODContextGetTypeID", "", _oDContextGetTypeIDErr)
	}
	return _oDContextGetTypeID(), nil
}

// ODContextGetTypeID returns the type ID for the Open Directory context.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODContextGetTypeID()
func ODContextGetTypeID() uint {
	result, callErr := tryODContextGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeAddAccountPolicy func(node ODNodeRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool
var _oDNodeAddAccountPolicyErr error

func tryODNodeAddAccountPolicy(node ODNodeRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeAddAccountPolicy == nil {
		return false, symbolCallError("ODNodeAddAccountPolicy", "10.10", _oDNodeAddAccountPolicyErr)
	}
	return _oDNodeAddAccountPolicy(node, policy, category, err), nil
}

// ODNodeAddAccountPolicy.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeAddAccountPolicy(_:_:_:_:)
func ODNodeAddAccountPolicy(node ODNodeRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeAddAccountPolicy(node, policy, category, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopyAccountPolicies func(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDNodeCopyAccountPoliciesErr error

func tryODNodeCopyAccountPolicies(node ODNodeRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDNodeCopyAccountPolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODNodeCopyAccountPolicies", "10.10", _oDNodeCopyAccountPoliciesErr)
	}
	return _oDNodeCopyAccountPolicies(node, err), nil
}

// ODNodeCopyAccountPolicies.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyAccountPolicies(_:_:)
func ODNodeCopyAccountPolicies(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODNodeCopyAccountPolicies(node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopyDetails func(node ODNodeRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDNodeCopyDetailsErr error

func tryODNodeCopyDetails(node ODNodeRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDNodeCopyDetails == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODNodeCopyDetails", "10.6", _oDNodeCopyDetailsErr)
	}
	return _oDNodeCopyDetails(node, keys, err), nil
}

// ODNodeCopyDetails returns a dictionary containing details about a node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyDetails(_:_:_:)
func ODNodeCopyDetails(node ODNodeRef, keys corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODNodeCopyDetails(node, keys, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopyPolicies func(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDNodeCopyPoliciesErr error

func tryODNodeCopyPolicies(node ODNodeRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDNodeCopyPolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODNodeCopyPolicies", "10.9", _oDNodeCopyPoliciesErr)
	}
	return _oDNodeCopyPolicies(node, err), nil
}

// ODNodeCopyPolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyPolicies(_:_:)
func ODNodeCopyPolicies(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODNodeCopyPolicies(node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopyRecord func(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, attributes corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) ODRecordRef
var _oDNodeCopyRecordErr error

func tryODNodeCopyRecord(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, attributes corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (ODRecordRef, error) {
	if _oDNodeCopyRecord == nil {
		return *new(ODRecordRef), symbolCallError("ODNodeCopyRecord", "10.6", _oDNodeCopyRecordErr)
	}
	return _oDNodeCopyRecord(node, recordType, recordName, attributes, err), nil
}

// ODNodeCopyRecord returns a reference to a record of a node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyRecord(_:_:_:_:_:)
func ODNodeCopyRecord(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, attributes corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) ODRecordRef {
	result, callErr := tryODNodeCopyRecord(node, recordType, recordName, attributes, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopySubnodeNames func(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDNodeCopySubnodeNamesErr error

func tryODNodeCopySubnodeNames(node ODNodeRef, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDNodeCopySubnodeNames == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODNodeCopySubnodeNames", "10.6", _oDNodeCopySubnodeNamesErr)
	}
	return _oDNodeCopySubnodeNames(node, err), nil
}

// ODNodeCopySubnodeNames returns the names of subnodes for a given node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySubnodeNames(_:_:)
func ODNodeCopySubnodeNames(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODNodeCopySubnodeNames(node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopySupportedAttributes func(node ODNodeRef, recordType unsafe.Pointer, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDNodeCopySupportedAttributesErr error

func tryODNodeCopySupportedAttributes(node ODNodeRef, recordType unsafe.Pointer, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDNodeCopySupportedAttributes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODNodeCopySupportedAttributes", "10.6", _oDNodeCopySupportedAttributesErr)
	}
	return _oDNodeCopySupportedAttributes(node, recordType, err), nil
}

// ODNodeCopySupportedAttributes returns an array of attribute types supported by a given node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySupportedAttributes(_:_:_:)
func ODNodeCopySupportedAttributes(node ODNodeRef, recordType unsafe.Pointer, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODNodeCopySupportedAttributes(node, recordType, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopySupportedPolicies func(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDNodeCopySupportedPoliciesErr error

func tryODNodeCopySupportedPolicies(node ODNodeRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDNodeCopySupportedPolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODNodeCopySupportedPolicies", "10.9", _oDNodeCopySupportedPoliciesErr)
	}
	return _oDNodeCopySupportedPolicies(node, err), nil
}

// ODNodeCopySupportedPolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySupportedPolicies(_:_:)
func ODNodeCopySupportedPolicies(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODNodeCopySupportedPolicies(node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopySupportedRecordTypes func(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDNodeCopySupportedRecordTypesErr error

func tryODNodeCopySupportedRecordTypes(node ODNodeRef, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDNodeCopySupportedRecordTypes == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODNodeCopySupportedRecordTypes", "10.6", _oDNodeCopySupportedRecordTypesErr)
	}
	return _oDNodeCopySupportedRecordTypes(node, err), nil
}

// ODNodeCopySupportedRecordTypes returns an array of the record types supported by a given node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopySupportedRecordTypes(_:_:)
func ODNodeCopySupportedRecordTypes(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODNodeCopySupportedRecordTypes(node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCopyUnreachableSubnodeNames func(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDNodeCopyUnreachableSubnodeNamesErr error

func tryODNodeCopyUnreachableSubnodeNames(node ODNodeRef, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDNodeCopyUnreachableSubnodeNames == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODNodeCopyUnreachableSubnodeNames", "10.6", _oDNodeCopyUnreachableSubnodeNamesErr)
	}
	return _oDNodeCopyUnreachableSubnodeNames(node, err), nil
}

// ODNodeCopyUnreachableSubnodeNames returns an array of the subnodes of a given node that are currently unreachable.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCopyUnreachableSubnodeNames(_:_:)
func ODNodeCopyUnreachableSubnodeNames(node ODNodeRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODNodeCopyUnreachableSubnodeNames(node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCreateCopy func(allocator corefoundation.CFAllocatorRef, node ODNodeRef, err *corefoundation.CFErrorRef) ODNodeRef
var _oDNodeCreateCopyErr error

func tryODNodeCreateCopy(allocator corefoundation.CFAllocatorRef, node ODNodeRef, err *corefoundation.CFErrorRef) (ODNodeRef, error) {
	if _oDNodeCreateCopy == nil {
		return *new(ODNodeRef), symbolCallError("ODNodeCreateCopy", "10.6", _oDNodeCreateCopyErr)
	}
	return _oDNodeCreateCopy(allocator, node, err), nil
}

// ODNodeCreateCopy returns a copy of an existing node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateCopy(_:_:_:)
func ODNodeCreateCopy(allocator corefoundation.CFAllocatorRef, node ODNodeRef, err *corefoundation.CFErrorRef) ODNodeRef {
	result, callErr := tryODNodeCreateCopy(allocator, node, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCreateRecord func(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, attributeDict corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) ODRecordRef
var _oDNodeCreateRecordErr error

func tryODNodeCreateRecord(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, attributeDict corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (ODRecordRef, error) {
	if _oDNodeCreateRecord == nil {
		return *new(ODRecordRef), symbolCallError("ODNodeCreateRecord", "10.6", _oDNodeCreateRecordErr)
	}
	return _oDNodeCreateRecord(node, recordType, recordName, attributeDict, err), nil
}

// ODNodeCreateRecord creates a record in a specified node with specified properties.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateRecord(_:_:_:_:_:)
func ODNodeCreateRecord(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, attributeDict corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) ODRecordRef {
	result, callErr := tryODNodeCreateRecord(node, recordType, recordName, attributeDict, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCreateWithName func(allocator corefoundation.CFAllocatorRef, session ODSessionRef, nodeName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) ODNodeRef
var _oDNodeCreateWithNameErr error

func tryODNodeCreateWithName(allocator corefoundation.CFAllocatorRef, session ODSessionRef, nodeName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (ODNodeRef, error) {
	if _oDNodeCreateWithName == nil {
		return *new(ODNodeRef), symbolCallError("ODNodeCreateWithName", "10.6", _oDNodeCreateWithNameErr)
	}
	return _oDNodeCreateWithName(allocator, session, nodeName, err), nil
}

// ODNodeCreateWithName returns a new node created with a specified name.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateWithName(_:_:_:_:)
func ODNodeCreateWithName(allocator corefoundation.CFAllocatorRef, session ODSessionRef, nodeName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) ODNodeRef {
	result, callErr := tryODNodeCreateWithName(allocator, session, nodeName, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCreateWithNodeType func(allocator corefoundation.CFAllocatorRef, session ODSessionRef, nodeType ODNodeType, err *corefoundation.CFErrorRef) ODNodeRef
var _oDNodeCreateWithNodeTypeErr error

func tryODNodeCreateWithNodeType(allocator corefoundation.CFAllocatorRef, session ODSessionRef, nodeType ODNodeType, err *corefoundation.CFErrorRef) (ODNodeRef, error) {
	if _oDNodeCreateWithNodeType == nil {
		return *new(ODNodeRef), symbolCallError("ODNodeCreateWithNodeType", "10.6", _oDNodeCreateWithNodeTypeErr)
	}
	return _oDNodeCreateWithNodeType(allocator, session, nodeType, err), nil
}

// ODNodeCreateWithNodeType returns a new node created with a specified type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCreateWithNodeType(_:_:_:_:)
func ODNodeCreateWithNodeType(allocator corefoundation.CFAllocatorRef, session ODSessionRef, nodeType ODNodeType, err *corefoundation.CFErrorRef) ODNodeRef {
	result, callErr := tryODNodeCreateWithNodeType(allocator, session, nodeType, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCustomCall func(node ODNodeRef, customCode int, data corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef
var _oDNodeCustomCallErr error

func tryODNodeCustomCall(node ODNodeRef, customCode int, data corefoundation.CFDataRef, err *corefoundation.CFErrorRef) (corefoundation.CFDataRef, error) {
	if _oDNodeCustomCall == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("ODNodeCustomCall", "10.6", _oDNodeCustomCallErr)
	}
	return _oDNodeCustomCall(node, customCode, data, err), nil
}

// ODNodeCustomCall returns the result of a custom call to a node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCustomCall(_:_:_:_:)
func ODNodeCustomCall(node ODNodeRef, customCode int, data corefoundation.CFDataRef, err *corefoundation.CFErrorRef) corefoundation.CFDataRef {
	result, callErr := tryODNodeCustomCall(node, customCode, data, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeCustomFunction func(node ODNodeRef, function corefoundation.CFStringRef, payload corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef
var _oDNodeCustomFunctionErr error

func tryODNodeCustomFunction(node ODNodeRef, function corefoundation.CFStringRef, payload corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (corefoundation.CFTypeRef, error) {
	if _oDNodeCustomFunction == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("ODNodeCustomFunction", "10.9", _oDNodeCustomFunctionErr)
	}
	return _oDNodeCustomFunction(node, function, payload, err), nil
}

// ODNodeCustomFunction.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeCustomFunction(_:_:_:_:)
func ODNodeCustomFunction(node ODNodeRef, function corefoundation.CFStringRef, payload corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) corefoundation.CFTypeRef {
	result, callErr := tryODNodeCustomFunction(node, function, payload, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeGetName func(node ODNodeRef) corefoundation.CFStringRef
var _oDNodeGetNameErr error

func tryODNodeGetName(node ODNodeRef) (corefoundation.CFStringRef, error) {
	if _oDNodeGetName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("ODNodeGetName", "10.6", _oDNodeGetNameErr)
	}
	return _oDNodeGetName(node), nil
}

// ODNodeGetName returns the name of a node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeGetName(_:)
func ODNodeGetName(node ODNodeRef) corefoundation.CFStringRef {
	result, callErr := tryODNodeGetName(node)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeGetTypeID func() uint
var _oDNodeGetTypeIDErr error

func tryODNodeGetTypeID() (uint, error) {
	if _oDNodeGetTypeID == nil {
		return 0, symbolCallError("ODNodeGetTypeID", "10.6", _oDNodeGetTypeIDErr)
	}
	return _oDNodeGetTypeID(), nil
}

// ODNodeGetTypeID returns the type ID for an Open Directory node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeGetTypeID()
func ODNodeGetTypeID() uint {
	result, callErr := tryODNodeGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodePasswordContentCheck func(node ODNodeRef, password corefoundation.CFStringRef, recordName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDNodePasswordContentCheckErr error

func tryODNodePasswordContentCheck(node ODNodeRef, password corefoundation.CFStringRef, recordName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodePasswordContentCheck == nil {
		return false, symbolCallError("ODNodePasswordContentCheck", "10.10", _oDNodePasswordContentCheckErr)
	}
	return _oDNodePasswordContentCheck(node, password, recordName, err), nil
}

// ODNodePasswordContentCheck.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodePasswordContentCheck(_:_:_:_:)
func ODNodePasswordContentCheck(node ODNodeRef, password corefoundation.CFStringRef, recordName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodePasswordContentCheck(node, password, recordName, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeRemoveAccountPolicy func(node ODNodeRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool
var _oDNodeRemoveAccountPolicyErr error

func tryODNodeRemoveAccountPolicy(node ODNodeRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeRemoveAccountPolicy == nil {
		return false, symbolCallError("ODNodeRemoveAccountPolicy", "10.10", _oDNodeRemoveAccountPolicyErr)
	}
	return _oDNodeRemoveAccountPolicy(node, policy, category, err), nil
}

// ODNodeRemoveAccountPolicy.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeRemoveAccountPolicy(_:_:_:_:)
func ODNodeRemoveAccountPolicy(node ODNodeRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeRemoveAccountPolicy(node, policy, category, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeRemovePolicy func(node ODNodeRef, policyType ODPolicyType, err *corefoundation.CFErrorRef) bool
var _oDNodeRemovePolicyErr error

func tryODNodeRemovePolicy(node ODNodeRef, policyType ODPolicyType, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeRemovePolicy == nil {
		return false, symbolCallError("ODNodeRemovePolicy", "10.9", _oDNodeRemovePolicyErr)
	}
	return _oDNodeRemovePolicy(node, policyType, err), nil
}

// ODNodeRemovePolicy.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeRemovePolicy(_:_:_:)
func ODNodeRemovePolicy(node ODNodeRef, policyType ODPolicyType, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeRemovePolicy(node, policyType, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeSetAccountPolicies func(node ODNodeRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool
var _oDNodeSetAccountPoliciesErr error

func tryODNodeSetAccountPolicies(node ODNodeRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeSetAccountPolicies == nil {
		return false, symbolCallError("ODNodeSetAccountPolicies", "10.10", _oDNodeSetAccountPoliciesErr)
	}
	return _oDNodeSetAccountPolicies(node, policies, err), nil
}

// ODNodeSetAccountPolicies.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetAccountPolicies(_:_:_:)
func ODNodeSetAccountPolicies(node ODNodeRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeSetAccountPolicies(node, policies, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeSetCredentials func(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDNodeSetCredentialsErr error

func tryODNodeSetCredentials(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeSetCredentials == nil {
		return false, symbolCallError("ODNodeSetCredentials", "10.6", _oDNodeSetCredentialsErr)
	}
	return _oDNodeSetCredentials(node, recordType, recordName, password, err), nil
}

// ODNodeSetCredentials sets credentials for interacting with a node.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetCredentials(_:_:_:_:_:)
func ODNodeSetCredentials(node ODNodeRef, recordType unsafe.Pointer, recordName corefoundation.CFStringRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeSetCredentials(node, recordType, recordName, password, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeSetCredentialsExtended func(node ODNodeRef, recordType unsafe.Pointer, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) bool
var _oDNodeSetCredentialsExtendedErr error

func tryODNodeSetCredentialsExtended(node ODNodeRef, recordType unsafe.Pointer, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeSetCredentialsExtended == nil {
		return false, symbolCallError("ODNodeSetCredentialsExtended", "10.6", _oDNodeSetCredentialsExtendedErr)
	}
	return _oDNodeSetCredentialsExtended(node, recordType, authType, authItems, outAuthItems, outContext, err), nil
}

// ODNodeSetCredentialsExtended sets credentials for interacting with a node using a specified authentication method.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetCredentialsExtended(_:_:_:_:_:_:_:)
func ODNodeSetCredentialsExtended(node ODNodeRef, recordType unsafe.Pointer, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeSetCredentialsExtended(node, recordType, authType, authItems, outAuthItems, outContext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeSetCredentialsUsingKerberosCache func(node ODNodeRef, cacheName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDNodeSetCredentialsUsingKerberosCacheErr error

func tryODNodeSetCredentialsUsingKerberosCache(node ODNodeRef, cacheName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeSetCredentialsUsingKerberosCache == nil {
		return false, symbolCallError("ODNodeSetCredentialsUsingKerberosCache", "10.6", _oDNodeSetCredentialsUsingKerberosCacheErr)
	}
	return _oDNodeSetCredentialsUsingKerberosCache(node, cacheName, err), nil
}

// ODNodeSetCredentialsUsingKerberosCache sets credentials for interacting with a node with the Kerberos cache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetCredentialsUsingKerberosCache
func ODNodeSetCredentialsUsingKerberosCache(node ODNodeRef, cacheName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeSetCredentialsUsingKerberosCache(node, cacheName, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeSetPolicies func(node ODNodeRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool
var _oDNodeSetPoliciesErr error

func tryODNodeSetPolicies(node ODNodeRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeSetPolicies == nil {
		return false, symbolCallError("ODNodeSetPolicies", "10.9", _oDNodeSetPoliciesErr)
	}
	return _oDNodeSetPolicies(node, policies, err), nil
}

// ODNodeSetPolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetPolicies(_:_:_:)
func ODNodeSetPolicies(node ODNodeRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeSetPolicies(node, policies, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDNodeSetPolicy func(node ODNodeRef, policyType ODPolicyType, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool
var _oDNodeSetPolicyErr error

func tryODNodeSetPolicy(node ODNodeRef, policyType ODPolicyType, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDNodeSetPolicy == nil {
		return false, symbolCallError("ODNodeSetPolicy", "10.9", _oDNodeSetPolicyErr)
	}
	return _oDNodeSetPolicy(node, policyType, value, err), nil
}

// ODNodeSetPolicy.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeSetPolicy(_:_:_:_:)
func ODNodeSetPolicy(node ODNodeRef, policyType ODPolicyType, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODNodeSetPolicy(node, policyType, value, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDQueryCopyResults func(query ODQueryRef, allowPartialResults bool, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDQueryCopyResultsErr error

func tryODQueryCopyResults(query ODQueryRef, allowPartialResults bool, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDQueryCopyResults == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODQueryCopyResults", "10.6", _oDQueryCopyResultsErr)
	}
	return _oDQueryCopyResults(query, allowPartialResults, err), nil
}

// ODQueryCopyResults returns results from a query synchronously.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryCopyResults(_:_:_:)
func ODQueryCopyResults(query ODQueryRef, allowPartialResults bool, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODQueryCopyResults(query, allowPartialResults, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDQueryCreateWithNode func(allocator corefoundation.CFAllocatorRef, node ODNodeRef, recordTypeOrList corefoundation.CFTypeRef, attribute unsafe.Pointer, matchType ODMatchType, queryValueOrList corefoundation.CFTypeRef, returnAttributeOrList corefoundation.CFTypeRef, maxResults int, err *corefoundation.CFErrorRef) ODQueryRef
var _oDQueryCreateWithNodeErr error

func tryODQueryCreateWithNode(allocator corefoundation.CFAllocatorRef, node ODNodeRef, recordTypeOrList corefoundation.CFTypeRef, attribute unsafe.Pointer, matchType ODMatchType, queryValueOrList corefoundation.CFTypeRef, returnAttributeOrList corefoundation.CFTypeRef, maxResults int, err *corefoundation.CFErrorRef) (ODQueryRef, error) {
	if _oDQueryCreateWithNode == nil {
		return *new(ODQueryRef), symbolCallError("ODQueryCreateWithNode", "10.6", _oDQueryCreateWithNodeErr)
	}
	return _oDQueryCreateWithNode(allocator, node, recordTypeOrList, attribute, matchType, queryValueOrList, returnAttributeOrList, maxResults, err), nil
}

// ODQueryCreateWithNode creates a query with a node using provided parameters.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryCreateWithNode(_:_:_:_:_:_:_:_:_:)
func ODQueryCreateWithNode(allocator corefoundation.CFAllocatorRef, node ODNodeRef, recordTypeOrList corefoundation.CFTypeRef, attribute unsafe.Pointer, matchType ODMatchType, queryValueOrList corefoundation.CFTypeRef, returnAttributeOrList corefoundation.CFTypeRef, maxResults int, err *corefoundation.CFErrorRef) ODQueryRef {
	result, callErr := tryODQueryCreateWithNode(allocator, node, recordTypeOrList, attribute, matchType, queryValueOrList, returnAttributeOrList, maxResults, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDQueryCreateWithNodeType func(allocator corefoundation.CFAllocatorRef, nodeType ODNodeType, recordTypeOrList corefoundation.CFTypeRef, attribute unsafe.Pointer, matchType ODMatchType, queryValueOrList corefoundation.CFTypeRef, returnAttributeOrList corefoundation.CFTypeRef, maxResults int, err *corefoundation.CFErrorRef) ODQueryRef
var _oDQueryCreateWithNodeTypeErr error

func tryODQueryCreateWithNodeType(allocator corefoundation.CFAllocatorRef, nodeType ODNodeType, recordTypeOrList corefoundation.CFTypeRef, attribute unsafe.Pointer, matchType ODMatchType, queryValueOrList corefoundation.CFTypeRef, returnAttributeOrList corefoundation.CFTypeRef, maxResults int, err *corefoundation.CFErrorRef) (ODQueryRef, error) {
	if _oDQueryCreateWithNodeType == nil {
		return *new(ODQueryRef), symbolCallError("ODQueryCreateWithNodeType", "10.6", _oDQueryCreateWithNodeTypeErr)
	}
	return _oDQueryCreateWithNodeType(allocator, nodeType, recordTypeOrList, attribute, matchType, queryValueOrList, returnAttributeOrList, maxResults, err), nil
}

// ODQueryCreateWithNodeType creates a query for a particular node type using provided parameters.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryCreateWithNodeType(_:_:_:_:_:_:_:_:_:)
func ODQueryCreateWithNodeType(allocator corefoundation.CFAllocatorRef, nodeType ODNodeType, recordTypeOrList corefoundation.CFTypeRef, attribute unsafe.Pointer, matchType ODMatchType, queryValueOrList corefoundation.CFTypeRef, returnAttributeOrList corefoundation.CFTypeRef, maxResults int, err *corefoundation.CFErrorRef) ODQueryRef {
	result, callErr := tryODQueryCreateWithNodeType(allocator, nodeType, recordTypeOrList, attribute, matchType, queryValueOrList, returnAttributeOrList, maxResults, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDQueryGetTypeID func() uint
var _oDQueryGetTypeIDErr error

func tryODQueryGetTypeID() (uint, error) {
	if _oDQueryGetTypeID == nil {
		return 0, symbolCallError("ODQueryGetTypeID", "10.6", _oDQueryGetTypeIDErr)
	}
	return _oDQueryGetTypeID(), nil
}

// ODQueryGetTypeID returns the type ID for an Open Directory query.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryGetTypeID()
func ODQueryGetTypeID() uint {
	result, callErr := tryODQueryGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDQueryScheduleWithRunLoop func(query ODQueryRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _oDQueryScheduleWithRunLoopErr error

func tryODQueryScheduleWithRunLoop(query ODQueryRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _oDQueryScheduleWithRunLoop == nil {
		return symbolCallError("ODQueryScheduleWithRunLoop", "10.6", _oDQueryScheduleWithRunLoopErr)
	}
	_oDQueryScheduleWithRunLoop(query, runLoop, runLoopMode)
	return nil
}

// ODQueryScheduleWithRunLoop retrieves results from a query asynchronously by scheduling the query in a run loop.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryScheduleWithRunLoop(_:_:_:)
func ODQueryScheduleWithRunLoop(query ODQueryRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryODQueryScheduleWithRunLoop(query, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _oDQuerySetCallback func(query ODQueryRef, callback ODQueryCallback, userInfo unsafe.Pointer)
var _oDQuerySetCallbackErr error

func tryODQuerySetCallback(query ODQueryRef, callback ODQueryCallback, userInfo unsafe.Pointer) error {
	if _oDQuerySetCallback == nil {
		return symbolCallError("ODQuerySetCallback", "10.6", _oDQuerySetCallbackErr)
	}
	_oDQuerySetCallback(query, callback, userInfo)
	return nil
}

// ODQuerySetCallback sets the callback for an asynchronous query.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuerySetCallback(_:_:_:)
func ODQuerySetCallback(query ODQueryRef, callback ODQueryCallback, userInfo unsafe.Pointer) {
	if callErr := tryODQuerySetCallback(query, callback, userInfo); callErr != nil {
		panic(callErr)
	}
}

var _oDQuerySetDispatchQueue func(query ODQueryRef, queue uintptr)
var _oDQuerySetDispatchQueueErr error

func tryODQuerySetDispatchQueue(query ODQueryRef, queue dispatch.Queue) error {
	if _oDQuerySetDispatchQueue == nil {
		return symbolCallError("ODQuerySetDispatchQueue", "10.6", _oDQuerySetDispatchQueueErr)
	}
	_oDQuerySetDispatchQueue(query, uintptr(queue.Handle()))
	return nil
}

// ODQuerySetDispatchQueue retrieves results from a query asynchronously by adding the query to a dispatch queue.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuerySetDispatchQueue(_:_:)
func ODQuerySetDispatchQueue(query ODQueryRef, queue dispatch.Queue) {
	if callErr := tryODQuerySetDispatchQueue(query, queue); callErr != nil {
		panic(callErr)
	}
}

var _oDQuerySynchronize func(query ODQueryRef)
var _oDQuerySynchronizeErr error

func tryODQuerySynchronize(query ODQueryRef) error {
	if _oDQuerySynchronize == nil {
		return symbolCallError("ODQuerySynchronize", "10.6", _oDQuerySynchronizeErr)
	}
	_oDQuerySynchronize(query)
	return nil
}

// ODQuerySynchronize restarts a query, disposing of any results it has obtained.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuerySynchronize(_:)
func ODQuerySynchronize(query ODQueryRef) {
	if callErr := tryODQuerySynchronize(query); callErr != nil {
		panic(callErr)
	}
}

var _oDQueryUnscheduleFromRunLoop func(query ODQueryRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef)
var _oDQueryUnscheduleFromRunLoopErr error

func tryODQueryUnscheduleFromRunLoop(query ODQueryRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) error {
	if _oDQueryUnscheduleFromRunLoop == nil {
		return symbolCallError("ODQueryUnscheduleFromRunLoop", "10.6", _oDQueryUnscheduleFromRunLoopErr)
	}
	_oDQueryUnscheduleFromRunLoop(query, runLoop, runLoopMode)
	return nil
}

// ODQueryUnscheduleFromRunLoop removes a query from a specified run loop.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryUnscheduleFromRunLoop(_:_:_:)
func ODQueryUnscheduleFromRunLoop(query ODQueryRef, runLoop corefoundation.CFRunLoopRef, runLoopMode corefoundation.CFStringRef) {
	if callErr := tryODQueryUnscheduleFromRunLoop(query, runLoop, runLoopMode); callErr != nil {
		panic(callErr)
	}
}

var _oDRecordAddAccountPolicy func(record ODRecordRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool
var _oDRecordAddAccountPolicyErr error

func tryODRecordAddAccountPolicy(record ODRecordRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordAddAccountPolicy == nil {
		return false, symbolCallError("ODRecordAddAccountPolicy", "10.10", _oDRecordAddAccountPolicyErr)
	}
	return _oDRecordAddAccountPolicy(record, policy, category, err), nil
}

// ODRecordAddAccountPolicy.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordAddAccountPolicy(_:_:_:_:)
func ODRecordAddAccountPolicy(record ODRecordRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordAddAccountPolicy(record, policy, category, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordAddMember func(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) bool
var _oDRecordAddMemberErr error

func tryODRecordAddMember(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordAddMember == nil {
		return false, symbolCallError("ODRecordAddMember", "10.6", _oDRecordAddMemberErr)
	}
	return _oDRecordAddMember(group, member, err), nil
}

// ODRecordAddMember adds a record as a member of a group record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordAddMember(_:_:_:)
func ODRecordAddMember(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordAddMember(group, member, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordAddValue func(record ODRecordRef, attribute unsafe.Pointer, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool
var _oDRecordAddValueErr error

func tryODRecordAddValue(record ODRecordRef, attribute unsafe.Pointer, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordAddValue == nil {
		return false, symbolCallError("ODRecordAddValue", "10.6", _oDRecordAddValueErr)
	}
	return _oDRecordAddValue(record, attribute, value, err), nil
}

// ODRecordAddValue adds a value to an attribute of a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordAddValue(_:_:_:_:)
func ODRecordAddValue(record ODRecordRef, attribute unsafe.Pointer, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordAddValue(record, attribute, value, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordAuthenticationAllowed func(record ODRecordRef, err *corefoundation.CFErrorRef) bool
var _oDRecordAuthenticationAllowedErr error

func tryODRecordAuthenticationAllowed(record ODRecordRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordAuthenticationAllowed == nil {
		return false, symbolCallError("ODRecordAuthenticationAllowed", "10.10", _oDRecordAuthenticationAllowedErr)
	}
	return _oDRecordAuthenticationAllowed(record, err), nil
}

// ODRecordAuthenticationAllowed.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordAuthenticationAllowed(_:_:)
func ODRecordAuthenticationAllowed(record ODRecordRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordAuthenticationAllowed(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordChangePassword func(record ODRecordRef, oldPassword corefoundation.CFStringRef, newPassword corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDRecordChangePasswordErr error

func tryODRecordChangePassword(record ODRecordRef, oldPassword corefoundation.CFStringRef, newPassword corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordChangePassword == nil {
		return false, symbolCallError("ODRecordChangePassword", "10.6", _oDRecordChangePasswordErr)
	}
	return _oDRecordChangePassword(record, oldPassword, newPassword, err), nil
}

// ODRecordChangePassword changes the password of a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordChangePassword(_:_:_:_:)
func ODRecordChangePassword(record ODRecordRef, oldPassword corefoundation.CFStringRef, newPassword corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordChangePassword(record, oldPassword, newPassword, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordContainsMember func(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) bool
var _oDRecordContainsMemberErr error

func tryODRecordContainsMember(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordContainsMember == nil {
		return false, symbolCallError("ODRecordContainsMember", "10.6", _oDRecordContainsMemberErr)
	}
	return _oDRecordContainsMember(group, member, err), nil
}

// ODRecordContainsMember returns whether a group record contains a given record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordContainsMember(_:_:_:)
func ODRecordContainsMember(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordContainsMember(group, member, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopyAccountPolicies func(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDRecordCopyAccountPoliciesErr error

func tryODRecordCopyAccountPolicies(record ODRecordRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDRecordCopyAccountPolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODRecordCopyAccountPolicies", "10.10", _oDRecordCopyAccountPoliciesErr)
	}
	return _oDRecordCopyAccountPolicies(record, err), nil
}

// ODRecordCopyAccountPolicies.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyAccountPolicies(_:_:)
func ODRecordCopyAccountPolicies(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODRecordCopyAccountPolicies(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopyDetails func(record ODRecordRef, attributes corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDRecordCopyDetailsErr error

func tryODRecordCopyDetails(record ODRecordRef, attributes corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDRecordCopyDetails == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODRecordCopyDetails", "10.6", _oDRecordCopyDetailsErr)
	}
	return _oDRecordCopyDetails(record, attributes, err), nil
}

// ODRecordCopyDetails returns the values of a record’s attributes.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyDetails(_:_:_:)
func ODRecordCopyDetails(record ODRecordRef, attributes corefoundation.CFArrayRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODRecordCopyDetails(record, attributes, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopyEffectivePolicies func(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDRecordCopyEffectivePoliciesErr error

func tryODRecordCopyEffectivePolicies(record ODRecordRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDRecordCopyEffectivePolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODRecordCopyEffectivePolicies", "10.9", _oDRecordCopyEffectivePoliciesErr)
	}
	return _oDRecordCopyEffectivePolicies(record, err), nil
}

// ODRecordCopyEffectivePolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyEffectivePolicies(_:_:)
func ODRecordCopyEffectivePolicies(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODRecordCopyEffectivePolicies(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopyPasswordPolicy func(allocator corefoundation.CFAllocatorRef, record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDRecordCopyPasswordPolicyErr error

func tryODRecordCopyPasswordPolicy(allocator corefoundation.CFAllocatorRef, record ODRecordRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDRecordCopyPasswordPolicy == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODRecordCopyPasswordPolicy", "10.6", _oDRecordCopyPasswordPolicyErr)
	}
	return _oDRecordCopyPasswordPolicy(allocator, record, err), nil
}

// ODRecordCopyPasswordPolicy returns the password policies of a record.
//
// Deprecated: Deprecated since macOS 10.9.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyPasswordPolicy
func ODRecordCopyPasswordPolicy(allocator corefoundation.CFAllocatorRef, record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODRecordCopyPasswordPolicy(allocator, record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopyPolicies func(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDRecordCopyPoliciesErr error

func tryODRecordCopyPolicies(record ODRecordRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDRecordCopyPolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODRecordCopyPolicies", "10.9", _oDRecordCopyPoliciesErr)
	}
	return _oDRecordCopyPolicies(record, err), nil
}

// ODRecordCopyPolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyPolicies(_:_:)
func ODRecordCopyPolicies(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODRecordCopyPolicies(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopySupportedPolicies func(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef
var _oDRecordCopySupportedPoliciesErr error

func tryODRecordCopySupportedPolicies(record ODRecordRef, err *corefoundation.CFErrorRef) (corefoundation.CFDictionaryRef, error) {
	if _oDRecordCopySupportedPolicies == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("ODRecordCopySupportedPolicies", "10.9", _oDRecordCopySupportedPoliciesErr)
	}
	return _oDRecordCopySupportedPolicies(record, err), nil
}

// ODRecordCopySupportedPolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopySupportedPolicies(_:_:)
func ODRecordCopySupportedPolicies(record ODRecordRef, err *corefoundation.CFErrorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryODRecordCopySupportedPolicies(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordCopyValues func(record ODRecordRef, attribute unsafe.Pointer, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDRecordCopyValuesErr error

func tryODRecordCopyValues(record ODRecordRef, attribute unsafe.Pointer, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDRecordCopyValues == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODRecordCopyValues", "10.6", _oDRecordCopyValuesErr)
	}
	return _oDRecordCopyValues(record, attribute, err), nil
}

// ODRecordCopyValues returns the value of a single attribute of a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordCopyValues(_:_:_:)
func ODRecordCopyValues(record ODRecordRef, attribute unsafe.Pointer, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODRecordCopyValues(record, attribute, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordDelete func(record ODRecordRef, err *corefoundation.CFErrorRef) bool
var _oDRecordDeleteErr error

func tryODRecordDelete(record ODRecordRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordDelete == nil {
		return false, symbolCallError("ODRecordDelete", "10.6", _oDRecordDeleteErr)
	}
	return _oDRecordDelete(record, err), nil
}

// ODRecordDelete deletes a record from a node and invalidates the record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordDelete(_:_:)
func ODRecordDelete(record ODRecordRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordDelete(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordGetRecordName func(record ODRecordRef) corefoundation.CFStringRef
var _oDRecordGetRecordNameErr error

func tryODRecordGetRecordName(record ODRecordRef) (corefoundation.CFStringRef, error) {
	if _oDRecordGetRecordName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("ODRecordGetRecordName", "10.6", _oDRecordGetRecordNameErr)
	}
	return _oDRecordGetRecordName(record), nil
}

// ODRecordGetRecordName returns the official name of a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordGetRecordName(_:)
func ODRecordGetRecordName(record ODRecordRef) corefoundation.CFStringRef {
	result, callErr := tryODRecordGetRecordName(record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordGetRecordType func(record ODRecordRef) corefoundation.CFStringRef
var _oDRecordGetRecordTypeErr error

func tryODRecordGetRecordType(record ODRecordRef) (corefoundation.CFStringRef, error) {
	if _oDRecordGetRecordType == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("ODRecordGetRecordType", "10.6", _oDRecordGetRecordTypeErr)
	}
	return _oDRecordGetRecordType(record), nil
}

// ODRecordGetRecordType returns the type of a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordGetRecordType(_:)
func ODRecordGetRecordType(record ODRecordRef) corefoundation.CFStringRef {
	result, callErr := tryODRecordGetRecordType(record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordGetTypeID func() uint
var _oDRecordGetTypeIDErr error

func tryODRecordGetTypeID() (uint, error) {
	if _oDRecordGetTypeID == nil {
		return 0, symbolCallError("ODRecordGetTypeID", "10.6", _oDRecordGetTypeIDErr)
	}
	return _oDRecordGetTypeID(), nil
}

// ODRecordGetTypeID returns the type ID for a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordGetTypeID()
func ODRecordGetTypeID() uint {
	result, callErr := tryODRecordGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordPasswordChangeAllowed func(record ODRecordRef, newPassword corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDRecordPasswordChangeAllowedErr error

func tryODRecordPasswordChangeAllowed(record ODRecordRef, newPassword corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordPasswordChangeAllowed == nil {
		return false, symbolCallError("ODRecordPasswordChangeAllowed", "10.10", _oDRecordPasswordChangeAllowedErr)
	}
	return _oDRecordPasswordChangeAllowed(record, newPassword, err), nil
}

// ODRecordPasswordChangeAllowed.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordPasswordChangeAllowed(_:_:_:)
func ODRecordPasswordChangeAllowed(record ODRecordRef, newPassword corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordPasswordChangeAllowed(record, newPassword, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordRemoveAccountPolicy func(record ODRecordRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool
var _oDRecordRemoveAccountPolicyErr error

func tryODRecordRemoveAccountPolicy(record ODRecordRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordRemoveAccountPolicy == nil {
		return false, symbolCallError("ODRecordRemoveAccountPolicy", "10.10", _oDRecordRemoveAccountPolicyErr)
	}
	return _oDRecordRemoveAccountPolicy(record, policy, category, err), nil
}

// ODRecordRemoveAccountPolicy.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemoveAccountPolicy(_:_:_:_:)
func ODRecordRemoveAccountPolicy(record ODRecordRef, policy corefoundation.CFDictionaryRef, category ODPolicyCategoryType, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordRemoveAccountPolicy(record, policy, category, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordRemoveMember func(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) bool
var _oDRecordRemoveMemberErr error

func tryODRecordRemoveMember(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordRemoveMember == nil {
		return false, symbolCallError("ODRecordRemoveMember", "10.6", _oDRecordRemoveMemberErr)
	}
	return _oDRecordRemoveMember(group, member, err), nil
}

// ODRecordRemoveMember removes a record as a member from a specified group record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemoveMember(_:_:_:)
func ODRecordRemoveMember(group ODRecordRef, member ODRecordRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordRemoveMember(group, member, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordRemovePolicy func(record ODRecordRef, policy ODPolicyType, err *corefoundation.CFErrorRef) bool
var _oDRecordRemovePolicyErr error

func tryODRecordRemovePolicy(record ODRecordRef, policy ODPolicyType, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordRemovePolicy == nil {
		return false, symbolCallError("ODRecordRemovePolicy", "10.9", _oDRecordRemovePolicyErr)
	}
	return _oDRecordRemovePolicy(record, policy, err), nil
}

// ODRecordRemovePolicy.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemovePolicy(_:_:_:)
func ODRecordRemovePolicy(record ODRecordRef, policy ODPolicyType, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordRemovePolicy(record, policy, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordRemoveValue func(record ODRecordRef, attribute unsafe.Pointer, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool
var _oDRecordRemoveValueErr error

func tryODRecordRemoveValue(record ODRecordRef, attribute unsafe.Pointer, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordRemoveValue == nil {
		return false, symbolCallError("ODRecordRemoveValue", "10.6", _oDRecordRemoveValueErr)
	}
	return _oDRecordRemoveValue(record, attribute, value, err), nil
}

// ODRecordRemoveValue removes a value from a record’s attribute.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordRemoveValue(_:_:_:_:)
func ODRecordRemoveValue(record ODRecordRef, attribute unsafe.Pointer, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordRemoveValue(record, attribute, value, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSecondsUntilAuthenticationsExpire func(record ODRecordRef) int64
var _oDRecordSecondsUntilAuthenticationsExpireErr error

func tryODRecordSecondsUntilAuthenticationsExpire(record ODRecordRef) (int64, error) {
	if _oDRecordSecondsUntilAuthenticationsExpire == nil {
		return 0, symbolCallError("ODRecordSecondsUntilAuthenticationsExpire", "10.10", _oDRecordSecondsUntilAuthenticationsExpireErr)
	}
	return _oDRecordSecondsUntilAuthenticationsExpire(record), nil
}

// ODRecordSecondsUntilAuthenticationsExpire.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSecondsUntilAuthenticationsExpire(_:)
func ODRecordSecondsUntilAuthenticationsExpire(record ODRecordRef) int64 {
	result, callErr := tryODRecordSecondsUntilAuthenticationsExpire(record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSecondsUntilPasswordExpires func(record ODRecordRef) int64
var _oDRecordSecondsUntilPasswordExpiresErr error

func tryODRecordSecondsUntilPasswordExpires(record ODRecordRef) (int64, error) {
	if _oDRecordSecondsUntilPasswordExpires == nil {
		return 0, symbolCallError("ODRecordSecondsUntilPasswordExpires", "10.10", _oDRecordSecondsUntilPasswordExpiresErr)
	}
	return _oDRecordSecondsUntilPasswordExpires(record), nil
}

// ODRecordSecondsUntilPasswordExpires.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSecondsUntilPasswordExpires(_:)
func ODRecordSecondsUntilPasswordExpires(record ODRecordRef) int64 {
	result, callErr := tryODRecordSecondsUntilPasswordExpires(record)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetAccountPolicies func(record ODRecordRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetAccountPoliciesErr error

func tryODRecordSetAccountPolicies(record ODRecordRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetAccountPolicies == nil {
		return false, symbolCallError("ODRecordSetAccountPolicies", "10.10", _oDRecordSetAccountPoliciesErr)
	}
	return _oDRecordSetAccountPolicies(record, policies, err), nil
}

// ODRecordSetAccountPolicies.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetAccountPolicies(_:_:_:)
func ODRecordSetAccountPolicies(record ODRecordRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetAccountPolicies(record, policies, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetNodeCredentials func(record ODRecordRef, username corefoundation.CFStringRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetNodeCredentialsErr error

func tryODRecordSetNodeCredentials(record ODRecordRef, username corefoundation.CFStringRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetNodeCredentials == nil {
		return false, symbolCallError("ODRecordSetNodeCredentials", "10.6", _oDRecordSetNodeCredentialsErr)
	}
	return _oDRecordSetNodeCredentials(record, username, password, err), nil
}

// ODRecordSetNodeCredentials sets node authentication credentials for a given record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetNodeCredentials(_:_:_:_:)
func ODRecordSetNodeCredentials(record ODRecordRef, username corefoundation.CFStringRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetNodeCredentials(record, username, password, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetNodeCredentialsExtended func(record ODRecordRef, recordType unsafe.Pointer, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetNodeCredentialsExtendedErr error

func tryODRecordSetNodeCredentialsExtended(record ODRecordRef, recordType unsafe.Pointer, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetNodeCredentialsExtended == nil {
		return false, symbolCallError("ODRecordSetNodeCredentialsExtended", "10.6", _oDRecordSetNodeCredentialsExtendedErr)
	}
	return _oDRecordSetNodeCredentialsExtended(record, recordType, authType, authItems, outAuthItems, outContext, err), nil
}

// ODRecordSetNodeCredentialsExtended sets node authentication credentials for a record using a specified authentication method.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetNodeCredentialsExtended(_:_:_:_:_:_:_:)
func ODRecordSetNodeCredentialsExtended(record ODRecordRef, recordType unsafe.Pointer, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetNodeCredentialsExtended(record, recordType, authType, authItems, outAuthItems, outContext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetNodeCredentialsUsingKerberosCache func(record ODRecordRef, cacheName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetNodeCredentialsUsingKerberosCacheErr error

func tryODRecordSetNodeCredentialsUsingKerberosCache(record ODRecordRef, cacheName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetNodeCredentialsUsingKerberosCache == nil {
		return false, symbolCallError("ODRecordSetNodeCredentialsUsingKerberosCache", "10.6", _oDRecordSetNodeCredentialsUsingKerberosCacheErr)
	}
	return _oDRecordSetNodeCredentialsUsingKerberosCache(record, cacheName, err), nil
}

// ODRecordSetNodeCredentialsUsingKerberosCache sets credentials for interacting with a record’s node with the Kerberos cache.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetNodeCredentialsUsingKerberosCache
func ODRecordSetNodeCredentialsUsingKerberosCache(record ODRecordRef, cacheName corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetNodeCredentialsUsingKerberosCache(record, cacheName, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetPolicies func(record ODRecordRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetPoliciesErr error

func tryODRecordSetPolicies(record ODRecordRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetPolicies == nil {
		return false, symbolCallError("ODRecordSetPolicies", "10.9", _oDRecordSetPoliciesErr)
	}
	return _oDRecordSetPolicies(record, policies, err), nil
}

// ODRecordSetPolicies.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetPolicies(_:_:_:)
func ODRecordSetPolicies(record ODRecordRef, policies corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetPolicies(record, policies, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetPolicy func(record ODRecordRef, policy ODPolicyType, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetPolicyErr error

func tryODRecordSetPolicy(record ODRecordRef, policy ODPolicyType, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetPolicy == nil {
		return false, symbolCallError("ODRecordSetPolicy", "10.9", _oDRecordSetPolicyErr)
	}
	return _oDRecordSetPolicy(record, policy, value, err), nil
}

// ODRecordSetPolicy.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetPolicy(_:_:_:_:)
func ODRecordSetPolicy(record ODRecordRef, policy ODPolicyType, value corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetPolicy(record, policy, value, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSetValue func(record ODRecordRef, attribute unsafe.Pointer, valueOrValues corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSetValueErr error

func tryODRecordSetValue(record ODRecordRef, attribute unsafe.Pointer, valueOrValues corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSetValue == nil {
		return false, symbolCallError("ODRecordSetValue", "10.6", _oDRecordSetValueErr)
	}
	return _oDRecordSetValue(record, attribute, valueOrValues, err), nil
}

// ODRecordSetValue sets one or more attribute values of a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSetValue(_:_:_:_:)
func ODRecordSetValue(record ODRecordRef, attribute unsafe.Pointer, valueOrValues corefoundation.CFTypeRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSetValue(record, attribute, valueOrValues, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordSynchronize func(record ODRecordRef, err *corefoundation.CFErrorRef) bool
var _oDRecordSynchronizeErr error

func tryODRecordSynchronize(record ODRecordRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordSynchronize == nil {
		return false, symbolCallError("ODRecordSynchronize", "10.6", _oDRecordSynchronizeErr)
	}
	return _oDRecordSynchronize(record, err), nil
}

// ODRecordSynchronize synchronizes a record with the directory to get current data and commit changes.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordSynchronize(_:_:)
func ODRecordSynchronize(record ODRecordRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordSynchronize(record, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordVerifyPassword func(record ODRecordRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _oDRecordVerifyPasswordErr error

func tryODRecordVerifyPassword(record ODRecordRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordVerifyPassword == nil {
		return false, symbolCallError("ODRecordVerifyPassword", "10.6", _oDRecordVerifyPasswordErr)
	}
	return _oDRecordVerifyPassword(record, password, err), nil
}

// ODRecordVerifyPassword verifies a given password for a record.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordVerifyPassword(_:_:_:)
func ODRecordVerifyPassword(record ODRecordRef, password corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordVerifyPassword(record, password, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordVerifyPasswordExtended func(record ODRecordRef, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) bool
var _oDRecordVerifyPasswordExtendedErr error

func tryODRecordVerifyPasswordExtended(record ODRecordRef, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _oDRecordVerifyPasswordExtended == nil {
		return false, symbolCallError("ODRecordVerifyPasswordExtended", "10.6", _oDRecordVerifyPasswordExtendedErr)
	}
	return _oDRecordVerifyPasswordExtended(record, authType, authItems, outAuthItems, outContext, err), nil
}

// ODRecordVerifyPasswordExtended verifies a given password for a record given a specified authentication method.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordVerifyPasswordExtended(_:_:_:_:_:_:)
func ODRecordVerifyPasswordExtended(record ODRecordRef, authType ODAuthenticationType, authItems corefoundation.CFArrayRef, outAuthItems *corefoundation.CFArrayRef, outContext *ODContextRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryODRecordVerifyPasswordExtended(record, authType, authItems, outAuthItems, outContext, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordWillAuthenticationsExpire func(record ODRecordRef, willExpireIn uint64) bool
var _oDRecordWillAuthenticationsExpireErr error

func tryODRecordWillAuthenticationsExpire(record ODRecordRef, willExpireIn uint64) (bool, error) {
	if _oDRecordWillAuthenticationsExpire == nil {
		return false, symbolCallError("ODRecordWillAuthenticationsExpire", "10.10", _oDRecordWillAuthenticationsExpireErr)
	}
	return _oDRecordWillAuthenticationsExpire(record, willExpireIn), nil
}

// ODRecordWillAuthenticationsExpire.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordWillAuthenticationsExpire(_:_:)
func ODRecordWillAuthenticationsExpire(record ODRecordRef, willExpireIn uint64) bool {
	result, callErr := tryODRecordWillAuthenticationsExpire(record, willExpireIn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDRecordWillPasswordExpire func(record ODRecordRef, willExpireIn uint64) bool
var _oDRecordWillPasswordExpireErr error

func tryODRecordWillPasswordExpire(record ODRecordRef, willExpireIn uint64) (bool, error) {
	if _oDRecordWillPasswordExpire == nil {
		return false, symbolCallError("ODRecordWillPasswordExpire", "10.10", _oDRecordWillPasswordExpireErr)
	}
	return _oDRecordWillPasswordExpire(record, willExpireIn), nil
}

// ODRecordWillPasswordExpire.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordWillPasswordExpire(_:_:)
func ODRecordWillPasswordExpire(record ODRecordRef, willExpireIn uint64) bool {
	result, callErr := tryODRecordWillPasswordExpire(record, willExpireIn)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDSessionCopyNodeNames func(allocator corefoundation.CFAllocatorRef, session ODSessionRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef
var _oDSessionCopyNodeNamesErr error

func tryODSessionCopyNodeNames(allocator corefoundation.CFAllocatorRef, session ODSessionRef, err *corefoundation.CFErrorRef) (corefoundation.CFArrayRef, error) {
	if _oDSessionCopyNodeNames == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("ODSessionCopyNodeNames", "10.6", _oDSessionCopyNodeNamesErr)
	}
	return _oDSessionCopyNodeNames(allocator, session, err), nil
}

// ODSessionCopyNodeNames returns the names of nodes registered in a given session.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionCopyNodeNames(_:_:_:)
func ODSessionCopyNodeNames(allocator corefoundation.CFAllocatorRef, session ODSessionRef, err *corefoundation.CFErrorRef) corefoundation.CFArrayRef {
	result, callErr := tryODSessionCopyNodeNames(allocator, session, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDSessionCreate func(allocator corefoundation.CFAllocatorRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) ODSessionRef
var _oDSessionCreateErr error

func tryODSessionCreate(allocator corefoundation.CFAllocatorRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (ODSessionRef, error) {
	if _oDSessionCreate == nil {
		return *new(ODSessionRef), symbolCallError("ODSessionCreate", "10.6", _oDSessionCreateErr)
	}
	return _oDSessionCreate(allocator, options, err), nil
}

// ODSessionCreate creates a session to be passed to node functions.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionCreate(_:_:_:)
func ODSessionCreate(allocator corefoundation.CFAllocatorRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) ODSessionRef {
	result, callErr := tryODSessionCreate(allocator, options, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _oDSessionGetTypeID func() uint
var _oDSessionGetTypeIDErr error

func tryODSessionGetTypeID() (uint, error) {
	if _oDSessionGetTypeID == nil {
		return 0, symbolCallError("ODSessionGetTypeID", "10.6", _oDSessionGetTypeIDErr)
	}
	return _oDSessionGetTypeID(), nil
}

// ODSessionGetTypeID returns the type ID for a session.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionGetTypeID()
func ODSessionGetTypeID() uint {
	result, callErr := tryODSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_oDContextGetTypeID, &_oDContextGetTypeIDErr, frameworkHandle, "ODContextGetTypeID", "")
	registerFunc(&_oDNodeAddAccountPolicy, &_oDNodeAddAccountPolicyErr, frameworkHandle, "ODNodeAddAccountPolicy", "10.10")
	registerFunc(&_oDNodeCopyAccountPolicies, &_oDNodeCopyAccountPoliciesErr, frameworkHandle, "ODNodeCopyAccountPolicies", "10.10")
	registerFunc(&_oDNodeCopyDetails, &_oDNodeCopyDetailsErr, frameworkHandle, "ODNodeCopyDetails", "10.6")
	registerFunc(&_oDNodeCopyPolicies, &_oDNodeCopyPoliciesErr, frameworkHandle, "ODNodeCopyPolicies", "10.9")
	registerFunc(&_oDNodeCopyRecord, &_oDNodeCopyRecordErr, frameworkHandle, "ODNodeCopyRecord", "10.6")
	registerFunc(&_oDNodeCopySubnodeNames, &_oDNodeCopySubnodeNamesErr, frameworkHandle, "ODNodeCopySubnodeNames", "10.6")
	registerFunc(&_oDNodeCopySupportedAttributes, &_oDNodeCopySupportedAttributesErr, frameworkHandle, "ODNodeCopySupportedAttributes", "10.6")
	registerFunc(&_oDNodeCopySupportedPolicies, &_oDNodeCopySupportedPoliciesErr, frameworkHandle, "ODNodeCopySupportedPolicies", "10.9")
	registerFunc(&_oDNodeCopySupportedRecordTypes, &_oDNodeCopySupportedRecordTypesErr, frameworkHandle, "ODNodeCopySupportedRecordTypes", "10.6")
	registerFunc(&_oDNodeCopyUnreachableSubnodeNames, &_oDNodeCopyUnreachableSubnodeNamesErr, frameworkHandle, "ODNodeCopyUnreachableSubnodeNames", "10.6")
	registerFunc(&_oDNodeCreateCopy, &_oDNodeCreateCopyErr, frameworkHandle, "ODNodeCreateCopy", "10.6")
	registerFunc(&_oDNodeCreateRecord, &_oDNodeCreateRecordErr, frameworkHandle, "ODNodeCreateRecord", "10.6")
	registerFunc(&_oDNodeCreateWithName, &_oDNodeCreateWithNameErr, frameworkHandle, "ODNodeCreateWithName", "10.6")
	registerFunc(&_oDNodeCreateWithNodeType, &_oDNodeCreateWithNodeTypeErr, frameworkHandle, "ODNodeCreateWithNodeType", "10.6")
	registerFunc(&_oDNodeCustomCall, &_oDNodeCustomCallErr, frameworkHandle, "ODNodeCustomCall", "10.6")
	registerFunc(&_oDNodeCustomFunction, &_oDNodeCustomFunctionErr, frameworkHandle, "ODNodeCustomFunction", "10.9")
	registerFunc(&_oDNodeGetName, &_oDNodeGetNameErr, frameworkHandle, "ODNodeGetName", "10.6")
	registerFunc(&_oDNodeGetTypeID, &_oDNodeGetTypeIDErr, frameworkHandle, "ODNodeGetTypeID", "10.6")
	registerFunc(&_oDNodePasswordContentCheck, &_oDNodePasswordContentCheckErr, frameworkHandle, "ODNodePasswordContentCheck", "10.10")
	registerFunc(&_oDNodeRemoveAccountPolicy, &_oDNodeRemoveAccountPolicyErr, frameworkHandle, "ODNodeRemoveAccountPolicy", "10.10")
	registerFunc(&_oDNodeRemovePolicy, &_oDNodeRemovePolicyErr, frameworkHandle, "ODNodeRemovePolicy", "10.9")
	registerFunc(&_oDNodeSetAccountPolicies, &_oDNodeSetAccountPoliciesErr, frameworkHandle, "ODNodeSetAccountPolicies", "10.10")
	registerFunc(&_oDNodeSetCredentials, &_oDNodeSetCredentialsErr, frameworkHandle, "ODNodeSetCredentials", "10.6")
	registerFunc(&_oDNodeSetCredentialsExtended, &_oDNodeSetCredentialsExtendedErr, frameworkHandle, "ODNodeSetCredentialsExtended", "10.6")
	registerFunc(&_oDNodeSetCredentialsUsingKerberosCache, &_oDNodeSetCredentialsUsingKerberosCacheErr, frameworkHandle, "ODNodeSetCredentialsUsingKerberosCache", "10.6")
	registerFunc(&_oDNodeSetPolicies, &_oDNodeSetPoliciesErr, frameworkHandle, "ODNodeSetPolicies", "10.9")
	registerFunc(&_oDNodeSetPolicy, &_oDNodeSetPolicyErr, frameworkHandle, "ODNodeSetPolicy", "10.9")
	registerFunc(&_oDQueryCopyResults, &_oDQueryCopyResultsErr, frameworkHandle, "ODQueryCopyResults", "10.6")
	registerFunc(&_oDQueryCreateWithNode, &_oDQueryCreateWithNodeErr, frameworkHandle, "ODQueryCreateWithNode", "10.6")
	registerFunc(&_oDQueryCreateWithNodeType, &_oDQueryCreateWithNodeTypeErr, frameworkHandle, "ODQueryCreateWithNodeType", "10.6")
	registerFunc(&_oDQueryGetTypeID, &_oDQueryGetTypeIDErr, frameworkHandle, "ODQueryGetTypeID", "10.6")
	registerFunc(&_oDQueryScheduleWithRunLoop, &_oDQueryScheduleWithRunLoopErr, frameworkHandle, "ODQueryScheduleWithRunLoop", "10.6")
	registerFunc(&_oDQuerySetCallback, &_oDQuerySetCallbackErr, frameworkHandle, "ODQuerySetCallback", "10.6")
	registerFunc(&_oDQuerySetDispatchQueue, &_oDQuerySetDispatchQueueErr, frameworkHandle, "ODQuerySetDispatchQueue", "10.6")
	registerFunc(&_oDQuerySynchronize, &_oDQuerySynchronizeErr, frameworkHandle, "ODQuerySynchronize", "10.6")
	registerFunc(&_oDQueryUnscheduleFromRunLoop, &_oDQueryUnscheduleFromRunLoopErr, frameworkHandle, "ODQueryUnscheduleFromRunLoop", "10.6")
	registerFunc(&_oDRecordAddAccountPolicy, &_oDRecordAddAccountPolicyErr, frameworkHandle, "ODRecordAddAccountPolicy", "10.10")
	registerFunc(&_oDRecordAddMember, &_oDRecordAddMemberErr, frameworkHandle, "ODRecordAddMember", "10.6")
	registerFunc(&_oDRecordAddValue, &_oDRecordAddValueErr, frameworkHandle, "ODRecordAddValue", "10.6")
	registerFunc(&_oDRecordAuthenticationAllowed, &_oDRecordAuthenticationAllowedErr, frameworkHandle, "ODRecordAuthenticationAllowed", "10.10")
	registerFunc(&_oDRecordChangePassword, &_oDRecordChangePasswordErr, frameworkHandle, "ODRecordChangePassword", "10.6")
	registerFunc(&_oDRecordContainsMember, &_oDRecordContainsMemberErr, frameworkHandle, "ODRecordContainsMember", "10.6")
	registerFunc(&_oDRecordCopyAccountPolicies, &_oDRecordCopyAccountPoliciesErr, frameworkHandle, "ODRecordCopyAccountPolicies", "10.10")
	registerFunc(&_oDRecordCopyDetails, &_oDRecordCopyDetailsErr, frameworkHandle, "ODRecordCopyDetails", "10.6")
	registerFunc(&_oDRecordCopyEffectivePolicies, &_oDRecordCopyEffectivePoliciesErr, frameworkHandle, "ODRecordCopyEffectivePolicies", "10.9")
	registerFunc(&_oDRecordCopyPasswordPolicy, &_oDRecordCopyPasswordPolicyErr, frameworkHandle, "ODRecordCopyPasswordPolicy", "10.6")
	registerFunc(&_oDRecordCopyPolicies, &_oDRecordCopyPoliciesErr, frameworkHandle, "ODRecordCopyPolicies", "10.9")
	registerFunc(&_oDRecordCopySupportedPolicies, &_oDRecordCopySupportedPoliciesErr, frameworkHandle, "ODRecordCopySupportedPolicies", "10.9")
	registerFunc(&_oDRecordCopyValues, &_oDRecordCopyValuesErr, frameworkHandle, "ODRecordCopyValues", "10.6")
	registerFunc(&_oDRecordDelete, &_oDRecordDeleteErr, frameworkHandle, "ODRecordDelete", "10.6")
	registerFunc(&_oDRecordGetRecordName, &_oDRecordGetRecordNameErr, frameworkHandle, "ODRecordGetRecordName", "10.6")
	registerFunc(&_oDRecordGetRecordType, &_oDRecordGetRecordTypeErr, frameworkHandle, "ODRecordGetRecordType", "10.6")
	registerFunc(&_oDRecordGetTypeID, &_oDRecordGetTypeIDErr, frameworkHandle, "ODRecordGetTypeID", "10.6")
	registerFunc(&_oDRecordPasswordChangeAllowed, &_oDRecordPasswordChangeAllowedErr, frameworkHandle, "ODRecordPasswordChangeAllowed", "10.10")
	registerFunc(&_oDRecordRemoveAccountPolicy, &_oDRecordRemoveAccountPolicyErr, frameworkHandle, "ODRecordRemoveAccountPolicy", "10.10")
	registerFunc(&_oDRecordRemoveMember, &_oDRecordRemoveMemberErr, frameworkHandle, "ODRecordRemoveMember", "10.6")
	registerFunc(&_oDRecordRemovePolicy, &_oDRecordRemovePolicyErr, frameworkHandle, "ODRecordRemovePolicy", "10.9")
	registerFunc(&_oDRecordRemoveValue, &_oDRecordRemoveValueErr, frameworkHandle, "ODRecordRemoveValue", "10.6")
	registerFunc(&_oDRecordSecondsUntilAuthenticationsExpire, &_oDRecordSecondsUntilAuthenticationsExpireErr, frameworkHandle, "ODRecordSecondsUntilAuthenticationsExpire", "10.10")
	registerFunc(&_oDRecordSecondsUntilPasswordExpires, &_oDRecordSecondsUntilPasswordExpiresErr, frameworkHandle, "ODRecordSecondsUntilPasswordExpires", "10.10")
	registerFunc(&_oDRecordSetAccountPolicies, &_oDRecordSetAccountPoliciesErr, frameworkHandle, "ODRecordSetAccountPolicies", "10.10")
	registerFunc(&_oDRecordSetNodeCredentials, &_oDRecordSetNodeCredentialsErr, frameworkHandle, "ODRecordSetNodeCredentials", "10.6")
	registerFunc(&_oDRecordSetNodeCredentialsExtended, &_oDRecordSetNodeCredentialsExtendedErr, frameworkHandle, "ODRecordSetNodeCredentialsExtended", "10.6")
	registerFunc(&_oDRecordSetNodeCredentialsUsingKerberosCache, &_oDRecordSetNodeCredentialsUsingKerberosCacheErr, frameworkHandle, "ODRecordSetNodeCredentialsUsingKerberosCache", "10.6")
	registerFunc(&_oDRecordSetPolicies, &_oDRecordSetPoliciesErr, frameworkHandle, "ODRecordSetPolicies", "10.9")
	registerFunc(&_oDRecordSetPolicy, &_oDRecordSetPolicyErr, frameworkHandle, "ODRecordSetPolicy", "10.9")
	registerFunc(&_oDRecordSetValue, &_oDRecordSetValueErr, frameworkHandle, "ODRecordSetValue", "10.6")
	registerFunc(&_oDRecordSynchronize, &_oDRecordSynchronizeErr, frameworkHandle, "ODRecordSynchronize", "10.6")
	registerFunc(&_oDRecordVerifyPassword, &_oDRecordVerifyPasswordErr, frameworkHandle, "ODRecordVerifyPassword", "10.6")
	registerFunc(&_oDRecordVerifyPasswordExtended, &_oDRecordVerifyPasswordExtendedErr, frameworkHandle, "ODRecordVerifyPasswordExtended", "10.6")
	registerFunc(&_oDRecordWillAuthenticationsExpire, &_oDRecordWillAuthenticationsExpireErr, frameworkHandle, "ODRecordWillAuthenticationsExpire", "10.10")
	registerFunc(&_oDRecordWillPasswordExpire, &_oDRecordWillPasswordExpireErr, frameworkHandle, "ODRecordWillPasswordExpire", "10.10")
	registerFunc(&_oDSessionCopyNodeNames, &_oDSessionCopyNodeNamesErr, frameworkHandle, "ODSessionCopyNodeNames", "10.6")
	registerFunc(&_oDSessionCreate, &_oDSessionCreateErr, frameworkHandle, "ODSessionCreate", "10.6")
	registerFunc(&_oDSessionGetTypeID, &_oDSessionGetTypeIDErr, frameworkHandle, "ODSessionGetTypeID", "10.6")
}
