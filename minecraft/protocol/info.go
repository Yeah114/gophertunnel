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

// Version returns the Minecraft version string registered for the protocol version.
func (i Info) Version() string {
	if name, ok := Pool[i]; ok {
		return name
	}
	return "unknown"
}

// Current Protocol and Version
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

// Minecraft 1.21.130.28
// Info1v21v130v28 is the protocol info for Minecraft 1.21.130.28.
var Info1v21v130v28 = NewInfo(Protocol1v21v130v28, Version1v21v130v28)

const (
	// Protocol1v21v130v28 is the protocol version for Minecraft 1.21.130.28.
	Protocol1v21v130v28 int32 = 897
	// Version1v21v130v28 is the Minecraft version string for protocol 897.
	Version1v21v130v28 = "1.21.130.28"
)

// Minecraft 1.21.124
// Info1v21v124 is the protocol info for Minecraft 1.21.124.
var Info1v21v124 = NewInfo(Protocol1v21v124, Version1v21v124)

const (
	// Protocol1v21v124 is the protocol version for Minecraft 1.21.124.
	Protocol1v21v124 int32 = 860
	// Version1v21v124 is the Minecraft version string for protocol 860.
	Version1v21v124 = "1.21.124"
)

// Minecraft 1.21.120
// Info1v21v120 is the protocol info for Minecraft 1.21.120.
var Info1v21v120 = NewInfo(Protocol1v21v120, Version1v21v120)

const (
	// Protocol1v21v120 is the protocol version for Minecraft 1.21.120.
	Protocol1v21v120 int32 = 859
	// Version1v21v120 is the Minecraft version string for protocol 859.
	Version1v21v120 = "1.21.120"
)

// Minecraft 1.21.110
// Info1v21v110 is the protocol info for Minecraft 1.21.110.
var Info1v21v110 = NewInfo(Protocol1v21v110, Version1v21v110)

const (
	// Protocol1v21v110 is the protocol version for Minecraft 1.21.110.
	Protocol1v21v110 int32 = 844
	// Version1v21v110 is the Minecraft version string for protocol 844.
	Version1v21v110 = "1.21.110"
)

// Minecraft 1.21.110.26
// Info1v21v110v26 is the protocol info for Minecraft 1.21.110.26.
var Info1v21v110v26 = NewInfo(Protocol1v21v110v26, Version1v21v110v26)

const (
	// Protocol1v21v110v26 is the protocol version for Minecraft 1.21.110.26.
	Protocol1v21v110v26 int32 = 843
	// Version1v21v110v26 is the Minecraft version string for protocol 843.
	Version1v21v110v26 = "1.21.110.26"
)

// Minecraft 1.21.100
// Info1v21v100 is the protocol info for Minecraft 1.21.100.
var Info1v21v100 = NewInfo(Protocol1v21v100, Version1v21v100)

const (
	// Protocol1v21v100 is the protocol version for Minecraft 1.21.100.
	Protocol1v21v100 int32 = 827
	// Version1v21v100 is the Minecraft version string for protocol 827.
	Version1v21v100 = "1.21.100"
)

// Minecraft 1.21.93
// Info1v21v93 is the protocol info for Minecraft 1.21.93.
var Info1v21v93 = NewInfo(Protocol1v21v93, Version1v21v93)

const (
	// Protocol1v21v93 is the protocol version for Minecraft 1.21.93.
	Protocol1v21v93 int32 = 819
	// Version1v21v93 is the Minecraft version string for protocol 819.
	Version1v21v93 = "1.21.93"
)

// Minecraft 1.21.90
// Info1v21v90 is the protocol info for Minecraft 1.21.90.
var Info1v21v90 = NewInfo(Protocol1v21v90, Version1v21v90)

const (
	// Protocol1v21v90 is the protocol version for Minecraft 1.21.90.
	Protocol1v21v90 int32 = 818
	// Version1v21v90 is the Minecraft version string for protocol 818.
	Version1v21v90 = "1.21.90"
)

// Minecraft 1.21.80
// Info1v21v80 is the protocol info for Minecraft 1.21.80.
var Info1v21v80 = NewInfo(Protocol1v21v80, Version1v21v80)

const (
	// Protocol1v21v80 is the protocol version for Minecraft 1.21.80.
	Protocol1v21v80 int32 = 800
	// Version1v21v80 is the Minecraft version string for protocol 800.
	Version1v21v80 = "1.21.80"
)

