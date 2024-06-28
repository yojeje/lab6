package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/yojeje/lab6"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBrokerServer
	pb.UnimplementedIngenierosServer
	pb.UnimplementedKaisServer
}

func (s *server) EnviarBroker(ctx context.Context, in *pb.Comando) (*pb.Direccion, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var n int = r.Intn(7-2)+2
	return &pb.Direccion{Dir: int64(n)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}
	s :=  grpc.NewServer()
	pb.RegisterIngenierosServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falla de servidor: %v", err)
	}
}