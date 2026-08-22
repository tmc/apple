// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKEffectNode] class.
var (
	_SKEffectNodeClass     SKEffectNodeClass
	_SKEffectNodeClassOnce sync.Once
)

func getSKEffectNodeClass() SKEffectNodeClass {
	_SKEffectNodeClassOnce.Do(func() {
		_SKEffectNodeClass = SKEffectNodeClass{class: objc.GetClass("SKEffectNode")}
	})
	return _SKEffectNodeClass
}

// GetSKEffectNodeClass returns the class object for SKEffectNode.
func GetSKEffectNodeClass() SKEffectNodeClass {
	return getSKEffectNodeClass()
}

type SKEffectNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKEffectNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKEffectNodeClass) Alloc() SKEffectNode {
	rv := objc.Send[SKEffectNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A node that renders its children into a separate buffer, optionally
// applying an effect, before drawing the final result.
//
// # Overview
//
// An [SKEffectNode] object renders its children into a buffer and optionally
// applies a Core Image filter to this rendered output. Because effect nodes
// conform to [SKWarpable], you can also use them to apply distortions to
// nodes that don’t implement the protocol, such as shape and video nodes.
// Use effect nodes to incorporate sophisticated special effects into a scene
// or to cache the contents of a static subtree for faster rendering
// performance.
//
// Each time a new frame is rendered using the effect node, the effect node
// follows these steps:
//
// - It draws its children into a private framebuffer. - It applies a Core
// Image effect to the private framebuffer. This stage is optional; see the
// [SKEffectNode.Filter] and [SKEffectNode.ShouldEnableEffects] properties. -
// It blends the contents of its private framebuffer into its parent’s
// framebuffer, using one of the standard sprite blend modes. - It discards
// its private framebuffer. This step is optional; see the
// [SKEffectNode.ShouldRasterize] property.
//
// # Applying Core Image Filters with an Effect Node
//
//   - [SKEffectNode.Filter]: The Core Image filter to apply.
//   - [SKEffectNode.SetFilter]
//   - [SKEffectNode.ShouldEnableEffects]: A Boolean value that determines whether the effect node applies the filter to its children as they are drawn.
//   - [SKEffectNode.SetShouldEnableEffects]
//   - [SKEffectNode.ShouldCenterFilter]: A Boolean value that determines whether the effect node automatically sets the filter’s image center.
//   - [SKEffectNode.SetShouldCenterFilter]
//
// # Applying a Shader with an Effect Node
//
//   - [SKEffectNode.Shader]: A custom shader that is called when the effect node is blended into the parent’s framebuffer.
//   - [SKEffectNode.SetShader]
//   - [SKEffectNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [SKEffectNode.SetAttributeValues]
//   - [SKEffectNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [SKEffectNode.ValueForAttributeNamed]: Gets the value of a shader attribute.
//
// # Flattening an Effect Node’s Child Tree for Performance Improvement
//
//   - [SKEffectNode.ShouldRasterize]: A Boolean value that indicates whether the results of rendering the child nodes should be cached.
//   - [SKEffectNode.SetShouldRasterize]
//
// # Configuring Alpha Blending
//
//   - [SKEffectNode.BlendMode]: The blend mode used to draw the node’s contents into its parent’s framebuffer.
//   - [SKEffectNode.SetBlendMode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode
type SKEffectNode struct {
	SKNode
}

// SKEffectNodeFromID constructs a [SKEffectNode] from an objc.ID.
//
// A node that renders its children into a separate buffer, optionally
// applying an effect, before drawing the final result.
func SKEffectNodeFromID(id objc.ID) SKEffectNode {
	return SKEffectNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKEffectNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKEffectNode] class.
//
// # Applying Core Image Filters with an Effect Node
//
//   - [ISKEffectNode.Filter]: The Core Image filter to apply.
//   - [ISKEffectNode.SetFilter]
//   - [ISKEffectNode.ShouldEnableEffects]: A Boolean value that determines whether the effect node applies the filter to its children as they are drawn.
//   - [ISKEffectNode.SetShouldEnableEffects]
//   - [ISKEffectNode.ShouldCenterFilter]: A Boolean value that determines whether the effect node automatically sets the filter’s image center.
//   - [ISKEffectNode.SetShouldCenterFilter]
//
// # Applying a Shader with an Effect Node
//
//   - [ISKEffectNode.Shader]: A custom shader that is called when the effect node is blended into the parent’s framebuffer.
//   - [ISKEffectNode.SetShader]
//   - [ISKEffectNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [ISKEffectNode.SetAttributeValues]
//   - [ISKEffectNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [ISKEffectNode.ValueForAttributeNamed]: Gets the value of a shader attribute.
//
// # Flattening an Effect Node’s Child Tree for Performance Improvement
//
//   - [ISKEffectNode.ShouldRasterize]: A Boolean value that indicates whether the results of rendering the child nodes should be cached.
//   - [ISKEffectNode.SetShouldRasterize]
//
// # Configuring Alpha Blending
//
//   - [ISKEffectNode.BlendMode]: The blend mode used to draw the node’s contents into its parent’s framebuffer.
//   - [ISKEffectNode.SetBlendMode]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode
type ISKEffectNode interface {
	ISKNode

	// Topic: Applying Core Image Filters with an Effect Node

	// The Core Image filter to apply.
	Filter() coreimage.CIFilter
	SetFilter(value coreimage.CIFilter)
	// A Boolean value that determines whether the effect node applies the filter to its children as they are drawn.
	ShouldEnableEffects() bool
	SetShouldEnableEffects(value bool)
	// A Boolean value that determines whether the effect node automatically sets the filter’s image center.
	ShouldCenterFilter() bool
	SetShouldCenterFilter(value bool)

	// Topic: Applying a Shader with an Effect Node

	// A custom shader that is called when the effect node is blended into the parent’s framebuffer.
	Shader() ISKShader
	SetShader(value ISKShader)
	// The values of each attribute associated with the node’s attached shader.
	AttributeValues() foundation.INSDictionary
	SetAttributeValues(value foundation.INSDictionary)
	// Sets an attribute value for an attached shader.
	SetValueForAttributeNamed(value ISKAttributeValue, key string)
	// Gets the value of a shader attribute.
	ValueForAttributeNamed(key string) ISKAttributeValue

	// Topic: Flattening an Effect Node’s Child Tree for Performance Improvement

	// A Boolean value that indicates whether the results of rendering the child nodes should be cached.
	ShouldRasterize() bool
	SetShouldRasterize(value bool)

	// Topic: Configuring Alpha Blending

	// The blend mode used to draw the node’s contents into its parent’s framebuffer.
	BlendMode() SKBlendMode
	SetBlendMode(value SKBlendMode)

	// The maximum number of subdivision iterations used to generate the final vertices.
	SubdivisionLevels() int
	// The warp geometry used to define the distortion.
	WarpGeometry() ISKWarpGeometry
}

// Init initializes the instance.
func (e SKEffectNode) Init() SKEffectNode {
	rv := objc.Send[SKEffectNode](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e SKEffectNode) Autorelease() SKEffectNode {
	rv := objc.Send[SKEffectNode](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKEffectNode creates a new SKEffectNode instance.
func NewSKEffectNode() SKEffectNode {
	class := getSKEffectNodeClass()
	rv := objc.Send[SKEffectNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewEffectNodeWithCoder(aDecoder foundation.INSCoder) SKEffectNode {
	instance := getSKEffectNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKEffectNodeFromID(rv)
}

// Creates a new node by loading an archive file from the game’s main
// bundle.
//
// filename: The name of the file, without a file extension. The file must be in the
// app’s main bundle and have a `XCUIElementTypeSks` filename extension.
//
// # Return Value
//
// The unarchived node object.
//
// # Discussion
//
// If you call this method on a subclass of the [SKScene] class and the object
// in the archive is an [SKScene] object, the returned object is initialized
// as if it is a member of the subclass. You use this behavior to create scene
// layouts in the Xcode Editor and provide custom behaviors in your subclass.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:)
func NewEffectNodeWithFileNamed(filename string) SKEffectNode {
	rv := objc.Send[objc.ID](objc.ID(getSKEffectNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKEffectNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewEffectNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKEffectNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKEffectNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKEffectNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKEffectNode{}, objc.ErrInitFailed
	}
	return SKEffectNodeFromID(rv), nil
}

// Sets an attribute value for an attached shader.
//
// value: An attribute value object containing the scalar or vector value to set in
// the attached shader.
//
// key: The attribute name.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/setValue(_:forAttribute:)
func (e SKEffectNode) SetValueForAttributeNamed(value ISKAttributeValue, key string) {
	objc.Send[objc.ID](e.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}

// Gets the value of a shader attribute.
//
// key: The attribute name.
//
// # Return Value
//
// An attribute value object containing the scalar or vector value or `nil` if
// no such attribute exists.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/value(forAttributeNamed:)
func (e SKEffectNode) ValueForAttributeNamed(key string) ISKAttributeValue {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("valueForAttributeNamed:"), objc.String(key))
	return SKAttributeValueFromID(rv)
}

// The maximum number of subdivision iterations used to generate the final
// vertices.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/subdivisionLevels
func (e SKEffectNode) SubdivisionLevels() int {
	rv := objc.Send[int](e.ID, objc.Sel("subdivisionLevels"))
	return rv
}

// The warp geometry used to define the distortion.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/warpGeometry
func (e SKEffectNode) WarpGeometry() ISKWarpGeometry {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("warpGeometry"))
	return SKWarpGeometryFromID(rv)
}

// The Core Image filter to apply.
//
// # Discussion
//
// The Core Image filter must have a single `inputImage` parameter and produce
// a single `outputImage` parameter. The default value is `nil`. If the value
// is `nil` and the effect node is enabled, no filtering takes place. However,
// its children are still rendered in a separate pass and blended to the
// parent’s framebuffer.
//
// If you wish to use a Core Image filter that doesn’t have an `inputImage`
// parameter, such as a sunbeams generator, you can subclass [CIFilter] and
// add an `inputImage` property. The input image’s extent can be used to
// define properties such as radius on the filter. The following code creates
// a filter based on [CISunbeamsGenerator] which can be used as an effect
// node’s filter:
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/filter
//
// [CIFilter]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
// [CISunbeamsGenerator]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CISunbeamsGenerator
func (e SKEffectNode) Filter() coreimage.CIFilter {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("filter"))
	return coreimage.CIFilterFromID(objc.ID(rv))
}
func (e SKEffectNode) SetFilter(value coreimage.CIFilter) {
	objc.Send[struct{}](e.ID, objc.Sel("setFilter:"), value)
}

// A Boolean value that determines whether the effect node applies the filter
// to its children as they are drawn.
//
// # Discussion
//
// If the value of this property is true, the effect node applies the filter
// and blends the results. If the value is false, the effect node is ignored
// and its children are rendered normally. The default value is false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/shouldEnableEffects
func (e SKEffectNode) ShouldEnableEffects() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("shouldEnableEffects"))
	return rv
}
func (e SKEffectNode) SetShouldEnableEffects(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setShouldEnableEffects:"), value)
}

// A Boolean value that determines whether the effect node automatically sets
// the filter’s image center.
//
// # Discussion
//
// If the value of this property is true and the filter has an `inputCenter`
// parameter, the effect node automatically sets the filter’s input center
// to the effect node’s origin. The default value is true.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/shouldCenterFilter
func (e SKEffectNode) ShouldCenterFilter() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("shouldCenterFilter"))
	return rv
}
func (e SKEffectNode) SetShouldCenterFilter(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setShouldCenterFilter:"), value)
}

// A custom shader that is called when the effect node is blended into the
// parent’s framebuffer.
//
// # Discussion
//
// The default value is `nil`, meaning that default blending behavior
// executes. If a shader is specified, it is called when the rasterized image
// is blended into the parent’s framebuffer.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/shader
func (e SKEffectNode) Shader() ISKShader {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shader"))
	return SKShaderFromID(objc.ID(rv))
}
func (e SKEffectNode) SetShader(value ISKShader) {
	objc.Send[struct{}](e.ID, objc.Sel("setShader:"), value)
}

// The values of each attribute associated with the node’s attached shader.
//
// # Discussion
//
// All nodes have their own copy of an attribute value and therefore the
// attribute values can be different per-node across the same [SKShader]. If
// instead you need all nodes to share the same value, use [SKUniform].
// Uniforms can change values every frame, but uniforms cannot vary per-node
// like attributes can.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/attributeValues
func (e SKEffectNode) AttributeValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("attributeValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e SKEffectNode) SetAttributeValues(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setAttributeValues:"), value)
}

