package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PredictionTypePlayer = iota
	PredictionTypeVehicle
)

// CorrectPlayerMovePrediction is sent by the server if and only if StartGame.ServerAuthoritativeMovementMode
// is set to AuthoritativeMovementModeServerWithRewind. The packet is used to correct movement at a specific
// point in time.
//
// Added: v1.16.100
type CorrectPlayerMovePrediction struct {
	// PredictionType is the type of prediction that was corrected. It is one of the constants above.
	//
	// Added: v1.20.70
	PredictionType byte
	// Position is the position that the player is supposed to be at the tick written in the field below.
	// The client will change its current position based on movement after that tick starting from the
	// Position.
	//
	// Added: v1.16.100
	Position mgl32.Vec3
	// Delta is the change in position compared to what the client sent as its position at that specific tick.
	//
	// Added: v1.16.100
	Delta mgl32.Vec3
	// Rotation is the rotation of the player at the tick written in the field below.
	//
	// Added: v1.20.70
	// Changed: v1.21.100, encoded for all prediction types instead of only PredictionTypeVehicle.
	Rotation mgl32.Vec2
	// VehicleAngularVelocity is the angular velocity of the vehicle that the rider is riding.
	//
	// Added: v1.21.20
	// Changed: v1.21.100, encoded for all prediction types as an optional field instead of only PredictionTypeVehicle.
	VehicleAngularVelocity protocol.Optional[float32]
	// OnGround specifies if the player was on the ground at the time of the tick below.
	//
	// Added: v1.16.100
	OnGround bool
	// Tick is the tick of the movement which was corrected by this packet.
	//
	// Added: v1.16.100
	Tick uint64
}

// ID ...
func (*CorrectPlayerMovePrediction) ID() uint32 {
	return IDCorrectPlayerMovePrediction
}

func (pk *CorrectPlayerMovePrediction) Marshal(io protocol.IO) {
	io.Uint8(&pk.PredictionType)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Delta)
	io.Vec2(&pk.Rotation)
	protocol.OptionalFunc(io, &pk.VehicleAngularVelocity, io.Float32)
	io.Bool(&pk.OnGround)
	io.Varuint64(&pk.Tick)
}
