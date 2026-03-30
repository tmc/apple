// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"testing"

	"github.com/tmc/apple/objectivec"
)

func TestSetCompletionHandlerRetainedNilRequest(t *testing.T) {
	var req ANERequest
	b := req.SetCompletionHandlerRetained(func(bool, error) {})
	if b != nil {
		t.Fatalf("SetCompletionHandlerRetained() on nil request = %v, want nil", b)
	}
}

func TestCompletionHandlerBindingReleaseNil(t *testing.T) {
	var b *CompletionHandlerBinding
	b.Release()
}

func TestCompletionHandlerBindingReleaseIdempotent(t *testing.T) {
	b := &CompletionHandlerBinding{}
	b.Release()
	b.Release()
}

func TestCompletionEventWrapperCompile(t *testing.T) {
	var _ func(ANEVirtualClient, objectivec.IObject, objectivec.IObject, objectivec.IObject, uint32, BoolErrorHandler) (bool, func(), error) = ANEVirtualClient.DoEvaluateWithModelOptionsRequestQosCompletionEventHandlerError
	var _ func(ANEVirtualClient, objectivec.IObject, objectivec.IObject, objectivec.IObject, uint32, BoolErrorHandler) (bool, func(), error) = ANEVirtualClient.DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventHandlerError
}

func TestCompletionEventWrapperZeroValueNoPanic(t *testing.T) {
	var c ANEVirtualClient
	ok, release, err := c.DoEvaluateWithModelOptionsRequestQosCompletionEventHandlerError(objectivec.Object{}, objectivec.Object{}, objectivec.Object{}, 0, func(bool, error) {})
	if ok {
		t.Fatalf("zero-value wrapper returned ok=true, want false")
	}
	if release != nil {
		t.Fatalf("zero-value wrapper returned non-nil release func")
	}
	if err != nil {
		t.Fatalf("zero-value wrapper returned err=%v, want nil", err)
	}
}
