package protocol

import (
	"strconv"
	"strings"
)

// Profile 表示一个 Minecraft 协议版本配置。
//
// Profile represents a Minecraft protocol version profile.
type Profile struct {
	Protocol int32
	Version  string
}

// NewProfile 创建协议版本配置。
//
// NewProfile creates a protocol profile.
func NewProfile(protocol int32, version string) Profile {
	return Profile{Protocol: protocol, Version: version}
}

// ID 返回 int32 形式的协议版本号。
//
// ID returns the protocol version as an int32.
func (p Profile) ID() int32 {
	return p.Protocol
}

// Ver 返回该协议配置对应的 Minecraft 版本字符串。
//
// Ver returns the Minecraft version string for the protocol profile.
func (p Profile) Ver() string {
	return p.Version
}

// BlockStateVersion 将该配置的版本号转换为方块状态版本 ID。
//
// BlockStateVersion returns the profile version as a block state version ID.
func (p Profile) BlockStateVersion() int32 {
	v := make([]string, 4)
	copy(v, strings.Split(p.Ver(), "."))
	major, _ := strconv.ParseInt(v[0], 10, 32)
	minor, _ := strconv.ParseInt(v[1], 10, 32)
	patch, _ := strconv.ParseInt(v[2], 10, 32)
	revision, _ := strconv.ParseInt(v[3], 10, 32)
	return int32(major)<<24 | int32(minor)<<16 | int32(patch)<<8 | int32(revision)
}

// CurrentProfile 是当前默认使用的协议配置。
//
// CurrentProfile is the currently default protocol profile.
var CurrentProfile = NewProfile(CurrentProtocol, CurrentVersion)

const (
	// CurrentProtocol 是下方 CurrentVersion 对应的当前协议版本。
	//
	// CurrentProtocol is the current protocol version for the version below.
	CurrentProtocol = Protocol1v26v20
	// CurrentVersion 是 `packet` 包当前支持的 Minecraft 版本。
	//
	// CurrentVersion is the current version of Minecraft as supported by the `packet` package.
	CurrentVersion = Version1v26v20
)

// Profile1v26v20 是 Minecraft 1.26.20 的协议配置。
//
// Profile1v26v20 is the protocol profile for Minecraft 1.26.20.
var Profile1v26v20 = NewProfile(Protocol1v26v20, Version1v26v20)

const (
	// Protocol1v26v20 是 Minecraft 1.26.20 的协议版本。
	//
	// Protocol1v26v20 is the protocol version for Minecraft 1.26.20.
	Protocol1v26v20 int32 = 975
	// Version1v26v20 是协议 975 对应的 Minecraft 版本字符串。
	//
	// Version1v26v20 is the Minecraft version string for protocol 975.
	Version1v26v20 = "1.26.20"
)

// Profile1v26v20v26 是 Minecraft 1.26.20.26 的协议配置。
//
// Profile1v26v20v26 is the protocol profile for Minecraft 1.26.20.26.
var Profile1v26v20v26 = NewProfile(Protocol1v26v20v26, Version1v26v20v26)

const (
	// Protocol1v26v20v26 是 Minecraft 1.26.20.26 的协议版本。
	//
	// Protocol1v26v20v26 is the protocol version for Minecraft 1.26.20.26.
	Protocol1v26v20v26 int32 = 974
	// Version1v26v20v26 是协议 974 对应的 Minecraft 版本字符串。
	//
	// Version1v26v20v26 is the Minecraft version string for protocol 974.
	Version1v26v20v26 = "1.26.20.26"
)

// Profile1v26v10 是 Minecraft 1.26.10 的协议配置。
//
// Profile1v26v10 is the protocol profile for Minecraft 1.26.10.
var Profile1v26v10 = NewProfile(Protocol1v26v10, Version1v26v10)

const (
	// Protocol1v26v10 是 Minecraft 1.26.10 的协议版本。
	//
	// Protocol1v26v10 is the protocol version for Minecraft 1.26.10.
	Protocol1v26v10 int32 = 944
	// Version1v26v10 是协议 944 对应的 Minecraft 版本字符串。
	//
	// Version1v26v10 is the Minecraft version string for protocol 944.
	Version1v26v10 = "1.26.10"
)

// Profile1v26v0 是 Minecraft 1.26.0 的协议配置。
//
// Profile1v26v0 is the protocol profile for Minecraft 1.26.0.
var Profile1v26v0 = NewProfile(Protocol1v26v0, Version1v26v0)

const (
	// Protocol1v26v0 是 Minecraft 1.26.0 的协议版本。
	//
	// Protocol1v26v0 is the protocol version for Minecraft 1.26.0.
	Protocol1v26v0 int32 = 924
	// Version1v26v0 是协议 924 对应的 Minecraft 版本字符串。
	//
	// Version1v26v0 is the Minecraft version string for protocol 924.
	Version1v26v0 = "1.26.0"
)

// Profile1v21v130 是 Minecraft 1.21.130 的协议配置。
//
// Profile1v21v130 is the protocol profile for Minecraft 1.21.130.
var Profile1v21v130 = NewProfile(Protocol1v21v130, Version1v21v130)

const (
	// Protocol1v21v130 是 Minecraft 1.21.130 的协议版本。
	//
	// Protocol1v21v130 is the protocol version for Minecraft 1.21.130.
	Protocol1v21v130 int32 = 898
	// Version1v21v130 是协议 898 对应的 Minecraft 版本字符串。
	//
	// Version1v21v130 is the Minecraft version string for protocol 898.
	Version1v21v130 = "1.21.130"
)

// Profile1v21v130v28 是 Minecraft 1.21.130.28 的协议配置。
//
// Profile1v21v130v28 is the protocol profile for Minecraft 1.21.130.28.
var Profile1v21v130v28 = NewProfile(Protocol1v21v130v28, Version1v21v130v28)

const (
	// Protocol1v21v130v28 是 Minecraft 1.21.130.28 的协议版本。
	//
	// Protocol1v21v130v28 is the protocol version for Minecraft 1.21.130.28.
	Protocol1v21v130v28 int32 = 897
	// Version1v21v130v28 是协议 897 对应的 Minecraft 版本字符串。
	//
	// Version1v21v130v28 is the Minecraft version string for protocol 897.
	Version1v21v130v28 = "1.21.130.28"
)

// Profile1v21v124 是 Minecraft 1.21.124 的协议配置。
//
// Profile1v21v124 is the protocol profile for Minecraft 1.21.124.
var Profile1v21v124 = NewProfile(Protocol1v21v124, Version1v21v124)

const (
	// Protocol1v21v124 是 Minecraft 1.21.124 的协议版本。
	//
	// Protocol1v21v124 is the protocol version for Minecraft 1.21.124.
	Protocol1v21v124 int32 = 860
	// Version1v21v124 是协议 860 对应的 Minecraft 版本字符串。
	//
	// Version1v21v124 is the Minecraft version string for protocol 860.
	Version1v21v124 = "1.21.124"
)

// Profile1v21v120 是 Minecraft 1.21.120 的协议配置。
//
// Profile1v21v120 is the protocol profile for Minecraft 1.21.120.
var Profile1v21v120 = NewProfile(Protocol1v21v120, Version1v21v120)

const (
	// Protocol1v21v120 是 Minecraft 1.21.120 的协议版本。
	//
	// Protocol1v21v120 is the protocol version for Minecraft 1.21.120.
	Protocol1v21v120 int32 = 859
	// Version1v21v120 是协议 859 对应的 Minecraft 版本字符串。
	//
	// Version1v21v120 is the Minecraft version string for protocol 859.
	Version1v21v120 = "1.21.120"
)

// Profile1v21v110 是 Minecraft 1.21.110 的协议配置。
//
// Profile1v21v110 is the protocol profile for Minecraft 1.21.110.
var Profile1v21v110 = NewProfile(Protocol1v21v110, Version1v21v110)

