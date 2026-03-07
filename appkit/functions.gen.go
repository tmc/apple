// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: AppKit: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: AppKit: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: AppKit: symbol %s not found\n", name)
		return
	}
	*dst = sym
}


var _nSAccessibilityActionDescription func(action NSAccessibilityActionName) foundation.NSString

// NSAccessibilityActionDescription returns a standard description for an action.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/description
func NSAccessibilityActionDescription(action NSAccessibilityActionName) foundation.NSString {
	if _nSAccessibilityActionDescription == nil {
		panic("AppKit: symbol NSAccessibilityActionDescription not loaded")
	}
	return _nSAccessibilityActionDescription(action)
}


var _nSAccessibilityFrameInView func(parentView *NSView, frame corefoundation.CGRect) corefoundation.CGRect

// NSAccessibilityFrameInView returns the frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenRect(fromView:rect:)
func NSAccessibilityFrameInView(parentView *NSView, frame corefoundation.CGRect) corefoundation.CGRect {
	if _nSAccessibilityFrameInView == nil {
		panic("AppKit: symbol NSAccessibilityFrameInView not loaded")
	}
	return _nSAccessibilityFrameInView(parentView, frame)
}


var _nSAccessibilityPointInView func(parentView *NSView, point corefoundation.CGPoint) corefoundation.CGPoint

// NSAccessibilityPointInView returns the point in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenPoint(fromView:point:)
func NSAccessibilityPointInView(parentView *NSView, point corefoundation.CGPoint) corefoundation.CGPoint {
	if _nSAccessibilityPointInView == nil {
		panic("AppKit: symbol NSAccessibilityPointInView not loaded")
	}
	return _nSAccessibilityPointInView(parentView, point)
}


var _nSAccessibilityPostNotification func(element objectivec.Object, notification NSAccessibilityNotificationName) 

// NSAccessibilityPostNotification sends a notification to any observing assistive apps.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:)
func NSAccessibilityPostNotification(element objectivec.Object, notification NSAccessibilityNotificationName)  {
	if _nSAccessibilityPostNotification == nil {
		panic("AppKit: symbol NSAccessibilityPostNotification not loaded")
	}
	_nSAccessibilityPostNotification(element, notification)
}



var _nSAccessibilityRoleDescription func(role NSAccessibilityRole, subrole NSAccessibilitySubrole) foundation.NSString

// NSAccessibilityRoleDescription returns a standard description for a role and subrole.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(with:)
func NSAccessibilityRoleDescription(role NSAccessibilityRole, subrole NSAccessibilitySubrole) foundation.NSString {
	if _nSAccessibilityRoleDescription == nil {
		panic("AppKit: symbol NSAccessibilityRoleDescription not loaded")
	}
	return _nSAccessibilityRoleDescription(role, subrole)
}


var _nSAccessibilityRoleDescriptionForUIElement func(element objectivec.Object) foundation.NSString

// NSAccessibilityRoleDescriptionForUIElement returns a standard role description for a user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(for:)
func NSAccessibilityRoleDescriptionForUIElement(element objectivec.Object) foundation.NSString {
	if _nSAccessibilityRoleDescriptionForUIElement == nil {
		panic("AppKit: symbol NSAccessibilityRoleDescriptionForUIElement not loaded")
	}
	return _nSAccessibilityRoleDescriptionForUIElement(element)
}


var _nSAccessibilitySetMayContainProtectedContent func(flag bool) bool

// NSAccessibilitySetMayContainProtectedContent sets whether the app may have protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/setMayContainProtectedContent(_:)
func NSAccessibilitySetMayContainProtectedContent(flag bool) bool {
	if _nSAccessibilitySetMayContainProtectedContent == nil {
		panic("AppKit: symbol NSAccessibilitySetMayContainProtectedContent not loaded")
	}
	return _nSAccessibilitySetMayContainProtectedContent(flag)
}


var _nSAccessibilityUnignoredAncestor func(element objectivec.Object) objectivec.Object

// NSAccessibilityUnignoredAncestor returns an unignored accessibility object, ascending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredAncestor(of:)
func NSAccessibilityUnignoredAncestor(element objectivec.Object) objectivec.Object {
	if _nSAccessibilityUnignoredAncestor == nil {
		panic("AppKit: symbol NSAccessibilityUnignoredAncestor not loaded")
	}
	return _nSAccessibilityUnignoredAncestor(element)
}


var _nSAccessibilityUnignoredChildren func(originalChildren foundation.NSArray) foundation.NSArray

// NSAccessibilityUnignoredChildren returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildren(from:)
func NSAccessibilityUnignoredChildren(originalChildren foundation.NSArray) foundation.NSArray {
	if _nSAccessibilityUnignoredChildren == nil {
		panic("AppKit: symbol NSAccessibilityUnignoredChildren not loaded")
	}
	return _nSAccessibilityUnignoredChildren(originalChildren)
}


var _nSAccessibilityUnignoredChildrenForOnlyChild func(originalChild objectivec.Object) foundation.NSArray

// NSAccessibilityUnignoredChildrenForOnlyChild returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildrenForOnlyChild(from:)
func NSAccessibilityUnignoredChildrenForOnlyChild(originalChild objectivec.Object) foundation.NSArray {
	if _nSAccessibilityUnignoredChildrenForOnlyChild == nil {
		panic("AppKit: symbol NSAccessibilityUnignoredChildrenForOnlyChild not loaded")
	}
	return _nSAccessibilityUnignoredChildrenForOnlyChild(originalChild)
}


