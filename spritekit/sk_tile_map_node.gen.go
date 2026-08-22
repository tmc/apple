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

// The class instance for the [SKTileMapNode] class.
var (
	_SKTileMapNodeClass     SKTileMapNodeClass
	_SKTileMapNodeClassOnce sync.Once
)

func getSKTileMapNodeClass() SKTileMapNodeClass {
	_SKTileMapNodeClassOnce.Do(func() {
		_SKTileMapNodeClass = SKTileMapNodeClass{class: objc.GetClass("SKTileMapNode")}
	})
	return _SKTileMapNodeClass
}

// GetSKTileMapNodeClass returns the class object for SKTileMapNode.
func GetSKTileMapNodeClass() SKTileMapNodeClass {
	return getSKTileMapNodeClass()
}

type SKTileMapNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTileMapNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTileMapNodeClass) Alloc() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A two-dimensional array of images.
//
// # Overview
//
// [SKTileMapNode] does the work of laying out predefined tiles in a grid of
// any size. Typically, you configure 9-slice images (tile groups) in
// Xcode’s SpriteKit scene editor and paint the look of your tile map ahead
// of time versus configuring the tile map in code.
//
// As with sprite nodes, you can layer tile maps with different blend modes or
// control it with actions and physics, for example, for the purpose of
// parallax scrolling. The rendered tile map can be post processed with an
// [SKShader] to add effects such as motion blur or atmospheric perspective.
//
// To work with a tile map programmatically, you supply [SKTileMapNode] with a
// tile set that defines the tile definitions it can render. Then, fill each
// tile in the tile map with the [fill(with:)] method and set individual tiles
// with [setTileGroup(_:andTileDefinition:forColumn:row:)].
//
// # Controlling a Tile Map’s On-Screen Position Relative to its Origin
//
//   - [SKTileMapNode.AnchorPoint]: Defines the point in the tile map that corresponds to its [position](<https://developer.apple.com/documentation/SpriteKit/SKNode/position>).
//   - [SKTileMapNode.SetAnchorPoint]
//
// # Reading or Manually Configuring the Tile Map’s Size
//
//   - [SKTileMapNode.TileSize]: The size of each tile in points.
//   - [SKTileMapNode.SetTileSize]
//   - [SKTileMapNode.TileSet]: The tile set being used by this tile map. The tile map object can only display tiles that exist in this set.
//   - [SKTileMapNode.SetTileSet]
//   - [SKTileMapNode.NumberOfColumns]: The number of columns in the tile map
//   - [SKTileMapNode.SetNumberOfColumns]
//   - [SKTileMapNode.NumberOfRows]: The number of rows in the tile map.
//   - [SKTileMapNode.SetNumberOfRows]
//
// # Querying the Tile Map’s Properties
//
//   - [SKTileMapNode.CenterOfTileAtColumnRow]
//   - [SKTileMapNode.TileColumnIndexFromPosition]
//   - [SKTileMapNode.TileDefinitionAtColumnRow]
//   - [SKTileMapNode.TileGroupAtColumnRow]
//   - [SKTileMapNode.TileRowIndexFromPosition]: Returns the tile map node object’s tile row index for the specified position in points.
//   - [SKTileMapNode.MapSize]: The overall size of the tile map.
//
// # Tinting a Tile Map
//
//   - [SKTileMapNode.Color]: The base color for the tile map. The influence of the color over the tile map node’s textures is controlled by [colorBlendFactor](<https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/colorBlendFactor>).
//   - [SKTileMapNode.SetColor]
//   - [SKTileMapNode.ColorBlendFactor]: Controls the blending between the texture and the tile map object’s [color](<https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/color>). Values are clamped between zero and one where zero has no color blending and one has the maximum color blending.
//   - [SKTileMapNode.SetColorBlendFactor]
//
// # Lighting a Tile Map
//
//   - [SKTileMapNode.LightingBitMask]: A mask that defines how the tile map is lit by light nodes in the scene.
//   - [SKTileMapNode.SetLightingBitMask]
//
// # Configuring How Alpha Values Blend the Sprite
//
//   - [SKTileMapNode.BlendMode]: Defines the blend mode to use when compositing the tile map over other nodes.
//   - [SKTileMapNode.SetBlendMode]
//
// # Working with Custom Shaders
//
//   - [SKTileMapNode.Shader]: Defines a shader which is applied to each tile of the tile map.
//   - [SKTileMapNode.SetShader]
//   - [SKTileMapNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [SKTileMapNode.SetAttributeValues]
//   - [SKTileMapNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [SKTileMapNode.ValueForAttributeNamed]: The value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode
//
// [fill(with:)]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/fill(with:)
// [setTileGroup(_:andTileDefinition:forColumn:row:)]: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/setTileGroup(_:andTileDefinition:forColumn:row:)
type SKTileMapNode struct {
	SKNode
}

