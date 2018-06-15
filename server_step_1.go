package main

import (
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"log"
)



func HelloPath(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func UploadPath(w http.ResponseWriter, req *http.Request) {
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
	}else {
		defer file.Close()
		io.Copy(&Buf, file)
		data := Buf.Bytes()
		client :=InitHttpClientNoKeepAlive()
		upload:= UploadInfo{
			fullPath:"/"+GetTimeStamp()+"_"+header.Filename,
			token:token,
			data:data,
		}
		responseInfo :=upload.UploadFile(client)
		if(responseInfo.err!=nil) {
			response,_:=json.Marshal(HttpResponse{
				Status:"fail",
				Data:"",
				Error:responseInfo.err.Error(),
			})
			w.Write([]byte(response))
		}else{
			response,_:=json.Marshal(HttpResponse{
				Status:"success",
				Data:responseInfo.response,
				Error:"",
			})
			w.Write([]byte(response))
		}
	}
}

func RunServerStep1() {
	http.HandleFunc("/hello", HelloPath)
	http.HandleFunc("/upload/dropbox", UploadPath)
	err := http.ListenAndServeTLS(":8443", "ssl/server.crt", "ssl/server.key", nil)
	//err := http.ListenAndServe(":8443",  nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