const (
	// Protocol1v21v110 是 Minecraft 1.21.110 的协议版本。
	//
	// Protocol1v21v110 is the protocol version for Minecraft 1.21.110.
	Protocol1v21v110 int32 = 844
	// Version1v21v110 是协议 844 对应的 Minecraft 版本字符串。
	//
	// Version1v21v110 is the Minecraft version string for protocol 844.
	Version1v21v110 = "1.21.110"
)

// Profile1v21v110v26 是 Minecraft 1.21.110.26 的协议配置。
//
// Profile1v21v110v26 is the protocol profile for Minecraft 1.21.110.26.
var Profile1v21v110v26 = NewProfile(Protocol1v21v110v26, Version1v21v110v26)

const (
	// Protocol1v21v110v26 是 Minecraft 1.21.110.26 的协议版本。
	//
	// Protocol1v21v110v26 is the protocol version for Minecraft 1.21.110.26.
	Protocol1v21v110v26 int32 = 843
	// Version1v21v110v26 是协议 843 对应的 Minecraft 版本字符串。
	//
	// Version1v21v110v26 is the Minecraft version string for protocol 843.
	Version1v21v110v26 = "1.21.110.26"
)

// Profile1v21v100 是 Minecraft 1.21.100 的协议配置。
//
// Profile1v21v100 is the protocol profile for Minecraft 1.21.100.
var Profile1v21v100 = NewProfile(Protocol1v21v100, Version1v21v100)

const (
	// Protocol1v21v100 是 Minecraft 1.21.100 的协议版本。
	//
	// Protocol1v21v100 is the protocol version for Minecraft 1.21.100.
	Protocol1v21v100 int32 = 827
	// Version1v21v100 是协议 827 对应的 Minecraft 版本字符串。
	//
	// Version1v21v100 is the Minecraft version string for protocol 827.
	Version1v21v100 = "1.21.100"
)

// Profile1v21v93 是 Minecraft 1.21.93 的协议配置。
//
// Profile1v21v93 is the protocol profile for Minecraft 1.21.93.
var Profile1v21v93 = NewProfile(Protocol1v21v93, Version1v21v93)

const (
	// Protocol1v21v93 是 Minecraft 1.21.93 的协议版本。
	//
	// Protocol1v21v93 is the protocol version for Minecraft 1.21.93.
	Protocol1v21v93 int32 = 819
	// Version1v21v93 是协议 819 对应的 Minecraft 版本字符串。
	//
	// Version1v21v93 is the Minecraft version string for protocol 819.
	Version1v21v93 = "1.21.93"
)

// Profile1v21v90 是 Minecraft 1.21.90 的协议配置。
//
// Profile1v21v90 is the protocol profile for Minecraft 1.21.90.
var Profile1v21v90 = NewProfile(Protocol1v21v90, Version1v21v90)

const (
	// Protocol1v21v90 是 Minecraft 1.21.90 的协议版本。
	//
	// Protocol1v21v90 is the protocol version for Minecraft 1.21.90.
	Protocol1v21v90 int32 = 818
	// Version1v21v90 是协议 818 对应的 Minecraft 版本字符串。
	//
	// Version1v21v90 is the Minecraft version string for protocol 818.
	Version1v21v90 = "1.21.90"
)

// Profile1v21v80 是 Minecraft 1.21.80 的协议配置。
//
// Profile1v21v80 is the protocol profile for Minecraft 1.21.80.
var Profile1v21v80 = NewProfile(Protocol1v21v80, Version1v21v80)

const (
	// Protocol1v21v80 是 Minecraft 1.21.80 的协议版本。
	//
	// Protocol1v21v80 is the protocol version for Minecraft 1.21.80.
	Protocol1v21v80 int32 = 800
	// Version1v21v80 是协议 800 对应的 Minecraft 版本字符串。
	//
	// Version1v21v80 is the Minecraft version string for protocol 800.
	Version1v21v80 = "1.21.80"
)

// Profile1v21v70 是 Minecraft 1.21.70 的协议配置。
//
// Profile1v21v70 is the protocol profile for Minecraft 1.21.70.
var Profile1v21v70 = NewProfile(Protocol1v21v70, Version1v21v70)

const (
	// Protocol1v21v70 是 Minecraft 1.21.70 的协议版本。
	//
	// Protocol1v21v70 is the protocol version for Minecraft 1.21.70.
	Protocol1v21v70 int32 = 786
	// Version1v21v70 是协议 786 对应的 Minecraft 版本字符串。
	//
	// Version1v21v70 is the Minecraft version string for protocol 786.
	Version1v21v70 = "1.21.70"
)

// Profile1v21v70v24 是 Minecraft 1.21.70.24 的协议配置。
//
// Profile1v21v70v24 is the protocol profile for Minecraft 1.21.70.24.
var Profile1v21v70v24 = NewProfile(Protocol1v21v70v24, Version1v21v70v24)

const (
	// Protocol1v21v70v24 是 Minecraft 1.21.70.24 的协议版本。
	//
	// Protocol1v21v70v24 is the protocol version for Minecraft 1.21.70.24.
	Protocol1v21v70v24 int32 = 785
	// Version1v21v70v24 是协议 785 对应的 Minecraft 版本字符串。
	//
	// Version1v21v70v24 is the Minecraft version string for protocol 785.
	Version1v21v70v24 = "1.21.70.24"
)

// Profile1v21v60 是 Minecraft 1.21.60 的协议配置。
//
// Profile1v21v60 is the protocol profile for Minecraft 1.21.60.
var Profile1v21v60 = NewProfile(Protocol1v21v60, Version1v21v60)

const (
	// Protocol1v21v60 是 Minecraft 1.21.60 的协议版本。
	//
	// Protocol1v21v60 is the protocol version for Minecraft 1.21.60.
	Protocol1v21v60 int32 = 776
	// Version1v21v60 是协议 776 对应的 Minecraft 版本字符串。
	//
	// Version1v21v60 is the Minecraft version string for protocol 776.
	Version1v21v60 = "1.21.60"
)

// Profile1v21v50 是 Minecraft 1.21.50 的协议配置。
//
// Profile1v21v50 is the protocol profile for Minecraft 1.21.50.
var Profile1v21v50 = NewProfile(Protocol1v21v50, Version1v21v50)

const (
	// Protocol1v21v50 是 Minecraft 1.21.50 的协议版本。
	//
	// Protocol1v21v50 is the protocol version for Minecraft 1.21.50.
	Protocol1v21v50 int32 = 766
	// Version1v21v50 是协议 766 对应的 Minecraft 版本字符串。
	//
	// Version1v21v50 is the Minecraft version string for protocol 766.
	Version1v21v50 = "1.21.50"
)

// Profile1v21v50v26 是 Minecraft 1.21.50.26 的协议配置。
//
// Profile1v21v50v26 is the protocol profile for Minecraft 1.21.50.26.
var Profile1v21v50v26 = NewProfile(Protocol1v21v50v26, Version1v21v50v26)

const (
	// Protocol1v21v50v26 是 Minecraft 1.21.50.26 的协议版本。
	//
	// Protocol1v21v50v26 is the protocol version for Minecraft 1.21.50.26.
	Protocol1v21v50v26 int32 = 765
	// Version1v21v50v26 是协议 765 对应的 Minecraft 版本字符串。
	//
	// Version1v21v50v26 is the Minecraft version string for protocol 765.
	Version1v21v50v26 = "1.21.50.26"
)

// Profile1v21v40 是 Minecraft 1.21.40 的协议配置。
//
// Profile1v21v40 is the protocol profile for Minecraft 1.21.40.
var Profile1v21v40 = NewProfile(Protocol1v21v40, Version1v21v40)

const (
	// Protocol1v21v40 是 Minecraft 1.21.40 的协议版本。
	//
	// Protocol1v21v40 is the protocol version for Minecraft 1.21.40.
	Protocol1v21v40 int32 = 748
	// Version1v21v40 是协议 748 对应的 Minecraft 版本字符串。
	//
	// Version1v21v40 is the Minecraft version string for protocol 748.
	Version1v21v40 = "1.21.40"
)

