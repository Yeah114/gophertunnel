package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ServerStoreInfo is sent by the server to provide the client with a store entry point. Like the ShowStoreOffer packet,
// this only has an effect on partnered servers.
//
// Added: v1.26.20
type ServerStoreInfo struct {
	// StoreInfo is the store info to set, or nothing to fall back to the default.
	//
	// Added: v1.26.20
	StoreInfo protocol.Optional[protocol.StoreEntryPointInfo]
}

// ID ...
func (*ServerStoreInfo) ID() uint32 {
	return IDServerStoreInfo
}

func (pk *ServerStoreInfo) Marshal(io protocol.IO) {
	protocol.OptionalMarshaler(io, &pk.StoreInfo)
}
