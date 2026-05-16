package protocol

import (
	"image/color"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	AimAssistTargetModeAngle = iota
	AimAssistTargetModeDistance
)

const (
	AudioListenerCamera = iota
	AudioListenerPlayer
)

const (
	EasingTypeLinear = iota
	EasingTypeSpring
	EasingTypeInQuad
	EasingTypeOutQuad
	EasingTypeInOutQuad
	EasingTypeInCubic
	EasingTypeOutCubic
	EasingTypeInOutCubic
	EasingTypeInQuart
	EasingTypeOutQuart
	EasingTypeInOutQuart
	EasingTypeInQuint
	EasingTypeOutQuint
	EasingTypeInOutQuint
	EasingTypeInSine
	EasingTypeOutSine
	EasingTypeInOutSine
	EasingTypeInExpo
	EasingTypeOutExpo
	EasingTypeInOutExpo
	EasingTypeInCirc
	EasingTypeOutCirc
	EasingTypeInOutCirc
	EasingTypeInBounce
	EasingTypeOutBounce
	EasingTypeInOutBounce
	EasingTypeInBack
	EasingTypeOutBack
	EasingTypeInOutBack
	EasingTypeInElastic
	EasingTypeOutElastic
	EasingTypeInOutElastic
	EasingTypeInverseLerp
)

// easingTypeFromString looks up an easing type from a string and writes the result to x.
func easingTypeFromString(io IO, x *int32, s string) {
	switch s {
	case "linear":
		*x = EasingTypeLinear
	case "spring":
		*x = EasingTypeSpring
	case "in_quad":
		*x = EasingTypeInQuad
	case "out_quad":
		*x = EasingTypeOutQuad
	case "in_out_quad":
		*x = EasingTypeInOutQuad
	case "in_cubic":
		*x = EasingTypeInCubic
	case "out_cubic":
		*x = EasingTypeOutCubic
	case "in_out_cubic":
		*x = EasingTypeInOutCubic
	case "in_quart":
		*x = EasingTypeInQuart
	case "out_quart":
		*x = EasingTypeOutQuart
	case "in_out_quart":
		*x = EasingTypeInOutQuart
	case "in_quint":
		*x = EasingTypeInQuint
	case "out_quint":
		*x = EasingTypeOutQuint
	case "in_out_quint":
		*x = EasingTypeInOutQuint
	case "in_sine":
		*x = EasingTypeInSine
	case "out_sine":
		*x = EasingTypeOutSine
	case "in_out_sine":
		*x = EasingTypeInOutSine
	case "in_expo":
		*x = EasingTypeInExpo
	case "out_expo":
		*x = EasingTypeOutExpo
	case "in_out_expo":
		*x = EasingTypeInOutExpo
	case "in_circ":
		*x = EasingTypeInCirc
	case "out_circ":
		*x = EasingTypeOutCirc
	case "in_out_circ":
		*x = EasingTypeInOutCirc
	case "in_back":
		*x = EasingTypeInBack
	case "out_back":
		*x = EasingTypeOutBack
	case "in_out_back":
		*x = EasingTypeInOutBack
	case "in_elastic":
		*x = EasingTypeInElastic
	case "out_elastic":
		*x = EasingTypeOutElastic
	case "in_out_elastic":
		*x = EasingTypeInOutElastic
	case "in_bounce":
		*x = EasingTypeInBounce
	case "out_bounce":
		*x = EasingTypeOutBounce
	case "in_out_bounce":
		*x = EasingTypeInOutBounce
	case "inverse_lerp":
		*x = EasingTypeInverseLerp
	default:
		io.InvalidValue(s, "easingType", "unknown easing type")
	}
}

