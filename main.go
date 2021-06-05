package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-cli/services"
	"io/ioutil"
	"log"
)

func main(){
	//客户端加载证书
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "leixiaotian")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//客户端加载证书,双向验证
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: 				 []tls.Certificate{cert}, //客户端证书
		ServerName:                  "localhost",
		RootCAs:                     certPool,
	})

	//conn, err := grpc.Dial(":8081", grpc.WithInsecure())		//无证书连接
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	goodsClient := services.NewGoodsServiceClient(conn)
	goodsRes, err := goodsClient.GetGoodsStock(context.Background(),&services.GoodsRequest{GoodsId:12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(goodsRes.GoodsStock)
}
