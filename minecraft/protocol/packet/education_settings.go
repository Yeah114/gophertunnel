package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// EducationSettings is a packet sent by the server to update Minecraft: Education Edition related settings.
// It is unused by the normal base game.
//
// Added: v1.13
type EducationSettings struct {
	// CodeBuilderDefaultURI is the default URI that the code builder is ran on. Using this, a Code Builder program can
	// make code directly affect the server.
	//
	// Added: v1.13
	CodeBuilderDefaultURI string
	// CodeBuilderTitle is the title of the code builder shown when connected to the CodeBuilderDefaultURI.
	//
	// Added: v1.16
	CodeBuilderTitle string
	// CanResizeCodeBuilder specifies if clients connected to the world should be able to resize the code
	// builder when it is opened.
	//
	// Added: v1.16
	CanResizeCodeBuilder bool
	// DisableLegacyTitleBar ...
	//
	// Added: v1.17.30
	DisableLegacyTitleBar bool
	// PostProcessFilter ...
	//
	// Added: v1.17.30
	PostProcessFilter string
	// ScreenshotBorderPath ...
	//
	// Added: v1.17.30
	ScreenshotBorderPath string
	// CanModifyBlocks ...
	//
	// Added: v1.18.30
	// Changed: v1.18.30, renamed from AgentCapabilities and encoded as an optional bool.
	CanModifyBlocks protocol.Optional[bool]
	// OverrideURI ...
	//
	// Added: v1.17.0
	// Changed: v1.18.30, encoded as an optional string instead of a plain string.
	OverrideURI protocol.Optional[string]
	// HasQuiz specifies if the world has a quiz connected to it.
	//
	// Added: v1.13
	HasQuiz bool
	// ExternalLinkSettings ...
	//
	// Added: v1.17.30
	// Changed: v1.18.30, encoded as an optional EducationExternalLinkSettings value instead of a pointer value.
	ExternalLinkSettings protocol.Optional[protocol.EducationExternalLinkSettings]
}

// ID ...
func (*EducationSettings) ID() uint32 {
	return IDEducationSettings
}

func (pk *EducationSettings) Marshal(io protocol.IO) {
	io.String(&pk.CodeBuilderDefaultURI)
	io.String(&pk.CodeBuilderTitle)
	io.Bool(&pk.CanResizeCodeBuilder)
	io.Bool(&pk.DisableLegacyTitleBar)
	io.String(&pk.PostProcessFilter)
	io.String(&pk.ScreenshotBorderPath)
	protocol.OptionalFunc(io, &pk.CanModifyBlocks, io.Bool)
	protocol.OptionalFunc(io, &pk.OverrideURI, io.String)
	io.Bool(&pk.HasQuiz)
	protocol.OptionalMarshaler(io, &pk.ExternalLinkSettings)
}
