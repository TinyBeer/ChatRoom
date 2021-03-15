package processes

import (
	"ChartRoom/common/message"
	"ChartRoom/common/utils"
	"fmt"
	"log"
	"net"
)

// 和服务器保持通信
func ServerMesProcess(conn net.Conn) {
	defer conn.Close()
	// 创建一个Transfer 不停的读取消息
	tf := utils.NewTransfer(conn)
	for {
		// fmt.Println("客户端正在读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			log.Println("tf.ReadPkg failed, err=", err.Error())
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 处理用户状态更新消息
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = utils.Unpack(&mes, &notifyUserStatusMes)
			if err != nil {
				log.Println("Unpack failed, err=", err.Error())
				continue
			}
			updateUserStatus(&notifyUserStatusMes)
			OutputOnlineUsers()
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("获取到未知消息类型")
		}

	}
}