// easingTypeToString looks up an easing type constant and returns the string representation.
func easingTypeToString(x int32) string {
	switch x {
	case EasingTypeLinear:
		return "linear"
	case EasingTypeSpring:
		return "spring"
	case EasingTypeInQuad:
		return "in_quad"
	case EasingTypeOutQuad:
		return "out_quad"
	case EasingTypeInOutQuad:
		return "in_out_quad"
	case EasingTypeInCubic:
		return "in_cubic"
	case EasingTypeOutCubic:
		return "out_cubic"
	case EasingTypeInOutCubic:
		return "in_out_cubic"
	case EasingTypeInQuart:
		return "in_quart"
	case EasingTypeOutQuart:
		return "out_quart"
	case EasingTypeInOutQuart:
		return "in_out_quart"
	case EasingTypeInQuint:
		return "in_quint"
	case EasingTypeOutQuint:
		return "out_quint"
	case EasingTypeInOutQuint:
		return "in_out_quint"
	case EasingTypeInSine:
		return "in_sine"
	case EasingTypeOutSine:
		return "out_sine"
	case EasingTypeInOutSine:
		return "in_out_sine"
	case EasingTypeInExpo:
		return "in_expo"
	case EasingTypeOutExpo:
		return "out_expo"
	case EasingTypeInOutExpo:
		return "in_out_expo"
	case EasingTypeInCirc:
		return "in_circ"
	case EasingTypeOutCirc:
		return "out_circ"
	case EasingTypeInOutCirc:
		return "in_out_circ"
	case EasingTypeInBack:
		return "in_back"
	case EasingTypeOutBack:
		return "out_back"
	case EasingTypeInOutBack:
		return "in_out_back"
	case EasingTypeInElastic:
		return "in_elastic"
	case EasingTypeOutElastic:
		return "out_elastic"
	case EasingTypeInOutElastic:
		return "in_out_elastic"
	case EasingTypeInBounce:
		return "in_bounce"
	case EasingTypeOutBounce:
		return "out_bounce"
	case EasingTypeInOutBounce:
		return "in_out_bounce"
	case EasingTypeInverseLerp:
		return "inverse_lerp"
	default:
		return "unknown"
	}
}

const (
	SplineTypeCatmullRom = "catmullrom"
	SplineTypeLinear     = "linear"
)

func splineTypeFromUint8(io IO, x *string, v uint8) {
	switch v {
	case 0:
		*x = SplineTypeCatmullRom
	case 1:
		*x = SplineTypeLinear
	default:
		io.InvalidValue(v, "splineType", "unknown spline type")
	}
}

func splineTypeToUint8(io IO, x string) uint8 {
	switch x {
	case SplineTypeCatmullRom:
		return 0
	case SplineTypeLinear:
		return 1
	default:
		io.InvalidValue(x, "splineType", "unknown spline type")
		return 0
	}
}

// CameraEase represents an easing function that can be used by a CameraInstructionSet.
//
// Added: v1.20.30
type CameraEase struct {
	// Type is the type of easing function used. This is one of the constants above.
	//
	// Added: v1.20.30
	Type uint8
	// Duration is the time in seconds that the easing function should take.
	//
	// Added: v1.20.30
	Duration float32
}

// Marshal encodes/decodes a CameraEase.
func (x *CameraEase) Marshal(r IO) {
	r.Uint8(&x.Type)
	r.Float32(&x.Duration)
}

// CameraInstructionSet represents a camera instruction that sets the camera to a specified preset and can be extended
// with easing functions and translations to the camera's position and rotation.
//
// Added: v1.20.30
type CameraInstructionSet struct {
	// Preset is the index of the preset in the CameraPresets packet sent to the player.
	//
	// Added: v1.20.30
	Preset uint32
	// Ease represents the easing function that is used by the instruction.
	//
	// Added: v1.20.30
	Ease Optional[CameraEase]
	// Position represents the position of the camera.
	//
	// Added: v1.20.30
	Position Optional[mgl32.Vec3]
	// Rotation represents the rotation of the camera.
	//
	// Added: v1.20.30
	Rotation Optional[mgl32.Vec2]
	// Facing is a vector that the camera will always face towards during the duration of the instruction.
	//
	// Added: v1.20.30
	Facing Optional[mgl32.Vec3]
	// ViewOffset is an offset based on a pivot point to the player, causing the camera to be shifted in a
	// certain direction.
	//
	// Added: v1.21.30
	ViewOffset Optional[mgl32.Vec2]
	// EntityOffset is an offset from the entity that the camera should be rendered at.
	//
	// Added: v1.21.30
	EntityOffset Optional[mgl32.Vec3]
	// Default determines whether the camera is a default camera or not.
	//
	// Added: v1.20.30
	Default Optional[bool]
	// IgnoreStartingValuesComponent behavior is currently unknown.
	//
	// Added: v1.21.90
	IgnoreStartingValuesComponent bool
}

