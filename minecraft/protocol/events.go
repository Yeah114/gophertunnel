package protocol

const (
	EventTypeAchievementAwarded = iota
	EventTypeEntityInteract
	EventTypePortalBuilt
	EventTypePortalUsed
	EventTypeMobKilled
	EventTypeCauldronUsed
	EventTypePlayerDied
	EventTypeBossKilled
	EventTypeAgentCommand
	EventTypeAgentCreated
	EventTypePatternRemoved
	EventTypeSlashCommandExecuted
	EventTypeFishBucketed
	EventTypeMobBorn
	EventTypePetDied
	EventTypeCauldronInteract
	EventTypeComposterInteract
	EventTypeBellUsed
	EventTypeEntityDefinitionTrigger
	EventTypeRaidUpdate
	EventTypeMovementAnomaly
	EventTypeMovementCorrected
	EventTypeExtractHoney
	EventTypeTargetBlockHit
	EventTypePiglinBarter
	EventTypePlayerWaxedOrUnwaxedCopper
	EventTypeCodeBuilderRuntimeAction
	EventTypeCodeBuilderScoreboard
	EventTypeStriderRiddenInLavaInOverworld
	EventTypeSneakCloseToSculkSensor
	EventTypeCarefulRestoration
	EventTypeItemUsed
)

// lookupEvent looks up an Event matching the event type passed. False is
// returned if no such event data exists.
func lookupEvent(eventType int32, x *Event) bool {
	switch eventType {
	case EventTypeAchievementAwarded:
		*x = &AchievementAwardedEvent{}
	case EventTypeEntityInteract:
		*x = &EntityInteractEvent{}
	case EventTypePortalBuilt:
		*x = &PortalBuiltEvent{}
	case EventTypePortalUsed:
		*x = &PortalUsedEvent{}
	case EventTypeMobKilled:
		*x = &MobKilledEvent{}
	case EventTypeCauldronUsed:
		*x = &CauldronUsedEvent{}
	case EventTypePlayerDied:
		*x = &PlayerDiedEvent{}
	case EventTypeBossKilled:
		*x = &BossKilledEvent{}
	case EventTypeAgentCommand:
		*x = &AgentCommandEvent{}
	case EventTypeAgentCreated:
		*x = &AgentCreatedEvent{}
	case EventTypePatternRemoved:
		*x = &PatternRemovedEvent{}
	case EventTypeSlashCommandExecuted:
		*x = &SlashCommandExecutedEvent{}
	case EventTypeFishBucketed:
		*x = &FishBucketedEvent{}
	case EventTypeMobBorn:
		*x = &MobBornEvent{}
	case EventTypePetDied:
		*x = &PetDiedEvent{}
	case EventTypeCauldronInteract:
		*x = &CauldronInteractEvent{}
	case EventTypeComposterInteract:
		*x = &ComposterInteractEvent{}
	case EventTypeBellUsed:
		*x = &BellUsedEvent{}
	case EventTypeEntityDefinitionTrigger:
		*x = &EntityDefinitionTriggerEvent{}
	case EventTypeRaidUpdate:
		*x = &RaidUpdateEvent{}
	case EventTypeMovementAnomaly:
		*x = &MovementAnomalyEvent{}
	case EventTypeMovementCorrected:
		*x = &MovementCorrectedEvent{}
	case EventTypeExtractHoney:
		*x = &ExtractHoneyEvent{}
	case EventTypeTargetBlockHit:
		*x = &TargetBlockHitEvent{}
	case EventTypePiglinBarter:
		*x = &PiglinBarterEvent{}
	case EventTypePlayerWaxedOrUnwaxedCopper:
		*x = &WaxedOrUnwaxedCopperEvent{}
	case EventTypeCodeBuilderRuntimeAction:
		*x = &CodeBuilderRuntimeActionEvent{}
	case EventTypeCodeBuilderScoreboard:
		*x = &CodeBuilderScoreboardEvent{}
	case EventTypeStriderRiddenInLavaInOverworld:
		*x = &StriderRiddenInLavaInOverworldEvent{}
	case EventTypeSneakCloseToSculkSensor:
		*x = &SneakCloseToSculkSensorEvent{}
	case EventTypeCarefulRestoration:
		*x = &CarefulRestorationEvent{}
	case EventTypeItemUsed:
		*x = &ItemUsedEvent{}
	default:
		return false
	}
	return true
}

