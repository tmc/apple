// Code generated from Apple documentation. DO NOT EDIT.

package screencapturekit

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// SCStreamErrorDomain is a string representation of the error domain.
	//
	// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamErrorDomain
	SCStreamErrorDomain string
)

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoBoundingRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.BoundingRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoContentRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.ContentRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoContentScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.ContentScale = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoDirtyRects"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.DirtyRects = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoDisplayTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.DisplayTime = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoPresenterOverlayContentRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.PresenterOverlayContentRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoScaleFactor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.ScaleFactor = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoScreenRect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.ScreenRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoStatus"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamFrameInfos.Status = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

}

// SCStreamFrameInfos provides typed accessors for [SCStreamFrameInfo] constants.
var SCStreamFrameInfos struct {
	// BoundingRect: A key to retrieve the bounding rectangle for a video frame.
	BoundingRect SCStreamFrameInfo
	// ContentRect: A key to retrieve the content rectangle of a video frame.
	ContentRect SCStreamFrameInfo
	// ContentScale: A key to retrieve the content scale of a video frame.
	ContentScale SCStreamFrameInfo
	// DirtyRects: A key to retrieve the areas of a video frame that contain changes.
	DirtyRects SCStreamFrameInfo
	// DisplayTime: A key to retrieve the display time of a video frame.
	DisplayTime                 SCStreamFrameInfo
	PresenterOverlayContentRect SCStreamFrameInfo
	// ScaleFactor: A key to retrieve the scale factor of a video frame.
	ScaleFactor SCStreamFrameInfo
	// ScreenRect: A key to retrieve the onscreen location of captured content.
	ScreenRect SCStreamFrameInfo
	// Status: A key to retrieve the status of a video frame.
	Status SCStreamFrameInfo
}