// Profile1v21v30 是 Minecraft 1.21.30 的协议配置。
//
// Profile1v21v30 is the protocol profile for Minecraft 1.21.30.
var Profile1v21v30 = NewProfile(Protocol1v21v30, Version1v21v30)

const (
	// Protocol1v21v30 是 Minecraft 1.21.30 的协议版本。
	//
	// Protocol1v21v30 is the protocol version for Minecraft 1.21.30.
	Protocol1v21v30 int32 = 729
	// Version1v21v30 是协议 729 对应的 Minecraft 版本字符串。
	//
	// Version1v21v30 is the Minecraft version string for protocol 729.
	Version1v21v30 = "1.21.30"
)

// Profile1v21v20 是 Minecraft 1.21.20 的协议配置。
//
// Profile1v21v20 is the protocol profile for Minecraft 1.21.20.
var Profile1v21v20 = NewProfile(Protocol1v21v20, Version1v21v20)

const (
	// Protocol1v21v20 是 Minecraft 1.21.20 的协议版本。
	//
	// Protocol1v21v20 is the protocol version for Minecraft 1.21.20.
	Protocol1v21v20 int32 = 712
	// Version1v21v20 是协议 712 对应的 Minecraft 版本字符串。
	//
	// Version1v21v20 is the Minecraft version string for protocol 712.
	Version1v21v20 = "1.21.20"
)

// Profile1v21v2 是 Minecraft 1.21.2 的协议配置。
//
// Profile1v21v2 is the protocol profile for Minecraft 1.21.2.
var Profile1v21v2 = NewProfile(Protocol1v21v2, Version1v21v2)

const (
	// Protocol1v21v2 是 Minecraft 1.21.2 的协议版本。
	//
	// Protocol1v21v2 is the protocol version for Minecraft 1.21.2.
	Protocol1v21v2 int32 = 686
	// Version1v21v2 是协议 686 对应的 Minecraft 版本字符串。
	//
	// Version1v21v2 is the Minecraft version string for protocol 686.
	Version1v21v2 = "1.21.2"
)

// Profile1v21v0 是 Minecraft 1.21.0 的协议配置。
//
// Profile1v21v0 is the protocol profile for Minecraft 1.21.0.
var Profile1v21v0 = NewProfile(Protocol1v21v0, Version1v21v0)

const (
	// Protocol1v21v0 是 Minecraft 1.21.0 的协议版本。
	//
	// Protocol1v21v0 is the protocol version for Minecraft 1.21.0.
	Protocol1v21v0 int32 = 685
	// Version1v21v0 是协议 685 对应的 Minecraft 版本字符串。
	//
	// Version1v21v0 is the Minecraft version string for protocol 685.
	Version1v21v0 = "1.21.0"
)

// Profile1v20v80 是 Minecraft 1.20.80 的协议配置。
//
// Profile1v20v80 is the protocol profile for Minecraft 1.20.80.
var Profile1v20v80 = NewProfile(Protocol1v20v80, Version1v20v80)

const (
	// Protocol1v20v80 是 Minecraft 1.20.80 的协议版本。
	//
	// Protocol1v20v80 is the protocol version for Minecraft 1.20.80.
	Protocol1v20v80 int32 = 671
	// Version1v20v80 是协议 671 对应的 Minecraft 版本字符串。
	//
	// Version1v20v80 is the Minecraft version string for protocol 671.
	Version1v20v80 = "1.20.80"
)

// Profile1v20v70 是 Minecraft 1.20.70 的协议配置。
//
// Profile1v20v70 is the protocol profile for Minecraft 1.20.70.
var Profile1v20v70 = NewProfile(Protocol1v20v70, Version1v20v70)

const (
	// Protocol1v20v70 是 Minecraft 1.20.70 的协议版本。
	//
	// Protocol1v20v70 is the protocol version for Minecraft 1.20.70.
	Protocol1v20v70 int32 = 662
	// Version1v20v70 是协议 662 对应的 Minecraft 版本字符串。
	//
	// Version1v20v70 is the Minecraft version string for protocol 662.
	Version1v20v70 = "1.20.70"
)

// Profile1v20v60 是 Minecraft 1.20.60 的协议配置。
//
// Profile1v20v60 is the protocol profile for Minecraft 1.20.60.
var Profile1v20v60 = NewProfile(Protocol1v20v60, Version1v20v60)

const (
	// Protocol1v20v60 是 Minecraft 1.20.60 的协议版本。
	//
	// Protocol1v20v60 is the protocol version for Minecraft 1.20.60.
	Protocol1v20v60 int32 = 649
	// Version1v20v60 是协议 649 对应的 Minecraft 版本字符串。
	//
	// Version1v20v60 is the Minecraft version string for protocol 649.
	Version1v20v60 = "1.20.60"
)

// Profile1v20v50 是 Minecraft 1.20.50 的协议配置。
//
// Profile1v20v50 is the protocol profile for Minecraft 1.20.50.
var Profile1v20v50 = NewProfile(Protocol1v20v50, Version1v20v50)

const (
	// Protocol1v20v50 是 Minecraft 1.20.50 的协议版本。
	//
	// Protocol1v20v50 is the protocol version for Minecraft 1.20.50.
	Protocol1v20v50 int32 = 630
	// Version1v20v50 是协议 630 对应的 Minecraft 版本字符串。
	//
	// Version1v20v50 is the Minecraft version string for protocol 630.
	Version1v20v50 = "1.20.50"
)

// Profile1v20v40 是 Minecraft 1.20.40 的协议配置。
//
// Profile1v20v40 is the protocol profile for Minecraft 1.20.40.
var Profile1v20v40 = NewProfile(Protocol1v20v40, Version1v20v40)

const (
	// Protocol1v20v40 是 Minecraft 1.20.40 的协议版本。
	//
	// Protocol1v20v40 is the protocol version for Minecraft 1.20.40.
	Protocol1v20v40 int32 = 622
	// Version1v20v40 是协议 622 对应的 Minecraft 版本字符串。
	//
	// Version1v20v40 is the Minecraft version string for protocol 622.
	Version1v20v40 = "1.20.40"
)

// Profile1v20v30 是 Minecraft 1.20.30 的协议配置。
//
// Profile1v20v30 is the protocol profile for Minecraft 1.20.30.
var Profile1v20v30 = NewProfile(Protocol1v20v30, Version1v20v30)

const (
	// Protocol1v20v30 是 Minecraft 1.20.30 的协议版本。
	//
	// Protocol1v20v30 is the protocol version for Minecraft 1.20.30.
	Protocol1v20v30 int32 = 618
	// Version1v20v30 是协议 618 对应的 Minecraft 版本字符串。
	//
	// Version1v20v30 is the Minecraft version string for protocol 618.
	Version1v20v30 = "1.20.30"
)

// Profile1v20v30v24 是 Minecraft 1.20.30.24 的协议配置。
//
// Profile1v20v30v24 is the protocol profile for Minecraft 1.20.30.24.
var Profile1v20v30v24 = NewProfile(Protocol1v20v30v24, Version1v20v30v24)

const (
	// Protocol1v20v30v24 是 Minecraft 1.20.30.24 的协议版本。
	//
	// Protocol1v20v30v24 is the protocol version for Minecraft 1.20.30.24.
	Protocol1v20v30v24 int32 = 617
	// Version1v20v30v24 是协议 617 对应的 Minecraft 版本字符串。
	//
	// Version1v20v30v24 is the Minecraft version string for protocol 617.
	Version1v20v30v24 = "1.20.30.24"
)

// Profile1v20v10 是 Minecraft 1.20.10 的协议配置。
//
// Profile1v20v10 is the protocol profile for Minecraft 1.20.10.
var Profile1v20v10 = NewProfile(Protocol1v20v10, Version1v20v10)

