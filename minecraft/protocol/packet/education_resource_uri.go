package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// EducationResourceURI is a packet that transmits education resource settings to all clients.
//
// Added: v1.17.30
type EducationResourceURI struct {
	// Resource is the resource that is being referenced.
	//
	// Added: v1.17.30
	Resource protocol.EducationSharedResourceURI
}

// ID ...
func (*EducationResourceURI) ID() uint32 {
	return IDEducationResourceURI
}

func (pk *EducationResourceURI) Marshal(io protocol.IO) {
	protocol.Single(io, &pk.Resource)
}
