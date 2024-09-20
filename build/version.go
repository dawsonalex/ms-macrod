package build

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed version
var Version string

var (
	Commit      = "939jf93k92je"
	Branch      = "main"
	Host        = "main"
	Environment = "local"
)

type VersionInfo struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
	Patch string `json:"patch"`
}

func (v VersionInfo) Sprint() string {
	return fmt.Sprintf("%s.%s.%s", v.Major, v.Minor, v.Patch)
}

type BuildInfo struct {
	Version     VersionInfo `json:"version"`
	Commit      string      `json:"commit"`
	Branch      string      `json:"branch"`
	Host        string      `json:"host"`
	Environment string      `json:"environment"`
}

func Info() BuildInfo {
	splitVersion := strings.Split(Version, ".")

	return BuildInfo{
		Version: VersionInfo{
			Major: splitVersion[0],
			Minor: splitVersion[1],
			Patch: splitVersion[2],
		},
		Commit:      Commit,
		Branch:      Branch,
		Host:        Host,
		Environment: Environment,
	}
}
