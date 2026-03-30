// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that provides a container in which the find bar is displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer
type NSTextFinderBarContainer interface {
	objectivec.IObject

	// The view assigned by the text bar as the find bar view for the container.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarView
	FindBarView() INSView

	// Returns whether the container should display its find bar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/isFindBarVisible
	IsFindBarVisible() bool

	// Notifies the find bar container that the find bar has changed its height.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarViewDidChangeHeight()
	FindBarViewDidChangeHeight()

	// The view assigned by the text bar as the find bar view for the container.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarView
	SetFindBarView(value INSView)

	// Returns whether the container should display its find bar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/isFindBarVisible
	SetFindBarVisible(value bool)
}

// NSTextFinderBarContainerObject wraps an existing Objective-C object that conforms to the NSTextFinderBarContainer protocol.
type NSTextFinderBarContainerObject struct {
	objectivec.Object
}

func (o NSTextFinderBarContainerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextFinderBarContainerObjectFromID constructs a [NSTextFinderBarContainerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextFinderBarContainerObjectFromID(id objc.ID) NSTextFinderBarContainerObject {
	return NSTextFinderBarContainerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The view assigned by the text bar as the find bar view for the container.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarView
func (o NSTextFinderBarContainerObject) FindBarView() INSView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("findBarView"))
	return NSViewFromID(rv)
}

// Returns whether the container should display its find bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/isFindBarVisible
func (o NSTextFinderBarContainerObject) IsFindBarVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isFindBarVisible"))
	return rv
}

// Notifies the find bar container that the find bar has changed its height.
//
// # Discussion
//
// Upon receiving this message the container may be required to re-tile its
// contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarViewDidChangeHeight()
func (o NSTextFinderBarContainerObject) FindBarViewDidChangeHeight() {
	objc.Send[struct{}](o.ID, objc.Sel("findBarViewDidChangeHeight"))
}

// A view hierarchy that contains all the views which display the contents
// being searched.
//
// # Return Value
//
// The root view of the content view hierarchy.
//
// # Discussion
//
// This content view defines the view hierarchy to be dimmed during
// incremental search, if the [NSTextFinder] instance’s
// [IncrementalSearchingShouldDimContentView] is true. If this method is not
// implemented or returns `nil`, then the [NSTextFinder] instance will act as
// if [IncrementalSearchingShouldDimContentView] is false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/contentView()
func (o NSTextFinderBarContainerObject) ContentView() INSView {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentView"))
	return NSViewFromID(rv)
}

// The view assigned by the text bar as the find bar view for the container.
//
// # Discussion
//
// This property is managed by [NSTextFinder] and you must not set this
// property.
//
// The container may freely modify the view’s width, but should not modify
// its height.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarView
func (o NSTextFinderBarContainerObject) SetFindBarView(value INSView) {
	objc.Send[struct{}](o.ID, objc.Sel("setFindBarView:"), value)
}

// Returns whether the container should display its find bar.
//
// # Discussion
//
// When this property is true and the [FindBarView] property is set, then the
// find bar is displayed by the container. Otherwise, the find bar is not
// displayed.
//
// The default value should be false.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/isFindBarVisible
func (o NSTextFinderBarContainerObject) SetFindBarVisible(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setFindBarVisible:"), value)
}