// lookupEventType looks up an event type that matches the Event passed.
func lookupEventType(x Event, eventType *int32) bool {
	switch x.(type) {
	case *AchievementAwardedEvent:
		*eventType = EventTypeAchievementAwarded
	case *EntityInteractEvent:
		*eventType = EventTypeEntityInteract
	case *PortalBuiltEvent:
		*eventType = EventTypePortalBuilt
	case *PortalUsedEvent:
		*eventType = EventTypePortalUsed
	case *MobKilledEvent:
		*eventType = EventTypeMobKilled
	case *CauldronUsedEvent:
		*eventType = EventTypeCauldronUsed
	case *PlayerDiedEvent:
		*eventType = EventTypePlayerDied
	case *BossKilledEvent:
		*eventType = EventTypeBossKilled
	case *AgentCommandEvent:
		*eventType = EventTypeAgentCommand
	case *AgentCreatedEvent:
		*eventType = EventTypeAgentCreated
	case *PatternRemovedEvent:
		*eventType = EventTypePatternRemoved
	case *SlashCommandExecutedEvent:
		*eventType = EventTypeSlashCommandExecuted
	case *FishBucketedEvent:
		*eventType = EventTypeFishBucketed
	case *MobBornEvent:
		*eventType = EventTypeMobBorn
	case *PetDiedEvent:
		*eventType = EventTypePetDied
	case *CauldronInteractEvent:
		*eventType = EventTypeCauldronInteract
	case *ComposterInteractEvent:
		*eventType = EventTypeComposterInteract
	case *BellUsedEvent:
		*eventType = EventTypeBellUsed
	case *EntityDefinitionTriggerEvent:
		*eventType = EventTypeEntityDefinitionTrigger
	case *RaidUpdateEvent:
		*eventType = EventTypeRaidUpdate
	case *MovementAnomalyEvent:
		*eventType = EventTypeMovementAnomaly
	case *MovementCorrectedEvent:
		*eventType = EventTypeMovementCorrected
	case *ExtractHoneyEvent:
		*eventType = EventTypeExtractHoney
	case *TargetBlockHitEvent:
		*eventType = EventTypeTargetBlockHit
	case *PiglinBarterEvent:
		*eventType = EventTypePiglinBarter
	case *WaxedOrUnwaxedCopperEvent:
		*eventType = EventTypePlayerWaxedOrUnwaxedCopper
	case *CodeBuilderRuntimeActionEvent:
		*eventType = EventTypeCodeBuilderRuntimeAction
	case *CodeBuilderScoreboardEvent:
		*eventType = EventTypeCodeBuilderScoreboard
	case *StriderRiddenInLavaInOverworldEvent:
		*eventType = EventTypeStriderRiddenInLavaInOverworld
	case *SneakCloseToSculkSensorEvent:
		*eventType = EventTypeSneakCloseToSculkSensor
	case *CarefulRestorationEvent:
		*eventType = EventTypeCarefulRestoration
	case *ItemUsedEvent:
		*eventType = EventTypeItemUsed
	default:
		return false
	}
	return true
}

