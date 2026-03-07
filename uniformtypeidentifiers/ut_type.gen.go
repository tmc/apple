// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

package uniformtypeidentifiers

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
// [Defining file and data types for your app]: https://developer.apple.com/documentation/UniformTypeIdentifiers/defining-file-and-data-types-for-your-app
// [identifier]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/identifier
// [preferredFilenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/preferredFilenameExtension
// [preferredMIMEType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/preferredMIMEType
//
// # Identifying a type
//
//   - [UTType.Identifier]: The string that represents the type.
//   - [UTType.SetIdentifier]
//
// # Obtaining tags
//
//   - [UTType.PreferredFilenameExtension]: The preferred filename extension for the type.
//   - [UTType.SetPreferredFilenameExtension]
//   - [UTType.PreferredMIMEType]: The preferred MIME type for the type.
//   - [UTType.SetPreferredMIMEType]
//   - [UTType.Tags]: The tag specification dictionary of the type.
//   - [UTType.SetTags]
//
// # Obtaining additional type information
//
//   - [UTType.Declared]: A Boolean value that indicates whether the system declares the type.
//   - [UTType.Dynamic]: A Boolean value that indicates whether the system generates the type.
//   - [UTType.PublicType]: A Boolean value that indicates whether the type is in the public domain.
//   - [UTType.ReferenceURL]: The reference URL for the type.
//   - [UTType.SetReferenceURL]
//   - [UTType.Version]: The type’s version, if available.
//   - [UTType.SetVersion]
//
// # Checking a type’s relationship to another type
//
//   - [UTType.Supertypes]: The set of types the type directly or indirectly conforms to.
//   - [UTType.SetSupertypes]
//   - [UTType.ConformsToType]: Returns a Boolean value that indicates whether a type conforms to the type.
//   - [UTType.IsSubtypeOfType]: Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
//   - [UTType.IsSupertypeOfType]: Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
//
// # Describing a type
//
//   - [UTType.LocalizedDescription]: A localized description of the type.
//   - [UTType.SetLocalizedDescription]
//
// # Type Properties
//
//   - [UTType.ShazamCustomCatalog]: A type that represents a custom catalog.
//   - [UTType.SetShazamCustomCatalog]
//   - [UTType.ShazamSignature]: A type that represents a signature.
//   - [UTType.SetShazamSignature]
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference
type UTType struct {
	objectivec.Object
}

