// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTileDefinition] class.
var (
	_SKTileDefinitionClass     SKTileDefinitionClass
	_SKTileDefinitionClassOnce sync.Once
)

func getSKTileDefinitionClass() SKTileDefinitionClass {
	_SKTileDefinitionClassOnce.Do(func() {
		_SKTileDefinitionClass = SKTileDefinitionClass{class: objc.GetClass("SKTileDefinition")}
	})
	return _SKTileDefinitionClass
}

// GetSKTileDefinitionClass returns the class object for SKTileDefinition.
func GetSKTileDefinitionClass() SKTileDefinitionClass {
	return getSKTileDefinitionClass()
}

type SKTileDefinitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTileDefinitionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTileDefinitionClass) Alloc() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A single tile that can be repeated in a tile map.
//
// # Overview
//
// To define the visual representation of a single tile, you create an
// [SKTileDefinition] object with texture and size information. Tile
// definitions support separate normal textures, for simulating 3D lighting,
// and arrays of textures for animation with speed controlled by the
// [SKTileDefinition.TimePerFrame] property. Textures can be rotated in 90˚
// increments or flipped either vertically or horizontally.
//
// Once a tile definition has been created, you encapsulate it in a
// [SKTileGroup] which is added to a [SKTileSet] which, in turn, is displayed
// in the scene with a [SKTileMapNode].
//
// # Creating a Tile with a Texture
//
//   - [SKTileDefinition.InitWithTexture]: Initializes a new tile definition with a single texture.
//
// # Creating a Tile with a Normal Texture
//
//   - [SKTileDefinition.InitWithTextureNormalTextureSize]: Initializes a new tile definition with a single texture and separate normal texture for simulating 3D lighting.
//
// # Creating a Tile with a Size
//
//   - [SKTileDefinition.InitWithTextureSize]: Initializes a new tile definition of a specified size with a single texture.
//
// # Creating an Animated Tile
//
//   - [SKTileDefinition.InitWithTexturesNormalTexturesSizeTimePerFrame]: Initializes a new tile definition with arrays of textures and normal textures for animation.
//   - [SKTileDefinition.InitWithTexturesSizeTimePerFrame]: Initializes a new tile definition with an array of textures for animation.
//
// # Flipping a Tile Vertically or Horizontally
//
//   - [SKTileDefinition.FlipHorizontally]: A Boolean that flips the definition’s image vertically.
//   - [SKTileDefinition.SetFlipHorizontally]
//   - [SKTileDefinition.FlipVertically]: A Boolean that flips the definition’s image horizontally.
//   - [SKTileDefinition.SetFlipVertically]
//
// # Rotating a Tile
//
//   - [SKTileDefinition.Rotation]: The rotation of the tile definition in 90˚ increments.
//   - [SKTileDefinition.SetRotation]
//
// # Configure Animated Tile Properties
//
//   - [SKTileDefinition.Textures]: An array of [SKTexture](<https://developer.apple.com/documentation/SpriteKit/SKTexture>) objects that defines the tile definition object’s content.
//   - [SKTileDefinition.SetTextures]
//   - [SKTileDefinition.NormalTextures]: An array of [SKTexture](<https://developer.apple.com/documentation/SpriteKit/SKTexture>) objects used to generate the normals for the tile to simulate 3D lighting.
//   - [SKTileDefinition.SetNormalTextures]
//   - [SKTileDefinition.TimePerFrame]: The duration, in seconds, that each texture in the textures array is displayed before switching to the next texture in the sequence.
//   - [SKTileDefinition.SetTimePerFrame]
//
// # Reading or Adding a Tile’s Custom Data
//
//   - [SKTileDefinition.UserData]: A dictionary containing arbitrary data.
//   - [SKTileDefinition.SetUserData]
//
// # Reading or Adjusting a Tile’s Instance Properties
//
//   - [SKTileDefinition.Name]: A name associated with the tile definition.
//   - [SKTileDefinition.SetName]
//   - [SKTileDefinition.PlacementWeight]: The placement weight of the tile definition.
//   - [SKTileDefinition.SetPlacementWeight]
//   - [SKTileDefinition.Size]: The size of the tile definition in points.
//   - [SKTileDefinition.SetSize]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition
type SKTileDefinition struct {
	objectivec.Object
}