var _nSAccessibilityUnignoredDescendant func(element objectivec.Object) objectivec.Object

// NSAccessibilityUnignoredDescendant returns an unignored accessibility object, descending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredDescendant(of:)
func NSAccessibilityUnignoredDescendant(element objectivec.Object) objectivec.Object {
	if _nSAccessibilityUnignoredDescendant == nil {
		panic("AppKit: symbol NSAccessibilityUnignoredDescendant not loaded")
	}
	return _nSAccessibilityUnignoredDescendant(element)
}


var _nSApplicationLoad func() bool

// NSApplicationLoad startup function to call when running Cocoa code from a Carbon application.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationLoad
func NSApplicationLoad() bool {
	if _nSApplicationLoad == nil {
		panic("AppKit: symbol NSApplicationLoad not loaded")
	}
	return _nSApplicationLoad()
}


var _nSApplicationMain func(argc int, argv *byte) int

// NSApplicationMain called by the main function to create and run the application.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationMain
func NSApplicationMain(argc int, argv *byte) int {
	if _nSApplicationMain == nil {
		panic("AppKit: symbol NSApplicationMain not loaded")
	}
	return _nSApplicationMain(argc, argv)
}


var _nSAvailableWindowDepths func() *NSWindowDepth

// NSAvailableWindowDepths returns the available window depth values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAvailableWindowDepths
func NSAvailableWindowDepths() *NSWindowDepth {
	if _nSAvailableWindowDepths == nil {
		panic("AppKit: symbol NSAvailableWindowDepths not loaded")
	}
	return _nSAvailableWindowDepths()
}


var _nSBeep func() 

// NSBeep plays the system beep.
//
// See: https://developer.apple.com/documentation/AppKit/NSBeep
func NSBeep()  {
	if _nSBeep == nil {
		panic("AppKit: symbol NSBeep not loaded")
	}
	_nSBeep()
}


var _nSBestDepth func(colorSpace NSColorSpaceName, bps int, bpp int, planar bool, exactMatch *bool) NSWindowDepth

// NSBestDepth attempts to return a window depth adequate for the specified parameters.
//
// See: https://developer.apple.com/documentation/AppKit/NSBestDepth
func NSBestDepth(colorSpace NSColorSpaceName, bps int, bpp int, planar bool, exactMatch *bool) NSWindowDepth {
	if _nSBestDepth == nil {
		panic("AppKit: symbol NSBestDepth not loaded")
	}
	return _nSBestDepth(colorSpace, bps, bpp, planar, exactMatch)
}


var _nSBitsPerPixelFromDepth func(depth NSWindowDepth) int

// NSBitsPerPixelFromDepth returns the bits per pixel for the specified window depth.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerPixel
func NSBitsPerPixelFromDepth(depth NSWindowDepth) int {
	if _nSBitsPerPixelFromDepth == nil {
		panic("AppKit: symbol NSBitsPerPixelFromDepth not loaded")
	}
	return _nSBitsPerPixelFromDepth(depth)
}


var _nSBitsPerSampleFromDepth func(depth NSWindowDepth) int

// NSBitsPerSampleFromDepth returns the bits per sample for the specified window depth.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerSample
func NSBitsPerSampleFromDepth(depth NSWindowDepth) int {
	if _nSBitsPerSampleFromDepth == nil {
		panic("AppKit: symbol NSBitsPerSampleFromDepth not loaded")
	}
	return _nSBitsPerSampleFromDepth(depth)
}


var _nSColorSpaceFromDepth func(depth NSWindowDepth) NSColorSpaceName

// NSColorSpaceFromDepth returns the name of the color space corresponding to the passed window depth.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/colorSpaceName
func NSColorSpaceFromDepth(depth NSWindowDepth) NSColorSpaceName {
	if _nSColorSpaceFromDepth == nil {
		panic("AppKit: symbol NSColorSpaceFromDepth not loaded")
	}
	return _nSColorSpaceFromDepth(depth)
}


var _nSConvertGlyphsToPackedGlyphs func(glBuf *NSGlyph, count int, packing NSMultibyteGlyphPacking, packedGlyphs *byte) int

// NSConvertGlyphsToPackedGlyphs prepares a set of glyphs for processing by character-based routines.
//
// See: https://developer.apple.com/documentation/AppKit/NSConvertGlyphsToPackedGlyphs(_:_:_:_:)
func NSConvertGlyphsToPackedGlyphs(glBuf *NSGlyph, count int, packing NSMultibyteGlyphPacking, packedGlyphs *byte) int {
	if _nSConvertGlyphsToPackedGlyphs == nil {
		panic("AppKit: symbol NSConvertGlyphsToPackedGlyphs not loaded")
	}
	return _nSConvertGlyphsToPackedGlyphs(glBuf, count, packing, packedGlyphs)
}


var _nSCopyBits func(srcGState int, srcRect corefoundation.CGRect, destPoint corefoundation.CGPoint) 

