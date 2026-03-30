// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

package uniformtypeidentifiers

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UTType] class.
var (
	_UTTypeClass     UTTypeClass
	_UTTypeClassOnce sync.Once
)

func getUTTypeClass() UTTypeClass {
	_UTTypeClassOnce.Do(func() {
		_UTTypeClass = UTTypeClass{class: objc.GetClass("UTType")}
	})
	return _UTTypeClass
}

// GetUTTypeClass returns the class object for UTType.
func GetUTTypeClass() UTTypeClass {
	return getUTTypeClass()
}

type UTTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UTTypeClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UTTypeClass) Alloc() UTType {
	rv := objc.Send[UTType](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a type of data to load, send, or receive.
//
// # Overview
//
// The [UTType] object may represent files on disk, abstract data types with
// no on-disk representation, or entirely unrelated hierarchical
// classification systems, such as hardware. Each instance has a unique
// [identifier], and helpful properties, [preferredFilenameExtension] and
// [preferredMIMEType].
//
// The [UTType] object may provide additional information related to the type.
// For example, it may include a localized user-facing description, a
// reference URL to technical documentation about the type, or its version
// number. You can look up types by their conformance to get either a type or
// a list of types that are relevant to your use case.
//
// To define your own types in your app’s `Info.Plist()`, see [Defining file
// and data types for your app].
//
// # Obtaining additional type information
//
//   - [UTType.Declared]: A Boolean value that indicates whether the system declares the type.
//   - [UTType.Dynamic]: A Boolean value that indicates whether the system generates the type.
//   - [UTType.PublicType]: A Boolean value that indicates whether the type is in the public domain.
//
// # Checking a type’s relationship to another type
//
//   - [UTType.ConformsToType]: Returns a Boolean value that indicates whether a type conforms to the type.
//   - [UTType.IsSubtypeOfType]: Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
//   - [UTType.IsSupertypeOfType]: Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
//
// # Type Properties
//
//   - [UTType.ShazamCustomCatalog]: A type that represents a custom catalog.
//   - [UTType.SetShazamCustomCatalog]
//   - [UTType.ShazamSignature]: A type that represents a signature.
//   - [UTType.SetShazamSignature]
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference
//
// [Defining file and data types for your app]: https://developer.apple.com/documentation/UniformTypeIdentifiers/defining-file-and-data-types-for-your-app
// [identifier]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/identifier
// [preferredFilenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/preferredFilenameExtension
// [preferredMIMEType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/preferredMIMEType
type UTType struct {
	objectivec.Object
}

// UTTypeFromID constructs a [UTType] from an objc.ID.
//
// An object that represents a type of data to load, send, or receive.
func UTTypeFromID(id objc.ID) UTType {
	return UTType{objectivec.Object{ID: id}}
}

// NOTE: UTType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UTType] class.
//
// # Obtaining additional type information
//
//   - [IUTType.Declared]: A Boolean value that indicates whether the system declares the type.
//   - [IUTType.Dynamic]: A Boolean value that indicates whether the system generates the type.
//   - [IUTType.PublicType]: A Boolean value that indicates whether the type is in the public domain.
//
// # Checking a type’s relationship to another type
//
//   - [IUTType.ConformsToType]: Returns a Boolean value that indicates whether a type conforms to the type.
//   - [IUTType.IsSubtypeOfType]: Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
//   - [IUTType.IsSupertypeOfType]: Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
//
// # Type Properties
//
//   - [IUTType.ShazamCustomCatalog]: A type that represents a custom catalog.
//   - [IUTType.SetShazamCustomCatalog]
//   - [IUTType.ShazamSignature]: A type that represents a signature.
//   - [IUTType.SetShazamSignature]
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference
type IUTType interface {
	objectivec.IObject

	// Topic: Obtaining additional type information

	// A Boolean value that indicates whether the system declares the type.
	Declared() bool
	// A Boolean value that indicates whether the system generates the type.
	Dynamic() bool
	// A Boolean value that indicates whether the type is in the public domain.
	PublicType() bool

	// Topic: Checking a type’s relationship to another type

	// Returns a Boolean value that indicates whether a type conforms to the type.
	ConformsToType(type_ IUTType) bool
	// Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
	IsSubtypeOfType(type_ IUTType) bool
	// Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
	IsSupertypeOfType(type_ IUTType) bool

	// Topic: Type Properties

	// A type that represents a custom catalog.
	ShazamCustomCatalog() IUTType
	SetShazamCustomCatalog(value IUTType)
	// A type that represents a signature.
	ShazamSignature() IUTType
	SetShazamSignature(value IUTType)

	// The string that represents the type.
	Identifier() string
	// A localized description of the type.
	LocalizedDescription() string
	// The preferred filename extension for the type.
	PreferredFilenameExtension() string
	// The preferred MIME type for the type.
	PreferredMIMEType() string
	// The reference URL for the type.
	ReferenceURL() foundation.INSURL
	// The set of types the type directly or indirectly conforms to.
	Supertypes() foundation.INSSet
	// The tag specification dictionary of the type.
	Tags() foundation.INSDictionary
	// The type’s version, if available.
	Version() foundation.NSNumber
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t UTType) Init() UTType {
	rv := objc.Send[UTType](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t UTType) Autorelease() UTType {
	rv := objc.Send[UTType](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewUTType creates a new UTType instance.
func NewUTType() UTType {
	class := getUTTypeClass()
	rv := objc.Send[UTType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a type your app owns based on an identifier.
//
// identifier: The identifier of your type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:)
func NewTypeExportedTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("exportedTypeWithIdentifier:"), objc.String(identifier))
	return UTTypeFromID(rv)
}

// Creates a type your app owns based on an identifier and a supertype that it
// conforms to.
//
// identifier: The identifier of your type.
//
// parentType: A type to extend for your own type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(exportedAs:conformingTo:)
func NewTypeExportedTypeWithIdentifierConformingToType(identifier string, parentType IUTType) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("exportedTypeWithIdentifier:conformingToType:"), objc.String(identifier), parentType)
	return UTTypeFromID(rv)
}

// Creates a type your app uses, but doesn’t own, based on an identifier.
//
// identifier: The identifier of your type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:)
func NewTypeImportedTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("importedTypeWithIdentifier:"), objc.String(identifier))
	return UTTypeFromID(rv)
}

// Creates a type your app uses, but doesn’t own, based on an identifier and
// a supertype that it conforms to.
//
// identifier: The identifier of your type.
//
// parentType: A type to extend with this type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(importedAs:conformingTo:)
func NewTypeImportedTypeWithIdentifierConformingToType(identifier string, parentType IUTType) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("importedTypeWithIdentifier:conformingToType:"), objc.String(identifier), parentType)
	return UTTypeFromID(rv)
}

