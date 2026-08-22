// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSharingSession] class.
var (
	_SLSharingSessionClass     SLSharingSessionClass
	_SLSharingSessionClassOnce sync.Once
)

func getSLSharingSessionClass() SLSharingSessionClass {
	_SLSharingSessionClassOnce.Do(func() {
		_SLSharingSessionClass = SLSharingSessionClass{class: objc.GetClass("SLSharingSession")}
	})
	return _SLSharingSessionClass
}

// GetSLSharingSessionClass returns the class object for SLSharingSession.
func GetSLSharingSessionClass() SLSharingSessionClass {
	return getSLSharingSessionClass()
}

type SLSharingSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSharingSessionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSharingSessionClass) Alloc() SLSharingSession {
	rv := objc.SendIfResponds[SLSharingSession](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSharingSession.Content]
//   - [SLSharingSession.GetUUID]
//   - [SLSharingSession.GetUUIDBytes]
//   - [SLSharingSession.IsEqualToSharingSession]
//   - [SLSharingSession.LifetimePort]
//   - [SLSharingSession.SetLifetimePort]
//   - [SLSharingSession.SetContent]
//   - [SLSharingSession.SetPresentationDisplayPrimaryEnableShowCursor]
//   - [SLSharingSession.ShowPicker]
//   - [SLSharingSession.Title]
//   - [SLSharingSession.Type]
//   - [SLSharingSession.SetType]
//   - [SLSharingSession.Uuid]
//   - [SLSharingSession.Uuid_internal]
//   - [SLSharingSession.SetUuid_internal]
//   - [SLSharingSession.InitFromUUID]
//   - [SLSharingSession.InitWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications]
//   - [SLSharingSession.InitWithUUIDTitleType]
type SLSharingSession struct {
	objectivec.Object
}

// SLSharingSessionFromID constructs a [SLSharingSession] from an objc.ID.
func SLSharingSessionFromID(id objc.ID) SLSharingSession {
	return SLSharingSession{objectivec.Object{ID: id}}
}

// Ensure SLSharingSession implements ISLSharingSession.
var _ ISLSharingSession = SLSharingSession{}

// An interface definition for the [SLSharingSession] class.
//
// # Methods
//
//   - [ISLSharingSession.Content]
//   - [ISLSharingSession.GetUUID]
//   - [ISLSharingSession.GetUUIDBytes]
//   - [ISLSharingSession.IsEqualToSharingSession]
//   - [ISLSharingSession.LifetimePort]
//   - [ISLSharingSession.SetLifetimePort]
//   - [ISLSharingSession.SetContent]
//   - [ISLSharingSession.SetPresentationDisplayPrimaryEnableShowCursor]
//   - [ISLSharingSession.ShowPicker]
//   - [ISLSharingSession.Title]
//   - [ISLSharingSession.Type]
//   - [ISLSharingSession.SetType]
//   - [ISLSharingSession.Uuid]
//   - [ISLSharingSession.Uuid_internal]
//   - [ISLSharingSession.SetUuid_internal]
//   - [ISLSharingSession.InitFromUUID]
//   - [ISLSharingSession.InitWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications]
//   - [ISLSharingSession.InitWithUUIDTitleType]
type ISLSharingSession interface {
	objectivec.IObject

	// Topic: Methods

	Content() objectivec.IObject
	GetUUID() objectivec.IObject
	GetUUIDBytes() unsafe.Pointer
	IsEqualToSharingSession(session objectivec.IObject) bool
	LifetimePort() uint32
	SetLifetimePort(value uint32)
	SetContent(content objectivec.IObject)
	SetPresentationDisplayPrimaryEnableShowCursor(display objectivec.IObject, primary objectivec.IObject, enable bool, cursor bool)
	ShowPicker()
	Title() string
	Type() int
	SetType(value int)
	Uuid() foundation.NSUUID
	Uuid_internal() foundation.NSUUID
	SetUuid_internal(value foundation.NSUUID)
	InitFromUUID(uuid objectivec.IObject) SLSharingSession
	InitWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) SLSharingSession
	InitWithUUIDTitleType(uuid objectivec.IObject, title objectivec.IObject, type_ int) SLSharingSession
}

// Init initializes the instance.
func (s SLSharingSession) Init() SLSharingSession {
	rv := objc.SendIfResponds[SLSharingSession](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSharingSession) Autorelease() SLSharingSession {
	rv := objc.SendIfResponds[SLSharingSession](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSharingSession creates a new SLSharingSession instance.
func NewSLSharingSession() SLSharingSession {
	class := getSLSharingSessionClass()
	rv := objc.SendIfResponds[SLSharingSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSharingSessionFromUUID(uuid objectivec.IObject) SLSharingSession {
	instance := getSLSharingSessionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromUUID:"), uuid)
	return SLSharingSessionFromID(rv)
}

func NewSLSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) SLSharingSession {
	instance := getSLSharingSessionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:"), title, indicator, notifications)
	return SLSharingSessionFromID(rv)
}

func NewSLSharingSessionWithUUIDTitleType(uuid objectivec.IObject, title objectivec.IObject, type_ int) SLSharingSession {
	instance := getSLSharingSessionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithUUID:title:type:"), uuid, title, type_)
	return SLSharingSessionFromID(rv)
}

func (s SLSharingSession) Content() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("content"))
	return objectivec.Object{ID: rv}
}
func (s SLSharingSession) GetUUID() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("getUUID"))
	return objectivec.Object{ID: rv}
}
func (s SLSharingSession) GetUUIDBytes() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("getUUIDBytes"))
	return rv
}
func (s SLSharingSession) IsEqualToSharingSession(session objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isEqualToSharingSession:"), session)
	return rv
}
func (s SLSharingSession) SetContent(content objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setContent:"), content)
}
func (s SLSharingSession) SetPresentationDisplayPrimaryEnableShowCursor(display objectivec.IObject, primary objectivec.IObject, enable bool, cursor bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setPresentationDisplay:primary:enable:showCursor:"), display, primary, enable, cursor)
}
func (s SLSharingSession) ShowPicker() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("showPicker"))
}
func (s SLSharingSession) InitFromUUID(uuid objectivec.IObject) SLSharingSession {
	rv := objc.SendIfResponds[SLSharingSession](s.ID, objc.Sel("initFromUUID:"), uuid)
	return rv
}
func (s SLSharingSession) InitWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) SLSharingSession {
	rv := objc.SendIfResponds[SLSharingSession](s.ID, objc.Sel("initWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:"), title, indicator, notifications)
	return rv
}
func (s SLSharingSession) InitWithUUIDTitleType(uuid objectivec.IObject, title objectivec.IObject, type_ int) SLSharingSession {
	rv := objc.SendIfResponds[SLSharingSession](s.ID, objc.Sel("initWithUUID:title:type:"), uuid, title, type_)
	return rv
}

func (s SLSharingSession) LifetimePort() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("lifetimePort"))
	return rv
}
func (s SLSharingSession) SetLifetimePort(value uint32) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setLifetimePort:"), value)
}
func (s SLSharingSession) Title() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSharingSession) Type() int {
	rv := objc.SendIfResponds[int](s.ID, objc.Sel("type"))
	return rv
}
func (s SLSharingSession) SetType(value int) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setType:"), value)
}
func (s SLSharingSession) Uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (s SLSharingSession) Uuid_internal() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("uuid_internal"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (s SLSharingSession) SetUuid_internal(value foundation.NSUUID) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setUuid_internal:"), value)
}