// Minecraft 1.21.70
// Info1v21v70 is the protocol info for Minecraft 1.21.70.
var Info1v21v70 = NewInfo(Protocol1v21v70, Version1v21v70)

const (
	// Protocol1v21v70 is the protocol version for Minecraft 1.21.70.
	Protocol1v21v70 int32 = 786
	// Version1v21v70 is the Minecraft version string for protocol 786.
	Version1v21v70 = "1.21.70"
)

// Minecraft 1.21.70.24
// Info1v21v70v24 is the protocol info for Minecraft 1.21.70.24.
var Info1v21v70v24 = NewInfo(Protocol1v21v70v24, Version1v21v70v24)

const (
	// Protocol1v21v70v24 is the protocol version for Minecraft 1.21.70.24.
	Protocol1v21v70v24 int32 = 785
	// Version1v21v70v24 is the Minecraft version string for protocol 785.
	Version1v21v70v24 = "1.21.70.24"
)

// Minecraft 1.21.60
// Info1v21v60 is the protocol info for Minecraft 1.21.60.
var Info1v21v60 = NewInfo(Protocol1v21v60, Version1v21v60)

const (
	// Protocol1v21v60 is the protocol version for Minecraft 1.21.60.
	Protocol1v21v60 int32 = 776
	// Version1v21v60 is the Minecraft version string for protocol 776.
	Version1v21v60 = "1.21.60"
)

// Minecraft 1.21.50
// Info1v21v50 is the protocol info for Minecraft 1.21.50.
var Info1v21v50 = NewInfo(Protocol1v21v50, Version1v21v50)

const (
	// Protocol1v21v50 is the protocol version for Minecraft 1.21.50.
	Protocol1v21v50 int32 = 766
	// Version1v21v50 is the Minecraft version string for protocol 766.
	Version1v21v50 = "1.21.50"
)

// Minecraft 1.21.50.26
// Info1v21v50v26 is the protocol info for Minecraft 1.21.50.26.
var Info1v21v50v26 = NewInfo(Protocol1v21v50v26, Version1v21v50v26)

const (
	// Protocol1v21v50v26 is the protocol version for Minecraft 1.21.50.26.
	Protocol1v21v50v26 int32 = 765
	// Version1v21v50v26 is the Minecraft version string for protocol 765.
	Version1v21v50v26 = "1.21.50.26"
)

// Minecraft 1.21.40
// Info1v21v40 is the protocol info for Minecraft 1.21.40.
var Info1v21v40 = NewInfo(Protocol1v21v40, Version1v21v40)

const (
	// Protocol1v21v40 is the protocol version for Minecraft 1.21.40.
	Protocol1v21v40 int32 = 748
	// Version1v21v40 is the Minecraft version string for protocol 748.
	Version1v21v40 = "1.21.40"
)

// Minecraft 1.21.30
// Info1v21v30 is the protocol info for Minecraft 1.21.30.
var Info1v21v30 = NewInfo(Protocol1v21v30, Version1v21v30)

const (
	// Protocol1v21v30 is the protocol version for Minecraft 1.21.30.
	Protocol1v21v30 int32 = 729
	// Version1v21v30 is the Minecraft version string for protocol 729.
	Version1v21v30 = "1.21.30"
)

// Minecraft 1.21.20
// Info1v21v20 is the protocol info for Minecraft 1.21.20.
var Info1v21v20 = NewInfo(Protocol1v21v20, Version1v21v20)

const (
	// Protocol1v21v20 is the protocol version for Minecraft 1.21.20.
	Protocol1v21v20 int32 = 712
	// Version1v21v20 is the Minecraft version string for protocol 712.
	Version1v21v20 = "1.21.20"
)

// Minecraft 1.21.2
// Info1v21v2 is the protocol info for Minecraft 1.21.2.
var Info1v21v2 = NewInfo(Protocol1v21v2, Version1v21v2)

const (
	// Protocol1v21v2 is the protocol version for Minecraft 1.21.2.
	Protocol1v21v2 int32 = 686
	// Version1v21v2 is the Minecraft version string for protocol 686.
	Version1v21v2 = "1.21.2"
)

// Minecraft 1.21.0
// Info1v21v0 is the protocol info for Minecraft 1.21.0.
var Info1v21v0 = NewInfo(Protocol1v21v0, Version1v21v0)

const (
	// Protocol1v21v0 is the protocol version for Minecraft 1.21.0.
	Protocol1v21v0 int32 = 685
	// Version1v21v0 is the Minecraft version string for protocol 685.
	Version1v21v0 = "1.21.0"
)

