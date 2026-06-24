package fskitbridge

import "testing"

func TestRegisterClassesMissingNames(t *testing.T) {
	tests := []struct {
		name string
		cfg  ClassConfig
	}{
		{"all missing", ClassConfig{}},
		{"missing file system", ClassConfig{VolumeName: "V", ItemName: "I"}},
		{"missing volume", ClassConfig{FileSystemName: "F", ItemName: "I"}},
		{"missing item", ClassConfig{FileSystemName: "F", VolumeName: "V"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := RegisterClasses(tt.cfg); err == nil {
				t.Fatal("RegisterClasses() error = nil, want error")
			}
		})
	}
}
