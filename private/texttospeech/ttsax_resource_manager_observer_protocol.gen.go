// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// TTSAXResourceManagerObserver protocol.
type TTSAXResourceManagerObserver interface {
	objectivec.IObject

	// DownloadProgressForVoiceIdProgressStorageSizeRequiredDiskSpace protocol.
	DownloadProgressForVoiceIdProgressStorageSizeRequiredDiskSpace(id objectivec.IObject, progress float32, size int64, space bool)

	// FinishedDeletingResource protocol.
	FinishedDeletingResource(resource objectivec.IObject)

	// FinishedDownloadingResourceWasCancelled protocol.
	FinishedDownloadingResourceWasCancelled(resource objectivec.IObject, cancelled bool)

	// ResourceCacheDidReceiveUpdate protocol.
	ResourceCacheDidReceiveUpdate()
}

// TTSAXResourceManagerObserverObject wraps an existing Objective-C object that conforms to the TTSAXResourceManagerObserver protocol.
type TTSAXResourceManagerObserverObject struct {
	objectivec.Object
}

func (o TTSAXResourceManagerObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// TTSAXResourceManagerObserverObjectFromID constructs a [TTSAXResourceManagerObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func TTSAXResourceManagerObserverObjectFromID(id objc.ID) TTSAXResourceManagerObserverObject {
	return TTSAXResourceManagerObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o TTSAXResourceManagerObserverObject) DownloadProgressForVoiceIdProgressStorageSizeRequiredDiskSpace(id objectivec.IObject, progress float32, size int64, space bool) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("downloadProgressForVoiceId:progress:storageSize:requiredDiskSpace:"), id, progress, size, space)
}
func (o TTSAXResourceManagerObserverObject) FinishedDeletingResource(resource objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("finishedDeletingResource:"), resource)
}
func (o TTSAXResourceManagerObserverObject) FinishedDownloadingResourceWasCancelled(resource objectivec.IObject, cancelled bool) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("finishedDownloadingResource:wasCancelled:"), resource, cancelled)
}
func (o TTSAXResourceManagerObserverObject) ResourceCacheDidReceiveUpdate() {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("resourceCacheDidReceiveUpdate"))
}
