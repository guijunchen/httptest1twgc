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

	//tassl
	caCart := "../assets/tasslcertdns/CA.crt"
	signCertFile := "../assets/tasslcertdns/SS.crt"
	signKeyFile := "../assets/tasslcertdns/SS.key"
	cipherCertFile := "../assets/tasslcertdns/SE.crt"
	cipherKeyFile := "../assets/tasslcertdns/SE.key"
	//su
	//caCart := "../assets/sudouble-cert/server_ca.crt"
	//signCertFile := "../assets/sudouble-cert/server_sign.crt"
	//signKeyFile := "../assets/sudouble-cert/server_sign.key"
	//cipherCertFile := "../assets/sudouble-cert/server_cipher.crt"
	//cipherKeyFile := "../assets/sudouble-cert/server_cipher.key"

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

	//err = s.ListenAndServeTLS(serverCrt, serverKey)
	err = s.ListenAndServeTLSWithDoubleCert(signCertFile, signKeyFile, cipherCertFile, cipherKeyFile)
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}
