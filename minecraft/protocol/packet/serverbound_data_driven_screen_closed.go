package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	DataDrivenScreenCloseReasonProgrammaticClose    = "programmatic_close"
	DataDrivenScreenCloseReasonProgrammaticCloseAll = "programmatic_close_all"
	DataDrivenScreenCloseReasonClientCanceled       = "client_canceled"
	DataDrivenScreenCloseReasonUserBusy             = "user_busy"
	DataDrivenScreenCloseReasonInvalidForm          = "invalid_form"
)

// ServerBoundDataDrivenScreenClosed is sent by the client when a data-driven UI screen is closed.
//
// Added: v1.26.10
type ServerBoundDataDrivenScreenClosed struct {
	// FormID is the unique instance ID of the form that was closed.
	//
	// Added: v1.26.10
	FormID int32
	// CloseReason is the reason the screen was closed. It is one of the DataDrivenScreenCloseReason constants.
	//
	// Added: v1.26.10
	CloseReason string
}

// ID ...
func (*ServerBoundDataDrivenScreenClosed) ID() uint32 {
	return IDServerBoundDataDrivenScreenClosed
}

func (pk *ServerBoundDataDrivenScreenClosed) Marshal(io protocol.IO) {
	io.Int32(&pk.FormID)
	io.String(&pk.CloseReason)
}
