// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelCollectionEntry] class.
var (
	_MLModelCollectionEntryClass     MLModelCollectionEntryClass
	_MLModelCollectionEntryClassOnce sync.Once
)

func getMLModelCollectionEntryClass() MLModelCollectionEntryClass {
	_MLModelCollectionEntryClassOnce.Do(func() {
		_MLModelCollectionEntryClass = MLModelCollectionEntryClass{class: objc.GetClass("MLModelCollectionEntry")}
	})
	return _MLModelCollectionEntryClass
}

// GetMLModelCollectionEntryClass returns the class object for MLModelCollectionEntry.
func GetMLModelCollectionEntryClass() MLModelCollectionEntryClass {
	return getMLModelCollectionEntryClass()
}

type MLModelCollectionEntryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelCollectionEntryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelCollectionEntryClass) Alloc() MLModelCollectionEntry {
	rv := objc.Send[MLModelCollectionEntry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelCollectionEntry._initWithModelIdentifierModelUrl]
//   - [MLModelCollectionEntry.IsEqualToModelCollectionEntry]
//   - [MLModelCollectionEntry.ModelIdentifier]
//   - [MLModelCollectionEntry.ModelURL]
// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry
type MLModelCollectionEntry struct {
	objectivec.Object
}

// MLModelCollectionEntryFromID constructs a [MLModelCollectionEntry] from an objc.ID.
func MLModelCollectionEntryFromID(id objc.ID) MLModelCollectionEntry {
	return MLModelCollectionEntry{objectivec.Object{ID: id}}
}
// Ensure MLModelCollectionEntry implements IMLModelCollectionEntry.
var _ IMLModelCollectionEntry = MLModelCollectionEntry{}

// An interface definition for the [MLModelCollectionEntry] class.
//
// # Methods
//
//   - [IMLModelCollectionEntry._initWithModelIdentifierModelUrl]
//   - [IMLModelCollectionEntry.IsEqualToModelCollectionEntry]
//   - [IMLModelCollectionEntry.ModelIdentifier]
//   - [IMLModelCollectionEntry.ModelURL]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry
type IMLModelCollectionEntry interface {
	objectivec.IObject

	// Topic: Methods

	_initWithModelIdentifierModelUrl(identifier objectivec.IObject, url foundation.INSURL) objectivec.IObject
	IsEqualToModelCollectionEntry(entry objectivec.IObject) bool
	ModelIdentifier() string
	ModelURL() foundation.INSURL
}

// Init initializes the instance.
func (m MLModelCollectionEntry) Init() MLModelCollectionEntry {
	rv := objc.Send[MLModelCollectionEntry](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelCollectionEntry) Autorelease() MLModelCollectionEntry {
	rv := objc.Send[MLModelCollectionEntry](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelCollectionEntry creates a new MLModelCollectionEntry instance.
func NewMLModelCollectionEntry() MLModelCollectionEntry {
	class := getMLModelCollectionEntryClass()
	rv := objc.Send[MLModelCollectionEntry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry/_initWithModelIdentifier:modelUrl:
func (m MLModelCollectionEntry) _initWithModelIdentifierModelUrl(identifier objectivec.IObject, url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_initWithModelIdentifier:modelUrl:"), identifier, url)
	return objectivec.Object{ID: rv}
}

// InitWithModelIdentifierModelUrl is an exported wrapper for the private method _initWithModelIdentifierModelUrl.
func (m MLModelCollectionEntry) InitWithModelIdentifierModelUrl(identifier objectivec.IObject, url foundation.INSURL) objectivec.IObject {
	return m._initWithModelIdentifierModelUrl(identifier, url)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry/isEqualToModelCollectionEntry:
func (m MLModelCollectionEntry) IsEqualToModelCollectionEntry(entry objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEqualToModelCollectionEntry:"), entry)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry/entryWithModelIdentifier:modelURL:
func (_MLModelCollectionEntryClass MLModelCollectionEntryClass) EntryWithModelIdentifierModelURL(identifier objectivec.IObject, url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelCollectionEntryClass.class), objc.Sel("entryWithModelIdentifier:modelURL:"), identifier, url)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry/modelIdentifier
func (m MLModelCollectionEntry) ModelIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModelCollectionEntry/modelURL
func (m MLModelCollectionEntry) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

