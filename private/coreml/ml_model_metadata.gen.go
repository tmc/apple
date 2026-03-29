// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelMetadata] class.
var (
	_MLModelMetadataClass     MLModelMetadataClass
	_MLModelMetadataClassOnce sync.Once
)

func getMLModelMetadataClass() MLModelMetadataClass {
	_MLModelMetadataClassOnce.Do(func() {
		_MLModelMetadataClass = MLModelMetadataClass{class: objc.GetClass("MLModelMetadata")}
	})
	return _MLModelMetadataClass
}

// GetMLModelMetadataClass returns the class object for MLModelMetadata.
func GetMLModelMetadataClass() MLModelMetadataClass {
	return getMLModelMetadataClass()
}

type MLModelMetadataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelMetadataClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelMetadataClass) Alloc() MLModelMetadata {
	rv := objc.Send[MLModelMetadata](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelMetadata.Author]
//   - [MLModelMetadata.CreatorDefined]
//   - [MLModelMetadata.License]
//   - [MLModelMetadata.Name]
//   - [MLModelMetadata.ShortDescription]
//   - [MLModelMetadata.VersionString]
//   - [MLModelMetadata.InitWithName]
//   - [MLModelMetadata.InitWithNameShortDescriptionVersionStringAuthorLicenseCreatorDefined]
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata
type MLModelMetadata struct {
	objectivec.Object
}

// MLModelMetadataFromID constructs a [MLModelMetadata] from an objc.ID.
func MLModelMetadataFromID(id objc.ID) MLModelMetadata {
	return MLModelMetadata{objectivec.Object{ID: id}}
}
// Ensure MLModelMetadata implements IMLModelMetadata.
var _ IMLModelMetadata = MLModelMetadata{}

// An interface definition for the [MLModelMetadata] class.
//
// # Methods
//
//   - [IMLModelMetadata.Author]
//   - [IMLModelMetadata.CreatorDefined]
//   - [IMLModelMetadata.License]
//   - [IMLModelMetadata.Name]
//   - [IMLModelMetadata.ShortDescription]
//   - [IMLModelMetadata.VersionString]
//   - [IMLModelMetadata.InitWithName]
//   - [IMLModelMetadata.InitWithNameShortDescriptionVersionStringAuthorLicenseCreatorDefined]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata
type IMLModelMetadata interface {
	objectivec.IObject

	// Topic: Methods

	Author() string
	CreatorDefined() foundation.INSDictionary
	License() string
	Name() string
	ShortDescription() string
	VersionString() string
	InitWithName(name objectivec.IObject) MLModelMetadata
	InitWithNameShortDescriptionVersionStringAuthorLicenseCreatorDefined(name objectivec.IObject, description objectivec.IObject, string_ objectivec.IObject, author objectivec.IObject, license objectivec.IObject, defined objectivec.IObject) MLModelMetadata
}

// Init initializes the instance.
func (m MLModelMetadata) Init() MLModelMetadata {
	rv := objc.Send[MLModelMetadata](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelMetadata) Autorelease() MLModelMetadata {
	rv := objc.Send[MLModelMetadata](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelMetadata creates a new MLModelMetadata instance.
func NewMLModelMetadata() MLModelMetadata {
	class := getMLModelMetadataClass()
	rv := objc.Send[MLModelMetadata](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/initWithName:
func NewModelMetadataWithName(name objectivec.IObject) MLModelMetadata {
	instance := getMLModelMetadataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), name)
	return MLModelMetadataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/initWithName:shortDescription:versionString:author:license:creatorDefined:
func NewModelMetadataWithNameShortDescriptionVersionStringAuthorLicenseCreatorDefined(name objectivec.IObject, description objectivec.IObject, string_ objectivec.IObject, author objectivec.IObject, license objectivec.IObject, defined objectivec.IObject) MLModelMetadata {
	instance := getMLModelMetadataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:shortDescription:versionString:author:license:creatorDefined:"), name, description, string_, author, license, defined)
	return MLModelMetadataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/initWithName:
func (m MLModelMetadata) InitWithName(name objectivec.IObject) MLModelMetadata {
	rv := objc.Send[MLModelMetadata](m.ID, objc.Sel("initWithName:"), name)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/initWithName:shortDescription:versionString:author:license:creatorDefined:
func (m MLModelMetadata) InitWithNameShortDescriptionVersionStringAuthorLicenseCreatorDefined(name objectivec.IObject, description objectivec.IObject, string_ objectivec.IObject, author objectivec.IObject, license objectivec.IObject, defined objectivec.IObject) MLModelMetadata {
	rv := objc.Send[MLModelMetadata](m.ID, objc.Sel("initWithName:shortDescription:versionString:author:license:creatorDefined:"), name, description, string_, author, license, defined)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/author
func (m MLModelMetadata) Author() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("author"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/creatorDefined
func (m MLModelMetadata) CreatorDefined() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("creatorDefined"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/license
func (m MLModelMetadata) License() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("license"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/name
func (m MLModelMetadata) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/shortDescription
func (m MLModelMetadata) ShortDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shortDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelMetadata/versionString
func (m MLModelMetadata) VersionString() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("versionString"))
	return foundation.NSStringFromID(rv).String()
}

