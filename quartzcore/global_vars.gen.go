// Code generated from Apple documentation. DO NOT EDIT.

package quartzcore

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
)

var (
	// See: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRange/default
	CAFrameRateRangeDefault CAFrameRateRange
)

var (
	// CATransform3DIdentity is the identity transform: `[1 0 0 0; 0 1 0 0; 0 0 1 0; 0 0 0 1]`.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransform3DIdentity
	CATransform3DIdentity CATransform3D
)

var (
	// KCAAlignmentCenter is text is visually center aligned.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerAlignmentMode/center
	KCAAlignmentCenter CATextLayerAlignmentMode
	// KCAAlignmentJustified is text is justified.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerAlignmentMode/justified
	KCAAlignmentJustified CATextLayerAlignmentMode
	// KCAAlignmentLeft is text is visually left aligned.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerAlignmentMode/left
	KCAAlignmentLeft CATextLayerAlignmentMode
	// KCAAlignmentNatural is use the natural alignment of the text’s script.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerAlignmentMode/natural
	KCAAlignmentNatural CATextLayerAlignmentMode
	// KCAAlignmentRight is text is visually right aligned.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerAlignmentMode/right
	KCAAlignmentRight CATextLayerAlignmentMode
)

var (
	// KCAAnimationCubic is smooth spline calculation between keyframe values.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/cubic
	KCAAnimationCubic CAAnimationCalculationMode
	// KCAAnimationCubicPaced is cubic keyframe values are interpolated to produce an even pace throughout the animation.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/cubicPaced
	KCAAnimationCubicPaced CAAnimationCalculationMode
	// KCAAnimationDiscrete is each keyframe value is used in turn, no interpolated values are calculated.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/discrete
	KCAAnimationDiscrete CAAnimationCalculationMode
	// KCAAnimationLinear is simple linear calculation between keyframe values.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/linear
	KCAAnimationLinear CAAnimationCalculationMode
	// KCAAnimationPaced is linear keyframe values are interpolated to produce an even pace throughout the animation.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/paced
	KCAAnimationPaced CAAnimationCalculationMode
)

var (
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationRotationMode/rotateAuto
	KCAAnimationRotateAuto CAAnimationRotationMode
	// See: https://developer.apple.com/documentation/QuartzCore/CAAnimationRotationMode/rotateAutoReverse
	KCAAnimationRotateAutoReverse CAAnimationRotationMode
)

var (
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFormat/automatic
	KCAContentsFormatAutomatic CALayerContentsFormat
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFormat/gray8Uint
	KCAContentsFormatGray8Uint CALayerContentsFormat
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFormat/RGBA16Float
	KCAContentsFormatRGBA16Float CALayerContentsFormat
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFormat/RGBA8Uint
	KCAContentsFormatRGBA8Uint CALayerContentsFormat
)

var (
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerCornerCurve/circular
	KCACornerCurveCircular CALayerCornerCurve
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerCornerCurve/continuous
	KCACornerCurveContinuous CALayerCornerCurve
)

var (
	// KCAEmitterLayerAdditive is the particles are rendered using source-additive compositing.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerRenderMode/additive
	KCAEmitterLayerAdditive CAEmitterLayerRenderMode
	// KCAEmitterLayerBackToFront is particles are rendered from back to front, sorted by z-position. This mode uses source-over compositing.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerRenderMode/backToFront
	KCAEmitterLayerBackToFront CAEmitterLayerRenderMode
	// KCAEmitterLayerOldestFirst is particles are rendered oldest first. This mode uses source-over compositing.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerRenderMode/oldestFirst
	KCAEmitterLayerOldestFirst CAEmitterLayerRenderMode
	// KCAEmitterLayerOldestLast is particles are rendered oldest last. This mode uses source-over compositing.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerRenderMode/oldestLast
	KCAEmitterLayerOldestLast CAEmitterLayerRenderMode
	// KCAEmitterLayerUnordered is particles are rendered unordered. This mode uses source-over compositing.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerRenderMode/unordered
	KCAEmitterLayerUnordered CAEmitterLayerRenderMode
)

