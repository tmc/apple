// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKSpriteNode] class.
var (
	_SKSpriteNodeClass     SKSpriteNodeClass
	_SKSpriteNodeClassOnce sync.Once
)

func getSKSpriteNodeClass() SKSpriteNodeClass {
	_SKSpriteNodeClassOnce.Do(func() {
		_SKSpriteNodeClass = SKSpriteNodeClass{class: objc.GetClass("SKSpriteNode")}
	})
	return _SKSpriteNodeClass
}

// GetSKSpriteNodeClass returns the class object for SKSpriteNode.
func GetSKSpriteNodeClass() SKSpriteNodeClass {
	return getSKSpriteNodeClass()
}

type SKSpriteNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKSpriteNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKSpriteNodeClass) Alloc() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An image or solid color.
//
// # Overview
//
// [SKSpriteNode] is an onscreen graphical element that can be initialized
// from an image or a solid color. SpriteKit adds functionality to its ability
// to display images using the functions discussed below.
//
// # Creating a Sprite from an Image Filename
//
//   - [SKSpriteNode.InitWithImageNamed]: Initializes a textured sprite using an image file.
//
// # Creating a Sprite from a Texture
//
//   - [SKSpriteNode.InitWithTexture]: Initializes a textured sprite using an existing texture object.
//   - [SKSpriteNode.InitWithTextureColorSize]: Initializes a textured sprite in color using an existing texture object.
//   - [SKSpriteNode.Texture]: The texture used to draw the sprite.
//   - [SKSpriteNode.SetTexture]
//
// # Creating a Solid-Color Sprite
//
//   - [SKSpriteNode.InitWithColorSize]: Initializes a single-color sprite node.
//
// # Setting a Sprite’s Size and Position
//
//   - [SKSpriteNode.Size]: The dimensions of the sprite, in points.
//   - [SKSpriteNode.SetSize]
//   - [SKSpriteNode.ScaleToSize]: Scales the sprite node to a specified size.
//   - [SKSpriteNode.AnchorPoint]: Defines the point in the sprite that corresponds to the node’s position.
//   - [SKSpriteNode.SetAnchorPoint]
//
// # Scaling a Sprite in Nine Parts
//
//   - [SKSpriteNode.CenterRect]: Enable nine-part stretching of the sprite’s texture.
//   - [SKSpriteNode.SetCenterRect]
//
// # Tinting a Sprite
//
//   - [SKSpriteNode.Color]: The sprite’s color.
//   - [SKSpriteNode.SetColor]
//   - [SKSpriteNode.ColorBlendFactor]: A floating-point value that describes how the color is blended with the sprite’s texture.
//   - [SKSpriteNode.SetColorBlendFactor]
//
// # Configuring Alpha Blendling
//
//   - [SKSpriteNode.BlendMode]: The blend mode used to draw the sprite into the parent’s framebuffer.
//   - [SKSpriteNode.SetBlendMode]
//
// # Lighting a Sprite
//
//   - [SKSpriteNode.LightingBitMask]: A mask that defines how this sprite is lit by light nodes in the scene.
//   - [SKSpriteNode.SetLightingBitMask]
//   - [SKSpriteNode.ShadowedBitMask]: A mask that defines which lights add shadows to the sprite.
//   - [SKSpriteNode.SetShadowedBitMask]
//   - [SKSpriteNode.ShadowCastBitMask]: A mask that defines which lights are occluded by this sprite.
//   - [SKSpriteNode.SetShadowCastBitMask]
//   - [SKSpriteNode.NormalTexture]: A texture that specifies the normal map for the sprite.
//   - [SKSpriteNode.SetNormalTexture]
//
// # Adding a Custom Shader to a Sprite
//
//   - [SKSpriteNode.Shader]: A text file that defines code that does custom per-pixel drawing or colorization.
//   - [SKSpriteNode.SetShader]
//   - [SKSpriteNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [SKSpriteNode.SetAttributeValues]
//   - [SKSpriteNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [SKSpriteNode.ValueForAttributeNamed]: Sets the value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode
type SKSpriteNode struct {
	SKNode
}

// SKSpriteNodeFromID constructs a [SKSpriteNode] from an objc.ID.
//
// An image or solid color.
func SKSpriteNodeFromID(id objc.ID) SKSpriteNode {
	return SKSpriteNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKSpriteNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKSpriteNode] class.
