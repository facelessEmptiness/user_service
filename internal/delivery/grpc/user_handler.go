package grpc

import (
	"context"
	"github.com/facelessEmptiness/user_service/internal/domain"
	"github.com/facelessEmptiness/user_service/internal/usecase"
	pb "github.com/facelessEmptiness/user_service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // на практике — хешируй!
	}
	id, err := h.uc.Register(user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: id, Message: "User registered successfully"}, nil
}

func (h *UserHandler) AuthenticateUser(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	user, err := h.uc.GetByEmail(req.Email)
	if err != nil || user.Password != req.Password {
		return nil, err // упростим: без токенов пока
	}
	return &pb.AuthResponse{Token: "mock-token"}, nil
}

func (h *UserHandler) GetUserProfile(ctx context.Context, req *pb.UserID) (*pb.UserProfile, error) {
	user, err := h.uc.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserProfile{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
