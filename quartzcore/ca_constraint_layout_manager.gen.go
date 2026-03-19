// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAConstraintLayoutManager] class.
var (
	_CAConstraintLayoutManagerClass     CAConstraintLayoutManagerClass
	_CAConstraintLayoutManagerClassOnce sync.Once
)

func getCAConstraintLayoutManagerClass() CAConstraintLayoutManagerClass {
	_CAConstraintLayoutManagerClassOnce.Do(func() {
		_CAConstraintLayoutManagerClass = CAConstraintLayoutManagerClass{class: objc.GetClass("CAConstraintLayoutManager")}
	})
	return _CAConstraintLayoutManagerClass
}

// GetCAConstraintLayoutManagerClass returns the class object for CAConstraintLayoutManager.
func GetCAConstraintLayoutManagerClass() CAConstraintLayoutManagerClass {
	return getCAConstraintLayoutManagerClass()
}

type CAConstraintLayoutManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAConstraintLayoutManagerClass) Alloc() CAConstraintLayoutManager {
	rv := objc.Send[CAConstraintLayoutManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides a constraint-based layout manager.
//
// # Overview
// 
// You use the shared instance of this object by assigning it to the
// [CAConstraintLayoutManager.LayoutManager] property of any layer objects to which you have added
// constraints. During a layout update, Core Animation uses the layout manager
// to update the size and position of the sublayers based on the registered
// set of constraints.
// 
// Constraints let you define a set of geometric relationships between a layer
// and its sibling layers or between a layer and its superlayer. These
// relationships are expressed using constraint objects, which are instances
// of the [CAConstraint] class. When creating constraints, you can reference a
// layer by name using that object’s [CAConstraintLayoutManager.Name] property. You can also use the
// special name `superlayer` to refer to the layer’s superlayer.
// 
// The following example shows how you can use [CAConstraintLayoutManager] to
// create a layer containing two constrained sublayers: `leftLayer` and
// `rightLayer`. A series of [CAConstraint] objects are created so that the
// sublayers match their superlayer’s height and are half of its width.
// `leftConstraint` matches the [CAConstraintAttribute.minX] attribute and
// `rightConstraint` matches the [CAConstraintAttribute.maxX] attribute.
// 
// The end result is that the two sublayers are always laid out so that
// `leftLayer` fills the left half of `layer` and `rightLayer` fills the right
// half of layer.
// 
// This class is not meant to be subclassed.
//
// [CAConstraintAttribute.maxX]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/maxX
// [CAConstraintAttribute.minX]: https://developer.apple.com/documentation/QuartzCore/CAConstraintAttribute/minX
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager
type CAConstraintLayoutManager struct {
	objectivec.Object
}

// CAConstraintLayoutManagerFromID constructs a [CAConstraintLayoutManager] from an objc.ID.
//
// An object that provides a constraint-based layout manager.
func CAConstraintLayoutManagerFromID(id objc.ID) CAConstraintLayoutManager {
	return CAConstraintLayoutManager{objectivec.Object{ID: id}}
}
// NOTE: CAConstraintLayoutManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAConstraintLayoutManager] class.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager
type ICAConstraintLayoutManager interface {
	objectivec.IObject
	CALayoutManager

	// The object responsible for laying out the layer’s sublayers.
	LayoutManager() CALayoutManager
	SetLayoutManager(value CALayoutManager)
	// The name of the receiver.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (c CAConstraintLayoutManager) Init() CAConstraintLayoutManager {
	rv := objc.Send[CAConstraintLayoutManager](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CAConstraintLayoutManager) Autorelease() CAConstraintLayoutManager {
	rv := objc.Send[CAConstraintLayoutManager](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAConstraintLayoutManager creates a new CAConstraintLayoutManager instance.
func NewCAConstraintLayoutManager() CAConstraintLayoutManager {
	class := getCAConstraintLayoutManagerClass()
	rv := objc.Send[CAConstraintLayoutManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Invalidates the layout of a layer so it knows to refresh its content on the
// next frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager/invalidateLayout(of:)
func (c CAConstraintLayoutManager) InvalidateLayoutOfLayer(layer ICALayer) {
	objc.Send[objc.ID](c.ID, objc.Sel("invalidateLayoutOfLayer:"), layer)
}
// Override to customize layout of sublayers whenever the layer needs
// redrawing.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager/layoutSublayers(of:)
func (c CAConstraintLayoutManager) LayoutSublayersOfLayer(layer ICALayer) {
	objc.Send[objc.ID](c.ID, objc.Sel("layoutSublayersOfLayer:"), layer)
}
// Override to customize layer size.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayoutManager/preferredSize(of:)
func (c CAConstraintLayoutManager) PreferredSizeOfLayer(layer ICALayer) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("preferredSizeOfLayer:"), layer)
	return corefoundation.CGSize(rv)
}

// The object responsible for laying out the layer’s sublayers.
//
// See: https://developer.apple.com/documentation/quartzcore/calayer/layoutmanager
func (c CAConstraintLayoutManager) LayoutManager() CALayoutManager {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutManager"))
	return CALayoutManagerObjectFromID(rv)
}
func (c CAConstraintLayoutManager) SetLayoutManager(value CALayoutManager) {
	objc.Send[struct{}](c.ID, objc.Sel("setLayoutManager:"), value)
}
// The name of the receiver.
//
// See: https://developer.apple.com/documentation/quartzcore/calayer/name
func (c CAConstraintLayoutManager) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (c CAConstraintLayoutManager) SetName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setName:"), objc.String(value))
}

			// Protocol methods for CALayoutManager
			

