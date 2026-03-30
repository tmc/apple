// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// OSSystemExtensionsWorkspaceObserver protocol.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspaceObserver
type OSSystemExtensionsWorkspaceObserver interface {
	objectivec.IObject
}

// OSSystemExtensionsWorkspaceObserverObject wraps an existing Objective-C object that conforms to the OSSystemExtensionsWorkspaceObserver protocol.
type OSSystemExtensionsWorkspaceObserverObject struct {
	objectivec.Object
}

func (o OSSystemExtensionsWorkspaceObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// OSSystemExtensionsWorkspaceObserverObjectFromID constructs a [OSSystemExtensionsWorkspaceObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func OSSystemExtensionsWorkspaceObserverObjectFromID(id objc.ID) OSSystemExtensionsWorkspaceObserverObject {
	return OSSystemExtensionsWorkspaceObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// # Discussion
//
// This delegate method will be called when the user disables an already
// enabled system extension, or when the system extension is first installed
// and is in the disabled state.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspaceObserver/systemExtensionWillBecomeDisabled(_:)
func (o OSSystemExtensionsWorkspaceObserverObject) SystemExtensionWillBecomeDisabled(systemExtensionInfo IOSSystemExtensionInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("systemExtensionWillBecomeDisabled:"), systemExtensionInfo)
}

// # Discussion
//
// This delegate method will be called when a system extension has been
// validated and allowed by the user to run.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspaceObserver/systemExtensionWillBecomeEnabled(_:)
func (o OSSystemExtensionsWorkspaceObserverObject) SystemExtensionWillBecomeEnabled(systemExtensionInfo IOSSystemExtensionInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("systemExtensionWillBecomeEnabled:"), systemExtensionInfo)
}

// # Discussion
//
// This delegate method will be called when a system extension is deactivated
// and is about to get uninstalled. The extension may still be running until
// the system is rebooted.
//
// See: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspaceObserver/systemExtensionWillBecomeInactive(_:)
func (o OSSystemExtensionsWorkspaceObserverObject) SystemExtensionWillBecomeInactive(systemExtensionInfo IOSSystemExtensionInfo) {
	objc.Send[struct{}](o.ID, objc.Sel("systemExtensionWillBecomeInactive:"), systemExtensionInfo)
}
