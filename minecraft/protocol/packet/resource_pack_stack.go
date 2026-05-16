package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ResourcePackStack is sent by the server to send the order in which resource packs and behaviour packs
// should be applied (and downloaded) by the client.
//
// Added: v1.11.1
type ResourcePackStack struct {
	// TexturePackRequired specifies if the client must accept the texture packs the server has in order to
	// join the server. If set to true, the client gets the option to either download the resource packs and
	// join, or quit entirely. Behaviour packs never have to be downloaded.
	//
	// Added: v1.11.1
	TexturePackRequired bool
	// BehaviourPacks is a list of behaviour packs that the client needs to download before joining the server.
	// All of these behaviour packs will be applied together, and the order does not necessarily matter.
	//
	// Added: v1.11.1
	// Removed: v1.21.130.28
	BehaviourPacks []protocol.StackResourcePack
	// TexturePacks is a list of texture packs that the client needs to download before joining the server.
	// The order of these texture packs specifies the order that they are applied in on the client side. The
	// first in the list will be applied first.
	//
	// Added: v1.11.1
	TexturePacks []protocol.StackResourcePack
	// Experimental specifies if experimental gameplay should be enabled for the pack stack in legacy protocols.
	//
	// Added: v1.8.0, Removed: v1.13.0
	Experimental bool
	// BaseGameVersion is the vanilla version that the client should set its resource pack stack to.
	//
	// Added: v1.13.0
	BaseGameVersion string
	// Experiments holds a list of experiments that are either enabled or disabled in the world that the
	// player spawns in.
	// It is not clear why experiments are sent both here and in the StartGame packet.
	//
	// Added: v1.16.100
	Experiments []protocol.ExperimentData
	// ExperimentsPreviouslyToggled specifies if any experiments were previously toggled in this world. It is
	// probably used for some kind of metrics.
	//
	// Added: v1.16.100
	ExperimentsPreviouslyToggled bool
	// IncludeEditorPacks specifies if vanilla editor packs should be included in the resource pack stack when
	// connecting to an editor world.
	//
	// Added: v1.20.80
	// Changed: v1.21.130.28, retained after behaviour packs were removed from this packet.
	IncludeEditorPacks bool
}

// ID ...
func (*ResourcePackStack) ID() uint32 {
	return IDResourcePackStack
}

func (pk *ResourcePackStack) Marshal(io protocol.IO) {
	io.Bool(&pk.TexturePackRequired)
	if io.Protocol() < protocol.Protocol1v21v130v28 {
		protocol.Slice(io, &pk.BehaviourPacks)
	}
	protocol.Slice(io, &pk.TexturePacks)
	if io.Protocol() >= protocol.Protocol1v13v0 {
		io.String(&pk.BaseGameVersion)
	} else if io.Protocol() >= protocol.Protocol1v8v0 {
		io.Bool(&pk.Experimental)
	}
	if io.Protocol() >= protocol.Protocol1v16v100 {
		protocol.SliceUint32Length(io, &pk.Experiments)
		io.Bool(&pk.ExperimentsPreviouslyToggled)
	}
	if io.Protocol() >= protocol.Protocol1v20v80 {
		io.Bool(&pk.IncludeEditorPacks)
	}
}
