package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"grpc-cli/helper"
	"grpc-cli/services"
	"io"
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
	//用户client
	usersClient := services.NewUserServiceClient(conn)

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

	//用户服务-获取用户积分
	usersRes, err := usersClient.GetUserScore(context.Background(),&services.UserScoreRequest{Users: []*services.UserInfo{{UserId:10,UserScore:20},{UserId:11,UserScore:20}}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usersRes.Users)

	//用户服务-服务端流模式获取用户积分
	//stream, err := usersClient.GetUserScoreByServerStream(context.Background(),&services.UserScoreRequest{Users: []*services.UserInfo{{UserId:11},{UserId:12},{UserId:13},{UserId:14},{UserId:15}}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	res, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(res.Users)
	//	//以下可以开协程处理逻辑
	//}

	////用户服务-客户端流模式获取用户积分
	//var i int32
	//stream, err := usersClient.GetUserScoreByClientStream(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for j:=1;j<=3;j++ {
	//	req := services.UserScoreRequest{}
	//	req.Users = make([]*services.UserInfo, 0)
	//	for i=1;i<=5;i++ {
	//		req.Users = append(req.Users, &services.UserInfo{UserId:i})
	//	}
	//	err := stream.Send(&req)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}
	//res,_ := stream.CloseAndRecv()
	//fmt.Println(res.Users)

	//用户服务-双向流模式获取用户积分
	var i int32
	stream, err := usersClient.GetUserScoreByTWStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var uid int32=1
	for j:=1;j<=3;j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)
		for i=1;i<=5;i++ {
			req.Users = append(req.Users, &services.UserInfo{UserId:uid})
			uid++
		}
		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}
		res,err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Println(res.Users)
	}
}
