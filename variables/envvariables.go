package variables

import "os"

// Init - Set env variables
func Init() {
	os.Setenv(
		"globeHost", "http://google.com",
	)
	os.Setenv(
		"timeoutGlobe", "5s",
	)
	os.Setenv(
		"redisHost", "localhost:6379",
	)
}
