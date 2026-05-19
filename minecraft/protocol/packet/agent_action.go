package packet

import (
	"github.com/Yeah114/gophertunnel/minecraft/protocol"
)

const (
	AgentActionTypeAttack = iota + 1
	AgentActionTypeCollect
	AgentActionTypeDestroy
	AgentActionTypeDetectRedstone
	AgentActionTypeDetectObstacle
	AgentActionTypeDrop
	AgentActionTypeDropAll
	AgentActionTypeInspect
	AgentActionTypeInspectData
	AgentActionTypeInspectItemCount
	AgentActionTypeInspectItemDetail
	AgentActionTypeInspectItemSpace
	AgentActionTypeInteract
	AgentActionTypeMove
	AgentActionTypePlaceBlock
	AgentActionTypeTill
	AgentActionTypeTransferItemTo
	AgentActionTypeTurn
)

// AgentAction is an Education Edition packet sent from the server to the client to return a response to a
// previously requested action.
//
// Added: v1.18.30
type AgentAction struct {
	// Identifier is a JSON identifier referenced in the initial action.
	//
	// Added: v1.18.30
	Identifier string
	// Action represents the action type that was requested. It is one of the constants defined above.
	//
	// Added: v1.18.30
	Action int32
	// Response is a JSON string containing the response to the action.
	//
	// Added: v1.18.30
	Response []byte
}

// ID ...
func (*AgentAction) ID() uint32 {
	return IDAgentAction
}

func (pk *AgentAction) Marshal(io protocol.IO) {
	io.String(&pk.Identifier)
	io.Int32(&pk.Action)
	io.ByteSlice(&pk.Response)
}