// SKTileDefinitionFromID constructs a [SKTileDefinition] from an objc.ID.
//
// A single tile that can be repeated in a tile map.
func SKTileDefinitionFromID(id objc.ID) SKTileDefinition {
	return SKTileDefinition{objectivec.Object{ID: id}}
}

// NOTE: SKTileDefinition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTileDefinition] class.
//
// # Creating a Tile with a Texture
//
//   - [ISKTileDefinition.InitWithTexture]: Initializes a new tile definition with a single texture.
//
// # Creating a Tile with a Normal Texture
//
//   - [ISKTileDefinition.InitWithTextureNormalTextureSize]: Initializes a new tile definition with a single texture and separate normal texture for simulating 3D lighting.
//
// # Creating a Tile with a Size
//
//   - [ISKTileDefinition.InitWithTextureSize]: Initializes a new tile definition of a specified size with a single texture.
//
// # Creating an Animated Tile
//
//   - [ISKTileDefinition.InitWithTexturesNormalTexturesSizeTimePerFrame]: Initializes a new tile definition with arrays of textures and normal textures for animation.
//   - [ISKTileDefinition.InitWithTexturesSizeTimePerFrame]: Initializes a new tile definition with an array of textures for animation.
//
// # Flipping a Tile Vertically or Horizontally
//
//   - [ISKTileDefinition.FlipHorizontally]: A Boolean that flips the definition’s image vertically.
//   - [ISKTileDefinition.SetFlipHorizontally]
//   - [ISKTileDefinition.FlipVertically]: A Boolean that flips the definition’s image horizontally.
//   - [ISKTileDefinition.SetFlipVertically]
//
// # Rotating a Tile
//
//   - [ISKTileDefinition.Rotation]: The rotation of the tile definition in 90˚ increments.
//   - [ISKTileDefinition.SetRotation]
//
// # Configure Animated Tile Properties
//
//   - [ISKTileDefinition.Textures]: An array of [SKTexture](<https://developer.apple.com/documentation/SpriteKit/SKTexture>) objects that defines the tile definition object’s content.
//   - [ISKTileDefinition.SetTextures]
//   - [ISKTileDefinition.NormalTextures]: An array of [SKTexture](<https://developer.apple.com/documentation/SpriteKit/SKTexture>) objects used to generate the normals for the tile to simulate 3D lighting.
//   - [ISKTileDefinition.SetNormalTextures]
//   - [ISKTileDefinition.TimePerFrame]: The duration, in seconds, that each texture in the textures array is displayed before switching to the next texture in the sequence.
//   - [ISKTileDefinition.SetTimePerFrame]
//
// # Reading or Adding a Tile’s Custom Data
//
//   - [ISKTileDefinition.UserData]: A dictionary containing arbitrary data.
//   - [ISKTileDefinition.SetUserData]
//
// # Reading or Adjusting a Tile’s Instance Properties
//
//   - [ISKTileDefinition.Name]: A name associated with the tile definition.
//   - [ISKTileDefinition.SetName]
//   - [ISKTileDefinition.PlacementWeight]: The placement weight of the tile definition.
//   - [ISKTileDefinition.SetPlacementWeight]
//   - [ISKTileDefinition.Size]: The size of the tile definition in points.
//   - [ISKTileDefinition.SetSize]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition
type ISKTileDefinition interface {
	objectivec.IObject

	// Topic: Creating a Tile with a Texture

	// Initializes a new tile definition with a single texture.
	InitWithTexture(texture ISKTexture) SKTileDefinition

	// Topic: Creating a Tile with a Normal Texture

	// Initializes a new tile definition with a single texture and separate normal texture for simulating 3D lighting.
	InitWithTextureNormalTextureSize(texture ISKTexture, normalTexture ISKTexture, size corefoundation.CGSize) SKTileDefinition

	// Topic: Creating a Tile with a Size

	// Initializes a new tile definition of a specified size with a single texture.
	InitWithTextureSize(texture ISKTexture, size corefoundation.CGSize) SKTileDefinition

	// Topic: Creating an Animated Tile

	// Initializes a new tile definition with arrays of textures and normal textures for animation.
	InitWithTexturesNormalTexturesSizeTimePerFrame(textures []SKTexture, normalTextures []SKTexture, size corefoundation.CGSize, timePerFrame float64) SKTileDefinition
	// Initializes a new tile definition with an array of textures for animation.
	InitWithTexturesSizeTimePerFrame(textures []SKTexture, size corefoundation.CGSize, timePerFrame float64) SKTileDefinition

	// Topic: Flipping a Tile Vertically or Horizontally

	// A Boolean that flips the definition’s image vertically.
	FlipHorizontally() bool
	SetFlipHorizontally(value bool)
	// A Boolean that flips the definition’s image horizontally.
	FlipVertically() bool
	SetFlipVertically(value bool)

	// Topic: Rotating a Tile

	// The rotation of the tile definition in 90˚ increments.
	Rotation() SKTileDefinitionRotation
	SetRotation(value SKTileDefinitionRotation)

	// Topic: Configure Animated Tile Properties

	// An array of [SKTexture](<https://developer.apple.com/documentation/SpriteKit/SKTexture>) objects that defines the tile definition object’s content.
	Textures() []SKTexture
	SetTextures(value []SKTexture)
	// An array of [SKTexture](<https://developer.apple.com/documentation/SpriteKit/SKTexture>) objects used to generate the normals for the tile to simulate 3D lighting.
	NormalTextures() []SKTexture
	SetNormalTextures(value []SKTexture)
	// The duration, in seconds, that each texture in the textures array is displayed before switching to the next texture in the sequence.
	TimePerFrame() float64
	SetTimePerFrame(value float64)

	// Topic: Reading or Adding a Tile’s Custom Data

	// A dictionary containing arbitrary data.
	UserData() foundation.INSDictionary
	SetUserData(value foundation.INSDictionary)

	// Topic: Reading or Adjusting a Tile’s Instance Properties

	// A name associated with the tile definition.
	Name() string
	SetName(value string)
	// The placement weight of the tile definition.
	PlacementWeight() uint
	SetPlacementWeight(value uint)
	// The size of the tile definition in points.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SKTileDefinition) Init() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTileDefinition) Autorelease() SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileDefinition creates a new SKTileDefinition instance.
func NewSKTileDefinition() SKTileDefinition {
	class := getSKTileDefinitionClass()
	rv := objc.Send[SKTileDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new tile definition with a single texture.
//
// texture: The texture to reference for the definition’s size and content.
//
// # Return Value
//
// A new tile definition.
//
// # Discussion
//
// The size of the newly created tile definition will be the same as the
// texture used to initialize it.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:)
func NewTileDefinitionWithTexture(texture ISKTexture) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:"), texture)
	return SKTileDefinitionFromID(rv)
}

// Initializes a new tile definition with a single texture and separate normal
// texture for simulating 3D lighting.
//
// texture: The texture to reference for the definition’s content.
//
// normalTexture: The texture to reference for generating normals to simulate 3D lighting.
//
// size: The size of the tile in points.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:normalTexture:size:)
func NewTileDefinitionWithTextureNormalTextureSize(texture ISKTexture, normalTexture ISKTexture, size corefoundation.CGSize) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:normalTexture:size:"), texture, normalTexture, size)
	return SKTileDefinitionFromID(rv)
}

// Initializes a new tile definition of a specified size with a single
// texture.
//
// texture: The texture to reference for the definition’s content.
//
// size: The size of the tile in points.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:size:)
func NewTileDefinitionWithTextureSize(texture ISKTexture, size corefoundation.CGSize) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTexture:size:"), texture, size)
	return SKTileDefinitionFromID(rv)
}

