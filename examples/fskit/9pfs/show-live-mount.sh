#!/usr/bin/env bash
set -euo pipefail

mnt=${1:-}
if [[ -z "$mnt" ]]; then
	mnt=$(mount | awk '$0 ~ /\(9pfs[,)]/ {print $3; exit}')
fi

if [[ -z "$mnt" ]]; then
	echo "show-live-mount: no active 9pfs mount found" >&2
	exit 1
fi
if [[ ! -d "$mnt" ]]; then
	echo "show-live-mount: mount path does not exist: $mnt" >&2
	exit 1
fi

mount | awk -v mnt="$mnt" '$3 == mnt {print}'
echo
ls -la "$mnt"

read_file() {
	local name=$1
	local path=$mnt/$name
	if [[ -f "$path" ]]; then
		printf '%s: ' "$name"
		cat "$path"
		printf '\n'
	fi
}

echo
echo "first read:"
read_file README
read_file live-counter.txt
read_file live-time.txt
read_file dynamic-counter.txt
read_file dynamic-time.txt

sleep 2

echo
echo "second read:"
read_file live-counter.txt
read_file live-time.txt
read_file dynamic-counter.txt
read_file dynamic-time.txt

echo
echo "9pfs: live mount read-only demo ok"
