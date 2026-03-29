package vzkit

import snapshotx "github.com/tmc/apple/x/vzkit/snapshot"

// SnapshotMeta is the on-disk metadata format for VM state snapshots.
type SnapshotMeta = snapshotx.Meta

// SnapshotManager handles VM state snapshot operations.
type SnapshotManager = snapshotx.Manager

// NewSnapshotManager creates a snapshot manager for the given VM directory.
func NewSnapshotManager(vmDir string) *SnapshotManager { return snapshotx.NewManager(vmDir) }
