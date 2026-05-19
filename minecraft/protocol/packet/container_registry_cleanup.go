package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

// ContainerRegistryCleanup is sent by the server to trigger a client-side cleanup of the dynamic container
// registry.
//
// Added: v1.21.30
type ContainerRegistryCleanup struct {
	// RemovedContainers is a list of protocol.FullContainerName's that should be removed from the client-side
	// container registry.
	//
	// Added: v1.21.30
	RemovedContainers []protocol.FullContainerName
}

// ID ...
func (*ContainerRegistryCleanup) ID() uint32 {
	return IDContainerRegistryCleanup
}

func (pk *ContainerRegistryCleanup) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.RemovedContainers)
}
