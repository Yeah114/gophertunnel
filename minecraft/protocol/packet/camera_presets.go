package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CameraPresets gives the client a list of custom camera presets.
//
// Added: v1.20.30
type CameraPresets struct {
	// Presets is a list of camera presets that can be used by other cameras. The order of this list is important because
	// the index of presets is used as a pointer in the CameraInstruction packet.
	//
	// Added: v1.20.30
	Presets []protocol.CameraPreset
}

// ID ...
func (*CameraPresets) ID() uint32 {
	return IDCameraPresets
}

func (pk *CameraPresets) Marshal(io protocol.IO) {
	if io.Protocol() < protocol.Protocol1v20v30v24 {
		cameraPresetsNBT(io, &pk.Presets)
		return
	}
	protocol.Slice(io, &pk.Presets)
}

func cameraPresetsNBT(io protocol.IO, presets *[]protocol.CameraPreset) {
	data := map[string]any{
		"presets": cameraPresetsToNBT(*presets),
	}
	io.NBT(&data, nbt.NetworkLittleEndian)
	*presets = cameraPresetsFromNBT(data)
}

func cameraPresetsToNBT(presets []protocol.CameraPreset) []any {
	list := make([]any, 0, len(presets))
	for _, preset := range presets {
		presetData := map[string]any{
			"identifier":   preset.Name,
			"inherit_from": preset.Parent,
		}
		if posX, ok := preset.PosX.Value(); ok {
			presetData["pos_x"] = posX
		}
		if posY, ok := preset.PosY.Value(); ok {
			presetData["pos_y"] = posY
		}
		if posZ, ok := preset.PosZ.Value(); ok {
			presetData["pos_z"] = posZ
		}
		if rotX, ok := preset.RotX.Value(); ok {
			presetData["rot_x"] = rotX
		}
		if rotY, ok := preset.RotY.Value(); ok {
			presetData["rot_y"] = rotY
		}
		list = append(list, presetData)
	}
	return list
}

func cameraPresetsFromNBT(data map[string]any) []protocol.CameraPreset {
	list, ok := data["presets"].([]any)
	if !ok {
		return nil
	}
	presets := make([]protocol.CameraPreset, 0, len(list))
	for _, value := range list {
		presetData, ok := value.(map[string]any)
		if !ok {
			continue
		}
		presets = append(presets, cameraPresetFromNBT(presetData))
	}
	return presets
}

func cameraPresetFromNBT(data map[string]any) protocol.CameraPreset {
	var preset protocol.CameraPreset
	if name, ok := data["identifier"].(string); ok {
		preset.Name = name
	}
	if parent, ok := data["inherit_from"].(string); ok {
		preset.Parent = parent
	}
	if posX, ok := data["pos_x"].(float32); ok {
		preset.PosX = protocol.Option(posX)
	}
	if posY, ok := data["pos_y"].(float32); ok {
		preset.PosY = protocol.Option(posY)
	}
	if posZ, ok := data["pos_z"].(float32); ok {
		preset.PosZ = protocol.Option(posZ)
	}
	if rotX, ok := data["rot_x"].(float32); ok {
		preset.RotX = protocol.Option(rotX)
	}
	if rotY, ok := data["rot_y"].(float32); ok {
		preset.RotY = protocol.Option(rotY)
	}
	return preset
}
