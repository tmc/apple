// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

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
	rv := objc.Send[SLSharingSession](objc.ID(sc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession
type ISLSharingSession interface {
	objectivec.IObject

	// Topic: Methods

	Content() objectivec.IObject
	GetUUID() objectivec.IObject
	GetUUIDBytes() objectivec.IObject
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
	rv := objc.Send[SLSharingSession](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSharingSession) Autorelease() SLSharingSession {
	rv := objc.Send[SLSharingSession](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSharingSession creates a new SLSharingSession instance.
func NewSLSharingSession() SLSharingSession {
	class := getSLSharingSessionClass()
	rv := objc.Send[SLSharingSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/initFromUUID:
func NewSLSharingSessionFromUUID(uuid objectivec.IObject) SLSharingSession {
	instance := getSLSharingSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromUUID:"), uuid)
	return SLSharingSessionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/initWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:
func NewSLSharingSessionWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) SLSharingSession {
	instance := getSLSharingSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:"), title, indicator, notifications)
	return SLSharingSessionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/initWithUUID:title:type:
func NewSLSharingSessionWithUUIDTitleType(uuid objectivec.IObject, title objectivec.IObject, type_ int) SLSharingSession {
	instance := getSLSharingSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUUID:title:type:"), uuid, title, type_)
	return SLSharingSessionFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/content
func (s SLSharingSession) Content() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("content"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/getUUID
func (s SLSharingSession) GetUUID() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("getUUID"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/getUUIDBytes
func (s SLSharingSession) GetUUIDBytes() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("getUUIDBytes"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/isEqualToSharingSession:
func (s SLSharingSession) IsEqualToSharingSession(session objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEqualToSharingSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/setContent:
func (s SLSharingSession) SetContent(content objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setContent:"), content)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/setPresentationDisplay:primary:enable:showCursor:
func (s SLSharingSession) SetPresentationDisplayPrimaryEnableShowCursor(display objectivec.IObject, primary objectivec.IObject, enable bool, cursor bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setPresentationDisplay:primary:enable:showCursor:"), display, primary, enable, cursor)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/showPicker
func (s SLSharingSession) ShowPicker() {
	objc.Send[objc.ID](s.ID, objc.Sel("showPicker"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/initFromUUID:
func (s SLSharingSession) InitFromUUID(uuid objectivec.IObject) SLSharingSession {
	rv := objc.Send[SLSharingSession](s.ID, objc.Sel("initFromUUID:"), uuid)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/initWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:
func (s SLSharingSession) InitWithTitleSuppressWindowSharingIndicatorSuppressMenuBarSharingIndicatorNotifications(title objectivec.IObject, indicator bool, notifications bool) SLSharingSession {
	rv := objc.Send[SLSharingSession](s.ID, objc.Sel("initWithTitle:suppressWindowSharingIndicator:suppressMenuBarSharingIndicatorNotifications:"), title, indicator, notifications)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/initWithUUID:title:type:
func (s SLSharingSession) InitWithUUIDTitleType(uuid objectivec.IObject, title objectivec.IObject, type_ int) SLSharingSession {
	rv := objc.Send[SLSharingSession](s.ID, objc.Sel("initWithUUID:title:type:"), uuid, title, type_)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/lifetimePort
func (s SLSharingSession) LifetimePort() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("lifetimePort"))
	return rv
}
func (s SLSharingSession) SetLifetimePort(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLifetimePort:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/title
func (s SLSharingSession) Title() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/type
func (s SLSharingSession) Type() int {
	rv := objc.Send[int](s.ID, objc.Sel("type"))
	return rv
}
func (s SLSharingSession) SetType(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setType:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/uuid
func (s SLSharingSession) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSharingSession/uuid_internal
func (s SLSharingSession) Uuid_internal() foundation.NSUUID {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uuid_internal"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (s SLSharingSession) SetUuid_internal(value foundation.NSUUID) {
	objc.Send[struct{}](s.ID, objc.Sel("setUuid_internal:"), value)
}
