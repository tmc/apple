// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDraggingImageComponent] class.
var (
	_NSDraggingImageComponentClass     NSDraggingImageComponentClass
	_NSDraggingImageComponentClassOnce sync.Once
)

func getNSDraggingImageComponentClass() NSDraggingImageComponentClass {
	_NSDraggingImageComponentClassOnce.Do(func() {
		_NSDraggingImageComponentClass = NSDraggingImageComponentClass{class: objc.GetClass("NSDraggingImageComponent")}
	})
	return _NSDraggingImageComponentClass
}

// GetNSDraggingImageComponentClass returns the class object for NSDraggingImageComponent.
func GetNSDraggingImageComponentClass() NSDraggingImageComponentClass {
	return getNSDraggingImageComponentClass()
}

type NSDraggingImageComponentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDraggingImageComponentClass) Alloc() NSDraggingImageComponent {
	rv := objc.Send[NSDraggingImageComponent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A single object in a dragging item.
//
// # Overview
// 
// An array of [NSDraggingImageComponent] instances are composited together to
// create the dragging image for an [NSDraggingItem].
// [NSDraggingImageComponent] instances can simply be considered as named
// images with a location used by an [NSDraggingItem] instance.
//
// # Creating a Dragging Image Component
//
//   - [NSDraggingImageComponent.InitWithKey]: Initializes and returns a dragging image component with the specified key.
//
// # Dragging Image Component
//
//   - [NSDraggingImageComponent.Key]: The unique name of this image component instance.
//   - [NSDraggingImageComponent.SetKey]
//
// # Dragging Image Contents
//
//   - [NSDraggingImageComponent.Contents]: An object providing the image contents of the component.
//   - [NSDraggingImageComponent.SetContents]
//   - [NSDraggingImageComponent.Frame]: The coordinate space is the bounds of the parent dragging item.
//   - [NSDraggingImageComponent.SetFrame]
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent
type NSDraggingImageComponent struct {
	objectivec.Object
}

// NSDraggingImageComponentFromID constructs a [NSDraggingImageComponent] from an objc.ID.
//
// A single object in a dragging item.
func NSDraggingImageComponentFromID(id objc.ID) NSDraggingImageComponent {
	return NSDraggingImageComponent{objectivec.Object{ID: id}}
}
// NOTE: NSDraggingImageComponent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDraggingImageComponent] class.
//
// # Creating a Dragging Image Component
//
//   - [INSDraggingImageComponent.InitWithKey]: Initializes and returns a dragging image component with the specified key.
//
// # Dragging Image Component
//
//   - [INSDraggingImageComponent.Key]: The unique name of this image component instance.
//   - [INSDraggingImageComponent.SetKey]
//
// # Dragging Image Contents
//
//   - [INSDraggingImageComponent.Contents]: An object providing the image contents of the component.
//   - [INSDraggingImageComponent.SetContents]
//   - [INSDraggingImageComponent.Frame]: The coordinate space is the bounds of the parent dragging item.
//   - [INSDraggingImageComponent.SetFrame]
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent
type INSDraggingImageComponent interface {
	objectivec.IObject

	// Topic: Creating a Dragging Image Component

	// Initializes and returns a dragging image component with the specified key.
	InitWithKey(key NSDraggingImageComponentKey) NSDraggingImageComponent

	// Topic: Dragging Image Component

	// The unique name of this image component instance.
	Key() NSDraggingImageComponentKey
	SetKey(value NSDraggingImageComponentKey)

	// Topic: Dragging Image Contents

	// An object providing the image contents of the component.
	Contents() objectivec.IObject
	SetContents(value objectivec.IObject)
	// The coordinate space is the bounds of the parent dragging item.
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
}

// Init initializes the instance.
func (d NSDraggingImageComponent) Init() NSDraggingImageComponent {
	rv := objc.Send[NSDraggingImageComponent](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDraggingImageComponent) Autorelease() NSDraggingImageComponent {
	rv := objc.Send[NSDraggingImageComponent](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDraggingImageComponent creates a new NSDraggingImageComponent instance.
func NewNSDraggingImageComponent() NSDraggingImageComponent {
	class := getNSDraggingImageComponentClass()
	rv := objc.Send[NSDraggingImageComponent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a dragging image component with the specified key.
//
// key: The key.
//
// # Return Value
// 
// An initialized dragging image component with the specified key.
//
// # Discussion
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/init(key:)
func NewDraggingImageComponentWithKey(key NSDraggingImageComponentKey) NSDraggingImageComponent {
	instance := getNSDraggingImageComponentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKey:"), objc.String(string(key)))
	return NSDraggingImageComponentFromID(rv)
}

// Initializes and returns a dragging image component with the specified key.
//
// key: The key.
//
// # Return Value
// 
// An initialized dragging image component with the specified key.
//
// # Discussion
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/init(key:)
func (d NSDraggingImageComponent) InitWithKey(key NSDraggingImageComponentKey) NSDraggingImageComponent {
	rv := objc.Send[NSDraggingImageComponent](d.ID, objc.Sel("initWithKey:"), objc.String(string(key)))
	return rv
}

// The unique name of this image component instance.
//
// # Discussion
// 
// The key must be unique for each component in an [NSDraggingItem] instance.
// You can create your own named components, however the keys described in
// [NSDragImage Component Keys] have special meanings.
// 
// When an NSDraggingItem instances [ImageComponents] are changed by one of
// the `` methods the image associated with this key is morphed into the new
// image component’s image associated with the same key.
//
// [NSDragImage Component Keys]: https://developer.apple.com/documentation/AppKit/nsdragimage-component-keys
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/key
func (d NSDraggingImageComponent) Key() NSDraggingImageComponentKey {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("key"))
	return NSDraggingImageComponentKey(foundation.NSStringFromID(rv).String())
}
func (d NSDraggingImageComponent) SetKey(value NSDraggingImageComponentKey) {
	objc.Send[struct{}](d.ID, objc.Sel("setKey:"), objc.String(string(value)))
}

// An object providing the image contents of the component.
//
// # Discussion
// 
// Typically you set an [NSImage] instance or a [CGImage] as content.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/contents
func (d NSDraggingImageComponent) Contents() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("contents"))
	return objectivec.Object{ID: rv}
}
func (d NSDraggingImageComponent) SetContents(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setContents:"), value)
}

// The coordinate space is the bounds of the parent dragging item.
//
// # Discussion
// 
// The frame is {{0,0}, {`draggingFrame.SizeXCUIElementTypeWidth()`,
// `draggingFrame.SizeXCUIElementTypeHeight()`}}.
// 
// The coordinate space is the bounds of the parent [NSDraggingItem]
// instance’s [DraggingFrame].
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/frame
func (d NSDraggingImageComponent) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
func (d NSDraggingImageComponent) SetFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](d.ID, objc.Sel("setFrame:"), value)
}

