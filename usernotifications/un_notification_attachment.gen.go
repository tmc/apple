// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationAttachment] class.
var (
	_UNNotificationAttachmentClass     UNNotificationAttachmentClass
	_UNNotificationAttachmentClassOnce sync.Once
)

func getUNNotificationAttachmentClass() UNNotificationAttachmentClass {
	_UNNotificationAttachmentClassOnce.Do(func() {
		_UNNotificationAttachmentClass = UNNotificationAttachmentClass{class: objc.GetClass("UNNotificationAttachment")}
	})
	return _UNNotificationAttachmentClass
}

// GetUNNotificationAttachmentClass returns the class object for UNNotificationAttachment.
func GetUNNotificationAttachmentClass() UNNotificationAttachmentClass {
	return getUNNotificationAttachmentClass()
}

type UNNotificationAttachmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationAttachmentClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationAttachmentClass) Alloc() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A media file associated with a notification.
//
// # Overview
//
// Create a [UNNotificationAttachment] object when you want to include audio,
// image, or video content together in an alert-based notification. When
// creating the [UNNotificationAttachment] object, the file you specify must
// be on disk, and the file format must be one of the supported types.
//
// You’re responsible for supplying attachments before the system displays
// your notification’s alert. For local notifications, add attachments when
// creating the notification’s content. For remote notifications, use a
// notification service app extension to download the attached files and then
// add them to the notification’s content before delivery.
//
// The system validates attachments before displaying the associated
// notification. If you attach a file to a local notification request that’s
// corrupted, invalid, or of an unsupported file type, the system doesn’t
// schedule your request. For remote notifications, the system validates
// attachments after your notification service app extension finishes. Once
// validated, the system moves the attached files into the attachment data
// store so that the appropriate processes can access the files. The system
// copies attachments located inside an app’s bundle.
//
// # Supported File Types
//
// Table 1 lists the types of files you can include as an attachment and the
// supported file formats. The table also lists the maximum size allowed for
// attachments of each type. An image file may contain a static image or an
// animated image sequence.
//
// Table 1. Supported attachment file types
//
// [Table data omitted]
//
// When creating an attachment, you can specify optional details about how to
// present the thumbnail image for the image or movie. Use the
// [UNNotificationAttachmentOptionsThumbnailClippingRectKey] option to use
// only the specified portion of an image as a thumbnail. For animated images
// and movies, use the [UNNotificationAttachmentOptionsThumbnailTimeKey]
// option to select which frame to use for the thumbnail image.
//
// The system limits the amount of storage space allocated for attachments for
// each app. To delete attachments, use the methods of the
// [UNUserNotificationCenter] class to remove the notification requests that
// contain those attachments.
//
// # Getting the Attachment Contents
//
//   - [UNNotificationAttachment.Identifier]: The unique identifier for the attachment.
//   - [UNNotificationAttachment.URL]: The URL of the file for this attachment.
//   - [UNNotificationAttachment.Type]: The UTI type of the attachment.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment
//
// [UNNotificationAttachmentOptionsThumbnailClippingRectKey]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachmentOptionsThumbnailClippingRectKey
// [UNNotificationAttachmentOptionsThumbnailTimeKey]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachmentOptionsThumbnailTimeKey
type UNNotificationAttachment struct {
	objectivec.Object
}

// UNNotificationAttachmentFromID constructs a [UNNotificationAttachment] from an objc.ID.
//
// A media file associated with a notification.
func UNNotificationAttachmentFromID(id objc.ID) UNNotificationAttachment {
	return UNNotificationAttachment{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationAttachment] class.
//
// # Getting the Attachment Contents
//
//   - [IUNNotificationAttachment.Identifier]: The unique identifier for the attachment.
//   - [IUNNotificationAttachment.URL]: The URL of the file for this attachment.
//   - [IUNNotificationAttachment.Type]: The UTI type of the attachment.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment
type IUNNotificationAttachment interface {
	objectivec.IObject

	// Topic: Getting the Attachment Contents

	// The unique identifier for the attachment.
	Identifier() string
	// The URL of the file for this attachment.
	URL() foundation.NSURL
	// The UTI type of the attachment.
	Type() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationAttachment) Init() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationAttachment) Autorelease() UNNotificationAttachment {
	rv := objc.Send[UNNotificationAttachment](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationAttachment creates a new UNNotificationAttachment instance.
func NewUNNotificationAttachment() UNNotificationAttachment {
	class := getUNNotificationAttachmentClass()
	rv := objc.Send[UNNotificationAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an attachment object from the specified file and options.
//
// identifier: The unique identifier of the attachment. Use this string to identify the
// attachment later. If you specify an empty string, this method creates a
// unique identifier string for you.
//
// URL: The URL of the file you want to attach to the notification. The URL must be
// a file URL and the file must be readable by the current process. This
// parameter must not be `nil`. For a list of supported file types, see
// [UNNotificationAttachment].
//
// options: A dictionary of options related to the attached file. Use the options to
// specify meta information about the attachment, such as the clipping
// rectangle to use for the resulting thumbnail.
//
// # Return Value
//
// An attachment object containing information about the specified file or
// `nil` if the attachment could not be created.
//
// # Discussion
//
// This method verifies that the specified file is readable and that the file
// format is one of the supported types. When errors occur, the method
// provides an appropriate `error` object.
//
// When you schedule a notification request containing the attachment, the
// system moves the attachment’s file to a new location to facilitate access
// by the appropriate processes. After the move, the only way to access the
// file is using the methods of the [UNUserNotificationCenter] object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/init(identifier:url:options:)-83grx
func NewUNNotificationAttachmentWithIdentifierURLOptionsError(identifier string, URL foundation.NSURL, options foundation.INSDictionary) (UNNotificationAttachment, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationAttachmentClass().class), objc.Sel("attachmentWithIdentifier:URL:options:error:"), objc.String(identifier), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return UNNotificationAttachment{}, foundation.NSErrorFrom(errorPtr)
	}
	return UNNotificationAttachmentFromID(rv), nil
}

func (u UNNotificationAttachment) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The unique identifier for the attachment.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/identifier
func (u UNNotificationAttachment) Identifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The URL of the file for this attachment.
//
// # Discussion
//
// The file at the specified URL is security scoped to your app. Before you
// access it, call the [startAccessingSecurityScopedResource()] method of
// [NSURL].
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/url
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [startAccessingSecurityScopedResource()]: https://developer.apple.com/documentation/Foundation/URL/startAccessingSecurityScopedResource()
func (u UNNotificationAttachment) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The UTI type of the attachment.
//
// # Discussion
//
// The system derives the value of this property from the attachment data.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachment/type
func (u UNNotificationAttachment) Type() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
