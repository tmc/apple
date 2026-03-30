// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSRLanguageRow] class.
var (
	_SOSRLanguageRowClass     SOSRLanguageRowClass
	_SOSRLanguageRowClassOnce sync.Once
)

func getSOSRLanguageRowClass() SOSRLanguageRowClass {
	_SOSRLanguageRowClassOnce.Do(func() {
		_SOSRLanguageRowClass = SOSRLanguageRowClass{class: objc.GetClass("SOSRLanguageRow")}
	})
	return _SOSRLanguageRowClass
}

// GetSOSRLanguageRowClass returns the class object for SOSRLanguageRow.
func GetSOSRLanguageRowClass() SOSRLanguageRowClass {
	return getSOSRLanguageRowClass()
}

type SOSRLanguageRowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSRLanguageRowClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSRLanguageRowClass) Alloc() SOSRLanguageRow {
	rv := objc.Send[SOSRLanguageRow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSRLanguageRow.IsGroupRow]
//   - [SOSRLanguageRow.IsSelected]
//   - [SOSRLanguageRow.SetSelectedUsingLanguageCodeOnlyToSelectItems]
//   - [SOSRLanguageRow.SrLanguageItem]
//   - [SOSRLanguageRow.TableIndex]
//   - [SOSRLanguageRow.SetTableIndex]
//   - [SOSRLanguageRow.Title]
//   - [SOSRLanguageRow.InitWithTitleSrLanguageItem]
//   - [SOSRLanguageRow.GroupRow]
//   - [SOSRLanguageRow.SetGroupRow]
//   - [SOSRLanguageRow.Selected]
//   - [SOSRLanguageRow.SetSelected]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow
type SOSRLanguageRow struct {
	objectivec.Object
}

// SOSRLanguageRowFromID constructs a [SOSRLanguageRow] from an objc.ID.
func SOSRLanguageRowFromID(id objc.ID) SOSRLanguageRow {
	return SOSRLanguageRow{objectivec.Object{ID: id}}
}

// Ensure SOSRLanguageRow implements ISOSRLanguageRow.
var _ ISOSRLanguageRow = SOSRLanguageRow{}

// An interface definition for the [SOSRLanguageRow] class.
//
// # Methods
//
//   - [ISOSRLanguageRow.IsGroupRow]
//   - [ISOSRLanguageRow.IsSelected]
//   - [ISOSRLanguageRow.SetSelectedUsingLanguageCodeOnlyToSelectItems]
//   - [ISOSRLanguageRow.SrLanguageItem]
//   - [ISOSRLanguageRow.TableIndex]
//   - [ISOSRLanguageRow.SetTableIndex]
//   - [ISOSRLanguageRow.Title]
//   - [ISOSRLanguageRow.InitWithTitleSrLanguageItem]
//   - [ISOSRLanguageRow.GroupRow]
//   - [ISOSRLanguageRow.SetGroupRow]
//   - [ISOSRLanguageRow.Selected]
//   - [ISOSRLanguageRow.SetSelected]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow
type ISOSRLanguageRow interface {
	objectivec.IObject

	// Topic: Methods

	IsGroupRow() bool
	IsSelected() bool
	SetSelectedUsingLanguageCodeOnlyToSelectItems(items bool)
	SrLanguageItem() ISOSRLanguageItem
	TableIndex() uint64
	SetTableIndex(value uint64)
	Title() string
	InitWithTitleSrLanguageItem(title objectivec.IObject, item objectivec.IObject) SOSRLanguageRow
	GroupRow() bool
	SetGroupRow(value bool)
	Selected() bool
	SetSelected(value bool)
}

// Init initializes the instance.
func (s SOSRLanguageRow) Init() SOSRLanguageRow {
	rv := objc.Send[SOSRLanguageRow](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSRLanguageRow) Autorelease() SOSRLanguageRow {
	rv := objc.Send[SOSRLanguageRow](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSRLanguageRow creates a new SOSRLanguageRow instance.
func NewSOSRLanguageRow() SOSRLanguageRow {
	class := getSOSRLanguageRowClass()
	rv := objc.Send[SOSRLanguageRow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/initWithTitle:srLanguageItem:
func NewSOSRLanguageRowWithTitleSrLanguageItem(title objectivec.IObject, item objectivec.IObject) SOSRLanguageRow {
	instance := getSOSRLanguageRowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:srLanguageItem:"), title, item)
	return SOSRLanguageRowFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/isGroupRow
func (s SOSRLanguageRow) IsGroupRow() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isGroupRow"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/isSelected
func (s SOSRLanguageRow) IsSelected() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSelected"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/setSelectedUsingLanguageCodeOnlyToSelectItems:
func (s SOSRLanguageRow) SetSelectedUsingLanguageCodeOnlyToSelectItems(items bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSelectedUsingLanguageCodeOnlyToSelectItems:"), items)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/initWithTitle:srLanguageItem:
func (s SOSRLanguageRow) InitWithTitleSrLanguageItem(title objectivec.IObject, item objectivec.IObject) SOSRLanguageRow {
	rv := objc.Send[SOSRLanguageRow](s.ID, objc.Sel("initWithTitle:srLanguageItem:"), title, item)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/arrangedRowsFromSRLanguageItems:groupCountries:useLanguageCodeOnlyToSelectItems:showCurrentLocaleAtTop:
func (_SOSRLanguageRowClass SOSRLanguageRowClass) ArrangedRowsFromSRLanguageItemsGroupCountriesUseLanguageCodeOnlyToSelectItemsShowCurrentLocaleAtTop(items objectivec.IObject, countries bool, items2 bool, top bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSRLanguageRowClass.class), objc.Sel("arrangedRowsFromSRLanguageItems:groupCountries:useLanguageCodeOnlyToSelectItems:showCurrentLocaleAtTop:"), items, countries, items2, top)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/groupRow
func (s SOSRLanguageRow) GroupRow() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("groupRow"))
	return rv
}
func (s SOSRLanguageRow) SetGroupRow(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setGroupRow:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/selected
func (s SOSRLanguageRow) Selected() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("selected"))
	return rv
}
func (s SOSRLanguageRow) SetSelected(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelected:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/srLanguageItem
func (s SOSRLanguageRow) SrLanguageItem() ISOSRLanguageItem {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("srLanguageItem"))
	return SOSRLanguageItemFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/tableIndex
func (s SOSRLanguageRow) TableIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("tableIndex"))
	return rv
}
func (s SOSRLanguageRow) SetTableIndex(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setTableIndex:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSRLanguageRow/title
func (s SOSRLanguageRow) Title() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
