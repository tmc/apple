package platform

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOrCreateGenericMachineIdentifier(t *testing.T) {
	path := filepath.Join(t.TempDir(), "machine.id")

	first, created, err := LoadOrCreateGenericMachineIdentifier(path)
	if err != nil {
		t.Fatalf("LoadOrCreateGenericMachineIdentifier create: %v", err)
	}
	if first.ID == 0 {
		t.Fatal("created machine identifier is nil")
	}
	if !created {
		t.Fatal("created = false, want true")
	}
	if info, err := os.Stat(path); err != nil || info.Size() == 0 {
		t.Fatalf("machine identifier file = %v, %v; want non-empty", info, err)
	}

	second, created, err := LoadOrCreateGenericMachineIdentifier(path)
	if err != nil {
		t.Fatalf("LoadOrCreateGenericMachineIdentifier load: %v", err)
	}
	if second.ID == 0 {
		t.Fatal("loaded machine identifier is nil")
	}
	if created {
		t.Fatal("created = true, want false")
	}
}

func TestCreateEFIBootLoader(t *testing.T) {
	path := filepath.Join(t.TempDir(), "efi.nvram")

	first, created, err := CreateEFIBootLoader(path)
	if err != nil {
		t.Fatalf("CreateEFIBootLoader create: %v", err)
	}
	if first.ID == 0 {
		t.Fatal("created boot loader is nil")
	}
	if !created {
		t.Fatal("created = false, want true")
	}
	if info, err := os.Stat(path); err != nil || info.Size() == 0 {
		t.Fatalf("EFI variable store file = %v, %v; want non-empty", info, err)
	}

	second, created, err := CreateEFIBootLoader(path)
	if err != nil {
		t.Fatalf("CreateEFIBootLoader load: %v", err)
	}
	if second.ID == 0 {
		t.Fatal("loaded boot loader is nil")
	}
	if created {
		t.Fatal("created = true, want false")
	}
}
