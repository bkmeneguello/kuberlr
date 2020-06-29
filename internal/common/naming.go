package common

import (
	"fmt"

	"github.com/blang/semver"
)

// KubectlLocalNamingScheme holds the scheme used to name the kubectl binaries
// downloaded by kuberlr
const KubectlLocalNamingScheme = "kubectl-%d.%d.%d"

// KubectlSystemNamingScheme holds the scheme used to name the kubectl binaries
// installed system-wide
const KubectlSystemNamingScheme = "kubectl-%d.%d"

// BuildKubectlNameForLocalBin returns how kuberlr will name the kubectl binary
// with the specified version when downloading that to the user home
func BuildKubectlNameForLocalBin(v semver.Version) string {
	return fmt.Sprintf(KubectlLocalNamingScheme, v.Major, v.Minor, v.Patch)
}

// BuildKubectlNameForLocalBin returns how kuberlr expects system-wide
// kubectl binaries to be named
func BuildKubectlNameForSystemBin(version semver.Version) string {
	return fmt.Sprintf(KubectlSystemNamingScheme, version.Major, version.Minor)
}
