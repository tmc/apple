// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol used to provide information about the size and content insets of a layout’s container.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer
type NSCollectionLayoutContainer interface {
	objectivec.IObject

	// The size of the container before content insets are applied.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/contentSize
	ContentSize() corefoundation.CGSize

	// The size of the container after content insets are applied.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/effectiveContentSize
	EffectiveContentSize() corefoundation.CGSize

	// The amount of space added around the content of the container to adjust its final size.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/contentInsets
	ContentInsets() NSDirectionalEdgeInsets

	// The amount of space added around the content of the container to adjust its final size after item content insets are applied.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/effectiveContentInsets
	EffectiveContentInsets() NSDirectionalEdgeInsets
}

// NSCollectionLayoutContainerObject wraps an existing Objective-C object that conforms to the NSCollectionLayoutContainer protocol.
type NSCollectionLayoutContainerObject struct {
	objectivec.Object
}

func (o NSCollectionLayoutContainerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCollectionLayoutContainerObjectFromID constructs a [NSCollectionLayoutContainerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionLayoutContainerObjectFromID(id objc.ID) NSCollectionLayoutContainerObject {
	return NSCollectionLayoutContainerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The size of the container before content insets are applied.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/contentSize
func (o NSCollectionLayoutContainerObject) ContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("contentSize"))
	return corefoundation.CGSize(rv)
}

// The size of the container after content insets are applied.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/effectiveContentSize
func (o NSCollectionLayoutContainerObject) EffectiveContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("effectiveContentSize"))
	return corefoundation.CGSize(rv)
}

// The amount of space added around the content of the container to adjust its
// final size.
//
// # Discussion
//
// If the value of a content inset is less than `1.0`, the content inset is
// fractional. For example, a top content inset of `20.0` adds 20 points of
// space between the top of the layout and the top of its container. A top
// content inset of `0.2` adds space equal to 20% of the container’s height
// between the top of the layout and the top of its container.
//
// [media-3570380]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/contentInsets
func (o NSCollectionLayoutContainerObject) ContentInsets() NSDirectionalEdgeInsets {
	rv := objc.Send[NSDirectionalEdgeInsets](o.ID, objc.Sel("contentInsets"))
	return NSDirectionalEdgeInsets(rv)
}

// The amount of space added around the content of the container to adjust its
// final size after item content insets are applied.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutContainer/effectiveContentInsets
func (o NSCollectionLayoutContainerObject) EffectiveContentInsets() NSDirectionalEdgeInsets {
	rv := objc.Send[NSDirectionalEdgeInsets](o.ID, objc.Sel("effectiveContentInsets"))
	return NSDirectionalEdgeInsets(rv)
}