const (
	// Protocol1v20v10 是 Minecraft 1.20.10 的协议版本。
	//
	// Protocol1v20v10 is the protocol version for Minecraft 1.20.10.
	Protocol1v20v10 int32 = 594
	// Version1v20v10 是协议 594 对应的 Minecraft 版本字符串。
	//
	// Version1v20v10 is the Minecraft version string for protocol 594.
	Version1v20v10 = "1.20.10"
)

// Profile1v20v0 是 Minecraft 1.20.0 的协议配置。
//
// Profile1v20v0 is the protocol profile for Minecraft 1.20.0.
var Profile1v20v0 = NewProfile(Protocol1v20v0, Version1v20v0)

const (
	// Protocol1v20v0 是 Minecraft 1.20.0 的协议版本。
	//
	// Protocol1v20v0 is the protocol version for Minecraft 1.20.0.
	Protocol1v20v0 int32 = 589
	// Version1v20v0 是协议 589 对应的 Minecraft 版本字符串。
	//
	// Version1v20v0 is the Minecraft version string for protocol 589.
	Version1v20v0 = "1.20.0"
)

// Profile1v20v0v23 是 Minecraft 1.20.0.23 的协议配置。
//
// Profile1v20v0v23 is the protocol profile for Minecraft 1.20.0.23.
var Profile1v20v0v23 = NewProfile(Protocol1v20v0v23, Version1v20v0v23)

const (
	// Protocol1v20v0v23 是 Minecraft 1.20.0.23 的协议版本。
	//
	// Protocol1v20v0v23 is the protocol version for Minecraft 1.20.0.23.
	Protocol1v20v0v23 int32 = 588
	// Version1v20v0v23 是协议 588 对应的 Minecraft 版本字符串。
	//
	// Version1v20v0v23 is the Minecraft version string for protocol 588.
	Version1v20v0v23 = "1.20.0.23"
)

// Profile1v19v80 是 Minecraft 1.19.80 的协议配置。
//
// Profile1v19v80 is the protocol profile for Minecraft 1.19.80.
var Profile1v19v80 = NewProfile(Protocol1v19v80, Version1v19v80)

const (
	// Protocol1v19v80 是 Minecraft 1.19.80 的协议版本。
	//
	// Protocol1v19v80 is the protocol version for Minecraft 1.19.80.
	Protocol1v19v80 int32 = 582
	// Version1v19v80 是协议 582 对应的 Minecraft 版本字符串。
	//
	// Version1v19v80 is the Minecraft version string for protocol 582.
	Version1v19v80 = "1.19.80"
)

// Profile1v19v70 是 Minecraft 1.19.70 的协议配置。
//
// Profile1v19v70 is the protocol profile for Minecraft 1.19.70.
var Profile1v19v70 = NewProfile(Protocol1v19v70, Version1v19v70)

const (
	// Protocol1v19v70 是 Minecraft 1.19.70 的协议版本。
	//
	// Protocol1v19v70 is the protocol version for Minecraft 1.19.70.
	Protocol1v19v70 int32 = 575
	// Version1v19v70 是协议 575 对应的 Minecraft 版本字符串。
	//
	// Version1v19v70 is the Minecraft version string for protocol 575.
	Version1v19v70 = "1.19.70"
)

// Profile1v19v70v24 是 Minecraft 1.19.70.24 的协议配置。
//
// Profile1v19v70v24 is the protocol profile for Minecraft 1.19.70.24.
var Profile1v19v70v24 = NewProfile(Protocol1v19v70v24, Version1v19v70v24)

const (
	// Protocol1v19v70v24 是 Minecraft 1.19.70.24 的协议版本。
	//
	// Protocol1v19v70v24 is the protocol version for Minecraft 1.19.70.24.
	Protocol1v19v70v24 int32 = 574
	// Version1v19v70v24 是协议 574 对应的 Minecraft 版本字符串。
	//
	// Version1v19v70v24 is the Minecraft version string for protocol 574.
	Version1v19v70v24 = "1.19.70.24"
)

// Profile1v19v63 是 Minecraft 1.19.63 的协议配置。
//
// Profile1v19v63 is the protocol profile for Minecraft 1.19.63.
var Profile1v19v63 = NewProfile(Protocol1v19v63, Version1v19v63)

const (
	// Protocol1v19v63 是 Minecraft 1.19.63 的协议版本。
	//
	// Protocol1v19v63 is the protocol version for Minecraft 1.19.63.
	Protocol1v19v63 int32 = 568
	// Version1v19v63 是协议 568 对应的 Minecraft 版本字符串。
	//
	// Version1v19v63 is the Minecraft version string for protocol 568.
	Version1v19v63 = "1.19.63"
)

// Profile1v19v60 是 Minecraft 1.19.60 的协议配置。
//
// Profile1v19v60 is the protocol profile for Minecraft 1.19.60.
var Profile1v19v60 = NewProfile(Protocol1v19v60, Version1v19v60)

const (
	// Protocol1v19v60 是 Minecraft 1.19.60 的协议版本。
	//
	// Protocol1v19v60 is the protocol version for Minecraft 1.19.60.
	Protocol1v19v60 int32 = 567
	// Version1v19v60 是协议 567 对应的 Minecraft 版本字符串。
	//
	// Version1v19v60 is the Minecraft version string for protocol 567.
	Version1v19v60 = "1.19.60"
)

// Profile1v19v50 是 Minecraft 1.19.50 的协议配置。
//
// Profile1v19v50 is the protocol profile for Minecraft 1.19.50.
var Profile1v19v50 = NewProfile(Protocol1v19v50, Version1v19v50)

const (
	// Protocol1v19v50 是 Minecraft 1.19.50 的协议版本。
	//
	// Protocol1v19v50 is the protocol version for Minecraft 1.19.50.
	Protocol1v19v50 int32 = 560
	// Version1v19v50 是协议 560 对应的 Minecraft 版本字符串。
	//
	// Version1v19v50 is the Minecraft version string for protocol 560.
	Version1v19v50 = "1.19.50"
)

// Profile1v19v40 是 Minecraft 1.19.40 的协议配置。
//
// Profile1v19v40 is the protocol profile for Minecraft 1.19.40.
var Profile1v19v40 = NewProfile(Protocol1v19v40, Version1v19v40)

const (
	// Protocol1v19v40 是 Minecraft 1.19.40 的协议版本。
	//
	// Protocol1v19v40 is the protocol version for Minecraft 1.19.40.
	Protocol1v19v40 int32 = 557
	// Version1v19v40 是协议 557 对应的 Minecraft 版本字符串。
	//
	// Version1v19v40 is the Minecraft version string for protocol 557.
	Version1v19v40 = "1.19.40"
)

// Profile1v19v30 是 Minecraft 1.19.30 的协议配置。
//
// Profile1v19v30 is the protocol profile for Minecraft 1.19.30.
var Profile1v19v30 = NewProfile(Protocol1v19v30, Version1v19v30)

const (
	// Protocol1v19v30 是 Minecraft 1.19.30 的协议版本。
	//
	// Protocol1v19v30 is the protocol version for Minecraft 1.19.30.
	Protocol1v19v30 int32 = 554
	// Version1v19v30 是协议 554 对应的 Minecraft 版本字符串。
	//
	// Version1v19v30 is the Minecraft version string for protocol 554.
	Version1v19v30 = "1.19.30"
)

// Profile1v19v21 是 Minecraft 1.19.21 的协议配置。
//
// Profile1v19v21 is the protocol profile for Minecraft 1.19.21.
var Profile1v19v21 = NewProfile(Protocol1v19v21, Version1v19v21)

const (
	// Protocol1v19v21 是 Minecraft 1.19.21 的协议版本。
	//
	// Protocol1v19v21 is the protocol version for Minecraft 1.19.21.
	Protocol1v19v21 int32 = 545
	// Version1v19v21 是协议 545 对应的 Minecraft 版本字符串。
	//
	// Version1v19v21 is the Minecraft version string for protocol 545.
	Version1v19v21 = "1.19.21"
)

