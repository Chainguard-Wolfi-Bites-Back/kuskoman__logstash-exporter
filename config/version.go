package config

import (
	"fmt"
	"runtime"
)

var (
	// Version is the current version of Logstash Exporter.
	Version = "unknown"

	// GitCommit is the git commit hash of the current build.
	GitCommit = "unknown"

	// BuildDate is the date of the current build.
	BuildDate = "unknown"
)

// GetBuildInfo returns a VersionInfo struct with the current build information.
func GetBuildInfo() *VersionInfo {
	return &VersionInfo{
		Version:   Version,
		GitCommit: GitCommit,
		GoVersion: runtime.Version(),
		BuildArch: runtime.GOARCH,
		BuildOS:   runtime.GOOS,
		BuildDate: BuildDate,
	}
}

// VersionInfo contains the current build information.
type VersionInfo struct {
	Version   string
	GitCommit string
	GoVersion string
	BuildArch string
	BuildOS   string
	BuildDate string
}

// String returns a string representation of the VersionInfo struct.
func (v *VersionInfo) String() string {
	return fmt.Sprintf("%+v", *v)
}
