// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSecureCursorAssertion] class.
var (
	_SLSecureCursorAssertionClass     SLSecureCursorAssertionClass
	_SLSecureCursorAssertionClassOnce sync.Once
)

func getSLSecureCursorAssertionClass() SLSecureCursorAssertionClass {
	_SLSecureCursorAssertionClassOnce.Do(func() {
		_SLSecureCursorAssertionClass = SLSecureCursorAssertionClass{class: objc.GetClass("SLSecureCursorAssertion")}
	})
	return _SLSecureCursorAssertionClass
}

// GetSLSecureCursorAssertionClass returns the class object for SLSecureCursorAssertion.
func GetSLSecureCursorAssertionClass() SLSecureCursorAssertionClass {
	return getSLSecureCursorAssertionClass()
}

type SLSecureCursorAssertionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSecureCursorAssertionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSecureCursorAssertionClass) Alloc() SLSecureCursorAssertion {
	rv := objc.Send[SLSecureCursorAssertion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSecureCursorAssertion.Invalidate]
//   - [SLSecureCursorAssertion.IsValid]
//   - [SLSecureCursorAssertion.Uuid]
//   - [SLSecureCursorAssertion.SetUuid]
//   - [SLSecureCursorAssertion.Valid]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion
type SLSecureCursorAssertion struct {
	objectivec.Object
}

// SLSecureCursorAssertionFromID constructs a [SLSecureCursorAssertion] from an objc.ID.
func SLSecureCursorAssertionFromID(id objc.ID) SLSecureCursorAssertion {
	return SLSecureCursorAssertion{objectivec.Object{ID: id}}
}

// Ensure SLSecureCursorAssertion implements ISLSecureCursorAssertion.
var _ ISLSecureCursorAssertion = SLSecureCursorAssertion{}

// An interface definition for the [SLSecureCursorAssertion] class.
//
// # Methods
//
//   - [ISLSecureCursorAssertion.Invalidate]
//   - [ISLSecureCursorAssertion.IsValid]
//   - [ISLSecureCursorAssertion.Uuid]
//   - [ISLSecureCursorAssertion.SetUuid]
//   - [ISLSecureCursorAssertion.Valid]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion
type ISLSecureCursorAssertion interface {
	objectivec.IObject

	// Topic: Methods

	Invalidate()
	IsValid() bool
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	Valid() bool
}

// Init initializes the instance.
func (s SLSecureCursorAssertion) Init() SLSecureCursorAssertion {
	rv := objc.Send[SLSecureCursorAssertion](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSecureCursorAssertion) Autorelease() SLSecureCursorAssertion {
	rv := objc.Send[SLSecureCursorAssertion](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSecureCursorAssertion creates a new SLSecureCursorAssertion instance.
func NewSLSecureCursorAssertion() SLSecureCursorAssertion {
	class := getSLSecureCursorAssertionClass()
	rv := objc.Send[SLSecureCursorAssertion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion/invalidate
func (s SLSecureCursorAssertion) Invalidate() {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidate"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion/isValid
func (s SLSecureCursorAssertion) IsValid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isValid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion/assertion
func (_SLSecureCursorAssertionClass SLSecureCursorAssertionClass) Assertion() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSecureCursorAssertionClass.class), objc.Sel("assertion"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion/invalidateAll
func (_SLSecureCursorAssertionClass SLSecureCursorAssertionClass) InvalidateAll() {
	objc.Send[objc.ID](objc.ID(_SLSecureCursorAssertionClass.class), objc.Sel("invalidateAll"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion/uuid
func (s SLSecureCursorAssertion) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (s SLSecureCursorAssertion) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](s.ID, objc.Sel("setUuid:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertion/valid
func (s SLSecureCursorAssertion) Valid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("valid"))
	return rv
}
