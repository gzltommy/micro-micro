package handler

import (
	"github.com/gzltommy/user/domain/model"
	"github.com/gzltommy/user/domain/service"
	"github.com/gzltommy/user/proto/user"
	"context"
)

type User struct {
	UserDataService service.IUserDataService
}

// 注册
func (u User) Register(ctx context.Context, req *user.UserRegisterRequest, res *user.UserRegisterResponse) error {
	user := &model.User{
		//ID:           ,
		UserName:     req.UserName,
		FirstName:    req.FirstName,
		HashPassword: req.Pwd,
	}

	_, err := u.UserDataService.AddUser(user)
	if err != nil {
		return err
	}
	res.Message = "添加成功"
	return nil
}

// 登录
func (u User) Login(ctx context.Context, req *user.UserLoginRequest, res *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(req.UserName, req.Pwd)
	if err != nil {
		return err
	}
	res.IsSuccess = isOk
	return nil
}

// 获取用户信息
func (u User) GetUserInfo(ctx context.Context, req *user.UserInfoRequest, res *user.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(req.UserName)
	if err != nil {
		return err
	}
	res = UserForResponse(userInfo)
	return nil
}

// 类型转换
func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
