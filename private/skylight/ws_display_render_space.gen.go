// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [WSDisplayRenderSpace] class.
var (
	_WSDisplayRenderSpaceClass     WSDisplayRenderSpaceClass
	_WSDisplayRenderSpaceClassOnce sync.Once
)

func getWSDisplayRenderSpaceClass() WSDisplayRenderSpaceClass {
	_WSDisplayRenderSpaceClassOnce.Do(func() {
		_WSDisplayRenderSpaceClass = WSDisplayRenderSpaceClass{class: objc.GetClass("WSDisplayRenderSpace")}
	})
	return _WSDisplayRenderSpaceClass
}

// GetWSDisplayRenderSpaceClass returns the class object for WSDisplayRenderSpace.
func GetWSDisplayRenderSpaceClass() WSDisplayRenderSpaceClass {
	return getWSDisplayRenderSpaceClass()
}

type WSDisplayRenderSpaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSDisplayRenderSpaceClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSDisplayRenderSpaceClass) Alloc() WSDisplayRenderSpace {
	rv := objc.SendIfResponds[WSDisplayRenderSpace](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSDisplayRenderSpace._convertDesktopLocationToScreenLocationDisplayUUID]
//   - [WSDisplayRenderSpace._windowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid]
//   - [WSDisplayRenderSpace._windowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid]
//   - [WSDisplayRenderSpace.AddDisplayBlankingObserver]
//   - [WSDisplayRenderSpace.ClientTaskNameForTargetIDDisplayUUID]
//   - [WSDisplayRenderSpace.ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID]
//   - [WSDisplayRenderSpace.ConvertDesktopLocationToTargetLocationTargetID]
//   - [WSDisplayRenderSpace.ConvertDesktopRectToTargetRectTargetID]
//   - [WSDisplayRenderSpace.ConvertReferenceLocationToScreenLocationForDisplayUUID]
//   - [WSDisplayRenderSpace.ConvertScreenLocationToReferenceLocationForDisplayUUID]
//   - [WSDisplayRenderSpace.ConvertScreenLocationToTargetIDDisplayUUID]
//   - [WSDisplayRenderSpace.ConvertScreenLocationToDesktopLocationDisplayUUID]
//   - [WSDisplayRenderSpace.ConvertTargetLocationToDesktopLocationTargetID]
//   - [WSDisplayRenderSpace.ConvertTargetRectToDesktopRectTargetID]
//   - [WSDisplayRenderSpace.DisplayForUUID]
//   - [WSDisplayRenderSpace.DisplayIsBlanked]
//   - [WSDisplayRenderSpace.FromDesktopTransformValueForTargetID]
//   - [WSDisplayRenderSpace.GeometryForDisplayUUID]
//   - [WSDisplayRenderSpace.GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID]
//   - [WSDisplayRenderSpace.HitTestLayerInformationFromHitTestResults]
//   - [WSDisplayRenderSpace.HostTargetIDForEmbeddedTargetIDDisplayUUID]
//   - [WSDisplayRenderSpace.PidForTargetIdentifier]
//   - [WSDisplayRenderSpace.ScaleForDisplay]
//   - [WSDisplayRenderSpace.TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults]
//   - [WSDisplayRenderSpace.TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult]
//   - [WSDisplayRenderSpace.ToDesktopTransformValueForTargetID]
//   - [WSDisplayRenderSpace.TransformFromLayerIdContextIDDisplayUUID]
//   - [WSDisplayRenderSpace.DebugDescription]
//   - [WSDisplayRenderSpace.Description]
//   - [WSDisplayRenderSpace.Hash]
//   - [WSDisplayRenderSpace.Superclass]
type WSDisplayRenderSpace struct {
	objectivec.Object
}

// WSDisplayRenderSpaceFromID constructs a [WSDisplayRenderSpace] from an objc.ID.
func WSDisplayRenderSpaceFromID(id objc.ID) WSDisplayRenderSpace {
	return WSDisplayRenderSpace{objectivec.Object{ID: id}}
}

// Ensure WSDisplayRenderSpace implements IWSDisplayRenderSpace.
var _ IWSDisplayRenderSpace = WSDisplayRenderSpace{}

// An interface definition for the [WSDisplayRenderSpace] class.
//
// # Methods
//
//   - [IWSDisplayRenderSpace._convertDesktopLocationToScreenLocationDisplayUUID]
//   - [IWSDisplayRenderSpace._windowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid]
//   - [IWSDisplayRenderSpace._windowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid]
//   - [IWSDisplayRenderSpace.AddDisplayBlankingObserver]
//   - [IWSDisplayRenderSpace.ClientTaskNameForTargetIDDisplayUUID]
//   - [IWSDisplayRenderSpace.ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID]
//   - [IWSDisplayRenderSpace.ConvertDesktopLocationToTargetLocationTargetID]
//   - [IWSDisplayRenderSpace.ConvertDesktopRectToTargetRectTargetID]
//   - [IWSDisplayRenderSpace.ConvertReferenceLocationToScreenLocationForDisplayUUID]
//   - [IWSDisplayRenderSpace.ConvertScreenLocationToReferenceLocationForDisplayUUID]
//   - [IWSDisplayRenderSpace.ConvertScreenLocationToTargetIDDisplayUUID]
//   - [IWSDisplayRenderSpace.ConvertScreenLocationToDesktopLocationDisplayUUID]
//   - [IWSDisplayRenderSpace.ConvertTargetLocationToDesktopLocationTargetID]
//   - [IWSDisplayRenderSpace.ConvertTargetRectToDesktopRectTargetID]
//   - [IWSDisplayRenderSpace.DisplayForUUID]
//   - [IWSDisplayRenderSpace.DisplayIsBlanked]
//   - [IWSDisplayRenderSpace.FromDesktopTransformValueForTargetID]
//   - [IWSDisplayRenderSpace.GeometryForDisplayUUID]
//   - [IWSDisplayRenderSpace.GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID]
//   - [IWSDisplayRenderSpace.HitTestLayerInformationFromHitTestResults]
//   - [IWSDisplayRenderSpace.HostTargetIDForEmbeddedTargetIDDisplayUUID]
//   - [IWSDisplayRenderSpace.PidForTargetIdentifier]
//   - [IWSDisplayRenderSpace.ScaleForDisplay]
//   - [IWSDisplayRenderSpace.TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults]
//   - [IWSDisplayRenderSpace.TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult]
//   - [IWSDisplayRenderSpace.ToDesktopTransformValueForTargetID]
//   - [IWSDisplayRenderSpace.TransformFromLayerIdContextIDDisplayUUID]
//   - [IWSDisplayRenderSpace.DebugDescription]
//   - [IWSDisplayRenderSpace.Description]
//   - [IWSDisplayRenderSpace.Hash]
//   - [IWSDisplayRenderSpace.Superclass]
type IWSDisplayRenderSpace interface {
	objectivec.IObject

	// Topic: Methods

	_convertDesktopLocationToScreenLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint
	_windowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid(space unsafe.Pointer, valid *bool) WSWindowTransform
	_windowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid(space unsafe.Pointer, valid *bool) WSWindowTransform
	AddDisplayBlankingObserver(observer objectivec.IObject) objectivec.IObject
	ClientTaskNameForTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) uint32
	ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint
	ConvertDesktopLocationToTargetLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint
	ConvertDesktopRectToTargetRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect
	ConvertReferenceLocationToScreenLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint
	ConvertScreenLocationToReferenceLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint
	ConvertScreenLocationToTargetIDDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint
	ConvertScreenLocationToDesktopLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint
	ConvertTargetLocationToDesktopLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint
	ConvertTargetRectToDesktopRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect
	DisplayForUUID(uuid objectivec.IObject) objectivec.IObject
	DisplayIsBlanked(blanked objectivec.IObject) bool
	FromDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject
	GeometryForDisplayUUID(uuid objectivec.IObject) unsafe.Pointer
	GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID(port *uint32, identifier *uint64, id unsafe.Pointer, uuid objectivec.IObject)
	HitTestLayerInformationFromHitTestResults(results unsafe.Pointer) objectivec.IObject
	HostTargetIDForEmbeddedTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) unsafe.Pointer
	PidForTargetIdentifier(identifier unsafe.Pointer) int32
	ScaleForDisplay(display objectivec.IObject) float64
	TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer) bool
	TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer, result []objectivec.IObject) bool
	ToDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject
	TransformFromLayerIdContextIDDisplayUUID(id objectivec.IObject, id2 objectivec.IObject, uuid objectivec.IObject) quartzcore.CATransform3D
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSDisplayRenderSpace) Init() WSDisplayRenderSpace {
	rv := objc.SendIfResponds[WSDisplayRenderSpace](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSDisplayRenderSpace) Autorelease() WSDisplayRenderSpace {
	rv := objc.SendIfResponds[WSDisplayRenderSpace](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSDisplayRenderSpace creates a new WSDisplayRenderSpace instance.
func NewWSDisplayRenderSpace() WSDisplayRenderSpace {
	class := getWSDisplayRenderSpaceClass()
	rv := objc.SendIfResponds[WSDisplayRenderSpace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSDisplayRenderSpace) _convertDesktopLocationToScreenLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("_convertDesktopLocationToScreenLocation:displayUUID:"), location, uuid)
	return corefoundation.CGPoint(rv)
}

// ConvertDesktopLocationToScreenLocationDisplayUUID is an exported wrapper for the private method _convertDesktopLocationToScreenLocationDisplayUUID.
func (w WSDisplayRenderSpace) ConvertDesktopLocationToScreenLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) (corefoundation.CGPoint, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_convertDesktopLocationToScreenLocation:displayUUID:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_convertDesktopLocationToScreenLocation:displayUUID:"}
		return corefoundation.CGPoint{}, err
	}
	return w._convertDesktopLocationToScreenLocationDisplayUUID(location, uuid), nil
}

