# appnap

Compare an ordinary Go ticker with one protected by an
`NSProcessInfo.beginActivity(options:reason:)` activity:

```sh
go run ./examples/foundation/appnap
go run ./examples/foundation/appnap -protect
```

Leave the Mac idle and compare the reported intervals. The example is an
observation tool, not a timing guarantee; scheduler load and power policy
still affect both processes.