// Creates a type that represents the specified filename extension.
//
// filenameExtension: The filename extension.
//
// # Discussion
//
// If the system recognizes the filename extension, the intializer returns the
// corresponding type; otherwise, the initializer returns a dynamic type whose
// [isDeclared] and [isPublic] properties are both set to false.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:)
//
// [isDeclared]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isDeclared
// [isPublic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isPublic
func NewTypeWithFilenameExtension(filenameExtension string) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithFilenameExtension:"), objc.String(filenameExtension))
	return UTTypeFromID(rv)
}

// Creates a type that represents the specified filename extension and
// conforms to an existing type.
//
// filenameExtension: The filename extension.
//
// supertype: The type the resulting type must conform to, such as [data] or [package].
//
// # Discussion
//
// If the system recognizes the filename extension, the intializer returns the
// corresponding type; otherwise, the initializer returns a dynamic type whose
// [isDeclared] and [isPublic] properties are both set to false.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:conformingTo:)
//
// [data]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/data
// [package]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/package
// [isDeclared]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isDeclared
// [isPublic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isPublic
func NewTypeWithFilenameExtensionConformingToType(filenameExtension string, supertype IUTType) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithFilenameExtension:conformingToType:"), objc.String(filenameExtension), supertype)
	return UTTypeFromID(rv)
}

// Creates a type based on an identifier.
//
// identifier: The identifier of your type.
//
// # Discussion
//
// This initializer returns `nil` if the system doesn’t know the type
// identifier.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(_:)
func NewTypeWithIdentifier(identifier string) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithIdentifier:"), objc.String(identifier))
	return UTTypeFromID(rv)
}

