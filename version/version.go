package framework

import (
	"fmt"
	"strconv"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func (self Version) Undefined() bool {
	return (self.Major == 0 && self.Minor == 0 && self.Patch == 0)
}

func (self Version) OlderThan(v Version) bool {
	return (self.Major < v.Major || (self.Major == v.Major && self.Minor < v.Minor) ||
		(self.Major == v.Major && self.Minor == v.Minor && self.Patch < v.Patch))
}

func (self Version) NewerThan(v Version) bool {
	return (self.Major > v.Major || (self.Major == v.Major && self.Minor > v.Minor) ||
		(self.Major == v.Major && self.Minor == v.Minor && self.Patch > v.Patch))
}

func (self Version) String() string {
	return fmt.Sprintf("" + strconv.Itoa(self.Major) + "." + strconv.Itoa(self.Minor) + "." + strconv.Itoa(self.Patch))
}