var (
	// KCAEmitterLayerCircle is particles are emitted from a circle centered at (`emitterPosition.X()`, `emitterPosition.Y()`, `emitterZPosition`) of radius `emitterSize.Width()`.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/circle
	KCAEmitterLayerCircle CAEmitterLayerEmitterShape
	// KCAEmitterLayerCuboid is particles are emitted from a cuboid (3D rectangle) with opposite corners: [emitterPosition.x - emitterSize.width/2, emitterPosition.y - emitterSize.height/2, emitterZPosition - emitterDepth/2], [emitterPosition.x + emitterSize.width/2, emitterPosition.y + emitterSize.height/2, emitterZPosition+emitterDepth/2].
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/cuboid
	KCAEmitterLayerCuboid CAEmitterLayerEmitterShape
	// KCAEmitterLayerLine is particles are emitted along a line from (`emitterPosition.X() - emitterSize.Width()/2`, `emitterPosition.Y()`, `emitterZPosition`) to (`emitterPosition.X() + emitterSize.Width()/2`, `emitterPosition.Y()`, `emitterZPosition`).
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/line
	KCAEmitterLayerLine CAEmitterLayerEmitterShape
	// KCAEmitterLayerPoint is particles are emitted from a single point at (`emitterPosition.X()`, `emitterPosition.Y()`, `emitterZPosition`).
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/point
	KCAEmitterLayerPoint CAEmitterLayerEmitterShape
	// KCAEmitterLayerRectangle is particles are emitted from a rectangle with opposite corners [emitterPosition.x - emitterSize.width/2, emitterPosition.y - emitterSize.height/2, emitterZPosition], [emitterPosition.x + emitterSize.width/2, emitterPosition.y + emitterSize.height/2, emitterZPosition].
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/rectangle
	KCAEmitterLayerRectangle CAEmitterLayerEmitterShape
	// KCAEmitterLayerSphere is particles are emitted from a sphere centered at (`emitterPosition.X()`, `emitterPosition.Y()`, `emitterZPosition`) of radius `emitterSize.Width()`.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterShape/sphere
	KCAEmitterLayerSphere CAEmitterLayerEmitterShape
)

var (
	// KCAEmitterLayerOutline is particles are emitted from the outline of the particle emitter.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterMode/outline
	KCAEmitterLayerOutline CAEmitterLayerEmitterMode
	// KCAEmitterLayerPoints is particles are emitted from points on the particle emitter.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterMode/points
	KCAEmitterLayerPoints CAEmitterLayerEmitterMode
	// KCAEmitterLayerSurface is particles are emitted from the surface of the particle emitter.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterMode/surface
	KCAEmitterLayerSurface CAEmitterLayerEmitterMode
	// KCAEmitterLayerVolume is particles are emitted from the a position within the particle emitter.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAEmitterLayerEmitterMode/volume
	KCAEmitterLayerVolume CAEmitterLayerEmitterMode
)

var (
	// KCAFillModeBackwards is the receiver clamps values before zero to zero when the animation is completed.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/backwards
	KCAFillModeBackwards CAMediaTimingFillMode
	// KCAFillModeBoth is the receiver clamps values at both ends of the object’s time space.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/both
	KCAFillModeBoth CAMediaTimingFillMode
	// KCAFillModeForwards is the receiver remains visible in its final state when the animation is completed.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/forwards
	KCAFillModeForwards CAMediaTimingFillMode
	// KCAFillModeRemoved is the receiver is removed from the presentation when the animation is completed.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/removed
	KCAFillModeRemoved CAMediaTimingFillMode
)