// SKTileMapNodeFromID constructs a [SKTileMapNode] from an objc.ID.
//
// A two-dimensional array of images.
func SKTileMapNodeFromID(id objc.ID) SKTileMapNode {
	return SKTileMapNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKTileMapNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTileMapNode] class.
//
// # Controlling a Tile Map’s On-Screen Position Relative to its Origin
//
//   - [ISKTileMapNode.AnchorPoint]: Defines the point in the tile map that corresponds to its [position](<https://developer.apple.com/documentation/SpriteKit/SKNode/position>).
//   - [ISKTileMapNode.SetAnchorPoint]
//
// # Reading or Manually Configuring the Tile Map’s Size
//
//   - [ISKTileMapNode.TileSize]: The size of each tile in points.
//   - [ISKTileMapNode.SetTileSize]
//   - [ISKTileMapNode.TileSet]: The tile set being used by this tile map. The tile map object can only display tiles that exist in this set.
//   - [ISKTileMapNode.SetTileSet]
//   - [ISKTileMapNode.NumberOfColumns]: The number of columns in the tile map
//   - [ISKTileMapNode.SetNumberOfColumns]
//   - [ISKTileMapNode.NumberOfRows]: The number of rows in the tile map.
//   - [ISKTileMapNode.SetNumberOfRows]
//
// # Querying the Tile Map’s Properties
//
//   - [ISKTileMapNode.CenterOfTileAtColumnRow]
//   - [ISKTileMapNode.TileColumnIndexFromPosition]
//   - [ISKTileMapNode.TileDefinitionAtColumnRow]
//   - [ISKTileMapNode.TileGroupAtColumnRow]
//   - [ISKTileMapNode.TileRowIndexFromPosition]: Returns the tile map node object’s tile row index for the specified position in points.
//   - [ISKTileMapNode.MapSize]: The overall size of the tile map.
//
// # Tinting a Tile Map
//
//   - [ISKTileMapNode.Color]: The base color for the tile map. The influence of the color over the tile map node’s textures is controlled by [colorBlendFactor](<https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/colorBlendFactor>).
//   - [ISKTileMapNode.SetColor]
//   - [ISKTileMapNode.ColorBlendFactor]: Controls the blending between the texture and the tile map object’s [color](<https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/color>). Values are clamped between zero and one where zero has no color blending and one has the maximum color blending.
//   - [ISKTileMapNode.SetColorBlendFactor]
//
// # Lighting a Tile Map
//
//   - [ISKTileMapNode.LightingBitMask]: A mask that defines how the tile map is lit by light nodes in the scene.
//   - [ISKTileMapNode.SetLightingBitMask]
//
// # Configuring How Alpha Values Blend the Sprite
//
//   - [ISKTileMapNode.BlendMode]: Defines the blend mode to use when compositing the tile map over other nodes.
//   - [ISKTileMapNode.SetBlendMode]
//
// # Working with Custom Shaders
//
//   - [ISKTileMapNode.Shader]: Defines a shader which is applied to each tile of the tile map.
//   - [ISKTileMapNode.SetShader]
//   - [ISKTileMapNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [ISKTileMapNode.SetAttributeValues]
//   - [ISKTileMapNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [ISKTileMapNode.ValueForAttributeNamed]: The value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode
type ISKTileMapNode interface {
	ISKNode

	// Topic: Controlling a Tile Map’s On-Screen Position Relative to its Origin

	// Defines the point in the tile map that corresponds to its [position](<https://developer.apple.com/documentation/SpriteKit/SKNode/position>).
	AnchorPoint() corefoundation.CGPoint
	SetAnchorPoint(value corefoundation.CGPoint)

	// Topic: Reading or Manually Configuring the Tile Map’s Size

	// The size of each tile in points.
	TileSize() corefoundation.CGSize
	SetTileSize(value corefoundation.CGSize)
	// The tile set being used by this tile map. The tile map object can only display tiles that exist in this set.
	TileSet() ISKTileSet
	SetTileSet(value ISKTileSet)
	// The number of columns in the tile map
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)
	// The number of rows in the tile map.
	NumberOfRows() uint
	SetNumberOfRows(value uint)

	// Topic: Querying the Tile Map’s Properties

	CenterOfTileAtColumnRow(column uint, row uint) corefoundation.CGPoint
	TileColumnIndexFromPosition(position corefoundation.CGPoint) uint
	TileDefinitionAtColumnRow(column uint, row uint) ISKTileDefinition
	TileGroupAtColumnRow(column uint, row uint) ISKTileGroup
	// Returns the tile map node object’s tile row index for the specified position in points.
	TileRowIndexFromPosition(position corefoundation.CGPoint) uint
	// The overall size of the tile map.
	MapSize() corefoundation.CGSize

	// Topic: Tinting a Tile Map

	// The base color for the tile map. The influence of the color over the tile map node’s textures is controlled by [colorBlendFactor](<https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/colorBlendFactor>).
	Color() appkit.NSColor
	SetColor(value appkit.NSColor)
	// Controls the blending between the texture and the tile map object’s [color](<https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/color>). Values are clamped between zero and one where zero has no color blending and one has the maximum color blending.
	ColorBlendFactor() float64
	SetColorBlendFactor(value float64)

	// Topic: Lighting a Tile Map

	// A mask that defines how the tile map is lit by light nodes in the scene.
	LightingBitMask() uint32
	SetLightingBitMask(value uint32)

	// Topic: Configuring How Alpha Values Blend the Sprite

	// Defines the blend mode to use when compositing the tile map over other nodes.
	BlendMode() SKBlendMode
	SetBlendMode(value SKBlendMode)

	// Topic: Working with Custom Shaders

	// Defines a shader which is applied to each tile of the tile map.
	Shader() ISKShader
	SetShader(value ISKShader)
	// The values of each attribute associated with the node’s attached shader.
	AttributeValues() foundation.INSDictionary
	SetAttributeValues(value foundation.INSDictionary)
	// Sets an attribute value for an attached shader.
	SetValueForAttributeNamed(value ISKAttributeValue, key string)
	// The value of a shader attribute.
	ValueForAttributeNamed(key string) ISKAttributeValue

	// When creating a tile map node programmatically, specifies whether the tile map uses automapping behavior like the scene editor.
	EnableAutomapping() bool
	SetEnableAutomapping(value bool)
}