// lookupEventOrdinal looks up the event ordinal for the Event passed.
func lookupEventOrdinal(x Event, eventOrdinal *uint32) bool {
	switch x.(type) {
	case *AchievementAwardedEvent:
		*eventOrdinal = 0
	case *EntityInteractEvent:
		*eventOrdinal = 1
	case *PortalBuiltEvent:
		*eventOrdinal = 2
	case *PortalUsedEvent:
		*eventOrdinal = 3
	case *MobKilledEvent:
		*eventOrdinal = 4
	case *CauldronUsedEvent:
		*eventOrdinal = 5
	case *PlayerDiedEvent:
		*eventOrdinal = 6
	case *BossKilledEvent:
		*eventOrdinal = 7
	case *SlashCommandExecutedEvent:
		*eventOrdinal = 8
	case *MobBornEvent:
		*eventOrdinal = 9
	case *CauldronInteractEvent:
		*eventOrdinal = 10
	case *ComposterInteractEvent:
		*eventOrdinal = 11
	case *BellUsedEvent:
		*eventOrdinal = 12
	case *EntityDefinitionTriggerEvent:
		*eventOrdinal = 13
	case *RaidUpdateEvent:
		*eventOrdinal = 14
	case *TargetBlockHitEvent:
		*eventOrdinal = 15
	case *PiglinBarterEvent:
		*eventOrdinal = 16
	case *WaxedOrUnwaxedCopperEvent:
		*eventOrdinal = 17
	case *CodeBuilderRuntimeActionEvent:
		*eventOrdinal = 18
	case *CodeBuilderScoreboardEvent:
		*eventOrdinal = 19
	case *ItemUsedEvent:
		*eventOrdinal = 20
	default:
		return false
	}
	return true
}

// Event represents an object that holds data specific to an event.
// The data it holds depends on the type.
type Event interface {
	// Marshal encodes/decodes a serialised event data object.
	Marshal(r IO)
}

// AchievementAwardedEvent is the event data sent for achievements.
//
// Added: v1.17.10
type AchievementAwardedEvent struct {
	// AchievementID is the ID for the achievement.
	//
	// Added: v1.17.10
	AchievementID uint8
}

// Marshal ...
func (a *AchievementAwardedEvent) Marshal(r IO) {
	r.Uint8(&a.AchievementID)
}

// EntityInteractEvent is the event data sent for entity interactions.
//
// Added: v1.17.10
type EntityInteractEvent struct {
	// InteractedEntityID is the unique ID of the entity that was interacted with.
	//
	// Added: v1.17.10
	InteractedEntityID int64
	// InteractionType is the type of interaction that occurred.
	//
	// Added: v1.17.10
	InteractionType uint8
	// InteractionEntityType is the entity type runtime ID of the interacted entity.
	//
	// Added: v1.17.10
	InteractionEntityType int32
	// EntityVariant is the variant value of the interacted entity.
	//
	// Added: v1.17.10
	EntityVariant int32
	// EntityColour is the colour variant of the interacted entity, if applicable.
	//
	// Added: v1.17.10
	EntityColour uint8
}

// Marshal ...
func (e *EntityInteractEvent) Marshal(r IO) {
	r.Varint64(&e.InteractedEntityID)
	r.Uint8(&e.InteractionType)
	r.Varint32(&e.InteractionEntityType)
	r.Varint32(&e.EntityVariant)
	r.Uint8(&e.EntityColour)
}

// PortalBuiltEvent is the event data sent when a portal is built.
//
// Added: v1.17.10
type PortalBuiltEvent struct {
	// DimensionID is the numeric dimension ID where the portal was built.
	//
	// Added: v1.17.10
	DimensionID int32
}

// Marshal ...
func (p *PortalBuiltEvent) Marshal(r IO) {
	r.Varint32(&p.DimensionID)
}

// PortalUsedEvent is the event data sent when a portal is used.
//
// Added: v1.17.10
type PortalUsedEvent struct {
	// FromDimensionID is the source dimension ID of the portal transfer.
	//
	// Added: v1.17.10
	FromDimensionID int32
	// ToDimensionID is the destination dimension ID of the portal transfer.
	//
	// Added: v1.17.10
	ToDimensionID int32
}

