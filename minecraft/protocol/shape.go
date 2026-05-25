package protocol

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	ShapeDataLast = iota
	ShapeDataArrow
	ShapeDataText
	ShapeDataBox
	ShapeDataLine
	ShapeDataSphere
)

// lookupShapeData looks up an ShapeData matching the shape data type passed.
// False is returned if no such shape data exists.
func lookupShapeData(shapeDataType uint32, x *ShapeData) bool {
	switch shapeDataType {
	case ShapeDataLast:
		*x = &LastShape{}
	case ShapeDataArrow:
		*x = &ArrowShape{}
	case ShapeDataText:
		*x = &TextShape{}
	case ShapeDataBox:
		*x = &BoxShape{}
	case ShapeDataLine:
		*x = &LineShape{}
	case ShapeDataSphere:
		*x = &SphereShape{}
	default:
		return false
	}
	return true
}

// lookupShapeDataType looks up a debug shape type that matches the ShapeData passed.
func lookupShapeDataType(x ShapeData, shapeDataType *uint32) bool {
	switch x.(type) {
	case *LastShape:
		*shapeDataType = ShapeDataLast
	case *ArrowShape:
		*shapeDataType = ShapeDataArrow
	case *TextShape:
		*shapeDataType = ShapeDataText
	case *BoxShape:
		*shapeDataType = ShapeDataBox
	case *LineShape:
		*shapeDataType = ShapeDataLine
	case *SphereShape:
		*shapeDataType = ShapeDataSphere
	default:
		return false
	}
	return true
}

// ShapeData represents an object that holds data specific to a debug shape.
// The data it holds depends on the type.
//
// Added: v1.26.0
type ShapeData interface {
	// Marshal encodes/decodes a serialised debug shape data object.
	Marshal(r IO)
}

// LastShape points to using the last shape settings.
// If no shape has ever been set, then use the default one.
//
// Added: v1.26.0
type LastShape struct{}

// Marshal ...
func (shape *LastShape) Marshal(io IO) {}

// LineShape represents a line debug shape.
//
// Added: v1.26.0
type LineShape struct {
	// LineEndLocation is the line end location of the shape.
	//
	// Added: v1.26.0
	LineEndLocation mgl32.Vec3
}

// Marshal ...
func (shape *LineShape) Marshal(io IO) {
	io.Vec3(&shape.LineEndLocation)
}

// TextShape represents a text debug shape.
//
// Added: v1.26.0
type TextShape struct {
	// Text is the text of the debug text shape.
	//
	// Added: v1.26.0
	Text string
	// UseRotation is if the text should use the provided rotation, meaning it will be static and does not follow the
	// camera.
	//
	// Added: v1.26.20
	UseRotation bool
	// BackgroundColour is the optional RGBA colour to use for the text background.
	//
	// Added: v1.26.20
	BackgroundColour Optional[color.RGBA]
	// DepthTest determines whether the text is depth-tested against world geometry.
	//
	// Added: v1.26.20
	DepthTest bool
	// ShowBackface is if the background should render on the back side of the shape. This only has a visible effect when
	// UseRotation is true since you cannot see the back side of the text otherwise.
	//
	// Added: v1.26.20
	ShowBackface bool
}

// Marshal ...
func (shape *TextShape) Marshal(io IO) {
	io.String(&shape.Text)
	if io.Protocol() >= Protocol1v26v20 {
		io.Bool(&shape.UseRotation)
		OptionalFunc(io, &shape.BackgroundColour, io.ARGB)
		io.Bool(&shape.DepthTest)
		io.Bool(&shape.ShowBackface)
	}
}

// BoxShape represents a box debug shape.
//
// Added: v1.26.0
type BoxShape struct {
	// BoxBound is the box bound of the shape.
	//
	// Added: v1.26.0
	BoxBound mgl32.Vec3
}

// Marshal ...
func (shape *BoxShape) Marshal(io IO) {
	io.Vec3(&shape.BoxBound)
}

// SphereShape represents a circle or sphere debug shape.
//
// Added: v1.26.0
type SphereShape struct {
	// Segments is the segments that used for the debug circle or sphere.
	//
	// Added: v1.26.0
	Segments byte
}

