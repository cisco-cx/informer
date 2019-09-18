package example

import "runtime"

// Information set at build-time.
var (
	Program   = "none"
	License   = "none"
	URL       = "none"
	BuildUser = "none"
	BuildDate = "none"
	GoVersion = runtime.Version()
	Version   = "none"
	Revision  = "none"
	Branch    = "none"
)
