package main

import (
	"net/http"
	"io"
	"bytes"
	"time"
	"encoding/json"
	"log"
)

var jobRequestToGrpcs chan UploadInfoAsyn
func HelloPath2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func UploadPath2(w http.ResponseWriter, req *http.Request) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	token := req.Header.Get("Authorization")
	file, header, err := req.FormFile("image")
	w.Header().Set("Content-Type", "text/plain")
	if err != nil {
		response,_:=json.Marshal(HttpResponse{
			Status:"fail",
			Data:"",
			Error:err.Error(),
		})
		w.Write([]byte(response))
	} else {
		defer file.Close()
		io.Copy(&Buf, file)
		data := Buf.Bytes()
		responseChan :=make(chan ResponseInfo)
		upload:= UploadInfoAsyn{
			info:UploadInfo{
				fullPath: "/" + GetTimeStamp() + "_" + header.Filename,
				token:    token,
				data:     data,
			},
			reponse:responseChan,
		}
		jobRequestToGrpcs <-upload
		select {
		case responseInfo := <- responseChan:
			if (responseInfo.err!= nil) {
				response,_:=json.Marshal(HttpResponse{
					Status:"fail",
					Data:"",
					Error:responseInfo.err.Error(),
				})
				w.Write([]byte(response))
			} else {
				response,_:=json.Marshal(HttpResponse{
					Status:"success",
					Data:responseInfo.response,
					Error:"",
				})
				w.Write([]byte(response))
			}
		case <-time.After(2*timeOut):
			response,_:=json.Marshal(HttpResponse{
				Status:"fail",
				Data:"",
				Error:"Request Time Out To GRPC Server",
			})
			w.Write([]byte(response))
		}


	}
}

func RunServerStep2(nConnectionPool int,grpcAdress string)  {
	jobRequestToGrpcs = make(chan UploadInfoAsyn)
	for i := 1; i <= nConnectionPool; i++ {
		go workerSendRequestToGRPC(i,grpcAdress, jobRequestToGrpcs)
	}
	http.HandleFunc("/hello", HelloPath2)
	http.HandleFunc("/upload/dropbox", UploadPath2)
	err := http.ListenAndServeTLS(":8443", "ssl/server.crt", "ssl/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
