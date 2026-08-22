// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTileGroup] class.
var (
	_SKTileGroupClass     SKTileGroupClass
	_SKTileGroupClassOnce sync.Once
)

func getSKTileGroupClass() SKTileGroupClass {
	_SKTileGroupClassOnce.Do(func() {
		_SKTileGroupClass = SKTileGroupClass{class: objc.GetClass("SKTileGroup")}
	})
	return _SKTileGroupClass
}

// GetSKTileGroupClass returns the class object for SKTileGroup.
func GetSKTileGroupClass() SKTileGroupClass {
	return getSKTileGroupClass()
}

type SKTileGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTileGroupClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTileGroupClass) Alloc() SKTileGroup {
	rv := objc.Send[SKTileGroup](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A set of tiles that collectively define one type of terrain.
//
// # Overview
//
// An [SKTileGroup] object contains either the definition of a single tile or
// an array of [SKTileGroupRule] objects that define adjacency rules.
//
// You supply a tile group with either:
//
// - The definition of a single tile that can be used to populate a tile map
// node with a single texture. - An array of one or more tile group rules that
// allow for the automatic placement of textures dependent on their adjacency
// and the placement weights of their definitions. For example, a tile group
// may contain nine tile group rules containing the definitions of the central
// tile and eight edge tiles that, when placed adjacently, appear as a single
// object.
//
// The preferred method to create tile groups is to use the editor tools in
// Xcode. However, to work with SpriteKit’s tile support programmatically,
// see the following articles.
//
// # Creating Tile Groups
//
//   - [SKTileGroup.InitWithTileDefinition]: Creates and initializes a simple tile group with a single tile definition.
//   - [SKTileGroup.InitWithRules]: Creates and initializes a tile group with the specified tile group rules.
//
// # Accessing or Setting a Tile Group’s Properties
//
//   - [SKTileGroup.Name]: The receiver’s name.
//   - [SKTileGroup.SetName]
//   - [SKTileGroup.Rules]: An array of [SKTileGroupRule](<https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule>) objects that the tile group uses to determine tile placement.
//   - [SKTileGroup.SetRules]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup
type SKTileGroup struct {
	objectivec.Object
}

// SKTileGroupFromID constructs a [SKTileGroup] from an objc.ID.
//
// A set of tiles that collectively define one type of terrain.
func SKTileGroupFromID(id objc.ID) SKTileGroup {
	return SKTileGroup{objectivec.Object{ID: id}}
}

// NOTE: SKTileGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTileGroup] class.
//
// # Creating Tile Groups
//
//   - [ISKTileGroup.InitWithTileDefinition]: Creates and initializes a simple tile group with a single tile definition.
//   - [ISKTileGroup.InitWithRules]: Creates and initializes a tile group with the specified tile group rules.
//
// # Accessing or Setting a Tile Group’s Properties
//
//   - [ISKTileGroup.Name]: The receiver’s name.
//   - [ISKTileGroup.SetName]
//   - [ISKTileGroup.Rules]: An array of [SKTileGroupRule](<https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule>) objects that the tile group uses to determine tile placement.
//   - [ISKTileGroup.SetRules]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup
type ISKTileGroup interface {
	objectivec.IObject

	// Topic: Creating Tile Groups

	// Creates and initializes a simple tile group with a single tile definition.
	InitWithTileDefinition(tileDefinition ISKTileDefinition) SKTileGroup
	// Creates and initializes a tile group with the specified tile group rules.
	InitWithRules(rules []SKTileGroupRule) SKTileGroup

	// Topic: Accessing or Setting a Tile Group’s Properties

	// The receiver’s name.
	Name() string
	SetName(value string)
	// An array of [SKTileGroupRule](<https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule>) objects that the tile group uses to determine tile placement.
	Rules() []SKTileGroupRule
	SetRules(value []SKTileGroupRule)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SKTileGroup) Init() SKTileGroup {
	rv := objc.Send[SKTileGroup](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTileGroup) Autorelease() SKTileGroup {
	rv := objc.Send[SKTileGroup](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileGroup creates a new SKTileGroup instance.
func NewSKTileGroup() SKTileGroup {
	class := getSKTileGroupClass()
	rv := objc.Send[SKTileGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes a tile group with the specified tile group rules.
//
// rules: The tile group rules to determine tile placement.
//
// # Return Value
//
// A new tile group.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/init(rules:)
func NewTileGroupWithRules(rules []SKTileGroupRule) SKTileGroup {
	instance := getSKTileGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRules:"), objectivec.IObjectSliceToNSArray(rules))
	return SKTileGroupFromID(rv)
}

// Creates and initializes a simple tile group with a single tile definition.
//
// tileDefinition: The tile definition to place in a tile map.
//
// # Return Value
//
// A new tile group.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/init(tileDefinition:)
func NewTileGroupWithTileDefinition(tileDefinition ISKTileDefinition) SKTileGroup {
	instance := getSKTileGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTileDefinition:"), tileDefinition)
	return SKTileGroupFromID(rv)
}

// Creates and initializes a simple tile group with a single tile definition.
//
// tileDefinition: The tile definition to place in a tile map.
//
// # Return Value
//
// A new tile group.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/init(tileDefinition:)
func (t SKTileGroup) InitWithTileDefinition(tileDefinition ISKTileDefinition) SKTileGroup {
	rv := objc.Send[SKTileGroup](t.ID, objc.Sel("initWithTileDefinition:"), tileDefinition)
	return rv
}

// Creates and initializes a tile group with the specified tile group rules.
//
// rules: The tile group rules to determine tile placement.
//
// # Return Value
//
// A new tile group.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/init(rules:)
func (t SKTileGroup) InitWithRules(rules []SKTileGroupRule) SKTileGroup {
	rv := objc.Send[SKTileGroup](t.ID, objc.Sel("initWithRules:"), objectivec.IObjectSliceToNSArray(rules))
	return rv
}
func (t SKTileGroup) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates an empty tile that erases the existing tile at that location on a
// tile map.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/empty()
func (_SKTileGroupClass SKTileGroupClass) EmptyTileGroup() SKTileGroup {
	rv := objc.Send[objc.ID](objc.ID(_SKTileGroupClass.class), objc.Sel("emptyTileGroup"))
	return SKTileGroupFromID(rv)
}

// The receiver’s name.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/name
func (t SKTileGroup) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t SKTileGroup) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}

// An array of [SKTileGroupRule] objects that the tile group uses to determine
// tile placement.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroup/rules
func (t SKTileGroup) Rules() []SKTileGroupRule {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("rules"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKTileGroupRule {
		return SKTileGroupRuleFromID(id)
	})
}
func (t SKTileGroup) SetRules(value []SKTileGroupRule) {
	objc.Send[struct{}](t.ID, objc.Sel("setRules:"), objectivec.IObjectSliceToNSArray(value))
}