var (
	// KCAFillRuleEvenOdd is specifies the even-odd winding rule. Count the total number of path crossings. If the number of crossings is even, the point is outside the path. If the number of crossings is odd, the point is inside the path and the region containing it should be filled.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerFillRule/evenOdd
	KCAFillRuleEvenOdd CAShapeLayerFillRule
	// KCAFillRuleNonZero is specifies the non-zero winding rule. Count each left-to-right path as +1 and each right-to-left path as -1. If the sum of all crossings is 0, the point is outside the path. If the sum is nonzero, the point is inside the path and the region containing it is filled.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerFillRule/nonZero
	KCAFillRuleNonZero CAShapeLayerFillRule
)

var (
	// KCAFilterLinear is linear interpolation filter.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/linear
	KCAFilterLinear CALayerContentsFilter
	// KCAFilterNearest is nearest neighbor interpolation filter.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/nearest
	KCAFilterNearest CALayerContentsFilter
	// KCAFilterTrilinear is trilinear minification filter. Enables mipmap generation. Some renderers may ignore this, or impose additional restrictions, such as source images requiring power-of-two dimensions.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsFilter/trilinear
	KCAFilterTrilinear CALayerContentsFilter
)

var (
	// KCAGradientLayerAxial is an axial gradient (also called a linear gradient) varies along an axis between two defined end points. All points that lie on a line perpendicular to the axis have the same color value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayerType/axial
	KCAGradientLayerAxial CAGradientLayerType
	// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayerType/conic
	KCAGradientLayerConic CAGradientLayerType
	// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayerType/radial
	KCAGradientLayerRadial CAGradientLayerType
)

var (
	// KCAGravityBottom is the content is horizontally centered at the bottom-edge of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/bottom
	KCAGravityBottom CALayerContentsGravity
	// KCAGravityBottomLeft is the content is positioned in the bottom-left corner of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/bottomLeft
	KCAGravityBottomLeft CALayerContentsGravity
	// KCAGravityBottomRight is the content is positioned in the bottom-right corner of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/bottomRight
	KCAGravityBottomRight CALayerContentsGravity
	// KCAGravityCenter is the content is horizontally and vertically centered in the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/center
	KCAGravityCenter CALayerContentsGravity
	// KCAGravityLeft is the content is vertically centered at the left-edge of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/left
	KCAGravityLeft CALayerContentsGravity
	// KCAGravityResize is the content is resized to fit the entire bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resize
	KCAGravityResize CALayerContentsGravity
	// KCAGravityResizeAspect is the content is resized to fit the bounds rectangle, preserving the aspect of the content. If the content does not completely fill the bounds rectangle, the content is centered in the partial axis.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resizeAspect
	KCAGravityResizeAspect CALayerContentsGravity
	// KCAGravityResizeAspectFill is the content is resized to completely fill the bounds rectangle, while still preserving the aspect of the content. The content is centered in the axis it exceeds.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resizeAspectFill
	KCAGravityResizeAspectFill CALayerContentsGravity
	// KCAGravityRight is the content is vertically centered at the right-edge of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/right
	KCAGravityRight CALayerContentsGravity
	// KCAGravityTop is the content is horizontally centered at the top-edge of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/top
	KCAGravityTop CALayerContentsGravity
	// KCAGravityTopLeft is the content is positioned in the top-left corner of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/topLeft
	KCAGravityTopLeft CALayerContentsGravity
	// KCAGravityTopRight is the content is positioned in the top-right corner of the bounds rectangle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/topRight
	KCAGravityTopRight CALayerContentsGravity
)

var (
	// KCALineCapButt is specifies a butt line cap style for endpoints for an open path when stroked.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineCap/butt
	KCALineCapButt CAShapeLayerLineCap
	// KCALineCapRound is specifies a round line cap style for endpoints for an open path when stroked.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineCap/round
	KCALineCapRound CAShapeLayerLineCap
	// KCALineCapSquare is specifies a square line cap style for endpoints for an open path when stroked.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineCap/square
	KCALineCapSquare CAShapeLayerLineCap
)

