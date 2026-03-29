// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationActionIcon] class.
var (
	_UNNotificationActionIconClass     UNNotificationActionIconClass
	_UNNotificationActionIconClassOnce sync.Once
)

func getUNNotificationActionIconClass() UNNotificationActionIconClass {
	_UNNotificationActionIconClassOnce.Do(func() {
		_UNNotificationActionIconClass = UNNotificationActionIconClass{class: objc.GetClass("UNNotificationActionIcon")}
	})
	return _UNNotificationActionIconClass
}

// GetUNNotificationActionIconClass returns the class object for UNNotificationActionIcon.
func GetUNNotificationActionIconClass() UNNotificationActionIconClass {
	return getUNNotificationActionIconClass()
}

type UNNotificationActionIconClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationActionIconClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationActionIconClass) Alloc() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An icon associated with an action.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon
type UNNotificationActionIcon struct {
	objectivec.Object
}

// UNNotificationActionIconFromID constructs a [UNNotificationActionIcon] from an objc.ID.
//
// An icon associated with an action.
func UNNotificationActionIconFromID(id objc.ID) UNNotificationActionIcon {
	return UNNotificationActionIcon{objectivec.Object{ID: id}}
}
// NOTE: UNNotificationActionIcon adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationActionIcon] class.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon
type IUNNotificationActionIcon interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationActionIcon) Init() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationActionIcon) Autorelease() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationActionIcon creates a new UNNotificationActionIcon instance.
func NewUNNotificationActionIcon() UNNotificationActionIcon {
	class := getUNNotificationActionIconClass()
	rv := objc.Send[UNNotificationActionIcon](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an action icon by using a system symbol image.
//
// systemImageName: The name of the system symbol image. Use the SF Symbols app to look up the
// names of system symbol images. Download this app from the design resources
// page at [developer.apple.com].
// //
// [developer.apple.com]: https://developer.apple.com/design/resources/
//
// # Return Value
// 
// An action icon that the system initializes with the system symbol image
// that your app specifies.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(systemImageName:)
func NewUNNotificationActionIconWithSystemImageName(systemImageName string) UNNotificationActionIcon {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationActionIconClass().class), objc.Sel("iconWithSystemImageName:"), objc.String(systemImageName))
	return UNNotificationActionIconFromID(rv)
}

// Creates an action icon based on an image in your app’s bundle, preferably
// in an asset catalog.
//
// templateImageName: The name of a custom image in the app’s asset catalog. If the image
// isn’t in your app’s asset catalog, this method searches the app bundle
// for the image.
// 
// You don’t need to specify the filename extension or the `@2x` or `@3x`
// modifiers for this name. This method retrieves the appropriate image based
// on the system and the available image resources.
//
// # Return Value
// 
// An action icon initialized with the specified template image provided by
// your app.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(templateImageName:)
func NewUNNotificationActionIconWithTemplateImageName(templateImageName string) UNNotificationActionIcon {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationActionIconClass().class), objc.Sel("iconWithTemplateImageName:"), objc.String(templateImageName))
	return UNNotificationActionIconFromID(rv)
}

func (u UNNotificationActionIcon) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

