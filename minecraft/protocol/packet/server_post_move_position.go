package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/go-gl/mathgl/mgl32"
)

// ServerPostMovePosition is sent by the server to update the client's post-move position.
//
// Added: v1.20.50
type ServerPostMovePosition struct {
	// Position is the position sent after movement has been processed server-side.
	//
	// Added: v1.20.50
	Position mgl32.Vec3
}

// ID ...
func (*ServerPostMovePosition) ID() uint32 {
	return IDServerPostMovePosition
}

func (pk *ServerPostMovePosition) Marshal(io protocol.IO) {
	io.Vec3(&pk.Position)
}