var (
	// KCALineJoinBevel is specifies a bevel line shape of the joints between connected segments of a stroked path.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineJoin/bevel
	KCALineJoinBevel CAShapeLayerLineJoin
	// KCALineJoinMiter is specifies a miter line shape of the joints between connected segments of a stroked path.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineJoin/miter
	KCALineJoinMiter CAShapeLayerLineJoin
	// KCALineJoinRound is specifies a round line shape of the joints between connected segments of a stroked path.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineJoin/round
	KCALineJoinRound CAShapeLayerLineJoin
)

var (
	// KCAMediaTimingFunctionDefault is the system default timing function. Use this function to ensure that the timing of your animations matches that of most system animations.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunctionName/default
	KCAMediaTimingFunctionDefault CAMediaTimingFunctionName
	// KCAMediaTimingFunctionEaseIn is ease-in pacing, which causes an animation to begin slowly and then speed up as it progresses.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunctionName/easeIn
	KCAMediaTimingFunctionEaseIn CAMediaTimingFunctionName
	// KCAMediaTimingFunctionEaseInEaseOut is ease-in-ease-out pacing, which causes an animation to begin slowly, accelerate through the middle of its duration, and then slow again before completing.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunctionName/easeInEaseOut
	KCAMediaTimingFunctionEaseInEaseOut CAMediaTimingFunctionName
	// KCAMediaTimingFunctionEaseOut is ease-out pacing, which causes an animation to begin quickly and then slow as it progresses.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunctionName/easeOut
	KCAMediaTimingFunctionEaseOut CAMediaTimingFunctionName
	// KCAMediaTimingFunctionLinear is linear pacing, which causes an animation to occur evenly over its duration.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunctionName/linear
	KCAMediaTimingFunctionLinear CAMediaTimingFunctionName
)

var (
	// KCAOnOrderIn is the identifier that represents the action taken when a layer becomes visible, either as a result being inserted into the visible layer hierarchy or the layer is no longer set as hidden.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/kCAOnOrderIn
	KCAOnOrderIn string
	// KCAOnOrderOut is the identifier that represents the action taken when the layer is removed from the layer hierarchy or is hidden.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/kCAOnOrderOut
	KCAOnOrderOut string
	// See: https://developer.apple.com/documentation/QuartzCore/kCARendererColorSpace
	KCARendererColorSpace string
	// See: https://developer.apple.com/documentation/QuartzCore/kCARendererMetalCommandQueue
	KCARendererMetalCommandQueue string
	// KCATransactionAnimationDuration is duration, in seconds, for animations triggered within the transaction group.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationDuration
	KCATransactionAnimationDuration string
	// See: https://developer.apple.com/documentation/QuartzCore/kCATransactionAnimationTimingFunction
	KCATransactionAnimationTimingFunction string
	// See: https://developer.apple.com/documentation/QuartzCore/kCATransactionCompletionBlock
	KCATransactionCompletionBlock string
	// KCATransactionDisableActions is a key whose value indicates whether implicit actions for property changes made within the transaction group are suppressed.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/kCATransactionDisableActions
	KCATransactionDisableActions string
	// KCATransition is the identifier that represents a transition animation.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/kCATransition
	KCATransition string
)

var (
	// KCAScrollBoth is the receiver is able to scroll both horizontally and vertically.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayerScrollMode/both
	KCAScrollBoth CAScrollLayerScrollMode
	// KCAScrollHorizontally is the receiver is able to scroll horizontally.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayerScrollMode/horizontally
	KCAScrollHorizontally CAScrollLayerScrollMode
	// KCAScrollNone is the receiver is unable to scroll.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayerScrollMode/none
	KCAScrollNone CAScrollLayerScrollMode
	// KCAScrollVertically is the receiver is able to scroll vertically.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayerScrollMode/vertically
	KCAScrollVertically CAScrollLayerScrollMode
)

var (
	// KCATransitionFade is the layer’s content fades as it becomes visible or hidden.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionType/fade
	KCATransitionFade CATransitionType
	// KCATransitionMoveIn is the layer’s content slides into place over any existing content.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionType/moveIn
	KCATransitionMoveIn CATransitionType
	// KCATransitionPush is the layer’s content pushes any existing content as it slides into place.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionType/push
	KCATransitionPush CATransitionType
	// KCATransitionReveal is the layer’s content is revealed gradually in the direction specified by the transition subtype.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionType/reveal
	KCATransitionReveal CATransitionType
)

