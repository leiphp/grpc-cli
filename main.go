package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"grpc-cli/helper"
	"grpc-cli/services"
	"log"
	"time"
)

func main(){
	//conn, err := grpc.Dial(":8081", grpc.WithInsecure())		//无证书连接
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//商品client
	goodsClient := services.NewGoodsServiceClient(conn)
	//订单client
	ordersClient := services.NewOrdersServiceClient(conn)

	//商品服务-获取单个商品
	goodsRes, err := goodsClient.GetGoodsStock(context.Background(),&services.GoodsRequest{GoodsId:10,GoodsArea:services.GoodsAreas_B})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(goodsRes.GoodsStock)

	//商品服务-获取多个商品
	//response, err := goodsClient.GetGoodsStocks(context.Background(), &services.GoodsSize{Size:10})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(response.Goodsres)

	//订单服务-创建订单
	t := timestamp.Timestamp{Seconds:time.Now().Unix()}
	ordersRes, err := ordersClient.CreateOrder(context.Background(),&services.OrderMain{OrderId:1001,OrderNo:"0001",OrderMoney:90,OrderTime:&t})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ordersRes)

}