// Minecraft 1.20.80
// Info1v20v80 is the protocol info for Minecraft 1.20.80.
var Info1v20v80 = NewInfo(Protocol1v20v80, Version1v20v80)

const (
	// Protocol1v20v80 is the protocol version for Minecraft 1.20.80.
	Protocol1v20v80 int32 = 671
	// Version1v20v80 is the Minecraft version string for protocol 671.
	Version1v20v80 = "1.20.80"
)

// Minecraft 1.20.70
// Info1v20v70 is the protocol info for Minecraft 1.20.70.
var Info1v20v70 = NewInfo(Protocol1v20v70, Version1v20v70)

const (
	// Protocol1v20v70 is the protocol version for Minecraft 1.20.70.
	Protocol1v20v70 int32 = 662
	// Version1v20v70 is the Minecraft version string for protocol 662.
	Version1v20v70 = "1.20.70"
)

// Minecraft 1.20.60
// Info1v20v60 is the protocol info for Minecraft 1.20.60.
var Info1v20v60 = NewInfo(Protocol1v20v60, Version1v20v60)

const (
	// Protocol1v20v60 is the protocol version for Minecraft 1.20.60.
	Protocol1v20v60 int32 = 649
	// Version1v20v60 is the Minecraft version string for protocol 649.
	Version1v20v60 = "1.20.60"
)

// Minecraft 1.20.50
// Info1v20v50 is the protocol info for Minecraft 1.20.50.
var Info1v20v50 = NewInfo(Protocol1v20v50, Version1v20v50)

const (
	// Protocol1v20v50 is the protocol version for Minecraft 1.20.50.
	Protocol1v20v50 int32 = 630
	// Version1v20v50 is the Minecraft version string for protocol 630.
	Version1v20v50 = "1.20.50"
)

// Minecraft 1.20.40
// Info1v20v40 is the protocol info for Minecraft 1.20.40.
var Info1v20v40 = NewInfo(Protocol1v20v40, Version1v20v40)

const (
	// Protocol1v20v40 is the protocol version for Minecraft 1.20.40.
	Protocol1v20v40 int32 = 622
	// Version1v20v40 is the Minecraft version string for protocol 622.
	Version1v20v40 = "1.20.40"
)

// Minecraft 1.20.30
// Info1v20v30 is the protocol info for Minecraft 1.20.30.
var Info1v20v30 = NewInfo(Protocol1v20v30, Version1v20v30)

const (
	// Protocol1v20v30 is the protocol version for Minecraft 1.20.30.
	Protocol1v20v30 int32 = 618
	// Version1v20v30 is the Minecraft version string for protocol 618.
	Version1v20v30 = "1.20.30"
)

// Minecraft 1.20.30.24
// Info1v20v30v24 is the protocol info for Minecraft 1.20.30.24.
var Info1v20v30v24 = NewInfo(Protocol1v20v30v24, Version1v20v30v24)

const (
	// Protocol1v20v30v24 is the protocol version for Minecraft 1.20.30.24.
	Protocol1v20v30v24 int32 = 617
	// Version1v20v30v24 is the Minecraft version string for protocol 617.
	Version1v20v30v24 = "1.20.30.24"
)

// Minecraft 1.20.10
// Info1v20v10 is the protocol info for Minecraft 1.20.10.
var Info1v20v10 = NewInfo(Protocol1v20v10, Version1v20v10)

const (
	// Protocol1v20v10 is the protocol version for Minecraft 1.20.10.
	Protocol1v20v10 int32 = 594
	// Version1v20v10 is the Minecraft version string for protocol 594.
	Version1v20v10 = "1.20.10"
)

// Minecraft 1.20.0
// Info1v20v0 is the protocol info for Minecraft 1.20.0.
var Info1v20v0 = NewInfo(Protocol1v20v0, Version1v20v0)

const (
	// Protocol1v20v0 is the protocol version for Minecraft 1.20.0.
	Protocol1v20v0 int32 = 589
	// Version1v20v0 is the Minecraft version string for protocol 589.
	Version1v20v0 = "1.20.0"
)

// Minecraft 1.20.0.23
// Info1v20v0v23 is the protocol info for Minecraft 1.20.0.23.
var Info1v20v0v23 = NewInfo(Protocol1v20v0v23, Version1v20v0v23)