var (
	// KCATransitionFromBottom is the transition begins at the bottom of the layer.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionSubtype/fromBottom
	KCATransitionFromBottom CATransitionSubtype
	// KCATransitionFromLeft is the transition begins at the left side of the layer.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionSubtype/fromLeft
	KCATransitionFromLeft CATransitionSubtype
	// KCATransitionFromRight is the transition begins at the right side of the layer.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionSubtype/fromRight
	KCATransitionFromRight CATransitionSubtype
	// KCATransitionFromTop is the transition begins at the top of the layer.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATransitionSubtype/fromTop
	KCATransitionFromTop CATransitionSubtype
)

var (
	// KCATruncationEnd is each line is displayed so that the beginning fits in the container and the missing text is indicated by some kind of ellipsis glyph.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerTruncationMode/end
	KCATruncationEnd CATextLayerTruncationMode
	// KCATruncationMiddle is each line is displayed so that the beginning and end fit in the container and the missing text is indicated by some kind of ellipsis glyph in the middle.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerTruncationMode/middle
	KCATruncationMiddle CATextLayerTruncationMode
	// KCATruncationNone is each line is displayed so that the text is either wrapped or clipped.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerTruncationMode/none
	KCATruncationNone CATextLayerTruncationMode
	// KCATruncationStart is each line is displayed so that the end fits in the container and the missing text is indicated by some kind of ellipsis glyph.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CATextLayerTruncationMode/start
	KCATruncationStart CATextLayerTruncationMode
)

