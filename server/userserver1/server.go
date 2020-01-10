package userserver1

import (
	"context"
	"fmt"
	"sniper/dao/user"
	pb "sniper/rpc/user/v1"
)

// Server 实现/twirp/user.v1.User接口
type Server struct {
	Dao *user.Dao
}

// Register 注册账户
func (s *Server) Register(ctx context.Context, req *pb.RegistReq) (*pb.RegistResp, error) {
	// 查询用户名是否已存在
	userInfo, err := s.Dao.GetUserByName(ctx, req.GetAccount())
	if err != nil {
		return &pb.RegistResp{
			Code: 1,
			Msg:  "query user info by name error",
		}, err
	}
	if userInfo != nil {
		return &pb.RegistResp{
			Code: 1,
			Msg:  "duplicate account",
		}, nil
	}
	// 账号写入DB
	err = s.Dao.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.RegistResp{
		Code: 0,
		Msg:  fmt.Sprintf("success for %s", req.GetAccount()),
	}, nil
}