// NSCopyBits copies a bitmap image to the location specified by a destination point.
//
// See: https://developer.apple.com/documentation/AppKit/NSCopyBits(_:_:_:)
func NSCopyBits(srcGState int, srcRect corefoundation.CGRect, destPoint corefoundation.CGPoint)  {
	if _nSCopyBits == nil {
		panic("AppKit: symbol NSCopyBits not loaded")
	}
	_nSCopyBits(srcGState, srcRect, destPoint)
}


var _nSCountWindows func(count *int) 

// NSCountWindows counts the number of onscreen windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSCountWindows
func NSCountWindows(count *int)  {
	if _nSCountWindows == nil {
		panic("AppKit: symbol NSCountWindows not loaded")
	}
	_nSCountWindows(count)
}


var _nSCountWindowsForContext func(context int, count *int) 

// NSCountWindowsForContext counts the number of onscreen windows belonging to a particular application.
//
// See: https://developer.apple.com/documentation/AppKit/NSCountWindowsForContext
func NSCountWindowsForContext(context int, count *int)  {
	if _nSCountWindowsForContext == nil {
		panic("AppKit: symbol NSCountWindowsForContext not loaded")
	}
	_nSCountWindowsForContext(context, count)
}


var _nSCreateFileContentsPboardType func(fileType foundation.NSString) NSPasteboardType

// NSCreateFileContentsPboardType returns a pasteboard type based on the passed file type.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContentsType(forPathExtension:)
func NSCreateFileContentsPboardType(fileType foundation.NSString) NSPasteboardType {
	if _nSCreateFileContentsPboardType == nil {
		panic("AppKit: symbol NSCreateFileContentsPboardType not loaded")
	}
	return _nSCreateFileContentsPboardType(fileType)
}


var _nSCreateFilenamePboardType func(fileType foundation.NSString) NSPasteboardType

// NSCreateFilenamePboardType returns a pasteboard type based on the passed file type.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileNameType(forPathExtension:)
func NSCreateFilenamePboardType(fileType foundation.NSString) NSPasteboardType {
	if _nSCreateFilenamePboardType == nil {
		panic("AppKit: symbol NSCreateFilenamePboardType not loaded")
	}
	return _nSCreateFilenamePboardType(fileType)
}



var _nSDottedFrameRect func(rect corefoundation.CGRect) 

// NSDottedFrameRect draws a bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSDottedFrameRect(_:)
func NSDottedFrameRect(rect corefoundation.CGRect)  {
	if _nSDottedFrameRect == nil {
		panic("AppKit: symbol NSDottedFrameRect not loaded")
	}
	_nSDottedFrameRect(rect)
}


var _nSDrawBitmap func(rect corefoundation.CGRect, width int, height int, bps int, spp int, bpp int, bpr int, isPlanar bool, hasAlpha bool, colorSpaceName NSColorSpaceName, data *byte) 

// NSDrawBitmap draws a bitmap image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawBitmap(_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawBitmap(rect corefoundation.CGRect, width int, height int, bps int, spp int, bpp int, bpr int, isPlanar bool, hasAlpha bool, colorSpaceName NSColorSpaceName, data *byte)  {
	if _nSDrawBitmap == nil {
		panic("AppKit: symbol NSDrawBitmap not loaded")
	}
	_nSDrawBitmap(rect, width, height, bps, spp, bpp, bpr, isPlanar, hasAlpha, colorSpaceName, data)
}


var _nSDrawButton func(rect corefoundation.CGRect, clipRect corefoundation.CGRect) 

// NSDrawButton draws a gray-filled rectangle representing a user-interface button.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawButton(_:_:)
func NSDrawButton(rect corefoundation.CGRect, clipRect corefoundation.CGRect)  {
	if _nSDrawButton == nil {
		panic("AppKit: symbol NSDrawButton not loaded")
	}
	_nSDrawButton(rect, clipRect)
}


var _nSDrawColorTiledRects func(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, colors *NSColor, count int) corefoundation.CGRect

// NSDrawColorTiledRects draws a single-color, bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawColorTiledRects(_:_:_:_:_:)
func NSDrawColorTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, colors *NSColor, count int) corefoundation.CGRect {
	if _nSDrawColorTiledRects == nil {
		panic("AppKit: symbol NSDrawColorTiledRects not loaded")
	}
	return _nSDrawColorTiledRects(boundsRect, clipRect, sides, colors, count)
}


var _nSDrawDarkBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect) 

// NSDrawDarkBezel draws a dark gray-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawDarkBezel(_:_:)
func NSDrawDarkBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect)  {
	if _nSDrawDarkBezel == nil {
		panic("AppKit: symbol NSDrawDarkBezel not loaded")
	}
	_nSDrawDarkBezel(rect, clipRect)
}


var _nSDrawGrayBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect) 

// NSDrawGrayBezel draws a gray-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawGrayBezel(_:_:)
func NSDrawGrayBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect)  {
	if _nSDrawGrayBezel == nil {
		panic("AppKit: symbol NSDrawGrayBezel not loaded")
	}
	_nSDrawGrayBezel(rect, clipRect)
}


var _nSDrawGroove func(rect corefoundation.CGRect, clipRect corefoundation.CGRect) 

