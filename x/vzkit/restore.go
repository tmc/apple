package vzkit

import (
	"context"
	"fmt"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// RestoreImageInfo describes a macOS restore image (.ipsw) that can be used
// to install macOS on a virtual machine.
type RestoreImageInfo struct {
	// URL is the location of the restore image. For images fetched from
	// the network, this is a remote URL. For local images, it is a file URL.
	URL string

	// BuildVersion is the build version string (e.g., "22A380").
	BuildVersion string

	// OSVersion is the macOS version contained in this image.
	OSVersion foundation.NSOperatingSystemVersion

	// Supported reports whether the current host can run this image.
	Supported bool

	// MinCPUCount is the minimum number of CPUs required by the most
	// featureful supported configuration, or 0 if unavailable.
	MinCPUCount uint

	// MinMemorySize is the minimum memory in bytes required by the most
	// featureful supported configuration, or 0 if unavailable.
	MinMemorySize uint64

	// Image is the underlying VZMacOSRestoreImage for use with VZMacOSInstaller.
	Image vz.VZMacOSRestoreImage
}

func restoreImageInfoFrom(img vz.VZMacOSRestoreImage) RestoreImageInfo {
	info := RestoreImageInfo{
		BuildVersion: img.BuildVersion(),
		OSVersion:    img.OperatingSystemVersion(),
		Supported:    img.Supported(),
		Image:        img,
	}

	if u := img.URL(); u != nil && u.GetID() != 0 {
		nsurl := foundation.NSURLFromID(u.GetID())
		info.URL = nsurl.AbsoluteString()
	}

	if cfg := img.MostFeaturefulSupportedConfiguration(); cfg != nil && cfg.GetID() != 0 {
		reqs := vz.VZMacOSConfigurationRequirementsFromID(cfg.GetID())
		info.MinCPUCount = reqs.MinimumSupportedCPUCount()
		info.MinMemorySize = reqs.MinimumSupportedMemorySize()
	}

	return info
}

// FetchLatestRestoreImage fetches information about the latest macOS restore
// image supported by this host from Apple's servers. The returned info
// contains a network URL; use the URL to download the .ipsw before installing.
func FetchLatestRestoreImage(ctx context.Context) (RestoreImageInfo, error) {
	type result struct {
		img *vz.VZMacOSRestoreImage
		err error
	}
	ch := make(chan result, 1)

	cls := vz.GetVZMacOSRestoreImageClass()
	cls.FetchLatestSupportedWithCompletionHandler(func(img *vz.VZMacOSRestoreImage, err error) {
		ch <- result{img, err}
	})

	select {
	case r := <-ch:
		if r.err != nil {
			return RestoreImageInfo{}, fmt.Errorf("fetch latest restore image: %w", r.err)
		}
		if r.img == nil {
			return RestoreImageInfo{}, fmt.Errorf("fetch latest restore image: nil result")
		}
		return restoreImageInfoFrom(*r.img), nil
	case <-ctx.Done():
		return RestoreImageInfo{}, ctx.Err()
	}
}

// LoadRestoreImage loads a macOS restore image from a local .ipsw file.
func LoadRestoreImage(ctx context.Context, path string) (RestoreImageInfo, error) {
	type result struct {
		img *vz.VZMacOSRestoreImage
		err error
	}
	ch := make(chan result, 1)

	fileURL := foundation.NewURLFileURLWithPath(path)
	cls := vz.GetVZMacOSRestoreImageClass()
	cls.LoadFileURLCompletionHandler(fileURL, func(img *vz.VZMacOSRestoreImage, err error) {
		ch <- result{img, err}
	})

	select {
	case r := <-ch:
		if r.err != nil {
			return RestoreImageInfo{}, fmt.Errorf("load restore image: %w", r.err)
		}
		if r.img == nil {
			return RestoreImageInfo{}, fmt.Errorf("load restore image: nil result")
		}
		return restoreImageInfoFrom(*r.img), nil
	case <-ctx.Done():
		return RestoreImageInfo{}, ctx.Err()
	}
}
