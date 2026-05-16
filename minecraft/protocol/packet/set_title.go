package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	TitleActionClear = iota
	TitleActionReset
	TitleActionSetTitle
	TitleActionSetSubtitle
	TitleActionSetActionBar
	TitleActionSetDurations
	TitleActionTitleTextObject
	TitleActionSubtitleTextObject
	TitleActionActionbarTextObject
)

// SetTitle is sent by the server to make a title, subtitle or action bar shown to a player. It has several
// fields that allow setting the duration of the titles.
//
// Added: v1.12
type SetTitle struct {
	// ActionType is the type of the action that should be executed upon the title of a player. It is one of
	// the constants above and specifies the response of the client to the packet.
	//
	// Added: v1.12
	ActionType int32
	// Text is the text of the title, which has a different meaning depending on the ActionType that the
	// packet has. The text is the text of a title, subtitle or action bar, depending on the type set.
	//
	// Added: v1.12
	Text string
	// FadeInDuration is the duration that the title takes to fade in on the screen of the player. It is
	// measured in 20ths of a second (AKA in ticks).
	//
	// Added: v1.12
	FadeInDuration int32
	// RemainDuration is the duration that the title remains on the screen of the player. It is measured in
	// 20ths of a second (AKA in ticks).
	//
	// Added: v1.12
	RemainDuration int32
	// FadeOutDuration is the duration that the title takes to fade out of the screen of the player. It is
	// measured in 20ths of a second (AKA in ticks).
	//
	// Added: v1.12
	FadeOutDuration int32
	// XUID is the XBOX Live user ID of the player, which will remain consistent as long as the player is
	// logged in with the XBOX Live account. It is empty if the user is not logged into its XBL account.
	//
	// Added: v1.17.10
	XUID string
	// PlatformOnlineID is either a uint64 or an empty string.
	//
	// Added: v1.17.10
	PlatformOnlineID string
	// FilteredMessage is a filtered version of Message with all the profanity removed. The client will use
	// this over Message if this field is not empty and they have the "Filter Profanity" setting enabled.
	//
	// Added: v1.21.20
	// Changed: v1.21.50, used as the profanity-filtered replacement for Text on clients with text filtering enabled.
	FilteredMessage string
}

// ID ...
func (*SetTitle) ID() uint32 {
	return IDSetTitle
}

func (pk *SetTitle) Marshal(io protocol.IO) {
	io.Varint32(&pk.ActionType)
	io.String(&pk.Text)
	io.Varint32(&pk.FadeInDuration)
	io.Varint32(&pk.RemainDuration)
	io.Varint32(&pk.FadeOutDuration)
	io.String(&pk.XUID)
	io.String(&pk.PlatformOnlineID)
	io.String(&pk.FilteredMessage)
}
