package userserver1

import (
	"context"
	"fmt"
	pb "sniper/rpc/user/v1"
)

// Server 实现/twirp/user.v1.User接口
type Server struct{}

// Register 注册账户
func (s *Server) Register(ctx context.Context, req *pb.RegistReq) (*pb.RegistResp, error) {
	// todo
	return &pb.RegistResp{
		Code: 0, 
		Msg: fmt.Sprintf("success for %s", req.Account),
		}, nil
}
