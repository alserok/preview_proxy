package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const (
	async = iota + 1
	localAsync
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	s := flag.Int("async", 0, "")
	flag.Parse()

	cc, err := grpc.NewClient(os.Getenv("SERVER_ADDR"))
	if err != nil {
		panic("failed to dial with server: " + err.Error())
	}

	client := proto.NewPreviewProxyClient(cc)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
		}

		val := scanner.Text()
		videoURLs := strings.Split(val, " ")

		switch *s {
		case async:
			res, err := client.GetThumbnails(context.Background(), &proto.GetThumbnailReq{
				VideoUrls: videoURLs,
				Async:     true,
			})
			if err != nil {
				fmt.Printf("failed to get thumbnails: %s\n", err.Error())
				continue
			}

			for _, data := range res.Videos {
				fmt.Printf("video url: %s\tthumbnail: %s\n", data.VideoUrl, data.Thumbnail)
			}
			fmt.Println()
			fmt.Printf("response:\nsuccessfully:%d\tfailed:%d\n", res.Total-res.Failed, res.Failed)
			fmt.Println()
		case localAsync:
			wg := sync.WaitGroup{}
			wg.Add(len(videoURLs))
			chData := make(chan *proto.GetThumbnailRes, len(videoURLs))

			for _, url := range videoURLs {
				go func() {
					defer wg.Done()

					res, err := client.GetThumbnails(context.Background(), &proto.GetThumbnailReq{
						VideoUrls: []string{url},
						Async:     false,
					})
					if err == nil {
						chData <- res
					}
				}()
			}

			go func() {
				wg.Wait()
				close(chData)
			}()

			succeeded := 0
			for data := range chData {
				if data.Failed == 0 {
					succeeded++
					continue
				}

				fmt.Printf("video url: %s\tthumbnail: %s\n", data.Videos[0].VideoUrl, data.Videos[0].Thumbnail)
			}
			fmt.Println()
			fmt.Printf("response:\nsuccessfully:%d\tfailed:%d\n", succeeded, len(videoURLs)-succeeded)
			fmt.Println()
		default:
			res, err := client.GetThumbnails(context.Background(), &proto.GetThumbnailReq{
				VideoUrls: strings.Split(val, " "),
				Async:     false,
			})
			if err != nil {
				fmt.Printf("failed to get thumbnails: %s\n", err.Error())
				continue
			}

			for _, data := range res.Videos {
				fmt.Printf("video url: %s\tthumbnail: %s\n", data.VideoUrl, data.Thumbnail)
			}
			fmt.Println()
			fmt.Printf("response:\nsuccessfully:%d\tfailed:%d\n", res.Total-res.Failed, res.Failed)
			fmt.Println()
		}
	}
}
