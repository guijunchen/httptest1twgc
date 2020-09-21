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
	//caCertPath := "./assets/client/tls/ca.crt"
	//clientCrt := "./assets/client/tls/client.crt"
	//clientKey := "./assets/client/tls/client.key"
	//gmssl
	//caCertPath := "./assets/client/gmtls/ca.crt"
	//clientCrt := "./assets/client/gmtls/client.crt"
	//clientKey := "./assets/client/gmtls/client.key"
	//tassl
	//caCertPath := "./assets/tasslcerts/CA.crt"
	//clientCrt := "./assets/tasslcerts/CS.crt"
	//clientKey := "./assets/tasslcerts/CS.key"
	//su
	//caCertPath := "./assets/sudouble-cert/client_ca.crt"
	//clientCrt := "./assets/sudouble-cert/client_sign.crt"
	//clientKey := "./assets/sudouble-cert/client_sign.key"
	//testcryptogen
	caCertPath := "./assets/testcryptogen/tlsuser/ca.crt"
	clientCrt := "./assets/testcryptogen/tlsuser/client.crt"
	clientKey := "./assets/testcryptogen/tlsuser/client.key"


	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("Readfile Err", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		fmt.Println("Loadx509keypair err", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			Certificates: []tls.Certificate{cliCrt},
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
