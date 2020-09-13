package version

import "fmt"

var (
	Version   = ""
	GitCommit = ""
	BuildTime = ""
)

func GetVersion() string {
	var sha string
	if GitCommit != "" {
		sha = GitCommit[:7]
	}
	return fmt.Sprintf("%s@%s+%s", Version, BuildTime, sha)
}
