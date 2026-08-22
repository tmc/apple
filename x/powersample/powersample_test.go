package powersample

import (
	"bufio"
	"math"
	"os"
	"testing"
	"time"
)

func parseAll(t *testing.T, path string) []Sample {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var p parser
	var samples []Sample
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if s, ok := p.line(sc.Text()); ok {
			samples = append(samples, s)
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	if s, ok := p.flush(); ok {
		samples = append(samples, s)
	}
	return samples
}

func TestParseMacOS26(t *testing.T) {
	samples := parseAll(t, "testdata/macos26.txt")
	want := []Sample{
		{
			Elapsed: time.Duration(502.45 * float64(time.Millisecond)),
			Power:   Power{CPU: 0.458, GPU: 0.061, ANE: 0, Combined: 0.519},
		},
		{
			Elapsed: time.Duration(500.11 * float64(time.Millisecond)),
			Power:   Power{CPU: 1.0, GPU: 4.5, ANE: 2.0, Combined: 7.5},
		},
	}
	if len(samples) != len(want) {
		t.Fatalf("got %d samples, want %d", len(samples), len(want))
	}
	for i, w := range want {
		g := samples[i]
		if g.Elapsed != w.Elapsed {
			t.Errorf("sample %d: elapsed = %v, want %v", i, g.Elapsed, w.Elapsed)
		}
		for _, c := range []struct {
			name      string
			got, want float64
		}{
			{"CPU", g.Power.CPU, w.Power.CPU},
			{"GPU", g.Power.GPU, w.Power.GPU},
			{"ANE", g.Power.ANE, w.Power.ANE},
			{"Combined", g.Power.Combined, w.Power.Combined},
		} {
			if math.Abs(c.got-c.want) > 1e-9 {
				t.Errorf("sample %d: %s = %v W, want %v W", i, c.name, c.got, c.want)
			}
		}
	}
}

// TestParseDuplicateRailsNotDoubleCounted pins the first-occurrence rule:
// the gpu_power and ane_power samplers restate rails that cpu_power's
// block already printed (with slightly different values in testdata so a
// regression is visible, not coincidentally equal).
func TestParseDuplicateRailsNotDoubleCounted(t *testing.T) {
	samples := parseAll(t, "testdata/macos26.txt")
	if len(samples) != 2 {
		t.Fatalf("got %d samples, want 2", len(samples))
	}
	if got := samples[1].Power.GPU; got != 4.5 {
		t.Errorf("GPU = %v W, want 4.5 (first occurrence, not the later 4.501)", got)
	}
	if got := samples[1].Power.ANE; got != 2.0 {
		t.Errorf("ANE = %v W, want 2.0 (first occurrence, not the later 2.001)", got)
	}
}

func TestParseWatts(t *testing.T) {
	tests := []struct {
		in   string
		want float64
		ok   bool
	}{
		{" 458 mW", 0.458, true},
		{" 4.5 W", 4.5, true},
		{" 0 mW", 0, true},
		{" garbage", 0, false},
		{"", 0, false},
		{" 12 kW", 0, false},
	}
	for _, tt := range tests {
		got, ok := parseWatts(tt.in)
		if got != tt.want || ok != tt.ok {
			t.Errorf("parseWatts(%q) = %v, %v, want %v, %v", tt.in, got, ok, tt.want, tt.ok)
		}
	}
}

func TestParseElapsed(t *testing.T) {
	got := parseElapsed("*** Sampled system activity (Tue Aug 12 10:00:00 2026 -0700) (502.45ms elapsed) ***")
	if want := time.Duration(502.45 * float64(time.Millisecond)); got != want {
		t.Errorf("parseElapsed = %v, want %v", got, want)
	}
	if got := parseElapsed("*** Sampled system activity ***"); got != 0 {
		t.Errorf("parseElapsed without elapsed = %v, want 0", got)
	}
}
