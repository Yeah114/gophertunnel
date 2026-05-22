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
	// Changed: v1.21.80, flattened from grouped categories to a direct category list.
	Categories []protocol.CameraAimAssistCategory
	// CategoryGroups is a list of named category groups used by legacy clients.
	//
	// Added: v1.21.50
	// Removed: v1.21.80
	CategoryGroups []protocol.CameraAimAssistCategoryGroup
	// Presets is a list of presets which define a base for how aim assist should behave
	//
	// Added: v1.21.40
	Presets []protocol.CameraAimAssistPreset
	// Operation is the operation to perform with the presets. It is one of the constants above.
	//
	// Added: v1.21.60
	Operation byte
}

// ID ...
func (*CameraAimAssistPresets) ID() uint32 {
	return IDCameraAimAssistPresets
}

func (pk *CameraAimAssistPresets) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v21v80 {
		protocol.Slice(io, &pk.Categories)
	} else {
		if len(pk.CategoryGroups) == 0 && len(pk.Categories) != 0 {
			pk.CategoryGroups = []protocol.CameraAimAssistCategoryGroup{{Categories: pk.Categories}}
		}
		protocol.Slice(io, &pk.CategoryGroups)
		if len(pk.Categories) == 0 && len(pk.CategoryGroups) != 0 {
			for _, group := range pk.CategoryGroups {
				pk.Categories = append(pk.Categories, group.Categories...)
			}
		}
	}
	protocol.Slice(io, &pk.Presets)
	if io.Protocol() >= protocol.Protocol1v21v60 {
		io.Uint8(&pk.Operation)
	}
}
