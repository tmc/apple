//go:build darwin

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/x/fskitbridge"
)

var (
	serverMu sync.Mutex
	server   *fskitbridge.Server

	shimHandle uintptr
	replyProbe func(objc.ID, objc.ID, objc.ID)
	shimVolume func() objc.ID
)

// ensureServer registers the TinyFS class set once and returns the bridge
// server. A failed attempt is not sticky: the next call retries.
func ensureServer(existingFSClass objc.Class) (*fskitbridge.Server, error) {
	serverMu.Lock()
	defer serverMu.Unlock()
	if server != nil {
		return server, nil
	}
	srv, err := fskitbridge.NewServer(fskitbridge.ServerConfig{
		FileSystemName:     "GoTinyFSFileSystem",
		VolumeName:         "GoTinyFSVolume",
		ItemName:           "GoTinyFSItem",
		ExistingFileSystem: existingFSClass,
		Shims: fskitbridge.ReplyBlockShims{
			Error:  "TinyFSShimReplyError",
			Object: "TinyFSShimReplyVolume",
		},
		FileSystem: tinyFS{},
		Logf: func(format string, args ...any) {
			extensionLog(fmt.Sprintf(format, args...))
		},
	})
	if err != nil {
		return nil, err
	}
	server = srv
	return srv, nil
}

// tinyFS is a minimal FSKit file system: a read-only volume holding a
// single empty root directory.
type tinyFS struct{}

var _ fskitbridge.UnaryFileSystem = tinyFS{}

func (tinyFS) Probe(resource fskit.FSResource) (fskitbridge.ProbeResult, error) {
	extensionLog("Go probeResource")
	return fskitbridge.ProbeResult{Name: "TinyFS"}, nil
}

func (tinyFS) Load(resource fskit.FSResource) (fskitbridge.Volume, error) {
	extensionLog("Go loadResource")
	return tinyVolume{}, nil
}

func (tinyFS) Unload() error { return nil }

type tinyVolume struct{}

type tinyItem struct {
	id       fskit.FSItemID
	parentID fskit.FSItemID
	typ      fskit.FSItemType
}

var (
	_ fskitbridge.Volume             = tinyVolume{}
	_ fskitbridge.CapabilitiesVolume = tinyVolume{}
	_ fskitbridge.StatisticsVolume   = tinyVolume{}
)

func (tinyVolume) VolumeName() string { return "TinyFS" }

func (tinyVolume) Root() (fskitbridge.Item, error) {
	return tinyItem{
		id:       fskit.FSItemIDRootDirectory,
		parentID: fskit.FSItemIDParentOfRoot,
		typ:      fskit.FSItemTypeDirectory,
	}, nil
}

func (tinyVolume) Lookup(dir fskitbridge.Item, name string) (fskitbridge.Item, error) {
	return nil, syscall.ENOENT
}

func (tinyVolume) Reclaim(item fskitbridge.Item) {}

func (tinyVolume) Attributes(item fskitbridge.Item) (fskit.FSItemAttributes, error) {
	it, ok := item.(tinyItem)
	if !ok {
		return fskit.FSItemAttributes{}, syscall.EINVAL
	}
	mode := uint32(0444)
	if it.typ == fskit.FSItemTypeDirectory {
		mode = 0555
	}
	return fskitbridge.NewItemAttributes().
		Type(it.typ).
		Mode(mode).
		LinkCount(1).
		UID(uint32(os.Getuid())).
		GID(uint32(os.Getgid())).
		Size(0).
		AllocSize(0).
		ID(it.id).
		ParentID(it.parentID).
		Build(), nil
}

func (tinyVolume) ReadDir(dir fskitbridge.Item) ([]fskitbridge.DirEntry, error) {
	return nil, nil
}

func (tinyVolume) Read(file fskitbridge.Item, offset int64, buf []byte) (int, error) {
	return 0, errors.ErrUnsupported
}

