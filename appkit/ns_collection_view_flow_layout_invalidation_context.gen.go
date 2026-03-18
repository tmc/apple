// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSCollectionViewFlowLayoutInvalidationContext] class.
var (
	_NSCollectionViewFlowLayoutInvalidationContextClass     NSCollectionViewFlowLayoutInvalidationContextClass
	_NSCollectionViewFlowLayoutInvalidationContextClassOnce sync.Once
)

func getNSCollectionViewFlowLayoutInvalidationContextClass() NSCollectionViewFlowLayoutInvalidationContextClass {
	_NSCollectionViewFlowLayoutInvalidationContextClassOnce.Do(func() {
		_NSCollectionViewFlowLayoutInvalidationContextClass = NSCollectionViewFlowLayoutInvalidationContextClass{class: objc.GetClass("NSCollectionViewFlowLayoutInvalidationContext")}
	})
	return _NSCollectionViewFlowLayoutInvalidationContextClass
}

// GetNSCollectionViewFlowLayoutInvalidationContextClass returns the class object for NSCollectionViewFlowLayoutInvalidationContext.
func GetNSCollectionViewFlowLayoutInvalidationContextClass() NSCollectionViewFlowLayoutInvalidationContextClass {
	return getNSCollectionViewFlowLayoutInvalidationContextClass()
}

