// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The interface implemented by a Smart Card user interaction delegate to handle user interaction events.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate
type TKSmartCardUserInteractionDelegate interface {
	objectivec.IObject
}

// TKSmartCardUserInteractionDelegateObject wraps an existing Objective-C object that conforms to the TKSmartCardUserInteractionDelegate protocol.
type TKSmartCardUserInteractionDelegateObject struct {
	objectivec.Object
}

func (o TKSmartCardUserInteractionDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// TKSmartCardUserInteractionDelegateObjectFromID constructs a [TKSmartCardUserInteractionDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TKSmartCardUserInteractionDelegateObjectFromID(id objc.ID) TKSmartCardUserInteractionDelegateObject {
	return TKSmartCardUserInteractionDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that a valid character has been entered.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/characterEntered(in:)
func (o TKSmartCardUserInteractionDelegateObject) CharacterEnteredInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("characterEnteredInUserInteraction:"), interaction)
}

// Tells the delegate that a correction key has been pressed.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/correctionKeyPressed(in:)
func (o TKSmartCardUserInteractionDelegateObject) CorrectionKeyPressedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("correctionKeyPressedInUserInteraction:"), interaction)
}

// Tells the delegate that the validation key has been pressed, indicating the
// end of PIN entry.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/validationKeyPressed(in:)
func (o TKSmartCardUserInteractionDelegateObject) ValidationKeyPressedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("validationKeyPressedInUserInteraction:"), interaction)
}

// Tells the delegate that an invalid character has been entered.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/invalidCharacterEntered(in:)
func (o TKSmartCardUserInteractionDelegateObject) InvalidCharacterEnteredInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("invalidCharacterEnteredInUserInteraction:"), interaction)
}

// Tells the delegate that the old PIN needs to be entered.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/oldPINRequested(in:)
func (o TKSmartCardUserInteractionDelegateObject) OldPINRequestedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("oldPINRequestedInUserInteraction:"), interaction)
}

// Tells the delegate that the new PIN needs to be entered.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/newPINRequested(in:)
func (o TKSmartCardUserInteractionDelegateObject) NewPINRequestedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("newPINRequestedInUserInteraction:"), interaction)
}

// Tells the delegate that the new PIN needs to be re-entered for
// confirmation.
//
// interaction: The user interaction.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteractionDelegate/newPINConfirmationRequested(in:)
func (o TKSmartCardUserInteractionDelegateObject) NewPINConfirmationRequestedInUserInteraction(interaction ITKSmartCardUserInteraction) {
	objc.Send[struct{}](o.ID, objc.Sel("newPINConfirmationRequestedInUserInteraction:"), interaction)
}

// TKSmartCardUserInteractionDelegateConfig holds optional typed callbacks for [TKSmartCardUserInteractionDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcarduserinteractiondelegate
type TKSmartCardUserInteractionDelegateConfig struct {

	// Other Methods
	// CharacterEnteredInUserInteraction — Tells the delegate that a valid character has been entered.
	CharacterEnteredInUserInteraction func(interaction TKSmartCardUserInteraction)
	// CorrectionKeyPressedInUserInteraction — Tells the delegate that a correction key has been pressed.
	CorrectionKeyPressedInUserInteraction func(interaction TKSmartCardUserInteraction)
	// ValidationKeyPressedInUserInteraction — Tells the delegate that the validation key has been pressed, indicating the end of PIN entry.
	ValidationKeyPressedInUserInteraction func(interaction TKSmartCardUserInteraction)
	// InvalidCharacterEnteredInUserInteraction — Tells the delegate that an invalid character has been entered.
	InvalidCharacterEnteredInUserInteraction func(interaction TKSmartCardUserInteraction)
	// OldPINRequestedInUserInteraction — Tells the delegate that the old PIN needs to be entered.
	OldPINRequestedInUserInteraction func(interaction TKSmartCardUserInteraction)
	// NewPINRequestedInUserInteraction — Tells the delegate that the new PIN needs to be entered.
	NewPINRequestedInUserInteraction func(interaction TKSmartCardUserInteraction)
	// NewPINConfirmationRequestedInUserInteraction — Tells the delegate that the new PIN needs to be re-entered for confirmation.
	NewPINConfirmationRequestedInUserInteraction func(interaction TKSmartCardUserInteraction)
}

// NewTKSmartCardUserInteractionDelegate creates an Objective-C object implementing the [TKSmartCardUserInteractionDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [TKSmartCardUserInteractionDelegateObject] satisfies the [TKSmartCardUserInteractionDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/cryptotokenkit/tksmartcarduserinteractiondelegate
func NewTKSmartCardUserInteractionDelegate(config TKSmartCardUserInteractionDelegateConfig) TKSmartCardUserInteractionDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoTKSmartCardUserInteractionDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CharacterEnteredInUserInteraction != nil {
		fn := config.CharacterEnteredInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("characterEnteredInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "characterEnteredInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	if config.CorrectionKeyPressedInUserInteraction != nil {
		fn := config.CorrectionKeyPressedInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("correctionKeyPressedInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "correctionKeyPressedInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	if config.ValidationKeyPressedInUserInteraction != nil {
		fn := config.ValidationKeyPressedInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("validationKeyPressedInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "validationKeyPressedInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	if config.InvalidCharacterEnteredInUserInteraction != nil {
		fn := config.InvalidCharacterEnteredInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("invalidCharacterEnteredInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "invalidCharacterEnteredInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	if config.OldPINRequestedInUserInteraction != nil {
		fn := config.OldPINRequestedInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("oldPINRequestedInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "oldPINRequestedInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	if config.NewPINRequestedInUserInteraction != nil {
		fn := config.NewPINRequestedInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("newPINRequestedInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "newPINRequestedInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	if config.NewPINConfirmationRequestedInUserInteraction != nil {
		fn := config.NewPINConfirmationRequestedInUserInteraction
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("newPINConfirmationRequestedInUserInteraction:"),
			Fn: func(self objc.ID, _cmd objc.SEL, interactionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("TKSmartCardUserInteractionDelegate", "newPINConfirmationRequestedInUserInteraction:")
					}
				}()
				interaction := TKSmartCardUserInteractionFromID(interactionID)
				fn(interaction)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("TKSmartCardUserInteractionDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewTKSmartCardUserInteractionDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return TKSmartCardUserInteractionDelegateObjectFromID(instance)
}
