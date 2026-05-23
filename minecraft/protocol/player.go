package protocol

import (
	"github.com/google/uuid"
	"image/color"
)

const (
	PlayerActionStartBreak = iota
	PlayerActionAbortBreak
	PlayerActionStopBreak
	PlayerActionGetUpdatedBlock
	PlayerActionDropItem
	PlayerActionStartSleeping
	PlayerActionStopSleeping
	PlayerActionRespawn
	PlayerActionJump
	PlayerActionStartSprint
	PlayerActionStopSprint
	PlayerActionStartSneak
	PlayerActionStopSneak
	PlayerActionCreativePlayerDestroyBlock
	PlayerActionDimensionChangeDone
	PlayerActionStartGlide
	PlayerActionStopGlide
	PlayerActionBuildDenied
	PlayerActionCrackBreak
	PlayerActionChangeSkin
	PlayerActionSetEnchantmentSeed
	PlayerActionStartSwimming
	PlayerActionStopSwimming
	PlayerActionStartSpinAttack
	PlayerActionStopSpinAttack
	PlayerActionStartBuildingBlock
	PlayerActionPredictDestroyBlock
	PlayerActionContinueDestroyBlock
	PlayerActionStartItemUseOn
	PlayerActionStopItemUseOn
	PlayerActionHandledTeleport
	PlayerActionMissedSwing
	PlayerActionStartCrawling
	PlayerActionStopCrawling
	PlayerActionStartFlying
	PlayerActionStopFlying
	PlayerActionReceivedServerData
	_
	PlayerActionStartUsingItem
)

const (
	PlayerMovementModeClient = iota
	PlayerMovementModeServer
	PlayerMovementModeServerWithRewind
)

// PlayerListEntry is an entry found in the PlayerList packet. It represents a single player using the UUID
// found in the entry, and contains several properties such as the skin.
//
// Added: v1.16
type PlayerListEntry struct {
	// UUID is the UUID of the player as sent in the Login packet when the client joined the server. It must
	// match this UUID exactly for the correct XBOX Live icon to show up in the list.
	//
	// Added: v1.16
	UUID uuid.UUID
	// EntityUniqueID is the unique entity ID of the player. This ID typically stays consistent during the
	// lifetime of a world, but servers often send the runtime ID for this.
	//
	// Added: v1.16
	EntityUniqueID int64
	// Username is the username that is shown in the player list of the player that obtains a PlayerList
	// packet with this entry. It does not have to be the same as the actual username of the player.
	//
	// Added: v1.16
	Username string
	// XUID is the XBOX Live user ID of the player, which will remain consistent as long as the player is
	// logged in with the XBOX Live account.
	//
	// Added: v1.16
	XUID string
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	//
	// Added: v1.16
	PlatformChatID string
	// BuildPlatform is the platform of the player as sent by that player in the Login packet.
	//
	// Added: v1.16
	BuildPlatform int32
	// Skin is the skin of the player that should be added to the player list. Once sent here, it will not
	// have to be sent again.
	//
	// Added: v1.16
	Skin Skin
	// Teacher is a Minecraft: Education Edition field. It specifies if the player to be added to the player
	// list is a teacher.
	//
	// Added: v1.16
	Teacher bool
	// Host specifies if the player that is added to the player list is the host of the game.
	//
	// Added: v1.16
	Host bool
	// SubClient specifies if the player that is added to the player list is a sub-client of another player.
	//
	// Added: v1.20.60
	SubClient bool
	// PlayerColour is the colour of the player that is shown in UI elements, currently only used for the
	// player locator bar.
	//
	// Added: v1.21.80
	PlayerColour color.RGBA
}

// Marshal encodes/decodes a PlayerListEntry.
func (x *PlayerListEntry) Marshal(r IO) {
	r.UUID(&x.UUID)
	r.Varint64(&x.EntityUniqueID)
	r.String(&x.Username)
	if r.Protocol() >= Protocol1v2v13 && r.Protocol() <= Protocol1v6v0 {
		thirdPartyName := ""
		thirdPartyID := int32(0)
		r.String(&thirdPartyName)
		r.Varint32(&thirdPartyID)
	}
	if r.Protocol() < Protocol1v13v0 {
		Single(r, &x.Skin)
		if r.Protocol() < Protocol1v2v13 {
			legacySkinData := []byte(nil)
			r.ByteSlice(&legacySkinData)
		}
	}
	r.String(&x.XUID)
	if r.Protocol() >= Protocol1v2v13 {
		r.String(&x.PlatformChatID)
		if r.Protocol() >= Protocol1v13v0 {
			r.Int32(&x.BuildPlatform)
			Single(r, &x.Skin)
			r.Bool(&x.Teacher)
			r.Bool(&x.Host)
			if r.Protocol() >= Protocol1v20v60 {
				r.Bool(&x.SubClient)
			}
			if r.Protocol() >= Protocol1v21v80 {
				r.ARGB(&x.PlayerColour)
			}
		}
	}
}

