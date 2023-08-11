package intistelecom

import "testing"

func TestNewClient(t *testing.T) {
	client := getClient()
	client.UserName = "username"
	client.ApiKey = "key"
	if client.Authorization() != "Basic dXNlcm5hbWU6a2V5" {
		t.Errorf("Invalid credentials: %s", client.Authorization())
	}
}
