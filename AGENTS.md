# RDMA hardware safety

Prefer `mac2` for RDMA or JACCL hardware experiments when practical. This host
is also suitable when a test requires both machines or `mac2` is unavailable.

- Never run two ranks against one local RDMA interface. In particular, do not
  use the loopback `devices-local.json` configuration: it can wedge
  AppleThunderboltRDMA and break login system-wide.
- Run every RDMA hardware test under a hard 60-second timeout or watchdog.
  Afterwards, inspect `ps` and confirm no related process is in state `U`.
  If one is, stop immediately, report it to the user, and do not retry or
  launch another rank.