func (tinyVolume) SupportedCapabilities() fskit.FSVolumeSupportedCapabilities {
	capabilities := fskit.NewFSVolumeSupportedCapabilities()
	capabilities.SetSupports64BitObjectIDs(true)
	capabilities.SetSupportsFastStatFS(true)
	capabilities.SetSupportsHiddenFiles(true)
	capabilities.SetDoesNotSupportSettingFilePermissions(true)
	capabilities.SetCaseFormat(fskit.FSVolumeCaseFormatSensitive)
	return capabilities
}

func (tinyVolume) Statistics() fskit.FSStatFSResult {
	result := fskit.NewStatFSResultWithFileSystemTypeName("tinyfs")
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

func main() {
	runtime.LockOSThread()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) > 1 && os.Args[1] == "--extension-main-probe" {
		return extensionMainProbe()
	}
	if len(os.Args) > 1 && os.Args[1] == "--extension-main" {
		return extensionMain()
	}
	if runningInAppExtension() {
		return extensionMain()
	}
	var err error
	objc.AutoreleasePool(func() {
		_, err = ensureServer(0)
	})
	if err != nil {
		return err
	}
	if len(os.Args) > 1 && os.Args[1] == "--smoke" {
		return smoke()
	}
	if len(os.Args) > 1 && os.Args[1] == "--fskit-list" {
		return listModules()
	}
	if len(os.Args) > 2 && os.Args[1] == "--fskit-enable" {
		return setModuleEnabled(os.Args[2], true)
	}
	if len(os.Args) > 2 && os.Args[1] == "--fskit-disable" {
		return setModuleEnabled(os.Args[2], false)
	}
	fmt.Println("tinyfsgo: pure Go FSKit classes registered")
	fmt.Println("tinyfsgo: run with --smoke to exercise probe/load/volume callbacks")
	fmt.Println("tinyfsgo: run with --fskit-list to list FSKit modules")
	fmt.Println("tinyfsgo: run with --extension-main-probe to verify the Swift entrypoint shim")
	return nil
}

func listModules() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	modules, err := fskit.GetFSClientClass().SharedInstance().FetchInstalledExtensions(ctx)
	if err != nil {
		return fmt.Errorf("fetch installed extensions: %w", err)
	}
	for _, module := range modules {
		fmt.Printf("fskit: %s enabled=%v\n", module.BundleIdentifier(), module.IsEnabled())
	}
	return nil
}

func setModuleEnabled(identifier string, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := fskit.GetFSClientClass().SharedInstance().SetEnabledState(ctx, identifier, enabled); err != nil {
		return fmt.Errorf("set enabled state for %s: %w", identifier, err)
	}
	fmt.Printf("fskit: %s enabled=%v\n", identifier, enabled)
	return nil
}

func extensionMainProbe() error {
	if err := loadSwiftShim(); err != nil {
		return err
	}
	if err := registerReplyShims(); err != nil {
		return err
	}
	shimClass := objc.GetClass("GoTinyFSFileSystem")
	if shimClass == 0 {
		return errors.New("swift shim did not register GoTinyFSFileSystem")
	}
	var err error
	objc.AutoreleasePool(func() {
		_, err = ensureServer(shimClass)
	})
	if err != nil {
		return err
	}
	if err := addTinyFileSystemCallbacks(); err != nil {
		return err
	}
	var hasClass func() bool
	if err := registerShimFunc(&hasClass, "TinyFSShimHasFileSystemClass"); err != nil {
		return err
	}
	if !hasClass() {
		return errors.New("TinyFSShimHasFileSystemClass returned false")
	}
	var runMain func()
	if err := registerShimFunc(&runMain, "TinyFSRunExtensionMain"); err != nil {
		return err
	}
	if err := smoke(); err != nil {
		return err
	}
	fmt.Println("tinyfsgo: swift shim loaded")
	fmt.Println("tinyfsgo: attached Go FSKit operation IMPs to GoTinyFSFileSystem")
	fmt.Println("tinyfsgo: resolved TinyFSRunExtensionMain without entering ExtensionFoundation main")
	return nil
}