//
// # Creating a Sprite from an Image Filename
//
//   - [ISKSpriteNode.InitWithImageNamed]: Initializes a textured sprite using an image file.
//
// # Creating a Sprite from a Texture
//
//   - [ISKSpriteNode.InitWithTexture]: Initializes a textured sprite using an existing texture object.
//   - [ISKSpriteNode.InitWithTextureColorSize]: Initializes a textured sprite in color using an existing texture object.
//   - [ISKSpriteNode.Texture]: The texture used to draw the sprite.
//   - [ISKSpriteNode.SetTexture]
//
// # Creating a Solid-Color Sprite
//
//   - [ISKSpriteNode.InitWithColorSize]: Initializes a single-color sprite node.
//
// # Setting a Sprite’s Size and Position
//
//   - [ISKSpriteNode.Size]: The dimensions of the sprite, in points.
//   - [ISKSpriteNode.SetSize]
//   - [ISKSpriteNode.ScaleToSize]: Scales the sprite node to a specified size.
//   - [ISKSpriteNode.AnchorPoint]: Defines the point in the sprite that corresponds to the node’s position.
//   - [ISKSpriteNode.SetAnchorPoint]
//
// # Scaling a Sprite in Nine Parts
//
//   - [ISKSpriteNode.CenterRect]: Enable nine-part stretching of the sprite’s texture.
//   - [ISKSpriteNode.SetCenterRect]
//
// # Tinting a Sprite
//
//   - [ISKSpriteNode.Color]: The sprite’s color.
//   - [ISKSpriteNode.SetColor]
//   - [ISKSpriteNode.ColorBlendFactor]: A floating-point value that describes how the color is blended with the sprite’s texture.
//   - [ISKSpriteNode.SetColorBlendFactor]
//
// # Configuring Alpha Blendling
//
//   - [ISKSpriteNode.BlendMode]: The blend mode used to draw the sprite into the parent’s framebuffer.
//   - [ISKSpriteNode.SetBlendMode]
//
// # Lighting a Sprite
//
//   - [ISKSpriteNode.LightingBitMask]: A mask that defines how this sprite is lit by light nodes in the scene.
//   - [ISKSpriteNode.SetLightingBitMask]
//   - [ISKSpriteNode.ShadowedBitMask]: A mask that defines which lights add shadows to the sprite.
//   - [ISKSpriteNode.SetShadowedBitMask]
//   - [ISKSpriteNode.ShadowCastBitMask]: A mask that defines which lights are occluded by this sprite.
//   - [ISKSpriteNode.SetShadowCastBitMask]
//   - [ISKSpriteNode.NormalTexture]: A texture that specifies the normal map for the sprite.
//   - [ISKSpriteNode.SetNormalTexture]
//
// # Adding a Custom Shader to a Sprite
//
//   - [ISKSpriteNode.Shader]: A text file that defines code that does custom per-pixel drawing or colorization.
//   - [ISKSpriteNode.SetShader]
//   - [ISKSpriteNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [ISKSpriteNode.SetAttributeValues]
//   - [ISKSpriteNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [ISKSpriteNode.ValueForAttributeNamed]: Sets the value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode
type ISKSpriteNode interface {
	ISKNode

	// Topic: Creating a Sprite from an Image Filename

	// Initializes a textured sprite using an image file.
	InitWithImageNamed(name string) SKSpriteNode

	// Topic: Creating a Sprite from a Texture

	// Initializes a textured sprite using an existing texture object.
	InitWithTexture(texture ISKTexture) SKSpriteNode
	// Initializes a textured sprite in color using an existing texture object.
	InitWithTextureColorSize(texture ISKTexture, color appkit.NSColor, size corefoundation.CGSize) SKSpriteNode
	// The texture used to draw the sprite.
	Texture() ISKTexture
	SetTexture(value ISKTexture)

	// Topic: Creating a Solid-Color Sprite

	// Initializes a single-color sprite node.
	InitWithColorSize(color appkit.NSColor, size corefoundation.CGSize) SKSpriteNode

	// Topic: Setting a Sprite’s Size and Position

	// The dimensions of the sprite, in points.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	// Scales the sprite node to a specified size.
	ScaleToSize(size corefoundation.CGSize)
	// Defines the point in the sprite that corresponds to the node’s position.
	AnchorPoint() corefoundation.CGPoint
	SetAnchorPoint(value corefoundation.CGPoint)

	// Topic: Scaling a Sprite in Nine Parts

	// Enable nine-part stretching of the sprite’s texture.
	CenterRect() corefoundation.CGRect
	SetCenterRect(value corefoundation.CGRect)

	// Topic: Tinting a Sprite

	// The sprite’s color.
	Color() appkit.NSColor
	SetColor(value appkit.NSColor)
	// A floating-point value that describes how the color is blended with the sprite’s texture.
	ColorBlendFactor() float64
	SetColorBlendFactor(value float64)

	// Topic: Configuring Alpha Blendling

	// The blend mode used to draw the sprite into the parent’s framebuffer.
	BlendMode() SKBlendMode
	SetBlendMode(value SKBlendMode)

	// Topic: Lighting a Sprite

	// A mask that defines how this sprite is lit by light nodes in the scene.
	LightingBitMask() uint32
	SetLightingBitMask(value uint32)
	// A mask that defines which lights add shadows to the sprite.
	ShadowedBitMask() uint32
	SetShadowedBitMask(value uint32)
	// A mask that defines which lights are occluded by this sprite.
	ShadowCastBitMask() uint32
	SetShadowCastBitMask(value uint32)
	// A texture that specifies the normal map for the sprite.
	NormalTexture() ISKTexture
	SetNormalTexture(value ISKTexture)

	// Topic: Adding a Custom Shader to a Sprite

	// A text file that defines code that does custom per-pixel drawing or colorization.
	Shader() ISKShader
	SetShader(value ISKShader)
	// The values of each attribute associated with the node’s attached shader.
	AttributeValues() foundation.INSDictionary
	SetAttributeValues(value foundation.INSDictionary)
	// Sets an attribute value for an attached shader.
	SetValueForAttributeNamed(value ISKAttributeValue, key string)
	// Sets the value of a shader attribute.
	ValueForAttributeNamed(key string) ISKAttributeValue

	// The maximum number of subdivision iterations used to generate the final vertices.
	SubdivisionLevels() int
	// The warp geometry used to define the distortion.
	WarpGeometry() ISKWarpGeometry
}