// CanConvertDesktopLocationToScreenLocationDisplayUUID reports whether the receiver responds to the private selector _convertDesktopLocationToScreenLocation:displayUUID:.
func (w WSDisplayRenderSpace) CanConvertDesktopLocationToScreenLocationDisplayUUID() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_convertDesktopLocationToScreenLocation:displayUUID:"))
}
func (w WSDisplayRenderSpace) _windowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid(space unsafe.Pointer, valid *bool) WSWindowTransform {
	rv := objc.SendIfResponds[WSWindowTransform](w.ID, objc.Sel("_windowTransformToConvertFromDesktopSpaceToTargetSpace:outIsValid:"), space, valid)
	return WSWindowTransform(rv)
}

// WindowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid is an exported wrapper for the private method _windowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid.
func (w WSDisplayRenderSpace) WindowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid(space unsafe.Pointer, valid *bool) (WSWindowTransform, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_windowTransformToConvertFromDesktopSpaceToTargetSpace:outIsValid:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_windowTransformToConvertFromDesktopSpaceToTargetSpace:outIsValid:"}
		return WSWindowTransform{}, err
	}
	return w._windowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid(space, valid), nil
}

// CanWindowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid reports whether the receiver responds to the private selector _windowTransformToConvertFromDesktopSpaceToTargetSpace:outIsValid:.
func (w WSDisplayRenderSpace) CanWindowTransformToConvertFromDesktopSpaceToTargetSpaceOutIsValid() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_windowTransformToConvertFromDesktopSpaceToTargetSpace:outIsValid:"))
}
func (w WSDisplayRenderSpace) _windowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid(space unsafe.Pointer, valid *bool) WSWindowTransform {
	rv := objc.SendIfResponds[WSWindowTransform](w.ID, objc.Sel("_windowTransformToConvertFromTargetSpaceToDesktopSpace:outIsValid:"), space, valid)
	return WSWindowTransform(rv)
}

// WindowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid is an exported wrapper for the private method _windowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid.
func (w WSDisplayRenderSpace) WindowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid(space unsafe.Pointer, valid *bool) (WSWindowTransform, error) {
	if !objc.RespondsToSelector(w.ID, objc.Sel("_windowTransformToConvertFromTargetSpaceToDesktopSpace:outIsValid:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_windowTransformToConvertFromTargetSpaceToDesktopSpace:outIsValid:"}
		return WSWindowTransform{}, err
	}
	return w._windowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid(space, valid), nil
}

// CanWindowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid reports whether the receiver responds to the private selector _windowTransformToConvertFromTargetSpaceToDesktopSpace:outIsValid:.
func (w WSDisplayRenderSpace) CanWindowTransformToConvertFromTargetSpaceToDesktopSpaceOutIsValid() bool {
	return objc.RespondsToSelector(w.ID, objc.Sel("_windowTransformToConvertFromTargetSpaceToDesktopSpace:outIsValid:"))
}
func (w WSDisplayRenderSpace) AddDisplayBlankingObserver(observer objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("addDisplayBlankingObserver:"), observer)
	return objectivec.Object{ID: rv}
}
func (w WSDisplayRenderSpace) ClientTaskNameForTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) uint32 {
	rv := objc.SendIfResponds[uint32](w.ID, objc.Sel("clientTaskNameForTargetID:displayUUID:"), id, uuid)
	return rv
}
func (w WSDisplayRenderSpace) ConvertContextLocationInTargetIDToScreenLocationForDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertContextLocation:inTargetID:toScreenLocationForDisplayUUID:"), location, id, uuid)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertDesktopLocationToTargetLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertDesktopLocationToTargetLocation:targetID:"), location, id)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertDesktopRectToTargetRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect {
	rv := objc.SendIfResponds[corefoundation.CGRect](w.ID, objc.Sel("convertDesktopRectToTargetRect:targetID:"), rect, id)
	return corefoundation.CGRect(rv)
}
func (w WSDisplayRenderSpace) ConvertReferenceLocationToScreenLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertReferenceLocation:toScreenLocationForDisplayUUID:"), location, uuid)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertScreenLocationToReferenceLocationForDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertScreenLocation:toReferenceLocationForDisplayUUID:"), location, uuid)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertScreenLocationToTargetIDDisplayUUID(location corefoundation.CGPoint, id unsafe.Pointer, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertScreenLocation:toTargetID:displayUUID:"), location, id, uuid)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertScreenLocationToDesktopLocationDisplayUUID(location corefoundation.CGPoint, uuid objectivec.IObject) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertScreenLocationToDesktopLocation:displayUUID:"), location, uuid)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertTargetLocationToDesktopLocationTargetID(location corefoundation.CGPoint, id unsafe.Pointer) corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](w.ID, objc.Sel("convertTargetLocationToDesktopLocation:targetID:"), location, id)
	return corefoundation.CGPoint(rv)
}
func (w WSDisplayRenderSpace) ConvertTargetRectToDesktopRectTargetID(rect corefoundation.CGRect, id unsafe.Pointer) corefoundation.CGRect {
	rv := objc.SendIfResponds[corefoundation.CGRect](w.ID, objc.Sel("convertTargetRectToDesktopRect:targetID:"), rect, id)
	return corefoundation.CGRect(rv)
}
func (w WSDisplayRenderSpace) DisplayForUUID(uuid objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("displayForUUID:"), uuid)
	return objectivec.Object{ID: rv}
}
func (w WSDisplayRenderSpace) DisplayIsBlanked(blanked objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("displayIsBlanked:"), blanked)
	return rv
}
func (w WSDisplayRenderSpace) FromDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("fromDesktopTransformValueForTargetID:"), id)
	return objectivec.Object{ID: rv}
}
func (w WSDisplayRenderSpace) GeometryForDisplayUUID(uuid objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("geometryForDisplayUUID:"), uuid)
	return rv
}
func (w WSDisplayRenderSpace) GetClientTaskNamePortClientConnectionIdentifierForTargetIDDisplayUUID(port *uint32, identifier *uint64, id unsafe.Pointer, uuid objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("getClientTaskNamePort:clientConnectionIdentifier:forTargetID:displayUUID:"), port, identifier, id, uuid)
}
func (w WSDisplayRenderSpace) HitTestLayerInformationFromHitTestResults(results unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("hitTestLayerInformationFromHitTestResults:"), results)
	return objectivec.Object{ID: rv}
}
func (w WSDisplayRenderSpace) HostTargetIDForEmbeddedTargetIDDisplayUUID(id unsafe.Pointer, uuid objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("hostTargetIDForEmbeddedTargetID:displayUUID:"), id, uuid)
	return rv
}
func (w WSDisplayRenderSpace) PidForTargetIdentifier(identifier unsafe.Pointer) int32 {
	rv := objc.SendIfResponds[int32](w.ID, objc.Sel("pidForTargetIdentifier:"), identifier)
	return rv
}
func (w WSDisplayRenderSpace) ScaleForDisplay(display objectivec.IObject) float64 {
	rv := objc.SendIfResponds[float64](w.ID, objc.Sel("scaleForDisplay:"), display)
	return rv
}
func (w WSDisplayRenderSpace) TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResults(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("targetIDAtScreenLocation:displayUUID:options:securityAnalysis:results:"), location, uuid, options, objectivec.IObjectSliceToNSArray(analysis), results)
	return rv
}
func (w WSDisplayRenderSpace) TargetIDAtScreenLocationDisplayUUIDOptionsSecurityAnalysisResultsHitTestPassthroughTargetsResult(location corefoundation.CGPoint, uuid objectivec.IObject, options objectivec.IObject, analysis []objectivec.IObject, results unsafe.Pointer, result []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("targetIDAtScreenLocation:displayUUID:options:securityAnalysis:results:hitTestPassthroughTargetsResult:"), location, uuid, options, objectivec.IObjectSliceToNSArray(analysis), results, objectivec.IObjectSliceToNSArray(result))
	return rv
}
func (w WSDisplayRenderSpace) ToDesktopTransformValueForTargetID(id unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("toDesktopTransformValueForTargetID:"), id)
	return objectivec.Object{ID: rv}
}
func (w WSDisplayRenderSpace) TransformFromLayerIdContextIDDisplayUUID(id objectivec.IObject, id2 objectivec.IObject, uuid objectivec.IObject) quartzcore.CATransform3D {
	rv := objc.SendIfResponds[quartzcore.CATransform3D](w.ID, objc.Sel("transformFromLayerId:contextID:displayUUID:"), id, id2, uuid)
	return quartzcore.CATransform3D(rv)
}

func (w WSDisplayRenderSpace) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSDisplayRenderSpace) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSDisplayRenderSpace) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSDisplayRenderSpace) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
