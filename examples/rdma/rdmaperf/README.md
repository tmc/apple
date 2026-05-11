# rdmaperf

`rdmaperf` compares TCP over ordinary network links, TCP over Thunderbolt
Bridge, and RDMA userspace readiness.

Run a server on the receiving Mac:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf serve -listen 169.254.x.y:9000
```

Run clients from the other Mac:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf tcp -addr 169.254.x.y:9000 -pattern stream -size 1M -duration 30s
GOWORK=off go run ./examples/rdma/rdmaperf tcp -addr 169.254.x.y:9000 -pattern pingpong -size 64 -duration 30s
GOWORK=off go run ./examples/rdma/rdmaperf sweep -addr 169.254.x.y:9000 -pattern stream -duration 10s
```

Use `interfaces` to find local addresses:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf interfaces
```

Use the same commands against Wi-Fi, Ethernet, and Thunderbolt Bridge addresses.
The TCP tests measure the chosen IP path; they do not prove RDMA datapath use.

Check RDMA userspace readiness separately:

```sh
GOWORK=off go run ./examples/rdma/rdmaperf rdma-probe
```

True RDMA datapath benchmarking requires successful protection-domain,
completion-queue, memory-region, and queue-pair lifecycle. When those verbs are
unavailable, `rdma-probe` reports the failing step instead of pretending to
measure RDMA throughput.
