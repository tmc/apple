// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTileGroupRule] class.
var (
	_SKTileGroupRuleClass     SKTileGroupRuleClass
	_SKTileGroupRuleClassOnce sync.Once
)

func getSKTileGroupRuleClass() SKTileGroupRuleClass {
	_SKTileGroupRuleClassOnce.Do(func() {
		_SKTileGroupRuleClass = SKTileGroupRuleClass{class: objc.GetClass("SKTileGroupRule")}
	})
	return _SKTileGroupRuleClass
}

// GetSKTileGroupRuleClass returns the class object for SKTileGroupRule.
func GetSKTileGroupRuleClass() SKTileGroupRuleClass {
	return getSKTileGroupRuleClass()
}

type SKTileGroupRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTileGroupRuleClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTileGroupRuleClass) Alloc() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// Rules that describe how various tiles should be placed in a map.
//
// # Overview
//
// When a tile is filled in a tile map, the tile group rule defines how
// neighboring tiles are populated based on adjacency rules. A rule with
// multiple definitions uses the placement weights of the definitions to
// randomly select which to use.
//
// # Creating a Tile Group Rule
//
//   - [SKTileGroupRule.InitWithAdjacencyTileDefinitions]: Initializes a new tile group rule with adjacency rules and tile definitions.
//
// # Accessing or Setting Tile Group Rule Properties
//
//   - [SKTileGroupRule.Adjacency]: The adjacency requirement for this rule.
//   - [SKTileGroupRule.SetAdjacency]
//   - [SKTileGroupRule.Name]: A name associated with the tile group rule.
//   - [SKTileGroupRule.SetName]
//   - [SKTileGroupRule.TileDefinitions]: The tile definitions used for this rule.
//   - [SKTileGroupRule.SetTileDefinitions]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule
type SKTileGroupRule struct {
	objectivec.Object
}

// SKTileGroupRuleFromID constructs a [SKTileGroupRule] from an objc.ID.
//
// Rules that describe how various tiles should be placed in a map.
func SKTileGroupRuleFromID(id objc.ID) SKTileGroupRule {
	return SKTileGroupRule{objectivec.Object{ID: id}}
}

// NOTE: SKTileGroupRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTileGroupRule] class.
//
// # Creating a Tile Group Rule
//
//   - [ISKTileGroupRule.InitWithAdjacencyTileDefinitions]: Initializes a new tile group rule with adjacency rules and tile definitions.
//
// # Accessing or Setting Tile Group Rule Properties
//
//   - [ISKTileGroupRule.Adjacency]: The adjacency requirement for this rule.
//   - [ISKTileGroupRule.SetAdjacency]
//   - [ISKTileGroupRule.Name]: A name associated with the tile group rule.
//   - [ISKTileGroupRule.SetName]
//   - [ISKTileGroupRule.TileDefinitions]: The tile definitions used for this rule.
//   - [ISKTileGroupRule.SetTileDefinitions]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule
type ISKTileGroupRule interface {
	objectivec.IObject

	// Topic: Creating a Tile Group Rule

	// Initializes a new tile group rule with adjacency rules and tile definitions.
	InitWithAdjacencyTileDefinitions(adjacency SKTileAdjacencyMask, tileDefinitions []SKTileDefinition) SKTileGroupRule

	// Topic: Accessing or Setting Tile Group Rule Properties

	// The adjacency requirement for this rule.
	Adjacency() SKTileAdjacencyMask
	SetAdjacency(value SKTileAdjacencyMask)
	// A name associated with the tile group rule.
	Name() string
	SetName(value string)
	// The tile definitions used for this rule.
	TileDefinitions() []SKTileDefinition
	SetTileDefinitions(value []SKTileDefinition)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t SKTileGroupRule) Init() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTileGroupRule) Autorelease() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileGroupRule creates a new SKTileGroupRule instance.
func NewSKTileGroupRule() SKTileGroupRule {
	class := getSKTileGroupRuleClass()
	rv := objc.Send[SKTileGroupRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a new tile group rule with adjacency rules and tile
// definitions.
//
// adjacency: The adjacency requirements for this rule.
//
// tileDefinitions: The tile definitions used for this rule.
//
// # Return Value
//
// A new tile group rule.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/init(adjacency:tileDefinitions:)
func NewTileGroupRuleWithAdjacencyTileDefinitions(adjacency SKTileAdjacencyMask, tileDefinitions []SKTileDefinition) SKTileGroupRule {
	instance := getSKTileGroupRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAdjacency:tileDefinitions:"), adjacency, objectivec.IObjectSliceToNSArray(tileDefinitions))
	return SKTileGroupRuleFromID(rv)
}

// Initializes a new tile group rule with adjacency rules and tile
// definitions.
//
// adjacency: The adjacency requirements for this rule.
//
// tileDefinitions: The tile definitions used for this rule.
//
// # Return Value
//
// A new tile group rule.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/init(adjacency:tileDefinitions:)
func (t SKTileGroupRule) InitWithAdjacencyTileDefinitions(adjacency SKTileAdjacencyMask, tileDefinitions []SKTileDefinition) SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](t.ID, objc.Sel("initWithAdjacency:tileDefinitions:"), adjacency, objectivec.IObjectSliceToNSArray(tileDefinitions))
	return rv
}
func (t SKTileGroupRule) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The adjacency requirement for this rule.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/adjacency
func (t SKTileGroupRule) Adjacency() SKTileAdjacencyMask {
	rv := objc.Send[SKTileAdjacencyMask](t.ID, objc.Sel("adjacency"))
	return SKTileAdjacencyMask(rv)
}
func (t SKTileGroupRule) SetAdjacency(value SKTileAdjacencyMask) {
	objc.Send[struct{}](t.ID, objc.Sel("setAdjacency:"), value)
}

// A name associated with the tile group rule.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/name
func (t SKTileGroupRule) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (t SKTileGroupRule) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}

// The tile definitions used for this rule.
//
// # Discussion
//
// When this rule is evaluated and its conditions met, one of the definitions
// is randomly selected for placement based on their placement weights.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/tileDefinitions
func (t SKTileGroupRule) TileDefinitions() []SKTileDefinition {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("tileDefinitions"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKTileDefinition {
		return SKTileDefinitionFromID(id)
	})
}
func (t SKTileGroupRule) SetTileDefinitions(value []SKTileDefinition) {
	objc.Send[struct{}](t.ID, objc.Sel("setTileDefinitions:"), objectivec.IObjectSliceToNSArray(value))
}
