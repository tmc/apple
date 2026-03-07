// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSharingService] class.
var (
	_NSSharingServiceClass     NSSharingServiceClass
	_NSSharingServiceClassOnce sync.Once
)

func getNSSharingServiceClass() NSSharingServiceClass {
	_NSSharingServiceClassOnce.Do(func() {
		_NSSharingServiceClass = NSSharingServiceClass{class: objc.GetClass("NSSharingService")}
	})
	return _NSSharingServiceClass
}

// GetNSSharingServiceClass returns the class object for NSSharingService.
func GetNSSharingServiceClass() NSSharingServiceClass {
	return getNSSharingServiceClass()
}

type NSSharingServiceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSharingServiceClass) Alloc() NSSharingService {
	rv := objc.Send[NSSharingService](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that facilitates the sharing of content with social media
// services, or with apps like Mail or Safari.
//
// # Overview
// 
// An [NSSharingService] object provides a consistent user experience for
// sharing items—[NSURL] objects, [NSString] objects, [NSImage] objects,
// video (through file URLs), of any object that implements the
// [NSPasteboardWriting] protocol—in macOS.
// 
// For any item or group of items, the [NSSharingService] displays a sheet
// with the content to share. A sharing service can create a post on a social
// network like Twitter or Facebook, send a message by email or iMessage,
// upload videos to viewing services, or send a file using AirDrop.
// 
// You can use [NSSharingService] objects directly in your app. The following
// example shows how to create a button that shares content directly to a
// social media service.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// # Creating a Sharing Service
//
//   - [NSSharingService.InitWithTitleImageAlternateImageHandler]: Creates a custom sharing service object.
//
// # Managing the Delegate
//
//   - [NSSharingService.Delegate]: Specifies the delegate of the sharing service.
//   - [NSSharingService.SetDelegate]
//
// # Querying Service Availability
//
//   - [NSSharingService.CanPerformWithItems]: Returns whether the service can share all the specified items.
//
// # Getting the Service’s Details
//
//   - [NSSharingService.AccountName]: The account name used for posting on Twitter or Sina Weibo.
//   - [NSSharingService.AlternateImage]: The alternate image representing the sharing service.
//   - [NSSharingService.Image]: The primary image representing the sharing service.
//   - [NSSharingService.Title]: The title of the sharing service.
//
// # Configuring the Service
//
//   - [NSSharingService.MenuItemTitle]: The title of the service in the Share menu.
//   - [NSSharingService.SetMenuItemTitle]
//   - [NSSharingService.Recipients]: An array containing the user handles of the desired recipients.
//   - [NSSharingService.SetRecipients]
//   - [NSSharingService.Subject]: The subject of the post.
//   - [NSSharingService.SetSubject]
//
// # Sharing Items
//
//   - [NSSharingService.PerformWithItems]: Manually performs the service on the provided items.
//
// # Getting the Shared Items
//
//   - [NSSharingService.AttachmentFileURLs]: An array of NSURL objects representing the files that were shared.
//   - [NSSharingService.MessageBody]: The message body as a string.
//   - [NSSharingService.PermanentLink]: A permanent URL (permalink) that your app can use to access the post.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService
type NSSharingService struct {
	objectivec.Object
}

// NSSharingServiceFromID constructs a [NSSharingService] from an objc.ID.
//
// An object that facilitates the sharing of content with social media
// services, or with apps like Mail or Safari.
func NSSharingServiceFromID(id objc.ID) NSSharingService {
	return NSSharingService{objectivec.Object{id}}
}
// NOTE: NSSharingService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSSharingService] class.
//
// # Creating a Sharing Service
//
//   - [INSSharingService.InitWithTitleImageAlternateImageHandler]: Creates a custom sharing service object.
//
// # Managing the Delegate
//
//   - [INSSharingService.Delegate]: Specifies the delegate of the sharing service.
//   - [INSSharingService.SetDelegate]
//
// # Querying Service Availability
//
//   - [INSSharingService.CanPerformWithItems]: Returns whether the service can share all the specified items.
//
// # Getting the Service’s Details
//
//   - [INSSharingService.AccountName]: The account name used for posting on Twitter or Sina Weibo.
//   - [INSSharingService.AlternateImage]: The alternate image representing the sharing service.
//   - [INSSharingService.Image]: The primary image representing the sharing service.
//   - [INSSharingService.Title]: The title of the sharing service.
//
// # Configuring the Service
//
//   - [INSSharingService.MenuItemTitle]: The title of the service in the Share menu.
//   - [INSSharingService.SetMenuItemTitle]
//   - [INSSharingService.Recipients]: An array containing the user handles of the desired recipients.
//   - [INSSharingService.SetRecipients]
//   - [INSSharingService.Subject]: The subject of the post.
//   - [INSSharingService.SetSubject]
//
// # Sharing Items
//
//   - [INSSharingService.PerformWithItems]: Manually performs the service on the provided items.
//
// # Getting the Shared Items
//
//   - [INSSharingService.AttachmentFileURLs]: An array of NSURL objects representing the files that were shared.
//   - [INSSharingService.MessageBody]: The message body as a string.
//   - [INSSharingService.PermanentLink]: A permanent URL (permalink) that your app can use to access the post.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService
type INSSharingService interface {
	objectivec.IObject

	// Topic: Creating a Sharing Service

	// Creates a custom sharing service object.
	InitWithTitleImageAlternateImageHandler(title string, image INSImage, alternateImage INSImage, block VoidHandler) NSSharingService

	// Topic: Managing the Delegate

	// Specifies the delegate of the sharing service.
	Delegate() NSSharingServiceDelegate
	SetDelegate(value NSSharingServiceDelegate)

	// Topic: Querying Service Availability

	// Returns whether the service can share all the specified items.
	CanPerformWithItems(items foundation.INSArray) bool

	// Topic: Getting the Service’s Details

	// The account name used for posting on Twitter or Sina Weibo.
	AccountName() string
	// The alternate image representing the sharing service.
	AlternateImage() INSImage
	// The primary image representing the sharing service.
	Image() INSImage
	// The title of the sharing service.
	Title() string

	// Topic: Configuring the Service

	// The title of the service in the Share menu.
	MenuItemTitle() string
	SetMenuItemTitle(value string)
	// An array containing the user handles of the desired recipients.
	Recipients() []string
	SetRecipients(value []string)
	// The subject of the post.
	Subject() string
	SetSubject(value string)

	// Topic: Sharing Items

	// Manually performs the service on the provided items.
	PerformWithItems(items foundation.INSArray)

	// Topic: Getting the Shared Items

	// An array of NSURL objects representing the files that were shared.
	AttachmentFileURLs() []foundation.NSURL
	// The message body as a string.
	MessageBody() string
	// A permanent URL (permalink) that your app can use to access the post.
	PermanentLink() foundation.INSURL
}





// Init initializes the instance.
func (s NSSharingService) Init() NSSharingService {
	rv := objc.Send[NSSharingService](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSharingService) Autorelease() NSSharingService {
	rv := objc.Send[NSSharingService](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSharingService creates a new NSSharingService instance.
func NewNSSharingService() NSSharingService {
	class := getNSSharingServiceClass()
	rv := objc.Send[NSSharingService](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Returns a sharing service instance representing the specified service name.
//
// serviceName: The service name. The possible system provided values are listed in
// `Available Sharing Services`.
//
// # Return Value
// 
// An instance of [NSSharingService] for the specified service name.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/init(named:)
func NewSharingServiceNamed(serviceName NSSharingServiceName) NSSharingService {
	rv := objc.Send[objc.ID](objc.ID(getNSSharingServiceClass().class), objc.Sel("sharingServiceNamed:"), objc.String(string(serviceName)))
	return NSSharingServiceFromID(rv)
}







// Creates a custom sharing service object.
//
// title: The custom sharing service name.
//
// image: The image that represents the sharing service
//
// alternateImage: The alternate image that represents the sharing service
//
// block: The block that actually interacts with the service.
//
// # Return Value
// 
// An instance of the custom sharing object.
//
// # Discussion
// 
// Custom sharing services can be added to the [NSSharingServicePicker] with
// the [SharingServicePickerSharingServicesForItemsProposedSharingServices]
// delegate method.
// 
// When implementing this method, consider subclassing [NSSharingService] so
// the [CanPerformWithItems] and [sharingServices(forItems:)] can provide
// accurate results.
//
// [sharingServices(forItems:)]: https://developer.apple.com/documentation/AppKit/NSSharingService/sharingServices(forItems:)
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/init(title:image:alternateImage:handler:)
func (s NSSharingService) InitWithTitleImageAlternateImageHandler(title string, image INSImage, alternateImage INSImage, block VoidHandler) NSSharingService {
		_block3, _cleanup3 := NewVoidBlock(block)
	defer _cleanup3()
		rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithTitle:image:alternateImage:handler:"), objc.String(title), image, alternateImage, _block3)
	return NSSharingServiceFromID(rv)
}

// Returns whether the service can share all the specified items.
//
// items: The items to share.
//
// # Return Value
// 
// [true] if the service can share all the items; [false] otherwise. If
// `items` is `nil`, the method will return [true] when the service is
// configured.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method can be used to validate a custom user interface such as a
// dedicated Twitter button. Therefore you could call it once at launch time
// with `nil` items to check whether to display the button or not, and then
// with real items to enable and disable the button depending on the context
// or selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/canPerform(withItems:)
func (s NSSharingService) CanPerformWithItems(items foundation.INSArray) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canPerformWithItems:"), items)
	return rv
}

// Manually performs the service on the provided items.
//
// items: The items to share.
//
// # Discussion
// 
// In most cases this will display a sharing window.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/perform(withItems:)
func (s NSSharingService) PerformWithItems(items foundation.INSArray) {
	objc.Send[objc.ID](s.ID, objc.Sel("performWithItems:"), items)
}












// Specifies the delegate of the sharing service.
//
// # Discussion
// 
// The delegate class must conform to the [NSSharingServiceDelegate] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/delegate
func (s NSSharingService) Delegate() NSSharingServiceDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSSharingServiceDelegateObjectFromID(rv)
}
func (s NSSharingService) SetDelegate(value NSSharingServiceDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}



// The account name used for posting on Twitter or Sina Weibo.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/accountName
func (s NSSharingService) AccountName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("accountName"))
	return foundation.NSStringFromID(rv).String()
}



// The alternate image representing the sharing service.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/alternateImage
func (s NSSharingService) AlternateImage() INSImage {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("alternateImage"))
	return NSImageFromID(objc.ID(rv))
}



// The primary image representing the sharing service.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/image
func (s NSSharingService) Image() INSImage {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}



// The title of the sharing service.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/title
func (s NSSharingService) Title() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}



// The title of the service in the Share menu.
//
// # Discussion
// 
// By default, this title is the same as the value of the [Title] property.
// Your app can modify this value.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/menuItemTitle
func (s NSSharingService) MenuItemTitle() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("menuItemTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSharingService) SetMenuItemTitle(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setMenuItemTitle:"), objc.String(value))
}



