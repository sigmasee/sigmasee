// Package util implements different utilities
package util

import "runtime"

// Set by linker
var (
	version  = "undefined"
	platform = "undefined"
	commit   = "undefined"
	date     = "undefined"
)

// Version defines the structure containns all information to be printed when 'version' command is requested.
type Version struct {
	Version       string
	Platform      string
	Commit        string
	Date          string
	GolangVersion string
}

// GetVersion returns the version information
// Returns the version inforamtion
func GetVersion() Version {
	return Version{
		Version:       version,
		Platform:      platform,
		Commit:        commit,
		Date:          date,
		GolangVersion: runtime.Version(),
	}
}
