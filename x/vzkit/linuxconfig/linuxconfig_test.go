package linuxconfig

import (
	"strings"
	"testing"

	displayx "github.com/tmc/apple/x/vzkit/display"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr string
	}{
		{
			name: "valid",
			config: Config{
				CPUCount: 2,
				MemoryGB: 4,
				Display:  []displayx.Config{{Width: 1024, Height: 768, PPI: 144}},
			},
		},
		{
			name:    "zero cpu",
			config:  Config{MemoryGB: 4},
			wantErr: "cpu count",
		},
		{
			name:    "zero memory",
			config:  Config{CPUCount: 2},
			wantErr: "memory",
		},
		{
			name: "bad display",
			config: Config{
				CPUCount: 2,
				MemoryGB: 4,
				Display:  []displayx.Config{{Height: 768}},
			},
			wantErr: "display[0]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatal("expected error")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %q, want %q", err, tt.wantErr)
			}
		})
	}
}
