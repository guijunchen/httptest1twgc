package main

import (
	//"crypto/tls"
	//"crypto/x509"
	"github.com/Hyperledger-TWGC/ccs-gm/tls"
	"github.com/Hyperledger-TWGC/ccs-gm/x509"
	"fmt"
	"io/ioutil"
	//"net/http"
	"github.com/Hyperledger-TWGC/net-go-gm/http"
)

func main() {
	pool := x509.NewCertPool()
	//caCart := "../assets/testcryptogen/tlspeer0/ca.crt"
	//caCart := "../assets/susingle-cert/ca.crt"
	caCart := "../assets/tasslcerts/CA.crt"

	caCrt, err := ioutil.ReadFile(caCart)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool,GMSupport: &tls.GMSupport{}},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
