package main

import (
	"context"
	proto "easy-go-iot/user-srv/proto"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserById(id int) {
	rsp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
		Id: int32(id),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Mobile, rsp.NickName, rsp.Password)
}

func TestGetUserByMobile(mobile string) {
	rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobilerRequest{
		Mobile: mobile,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Mobile, rsp.NickName, rsp.Password)
}

func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "root",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestUpdateUser() {
	_, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       1,
		NickName: "zhaojun",
		Gender:   "female",
		Birthday: uint64(time.Now().Unix()),
	})
	if err != nil {
		panic(err)
	}

}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("zzjj%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			Password: "root",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}

func main() {
	Init()
	//TestCreateUser()
	//TestGetUserList()
	//TestGetUserById(11)
	//TestGetUserByMobile("18237509999")
	//TestUpdateUser()
	//TestGetUserById(1)

	conn.Close()
}