// Profile1v19v20 是 Minecraft 1.19.20 的协议配置。
//
// Profile1v19v20 is the protocol profile for Minecraft 1.19.20.
var Profile1v19v20 = NewProfile(Protocol1v19v20, Version1v19v20)

const (
	// Protocol1v19v20 是 Minecraft 1.19.20 的协议版本。
	//
	// Protocol1v19v20 is the protocol version for Minecraft 1.19.20.
	Protocol1v19v20 int32 = 544
	// Version1v19v20 是协议 544 对应的 Minecraft 版本字符串。
	//
	// Version1v19v20 is the Minecraft version string for protocol 544.
	Version1v19v20 = "1.19.20"
)

// Profile1v19v10 是 Minecraft 1.19.10 的协议配置。
//
// Profile1v19v10 is the protocol profile for Minecraft 1.19.10.
var Profile1v19v10 = NewProfile(Protocol1v19v10, Version1v19v10)

const (
	// Protocol1v19v10 是 Minecraft 1.19.10 的协议版本。
	//
	// Protocol1v19v10 is the protocol version for Minecraft 1.19.10.
	Protocol1v19v10 int32 = 534
	// Version1v19v10 是协议 534 对应的 Minecraft 版本字符串。
	//
	// Version1v19v10 is the Minecraft version string for protocol 534.
	Version1v19v10 = "1.19.10"
)

// Profile1v19v0 是 Minecraft 1.19.0 的协议配置。
//
// Profile1v19v0 is the protocol profile for Minecraft 1.19.0.
var Profile1v19v0 = NewProfile(Protocol1v19v0, Version1v19v0)

const (
	// Protocol1v19v0 是 Minecraft 1.19.0 的协议版本。
	//
	// Protocol1v19v0 is the protocol version for Minecraft 1.19.0.
	Protocol1v19v0 int32 = 527
	// Version1v19v0 是协议 527 对应的 Minecraft 版本字符串。
	//
	// Version1v19v0 is the Minecraft version string for protocol 527.
	Version1v19v0 = "1.19.0"
)

// Profile1v19v0v29 是 Minecraft 1.19.0.29 的协议配置。
//
// Profile1v19v0v29 is the protocol profile for Minecraft 1.19.0.29.
var Profile1v19v0v29 = NewProfile(Protocol1v19v0v29, Version1v19v0v29)

const (
	// Protocol1v19v0v29 是 Minecraft 1.19.0.29 的协议版本。
	//
	// Protocol1v19v0v29 is the protocol version for Minecraft 1.19.0.29.
	Protocol1v19v0v29 int32 = 524
	// Version1v19v0v29 是协议 524 对应的 Minecraft 版本字符串。
	//
	// Version1v19v0v29 is the Minecraft version string for protocol 524.
	Version1v19v0v29 = "1.19.0.29"
)

// Profile1v18v30 是 Minecraft 1.18.30 的协议配置。
//
// Profile1v18v30 is the protocol profile for Minecraft 1.18.30.
var Profile1v18v30 = NewProfile(Protocol1v18v30, Version1v18v30)

const (
	// Protocol1v18v30 是 Minecraft 1.18.30 的协议版本。
	//
	// Protocol1v18v30 is the protocol version for Minecraft 1.18.30.
	Protocol1v18v30 int32 = 503
	// Version1v18v30 是协议 503 对应的 Minecraft 版本字符串。
	//
	// Version1v18v30 is the Minecraft version string for protocol 503.
	Version1v18v30 = "1.18.30"
)

// Profile1v18v10 是 Minecraft 1.18.10 的协议配置。
//
// Profile1v18v10 is the protocol profile for Minecraft 1.18.10.
var Profile1v18v10 = NewProfile(Protocol1v18v10, Version1v18v10)

const (
	// Protocol1v18v10 是 Minecraft 1.18.10 的协议版本。
	//
	// Protocol1v18v10 is the protocol version for Minecraft 1.18.10.
	Protocol1v18v10 int32 = 486
	// Version1v18v10 是协议 486 对应的 Minecraft 版本字符串。
	//
	// Version1v18v10 is the Minecraft version string for protocol 486.
	Version1v18v10 = "1.18.10"
)

// Profile1v18v10v26 是 Minecraft 1.18.10.26 的协议配置。
//
// Profile1v18v10v26 is the protocol profile for Minecraft 1.18.10.26.
var Profile1v18v10v26 = NewProfile(Protocol1v18v10v26, Version1v18v10v26)

const (
	// Protocol1v18v10v26 是 Minecraft 1.18.10.26 的协议版本。
	//
	// Protocol1v18v10v26 is the protocol version for Minecraft 1.18.10.26.
	Protocol1v18v10v26 int32 = 485
	// Version1v18v10v26 是协议 485 对应的 Minecraft 版本字符串。
	//
	// Version1v18v10v26 is the Minecraft version string for protocol 485.
	Version1v18v10v26 = "1.18.10.26"
)

// Profile1v18v0 是 Minecraft 1.18.0 的协议配置。
//
// Profile1v18v0 is the protocol profile for Minecraft 1.18.0.
var Profile1v18v0 = NewProfile(Protocol1v18v0, Version1v18v0)

const (
	// Protocol1v18v0 是 Minecraft 1.18.0 的协议版本。
	//
	// Protocol1v18v0 is the protocol version for Minecraft 1.18.0.
	Protocol1v18v0 int32 = 475
	// Version1v18v0 是协议 475 对应的 Minecraft 版本字符串。
	//
	// Version1v18v0 is the Minecraft version string for protocol 475.
	Version1v18v0 = "1.18.0"
)

// Profile1v17v40 是 Minecraft 1.17.40 的协议配置。
//
// Profile1v17v40 is the protocol profile for Minecraft 1.17.40.
var Profile1v17v40 = NewProfile(Protocol1v17v40, Version1v17v40)

const (
	// Protocol1v17v40 是 Minecraft 1.17.40 的协议版本。
	//
	// Protocol1v17v40 is the protocol version for Minecraft 1.17.40.
	Protocol1v17v40 int32 = 471
	// Version1v17v40 是协议 471 对应的 Minecraft 版本字符串。
	//
	// Version1v17v40 is the Minecraft version string for protocol 471.
	Version1v17v40 = "1.17.40"
)

// Profile1v17v30 是 Minecraft 1.17.30 的协议配置。
//
// Profile1v17v30 is the protocol profile for Minecraft 1.17.30.
var Profile1v17v30 = NewProfile(Protocol1v17v30, Version1v17v30)

const (
	// Protocol1v17v30 是 Minecraft 1.17.30 的协议版本。
	//
	// Protocol1v17v30 is the protocol version for Minecraft 1.17.30.
	Protocol1v17v30 int32 = 465
	// Version1v17v30 是协议 465 对应的 Minecraft 版本字符串。
	//
	// Version1v17v30 is the Minecraft version string for protocol 465.
	Version1v17v30 = "1.17.30"
)

// Profile1v17v20v20 是 Minecraft 1.17.20.20 的协议配置。
//
// Profile1v17v20v20 is the protocol profile for Minecraft 1.17.20.20.
var Profile1v17v20v20 = NewProfile(Protocol1v17v20v20, Version1v17v20v20)

const (
	// Protocol1v17v20v20 是 Minecraft 1.17.20.20 的协议版本。
	//
	// Protocol1v17v20v20 is the protocol version for Minecraft 1.17.20.20.
	Protocol1v17v20v20 int32 = 453
	// Version1v17v20v20 是协议 453 对应的 Minecraft 版本字符串。
	//
	// Version1v17v20v20 is the Minecraft version string for protocol 453.
	Version1v17v20v20 = "1.17.20.20"
)

// Profile1v17v10 是 Minecraft 1.17.10 的协议配置。
//
// Profile1v17v10 is the protocol profile for Minecraft 1.17.10.
var Profile1v17v10 = NewProfile(Protocol1v17v10, Version1v17v10)

