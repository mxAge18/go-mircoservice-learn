package helpers

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func GetClientCredential(certFile string, keyFile string, caFile string) *credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalln("tls.LoadX509KeyPair err,", err)
	}
	ca, err := ioutil.ReadFile(caFile)

	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(ca)

	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //加载客户端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return &cred
}