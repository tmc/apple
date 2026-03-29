package snapshot

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/vm"
)

// Meta is the on-disk metadata format for VM state snapshots.
type Meta struct {
	Name     string    `json:"name"`
	Created  time.Time `json:"created"`
	Size     int64     `json:"size"`
	VMState  string    `json:"vmState"`
	FilePath string    `json:"filePath"`
}

// Manager handles VM state snapshot operations.
type Manager struct {
	vmDir string
}

// NewManager creates a snapshot manager for vmDir.
func NewManager(vmDir string) *Manager {
	return &Manager{vmDir: vmDir}
}

func (m *Manager) snapshotsDir() string {
	return filepath.Join(m.vmDir, "snapshots")
}

func (m *Manager) snapshotPath(name string) string {
	return filepath.Join(m.snapshotsDir(), name+".vmstate")
}

func (m *Manager) metadataPath(name string) string {
	return filepath.Join(m.snapshotsDir(), name+".json")
}

func (m *Manager) ensureDir() error {
	return os.MkdirAll(m.snapshotsDir(), 0755)
}

// Save saves the current VM state to a named snapshot.
func (m *Manager) Save(machine vz.VZVirtualMachine, queue *vm.Queue, name string) error {
	if err := m.ensureDir(); err != nil {
		return fmt.Errorf("create snapshots directory: %w", err)
	}
	state := vm.State(queue, machine)
	wasPaused := state == vz.VZVirtualMachineStatePaused
	if state == vz.VZVirtualMachineStateRunning {
		errCh := make(chan error, 1)
		queue.Sync(func() {
			machine.PauseWithCompletionHandler(func(err error) {
				errCh <- err
			})
		})
		if err := <-errCh; err != nil {
			return fmt.Errorf("pause VM: %w", err)
		}
	} else if state != vz.VZVirtualMachineStatePaused {
		return fmt.Errorf("VM must be running or paused to save snapshot (current state: %s)", state.String())
	}

	snapshotFile := m.snapshotPath(name)
	saveURL := foundation.NewURLFileURLWithPath(snapshotFile)
	saveURL.Retain()

	errCh := make(chan error, 1)
	queue.Sync(func() {
		machine.SaveMachineStateToURLCompletionHandler(saveURL, func(err error) {
			errCh <- err
		})
	})
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("save snapshot: %w", err)
		}
	case <-time.After(60 * time.Second):
		return fmt.Errorf("save snapshot timed out")
	}

	var size int64
	if info, err := os.Stat(snapshotFile); err == nil {
		size = info.Size()
	}
	metadata := Meta{
		Name:     name,
		Created:  time.Now(),
		Size:     size,
		VMState:  "paused",
		FilePath: snapshotFile,
	}
	metaBytes, _ := json.MarshalIndent(metadata, "", "  ")
	os.WriteFile(m.metadataPath(name), metaBytes, 0644)

	if !wasPaused {
		resumeErrCh := make(chan error, 1)
		queue.Sync(func() {
			machine.ResumeWithCompletionHandler(func(err error) {
				resumeErrCh <- err
			})
		})
		if err := <-resumeErrCh; err != nil {
			return fmt.Errorf("resume VM after snapshot: %w", err)
		}
	}
	return nil
}

// Restore restores a VM from a previously saved snapshot.
func (m *Manager) Restore(machine vz.VZVirtualMachine, queue *vm.Queue, name string) error {
	snapshotFile := m.snapshotPath(name)
	if _, err := os.Stat(snapshotFile); os.IsNotExist(err) {
		return fmt.Errorf("snapshot '%s' not found", name)
	}
	state := vm.State(queue, machine)
	if state != vz.VZVirtualMachineStateStopped {
		return fmt.Errorf("VM must be stopped to restore snapshot (current state: %s)", state.String())
	}

	restoreURL := foundation.NewURLFileURLWithPath(snapshotFile)
	restoreURL.Retain()
	errCh := make(chan error, 1)
	queue.Sync(func() {
		machine.RestoreMachineStateFromURLCompletionHandler(restoreURL, func(err error) {
			errCh <- err
		})
	})
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("restore snapshot: %w", err)
		}
	case <-time.After(60 * time.Second):
		return fmt.Errorf("restore snapshot timed out")
	}
	return nil
}

// List returns all available snapshots sorted by creation time.
func (m *Manager) List() ([]Meta, error) {
	dir := m.snapshotsDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read snapshots directory: %w", err)
	}
	var snapshots []Meta
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".vmstate" {
			continue
		}
		name := entry.Name()[:len(entry.Name())-8]
		var info Meta
		metaPath := m.metadataPath(name)
		if metaBytes, err := os.ReadFile(metaPath); err == nil {
			json.Unmarshal(metaBytes, &info)
		}
		if info.Name == "" {
			info.Name = name
		}
		if fileInfo, err := entry.Info(); err == nil {
			info.Size = fileInfo.Size()
			if info.Created.IsZero() {
				info.Created = fileInfo.ModTime()
			}
		}
		info.FilePath = filepath.Join(dir, entry.Name())
		snapshots = append(snapshots, info)
	}
	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].Created.After(snapshots[j].Created)
	})
	return snapshots, nil
}

// Delete removes a named snapshot.
func (m *Manager) Delete(name string) error {
	snapshotFile := m.snapshotPath(name)
	if _, err := os.Stat(snapshotFile); os.IsNotExist(err) {
		return fmt.Errorf("snapshot '%s' not found", name)
	}
	if err := os.Remove(snapshotFile); err != nil {
		return fmt.Errorf("remove snapshot file: %w", err)
	}
	os.Remove(m.metadataPath(name))
	return nil
}
