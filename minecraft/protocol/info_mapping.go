package protocol

// GetProfile 按协议版本查找官方 Profile。
//
// GetProfile looks up an official profile by its protocol version.
func GetProfile(protocolVersion int32) (Profile, bool) {
	for _, profile := range profiles {
		if profile.Protocol == protocolVersion {
			return profile, true
		}
	}
	return Profile{}, false
}

// GetProtocol 按 Minecraft 版本字符串查找官方协议版本。
//
// GetProtocol looks up an official protocol version by its Minecraft version string.
func GetProtocol(version string) (int32, bool) {
	for _, profile := range profiles {
		if profile.Version == version {
			return profile.Protocol, true
		}
	}
	return 0, false
}

// GetVersion 按协议版本查找官方 Minecraft 版本字符串。
//
// GetVersion looks up an official Minecraft version string by its protocol version.
func GetVersion(protocolVersion int32) (string, bool) {
	profile, ok := GetProfile(protocolVersion)
	if !ok {
		return "", false
	}
	return profile.Version, true
}

// Profiles 返回所有协议 Profile。
//
// Profiles returns all protocol profiles.
func Profiles() []Profile {
	return append([]Profile(nil), profiles...)
}

var profiles = []Profile{
	// 1.26.x 版本。
	//
	// 1.26.x versions.
	Profile1v26v20,
	Profile1v26v20v26,
	Profile1v26v10,
	Profile1v26v0,

	// 1.21.x 版本。
	//
	// 1.21.x versions.
	Profile1v21v130,
	Profile1v21v130v28,
	Profile1v21v124,
	Profile1v21v120,
	Profile1v21v110,
	Profile1v21v110v26,
	Profile1v21v100,
	Profile1v21v93,
	Profile1v21v90,
	Profile1v21v80,
	Profile1v21v70,
	Profile1v21v70v24,
	Profile1v21v60,
	Profile1v21v50,
	Profile1v21v50v26,
	Profile1v21v40,
	Profile1v21v30,
	Profile1v21v20,
	Profile1v21v2,
	Profile1v21v0,

	// 1.20.x 版本。
	//
	// 1.20.x versions.
	Profile1v20v80,
	Profile1v20v70,
	Profile1v20v60,
	Profile1v20v50,
	Profile1v20v40,
	Profile1v20v30,
	Profile1v20v30v24,
	Profile1v20v10,
	Profile1v20v0,
	Profile1v20v0v23,

	// 1.19.x 版本。
	//
	// 1.19.x versions.
	Profile1v19v80,
	Profile1v19v70,
	Profile1v19v70v24,
	Profile1v19v63,
	Profile1v19v60,
	Profile1v19v50,
	Profile1v19v40,
	Profile1v19v30,
	Profile1v19v21,
	Profile1v19v20,
	Profile1v19v10,
	Profile1v19v0,
	Profile1v19v0v29,

	// 1.18.x 版本。
	//
	// 1.18.x versions.
	Profile1v18v30,
	Profile1v18v10,
	Profile1v18v10v26,
	Profile1v18v0,

	// 1.17.x 版本。
	//
	// 1.17.x versions.
	Profile1v17v40,
	Profile1v17v30,
	Profile1v17v20v20,
	Profile1v17v10,
	Profile1v17v0,

	// 1.16.x 版本。
	//
	// 1.16.x versions.
	Profile1v16v230v54,
	Profile1v16v230,
	Profile1v16v230v50,
	Profile1v16v220,
	Profile1v16v210,
	Profile1v16v210v50,
	Profile1v16v200,
	Profile1v16v200v51,
	Profile1v16v100,
	Profile1v16v20,
	Profile1v16v0,

	// 1.14.x 版本。
	//
	// 1.14.x versions.
	Profile1v14v60,
	Profile1v14v0,

	// 1.13.x 版本。
	//
	// 1.13.x versions.
	Profile1v13v0,

	// 1.12.x 版本。
	//
	// 1.12.x versions.
	Profile1v12v0,

	// 1.11.x 版本。
	//
	// 1.11.x versions.
	Profile1v11v0,

	// 1.10.x 版本。
	//
	// 1.10.x versions.
	Profile1v10v0,

	// 1.9.x 版本。
	//
	// 1.9.x versions.
	Profile1v9v0,

	// 1.8.x 版本。
	//
	// 1.8.x versions.
	Profile1v8v0,

	// 1.7.x 版本。
	//
	// 1.7.x versions.
	Profile1v7v0,

	// 1.6.x 版本。
	//
	// 1.6.x versions.
	Profile1v6v0,
	Profile1v6v0v5,

	// 1.5.x 版本。
	//
	// 1.5.x versions.
	Profile1v5v0,

	// 1.4.x 版本。
	//
	// 1.4.x versions.
	Profile1v4v0,

	// 1.2.x 版本。
	//
	// 1.2.x versions.
	Profile1v2v13v11,
	Profile1v2v13,
	Profile1v2v10,
	Profile1v2v7,
	Profile1v2v6,
	Profile1v2v5v11,
	Profile1v2v5,
	Profile1v2v0,

	// 1.1.x 版本。
	//
	// 1.1.x versions.
	Profile1v1v0,
}
