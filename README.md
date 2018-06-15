# example_https_server_golang
# run test step 1
 
 server :
  
    go run dropbox.go worker.go server_step_1.go util.go
 
 client :												
	
	curl -k -X POST  "https://127.0.0.1:8443/upload/dropbox" -F "image=@WorldCup.jpg" -H "Authorization : zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm"
   =>>> response 
	
	{"data":"https://dl.dropboxusercontent.com/apitl/1/AADLZgLUBzCKoSE5XD6Pm1Qy2Yw4GdNICaeIeXZDwknXzexVw4fu2Zdf-yORKU36S9qGP54hL3_I4kugspGP-yVk2bIe7J7MclbqfWUa5MZPZ8G7J1z7oQvDsTHcsK50KZ4xGYgBnOJPacfv7oZPjCwrkq9vTR4QOuRiuh0uq8YkkdVnda8FDCKvNI8dCtl6zSfVDFP7G8yIChPn7SEiZXyQWpEvS2GwLUUPgtDWtmrzSCX0pkj5cSw4M8amP58Lc9hIBfRihEv8Os150ZpV4g3s","error":"","status":"success"}


# run test step 2 : worker pool with chan
 server :
  
    go run dropbox.go worker.go server_step_2.go util.go
    worker 5 started  upload
    worker 5 finished job
    worker 1 started  upload
    worker 1 finished job
 
 client :
  
		curl -k -X POST  "https://127.0.0.1:8443/upload/dropbox" -F "image=@WorldCup.jpg" -H "Authorization : zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnm"
  ==> response 
  
		{"data":"https://dl.dropboxusercontent.com/apitl/1/AADLZgLUBzCKoSE5XD6Pm1Qy2Yw4GdNICaeIeXZDwknXzexVw4fu2Zdf-yORKU36S9qGP54hL3_I4kugspGP-yVk2bIe7J7MclbqfWUa5MZPZ8G7J1z7oQvDsTHcsK50KZ4xGYgBnOJPacfv7oZPjCwrkq9vTR4QOuRiuh0uq8YkkdVnda8FDCKvNI8dCtl6zSfVDFP7G8yIChPn7SEiZXyQWpEvS2GwLUUPgtDWtmrzSCX0pkj5cSw4M8amP58Lc9hIBfRihEv8Os150ZpV4g3s","error":"","status":"success"}
  
cmd2
  
		curl -k -X POST  "https://127.0.0.1:8443/upload/dropbox" -F "image=@WorldCup.jpg" -H "Authorization : zWnEITOt3sAAAAAAAAAAEYDxegNY7vq2usYt4dJcd-UZ6JTfPCNEh1QMeUYaeMnmx"
=>> response
 
	{"data":"","error":"Error in call to API function \"files/upload\": The given OAuth 2 access token is malformed.","status":"fail"}
  
  
  
  
  
