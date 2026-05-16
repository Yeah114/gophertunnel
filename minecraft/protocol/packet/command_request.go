package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CommandRequest is sent by the client to request the execution of a server-side command. Although some
// servers support sending commands using the Text packet, this packet is guaranteed to have the correct
// result.
//
// Added: v1.11.1
type CommandRequest struct {
	// CommandLine is the raw entered command line. The client does no parsing of the command line by itself
	// (unlike it did in the early stages), but lets the server do that.
	//
	// Added: v1.11.1
	CommandLine string
	// CommandOrigin is the data specifying the origin of the command. In other words, the source that the
	// command was from, such as the player itself or a websocket server.
	//
	// Added: v1.11.1
	CommandOrigin protocol.CommandOrigin
	// Internal specifies if the command request internal. Setting it to false seems to work and the usage of
	// this field is not known.
	//
	// Added: v1.11.1
	Internal bool
	// Version is the version of the command that is being executed. This field currently has no purpose or functionality.
	//
	// Added: v1.19.60
	// Changed: v1.21.130.28, encoded as a string instead of a varint32.
	Version string
}

// ID ...
func (*CommandRequest) ID() uint32 {
	return IDCommandRequest
}

func (pk *CommandRequest) Marshal(io protocol.IO) {
	io.String(&pk.CommandLine)
	protocol.CommandOriginData(io, &pk.CommandOrigin)
	io.Bool(&pk.Internal)
	if io.Protocol() >= protocol.Protocol1v21v130v28 {
		io.String(&pk.Version)
	} else if io.Protocol() >= protocol.Protocol1v19v60 {
		version := int32(0)
		io.Varint32(&version)
	}
}
