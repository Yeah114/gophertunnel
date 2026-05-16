package protocol

import (
	"github.com/google/uuid"
)

const (
	CommandPermissionLevelAny = iota
	CommandPermissionLevelGameDirectors
	CommandPermissionLevelAdmin
	CommandPermissionLevelHost
	CommandPermissionLevelOwner
	CommandPermissionLevelInternal
)

// Command holds the data that a command requires to be shown to a player client-side. The command is shown in
// the /help command and auto-completed using this data.
//
// Added: v1.12
type Command struct {
	// Name is the name of the command. The command may be executed using this name, and will be shown in the
	// /help list with it. It currently seems that the client crashes if the Name contains uppercase letters.
	//
	// Added: v1.12
	Name string
	// Description is the description of the command. It is shown in the /help list and when starting to write
	// a command.
	//
	// Added: v1.12
	Description string
	// Flags is a combination of flags not currently known. Leaving the Flags field empty appears to work.
	//
	// Added: v1.12
	Flags uint16
	// PermissionLevel is the command permission level that the player required to execute this command. The
	// field no longer seems to serve a purpose, as the client does not handle the execution of commands
	// anymore: The permissions should be checked server-side.
	//
	// Added: v1.12
	PermissionLevel byte
	// AliasesOffset is the offset to a CommandEnum that holds the values that
	// should be used as aliases for this command.
	//
	// Added: v1.12
	AliasesOffset uint32
	// ChainedSubcommandOffsets is a slice of offsets that all point to a different ChainedSubcommand from the
	// ChainedSubcommands slice in the AvailableCommands packet.
	//
	// Added: v1.20.10
	ChainedSubcommandOffsets []uint32
	// Overloads is a list of command overloads that specify the ways in which a command may be executed. The
	// overloads may be completely different.
	//
	// Added: v1.12
	Overloads []CommandOverload
}

func (c *Command) Marshal(r IO) {
	CommandContext{ChainedSubcommands: true}.Marshal(r, c)
}

// CommandContext holds context required for encoding commands.
//
// Added: v1.20.10
type CommandContext struct {
	// ChainedSubcommands specifies if chained subcommand offsets and overload chaining are encoded.
	//
	// Added: v1.20.10
	ChainedSubcommands bool
	// ShortFlags specifies if Flags is encoded as a little-endian uint16 instead of the legacy single byte.
	//
	// Added: v1.17.10
	ShortFlags bool
}

// Marshal encodes/decodes a Command.
func (ctx CommandContext) Marshal(r IO, c *Command) {
	permissionStr := commandPermissionToString(c.PermissionLevel)
	r.String(&c.Name)
	r.String(&c.Description)
	if ctx.ShortFlags {
		r.Uint16(&c.Flags)
	} else {
		flags := byte(c.Flags)
		r.Uint8(&flags)
		c.Flags = uint16(flags)
	}
	r.String(&permissionStr)
	commandPermissionFromString(r, &c.PermissionLevel, permissionStr)
	r.Uint32(&c.AliasesOffset)
	if ctx.ChainedSubcommands {
		FuncSlice(r, &c.ChainedSubcommandOffsets, r.Uint32)
	}
	FuncIOSlice(r, &c.Overloads, CommandOverloadContext{Chaining: ctx.ChainedSubcommands}.Marshal)
}

// CommandOverload represents an overload of a command. This overload can be compared to function overloading
// in languages such as java. It represents a single usage of the command. A command may have multiple
// different overloads, which are handled differently.
//
// Added: v1.12
type CommandOverload struct {
	// Chaining determines if the parameters use chained subcommands or not.
	//
	// Added: v1.12
	Chaining bool
	// Parameters is a list of command parameters that are part of the overload. These parameters specify the
	// usage of the command when this overload is applied.
	//
	// Added: v1.12
	Parameters []CommandParameter
}

func (c *CommandOverload) Marshal(r IO) {
	CommandOverloadContext{Chaining: true}.Marshal(r, c)

}

// CommandOverloadContext holds context required for encoding command overloads.
//
// Added: v1.20.10
type CommandOverloadContext struct {
	// Chaining specifies if the overload chaining flag is encoded.
	//
	// Added: v1.20.10
	Chaining bool
}

// Marshal encodes/decodes a CommandOverload.
func (ctx CommandOverloadContext) Marshal(r IO, c *CommandOverload) {
	if ctx.Chaining {
		r.Bool(&c.Chaining)
	}
	Slice(r, &c.Parameters)
}

