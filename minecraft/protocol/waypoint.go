package protocol

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

const (
	WaypointActionNone = iota
	WaypointActionAdd
	WaypointActionRemove
	WaypointActionUpdate
)

const (
	WaypointUpdateFlagVisible = 1 << iota
	WaypointUpdateFlagPosition
	WaypointUpdateFlagTextureID
	WaypointUpdateFlagColour
	WaypointUpdateFlagClientPositionAuthority
	WaypointUpdateFlagActorUniqueID
)

const (
	WaypointTextureSquare = iota + 2
	WaypointTextureCircle
	WaypointTextureSmallSquare
	WaypointTextureSmallStar
)

// LocatorBarWaypoint represents a waypoint entry in the locator bar packet.
// Added: v1.26.10
type LocatorBarWaypoint struct {
	// GroupHandle is the UUID handle for the waypoint group.
	// Added: v1.26.10
	GroupHandle uuid.UUID
	// Waypoint contains the waypoint data.
	// Added: v1.26.10
	Waypoint Waypoint
	// Action determines the action for this waypoint. It is one of the WaypointAction constants.
	// Added: v1.26.10
	Action uint8
}

// Marshal encodes/decodes a LocatorBarWaypoint.
func (x *LocatorBarWaypoint) Marshal(r IO) {
	r.UUID(&x.GroupHandle)
	Single(r, &x.Waypoint)
	r.Uint8(&x.Action)
}

// WaypointWorldPosition holds a position and dimension for a waypoint.
// Added: v1.26.10
type WaypointWorldPosition struct {
	// Position is the world position of the waypoint.
	// Added: v1.26.10
	Position mgl32.Vec3
	// DimensionID is the dimension the waypoint is in.
	// Added: v1.26.10
	DimensionID int32
}

// Marshal encodes/decodes a WaypointWorldPosition.
func (x *WaypointWorldPosition) Marshal(r IO) {
	r.Vec3(&x.Position)
	r.Varint32(&x.DimensionID)
}

// Waypoint holds optional data for a locator bar waypoint.
// Added: v1.26.10
type Waypoint struct {
	// UpdateFlag is a bitmask indicating which optional fields are set.
	// Added: v1.26.10
	UpdateFlag uint32
	// Visible determines whether the waypoint is shown.
	// Added: v1.26.10
	Visible Optional[bool]
	// WorldPosition is the position and dimension of the waypoint.
	// Added: v1.26.10
	WorldPosition Optional[WaypointWorldPosition]
	// TextureID is the legacy icon texture identifier of the waypoint. It is one of the WaypointTexture constants.
	// Removed: v1.26.20.26
	TextureID Optional[uint32]
	// TexturePath is the resource path for the waypoint icon texture.
	// Added: v1.26.20.26
	TexturePath Optional[string]
	// IconSize is the rendered size of the waypoint icon.
	// Added: v1.26.20.26
	IconSize Optional[mgl32.Vec2]
	// Colour is the RGB colour used to tint the waypoint icon.
	// Added: v1.26.10
	Colour Optional[int32]
	// ClientPositionAuthority determines whether the client has authority over the waypoint position.
	// Added: v1.26.10
	ClientPositionAuthority Optional[bool]
	// ActorUniqueID is the unique ID of the entity the waypoint tracks.
	// Added: v1.26.10
	ActorUniqueID Optional[int64]
}

// Marshal encodes/decodes a Waypoint.
func (x *Waypoint) Marshal(r IO) {
	r.Uint32(&x.UpdateFlag)
	OptionalFunc(r, &x.Visible, r.Bool)
	OptionalMarshaler(r, &x.WorldPosition)
	if r.Protocol() >= Protocol1v26v20v26 {
		OptionalFunc(r, &x.TexturePath, r.String)
		OptionalFunc(r, &x.IconSize, r.Vec2)
	} else {
		OptionalFunc(r, &x.TextureID, r.Uint32)
	}
	OptionalFunc(r, &x.Colour, r.Int32)
	OptionalFunc(r, &x.ClientPositionAuthority, r.Bool)
	OptionalFunc(r, &x.ActorUniqueID, r.Varint64)
}
