## Steps to reproduce
1. From the repository root, run `CGO_ENABLED=0 go test -tags 'containers_image_openpgp exclude_graphdriver_btrfs' ./pkg/api/server -run TestBuildCancelEndpointRegistered -count=1`.
2. The test sends `POST` requests to `/build/cancel?id=test-build` and `/v1.41/build/cancel?id=test-build` through the registered API router, which mirrors the Docker client calling the BuildKit cancel API.
3. Observe the router response codes before the fix.

## Observed
Both requests returned HTTP 404 instead of reaching a handler. The failing test output showed `expected: 204` and `actual: 404` for `/build/cancel` and `/v1.41/build/cancel`. This matches the issue report where `curl -X POST ... http://v1.41/build/cancel` returns `Not Found`, so a Docker BuildKit client cannot rely on the API being present.

## Expected
The compatibility API should register `/build/cancel` for both versioned and unversioned Docker-style paths so BuildKit clients can call it successfully. The expected result for the compatibility endpoint is an HTTP 204 no-content response instead of a router-level 404, allowing `DOCKER_BUILDKIT=1 docker build` to complete without failing on the missing API path.
