# example_https_server_golang
# run test step 1
 
init https server :
  
    go run dropbox.go worker.go server_step_1.go util.go
 
 client request:												
	
	curl -k -X POST  "https://127.0.0.1:8443/upload/dropbox" -F "image=@WorldCup.jpg" -H "Authorization : zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm"
   =>>> response 
	
	{"data":"https://dl.dropboxusercontent.com/apitl/1/AADLZgLUBzCKoSE5XD6Pm1Qy2Yw4GdNICaeIeXZDwknXzexVw4fu2Zdf-yORKU36S9qGP54hL3_I4kugspGP-yVk2bIe7J7MclbqfWUa5MZPZ8G7J1z7oQvDsTHcsK50KZ4xGYgBnOJPacfv7oZPjCwrkq9vTR4QOuRiuh0uq8YkkdVnda8FDCKvNI8dCtl6zSfVDFP7G8yIChPn7SEiZXyQWpEvS2GwLUUPgtDWtmrzSCX0pkj5cSw4M8amP58Lc9hIBfRihEv8Os150ZpV4g3s","error":"","status":"success"}


# run test step 2 : upload to dropbox through a grpc server 
  init GRPC server first ( have some workers pool upload file to Dropbox - it keep alive connection with dropbox)
	
	go run util.go pool_worker.go dropbox.go grpc.go server.go server_step_1.go server_step_2.go --step 3
	
  init https server (have some workers pool send request to grpc server - it keep alive connection with grpc server ):
  
    go run util.go pool_worker.go dropbox.go grpc.go server.go server_step_1.go server_step_2.go --step 2
    worker 5 started  upload
    worker 5 finished job
    worker 1 started  upload
    worker 1 finished job
 
 client request 1 :
  
		curl -k -X POST  "https://127.0.0.1:8443/upload/dropbox" -F "image=@WorldCup.jpg" -H "Authorization : zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm"
  ==> response 
  
		{"data":"https://dl.dropboxusercontent.com/apitl/1/AADLZgLUBzCKoSE5XD6Pm1Qy2Yw4GdNICaeIeXZDwknXzexVw4fu2Zdf-yORKU36S9qGP54hL3_I4kugspGP-yVk2bIe7J7MclbqfWUa5MZPZ8G7J1z7oQvDsTHcsK50KZ4xGYgBnOJPacfv7oZPjCwrkq9vTR4QOuRiuh0uq8YkkdVnda8FDCKvNI8dCtl6zSfVDFP7G8yIChPn7SEiZXyQWpEvS2GwLUUPgtDWtmrzSCX0pkj5cSw4M8amP58Lc9hIBfRihEv8Os150ZpV4g3s","error":"","status":"success"}
  
client request 2
  
		curl -k -X POST  "https://127.0.0.1:8443/upload/dropbox" -F "image=@WorldCup.jpg" -H "Authorization : zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnmx"
=>> response
 
	{"data":"","error":"Error in call to API function \"files/upload\": The given OAuth 2 access token is malformed.","status":"fail"}
  
  
  
  
  
