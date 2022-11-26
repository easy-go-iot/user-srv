package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"easy-go-iot/user-srv/global"
	"easy-go-iot/user-srv/handler"
	"easy-go-iot/user-srv/initialize"
	proto "easy-go-iot/user-srv/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//IP := flag.String("ip", GetLocalIP(), "ip地址")
	IP := flag.String("ip", GetLocalIP(), "ip地址")
	Port := flag.Int("port", 0, "端口号")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	flag.Parse()
	zap.S().Info("user-srv start at ip: ", *IP)
	if *Port == 0 {
		*Port = global.ServerConfig.Port
	}

	zap.S().Info("user-srv start at port: ", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", GetLocalIP(), Port),
			grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		userClient := proto.NewUserClient(conn)
		rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
			Id: int32(1),
		})
		if err != nil {
			zap.S().Error(err)
		}
		zap.S().Info(rsp.Mobile, rsp.NickName, rsp.Password)
		time.Sleep(5 * time.Second)
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

// GetLocalIP 获取本地eth0 IP
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsLinkLocalUnicast() {
			if ipnet.IP.To4() != nil { //本地网卡eth0的ip
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
