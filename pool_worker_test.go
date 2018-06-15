package main

import (
	"io/ioutil"
	"log"
	"testing"
	"time"
	"sync"
	"fmt"
)

func TestWorkerUploadFileToDropBox(t *testing.T) {

	fmt.Println(" Start with test TestWorkerUploadFileToDropBox ")
	dataFile, err := ioutil.ReadFile("WorldCup.jpg")
	if (err != nil) {
		log.Println(err)
	}
	responseChan := make(chan ResponseInfo)
	upload := UploadInfoAsyn{
		info: UploadInfo{
			data:     dataFile,
			token:    "zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm",
			fullPath: "/" + GetTimeStamp() + "_WorldCup.jpg" ,
		},
		reponse: responseChan,
	}
	jobs:= make(chan UploadInfoAsyn)
	go workerUploadFileToDropBox(1, jobs)
	jobs <- upload
	select {
	case responseInfo := <-responseChan:
		if (responseInfo.err != nil) {
			t.Error(" Error : ", responseInfo.err)
		}
	case <-time.After(timeOut):
		t.Error(" Request Time Out ")
	}

	fmt.Println(" Done with test TestWorkerUploadFileToDropBox ")
}



func TestWorkerUploadFileToDropBoxThroughGRPC(t *testing.T) {

	fmt.Println(" Start with test TestWorkerUploadFileToDropBoxThroughGRPC ")

	dataFile, err := ioutil.ReadFile("WorldCup.jpg")
	if (err != nil) {
		log.Println(err)
	}
	responseChan := make(chan ResponseInfo)
	upload := UploadInfoAsyn{
		info: UploadInfo{
			data:     dataFile,
			token:    "zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm",
			fullPath: "/" + GetTimeStamp() + "_WorldCup.jpg" ,
		},
		reponse: responseChan,
	}
	jobs := make(chan UploadInfoAsyn)
	go InitGRPCServer(1,":7777")
	go workerSendRequestToGRPC(1,"127.0.0.1:7777", jobs)

	jobs <- upload
	select {
	case responseInfo := <-responseChan:
		if (responseInfo.err != nil) {
			t.Error(" Error : ", responseInfo.err)
		}
	case <-time.After(timeOut):
		t.Error(" Request Time Out ")
	}

	fmt.Println(" Done with test TestWorkerUploadFileToDropBoxThroughGRPC ")
}


func BenchmarkWorkerUploadFile(b *testing.B) {
	dataFile, err := ioutil.ReadFile("WorldCup.jpg")
	if (err != nil) {
		log.Println(err)
	}
	responseChan := make(chan ResponseInfo)
	upload := UploadInfoAsyn{
		info: UploadInfo{
			data:     dataFile,
			token:    "zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm",
			fullPath: "/" + GetTimeStamp() + "_WorldCup.jpg" ,
		},
		reponse: responseChan,
	}
	jobRequestToGrpcs = make(chan UploadInfoAsyn)
	go workerUploadFileToDropBox(20, jobRequestToGrpcs)
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		jobRequestToGrpcs <- upload
		select {
		case <-responseChan:
			wg.Done()
		case <-time.After(10 * time.Minute):
			wg.Done()
		}
	}
	wg.Wait()
}
