package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	proto "github.com/alserok/preview_proxy/server/pkg/protobuf"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"os/signal"
	"path"
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

	cc, err := grpc.NewClient(os.Getenv("SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to dial with server: " + err.Error())
	}

	client := proto.NewPreviewProxyClient(cc)

	if _, err = os.Stat(filesDir); err != nil {
		if err = os.Mkdir(filesDir, 755); os.IsNotExist(err) {
			panic("failed to init files folder: " + err.Error())
		}
	}

	fmt.Println("ready")
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
				path, err := saveTo(context.Background(), strings.Split(data.VideoUrl, "?v=")[1], data.Thumbnail)
				if err != nil {
					res.Failed++
					fmt.Printf("failed to save thumbnail: %s\n", err.Error())
					continue
				}

				fmt.Printf("video url: %s\tthumbnail path: %s\n", data.VideoUrl, path)
			}
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
				if data.Failed != 0 && len(data.Videos) == 1 {
					path, err := saveTo(context.Background(), strings.Split(data.Videos[0].VideoUrl, "?v=")[1], data.Videos[0].Thumbnail)
					if err != nil {
						fmt.Printf("failed to save thumbnail: %s\n", err.Error())
						continue
					}

					fmt.Printf("video url: %s\tthumbnail path: %s\n", data.Videos[0].VideoUrl, path)
					succeeded++
				}
			}
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
				path, err := saveTo(context.Background(), strings.Split(data.VideoUrl, "?v=")[1], data.Thumbnail)
				if err != nil {
					res.Failed++
					fmt.Printf("failed to save thumbnail: %s\n", err.Error())
					continue
				}

				fmt.Printf("video url: %s\tthumbnail path: %s\n", data.VideoUrl, path)
			}
			fmt.Printf("response:\nsuccessfully:%d\tfailed:%d\n", res.Total-res.Failed, res.Failed)
			fmt.Println()
		}
	}
}

const filesDir = "files"

func saveTo(ctx context.Context, name string, data []byte) (string, error) {
	target := path.Join(filesDir, name+".jpg")

	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}

	if _, err = f.Write(data); err != nil {
		return "", err
	}

	return target, nil
}
