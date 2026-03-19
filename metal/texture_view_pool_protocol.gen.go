// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A pool of lightweight texture views.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool
type MTLTextureViewPool interface {
	objectivec.IObject
	MTLResourceViewPool

	// Creates a new lightweight texture view of a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool/setTextureView(buffer:descriptor:offset:bytesPerRow:index:)
	SetTextureViewFromBufferDescriptorOffsetBytesPerRowAtIndex(buffer MTLBuffer, descriptor IMTLTextureDescriptor, offset uint, bytesPerRow uint, index uint) MTLResourceID

	// Creates a new lightweight texture view.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool/setTextureView(texture:descriptor:index:)
	SetTextureViewDescriptorAtIndex(texture MTLTexture, descriptor IMTLTextureViewDescriptor, index uint) MTLResourceID

	// Copies a default texture view to a slot in this texture view pool at an index provided.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool/setTextureView(texture:index:)
	SetTextureViewAtIndex(texture MTLTexture, index uint) MTLResourceID
}

// MTLTextureViewPoolObject wraps an existing Objective-C object that conforms to the MTLTextureViewPool protocol.
type MTLTextureViewPoolObject struct {
	objectivec.Object
}
func (o MTLTextureViewPoolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTextureViewPoolObjectFromID constructs a [MTLTextureViewPoolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTextureViewPoolObjectFromID(id objc.ID) MTLTextureViewPoolObject {
	return MTLTextureViewPoolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Creates a new lightweight texture view of a buffer.
//
// buffer: An [MTLBuffer] instance for which to create a new texture view.
//
// descriptor: A descriptor specifying properties of the texture view to create.
//
// offset: A byte offset, within the `buffer` parameter, at which the data for the
// texture view starts.
//
// bytesPerRow: The number of bytes between adjacent rows of pixels in the source
// buffer’s memory.
//
// index: An index of a slot in the table into which this method writes the new
// texture view.
//
// # Return Value
// 
// The [MTLResourceID] of a new buffer view in this pool.
//
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// # Discussion
// 
// This method creates a lightweight texture view over a buffer, according to
// a descriptor you provide. It then associates the texture view with a slot
// in this texture view pool at the index you specify.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool/setTextureView(buffer:descriptor:offset:bytesPerRow:index:)
func (o MTLTextureViewPoolObject) SetTextureViewFromBufferDescriptorOffsetBytesPerRowAtIndex(buffer MTLBuffer, descriptor IMTLTextureDescriptor, offset uint, bytesPerRow uint, index uint) MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("setTextureViewFromBuffer:descriptor:offset:bytesPerRow:atIndex:"), buffer, descriptor, offset, bytesPerRow, index)
	return rv
	}
// Creates a new lightweight texture view.
//
// texture: An [MTLTexture] instance for which to create a new lightweight texture
// view.
//
// descriptor: A descriptor specifying properties of the texture view to create.
//
// index: An index of a slot in the texture pool into which this method writes the
// new texture view.
//
// # Return Value
// 
// The [MTLResourceID] of a newly created texture view in this pool.
//
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// # Discussion
// 
// This method creates a lightweight texture view over a texture according to
// a descriptor you provide. It then associates the texture view with a slot
// in this texture view pool at the index you specify.
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool/setTextureView(texture:descriptor:index:)
func (o MTLTextureViewPoolObject) SetTextureViewDescriptorAtIndex(texture MTLTexture, descriptor IMTLTextureViewDescriptor, index uint) MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("setTextureView:descriptor:atIndex:"), texture, descriptor, index)
	return rv
	}
// Copies a default texture view to a slot in this texture view pool at an
// index provided.
//
// texture: An [MTLTexture] instance for which to copy its texture view.
//
// index: An index of a slot in this texture pool into which this method copies the
// texture view.
//
// # Return Value
// 
// The [MTLResourceID] of a newly created texture view in this pool.
//
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// See: https://developer.apple.com/documentation/Metal/MTLTextureViewPool/setTextureView(texture:index:)
func (o MTLTextureViewPoolObject) SetTextureViewAtIndex(texture MTLTexture, index uint) MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("setTextureView:atIndex:"), texture, index)
	return rv
	}
// Obtains the resource ID corresponding to the resource view at index 0 in
// this resource view pool.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/baseResourceID
func (o MTLTextureViewPoolObject) BaseResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("baseResourceID"))
	return rv
	}
// Obtains a reference to the GPU device this pool belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/device
func (o MTLTextureViewPoolObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// Queries the optional debug label of this resource view pool.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/label
func (o MTLTextureViewPoolObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Queries the number of resource views that this pool contains.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/resourceViewCount
func (o MTLTextureViewPoolObject) ResourceViewCount() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("resourceViewCount"))
	return rv
	}
// Copies a range of resource views from a source view pool to a destination
// location in this view pool.
//
// sourcePool: Resource view pool from which to copy resource views.
//
// sourceRange: The range in the source resource view pool to copy.
//
// destinationIndex: The starting index in this destination view pool into which to copy the
// source range of resource views.
//
// # Return Value
// 
// The [MTLResourceID] of the resource view corresponding to
// `destinationIndex` of the copy in this resource view pool.
//
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceViewPool/copyResourceViewsFromPool:sourceRange:destinationIndex:
func (o MTLTextureViewPoolObject) CopyResourceViewsFromPoolSourceRangeDestinationIndex(sourcePool MTLResourceViewPool, sourceRange foundation.NSRange, destinationIndex uint) MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("copyResourceViewsFromPool:sourceRange:destinationIndex:"), sourcePool, sourceRange, destinationIndex)
	return rv
	}

