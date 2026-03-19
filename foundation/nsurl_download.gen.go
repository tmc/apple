// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSURLDownload] class.
var (
	_NSURLDownloadClass     NSURLDownloadClass
	_NSURLDownloadClassOnce sync.Once
)

func getNSURLDownloadClass() NSURLDownloadClass {
	_NSURLDownloadClassOnce.Do(func() {
		_NSURLDownloadClass = NSURLDownloadClass{class: objc.GetClass("NSURLDownload")}
	})
	return _NSURLDownloadClass
}

// GetNSURLDownloadClass returns the class object for NSURLDownload.
func GetNSURLDownloadClass() NSURLDownloadClass {
	return getNSURLDownloadClass()
}

type NSURLDownloadClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSURLDownloadClass) Alloc() NSURLDownload {
	rv := objc.Send[NSURLDownload](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that downloads a resource asynchronously and saves the data to a
// file.
//
// # Overview
// 
// The interface for [NSURLDownload] provides methods to initialize a
// download, set the destination path and cancel loading the request.
// 
// The delegate object assigned to each instance of this class should
// implement the methods defined by the [NSURLDownloadDelegate] protocol.
// These methods provide the delegate with the current status of in-progress
// asynchronous downloads and allow the delegate to customize the URL loading
// process. These delegate methods are called on the thread that started the
// asynchronous load operation for the associated [NSURLDownload] object.
//
// # Creating and configuring a download instance
//
//   - [NSURLDownload.SetDestinationAllowOverwrite]: Sets the destination path of the downloaded file.
//
// # Resuming partial downloads
//
//   - [NSURLDownload.ResumeData]: Returns the resume data for a download that is not yet complete.
//   - [NSURLDownload.DeletesFileUponFailure]: Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//   - [NSURLDownload.SetDeletesFileUponFailure]
//
// # Canceling a download
//
//   - [NSURLDownload.Cancel]: Cancels the receiver’s download and deletes the downloaded file.
//
// # Getting download properties
//
//   - [NSURLDownload.Request]: Returns the request that initiated the receiver’s download.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload
type NSURLDownload struct {
	objectivec.Object
}

// NSURLDownloadFromID constructs a [NSURLDownload] from an objc.ID.
//
// An object that downloads a resource asynchronously and saves the data to a
// file.
func NSURLDownloadFromID(id objc.ID) NSURLDownload {
	return NSURLDownload{objectivec.Object{ID: id}}
}
// NOTE: NSURLDownload adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSURLDownload] class.
//
// # Creating and configuring a download instance
//
//   - [INSURLDownload.SetDestinationAllowOverwrite]: Sets the destination path of the downloaded file.
//
// # Resuming partial downloads
//
//   - [INSURLDownload.ResumeData]: Returns the resume data for a download that is not yet complete.
//   - [INSURLDownload.DeletesFileUponFailure]: Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
//   - [INSURLDownload.SetDeletesFileUponFailure]
//
// # Canceling a download
//
//   - [INSURLDownload.Cancel]: Cancels the receiver’s download and deletes the downloaded file.
//
// # Getting download properties
//
//   - [INSURLDownload.Request]: Returns the request that initiated the receiver’s download.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload
type INSURLDownload interface {
	objectivec.IObject

	// Topic: Creating and configuring a download instance

	// Sets the destination path of the downloaded file.
	SetDestinationAllowOverwrite(path string, allowOverwrite bool)

	// Topic: Resuming partial downloads

	// Returns the resume data for a download that is not yet complete.
	ResumeData() INSData
	// Returns whether the receiver deletes partially downloaded files when a download stops prematurely.
	DeletesFileUponFailure() bool
	SetDeletesFileUponFailure(value bool)

	// Topic: Canceling a download

	// Cancels the receiver’s download and deletes the downloaded file.
	Cancel()

	// Topic: Getting download properties

	// Returns the request that initiated the receiver’s download.
	Request() INSURLRequest
}

// Init initializes the instance.
func (u NSURLDownload) Init() NSURLDownload {
	rv := objc.Send[NSURLDownload](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSURLDownload) Autorelease() NSURLDownload {
	rv := objc.Send[NSURLDownload](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSURLDownload creates a new NSURLDownload instance.
func NewNSURLDownload() NSURLDownload {
	class := getNSURLDownloadClass()
	rv := objc.Send[NSURLDownload](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the destination path of the downloaded file.
//
// path: The path for the downloaded file.
//
// allowOverwrite: [true] if an existing file at `path` can be replaced, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If `allowOverwrite` is [false] and a file already exists at `path`, a
// unique filename will be created for the downloaded file by appending a
// number to the filename. The delegate can implement the
// [DownloadDidCreateDestination] delegate method to determine the filename
// used when the file is written to disk.
// 
// # Special Considerations
// 
// An [NSURLDownload] instance ignores multiple calls to this method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload/setDestination(_:allowOverwrite:)
func (u NSURLDownload) SetDestinationAllowOverwrite(path string, allowOverwrite bool) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDestination:allowOverwrite:"), objc.String(path), allowOverwrite)
}
// Cancels the receiver’s download and deletes the downloaded file.
//
// # Discussion
// 
// This method deletes the partially downloaded file unless you have
// previously called [DeletesFileUponFailure], passing [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload/cancel()
func (u NSURLDownload) Cancel() {
	objc.Send[objc.ID](u.ID, objc.Sel("cancel"))
}

// Returns whether a URL download object can resume a download that was
// decoded with the specified MIME type.
//
// MIMEType: The MIME type the caller wants to know about.
//
// # Return Value
// 
// [true] if the URL download object can resume a download that was decoded
// with the specified MIME type, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The MIME type of a file, in conjunction with the value returned by the
// [DownloadShouldDecodeSourceDataOfMIMEType] delegate method, determines
// whether the [NSURLDownload] class should decode or decompress the incoming
// data as it is received.
// 
// Some compression techniques, such as the [DEFLATE] algorithm (`gzip`) use
// symbol dictionaries that vary during the compression process, making it
// impractical to decompress only a portion of the data starting in the
// middle. For this reason, this method returns [false] unless both of the
// following conditions are met:
// 
// - The MIME type is of a type that the [NSURLDownload] class knows how to
// decompress or decode. - The decoding can be safely resumed.
// 
// In practice, this method returns [true] for MacBinary and BinHex, otherwise
// [false].
// 
// If your app needs to be able to resume file downloads in `gzip` format,
// your [DownloadShouldDecodeSourceDataOfMIMEType] method must return [false],
// and you must decode the resulting file yourself after you finish
// downloading it in its entirety.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload/canResumeDownloadDecoded(withEncodingMIMEType:)
func (_NSURLDownloadClass NSURLDownloadClass) CanResumeDownloadDecodedWithEncodingMIMEType(MIMEType string) bool {
	rv := objc.Send[bool](objc.ID(_NSURLDownloadClass.class), objc.Sel("canResumeDownloadDecodedWithEncodingMIMEType:"), objc.String(MIMEType))
	return rv
}

// Returns the resume data for a download that is not yet complete.
//
// # Return Value
// 
// The resume data for a download that is not yet complete. This data
// represents the necessary state information that an [NSURLDownload] object
// needs to resume a download. The resume data can later be used when
// initializing a download with [init(resumeData:delegate:path:)]. Returns
// `nil` if the download is not able to be resumed.
// 
// # Discussion
// 
// Resume data is returned only if both the protocol and the server support
// resuming. For details on how to resume a connection, see the documentation
// for [init(resumeData:delegate:path:)].
//
// [init(resumeData:delegate:path:)]: https://developer.apple.com/documentation/Foundation/NSURLDownload/init(resumeData:delegate:path:)
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload/resumeData
func (u NSURLDownload) ResumeData() INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("resumeData"))
	return NSDataFromID(objc.ID(rv))
}
// Returns the request that initiated the receiver’s download.
//
// # Return Value
// 
// The URL request that initiated the receiver’s download.
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload/request
func (u NSURLDownload) Request() INSURLRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("request"))
	return NSURLRequestFromID(objc.ID(rv))
}
// Returns whether the receiver deletes partially downloaded files when a
// download stops prematurely.
//
// # Return Value
// 
// [true] if partially downloaded files should be deleted when a download
// stops prematurely, [false] otherwise. The default is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLDownload/deletesFileUponFailure
func (u NSURLDownload) DeletesFileUponFailure() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("deletesFileUponFailure"))
	return rv
}
func (u NSURLDownload) SetDeletesFileUponFailure(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setDeletesFileUponFailure:"), value)
}

