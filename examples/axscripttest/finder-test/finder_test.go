// Package finder_test exercises axscripttest against Finder.
//
// Run without macgo (relies on terminal's existing accessibility trust):
//
//	go test -v -run TestFinder
//
// Run with macgo to create a .app bundle for TCC registration:
//
//	go test -v -run TestFinder -axprompt
//
// When the terminal (e.g. iTerm2) already has accessibility permission
// granted via System Settings, the test passes without macgo. The macgo
// path is needed when running from a context that lacks a TCC identity
// (e.g. a bare go test binary with no bundle ID).
package finder_test

import (
	"flag"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/tmc/apple/x/axscripttest"
	"github.com/tmc/apple/x/axuiautomation"
)

var axprompt = flag.Bool("axprompt", false, "use macgo to create .app bundle for TCC identity")

func TestMain(m *testing.M) {
	runtime.LockOSThread()
	flag.Parse()
	if *axprompt {
		if err := axscripttest.EnsureTrusted(); err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(m.Run())
}

func TestFinder(t *testing.T) {
	t.Logf("accessibility trusted: %v, macgo: %v", axuiautomation.IsProcessTrusted(), *axprompt)
	axscripttest.RunFile(t, axscripttest.Config{
		App: "com.apple.finder",
	}, "testdata/finder.txt")
}
