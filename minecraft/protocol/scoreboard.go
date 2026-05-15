package protocol

const (
	ScoreboardIdentityPlayer = iota + 1
	ScoreboardIdentityEntity
	ScoreboardIdentityFakePlayer
)

// ScoreboardEntry represents a single entry that may be found on a scoreboard. These entries represent a
// line on the scoreboard each.
//
// Added: v1.16
type ScoreboardEntry struct {
	// EntryID is a unique identifier of this entry. Each entry that represents a different value should get
	// its own entry ID. When modifying a scoreboard, entries that represent the same line should have the
	// same entry ID.
	//
	// Added: v1.16
	EntryID int64
	// ObjectiveName is the name of the objective that this scoreboard entry is for. It must be identical to
	// the one set in the SetDisplayObjective packet previously sent.
	//
	// Added: v1.16
	ObjectiveName string
	// Score is the score that the entry represents. Scoreboard entries are ordered using this score, so in
	// order to get the scoreboard to be ordered as expected when trying to write non-score related text on
	// a scoreboard, this score should be incremented for each entry.
	//
	// Added: v1.16
	Score int32
	// IdentityType is the identity type of the scoreboard entry. The entry may represent an entity, player or
	// a fake player, as the constants above indicate.
	// In order to write plain text to the scoreboard, ScoreboardIdentityFakePlayer should always be used, in
	// combination with the DisplayName field. A different identity type will use the name of the entity.
	//
	// Added: v1.16
	IdentityType byte
	// EntityUniqueID is the unique ID of either the player or the entity represented by the scoreboard entry.
	// This field is only used if IdentityType is either ScoreboardIdentityEntity or ScoreboardIdentityPlayer.
	//
	// Added: v1.16
	EntityUniqueID int64
	// DisplayName is the custom name of the scoreboard entry. This field is only used if IdentityType is
	// ScoreboardIdentityFakePlayer. If this identity type is not used, the name of the entity/player will be
	// shown instead.
	//
	// Added: v1.16
	DisplayName string
}

// Marshal encodes/decodes a ScoreboardEntry x as an entry for modification.
func (x *ScoreboardEntry) Marshal(r IO) {
	ScoreRemoveEntry(r, x)
	r.Uint8(&x.IdentityType)
	switch x.IdentityType {
	case ScoreboardIdentityEntity, ScoreboardIdentityPlayer:
		r.Varint64(&x.EntityUniqueID)
	case ScoreboardIdentityFakePlayer:
		r.String(&x.DisplayName)
	default:
		r.UnknownEnumOption(x.IdentityType, "scoreboard entry identity type")
	}
}

// ScoreRemoveEntry encodes/decodes a ScoreboardEntry x as an entry for removal.
func ScoreRemoveEntry(r IO, x *ScoreboardEntry) {
	r.Varint64(&x.EntryID)
	r.String(&x.ObjectiveName)
	r.Int32(&x.Score)
}

// ScoreboardIdentityEntry holds an entry to either associate an identity with one of the entries in a
// scoreboard, or to remove associations.
//
// Added: v1.16
type ScoreboardIdentityEntry struct {
	// EntryID is the unique identifier of the entry that the identity should be associated with, or that
	// associations should be cleared from.
	//
	// Added: v1.16
	EntryID int64
	// EntityUniqueID is the unique ID that the entry should be associated with. It is empty if the
	// SetScoreboardIdentity packet is sent to remove associations with identities.
	//
	// Added: v1.16
	EntityUniqueID int64
}

// Marshal encodes/decodes a ScoreboardIdentityEntry.
func (x *ScoreboardIdentityEntry) Marshal(r IO) {
	r.Varint64(&x.EntryID)
	r.Varint64(&x.EntityUniqueID)
}

// ScoreboardIdentityClearEntry encodes/decodes a ScoreboardIdentityEntry for clearing the entry.
func ScoreboardIdentityClearEntry(r IO, x *ScoreboardIdentityEntry) {
	r.Varint64(&x.EntryID)
}