const (
	// Protocol1v17v10 是 Minecraft 1.17.10 的协议版本。
	//
	// Protocol1v17v10 is the protocol version for Minecraft 1.17.10.
	Protocol1v17v10 int32 = 448
	// Version1v17v10 是协议 448 对应的 Minecraft 版本字符串。
	//
	// Version1v17v10 is the Minecraft version string for protocol 448.
	Version1v17v10 = "1.17.10"
)

// Profile1v17v0 是 Minecraft 1.17.0 的协议配置。
//
// Profile1v17v0 is the protocol profile for Minecraft 1.17.0.
var Profile1v17v0 = NewProfile(Protocol1v17v0, Version1v17v0)

const (
	// Protocol1v17v0 是 Minecraft 1.17.0 的协议版本。
	//
	// Protocol1v17v0 is the protocol version for Minecraft 1.17.0.
	Protocol1v17v0 int32 = 440
	// Version1v17v0 是协议 440 对应的 Minecraft 版本字符串。
	//
	// Version1v17v0 is the Minecraft version string for protocol 440.
	Version1v17v0 = "1.17.0"
)

// Profile1v16v230v54 是 Minecraft 1.16.230.54 的协议配置。
//
// Profile1v16v230v54 is the protocol profile for Minecraft 1.16.230.54.
var Profile1v16v230v54 = NewProfile(Protocol1v16v230v54, Version1v16v230v54)

const (
	// Protocol1v16v230v54 是 Minecraft 1.16.230.54 的协议版本。
	//
	// Protocol1v16v230v54 is the protocol version for Minecraft 1.16.230.54.
	Protocol1v16v230v54 int32 = 435
	// Version1v16v230v54 是协议 435 对应的 Minecraft 版本字符串。
	//
	// Version1v16v230v54 is the Minecraft version string for protocol 435.
	Version1v16v230v54 = "1.16.230.54"
)

// Profile1v16v230 是 Minecraft 1.16.230 的协议配置。
//
// Profile1v16v230 is the protocol profile for Minecraft 1.16.230.
var Profile1v16v230 = NewProfile(Protocol1v16v230, Version1v16v230)

const (
	// Protocol1v16v230 是 Minecraft 1.16.230 的协议版本。
	//
	// Protocol1v16v230 is the protocol version for Minecraft 1.16.230.
	Protocol1v16v230 int32 = 434
	// Version1v16v230 是协议 434 对应的 Minecraft 版本字符串。
	//
	// Version1v16v230 is the Minecraft version string for protocol 434.
	Version1v16v230 = "1.16.230"
)

// Profile1v16v230v50 是 Minecraft 1.16.230.50 的协议配置。
//
// Profile1v16v230v50 is the protocol profile for Minecraft 1.16.230.50.
var Profile1v16v230v50 = NewProfile(Protocol1v16v230v50, Version1v16v230v50)

const (
	// Protocol1v16v230v50 是 Minecraft 1.16.230.50 的协议版本。
	//
	// Protocol1v16v230v50 is the protocol version for Minecraft 1.16.230.50.
	Protocol1v16v230v50 int32 = 433
	// Version1v16v230v50 是协议 433 对应的 Minecraft 版本字符串。
	//
	// Version1v16v230v50 is the Minecraft version string for protocol 433.
	Version1v16v230v50 = "1.16.230.50"
)

// Profile1v16v220 是 Minecraft 1.16.220 的协议配置。
//
// Profile1v16v220 is the protocol profile for Minecraft 1.16.220.
var Profile1v16v220 = NewProfile(Protocol1v16v220, Version1v16v220)

const (
	// Protocol1v16v220 是 Minecraft 1.16.220 的协议版本。
	//
	// Protocol1v16v220 is the protocol version for Minecraft 1.16.220.
	Protocol1v16v220 int32 = 431
	// Version1v16v220 是协议 431 对应的 Minecraft 版本字符串。
	//
	// Version1v16v220 is the Minecraft version string for protocol 431.
	Version1v16v220 = "1.16.220"
)

// Profile1v16v210 是 Minecraft 1.16.210 的协议配置。
//
// Profile1v16v210 is the protocol profile for Minecraft 1.16.210.
var Profile1v16v210 = NewProfile(Protocol1v16v210, Version1v16v210)

const (
	// Protocol1v16v210 是 Minecraft 1.16.210 的协议版本。
	//
	// Protocol1v16v210 is the protocol version for Minecraft 1.16.210.
	Protocol1v16v210 int32 = 428
	// Version1v16v210 是协议 428 对应的 Minecraft 版本字符串。
	//
	// Version1v16v210 is the Minecraft version string for protocol 428.
	Version1v16v210 = "1.16.210"
)

// Profile1v16v210v50 是 Minecraft 1.16.210.50 的协议配置。
//
// Profile1v16v210v50 is the protocol profile for Minecraft 1.16.210.50.
var Profile1v16v210v50 = NewProfile(Protocol1v16v210v50, Version1v16v210v50)

const (
	// Protocol1v16v210v50 是 Minecraft 1.16.210.50 的协议版本。
	//
	// Protocol1v16v210v50 is the protocol version for Minecraft 1.16.210.50.
	Protocol1v16v210v50 int32 = 423
	// Version1v16v210v50 是协议 423 对应的 Minecraft 版本字符串。
	//
	// Version1v16v210v50 is the Minecraft version string for protocol 423.
	Version1v16v210v50 = "1.16.210.50"
)

// Profile1v16v200 是 Minecraft 1.16.200 的协议配置。
//
// Profile1v16v200 is the protocol profile for Minecraft 1.16.200.
var Profile1v16v200 = NewProfile(Protocol1v16v200, Version1v16v200)

const (
	// Protocol1v16v200 是 Minecraft 1.16.200 的协议版本。
	//
	// Protocol1v16v200 is the protocol version for Minecraft 1.16.200.
	Protocol1v16v200 int32 = 422
	// Version1v16v200 是协议 422 对应的 Minecraft 版本字符串。
	//
	// Version1v16v200 is the Minecraft version string for protocol 422.
	Version1v16v200 = "1.16.200"
)

// Profile1v16v200v51 是 Minecraft 1.16.200.51 的协议配置。
//
// Profile1v16v200v51 is the protocol profile for Minecraft 1.16.200.51.
var Profile1v16v200v51 = NewProfile(Protocol1v16v200v51, Version1v16v200v51)

const (
	// Protocol1v16v200v51 是 Minecraft 1.16.200.51 的协议版本。
	//
	// Protocol1v16v200v51 is the protocol version for Minecraft 1.16.200.51.
	Protocol1v16v200v51 int32 = 420
	// Version1v16v200v51 是协议 420 对应的 Minecraft 版本字符串。
	//
	// Version1v16v200v51 is the Minecraft version string for protocol 420.
	Version1v16v200v51 = "1.16.200.51"
)

// Profile1v16v100 是 Minecraft 1.16.100 的协议配置。
//
// Profile1v16v100 is the protocol profile for Minecraft 1.16.100.
var Profile1v16v100 = NewProfile(Protocol1v16v100, Version1v16v100)

const (
	// Protocol1v16v100 是 Minecraft 1.16.100 的协议版本。
	//
	// Protocol1v16v100 is the protocol version for Minecraft 1.16.100.
	Protocol1v16v100 int32 = 419
	// Version1v16v100 是协议 419 对应的 Minecraft 版本字符串。
	//
	// Version1v16v100 is the Minecraft version string for protocol 419.
	Version1v16v100 = "1.16.100"
)

// Profile1v16v20 是 Minecraft 1.16.20 的协议配置。
//
// Profile1v16v20 is the protocol profile for Minecraft 1.16.20.
var Profile1v16v20 = NewProfile(Protocol1v16v20, Version1v16v20)

const (
	// Protocol1v16v20 是 Minecraft 1.16.20 的协议版本。
	//
	// Protocol1v16v20 is the protocol version for Minecraft 1.16.20.
	Protocol1v16v20 int32 = 408
	// Version1v16v20 是协议 408 对应的 Minecraft 版本字符串。
	//
	// Version1v16v20 is the Minecraft version string for protocol 408.
	Version1v16v20 = "1.16.20"
)

