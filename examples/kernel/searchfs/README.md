# searchfs

Search an APFS volume catalog without walking every directory:

```sh
go run ./examples/kernel/searchfs -v . -name go -f
go run ./examples/kernel/searchfs -v / -name Applications -e -d
```

The example demonstrates the bounded `searchfs(2)` session: partial results
are emitted before handling `EAGAIN`, `EBUSY` restarts the catalog search, and
`ENOBUFS` grows the return buffer. It resolves returned file IDs with
`fsgetpath(2)`. Both calls are SDK-declared but not generated in this tree, so
the small private binding remains local to the example.
