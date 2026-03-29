// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScrubber] class.
var (
	_NSScrubberClass     NSScrubberClass
	_NSScrubberClassOnce sync.Once
)

func getNSScrubberClass() NSScrubberClass {
	_NSScrubberClassOnce.Do(func() {
		_NSScrubberClass = NSScrubberClass{class: objc.GetClass("NSScrubber")}
	})
	return _NSScrubberClass
}

// GetNSScrubberClass returns the class object for NSScrubber.
func GetNSScrubberClass() NSScrubberClass {
	return getNSScrubberClass()
}

type NSScrubberClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScrubberClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberClass) Alloc() NSScrubber {
	rv := objc.Send[NSScrubber](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A customizable item picker control for the Touch Bar.
//
// # Overview
// 
// On supported MacBook Pro models, you can use a scrubber (an instance of the
// [NSScrubber] class) to provide a horizontally-oriented, item-picker control
// in the Touch Bar. Use a scrubber to let the user pick an item from a
// related collection, such as a photo from a library or a date from a date
// range.
// 
// Refer to the following sample code projects which demonstrate how to use
// [NSTouchBar] and related classes, including the [NSScrubber] class:
// 
// - [Creating and Customizing the Touch Bar] - [Integrating a Toolbar and
// Touch Bar into Your App]
// 
// Each item that appears in a scrubber is a specialized view that supports
// selection and scrubber-appropriate decorations. The scrubber keeps track of
// its items by their index positions.
// 
// There are many classes in the scrubber API, as well as a delegate protocol,
// a data source protocol, and a callback-based layout API. The design pattern
// is reminiscent of that used for a collection view (an instance of the
// [NSCollectionView] class). You might find it helpful to refer to the
// [NSCollectionView] overview for background. Be aware, though of the
// differences. For example, while scrubbers and collection views both employ
// a [ItemWithIdentifierOwner] method, and both employ a reuse queue, a
// scrubber is subclassed from the [NSView] class while a collection view is
// subclassed from the [NSViewController] class.
// 
// A scrubber employs:
// 
// - The itself (an instance of the [NSScrubber] class), which serves as a
// container view that shows a subview for each scrubber item, and which
// employs a reuse-queue pattern for efficiency and performance. - A
// (conforming to the [NSScrubberDataSource] protocol), which provides
// scrubber items to the scrubber, on demand, from an associated data
// collection in your app. Specify the data source in the scrubber’s
// [NSScrubber.DataSource] property - A (conforming to the [NSScrubberDelegate]
// protocol), which responds to user interaction — such as with its
// [DidBeginInteractingWithScrubber] and [DidCancelInteractingWithScrubber]
// methods. Specify the delegate in the scrubber’s [NSScrubber.Delegate] property. You
// can also use the delegate to respond to the highlighting and selection of
// scrubber items, and to respond to changes in which items are visible in the
// scrubber. - A (an instance of a subclass of the [NSScrubberLayout] abstract
// class, typically the [NSScrubberFlowLayout] concrete subclass). You
// implement a layout to respond to calls, from the system, to return view
// specifications for the items to be displayed in the scrubber. The layout,
// in this way, assists in arranging and decorating the scrubber’s contained
// items, and in providing appearance changes in response to user interaction.
// Specify the layout in the scrubber’s [ScrubberLayout] property.
// 
// Before learning how to use a scrubber in the Touch Bar, be sure you read
// the overview for the [NSTouchBar] class.
// 
// # Scrubber data source and delegate
// 
// A scrubber employs a data source and a delegate, using a pattern similar to
// that used for collection views, as follows:
// 
// To supply items for a scrubber, implement an object that conforms to the
// [NSScrubberDataSource] protocol and specify that object in the scrubber’s
// [NSScrubber.DataSource] property. There are two built-in item types, provided by the
// [NSScrubberTextItemView] and [NSScrubberImageItemView] concrete classes.
// For more on scrubber items, see [NSScrubber].
// 
// The following code shows an example implementation of the [NSScrubber.NumberOfItems]
// datasource method, returning the count of items displayed by the scrubber.
// 
// In addition to the count of scrubber items, you use the datasource method
// to provide individual items with the [ScrubberViewForItemAtIndex] method.
// An example implementation is shown in the following code.
// 
// To optimize resource usage and performance, a scrubber employs a reuse
// queue that’s similar to the reuse queue for an [NSCollectionView] object.
// 
// To respond to user interactions and to visibility, highlighting, and
// selection changes, implement a delegate object that conforms to the
// [NSScrubberDelegate] protocol and specify that object in the scrubber’s
// [NSScrubber.Delegate] property.
// 
// The following code shows a minimal implementation of the
// [ScrubberDidSelectItemAtIndex] delegate method for a scrubber.
// 
// # Choose a scrubber touch-interaction model
// 
// A scrubber offers many built-in permutations for touch interaction. By
// subclassing a scrubber, you can customize touch interaction.
// 
// To specify a scrubber’s touch-interaction model, set values for the
// following, cooperating scrubber properties: [NSScrubber.Mode], [NSScrubber.Continuous], and
// [NSScrubber.ItemAlignment]. Here’s how to choose the right permutation of values for
// these properties:
// 
// Decide whether you want the scrubber to to track horizontal finger movement
// across the scrubber, or to remain in place as the finger moves.
// 
// - For scrolling, specify the [NSScrubber.Mode.free] value for the
// scrubber’s [NSScrubber.Mode] property. - For a fixed scrubber, specify the
// [NSScrubber.Mode.fixed] value for the [NSScrubber.Mode] property (this is the default
// value). In this case, if the user’s finger reaches the left or right edge
// of the scrubber view and there are items beyond the edge, the scrubber
// automatically scrolls to bring those items into view.
// 
// Decide whether you want item selection to take place only upon a deliberate
// selection gesture, or continuously during horizontal finger movement on the
// scrubber.
// 
// - For deliberate selection, specify a value of [false] for the scrubber’s
// [NSScrubber.Continuous] property (this is the default value). In (scrolling) mode, the
// user must then tap an item to highlight and select it. In (non-scrolling)
// mode, ending interaction with the scrubber, by lifting the finger, selects
// the most-recently highlighted item. However, if there is already a
// highlighted item before interaction starts, and the user resumes
// interacting with the (fixed mode) scrubber on that item, selection changes
// continuously, tracking the user’s finger — even though the [NSScrubber.Continuous]
// property value is [false]. - For continuous selection, specify a value of
// [true] for the [NSScrubber.Continuous] property. Item selection behavior then depends
// on the [NSScrubber.Mode] and [NSScrubber.ItemAlignment] property values, as described in
// [NSScrubber].
// 
// The setting in the scrubber’s [NSScrubber.ItemAlignment] property affects two
// things: 1) item highlighting and selection, and 2) the resting position of
// scrubber items after manual or automatic scrolling. Available values for
// this property are [NSScrubber.Alignment.leading],
// [NSScrubber.Alignment.center], [NSScrubber.Alignment.trailing], and
// [NSScrubber.Alignment.none]. See the [NSScrubber.Alignment] enumeration for
// details on how these constants work.
// 
// Your choices for scrolling, selection, and alignment jointly impact
// highlighting and selection behavior. For details on highlighting and
// selection, see [NSScrubber]. Your choice for alignment also impacts
// scrubber-item resting-position behavior following a scroll interaction. For
// details on resting position, see [NSScrubber].
// 
// # Position-based scrubber item selection
// 
// In free mode with continuous selection style (the [NSScrubber.Mode] property value is
// [NSScrubber.Mode.free] and the [NSScrubber.Continuous] property value is [YES] for
// this configuration), the scrubber item on the alignment axis is
// automatically highlighted and selected. The is the left edge, right edge,
// or center of the scrubber, as you specify by setting the value of the
// [NSScrubber.ItemAlignment] property using constants from the [NSScrubber.Alignment]
// enumeration. Specifying an alignment axis of [NSScrubber.Alignment.none] is
// equivalent to a value of [NSScrubber.Alignment.center] for position-based
// item selection.
// 
// In free mode with deliberate selection style (the [NSScrubber.Mode] property value is
// [NSScrubber.Mode.free] and the [NSScrubber.Continuous] property value is [NO] for this
// configuration), the system ignores the [NSScrubber.ItemAlignment] property value in
// terms of item selection.
// 
// In fixed mode (the [NSScrubber.Mode] property value is [NSScrubber.Mode.fixed] for
// this configuration), the system ignores the [NSScrubber.ItemAlignment] property value
// in terms of item selection — no matter which value you specify for the
// [NSScrubber.Continuous] property.
// 
// # Scrubber item resting position
// 
// The value you provide in the [NSScrubber.ItemAlignment] property specifies the
// automatic scrubber item resting position that follows manual or automatic
// scrolling. (This value also affects item highlighting and selection, as
// described in [NSScrubber].) The system respects your setting for resting
// position irrespective of the values of the [NSScrubber.Mode] and [NSScrubber.Continuous]
// properties.
// 
// Specifically:
// 
// - [NSScrubber.Alignment.leading] — In a left-to-right language, the
// scrubber comes to rest, following manual or automatic scrolling, so that
// the left edge of the leftmost scrubber item is coincident with the left
// edge of the scrubber. - [NSScrubber.Alignment.center] — The scrubber
// comes to rest, following manual or automatic scrolling, so that a scrubber
// item is perfectly centered in the scrubber. -
// [NSScrubber.Alignment.trailing] — In a left-to-right language, the
// scrubber comes to rest, following manual or automatic scrolling, so that
// the right edge of the rightmost scrubber item is coincident with the right
// edge of the scrubber. - [NSScrubber.Alignment.none] — Following manual or
// automatic scrolling, the scrubber comes to rest without attempting to align
// any scrubber item.
// 
// # Scrubber layout
// 
// A scrubber configures the views for its items with the help of two classes,
// [NSScrubberLayout] and [NSScrubberLayoutAttributes], as described in this
// section.
// 
// # Layout implementation
// 
// A is a concrete implementation of the [NSScrubberLayout] abstract class.
// AppKit provides two concrete, preconfigured layout subclasses:
// [NSScrubberFlowLayout] and [NSScrubberProportionalLayout]. If you use one
// of these built-in layout types, there’s no additional layout code to
// write, apart from adding your choice of built-in layout to the scrubber’s
// [ScrubberLayout] property. This Swift example shows this simple step for
// the flow layout:
// 
// To create a custom layout, subclass the [NSScrubberLayout] class and
// implement its callback methods. Unlike a view delegate (such as used for a
// table view), which provides on demand, scrubber layout callbacks provide on
// demand. Using these callbacks, you specify:
// 
// - Scrubber item geometry - Scrubber item appearance - Layout life cycle for
// state management
// 
// Specify the overall visual dimensions of a custom scrubber when you create
// it, using the [NSScrubber.InitWithFrame] or [NSScrubber.InitWithCoder] initializer, or by using
// Interface Builder.
// 
// Return the total width and height for the elements in a custom scrubber,
// including those not currently visible, using the [NSScrubber.ScrubberContentSize]
// property in your layout. Specify height and width in points. To use the
// standard height, specify a value of `30`.
// 
// Specify the geometry and appearance for items in your custom scrubber,
// using the two required callback methods that each return instances of the
// [NSScrubberLayoutAttributes] class. The system calls one or another of
// these methods, as it needs to, as a user interacts with a layout’s owning
// scrubber:
// 
// [Table data omitted]
// 
// You can explicitly invalidate a layout by calling the [InvalidateLayout]
// method. Do this whenever your app changes a scrubber’s information in a
// way that requires a layout update. For example, if you change the text
// shown in one or more items, invalidate the layout.
// 
// You can specify layout life cycle in terms of the conditions under which a
// layout should be automatically invalidated, such as when the user selects
// something different in the layout’s owning scrubber. The API for
// automatic invalidation consists of the following two properties and one
// method:
// 
// - [NSScrubber.ShouldInvalidateLayoutForSelectionChange] -
// [NSScrubber.ShouldInvalidateLayoutForHighlightChange] -
// [ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect]
// 
// For example, if you design a scrubber’s layout characteristics to depend
// on which of its items is selected by the user, return a value of [true]
// from the scrubber’s [NSScrubber.ShouldInvalidateLayoutForSelectionChange] method.A
// object is an instance of the [NSScrubberLayoutAttributes] class, which you
// configure to describe the view for a single item. The class offers the
// following built-in attributes for you to work with:
// 
// - [NSScrubber.ItemIndex] — The item’s index position within the scrubber - [NSScrubber.Frame]
// — The item’s frame rectangle - [NSScrubber.Alpha] — The item’s transparency
// 
// You can specify additional item attributes by subclassing the
// [NSScrubberLayoutAttributes] class. For example, you could specify a
// geometric transform attribute.
// 
// If you’re using a custom [NSScrubberLayout] subclass, provide an
// implementation for the [InvalidateLayout] method to clear any custom layout
// state, such as by discarding cached data.
// 
// # Prepare for redrawing
// 
// The flip side of layout invalidation (as described in [NSScrubber]) is
// preparation for redrawing, which you perform in a layout’s
// [PrepareLayout] method. The goal of layout preparation is to optimize
// performance. A scrubber calls the [PrepareLayout] method exactly once
// between invalidation and redrawing. Complete as much one-time, up-front
// layout work as you can, in advance of redrawing, in this step. For example,
// your [PrepareLayout] implementation should perform initial layout
// calculations and should fill caches needed during drawing.
// 
// After the [PrepareLayout] method returns, the system updates the scrubber
// view hierarchy with repeated calls to three [NSScrubberLayout] methods:
// [LayoutAttributesForItemAtIndex], [LayoutAttributesForItemsInRect], and
// [NSScrubber.ScrubberContentSize]. Implement these methods to provide return values as
// quickly as possible, taking advantage of the work you did during layout
// preparation.
// 
// # Scrubber items
// 
// The view that represents a scrubber item is provided by your data source
// object, using the [ScrubberViewForItemAtIndex] protocol method. AppKit
// provides two purpose-built view classes you can use, both of which are
// concrete subclasses of the abstract [NSScrubberItemView] class:
// 
// - [NSScrubberImageItemView] has `image`, [ImageView], and [NSScrubber.ImageAlignment]
// properties - [NSScrubberTextItemView] has [TextField] and `title`
// properties
// 
// To create a custom item, subclass these or their abstract superclass,
// [NSScrubberItemView].
// 
// # Scrubbers and the responder chain
// 
// To show a scrubber, associate it with an [NSTouchBar] object (adding it, as
// the view for a custom item or popover item, to the bar) and then associate
// the bar with the appropriate responder object in your app. The system then
// shows the scrubber in the Touch Bar only at appropriate times. For more
// information on bars and the responder chain, read the overview for the
// [NSTouchBar] class.
// 
// # Choose between a scrubber and a scroll view
// 
// When choosing between a scrubber and a scroll view, use a scrubber unless
// the amount of content, or the nature of your content, doesn’t work well
// in a scrubber. Scrubber interaction is optimized for the Touch Bar,
// typically making a scrubber the better option for letting the user pick
// from among several choices, such as dates in a calendar.
//
// [Creating and Customizing the Touch Bar]: https://developer.apple.com/documentation/AppKit/creating-and-customizing-the-touch-bar
// [Integrating a Toolbar and Touch Bar into Your App]: https://developer.apple.com/documentation/AppKit/integrating-a-toolbar-and-touch-bar-into-your-app
// [NSScrubber.Alignment.center]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/center
// [NSScrubber.Alignment.leading]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/leading
// [NSScrubber.Alignment.none]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/none
// [NSScrubber.Alignment.trailing]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/trailing
// [NSScrubber.Alignment]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment
// [NSScrubber.Mode.fixed]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/fixed
// [NSScrubber.Mode.free]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/free
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring the scrubber
//
//   - [NSScrubber.DataSource]: The object that provides the data for the scrubber.
//   - [NSScrubber.SetDataSource]
//   - [NSScrubber.Delegate]: The object that acts as the delegate of the scrubber.
//   - [NSScrubber.SetDelegate]
//
// # Creating scrubber items
//
//   - [NSScrubber.RegisterClassForItemIdentifier]: Registers a class for the scrubber to use when it creates new items.
//   - [NSScrubber.RegisterNibForItemIdentifier]: Registers a nib file for the scrubber to use when it creates new items in the scrubber.
//   - [NSScrubber.MakeItemWithIdentifierOwner]: Creates or returns a reusable item object with the specified identifier.
//
// # Changing the layout
//
//   - [NSScrubber.ScrubberLayout]: An object used to describe the layout of items within the scrubber.
//   - [NSScrubber.SetScrubberLayout]
//   - [NSScrubber.Mode]: A setting that determines whether interaction with the scrubber is fixed or free.
//   - [NSScrubber.SetMode]
//   - [NSScrubber.ItemAlignment]: A setting that specifies the snapping behavior of items in the scrubber.
//   - [NSScrubber.SetItemAlignment]
//   - [NSScrubber.Continuous]: A Boolean value that, together with the [mode](<doc://com.apple.appkit/documentation/AppKit/NSScrubber/mode-swift.property>) property, determines scrubber interaction style.
//   - [NSScrubber.SetContinuous]
//
// # Configuring the scrubber’s appearance
//
//   - [NSScrubber.BackgroundColor]: The color displayed behind the scrubber content.
//   - [NSScrubber.SetBackgroundColor]
//   - [NSScrubber.BackgroundView]: A view that is displayed behind the scrubber content.
//   - [NSScrubber.SetBackgroundView]
//   - [NSScrubber.ShowsAdditionalContentIndicators]: A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
//   - [NSScrubber.SetShowsAdditionalContentIndicators]
//   - [NSScrubber.ShowsArrowButtons]: A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
//   - [NSScrubber.SetShowsArrowButtons]
//
// # Configuring the selection appearance
//
//   - [NSScrubber.FloatsSelectionViews]: A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
//   - [NSScrubber.SetFloatsSelectionViews]
//   - [NSScrubber.SelectionOverlayStyle]: The style overlaid on selected items.
//   - [NSScrubber.SetSelectionOverlayStyle]
//   - [NSScrubber.SelectionBackgroundStyle]: The style applied to the background of selected items.
//   - [NSScrubber.SetSelectionBackgroundStyle]
//
// # Reloading content
//
//   - [NSScrubber.ReloadData]: Reloads the content of the entire scrubber, and deselects the currently selected item.
//   - [NSScrubber.ReloadItemsAtIndexes]: Reloads the items at the specified indexes.
//
// # Getting the state of the scrubber
//
//   - [NSScrubber.NumberOfItems]: The number of items represented by the scrubber.
//   - [NSScrubber.HighlightedIndex]: The index of the highlighted item in the scrubber.
//   - [NSScrubber.SelectedIndex]: The index of the selected item in the scrubber.
//   - [NSScrubber.SetSelectedIndex]
//
// # Inserting, moving, and deleting items
//
//   - [NSScrubber.InsertItemsAtIndexes]: Inserts new items at the specified indexes into the scrubber.
//   - [NSScrubber.MoveItemAtIndexToIndex]: Moves an item from one index to another in the scrubber.
//   - [NSScrubber.RemoveItemsAtIndexes]: Removes the items at the specified indexes from the scrubber.
//
// # Animating multiple changes to the scrubber
//
//   - [NSScrubber.PerformSequentialBatchUpdates]: Combines multiple scrubber content updates into a single action.
//
// # Scrolling items
//
//   - [NSScrubber.ScrollItemAtIndexToAlignment]: Scrolls an item to a specified alignment within the scrubber.
//
// # Locating items in the scrubber
//
//   - [NSScrubber.ItemViewForItemAtIndex]: Returns the view for the item at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber
type NSScrubber struct {
	NSView
}

// NSScrubberFromID constructs a [NSScrubber] from an objc.ID.
//
// A customizable item picker control for the Touch Bar.
func NSScrubberFromID(id objc.ID) NSScrubber {
	return NSScrubber{NSView: NSViewFromID(id)}
}
// NOTE: NSScrubber adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubber] class.
//
// # Configuring the scrubber
//
//   - [INSScrubber.DataSource]: The object that provides the data for the scrubber.
//   - [INSScrubber.SetDataSource]
//   - [INSScrubber.Delegate]: The object that acts as the delegate of the scrubber.
//   - [INSScrubber.SetDelegate]
//
// # Creating scrubber items
//
//   - [INSScrubber.RegisterClassForItemIdentifier]: Registers a class for the scrubber to use when it creates new items.
//   - [INSScrubber.RegisterNibForItemIdentifier]: Registers a nib file for the scrubber to use when it creates new items in the scrubber.
//   - [INSScrubber.MakeItemWithIdentifierOwner]: Creates or returns a reusable item object with the specified identifier.
//
// # Changing the layout
//
//   - [INSScrubber.ScrubberLayout]: An object used to describe the layout of items within the scrubber.
//   - [INSScrubber.SetScrubberLayout]
//   - [INSScrubber.Mode]: A setting that determines whether interaction with the scrubber is fixed or free.
//   - [INSScrubber.SetMode]
//   - [INSScrubber.ItemAlignment]: A setting that specifies the snapping behavior of items in the scrubber.
//   - [INSScrubber.SetItemAlignment]
//   - [INSScrubber.Continuous]: A Boolean value that, together with the [mode](<doc://com.apple.appkit/documentation/AppKit/NSScrubber/mode-swift.property>) property, determines scrubber interaction style.
//   - [INSScrubber.SetContinuous]
//
// # Configuring the scrubber’s appearance
//
//   - [INSScrubber.BackgroundColor]: The color displayed behind the scrubber content.
//   - [INSScrubber.SetBackgroundColor]
//   - [INSScrubber.BackgroundView]: A view that is displayed behind the scrubber content.
//   - [INSScrubber.SetBackgroundView]
//   - [INSScrubber.ShowsAdditionalContentIndicators]: A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
//   - [INSScrubber.SetShowsAdditionalContentIndicators]
//   - [INSScrubber.ShowsArrowButtons]: A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
//   - [INSScrubber.SetShowsArrowButtons]
//
// # Configuring the selection appearance
//
//   - [INSScrubber.FloatsSelectionViews]: A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
//   - [INSScrubber.SetFloatsSelectionViews]
//   - [INSScrubber.SelectionOverlayStyle]: The style overlaid on selected items.
//   - [INSScrubber.SetSelectionOverlayStyle]
//   - [INSScrubber.SelectionBackgroundStyle]: The style applied to the background of selected items.
//   - [INSScrubber.SetSelectionBackgroundStyle]
//
// # Reloading content
//
//   - [INSScrubber.ReloadData]: Reloads the content of the entire scrubber, and deselects the currently selected item.
//   - [INSScrubber.ReloadItemsAtIndexes]: Reloads the items at the specified indexes.
//
// # Getting the state of the scrubber
//
//   - [INSScrubber.NumberOfItems]: The number of items represented by the scrubber.
//   - [INSScrubber.HighlightedIndex]: The index of the highlighted item in the scrubber.
//   - [INSScrubber.SelectedIndex]: The index of the selected item in the scrubber.
//   - [INSScrubber.SetSelectedIndex]
//
// # Inserting, moving, and deleting items
//
//   - [INSScrubber.InsertItemsAtIndexes]: Inserts new items at the specified indexes into the scrubber.
//   - [INSScrubber.MoveItemAtIndexToIndex]: Moves an item from one index to another in the scrubber.
//   - [INSScrubber.RemoveItemsAtIndexes]: Removes the items at the specified indexes from the scrubber.
//
// # Animating multiple changes to the scrubber
//
//   - [INSScrubber.PerformSequentialBatchUpdates]: Combines multiple scrubber content updates into a single action.
//
// # Scrolling items
//
//   - [INSScrubber.ScrollItemAtIndexToAlignment]: Scrolls an item to a specified alignment within the scrubber.
//
// # Locating items in the scrubber
//
//   - [INSScrubber.ItemViewForItemAtIndex]: Returns the view for the item at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber
type INSScrubber interface {
	INSView

	// Topic: Configuring the scrubber

	// The object that provides the data for the scrubber.
	DataSource() NSScrubberDataSource
	SetDataSource(value NSScrubberDataSource)
	// The object that acts as the delegate of the scrubber.
	Delegate() NSScrubberDelegate
	SetDelegate(value NSScrubberDelegate)

	// Topic: Creating scrubber items

	// Registers a class for the scrubber to use when it creates new items.
	RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier NSUserInterfaceItemIdentifier)
	// Registers a nib file for the scrubber to use when it creates new items in the scrubber.
	RegisterNibForItemIdentifier(nib INSNib, itemIdentifier NSUserInterfaceItemIdentifier)
	// Creates or returns a reusable item object with the specified identifier.
	MakeItemWithIdentifierOwner(itemIdentifier NSUserInterfaceItemIdentifier, owner objectivec.IObject) INSScrubberItemView

	// Topic: Changing the layout

	// An object used to describe the layout of items within the scrubber.
	ScrubberLayout() INSScrubberLayout
	SetScrubberLayout(value INSScrubberLayout)
	// A setting that determines whether interaction with the scrubber is fixed or free.
	Mode() NSScrubberMode
	SetMode(value NSScrubberMode)
	// A setting that specifies the snapping behavior of items in the scrubber.
	ItemAlignment() NSScrubberAlignment
	SetItemAlignment(value NSScrubberAlignment)
	// A Boolean value that, together with the [mode](<doc://com.apple.appkit/documentation/AppKit/NSScrubber/mode-swift.property>) property, determines scrubber interaction style.
	Continuous() bool
	SetContinuous(value bool)

	// Topic: Configuring the scrubber’s appearance

	// The color displayed behind the scrubber content.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A view that is displayed behind the scrubber content.
	BackgroundView() INSView
	SetBackgroundView(value INSView)
	// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges.
	ShowsAdditionalContentIndicators() bool
	SetShowsAdditionalContentIndicators(value bool)
	// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber.
	ShowsArrowButtons() bool
	SetShowsArrowButtons(value bool)

	// Topic: Configuring the selection appearance

	// A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes.
	FloatsSelectionViews() bool
	SetFloatsSelectionViews(value bool)
	// The style overlaid on selected items.
	SelectionOverlayStyle() INSScrubberSelectionStyle
	SetSelectionOverlayStyle(value INSScrubberSelectionStyle)
	// The style applied to the background of selected items.
	SelectionBackgroundStyle() INSScrubberSelectionStyle
	SetSelectionBackgroundStyle(value INSScrubberSelectionStyle)

	// Topic: Reloading content

	// Reloads the content of the entire scrubber, and deselects the currently selected item.
	ReloadData()
	// Reloads the items at the specified indexes.
	ReloadItemsAtIndexes(indexes foundation.NSIndexSet)

	// Topic: Getting the state of the scrubber

	// The number of items represented by the scrubber.
	NumberOfItems() int
	// The index of the highlighted item in the scrubber.
	HighlightedIndex() int
	// The index of the selected item in the scrubber.
	SelectedIndex() int
	SetSelectedIndex(value int)

	// Topic: Inserting, moving, and deleting items

	// Inserts new items at the specified indexes into the scrubber.
	InsertItemsAtIndexes(indexes foundation.NSIndexSet)
	// Moves an item from one index to another in the scrubber.
	MoveItemAtIndexToIndex(oldIndex int, newIndex int)
	// Removes the items at the specified indexes from the scrubber.
	RemoveItemsAtIndexes(indexes foundation.NSIndexSet)

	// Topic: Animating multiple changes to the scrubber

	// Combines multiple scrubber content updates into a single action.
	PerformSequentialBatchUpdates(updateBlock VoidHandler)

	// Topic: Scrolling items

	// Scrolls an item to a specified alignment within the scrubber.
	ScrollItemAtIndexToAlignment(index int, alignment NSScrubberAlignment)

	// Topic: Locating items in the scrubber

	// Returns the view for the item at the specified index.
	ItemViewForItemAtIndex(index int) INSScrubberItemView

	// The item’s alpha value.
	Alpha() float64
	SetAlpha(value float64)
	// The alignment of the image within the scrubber item.
	ImageAlignment() NSImageAlignment
	SetImageAlignment(value NSImageAlignment)
	// The image view that the scrubber item uses to display its image.
	ImageView() INSImageView
	SetImageView(value INSImageView)
	// The index of the scrubber item that is represented by the item’s layout attributes.
	ItemIndex() int
	SetItemIndex(value int)
	// The size required to contain all elements within the scrubber.
	ScrubberContentSize() corefoundation.CGSize
	SetScrubberContentSize(value corefoundation.CGSize)
	// Determines whether the scrubber should refresh its layout when an item is highlighted.
	ShouldInvalidateLayoutForHighlightChange() bool
	SetShouldInvalidateLayoutForHighlightChange(value bool)
	// Determines whether the scrubber should refresh its layout when the selection changes.
	ShouldInvalidateLayoutForSelectionChange() bool
	SetShouldInvalidateLayoutForSelectionChange(value bool)
	// The text field that the scrubber item uses to display its text.
	TextField() INSTextField
	SetTextField(value INSTextField)
}

// Init initializes the instance.
func (s NSScrubber) Init() NSScrubber {
	rv := objc.Send[NSScrubber](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubber) Autorelease() NSScrubber {
	rv := objc.Send[NSScrubber](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubber creates a new NSScrubber instance.
func NewNSScrubber() NSScrubber {
	class := getNSScrubberClass()
	rv := objc.Send[NSScrubber](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a newly allocated scrubber object from a storyboard
// or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/init(coder:)
func NewScrubberWithCoder(coder foundation.INSCoder) NSScrubber {
	instance := getNSScrubberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberFromID(rv)
}

// Initializes and returns a newly allocated scrubber object with the
// specified frame rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/init(frame:)
func NewScrubberWithFrame(frameRect corefoundation.CGRect) NSScrubber {
	instance := getNSScrubberClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrubberFromID(rv)
}

// Registers a class for the scrubber to use when it creates new items.
//
// itemViewClass: A class to use for creating items. The class must be descended from
// [NSScrubberItemView]. Specify `nil` to unregister a previously registered
// class.
//
// itemIdentifier: The string that identifies the type of item. You use this string later when
// requesting new item views. The string must be unique among the other
// registered item view classes of this scrubber. This parameter must not be
// an empty string or `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-2rb69
func (s NSScrubber) RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](s.ID, objc.Sel("registerClass:forItemIdentifier:"), itemViewClass, objc.String(string(itemIdentifier)))
}
// Registers a nib file for the scrubber to use when it creates new items in
// the scrubber.
//
// nib: The nib object containing the item object. The nib file must contain
// exactly one top-level [NSScrubberItemView] object. You can use a custom
// subclass when configuring the object in the nib file. Specify `nil` to
// unregister a previously registered file.
//
// itemIdentifier: The string that identifies the type of items. You use this string later
// when requesting new items. The string must be unique among the other
// registered item view classes of this scrubber. This parameter must not be
// an empty string or `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-6jye0
func (s NSScrubber) RegisterNibForItemIdentifier(nib INSNib, itemIdentifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](s.ID, objc.Sel("registerNib:forItemIdentifier:"), nib, objc.String(string(itemIdentifier)))
}
// Creates or returns a reusable item object with the specified identifier.
//
// itemIdentifier: The string that identifies the type of item you want. This is the
// identifier you specified when registering the item view. The parameter must
// not be `nil`.
//
// owner: The owner of this item. If the scrubber item is loaded from a nib then this
// object is set as the nib’s File’s Owner object. Set this parameter to
// `nil` for scrubber items loaded from classes.
//
// # Return Value
// 
// A valid [NSScrubberItemView] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/makeItem(withIdentifier:owner:)
func (s NSScrubber) MakeItemWithIdentifierOwner(itemIdentifier NSUserInterfaceItemIdentifier, owner objectivec.IObject) INSScrubberItemView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeItemWithIdentifier:owner:"), objc.String(string(itemIdentifier)), owner)
	return NSScrubberItemViewFromID(rv)
}
// Reloads the content of the entire scrubber, and deselects the currently
// selected item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/reloadData()
func (s NSScrubber) ReloadData() {
	objc.Send[objc.ID](s.ID, objc.Sel("reloadData"))
}
// Reloads the items at the specified indexes.
//
// indexes: The indexes of the items to reload, provided in an index set ([IndexSet]).
// //
// [IndexSet]: https://developer.apple.com/documentation/Foundation/IndexSet
//
// # Discussion
// 
// The scrubber makes requests for updated views from the data source and then
// animates their display.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/reloadItems(at:)
func (s NSScrubber) ReloadItemsAtIndexes(indexes foundation.NSIndexSet) {
	objc.Send[objc.ID](s.ID, objc.Sel("reloadItemsAtIndexes:"), indexes)
}
// Inserts new items at the specified indexes into the scrubber.
//
// indexes: An index set ([IndexSet]) of the indexes of the items to insert.
// //
// [IndexSet]: https://developer.apple.com/documentation/Foundation/IndexSet
//
// # Discussion
// 
// The scrubber requests the view for each index from its data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/insertItems(at:)
func (s NSScrubber) InsertItemsAtIndexes(indexes foundation.NSIndexSet) {
	objc.Send[objc.ID](s.ID, objc.Sel("insertItemsAtIndexes:"), indexes)
}
// Moves an item from one index to another in the scrubber.
//
// oldIndex: The index of the item that you want to move.
//
// newIndex: The index of the item’s new location.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/moveItem(at:to:)
func (s NSScrubber) MoveItemAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](s.ID, objc.Sel("moveItemAtIndex:toIndex:"), oldIndex, newIndex)
}
// Removes the items at the specified indexes from the scrubber.
//
// indexes: An index set ([IndexSet]) of the indexes of the items to remove.
// //
// [IndexSet]: https://developer.apple.com/documentation/Foundation/IndexSet
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/removeItems(at:)
func (s NSScrubber) RemoveItemsAtIndexes(indexes foundation.NSIndexSet) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeItemsAtIndexes:"), indexes)
}
// Combines multiple scrubber content updates into a single action.
//
// updateBlock: A block that represents the combination of insertion, removal, moving, and
// reloading instructions that should be performed simultaneously.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/performSequentialBatchUpdates(_:)
func (s NSScrubber) PerformSequentialBatchUpdates(updateBlock VoidHandler) {
_block0, _ := NewVoidBlock(updateBlock)
	objc.Send[objc.ID](s.ID, objc.Sel("performSequentialBatchUpdates:"), _block0)
}
// Scrolls an item to a specified alignment within the scrubber.
//
// index: The index of the item to be scrolled. Indexes range between `0` and `n-1`,
// where `n` is the number of items displayed in the scrubber.
//
// alignment: The position the item should be scrolled to. For possible values, see
// [NSScrubber.Alignment]. If [NSScrubber.Alignment.none], the scrubber
// scrolls the minimum distance required to make the item visible.
// //
// [NSScrubber.Alignment.none]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/none
// [NSScrubber.Alignment]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment
//
// # Discussion
// 
// To animate the scroll, call this method on the animator proxy. See
// [NSAnimatablePropertyContainer] for details.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/scrollItem(at:to:)
func (s NSScrubber) ScrollItemAtIndexToAlignment(index int, alignment NSScrubberAlignment) {
	objc.Send[objc.ID](s.ID, objc.Sel("scrollItemAtIndex:toAlignment:"), index, alignment)
}
// Returns the view for the item at the specified index.
//
// index: The index of the item whose view you want.
//
// # Return Value
// 
// The view for the specified index or `nil` if the item is not currently
// visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/itemViewForItem(at:)
func (s NSScrubber) ItemViewForItemAtIndex(index int) INSScrubberItemView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("itemViewForItemAtIndex:"), index)
	return NSScrubberItemViewFromID(rv)
}