// A Boolean value that indicates whether the results of rendering the child
// nodes should be cached.
//
// # Discussion
//
// If the value of this property is true, the effect node caches the filtered
// image for use in future frames. If the value is false, then SpriteKit
// discards the rendered image and redraws it from scratch the next time the
// node is rendered. The default value is false. Caching the rendered image
// uses more memory and may take more time to render. However, if the effect
// node’s descendants rarely change, caching can improve performance.
//
// When caching is enabled, changes to the effect node’s children trigger
// updates to the cached image in the next frame of animation. However,
// changing the filter’s properties does not.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/shouldRasterize
func (e SKEffectNode) ShouldRasterize() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("shouldRasterize"))
	return rv
}
func (e SKEffectNode) SetShouldRasterize(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setShouldRasterize:"), value)
}

// The blend mode used to draw the node’s contents into its parent’s
// framebuffer.
//
// # Discussion
//
// The default value is [SKBlendMode.alpha].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEffectNode/blendMode
//
// [SKBlendMode.alpha]: https://developer.apple.com/documentation/SpriteKit/SKBlendMode/alpha
func (e SKEffectNode) BlendMode() SKBlendMode {
	rv := objc.Send[SKBlendMode](e.ID, objc.Sel("blendMode"))
	return SKBlendMode(rv)
}
func (e SKEffectNode) SetBlendMode(value SKBlendMode) {
	objc.Send[struct{}](e.ID, objc.Sel("setBlendMode:"), value)
}

// Protocol methods for SKWarpable

// The maximum number of subdivision iterations used to generate the final
// vertices.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/subdivisionLevels
func (o SKEffectNode) SetSubdivisionLevels(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setSubdivisionLevels:"), value)
}

// The warp geometry used to define the distortion.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/warpGeometry
func (o SKEffectNode) SetWarpGeometry(value ISKWarpGeometry) {
	objc.Send[struct{}](o.ID, objc.Sel("setWarpGeometry:"), value)
}
