package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ChangeMobProperty is a packet sent from the server to the client to change one of the properties of a mob client-side.
//
// Added: v1.18.30
type ChangeMobProperty struct {
	// EntityUniqueID is the unique ID of the entity whose property is being changed.
	//
	// Added: v1.18.30
	EntityUniqueID int64
	// Property is the name of the property being updated.
	//
	// Added: v1.18.30
	Property string
	// BoolValue is set if the property value is a bool type. If the type is not a bool, this field is ignored.
	//
	// Added: v1.18.30
	BoolValue bool
	// StringValue is set if the property value is a string type. If the type is not a string, this field is ignored.
	//
	// Added: v1.18.30
	StringValue string
	// IntValue is set if the property value is an int type. If the type is not an int, this field is ignored.
	//
	// Added: v1.18.30
	IntValue int32
	// FloatValue is set if the property value is a float type. If the type is not a float, this field is ignored.
	//
	// Added: v1.18.30
	FloatValue float32
}

// ID ...
func (*ChangeMobProperty) ID() uint32 {
	return IDChangeMobProperty
}

func (pk *ChangeMobProperty) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
	io.String(&pk.Property)
	io.Bool(&pk.BoolValue)
	io.String(&pk.StringValue)
	io.Varint32(&pk.IntValue)
	io.Float32(&pk.FloatValue)
}