// Creates a type based on a MIME type.
//
// mimeType: A string that represents the MIME type.
//
// # Discussion
//
// This initializer returns `nil` if the system doesn’t know the MIME type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:)
func NewTypeWithMIMEType(mimeType string) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithMIMEType:"), objc.String(mimeType))
	return UTTypeFromID(rv)
}

// Creates a type based on a MIME type and a supertype that it conforms to.
//
// mimeType: A string that represents the MIME type.
//
// supertype: Another [UTType] instance that the resulting type must conform to; for
// example, [UTTypeData].
//
// # Discussion
//
// This initializer returns `nil` if the system doesn’t know the MIME type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:conformingTo:)
//
// [UTTypeData]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeData
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
func NewTypeWithMIMETypeConformingToType(mimeType string, supertype IUTType) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithMIMEType:conformingToType:"), objc.String(mimeType), supertype)
	return UTTypeFromID(rv)
}

// Creates a type that represents the specified tag and tag class and which
// conforms to an existing type.
//
// tag: The tag, such as a filename extension.
//
// tagClass: The appropriate tag class, such as [filenameExtension].
//
// supertype: The type the resulting type must conform to, such as [data].
//
// # Discussion
//
// If the system recognizes the filename extension, the intializer returns the
// corresponding type; otherwise, the initializer returns a dynamic type whose
// [isDeclared] and [isPublic] properties are both set to false.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(tag:tagClass:conformingToType:)
//
// [filenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClass/filenameExtension
// [data]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/data
// [isDeclared]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isDeclared
// [isPublic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isPublic
func NewTypeWithTagTagClassConformingToType(tag string, tagClass string, supertype IUTType) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return UTTypeFromID(rv)
}

// Returns a Boolean value that indicates whether a type conforms to the type.
//
// type: An [UTType] instance.
//
// # Return Value
//
// true if the type directly or indirectly conforms to `type`, or if it’s
// equal to `type`.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/conforms(to:)
//
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
func (t UTType) ConformsToType(type_ IUTType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("conformsToType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a type is higher in a
// hierarchy than the type.
//
// type: A [UTType] instance.
//
// # Return Value
//
// true if the type directly or indirectly conforms to `type`, but returns
// false if it’s equal to `type`.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSubtype(of:)
//
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
func (t UTType) IsSubtypeOfType(type_ IUTType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSubtypeOfType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a type is lower in a
// hierarchy than the type.
//
// type: A [UTType] instance.
//
// # Return Value
//
// true if `type` directly or indirectly conforms to the type, but returns
// false if it’s equal to the type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSupertype(of:)
//
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
func (t UTType) IsSupertypeOfType(type_ IUTType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSupertypeOfType:"), type_)
	return rv
}
func (t UTType) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns an array of types from the provided tag and tag class.
//
// tag: The desired tag, such as a filename extension.
//
// tagClass: The tag class, such as [UTTagClassFilenameExtension].
//
// supertype: Another type that the resulting type must conform to; for example,
// [UTTypeData].
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/types(tag:tagClass:conformingTo:)
//
// [UTTagClassFilenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClassFilenameExtension
// [UTTypeData]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeData
func (_UTTypeClass UTTypeClass) TypesWithTagTagClassConformingToType(tag string, tagClass string, supertype IUTType) []UTType {
	rv := objc.Send[[]objc.ID](objc.ID(_UTTypeClass.class), objc.Sel("typesWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return objc.ConvertSlice(rv, func(id objc.ID) UTType {
		return UTTypeFromID(id)
	})
}

// A Boolean value that indicates whether the system declares the type.
//
// # Discussion
//
// The system either declares a type or dynamically generates a type, but not
// both.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isDeclared
func (t UTType) Declared() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isDeclared"))
	return rv
}

// A Boolean value that indicates whether the system generates the type.
//
// # Discussion
//
// The system recognizes dynamic types, but they may not be directly declared
// or claimed by an app. The system returns dynamic types when it encounters a
// file whose metadata doesn’t have a corresponding type known to the
// system.
//
// The system either declares a type or dynamically generates a type, but not
// both.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isDynamic
func (t UTType) Dynamic() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isDynamic"))
	return rv
}

// A Boolean value that indicates whether the type is in the public domain.
//
// # Discussion
//
// Types in the public domain have identifiers starting with `public`, and are
// generally defined by a standards body or by convention. Public types
// aren’t dynamic.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isPublic
func (t UTType) PublicType() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isPublicType"))
	return rv
}

