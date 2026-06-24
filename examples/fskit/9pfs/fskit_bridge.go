//go:build darwin

package main

import (
	"fmt"
	"io/fs"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"9fans.net/go/plan9"
	p9 "github.com/hugelgupf/p9/p9"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/x/fskitbridge"
)

var ninepShims = fskitbridge.ReplyBlockShims{
	Error:    "NinePFSInvokeErrorBlock",
	Object:   "NinePFSInvokeObjectErrorBlock",
	ItemName: "NinePFSInvokeItemNameErrorBlock",
	Verifier: "NinePFSInvokeVerifierErrorBlock",
	Bool:     "NinePFSInvokeBoolErrorBlock",
	Size:     "NinePFSInvokeSizeErrorBlock",
}

var (
	serverMu sync.Mutex
	server   *fskitbridge.Server
)

// ensureServer registers the FSKit class set serving impl once and returns
// the process's bridge server. A failed attempt is not sticky: the next
// call retries, so a transient startup-ordering failure does not poison
// the process. Later calls return the existing server; their arguments are
// ignored.
func ensureServer(existingFSClass objc.Class, impl *ninepFileSystem) (*fskitbridge.Server, error) {
	serverMu.Lock()
	defer serverMu.Unlock()
	if server != nil {
		return server, nil
	}
	srv, err := fskitbridge.NewServer(fskitbridge.ServerConfig{
		FileSystemName:     "NinePFileSystem",
		VolumeName:         "NinePFSVolume",
		ItemName:           "NinePFSItem",
		ExistingFileSystem: existingFSClass,
		Shims:              ninepShims,
		FileSystem:         impl,
		Logf:               logBridgef,
	})
	if err != nil {
		return nil, err
	}
	server = srv
	return srv, nil
}

// currentServer returns the bridge server, or nil before ensureServer
// succeeds.
func currentServer() *fskitbridge.Server {
	serverMu.Lock()
	defer serverMu.Unlock()
	return server
}

// ninepFileSystem serves a 9p server as an FSKit unary file system. The
// backend is dialed on the first load unless one was provided directly.
type ninepFileSystem struct {
	mu      sync.Mutex
	config  fsConfig
	backend backend
}

var _ fskitbridge.UnaryFileSystem = (*ninepFileSystem)(nil)

type fsConfig struct {
	dialect string
	network string
	addr    string
	aname   string
}

func (f *ninepFileSystem) Probe(resource fskit.FSResource) (fskitbridge.ProbeResult, error) {
	if id := resource.GetID(); id != 0 {
		if _, err := fsConfigForResource(fsConfig{}, id); err != nil {
			return fskitbridge.ProbeResult{}, fmt.Errorf("probe 9p resource: %v: %w", err, fs.ErrInvalid)
		}
	}
	return fskitbridge.ProbeResult{Name: "9pfs"}, nil
}

func (f *ninepFileSystem) Load(resource fskit.FSResource) (fskitbridge.Volume, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.backend == nil {
		config, err := fsConfigForResource(f.config, resource.GetID())
		if err != nil {
			return nil, fmt.Errorf("parse 9p resource: %v: %w", err, fs.ErrInvalid)
		}
		logBridge(fmt.Sprintf("load 9p resource dialect=%s network=%s addr=%s aname=%q", config.dialect, config.network, config.addr, config.aname))
		backend, err := dialBackend(config.dialect, config.network, config.addr, config.aname)
		if err != nil {
			return nil, fmt.Errorf("dial 9p resource: %w", err)
		}
		f.config = config
		f.backend = backend
	}
	return newNinepVolume(f.backend), nil
}

func (f *ninepFileSystem) Unload() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.backend != nil {
		_ = f.backend.Close()
		f.backend = nil
	}
	return nil
}

// ninepVolume adapts a 9p backend to the fskitbridge volume interfaces.
// Items carry their cleaned 9p path; the volume keeps a path-to-ID table
// so an item keeps the same FSItemID across lookups and enumerations, and
// a set of live items whose paths are rewritten when a directory moves.
type ninepVolume struct {
	backend backend

	mu   sync.Mutex
	ids  map[string]fskit.FSItemID
	next fskit.FSItemID
	live map[*ninepItem]struct{}
}

