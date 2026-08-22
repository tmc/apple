// Code generated from Apple documentation. DO NOT EDIT.

package cryptotokenkit

import (
	"github.com/tmc/apple/objectivec"
)

// TKTLVTag is the type used to identify TLV format tags.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTLVTag
type TKTLVTag = uint64

// TKTokenDriverClassID is the type of the class identifier for the token driver.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/ClassID
type TKTokenDriverClassID = string

// TKTokenInstanceID is a type that represents the instance identifier of a token.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/InstanceID
type TKTokenInstanceID = string

// TKTokenObjectID is a unique and persistent identifier of a particular token object.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/ObjectID
type TKTokenObjectID = objectivec.Object

// TKTokenOperationConstraint is a token’s authentication constraint for a specific operation.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenOperationConstraint
type TKTokenOperationConstraint = objectivec.Object
