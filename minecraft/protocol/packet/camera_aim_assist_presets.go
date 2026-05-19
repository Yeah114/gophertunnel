package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	CameraAunAssistPresetOperationSet = iota
	CameraAunAssistPresetOperationAddToExisting
)

// CameraAimAssistPresets is sent by the server to the client to provide a list of categories and presets
// that can be used when sending a CameraAimAssist packet or a CameraInstruction including aim assist.
//
// Added: v1.21.40
type CameraAimAssistPresets struct {
	// Categories is a list of categories which can be referenced by one of the Presets.
	//
	// Added: v1.21.40
	Categories []protocol.CameraAimAssistCategory
	// Presets is a list of presets which define a base for how aim assist should behave
	//
	// Added: v1.21.40
	Presets []protocol.CameraAimAssistPreset
	// Operation is the operation to perform with the presets. It is one of the constants above.
	//
	// Added: v1.21.40
	Operation byte
}

// ID ...
func (*CameraAimAssistPresets) ID() uint32 {
	return IDCameraAimAssistPresets
}

func (pk *CameraAimAssistPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Categories)
	protocol.Slice(io, &pk.Presets)
	io.Uint8(&pk.Operation)
}
