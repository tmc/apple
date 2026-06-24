#!/usr/bin/env bash

ninepfs_active_mounts() {
	mount | awk '$0 ~ /\(9pfs[,)]/ {print $3}'
}

require_no_active_9pfs_mounts() {
	local mounts

	if [[ "${NINEPFS_ALLOW_ACTIVE_MOUNTS:-}" == "yes" ]]; then
		return 0
	fi
	mounts=$(ninepfs_active_mounts)
	if [[ -n "$mounts" ]]; then
		echo "9pfs: active 9pfs mount(s) present; refusing to continue" >&2
		printf '%s\n' "$mounts" >&2
		echo "9pfs: unmount them first, or set NINEPFS_ALLOW_ACTIVE_MOUNTS=yes for a deliberate probe window" >&2
		return 1
	fi
}