// Marshal ...
func (shape *SphereShape) Marshal(io IO) {
	io.Uint8(&shape.Segments)
}

// ArrowShape represents an arrow debug shape.
//
// Added: v1.26.0
type ArrowShape struct {
	// ArrowEndLocation is the arrow end location of the shape.
	//
	// Added: v1.26.0
	ArrowEndLocation Optional[mgl32.Vec3]
	// ArrowHeadLength is the arrow head length of the shape.
	//
	// Added: v1.26.0
	ArrowHeadLength Optional[float32]
	// ArrowHeadRadius is the arrow head radius of the shape.
	//
	// Added: v1.26.0
	ArrowHeadRadius Optional[float32]
	// Segments is the segments that used for the debug arrow's head.
	//
	// Added: v1.26.0
	Segments Optional[byte]
}

// Marshal ...
func (shape *ArrowShape) Marshal(io IO) {
	OptionalFunc(io, &shape.ArrowEndLocation, io.Vec3)
	OptionalFunc(io, &shape.ArrowHeadLength, io.Float32)
	OptionalFunc(io, &shape.ArrowHeadRadius, io.Float32)
	OptionalFunc(io, &shape.Segments, io.Uint8)
}

const (
	PrimitiveShapeLine uint8 = iota
	PrimitiveShapeBox
	PrimitiveShapeSphere
	PrimitiveShapeCircle
	PrimitiveShapeText
	PrimitiveShapeArrow
)

// PrimitiveShape defines a single shape to be rendered on the client. Each shape has a unique NetworkID and a set of
// optional parameters depending on its type.
//
// Added: v1.26.20
type PrimitiveShape struct {
	// NetworkID is the network ID of the shape.
	//
	// Added: v1.26.20
	NetworkID uint64
	// DimensionID is the optional dimension ID where the shape is rendered.
	//
	// Added: v1.26.20
	DimensionID Optional[int32]
	// AttachedToEntityID is the optional unique ID of the entity the shape is attached to.
	//
	// Added: v1.26.20
	AttachedToEntityID Optional[int64]
	// Type is the type of the shape.
	// If not set, the set shape will be cleared.
	//
	// Added: v1.26.20
	Type Optional[uint8]
	// Location is the location of the shape.
	//
	// Added: v1.26.20
	Location Optional[mgl32.Vec3]
	// Scale is the scale of the shape.
	//
	// Added: v1.26.20
	Scale Optional[float32]
	// Rotation is the rotation of the shape.
	//
	// Added: v1.26.20
	Rotation Optional[mgl32.Vec3]
	// TotalTimeLeft is the remaining lifetime of the shape, in seconds.
	//
	// Added: v1.26.20
	TotalTimeLeft Optional[float32]
	// MaxRenderDistance is the maximum distance from the camera at which the shape should render.
	//
	// Added: v1.26.20
	MaxRenderDistance Optional[float32]
	// Colour is the ARGB colour of the shape.
	//
	// Added: v1.26.20
	Colour Optional[color.RGBA]
	// ExtraShapeData holds data specific to the shape type, such as the text payload for text shapes.
	//
	// Added: v1.26.20
	ExtraShapeData ShapeData
}

// Marshal ...
func (x *PrimitiveShape) Marshal(io IO) {
	io.Varuint64(&x.NetworkID)
	OptionalFunc(io, &x.Type, io.Uint8)
	OptionalFunc(io, &x.Location, io.Vec3)
	OptionalFunc(io, &x.Scale, io.Float32)
	OptionalFunc(io, &x.Rotation, io.Vec3)
	OptionalFunc(io, &x.TotalTimeLeft, io.Float32)
	if io.Protocol() >= Protocol1v26v20 {
		OptionalFunc(io, &x.MaxRenderDistance, io.Float32)
	}
	OptionalFunc(io, &x.Colour, io.ARGB)
	OptionalFunc(io, &x.DimensionID, io.Varint32)
	OptionalFunc(io, &x.AttachedToEntityID, io.Varint64)
	io.ShapeData(&x.ExtraShapeData)
}