func extensionMain() error {
	extensionLog("extensionMain start")
	if err := loadSwiftShim(); err != nil {
		return err
	}
	extensionLog("swift shim loaded in extension")
	if err := registerReplyShims(); err != nil {
		return err
	}
	shimClass := objc.GetClass("GoTinyFSFileSystem")
	if shimClass == 0 {
		return errors.New("swift shim did not register GoTinyFSFileSystem")
	}
	var err error
	objc.AutoreleasePool(func() {
		_, err = ensureServer(shimClass)
	})
	if err != nil {
		return err
	}
	if err := addTinyFileSystemCallbacks(); err != nil {
		return err
	}
	extensionLog("Go FSKit IMPs attached in extension")
	if os.Getenv("TINYFSGO_USE_APP_EXTENSION_MAIN") != "1" {
		extensionLog("calling NSExtensionMain")
		return callNSExtensionMain()
	}
	var runMain func()
	if err := registerShimFunc(&runMain, "TinyFSRunExtensionMain"); err != nil {
		return err
	}
	extensionLog("calling TinyFSRunExtensionMain")
	runMain()
	extensionLog("TinyFSRunExtensionMain returned; entering CFRunLoopRun")
	if err := tryPrivateRunningExtensionResume(); err != nil {
		extensionLog("private resume failed: " + err.Error())
	}
	corefoundation.CFRunLoopRun()
	return nil
}

func extensionLog(msg string) {
	line := "tinyfsgo: " + msg
	fmt.Fprintln(os.Stderr, line)
	// The on-disk log + marker files exist to prove an installed extension
	// ran when its stdout/os_log are hard to reach. They probe many
	// candidate directories (including "."), so only write them when actually
	// running inside an app extension — otherwise a local --smoke or CLI run
	// litters the working directory with un-gitignored files.
	if runningInAppExtension() {
		for _, dir := range extensionLogDirs() {
			_ = os.MkdirAll(dir, 0777)
			if f, err := os.OpenFile(filepath.Join(dir, "tinyfsgo-extension.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
				fmt.Fprintln(f, time.Now().Format(time.RFC3339Nano), line)
				f.Close()
				_ = os.WriteFile(filepath.Join(dir, "tinyfsgo-"+markerName(msg)), []byte(line+"\n"), 0666)
				break
			}
		}
	}
	str := foundation.NewConstantStringWithUTF8String(line)
	foundation.NSLog(foundation.NSStringFromID(str.ID))
}

func extensionLogDirs() []string {
	dirs := []string{
		".",
		filepath.Join("Library", "Caches"),
		filepath.Join("Library", "Application Support"),
		"tmp",
	}
	if home, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs,
			filepath.Join(home, "Library", "Caches"),
			filepath.Join(home, "Library", "Application Support"),
			filepath.Join(home, "tmp"),
		)
	}
	dirs = append(dirs,
		"/Users/tmc/Library/Containers/dev.tmc.apple.examples.fskit.tinyfs.extension/Data/Library/Caches",
		"/Users/tmc/Library/Containers/dev.tmc.apple.examples.fskit.tinyfs.extension/Data/Library/Application Support",
		"/Users/tmc/Library/Containers/dev.tmc.apple.examples.fskit.tinyfs.extension/Data/tmp",
		"/Users/tmc/Library/Containers/dev.tmc.apple.examples.fskit.tinyfs/Data/Library/Caches",
		"/Users/tmc/Library/Containers/dev.tmc.apple.examples.fskit.tinyfs/Data/tmp",
	)
	if dir, err := os.UserCacheDir(); err == nil {
		dirs = append(dirs, dir)
	}
	if dir, err := os.UserConfigDir(); err == nil {
		dirs = append(dirs, dir)
	}
	dirs = append(dirs, os.TempDir())
	return dirs
}

func markerName(msg string) string {
	name := strings.ToLower(msg)
	var b strings.Builder
	for _, r := range name {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			b.WriteRune(r)
			continue
		}
		if b.Len() > 0 && b.String()[b.Len()-1] != '-' {
			b.WriteByte('-')
		}
	}
	out := strings.Trim(b.String(), "-")
	if out == "" {
		out = "event"
	}
	return out + ".marker"
}

func registerShimFunc(fptr any, name string) error {
	sym, err := purego.Dlsym(shimHandle, name)
	if err != nil {
		return fmt.Errorf("resolve %s: %w", name, err)
	}
	purego.RegisterFunc(fptr, sym)
	return nil
}

