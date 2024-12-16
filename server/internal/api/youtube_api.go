package api

import (
	"context"
	"fmt"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"github.com/alserok/preview_proxy/server/internal/utils"
	"io"
	"net/http"
	"time"
)

type YoutubeAPI interface {
	GetThumbnail(ctx context.Context, videoID string) ([]byte, error)
}

func NewYoutubeAPIClient() *youtubeAPIClient {
	cl := http.DefaultClient
	cl.Transport = &retryTransport{
		RoundTripper: http.DefaultTransport,
		maxRetries:   3,
		delay:        200 * time.Millisecond,
	}
	cl.Timeout = time.Millisecond * 600

	return &youtubeAPIClient{
		cl: cl,
	}
}

type retryTransport struct {
	RoundTripper http.RoundTripper
	maxRetries   int
	delay        time.Duration
}

func (rt *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var err error
	var resp *http.Response

	for i := 0; i < rt.maxRetries; i++ {
		resp, err = rt.RoundTripper.RoundTrip(req)
		if err == nil {
			return resp, nil
		}
		time.Sleep(rt.delay)
	}

	return nil, err
}

type youtubeAPIClient struct {
	cl *http.Client
}

func (cl *youtubeAPIClient) GetThumbnail(ctx context.Context, videoID string) ([]byte, error) {
	log := logger.FromContext(ctx)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://i.ytimg.com/vi/%s/%s.jpg", videoID, "maxresdefault"),
		nil)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	log.Debug("sending api request", logger.WithArg("addr", req.URL.Host))

	res, err := cl.cl.Do(req)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	log.Debug("received response", logger.WithArg("status", res.StatusCode))

	switch res.StatusCode {
	case http.StatusOK:
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, utils.NewError(fmt.Sprintf("failed to read body: %s", err.Error()), utils.Internal)
		}

		return b, nil
	case http.StatusNotFound:
		return nil, utils.NewError("video not found", utils.NotFound)
	case http.StatusBadRequest:
		return nil, utils.NewError("invalid data provided", utils.BadRequest)
	default:
		return nil, utils.NewError("internal error", utils.Internal)
	}
}