// Marshal encodes/decodes a CameraInstructionSet.
func (x *CameraInstructionSet) Marshal(r IO) {
	r.Uint32(&x.Preset)
	OptionalMarshaler(r, &x.Ease)
	OptionalFunc(r, &x.Position, r.Vec3)
	OptionalFunc(r, &x.Rotation, r.Vec2)
	OptionalFunc(r, &x.Facing, r.Vec3)
	OptionalFunc(r, &x.ViewOffset, r.Vec2)
	OptionalFunc(r, &x.EntityOffset, r.Vec3)
	OptionalFunc(r, &x.Default, r.Bool)
	r.Bool(&x.IgnoreStartingValuesComponent)
}

// CameraFadeTimeData represents the time data for a CameraInstructionFade.
//
// Added: v1.21.20
type CameraFadeTimeData struct {
	// FadeInDuration is the time in seconds for the screen to fully fade in.
	//
	// Added: v1.21.20
	FadeInDuration float32
	// WaitDuration is time in seconds to wait before fading out.
	//
	// Added: v1.21.20
	WaitDuration float32
	// FadeOutDuration is the time in seconds for the screen to fully fade out.
	//
	// Added: v1.21.20
	FadeOutDuration float32
}

// Marshal encodes/decodes a CameraFadeTimeData.
func (x *CameraFadeTimeData) Marshal(r IO) {
	r.Float32(&x.FadeInDuration)
	r.Float32(&x.WaitDuration)
	r.Float32(&x.FadeOutDuration)
}

// CameraInstructionFade represents a camera instruction that fades the screen to a specified colour.
//
// Added: v1.21.20
type CameraInstructionFade struct {
	// TimeData is the time data for the fade, which includes the fade in duration, wait duration and fade out
	// duration.
	//
	// Added: v1.21.20
	TimeData Optional[CameraFadeTimeData]
	// Colour is the colour of the screen to fade to. This only uses the red, green and blue components.
	//
	// Added: v1.21.20
	Colour Optional[color.RGBA]
}

// Marshal encodes/decodes a CameraInstructionFade.
func (x *CameraInstructionFade) Marshal(r IO) {
	OptionalMarshaler(r, &x.TimeData)
	OptionalFunc(r, &x.Colour, r.RGB)
}

// CameraInstructionTarget represents a camera instruction that targets a specific entity.
//
// Added: v1.21.20
type CameraInstructionTarget struct {
	// CenterOffset is the offset from the center of the entity that the camera should target.
	//
	// Added: v1.21.20
	CenterOffset Optional[mgl32.Vec3]
	// EntityUniqueID is the unique ID of the entity that the camera should target.
	//
	// Added: v1.21.20
	EntityUniqueID int64
}

// Marshal encodes/decodes a CameraInstructionTarget.
func (x *CameraInstructionTarget) Marshal(r IO) {
	OptionalFunc(r, &x.CenterOffset, r.Vec3)
	r.Int64(&x.EntityUniqueID)
}

// CameraInstructionFieldOfView represents a camera instruction that updates the field of view.
//
// Added: v1.21.100
type CameraInstructionFieldOfView struct {
	// FieldOfView is the field of view of the camera.
	//
	// Added: v1.21.100
	FieldOfView float32
	// EaseTime is the time in seconds that the easing function should take.
	//
	// Added: v1.21.100
	EaseTime float32
	// EaseType is the type of easing function used. This is one of the constants above.
	//
	// Changed: v1.26.10, encoded as a string easing identifier instead of a uint8 enum value.
	EaseType int32
	// Clear can be set to true to clear the current instruction.
	//
	// Added: v1.21.100
	Clear bool
}

// Marshal encodes/decodes a CameraInstructionFieldOfView.
func (x *CameraInstructionFieldOfView) Marshal(r IO) {
	r.Float32(&x.FieldOfView)
	r.Float32(&x.EaseTime)
	if r.Protocol() >= Protocol1v26v10 {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	} else {
		easeType := uint8(x.EaseType)
		r.Uint8(&easeType)
		x.EaseType = int32(easeType)
	}
	r.Bool(&x.Clear)
}