// PlayerListRemoveEntry encodes/decodes a PlayerListEntry for removal from the list.
func PlayerListRemoveEntry(r IO, x *PlayerListEntry) {
	r.UUID(&x.UUID)
}

// PlayerMovementSettings represents the different server authoritative movement settings. These control how
// the client will provide input to the server.
//
// Added: v1.16
type PlayerMovementSettings struct {
	// MovementType specifies the way the server handles player movement. Available options are
	// PlayerMovementModeClient, PlayerMovementModeServer and PlayerMovementModeServerWithRewind.
	//
	// Added: v1.16.0, Changed: v1.16.100, encoded as a movement mode instead of a legacy bool.
	MovementType int32
	// RewindHistorySize is the amount of history to keep at maximum if MovementType is
	// PlayerMovementModeServerWithRewind.
	//
	// Added: v1.16.210
	RewindHistorySize int32
	// ServerAuthoritativeBlockBreaking specifies if block breaking should be sent through
	// packet.PlayerAuthInput or not.
	//
	// Added: v1.16.210
	ServerAuthoritativeBlockBreaking bool
}

// PlayerMoveSettings reads/writes PlayerMovementSettings x to/from IO r.
func PlayerMoveSettings(r IO, x *PlayerMovementSettings) {
	if r.Protocol() >= Protocol1v26v20v26 {
		movementType := uint32(x.MovementType)
		r.Varuint32(&movementType)
		x.MovementType = int32(movementType)
		r.Varint32(&x.RewindHistorySize)
		r.Bool(&x.ServerAuthoritativeBlockBreaking)
	} else if r.Protocol() >= Protocol1v16v210 {
		r.Varint32(&x.RewindHistorySize)
		r.Bool(&x.ServerAuthoritativeBlockBreaking)
	} else if r.Protocol() >= Protocol1v16v100 {
		r.Varint32(&x.MovementType)
	} else {
		serverAuthoritativeMovement := x.MovementType != PlayerMovementModeClient
		r.Bool(&serverAuthoritativeMovement)
		if serverAuthoritativeMovement {
			x.MovementType = PlayerMovementModeServer
		} else {
			x.MovementType = PlayerMovementModeClient
		}
	}
}

// PlayerBlockAction represents a block-related action initiated by the player.
//
// Added: v1.16
type PlayerBlockAction struct {
	// Action is the action to be performed, and is one of the constants listed above.
	//
	// Added: v1.16
	Action int32
	// BlockPos is the position of the block that was interacted with.
	//
	// Added: v1.16
	BlockPos BlockPos
	// Face is the face of the block that was interacted with.
	//
	// Added: v1.16
	Face int32
}

// Marshal encodes/decodes a PlayerBlockAction.
func (x *PlayerBlockAction) Marshal(r IO) {
	r.Varint32(&x.Action)
	switch x.Action {
	case PlayerActionStartBreak, PlayerActionAbortBreak, PlayerActionCrackBreak, PlayerActionPredictDestroyBlock, PlayerActionContinueDestroyBlock:
		r.BlockPos(&x.BlockPos)
		r.Varint32(&x.Face)
	}
}

// PlayerArmourDamageEntry represents an entry for a single piece of armour that should be damaged.
//
// Added: v1.16
type PlayerArmourDamageEntry struct {
	// ArmourSlot is the index of the armour slot to damage.
	//
	// Added: v1.16
	ArmourSlot int32
	// Damage is the amount of damage to apply to the armour in the specified slot.
	//
	// Added: v1.16
	Damage int16
}

// Marshal encodes/decodes a PlayerArmourDamageEntry.
func (x *PlayerArmourDamageEntry) Marshal(r IO) {
	r.Varint32(&x.ArmourSlot)
	r.Int16(&x.Damage)
}
