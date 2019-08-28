/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Binary server is an example server.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/status"

	"github.com/getsentry/sentry-go"
)

var port = flag.Int("port", 50052, "port number")

// server is used to implement helloworld.GreeterServer.
type server struct {
	mu    sync.Mutex
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
	ds, err := st.WithDetails(
		&epb.QuotaFailure{
			Violations: []*epb.QuotaFailure_Violation{{
				Subject:     fmt.Sprintf("name:%s", in.Name),
				Description: "Limit one greeting per person",
			}},
		},
	)
	if err != nil {
		return nil, st.Err()
	}

	log.Printf("%v", ds.Err())

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtra("Details", ds.Details())
	})
	sentry.CaptureException(ds.Err())
	return nil, ds.Err()
	// }
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// SENTRY INSTALLATION
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://a4efaa11ca764dd8a91d790c0926f810@sentry.io/1511084",
	})

	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}