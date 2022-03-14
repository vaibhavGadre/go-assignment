package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	pb "example.com/go-streams-grpc/proto"
	"google.golang.org/grpc"
)

const PORT = ":50052"

type DataServiceServer struct {
	pb.UnimplementedDataServiceServer
}

func NewDataServiceServer() *DataServiceServer {
	return &DataServiceServer{}
}

func (server *DataServiceServer) Run() error {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, server)
	fmt.Printf("Server is listening on port: %v", lis.Addr())
	return s.Serve(lis)
}

func main() {
	var stream_server *DataServiceServer = NewDataServiceServer()

	if err := stream_server.Run(); err != nil {
		fmt.Errorf("server run failed: %v", err)
	}

}

func (server *DataServiceServer) SumNumbers(in pb.DataService_SumNumbersServer) error {
	if in == nil {
		return fmt.Errorf("SumNumbers called with nil input value")
	}

	var total int64
	v, err := in.Recv()
	for err == nil {
		total += v.Value
		v, err = in.Recv()
	}

	if err != io.EOF {
		return fmt.Errorf("SumNumbers Recv err: %v", err)
	}
	in.SendAndClose(&pb.SumNumResponse{Total: total})
	return nil
}

var src = rand.NewSource(time.Now().Unix())
var rNum = rand.New(src)

func (server *DataServiceServer) RandomNumbers(in *pb.RandomNumRequest, out pb.DataService_RandomNumbersServer) error {
	if server == nil {
		return fmt.Errorf("RandomNumbers called on empty object")
	}

	if in == nil {
		return fmt.Errorf("RandomNumbers called with nil input value")
	}

	count := int(in.Count)

	if in.IsRange {
		for i := 0; i < count; i++ {
			v := rNum.Int63n(in.MaxValue-in.MinValue) + in.MinValue
			out.Send(&pb.RandomNumResponse{Value: v})
		}
	} else {
		for i := 0; i < count; i++ {
			v := rNum.Int63()
			out.Send(&pb.RandomNumResponse{Value: v})
		}
	}

	return nil
}
