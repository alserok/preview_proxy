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
	"path"
	"strings"
	"sync"
	"time"
)

const (
	async mode = iota + 1
	localAsync
)

type mode uint

func (m mode) String() string {
	switch m {
	case async:
		return "server async"
	case localAsync:
		return "client async"
	default:
		return "sync"
	}
}

func main() {
	s := flag.Int("async", 0, "")
	flag.Parse()
	m := mode(uint(*s))

	cc, err := grpc.NewClient(os.Getenv("SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to dial with server: " + err.Error())
	}

	client := proto.NewPreviewProxyClient(cc)

	if _, err = os.Stat(filesDir); err != nil {
		if err = os.Mkdir(filesDir, 0777); os.IsNotExist(err) {
			panic("failed to init files folder: " + err.Error())
		}
	}

	fmt.Printf("ready\nmode: %s\n\n", m.String())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		switch val {
		case "":
			continue
		case "q":
			return
		default:
			videoURLs := strings.Split(val, " ")

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

			switch m {
			case async:
				serverAsync(ctx, client, videoURLs)
			case localAsync:
				clientAsync(ctx, client, videoURLs)
			default:
				defaultSync(ctx, client, videoURLs)
			}

			cancel()
		}
	}
}

func serverAsync(ctx context.Context, client proto.PreviewProxyClient, videoURLs []string) {
	res, err := client.GetThumbnails(ctx, &proto.GetThumbnailReq{
		VideoUrls: videoURLs,
		Async:     true,
	})
	if err != nil {
		fmt.Printf("failed to get thumbnails: %s\n", err.Error())
		return
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

	fmt.Printf("successfully:%d\tfailed:%d\n", res.Total-res.Failed, res.Failed)
	fmt.Println()
}

func clientAsync(ctx context.Context, client proto.PreviewProxyClient, videoURLs []string) {
	wg := sync.WaitGroup{}
	wg.Add(len(videoURLs))
	chData := make(chan *proto.GetThumbnailRes, len(videoURLs))

	for _, url := range videoURLs {
		go func() {
			defer wg.Done()

			res, err := client.GetThumbnails(ctx, &proto.GetThumbnailReq{
				VideoUrls: []string{url},
				Async:     false,
			})
			if err != nil {
				fmt.Printf("failed to get thumbnails: %s\n", err.Error())
				return
			}

			chData <- res
		}()
	}

	go func() {
		wg.Wait()
		close(chData)
	}()

	succeeded := 0
	for data := range chData {
		if data.Failed == 0 && len(data.Videos) == 1 {
			path, err := saveTo(context.Background(), strings.Split(data.Videos[0].VideoUrl, "?v=")[1], data.Videos[0].Thumbnail)
			if err != nil {
				fmt.Printf("failed to save thumbnail: %s\n", err.Error())
				continue
			}

			fmt.Printf("video url: %s\tthumbnail path: %s\n", data.Videos[0].VideoUrl, path)
			succeeded++

			continue
		}
	}
	fmt.Printf("successfully:%d\tfailed:%d\n", succeeded, len(videoURLs)-succeeded)
	fmt.Println()
}

func defaultSync(ctx context.Context, client proto.PreviewProxyClient, videoURLS []string) {
	res, err := client.GetThumbnails(ctx, &proto.GetThumbnailReq{
		VideoUrls: videoURLS,
		Async:     false,
	})
	if err != nil {
		fmt.Printf("failed to get thumbnails: %s\n", err.Error())
		return
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
	fmt.Printf("successfully:%d\tfailed:%d\n", res.Total-res.Failed, res.Failed)
	fmt.Println()
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
