// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DISLAFrontend] class.
var (
	_DISLAFrontendClass     DISLAFrontendClass
	_DISLAFrontendClassOnce sync.Once
)

func getDISLAFrontendClass() DISLAFrontendClass {
	_DISLAFrontendClassOnce.Do(func() {
		_DISLAFrontendClass = DISLAFrontendClass{class: objc.GetClass("DISLAFrontend")}
	})
	return _DISLAFrontendClass
}

// GetDISLAFrontendClass returns the class object for DISLAFrontend.
func GetDISLAFrontendClass() DISLAFrontendClass {
	return getDISLAFrontendClass()
}

type DISLAFrontendClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DISLAFrontendClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DISLAFrontendClass) Alloc() DISLAFrontend {
	rv := objc.SendIfResponds[DISLAFrontend](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DISLAFrontend struct {
	objectivec.Object
}

// DISLAFrontendFromID constructs a [DISLAFrontend] from an objc.ID.
func DISLAFrontendFromID(id objc.ID) DISLAFrontend {
	return DISLAFrontend{objectivec.Object{ID: id}}
}

// Ensure DISLAFrontend implements IDISLAFrontend.
var _ IDISLAFrontend = DISLAFrontend{}

// An interface definition for the [DISLAFrontend] class.
type IDISLAFrontend interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DISLAFrontend) Init() DISLAFrontend {
	rv := objc.SendIfResponds[DISLAFrontend](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DISLAFrontend) Autorelease() DISLAFrontend {
	rv := objc.SendIfResponds[DISLAFrontend](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDISLAFrontend creates a new DISLAFrontend instance.
func NewDISLAFrontend() DISLAFrontend {
	class := getDISLAFrontendClass()
	rv := objc.SendIfResponds[DISLAFrontend](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_DISLAFrontendClass DISLAFrontendClass) DisplaySLATextError(sLAText string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DISLAFrontendClass.class), objc.Sel("displaySLAText:error:"), unsafe.Pointer(unsafe.StringData(sLAText+"\x00")), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("displaySLAText:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DISLAFrontendClass DISLAFrontendClass) IsStdoutQuietMode() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DISLAFrontendClass.class), objc.Sel("isStdoutQuietMode"))
	return rv
}
func (_DISLAFrontendClass DISLAFrontendClass) PromptForAcceptance(acceptance []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_DISLAFrontendClass.class), objc.Sel("promptForAcceptance:"), objectivec.IObjectSliceToNSArray(acceptance))
	return rv
}
func (_DISLAFrontendClass DISLAFrontendClass) PromptUserForSLAAcceptanceError(sLAAcceptance objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DISLAFrontendClass.class), objc.Sel("promptUserForSLAAcceptance:error:"), sLAAcceptance, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("promptUserForSLAAcceptance:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_DISLAFrontendClass DISLAFrontendClass) RedirectStdoutToTTY() int32 {
	rv := objc.SendIfResponds[int32](objc.ID(_DISLAFrontendClass.class), objc.Sel("redirectStdoutToTTY"))
	return rv
}
func (_DISLAFrontendClass DISLAFrontendClass) RestoreStdout(stdout int32) {
	objc.SendIfResponds[objc.ID](objc.ID(_DISLAFrontendClass.class), objc.Sel("restoreStdout:"), stdout)
}
