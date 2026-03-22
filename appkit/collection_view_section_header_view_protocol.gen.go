// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines a button to control the collapse of a collection view’s section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewSectionHeaderView
type NSCollectionViewSectionHeaderView interface {
	objectivec.IObject
	NSCollectionViewElement
	NSUserInterfaceItemIdentification

	// A control that lets users collapse and open a collection view section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewSectionHeaderView/sectionCollapseButton
	SectionCollapseButton() INSButton

	// A control that lets users collapse and open a collection view section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewSectionHeaderView/sectionCollapseButton
	SetSectionCollapseButton(value INSButton)
}

// NSCollectionViewSectionHeaderViewObject wraps an existing Objective-C object that conforms to the NSCollectionViewSectionHeaderView protocol.
type NSCollectionViewSectionHeaderViewObject struct {
	objectivec.Object
}
func (o NSCollectionViewSectionHeaderViewObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCollectionViewSectionHeaderViewObjectFromID constructs a [NSCollectionViewSectionHeaderViewObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionViewSectionHeaderViewObjectFromID(id objc.ID) NSCollectionViewSectionHeaderViewObject {
	return NSCollectionViewSectionHeaderViewObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A control that lets users collapse and open a collection view section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewSectionHeaderView/sectionCollapseButton
func (o NSCollectionViewSectionHeaderViewObject) SectionCollapseButton() INSButton {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sectionCollapseButton"))
	return NSButtonFromID(rv)
	}
// Performs any necessary cleanup to prepare the element for use again.
//
// # Discussion
// 
// The recycling of content is an important technique for improving
// performance of a collection view. Instead of creating all views from
// scratch, the collection view recycles views and view controllers that move
// offscreen. When your app subsequently calls the
// [ItemWithIdentifierForIndexPath] or
// [SupplementaryViewOfKindWithIdentifierForIndexPath] method, the collection
// view retrieves a recycled object from the appropriate storage, calls this
// method, and then returns the object to your app.
// 
// Implement this method when you need to delete old data or when you want to
// restore your recycled views to a standard initial state prior to their
// reuse. For example, you might use this method to restore the size of a view
// to some standard size or reset the alpha to 1.0 to ensure that the view is
// fully opaque. Do not use this method to configure the view with new data.
// Restoring your views to a default state in this method simplifies the
// configuration code you must write in your data source object later.
// 
// If you implement this method, you must call `super` at some point.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewElement/prepareForReuse()
func (o NSCollectionViewSectionHeaderViewObject) PrepareForReuse() {
	objc.Send[struct{}](o.ID, objc.Sel("prepareForReuse"))
	}
// Asks your element if it wants to modify any layout attributes before they
// are applied.
//
// layoutAttributes: The attributes provided by the layout object. These attributes represent
// the values that the layout object intends to apply to the element.
//
// # Return Value
// 
// The final attributes to apply to the element.
//
// # Discussion
// 
// The default implementation of this method returns the same attributes that
// are in the `layoutAttributes` parameter. You can override this method in
// subclasses and use it to return a different set of attributes. If you
// override this method, call `super` first to give the system the opportunity
// to make changes, then modify the returned attributes.
// 
// # Special Considerations
// 
// In OS X 10.11, this method is never called.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewElement/preferredLayoutAttributesFitting(_:)
func (o NSCollectionViewSectionHeaderViewObject) PreferredLayoutAttributesFittingAttributes(layoutAttributes INSCollectionViewLayoutAttributes) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("preferredLayoutAttributesFittingAttributes:"), layoutAttributes)
	return NSCollectionViewLayoutAttributesFromID(rv)
	}
// Applies the specified layout attributes to the element.
//
// layoutAttributes: The layout attributes to apply.
//
// # Discussion
// 
// In your custom elements, you can use this method to apply the specified
// attributes to your content. For example, if your element object is a view
// controller, you would override this method and use it to apply the
// attributes to the root view object. When using your element with a layout
// object that supports custom attributes, you would also use this method to
// apply those custom attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewElement/apply(_:)
func (o NSCollectionViewSectionHeaderViewObject) ApplyLayoutAttributes(layoutAttributes INSCollectionViewLayoutAttributes) {
	objc.Send[struct{}](o.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
	}
// Tells the element that the layout object of the collection view is about to
// change.
//
// oldLayout: The current layout object used by the collection view.
//
// newLayout: The new layout object that is about to be used by the collection view.
//
// # Discussion
// 
// The default implementation of this method does nothing. Subclasses can
// override it and use it to prepare for the change in layouts.
// 
// # Special Considerations
// 
// In OS X 10.11, this method is never called.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewElement/willTransition(from:to:)
func (o NSCollectionViewSectionHeaderViewObject) WillTransitionFromLayoutToLayout(oldLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) {
	objc.Send[struct{}](o.ID, objc.Sel("willTransitionFromLayout:toLayout:"), oldLayout, newLayout)
	}
// Tells the element that the layout object of the collection view changed.
//
// oldLayout: The collection view’s previous layout object.
//
// newLayout: The current layout object associated with the collection view.
//
// # Discussion
// 
// The default implementation of this method does nothing. Subclasses can
// override it and use it to finalize any behaviors associated with the change
// in layouts.
// 
// # Special Considerations
// 
// In OS X 10.11, this method is never called.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewElement/didTransition(from:to:)
func (o NSCollectionViewSectionHeaderViewObject) DidTransitionFromLayoutToLayout(oldLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) {
	objc.Send[struct{}](o.ID, objc.Sel("didTransitionFromLayout:toLayout:"), oldLayout, newLayout)
	}
// A string that identifies the user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (o NSCollectionViewSectionHeaderViewObject) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
	}

func (o NSCollectionViewSectionHeaderViewObject) SetSectionCollapseButton(value INSButton) {
	objc.Send[struct{}](o.ID, objc.Sel("setSectionCollapseButton:"), value)
}

func (o NSCollectionViewSectionHeaderViewObject) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

