package display

import "testing"

func TestParseSpec(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Config
		wantErr bool
	}{
		{name: "preset", input: "4k", want: Config{Width: 3840, Height: 2160, PPI: 144}},
		{name: "explicit", input: "1920x1080@160", want: Config{Width: 1920, Height: 1080, PPI: 160}},
		{name: "invalid", input: "bad", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSpec(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseSpec(%q) error = nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseSpec(%q) error = %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("ParseSpec(%q) = %#v, want %#v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSliceString(t *testing.T) {
	var s Slice
	if got := s.String(); got != "" {
		t.Fatalf("String() = %q, want empty", got)
	}
	if err := s.Set("1920x1080"); err != nil {
		t.Fatalf("Set: %v", err)
	}
	if got := s.String(); got != "1920x1080@144" {
		t.Fatalf("String() = %q, want %q", got, "1920x1080@144")
	}
}