// Marshal ...
func (p *PortalUsedEvent) Marshal(r IO) {
	r.Varint32(&p.FromDimensionID)
	r.Varint32(&p.ToDimensionID)
}

// MobKilledEvent is the event data sent when a mob is killed.
//
// Added: v1.17.10
type MobKilledEvent struct {
	// KillerEntityUniqueID is the unique ID of the entity that dealt the killing blow.
	//
	// Added: v1.17.10
	KillerEntityUniqueID int64
	// VictimEntityUniqueID is the unique ID of the entity that died.
	//
	// Added: v1.17.10
	VictimEntityUniqueID int64
	// KillerEntityType is the entity type runtime ID of the killer.
	//
	// Added: v1.17.10
	KillerEntityType int32
	// EntityDamageCause is the damage cause that led to the kill.
	//
	// Added: v1.17.10
	EntityDamageCause int32
	// VillagerTradeTier is the villager trade tier, or -1 if the victim was not a trading actor.
	//
	// Added: v1.17.10
	VillagerTradeTier int32
	// VillagerDisplayName is the villager name, or empty if the victim was not a trading actor.
	//
	// Added: v1.17.10
	VillagerDisplayName string
}

// Marshal ...
func (m *MobKilledEvent) Marshal(r IO) {
	r.Varint64(&m.KillerEntityUniqueID)
	r.Varint64(&m.VictimEntityUniqueID)
	r.Varint32(&m.KillerEntityType)
	r.Varint32(&m.EntityDamageCause)
	r.Varint32(&m.VillagerTradeTier)
	r.String(&m.VillagerDisplayName)
}

// CauldronUsedEvent is the event data sent when a cauldron is used.
//
// Added: v1.17.10
type CauldronUsedEvent struct {
	// Colour is the ARGB colour value associated with the cauldron contents.
	//
	// Added: v1.17.10
	Colour uint32
	// PotionID is the potion type ID associated with the cauldron contents.
	//
	// Added: v1.17.10
	PotionID int16
	// FillLevel is the fill level of the cauldron after the action.
	//
	// Added: v1.17.10
	FillLevel int16
}

// Marshal ...
func (c *CauldronUsedEvent) Marshal(r IO) {
	r.Varuint32(&c.Colour)
	r.Int16(&c.PotionID)
	r.Int16(&c.FillLevel)
}

// PlayerDiedEvent is the event data sent when a player dies.
//
// Added: v1.17.10
type PlayerDiedEvent struct {
	// AttackerEntityID is the runtime ID of the attacking entity, if any.
	//
	// Added: v1.17.10
	AttackerEntityID int32
	// AttackerVariant is the variant value of the attacker.
	//
	// Added: v1.17.10
	AttackerVariant int32
	// EntityDamageCause is the damage cause that killed the player.
	//
	// Added: v1.17.10
	EntityDamageCause int32
	// InRaid specifies if the death occurred while the player was participating in a raid.
	//
	// Added: v1.19.21
	InRaid bool
}

// Marshal ...
func (p *PlayerDiedEvent) Marshal(r IO) {
	r.Varint32(&p.AttackerEntityID)
	r.Varint32(&p.AttackerVariant)
	r.Varint32(&p.EntityDamageCause)
	r.Bool(&p.InRaid)
}

// BossKilledEvent is the event data sent when a boss dies.
//
// Added: v1.17.10
type BossKilledEvent struct {
	// BossEntityUniqueID is the unique ID of the boss that was killed.
	//
	// Added: v1.17.10
	BossEntityUniqueID int64
	// PlayerPartySize is the size of the player party credited for the kill.
	//
	// Added: v1.17.10
	PlayerPartySize int32
	// InteractionEntityType is the entity type runtime ID of the boss.
	//
	// Added: v1.17.10
	InteractionEntityType int32
}

