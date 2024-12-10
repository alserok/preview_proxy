package service

import (
	"context"
	"github.com/alserok/preview_proxy/server/internal/api"
	"github.com/alserok/preview_proxy/server/internal/service/models"
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
			thumbnailURL, err := s.youtubeAPIClient.GetThumbnail(ctx, videoURL)
			if err != nil {
				failed++
				continue
			}

			videos = append(videos, models.Video{
				VideoURL:     videoURL,
				ThumbnailURL: thumbnailURL,
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

		chData := make(chan [2]string, workers)
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

						thumbnailURL, err := s.youtubeAPIClient.GetThumbnail(ctx, videoURL)
						if err != nil {
							atomic.AddInt64(&failed, 1)
							continue
						}

						chData <- [2]string{videoURL, thumbnailURL}
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
				VideoURL:     data[0],
				ThumbnailURL: data[1],
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
