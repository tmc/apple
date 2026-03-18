// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorSampler] class.
var (
	_NSColorSamplerClass     NSColorSamplerClass
	_NSColorSamplerClassOnce sync.Once
)

func getNSColorSamplerClass() NSColorSamplerClass {
	_NSColorSamplerClassOnce.Do(func() {
		_NSColorSamplerClass = NSColorSamplerClass{class: objc.GetClass("NSColorSampler")}
	})
	return _NSColorSamplerClass
}

// GetNSColorSamplerClass returns the class object for NSColorSampler.
func GetNSColorSamplerClass() NSColorSamplerClass {
	return getNSColorSamplerClass()
}

type NSColorSamplerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorSamplerClass) Alloc() NSColorSampler {
	rv := objc.Send[NSColorSampler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that displays the system’s color-sampling interface and returns
// the selected color to your app.
//
// # Overview
// 
// Create an [NSColorSampler] object when you want the user to select a color
// based on existing onscreen colors. When you call the
// [NSColorSampler.ShowSamplerWithSelectionHandler] method, AppKit shows the system’s color
// sampler interface and reports the selected color back to the provided
// block.
//
// # Capturing a Color Sample
//
//   - [NSColorSampler.ShowSamplerWithSelectionHandler]: Displays the system color-sampling interface asynchronously and reports the selected color back to your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSampler
type NSColorSampler struct {
	objectivec.Object
}

// NSColorSamplerFromID constructs a [NSColorSampler] from an objc.ID.
//
// An object that displays the system’s color-sampling interface and returns
// the selected color to your app.
func NSColorSamplerFromID(id objc.ID) NSColorSampler {
	return NSColorSampler{objectivec.Object{ID: id}}
}
// NOTE: NSColorSampler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSColorSampler] class.
//
// # Capturing a Color Sample
//
//   - [INSColorSampler.ShowSamplerWithSelectionHandler]: Displays the system color-sampling interface asynchronously and reports the selected color back to your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSampler
type INSColorSampler interface {
	objectivec.IObject

	// Topic: Capturing a Color Sample

	// Displays the system color-sampling interface asynchronously and reports the selected color back to your app.
	ShowSamplerWithSelectionHandler(selectionHandler ColorHandler)
}

// Init initializes the instance.
func (c NSColorSampler) Init() NSColorSampler {
	rv := objc.Send[NSColorSampler](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorSampler) Autorelease() NSColorSampler {
	rv := objc.Send[NSColorSampler](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorSampler creates a new NSColorSampler instance.
func NewNSColorSampler() NSColorSampler {
	class := getNSColorSamplerClass()
	rv := objc.Send[NSColorSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Displays the system color-sampling interface asynchronously and reports the
// selected color back to your app.
//
// selectionHandler: The handler block for processing the user-selected color. AppKit calls this
// block on your app’s main thread. This block has no return value and takes
// the following parameter:
// 
// selectedColor: The selected color.
//
// # Discussion
// 
// This method displays the color-sampling interface and returns immediately.
// The color-sampling interface magnifies the onscreen pixels and makes it
// easier for the user to select a single pixel. When the user clicks any
// mouse button, AppKit dismisses the interface and calls `selectionHandler`
// with the results.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorSampler/show(selectionHandler:)
func (c NSColorSampler) ShowSamplerWithSelectionHandler(selectionHandler ColorHandler) {
_block0, _cleanup0 := NewColorBlock(selectionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](c.ID, objc.Sel("showSamplerWithSelectionHandler:"), _block0)
}

// ShowSamplerWithSelectionHandlerSync is a synchronous wrapper around [NSColorSampler.ShowSamplerWithSelectionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c NSColorSampler) ShowSamplerWithSelectionHandlerSync(ctx context.Context) (*NSColor, error) {
	done := make(chan *NSColor, 1)
	c.ShowSamplerWithSelectionHandler(func(val *NSColor) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

