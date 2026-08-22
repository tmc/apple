// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSRLanguageItem] class.
var (
	_SOSRLanguageItemClass     SOSRLanguageItemClass
	_SOSRLanguageItemClassOnce sync.Once
)

func getSOSRLanguageItemClass() SOSRLanguageItemClass {
	_SOSRLanguageItemClassOnce.Do(func() {
		_SOSRLanguageItemClass = SOSRLanguageItemClass{class: objc.GetClass("SOSRLanguageItem")}
	})
	return _SOSRLanguageItemClass
}

// GetSOSRLanguageItemClass returns the class object for SOSRLanguageItem.
func GetSOSRLanguageItemClass() SOSRLanguageItemClass {
	return getSOSRLanguageItemClass()
}

type SOSRLanguageItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRLanguageItemClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRLanguageItemClass) Alloc() SOSRLanguageItem {
	rv := objc.SendIfResponds[SOSRLanguageItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSRLanguageItem.InitWithLocaleIdentifier]
//   - [SOSRLanguageItem.InitWithLocaleIdentifierUsingOffline]
type SOSRLanguageItem struct {
	SOSpeechItem
}

// SOSRLanguageItemFromID constructs a [SOSRLanguageItem] from an objc.ID.
func SOSRLanguageItemFromID(id objc.ID) SOSRLanguageItem {
	return SOSRLanguageItem{SOSpeechItem: SOSpeechItemFromID(id)}
}

// Ensure SOSRLanguageItem implements ISOSRLanguageItem.
var _ ISOSRLanguageItem = SOSRLanguageItem{}

// An interface definition for the [SOSRLanguageItem] class.
//
// # Methods
//
//   - [ISOSRLanguageItem.InitWithLocaleIdentifier]
//   - [ISOSRLanguageItem.InitWithLocaleIdentifierUsingOffline]
type ISOSRLanguageItem interface {
	ISOSpeechItem

	// Topic: Methods

	InitWithLocaleIdentifier(identifier objectivec.IObject) SOSRLanguageItem
	InitWithLocaleIdentifierUsingOffline(identifier objectivec.IObject, offline bool) SOSRLanguageItem
}

// Init initializes the instance.
func (s SOSRLanguageItem) Init() SOSRLanguageItem {
	rv := objc.SendIfResponds[SOSRLanguageItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRLanguageItem) Autorelease() SOSRLanguageItem {
	rv := objc.SendIfResponds[SOSRLanguageItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRLanguageItem creates a new SOSRLanguageItem instance.
func NewSOSRLanguageItem() SOSRLanguageItem {
	class := getSOSRLanguageItemClass()
	rv := objc.SendIfResponds[SOSRLanguageItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSOSRLanguageItemWithDownloadableBundleIdentifierProperties(identifier objectivec.IObject, properties objectivec.IObject) SOSRLanguageItem {
	instance := getSOSRLanguageItemClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDownloadableBundleIdentifier:properties:"), identifier, properties)
	return SOSRLanguageItemFromID(rv)
}

func NewSOSRLanguageItemWithLocaleIdentifier(identifier objectivec.IObject) SOSRLanguageItem {
	instance := getSOSRLanguageItemClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLocaleIdentifier:"), identifier)
	return SOSRLanguageItemFromID(rv)
}

func NewSOSRLanguageItemWithLocaleIdentifierUsingOffline(identifier objectivec.IObject, offline bool) SOSRLanguageItem {
	instance := getSOSRLanguageItemClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLocaleIdentifier:usingOffline:"), identifier, offline)
	return SOSRLanguageItemFromID(rv)
}

func (s SOSRLanguageItem) InitWithLocaleIdentifier(identifier objectivec.IObject) SOSRLanguageItem {
	rv := objc.SendIfResponds[SOSRLanguageItem](s.ID, objc.Sel("initWithLocaleIdentifier:"), identifier)
	return rv
}
func (s SOSRLanguageItem) InitWithLocaleIdentifierUsingOffline(identifier objectivec.IObject, offline bool) SOSRLanguageItem {
	rv := objc.SendIfResponds[SOSRLanguageItem](s.ID, objc.Sel("initWithLocaleIdentifier:usingOffline:"), identifier, offline)
	return rv
}

func (_SOSRLanguageItemClass SOSRLanguageItemClass) AvailableLocalRecognizerLanguageItems() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("availableLocalRecognizerLanguageItems"))
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) BestIndexFromLanguageItemsForLocaleIdentifier(items objectivec.IObject, identifier objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("bestIndexFromLanguageItems:forLocaleIdentifier:"), items, identifier)
	return rv
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) DownloadableLocalSRLanguageItems() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("downloadableLocalSRLanguageItems"))
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) EngineIdentifierFromLocaleIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("engineIdentifierFromLocaleIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) LanguageItemFromLanguageItemsMatchingLocaleIdentifier(items objectivec.IObject, identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("languageItemFromLanguageItems:matchingLocaleIdentifier:"), items, identifier)
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) LanguageItemsFromLocaleIdentifiers(identifiers objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("languageItemsFromLocaleIdentifiers:"), identifiers)
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) LanguageItemsFromLocaleIdentifiersUsingOffline(identifiers objectivec.IObject, offline bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("languageItemsFromLocaleIdentifiers:usingOffline:"), identifiers, offline)
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) PreferredDictationLocaleIdentifierFromAvaiableLocaleIdentifiersDefaultLocaleIdentifier(identifiers objectivec.IObject, identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("preferredDictationLocaleIdentifierFromAvaiableLocaleIdentifiers:defaultLocaleIdentifier:"), identifiers, identifier)
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) SetVisibilityValueForLocaleIdentifierUsingOffline(value uint64, identifier objectivec.IObject, offline bool) {
	objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("setVisibilityValue:forLocaleIdentifier:usingOffline:"), value, identifier, offline)
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) SetVisibleSRLanguageItemsTableUsingOffline(table objectivec.IObject, offline bool) {
	objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("setVisibleSRLanguageItemsTable:usingOffline:"), table, offline)
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) TagNameFromLocaleIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("tagNameFromLocaleIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) VisibiltyValueForLocaleIdentifierEntryExistsUsingOffline(identifier objectivec.IObject, exists *bool, offline bool) uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("visibiltyValueForLocaleIdentifier:entryExists:usingOffline:"), identifier, exists, offline)
	return rv
}
func (_SOSRLanguageItemClass SOSRLanguageItemClass) VisibleSRLanguageItemsTableUsingOffline(offline bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOSRLanguageItemClass.class), objc.Sel("visibleSRLanguageItemsTableUsingOffline:"), offline)
	return objectivec.Object{ID: rv}
}
