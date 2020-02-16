package common

import (
	"fmt"
	"runtime"
	"strings"
)

type Version struct {
	Major int

	Minor int

	// Increment this for bug releases
	Patch int
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func BuildCurrentVersionString(short bool) string {
	version := "v" + CurrentVersion.String()

	if short {
		return version
	}

	osArch := runtime.GOOS + "/" + runtime.GOARCH

	return fmt.Sprintf("%s (%s)", version, osArch)

}

func GetUserAgent() string {
	return "dtags/" + strings.TrimPrefix(BuildCurrentVersionString(true), "v")
}