// CameraPreset represents a basic preset that can be extended upon by more complex instructions.
//
// Added: v1.20.30
type CameraPreset struct {
	// Name is the name of the preset. Each preset must have their own unique name.
	//
	// Added: v1.20.30
	Name string
	// Parent is the name of the preset that this preset extends upon. This can be left empty.
	//
	// Added: v1.20.30
	Parent string
	// PosX is the default X position of the camera.
	//
	// Added: v1.20.30
	PosX Optional[float32]
	// PosY is the default Y position of the camera.
	//
	// Added: v1.20.30
	PosY Optional[float32]
	// PosZ is the default Z position of the camera.
	//
	// Added: v1.20.30
	PosZ Optional[float32]
	// RotX is the default pitch of the camera.
	//
	// Added: v1.20.30
	RotX Optional[float32]
	// RotY is the default yaw of the camera.
	//
	// Added: v1.20.30
	RotY Optional[float32]
	// RotationSpeed is the speed at which the camera should rotate.
	//
	// Added: v1.21.30
	RotationSpeed Optional[float32]
	// SnapToTarget determines whether the camera should snap to the target entity or not.
	//
	// Added: v1.21.30
	SnapToTarget Optional[bool]
	// HorizontalRotationLimit is the horizontal rotation limit of the camera.
	//
	// Added: v1.21.40
	HorizontalRotationLimit Optional[mgl32.Vec2]
	// VerticalRotationLimit is the vertical rotation limit of the camera.
	//
	// Added: v1.21.40
	VerticalRotationLimit Optional[mgl32.Vec2]
	// ContinueTargeting determines whether the camera should continue targeting when using aim assist.
	//
	// Added: v1.21.40
	ContinueTargeting Optional[bool]
	// TrackingRadius is the radius around the camera that the aim assist should track targets.
	//
	// Added: v1.21.40
	TrackingRadius Optional[float32]
	// ViewOffset is only used in a follow_orbit camera and controls an offset based on a pivot point to the
	// player, causing it to be shifted in a certain direction.
	//
	// Added: v1.21.20
	ViewOffset Optional[mgl32.Vec2]
	// EntityOffset controls the offset from the entity that the camera should be rendered at.
	//
	// Added: v1.21.30
	EntityOffset Optional[mgl32.Vec3]
	// Radius is only used in a follow_orbit camera and controls how far away from the player the camera should
	// be rendered.
	//
	// Added: v1.21.20
	Radius Optional[float32]
	// MinYawLimit is the minimum yaw limit of the camera.
	//
	// Added: v1.21.60
	MinYawLimit Optional[float32]
	// MaxYawLimit is the maximum yaw limit of the camera.
	//
	// Added: v1.21.60
	MaxYawLimit Optional[float32]
	// AudioListener defines where the audio should be played from when using this preset. This is one of the
	// constants above.
	//
	// Added: v1.20.30
	AudioListener Optional[byte]
	// PlayerEffects is currently unknown.
	//
	// Added: v1.20.30
	PlayerEffects Optional[bool]
	// AimAssist defines the aim assist to use when using this preset.
	//
	// Added: v1.21.50
	AimAssist Optional[CameraPresetAimAssist]
	// ControlScheme is the control scheme that the client should use in this camera. It is one of the following:
	//  - ControlSchemeLockedPlayerRelativeStrafe is the default behaviour, this cannot be set when the client
	//    is in a custom camera.
	//  - ControlSchemeCameraRelative makes movement relative to the camera's transform, with the client's
	//    rotation being relative to the client's movement.
	//  - ControlSchemeCameraRelativeStrafe makes movement relative to the camera's transform, with the
	//    client's rotation being locked.
	//  - ControlSchemePlayerRelative makes movement relative to the player's transform, meaning holding
	//    left/right will make the player turn in a circle.
	//  - ControlSchemePlayerRelativeStrafe makes movement the same as the default behaviour, but can be
	//    used in a custom camera.
	//
	// Added: v1.21.80
	ControlScheme Optional[byte]
}