// Profile1v16v0 是 Minecraft 1.16.0 的协议配置。
//
// Profile1v16v0 is the protocol profile for Minecraft 1.16.0.
var Profile1v16v0 = NewProfile(Protocol1v16v0, Version1v16v0)

const (
	// Protocol1v16v0 是 Minecraft 1.16.0 的协议版本。
	//
	// Protocol1v16v0 is the protocol version for Minecraft 1.16.0.
	Protocol1v16v0 int32 = 407
	// Version1v16v0 是协议 407 对应的 Minecraft 版本字符串。
	//
	// Version1v16v0 is the Minecraft version string for protocol 407.
	Version1v16v0 = "1.16.0"
)

// Profile1v14v60 是 Minecraft 1.14.60 的协议配置。
//
// Profile1v14v60 is the protocol profile for Minecraft 1.14.60.
var Profile1v14v60 = NewProfile(Protocol1v14v60, Version1v14v60)

const (
	// Protocol1v14v60 是 Minecraft 1.14.60 的协议版本。
	//
	// Protocol1v14v60 is the protocol version for Minecraft 1.14.60.
	Protocol1v14v60 int32 = 390
	// Version1v14v60 是协议 390 对应的 Minecraft 版本字符串。
	//
	// Version1v14v60 is the Minecraft version string for protocol 390.
	Version1v14v60 = "1.14.60"
)

// Profile1v14v0 是 Minecraft 1.14.0 的协议配置。
//
// Profile1v14v0 is the protocol profile for Minecraft 1.14.0.
var Profile1v14v0 = NewProfile(Protocol1v14v0, Version1v14v0)

const (
	// Protocol1v14v0 是 Minecraft 1.14.0 的协议版本。
	//
	// Protocol1v14v0 is the protocol version for Minecraft 1.14.0.
	Protocol1v14v0 int32 = 389
	// Version1v14v0 是协议 389 对应的 Minecraft 版本字符串。
	//
	// Version1v14v0 is the Minecraft version string for protocol 389.
	Version1v14v0 = "1.14.0"
)

// Profile1v13v0 是 Minecraft 1.13.0 的协议配置。
//
// Profile1v13v0 is the protocol profile for Minecraft 1.13.0.
var Profile1v13v0 = NewProfile(Protocol1v13v0, Version1v13v0)

const (
	// Protocol1v13v0 是 Minecraft 1.13.0 的协议版本。
	//
	// Protocol1v13v0 is the protocol version for Minecraft 1.13.0.
	Protocol1v13v0 int32 = 388
	// Version1v13v0 是协议 388 对应的 Minecraft 版本字符串。
	//
	// Version1v13v0 is the Minecraft version string for protocol 388.
	Version1v13v0 = "1.13.0"
)

// Profile1v12v0 是 Minecraft 1.12.0 的协议配置。
//
// Profile1v12v0 is the protocol profile for Minecraft 1.12.0.
var Profile1v12v0 = NewProfile(Protocol1v12v0, Version1v12v0)

const (
	// Protocol1v12v0 是 Minecraft 1.12.0 的协议版本。
	//
	// Protocol1v12v0 is the protocol version for Minecraft 1.12.0.
	Protocol1v12v0 int32 = 361
	// Version1v12v0 是协议 361 对应的 Minecraft 版本字符串。
	//
	// Version1v12v0 is the Minecraft version string for protocol 361.
	Version1v12v0 = "1.12.0"
)

// Profile1v11v0 是 Minecraft 1.11.0 的协议配置。
//
// Profile1v11v0 is the protocol profile for Minecraft 1.11.0.
var Profile1v11v0 = NewProfile(Protocol1v11v0, Version1v11v0)

const (
	// Protocol1v11v0 是 Minecraft 1.11.0 的协议版本。
	//
	// Protocol1v11v0 is the protocol version for Minecraft 1.11.0.
	Protocol1v11v0 int32 = 354
	// Version1v11v0 是协议 354 对应的 Minecraft 版本字符串。
	//
	// Version1v11v0 is the Minecraft version string for protocol 354.
	Version1v11v0 = "1.11.0"
)

// Profile1v10v0 是 Minecraft 1.10.0 的协议配置。
//
// Profile1v10v0 is the protocol profile for Minecraft 1.10.0.
var Profile1v10v0 = NewProfile(Protocol1v10v0, Version1v10v0)

const (
	// Protocol1v10v0 是 Minecraft 1.10.0 的协议版本。
	//
	// Protocol1v10v0 is the protocol version for Minecraft 1.10.0.
	Protocol1v10v0 int32 = 340
	// Version1v10v0 是协议 340 对应的 Minecraft 版本字符串。
	//
	// Version1v10v0 is the Minecraft version string for protocol 340.
	Version1v10v0 = "1.10.0"
)

// Profile1v9v0 是 Minecraft 1.9.0 的协议配置。
//
// Profile1v9v0 is the protocol profile for Minecraft 1.9.0.
var Profile1v9v0 = NewProfile(Protocol1v9v0, Version1v9v0)

const (
	// Protocol1v9v0 是 Minecraft 1.9.0 的协议版本。
	//
	// Protocol1v9v0 is the protocol version for Minecraft 1.9.0.
	Protocol1v9v0 int32 = 332
	// Version1v9v0 是协议 332 对应的 Minecraft 版本字符串。
	//
	// Version1v9v0 is the Minecraft version string for protocol 332.
	Version1v9v0 = "1.9.0"
)

// Profile1v8v0 是 Minecraft 1.8.0 的协议配置。
//
// Profile1v8v0 is the protocol profile for Minecraft 1.8.0.
var Profile1v8v0 = NewProfile(Protocol1v8v0, Version1v8v0)

const (
	// Protocol1v8v0 是 Minecraft 1.8.0 的协议版本。
	//
	// Protocol1v8v0 is the protocol version for Minecraft 1.8.0.
	Protocol1v8v0 int32 = 313
	// Version1v8v0 是协议 313 对应的 Minecraft 版本字符串。
	//
	// Version1v8v0 is the Minecraft version string for protocol 313.
	Version1v8v0 = "1.8.0"
)

// Profile1v7v0 是 Minecraft 1.7.0 的协议配置。
//
// Profile1v7v0 is the protocol profile for Minecraft 1.7.0.
var Profile1v7v0 = NewProfile(Protocol1v7v0, Version1v7v0)

const (
	// Protocol1v7v0 是 Minecraft 1.7.0 的协议版本。
	//
	// Protocol1v7v0 is the protocol version for Minecraft 1.7.0.
	Protocol1v7v0 int32 = 291
	// Version1v7v0 是协议 291 对应的 Minecraft 版本字符串。
	//
	// Version1v7v0 is the Minecraft version string for protocol 291.
	Version1v7v0 = "1.7.0"
)

// Profile1v6v0 是 Minecraft 1.6.0 的协议配置。
//
// Profile1v6v0 is the protocol profile for Minecraft 1.6.0.
var Profile1v6v0 = NewProfile(Protocol1v6v0, Version1v6v0)

const (
	// Protocol1v6v0 是 Minecraft 1.6.0 的协议版本。
	//
	// Protocol1v6v0 is the protocol version for Minecraft 1.6.0.
	Protocol1v6v0 int32 = 282
	// Version1v6v0 是协议 282 对应的 Minecraft 版本字符串。
	//
	// Version1v6v0 is the Minecraft version string for protocol 282.
	Version1v6v0 = "1.6.0"
)

// Profile1v6v0v5 是 Minecraft 1.6.0.5 的协议配置。
//
// Profile1v6v0v5 is the protocol profile for Minecraft 1.6.0.5.
var Profile1v6v0v5 = NewProfile(Protocol1v6v0v5, Version1v6v0v5)

