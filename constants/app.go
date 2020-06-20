// Package constants global config constant
package constants

const (
	ExportConfigPath = "bak_conf/"

	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

var (
	AppModeMap = map[string]string{
		DebugMode:   DebugMode,
		ReleaseMode: ReleaseMode,
		TestMode:    TestMode,
	}
)