// Marshal encodes/decodes a CameraPreset.
func (x *CameraPreset) Marshal(r IO) {
	r.String(&x.Name)
	r.String(&x.Parent)
	OptionalFunc(r, &x.PosX, r.Float32)
	OptionalFunc(r, &x.PosY, r.Float32)
	OptionalFunc(r, &x.PosZ, r.Float32)
	OptionalFunc(r, &x.RotX, r.Float32)
	OptionalFunc(r, &x.RotY, r.Float32)
	OptionalFunc(r, &x.RotationSpeed, r.Float32)
	OptionalFunc(r, &x.SnapToTarget, r.Bool)
	OptionalFunc(r, &x.HorizontalRotationLimit, r.Vec2)
	OptionalFunc(r, &x.VerticalRotationLimit, r.Vec2)
	OptionalFunc(r, &x.ContinueTargeting, r.Bool)
	OptionalFunc(r, &x.TrackingRadius, r.Float32)
	OptionalFunc(r, &x.ViewOffset, r.Vec2)
	OptionalFunc(r, &x.EntityOffset, r.Vec3)
	OptionalFunc(r, &x.Radius, r.Float32)
	OptionalFunc(r, &x.MinYawLimit, r.Float32)
	OptionalFunc(r, &x.MaxYawLimit, r.Float32)
	OptionalFunc(r, &x.AudioListener, r.Uint8)
	OptionalFunc(r, &x.PlayerEffects, r.Bool)
	OptionalMarshaler(r, &x.AimAssist)
	OptionalFunc(r, &x.ControlScheme, r.Uint8)
}

// CameraPresetAimAssist represents a preset for aim assist settings.
//
// Added: v1.21.50
type CameraPresetAimAssist struct {
	// Preset is the ID of the preset that has previously been defined in the CameraAimAssistPresets packet.
	//
	// Added: v1.21.50
	Preset Optional[string]
	// TargetMode is the mode that the camera should use for detecting targets. This is one of the constants
	// above.
	//
	// Added: v1.21.50
	TargetMode Optional[int32]
	// Angle is the maximum angle around the playes's cursor that the aim assist should check for a target,
	// if TargetMode is set to protocol.AimAssistTargetModeAngle.
	//
	// Added: v1.21.50
	Angle Optional[mgl32.Vec2]
	// Distance is the maximum distance from the player's cursor should check for a target, if TargetMode is
	// set to protocol.AimAssistTargetModeDistance.
	//
	// Added: v1.21.50
	Distance Optional[float32]
}

// Marshal encodes/decodes a CameraPresetAimAssist.
func (x *CameraPresetAimAssist) Marshal(r IO) {
	OptionalFunc(r, &x.Preset, r.String)
	OptionalFunc(r, &x.TargetMode, r.Int32)
	OptionalFunc(r, &x.Angle, r.Vec2)
	OptionalFunc(r, &x.Distance, r.Float32)
}

// CameraAimAssistCategory is an aim assist category that defines priorities for specific blocks and entities.
//
// Added: v1.21.50
type CameraAimAssistCategory struct {
	// Name is the name of the category which can be used by a CameraAimAssistPreset.
	//
	// Added: v1.21.50
	Name string
	// Priorities represents the block and entity specific priorities as well as the default priorities for
	// this category.
	//
	// Added: v1.21.50
	Priorities CameraAimAssistPriorities
}

// Marshal encodes/decodes a CameraAimAssistCategory.
func (x *CameraAimAssistCategory) Marshal(r IO) {
	r.String(&x.Name)
	Single(r, &x.Priorities)
}

// CameraAimAssistPriorities represents the block and entity specific priorities for targetting. The aim
// assist will select the block or entity with the highest priority within the specified thresholds.
//
// Added: v1.21.50
type CameraAimAssistPriorities struct {
	// Entities is a list of priorities for specific entity identifiers.
	//
	// Added: v1.21.50
	Entities []CameraAimAssistPriority
	// Blocks is a list of priorities for specific block identifiers.
	//
	// Added: v1.21.50
	Blocks []CameraAimAssistPriority
	// BlockTags is a list of priorities for specific block tags.
	//
	// Added: v1.21.130
	BlockTags []CameraAimAssistPriority
	// EntityTypeFamilies is a list of priorities for specific entity type families.
	//
	// Added: v1.26.0
	EntityTypeFamilies []CameraAimAssistPriority
	// EntityDefault is the default priority for entities.
	//
	// Added: v1.21.50
	EntityDefault Optional[int32]
	// BlockDefault is the default priority for blocks.
	//
	// Added: v1.21.50
	BlockDefault Optional[int32]
}

