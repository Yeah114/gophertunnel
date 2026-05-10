package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ServerStoreInfo is sent by the server to provide the client with a store entry point. Like the ShowStoreOffer packet,
// this only has an effect on partnered servers. This packet was added in v1.26.20.
type ServerStoreInfo struct {
	// StoreInfo is the store info to set, or nothing to fall back to the default.
	StoreInfo protocol.Optional[protocol.StoreEntryPointInfo]
}

// ID ...
func (*ServerStoreInfo) ID() uint32 {
	return IDServerStoreInfo
}

func (pk *ServerStoreInfo) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.StoreInfo)
}
