package protocol

// Info represents a Minecraft protocol version identifier.
type Info int32

// Pool holds the Minecraft version string for each registered protocol version.
var Pool = make(map[Info]string)

// NewInfo registers a protocol version and returns it as an Info value.
func NewInfo(protocol int32, version string) Info {
	info := Info(protocol)
	Pool[info] = version
	return info
}

// ID returns the protocol version as an int32.
func (i Info) ID() int32 {
	return int32(i)
}

// Ver returns the Minecraft version string registered for the protocol version.
func (i Info) Ver() string {
	if name, ok := Pool[i]; ok {
		return name
	}
	return "unknown"
}

// Current Protocol and Version
var CurrentInfo = NewInfo(CurrentProtocol, CurrentVersion)

const (
	// CurrentProtocol is the current protocol version for the version below.
	CurrentProtocol = Protocol1v26v20
	// CurrentVersion is the current version of Minecraft as supported by the `packet` package.
	CurrentVersion = Version1v26v20
)

// Minecraft 1.26.20
// Info1v26v20 is the protocol info for Minecraft 1.26.20.
var Info1v26v20 = NewInfo(Protocol1v26v20, Version1v26v20)

const (
	// Protocol1v26v20 is the protocol version for Minecraft 1.26.20.
	Protocol1v26v20 int32 = 975
	// Version1v26v20 is the Minecraft version string for protocol 975.
	Version1v26v20 = "1.26.20"
)

// Minecraft 1.26.20.26
// Info1v26v20v26 is the protocol info for Minecraft 1.26.20.26.
var Info1v26v20v26 = NewInfo(Protocol1v26v20v26, Version1v26v20v26)

const (
	// Protocol1v26v20v26 is the protocol version for Minecraft 1.26.20.26.
	Protocol1v26v20v26 int32 = 974
	// Version1v26v20v26 is the Minecraft version string for protocol 974.
	Version1v26v20v26 = "1.26.20.26"
)

// Minecraft 1.26.10
// Info1v26v10 is the protocol info for Minecraft 1.26.10.
var Info1v26v10 = NewInfo(Protocol1v26v10, Version1v26v10)

const (
	// Protocol1v26v10 is the protocol version for Minecraft 1.26.10.
	Protocol1v26v10 int32 = 944
	// Version1v26v10 is the Minecraft version string for protocol 944.
	Version1v26v10 = "1.26.10"
)

// Minecraft 1.26.0
// Info1v26v0 is the protocol info for Minecraft 1.26.0.
var Info1v26v0 = NewInfo(Protocol1v26v0, Version1v26v0)

const (
	// Protocol1v26v0 is the protocol version for Minecraft 1.26.0.
	Protocol1v26v0 int32 = 924
	// Version1v26v0 is the Minecraft version string for protocol 924.
	Version1v26v0 = "1.26.0"
)

// Minecraft 1.21.130
// Info1v21v130 is the protocol info for Minecraft 1.21.130.
var Info1v21v130 = NewInfo(Protocol1v21v130, Version1v21v130)

const (
	// Protocol1v21v130 is the protocol version for Minecraft 1.21.130.
	Protocol1v21v130 int32 = 898
	// Version1v21v130 is the Minecraft version string for protocol 898.
	Version1v21v130 = "1.21.130"
)