// Marshal encodes/decodes a CameraAimAssistPriorities.
func (x *CameraAimAssistPriorities) Marshal(r IO) {
	Slice(r, &x.Entities)
	Slice(r, &x.Blocks)
	if r.Protocol() >= Protocol1v21v130 {
		Slice(r, &x.BlockTags)
	}
	if r.Protocol() >= Protocol1v26v0 {
		Slice(r, &x.EntityTypeFamilies)
	}
	OptionalFunc(r, &x.EntityDefault, r.Int32)
	OptionalFunc(r, &x.BlockDefault, r.Int32)
}

// CameraAimAssistPriority represents a non-default priority for a specific target.
//
// Added: v1.21.50
type CameraAimAssistPriority struct {
	// Identifier is the identifier of a target to define the priority for.
	//
	// Added: v1.21.50
	Identifier string
	// Priority is the priority for this specific target.
	//
	// Added: v1.21.50
	Priority int32
}

// Marshal encodes/decodes a CameraAimAssistPriority.
func (x *CameraAimAssistPriority) Marshal(r IO) {
	r.String(&x.Identifier)
	r.Int32(&x.Priority)
}

// CameraAimAssistPreset defines a base preset that can be extended upon when sending an aim assist.
//
// Added: v1.21.50
type CameraAimAssistPreset struct {
	// Identifier represents the identifier of this preset.
	//
	// Added: v1.21.50
	Identifier string
	// BlockExclusions is a list of block identifiers that should be ignored by the aim assist.
	//
	// Added: v1.21.50
	BlockExclusions []string
	// EntityExclusions is a list of entity identifiers that should be ignored by the aim assist.
	//
	// Added: v1.21.130
	EntityExclusions []string
	// BlockTagExclusions is a list of block tags that should be ignored by the aim assist.
	//
	// Added: v1.21.130
	BlockTagExclusions []string
	// EntityTypeFamilyExclusions is a list of entity type families that should be ignored by the aim assist.
	//
	// Added: v1.26.0
	EntityTypeFamilyExclusions []string
	// LiquidTargets is a list of entity identifiers that should be targetted when inside of a liquid.
	//
	// Added: v1.21.50
	LiquidTargets []string
	// ItemSettings is a list of settings for specific item identifiers. If an item is not listed here, it
	// will fallback to DefaultItemSettings or HandSettings if no item is held.
	//
	// Added: v1.21.50
	ItemSettings []CameraAimAssistItemSettings
	// DefaultItemSettings is the identifier of a category to use when the player is not holding an item
	// listed in ItemSettings. This must be the identifier of a category within the Categories slice.
	//
	// Added: v1.21.50
	DefaultItemSettings Optional[string]
	// HandSettings is the identifier of a category to use when the player is not holding an item. This must
	// be the identifier of a category within Categories slice.
	//
	// Added: v1.21.50
	HandSettings Optional[string]
}

// Marshal encodes/decodes a CameraAimAssistPreset.
func (x *CameraAimAssistPreset) Marshal(r IO) {
	r.String(&x.Identifier)
	FuncSlice(r, &x.BlockExclusions, r.String)
	if r.Protocol() >= Protocol1v21v130 {
		FuncSlice(r, &x.EntityExclusions, r.String)
		FuncSlice(r, &x.BlockTagExclusions, r.String)
	}
	if r.Protocol() >= Protocol1v26v0 {
		FuncSlice(r, &x.EntityTypeFamilyExclusions, r.String)
	}
	FuncSlice(r, &x.LiquidTargets, r.String)
	Slice(r, &x.ItemSettings)
	OptionalFunc(r, &x.DefaultItemSettings, r.String)
	OptionalFunc(r, &x.HandSettings, r.String)
}

// CameraAimAssistItemSettings defines settings for how specific items should behave when using aim assist.
//
// Added: v1.21.50
type CameraAimAssistItemSettings struct {
	// Item is the identifier of the item to apply the settings to.
	//
	// Added: v1.21.50
	Item string
	// Category is the identifier of a category to use which has been defined by a CameraAimAssistCategory.
	// Only categories defined in the Categories slice used by the CameraAimAssistPreset can be
	// used here.
	//
	// Added: v1.21.50
	Category string
}

// Marshal encodes/decodes a CameraAimAssistItemSettings.
func (x *CameraAimAssistItemSettings) Marshal(r IO) {
	r.String(&x.Item)
	r.String(&x.Category)
}

