# fsevents

Watch a directory with the journal-backed FSEvents service:

```sh
go run ./examples/coreservices/fsevents -path . -duration 30s
touch a-file
```

The callback prints event IDs, raw event flags, and paths. It requests
file-level events and uses the dispatch-queue API. The example resolves the
FSEvents subframework directly because current macOS releases do not expose
these symbols through the CoreServices umbrella dylib.
