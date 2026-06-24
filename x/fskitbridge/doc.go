// Package fskitbridge helps implement FSKit file systems in Go.
//
// An FSKit file system extension exposes three Objective-C classes: a file
// system that subclasses FSUnaryFileSystem, a volume that subclasses
// FSVolume, and an item that subclasses FSItem. A [Server] registers this
// class set and implements its operation selectors on top of a Go
// [UnaryFileSystem]: implementations work with Go values and errors, and
// the Server handles FSKit object identity, reply blocks, set-attribute
// requests, directory packing, and errno reporting. A minimal volume
// implements [Volume]; optional interfaces such as [MutableVolume] and
// [XattrVolume] add the remaining FSKit operations.
//
// The lower layer remains available for file systems that need direct
// control of the selector set. RegisterClasses registers a class set with
// method implementations written in Go. FSKit passes each operation a
// reply block that the implementation must invoke to deliver its result;
// ReplyBlocks invokes the common reply block shapes and reports a reply
// that could not be delivered. When typed C shim functions are linked into
// the process, NewReplyBlocksWithShims routes invocation through them;
// otherwise the block's invoke pointer is called directly.
//
// ItemAttributesBuilder assembles the FSItemAttributes values returned by
// lookup, enumeration, and attribute operations, and POSIXError constructs
// the POSIX-domain NSError values that reply blocks accept.
//
// The file systems under examples/fskit use this package; see
// examples/fskit/9pfs for a complete example.
package fskitbridge