func registerReplyShims() error {
	if replyProbe == nil {
		if err := registerShimFunc(&replyProbe, "TinyFSShimReplyProbe"); err != nil {
			return err
		}
	}
	if shimVolume == nil {
		if err := registerShimFunc(&shimVolume, "TinyFSShimNewVolume"); err != nil {
			return err
		}
	}
	return nil
}

func loadSwiftShim() error {
	path := os.Getenv("TINYFSGO_SWIFTSHIM")
	if path == "" {
		var err error
		path, err = defaultSwiftShimPath()
		if err != nil {
			return err
		}
	}
	handle, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("load swift shim %s: %w", path, err)
	}
	shimHandle = handle
	return nil
}

func defaultSwiftShimPath() (string, error) {
	candidates := []string{
		filepath.Join("examples", "fskit", "tinyfsgo", "swiftshim", ".build", "release", "libTinyFSShim.dylib"),
		filepath.Join("examples", "fskit", "tinyfsgo", "swiftshim", ".build", "arm64-apple-macosx", "release", "libTinyFSShim.dylib"),
		filepath.Join("examples", "fskit", "tinyfsgo", "swiftshim", ".build", "x86_64-apple-macosx", "release", "libTinyFSShim.dylib"),
	}
	if exe, err := os.Executable(); err == nil {
		contents := filepath.Dir(filepath.Dir(exe))
		candidates = append(candidates,
			filepath.Join(contents, "Frameworks", "libTinyFSShim.dylib"),
			filepath.Join(contents, "MacOS", "libTinyFSShim.dylib"),
		)
	}
	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("swift shim not built; run: (cd examples/fskit/tinyfsgo/swiftshim && swift build -c release --product TinyFSShim)")
}

func runningInAppExtension() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	for dir := filepath.Clean(exe); dir != "." && dir != string(filepath.Separator); dir = filepath.Dir(dir) {
		if filepath.Ext(dir) == ".appex" {
			return true
		}
	}
	return false
}

// addTinyFileSystemCallbacks routes the Swift shim's file system hooks to
// Go. The shim's class is the registered file system class, so the
// operation selectors stay with the shim and call back into Go here.
func addTinyFileSystemCallbacks() error {
	callbacks := []struct {
		name string
		fn   any
	}{
		{"TinyFSShimSetProbeResourceResultCallback", tinyFSProbeResourceResult},
		{"TinyFSShimSetLoadResourceResultCallback", tinyFSLoadResourceResult},
		{"TinyFSShimSetUnloadResourceResultCallback", tinyFSUnloadResourceResult},
		{"TinyFSShimSetDidFinishLoadingCallback", tinyFSDidFinishLoading},
		{"TinyFSShimSetVolumeActivateCallback", tinyShimVolumeActivate},
		{"TinyFSShimSetVolumeMountCallback", tinyShimVolumeMount},
		{"TinyFSShimSetVolumeUnmountCallback", tinyShimVolumeUnmount},
	}
	for _, callback := range callbacks {
		var set func(uintptr)
		if err := registerShimFunc(&set, callback.name); err != nil {
			return err
		}
		set(purego.NewCallback(callback.fn))
	}
	return nil
}

func tinyShimVolumeActivate(self objc.ID, _ objc.SEL) {
	extensionLog("Go activate")
}

func tinyShimVolumeMount(self objc.ID, _ objc.SEL) {
	extensionLog("Go mount")
}

func tinyShimVolumeUnmount(self objc.ID, _ objc.SEL) {
	extensionLog("Go unmount")
}

func tinyFSDidFinishLoading(self objc.ID, _ objc.SEL) {
	extensionLog("Go didFinishLoading")
}

func tinyFSProbeResourceResult(self objc.ID, _ objc.SEL, resource objc.ID, errp unsafe.Pointer) objc.ID {
	extensionLog("Go probeResource")
	containerID := fskit.NewContainerIdentifierWithUUID(foundation.NewNSUUID())
	result := fskit.GetFSProbeResultClass().UsableProbeResultWithNameContainerID("TinyFS", containerID)
	return result.GetID()
}

