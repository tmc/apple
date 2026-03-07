// Code generated from Apple documentation. DO NOT EDIT.

package screencapturekit

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var SCStreamErrorDomain string

var sCStreamFrameInfoBoundingRect SCStreamFrameInfo

var sCStreamFrameInfoContentRect SCStreamFrameInfo

var sCStreamFrameInfoContentScale SCStreamFrameInfo

var sCStreamFrameInfoDirtyRects SCStreamFrameInfo

var sCStreamFrameInfoDisplayTime SCStreamFrameInfo

var sCStreamFrameInfoPresenterOverlayContentRect SCStreamFrameInfo

var sCStreamFrameInfoScaleFactor SCStreamFrameInfo

var sCStreamFrameInfoScreenRect SCStreamFrameInfo

var sCStreamFrameInfoStatus SCStreamFrameInfo

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				SCStreamErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoBoundingRect"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoBoundingRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoContentRect"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoContentRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoContentScale"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoContentScale = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoDirtyRects"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoDirtyRects = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoDisplayTime"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoDisplayTime = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoPresenterOverlayContentRect"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoPresenterOverlayContentRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoScaleFactor"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoScaleFactor = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoScreenRect"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoScreenRect = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "SCStreamFrameInfoStatus"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				sCStreamFrameInfoStatus = SCStreamFrameInfo(objc.GoString(cstr))
			}
		}
	}

}

type SCStreamFrameInfoValues struct{}

// SCStreamFrameInfos provides typed accessors for [SCStreamFrameInfo] constants.
var SCStreamFrameInfos SCStreamFrameInfoValues

// BoundingRect returns A key to retrieve the bounding rectangle for a video frame.
func (SCStreamFrameInfoValues) BoundingRect() SCStreamFrameInfo { return sCStreamFrameInfoBoundingRect }

// ContentRect returns A key to retrieve the content rectangle of a video frame.
func (SCStreamFrameInfoValues) ContentRect() SCStreamFrameInfo { return sCStreamFrameInfoContentRect }

// ContentScale returns A key to retrieve the content scale of a video frame.
func (SCStreamFrameInfoValues) ContentScale() SCStreamFrameInfo { return sCStreamFrameInfoContentScale }

// DirtyRects returns A key to retrieve the areas of a video frame that contain changes.
func (SCStreamFrameInfoValues) DirtyRects() SCStreamFrameInfo { return sCStreamFrameInfoDirtyRects }

// DisplayTime returns A key to retrieve the display time of a video frame.
func (SCStreamFrameInfoValues) DisplayTime() SCStreamFrameInfo { return sCStreamFrameInfoDisplayTime }

func (SCStreamFrameInfoValues) PresenterOverlayContentRect() SCStreamFrameInfo { return sCStreamFrameInfoPresenterOverlayContentRect }

// ScaleFactor returns A key to retrieve the scale factor of a video frame.
func (SCStreamFrameInfoValues) ScaleFactor() SCStreamFrameInfo { return sCStreamFrameInfoScaleFactor }

// ScreenRect returns A key to retrieve the onscreen location of captured content.
func (SCStreamFrameInfoValues) ScreenRect() SCStreamFrameInfo { return sCStreamFrameInfoScreenRect }

// Status returns A key to retrieve the status of a video frame.
func (SCStreamFrameInfoValues) Status() SCStreamFrameInfo { return sCStreamFrameInfoStatus }


