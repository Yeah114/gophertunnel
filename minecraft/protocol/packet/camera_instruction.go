package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// CameraInstruction gives a custom camera specific instructions to operate.
//
// Added: v1.20.30
type CameraInstruction struct {
	// Set is a camera instruction that sets the camera to a specified preset.
	//
	// Added: v1.20.30
	Set protocol.Optional[protocol.CameraInstructionSet]
	// Clear can be set to true to clear all the current camera instructions.
	//
	// Added: v1.20.30
	Clear protocol.Optional[bool]
	// Fade is a camera instruction that fades the screen to a specified colour.
	//
	// Added: v1.20.30
	Fade protocol.Optional[protocol.CameraInstructionFade]
	// Target is a camera instruction that targets a specific entity.
	//
	// Added: v1.21.20
	Target protocol.Optional[protocol.CameraInstructionTarget]
	// RemoveTarget can be set to true to remove the current aim assist target.
	//
	// Added: v1.21.20
	RemoveTarget protocol.Optional[bool]
	// FieldOfView is a camera instruction that updates the field of view for the camera.
	//
	// Added: v1.21.100
	FieldOfView protocol.Optional[protocol.CameraInstructionFieldOfView]
	// Spline is a camera instruction that creates a spline path for the camera to follow.
	//
	// Added: v1.21.111
	Spline protocol.Optional[protocol.CameraSplineInstruction]
	// AttachToEntity is the entity ID to attach the camera to.
	//
	// Added: v1.21.111
	AttachToEntity protocol.Optional[int64]
	// DetachFromEntity can be set to true to detach the camera from the current entity.
	//
	// Added: v1.21.111
	DetachFromEntity protocol.Optional[bool]
}

// ID ...
func (*CameraInstruction) ID() uint32 {
	return IDCameraInstruction
}

func (pk *CameraInstruction) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.Set)
	protocol.OptionalFunc(io, &pk.Clear, io.Bool)
	protocol.OptionalMarshaler(io, &pk.Fade)
	if io.Protocol() >= protocol.Protocol1v21v20 {
		protocol.OptionalMarshaler(io, &pk.Target)
		protocol.OptionalFunc(io, &pk.RemoveTarget, io.Bool)
	}
	if io.Protocol() >= protocol.Protocol1v21v100 {
		protocol.OptionalMarshaler(io, &pk.FieldOfView)
	}
	if io.Protocol() >= protocol.Protocol1v21v120 {
		protocol.OptionalMarshaler(io, &pk.Spline)
		protocol.OptionalFunc(io, &pk.AttachToEntity, io.Int64)
		protocol.OptionalFunc(io, &pk.DetachFromEntity, io.Bool)
	}
}
