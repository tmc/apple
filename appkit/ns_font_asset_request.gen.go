// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFontAssetRequest] class.
var (
	_NSFontAssetRequestClass     NSFontAssetRequestClass
	_NSFontAssetRequestClassOnce sync.Once
)

func getNSFontAssetRequestClass() NSFontAssetRequestClass {
	_NSFontAssetRequestClassOnce.Do(func() {
		_NSFontAssetRequestClass = NSFontAssetRequestClass{class: objc.GetClass("NSFontAssetRequest")}
	})
	return _NSFontAssetRequestClass
}

// GetNSFontAssetRequestClass returns the class object for NSFontAssetRequest.
func GetNSFontAssetRequestClass() NSFontAssetRequestClass {
	return getNSFontAssetRequestClass()
}

type NSFontAssetRequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFontAssetRequestClass) Alloc() NSFontAssetRequest {
	rv := objc.Send[NSFontAssetRequest](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

//
// # Creating a Font Asset Request
//
//   - [NSFontAssetRequest.InitWithFontDescriptorsOptions]
//
// # Downloading a Font Asset
//
//   - [NSFontAssetRequest.DownloadFontAssetsWithCompletionHandler]
//   - [NSFontAssetRequest.DownloadedFontDescriptors]
//
// # Getting the Download Progress
//
//   - [NSFontAssetRequest.Progress]
// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest
type NSFontAssetRequest struct {
	objectivec.Object
}

// NSFontAssetRequestFromID constructs a [NSFontAssetRequest] from an objc.ID.
func NSFontAssetRequestFromID(id objc.ID) NSFontAssetRequest {
	return NSFontAssetRequest{objectivec.Object{ID: id}}
}
// NOTE: NSFontAssetRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFontAssetRequest] class.
//
// # Creating a Font Asset Request
//
//   - [INSFontAssetRequest.InitWithFontDescriptorsOptions]
//
// # Downloading a Font Asset
//
//   - [INSFontAssetRequest.DownloadFontAssetsWithCompletionHandler]
//   - [INSFontAssetRequest.DownloadedFontDescriptors]
//
// # Getting the Download Progress
//
//   - [INSFontAssetRequest.Progress]
//
// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest
type INSFontAssetRequest interface {
	objectivec.IObject

	// Topic: Creating a Font Asset Request

	InitWithFontDescriptorsOptions(fontDescriptors []NSFontDescriptor, options NSFontAssetRequestOptions) NSFontAssetRequest

	// Topic: Downloading a Font Asset

	DownloadFontAssetsWithCompletionHandler(completionHandler ErrorHandler)
	DownloadedFontDescriptors() []NSFontDescriptor

	// Topic: Getting the Download Progress

	Progress() foundation.NSProgress
}

// Init initializes the instance.
func (f NSFontAssetRequest) Init() NSFontAssetRequest {
	rv := objc.Send[NSFontAssetRequest](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFontAssetRequest) Autorelease() NSFontAssetRequest {
	rv := objc.Send[NSFontAssetRequest](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFontAssetRequest creates a new NSFontAssetRequest instance.
func NewNSFontAssetRequest() NSFontAssetRequest {
	class := getNSFontAssetRequestClass()
	rv := objc.Send[NSFontAssetRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/init(fontDescriptors:options:)
func NewFontAssetRequestWithFontDescriptorsOptions(fontDescriptors []NSFontDescriptor, options NSFontAssetRequestOptions) NSFontAssetRequest {
	instance := getNSFontAssetRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFontDescriptors:options:"), objectivec.IObjectSliceToNSArray(fontDescriptors), options)
	return NSFontAssetRequestFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/init(fontDescriptors:options:)
func (f NSFontAssetRequest) InitWithFontDescriptorsOptions(fontDescriptors []NSFontDescriptor, options NSFontAssetRequestOptions) NSFontAssetRequest {
	rv := objc.Send[NSFontAssetRequest](f.ID, objc.Sel("initWithFontDescriptors:options:"), objectivec.IObjectSliceToNSArray(fontDescriptors), options)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/download(withCompletionHandler:)
func (f NSFontAssetRequest) DownloadFontAssetsWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](f.ID, objc.Sel("downloadFontAssetsWithCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/downloadedFontDescriptors
func (f NSFontAssetRequest) DownloadedFontDescriptors() []NSFontDescriptor {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("downloadedFontDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSFontDescriptor {
		return NSFontDescriptorFromID(id)
	})
}

// See: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/progress
func (f NSFontAssetRequest) Progress() foundation.NSProgress {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("progress"))
	return foundation.NSProgressFromID(objc.ID(rv))
}

// DownloadFontAssets is a synchronous wrapper around [NSFontAssetRequest.DownloadFontAssetsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFontAssetRequest) DownloadFontAssets(ctx context.Context) error {
	done := make(chan error, 1)
	f.DownloadFontAssetsWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

