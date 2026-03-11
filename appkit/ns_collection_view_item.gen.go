// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSCollectionViewItem] class.
var (
	_NSCollectionViewItemClass     NSCollectionViewItemClass
	_NSCollectionViewItemClassOnce sync.Once
)

func getNSCollectionViewItemClass() NSCollectionViewItemClass {
	_NSCollectionViewItemClassOnce.Do(func() {
		_NSCollectionViewItemClass = NSCollectionViewItemClass{class: objc.GetClass("NSCollectionViewItem")}
	})
	return _NSCollectionViewItemClass
}

// GetNSCollectionViewItemClass returns the class object for NSCollectionViewItem.
func GetNSCollectionViewItemClass() NSCollectionViewItemClass {
	return getNSCollectionViewItemClass()
}

type NSCollectionViewItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewItemClass) Alloc() NSCollectionViewItem {
	rv := objc.Send[NSCollectionViewItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The visual representation for a single data element in a collection view.
//
// # Overview
// 
// Item objects are view controllers, and you use their view hierarchies to
// display your content. The default implementation of this class supports the
// creation of a simple item that displays a single image or string. If the
// appearance or layout of your items is more sophisticated, you can subclass
// and configure the view hierarchy based on your needs.
// 
// Items are the most common types of elements displayed by a collection view,
// and every collection view must have at least one type of item. You use
// items to represent the main content of your collection view interface. For
// example, a photo browser app would use items to display individual photos.
// Remember that items are only the visual interpretation of your app’s
// underlying data. The actual data is always managed by your app and exposed
// to the collection view through the data source object, which uses the data
// to configure the items that are displayed.
// 
// The use of items with a collection view requires doing the following:
// 
// - Define the visual appearance of your items by specifying what views they
// contain and how those views are arranged. - When your interface is first
// loaded, register your items with the collection view. (You must register
// your items before the collection view tries to display any content.) - In
// your data source object, create and configure items when the collection
// view asks for them; see [NSCollectionViewDataSource].
// 
// At runtime, items merely present the data they are given. Your app’s data
// structures are always the original source of content, and the item contains
// only a copy of that data to present to the user. When the underlying data
// associated with an item changes, the data source should invalidate the item
// by calling the [ReloadItemsAtIndexPaths] method of the collection view.
// Invalidating an item forces the collection view to dispose of it so that
// the collection view can create a new one with the updated content.
// 
// For information about how the collection view displays items to the user,
// see [NSCollectionView].
// 
// # Configuring an Item’s Views
// 
// You configure the views of an item in one of two ways:
// 
// - Subclass [NSCollectionViewItem] and create any custom views
// programmatically. - Create a nib file containing a single top-level
// [NSCollectionViewItem] object. Embed any custom views in the root view of
// the item.
// 
// When creating the views programmatically, you typically override the
// item’s [LoadView] method as you would for any view controller. In your
// implementation, create the views and add them as subviews to the view
// controller’s root view. Add accessor properties to your subclass so that
// you can access the views later and configure them.
// 
// When using a nib file, you can use a generic [NSCollectionViewItem] object
// if your item contains only an image or text field. For more complex item
// content, subclass [NSCollectionViewItem] and add outlets for any views you
// need to access later. In Interface Builder, connect your outlets to the
// views you add to the nib file.
// 
// # Registering Items
// 
// Before you can ask the collection view to create items, you must register
// those items using one of the following methods:
// 
// - Use the [RegisterClassForItemWithIdentifier] method when your
// [NSCollectionViewItem] subclass handles the creation of its own views. -
// Use the [RegisterNibForItemWithIdentifier] method when you store the
// item’s views in a nib file.
// 
// You must register at least one item type before trying to display content
// from your collection view. The collection view’s data source uses the
// [ItemWithIdentifierForIndexPath] method to fetch an empty item for
// configuration. During the initial configuration of the collection view,
// that method creates all items using the registered classes and nib files
// you provided. Later on, the method may return a recycled item that can be
// repurposed with new data.
// 
// For more information about how how to register items, see
// [NSCollectionView]. For information about how the data source object
// creates and configures items, see [NSCollectionViewDataSource].
// 
// # Legacy Item Support
// 
// For apps built before OS X 10.11, you created a template item and assigned
// it to the [itemPrototype] property of your collection view. To create new
// instances of the item, you called the collection view’s
// [newItem(forRepresentedObject:)] method. For more information about how to
// support older collection view configurations, see Collection View
// Programming Guide for macOS.
//
// [itemPrototype]: https://developer.apple.com/documentation/AppKit/NSCollectionView/itemPrototype
// [newItem(forRepresentedObject:)]: https://developer.apple.com/documentation/AppKit/NSCollectionView/newItem(forRepresentedObject:)
//
// # Getting and Setting Image and Text Fields
//
//   - [NSCollectionViewItem.ImageView]: An image view outlet that you can use to display images.
//   - [NSCollectionViewItem.SetImageView]
//   - [NSCollectionViewItem.TextField]: A text field outlet that you can use to display a string.
//   - [NSCollectionViewItem.SetTextField]
//
// # Managing the Selection and Highlight States
//
//   - [NSCollectionViewItem.Selected]: A Boolean indicating whether the item is currently selected.
//   - [NSCollectionViewItem.SetSelected]
//   - [NSCollectionViewItem.HighlightState]: The highlight state currently applied to the item.
//   - [NSCollectionViewItem.SetHighlightState]
//
// # Getting the Parent Collection View
//
//   - [NSCollectionViewItem.CollectionView]: The collection view that owns the item.
//
// # Dragging Components
//
//   - [NSCollectionViewItem.DraggingImageComponents]: Dragging images for multi-image drag and drop support.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem
type NSCollectionViewItem struct {
	NSViewController
}

// NSCollectionViewItemFromID constructs a [NSCollectionViewItem] from an objc.ID.
//
// The visual representation for a single data element in a collection view.
func NSCollectionViewItemFromID(id objc.ID) NSCollectionViewItem {
	return NSCollectionViewItem{NSViewController: NSViewControllerFromID(id)}
}
// NOTE: NSCollectionViewItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionViewItem] class.
//
// # Getting and Setting Image and Text Fields
//
//   - [INSCollectionViewItem.ImageView]: An image view outlet that you can use to display images.
//   - [INSCollectionViewItem.SetImageView]
//   - [INSCollectionViewItem.TextField]: A text field outlet that you can use to display a string.
//   - [INSCollectionViewItem.SetTextField]
//
// # Managing the Selection and Highlight States
//
//   - [INSCollectionViewItem.Selected]: A Boolean indicating whether the item is currently selected.
//   - [INSCollectionViewItem.SetSelected]
//   - [INSCollectionViewItem.HighlightState]: The highlight state currently applied to the item.
//   - [INSCollectionViewItem.SetHighlightState]
//
// # Getting the Parent Collection View
//
//   - [INSCollectionViewItem.CollectionView]: The collection view that owns the item.
//
// # Dragging Components
//
//   - [INSCollectionViewItem.DraggingImageComponents]: Dragging images for multi-image drag and drop support.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem
type INSCollectionViewItem interface {
	INSViewController
	NSCollectionViewElement

	// Topic: Getting and Setting Image and Text Fields

	// An image view outlet that you can use to display images.
	ImageView() INSImageView
	SetImageView(value INSImageView)
	// A text field outlet that you can use to display a string.
	TextField() INSTextField
	SetTextField(value INSTextField)

	// Topic: Managing the Selection and Highlight States

	// A Boolean indicating whether the item is currently selected.
	Selected() bool
	SetSelected(value bool)
	// The highlight state currently applied to the item.
	HighlightState() NSCollectionViewItemHighlightState
	SetHighlightState(value NSCollectionViewItemHighlightState)

	// Topic: Getting the Parent Collection View

	// The collection view that owns the item.
	CollectionView() INSCollectionView

	// Topic: Dragging Components

	// Dragging images for multi-image drag and drop support.
	DraggingImageComponents() []NSDraggingImageComponent

	// The receiver’s collection view item prototype.
	ItemPrototype() INSCollectionViewItem
	SetItemPrototype(value INSCollectionViewItem)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSCollectionViewItem) Init() NSCollectionViewItem {
	rv := objc.Send[NSCollectionViewItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewItem) Autorelease() NSCollectionViewItem {
	rv := objc.Send[NSCollectionViewItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewItem creates a new NSCollectionViewItem instance.
func NewNSCollectionViewItem() NSCollectionViewItem {
	class := getNSCollectionViewItemClass()
	rv := objc.Send[NSCollectionViewItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewCollectionViewItemWithCoder(coder foundation.INSCoder) NSCollectionViewItem {
	instance := getNSCollectionViewItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCollectionViewItemFromID(rv)
}


// Returns a view controller object initialized to the nib file in the
// specified bundle.
//
// nibNameOrNil: The name of the nib file, without any leading path information.
//
// nibBundleOrNil: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
// 
// The initialized [NSViewController] object.
//
// # Discussion
// 
// The [NSViewController] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the Resources
// directory.
// 
// The specified nib file should typically have the class of the file’s
// owner set to [NSViewController], or a custom subclass, with the `view`
// outlet connected to a view.
// 
// If you pass in `nil` for `nibNameOrNil`, [NibName] returns `nil` and
// [LoadView] throws an exception; in this case you must set [View] before
// [View] is invoked, or override [LoadView].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewCollectionViewItemWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSCollectionViewItem {
	instance := getNSCollectionViewItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSCollectionViewItemFromID(rv)
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
func (c NSCollectionViewItem) ApplyLayoutAttributes(layoutAttributes INSCollectionViewLayoutAttributes) {
	objc.Send[objc.ID](c.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
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
func (c NSCollectionViewItem) DidTransitionFromLayoutToLayout(oldLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) {
	objc.Send[objc.ID](c.ID, objc.Sel("didTransitionFromLayout:toLayout:"), oldLayout, newLayout)
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
func (c NSCollectionViewItem) PreferredLayoutAttributesFittingAttributes(layoutAttributes INSCollectionViewLayoutAttributes) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("preferredLayoutAttributesFittingAttributes:"), layoutAttributes)
	return NSCollectionViewLayoutAttributesFromID(rv)
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
func (c NSCollectionViewItem) PrepareForReuse() {
	objc.Send[objc.ID](c.ID, objc.Sel("prepareForReuse"))
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
func (c NSCollectionViewItem) WillTransitionFromLayoutToLayout(oldLayout INSCollectionViewLayout, newLayout INSCollectionViewLayout) {
	objc.Send[objc.ID](c.ID, objc.Sel("willTransitionFromLayout:toLayout:"), oldLayout, newLayout)
}
func (c NSCollectionViewItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// An image view outlet that you can use to display images.
//
// # Discussion
// 
// This is a convenience property for accessing an image view in your item’s
// view hierarchy. Normally, you configure this property in Interface Builder
// by connecting it to one of your item’s image views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/imageView
func (c NSCollectionViewItem) ImageView() INSImageView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("imageView"))
	return NSImageViewFromID(objc.ID(rv))
}
func (c NSCollectionViewItem) SetImageView(value INSImageView) {
	objc.Send[struct{}](c.ID, objc.Sel("setImageView:"), value)
}



// A text field outlet that you can use to display a string.
//
// # Discussion
// 
// This is a convenience property for accessing a text field in your item’s
// view hierarchy. Normally, you configure this property in Interface Builder
// by connecting it to one of your item’s text fields.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/textField
func (c NSCollectionViewItem) TextField() INSTextField {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("textField"))
	return NSTextFieldFromID(objc.ID(rv))
}
func (c NSCollectionViewItem) SetTextField(value INSTextField) {
	objc.Send[struct{}](c.ID, objc.Sel("setTextField:"), value)
}



// A Boolean indicating whether the item is currently selected.
//
// # Discussion
// 
// The value of this property is [true] when the item is selected or [false]
// when it is not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/isSelected
func (c NSCollectionViewItem) Selected() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSelected"))
	return rv
}
func (c NSCollectionViewItem) SetSelected(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelected:"), value)
}



// The highlight state currently applied to the item.
//
// # Discussion
// 
// The highlight state provides a visual indication of operations happening to
// items in the collection view. The highlight state normally toggles between
// the [NSCollectionViewItem.HighlightState.none] and
// [NSCollectionViewItem.HighlightState.forSelection] states, but other states
// may be applied to indicate transient conditions. For example, the
// [NSCollectionViewItem.HighlightState.forDeselection] state is applied
// during interactive selections when a currently selected item is about to be
// deselected.
//
// [NSCollectionViewItem.HighlightState.forDeselection]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/forDeselection
// [NSCollectionViewItem.HighlightState.forSelection]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/forSelection
// [NSCollectionViewItem.HighlightState.none]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/none
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/highlightState-swift.property
func (c NSCollectionViewItem) HighlightState() NSCollectionViewItemHighlightState {
	rv := objc.Send[NSCollectionViewItemHighlightState](c.ID, objc.Sel("highlightState"))
	return NSCollectionViewItemHighlightState(rv)
}
func (c NSCollectionViewItem) SetHighlightState(value NSCollectionViewItemHighlightState) {
	objc.Send[struct{}](c.ID, objc.Sel("setHighlightState:"), value)
}



// The collection view that owns the item.
//
// # Discussion
// 
// Use this property as a convenient way to access the collection view that
// owns the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/collectionView
func (c NSCollectionViewItem) CollectionView() INSCollectionView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collectionView"))
	return NSCollectionViewFromID(objc.ID(rv))
}



// Dragging images for multi-image drag and drop support.
//
// # Discussion
// 
// The component frames are relative to a coordinate system that has its
// origin at the bottom left, so you need to take into account the flipped
// state of your view when computing the component frames.
// 
// This methods can be subclassed and overridden to provide a custom set of
// [NSDraggingImageComponent] objects to create the drag image.
// 
// The default implementation will return an array of up to two
// [NSDraggingImageComponent] instances – one for the [ImageView] and
// another for the [TextField] (if not `nil`).
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/draggingImageComponents
func (c NSCollectionViewItem) DraggingImageComponents() []NSDraggingImageComponent {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("draggingImageComponents"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSDraggingImageComponent {
		return NSDraggingImageComponentFromID(id)
	})
}



// The receiver’s collection view item prototype.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/itemprototype
func (c NSCollectionViewItem) ItemPrototype() INSCollectionViewItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemPrototype"))
	return NSCollectionViewItemFromID(objc.ID(rv))
}
func (c NSCollectionViewItem) SetItemPrototype(value INSCollectionViewItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setItemPrototype:"), value)
}














			// Protocol methods for NSCollectionViewElement
			
























