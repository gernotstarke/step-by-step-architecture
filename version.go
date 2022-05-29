package version

import "time"

const VersionStr = "0.0.4b"
const status = "pre-alpha"

func Status() string {
	return status + " - " + time.Now().Format("2. Jan 2006")
}
