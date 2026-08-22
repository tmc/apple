// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTileSet] class.
var (
	_SKTileSetClass     SKTileSetClass
	_SKTileSetClassOnce sync.Once
)

func getSKTileSetClass() SKTileSetClass {
	_SKTileSetClassOnce.Do(func() {
		_SKTileSetClass = SKTileSetClass{class: objc.GetClass("SKTileSet")}
	})
	return _SKTileSetClass
}

// GetSKTileSetClass returns the class object for SKTileSet.
func GetSKTileSetClass() SKTileSetClass {
	return getSKTileSetClass()
}

type SKTileSetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTileSetClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTileSetClass) Alloc() SKTileSet {
	rv := objc.Send[SKTileSet](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A container for related tile groups.
//
// # Overview
//
// An [SKTileSet] object contains an array of tile groups that define which
// tile definitions are available for use in a tile map.
//
// Tile sets also define the arrangement of tiles within a tile map. In
// addition to the default rectangular grid layout, tile sets can also define
// hexagonal and isometric layouts.
//
// # Creating a Tile Set Programmatically
//
//   - [SKTileSet.InitWithTileGroups]: Initializes a new tile set with an array of tile groups and rectangular grid layout.
//   - [SKTileSet.InitWithTileGroupsTileSetType]: Initializes a new tile set with an array of tile groups and specified layout.
//
// # Accessing or Reading a Tile Set’s Properties
//
//   - [SKTileSet.DefaultTileGroup]: The tile set’s default tile group.
//   - [SKTileSet.SetDefaultTileGroup]
//   - [SKTileSet.DefaultTileSize]: The tile set’s default tile size.
//   - [SKTileSet.SetDefaultTileSize]
//   - [SKTileSet.Name]: A name associated with the tile set.
//   - [SKTileSet.SetName]
//   - [SKTileSet.TileGroups]: The tile set’s array of tile group objects.
//   - [SKTileSet.SetTileGroups]
//   - [SKTileSet.Type]: The tile set’s type.
//   - [SKTileSet.SetType]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet
type SKTileSet struct {
	objectivec.Object
}

// SKTileSetFromID constructs a [SKTileSet] from an objc.ID.
//
// A container for related tile groups.
func SKTileSetFromID(id objc.ID) SKTileSet {
	return SKTileSet{objectivec.Object{ID: id}}
}

// NOTE: SKTileSet adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTileSet] class.
//
// # Creating a Tile Set Programmatically
//
//   - [ISKTileSet.InitWithTileGroups]: Initializes a new tile set with an array of tile groups and rectangular grid layout.
//   - [ISKTileSet.InitWithTileGroupsTileSetType]: Initializes a new tile set with an array of tile groups and specified layout.
//
// # Accessing or Reading a Tile Set’s Properties
//
//   - [ISKTileSet.DefaultTileGroup]: The tile set’s default tile group.
//   - [ISKTileSet.SetDefaultTileGroup]
//   - [ISKTileSet.DefaultTileSize]: The tile set’s default tile size.
//   - [ISKTileSet.SetDefaultTileSize]
//   - [ISKTileSet.Name]: A name associated with the tile set.
//   - [ISKTileSet.SetName]
//   - [ISKTileSet.TileGroups]: The tile set’s array of tile group objects.
//   - [ISKTileSet.SetTileGroups]
//   - [ISKTileSet.Type]: The tile set’s type.
//   - [ISKTileSet.SetType]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet
type ISKTileSet interface {
	objectivec.IObject

	// Topic: Creating a Tile Set Programmatically

	// Initializes a new tile set with an array of tile groups and rectangular grid layout.
	InitWithTileGroups(tileGroups []SKTileGroup) SKTileSet
	// Initializes a new tile set with an array of tile groups and specified layout.
	InitWithTileGroupsTileSetType(tileGroups []SKTileGroup, tileSetType SKTileSetType) SKTileSet

	// Topic: Accessing or Reading a Tile Set’s Properties

	// The tile set’s default tile group.
	DefaultTileGroup() ISKTileGroup
	SetDefaultTileGroup(value ISKTileGroup)
	// The tile set’s default tile size.
	DefaultTileSize() corefoundation.CGSize
	SetDefaultTileSize(value corefoundation.CGSize)
	// A name associated with the tile set.
	Name() string
	SetName(value string)
	// The tile set’s array of tile group objects.
	TileGroups() []SKTileGroup
	SetTileGroups(value []SKTileGroup)
	// The tile set’s type.
	Type() SKTileSetType
	SetType(value SKTileSetType)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SKTileSet) Init() SKTileSet {
	rv := objc.Send[SKTileSet](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTileSet) Autorelease() SKTileSet {
	rv := objc.Send[SKTileSet](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileSet creates a new SKTileSet instance.
func NewSKTileSet() SKTileSet {
	class := getSKTileSetClass()
	rv := objc.Send[SKTileSet](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a tile set from a URL to an archived .sks file.
//
// url: The URL of a tile set file.
//
// # Return Value
//
// A new tile set or `nil` if the URL doesn’t point to a valid tile set
// file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(from:)
func NewTileSetFromURL(url foundation.NSURL) SKTileSet {
	rv := objc.Send[objc.ID](objc.ID(getSKTileSetClass().class), objc.Sel("tileSetFromURL:"), url)
	return SKTileSetFromID(rv)
}

// Initializes a tile set by searching the app bundle for an archived
// `XCUIElementTypeSks` file by name.
//
// name: The name of the tile set to search for.
//
// # Return Value
//
// A new tile set or `nil` if a tile set with a matching name cannot be found.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(named:)
func NewTileSetNamed(name string) SKTileSet {
	rv := objc.Send[objc.ID](objc.ID(getSKTileSetClass().class), objc.Sel("tileSetNamed:"), objc.String(name))
	return SKTileSetFromID(rv)
}

// Initializes a new tile set with an array of tile groups and rectangular
// grid layout.
//
// tileGroups: An array of [SKTileGroup] objects from which to create the tile set from.
//
// # Return Value
//
// A new tile set.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(tileGroups:)
func NewTileSetWithTileGroups(tileGroups []SKTileGroup) SKTileSet {
	instance := getSKTileSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTileGroups:"), objectivec.IObjectSliceToNSArray(tileGroups))
	return SKTileSetFromID(rv)
}

// Initializes a new tile set with an array of tile groups and specified
// layout.
//
// tileGroups: An array of [SKTileGroup] objects from which to create the tile set from.
//
// tileSetType: The arrangement of the tiles.
//
// # Return Value
//
// A new tile set.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(tileGroups:tileSetType:)
func NewTileSetWithTileGroupsTileSetType(tileGroups []SKTileGroup, tileSetType SKTileSetType) SKTileSet {
	instance := getSKTileSetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTileGroups:tileSetType:"), objectivec.IObjectSliceToNSArray(tileGroups), tileSetType)
	return SKTileSetFromID(rv)
}

// Initializes a new tile set with an array of tile groups and rectangular
// grid layout.
//
// tileGroups: An array of [SKTileGroup] objects from which to create the tile set from.
//
// # Return Value
//
// A new tile set.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(tileGroups:)
func (t SKTileSet) InitWithTileGroups(tileGroups []SKTileGroup) SKTileSet {
	rv := objc.Send[SKTileSet](t.ID, objc.Sel("initWithTileGroups:"), objectivec.IObjectSliceToNSArray(tileGroups))
	return rv
}

// Initializes a new tile set with an array of tile groups and specified
// layout.
//
// tileGroups: An array of [SKTileGroup] objects from which to create the tile set from.
//
// tileSetType: The arrangement of the tiles.
//
// # Return Value
//
// A new tile set.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/init(tileGroups:tileSetType:)
func (t SKTileSet) InitWithTileGroupsTileSetType(tileGroups []SKTileGroup, tileSetType SKTileSetType) SKTileSet {
	rv := objc.Send[SKTileSet](t.ID, objc.Sel("initWithTileGroups:tileSetType:"), objectivec.IObjectSliceToNSArray(tileGroups), tileSetType)
	return rv
}
func (t SKTileSet) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The tile set’s default tile group.
//
// # Discussion
//
// With auto-mapping enabled, it is possible for some tiles to be removed
// because there is either no valid rule or a missing tile group for the
// required adjacency rule. In this situation, those tiles are replaced by the
// tile group specified by [SKTileSet.DefaultTileGroup].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/defaultTileGroup
func (t SKTileSet) DefaultTileGroup() ISKTileGroup {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultTileGroup"))
	return SKTileGroupFromID(objc.ID(rv))
}
func (t SKTileSet) SetDefaultTileGroup(value ISKTileGroup) {
	objc.Send[struct{}](t.ID, objc.Sel("setDefaultTileGroup:"), value)
}

// The tile set’s default tile size.
//
// # Discussion
//
// The default tile size an [SKTileMapNode] will use for its tiles
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/defaultTileSize
func (t SKTileSet) DefaultTileSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("defaultTileSize"))
	return corefoundation.CGSize(rv)
}
func (t SKTileSet) SetDefaultTileSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setDefaultTileSize:"), value)
}

// A name associated with the tile set.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/name
func (t SKTileSet) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t SKTileSet) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}

// The tile set’s array of tile group objects.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/tileGroups
func (t SKTileSet) TileGroups() []SKTileGroup {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tileGroups"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKTileGroup {
		return SKTileGroupFromID(id)
	})
}
func (t SKTileSet) SetTileGroups(value []SKTileGroup) {
	objc.Send[struct{}](t.ID, objc.Sel("setTileGroups:"), objectivec.IObjectSliceToNSArray(value))
}

// The tile set’s type.
//
// # Discussion
//
// The tile set’s type specifies how the tiles in the set will be arranged
// when placed in a tile map.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileSet/type
func (t SKTileSet) Type() SKTileSetType {
	rv := objc.Send[SKTileSetType](t.ID, objc.Sel("type"))
	return SKTileSetType(rv)
}
func (t SKTileSet) SetType(value SKTileSetType) {
	objc.Send[struct{}](t.ID, objc.Sel("setType:"), value)
}
