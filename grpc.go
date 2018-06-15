package main

import (
	"upload_dropbox_service"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
	"fmt"
	"context"
	"errors"
)

type server struct{}

var jobUploads chan UploadInfoAsyn

func (s *server) Upload(ctx context.Context, request *upload_dropbox_service.UploadInfo) (*upload_dropbox_service.ResponseResult, error) {
	responseChan := make(chan ResponseInfo)
	upload := UploadInfoAsyn{
		info: UploadInfo{
			fullPath: request.FullPath,
			token:    request.Token,
			data:     []byte(request.Data),
		},
		reponse: responseChan,
	}
	jobUploads <- upload
	select {
	case responseInfo := <-responseChan:
		if (responseInfo.err != nil) {
			return nil, responseInfo.err
		} else {
			return &upload_dropbox_service.ResponseResult{Message: responseInfo.response}, nil
		}
	case <-time.After(timeOut):
		return nil,errors.New("Request upload file to dropbox time out")
	}
}

func (s *server) Ping(ctx context.Context, request *upload_dropbox_service.UploadInfo) (*upload_dropbox_service.ResponseResult, error) {
	return &upload_dropbox_service.ResponseResult{Message: "Pong"}, nil
}

func InitGRPCServer(nWoker int, port string) {
	jobUploads = make(chan UploadInfoAsyn)
	for i := 1; i <= nWoker; i++ {
		go workerUploadFileToDropBox(i, jobUploads)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	upload_dropbox_service.RegisterDropboxServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitGRPCClient(adress string) (upload_dropbox_service.DropboxServiceClient) {
	conn, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("connection success")
	return upload_dropbox_service.NewDropboxServiceClient(conn)
}
