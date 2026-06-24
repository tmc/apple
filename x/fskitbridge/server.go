package fskitbridge

import (
	"errors"
	"log"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
)

// ServerConfig configures a Server.
type ServerConfig struct {
	// FileSystemName, VolumeName, and ItemName name the Objective-C
	// classes the Server registers.
	FileSystemName string
	VolumeName     string
	ItemName       string

	// ExistingFileSystem, if nonzero, is used as the file system class
	// instead of registering one; see [ClassConfig]. The host shim then
	// routes operations to the Server through [Server.ProbeResource],
	// [Server.LoadResource], and [Server.UnloadResource].
	ExistingFileSystem objc.Class

	// Shims names linked reply block shims; see [ReplyBlockShims].
	Shims ReplyBlockShims

	// FileSystem is the file system implementation.
	FileSystem UnaryFileSystem

	// Logf, if non-nil, receives diagnostic messages. If nil, failures
	// to deliver a reply are reported through the log package and other
	// diagnostics are dropped.
	Logf func(format string, args ...any)
}

// A Server adapts a [UnaryFileSystem] to FSKit. It registers the
// Objective-C class set, implements the FSKit operation selectors, tracks
// the FSItem and FSVolume objects it hands to FSKit, and reports operation
// errors as POSIX-domain NSErrors.
//
// The Server declares the FSKit volume operation protocols and answers
// each protocol's inhibited query according to the optional interfaces the
// loaded [Volume] implements. Volume renaming is always inhibited.
//
// Register a class set once per process: a second Server with the same
// class names takes over that class set's operations.
type Server struct {
	cfg     ServerConfig
	classes ClassSet
	reply   *ReplyBlocks

	volumes sync.Map // objc.ID (FSVolume) -> *serverVolume
	items   sync.Map // objc.ID (FSItem) -> *serverItem
	loaded  sync.Map // objc.ID (file system) -> objc.ID (FSVolume)
}

type serverVolume struct {
	impl Volume
	root objc.ID
}

type serverItem struct {
	vol  Volume
	item Item
}

// NewServer registers the class set for cfg and returns its Server.
func NewServer(cfg ServerConfig) (*Server, error) {
	if cfg.FileSystem == nil {
		return nil, errors.New("missing file system implementation")
	}
	s := &Server{cfg: cfg, reply: NewReplyBlocksWithShims(cfg.Shims)}
	classes, err := RegisterClasses(ClassConfig{
		FileSystemName:      cfg.FileSystemName,
		VolumeName:          cfg.VolumeName,
		ItemName:            cfg.ItemName,
		ExistingFileSystem:  cfg.ExistingFileSystem,
		FileSystemProtocols: []string{"FSUnaryFileSystemOperations"},
		VolumeProtocols: []string{
			"FSVolumeOperations",
			"FSVolumeReadWriteOperations",
			"FSVolumeOpenCloseOperations",
			"FSVolumeAccessCheckOperations",
			"FSVolumePathConfOperations",
			"FSVolumeXattrOperations",
			"FSVolumePreallocateOperations",
			"FSVolumeRenameOperations",
		},
		FileSystemMethods: s.fileSystemMethods(),
		VolumeMethods:     s.volumeMethods(),
		ItemMethods: []objc.MethodDef{
			{Cmd: objc.Sel("dealloc"), Fn: s.impItemDealloc},
		},
	})
	if err != nil {
		return nil, err
	}
	s.classes = classes
	return s, nil
}

// Classes returns the registered class set.
func (s *Server) Classes() ClassSet {
	return s.classes
}

// NewFileSystem returns a new instance of the file system class.
func (s *Server) NewFileSystem() objc.ID {
	fs := objc.Send[objc.ID](objc.ID(s.classes.FileSystem), objc.Sel("alloc"))
	return objc.Send[objc.ID](fs, objc.Sel("init"))
}

// NewVolume returns a new instance of the volume class serving v,
// with its root item already resolved.
func (s *Server) NewVolume(v Volume) (objc.ID, error) {
	root, err := v.Root()
	if err != nil {
		return 0, err
	}
	volumeID := fskit.NewVolumeIdentifierWithUUID(foundation.NewNSUUID())
	name := fskit.NewFileNameWithString(v.VolumeName())
	volume := objc.Send[objc.ID](objc.ID(s.classes.Volume), objc.Sel("alloc"))
	volume = objc.Send[objc.ID](volume, objc.Sel("initWithVolumeID:volumeName:"), volumeID, name)
	s.volumes.Store(volume, &serverVolume{impl: v, root: s.newItem(v, root)})
	return volume, nil
}

func (s *Server) newItem(v Volume, item Item) objc.ID {
	obj := objc.Send[objc.ID](objc.ID(s.classes.Item), objc.Sel("alloc"))
	obj = objc.Send[objc.ID](obj, objc.Sel("init"))
	s.items.Store(obj, &serverItem{vol: v, item: item})
	return obj
}

