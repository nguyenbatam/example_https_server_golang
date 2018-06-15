package main

import (
	"fmt"
	"context"
	"upload_dropbox_service"
)

func workerUploadFileToDropBox(id int, jobs <-chan UploadInfoAsyn) {
	client := InitHttpClientKeepAlive()
	for uploadInfoAsyn := range jobs {
		fmt.Println("worker", id, "start  upload")
		responseInfo := uploadInfoAsyn.info.UploadFile(client)
		fmt.Println("worker", id, "finished job")
		uploadInfoAsyn.reponse <- responseInfo
		close(uploadInfoAsyn.reponse)
	}
}

func workerSendRequestToGRPC(id int, grpcServerAdress string, jobs <-chan UploadInfoAsyn) {
	client := InitGRPCClient(grpcServerAdress)
	for uploadInfoAsyn := range jobs {
		fmt.Println("worker", id, "start send request to grpc server")
		response, err := sendRequestToGPRC(client, uploadInfoAsyn.info)
		fmt.Println("worker", id, "finished recevice response from grpc server")
		uploadInfoAsyn.reponse <- ResponseInfo{response: response, err: err}
		close(uploadInfoAsyn.reponse)
	}
}

func sendRequestToGPRC(client upload_dropbox_service.DropboxServiceClient, info UploadInfo) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	response, err := client.Upload(ctx, &upload_dropbox_service.UploadInfo{FullPath: info.fullPath, Data: info.data, Token: info.token})
	if err != nil {
		return "", err
	}
	return response.Message, nil
}
