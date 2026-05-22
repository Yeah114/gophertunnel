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
	// 1.26.x
	Protocol1v26v20:    Version1v26v20,
	Protocol1v26v20v26: Version1v26v20v26,
	Protocol1v26v10:    Version1v26v10,
	Protocol1v26v0:     Version1v26v0,

	// 1.21.x
	Protocol1v21v130:    Version1v21v130,
	Protocol1v21v130v28: Version1v21v130v28,
	Protocol1v21v124:    Version1v21v124,
	Protocol1v21v120:    Version1v21v120,
	Protocol1v21v110:    Version1v21v110,
	Protocol1v21v110v26: Version1v21v110v26,
	Protocol1v21v100:    Version1v21v100,
	Protocol1v21v93:     Version1v21v93,
	Protocol1v21v90:     Version1v21v90,
	Protocol1v21v80:     Version1v21v80,
	Protocol1v21v70:     Version1v21v70,
	Protocol1v21v70v24:  Version1v21v70v24,
	Protocol1v21v60:     Version1v21v60,
	Protocol1v21v50:     Version1v21v50,
	Protocol1v21v50v26:  Version1v21v50v26,
	Protocol1v21v40:     Version1v21v40,
	Protocol1v21v30:     Version1v21v30,
	Protocol1v21v20:     Version1v21v20,
	Protocol1v21v2:      Version1v21v2,
	Protocol1v21v0:      Version1v21v0,

	// 1.20.x
	Protocol1v20v80:    Version1v20v80,
	Protocol1v20v70:    Version1v20v70,
	Protocol1v20v60:    Version1v20v60,
	Protocol1v20v50:    Version1v20v50,
	Protocol1v20v40:    Version1v20v40,
	Protocol1v20v30:    Version1v20v30,
	Protocol1v20v30v24: Version1v20v30v24,
	Protocol1v20v10:    Version1v20v10,
	Protocol1v20v0:     Version1v20v0,
	Protocol1v20v0v23:  Version1v20v0v23,

	// 1.19.x
	Protocol1v19v80:    Version1v19v80,
	Protocol1v19v70:    Version1v19v70,
	Protocol1v19v70v24: Version1v19v70v24,
	Protocol1v19v63:    Version1v19v63,
	Protocol1v19v60:    Version1v19v60,
	Protocol1v19v50:    Version1v19v50,
	Protocol1v19v40:    Version1v19v40,
	Protocol1v19v30:    Version1v19v30,
	Protocol1v19v21:    Version1v19v21,
	Protocol1v19v20:    Version1v19v20,
	Protocol1v19v10:    Version1v19v10,
	Protocol1v19v0:     Version1v19v0,
	Protocol1v19v0v29:  Version1v19v0v29,

	// 1.18.x
	Protocol1v18v30:    Version1v18v30,
	Protocol1v18v10:    Version1v18v10,
	Protocol1v18v10v26: Version1v18v10v26,
	Protocol1v18v0:     Version1v18v0,

	// 1.17.x
	Protocol1v17v40:    Version1v17v40,
	Protocol1v17v30:    Version1v17v30,
	Protocol1v17v20v20: Version1v17v20v20,
	Protocol1v17v10:    Version1v17v10,
	Protocol1v17v0:     Version1v17v0,

	// 1.16.x
	Protocol1v16v230v54: Version1v16v230v54,
	Protocol1v16v230:    Version1v16v230,
	Protocol1v16v230v50: Version1v16v230v50,
	Protocol1v16v220:    Version1v16v220,
	Protocol1v16v210:    Version1v16v210,
	Protocol1v16v210v50: Version1v16v210v50,
	Protocol1v16v200:    Version1v16v200,
	Protocol1v16v200v51: Version1v16v200v51,
	Protocol1v16v100:    Version1v16v100,
	Protocol1v16v20:     Version1v16v20,
	Protocol1v16v0:      Version1v16v0,

	// 1.14.x
	Protocol1v14v60: Version1v14v60,
	Protocol1v14v0:  Version1v14v0,

	// 1.13.x
	Protocol1v13v0: Version1v13v0,

	// 1.12.x
	Protocol1v12v0: Version1v12v0,

	// 1.11.x
	Protocol1v11v0: Version1v11v0,

	// 1.10.x
	Protocol1v10v0: Version1v10v0,

	// 1.9.x
	Protocol1v9v0: Version1v9v0,

	// 1.8.x
	Protocol1v8v0: Version1v8v0,

	// 1.7.x
	Protocol1v7v0: Version1v7v0,

	// 1.6.x
	Protocol1v6v0:   Version1v6v0,
	Protocol1v6v0v5: Version1v6v0v5,

	// 1.5.x
	Protocol1v5v0: Version1v5v0,

	// 1.4.x
	Protocol1v4v0: Version1v4v0,

	// 1.2.x
	Protocol1v2v13v11: Version1v2v13v11,
	Protocol1v2v13:    Version1v2v13,
	Protocol1v2v10:    Version1v2v10,
	Protocol1v2v7:     Version1v2v7,
	Protocol1v2v6:     Version1v2v6,
	Protocol1v2v5v11:  Version1v2v5v11,
	Protocol1v2v5:     Version1v2v5,
	Protocol1v2v0:     Version1v2v0,

	// 1.1.x
	Protocol1v1v0: Version1v1v0,
}

