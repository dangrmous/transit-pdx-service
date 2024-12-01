package trimet

import (
	"fmt"
	"io"
	"net/http"
)

// NewClient returns a new instance of Client.
func NewClient() *Client {
	return &Client{}
}

// TrimetClient defines the interface for accessing Trimet data.
type TrimetClient interface {
	GetArrivals(locIDs string) (string, error)
}

// Client implements the TrimetClient interface.
type Client struct{}

// GetArrivals retrieves arrival information for the specified location IDs.
func (c *Client) GetArrivals(locIDs string) (string, error) {
	url := fmt.Sprintf("https://developer.trimet.org/ws/v2/arrivals?locIDs=%s", locIDs)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
