package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo14-Laboratorio-3/proto"

	"google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedBrokerServer
}

// Initialize the random number generator
func init() {
    rand.Seed(time.Now().UnixNano())
}

var fulcrumServers = []string{
    "dist054:50056",
    "dist055:50057",
    "dist056:50058",
}

func (s *server) RedirectInformant(ctx context.Context, in *pb.InformantRequest) (*pb.FulcrumAddress, error) {
    // Choose a random Fulcrum server
    address := fulcrumServers[rand.Intn(len(fulcrumServers))]
    fmt.Printf("Redirecting Informant to %v\n", address)

    return &pb.FulcrumAddress{Address: address}, nil
}

func (s *server) Mediate(ctx context.Context, in *pb.Message) (*pb.Acknowledgement, error) {
    // Choose a random Fulcrum server
    address := fulcrumServers[rand.Intn(len(fulcrumServers))]

    // Connect to the Fulcrum server
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    // Create a Fulcrum client
    client := pb.NewFulcrumClient(conn)

    // Forward the message to the Fulcrum server
    message := &pb.Message{
        Sector: in.GetSector(),
        Base: in.GetBase(),
        VectorClock: in.GetVectorClock(), // Add this line
    }
    ack, err := client.ProcessVanguardMessage(ctx, message)
    fmt.Printf("Mediating message %v, %v\n", in.GetSector(), in.GetBase())
    if err != nil {
        return nil, err
    }

    // Return the acknowledgement from the Fulcrum server
    return ack, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterBrokerServer(s, &server{})

    fmt.Printf("Broker listening at %v\n", lis.Addr().String())

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}