// Code generated from Apple documentation. DO NOT EDIT.

package objectivec

import (
	"unsafe"
	"github.com/tmc/apple/objc"
)

// BOOL is type to represent a Boolean value.
//
// See: https://developer.apple.com/documentation/ObjectiveC/BOOL
type BOOL = bool

// Category is an opaque type that represents a category.
//
// See: https://developer.apple.com/documentation/ObjectiveC/Category
type Category = uintptr

// Class is an opaque type that represents an Objective-C class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/Class
type Class = objc.Class

// IMP is a pointer to the start of a method implementation.
//
// See: https://developer.apple.com/documentation/ObjectiveC/IMP
type IMP = func()

// Ivar is an opaque type that represents an instance variable.
//
// See: https://developer.apple.com/documentation/ObjectiveC/Ivar
type Ivar = uintptr

// Method is an opaque type that represents a method in a class definition.
//
// See: https://developer.apple.com/documentation/ObjectiveC/Method
type Method = uintptr

// NSInteger is describes an integer.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSInteger
type NSInteger = int

// NSUInteger is describes an unsigned integer.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSUInteger
type NSUInteger = uint

// SEL is defines an opaque type that represents a method selector.
//
// See: https://developer.apple.com/documentation/ObjectiveC/SEL
type SEL = uintptr

// Id is a pointer to an instance of a class.
//
// See: https://developer.apple.com/documentation/ObjectiveC/id
type Id = Object

// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_handler
type Objc_exception_handler = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_matcher
type Objc_exception_matcher = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_exception_preprocessor
type Objc_exception_preprocessor = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_func_loadImage
type Objc_func_loadImage = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getClass
type Objc_hook_getClass = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getImageName
type Objc_hook_getImageName = string

// See: https://developer.apple.com/documentation/ObjectiveC/objc_hook_lazyClassNamer
type Objc_hook_lazyClassNamer = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_objectptr_t
type Objc_objectptr_t = unsafe.Pointer

// Objc_property_t is an opaque type that represents an Objective-C declared property.
//
// See: https://developer.apple.com/documentation/ObjectiveC/objc_property_t
type Objc_property_t = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_uncaught_exception_handler
type Objc_uncaught_exception_handler = unsafe.Pointer

// See: https://developer.apple.com/documentation/ObjectiveC/objc_zone_t
type Objc_zone_t = uintptr

