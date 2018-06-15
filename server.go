package main

import (
	"fmt"
	"flag"
	"time"
)

var step = flag.String("step", "1", "")
var timeOut = 10 * time.Second
func main() {
	flag.Parse()

	port:=":7777"
	switch *step {
	case "1":
		fmt.Println(" Run https server with step " + *step)
		RunServerStep1()
	case "2":
		fmt.Println(" Run https server with step " + *step)
		RunServerStep2(10,"127.0.0.1"+port)
	case "3":
		fmt.Println(" Init GRPC Server Upload File To Drop Box")
		InitGRPCServer(10,port)
	default:
		fmt.Println(" Not Found Step "+*step)
	}

}
