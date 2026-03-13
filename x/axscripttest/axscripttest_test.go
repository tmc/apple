package axscripttest_test

import (
	"flag"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/tmc/apple/x/axscripttest"
)

var axprompt = flag.Bool("axprompt", false, "set up TCC identity and trigger system accessibility permission dialog")

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
	axscripttest.RunFile(t, axscripttest.Config{
		App: "com.apple.finder",
	}, "testdata/finder.txt")
}
