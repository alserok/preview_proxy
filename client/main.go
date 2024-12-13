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
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	async := flag.Bool("async", false, "")
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
		res, err := client.GetThumbnails(context.Background(), &proto.GetThumbnailReq{
			VideoUrls: strings.Split(val, " "),
			Async:     *async,
		})
		if err != nil {
			fmt.Printf("failed to get thumbnails: %s\n", err.Error())
			continue
		}

		fmt.Printf("response:\nsuccessfully:%d\tfailed:%d\n", res.Total-res.Failed, res.Failed)
		for _, data := range res.Videos {
			fmt.Printf("video url: %s\tthumbnail: %s\n", data.VideoUrl, data.Thumbnail)
		}
		fmt.Println()
	}
}
