package version42

import "time"

const VersionStr = "0.0.4b"

const AppName = "PDFminion"

var Status = "pre-alpha - " + time.Now().Format("2. Jan 2006")

func Version() string {
	return VersionStr + Status
}
