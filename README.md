# grpc-cli
grpc-cli是基于grpc-gateway是客户端项目，是调用grpc服务端的客户端，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点，让你快速了解gepc的使用

### 环境依赖
golang v1.14.4    

### 部署步骤
1. 加载客户端证书  
   ```go
    creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "lxtkj.cn") //第二个参数是生成证书时填写的Common Name
    ``` 
