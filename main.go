package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-cli/helper"
	"grpc-cli/services"
	"log"
)

func main(){
	//conn, err := grpc.Dial(":8081", grpc.WithInsecure())		//无证书连接
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	goodsClient := services.NewGoodsServiceClient(conn)
	//获取单个商品
	//goodsRes, err := goodsClient.GetGoodsStock(context.Background(),&services.GoodsRequest{GoodsId:12})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(goodsRes.GoodsStock)

	//获取多个商品
	response, err := goodsClient.GetGoodsStocks(context.Background(), &services.GoodsSize{Size:10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Goodsres)
}
