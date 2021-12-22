package service

import (
	"log"
	"os"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/Yusuf1999-hub/to-do-service/genproto"
)

var task pb.TaskServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	task = pb.NewTaskServiceClient(conn)
	os.Exit(m.Run())
}
