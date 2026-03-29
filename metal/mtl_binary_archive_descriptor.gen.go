// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLBinaryArchiveDescriptor] class.
var (
	_MTLBinaryArchiveDescriptorClass     MTLBinaryArchiveDescriptorClass
	_MTLBinaryArchiveDescriptorClassOnce sync.Once
)

func getMTLBinaryArchiveDescriptorClass() MTLBinaryArchiveDescriptorClass {
	_MTLBinaryArchiveDescriptorClassOnce.Do(func() {
		_MTLBinaryArchiveDescriptorClass = MTLBinaryArchiveDescriptorClass{class: objc.GetClass("MTLBinaryArchiveDescriptor")}
	})
	return _MTLBinaryArchiveDescriptorClass
}

// GetMTLBinaryArchiveDescriptorClass returns the class object for MTLBinaryArchiveDescriptor.
func GetMTLBinaryArchiveDescriptorClass() MTLBinaryArchiveDescriptorClass {
	return getMTLBinaryArchiveDescriptorClass()
}

type MTLBinaryArchiveDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLBinaryArchiveDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLBinaryArchiveDescriptorClass) Alloc() MTLBinaryArchiveDescriptor {
	rv := objc.Send[MTLBinaryArchiveDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a binary shader archive that you want to create.
//
// # Choosing an archive file
//
//   - [MTLBinaryArchiveDescriptor.Url]: A URL to a Metal binary archive file.
//   - [MTLBinaryArchiveDescriptor.SetUrl]
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor
type MTLBinaryArchiveDescriptor struct {
	objectivec.Object
}

// MTLBinaryArchiveDescriptorFromID constructs a [MTLBinaryArchiveDescriptor] from an objc.ID.
//
// A description of a binary shader archive that you want to create.
func MTLBinaryArchiveDescriptorFromID(id objc.ID) MTLBinaryArchiveDescriptor {
	return MTLBinaryArchiveDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLBinaryArchiveDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLBinaryArchiveDescriptor] class.
//
// # Choosing an archive file
//
//   - [IMTLBinaryArchiveDescriptor.Url]: A URL to a Metal binary archive file.
//   - [IMTLBinaryArchiveDescriptor.SetUrl]
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor
type IMTLBinaryArchiveDescriptor interface {
	objectivec.IObject

	// Topic: Choosing an archive file

	// A URL to a Metal binary archive file.
	Url() foundation.INSURL
	SetUrl(value foundation.INSURL)

	// The domain for Metal binary archive errors.
	MTLBinaryArchiveDomain() string
}

// Init initializes the instance.
func (b MTLBinaryArchiveDescriptor) Init() MTLBinaryArchiveDescriptor {
	rv := objc.Send[MTLBinaryArchiveDescriptor](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MTLBinaryArchiveDescriptor) Autorelease() MTLBinaryArchiveDescriptor {
	rv := objc.Send[MTLBinaryArchiveDescriptor](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLBinaryArchiveDescriptor creates a new MTLBinaryArchiveDescriptor instance.
func NewMTLBinaryArchiveDescriptor() MTLBinaryArchiveDescriptor {
	class := getMTLBinaryArchiveDescriptorClass()
	rv := objc.Send[MTLBinaryArchiveDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A URL to a Metal binary archive file.
//
// # Discussion
// 
// You can use this method to load a binary archive you created with an
// [MTLBinaryArchive] instance’s [SerializeToURLError] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor/url
func (b MTLBinaryArchiveDescriptor) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (b MTLBinaryArchiveDescriptor) SetUrl(value foundation.INSURL) {
	objc.Send[struct{}](b.ID, objc.Sel("setUrl:"), value)
}
// The domain for Metal binary archive errors.
//
// See: https://developer.apple.com/documentation/metal/mtlbinaryarchivedomain
func (b MTLBinaryArchiveDescriptor) MTLBinaryArchiveDomain() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("MTLBinaryArchiveDomain"))
	return foundation.NSStringFromID(rv).String()
}

