package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	TextTypeRaw = iota
	TextTypeChat
	TextTypeTranslation
	TextTypePopup
	TextTypeJukeboxPopup
	TextTypeTip
	TextTypeSystem
	TextTypeWhisper
	TextTypeAnnouncement
	TextTypeObject
	TextTypeObjectWhisper
	TextTypeObjectAnnouncement
)

const (
	TextCategoryMessageOnly = iota
	TextCategoryAuthoredMessage
	TextCategoryMessageWithParameters
)

// Text is sent by the client to the server to send chat messages, and by the server to the client to forward
// or send messages, which may be chat, popups, tips etc.
//
// Added: v1.11.1
type Text struct {
	// TextType is the type of the text sent. When a client sends this to the server, it should always be
	// TextTypeChat. If the server sends it, it may be one of the other text types above.
	//
	// Added: v1.11.1
	TextType byte
	// NeedsTranslation specifies if any of the messages need to be translated. It seems that where % is found
	// in translatable text types, these are translated regardless of this bool. Translatable text types
	// include TextTypeTranslation, TextTypeTip, TextTypePopup and TextTypeJukeboxPopup.
	//
	// Added: v1.11.1
	NeedsTranslation bool
	// SourceName is the name of the source of the messages. This source is displayed in text types such as
	// the TextTypeChat and TextTypeWhisper, where typically the username is shown.
	//
	// Added: v1.11.1
	SourceName string
	// Message is the message of the packet. This field is set for each TextType and is the main component of
	// the packet.
	//
	// Added: v1.11.1
	Message string
	// Parameters is a list of parameters that should be filled into the message. These parameters are only
	// written if the type of the packet is TextTypeTranslation, TextTypeTip, TextTypePopup or TextTypeJukeboxPopup.
	//
	// Added: v1.11.1
	Parameters []string
	// XUID is the XBOX Live user ID of the player that sent the message. It is only set for packets of
	// TextTypeChat. When sent to a player, the player will only be shown the chat message if a player with
	// this XUID is present in the player list and not muted, or if the XUID is empty.
	//
	// Added: v1.12.0
	XUID string
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	//
	// Added: v1.2.13, Changed: v1.19.50
	PlatformChatID string
	// FilteredMessage is a filtered version of Message with all the profanity removed. The client will use
	// this over Message if this field is not empty and they have the "Filter Profanity" setting enabled.
	//
	// Added: v1.21.0
	FilteredMessage protocol.Optional[string]
}

// ID ...
func (*Text) ID() uint32 {
	return IDText
}

func (pk *Text) Marshal(io protocol.IO) {
	if io.Protocol() >= protocol.Protocol1v26v0 {
		pk.Marshal1v26v0(io)
		return
	}
	if io.Protocol() >= protocol.Protocol1v21v130 {
		pk.Marshal1v21v130(io)
		return
	}
	if io.Protocol() >= protocol.Protocol1v21v50 {
		pk.Marshal1v21v50(io)
		return
	}
	// panic("Marshal: Unsupport Text packet")
}

func (pk *Text) Marshal1v26v0(io protocol.IO) {
	io.Bool(&pk.NeedsTranslation)

	var categoryType uint8
	switch pk.TextType {
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObjectWhisper, TextTypeObjectAnnouncement, TextTypeObject:
		categoryType = TextCategoryMessageOnly
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		categoryType = TextCategoryAuthoredMessage
	default:
		categoryType = TextCategoryMessageWithParameters
	}
	io.Uint8(&categoryType)
	io.Uint8(&pk.TextType)
	switch pk.TextType {
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		io.String(&pk.SourceName)
		io.String(&pk.Message)
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObject, TextTypeObjectWhisper, TextTypeObjectAnnouncement:
		io.String(&pk.Message)
	case TextTypeTranslation, TextTypePopup, TextTypeJukeboxPopup:
		io.String(&pk.Message)
		protocol.FuncSlice(io, &pk.Parameters, io.String)
	}

	if len(pk.Message) == 0 {
		io.InvalidValue(pk.Message, "message", "string cannot be empty")
	}
	io.String(&pk.XUID)
	io.String(&pk.PlatformChatID)
	protocol.OptionalFunc(io, &pk.FilteredMessage, io.String)
}

func (pk *Text) Marshal1v21v130(io protocol.IO) {
	io.Bool(&pk.NeedsTranslation)

	var categoryType uint8
	switch pk.TextType {
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObjectWhisper, TextTypeObjectAnnouncement, TextTypeObject:
		categoryType = TextCategoryMessageOnly
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		categoryType = TextCategoryAuthoredMessage
	default:
		categoryType = TextCategoryMessageWithParameters
	}
	io.Uint8(&categoryType)

	switch categoryType {
	case TextCategoryMessageOnly:
		io.StringConst("raw")
		io.StringConst("tip")
		io.StringConst("systemMessage")
		io.StringConst("textObjectWhisper")
		io.StringConst("textObjectAnnouncement")
		io.StringConst("textObject")
	case TextCategoryAuthoredMessage:
		io.StringConst("chat")
		io.StringConst("whisper")
		io.StringConst("announcement")
	default:
		io.StringConst("translate")
		io.StringConst("popup")
		io.StringConst("jukeboxPopup")
	}

	io.Uint8(&pk.TextType)
	switch pk.TextType {
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		io.String(&pk.SourceName)
		io.String(&pk.Message)
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObject, TextTypeObjectWhisper, TextTypeObjectAnnouncement:
		io.String(&pk.Message)
	case TextTypeTranslation, TextTypePopup, TextTypeJukeboxPopup:
		io.String(&pk.Message)
		protocol.FuncSlice(io, &pk.Parameters, io.String)
	}

	if len(pk.Message) == 0 {
		io.InvalidValue(pk.Message, "message", "string cannot be empty")
	}
	io.String(&pk.XUID)
	io.String(&pk.PlatformChatID)
	protocol.OptionalFunc(io, &pk.FilteredMessage, io.String)
}

func (pk *Text) Marshal1v21v50(io protocol.IO) {
	io.Uint8(&pk.TextType)
	io.Bool(&pk.NeedsTranslation)
	switch pk.TextType {
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		io.String(&pk.SourceName)
		io.String(&pk.Message)
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObject, TextTypeObjectWhisper, TextTypeObjectAnnouncement:
		io.String(&pk.Message)
	case TextTypeTranslation, TextTypePopup, TextTypeJukeboxPopup:
		io.String(&pk.Message)
		protocol.FuncSlice(io, &pk.Parameters, io.String)
	}
	io.String(&pk.XUID)
	io.String(&pk.PlatformChatID)
	filteredMessage, _ := pk.FilteredMessage.Value()
	io.String(&filteredMessage)
	pk.FilteredMessage = protocol.Option(filteredMessage)
}
