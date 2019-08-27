# Golang gRPC & sentry-go Demo

## Setup
Recommended:  

1. .bashrc
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```
2. `cd` into `$HOME/go/github.com/your_username` or wherever your Go workspace and packages are, and clone repo into here

Mandatory:
1. `git clone git@github.com:thinkocapo/golang-grpc.git`
2. `go get github.com/getsentry/sentry-go`  
3. paste your Sentry key into server.go
```
	// SENTRY INSTALLATION
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://a4efaa11ca764dd8a91d790c0926f810@sentry.io/1511084",
	})
```
## Run
#### handled error:
1. `go run server/main.go`
2. `go run client/main.go`
3. check sentry.io to see your Event

#### unhandled error:
1. `go run server/server.go`
2. `go run client/client.go`
3. check sentry.io to see your Event

## GIF
handled RPC error  
![Alt Text](go-grpc-handled.gif)  
unhandled error  
![Alt Text](go-grpc-unhandled.gif)

## Documentation  
- https://grpc.io/docs/guides/error/  
- https://github.com/grpc/grpc-go
- https://github.com/grpc/grpc-go/blob/master/Documentation/rpc-errors.md  
- [sentry-go SDK](https://docs.sentry.io/platforms/go/#install)
- [RPC 'Status' Error](https://godoc.org/google.golang.org/grpc/status#Status)

## Troubleshooting
- Stuck running Go? [Note that this differs from other programming environments](https://golang.org/doc/code.html#Overview) in which every project has a separate workspace and workspaces are closely tied to version control repositories.
- Note that GOPATH must not be the same path as your Go installation.