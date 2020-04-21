package main

import (
	"context"
	"github.com/aimpsmile/x-apisix-example/gateway"
	hello "github.com/aimpsmile/x-apisix-example/grpc/proto/aimgo/passport/http2/v1"
	"github.com/micro/go-micro/v2"
	gsrv "github.com/micro/go-micro/v2/server/grpc"

	"log"
	"time"
)

type SayService struct{}

func (s *SayService) Hello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloResponse) error {
	log.Print("Received Say.Hello request")
	rsp.Code = 200
	rsp.Msg = "Hello " + req.Name
	return nil
}
func (s *SayService) Update(ctx context.Context, req *hello.UpdateRequest, rsp *hello.UpdateResponse) error {
	log.Print("Received Say.Hello request")
	rsp.Code = 200
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	srv := micro.NewService(
		micro.Server(gsrv.NewServer()),
		micro.Name("aimgo.http2.v1.passport"),
		micro.Version("v1.22.3"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	srv.Init()

	err := micro.RegisterHandler(srv.Server(), new(SayService), gateway.WithEndpoint(&gateway.Endpoint{
		// The SRV method
		Name: "SayService.Hello",
		// The HTTP paths. This can be a POSIX regex
		Path: "/say/hello",
		// The HTTP Methods for this endpoint
		Method: []string{"GET", "POST"},
		// The API handler to use
		Handler: "http2",
	}),
		gateway.WithEndpoint(&gateway.Endpoint{
			// The SRV method
			Name: "SayService.Update",
			// The HTTP paths. This can be a POSIX regex
			Path: "/say/update",
			// The HTTP Methods for this endpoint
			Method: []string{"GET", "POST"},
			// The API handler to use
			Handler: "http2",
		}),
	)
	// Register Handlers
	if err != nil {
		log.Fatal("handler:%v", err)
	}

	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal("run:%v", err)
	}
}
