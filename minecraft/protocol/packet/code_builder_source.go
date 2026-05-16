package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CodeBuilderOperationNone = iota
	CodeBuilderOperationGet
	CodeBuilderOperationSet
	CodeBuilderOperationReset
)

const (
	CodeBuilderCategoryNone = iota
	CodeBuilderCategoryStatus
	CodeBuilderCategoryInstantiation
)

const (
	CodeBuilderStatusNone = iota
	CodeBuilderStatusNotStarted
	CodeBuilderStatusInProgress
	CodeBuilderStatusPaused
	CodeBuilderStatusError
	CodeBuilderStatusSucceeded
)

// CodeBuilderSource is an Education Edition packet sent by the client to the server to run an operation with a
// code builder.
//
// Added: v1.18.10
// Changed: v1.21.0, the Value payload field was replaced with the CodeStatus field.
type CodeBuilderSource struct {
	// Operation is used to distinguish the operation performed. It is always one of the constants listed above.
	//
	// Added: v1.18.10
	Operation byte
	// Category is used to distinguish the category of the operation performed. It is always one of the constants
	// listed above.
	//
	// Added: v1.18.10
	Category byte
	// CodeStatus is the status of the code builder. It is always one of the constants listed above.
	//
	// Added: v1.21.0
	CodeStatus byte
}

// ID ...
func (pk *CodeBuilderSource) ID() uint32 {
	return IDCodeBuilderSource
}

func (pk *CodeBuilderSource) Marshal(io protocol.IO) {
	io.Uint8(&pk.Operation)
	io.Uint8(&pk.Category)
	io.Uint8(&pk.CodeStatus)
}
