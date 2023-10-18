package notify

import (
	"context"
	"fmt"
	c "github.com/ncostamagna/go_http_client/client"
	"net/http"
	"net/url"
	"time"
)

type (
	DataResponse struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
		Data   string `json:"data"`
	}

	PushReq struct {
		Title   string `json:"title,omitempty"`
		Message string `json:"message,omitempty"`
		Url     string `json:"url,omitempty"`
	}

	Transport interface {
		Push(ctx context.Context, title, message, urlNotify string) error
	}

	clientHTTP struct {
		client c.Transport
	}
)

func NewHttpClient(baseURL, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: c.New(header, baseURL, 5000*time.Millisecond, true),
	}
}

func (c *clientHTTP) Push(_ context.Context, title, message, urlNotify string) error {

	dataResponse := DataResponse{}

	u := url.URL{}
	u.Path = "/push"

	request := PushReq{
		Title:   title,
		Message: message,
		Url:     urlNotify,
	}
	reps := c.client.Post(u.String(), request)

	if reps.Err != nil {
		return reps.Err
	}

	if err := reps.FillUp(&dataResponse); err != nil {
		return fmt.Errorf("%s", reps)
	}

	if reps.StatusCode > 299 {
		return fmt.Errorf("%s", dataResponse.Error)
	}

	return nil
}
