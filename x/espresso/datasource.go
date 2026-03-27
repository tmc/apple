//go:build darwin

package espresso

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	pe "github.com/tmc/apple/private/espresso"
)

// DataSource wraps an Espresso data source for providing training
// or inference data to the training loop.
type DataSource struct {
	buf    pe.ETDataSourceBuf
	folder pe.ETDataSourceFromFolderData
	cached pe.ETDataSourceWithCache

	// obj is the underlying object to pass to training methods.
	obj objectivec.Object
}

// NewBufferDataSource creates an in-memory data source.
// Use SetBlobs to populate it with tensor data.
func NewBufferDataSource() *DataSource {
	buf := pe.NewETDataSourceBuf()
	return &DataSource{
		buf: buf,
		obj: objectivec.Object{ID: buf.GetID()},
	}
}

// FromFolder creates a data source from a directory of images
// organized into class subdirectories.
func FromFolder(path string, balanceClasses bool) *DataSource {
	nsPath := objectivec.Object{ID: objc.String(path)}
	folder := pe.NewETDataSourceFromFolderDataWithFolderBalanceClassesForTraining(nsPath, balanceClasses)
	return &DataSource{
		folder: folder,
		obj:    objectivec.Object{ID: folder.GetID()},
	}
}

// WithCache wraps a data source with caching.
func (d *DataSource) WithCache() *DataSource {
	cached := pe.NewETDataSourceWithCacheWithDataSource(d.obj)
	return &DataSource{
		cached: cached,
		obj:    objectivec.Object{ID: cached.GetID()},
	}
}

// Count returns the number of data points.
func (d *DataSource) Count() int {
	if d.buf.GetID() != 0 {
		return d.buf.NumberOfDataPoints()
	}
	if d.folder.GetID() != 0 {
		return d.folder.NumberOfDataPoints()
	}
	if d.cached.GetID() != 0 {
		return d.cached.NumberOfDataPoints()
	}
	return 0
}

// NumClasses returns the number of classes (folder data sources only).
func (d *DataSource) NumClasses() int {
	if d.folder.GetID() != 0 {
		return d.folder.NumberOfClasses()
	}
	return 0
}

// ClassNames returns the class label names (folder data sources only).
func (d *DataSource) ClassNames() []string {
	if d.folder.GetID() == 0 {
		return nil
	}
	arr := d.folder.ClassNames()
	if arr == nil {
		return nil
	}
	return nsArrayToStrings(arr)
}

// ObjcID returns the underlying ObjC object pointer.
func (d *DataSource) ObjcID() uintptr { return uintptr(d.obj.ID) }
