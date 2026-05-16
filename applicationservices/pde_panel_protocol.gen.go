// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

package applicationservices

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// PDEPanel protocol.
//
// See: https://developer.apple.com/documentation/applicationservices/pdepanel
type PDEPanel interface {
	objectivec.IObject
}

// PDEPanelObject wraps an existing Objective-C object that conforms to the PDEPanel protocol.
type PDEPanelObject struct {
	objectivec.Object
}

func (o PDEPanelObject) BaseObject() objectivec.Object {
	return o.Object
}

// PDEPanelObjectFromID constructs a [PDEPanelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func PDEPanelObjectFromID(id objc.ID) PDEPanelObject {
	return PDEPanelObject{
		Object: objectivec.ObjectFromID(id),
	}
}