// Init initializes the instance.
func (s SKSpriteNode) Init() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SKSpriteNode) Autorelease() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKSpriteNode creates a new SKSpriteNode instance.
func NewSKSpriteNode() SKSpriteNode {
	class := getSKSpriteNodeClass()
	rv := objc.Send[SKSpriteNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells you when to initialize a sprite from an archive.
//
// # Discussion
//
// Don’t call this function directly; the system calls this function when
// you should initialize your sprite from the argument archived data.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(coder:)
func NewSpriteNodeWithCoder(aDecoder foundation.INSCoder) SKSpriteNode {
	instance := getSKSpriteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKSpriteNodeFromID(rv)
}

// Initializes a single-color sprite node.
//
// color: The color for the resulting sprite node.
//
// size: The size of the sprite node in points.
//
// # Return Value
//
// A newly initialized sprite node.
//
// # Discussion
//
// Although textured nodes are the most common way to use the [SKSpriteNode]
// class, you can also create sprite nodes without a texture. The behavior of
// the class changes when the node lacks a texture:
//
// - The sprite node that is returned from this method has its
// [SKSpriteNode.Texture] property set to `nil`. - There is no texture to
// stretch, so the [SKSpriteNode.CenterRect] parameter is ignored. - There is
// no colorization step; the [SKSpriteNode.Color] property is used as the
// sprite’s color. - The sprite node’s [SKNode.Alpha] component is used to
// determine how it is blended into the buffer.
//
// Listing 1 shows how to create a red sprite node `100 x 100` points in size.
//
// Listing 1. Creating a non-textured sprite node
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(color:size:)
func NewSpriteNodeWithColorSize(color appkit.NSColor, size corefoundation.CGSize) SKSpriteNode {
	instance := getSKSpriteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithColor:size:"), color, size)
	return SKSpriteNodeFromID(rv)
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
func NewSpriteNodeWithFileNamed(filename string) SKSpriteNode {
	rv := objc.Send[objc.ID](objc.ID(getSKSpriteNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKSpriteNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewSpriteNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKSpriteNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKSpriteNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKSpriteNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKSpriteNode{}, objc.ErrInitFailed
	}
	return SKSpriteNodeFromID(rv), nil
}

// Initializes a textured sprite using an image file.
//
// name: The name of an image file stored in the app bundle.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// This method creates a new texture object from the image file and assigns
// that texture to the [SKSpriteNode.Texture] property, the
// [SKSpriteNode.NormalTexture] properties is set to `nil`. The
// [SKSpriteNode.Size] property of the sprite is set to the dimensions of the
// image. The [SKSpriteNode.Color] property is set to white with an alpha of
// zero `(1.0,1.0,1.0,0.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(imageNamed:)
func NewSpriteNodeWithImageNamed(name string) SKSpriteNode {
	instance := getSKSpriteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImageNamed:"), objc.String(name))
	return SKSpriteNodeFromID(rv)
}

// Initializes a textured sprite using an image file, optionally adding a
// normal map to simulate 3D lighting.
//
// name: The name of an image file stored in the app bundle.
//
// generateNormalMap: If true, a normal map is generated from the image texture without applying
// any filter to it (SKTextureNormalMapFilteringTypeNone). If false, no normal
// map is generated (matching the behavior of the [spriteNodeWithImageNamed:]
// class method).
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// The normal map is used only when lighting is enabled in the scene. For more
// information, see [SKSpriteNode] and [SKLightNode].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(imageNamed:normalMapped:)
//
// [spriteNodeWithImageNamed:]: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/spriteNodeWithImageNamed:
func NewSpriteNodeWithImageNamedNormalMapped(name string, generateNormalMap bool) SKSpriteNode {
	rv := objc.Send[objc.ID](objc.ID(getSKSpriteNodeClass().class), objc.Sel("spriteNodeWithImageNamed:normalMapped:"), objc.String(name), generateNormalMap)
	return SKSpriteNodeFromID(rv)
}

// Initializes a textured sprite using an existing texture object.
//
// texture: A SpriteKit texture.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// The [SKSpriteNode.Size] property of the sprite is set to the dimensions of
// the texture. The [SKSpriteNode.Color] property is set to white with an
// alpha of zero `(1.0,1.0,1.0,0.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(texture:)
func NewSpriteNodeWithTexture(texture ISKTexture) SKSpriteNode {
	instance := getSKSpriteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:"), texture)
	return SKSpriteNodeFromID(rv)
}

// Initializes a textured sprite in color using an existing texture object.
//
// texture: A texture to apply to the sprite.
//
// color: The color for the new sprite.
//
// size: The size for the new sprite.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// To colorize your texture, you also need to set the
// [SKSpriteNode.ColorBlendFactor] property of the sprite.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(texture:color:size:)
func NewSpriteNodeWithTextureColorSize(texture ISKTexture, color appkit.NSColor, size corefoundation.CGSize) SKSpriteNode {
	instance := getSKSpriteNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:color:size:"), texture, color, size)
	return SKSpriteNodeFromID(rv)
}

// Initializes a textured sprite with a normal map to simulate 3D lighting.
//
// texture: A SpriteKit texture used to draw the sprite.
//
// normalMap: A SpriteKit texture used to add lighting behavior to the sprite.
//
// # Return Value
//
// A newly initialized sprite object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(texture:normalMap:)
func NewSpriteNodeWithTextureNormalMap(texture ISKTexture, normalMap ISKTexture) SKSpriteNode {
	rv := objc.Send[objc.ID](objc.ID(getSKSpriteNodeClass().class), objc.Sel("spriteNodeWithTexture:normalMap:"), texture, normalMap)
	return SKSpriteNodeFromID(rv)
}

// Initializes a textured sprite using an existing texture object but with a
// specified size.
//
// texture: A SpriteKit texture.
//
// size: The size of the sprite in points.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// The sprite is initialized using the texture, but the texture’s dimensions
// are not used. Instead, the size passed into the constructor method is used.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(texture:size:)
func NewSpriteNodeWithTextureSize(texture ISKTexture, size corefoundation.CGSize) SKSpriteNode {
	rv := objc.Send[objc.ID](objc.ID(getSKSpriteNodeClass().class), objc.Sel("spriteNodeWithTexture:size:"), texture, size)
	return SKSpriteNodeFromID(rv)
}

// Initializes a textured sprite using an image file.
//
// name: The name of an image file stored in the app bundle.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// This method creates a new texture object from the image file and assigns
// that texture to the [SKSpriteNode.Texture] property, the
// [SKSpriteNode.NormalTexture] properties is set to `nil`. The
// [SKSpriteNode.Size] property of the sprite is set to the dimensions of the
// image. The [SKSpriteNode.Color] property is set to white with an alpha of
// zero `(1.0,1.0,1.0,0.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(imageNamed:)
func (s SKSpriteNode) InitWithImageNamed(name string) SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s.ID, objc.Sel("initWithImageNamed:"), objc.String(name))
	return rv
}

// Initializes a textured sprite using an existing texture object.
//
// texture: A SpriteKit texture.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// The [SKSpriteNode.Size] property of the sprite is set to the dimensions of
// the texture. The [SKSpriteNode.Color] property is set to white with an
// alpha of zero `(1.0,1.0,1.0,0.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(texture:)
func (s SKSpriteNode) InitWithTexture(texture ISKTexture) SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s.ID, objc.Sel("initWithTexture:"), texture)
	return rv
}

