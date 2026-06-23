package groups_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/cmd/groups"
	"github.com/mic615/chill-crate-cli/internal/client"
	"github.com/mic615/chill-crate-cli/internal/testutil"
)

const testGroupID = "group-1"

// newRoot builds a minimal cobra tree for each test so tests don't share state.
func newRoot() *cobra.Command {
	root := &cobra.Command{Use: "chill"}
	root.AddCommand(groups.GroupsCmd())
	return root
}

// run executes the given args against a fresh root and returns combined output.
func run(t *testing.T, args ...string) (string, error) {
	t.Helper()
	var buf bytes.Buffer
	root := newRoot()
	root.SetOut(&buf)
	root.SetErr(&buf)
	root.SetArgs(args)
	_, err := root.ExecuteC()
	return buf.String(), err
}

// --- groups create ---

func TestGroupCreate_success(t *testing.T) {
	srv := testutil.NewAPIServer(t, "")
	srv.Mux.HandleFunc("POST /groups", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(client.Group{Name: "my-group", ID: "group-1"})
	})

	out, err := run(t, "groups", "create", "my-group")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "my-group") {
		t.Errorf("expected output to contain group name, got: %q", out)
	}
}

func TestGroupCreate_apiError(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc("POST /groups", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "group already exists"})
	})

	_, err := run(t, "groups", "create", "my-group")
	if err == nil {
		t.Fatal("expected error from API")
	}
	if !strings.Contains(err.Error(), "group already exists") {
		t.Errorf("unexpected error: %v", err)
	}
}

// --- groups list ---

func TestGroupList_success(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc("GET /groups", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]client.Group{
			{Name: "my-group", ID: "group-1"},
		})
	})

	out, err := run(t, "groups", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "my-group") {
		t.Errorf("expected output to contain group name, got: %q", out)
	}
}

func TestGroupList_empty(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc("GET /groups", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]client.Group{})
	})

	out, err := run(t, "groups", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "No groups yet") {
		t.Errorf("expected empty-state message, got: %q", out)
	}
}

func TestGroupList_currentMarker(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc("GET /groups", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]client.Group{
			{Name: "active-group", ID: testGroupID},
			{Name: "other-group", ID: "group-2"},
		})
	})

	out, err := run(t, "groups", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines of output, got %d: %q", len(lines), out)
	}
	if !strings.HasPrefix(lines[0], "*") {
		t.Errorf("expected active group to have '*' marker, got: %q", lines[0])
	}
	if strings.HasPrefix(lines[1], "*") {
		t.Errorf("expected inactive group to have no '*' marker, got: %q", lines[1])
	}
}

// --- groups current ---

func TestGroupCurrent_success(t *testing.T) {
	testutil.NewAPIServer(t, testGroupID)
	viper.Set("current_group_name", "my-group")

	out, err := run(t, "groups", "current")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "my-group") {
		t.Errorf("expected output to contain group name, got: %q", out)
	}
}

func TestGroupCurrent_noGroup(t *testing.T) {
	testutil.NewAPIServer(t, "")
	viper.Set("current_group_id", "")

	_, err := run(t, "groups", "current")
	if err == nil {
		t.Fatal("expected error when no group is set")
	}
	if !strings.Contains(err.Error(), "no group selected") {
		t.Errorf("unexpected error: %v", err)
	}
}
