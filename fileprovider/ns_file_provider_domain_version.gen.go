// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderDomainVersion] class.
var (
	_NSFileProviderDomainVersionClass     NSFileProviderDomainVersionClass
	_NSFileProviderDomainVersionClassOnce sync.Once
)

func getNSFileProviderDomainVersionClass() NSFileProviderDomainVersionClass {
	_NSFileProviderDomainVersionClassOnce.Do(func() {
		_NSFileProviderDomainVersionClass = NSFileProviderDomainVersionClass{class: objc.GetClass("NSFileProviderDomainVersion")}
	})
	return _NSFileProviderDomainVersionClass
}

// GetNSFileProviderDomainVersionClass returns the class object for NSFileProviderDomainVersion.
func GetNSFileProviderDomainVersionClass() NSFileProviderDomainVersionClass {
	return getNSFileProviderDomainVersionClass()
}

type NSFileProviderDomainVersionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderDomainVersionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderDomainVersionClass) Alloc() NSFileProviderDomainVersion {
	rv := objc.Send[NSFileProviderDomainVersion](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An opaque object that identifies a specific version of a domain.
//
// # Overview
//
// The file provider extension is responsible for assigning and updating the
// domain version. To specify the domain version, adopt the
// [NSFileProviderDomainState] protocol. The system then calls your
// extension’s [DomainVersion] method to read the current version.
//
// The system reads the domain version after you call:
//
// - The
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
// completion handler - The
// [ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler]
// completion handler - The
// [DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler]
// completion handler - The [ItemForIdentifierRequestCompletionHandler]
// completion handler - The [FinishEnumeratingUpToPage] or
// [FinishEnumeratingWithError] method when enumerating the materialized set.
//
// The system always reads the domain version on the same dispatch queue as
// the completion handler.
//
// Your extension defines when the domain version changes. When you update the
// version, call the
// [NSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler]
// and passing the [workingSet] constant as the `containerItemIdentifier`
// property. This notifies the system of the update. The system ignores any
// lower versions.
//
// When the system discovers a change on disk, it associates that change with
// the current domain version. It then includes the version in the
// [NSFileProviderRequest] object passed to the file provider extension.
//
// Only file provider extensions based on the
// [NSFileProviderReplicatedExtension] use instances of this class. Each
// version object is immutable. You can use them as keys in a dictionary.
//
// # Creating Versions
//
//   - [NSFileProviderDomainVersion.Next]: Creates a new version that supersedes the current version.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion
//
// [workingSet]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/workingSet
type NSFileProviderDomainVersion struct {
	objectivec.Object
}

// NSFileProviderDomainVersionFromID constructs a [NSFileProviderDomainVersion] from an objc.ID.
//
// An opaque object that identifies a specific version of a domain.
func NSFileProviderDomainVersionFromID(id objc.ID) NSFileProviderDomainVersion {
	return NSFileProviderDomainVersion{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderDomainVersion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderDomainVersion] class.
//
// # Creating Versions
//
//   - [INSFileProviderDomainVersion.Next]: Creates a new version that supersedes the current version.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion
type INSFileProviderDomainVersion interface {
	objectivec.IObject

	// Topic: Creating Versions

	// Creates a new version that supersedes the current version.
	Next() INSFileProviderDomainVersion

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NSFileProviderDomainVersion) Init() NSFileProviderDomainVersion {
	rv := objc.Send[NSFileProviderDomainVersion](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderDomainVersion) Autorelease() NSFileProviderDomainVersion {
	rv := objc.Send[NSFileProviderDomainVersion](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderDomainVersion creates a new NSFileProviderDomainVersion instance.
func NewNSFileProviderDomainVersion() NSFileProviderDomainVersion {
	class := getNSFileProviderDomainVersionClass()
	rv := objc.Send[NSFileProviderDomainVersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new version that supersedes the current version.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomainVersion/next()
func (f NSFileProviderDomainVersion) Next() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("next"))
	return NSFileProviderDomainVersionFromID(rv)
}
func (f NSFileProviderDomainVersion) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}