// Marshal ...
func (b *BossKilledEvent) Marshal(r IO) {
	r.Varint64(&b.BossEntityUniqueID)
	r.Varint32(&b.PlayerPartySize)
	r.Varint32(&b.InteractionEntityType)
}

// AgentCommandEvent is an event used in Education Edition.
//
// Added: v1.17.10
type AgentCommandEvent struct {
	// AgentResult is the result code of the agent command execution.
	//
	// Added: v1.17.10
	AgentResult int32
	// DataValue is an auxiliary numeric value associated with the command.
	//
	// Added: v1.17.10
	DataValue int32
	// Command is the agent command name that was executed.
	//
	// Added: v1.17.10
	Command string
	// DataKey is the key associated with the command payload.
	//
	// Added: v1.17.10
	DataKey string
	// Output is the textual output returned by the agent command.
	//
	// Added: v1.17.10
	Output string
}

// Marshal ...
func (a *AgentCommandEvent) Marshal(r IO) {
	r.Varint32(&a.AgentResult)
	r.Varint32(&a.DataValue)
	r.String(&a.Command)
	r.String(&a.DataKey)
	r.String(&a.Output)
}

// AgentCreatedEvent is the event data sent when an agent is created.
//
// Added: v1.21.111
type AgentCreatedEvent struct{}

// Marshal ...
func (a *AgentCreatedEvent) Marshal(r IO) {}

// PatternRemovedEvent is the event data sent when a pattern is removed.
//
// Added: v1.19.70
type PatternRemovedEvent struct{}

// Marshal ...
func (p *PatternRemovedEvent) Marshal(r IO) {}

// SlashCommandExecutedEvent is the event data sent when a slash command is executed.
//
// Added: v1.17.10
type SlashCommandExecutedEvent struct {
	// SuccessCount is the number of successful command executions.
	//
	// Added: v1.17.10
	SuccessCount int32
	// MessageCount indicates the amount of OutputMessages present.
	//
	// Added: v1.17.10
	MessageCount int32
	// CommandName is the slash command that was executed.
	//
	// Added: v1.17.10
	CommandName string
	// OutputMessages is a list of messages joined with `;`.
	//
	// Added: v1.17.10
	OutputMessages string
}

// Marshal ...
func (s *SlashCommandExecutedEvent) Marshal(r IO) {
	r.Varint32(&s.SuccessCount)
	r.Varint32(&s.MessageCount)
	r.String(&s.CommandName)
	r.String(&s.OutputMessages)
}

// FishBucketedEvent is the event data sent when a fish is bucketed.
//
// Added: v1.19.70
type FishBucketedEvent struct{}

// Marshal ...
func (f *FishBucketedEvent) Marshal(r IO) {}

// MobBornEvent is the event data sent when a mob is born.
//
// Added: v1.19.70
type MobBornEvent struct {
	// EntityType is the entity type runtime ID of the newly spawned mob.
	//
	// Added: v1.19.70
	EntityType int32
	// Variant is the mob variant value.
	//
	// Added: v1.19.70
	Variant int32
	// Colour is the colour variant of the mob, if applicable.
	//
	// Added: v1.19.70
	Colour uint8
}

// Marshal ...
func (m *MobBornEvent) Marshal(r IO) {
	r.Varint32(&m.EntityType)
	r.Varint32(&m.Variant)
	r.Uint8(&m.Colour)
}

// PetDiedEvent is the event data sent when a pet dies.
//
// Added: v1.19.70
type PetDiedEvent struct{}

// Marshal ...
func (p *PetDiedEvent) Marshal(r IO) {}

// CauldronInteractEvent is the event data sent when a cauldron is interacted with.
//
// Added: v1.19.70
type CauldronInteractEvent struct {
	// BlockInteractionType is the type of cauldron interaction that occurred.
	//
	// Added: v1.19.70
	BlockInteractionType uint8
	// ItemID is the numeric item ID involved in the interaction.
	//
	// Added: v1.19.70
	ItemID int16
}