// VersionToProtocol is a map of Minecraft version strings to protocol versions. It is used to lookup the protocol version associated with a Minecraft version string. It is not used for any other purpose, and may be modified or removed without warning.
var VersionToProtocol = map[string]int32{
	// 1.26.x
	Version1v26v20:    Protocol1v26v20,
	Version1v26v20v26: Protocol1v26v20v26,
	Version1v26v10:    Protocol1v26v10,
	Version1v26v0:     Protocol1v26v0,

	// 1.21.x
	Version1v21v130:    Protocol1v21v130,
	Version1v21v130v28: Protocol1v21v130v28,
	Version1v21v124:    Protocol1v21v124,
	Version1v21v120:    Protocol1v21v120,
	Version1v21v110:    Protocol1v21v110,
	Version1v21v110v26: Protocol1v21v110v26,
	Version1v21v100:    Protocol1v21v100,
	Version1v21v93:     Protocol1v21v93,
	Version1v21v90:     Protocol1v21v90,
	Version1v21v80:     Protocol1v21v80,
	Version1v21v70:     Protocol1v21v70,
	Version1v21v70v24:  Protocol1v21v70v24,
	Version1v21v60:     Protocol1v21v60,
	Version1v21v50:     Protocol1v21v50,
	Version1v21v50v26:  Protocol1v21v50v26,
	Version1v21v40:     Protocol1v21v40,
	Version1v21v30:     Protocol1v21v30,
	Version1v21v20:     Protocol1v21v20,
	Version1v21v2:      Protocol1v21v2,
	Version1v21v0:      Protocol1v21v0,

	// 1.20.x
	Version1v20v80:    Protocol1v20v80,
	Version1v20v70:    Protocol1v20v70,
	Version1v20v60:    Protocol1v20v60,
	Version1v20v50:    Protocol1v20v50,
	Version1v20v40:    Protocol1v20v40,
	Version1v20v30:    Protocol1v20v30,
	Version1v20v30v24: Protocol1v20v30v24,
	Version1v20v10:    Protocol1v20v10,
	Version1v20v0:     Protocol1v20v0,
	Version1v20v0v23:  Protocol1v20v0v23,

	// 1.19.x
	Version1v19v80:    Protocol1v19v80,
	Version1v19v70:    Protocol1v19v70,
	Version1v19v70v24: Protocol1v19v70v24,
	Version1v19v63:    Protocol1v19v63,
	Version1v19v60:    Protocol1v19v60,
	Version1v19v50:    Protocol1v19v50,
	Version1v19v40:    Protocol1v19v40,
	Version1v19v30:    Protocol1v19v30,
	Version1v19v21:    Protocol1v19v21,
	Version1v19v20:    Protocol1v19v20,
	Version1v19v10:    Protocol1v19v10,
	Version1v19v0:     Protocol1v19v0,
	Version1v19v0v29:  Protocol1v19v0v29,

	// 1.18.x
	Version1v18v30:    Protocol1v18v30,
	Version1v18v10:    Protocol1v18v10,
	Version1v18v10v26: Protocol1v18v10v26,
	Version1v18v0:     Protocol1v18v0,

	// 1.17.x
	Version1v17v40:    Protocol1v17v40,
	Version1v17v30:    Protocol1v17v30,
	Version1v17v20v20: Protocol1v17v20v20,
	Version1v17v10:    Protocol1v17v10,
	Version1v17v0:     Protocol1v17v0,

	// 1.16.x
	Version1v16v230v54: Protocol1v16v230v54,
	Version1v16v230:    Protocol1v16v230,
	Version1v16v230v50: Protocol1v16v230v50,
	Version1v16v220:    Protocol1v16v220,
	Version1v16v210:    Protocol1v16v210,
	Version1v16v210v50: Protocol1v16v210v50,
	Version1v16v200:    Protocol1v16v200,
	Version1v16v200v51: Protocol1v16v200v51,
	Version1v16v100:    Protocol1v16v100,
	Version1v16v20:     Protocol1v16v20,
	Version1v16v0:      Protocol1v16v0,

	// 1.14.x
	Version1v14v60: Protocol1v14v60,
	Version1v14v0:  Protocol1v14v0,

	// 1.13.x
	Version1v13v0: Protocol1v13v0,

	// 1.12.x
	Version1v12v0: Protocol1v12v0,

	// 1.11.x
	Version1v11v0: Protocol1v11v0,

	// 1.10.x
	Version1v10v0: Protocol1v10v0,

	// 1.9.x
	Version1v9v0: Protocol1v9v0,

	// 1.8.x
	Version1v8v0: Protocol1v8v0,

	// 1.7.x
	Version1v7v0: Protocol1v7v0,

	// 1.6.x
	Version1v6v0:   Protocol1v6v0,
	Version1v6v0v5: Protocol1v6v0v5,

	// 1.5.x
	Version1v5v0: Protocol1v5v0,

	// 1.4.x
	Version1v4v0: Protocol1v4v0,

	// 1.2.x
	Version1v2v13v11: Protocol1v2v13v11,
	Version1v2v13:    Protocol1v2v13,
	Version1v2v10:    Protocol1v2v10,
	Version1v2v7:     Protocol1v2v7,
	Version1v2v6:     Protocol1v2v6,
	Version1v2v5v11:  Protocol1v2v5v11,
	Version1v2v5:     Protocol1v2v5,
	Version1v2v0:     Protocol1v2v0,

	// 1.1.x
	Version1v1v0: Protocol1v1v0,
}
