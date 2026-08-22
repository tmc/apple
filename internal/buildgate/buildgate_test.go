// Copyright 2026 The tmc/apple Authors. All rights reserved.

package buildgate

import (
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

// frameworks are the packages the runtime smoke gate depends on. A build
// failure in any of them blanks that gate, so each is reported individually.
var frameworks = []string{
	"appkit",
	"coreimage",
	"foundation",
	"iosurface",
	"naturallanguage",
	"objectivec",
	"vision",
}

// modulePath is the import path prefix. Building by import path rather than by
// relative directory matters: the test's working directory is this package, so
// "./vision/" resolves to nothing and every framework would report UNBUILDABLE
// for a reason that has nothing to do with the framework.
const modulePath = "github.com/tmc/apple/"

func build(pkg string) (string, error) {
	out, err := exec.Command("go", "build", modulePath+pkg).CombinedOutput()
	return string(out), err
}

func TestFrameworksBuild(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Apple frameworks are only available on darwin")
	}
	// Canary: this package has no dependencies and must always build. If it
	// does not, the harness is broken rather than the tree, and every
	// UNBUILDABLE below would be meaningless.
	if out, err := build("internal/buildgate"); err != nil {
		t.Fatalf("harness is broken, not the tree: building the gate's own "+
			"package failed, so no verdict below can be trusted: %v\n%s",
			err, indent(out))
	}
	var ok, broken int
	for _, fw := range frameworks {
		t.Run(fw, func(t *testing.T) {
			out, err := build(fw)
			if err != nil {
				broken++
				t.Errorf("UNBUILDABLE (not measured, not clean): %v\n%s",
					err, indent(out))
				return
			}
			ok++
		})
	}
	t.Logf("buildgate summary: %d build, %d UNBUILDABLE of %d frameworks",
		ok, broken, len(frameworks))
	if broken > 0 {
		t.Logf("a framework listed UNBUILDABLE above is why smoketest reports " +
			"[build failed] rather than a per-framework board")
	}
}

func indent(s string) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	return "\t" + strings.Join(lines, "\n\t")
}