// Initializes a textured sprite in color using an existing texture object.
//
// texture: A texture to apply to the sprite.
//
// color: The color for the new sprite.
//
// size: The size for the new sprite.
//
// # Return Value
//
// A newly initialized sprite object.
//
// # Discussion
//
// To colorize your texture, you also need to set the
// [SKSpriteNode.ColorBlendFactor] property of the sprite.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(texture:color:size:)
func (s SKSpriteNode) InitWithTextureColorSize(texture ISKTexture, color appkit.NSColor, size corefoundation.CGSize) SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s.ID, objc.Sel("initWithTexture:color:size:"), texture, color, size)
	return rv
}

// Initializes a single-color sprite node.
//
// color: The color for the resulting sprite node.
//
// size: The size of the sprite node in points.
//
// # Return Value
//
// A newly initialized sprite node.
//
// # Discussion
//
// Although textured nodes are the most common way to use the [SKSpriteNode]
// class, you can also create sprite nodes without a texture. The behavior of
// the class changes when the node lacks a texture:
//
// - The sprite node that is returned from this method has its
// [SKSpriteNode.Texture] property set to `nil`. - There is no texture to
// stretch, so the [SKSpriteNode.CenterRect] parameter is ignored. - There is
// no colorization step; the [SKSpriteNode.Color] property is used as the
// sprite’s color. - The sprite node’s [SKNode.Alpha] component is used to
// determine how it is blended into the buffer.
//
// Listing 1 shows how to create a red sprite node `100 x 100` points in size.
//
// Listing 1. Creating a non-textured sprite node
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/init(color:size:)
func (s SKSpriteNode) InitWithColorSize(color appkit.NSColor, size corefoundation.CGSize) SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s.ID, objc.Sel("initWithColor:size:"), color, size)
	return rv
}

