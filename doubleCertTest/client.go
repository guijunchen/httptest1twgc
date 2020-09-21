package main

import (
	//"crypto/tls"
	//"github.com/cetcxinlian/cryptogm/tls"
	//"crypto/x509"
	//"github.com/cetcxinlian/cryptogm/x509"
	"github.com/Hyperledger-TWGC/ccs-gm/tls"
	"github.com/Hyperledger-TWGC/ccs-gm/x509"
	"fmt"
	"io/ioutil"
	"github.com/Hyperledger-TWGC/net-go-gm/http"
)

func main()  {
	pool := x509.NewCertPool()
	//tassl
	caCertPath := "../assets/tasslcertdns/CA.crt"
	signCertFile := "../assets/tasslcertdns/CS.crt"
	signKeyFile := "../assets/tasslcertdns/CS.key"
	cipherCertFile := "../assets/tasslcertdns/CE.crt"
	cipherKeyFile := "../assets/tasslcertdns/CE.key"

	//caCertPath := "../assets/sudouble-cert/client_ca.crt"
	//signCertFile := "../assets/sudouble-cert/client_sign.crt"
	//signKeyFile := "../assets/sudouble-cert/client_sign.key"
	//cipherCertFile := "../assets/sudouble-cert/client_cipher.crt"
	//cipherKeyFile := "../assets/sudouble-cert/client_cipher.key"


	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("Readfile Err", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	signCert, err := tls.LoadX509KeyPair(signCertFile, signKeyFile)
	if err != nil {
		fmt.Println("Loadx509keypair err", err)
		return
	}
	cipherCert, err := tls.LoadX509KeyPair(cipherCertFile, cipherKeyFile)
	if err != nil {
		fmt.Println("Loadx509keypair err", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			Certificates: []tls.Certificate{signCert,cipherCert},
			GMSupport: &tls.GMSupport{},
		},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