func tinyFSLoadResourceResult(self objc.ID, _ objc.SEL, resource objc.ID, options objc.ID, errp unsafe.Pointer) objc.ID {
	extensionLog("Go loadResource")
	volume := newVolume()
	extensionLog("Go loadResource created volume")
	fskit.FSUnaryFileSystemFromID(self).SetContainerStatus(fskit.GetFSContainerStatusClass().Ready())
	extensionLog("Go loadResource set ready status")
	return volume
}

func tinyFSUnloadResourceResult(self objc.ID, _ objc.SEL, resource objc.ID, options objc.ID) objc.ID {
	return 0
}

// newVolume returns the volume to serve: the Swift shim's volume when the
// shim provides one, otherwise a bridge volume serving tinyVolume.
func newVolume() objc.ID {
	if shimVolume != nil {
		return shimVolume()
	}
	srv := currentServer()
	if srv == nil {
		return 0
	}
	volume, err := srv.NewVolume(tinyVolume{})
	if err != nil {
		extensionLog("new volume: " + err.Error())
		return 0
	}
	return volume
}

func currentServer() *fskitbridge.Server {
	serverMu.Lock()
	defer serverMu.Unlock()
	return server
}

func smoke() error {
	var err error
	objc.AutoreleasePool(func() {
		err = smokeInPool()
	})
	return err
}

func smokeInPool() error {
	srv := currentServer()
	if srv == nil {
		return errors.New("bridge server not registered")
	}
	var callbackErr error
	fs := srv.NewFileSystem()
	defer objc.Send[struct{}](fs, objc.Sel("release"))

	var probeName string
	probeReply := objc.NewBlock(func(_ objc.Block, result objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("probe returned error %d", errID)
			return
		}
		probeName = fskit.FSProbeResultFromID(result).Name()
	})
	defer probeReply.Release()
	objc.Send[struct{}](fs, objc.Sel("probeResource:replyHandler:"), objc.ID(0), objc.ID(probeReply))
	if probeName != "TinyFS" {
		return fmt.Errorf("probe name = %q, want TinyFS", probeName)
	}

	var volume objc.ID
	loadReply := objc.NewBlock(func(_ objc.Block, result objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("load returned error %d", errID)
			return
		}
		volume = result
	})
	defer loadReply.Release()
	objc.Send[struct{}](fs, objc.Sel("loadResource:options:replyHandler:"), objc.ID(0), objc.ID(0), objc.ID(loadReply))
	if callbackErr != nil {
		return callbackErr
	}
	if volume == 0 {
		return errors.New("load returned nil volume")
	}
	if state := fskit.FSUnaryFileSystemFromID(fs).ContainerStatus().State(); state != fskit.FSContainerStateReady {
		return fmt.Errorf("container state = %v, want ready", state)
	}

	var root objc.ID
	activateReply := objc.NewBlock(func(_ objc.Block, item objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("activate returned error %d", errID)
			return
		}
		root = item
	})
	defer activateReply.Release()
	objc.Send[struct{}](volume, objc.Sel("activateWithOptions:replyHandler:"), objc.ID(0), objc.ID(activateReply))
	if callbackErr != nil {
		return callbackErr
	}
	if root == 0 {
		return errors.New("activate returned nil root")
	}

	var mode uint32
	attrsReply := objc.NewBlock(func(_ objc.Block, attrs objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("get attributes returned error %d", errID)
			return
		}
		mode = fskit.FSItemAttributesFromID(attrs).Mode()
	})
	defer attrsReply.Release()
	objc.Send[struct{}](volume, objc.Sel("getAttributes:ofItem:replyHandler:"), objc.ID(0), root, objc.ID(attrsReply))
	if callbackErr != nil {
		return callbackErr
	}
	if mode != 0555 {
		return fmt.Errorf("root mode = %#o, want 0555", mode)
	}

	fmt.Println("tinyfsgo: smoke ok")
	fmt.Println("tinyfsgo: registered FSUnaryFileSystem, FSVolume, and FSItem subclasses from pure Go")
	fmt.Println("tinyfsgo: exercised FSKit reply blocks without cgo")
	return nil
}
