// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSEventAuthenticationMessageVersionedPID] class.
var (
	_SLSEventAuthenticationMessageVersionedPIDClass     SLSEventAuthenticationMessageVersionedPIDClass
	_SLSEventAuthenticationMessageVersionedPIDClassOnce sync.Once
)

func getSLSEventAuthenticationMessageVersionedPIDClass() SLSEventAuthenticationMessageVersionedPIDClass {
	_SLSEventAuthenticationMessageVersionedPIDClassOnce.Do(func() {
		_SLSEventAuthenticationMessageVersionedPIDClass = SLSEventAuthenticationMessageVersionedPIDClass{class: objc.GetClass("SLSEventAuthenticationMessageVersionedPID")}
	})
	return _SLSEventAuthenticationMessageVersionedPIDClass
}

// GetSLSEventAuthenticationMessageVersionedPIDClass returns the class object for SLSEventAuthenticationMessageVersionedPID.
func GetSLSEventAuthenticationMessageVersionedPIDClass() SLSEventAuthenticationMessageVersionedPIDClass {
	return getSLSEventAuthenticationMessageVersionedPIDClass()
}

type SLSEventAuthenticationMessageVersionedPIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSEventAuthenticationMessageVersionedPIDClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSEventAuthenticationMessageVersionedPIDClass) Alloc() SLSEventAuthenticationMessageVersionedPID {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageVersionedPID](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSEventAuthenticationMessageVersionedPID.AddToSigningContext]
//   - [SLSEventAuthenticationMessageVersionedPID.EncodeWithCoder]
//   - [SLSEventAuthenticationMessageVersionedPID.Pid]
//   - [SLSEventAuthenticationMessageVersionedPID.Token]
//   - [SLSEventAuthenticationMessageVersionedPID.InitWithCoder]
//   - [SLSEventAuthenticationMessageVersionedPID.InitWithPIDVersion]
//   - [SLSEventAuthenticationMessageVersionedPID.Version]
type SLSEventAuthenticationMessageVersionedPID struct {
	objectivec.Object
}

// SLSEventAuthenticationMessageVersionedPIDFromID constructs a [SLSEventAuthenticationMessageVersionedPID] from an objc.ID.
func SLSEventAuthenticationMessageVersionedPIDFromID(id objc.ID) SLSEventAuthenticationMessageVersionedPID {
	return SLSEventAuthenticationMessageVersionedPID{objectivec.Object{ID: id}}
}

// Ensure SLSEventAuthenticationMessageVersionedPID implements ISLSEventAuthenticationMessageVersionedPID.
var _ ISLSEventAuthenticationMessageVersionedPID = SLSEventAuthenticationMessageVersionedPID{}

// An interface definition for the [SLSEventAuthenticationMessageVersionedPID] class.
//
// # Methods
//
//   - [ISLSEventAuthenticationMessageVersionedPID.AddToSigningContext]
//   - [ISLSEventAuthenticationMessageVersionedPID.EncodeWithCoder]
//   - [ISLSEventAuthenticationMessageVersionedPID.Pid]
//   - [ISLSEventAuthenticationMessageVersionedPID.Token]
//   - [ISLSEventAuthenticationMessageVersionedPID.InitWithCoder]
//   - [ISLSEventAuthenticationMessageVersionedPID.InitWithPIDVersion]
//   - [ISLSEventAuthenticationMessageVersionedPID.Version]
type ISLSEventAuthenticationMessageVersionedPID interface {
	objectivec.IObject

	// Topic: Methods

	AddToSigningContext(context objectivec.IObject)
	EncodeWithCoder(coder foundation.INSCoder)
	Pid() int
	Token() uint64
	InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessageVersionedPID
	InitWithPIDVersion(pid int, version uint32) SLSEventAuthenticationMessageVersionedPID
	Version() uint32
}

// Init initializes the instance.
func (s SLSEventAuthenticationMessageVersionedPID) Init() SLSEventAuthenticationMessageVersionedPID {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageVersionedPID](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSEventAuthenticationMessageVersionedPID) Autorelease() SLSEventAuthenticationMessageVersionedPID {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageVersionedPID](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSEventAuthenticationMessageVersionedPID creates a new SLSEventAuthenticationMessageVersionedPID instance.
func NewSLSEventAuthenticationMessageVersionedPID() SLSEventAuthenticationMessageVersionedPID {
	class := getSLSEventAuthenticationMessageVersionedPIDClass()
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageVersionedPID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSEventAuthenticationMessageVersionedPIDWithCoder(coder objectivec.IObject) SLSEventAuthenticationMessageVersionedPID {
	instance := getSLSEventAuthenticationMessageVersionedPIDClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSEventAuthenticationMessageVersionedPIDFromID(rv)
}

func NewSLSEventAuthenticationMessageVersionedPIDWithPIDVersion(pid int, version uint32) SLSEventAuthenticationMessageVersionedPID {
	instance := getSLSEventAuthenticationMessageVersionedPIDClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithPID:version:"), pid, version)
	return SLSEventAuthenticationMessageVersionedPIDFromID(rv)
}

func (s SLSEventAuthenticationMessageVersionedPID) AddToSigningContext(context objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("addToSigningContext:"), context)
}
func (s SLSEventAuthenticationMessageVersionedPID) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (s SLSEventAuthenticationMessageVersionedPID) InitWithCoder(coder foundation.INSCoder) SLSEventAuthenticationMessageVersionedPID {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageVersionedPID](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s SLSEventAuthenticationMessageVersionedPID) InitWithPIDVersion(pid int, version uint32) SLSEventAuthenticationMessageVersionedPID {
	rv := objc.SendIfResponds[SLSEventAuthenticationMessageVersionedPID](s.ID, objc.Sel("initWithPID:version:"), pid, version)
	return rv
}

func (_SLSEventAuthenticationMessageVersionedPIDClass SLSEventAuthenticationMessageVersionedPIDClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_SLSEventAuthenticationMessageVersionedPIDClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (s SLSEventAuthenticationMessageVersionedPID) Pid() int {
	rv := objc.SendIfResponds[int](s.ID, objc.Sel("pid"))
	return rv
}
func (s SLSEventAuthenticationMessageVersionedPID) Token() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("token"))
	return rv
}
func (s SLSEventAuthenticationMessageVersionedPID) Version() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("version"))
	return rv
}
