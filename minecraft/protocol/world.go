package protocol

const (
	GeneratorLegacy    = 0
	GeneratorOverworld = 1
	GeneratorFlat      = 2
	GeneratorNether    = 3
	GeneratorEnd       = 4
	GeneratorVoid      = 5
)

// DimensionDefinition contains information specifying dimension-specific properties, used for data-driven dimensions.
// These include the range (the height min/max), generator variant, and more.
//
// Added: v1.18.30
type DimensionDefinition struct {
	// Name specifies the name of the dimension.
	//
	// Added: v1.18.30
	Name string
	// Range is the height range of the dimension, where the first value is the minimum and the second is the maximum.
	//
	// Added: v1.18.30
	Range [2]int32
	// Generator is the variant of generator that exists in the provided dimension. These can be one of the constants
	// defined above. If this is set to GeneratorLegacy, the legacy horizontal world limits will be enforced.
	//
	// Added: v1.18.30
	Generator int32
	// DimensionType is the numeric identifier of the dimension. This cannot override a vanilla dimension (0-2), but
	// custom dimensions should start from 1000 like vanilla.
	//
	// Added: v1.18.30
	DimensionType int32
}

// Marshal encodes/decodes a DimensionDefinition.
func (x *DimensionDefinition) Marshal(r IO) {
	r.String(&x.Name)
	r.Varint32(&x.Range[0])
	r.Varint32(&x.Range[1])
	r.Varint32(&x.Generator)
	r.Varint32(&x.DimensionType)
}

// GenerationFeature represents a world generation feature, used when encoding the FeatureRegistry to the client.
//
// Added: v1.19.20
type GenerationFeature struct {
	// Name is the name of the feature.
	//
	// Added: v1.19.20
	Name string
	// JSON is the encoded JSON data instructing the client on how to generate the feature.
	//
	// Added: v1.19.20
	JSON []byte
}

// Marshal encodes/decodes a GenerationFeature.
func (x *GenerationFeature) Marshal(r IO) {
	r.String(&x.Name)
	r.ByteSlice(&x.JSON)
}