// CameraRotationOption represents a rotation option for camera spline instructions.
//
// Added: v1.21.120
type CameraRotationOption struct {
	// Value is the rotation value.
	//
	// Added: v1.21.120
	Value mgl32.Vec3
	// Time is the time for this rotation option.
	//
	// Added: v1.21.120
	Time float32
	// EaseType is the easing function used to interpolate towards this rotation key frame.
	//
	// Changed: v1.26.10, encoded as a string easing identifier instead of a uint8 enum value.
	EaseType int32
}

// Marshal encodes/decodes a CameraRotationOption.
func (x *CameraRotationOption) Marshal(r IO) {
	r.Vec3(&x.Value)
	r.Float32(&x.Time)
	if r.Protocol() >= Protocol1v26v10 {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	} else {
		easeType := uint8(x.EaseType)
		r.Uint8(&easeType)
		x.EaseType = int32(easeType)
	}
}

// CameraProgressOption represents a progress keyframe option for camera spline instructions.
//
// Added: v1.26.0
type CameraProgressOption struct {
	// Value is the progress value.
	//
	// Added: v1.26.0
	Value float32
	// Time is the time for this progress option.
	//
	// Added: v1.26.0
	Time float32
	// EaseType is the easing function used to interpolate towards this progress key frame.
	//
	// Changed: v1.26.10, encoded as a string easing identifier instead of a uint8 enum value.
	EaseType int32
}

// Marshal encodes/decodes a CameraProgressOption.
func (x *CameraProgressOption) Marshal(r IO) {
	r.Float32(&x.Value)
	r.Float32(&x.Time)
	if r.Protocol() >= Protocol1v26v10 {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	} else {
		easeType := uint8(x.EaseType)
		r.Uint8(&easeType)
		x.EaseType = int32(easeType)
	}
}

// CameraSplineInstruction represents a camera instruction that creates a spline path for the camera to follow.
//
// Added: v1.21.120
type CameraSplineInstruction struct {
	// TotalTime is the total time for the spline animation.
	//
	// Added: v1.21.120
	TotalTime float32
	// SplineType is the optional spline interpolation type used for the spline curve.
	//
	// Changed: v1.26.10, encoded as an optional string identifier instead of an optional uint8 enum value.
	SplineType Optional[string]
	// Curve is a list of points that define the spline curve.
	//
	// Added: v1.21.120
	Curve []mgl32.Vec3
	// ProgressKeyFrames is a list of progress key frames for the spline.
	//
	// Added: v1.21.120
	ProgressKeyFrames []CameraProgressOption
	// RotationOptions is a list of rotation options for the spline.
	//
	// Added: v1.21.120
	RotationOptions []CameraRotationOption
	// SplineIdentifier is an optional identifier for referencing the spline by name.
	//
	// Added: v1.26.10
	SplineIdentifier Optional[string]
	// LoadFromJson optionally determines whether the spline should be loaded from a JSON definition.
	//
	// Added: v1.26.10
	LoadFromJson Optional[bool]
}

// Marshal encodes/decodes a CameraSplineInstruction.
func (x *CameraSplineInstruction) Marshal(r IO) {
	r.Float32(&x.TotalTime)
	if r.Protocol() >= Protocol1v26v10 {
		OptionalFunc(r, &x.SplineType, r.String)
	} else {
		var splineType Optional[uint8]
		if value, ok := x.SplineType.Value(); ok {
			splineType = Option(splineTypeToUint8(r, value))
		}
		OptionalFunc(r, &splineType, r.Uint8)
		if value, ok := splineType.Value(); ok {
			x.SplineType = Option("")
			splineTypeFromUint8(r, &x.SplineType.val, value)
			x.SplineType.set = true
		} else {
			x.SplineType = Optional[string]{}
		}
	}
	FuncSlice(r, &x.Curve, r.Vec3)
	Slice(r, &x.ProgressKeyFrames)
	Slice(r, &x.RotationOptions)
	if r.Protocol() >= Protocol1v26v10 {
		OptionalFunc(r, &x.SplineIdentifier, r.String)
		OptionalFunc(r, &x.LoadFromJson, r.Bool)
	}
}

