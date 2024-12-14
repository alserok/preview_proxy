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
		cl: cl,
	}
}

type youtubeAPIClient struct {
	cl *http.Client
}

func (cl *youtubeAPIClient) GetThumbnail(ctx context.Context, videoID string) ([]byte, error) {
	log := logger.FromContext(ctx)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://i.ytimg.com/vi/%s/%s.jpg", videoID, "default"),
		nil)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	log.Debug("sending api request", logger.WithArg("addr", req.URL.Host))

	res, err := cl.cl.Do(req)
	defer func() {
		_ = res.Body.Close()
	}()
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	log.Debug("received response", logger.WithArg("status", res.StatusCode))

	switch res.StatusCode {
	case http.StatusOK:
		b, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
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
