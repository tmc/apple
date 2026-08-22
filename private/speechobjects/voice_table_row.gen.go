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
	rv := objc.SendIfResponds[VoiceTableRow](objc.ID(vc.class), objc.Sel("alloc"))
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
	rv := objc.SendIfResponds[VoiceTableRow](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VoiceTableRow) Autorelease() VoiceTableRow {
	rv := objc.SendIfResponds[VoiceTableRow](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVoiceTableRow creates a new VoiceTableRow instance.
func NewVoiceTableRow() VoiceTableRow {
	class := getVoiceTableRowClass()
	rv := objc.SendIfResponds[VoiceTableRow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVoiceTableRowWithTitleVoiceObject(title objectivec.IObject, object objectivec.IObject) VoiceTableRow {
	instance := getVoiceTableRowClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTitle:voiceObject:"), title, object)
	return VoiceTableRowFromID(rv)
}

func (v VoiceTableRow) IsGroupRow() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isGroupRow"))
	return rv
}
func (v VoiceTableRow) IsSelected() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isSelected"))
	return rv
}
func (v VoiceTableRow) SetSelectedUsingLanguageCodeOnlyToSelectVoices(voices bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("setSelectedUsingLanguageCodeOnlyToSelectVoices:"), voices)
}
func (v VoiceTableRow) InitWithTitleVoiceObject(title objectivec.IObject, object objectivec.IObject) VoiceTableRow {
	rv := objc.SendIfResponds[VoiceTableRow](v.ID, objc.Sel("initWithTitle:voiceObject:"), title, object)
	return rv
}

func (_VoiceTableRowClass VoiceTableRowClass) ArrangedRowsFromVoiceObjectsUseLanguageCodeOnlyToSelectVoicesShowIndividualQualitiesShowFullGroupNamesShowCurrentLocaleAtTop(objects objectivec.IObject, voices bool, qualities bool, names bool, top bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_VoiceTableRowClass.class), objc.Sel("arrangedRowsFromVoiceObjects:useLanguageCodeOnlyToSelectVoices:showIndividualQualities:showFullGroupNames:showCurrentLocaleAtTop:"), objects, voices, qualities, names, top)
	return objectivec.Object{ID: rv}
}

func (v VoiceTableRow) GroupRow() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("groupRow"))
	return rv
}
func (v VoiceTableRow) SetGroupRow(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setGroupRow:"), value)
}
func (v VoiceTableRow) Selected() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("selected"))
	return rv
}
func (v VoiceTableRow) SetSelected(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setSelected:"), value)
}
func (v VoiceTableRow) TableIndex() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("tableIndex"))
	return rv
}
func (v VoiceTableRow) SetTableIndex(value uint64) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setTableIndex:"), value)
}
func (v VoiceTableRow) Title() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (v VoiceTableRow) VoiceObject() ISOVoiceObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("voiceObject"))
	return SOVoiceObjectFromID(objc.ID(rv))
}
