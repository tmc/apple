package appleneuralengine

import (
	"testing"

	"github.com/tmc/apple/objc"
)

func TestANEFSharedEventOptionKeysLoaded(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("AppleNeuralEngine framework is not loaded on this host")
	}

	tests := []struct {
		name string
		obj  interface {
			GetID() objc.ID
			Description() string
		}
	}{
		{
			name: "KANEFDisableIOFencesUseSharedEventsKey",
			obj:  KANEFDisableIOFencesUseSharedEventsKey,
		},
		{
			name: "KANEFEnableFWToFWSignal",
			obj:  KANEFEnableFWToFWSignal,
		},
	}

	for _, tc := range tests {
		if tc.obj.GetID() == 0 {
			t.Fatalf("%s was not resolved via dlsym", tc.name)
		}
		if desc := tc.obj.Description(); desc == "" {
			t.Fatalf("%s has empty Objective-C description", tc.name)
		}
	}
}
