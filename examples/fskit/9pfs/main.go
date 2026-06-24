//go:build darwin

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"9fans.net/go/plan9/client"
)

func cliMain() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var dialect, network, addr, aname string
	var lsPath, catPath, createPath, mkdirPath, writePath, writeData, removePath, renameFrom, renameTo string
	var statPath, chmodPath, chmodMode, truncatePath, truncateSize, mtimePath, mtimeValue string
	var symlinkPath, symlinkTarget, readlinkPath, linkFrom, linkTo string
	var xattrPath, xattrName, xattrValue, getXattrPath, listXattrPath, removeXattrPath string
	flag.StringVar(&dialect, "dialect", "9p2000", "9p dialect: 9p2000 or 9p2000l")
	flag.StringVar(&network, "net", "tcp", "network for the 9p server")
	flag.StringVar(&addr, "addr", "127.0.0.1:5640", "address for the 9p server")
	flag.StringVar(&aname, "aname", "", "9p attach name")
	flag.StringVar(&lsPath, "ls", "/", "list a directory on the 9p server")
	flag.StringVar(&catPath, "cat", "", "read a file on the 9p server")
	flag.StringVar(&createPath, "create", "", "create a file on the 9p server")
	flag.StringVar(&mkdirPath, "mkdir", "", "create a directory on the 9p server")
	flag.StringVar(&writePath, "write", "", "write data to a file on the 9p server")
	flag.StringVar(&writeData, "data", "", "data for -write")
	flag.StringVar(&removePath, "rm", "", "remove a file on the 9p server")
	flag.StringVar(&renameFrom, "rename", "", "rename a path on the 9p server")
	flag.StringVar(&renameTo, "to", "", "destination for -rename")
	flag.StringVar(&statPath, "stat", "", "stat a path on the 9p server")
	flag.StringVar(&chmodPath, "chmod", "", "change file mode on the 9p server")
	flag.StringVar(&chmodMode, "mode", "", "octal mode for -chmod")
	flag.StringVar(&truncatePath, "truncate", "", "truncate a file on the 9p server")
	flag.StringVar(&truncateSize, "size", "", "size for -truncate")
	flag.StringVar(&mtimePath, "mtime", "", "set modification time on the 9p server")
	flag.StringVar(&mtimeValue, "mtime-seconds", "", "Unix seconds for -mtime")
	flag.StringVar(&symlinkPath, "symlink", "", "create a symlink on the 9p server")
	flag.StringVar(&symlinkTarget, "target", "", "target for -symlink")
	flag.StringVar(&readlinkPath, "readlink", "", "read a symlink on the 9p server")
	flag.StringVar(&linkFrom, "link", "", "create a hard link to this path")
	flag.StringVar(&linkTo, "link-to", "", "destination for -link")
	flag.StringVar(&xattrPath, "setxattr", "", "set an extended attribute on this path")
	flag.StringVar(&xattrName, "xattr", "", "extended attribute name for xattr operations")
	flag.StringVar(&xattrValue, "xattr-data", "", "extended attribute value for -setxattr")
	flag.StringVar(&getXattrPath, "getxattr", "", "get an extended attribute from this path")
	flag.StringVar(&listXattrPath, "listxattr", "", "list extended attributes on this path")
	flag.StringVar(&removeXattrPath, "rmxattr", "", "remove an extended attribute from this path")
	fskitSmokeFlag := flag.Bool("fskit-smoke", false, "exercise FSKit volume callbacks with an in-memory 9p tree")
	extensionMainProbeFlag := flag.Bool("extension-main-probe", false, "verify the Swift FSKit entrypoint shim without entering ExtensionFoundation main")
	flag.Parse()

	if *extensionMainProbeFlag {
		return extensionMainProbe()
	}
	if *fskitSmokeFlag {
		return fskitSmoke()
	}

	fs, err := dialBackend(dialect, network, addr, aname)
	if err != nil {
		return err
	}
	defer fs.Close()

	if createPath != "" {
		if _, err := fs.Create(createPath, 0666, false); err != nil {
			return err
		}
	}
	if mkdirPath != "" {
		if _, err := fs.Create(mkdirPath, 0777, true); err != nil {
			return err
		}
	}
	if writePath != "" {
		if _, err := fs.WriteFile(writePath, 0, []byte(writeData)); err != nil {
			return err
		}
	}
	if renameFrom != "" {
		if renameTo == "" {
			return fmt.Errorf("-rename requires -to")
		}
		if err := fs.Rename(renameFrom, renameTo); err != nil {
			return err
		}
	}
	if chmodPath != "" {
		if chmodMode == "" {
			return fmt.Errorf("-chmod requires -mode")
		}
		mode64, err := strconv.ParseUint(chmodMode, 8, 32)
		if err != nil {
			return fmt.Errorf("parse -mode: %w", err)
		}
		mode := uint32(mode64)
		if _, err := fs.SetAttr(chmodPath, setAttr{Mode: &mode}); err != nil {
			return err
		}
	}
	if truncatePath != "" {
		if truncateSize == "" {
			return fmt.Errorf("-truncate requires -size")
		}
		size, err := strconv.ParseUint(truncateSize, 10, 64)
		if err != nil {
			return fmt.Errorf("parse -size: %w", err)
		}
		if _, err := fs.SetAttr(truncatePath, setAttr{Size: &size}); err != nil {
			return err
		}
	}
	if mtimePath != "" {
		if mtimeValue == "" {
			return fmt.Errorf("-mtime requires -mtime-seconds")
		}
		modified, err := strconv.ParseUint(mtimeValue, 10, 64)
		if err != nil {
			return fmt.Errorf("parse -mtime-seconds: %w", err)
		}
		if _, err := fs.SetAttr(mtimePath, setAttr{Modified: &modified}); err != nil {
			return err
		}
	}
	if symlinkPath != "" {
		if symlinkTarget == "" {
			return fmt.Errorf("-symlink requires -target")
		}
		if _, err := fs.CreateSymlink(symlinkPath, symlinkTarget); err != nil {
			return err
		}
	}
	if linkFrom != "" {
		if linkTo == "" {
			return fmt.Errorf("-link requires -link-to")
		}
		if _, err := fs.CreateLink(linkFrom, linkTo); err != nil {
			return err
		}
	}
	if xattrPath != "" {
		if xattrName == "" {
			return fmt.Errorf("-setxattr requires -xattr")
		}
		if err := fs.SetXattr(xattrPath, xattrName, []byte(xattrValue)); err != nil {
			return err
		}
	}
	if removeXattrPath != "" {
		if xattrName == "" {
			return fmt.Errorf("-rmxattr requires -xattr")
		}
		if err := fs.RemoveXattr(removeXattrPath, xattrName); err != nil {
			return err
		}
	}
	if removePath != "" {
		if err := fs.Remove(removePath); err != nil {
			return err
		}
	}
	if readlinkPath != "" {
		target, err := fs.Readlink(readlinkPath)
		if err != nil {
			return err
		}
		fmt.Println(target)
		return nil
	}
	if getXattrPath != "" {
		if xattrName == "" {
			return fmt.Errorf("-getxattr requires -xattr")
		}
		data, err := fs.GetXattr(getXattrPath, xattrName)
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(data)
		return err
	}
	if listXattrPath != "" {
		names, err := fs.ListXattr(listXattrPath)
		if err != nil {
			return err
		}
		if len(names) > 0 {
			fmt.Println(strings.Join(names, "\n"))
		}
		return nil
	}
	if statPath != "" {
		info, err := fs.Stat(statPath)
		if err != nil {
			return err
		}
		fmt.Printf("%s %s %d %d %04o\n", modeString(info.Mode), info.Name, info.Length, info.Modified, info.Mode&0777)
		return nil
	}
	if catPath != "" {
		data, err := fs.ReadFile(catPath)
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(data)
		return err
	}

	entries, err := fs.ReadDir(lsPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Printf("%s %s %d\n", modeString(entry.Mode), entry.Name, entry.Length)
	}
	return nil
}

func mount9P(network, addr, aname string) (*client.Fsys, error) {
	if aname == "" {
		fs, err := client.Mount(network, addr)
		if err != nil {
			return nil, fmt.Errorf("mount 9p %s %s: %w", network, addr, err)
		}
		return fs, nil
	}
	conn, err := client.Dial(network, addr)
	if err != nil {
		return nil, fmt.Errorf("dial 9p %s %s: %w", network, addr, err)
	}
	fs, err := conn.Attach(nil, os.Getenv("USER"), aname)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("attach 9p aname %q: %w", aname, err)
	}
	return fs, nil
}
