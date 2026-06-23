package objects_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mic615/chill-crate-cli/cmd/objects"
	"github.com/mic615/chill-crate-cli/internal/client"
	"github.com/mic615/chill-crate-cli/internal/testutil"
)

const (
	testGroupID  = "group-1"
	testBucketID = "bucket-1"
)

func newRoot() *cobra.Command {
	root := &cobra.Command{Use: "chill"}
	root.AddCommand(objects.ObjectsCmd())
	return root
}

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

// registerBucketLookup registers the bucket-by-name endpoint used by most object commands.
func registerBucketLookup(srv *testutil.Server, bucketName string) {
	srv.Mux.HandleFunc(
		"GET /groups/{groupID}/buckets/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(client.Bucket{Name: bucketName, ID: testBucketID})
		},
	)
}

// --- objects list ---

func TestObjectList_success(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	registerBucketLookup(srv, "photos")
	srv.Mux.HandleFunc(
		"GET /buckets/{bucketID}/objects",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode([]client.Object{
				{FileName: "cat.jpg", Version: 1},
				{FileName: "dog.jpg", Version: 3},
			})
		},
	)

	out, err := run(t, "objects", "list", "photos")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "cat.jpg") || !strings.Contains(out, "dog.jpg") {
		t.Errorf("expected both file names in output, got: %q", out)
	}
	if !strings.Contains(out, "3") {
		t.Errorf("expected version number in output, got: %q", out)
	}
}

func TestObjectList_empty(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	registerBucketLookup(srv, "photos")
	srv.Mux.HandleFunc(
		"GET /buckets/{bucketID}/objects",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode([]client.Object{})
		},
	)

	out, err := run(t, "objects", "list", "photos")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "No objects yet") {
		t.Errorf("expected empty-state message, got: %q", out)
	}
}

func TestObjectList_noGroup(t *testing.T) {
	testutil.NewAPIServer(t, "")
	viper.Set("current_group_id", "")

	_, err := run(t, "objects", "list", "photos")
	if err == nil {
		t.Fatal("expected error when no group is set")
	}
	if !strings.Contains(err.Error(), "no group selected") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestObjectList_bucketNotFound(t *testing.T) {
	srv := testutil.NewAPIServer(t, testGroupID)
	srv.Mux.HandleFunc(
		"GET /groups/{groupID}/buckets/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "bucket not found"})
		},
	)

	_, err := run(t, "objects", "list", "ghost")
	if err == nil {
		t.Fatal("expected error when bucket does not exist")
	}
	if !strings.Contains(err.Error(), "bucket not found") {
		t.Errorf("unexpected error: %v", err)
	}
}