// NSDrawGroove draws a gray-filled rectangle with a groove border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawGroove(_:_:)
func NSDrawGroove(rect corefoundation.CGRect, clipRect corefoundation.CGRect)  {
	if _nSDrawGroove == nil {
		panic("AppKit: symbol NSDrawGroove not loaded")
	}
	_nSDrawGroove(rect, clipRect)
}


var _nSDrawLightBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect) 

// NSDrawLightBezel draws a white-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawLightBezel(_:_:)
func NSDrawLightBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect)  {
	if _nSDrawLightBezel == nil {
		panic("AppKit: symbol NSDrawLightBezel not loaded")
	}
	_nSDrawLightBezel(rect, clipRect)
}


var _nSDrawNinePartImage func(frame corefoundation.CGRect, topLeftCorner *NSImage, topEdgeFill *NSImage, topRightCorner *NSImage, leftEdgeFill *NSImage, centerFill *NSImage, rightEdgeFill *NSImage, bottomLeftCorner *NSImage, bottomEdgeFill *NSImage, bottomRightCorner *NSImage, op NSCompositingOperation, alphaFraction float64, flipped bool) 

// NSDrawNinePartImage draws a nine-part tiled image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawNinePartImage(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawNinePartImage(frame corefoundation.CGRect, topLeftCorner *NSImage, topEdgeFill *NSImage, topRightCorner *NSImage, leftEdgeFill *NSImage, centerFill *NSImage, rightEdgeFill *NSImage, bottomLeftCorner *NSImage, bottomEdgeFill *NSImage, bottomRightCorner *NSImage, op NSCompositingOperation, alphaFraction float64, flipped bool)  {
	if _nSDrawNinePartImage == nil {
		panic("AppKit: symbol NSDrawNinePartImage not loaded")
	}
	_nSDrawNinePartImage(frame, topLeftCorner, topEdgeFill, topRightCorner, leftEdgeFill, centerFill, rightEdgeFill, bottomLeftCorner, bottomEdgeFill, bottomRightCorner, op, alphaFraction, flipped)
}


var _nSDrawThreePartImage func(frame corefoundation.CGRect, startCap *NSImage, centerFill *NSImage, endCap *NSImage, vertical bool, op NSCompositingOperation, alphaFraction float64, flipped bool) 

// NSDrawThreePartImage draws a three-part tiled image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawThreePartImage(_:_:_:_:_:_:_:_:)
func NSDrawThreePartImage(frame corefoundation.CGRect, startCap *NSImage, centerFill *NSImage, endCap *NSImage, vertical bool, op NSCompositingOperation, alphaFraction float64, flipped bool)  {
	if _nSDrawThreePartImage == nil {
		panic("AppKit: symbol NSDrawThreePartImage not loaded")
	}
	_nSDrawThreePartImage(frame, startCap, centerFill, endCap, vertical, op, alphaFraction, flipped)
}


var _nSDrawTiledRects func(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, grays *float64, count int) corefoundation.CGRect

// NSDrawTiledRects draws rectangles with borders.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawTiledRects(_:_:_:_:_:)
func NSDrawTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, grays *float64, count int) corefoundation.CGRect {
	if _nSDrawTiledRects == nil {
		panic("AppKit: symbol NSDrawTiledRects not loaded")
	}
	return _nSDrawTiledRects(boundsRect, clipRect, sides, grays, count)
}


var _nSDrawWhiteBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect) 

// NSDrawWhiteBezel draws a white-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawWhiteBezel(_:_:)
func NSDrawWhiteBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect)  {
	if _nSDrawWhiteBezel == nil {
		panic("AppKit: symbol NSDrawWhiteBezel not loaded")
	}
	_nSDrawWhiteBezel(rect, clipRect)
}


var _nSDrawWindowBackground func(rect corefoundation.CGRect) 

// NSDrawWindowBackground draws the window’s default background pattern into the specified rectangle of the currently focused view.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawWindowBackground(_:)
func NSDrawWindowBackground(rect corefoundation.CGRect)  {
	if _nSDrawWindowBackground == nil {
		panic("AppKit: symbol NSDrawWindowBackground not loaded")
	}
	_nSDrawWindowBackground(rect)
}


var _nSEraseRect func(rect corefoundation.CGRect) 

// NSEraseRect erases the specified rect by filling it with white.
//
// See: https://developer.apple.com/documentation/AppKit/NSEraseRect(_:)
func NSEraseRect(rect corefoundation.CGRect)  {
	if _nSEraseRect == nil {
		panic("AppKit: symbol NSEraseRect not loaded")
	}
	_nSEraseRect(rect)
}



var _nSFrameRect func(rect corefoundation.CGRect) 

// NSFrameRect draws a bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSFrameRect
func NSFrameRect(rect corefoundation.CGRect)  {
	if _nSFrameRect == nil {
		panic("AppKit: symbol NSFrameRect not loaded")
	}
	_nSFrameRect(rect)
}


var _nSFrameRectWithWidth func(rect corefoundation.CGRect, frameWidth float64) 

// NSFrameRectWithWidth draws a bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidth
func NSFrameRectWithWidth(rect corefoundation.CGRect, frameWidth float64)  {
	if _nSFrameRectWithWidth == nil {
		panic("AppKit: symbol NSFrameRectWithWidth not loaded")
	}
	_nSFrameRectWithWidth(rect, frameWidth)
}


