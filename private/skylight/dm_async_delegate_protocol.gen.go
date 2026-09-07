// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// DMAsyncDelegate protocol.
type DMAsyncDelegate interface {
	objectivec.IObject

	// DmAsyncFinishedForDiskMainErrorDetailErrorDictionary protocol.
	DmAsyncFinishedForDiskMainErrorDetailErrorDictionary(disk DADiskRef, mainError int, detailError int, dictionary objectivec.IObject)

	// DmAsyncMessageForDiskStringDictionary protocol.
	DmAsyncMessageForDiskStringDictionary(disk DADiskRef, string_ objectivec.IObject, dictionary objectivec.IObject)

	// DmAsyncProgressForDiskBarberPolePercent protocol.
	DmAsyncProgressForDiskBarberPolePercent(disk DADiskRef, pole bool, percent float32)

	// DmAsyncStartedForDisk protocol.
	DmAsyncStartedForDisk(disk DADiskRef)

	// DmInterruptibilityChanged protocol.
	DmInterruptibilityChanged(changed bool)
}

// DMAsyncDelegateObject wraps an existing Objective-C object that conforms to the DMAsyncDelegate protocol.
type DMAsyncDelegateObject struct {
	objectivec.Object
}

func (o DMAsyncDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// DMAsyncDelegateObjectFromID constructs a [DMAsyncDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func DMAsyncDelegateObjectFromID(id objc.ID) DMAsyncDelegateObject {
	return DMAsyncDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o DMAsyncDelegateObject) DmAsyncFinishedForDiskMainErrorDetailErrorDictionary(disk DADiskRef, mainError int, detailError int, dictionary objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("dmAsyncFinishedForDisk:mainError:detailError:dictionary:"), disk, mainError, detailError, dictionary)
}
func (o DMAsyncDelegateObject) DmAsyncMessageForDiskStringDictionary(disk DADiskRef, string_ objectivec.IObject, dictionary objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("dmAsyncMessageForDisk:string:dictionary:"), disk, string_, dictionary)
}
func (o DMAsyncDelegateObject) DmAsyncProgressForDiskBarberPolePercent(disk DADiskRef, pole bool, percent float32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("dmAsyncProgressForDisk:barberPole:percent:"), disk, pole, percent)
}
func (o DMAsyncDelegateObject) DmAsyncStartedForDisk(disk DADiskRef) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("dmAsyncStartedForDisk:"), disk)
}
func (o DMAsyncDelegateObject) DmInterruptibilityChanged(changed bool) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("dmInterruptibilityChanged:"), changed)
}

// DMAsyncDelegateConfig holds optional typed callbacks for [DMAsyncDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
type DMAsyncDelegateConfig struct {

	// Other Methods
	DmAsyncProgressForDiskBarberPolePercent func(disk DADiskRef, pole bool, percent float32)
	DmAsyncStartedForDisk                   func(disk DADiskRef)
	DmInterruptibilityChanged               func(changed bool)
}

// NewDMAsyncDelegate creates an Objective-C object implementing the [DMAsyncDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [DMAsyncDelegateObject] satisfies the [DMAsyncDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
func NewDMAsyncDelegate(config DMAsyncDelegateConfig) DMAsyncDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoDMAsyncDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DmAsyncProgressForDiskBarberPolePercent != nil {
		fn := config.DmAsyncProgressForDiskBarberPolePercent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("dmAsyncProgressForDisk:barberPole:percent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, disk DADiskRef, pole bool, percent float32) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("DMAsyncDelegate", "dmAsyncProgressForDisk:barberPole:percent:")
					}
				}()
				fn(disk, pole, percent)
				_delegateDone = true
			},
		})
	}

	if config.DmAsyncStartedForDisk != nil {
		fn := config.DmAsyncStartedForDisk
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("dmAsyncStartedForDisk:"),
			Fn: func(self objc.ID, _cmd objc.SEL, disk DADiskRef) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("DMAsyncDelegate", "dmAsyncStartedForDisk:")
					}
				}()
				fn(disk)
				_delegateDone = true
			},
		})
	}

	if config.DmInterruptibilityChanged != nil {
		fn := config.DmInterruptibilityChanged
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("dmInterruptibilityChanged:"),
			Fn: func(self objc.ID, _cmd objc.SEL, changed bool) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("DMAsyncDelegate", "dmInterruptibilityChanged:")
					}
				}()
				fn(changed)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("DMAsyncDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewDMAsyncDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return DMAsyncDelegateObjectFromID(instance)
}
