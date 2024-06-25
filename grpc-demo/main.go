package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-demo/proto"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users  map[int32]*pb.GetUserResponse
	nextID int32
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.nextID++
	user := &pb.GetUserResponse{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	s.users[s.nextID] = user
	return &pb.CreateUserResponse{Id: s.nextID}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return &pb.UpdateUserResponse{Success: false}, status.Errorf(codes.NotFound, "user not found")
	}
	user.Name = req.Name
	user.Email = req.Email
	return &pb.UpdateUserResponse{Success: true}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, exists := s.users[req.Id]
	if !exists {
		return &pb.DeleteUserResponse{Success: false}, status.Errorf(codes.NotFound, "user not found")
	}
	delete(s.users, req.Id)
	return &pb.DeleteUserResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{
		users:  make(map[int32]*pb.GetUserResponse),
		nextID: 0,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
