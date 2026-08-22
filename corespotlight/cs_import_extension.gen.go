// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSImportExtension] class.
var (
	_CSImportExtensionClass     CSImportExtensionClass
	_CSImportExtensionClassOnce sync.Once
)

func getCSImportExtensionClass() CSImportExtensionClass {
	_CSImportExtensionClassOnce.Do(func() {
		_CSImportExtensionClass = CSImportExtensionClass{class: objc.GetClass("CSImportExtension")}
	})
	return _CSImportExtensionClass
}

// GetCSImportExtensionClass returns the class object for CSImportExtension.
func GetCSImportExtensionClass() CSImportExtensionClass {
	return getCSImportExtensionClass()
}

type CSImportExtensionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSImportExtensionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSImportExtensionClass) Alloc() CSImportExtension {
	rv := objc.Send[CSImportExtension](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides searchable attributes for file types that the app
// supports.
//
// # Overview
//
// To create a Spotlight File Import extension, add a target to your app using
// the Spotlight File Import extension template in Xcode. The template project
// contains a subclass of [CSImportExtension]. To index content on a user’s
// device, Core Spotlight loads your extension and invokes the
// [CSImportExtension.UpdateAttributesForFileAtURLError] method. Core
// Spotlight passes a [CSSearchableItemAttributeSet] and URL of a file to the
// extension, and you set properties that are relevant for the file.
//
// Typically, your extension loads details about the file and uses that
// information to set properties of the attribute set. For example, if your
// app contains files that are notes the user creates, it does the following:
//
// To specify the file types your app supports, set the value of
// [CSSupportedContentTypes] in your extension’s `Info.Plist()` file to an
// array of file type identifiers. For more information about file type
// identifiers, see [Uniform Type Identifiers]. The app in the previous
// example configures the extension’s `Info.Plist()` as follows:
//
// # Providing searchable attributes
//
//   - [CSImportExtension.UpdateAttributesForFileAtURLError]: Provides searchable attributes for a file at the specified URL.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSImportExtension
//
// [CSSupportedContentTypes]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSExtension/NSExtensionAttributes/CSSupportedContentTypes
// [Uniform Type Identifiers]: https://developer.apple.com/documentation/UniformTypeIdentifiers
type CSImportExtension struct {
	objectivec.Object
}

// CSImportExtensionFromID constructs a [CSImportExtension] from an objc.ID.
//
// An object that provides searchable attributes for file types that the app
// supports.
func CSImportExtensionFromID(id objc.ID) CSImportExtension {
	return CSImportExtension{objectivec.Object{ID: id}}
}

// NOTE: CSImportExtension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSImportExtension] class.
//
// # Providing searchable attributes
//
//   - [ICSImportExtension.UpdateAttributesForFileAtURLError]: Provides searchable attributes for a file at the specified URL.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSImportExtension
type ICSImportExtension interface {
	objectivec.IObject

	// Topic: Providing searchable attributes

	// Provides searchable attributes for a file at the specified URL.
	UpdateAttributesForFileAtURLError(attributes ICSSearchableItemAttributeSet, contentURL foundation.NSURL) (bool, error)
}

// Init initializes the instance.
func (c CSImportExtension) Init() CSImportExtension {
	rv := objc.Send[CSImportExtension](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSImportExtension) Autorelease() CSImportExtension {
	rv := objc.Send[CSImportExtension](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSImportExtension creates a new CSImportExtension instance.
func NewCSImportExtension() CSImportExtension {
	class := getCSImportExtensionClass()
	rv := objc.Send[CSImportExtension](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides searchable attributes for a file at the specified URL.
//
// attributes: The attribute set for the file at `contentURL`.
//
// contentURL: The URL of the file to provide attributes for.
//
// # Discussion
//
// When Core Spotlight invokes this method, update the properties of the
// attribute set according to the content in the specified file. For a
// complete list of properties available, see [CSSearchableItemAttributeSet].
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSImportExtension/update(_:forFileAt:)
func (c CSImportExtension) UpdateAttributesForFileAtURLError(attributes ICSSearchableItemAttributeSet, contentURL foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("updateAttributes:forFileAtURL:error:"), attributes, contentURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateAttributes:forFileAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
