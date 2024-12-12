package service

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/api"
	"github.com/alserok/preview_proxy/server/internal/service/models"
	"github.com/alserok/preview_proxy/server/internal/utils"
	"net/url"
	"sync"
	"sync/atomic"
)

type Service interface {
	GetThumbnails(ctx context.Context, req models.DownloadThumbnailsReq) (models.DownloadThumbnailsRes, error)
}

func New(cls Clients) *service {
	return &service{
		youtubeAPIClient: cls.YoutubeAPIClient,
	}
}

type Clients struct {
	YoutubeAPIClient api.YoutubeAPI
}

type service struct {
	youtubeAPIClient api.YoutubeAPI
}

func (s *service) GetThumbnails(ctx context.Context, req models.DownloadThumbnailsReq) (models.DownloadThumbnailsRes, error) {
	var (
		failed int64
		videos []models.Video
	)

	syncCalls := func() {
		for _, videoURL := range req.VideoURLs {
			videoID, err := getVideoIDFromURL(videoURL)
			if err != nil {
				failed++
				continue
			}

			thumbnailURL, err := s.youtubeAPIClient.GetThumbnail(ctx, videoID)
			if err != nil {
				failed++
				continue
			}

			videos = append(videos, models.Video{
				VideoURL:  videoURL,
				Thumbnail: thumbnailURL,
			})
		}
	}
	asyncCalls := func() {
		workers := 3
		wg := &sync.WaitGroup{}

		chVideoURLs := make(chan string, workers)
		for _, videoURL := range req.VideoURLs {
			chVideoURLs <- videoURL
		}
		close(chVideoURLs)

		type videoData struct {
			videoURL  string
			thumbnail []byte
		}
		chData := make(chan videoData, workers)
		for i := 0; i < workers; i++ {
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()

				for {
					select {
					case videoURL, ok := <-chVideoURLs:
						if !ok {
							return
						}

						videoID, err := getVideoIDFromURL(videoURL)
						if err != nil {
							atomic.AddInt64(&failed, 1)
							continue
						}

						thumbnailURL, err := s.youtubeAPIClient.GetThumbnail(ctx, videoID)
						if err != nil {
							atomic.AddInt64(&failed, 1)
							continue
						}

						chData <- videoData{videoURL, thumbnailURL}
					case <-ctx.Done():
						return
					}
				}
			}(wg)
		}

		go func() {
			wg.Wait()
			close(chData)
		}()

		for data := range chData {
			videos = append(videos, models.Video{
				VideoURL:  data.videoURL,
				Thumbnail: data.thumbnail,
			})
		}
	}

	if req.Async {
		asyncCalls()
	} else {
		syncCalls()
	}

	return models.DownloadThumbnailsRes{
		Failed: uint32(failed),
		Total:  uint32(len(videos)),
		Videos: videos,
	}, nil
}

// https://www.youtube.com/watch?v=C91PNFPer_s
func getVideoIDFromURL(videoURL string) (string, error) {
	url, err := url.Parse(videoURL)
	if err != nil {
		return "", utils.NewError(err.Error(), utils.BadRequest)
	}

	videoID := url.Query().Get("v")
	if videoID == "" {
		return "", utils.NewError("video URL is not provided", utils.BadRequest)
	}

	return videoID, nil
}
