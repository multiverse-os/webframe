package version

import (
	"fmt"
	"strconv"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) Undefined() bool {
	return (v.Major == 0 && v.Minor == 0 && v.Patch == 0)
}

func (v Version) OlderThan(version Version) bool {
	return (v.Major < version.Major || (v.Major == version.Major && v.Minor < version.Minor) ||
		(v.Major == version.Major && v.Minor == version.Minor && v.Patch < version.Patch))
}

func (v Version) NewerThan(version Version) bool {
	return (v.Major > version.Major || (v.Major == version.Major && v.Minor > version.Minor) ||
		(v.Major == version.Major && v.Minor == version.Minor && v.Patch > version.Patch))
}

func (v Version) String() string {
	return fmt.Sprintf("" + strconv.Itoa(v.Major) + "." + strconv.Itoa(v.Minor) + "." + strconv.Itoa(v.Patch))
}
