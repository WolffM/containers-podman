package main

import (
	"strings"
	"testing"
)

func composeEnvLookup(env []string, key string) string {
	prefix := key + "="
	for _, entry := range env {
		if strings.HasPrefix(entry, prefix) {
			return strings.TrimPrefix(entry, prefix)
		}
	}
	return ""
}

func TestComposeEnvUsesDockerBuildkitFromEnvironment(t *testing.T) {
	t.Setenv("DOCKER_HOST", "unix:///tmp/podman.sock")
	t.Setenv("DOCKER_BUILDKIT", "1")

	env, err := composeEnv()
	if err != nil {
		t.Fatalf("composeEnv() error = %v", err)
	}

	if got := composeEnvLookup(env, "DOCKER_BUILDKIT"); got != "1" {
		t.Fatalf("expected DOCKER_BUILDKIT to be propagated, got %q", got)
	}
}

func TestComposeEnvDisablesBuildkitByDefault(t *testing.T) {
	t.Setenv("DOCKER_HOST", "unix:///tmp/podman.sock")

	env, err := composeEnv()
	if err != nil {
		t.Fatalf("composeEnv() error = %v", err)
	}

	if got := composeEnvLookup(env, "DOCKER_BUILDKIT"); got != "0" {
		t.Fatalf("expected DOCKER_BUILDKIT to default to 0, got %q", got)
	}
}
