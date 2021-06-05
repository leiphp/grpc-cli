package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-cli/services"
	"log"
)

func main(){
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
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
