package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "example.com/go-streams-grpc/proto"

	"google.golang.org/grpc"
)

const address = "localhost:50052"

func main() {
	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		fmt.Errorf("Failed to connect: %v", err)
	}

	// testRandomNumbersStream(conn)
	testSumNumbers(conn)
}

var src = rand.NewSource(time.Now().Unix())
var r = rand.New(src)

func testSumNumbers(conn *grpc.ClientConn) {
	client := pb.NewDataServiceClient(conn)

	defer conn.Close()

	nums, err := client.SumNumbers(context.Background())
	if err != nil {
		log.Fatalf("SumNumbers call err: %v", err)
	}

	for i := 0; i < 5; i++ {
		v := r.Int63n(100)
		fmt.Printf("%2v: %v\n", i, v)
		nums.Send(&pb.SumNumRequest{Value: v})
	}

	res, err := nums.CloseAndRecv()
	if err != nil {
		log.Fatalf("SumNumbers CloseAndRecv err: %v", err)
	}

	fmt.Printf("Total sum of numbers: %v\n", res.Total)
}

func testRandomNumbersStream(conn *grpc.ClientConn) {
	client := pb.NewDataServiceClient(conn)

	defer conn.Close()

	in := &pb.RandomNumRequest{Count: 5, IsRange: true, MinValue: 500, MaxValue: 700}

	randNums, err := client.RandomNumbers(context.Background(), in)

	if err != nil {
		log.Fatalf("RandomNumbers call failed: %v", err)
	}
	v, err := randNums.Recv()
	i := 1
	for err == nil {
		fmt.Printf("%2v. %v\n", i, v.Value)
		i++
		v, err = randNums.Recv()
	}
	if err != io.EOF {
		fmt.Printf("Error in receiver: %v", err)
	}
}
