package intistelecom

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer srv.Close()

	client := getClient()
	client.BaseUrl = srv.URL

	os.Exit(m.Run())
}
