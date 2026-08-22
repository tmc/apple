# vmremap

This example maps the current task's pages through `mach_vm_remap`, mutates
the returned mapping, and observes the shared-page mutation. The mapping is
outside the Go heap and is explicitly released with `mach.Unmap`.

```sh
go run ./examples/kernel/vmremap
```
