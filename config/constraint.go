package config

import "time"

const (
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	AccessTokenExpiresInTime time.Duration = 1 * 60 * 24 * time.Minute
)
