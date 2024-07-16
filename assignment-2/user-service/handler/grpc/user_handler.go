package grpc

import (
	"context"
	"fmt"
	"golang-advance/assignment-2/user-service/entity"
	pb "golang-advance/assignment-2/user-service/proto/user_service/v1"
	"golang-advance/assignment-2/user-service/service"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetUsersResponse, error) {
	users, err := u.userService.GetAllUsers(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var usersProto []*pb.User
	for _, user := range users {
		usersProto = append(usersProto, &pb.User{
			Id:          int32(user.ID),
			Name:        user.Name,
			Email:       user.Email,
			MobilePhone: user.Mobilephone,
			CreatedAt:   timestamppb.New(user.CreatedAt),
			UpdatedAt:   timestamppb.New(user.UpdatedAt),
		})
	}

	return &pb.GetUsersResponse{
		Users: usersProto,
	}, nil
}
func (u *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := u.userService.GetUserByID(ctx, int(req.GetId()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res := &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:          int32(user.ID),
			Name:        user.Name,
			Email:       user.Email,
			MobilePhone: user.Mobilephone,
			CreatedAt:   timestamppb.New(user.CreatedAt),
			UpdatedAt:   timestamppb.New(user.UpdatedAt),
		},
	}
	return res, nil
}

func (u *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.MutationResponse, error) {
	createdUser, err := u.userService.CreateUser(ctx, &entity.User{
		Name:        req.GetName(),
		Email:       req.GetEmail(),
		Mobilephone: req.GetMobilePhone(),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success created user with ID %d", createdUser.ID),
	}, nil
}
func (u *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.MutationResponse, error) {
	user := entity.User{
		ID:    int(req.GetId()),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	_, err := u.userService.UpdateUser(ctx, user.ID, user)
	if err != nil {
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success update user with ID %d", req.GetId()),
	}, nil
}
func (u *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.MutationResponse, error) {
	if err := u.userService.DeleteUser(ctx, int(req.GetId())); err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success deleted user with ID %d", req.GetId()),
	}, nil
}
