package api

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/utils"
	"net/http"
	"time"
)

type YoutubeAPI interface {
	GetThumbnail(ctx context.Context, videoURL string) (string, error)
}

func NewYoutubeAPIClient(addr string) *youtubeAPIClient {
	cl := http.DefaultClient
	cl.Timeout = 300 * time.Millisecond
	cl.Transport = struct {
		http.RoundTripper
		maxRetries int
		delay      time.Duration
	}{
		RoundTripper: http.DefaultTransport,
		maxRetries:   3,
		delay:        30 * time.Millisecond,
	}

	return &youtubeAPIClient{
		cl:   cl,
		addr: addr,
	}
}

type youtubeAPIClient struct {
	cl *http.Client

	addr string
}

func (cl *youtubeAPIClient) GetThumbnail(ctx context.Context, videoURL string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, videoURL, nil)
	if err != nil {
		return "", utils.NewError(err.Error(), utils.Internal)
	}

	res, err := cl.cl.Do(req)
	if err != nil {
		return "", utils.NewError(err.Error(), utils.Internal)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	return "video url", nil
}