// Init initializes the instance.
func (t SKTileMapNode) Init() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTileMapNode) Autorelease() SKTileMapNode {
	rv := objc.Send[SKTileMapNode](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileMapNode creates a new SKTileMapNode instance.
func NewSKTileMapNode() SKTileMapNode {
	class := getSKTileMapNodeClass()
	rv := objc.Send[SKTileMapNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewTileMapNodeWithCoder(aDecoder foundation.INSCoder) SKTileMapNode {
	instance := getSKTileMapNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKTileMapNodeFromID(rv)
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
func NewTileMapNodeWithFileNamed(filename string) SKTileMapNode {
	rv := objc.Send[objc.ID](objc.ID(getSKTileMapNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKTileMapNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewTileMapNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKTileMapNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKTileMapNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKTileMapNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKTileMapNode{}, objc.ErrInitFailed
	}
	return SKTileMapNodeFromID(rv), nil
}

// column: The column index of the tile.
//
// row: The row index of the tile.
//
// # Return Value
//
// The coordinates in points of the center of the tile for a given column and
// row.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/centerOfTile(atColumn:row:)
func (t SKTileMapNode) CenterOfTileAtColumnRow(column uint, row uint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("centerOfTileAtColumn:row:"), column, row)
	return corefoundation.CGPoint(rv)
}

// position: The position in the tile map to check.
//
// # Return Value
//
// The tile map node object’s tile column index for the specified position.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileColumnIndex(fromPosition:)
func (t SKTileMapNode) TileColumnIndexFromPosition(position corefoundation.CGPoint) uint {
	rv := objc.Send[uint](t.ID, objc.Sel("tileColumnIndexFromPosition:"), position)
	return rv
}

// column: The column index of the tile.
//
// row: The row index of the tile.
//
// # Return Value
//
// The tile definition for the tile at the specified column and row.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileDefinition(atColumn:row:)
func (t SKTileMapNode) TileDefinitionAtColumnRow(column uint, row uint) ISKTileDefinition {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tileDefinitionAtColumn:row:"), column, row)
	return SKTileDefinitionFromID(rv)
}

// column: The column index of the tile.
//
// row: The row index of the tile.
//
// # Return Value
//
// The tile group for the tile at the specified column and row.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileGroup(atColumn:row:)
func (t SKTileMapNode) TileGroupAtColumnRow(column uint, row uint) ISKTileGroup {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tileGroupAtColumn:row:"), column, row)
	return SKTileGroupFromID(rv)
}

// Returns the tile map node object’s tile row index for the specified
// position in points.
//
// position: The position in the tile map to check.
//
// # Return Value
//
// The tile map node object’s tile row index for the specified position.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileRowIndex(fromPosition:)
func (t SKTileMapNode) TileRowIndexFromPosition(position corefoundation.CGPoint) uint {
	rv := objc.Send[uint](t.ID, objc.Sel("tileRowIndexFromPosition:"), position)
	return rv
}

// Sets an attribute value for an attached shader.
//
// value: An attribute value object containing the scalar or vector value to set in
// the attached shader.
//
// key: The attribute name.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/setValue(_:forAttribute:)
func (t SKTileMapNode) SetValueForAttributeNamed(value ISKAttributeValue, key string) {
	objc.Send[objc.ID](t.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}

// The value of a shader attribute.
//
// key: The attribute name.
//
// # Return Value
//
// An attribute value object containing the scalar or vector value or `nil` if
// no such attribute exists.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/value(forAttributeNamed:)
func (t SKTileMapNode) ValueForAttributeNamed(key string) ISKAttributeValue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("valueForAttributeNamed:"), objc.String(key))
	return SKAttributeValueFromID(rv)
}

// Defines the point in the tile map that corresponds to its
// [SKNode.Position].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/anchorPoint
func (t SKTileMapNode) AnchorPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("anchorPoint"))
	return corefoundation.CGPoint(rv)
}
func (t SKTileMapNode) SetAnchorPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](t.ID, objc.Sel("setAnchorPoint:"), value)
}

