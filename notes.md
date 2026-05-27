## Steps to reproduce
1. Added a focused unit test in `/tmp/workspace/WolffM/containers-podman/cmd/podman/compose_test.go` that sets `DOCKER_HOST` and `DOCKER_BUILDKIT=1`, then calls `composeEnv()`.
2. Ran: `CGO_ENABLED=0 go test -tags "exclude_graphdriver_btrfs containers_image_openpgp" ./cmd/podman -run TestComposeEnvUsesDockerBuildkitFromEnvironment -count=1`.
3. Confirmed the test fails against the previous implementation before the fix.

## Observed
The command failed with: `expected DOCKER_BUILDKIT to be propagated, got "0"`. This trace shows that `podman compose` always forced `DOCKER_BUILDKIT=0` even when the user explicitly exported `DOCKER_BUILDKIT=1`. As a result, compose providers cannot use BuildKit-specific features even when the user opts in.

## Expected
When the user sets `DOCKER_BUILDKIT` before running `podman compose`, that value should be passed through to the external compose provider. The default should remain `DOCKER_BUILDKIT=0` when unset, but users should have a supported opt-in path to enable BuildKit features for local development and debugging workflows.
