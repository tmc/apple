# ANE Firmware-to-Firmware (FWToFW) Hardware Signaling Demo

This example demonstrates enabling **Firmware-to-Firmware (FWToFW)** signaling (`kANEFEnableFWToFWSignal`) and setting up hardware synchronization options for direct ANE-to-GPU event-driven execution.

## Background

When dispatching execution requests to the Apple Neural Engine (ANE), standard operating system IO fences can introduce thread context switching overhead. By enabling `kANEFEnableFWToFWSignal` and `kANEFDisableIOFencesUseSharedEventsKey`, hardware signaling events occur directly at the coprocessor firmware layer (`_ANESharedSignalEvent` and `_ANESharedWaitEvent`).

## Running the Example

```bash
go run ./examples/compute/ane_fwtofw
```