type ninepItem struct {
	// path and info are guarded by the volume mutex.
	path     string
	id       fskit.FSItemID
	parentID fskit.FSItemID
	info     nodeInfo
}

var (
	_ fskitbridge.MutableVolume      = (*ninepVolume)(nil)
	_ fskitbridge.SymlinkVolume      = (*ninepVolume)(nil)
	_ fskitbridge.LinkVolume         = (*ninepVolume)(nil)
	_ fskitbridge.XattrVolume        = (*ninepVolume)(nil)
	_ fskitbridge.PreallocateVolume  = (*ninepVolume)(nil)
	_ fskitbridge.OpenCloseVolume    = (*ninepVolume)(nil)
	_ fskitbridge.AccessCheckVolume  = (*ninepVolume)(nil)
	_ fskitbridge.CapabilitiesVolume = (*ninepVolume)(nil)
	_ fskitbridge.StatisticsVolume   = (*ninepVolume)(nil)
	_ fskitbridge.PathConfVolume     = (*ninepVolume)(nil)
)

func newNinepVolume(b backend) *ninepVolume {
	return &ninepVolume{
		backend: b,
		ids:     make(map[string]fskit.FSItemID),
		live:    make(map[*ninepItem]struct{}),
	}
}

func (v *ninepVolume) VolumeName() string { return "9pfs" }

func (v *ninepVolume) Root() (fskitbridge.Item, error) {
	info, err := v.backend.Stat("/")
	if err != nil {
		return nil, err
	}
	return v.newItem("", fskit.FSItemIDParentOfRoot, info), nil
}

func (v *ninepVolume) newItem(path string, parentID fskit.FSItemID, info nodeInfo) *ninepItem {
	path = clean9PPath(path)
	v.mu.Lock()
	defer v.mu.Unlock()
	it := &ninepItem{path: path, id: v.idForPathLocked(path), parentID: parentID, info: info}
	v.live[it] = struct{}{}
	return it
}

// idForPathLocked returns the stable FSItemID for a cleaned 9p path,
// minting a new ID on first use.
func (v *ninepVolume) idForPathLocked(path string) fskit.FSItemID {
	if path == "" {
		return fskit.FSItemIDRootDirectory
	}
	if id, ok := v.ids[path]; ok {
		return id
	}
	v.next++
	id := fskit.FSItemIDRootDirectory + v.next
	v.ids[path] = id
	return id
}

func (v *ninepVolume) idForPath(path string) fskit.FSItemID {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.idForPathLocked(clean9PPath(path))
}

// item returns x as a *ninepItem.
func (v *ninepVolume) item(x fskitbridge.Item) (*ninepItem, error) {
	it, ok := x.(*ninepItem)
	if !ok {
		return nil, syscall.EINVAL
	}
	return it, nil
}

func (v *ninepVolume) pathOf(it *ninepItem) string {
	v.mu.Lock()
	defer v.mu.Unlock()
	return it.path
}

func (v *ninepVolume) setInfo(it *ninepItem, info nodeInfo) {
	v.mu.Lock()
	defer v.mu.Unlock()
	it.info = info
}

func (v *ninepVolume) Lookup(dir fskitbridge.Item, name string) (fskitbridge.Item, error) {
	d, err := v.item(dir)
	if err != nil {
		return nil, err
	}
	childPath := child9PPath(v.pathOf(d), name)
	info, err := v.backend.Stat(childPath)
	if err != nil {
		// A lookup that cannot stat the child is reported as ENOENT
		// regardless of the backend errno: to the VFS a name it cannot
		// resolve is not-found. The underlying error survives in the text.
		return nil, fmt.Errorf("lookup %s: %v: %w", childPath, err, syscall.ENOENT)
	}
	return v.newItem(childPath, d.id, info), nil
}

func (v *ninepVolume) Reclaim(item fskitbridge.Item) {
	if it, ok := item.(*ninepItem); ok {
		v.mu.Lock()
		delete(v.live, it)
		v.mu.Unlock()
	}
}

