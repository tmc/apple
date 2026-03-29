package vzkit

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseNetworkMode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    NetworkConfig
		wantErr bool
	}{
		{
			name:  "empty defaults to nat",
			input: "",
			want:  NetworkConfig{Mode: NetworkModeNAT},
		},
		{
			name:  "nat",
			input: "nat",
			want:  NetworkConfig{Mode: NetworkModeNAT},
		},
		{
			name:  "none",
			input: "none",
			want:  NetworkConfig{Mode: NetworkModeNone},
		},
		{
			name:  "vmnet",
			input: "vmnet",
			want:  NetworkConfig{Mode: NetworkModeVMNet},
		},
		{
			name:  "host-only",
			input: "host-only",
			want:  NetworkConfig{Mode: NetworkModeHostOnly},
		},
		{
			name:  "bridged",
			input: "bridged:en0",
			want:  NetworkConfig{Mode: NetworkModeBridged, Interface: "en0"},
		},
		{
			name:    "bridged without interface",
			input:   "bridged:",
			wantErr: true,
		},
		{
			name:    "unknown mode",
			input:   "foo",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNetworkMode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseNetworkMode(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Fatalf("ParseNetworkMode(%q) = %+v, want %+v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseVolumeSpec(t *testing.T) {
	tmp := t.TempDir()
	hostDir := filepath.Join(tmp, "host")
	if err := os.Mkdir(hostDir, 0o755); err != nil {
		t.Fatalf("Mkdir host dir: %v", err)
	}

	hostFile := filepath.Join(tmp, "file.txt")
	if err := os.WriteFile(hostFile, []byte("x"), 0o644); err != nil {
		t.Fatalf("WriteFile host file: %v", err)
	}

	wantHostDir, err := filepath.Abs(hostDir)
	if err != nil {
		t.Fatalf("Abs host dir: %v", err)
	}
	if resolved, err := filepath.EvalSymlinks(wantHostDir); err == nil {
		wantHostDir = resolved
	}

	tests := []struct {
		name    string
		spec    string
		want    VolumeMount
		wantErr bool
	}{
		{
			name: "default rw mount",
			spec: hostDir,
			want: VolumeMount{HostPath: wantHostDir},
		},
		{
			name: "read-only mount",
			spec: hostDir + ":ro",
			want: VolumeMount{HostPath: wantHostDir, ReadOnly: true},
		},
		{
			name: "tagged mount",
			spec: hostDir + ":workspace:rw",
			want: VolumeMount{HostPath: wantHostDir, Tag: "workspace"},
		},
		{
			name:    "empty spec",
			spec:    "",
			wantErr: true,
		},
		{
			name:    "missing path",
			spec:    filepath.Join(tmp, "does-not-exist"),
			wantErr: true,
		},
		{
			name:    "not a directory",
			spec:    hostFile,
			wantErr: true,
		},
		{
			name:    "multiple tags",
			spec:    hostDir + ":one:two",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseVolumeSpec(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseVolumeSpec(%q) error = %v, wantErr %v", tt.spec, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ParseVolumeSpec(%q) = %+v, want %+v", tt.spec, got, tt.want)
			}
		})
	}
}
