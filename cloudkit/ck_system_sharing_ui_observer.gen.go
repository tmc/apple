// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSystemSharingUIObserver] class.
var (
	_CKSystemSharingUIObserverClass     CKSystemSharingUIObserverClass
	_CKSystemSharingUIObserverClassOnce sync.Once
)

func getCKSystemSharingUIObserverClass() CKSystemSharingUIObserverClass {
	_CKSystemSharingUIObserverClassOnce.Do(func() {
		_CKSystemSharingUIObserverClass = CKSystemSharingUIObserverClass{class: objc.GetClass("CKSystemSharingUIObserver")}
	})
	return _CKSystemSharingUIObserverClass
}

// GetCKSystemSharingUIObserverClass returns the class object for CKSystemSharingUIObserver.
func GetCKSystemSharingUIObserverClass() CKSystemSharingUIObserverClass {
	return getCKSystemSharingUIObserverClass()
}

type CKSystemSharingUIObserverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSystemSharingUIObserverClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSystemSharingUIObserverClass) Alloc() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object the system uses to monitor changes in sharing.
//
// # Overview
//
// Initialize a [CKSystemSharingUIObserver] instance with your [CKContainer]
// when preparing to share an item using the share sheet. Use your
// implementation to update the local state of a shared item when your app
// receives a [CKShare], or to delete a locally cached share when the system
// notifies your app about a share deletion.
//
// The system only propagates changes on the local device using
// [CKSystemSharingUIObserver]. The system doesn’t notify your app about any
// remote changes on the server. For more information about how to keep your
// local cache in sync with remote changes, see [Remote Records].
//
// # Creating a sharing observer
//
//   - [CKSystemSharingUIObserver.InitWithContainer]: Creates and initializes an observer using the provided container.
//
// # Accessing sharing blocks
//
//   - [CKSystemSharingUIObserver.SystemSharingUIDidSaveShareBlock]: A callback block the system invokes after the success or failure of a share save by the system sharing UI.
//   - [CKSystemSharingUIObserver.SetSystemSharingUIDidSaveShareBlock]
//   - [CKSystemSharingUIObserver.SystemSharingUIDidStopSharingBlock]: A callback block the system invokes after the success or failure of a share delete by the system sharing UI.
//   - [CKSystemSharingUIObserver.SetSystemSharingUIDidStopSharingBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver
//
// [Remote Records]: https://developer.apple.com/documentation/CloudKit/remote-records
type CKSystemSharingUIObserver struct {
	objectivec.Object
}

// CKSystemSharingUIObserverFromID constructs a [CKSystemSharingUIObserver] from an objc.ID.
//
// An object the system uses to monitor changes in sharing.
func CKSystemSharingUIObserverFromID(id objc.ID) CKSystemSharingUIObserver {
	return CKSystemSharingUIObserver{objectivec.Object{ID: id}}
}

// NOTE: CKSystemSharingUIObserver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSystemSharingUIObserver] class.
//
// # Creating a sharing observer
//
//   - [ICKSystemSharingUIObserver.InitWithContainer]: Creates and initializes an observer using the provided container.
//
// # Accessing sharing blocks
//
//   - [ICKSystemSharingUIObserver.SystemSharingUIDidSaveShareBlock]: A callback block the system invokes after the success or failure of a share save by the system sharing UI.
//   - [ICKSystemSharingUIObserver.SetSystemSharingUIDidSaveShareBlock]
//   - [ICKSystemSharingUIObserver.SystemSharingUIDidStopSharingBlock]: A callback block the system invokes after the success or failure of a share delete by the system sharing UI.
//   - [ICKSystemSharingUIObserver.SetSystemSharingUIDidStopSharingBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver
type ICKSystemSharingUIObserver interface {
	objectivec.IObject

	// Topic: Creating a sharing observer

	// Creates and initializes an observer using the provided container.
	InitWithContainer(container ICKContainer) CKSystemSharingUIObserver

	// Topic: Accessing sharing blocks

	// A callback block the system invokes after the success or failure of a share save by the system sharing UI.
	SystemSharingUIDidSaveShareBlock() unsafe.Pointer
	SetSystemSharingUIDidSaveShareBlock(value unsafe.Pointer)
	// A callback block the system invokes after the success or failure of a share delete by the system sharing UI.
	SystemSharingUIDidStopSharingBlock() unsafe.Pointer
	SetSystemSharingUIDidStopSharingBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKSystemSharingUIObserver) Init() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSystemSharingUIObserver) Autorelease() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSystemSharingUIObserver creates a new CKSystemSharingUIObserver instance.
func NewCKSystemSharingUIObserver() CKSystemSharingUIObserver {
	class := getCKSystemSharingUIObserverClass()
	rv := objc.Send[CKSystemSharingUIObserver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes an observer using the provided container.
//
// container: The [CKContainer] for the sharing observer.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/init(container:)
func NewCKSystemSharingUIObserverWithContainer(container ICKContainer) CKSystemSharingUIObserver {
	instance := getCKSystemSharingUIObserverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:"), container)
	return CKSystemSharingUIObserverFromID(rv)
}

// Creates and initializes an observer using the provided container.
//
// container: The [CKContainer] for the sharing observer.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/init(container:)
func (c CKSystemSharingUIObserver) InitWithContainer(container ICKContainer) CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](c.ID, objc.Sel("initWithContainer:"), container)
	return rv
}

// A callback block the system invokes after the success or failure of a share
// save by the system sharing UI.
//
// See: https://developer.apple.com/documentation/cloudkit/cksystemsharinguiobserver/systemsharinguididsaveshareblock-8c9vi
func (c CKSystemSharingUIObserver) SystemSharingUIDidSaveShareBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("systemSharingUIDidSaveShareBlock"))
	return rv
}
func (c CKSystemSharingUIObserver) SetSystemSharingUIDidSaveShareBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setSystemSharingUIDidSaveShareBlock:"), value)
}

// A callback block the system invokes after the success or failure of a share
// delete by the system sharing UI.
//
// See: https://developer.apple.com/documentation/cloudkit/cksystemsharinguiobserver/systemsharinguididstopsharingblock-7nmiw
func (c CKSystemSharingUIObserver) SystemSharingUIDidStopSharingBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("systemSharingUIDidStopSharingBlock"))
	return rv
}
func (c CKSystemSharingUIObserver) SetSystemSharingUIDidStopSharingBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setSystemSharingUIDidStopSharingBlock:"), value)
}
