package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CommandOutputTypeNone = iota
	CommandOutputTypeLastOutput
	CommandOutputTypeSilent
	CommandOutputTypeAllOutput
	CommandOutputTypeDataSet
)

// CommandOutput is sent by the server to the client to send text as output of a command. Most servers do not
// use this packet and instead simply send Text packets, but there is reason to send it.
// If the origin of a CommandRequest packet is not the player itself, but, for example, a websocket server,
// sending a Text packet will not do what is expected: The message should go to the websocket server, not to
// the client's chat. The CommandOutput packet will make sure the messages are relayed to the correct origin
// of the command request.
//
// Added: v1.11.1
type CommandOutput struct {
	// CommandOrigin is the data specifying the origin of the command. In other words, the source that the
	// command request was from, such as the player itself or a websocket server. The client forwards the
	// messages in this packet to the right origin, depending on what is sent here.
	//
	// Added: v1.11.1
	CommandOrigin protocol.CommandOrigin
	// OutputType specifies the type of output that is sent. The OutputType sent by vanilla games appears to
	// be CommandOutputTypeAllOutput, which seems to work.
	//
	// Added: v1.11.1
	// Changed: v1.21.130, encoded as a string output type identifier instead of a uint8 enum value.
	OutputType byte
	// SuccessCount is the amount of times that a command was executed successfully as a result of the command
	// that was requested. For servers, this is usually a rather meaningless fields, but for vanilla, this is
	// applicable for commands created with Functions.
	//
	// Added: v1.11.1
	// Changed: v1.21.130, encoded as a uint32 instead of a varuint32.
	SuccessCount uint32
	// OutputMessages is a list of all output messages that should be sent to the player. Whether they are
	// shown or not, depends on the type of the messages.
	//
	// Added: v1.11.1
	OutputMessages []protocol.CommandOutputMessage
	// DataSet ... TODO: Find out what this is for.
	//
	// Added: v1.11.1
	// Changed: v1.21.130, encoded as an optional string instead of a conditional string for CommandOutputTypeDataSet.
	DataSet protocol.Optional[string]
}

// ID ...
func (*CommandOutput) ID() uint32 {
	return IDCommandOutput
}

func (pk *CommandOutput) Marshal(io protocol.IO) {
	protocol.CommandOriginData(io, &pk.CommandOrigin)
	outputTypeStr := commandOutputTypeToString(pk.OutputType)
	io.String(&outputTypeStr)
	commandOutputTypeFromString(io, &pk.OutputType, outputTypeStr)
	io.Uint32(&pk.SuccessCount)
	protocol.Slice(io, &pk.OutputMessages)
	protocol.OptionalFunc(io, &pk.DataSet, io.String)
}

func commandOutputTypeToString(x byte) string {
	switch x {
	case CommandOutputTypeNone:
		return "none"
	case CommandOutputTypeLastOutput:
		return "lastoutput"
	case CommandOutputTypeSilent:
		return "silent"
	case CommandOutputTypeAllOutput:
		return "alloutput"
	case CommandOutputTypeDataSet:
		return "dataset"
	default:
		return "unknown"
	}
}

func commandOutputTypeFromString(io protocol.IO, x *byte, s string) {
	switch s {
	case "none":
		*x = CommandOutputTypeNone
	case "lastoutput":
		*x = CommandOutputTypeLastOutput
	case "silent":
		*x = CommandOutputTypeSilent
	case "alloutput":
		*x = CommandOutputTypeAllOutput
	case "dataset":
		*x = CommandOutputTypeDataSet
	default:
		io.InvalidValue(s, "outputType", "unknown type")
	}
}