var (
	// KCAValueFunctionRotateX is a value function that rotates by the input value, in radians, around the x-axis. This value function expects a single input value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/rotateX
	KCAValueFunctionRotateX CAValueFunctionName
	// KCAValueFunctionRotateY is a value function that rotates by the input value, in radians, around the y-axis. This value function expects a single input value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/rotateY
	KCAValueFunctionRotateY CAValueFunctionName
	// KCAValueFunctionRotateZ is a value function that rotates by the input value, in radians, around the z-axis. This value function expects a single input value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/rotateZ
	KCAValueFunctionRotateZ CAValueFunctionName
	// KCAValueFunctionScale is a value function scales by the input value along all three axis. Animations using this value transform function must provide animation values in an [NSArray] of three [NSNumber] instances that specify the (x, y, z) scale values.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/scale
	KCAValueFunctionScale CAValueFunctionName
	// KCAValueFunctionScaleX is a value function scales by the input value along the x-axis. Animations referencing this value transform function must provide a single animation value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/scaleX
	KCAValueFunctionScaleX CAValueFunctionName
	// KCAValueFunctionScaleY is a value function scales by the input value along the y-axis. Animations referencing this value function must provide a single animation value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/scaleY
	KCAValueFunctionScaleY CAValueFunctionName
	// KCAValueFunctionScaleZ is a value function that scales by the input value along the z-axis. Animations referencing this value function must provide a single animation value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/scaleZ
	KCAValueFunctionScaleZ CAValueFunctionName
	// KCAValueFunctionTranslate is a value function that translates by the input values along all three axis. Animations using this value transform function must provide animation values in an [NSArray] of three [NSNumber] instances that specify the (x, y, z) translate values.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/translate
	KCAValueFunctionTranslate CAValueFunctionName
	// KCAValueFunctionTranslateX is a value function translates by the input value along the x-axis. Animations referencing this value function must provide a single input value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/translateX
	KCAValueFunctionTranslateX CAValueFunctionName
	// KCAValueFunctionTranslateY is a value function translates by the input value along the y-axis. Animations referencing this value function must provide a single input value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/translateY
	KCAValueFunctionTranslateY CAValueFunctionName
	// KCAValueFunctionTranslateZ is a value function translates by the input value along the z-axis. Animations referencing this value function must provide a single input value.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/translateZ
	KCAValueFunctionTranslateZ CAValueFunctionName
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CADynamicRangeAutomatic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CADynamicRanges.Automatic = CADynamicRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CADynamicRangeConstrainedHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CADynamicRanges.ConstrainedHigh = CADynamicRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CADynamicRangeHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CADynamicRanges.High = CADynamicRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CADynamicRangeStandard"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				CADynamicRanges.Standard = CADynamicRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CAFrameRateRangeDefault"); err == nil && ptr != 0 {
		CAFrameRateRangeDefault = *(*CAFrameRateRange)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CATransform3DIdentity"); err == nil && ptr != 0 {
		CATransform3DIdentity = *(*CATransform3D)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAlignmentCenter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAlignmentCenter = CATextLayerAlignmentMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAlignmentJustified"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAlignmentJustified = CATextLayerAlignmentMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAlignmentLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAlignmentLeft = CATextLayerAlignmentMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAlignmentNatural"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAlignmentNatural = CATextLayerAlignmentMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAlignmentRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAlignmentRight = CATextLayerAlignmentMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationCubic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationCubic = CAAnimationCalculationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationCubicPaced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationCubicPaced = CAAnimationCalculationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationDiscrete"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationDiscrete = CAAnimationCalculationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationLinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationLinear = CAAnimationCalculationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationPaced"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationPaced = CAAnimationCalculationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationRotateAuto"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationRotateAuto = CAAnimationRotationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAAnimationRotateAutoReverse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAAnimationRotateAutoReverse = CAAnimationRotationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAContentsFormatAutomatic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAContentsFormatAutomatic = CALayerContentsFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAContentsFormatGray8Uint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAContentsFormatGray8Uint = CALayerContentsFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAContentsFormatRGBA16Float"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAContentsFormatRGBA16Float = CALayerContentsFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAContentsFormatRGBA8Uint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAContentsFormatRGBA8Uint = CALayerContentsFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCACornerCurveCircular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCACornerCurveCircular = CALayerCornerCurve(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCACornerCurveContinuous"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCACornerCurveContinuous = CALayerCornerCurve(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerAdditive"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerAdditive = CAEmitterLayerRenderMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerBackToFront"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerBackToFront = CAEmitterLayerRenderMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerCircle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerCircle = CAEmitterLayerEmitterShape(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerCuboid"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerCuboid = CAEmitterLayerEmitterShape(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerLine"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerLine = CAEmitterLayerEmitterShape(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerOldestFirst"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerOldestFirst = CAEmitterLayerRenderMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerOldestLast"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerOldestLast = CAEmitterLayerRenderMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerOutline"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerOutline = CAEmitterLayerEmitterMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerPoint"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerPoint = CAEmitterLayerEmitterShape(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerPoints"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerPoints = CAEmitterLayerEmitterMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerRectangle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerRectangle = CAEmitterLayerEmitterShape(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerSphere"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerSphere = CAEmitterLayerEmitterShape(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerSurface"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerSurface = CAEmitterLayerEmitterMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerUnordered"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerUnordered = CAEmitterLayerRenderMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAEmitterLayerVolume"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAEmitterLayerVolume = CAEmitterLayerEmitterMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFillModeBackwards"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFillModeBackwards = CAMediaTimingFillMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFillModeBoth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFillModeBoth = CAMediaTimingFillMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFillModeForwards"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFillModeForwards = CAMediaTimingFillMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFillModeRemoved"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFillModeRemoved = CAMediaTimingFillMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFillRuleEvenOdd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFillRuleEvenOdd = CAShapeLayerFillRule(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFillRuleNonZero"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFillRuleNonZero = CAShapeLayerFillRule(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFilterLinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFilterLinear = CALayerContentsFilter(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFilterNearest"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFilterNearest = CALayerContentsFilter(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAFilterTrilinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAFilterTrilinear = CALayerContentsFilter(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGradientLayerAxial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGradientLayerAxial = CAGradientLayerType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGradientLayerConic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGradientLayerConic = CAGradientLayerType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGradientLayerRadial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGradientLayerRadial = CAGradientLayerType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityBottom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityBottom = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityBottomLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityBottomLeft = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityBottomRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityBottomRight = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityCenter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityCenter = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityLeft = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityResize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityResize = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityResizeAspect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityResizeAspect = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityResizeAspectFill"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityResizeAspectFill = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityRight = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityTop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityTop = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityTopLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityTopLeft = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAGravityTopRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAGravityTopRight = CALayerContentsGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCALineCapButt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCALineCapButt = CAShapeLayerLineCap(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCALineCapRound"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCALineCapRound = CAShapeLayerLineCap(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCALineCapSquare"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCALineCapSquare = CAShapeLayerLineCap(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCALineJoinBevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCALineJoinBevel = CAShapeLayerLineJoin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCALineJoinMiter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCALineJoinMiter = CAShapeLayerLineJoin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCALineJoinRound"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCALineJoinRound = CAShapeLayerLineJoin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAMediaTimingFunctionDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAMediaTimingFunctionDefault = CAMediaTimingFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAMediaTimingFunctionEaseIn"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAMediaTimingFunctionEaseIn = CAMediaTimingFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAMediaTimingFunctionEaseInEaseOut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAMediaTimingFunctionEaseInEaseOut = CAMediaTimingFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAMediaTimingFunctionEaseOut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAMediaTimingFunctionEaseOut = CAMediaTimingFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAMediaTimingFunctionLinear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAMediaTimingFunctionLinear = CAMediaTimingFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAOnOrderIn"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAOnOrderIn = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAOnOrderOut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAOnOrderOut = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCARendererColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCARendererColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCARendererMetalCommandQueue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCARendererMetalCommandQueue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAScrollBoth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAScrollBoth = CAScrollLayerScrollMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAScrollHorizontally"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAScrollHorizontally = CAScrollLayerScrollMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAScrollNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAScrollNone = CAScrollLayerScrollMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAScrollVertically"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAScrollVertically = CAScrollLayerScrollMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransactionAnimationDuration"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransactionAnimationDuration = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransactionAnimationTimingFunction"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransactionAnimationTimingFunction = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransactionCompletionBlock"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransactionCompletionBlock = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransactionDisableActions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransactionDisableActions = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionFade"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionFade = CATransitionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionFromBottom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionFromBottom = CATransitionSubtype(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionFromLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionFromLeft = CATransitionSubtype(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionFromRight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionFromRight = CATransitionSubtype(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionFromTop"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionFromTop = CATransitionSubtype(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionMoveIn"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionMoveIn = CATransitionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionPush"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionPush = CATransitionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATransitionReveal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATransitionReveal = CATransitionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATruncationEnd"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATruncationEnd = CATextLayerTruncationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATruncationMiddle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATruncationMiddle = CATextLayerTruncationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATruncationNone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATruncationNone = CATextLayerTruncationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCATruncationStart"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCATruncationStart = CATextLayerTruncationMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionRotateX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionRotateX = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionRotateY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionRotateY = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionRotateZ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionRotateZ = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionScale"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionScale = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionScaleX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionScaleX = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionScaleY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionScaleY = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionScaleZ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionScaleZ = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionTranslate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionTranslate = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionTranslateX"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionTranslateX = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionTranslateY"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionTranslateY = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCAValueFunctionTranslateZ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCAValueFunctionTranslateZ = CAValueFunctionName(objc.GoString(cstr))
			}
		}
	}

}

// CADynamicRanges provides typed accessors for [CADynamicRange] constants.
var CADynamicRanges struct {
	Automatic CADynamicRange
	ConstrainedHigh CADynamicRange
	High CADynamicRange
	Standard CADynamicRange
}

