package protocol

// ExperimentData holds data on an experiment that is either enabled or disabled.
//
// Added: v1.16.100
type ExperimentData struct {
	// Name is the name of the experiment.
	//
	// Added: v1.16.100
	Name string
	// Enabled specifies if the experiment is enabled. Vanilla typically always sets this to true for any
	// experiments sent.
	//
	// Added: v1.16.100
	Enabled bool
}

// Marshal encodes/decodes an ExperimentData.
func (x *ExperimentData) Marshal(r IO) {
	r.String(&x.Name)
	r.Bool(&x.Enabled)
}
