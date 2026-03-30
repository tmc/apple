// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo
type CGBitmapInfo uint32

const (
	KCGBitmapAlphaInfoMask       CGBitmapInfo = 31
	KCGBitmapByteOrderInfoMask   CGBitmapInfo = 28672
	KCGBitmapComponentInfoMask   CGBitmapInfo = 3840
	KCGBitmapPixelFormatInfoMask CGBitmapInfo = 983040
	// Deprecated.
	KCGBitmapByteOrder16Big CGBitmapInfo = 12288
	// Deprecated.
	KCGBitmapByteOrder16Little CGBitmapInfo = 4096
	// Deprecated.
	KCGBitmapByteOrder32Big CGBitmapInfo = 12289
	// Deprecated.
	KCGBitmapByteOrder32Little CGBitmapInfo = 8192
	// Deprecated.
	KCGBitmapByteOrderDefault CGBitmapInfo = 0
	// Deprecated.
	KCGBitmapByteOrderMask CGBitmapInfo = 28672
	// Deprecated.
	KCGBitmapFloatComponents CGBitmapInfo = 256
	// Deprecated.
	KCGBitmapFloatInfoMask CGBitmapInfo = 3840
)

func (e CGBitmapInfo) String() string {
	switch e {
	case KCGBitmapAlphaInfoMask:
		return "KCGBitmapAlphaInfoMask"
	case KCGBitmapByteOrderInfoMask:
		return "KCGBitmapByteOrderInfoMask"
	case KCGBitmapComponentInfoMask:
		return "KCGBitmapComponentInfoMask"
	case KCGBitmapPixelFormatInfoMask:
		return "KCGBitmapPixelFormatInfoMask"
	case KCGBitmapByteOrder16Big:
		return "KCGBitmapByteOrder16Big"
	case KCGBitmapByteOrder16Little:
		return "KCGBitmapByteOrder16Little"
	case KCGBitmapByteOrder32Big:
		return "KCGBitmapByteOrder32Big"
	case KCGBitmapByteOrder32Little:
		return "KCGBitmapByteOrder32Little"
	case KCGBitmapByteOrderDefault:
		return "KCGBitmapByteOrderDefault"
	case KCGBitmapFloatComponents:
		return "KCGBitmapFloatComponents"
	default:
		return fmt.Sprintf("CGBitmapInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout
type CGBitmapLayout uint32

const (
	KCGBitmapLayoutABGR      CGBitmapLayout = 9
	KCGBitmapLayoutARGB      CGBitmapLayout = 4
	KCGBitmapLayoutAlphaOnly CGBitmapLayout = 0
	KCGBitmapLayoutBGRA      CGBitmapLayout = 7
	KCGBitmapLayoutBGRX      CGBitmapLayout = 8
	KCGBitmapLayoutCMYK      CGBitmapLayout = 11
	KCGBitmapLayoutGray      CGBitmapLayout = 1
	KCGBitmapLayoutGrayAlpha CGBitmapLayout = 2
	KCGBitmapLayoutRGBA      CGBitmapLayout = 3
	KCGBitmapLayoutRGBX      CGBitmapLayout = 5
	KCGBitmapLayoutXBGR      CGBitmapLayout = 10
	KCGBitmapLayoutXRGB      CGBitmapLayout = 6
)

func (e CGBitmapLayout) String() string {
	switch e {
	case KCGBitmapLayoutABGR:
		return "KCGBitmapLayoutABGR"
	case KCGBitmapLayoutARGB:
		return "KCGBitmapLayoutARGB"
	case KCGBitmapLayoutAlphaOnly:
		return "KCGBitmapLayoutAlphaOnly"
	case KCGBitmapLayoutBGRA:
		return "KCGBitmapLayoutBGRA"
	case KCGBitmapLayoutBGRX:
		return "KCGBitmapLayoutBGRX"
	case KCGBitmapLayoutCMYK:
		return "KCGBitmapLayoutCMYK"
	case KCGBitmapLayoutGray:
		return "KCGBitmapLayoutGray"
	case KCGBitmapLayoutGrayAlpha:
		return "KCGBitmapLayoutGrayAlpha"
	case KCGBitmapLayoutRGBA:
		return "KCGBitmapLayoutRGBA"
	case KCGBitmapLayoutRGBX:
		return "KCGBitmapLayoutRGBX"
	case KCGBitmapLayoutXBGR:
		return "KCGBitmapLayoutXBGR"
	case KCGBitmapLayoutXRGB:
		return "KCGBitmapLayoutXRGB"
	default:
		return fmt.Sprintf("CGBitmapLayout(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode
type CGBlendMode int32

const (
	// KCGBlendModeClear: `R = 0`
	KCGBlendModeClear CGBlendMode = 16
	// KCGBlendModeColor: Uses the luminance values of the background with the hue and saturation values of the source image.
	KCGBlendModeColor CGBlendMode = 14
	// KCGBlendModeColorBurn: Darkens the background image samples to reflect the source image samples.
	KCGBlendModeColorBurn CGBlendMode = 7
	// KCGBlendModeColorDodge: Brightens the background image samples to reflect the source image samples.
	KCGBlendModeColorDodge CGBlendMode = 6
	// KCGBlendModeCopy: `R = S`
	KCGBlendModeCopy CGBlendMode = 17
	// KCGBlendModeDarken: # Discussion
	KCGBlendModeDarken CGBlendMode = 4
	// KCGBlendModeDestinationAtop: `R = S*(1 - Da) + D*Sa`
	KCGBlendModeDestinationAtop CGBlendMode = 24
	// KCGBlendModeDestinationIn: `R = D*Sa`
	KCGBlendModeDestinationIn CGBlendMode = 22
	// KCGBlendModeDestinationOut: `R = D*(1 - Sa)`
	KCGBlendModeDestinationOut CGBlendMode = 23
	// KCGBlendModeDestinationOver: `R = S*(1 - Da) + D`
	KCGBlendModeDestinationOver CGBlendMode = 21
	// KCGBlendModeDifference: # Discussion
	KCGBlendModeDifference CGBlendMode = 10
	// KCGBlendModeExclusion: Produces an effect similar to that produced by CGBlendMode.difference, but with lower contrast.
	KCGBlendModeExclusion CGBlendMode = 11
	// KCGBlendModeHardLight: # Discussion
	KCGBlendModeHardLight CGBlendMode = 9
	// KCGBlendModeHue: Uses the luminance and saturation values of the background with the hue of the source image.
	KCGBlendModeHue CGBlendMode = 12
	// KCGBlendModeLighten: # Discussion
	KCGBlendModeLighten CGBlendMode = 5
	// KCGBlendModeLuminosity: Uses the hue and saturation of the background with the luminance of the source image.
	KCGBlendModeLuminosity CGBlendMode = 15
	// KCGBlendModeMultiply: Multiplies the source image samples with the background image samples.
	KCGBlendModeMultiply CGBlendMode = 1
	// KCGBlendModeNormal: Paints the source image samples over the background image samples.
	KCGBlendModeNormal CGBlendMode = 0
	// KCGBlendModeOverlay: # Discussion
	KCGBlendModeOverlay CGBlendMode = 3
	// KCGBlendModePlusDarker: `R = MAX(0, 1 - ((1 - D) + (1 - S)))`
	KCGBlendModePlusDarker CGBlendMode = 26
	// KCGBlendModePlusLighter: `R = MIN(1, S + D)`
	KCGBlendModePlusLighter CGBlendMode = 27
	// KCGBlendModeSaturation: Uses the luminance and hue values of the background with the saturation of the source image.
	KCGBlendModeSaturation CGBlendMode = 13
	// KCGBlendModeScreen: Multiplies the inverse of the source image samples with the inverse of the background image samples, resulting in colors that are at least as light as either of the two contributing sample colors.
	KCGBlendModeScreen CGBlendMode = 2
	// KCGBlendModeSoftLight: # Discussion
	KCGBlendModeSoftLight CGBlendMode = 8
	// KCGBlendModeSourceAtop: `R = S*Da + D*(1 - Sa)`
	KCGBlendModeSourceAtop CGBlendMode = 20
	// KCGBlendModeSourceIn: `R = S*Da`
	KCGBlendModeSourceIn CGBlendMode = 18
	// KCGBlendModeSourceOut: `R = S*(1 - Da)`
	KCGBlendModeSourceOut CGBlendMode = 19
	// KCGBlendModeXOR: `R = S*(1 - Da) + D*(1 - Sa)`.
	KCGBlendModeXOR CGBlendMode = 25
)

func (e CGBlendMode) String() string {
	switch e {
	case KCGBlendModeClear:
		return "KCGBlendModeClear"
	case KCGBlendModeColor:
		return "KCGBlendModeColor"
	case KCGBlendModeColorBurn:
		return "KCGBlendModeColorBurn"
	case KCGBlendModeColorDodge:
		return "KCGBlendModeColorDodge"
	case KCGBlendModeCopy:
		return "KCGBlendModeCopy"
	case KCGBlendModeDarken:
		return "KCGBlendModeDarken"
	case KCGBlendModeDestinationAtop:
		return "KCGBlendModeDestinationAtop"
	case KCGBlendModeDestinationIn:
		return "KCGBlendModeDestinationIn"
	case KCGBlendModeDestinationOut:
		return "KCGBlendModeDestinationOut"
	case KCGBlendModeDestinationOver:
		return "KCGBlendModeDestinationOver"
	case KCGBlendModeDifference:
		return "KCGBlendModeDifference"
	case KCGBlendModeExclusion:
		return "KCGBlendModeExclusion"
	case KCGBlendModeHardLight:
		return "KCGBlendModeHardLight"
	case KCGBlendModeHue:
		return "KCGBlendModeHue"
	case KCGBlendModeLighten:
		return "KCGBlendModeLighten"
	case KCGBlendModeLuminosity:
		return "KCGBlendModeLuminosity"
	case KCGBlendModeMultiply:
		return "KCGBlendModeMultiply"
	case KCGBlendModeNormal:
		return "KCGBlendModeNormal"
	case KCGBlendModeOverlay:
		return "KCGBlendModeOverlay"
	case KCGBlendModePlusDarker:
		return "KCGBlendModePlusDarker"
	case KCGBlendModePlusLighter:
		return "KCGBlendModePlusLighter"
	case KCGBlendModeSaturation:
		return "KCGBlendModeSaturation"
	case KCGBlendModeScreen:
		return "KCGBlendModeScreen"
	case KCGBlendModeSoftLight:
		return "KCGBlendModeSoftLight"
	case KCGBlendModeSourceAtop:
		return "KCGBlendModeSourceAtop"
	case KCGBlendModeSourceIn:
		return "KCGBlendModeSourceIn"
	case KCGBlendModeSourceOut:
		return "KCGBlendModeSourceOut"
	case KCGBlendModeXOR:
		return "KCGBlendModeXOR"
	default:
		return fmt.Sprintf("CGBlendMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions
type CGCaptureOptions uint32

const (
	// KCGCaptureNoFill: Disables fill with black.
	KCGCaptureNoFill CGCaptureOptions = 1
	// KCGCaptureNoOptions: The system should use the default fill behavior, which is fill with black.
	KCGCaptureNoOptions CGCaptureOptions = 0
	// KCGMouseEventSubtype: Key to access an integer field that encodes the mouse event subtype as a .
	KCGMouseEventSubtype CGCaptureOptions = 0
)

func (e CGCaptureOptions) String() string {
	switch e {
	case KCGCaptureNoFill:
		return "KCGCaptureNoFill"
	case KCGCaptureNoOptions:
		return "KCGCaptureNoOptions"
	default:
		return fmt.Sprintf("CGCaptureOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType
type CGColorConversionInfoTransformType uint32

const (
	// KCGColorConversionTransformApplySpace: Specifies a color conversion between one color profile and another.
	KCGColorConversionTransformApplySpace CGColorConversionInfoTransformType = 2
	// KCGColorConversionTransformFromSpace: Specifies a color conversion from a device color space to a color profile.
	KCGColorConversionTransformFromSpace CGColorConversionInfoTransformType = 0
	// KCGColorConversionTransformToSpace: Specifies a color conversion from a color profile to a device color space.
	KCGColorConversionTransformToSpace CGColorConversionInfoTransformType = 1
)

func (e CGColorConversionInfoTransformType) String() string {
	switch e {
	case KCGColorConversionTransformApplySpace:
		return "KCGColorConversionTransformApplySpace"
	case KCGColorConversionTransformFromSpace:
		return "KCGColorConversionTransformFromSpace"
	case KCGColorConversionTransformToSpace:
		return "KCGColorConversionTransformToSpace"
	default:
		return fmt.Sprintf("CGColorConversionInfoTransformType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGColorModel
type CGColorModel uint32

const (
	KCGColorModelCMYK       CGColorModel = 4
	KCGColorModelDeviceN    CGColorModel = 16
	KCGColorModelGray       CGColorModel = 1
	KCGColorModelLab        CGColorModel = 8
	KCGColorModelNoColorant CGColorModel = 0
	KCGColorModelRGB        CGColorModel = 2
)

func (e CGColorModel) String() string {
	switch e {
	case KCGColorModelCMYK:
		return "KCGColorModelCMYK"
	case KCGColorModelDeviceN:
		return "KCGColorModelDeviceN"
	case KCGColorModelGray:
		return "KCGColorModelGray"
	case KCGColorModelLab:
		return "KCGColorModelLab"
	case KCGColorModelNoColorant:
		return "KCGColorModelNoColorant"
	case KCGColorModelRGB:
		return "KCGColorModelRGB"
	default:
		return fmt.Sprintf("CGColorModel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent
type CGColorRenderingIntent int32

const (
	// KCGRenderingIntentAbsoluteColorimetric: # Discussion
	KCGRenderingIntentAbsoluteColorimetric CGColorRenderingIntent = 1
	// KCGRenderingIntentDefault: The default rendering intent for the graphics context.
	KCGRenderingIntentDefault CGColorRenderingIntent = 0
	// KCGRenderingIntentPerceptual: Preserve the visual relationship between colors by compressing the gamut of the graphics context to fit inside the gamut of the output device.
	KCGRenderingIntentPerceptual CGColorRenderingIntent = 3
	// KCGRenderingIntentRelativeColorimetric: # Discussion
	KCGRenderingIntentRelativeColorimetric CGColorRenderingIntent = 2
	// KCGRenderingIntentSaturation: # Discussion
	KCGRenderingIntentSaturation CGColorRenderingIntent = 4
)

func (e CGColorRenderingIntent) String() string {
	switch e {
	case KCGRenderingIntentAbsoluteColorimetric:
		return "KCGRenderingIntentAbsoluteColorimetric"
	case KCGRenderingIntentDefault:
		return "KCGRenderingIntentDefault"
	case KCGRenderingIntentPerceptual:
		return "KCGRenderingIntentPerceptual"
	case KCGRenderingIntentRelativeColorimetric:
		return "KCGRenderingIntentRelativeColorimetric"
	case KCGRenderingIntentSaturation:
		return "KCGRenderingIntentSaturation"
	default:
		return fmt.Sprintf("CGColorRenderingIntent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel
type CGColorSpaceModel int32

const (
	// KCGColorSpaceModelCMYK: A CMYK color space model.
	KCGColorSpaceModelCMYK CGColorSpaceModel = 2
	// KCGColorSpaceModelDeviceN: A DeviceN color space model.
	KCGColorSpaceModelDeviceN CGColorSpaceModel = 4
	// KCGColorSpaceModelIndexed: An indexed color space model.
	KCGColorSpaceModelIndexed CGColorSpaceModel = 5
	// KCGColorSpaceModelLab: A Lab color space model.
	KCGColorSpaceModelLab CGColorSpaceModel = 3
	// KCGColorSpaceModelMonochrome: A monochrome color space model.
	KCGColorSpaceModelMonochrome CGColorSpaceModel = 0
	// KCGColorSpaceModelPattern: A pattern color space model.
	KCGColorSpaceModelPattern CGColorSpaceModel = 6
	// KCGColorSpaceModelRGB: An RGB color space model.
	KCGColorSpaceModelRGB CGColorSpaceModel = 1
	// KCGColorSpaceModelUnknown: An unknown color space model.
	KCGColorSpaceModelUnknown CGColorSpaceModel = -1
	// KCGColorSpaceModelXYZ: An XYZ color space model.
	KCGColorSpaceModelXYZ CGColorSpaceModel = 7
)

func (e CGColorSpaceModel) String() string {
	switch e {
	case KCGColorSpaceModelCMYK:
		return "KCGColorSpaceModelCMYK"
	case KCGColorSpaceModelDeviceN:
		return "KCGColorSpaceModelDeviceN"
	case KCGColorSpaceModelIndexed:
		return "KCGColorSpaceModelIndexed"
	case KCGColorSpaceModelLab:
		return "KCGColorSpaceModelLab"
	case KCGColorSpaceModelMonochrome:
		return "KCGColorSpaceModelMonochrome"
	case KCGColorSpaceModelPattern:
		return "KCGColorSpaceModelPattern"
	case KCGColorSpaceModelRGB:
		return "KCGColorSpaceModelRGB"
	case KCGColorSpaceModelUnknown:
		return "KCGColorSpaceModelUnknown"
	case KCGColorSpaceModelXYZ:
		return "KCGColorSpaceModelXYZ"
	default:
		return fmt.Sprintf("CGColorSpaceModel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGComponent
type CGComponent uint32

const (
	KCGComponentFloat16Bit   CGComponent = 5
	KCGComponentFloat32Bit   CGComponent = 4
	KCGComponentInteger10Bit CGComponent = 6
	KCGComponentInteger16Bit CGComponent = 2
	KCGComponentInteger32Bit CGComponent = 3
	KCGComponentInteger8Bit  CGComponent = 1
	KCGComponentUnknown      CGComponent = 0
)

func (e CGComponent) String() string {
	switch e {
	case KCGComponentFloat16Bit:
		return "KCGComponentFloat16Bit"
	case KCGComponentFloat32Bit:
		return "KCGComponentFloat32Bit"
	case KCGComponentInteger10Bit:
		return "KCGComponentInteger10Bit"
	case KCGComponentInteger16Bit:
		return "KCGComponentInteger16Bit"
	case KCGComponentInteger32Bit:
		return "KCGComponentInteger32Bit"
	case KCGComponentInteger8Bit:
		return "KCGComponentInteger8Bit"
	case KCGComponentUnknown:
		return "KCGComponentUnknown"
	default:
		return fmt.Sprintf("CGComponent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption
type CGConfigureOption uint32

const (
	// KCGConfigureForAppOnly: Changes persist for the lifetime of the current application.
	KCGConfigureForAppOnly CGConfigureOption = 0
	// KCGConfigureForSession: Changes persist for the lifetime of the current login session.
	KCGConfigureForSession CGConfigureOption = 1
	// KCGConfigurePermanently: # Discussion
	KCGConfigurePermanently CGConfigureOption = 2
)

func (e CGConfigureOption) String() string {
	switch e {
	case KCGConfigureForAppOnly:
		return "KCGConfigureForAppOnly"
	case KCGConfigureForSession:
		return "KCGConfigureForSession"
	case KCGConfigurePermanently:
		return "KCGConfigurePermanently"
	default:
		return fmt.Sprintf("CGConfigureOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags
type CGDisplayChangeSummaryFlags uint32

const (
	// KCGDisplayAddFlag: The display has been added to the active display list.
	KCGDisplayAddFlag CGDisplayChangeSummaryFlags = 16
	// KCGDisplayBeginConfigurationFlag: The display configuration is about to change.
	KCGDisplayBeginConfigurationFlag CGDisplayChangeSummaryFlags = 1
	// KCGDisplayDesktopShapeChangedFlag: # Discussion
	KCGDisplayDesktopShapeChangedFlag CGDisplayChangeSummaryFlags = 4096
	// KCGDisplayDisabledFlag: The display has been disabled.
	KCGDisplayDisabledFlag CGDisplayChangeSummaryFlags = 512
	// KCGDisplayEnabledFlag: The display has been enabled.
	KCGDisplayEnabledFlag CGDisplayChangeSummaryFlags = 256
	// KCGDisplayMirrorFlag: The display is now mirroring another display.
	KCGDisplayMirrorFlag CGDisplayChangeSummaryFlags = 1024
	// KCGDisplayMovedFlag: The location of the upper-left corner of the display in the global display coordinate space has changed.
	KCGDisplayMovedFlag CGDisplayChangeSummaryFlags = 2
	// KCGDisplayRemoveFlag: The display has been removed from the active display list.
	KCGDisplayRemoveFlag CGDisplayChangeSummaryFlags = 32
	// KCGDisplaySetMainFlag: The display is now the main display.
	KCGDisplaySetMainFlag CGDisplayChangeSummaryFlags = 4
	// KCGDisplaySetModeFlag: The display mode has changed.
	KCGDisplaySetModeFlag CGDisplayChangeSummaryFlags = 8
	// KCGDisplayUnMirrorFlag: The display is no longer mirroring another display.
	KCGDisplayUnMirrorFlag CGDisplayChangeSummaryFlags = 2048
)

func (e CGDisplayChangeSummaryFlags) String() string {
	switch e {
	case KCGDisplayAddFlag:
		return "KCGDisplayAddFlag"
	case KCGDisplayBeginConfigurationFlag:
		return "KCGDisplayBeginConfigurationFlag"
	case KCGDisplayDesktopShapeChangedFlag:
		return "KCGDisplayDesktopShapeChangedFlag"
	case KCGDisplayDisabledFlag:
		return "KCGDisplayDisabledFlag"
	case KCGDisplayEnabledFlag:
		return "KCGDisplayEnabledFlag"
	case KCGDisplayMirrorFlag:
		return "KCGDisplayMirrorFlag"
	case KCGDisplayMovedFlag:
		return "KCGDisplayMovedFlag"
	case KCGDisplayRemoveFlag:
		return "KCGDisplayRemoveFlag"
	case KCGDisplaySetMainFlag:
		return "KCGDisplaySetMainFlag"
	case KCGDisplaySetModeFlag:
		return "KCGDisplaySetModeFlag"
	case KCGDisplayUnMirrorFlag:
		return "KCGDisplayUnMirrorFlag"
	default:
		return fmt.Sprintf("CGDisplayChangeSummaryFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus
type CGDisplayStreamFrameStatus int32

const (
	// KCGDisplayStreamFrameStatusFrameBlank: A new frame was not generated because the display has gone blank.
	KCGDisplayStreamFrameStatusFrameBlank CGDisplayStreamFrameStatus = 2
	// KCGDisplayStreamFrameStatusFrameComplete: A new frame was generated.
	KCGDisplayStreamFrameStatusFrameComplete CGDisplayStreamFrameStatus = 0
	// KCGDisplayStreamFrameStatusFrameIdle: A new frame was not generated because the display did not change.
	KCGDisplayStreamFrameStatusFrameIdle CGDisplayStreamFrameStatus = 1
	// KCGDisplayStreamFrameStatusStopped: The display stream was stopped.
	KCGDisplayStreamFrameStatusStopped CGDisplayStreamFrameStatus = 3
)

func (e CGDisplayStreamFrameStatus) String() string {
	switch e {
	case KCGDisplayStreamFrameStatusFrameBlank:
		return "KCGDisplayStreamFrameStatusFrameBlank"
	case KCGDisplayStreamFrameStatusFrameComplete:
		return "KCGDisplayStreamFrameStatusFrameComplete"
	case KCGDisplayStreamFrameStatusFrameIdle:
		return "KCGDisplayStreamFrameStatusFrameIdle"
	case KCGDisplayStreamFrameStatusStopped:
		return "KCGDisplayStreamFrameStatusStopped"
	default:
		return fmt.Sprintf("CGDisplayStreamFrameStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType
type CGDisplayStreamUpdateRectType int32

const (
	// KCGDisplayStreamUpdateDirtyRects: The union of both rectangles that were redrawn and rectangles that were moved.
	KCGDisplayStreamUpdateDirtyRects CGDisplayStreamUpdateRectType = 2
	// KCGDisplayStreamUpdateMovedRects: The rectangles for the portions of the display that were simply moved from one part of the display to another.
	KCGDisplayStreamUpdateMovedRects CGDisplayStreamUpdateRectType = 1
	// KCGDisplayStreamUpdateReducedDirtyRects: The union is calculated and then simplified.
	KCGDisplayStreamUpdateReducedDirtyRects CGDisplayStreamUpdateRectType = 3
	// KCGDisplayStreamUpdateRefreshedRects: The rectangles for the portions of the display that were redrawn.
	KCGDisplayStreamUpdateRefreshedRects CGDisplayStreamUpdateRectType = 0
)

func (e CGDisplayStreamUpdateRectType) String() string {
	switch e {
	case KCGDisplayStreamUpdateDirtyRects:
		return "KCGDisplayStreamUpdateDirtyRects"
	case KCGDisplayStreamUpdateMovedRects:
		return "KCGDisplayStreamUpdateMovedRects"
	case KCGDisplayStreamUpdateReducedDirtyRects:
		return "KCGDisplayStreamUpdateReducedDirtyRects"
	case KCGDisplayStreamUpdateRefreshedRects:
		return "KCGDisplayStreamUpdateRefreshedRects"
	default:
		return fmt.Sprintf("CGDisplayStreamUpdateRectType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGError
type CGError int

const (
	// KCGErrorCannotComplete: The requested operation is inappropriate for the parameters passed in, or the current system state.
	KCGErrorCannotComplete CGError = 1004
	// KCGErrorFailure: A general failure occurred.
	KCGErrorFailure CGError = 1000
	// KCGErrorIllegalArgument: One or more of the parameters passed to a function are invalid.
	KCGErrorIllegalArgument CGError = 1001
	// KCGErrorInvalidConnection: The parameter representing a connection to the window server is invalid.
	KCGErrorInvalidConnection CGError = 1002
	// KCGErrorInvalidContext: The [CPSProcessSerNum] or context identifier parameter is not valid.
	KCGErrorInvalidContext CGError = 1003
	// KCGErrorInvalidOperation: The requested operation is not valid for the parameters passed in, or the current system state.
	KCGErrorInvalidOperation CGError = 1010
	// KCGErrorNoneAvailable: The requested operation could not be completed as the indicated resources were not found.
	KCGErrorNoneAvailable CGError = 1011
	// KCGErrorNotImplemented: Return value from obsolete function stubs present for binary compatibility, but not typically called.
	KCGErrorNotImplemented CGError = 1006
	// KCGErrorRangeCheck: A parameter passed in has a value that is inappropriate, or which does not map to a useful operation or value.
	KCGErrorRangeCheck CGError = 1007
	// KCGErrorSuccess: The requested operation was completed successfully.
	KCGErrorSuccess CGError = 0
	// KCGErrorTypeCheck: A data type or token was encountered that did not match the expected type or token.
	KCGErrorTypeCheck CGError = 1008
)

func (e CGError) String() string {
	switch e {
	case KCGErrorCannotComplete:
		return "KCGErrorCannotComplete"
	case KCGErrorFailure:
		return "KCGErrorFailure"
	case KCGErrorIllegalArgument:
		return "KCGErrorIllegalArgument"
	case KCGErrorInvalidConnection:
		return "KCGErrorInvalidConnection"
	case KCGErrorInvalidContext:
		return "KCGErrorInvalidContext"
	case KCGErrorInvalidOperation:
		return "KCGErrorInvalidOperation"
	case KCGErrorNoneAvailable:
		return "KCGErrorNoneAvailable"
	case KCGErrorNotImplemented:
		return "KCGErrorNotImplemented"
	case KCGErrorRangeCheck:
		return "KCGErrorRangeCheck"
	case KCGErrorSuccess:
		return "KCGErrorSuccess"
	case KCGErrorTypeCheck:
		return "KCGErrorTypeCheck"
	default:
		return fmt.Sprintf("CGError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventField
type CGEventField uint32

const (
	// KCGEventSourceGroupID: Key to access a field that contains the event source Unix effective GID.
	KCGEventSourceGroupID CGEventField = 44
	// KCGEventSourceStateID: Key to access a field that contains the event source state ID used to create this event.
	KCGEventSourceStateID CGEventField = 45
	// KCGEventSourceUnixProcessID: Key to access a field that contains the event source Unix process ID.
	KCGEventSourceUnixProcessID CGEventField = 41
	// KCGEventSourceUserData: Key to access a field that contains the event source user-supplied data, up to 64 bits.
	KCGEventSourceUserData CGEventField = 42
	// KCGEventSourceUserID: Key to access a field that contains the event source Unix effective UID.
	KCGEventSourceUserID CGEventField = 43
	// KCGEventTargetProcessSerialNumber: Key to access a field that contains the event target process serial number.
	KCGEventTargetProcessSerialNumber CGEventField = 39
	// KCGEventTargetUnixProcessID: Key to access a field that contains the event target Unix process ID.
	KCGEventTargetUnixProcessID           CGEventField = 40
	KCGEventUnacceleratedPointerMovementX CGEventField = 170
	KCGEventUnacceleratedPointerMovementY CGEventField = 171
	// KCGKeyboardEventAutorepeat: Key to access an integer field, non-zero when this is an autorepeat of a key-down, and zero otherwise.
	KCGKeyboardEventAutorepeat CGEventField = 8
	// KCGKeyboardEventKeyboardType: Key to access an integer field that contains the keyboard type identifier.
	KCGKeyboardEventKeyboardType CGEventField = 10
	// KCGKeyboardEventKeycode: Key to access an integer field that contains the virtual keycode of the key-down or key-up event.
	KCGKeyboardEventKeycode CGEventField = 9
	// KCGMouseEventButtonNumber: Key to access an integer field that contains the mouse button number.
	KCGMouseEventButtonNumber CGEventField = 3
	// KCGMouseEventClickState: Key to access an integer field that contains the mouse button click state.
	KCGMouseEventClickState CGEventField = 1
	// KCGMouseEventDeltaX: Key to access an integer field that contains the horizontal mouse delta since the last mouse movement event.
	KCGMouseEventDeltaX CGEventField = 4
	// KCGMouseEventDeltaY: Key to access an integer field that contains the vertical mouse delta since the last mouse movement event.
	KCGMouseEventDeltaY CGEventField = 5
	// KCGMouseEventInstantMouser: Key to access an integer field.
	KCGMouseEventInstantMouser CGEventField = 6
	// KCGMouseEventNumber: Key to access an integer field that contains the mouse button event number.
	KCGMouseEventNumber CGEventField = 0
	// KCGMouseEventPressure: Key to access a double field that contains the mouse button pressure.
	KCGMouseEventPressure                                      CGEventField = 2
	KCGMouseEventWindowUnderMousePointer                       CGEventField = 91
	KCGMouseEventWindowUnderMousePointerThatCanHandleThisEvent CGEventField = 92
	KCGScrollWheelEventAcceleratedDeltaAxis1                   CGEventField = 176
	KCGScrollWheelEventAcceleratedDeltaAxis2                   CGEventField = 175
	// KCGScrollWheelEventDeltaAxis1: Key to access an integer field that contains scrolling data.
	KCGScrollWheelEventDeltaAxis1 CGEventField = 11
	// KCGScrollWheelEventDeltaAxis2: Key to access an integer field that contains scrolling data.
	KCGScrollWheelEventDeltaAxis2 CGEventField = 12
	// KCGScrollWheelEventDeltaAxis3: This field is not used.
	KCGScrollWheelEventDeltaAxis3 CGEventField = 13
	// KCGScrollWheelEventFixedPtDeltaAxis1: Key to access a field that contains scrolling data.
	KCGScrollWheelEventFixedPtDeltaAxis1 CGEventField = 93
	// KCGScrollWheelEventFixedPtDeltaAxis2: Key to access a field that contains scrolling data.
	KCGScrollWheelEventFixedPtDeltaAxis2 CGEventField = 94
	// KCGScrollWheelEventFixedPtDeltaAxis3: This field is not used.
	KCGScrollWheelEventFixedPtDeltaAxis3 CGEventField = 95
	// KCGScrollWheelEventInstantMouser: Key to access an integer field that indicates whether the event should be ignored by the Inkwell subsystem.
	KCGScrollWheelEventInstantMouser CGEventField = 14
	// KCGScrollWheelEventIsContinuous: Key to access an integer field that indicates whether a scrolling event contains continuous, pixel-based scrolling data.
	KCGScrollWheelEventIsContinuous        CGEventField = 88
	KCGScrollWheelEventMomentumOptionPhase CGEventField = 173
	KCGScrollWheelEventMomentumPhase       CGEventField = 123
	// KCGScrollWheelEventPointDeltaAxis1: Key to access an integer field that contains pixel-based scrolling data.
	KCGScrollWheelEventPointDeltaAxis1 CGEventField = 96
	// KCGScrollWheelEventPointDeltaAxis2: Key to access an integer field that contains pixel-based scrolling data.
	KCGScrollWheelEventPointDeltaAxis2 CGEventField = 97
	// KCGScrollWheelEventPointDeltaAxis3: This field is not used.
	KCGScrollWheelEventPointDeltaAxis3 CGEventField = 98
	KCGScrollWheelEventRawDeltaAxis1   CGEventField = 178
	KCGScrollWheelEventRawDeltaAxis2   CGEventField = 177
	KCGScrollWheelEventScrollCount     CGEventField = 100
	KCGScrollWheelEventScrollPhase     CGEventField = 99
	// KCGTabletEventDeviceID: Key to access an integer field that contains the system-assigned unique device ID.
	KCGTabletEventDeviceID CGEventField = 24
	// KCGTabletEventPointButtons: Key to access an integer field that contains the tablet button state.
	KCGTabletEventPointButtons CGEventField = 18
	// KCGTabletEventPointPressure: Key to access a double field that contains the tablet pen pressure.
	KCGTabletEventPointPressure CGEventField = 19
	// KCGTabletEventPointX: Key to access an integer field that contains the absolute X coordinate in tablet space at full tablet resolution.
	KCGTabletEventPointX CGEventField = 15
	// KCGTabletEventPointY: Key to access an integer field that contains the absolute Y coordinate in tablet space at full tablet resolution.
	KCGTabletEventPointY CGEventField = 16
	// KCGTabletEventPointZ: Key to access an integer field that contains the absolute Z coordinate in tablet space at full tablet resolution.
	KCGTabletEventPointZ CGEventField = 17
	// KCGTabletEventRotation: Key to access a double field that contains the tablet pen rotation.
	KCGTabletEventRotation CGEventField = 22
	// KCGTabletEventTangentialPressure: Key to access a double field that contains the tangential pressure on the device.
	KCGTabletEventTangentialPressure CGEventField = 23
	// KCGTabletEventTiltX: Key to access a double field that contains the horizontal tablet pen tilt.
	KCGTabletEventTiltX CGEventField = 20
	// KCGTabletEventTiltY: Key to access a double field that contains the vertical tablet pen tilt.
	KCGTabletEventTiltY CGEventField = 21
	// KCGTabletEventVendor1: Key to access an integer field that contains a vendor-specified value.
	KCGTabletEventVendor1 CGEventField = 25
	// KCGTabletEventVendor2: Key to access an integer field that contains a vendor-specified value.
	KCGTabletEventVendor2 CGEventField = 26
	// KCGTabletEventVendor3: Key to access an integer field that contains a vendor-specified value.
	KCGTabletEventVendor3 CGEventField = 27
	// KCGTabletProximityEventCapabilityMask: Key to access an integer field that contains the device capabilities mask.
	KCGTabletProximityEventCapabilityMask CGEventField = 36
	// KCGTabletProximityEventDeviceID: Key to access an integer field that contains the system-assigned device ID.
	KCGTabletProximityEventDeviceID CGEventField = 31
	// KCGTabletProximityEventEnterProximity: Key to access an integer field that indicates whether the pen is in proximity to the tablet.
	KCGTabletProximityEventEnterProximity CGEventField = 38
	// KCGTabletProximityEventPointerID: Key to access an integer field that contains the vendor-defined ID of the pointing device.
	KCGTabletProximityEventPointerID CGEventField = 30
	// KCGTabletProximityEventPointerType: Key to access an integer field that contains the pointer type.
	KCGTabletProximityEventPointerType CGEventField = 37
	// KCGTabletProximityEventSystemTabletID: Key to access an integer field that contains the system-assigned unique tablet ID.
	KCGTabletProximityEventSystemTabletID CGEventField = 32
	// KCGTabletProximityEventTabletID: Key to access an integer field that contains the vendor-defined tablet ID, typically the USB product ID.
	KCGTabletProximityEventTabletID CGEventField = 29
	// KCGTabletProximityEventVendorID: Key to access an integer field that contains the vendor-defined ID, typically the USB vendor ID.
	KCGTabletProximityEventVendorID CGEventField = 28
	// KCGTabletProximityEventVendorPointerSerialNumber: Key to access an integer field that contains the vendor-defined pointer serial number.
	KCGTabletProximityEventVendorPointerSerialNumber CGEventField = 34
	// KCGTabletProximityEventVendorPointerType: Key to access an integer field that contains the vendor-assigned pointer type.
	KCGTabletProximityEventVendorPointerType CGEventField = 33
	// KCGTabletProximityEventVendorUniqueID: Key to access an integer field that contains the vendor-defined unique ID.
	KCGTabletProximityEventVendorUniqueID CGEventField = 35
)

func (e CGEventField) String() string {
	switch e {
	case KCGEventSourceGroupID:
		return "KCGEventSourceGroupID"
	case KCGEventSourceStateID:
		return "KCGEventSourceStateID"
	case KCGEventSourceUnixProcessID:
		return "KCGEventSourceUnixProcessID"
	case KCGEventSourceUserData:
		return "KCGEventSourceUserData"
	case KCGEventSourceUserID:
		return "KCGEventSourceUserID"
	case KCGEventTargetProcessSerialNumber:
		return "KCGEventTargetProcessSerialNumber"
	case KCGEventTargetUnixProcessID:
		return "KCGEventTargetUnixProcessID"
	case KCGEventUnacceleratedPointerMovementX:
		return "KCGEventUnacceleratedPointerMovementX"
	case KCGEventUnacceleratedPointerMovementY:
		return "KCGEventUnacceleratedPointerMovementY"
	case KCGKeyboardEventAutorepeat:
		return "KCGKeyboardEventAutorepeat"
	case KCGKeyboardEventKeyboardType:
		return "KCGKeyboardEventKeyboardType"
	case KCGKeyboardEventKeycode:
		return "KCGKeyboardEventKeycode"
	case KCGMouseEventButtonNumber:
		return "KCGMouseEventButtonNumber"
	case KCGMouseEventClickState:
		return "KCGMouseEventClickState"
	case KCGMouseEventDeltaX:
		return "KCGMouseEventDeltaX"
	case KCGMouseEventDeltaY:
		return "KCGMouseEventDeltaY"
	case KCGMouseEventInstantMouser:
		return "KCGMouseEventInstantMouser"
	case KCGMouseEventNumber:
		return "KCGMouseEventNumber"
	case KCGMouseEventPressure:
		return "KCGMouseEventPressure"
	case KCGMouseEventWindowUnderMousePointer:
		return "KCGMouseEventWindowUnderMousePointer"
	case KCGMouseEventWindowUnderMousePointerThatCanHandleThisEvent:
		return "KCGMouseEventWindowUnderMousePointerThatCanHandleThisEvent"
	case KCGScrollWheelEventAcceleratedDeltaAxis1:
		return "KCGScrollWheelEventAcceleratedDeltaAxis1"
	case KCGScrollWheelEventAcceleratedDeltaAxis2:
		return "KCGScrollWheelEventAcceleratedDeltaAxis2"
	case KCGScrollWheelEventDeltaAxis1:
		return "KCGScrollWheelEventDeltaAxis1"
	case KCGScrollWheelEventDeltaAxis2:
		return "KCGScrollWheelEventDeltaAxis2"
	case KCGScrollWheelEventDeltaAxis3:
		return "KCGScrollWheelEventDeltaAxis3"
	case KCGScrollWheelEventFixedPtDeltaAxis1:
		return "KCGScrollWheelEventFixedPtDeltaAxis1"
	case KCGScrollWheelEventFixedPtDeltaAxis2:
		return "KCGScrollWheelEventFixedPtDeltaAxis2"
	case KCGScrollWheelEventFixedPtDeltaAxis3:
		return "KCGScrollWheelEventFixedPtDeltaAxis3"
	case KCGScrollWheelEventInstantMouser:
		return "KCGScrollWheelEventInstantMouser"
	case KCGScrollWheelEventIsContinuous:
		return "KCGScrollWheelEventIsContinuous"
	case KCGScrollWheelEventMomentumOptionPhase:
		return "KCGScrollWheelEventMomentumOptionPhase"
	case KCGScrollWheelEventMomentumPhase:
		return "KCGScrollWheelEventMomentumPhase"
	case KCGScrollWheelEventPointDeltaAxis1:
		return "KCGScrollWheelEventPointDeltaAxis1"
	case KCGScrollWheelEventPointDeltaAxis2:
		return "KCGScrollWheelEventPointDeltaAxis2"
	case KCGScrollWheelEventPointDeltaAxis3:
		return "KCGScrollWheelEventPointDeltaAxis3"
	case KCGScrollWheelEventRawDeltaAxis1:
		return "KCGScrollWheelEventRawDeltaAxis1"
	case KCGScrollWheelEventRawDeltaAxis2:
		return "KCGScrollWheelEventRawDeltaAxis2"
	case KCGScrollWheelEventScrollCount:
		return "KCGScrollWheelEventScrollCount"
	case KCGScrollWheelEventScrollPhase:
		return "KCGScrollWheelEventScrollPhase"
	case KCGTabletEventDeviceID:
		return "KCGTabletEventDeviceID"
	case KCGTabletEventPointButtons:
		return "KCGTabletEventPointButtons"
	case KCGTabletEventPointPressure:
		return "KCGTabletEventPointPressure"
	case KCGTabletEventPointX:
		return "KCGTabletEventPointX"
	case KCGTabletEventPointY:
		return "KCGTabletEventPointY"
	case KCGTabletEventPointZ:
		return "KCGTabletEventPointZ"
	case KCGTabletEventRotation:
		return "KCGTabletEventRotation"
	case KCGTabletEventTangentialPressure:
		return "KCGTabletEventTangentialPressure"
	case KCGTabletEventTiltX:
		return "KCGTabletEventTiltX"
	case KCGTabletEventTiltY:
		return "KCGTabletEventTiltY"
	case KCGTabletEventVendor1:
		return "KCGTabletEventVendor1"
	case KCGTabletEventVendor2:
		return "KCGTabletEventVendor2"
	case KCGTabletEventVendor3:
		return "KCGTabletEventVendor3"
	case KCGTabletProximityEventCapabilityMask:
		return "KCGTabletProximityEventCapabilityMask"
	case KCGTabletProximityEventDeviceID:
		return "KCGTabletProximityEventDeviceID"
	case KCGTabletProximityEventEnterProximity:
		return "KCGTabletProximityEventEnterProximity"
	case KCGTabletProximityEventPointerID:
		return "KCGTabletProximityEventPointerID"
	case KCGTabletProximityEventPointerType:
		return "KCGTabletProximityEventPointerType"
	case KCGTabletProximityEventSystemTabletID:
		return "KCGTabletProximityEventSystemTabletID"
	case KCGTabletProximityEventTabletID:
		return "KCGTabletProximityEventTabletID"
	case KCGTabletProximityEventVendorID:
		return "KCGTabletProximityEventVendorID"
	case KCGTabletProximityEventVendorPointerSerialNumber:
		return "KCGTabletProximityEventVendorPointerSerialNumber"
	case KCGTabletProximityEventVendorPointerType:
		return "KCGTabletProximityEventVendorPointerType"
	case KCGTabletProximityEventVendorUniqueID:
		return "KCGTabletProximityEventVendorUniqueID"
	default:
		return fmt.Sprintf("CGEventField(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask
type CGEventFilterMask uint32

const (
	KCGEventFilterMaskPermitLocalKeyboardEvents CGEventFilterMask = 2
	KCGEventFilterMaskPermitLocalMouseEvents    CGEventFilterMask = 1
	KCGEventFilterMaskPermitSystemDefinedEvents CGEventFilterMask = 4
)

func (e CGEventFilterMask) String() string {
	switch e {
	case KCGEventFilterMaskPermitLocalKeyboardEvents:
		return "KCGEventFilterMaskPermitLocalKeyboardEvents"
	case KCGEventFilterMaskPermitLocalMouseEvents:
		return "KCGEventFilterMaskPermitLocalMouseEvents"
	case KCGEventFilterMaskPermitSystemDefinedEvents:
		return "KCGEventFilterMaskPermitSystemDefinedEvents"
	default:
		return fmt.Sprintf("CGEventFilterMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags
type CGEventFlags uint64

const (
	// KCGEventFlagMaskAlphaShift: Indicates that the Caps Lock key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskAlphaShift CGEventFlags = 65536
	// KCGEventFlagMaskAlternate: Indicates that the Alt or Option key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskAlternate CGEventFlags = 524288
	// KCGEventFlagMaskCommand: Indicates that the Command key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskCommand CGEventFlags = 1048576
	// KCGEventFlagMaskControl: Indicates that the Control key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskControl CGEventFlags = 262144
	// KCGEventFlagMaskHelp: Indicates that the Help modifier key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskHelp CGEventFlags = 4194304
	// KCGEventFlagMaskNonCoalesced: Indicates that mouse and pen movement events are not being coalesced.
	KCGEventFlagMaskNonCoalesced CGEventFlags = 256
	// KCGEventFlagMaskNumericPad: Identifies key events from the numeric keypad area on extended keyboards.
	KCGEventFlagMaskNumericPad CGEventFlags = 2097152
	// KCGEventFlagMaskSecondaryFn: Indicates that the Fn (Function) key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskSecondaryFn CGEventFlags = 8388608
	// KCGEventFlagMaskShift: Indicates that the Shift key is down for a keyboard, mouse, or flag-changed event.
	KCGEventFlagMaskShift CGEventFlags = 131072
)

func (e CGEventFlags) String() string {
	switch e {
	case KCGEventFlagMaskAlphaShift:
		return "KCGEventFlagMaskAlphaShift"
	case KCGEventFlagMaskAlternate:
		return "KCGEventFlagMaskAlternate"
	case KCGEventFlagMaskCommand:
		return "KCGEventFlagMaskCommand"
	case KCGEventFlagMaskControl:
		return "KCGEventFlagMaskControl"
	case KCGEventFlagMaskHelp:
		return "KCGEventFlagMaskHelp"
	case KCGEventFlagMaskNonCoalesced:
		return "KCGEventFlagMaskNonCoalesced"
	case KCGEventFlagMaskNumericPad:
		return "KCGEventFlagMaskNumericPad"
	case KCGEventFlagMaskSecondaryFn:
		return "KCGEventFlagMaskSecondaryFn"
	case KCGEventFlagMaskShift:
		return "KCGEventFlagMaskShift"
	default:
		return fmt.Sprintf("CGEventFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype
type CGEventMouseSubtype uint32

const (
	// KCGEventMouseSubtypeDefault: Specifies that the event is an ordinary mouse event, and does not contain additional tablet device information.
	KCGEventMouseSubtypeDefault CGEventMouseSubtype = 0
	// KCGEventMouseSubtypeTabletPoint: Specifies that the mouse event originated from a tablet device, and that the various `kCGTabletEvent` field selectors may be used to obtain tablet-specific data from the mouse event.
	KCGEventMouseSubtypeTabletPoint CGEventMouseSubtype = 1
	// KCGEventMouseSubtypeTabletProximity: # Discussion
	KCGEventMouseSubtypeTabletProximity CGEventMouseSubtype = 2
)

func (e CGEventMouseSubtype) String() string {
	switch e {
	case KCGEventMouseSubtypeDefault:
		return "KCGEventMouseSubtypeDefault"
	case KCGEventMouseSubtypeTabletPoint:
		return "KCGEventMouseSubtypeTabletPoint"
	case KCGEventMouseSubtypeTabletProximity:
		return "KCGEventMouseSubtypeTabletProximity"
	default:
		return fmt.Sprintf("CGEventMouseSubtype(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID
type CGEventSourceStateID int32

const (
	// KCGEventSourceStateCombinedSessionState: Specifies that an event source should use the event state table that reflects the combined state of all event sources posting to the current user login session.
	KCGEventSourceStateCombinedSessionState CGEventSourceStateID = 0
	// KCGEventSourceStateHIDSystemState: Specifies that an event source should use the event state table that reflects the combined state of all hardware event sources posting from the HID system.
	KCGEventSourceStateHIDSystemState CGEventSourceStateID = 1
	// KCGEventSourceStatePrivate: Specifies that an event source should use a private event state table.
	KCGEventSourceStatePrivate CGEventSourceStateID = -1
)

func (e CGEventSourceStateID) String() string {
	switch e {
	case KCGEventSourceStateCombinedSessionState:
		return "KCGEventSourceStateCombinedSessionState"
	case KCGEventSourceStateHIDSystemState:
		return "KCGEventSourceStateHIDSystemState"
	case KCGEventSourceStatePrivate:
		return "KCGEventSourceStatePrivate"
	default:
		return fmt.Sprintf("CGEventSourceStateID(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState
type CGEventSuppressionState uint32

const (
	// KCGEventSuppressionStateRemoteMouseDrag: Specifies that certain local hardware events may be suppressed during a mouse drag operation (mouse movement with the left or only mouse button down).
	KCGEventSuppressionStateRemoteMouseDrag CGEventSuppressionState = 1
	// KCGEventSuppressionStateSuppressionInterval: Specifies that certain local hardware events may be suppressed for a short interval after posting an event.
	KCGEventSuppressionStateSuppressionInterval CGEventSuppressionState = 0
	KCGNumberOfEventSuppressionStates           CGEventSuppressionState = 2
)

func (e CGEventSuppressionState) String() string {
	switch e {
	case KCGEventSuppressionStateRemoteMouseDrag:
		return "KCGEventSuppressionStateRemoteMouseDrag"
	case KCGEventSuppressionStateSuppressionInterval:
		return "KCGEventSuppressionStateSuppressionInterval"
	case KCGNumberOfEventSuppressionStates:
		return "KCGNumberOfEventSuppressionStates"
	default:
		return fmt.Sprintf("CGEventSuppressionState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation
type CGEventTapLocation uint32

const (
	// KCGAnnotatedSessionEventTap: Specifies that an event tap is placed at the point where session events have been annotated to flow to an application.
	KCGAnnotatedSessionEventTap CGEventTapLocation = 2
	// KCGHIDEventTap: Specifies that an event tap is placed at the point where HID system events enter the window server.
	KCGHIDEventTap CGEventTapLocation = 0
	// KCGSessionEventTap: Specifies that an event tap is placed at the point where HID system and remote control events enter a login session.
	KCGSessionEventTap CGEventTapLocation = 1
)

func (e CGEventTapLocation) String() string {
	switch e {
	case KCGAnnotatedSessionEventTap:
		return "KCGAnnotatedSessionEventTap"
	case KCGHIDEventTap:
		return "KCGHIDEventTap"
	case KCGSessionEventTap:
		return "KCGSessionEventTap"
	default:
		return fmt.Sprintf("CGEventTapLocation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions
type CGEventTapOptions uint32

const (
	// KCGEventTapOptionDefault: # Discussion
	KCGEventTapOptionDefault CGEventTapOptions = 0
	// KCGEventTapOptionListenOnly: # Discussion
	KCGEventTapOptionListenOnly CGEventTapOptions = 1
)

func (e CGEventTapOptions) String() string {
	switch e {
	case KCGEventTapOptionDefault:
		return "KCGEventTapOptionDefault"
	case KCGEventTapOptionListenOnly:
		return "KCGEventTapOptionListenOnly"
	default:
		return fmt.Sprintf("CGEventTapOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement
type CGEventTapPlacement uint32

const (
	// KCGHeadInsertEventTap: Specifies that a new event tap should be inserted before any pre-existing event taps at the same location.
	KCGHeadInsertEventTap CGEventTapPlacement = 0
	// KCGTailAppendEventTap: Specifies that a new event tap should be inserted after any pre-existing event taps at the same location.
	KCGTailAppendEventTap CGEventTapPlacement = 1
)

func (e CGEventTapPlacement) String() string {
	switch e {
	case KCGHeadInsertEventTap:
		return "KCGHeadInsertEventTap"
	case KCGTailAppendEventTap:
		return "KCGTailAppendEventTap"
	default:
		return fmt.Sprintf("CGEventTapPlacement(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGEventType
type CGEventType uint32

const (
	// KCGEventFlagsChanged: Specifies a key changed event for a modifier or status key.
	KCGEventFlagsChanged CGEventType = 12
	// KCGEventKeyDown: Specifies a key down event.
	KCGEventKeyDown CGEventType = 10
	// KCGEventKeyUp: Specifies a key up event.
	KCGEventKeyUp CGEventType = 11
	// KCGEventLeftMouseDown: Specifies a mouse down event with the left button.
	KCGEventLeftMouseDown CGEventType = 1
	// KCGEventLeftMouseDragged: Specifies a mouse drag event with the left button down.
	KCGEventLeftMouseDragged CGEventType = 6
	// KCGEventLeftMouseUp: Specifies a mouse up event with the left button.
	KCGEventLeftMouseUp CGEventType = 2
	// KCGEventMouseMoved: Specifies a mouse moved event.
	KCGEventMouseMoved CGEventType = 5
	// KCGEventNull: Specifies a null event.
	KCGEventNull CGEventType = 0
	// KCGEventOtherMouseDown: Specifies a mouse down event with one of buttons 2-31.
	KCGEventOtherMouseDown CGEventType = 25
	// KCGEventOtherMouseDragged: Specifies a mouse drag event with one of buttons 2-31 down.
	KCGEventOtherMouseDragged CGEventType = 27
	// KCGEventOtherMouseUp: Specifies a mouse up event with one of buttons 2-31.
	KCGEventOtherMouseUp CGEventType = 26
	// KCGEventRightMouseDown: Specifies a mouse down event with the right button.
	KCGEventRightMouseDown CGEventType = 3
	// KCGEventRightMouseDragged: Specifies a mouse drag event with the right button down.
	KCGEventRightMouseDragged CGEventType = 7
	// KCGEventRightMouseUp: Specifies a mouse up event with the right button.
	KCGEventRightMouseUp CGEventType = 4
	// KCGEventScrollWheel: Specifies a scroll wheel moved event.
	KCGEventScrollWheel CGEventType = 22
	// KCGEventTabletPointer: Specifies a tablet pointer event.
	KCGEventTabletPointer CGEventType = 23
	// KCGEventTabletProximity: Specifies a tablet proximity event.
	KCGEventTabletProximity CGEventType = 24
	// KCGEventTapDisabledByTimeout: Specifies an event indicating the event tap is disabled because of timeout.
	KCGEventTapDisabledByTimeout CGEventType = 4294967294
	// KCGEventTapDisabledByUserInput: Specifies an event indicating the event tap is disabled because of user input.
	KCGEventTapDisabledByUserInput CGEventType = 4294967295
)

func (e CGEventType) String() string {
	switch e {
	case KCGEventFlagsChanged:
		return "KCGEventFlagsChanged"
	case KCGEventKeyDown:
		return "KCGEventKeyDown"
	case KCGEventKeyUp:
		return "KCGEventKeyUp"
	case KCGEventLeftMouseDown:
		return "KCGEventLeftMouseDown"
	case KCGEventLeftMouseDragged:
		return "KCGEventLeftMouseDragged"
	case KCGEventLeftMouseUp:
		return "KCGEventLeftMouseUp"
	case KCGEventMouseMoved:
		return "KCGEventMouseMoved"
	case KCGEventNull:
		return "KCGEventNull"
	case KCGEventOtherMouseDown:
		return "KCGEventOtherMouseDown"
	case KCGEventOtherMouseDragged:
		return "KCGEventOtherMouseDragged"
	case KCGEventOtherMouseUp:
		return "KCGEventOtherMouseUp"
	case KCGEventRightMouseDown:
		return "KCGEventRightMouseDown"
	case KCGEventRightMouseDragged:
		return "KCGEventRightMouseDragged"
	case KCGEventRightMouseUp:
		return "KCGEventRightMouseUp"
	case KCGEventScrollWheel:
		return "KCGEventScrollWheel"
	case KCGEventTabletPointer:
		return "KCGEventTabletPointer"
	case KCGEventTabletProximity:
		return "KCGEventTabletProximity"
	case KCGEventTapDisabledByTimeout:
		return "KCGEventTapDisabledByTimeout"
	case KCGEventTapDisabledByUserInput:
		return "KCGEventTapDisabledByUserInput"
	default:
		return fmt.Sprintf("CGEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat
type CGFontPostScriptFormat int32

const (
	// KCGFontPostScriptFormatType1: A Type 1 font format.
	KCGFontPostScriptFormatType1 CGFontPostScriptFormat = 1
	// KCGFontPostScriptFormatType3: A Type 3 PostScript format.
	KCGFontPostScriptFormatType3 CGFontPostScriptFormat = 3
	// KCGFontPostScriptFormatType42: A constant representing a Type 42 font format.
	KCGFontPostScriptFormatType42 CGFontPostScriptFormat = 42
)

func (e CGFontPostScriptFormat) String() string {
	switch e {
	case KCGFontPostScriptFormatType1:
		return "KCGFontPostScriptFormatType1"
	case KCGFontPostScriptFormatType3:
		return "KCGFontPostScriptFormatType3"
	case KCGFontPostScriptFormatType42:
		return "KCGFontPostScriptFormatType42"
	default:
		return fmt.Sprintf("CGFontPostScriptFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase
type CGGesturePhase uint32

const (
	KCGGesturePhaseBegan     CGGesturePhase = 1
	KCGGesturePhaseCancelled CGGesturePhase = 8
	KCGGesturePhaseChanged   CGGesturePhase = 2
	KCGGesturePhaseEnded     CGGesturePhase = 4
	KCGGesturePhaseMayBegin  CGGesturePhase = 128
	KCGGesturePhaseNone      CGGesturePhase = 0
)

func (e CGGesturePhase) String() string {
	switch e {
	case KCGGesturePhaseBegan:
		return "KCGGesturePhaseBegan"
	case KCGGesturePhaseCancelled:
		return "KCGGesturePhaseCancelled"
	case KCGGesturePhaseChanged:
		return "KCGGesturePhaseChanged"
	case KCGGesturePhaseEnded:
		return "KCGGesturePhaseEnded"
	case KCGGesturePhaseMayBegin:
		return "KCGGesturePhaseMayBegin"
	case KCGGesturePhaseNone:
		return "KCGGesturePhaseNone"
	default:
		return fmt.Sprintf("CGGesturePhase(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum
type CGGlyphDeprecatedEnum int32

const (
	// CGGlyphMax: Maximum font index value.
	CGGlyphMax CGGlyphDeprecatedEnum = 1
	// CGGlyphMin: Minimum font index value.
	CGGlyphMin CGGlyphDeprecatedEnum = 0
)

func (e CGGlyphDeprecatedEnum) String() string {
	switch e {
	case CGGlyphMax:
		return "CGGlyphMax"
	case CGGlyphMin:
		return "CGGlyphMin"
	default:
		return fmt.Sprintf("CGGlyphDeprecatedEnum(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions
type CGGradientDrawingOptions uint32

const (
	// KCGGradientDrawsAfterEndLocation: The fill should extend beyond the ending location.
	KCGGradientDrawsAfterEndLocation CGGradientDrawingOptions = 2
	// KCGGradientDrawsBeforeStartLocation: The fill should extend beyond the starting location.
	KCGGradientDrawsBeforeStartLocation CGGradientDrawingOptions = 1
)

func (e CGGradientDrawingOptions) String() string {
	switch e {
	case KCGGradientDrawsAfterEndLocation:
		return "KCGGradientDrawsAfterEndLocation"
	case KCGGradientDrawsBeforeStartLocation:
		return "KCGGradientDrawsBeforeStartLocation"
	default:
		return fmt.Sprintf("CGGradientDrawingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo
type CGImageAlphaInfo uint32

const (
	// KCGImageAlphaFirst: The alpha component is stored in the most significant bits of each pixel.
	KCGImageAlphaFirst CGImageAlphaInfo = 4
	// KCGImageAlphaLast: The alpha component is stored in the least significant bits of each pixel.
	KCGImageAlphaLast CGImageAlphaInfo = 3
	// KCGImageAlphaNone: There is no alpha channel.
	KCGImageAlphaNone CGImageAlphaInfo = 0
	// KCGImageAlphaNoneSkipFirst: There is no alpha channel.
	KCGImageAlphaNoneSkipFirst CGImageAlphaInfo = 6
	// KCGImageAlphaNoneSkipLast: There is no alpha channel.
	KCGImageAlphaNoneSkipLast CGImageAlphaInfo = 5
	// KCGImageAlphaOnly: There is no color data, only an alpha channel.
	KCGImageAlphaOnly CGImageAlphaInfo = 7
	// KCGImageAlphaPremultipliedFirst: The alpha component is stored in the most significant bits of each pixel and the color components have already been multiplied by this alpha value.
	KCGImageAlphaPremultipliedFirst CGImageAlphaInfo = 2
	// KCGImageAlphaPremultipliedLast: The alpha component is stored in the least significant bits of each pixel and the color components have already been multiplied by this alpha value.
	KCGImageAlphaPremultipliedLast CGImageAlphaInfo = 1
)

func (e CGImageAlphaInfo) String() string {
	switch e {
	case KCGImageAlphaFirst:
		return "KCGImageAlphaFirst"
	case KCGImageAlphaLast:
		return "KCGImageAlphaLast"
	case KCGImageAlphaNone:
		return "KCGImageAlphaNone"
	case KCGImageAlphaNoneSkipFirst:
		return "KCGImageAlphaNoneSkipFirst"
	case KCGImageAlphaNoneSkipLast:
		return "KCGImageAlphaNoneSkipLast"
	case KCGImageAlphaOnly:
		return "KCGImageAlphaOnly"
	case KCGImageAlphaPremultipliedFirst:
		return "KCGImageAlphaPremultipliedFirst"
	case KCGImageAlphaPremultipliedLast:
		return "KCGImageAlphaPremultipliedLast"
	default:
		return fmt.Sprintf("CGImageAlphaInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo
type CGImageByteOrderInfo uint32

const (
	KCGImageByteOrder16Big    CGImageByteOrderInfo = 12288
	KCGImageByteOrder16Host   CGImageByteOrderInfo = 4096
	KCGImageByteOrder16Little CGImageByteOrderInfo = 4096
	KCGImageByteOrder32Big    CGImageByteOrderInfo = 16384
	KCGImageByteOrder32Host   CGImageByteOrderInfo = 8192
	KCGImageByteOrder32Little CGImageByteOrderInfo = 8192
	KCGImageByteOrderDefault  CGImageByteOrderInfo = 0
	// Deprecated.
	KCGImageByteOrderMask CGImageByteOrderInfo = 0
)

func (e CGImageByteOrderInfo) String() string {
	switch e {
	case KCGImageByteOrder16Big:
		return "KCGImageByteOrder16Big"
	case KCGImageByteOrder16Host:
		return "KCGImageByteOrder16Host"
	case KCGImageByteOrder32Big:
		return "KCGImageByteOrder32Big"
	case KCGImageByteOrder32Host:
		return "KCGImageByteOrder32Host"
	case KCGImageByteOrderDefault:
		return "KCGImageByteOrderDefault"
	default:
		return fmt.Sprintf("CGImageByteOrderInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo
type CGImageComponentInfo uint32

const (
	KCGImageComponentFloat   CGImageComponentInfo = 256
	KCGImageComponentInteger CGImageComponentInfo = 0
)

func (e CGImageComponentInfo) String() string {
	switch e {
	case KCGImageComponentFloat:
		return "KCGImageComponentFloat"
	case KCGImageComponentInteger:
		return "KCGImageComponentInteger"
	default:
		return fmt.Sprintf("CGImageComponentInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo
type CGImagePixelFormatInfo uint32

const (
	KCGImagePixelFormatPacked    CGImagePixelFormatInfo = 0
	KCGImagePixelFormatRGB101010 CGImagePixelFormatInfo = 196608
	KCGImagePixelFormatRGB555    CGImagePixelFormatInfo = 65536
	KCGImagePixelFormatRGB565    CGImagePixelFormatInfo = 131072
	KCGImagePixelFormatRGBCIF10  CGImagePixelFormatInfo = 262144
	// Deprecated.
	KCGImagePixelFormatMask CGImagePixelFormatInfo = 0
)

func (e CGImagePixelFormatInfo) String() string {
	switch e {
	case KCGImagePixelFormatPacked:
		return "KCGImagePixelFormatPacked"
	case KCGImagePixelFormatRGB101010:
		return "KCGImagePixelFormatRGB101010"
	case KCGImagePixelFormatRGB555:
		return "KCGImagePixelFormatRGB555"
	case KCGImagePixelFormatRGB565:
		return "KCGImagePixelFormatRGB565"
	case KCGImagePixelFormatRGBCIF10:
		return "KCGImagePixelFormatRGBCIF10"
	default:
		return fmt.Sprintf("CGImagePixelFormatInfo(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality
type CGInterpolationQuality int32

const (
	// KCGInterpolationDefault: The default level of quality.
	KCGInterpolationDefault CGInterpolationQuality = 0
	// KCGInterpolationHigh: A high level of interpolation quality.
	KCGInterpolationHigh CGInterpolationQuality = 3
	// KCGInterpolationLow: A low level of interpolation quality.
	KCGInterpolationLow CGInterpolationQuality = 2
	// KCGInterpolationMedium: A medium level of interpolation quality.
	KCGInterpolationMedium CGInterpolationQuality = 4
	// KCGInterpolationNone: No interpolation.
	KCGInterpolationNone CGInterpolationQuality = 1
)

func (e CGInterpolationQuality) String() string {
	switch e {
	case KCGInterpolationDefault:
		return "KCGInterpolationDefault"
	case KCGInterpolationHigh:
		return "KCGInterpolationHigh"
	case KCGInterpolationLow:
		return "KCGInterpolationLow"
	case KCGInterpolationMedium:
		return "KCGInterpolationMedium"
	case KCGInterpolationNone:
		return "KCGInterpolationNone"
	default:
		return fmt.Sprintf("CGInterpolationQuality(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGLineCap
type CGLineCap int32

const (
	// KCGLineCapButt: A line with a squared-off end.
	KCGLineCapButt CGLineCap = 0
	// KCGLineCapRound: A line with a rounded end.
	KCGLineCapRound CGLineCap = 1
	// KCGLineCapSquare: A line with a squared-off end.
	KCGLineCapSquare CGLineCap = 2
)

func (e CGLineCap) String() string {
	switch e {
	case KCGLineCapButt:
		return "KCGLineCapButt"
	case KCGLineCapRound:
		return "KCGLineCapRound"
	case KCGLineCapSquare:
		return "KCGLineCapSquare"
	default:
		return fmt.Sprintf("CGLineCap(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin
type CGLineJoin int32

const (
	// KCGLineJoinBevel: A join with a squared-off end.
	KCGLineJoinBevel CGLineJoin = 2
	// KCGLineJoinMiter: # Discussion
	KCGLineJoinMiter CGLineJoin = 0
	// KCGLineJoinRound: A join with a rounded end.
	KCGLineJoinRound CGLineJoin = 1
)

func (e CGLineJoin) String() string {
	switch e {
	case KCGLineJoinBevel:
		return "KCGLineJoinBevel"
	case KCGLineJoinMiter:
		return "KCGLineJoinMiter"
	case KCGLineJoinRound:
		return "KCGLineJoinRound"
	default:
		return fmt.Sprintf("CGLineJoin(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase
type CGMomentumScrollPhase uint32

const (
	KCGMomentumScrollPhaseBegin    CGMomentumScrollPhase = 1
	KCGMomentumScrollPhaseContinue CGMomentumScrollPhase = 2
	KCGMomentumScrollPhaseEnd      CGMomentumScrollPhase = 3
	KCGMomentumScrollPhaseNone     CGMomentumScrollPhase = 0
)

func (e CGMomentumScrollPhase) String() string {
	switch e {
	case KCGMomentumScrollPhaseBegin:
		return "KCGMomentumScrollPhaseBegin"
	case KCGMomentumScrollPhaseContinue:
		return "KCGMomentumScrollPhaseContinue"
	case KCGMomentumScrollPhaseEnd:
		return "KCGMomentumScrollPhaseEnd"
	case KCGMomentumScrollPhaseNone:
		return "KCGMomentumScrollPhaseNone"
	default:
		return fmt.Sprintf("CGMomentumScrollPhase(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton
type CGMouseButton uint32

const (
	// KCGMouseButtonCenter: # Discussion
	KCGMouseButtonCenter CGMouseButton = 2
	// KCGMouseButtonLeft: # Discussion
	KCGMouseButtonLeft CGMouseButton = 0
	// KCGMouseButtonRight: # Discussion
	KCGMouseButtonRight CGMouseButton = 1
)

func (e CGMouseButton) String() string {
	switch e {
	case KCGMouseButtonCenter:
		return "KCGMouseButtonCenter"
	case KCGMouseButtonLeft:
		return "KCGMouseButtonLeft"
	case KCGMouseButtonRight:
		return "KCGMouseButtonRight"
	default:
		return fmt.Sprintf("CGMouseButton(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions
type CGPDFAccessPermissions uint32

const (
	KCGPDFAllowsCommenting           CGPDFAccessPermissions = 64
	KCGPDFAllowsContentAccessibility CGPDFAccessPermissions = 32
	KCGPDFAllowsContentCopying       CGPDFAccessPermissions = 16
	KCGPDFAllowsDocumentAssembly     CGPDFAccessPermissions = 8
	KCGPDFAllowsDocumentChanges      CGPDFAccessPermissions = 4
	KCGPDFAllowsFormFieldEntry       CGPDFAccessPermissions = 128
	KCGPDFAllowsHighQualityPrinting  CGPDFAccessPermissions = 2
	KCGPDFAllowsLowQualityPrinting   CGPDFAccessPermissions = 1
)

func (e CGPDFAccessPermissions) String() string {
	switch e {
	case KCGPDFAllowsCommenting:
		return "KCGPDFAllowsCommenting"
	case KCGPDFAllowsContentAccessibility:
		return "KCGPDFAllowsContentAccessibility"
	case KCGPDFAllowsContentCopying:
		return "KCGPDFAllowsContentCopying"
	case KCGPDFAllowsDocumentAssembly:
		return "KCGPDFAllowsDocumentAssembly"
	case KCGPDFAllowsDocumentChanges:
		return "KCGPDFAllowsDocumentChanges"
	case KCGPDFAllowsFormFieldEntry:
		return "KCGPDFAllowsFormFieldEntry"
	case KCGPDFAllowsHighQualityPrinting:
		return "KCGPDFAllowsHighQualityPrinting"
	case KCGPDFAllowsLowQualityPrinting:
		return "KCGPDFAllowsLowQualityPrinting"
	default:
		return fmt.Sprintf("CGPDFAccessPermissions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox
type CGPDFBox int32

const (
	// KCGPDFArtBox: The page art box—a rectangle, expressed in default user space units, defining the extent of the page’s meaningful content (including potential white space) as intended by the page’s creator.
	KCGPDFArtBox CGPDFBox = 4
	// KCGPDFBleedBox: The page bleed box—a rectangle, expressed in default user space units, that defines the region to which the contents of the page should be clipped when output in a production environment.
	KCGPDFBleedBox CGPDFBox = 2
	// KCGPDFCropBox: The page crop box—a rectangle, expressed in default user space units, that defines the visible region of default user space.
	KCGPDFCropBox CGPDFBox = 1
	// KCGPDFMediaBox: The page media box—a rectangle, expressed in default user space units, that defines the boundaries of the physical medium on which the page is intended to be displayed or printed
	KCGPDFMediaBox CGPDFBox = 0
	// KCGPDFTrimBox: The page trim box—a rectangle, expressed in default user space units, that defines the intended dimensions of the finished page after trimming.
	KCGPDFTrimBox CGPDFBox = 3
)

func (e CGPDFBox) String() string {
	switch e {
	case KCGPDFArtBox:
		return "KCGPDFArtBox"
	case KCGPDFBleedBox:
		return "KCGPDFBleedBox"
	case KCGPDFCropBox:
		return "KCGPDFCropBox"
	case KCGPDFMediaBox:
		return "KCGPDFMediaBox"
	case KCGPDFTrimBox:
		return "KCGPDFTrimBox"
	default:
		return fmt.Sprintf("CGPDFBox(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat
type CGPDFDataFormat int32

const (
	// CGPDFDataFormatJPEG2000: The data stream is encoded in JPEG-2000 format.
	CGPDFDataFormatJPEG2000 CGPDFDataFormat = 2
	// CGPDFDataFormatJPEGEncoded: The data stream is encoded in JPEG format.
	CGPDFDataFormatJPEGEncoded CGPDFDataFormat = 1
	// CGPDFDataFormatRaw: The data stream is not encoded.
	CGPDFDataFormatRaw CGPDFDataFormat = 0
)

func (e CGPDFDataFormat) String() string {
	switch e {
	case CGPDFDataFormatJPEG2000:
		return "CGPDFDataFormatJPEG2000"
	case CGPDFDataFormatJPEGEncoded:
		return "CGPDFDataFormatJPEGEncoded"
	case CGPDFDataFormatRaw:
		return "CGPDFDataFormatRaw"
	default:
		return fmt.Sprintf("CGPDFDataFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType
type CGPDFObjectType int32

const (
	// KCGPDFObjectTypeArray: Type for a PDF array.
	KCGPDFObjectTypeArray CGPDFObjectType = 7
	// KCGPDFObjectTypeBoolean: The type for a PDF Boolean.
	KCGPDFObjectTypeBoolean CGPDFObjectType = 2
	// KCGPDFObjectTypeDictionary: The type for a PDF dictionary.
	KCGPDFObjectTypeDictionary CGPDFObjectType = 8
	// KCGPDFObjectTypeInteger: The type for a PDF integer.
	KCGPDFObjectTypeInteger CGPDFObjectType = 3
	// KCGPDFObjectTypeName: Type for a PDF name.
	KCGPDFObjectTypeName CGPDFObjectType = 5
	// KCGPDFObjectTypeNull: The type for a PDF null.
	KCGPDFObjectTypeNull CGPDFObjectType = 1
	// KCGPDFObjectTypeReal: The type for a PDF real.
	KCGPDFObjectTypeReal CGPDFObjectType = 4
	// KCGPDFObjectTypeStream: The type for a PDF stream.
	KCGPDFObjectTypeStream CGPDFObjectType = 9
	// KCGPDFObjectTypeString: The type for a PDF string.
	KCGPDFObjectTypeString CGPDFObjectType = 6
)

func (e CGPDFObjectType) String() string {
	switch e {
	case KCGPDFObjectTypeArray:
		return "KCGPDFObjectTypeArray"
	case KCGPDFObjectTypeBoolean:
		return "KCGPDFObjectTypeBoolean"
	case KCGPDFObjectTypeDictionary:
		return "KCGPDFObjectTypeDictionary"
	case KCGPDFObjectTypeInteger:
		return "KCGPDFObjectTypeInteger"
	case KCGPDFObjectTypeName:
		return "KCGPDFObjectTypeName"
	case KCGPDFObjectTypeNull:
		return "KCGPDFObjectTypeNull"
	case KCGPDFObjectTypeReal:
		return "KCGPDFObjectTypeReal"
	case KCGPDFObjectTypeStream:
		return "KCGPDFObjectTypeStream"
	case KCGPDFObjectTypeString:
		return "KCGPDFObjectTypeString"
	default:
		return fmt.Sprintf("CGPDFObjectType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType
type CGPDFTagType int32

const (
	CGPDFTagTypeAnnotation         CGPDFTagType = 507
	CGPDFTagTypeArt                CGPDFTagType = 102
	CGPDFTagTypeBibliography       CGPDFTagType = 504
	CGPDFTagTypeBlockQuote         CGPDFTagType = 105
	CGPDFTagTypeCaption            CGPDFTagType = 106
	CGPDFTagTypeCode               CGPDFTagType = 505
	CGPDFTagTypeDiv                CGPDFTagType = 104
	CGPDFTagTypeDocument           CGPDFTagType = 100
	CGPDFTagTypeFigure             CGPDFTagType = 700
	CGPDFTagTypeForm               CGPDFTagType = 702
	CGPDFTagTypeFormula            CGPDFTagType = 701
	CGPDFTagTypeHeader             CGPDFTagType = 201
	CGPDFTagTypeHeader1            CGPDFTagType = 202
	CGPDFTagTypeHeader2            CGPDFTagType = 203
	CGPDFTagTypeHeader3            CGPDFTagType = 204
	CGPDFTagTypeHeader4            CGPDFTagType = 205
	CGPDFTagTypeHeader5            CGPDFTagType = 206
	CGPDFTagTypeHeader6            CGPDFTagType = 207
	CGPDFTagTypeLabel              CGPDFTagType = 302
	CGPDFTagTypeLink               CGPDFTagType = 506
	CGPDFTagTypeList               CGPDFTagType = 300
	CGPDFTagTypeListBody           CGPDFTagType = 303
	CGPDFTagTypeListItem           CGPDFTagType = 301
	CGPDFTagTypeNonStructure       CGPDFTagType = 110
	CGPDFTagTypeNote               CGPDFTagType = 502
	CGPDFTagTypeObject             CGPDFTagType = 800
	CGPDFTagTypeParagraph          CGPDFTagType = 200
	CGPDFTagTypePart               CGPDFTagType = 101
	CGPDFTagTypePrivate            CGPDFTagType = 111
	CGPDFTagTypeQuote              CGPDFTagType = 501
	CGPDFTagTypeReference          CGPDFTagType = 503
	CGPDFTagTypeRuby               CGPDFTagType = 600
	CGPDFTagTypeRubyAnnotationText CGPDFTagType = 602
	CGPDFTagTypeRubyBaseText       CGPDFTagType = 601
	CGPDFTagTypeRubyPunctuation    CGPDFTagType = 603
	CGPDFTagTypeSection            CGPDFTagType = 103
	CGPDFTagTypeSpan               CGPDFTagType = 500
	CGPDFTagTypeTOC                CGPDFTagType = 107
	CGPDFTagTypeTOCI               CGPDFTagType = 108
	CGPDFTagTypeTable              CGPDFTagType = 400
	CGPDFTagTypeTableBody          CGPDFTagType = 405
	CGPDFTagTypeTableDataCell      CGPDFTagType = 403
	CGPDFTagTypeTableFooter        CGPDFTagType = 406
	CGPDFTagTypeTableHeader        CGPDFTagType = 404
	CGPDFTagTypeTableHeaderCell    CGPDFTagType = 402
	CGPDFTagTypeTableRow           CGPDFTagType = 401
	CGPDFTagTypeWarichu            CGPDFTagType = 604
	CGPDFTagTypeWarichuPunctiation CGPDFTagType = 606
	CGPDFTagTypeWarichuText        CGPDFTagType = 605
)

func (e CGPDFTagType) String() string {
	switch e {
	case CGPDFTagTypeAnnotation:
		return "CGPDFTagTypeAnnotation"
	case CGPDFTagTypeArt:
		return "CGPDFTagTypeArt"
	case CGPDFTagTypeBibliography:
		return "CGPDFTagTypeBibliography"
	case CGPDFTagTypeBlockQuote:
		return "CGPDFTagTypeBlockQuote"
	case CGPDFTagTypeCaption:
		return "CGPDFTagTypeCaption"
	case CGPDFTagTypeCode:
		return "CGPDFTagTypeCode"
	case CGPDFTagTypeDiv:
		return "CGPDFTagTypeDiv"
	case CGPDFTagTypeDocument:
		return "CGPDFTagTypeDocument"
	case CGPDFTagTypeFigure:
		return "CGPDFTagTypeFigure"
	case CGPDFTagTypeForm:
		return "CGPDFTagTypeForm"
	case CGPDFTagTypeFormula:
		return "CGPDFTagTypeFormula"
	case CGPDFTagTypeHeader:
		return "CGPDFTagTypeHeader"
	case CGPDFTagTypeHeader1:
		return "CGPDFTagTypeHeader1"
	case CGPDFTagTypeHeader2:
		return "CGPDFTagTypeHeader2"
	case CGPDFTagTypeHeader3:
		return "CGPDFTagTypeHeader3"
	case CGPDFTagTypeHeader4:
		return "CGPDFTagTypeHeader4"
	case CGPDFTagTypeHeader5:
		return "CGPDFTagTypeHeader5"
	case CGPDFTagTypeHeader6:
		return "CGPDFTagTypeHeader6"
	case CGPDFTagTypeLabel:
		return "CGPDFTagTypeLabel"
	case CGPDFTagTypeLink:
		return "CGPDFTagTypeLink"
	case CGPDFTagTypeList:
		return "CGPDFTagTypeList"
	case CGPDFTagTypeListBody:
		return "CGPDFTagTypeListBody"
	case CGPDFTagTypeListItem:
		return "CGPDFTagTypeListItem"
	case CGPDFTagTypeNonStructure:
		return "CGPDFTagTypeNonStructure"
	case CGPDFTagTypeNote:
		return "CGPDFTagTypeNote"
	case CGPDFTagTypeObject:
		return "CGPDFTagTypeObject"
	case CGPDFTagTypeParagraph:
		return "CGPDFTagTypeParagraph"
	case CGPDFTagTypePart:
		return "CGPDFTagTypePart"
	case CGPDFTagTypePrivate:
		return "CGPDFTagTypePrivate"
	case CGPDFTagTypeQuote:
		return "CGPDFTagTypeQuote"
	case CGPDFTagTypeReference:
		return "CGPDFTagTypeReference"
	case CGPDFTagTypeRuby:
		return "CGPDFTagTypeRuby"
	case CGPDFTagTypeRubyAnnotationText:
		return "CGPDFTagTypeRubyAnnotationText"
	case CGPDFTagTypeRubyBaseText:
		return "CGPDFTagTypeRubyBaseText"
	case CGPDFTagTypeRubyPunctuation:
		return "CGPDFTagTypeRubyPunctuation"
	case CGPDFTagTypeSection:
		return "CGPDFTagTypeSection"
	case CGPDFTagTypeSpan:
		return "CGPDFTagTypeSpan"
	case CGPDFTagTypeTOC:
		return "CGPDFTagTypeTOC"
	case CGPDFTagTypeTOCI:
		return "CGPDFTagTypeTOCI"
	case CGPDFTagTypeTable:
		return "CGPDFTagTypeTable"
	case CGPDFTagTypeTableBody:
		return "CGPDFTagTypeTableBody"
	case CGPDFTagTypeTableDataCell:
		return "CGPDFTagTypeTableDataCell"
	case CGPDFTagTypeTableFooter:
		return "CGPDFTagTypeTableFooter"
	case CGPDFTagTypeTableHeader:
		return "CGPDFTagTypeTableHeader"
	case CGPDFTagTypeTableHeaderCell:
		return "CGPDFTagTypeTableHeaderCell"
	case CGPDFTagTypeTableRow:
		return "CGPDFTagTypeTableRow"
	case CGPDFTagTypeWarichu:
		return "CGPDFTagTypeWarichu"
	case CGPDFTagTypeWarichuPunctiation:
		return "CGPDFTagTypeWarichuPunctiation"
	case CGPDFTagTypeWarichuText:
		return "CGPDFTagTypeWarichuText"
	default:
		return fmt.Sprintf("CGPDFTagType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode
type CGPathDrawingMode int32

const (
	// KCGPathEOFill: Render the area within the path using the even-odd rule.
	KCGPathEOFill CGPathDrawingMode = 1
	// KCGPathEOFillStroke: First fill and then stroke the path, using the even-odd rule.
	KCGPathEOFillStroke CGPathDrawingMode = 4
	// KCGPathFill: Render the area contained within the path using the non-zero winding number rule.
	KCGPathFill CGPathDrawingMode = 0
	// KCGPathFillStroke: First fill and then stroke the path, using the nonzero winding number rule.
	KCGPathFillStroke CGPathDrawingMode = 3
	// KCGPathStroke: Render a line along the path.
	KCGPathStroke CGPathDrawingMode = 2
)

func (e CGPathDrawingMode) String() string {
	switch e {
	case KCGPathEOFill:
		return "KCGPathEOFill"
	case KCGPathEOFillStroke:
		return "KCGPathEOFillStroke"
	case KCGPathFill:
		return "KCGPathFill"
	case KCGPathFillStroke:
		return "KCGPathFillStroke"
	case KCGPathStroke:
		return "KCGPathStroke"
	default:
		return fmt.Sprintf("CGPathDrawingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType
type CGPathElementType int32

const (
	// KCGPathElementAddCurveToPoint: The path element that adds a cubic curve from the current point to the specified point.
	KCGPathElementAddCurveToPoint CGPathElementType = 3
	// KCGPathElementAddLineToPoint: The path element that adds a line from the current point to a new point.
	KCGPathElementAddLineToPoint CGPathElementType = 1
	// KCGPathElementAddQuadCurveToPoint: The path element that adds a quadratic curve from the current point to the specified point.
	KCGPathElementAddQuadCurveToPoint CGPathElementType = 2
	// KCGPathElementCloseSubpath: The path element that closes and completes a subpath.
	KCGPathElementCloseSubpath CGPathElementType = 4
	// KCGPathElementMoveToPoint: The path element that starts a new subpath.
	KCGPathElementMoveToPoint CGPathElementType = 0
)

func (e CGPathElementType) String() string {
	switch e {
	case KCGPathElementAddCurveToPoint:
		return "KCGPathElementAddCurveToPoint"
	case KCGPathElementAddLineToPoint:
		return "KCGPathElementAddLineToPoint"
	case KCGPathElementAddQuadCurveToPoint:
		return "KCGPathElementAddQuadCurveToPoint"
	case KCGPathElementCloseSubpath:
		return "KCGPathElementCloseSubpath"
	case KCGPathElementMoveToPoint:
		return "KCGPathElementMoveToPoint"
	default:
		return fmt.Sprintf("CGPathElementType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling
type CGPatternTiling int32

const (
	// KCGPatternTilingConstantSpacing: Pattern cells are spaced consistently, as with CGPatternTiling.constantSpacingMinimalDistortion.The pattern cell may be distorted additionally to permit a moreefficient implementation.
	KCGPatternTilingConstantSpacing CGPatternTiling = 2
	// KCGPatternTilingConstantSpacingMinimalDistortion: Pattern cells are spaced consistently.
	KCGPatternTilingConstantSpacingMinimalDistortion CGPatternTiling = 1
	// KCGPatternTilingNoDistortion: The pattern cell is not distorted when painted.The spacing between pattern cells may vary by as much as 1 devicepixel.
	KCGPatternTilingNoDistortion CGPatternTiling = 0
)

func (e CGPatternTiling) String() string {
	switch e {
	case KCGPatternTilingConstantSpacing:
		return "KCGPatternTilingConstantSpacing"
	case KCGPatternTilingConstantSpacingMinimalDistortion:
		return "KCGPatternTilingConstantSpacingMinimalDistortion"
	case KCGPatternTilingNoDistortion:
		return "KCGPatternTilingNoDistortion"
	default:
		return fmt.Sprintf("CGPatternTiling(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge
type CGRectEdge uint32

const (
	CGRectMaxXEdge CGRectEdge = 2
	CGRectMaxYEdge CGRectEdge = 3
	CGRectMinXEdge CGRectEdge = 0
	CGRectMinYEdge CGRectEdge = 1
)

func (e CGRectEdge) String() string {
	switch e {
	case CGRectMaxXEdge:
		return "CGRectMaxXEdge"
	case CGRectMaxYEdge:
		return "CGRectMaxYEdge"
	case CGRectMinXEdge:
		return "CGRectMinXEdge"
	case CGRectMinYEdge:
		return "CGRectMinYEdge"
	default:
		return fmt.Sprintf("CGRectEdge(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation
type CGScreenUpdateOperation uint32

const (
	// KCGScreenUpdateOperationMove: A screen-move operation.
	KCGScreenUpdateOperationMove CGScreenUpdateOperation = 1
	// KCGScreenUpdateOperationReducedDirtyRectangleCount: # Discussion
	KCGScreenUpdateOperationReducedDirtyRectangleCount CGScreenUpdateOperation = 2147483648
	// KCGScreenUpdateOperationRefresh: A screen-refresh operation.
	KCGScreenUpdateOperationRefresh CGScreenUpdateOperation = 0
)

func (e CGScreenUpdateOperation) String() string {
	switch e {
	case KCGScreenUpdateOperationMove:
		return "KCGScreenUpdateOperationMove"
	case KCGScreenUpdateOperationReducedDirtyRectangleCount:
		return "KCGScreenUpdateOperationReducedDirtyRectangleCount"
	case KCGScreenUpdateOperationRefresh:
		return "KCGScreenUpdateOperationRefresh"
	default:
		return fmt.Sprintf("CGScreenUpdateOperation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit
type CGScrollEventUnit uint32

const (
	// KCGScrollEventUnitLine: Specifies that the unit of measurement is lines.
	KCGScrollEventUnitLine CGScrollEventUnit = 1
	// KCGScrollEventUnitPixel: Specifies that the unit of measurement is pixels.
	KCGScrollEventUnitPixel CGScrollEventUnit = 0
)

func (e CGScrollEventUnit) String() string {
	switch e {
	case KCGScrollEventUnitLine:
		return "KCGScrollEventUnitLine"
	case KCGScrollEventUnitPixel:
		return "KCGScrollEventUnitPixel"
	default:
		return fmt.Sprintf("CGScrollEventUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase
type CGScrollPhase uint32

const (
	KCGScrollPhaseBegan     CGScrollPhase = 1
	KCGScrollPhaseCancelled CGScrollPhase = 8
	KCGScrollPhaseChanged   CGScrollPhase = 2
	KCGScrollPhaseEnded     CGScrollPhase = 4
	KCGScrollPhaseMayBegin  CGScrollPhase = 128
)

func (e CGScrollPhase) String() string {
	switch e {
	case KCGScrollPhaseBegan:
		return "KCGScrollPhaseBegan"
	case KCGScrollPhaseCancelled:
		return "KCGScrollPhaseCancelled"
	case KCGScrollPhaseChanged:
		return "KCGScrollPhaseChanged"
	case KCGScrollPhaseEnded:
		return "KCGScrollPhaseEnded"
	case KCGScrollPhaseMayBegin:
		return "KCGScrollPhaseMayBegin"
	default:
		return fmt.Sprintf("CGScrollPhase(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode
type CGTextDrawingMode int32

const (
	// KCGTextClip: Specifies to intersect the text with the current clipping path.
	KCGTextClip CGTextDrawingMode = 7
	// KCGTextFill: Perform a fill operation on the text.
	KCGTextFill CGTextDrawingMode = 0
	// KCGTextFillClip: Perform a fill operation, then intersect the text with the current clipping path.
	KCGTextFillClip CGTextDrawingMode = 4
	// KCGTextFillStroke: Perform fill, then stroke operations on the text.
	KCGTextFillStroke CGTextDrawingMode = 2
	// KCGTextFillStrokeClip: Perform fill then stroke operations, then intersect the text with the current clipping path.
	KCGTextFillStrokeClip CGTextDrawingMode = 6
	// KCGTextInvisible: Do not draw the text, but do update the text position.
	KCGTextInvisible CGTextDrawingMode = 3
	// KCGTextStroke: Perform a stroke operation on the text.
	KCGTextStroke CGTextDrawingMode = 1
	// KCGTextStrokeClip: Perform a stroke operation, then intersect the text with the current clipping path.
	KCGTextStrokeClip CGTextDrawingMode = 5
)

func (e CGTextDrawingMode) String() string {
	switch e {
	case KCGTextClip:
		return "KCGTextClip"
	case KCGTextFill:
		return "KCGTextFill"
	case KCGTextFillClip:
		return "KCGTextFillClip"
	case KCGTextFillStroke:
		return "KCGTextFillStroke"
	case KCGTextFillStrokeClip:
		return "KCGTextFillStrokeClip"
	case KCGTextInvisible:
		return "KCGTextInvisible"
	case KCGTextStroke:
		return "KCGTextStroke"
	case KCGTextStrokeClip:
		return "KCGTextStrokeClip"
	default:
		return fmt.Sprintf("CGTextDrawingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding
type CGTextEncoding int32

const (
	// Deprecated.
	KCGEncodingFontSpecific CGTextEncoding = 0
	// Deprecated.
	KCGEncodingMacRoman CGTextEncoding = 1
)

func (e CGTextEncoding) String() string {
	switch e {
	case KCGEncodingFontSpecific:
		return "KCGEncodingFontSpecific"
	case KCGEncodingMacRoman:
		return "KCGEncodingMacRoman"
	default:
		return fmt.Sprintf("CGTextEncoding(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping
type CGToneMapping uint32

const (
	KCGToneMappingDefault                  CGToneMapping = 0
	KCGToneMappingEXRGamma                 CGToneMapping = 4
	KCGToneMappingITURecommended           CGToneMapping = 3
	KCGToneMappingImageSpecificLumaScaling CGToneMapping = 1
	KCGToneMappingNone                     CGToneMapping = 5
	KCGToneMappingReferenceWhiteBased      CGToneMapping = 2
)

func (e CGToneMapping) String() string {
	switch e {
	case KCGToneMappingDefault:
		return "KCGToneMappingDefault"
	case KCGToneMappingEXRGamma:
		return "KCGToneMappingEXRGamma"
	case KCGToneMappingITURecommended:
		return "KCGToneMappingITURecommended"
	case KCGToneMappingImageSpecificLumaScaling:
		return "KCGToneMappingImageSpecificLumaScaling"
	case KCGToneMappingNone:
		return "KCGToneMappingNone"
	case KCGToneMappingReferenceWhiteBased:
		return "KCGToneMappingReferenceWhiteBased"
	default:
		return fmt.Sprintf("CGToneMapping(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType
type CGWindowBackingType uint32

const (
	// KCGBackingStoreBuffered: # Discussion
	KCGBackingStoreBuffered CGWindowBackingType = 2
	// KCGBackingStoreNonretained: # Discussion
	KCGBackingStoreNonretained CGWindowBackingType = 1
	// KCGBackingStoreRetained: # Discussion
	KCGBackingStoreRetained CGWindowBackingType = 0
)

func (e CGWindowBackingType) String() string {
	switch e {
	case KCGBackingStoreBuffered:
		return "KCGBackingStoreBuffered"
	case KCGBackingStoreNonretained:
		return "KCGBackingStoreNonretained"
	case KCGBackingStoreRetained:
		return "KCGBackingStoreRetained"
	default:
		return fmt.Sprintf("CGWindowBackingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption
type CGWindowImageOption uint32

const (
	// KCGWindowImageBestResolution: When capturing the window, return the best image resolution.
	KCGWindowImageBestResolution CGWindowImageOption = 8
	// KCGWindowImageBoundsIgnoreFraming: # Discussion
	KCGWindowImageBoundsIgnoreFraming CGWindowImageOption = 1
	// KCGWindowImageDefault: # Discussion
	KCGWindowImageDefault CGWindowImageOption = 0
	// KCGWindowImageNominalResolution: When capturing the window, return the nominal image resolution.
	KCGWindowImageNominalResolution CGWindowImageOption = 16
	// KCGWindowImageOnlyShadows: # Discussion
	KCGWindowImageOnlyShadows CGWindowImageOption = 4
	// KCGWindowImageShouldBeOpaque: # Discussion
	KCGWindowImageShouldBeOpaque CGWindowImageOption = 2
)

func (e CGWindowImageOption) String() string {
	switch e {
	case KCGWindowImageBestResolution:
		return "KCGWindowImageBestResolution"
	case KCGWindowImageBoundsIgnoreFraming:
		return "KCGWindowImageBoundsIgnoreFraming"
	case KCGWindowImageDefault:
		return "KCGWindowImageDefault"
	case KCGWindowImageNominalResolution:
		return "KCGWindowImageNominalResolution"
	case KCGWindowImageOnlyShadows:
		return "KCGWindowImageOnlyShadows"
	case KCGWindowImageShouldBeOpaque:
		return "KCGWindowImageShouldBeOpaque"
	default:
		return fmt.Sprintf("CGWindowImageOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey
type CGWindowLevelKey int32

const (
	// KCGAssistiveTechHighWindowLevelKey: # Discussion
	KCGAssistiveTechHighWindowLevelKey CGWindowLevelKey = 20
	// KCGBackstopMenuLevelKey: # Discussion
	KCGBackstopMenuLevelKey CGWindowLevelKey = 3
	// KCGBaseWindowLevelKey: # Discussion
	KCGBaseWindowLevelKey CGWindowLevelKey = 0
	// KCGCursorWindowLevelKey: # Discussion
	KCGCursorWindowLevelKey CGWindowLevelKey = 19
	// KCGDesktopIconWindowLevelKey: # Discussion
	KCGDesktopIconWindowLevelKey CGWindowLevelKey = 18
	// KCGDesktopWindowLevelKey: # Discussion
	KCGDesktopWindowLevelKey CGWindowLevelKey = 2
	// KCGDockWindowLevelKey: # Discussion
	KCGDockWindowLevelKey CGWindowLevelKey = 7
	// KCGDraggingWindowLevelKey: # Discussion
	KCGDraggingWindowLevelKey CGWindowLevelKey = 12
	// KCGFloatingWindowLevelKey: # Discussion
	KCGFloatingWindowLevelKey CGWindowLevelKey = 5
	// KCGHelpWindowLevelKey: # Discussion
	KCGHelpWindowLevelKey CGWindowLevelKey = 16
	// KCGMainMenuWindowLevelKey: # Discussion
	KCGMainMenuWindowLevelKey CGWindowLevelKey = 8
	// KCGMaximumWindowLevelKey: # Discussion
	KCGMaximumWindowLevelKey CGWindowLevelKey = 14
	// KCGMinimumWindowLevelKey: # Discussion
	KCGMinimumWindowLevelKey CGWindowLevelKey = 1
	// KCGModalPanelWindowLevelKey: # Discussion
	KCGModalPanelWindowLevelKey CGWindowLevelKey = 10
	// KCGNormalWindowLevelKey: # Discussion
	KCGNormalWindowLevelKey CGWindowLevelKey = 4
	// KCGNumberOfWindowLevelKeys: # Discussion
	KCGNumberOfWindowLevelKeys CGWindowLevelKey = 21
	// KCGOverlayWindowLevelKey: # Discussion
	KCGOverlayWindowLevelKey CGWindowLevelKey = 15
	// KCGPopUpMenuWindowLevelKey: # Discussion
	KCGPopUpMenuWindowLevelKey CGWindowLevelKey = 11
	// KCGScreenSaverWindowLevelKey: # Discussion
	KCGScreenSaverWindowLevelKey CGWindowLevelKey = 13
	// KCGStatusWindowLevelKey: # Discussion
	KCGStatusWindowLevelKey CGWindowLevelKey = 9
	// KCGTornOffMenuWindowLevelKey: # Discussion
	KCGTornOffMenuWindowLevelKey CGWindowLevelKey = 6
	// KCGUtilityWindowLevelKey: # Discussion
	KCGUtilityWindowLevelKey CGWindowLevelKey = 17
)

func (e CGWindowLevelKey) String() string {
	switch e {
	case KCGAssistiveTechHighWindowLevelKey:
		return "KCGAssistiveTechHighWindowLevelKey"
	case KCGBackstopMenuLevelKey:
		return "KCGBackstopMenuLevelKey"
	case KCGBaseWindowLevelKey:
		return "KCGBaseWindowLevelKey"
	case KCGCursorWindowLevelKey:
		return "KCGCursorWindowLevelKey"
	case KCGDesktopIconWindowLevelKey:
		return "KCGDesktopIconWindowLevelKey"
	case KCGDesktopWindowLevelKey:
		return "KCGDesktopWindowLevelKey"
	case KCGDockWindowLevelKey:
		return "KCGDockWindowLevelKey"
	case KCGDraggingWindowLevelKey:
		return "KCGDraggingWindowLevelKey"
	case KCGFloatingWindowLevelKey:
		return "KCGFloatingWindowLevelKey"
	case KCGHelpWindowLevelKey:
		return "KCGHelpWindowLevelKey"
	case KCGMainMenuWindowLevelKey:
		return "KCGMainMenuWindowLevelKey"
	case KCGMaximumWindowLevelKey:
		return "KCGMaximumWindowLevelKey"
	case KCGMinimumWindowLevelKey:
		return "KCGMinimumWindowLevelKey"
	case KCGModalPanelWindowLevelKey:
		return "KCGModalPanelWindowLevelKey"
	case KCGNormalWindowLevelKey:
		return "KCGNormalWindowLevelKey"
	case KCGNumberOfWindowLevelKeys:
		return "KCGNumberOfWindowLevelKeys"
	case KCGOverlayWindowLevelKey:
		return "KCGOverlayWindowLevelKey"
	case KCGPopUpMenuWindowLevelKey:
		return "KCGPopUpMenuWindowLevelKey"
	case KCGScreenSaverWindowLevelKey:
		return "KCGScreenSaverWindowLevelKey"
	case KCGStatusWindowLevelKey:
		return "KCGStatusWindowLevelKey"
	case KCGTornOffMenuWindowLevelKey:
		return "KCGTornOffMenuWindowLevelKey"
	case KCGUtilityWindowLevelKey:
		return "KCGUtilityWindowLevelKey"
	default:
		return fmt.Sprintf("CGWindowLevelKey(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption
type CGWindowListOption uint32

const (
	// KCGWindowListExcludeDesktopElements: # Discussion
	KCGWindowListExcludeDesktopElements CGWindowListOption = 16
	// KCGWindowListOptionAll: # Discussion
	KCGWindowListOptionAll CGWindowListOption = 0
	// KCGWindowListOptionIncludingWindow: # Discussion
	KCGWindowListOptionIncludingWindow CGWindowListOption = 8
	// KCGWindowListOptionOnScreenAboveWindow: # Discussion
	KCGWindowListOptionOnScreenAboveWindow CGWindowListOption = 2
	// KCGWindowListOptionOnScreenBelowWindow: # Discussion
	KCGWindowListOptionOnScreenBelowWindow CGWindowListOption = 4
	// KCGWindowListOptionOnScreenOnly: # Discussion
	KCGWindowListOptionOnScreenOnly CGWindowListOption = 1
)

func (e CGWindowListOption) String() string {
	switch e {
	case KCGWindowListExcludeDesktopElements:
		return "KCGWindowListExcludeDesktopElements"
	case KCGWindowListOptionAll:
		return "KCGWindowListOptionAll"
	case KCGWindowListOptionIncludingWindow:
		return "KCGWindowListOptionIncludingWindow"
	case KCGWindowListOptionOnScreenAboveWindow:
		return "KCGWindowListOptionOnScreenAboveWindow"
	case KCGWindowListOptionOnScreenBelowWindow:
		return "KCGWindowListOptionOnScreenBelowWindow"
	case KCGWindowListOptionOnScreenOnly:
		return "KCGWindowListOptionOnScreenOnly"
	default:
		return fmt.Sprintf("CGWindowListOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType
type CGWindowSharingType uint32

const (
	// KCGWindowSharingNone: # Discussion
	KCGWindowSharingNone CGWindowSharingType = 0
	// KCGWindowSharingReadOnly: # Discussion
	KCGWindowSharingReadOnly CGWindowSharingType = 1
	// KCGWindowSharingReadWrite: # Discussion
	KCGWindowSharingReadWrite CGWindowSharingType = 2
)

func (e CGWindowSharingType) String() string {
	switch e {
	case KCGWindowSharingNone:
		return "KCGWindowSharingNone"
	case KCGWindowSharingReadOnly:
		return "KCGWindowSharingReadOnly"
	case KCGWindowSharingReadWrite:
		return "KCGWindowSharingReadWrite"
	default:
		return fmt.Sprintf("CGWindowSharingType(%d)", e)
	}
}