func (s *Server) fileSystemMethods() []objc.MethodDef {
	return []objc.MethodDef{
		{Cmd: objc.Sel("probeResource:replyHandler:"), Fn: s.impProbeResource},
		{Cmd: objc.Sel("loadResource:options:replyHandler:"), Fn: s.impLoadResource},
		{Cmd: objc.Sel("unloadResource:options:replyHandler:"), Fn: s.impUnloadResource},
		{Cmd: objc.Sel("dealloc"), Fn: s.impFileSystemDealloc},
	}
}

func (s *Server) impProbeResource(self objc.ID, _ objc.SEL, resource, reply objc.ID) {
	s.ProbeResource(self, resource, reply)
}

func (s *Server) impLoadResource(self objc.ID, _ objc.SEL, resource, options, reply objc.ID) {
	s.LoadResource(self, resource, options, reply)
}

func (s *Server) impUnloadResource(self objc.ID, _ objc.SEL, resource, options, reply objc.ID) {
	s.UnloadResource(self, resource, options, reply)
}

func (s *Server) impFileSystemDealloc(self objc.ID, _ objc.SEL) {
	s.loaded.Delete(self)
	if err := s.cfg.FileSystem.Unload(); err != nil {
		s.logf("unload on dealloc: %v", err)
	}
}

// ProbeResource probes a resource and replies with an FSProbeResult.
// It is the implementation of probeResource:replyHandler:, exported for
// host shims that route the selector to Go themselves.
func (s *Server) ProbeResource(self, resource, reply objc.ID) {
	result, err := s.cfg.FileSystem.Probe(fskit.FSResourceFromID(resource))
	if err != nil {
		s.replyErr("probe", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	containerID := fskit.NewContainerIdentifierWithUUID(foundation.NewNSUUID())
	probe := fskit.GetFSProbeResultClass().UsableProbeResultWithNameContainerID(result.Name, containerID)
	s.replyErr("probe", s.reply.ObjectError(reply, probe.GetID(), 0))
}

// LoadResource loads a resource, replies with its FSVolume, and marks the
// container ready. It is the implementation of
// loadResource:options:replyHandler:, exported for host shims that route
// the selector to Go themselves.
func (s *Server) LoadResource(self, resource, options, reply objc.ID) {
	v, err := s.cfg.FileSystem.Load(fskit.FSResourceFromID(resource))
	if err != nil {
		s.logf("load resource: %v", err)
		s.replyErr("load", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	volume, err := s.NewVolume(v)
	if err != nil {
		s.logf("load resource root: %v", err)
		s.replyErr("load", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	s.loaded.Store(self, volume)
	fskit.FSUnaryFileSystemFromID(self).SetContainerStatus(fskit.GetFSContainerStatusClass().Ready())
	s.replyErr("load", s.reply.ObjectError(reply, volume, 0))
}

// UnloadResource unloads the resource and replies. It is the
// implementation of unloadResource:options:replyHandler:, exported for
// host shims that route the selector to Go themselves.
func (s *Server) UnloadResource(self, resource, options, reply objc.ID) {
	if volume, ok := s.loaded.LoadAndDelete(self); ok {
		objc.Send[struct{}](volume.(objc.ID), objc.Sel("release"))
	}
	err := s.cfg.FileSystem.Unload()
	s.replyErr("unload", s.reply.Error(reply, s.errorFor(err)))
}

func (s *Server) volumeFor(self objc.ID) (*serverVolume, bool) {
	v, ok := s.volumes.Load(self)
	if !ok {
		return nil, false
	}
	return v.(*serverVolume), true
}

func (s *Server) itemFor(obj objc.ID) (Item, bool) {
	v, ok := s.items.Load(obj)
	if !ok {
		return nil, false
	}
	return v.(*serverItem).item, true
}

// errorFor returns the POSIX-domain NSError for err, or 0 for nil.
func (s *Server) errorFor(err error) objc.ID {
	if err == nil {
		return 0
	}
	return POSIXError(errnoFor(err))
}

func (s *Server) xattrErrorFor(err error) objc.ID {
	if err == nil {
		return 0
	}
	return POSIXError(xattrErrnoFor(err))
}

func (s *Server) logf(format string, args ...any) {
	if s.cfg.Logf != nil {
		s.cfg.Logf(format, args...)
	}
}

// replyErr reports a reply that could not be delivered: the FSKit caller
// is left without a response, so it is always reported.
func (s *Server) replyErr(op string, err error) {
	if err == nil {
		return
	}
	if s.cfg.Logf != nil {
		s.cfg.Logf("%s reply failed: %v", op, err)
		return
	}
	log.Printf("fskitbridge: %s reply failed: %v", op, err)
}

func fileNameString(name objc.ID) string {
	return fskit.FSFileNameFromID(name).String()
}

func bytesFromNSData(id objc.ID) []byte {
	if id == 0 {
		return nil
	}
	data := foundation.NSDataFromID(id)
	n := data.Length()
	if n == 0 {
		return nil
	}
	src := data.Bytes()
	out := make([]byte, n)
	copy(out, unsafe.Slice((*byte)(src), n))
	return out
}
