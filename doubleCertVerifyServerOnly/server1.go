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
	signCertFile := "../assets/tasslcertdns/SS.crt"
	signKeyFile := "../assets/tasslcertdns/SS.key"
	cipherCertFile := "../assets/tasslcertdns/SE.crt"
	cipherKeyFile := "../assets/tasslcertdns/SE.key"

	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLSWithDoubleCert(":8081",
		signCertFile, signKeyFile, cipherCertFile, cipherKeyFile, nil)
	if err != nil {
		fmt.Println("ListenAndServeTLSWithDoubleCert err:", err)
	}
}
