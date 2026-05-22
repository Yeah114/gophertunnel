package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/go-gl/mathgl/mgl32"
)

const PlayerAuthInputBitsetSize = 65

const (
	InputFlagAscend = iota
	InputFlagDescend
	InputFlagNorthJump
	InputFlagJumpDown
	InputFlagSprintDown
	InputFlagChangeHeight
	InputFlagJumping
	InputFlagAutoJumpingInWater
	InputFlagSneaking
	InputFlagSneakDown
	InputFlagUp
	InputFlagDown
	InputFlagLeft
	InputFlagRight
	InputFlagUpLeft
	InputFlagUpRight
	InputFlagWantUp
	InputFlagWantDown
	InputFlagWantDownSlow
	InputFlagWantUpSlow
	InputFlagSprinting
	InputFlagAscendBlock
	InputFlagDescendBlock
	InputFlagSneakToggleDown
	InputFlagPersistSneak
	InputFlagStartSprinting
	InputFlagStopSprinting
	InputFlagStartSneaking
	InputFlagStopSneaking
	InputFlagStartSwimming
	InputFlagStopSwimming
	InputFlagStartJumping
	InputFlagStartGliding
	InputFlagStopGliding
	InputFlagPerformItemInteraction
	InputFlagPerformBlockActions
	InputFlagPerformItemStackRequest
	InputFlagHandledTeleport
	InputFlagEmoting
	InputFlagMissedSwing
	InputFlagStartCrawling
	InputFlagStopCrawling
	InputFlagStartFlying
	InputFlagStopFlying
	InputFlagClientAckServerData
	InputFlagClientPredictedVehicle
	InputFlagPaddlingLeft
	InputFlagPaddlingRight
	InputFlagBlockBreakingDelayEnabled
	InputFlagHorizontalCollision
	InputFlagVerticalCollision
	InputFlagDownLeft
	InputFlagDownRight
	InputFlagStartUsingItem
	InputFlagCameraRelativeMovementEnabled
	InputFlagRotControlledByMoveDirection
	InputFlagStartSpinAttack
	InputFlagStopSpinAttack
	InputFlagIsHotbarTouchOnly
	InputFlagJumpReleasedRaw
	InputFlagJumpPressedRaw
	InputFlagJumpCurrentRaw
	InputFlagSneakReleasedRaw
	InputFlagSneakPressedRaw
	InputFlagSneakCurrentRaw
)

const (
	InputModeMouse = iota + 1
	InputModeTouch
	InputModeGamePad
)

const (
	PlayModeNormal = iota
	PlayModeTeaser
	PlayModeScreen
	PlayModeViewer
	PlayModeReality
	_
	_
	_
	PlayModeExitLevel
	_
	PlayModeNumModes
)

const (
	InteractionModelTouch = iota
	InteractionModelCrosshair
	InteractionModelClassic
)

// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to synchronise
// the player input with the position server-side.
// The client sends this packet when the ServerAuthoritativeMovementMode field in the StartGame packet is set
// to true, instead of the MovePlayer packet. The client will send this packet once every tick.
//
// Added: v1.13
type PlayerAuthInput struct {
	// Pitch and Yaw hold the rotation that the player reports it has.
	Pitch, Yaw float32
	// Position holds the position that the player reports it has.
	//
	// Added: v1.13
	Position mgl32.Vec3
	// MoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of X/Z
	// values which are created using the WASD/controller stick state.
	//
	// Added: v1.13
	MoveVector mgl32.Vec2
	// HeadYaw is the horizontal rotation of the head that the player reports it has.
	//
	// Added: v1.13
	HeadYaw float32
	// InputData is a combination of bit flags that together specify the way the player moved last tick. It
	// is a combination of the flags above.
	//
	// Added: v1.13
	// Changed: v1.21.50, represented using protocol.Bitset instead of a raw integer bitfield.
	InputData protocol.Bitset
	// InputMode specifies the way that the client inputs data to the screen. It is one of the constants that
	// may be found above.
	//
	// Added: v1.13
	InputMode uint32
	// PlayMode specifies the way that the player is playing. The values it holds, which are rather random,
	// may be found above.
	//
	// Added: v1.13
	PlayMode uint32
	// InteractionModel is a constant representing the interaction model the player is using. It is one of the
	// constants that may be found above.
	//
	// Added: v1.20.80
	InteractionModel uint32
	// InteractPitch and interactYaw is the rotation the player is looking that they intend to use for
	// interactions. This is only different to Pitch and Yaw in cases such as VR or when custom cameras
	// being used.
	//
	// Added: v1.21.40
	InteractPitch, InteractYaw float32
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	//
	// Added: v1.16.100
	Tick uint64
	// Delta was the delta between the old and the new position. There isn't any practical use for this field
	// as it can be calculated by the server itself.
	//
	// Added: v1.16.100
	Delta mgl32.Vec3
	// ItemInteractionData is the transaction data if the InputData includes an item interaction.
	//
	// Added: v1.16.210
	ItemInteractionData protocol.UseItemTransactionData
	// ItemStackRequest is sent by the client to change an item in their inventory.
	//
	// Added: v1.16.210
	ItemStackRequest protocol.ItemStackRequest
	// BlockActions is a slice of block actions that the client has interacted with.
	//
	// Added: v1.16.210
	BlockActions []protocol.PlayerBlockAction
	// VehicleRotation is the rotation of the vehicle that the player is in, if any.
	//
	// Added: v1.20.70
	VehicleRotation mgl32.Vec2
	// ClientPredictedVehicle is the unique ID of the vehicle that the client predicts the player to be in.
	//
	// Added: v1.20.60
	ClientPredictedVehicle int64
	// AnalogueMoveVector is a Vec2 that specifies the direction in which the player moved, as a combination
	// of X/Z values which are created using an analogue input.
	//
	// Added: v1.19.70.24
	AnalogueMoveVector mgl32.Vec2
	// CameraOrientation is the vector that represents the camera's forward direction which can be used to
	// transform movement to be camera relative.
	//
	// Added: v1.21.50
	CameraOrientation mgl32.Vec3
	// RawMoveVector is the value of MoveVector before it is affected by input permissions, sneaking/fly
	// speeds and isn't normalised for analogue inputs.
	//
	// Added: v1.21.50
	RawMoveVector mgl32.Vec2
}

// ID ...
func (pk *PlayerAuthInput) ID() uint32 {
	return IDPlayerAuthInput
}

func (pk *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.HeadYaw)
	io.Bitset(&pk.InputData, PlayerAuthInputBitsetSize)
	io.Varuint32(&pk.InputMode)
	io.Varuint32(&pk.PlayMode)
	if io.Protocol() >= protocol.Protocol1v19v0 {
		io.Varuint32(&pk.InteractionModel)
	}
	if io.Protocol() >= protocol.Protocol1v21v40 {
		io.Float32(&pk.InteractPitch)
		io.Float32(&pk.InteractYaw)
	} else if pk.PlayMode == PlayModeReality {
		io.Vec3(&pk.CameraOrientation)
	}
	io.Varuint64(&pk.Tick)
	io.Vec3(&pk.Delta)

	if pk.InputData.Load(InputFlagPerformItemInteraction) {
		io.PlayerInventoryAction(&pk.ItemInteractionData)
	}

	if pk.InputData.Load(InputFlagPerformItemStackRequest) {
		protocol.Single(io, &pk.ItemStackRequest)
	}

	if pk.InputData.Load(InputFlagPerformBlockActions) {
		protocol.SliceVarint32Length(io, &pk.BlockActions)
	}

	if io.Protocol() >= protocol.Protocol1v20v60 && pk.InputData.Load(InputFlagClientPredictedVehicle) {
		if io.Protocol() >= protocol.Protocol1v20v70 {
			io.Vec2(&pk.VehicleRotation)
		}
		io.Varint64(&pk.ClientPredictedVehicle)
	}

	if io.Protocol() >= protocol.Protocol1v19v70v24 {
		io.Vec2(&pk.AnalogueMoveVector)
	}
	if io.Protocol() >= protocol.Protocol1v21v40 {
		io.Vec3(&pk.CameraOrientation)
	}
	if io.Protocol() >= protocol.Protocol1v21v50 {
		io.Vec2(&pk.RawMoveVector)
	}
}