const (
	CommandArgValid    = 0x100000
	CommandArgEnum     = 0x200000
	CommandArgSuffixed = 0x1000000
	CommandArgSoftEnum = 0x4000000

	CommandArgTypeInt             = 1
	CommandArgTypeFloat           = 3
	CommandArgTypeValue           = 4
	CommandArgTypeWildcardInt     = 5
	CommandArgTypeOperator        = 6
	CommandArgTypeCompareOperator = 7
	CommandArgTypeTarget          = 8
	CommandArgTypeWildcardTarget  = 10
	CommandArgTypeFilepath        = 17
	CommandArgTypeIntegerRange    = 23
	CommandArgTypeEquipmentSlots  = 47
	CommandArgTypeString          = 56
	CommandArgTypeBlockPosition   = 64
	CommandArgTypePosition        = 65
	CommandArgTypeMessage         = 67
	CommandArgTypeRawText         = 70
	CommandArgTypeJSON            = 74
	CommandArgTypeBlockStates     = 83
	CommandArgTypeCommand         = 87
)

const (
	// ParamOptionCollapseEnum specifies if the enum (only if the Type is actually an enum type. If not,
	// setting this to true has no effect) should be collapsed. This means that the options of the enum are
	// never shown in the actual usage of the command, but only as auto-completion, like it automatically does
	// with enums that have a big amount of options. To illustrate, it can make
	// <false|true|yes|no> <$Name: bool>.
	ParamOptionCollapseEnum = iota + 1
	ParamOptionHasSemanticConstraint
	_
	ParamOptionAsChainedCommand
)

// CommandParameter represents a single parameter of a command overload, which accepts a certain type of input
// values. It has a name and a type which show up client-side when a player is entering the command.
//
// Added: v1.12
type CommandParameter struct {
	// Name is the name of the command parameter. It shows up in the usage like <$Name: $Type>, with the
	// exception of enum types, which show up simply as a list of options if the list is short enough and
	// Options is set to false.
	//
	// Added: v1.12
	Name string
	// Type is a rather odd combination of type(flag)s that result in a certain parameter type to show up
	// client-side. It is a combination of the flags above. The basic types must be combined with the
	// ArgumentTypeFlagBasic flag (and integers with a suffix ArgumentTypeFlagSuffixed), whereas enums are
	// combined with the ArgumentTypeFlagEnum flag.
	//
	// Added: v1.12
	Type uint32
	// Optional specifies if the command parameter is optional to enter. Note that no non-optional parameter
	// should ever be present in a command overload after an optional parameter. When optional, the parameter
	// shows up like so: [$Name: $Type], whereas when mandatory, it shows up like so: <$Name: $Type>.
	//
	// Added: v1.12
	Optional bool
	// Options holds a combinations of options that additionally apply to the command parameter. The list of
	// options can be found above.
	//
	// Added: v1.12
	Options byte
}

func (c *CommandParameter) Marshal(r IO) {
	r.String(&c.Name)
	r.Uint32(&c.Type)
	r.Bool(&c.Optional)
	r.Uint8(&c.Options)
}

// CommandEnum represents an enum in a command usage. The enum typically has a type and a set of options that
// are valid. A value that is not one of the options results in a failure during execution.
//
// Added: v1.12
type CommandEnum struct {
	// Type is the type of the command enum. The type will show up in the command usage as the type of the
	// argument if it has a certain amount of arguments, or when Options is set to true in the
	// command holding the enum.
	//
	// Added: v1.12
	Type string
	// ValueIndices holds a list of indices that point to the EnumValues slice in the
	// AvailableCommandsPacket. These represent the options of the enum.
	//
	// Added: v1.12
	ValueIndices []uint32
}

// CommandEnumContext holds context required for encoding command enums.
//
// Added: v1.19.80
type CommandEnumContext struct {
	// EnumValues is the shared string table that command enum value indices refer to during encoding.
	//
	// Added: v1.19.80
	EnumValues []string
}

// Marshal encodes/decodes a CommandEnum.
func (ctx CommandEnumContext) Marshal(r IO, x *CommandEnum) {
	r.String(&x.Type)
	FuncSlice(r, &x.ValueIndices, r.Uint32)
}

// ChainedSubcommand represents a subcommand that can have chained commands, such as /execute which allows you to run
// another command as another entity or at a different position etc.
//
// Added: v1.20.10
type ChainedSubcommand struct {
	// Name is the name of the chained subcommand and shows up in the list as a regular subcommand enum.
	//
	// Added: v1.20.10
	Name string
	// Values contains the index and parameter type of the chained subcommand.
	//
	// Added: v1.20.10
	Values []ChainedSubcommandValue
}

func (x *ChainedSubcommand) Marshal(r IO) {
	r.String(&x.Name)
	Slice(r, &x.Values)
}

