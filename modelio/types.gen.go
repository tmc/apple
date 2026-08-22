// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

// C struct types

// MDLAxisAlignedBoundingBox - The minimal volume containing an object, used by the [boundingBox(atTime:)](<https://developer.apple.com/documentation/ModelIO/MDLObject/boundingBox(atTime:)>) method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLAxisAlignedBoundingBox
type MDLAxisAlignedBoundingBox struct {
	MaxBounds float32 // The corner of the bounding box with the highest x-, y-, and z-coordinate values.
	MinBounds float32 // The corner of the bounding box with the lowest x-, y-, and z-coordinate values.

}

// MDLVoxelIndexExtent - The corner voxel indices defining a solid rectangular volume of voxels. Used by the [voxelIndexExtent](<https://developer.apple.com/documentation/ModelIO/MDLVoxelArray/voxelIndexExtent>) property and [voxels(within:)](<https://developer.apple.com/documentation/ModelIO/MDLVoxelArray/voxels(within:)>) method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVoxelIndexExtent
type MDLVoxelIndexExtent struct {
	MinimumExtent MDLVoxelIndex // The lowest x, y, and z coordinates in the volume.
	MaximumExtent MDLVoxelIndex // The highest x, y, and z coordinates in the volume.

}