var _nSFrameRectWithWidthUsingOperation func(rect corefoundation.CGRect, frameWidth float64, op NSCompositingOperation) 

// NSFrameRectWithWidthUsingOperation draws a bordered rectangle using the specified compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidthUsingOperation
func NSFrameRectWithWidthUsingOperation(rect corefoundation.CGRect, frameWidth float64, op NSCompositingOperation)  {
	if _nSFrameRectWithWidthUsingOperation == nil {
		panic("AppKit: symbol NSFrameRectWithWidthUsingOperation not loaded")
	}
	_nSFrameRectWithWidthUsingOperation(rect, frameWidth, op)
}


var _nSGetFileType func(pboardType NSPasteboardType) foundation.NSString

// NSGetFileType a file type based on the passed pasteboard type.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtension
func NSGetFileType(pboardType NSPasteboardType) foundation.NSString {
	if _nSGetFileType == nil {
		panic("AppKit: symbol NSGetFileType not loaded")
	}
	return _nSGetFileType(pboardType)
}


var _nSGetFileTypes func(pboardTypes *foundation.NSString) []foundation.NSString

// NSGetFileTypes returns an array of file types based on the passed pasteboard types.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtensions(from:)
func NSGetFileTypes(pboardTypes []foundation.NSString) []foundation.NSString {
	if _nSGetFileTypes == nil {
		panic("AppKit: symbol NSGetFileTypes not loaded")
	}
	return _nSGetFileTypes(unsafe.SliceData(pboardTypes))
}


var _nSHighlightRect func(rect corefoundation.CGRect) 

// NSHighlightRect highlights the specified rect by filling it with white.
//
// See: https://developer.apple.com/documentation/AppKit/NSHighlightRect
func NSHighlightRect(rect corefoundation.CGRect)  {
	if _nSHighlightRect == nil {
		panic("AppKit: symbol NSHighlightRect not loaded")
	}
	_nSHighlightRect(rect)
}



var _nSIsControllerMarker func(object objectivec.Object) bool

// NSIsControllerMarker tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.
//
// See: https://developer.apple.com/documentation/AppKit/NSIsControllerMarker(_:)
func NSIsControllerMarker(object objectivec.Object) bool {
	if _nSIsControllerMarker == nil {
		panic("AppKit: symbol NSIsControllerMarker not loaded")
	}
	return _nSIsControllerMarker(object)
}


var _nSNumberOfColorComponents func(colorSpaceName NSColorSpaceName) int

// NSNumberOfColorComponents returns the number of color components in the specified color space.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/numberOfColorComponents
func NSNumberOfColorComponents(colorSpaceName NSColorSpaceName) int {
	if _nSNumberOfColorComponents == nil {
		panic("AppKit: symbol NSNumberOfColorComponents not loaded")
	}
	return _nSNumberOfColorComponents(colorSpaceName)
}


var _nSPerformService func(itemName foundation.NSString, pboard *NSPasteboard) bool

// NSPerformService programmatically invokes a Services menu service.
//
// See: https://developer.apple.com/documentation/AppKit/NSPerformService(_:_:)
func NSPerformService(itemName foundation.NSString, pboard *NSPasteboard) bool {
	if _nSPerformService == nil {
		panic("AppKit: symbol NSPerformService not loaded")
	}
	return _nSPerformService(itemName, pboard)
}


var _nSPlanarFromDepth func(depth NSWindowDepth) bool

// NSPlanarFromDepth returns whether the specified window depth is planar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/isPlanar
func NSPlanarFromDepth(depth NSWindowDepth) bool {
	if _nSPlanarFromDepth == nil {
		panic("AppKit: symbol NSPlanarFromDepth not loaded")
	}
	return _nSPlanarFromDepth(depth)
}


var _nSRectClip func(rect corefoundation.CGRect) 

// NSRectClip modifies the current clipping path by intersecting it with the passed rect.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectClip
func NSRectClip(rect corefoundation.CGRect)  {
	if _nSRectClip == nil {
		panic("AppKit: symbol NSRectClip not loaded")
	}
	_nSRectClip(rect)
}


var _nSRectClipList func(rects *corefoundation.CGRect, count int) 

// NSRectClipList modifies the current clipping path by intersecting it with the passed rect.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectClipList
func NSRectClipList(rects *corefoundation.CGRect, count int)  {
	if _nSRectClipList == nil {
		panic("AppKit: symbol NSRectClipList not loaded")
	}
	_nSRectClipList(rects, count)
}


var _nSRectFill func(rect corefoundation.CGRect) 

// NSRectFill fills the passed rectangle with the current color.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFill
func NSRectFill(rect corefoundation.CGRect)  {
	if _nSRectFill == nil {
		panic("AppKit: symbol NSRectFill not loaded")
	}
	_nSRectFill(rect)
}


var _nSRectFillList func(rects *corefoundation.CGRect, count int) 

// NSRectFillList fills the rectangles in the passed list with the current fill color.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillList
func NSRectFillList(rects *corefoundation.CGRect, count int)  {
	if _nSRectFillList == nil {
		panic("AppKit: symbol NSRectFillList not loaded")
	}
	_nSRectFillList(rects, count)
}


var _nSRectFillListUsingOperation func(rects *corefoundation.CGRect, count int, op NSCompositingOperation) 

