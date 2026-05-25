package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

// TickSync is sent to synchronise client and server tick timestamps.
//
// Added: v1.11.1
type TickSync struct {
	// RequestTimestamp is the timestamp sent by the requester.
	RequestTimestamp int64
	// ResponseTimestamp is the timestamp sent in response.
	ResponseTimestamp int64
}

// ID ...
func (*TickSync) ID() uint32 {
	return IDTickSync
}

func (pk *TickSync) Marshal(io protocol.IO) {
	io.Int64(&pk.RequestTimestamp)
	io.Int64(&pk.ResponseTimestamp)
}
