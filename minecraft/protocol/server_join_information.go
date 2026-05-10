package protocol

import "github.com/google/uuid"

// GatheringJoinInfo contains information about the gathering (experience) the player is joining. This type was added
// in v1.26.0.
type GatheringJoinInfo struct {
	// ExperienceID is the UUID of the experience.
	ExperienceID uuid.UUID
	// ExperienceID1v26v0 is the ID of the experience.
	// This field was removed in v1.26.10.
	ExperienceID1v26v0 string
	// ExperienceName is the name of the experience.
	ExperienceName string
	// ExperienceWorldID is the UUID of the experience world.
	ExperienceWorldID uuid.UUID
	// ExperienceWorldID1v26v0 is the world ID of the experience.
	// This field was removed in v1.26.10.
	ExperienceWorldID1v26v0 string
	// ExperienceWorldName is the world name of the experience.
	ExperienceWorldName string
	// CreatorID is the ID of the creator.
	CreatorID string
	// StoreID is the store ID.
	// This field was removed in v1.26.10.
	StoreID string
	// TargetID is the session ID of the experience. This field was added in v1.26.20.26.
	TargetID uuid.UUID
	// ScenarioID is the scenario ID of experience. This field was added in v1.26.20.26.
	ScenarioID string
	// UnknownUUID1 is an unknown UUID field.
	// This field was added in v1.26.10 and removed in v1.26.20.26.
	UnknownUUID1 uuid.UUID
	// UnknownUUID2 is an unknown UUID field.
	// This field was added in v1.26.10 and removed in v1.26.20.26.
	UnknownUUID2 uuid.UUID
	// ServerID is the server identifier. This field was added in v1.26.20.26.
	ServerID string
}

// Marshal encodes/decodes a GatheringJoinInfo.
func (x *GatheringJoinInfo) Marshal(r IO) {
	if r.Protocol() < Protocol1v26v10 {
		r.String(&x.ExperienceID1v26v0)
	} else {
		r.UUID(&x.ExperienceID)
	}
	r.String(&x.ExperienceName)
	if r.Protocol() < Protocol1v26v10 {
		r.String(&x.ExperienceWorldID1v26v0)
	} else {
		r.UUID(&x.ExperienceWorldID)
	}
	r.String(&x.ExperienceWorldName)
	r.String(&x.CreatorID)
	if r.Protocol() < Protocol1v26v10 {
		r.String(&x.StoreID)
	}
	if r.Protocol() >= Protocol1v26v20v26 {
		r.UUID(&x.TargetID)
		r.String(&x.ScenarioID)
	} else if r.Protocol() >= Protocol1v26v10 {
		r.UUID(&x.UnknownUUID1)
		r.UUID(&x.UnknownUUID2)
	}
	if r.Protocol() >= Protocol1v26v20v26 {
		r.String(&x.ServerID)
	}
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
	if r.Protocol() >= Protocol1v26v10 {
		OptionalMarshaler(r, &x.StoreEntryPointInfo)
		OptionalMarshaler(r, &x.PresenceInfo)
	}
}
