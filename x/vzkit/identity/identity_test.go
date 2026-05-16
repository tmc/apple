package identity

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadMacMachineIdentifierErrors(t *testing.T) {
	_, err := LoadMacMachineIdentifier(filepath.Join(t.TempDir(), "missing"))
	if err == nil || !strings.Contains(err.Error(), "read machine identifier") {
		t.Fatalf("missing error = %v", err)
	}
	path := filepath.Join(t.TempDir(), "machine.id")
	if err := os.WriteFile(path, nil, 0644); err != nil {
		t.Fatal(err)
	}
	_, err = LoadMacMachineIdentifier(path)
	if err == nil || !strings.Contains(err.Error(), "empty") {
		t.Fatalf("empty error = %v", err)
	}
}

func TestCreateMacMachineIdentifier(t *testing.T) {
	path := filepath.Join(t.TempDir(), "machine.id")
	if err := CreateMacMachineIdentifier(path); err != nil {
		t.Fatalf("CreateMacMachineIdentifier: %v", err)
	}
	if _, err := LoadMacMachineIdentifier(path); err != nil {
		t.Fatalf("LoadMacMachineIdentifier: %v", err)
	}
}
