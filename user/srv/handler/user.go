package handler

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"user/pkg/model"
	db "user/pkg/mysql"

	"github.com/micro/go-micro/v2/util/log"

	pb "user/proto/user"
)

type UserSrv struct{}

func (e *UserSrv) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	var user model.User

	err := db.DB.Where("account = ?", req.Account).Take(&user).Error

	if err != nil {
		log.Error("查找用户错误:", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		log.Error("CompareHashAndPassword error:", err)
		return err
	}



	rsp.User = &pb.UserResponse{
		Id:                  int64(user.ID),
		Name:                 user.Name,
		Gender:               int64(user.Gender),
		CreateAt:             user.CreatedAt.Format("2006-01-02 15:04:05"),
		Avatar:               user.Avatar,
		Emial:                user.Email,
	}


	return nil
}

func (e *UserSrv) User(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	var user model.User
	err := db.DB.Take(&user, req.Id).Error
	if err != nil {
		return err
	}

	rsp.Id = int64(user.ID)
	rsp.Name = user.Name
	rsp.Emial = user.Email
	rsp.Avatar = user.Avatar
	rsp.Gender = int64(user.Gender)
	rsp.CreateAt = user.CreatedAt.Format("2006-01-02 15:04:05")


	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *UserSrv) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.UserSrv_StreamStream) error {
	log.Logf("Received User.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *UserSrv) PingPong(ctx context.Context, stream pb.UserSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
