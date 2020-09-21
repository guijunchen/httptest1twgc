package  main

import (
	"fmt"
	//"net/http"
	"github.com/Hyperledger-TWGC/net-go-gm/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!")
}

func main() {
	//serverCrt := "../assets/testcryptogen/tlspeer0/server.crt"
	//serverKey := "../assets/testcryptogen/tlspeer0/server.key"
	serverCrt := "../assets/susingle-cert/server.crt"
	serverKey := "../assets/susingle-cert/server.key"
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":8081",
		serverCrt, serverKey, nil)
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