// ChainedSubcommandValue represents the value for a chained subcommand argument.
//
// Added: v1.20.10
type ChainedSubcommandValue struct {
	// Index is the index of the argument in the ChainedSubcommandValues slice from the AvailableCommands packet. This is
	// then used to set the type specified by the Value field below.
	//
	// Added: v1.20.10
	Index uint32
	// Value is a combination of the flags above and specified the type of argument. Unlike regular parameter types,
	// this should NOT contain any of the special flags (valid, enum, suffixed or soft enum) but only the basic types.
	//
	// Added: v1.20.10
	Value uint32
}

func (x *ChainedSubcommandValue) Marshal(r IO) {
	r.Varuint32(&x.Index)
	r.Varuint32(&x.Value)
}

// DynamicEnum is an enum variant that can have its options changed during runtime,
// without sending a new AvailableCommands packet.
//
// Added: v1.19.80
type DynamicEnum struct {
	// Type is the type of the command enum. The type will show up in the command usage as the type of the
	// argument if it has a certain amount of arguments, or when Options is set to true in the
	// command holding the enum.
	//
	// Added: v1.19.80
	Type string
	// Values is a slice of possible options for the enum.
	//
	// Added: v1.19.80
	Values []string
}

// Marshal encodes/decodes a DynamicEnum.
func (x *DynamicEnum) Marshal(r IO) {
	r.String(&x.Type)
	FuncSlice(r, &x.Values, r.String)
}

const (
	CommandEnumConstraintCheatsEnabled = iota
	CommandEnumConstraintOperatorPermissions
	CommandEnumConstraintHostPermissions
	_
)

// CommandEnumConstraint is sent in the AvailableCommands packet to limit what values of an enum may be used
// taking in account things such as whether cheats are enabled.
//
// Added: v1.16.220
type CommandEnumConstraint struct {
	// EnumValueIndex points to an enum value in the AvailableCommands packet that this
	// constraint should apply to.
	//
	// Added: v1.16.220
	EnumValueIndex uint32
	// EnumIndex points to an enum in the AvailableCommands packet to which this
	// constraint should apply to.
	//
	// Added: v1.16.220
	EnumIndex uint32
	// Constraints holds a slice of constraints as present in the constants above.
	//
	// Added: v1.16.220
	Constraints []byte
}

// Marshal encodes/decodes a CommandEnumConstraint.
func (c *CommandEnumConstraint) Marshal(r IO) {
	r.Uint32(&c.EnumValueIndex)
	r.Uint32(&c.EnumIndex)
	r.ByteSlice(&c.Constraints)
}

const (
	CommandOriginPlayer = iota
	CommandOriginBlock
	CommandOriginMinecartBlock
	CommandOriginDevConsole
	CommandOriginTest
	CommandOriginAutomationPlayer
	CommandOriginClientAutomation
	CommandOriginDedicatedServer
	CommandOriginEntity
	CommandOriginVirtual
	CommandOriginGameArgument
	CommandOriginEntityServer
	CommandOriginPrecompiled
	CommandOriginGameDirectorEntityServer
	CommandOriginScript
	CommandOriginExecutor
)

// CommandOrigin holds data that identifies the origin of the requesting of a command. It holds several
// fields that may be used to get specific information.
// When sent in a CommandRequest packet, the same CommandOrigin should be sent in a CommandOutput packet.
//
// Added: v1.12
type CommandOrigin struct {
	// Origin is one of the values above that specifies the origin of the command. The origin may change,
	// depending on what part of the client actually called the command. The command may be issued by a
	// websocket server, for example.
	//
	// Added: v1.12
	Origin uint32
	// UUID is a unique identifier for every instantiation of a command.
	//
	// Added: v1.12
	UUID uuid.UUID
	// RequestID is an ID that identifies the request of the client. The server should send a CommandOrigin
	// with the same request ID to ensure it can be matched with the request by the caller of the command.
	// This is especially important for websocket servers and it seems that this field is only non-empty for
	// these websocket servers.
	//
	// Added: v1.12
	RequestID string
	// PlayerUniqueID is an ID that identifies the player, the same as the one found in the AdventureSettings
	// packet. Filling it out with 0 seems to work.
	// PlayerUniqueID is only written if Origin is CommandOriginDevConsole or CommandOriginTest.
	//
	// Added: v1.12
	PlayerUniqueID int64
}

// CommandOriginData reads/writes a CommandOrigin x using IO r.
func CommandOriginData(r IO, x *CommandOrigin) {
	if r.Protocol() >= Protocol1v21v130v28 {
		originStr := commandOriginToString(x.Origin)
		r.String(&originStr)
		commandOriginFromString(r, &x.Origin, originStr)
	} else {
		r.Varuint32(&x.Origin)
	}
	r.UUID(&x.UUID)
	r.String(&x.RequestID)
	if r.Protocol() >= Protocol1v21v130v28 {
		r.Int64(&x.PlayerUniqueID)
	} else if x.Origin == CommandOriginDevConsole || x.Origin == CommandOriginTest {
		r.Varint64(&x.PlayerUniqueID)
	}
}

