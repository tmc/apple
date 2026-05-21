// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NWPrettyDescription protocol.
type NWPrettyDescription interface {
	objectivec.IObject
}

// NWPrettyDescriptionObject wraps an existing Objective-C object that conforms to the NWPrettyDescription protocol.
type NWPrettyDescriptionObject struct {
	objectivec.Object
}

func (o NWPrettyDescriptionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NWPrettyDescriptionObjectFromID constructs a [NWPrettyDescriptionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NWPrettyDescriptionObjectFromID(id objc.ID) NWPrettyDescriptionObject {
	return NWPrettyDescriptionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o NWPrettyDescriptionObject) DescriptionWithIndentShowFullContent(indent int, content bool) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("descriptionWithIndent:showFullContent:"), indent, content)
	return objectivec.Object{ID: rv}
}
