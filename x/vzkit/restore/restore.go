package restore

import (
	"context"
	"fmt"

	"github.com/tmc/apple/foundation"
	vz "github.com/tmc/apple/virtualization"
)

// ImageInfo describes a macOS restore image.
type ImageInfo struct {
	URL           string
	BuildVersion  string
	OSVersion     foundation.NSOperatingSystemVersion
	Supported     bool
	MinCPUCount   uint
	MinMemorySize uint64
	Image         vz.VZMacOSRestoreImage
}

func imageInfoFrom(img vz.VZMacOSRestoreImage) ImageInfo {
	info := ImageInfo{
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

// FetchLatest fetches the latest restore image supported by the current host.
func FetchLatest(ctx context.Context) (ImageInfo, error) {
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
			return ImageInfo{}, fmt.Errorf("fetch latest restore image: %w", r.err)
		}
		if r.img == nil {
			return ImageInfo{}, fmt.Errorf("fetch latest restore image: nil result")
		}
		return imageInfoFrom(*r.img), nil
	case <-ctx.Done():
		return ImageInfo{}, ctx.Err()
	}
}

// Load loads a restore image from a local IPSW.
func Load(ctx context.Context, path string) (ImageInfo, error) {
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
			return ImageInfo{}, fmt.Errorf("load restore image: %w", r.err)
		}
		if r.img == nil {
			return ImageInfo{}, fmt.Errorf("load restore image: nil result")
		}
		return imageInfoFrom(*r.img), nil
	case <-ctx.Done():
		return ImageInfo{}, ctx.Err()
	}
}
