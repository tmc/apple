// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLDataTimelineProcessWindows protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows
type SLDataTimelineProcessWindows interface {
	objectivec.IObject

	// OffScreen protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/offScreen
	OffScreen() uint64

	// OnScreenOccluded protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/onScreenOccluded
	OnScreenOccluded() uint64

	// OnScreenVisible protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/onScreenVisible
	OnScreenVisible() uint64

	// OrderedOut protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/orderedOut
	OrderedOut() uint64
}

// SLDataTimelineProcessWindowsObject wraps an existing Objective-C object that conforms to the SLDataTimelineProcessWindows protocol.
type SLDataTimelineProcessWindowsObject struct {
	objectivec.Object
}

func (o SLDataTimelineProcessWindowsObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLDataTimelineProcessWindowsObjectFromID constructs a [SLDataTimelineProcessWindowsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLDataTimelineProcessWindowsObjectFromID(id objc.ID) SLDataTimelineProcessWindowsObject {
	return SLDataTimelineProcessWindowsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/offScreen
func (o SLDataTimelineProcessWindowsObject) OffScreen() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("offScreen"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/onScreenOccluded
func (o SLDataTimelineProcessWindowsObject) OnScreenOccluded() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("onScreenOccluded"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/onScreenVisible
func (o SLDataTimelineProcessWindowsObject) OnScreenVisible() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("onScreenVisible"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLDataTimelineProcessWindows/orderedOut
func (o SLDataTimelineProcessWindowsObject) OrderedOut() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("orderedOut"))
	return rv
}
