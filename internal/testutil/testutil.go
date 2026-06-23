package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

// Server wraps httptest.Server with a mux for registering per-test handlers.
type Server struct {
	*httptest.Server
	Mux *http.ServeMux
}

// NewAPIServer starts a fake HTTP API server, points viper's api_url at it,
// and sets current_group_id to groupID. Both are reset when the test ends.
func NewAPIServer(t *testing.T, groupID string) *Server {
	t.Helper()
	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	t.Cleanup(viper.Reset)
	viper.Set("api_url", srv.URL)
	viper.Set("current_group_id", groupID)
	return &Server{Server: srv, Mux: mux}
}