// The size of each tile in points.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileSize
func (t SKTileMapNode) TileSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("tileSize"))
	return corefoundation.CGSize(rv)
}
func (t SKTileMapNode) SetTileSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setTileSize:"), value)
}

// The tile set being used by this tile map. The tile map object can only
// display tiles that exist in this set.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/tileSet
func (t SKTileMapNode) TileSet() ISKTileSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tileSet"))
	return SKTileSetFromID(objc.ID(rv))
}
func (t SKTileMapNode) SetTileSet(value ISKTileSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setTileSet:"), value)
}

// The number of columns in the tile map
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/numberOfColumns
func (t SKTileMapNode) NumberOfColumns() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("numberOfColumns"))
	return rv
}
func (t SKTileMapNode) SetNumberOfColumns(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setNumberOfColumns:"), value)
}

// The number of rows in the tile map.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/numberOfRows
func (t SKTileMapNode) NumberOfRows() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("numberOfRows"))
	return rv
}
func (t SKTileMapNode) SetNumberOfRows(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setNumberOfRows:"), value)
}

// The overall size of the tile map.
//
// # Discussion
//
// For a grid set type, the overall size, in points, of the node will be
// [SKTileMapNode.NumberOfColumns] `*` [SKTileMapNode.TileSize] `.` [width]
// wide and [SKWarpGeometryGrid.NumberOfRows] `*` [SKTileMapNode.TileSize] `.`
// [height] high.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/mapSize
//
// [height]: https://developer.apple.com/documentation/CoreFoundation/CGSize/height
// [width]: https://developer.apple.com/documentation/CoreFoundation/CGSize/width
func (t SKTileMapNode) MapSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("mapSize"))
	return corefoundation.CGSize(rv)
}