func (v *ninepVolume) Attributes(item fskitbridge.Item) (fskit.FSItemAttributes, error) {
	it, err := v.item(item)
	if err != nil {
		return fskit.FSItemAttributes{}, err
	}
	v.mu.Lock()
	info, id, parentID := it.info, it.id, it.parentID
	v.mu.Unlock()
	return attributesForNode(info, id, parentID), nil
}

func (v *ninepVolume) ReadDir(dir fskitbridge.Item) ([]fskitbridge.DirEntry, error) {
	d, err := v.item(dir)
	if err != nil {
		return nil, err
	}
	dirPath := v.pathOf(d)
	entries, err := v.backend.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	out := make([]fskitbridge.DirEntry, 0, len(entries))
	for _, entry := range entries {
		out = append(out, fskitbridge.DirEntry{
			Name:       entry.Name,
			Type:       itemTypeFor(entry),
			Attributes: attributesForNode(entry, v.idForPath(child9PPath(dirPath, entry.Name)), d.id),
		})
	}
	return out, nil
}

func (v *ninepVolume) Read(file fskitbridge.Item, offset int64, buf []byte) (int, error) {
	it, err := v.item(file)
	if err != nil {
		return 0, err
	}
	return v.backend.ReadFileAt(v.pathOf(it), offset, buf)
}

func (v *ninepVolume) Write(file fskitbridge.Item, offset int64, data []byte) (int, error) {
	it, err := v.item(file)
	if err != nil {
		return 0, err
	}
	path := v.pathOf(it)
	n, err := v.backend.WriteFile(path, offset, data)
	if err != nil {
		return n, err
	}
	if info, err := v.backend.Stat(path); err == nil {
		v.setInfo(it, info)
	}
	return n, nil
}

func (v *ninepVolume) Create(dir fskitbridge.Item, name string, typ fskit.FSItemType, mode uint32) (fskitbridge.Item, error) {
	d, err := v.item(dir)
	if err != nil {
		return nil, err
	}
	var isDir bool
	switch typ {
	case fskit.FSItemTypeFile:
	case fskit.FSItemTypeDirectory:
		isDir = true
	default:
		return nil, fmt.Errorf("create item type %d: %w", typ, errUnsupported)
	}
	childPath := child9PPath(v.pathOf(d), name)
	info, err := v.backend.Create(childPath, mode, isDir)
	if err != nil {
		return nil, err
	}
	return v.newItem(childPath, d.id, info), nil
}

func (v *ninepVolume) Remove(dir fskitbridge.Item, name string, item fskitbridge.Item) error {
	it, err := v.item(item)
	if err != nil {
		return err
	}
	path := v.pathOf(it)
	if err := v.backend.Remove(path); err != nil {
		return err
	}
	v.mu.Lock()
	delete(v.ids, path)
	v.mu.Unlock()
	return nil
}

func (v *ninepVolume) Rename(item, srcDir fskitbridge.Item, srcName string, dstDir fskitbridge.Item, dstName string, over fskitbridge.Item) error {
	it, err := v.item(item)
	if err != nil {
		return err
	}
	dd, err := v.item(dstDir)
	if err != nil {
		return err
	}
	oldPath := v.pathOf(it)
	dstPath := child9PPath(v.pathOf(dd), dstName)
	if err := v.backend.Rename(oldPath, dstPath); err != nil {
		return err
	}
	info, err := v.backend.Stat(dstPath)
	if err != nil {
		return err
	}
	v.mu.Lock()
	v.movePathsLocked(oldPath, dstPath)
	it.path = dstPath
	it.parentID = dd.id
	it.info = info
	v.mu.Unlock()
	return nil
}

// movePathsLocked rewrites the cached paths of oldPath and everything
// beneath it after a rename, in both the ID table and the live item set.
func (v *ninepVolume) movePathsLocked(oldPath, newPath string) {
	moved := func(p string) (string, bool) {
		if p == oldPath {
			return newPath, true
		}
		if strings.HasPrefix(p, oldPath+"/") {
			return newPath + p[len(oldPath):], true
		}
		return "", false
	}
	type idMove struct {
		from, to string
	}
	var idMoves []idMove
	for p := range v.ids {
		if np, ok := moved(p); ok {
			idMoves = append(idMoves, idMove{p, np})
		}
	}
	for _, m := range idMoves {
		id := v.ids[m.from]
		delete(v.ids, m.from)
		v.ids[m.to] = id
	}
	for it := range v.live {
		if np, ok := moved(it.path); ok {
			it.path = np
		}
	}
}