// An array containing the user handles of the desired recipients.
//
// # Discussion
// 
// Each object in the array is an [NSString] object that contains the handle
// of a single recipient. The specific format of these handle varies from
// service to service. For example, some services use email addresses as
// handles.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/recipients
func (s NSSharingService) Recipients() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("recipients"))
	return objc.ConvertSliceToStrings(rv)
}
func (s NSSharingService) SetRecipients(value []string) {
	objc.Send[struct{}](s.ID, objc.Sel("setRecipients:"), objectivec.StringSliceToNSArray(value))
}



// The subject of the post.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/subject
func (s NSSharingService) Subject() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("subject"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSharingService) SetSubject(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setSubject:"), objc.String(value))
}



// An array of NSURL objects representing the files that were shared.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/attachmentFileURLs
func (s NSSharingService) AttachmentFileURLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("attachmentFileURLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}



// The message body as a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/messageBody
func (s NSSharingService) MessageBody() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("messageBody"))
	return foundation.NSStringFromID(rv).String()
}



// A permanent URL (permalink) that your app can use to access the post.
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingService/permanentLink
func (s NSSharingService) PermanentLink() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("permanentLink"))
	return foundation.NSURLFromID(objc.ID(rv))
}


















// InitWithTitleImageAlternateImageHandlerSync is a synchronous wrapper around [NSSharingService.InitWithTitleImageAlternateImageHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSharingService) InitWithTitleImageAlternateImageHandlerSync(ctx context.Context, title string, image INSImage, alternateImage INSImage) error {
	done := make(chan struct{}, 1)
	s.InitWithTitleImageAlternateImageHandler(title, image, alternateImage, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