const (
	// Protocol1v20v0v23 is the protocol version for Minecraft 1.20.0.23.
	Protocol1v20v0v23 int32 = 588
	// Version1v20v0v23 is the Minecraft version string for protocol 588.
	Version1v20v0v23 = "1.20.0.23"
)

// Minecraft 1.19.80
// Info1v19v80 is the protocol info for Minecraft 1.19.80.
var Info1v19v80 = NewInfo(Protocol1v19v80, Version1v19v80)

const (
	// Protocol1v19v80 is the protocol version for Minecraft 1.19.80.
	Protocol1v19v80 int32 = 582
	// Version1v19v80 is the Minecraft version string for protocol 582.
	Version1v19v80 = "1.19.80"
)

// Minecraft 1.19.70
// Info1v19v70 is the protocol info for Minecraft 1.19.70.
var Info1v19v70 = NewInfo(Protocol1v19v70, Version1v19v70)

const (
	// Protocol1v19v70 is the protocol version for Minecraft 1.19.70.
	Protocol1v19v70 int32 = 575
	// Version1v19v70 is the Minecraft version string for protocol 575.
	Version1v19v70 = "1.19.70"
)

// Minecraft 1.19.70.24
// Info1v19v70v24 is the protocol info for Minecraft 1.19.70.24.
var Info1v19v70v24 = NewInfo(Protocol1v19v70v24, Version1v19v70v24)

const (
	// Protocol1v19v70v24 is the protocol version for Minecraft 1.19.70.24.
	Protocol1v19v70v24 int32 = 574
	// Version1v19v70v24 is the Minecraft version string for protocol 574.
	Version1v19v70v24 = "1.19.70.24"
)

// Minecraft 1.19.63
// Info1v19v63 is the protocol info for Minecraft 1.19.63.
var Info1v19v63 = NewInfo(Protocol1v19v63, Version1v19v63)

const (
	// Protocol1v19v63 is the protocol version for Minecraft 1.19.63.
	Protocol1v19v63 int32 = 568
	// Version1v19v63 is the Minecraft version string for protocol 568.
	Version1v19v63 = "1.19.63"
)

// Minecraft 1.19.60
// Info1v19v60 is the protocol info for Minecraft 1.19.60.
var Info1v19v60 = NewInfo(Protocol1v19v60, Version1v19v60)

const (
	// Protocol1v19v60 is the protocol version for Minecraft 1.19.60.
	Protocol1v19v60 int32 = 567
	// Version1v19v60 is the Minecraft version string for protocol 567.
	Version1v19v60 = "1.19.60"
)

// Minecraft 1.19.50
// Info1v19v50 is the protocol info for Minecraft 1.19.50.
var Info1v19v50 = NewInfo(Protocol1v19v50, Version1v19v50)

const (
	// Protocol1v19v50 is the protocol version for Minecraft 1.19.50.
	Protocol1v19v50 int32 = 560
	// Version1v19v50 is the Minecraft version string for protocol 560.
	Version1v19v50 = "1.19.50"
)

// Minecraft 1.19.40
// Info1v19v40 is the protocol info for Minecraft 1.19.40.
var Info1v19v40 = NewInfo(Protocol1v19v40, Version1v19v40)

const (
	// Protocol1v19v40 is the protocol version for Minecraft 1.19.40.
	Protocol1v19v40 int32 = 557
	// Version1v19v40 is the Minecraft version string for protocol 557.
	Version1v19v40 = "1.19.40"
)

// Minecraft 1.19.30
// Info1v19v30 is the protocol info for Minecraft 1.19.30.
var Info1v19v30 = NewInfo(Protocol1v19v30, Version1v19v30)

const (
	// Protocol1v19v30 is the protocol version for Minecraft 1.19.30.
	Protocol1v19v30 int32 = 554
	// Version1v19v30 is the Minecraft version string for protocol 554.
	Version1v19v30 = "1.19.30"
)

// Minecraft 1.19.21
// Info1v19v21 is the protocol info for Minecraft 1.19.21.
var Info1v19v21 = NewInfo(Protocol1v19v21, Version1v19v21)

const (
	// Protocol1v19v21 is the protocol version for Minecraft 1.19.21.
	Protocol1v19v21 int32 = 545
	// Version1v19v21 is the Minecraft version string for protocol 545.
	Version1v19v21 = "1.19.21"
)

// Minecraft 1.19.20
// Info1v19v20 is the protocol info for Minecraft 1.19.20.
var Info1v19v20 = NewInfo(Protocol1v19v20, Version1v19v20)