// Initializes a new tile definition with arrays of textures and normal
// textures for animation.
//
// textures: An array of textures to reference for the definition’s content.
//
// normalTextures: An array of textures to reference for generating normals to simulate 3D
// lighting.
//
// size: The size of the tile in points.
//
// timePerFrame: The duration, in seconds, that each texture is displayed.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(textures:normalTextures:size:timePerFrame:)
func NewTileDefinitionWithTexturesNormalTexturesSizeTimePerFrame(textures []SKTexture, normalTextures []SKTexture, size corefoundation.CGSize, timePerFrame float64) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextures:normalTextures:size:timePerFrame:"), objectivec.IObjectSliceToNSArray(textures), objectivec.IObjectSliceToNSArray(normalTextures), size, timePerFrame)
	return SKTileDefinitionFromID(rv)
}

// Initializes a new tile definition with an array of textures for animation.
//
// textures: An array of textures to reference for the definition’s size and content.
//
// size: The size of the tile in points.
//
// timePerFrame: The duration, in seconds, that each texture is displayed.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(textures:size:timePerFrame:)
func NewTileDefinitionWithTexturesSizeTimePerFrame(textures []SKTexture, size corefoundation.CGSize, timePerFrame float64) SKTileDefinition {
	instance := getSKTileDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextures:size:timePerFrame:"), objectivec.IObjectSliceToNSArray(textures), size, timePerFrame)
	return SKTileDefinitionFromID(rv)
}

