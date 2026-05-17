package protocol

const (
	AttributeLayerPayloadTypeUpdateLayers = iota
	AttributeLayerPayloadTypeUpdateSettings
	AttributeLayerPayloadTypeUpdateEnvironment
	AttributeLayerPayloadTypeRemoveEnvironment
)

const (
	AttributeDataTypeBool = iota
	AttributeDataTypeFloat
	AttributeDataTypeColour
)

const (
	AttributeBoolOperationOverride = iota
	AttributeBoolOperationAlphaBlend
	AttributeBoolOperationAnd
	AttributeBoolOperationNand
	AttributeBoolOperationOr
	AttributeBoolOperationNor
	AttributeBoolOperationXor
	AttributeBoolOperationXnor
)

const (
	AttributeFloatOperationOverride = iota
	AttributeFloatOperationAlphaBlend
	AttributeFloatOperationAdd
	AttributeFloatOperationSubtract
	AttributeFloatOperationMultiply
	AttributeFloatOperationMinimum
	AttributeFloatOperationMaximum
)

const (
	AttributeColourOperationOverride = iota
	AttributeColourOperationAlphaBlend
	AttributeColourOperationAdd
	AttributeColourOperationSubtract
	AttributeColourOperationMultiply
)

const (
	AttributeColourValueTypeString = iota
	AttributeColourValueTypeArray
)

const (
	AttributeLayerWeightTypeFloat = iota
	AttributeLayerWeightTypeString
)

// AttributeData represents a polymorphic attribute value.
//
// Added: v1.26.10
type AttributeData struct {
	// Type is the attribute data type. It is one of the AttributeDataType constants.
	//
	// Added: v1.26.10
	Type uint32
	// BoolValue is the boolean value if Type is AttributeDataTypeBool.
	//
	// Added: v1.26.10
	BoolValue bool
	// BoolOperation is the optional operation for boolean attributes.
	//
	// Added: v1.26.10
	BoolOperation Optional[int32]
	// FloatValue is the float value if Type is AttributeDataTypeFloat.
	//
	// Added: v1.26.10
	FloatValue float32
	// FloatOperation is the optional operation for float attributes.
	//
	// Added: v1.26.10
	FloatOperation Optional[int32]
	// FloatConstraintMin is the optional minimum constraint for float attributes.
	//
	// Added: v1.26.10
	FloatConstraintMin Optional[float32]
	// FloatConstraintMax is the optional maximum constraint for float attributes.
	//
	// Added: v1.26.10
	FloatConstraintMax Optional[float32]
	// ColourValue is the colour value if Type is AttributeDataTypeColour.
	//
	// Deprecated: Use ColourType, ColourString and ColourValues for the protocol union representation.
	ColourValue int32
	// ColourType is the colour value type if Type is AttributeDataTypeColour. It is one of the
	// AttributeColourValueType constants.
	//
	// Added: v1.26.10
	ColourType uint32
	// ColourString is the colour value if ColourType is AttributeColourValueTypeString.
	//
	// Added: v1.26.10
	ColourString string
	// ColourValues is the RGBA colour value if ColourType is AttributeColourValueTypeArray.
	//
	// Added: v1.26.10
	ColourValues [4]int32
	// ColourOperation is the optional operation for colour attributes.
	//
	// Added: v1.26.10
	ColourOperation Optional[int32]
}

// Marshal encodes/decodes an AttributeData.
func (x *AttributeData) Marshal(r IO) {
	r.Varuint32(&x.Type)
	switch x.Type {
	case AttributeDataTypeBool:
		r.Bool(&x.BoolValue)
		operation := attributeOperationToString(r, x.BoolOperation, attributeBoolOperations, "bool attribute operation")
		r.String(&operation)
		attributeOperationFromString(r, &x.BoolOperation, operation, attributeBoolOperations, "bool attribute operation")
	case AttributeDataTypeFloat:
		r.Float32(&x.FloatValue)
		operation := attributeOperationToString(r, x.FloatOperation, attributeFloatOperations, "float attribute operation")
		r.String(&operation)
		attributeOperationFromString(r, &x.FloatOperation, operation, attributeFloatOperations, "float attribute operation")
		OptionalFunc(r, &x.FloatConstraintMin, r.Float32)
		OptionalFunc(r, &x.FloatConstraintMax, r.Float32)
	case AttributeDataTypeColour:
		colourType := x.ColourType
		if colourType == AttributeColourValueTypeString && x.ColourValue != 0 {
			colourType = AttributeColourValueTypeArray
		}
		r.Varuint32(&colourType)
		x.ColourType = colourType
		switch x.ColourType {
		case AttributeColourValueTypeString:
			r.String(&x.ColourString)
		case AttributeColourValueTypeArray:
			if x.ColourValues == [4]int32{} && x.ColourValue != 0 {
				x.ColourValues = [4]int32{x.ColourValue}
			}
			for i := range x.ColourValues {
				r.Int32(&x.ColourValues[i])
			}
			if x.ColourValue == 0 {
				x.ColourValue = x.ColourValues[0]
			}
		default:
			r.UnknownEnumOption(x.ColourType, "attribute colour value type")
		}
		operation := attributeOperationToString(r, x.ColourOperation, attributeColourOperations, "colour attribute operation")
		r.String(&operation)
		attributeOperationFromString(r, &x.ColourOperation, operation, attributeColourOperations, "colour attribute operation")
	default:
		r.UnknownEnumOption(x.Type, "attribute data type")
	}
}

