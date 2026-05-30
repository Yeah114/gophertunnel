package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// UpdateBlockProperties is reserved for the legacy update block properties packet.
type UpdateBlockProperties struct{}

// ID ...
func (*UpdateBlockProperties) ID() uint32 {
	return IDUpdateBlockProperties
}

func (*UpdateBlockProperties) Marshal(protocol.IO) {}