func (v *ninepVolume) movePaths(oldPath, newPath string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.movePathsLocked(oldPath, newPath)
}

func (v *ninepVolume) SetAttributes(item fskitbridge.Item, set fskitbridge.SetAttributes) error {
	it, err := v.item(item)
	if err != nil {
		return err
	}
	attr := setAttr{Mode: set.Mode, Size: set.Size}
	if set.AccessTime != nil {
		if set.AccessTime.Sec < 0 {
			return syscall.EINVAL
		}
		seconds := uint64(set.AccessTime.Sec)
		attr.Accessed = &seconds
	}
	if set.ModifyTime != nil {
		if set.ModifyTime.Sec < 0 {
			return syscall.EINVAL
		}
		seconds := uint64(set.ModifyTime.Sec)
		attr.Modified = &seconds
	}
	info, err := v.backend.SetAttr(v.pathOf(it), attr)
	if err != nil {
		return err
	}
	v.setInfo(it, info)
	return nil
}

func (v *ninepVolume) Readlink(item fskitbridge.Item) (string, error) {
	it, err := v.item(item)
	if err != nil {
		return "", err
	}
	return v.backend.Readlink(v.pathOf(it))
}

func (v *ninepVolume) Symlink(dir fskitbridge.Item, name, target string) (fskitbridge.Item, error) {
	d, err := v.item(dir)
	if err != nil {
		return nil, err
	}
	childPath := child9PPath(v.pathOf(d), name)
	info, err := v.backend.CreateSymlink(childPath, target)
	if err != nil {
		return nil, err
	}
	return v.newItem(childPath, d.id, info), nil
}

func (v *ninepVolume) Link(item, dir fskitbridge.Item, name string) error {
	it, err := v.item(item)
	if err != nil {
		return err
	}
	d, err := v.item(dir)
	if err != nil {
		return err
	}
	_, err = v.backend.CreateLink(v.pathOf(it), child9PPath(v.pathOf(d), name))
	return err
}

func (v *ninepVolume) GetXattr(item fskitbridge.Item, name string) ([]byte, error) {
	it, err := v.item(item)
	if err != nil {
		return nil, err
	}
	data, err := v.backend.GetXattr(v.pathOf(it), name)
	if err != nil {
		// Report a missing attribute (ENOATTR) for any backend failure;
		// backend errnos use the server's values, not this host's.
		return nil, fmt.Errorf("get xattr %s: %v: %w", name, err, fs.ErrNotExist)
	}
	return data, nil
}

func (v *ninepVolume) SetXattr(item fskitbridge.Item, name string, data []byte) error {
	it, err := v.item(item)
	if err != nil {
		return err
	}
	return v.backend.SetXattr(v.pathOf(it), name, data)
}

func (v *ninepVolume) RemoveXattr(item fskitbridge.Item, name string) error {
	it, err := v.item(item)
	if err != nil {
		return err
	}
	return v.backend.RemoveXattr(v.pathOf(it), name)
}

func (v *ninepVolume) ListXattr(item fskitbridge.Item) ([]string, error) {
	it, err := v.item(item)
	if err != nil {
		return nil, err
	}
	return v.backend.ListXattr(v.pathOf(it))
}

func (v *ninepVolume) Preallocate(file fskitbridge.Item, offset int64, length uint64) (uint64, error) {
	it, err := v.item(file)
	if err != nil {
		return 0, err
	}
	return v.backend.Preallocate(v.pathOf(it), offset, length)
}

func (v *ninepVolume) Open(item fskitbridge.Item, modes fskit.FSVolumeOpenModes) error  { return nil }
func (v *ninepVolume) Close(item fskitbridge.Item, modes fskit.FSVolumeOpenModes) error { return nil }

func (v *ninepVolume) CheckAccess(item fskitbridge.Item, access fskit.FSAccessMask) (bool, error) {
	return true, nil
}

