package protocol

// VoxelCells represents a 3D grid of voxel cell data.
// Added: v1.26.0
type VoxelCells struct {
	// XSize is the size of the grid along the X axis.
	// Added: v1.26.0
	XSize uint8
	// YSize is the size of the grid along the Y axis.
	// Added: v1.26.0
	YSize uint8
	// ZSize is the size of the grid along the Z axis.
	// Added: v1.26.0
	ZSize uint8
	// Storage is the raw cell data stored in the grid.
	// Added: v1.26.0
	Storage []uint8
}

// Marshal encodes/decodes a VoxelCells.
func (x *VoxelCells) Marshal(r IO) {
	r.Uint8(&x.XSize)
	r.Uint8(&x.YSize)
	r.Uint8(&x.ZSize)
	FuncSlice(r, &x.Storage, r.Uint8)
}

// VoxelShapeNameEntry represents a name-to-ID mapping entry for voxel shapes.
// Added: v1.26.0
type VoxelShapeNameEntry struct {
	// Name is the name of the voxel shape.
	// Added: v1.26.0
	Name string
	// ID is the numeric ID of the voxel shape.
	// Added: v1.26.0
	ID uint16
}

// Marshal encodes/decodes a VoxelShapeNameEntry.
func (x *VoxelShapeNameEntry) Marshal(r IO) {
	r.String(&x.Name)
	r.Uint16(&x.ID)
}

// VoxelShape represents a voxel shape with cells and coordinate axes.
// Added: v1.26.0
type VoxelShape struct {
	// Cells is the grid of cells representing solid and empty regions.
	// Added: v1.26.0
	Cells VoxelCells
	// XCoordinates is a list of X axis coordinates for the shape.
	// Added: v1.26.0
	XCoordinates []float32
	// YCoordinates is a list of Y axis coordinates for the shape.
	// Added: v1.26.0
	YCoordinates []float32
	// ZCoordinates is a list of Z axis coordinates for the shape.
	// Added: v1.26.0
	ZCoordinates []float32
}

// Marshal encodes/decodes a VoxelShape.
func (x *VoxelShape) Marshal(r IO) {
	Single(r, &x.Cells)
	FuncSlice(r, &x.XCoordinates, r.Float32)
	FuncSlice(r, &x.YCoordinates, r.Float32)
	FuncSlice(r, &x.ZCoordinates, r.Float32)
}
