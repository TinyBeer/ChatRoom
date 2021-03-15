package processes

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"errors"
	"fmt"
	"log"
	"net"
)

type UserProcess struct {
	// 暂时不需要字段
}

// 完成注册任务
func (up *UserProcess) Register(userID int, userPwd, userName string) (err error) {
	conn, err := net.Dial("tcp", "192.168.68.166:8889")
	if err != nil {
		return err
	}
	// 延迟断开
	defer conn.Close()

	// 2.准备通过conn发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.创建registerMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 封包
	err = utils.Pack(&mes, &registerMes)

	if err != nil {
		return err
	}

	// 使用Transfer发送数据
	tf := utils.NewTransfer(conn)
	err = tf.WritePkg(&mes)
	if err != nil {
		fmt.Println("注册消息发送失败")
		return
	}

	// 读取客服务端返回的mes
	resMes, err := tf.ReadPkg()
	if err != nil {
		return
	}

	// 解包
	var registerResMes message.RegisterResMes
	err = utils.Unpack(&resMes, &registerResMes)
	if err != nil {
		log.Println("Unpack failed, err=", err.Error())
		return
	}

	if registerResMes.Code != 200 {
		err = errors.New(registerResMes.Error)
	}
	return
}

func (up *UserProcess) Logout() {

	// 1.创建mes
	var mes message.Message
	mes.Type = message.LogoutMesType
	// 2.创建logoutMes
	var logoutMes message.LogoutMes
	logoutMes.User = CurUser.User
	// 3.封包
	err := utils.Pack(&mes, &logoutMes)
	if err != nil {
		fmt.Println("Pack failed, err=", err.Error())
		return
	}

	// 4.发送
	tf := utils.NewTransfer(CurUser.Conn)
	tf.WritePkg(&mes)
}

func (up *UserProcess) Login(userID int, userPwd string) (conn net.Conn, err error) {
	// 1.连接到服务器
	conn, err = net.Dial("tcp", "192.168.68.166:8889")
	if err != nil {
		return
	}

	// 2.准备通过conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	// 4.封包
	err = utils.Pack(&mes, &loginMes)

	// 使用Transfer发送数据
	tf := utils.NewTransfer(conn)
	err = tf.WritePkg(&mes)
	if err != nil {
		fmt.Println("登录消息发送失败")
		return
	}

	// 读取客服务端返回的mes
	resMes, err := tf.ReadPkg()
	if err != nil {
		// fmt.Println("err=", err.Error())
		return
	}
	// 解包
	var loginResMes message.LoginResMes
	err = utils.Unpack(&resMes, &loginResMes)
	if err != nil {
		fmt.Println("Unpack failed, err=", err.Error())
		return
	}
	if loginResMes.Code == 200 {

		CurUser.Conn = conn
		CurUser.UserID = userID
		CurUser.UserName = loginResMes.UserName
		CurUser.UserStatus = message.USER_ONLINE

		// 初始化在线用户列表
		for _, onlineUserID := range loginResMes.OnlineUsersID {
			// 初始化onlineUsers
			user := &message.User{
				UserID:     onlineUserID,
				UserStatus: message.USER_ONLINE,
			}
			onlineUsers[onlineUserID] = user
		}
	} else {
		err = errors.New(loginResMes.Error)
	}
	return
}
