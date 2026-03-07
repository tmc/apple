
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

// Package appkit provides Go bindings for the AppKit framework.
//
// Construct and manage a graphical, event-driven user interface for your macOS app.
//
// AppKit contains the objects you need to build the user interface for a macOS app. In addition to drawing windows, buttons, panels, and text fields, it handles all the event management and interaction between your app, people, and macOS.
//
// # Essentials
//
//   - Adopting Liquid Glass: Find out how to bring the new material to your app.
//   - AppKit updates: Learn about important changes to AppKit.
//   - Protecting the User’s Privacy: Secure personal data, and respect user preferences for how data is used.
//   - Porting your macOS apps to Apple silicon: Create a version of your macOS app that runs on both Apple silicon and Intel-based Mac computers.
//
// # App Structure
//
//   - App and Environment: Learn about the objects that you use to interact with the system. ([NSApplication], [NSRunningApplication], [NSApplicationDelegate], [NSWorkspace], [NSAppKitVersion])
//   - Documents, Data, and Pasteboard: Organize your app’s data and preferences, and share that data on the pasteboard or in iCloud. ([NSDocument], [NSDocumentController], [NSPersistentDocument], [NSUserDefaultsController], [NSPasteboard])
//   - Cocoa Bindings: Automatically synchronize your data model with your app’s interface using Cocoa Bindings. ([NSObjectController], [NSController], [NSTreeController], [NSTreeNode], [NSArrayController])
//   - Resource Management: Manage the storyboards and nib files containing your app’s user interface, and learn how to load data that is stored in resource files. ([NSStoryboard], [NSStoryboardSegue], [NSSeguePerforming], [NSDataAsset], [NSNib])
//   - App Extensions: Extend your app’s basic functionality to other parts of the system.
//
// # User Interface
//
//   - Views and Controls: Present your content onscreen and handle user input and events. ([NSView], [NSControl], [NSCell], [NSActionCell], [NSSplitView])
//   - View Management: Manage your user interface, including the size and position of views in a window. ([NSWindowController], [NSViewController], [NSTitlebarAccessoryViewController], [NSSplitViewController], [NSSplitView])
//   - View Layout: Position and size views using a stack view or Auto Layout constraints. ([NSStackView], [NSLayoutConstraint], [NSLayoutGuide], [NSLayoutDimension], [NSLayoutAnchor])
//   - Appearance Customization: Add Dark Mode support to your app, and use appearance proxies to modify your UI. ([NSAppearance], [NSAppearanceCustomization])
//   - Animation: Animate your views and other content to create a more engaging experience for users. ([NSViewAnimation], [NSAnimatablePropertyContainer], [NSAnimationContext], [NSAnimationEffect], [NSViewControllerPresentationAnimator])
//   - Windows, Panels, and Screens: Organize your view hierarchies and facilitate their display onscreen. ([NSWindow], [NSPanel], [NSWindowDelegate], [NSWindowTab], [NSWindowTabGroup])
//   - Sound, Speech, and Haptics: Play sounds and haptic feedback, and incorporate speech recognition and synthesis into your interface. ([NSSound], [NSSpeechRecognizer], [NSSpeechSynthesizer], [NSHapticFeedbackManager], [NSHapticFeedbackPerformer])
//   - Supporting Continuity Camera in Your Mac App: Incorporate scanned documents and pictures from a user’s iPhone, iPad, or iPod touch into your Mac app using Continuity Camera.
//
// # User Interactions
//
//   - Mouse, Keyboard, and Trackpad: Handle events related to mouse, keyboard, and trackpad input. ([NSResponder], [NSEvent], [NSTouch], [NSPressureConfiguration], [NSHapticFeedbackManager])
//   - Menus, Cursors, and the Dock: Implement menus and cursors to facilitate interactions with your app, and use your app’s Dock tile to convey updated information. ([NSMenu], [NSMenuItem], [NSMenuItemBadge], [NSMenuDelegate], [NSMenuItemValidation])
//   - Gestures: Encapsulate your app’s event-handling logic in gesture recognizers so that you can reuse that code throughout your app. ([NSClickGestureRecognizer], [NSPressGestureRecognizer], [NSPanGestureRecognizer], [NSRotationGestureRecognizer], [NSMagnificationGestureRecognizer])
//   - Touch Bar: Display interactive content and controls in the Touch Bar. ([NSTouchBar], [NSTouchBarDelegate], [NSTouchBarProvider], [NSTouchBarItem], [NSCandidateListTouchBarItem])
//   - Drag and Drop: Support the direct manipulation of your app’s content using drag and drop. ([NSDraggingSource], [NSDraggingItem], [NSDraggingSession], [NSDraggingImageComponent], [NSDraggingDestination])
//   - Accessibility for AppKit: Make your AppKit apps accessible to everyone who uses macOS. ([NSAccessibilityProtocol], [NSAccessibility], [NSAccessibilityElement], [NSAccessibilityAnnotationPosition], [NSAccessibilityOrientation])
//
// # Graphics, Drawing, Color, and Printing
//
//   - Images and PDF: Create and manage images, in bitmap, PDF, and other formats. ([NSImage], [NSImageDelegate], [NSImageRep], [NSBitmapImageRep], [NSCIImageRep])
//   - Drawing: Draw shapes, images, and other content on the screen. ([NSGraphicsContext], [NSBezierPath], [NSStringDrawingContext], [NSStringDrawingOptions], [NSGradient])
//   - Color: Represent colors using built-in or custom formats, and give users options for selecting and applying colors. ([NSColor], [NSColorList], [NSColorSpace], [NSColorPicker], [NSColorWell])
//   - Printing: Display the system print panels and manage the printing process. ([NSPrintPanel], [NSPageLayout], [NSPrinter], [NSPrintInfo], [NSPrintOperation])
//
// # Text
//
//   - Text Display: Display text and check spelling. ([NSTextField], [NSTextFieldDelegate], [NSTextView], [NSTextViewDelegate], [NSTextDelegate])
//   - TextKit: Manage text storage and perform custom layout of text-based content in your app’s views. ([NSTextContentStorage], [NSTextContentManager], [NSParagraphStyle], [NSMutableParagraphStyle], [NSTextTab])
//   - Fonts: Manage the fonts used to display text. ([NSFont], [NSFontDescriptor], [NSFontTraitMask], [NSFontFamilyClass], [NSFontAssetRequest])
//   - Writing Tools: Add support for Writing Tools to your app’s text views. ([NSWritingToolsBehavior], [NSWritingToolsResultOptions], [NSWritingToolsCoordinator], [NSTextPreview])
//
// # Deprecated
//
//   - Deprecated Symbols: Review symbols that are no longer supported, and find the replacements to use instead. ([NSFormCell], [NSMenuItemCell], [NSEditorRegistration], [NSInputServiceProvider], [NSInputServerMouseTracker])
//
// # Key Types
//
//   - [NSView] - The infrastructure for drawing, printing, and handling events in an app.
//   - [NSWindow] - A window that an app displays on the screen.
//   - [NSTextView] - A view that draws text and handles user interactions with that text.
//   - [NSTableView] - A set of related records, displayed in rows that represent individual records and columns that represent the attributes of those records.
//   - [NSApplication] - An object that manages an app’s main event loop and resources used by all of that app’s objects.
//   - [NSBrowser] - An interface that displays a hierarchically organized list of data items that can be navigated and selected.
//   - [NSOutlineView] - A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed.
//   - [NSMatrix] - A legacy interface for grouping radio buttons or other types of cells together.
//   - [NSCollectionView] - An ordered collection of data items displayed in a customizable layout.
//   - [NSScrollView] - A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.
//
// [AppKit Documentation]: https://developer.apple.com/documentation/AppKit
package appkit

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AppKit.framework/AppKit"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: AppKit: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

