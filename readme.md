# Task

**Rules**

- Use the Go programming language in its latest stable version at the time of working on the task.
- Only use Go's standard libraries. External dependencies can be used for logging, implementing client-server features (HTTP, gRPC), and writing tests. External dependencies that implement the task's logic, such as algorithms, cannot be used.

**Criteria**

- Does the program meet the task requirements?
- Are the results of the program's execution correct?
- Readability and logical organization of functions and packages in the source code.
- Did the developer try to follow the main rules of Go code writing? The rules are available on the official Go documentation site and in best practice guides like the "Uber Go Style Guide." We will not nitpick details, but candidates who do not adhere to any coding structure principles or use an inherited package organization from projects like Java or C++ will be disqualified.
- Are the results of the program's execution correct?
- Are there tests available?
- Is there a written pipeline for testing the code (GitHub Actions, Gitlab CI/CD)?
- Was the program able to run on the first attempt according to the instructions in README.md, and were the results reproducible?
- Is there no excessive use of memory or CPU time during the execution of the program?
- The repository or folder containing the solution should not include any binary build artifacts.

**Task**

**Task Description**

Write a gRPC proxy service to download thumbnails (video previews) from YouTube videos. Upon repeated requests for the same video, the service should return a cached response (you can use runtime caching, but it will be a plus if a temporary storage like SQLite is used). It is also suggested to write a client-side utility as a command-line tool that accepts video links as parameters. The command-line utility should include a --async flag, which allows downloading a large number of files asynchronously.

---
# Solution

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

    
