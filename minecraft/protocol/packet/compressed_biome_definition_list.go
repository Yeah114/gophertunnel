package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// CompressedBiomeDefinitionList is sent by the server to send compressed biome definitions to the client.
//
// Added: v1.21.0
type CompressedBiomeDefinitionList struct {
	// Payload is the compressed biome definition payload, including the compression indicator and dictionary
	// data. It is kept raw so connections can forward it without having to rebuild Mojang's dictionary format.
	//
	// Added: v1.21.0
	Payload []byte
}

// ID ...
func (*CompressedBiomeDefinitionList) ID() uint32 {
	return IDCompressedBiomeDefinitionList
}

func (pk *CompressedBiomeDefinitionList) Marshal(io protocol.IO) {
	io.ByteSlice(&pk.Payload)
}