// CommandOutputMessage represents a message sent by a command that holds the output of one of the commands
// executed.
//
// Added: v1.12
type CommandOutputMessage struct {
	// Success indicates if the output message was one of a successful command execution. If set to true, the
	// output message is by default coloured white, whereas if set to false, the message is by default
	// coloured red.
	//
	// Added: v1.12
	Success bool
	// Message is the message that is sent to the client in the chat window. It may either be simply a
	// message or a translated built-in string like 'commands.tp.success.coordinates', combined with specific
	// parameters below.
	//
	// Added: v1.12
	Message string
	// Parameters is a list of parameters that serve to supply the message sent with additional information,
	// such as the position that a player was teleported to or the effect that was applied to an entity.
	// These parameters only apply for the Minecraft built-in command output.
	//
	// Added: v1.12
	Parameters []string
}

// Marshal encodes/decodes a CommandOutputMessage.
func (x *CommandOutputMessage) Marshal(r IO) {
	if r.Protocol() >= Protocol1v21v130v28 {
		r.String(&x.Message)
		r.Bool(&x.Success)
	} else {
		r.Bool(&x.Success)
		r.String(&x.Message)
	}
	FuncSlice(r, &x.Parameters, r.String)
}

func commandOriginToString(x uint32) string {
	switch x {
	case CommandOriginPlayer:
		return "player"
	case CommandOriginBlock:
		return "commandblock"
	case CommandOriginMinecartBlock:
		return "minecartcommandblock"
	case CommandOriginDevConsole:
		return "devconsole"
	case CommandOriginTest:
		return "test"
	case CommandOriginAutomationPlayer:
		return "automationplayer"
	case CommandOriginClientAutomation:
		return "clientautomation"
	case CommandOriginDedicatedServer:
		return "dedicatedserver"
	case CommandOriginEntity:
		return "entity"
	case CommandOriginVirtual:
		return "virtual"
	case CommandOriginGameArgument:
		return "gameargument"
	case CommandOriginEntityServer:
		return "entityserver"
	case CommandOriginPrecompiled:
		return "precompiled"
	case CommandOriginGameDirectorEntityServer:
		return "gamedirectorentityserver"
	case CommandOriginScript:
		return "scripting"
	case CommandOriginExecutor:
		return "executecontext"
	default:
		return "unknown"
	}
}

func commandOriginFromString(r IO, x *uint32, s string) {
	switch s {
	case "player":
		*x = CommandOriginPlayer
	case "commandblock":
		*x = CommandOriginBlock
	case "minecartcommandblock":
		*x = CommandOriginMinecartBlock
	case "devconsole":
		*x = CommandOriginDevConsole
	case "test":
		*x = CommandOriginTest
	case "automationplayer":
		*x = CommandOriginAutomationPlayer
	case "clientautomation":
		*x = CommandOriginClientAutomation
	case "dedicatedserver":
		*x = CommandOriginDedicatedServer
	case "entity":
		*x = CommandOriginEntity
	case "virtual":
		*x = CommandOriginVirtual
	case "gameargument":
		*x = CommandOriginGameArgument
	case "entityserver":
		*x = CommandOriginEntityServer
	case "precompiled":
		*x = CommandOriginPrecompiled
	case "gamedirectorentityserver":
		*x = CommandOriginGameDirectorEntityServer
	case "scripting":
		*x = CommandOriginScript
	case "executecontext":
		*x = CommandOriginExecutor
	default:
		r.InvalidValue(s, "origin", "unknown origin")
	}
}

func commandPermissionToString(x byte) string {
	switch x {
	case CommandPermissionLevelAny:
		return "any"
	case CommandPermissionLevelGameDirectors:
		return "gamedirectors"
	case CommandPermissionLevelAdmin:
		return "admin"
	case CommandPermissionLevelHost:
		return "host"
	case CommandPermissionLevelOwner:
		return "owner"
	case CommandPermissionLevelInternal:
		return "internal"
	default:
		return "unknown"
	}
}

func commandPermissionFromString(r IO, x *byte, s string) {
	switch s {
	case "any":
		*x = CommandPermissionLevelAny
	case "gamedirectors":
		*x = CommandPermissionLevelGameDirectors
	case "admin":
		*x = CommandPermissionLevelAdmin
	case "host":
		*x = CommandPermissionLevelHost
	case "owner":
		*x = CommandPermissionLevelOwner
	case "internal":
		*x = CommandPermissionLevelInternal
	default:
		r.InvalidValue(s, "permissionLevel", "unknown permission level")
	}
}
