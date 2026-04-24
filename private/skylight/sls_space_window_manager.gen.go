// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSSpaceWindowManager] class.
var (
	_SLSSpaceWindowManagerClass     SLSSpaceWindowManagerClass
	_SLSSpaceWindowManagerClassOnce sync.Once
)

func getSLSSpaceWindowManagerClass() SLSSpaceWindowManagerClass {
	_SLSSpaceWindowManagerClassOnce.Do(func() {
		_SLSSpaceWindowManagerClass = SLSSpaceWindowManagerClass{class: objc.GetClass("SLSSpaceWindowManager")}
	})
	return _SLSSpaceWindowManagerClass
}

// GetSLSSpaceWindowManagerClass returns the class object for SLSSpaceWindowManager.
func GetSLSSpaceWindowManagerClass() SLSSpaceWindowManagerClass {
	return getSLSSpaceWindowManagerClass()
}

type SLSSpaceWindowManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSSpaceWindowManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSSpaceWindowManagerClass) Alloc() SLSSpaceWindowManager {
	rv := objc.Send[SLSSpaceWindowManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSSpaceWindowManager._beginBatch]
//   - [SLSSpaceWindowManager._checkDisplayState]
//   - [SLSSpaceWindowManager._checkSpaceMovedToDisplayDisplayUUID]
//   - [SLSSpaceWindowManager._endBatch]
//   - [SLSSpaceWindowManager._fullRebuildSpaceChange]
//   - [SLSSpaceWindowManager._fullRebuildSpacesChanged]
//   - [SLSSpaceWindowManager._getDisplayUUIDForSpace]
//   - [SLSSpaceWindowManager._performBatchingCallouts]
//   - [SLSSpaceWindowManager._postActiveDisplayChange]
//   - [SLSSpaceWindowManager._removeSpace]
//   - [SLSSpaceWindowManager._spaceAddWindow]
//   - [SLSSpaceWindowManager._spaceRemoveWindow]
//   - [SLSSpaceWindowManager._spaceChangedDisplay]
//   - [SLSSpaceWindowManager._spaceWithIDCreateIfNeeded]
//   - [SLSSpaceWindowManager._updateSpaceWithData]
//   - [SLSSpaceWindowManager.Activate]
//   - [SLSSpaceWindowManager.AddWindowsToSpacesRemovingFromTransaction]
//   - [SLSSpaceWindowManager.BatchedDelegate]
//   - [SLSSpaceWindowManager.BeganBatch]
//   - [SLSSpaceWindowManager.SetBeganBatch]
//   - [SLSSpaceWindowManager.Capabilities]
//   - [SLSSpaceWindowManager.SetCapabilities]
//   - [SLSSpaceWindowManager.ConnectionID]
//   - [SLSSpaceWindowManager.SetConnectionID]
//   - [SLSSpaceWindowManager.Delegate]
//   - [SLSSpaceWindowManager.SetDelegate]
//   - [SLSSpaceWindowManager.DisplayCurrentSpaces]
//   - [SLSSpaceWindowManager.SetDisplayCurrentSpaces]
//   - [SLSSpaceWindowManager.DisplaySpaceList]
//   - [SLSSpaceWindowManager.SetDisplaySpaceList]
//   - [SLSSpaceWindowManager.DisplaysHaveSeparateSpaces]
//   - [SLSSpaceWindowManager.SetDisplaysHaveSeparateSpaces]
//   - [SLSSpaceWindowManager.Invalidate]
//   - [SLSSpaceWindowManager.IsWindowPresentOnUnmanagedSpaces]
//   - [SLSSpaceWindowManager.MoveDraggedWindowToPointMouseLocationTimestampTransaction]
//   - [SLSSpaceWindowManager.NestedCalloutCount]
//   - [SLSSpaceWindowManager.SetNestedCalloutCount]
//   - [SLSSpaceWindowManager.RebuildMenuBarOnSpaceFrontConnectionTransaction]
//   - [SLSSpaceWindowManager.RequestSpaceSwitchForWindowTransaction]
//   - [SLSSpaceWindowManager.SetGlobalWindowVisibilityListTransaction]
//   - [SLSSpaceWindowManager.SpaceWithID]
//   - [SLSSpaceWindowManager.Spaces]
//   - [SLSSpaceWindowManager.SetSpaces]
//   - [SLSSpaceWindowManager.Synchronize]
//   - [SLSSpaceWindowManager.Valid]
//   - [SLSSpaceWindowManager.SetValid]
//   - [SLSSpaceWindowManager.InitWithConnectionIDDelegateCapabilities]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager
type SLSSpaceWindowManager struct {
	objectivec.Object
}

// SLSSpaceWindowManagerFromID constructs a [SLSSpaceWindowManager] from an objc.ID.
func SLSSpaceWindowManagerFromID(id objc.ID) SLSSpaceWindowManager {
	return SLSSpaceWindowManager{objectivec.Object{ID: id}}
}

// Ensure SLSSpaceWindowManager implements ISLSSpaceWindowManager.
var _ ISLSSpaceWindowManager = SLSSpaceWindowManager{}

// An interface definition for the [SLSSpaceWindowManager] class.
//
// # Methods
//
//   - [ISLSSpaceWindowManager._beginBatch]
//   - [ISLSSpaceWindowManager._checkDisplayState]
//   - [ISLSSpaceWindowManager._checkSpaceMovedToDisplayDisplayUUID]
//   - [ISLSSpaceWindowManager._endBatch]
//   - [ISLSSpaceWindowManager._fullRebuildSpaceChange]
//   - [ISLSSpaceWindowManager._fullRebuildSpacesChanged]
//   - [ISLSSpaceWindowManager._getDisplayUUIDForSpace]
//   - [ISLSSpaceWindowManager._performBatchingCallouts]
//   - [ISLSSpaceWindowManager._postActiveDisplayChange]
//   - [ISLSSpaceWindowManager._removeSpace]
//   - [ISLSSpaceWindowManager._spaceAddWindow]
//   - [ISLSSpaceWindowManager._spaceRemoveWindow]
//   - [ISLSSpaceWindowManager._spaceChangedDisplay]
//   - [ISLSSpaceWindowManager._spaceWithIDCreateIfNeeded]
//   - [ISLSSpaceWindowManager._updateSpaceWithData]
//   - [ISLSSpaceWindowManager.Activate]
//   - [ISLSSpaceWindowManager.AddWindowsToSpacesRemovingFromTransaction]
//   - [ISLSSpaceWindowManager.BatchedDelegate]
//   - [ISLSSpaceWindowManager.BeganBatch]
//   - [ISLSSpaceWindowManager.SetBeganBatch]
//   - [ISLSSpaceWindowManager.Capabilities]
//   - [ISLSSpaceWindowManager.SetCapabilities]
//   - [ISLSSpaceWindowManager.ConnectionID]
//   - [ISLSSpaceWindowManager.SetConnectionID]
//   - [ISLSSpaceWindowManager.Delegate]
//   - [ISLSSpaceWindowManager.SetDelegate]
//   - [ISLSSpaceWindowManager.DisplayCurrentSpaces]
//   - [ISLSSpaceWindowManager.SetDisplayCurrentSpaces]
//   - [ISLSSpaceWindowManager.DisplaySpaceList]
//   - [ISLSSpaceWindowManager.SetDisplaySpaceList]
//   - [ISLSSpaceWindowManager.DisplaysHaveSeparateSpaces]
//   - [ISLSSpaceWindowManager.SetDisplaysHaveSeparateSpaces]
//   - [ISLSSpaceWindowManager.Invalidate]
//   - [ISLSSpaceWindowManager.IsWindowPresentOnUnmanagedSpaces]
//   - [ISLSSpaceWindowManager.MoveDraggedWindowToPointMouseLocationTimestampTransaction]
//   - [ISLSSpaceWindowManager.NestedCalloutCount]
//   - [ISLSSpaceWindowManager.SetNestedCalloutCount]
//   - [ISLSSpaceWindowManager.RebuildMenuBarOnSpaceFrontConnectionTransaction]
//   - [ISLSSpaceWindowManager.RequestSpaceSwitchForWindowTransaction]
//   - [ISLSSpaceWindowManager.SetGlobalWindowVisibilityListTransaction]
//   - [ISLSSpaceWindowManager.SpaceWithID]
//   - [ISLSSpaceWindowManager.Spaces]
//   - [ISLSSpaceWindowManager.SetSpaces]
//   - [ISLSSpaceWindowManager.Synchronize]
//   - [ISLSSpaceWindowManager.Valid]
//   - [ISLSSpaceWindowManager.SetValid]
//   - [ISLSSpaceWindowManager.InitWithConnectionIDDelegateCapabilities]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager
type ISLSSpaceWindowManager interface {
	objectivec.IObject

	// Topic: Methods

	_beginBatch()
	_checkDisplayState(state objectivec.IObject)
	_checkSpaceMovedToDisplayDisplayUUID(display objectivec.IObject, uuid objectivec.IObject)
	_endBatch()
	_fullRebuildSpaceChange(change uint64)
	_fullRebuildSpacesChanged()
	_getDisplayUUIDForSpace(space uint64) objectivec.IObject
	_performBatchingCallouts(callouts VoidHandler)
	_postActiveDisplayChange()
	_removeSpace(space uint64)
	_spaceAddWindow(_space uint64, window uint32)
	_spaceRemoveWindow(_space uint64, window uint32)
	_spaceChangedDisplay(display uint64)
	_spaceWithIDCreateIfNeeded(id uint64, needed bool) objectivec.IObject
	_updateSpaceWithData(data objectivec.IObject)
	Activate()
	AddWindowsToSpacesRemovingFromTransaction(windows objectivec.IObject, spaces objectivec.IObject, from uint32, transaction SLSTransactionRef)
	BatchedDelegate() objectivec.IObject
	BeganBatch() bool
	SetBeganBatch(value bool)
	Capabilities() uint64
	SetCapabilities(value uint64)
	ConnectionID() uint32
	SetConnectionID(value uint32)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	DisplayCurrentSpaces() foundation.INSDictionary
	SetDisplayCurrentSpaces(value foundation.INSDictionary)
	DisplaySpaceList() foundation.INSDictionary
	SetDisplaySpaceList(value foundation.INSDictionary)
	DisplaysHaveSeparateSpaces() bool
	SetDisplaysHaveSeparateSpaces(value bool)
	Invalidate()
	IsWindowPresentOnUnmanagedSpaces(spaces uint32) bool
	MoveDraggedWindowToPointMouseLocationTimestampTransaction(window uint32, point corefoundation.CGPoint, location corefoundation.CGPoint, timestamp uint64, transaction SLSTransactionRef)
	NestedCalloutCount() int64
	SetNestedCalloutCount(value int64)
	RebuildMenuBarOnSpaceFrontConnectionTransaction(space uint64, connection uint32, transaction SLSTransactionRef)
	RequestSpaceSwitchForWindowTransaction(window uint32, transaction SLSTransactionRef)
	SetGlobalWindowVisibilityListTransaction(list objectivec.IObject, transaction SLSTransactionRef)
	SpaceWithID(id uint64) objectivec.IObject
	Spaces() foundation.INSDictionary
	SetSpaces(value foundation.INSDictionary)
	Synchronize()
	Valid() bool
	SetValid(value bool)
	InitWithConnectionIDDelegateCapabilities(id uint32, delegate objectivec.IObject, capabilities uint64) SLSSpaceWindowManager
}

// Init initializes the instance.
func (s SLSSpaceWindowManager) Init() SLSSpaceWindowManager {
	rv := objc.Send[SLSSpaceWindowManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSSpaceWindowManager) Autorelease() SLSSpaceWindowManager {
	rv := objc.Send[SLSSpaceWindowManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSSpaceWindowManager creates a new SLSSpaceWindowManager instance.
func NewSLSSpaceWindowManager() SLSSpaceWindowManager {
	class := getSLSSpaceWindowManagerClass()
	rv := objc.Send[SLSSpaceWindowManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/initWithConnectionID:delegate:capabilities:
func NewSLSSpaceWindowManagerWithConnectionIDDelegateCapabilities(id uint32, delegate objectivec.IObject, capabilities uint64) SLSSpaceWindowManager {
	instance := getSLSSpaceWindowManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConnectionID:delegate:capabilities:"), id, delegate, capabilities)
	return SLSSpaceWindowManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_beginBatch
func (s SLSSpaceWindowManager) _beginBatch() {
	objc.Send[objc.ID](s.ID, objc.Sel("_beginBatch"))
}

// BeginBatch is an exported wrapper for the private method _beginBatch.
func (s SLSSpaceWindowManager) BeginBatch() {
	s._beginBatch()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_checkDisplayState:
func (s SLSSpaceWindowManager) _checkDisplayState(state objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_checkDisplayState:"), state)
}

// CheckDisplayState is an exported wrapper for the private method _checkDisplayState.
func (s SLSSpaceWindowManager) CheckDisplayState(state objectivec.IObject) {
	s._checkDisplayState(state)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_checkSpaceMovedToDisplay:displayUUID:
func (s SLSSpaceWindowManager) _checkSpaceMovedToDisplayDisplayUUID(display objectivec.IObject, uuid objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_checkSpaceMovedToDisplay:displayUUID:"), display, uuid)
}

// CheckSpaceMovedToDisplayDisplayUUID is an exported wrapper for the private method _checkSpaceMovedToDisplayDisplayUUID.
func (s SLSSpaceWindowManager) CheckSpaceMovedToDisplayDisplayUUID(display objectivec.IObject, uuid objectivec.IObject) {
	s._checkSpaceMovedToDisplayDisplayUUID(display, uuid)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_endBatch
func (s SLSSpaceWindowManager) _endBatch() {
	objc.Send[objc.ID](s.ID, objc.Sel("_endBatch"))
}

// EndBatch is an exported wrapper for the private method _endBatch.
func (s SLSSpaceWindowManager) EndBatch() {
	s._endBatch()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_fullRebuildSpaceChange:
func (s SLSSpaceWindowManager) _fullRebuildSpaceChange(change uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("_fullRebuildSpaceChange:"), change)
}

// FullRebuildSpaceChange is an exported wrapper for the private method _fullRebuildSpaceChange.
func (s SLSSpaceWindowManager) FullRebuildSpaceChange(change uint64) {
	s._fullRebuildSpaceChange(change)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_fullRebuildSpacesChanged
func (s SLSSpaceWindowManager) _fullRebuildSpacesChanged() {
	objc.Send[objc.ID](s.ID, objc.Sel("_fullRebuildSpacesChanged"))
}

// FullRebuildSpacesChanged is an exported wrapper for the private method _fullRebuildSpacesChanged.
func (s SLSSpaceWindowManager) FullRebuildSpacesChanged() {
	s._fullRebuildSpacesChanged()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_getDisplayUUIDForSpace:
func (s SLSSpaceWindowManager) _getDisplayUUIDForSpace(space uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_getDisplayUUIDForSpace:"), space)
	return objectivec.Object{ID: rv}
}

// GetDisplayUUIDForSpace is an exported wrapper for the private method _getDisplayUUIDForSpace.
func (s SLSSpaceWindowManager) GetDisplayUUIDForSpace(space uint64) objectivec.IObject {
	return s._getDisplayUUIDForSpace(space)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_performBatchingCallouts:
func (s SLSSpaceWindowManager) _performBatchingCallouts(callouts VoidHandler) {
	_block0, _ := NewVoidBlock(callouts)
	objc.Send[objc.ID](s.ID, objc.Sel("_performBatchingCallouts:"), _block0)
}

// PerformBatchingCallouts is an exported wrapper for the private method _performBatchingCallouts.
func (s SLSSpaceWindowManager) PerformBatchingCallouts(callouts VoidHandler) {
	s._performBatchingCallouts(callouts)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_postActiveDisplayChange
func (s SLSSpaceWindowManager) _postActiveDisplayChange() {
	objc.Send[objc.ID](s.ID, objc.Sel("_postActiveDisplayChange"))
}

// PostActiveDisplayChange is an exported wrapper for the private method _postActiveDisplayChange.
func (s SLSSpaceWindowManager) PostActiveDisplayChange() {
	s._postActiveDisplayChange()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_removeSpace:
func (s SLSSpaceWindowManager) _removeSpace(space uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("_removeSpace:"), space)
}

// RemoveSpace is an exported wrapper for the private method _removeSpace.
func (s SLSSpaceWindowManager) RemoveSpace(space uint64) {
	s._removeSpace(space)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_space:addWindow:
func (s SLSSpaceWindowManager) _spaceAddWindow(_space uint64, window uint32) {
	objc.Send[objc.ID](s.ID, objc.Sel("_space:addWindow:"), _space, window)
}

// SpaceAddWindow is an exported wrapper for the private method _spaceAddWindow.
func (s SLSSpaceWindowManager) SpaceAddWindow(_space uint64, window uint32) {
	s._spaceAddWindow(_space, window)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_space:removeWindow:
func (s SLSSpaceWindowManager) _spaceRemoveWindow(_space uint64, window uint32) {
	objc.Send[objc.ID](s.ID, objc.Sel("_space:removeWindow:"), _space, window)
}

// SpaceRemoveWindow is an exported wrapper for the private method _spaceRemoveWindow.
func (s SLSSpaceWindowManager) SpaceRemoveWindow(_space uint64, window uint32) {
	s._spaceRemoveWindow(_space, window)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_spaceChangedDisplay:
func (s SLSSpaceWindowManager) _spaceChangedDisplay(display uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("_spaceChangedDisplay:"), display)
}

// SpaceChangedDisplay is an exported wrapper for the private method _spaceChangedDisplay.
func (s SLSSpaceWindowManager) SpaceChangedDisplay(display uint64) {
	s._spaceChangedDisplay(display)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_spaceWithID:createIfNeeded:
func (s SLSSpaceWindowManager) _spaceWithIDCreateIfNeeded(id uint64, needed bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_spaceWithID:createIfNeeded:"), id, needed)
	return objectivec.Object{ID: rv}
}

// SpaceWithIDCreateIfNeeded is an exported wrapper for the private method _spaceWithIDCreateIfNeeded.
func (s SLSSpaceWindowManager) SpaceWithIDCreateIfNeeded(id uint64, needed bool) objectivec.IObject {
	return s._spaceWithIDCreateIfNeeded(id, needed)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/_updateSpaceWithData:
func (s SLSSpaceWindowManager) _updateSpaceWithData(data objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_updateSpaceWithData:"), data)
}

// UpdateSpaceWithData is an exported wrapper for the private method _updateSpaceWithData.
func (s SLSSpaceWindowManager) UpdateSpaceWithData(data objectivec.IObject) {
	s._updateSpaceWithData(data)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/activate
func (s SLSSpaceWindowManager) Activate() {
	objc.Send[objc.ID](s.ID, objc.Sel("activate"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/addWindows:toSpaces:removingFrom:transaction:
func (s SLSSpaceWindowManager) AddWindowsToSpacesRemovingFromTransaction(windows objectivec.IObject, spaces objectivec.IObject, from uint32, transaction SLSTransactionRef) {
	objc.Send[objc.ID](s.ID, objc.Sel("addWindows:toSpaces:removingFrom:transaction:"), windows, spaces, from, transaction)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/batchedDelegate
func (s SLSSpaceWindowManager) BatchedDelegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("batchedDelegate"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/invalidate
func (s SLSSpaceWindowManager) Invalidate() {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidate"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/isWindowPresentOnUnmanagedSpaces:
func (s SLSSpaceWindowManager) IsWindowPresentOnUnmanagedSpaces(spaces uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isWindowPresentOnUnmanagedSpaces:"), spaces)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/moveDraggedWindow:toPoint:mouseLocation:timestamp:transaction:
func (s SLSSpaceWindowManager) MoveDraggedWindowToPointMouseLocationTimestampTransaction(window uint32, point corefoundation.CGPoint, location corefoundation.CGPoint, timestamp uint64, transaction SLSTransactionRef) {
	objc.Send[objc.ID](s.ID, objc.Sel("moveDraggedWindow:toPoint:mouseLocation:timestamp:transaction:"), window, point, location, timestamp, transaction)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/rebuildMenuBarOnSpace:frontConnection:transaction:
func (s SLSSpaceWindowManager) RebuildMenuBarOnSpaceFrontConnectionTransaction(space uint64, connection uint32, transaction SLSTransactionRef) {
	objc.Send[objc.ID](s.ID, objc.Sel("rebuildMenuBarOnSpace:frontConnection:transaction:"), space, connection, transaction)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/requestSpaceSwitchForWindow:transaction:
func (s SLSSpaceWindowManager) RequestSpaceSwitchForWindowTransaction(window uint32, transaction SLSTransactionRef) {
	objc.Send[objc.ID](s.ID, objc.Sel("requestSpaceSwitchForWindow:transaction:"), window, transaction)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/setGlobalWindowVisibilityList:transaction:
func (s SLSSpaceWindowManager) SetGlobalWindowVisibilityListTransaction(list objectivec.IObject, transaction SLSTransactionRef) {
	objc.Send[objc.ID](s.ID, objc.Sel("setGlobalWindowVisibilityList:transaction:"), list, transaction)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/spaceWithID:
func (s SLSSpaceWindowManager) SpaceWithID(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spaceWithID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/synchronize
func (s SLSSpaceWindowManager) Synchronize() {
	objc.Send[objc.ID](s.ID, objc.Sel("synchronize"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/initWithConnectionID:delegate:capabilities:
func (s SLSSpaceWindowManager) InitWithConnectionIDDelegateCapabilities(id uint32, delegate objectivec.IObject, capabilities uint64) SLSSpaceWindowManager {
	rv := objc.Send[SLSSpaceWindowManager](s.ID, objc.Sel("initWithConnectionID:delegate:capabilities:"), id, delegate, capabilities)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/beganBatch
func (s SLSSpaceWindowManager) BeganBatch() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("beganBatch"))
	return rv
}
func (s SLSSpaceWindowManager) SetBeganBatch(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setBeganBatch:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/capabilities
func (s SLSSpaceWindowManager) Capabilities() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("capabilities"))
	return rv
}
func (s SLSSpaceWindowManager) SetCapabilities(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setCapabilities:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/connectionID
func (s SLSSpaceWindowManager) ConnectionID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("connectionID"))
	return rv
}
func (s SLSSpaceWindowManager) SetConnectionID(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setConnectionID:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/delegate
func (s SLSSpaceWindowManager) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (s SLSSpaceWindowManager) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/displayCurrentSpaces
func (s SLSSpaceWindowManager) DisplayCurrentSpaces() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayCurrentSpaces"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLSSpaceWindowManager) SetDisplayCurrentSpaces(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplayCurrentSpaces:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/displaySpaceList
func (s SLSSpaceWindowManager) DisplaySpaceList() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displaySpaceList"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLSSpaceWindowManager) SetDisplaySpaceList(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplaySpaceList:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/displaysHaveSeparateSpaces
func (s SLSSpaceWindowManager) DisplaysHaveSeparateSpaces() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("displaysHaveSeparateSpaces"))
	return rv
}
func (s SLSSpaceWindowManager) SetDisplaysHaveSeparateSpaces(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplaysHaveSeparateSpaces:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/nestedCalloutCount
func (s SLSSpaceWindowManager) NestedCalloutCount() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("nestedCalloutCount"))
	return rv
}
func (s SLSSpaceWindowManager) SetNestedCalloutCount(value int64) {
	objc.Send[struct{}](s.ID, objc.Sel("setNestedCalloutCount:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/spaces
func (s SLSSpaceWindowManager) Spaces() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SLSSpaceWindowManager) SetSpaces(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpaces:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSSpaceWindowManager/valid
func (s SLSSpaceWindowManager) Valid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("valid"))
	return rv
}
func (s SLSSpaceWindowManager) SetValid(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setValid:"), value)
}

// _performBatchingCalloutsSync is a synchronous wrapper around [SLSSpaceWindowManager._performBatchingCallouts].
// It blocks until the completion handler fires or the context is cancelled.
func (s SLSSpaceWindowManager) _performBatchingCalloutsSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	s._performBatchingCallouts(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
