## Server cache
  Server uses Redis cache. To launch it:

    docker compose -f server/docker-compose.redis.yaml up -d

## Server
  Interacts with external api to get videos thumbnail, returns response with slices
  of bytes and amount of failed requests.

    cd server
    go run main.go

## Client

Firstly, launch the server!

Then: 
* Server async
      
      cd client
      go run main.go --async 1

* Client async

      cd client
      go run main.go --async 2

* Sync

      cd client
      go run main.go

In the terminal you will see `ready`.  After that input your addresses in such format:
  
    https://www.youtube.com/watch?v=aalb4uiuRyo https://www.youtube.com/watch?v=e-or_D-qNqM

Each address should be separated with space.

Server response:
```protobuf
message GetThumbnailRes {
  uint32 failed = 1;
  uint32 total = 2;
  repeated Video videos = 3;
}

message Video {
  string video_url = 1;
  bytes thumbnail = 2;
}
```

All files are stored in `client/files`

### Config

All configuration is provided via env variables.

    
