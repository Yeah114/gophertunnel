package protocol

import "github.com/google/uuid"

// GatheringJoinInfo contains information about the gathering (experience) the player is joining. This type was added
// in v1.26.0.
type GatheringJoinInfo struct {
	// ExperienceID is the UUID of the experience.
	ExperienceID uuid.UUID
	// ExperienceName is the name of the experience.
	ExperienceName string
	// ExperienceWorldID is the UUID of the experience world.
	ExperienceWorldID uuid.UUID
	// ExperienceWorldName is the world name of the experience.
	ExperienceWorldName string
	// CreatorID is the ID of the creator.
	CreatorID string
	// TargetID is the session ID of the experience. This field was added in v1.26.20.
	TargetID uuid.UUID
	// ScenarioID is the scenario ID of experience. This field was added in v1.26.20.
	ScenarioID string
	// ServerID is the server identifier. This field was added in v1.26.20.
	ServerID string
}

// Marshal encodes/decodes a GatheringJoinInfo.
func (x *GatheringJoinInfo) Marshal(r IO) {
	r.UUID(&x.ExperienceID)
	r.String(&x.ExperienceName)
	r.UUID(&x.ExperienceWorldID)
	r.String(&x.ExperienceWorldName)
	r.String(&x.CreatorID)
	r.UUID(&x.TargetID)
	r.String(&x.ScenarioID)
	r.String(&x.ServerID)
}

// StoreEntryPointInfo contains information about the store entry point. This type was added in v1.26.10.
type StoreEntryPointInfo struct {
	// StoreID is the store identifier.
	StoreID string
	// StoreName is the store name.
	StoreName string
}

// Marshal encodes/decodes a StoreEntryPointInfo.
func (x *StoreEntryPointInfo) Marshal(r IO) {
	r.String(&x.StoreID)
	r.String(&x.StoreName)
}

// PresenceInfo contains presence information about the experience. This type was added in v1.26.10.
type PresenceInfo struct {
	// ExperienceName is the name of the experience.
	ExperienceName string
	// WorldName is the name of the world.
	WorldName string
}

// Marshal encodes/decodes a PresenceInfo.
func (x *PresenceInfo) Marshal(r IO) {
	r.String(&x.ExperienceName)
	r.String(&x.WorldName)
}

// ServerJoinInformation contains optional information about the server the player is joining. This type was added in
// v1.26.0.
type ServerJoinInformation struct {
	// GatheringJoinInfo is optional information about the gathering being joined.
	GatheringJoinInfo Optional[GatheringJoinInfo]
	// StoreEntryPointInfo is optional information about the store entry point. This field was added in v1.26.10.
	StoreEntryPointInfo Optional[StoreEntryPointInfo]
	// PresenceInfo is optional presence information. This field was added in v1.26.10.
	PresenceInfo Optional[PresenceInfo]
}

// Marshal encodes/decodes a ServerJoinInformation.
func (x *ServerJoinInformation) Marshal(r IO) {
	OptionalMarshaler(r, &x.GatheringJoinInfo)
	OptionalMarshaler(r, &x.StoreEntryPointInfo)
	OptionalMarshaler(r, &x.PresenceInfo)
}
