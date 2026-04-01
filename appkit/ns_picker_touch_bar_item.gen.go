// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPickerTouchBarItem] class.
var (
	_NSPickerTouchBarItemClass     NSPickerTouchBarItemClass
	_NSPickerTouchBarItemClassOnce sync.Once
)

func getNSPickerTouchBarItemClass() NSPickerTouchBarItemClass {
	_NSPickerTouchBarItemClassOnce.Do(func() {
		_NSPickerTouchBarItemClass = NSPickerTouchBarItemClass{class: objc.GetClass("NSPickerTouchBarItem")}
	})
	return _NSPickerTouchBarItemClass
}

// GetNSPickerTouchBarItemClass returns the class object for NSPickerTouchBarItem.
func GetNSPickerTouchBarItemClass() NSPickerTouchBarItemClass {
	return getNSPickerTouchBarItemClass()
}

type NSPickerTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPickerTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPickerTouchBarItemClass) Alloc() NSPickerTouchBarItem {
	rv := objc.Send[NSPickerTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that provides a picker control with multiple options.
//
// # Configuring picker appearance
//
//   - [NSPickerTouchBarItem.NumberOfOptions]
//   - [NSPickerTouchBarItem.SetNumberOfOptions]
//   - [NSPickerTouchBarItem.SetLabelAtIndex]
//   - [NSPickerTouchBarItem.LabelAtIndex]
//   - [NSPickerTouchBarItem.SetImageAtIndex]
//   - [NSPickerTouchBarItem.ImageAtIndex]
//   - [NSPickerTouchBarItem.CollapsedRepresentationImage]
//   - [NSPickerTouchBarItem.SetCollapsedRepresentationImage]
//   - [NSPickerTouchBarItem.CollapsedRepresentationLabel]
//   - [NSPickerTouchBarItem.SetCollapsedRepresentationLabel]
//   - [NSPickerTouchBarItem.ControlRepresentation]
//   - [NSPickerTouchBarItem.SetControlRepresentation]
//
// # Configuring picker state
//
//   - [NSPickerTouchBarItem.Enabled]
//   - [NSPickerTouchBarItem.SetEnabled]
//   - [NSPickerTouchBarItem.IsEnabledAtIndex]
//   - [NSPickerTouchBarItem.SetEnabledAtIndex]
//
// # Handling selection
//
//   - [NSPickerTouchBarItem.SelectedIndex]
//   - [NSPickerTouchBarItem.SetSelectedIndex]
//   - [NSPickerTouchBarItem.SelectionColor]
//   - [NSPickerTouchBarItem.SetSelectionColor]
//   - [NSPickerTouchBarItem.SelectionMode]
//   - [NSPickerTouchBarItem.SetSelectionMode]
//
// # Handling picker interaction
//
//   - [NSPickerTouchBarItem.Action]
//   - [NSPickerTouchBarItem.SetAction]
//   - [NSPickerTouchBarItem.Target]
//   - [NSPickerTouchBarItem.SetTarget]
//
// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem
type NSPickerTouchBarItem struct {
	NSTouchBarItem
}

// NSPickerTouchBarItemFromID constructs a [NSPickerTouchBarItem] from an objc.ID.
//
// A bar item that provides a picker control with multiple options.
func NSPickerTouchBarItemFromID(id objc.ID) NSPickerTouchBarItem {
	return NSPickerTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}

// NOTE: NSPickerTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPickerTouchBarItem] class.
//
// # Configuring picker appearance
//
//   - [INSPickerTouchBarItem.NumberOfOptions]
//   - [INSPickerTouchBarItem.SetNumberOfOptions]
//   - [INSPickerTouchBarItem.SetLabelAtIndex]
//   - [INSPickerTouchBarItem.LabelAtIndex]
//   - [INSPickerTouchBarItem.SetImageAtIndex]
//   - [INSPickerTouchBarItem.ImageAtIndex]
//   - [INSPickerTouchBarItem.CollapsedRepresentationImage]
//   - [INSPickerTouchBarItem.SetCollapsedRepresentationImage]
//   - [INSPickerTouchBarItem.CollapsedRepresentationLabel]
//   - [INSPickerTouchBarItem.SetCollapsedRepresentationLabel]
//   - [INSPickerTouchBarItem.ControlRepresentation]
//   - [INSPickerTouchBarItem.SetControlRepresentation]
//
// # Configuring picker state
//
//   - [INSPickerTouchBarItem.Enabled]
//   - [INSPickerTouchBarItem.SetEnabled]
//   - [INSPickerTouchBarItem.IsEnabledAtIndex]
//   - [INSPickerTouchBarItem.SetEnabledAtIndex]
//
// # Handling selection
//
//   - [INSPickerTouchBarItem.SelectedIndex]
//   - [INSPickerTouchBarItem.SetSelectedIndex]
//   - [INSPickerTouchBarItem.SelectionColor]
//   - [INSPickerTouchBarItem.SetSelectionColor]
//   - [INSPickerTouchBarItem.SelectionMode]
//   - [INSPickerTouchBarItem.SetSelectionMode]
//
// # Handling picker interaction
//
//   - [INSPickerTouchBarItem.Action]
//   - [INSPickerTouchBarItem.SetAction]
//   - [INSPickerTouchBarItem.Target]
//   - [INSPickerTouchBarItem.SetTarget]
//
// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem
type INSPickerTouchBarItem interface {
	INSTouchBarItem

	// Topic: Configuring picker appearance

	NumberOfOptions() int
	SetNumberOfOptions(value int)
	SetLabelAtIndex(label string, index int)
	LabelAtIndex(index int) string
	SetImageAtIndex(image INSImage, index int)
	ImageAtIndex(index int) INSImage
	CollapsedRepresentationImage() INSImage
	SetCollapsedRepresentationImage(value INSImage)
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)
	ControlRepresentation() NSPickerTouchBarItemControlRepresentation
	SetControlRepresentation(value NSPickerTouchBarItemControlRepresentation)

	// Topic: Configuring picker state

	Enabled() bool
	SetEnabled(value bool)
	IsEnabledAtIndex(index int) bool
	SetEnabledAtIndex(enabled bool, index int)

	// Topic: Handling selection

	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectionColor() INSColor
	SetSelectionColor(value INSColor)
	SelectionMode() NSPickerTouchBarItemSelectionMode
	SetSelectionMode(value NSPickerTouchBarItemSelectionMode)

	// Topic: Handling picker interaction

	Action() objc.SEL
	SetAction(value objc.SEL)
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
}

// Init initializes the instance.
func (p NSPickerTouchBarItem) Init() NSPickerTouchBarItem {
	rv := objc.Send[NSPickerTouchBarItem](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPickerTouchBarItem) Autorelease() NSPickerTouchBarItem {
	rv := objc.Send[NSPickerTouchBarItem](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPickerTouchBarItem creates a new NSPickerTouchBarItem instance.
func NewNSPickerTouchBarItem() NSPickerTouchBarItem {
	class := getNSPickerTouchBarItemClass()
	rv := objc.Send[NSPickerTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewPickerTouchBarItemWithCoder(coder foundation.INSCoder) NSPickerTouchBarItem {
	instance := getNSPickerTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPickerTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
//
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewPickerTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSPickerTouchBarItem {
	instance := getNSPickerTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSPickerTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:images:selectionMode:target:action:)
func NewPickerTouchBarItemWithIdentifierImagesSelectionModeTargetAction(identifier NSTouchBarItemIdentifier, images []NSImage, selectionMode NSPickerTouchBarItemSelectionMode, target objectivec.IObject, action objc.SEL) NSPickerTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSPickerTouchBarItemClass().class), objc.Sel("pickerTouchBarItemWithIdentifier:images:selectionMode:target:action:"), objc.String(string(identifier)), objectivec.IObjectSliceToNSArray(images), selectionMode, target, action)
	return NSPickerTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func NewPickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier NSTouchBarItemIdentifier, labels []string, selectionMode NSPickerTouchBarItemSelectionMode, target objectivec.IObject, action objc.SEL) NSPickerTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(getNSPickerTouchBarItemClass().class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), objc.String(string(identifier)), objectivec.StringSliceToNSArray(labels), selectionMode, target, action)
	return NSPickerTouchBarItemFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/setLabel(_:at:)
func (p NSPickerTouchBarItem) SetLabelAtIndex(label string, index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setLabel:atIndex:"), objc.String(label), index)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/label(at:)
func (p NSPickerTouchBarItem) LabelAtIndex(index int) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("labelAtIndex:"), index)
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/setImage(_:at:)
func (p NSPickerTouchBarItem) SetImageAtIndex(image INSImage, index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setImage:atIndex:"), image, index)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/image(at:)
func (p NSPickerTouchBarItem) ImageAtIndex(index int) INSImage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("imageAtIndex:"), index)
	return NSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/isEnabled(at:)
func (p NSPickerTouchBarItem) IsEnabledAtIndex(index int) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isEnabledAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/setEnabled(_:at:)
func (p NSPickerTouchBarItem) SetEnabledAtIndex(enabled bool, index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setEnabled:atIndex:"), enabled, index)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/numberOfOptions
func (p NSPickerTouchBarItem) NumberOfOptions() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfOptions"))
	return rv
}
func (p NSPickerTouchBarItem) SetNumberOfOptions(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNumberOfOptions:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p NSPickerTouchBarItem) CollapsedRepresentationImage() INSImage {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("collapsedRepresentationImage"))
	return NSImageFromID(objc.ID(rv))
}
func (p NSPickerTouchBarItem) SetCollapsedRepresentationImage(value INSImage) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p NSPickerTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("collapsedRepresentationLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (p NSPickerTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setCollapsedRepresentationLabel:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/controlRepresentation-swift.property
func (p NSPickerTouchBarItem) ControlRepresentation() NSPickerTouchBarItemControlRepresentation {
	rv := objc.Send[NSPickerTouchBarItemControlRepresentation](p.ID, objc.Sel("controlRepresentation"))
	return NSPickerTouchBarItemControlRepresentation(rv)
}
func (p NSPickerTouchBarItem) SetControlRepresentation(value NSPickerTouchBarItemControlRepresentation) {
	objc.Send[struct{}](p.ID, objc.Sel("setControlRepresentation:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/isEnabled
func (p NSPickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isEnabled"))
	return rv
}
func (p NSPickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setEnabled:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectedIndex
func (p NSPickerTouchBarItem) SelectedIndex() int {
	rv := objc.Send[int](p.ID, objc.Sel("selectedIndex"))
	return rv
}
func (p NSPickerTouchBarItem) SetSelectedIndex(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectedIndex:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectionColor
func (p NSPickerTouchBarItem) SelectionColor() INSColor {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectionColor"))
	return NSColorFromID(objc.ID(rv))
}
func (p NSPickerTouchBarItem) SetSelectionColor(value INSColor) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectionColor:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectionMode-swift.property
func (p NSPickerTouchBarItem) SelectionMode() NSPickerTouchBarItemSelectionMode {
	rv := objc.Send[NSPickerTouchBarItemSelectionMode](p.ID, objc.Sel("selectionMode"))
	return NSPickerTouchBarItemSelectionMode(rv)
}
func (p NSPickerTouchBarItem) SetSelectionMode(value NSPickerTouchBarItemSelectionMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectionMode:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/action
func (p NSPickerTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](p.ID, objc.Sel("action"))
	return rv
}
func (p NSPickerTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[struct{}](p.ID, objc.Sel("setAction:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/target
func (p NSPickerTouchBarItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (p NSPickerTouchBarItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setTarget:"), value)
}
