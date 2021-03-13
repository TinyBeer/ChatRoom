package processes

import "fmt"

// UserMgr实例有且仅有一个  定义为全局变量
var (
	userMgr *UserMgr
)

// UserMgr结构体  用于管理用户连接
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr的初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加  修改
func (uMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	uMgr.onlineUsers[up.UserID] = up
}

// 删除
func (uMgr *UserMgr) DelOnlineUser(userID int) {
	delete(uMgr.onlineUsers, userID)
}

// 返回所欲在线用户
func (uMgr *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return uMgr.onlineUsers
}

// 根据id返回在线用户
func (uMgr *UserMgr) GetOnlineUserById(userID int) (up *UserProcess, err error) {
	// 待检测的 从map中获取数据
	up, ok := uMgr.onlineUsers[userID]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userID)
	}
	return
}
