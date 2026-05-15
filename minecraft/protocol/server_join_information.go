package protocol

import "github.com/google/uuid"

// GatheringJoinInfo contains information about the gathering (experience) the player is joining.
// Added: v1.26.0
type GatheringJoinInfo struct {
	// ExperienceID is the UUID of the experience.
	// Added: v1.26.10
	ExperienceID uuid.UUID
	// ExperienceID1v26v0 is the legacy string form of the experience identifier.
	// Removed: v1.26.10
	ExperienceID1v26v0 string
	// ExperienceName is the name of the experience.
	// Added: v1.26.0
	ExperienceName string
	// ExperienceWorldID is the UUID of the experience world.
	// Added: v1.26.10
	ExperienceWorldID uuid.UUID
	// ExperienceWorldID1v26v0 is the legacy string form of the experience world identifier.
	// Removed: v1.26.10
	ExperienceWorldID1v26v0 string
	// ExperienceWorldName is the world name of the experience.
	// Added: v1.26.0
	ExperienceWorldName string
	// CreatorID is the ID of the creator.
	// Added: v1.26.0
	CreatorID string
	// StoreID is the legacy store identifier bundled directly with the gathering info.
	// Removed: v1.26.10
	StoreID string
	// TargetID is the session identifier for the target experience instance.
	// Added: v1.26.20.26
	TargetID uuid.UUID
	// ScenarioID is the scenario identifier associated with the experience.
	// Added: v1.26.20.26
	ScenarioID string
	// UnknownUUID1 is an unknown UUID field present in the v1.26.10 through v1.26.20 payload layout.
	// Added: v1.26.10, Removed: v1.26.20.26
	UnknownUUID1 uuid.UUID
	// UnknownUUID2 is an unknown UUID field present in the v1.26.10 through v1.26.20 payload layout.
	// Added: v1.26.10, Removed: v1.26.20.26
	UnknownUUID2 uuid.UUID
	// ServerID is the server identifier associated with the target experience.
	// Added: v1.26.20.26
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

// StoreEntryPointInfo contains information about the store entry point.
// Added: v1.26.10
type StoreEntryPointInfo struct {
	// StoreID is the store identifier.
	// Added: v1.26.10
	StoreID string
	// StoreName is the display name of the store entry point.
	// Added: v1.26.10
	StoreName string
}

// Marshal encodes/decodes a StoreEntryPointInfo.
func (x *StoreEntryPointInfo) Marshal(r IO) {
	r.String(&x.StoreID)
	r.String(&x.StoreName)
}

// PresenceInfo contains presence information about the experience.
// Added: v1.26.10
type PresenceInfo struct {
	// ExperienceName is the experience name shown in presence data.
	// Added: v1.26.10
	ExperienceName string
	// WorldName is the world name shown in presence data.
	// Added: v1.26.10
	WorldName string
}

// Marshal encodes/decodes a PresenceInfo.
func (x *PresenceInfo) Marshal(r IO) {
	r.String(&x.ExperienceName)
	r.String(&x.WorldName)
}

// ServerJoinInformation contains optional information about the server the player is joining.
// Added: v1.26.0
type ServerJoinInformation struct {
	// GatheringJoinInfo is optional information about the gathering being joined.
	// Added: v1.26.0
	GatheringJoinInfo Optional[GatheringJoinInfo]
	// StoreEntryPointInfo is optional information about the store entry point that opened the join flow.
	// Added: v1.26.10
	StoreEntryPointInfo Optional[StoreEntryPointInfo]
	// PresenceInfo is optional rich presence information for the experience and world being joined.
	// Added: v1.26.10
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
