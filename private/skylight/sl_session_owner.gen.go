// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSessionOwner] class.
var (
	_SLSessionOwnerClass     SLSessionOwnerClass
	_SLSessionOwnerClassOnce sync.Once
)

func getSLSessionOwnerClass() SLSessionOwnerClass {
	_SLSessionOwnerClassOnce.Do(func() {
		_SLSessionOwnerClass = SLSessionOwnerClass{class: objc.GetClass("SLSessionOwner")}
	})
	return _SLSessionOwnerClass
}

// GetSLSessionOwnerClass returns the class object for SLSessionOwner.
func GetSLSessionOwnerClass() SLSessionOwnerClass {
	return getSLSessionOwnerClass()
}

type SLSessionOwnerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSessionOwnerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSessionOwnerClass) Alloc() SLSessionOwner {
	rv := objc.Send[SLSessionOwner](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSessionOwner.AuditSessionID]
//   - [SLSessionOwner.SetAuditSessionID]
//   - [SLSessionOwner.CreateXPCSerializationAndInvalidate]
//   - [SLSessionOwner.IsValid]
//   - [SLSessionOwner.Port]
//   - [SLSessionOwner.SetPort]
//   - [SLSessionOwner.SessionID]
//   - [SLSessionOwner.SetSessionID]
//   - [SLSessionOwner.InitWithPortAuditSessionIDCgSessionID]
//   - [SLSessionOwner.InitWithXPCSerialization]
//   - [SLSessionOwner.Valid]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner
type SLSessionOwner struct {
	objectivec.Object
}

// SLSessionOwnerFromID constructs a [SLSessionOwner] from an objc.ID.
func SLSessionOwnerFromID(id objc.ID) SLSessionOwner {
	return SLSessionOwner{objectivec.Object{ID: id}}
}

// Ensure SLSessionOwner implements ISLSessionOwner.
var _ ISLSessionOwner = SLSessionOwner{}

// An interface definition for the [SLSessionOwner] class.
//
// # Methods
//
//   - [ISLSessionOwner.AuditSessionID]
//   - [ISLSessionOwner.SetAuditSessionID]
//   - [ISLSessionOwner.CreateXPCSerializationAndInvalidate]
//   - [ISLSessionOwner.IsValid]
//   - [ISLSessionOwner.Port]
//   - [ISLSessionOwner.SetPort]
//   - [ISLSessionOwner.SessionID]
//   - [ISLSessionOwner.SetSessionID]
//   - [ISLSessionOwner.InitWithPortAuditSessionIDCgSessionID]
//   - [ISLSessionOwner.InitWithXPCSerialization]
//   - [ISLSessionOwner.Valid]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner
type ISLSessionOwner interface {
	objectivec.IObject

	// Topic: Methods

	AuditSessionID() foundation.NSNumber
	SetAuditSessionID(value foundation.NSNumber)
	CreateXPCSerializationAndInvalidate() objectivec.IObject
	IsValid() bool
	Port() uint32
	SetPort(value uint32)
	SessionID() foundation.NSNumber
	SetSessionID(value foundation.NSNumber)
	InitWithPortAuditSessionIDCgSessionID(port uint32, id int, id2 uint32) SLSessionOwner
	InitWithXPCSerialization(xPCSerialization objectivec.IObject) SLSessionOwner
	Valid() bool
}

// Init initializes the instance.
func (s SLSessionOwner) Init() SLSessionOwner {
	rv := objc.Send[SLSessionOwner](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSessionOwner) Autorelease() SLSessionOwner {
	rv := objc.Send[SLSessionOwner](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSessionOwner creates a new SLSessionOwner instance.
func NewSLSessionOwner() SLSessionOwner {
	class := getSLSessionOwnerClass()
	rv := objc.Send[SLSessionOwner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/initWithPort:auditSessionID:cgSessionID:
func NewSLSessionOwnerWithPortAuditSessionIDCgSessionID(port uint32, id int, id2 uint32) SLSessionOwner {
	instance := getSLSessionOwnerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPort:auditSessionID:cgSessionID:"), port, id, id2)
	return SLSessionOwnerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/initWithXPCSerialization:
func NewSLSessionOwnerWithXPCSerialization(xPCSerialization objectivec.IObject) SLSessionOwner {
	instance := getSLSessionOwnerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCSerialization:"), xPCSerialization)
	return SLSessionOwnerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/createXPCSerializationAndInvalidate
func (s SLSessionOwner) CreateXPCSerializationAndInvalidate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("createXPCSerializationAndInvalidate"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/isValid
func (s SLSessionOwner) IsValid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isValid"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/initWithPort:auditSessionID:cgSessionID:
func (s SLSessionOwner) InitWithPortAuditSessionIDCgSessionID(port uint32, id int, id2 uint32) SLSessionOwner {
	rv := objc.Send[SLSessionOwner](s.ID, objc.Sel("initWithPort:auditSessionID:cgSessionID:"), port, id, id2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/initWithXPCSerialization:
func (s SLSessionOwner) InitWithXPCSerialization(xPCSerialization objectivec.IObject) SLSessionOwner {
	rv := objc.Send[SLSessionOwner](s.ID, objc.Sel("initWithXPCSerialization:"), xPCSerialization)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/sessionOwnerBySettingLoginwindowConnection:
func (_SLSessionOwnerClass SLSessionOwnerClass) SessionOwnerBySettingLoginwindowConnection(connection uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSessionOwnerClass.class), objc.Sel("sessionOwnerBySettingLoginwindowConnection:"), connection)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/sessionOwnerForNewSessionWithAuditSessionID:
func (_SLSessionOwnerClass SLSessionOwnerClass) SessionOwnerForNewSessionWithAuditSessionID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSessionOwnerClass.class), objc.Sel("sessionOwnerForNewSessionWithAuditSessionID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/sessionOwnerForNewSessionWithAuditSessionID:launchData:
func (_SLSessionOwnerClass SLSessionOwnerClass) SessionOwnerForNewSessionWithAuditSessionIDLaunchData(id objectivec.IObject, data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSessionOwnerClass.class), objc.Sel("sessionOwnerForNewSessionWithAuditSessionID:launchData:"), id, data)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/sessionOwnerWithPort:auditSessionID:cgSessionID:
func (_SLSessionOwnerClass SLSessionOwnerClass) SessionOwnerWithPortAuditSessionIDCgSessionID(port uint32, id int, id2 uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSessionOwnerClass.class), objc.Sel("sessionOwnerWithPort:auditSessionID:cgSessionID:"), port, id, id2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/sessionOwnerWithXPCSerialization:
func (_SLSessionOwnerClass SLSessionOwnerClass) SessionOwnerWithXPCSerialization(xPCSerialization objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSessionOwnerClass.class), objc.Sel("sessionOwnerWithXPCSerialization:"), xPCSerialization)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/auditSessionID
func (s SLSessionOwner) AuditSessionID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("auditSessionID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLSessionOwner) SetAuditSessionID(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setAuditSessionID:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/port
func (s SLSessionOwner) Port() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("port"))
	return rv
}
func (s SLSessionOwner) SetPort(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setPort:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/sessionID
func (s SLSessionOwner) SessionID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sessionID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLSessionOwner) SetSessionID(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setSessionID:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSessionOwner/valid
func (s SLSessionOwner) Valid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("valid"))
	return rv
}
