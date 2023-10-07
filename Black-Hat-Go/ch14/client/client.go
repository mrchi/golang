package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mrchi/golang/Black-Hat-Go/ch14/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial("127.0.0.1:9999", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := grpcapi.NewAdminClient(conn)

	cmd := grpcapi.Command{In: os.Args[1]}
	ctx := context.Background()
	respCmd, err := client.RunCommand(ctx, &cmd)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(respCmd.Out)
}