// NSRectFillListUsingOperation fills the rectangles in a list using the current fill color and specified compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListUsingOperation
func NSRectFillListUsingOperation(rects *corefoundation.CGRect, count int, op NSCompositingOperation)  {
	if _nSRectFillListUsingOperation == nil {
		panic("AppKit: symbol NSRectFillListUsingOperation not loaded")
	}
	_nSRectFillListUsingOperation(rects, count, op)
}


var _nSRectFillListWithColors func(rects *corefoundation.CGRect, colors NSColor, num int) 

// NSRectFillListWithColors fills the rectangles in the passed list with the passed list of colors.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColors
func NSRectFillListWithColors(rects *corefoundation.CGRect, colors NSColor, num int)  {
	if _nSRectFillListWithColors == nil {
		panic("AppKit: symbol NSRectFillListWithColors not loaded")
	}
	_nSRectFillListWithColors(rects, colors, num)
}


var _nSRectFillListWithColorsUsingOperation func(rects *corefoundation.CGRect, colors NSColor, num int, op NSCompositingOperation) 

// NSRectFillListWithColorsUsingOperation fills the rectangles in a list using the specified colors and compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColorsUsingOperation
func NSRectFillListWithColorsUsingOperation(rects *corefoundation.CGRect, colors NSColor, num int, op NSCompositingOperation)  {
	if _nSRectFillListWithColorsUsingOperation == nil {
		panic("AppKit: symbol NSRectFillListWithColorsUsingOperation not loaded")
	}
	_nSRectFillListWithColorsUsingOperation(rects, colors, num, op)
}


var _nSRectFillListWithGrays func(rects *corefoundation.CGRect, grays *float64, num int) 

// NSRectFillListWithGrays fills the rectangles in the passed list with the passed list of grays.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListWithGrays
func NSRectFillListWithGrays(rects *corefoundation.CGRect, grays *float64, num int)  {
	if _nSRectFillListWithGrays == nil {
		panic("AppKit: symbol NSRectFillListWithGrays not loaded")
	}
	_nSRectFillListWithGrays(rects, grays, num)
}


var _nSRectFillUsingOperation func(rect corefoundation.CGRect, op NSCompositingOperation) 

// NSRectFillUsingOperation fills a rectangle using the current fill color and the specified compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillUsingOperation
func NSRectFillUsingOperation(rect corefoundation.CGRect, op NSCompositingOperation)  {
	if _nSRectFillUsingOperation == nil {
		panic("AppKit: symbol NSRectFillUsingOperation not loaded")
	}
	_nSRectFillUsingOperation(rect, op)
}


var _nSRegisterServicesProvider func(provider objectivec.Object, name NSServiceProviderName) 

// NSRegisterServicesProvider registers a service provider.
//
// See: https://developer.apple.com/documentation/AppKit/NSRegisterServicesProvider(_:_:)
func NSRegisterServicesProvider(provider objectivec.Object, name NSServiceProviderName)  {
	if _nSRegisterServicesProvider == nil {
		panic("AppKit: symbol NSRegisterServicesProvider not loaded")
	}
	_nSRegisterServicesProvider(provider, name)
}