// A type that represents a custom catalog.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/shazamcustomcatalog
func (t UTType) ShazamCustomCatalog() IUTType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("shazamCustomCatalog"))
	return UTTypeFromID(objc.ID(rv))
}
func (t UTType) SetShazamCustomCatalog(value IUTType) {
	objc.Send[struct{}](t.ID, objc.Sel("setShazamCustomCatalog:"), value)
}

// A type that represents a signature.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/shazamsignature
func (t UTType) ShazamSignature() IUTType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("shazamSignature"))
	return UTTypeFromID(objc.ID(rv))
}
func (t UTType) SetShazamSignature(value IUTType) {
	objc.Send[struct{}](t.ID, objc.Sel("setShazamSignature:"), value)
}

// The string that represents the type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/identifier
func (t UTType) Identifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// A localized description of the type.
//
// # Discussion
//
// If the type doesn’t provide a description, the system searches its
// supertypes. A dynamic type doesn’t have localized description, even if
// its supertypes do.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/localizedDescription
func (t UTType) LocalizedDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}

// The preferred filename extension for the type.
//
// # Discussion
//
// If available, the preferred (first available) tag of class
// [filenameExtension].
//
// Many types require the generation of a filename; for example, when saving a
// file to disk. If not `nil`, the value of this property is the best
// available filename extension for this type.
//
// The value of this property is equivalent to, but more efficient than:
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/preferredFilenameExtension
//
// [filenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClass/filenameExtension
func (t UTType) PreferredFilenameExtension() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("preferredFilenameExtension"))
	return foundation.NSStringFromID(rv).String()
}

// The preferred MIME type for the type.
//
// # Discussion
//
// If available, the preferred (first available) tag of class [mimeType]. If
// not `nil`, the value of this property is the best available MIME type value
// for this type.
//
// The value of this property is equivalent to, but more efficient than:
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/preferredMIMEType
//
// [mimeType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClass/mimeType
func (t UTType) PreferredMIMEType() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("preferredMIMEType"))
	return foundation.NSStringFromID(rv).String()
}

// The reference URL for the type.
//
// # Discussion
//
// A reference URL is a human-readable document that describes a type. Most
// types don’t specify reference URLs.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/referenceURL
func (t UTType) ReferenceURL() foundation.INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("referenceURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The set of types the type directly or indirectly conforms to.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/supertypes
func (t UTType) Supertypes() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("supertypes"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The tag specification dictionary of the type.
//
// # Discussion
//
// Uniform Type Identifiers don’t store tag information for nonstandard tag
// classes. Identifiers normalize string values into arrays that contain those
// strings. For example, a tag specification dictionary with values of:
//
// Normalizes to:
//
// Use the tag class [mimeType] or [filenameExtension] to retrieve the list of
// supporting MIME types or filename extensions for your type. For example,
// the following example retrieves a list of the filename extensions for the
// [mpeg] type:
//
// Types that have no tags for the requested tag class return a nil, not an
// empty array.
//
// To get the preferred filename extension or MIME type, use the tag class
// [preferredFilenameExtension] or [preferredMIMEType], respectively.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/tags
//
// [filenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClass/filenameExtension
// [mimeType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClass/mimeType
// [mpeg]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/mpeg
// [preferredFilenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/preferredFilenameExtension
// [preferredMIMEType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/preferredMIMEType
func (t UTType) Tags() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tags"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The type’s version, if available.
//
// # Discussion
//
// Most types don’t have a version.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/version
func (t UTType) Version() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("version"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A type representing the @c SHCustomCatalog file format with the
// .shazamcatalog extension
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHCustomCatalogContentType
func (_UTTypeClass UTTypeClass) SHCustomCatalogContentType() UTType {
	rv := objc.Send[objc.ID](objc.ID(_UTTypeClass.class), objc.Sel("SHCustomCatalogContentType"))
	return UTTypeFromID(objc.ID(rv))
}

// A type representing the @c SHSignature file format with the
// .shazamsignature extension
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-c.class/SHSignatureContentType
func (_UTTypeClass UTTypeClass) SHSignatureContentType() UTType {
	rv := objc.Send[objc.ID](objc.ID(_UTTypeClass.class), objc.Sel("SHSignatureContentType"))
	return UTTypeFromID(objc.ID(rv))
}
