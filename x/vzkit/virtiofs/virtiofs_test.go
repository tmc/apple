package virtiofs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseMount(t *testing.T) {
	dir := t.TempDir()
	mount, err := ParseMount(dir + ":MyTag:ro:cache=none,noac")
	if err != nil {
		t.Fatalf("ParseMount: %v", err)
	}
	if mount.Tag != "MyTag" || !mount.ReadOnly || len(mount.MountOpts) != 2 {
		t.Fatalf("ParseMount = %#v", mount)
	}
}

func TestCreateDevicesEmpty(t *testing.T) {
	got, err := CreateDevices(nil)
	if err != nil {
		t.Fatalf("CreateDevices: %v", err)
	}
	if got != nil {
		t.Fatalf("CreateDevices(nil) = %#v, want nil", got)
	}
}

func TestParseMountRejectsMissingDir(t *testing.T) {
	_, err := ParseMount(filepath.Join(os.TempDir(), "definitely-not-here"))
	if err == nil {
		t.Fatal("ParseMount missing path = nil error")
	}
}
