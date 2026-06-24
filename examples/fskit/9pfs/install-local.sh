#!/usr/bin/env bash
set -euo pipefail

build_dir=${1:-/tmp/9pfs-build}
app=$build_dir/NinePFSHost.app
fsbundle=$build_dir/9pfs.fs
script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
source "$script_dir/scriptlib.sh"

cat <<EOF
9pfs local install plan

Required artifacts:
  $app
  $fsbundle

Signing prerequisites:
  Build with CODESIGN_IDENTITY and matching provisioning profiles. The build
  script auto-discovers profiles from ~/Library/MobileDevice/Provisioning Profiles;
  use NINEPFS_REQUIRE_PROFILES=yes to fail early when they are missing.
  The extension profile must match dev.tmc.apple.examples.fskit.9pfs.extension
  and include com.apple.developer.fskit.fsmodule.

Commands:
  sudo rm -rf /Applications/NinePFSHost.app /Library/Filesystems/9pfs.fs
  sudo cp -R "$app" /Applications/NinePFSHost.app
  sudo cp -R "$fsbundle" /Library/Filesystems/9pfs.fs
  pluginkit -a /Applications/NinePFSHost.app/Contents/Extensions/NinePFSExtension.appex
  /Applications/NinePFSHost.app/Contents/MacOS/NinePFSHost --open-fskit-settings
  /Applications/NinePFSHost.app/Contents/MacOS/NinePFSHost --fskit-probe

After System Settings shows dev.tmc.apple.examples.fskit.9pfs.extension enabled:
  ./test-installed-live-9p2000l.sh
EOF

if [[ "${CONFIRM_9PFS_INSTALL:-}" != "yes" ]]; then
	echo
	echo "Set CONFIRM_9PFS_INSTALL=yes to run the copy/register commands." >&2
	exit 0
fi

[[ -d "$app" ]] || { echo "missing $app" >&2; exit 1; }
[[ -d "$fsbundle" ]] || { echo "missing $fsbundle" >&2; exit 1; }

require_no_active_9pfs_mounts

sudo rm -rf /Applications/NinePFSHost.app /Library/Filesystems/9pfs.fs
sudo cp -R "$app" /Applications/NinePFSHost.app
sudo cp -R "$fsbundle" /Library/Filesystems/9pfs.fs
pluginkit -a /Applications/NinePFSHost.app/Contents/Extensions/NinePFSExtension.appex
/Applications/NinePFSHost.app/Contents/MacOS/NinePFSHost --open-fskit-settings
/Applications/NinePFSHost.app/Contents/MacOS/NinePFSHost --fskit-probe
