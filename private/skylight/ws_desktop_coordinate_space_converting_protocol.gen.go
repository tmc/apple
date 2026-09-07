// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WSDesktopCoordinateSpaceConverting protocol.
type WSDesktopCoordinateSpaceConverting interface {
	objectivec.IObject

	// ConvertDesktopLocationToTargetLocationTargetID protocol.
	ConvertDesktopLocationToTargetLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint

	// ConvertDesktopRectToTargetRectTargetID protocol.
	ConvertDesktopRectToTargetRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect

	// ConvertReferenceLocationToScreenLocationForDisplayUUID protocol.
	ConvertReferenceLocationToScreenLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint

	// ConvertScreenLocationToDesktopLocationDisplayUUID protocol.
	ConvertScreenLocationToDesktopLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint

	// ConvertTargetLocationToDesktopLocationTargetID protocol.
	ConvertTargetLocationToDesktopLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint

	// ConvertTargetRectToDesktopRectTargetID protocol.
	ConvertTargetRectToDesktopRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect

	// FromDesktopTransformValueForTargetID protocol.
	FromDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject

	// ToDesktopTransformValueForTargetID protocol.
	ToDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject
}

// WSDesktopCoordinateSpaceConvertingObject wraps an existing Objective-C object that conforms to the WSDesktopCoordinateSpaceConverting protocol.
type WSDesktopCoordinateSpaceConvertingObject struct {
	objectivec.Object
}

func (o WSDesktopCoordinateSpaceConvertingObject) BaseObject() objectivec.Object {
	return o.Object
}

// WSDesktopCoordinateSpaceConvertingObjectFromID constructs a [WSDesktopCoordinateSpaceConvertingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WSDesktopCoordinateSpaceConvertingObjectFromID(id objc.ID) WSDesktopCoordinateSpaceConvertingObject {
	return WSDesktopCoordinateSpaceConvertingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WSDesktopCoordinateSpaceConvertingObject) ConvertDesktopLocationToTargetLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertDesktopLocationToTargetLocation:targetID:"), location, id)
	return rv
}
func (o WSDesktopCoordinateSpaceConvertingObject) ConvertDesktopRectToTargetRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect {
	rv := objc.SendIfResponds[corefoundation.CGRect](o.ID, objc.Sel("convertDesktopRectToTargetRect:targetID:"), rect, id)
	return rv
}
func (o WSDesktopCoordinateSpaceConvertingObject) ConvertReferenceLocationToScreenLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertReferenceLocation:toScreenLocationForDisplayUUID:"), location, uuid)
	return rv
}
func (o WSDesktopCoordinateSpaceConvertingObject) ConvertScreenLocationToDesktopLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertScreenLocationToDesktopLocation:displayUUID:"), location, uuid)
	return rv
}
func (o WSDesktopCoordinateSpaceConvertingObject) ConvertTargetLocationToDesktopLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertTargetLocationToDesktopLocation:targetID:"), location, id)
	return rv
}
func (o WSDesktopCoordinateSpaceConvertingObject) ConvertTargetRectToDesktopRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect {
	rv := objc.SendIfResponds[corefoundation.CGRect](o.ID, objc.Sel("convertTargetRectToDesktopRect:targetID:"), rect, id)
	return rv
}
func (o WSDesktopCoordinateSpaceConvertingObject) FromDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("fromDesktopTransformValueForTargetID:"), id)
	return objectivec.Object{ID: rv}
}
func (o WSDesktopCoordinateSpaceConvertingObject) ToDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("toDesktopTransformValueForTargetID:"), id)
	return objectivec.Object{ID: rv}
}
