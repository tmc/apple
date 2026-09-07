# Recovery USB transport

`Conn.Upload` sends image bytes through a connection selected by ECID. It holds
exclusive connection access throughout the transfer. Other context-bearing
operations can cancel while waiting for that access.

- DFU: 2048-byte logical blocks, Apple footer, and download-idle polling after
  each block. If the footer does not fit, it is sent separately using the final
  block's ID, matching the pinned libirecovery implementation.
- Recovery: 32768-byte bulk chunks and a terminating zero-length packet for
  payload lengths divisible by 512.
- Short transfers fail without automatic retry. DFU `bStatus` errors are exposed
  through `StatusError`. Polling honors the full 24-bit millisecond timeout.
- Empty images, unsupported modes and DFU images exceeding the 16-bit block
  counter fail. Recovery bulk uploads have no DFU block-count limit.
- A non-idle initial DFU state triggers CLRSTATUS (state 10) or ABORT, followed
  by an error requiring an explicit retry. Recovery-command failures propagate.
  The transfer has a two-minute ceiling or the caller's deadline.

Upload completion is not firmware execution or a completed restore. DFU
manifestation notification, reset and ECID-matched re-enumeration remain separate,
unimplemented steps. Recovery uploads need their subsequent execution command.
No physical or research-VM upload has been qualified.

## Protocol evidence

Packet layout and CRC follow [libirecovery 1.3.0, commit 29592eb6](https://github.com/libimobiledevice/libirecovery/blob/29592eb6cae8b25b214aa1e3cfb6ef4a6d555d43/src/libirecovery.c),
function `irecv_send_buffer`. The footer's first 12 bytes are
`ffffffffac05000155464410`. Its last four bytes are the little-endian CRC obtained
by complementing the standard IEEE CRC32 of the image followed by those 12 bytes.

The cloned pymobiledevice3 revision `a16ffc51` uses a different CRC initialization
in `IRecv.send_buffer`: its `binascii.crc32(image, -1)` path differs from the
pinned libirecovery algorithm. For image `123456789`, the footer CRC values are
`14c2daa4` (libirecovery) and `55c61f2f` (that Python path). Tests preserve the
libirecovery value; this disagreement still needs live-device qualification.

DFU status uses byte 0 for `bStatus`, bytes 1–3 for the little-endian poll delay,
and byte 4 for state. State 5 is download-idle; manifestation-sync is state 6.
See the [USB DFU protocol state diagrams](https://www.st.com/resource/en/application_note/cd00264379-usb-dfu-protocol-used-in-the-stm32-bootloader-stmicroelectronics.pdf).
