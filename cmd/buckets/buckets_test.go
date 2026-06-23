package buckets_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/cmd/buckets"
	"github.com/mic615/chill-crate-cli/internal/client"
	"github.com/mic615/chill-crate-cli/internal/testutil"
)

const testGroupID = "group-1"

// newRoot builds a minimal cobra tree for each test so tests don't share state.
func newRoot() *cobra.Command {
	root := &cobra.Command{Use: "chill"}
	root.AddCommand(buckets.BucketsCmd())
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

// --- buckets create ---

func TestBucketCreate_success(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc("POST /buckets", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(client.Bucket{Name: "photos", ID: "bucket-1"})
	})

	out, err := run(t, "buckets", "create", "photos")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "photos") {
		t.Errorf("expected output to contain bucket name, got: %q", out)
	}
}

func TestBucketCreate_noGroup(t *testing.T) {
	testutil.NewAPIServer(t, "" /* no group */)
	viper.Set("current_group_id", "")

	_, err := run(t, "buckets", "create", "photos")
	if err == nil {
		t.Fatal("expected error when no group is set")
	}
	if !strings.Contains(err.Error(), "no group selected") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBucketCreate_apiError(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc("POST /buckets", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "bucket already exists"})
	})

	_, err := run(t, "buckets", "create", "photos")
	if err == nil {
		t.Fatal("expected error from API")
	}
	if !strings.Contains(err.Error(), "bucket already exists") {
		t.Errorf("unexpected error: %v", err)
	}
}

// --- buckets list ---

func TestBucketList_success(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc(
		"GET /groups/{groupID}/buckets",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode([]client.Bucket{
				{Name: "photos", ID: "b-1"},
				{Name: "videos", ID: "b-2"},
			})
		},
	)

	out, err := run(t, "buckets", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "photos") || !strings.Contains(out, "videos") {
		t.Errorf("expected both bucket names in output, got: %q", out)
	}
}

func TestBucketList_empty(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc(
		"GET /groups/{groupID}/buckets",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode([]client.Bucket{})
		},
	)

	out, err := run(t, "buckets", "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "No buckets yet") {
		t.Errorf("expected empty-state message, got: %q", out)
	}
}

func TestBucketList_noGroup(t *testing.T) {
	testutil.NewAPIServer(t, "")
	viper.Set("current_group_id", "")

	_, err := run(t, "buckets", "list")
	if err == nil {
		t.Fatal("expected error when no group is set")
	}
	if !strings.Contains(err.Error(), "no group selected") {
		t.Errorf("unexpected error: %v", err)
	}
}

// --- buckets delete (error paths that don't reach the interactive prompt) ---

func TestBucketDelete_noGroup(t *testing.T) {
	testutil.NewAPIServer(t, "")
	viper.Set("current_group_id", "")

	_, err := run(t, "buckets", "delete", "photos")
	if err == nil {
		t.Fatal("expected error when no group is set")
	}
	if !strings.Contains(err.Error(), "no group selected") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBucketDelete_bucketNotFound(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc(
		"GET /groups/{groupID}/buckets/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "bucket not found"})
		},
	)

	_, err := run(t, "buckets", "delete", "ghost")
	if err == nil {
		t.Fatal("expected error when bucket does not exist")
	}
	if !strings.Contains(err.Error(), "bucket not found") {
		t.Errorf("unexpected error: %v", err)
	}
}
