
// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "testing"

func TestSharedEventNotificationBindingReleaseNil(t *testing.T) {
	var b *SharedEventNotificationBinding
	b.Release()
}

func TestSharedEventNotificationBindingReleaseIdempotent(t *testing.T) {
	b := &SharedEventNotificationBinding{}
	b.Release()
	b.Release()
}

func TestNotifyListenerAtValueRetainedNilReceiver(t *testing.T) {
	var ev MTLSharedEventObject
	b := ev.NotifyListenerAtValueRetained(nil, 1, func(_ MTLSharedEvent, _ uint64) {})
	if b != nil {
		t.Fatalf("NotifyListenerAtValueRetained() on nil receiver = %v, want nil", b)
	}
}

func TestMetalExtensionWrapperCompile(t *testing.T) {
	var _ func(MTLDeviceObject, uint32) MTLSharedEvent = MTLDeviceObject.NewSharedEventWithMachPort
	var _ func(MTLSharedEventObject, IMTLSharedEventListener, uint64, MTLSharedEventNotificationBlock) *SharedEventNotificationBinding =
		MTLSharedEventObject.NotifyListenerAtValueRetained
}