// Initializes a new tile definition with a single texture.
//
// texture: The texture to reference for the definition’s size and content.
//
// # Return Value
//
// A new tile definition.
//
// # Discussion
//
// The size of the newly created tile definition will be the same as the
// texture used to initialize it.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:)
func (t SKTileDefinition) InitWithTexture(texture ISKTexture) SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("initWithTexture:"), texture)
	return rv
}

// Initializes a new tile definition with a single texture and separate normal
// texture for simulating 3D lighting.
//
// texture: The texture to reference for the definition’s content.
//
// normalTexture: The texture to reference for generating normals to simulate 3D lighting.
//
// size: The size of the tile in points.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:normalTexture:size:)
func (t SKTileDefinition) InitWithTextureNormalTextureSize(texture ISKTexture, normalTexture ISKTexture, size corefoundation.CGSize) SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("initWithTexture:normalTexture:size:"), texture, normalTexture, size)
	return rv
}

// Initializes a new tile definition of a specified size with a single
// texture.
//
// texture: The texture to reference for the definition’s content.
//
// size: The size of the tile in points.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(texture:size:)
func (t SKTileDefinition) InitWithTextureSize(texture ISKTexture, size corefoundation.CGSize) SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("initWithTexture:size:"), texture, size)
	return rv
}

// Initializes a new tile definition with arrays of textures and normal
// textures for animation.
//
// textures: An array of textures to reference for the definition’s content.
//
// normalTextures: An array of textures to reference for generating normals to simulate 3D
// lighting.
//
// size: The size of the tile in points.
//
// timePerFrame: The duration, in seconds, that each texture is displayed.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(textures:normalTextures:size:timePerFrame:)
func (t SKTileDefinition) InitWithTexturesNormalTexturesSizeTimePerFrame(textures []SKTexture, normalTextures []SKTexture, size corefoundation.CGSize, timePerFrame float64) SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("initWithTextures:normalTextures:size:timePerFrame:"), objectivec.IObjectSliceToNSArray(textures), objectivec.IObjectSliceToNSArray(normalTextures), size, timePerFrame)
	return rv
}

