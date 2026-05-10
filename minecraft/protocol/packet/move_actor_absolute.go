package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	MoveFlagOnGround = 1 << iota
	MoveFlagTeleport
	// MoveFlagForceCompletion forces the move to complete. This flag was added in v1.26.20.26.
	MoveFlagForceCompletion
)

// MoveActorAbsolute is sent by the server to move an entity to an absolute position. It is typically used
// for movements where high accuracy isn't needed, such as for long range teleporting.
type MoveActorAbsolute struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// Flags is a combination of flags that specify details of the movement. It is a combination of the flags
	// above.
	Flags byte
	// ForceCompletion forces the move to complete. This field was added in v1.26.20.26.
	ForceCompletion bool
	// Position is the position to spawn the entity on. If the entity is on a distance that the player cannot
	// see it, the entity will still show up if the player moves closer.
	Position mgl32.Vec3
	// Rotation is a Vec3 holding the X, Y and Z rotation of the entity after the movement. This is a Vec3 for
	// the reason that projectiles like arrows don't have yaw/pitch, but do have roll.
	Rotation mgl32.Vec3
}

// ID ...
func (*MoveActorAbsolute) ID() uint32 {
	return IDMoveActorAbsolute
}

func (pk *MoveActorAbsolute) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	flags := pk.Flags
	if io.Protocol() >= protocol.Protocol1v26v20v26 && pk.ForceCompletion {
		flags |= MoveFlagForceCompletion
	}
	io.Uint8(&flags)
	pk.Flags = flags
	pk.ForceCompletion = flags&MoveFlagForceCompletion != 0
	io.Vec3(&pk.Position)
	io.ByteFloat(&pk.Rotation[0])
	io.ByteFloat(&pk.Rotation[1])
	io.ByteFloat(&pk.Rotation[2])
}