// Marshal ...
func (c *CauldronInteractEvent) Marshal(r IO) {
	r.Uint8(&c.BlockInteractionType)
	r.Int16(&c.ItemID)
}

// ComposterInteractEvent is the event data sent when a composter is interacted with.
//
// Added: v1.19.70
type ComposterInteractEvent struct {
	// BlockInteractionType is the type of composter interaction that occurred.
	//
	// Added: v1.19.70
	BlockInteractionType uint8
	// ItemID is the numeric item ID involved in the interaction.
	//
	// Added: v1.19.70
	ItemID int16
}

// Marshal ...
func (c *ComposterInteractEvent) Marshal(r IO) {
	r.Uint8(&c.BlockInteractionType)
	r.Int16(&c.ItemID)
}

// BellUsedEvent is the event data sent when a bell is used.
//
// Added: v1.19.70
type BellUsedEvent struct {
	// ItemID is the numeric item ID involved in ringing the bell.
	//
	// Added: v1.19.70
	ItemID int16
}

// Marshal ...
func (b *BellUsedEvent) Marshal(r IO) {
	r.Int16(&b.ItemID)
}

// EntityDefinitionTriggerEvent is an event used for an unknown purpose.
//
// Added: v1.19.70
type EntityDefinitionTriggerEvent struct {
	// EventName is the identifier of the triggered entity definition event.
	//
	// Added: v1.19.70
	EventName string
}

// Marshal ...
func (e *EntityDefinitionTriggerEvent) Marshal(r IO) {
	r.String(&e.EventName)
}

// RaidUpdateEvent is an event used to update raid progress client side.
//
// Added: v1.19.70
type RaidUpdateEvent struct {
	// CurrentRaidWave is the current raid wave number.
	//
	// Added: v1.19.70
	CurrentRaidWave int32
	// TotalRaidWaves is the total amount of raid waves.
	//
	// Added: v1.19.70
	TotalRaidWaves int32
	// WonRaid specifies if the raid was won.
	//
	// Added: v1.19.70
	WonRaid bool
}

// Marshal ...
func (ra *RaidUpdateEvent) Marshal(r IO) {
	r.Varint32(&ra.CurrentRaidWave)
	r.Varint32(&ra.TotalRaidWaves)
	r.Bool(&ra.WonRaid)
}

// MovementAnomalyEvent is an event used to detect movement anomalies.
//
// Added: v1.19.70
type MovementAnomalyEvent struct{}

// Marshal ...
func (m *MovementAnomalyEvent) Marshal(r IO) {}

// MovementCorrectedEvent is an event used to correct movement anomalies.
//
// Added: v1.19.70
type MovementCorrectedEvent struct{}

// Marshal ...
func (m *MovementCorrectedEvent) Marshal(r IO) {}

// ExtractHoneyEvent is an event used to extract honey from a hive.
//
// Added: v1.19.70
type ExtractHoneyEvent struct{}

// Marshal ...
func (e *ExtractHoneyEvent) Marshal(r IO) {}

// TargetBlockHitEvent is an event used when a target block is hit by an arrow.
//
// Added: v1.21.20
type TargetBlockHitEvent struct {
	// RedstoneLevel is the redstone strength produced by the hit target block.
	//
	// Added: v1.21.20
	RedstoneLevel int32
}

// Marshal ...
func (t *TargetBlockHitEvent) Marshal(r IO) {
	r.Varint32(&t.RedstoneLevel)
}

// PiglinBarterEvent is called when a player drops gold ingots to a piglin to initiate a trade for an item.
//
// Added: v1.21.20
type PiglinBarterEvent struct {
	// ItemID is the numeric item ID that was bartered.
	//
	// Added: v1.21.20
	ItemID int32
	// WasTargetingBarteringPlayer specifies if the piglin was targeting the player who started the barter.
	//
	// Added: v1.21.20
	WasTargetingBarteringPlayer bool
}

