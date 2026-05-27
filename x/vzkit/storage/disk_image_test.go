//go:build darwin

package storage

import (
	"strings"
	"testing"

	vz "github.com/tmc/apple/virtualization"
)

func TestDiskImageModes(t *testing.T) {
	tests := []struct {
		name        string
		policy      CachePolicy
		wantCaching vz.VZDiskImageCachingMode
		wantSync    vz.VZDiskImageSynchronizationMode
	}{
		{"durable", CacheDurable, vz.VZDiskImageCachingModeCached, vz.VZDiskImageSynchronizationModeFsync},
		{"ephemeral", CacheEphemeral, vz.VZDiskImageCachingModeCached, vz.VZDiskImageSynchronizationModeNone},
		{"read-only", CacheReadOnly, vz.VZDiskImageCachingModeAutomatic, vz.VZDiskImageSynchronizationModeFull},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCaching, gotSync, err := DiskImageModes(tt.policy)
			if err != nil {
				t.Fatal(err)
			}
			if gotCaching != tt.wantCaching || gotSync != tt.wantSync {
				t.Fatalf("DiskImageModes(%v) = %v, %v, want %v, %v", tt.policy, gotCaching, gotSync, tt.wantCaching, tt.wantSync)
			}
		})
	}
}

func TestDiskImageModesUnknownPolicy(t *testing.T) {
	_, _, err := DiskImageModes(CachePolicy(99))
	if err == nil || !strings.Contains(err.Error(), "unknown disk cache policy") {
		t.Fatalf("err = %v, want unknown disk cache policy", err)
	}
}