// The base color for the tile map. The influence of the color over the tile
// map node’s textures is controlled by [SKTileMapNode.ColorBlendFactor].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/color
func (t SKTileMapNode) Color() appkit.NSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("color"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (t SKTileMapNode) SetColor(value appkit.NSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setColor:"), value)
}

// Controls the blending between the texture and the tile map object’s
// [SKTileMapNode.Color]. Values are clamped between zero and one where zero
// has no color blending and one has the maximum color blending.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/colorBlendFactor
func (t SKTileMapNode) ColorBlendFactor() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("colorBlendFactor"))
	return rv
}
func (t SKTileMapNode) SetColorBlendFactor(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setColorBlendFactor:"), value)
}

// A mask that defines how the tile map is lit by light nodes in the scene.
//
// # Discussion
//
// To determine whether this sprite is lit by a light node, the sprite’s
// `lightingBitMask` property is tested against the light’s
// [categoryBitMask] property by performing a logical AND operation. If the
// comparison results in a nonzero value, the sprite is lit by this light.
//
// The default value is 0x00000000 (all bits cleared).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/lightingBitMask
//
// [categoryBitMask]: https://developer.apple.com/documentation/SceneKit/SCNLight/categoryBitMask
func (t SKTileMapNode) LightingBitMask() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("lightingBitMask"))
	return rv
}
func (t SKTileMapNode) SetLightingBitMask(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setLightingBitMask:"), value)
}

// Defines the blend mode to use when compositing the tile map over other
// nodes.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/blendMode
func (t SKTileMapNode) BlendMode() SKBlendMode {
	rv := objc.Send[SKBlendMode](t.ID, objc.Sel("blendMode"))
	return SKBlendMode(rv)
}
func (t SKTileMapNode) SetBlendMode(value SKBlendMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setBlendMode:"), value)
}

// Defines a shader which is applied to each tile of the tile map.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/shader
func (t SKTileMapNode) Shader() ISKShader {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("shader"))
	return SKShaderFromID(objc.ID(rv))
}
func (t SKTileMapNode) SetShader(value ISKShader) {
	objc.Send[struct{}](t.ID, objc.Sel("setShader:"), value)
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
// See: https://developer.apple.com/documentation/SpriteKit/SKTileMapNode/attributeValues
func (t SKTileMapNode) AttributeValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributeValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t SKTileMapNode) SetAttributeValues(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setAttributeValues:"), value)
}

// When creating a tile map node programmatically, specifies whether the tile
// map uses automapping behavior like the scene editor.
//
// See: https://developer.apple.com/documentation/spritekit/sktilemapnode/enableautomapping
func (t SKTileMapNode) EnableAutomapping() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("enableAutomapping"))
	return rv
}
func (t SKTileMapNode) SetEnableAutomapping(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEnableAutomapping:"), value)
}
