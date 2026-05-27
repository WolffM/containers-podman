//go:build !remote && (linux || freebsd)

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestBuildCancelEndpointRegistered(t *testing.T) {
	for _, path := range []string{"/build/cancel", "/v1.41/build/cancel"} {
		t.Run(path, func(t *testing.T) {
			router := mux.NewRouter()
			server := &APIServer{}

			require.NoError(t, server.registerImagesHandlers(router))

			request := httptest.NewRequest(http.MethodPost, path+"?id=test-build", nil)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			require.Equal(t, http.StatusNoContent, response.Code)
		})
	}
}