// Marshal ...
func (p *PiglinBarterEvent) Marshal(r IO) {
	r.Varint32(&p.ItemID)
	r.Bool(&p.WasTargetingBarteringPlayer)
}

const (
	WaxNotOxidised   = uint16(0xa609)
	WaxExposed       = uint16(0xa809)
	WaxWeathered     = uint16(0xaa09)
	WaxOxidised      = uint16(0xac09)
	UnWaxNotOxidised = uint16(0xae09)
	UnWaxExposed     = uint16(0xb009)
	UnWaxWeathered   = uint16(0xb209)
	UnWaxOxidised    = uint16(0xfa0a)
)

// WaxedOrUnwaxedCopperEvent is an event sent by the server when a copper block is waxed or unwaxed.
//
// Added: v1.19.70
type WaxedOrUnwaxedCopperEvent struct {
	// CopperBlockID is the block runtime ID of the waxed or unwaxed copper block.
	//
	// Added: v1.19.70
	CopperBlockID int32
}

// Marshal ...
func (w *WaxedOrUnwaxedCopperEvent) Marshal(r IO) {
	r.Varint32(&w.CopperBlockID)
}

// CodeBuilderRuntimeActionEvent is an event sent by the server when a code builder runtime action is performed.
//
// Added: v1.21.20
type CodeBuilderRuntimeActionEvent struct {
	// Action is the code builder runtime action that was performed.
	//
	// Added: v1.21.20
	Action string
}

// Marshal ...
func (c *CodeBuilderRuntimeActionEvent) Marshal(r IO) {
	r.String(&c.Action)
}

// CodeBuilderScoreboardEvent is an event sent by the server when a code builder scoreboard is updated.
//
// Added: v1.21.20
type CodeBuilderScoreboardEvent struct {
	// ObjectiveName is the objective name of the updated code builder scoreboard.
	//
	// Added: v1.21.20
	ObjectiveName string
	// Score is the updated score value.
	//
	// Added: v1.21.20
	Score int32
}

// Marshal ...
func (c *CodeBuilderScoreboardEvent) Marshal(r IO) {
	r.String(&c.ObjectiveName)
	r.Varint32(&c.Score)
}

// StriderRiddenInLavaInOverworldEvent is an event sent by the server when a strider is ridden in lava in the overworld.
//
// Added: v1.21.20
type StriderRiddenInLavaInOverworldEvent struct{}

// Marshal ...
func (s *StriderRiddenInLavaInOverworldEvent) Marshal(r IO) {}

// SneakCloseToSculkSensorEvent is an event sent by the server when a player sneaks close to a sculk sensor.
//
// Added: v1.19.70
type SneakCloseToSculkSensorEvent struct{}

// Marshal ...
func (s *SneakCloseToSculkSensorEvent) Marshal(r IO) {}

// CarefulRestorationEvent is an event sent by the server when a player performs a careful restoration.
//
// Added: v1.21.20
type CarefulRestorationEvent struct{}

// Marshal ...
func (c *CarefulRestorationEvent) Marshal(r IO) {}

// ItemUsedEvent is sent when a player right clicks an item.
//
// Added: v1.21.0
type ItemUsedEvent struct {
	// ItemID is the numeric item ID of the used item.
	//
	// Added: v1.21.0
	ItemID int16
	// ItemAux is the auxiliary data value of the used item.
	//
	// Added: v1.21.0
	ItemAux int32
	// UseMethod identifies how the item was used.
	//
	// Added: v1.21.0
	UseMethod int32
	// UseCount is the amount of use operations that were performed.
	//
	// Added: v1.21.0
	UseCount int32
}

// Marshal ...
func (i *ItemUsedEvent) Marshal(r IO) {
	r.Int16(&i.ItemID)
	r.Int32(&i.ItemAux)
	r.Int32(&i.UseMethod)
	r.Int32(&i.UseCount)
}