func (v *ninepVolume) SupportedCapabilities() fskit.FSVolumeSupportedCapabilities {
	capabilities := fskit.NewFSVolumeSupportedCapabilities()
	capabilities.SetSupports64BitObjectIDs(true)
	capabilities.SetSupportsHiddenFiles(true)
	capabilities.SetSupportsSparseFiles(true)
	if supportsSymlinks(v.backend) {
		capabilities.SetSupportsSymbolicLinks(true)
	}
	if supportsHardLinks(v.backend) {
		capabilities.SetSupportsHardLinks(true)
	}
	capabilities.SetCaseFormat(fskit.FSVolumeCaseFormatSensitive)
	return capabilities
}

func (v *ninepVolume) Statistics() fskit.FSStatFSResult {
	result := fskit.NewStatFSResultWithFileSystemTypeName("9pfs")
	result.SetBlockSize(4096)
	result.SetIoSize(4096)
	result.SetTotalBlocks(1)
	result.SetAvailableBlocks(0)
	result.SetFreeBlocks(0)
	result.SetUsedBlocks(1)
	result.SetTotalFiles(1)
	result.SetFreeFiles(0)
	return result
}

func (v *ninepVolume) PathConf() fskitbridge.PathConf {
	conf := fskitbridge.DefaultPathConf()
	if supportsHardLinks(v.backend) {
		conf.MaximumLinkCount = 32767
	}
	return conf
}

func attributesForNode(info nodeInfo, id, parentID fskit.FSItemID) fskit.FSItemAttributes {
	typ := itemTypeFor(info)
	mode := info.Mode & 0777
	if typ == fskit.FSItemTypeDirectory {
		mode |= 0111
	}
	ts := syscall.Timespec{Sec: int64(info.Modified)}
	return fskitbridge.NewItemAttributes().
		Type(typ).
		Mode(mode).
		LinkCount(1).
		UID(uint32(os.Getuid())).
		GID(uint32(os.Getgid())).
		Flags(0).
		Size(info.Length).
		AllocSize(info.Length).
		AccessTime(ts).
		ModifyTime(ts).
		ChangeTime(ts).
		BirthTime(ts).
		ID(id).
		ParentID(parentID).
		Build()
}

func itemTypeFor(info nodeInfo) fskit.FSItemType {
	if info.Mode&uint32(plan9.DMDIR) != 0 || p9FileModeIsDir(info.Mode) {
		return fskit.FSItemTypeDirectory
	}
	if info.Mode&uint32(plan9.DMSYMLINK) != 0 || p9FileModeIsSymlink(info.Mode) {
		return fskit.FSItemTypeSymlink
	}
	mode := p9.FileMode(info.Mode)
	switch {
	case mode.IsNamedPipe():
		return fskit.FSItemTypeFIFO
	case mode.IsCharacterDevice():
		return fskit.FSItemTypeCharDevice
	case mode.IsBlockDevice():
		return fskit.FSItemTypeBlockDevice
	case mode.IsSocket():
		return fskit.FSItemTypeSocket
	}
	return fskit.FSItemTypeFile
}

func p9FileModeIsDir(mode uint32) bool {
	return mode&0170000 == 0040000
}

func p9FileModeIsSymlink(mode uint32) bool {
	return mode&0170000 == 0120000
}

// Capabilities a backend opts into by implementing these interfaces. A
// backend reports a capability by returning true; the type switch this
// replaced broke when the backend was wrapped (errnoBackend), so capability
// is a behavior the backend declares, forwarded transparently through the
// embedding wrapper.
type symlinkCapable interface{ supportsSymlinks() bool }
type hardLinkCapable interface{ supportsHardLinks() bool }

func supportsSymlinks(b backend) bool {
	c, ok := b.(symlinkCapable)
	return ok && c.supportsSymlinks()
}

func supportsHardLinks(b backend) bool {
	c, ok := b.(hardLinkCapable)
	return ok && c.supportsHardLinks()
}

func posixError(errno syscall.Errno) objc.ID {
	return fskitbridge.POSIXError(errno)
}

