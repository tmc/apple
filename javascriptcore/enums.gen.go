// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/JavaScriptCore/JSRelationCondition
type JSRelationCondition uint32

const ()

type KJSClassAttribute uint

const (
	// KJSClassAttributeNoAutomaticPrototype: An attribute that specifies that a class doesn’t automatically generate a shared prototype for its instance objects.
	KJSClassAttributeNoAutomaticPrototype KJSClassAttribute = 2
	// KJSClassAttributeNone: An attribute that specifies that a class has no special attributes.
	KJSClassAttributeNone KJSClassAttribute = 0
)

func (e KJSClassAttribute) String() string {
	switch e {
	case KJSClassAttributeNoAutomaticPrototype:
		return "KJSClassAttributeNoAutomaticPrototype"
	case KJSClassAttributeNone:
		return "KJSClassAttributeNone"
	default:
		return fmt.Sprintf("KJSClassAttribute(%d)", e)
	}
}

type KJSPropertyAttribute uint

const (
	// KJSPropertyAttributeDontDelete: An attribute that specifies that the delete operation fails on a property.
	KJSPropertyAttributeDontDelete KJSPropertyAttribute = 8
	// KJSPropertyAttributeDontEnum: An attribute that specifies that property enumerators and JavaScript for-in loops don’t enumerate a property.
	KJSPropertyAttributeDontEnum KJSPropertyAttribute = 4
	// KJSPropertyAttributeNone: An attribute that specifies that a property has no special attributes.
	KJSPropertyAttributeNone KJSPropertyAttribute = 0
	// KJSPropertyAttributeReadOnly: An attribute that specifies that a property is read-only.
	KJSPropertyAttributeReadOnly KJSPropertyAttribute = 2
)

func (e KJSPropertyAttribute) String() string {
	switch e {
	case KJSPropertyAttributeDontDelete:
		return "KJSPropertyAttributeDontDelete"
	case KJSPropertyAttributeDontEnum:
		return "KJSPropertyAttributeDontEnum"
	case KJSPropertyAttributeNone:
		return "KJSPropertyAttributeNone"
	case KJSPropertyAttributeReadOnly:
		return "KJSPropertyAttributeReadOnly"
	default:
		return fmt.Sprintf("KJSPropertyAttribute(%d)", e)
	}
}

type KJSType uint

const (
	KJSTypeBigInt KJSType = 7
	// KJSTypeBoolean: A primitive Boolean value.
	KJSTypeBoolean KJSType = 2
	// KJSTypeNull: The unique null value.
	KJSTypeNull KJSType = 1
	// KJSTypeNumber: A primitive number value.
	KJSTypeNumber KJSType = 3
	// KJSTypeObject: An object value.
	KJSTypeObject KJSType = 5
	// KJSTypeString: A primitive string value.
	KJSTypeString KJSType = 4
	// KJSTypeSymbol: A primitive symbol value.
	KJSTypeSymbol KJSType = 6
	// KJSTypeUndefined: The unique undefined value.
	KJSTypeUndefined KJSType = 0
)

func (e KJSType) String() string {
	switch e {
	case KJSTypeBigInt:
		return "KJSTypeBigInt"
	case KJSTypeBoolean:
		return "KJSTypeBoolean"
	case KJSTypeNull:
		return "KJSTypeNull"
	case KJSTypeNumber:
		return "KJSTypeNumber"
	case KJSTypeObject:
		return "KJSTypeObject"
	case KJSTypeString:
		return "KJSTypeString"
	case KJSTypeSymbol:
		return "KJSTypeSymbol"
	case KJSTypeUndefined:
		return "KJSTypeUndefined"
	default:
		return fmt.Sprintf("KJSType(%d)", e)
	}
}

type KJSTypedArrayType uint

const (
	// KJSTypedArrayTypeArrayBuffer: An array buffer type.
	KJSTypedArrayTypeArrayBuffer    KJSTypedArrayType = 9
	KJSTypedArrayTypeBigInt64Array  KJSTypedArrayType = 11
	KJSTypedArrayTypeBigUint64Array KJSTypedArrayType = 12
	// KJSTypedArrayTypeFloat32Array: A 32-bit floating point array type.
	KJSTypedArrayTypeFloat32Array KJSTypedArrayType = 7
	// KJSTypedArrayTypeFloat64Array: A 64-bit floating point array type.
	KJSTypedArrayTypeFloat64Array KJSTypedArrayType = 8
	// KJSTypedArrayTypeInt16Array: A 16-bit integer array type.
	KJSTypedArrayTypeInt16Array KJSTypedArrayType = 1
	// KJSTypedArrayTypeInt32Array: A 32-bit integer array type.
	KJSTypedArrayTypeInt32Array KJSTypedArrayType = 2
	// KJSTypedArrayTypeInt8Array: An 8-bit integer array type.
	KJSTypedArrayTypeInt8Array KJSTypedArrayType = 0
	// KJSTypedArrayTypeNone: Not a typed array.
	KJSTypedArrayTypeNone KJSTypedArrayType = 10
	// KJSTypedArrayTypeUint16Array: A 16-bit unsigned integer array type.
	KJSTypedArrayTypeUint16Array KJSTypedArrayType = 5
	// KJSTypedArrayTypeUint32Array: A 32-bit unsigned integer array type.
	KJSTypedArrayTypeUint32Array KJSTypedArrayType = 6
	// KJSTypedArrayTypeUint8Array: An 8-bit unsigned integer array type.
	KJSTypedArrayTypeUint8Array KJSTypedArrayType = 3
	// KJSTypedArrayTypeUint8ClampedArray: An 8-bit unsigned integer clamped array type.
	KJSTypedArrayTypeUint8ClampedArray KJSTypedArrayType = 4
)

func (e KJSTypedArrayType) String() string {
	switch e {
	case KJSTypedArrayTypeArrayBuffer:
		return "KJSTypedArrayTypeArrayBuffer"
	case KJSTypedArrayTypeBigInt64Array:
		return "KJSTypedArrayTypeBigInt64Array"
	case KJSTypedArrayTypeBigUint64Array:
		return "KJSTypedArrayTypeBigUint64Array"
	case KJSTypedArrayTypeFloat32Array:
		return "KJSTypedArrayTypeFloat32Array"
	case KJSTypedArrayTypeFloat64Array:
		return "KJSTypedArrayTypeFloat64Array"
	case KJSTypedArrayTypeInt16Array:
		return "KJSTypedArrayTypeInt16Array"
	case KJSTypedArrayTypeInt32Array:
		return "KJSTypedArrayTypeInt32Array"
	case KJSTypedArrayTypeInt8Array:
		return "KJSTypedArrayTypeInt8Array"
	case KJSTypedArrayTypeNone:
		return "KJSTypedArrayTypeNone"
	case KJSTypedArrayTypeUint16Array:
		return "KJSTypedArrayTypeUint16Array"
	case KJSTypedArrayTypeUint32Array:
		return "KJSTypedArrayTypeUint32Array"
	case KJSTypedArrayTypeUint8Array:
		return "KJSTypedArrayTypeUint8Array"
	case KJSTypedArrayTypeUint8ClampedArray:
		return "KJSTypedArrayTypeUint8ClampedArray"
	default:
		return fmt.Sprintf("KJSTypedArrayType(%d)", e)
	}
}

// JSType is an alias for referenced enum type KJSType.
type JSType = KJSType

// JSTypedArrayType is an alias for referenced enum type KJSTypedArrayType.
type JSTypedArrayType = KJSTypedArrayType