// UTTypeFromID constructs a [UTType] from an objc.ID.
//
// An object that represents a type of data to load, send, or receive.
func UTTypeFromID(id objc.ID) UTType {
	return UTType{objectivec.Object{id}}
}
// NOTE: UTType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [UTType] class.
//
// # Identifying a type
//
//   - [IUTType.Identifier]: The string that represents the type.
//   - [IUTType.SetIdentifier]
//
// # Obtaining tags
//
//   - [IUTType.PreferredFilenameExtension]: The preferred filename extension for the type.
//   - [IUTType.SetPreferredFilenameExtension]
//   - [IUTType.PreferredMIMEType]: The preferred MIME type for the type.
//   - [IUTType.SetPreferredMIMEType]
//   - [IUTType.Tags]: The tag specification dictionary of the type.
//   - [IUTType.SetTags]
//
// # Obtaining additional type information
//
//   - [IUTType.Declared]: A Boolean value that indicates whether the system declares the type.
//   - [IUTType.Dynamic]: A Boolean value that indicates whether the system generates the type.
//   - [IUTType.PublicType]: A Boolean value that indicates whether the type is in the public domain.
//   - [IUTType.ReferenceURL]: The reference URL for the type.
//   - [IUTType.SetReferenceURL]
//   - [IUTType.Version]: The type’s version, if available.
//   - [IUTType.SetVersion]
//
// # Checking a type’s relationship to another type
//
//   - [IUTType.Supertypes]: The set of types the type directly or indirectly conforms to.
//   - [IUTType.SetSupertypes]
//   - [IUTType.ConformsToType]: Returns a Boolean value that indicates whether a type conforms to the type.
//   - [IUTType.IsSubtypeOfType]: Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
//   - [IUTType.IsSupertypeOfType]: Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
//
// # Describing a type
//
//   - [IUTType.LocalizedDescription]: A localized description of the type.
//   - [IUTType.SetLocalizedDescription]
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

	// Topic: Identifying a type

	// The string that represents the type.
	Identifier() string
	SetIdentifier(value string)

	// Topic: Obtaining tags

	// The preferred filename extension for the type.
	PreferredFilenameExtension() string
	SetPreferredFilenameExtension(value string)
	// The preferred MIME type for the type.
	PreferredMIMEType() string
	SetPreferredMIMEType(value string)
	// The tag specification dictionary of the type.
	Tags() string
	SetTags(value string)

	// Topic: Obtaining additional type information

	// A Boolean value that indicates whether the system declares the type.
	Declared() bool
	// A Boolean value that indicates whether the system generates the type.
	Dynamic() bool
	// A Boolean value that indicates whether the type is in the public domain.
	PublicType() bool
	// The reference URL for the type.
	ReferenceURL() foundation.INSURL
	SetReferenceURL(value foundation.INSURL)
	// The type’s version, if available.
	Version() int
	SetVersion(value int)

	// Topic: Checking a type’s relationship to another type

	// The set of types the type directly or indirectly conforms to.
	Supertypes() IUTType
	SetSupertypes(value IUTType)
	// Returns a Boolean value that indicates whether a type conforms to the type.
	ConformsToType(type_ IUTType) bool
	// Returns a Boolean value that indicates whether a type is higher in a hierarchy than the type.
	IsSubtypeOfType(type_ IUTType) bool
	// Returns a Boolean value that indicates whether a type is lower in a hierarchy than the type.
	IsSupertypeOfType(type_ IUTType) bool

	// Topic: Describing a type

	// A localized description of the type.
	LocalizedDescription() string
	SetLocalizedDescription(value string)

	// Topic: Type Properties

	// A type that represents a custom catalog.
	ShazamCustomCatalog() IUTType
	SetShazamCustomCatalog(value IUTType)
	// A type that represents a signature.
	ShazamSignature() IUTType
	SetShazamSignature(value IUTType)

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
// [isDeclared] and [isPublic] properties are both set to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isDeclared]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isDeclared
// [isPublic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isPublic
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:)
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
// //
// [data]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/data
// [package]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/package
//
// # Discussion
// 
// If the system recognizes the filename extension, the intializer returns the
// corresponding type; otherwise, the initializer returns a dynamic type whose
// [isDeclared] and [isPublic] properties are both set to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isDeclared]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isDeclared
// [isPublic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isPublic
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(filenameExtension:conformingTo:)
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
// //
// [UTTypeData]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeData
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
//
// # Discussion
// 
// This initializer returns `nil` if the system doesn’t know the MIME type.
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(mimeType:conformingTo:)
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
// //
// [filenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClass/filenameExtension
//
// supertype: The type the resulting type must conform to, such as [data].
// //
// [data]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/data
//
// # Discussion
// 
// If the system recognizes the filename extension, the intializer returns the
// corresponding type; otherwise, the initializer returns a dynamic type whose
// [isDeclared] and [isPublic] properties are both set to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isDeclared]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isDeclared
// [isPublic]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct/isPublic
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/init(tag:tagClass:conformingToType:)
func NewTypeWithTagTagClassConformingToType(tag string, tagClass string, supertype IUTType) UTType {
	rv := objc.Send[objc.ID](objc.ID(getUTTypeClass().class), objc.Sel("typeWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return UTTypeFromID(rv)
}







// Returns a Boolean value that indicates whether a type conforms to the type.
//
// type: An [UTType] instance.
// //
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
//
// # Return Value
// 
// [true] if the type directly or indirectly conforms to `type`, or if it’s
// equal to `type`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/conforms(to:)
func (t UTType) ConformsToType(type_ IUTType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("conformsToType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a type is higher in a
// hierarchy than the type.
//
// type: A [UTType] instance.
// //
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
//
// # Return Value
// 
// [true] if the type directly or indirectly conforms to `type`, but returns
// [false] if it’s equal to `type`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSubtype(of:)
func (t UTType) IsSubtypeOfType(type_ IUTType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSubtypeOfType:"), type_)
	return rv
}

// Returns a Boolean value that indicates whether a type is lower in a
// hierarchy than the type.
//
// type: A [UTType] instance.
// //
// [UTType]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTType-swift.struct
//
// # Return Value
// 
// [true] if `type` directly or indirectly conforms to the type, but returns
// [false] if it’s equal to the type.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/isSupertype(of:)
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
// //
// [UTTagClassFilenameExtension]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTagClassFilenameExtension
//
// supertype: Another type that the resulting type must conform to; for example,
// [UTTypeData].
// //
// [UTTypeData]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeData
//
// See: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeReference/types(tag:tagClass:conformingTo:)
func (_UTTypeClass UTTypeClass) TypesWithTagTagClassConformingToType(tag string, tagClass string, supertype IUTType) []UTType {
	rv := objc.Send[[]objc.ID](objc.ID(_UTTypeClass.class), objc.Sel("typesWithTag:tagClass:conformingToType:"), objc.String(tag), objc.String(tagClass), supertype)
	return objc.ConvertSlice(rv, func(id objc.ID) UTType {
		return UTTypeFromID(id)
	})
}








// The string that represents the type.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/identifier
func (t UTType) Identifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t UTType) SetIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setIdentifier:"), objc.String(value))
}



// The preferred filename extension for the type.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/preferredfilenameextension
func (t UTType) PreferredFilenameExtension() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("preferredFilenameExtension"))
	return foundation.NSStringFromID(rv).String()
}
func (t UTType) SetPreferredFilenameExtension(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreferredFilenameExtension:"), objc.String(value))
}



// The preferred MIME type for the type.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/preferredmimetype
func (t UTType) PreferredMIMEType() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("preferredMIMEType"))
	return foundation.NSStringFromID(rv).String()
}
func (t UTType) SetPreferredMIMEType(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreferredMIMEType:"), objc.String(value))
}



// The tag specification dictionary of the type.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/tags
func (t UTType) Tags() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tags"))
	return foundation.NSStringFromID(rv).String()
}
func (t UTType) SetTags(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setTags:"), objc.String(value))
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



// The reference URL for the type.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/referenceurl
func (t UTType) ReferenceURL() foundation.INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("referenceURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (t UTType) SetReferenceURL(value foundation.INSURL) {
	objc.Send[struct{}](t.ID, objc.Sel("setReferenceURL:"), value)
}



// The type’s version, if available.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/version
func (t UTType) Version() int {
	rv := objc.Send[int](t.ID, objc.Sel("version"))
	return rv
}
func (t UTType) SetVersion(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setVersion:"), value)
}



// The set of types the type directly or indirectly conforms to.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/supertypes
func (t UTType) Supertypes() IUTType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("supertypes"))
	return UTTypeFromID(objc.ID(rv))
}
func (t UTType) SetSupertypes(value IUTType) {
	objc.Send[struct{}](t.ID, objc.Sel("setSupertypes:"), value)
}



// A localized description of the type.
//
// See: https://developer.apple.com/documentation/uniformtypeidentifiers/uttype-swift.struct/localizeddescription
func (t UTType) LocalizedDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("localizedDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (t UTType) SetLocalizedDescription(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
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































