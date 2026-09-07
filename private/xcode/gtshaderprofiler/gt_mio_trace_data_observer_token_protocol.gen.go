// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioTraceDataObserverToken protocol.
type GTMioTraceDataObserverToken interface {
	objectivec.IObject

	// Cancel protocol.
	Cancel()
}

// GTMioTraceDataObserverTokenObject wraps an existing Objective-C object that conforms to the GTMioTraceDataObserverToken protocol.
type GTMioTraceDataObserverTokenObject struct {
	objectivec.Object
}

func (o GTMioTraceDataObserverTokenObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioTraceDataObserverTokenObjectFromID constructs a [GTMioTraceDataObserverTokenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioTraceDataObserverTokenObjectFromID(id objc.ID) GTMioTraceDataObserverTokenObject {
	return GTMioTraceDataObserverTokenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTMioTraceDataObserverTokenObject) Cancel() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("cancel"))
}
