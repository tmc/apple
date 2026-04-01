// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationAttributedMessageContext] class.
var (
	_UNNotificationAttributedMessageContextClass     UNNotificationAttributedMessageContextClass
	_UNNotificationAttributedMessageContextClassOnce sync.Once
)

func getUNNotificationAttributedMessageContextClass() UNNotificationAttributedMessageContextClass {
	_UNNotificationAttributedMessageContextClassOnce.Do(func() {
		_UNNotificationAttributedMessageContextClass = UNNotificationAttributedMessageContextClass{class: objc.GetClass("UNNotificationAttributedMessageContext")}
	})
	return _UNNotificationAttributedMessageContextClass
}

// GetUNNotificationAttributedMessageContextClass returns the class object for UNNotificationAttributedMessageContext.
func GetUNNotificationAttributedMessageContextClass() UNNotificationAttributedMessageContextClass {
	return getUNNotificationAttributedMessageContextClass()
}

type UNNotificationAttributedMessageContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationAttributedMessageContextClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationAttributedMessageContextClass) Alloc() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext
type UNNotificationAttributedMessageContext struct {
	objectivec.Object
}

// UNNotificationAttributedMessageContextFromID constructs a [UNNotificationAttributedMessageContext] from an objc.ID.
func UNNotificationAttributedMessageContextFromID(id objc.ID) UNNotificationAttributedMessageContext {
	return UNNotificationAttributedMessageContext{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationAttributedMessageContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationAttributedMessageContext] class.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext
type IUNNotificationAttributedMessageContext interface {
	objectivec.IObject
	UNNotificationContentProviding
}

// Init initializes the instance.
func (u UNNotificationAttributedMessageContext) Init() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationAttributedMessageContext) Autorelease() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationAttributedMessageContext creates a new UNNotificationAttributedMessageContext instance.
func NewUNNotificationAttributedMessageContext() UNNotificationAttributedMessageContext {
	class := getUNNotificationAttributedMessageContextClass()
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext/init(sendMessageIntent:attributedContent:)
func NewUNNotificationAttributedMessageContextWithSendMessageIntentAttributedContent(sendMessageIntent unsafe.Pointer, attributedContent foundation.NSAttributedString) UNNotificationAttributedMessageContext {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationAttributedMessageContextClass().class), objc.Sel("contextWithSendMessageIntent:attributedContent:"), sendMessageIntent, attributedContent)
	return UNNotificationAttributedMessageContextFromID(rv)
}

// Protocol methods for UNNotificationContentProviding
