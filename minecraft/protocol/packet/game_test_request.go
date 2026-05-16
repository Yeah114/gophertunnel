package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	GameTestRequestRotation0 = iota
	GameTestRequestRotation90
	GameTestRequestRotation180
	GameTestRequestRotation270
	GameTestRequestRotation360
)

// GameTestRequest requests that the server run a game test instance with the parameters provided.
//
// Added: v1.19.40
type GameTestRequest struct {
	// Name represents the name of the test.
	//
	// Added: v1.19.40
	Name string
	// Rotation represents the rotation of the test. It is one of the constants above.
	//
	// Added: v1.19.40
	Rotation uint8
	// Repetitions represents the amount of times the test will be run.
	//
	// Added: v1.19.40
	Repetitions int32
	// Position is the position at which the test will be performed.
	//
	// Added: v1.19.40
	Position protocol.BlockPos
	// StopOnError indicates whether the test should immediately stop when an error is encountered.
	//
	// Added: v1.19.40
	StopOnError bool
	// TestsPerRow is the amount of tests that should be placed on each row when building repeated test instances.
	//
	// Added: v1.19.40
	TestsPerRow int32
	// MaxTestsPerBatch is the maximum amount of test instances that may be processed in a single batch.
	//
	// Added: v1.19.40
	MaxTestsPerBatch int32
}

// ID ...
func (pk *GameTestRequest) ID() uint32 {
	return IDGameTestRequest
}

func (pk *GameTestRequest) Marshal(io protocol.IO) {
	io.Varint32(&pk.MaxTestsPerBatch)
	io.Varint32(&pk.Repetitions)
	io.Uint8(&pk.Rotation)
	io.Bool(&pk.StopOnError)
	io.BlockPos(&pk.Position)
	io.Varint32(&pk.TestsPerRow)
	io.String(&pk.Name)
}
