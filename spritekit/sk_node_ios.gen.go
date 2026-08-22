// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.
//go:build ios
// +build ios

package spritekit

import (
	"github.com/tmc/apple/objc"
)

// The focus behavior for a node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/focusBehavior
func (n SKNode) FocusBehavior() SKNodeFocusBehavior {
	rv := objc.Send[SKNodeFocusBehavior](n.ID, objc.Sel("focusBehavior"))
	return SKNodeFocusBehavior(rv)
}
func (n SKNode) SetFocusBehavior(value SKNodeFocusBehavior) {
	objc.Send[struct{}](n.ID, objc.Sel("setFocusBehavior:"), value)
}
