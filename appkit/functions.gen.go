// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("AppKit: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("AppKit: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("AppKit: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("AppKit: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _nSAccessibilityActionDescription func(action NSAccessibilityActionName) foundation.NSString
var _nSAccessibilityActionDescriptionErr error

func tryNSAccessibilityActionDescription(action NSAccessibilityActionName) (foundation.NSString, error) {
	if _nSAccessibilityActionDescription == nil {
		return foundation.NSString{}, symbolCallError("NSAccessibilityActionDescription", "", _nSAccessibilityActionDescriptionErr)
	}
	return _nSAccessibilityActionDescription(action), nil
}

// NSAccessibilityActionDescription returns a standard description for an action.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action/description
func NSAccessibilityActionDescription(action NSAccessibilityActionName) foundation.NSString {
	result, callErr := tryNSAccessibilityActionDescription(action)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityFrameInView func(parentView *NSView, frame corefoundation.CGRect) corefoundation.CGRect
var _nSAccessibilityFrameInViewErr error

func tryNSAccessibilityFrameInView(parentView *NSView, frame corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _nSAccessibilityFrameInView == nil {
		return corefoundation.CGRect{}, symbolCallError("NSAccessibilityFrameInView", "10.10", _nSAccessibilityFrameInViewErr)
	}
	return _nSAccessibilityFrameInView(parentView, frame), nil
}

// NSAccessibilityFrameInView returns the frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenRect(fromView:rect:)
func NSAccessibilityFrameInView(parentView *NSView, frame corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryNSAccessibilityFrameInView(parentView, frame)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityPointInView func(parentView *NSView, point corefoundation.CGPoint) corefoundation.CGPoint
var _nSAccessibilityPointInViewErr error

func tryNSAccessibilityPointInView(parentView *NSView, point corefoundation.CGPoint) (corefoundation.CGPoint, error) {
	if _nSAccessibilityPointInView == nil {
		return corefoundation.CGPoint{}, symbolCallError("NSAccessibilityPointInView", "10.10", _nSAccessibilityPointInViewErr)
	}
	return _nSAccessibilityPointInView(parentView, point), nil
}

// NSAccessibilityPointInView returns the point in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenPoint(fromView:point:)
func NSAccessibilityPointInView(parentView *NSView, point corefoundation.CGPoint) corefoundation.CGPoint {
	result, callErr := tryNSAccessibilityPointInView(parentView, point)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityPostNotification func(element objectivec.Object, notification NSAccessibilityNotificationName)
var _nSAccessibilityPostNotificationErr error

func tryNSAccessibilityPostNotification(element objectivec.Object, notification NSAccessibilityNotificationName) error {
	if _nSAccessibilityPostNotification == nil {
		return symbolCallError("NSAccessibilityPostNotification", "", _nSAccessibilityPostNotificationErr)
	}
	_nSAccessibilityPostNotification(element, notification)
	return nil
}

// NSAccessibilityPostNotification sends a notification to any observing assistive apps.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:)
func NSAccessibilityPostNotification(element objectivec.Object, notification NSAccessibilityNotificationName) {
	if callErr := tryNSAccessibilityPostNotification(element, notification); callErr != nil {
		panic(callErr)
	}
}

var _nSAccessibilityPostNotificationWithUserInfo func(element objectivec.Object, notification NSAccessibilityNotificationName, userInfo uintptr)
var _nSAccessibilityPostNotificationWithUserInfoErr error

func tryNSAccessibilityPostNotificationWithUserInfo(element objectivec.Object, notification NSAccessibilityNotificationName, userInfo uintptr) error {
	if _nSAccessibilityPostNotificationWithUserInfo == nil {
		return symbolCallError("NSAccessibilityPostNotificationWithUserInfo", "10.7", _nSAccessibilityPostNotificationWithUserInfoErr)
	}
	_nSAccessibilityPostNotificationWithUserInfo(element, notification, userInfo)
	return nil
}

// NSAccessibilityPostNotificationWithUserInfo sends a notification and an optional user info dictionary to any observing assistive apps.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:userInfo:)
func NSAccessibilityPostNotificationWithUserInfo(element objectivec.Object, notification NSAccessibilityNotificationName, userInfo uintptr) {
	if callErr := tryNSAccessibilityPostNotificationWithUserInfo(element, notification, userInfo); callErr != nil {
		panic(callErr)
	}
}

var _nSAccessibilityRoleDescription func(role NSAccessibilityRole, subrole NSAccessibilitySubrole) foundation.NSString
var _nSAccessibilityRoleDescriptionErr error

func tryNSAccessibilityRoleDescription(role NSAccessibilityRole, subrole NSAccessibilitySubrole) (foundation.NSString, error) {
	if _nSAccessibilityRoleDescription == nil {
		return foundation.NSString{}, symbolCallError("NSAccessibilityRoleDescription", "", _nSAccessibilityRoleDescriptionErr)
	}
	return _nSAccessibilityRoleDescription(role, subrole), nil
}

// NSAccessibilityRoleDescription returns a standard description for a role and subrole.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(with:)
func NSAccessibilityRoleDescription(role NSAccessibilityRole, subrole NSAccessibilitySubrole) foundation.NSString {
	result, callErr := tryNSAccessibilityRoleDescription(role, subrole)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityRoleDescriptionForUIElement func(element objectivec.Object) foundation.NSString
var _nSAccessibilityRoleDescriptionForUIElementErr error

func tryNSAccessibilityRoleDescriptionForUIElement(element objectivec.Object) (foundation.NSString, error) {
	if _nSAccessibilityRoleDescriptionForUIElement == nil {
		return foundation.NSString{}, symbolCallError("NSAccessibilityRoleDescriptionForUIElement", "", _nSAccessibilityRoleDescriptionForUIElementErr)
	}
	return _nSAccessibilityRoleDescriptionForUIElement(element), nil
}

// NSAccessibilityRoleDescriptionForUIElement returns a standard role description for a user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/description(for:)
func NSAccessibilityRoleDescriptionForUIElement(element objectivec.Object) foundation.NSString {
	result, callErr := tryNSAccessibilityRoleDescriptionForUIElement(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilitySetMayContainProtectedContent func(flag bool) bool
var _nSAccessibilitySetMayContainProtectedContentErr error

func tryNSAccessibilitySetMayContainProtectedContent(flag bool) (bool, error) {
	if _nSAccessibilitySetMayContainProtectedContent == nil {
		return false, symbolCallError("NSAccessibilitySetMayContainProtectedContent", "", _nSAccessibilitySetMayContainProtectedContentErr)
	}
	return _nSAccessibilitySetMayContainProtectedContent(flag), nil
}

// NSAccessibilitySetMayContainProtectedContent sets whether the app may have protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/setMayContainProtectedContent(_:)
func NSAccessibilitySetMayContainProtectedContent(flag bool) bool {
	result, callErr := tryNSAccessibilitySetMayContainProtectedContent(flag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityUnignoredAncestor func(element objectivec.Object) objectivec.Object
var _nSAccessibilityUnignoredAncestorErr error

func tryNSAccessibilityUnignoredAncestor(element objectivec.Object) (objectivec.Object, error) {
	if _nSAccessibilityUnignoredAncestor == nil {
		return objectivec.Object{}, symbolCallError("NSAccessibilityUnignoredAncestor", "", _nSAccessibilityUnignoredAncestorErr)
	}
	return _nSAccessibilityUnignoredAncestor(element), nil
}

// NSAccessibilityUnignoredAncestor returns an unignored accessibility object, ascending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredAncestor(of:)
func NSAccessibilityUnignoredAncestor(element objectivec.Object) objectivec.Object {
	result, callErr := tryNSAccessibilityUnignoredAncestor(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityUnignoredChildren func(originalChildren foundation.NSArray) foundation.NSArray
var _nSAccessibilityUnignoredChildrenErr error

func tryNSAccessibilityUnignoredChildren(originalChildren foundation.NSArray) (foundation.NSArray, error) {
	if _nSAccessibilityUnignoredChildren == nil {
		return foundation.NSArray{}, symbolCallError("NSAccessibilityUnignoredChildren", "", _nSAccessibilityUnignoredChildrenErr)
	}
	return _nSAccessibilityUnignoredChildren(originalChildren), nil
}

// NSAccessibilityUnignoredChildren returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildren(from:)
func NSAccessibilityUnignoredChildren(originalChildren foundation.NSArray) foundation.NSArray {
	result, callErr := tryNSAccessibilityUnignoredChildren(originalChildren)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityUnignoredChildrenForOnlyChild func(originalChild objectivec.Object) foundation.NSArray
var _nSAccessibilityUnignoredChildrenForOnlyChildErr error

func tryNSAccessibilityUnignoredChildrenForOnlyChild(originalChild objectivec.Object) (foundation.NSArray, error) {
	if _nSAccessibilityUnignoredChildrenForOnlyChild == nil {
		return foundation.NSArray{}, symbolCallError("NSAccessibilityUnignoredChildrenForOnlyChild", "", _nSAccessibilityUnignoredChildrenForOnlyChildErr)
	}
	return _nSAccessibilityUnignoredChildrenForOnlyChild(originalChild), nil
}

// NSAccessibilityUnignoredChildrenForOnlyChild returns a list of unignored accessibility objects, descending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredChildrenForOnlyChild(from:)
func NSAccessibilityUnignoredChildrenForOnlyChild(originalChild objectivec.Object) foundation.NSArray {
	result, callErr := tryNSAccessibilityUnignoredChildrenForOnlyChild(originalChild)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAccessibilityUnignoredDescendant func(element objectivec.Object) objectivec.Object
var _nSAccessibilityUnignoredDescendantErr error

func tryNSAccessibilityUnignoredDescendant(element objectivec.Object) (objectivec.Object, error) {
	if _nSAccessibilityUnignoredDescendant == nil {
		return objectivec.Object{}, symbolCallError("NSAccessibilityUnignoredDescendant", "", _nSAccessibilityUnignoredDescendantErr)
	}
	return _nSAccessibilityUnignoredDescendant(element), nil
}

// NSAccessibilityUnignoredDescendant returns an unignored accessibility object, descending the hierarchy, if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/unignoredDescendant(of:)
func NSAccessibilityUnignoredDescendant(element objectivec.Object) objectivec.Object {
	result, callErr := tryNSAccessibilityUnignoredDescendant(element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSApplicationLoad func() bool
var _nSApplicationLoadErr error

func tryNSApplicationLoad() (bool, error) {
	if _nSApplicationLoad == nil {
		return false, symbolCallError("NSApplicationLoad", "", _nSApplicationLoadErr)
	}
	return _nSApplicationLoad(), nil
}

// NSApplicationLoad startup function to call when running Cocoa code from a Carbon application.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationLoad
func NSApplicationLoad() bool {
	result, callErr := tryNSApplicationLoad()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSApplicationMain func(argc int, argv string) int
var _nSApplicationMainErr error

func tryNSApplicationMain(argc int, argv string) (int, error) {
	if _nSApplicationMain == nil {
		return 0, symbolCallError("NSApplicationMain", "", _nSApplicationMainErr)
	}
	return _nSApplicationMain(argc, argv), nil
}

// NSApplicationMain called by the main function to create and run the application.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationMain
func NSApplicationMain(argc int, argv string) int {
	result, callErr := tryNSApplicationMain(argc, argv)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSAvailableWindowDepths func() *NSWindowDepth
var _nSAvailableWindowDepthsErr error

func tryNSAvailableWindowDepths() (*NSWindowDepth, error) {
	if _nSAvailableWindowDepths == nil {
		return nil, symbolCallError("NSAvailableWindowDepths", "", _nSAvailableWindowDepthsErr)
	}
	return _nSAvailableWindowDepths(), nil
}

// NSAvailableWindowDepths returns the available window depth values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAvailableWindowDepths
func NSAvailableWindowDepths() *NSWindowDepth {
	result, callErr := tryNSAvailableWindowDepths()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSBeep func()
var _nSBeepErr error

func tryNSBeep() error {
	if _nSBeep == nil {
		return symbolCallError("NSBeep", "", _nSBeepErr)
	}
	_nSBeep()
	return nil
}

// NSBeep plays the system beep.
//
// See: https://developer.apple.com/documentation/AppKit/NSBeep
func NSBeep() {
	if callErr := tryNSBeep(); callErr != nil {
		panic(callErr)
	}
}

var _nSBestDepth func(colorSpace NSColorSpaceName, bps int, bpp int, planar bool, exactMatch *bool) NSWindowDepth
var _nSBestDepthErr error

func tryNSBestDepth(colorSpace NSColorSpaceName, bps int, bpp int, planar bool, exactMatch *bool) (NSWindowDepth, error) {
	if _nSBestDepth == nil {
		return *new(NSWindowDepth), symbolCallError("NSBestDepth", "", _nSBestDepthErr)
	}
	return _nSBestDepth(colorSpace, bps, bpp, planar, exactMatch), nil
}

// NSBestDepth attempts to return a window depth adequate for the specified parameters.
//
// See: https://developer.apple.com/documentation/AppKit/NSBestDepth
func NSBestDepth(colorSpace NSColorSpaceName, bps int, bpp int, planar bool, exactMatch *bool) NSWindowDepth {
	result, callErr := tryNSBestDepth(colorSpace, bps, bpp, planar, exactMatch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSBitsPerPixelFromDepth func(depth NSWindowDepth) int
var _nSBitsPerPixelFromDepthErr error

func tryNSBitsPerPixelFromDepth(depth NSWindowDepth) (int, error) {
	if _nSBitsPerPixelFromDepth == nil {
		return 0, symbolCallError("NSBitsPerPixelFromDepth", "", _nSBitsPerPixelFromDepthErr)
	}
	return _nSBitsPerPixelFromDepth(depth), nil
}

// NSBitsPerPixelFromDepth returns the bits per pixel for the specified window depth.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerPixel
func NSBitsPerPixelFromDepth(depth NSWindowDepth) int {
	result, callErr := tryNSBitsPerPixelFromDepth(depth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSBitsPerSampleFromDepth func(depth NSWindowDepth) int
var _nSBitsPerSampleFromDepthErr error

func tryNSBitsPerSampleFromDepth(depth NSWindowDepth) (int, error) {
	if _nSBitsPerSampleFromDepth == nil {
		return 0, symbolCallError("NSBitsPerSampleFromDepth", "", _nSBitsPerSampleFromDepthErr)
	}
	return _nSBitsPerSampleFromDepth(depth), nil
}

// NSBitsPerSampleFromDepth returns the bits per sample for the specified window depth.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerSample
func NSBitsPerSampleFromDepth(depth NSWindowDepth) int {
	result, callErr := tryNSBitsPerSampleFromDepth(depth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSColorSpaceFromDepth func(depth NSWindowDepth) NSColorSpaceName
var _nSColorSpaceFromDepthErr error

func tryNSColorSpaceFromDepth(depth NSWindowDepth) (NSColorSpaceName, error) {
	if _nSColorSpaceFromDepth == nil {
		return *new(NSColorSpaceName), symbolCallError("NSColorSpaceFromDepth", "", _nSColorSpaceFromDepthErr)
	}
	return _nSColorSpaceFromDepth(depth), nil
}

// NSColorSpaceFromDepth returns the name of the color space corresponding to the passed window depth.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/colorSpaceName
func NSColorSpaceFromDepth(depth NSWindowDepth) NSColorSpaceName {
	result, callErr := tryNSColorSpaceFromDepth(depth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSConvertGlyphsToPackedGlyphs func(glBuf *NSGlyph, count int, packing NSMultibyteGlyphPacking, packedGlyphs string) int
var _nSConvertGlyphsToPackedGlyphsErr error

func tryNSConvertGlyphsToPackedGlyphs(glBuf *NSGlyph, count int, packing NSMultibyteGlyphPacking, packedGlyphs string) (int, error) {
	if _nSConvertGlyphsToPackedGlyphs == nil {
		return 0, symbolCallError("NSConvertGlyphsToPackedGlyphs", "10.0", _nSConvertGlyphsToPackedGlyphsErr)
	}
	return _nSConvertGlyphsToPackedGlyphs(glBuf, count, packing, packedGlyphs), nil
}

// NSConvertGlyphsToPackedGlyphs prepares a set of glyphs for processing by character-based routines.
//
// Deprecated: Deprecated since macOS 10.13.
//
// See: https://developer.apple.com/documentation/AppKit/NSConvertGlyphsToPackedGlyphs(_:_:_:_:)
func NSConvertGlyphsToPackedGlyphs(glBuf *NSGlyph, count int, packing NSMultibyteGlyphPacking, packedGlyphs string) int {
	result, callErr := tryNSConvertGlyphsToPackedGlyphs(glBuf, count, packing, packedGlyphs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCopyBits func(srcGState int, srcRect corefoundation.CGRect, destPoint corefoundation.CGPoint)
var _nSCopyBitsErr error

func tryNSCopyBits(srcGState int, srcRect corefoundation.CGRect, destPoint corefoundation.CGPoint) error {
	if _nSCopyBits == nil {
		return symbolCallError("NSCopyBits", "10.0", _nSCopyBitsErr)
	}
	_nSCopyBits(srcGState, srcRect, destPoint)
	return nil
}

// NSCopyBits copies a bitmap image to the location specified by a destination point.
//
// Deprecated: Deprecated since macOS 10.10.
//
// See: https://developer.apple.com/documentation/AppKit/NSCopyBits(_:_:_:)
func NSCopyBits(srcGState int, srcRect corefoundation.CGRect, destPoint corefoundation.CGPoint) {
	if callErr := tryNSCopyBits(srcGState, srcRect, destPoint); callErr != nil {
		panic(callErr)
	}
}

var _nSCountWindows func(count *int)
var _nSCountWindowsErr error

func tryNSCountWindows(count *int) error {
	if _nSCountWindows == nil {
		return symbolCallError("NSCountWindows", "10.0", _nSCountWindowsErr)
	}
	_nSCountWindows(count)
	return nil
}

// NSCountWindows counts the number of onscreen windows.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AppKit/NSCountWindows
func NSCountWindows(count *int) {
	if callErr := tryNSCountWindows(count); callErr != nil {
		panic(callErr)
	}
}

var _nSCountWindowsForContext func(context int, count *int)
var _nSCountWindowsForContextErr error

func tryNSCountWindowsForContext(context int, count *int) error {
	if _nSCountWindowsForContext == nil {
		return symbolCallError("NSCountWindowsForContext", "10.0", _nSCountWindowsForContextErr)
	}
	_nSCountWindowsForContext(context, count)
	return nil
}

// NSCountWindowsForContext counts the number of onscreen windows belonging to a particular application.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AppKit/NSCountWindowsForContext
func NSCountWindowsForContext(context int, count *int) {
	if callErr := tryNSCountWindowsForContext(context, count); callErr != nil {
		panic(callErr)
	}
}

var _nSCreateFileContentsPboardType func(fileType foundation.NSString) NSPasteboardType
var _nSCreateFileContentsPboardTypeErr error

func tryNSCreateFileContentsPboardType(fileType foundation.NSString) (NSPasteboardType, error) {
	if _nSCreateFileContentsPboardType == nil {
		return *new(NSPasteboardType), symbolCallError("NSCreateFileContentsPboardType", "", _nSCreateFileContentsPboardTypeErr)
	}
	return _nSCreateFileContentsPboardType(fileType), nil
}

// NSCreateFileContentsPboardType returns a pasteboard type based on the passed file type.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContentsType(forPathExtension:)
func NSCreateFileContentsPboardType(fileType foundation.NSString) NSPasteboardType {
	result, callErr := tryNSCreateFileContentsPboardType(fileType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSCreateFilenamePboardType func(fileType foundation.NSString) NSPasteboardType
var _nSCreateFilenamePboardTypeErr error

func tryNSCreateFilenamePboardType(fileType foundation.NSString) (NSPasteboardType, error) {
	if _nSCreateFilenamePboardType == nil {
		return *new(NSPasteboardType), symbolCallError("NSCreateFilenamePboardType", "", _nSCreateFilenamePboardTypeErr)
	}
	return _nSCreateFilenamePboardType(fileType), nil
}

// NSCreateFilenamePboardType returns a pasteboard type based on the passed file type.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileNameType(forPathExtension:)
func NSCreateFilenamePboardType(fileType foundation.NSString) NSPasteboardType {
	result, callErr := tryNSCreateFilenamePboardType(fileType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDottedFrameRect func(rect corefoundation.CGRect)
var _nSDottedFrameRectErr error

func tryNSDottedFrameRect(rect corefoundation.CGRect) error {
	if _nSDottedFrameRect == nil {
		return symbolCallError("NSDottedFrameRect", "", _nSDottedFrameRectErr)
	}
	_nSDottedFrameRect(rect)
	return nil
}

// NSDottedFrameRect draws a bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSDottedFrameRect(_:)
func NSDottedFrameRect(rect corefoundation.CGRect) {
	if callErr := tryNSDottedFrameRect(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawBitmap func(rect corefoundation.CGRect, width int, height int, bps int, spp int, bpp int, bpr int, isPlanar bool, hasAlpha bool, colorSpaceName NSColorSpaceName, data string)
var _nSDrawBitmapErr error

func tryNSDrawBitmap(rect corefoundation.CGRect, width int, height int, bps int, spp int, bpp int, bpr int, isPlanar bool, hasAlpha bool, colorSpaceName NSColorSpaceName, data string) error {
	if _nSDrawBitmap == nil {
		return symbolCallError("NSDrawBitmap", "", _nSDrawBitmapErr)
	}
	_nSDrawBitmap(rect, width, height, bps, spp, bpp, bpr, isPlanar, hasAlpha, colorSpaceName, data)
	return nil
}

// NSDrawBitmap draws a bitmap image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawBitmap(_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawBitmap(rect corefoundation.CGRect, width int, height int, bps int, spp int, bpp int, bpr int, isPlanar bool, hasAlpha bool, colorSpaceName NSColorSpaceName, data string) {
	if callErr := tryNSDrawBitmap(rect, width, height, bps, spp, bpp, bpr, isPlanar, hasAlpha, colorSpaceName, data); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawButton func(rect corefoundation.CGRect, clipRect corefoundation.CGRect)
var _nSDrawButtonErr error

func tryNSDrawButton(rect corefoundation.CGRect, clipRect corefoundation.CGRect) error {
	if _nSDrawButton == nil {
		return symbolCallError("NSDrawButton", "", _nSDrawButtonErr)
	}
	_nSDrawButton(rect, clipRect)
	return nil
}

// NSDrawButton draws a gray-filled rectangle representing a user-interface button.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawButton(_:_:)
func NSDrawButton(rect corefoundation.CGRect, clipRect corefoundation.CGRect) {
	if callErr := tryNSDrawButton(rect, clipRect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawColorTiledRects func(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, colors *NSColor, count int) corefoundation.CGRect
var _nSDrawColorTiledRectsErr error

func tryNSDrawColorTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, colors *NSColor, count int) (corefoundation.CGRect, error) {
	if _nSDrawColorTiledRects == nil {
		return corefoundation.CGRect{}, symbolCallError("NSDrawColorTiledRects", "", _nSDrawColorTiledRectsErr)
	}
	return _nSDrawColorTiledRects(boundsRect, clipRect, sides, colors, count), nil
}

// NSDrawColorTiledRects draws a single-color, bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawColorTiledRects(_:_:_:_:_:)
func NSDrawColorTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, colors *NSColor, count int) corefoundation.CGRect {
	result, callErr := tryNSDrawColorTiledRects(boundsRect, clipRect, sides, colors, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDrawDarkBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect)
var _nSDrawDarkBezelErr error

func tryNSDrawDarkBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) error {
	if _nSDrawDarkBezel == nil {
		return symbolCallError("NSDrawDarkBezel", "", _nSDrawDarkBezelErr)
	}
	_nSDrawDarkBezel(rect, clipRect)
	return nil
}

// NSDrawDarkBezel draws a dark gray-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawDarkBezel(_:_:)
func NSDrawDarkBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) {
	if callErr := tryNSDrawDarkBezel(rect, clipRect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawGrayBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect)
var _nSDrawGrayBezelErr error

func tryNSDrawGrayBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) error {
	if _nSDrawGrayBezel == nil {
		return symbolCallError("NSDrawGrayBezel", "", _nSDrawGrayBezelErr)
	}
	_nSDrawGrayBezel(rect, clipRect)
	return nil
}

// NSDrawGrayBezel draws a gray-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawGrayBezel(_:_:)
func NSDrawGrayBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) {
	if callErr := tryNSDrawGrayBezel(rect, clipRect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawGroove func(rect corefoundation.CGRect, clipRect corefoundation.CGRect)
var _nSDrawGrooveErr error

func tryNSDrawGroove(rect corefoundation.CGRect, clipRect corefoundation.CGRect) error {
	if _nSDrawGroove == nil {
		return symbolCallError("NSDrawGroove", "", _nSDrawGrooveErr)
	}
	_nSDrawGroove(rect, clipRect)
	return nil
}

// NSDrawGroove draws a gray-filled rectangle with a groove border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawGroove(_:_:)
func NSDrawGroove(rect corefoundation.CGRect, clipRect corefoundation.CGRect) {
	if callErr := tryNSDrawGroove(rect, clipRect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawLightBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect)
var _nSDrawLightBezelErr error

func tryNSDrawLightBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) error {
	if _nSDrawLightBezel == nil {
		return symbolCallError("NSDrawLightBezel", "", _nSDrawLightBezelErr)
	}
	_nSDrawLightBezel(rect, clipRect)
	return nil
}

// NSDrawLightBezel draws a white-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawLightBezel(_:_:)
func NSDrawLightBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) {
	if callErr := tryNSDrawLightBezel(rect, clipRect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawNinePartImage func(frame corefoundation.CGRect, topLeftCorner *NSImage, topEdgeFill *NSImage, topRightCorner *NSImage, leftEdgeFill *NSImage, centerFill *NSImage, rightEdgeFill *NSImage, bottomLeftCorner *NSImage, bottomEdgeFill *NSImage, bottomRightCorner *NSImage, op NSCompositingOperation, alphaFraction float64, flipped bool)
var _nSDrawNinePartImageErr error

func tryNSDrawNinePartImage(frame corefoundation.CGRect, topLeftCorner *NSImage, topEdgeFill *NSImage, topRightCorner *NSImage, leftEdgeFill *NSImage, centerFill *NSImage, rightEdgeFill *NSImage, bottomLeftCorner *NSImage, bottomEdgeFill *NSImage, bottomRightCorner *NSImage, op NSCompositingOperation, alphaFraction float64, flipped bool) error {
	if _nSDrawNinePartImage == nil {
		return symbolCallError("NSDrawNinePartImage", "10.5", _nSDrawNinePartImageErr)
	}
	_nSDrawNinePartImage(frame, topLeftCorner, topEdgeFill, topRightCorner, leftEdgeFill, centerFill, rightEdgeFill, bottomLeftCorner, bottomEdgeFill, bottomRightCorner, op, alphaFraction, flipped)
	return nil
}

// NSDrawNinePartImage draws a nine-part tiled image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawNinePartImage(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawNinePartImage(frame corefoundation.CGRect, topLeftCorner *NSImage, topEdgeFill *NSImage, topRightCorner *NSImage, leftEdgeFill *NSImage, centerFill *NSImage, rightEdgeFill *NSImage, bottomLeftCorner *NSImage, bottomEdgeFill *NSImage, bottomRightCorner *NSImage, op NSCompositingOperation, alphaFraction float64, flipped bool) {
	if callErr := tryNSDrawNinePartImage(frame, topLeftCorner, topEdgeFill, topRightCorner, leftEdgeFill, centerFill, rightEdgeFill, bottomLeftCorner, bottomEdgeFill, bottomRightCorner, op, alphaFraction, flipped); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawThreePartImage func(frame corefoundation.CGRect, startCap *NSImage, centerFill *NSImage, endCap *NSImage, vertical bool, op NSCompositingOperation, alphaFraction float64, flipped bool)
var _nSDrawThreePartImageErr error

func tryNSDrawThreePartImage(frame corefoundation.CGRect, startCap *NSImage, centerFill *NSImage, endCap *NSImage, vertical bool, op NSCompositingOperation, alphaFraction float64, flipped bool) error {
	if _nSDrawThreePartImage == nil {
		return symbolCallError("NSDrawThreePartImage", "10.5", _nSDrawThreePartImageErr)
	}
	_nSDrawThreePartImage(frame, startCap, centerFill, endCap, vertical, op, alphaFraction, flipped)
	return nil
}

// NSDrawThreePartImage draws a three-part tiled image.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawThreePartImage(_:_:_:_:_:_:_:_:)
func NSDrawThreePartImage(frame corefoundation.CGRect, startCap *NSImage, centerFill *NSImage, endCap *NSImage, vertical bool, op NSCompositingOperation, alphaFraction float64, flipped bool) {
	if callErr := tryNSDrawThreePartImage(frame, startCap, centerFill, endCap, vertical, op, alphaFraction, flipped); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawTiledRects func(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, grays *float64, count int) corefoundation.CGRect
var _nSDrawTiledRectsErr error

func tryNSDrawTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, grays *float64, count int) (corefoundation.CGRect, error) {
	if _nSDrawTiledRects == nil {
		return corefoundation.CGRect{}, symbolCallError("NSDrawTiledRects", "", _nSDrawTiledRectsErr)
	}
	return _nSDrawTiledRects(boundsRect, clipRect, sides, grays, count), nil
}

// NSDrawTiledRects draws rectangles with borders.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawTiledRects(_:_:_:_:_:)
func NSDrawTiledRects(boundsRect corefoundation.CGRect, clipRect corefoundation.CGRect, sides *foundation.NSRectEdge, grays *float64, count int) corefoundation.CGRect {
	result, callErr := tryNSDrawTiledRects(boundsRect, clipRect, sides, grays, count)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSDrawWhiteBezel func(rect corefoundation.CGRect, clipRect corefoundation.CGRect)
var _nSDrawWhiteBezelErr error

func tryNSDrawWhiteBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) error {
	if _nSDrawWhiteBezel == nil {
		return symbolCallError("NSDrawWhiteBezel", "", _nSDrawWhiteBezelErr)
	}
	_nSDrawWhiteBezel(rect, clipRect)
	return nil
}

// NSDrawWhiteBezel draws a white-filled rectangle with a bezel border.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawWhiteBezel(_:_:)
func NSDrawWhiteBezel(rect corefoundation.CGRect, clipRect corefoundation.CGRect) {
	if callErr := tryNSDrawWhiteBezel(rect, clipRect); callErr != nil {
		panic(callErr)
	}
}

var _nSDrawWindowBackground func(rect corefoundation.CGRect)
var _nSDrawWindowBackgroundErr error

func tryNSDrawWindowBackground(rect corefoundation.CGRect) error {
	if _nSDrawWindowBackground == nil {
		return symbolCallError("NSDrawWindowBackground", "", _nSDrawWindowBackgroundErr)
	}
	_nSDrawWindowBackground(rect)
	return nil
}

// NSDrawWindowBackground draws the window’s default background pattern into the specified rectangle of the currently focused view.
//
// See: https://developer.apple.com/documentation/AppKit/NSDrawWindowBackground(_:)
func NSDrawWindowBackground(rect corefoundation.CGRect) {
	if callErr := tryNSDrawWindowBackground(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSEraseRect func(rect corefoundation.CGRect)
var _nSEraseRectErr error

func tryNSEraseRect(rect corefoundation.CGRect) error {
	if _nSEraseRect == nil {
		return symbolCallError("NSEraseRect", "", _nSEraseRectErr)
	}
	_nSEraseRect(rect)
	return nil
}

// NSEraseRect erases the specified rect by filling it with white.
//
// See: https://developer.apple.com/documentation/AppKit/NSEraseRect(_:)
func NSEraseRect(rect corefoundation.CGRect) {
	if callErr := tryNSEraseRect(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSFrameRect func(rect corefoundation.CGRect)
var _nSFrameRectErr error

func tryNSFrameRect(rect corefoundation.CGRect) error {
	if _nSFrameRect == nil {
		return symbolCallError("NSFrameRect", "", _nSFrameRectErr)
	}
	_nSFrameRect(rect)
	return nil
}

// NSFrameRect draws a bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSFrameRect
func NSFrameRect(rect corefoundation.CGRect) {
	if callErr := tryNSFrameRect(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSFrameRectWithWidth func(rect corefoundation.CGRect, frameWidth float64)
var _nSFrameRectWithWidthErr error

func tryNSFrameRectWithWidth(rect corefoundation.CGRect, frameWidth float64) error {
	if _nSFrameRectWithWidth == nil {
		return symbolCallError("NSFrameRectWithWidth", "", _nSFrameRectWithWidthErr)
	}
	_nSFrameRectWithWidth(rect, frameWidth)
	return nil
}

// NSFrameRectWithWidth draws a bordered rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidth
func NSFrameRectWithWidth(rect corefoundation.CGRect, frameWidth float64) {
	if callErr := tryNSFrameRectWithWidth(rect, frameWidth); callErr != nil {
		panic(callErr)
	}
}

var _nSFrameRectWithWidthUsingOperation func(rect corefoundation.CGRect, frameWidth float64, op NSCompositingOperation)
var _nSFrameRectWithWidthUsingOperationErr error

func tryNSFrameRectWithWidthUsingOperation(rect corefoundation.CGRect, frameWidth float64, op NSCompositingOperation) error {
	if _nSFrameRectWithWidthUsingOperation == nil {
		return symbolCallError("NSFrameRectWithWidthUsingOperation", "", _nSFrameRectWithWidthUsingOperationErr)
	}
	_nSFrameRectWithWidthUsingOperation(rect, frameWidth, op)
	return nil
}

// NSFrameRectWithWidthUsingOperation draws a bordered rectangle using the specified compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidthUsingOperation
func NSFrameRectWithWidthUsingOperation(rect corefoundation.CGRect, frameWidth float64, op NSCompositingOperation) {
	if callErr := tryNSFrameRectWithWidthUsingOperation(rect, frameWidth, op); callErr != nil {
		panic(callErr)
	}
}

var _nSGetFileType func(pboardType NSPasteboardType) foundation.NSString
var _nSGetFileTypeErr error

func tryNSGetFileType(pboardType NSPasteboardType) (foundation.NSString, error) {
	if _nSGetFileType == nil {
		return foundation.NSString{}, symbolCallError("NSGetFileType", "", _nSGetFileTypeErr)
	}
	return _nSGetFileType(pboardType), nil
}

// NSGetFileType a file type based on the passed pasteboard type.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtension
func NSGetFileType(pboardType NSPasteboardType) foundation.NSString {
	result, callErr := tryNSGetFileType(pboardType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSGetFileTypes func(pboardTypes *foundation.NSString) []foundation.NSString
var _nSGetFileTypesErr error

func tryNSGetFileTypes(pboardTypes []foundation.NSString) ([]foundation.NSString, error) {
	if _nSGetFileTypes == nil {
		return nil, symbolCallError("NSGetFileTypes", "", _nSGetFileTypesErr)
	}
	return _nSGetFileTypes(unsafe.SliceData(pboardTypes)), nil
}

// NSGetFileTypes returns an array of file types based on the passed pasteboard types.
//
// Deprecated: The file contents pboard type allowed you to synthesize a pboard type for a file’s contents based on the file’s extension. Using the UTI of a file to represent its contents now replaces this functionality.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtensions(from:)
func NSGetFileTypes(pboardTypes []foundation.NSString) []foundation.NSString {
	result, callErr := tryNSGetFileTypes(pboardTypes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSHighlightRect func(rect corefoundation.CGRect)
var _nSHighlightRectErr error

func tryNSHighlightRect(rect corefoundation.CGRect) error {
	if _nSHighlightRect == nil {
		return symbolCallError("NSHighlightRect", "10.0", _nSHighlightRectErr)
	}
	_nSHighlightRect(rect)
	return nil
}

// NSHighlightRect highlights the specified rect by filling it with white.
//
// Deprecated: Deprecated since macOS 10.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSHighlightRect
func NSHighlightRect(rect corefoundation.CGRect) {
	if callErr := tryNSHighlightRect(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSInterfaceStyleForKey func(key foundation.NSString, responder *NSResponder) unsafe.Pointer
var _nSInterfaceStyleForKeyErr error

func tryNSInterfaceStyleForKey(key foundation.NSString, responder *NSResponder) (unsafe.Pointer, error) {
	if _nSInterfaceStyleForKey == nil {
		return nil, symbolCallError("NSInterfaceStyleForKey", "10.0", _nSInterfaceStyleForKeyErr)
	}
	return _nSInterfaceStyleForKey(key, responder), nil
}

// NSInterfaceStyleForKey returns an interface style value for the specified key and responder.
//
// Deprecated: Deprecated since macOS 10.8.
//
// See: https://developer.apple.com/documentation/AppKit/NSInterfaceStyleForKey
func NSInterfaceStyleForKey(key foundation.NSString, responder *NSResponder) unsafe.Pointer {
	result, callErr := tryNSInterfaceStyleForKey(key, responder)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSIsControllerMarker func(object objectivec.Object) bool
var _nSIsControllerMarkerErr error

func tryNSIsControllerMarker(object objectivec.Object) (bool, error) {
	if _nSIsControllerMarker == nil {
		return false, symbolCallError("NSIsControllerMarker", "", _nSIsControllerMarkerErr)
	}
	return _nSIsControllerMarker(object), nil
}

// NSIsControllerMarker tests whether a given object is special marker object used for indicating the state of a selection in relation to a key.
//
// See: https://developer.apple.com/documentation/AppKit/NSIsControllerMarker(_:)
func NSIsControllerMarker(object objectivec.Object) bool {
	result, callErr := tryNSIsControllerMarker(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSNumberOfColorComponents func(colorSpaceName NSColorSpaceName) int
var _nSNumberOfColorComponentsErr error

func tryNSNumberOfColorComponents(colorSpaceName NSColorSpaceName) (int, error) {
	if _nSNumberOfColorComponents == nil {
		return 0, symbolCallError("NSNumberOfColorComponents", "", _nSNumberOfColorComponentsErr)
	}
	return _nSNumberOfColorComponents(colorSpaceName), nil
}

// NSNumberOfColorComponents returns the number of color components in the specified color space.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/numberOfColorComponents
func NSNumberOfColorComponents(colorSpaceName NSColorSpaceName) int {
	result, callErr := tryNSNumberOfColorComponents(colorSpaceName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSPerformService func(itemName foundation.NSString, pboard *NSPasteboard) bool
var _nSPerformServiceErr error

func tryNSPerformService(itemName foundation.NSString, pboard *NSPasteboard) (bool, error) {
	if _nSPerformService == nil {
		return false, symbolCallError("NSPerformService", "", _nSPerformServiceErr)
	}
	return _nSPerformService(itemName, pboard), nil
}

// NSPerformService programmatically invokes a Services menu service.
//
// See: https://developer.apple.com/documentation/AppKit/NSPerformService(_:_:)
func NSPerformService(itemName foundation.NSString, pboard *NSPasteboard) bool {
	result, callErr := tryNSPerformService(itemName, pboard)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSPlanarFromDepth func(depth NSWindowDepth) bool
var _nSPlanarFromDepthErr error

func tryNSPlanarFromDepth(depth NSWindowDepth) (bool, error) {
	if _nSPlanarFromDepth == nil {
		return false, symbolCallError("NSPlanarFromDepth", "", _nSPlanarFromDepthErr)
	}
	return _nSPlanarFromDepth(depth), nil
}

// NSPlanarFromDepth returns whether the specified window depth is planar.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/isPlanar
func NSPlanarFromDepth(depth NSWindowDepth) bool {
	result, callErr := tryNSPlanarFromDepth(depth)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRectClip func(rect corefoundation.CGRect)
var _nSRectClipErr error

func tryNSRectClip(rect corefoundation.CGRect) error {
	if _nSRectClip == nil {
		return symbolCallError("NSRectClip", "", _nSRectClipErr)
	}
	_nSRectClip(rect)
	return nil
}

// NSRectClip modifies the current clipping path by intersecting it with the passed rect.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectClip
func NSRectClip(rect corefoundation.CGRect) {
	if callErr := tryNSRectClip(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSRectClipList func(rects *corefoundation.CGRect, count int)
var _nSRectClipListErr error

func tryNSRectClipList(rects *corefoundation.CGRect, count int) error {
	if _nSRectClipList == nil {
		return symbolCallError("NSRectClipList", "", _nSRectClipListErr)
	}
	_nSRectClipList(rects, count)
	return nil
}

// NSRectClipList modifies the current clipping path by intersecting it with the passed rect.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectClipList
func NSRectClipList(rects *corefoundation.CGRect, count int) {
	if callErr := tryNSRectClipList(rects, count); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFill func(rect corefoundation.CGRect)
var _nSRectFillErr error

func tryNSRectFill(rect corefoundation.CGRect) error {
	if _nSRectFill == nil {
		return symbolCallError("NSRectFill", "", _nSRectFillErr)
	}
	_nSRectFill(rect)
	return nil
}

// NSRectFill fills the passed rectangle with the current color.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFill
func NSRectFill(rect corefoundation.CGRect) {
	if callErr := tryNSRectFill(rect); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFillList func(rects *corefoundation.CGRect, count int)
var _nSRectFillListErr error

func tryNSRectFillList(rects *corefoundation.CGRect, count int) error {
	if _nSRectFillList == nil {
		return symbolCallError("NSRectFillList", "", _nSRectFillListErr)
	}
	_nSRectFillList(rects, count)
	return nil
}

// NSRectFillList fills the rectangles in the passed list with the current fill color.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillList
func NSRectFillList(rects *corefoundation.CGRect, count int) {
	if callErr := tryNSRectFillList(rects, count); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFillListUsingOperation func(rects *corefoundation.CGRect, count int, op NSCompositingOperation)
var _nSRectFillListUsingOperationErr error

func tryNSRectFillListUsingOperation(rects *corefoundation.CGRect, count int, op NSCompositingOperation) error {
	if _nSRectFillListUsingOperation == nil {
		return symbolCallError("NSRectFillListUsingOperation", "", _nSRectFillListUsingOperationErr)
	}
	_nSRectFillListUsingOperation(rects, count, op)
	return nil
}

// NSRectFillListUsingOperation fills the rectangles in a list using the current fill color and specified compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListUsingOperation
func NSRectFillListUsingOperation(rects *corefoundation.CGRect, count int, op NSCompositingOperation) {
	if callErr := tryNSRectFillListUsingOperation(rects, count, op); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFillListWithColors func(rects *corefoundation.CGRect, colors NSColor, num int)
var _nSRectFillListWithColorsErr error

func tryNSRectFillListWithColors(rects *corefoundation.CGRect, colors NSColor, num int) error {
	if _nSRectFillListWithColors == nil {
		return symbolCallError("NSRectFillListWithColors", "", _nSRectFillListWithColorsErr)
	}
	_nSRectFillListWithColors(rects, colors, num)
	return nil
}

// NSRectFillListWithColors fills the rectangles in the passed list with the passed list of colors.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColors
func NSRectFillListWithColors(rects *corefoundation.CGRect, colors NSColor, num int) {
	if callErr := tryNSRectFillListWithColors(rects, colors, num); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFillListWithColorsUsingOperation func(rects *corefoundation.CGRect, colors NSColor, num int, op NSCompositingOperation)
var _nSRectFillListWithColorsUsingOperationErr error

func tryNSRectFillListWithColorsUsingOperation(rects *corefoundation.CGRect, colors NSColor, num int, op NSCompositingOperation) error {
	if _nSRectFillListWithColorsUsingOperation == nil {
		return symbolCallError("NSRectFillListWithColorsUsingOperation", "", _nSRectFillListWithColorsUsingOperationErr)
	}
	_nSRectFillListWithColorsUsingOperation(rects, colors, num, op)
	return nil
}

// NSRectFillListWithColorsUsingOperation fills the rectangles in a list using the specified colors and compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColorsUsingOperation
func NSRectFillListWithColorsUsingOperation(rects *corefoundation.CGRect, colors NSColor, num int, op NSCompositingOperation) {
	if callErr := tryNSRectFillListWithColorsUsingOperation(rects, colors, num, op); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFillListWithGrays func(rects *corefoundation.CGRect, grays *float64, num int)
var _nSRectFillListWithGraysErr error

func tryNSRectFillListWithGrays(rects *corefoundation.CGRect, grays *float64, num int) error {
	if _nSRectFillListWithGrays == nil {
		return symbolCallError("NSRectFillListWithGrays", "", _nSRectFillListWithGraysErr)
	}
	_nSRectFillListWithGrays(rects, grays, num)
	return nil
}

// NSRectFillListWithGrays fills the rectangles in the passed list with the passed list of grays.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillListWithGrays
func NSRectFillListWithGrays(rects *corefoundation.CGRect, grays *float64, num int) {
	if callErr := tryNSRectFillListWithGrays(rects, grays, num); callErr != nil {
		panic(callErr)
	}
}

var _nSRectFillUsingOperation func(rect corefoundation.CGRect, op NSCompositingOperation)
var _nSRectFillUsingOperationErr error

func tryNSRectFillUsingOperation(rect corefoundation.CGRect, op NSCompositingOperation) error {
	if _nSRectFillUsingOperation == nil {
		return symbolCallError("NSRectFillUsingOperation", "", _nSRectFillUsingOperationErr)
	}
	_nSRectFillUsingOperation(rect, op)
	return nil
}

// NSRectFillUsingOperation fills a rectangle using the current fill color and the specified compositing operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSRectFillUsingOperation
func NSRectFillUsingOperation(rect corefoundation.CGRect, op NSCompositingOperation) {
	if callErr := tryNSRectFillUsingOperation(rect, op); callErr != nil {
		panic(callErr)
	}
}

var _nSRegisterServicesProvider func(provider objectivec.Object, name NSServiceProviderName)
var _nSRegisterServicesProviderErr error

func tryNSRegisterServicesProvider(provider objectivec.Object, name NSServiceProviderName) error {
	if _nSRegisterServicesProvider == nil {
		return symbolCallError("NSRegisterServicesProvider", "", _nSRegisterServicesProviderErr)
	}
	_nSRegisterServicesProvider(provider, name)
	return nil
}

// NSRegisterServicesProvider registers a service provider.
//
// See: https://developer.apple.com/documentation/AppKit/NSRegisterServicesProvider(_:_:)
func NSRegisterServicesProvider(provider objectivec.Object, name NSServiceProviderName) {
	if callErr := tryNSRegisterServicesProvider(provider, name); callErr != nil {
		panic(callErr)
	}
}

var _nSRunAlertPanelRelativeToWindow func(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int
var _nSRunAlertPanelRelativeToWindowErr error

func tryNSRunAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) (int, error) {
	if _nSRunAlertPanelRelativeToWindow == nil {
		return 0, symbolCallError("NSRunAlertPanelRelativeToWindow", "10.0", _nSRunAlertPanelRelativeToWindowErr)
	}
	return _nSRunAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow), nil
}

// NSRunAlertPanelRelativeToWindow.
//
// Deprecated: Deprecated since macOS 10.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunAlertPanelRelativeToWindow
func NSRunAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int {
	result, callErr := tryNSRunAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRunCriticalAlertPanelRelativeToWindow func(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int
var _nSRunCriticalAlertPanelRelativeToWindowErr error

func tryNSRunCriticalAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) (int, error) {
	if _nSRunCriticalAlertPanelRelativeToWindow == nil {
		return 0, symbolCallError("NSRunCriticalAlertPanelRelativeToWindow", "10.0", _nSRunCriticalAlertPanelRelativeToWindowErr)
	}
	return _nSRunCriticalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow), nil
}

// NSRunCriticalAlertPanelRelativeToWindow.
//
// Deprecated: Deprecated since macOS 10.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunCriticalAlertPanelRelativeToWindow
func NSRunCriticalAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int {
	result, callErr := tryNSRunCriticalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSRunInformationalAlertPanelRelativeToWindow func(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int
var _nSRunInformationalAlertPanelRelativeToWindowErr error

func tryNSRunInformationalAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) (int, error) {
	if _nSRunInformationalAlertPanelRelativeToWindow == nil {
		return 0, symbolCallError("NSRunInformationalAlertPanelRelativeToWindow", "10.0", _nSRunInformationalAlertPanelRelativeToWindowErr)
	}
	return _nSRunInformationalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow), nil
}

// NSRunInformationalAlertPanelRelativeToWindow.
//
// Deprecated: Deprecated since macOS 10.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunInformationalAlertPanelRelativeToWindow
func NSRunInformationalAlertPanelRelativeToWindow(title foundation.NSString, msgFormat foundation.NSString, defaultButton foundation.NSString, alternateButton foundation.NSString, otherButton foundation.NSString, docWindow *NSWindow) int {
	result, callErr := tryNSRunInformationalAlertPanelRelativeToWindow(title, msgFormat, defaultButton, alternateButton, otherButton, docWindow)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSSetFocusRingStyle func(placement NSFocusRingPlacement)
var _nSSetFocusRingStyleErr error

func tryNSSetFocusRingStyle(placement NSFocusRingPlacement) error {
	if _nSSetFocusRingStyle == nil {
		return symbolCallError("NSSetFocusRingStyle", "", _nSSetFocusRingStyleErr)
	}
	_nSSetFocusRingStyle(placement)
	return nil
}

// NSSetFocusRingStyle specifies how the system draws the focus ring.
//
// See: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/set()
func NSSetFocusRingStyle(placement NSFocusRingPlacement) {
	if callErr := tryNSSetFocusRingStyle(placement); callErr != nil {
		panic(callErr)
	}
}

var _nSSetShowsServicesMenuItem func(itemName foundation.NSString, enabled bool) int
var _nSSetShowsServicesMenuItemErr error

func tryNSSetShowsServicesMenuItem(itemName foundation.NSString, enabled bool) (int, error) {
	if _nSSetShowsServicesMenuItem == nil {
		return 0, symbolCallError("NSSetShowsServicesMenuItem", "", _nSSetShowsServicesMenuItemErr)
	}
	return _nSSetShowsServicesMenuItem(itemName, enabled), nil
}

// NSSetShowsServicesMenuItem specifies whether an item should be included in Services menus.
//
// See: https://developer.apple.com/documentation/AppKit/NSSetShowsServicesMenuItem(_:_:)
func NSSetShowsServicesMenuItem(itemName foundation.NSString, enabled bool) int {
	result, callErr := tryNSSetShowsServicesMenuItem(itemName, enabled)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSShowAnimationEffect func(animationEffect NSAnimationEffect, centerLocation corefoundation.CGPoint, size corefoundation.CGSize, animationDelegate objectivec.Object, didEndSelector objectivec.SEL, contextInfo uintptr)
var _nSShowAnimationEffectErr error

func tryNSShowAnimationEffect(animationEffect NSAnimationEffect, centerLocation corefoundation.CGPoint, size corefoundation.CGSize, animationDelegate objectivec.Object, didEndSelector objectivec.SEL, contextInfo uintptr) error {
	if _nSShowAnimationEffect == nil {
		return symbolCallError("NSShowAnimationEffect", "10.3", _nSShowAnimationEffectErr)
	}
	_nSShowAnimationEffect(animationEffect, centerLocation, size, animationDelegate, didEndSelector, contextInfo)
	return nil
}

// NSShowAnimationEffect runs a system animation effect.
//
// Deprecated: Deprecated since macOS 14.0. Use [disappearingItem](<doc://com.apple.appkit/documentation/AppKit/NSCursor/disappearingItem>) instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSShowAnimationEffect
func NSShowAnimationEffect(animationEffect NSAnimationEffect, centerLocation corefoundation.CGPoint, size corefoundation.CGSize, animationDelegate objectivec.Object, didEndSelector objectivec.SEL, contextInfo uintptr) {
	if callErr := tryNSShowAnimationEffect(animationEffect, centerLocation, size, animationDelegate, didEndSelector, contextInfo); callErr != nil {
		panic(callErr)
	}
}

var _nSShowsServicesMenuItem func(itemName foundation.NSString) bool
var _nSShowsServicesMenuItemErr error

func tryNSShowsServicesMenuItem(itemName foundation.NSString) (bool, error) {
	if _nSShowsServicesMenuItem == nil {
		return false, symbolCallError("NSShowsServicesMenuItem", "", _nSShowsServicesMenuItemErr)
	}
	return _nSShowsServicesMenuItem(itemName), nil
}

// NSShowsServicesMenuItem specifies whether a Services menu item is currently enabled.
//
// See: https://developer.apple.com/documentation/AppKit/NSShowsServicesMenuItem(_:)
func NSShowsServicesMenuItem(itemName foundation.NSString) bool {
	result, callErr := tryNSShowsServicesMenuItem(itemName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _nSUnregisterServicesProvider func(name NSServiceProviderName)
var _nSUnregisterServicesProviderErr error

func tryNSUnregisterServicesProvider(name NSServiceProviderName) error {
	if _nSUnregisterServicesProvider == nil {
		return symbolCallError("NSUnregisterServicesProvider", "", _nSUnregisterServicesProviderErr)
	}
	_nSUnregisterServicesProvider(name)
	return nil
}

// NSUnregisterServicesProvider unregisters a service provider.
//
// See: https://developer.apple.com/documentation/AppKit/NSUnregisterServicesProvider(_:)
func NSUnregisterServicesProvider(name NSServiceProviderName) {
	if callErr := tryNSUnregisterServicesProvider(name); callErr != nil {
		panic(callErr)
	}
}

var _nSUpdateDynamicServices func()
var _nSUpdateDynamicServicesErr error

func tryNSUpdateDynamicServices() error {
	if _nSUpdateDynamicServices == nil {
		return symbolCallError("NSUpdateDynamicServices", "", _nSUpdateDynamicServicesErr)
	}
	_nSUpdateDynamicServices()
	return nil
}

// NSUpdateDynamicServices causes the services information for the system to be updated.
//
// See: https://developer.apple.com/documentation/AppKit/NSUpdateDynamicServices()
func NSUpdateDynamicServices() {
	if callErr := tryNSUpdateDynamicServices(); callErr != nil {
		panic(callErr)
	}
}

var _nSWindowList func(size int, list int)
var _nSWindowListErr error

func tryNSWindowList(size int, list int) error {
	if _nSWindowList == nil {
		return symbolCallError("NSWindowList", "10.0", _nSWindowListErr)
	}
	_nSWindowList(size, list)
	return nil
}

// NSWindowList gets information about onscreen windows.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowList
func NSWindowList(size int, list int) {
	if callErr := tryNSWindowList(size, list); callErr != nil {
		panic(callErr)
	}
}

var _nSWindowListForContext func(context int, size int, list int)
var _nSWindowListForContextErr error

func tryNSWindowListForContext(context int, size int, list int) error {
	if _nSWindowListForContext == nil {
		return symbolCallError("NSWindowListForContext", "10.0", _nSWindowListForContextErr)
	}
	_nSWindowListForContext(context, size, list)
	return nil
}

// NSWindowListForContext gets information about an application’s onscreen windows.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowListForContext
func NSWindowListForContext(context int, size int, list int) {
	if callErr := tryNSWindowListForContext(context, size, list); callErr != nil {
		panic(callErr)
	}
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_nSAccessibilityActionDescription, &_nSAccessibilityActionDescriptionErr, frameworkHandle, "NSAccessibilityActionDescription", "")
	registerFunc(&_nSAccessibilityFrameInView, &_nSAccessibilityFrameInViewErr, frameworkHandle, "NSAccessibilityFrameInView", "10.10")
	registerFunc(&_nSAccessibilityPointInView, &_nSAccessibilityPointInViewErr, frameworkHandle, "NSAccessibilityPointInView", "10.10")
	registerFunc(&_nSAccessibilityPostNotification, &_nSAccessibilityPostNotificationErr, frameworkHandle, "NSAccessibilityPostNotification", "")
	registerFunc(&_nSAccessibilityPostNotificationWithUserInfo, &_nSAccessibilityPostNotificationWithUserInfoErr, frameworkHandle, "NSAccessibilityPostNotificationWithUserInfo", "10.7")
	registerFunc(&_nSAccessibilityRoleDescription, &_nSAccessibilityRoleDescriptionErr, frameworkHandle, "NSAccessibilityRoleDescription", "")
	registerFunc(&_nSAccessibilityRoleDescriptionForUIElement, &_nSAccessibilityRoleDescriptionForUIElementErr, frameworkHandle, "NSAccessibilityRoleDescriptionForUIElement", "")
	registerFunc(&_nSAccessibilitySetMayContainProtectedContent, &_nSAccessibilitySetMayContainProtectedContentErr, frameworkHandle, "NSAccessibilitySetMayContainProtectedContent", "")
	registerFunc(&_nSAccessibilityUnignoredAncestor, &_nSAccessibilityUnignoredAncestorErr, frameworkHandle, "NSAccessibilityUnignoredAncestor", "")
	registerFunc(&_nSAccessibilityUnignoredChildren, &_nSAccessibilityUnignoredChildrenErr, frameworkHandle, "NSAccessibilityUnignoredChildren", "")
	registerFunc(&_nSAccessibilityUnignoredChildrenForOnlyChild, &_nSAccessibilityUnignoredChildrenForOnlyChildErr, frameworkHandle, "NSAccessibilityUnignoredChildrenForOnlyChild", "")
	registerFunc(&_nSAccessibilityUnignoredDescendant, &_nSAccessibilityUnignoredDescendantErr, frameworkHandle, "NSAccessibilityUnignoredDescendant", "")
	registerFunc(&_nSApplicationLoad, &_nSApplicationLoadErr, frameworkHandle, "NSApplicationLoad", "")
	registerFunc(&_nSApplicationMain, &_nSApplicationMainErr, frameworkHandle, "NSApplicationMain", "")
	registerFunc(&_nSAvailableWindowDepths, &_nSAvailableWindowDepthsErr, frameworkHandle, "NSAvailableWindowDepths", "")
	registerFunc(&_nSBeep, &_nSBeepErr, frameworkHandle, "NSBeep", "")
	registerFunc(&_nSBestDepth, &_nSBestDepthErr, frameworkHandle, "NSBestDepth", "")
	registerFunc(&_nSBitsPerPixelFromDepth, &_nSBitsPerPixelFromDepthErr, frameworkHandle, "NSBitsPerPixelFromDepth", "")
	registerFunc(&_nSBitsPerSampleFromDepth, &_nSBitsPerSampleFromDepthErr, frameworkHandle, "NSBitsPerSampleFromDepth", "")
	registerFunc(&_nSColorSpaceFromDepth, &_nSColorSpaceFromDepthErr, frameworkHandle, "NSColorSpaceFromDepth", "")
	registerFunc(&_nSConvertGlyphsToPackedGlyphs, &_nSConvertGlyphsToPackedGlyphsErr, frameworkHandle, "NSConvertGlyphsToPackedGlyphs", "10.0")
	registerFunc(&_nSCopyBits, &_nSCopyBitsErr, frameworkHandle, "NSCopyBits", "10.0")
	registerFunc(&_nSCountWindows, &_nSCountWindowsErr, frameworkHandle, "NSCountWindows", "10.0")
	registerFunc(&_nSCountWindowsForContext, &_nSCountWindowsForContextErr, frameworkHandle, "NSCountWindowsForContext", "10.0")
	registerFunc(&_nSCreateFileContentsPboardType, &_nSCreateFileContentsPboardTypeErr, frameworkHandle, "NSCreateFileContentsPboardType", "")
	registerFunc(&_nSCreateFilenamePboardType, &_nSCreateFilenamePboardTypeErr, frameworkHandle, "NSCreateFilenamePboardType", "")
	registerFunc(&_nSDottedFrameRect, &_nSDottedFrameRectErr, frameworkHandle, "NSDottedFrameRect", "")
	registerFunc(&_nSDrawBitmap, &_nSDrawBitmapErr, frameworkHandle, "NSDrawBitmap", "")
	registerFunc(&_nSDrawButton, &_nSDrawButtonErr, frameworkHandle, "NSDrawButton", "")
	registerFunc(&_nSDrawColorTiledRects, &_nSDrawColorTiledRectsErr, frameworkHandle, "NSDrawColorTiledRects", "")
	registerFunc(&_nSDrawDarkBezel, &_nSDrawDarkBezelErr, frameworkHandle, "NSDrawDarkBezel", "")
	registerFunc(&_nSDrawGrayBezel, &_nSDrawGrayBezelErr, frameworkHandle, "NSDrawGrayBezel", "")
	registerFunc(&_nSDrawGroove, &_nSDrawGrooveErr, frameworkHandle, "NSDrawGroove", "")
	registerFunc(&_nSDrawLightBezel, &_nSDrawLightBezelErr, frameworkHandle, "NSDrawLightBezel", "")
	registerFunc(&_nSDrawNinePartImage, &_nSDrawNinePartImageErr, frameworkHandle, "NSDrawNinePartImage", "10.5")
	registerFunc(&_nSDrawThreePartImage, &_nSDrawThreePartImageErr, frameworkHandle, "NSDrawThreePartImage", "10.5")
	registerFunc(&_nSDrawTiledRects, &_nSDrawTiledRectsErr, frameworkHandle, "NSDrawTiledRects", "")
	registerFunc(&_nSDrawWhiteBezel, &_nSDrawWhiteBezelErr, frameworkHandle, "NSDrawWhiteBezel", "")
	registerFunc(&_nSDrawWindowBackground, &_nSDrawWindowBackgroundErr, frameworkHandle, "NSDrawWindowBackground", "")
	registerFunc(&_nSEraseRect, &_nSEraseRectErr, frameworkHandle, "NSEraseRect", "")
	registerFunc(&_nSFrameRect, &_nSFrameRectErr, frameworkHandle, "NSFrameRect", "")
	registerFunc(&_nSFrameRectWithWidth, &_nSFrameRectWithWidthErr, frameworkHandle, "NSFrameRectWithWidth", "")
	registerFunc(&_nSFrameRectWithWidthUsingOperation, &_nSFrameRectWithWidthUsingOperationErr, frameworkHandle, "NSFrameRectWithWidthUsingOperation", "")
	registerFunc(&_nSGetFileType, &_nSGetFileTypeErr, frameworkHandle, "NSGetFileType", "")
	registerFunc(&_nSGetFileTypes, &_nSGetFileTypesErr, frameworkHandle, "NSGetFileTypes", "")
	registerFunc(&_nSHighlightRect, &_nSHighlightRectErr, frameworkHandle, "NSHighlightRect", "10.0")
	registerFunc(&_nSInterfaceStyleForKey, &_nSInterfaceStyleForKeyErr, frameworkHandle, "NSInterfaceStyleForKey", "10.0")
	registerFunc(&_nSIsControllerMarker, &_nSIsControllerMarkerErr, frameworkHandle, "NSIsControllerMarker", "")
	registerFunc(&_nSNumberOfColorComponents, &_nSNumberOfColorComponentsErr, frameworkHandle, "NSNumberOfColorComponents", "")
	registerFunc(&_nSPerformService, &_nSPerformServiceErr, frameworkHandle, "NSPerformService", "")
	registerFunc(&_nSPlanarFromDepth, &_nSPlanarFromDepthErr, frameworkHandle, "NSPlanarFromDepth", "")
	registerFunc(&_nSRectClip, &_nSRectClipErr, frameworkHandle, "NSRectClip", "")
	registerFunc(&_nSRectClipList, &_nSRectClipListErr, frameworkHandle, "NSRectClipList", "")
	registerFunc(&_nSRectFill, &_nSRectFillErr, frameworkHandle, "NSRectFill", "")
	registerFunc(&_nSRectFillList, &_nSRectFillListErr, frameworkHandle, "NSRectFillList", "")
	registerFunc(&_nSRectFillListUsingOperation, &_nSRectFillListUsingOperationErr, frameworkHandle, "NSRectFillListUsingOperation", "")
	registerFunc(&_nSRectFillListWithColors, &_nSRectFillListWithColorsErr, frameworkHandle, "NSRectFillListWithColors", "")
	registerFunc(&_nSRectFillListWithColorsUsingOperation, &_nSRectFillListWithColorsUsingOperationErr, frameworkHandle, "NSRectFillListWithColorsUsingOperation", "")
	registerFunc(&_nSRectFillListWithGrays, &_nSRectFillListWithGraysErr, frameworkHandle, "NSRectFillListWithGrays", "")
	registerFunc(&_nSRectFillUsingOperation, &_nSRectFillUsingOperationErr, frameworkHandle, "NSRectFillUsingOperation", "")
	registerFunc(&_nSRegisterServicesProvider, &_nSRegisterServicesProviderErr, frameworkHandle, "NSRegisterServicesProvider", "")
	registerFunc(&_nSRunAlertPanelRelativeToWindow, &_nSRunAlertPanelRelativeToWindowErr, frameworkHandle, "NSRunAlertPanelRelativeToWindow", "10.0")
	registerFunc(&_nSRunCriticalAlertPanelRelativeToWindow, &_nSRunCriticalAlertPanelRelativeToWindowErr, frameworkHandle, "NSRunCriticalAlertPanelRelativeToWindow", "10.0")
	registerFunc(&_nSRunInformationalAlertPanelRelativeToWindow, &_nSRunInformationalAlertPanelRelativeToWindowErr, frameworkHandle, "NSRunInformationalAlertPanelRelativeToWindow", "10.0")
	registerFunc(&_nSSetFocusRingStyle, &_nSSetFocusRingStyleErr, frameworkHandle, "NSSetFocusRingStyle", "")
	registerFunc(&_nSSetShowsServicesMenuItem, &_nSSetShowsServicesMenuItemErr, frameworkHandle, "NSSetShowsServicesMenuItem", "")
	registerFunc(&_nSShowAnimationEffect, &_nSShowAnimationEffectErr, frameworkHandle, "NSShowAnimationEffect", "10.3")
	registerFunc(&_nSShowsServicesMenuItem, &_nSShowsServicesMenuItemErr, frameworkHandle, "NSShowsServicesMenuItem", "")
	registerFunc(&_nSUnregisterServicesProvider, &_nSUnregisterServicesProviderErr, frameworkHandle, "NSUnregisterServicesProvider", "")
	registerFunc(&_nSUpdateDynamicServices, &_nSUpdateDynamicServicesErr, frameworkHandle, "NSUpdateDynamicServices", "")
	registerFunc(&_nSWindowList, &_nSWindowListErr, frameworkHandle, "NSWindowList", "10.0")
	registerFunc(&_nSWindowListForContext, &_nSWindowListForContextErr, frameworkHandle, "NSWindowListForContext", "10.0")
}
