package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/go-gl/mathgl/mgl32"
)

// AddActor is sent by the server to the client to spawn an entity to the player. It is used for every entity
// except other players, for which the AddPlayer packet is used.
//
// Added: v1.16
type AddActor struct {
	// EntityUniqueID is the unique ID of the entity. The unique ID is a value that remains consistent across
	// different sessions of the same world, but most servers simply fill the runtime ID of the entity out for
	// this field.
	//
	// Added: v1.16
	EntityUniqueID int64
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.16
	EntityRuntimeID uint64
	// EntityType is the string entity type of the entity, for example 'minecraft:skeleton'. A list of these
	// entities may be found online.
	//
	// Added: v1.16
	EntityType string
	// LegacyEntityType is the numeric entity type written before string entity identifiers were introduced.
	//
	// Removed: v1.8.0
	LegacyEntityType uint32
	// Position is the position to spawn the entity on. If the entity is on a distance that the player cannot
	// see it, the entity will still show up if the player moves closer.
	//
	// Added: v1.16
	Position mgl32.Vec3
	// Velocity is the initial velocity the entity spawns with. This velocity will initiate client side
	// movement of the entity.
	//
	// Added: v1.16
	Velocity mgl32.Vec3
	// Pitch is the vertical rotation of the entity. Facing straight forward yields a pitch of 0. Pitch is
	// measured in degrees.
	//
	// Added: v1.16
	Pitch float32
	// Yaw is the horizontal rotation of the entity. Yaw is also measured in degrees.
	//
	// Added: v1.16
	Yaw float32
	// HeadYaw is the same as Yaw, except that it applies specifically to the head of the entity. A different value for
	// HeadYaw than Yaw means that the entity will have its head turned.
	//
	// Added: v1.16
	HeadYaw float32
	// BodyYaw is the same as Yaw, except that it applies specifically to the body of the entity. A different value for
	// BodyYaw than HeadYaw means that the entity will have its body turned, although it is unclear what the difference
	// between BodyYaw and Yaw is.
	//
	// Added: v1.16
	BodyYaw float32
	// Attributes is a slice of attributes that the entity has. It includes attributes such as its health,
	// movement speed, etc.
	//
	// Added: v1.16
	Attributes []protocol.AttributeValue
	// EntityMetadata is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the entity looks. Flags include ones such as 'on fire' and 'sprinting'.
	// The metadata values are indexed by their property key.
	//
	// Added: v1.16
	EntityMetadata protocol.EntityMetadata
	// EntityProperties is a list of properties that the entity inhibits. These properties define and alter specific
	// attributes of the entity.
	//
	// Added: v1.19.20
	EntityProperties protocol.EntityProperties
	// EntityLinks is a list of entity links that are currently active on the entity. These links alter the
	// way the entity shows up when first spawned in terms of it shown as riding an entity. Setting these
	// links is important for new viewers to see the entity is riding another entity.
	//
	// Added: v1.16
	EntityLinks []protocol.EntityLink
}

// ID ...
func (*AddActor) ID() uint32 {
	return IDAddActor
}

func (pk *AddActor) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	if io.Protocol() < protocol.Protocol1v8v0 {
		io.Varuint32(&pk.LegacyEntityType)
	} else {
		io.String(&pk.EntityType)
	}
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	if io.Protocol() >= protocol.Protocol1v5v0 {
		io.Float32(&pk.HeadYaw)
		if io.Protocol() >= protocol.Protocol1v19v10 {
			io.Float32(&pk.BodyYaw)
		}
	}
	protocol.Slice(io, &pk.Attributes)
	io.EntityMetadata(&pk.EntityMetadata)
	if io.Protocol() >= protocol.Protocol1v19v40 {
		protocol.Single(io, &pk.EntityProperties)
	}
	protocol.Slice(io, &pk.EntityLinks)
}
