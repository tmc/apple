// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Msgid] class.
var (
	_MsgidClass     MsgidClass
	_MsgidClassOnce sync.Once
)

func getMsgidClass() MsgidClass {
	_MsgidClassOnce.Do(func() {
		_MsgidClass = MsgidClass{class: objc.GetClass("msgid")}
	})
	return _MsgidClass
}

// GetMsgidClass returns the class object for msgid.
func GetMsgidClass() MsgidClass {
	return getMsgidClass()
}

type MsgidClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MsgidClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MsgidClass) Alloc() Msgid {
	rv := objc.Send[Msgid](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/msgid-c.ivar
type Msgid struct {
	objectivec.Object
}

// MsgidFromID constructs a [Msgid] from an objc.ID.
func MsgidFromID(id objc.ID) Msgid {
	return Msgid{objectivec.Object{ID: id}}
}
// Ensure Msgid implements IMsgid.
var _ IMsgid = Msgid{}

// An interface definition for the [Msgid] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSPortMessage/msgid-c.ivar
type IMsgid interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m Msgid) Init() Msgid {
	rv := objc.Send[Msgid](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m Msgid) Autorelease() Msgid {
	rv := objc.Send[Msgid](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMsgid creates a new Msgid instance.
func NewMsgid() Msgid {
	class := getMsgidClass()
	rv := objc.Send[Msgid](objc.ID(class.class), objc.Sel("new"))
	return rv
}

