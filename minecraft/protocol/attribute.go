package protocol

const (
	AttributeModifierOperationAddition = iota
	AttributeModifierOperationMultiplyBase
	AttributeModifierOperationMultiplyTotal
	AttributeModifierOperationCap
)

const (
	AttributeModifierOperandMin = iota
	AttributeModifierOperandMax
	AttributeModifierOperandCurrent
)

// AttributeValue holds the value of an attribute, including the min and max.
//
// Added: v1.16
type AttributeValue struct {
	// Name is the name of the attribute, for example 'minecraft:health'. These names must be identical to
	// the ones defined client-side.
	//
	// Added: v1.16
	Name string
	// Value is the current value of the attribute. This value will be applied to the entity when sent in a
	// packet.
	//
	// Added: v1.16
	Value float32
	// Max and Min specify the boundaries within the value of the attribute must be. The definition of these
	// fields differ per attribute. The maximum health of an entity may be changed, whereas the maximum
	// movement speed for example may not be.
	//
	// Added: v1.16
	Max, Min float32
}

// Marshal encodes/decodes an AttributeValue.
func (x *AttributeValue) Marshal(r IO) {
	r.String(&x.Name)
	r.Float32(&x.Min)
	r.Float32(&x.Value)
	r.Float32(&x.Max)
}

// Attribute is an entity attribute, that holds specific data such as the health of the entity. Each attribute
// holds a default value, maximum and minimum value, name and its current value.
//
// Added: v1.16
type Attribute struct {
	// AttributeValue holds the attribute name and its current value range.
	//
	// Added: v1.16
	AttributeValue
	// DefaultMin is the default minimum value of the attribute. It's not clear why this field must be sent to
	// the client, but it is required regardless.
	//
	// Added: v1.16
	DefaultMin float32
	// DefaultMax is the default maximum value of the attribute. It's not clear why this field must be sent to
	// the client, but it is required regardless.
	//
	// Added: v1.16
	DefaultMax float32
	// Default is the default value of the attribute. It's not clear why this field must be sent to the
	// client, but it is required regardless.
	//
	// Added: v1.16
	Default float32
	// Modifiers is a slice of AttributeModifiers that are applied to the attribute.
	//
	// Added: v1.19.20
	Modifiers []AttributeModifier
}

// Marshal encodes/decodes an Attribute.
func (x *Attribute) Marshal(r IO) {
	r.Float32(&x.Min)
	r.Float32(&x.Max)
	r.Float32(&x.Value)
	r.Float32(&x.DefaultMin)
	r.Float32(&x.DefaultMax)
	r.Float32(&x.Default)
	r.String(&x.Name)
	Slice(r, &x.Modifiers)
}

// AttributeModifier temporarily buffs/debuffs a given attribute until the modifier is used. In vanilla, these are
// mainly used for effects.
//
// Added: v1.19.20
type AttributeModifier struct {
	// ID is the unique ID of the modifier. It is used to identify the modifier in the packet.
	//
	// Added: v1.19.20
	ID string
	// Name is the name of the attribute that is modified.
	//
	// Added: v1.19.20
	Name string
	// Amount is the amount of difference between the current value of the attribute and the new value.
	//
	// Added: v1.19.20
	Amount float32
	// Operation is the operation that is performed on the attribute. It can be addition, multiply base, multiply total
	// or cap.
	//
	// Added: v1.19.20
	Operation int32
	// Operand identifies which attribute operand the modifier targets.
	// TODO: Figure out what this field is used for exactly in vanilla.
	//
	// Added: v1.19.20
	Operand int32
	// Serializable specifies if the modifier should be preserved when it is stored or forwarded.
	// TODO: Figure out what this field is used for exactly in vanilla.
	//
	// Added: v1.19.20
	Serializable bool
}

// Marshal encodes/decodes an AttributeModifier.
func (x *AttributeModifier) Marshal(r IO) {
	r.String(&x.ID)
	r.String(&x.Name)
	r.Float32(&x.Amount)
	r.Int32(&x.Operation)
	r.Int32(&x.Operand)
	r.Bool(&x.Serializable)
}
