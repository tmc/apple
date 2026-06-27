package vzkit

import (
	"strings"
	"testing"

	"github.com/tmc/apple/x/vzkit/virtiofs"
)

func TestRequireRootVolume(t *testing.T) {
	tests := []struct {
		name    string
		cfg     LinuxVMConfig
		wantErr bool
	}{
		{
			name: "default tag present",
			cfg: LinuxVMConfig{
				Volumes: []VolumeMount{{HostPath: "/r", Tag: DefaultRootVolumeTag}},
			},
		},
		{
			name: "custom tag present",
			cfg: LinuxVMConfig{
				RootVolumeTag: "myroot",
				Volumes:       []VolumeMount{{HostPath: "/r", Tag: "myroot"}},
			},
		},
		{
			name: "root among several volumes",
			cfg: LinuxVMConfig{
				Volumes: []VolumeMount{
					{HostPath: "/a", Tag: "hostmount0"},
					{HostPath: "/r", Tag: DefaultRootVolumeTag},
				},
			},
		},
		{
			name:    "no volumes at all",
			cfg:     LinuxVMConfig{},
			wantErr: true,
		},
		{
			name: "volumes present but none tagged root",
			cfg: LinuxVMConfig{
				Volumes: []VolumeMount{{HostPath: "/a", Tag: "hostmount0"}},
			},
			wantErr: true,
		},
		{
			name: "custom tag set but volume uses default",
			cfg: LinuxVMConfig{
				RootVolumeTag: "myroot",
				Volumes:       []VolumeMount{{HostPath: "/r", Tag: DefaultRootVolumeTag}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := requireRootVolume(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("requireRootVolume() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), "root filesystem") {
				t.Errorf("error %q should mention the root filesystem requirement", err)
			}
		})
	}
}

// Ensure the virtiofs.Mount alias is the type used for Volumes, so a root
// volume can be expressed with the same struct callers already build.
var _ = VolumeMount(virtiofs.Mount{})
