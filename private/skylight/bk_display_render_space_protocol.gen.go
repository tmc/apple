// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// BKDisplayRenderSpace protocol.
type BKDisplayRenderSpace interface {
	objectivec.IObject

	// AddDisplayBlankingObserver protocol.
	AddDisplayBlankingObserver(observer objectivec.IObject) objectivec.IObject

	// ClientTaskNameForTargetIDDisplayUUID protocol.
	ClientTaskNameForTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) uint32

	// ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID protocol.
	ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint

	// ConvertReferenceLocationToScreenLocationForDisplayUUID protocol.
	ConvertReferenceLocationToScreenLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint

	// ConvertScreenLocationToReferenceLocationForDisplayUUID protocol.
	ConvertScreenLocationToReferenceLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint

	// ConvertScreenLocationToTargetIDDisplayUUID protocol.
	ConvertScreenLocationToTargetIDDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint

	// DisplayForUUID protocol.
	DisplayForUUID(uuid objectivec.IObject) objectivec.IObject

	// DisplayIsBlanked protocol.
	DisplayIsBlanked(blanked objectivec.IObject) bool

	// GeometryForDisplayUUID protocol.
	GeometryForDisplayUUID(uuid objectivec.IObject) unsafe.Pointer

	// GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID protocol.
	GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID(port *uint32, identifier *uint64, id unsafe.Pointer, uuid objectivec.IObject)

	// HitTestLayerInformationFromHitTestResults protocol.
	HitTestLayerInformationFromHitTestResults(results unsafe.Pointer) objectivec.IObject

	// HostTargetIDForEmbeddedTargetIDDisplayUUID protocol.
	HostTargetIDForEmbeddedTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) unsafe.Pointer

	// PidForTargetIdentifier protocol.
	PidForTargetIdentifier(identifier unsafe.Pointer) int

	// ScaleForDisplay protocol.
	ScaleForDisplay(display objectivec.IObject) float64

	// TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults protocol.
	TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer) bool

	// TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult protocol.
	TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer, result []objectivec.IObject) bool

	// TransformFromLayerIdContextIDDisplayUUID protocol.
	TransformFromLayerIdContextIDDisplayUUID(id objectivec.IObject, id2 objectivec.IObject, uuid objectivec.IObject) quartzcore.CATransform3D
}

// BKDisplayRenderSpaceObject wraps an existing Objective-C object that conforms to the BKDisplayRenderSpace protocol.
type BKDisplayRenderSpaceObject struct {
	objectivec.Object
}

func (o BKDisplayRenderSpaceObject) BaseObject() objectivec.Object {
	return o.Object
}

// BKDisplayRenderSpaceObjectFromID constructs a [BKDisplayRenderSpaceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func BKDisplayRenderSpaceObjectFromID(id objc.ID) BKDisplayRenderSpaceObject {
	return BKDisplayRenderSpaceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o BKDisplayRenderSpaceObject) AddDisplayBlankingObserver(observer objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("addDisplayBlankingObserver:"), observer)
	return objectivec.Object{ID: rv}
}
func (o BKDisplayRenderSpaceObject) ClientTaskNameForTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("clientTaskNameForTargetID:displayUUID:"), id, uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertContextLocation:inTargetID:toScreenLocationForDisplayUUID:"), location, id, uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) ConvertReferenceLocationToScreenLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertReferenceLocation:toScreenLocationForDisplayUUID:"), location, uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) ConvertScreenLocationToReferenceLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertScreenLocation:toReferenceLocationForDisplayUUID:"), location, uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) ConvertScreenLocationToTargetIDDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](o.ID, objc.Sel("convertScreenLocation:toTargetID:displayUUID:"), location, id, uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) DisplayForUUID(uuid objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("displayForUUID:"), uuid)
	return objectivec.Object{ID: rv}
}
func (o BKDisplayRenderSpaceObject) DisplayIsBlanked(blanked objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("displayIsBlanked:"), blanked)
	return rv
}
func (o BKDisplayRenderSpaceObject) GeometryForDisplayUUID(uuid objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("geometryForDisplayUUID:"), uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID(port *uint32, identifier *uint64, id unsafe.Pointer, uuid objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("getClientTaskNamePort:clientConnectionIdentifier:forTargetID:displayUUID:"), port, identifier, id, uuid)
}
func (o BKDisplayRenderSpaceObject) HitTestLayerInformationFromHitTestResults(results unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("hitTestLayerInformationFromHitTestResults:"), results)
	return objectivec.Object{ID: rv}
}
func (o BKDisplayRenderSpaceObject) HostTargetIDForEmbeddedTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("hostTargetIDForEmbeddedTargetID:displayUUID:"), id, uuid)
	return rv
}
func (o BKDisplayRenderSpaceObject) PidForTargetIdentifier(identifier unsafe.Pointer) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("pidForTargetIdentifier:"), identifier)
	return rv
}
func (o BKDisplayRenderSpaceObject) ScaleForDisplay(display objectivec.IObject) float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("scaleForDisplay:"), display)
	return rv
}
func (o BKDisplayRenderSpaceObject) TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("targetIDAtScreenLocation:displayUUID:options:securityAnalysis:results:"), location, uuid, options, objectivec.IObjectSliceToNSArray(analysis), results)
	return rv
}
func (o BKDisplayRenderSpaceObject) TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer, result []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("targetIDAtScreenLocation:displayUUID:options:securityAnalysis:results:hitTestPassthroughTargetsResult:"), location, uuid, options, objectivec.IObjectSliceToNSArray(analysis), results, objectivec.IObjectSliceToNSArray(result))
	return rv
}
func (o BKDisplayRenderSpaceObject) TransformFromLayerIdContextIDDisplayUUID(id objectivec.IObject, id2 objectivec.IObject, uuid objectivec.IObject) quartzcore.CATransform3D {
	rv := objc.SendIfResponds[quartzcore.CATransform3D](o.ID, objc.Sel("transformFromLayerId:contextID:displayUUID:"), id, id2, uuid)
	return rv
}
