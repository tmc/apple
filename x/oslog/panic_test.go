package oslog

import (
	"strings"
	"testing"
)

func TestTrimStack(t *testing.T) {
	// A synthetic trace: goroutine header, two oslog frames, the panic frame,
	// then the real frames. trimStack should leave only the header and the real
	// frames.
	in := strings.Join([]string{
		"goroutine 1 [running]:",
		"github.com/tmc/apple/x/oslog.Stack()",
		"\t/x/oslog/panic.go:15 +0x64",
		"github.com/tmc/apple/x/oslog.(*Logger).logPanic(...)",
		"\t/x/oslog/panic.go:65 +0x48",
		"panic({0x1, 0x2})",
		"\t/usr/local/go/src/runtime/panic.go:860 +0x12c",
		"main.boom(...)",
		"\t/app/main.go:9",
		"main.main()",
		"\t/app/main.go:18 +0x3c",
		"",
	}, "\n")

	got := trimStack(in)
	if !strings.HasPrefix(got, "goroutine 1 [running]:") {
		t.Errorf("trimStack dropped the goroutine header: %q", got)
	}
	if strings.Contains(got, "x/oslog.Stack") || strings.Contains(got, "logPanic") {
		t.Errorf("trimStack kept oslog frames:\n%s", got)
	}
	if strings.Contains(got, "panic(") {
		t.Errorf("trimStack kept the runtime panic frame:\n%s", got)
	}
	if !strings.Contains(got, "main.boom") {
		t.Errorf("trimStack dropped the real frames:\n%s", got)
	}
}

func TestStackNonEmpty(t *testing.T) {
	s := Stack()
	if !strings.Contains(s, "goroutine") {
		t.Errorf("Stack() = %q, want a goroutine trace", s)
	}
}
