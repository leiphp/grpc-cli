package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-cli/services"
	"log"
)

func main(){
	//客户端加载证书
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "leixiaotian")
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
