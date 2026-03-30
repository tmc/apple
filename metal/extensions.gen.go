// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// NewSharedEventWithMachPort is a private SPI used for IOSurfaceSharedEvent
// interop. It may be unavailable on some runtimes.
func (o MTLDeviceObject) NewSharedEventWithMachPort(port uint32) MTLSharedEvent {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedEventWithMachPort:"), port)
	return MTLSharedEventObjectFromID(rv)
}

// SharedEventNotificationBinding keeps a notify block alive.
type SharedEventNotificationBinding struct {
	cleanup func()
	once    sync.Once
}

// Release releases the underlying Objective-C block.
func (b *SharedEventNotificationBinding) Release() {
	if b == nil {
		return
	}
	b.once.Do(func() {
		if b.cleanup != nil {
			b.cleanup()
		}
	})
}

// NotifyListenerAtValueRetained registers a notification block and returns a
// binding that must be released when notifications are no longer needed.
func (o MTLSharedEventObject) NotifyListenerAtValueRetained(listener IMTLSharedEventListener, value uint64, block MTLSharedEventNotificationBlock) *SharedEventNotificationBinding {
	if o.ID == 0 || block == nil {
		return nil
	}
	b := objc.NewBlock(func(_ objc.Block, event objc.ID, signaledValue uint64) {
		block(MTLSharedEventObjectFromID(event), signaledValue)
	})
	objc.Send[struct{}](o.ID, objc.Sel("notifyListener:atValue:block:"), listener, value, objc.ID(b))
	return &SharedEventNotificationBinding{cleanup: b.Release}
}
