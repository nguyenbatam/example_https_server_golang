package main

import (
	"testing"
	"io/ioutil"
	"log"
	"fmt"
)

func TestUploadFile(t *testing.T) {
	dataFile, err := ioutil.ReadFile("WorldCup.jpg")
	if (err != nil) {
		log.Println(err)
	}
	uploadInfo := UploadInfo{
		data:     dataFile,
		token:    "zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm",
		fullPath: "/" + GetTimeStamp() + "_WorldCup.jpg",
	}
	client := InitHttpClientNoKeepAlive()
	responseInfo := uploadInfo.UploadFile(client)
	if responseInfo.err != nil {
		t.Error(" Error : ", responseInfo.err)
	}
	fmt.Println(" Done with test TestUploadFile ")
}

func TestGetLinkFile(t *testing.T) {
	client := InitHttpClientNoKeepAlive()
	responseInfo := getLinkFile(client, "zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm", "/1528995906933_WorldCup.jpg")
	if responseInfo.err != nil {
		t.Error(" Error : ", responseInfo.err)
	}
	fmt.Println(" Done with test TestGetLinkFile ")
}

func BenchmarkUploadFile(b *testing.B) {
	dataFile, err := ioutil.ReadFile("WorldCup.jpg")
	if (err != nil) {
		log.Println(err)
	}
	uploadInfo := UploadInfo{
		data:     dataFile,
		token:    "zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm",
		fullPath: "/" + GetTimeStamp() + "_WorldCup.jpg",
	}
	client := InitHttpClientNoKeepAlive()
	for i := 0; i < b.N; i++ {
		uploadInfo.UploadFile(client)
	}
}
