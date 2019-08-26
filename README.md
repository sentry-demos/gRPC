# Golang gRPC & sentry-go Demo

## Setup
Typical setup:
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```
Run these commands:
1. `cd $HOME/go/github.com/your_username` or `cd` to wherever your working Go workspace is
1. git clone git@github.com:thinkocapo/golang-grpc.git
2. go get github.com/getsentry/sentry-go  
3. paste your Sentry key into server.go
```
	// SENTRY INSTALLATION
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://a4efaa11ca764dd8a91d790c0926f810@sentry.io/1511084",
	})
```
## Run
for handled error:
1. `go run server/main.go`
2. `go run client/main.go`
3. check sentry.io to see your Event

for unhandled error:
1. `go run server/server.go`
2. `go run client/client.go`
3. check sentry.io to see your Event

## RPC Errors
https://grpc.io/docs/guides/error/  
https://github.com/grpc/grpc-go/blob/master/Documentation/rpc-errors.md  
https://godoc.org/google.golang.org/grpc/status#Status  

## GIF
handled RPC error  
![Alt Text](go-grpc-handled.gif)  
unhandled error  
![Alt Text](go-grpc-unhandled.gif)

## Documentation  
- Stuck running Go? [Note that this differs from other programming environments](https://golang.org/doc/code.html#Overview) in which every project has a separate workspace and workspaces are closely tied to version control repositories.
- Note that GOPATH must not be the same path as your Go installation.
- https://docs.sentry.io/platforms/go/#install sentry-go SDK
- https://github.com/grpc/grpc-go