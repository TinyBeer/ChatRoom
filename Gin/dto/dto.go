package dto

import "ChatRoom/Gin/model"

func ToUserInfo(user *model.User) *model.UserInfo {
	return &model.UserInfo{
		ID:   user.ID,
		Name: user.Name,
	}
}
