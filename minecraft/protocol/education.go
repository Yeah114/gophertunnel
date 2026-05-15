package protocol

// EducationSharedResourceURI is an education edition feature that is used for transmitting
// education resource settings to clients. It contains a button name and a link URL.
//
// Added: v1.17.30
type EducationSharedResourceURI struct {
	// ButtonName is the button name of the resource URI.
	//
	// Added: v1.17.30
	ButtonName string
	// LinkURI is the link URI for the resource URI.
	//
	// Added: v1.17.30
	LinkURI string
}

// Marshal reads/writes an EducationSharedResourceURI to an IO.
func (x *EducationSharedResourceURI) Marshal(r IO) {
	r.String(&x.ButtonName)
	r.String(&x.LinkURI)
}

// EducationExternalLinkSettings holds the external link metadata shown in Education Edition UI.
//
// Added: v1.17.30
type EducationExternalLinkSettings struct {
	// URL is the external link URL.
	//
	// Added: v1.17.30
	URL string
	// DisplayName is the display name in game.
	//
	// Added: v1.17.30
	DisplayName string
}

// Marshal encodes/decodes an EducationExternalLinkSettings.
func (x *EducationExternalLinkSettings) Marshal(r IO) {
	r.String(&x.URL)
	r.String(&x.DisplayName)
}
