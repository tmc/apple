// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var completionHandlerAssocKey struct{}

// CompletionHandlerBinding keeps an Objective-C completion block alive while
// a request is in flight. Call Release once completion has fired (or on error).
type CompletionHandlerBinding struct {
	requestID objc.ID
	cleanup   func()
	once      sync.Once
}

// Release detaches and releases the retained completion block.
func (b *CompletionHandlerBinding) Release() {
	if b == nil {
		return
	}
	b.once.Do(func() {
		if b.requestID != 0 {
			objc.Send[struct{}](b.requestID, objc.Sel("setCompletionHandler:"), objc.ID(0))
			objectivec.Objc_setAssociatedObject(
				objectivec.Object{ID: b.requestID},
				unsafe.Pointer(&completionHandlerAssocKey),
				objectivec.Object{},
				objectivec.OBJC_ASSOCIATION_RETAIN_NONATOMIC,
			)
		}
		if b.cleanup != nil {
			b.cleanup()
		}
	})
}

// SetCompletionHandlerRetained sets a completion handler and retains the block
// via associated object to avoid premature release on async paths.
func (a ANERequest) SetCompletionHandlerRetained(handler BoolErrorHandler) *CompletionHandlerBinding {
	if a.ID == 0 {
		return nil
	}
	blockID, blockCleanup := NewBoolErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("setCompletionHandler:"), blockID)
	objectivec.Objc_setAssociatedObject(
		objectivec.Object{ID: a.ID},
		unsafe.Pointer(&completionHandlerAssocKey),
		objectivec.Object{ID: blockID},
		objectivec.OBJC_ASSOCIATION_RETAIN_NONATOMIC,
	)
	return &CompletionHandlerBinding{requestID: a.ID, cleanup: blockCleanup}
}

// ClearCompletionHandler clears and releases any retained completion handler.
func (a ANERequest) ClearCompletionHandler() {
	(&CompletionHandlerBinding{requestID: a.ID}).Release()
}

// DoEvaluateWithModelOptionsRequestQosCompletionEventHandlerError wraps the
// completionEvent argument with a typed BoolError handler block.
//
// The caller must keep the returned release func alive until completion has
// fired, then call it exactly once.
func (a ANEVirtualClient) DoEvaluateWithModelOptionsRequestQosCompletionEventHandlerError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, handler BoolErrorHandler) (ok bool, release func(), err error) {
	if a.ID == 0 {
		return false, nil, nil
	}
	blockID, blockCleanup := NewBoolErrorBlock(handler)
	ok, err = a.DoEvaluateWithModelOptionsRequestQosCompletionEventError(model, options, request, qos, objectivec.Object{ID: blockID})
	if !ok || err != nil {
		blockCleanup()
		return ok, nil, err
	}
	return ok, blockCleanup, nil
}

// DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventHandlerError wraps
// legacy completionEvent with a typed BoolError handler block.
//
// The caller must keep the returned release func alive until completion has
// fired, then call it exactly once.
func (a ANEVirtualClient) DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventHandlerError(legacy objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, handler BoolErrorHandler) (ok bool, release func(), err error) {
	if a.ID == 0 {
		return false, nil, nil
	}
	blockID, blockCleanup := NewBoolErrorBlock(handler)
	ok, err = a.DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventError(legacy, options, request, qos, objectivec.Object{ID: blockID})
	if !ok || err != nil {
		blockCleanup()
		return ok, nil, err
	}
	return ok, blockCleanup, nil
}

// ModelWithMILTextWeightsOptionsPlistIsMILModel calls
// modelWithMILText:weights:optionsPlist:isMILModel: when available.
func (_ANEInMemoryModelDescriptorClass ANEInMemoryModelDescriptorClass) ModelWithMILTextWeightsOptionsPlistIsMILModel(mILText objectivec.IObject, weights objectivec.IObject, plist objectivec.IObject, isMILModel bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEInMemoryModelDescriptorClass.class), objc.Sel("modelWithMILText:weights:optionsPlist:isMILModel:"), mILText, weights, plist, isMILModel)
	return objectivec.Object{ID: rv}
}
