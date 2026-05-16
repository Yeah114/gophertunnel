package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	StructureTemplateRequestExportFromSave = iota + 1
	StructureTemplateRequestExportFromLoad
	StructureTemplateRequestQuerySavedStructure
)

// StructureTemplateDataRequest is sent by the client to request data of a structure.
//
// Added: v1.19.50
type StructureTemplateDataRequest struct {
	// StructureName is the name of the structure that was set in the structure block's UI. This is the name
	// used to export the structure to a file.
	//
	// Added: v1.19.50
	StructureName string
	// Position is the position of the structure block that has its template data requested.
	//
	// Added: v1.19.50
	Position protocol.BlockPos
	// Settings is a struct of settings that should be used for exporting the structure. These settings are
	// identical to the last sent in the StructureBlockUpdate packet by the client.
	//
	// Added: v1.19.50
	Settings protocol.StructureSettings
	// RequestType specifies the type of template data request that the player sent. It is one of the
	// constants found above.
	//
	// Added: v1.19.50
	RequestType byte
}

// ID ...
func (pk *StructureTemplateDataRequest) ID() uint32 {
	return IDStructureTemplateDataRequest
}

func (pk *StructureTemplateDataRequest) Marshal(io protocol.IO) {
	io.String(&pk.StructureName)
	io.UBlockPos(&pk.Position)
	protocol.Single(io, &pk.Settings)
	io.Uint8(&pk.RequestType)
}
