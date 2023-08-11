package intistelecom

var client *Client

func init() {
	if client == nil {
		client = NewClient()
	}
}
