package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
	"encoding/json"
	"errors"
)

type UploadInfo struct {
	data     []byte
	fullPath string
	token    string
}
type UploadInfoAsyn struct {
	info    UploadInfo
	reponse chan ResponseInfo
}

type ResponseInfo struct {
	response string
	err      error
}

func (uploadInfo *UploadInfo) UploadFile(client *http.Client) ResponseInfo {
	req, err := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", bytes.NewReader(uploadInfo.data))
	if err != nil {
		return ResponseInfo{"", err}
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Authorization", "Bearer "+uploadInfo.token)
	req.Header.Set("Dropbox-API-Arg", "{\"path\":\""+uploadInfo.fullPath+"\"}")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("fail to read response data")
		return ResponseInfo{"", err}
	}
	var mapObject = make(map[string]interface{})
	err = json.Unmarshal(respByte, &mapObject)
	if (err != nil) {
		return ResponseInfo{"", errors.New(string(respByte))}
	}
	fileName, check := mapObject["path_display"].(string)
	if (! check) {
		return ResponseInfo{"", errors.New(string(respByte))}
	}
	responseInfo := getLinkFile(client, uploadInfo.token, fileName)
	if (responseInfo.err != nil) {
		return ResponseInfo{"", err}
	}
	return ResponseInfo{responseInfo.response, nil}
}

func getLinkFile(client *http.Client, token string, fileName string) ResponseInfo {
	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/get_temporary_link",
		bytes.NewReader([]byte("{\"path\":\""+fileName+"\"}")))
	if err != nil {
		return ResponseInfo{"", err}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseInfo{"", err}
	}
	var mapObject = make(map[string]interface{})
	err = json.Unmarshal(respByte, &mapObject)
	if (err != nil) {
		return ResponseInfo{"", errors.New(string(respByte))}
	}
	link, check := mapObject["link"].(string)
	if (!check) {
		return ResponseInfo{"",errors.New(string(respByte))}
	}
	return ResponseInfo{link, nil}
}

//func main() {
//	client:= InitHttpClient();
//	data,err:=ioutil.ReadFile("WorldCup.jpg")
//	if(err!=nil){
//		log.Println(err)
//	}
//	response,err:=UploadFile(client,"/WorldCup_" + strconv.FormatInt(time.Now().UnixNano() / int64(time.Millisecond),10)+".jpg",data,"zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm")
//	if(err!=nil){
//		log.Println(err)
//	}
//	fmt.Println(response)
//}
