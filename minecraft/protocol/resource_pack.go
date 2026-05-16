package protocol

import "github.com/google/uuid"

const (
	PackSettingTypeFloat = iota
	PackSettingTypeBool
	PackSettingTypeString
)

// TexturePackInfo represents a texture pack's info sent over network. It holds information about the
// texture pack such as its name, description and version.
//
// Added: v1.16.200
type TexturePackInfo struct {
	// UUID is the UUID of the texture pack. Each texture pack downloaded must have a different UUID in
	// order for the client to be able to handle them properly.
	//
	// Added: v1.16.200
	UUID uuid.UUID
	// Version is the version of the texture pack. The client will cache texture packs sent by the server as
	// long as they carry the same version. Sending a texture pack with a different version than previously
	// will force the client to re-download it.
	//
	// Added: v1.16.200
	Version string
	// Size is the total size in bytes that the texture pack occupies. This is the size of the compressed
	// archive (zip) of the texture pack.
	//
	// Added: v1.16.200
	Size uint64
	// ContentKey is the key used to decrypt the behaviour pack if it is encrypted. This is generally the case
	// for marketplace texture packs.
	//
	// Added: v1.16.200
	ContentKey string
	// SubPackName is the selected sub-pack name inside the texture pack, if any.
	//
	// Added: v1.16.200
	SubPackName string
	// ContentIdentity is another UUID for the resource pack, and is generally set for marketplace texture
	// packs. It is also required for client-side validations when the resource pack is encrypted.
	//
	// Added: v1.6.0.5
	ContentIdentity string
	// HasScripts specifies if the texture packs has any scripts in it. A client will only download the
	// behaviour pack if it supports scripts, which, up to 1.11, only includes Windows 10.
	//
	// Added: v1.9.0
	HasScripts bool
	// AddonPack specifies if the texture pack is from an addon.
	//
	// Added: v1.16.200
	AddonPack bool
	// RTXEnabled specifies if the texture pack uses the raytracing technology introduced in 1.16.200.
	//
	// Added: v1.16.200
	RTXEnabled bool
	// DownloadURL is a URL that the client can use to download the pack instead of the server sending it in
	// chunks, which it will continue to do if this field is left empty.
	//
	// Added: v1.16.200
	DownloadURL string
}

// Marshal encodes/decodes a TexturePackInfo.
func (x *TexturePackInfo) Marshal(r IO) {
	if r.Protocol() >= Protocol1v21v50 {
		r.UUID(&x.UUID)
	} else {
		u := x.UUID.String()
		r.String(&u)
		if parsed, err := uuid.Parse(u); err == nil {
			x.UUID = parsed
		}
	}
	r.String(&x.Version)
	r.Uint64(&x.Size)
	r.String(&x.ContentKey)
	r.String(&x.SubPackName)
	if r.Protocol() > Protocol1v5v0 {
		r.String(&x.ContentIdentity)
	}
	if r.Protocol() >= Protocol1v9v0 {
		r.Bool(&x.HasScripts)
	}
	if r.Protocol() >= Protocol1v16v200 {
		r.Bool(&x.AddonPack)
		r.Bool(&x.RTXEnabled)
	}
	if r.Protocol() >= Protocol1v21v40 {
		r.String(&x.DownloadURL)
	}
}

// StackResourcePack represents a resource pack sent on the stack of the client. When sent, the client will
// apply them in the order of the stack sent.
//
// Added: v1.13.0
type StackResourcePack struct {
	// UUID is the UUID of the resource pack. Each resource pack downloaded must have a different UUID in
	// order for the client to be able to handle them properly.
	//
	// Added: v1.13.0
	UUID string
	// Version is the version of the resource pack. The client will cache resource packs sent by the server as
	// long as they carry the same version. Sending a resource pack with a different version than previously
	// will force the client to re-download it.
	//
	// Added: v1.13.0
	Version string
	// SubPackName ...
	//
	// Added: v1.13.0
	SubPackName string
}

// Marshal encodes/decodes a StackResourcePack.
func (x *StackResourcePack) Marshal(r IO) {
	r.String(&x.UUID)
	r.String(&x.Version)
	if r.Protocol() >= Protocol1v8v0 {
		r.String(&x.SubPackName)
	}
}

// PackURL represents a resource pack that is being served from a HTTP server rather than being sent over
// the Minecraft protocol.
//
// Added: v1.20.30
type PackURL struct {
	// UUIDVersion is a combination of the UUID and version of the resource pack in the format UUID_Version.
	// The client will only attempt to download the resource pack if it does not already have it cached.
	//
	// Added: v1.20.30
	UUIDVersion string
	// URL is the URL from which the resource pack is downloaded. This URL must serve a zip file containing
	// a manifest.json file inside another folder. The manifest cannot be in the root of the zip file.
	//
	// Added: v1.20.30
	URL string
}

// Marshal encodes/decodes a CDNURL.
func (x *PackURL) Marshal(r IO) {
	r.String(&x.UUIDVersion)
	r.String(&x.URL)
}

// PackSetting represents a single setting from the pack settings UI. It holds information
// about the setting that was changed, including its name and the new value.
//
// Added: v1.21.111
type PackSetting struct {
	// Name is the name of the pack setting.
	//
	// Added: v1.21.111
	Name string
	// Value is the new value of the setting. This is either a float32, bool or string.
	//
	// Added: v1.21.111
	Value any
}
