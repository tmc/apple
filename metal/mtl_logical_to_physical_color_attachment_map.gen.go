// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLLogicalToPhysicalColorAttachmentMap] class.
var (
	_MTLLogicalToPhysicalColorAttachmentMapClass     MTLLogicalToPhysicalColorAttachmentMapClass
	_MTLLogicalToPhysicalColorAttachmentMapClassOnce sync.Once
)

func getMTLLogicalToPhysicalColorAttachmentMapClass() MTLLogicalToPhysicalColorAttachmentMapClass {
	_MTLLogicalToPhysicalColorAttachmentMapClassOnce.Do(func() {
		_MTLLogicalToPhysicalColorAttachmentMapClass = MTLLogicalToPhysicalColorAttachmentMapClass{class: objc.GetClass("MTLLogicalToPhysicalColorAttachmentMap")}
	})
	return _MTLLogicalToPhysicalColorAttachmentMapClass
}

// GetMTLLogicalToPhysicalColorAttachmentMapClass returns the class object for MTLLogicalToPhysicalColorAttachmentMap.
func GetMTLLogicalToPhysicalColorAttachmentMapClass() MTLLogicalToPhysicalColorAttachmentMapClass {
	return getMTLLogicalToPhysicalColorAttachmentMapClass()
}

type MTLLogicalToPhysicalColorAttachmentMapClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLLogicalToPhysicalColorAttachmentMapClass) Alloc() MTLLogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[MTLLogicalToPhysicalColorAttachmentMap](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Allows you to easily specify color attachment remapping from logical to
// physical indices.
//
// # Instance Methods
//
//   - [MTLLogicalToPhysicalColorAttachmentMap.Reset]
//
// See: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap
type MTLLogicalToPhysicalColorAttachmentMap struct {
	objectivec.Object
}

// MTLLogicalToPhysicalColorAttachmentMapFromID constructs a [MTLLogicalToPhysicalColorAttachmentMap] from an objc.ID.
//
// Allows you to easily specify color attachment remapping from logical to
// physical indices.
func MTLLogicalToPhysicalColorAttachmentMapFromID(id objc.ID) MTLLogicalToPhysicalColorAttachmentMap {
	return MTLLogicalToPhysicalColorAttachmentMap{objectivec.Object{id}}
}
// NOTE: MTLLogicalToPhysicalColorAttachmentMap adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLLogicalToPhysicalColorAttachmentMap] class.
//
// # Instance Methods
//
//   - [IMTLLogicalToPhysicalColorAttachmentMap.Reset]
//
// See: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap
type IMTLLogicalToPhysicalColorAttachmentMap interface {
	objectivec.IObject

	// Topic: Instance Methods

	Reset()

	// Queries the physical color attachment index corresponding to a logical index.
	GetPhysicalIndexForLogicalIndex(logicalIndex uint) uint
	// Maps a physical color attachment index to a logical index.
	SetPhysicalIndexForLogicalIndex(physicalIndex uint, logicalIndex uint)
}





// Init initializes the instance.
func (l MTLLogicalToPhysicalColorAttachmentMap) Init() MTLLogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[MTLLogicalToPhysicalColorAttachmentMap](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MTLLogicalToPhysicalColorAttachmentMap) Autorelease() MTLLogicalToPhysicalColorAttachmentMap {
	rv := objc.Send[MTLLogicalToPhysicalColorAttachmentMap](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLLogicalToPhysicalColorAttachmentMap creates a new MTLLogicalToPhysicalColorAttachmentMap instance.
func NewMTLLogicalToPhysicalColorAttachmentMap() MTLLogicalToPhysicalColorAttachmentMap {
	class := getMTLLogicalToPhysicalColorAttachmentMapClass()
	rv := objc.Send[MTLLogicalToPhysicalColorAttachmentMap](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// See: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/reset()
func (l MTLLogicalToPhysicalColorAttachmentMap) Reset() {
	objc.Send[objc.ID](l.ID, objc.Sel("reset"))
}

// Queries the physical color attachment index corresponding to a logical
// index.
//
// See: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/getPhysicalIndexForLogicalIndex:
func (l MTLLogicalToPhysicalColorAttachmentMap) GetPhysicalIndexForLogicalIndex(logicalIndex uint) uint {
	rv := objc.Send[uint](l.ID, objc.Sel("getPhysicalIndexForLogicalIndex:"), logicalIndex)
	return rv
}

// Maps a physical color attachment index to a logical index.
//
// physicalIndex: Index of the color attachment’s physical mapping.
//
// logicalIndex: Index of the color attachment’s logical mapping.
//
// See: https://developer.apple.com/documentation/Metal/MTLLogicalToPhysicalColorAttachmentMap/setPhysicalIndex:forLogicalIndex:
func (l MTLLogicalToPhysicalColorAttachmentMap) SetPhysicalIndexForLogicalIndex(physicalIndex uint, logicalIndex uint) {
	objc.Send[objc.ID](l.ID, objc.Sel("setPhysicalIndex:forLogicalIndex:"), physicalIndex, logicalIndex)
}

































