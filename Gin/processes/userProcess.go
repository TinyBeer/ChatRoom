package processes

import (
	"ChatRoom/Gin/message"
	"ChatRoom/Gin/utils"
	"encoding/json"
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
	registerMes.UserID = userID
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	// 封包
	err = message.Pack(&mes, &registerMes)
	if err != nil {
		return err
	}

	// 序列化
	data, err := json.Marshal(&mes)
	if err != nil {
		return
	}

	// 使用Transfer发送数据
	tf := utils.NewTransfer(conn)
	err = tf.WriteData(data)
	if err != nil {
		fmt.Println("注册消息发送失败")
		return
	}

	resData, err := tf.ReadDate()
	if err != nil {
		log.Println("tf.ReadDate failed, err=", err.Error())
		return
	}
	var resMes message.Message
	err = json.Unmarshal(resData, &resMes)
	if err != nil {
		log.Println("json.Unmarshal failed, err=", err.Error())
		return
	}

	// 解包
	var registerResMes message.RegisterResMes
	err = message.Unpack(&resMes, &registerResMes)
	if err != nil {
		log.Println("Unpack failed, err=", err.Error())
		return
	}

	fmt.Println(resMes)

	if registerResMes.Code != 200 {
		err = errors.New(registerResMes.Error)
	}
	return
}

func (up *UserProcess) Check(userID int, userPwd string) (string, error) {
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "192.168.68.166:8889")
	if err != nil {
		return "", err
	}

	// 2.准备通过conn发送消息
	var mes message.Message
	mes.Type = message.CheckMesType

	var checkMes message.CheckMes
	checkMes.UserID = userID
	checkMes.UserPwd = userPwd

	// 4.封包
	err = message.Pack(&mes, &checkMes)
	if err != nil {
		log.Println("pack failed, err=", err.Error())
		return "", err
	}
	// 序列化
	data, err := json.Marshal(&mes)
	if err != nil {
		return "", err
	}

	// 使用Transfer发送数据
	tf := utils.NewTransfer(conn)
	err = tf.WriteData(data)
	if err != nil {
		fmt.Println("验证消息发送失败")
		return "", err
	}

	// 读取返回消息
	resData, err := tf.ReadDate()
	if err != nil {
		log.Println("tf.ReadDate failed, err=", err.Error())
		return "", err
	}

	var resMes message.Message
	err = json.Unmarshal(resData, &resMes)
	if err != nil {
		log.Println("json.Unmarshal failed, err=", err.Error())
		return "", err
	}

	// 解包
	var checkResMes message.CheckResMes
	err = message.Unpack(&resMes, &checkResMes)
	if err != nil {
		fmt.Println("Unpack failed, err=", err.Error())
		return "", err
	}

	if checkResMes.Code != 200 {
		err = errors.New(checkResMes.Error)
	}

	return checkResMes.UserName, err
}