const (
	// Protocol1v19v20 is the protocol version for Minecraft 1.19.20.
	Protocol1v19v20 int32 = 544
	// Version1v19v20 is the Minecraft version string for protocol 544.
	Version1v19v20 = "1.19.20"
)

// Minecraft 1.19.10
// Info1v19v10 is the protocol info for Minecraft 1.19.10.
var Info1v19v10 = NewInfo(Protocol1v19v10, Version1v19v10)

const (
	// Protocol1v19v10 is the protocol version for Minecraft 1.19.10.
	Protocol1v19v10 int32 = 534
	// Version1v19v10 is the Minecraft version string for protocol 534.
	Version1v19v10 = "1.19.10"
)

// Minecraft 1.19.0
// Info1v19v0 is the protocol info for Minecraft 1.19.0.
var Info1v19v0 = NewInfo(Protocol1v19v0, Version1v19v0)

const (
	// Protocol1v19v0 is the protocol version for Minecraft 1.19.0.
	Protocol1v19v0 int32 = 527
	// Version1v19v0 is the Minecraft version string for protocol 527.
	Version1v19v0 = "1.19.0"
)

// Minecraft 1.18.30
// Info1v18v30 is the protocol info for Minecraft 1.18.30.
var Info1v18v30 = NewInfo(Protocol1v18v30, Version1v18v30)

const (
	// Protocol1v18v30 is the protocol version for Minecraft 1.18.30.
	Protocol1v18v30 int32 = 503
	// Version1v18v30 is the Minecraft version string for protocol 503.
	Version1v18v30 = "1.18.30"
)

// Minecraft 1.18.10
// Info1v18v10 is the protocol info for Minecraft 1.18.10.
var Info1v18v10 = NewInfo(Protocol1v18v10, Version1v18v10)

const (
	// Protocol1v18v10 is the protocol version for Minecraft 1.18.10.
	Protocol1v18v10 int32 = 486
	// Version1v18v10 is the Minecraft version string for protocol 486.
	Version1v18v10 = "1.18.10"
)

// Minecraft 1.18.0
// Info1v18v0 is the protocol info for Minecraft 1.18.0.
var Info1v18v0 = NewInfo(Protocol1v18v0, Version1v18v0)

const (
	// Protocol1v18v0 is the protocol version for Minecraft 1.18.0.
	Protocol1v18v0 int32 = 475
	// Version1v18v0 is the Minecraft version string for protocol 475.
	Version1v18v0 = "1.18.0"
)

// Minecraft 1.17.40
// Info1v17v40 is the protocol info for Minecraft 1.17.40.
var Info1v17v40 = NewInfo(Protocol1v17v40, Version1v17v40)

const (
	// Protocol1v17v40 is the protocol version for Minecraft 1.17.40.
	Protocol1v17v40 int32 = 471
	// Version1v17v40 is the Minecraft version string for protocol 471.
	Version1v17v40 = "1.17.40"
)

// Minecraft 1.17.30
// Info1v17v30 is the protocol info for Minecraft 1.17.30.
var Info1v17v30 = NewInfo(Protocol1v17v30, Version1v17v30)

const (
	// Protocol1v17v30 is the protocol version for Minecraft 1.17.30.
	Protocol1v17v30 int32 = 465
	// Version1v17v30 is the Minecraft version string for protocol 465.
	Version1v17v30 = "1.17.30"
)

// Minecraft 1.17.20.20
// Info1v17v20v20 is the protocol info for Minecraft 1.17.20.20.
var Info1v17v20v20 = NewInfo(Protocol1v17v20v20, Version1v17v20v20)

const (
	// Protocol1v17v20v20 is the protocol version for Minecraft 1.17.20.20.
	Protocol1v17v20v20 int32 = 453
	// Version1v17v20v20 is the Minecraft version string for protocol 453.
	Version1v17v20v20 = "1.17.20.20"
)

// Minecraft 1.17.10
// Info1v17v10 is the protocol info for Minecraft 1.17.10.
var Info1v17v10 = NewInfo(Protocol1v17v10, Version1v17v10)

const (
	// Protocol1v17v10 is the protocol version for Minecraft 1.17.10.
	Protocol1v17v10 int32 = 448
	// Version1v17v10 is the Minecraft version string for protocol 448.
	Version1v17v10 = "1.17.10"
)

// Minecraft 1.17.0
// Info1v17v0 is the protocol info for Minecraft 1.17.0.
var Info1v17v0 = NewInfo(Protocol1v17v0, Version1v17v0)

