package fskitbridge

import (
	"errors"
	"fmt"

	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objc/objcbridge"
)

// ClassSet is the set of Objective-C classes that implements an FSKit
// file system.
type ClassSet struct {
	FileSystem objc.Class
	Volume     objc.Class
	Item       objc.Class
}

// ClassConfig describes the Objective-C classes and method implementations
// for an FSKit file system.
//
// If ExistingFileSystem is nonzero it is used as the file system class as is
// and FileSystemMethods are not attached; this supports extensions whose
// file system class is provided by a host shim that routes operations to Go
// itself. If a class named VolumeName or ItemName is already registered,
// its methods are added or replaced instead of registering a new class, so
// registration can run again in the same process.
type ClassConfig struct {
	FileSystemName      string
	VolumeName          string
	ItemName            string
	ExistingFileSystem  objc.Class
	FileSystemProtocols []string
	VolumeProtocols     []string
	FileSystemMethods   []objc.MethodDef
	VolumeMethods       []objc.MethodDef
	ItemMethods         []objc.MethodDef
}

// RegisterClasses registers or extends the Objective-C class set for an
// FSKit file system.
func RegisterClasses(cfg ClassConfig) (ClassSet, error) {
	switch {
	case cfg.FileSystemName == "":
		return ClassSet{}, errors.New("missing file system class name")
	case cfg.VolumeName == "":
		return ClassSet{}, errors.New("missing volume class name")
	case cfg.ItemName == "":
		return ClassSet{}, errors.New("missing item class name")
	}
	var set ClassSet
	var err error

	set.Item = objc.GetClass(cfg.ItemName)
	if set.Item != 0 {
		if err := objcbridge.AddMethods(set.Item, cfg.ItemName, cfg.ItemMethods); err != nil {
			return ClassSet{}, err
		}
	} else {
		set.Item, err = objc.RegisterClass(cfg.ItemName, fskit.GetFSItemClass().Class(), nil, nil, cfg.ItemMethods)
		if err != nil {
			return ClassSet{}, fmt.Errorf("register item class %s: %w", cfg.ItemName, err)
		}
	}

	set.Volume = objc.GetClass(cfg.VolumeName)
	volumeProtocols := objcbridge.ProtocolsByName(cfg.VolumeProtocols...)
	if set.Volume != 0 {
		if err := objcbridge.AddMethods(set.Volume, cfg.VolumeName, cfg.VolumeMethods); err != nil {
			return ClassSet{}, err
		}
	} else {
		set.Volume, err = objc.RegisterClass(cfg.VolumeName, fskit.GetFSVolumeClass().Class(), volumeProtocols, nil, cfg.VolumeMethods)
		if err != nil {
			return ClassSet{}, fmt.Errorf("register volume class %s: %w", cfg.VolumeName, err)
		}
	}

	if cfg.ExistingFileSystem != 0 {
		set.FileSystem = cfg.ExistingFileSystem
	} else {
		fsProtocols := objcbridge.ProtocolsByName(cfg.FileSystemProtocols...)
		set.FileSystem, err = objc.RegisterClass(cfg.FileSystemName, fskit.GetFSUnaryFileSystemClass().Class(), fsProtocols, nil, cfg.FileSystemMethods)
		if err != nil {
			return ClassSet{}, fmt.Errorf("register file system class %s: %w", cfg.FileSystemName, err)
		}
	}
	return set, nil
}
