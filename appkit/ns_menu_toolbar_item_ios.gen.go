// Code generated from Apple documentation for AppKit. DO NOT EDIT.
// +build ios

package appkit

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/objectivec"
)

// See: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/itemMenu
func (m NSMenuToolbarItem) ItemMenu() objectivec.IObject {
rv := objc.Send[objc.ID](m.ID, objc.Sel("itemMenu"))
return objectivec.Object{ID: rv}
}
func (m NSMenuToolbarItem) SetItemMenu(value objectivec.IObject) {
objc.Send[struct{}](m.ID, objc.Sel("setItemMenu:"), value)
}

