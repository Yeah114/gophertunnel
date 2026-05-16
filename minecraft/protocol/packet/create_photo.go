package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// CreatePhoto is a packet that allows players to export photos from their portfolios into items in their inventory.
// This packet only works on the Education Edition version of Minecraft.
//
// Added: v1.17.30
type CreatePhoto struct {
	// EntityUniqueID is the unique ID of the entity.
	//
	// Added: v1.17.30
	EntityUniqueID int64
	// PhotoName is the name of the photo.
	//
	// Added: v1.17.30
	PhotoName string
	// ItemName is the name of the photo as an item.
	//
	// Added: v1.17.30
	ItemName string
}

// ID ...
func (*CreatePhoto) ID() uint32 {
	return IDCreatePhoto
}

func (pk *CreatePhoto) Marshal(io protocol.IO) {
	io.Int64(&pk.EntityUniqueID)
	io.String(&pk.PhotoName)
	io.String(&pk.ItemName)
}
