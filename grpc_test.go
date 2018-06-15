package main

import (
	"testing"
	"upload_dropbox_service"
	"context"
	"fmt"
)

func TestGRPC(t *testing.T) {

	fmt.Println(" Start with test GRPC ")

	go InitGRPCServer(1, ":8888")
	client := InitGRPCClient("127.0.0.1:8888")
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	response, err := client.Ping(ctx, &upload_dropbox_service.UploadInfo{})
	if err != nil {
		t.Error(" Error Ping to GRPC Server: ", err)
	}
	if (response.Message != "Pong") {
		t.Error(" Error Ping to GRPC Server: ", err)
	}

	fmt.Println(" Done with test GRPC ")

}