// The object that provides the data for the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/dataSource
func (s NSScrubber) DataSource() NSScrubberDataSource {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dataSource"))
	return NSScrubberDataSourceObjectFromID(rv)
}
func (s NSScrubber) SetDataSource(value NSScrubberDataSource) {
	objc.Send[struct{}](s.ID, objc.Sel("setDataSource:"), value)
}
// The object that acts as the delegate of the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/delegate
func (s NSScrubber) Delegate() NSScrubberDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSScrubberDelegateObjectFromID(rv)
}
func (s NSScrubber) SetDelegate(value NSScrubberDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
// An object used to describe the layout of items within the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/scrubberLayout
func (s NSScrubber) ScrubberLayout() INSScrubberLayout {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("scrubberLayout"))
	return NSScrubberLayoutFromID(objc.ID(rv))
}
func (s NSScrubber) SetScrubberLayout(value INSScrubberLayout) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrubberLayout:"), value)
}
// A setting that determines whether interaction with the scrubber is fixed or
// free.
//
// # Discussion
// 
// When set to [NSScrubber.Mode.fixed], the scrubber’s content doesn’t
// scroll as the user pans. Instead, the element under the user’s finger is
// highlighted. The highlighted item becomes selected when the user completes
// the pan gesture. The default value is [NSScrubber.Mode.fixed].
// 
// When [Mode] is set to [NSScrubber.Mode.free], panning over the scrubber
// scrolls the scrubber’s content. A user selects items by tapping on them.
//
// [NSScrubber.Mode.fixed]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/fixed
// [NSScrubber.Mode.free]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/free
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/mode-swift.property
func (s NSScrubber) Mode() NSScrubberMode {
	rv := objc.Send[NSScrubberMode](s.ID, objc.Sel("mode"))
	return NSScrubberMode(rv)
}
func (s NSScrubber) SetMode(value NSScrubberMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setMode:"), value)
}
// A setting that specifies the snapping behavior of items in the scrubber.
//
// # Discussion
// 
// Set the value to something other than [NSScrubber.Alignment.none] to ensure
// that an item is aligned with the specified position when the scrubber comes
// to rest.
// 
// The default value is [NSScrubber.Alignment.none].
//
// [NSScrubber.Alignment.none]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/none
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/itemAlignment
func (s NSScrubber) ItemAlignment() NSScrubberAlignment {
	rv := objc.Send[NSScrubberAlignment](s.ID, objc.Sel("itemAlignment"))
	return NSScrubberAlignment(rv)
}
func (s NSScrubber) SetItemAlignment(value NSScrubberAlignment) {
	objc.Send[struct{}](s.ID, objc.Sel("setItemAlignment:"), value)
}
// A Boolean value that, together with the [Mode] property, determines
// scrubber interaction style.
//
// # Discussion
// 
// For details on how to choose the right [Continuous] value for your app, see
// [NSScrubber].
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/isContinuous
func (s NSScrubber) Continuous() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isContinuous"))
	return rv
}
func (s NSScrubber) SetContinuous(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setContinuous:"), value)
}
// The color displayed behind the scrubber content.
//
// # Discussion
// 
// The value of this property is ignored if the value of the [BackgroundView]
// property is non-`nil`. The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/backgroundColor
func (s NSScrubber) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (s NSScrubber) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setBackgroundColor:"), value)
}
// A view that is displayed behind the scrubber content.
//
// # Discussion
// 
// The scrubber manages the layout of the background view to match the size of
// the content area. If this property is non-`nil`, the value of the
// [BackgroundColor] property is ignored.
// 
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/backgroundView
func (s NSScrubber) BackgroundView() INSView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("backgroundView"))
	return NSViewFromID(objc.ID(rv))
}
func (s NSScrubber) SetBackgroundView(value INSView) {
	objc.Send[struct{}](s.ID, objc.Sel("setBackgroundView:"), value)
}
// A Boolean value that specifies whether the scrubber should display the
// existence of additional items beyond the leading and trailing edges.
//
// # Discussion
// 
// When [true], the scrubber uses a fade effect to indicate that there is
// additional content the user can scroll to. The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/showsAdditionalContentIndicators
func (s NSScrubber) ShowsAdditionalContentIndicators() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsAdditionalContentIndicators"))
	return rv
}
func (s NSScrubber) SetShowsAdditionalContentIndicators(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsAdditionalContentIndicators:"), value)
}
// A Boolean value that specifies whether arrow buttons should be displayed at
// the leading and trailing edges of the scrubber.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/showsArrowButtons
func (s NSScrubber) ShowsArrowButtons() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsArrowButtons"))
	return rv
}
func (s NSScrubber) SetShowsArrowButtons(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsArrowButtons:"), value)
}
// A Boolean value that determines the behavior of the item selection
// decorations as the scrubber’s selection changes.
//
// # Discussion
// 
// When scrubber items are selected, they are decorated with background and
// overlay views, as determined by the [SelectionBackgroundStyle] and
// [SelectionOverlayStyle] properties.
// 
// As the selection changes, the behavior of these selection decoration views
// is determined by the [FloatsSelectionViews] property, as follows:
// 
// - [true] The overlay and background views float smoothly between the
// previously selected item and the newly selected item. - [false] The overlay
// and background views cross-fade from the previously selected item to the
// newly selected item.
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/floatsSelectionViews
func (s NSScrubber) FloatsSelectionViews() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("floatsSelectionViews"))
	return rv
}
func (s NSScrubber) SetFloatsSelectionViews(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setFloatsSelectionViews:"), value)
}
// The style overlaid on selected items.
//
// # Discussion
// 
// The default value is `nil`, which specifies no overlay decoration.
// 
// You can either choose from one of the built-in selection styles
// ([OutlineOverlayStyle] or [RoundedBackgroundStyle]), or you can subclass
// [NSScrubberSelectionStyle] to create your own custom selection style.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/selectionOverlayStyle
func (s NSScrubber) SelectionOverlayStyle() INSScrubberSelectionStyle {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectionOverlayStyle"))
	return NSScrubberSelectionStyleFromID(objc.ID(rv))
}
func (s NSScrubber) SetSelectionOverlayStyle(value INSScrubberSelectionStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelectionOverlayStyle:"), value)
}
// The style applied to the background of selected items.
//
// # Discussion
// 
// The default value is `nil`, which specifies no background decoration.
// 
// You can either choose from one of the built-in selection styles
// ([OutlineOverlayStyle] or [RoundedBackgroundStyle]), or you can subclass
// [NSScrubberSelectionStyle] to create your own custom selection style.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/selectionBackgroundStyle
func (s NSScrubber) SelectionBackgroundStyle() INSScrubberSelectionStyle {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("selectionBackgroundStyle"))
	return NSScrubberSelectionStyleFromID(objc.ID(rv))
}
func (s NSScrubber) SetSelectionBackgroundStyle(value INSScrubberSelectionStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelectionBackgroundStyle:"), value)
}
// The number of items represented by the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/numberOfItems
func (s NSScrubber) NumberOfItems() int {
	rv := objc.Send[int](s.ID, objc.Sel("numberOfItems"))
	return rv
}
// The index of the highlighted item in the scrubber.
//
// # Discussion
// 
// If no item is highlighted, the value of this property is `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/highlightedIndex
func (s NSScrubber) HighlightedIndex() int {
	rv := objc.Send[int](s.ID, objc.Sel("highlightedIndex"))
	return rv
}
// The index of the selected item in the scrubber.
//
// # Discussion
// 
// If no item is selected, the value of this property is `-1`. If you set this
// property through the scrubber’s animator proxy, the selection change
// animates.
// 
// To use a scrubber’s animator proxy when changing the selected item,
// employ the [NSAnimatablePropertyContainer] protocol, using code like this:
// `scrubber.AnimatorXCUIElementTypeSelectedIndex() = 123`
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubber/selectedIndex
func (s NSScrubber) SelectedIndex() int {
	rv := objc.Send[int](s.ID, objc.Sel("selectedIndex"))
	return rv
}
func (s NSScrubber) SetSelectedIndex(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelectedIndex:"), value)
}
// The item’s alpha value.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/alpha
func (s NSScrubber) Alpha() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("alpha"))
	return rv
}
func (s NSScrubber) SetAlpha(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAlpha:"), value)
}
// The alignment of the image within the scrubber item.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imagealignment
func (s NSScrubber) ImageAlignment() NSImageAlignment {
	rv := objc.Send[NSImageAlignment](s.ID, objc.Sel("imageAlignment"))
	return NSImageAlignment(rv)
}
func (s NSScrubber) SetImageAlignment(value NSImageAlignment) {
	objc.Send[struct{}](s.ID, objc.Sel("setImageAlignment:"), value)
}
// The image view that the scrubber item uses to display its image.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberimageitemview/imageview
func (s NSScrubber) ImageView() INSImageView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("imageView"))
	return NSImageViewFromID(objc.ID(rv))
}
func (s NSScrubber) SetImageView(value INSImageView) {
	objc.Send[struct{}](s.ID, objc.Sel("setImageView:"), value)
}
// The index of the scrubber item that is represented by the item’s layout
// attributes.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberlayoutattributes/itemindex
func (s NSScrubber) ItemIndex() int {
	rv := objc.Send[int](s.ID, objc.Sel("itemIndex"))
	return rv
}
func (s NSScrubber) SetItemIndex(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setItemIndex:"), value)
}
// The size required to contain all elements within the scrubber.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberlayout/scrubbercontentsize
func (s NSScrubber) ScrubberContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("scrubberContentSize"))
	return corefoundation.CGSize(rv)
}
func (s NSScrubber) SetScrubberContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrubberContentSize:"), value)
}
// Determines whether the scrubber should refresh its layout when an item is
// highlighted.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforhighlightchange
func (s NSScrubber) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}
func (s NSScrubber) SetShouldInvalidateLayoutForHighlightChange(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShouldInvalidateLayoutForHighlightChange:"), value)
}
// Determines whether the scrubber should refresh its layout when the
// selection changes.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubberlayout/shouldinvalidatelayoutforselectionchange
func (s NSScrubber) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}
func (s NSScrubber) SetShouldInvalidateLayoutForSelectionChange(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShouldInvalidateLayoutForSelectionChange:"), value)
}
// The text field that the scrubber item uses to display its text.
//
// See: https://developer.apple.com/documentation/appkit/nsscrubbertextitemview/textfield
func (s NSScrubber) TextField() INSTextField {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("textField"))
	return NSTextFieldFromID(objc.ID(rv))
}
func (s NSScrubber) SetTextField(value INSTextField) {
	objc.Send[struct{}](s.ID, objc.Sel("setTextField:"), value)
}

// PerformSequentialBatchUpdatesSync is a synchronous wrapper around [NSScrubber.PerformSequentialBatchUpdates].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSScrubber) PerformSequentialBatchUpdatesSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s.PerformSequentialBatchUpdates(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

