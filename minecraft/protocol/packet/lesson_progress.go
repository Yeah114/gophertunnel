package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	LessonActionStart = iota
	LessonActionComplete
	LessonActionRestart
)

// LessonProgress is a packet sent by the server to the client to inform the client of updated progress on a lesson.
// This packet only functions on the Minecraft: Education Edition version of the game.
//
// Added: v1.18.30
type LessonProgress struct {
	// Identifier is the identifier of the lesson that is being progressed.
	//
	// Added: v1.18.30
	Identifier string
	// Action is the action the client should perform to show progress. This is one of the constants defined above.
	//
	// Added: v1.18.30
	Action int32
	// Score is the score the client should use when displaying the progress.
	//
	// Added: v1.18.30
	Score int32
}

// ID ...
func (*LessonProgress) ID() uint32 {
	return IDLessonProgress
}

func (pk *LessonProgress) Marshal(io protocol.IO) {
	io.Varint32(&pk.Action)
	io.Varint32(&pk.Score)
	io.String(&pk.Identifier)
}
