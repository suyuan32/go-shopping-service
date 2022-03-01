//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"io"
	"sys_srv/global"
	"sys_srv/model"
	"sys_srv/proto"
	"time"
)

func UserModelToResponse(m model.User) *proto.UserInfoResponse {
	return &proto.UserInfoResponse{
		Id:           uint64(m.ID),
		NickName:     m.NickName,
		LoginName:    m.LoginName,
		Email:        m.Email,
		Mobile:       m.Mobile,
		Pic:          m.Pic,
		Status:       m.Status,
		LastLoginAt:  m.LastLoginAt.Unix(),
		LastLogin_IP: m.LastLoginIP,
	}
}

func (*SystemServer) GetUserList(ctx context.Context, page *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Scopes(Paginate(page.PageNum, page.PageSize)).Find(&users)
	if result.Error != nil {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}
	var response proto.UserListResponse
	response.Total = result.RowsAffected

	for _, value := range users {
		response.Data = append(response.Data, UserModelToResponse(value))
	}

	return &response, nil
}

func (*SystemServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.GetMobile()}).First(&user)
	if result.Error != nil {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}
	return UserModelToResponse(user), nil
}

func (*SystemServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.GetId())
	if result.Error != nil {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, global.GetMessage("UserNotFound"))
	}

	return UserModelToResponse(user), nil
}

func (*SystemServer) GetUserByLoginName(ctx context.Context, req *proto.BaseRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{LoginName: req.GetMsg()}).First(&user)
	if result.Error != nil {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, global.GetMessage("UserNotFound"))
	}

	return UserModelToResponse(user), nil
}

func (*SystemServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.BaseResponse, error) {
	user := &model.User{
		Model:         gorm.Model{},
		NickName:      *req.NickName,
		LoginName:     req.LoginName,
		Email:         *req.Email,
		Mobile:        *req.Mobile,
		LoginPassword: req.Password,
		Pic:           "",
		Status:        0,
		LastLoginAt:   time.Time{},
		LastLoginIP:   "",
		RoleID:        1,
		Role:          model.Role{},
	}

	var check model.User
	if res := global.DB.Where(&model.User{LoginName: req.LoginName}).Find(&check); res.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, global.GetMessage("UserAlreadyExist"))
	}

	result := global.DB.Create(&user)
	if result.Error != nil {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}

	if result.RowsAffected == 1 {
		return &proto.BaseResponse{
			Id:         uint64(user.ID),
			AffectRows: &result.RowsAffected,
		}, nil
	} else {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}
}

func (*SystemServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*proto.BaseResponse, error) {
	var passEncrypt string
	var err error
	if req.Password != "" {
		passEncrypt, err = passwordEncrypt(req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, global.GetMessage("FailMD5"))
		}
		passEncrypt = fmt.Sprintf("%s$%s", global.ServerConfig.PasswordEncrypt, passEncrypt)
	} else {
		passEncrypt = ""
	}
	user := &model.User{
		NickName:      req.NickName,
		LoginName:     req.LoginName,
		Email:         req.Email,
		Mobile:        req.Mobile,
		LoginPassword: passEncrypt,
		Pic:           req.Pic,
		Status:        req.Status,
		LastLoginAt:   time.Unix(req.LastLoginAt, 0),
		LastLoginIP:   req.LastLogin_IP,
		RoleID:        uint(req.RoleId),
	}

	result := global.DB.Save(&user)
	if result.Error != nil {
		msg := global.GetMessage("GormError") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Internal, msg)
	}

	if result.RowsAffected == 1 {
		return &proto.BaseResponse{
			Id:         uint64(user.ID),
			AffectRows: &result.RowsAffected,
		}, nil
	} else {
		msg := global.GetMessage("FailUpdate") + " err:" + result.Error.Error()
		zap.S().Error(msg)
		return nil, status.Errorf(codes.Unimplemented, msg)
	}
}

func (*SystemServer) CheckPassword(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	tmp, err := passwordEncrypt(req.Password)
	if tmp == req.EncryptedPassword {
		return &proto.CheckResponse{Success: true}, nil
	} else {
		if err != nil {
			return &proto.CheckResponse{Success: false}, status.Error(codes.Internal, err.Error())
		} else {
			return &proto.CheckResponse{Success: false}, status.Error(codes.OK, global.GetMessage("WrongPassword"))
		}
	}
}

// use md5 to encrypt password

func passwordEncrypt(origin string) (string, error) {
	hash := md5.New()
	_, _ = io.WriteString(hash, global.ServerConfig.PasswordEncrypt)
	_, err := io.WriteString(hash, origin)
	if err != nil {
		zap.S().Error(global.GetMessage("FailMD5"))
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
