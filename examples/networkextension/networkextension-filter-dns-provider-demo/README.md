# NetworkExtension filter and DNS provider demo

This example checks generated DNS proxy and content-filter provider callback
surfaces against a local Objective-C object. It invokes provider lifecycle
completion blocks, DNS proxy flow decisions, filter lifecycle callbacks,
filter reports/configuration, filter data verdict methods, settings
application, flow-resume/update methods, packet handler methods, and default
content-filter value objects. It does not install provider extensions or
process live network traffic.

Run:

```sh
GOWORK=off go test ./examples/networkextension/networkextension-filter-dns-provider-demo -count=1
GOWORK=off go run ./examples/networkextension/networkextension-filter-dns-provider-demo
```