// Initializes a new tile definition with an array of textures for animation.
//
// textures: An array of textures to reference for the definition’s size and content.
//
// size: The size of the tile in points.
//
// timePerFrame: The duration, in seconds, that each texture is displayed.
//
// # Return Value
//
// A new tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/init(textures:size:timePerFrame:)
func (t SKTileDefinition) InitWithTexturesSizeTimePerFrame(textures []SKTexture, size corefoundation.CGSize, timePerFrame float64) SKTileDefinition {
	rv := objc.Send[SKTileDefinition](t.ID, objc.Sel("initWithTextures:size:timePerFrame:"), objectivec.IObjectSliceToNSArray(textures), size, timePerFrame)
	return rv
}
func (t SKTileDefinition) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean that flips the definition’s image vertically.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/flipHorizontally
func (t SKTileDefinition) FlipHorizontally() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("flipHorizontally"))
	return rv
}
func (t SKTileDefinition) SetFlipHorizontally(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFlipHorizontally:"), value)
}

// A Boolean that flips the definition’s image horizontally.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/flipVertically
func (t SKTileDefinition) FlipVertically() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("flipVertically"))
	return rv
}
func (t SKTileDefinition) SetFlipVertically(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFlipVertically:"), value)
}

// The rotation of the tile definition in 90˚ increments.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/rotation
func (t SKTileDefinition) Rotation() SKTileDefinitionRotation {
	rv := objc.Send[SKTileDefinitionRotation](t.ID, objc.Sel("rotation"))
	return SKTileDefinitionRotation(rv)
}
func (t SKTileDefinition) SetRotation(value SKTileDefinitionRotation) {
	objc.Send[struct{}](t.ID, objc.Sel("setRotation:"), value)
}

// An array of [SKTexture] objects that defines the tile definition object’s
// content.
//
// # Discussion
//
// If the tile is non-animated, this will be an array containing one textures.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/textures
func (t SKTileDefinition) Textures() []SKTexture {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textures"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKTexture {
		return SKTextureFromID(id)
	})
}
func (t SKTileDefinition) SetTextures(value []SKTexture) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextures:"), objectivec.IObjectSliceToNSArray(value))
}

// An array of [SKTexture] objects used to generate the normals for the tile
// to simulate 3D lighting.
//
// # Discussion
//
// If the tile is non-animated, this will be an array containing one texture.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/normalTextures
func (t SKTileDefinition) NormalTextures() []SKTexture {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("normalTextures"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKTexture {
		return SKTextureFromID(id)
	})
}
func (t SKTileDefinition) SetNormalTextures(value []SKTexture) {
	objc.Send[struct{}](t.ID, objc.Sel("setNormalTextures:"), objectivec.IObjectSliceToNSArray(value))
}

// The duration, in seconds, that each texture in the textures array is
// displayed before switching to the next texture in the sequence.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/timePerFrame
func (t SKTileDefinition) TimePerFrame() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("timePerFrame"))
	return rv
}
func (t SKTileDefinition) SetTimePerFrame(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setTimePerFrame:"), value)
}

// A dictionary containing arbitrary data.
//
// # Discussion
//
// You use this property to store your own data in a tile definition. For
// example, you might use this property to specify whether this tile is a
// platform that a player can land on.
//
// SpriteKit doesn’t do anything with this data. However, the data is
// archived when the tile definition is archived.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/userData
func (t SKTileDefinition) UserData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("userData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t SKTileDefinition) SetUserData(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setUserData:"), value)
}

// A name associated with the tile definition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/name
func (t SKTileDefinition) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t SKTileDefinition) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}

// The placement weight of the tile definition.
//
// # Discussion
//
// This value is used to determine how likely this tile definition is to be
// chosen for placement when a [SKTileGroupRule] has multiple tile definitions
// assigned to it. A higher value relative to the other definitions assigned
// to the rule make it more likely for this definition to be selected; lower
// values make it less likely. Defaults to 1. When set to 0, the definition
// will never be chosen as long as there is at least one other definition with
// a [SKTileDefinition.PlacementWeight] above 0.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/placementWeight
func (t SKTileDefinition) PlacementWeight() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("placementWeight"))
	return rv
}
func (t SKTileDefinition) SetPlacementWeight(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlacementWeight:"), value)
}

// The size of the tile definition in points.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinition/size
func (t SKTileDefinition) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (t SKTileDefinition) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setSize:"), value)
}
