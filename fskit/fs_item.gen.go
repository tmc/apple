// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSItem] class.
var (
	_FSItemClass     FSItemClass
	_FSItemClassOnce sync.Once
)

func getFSItemClass() FSItemClass {
	_FSItemClassOnce.Do(func() {
		_FSItemClass = FSItemClass{class: objc.GetClass("FSItem")}
	})
	return _FSItemClass
}

// GetFSItemClass returns the class object for FSItem.
func GetFSItemClass() FSItemClass {
	return getFSItemClass()
}

type FSItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSItemClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSItemClass) Alloc() FSItem {
	rv := objc.Send[FSItem](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A distinct object in a file hierarchy, such as a file, directory, symlink,
// socket, and more.
//
// # Overview
//
// An [FSItem] is a mostly opaque object, which your file system
// implementation defines as needed.
//
// The [FSItemAttributes] class defines nonatomic properties to support
// [FSItem] instances. An [FSItemAttributes] instance contains a snapshot of
// the attributes of an [FSItem] at one point in time. The [FSItemAttributes]
// properties have no explicit thread safety provisions, since the operations
// that either get or set these properties enforce thread safety.
//
// You test an attribute’s validity with the the method
// [FSItemAttributes.IsValid]. If the value is `true` (Swift) or [YES]
// (Objective-C), it’s safe to use the attribute.
//
// Methods that get or set an item’s attribute use
// [FSItemGetAttributesRequest] or [FSItemSetAttributesRequest], respectively.
// Both are subclasses of [FSItemAttributes]. An [FSItemGetAttributesRequest]
// contains a [FSItemGetAttributesRequest.WantedAttributes] property to
// indicate the attributes a file system provides for the request. Similarly,
// [FSItemSetAttributesRequest] uses the property
// [FSItemSetAttributesRequest.ConsumedAttributes] for a file system to signal
// back which attributes it successfully used.
//
// [FSItem] is the FSKit equivelant of a vnode in the kernel. For every FSKit
// vnode in the kernel, the [FSModule] hosting the volume has an instantiated
// [FSItem].
//
// See: https://developer.apple.com/documentation/FSKit/FSItem
type FSItem struct {
	objectivec.Object
}

// FSItemFromID constructs a [FSItem] from an objc.ID.
//
// A distinct object in a file hierarchy, such as a file, directory, symlink,
// socket, and more.
func FSItemFromID(id objc.ID) FSItem {
	return FSItem{objectivec.Object{ID: id}}
}

// NOTE: FSItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSItem] class.
//
// See: https://developer.apple.com/documentation/FSKit/FSItem
type IFSItem interface {
	objectivec.IObject

	// The attributes successfully used by the file system.
	ConsumedAttributes() unsafe.Pointer
	SetConsumedAttributes(value unsafe.Pointer)
	// The attributes requested by the request.
	WantedAttributes() unsafe.Pointer
	SetWantedAttributes(value unsafe.Pointer)
}

// Init initializes the instance.
func (i FSItem) Init() FSItem {
	rv := objc.Send[FSItem](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i FSItem) Autorelease() FSItem {
	rv := objc.Send[FSItem](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItem creates a new FSItem instance.
func NewFSItem() FSItem {
	class := getFSItemClass()
	rv := objc.Send[FSItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The attributes successfully used by the file system.
//
// See: https://developer.apple.com/documentation/fskit/fsitem/setattributesrequest/consumedattributes
func (i FSItem) ConsumedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("consumedAttributes"))
	return rv
}
func (i FSItem) SetConsumedAttributes(value unsafe.Pointer) {
	objc.Send[struct{}](i.ID, objc.Sel("setConsumedAttributes:"), value)
}

// The attributes requested by the request.
//
// See: https://developer.apple.com/documentation/fskit/fsitem/getattributesrequest/wantedattributes
func (i FSItem) WantedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("wantedAttributes"))
	return rv
}
func (i FSItem) SetWantedAttributes(value unsafe.Pointer) {
	objc.Send[struct{}](i.ID, objc.Sel("setWantedAttributes:"), value)
}
