// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKWarpGeometryGrid] class.
var (
	_SKWarpGeometryGridClass     SKWarpGeometryGridClass
	_SKWarpGeometryGridClassOnce sync.Once
)

func getSKWarpGeometryGridClass() SKWarpGeometryGridClass {
	_SKWarpGeometryGridClassOnce.Do(func() {
		_SKWarpGeometryGridClass = SKWarpGeometryGridClass{class: objc.GetClass("SKWarpGeometryGrid")}
	})
	return _SKWarpGeometryGridClass
}

// GetSKWarpGeometryGridClass returns the class object for SKWarpGeometryGrid.
func GetSKWarpGeometryGridClass() SKWarpGeometryGridClass {
	return getSKWarpGeometryGridClass()
}

type SKWarpGeometryGridClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKWarpGeometryGridClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKWarpGeometryGridClass) Alloc() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A definition for a grid-based deformation of nodes that conform to
// [SKWarpable].
//
// # Overview
//
// An [SKWarpGeometryGrid] exposes a 2D array of source positions, and set of
// destination positions with matching size, that allow you to define which
// sections of a node should be translated from the source positions to the
// destination positions. Conceptually, this forms two grids—a source grid
// and a destination grid—where the visual warping is accomplished by
// stretching or shrinking each section of the node as the source positions of
// the grid interpolate to their corresponding destination positions.
//
// # Accessing or Setting Warp Geometry Grid Size
//
//   - [SKWarpGeometryGrid.NumberOfColumns]: The object’s number of columns.
//   - [SKWarpGeometryGrid.NumberOfRows]: The object’s number of rows.
//   - [SKWarpGeometryGrid.VertexCount]: The object’s total number of vertices.
//
// # Accessing or Setting Grid Vertices
//
//   - [SKWarpGeometryGrid.DestPositionAtIndex]: Returns the destination position of a vertex.
//   - [SKWarpGeometryGrid.SourcePositionAtIndex]: Returns the source position of a vertex.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid
type SKWarpGeometryGrid struct {
	SKWarpGeometry
}

// SKWarpGeometryGridFromID constructs a [SKWarpGeometryGrid] from an objc.ID.
//
// A definition for a grid-based deformation of nodes that conform to
// [SKWarpable].
func SKWarpGeometryGridFromID(id objc.ID) SKWarpGeometryGrid {
	return SKWarpGeometryGrid{SKWarpGeometry: SKWarpGeometryFromID(id)}
}

// NOTE: SKWarpGeometryGrid adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKWarpGeometryGrid] class.
//
// # Accessing or Setting Warp Geometry Grid Size
//
//   - [ISKWarpGeometryGrid.NumberOfColumns]: The object’s number of columns.
//   - [ISKWarpGeometryGrid.NumberOfRows]: The object’s number of rows.
//   - [ISKWarpGeometryGrid.VertexCount]: The object’s total number of vertices.
//
// # Accessing or Setting Grid Vertices
//
//   - [ISKWarpGeometryGrid.DestPositionAtIndex]: Returns the destination position of a vertex.
//   - [ISKWarpGeometryGrid.SourcePositionAtIndex]: Returns the source position of a vertex.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid
type ISKWarpGeometryGrid interface {
	ISKWarpGeometry

	// Topic: Accessing or Setting Warp Geometry Grid Size

	// The object’s number of columns.
	NumberOfColumns() int
	// The object’s number of rows.
	NumberOfRows() int
	// The object’s total number of vertices.
	VertexCount() int

	// Topic: Accessing or Setting Grid Vertices

	// Returns the destination position of a vertex.
	DestPositionAtIndex(index int) [2]float32
	// Returns the source position of a vertex.
	SourcePositionAtIndex(index int) [2]float32
}

// Init initializes the instance.
func (w SKWarpGeometryGrid) Init() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w SKWarpGeometryGrid) Autorelease() SKWarpGeometryGrid {
	rv := objc.Send[SKWarpGeometryGrid](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKWarpGeometryGrid creates a new SKWarpGeometryGrid instance.
func NewSKWarpGeometryGrid() SKWarpGeometryGrid {
	class := getSKWarpGeometryGridClass()
	rv := objc.Send[SKWarpGeometryGrid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells you when to intialize a grid that was loaded from an archive.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/init(coder:)
func NewWarpGeometryGridWithCoder(aDecoder foundation.INSCoder) SKWarpGeometryGrid {
	instance := getSKWarpGeometryGridClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKWarpGeometryGridFromID(rv)
}

// Creates a warp geometry grid of a specified size.
//
// cols: The number of columns in the grid.
//
// rows: The number of rows in the grid.
//
// # Return Value
//
// A new warp geometry grid object.
//
// # Discussion
//
// Creating a warp geometry grid without explicit source and destination
// positions automatically generates the required position arrays. For
// example, a 2 column by 2 row grid would create two arrays containing nine
// positions each, beginning at `[0,0]` - for the bottom left position - and
// ending at `[1,1]` - for the top left position.
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/init(columns:rows:)
func NewWarpGeometryGridWithColumnsRows(cols int, rows int) SKWarpGeometryGrid {
	rv := objc.Send[objc.ID](objc.ID(getSKWarpGeometryGridClass().class), objc.Sel("gridWithColumns:rows:"), cols, rows)
	return SKWarpGeometryGridFromID(rv)
}

// Returns the destination position of a vertex.
//
// index: The index of the position vertex to query.
//
// # Return Value
//
// The normalized position of the specified vertex in `destPositions`.
//
// # Discussion
//
// The specified index must be between 0 and the warp geometry grid’s
// [SKWarpGeometryGrid.VertexCount] `- 1`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/destPosition(at:)
func (w SKWarpGeometryGrid) DestPositionAtIndex(index int) [2]float32 {
	rv := objc.Send[[2]float32](w.ID, objc.Sel("destPositionAtIndex:"), index)
	return [2]float32(rv)
}

// Returns the source position of a vertex.
//
// index: The index of the position vertex to query.
//
// # Return Value
//
// The normalized position of the specified vertex in `sourcePositions`.
//
// # Discussion
//
// The specified index must be between 0 and the warp geometry grid’s
// [SKWarpGeometryGrid.VertexCount] `- 1`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/sourcePosition(at:)
func (w SKWarpGeometryGrid) SourcePositionAtIndex(index int) [2]float32 {
	rv := objc.Send[[2]float32](w.ID, objc.Sel("sourcePositionAtIndex:"), index)
	return [2]float32(rv)
}

// The object’s number of columns.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/numberOfColumns
func (w SKWarpGeometryGrid) NumberOfColumns() int {
	rv := objc.Send[int](w.ID, objc.Sel("numberOfColumns"))
	return rv
}

// The object’s number of rows.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/numberOfRows
func (w SKWarpGeometryGrid) NumberOfRows() int {
	rv := objc.Send[int](w.ID, objc.Sel("numberOfRows"))
	return rv
}

// The object’s total number of vertices.
//
// # Discussion
//
// The number of vertices is the same as `(```numberOfColumns“ `+ 1) *
// (```numberOfRows“ `+ 1)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKWarpGeometryGrid/vertexCount
func (w SKWarpGeometryGrid) VertexCount() int {
	rv := objc.Send[int](w.ID, objc.Sel("vertexCount"))
	return rv
}
