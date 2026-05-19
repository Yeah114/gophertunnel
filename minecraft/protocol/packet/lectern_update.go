package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// LecternUpdate is sent by the client to update the server on which page was opened in a book on a lectern,
// or if the book should be removed from it.
//
// Added: v1.11.1
type LecternUpdate struct {
	// Page is the page number in the book that was opened by the player on the lectern.
	//
	// Added: v1.11.1
	Page byte
	// PageCount is the number of pages that the book opened in the lectern has.
	//
	// Added: v1.11.1
	PageCount byte
	// Position is the position of the lectern that was updated. If no lectern is at the block position,
	// the packet should be ignored.
	//
	// Added: v1.11.1
	Position protocol.BlockPos
}

// ID ...
func (*LecternUpdate) ID() uint32 {
	return IDLecternUpdate
}

func (pk *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&pk.Page)
	io.Uint8(&pk.PageCount)
	io.UBlockPos(&pk.Position)
}