// CameraSplineDefinition represents a named camera spline definition.
//
// Added: v1.26.0
type CameraSplineDefinition struct {
	// Name is the name of the spline definition.
	//
	// Added: v1.26.0
	Name string
	// Instruction is the nested spline instruction payload used by the v1.26.0 layout.
	//
	// Removed: v1.26.10
	Instruction CameraSplineInstruction
	// TotalTime is the total time for the spline animation in the flattened v1.26.10+ layout.
	//
	// Added: v1.26.10
	TotalTime float32
	// SplineType is the optional spline interpolation type in the flattened v1.26.10+ layout.
	//
	// Added: v1.26.10
	SplineType Optional[string]
	// ControlPoints is a list of points that define the spline curve in the flattened v1.26.10+ layout.
	//
	// Added: v1.26.10
	ControlPoints []mgl32.Vec3
	// ProgressKeyFrames is a list of progress key frames for the spline in the flattened v1.26.10+ layout.
	//
	// Added: v1.26.10
	ProgressKeyFrames []CameraProgressOption
	// RotationKeyFrames is a list of rotation key frames for the spline in the flattened v1.26.10+ layout.
	//
	// Added: v1.26.10
	RotationKeyFrames []CameraRotationOption
}

// Marshal encodes/decodes a CameraSplineDefinition.
func (x *CameraSplineDefinition) Marshal(r IO) {
	r.String(&x.Name)
	if r.Protocol() >= Protocol1v26v10 {
		if x.Instruction.TotalTime != 0 || x.Instruction.SplineType.set || len(x.Instruction.Curve) != 0 ||
			len(x.Instruction.ProgressKeyFrames) != 0 || len(x.Instruction.RotationOptions) != 0 {
			x.TotalTime = x.Instruction.TotalTime
			x.SplineType = x.Instruction.SplineType
			x.ControlPoints = x.Instruction.Curve
			x.ProgressKeyFrames = x.Instruction.ProgressKeyFrames
			x.RotationKeyFrames = x.Instruction.RotationOptions
		}
		r.Float32(&x.TotalTime)
		OptionalFunc(r, &x.SplineType, r.String)
		FuncSlice(r, &x.ControlPoints, r.Vec3)
		Slice(r, &x.ProgressKeyFrames)
		Slice(r, &x.RotationKeyFrames)
		x.Instruction = CameraSplineInstruction{
			TotalTime:         x.TotalTime,
			SplineType:        x.SplineType,
			Curve:             x.ControlPoints,
			ProgressKeyFrames: x.ProgressKeyFrames,
			RotationOptions:   x.RotationKeyFrames,
		}
		return
	}

	instruction := x.Instruction
	if instruction.TotalTime == 0 && !instruction.SplineType.set && len(instruction.Curve) == 0 &&
		len(instruction.ProgressKeyFrames) == 0 && len(instruction.RotationOptions) == 0 {
		instruction = CameraSplineInstruction{
			TotalTime:         x.TotalTime,
			SplineType:        x.SplineType,
			Curve:             x.ControlPoints,
			ProgressKeyFrames: x.ProgressKeyFrames,
			RotationOptions:   x.RotationKeyFrames,
		}
	}
	Single(r, &instruction)
	x.Instruction = instruction
	x.TotalTime = instruction.TotalTime
	x.SplineType = instruction.SplineType
	x.ControlPoints = instruction.Curve
	x.ProgressKeyFrames = instruction.ProgressKeyFrames
	x.RotationKeyFrames = instruction.RotationOptions
}

// CameraAimAssistActorPriorityData represents priority data for aim assist actor targeting.
//
// Added: v1.26.0
type CameraAimAssistActorPriorityData struct {
	// PresetIndex is the index of the aim assist preset.
	//
	// Added: v1.26.0
	PresetIndex int32
	// CategoryIndex is the index of the aim assist category.
	//
	// Added: v1.26.0
	CategoryIndex int32
	// ActorIndex is the index of the actor.
	//
	// Added: v1.26.0
	ActorIndex int32
	// Priority is the priority value for this actor.
	//
	// Added: v1.26.0
	Priority int32
}

// Marshal encodes/decodes a CameraAimAssistActorPriorityData.
func (x *CameraAimAssistActorPriorityData) Marshal(r IO) {
	r.Int32(&x.PresetIndex)
	r.Int32(&x.CategoryIndex)
	r.Int32(&x.ActorIndex)
	r.Int32(&x.Priority)
}
