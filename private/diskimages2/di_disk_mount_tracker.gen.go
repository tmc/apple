// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIDiskMountTracker] class.
var (
	_DIDiskMountTrackerClass     DIDiskMountTrackerClass
	_DIDiskMountTrackerClassOnce sync.Once
)

func getDIDiskMountTrackerClass() DIDiskMountTrackerClass {
	_DIDiskMountTrackerClassOnce.Do(func() {
		_DIDiskMountTrackerClass = DIDiskMountTrackerClass{class: objc.GetClass("DIDiskMountTracker")}
	})
	return _DIDiskMountTrackerClass
}

// GetDIDiskMountTrackerClass returns the class object for DIDiskMountTracker.
func GetDIDiskMountTrackerClass() DIDiskMountTrackerClass {
	return getDIDiskMountTrackerClass()
}

type DIDiskMountTrackerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIDiskMountTrackerClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIDiskMountTrackerClass) Alloc() DIDiskMountTracker {
	rv := objc.SendIfResponds[DIDiskMountTracker](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIDiskMountTracker.AppearedDiskCount]
//   - [DIDiskMountTracker.SetAppearedDiskCount]
//   - [DIDiskMountTracker.AutoMount]
//   - [DIDiskMountTracker.CountIOMediaObjectsWithError]
//   - [DIDiskMountTracker.DiskArb]
//   - [DIDiskMountTracker.MountableDiskCount]
//   - [DIDiskMountTracker.SetMountableDiskCount]
//   - [DIDiskMountTracker.MountedDiskCount]
//   - [DIDiskMountTracker.SetMountedDiskCount]
//   - [DIDiskMountTracker.MyDisks]
//   - [DIDiskMountTracker.SetMyDisks]
//   - [DIDiskMountTracker.OnDiskAppearedWithDiskNameIsMountable]
//   - [DIDiskMountTracker.OnDiskDescriptionChangedWithDiskNameIsMounted]
//   - [DIDiskMountTracker.RegEntryID]
//   - [DIDiskMountTracker.TrackIfMyDisk]
//   - [DIDiskMountTracker.WaitForDAMountWithError]
//   - [DIDiskMountTracker.InitWithRegEntryIDAutoMountError]
type DIDiskMountTracker struct {
	objectivec.Object
}

// DIDiskMountTrackerFromID constructs a [DIDiskMountTracker] from an objc.ID.
func DIDiskMountTrackerFromID(id objc.ID) DIDiskMountTracker {
	return DIDiskMountTracker{objectivec.Object{ID: id}}
}

// Ensure DIDiskMountTracker implements IDIDiskMountTracker.
var _ IDIDiskMountTracker = DIDiskMountTracker{}

// An interface definition for the [DIDiskMountTracker] class.
//
// # Methods
//
//   - [IDIDiskMountTracker.AppearedDiskCount]
//   - [IDIDiskMountTracker.SetAppearedDiskCount]
//   - [IDIDiskMountTracker.AutoMount]
//   - [IDIDiskMountTracker.CountIOMediaObjectsWithError]
//   - [IDIDiskMountTracker.DiskArb]
//   - [IDIDiskMountTracker.MountableDiskCount]
//   - [IDIDiskMountTracker.SetMountableDiskCount]
//   - [IDIDiskMountTracker.MountedDiskCount]
//   - [IDIDiskMountTracker.SetMountedDiskCount]
//   - [IDIDiskMountTracker.MyDisks]
//   - [IDIDiskMountTracker.SetMyDisks]
//   - [IDIDiskMountTracker.OnDiskAppearedWithDiskNameIsMountable]
//   - [IDIDiskMountTracker.OnDiskDescriptionChangedWithDiskNameIsMounted]
//   - [IDIDiskMountTracker.RegEntryID]
//   - [IDIDiskMountTracker.TrackIfMyDisk]
//   - [IDIDiskMountTracker.WaitForDAMountWithError]
//   - [IDIDiskMountTracker.InitWithRegEntryIDAutoMountError]
type IDIDiskMountTracker interface {
	objectivec.IObject

	// Topic: Methods

	AppearedDiskCount() int64
	SetAppearedDiskCount(value int64)
	AutoMount() bool
	CountIOMediaObjectsWithError() (objectivec.IObject, error)
	DiskArb() IDIDiskArb
	MountableDiskCount() int64
	SetMountableDiskCount(value int64)
	MountedDiskCount() int64
	SetMountedDiskCount(value int64)
	MyDisks() foundation.INSSet
	SetMyDisks(value foundation.INSSet)
	OnDiskAppearedWithDiskNameIsMountable(name objectivec.IObject, mountable bool)
	OnDiskDescriptionChangedWithDiskNameIsMounted(name objectivec.IObject, mounted bool)
	RegEntryID() uint64
	TrackIfMyDisk(disk objectivec.IObject) bool
	WaitForDAMountWithError() (bool, error)
	InitWithRegEntryIDAutoMountError(id uint64, mount bool) (DIDiskMountTracker, error)
}

// Init initializes the instance.
func (d DIDiskMountTracker) Init() DIDiskMountTracker {
	rv := objc.SendIfResponds[DIDiskMountTracker](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDiskMountTracker) Autorelease() DIDiskMountTracker {
	rv := objc.SendIfResponds[DIDiskMountTracker](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDiskMountTracker creates a new DIDiskMountTracker instance.
func NewDIDiskMountTracker() DIDiskMountTracker {
	class := getDIDiskMountTrackerClass()
	rv := objc.SendIfResponds[DIDiskMountTracker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDIDiskMountTrackerWithRegEntryIDAutoMountError(id uint64, mount bool) (DIDiskMountTracker, error) {
	var errorPtr objc.ID
	instance := getDIDiskMountTrackerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRegEntryID:autoMount:error:"), id, mount, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIDiskMountTracker{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return DIDiskMountTracker{}, objc.ErrInitFailed
	}
	return DIDiskMountTrackerFromID(rv), nil
}

func (d DIDiskMountTracker) CountIOMediaObjectsWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("countIOMediaObjectsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (d DIDiskMountTracker) OnDiskAppearedWithDiskNameIsMountable(name objectivec.IObject, mountable bool) {
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("onDiskAppearedWithDiskName:isMountable:"), name, mountable)
}
func (d DIDiskMountTracker) OnDiskDescriptionChangedWithDiskNameIsMounted(name objectivec.IObject, mounted bool) {
	objc.SendIfResponds[objc.ID](d.ID, objc.Sel("onDiskDescriptionChangedWithDiskName:isMounted:"), name, mounted)
}
func (d DIDiskMountTracker) TrackIfMyDisk(disk objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("trackIfMyDisk:"), disk)
	return rv
}
func (d DIDiskMountTracker) WaitForDAMountWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("waitForDAMountWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("waitForDAMountWithError: returned NO with nil NSError")
	}
	return rv, nil

}
func (d DIDiskMountTracker) InitWithRegEntryIDAutoMountError(id uint64, mount bool) (DIDiskMountTracker, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithRegEntryID:autoMount:error:"), id, mount, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIDiskMountTracker{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIDiskMountTrackerFromID(rv), nil

}

func (d DIDiskMountTracker) AppearedDiskCount() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("appearedDiskCount"))
	return rv
}
func (d DIDiskMountTracker) SetAppearedDiskCount(value int64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setAppearedDiskCount:"), value)
}
func (d DIDiskMountTracker) AutoMount() bool {
	rv := objc.SendIfResponds[bool](d.ID, objc.Sel("autoMount"))
	return rv
}
func (d DIDiskMountTracker) DiskArb() IDIDiskArb {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("diskArb"))
	return DIDiskArbFromID(objc.ID(rv))
}
func (d DIDiskMountTracker) MountableDiskCount() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("mountableDiskCount"))
	return rv
}
func (d DIDiskMountTracker) SetMountableDiskCount(value int64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setMountableDiskCount:"), value)
}
func (d DIDiskMountTracker) MountedDiskCount() int64 {
	rv := objc.SendIfResponds[int64](d.ID, objc.Sel("mountedDiskCount"))
	return rv
}
func (d DIDiskMountTracker) SetMountedDiskCount(value int64) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setMountedDiskCount:"), value)
}
func (d DIDiskMountTracker) MyDisks() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](d.ID, objc.Sel("myDisks"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (d DIDiskMountTracker) SetMyDisks(value foundation.INSSet) {
	objc.SendIfResponds[struct{}](d.ID, objc.Sel("setMyDisks:"), value)
}
func (d DIDiskMountTracker) RegEntryID() uint64 {
	rv := objc.SendIfResponds[uint64](d.ID, objc.Sel("regEntryID"))
	return rv
}
