package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CameraAimAssistActionSet = iota
	CameraAimAssistActionClear
)

// CameraAimAssist is sent by the server to the client to set up aim assist for the client's camera.
//
// Added: v1.21.30
type CameraAimAssist struct {
	// Preset is the ID of the preset that has previously been defined in the CameraAimAssistPresets packet.
	//
	// Added: v1.21.50
	Preset string
	// Angle is the maximum angle around the playes's cursor that the aim assist should check for a target,
	// if TargetMode is set to protocol.AimAssistTargetModeAngle.
	//
	// Added: v1.21.40
	Angle mgl32.Vec2
	// Distance is the maximum distance from the player's cursor should check for a target, if TargetMode is
	// set to protocol.AimAssistTargetModeDistance.
	//
	// Added: v1.21.30
	Distance float32
	// TargetMode is the mode that the camera should use for detecting targets. This is currently one of
	// protocol.AimAssistTargetModeAngle or protocol.AimAssistTargetModeDistance.
	//
	// Added: v1.21.30
	TargetMode byte
	// Action is the action that should be performed with the aim assist. This is one of the constants above.
	//
	// Added: v1.21.30
	Action byte
	// ShowDebugRender specifies if debug render should be shown.
	//
	// Added: v1.21.100
	ShowDebugRender bool
}

// ID ...
func (*CameraAimAssist) ID() uint32 {
	return IDCameraAimAssist
}

func (pk *CameraAimAssist) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v21v50 {
		io.String(&pk.Preset)
	}
	io.Vec2(&pk.Angle)
	io.Float32(&pk.Distance)
	io.Uint8(&pk.TargetMode)
	io.Uint8(&pk.Action)
	if io.Protocol() >= protocol.Protocol1v21v100 {
		io.Bool(&pk.ShowDebugRender)
	}
}
