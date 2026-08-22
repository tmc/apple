# rdmainfo

`rdmainfo` reports Apple RDMA provider state and performs bounded local
diagnostics.

Use discovery and preflight commands first:

```sh
GOWORK=off go run ./examples/rdma/rdmainfo status -json
GOWORK=off go run ./examples/rdma/rdmainfo preflight -json
GOWORK=off go run ./examples/rdma/rdmainfo query-device -name rdma_en3 -json
```

`status`, `list`, `features`, `preflight`, `scan`, `matrix`, `open`, `query`,
and `query-device` do not allocate provider PD, CQ, QP, or MR resources.

`lifecycle` allocates one PD and one CQ and may register one MR. It is a local
capability observation, not a datapath test. It requires both acknowledgements:

```sh
CONFIRM_RDMA_RESOURCE_LIFECYCLE=one-shot-resource-lifecycle \
  GOWORK=off go run ./examples/rdma/rdmainfo lifecycle \
  -name rdma_en3 -allow-resource-lifecycle -json
```

Do not use `rdmainfo` to drive RTR/RTS or post RDMA work. Use the separately
gated `rdmaperf` commands for a reviewed two-host experiment.