// Scales the sprite node to a specified size.
//
// # Discussion
//
// This method works by setting the sprite node’s [SKNode.XScale] and
// [SKNode.YScale] to achieve the specified size in its parent’s coordinate
// space.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/scale(to:)
func (s SKSpriteNode) ScaleToSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](s.ID, objc.Sel("scaleToSize:"), size)
}

// Sets an attribute value for an attached shader.
//
// value: An attribute value object containing the scalar or vector value to set in
// the attached shader.
//
// key: The attribute name.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/setValue(_:forAttribute:)
func (s SKSpriteNode) SetValueForAttributeNamed(value ISKAttributeValue, key string) {
	objc.Send[objc.ID](s.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}

// Sets the value of a shader attribute.
//
// key: The attribute name.
//
// # Return Value
//
// An attribute value object containing the scalar or vector value or `nil` if
// no such attribute exists.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/value(forAttributeNamed:)
func (s SKSpriteNode) ValueForAttributeNamed(key string) ISKAttributeValue {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("valueForAttributeNamed:"), objc.String(key))
	return SKAttributeValueFromID(rv)
}

// The maximum number of subdivision iterations used to generate the final
// vertices.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/subdivisionLevels
func (s SKSpriteNode) SubdivisionLevels() int {
	rv := objc.Send[int](s.ID, objc.Sel("subdivisionLevels"))
	return rv
}

