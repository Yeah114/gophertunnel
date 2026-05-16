package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddPlayer is sent by the server to the client to make a player entity show up client-side. It is one of the
// few entities that cannot be sent using the AddActor packet.
//
// Added: v1.16
type AddPlayer struct {
	// UUID is the UUID of the player. It is the same UUID that the client sent in the Login packet at the
	// start of the session. A player with this UUID must exist in the player list (built up using the
	// PlayerList packet), for it to show up in-game.
	//
	// Added: v1.16
	UUID uuid.UUID
	// Username is the name of the player. This username is the username that will be set as the initial
	// name tag of the player.
	//
	// Added: v1.16
	Username string
	// EntityUniqueID is the unique ID of the player. The unique ID is a value that remains consistent across
	// different sessions of the same world, but most servers simply fill the runtime ID of the player out for
	// this field.
	//
	// Added: v1.16
	// Removed: v1.19.10
	EntityUniqueID int64
	// EntityRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	//
	// Added: v1.16
	EntityRuntimeID uint64
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	//
	// Added: v1.2.13
	PlatformChatID string
	// Position is the position to spawn the player on. If the player is on a distance that the viewer cannot
	// see it, the player will still show up if the viewer moves closer.
	//
	// Added: v1.16
	Position mgl32.Vec3
	// Velocity is the initial velocity the player spawns with. This velocity will initiate client side
	// movement of the player.
	//
	// Added: v1.16
	Velocity mgl32.Vec3
	// Pitch is the vertical rotation of the player. Facing straight forward yields a pitch of 0. Pitch is
	// measured in degrees.
	//
	// Added: v1.16
	Pitch float32
	// Yaw is the horizontal rotation of the player. Yaw is also measured in degrees.
	//
	// Added: v1.16
	Yaw float32
	// HeadYaw is the same as Yaw, except that it applies specifically to the head of the player. A different
	// value for HeadYaw than Yaw means that the player will have its head turned.
	//
	// Added: v1.16
	HeadYaw float32
	// HeldItem is the item that the player is holding. The item is shown to the viewer as soon as the player
	// itself shows up. Needless to say that this field is rather pointless, as additional packets still must
	// be sent for armour to show up.
	//
	// Added: v1.16
	HeldItem protocol.ItemInstance
	// GameType is the game type of the player. If set to GameTypeSpectator, the player will not be shown to viewers.
	//
	// Added: v1.18.30
	GameType int32
	// EntityMetadata is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the player looks. Flags include ones such as 'on fire' and 'sprinting'.
	// The metadata values are indexed by their property key.
	//
	// Added: v1.16
	EntityMetadata protocol.EntityMetadata
	// EntityProperties is a list of properties that the entity inhibits. These properties define and alter specific
	// attributes of the entity.
	//
	// Added: v1.19.40
	EntityProperties protocol.EntityProperties
	// LegacyAbilityFlags is a set of flags that specify certain properties of the player, such as whether it can fly
	// and/or move through blocks.
	//
	// Added: v1.16
	// Removed: v1.19.10
	LegacyAbilityFlags uint32
	// LegacyCommandPermissionLevel specifies what commands a player is allowed to execute before ability data was moved
	// into its own structure.
	//
	// Added: v1.16
	// Removed: v1.19.10
	LegacyCommandPermissionLevel uint32
	// LegacyActionPermissions is a set of flags that specify actions that the player is allowed to undertake.
	//
	// Added: v1.16
	// Removed: v1.19.10
	LegacyActionPermissions uint32
	// LegacyPermissionLevel is the permission level of the player as it shows up in the player list built up using the
	// PlayerList packet.
	//
	// Added: v1.16
	// Removed: v1.19.10
	LegacyPermissionLevel uint32
	// LegacyCustomStoredPermissions is an old permission field that was replaced by AbilityData.
	//
	// Added: v1.16
	// Removed: v1.19.10
	LegacyCustomStoredPermissions uint32
	// AbilityData represents various data about the abilities of a player, such as ability layers or permissions.
	//
	// Added: v1.19.10
	AbilityData protocol.AbilityData
	// EntityLinks is a list of entity links that are currently active on the player. These links alter the
	// way the player shows up when first spawned in terms of it shown as riding an entity. Setting these
	// links is important for new viewers to see the player is riding another entity.
	//
	// Added: v1.16
	EntityLinks []protocol.EntityLink
	// DeviceID is the device ID set in one of the files found in the storage of the device of the player. It
	// may be changed freely, so it should not be relied on for anything.
	//
	// Added: v1.16
	DeviceID string
	// BuildPlatform is the build platform/device OS of the player that is about to be added, as it sent in
	// the Login packet when joining.
	//
	// Added: v1.16
	BuildPlatform int32
}

// ID ...
func (*AddPlayer) ID() uint32 {
	return IDAddPlayer
}

func (pk *AddPlayer) Marshal(io protocol.IO) {
	io.UUID(&pk.UUID)
	io.String(&pk.Username)
	if io.Protocol() < protocol.Protocol1v19v10 {
		io.Varint64(&pk.EntityUniqueID)
	}
	io.Varuint64(&pk.EntityRuntimeID)
	if io.Protocol() >= protocol.Protocol1v2v13 {
		io.String(&pk.PlatformChatID)
	}
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Float32(&pk.HeadYaw)
	io.ItemInstance(&pk.HeldItem)
	if io.Protocol() >= protocol.Protocol1v18v30 {
		io.Varint32(&pk.GameType)
	}
	io.EntityMetadata(&pk.EntityMetadata)
	if io.Protocol() >= protocol.Protocol1v19v40 {
		protocol.Single(io, &pk.EntityProperties)
	}
	if io.Protocol() >= protocol.Protocol1v19v10 {
		protocol.Single(io, &pk.AbilityData)
	} else {
		io.Varuint32(&pk.LegacyAbilityFlags)
		io.Varuint32(&pk.LegacyCommandPermissionLevel)
		io.Varuint32(&pk.LegacyActionPermissions)
		io.Varuint32(&pk.LegacyPermissionLevel)
		io.Varuint32(&pk.LegacyCustomStoredPermissions)
		io.Int64(&pk.EntityUniqueID)
	}
	protocol.Slice(io, &pk.EntityLinks)
	io.String(&pk.DeviceID)
	io.Int32(&pk.BuildPlatform)
}