var _nSRunAlertPanelRelativeToWindow func(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int

// NSRunAlertPanelRelativeToWindow.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunAlertPanelRelativeToWindow
func NSRunAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int {
	if _nSRunAlertPanelRelativeToWindow == nil {
		panic("AppKit: symbol NSRunAlertPanelRelativeToWindow not loaded")
	}
	return _nSRunAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}


var _nSRunCriticalAlertPanelRelativeToWindow func(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int

// NSRunCriticalAlertPanelRelativeToWindow.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunCriticalAlertPanelRelativeToWindow
func NSRunCriticalAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int {
	if _nSRunCriticalAlertPanelRelativeToWindow == nil {
		panic("AppKit: symbol NSRunCriticalAlertPanelRelativeToWindow not loaded")
	}
	return _nSRunCriticalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}


var _nSRunInformationalAlertPanelRelativeToWindow func(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int

// NSRunInformationalAlertPanelRelativeToWindow.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunInformationalAlertPanelRelativeToWindow
func NSRunInformationalAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int {
	if _nSRunInformationalAlertPanelRelativeToWindow == nil {
		panic("AppKit: symbol NSRunInformationalAlertPanelRelativeToWindow not loaded")
	}
	return _nSRunInformationalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
}


var _nSSetShowsServicesMenuItem func(itemName foundation.NSString, enabled bool) int

// NSSetShowsServicesMenuItem specifies whether an item should be included in Services menus.
//
// See: https://developer.apple.com/documentation/AppKit/NSSetShowsServicesMenuItem(_:_:)
func NSSetShowsServicesMenuItem(itemName foundation.NSString, enabled bool) int {
	if _nSSetShowsServicesMenuItem == nil {
		panic("AppKit: symbol NSSetShowsServicesMenuItem not loaded")
	}
	return _nSSetShowsServicesMenuItem(itemName, enabled)
}


var _nSShowAnimationEffect func(animationEffect NSAnimationEffect, centerLocation corefoundation.CGPoint, size corefoundation.CGSize, animationDelegate objectivec.Object, didEndSelector objectivec.SEL, contextInfo uintptr) 

// NSShowAnimationEffect runs a system animation effect.
//
// Deprecated: Deprecated since macOS 14.0. Use [disappearingItem](<doc://com.apple.appkit/documentation/AppKit/NSCursor/disappearingItem>) instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSShowAnimationEffect
func NSShowAnimationEffect(animationEffect NSAnimationEffect, centerLocation corefoundation.CGPoint, size corefoundation.CGSize, animationDelegate objectivec.Object, didEndSelector objectivec.SEL, contextInfo uintptr)  {
	if _nSShowAnimationEffect == nil {
		panic("AppKit: symbol NSShowAnimationEffect not loaded")
	}
	_nSShowAnimationEffect(animationEffect, centerLocation, size, animationDelegate, didEndSelector, contextInfo)
}


var _nSShowsServicesMenuItem func(itemName foundation.NSString) bool

// NSShowsServicesMenuItem specifies whether a Services menu item is currently enabled.
//
// See: https://developer.apple.com/documentation/AppKit/NSShowsServicesMenuItem(_:)
func NSShowsServicesMenuItem(itemName foundation.NSString) bool {
	if _nSShowsServicesMenuItem == nil {
		panic("AppKit: symbol NSShowsServicesMenuItem not loaded")
	}
	return _nSShowsServicesMenuItem(itemName)
}



var _nSUnregisterServicesProvider func(name NSServiceProviderName) 

// NSUnregisterServicesProvider unregisters a service provider.
//
// See: https://developer.apple.com/documentation/AppKit/NSUnregisterServicesProvider(_:)
func NSUnregisterServicesProvider(name NSServiceProviderName)  {
	if _nSUnregisterServicesProvider == nil {
		panic("AppKit: symbol NSUnregisterServicesProvider not loaded")
	}
	_nSUnregisterServicesProvider(name)
}


var _nSUpdateDynamicServices func() 

// NSUpdateDynamicServices causes the services information for the system to be updated.
//
// See: https://developer.apple.com/documentation/AppKit/NSUpdateDynamicServices()
func NSUpdateDynamicServices()  {
	if _nSUpdateDynamicServices == nil {
		panic("AppKit: symbol NSUpdateDynamicServices not loaded")
	}
	_nSUpdateDynamicServices()
}


var _nSWindowList func(size int, list int) 

// NSWindowList gets information about onscreen windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowList
func NSWindowList(size int, list int)  {
	if _nSWindowList == nil {
		panic("AppKit: symbol NSWindowList not loaded")
	}
	_nSWindowList(size, list)
}


var _nSWindowListForContext func(context int, size int, list int) 

// NSWindowListForContext gets information about an application’s onscreen windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowListForContext
func NSWindowListForContext(context int, size int, list int)  {
	if _nSWindowListForContext == nil {
		panic("AppKit: symbol NSWindowListForContext not loaded")
	}
	_nSWindowListForContext(context, size, list)
}



func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_nSAccessibilityActionDescription, frameworkHandle, "NSAccessibilityActionDescription")
		registerFunc(&_nSAccessibilityFrameInView, frameworkHandle, "NSAccessibilityFrameInView")
		registerFunc(&_nSAccessibilityPointInView, frameworkHandle, "NSAccessibilityPointInView")
		registerFunc(&_nSAccessibilityPostNotification, frameworkHandle, "NSAccessibilityPostNotification")
		registerFunc(&_nSAccessibilityRoleDescription, frameworkHandle, "NSAccessibilityRoleDescription")
		registerFunc(&_nSAccessibilityRoleDescriptionForUIElement, frameworkHandle, "NSAccessibilityRoleDescriptionForUIElement")
		registerFunc(&_nSAccessibilitySetMayContainProtectedContent, frameworkHandle, "NSAccessibilitySetMayContainProtectedContent")
		registerFunc(&_nSAccessibilityUnignoredAncestor, frameworkHandle, "NSAccessibilityUnignoredAncestor")
		registerFunc(&_nSAccessibilityUnignoredChildren, frameworkHandle, "NSAccessibilityUnignoredChildren")
		registerFunc(&_nSAccessibilityUnignoredChildrenForOnlyChild, frameworkHandle, "NSAccessibilityUnignoredChildrenForOnlyChild")
		registerFunc(&_nSAccessibilityUnignoredDescendant, frameworkHandle, "NSAccessibilityUnignoredDescendant")
		registerFunc(&_nSApplicationLoad, frameworkHandle, "NSApplicationLoad")
		registerFunc(&_nSApplicationMain, frameworkHandle, "NSApplicationMain")
		registerFunc(&_nSAvailableWindowDepths, frameworkHandle, "NSAvailableWindowDepths")
		registerFunc(&_nSBeep, frameworkHandle, "NSBeep")
		registerFunc(&_nSBestDepth, frameworkHandle, "NSBestDepth")
		registerFunc(&_nSBitsPerPixelFromDepth, frameworkHandle, "NSBitsPerPixelFromDepth")
		registerFunc(&_nSBitsPerSampleFromDepth, frameworkHandle, "NSBitsPerSampleFromDepth")
		registerFunc(&_nSColorSpaceFromDepth, frameworkHandle, "NSColorSpaceFromDepth")
		registerFunc(&_nSConvertGlyphsToPackedGlyphs, frameworkHandle, "NSConvertGlyphsToPackedGlyphs")
		registerFunc(&_nSCopyBits, frameworkHandle, "NSCopyBits")
		registerFunc(&_nSCountWindows, frameworkHandle, "NSCountWindows")
		registerFunc(&_nSCountWindowsForContext, frameworkHandle, "NSCountWindowsForContext")
		registerFunc(&_nSCreateFileContentsPboardType, frameworkHandle, "NSCreateFileContentsPboardType")
		registerFunc(&_nSCreateFilenamePboardType, frameworkHandle, "NSCreateFilenamePboardType")
		registerFunc(&_nSDottedFrameRect, frameworkHandle, "NSDottedFrameRect")
		registerFunc(&_nSDrawBitmap, frameworkHandle, "NSDrawBitmap")
		registerFunc(&_nSDrawButton, frameworkHandle, "NSDrawButton")
		registerFunc(&_nSDrawColorTiledRects, frameworkHandle, "NSDrawColorTiledRects")
		registerFunc(&_nSDrawDarkBezel, frameworkHandle, "NSDrawDarkBezel")
		registerFunc(&_nSDrawGrayBezel, frameworkHandle, "NSDrawGrayBezel")
		registerFunc(&_nSDrawGroove, frameworkHandle, "NSDrawGroove")
		registerFunc(&_nSDrawLightBezel, frameworkHandle, "NSDrawLightBezel")
		registerFunc(&_nSDrawNinePartImage, frameworkHandle, "NSDrawNinePartImage")
		registerFunc(&_nSDrawThreePartImage, frameworkHandle, "NSDrawThreePartImage")
		registerFunc(&_nSDrawTiledRects, frameworkHandle, "NSDrawTiledRects")
		registerFunc(&_nSDrawWhiteBezel, frameworkHandle, "NSDrawWhiteBezel")
		registerFunc(&_nSDrawWindowBackground, frameworkHandle, "NSDrawWindowBackground")
		registerFunc(&_nSEraseRect, frameworkHandle, "NSEraseRect")
		registerFunc(&_nSFrameRect, frameworkHandle, "NSFrameRect")
		registerFunc(&_nSFrameRectWithWidth, frameworkHandle, "NSFrameRectWithWidth")
		registerFunc(&_nSFrameRectWithWidthUsingOperation, frameworkHandle, "NSFrameRectWithWidthUsingOperation")
		registerFunc(&_nSGetFileType, frameworkHandle, "NSGetFileType")
		registerFunc(&_nSGetFileTypes, frameworkHandle, "NSGetFileTypes")
		registerFunc(&_nSHighlightRect, frameworkHandle, "NSHighlightRect")
		registerFunc(&_nSIsControllerMarker, frameworkHandle, "NSIsControllerMarker")
		registerFunc(&_nSNumberOfColorComponents, frameworkHandle, "NSNumberOfColorComponents")
		registerFunc(&_nSPerformService, frameworkHandle, "NSPerformService")
		registerFunc(&_nSPlanarFromDepth, frameworkHandle, "NSPlanarFromDepth")
		registerFunc(&_nSRectClip, frameworkHandle, "NSRectClip")
		registerFunc(&_nSRectClipList, frameworkHandle, "NSRectClipList")
		registerFunc(&_nSRectFill, frameworkHandle, "NSRectFill")
		registerFunc(&_nSRectFillList, frameworkHandle, "NSRectFillList")
		registerFunc(&_nSRectFillListUsingOperation, frameworkHandle, "NSRectFillListUsingOperation")
		registerFunc(&_nSRectFillListWithColors, frameworkHandle, "NSRectFillListWithColors")
		registerFunc(&_nSRectFillListWithColorsUsingOperation, frameworkHandle, "NSRectFillListWithColorsUsingOperation")
		registerFunc(&_nSRectFillListWithGrays, frameworkHandle, "NSRectFillListWithGrays")
		registerFunc(&_nSRectFillUsingOperation, frameworkHandle, "NSRectFillUsingOperation")
		registerFunc(&_nSRegisterServicesProvider, frameworkHandle, "NSRegisterServicesProvider")
		registerFunc(&_nSRunAlertPanelRelativeToWindow, frameworkHandle, "NSRunAlertPanelRelativeToWindow")
		registerFunc(&_nSRunCriticalAlertPanelRelativeToWindow, frameworkHandle, "NSRunCriticalAlertPanelRelativeToWindow")
		registerFunc(&_nSRunInformationalAlertPanelRelativeToWindow, frameworkHandle, "NSRunInformationalAlertPanelRelativeToWindow")
		registerFunc(&_nSSetShowsServicesMenuItem, frameworkHandle, "NSSetShowsServicesMenuItem")
		registerFunc(&_nSShowAnimationEffect, frameworkHandle, "NSShowAnimationEffect")
		registerFunc(&_nSShowsServicesMenuItem, frameworkHandle, "NSShowsServicesMenuItem")
		registerFunc(&_nSUnregisterServicesProvider, frameworkHandle, "NSUnregisterServicesProvider")
		registerFunc(&_nSUpdateDynamicServices, frameworkHandle, "NSUpdateDynamicServices")
		registerFunc(&_nSWindowList, frameworkHandle, "NSWindowList")
		registerFunc(&_nSWindowListForContext, frameworkHandle, "NSWindowListForContext")
	}