// The warp geometry used to define the distortion.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/warpGeometry
func (s SKSpriteNode) WarpGeometry() ISKWarpGeometry {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("warpGeometry"))
	return SKWarpGeometryFromID(rv)
}

// The texture used to draw the sprite.
//
// # Discussion
//
// If the value is `nil`, the sprite is drawn as a single-color rectangle
// using its [SKSpriteNode.Color] property. Otherwise, the texture is used to
// draw the sprite. The related properties affect how the texture is applied.
//
// SpriteKit automatically generates a texture for sprites when they are
// initialized with [SKSpriteNode.InitWithImageNamed].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/texture
func (s SKSpriteNode) Texture() ISKTexture {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("texture"))
	return SKTextureFromID(objc.ID(rv))
}
func (s SKSpriteNode) SetTexture(value ISKTexture) {
	objc.Send[struct{}](s.ID, objc.Sel("setTexture:"), value)
}

// The dimensions of the sprite, in points.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/size
func (s SKSpriteNode) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (s SKSpriteNode) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](s.ID, objc.Sel("setSize:"), value)
}

// Defines the point in the sprite that corresponds to the node’s position.
//
// # Discussion
//
// You specify the value for this property in the unit coordinate space. The
// default value is `(0.5,0.5)`, which means that the sprite is centered on
// its position.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/anchorPoint
func (s SKSpriteNode) AnchorPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("anchorPoint"))
	return corefoundation.CGPoint(rv)
}
func (s SKSpriteNode) SetAnchorPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](s.ID, objc.Sel("setAnchorPoint:"), value)
}

// Enable nine-part stretching of the sprite’s texture.
//
// # Discussion
//
// Controls how the texture is stretched to fill the SKSpriteNode.
//
// The argument rectangle is in the unit coordinate space with a default value
// of `(0,0)-(1.0,1.0)`, which indicates that the entire texture is stretched
// to fill the sprite.
//
// If instead you define a different rectangle, its coordinates are used to
// break the texture into a 3 x 3 grid that is scaled like the following:
//
// - The four corners of this grid are applied without performing any scaling.
// - The upper and lower-middle parts are scaled horizontally - The left and
// right-middle parts are scaled vertically - The center is scaled in all
// directions.
//
// This is what’s referred to as a 9-part scaling algorithm.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/centerRect
func (s SKSpriteNode) CenterRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("centerRect"))
	return corefoundation.CGRect(rv)
}
func (s SKSpriteNode) SetCenterRect(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setCenterRect:"), value)
}

// The sprite’s color.
//
// # Discussion
//
// If the [SKSpriteNode.Texture] property is non-`nil`, the red, green, and
// blue values of the color property are blended with the texture before the
// texture is drawn and the alpha property is ignored. If the
// [SKSpriteNode.Texture] property is `nil`, the color (including the alpha
// component) is used to draw a single-color rectangle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/color
func (s SKSpriteNode) Color() appkit.NSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("color"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (s SKSpriteNode) SetColor(value appkit.NSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setColor:"), value)
}

// A floating-point value that describes how the color is blended with the
// sprite’s texture.
//
// # Discussion
//
// The value must be a number between `0.0` and `1.0`, inclusive. The default
// value (`0.0`) indicates the color property is ignored and that the
// texture’s values should be used unmodified. For values greater than
// `0.0`, the texture is blended with the color before being drawn to the
// scene.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/colorBlendFactor
func (s SKSpriteNode) ColorBlendFactor() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("colorBlendFactor"))
	return rv
}
func (s SKSpriteNode) SetColorBlendFactor(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setColorBlendFactor:"), value)
}

// The blend mode used to draw the sprite into the parent’s framebuffer.
//
// # Discussion
//
// The default value is [SKBlendMode.alpha].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/blendMode
//
// [SKBlendMode.alpha]: https://developer.apple.com/documentation/SpriteKit/SKBlendMode/alpha
func (s SKSpriteNode) BlendMode() SKBlendMode {
	rv := objc.Send[SKBlendMode](s.ID, objc.Sel("blendMode"))
	return SKBlendMode(rv)
}
func (s SKSpriteNode) SetBlendMode(value SKBlendMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setBlendMode:"), value)
}

