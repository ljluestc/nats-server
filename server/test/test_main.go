//go:build !someotherbuild
// +build !someotherbuild

package test

import (
	"os"
	"testing"
)

// TestMain is the main entry point for testing
func TestMain(m *testing.M) {
	// Test setup code

	// Run tests
	code := m.Run()

	// Test teardown code

	os.Exit(code)
}
