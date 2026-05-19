package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	ClientInputLockCamera = 1 << (iota + 1)
	ClientInputLockMovement
	_
	ClientInputLockLateralMovement
	ClientInputLockSneak
	ClientInputLockJump
	ClientInputLockMount
	ClientInputLockDismount
	ClientInputLockMoveForward
	ClientInputLockMoveBackward
	ClientInputLockMoveLeft
	ClientInputLockMoveRight
)

// UpdateClientInputLocks is sent by the server to the client to lock specific player inputs such as camera
// rotation, movement, jumping, sneaking, mounting or individual directional movement.
//
// Added: v1.19.50
type UpdateClientInputLocks struct {
	// Locks is a set of flags that specify which client inputs are disabled, such as whether the player can
	// move, rotate the camera, jump, sneak or mount/dismount entities. It is a combination of the
	// ClientInputLock constants above.
	//
	// Added: v1.19.50
	Locks uint32
	// Position is the server's position of the client at the time the packet was sent.
	//
	// Removed: v1.26.10
	Position mgl32.Vec3
}

// ID ...
func (pk *UpdateClientInputLocks) ID() uint32 {
	return IDUpdateClientInputLocks
}

func (pk *UpdateClientInputLocks) Marshal(io protocol.IO) {
	io.Varuint32(&pk.Locks)
	if io.Protocol() < protocol.Protocol1v26v10 {
		io.Vec3(&pk.Position)
	}
}