// logBridge logs a bridge diagnostic. os_log (nativeExtensionLog) is always
// emitted: it is the running extension's proper channel and goes to the
// unified log, not a file. The stderr and on-disk sinks are gated behind
// NINEPFS_DEBUG, so a normal mount does not spew to stderr or append
// uncapped to log files; set NINEPFS_DEBUG=1 to enable them while debugging.
func logBridge(msg string) {
	nativeExtensionLog(msg)
	if !bridgeDebug {
		return
	}
	line := "9pfs: " + msg
	fmt.Fprintln(os.Stderr, line)
	bridgeLogOnce.Do(openBridgeLogs)
	bridgeLogMu.Lock()
	defer bridgeLogMu.Unlock()
	for _, f := range bridgeLogFiles {
		fmt.Fprintln(f, line)
	}
}

func logBridgef(format string, args ...any) {
	logBridge(fmt.Sprintf(format, args...))
}

var (
	bridgeDebug    = os.Getenv("NINEPFS_DEBUG") != ""
	bridgeLogOnce  sync.Once
	bridgeLogMu    sync.Mutex
	bridgeLogFiles []*os.File
)

func openBridgeLogs() {
	addBridgeLog := func(name string) {
		f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err == nil {
			bridgeLogFiles = append(bridgeLogFiles, f)
		}
	}
	addBridgeLog("/tmp/9pfs-extension.log")
	if dir, err := os.UserCacheDir(); err == nil {
		_ = os.MkdirAll(dir, 0700)
		addBridgeLog(filepath.Join(dir, "9pfs-extension.log"))
	}
}

// defaultFSConfigFromEnv reads the mount defaults from NINEPFS_* environment
// variables, falling back to a classic 9P2000 server on the default port.
func defaultFSConfigFromEnv() fsConfig {
	return fsConfig{
		dialect: envOr("NINEPFS_DIALECT", "9p2000"),
		network: envOr("NINEPFS_NET", "tcp"),
		addr:    envOr("NINEPFS_ADDR", "127.0.0.1:5640"),
		aname:   os.Getenv("NINEPFS_ANAME"),
	}
}

func envOr(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}

func fsConfigForResource(base fsConfig, resource objc.ID) (fsConfig, error) {
	if base.dialect == "" {
		base.dialect = "9p2000"
	}
	if base.network == "" {
		base.network = "tcp"
	}
	if base.addr == "" {
		base.addr = "127.0.0.1:5640"
	}
	if resource == 0 {
		return base, nil
	}
	u, err := urlForResource(resource)
	if err != nil {
		return fsConfig{}, err
	}
	if u == "" {
		return base, nil
	}
	return fsConfigForURL(base, u)
}

func urlForResource(resource objc.ID) (string, error) {
	urlID := objc.Send[objc.ID](resource, objc.Sel("url"))
	if urlID == 0 {
		return "", fmt.Errorf("resource has no url")
	}
	u := foundation.NSURLFromID(urlID)
	s := u.AbsoluteString()
	if s == "" {
		return "", fmt.Errorf("resource url is empty")
	}
	return s, nil
}

func fsConfigForURL(base fsConfig, raw string) (fsConfig, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return fsConfig{}, fmt.Errorf("parse %q: %w", raw, err)
	}
	config := base
	if v := u.Query().Get("dialect"); v != "" {
		config.dialect = v
	}
	if v := u.Query().Get("aname"); v != "" {
		config.aname = v
	}
	switch strings.ToLower(u.Scheme) {
	case "9p", "ninep", "tcp":
		config.network = "tcp"
		host := u.Hostname()
		if host == "" {
			return fsConfig{}, fmt.Errorf("missing host in %q", raw)
		}
		port := u.Port()
		if port == "" {
			port = "5640"
		}
		if _, err := strconv.Atoi(port); err != nil {
			return fsConfig{}, fmt.Errorf("bad port %q in %q", port, raw)
		}
		config.addr = net.JoinHostPort(host, port)
		if config.aname == "" {
			config.aname = strings.TrimPrefix(u.EscapedPath(), "/")
			if v, err := url.PathUnescape(config.aname); err == nil {
				config.aname = v
			}
		}
	case "unix":
		config.network = "unix"
		config.addr = u.Path
		if config.addr == "" {
			config.addr = u.Host
		}
		if config.addr == "" {
			return fsConfig{}, fmt.Errorf("missing unix socket path in %q", raw)
		}
	default:
		return fsConfig{}, fmt.Errorf("unsupported 9p resource scheme %q", u.Scheme)
	}
	return config, nil
}
