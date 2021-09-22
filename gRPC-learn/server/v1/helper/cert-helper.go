package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

// GetServerCredentials 根据证书分拣生成服务端credentials
func GetServerCredentials(certFile string, keyFile string, caFile string) *credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalln("tls.LoadX509KeyPair err,", err)
	}
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalln("ioutil.ReadFile err,", err)
	}

	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return &creds
}

// GetClientCredential 根据证书分拣生成客户端credentials
func GetClientCredential(certFile string, keyFile string, caFile string) *credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// https://blog.csdn.net/ma_jiang/article/details/111992609
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
