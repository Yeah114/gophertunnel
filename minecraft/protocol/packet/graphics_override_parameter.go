package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// GraphicsOverrideParameter is sent by the server to override graphics parameters.
//
// Added: v1.21.120
type GraphicsOverrideParameter struct {
	// Values is a list of parameter keyframe values.
	//
	// Added: v1.21.120
	Values []protocol.ParameterKeyframeValue
	// FloatValue is an optional single float graphics parameter to be overridden.
	//
	// Added: v1.26.0
	// Changed: v1.26.10, encoded as an optional single float value instead of an always-present float value.
	FloatValue protocol.Optional[float32]
	// Vec3Value is an optional single Vec3 graphics parameter to be overridden.
	//
	// Added: v1.26.0
	// Changed: v1.26.10, encoded as an optional single Vec3 value instead of an always-present Vec3 value.
	Vec3Value protocol.Optional[mgl32.Vec3]
	// BiomeIdentifier is the identifier of the biome for which the parameters apply.
	//
	// Added: v1.21.120
	BiomeIdentifier string
	// ParameterType is the type of parameter being overridden.
	//
	// Added: v1.21.120
	ParameterType uint8
	// Reset indicates whether to reset the parameters.
	//
	// Added: v1.21.120
	Reset bool
}

// ID ...
func (*GraphicsOverrideParameter) ID() uint32 {
	return IDGraphicsOverrideParameter
}

func (pk *GraphicsOverrideParameter) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Values)
	protocol.OptionalFunc(io, &pk.FloatValue, io.Float32)
	protocol.OptionalFunc(io, &pk.Vec3Value, io.Vec3)
	io.String(&pk.BiomeIdentifier)
	io.Uint8(&pk.ParameterType)
	io.Bool(&pk.Reset)
}
