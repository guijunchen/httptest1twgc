package main

import (
	//"crypto/tls"
	"github.com/Hyperledger-TWGC/ccs-gm/tls"
	"fmt"
	//"github.com/cetcxinlian/cryptogm/tls"
	//"crypto/x509"
	//"github.com/cetcxinlian/cryptogm/x509"
	"github.com/Hyperledger-TWGC/ccs-gm/x509"
	"github.com/Hyperledger-TWGC/net-go-gm/http"
	"io/ioutil"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!\n")
}

func main(){
	pool := x509.NewCertPool()
	//caCart := "./assets/server/tls/ca.crt"
	//serverCrt := "./assets/server/tls/server.crt"
	//serverKey := "./assets/server/tls/server.key"
	//gmssl
	//caCart := "./assets/server/gmtls/ca.crt"
	//serverCrt := "./assets/server/gmtls/server.crt"
	//serverKey := "./assets/server/gmtls/server.key"
	//tassl
	caCart := "./assets/tasslcerts/CA.crt"
	serverCrt := "./assets/tasslcerts/Server.crt"
	serverKey := "./assets/tasslcerts/Server.key"
	//su
	//caCart := "./assets/sudouble-cert/server_ca.crt"
	//serverCrt := "./assets/sudouble-cert/server_sign.crt"
	//serverKey := "./assets/sudouble-cert/server_sign.key"
	//testcryptogen
	//caCart := "./assets/testcryptogen/tlspeer0/ca.crt"
	//serverCrt := "./assets/testcryptogen/tlspeer0/server.crt"
	//serverKey := "./assets/testcryptogen/tlspeer0/server.key"


	caCrt ,err := ioutil.ReadFile(caCart)
	if err != nil {
		fmt.Println("read err", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: "localhost:8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs: pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS(serverCrt, serverKey)
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
