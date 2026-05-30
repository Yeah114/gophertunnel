package packet

import "github.com/Yeah114/gophertunnel/minecraft/protocol"

const (
	VideoStreamActionOpen byte = iota
	VideoStreamActionClose
)

// VideoStreamConnect is sent by the server to open or close a video stream connection.
type VideoStreamConnect struct {
	// Address is the stream address the client should connect to.
	Address string
	// ScreenshotFrequency is the interval at which screenshots are sent.
	ScreenshotFrequency float32
	// Action is either VideoStreamActionOpen or VideoStreamActionClose.
	Action byte
	// Width and Height specify the stream dimensions.
	Width, Height int32
}

// ID ...
func (*VideoStreamConnect) ID() uint32 {
	return IDVideoStreamConnect
}

func (pk *VideoStreamConnect) Marshal(io protocol.IO) {
	io.String(&pk.Address)
	io.Float32(&pk.ScreenshotFrequency)
	io.Uint8(&pk.Action)
	if io.Protocol() >= protocol.Protocol1v12v0 {
		io.Int32(&pk.Width)
		io.Int32(&pk.Height)
	}
}
