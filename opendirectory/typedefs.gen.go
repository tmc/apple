// Code generated from Apple documentation. DO NOT EDIT.

package opendirectory

import (
	"unsafe"
)

// ODAuthenticationType is an Open Directory authentication type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODAuthenticationType
type ODAuthenticationType = string

// ODContextRef is an Open Directory context type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODContext
type ODContextRef uintptr

// See: https://developer.apple.com/documentation/OpenDirectory/ODErrorUserInfoKeyType
type ODErrorUserInfoKeyType = string

// ODMatchType is an Open Directory match type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODMatchType
type ODMatchType = uint32

// ODNodeRef is an Open Directory node type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeRef
type ODNodeRef uintptr

// ODNodeType is an Open Directory node type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODNodeType
type ODNodeType = uint32

// See: https://developer.apple.com/documentation/OpenDirectory/ODOptionKeyType
type ODOptionKeyType = string

// See: https://developer.apple.com/documentation/OpenDirectory/ODPolicyAttributeType
type ODPolicyAttributeType = string

// See: https://developer.apple.com/documentation/OpenDirectory/ODPolicyCategoryType
type ODPolicyCategoryType = string

// See: https://developer.apple.com/documentation/OpenDirectory/ODPolicyKeyType
type ODPolicyKeyType = string

// See: https://developer.apple.com/documentation/OpenDirectory/ODPolicyType
type ODPolicyType = string

// ODQueryCallback is a callback function called as results from a scheduled query are returned.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryCallback
type ODQueryCallback = func(ODQueryRef, uintptr, uintptr, unsafe.Pointer)

// ODQueryRef is an Open Directory query type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQueryRef
type ODQueryRef uintptr

// ODRecordRef is an Open Directory record type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODRecordRef
type ODRecordRef uintptr

// ODSessionRef is an Open Directory session type.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODSessionRef
type ODSessionRef uintptr
