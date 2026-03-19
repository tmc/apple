// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAppearance] class.
var (
	_NSAppearanceClass     NSAppearanceClass
	_NSAppearanceClassOnce sync.Once
)

func getNSAppearanceClass() NSAppearanceClass {
	_NSAppearanceClassOnce.Do(func() {
		_NSAppearanceClass = NSAppearanceClass{class: objc.GetClass("NSAppearance")}
	})
	return _NSAppearanceClass
}

// GetNSAppearanceClass returns the class object for NSAppearance.
func GetNSAppearanceClass() NSAppearanceClass {
	return getNSAppearanceClass()
}

type NSAppearanceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAppearanceClass) Alloc() NSAppearance {
	rv := objc.Send[NSAppearance](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages standard appearance attributes for UI elements in an
// app.
//
// # Overview
// 
// An [NSAppearance] object manages how AppKit renders your app’s UI
// elements. Specifically, appearance objects determine which colors and
// images AppKit uses when drawing windows, views, and controls. Although you
// can use an appearance object to determine how to draw custom views and
// controls, a better approach is to choose colors and images that adapt
// automatically to the current appearance. For example, define a color asset
// whose actual color value changes for light and dark appearances. You can
// assign specific appearances to your views in Interface Builder.
// 
// The user chooses the default appearance for the system, but you can
// override that appearance for all or part of your app. Apps inherit the
// default system appearance, windows inherit their app’s appearance, and
// views inherit the appearance of their nearest ancestor (either a superview
// or window). To force a window or view to adopt an appearance, assign a
// specific appearance object to its [Appearance] property.
// 
// When AppKit draws a control, it automatically sets the current appearance
// on the current thread to the control’s appearance. The current appearance
// influences the drawing path and return values you get when you access
// system fonts and colors. The current appearance also affects the appearance
// of text and images, such as the text and template images in a toolbar.
//
// # Creating an Appearance
//
//   - [NSAppearance.InitWithAppearanceNamedBundle]: Creates an appearance object from the named appearance file located in the specified bundle.
//   - [NSAppearance.InitWithCoder]
//
// # Getting the Appearance Name
//
//   - [NSAppearance.Name]: The name of the appearance.
//
// # Determining the Most Appropriate Appearance
//
//   - [NSAppearance.BestMatchFromAppearancesWithNames]: Returns the appearance name that most closely matches the current appearance object.
//
// # Getting and Setting the Current Appearance
//
//   - [NSAppearance.PerformAsCurrentDrawingAppearance]: Sets the appearance to be the active drawing appearance and perform the specified block.
//
// # Managing Vibrancy
//
//   - [NSAppearance.AllowsVibrancy]: Specifies whether the current appearance allows vibrancy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance
type NSAppearance struct {
	objectivec.Object
}

// NSAppearanceFromID constructs a [NSAppearance] from an objc.ID.
//
// An object that manages standard appearance attributes for UI elements in an
// app.
func NSAppearanceFromID(id objc.ID) NSAppearance {
	return NSAppearance{objectivec.Object{ID: id}}
}
// NOTE: NSAppearance adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAppearance] class.
//
// # Creating an Appearance
//
//   - [INSAppearance.InitWithAppearanceNamedBundle]: Creates an appearance object from the named appearance file located in the specified bundle.
//   - [INSAppearance.InitWithCoder]
//
// # Getting the Appearance Name
//
//   - [INSAppearance.Name]: The name of the appearance.
//
// # Determining the Most Appropriate Appearance
//
//   - [INSAppearance.BestMatchFromAppearancesWithNames]: Returns the appearance name that most closely matches the current appearance object.
//
// # Getting and Setting the Current Appearance
//
//   - [INSAppearance.PerformAsCurrentDrawingAppearance]: Sets the appearance to be the active drawing appearance and perform the specified block.
//
// # Managing Vibrancy
//
//   - [INSAppearance.AllowsVibrancy]: Specifies whether the current appearance allows vibrancy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance
type INSAppearance interface {
	objectivec.IObject

	// Topic: Creating an Appearance

	// Creates an appearance object from the named appearance file located in the specified bundle.
	InitWithAppearanceNamedBundle(name NSAppearanceName, bundle foundation.NSBundle) NSAppearance
	InitWithCoder(coder foundation.INSCoder) NSAppearance

	// Topic: Getting the Appearance Name

	// The name of the appearance.
	Name() NSAppearanceName

	// Topic: Determining the Most Appropriate Appearance

	// Returns the appearance name that most closely matches the current appearance object.
	BestMatchFromAppearancesWithNames(appearances []string) NSAppearanceName

	// Topic: Getting and Setting the Current Appearance

	// Sets the appearance to be the active drawing appearance and perform the specified block.
	PerformAsCurrentDrawingAppearance(block VoidHandler)

	// Topic: Managing Vibrancy

	// Specifies whether the current appearance allows vibrancy.
	AllowsVibrancy() bool

	// The appearance of the receiver, in an 
	Appearance() INSAppearance
	SetAppearance(value INSAppearance)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a NSAppearance) Init() NSAppearance {
	rv := objc.Send[NSAppearance](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAppearance) Autorelease() NSAppearance {
	rv := objc.Send[NSAppearance](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAppearance creates a new NSAppearance instance.
func NewNSAppearance() NSAppearance {
	class := getNSAppearanceClass()
	rv := objc.Send[NSAppearance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an appearance object based on the name of one of the standard
// system appearances.
//
// name: The name of a standard appearance. See [NSAppearanceName] for the list of
// standard appearance names.
//
// # Return Value
// 
// A standard [NSAppearance] object.
//
// # Discussion
// 
// When you specify a standard appearance name—such as [aqua]—this method
// returns a built-in appearance.
//
// [aqua]: https://developer.apple.com/documentation/AppKit/NSAppearance/Name-swift.struct/aqua
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/init(named:)
func NewAppearanceNamed(name NSAppearanceName) NSAppearance {
	rv := objc.Send[objc.ID](objc.ID(getNSAppearanceClass().class), objc.Sel("appearanceNamed:"), objc.String(string(name)))
	return NSAppearanceFromID(rv)
}

// Creates an appearance object from the named appearance file located in the
// specified bundle.
//
// name: The name of the appearance file to retrieve. Do not include any path
// information in the name.
//
// bundle: The bundle in which to search for the named appearance file. Specify `nil`
// to search for the appearance file in the main bundle.
//
// # Return Value
// 
// An initialized appearance object, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/init(appearanceNamed:bundle:)
func NewAppearanceWithAppearanceNamedBundle(name NSAppearanceName, bundle foundation.NSBundle) NSAppearance {
	instance := getNSAppearanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAppearanceNamed:bundle:"), objc.String(string(name)), bundle)
	return NSAppearanceFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/init(coder:)
func NewAppearanceWithCoder(coder foundation.INSCoder) NSAppearance {
	instance := getNSAppearanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSAppearanceFromID(rv)
}

// Creates an appearance object from the named appearance file located in the
// specified bundle.
//
// name: The name of the appearance file to retrieve. Do not include any path
// information in the name.
//
// bundle: The bundle in which to search for the named appearance file. Specify `nil`
// to search for the appearance file in the main bundle.
//
// # Return Value
// 
// An initialized appearance object, or `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/init(appearanceNamed:bundle:)
func (a NSAppearance) InitWithAppearanceNamedBundle(name NSAppearanceName, bundle foundation.NSBundle) NSAppearance {
	rv := objc.Send[NSAppearance](a.ID, objc.Sel("initWithAppearanceNamed:bundle:"), objc.String(string(name)), bundle)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/init(coder:)
func (a NSAppearance) InitWithCoder(coder foundation.INSCoder) NSAppearance {
	rv := objc.Send[NSAppearance](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Returns the appearance name that most closely matches the current
// appearance object.
//
// appearances: An array of appearance names, representing the appearances that your app
// supports.
//
// # Return Value
// 
// The name of the appearance that most closely matches the current appearance
// object.
//
// # Discussion
// 
// You can use this method in situations where your app doesn’t fully
// support the current appearance, but supports a different appearance object
// that has similar qualities. This method returns the name from the
// `appearances` array that comes closest to matching the current appearance
// object’s attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/bestMatch(from:)
func (a NSAppearance) BestMatchFromAppearancesWithNames(appearances []string) NSAppearanceName {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bestMatchFromAppearancesWithNames:"), objectivec.StringSliceToNSArray(appearances))
	return NSAppearanceName(foundation.NSStringFromID(rv).String())
}
// Sets the appearance to be the active drawing appearance and perform the
// specified block.
//
// block: The block to invoke after setting the appearance to be the current drawing
// appearance.
//
// # Discussion
// 
// This method saves and restores the previous current appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/performAsCurrentDrawingAppearance(_:)
func (a NSAppearance) PerformAsCurrentDrawingAppearance(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](a.ID, objc.Sel("performAsCurrentDrawingAppearance:"), _block0)
}
func (a NSAppearance) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The appearance that the system uses for color and asset resolution, and
// that’s active for drawing, usually from locking focus on a view.
//
// # Return Value
// 
// The current appearance used for drawing.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/currentDrawing()
func (_NSAppearanceClass NSAppearanceClass) CurrentDrawingAppearance() NSAppearance {
	rv := objc.Send[objc.ID](objc.ID(_NSAppearanceClass.class), objc.Sel("currentDrawingAppearance"))
	return NSAppearanceFromID(rv)
}

// The name of the appearance.
//
// # Discussion
// 
// For a list of standard appearance names, see [NSAppearanceName].
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/name-swift.property
func (a NSAppearance) Name() NSAppearanceName {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return NSAppearanceName(foundation.NSStringFromID(rv).String())
}
// Specifies whether the current appearance allows vibrancy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearance/allowsVibrancy
func (a NSAppearance) AllowsVibrancy() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowsVibrancy"))
	return rv
}
// The appearance of the receiver, in an
//
// See: https://developer.apple.com/documentation/appkit/nsappearancecustomization/appearance
func (a NSAppearance) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (a NSAppearance) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](a.ID, objc.Sel("setAppearance:"), value)
}

// PerformAsCurrentDrawingAppearanceSync is a synchronous wrapper around [NSAppearance.PerformAsCurrentDrawingAppearance].
// It blocks until the completion handler fires or the context is cancelled.
func (a NSAppearance) PerformAsCurrentDrawingAppearanceSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.PerformAsCurrentDrawingAppearance(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

