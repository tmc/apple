// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CALayer] class.
var (
	_CALayerClass     CALayerClass
	_CALayerClassOnce sync.Once
)

func getCALayerClass() CALayerClass {
	_CALayerClassOnce.Do(func() {
		_CALayerClass = CALayerClass{class: objc.GetClass("CALayer")}
	})
	return _CALayerClass
}

// GetCALayerClass returns the class object for CALayer.
func GetCALayerClass() CALayerClass {
	return getCALayerClass()
}

type CALayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CALayerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CALayerClass) Alloc() CALayer {
	rv := objc.Send[CALayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages image-based content and allows you to perform
// animations on that content.
//
// # Overview
// 
// Layers are often used to provide the backing store for views but can also
// be used without a view to display content. A layer’s main job is to
// manage the visual content that you provide but the layer itself has visual
// attributes that can be set, such as a background color, border, and shadow.
// In addition to managing visual content, the layer also maintains
// information about the geometry of its content (such as its position, size,
// and transform) that is used to present that content onscreen. Modifying the
// properties of the layer is how you initiate animations on the layer’s
// content or geometry. A layer object encapsulates the duration and pacing of
// a layer and its animations by adopting the [CAMediaTiming] protocol, which
// defines the layer’s timing information.
// 
// If the layer object was created by a view, the view typically assigns
// itself as the layer’s delegate automatically, and you should not change
// that relationship. For layers you create yourself, you can assign a
// [CALayer.Delegate] object and use that object to provide the contents of the layer
// dynamically and perform other tasks. A layer may also have a layout manager
// object (assigned to the [CALayer.LayoutManager] property) to manage the layout of
// subviews separately.
//
// # Creating a layer
//
//   - [CALayer.InitWithLayer]: Override to copy or initialize custom fields of the specified layer.
//
// # Accessing related layer objects
//
//   - [CALayer.PresentationLayer]: Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
//   - [CALayer.ModelLayer]: Returns the model layer object associated with the receiver, if any.
//
// # Accessing the delegate
//
//   - [CALayer.Delegate]: The layer’s delegate object.
//   - [CALayer.SetDelegate]
//
// # Providing the layer’s content
//
//   - [CALayer.Contents]: An object that provides the contents of the layer. Animatable.
//   - [CALayer.SetContents]
//   - [CALayer.ContentsRect]: The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//   - [CALayer.SetContentsRect]
//   - [CALayer.ContentsCenter]: The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//   - [CALayer.SetContentsCenter]
//   - [CALayer.Display]: Reloads the content of this layer.
//   - [CALayer.DrawInContext]: Draws the layer’s content using the specified graphics context.
//
// # Modifying the layer’s appearance
//
//   - [CALayer.ContentsGravity]: A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//   - [CALayer.SetContentsGravity]
//   - [CALayer.Opacity]: The opacity of the receiver. Animatable.
//   - [CALayer.SetOpacity]
//   - [CALayer.Hidden]: A Boolean indicating whether the layer is displayed. Animatable.
//   - [CALayer.SetHidden]
//   - [CALayer.MasksToBounds]: A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//   - [CALayer.SetMasksToBounds]
//   - [CALayer.Mask]: An optional layer whose alpha channel is used to mask the layer’s content.
//   - [CALayer.SetMask]
//   - [CALayer.DoubleSided]: A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//   - [CALayer.SetDoubleSided]
//   - [CALayer.CornerRadius]: The radius to use when drawing rounded corners for the layer’s background. Animatable.
//   - [CALayer.SetCornerRadius]
//   - [CALayer.MaskedCorners]
//   - [CALayer.SetMaskedCorners]
//   - [CALayer.BorderWidth]: The width of the layer’s border. Animatable.
//   - [CALayer.SetBorderWidth]
//   - [CALayer.BorderColor]: The color of the layer’s border. Animatable.
//   - [CALayer.SetBorderColor]
//   - [CALayer.BackgroundColor]: The background color of the receiver. Animatable.
//   - [CALayer.SetBackgroundColor]
//   - [CALayer.ShadowOpacity]: The opacity of the layer’s shadow. Animatable.
//   - [CALayer.SetShadowOpacity]
//   - [CALayer.ShadowRadius]: The blur radius (in points) used to render the layer’s shadow. Animatable.
//   - [CALayer.SetShadowRadius]
//   - [CALayer.ShadowOffset]: The offset (in points) of the layer’s shadow. Animatable.
//   - [CALayer.SetShadowOffset]
//   - [CALayer.ShadowColor]: The color of the layer’s shadow. Animatable.
//   - [CALayer.SetShadowColor]
//   - [CALayer.ShadowPath]: The shape of the layer’s shadow. Animatable.
//   - [CALayer.SetShadowPath]
//   - [CALayer.Style]: An optional dictionary used to store property values that aren’t explicitly defined by the layer.
//   - [CALayer.SetStyle]
//   - [CALayer.AllowsEdgeAntialiasing]: A Boolean indicating whether the layer is allowed to perform edge antialiasing.
//   - [CALayer.SetAllowsEdgeAntialiasing]
//   - [CALayer.AllowsGroupOpacity]: A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//   - [CALayer.SetAllowsGroupOpacity]
//
// # Layer filters
//
//   - [CALayer.Filters]: An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//   - [CALayer.SetFilters]
//   - [CALayer.CompositingFilter]: A CoreImage filter used to composite the layer and the content behind it. Animatable.
//   - [CALayer.SetCompositingFilter]
//   - [CALayer.BackgroundFilters]: An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//   - [CALayer.SetBackgroundFilters]
//   - [CALayer.MinificationFilter]: The filter used when reducing the size of the content.
//   - [CALayer.SetMinificationFilter]
//   - [CALayer.MinificationFilterBias]: The bias factor used by the minification filter to determine the levels of detail.
//   - [CALayer.SetMinificationFilterBias]
//   - [CALayer.MagnificationFilter]: The filter used when increasing the size of the content.
//   - [CALayer.SetMagnificationFilter]
//
// # Configuring the layer’s rendering behavior
//
//   - [CALayer.Opaque]: A Boolean value indicating whether the layer contains completely opaque content.
//   - [CALayer.SetOpaque]
//   - [CALayer.EdgeAntialiasingMask]: A bitmask defining how the edges of the receiver are rasterized.
//   - [CALayer.SetEdgeAntialiasingMask]
//   - [CALayer.ContentsAreFlipped]: Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
//   - [CALayer.GeometryFlipped]: A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//   - [CALayer.SetGeometryFlipped]
//   - [CALayer.DrawsAsynchronously]: A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//   - [CALayer.SetDrawsAsynchronously]
//   - [CALayer.ShouldRasterize]: A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//   - [CALayer.SetShouldRasterize]
//   - [CALayer.RasterizationScale]: The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//   - [CALayer.SetRasterizationScale]
//   - [CALayer.ContentsFormat]: A hint for the desired storage format of the layer contents.
//   - [CALayer.SetContentsFormat]
//   - [CALayer.RenderInContext]: Renders the layer and its sublayers into the specified context.
//
// # Modifying the layer geometry
//
//   - [CALayer.Frame]: The layer’s frame rectangle.
//   - [CALayer.SetFrame]
//   - [CALayer.Bounds]: The layer’s bounds rectangle. Animatable.
//   - [CALayer.SetBounds]
//   - [CALayer.Position]: The layer’s position in its superlayer’s coordinate space. Animatable.
//   - [CALayer.SetPosition]
//   - [CALayer.ZPosition]: The layer’s position on the z axis. Animatable.
//   - [CALayer.SetZPosition]
//   - [CALayer.AnchorPointZ]: The anchor point for the layer’s position along the z axis. Animatable.
//   - [CALayer.SetAnchorPointZ]
//   - [CALayer.AnchorPoint]: Defines the anchor point of the layer’s bounds rectangle. Animatable.
//   - [CALayer.SetAnchorPoint]
//   - [CALayer.ContentsScale]: The scale factor applied to the layer.
//   - [CALayer.SetContentsScale]
//
// # Managing the layer’s transform
//
//   - [CALayer.Transform]: The transform applied to the layer’s contents. Animatable.
//   - [CALayer.SetTransform]
//   - [CALayer.SublayerTransform]: Specifies the transform to apply to sublayers when rendering. Animatable.
//   - [CALayer.SetSublayerTransform]
//   - [CALayer.AffineTransform]: Returns an affine version of the layer’s transform.
//   - [CALayer.SetAffineTransform]: Sets the layer’s transform to the specified affine transform.
//
// # Managing the layer hierarchy
//
//   - [CALayer.Sublayers]: An array containing the layer’s sublayers.
//   - [CALayer.SetSublayers]
//   - [CALayer.Superlayer]: The superlayer of the layer.
//   - [CALayer.AddSublayer]: Appends the layer to the layer’s list of sublayers.
//   - [CALayer.RemoveFromSuperlayer]: Detaches the layer from its parent layer.
//   - [CALayer.InsertSublayerAtIndex]: Inserts the specified layer into the receiver’s list of sublayers at the specified index.
//   - [CALayer.InsertSublayerBelow]: Inserts the specified sublayer below a different sublayer that already belongs to the receiver.
//   - [CALayer.InsertSublayerAbove]: Inserts the specified sublayer above a different sublayer that already belongs to the receiver.
//   - [CALayer.ReplaceSublayerWith]: Replaces the specified sublayer with a different layer object.
//
// # Updating layer display
//
//   - [CALayer.SetNeedsDisplay]: Marks the layer’s contents as needing to be updated.
//   - [CALayer.SetNeedsDisplayInRect]: Marks the region within the specified rectangle as needing to be updated.
//   - [CALayer.NeedsDisplayOnBoundsChange]: A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//   - [CALayer.SetNeedsDisplayOnBoundsChange]
//   - [CALayer.DisplayIfNeeded]: Initiates the update process for a layer if it is currently marked as needing an update.
//   - [CALayer.NeedsDisplay]: Returns a Boolean indicating whether the layer has been marked as needing an update.
//
// # Layer animations
//
//   - [CALayer.AddAnimationForKey]: Add the specified animation object to the layer’s render tree.
//   - [CALayer.AnimationForKey]: Returns the animation object with the specified identifier.
//   - [CALayer.RemoveAllAnimations]: Remove all animations attached to the layer.
//   - [CALayer.RemoveAnimationForKey]: Remove the animation object with the specified key.
//   - [CALayer.AnimationKeys]: Returns an array of strings that identify the animations currently attached to the layer.
//
// # Managing layer resizing and layout
//
//   - [CALayer.LayoutManager]: The object responsible for laying out the layer’s sublayers.
//   - [CALayer.SetLayoutManager]
//   - [CALayer.SetNeedsLayout]: Invalidates the layer’s layout and marks it as needing an update.
//   - [CALayer.LayoutSublayers]: Tells the layer to update its layout.
//   - [CALayer.LayoutIfNeeded]: Recalculate the receiver’s layout, if required.
//   - [CALayer.NeedsLayout]: Returns a Boolean indicating whether the layer has been marked as needing a layout update.
//   - [CALayer.AutoresizingMask]: A bitmask defining how the layer is resized when the bounds of its superlayer changes.
//   - [CALayer.SetAutoresizingMask]
//   - [CALayer.ResizeWithOldSuperlayerSize]: Informs the receiver that the size of its superlayer changed.
//   - [CALayer.ResizeSublayersWithOldSize]: Informs the receiver’s sublayers that the receiver’s size has changed.
//   - [CALayer.PreferredFrameSize]: Returns the preferred size of the layer in the coordinate space of its superlayer.
//
// # Managing layer constraints
//
//   - [CALayer.Constraints]: The constraints used to position current layer’s sublayers.
//   - [CALayer.SetConstraints]
//   - [CALayer.AddConstraint]: Adds the specified constraint to the layer.
//
// # Getting the layer’s actions
//
//   - [CALayer.ActionForKey]: Returns the action object assigned to the specified key.
//   - [CALayer.Actions]: A dictionary containing layer actions.
//   - [CALayer.SetActions]
//
// # Mapping between coordinate and time spaces
//
//   - [CALayer.ConvertPointFromLayer]: Converts the point from the specified layer’s coordinate system to the receiver’s coordinate system.
//   - [CALayer.ConvertPointToLayer]: Converts the point from the receiver’s coordinate system to the specified layer’s coordinate system.
//   - [CALayer.ConvertRectFromLayer]: Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
//   - [CALayer.ConvertRectToLayer]: Converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
//   - [CALayer.ConvertTimeFromLayer]: Converts the time interval from the specified layer’s time space to the receiver’s time space.
//   - [CALayer.ConvertTimeToLayer]: Converts the time interval from the receiver’s time space to the specified layer’s time space
//
// # Hit testing
//
//   - [CALayer.HitTest]: Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point.
//   - [CALayer.ContainsPoint]: Returns whether the receiver contains a specified point.
//
// # Scrolling
//
//   - [CALayer.VisibleRect]: The visible region of the layer in its own coordinate space.
//   - [CALayer.ScrollPoint]: Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified point lies at the origin of the scroll layer.
//   - [CALayer.ScrollRectToVisible]: Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
//
// # Identifying the layer
//
//   - [CALayer.Name]: The name of the receiver.
//   - [CALayer.SetName]
//
// # Key-value coding extensions
//
//   - [CALayer.ShouldArchiveValueForKey]: Returns a Boolean indicating whether the value of the specified key should be archived.
//
// # High dynamic range
//
//   - [CALayer.PreferredDynamicRange]
//   - [CALayer.SetPreferredDynamicRange]
//   - [CALayer.ContentsHeadroom]
//   - [CALayer.SetContentsHeadroom]
//   - [CALayer.WantsExtendedDynamicRangeContent]
//   - [CALayer.SetWantsExtendedDynamicRangeContent]
//
// # Instance properties
//
//   - [CALayer.CornerCurve]
//   - [CALayer.SetCornerCurve]
//
// # Instance Properties
//
//   - [CALayer.ToneMapMode]
//   - [CALayer.SetToneMapMode]
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer
type CALayer struct {
	objectivec.Object
}

// CALayerFromID constructs a [CALayer] from an objc.ID.
//
// An object that manages image-based content and allows you to perform
// animations on that content.
func CALayerFromID(id objc.ID) CALayer {
	return CALayer{objectivec.Object{ID: id}}
}
// NOTE: CALayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CALayer] class.
//
// # Creating a layer
//
//   - [ICALayer.InitWithLayer]: Override to copy or initialize custom fields of the specified layer.
//
// # Accessing related layer objects
//
//   - [ICALayer.PresentationLayer]: Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
//   - [ICALayer.ModelLayer]: Returns the model layer object associated with the receiver, if any.
//
// # Accessing the delegate
//
//   - [ICALayer.Delegate]: The layer’s delegate object.
//   - [ICALayer.SetDelegate]
//
// # Providing the layer’s content
//
//   - [ICALayer.Contents]: An object that provides the contents of the layer. Animatable.
//   - [ICALayer.SetContents]
//   - [ICALayer.ContentsRect]: The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//   - [ICALayer.SetContentsRect]
//   - [ICALayer.ContentsCenter]: The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//   - [ICALayer.SetContentsCenter]
//   - [ICALayer.Display]: Reloads the content of this layer.
//   - [ICALayer.DrawInContext]: Draws the layer’s content using the specified graphics context.
//
// # Modifying the layer’s appearance
//
//   - [ICALayer.ContentsGravity]: A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//   - [ICALayer.SetContentsGravity]
//   - [ICALayer.Opacity]: The opacity of the receiver. Animatable.
//   - [ICALayer.SetOpacity]
//   - [ICALayer.Hidden]: A Boolean indicating whether the layer is displayed. Animatable.
//   - [ICALayer.SetHidden]
//   - [ICALayer.MasksToBounds]: A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//   - [ICALayer.SetMasksToBounds]
//   - [ICALayer.Mask]: An optional layer whose alpha channel is used to mask the layer’s content.
//   - [ICALayer.SetMask]
//   - [ICALayer.DoubleSided]: A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//   - [ICALayer.SetDoubleSided]
//   - [ICALayer.CornerRadius]: The radius to use when drawing rounded corners for the layer’s background. Animatable.
//   - [ICALayer.SetCornerRadius]
//   - [ICALayer.MaskedCorners]
//   - [ICALayer.SetMaskedCorners]
//   - [ICALayer.BorderWidth]: The width of the layer’s border. Animatable.
//   - [ICALayer.SetBorderWidth]
//   - [ICALayer.BorderColor]: The color of the layer’s border. Animatable.
//   - [ICALayer.SetBorderColor]
//   - [ICALayer.BackgroundColor]: The background color of the receiver. Animatable.
//   - [ICALayer.SetBackgroundColor]
//   - [ICALayer.ShadowOpacity]: The opacity of the layer’s shadow. Animatable.
//   - [ICALayer.SetShadowOpacity]
//   - [ICALayer.ShadowRadius]: The blur radius (in points) used to render the layer’s shadow. Animatable.
//   - [ICALayer.SetShadowRadius]
//   - [ICALayer.ShadowOffset]: The offset (in points) of the layer’s shadow. Animatable.
//   - [ICALayer.SetShadowOffset]
//   - [ICALayer.ShadowColor]: The color of the layer’s shadow. Animatable.
//   - [ICALayer.SetShadowColor]
//   - [ICALayer.ShadowPath]: The shape of the layer’s shadow. Animatable.
//   - [ICALayer.SetShadowPath]
//   - [ICALayer.Style]: An optional dictionary used to store property values that aren’t explicitly defined by the layer.
//   - [ICALayer.SetStyle]
//   - [ICALayer.AllowsEdgeAntialiasing]: A Boolean indicating whether the layer is allowed to perform edge antialiasing.
//   - [ICALayer.SetAllowsEdgeAntialiasing]
//   - [ICALayer.AllowsGroupOpacity]: A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//   - [ICALayer.SetAllowsGroupOpacity]
//
// # Layer filters
//
//   - [ICALayer.Filters]: An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//   - [ICALayer.SetFilters]
//   - [ICALayer.CompositingFilter]: A CoreImage filter used to composite the layer and the content behind it. Animatable.
//   - [ICALayer.SetCompositingFilter]
//   - [ICALayer.BackgroundFilters]: An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//   - [ICALayer.SetBackgroundFilters]
//   - [ICALayer.MinificationFilter]: The filter used when reducing the size of the content.
//   - [ICALayer.SetMinificationFilter]
//   - [ICALayer.MinificationFilterBias]: The bias factor used by the minification filter to determine the levels of detail.
//   - [ICALayer.SetMinificationFilterBias]
//   - [ICALayer.MagnificationFilter]: The filter used when increasing the size of the content.
//   - [ICALayer.SetMagnificationFilter]
//
// # Configuring the layer’s rendering behavior
//
//   - [ICALayer.Opaque]: A Boolean value indicating whether the layer contains completely opaque content.
//   - [ICALayer.SetOpaque]
//   - [ICALayer.EdgeAntialiasingMask]: A bitmask defining how the edges of the receiver are rasterized.
//   - [ICALayer.SetEdgeAntialiasingMask]
//   - [ICALayer.ContentsAreFlipped]: Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
//   - [ICALayer.GeometryFlipped]: A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//   - [ICALayer.SetGeometryFlipped]
//   - [ICALayer.DrawsAsynchronously]: A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//   - [ICALayer.SetDrawsAsynchronously]
//   - [ICALayer.ShouldRasterize]: A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//   - [ICALayer.SetShouldRasterize]
//   - [ICALayer.RasterizationScale]: The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//   - [ICALayer.SetRasterizationScale]
//   - [ICALayer.ContentsFormat]: A hint for the desired storage format of the layer contents.
//   - [ICALayer.SetContentsFormat]
//   - [ICALayer.RenderInContext]: Renders the layer and its sublayers into the specified context.
//
// # Modifying the layer geometry
//
//   - [ICALayer.Frame]: The layer’s frame rectangle.
//   - [ICALayer.SetFrame]
//   - [ICALayer.Bounds]: The layer’s bounds rectangle. Animatable.
//   - [ICALayer.SetBounds]
//   - [ICALayer.Position]: The layer’s position in its superlayer’s coordinate space. Animatable.
//   - [ICALayer.SetPosition]
//   - [ICALayer.ZPosition]: The layer’s position on the z axis. Animatable.
//   - [ICALayer.SetZPosition]
//   - [ICALayer.AnchorPointZ]: The anchor point for the layer’s position along the z axis. Animatable.
//   - [ICALayer.SetAnchorPointZ]
//   - [ICALayer.AnchorPoint]: Defines the anchor point of the layer’s bounds rectangle. Animatable.
//   - [ICALayer.SetAnchorPoint]
//   - [ICALayer.ContentsScale]: The scale factor applied to the layer.
//   - [ICALayer.SetContentsScale]
//
// # Managing the layer’s transform
//
//   - [ICALayer.Transform]: The transform applied to the layer’s contents. Animatable.
//   - [ICALayer.SetTransform]
//   - [ICALayer.SublayerTransform]: Specifies the transform to apply to sublayers when rendering. Animatable.
//   - [ICALayer.SetSublayerTransform]
//   - [ICALayer.AffineTransform]: Returns an affine version of the layer’s transform.
//   - [ICALayer.SetAffineTransform]: Sets the layer’s transform to the specified affine transform.
//
// # Managing the layer hierarchy
//
//   - [ICALayer.Sublayers]: An array containing the layer’s sublayers.
//   - [ICALayer.SetSublayers]
//   - [ICALayer.Superlayer]: The superlayer of the layer.
//   - [ICALayer.AddSublayer]: Appends the layer to the layer’s list of sublayers.
//   - [ICALayer.RemoveFromSuperlayer]: Detaches the layer from its parent layer.
//   - [ICALayer.InsertSublayerAtIndex]: Inserts the specified layer into the receiver’s list of sublayers at the specified index.
//   - [ICALayer.InsertSublayerBelow]: Inserts the specified sublayer below a different sublayer that already belongs to the receiver.
//   - [ICALayer.InsertSublayerAbove]: Inserts the specified sublayer above a different sublayer that already belongs to the receiver.
//   - [ICALayer.ReplaceSublayerWith]: Replaces the specified sublayer with a different layer object.
//
// # Updating layer display
//
//   - [ICALayer.SetNeedsDisplay]: Marks the layer’s contents as needing to be updated.
//   - [ICALayer.SetNeedsDisplayInRect]: Marks the region within the specified rectangle as needing to be updated.
//   - [ICALayer.NeedsDisplayOnBoundsChange]: A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//   - [ICALayer.SetNeedsDisplayOnBoundsChange]
//   - [ICALayer.DisplayIfNeeded]: Initiates the update process for a layer if it is currently marked as needing an update.
//   - [ICALayer.NeedsDisplay]: Returns a Boolean indicating whether the layer has been marked as needing an update.
//
// # Layer animations
//
//   - [ICALayer.AddAnimationForKey]: Add the specified animation object to the layer’s render tree.
//   - [ICALayer.AnimationForKey]: Returns the animation object with the specified identifier.
//   - [ICALayer.RemoveAllAnimations]: Remove all animations attached to the layer.
//   - [ICALayer.RemoveAnimationForKey]: Remove the animation object with the specified key.
//   - [ICALayer.AnimationKeys]: Returns an array of strings that identify the animations currently attached to the layer.
//
// # Managing layer resizing and layout
//
//   - [ICALayer.LayoutManager]: The object responsible for laying out the layer’s sublayers.
//   - [ICALayer.SetLayoutManager]
//   - [ICALayer.SetNeedsLayout]: Invalidates the layer’s layout and marks it as needing an update.
//   - [ICALayer.LayoutSublayers]: Tells the layer to update its layout.
//   - [ICALayer.LayoutIfNeeded]: Recalculate the receiver’s layout, if required.
//   - [ICALayer.NeedsLayout]: Returns a Boolean indicating whether the layer has been marked as needing a layout update.
//   - [ICALayer.AutoresizingMask]: A bitmask defining how the layer is resized when the bounds of its superlayer changes.
//   - [ICALayer.SetAutoresizingMask]
//   - [ICALayer.ResizeWithOldSuperlayerSize]: Informs the receiver that the size of its superlayer changed.
//   - [ICALayer.ResizeSublayersWithOldSize]: Informs the receiver’s sublayers that the receiver’s size has changed.
//   - [ICALayer.PreferredFrameSize]: Returns the preferred size of the layer in the coordinate space of its superlayer.
//
// # Managing layer constraints
//
//   - [ICALayer.Constraints]: The constraints used to position current layer’s sublayers.
//   - [ICALayer.SetConstraints]
//   - [ICALayer.AddConstraint]: Adds the specified constraint to the layer.
//
// # Getting the layer’s actions
//
//   - [ICALayer.ActionForKey]: Returns the action object assigned to the specified key.
//   - [ICALayer.Actions]: A dictionary containing layer actions.
//   - [ICALayer.SetActions]
//
// # Mapping between coordinate and time spaces
//
//   - [ICALayer.ConvertPointFromLayer]: Converts the point from the specified layer’s coordinate system to the receiver’s coordinate system.
//   - [ICALayer.ConvertPointToLayer]: Converts the point from the receiver’s coordinate system to the specified layer’s coordinate system.
//   - [ICALayer.ConvertRectFromLayer]: Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
//   - [ICALayer.ConvertRectToLayer]: Converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
//   - [ICALayer.ConvertTimeFromLayer]: Converts the time interval from the specified layer’s time space to the receiver’s time space.
//   - [ICALayer.ConvertTimeToLayer]: Converts the time interval from the receiver’s time space to the specified layer’s time space
//
// # Hit testing
//
//   - [ICALayer.HitTest]: Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point.
//   - [ICALayer.ContainsPoint]: Returns whether the receiver contains a specified point.
//
// # Scrolling
//
//   - [ICALayer.VisibleRect]: The visible region of the layer in its own coordinate space.
//   - [ICALayer.ScrollPoint]: Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified point lies at the origin of the scroll layer.
//   - [ICALayer.ScrollRectToVisible]: Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
//
// # Identifying the layer
//
//   - [ICALayer.Name]: The name of the receiver.
//   - [ICALayer.SetName]
//
// # Key-value coding extensions
//
//   - [ICALayer.ShouldArchiveValueForKey]: Returns a Boolean indicating whether the value of the specified key should be archived.
//
// # High dynamic range
//
//   - [ICALayer.PreferredDynamicRange]
//   - [ICALayer.SetPreferredDynamicRange]
//   - [ICALayer.ContentsHeadroom]
//   - [ICALayer.SetContentsHeadroom]
//   - [ICALayer.WantsExtendedDynamicRangeContent]
//   - [ICALayer.SetWantsExtendedDynamicRangeContent]
//
// # Instance properties
//
//   - [ICALayer.CornerCurve]
//   - [ICALayer.SetCornerCurve]
//
// # Instance Properties
//
//   - [ICALayer.ToneMapMode]
//   - [ICALayer.SetToneMapMode]
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer
type ICALayer interface {
	objectivec.IObject
	CAMediaTiming

	// Topic: Creating a layer

	// Override to copy or initialize custom fields of the specified layer.
	InitWithLayer(layer objectivec.IObject) CALayer

	// Topic: Accessing related layer objects

	// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
	PresentationLayer() ICALayer
	// Returns the model layer object associated with the receiver, if any.
	ModelLayer() ICALayer

	// Topic: Accessing the delegate

	// The layer’s delegate object.
	Delegate() CALayerDelegate
	SetDelegate(value CALayerDelegate)

	// Topic: Providing the layer’s content

	// An object that provides the contents of the layer. Animatable.
	Contents() objectivec.IObject
	SetContents(value objectivec.IObject)
	// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
	ContentsRect() corefoundation.CGRect
	SetContentsRect(value corefoundation.CGRect)
	// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
	ContentsCenter() corefoundation.CGRect
	SetContentsCenter(value corefoundation.CGRect)
	// Reloads the content of this layer.
	Display()
	// Draws the layer’s content using the specified graphics context.
	DrawInContext(ctx coregraphics.CGContextRef)

	// Topic: Modifying the layer’s appearance

	// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
	ContentsGravity() CALayerContentsGravity
	SetContentsGravity(value CALayerContentsGravity)
	// The opacity of the receiver. Animatable.
	Opacity() float32
	SetOpacity(value float32)
	// A Boolean indicating whether the layer is displayed. Animatable.
	Hidden() bool
	SetHidden(value bool)
	// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
	MasksToBounds() bool
	SetMasksToBounds(value bool)
	// An optional layer whose alpha channel is used to mask the layer’s content.
	Mask() ICALayer
	SetMask(value ICALayer)
	// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
	DoubleSided() bool
	SetDoubleSided(value bool)
	// The radius to use when drawing rounded corners for the layer’s background. Animatable.
	CornerRadius() float64
	SetCornerRadius(value float64)
	MaskedCorners() CACornerMask
	SetMaskedCorners(value CACornerMask)
	// The width of the layer’s border. Animatable.
	BorderWidth() float64
	SetBorderWidth(value float64)
	// The color of the layer’s border. Animatable.
	BorderColor() coregraphics.CGColorRef
	SetBorderColor(value coregraphics.CGColorRef)
	// The background color of the receiver. Animatable.
	BackgroundColor() coregraphics.CGColorRef
	SetBackgroundColor(value coregraphics.CGColorRef)
	// The opacity of the layer’s shadow. Animatable.
	ShadowOpacity() float32
	SetShadowOpacity(value float32)
	// The blur radius (in points) used to render the layer’s shadow. Animatable.
	ShadowRadius() float64
	SetShadowRadius(value float64)
	// The offset (in points) of the layer’s shadow. Animatable.
	ShadowOffset() corefoundation.CGSize
	SetShadowOffset(value corefoundation.CGSize)
	// The color of the layer’s shadow. Animatable.
	ShadowColor() coregraphics.CGColorRef
	SetShadowColor(value coregraphics.CGColorRef)
	// The shape of the layer’s shadow. Animatable.
	ShadowPath() coregraphics.CGPathRef
	SetShadowPath(value coregraphics.CGPathRef)
	// An optional dictionary used to store property values that aren’t explicitly defined by the layer.
	Style() foundation.INSDictionary
	SetStyle(value foundation.INSDictionary)
	// A Boolean indicating whether the layer is allowed to perform edge antialiasing.
	AllowsEdgeAntialiasing() bool
	SetAllowsEdgeAntialiasing(value bool)
	// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
	AllowsGroupOpacity() bool
	SetAllowsGroupOpacity(value bool)

	// Topic: Layer filters

	// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
	Filters() foundation.INSArray
	SetFilters(value foundation.INSArray)
	// A CoreImage filter used to composite the layer and the content behind it. Animatable.
	CompositingFilter() objectivec.IObject
	SetCompositingFilter(value objectivec.IObject)
	// An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
	BackgroundFilters() foundation.INSArray
	SetBackgroundFilters(value foundation.INSArray)
	// The filter used when reducing the size of the content.
	MinificationFilter() CALayerContentsFilter
	SetMinificationFilter(value CALayerContentsFilter)
	// The bias factor used by the minification filter to determine the levels of detail.
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	// The filter used when increasing the size of the content.
	MagnificationFilter() CALayerContentsFilter
	SetMagnificationFilter(value CALayerContentsFilter)

	// Topic: Configuring the layer’s rendering behavior

	// A Boolean value indicating whether the layer contains completely opaque content.
	Opaque() bool
	SetOpaque(value bool)
	// A bitmask defining how the edges of the receiver are rasterized.
	EdgeAntialiasingMask() CAEdgeAntialiasingMask
	SetEdgeAntialiasingMask(value CAEdgeAntialiasingMask)
	// Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
	ContentsAreFlipped() bool
	// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
	GeometryFlipped() bool
	SetGeometryFlipped(value bool)
	// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
	DrawsAsynchronously() bool
	SetDrawsAsynchronously(value bool)
	// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
	ShouldRasterize() bool
	SetShouldRasterize(value bool)
	// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
	RasterizationScale() float64
	SetRasterizationScale(value float64)
	// A hint for the desired storage format of the layer contents.
	ContentsFormat() CALayerContentsFormat
	SetContentsFormat(value CALayerContentsFormat)
	// Renders the layer and its sublayers into the specified context.
	RenderInContext(ctx coregraphics.CGContextRef)

	// Topic: Modifying the layer geometry

	// The layer’s frame rectangle.
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
	// The layer’s bounds rectangle. Animatable.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	// The layer’s position in its superlayer’s coordinate space. Animatable.
	Position() corefoundation.CGPoint
	SetPosition(value corefoundation.CGPoint)
	// The layer’s position on the z axis. Animatable.
	ZPosition() float64
	SetZPosition(value float64)
	// The anchor point for the layer’s position along the z axis. Animatable.
	AnchorPointZ() float64
	SetAnchorPointZ(value float64)
	// Defines the anchor point of the layer’s bounds rectangle. Animatable.
	AnchorPoint() corefoundation.CGPoint
	SetAnchorPoint(value corefoundation.CGPoint)
	// The scale factor applied to the layer.
	ContentsScale() float64
	SetContentsScale(value float64)

	// Topic: Managing the layer’s transform

	// The transform applied to the layer’s contents. Animatable.
	Transform() CATransform3D
	SetTransform(value CATransform3D)
	// Specifies the transform to apply to sublayers when rendering. Animatable.
	SublayerTransform() CATransform3D
	SetSublayerTransform(value CATransform3D)
	// Returns an affine version of the layer’s transform.
	AffineTransform() corefoundation.CGAffineTransform
	// Sets the layer’s transform to the specified affine transform.
	SetAffineTransform(m corefoundation.CGAffineTransform)

	// Topic: Managing the layer hierarchy

	// An array containing the layer’s sublayers.
	Sublayers() []CALayer
	SetSublayers(value []CALayer)
	// The superlayer of the layer.
	Superlayer() ICALayer
	// Appends the layer to the layer’s list of sublayers.
	AddSublayer(layer ICALayer)
	// Detaches the layer from its parent layer.
	RemoveFromSuperlayer()
	// Inserts the specified layer into the receiver’s list of sublayers at the specified index.
	InsertSublayerAtIndex(layer ICALayer, idx uint32)
	// Inserts the specified sublayer below a different sublayer that already belongs to the receiver.
	InsertSublayerBelow(layer ICALayer, sibling ICALayer)
	// Inserts the specified sublayer above a different sublayer that already belongs to the receiver.
	InsertSublayerAbove(layer ICALayer, sibling ICALayer)
	// Replaces the specified sublayer with a different layer object.
	ReplaceSublayerWith(oldLayer ICALayer, newLayer ICALayer)

	// Topic: Updating layer display

	// Marks the layer’s contents as needing to be updated.
	SetNeedsDisplay()
	// Marks the region within the specified rectangle as needing to be updated.
	SetNeedsDisplayInRect(r corefoundation.CGRect)
	// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
	NeedsDisplayOnBoundsChange() bool
	SetNeedsDisplayOnBoundsChange(value bool)
	// Initiates the update process for a layer if it is currently marked as needing an update.
	DisplayIfNeeded()
	// Returns a Boolean indicating whether the layer has been marked as needing an update.
	NeedsDisplay() bool

	// Topic: Layer animations

	// Add the specified animation object to the layer’s render tree.
	AddAnimationForKey(anim ICAAnimation, key string)
	// Returns the animation object with the specified identifier.
	AnimationForKey(key string) ICAAnimation
	// Remove all animations attached to the layer.
	RemoveAllAnimations()
	// Remove the animation object with the specified key.
	RemoveAnimationForKey(key string)
	// Returns an array of strings that identify the animations currently attached to the layer.
	AnimationKeys() []string

	// Topic: Managing layer resizing and layout

	// The object responsible for laying out the layer’s sublayers.
	LayoutManager() CALayoutManager
	SetLayoutManager(value CALayoutManager)
	// Invalidates the layer’s layout and marks it as needing an update.
	SetNeedsLayout()
	// Tells the layer to update its layout.
	LayoutSublayers()
	// Recalculate the receiver’s layout, if required.
	LayoutIfNeeded()
	// Returns a Boolean indicating whether the layer has been marked as needing a layout update.
	NeedsLayout() bool
	// A bitmask defining how the layer is resized when the bounds of its superlayer changes.
	AutoresizingMask() CAAutoresizingMask
	SetAutoresizingMask(value CAAutoresizingMask)
	// Informs the receiver that the size of its superlayer changed.
	ResizeWithOldSuperlayerSize(size corefoundation.CGSize)
	// Informs the receiver’s sublayers that the receiver’s size has changed.
	ResizeSublayersWithOldSize(size corefoundation.CGSize)
	// Returns the preferred size of the layer in the coordinate space of its superlayer.
	PreferredFrameSize() corefoundation.CGSize

	// Topic: Managing layer constraints

	// The constraints used to position current layer’s sublayers.
	Constraints() []CAConstraint
	SetConstraints(value []CAConstraint)
	// Adds the specified constraint to the layer.
	AddConstraint(c ICAConstraint)

	// Topic: Getting the layer’s actions

	// Returns the action object assigned to the specified key.
	ActionForKey(event string) CAAction
	// A dictionary containing layer actions.
	Actions() foundation.INSDictionary
	SetActions(value foundation.INSDictionary)

	// Topic: Mapping between coordinate and time spaces

	// Converts the point from the specified layer’s coordinate system to the receiver’s coordinate system.
	ConvertPointFromLayer(p corefoundation.CGPoint, l ICALayer) corefoundation.CGPoint
	// Converts the point from the receiver’s coordinate system to the specified layer’s coordinate system.
	ConvertPointToLayer(p corefoundation.CGPoint, l ICALayer) corefoundation.CGPoint
	// Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
	ConvertRectFromLayer(r corefoundation.CGRect, l ICALayer) corefoundation.CGRect
	// Converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
	ConvertRectToLayer(r corefoundation.CGRect, l ICALayer) corefoundation.CGRect
	// Converts the time interval from the specified layer’s time space to the receiver’s time space.
	ConvertTimeFromLayer(t float64, l ICALayer) float64
	// Converts the time interval from the receiver’s time space to the specified layer’s time space
	ConvertTimeToLayer(t float64, l ICALayer) float64

	// Topic: Hit testing

	// Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point.
	HitTest(p corefoundation.CGPoint) ICALayer
	// Returns whether the receiver contains a specified point.
	ContainsPoint(p corefoundation.CGPoint) bool

	// Topic: Scrolling

	// The visible region of the layer in its own coordinate space.
	VisibleRect() corefoundation.CGRect
	// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified point lies at the origin of the scroll layer.
	ScrollPoint(p corefoundation.CGPoint)
	// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
	ScrollRectToVisible(r corefoundation.CGRect)

	// Topic: Identifying the layer

	// The name of the receiver.
	Name() string
	SetName(value string)

	// Topic: Key-value coding extensions

	// Returns a Boolean indicating whether the value of the specified key should be archived.
	ShouldArchiveValueForKey(key string) bool

	// Topic: High dynamic range

	PreferredDynamicRange() CADynamicRange
	SetPreferredDynamicRange(value CADynamicRange)
	ContentsHeadroom() float64
	SetContentsHeadroom(value float64)
	WantsExtendedDynamicRangeContent() bool
	SetWantsExtendedDynamicRangeContent(value bool)

	// Topic: Instance properties

	CornerCurve() CALayerCornerCurve
	SetCornerCurve(value CALayerCornerCurve)

	// Topic: Instance Properties

	ToneMapMode() uint
	SetToneMapMode(value uint)

	EncodeWithCoder(coder foundation.INSCoder)
	// Sets the value of the property identified by the given key.
	SetValueForKey(value objectivec.IObject, key string)
	// Returns the value of the property identified by the given key.
	ValueForKey(key string) objectivec.IObject
}

// Init initializes the instance.
func (l CALayer) Init() CALayer {
	rv := objc.Send[CALayer](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l CALayer) Autorelease() CALayer {
	rv := objc.Send[CALayer](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewCALayer creates a new CALayer instance.
func NewCALayer() CALayer {
	class := getCALayerClass()
	rv := objc.Send[CALayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
// 
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
// 
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
// 
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
// 
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewLayerWithLayer(layer objectivec.IObject) CALayer {
	instance := getCALayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CALayerFromID(rv)
}

// Initializes a layer with a remote client ID.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(remoteClientId:)
func NewLayerWithRemoteClientId(client_id uint32) CALayer {
	rv := objc.Send[objc.ID](objc.ID(getCALayerClass().class), objc.Sel("layerWithRemoteClientId:"), client_id)
	return CALayerFromID(rv)
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
// 
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
// 
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
// 
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
// 
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func (l CALayer) InitWithLayer(layer objectivec.IObject) CALayer {
	rv := objc.Send[CALayer](l.ID, objc.Sel("initWithLayer:"), layer)
	return rv
}
// Returns a copy of the presentation layer object that represents the state
// of the layer as it currently appears onscreen.
//
// # Return Value
// 
// A copy of the current presentation layer object.
//
// # Discussion
// 
// The layer object returned by this method provides a close approximation of
// the layer that is currently being displayed onscreen. While an animation is
// in progress, you can retrieve this object and use it to get the current
// values for those animations.
// 
// The [Sublayers], [Mask], and [Superlayer] properties of the returned layer
// return the corresponding objects from the presentation tree (not the model
// tree). This pattern also applies to any read-only layer methods. For
// example, the [HitTest] method of the returned object queries the layer
// objects in the presentation tree.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/presentation()
func (l CALayer) PresentationLayer() ICALayer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("presentationLayer"))
	return CALayerFromID(rv)
}
// Returns the model layer object associated with the receiver, if any.
//
// # Return Value
// 
// A layer instance representing the underlying model layer.
//
// # Discussion
// 
// Calling this method on a layer in the presentation tree returns the
// corresponding layer object in the model tree. This method returns a value
// only when a transaction involving changes to the presentation layer is in
// progress. If no transaction is in progress, the results of calling this
// method are undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/model()
func (l CALayer) ModelLayer() ICALayer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelLayer"))
	return CALayerFromID(rv)
}
// Reloads the content of this layer.
//
// # Discussion
// 
// Do not call this method directly. The layer calls this method at
// appropriate times to update the layer’s content. If the layer has a
// delegate object, this method attempts to call the delegate’s
// [DisplayLayer] method, which the delegate can use to update the layer’s
// contents. If the delegate does not implement the [DisplayLayer] method,
// this method creates a backing store and calls the layer’s [DrawInContext]
// method to fill that backing store with content. The new backing store
// replaces the previous contents of the layer.
// 
// Subclasses can override this method and use it to set the layer’s
// [Contents] property directly. You might do this if your custom layer
// subclass handles layer updates differently.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/display()
func (l CALayer) Display() {
	objc.Send[objc.ID](l.ID, objc.Sel("display"))
}
// Draws the layer’s content using the specified graphics context.
//
// ctx: The graphics context in which to draw the content. The context may be
// clipped to protect valid layer content. Subclasses that wish to find the
// actual region to draw can call [boundingBoxOfClipPath].
// //
// [boundingBoxOfClipPath]: https://developer.apple.com/documentation/CoreGraphics/CGContext/boundingBoxOfClipPath
//
// # Discussion
// 
// The default implementation of this method does not do any drawing itself.
// If the layer’s delegate implements the [DrawLayerInContext] method, that
// method is called to do the actual drawing.
// 
// Subclasses can override this method and use it to draw the layer’s
// content. When drawing, all coordinates should be specified in points in the
// logical coordinate space.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/draw(in:)
func (l CALayer) DrawInContext(ctx coregraphics.CGContextRef) {
	objc.Send[objc.ID](l.ID, objc.Sel("drawInContext:"), ctx)
}
// Returns a Boolean indicating whether the layer content is implicitly
// flipped when rendered.
//
// # Return Value
// 
// [true] if the layer contents are implicitly flipped when rendered or
// [false] if they are not. This method returns [false] by default.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method provides information about whether the layer’s contents are
// being flipped during drawing. You should not attempt to override this
// method and return a different value.
// 
// If the layer needs to flip its content, it returns [true] from this method
// and applies a y-flip transform to the graphics context before passing it to
// the layer’s [DrawInContext] method. Similarly, the layer converts any
// rectangles passed to its [SetNeedsDisplayInRect] into the flipped
// coordinate space.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsAreFlipped()
func (l CALayer) ContentsAreFlipped() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("contentsAreFlipped"))
	return rv
}
// Renders the layer and its sublayers into the specified context.
//
// ctx: The graphics context to use to render the layer.
//
// # Discussion
// 
// This method renders directly from the layer tree, ignoring any animations
// added to the render tree. Renders in the coordinate space of the layer.
// 
// The following code shows how you can use [RenderInContext] to create a
// [UIImage] from a [CAShapeLayer] with a [Path] that describes a circle.
// After creating the layer, the code creates a [CGContext] into which the
// circle is rendered. After rendering,
// [UIGraphicsGetImageFromCurrentImageContext()] generates the image.
//
// [CGContext]: https://developer.apple.com/documentation/CoreGraphics/CGContext
// [UIGraphicsGetImageFromCurrentImageContext()]: https://developer.apple.com/documentation/UIKit/UIGraphicsGetImageFromCurrentImageContext()
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/render(in:)
func (l CALayer) RenderInContext(ctx coregraphics.CGContextRef) {
	objc.Send[objc.ID](l.ID, objc.Sel("renderInContext:"), ctx)
}
// Returns an affine version of the layer’s transform.
//
// # Return Value
// 
// The affine transform structure that corresponds to the value in the
// layer’s [Transform] property.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/affineTransform()
func (l CALayer) AffineTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](l.ID, objc.Sel("affineTransform"))
	return corefoundation.CGAffineTransform(rv)
}
// Sets the layer’s transform to the specified affine transform.
//
// m: The affine transform to use for the layer’s transform.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/setAffineTransform(_:)
func (l CALayer) SetAffineTransform(m corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](l.ID, objc.Sel("setAffineTransform:"), m)
}
// Appends the layer to the layer’s list of sublayers.
//
// layer: The layer to be added.
//
// # Discussion
// 
// If the array in the sublayers property is `nil`, calling this method
// creates an array for that property and adds the specified layer to it.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/addSublayer(_:)
func (l CALayer) AddSublayer(layer ICALayer) {
	objc.Send[objc.ID](l.ID, objc.Sel("addSublayer:"), layer)
}
// Detaches the layer from its parent layer.
//
// # Discussion
// 
// You can use this method to remove a layer (and all of its sublayers) from a
// layer hierarchy. This method updates both the superlayer’s list of
// sublayers and sets this layer’s [Superlayer] property to `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/removeFromSuperlayer()
func (l CALayer) RemoveFromSuperlayer() {
	objc.Send[objc.ID](l.ID, objc.Sel("removeFromSuperlayer"))
}
// Inserts the specified layer into the receiver’s list of sublayers at the
// specified index.
//
// layer: The sublayer to be inserted into the current layer.
//
// idx: The index at which to insert `aLayer`. This value must be a valid 0-based
// index into the [Sublayers] array.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:at:)
func (l CALayer) InsertSublayerAtIndex(layer ICALayer, idx uint32) {
	objc.Send[objc.ID](l.ID, objc.Sel("insertSublayer:atIndex:"), layer, idx)
}
// Inserts the specified sublayer below a different sublayer that already
// belongs to the receiver.
//
// layer: The sublayer to be inserted into the current layer.
//
// sibling: An existing sublayer in the current layer. The layer in `aLayer` is
// inserted before this layer in the [Sublayers] array, and thus appears
// behind it visually.
//
// # Discussion
// 
// If `sublayer` is not in the receiver’s [Sublayers] array, this method
// raises an exception.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:below:)
func (l CALayer) InsertSublayerBelow(layer ICALayer, sibling ICALayer) {
	objc.Send[objc.ID](l.ID, objc.Sel("insertSublayer:below:"), layer, sibling)
}
// Inserts the specified sublayer above a different sublayer that already
// belongs to the receiver.
//
// layer: The sublayer to be inserted into the current layer.
//
// sibling: An existing sublayer in the current layer. The layer in `aLayer` is
// inserted after this layer in the [Sublayers] array, and thus appears in
// front of it visually.
//
// # Discussion
// 
// If `sublayer` is not in the receiver’s [Sublayers] array, this method
// raises an exception.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:above:)
func (l CALayer) InsertSublayerAbove(layer ICALayer, sibling ICALayer) {
	objc.Send[objc.ID](l.ID, objc.Sel("insertSublayer:above:"), layer, sibling)
}
// Replaces the specified sublayer with a different layer object.
//
// oldLayer: The layer to be replaced.
//
// newLayer: The layer with which to replace `oldLayer`.
//
// # Discussion
// 
// If `oldLayer` is not in the receiver’s [Sublayers] array, the behavior of
// this method is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/replaceSublayer(_:with:)
func (l CALayer) ReplaceSublayerWith(oldLayer ICALayer, newLayer ICALayer) {
	objc.Send[objc.ID](l.ID, objc.Sel("replaceSublayer:with:"), oldLayer, newLayer)
}
// Marks the layer’s contents as needing to be updated.
//
// # Discussion
// 
// Calling this method causes the layer to recache its content. This results
// in the layer potentially calling either the [DisplayLayer] or
// [DrawLayerInContext] method of its delegate. The existing content in the
// layer’s [Contents] property is removed to make way for the new content.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay()
func (l CALayer) SetNeedsDisplay() {
	objc.Send[objc.ID](l.ID, objc.Sel("setNeedsDisplay"))
}
// Marks the region within the specified rectangle as needing to be updated.
//
// r: The rectangular region of the layer to mark as invalid. You must specify
// this rectangle in the layer’s own coordinate system.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay(_:)
func (l CALayer) SetNeedsDisplayInRect(r corefoundation.CGRect) {
	objc.Send[objc.ID](l.ID, objc.Sel("setNeedsDisplayInRect:"), r)
}
// Initiates the update process for a layer if it is currently marked as
// needing an update.
//
// # Discussion
// 
// You can call this method as needed to force an update to your layer’s
// contents outside of the normal update cycle. Doing so is generally not
// needed, though. The preferred way to update a layer is to call
// [SetNeedsDisplay] and let the system update the layer during the next
// cycle.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/displayIfNeeded()
func (l CALayer) DisplayIfNeeded() {
	objc.Send[objc.ID](l.ID, objc.Sel("displayIfNeeded"))
}
// Returns a Boolean indicating whether the layer has been marked as needing
// an update.
//
// # Return Value
// 
// [true] if the layer needs to be updated.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplay()
func (l CALayer) NeedsDisplay() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("needsDisplay"))
	return rv
}
// Add the specified animation object to the layer’s render tree.
//
// anim: The animation to be added to the render tree. This object is copied by the
// render tree, not referenced. Therefore, subsequent modifications to the
// object are not propagated into the render tree.
//
// key: A string that identifies the animation. Only one animation per unique key
// is added to the layer. The special key [kCATransition] is automatically
// used for transition animations. You may specify `nil` for this parameter.
// //
// [kCATransition]: https://developer.apple.com/documentation/QuartzCore/kCATransition
//
// # Discussion
// 
// If the `duration` property of the animation is zero or negative, the
// duration is changed to the current value of the
// [kCATransactionAnimationDuration] transaction property (if set) or to the
// default value of `0.25` seconds.
//
// [kCATransactionAnimationDuration]: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationDuration
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/add(_:forKey:)
func (l CALayer) AddAnimationForKey(anim ICAAnimation, key string) {
	objc.Send[objc.ID](l.ID, objc.Sel("addAnimation:forKey:"), anim, objc.String(key))
}
// Returns the animation object with the specified identifier.
//
// key: A string that specifies the identifier of the animation. This string
// corresponds to the identifier string you passed to the [AddAnimationForKey]
// method.
//
// # Return Value
// 
// The animation object matching the identifier, or `nil` if no such animation
// exists.
//
// # Discussion
// 
// Use this method to retrieve only animation objects already associated with
// a layer. Modifying any properties of the returned object results in
// undefined behavior.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/animation(forKey:)
func (l CALayer) AnimationForKey(key string) ICAAnimation {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("animationForKey:"), objc.String(key))
	return CAAnimationFromID(rv)
}
// Remove all animations attached to the layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/removeAllAnimations()
func (l CALayer) RemoveAllAnimations() {
	objc.Send[objc.ID](l.ID, objc.Sel("removeAllAnimations"))
}
// Remove the animation object with the specified key.
//
// key: The identifier of the animation to remove.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/removeAnimation(forKey:)
func (l CALayer) RemoveAnimationForKey(key string) {
	objc.Send[objc.ID](l.ID, objc.Sel("removeAnimationForKey:"), objc.String(key))
}
// Returns an array of strings that identify the animations currently attached
// to the layer.
//
// # Return Value
// 
// An array of [NSString] objects identifying the current animations.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// # Discussion
// 
// The order of the array matches the order in which animations will be
// applied to the layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/animationKeys()
func (l CALayer) AnimationKeys() []string {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("animationKeys"))
	return objc.ConvertSliceToStrings(rv)
}
// Invalidates the layer’s layout and marks it as needing an update.
//
// # Discussion
// 
// You can call this method to indicate that the layout of a layer’s
// sublayers has changed and must be updated. The system typically calls this
// method automatically when the layer’s bounds change or when sublayers are
// added or removed. In macOS, if your layer’s [LayoutManager] property
// contains an object that implements the `invalidateLayout()` method in Swift
// or the `` method in Objective-C, the system calls that method too.
// 
// During the next update cycle, the system calls the [LayoutSublayers] method
// of any layers requiring layout updates.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsLayout()
func (l CALayer) SetNeedsLayout() {
	objc.Send[objc.ID](l.ID, objc.Sel("setNeedsLayout"))
}
// Tells the layer to update its layout.
//
// # Discussion
// 
// Subclasses can override this method and use it to implement their own
// layout algorithm. Your implementation must set the frame of each sublayer
// managed by the receiver.
// 
// The default implementation of this method calls the `layoutSublayers()`
// method in Swift or `` method in Objective-C of the layer’s delegate
// object. If there is no delegate object, or the delegate does not implement
// that method, this method calls the `layoutSublayers()` method in Swift or
// `` method in Objective-C of the object in the [LayoutManager] property.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutSublayers()
func (l CALayer) LayoutSublayers() {
	objc.Send[objc.ID](l.ID, objc.Sel("layoutSublayers"))
}
// Recalculate the receiver’s layout, if required.
//
// # Discussion
// 
// When this message is received, the layer’s super layers are traversed
// until a ancestor layer is found that does not require layout. Then layout
// is performed on the entire layer-tree beneath that ancestor.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutIfNeeded()
func (l CALayer) LayoutIfNeeded() {
	objc.Send[objc.ID](l.ID, objc.Sel("layoutIfNeeded"))
}
// Returns a Boolean indicating whether the layer has been marked as needing a
// layout update.
//
// # Return Value
// 
// [true] if the layer has been marked as requiring a layout update.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/needsLayout()
func (l CALayer) NeedsLayout() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("needsLayout"))
	return rv
}
// Informs the receiver that the size of its superlayer changed.
//
// size: The previous size of the superlayer.
//
// # Discussion
// 
// When the [AutoresizingMask] property is used for resizing and the bounds of
// a layer change, that layer calls this method on each of its sublayers.
// Sublayers use this method to adjust their own frame rectangles to reflect
// the new superlayer bounds, which can be retrieved directly from the
// superlayer. The old size of the superlayer is passed to this method so that
// the sublayer has that information for any calculations it must make.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/resize(withOldSuperlayerSize:)
func (l CALayer) ResizeWithOldSuperlayerSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](l.ID, objc.Sel("resizeWithOldSuperlayerSize:"), size)
}
// Informs the receiver’s sublayers that the receiver’s size has changed.
//
// size: The previous size of the current layer.
//
// # Discussion
// 
// When the [AutoresizingMask] property is used for resizing and the bounds of
// this layer change, the layer calls this method. The default implementation
// calls the [ResizeWithOldSuperlayerSize] method of each sublayer to let it
// know its superlayer’s bounds changed. You should not need to call or
// override this method directly.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/resizeSublayers(withOldSize:)
func (l CALayer) ResizeSublayersWithOldSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](l.ID, objc.Sel("resizeSublayersWithOldSize:"), size)
}
// Returns the preferred size of the layer in the coordinate space of its
// superlayer.
//
// # Return Value
// 
// The layer’s preferred frame size.
//
// # Discussion
// 
// In macOS, the default implementation of this method calls the
// `preferredSize()` method in Swift or the `` method in Objective-C of its
// layout manager—that is, the object in its [LayoutManager] property. If
// that object does not exist or does not implement that method, this method
// returns the size of the layer’s current [Bounds] rectangle mapped into
// the coordinate space of its [Superlayer].
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredFrameSize()
func (l CALayer) PreferredFrameSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](l.ID, objc.Sel("preferredFrameSize"))
	return corefoundation.CGSize(rv)
}
// Adds the specified constraint to the layer.
//
// c: The constraint object to add to the receiver’s array of constraint
// objects.
//
// # Discussion
// 
// In macOS, you typically add constraints to a layer to manage the size and
// position of that layer’s sublayers. Before constraints can be applied,
// you must also assign a [CAConstraintLayoutManager] object to the
// [LayoutManager] property of the layer. For more information about managing
// layer-based constraints, see [Core Animation Programming Guide].
// 
// iOS apps do not support layer-based constraints.
//
// [Core Animation Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004514
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/addConstraint(_:)
func (l CALayer) AddConstraint(c ICAConstraint) {
	objc.Send[objc.ID](l.ID, objc.Sel("addConstraint:"), c)
}
// Returns the action object assigned to the specified key.
//
// event: The identifier of the action.
//
// # Return Value
// 
// Returns the object that provides the action for `key`. The object must
// implement the [CAAction] protocol.
//
// # Discussion
// 
// This method searches for the given action object of the layer. Actions
// define dynamic behaviors for a layer. For example, the animatable
// properties of a layer typically have corresponding action objects to
// initiate the actual animations. When that property changes, the layer looks
// for the action object associated with the property name and executes it.
// You can also associate custom action objects with your layer to implement
// app-specific actions.
// 
// This method searches for the layer’s associated actions in the following
// order:
// 
// - If the layer has a delegate that implements the [ActionForLayerForKey]
// method, the layer calls that method. The delegate must do one of the
// following:
// 
// - Return the action object for the given key. - Return the [NSNull] object
// if it does not handle the action.
// 
// - The layer looks in the layer’s [Actions] dictionary for a matching
// key/action pair. - The layer looks in the [Style] dictionary for an
// [Actions] dictionary for a matching key/action pair. - The layer calls the
// [DefaultActionForKey] class method to look for any class-defined actions.
// 
// If any of the above steps returns an instance of [NSNull], it is converted
// to `nil` before continuing.
// 
// When an action object is invoked it receives three parameters: the name of
// the event, the object on which the event happened (the layer), and a
// dictionary of named arguments specific to each event kind.
//
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
func (l CALayer) ActionForKey(event string) CAAction {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("actionForKey:"), objc.String(event))
	return CAActionObjectFromID(rv)
}
// Converts the point from the specified layer’s coordinate system to the
// receiver’s coordinate system.
//
// p: A point specifying a location in the coordinate system of `l`.
//
// l: The layer with `p` in its coordinate system. The receiver and `l` and must
// share a common parent layer. This parameter may be `nil`.
//
// # Return Value
// 
// The point converted to the receiver’s coordinate system.
//
// # Discussion
// 
// If you specify `nil` for the `l` parameter, this method returns the
// original point subtracted from the layer’s frame’s origin.
// 
// The following example shows code that creates two layers, `redLayer` and
// `yellowLayer`. `yellowLayer` is scaled so that it is half of its original
// size.
// 
// The following figure shows the two layers and an overlaid point (rendered
// as a blue cross) with a position of `(50.0, 50.0)` in the red layer’s
// coordinate system.
// 
// [media-2850329]
// 
// The following code shows how you can find the coordinates of that point in
// the yellow layer’s coordinate system.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:from:)-8kl76
func (l_ CALayer) ConvertPointFromLayer(p corefoundation.CGPoint, l ICALayer) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l_.ID, objc.Sel("convertPoint:fromLayer:"), p, l)
	return corefoundation.CGPoint(rv)
}
// Converts the point from the receiver’s coordinate system to the specified
// layer’s coordinate system.
//
// p: A point specifying a location in the coordinate system of `l`.
//
// l: The layer into whose coordinate system `p` is to be converted. The receiver
// and `l` must share a common parent layer. This parameter may be `nil`.
//
// # Return Value
// 
// The point converted to the coordinate system of `layer`.
//
// # Discussion
// 
// If you specify `nil` for the `l` parameter, this method returns the
// original point added to the layer’s frame’s origin.
// 
// The following example shows code that creates two layers, `redLayer` and
// `yellowLayer`. `yellowLayer` is scaled so that it is half of its original
// size.
// 
// The following figure shows the two layers and an overlaid point (rendered
// as a blue cross) with a position of `(50.0, 50.0)` in the red layer’s
// coordinate system.
// 
// [media-2850332]
// 
// The following code shows how you can find the coordinates of that point in
// the yellow layer’s coordinate system.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:to:)-7dcke
func (l_ CALayer) ConvertPointToLayer(p corefoundation.CGPoint, l ICALayer) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l_.ID, objc.Sel("convertPoint:toLayer:"), p, l)
	return corefoundation.CGPoint(rv)
}
// Converts the rectangle from the specified layer’s coordinate system to
// the receiver’s coordinate system.
//
// r: A point specifying a location in the coordinate system of `l`.
//
// l: The layer with `r` in its coordinate system. The receiver and `l` and must
// share a common parent layer. This parameter may be `nil`.
//
// # Return Value
// 
// The rectangle converted to the receiver’s coordinate system.
//
// # Discussion
// 
// If you specify `nil` for the `l` parameter, this method returns the
// original rect with an origin subtracted from the layer’s frame’s
// origin.
// 
// The following example shows code that creates two layers, `redLayer` and
// `yellowLayer`. `yellowLayer` is scaled so that it is half of its original
// size.
// 
// The following figure shows the two layers and an overlaid rectangle with a
// frame of `(50, 50, 200, 200)` in the red layer’s coordinate system.
// 
// [media-2850326]
// 
// The following code shows how you can find the coordinates of that rectangle
// in the yellow layer’s coordinate system.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:from:)-4kx9l
func (l_ CALayer) ConvertRectFromLayer(r corefoundation.CGRect, l ICALayer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("convertRect:fromLayer:"), r, l)
	return corefoundation.CGRect(rv)
}
// Converts the rectangle from the receiver’s coordinate system to the
// specified layer’s coordinate system.
//
// r: A point specifying a location in the coordinate system of `l`.
//
// l: The layer into whose coordinate system `r` is to be converted. The receiver
// and `l` and must share a common parent layer. This parameter may be `nil`.
//
// # Return Value
// 
// The rectangle converted to the coordinate system of `l`.
//
// # Discussion
// 
// If you specify `nil` for the `l` parameter, this method returns the
// original rect with an origin added to the layer’s frame’s origin.
// 
// The following example shows code that creates two layers, `redLayer` and
// `yellowLayer`. `yellowLayer` is scaled so that it is half of its original
// size.
// 
// The following figure shows the two layers and an overlaid rectangle with a
// frame of `(50, 50, 200, 200)` in the red layer’s coordinate system.
// 
// [media-2850323]
// 
// The following code shows how you can find the coordinates of that rectangle
// in the yellow layer’s coordinate system.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:to:)-tly5
func (l_ CALayer) ConvertRectToLayer(r corefoundation.CGRect, l ICALayer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("convertRect:toLayer:"), r, l)
	return corefoundation.CGRect(rv)
}
// Converts the time interval from the specified layer’s time space to the
// receiver’s time space.
//
// t: A point specifying a location in the coordinate system of `l`.
//
// l: The layer with `t` in its time space. The receiver and `l` and must share a
// common parent layer.
//
// # Return Value
// 
// The time interval converted to the receiver’s time space.
//
// # Discussion
// 
// The following code shows the creation of two layers, layer and
// `offsetSlowMoLayer`. `offsetSlowMoLayer` has an offset time of 1 second and
// its [Speed] is set to `0.5`. The last line converts and prints a time
// interval of 0.5 seconds converted from the time space of `layer` to the
// time space of `offsetSlowMoLayer`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/convertTime(_:from:)
func (l_ CALayer) ConvertTimeFromLayer(t float64, l ICALayer) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("convertTime:fromLayer:"), t, l)
	return rv
}
// Converts the time interval from the receiver’s time space to the
// specified layer’s time space
//
// t: A point specifying a location in the coordinate system of `l`.
//
// l: The layer into whose time space `t` is to be converted. The receiver and
// `l` and must share a common parent layer.
//
// # Return Value
// 
// The time interval converted to the time space of `layer`.
//
// # Discussion
// 
// The following code shows the creation of two layers, `layer` and
// `offsetSlowMoLayer`. `offsetSlowMoLayer` has an offset time of 1 second and
// its [Speed] is set to `0.5`. The last line converts and prints a time
// interval of 0.5 seconds converted from the time space of `layer` to the
// time space of `offsetSlowMoLayer`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/convertTime(_:to:)
func (l_ CALayer) ConvertTimeToLayer(t float64, l ICALayer) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("convertTime:toLayer:"), t, l)
	return rv
}
// Returns the farthest descendant of the receiver in the layer hierarchy
// (including itself) that contains the specified point.
//
// p: A point in the coordinate system of the receiver’s superlayer.
//
// # Return Value
// 
// The layer that contains `thePoint` or `nil` if the point lies outside the
// receiver’s bounds rectangle.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/hitTest(_:)
func (l CALayer) HitTest(p corefoundation.CGPoint) ICALayer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("hitTest:"), p)
	return CALayerFromID(rv)
}
// Returns whether the receiver contains a specified point.
//
// p: A point in the receiver’s coordinate system.
//
// # Return Value
// 
// [true] if the bounds of the layer contains the point.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contains(_:)
func (l CALayer) ContainsPoint(p corefoundation.CGPoint) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("containsPoint:"), p)
	return rv
}
// Initiates a scroll in the layer’s closest ancestor scroll layer so that
// the specified point lies at the origin of the scroll layer.
//
// p: The point in the current layer that should be scrolled into position.
//
// # Discussion
// 
// If the layer is not contained by a [CAScrollLayer] object, this method does
// nothing.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/scroll(_:)
func (l CALayer) ScrollPoint(p corefoundation.CGPoint) {
	objc.Send[objc.ID](l.ID, objc.Sel("scrollPoint:"), p)
}
// Initiates a scroll in the layer’s closest ancestor scroll layer so that
// the specified rectangle becomes visible.
//
// r: The rectangle to be made visible.
//
// # Discussion
// 
// If the layer is not contained by a [CAScrollLayer] object, this method does
// nothing.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/scrollRectToVisible(_:)
func (l CALayer) ScrollRectToVisible(r corefoundation.CGRect) {
	objc.Send[objc.ID](l.ID, objc.Sel("scrollRectToVisible:"), r)
}
// Returns a Boolean indicating whether the value of the specified key should
// be archived.
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
// 
// [true] if the specified property should be archived or [false] if it should
// not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default implementation returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldArchiveValue(forKey:)
func (l CALayer) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("shouldArchiveValueForKey:"), objc.String(key))
	return rv
}
func (l CALayer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}
// Sets the value of the property identified by the given key. [Full Topic]
func (l CALayer) SetValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](l.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}
// Returns the value of the property identified by the given key. [Full Topic]
func (l CALayer) ValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("valueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns a Boolean indicating whether changes to the specified key require
// the layer to be redisplayed.
//
// key: A string that specifies an attribute of the layer.
//
// # Return Value
// 
// [true] if the layer requires a redisplay.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Subclasses can override this method and return [true] if the layer should
// be redisplayed when the value of the specified attribute changes.
// Animations changing the value of the attribute also trigger redisplay.
// 
// The default implementation of this method returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplay(forKey:)
func (_CALayerClass CALayerClass) NeedsDisplayForKey(key string) bool {
	rv := objc.Send[bool](objc.ID(_CALayerClass.class), objc.Sel("needsDisplayForKey:"), objc.String(key))
	return rv
}
// Returns the default action for the current class.
//
// event: The identifier of the action.
//
// # Return Value
// 
// Returns a suitable action object for the given key or `nil` of no action
// object was associated with that key.
//
// # Discussion
// 
// Classes that want to provide default actions can override this method and
// use it to return those actions.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/defaultAction(forKey:)
func (_CALayerClass CALayerClass) DefaultActionForKey(event string) CAAction {
	rv := objc.Send[objc.ID](objc.ID(_CALayerClass.class), objc.Sel("defaultActionForKey:"), objc.String(event))
	return CAActionObjectFromID(rv)
}
// Specifies the default value associated with the specified key.
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
// 
// The default value for the named property. Returns `nil` if no default value
// has been set.
//
// # Discussion
// 
// If you define custom properties for a layer but do not set a value, this
// method returns a suitable “zero” default value based on the expected
// value of the `key`. For example, if the value for `key` is a [CGSize]
// struct, the method returns a size struct containing (0.0,0.0) wrapped in an
// [NSValue] object. For a [CGRect] an empty rectangle is returned. For
// [CGAffineTransform] and [CATransform3D], the appropriate identity matrix is
// returned.
// 
// # Special Considerations
// 
// If `key` is not a known for property of the class, the result of the method
// is undefined.
//
// [CATransform3D]: https://developer.apple.com/documentation/QuartzCore/CATransform3D
// [CGAffineTransform]: https://developer.apple.com/documentation/CoreFoundation/CGAffineTransform
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
// [CGSize]: https://developer.apple.com/documentation/CoreFoundation/CGSize
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/defaultValue(forKey:)
func (_CALayerClass CALayerClass) DefaultValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_CALayerClass.class), objc.Sel("defaultValueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurveExpansionFactor(_:)
func (_CALayerClass CALayerClass) CornerCurveExpansionFactor(curve CALayerCornerCurve) float64 {
	rv := objc.Send[float64](objc.ID(_CALayerClass.class), objc.Sel("cornerCurveExpansionFactor:"), objc.String(string(curve)))
	return rv
}

// The layer’s delegate object.
//
// # Discussion
// 
// You can use a delegate object to provide the layer’s contents, handle the
// layout of any sublayers, and provide custom actions in response to
// layer-related changes. The object you assign to this property should
// implement one or more of the methods of the [CALayerDelegate] informal
// protocol. For more information about that protocol, see [CALayerDelegate]
// 
// In iOS, if the layer is associated with a [UIView] object, this property be
// set to the view that owns the layer.
//
// [UIView]: https://developer.apple.com/documentation/UIKit/UIView
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l CALayer) Delegate() CALayerDelegate {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("delegate"))
	return CALayerDelegateObjectFromID(rv)
}
func (l CALayer) SetDelegate(value CALayerDelegate) {
	objc.Send[struct{}](l.ID, objc.Sel("setDelegate:"), value)
}
// An object that provides the contents of the layer. Animatable.
//
// # Discussion
// 
// The default value of this property is `nil`.
// 
// If you are using the layer to display a static image, you can set this
// property to the [CGImage] containing the image you want to display. (In
// macOS 10.6 and later, you can also set the property to an [NSImage]
// object.) Assigning a value to this property causes the layer to use your
// image rather than create a separate backing store.
// 
// If the layer object is tied to a view object, you should avoid setting the
// contents of this property directly. The interplay between views and layers
// usually results in the view replacing the contents of this property during
// a subsequent update.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (l CALayer) Contents() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("contents"))
	return objectivec.Object{ID: rv}
}
func (l CALayer) SetContents(value objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setContents:"), value)
}
// The rectangle, in the unit coordinate space, that defines the portion of
// the layer’s contents that should be used. Animatable.
//
// # Discussion
// 
// Defaults to the unit rectangle (0.0, 0.0, 1.0, 1.0).
// 
// If pixels outside the unit rectangle are requested, the edge pixels of the
// contents image will be extended outwards.
// 
// If an empty rectangle is provided, the results are undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsRect
func (l CALayer) ContentsRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("contentsRect"))
	return corefoundation.CGRect(rv)
}
func (l CALayer) SetContentsRect(value corefoundation.CGRect) {
	objc.Send[struct{}](l.ID, objc.Sel("setContentsRect:"), value)
}
// The rectangle that defines how the layer contents are scaled if the
// layer’s contents are resized. Animatable.
//
// # Discussion
// 
// You can use this property to subdivide the layer’s content into a 3x3
// grid. The value in this property specifies the location and size of the
// center rectangle in that grid. If the layer’s [ContentsGravity] property
// is set to one of the resizing modes, resizing the layer causes scaling to
// occur differently in each rectangle of the grid. The center rectangle is
// stretched in both dimensions, the top-center and bottom-center rectangles
// are stretched only horizontally, the left-center and right-center
// rectangles are stretched only vertically, and the four corner rectangles
// are not stretched at all. Therefore, you can use this technique to
// implement stretchable backgrounds or images using a three-part or nine-part
// image.
// 
// The value in this property is set to the unit rectangle `(0.0,0.0)
// (1.0,1.0)` by default, which causes the entire image to scale in both
// dimensions. If you specify a rectangle that extends outside the unit
// rectangle, the result is undefined. The rectangle you specify is applied
// only after the [ContentsRect] property has been applied to the image.
// 
// The following code shows how you can create a [CALayer] with the button
// image shown in the following figure set as its contents. The corner radii
// of the button image are set to one quarter of the length of its side.
// 
// [media-2852114]
// 
// By setting the layer’s [ContentsCenter] to `(0.25,0.25) (0.5,0.5)`, the
// button’s corner radius remains unchanged, whatever size the layer is set
// to.
// 
// The following figure shows the layer created in the previous code resized
// to 400 x 400, 200 x 200 and 100 x 100 points.
// 
// [media-2852112]
// 
// The following figure shows the layer at the same sizes but without
// explicitly setting the layer’s [ContentsCenter]: the entire button image
// is scaled, including the rounded corners.
// 
// [media-2852113]
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsCenter
func (l CALayer) ContentsCenter() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("contentsCenter"))
	return corefoundation.CGRect(rv)
}
func (l CALayer) SetContentsCenter(value corefoundation.CGRect) {
	objc.Send[struct{}](l.ID, objc.Sel("setContentsCenter:"), value)
}
// A constant that specifies how the layer’s contents are positioned or
// scaled within its bounds.
//
// # Discussion
// 
// The possible values for this property are listed in [Contents Gravity
// Values].
// 
// The default value of this property is [resize].
// 
// [Figure 1] shows four examples of the effect of setting different values
// for a layer’s [ContentsGravity] property.
// 
// [media-2851774]
// 
// - Contents gravity is [resize] - the default - Contents gravity is [center]
// - Contents gravity is [ContentsAreFlipped] `?` [top] : [bottom] - Contents
// gravity is [ContentsAreFlipped] `?` [bottomLeft] : [topLeft]
//
// [Contents Gravity Values]: https://developer.apple.com/documentation/QuartzCore/contents-gravity-values
// [bottomLeft]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/bottomLeft
// [bottom]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/bottom
// [center]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/center
// [resize]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resize
// [topLeft]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/topLeft
// [top]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/top
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (l CALayer) ContentsGravity() CALayerContentsGravity {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("contentsGravity"))
	return CALayerContentsGravity(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetContentsGravity(value CALayerContentsGravity) {
	objc.Send[struct{}](l.ID, objc.Sel("setContentsGravity:"), objc.String(string(value)))
}
// The opacity of the receiver. Animatable.
//
// # Discussion
// 
// The value of this property must be in the range `0.0` (transparent) to
// `1.0` (opaque). Values outside that range are clamped to the minimum or
// maximum. The default value of this property is `1.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l CALayer) Opacity() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("opacity"))
	return rv
}
func (l CALayer) SetOpacity(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setOpacity:"), value)
}
// A Boolean indicating whether the layer is displayed. Animatable.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/isHidden
func (l CALayer) Hidden() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isHidden"))
	return rv
}
func (l CALayer) SetHidden(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setHidden:"), value)
}
// A Boolean indicating whether sublayers are clipped to the layer’s bounds.
// Animatable.
//
// # Discussion
// 
// When the value of this property is [true], Core Animation creates an
// implicit clipping mask that matches the bounds of the layer and includes
// any corner radius effects. If a value for the [Mask] property is also
// specified, the two masks are multiplied to get the final mask value.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/masksToBounds
func (l CALayer) MasksToBounds() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("masksToBounds"))
	return rv
}
func (l CALayer) SetMasksToBounds(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setMasksToBounds:"), value)
}
// An optional layer whose alpha channel is used to mask the layer’s
// content.
//
// # Discussion
// 
// The layer’s alpha channel determines how much of the layer’s content
// and background shows through. Fully or partially opaque pixels allow the
// underlying content to show through, but fully transparent pixels block that
// content.
// 
// The default value of this property is `nil`. When configuring a mask,
// remember to set the size and position of the mask layer to ensure it is
// aligned properly with the layer it masks.
// 
// # Special Considerations
// 
// The layer you assign to this property must not have a superlayer. If it
// does, the behavior is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/mask
func (l CALayer) Mask() ICALayer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("mask"))
	return CALayerFromID(objc.ID(rv))
}
func (l CALayer) SetMask(value ICALayer) {
	objc.Send[struct{}](l.ID, objc.Sel("setMask:"), value)
}
// A Boolean indicating whether the layer displays its content when facing
// away from the viewer. Animatable.
//
// # Discussion
// 
// When the value in this property is [false], the layer hides its content
// when it faces away from the viewer. The default value of this property is
// [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/isDoubleSided
func (l CALayer) DoubleSided() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isDoubleSided"))
	return rv
}
func (l CALayer) SetDoubleSided(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setDoubleSided:"), value)
}
// The radius to use when drawing rounded corners for the layer’s
// background. Animatable.
//
// # Discussion
// 
// Setting the radius to a value greater than `0.0` causes the layer to begin
// drawing rounded corners on its background. By default, the corner radius
// does not apply to the image in the layer’s [Contents] property; it
// applies only to the background color and border of the layer. However,
// setting the [MasksToBounds] property to [true] causes the content to be
// clipped to the rounded corners.
// 
// The default value of this property is `0.0`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerRadius
func (l CALayer) CornerRadius() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("cornerRadius"))
	return rv
}
func (l CALayer) SetCornerRadius(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setCornerRadius:"), value)
}
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/maskedCorners
func (l CALayer) MaskedCorners() CACornerMask {
	rv := objc.Send[CACornerMask](l.ID, objc.Sel("maskedCorners"))
	return CACornerMask(rv)
}
func (l CALayer) SetMaskedCorners(value CACornerMask) {
	objc.Send[struct{}](l.ID, objc.Sel("setMaskedCorners:"), value)
}
// The width of the layer’s border. Animatable.
//
// # Discussion
// 
// When this value is greater than 0.0, the layer draws a border using the
// current [BorderColor] value. The border is drawn inset from the
// receiver’s bounds by the value specified in this property. It is
// composited above the receiver’s contents and sublayers and includes the
// effects of the [CornerRadius] property.
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/borderWidth
func (l CALayer) BorderWidth() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("borderWidth"))
	return rv
}
func (l CALayer) SetBorderWidth(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setBorderWidth:"), value)
}
// The color of the layer’s border. Animatable.
//
// # Discussion
// 
// The default value of this property is an opaque black color.
// 
// The value of this property is retained using the Core Foundation
// retain/release semantics. This behavior occurs despite the fact that the
// property declaration appears to use the default assign semantics for object
// retention.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/borderColor
func (l CALayer) BorderColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](l.ID, objc.Sel("borderColor"))
	return coregraphics.CGColorRef(rv)
}
func (l CALayer) SetBorderColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](l.ID, objc.Sel("setBorderColor:"), value)
}
// The background color of the receiver. Animatable.
//
// # Discussion
// 
// The default value of this property is `nil`.
// 
// The value of this property is retained using the Core Foundation
// retain/release semantics. This behavior occurs despite the fact that the
// property declaration appears to use the default assign semantics for object
// retention.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundColor
func (l CALayer) BackgroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](l.ID, objc.Sel("backgroundColor"))
	return coregraphics.CGColorRef(rv)
}
func (l CALayer) SetBackgroundColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](l.ID, objc.Sel("setBackgroundColor:"), value)
}
// The opacity of the layer’s shadow. Animatable.
//
// # Discussion
// 
// The value in this property must be in the range `0.0` (transparent) to
// `1.0` (opaque). The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOpacity
func (l CALayer) ShadowOpacity() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("shadowOpacity"))
	return rv
}
func (l CALayer) SetShadowOpacity(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setShadowOpacity:"), value)
}
// The blur radius (in points) used to render the layer’s shadow.
// Animatable.
//
// # Discussion
// 
// You specify the radius The default value of this property is 3.0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowRadius
func (l CALayer) ShadowRadius() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("shadowRadius"))
	return rv
}
func (l CALayer) SetShadowRadius(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setShadowRadius:"), value)
}
// The offset (in points) of the layer’s shadow. Animatable.
//
// # Discussion
// 
// The default value of this property is (`0.0`, `-3.0`).
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOffset
func (l CALayer) ShadowOffset() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](l.ID, objc.Sel("shadowOffset"))
	return corefoundation.CGSize(rv)
}
func (l CALayer) SetShadowOffset(value corefoundation.CGSize) {
	objc.Send[struct{}](l.ID, objc.Sel("setShadowOffset:"), value)
}
// The color of the layer’s shadow. Animatable.
//
// # Discussion
// 
// The default value of this property is an opaque black color.
// 
// The value of this property is retained using the Core Foundation
// retain/release semantics. This behavior occurs despite the fact that the
// property declaration appears to use the default assign semantics for object
// retention.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowColor
func (l CALayer) ShadowColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](l.ID, objc.Sel("shadowColor"))
	return coregraphics.CGColorRef(rv)
}
func (l CALayer) SetShadowColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](l.ID, objc.Sel("setShadowColor:"), value)
}
// The shape of the layer’s shadow. Animatable.
//
// # Discussion
// 
// The default value of this property is `nil`, which causes the layer to use
// a standard shadow shape. If you specify a value for this property, the
// layer creates its shadow using the specified path instead of the layer’s
// composited alpha channel. The path you provide defines the outline of the
// shadow. It is filled using the non-zero winding rule and the current shadow
// color, opacity, and blur radius.
// 
// Unlike most animatable properties, this property (as with all [CGPathRef]
// animatable properties) does not support implicit animation. However, the
// path object may be animated using any of the concrete subclasses of
// [CAPropertyAnimation]. Paths will interpolate as a linear blend of the
// “on-line” points; “off-line” points may be interpolated
// non-linearly (to preserve continuity of the curve’s derivative). If the
// two paths have a different number of control points or segments, the
// results are undefined. If the path extends outside the layer bounds it will
// not automatically be clipped to the layer, only if the normal layer masking
// rules cause that.
// 
// Specifying an explicit path usually improves rendering performance.
// 
// The value of this property is retained using the Core Foundation
// retain/release semantics. This behavior occurs despite the fact that the
// property declaration appears to use the default assign semantics for object
// retention.
// 
// # Using Shadow Path for Special Effects
// 
// You can use a layer’s shadow path to create special effects such as
// simulating the shadows available in [Pages].
// 
// The following code shows the code required to add an elliptical shadow to
// the bottom of a layer to simulate the Pages effect.
// 
// [media-2851604]
// 
// The following code shows how to create a path to simulate the Pages . The
// left, top and right sides of the path are straight lines, and the bottom is
// a concave curve as illustrated in the following figure.
// 
// [media-2851608]
// 
// [media-2851607]
//
// [Pages]: http://www.apple.com/pages/
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowPath
func (l CALayer) ShadowPath() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](l.ID, objc.Sel("shadowPath"))
	return coregraphics.CGPathRef(rv)
}
func (l CALayer) SetShadowPath(value coregraphics.CGPathRef) {
	objc.Send[struct{}](l.ID, objc.Sel("setShadowPath:"), value)
}
// An optional dictionary used to store property values that aren’t
// explicitly defined by the layer.
//
// # Discussion
// 
// This dictionary may in turn have a `style` key, forming a hierarchy of
// default values. In the case of hierarchical style dictionaries the
// shallowest value for a property is used. For example, the value for
// “style.someValue” takes precedence over “style.style.someValue”.
// 
// If the style dictionary does not define a value for an attribute, the
// receiver’s [DefaultValueForKey] method is called. The default value of
// this property is `nil`.
// 
// The style dictionary is not consulted for the following keys: `bounds`,
// `frame`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/style
func (l CALayer) Style() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("style"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (l CALayer) SetStyle(value foundation.INSDictionary) {
	objc.Send[struct{}](l.ID, objc.Sel("setStyle:"), value)
}
// A Boolean indicating whether the layer is allowed to perform edge
// antialiasing.
//
// # Discussion
// 
// When the value is [true], the layer is allowed to antialias its edges, as
// requested by the value in the layer’s [EdgeAntialiasingMask] property.
// The default value is read from the boolean [UIViewEdgeAntialiasing]
// property in the main bundle’s `Info.Plist()` file. If no value is found,
// the default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsEdgeAntialiasing
func (l CALayer) AllowsEdgeAntialiasing() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("allowsEdgeAntialiasing"))
	return rv
}
func (l CALayer) SetAllowsEdgeAntialiasing(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setAllowsEdgeAntialiasing:"), value)
}
// A Boolean indicating whether the layer is allowed to composite itself as a
// group separate from its parent.
//
// # Discussion
// 
// When the value is [true] and the layer’s opacity property value is less
// than `1.0`, the layer is allowed to composite itself as a group separate
// from its parent. This gives correct results when the layer contains
// multiple opaque components, but may reduce performance.
// 
// The default value is read from the boolean [UIViewGroupOpacity] property in
// the main bundle’s `Info.Plist()` file. If no value is found, the default
// value is [true] for apps linked against the iOS 7 SDK or later and [false]
// for apps linked against an earlier SDK.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsGroupOpacity
func (l CALayer) AllowsGroupOpacity() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("allowsGroupOpacity"))
	return rv
}
func (l CALayer) SetAllowsGroupOpacity(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setAllowsGroupOpacity:"), value)
}
// An array of Core Image filters to apply to the contents of the layer and
// its sublayers. Animatable.
//
// # Discussion
// 
// The filters you add to this property affect the content of the layer,
// including its border, filled background and sublayers. The default value of
// this property is `nil`.
// 
// Changing the inputs of the [CIFilter] object directly after it is attached
// to the layer causes undefined behavior. It is possible to modify filter
// parameters after attaching them to the layer but you must use the layer’s
// [setValue(_:forKeyPath:)] method to do so. In addition, you must assign a
// name to the filter so that you can identify it in the array. For example,
// to change the `inputRadius` parameter of the filter, you could use code
// similar to the following:
// 
// The following code shows how to create a text layer and apply a
// [CIPointillize] filter to it.
// 
// The following figure shows the result: a pointillist effect is added to the
// text.
// 
// [media-2851431]
// 
// # Special Considerations
// 
// This property is not supported on layers in iOS.
//
// [CIFilter]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
// [CIPointillize]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CIPointillize
// [setValue(_:forKeyPath:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKeyPath:)
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/filters
func (l CALayer) Filters() foundation.INSArray {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("filters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (l CALayer) SetFilters(value foundation.INSArray) {
	objc.Send[struct{}](l.ID, objc.Sel("setFilters:"), value)
}
// A CoreImage filter used to composite the layer and the content behind it.
// Animatable.
//
// # Discussion
// 
// The default value of this property is `nil`, which causes the layer to use
// source-over compositing. Although you can use any Core Image filter as a
// layer’s compositing filter, for best results, use those in the
// [CICategoryCompositeOperation] category.
// 
// In macOS, it is possible to modify the filter’s parameters after
// attaching it to the layer but you must use the layer’s
// [setValue(_:forKeyPath:)] method to do so. For example, to change the
// `inputRadius` parameter of the filter, you could use code similar to the
// following:
// 
// Changing the inputs of the [CIFilter] object directly after it is attached
// to the layer causes undefined behavior.
// 
// The following code shows how to create two overlapping text layers,
// background and foreground. Addition compositing is used to composite the
// foreground over the background.
// 
// The following figure shows the result: the identical background colors of
// the two layers are added together so that a brighter gray is produced where
// the layers overlap.
// 
// [media-2851428]
// 
// The following figure shows the default result when the foreground layer’s
// compositing filter is `nil` or [CISourceOverCompositing].
// 
// [media-2851429]
// 
// # Special Considerations
// 
// This property is not supported on layers in iOS.
//
// [CICategoryCompositeOperation]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/uid/TP30000136-SW71
// [CIFilter]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
// [CISourceOverCompositing]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CISourceOverCompositing
// [setValue(_:forKeyPath:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKeyPath:)
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/compositingFilter
func (l CALayer) CompositingFilter() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("compositingFilter"))
	return objectivec.Object{ID: rv}
}
func (l CALayer) SetCompositingFilter(value objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setCompositingFilter:"), value)
}
// An array of Core Image filters to apply to the content immediately behind
// the layer. Animatable.
//
// # Discussion
// 
// Background filters affect the content behind the layer that shows through
// into the layer itself. Typically this content belongs to the superlayer
// that acts as the parent of the layer. These filters do not affect the
// content of the layer itself, including the layer’s background color and
// border.
// 
// The default value of this property is `nil`.
// 
// Changing the inputs of the [CIFilter] object directly after it is attached
// to the layer causes undefined behavior. In macOS, it is possible to modify
// filter parameters after attaching them to the layer but you must use the
// layer’s [setValue(_:forKeyPath:)] method to do so. In addition, you must
// assign a name to the filter so that you can identify it in the array. For
// example, to change the `inputRadius` parameter of the filter, you could use
// code similar to the following:
// 
// You use the layer’s [MasksToBounds] to control the extent of its
// background filter’s effect.
// 
// The following code shows how to create two overlapping text layers,
// `background` and `foreground`. A Gaussian blur filter is added to the
// foreground layer’s [BackgroundFilters] array and its [MasksToBounds] is
// set to [true]:
// 
// The following figure shows the result: the background layer is only blurred
// where it is overlapped by the foreground layer.
// 
// [media-2851423]
// 
// However, if the foreground layer’s [MasksToBounds] is set to [false], the
// entire background layer is blurred as illustrated in the following figure.
// 
// [media-2851424]
// 
// # Special Considerations
// 
// This property is not supported on layers in iOS.
//
// [CIFilter]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
// [false]: https://developer.apple.com/documentation/Swift/false
// [setValue(_:forKeyPath:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKeyPath:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundFilters
func (l CALayer) BackgroundFilters() foundation.INSArray {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("backgroundFilters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (l CALayer) SetBackgroundFilters(value foundation.INSArray) {
	objc.Send[struct{}](l.ID, objc.Sel("setBackgroundFilters:"), value)
}
// The filter used when reducing the size of the content.
//
// # Discussion
// 
// The possible values for this property are listed in [Scaling Filters].
// 
// The default value of this property is [linear].
//
// [Scaling Filters]: https://developer.apple.com/documentation/QuartzCore/scaling-filters
// [linear]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/linear
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilter
func (l CALayer) MinificationFilter() CALayerContentsFilter {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("minificationFilter"))
	return CALayerContentsFilter(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetMinificationFilter(value CALayerContentsFilter) {
	objc.Send[struct{}](l.ID, objc.Sel("setMinificationFilter:"), objc.String(string(value)))
}
// The bias factor used by the minification filter to determine the levels of
// detail.
//
// # Discussion
// 
// This value is used by the [MinificationFilter] when it is set to
// [trilinear].
// 
// The default value of this property is `0.0`.
//
// [trilinear]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/trilinear
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilterBias
func (l CALayer) MinificationFilterBias() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("minificationFilterBias"))
	return rv
}
func (l CALayer) SetMinificationFilterBias(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setMinificationFilterBias:"), value)
}
// The filter used when increasing the size of the content.
//
// # Discussion
// 
// The possible values for this property are listed in [Scaling Filters].
// 
// The default value of this property is [linear].
// 
// [Figure 1] shows the difference between linear and nearest filtering when a
// 10 x 10 point image of a circle is magnified by a scale of 10.
// 
// [media-2851435]
// 
// The circle on the left uses [linear] and the circle on the right uses
// [nearest].
//
// [Scaling Filters]: https://developer.apple.com/documentation/QuartzCore/scaling-filters
// [linear]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/linear
// [nearest]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/nearest
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/magnificationFilter
func (l CALayer) MagnificationFilter() CALayerContentsFilter {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("magnificationFilter"))
	return CALayerContentsFilter(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetMagnificationFilter(value CALayerContentsFilter) {
	objc.Send[struct{}](l.ID, objc.Sel("setMagnificationFilter:"), objc.String(string(value)))
}
// A Boolean value indicating whether the layer contains completely opaque
// content.
//
// # Discussion
// 
// The default value of this property is [false]. If your app draws completely
// opaque content that fills the layer’s bounds, setting this property to
// [true] lets the system optimize the rendering behavior for the layer.
// Specifically, when the layer creates the backing store for your drawing
// commands, Core Animation omits the alpha channel of that backing store.
// Doing so can improve the performance of compositing operations. If you set
// the value of this property to [true], you must fill the layer’s bounds
// with opaque content.
// 
// Setting this property affects only the backing store managed by Core
// Animation. If you assign an image with an alpha channel to the layer’s
// [Contents] property, that image retains its alpha channel regardless of the
// value of this property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/isOpaque
func (l CALayer) Opaque() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isOpaque"))
	return rv
}
func (l CALayer) SetOpaque(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setOpaque:"), value)
}
// A bitmask defining how the edges of the receiver are rasterized.
//
// # Discussion
// 
// This property specifies which edges of the layer are antialiased and is a
// combination of the constants defined in [CAEdgeAntialiasingMask]. You can
// enable or disable antialiasing for each edge (top, left, bottom, right)
// separately. By default antialiasing is enabled for all edges.
// 
// Typically, you would use this property to disable antialiasing for edges
// that abut edges of other layers, to eliminate the seams that would
// otherwise occur.
//
// [CAEdgeAntialiasingMask]: https://developer.apple.com/documentation/QuartzCore/CAEdgeAntialiasingMask
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/edgeAntialiasingMask
func (l CALayer) EdgeAntialiasingMask() CAEdgeAntialiasingMask {
	rv := objc.Send[CAEdgeAntialiasingMask](l.ID, objc.Sel("edgeAntialiasingMask"))
	return CAEdgeAntialiasingMask(rv)
}
func (l CALayer) SetEdgeAntialiasingMask(value CAEdgeAntialiasingMask) {
	objc.Send[struct{}](l.ID, objc.Sel("setEdgeAntialiasingMask:"), value)
}
// A Boolean that indicates whether the geometry of the layer and its
// sublayers is flipped vertically.
//
// # Discussion
// 
// If the layer is providing the backing for a layer-backed view, the view is
// responsible for managing the value in this property. For standalone layers,
// this property controls whether geometry values for the layer are
// interpreted using the standard or flipped coordinate system. The value of
// this property does not affect the rendering of the layer’s content.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/isGeometryFlipped
func (l CALayer) GeometryFlipped() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isGeometryFlipped"))
	return rv
}
func (l CALayer) SetGeometryFlipped(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setGeometryFlipped:"), value)
}
// A Boolean indicating whether drawing commands are deferred and processed
// asynchronously in a background thread.
//
// # Discussion
// 
// When this property is set to [true], the graphics context used to draw the
// layer’s contents queues drawing commands and executes them on a
// background thread rather than executing them synchronously. Performing
// these commands asynchronously can improve performance in some apps.
// However, you should always measure the actual performance benefits before
// enabling this capability.
// 
// The default value for this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/drawsAsynchronously
func (l CALayer) DrawsAsynchronously() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("drawsAsynchronously"))
	return rv
}
func (l CALayer) SetDrawsAsynchronously(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setDrawsAsynchronously:"), value)
}
// A Boolean that indicates whether the layer is rendered as a bitmap before
// compositing. Animatable
//
// # Discussion
// 
// When the value of this property is [true], the layer is rendered as a
// bitmap in its local coordinate space and then composited to the destination
// with any other content. Shadow effects and any filters in the [Filters]
// property are rasterized and included in the bitmap. However, the current
// opacity of the layer is not rasterized. If the rasterized bitmap requires
// scaling during compositing, the filters in the [MinificationFilter] and
// [MagnificationFilter] properties are applied as needed.
// 
// When the value of this property is [false], the layer is composited
// directly into the destination whenever possible. The layer may still be
// rasterized prior to compositing if certain features of the compositing
// model (such as the inclusion of filters) require it.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldRasterize
func (l CALayer) ShouldRasterize() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("shouldRasterize"))
	return rv
}
func (l CALayer) SetShouldRasterize(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setShouldRasterize:"), value)
}
// The scale at which to rasterize content, relative to the coordinate space
// of the layer. Animatable
//
// # Discussion
// 
// When the value in the [ShouldRasterize] property is [true], the layer uses
// this property to determine whether to scale the rasterized content (and by
// how much). The default value of this property is `1.0`, which indicates
// that the layer should be rasterized at its current size. Larger values
// magnify the content and smaller values shrink it.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/rasterizationScale
func (l CALayer) RasterizationScale() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("rasterizationScale"))
	return rv
}
func (l CALayer) SetRasterizationScale(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setRasterizationScale:"), value)
}
// A hint for the desired storage format of the layer contents.
//
// # Discussion
// 
// The default value of this property is [RGBA8Uint].
// 
// [UIView] and layer-backed [NSView] objects may change the value to a format
// appropriate for the current device.
//
// [NSView]: https://developer.apple.com/documentation/AppKit/NSView
// [RGBA8Uint]: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFormat/RGBA8Uint
// [UIView]: https://developer.apple.com/documentation/UIKit/UIView
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsFormat
func (l CALayer) ContentsFormat() CALayerContentsFormat {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("contentsFormat"))
	return CALayerContentsFormat(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetContentsFormat(value CALayerContentsFormat) {
	objc.Send[struct{}](l.ID, objc.Sel("setContentsFormat:"), objc.String(string(value)))
}
// The layer’s frame rectangle.
//
// # Discussion
// 
// The frame rectangle is position and size of the layer specified in the
// superlayer’s coordinate space. For layers, the frame rectangle is a
// computed property that is derived from the values in the[Bounds],
// [AnchorPoint] and [Position] properties. When you assign a new value to
// this property, the layer changes its [Position] and [Bounds] properties to
// match the rectangle you specified. The values of each coordinate in the
// rectangle are measured in points.
// 
// Do not set the frame if the [Transform] property applies a rotation
// transform that is not a multiple of 90 degrees.
// 
// For more information about the relationship between the [Frame], [Bounds],
// [AnchorPoint] and [Position] properties, see [Core Animation Programming
// Guide].
//
// [Core Animation Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004514
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/frame
func (l CALayer) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
func (l CALayer) SetFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](l.ID, objc.Sel("setFrame:"), value)
}
// The layer’s bounds rectangle. Animatable.
//
// # Discussion
// 
// The bounds rectangle is the origin and size of the layer in its own
// coordinate space. When you create a new standalone layer, the default value
// for this property is an empty rectangle, which you must change before using
// the layer. The values of each coordinate in the rectangle are measured in
// points.
// 
// For more information about the relationship between the [Frame], [Bounds],
// [AnchorPoint] and [Position] properties, see [Core Animation Programming
// Guide].
//
// [Core Animation Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004514
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/bounds
func (l CALayer) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (l CALayer) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](l.ID, objc.Sel("setBounds:"), value)
}
// The layer’s position in its superlayer’s coordinate space. Animatable.
//
// # Discussion
// 
// The value of this property is specified in points and is always specified
// relative to the value in the [AnchorPoint] property. For new standalone
// layers, the default position is set to (0.0, 0.0). Changing the [Frame]
// property also updates the value in this property.
// 
// For more information about the relationship between the [Frame], [Bounds],
// [AnchorPoint] and [Position] properties, see [Core Animation Programming
// Guide].
//
// [Core Animation Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004514
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/position
func (l CALayer) Position() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l.ID, objc.Sel("position"))
	return corefoundation.CGPoint(rv)
}
func (l CALayer) SetPosition(value corefoundation.CGPoint) {
	objc.Send[struct{}](l.ID, objc.Sel("setPosition:"), value)
}
// The layer’s position on the z axis. Animatable.
//
// # Discussion
// 
// The default value of this property is `0`. Changing the value of this
// property changes the front-to-back ordering of layers onscreen. Higher
// values place this layer visually closer to the viewer than layers with
// lower values. This can affect the visibility of layers whose frame
// rectangles overlap.
// 
// The value of this property is measured in points. The range of this
// property is single-precision, floating-point `-`[greatestFiniteMagnitude]
// to [greatestFiniteMagnitude].
//
// [greatestFiniteMagnitude]: https://developer.apple.com/documentation/Swift/Float/greatestFiniteMagnitude
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/zPosition
func (l CALayer) ZPosition() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("zPosition"))
	return rv
}
func (l CALayer) SetZPosition(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setZPosition:"), value)
}
// The anchor point for the layer’s position along the z axis. Animatable.
//
// # Discussion
// 
// This property specifies the anchor point on the z axis around which
// geometric manipulations occur. The point is expressed as a distance
// (measured in points) along the z axis. The default value of this property
// is `0`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPointZ
func (l CALayer) AnchorPointZ() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("anchorPointZ"))
	return rv
}
func (l CALayer) SetAnchorPointZ(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setAnchorPointZ:"), value)
}
// Defines the anchor point of the layer’s bounds rectangle. Animatable.
//
// # Discussion
// 
// You specify the value for this property using the unit coordinate space.
// The default value of this property is (0.5, 0.5), which represents the
// center of the layer’s bounds rectangle. All geometric manipulations to
// the view occur about the specified point. For example, applying a rotation
// transform to a layer with the default anchor point causes the layer to
// rotate around its center. Changing the anchor point to a different location
// would cause the layer to rotate around that new point.
// 
// For more information about the relationship between the [Frame], [Bounds],
// [AnchorPoint] and [Position] properties, see [Core Animation Programming
// Guide].
//
// [Core Animation Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004514
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPoint
func (l CALayer) AnchorPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l.ID, objc.Sel("anchorPoint"))
	return corefoundation.CGPoint(rv)
}
func (l CALayer) SetAnchorPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](l.ID, objc.Sel("setAnchorPoint:"), value)
}
// The scale factor applied to the layer.
//
// # Discussion
// 
// This value defines the mapping between the logical coordinate space of the
// layer (measured in points) and the physical coordinate space (measured in
// pixels). Higher scale factors indicate that each point in the layer is
// represented by more than one pixel at render time. For example, if the
// scale factor is `2.0` and the layer’s bounds are 50 x 50 points, the size
// of the bitmap used to present the layer’s content is 100 x 100 pixels.
// 
// The default value of this property is 1.0. For layers attached to a view,
// the view changes the scale factor automatically to a value that is
// appropriate for the current screen. For layers you create and manage
// yourself, you must set the value of this property yourself based on the
// resolution of the screen and the content you are providing. Core Animation
// uses the value you specify as a cue to determine how to render your
// content.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
func (l CALayer) ContentsScale() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("contentsScale"))
	return rv
}
func (l CALayer) SetContentsScale(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setContentsScale:"), value)
}
// The transform applied to the layer’s contents. Animatable.
//
// # Discussion
// 
// This property is set to the identity transform by default. Any
// transformations you apply to the layer occur relative to the layer’s
// anchor point.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/transform
func (l CALayer) Transform() CATransform3D {
	rv := objc.Send[CATransform3D](l.ID, objc.Sel("transform"))
	return CATransform3D(rv)
}
func (l CALayer) SetTransform(value CATransform3D) {
	objc.Send[struct{}](l.ID, objc.Sel("setTransform:"), value)
}
// Specifies the transform to apply to sublayers when rendering. Animatable.
//
// # Discussion
// 
// You typically use this property to add perspective and other viewing
// effects to embedded layers. You add perspective by setting the sublayer
// transform to the desired projection matrix. The default value of this
// property is the identity transform.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayerTransform
func (l CALayer) SublayerTransform() CATransform3D {
	rv := objc.Send[CATransform3D](l.ID, objc.Sel("sublayerTransform"))
	return CATransform3D(rv)
}
func (l CALayer) SetSublayerTransform(value CATransform3D) {
	objc.Send[struct{}](l.ID, objc.Sel("setSublayerTransform:"), value)
}
// An array containing the layer’s sublayers.
//
// # Discussion
// 
// The sublayers are listed in back to front order. The default value of this
// property is `nil`.
// 
// # Special Considerations
// 
// When setting the [Sublayers] property to an array populated with layer
// objects, each layer in the array must not already have a superlayer—that
// is, its [Superlayer] property must currently be `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayers
func (l CALayer) Sublayers() []CALayer {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("sublayers"))
	return objc.ConvertSlice(rv, func(id objc.ID) CALayer {
		return CALayerFromID(id)
	})
}
func (l CALayer) SetSublayers(value []CALayer) {
	objc.Send[struct{}](l.ID, objc.Sel("setSublayers:"), objectivec.IObjectSliceToNSArray(value))
}
// The superlayer of the layer.
//
// # Discussion
// 
// The superlayer manages the layout of its sublayers.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/superlayer
func (l CALayer) Superlayer() ICALayer {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("superlayer"))
	return CALayerFromID(objc.ID(rv))
}
// A Boolean indicating whether the layer contents must be updated when its
// bounds rectangle changes.
//
// # Discussion
// 
// When this property is set to [true], the layer automatically calls its
// [SetNeedsDisplay] method whenever its [Bounds] property changes. The
// default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplayOnBoundsChange
func (l CALayer) NeedsDisplayOnBoundsChange() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("needsDisplayOnBoundsChange"))
	return rv
}
func (l CALayer) SetNeedsDisplayOnBoundsChange(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setNeedsDisplayOnBoundsChange:"), value)
}
// The object responsible for laying out the layer’s sublayers.
//
// # Discussion
// 
// The object you assign to this property must nominally implement the
// CALayoutManager Informal Protocol informal protocol. If the layer’s
// delegate does not handle layout updates, the object assigned to this
// property is given a chance to update the layout of the layer’s sublayers.
// 
// In macOS, assign an instance of the [CAConstraintLayoutManager] class to
// this property if your layer uses layer-based constraints to handle layout
// changes.
// 
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l CALayer) LayoutManager() CALayoutManager {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("layoutManager"))
	return CALayoutManagerObjectFromID(rv)
}
func (l CALayer) SetLayoutManager(value CALayoutManager) {
	objc.Send[struct{}](l.ID, objc.Sel("setLayoutManager:"), value)
}
// A bitmask defining how the layer is resized when the bounds of its
// superlayer changes.
//
// # Discussion
// 
// If your app does not use a layout manager or constraints to handle layout
// changes, you can assign a value to this property to adjust the layer’s
// size in response to changes in the superlayer’s bounds. For a list of
// possible values, see [CAAutoresizingMask].
// 
// The default value of this property is [kCALayerNotSizable].
//
// [CAAutoresizingMask]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask
// [kCALayerNotSizable]: https://developer.apple.com/documentation/QuartzCore/CAAutoresizingMask/kCALayerNotSizable
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/autoresizingMask
func (l CALayer) AutoresizingMask() CAAutoresizingMask {
	rv := objc.Send[CAAutoresizingMask](l.ID, objc.Sel("autoresizingMask"))
	return CAAutoresizingMask(rv)
}
func (l CALayer) SetAutoresizingMask(value CAAutoresizingMask) {
	objc.Send[struct{}](l.ID, objc.Sel("setAutoresizingMask:"), value)
}
// The constraints used to position current layer’s sublayers.
//
// # Discussion
// 
// macOS apps can use this property to access their layer-based constraints.
// Before constraints can be applied, you must also assign a
// [CAConstraintLayoutManager] object to the [LayoutManager] property of the
// layer.
// 
// iOS apps do not support layer-based constraints.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/constraints
func (l CALayer) Constraints() []CAConstraint {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("constraints"))
	return objc.ConvertSlice(rv, func(id objc.ID) CAConstraint {
		return CAConstraintFromID(id)
	})
}
func (l CALayer) SetConstraints(value []CAConstraint) {
	objc.Send[struct{}](l.ID, objc.Sel("setConstraints:"), objectivec.IObjectSliceToNSArray(value))
}
// A dictionary containing layer actions.
//
// # Discussion
// 
// The default value of this property is `nil`. You can use this dictionary to
// store custom actions for your layer. The contents of this dictionary
// searched as part of the standard implementation of the [ActionForKey]
// method.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/actions
func (l CALayer) Actions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("actions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (l CALayer) SetActions(value foundation.INSDictionary) {
	objc.Send[struct{}](l.ID, objc.Sel("setActions:"), value)
}
// The visible region of the layer in its own coordinate space.
//
// # Discussion
// 
// The visible region is the area not clipped by the containing scroll layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/visibleRect
func (l CALayer) VisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("visibleRect"))
	return corefoundation.CGRect(rv)
}
// The name of the receiver.
//
// # Discussion
// 
// The layer name is used by some layout managers to identify a layer. The
// default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/name
func (l CALayer) Name() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (l CALayer) SetName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredDynamicRange
func (l CALayer) PreferredDynamicRange() CADynamicRange {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("preferredDynamicRange"))
	return CADynamicRange(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetPreferredDynamicRange(value CADynamicRange) {
	objc.Send[struct{}](l.ID, objc.Sel("setPreferredDynamicRange:"), objc.String(string(value)))
}
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsHeadroom
func (l CALayer) ContentsHeadroom() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("contentsHeadroom"))
	return rv
}
func (l CALayer) SetContentsHeadroom(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setContentsHeadroom:"), value)
}
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsExtendedDynamicRangeContent
func (l CALayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}
func (l CALayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurve
func (l CALayer) CornerCurve() CALayerCornerCurve {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("cornerCurve"))
	return CALayerCornerCurve(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetCornerCurve(value CALayerCornerCurve) {
	objc.Send[struct{}](l.ID, objc.Sel("setCornerCurve:"), objc.String(string(value)))
}
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/toneMapMode-swift.property
func (l CALayer) ToneMapMode() uint {
	rv := objc.Send[uint](l.ID, objc.Sel("toneMapMode"))
	return rv
}
func (l CALayer) SetToneMapMode(value uint) {
	objc.Send[struct{}](l.ID, objc.Sel("setToneMapMode:"), value)
}
// Determines if the receiver plays in the reverse upon completion.
//
// # Discussion
// 
// When [true], the receiver plays backwards after playing forwards. Defaults
// to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/autoreverses
func (l CALayer) Autoreverses() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("autoreverses"))
	return rv
}
func (l CALayer) SetAutoreverses(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setAutoreverses:"), value)
}
// Specifies the begin time of the receiver in relation to its parent object,
// if applicable.
//
// # Discussion
// 
// Defaults to 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (l CALayer) BeginTime() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("beginTime"))
	return rv
}
func (l CALayer) SetBeginTime(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setBeginTime:"), value)
}
// Specifies the basic duration of the animation, in seconds.
//
// # Discussion
// 
// Defaults to 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/duration
func (l CALayer) Duration() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("duration"))
	return rv
}
func (l CALayer) SetDuration(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setDuration:"), value)
}
// Determines if the receiver’s presentation is frozen or removed once its
// active duration has completed.
//
// # Discussion
// 
// The possible values are described in [Fill Modes]. The default is
// [removed].
//
// [Fill Modes]: https://developer.apple.com/documentation/QuartzCore/fill-modes
// [removed]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/removed
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/fillMode
func (l CALayer) FillMode() CAMediaTimingFillMode {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("fillMode"))
	return CAMediaTimingFillMode(foundation.NSStringFromID(rv).String())
}
func (l CALayer) SetFillMode(value CAMediaTimingFillMode) {
	objc.Send[struct{}](l.ID, objc.Sel("setFillMode:"), objc.String(string(value)))
}
// Determines the number of times the animation will repeat.
//
// # Discussion
// 
// May be fractional. If the `repeatCount` is 0, it is ignored. Defaults to 0.
// If both [RepeatDuration] and [RepeatCount] are specified the behavior is
// undefined.
// 
// Setting this property to [greatestFiniteMagnitude] will cause the animation
// to repeat forever.
//
// [greatestFiniteMagnitude]: https://developer.apple.com/documentation/Swift/Float/greatestFiniteMagnitude
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatCount
func (l CALayer) RepeatCount() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("repeatCount"))
	return rv
}
func (l CALayer) SetRepeatCount(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setRepeatCount:"), value)
}
// Determines how many seconds the animation will repeat for.
//
// # Discussion
// 
// Defaults to 0. If the `repeatDuration` is 0, it is ignored. If both
// [RepeatDuration] and [RepeatCount] are specified the behavior is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatDuration
func (l CALayer) RepeatDuration() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("repeatDuration"))
	return rv
}
func (l CALayer) SetRepeatDuration(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setRepeatDuration:"), value)
}
// Specifies how time is mapped to receiver’s time space from the parent
// time space.
//
// # Discussion
// 
// For example, if `speed` is 2.0 local time progresses twice as fast as
// parent time. Defaults to 1.0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/speed
func (l CALayer) Speed() float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("speed"))
	return rv
}
func (l CALayer) SetSpeed(value float32) {
	objc.Send[struct{}](l.ID, objc.Sel("setSpeed:"), value)
}
// Specifies an additional time offset in active local time.
//
// # Discussion
// 
// Defaults to 0. .
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/timeOffset
func (l CALayer) TimeOffset() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("timeOffset"))
	return rv
}
func (l CALayer) SetTimeOffset(value float64) {
	objc.Send[struct{}](l.ID, objc.Sel("setTimeOffset:"), value)
}

			// Protocol methods for CAMediaTiming
			

