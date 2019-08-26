# Golang GRPC & sentry-go Demo

## Instructions
Typical setup:
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

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
3. go run server/server.go
4. go run client/client.go
5. check logs
6. visit sentry.io to see your Event

## RPC Errors
https://grpc.io/docs/guides/error/  
https://github.com/grpc/grpc-go/blob/master/Documentation/rpc-errors.md  
https://godoc.org/google.golang.org/grpc/status#Status  

## Documentation  
- Stuck running Go? [Note that this differs from other programming environments](https://golang.org/doc/code.html#Overview) in which every project has a separate workspace and workspaces are closely tied to version control repositories.
- Note that GOPATH must not be the same path as your Go installation.
- https://docs.sentry.io/platforms/go/#install sentry-go SDK
- https://github.com/grpc/grpc-go
- https://github.com/grpc/grpc-go/tree/master/examples  