// A mask that defines how this sprite is lit by light nodes in the scene.
//
// # Discussion
//
// To determine whether this sprite is lit by a light node, the sprite’s
// [SKSpriteNode.LightingBitMask] property is tested against the light’s
// [categoryBitMask] property by performing a logical AND operation. If the
// comparison results in a nonzero value, the sprite is lit by this light.
//
// The default value is 0x00000000 (all bits cleared).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/lightingBitMask
//
// [categoryBitMask]: https://developer.apple.com/documentation/SceneKit/SCNLight/categoryBitMask
func (s SKSpriteNode) LightingBitMask() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("lightingBitMask"))
	return rv
}
func (s SKSpriteNode) SetLightingBitMask(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLightingBitMask:"), value)
}

// A mask that defines which lights add shadows to the sprite.
//
// # Discussion
//
// To determine whether this sprite is affected by being a shadow generated by
// a light, its [SKSpriteNode.ShadowCastBitMask] property is tested against
// the light’s [categoryBitMask] property by performing a logical AND
// operation. If the comparison results in a nonzero value, the sprite is
// drawn using a shadowed effect.
//
// The default value is 0x00000000 (all bits cleared).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/shadowedBitMask
//
// [categoryBitMask]: https://developer.apple.com/documentation/SceneKit/SCNLight/categoryBitMask
func (s SKSpriteNode) ShadowedBitMask() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("shadowedBitMask"))
	return rv
}
func (s SKSpriteNode) SetShadowedBitMask(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setShadowedBitMask:"), value)
}

// A mask that defines which lights are occluded by this sprite.
//
// # Discussion
//
// To determine whether this sprite blocks the light (casting a shadow), the
// sprite’s [SKSpriteNode.ShadowedBitMask] property is tested against the
// light’s [categoryBitMask] property by performing a logical AND operation.
// If the comparison results in a nonzero value, the sprite casts a shadow
// past itself.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/shadowCastBitMask
//
// [categoryBitMask]: https://developer.apple.com/documentation/SceneKit/SCNLight/categoryBitMask
func (s SKSpriteNode) ShadowCastBitMask() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("shadowCastBitMask"))
	return rv
}
func (s SKSpriteNode) SetShadowCastBitMask(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setShadowCastBitMask:"), value)
}

// A texture that specifies the normal map for the sprite.
//
// # Discussion
//
// A normal map texture is used when a sprite is lit, giving it a more
// realistic look with shadows and specular highlights. The texture must be a
// normal map texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/normalTexture
func (s SKSpriteNode) NormalTexture() ISKTexture {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("normalTexture"))
	return SKTextureFromID(objc.ID(rv))
}
func (s SKSpriteNode) SetNormalTexture(value ISKTexture) {
	objc.Send[struct{}](s.ID, objc.Sel("setNormalTexture:"), value)
}

// A text file that defines code that does custom per-pixel drawing or
// colorization.
//
// # Discussion
//
// The default value is `nil`, which means the default behavior for sprite
// rendering is performed. SpriteKit implements many sprite features using a
// default shader, such as:
//
// - Animations on [SKNode.Alpha]. - [SKTexture] [SKTexture.FilteringMode]. -
// Light from [SKLightNode].
//
// If you supply a custom value for `shader`, your custom shader overrides the
// default shader which neutralizes the default features. It is the
// responsibility of your custom shader to implement any of the features your
// sprites require.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/shader
func (s SKSpriteNode) Shader() ISKShader {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("shader"))
	return SKShaderFromID(objc.ID(rv))
}
func (s SKSpriteNode) SetShader(value ISKShader) {
	objc.Send[struct{}](s.ID, objc.Sel("setShader:"), value)
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
// See: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode/attributeValues
func (s SKSpriteNode) AttributeValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attributeValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SKSpriteNode) SetAttributeValues(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setAttributeValues:"), value)
}

// Protocol methods for SKWarpable

// The maximum number of subdivision iterations used to generate the final
// vertices.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/subdivisionLevels
func (o SKSpriteNode) SetSubdivisionLevels(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setSubdivisionLevels:"), value)
}

// The warp geometry used to define the distortion.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpable/warpGeometry
func (o SKSpriteNode) SetWarpGeometry(value ISKWarpGeometry) {
	objc.Send[struct{}](o.ID, objc.Sel("setWarpGeometry:"), value)
}
