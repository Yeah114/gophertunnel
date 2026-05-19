package protocol

// NewInfoByProtocol creates a new Info with the given protocol version.
// The Minecraft version string registered for the protocol version is looked up in ProtocolToVersion,
// and if it is not found, the CurrentVersion is used as a fallback.
func NewInfoByProtocol(protocol int32) Info {
	if version, ok := ProtocolToVersion[protocol]; ok {
		return NewInfo(protocol, version)
	}
	return CurrentInfo
}

// NewInfoByVersion creates a new Info with the given Minecraft version string.
// The protocol version registered for the Minecraft version string is looked up in VersionToProtocol,
// and if it is not found, the CurrentProtocol is used as a fallback.
func NewInfoByVersion(version string) Info {
	if protocol, ok := VersionToProtocol[version]; ok {
		return NewInfo(protocol, version)
	}
	return CurrentInfo
}

// ProtocolToVersion is a map of protocol versions to Minecraft version strings. It is used to lookup the Minecraft version string associated with a protocol version. It is not used for any other purpose, and may be modified or removed without warning.
var ProtocolToVersion = map[int32]string{
	Protocol1v26v20:    Version1v26v20,
	Protocol1v26v20v26: Version1v26v20v26,
	Protocol1v26v10:    Version1v26v10,
	Protocol1v26v0:     Version1v26v0,
	Protocol1v21v130:   Version1v21v130,
}

// VersionToProtocol is a map of Minecraft version strings to protocol versions. It is used to lookup the protocol version associated with a Minecraft version string. It is not used for any other purpose, and may be modified or removed without warning.
var VersionToProtocol = map[string]int32{
	Version1v26v20:    Protocol1v26v20,
	Version1v26v20v26: Protocol1v26v20v26,
	Version1v26v10:    Protocol1v26v10,
	Version1v26v0:     Protocol1v26v0,
	Version1v21v130:   Protocol1v21v130,
}
