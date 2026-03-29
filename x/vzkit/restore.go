package vzkit

import (
	"context"

	restorex "github.com/tmc/apple/x/vzkit/restore"
)

// RestoreImageInfo describes a macOS restore image (.ipsw).
type RestoreImageInfo = restorex.ImageInfo

// FetchLatestRestoreImage fetches information about the latest supported restore image.
func FetchLatestRestoreImage(ctx context.Context) (RestoreImageInfo, error) {
	return restorex.FetchLatest(ctx)
}

// LoadRestoreImage loads a macOS restore image from a local .ipsw file.
func LoadRestoreImage(ctx context.Context, path string) (RestoreImageInfo, error) {
	return restorex.Load(ctx, path)
}