type NSCollectionViewFlowLayoutInvalidationContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewFlowLayoutInvalidationContextClass) Alloc() NSCollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[NSCollectionViewFlowLayoutInvalidationContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that identifies the portions of a flow layout object that need to
// be updated.
//
// # Overview
// 
// Layout objects use invalidation contexts to optimize the layout process and
// avoid unnecessary work. You use this class to specify whether the
// [NSCollectionViewFlowLayout] object should fetch new size information from
// its delegate. You can also prevent the flow layout object from updating its
// layout information altogether.
// 
// When you want to invalidate your flow layout object, call the
// [NSCollectionViewFlowLayoutInvalidationContext.InvalidationContextClass] method of your layout object and instantiate the
// resulting class. (The implementation of that method in
// [NSCollectionViewFlowLayout] returns this class.) After instantiating this
// class, set the properties to appropriate values and pass the object to the
// [InvalidateLayoutWithContext] method of the layout object.
//
// # Invalidating the Flow Layout
//
//   - [NSCollectionViewFlowLayoutInvalidationContext.InvalidateFlowLayoutAttributes]: A Boolean value indicating whether the flow layout object should invalidate its current attributes.
//   - [NSCollectionViewFlowLayoutInvalidationContext.SetInvalidateFlowLayoutAttributes]
//   - [NSCollectionViewFlowLayoutInvalidationContext.InvalidateFlowLayoutDelegateMetrics]: A Boolean value indicating whether the flow layout object should fetch new size information from its delegate.
//   - [NSCollectionViewFlowLayoutInvalidationContext.SetInvalidateFlowLayoutDelegateMetrics]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext
type NSCollectionViewFlowLayoutInvalidationContext struct {
	NSCollectionViewLayoutInvalidationContext
}

// NSCollectionViewFlowLayoutInvalidationContextFromID constructs a [NSCollectionViewFlowLayoutInvalidationContext] from an objc.ID.
//
// An object that identifies the portions of a flow layout object that need to
// be updated.
func NSCollectionViewFlowLayoutInvalidationContextFromID(id objc.ID) NSCollectionViewFlowLayoutInvalidationContext {
	return NSCollectionViewFlowLayoutInvalidationContext{NSCollectionViewLayoutInvalidationContext: NSCollectionViewLayoutInvalidationContextFromID(id)}
}
// NOTE: NSCollectionViewFlowLayoutInvalidationContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewFlowLayoutInvalidationContext] class.
//
// # Invalidating the Flow Layout
//
//   - [INSCollectionViewFlowLayoutInvalidationContext.InvalidateFlowLayoutAttributes]: A Boolean value indicating whether the flow layout object should invalidate its current attributes.
//   - [INSCollectionViewFlowLayoutInvalidationContext.SetInvalidateFlowLayoutAttributes]
//   - [INSCollectionViewFlowLayoutInvalidationContext.InvalidateFlowLayoutDelegateMetrics]: A Boolean value indicating whether the flow layout object should fetch new size information from its delegate.
//   - [INSCollectionViewFlowLayoutInvalidationContext.SetInvalidateFlowLayoutDelegateMetrics]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext
type INSCollectionViewFlowLayoutInvalidationContext interface {
	INSCollectionViewLayoutInvalidationContext

	// Topic: Invalidating the Flow Layout

	// A Boolean value indicating whether the flow layout object should invalidate its current attributes.
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	// A Boolean value indicating whether the flow layout object should fetch new size information from its delegate.
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
}

// Init initializes the instance.
func (c NSCollectionViewFlowLayoutInvalidationContext) Init() NSCollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[NSCollectionViewFlowLayoutInvalidationContext](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewFlowLayoutInvalidationContext) Autorelease() NSCollectionViewFlowLayoutInvalidationContext {
	rv := objc.Send[NSCollectionViewFlowLayoutInvalidationContext](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewFlowLayoutInvalidationContext creates a new NSCollectionViewFlowLayoutInvalidationContext instance.
func NewNSCollectionViewFlowLayoutInvalidationContext() NSCollectionViewFlowLayoutInvalidationContext {
	class := getNSCollectionViewFlowLayoutInvalidationContextClass()
	rv := objc.Send[NSCollectionViewFlowLayoutInvalidationContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value indicating whether the flow layout object should invalidate
// its current attributes.
//
// # Discussion
// 
// Setting this property to [false] tells the flow layout object to keep its
// existing layout information, effectively stopping the invalidation process.
// Typically, you set this property to [false] only if you subclass
// [NSCollectionViewFlowLayout] and update changed layout information
// directly.
// 
// The default value of this property is [true], which causes the flow layout
// object to throw out its existing layout information and recompute it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext/invalidateFlowLayoutAttributes
func (c NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("invalidateFlowLayoutAttributes"))
	return rv
}
func (c NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setInvalidateFlowLayoutAttributes:"), value)
}

// A Boolean value indicating whether the flow layout object should fetch new
// size information from its delegate.
//
// # Discussion
// 
// As part of the invalidation process, the flow layout object normally asks
// its delegate to provide size information for the items in the flow layout.
// This behavior is necessary when the size of the items can change because it
// ensures that the corresponding layout attributes are always updated.
// However, if you know that the size of items has not changed, you can set
// this property to [false]. Doing so causes the flow layout to use its
// existing size information rather than querying the delegate, which saves
// time.
// 
// The default value of this property is [true], which causes the flow layout
// object to query the delegate for new size information.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayoutInvalidationContext/invalidateFlowLayoutDelegateMetrics
func (c NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("invalidateFlowLayoutDelegateMetrics"))
	return rv
}
func (c NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setInvalidateFlowLayoutDelegateMetrics:"), value)
}

// Returns the class to use when creating an invalidation context object for
// the layout.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionviewlayout/invalidationcontextclass
func (_NSCollectionViewFlowLayoutInvalidationContextClass NSCollectionViewFlowLayoutInvalidationContextClass) InvalidationContextClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSCollectionViewFlowLayoutInvalidationContextClass.class), objc.Sel("invalidationContextClass"))
	return rv
}
func (_NSCollectionViewFlowLayoutInvalidationContextClass NSCollectionViewFlowLayoutInvalidationContextClass) SetInvalidationContextClass(value objc.Class) {
	objc.Send[struct{}](objc.ID(_NSCollectionViewFlowLayoutInvalidationContextClass.class), objc.Sel("setInvalidationContextClass:"), value)
}

