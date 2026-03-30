// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VoiceTableRow] class.
var (
	_VoiceTableRowClass     VoiceTableRowClass
	_VoiceTableRowClassOnce sync.Once
)

func getVoiceTableRowClass() VoiceTableRowClass {
	_VoiceTableRowClassOnce.Do(func() {
		_VoiceTableRowClass = VoiceTableRowClass{class: objc.GetClass("VoiceTableRow")}
	})
	return _VoiceTableRowClass
}

// GetVoiceTableRowClass returns the class object for VoiceTableRow.
func GetVoiceTableRowClass() VoiceTableRowClass {
	return getVoiceTableRowClass()
}

type VoiceTableRowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VoiceTableRowClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VoiceTableRowClass) Alloc() VoiceTableRow {
	rv := objc.Send[VoiceTableRow](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VoiceTableRow.IsGroupRow]
//   - [VoiceTableRow.IsSelected]
//   - [VoiceTableRow.SetSelectedUsingLanguageCodeOnlyToSelectVoices]
//   - [VoiceTableRow.TableIndex]
//   - [VoiceTableRow.SetTableIndex]
//   - [VoiceTableRow.Title]
//   - [VoiceTableRow.VoiceObject]
//   - [VoiceTableRow.InitWithTitleVoiceObject]
//   - [VoiceTableRow.GroupRow]
//   - [VoiceTableRow.SetGroupRow]
//   - [VoiceTableRow.Selected]
//   - [VoiceTableRow.SetSelected]
//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow
type VoiceTableRow struct {
	objectivec.Object
}

// VoiceTableRowFromID constructs a [VoiceTableRow] from an objc.ID.
func VoiceTableRowFromID(id objc.ID) VoiceTableRow {
	return VoiceTableRow{objectivec.Object{ID: id}}
}

// Ensure VoiceTableRow implements IVoiceTableRow.
var _ IVoiceTableRow = VoiceTableRow{}

// An interface definition for the [VoiceTableRow] class.
//
// # Methods
//
//   - [IVoiceTableRow.IsGroupRow]
//   - [IVoiceTableRow.IsSelected]
//   - [IVoiceTableRow.SetSelectedUsingLanguageCodeOnlyToSelectVoices]
//   - [IVoiceTableRow.TableIndex]
//   - [IVoiceTableRow.SetTableIndex]
//   - [IVoiceTableRow.Title]
//   - [IVoiceTableRow.VoiceObject]
//   - [IVoiceTableRow.InitWithTitleVoiceObject]
//   - [IVoiceTableRow.GroupRow]
//   - [IVoiceTableRow.SetGroupRow]
//   - [IVoiceTableRow.Selected]
//   - [IVoiceTableRow.SetSelected]
//
// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow
type IVoiceTableRow interface {
	objectivec.IObject

	// Topic: Methods

	IsGroupRow() bool
	IsSelected() bool
	SetSelectedUsingLanguageCodeOnlyToSelectVoices(voices bool)
	TableIndex() uint64
	SetTableIndex(value uint64)
	Title() string
	VoiceObject() ISOVoiceObject
	InitWithTitleVoiceObject(title objectivec.IObject, object objectivec.IObject) VoiceTableRow
	GroupRow() bool
	SetGroupRow(value bool)
	Selected() bool
	SetSelected(value bool)
}

// Init initializes the instance.
func (v VoiceTableRow) Init() VoiceTableRow {
	rv := objc.Send[VoiceTableRow](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceTableRow) Autorelease() VoiceTableRow {
	rv := objc.Send[VoiceTableRow](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceTableRow creates a new VoiceTableRow instance.
func NewVoiceTableRow() VoiceTableRow {
	class := getVoiceTableRowClass()
	rv := objc.Send[VoiceTableRow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/initWithTitle:voiceObject:
func NewVoiceTableRowWithTitleVoiceObject(title objectivec.IObject, object objectivec.IObject) VoiceTableRow {
	instance := getVoiceTableRowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:voiceObject:"), title, object)
	return VoiceTableRowFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/isGroupRow
func (v VoiceTableRow) IsGroupRow() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isGroupRow"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/isSelected
func (v VoiceTableRow) IsSelected() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isSelected"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/setSelectedUsingLanguageCodeOnlyToSelectVoices:
func (v VoiceTableRow) SetSelectedUsingLanguageCodeOnlyToSelectVoices(voices bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("setSelectedUsingLanguageCodeOnlyToSelectVoices:"), voices)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/initWithTitle:voiceObject:
func (v VoiceTableRow) InitWithTitleVoiceObject(title objectivec.IObject, object objectivec.IObject) VoiceTableRow {
	rv := objc.Send[VoiceTableRow](v.ID, objc.Sel("initWithTitle:voiceObject:"), title, object)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/arrangedRowsFromVoiceObjects:useLanguageCodeOnlyToSelectVoices:showIndividualQualities:showFullGroupNames:showCurrentLocaleAtTop:
func (_VoiceTableRowClass VoiceTableRowClass) ArrangedRowsFromVoiceObjectsUseLanguageCodeOnlyToSelectVoicesShowIndividualQualitiesShowFullGroupNamesShowCurrentLocaleAtTop(objects objectivec.IObject, voices bool, qualities bool, names bool, top bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_VoiceTableRowClass.class), objc.Sel("arrangedRowsFromVoiceObjects:useLanguageCodeOnlyToSelectVoices:showIndividualQualities:showFullGroupNames:showCurrentLocaleAtTop:"), objects, voices, qualities, names, top)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/groupRow
func (v VoiceTableRow) GroupRow() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("groupRow"))
	return rv
}
func (v VoiceTableRow) SetGroupRow(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setGroupRow:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/selected
func (v VoiceTableRow) Selected() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("selected"))
	return rv
}
func (v VoiceTableRow) SetSelected(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSelected:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/tableIndex
func (v VoiceTableRow) TableIndex() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("tableIndex"))
	return rv
}
func (v VoiceTableRow) SetTableIndex(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setTableIndex:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/title
func (v VoiceTableRow) Title() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/VoiceTableRow/voiceObject
func (v VoiceTableRow) VoiceObject() ISOVoiceObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("voiceObject"))
	return SOVoiceObjectFromID(objc.ID(rv))
}
