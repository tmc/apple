#!/usr/bin/env bash
set -euo pipefail

url=${1:-ninep://127.0.0.1:5640?dialect=9p2000l}
mnt=${2:-/Volumes/9pfs-test}
script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
source "$script_dir/scriptlib.sh"

die() {
	echo "test-installed-mount: $*" >&2
	exit 1
}

[[ -x /Applications/NinePFSHost.app/Contents/MacOS/NinePFSHost ]] ||
	die "install signed NinePFSHost.app in /Applications first"

[[ -d "$mnt" ]] || die "create mount point $mnt first"

is_mounted() {
	mount | awk '{print $3}' | grep -qx "$mnt"
}

if is_mounted; then
	die "$mnt is already mounted"
fi

require_no_active_9pfs_mounts
"$script_dir/preflight-installed.sh"

/Applications/NinePFSHost.app/Contents/MacOS/NinePFSHost --fskit-probe

if ! /sbin/mount -F -t 9pfs "$url" "$mnt"; then
	is_mounted || die "mount failed and $mnt is not mounted"
	echo "test-installed-mount: mount returned nonzero after mounting $mnt" >&2
fi
trap 'umount "$mnt" 2>/dev/null || true' EXIT

ls -la "$mnt"
mkdir "$mnt/codex-dir"
printf 'written through mounted 9pfs\n' > "$mnt/codex-dir/a.txt"
cat "$mnt/codex-dir/a.txt"
mv "$mnt/codex-dir/a.txt" "$mnt/codex-dir/b.txt"
truncate -s 8 "$mnt/codex-dir/b.txt"
test "$(cat "$mnt/codex-dir/b.txt")" = "written "
chmod 600 "$mnt/codex-dir/b.txt"
touch -t 202001020304 "$mnt/codex-dir/b.txt"
mtime=$(stat -f '%m' "$mnt/codex-dir/b.txt")
test "$mtime" = "1577963040" || die "mtime got $mtime, want 1577963040"

if [[ "$url" == *9p2000l* ]]; then
	ln -s b.txt "$mnt/codex-dir/sym"
	test "$(readlink "$mnt/codex-dir/sym")" = "b.txt"
	ln "$mnt/codex-dir/b.txt" "$mnt/codex-dir/hard"
	xattr -w user.codex value "$mnt/codex-dir/b.txt"
	test "$(xattr -p user.codex "$mnt/codex-dir/b.txt")" = "value"
	xattr -d user.codex "$mnt/codex-dir/b.txt"
	rm "$mnt/codex-dir/sym" "$mnt/codex-dir/hard"
fi

rm "$mnt/codex-dir/b.txt"
rmdir "$mnt/codex-dir"

echo "9pfs: installed FSKit mount read/write/rename/chmod/mtime/truncate/link/xattr/remove ok"