var (
	attributeBoolOperations   = []string{"override", "alpha_blend", "and", "nand", "or", "nor", "xor", "xnor"}
	attributeFloatOperations  = []string{"override", "alpha_blend", "add", "subtract", "multiply", "minimum", "maximum"}
	attributeColourOperations = []string{"override", "alpha_blend", "add", "subtract", "multiply"}
)

func attributeOperationToString(r IO, x Optional[int32], values []string, field string) string {
	operation, _ := x.Value()
	if operation < 0 || int(operation) >= len(values) {
		r.InvalidValue(operation, field, "unknown attribute operation")
		return values[0]
	}
	return values[operation]
}

func attributeOperationFromString(r IO, x *Optional[int32], value string, values []string, field string) {
	for operation, name := range values {
		if name == value {
			*x = Option(int32(operation))
			return
		}
	}
	r.InvalidValue(value, field, "unknown attribute operation")
}

// EnvironmentAttributeData represents an environment attribute with optional transition data.
//
// Added: v1.26.10
type EnvironmentAttributeData struct {
	// AttributeName is the name of the attribute.
	//
	// Added: v1.26.10
	AttributeName string
	// FromAttribute is the optional starting attribute for transitions.
	//
	// Added: v1.26.10
	FromAttribute Optional[AttributeData]
	// Attribute is the current attribute value.
	//
	// Added: v1.26.10
	Attribute AttributeData
	// ToAttribute is the optional target attribute for transitions.
	//
	// Added: v1.26.10
	ToAttribute Optional[AttributeData]
	// CurrentTransitionTicks is the number of ticks elapsed in the current transition.
	//
	// Added: v1.26.10
	CurrentTransitionTicks uint32
	// TotalTransitionTicks is the total number of ticks for the transition.
	//
	// Added: v1.26.10
	TotalTransitionTicks uint32
	// EaseType is the easing function used for the transition. It is one of the EasingType constants.
	//
	// Added: v1.26.10
	EaseType int32
}

// Marshal encodes/decodes an EnvironmentAttributeData.
func (x *EnvironmentAttributeData) Marshal(r IO) {
	easingType := easingTypeToString(x.EaseType)
	r.String(&x.AttributeName)
	OptionalMarshaler(r, &x.FromAttribute)
	Single(r, &x.Attribute)
	OptionalMarshaler(r, &x.ToAttribute)
	r.Uint32(&x.CurrentTransitionTicks)
	r.Uint32(&x.TotalTransitionTicks)
	r.String(&easingType)
	easingTypeFromString(r, &x.EaseType, easingType)
}

// AttributeLayerSettings represents settings for an attribute layer.
//
// Added: v1.26.10
type AttributeLayerSettings struct {
	// Priority is the priority of the layer.
	//
	// Added: v1.26.10
	Priority int32
	// WeightType determines whether the weight is a float or string. It is one of the
	// AttributeLayerWeightType constants.
	//
	// Removed: v1.26.20.26
	WeightType uint32
	// FloatWeight is the numeric weight value for the layer.
	// When WeightType is present, this field is used when WeightType is AttributeLayerWeightTypeFloat.
	//
	// Changed: v1.26.20.26, encoded directly as a float and no longer selected through WeightType.
	FloatWeight float32
	// StringWeight is the weight if WeightType is AttributeLayerWeightTypeString.
	//
	// Removed: v1.26.20.26
	StringWeight string
	// Enabled indicates if the layer is enabled.
	//
	// Added: v1.26.10
	Enabled bool
	// TransitionsPaused indicates if transitions are paused for this layer.
	//
	// Added: v1.26.10
	TransitionsPaused bool
}

// Marshal encodes/decodes an AttributeLayerSettings.
func (x *AttributeLayerSettings) Marshal(r IO) {
	r.Int32(&x.Priority)
	if r.Protocol() >= Protocol1v26v20v26 {
		r.Float32(&x.FloatWeight)
	} else {
		r.Varuint32(&x.WeightType)
		switch x.WeightType {
		case AttributeLayerWeightTypeFloat:
			r.Float32(&x.FloatWeight)
		case AttributeLayerWeightTypeString:
			r.String(&x.StringWeight)
		default:
			r.UnknownEnumOption(x.WeightType, "attribute layer weight type")
		}
	}
	r.Bool(&x.Enabled)
	r.Bool(&x.TransitionsPaused)
}

// AttributeLayerData represents a complete attribute layer.
//
// Added: v1.26.10
type AttributeLayerData struct {
	// Name is the name of the attribute layer.
	//
	// Added: v1.26.10
	Name string
	// DimensionID is the dimension the layer applies to.
	//
	// Added: v1.26.10
	DimensionID int32
	// Settings is the layer's settings.
	//
	// Added: v1.26.10
	Settings AttributeLayerSettings
	// EnvironmentAttributes is the list of environment attributes in this layer.
	//
	// Added: v1.26.10
	EnvironmentAttributes []EnvironmentAttributeData
}

// Marshal encodes/decodes an AttributeLayerData.
func (x *AttributeLayerData) Marshal(r IO) {
	r.String(&x.Name)
	r.Varint32(&x.DimensionID)
	Single(r, &x.Settings)
	Slice(r, &x.EnvironmentAttributes)
}
