package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	EmoteFlagServerSide = 1 << iota
	EmoteFlagMuteChat
)

// Emote is sent by both the server and the client. When the client sends an emote, it sends this packet to
// the server, after which the server will broadcast the packet to other players online.
//
// Added: v1.13
type Emote struct {
	// EntityRuntimeID is the entity that sent the emote. When a player sends this packet, it has this field
	// set as its own entity runtime ID.
	//
	// Added: v1.13
	EntityRuntimeID uint64
	// EmoteLength is the number of ticks that the emote lasts for.
	//
	// Added: v1.21.30
	EmoteLength uint32
	// EmoteID is the ID of the emote to send.
	//
	// Added: v1.13
	EmoteID string
	// XUID is the Xbox User ID of the player that sent the emote. It is only set when the emote is used by a player that
	// is authenticated with Xbox Live.
	//
	// Added: v1.20.0
	XUID string
	// PlatformID is an identifier only set for particular platforms when using an emote (presumably only for Nintendo
	// Switch). It is otherwise an empty string, and is used to decide which players are able to emote with each other.
	//
	// Added: v1.20.0
	PlatformID string
	// Flags is a combination of flags that change the way the Emote packet operates. When the server sends
	// this packet to other players, EmoteFlagServerSide must be present.
	//
	// Added: v1.13
	Flags byte
}

// ID ...
func (*Emote) ID() uint32 {
	return IDEmote
}

func (pk *Emote) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.String(&pk.EmoteID)
	if io.Protocol() >= protocol.Protocol1v21v30 {
		io.Varuint32(&pk.EmoteLength)
	}
	if io.Protocol() >= protocol.Protocol1v20v0v23 {
		io.String(&pk.XUID)
		io.String(&pk.PlatformID)
	}
	io.Uint8(&pk.Flags)
}