const (
	// Protocol1v6v0v5 是 Minecraft 1.6.0.5 的协议版本。
	//
	// Protocol1v6v0v5 is the protocol version for Minecraft 1.6.0.5.
	Protocol1v6v0v5 int32 = 281
	// Version1v6v0v5 是协议 281 对应的 Minecraft 版本字符串。
	//
	// Version1v6v0v5 is the Minecraft version string for protocol 281.
	Version1v6v0v5 = "1.6.0.5"
)

// Profile1v5v0 是 Minecraft 1.5.0 的协议配置。
//
// Profile1v5v0 is the protocol profile for Minecraft 1.5.0.
var Profile1v5v0 = NewProfile(Protocol1v5v0, Version1v5v0)

const (
	// Protocol1v5v0 是 Minecraft 1.5.0 的协议版本。
	//
	// Protocol1v5v0 is the protocol version for Minecraft 1.5.0.
	Protocol1v5v0 int32 = 274
	// Version1v5v0 是协议 274 对应的 Minecraft 版本字符串。
	//
	// Version1v5v0 is the Minecraft version string for protocol 274.
	Version1v5v0 = "1.5.0"
)

// Profile1v4v0 是 Minecraft 1.4.0 的协议配置。
//
// Profile1v4v0 is the protocol profile for Minecraft 1.4.0.
var Profile1v4v0 = NewProfile(Protocol1v4v0, Version1v4v0)

const (
	// Protocol1v4v0 是 Minecraft 1.4.0 的协议版本。
	//
	// Protocol1v4v0 is the protocol version for Minecraft 1.4.0.
	Protocol1v4v0 int32 = 261
	// Version1v4v0 是协议 261 对应的 Minecraft 版本字符串。
	//
	// Version1v4v0 is the Minecraft version string for protocol 261.
	Version1v4v0 = "1.4.0"
)

// Profile1v2v13v11 是 Minecraft 1.2.13.11 的协议配置。
//
// Profile1v2v13v11 is the protocol profile for Minecraft 1.2.13.11.
var Profile1v2v13v11 = NewProfile(Protocol1v2v13v11, Version1v2v13v11)

const (
	// Protocol1v2v13v11 是 Minecraft 1.2.13.11 的协议版本。
	//
	// Protocol1v2v13v11 is the protocol version for Minecraft 1.2.13.11.
	Protocol1v2v13v11 int32 = 224
	// Version1v2v13v11 是协议 224 对应的 Minecraft 版本字符串。
	//
	// Version1v2v13v11 is the Minecraft version string for protocol 224.
	Version1v2v13v11 = "1.2.13.11"
)

// Profile1v2v13 是 Minecraft 1.2.13 的协议配置。
//
// Profile1v2v13 is the protocol profile for Minecraft 1.2.13.
var Profile1v2v13 = NewProfile(Protocol1v2v13, Version1v2v13)

const (
	// Protocol1v2v13 是 Minecraft 1.2.13 的协议版本。
	//
	// Protocol1v2v13 is the protocol version for Minecraft 1.2.13.
	Protocol1v2v13 int32 = 223
	// Version1v2v13 是协议 223 对应的 Minecraft 版本字符串。
	//
	// Version1v2v13 is the Minecraft version string for protocol 223.
	Version1v2v13 = "1.2.13"
)

// Profile1v2v10 是 Minecraft 1.2.10 的协议配置。
//
// Profile1v2v10 is the protocol profile for Minecraft 1.2.10.
var Profile1v2v10 = NewProfile(Protocol1v2v10, Version1v2v10)

const (
	// Protocol1v2v10 是 Minecraft 1.2.10 的协议版本。
	//
	// Protocol1v2v10 is the protocol version for Minecraft 1.2.10.
	Protocol1v2v10 int32 = 201
	// Version1v2v10 是协议 201 对应的 Minecraft 版本字符串。
	//
	// Version1v2v10 is the Minecraft version string for protocol 201.
	Version1v2v10 = "1.2.10"
)

// Profile1v2v7 是 Minecraft 1.2.7 的协议配置。
//
// Profile1v2v7 is the protocol profile for Minecraft 1.2.7.
var Profile1v2v7 = NewProfile(Protocol1v2v7, Version1v2v7)

const (
	// Protocol1v2v7 是 Minecraft 1.2.7 的协议版本。
	//
	// Protocol1v2v7 is the protocol version for Minecraft 1.2.7.
	Protocol1v2v7 int32 = 160
	// Version1v2v7 是协议 160 对应的 Minecraft 版本字符串。
	//
	// Version1v2v7 is the Minecraft version string for protocol 160.
	Version1v2v7 = "1.2.7"
)

// Profile1v2v6 是 Minecraft 1.2.6 的协议配置。
//
// Profile1v2v6 is the protocol profile for Minecraft 1.2.6.
var Profile1v2v6 = NewProfile(Protocol1v2v6, Version1v2v6)

const (
	// Protocol1v2v6 是 Minecraft 1.2.6 的协议版本。
	//
	// Protocol1v2v6 is the protocol version for Minecraft 1.2.6.
	Protocol1v2v6 int32 = 150
	// Version1v2v6 是协议 150 对应的 Minecraft 版本字符串。
	//
	// Version1v2v6 is the Minecraft version string for protocol 150.
	Version1v2v6 = "1.2.6"
)

// Profile1v2v5v11 是 Minecraft 1.2.5.11 的协议配置。
//
// Profile1v2v5v11 is the protocol profile for Minecraft 1.2.5.11.
var Profile1v2v5v11 = NewProfile(Protocol1v2v5v11, Version1v2v5v11)

const (
	// Protocol1v2v5v11 是 Minecraft 1.2.5.11 的协议版本。
	//
	// Protocol1v2v5v11 is the protocol version for Minecraft 1.2.5.11.
	Protocol1v2v5v11 int32 = 140
	// Version1v2v5v11 是协议 140 对应的 Minecraft 版本字符串。
	//
	// Version1v2v5v11 is the Minecraft version string for protocol 140.
	Version1v2v5v11 = "1.2.5.11"
)

// Profile1v2v5 是 Minecraft 1.2.5 的协议配置。
//
// Profile1v2v5 is the protocol profile for Minecraft 1.2.5.
var Profile1v2v5 = NewProfile(Protocol1v2v5, Version1v2v5)

const (
	// Protocol1v2v5 是 Minecraft 1.2.5 的协议版本。
	//
	// Protocol1v2v5 is the protocol version for Minecraft 1.2.5.
	Protocol1v2v5 int32 = 141
	// Version1v2v5 是协议 141 对应的 Minecraft 版本字符串。
	//
	// Version1v2v5 is the Minecraft version string for protocol 141.
	Version1v2v5 = "1.2.5"
)

// Profile1v2v0 是 Minecraft 1.2.0 的协议配置。
//
// Profile1v2v0 is the protocol profile for Minecraft 1.2.0.
var Profile1v2v0 = NewProfile(Protocol1v2v0, Version1v2v0)

const (
	// Protocol1v2v0 是 Minecraft 1.2.0 的协议版本。
	//
	// Protocol1v2v0 is the protocol version for Minecraft 1.2.0.
	Protocol1v2v0 int32 = 137
	// Version1v2v0 是协议 137 对应的 Minecraft 版本字符串。
	//
	// Version1v2v0 is the Minecraft version string for protocol 137.
	Version1v2v0 = "1.2.0"
)

// Profile1v1v0 是 Minecraft 1.1.0 的协议配置。
//
// Profile1v1v0 is the protocol profile for Minecraft 1.1.0.
var Profile1v1v0 = NewProfile(Protocol1v1v0, Version1v1v0)

const (
	// Protocol1v1v0 是 Minecraft 1.1.0 的协议版本。
	//
	// Protocol1v1v0 is the protocol version for Minecraft 1.1.0.
	Protocol1v1v0 int32 = 113
	// Version1v1v0 是协议 113 对应的 Minecraft 版本字符串。
	//
	// Version1v1v0 is the Minecraft version string for protocol 113.
	Version1v1v0 = "1.1.0"
)
