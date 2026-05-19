package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ClientBoundAttributeLayerSync is sent by the server to synchronise attribute layers with the client.
//
// Added: v1.26.10
type ClientBoundAttributeLayerSync struct {
	// PayloadType is the type of attribute layer payload. It is one of the protocol.AttributeLayerPayloadType constants.
	//
	// Added: v1.26.10
	PayloadType uint32
	// Layers is set if PayloadType is AttributeLayerPayloadTypeUpdateLayers.
	//
	// Added: v1.26.10
	Layers []protocol.AttributeLayerData
	// LayerName is the attribute layer name, used for UpdateSettings, UpdateEnvironment and RemoveEnvironment.
	//
	// Added: v1.26.10
	LayerName string
	// DimensionID is the dimension ID, used for UpdateSettings, UpdateEnvironment and RemoveEnvironment.
	//
	// Added: v1.26.10
	DimensionID int32
	// Settings is set if PayloadType is AttributeLayerPayloadTypeUpdateSettings.
	//
	// Added: v1.26.10
	Settings protocol.AttributeLayerSettings
	// EnvironmentAttributes is set if PayloadType is AttributeLayerPayloadTypeUpdateEnvironment.
	//
	// Added: v1.26.10
	EnvironmentAttributes []protocol.EnvironmentAttributeData
	// RemoveAttributeNames is set if PayloadType is AttributeLayerPayloadTypeRemoveEnvironment.
	//
	// Added: v1.26.10
	RemoveAttributeNames []string
}

// ID ...
func (*ClientBoundAttributeLayerSync) ID() uint32 {
	return IDClientBoundAttributeLayerSync
}

func (pk *ClientBoundAttributeLayerSync) Marshal(io protocol.IO) {
	io.Varuint32(&pk.PayloadType)
	switch pk.PayloadType {
	case protocol.AttributeLayerPayloadTypeUpdateLayers:
		protocol.Slice(io, &pk.Layers)
	case protocol.AttributeLayerPayloadTypeUpdateSettings:
		io.String(&pk.LayerName)
		io.Varint32(&pk.DimensionID)
		protocol.Single(io, &pk.Settings)
	case protocol.AttributeLayerPayloadTypeUpdateEnvironment:
		io.String(&pk.LayerName)
		io.Varint32(&pk.DimensionID)
		protocol.Slice(io, &pk.EnvironmentAttributes)
	case protocol.AttributeLayerPayloadTypeRemoveEnvironment:
		io.String(&pk.LayerName)
		io.Varint32(&pk.DimensionID)
		protocol.FuncSlice(io, &pk.RemoveAttributeNames, io.String)
	default:
		io.UnknownEnumOption(pk.PayloadType, "attribute layer payload type")
	}
}
