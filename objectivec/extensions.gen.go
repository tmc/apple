// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import "github.com/tmc/apple/objc"

// Retain increments the reference count of the object.
func (o Object) Retain() {
	if o.ID != 0 {
		objc.Send[objc.ID](o.ID, objc.Sel("retain"))
	}
}

// Release decrements the reference count of the object.
func (o Object) Release() {
	if o.ID != 0 {
		objc.Send[objc.ID](o.ID, objc.Sel("release"))
	}
}

// Autorelease adds the object to the current autorelease pool.
func (o Object) Autorelease() {
	if o.ID != 0 {
		objc.Send[objc.ID](o.ID, objc.Sel("autorelease"))
	}
}