const (
	// Protocol1v17v0 is the protocol version for Minecraft 1.17.0.
	Protocol1v17v0 int32 = 440
	// Version1v17v0 is the Minecraft version string for protocol 440.
	Version1v17v0 = "1.17.0"
)

// Minecraft 1.16.230.54
// Info1v16v230v54 is the protocol info for Minecraft 1.16.230.54.
var Info1v16v230v54 = NewInfo(Protocol1v16v230v54, Version1v16v230v54)

const (
	// Protocol1v16v230v54 is the protocol version for Minecraft 1.16.230.54.
	Protocol1v16v230v54 int32 = 435
	// Version1v16v230v54 is the Minecraft version string for protocol 435.
	Version1v16v230v54 = "1.16.230.54"
)

// Minecraft 1.16.230
// Info1v16v230 is the protocol info for Minecraft 1.16.230.
var Info1v16v230 = NewInfo(Protocol1v16v230, Version1v16v230)

const (
	// Protocol1v16v230 is the protocol version for Minecraft 1.16.230.
	Protocol1v16v230 int32 = 434
	// Version1v16v230 is the Minecraft version string for protocol 434.
	Version1v16v230 = "1.16.230"
)

// Minecraft 1.16.230.50
// Info1v16v230v50 is the protocol info for Minecraft 1.16.230.50.
var Info1v16v230v50 = NewInfo(Protocol1v16v230v50, Version1v16v230v50)

const (
	// Protocol1v16v230v50 is the protocol version for Minecraft 1.16.230.50.
	Protocol1v16v230v50 int32 = 433
	// Version1v16v230v50 is the Minecraft version string for protocol 433.
	Version1v16v230v50 = "1.16.230.50"
)

// Minecraft 1.16.220
// Info1v16v220 is the protocol info for Minecraft 1.16.220.
var Info1v16v220 = NewInfo(Protocol1v16v220, Version1v16v220)

const (
	// Protocol1v16v220 is the protocol version for Minecraft 1.16.220.
	Protocol1v16v220 int32 = 431
	// Version1v16v220 is the Minecraft version string for protocol 431.
	Version1v16v220 = "1.16.220"
)

// Minecraft 1.16.210
// Info1v16v210 is the protocol info for Minecraft 1.16.210.
var Info1v16v210 = NewInfo(Protocol1v16v210, Version1v16v210)

const (
	// Protocol1v16v210 is the protocol version for Minecraft 1.16.210.
	Protocol1v16v210 int32 = 428
	// Version1v16v210 is the Minecraft version string for protocol 428.
	Version1v16v210 = "1.16.210"
)

// Minecraft 1.16.210.50
// Info1v16v210v50 is the protocol info for Minecraft 1.16.210.50.
var Info1v16v210v50 = NewInfo(Protocol1v16v210v50, Version1v16v210v50)

const (
	// Protocol1v16v210v50 is the protocol version for Minecraft 1.16.210.50.
	Protocol1v16v210v50 int32 = 423
	// Version1v16v210v50 is the Minecraft version string for protocol 423.
	Version1v16v210v50 = "1.16.210.50"
)

// Minecraft 1.16.200
// Info1v16v200 is the protocol info for Minecraft 1.16.200.
var Info1v16v200 = NewInfo(Protocol1v16v200, Version1v16v200)

const (
	// Protocol1v16v200 is the protocol version for Minecraft 1.16.200.
	Protocol1v16v200 int32 = 422
	// Version1v16v200 is the Minecraft version string for protocol 422.
	Version1v16v200 = "1.16.200"
)

// Minecraft 1.16.200.51
// Info1v16v200v51 is the protocol info for Minecraft 1.16.200.51.
var Info1v16v200v51 = NewInfo(Protocol1v16v200v51, Version1v16v200v51)

const (
	// Protocol1v16v200v51 is the protocol version for Minecraft 1.16.200.51.
	Protocol1v16v200v51 int32 = 420
	// Version1v16v200v51 is the Minecraft version string for protocol 420.
	Version1v16v200v51 = "1.16.200.51"
)

// Minecraft 1.16.100
// Info1v16v100 is the protocol info for Minecraft 1.16.100.
var Info1v16v100 = NewInfo(Protocol1v16v100, Version1v16v100)

const (
	// Protocol1v16v100 is the protocol version for Minecraft 1.16.100.
	Protocol1v16v100 int32 = 419
	// Version1v16v100 is the Minecraft version string for protocol 419.
	Version1v16v100 = "1.16.100"
)
